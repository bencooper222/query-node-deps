package yarn

type Package struct {
	Name              string
	SemverVersionSpec string
	ResolvedVersion   string
}

type YarnLockJsonSchema struct {
	Type string `json:"type"`

	Object map[string]struct {
		Version      string            `json:"version"`
		Resolved     string            `json:"resolved"`
		Integrity    string            `json:"integrity"`
		Dependencies map[string]string `json:"dependencies"`
	} `json:"object"`
}

type YarnLockBerryJsonSchema map[string]interface{}

type LockPackage struct {
	Version      string            `json:"version"`
	Resolution   string            `json:"resolution,omitempty"`
	Dependencies map[string]string `json:"dependencies,omitempty"`
	Checksum     string            `json:"checksum"`
	LanguageName string            `json:"languageName"`
	LinkType     string            `json:"linkType"`
}

// This was close to working and would be safer than assertion:
// type YarnLockBerryJsonSchema struct {
// 	Metadata struct {
// 		Version  int `json:"version"`
// 		CacheKey int `json:"cacheKey"`
// 	} `json:"__metadata"`
// 	Object map[string]LockPackage `json:"-"`
// }
