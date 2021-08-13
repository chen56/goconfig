package json

import (
	"encoding/json"

	"github.com/chen56/goconfig"
)

type codec struct {
}

func New() goconfig.Codec {
	return &codec{}
}

func (t *codec) Marshal(value interface{}) (out []byte, err error) {
	return json.Marshal(value)
}

func (t *codec) Unmarshal(in []byte, out interface{}) error {
	return json.Unmarshal(in, out)
}

func (t *codec) String() string {
	return "json"
}
