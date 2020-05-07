package main

import (
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	timestamp := time.Now().UnixNano()
	t.Log(timestamp)
}
