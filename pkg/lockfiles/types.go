package yarn

type Package struct {
	Name              string
	SemverVersionSpec string
	ResolvedVersion   string
}

type JsonSchema struct {
	Type string `json:"type"`

	Object map[string]struct {
		Version      string            `json:"version"`
		Resolved     string            `json:"resolved"`
		Integrity    string            `json:"integrity"`
		Dependencies map[string]string `json:"dependencies"`
	} `json:"object"`
}
