package backend1

import "fmt"

func (t backend1Backend) Process() error {
	fmt.Printf("Hello, from %s\n", t.Name())
	return nil
}
