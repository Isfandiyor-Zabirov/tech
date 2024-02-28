package service

import (
	"tech/internal/entities/utils"
	"testing"
)

func TestCreateTask(t *testing.T) {
	utils.InitMutex()
	lengthBefore := len(taskList)
	taskIDBefore := taskID
	createTask()
	lengthAfter := len(taskList)
	taskIDAfter := taskID

	differenceGot := lengthAfter - lengthBefore
	differenceWant := 1

	if differenceGot != differenceWant {
		t.Errorf("got length difference: %d, wanted length difference: %d", differenceGot, differenceWant)
	}

	if taskIDAfter != taskIDBefore+1 {
		t.Errorf("got task ID: %d, wanted task ID: %d", taskIDAfter, taskIDBefore+1)
	}
}

func TestGetTaskByID(t *testing.T) {
	utils.InitMutex()

	request := CreateTaskRequest{
		Method: "GET",
		Url:    "https://google.com",
		Headers: Headers{
			Authentication: "Bearer ef903.eu8euf.921312",
		},
	}

	taskId := createTask()
	SendRequest(taskId, &request)

	task := getTaskByID(taskId)

	if task.Status != "done" {
		t.Errorf("got task status: %s, wanted task status: %s", task.Status, "done")
	}

	request = CreateTaskRequest{
		Method: "GET",
		Url:    "https://google.coms",
		Headers: Headers{
			Authentication: "Bearer ef903.eu8euf.921312",
		},
	}

	taskId = createTask()
	SendRequest(taskId, &request)

	task = getTaskByID(taskId)

	if task.Status != "error" {
		t.Errorf("got task status: %s, wanted task status: %s", task.Status, "error")
	}
}
