package metrics

import (
	"strconv"

	"github.com/seyio91/kube-deprecated-apis/pkg/models"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	deprecatedAPIs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kubernetes_deprecated_apis",
			Help: "Count of Kubernetes resources using deprecated APIs.",
		},
		[]string{"cluster", "name", "namespace", "apiVersion", "kind", "deprecated", "removed", "replacement_api"},
	)
)

func init() {
	// Register the metric with Prometheus's default registry
	prometheus.MustRegister(deprecatedAPIs)
}

func UpdateDeprecatedAPIMetrics(result *models.PlutoOutput, cluster string) error {
	for _, api := range result.DeprecatedAPIs {
		deprecatedAPIs.With(prometheus.Labels{"cluster": cluster, "name": api.Name, "namespace": api.Namespace, "apiVersion": api.API.Version, "kind": api.API.Kind, "deprecated": strconv.FormatBool(api.Deprecated), "removed": strconv.FormatBool(api.Removed), "replacement_api": strconv.FormatBool(api.ReplacementAvailable)}).Set(1)
	}
	return nil
}
