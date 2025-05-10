package markdown

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadDir(base string, path string) (Node, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	cwd, _ := os.Getwd()
	relpath, err := filepath.Rel(cwd, path)

	if err != nil {
		return nil, err
	}

	dir := Directory{
		Path:    path,
		RelPath: strings.Replace(relpath, base, "", 1),
		Name:    info.Name(),
		Nodes:   []Node{},
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		node, err := Read(base, filepath.Join(path, entry.Name()))

		if err != nil {
			return nil, err
		}

		if node == nil {
			continue
		}

		dir.Add(node)
	}

	return dir, nil
}
