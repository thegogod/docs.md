package collections

import "encoding/json"

type Dictionary map[string]any

func (self Dictionary) Has(key string) bool {
	_, exists := self[key]
	return exists
}

func (self Dictionary) Set(key string, value any) Dictionary {
	self[key] = value
	return self
}

func (self Dictionary) Get(key string) any {
	if value, exists := self[key]; exists {
		return value
	}

	return nil
}

func (self Dictionary) String() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}
