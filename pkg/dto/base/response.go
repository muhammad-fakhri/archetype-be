package base

type Response struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors *ErrorResponse `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"-"`
	Detail  string `json:"detail,omitempty"`
}
