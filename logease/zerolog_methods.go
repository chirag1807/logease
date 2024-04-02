package logease

type ZerologMethods interface {
	Info(message string, params map[string]interface{})
	Warning(message string, params map[string]interface{})
	Debug(message string, params map[string]interface{})
	Error(message string, params map[string]interface{})
}

func (z ZerologLoggerInstance) Info(message string, params map[string]interface{}) {
	z.ZerologLoggerInstance.Info().Timestamp().Str("message", message).Fields(params)
}

func (z ZerologLoggerInstance) Warning(message string, params map[string]interface{}) {
	z.ZerologLoggerInstance.Warn().Timestamp().Str("message", message).Fields(params).Msg("")
}

func (z ZerologLoggerInstance) Debug(message string, params map[string]interface{}) {
	z.ZerologLoggerInstance.Debug().Timestamp().Str("message", message).Fields(params)
}

func (z ZerologLoggerInstance) Error(message string, params map[string]interface{}) {
	z.ZerologLoggerInstance.Error().Timestamp().Str("message", message).Fields(params)
}
