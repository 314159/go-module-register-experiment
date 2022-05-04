package main

import (
	"fmt"

	"github.com/314159/go-module-register-experiment/backends"
)

func main() {
	fmt.Println("Hello, world")
	backend1, err := backends.GetBackend("backend1")
	if err != nil {
		panic(err.Error())
	}
	backend2, err := backends.GetBackend("backend2")
	if err != nil {
		panic(err.Error())
	}
	if err := backend1.Process(); err != nil {
		panic(err.Error())
	}
	if err := backend2.Process(); err != nil {
		panic(err.Error())
	}
}
