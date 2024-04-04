package logease

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var excludedWords = []string{
	"token",
	"password",
	"requestMethod",
	"RequestMethod",
}

// SendErrorDetailsToTeams is generalized function for sending error log/message to teams.
func SendErrorDetailsToTeams(r *http.Request, errorStack error, message string, params ...interface{}) error {
	paramsString, err := json.Marshal(params)
	if err != nil {
		return errors.New(JSONMarshalError)
	}

	var requestBody map[string]interface{}
	if r != nil {
		if r.Body != nil && r.Body != http.NoBody {
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			if err != nil {
				return errors.New(ReadBodyError)
			}

			for _, field := range excludedWords {
				delete(requestBody, field)
			}

			// r.Body = io.NopCloser(bytes.NewReader(newBody))
		}
	}
	requestBodyString, err := json.Marshal(requestBody)
	if err != nil {
		return errors.New(JSONMarshalError)
	}

	queryParams := make(map[string][]string)
	if r != nil && r.URL != nil {
		queryParams = r.URL.Query()
	}
	for _, key := range excludedWords {
		delete(queryParams, key)
	}
	queryParamsString, err := json.Marshal(queryParams)
	if err != nil {
		return errors.New(JSONMarshalError)
	}

	teamsMessage := TeamsMessage{
		Type:       Type,
		Context:    Context,
		Summary:    Summary,
		ThemeColor: ThemeColor,
		Sections: []TeamsMessageSection{
			{
				ActivityTitle: ActivityTitle,
				Facts: []TeamsMessageFact{
					{Name: ReqMethod, Value: r.Method},
					{Name: ReqURL, Value: r.URL.Scheme + "://" + r.URL.Host + r.URL.Path},
					{Name: QueryParams, Value: string(queryParamsString)},
					{Name: Time, Value: time.Now().Format(time.RFC1123)},
					{Name: Message, Value: message},
					{Name: ReqBody, Value: string(requestBodyString)},
					{Name: Parameters, Value: string(paramsString)},
					{Name: ErrorStack, Value: string(errorStack.Error())},
				},
			},
		},
	}

	teamsMessageJSON, err := json.Marshal(teamsMessage)
	if err != nil {
		return errors.New(JSONMarshalError)
	}

	req, err := http.NewRequest("POST", TeamsURL, bytes.NewBuffer(teamsMessageJSON))
	if err != nil {
		return errors.New(RequestNotCreatedError)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return errors.New(RequestNotSentError)
	}
	defer response.Body.Close()

	fmt.Println(response.Status)
	if response.StatusCode != http.StatusOK {
		return errors.New(RequestNotSentError)
	}

	return nil
}
