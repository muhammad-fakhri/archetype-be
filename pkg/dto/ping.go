package dto

type ResponsePing struct {
	AppName         string `json:"app_name"`
	Environment     string `json:"environment"`
	Message         string `json:"message"`
	ServerTimestamp int64  `json:"server_timestamp"`
}
