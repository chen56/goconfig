package goconfig

import (
	"fmt"
)

//Source 配置源，配置可以从字符串/文件/网络读取
type Source interface {
	fmt.Stringer
	Read() ([]byte, error)
}
type Codec interface {
	fmt.Stringer
	Marshal(value interface{}) (out []byte, err error)
	Unmarshal(in []byte, out interface{}) error
}
type NewBuilder struct {
	Source Source
	Codec  Codec
}
type Config struct {
	fmt.Stringer
	source Source
	codec  Codec
}

func (builder NewBuilder) Build() *Config {
	return &Config{
		source: builder.Source,
		codec:  builder.Codec,
	}
}
func (self *Config) String() string {
	return fmt.Sprintf("%s:%s", self.source, self.codec)
}
