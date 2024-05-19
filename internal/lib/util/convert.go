package util

import "time"

func ToTimeDurationSec[T uint | int](value T) time.Duration {
	return time.Duration(value) * time.Second
}

func WithPointer[T any](value T) *T {
	return &value
}
