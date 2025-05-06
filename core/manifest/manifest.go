package manifest

import (
	"encoding/json"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/thegogod/docs.md/core/collections"
)

type Manifest struct {
	Name    string                            `toml:"name" json:"name"`
	Version *string                           `json:"version,omitempty" json:"version"`
	Build   Build                             `toml:"build" json:"build"`
	Plugins map[string]collections.Dictionary `toml:"plugins" json:"plugins"`
}

func Load(path string) (Manifest, error) {
	value := Manifest{
		Build: Build{
			OutDir: "dist",
		},
		Plugins: map[string]collections.Dictionary{},
	}

	if _, err := toml.DecodeFile(filepath.Join(path, "manifest.toml"), &value); err != nil {
		return value, err
	}

	return value, nil
}

func (self Manifest) String() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}
