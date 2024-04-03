package mod3

import (
	"fmt"

	"github.com/google/uuid"
)

var version = "v2.0.0"

func F1() {
	fmt.Println("This function F1 from mod3 version: ", version)
}

func F2() {
	fmt.Println("This function F2 from mod3 version:", version)
}

func getId() string {
	uuid := uuid.New()
	return uuid.String()
}

func main() {
	fmt.Println("Run mod3 main function")
}
