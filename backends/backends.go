package backends

import (
	"fmt"

	"github.com/pkg/errors"
)

type Backend interface {
	Name() string
	Process() error
}

var registeredBackends = make(map[string]Backend)

func Register(name string, backend Backend) {
	registeredBackends[name] = backend
}

func GetBackend(name string) (Backend, error) {
	if backend, ok := registeredBackends[name]; ok {
		return backend, nil
	}

	return nil, errors.New(fmt.Sprintf("unknown backend %s", name))
}
