package splitter

import (
	"fmt"
	"path/filepath"
	"strings"

	"splitpdf/internal/parser"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func SplitPDF(pdfPath string, chapters []parser.Chapter, prefix string) error {
	outDir := filepath.Dir(pdfPath)
	conf := model.NewDefaultConfiguration()

	for _, ch := range chapters {
		title := sanitize(ch.Title)
		filename := fmt.Sprintf("%s_%s.pdf", ch.Chapter, title)
		if prefix != "" {
			filename = fmt.Sprintf("%s_%s", prefix, filename)
		}

		outputPath := filepath.Join(outDir, filename)
		pageRange := fmt.Sprintf("%d-%d", ch.FirstPage, ch.LastPage)

		if err := api.TrimFile(pdfPath, outputPath, []string{pageRange}, conf); err != nil {
			return fmt.Errorf("error trimming %s: %w", filename, err)
		}
	}

	return nil
}

func sanitize(title string) string {
	replacer := strings.NewReplacer(
		" ", "", "-", "", ",", "", ".", "", "\"", "",
		"'", "", "/", "", "\\", "", "(", "", ")", "",
		":", "", ";", "", "&", "and",
	)
	return replacer.Replace(title)
}
