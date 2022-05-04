package backends

// Example of a backend that requires its own package

import "github.com/314159/go-module-register-experiment/backends/backend1"

func init() {
	Register("backend1", backend1.New())
}
