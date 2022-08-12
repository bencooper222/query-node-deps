package process

import (
	"fmt"
	"log"

	"github.com/bencooper222/query-node-deps/pkg/db/types"
	gh "github.com/bencooper222/query-node-deps/pkg/github"
	yarn "github.com/bencooper222/query-node-deps/pkg/lockfiles"
	"github.com/bencooper222/query-node-deps/pkg/util"
	"github.com/google/go-github/v45/github"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ProcessLockAndPackageForLatestCommit(ghClient github.Client, db gorm.DB, org string, repo string) error {
	latestCommit, err := gh.GetLatestCommit(ghClient, org, repo, nil)
	if err != nil {
		return fmt.Errorf("error getting latest commit: %v", err)
	}
	err = ProcessLockfileForRepoLatestCommit(ghClient, db, org, repo, *latestCommit)
	if err != nil {
		return fmt.Errorf("error processing lockfile: %v", err)
	}
	err = ProcessPackageJSONForRepoLatestCommit(ghClient, db, org, repo, *latestCommit)
	if err != nil {
		return fmt.Errorf("error processing package.json: %v", err)
	}
	return nil
}

func ProcessPackageJSONForRepoLatestCommit(ghClient github.Client, db gorm.DB, org string, repo string, latestCommit github.RepositoryCommit) error {
	gitSlug := "gh/" + org + "/" + repo
	pjson, err := gh.GetStringifiedFileContents(ghClient, org, repo, "package.json", latestCommit.SHA)
	if err != nil {
		return err
	}
	deps := yarn.GetPackageJSONDependencies(pjson.Contents)
	packageDeps := []types.Dependency{}
	for key, value := range deps {
		packageDeps = append(packageDeps, types.Dependency{
			Fully_qualified_git_slug: gitSlug,
			Source:                   "PACKAGE_JSON",
			Relative_repo_file_path:  "/package.json",
			Name:                     key,
			Semver_version_spec:      value,
			Resolved_version:         types.Semver("0.0.0"),
			Sha:                      *latestCommit.SHA,
		})
	}
	db.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&packageDeps).Error
		return err
	})
	return nil
}

func ProcessLockfileForRepoLatestCommit(ghClient github.Client, db gorm.DB, org string, repo string, latestCommit github.RepositoryCommit) error {
	gitSlug := "gh/" + org + "/" + repo
	lockfileContents, err := gh.GetStringifiedFileContents(ghClient, org, repo, "yarn.lock", latestCommit.SHA)
	if err != nil {
		return err
	}

	rawParsedYarnDependencies := yarn.GetParsedYarnLockfileFromAlreadyStringifiedLockfile(lockfileContents.Contents)
	mappedYarnDependencies := util.MapSlice(rawParsedYarnDependencies, func(p yarn.Package) types.Dependency {

		return types.Dependency{
			Fully_qualified_git_slug: gitSlug,
			Source:                   "YARN_LOCK",
			Relative_repo_file_path:  "/yarn.lock",
			Name:                     p.Name,
			Semver_version_spec:      p.SemverVersionSpec,
			Resolved_version:         types.Semver(p.ResolvedVersion),
			Sha:                      *latestCommit.SHA,
		}
	})
	if len(mappedYarnDependencies) == 0 {
		log.Printf("no dependencies found in yarn.lock for %s/%s", org, repo)
		return nil
	} else {
		log.Printf("found %d dependencies in yarn.lock for %s/%s", len(mappedYarnDependencies), org, repo)
	}

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&types.Repository{
			Git_svc: "gh",
			Org:     org,
			Repo:    repo,
		}).Error; err != nil {
			return err
		}

		// log.Println(latestCommit.Commit.Committer.Date)
		if err := tx.Create(&types.Commit{
			Fully_qualified_git_slug: gitSlug,
			Sha:                      *latestCommit.SHA,
			Datetime:                 *latestCommit.Commit.Committer.Date,
		}).Error; err != nil {
			log.Println("If this was a conflict, make sure to check if the commit is already in the DB", org, repo, lockfileContents.SHA)
			return err
		}

		if err := tx.Create(mappedYarnDependencies).Error; err != nil {
			return err
		}

		return nil
	})
	return nil

}
