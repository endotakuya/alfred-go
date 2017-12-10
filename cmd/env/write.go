package env

import (
	"fmt"
	"os"
)

func (c *Config) Write(f *os.File) {
	fmt.Fprintf(f, `package %s

func Get(k string) string {
	return List()[k]
}

func List() map[string]string {
	return map[string]string{`, c.PackageName)

	for k, v := range c.Data {
		fmt.Fprintf(f, `
		"%s": "%s",`, k, v)
	}

	fmt.Fprint(f, `
	}
}`)
}
