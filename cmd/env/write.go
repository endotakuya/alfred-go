package env

import (
	"bufio"
	"fmt"
	"os"
)

func (c *Config) Write(f *os.File) {
	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, `package %s

func Get(k string) string {
	return List()[k]
}

func List() map[string]string {
	return map[string]string{`, c.PackageName)

	for k, v := range c.Data {
		fmt.Fprintf(w, `
		"%s": "%s",`, k, v)
	}

	fmt.Fprint(w, `
	}
}`)
}
