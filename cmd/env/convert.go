package env

import (
	"log"
	"os"
	"path/filepath"
)

// Convert to .go file
func (c *Config) Convert() {

	// Is directory exists?
	if _, err := os.Stat(c.DirName); err != nil {
		if err := os.Mkdir(c.DirName, 0777); err != nil {
			log.Fatal(err)
		}
	}

	path := c.filePath()
	// Is file exists?
	if _, err := os.Stat(path); err == nil {
		log.Fatalf("Already %s is exists.", c.FileName)
	}

	var f *os.File
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	c.Write(f)
}

func (c *Config) filePath() string {
	return filepath.Join(c.DirName, c.FileName)
}
