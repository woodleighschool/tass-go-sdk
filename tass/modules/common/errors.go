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
