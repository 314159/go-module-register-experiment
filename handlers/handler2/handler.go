package handler2

import "fmt"

type handler2 struct{}

const HandlerName = "handler2"

func New() *handler2 {
	return &handler2{}
}
func (b handler2) Name() string {
	return HandlerName
}

func (b handler2) Process() error {
	fmt.Printf("Hello from handler %s\n", b.Name())
	return nil
}
