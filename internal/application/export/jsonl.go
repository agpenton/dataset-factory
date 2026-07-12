package export

import (
	"os"
	"strings"
)

func JSONL(path string, records []string) error {
	content := strings.Join(records, "\n") + "\n"

	return os.WriteFile(
		path,
		[]byte(content),
		0o644,
	)
}
