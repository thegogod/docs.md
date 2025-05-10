package markdown

import (
	"os"
)

func Read(base string, path string) (Node, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return ReadDir(base, path)
	}

	return ReadFile(base, path)
}
