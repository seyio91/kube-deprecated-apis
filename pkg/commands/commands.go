package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/seyio91/kube-deprecated-apis/pkg/models"
)

func RunPlutoExec(k8s_version string) (*models.PlutoOutput, error) {
	targetVersion := fmt.Sprintf("k8s=v%s", k8s_version)
	cmd := exec.Command("pluto", "detect-all-in-cluster", "-t", targetVersion, "-o", "json")

	ignoreDeprecations := "PLUTO_IGNORE_DEPRECATIONS=true"
	ignoreRemovals := "PLUTO_IGNORE_REMOVALS=true"
	newEnv := append(os.Environ(), ignoreDeprecations, ignoreRemovals)
	cmd.Env = newEnv
	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("failed to run pluto command: %v", err)
	}
	var result models.PlutoOutput
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
