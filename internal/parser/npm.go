package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/xSPRV/xRadar/internal/graph"
)

type NPMLockfile struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Packages map[string]struct {
		Version      string            `json:"version"`
		Dependencies map[string]string `json:"dependencies"`
	} `json:"packages"`
}

type NPMParser struct{}

func (p *NPMParser) Parse(filePath string, g *graph.Graph) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read lockfile: %w", err)
	}

	var lockfile NPMLockfile
	if err := json.Unmarshal(data, &lockfile); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	for path, pkgData := range lockfile.Packages {
		if path == "" {
			continue
		}

		parts := strings.Split(path, "node_modules/")
		pkgName := parts[len(parts)-1]

		node := &graph.DependencyNode{
			ID:        fmt.Sprintf("npm:%s:%s", pkgName, pkgData.Version),
			Ecosystem: graph.EcosystemNPM,
			Name:      pkgName,
			Version:   pkgData.Version,
			Direct:    !strings.Contains(path, "node_modules/"),
		}
		g.AddNode(node)
	}

	// Note: A complete implementation requires a second pass here to map the AddEdge
	// relationships based on the 'dependencies' map in each package.

	return nil
}
