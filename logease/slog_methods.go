package logease

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LogEntry struct {
	Time   time.Time              `json:"time"`
	Level  string                 `json:"level"`
	Msg    string                 `json:"msg"`
	Params map[string]interface{} `json:"params"`
}

type TeamsMessage struct {
	Type       string                `json:"@type"`
	Context    string                `json:"@context"`
	Summary    string                `json:"summary,omitempty"`
	ThemeColor string                `json:"themeColor,omitempty"`
	Sections   []TeamsMessageSection `json:"sections"`
}

type TeamsMessageSection struct {
	ActivityTitle    string             `json:"activityTitle"`
	ActivitySubtitle string             `json:"activitySubtitle"`
	ActivityText     string             `json:"activityText"`
	Facts            []TeamsMessageFact `json:"facts,omitempty"`
}

type TeamsMessageFact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SlogMethods interface {
	Info(message string, params map[string]interface{})
	Warning(message string, params map[string]interface{})
	Debug(message string, params map[string]interface{})
	Error(message string, params map[string]interface{})
}

func (s SlogLoggerInstance) Info(message string, params map[string]interface{}) {
	s.SlogLoggerInstance.Info(message, "params", params)
}

func (s SlogLoggerInstance) Warning(message string, params map[string]interface{}) {
	s.SlogLoggerInstance.Warn(message, "params", params)
}

func (s SlogLoggerInstance) Debug(message string, params map[string]interface{}) {
	s.SlogLoggerInstance.Debug(message, "params", params)
}

func (s SlogLoggerInstance) Error(message string, params map[string]interface{}) {
	if EnableNotification {
		logEntry := LogEntry{
			Time:   time.Now(),
			Level:  "ERROR",
			Msg:    message,
			Params: params,
		}

		// teamsMessage := TeamsMessage{
		// 	Text: fmt.Sprintf("Time: %s\nMessage: %s\nParams: %v", time.Now().Format(time.RFC3339), message, params),
		// }

		teamsMessage := TeamsMessage{
			Type:    "Error Log Card",
			Context: "http://schema.org/extensions",
			Sections: []TeamsMessageSection{
				{
					ActivityTitle:    "Error Log",
					ActivitySubtitle: "Request Details",
					ActivityText:     fmt.Sprintf("Error occurred at %s", logEntry.Time),
					Facts: []TeamsMessageFact{
						{Name: "Request Method", Value: "POST"},
						{Name: "Request URL", Value: "https://example.com/api/endpoint"},
						{Name: "Message", Value: logEntry.Msg},
						{Name: "Parameters", Value: fmt.Sprintf("%v", logEntry.Params)},
					},
				},
			},
		}

		logEntryJSON, err := json.Marshal(teamsMessage)
		if err != nil {
			fmt.Println("Error marshaling log entry:", err)
			return
		}

		req, err := http.NewRequest("POST", TeamsURL, bytes.NewBuffer(logEntryJSON))

		if err != nil {
			fmt.Println("Error creating POST request.", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending POST request.", err)
			return
		}
		defer response.Body.Close()

		fmt.Println(response.Status)
		if response.StatusCode != http.StatusOK {
			fmt.Println("Error log can not be POST to teams.", err)
			return
		}
	} else {
		s.SlogLoggerInstance.Error(message, "params", params)
	}
}
