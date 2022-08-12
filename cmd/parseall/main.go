package main

import (
	"context"
	"log"
	"strings"

	"github.com/bencooper222/query-node-deps/pkg/db"
	"github.com/bencooper222/query-node-deps/pkg/env"
	gh "github.com/bencooper222/query-node-deps/pkg/github"
	"github.com/bencooper222/query-node-deps/pkg/process"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: env.Get("GITHUB_ACCESS_TOKEN", "not a token")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	pgclient, _ := db.BuildPostgresClient("localhost", "postgres", "password", "postgres", "5432")

	org := "convoyinc"

	log.Print("Finding all the repos in ", org)
	repos := gh.GetOrgRepos(*client, org)
	log.Printf("Found %d repos", len(repos))

	for i, repo := range repos {
		log.Printf("(%d/%d) Pulling package.json+yarn.lock info for %s/%s", i, len(repos), org, repo.GetName())

		err := process.ProcessLockAndPackageForLatestCommit(*client, pgclient, org, repo.GetName())
		if err != nil {
			if strings.Contains(err.Error(), "404 Not Found") {
				log.Printf("package.json or yarn.lock not found for %s/%s", org, repo.GetName())
			} else {
				log.Printf("Error processing %s/%s, continuing. %s", org, repo.GetName(), err)
			}
		} else {
			log.Printf("Processed %s/%s successfully", org, repo.GetName())
		}
	}
}
