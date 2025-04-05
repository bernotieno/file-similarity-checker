package checker

import (
	"os"
	"strings"
)

// readFileContent reads entire file content
func readFileContent(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// calculateSimilarity computes similarity between two files
func calculateSimilarity(content1, content2 string) float64 {
	if len(content1) > 10000 {
		content1 = content1[:10000]
	}
	if len(content2) > 10000 {
		content2 = content2[:10000]
	}

	lines1 := strings.Split(content1, "\n")
	lines2 := strings.Split(content2, "\n")

	matchingLines := 0
	for _, line1 := range lines1 {
		for _, line2 := range lines2 {
			if strings.TrimSpace(line1) == strings.TrimSpace(line2) {
				matchingLines++
				break
			}
		}
	}

	totalLines := len(lines1) + len(lines2)
	if totalLines == 0 {
		return 0
	}

	return (2.0 * float64(matchingLines)) / float64(totalLines) * 100
}

// categorizeSimilarity determines similarity category
func categorizeSimilarity(similarityScore float64) string {
	switch {
	case similarityScore > 70:
		return "Very Similar"
	case similarityScore > 30:
		return "Similar"
	default:
		return "Not Similar"
	}
}
