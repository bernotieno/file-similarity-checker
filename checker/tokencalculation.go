package checker

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// calculateSmartTokenSimilarity computes similarity between tokenized Go source code
func calculateSmartTokenSimilarity(code1, code2 string) float64 {
	tokens1, _ := smartTokenizeGo(code1)
	tokens2, _ := smartTokenizeGo(code2)

	set1 := make(map[string]struct{})
	for _, t := range tokens1 {
		set1[t] = struct{}{}
	}
	set2 := make(map[string]struct{})
	for _, t := range tokens2 {
		set2[t] = struct{}{}
	}

	intersection := 0
	for token := range set1 {
		if _, exists := set2[token]; exists {
			intersection++
		}
	}

	union := make(map[string]struct{})
	for token := range set1 {
		union[token] = struct{}{}
	}
	for token := range set2 {
		union[token] = struct{}{}
	}

	if len(union) == 0 {
		return 0
	}

	return float64(intersection) / float64(len(union)) * 100
}

// smartTokenizeGo returns a normalized list of Go code tokens
func smartTokenizeGo(code string) ([]string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	var tokens []string
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			tokens = append(tokens, "_id")
		case *ast.BasicLit:
			tokens = append(tokens, "_val")
		case *ast.BinaryExpr:
			tokens = append(tokens, x.Op.String())
		case *ast.UnaryExpr:
			tokens = append(tokens, x.Op.String())
		case *ast.AssignStmt:
			tokens = append(tokens, "=")
		case *ast.IfStmt:
			tokens = append(tokens, "if")
		case *ast.ForStmt:
			tokens = append(tokens, "for")
		case *ast.ReturnStmt:
			tokens = append(tokens, "return")
		case *ast.CallExpr:
			tokens = append(tokens, "call")
		}
		return true
	})

	return tokens, nil
}
