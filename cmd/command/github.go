package command

import (
	"encoding/json"
	"github.com/goutils/pkg/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

var gitConfig github.Config

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Get list of issues",
	Long:  "Get the list of issues in the specified repo",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&gitConfig); err != nil {
			log.Fatal("Unmarshal config file error:", err)
		}
		var result []interface{}
		for _, repoName := range gitConfig.Projects {
			result = append(result, repoName, github.ListIssues(gitConfig.AuthToken, gitConfig.Owner, repoName))
		}

		file, _ := json.MarshalIndent(result, "", "")
		_ = ioutil.WriteFile("test.json", file, 0644)
	},
}
