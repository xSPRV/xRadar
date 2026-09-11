package analyzer

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

type ReachabilityEngine struct{}

func NewReachabilityEngine() *ReachabilityEngine {
	return &ReachabilityEngine{}
}

func (e *ReachabilityEngine) ParseFile(ctx context.Context, sourceCode []byte) (*sitter.Node, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(javascript.GetLanguage())

	tree, err := parser.ParseCtx(ctx, nil, sourceCode)
	if err != nil {
		return nil, fmt.Errorf("failed to parse source code: %w", err)
	}

	return tree.RootNode(), nil
}
