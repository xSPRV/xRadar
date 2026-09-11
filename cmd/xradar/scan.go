package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xSPRV/xRadar/internal/analyzer"
	"github.com/xSPRV/xRadar/internal/graph"
	"github.com/xSPRV/xRadar/internal/parser"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

var targetDir string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a directory for dependencies and vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s[INFO]%s Starting xRadar scan on directory: %s\n", colorCyan, colorReset, targetDir)

		g := graph.NewGraph()
		npmParser := &parser.NPMParser{}
		lockfilePath := filepath.Join(targetDir, "package-lock.json")

		if err := npmParser.Parse(lockfilePath, g); err != nil {
			fmt.Printf("%s[WARN]%s Could not parse package-lock.json: %v\n", colorYellow, colorReset, err)
		} else {
			fmt.Printf("%s[INFO]%s Parsed %d dependencies from lockfile.\n", colorCyan, colorReset, len(g.Nodes))
		}

		engine := analyzer.NewReachabilityEngine()
		ctx := context.Background()

		targetLib := "lodash"
		vulnerableMethod := "template"

		fmt.Printf("%s[INFO]%s Scanning for reachability of %s.%s()...\n", colorCyan, colorReset, targetLib, vulnerableMethod)

		reachableFiles := []string{}

		err := filepath.WalkDir(targetDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() && d.Name() == "node_modules" {
				return filepath.SkipDir
			}
			if !d.IsDir() && strings.HasSuffix(d.Name(), ".js") {
				sourceCode, err := os.ReadFile(path)
				if err != nil {
					return nil
				}

				isReachable, err := engine.IsVulnerableFunctionReachable(ctx, sourceCode, targetLib, vulnerableMethod)
				if err == nil && isReachable {
					reachableFiles = append(reachableFiles, path)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("%s[ERROR]%s Error walking directory: %v\n", colorRed, colorReset, err)
			return
		}

		if len(reachableFiles) > 0 {
			fmt.Printf("%s[CRITICAL] Vulnerability reachable! Found in %d files:%s\n", colorRed, len(reachableFiles), colorReset)
			for _, f := range reachableFiles {
				fmt.Printf("  - %s\n", f)
			}
		} else {
			fmt.Printf("%s[OK] No reachable paths found for this vulnerability.%s\n", colorGreen, colorReset)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&targetDir, "dir", "d", ".", "Target directory to scan")
}
