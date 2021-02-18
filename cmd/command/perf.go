package command

import (
	"github.com/goutils/pkg/performance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var perfConfig performance.Config

var cmdPERF = &cobra.Command{
	Use:   "perf",
	Short: "performance test",
	Long:  "performance testing",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&perfConfig); err != nil {
			log.Fatal("Unmarshal config file error:", err)
		}
		performance.TestRunner(&perfConfig)
	},
}
