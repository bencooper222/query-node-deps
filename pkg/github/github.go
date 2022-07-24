package github

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"time"

	gh "github.com/google/go-github/v45/github"
)

type FileContentsRtn struct {
	SHA      string
	Contents string
}

func GetStringifiedFileContents(client gh.Client, org string, repo string, loc string, commit *string) (FileContentsRtn, error) {
	ctx := context.Background()
	fileContent, _, _, err := client.Repositories.GetContents(ctx, org, repo, loc, &gh.RepositoryContentGetOptions{
		Ref: *commit,
	})
	if err != nil {
		return FileContentsRtn{}, err
	}

	if fileContent == nil {
		return FileContentsRtn{}, errors.New("file content is nil")
	}

	decoded, err := base64.StdEncoding.DecodeString(*fileContent.Content)

	return FileContentsRtn{Contents: string(decoded), SHA: *fileContent.SHA}, nil
}

func GetCommitDatetime(client gh.Client, org string, repo string, ref string) (time.Time, error) {
	ctx := context.Background()
	refData, res, err := client.Repositories.GetCommit(ctx, org, repo, ref, nil)
	if err != nil {
		return time.Time{}, err
	}
	log.Println(res)
	return *refData.Commit.Committer.Date, nil
}

func GetLatestCommit(client gh.Client, org string, repo string, branch *string) (gh.RepositoryCommit, error) {
	if branch == nil {
		master := "master"
		branch = &master
	}
	ctx := context.Background()
	refData, _, err := client.Repositories.ListCommits(ctx, org, repo, &gh.CommitsListOptions{
		SHA: *branch,
		ListOptions: gh.ListOptions{
			PerPage: 1,
		},
	})

	return *refData[0], err
}
