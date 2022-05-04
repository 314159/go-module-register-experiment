package handler1

import "fmt"

type handler1 struct{}

const HandlerName = "handler1"

func New() *handler1 {
	return &handler1{}
}
func (b handler1) Name() string {
	return HandlerName
}

func (b handler1) Process() error {
	fmt.Printf("Hello from handler %s\n", b.Name())
	return nil
}
