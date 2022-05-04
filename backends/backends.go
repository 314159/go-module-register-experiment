package backends

import (
	"fmt"

	"github.com/pkg/errors"
)

type Backend interface {
	Name() string
	Process() error
}

var registry = make(map[string]Backend)

func Register(name string, backend Backend) {
	registry[name] = backend
}

func GetBackend(name string) (Backend, error) {
	if backend, ok := registry[name]; ok {
		return backend, nil
	}

	return nil, errors.New(fmt.Sprintf("unknown backend %s", name))
}
