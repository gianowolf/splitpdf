package parser

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Chapter struct {
	Chapter   string
	Title     string
	FirstPage int
	LastPage  int
}

func ParseCSV(filePath string) ([]Chapter, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var chapters []Chapter
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 4 {
			continue
		}

		first, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			return nil, err
		}
		last, err := strconv.Atoi(strings.TrimSpace(record[3]))
		if err != nil {
			return nil, err
		}

		chapters = append(chapters, Chapter{
			Chapter:   strings.TrimSpace(record[0]),
			Title:     strings.TrimSpace(record[1]),
			FirstPage: first,
			LastPage:  last,
		})
	}

	return chapters, nil
}

