package input

import (
	"bufio"
	"encoding/json"
	"os"
)

type Record struct {
	Answer string `json:"answer"`
}

func ReadAnswers(path string) ([]Record, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records []Record

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var r Record

		if err := json.Unmarshal(scanner.Bytes(), &r); err != nil {
			return nil, err
		}

		records = append(records, r)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
