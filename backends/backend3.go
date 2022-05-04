package backends

import "fmt"

// example of a backend that doesn't require its own package

type backend3 struct{}

const backend3Name = "backend3"

func init() {
	Register(backend3Name, &backend3{})
}

func (b backend3) Name() string {
	return backend3Name
}

func (b backend3) Process() error {
	fmt.Printf("Hello from %s\n", b.Name())
	return nil
}
