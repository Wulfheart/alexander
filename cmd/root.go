package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	variant string
)

var rootCmd = &cobra.Command{
	Use:   "alex",
	Short: "alex is a cli for adjudicating diplomacy games",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
}
