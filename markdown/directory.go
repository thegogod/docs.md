package markdown

import "encoding/json"

type Directory struct {
	Name  string `json:"name"`
	Nodes []Node `json:"nodes,omitempty"`
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
