package markdown

import (
	"bytes"
	"encoding/json"
	"html/template"
	"time"
)

type File struct {
	Path      string    `json:"path"`
	RelPath   string    `json:"rel_path"`
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Src       []byte    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (self File) GetPath() string {
	return self.Path
}

func (self File) GetRelPath() string {
	return self.RelPath
}

func (self File) GetName() string {
	return self.Name
}

func (self File) GetNodes() []Node {
	return []Node{}
}

func (self File) Render() ([]byte, error) {
	var buf bytes.Buffer

	if err := Parser.Convert(self.Src, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (self File) Parse(template *template.Template) (*template.Template, error) {
	_, err := template.New(self.Name).Parse(string(self.Src))
	return template, err
}

func (self File) String() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}
