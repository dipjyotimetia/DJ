package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	ListBatch("TOKEN", "dipjyotimetia", "DJ")

	// fmt.Println(pp)
	// CreateIssues(TOKEN, "dipjyotimetia", "DJ")

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
		Title:    Ptr("New Tracker"),
		Labels:   PtrArr([]string{"storage"}),
		Assignee: Ptr("dipjyotimetia"),
		Body:     Ptr(constructData([]string{"A1234567", "C2342343", "D6575665", "E2342356"})),
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(issues.GetNumber())
}

func constructData(data []string) string {
	return strings.Join(data, "\n")
}

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
	issues, _, err := client.Issues.ListByRepo(ctx, owner, repos, &github.IssueListByRepoOptions{
		Labels: []string{"storage"},
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(issues[0].GetBody())
	return ""
}

// ListIssues get list of issues
func ListBatch(token string, owner string, repos string) {
	var env string
	ctx := context.Background()
	client := AuthGithubAPI(ctx, token)
	issues, _, err := client.Issues.ListByRepo(ctx, owner, repos, &github.IssueListByRepoOptions{
		Labels: []string{"batch"},
	})
	if err != nil {
		log.Println(err)
	}
	for _, v := range issues {
		if v.GetTitle() == "BatchWindow" {
			switch strings.ToUpper(v.GetBody()) {
			case "SIT-L":
				env = "SIT-L"
			case "SIT-K":
				env = "SIT-K"
			}
		}
	}
	fmt.Println(env)
}
