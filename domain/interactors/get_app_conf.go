package interactors

import "github.com/RobyFerro/go-web-framework/domain/entities"

// GetAppConfig returns an applicazion config structure
type GetAppConfig struct{}

// Call executes usecase logic
func (c GetAppConfig) Call() entities.Config {
	return entities.Config{
		Port:    8005,
		SSL:     false,
		SSLCert: "storage/certs/tls.crt",
		SSLKey:  "storage/certs/tls.key",
		Key:     "REPLACE_WITH_CUSTOM_APP_KEY",
	}
}
