package exloggo

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type GoroutineData struct {
	RequestId       map[int]string
	RequestClientId map[int]string
	sync.RWMutex
}

func NewGoroutineData() *GoroutineData {
	return &GoroutineData{
		RequestId:       make(map[int]string),
		RequestClientId: make(map[int]string),
	}
}

var goroutineData = NewGoroutineData()

func GetGoroutineId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func GetGoroutineRequestId() string {
	goroutineId := GetGoroutineId()
	return goroutineData.RequestId[goroutineId]
}

func GetGoroutineRequestClientId() string {
	goroutineId := GetGoroutineId()
	return goroutineData.RequestClientId[goroutineId]
}

func BindGoroutineRequestId(requestId string, clientRequestId string) {
	goroutineId := GetGoroutineId()

	goroutineData.Lock()
	defer goroutineData.Unlock()

	goroutineData.RequestId[goroutineId] = requestId
	goroutineData.RequestClientId[goroutineId] = clientRequestId
}

func UntieGoroutineRequestId() {
	goroutineId := GetGoroutineId()

	goroutineData.Lock()
	defer goroutineData.Unlock()

	delete(goroutineData.RequestId, goroutineId)
	delete(goroutineData.RequestClientId, goroutineId)
}
