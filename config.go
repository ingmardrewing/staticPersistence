package staticPersistence

import (
	"strings"

	"github.com/ingmardrewing/fs"
)

func NewConfig(path string) *Config {
	c := new(Config)
	c.ReadConfigFile(path)
	return c
}

// Config reading from json data
type Config struct {
	Json
	data []byte
}

func (c *Config) ReadDir(keys ...string) string {
	path := c.Read(keys...)
	if !strings.HasSuffix(path, "/") {
		return path + "/"
	}
	return path
}

func (c *Config) Read(keys ...string) string {
	return c.ReadString(c.data, keys...)
}

func (c *Config) ReadConfigFile(path string) {
	bytes := fs.ReadByteArrayFromFile(path)
	c.data = bytes
}
