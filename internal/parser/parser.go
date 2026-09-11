package parser

import "github.com/xSPRV/xRadar/internal/graph"

type LockfileParser interface {
	Parse(filePath string, g *graph.Graph) error
}
