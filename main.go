package gomodone

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	version  = "v0.1.1"
	nameRepo = "gomodone"
)

func F1() {
	fmt.Printf("This function F1 from %v version: %v", nameRepo, version)
}

func NewId() string {
	uuid := uuid.New()
	return uuid.String()
}
