package markdown

import (
	"os"
	"path/filepath"
)

func ReadFile(path string) (Node, error) {
	if filepath.Ext(path) != ".md" {
		return nil, nil
	}

	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return File{
		Name:      info.Name(),
		Size:      info.Size(),
		Src:       b,
		UpdatedAt: info.ModTime(),
	}, nil
}
