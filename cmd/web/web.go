package main

import (
	"fmt"

	"github.com/simple-distributed-project/web"
)

func main() {
	fmt.Println("Starting service!")
	// web.StartService()
	web.StartGRPCServer()
}
