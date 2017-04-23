package poi

import "fmt"

type Cleaner interface {
	Clean(path string) error
}

type noexec struct{}
type exec struct{}

var NOEXEC Cleaner = (*noexec)(nil)
var EXEC Cleaner = (*exec)(nil)

func (n *noexec) Clean(path string) error {
	fmt.Println("noexec cleaner: " + path)

	return nil
}

func (n *exec) Clean(path string) error {
	fmt.Println("exec cleaner: " + path)

	return nil
}
