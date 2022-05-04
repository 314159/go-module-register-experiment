package backends

import "fmt"

// example of a backend that doesn't require its own package

func init() {
	Register("backend2", New())
}

type backend2 struct{}

func New() Backend {
	return &backend2{}
}

func (b backend2) Name() string {
	return "backend2"
}

func (b backend2) Process() error {
	fmt.Printf("Hello from %s\n", b.Name())
	return nil
}
