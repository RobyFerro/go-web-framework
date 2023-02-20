package entities

// AppConf contains main GoWeb configuration
type AppConf struct {
	Name    string
	Port    int
	SSL     bool
	SSLCert string
	SSLKey  string
	Key     string
}
