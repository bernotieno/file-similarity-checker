package checker

import (
	"fmt"
	"os"
	"path/filepath"
)

// SimilarityResult stores comparison results between two files
type SimilarityResult struct {
	File1      string
	File2      string
	Similarity float64
	Category   string
}

// CodeSimilarityChecker manages file comparisons
type CodeSimilarityChecker struct {
	directory string
	files     []string
}

// New creates a new CodeSimilarityChecker
func New(directory string) (*CodeSimilarityChecker, error) {
	var err error
	if directory == "" {
		directory, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	} else {
		directory, err = filepath.Abs(directory)
		if err != nil {
			return nil, err
		}
	}

	checker := &CodeSimilarityChecker{directory: directory}
	err = checker.findFiles()
	if err != nil {
		return nil, err
	}
	return checker, nil
}

// findFiles identifies relevant code files
func (cc *CodeSimilarityChecker) findFiles() error {
	extensions := []string{".go"}

	err := filepath.Walk(cc.directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		for _, ext := range extensions {
			if filepath.Ext(path) == ext {
				cc.files = append(cc.files, path)
				break
			}
		}
		return nil
	})

	if len(cc.files) < 2 {
		return fmt.Errorf("need at least two files to compare")
	}

	if len(cc.files) > 50 {
		cc.files = cc.files[:50]
	}

	return err
}

// CompareFiles compares tokenized files
func (cc *CodeSimilarityChecker) CompareFiles() ([]SimilarityResult, error) {
	var results []SimilarityResult

	for i := 0; i < len(cc.files); i++ {
		for j := i + 1; j < len(cc.files); j++ {
			content1, err1 := os.ReadFile(cc.files[i])
			content2, err2 := os.ReadFile(cc.files[j])
			if err1 != nil || err2 != nil {
				continue
			}

			similarityScore := calculateSmartTokenSimilarity(string(content1), string(content2))
			category := categorizeSimilarity(similarityScore)

			result := SimilarityResult{
				File1:      filepath.Base(cc.files[i]),
				File2:      filepath.Base(cc.files[j]),
				Similarity: similarityScore,
				Category:   category,
			}
			results = append(results, result)
		}
	}

	return results, nil
}

// Directory returns the directory path
func (cc *CodeSimilarityChecker) Directory() string {
	return cc.directory
}

// Files returns the list of files
func (cc *CodeSimilarityChecker) Files() []string {
	return cc.files
}

// categorizeSimilarity categorizes the similarity score
func categorizeSimilarity(score float64) string {
	switch {
	case score >= 90:
		return "Very High"
	case score >= 70:
		return "High"
	case score >= 50:
		return "Moderate"
	case score >= 30:
		return "Low"
	default:
		return "Very Low"
	}
}
