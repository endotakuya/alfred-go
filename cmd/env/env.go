package env

type (
	Config struct {
		PackageName string
		FileName    string
		DirName     string
		Data        map[string]string
	}
)

func New(p, f, d string) *Config {
	return &Config{
		PackageName: p,
		FileName:    f,
		DirName:     d,
		Data:        map[string]string{},
	}
}

func (c *Config) Add(data map[string]string) {
	for k, v := range data {
		c.Data[k] = v
	}
}
