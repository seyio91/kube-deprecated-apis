# Kube Deprecated API
The kube-deprecated-api tool is a script that run on each kubernetes cluster to expose the deprecated API's as prometheus metrics.
The metric exposed is `kubernetes_deprecated_apis` gauge metrics, which can be used to identify API's that need upgrading before performing Kubernetes Upgrade.

### Prerequisites
- K8S_VERSION: (Optional) This is the target kubernetes version to be upgraded to. default is 1.26. Change to check compatible api's with other kubernetes versions
- K8S_CLUSTER_NAME: (Required) Identifier for each kubernetes cluster on NewRelic.
- METRICS_CRON_INTERVAL: (Optional) cron option for querying the cluster for deprecated api versions. current default is `@every 1h`