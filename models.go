package logease

import (
	"log/slog"

	"github.com/rs/zerolog"
)

// ZerologLoggerInstance includes instance of zerolog's logger which will be used for implementing its native log methods.
type ZerologLoggerInstance struct {
	ZerologLoggerInstance *zerolog.Logger
}

// SlogLoggerInstance includes instance of slog's logger which will be used for implementing its native log methods.
type SlogLoggerInstance struct {
	SlogLoggerInstance slog.Logger
}

// TeamsMessage includes details of card that will be shown to members of team's channel like type and context of card, summary, theme color and message section.
type TeamsMessage struct {
	Type       string                `json:"@type"`
	Context    string                `json:"@context"`
	Summary    string                `json:"summary,omitempty"`
	ThemeColor string                `json:"themeColor,omitempty"`
	Sections   []TeamsMessageSection `json:"sections"`
}

// TeamsMessageSection includes activity title, subtitle, text and facts about error log.
type TeamsMessageSection struct {
	ActivityTitle    string             `json:"activityTitle"`
	ActivitySubtitle string             `json:"activitySubtitle"`
	ActivityText     string             `json:"activityText"`
	Facts            []TeamsMessageFact `json:"facts,omitempty"`
}

// TeamsMessageFact includes details of error log in key-value pair.
type TeamsMessageFact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
