package goconfig

import (
	"fmt"
	"io/ioutil"
)

//NewSourceFromFile 文件配置源
func NewSourceFromFile(path string) Source {
	return &fileSource{path: path}
}

type fileSource struct {
	path string
}

func (t *fileSource) Read() ([]byte, error) {
	return ioutil.ReadFile(t.path)
}

func (t *fileSource) String() string {
	return "file:" + t.path
}

//NewMemoryConfigSource 内存配置源
func NewSourceFromBytes(data []byte) Source {
	return &memorySource{data: data}
}

type memorySource struct {
	data []byte
}

func (t *memorySource) Read() ([]byte, error) {
	return t.data, nil
}
func (t *memorySource) String() string {
	return "memory"
}

//NewMemoryConfigSource 内存配置源
func NewInterfaceConfigSource(conf interface{}) Source {
	return &interfaceConfigSource{conf: conf}
}

type interfaceConfigSource struct {
	conf interface{}
}

func (t *interfaceConfigSource) Read() ([]byte, error) {
	return marshal(t.conf)
}

func (t *interfaceConfigSource) String() string {
	return fmt.Sprintf("interface: %T", t.conf)
}
