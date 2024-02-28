package service

import "tech/internal/entities/utils"

var taskID uint = 1
var taskList []Task

// createTask creates task ID and saves to list
func createTask() uint {

	// get global mutex to lock and unlock taskID and task list
	mu := utils.GetMutex()

	mu.Lock()

	var task Task
	task.ID = taskID
	task.Status = "in_process"

	saveTask(task)
	incrementID()

	mu.Unlock()

	return task.ID
}

// incrementID increments created task IDs
func incrementID() {
	taskID++
}

// saveTask saves the created task to task list
func saveTask(task Task) {
	taskList = append(taskList, task)
}

// updateTask updates task info in the list after sending request to third party
func updateTask(id uint, statusCode uint, headers interface{}, length uint64, status string) {
	for i, task := range taskList {
		if task.ID == id {
			taskList[i].HttpStatusCode = statusCode
			taskList[i].Headers = headers
			taskList[i].Length = length
			taskList[i].Status = status
		}
	}
}

func getTaskByID(taskID uint) Task {
	var task Task
	for _, t := range taskList {
		if t.ID == taskID {
			task = t
			break
		}
	}

	return task
}
