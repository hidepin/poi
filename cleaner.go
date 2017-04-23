package poi

import "fmt"

type Cleaner interface {
	Clean(path string) error
}

type noexec struct{}

var NOEXEC Cleaner = (*noexec)(nil)

func (n *noexec) Clean(path string) error {
	fmt.Println("cleaner: " + path)

	return nil
}
