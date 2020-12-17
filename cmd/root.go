package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"wulfheartalexander/logging"
)

var (
	variant string
	logger *zap.Logger
	debug bool
)

var rootCmd = &cobra.Command{
	Use:   "alex",
	Short: "alex is a cli for adjudicating diplomacy games",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ! debug {
			logger, err := zap.NewProduction()
			if err != nil {
				panic(err)
			}
			logging.Instantiate(logger)
		} else {
			logger, err := zap.NewDevelopment()
			if err != nil {
				panic(err)
			}
			logging.Instantiate(logger)
		}
	},
	Version: "0.3.0",
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug,"debug",false, "Print some debug messages")
}
