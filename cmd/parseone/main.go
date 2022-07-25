package main

import (
	"context"

	"github.com/bencooper222/query-node-deps/pkg/db"
	"github.com/bencooper222/query-node-deps/pkg/env"
	"github.com/bencooper222/query-node-deps/pkg/process"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func main() {
	// log.Println("Hello, world!")
	// str := yarn.GetParsedYarnLockfile("./yarn.lock")
	// log.Println(str)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: env.Get("GITHUB_ACCESS_TOKEN", "not a token")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	pgclient, _ := db.BuildPostgresClient("localhost", "postgres", "password", "postgres", "5432")

	process.ProcessLockfileForRepoLatestCommit(*client, pgclient, "bencooper222", "hibp-bot")
}
