package kernel

// RetrieveAppConf returns a `Conf` struct by parsing the main config.yml file.
func RetrieveAppConf() *ServerConf {
	return config
}
