package manifest

import "encoding/json"

type Build struct {
	OutDir string `toml:"outdir" json:"outdir"`
}

func (self Build) String() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}
