package kernel

// RetrieveAppConf returns a `Conf` struct by parsing the main config.yml file.
func RetrieveAppConf() *ServerConf {
	return config
}

/*// RetrieveRoutingConf will parse router.yml file (present in Go-Web root dir) and return a Router structure.
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
}*/
