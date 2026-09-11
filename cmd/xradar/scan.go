package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var targetDir string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a directory for dependencies and vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting xRadar scan on directory: %s\n", targetDir)

		// TODO: Call internal/parser and internal/graph logic here
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&targetDir, "dir", "d", ".", "Target directory to scan")
}
