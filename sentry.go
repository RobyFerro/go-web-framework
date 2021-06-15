package gwf

import (
	"github.com/RobyFerro/go-web-framework/helper"
	"github.com/getsentry/sentry-go"
)

// Send an error to Sentry.
// This requires sentry endpoint configured into the config.yml file
func SentryReport(report error, conf helper.ErrorConfiguration) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: conf.Exception.Sentry,
	}); err != nil {
		helper.Log(err.Error())
	}

	sentry.CaptureException(report)
}
