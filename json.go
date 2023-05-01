package gophig

import "github.com/goccy/go-json"

// JSONMarshaler is a Marshaler that uses the goccy/go-json package.
type JSONMarshaler struct{}

// Marshal ...
func (JSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal ...
func (JSONMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
