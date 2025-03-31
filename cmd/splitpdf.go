package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"splitpdf/internal/parser"
	"splitpdf/internal/splitter"
)

func Execute() {
	inputPtr := flag.String("i", "", "Input PDF file")
	csvPtr := flag.String("f", "", "CSV file with chapters")
	prefixPtr := flag.String("n", "", "Optional prefix for output files")
	flag.Parse()

	var pdfPath, csvPath string

	// Case 1: PDF and CSV provided
	if *inputPtr != "" && *csvPtr != "" {
		pdfPath = *inputPtr
		csvPath = *csvPtr
	} else if flag.NArg() == 1 {
		// Case 2: Directory mode
		dir := flag.Arg(0)
		_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				switch filepath.Ext(path) {
				case ".pdf":
					if pdfPath == "" {
						pdfPath = path
					}
				case ".csv":
					if csvPath == "" {
						csvPath = path
					}
				}
			}
			return nil
		})
	} else {
		fmt.Fprintln(os.Stderr, "Usage: splitpdf -i input.pdf -f chapters.csv [-n prefix] OR splitpdf <directory>")
		os.Exit(1)
	}

	if pdfPath == "" || csvPath == "" {
		fmt.Fprintln(os.Stderr, "PDF or CSV file not found.")
		os.Exit(1)
	}

	chapters, err := parser.ParseCSV(csvPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CSV error: %v\n", err)
		os.Exit(1)
	}

	err = splitter.SplitPDF(pdfPath, chapters, *prefixPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Split error: %v\n", err)
		os.Exit(1)
	}
}

