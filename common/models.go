package common

// ErrorResponse default structure from error responses
type ErrorResponse struct {
	UserMessage     string `json:"userMessage"`
	RandomGenerate  int    `json:"randomGenerate,omitempty"`
	InternalMessage string `json:"internalMessage"`
	MoreInfo        string `json:"moreInfo"`
}

// Response default structure response
type Response struct {
	Param  int    `json:"param,omitempty"`
	Random int    `json:"random"`
	PiCalc string `json:"PiCalc"`
}
