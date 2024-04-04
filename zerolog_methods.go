package logease

import "net/http"

type ZerologMethods interface {
	Info(message string, params ...interface{})
	Warning(message string, params ...interface{})
	Debug(message string, params ...interface{})
	Error(r *http.Request, errorStack error, message string, params ...interface{})
}

// Info method implemented on zerolog package's logger instance and print info level log to console.
func (z ZerologLoggerInstance) Info(message string, params ...interface{}) {
	z.ZerologLoggerInstance.Info().Timestamp().Str("message", message).Fields(params)
}

// Warning method implemented on zerolog package's logger instance and print warning level log to console.
func (z ZerologLoggerInstance) Warning(message string, params ...interface{}) {
	z.ZerologLoggerInstance.Warn().Timestamp().Str("message", message).Fields(params).Msg("")
}

// Debug method implemented on zerolog package's logger instance and print debug level log to console.
func (z ZerologLoggerInstance) Debug(message string, params ...interface{}) {
	z.ZerologLoggerInstance.Debug().Timestamp().Str("message", message).Fields(params)
}

// Error method implemented on zerolog package's logger instance and send error details to teams in case of SendErrorToTeams enabled otherwise it will print error level log to console.
func (z ZerologLoggerInstance) Error(r *http.Request, errorStack error, message string, params ...interface{}) {
	if SendErrorToTeams {
		err := SendErrorDetailsToTeams(r, errorStack, message, params...)
		if err != nil {
			z.ZerologLoggerInstance.Error().Timestamp().Str("message", err.Error())
			z.ZerologLoggerInstance.Error().Timestamp().Str("message", message).Fields(params)
			return
		}
	} else {
		z.ZerologLoggerInstance.Error().Timestamp().Str("message", message).Fields(params)
	}
}
