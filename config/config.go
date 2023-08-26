package config

import (
	"os"

	"github.com/bytedance/sonic"
)

var Config *config

type config struct {
	Db struct {
		Uri    string `json:"uri"`
		DbName string `json:"dbName"`
	} `json:"db"`
	Api struct {
		Addr         string `json:"addr"`
		LogLevel     string `json:"logLevel"`
		JwtSecret    string `json:"jwtSecret"`
		CookieSecret string `json:"cookieSecret"`
	} `json:"api"`
}

type ConfigParser interface {
	Parse() *config
}

type jsonConfigParser struct {
	path string
}

func NewJsonConfigParser(path string) ConfigParser {
	return &jsonConfigParser{path: path}
}

func (parser *jsonConfigParser) Parse() *config {
	raw, err := os.ReadFile(parser.path)
	if err != nil {
		panic(err)
	}
	var conf config
	if err := sonic.Unmarshal(raw, &conf); err != nil {
		panic(err)
	}
	Config = &conf
	return &conf
}
