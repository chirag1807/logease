package logease

import "net/http"

type SlogMethods interface {
	Info(message string, params ...interface{})
	Warning(message string, params ...interface{})
	Debug(message string, params ...interface{})
	Error(r *http.Request, errorStack error, message string, params ...interface{})
}

// Info method implemented on slog package's logger instance and print info level log to console.
func (s SlogLoggerInstance) Info(message string, params ...interface{}) {
	s.SlogLoggerInstance.Info(message, "params", params)
}

// Warning method implemented on slog package's logger instance and print warning level log to console.
func (s SlogLoggerInstance) Warning(message string, params ...interface{}) {
	s.SlogLoggerInstance.Warn(message, "params", params)
}

// Debug method implemented on slog package's logger instance and print debug level log to console.
func (s SlogLoggerInstance) Debug(message string, params ...interface{}) {
	s.SlogLoggerInstance.Debug(message, "params", params)
}

// Error method implemented on slog package's logger instance and send error details to teams in case of SendErrorToTeams enabled otherwise it will print error level log to console.
func (s SlogLoggerInstance) Error(r *http.Request, errorStack error, message string, params ...interface{}) {
	if SendErrorToTeams {
		err := SendErrorDetailsToTeams(r, errorStack, message, params...)
		if err != nil {
			s.SlogLoggerInstance.Error(err.Error())
			s.SlogLoggerInstance.Error(message, "params", params)
			return
		}
	} else {
		s.SlogLoggerInstance.Error(message, "params", params)
	}
}
