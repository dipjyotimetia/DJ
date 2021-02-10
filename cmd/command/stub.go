package command

import (
	"github.com/goutils/pkg/stubs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var stubConfig stubs.Config

var cmdSTUBS = &cobra.Command{
	Use:   "stubs",
	Short: "stub server",
	Long:  `stubs server:`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&stubConfig); err != nil {
			log.Fatal("Unmarshal config file error:", err)
		}

		if err := stubs.ListenAndServe(&stubConfig); err != nil {
			log.Fatal(err)
		}
	},
}
