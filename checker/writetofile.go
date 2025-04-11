package checker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/jung-kurt/gofpdf"
)

// WriteResultsToFile saves comparison results in different formats: text, markdown, or html
func WriteResultsToFile(results []SimilarityResult, outputFile, format string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	switch strings.ToLower(format) {
	case "html":
		writer.WriteString(`<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Code Similarity Report</title>
	<style>
		body { font-family: Arial, sans-serif; }
		table { width: 100%; border-collapse: collapse; margin-top: 20px; }
		th, td { border: 1px solid #ccc; padding: 8px; text-align: left; }
		th { background-color: #f4f4f4; }
		tr:nth-child(even) { background-color: #f9f9f9; }
	</style>
</head>
<body>
	<h2>Code Similarity Report</h2>
	<table>
		<tr><th>File 1</th><th>File 2</th><th>Similarity (%)</th><th>Category</th></tr>
`)
		for _, result := range results {
			row := fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%.2f</td><td>%s</td></tr>\n",
				result.File1, result.File2, result.Similarity, result.Category)
			writer.WriteString(row)
		}
		writer.WriteString(`	</table>
</body>
</html>`)
	case "pdf":
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Code Similarity Report")
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(50, 10, "File 1", "1", 0, "", false, 0, "")
	pdf.CellFormat(50, 10, "File 2", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Similarity %", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Category", "1", 1, "", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	for _, result := range results {
		pdf.CellFormat(50, 10, result.File1, "1", 0, "", false, 0, "")
		pdf.CellFormat(50, 10, result.File2, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", result.Similarity), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, result.Category, "1", 1, "", false, 0, "")
	}

	err := pdf.OutputFileAndClose(outputFile)
	return err

	default: 
		header := fmt.Sprintf("%-30s %-30s %-15s %-15s\n", "File 1", "File 2", "Similarity %", "Category")
		writer.WriteString(header)
		writer.WriteString(strings.Repeat("-", 90) + "\n")
		for _, result := range results {
			line := fmt.Sprintf("%-30s %-30s %-15.2f %-15s\n",
				result.File1, result.File2, result.Similarity, result.Category)
			writer.WriteString(line)
		}
	}

	return nil
}
