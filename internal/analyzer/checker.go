package analyzer

import (
	"context"
	"strings"
)

func (e *ReachabilityEngine) IsVulnerableFunctionReachable(ctx context.Context, sourceCode []byte, targetLib, vulnerableMethod string) (bool, error) {
	rootNode, err := e.ParseFile(ctx, sourceCode)
	if err != nil {
		return false, err
	}

	imports, err := e.ExtractImports(rootNode, sourceCode)
	if err != nil {
		return false, err
	}

	libraryImported := false
	for _, imp := range imports {
		if strings.Contains(imp, targetLib) {
			libraryImported = true
			break
		}
	}

	if !libraryImported {
		return false, nil
	}

	calls, err := e.ExtractCalls(rootNode, sourceCode)
	if err != nil {
		return false, err
	}

	for _, call := range calls {
		if call.Method == vulnerableMethod {
			return true, nil
		}
	}

	return false, nil
}
