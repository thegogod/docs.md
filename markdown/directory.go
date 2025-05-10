package markdown

import "encoding/json"

type Directory struct {
	Path    string `json:"path"`
	RelPath string `json:"rel_path"`
	Name    string `json:"name"`
	Nodes   []Node `json:"nodes,omitempty"`
}

func (self Directory) GetPath() string {
	return self.Path
}

func (self Directory) GetRelPath() string {
	return self.RelPath
}

func (self Directory) GetName() string {
	return self.Name
}

func (self Directory) GetNodes() []Node {
	return self.Nodes
}

func (self *Directory) Add(nodes ...Node) *Directory {
	self.Nodes = append(self.Nodes, nodes...)
	return self
}

func (self Directory) String() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}
