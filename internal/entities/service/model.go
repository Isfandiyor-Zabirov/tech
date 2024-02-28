package service

type CreateTaskRequest struct {
	Method  string  `json:"method"`
	Url     string  `json:"url"`
	Headers Headers `json:"headers"`
}

type Headers struct {
	Authentication string `json:"Authentication"`
}

type Task struct {
	ID             uint        `json:"id"`
	Status         string      `json:"status"`
	HttpStatusCode uint        `json:"httpStatusCode"`
	Headers        interface{} `json:"headers"`
	Length         uint64      `json:"length"`
}
