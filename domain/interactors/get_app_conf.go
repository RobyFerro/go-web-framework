package interactors

import "github.com/RobyFerro/go-web-framework/domain/entities"

// GetAppConfig will reuturn a new app configuration structure
func GetAppConfig() entities.AppConf {
	return entities.AppConf{
		Port:    8005,
		SSL:     false,
		SSLCert: "storage/certs/tls.crt",
		SSLKey:  "storage/certs/tls.key",
		Key:     "REPLACE_WITH_CUSTOM_APP_KEY",
	}
}
