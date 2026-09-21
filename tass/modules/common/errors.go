package tasscommon

type ExceptionDetails struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type ValidationExceptionDetails struct {
	Errors Errors `json:"errors"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type Errors struct {
	AdditionalProperties []string `json:"additionalProperties"`
}

const DateError string = `%s: %s does not parse as a date: %w`
const MaxLengthError string = `%s: %s must be less than or equal to %s`
