package logease

type PackageToUse string

const (
	ZeroLogger PackageToUse = "ZeroLogger"
	Slog       PackageToUse = "Slog"
)

type RequestMethod string

const (
	POST   RequestMethod = "POST"
	GET    RequestMethod = "GET"
	PATCH  RequestMethod = "PATCH"
	PUT    RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
)

// Used By Teams Message Card
const (
	Type          string = "Error Log Card"
	Context       string = "http://schema.org/extensions"
	Summary       string = "GOLang Application's Errors"
	ThemeColor    string = "0076D7"
	ActivityTitle string = "Error-Details:"
	ReqMethod     string = "Request-Method:"
	ReqURL        string = "Request-URL:"
	QueryParams   string = "Query-Params:"
	Time          string = "Time:"
	Message       string = "Message:"
	ReqBody       string = "Request-Body:"
	Parameters    string = "Parameters:"
	ErrorStack    string = "Error-Stack:"
)

// Error From logease Package
const (
	JSONMarshalError       string = "Error Message from logease - Could not Send Error Log to Teams Due to Error Occured at JSON Marshalling"
	ReadBodyError          string = "Error Message from logease - Could not Send Error Log to Teams Due to Error Occured While Reading Request Body"
	RequestNotCreatedError string = "Error Message from logease - Could not Send Error Log to Teams Due to POST Reuqest Could not be Created"
	RequestNotSentError    string = "Error Message from logease - Could not Send Error Log to Teams"
)
