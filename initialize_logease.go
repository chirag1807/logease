package logease

import (
	"errors"
	"log/slog"
	"os"

	"github.com/rs/zerolog"
)

var TeamsURL string
var SendErrorToTeams bool

// InitLogease function takes 3 paramters, sendErrorToTeams means you want to print error details to teams, if yes then teams URL and 3rd is package that you want to use it's functionality either it is zerolog or slog.
func InitLogease(sendErrorToTeams bool, teamsURL string, packageToUse PackageToUse) (interface{}, error) {
	if packageToUse != ZeroLogger && packageToUse != Slog {
		return nil, errors.New("provide valid package name. field should be either ZeroLogger or Slog")
	}

	TeamsURL = teamsURL
	SendErrorToTeams = sendErrorToTeams

	switch packageToUse {
	case ZeroLogger:
		zerologLoggerInstance := zerolog.New(os.Stdout).With().Logger()
		return NewZerologLoggerInstance(zerologLoggerInstance), nil
	case Slog:
		slogLoggerInstance := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		return NewSlogLoggerInstance(slogLoggerInstance), nil
	}

	return nil, nil
}

func NewSlogLoggerInstance(slogLoggerInstance *slog.Logger) SlogMethods {
	return SlogLoggerInstance{
		SlogLoggerInstance: *slogLoggerInstance,
	}
}

func NewZerologLoggerInstance(zerologLoggerInstance zerolog.Logger) ZerologMethods {
	return ZerologLoggerInstance{
		ZerologLoggerInstance: &zerologLoggerInstance,
	}
}
