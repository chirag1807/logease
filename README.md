# logease

A Go package for logging application messages to the console based on their severity level, and additionally sending error level logs to Microsoft Teams with detailed error information.

## Installation
```
go get -u github.com/chirag1807/logease
```

## API
```
package main

import (
	"bytes"
        "fmt"
        "log"
	"net/http"
	"github.com/chirag1807/logease"
)

func main() {
        LoggerInstance, err := logease.InitLogease(true, "YOUR_TEAMS_CHANNEL_URL", logease.ZeroLogger)
        // You have the option to utilize the functionality of the log/slog package by replacing logease.ZeroLogger with logease.Slog.
        if err != nil {
            log.Println(err)
        }
        v := LoggerInstance.(logease.ZerologLoggerInstance)
        // If you opt to use the log/slog package, ensure to assert its type. Replace the previous assertion with logease.SlogLoggerInstance for asserting slog's type.

        dummyBody := []byte(`{"name": "chirag", "token":"eyjijk"}`)
        dummyReader := bytes.NewReader(dummyBody)
        dummyRequest, _ := http.NewRequest("GET", "http://www.example.com/api/v1/get-users-list?token=eyjijk&name=chirag", dummyReader)
        dummyError := fmt.Errorf("dummy error: something went wrong")
        dummyMessage := "Error occurred in the application"
        dummyParams := map[string]interface{}{
            "param1": 1,
            "param2": "chirag",
        }

        v.Error(dummyRequest, dummyError, dummyMessage, dummyParams)
        v.Info(dummyMessage, dummyParams)
        v.Warning(dummyMessage, dummyParams)
        v.Debug(dummyMessage, dummyParams)
}
```

## How It Works
Firstly InitLogease function accepts three parameters:
1. `sendErrorToTeams`: Indicates whether you want to send error details to Teams or only print them to the console.
2. `teamsURL`: Indicates the link to the Teams channel where error details will be sent by this package.
3. `packageToUse`: Indicates which package's functionality you want to utilize, either zerolog or log/slog.

After calling the function, it returns a logger instance of the corresponding package as an interface. To use the methods implemented on that logger instance, you need to perform a type assertion to the corresponding package. Once done, you can then utilize the methods as shown in the example code.

## Contribution
[<img alt="Chirag Makwana" src="https://github.com/chirag1807/task-management-system/assets/94277910/0e27ad00-c278-4eea-81df-8c3096c1ed2c" width="84" height="100" style="border-radius: 50%;" />](https://github.com/chirag1807)
