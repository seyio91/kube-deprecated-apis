package models

type PlutoOutput struct {
	DeprecatedAPIs []struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		API       struct {
			Version                string `json:"version"`
			Kind                   string `json:"kind"`
			DeprecatedIn           string `json:"deprecated-in"`
			RemovedIn              string `json:"removed-in"`
			ReplacementApi         string `json:"replacement-api"`
			ReplacementAvailableIn string `json:"replacement-available-in"`
			Component              string `json:"component"`
		} `json:"api"`
		Deprecated           bool `json:"deprecated"`
		Removed              bool `json:"removed"`
		ReplacementAvailable bool `json:"replacementAvailable"`
	} `json:"items"`
}