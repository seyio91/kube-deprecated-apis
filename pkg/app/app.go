package app

import (
	"net/http"

	"github.com/seyio91/kube-deprecated-apis/pkg/commands"
	"github.com/seyio91/kube-deprecated-apis/pkg/config"
	"github.com/seyio91/kube-deprecated-apis/pkg/logger"
	"github.com/seyio91/kube-deprecated-apis/pkg/metrics"
	"github.com/seyio91/kube-deprecated-apis/pkg/models"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger         *logrus.Logger
	Config         *config.Config
	DeprecatedAPIs *models.PlutoOutput
}

func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	logger := logger.GetLogger()

	app := &App{
		Logger: logger,
		Config: cfg,
	}

	if app.DeprecatedAPIs, err = commands.RunPlutoExec(cfg.K8sVersion); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	// initial
	if err := metrics.UpdateDeprecatedAPIMetrics(a.DeprecatedAPIs, a.Config.K8sClusterName); err != nil {
		a.Logger.Errorf("Error updating metrics:", err)
		return err
	}

	c := cron.New()

	_, err := c.AddFunc(a.Config.CronInterval, func() {
		// Update metrics with current data
		a.Logger.Info("collecting pluto metrics")
		if err := metrics.UpdateDeprecatedAPIMetrics(a.DeprecatedAPIs, a.Config.K8sClusterName); err != nil {
			a.Logger.Errorf("Error updating metrics:", err)
			return
		}
	})

	if err != nil {
		a.Logger.Errorf("Error scheduling the job:", err)
		return err
	}

	c.Start()

	defer c.Stop()

	a.Logger.Info("Cron scheduler started.")

	// Metrics Handler
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	a.Logger.Infof("Starting HTTP server on :%s", a.Config.Port)

	if err := http.ListenAndServe(":"+a.Config.Port, nil); err != nil {
		a.Logger.Errorf("Error starting server: %s", err)
	}
	return nil
}
