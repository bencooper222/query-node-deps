package main

import (
	"log"

	gh "github.com/bencooper222/query-node-deps/pkg/github"
	"github.com/google/go-github/v45/github"
)

func main() {
	// log.Println("Hello, world!")
	// str := yarn.GetParsedYarnLockfile("./yarn.lock")
	// log.Println(str)

	client := github.NewClient(nil)

	latestCommit, err := gh.GetLatestCommit(*client, "bencooper222", "hibp-bot", nil)
	if err != nil {
		panic(err)
	}

	fileContents, err := gh.GetStringifiedFileContents(*client, "bencooper222", "hibp-bot", "yarn.lock", latestCommit.SHA)

	if err != nil {
		panic(err)
	}
	log.Println(fileContents.Contents)
	log.Println(*latestCommit.SHA)

}
