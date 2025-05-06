package markdown

import (
	"os"
)

func Read(path string) (Node, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return ReadDir(path)
	}

	return ReadFile(path)
}
