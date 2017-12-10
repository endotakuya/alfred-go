package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/endotakuya/alfred-go/cmd/env"
)

const DEFAULT_ENVIRONMENT_FILE = ".env"

func main() {
	f, err := os.Open(DEFAULT_ENVIRONMENT_FILE)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	c := env.New("env", "env.go", "env")
	for s.Scan() {
		c.Add(getEnv(s.Text()))
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	if len(c.Data) == 0 {
		log.Fatalln("No environment data")
	}
	c.Convert()
}

func getEnv(s string) map[string]string {
	i := strings.Index(s, "#")

	// Exists comment
	if i != -1 {
		comment := s[i:(len(s))]
		s = strings.Replace(s, comment, "", -1)
	}

	// Delete space
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}

	e := strings.Split(s, "=")
	e[1] = strings.Replace(e[1], "\"", "", -1)
	return map[string]string{ e[0]: e[1] }
}
