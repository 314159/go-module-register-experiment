package handlers

// Another way is to simply register them here; most of the same benefits

import (
	"fmt"

	"github.com/314159/go-module-register-experiment/handlers/handler1"
	"github.com/314159/go-module-register-experiment/handlers/handler2"
	"github.com/pkg/errors"
)

type Handler interface {
	Name() string
	Process() error
}

var registry = make(map[string]Handler)

func Register(name string, backend Handler) {
	registry[name] = backend
}

func GetHandler(name string) (Handler, error) {
	if backend, ok := registry[name]; ok {
		return backend, nil
	}

	return nil, errors.New(fmt.Sprintf("unknown backend %s", name))
}

func init() {
	Register(handler1.HandlerName, handler1.New())
	Register(handler2.HandlerName, handler2.New())
}
