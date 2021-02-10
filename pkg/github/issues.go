package github

import (
	"context"
	"log"
	"time"
)

type Issues struct {
	ID        int64
	Title     string
	State     string
	CreatedAt time.Time
	URL       string
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
			CreatedAt: v.GetCreatedAt(),
			URL:       v.GetHTMLURL(),
		})
	}
	return issueList
}
