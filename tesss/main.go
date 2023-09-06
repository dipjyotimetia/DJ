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
	// _ = ListIssues("", "dipjyotimetia", "DJ")

	// fmt.Println(pp)
	CreateIssues("TOKEN", "dipjyotimetia", "DJ")

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
	Author    string
}

func CreateIssues(token string, owner string, repos string) {
	ctx := context.Background()
	client := AuthGithubAPI(ctx, token)
	issues, _, err := client.Issues.Create(ctx, owner, repos, &github.IssueRequest{
		Title:  Ptr("New Tracker"),
		Labels: PtrArr([]string{"storage"}),
		Body:   Ptr("New test trackers"),
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(issues)
}

// Ptr returns a pointer to the provided value.
func Ptr[T any](val T) *T {
	return &val
}

func PtrArr[T any](val []T) *[]T {
	return &val
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
			Author:    v.GetAuthorAssociation(),
		})
	}
	fmt.Println(issueList[0])
	return issueList
}
