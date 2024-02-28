package utils

import "sync"

var mu *sync.Mutex

func InitMutex() {
	mu = &sync.Mutex{}
}

func GetMutex() *sync.Mutex {
	return mu
}
