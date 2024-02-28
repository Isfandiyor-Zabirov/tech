package service

import (
	"log"
	"net/http"
)

func CreateTask() uint {
	return createTask()
}

func SendRequest(taskID uint, request *CreateTaskRequest) {
	var (
		client        = &http.Client{}
		statusCode    int
		contentLength uint64 = 0
		headers       http.Header
		status        = "done"
	)

	defer func(code *int, header *http.Header, length *uint64, updatedStatus *string) {
		updateTask(taskID, uint(*code), header, *length, *updatedStatus)
	}(&statusCode, &headers, &contentLength, &status)

	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		log.Println("SendRequest func http request error:", err.Error())
		status = "error"
		statusCode = http.StatusNotFound
		return
	}

	req.Header.Add("Authentication", request.Headers.Authentication)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("SendRequest request execution error:", err.Error())
		status = "error"
		statusCode = http.StatusInternalServerError
		return
	}

	headers = resp.Header
	statusCode = resp.StatusCode
	contentLength = uint64(resp.ContentLength)

	defer resp.Body.Close()
}

func GetTaskByID(taskID uint) Task {
	return getTaskByID(taskID)
}
