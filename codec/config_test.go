package goconfig_test

import (
	"log"
	"testing"

	"github.com/chen56/goconfig"
	"github.com/chen56/goconfig/codec/yaml"
)

func Test_Name(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Fatalf("%s: ", tt.name)
		})
	}
}
func Test_T(t *testing.T) {
	var source = goconfig.NewSourceFromBytes([]byte(`
redis:
	addr: localhost:6379
	password: 
	db: 1
`))
	config := goconfig.NewBuilder{
		Source: source,
		Codec:  yaml.New(),
	}.Build()
	log.Printf("test:%s", config.String())

}
