package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xradar",
	Short: "xRadar is a dependency reachability and health scanner",
	Long: `xRadar parses project dependencies, builds a directed acyclic graph (DAG), 
and uses Tree-sitter to perform reachability analysis on known vulnerabilities.`,
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
