package yaml

import (
	"github.com/chen56/goconfig"
	yaml2 "gopkg.in/yaml.v2"
)

type codec struct {
}

func New() goconfig.Codec {
	return &codec{}
}

func (t *codec) Marshal(value interface{}) (out []byte, err error) {
	return yaml2.Marshal(value)
}

func (t *codec) Unmarshal(in []byte, out interface{}) error {
	return yaml2.Unmarshal(in, out)
}

func (t *codec) String() string {
	return "yaml"
}
