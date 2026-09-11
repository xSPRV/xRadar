package analyzer

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

type CallSite struct {
	Receiver string
	Method   string
}

func (e *ReachabilityEngine) ExtractCalls(rootNode *sitter.Node, sourceCode []byte) ([]CallSite, error) {
	queryString := `
		(call_expression
			function: (member_expression
				object: (identifier) @receiver
				property: (property_identifier) @method
			)
		)
		(call_expression
			function: (identifier) @func_name
		)
	`

	lang := javascript.GetLanguage()
	query, err := sitter.NewQuery([]byte(queryString), lang)
	if err != nil {
		return nil, err
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(query, rootNode)

	var calls []CallSite

	for {
		match, ok := qc.NextMatch()
		if !ok {
			break
		}

		var site CallSite
		for _, capture := range match.Captures {
			captureName := query.CaptureNameForId(capture.Index)
			content := capture.Node.Content(sourceCode)

			switch captureName {
			case "receiver":
				site.Receiver = content
			case "method":
				site.Method = content
			case "func_name":
				site.Method = content
			}
		}
		if site.Method != "" {
			calls = append(calls, site)
		}
	}

	return calls, nil
}
