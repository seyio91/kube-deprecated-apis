package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	K8sVersion     string `env:"K8S_VERSION,default=1.26"`
	K8sClusterName string `env:"K8S_CLUSTER_NAME,required"`
	CronInterval   string `env:"METRICS_CRON_INTERVAL,default=@every 1h"`
	Port           string `env:"PORT,default=2112"`
}

func New() (*Config, error) {
	var c Config
	if err := envconfig.Process(context.TODO(), &c); err != nil {
		return nil, fmt.Errorf("failed to process config: %v", err)
	}

	return &c, nil
}
