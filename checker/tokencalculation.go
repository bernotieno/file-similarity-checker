package checker

import (
	"regexp"
	"strings"
)

// calculateTokenSimilarity compares two code files by token set similarity
func calculateTokenSimilarity(code1, code2 string) float64 {
	tokens1 := genericTokenize(code1)
	tokens2 := genericTokenize(code2)

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
		if _, found := set2[token]; found {
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

// genericTokenize performs language-agnostic tokenization
func genericTokenize(code string) []string {
	code = removeComments(code)

	// Replace strings and numbers with generic tokens
	code = regexp.MustCompile(`"[^"]*"|'[^']*'|`+"`[^`]*`").ReplaceAllString(code, "_str")
	code = regexp.MustCompile(`\b\d+(\.\d+)?\b`).ReplaceAllString(code, "_num")

	// Split by word boundaries and symbols
	rawTokens := regexp.MustCompile(`[A-Za-z_]\w*|\S`).FindAllString(code, -1)

	// Normalize identifiers (optional): Replace all variable names with _id
	keywords := map[string]struct{}{
		"if": {}, "else": {}, "for": {}, "while": {}, "return": {}, "switch": {}, "case": {}, "func": {},
		"var": {}, "let": {}, "const": {}, "class": {}, "struct": {}, "import": {}, "package": {}, "public": {},
		"private": {}, "protected": {}, "def": {}, "end": {}, "do": {}, "try": {}, "catch": {}, "finally": {},
	}

	var tokens []string
	for _, tok := range rawTokens {
		lower := strings.ToLower(tok)
		if _, isKeyword := keywords[lower]; isKeyword {
			tokens = append(tokens, lower)
		} else if regexp.MustCompile(`^[A-Za-z_]\w*$`).MatchString(tok) {
			tokens = append(tokens, "_id")
		} else {
			tokens = append(tokens, tok)
		}
	}

	return tokens
}

// removeComments strips simple single-line and multi-line comments
func removeComments(code string) string {
	// C/Java/JS style comments
	singleLine := regexp.MustCompile(`(?m)//.*$`)
	multiLine := regexp.MustCompile(`(?s)/\*.*?\*/`)

	// Python/Ruby/etc style comments
	pythonLine := regexp.MustCompile(`(?m)#.*$`)

	code = singleLine.ReplaceAllString(code, "")
	code = multiLine.ReplaceAllString(code, "")
	code = pythonLine.ReplaceAllString(code, "")

	return code
}
