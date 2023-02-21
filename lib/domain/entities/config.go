package entities

// Config contains main GoWeb configuration
type Config struct {
	Name    string
	Port    int
	SSL     bool
	SSLCert string
	SSLKey  string
	Key     string
}
