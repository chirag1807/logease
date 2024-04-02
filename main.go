package main

import (
	"log"

	"github.com/chirag1807/logease/logease"
)

func main() {
	// LoggerInstance, err := logease.InitLogease(true, "", logease.ZeroLogger)
	// if err != nil {
	// 	log.Println(err)
	// }
	// v := LoggerInstance.(logease.ZerologLoggerInstance)
	// message := "This is Chirag Makwana"
	// params := map[string]interface{}{
	// 	"param1": 1,
	// 	"param2": 11,
	// 	"param3": 1,
	// 	"param4": 1,
	// 	"param5": 1,
	// 	"param6": 1,
	// 	"param7": 1,
	// }
	// v.Warning(message, params)

	LoggerInstance, err := logease.InitLogease(true, "https://zuruinc.webhook.office.com/webhookb2/1f8c1832-693b-465e-b543-25aa918b2755@cc9c131c-98a0-48b2-85fe-dfd27d6c8ecc/IncomingWebhook/0ff0ce62b5934d9bbb033b2077a0810e/df1ac8ce-7ccb-43be-9bb1-d681858219f7", logease.Slog)
	if err != nil {
		log.Println(err)
	}
	v := LoggerInstance.(logease.SlogLoggerInstance)
	message := "This is Testing Purpose Error Log."
	params := map[string]interface{}{
		"param1": 1,
		"param2": 11,
	}
	v.Error(message, params)
}
