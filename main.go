package main

import (
	"fmt"

	"github.com/314159/go-module-register-experiment/backends"
	"github.com/314159/go-module-register-experiment/handlers"
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
	backend3, err := backends.GetBackend("backend3")
	if err != nil {
		panic(err.Error())
	}
	if err := backend1.Process(); err != nil {
		panic(err.Error())
	}
	if err := backend2.Process(); err != nil {
		panic(err.Error())
	}
	if err := backend3.Process(); err != nil {
		panic(err.Error())
	}
	if _, err := backends.GetBackend("unknown"); err == nil {
		panic("Oops, request for unknown backend succeeded!")
	}
	handler1, err := handlers.GetHandler("handler1")
	if err != nil {
		panic(err.Error())
	}
	handler1.Process()

	handler2, err := handlers.GetHandler("handler2")
	if err != nil {
		panic(err.Error())
	}
	handler2.Process()

}
