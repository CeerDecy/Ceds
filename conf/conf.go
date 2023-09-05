package conf

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type cdsConf struct {
	Ceds
}

type Ceds map[string]any

func (c Ceds) Host() string {
	return c["host"].(string)
}
func (c Ceds) Port() int64 {
	return c["port"].(int64)
}

var Config = &cdsConf{}

func init() {
	loadToml()
}

func loadToml() {
	configFile := flag.String("conf", "config/config.toml", "ceds config file")
	flag.Parse()
	if _, err := os.Stat(*configFile); err != nil {
		log.Println(err)
		return
	}
	_, err := toml.DecodeFile(*configFile, Config)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
