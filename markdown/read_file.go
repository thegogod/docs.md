package markdown

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(base string, path string) (Node, error) {
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

	cwd, _ := os.Getwd()
	relpath, err := filepath.Rel(cwd, path)

	if err != nil {
		return nil, err
	}

	return File{
		Path:      path,
		RelPath:   strings.Replace(relpath, base, "", 1),
		Name:      info.Name(),
		Size:      info.Size(),
		Src:       b,
		UpdatedAt: info.ModTime(),
	}, nil
}
