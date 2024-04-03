package gomodone

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	version  = "test"
	nameRepo = "gomodone"
)

func F1() {
	fmt.Printf("This function F1 from %v version: %v\n", nameRepo, version)
}

func F2() {
	fmt.Println("testing version")
}

func NewId() string {
	uuid := uuid.New()
	return uuid.String()
}
