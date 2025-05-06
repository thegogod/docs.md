package markdown

import (
	"os"
	"path/filepath"
)

func ReadDir(path string) (Node, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	dir := Directory{
		Name:  info.Name(),
		Nodes: []Node{},
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		node, err := Read(filepath.Join(path, entry.Name()))

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
