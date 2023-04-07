package module

type ResponseData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}