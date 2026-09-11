package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xSPRV/xRadar/internal/graph"
	"github.com/xSPRV/xRadar/internal/parser"
)

var targetDir string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a directory for dependencies and vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting xRadar scan on directory: %s\n", targetDir)

		g := graph.NewGraph()
		npmParser := &parser.NPMParser{}

		lockfilePath := filepath.Join(targetDir, "package-lock.json")

		err := npmParser.Parse(lockfilePath, g)
		if err != nil {
			fmt.Printf("Error parsing lockfile: %v\n", err)
			return
		}

		fmt.Printf("Successfully parsed graph!\n")
		fmt.Printf("Total Nodes: %d\n", len(g.Nodes))
		fmt.Printf("Direct Dependencies: %d\n", len(g.RootNodes))
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&targetDir, "dir", "d", ".", "Target directory to scan")
}
