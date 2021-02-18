package command

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var cfgFile string

var (
	rootCmd = &cobra.Command{
		Use:   "dj",
		Short: "Command line utility",
		Long:  "Single solution for common problems",
	}
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of github eyes",
	Long:  `All software has versions. This is dj`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DJ v0.1 -- HEAD")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "", "config file")
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	rootCmd.PersistentFlags().StringP("author", "a", "Dipjyoti Metia", "author name for copyright attribution")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(issueCmd)
	rootCmd.AddCommand(cmdSTUBS)
	rootCmd.AddCommand(cmdPERF)
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.AddConfigPath("../configs")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func err(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
