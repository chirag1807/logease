package logease

import (
	"errors"
	"log/slog"
	"os"

	"github.com/rs/zerolog"
)

type PackageToUse string

const (
	ZeroLogger PackageToUse = "ZeroLogger"
	Slog       PackageToUse = "Slog"
)

var TeamsURL string

// https://zuruinc.webhook.office.com/webhookb2/1f8c1832-693b-465e-b543-25aa918b2755@cc9c131c-98a0-48b2-85fe-dfd27d6c8ecc/IncomingWebhook/0ff0ce62b5934d9bbb033b2077a0810e/df1ac8ce-7ccb-43be-9bb1-d681858219f7
var EnableNotification bool

type ZerologLoggerInstance struct {
	ZerologLoggerInstance *zerolog.Logger
}

type SlogLoggerInstance struct {
	SlogLoggerInstance slog.Logger
}

func InitLogease(enableNotification bool, teamsURL string, packageToUse PackageToUse) (interface{}, error) {
	if packageToUse != ZeroLogger && packageToUse != Slog {
		return nil, errors.New("provide valid package name. field should be either ZeroLogger or Slog")
	}

	TeamsURL = teamsURL
	EnableNotification = enableNotification

	switch packageToUse {
	case ZeroLogger:
		zerologLoggerInstance := zerolog.New(os.Stdout).With().Logger()
		logease := ZerologLoggerInstance{
			ZerologLoggerInstance: &zerologLoggerInstance,
		}
		return logease, nil
	case Slog:
		slogLoggerInstance := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logease := SlogLoggerInstance{
			SlogLoggerInstance: *slogLoggerInstance,
		}
		return logease, nil
	}

	return nil, nil
}
