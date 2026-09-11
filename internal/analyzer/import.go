package analyzer

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

func (e *ReachabilityEngine) ExtractImports(rootNode *sitter.Node, sourceCode []byte) ([]string, error) {
	queryString := `
		(import_statement source: (string (string_fragment) @import_name))
		(call_expression
			function: (identifier) @func_name (#eq? @func_name "require")
			arguments: (arguments (string (string_fragment) @import_name))
		)
	`

	lang := javascript.GetLanguage()
	query, err := sitter.NewQuery([]byte(queryString), lang)
	if err != nil {
		return nil, err
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(query, rootNode)

	var imports []string

	for {
		match, ok := qc.NextMatch()
		if !ok {
			break
		}
		for _, capture := range match.Captures {
			if query.CaptureNameForId(capture.Index) == "import_name" {
				importName := capture.Node.Content(sourceCode)
				imports = append(imports, importName)
			}
		}
	}

	return imports, nil
}
