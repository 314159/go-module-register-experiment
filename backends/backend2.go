package backends

import "fmt"

// example of a backend that doesn't require its own package

type backend2 struct{}

const backend2Name = "backend2"

func init() {
	Register(backend2Name, &backend2{})
}

func (b backend2) Name() string {
	return backend2Name
}

func (b backend2) Process() error {
	fmt.Printf("Hello from %s\n", b.Name())
	return nil
}
