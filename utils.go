package exloggo

import uuid "github.com/uuid6/uuid6go-proto"

var uuidGenerator = uuid.UUIDv7Generator{CounterPrecisionLength: 12}

func GetUUIDv7() string {
	return uuidGenerator.Next().ToString()
}
