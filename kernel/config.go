package kernel

import (
	"github.com/RobyFerro/go-web-framework/tool"
	"gopkg.in/yaml.v3"
	"os"
)

// Conf represents the main configuration of Go-Web
// You can implement this method if wanna implement more configuration.
// Remember: this struct will be populated by parsing the config.yml file present into the Go-Web main directory.
// You've to implement both to works properly.
type Conf struct {
	Server struct {
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
		Ssl      bool   `yaml:"ssl"`
		SslCert  string `yaml:"sslcert"`
		SslKey   string `yaml:"sslkey"`
		RunUser  string `yaml:"run-user"`
		RunGroup string `yaml:"run-group"`
	} `yaml:"server"`
	App struct {
		Key string `yaml:"key"`
	} `yaml:"app"`
	Mail struct {
		From     string `yaml:"from"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
	} `yaml:"mail"`
}

// RetrieveRoutingConf will parse router.yml file (present in Go-Web root dir) and return a Router structure.
// This structure will be used by the HTTP kernel to setup every routes.
func RetrieveRoutingConf() (*Router, error) {
	var conf Router
	routePath := tool.GetDynamicPath("routing.yml")
	c, err := os.Open(routePath)

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

// RetrieveAppConf returns a `Conf` struct by parsing the main config.yml file.
func RetrieveAppConf() (*Conf, error) {
	var conf Conf
	confFile := tool.GetDynamicPath("config.yml")
	c, err := os.Open(confFile)

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
