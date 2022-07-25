package types

import "time"

type Repository struct {
	Git_svc                  string
	Org                      string
	Repo                     string
	Fully_qualified_git_slug string `gorm:"->;type:text GENERATED ALWAYS AS (git_svc || '/' || org || '/' || repo) STORED;primaryKey;"`
}

type Commit struct {
	Fully_qualified_git_slug string // implicit foreign key because I can't figure out how to get the below to work
	// Repository               Repository `gorm:"foreignKey:Fully_qualified_git_slug"`

	Sha      string
	Datetime time.Time
}

type Dependency struct {
	Fully_qualified_git_slug string // implicit foreign key because I can't figure out how to get the below to work
	// Repository               Repository `gorm:"foreignKey:Fully_qualified_git_slug"`

	Source                  string // PACKAGE_JSON or YARN_LOCK, not bothering with dealing with golang's shitty enum workarounds
	Relative_repo_file_path string // format: /x.a is a top-level file in the repo called "x.a"
	Name                    string
	Semver_version_spec     string
	Resolved_version        Semver
}
