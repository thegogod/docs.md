package markdown

import (
	"bytes"
	"encoding/json"
	"time"
)

type File struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Src       []byte    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
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

func (self File) String() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}
