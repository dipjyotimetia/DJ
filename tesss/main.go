package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	pp := ListIssues("", "", "")

	fmt.Println(pp)

}

// AuthGithubAPI authentication of github api
func AuthGithubAPI(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

type Issues struct {
	ID        int64
	Title     string
	State     string
	CreatedAt time.Time
	URL       string
	Body      string
}

// ListIssues get list of issues
func ListIssues(token string, owner string, repos string) interface{} {
	ctx := context.Background()
	client := AuthGithubAPI(ctx, token)
	issues, _, err := client.Issues.ListByRepo(ctx, owner, repos, nil)
	if err != nil {
		log.Println(err)
	}

	var issueList []interface{}
	for _, v := range issues {
		issueList = append(issueList, &Issues{
			ID:        v.GetID(),
			Title:     v.GetTitle(),
			State:     v.GetState(),
			CreatedAt: v.GetCreatedAt().Time,
			URL:       v.GetHTMLURL(),
			Body:      v.GetBody(),
		})
	}
	fmt.Println(issueList...)
	return issueList
}
