package exloggo

import "github.com/google/uuid"

func getUUID() string {
	return uuid.New().String()
}
