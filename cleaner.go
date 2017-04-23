package poi

import (
	"fmt"
	"os"
)

type Cleaner interface {
	Clean(path string) error
}

type noexec struct{}
type exec struct{}

var NOEXEC Cleaner = (*noexec)(nil)
var EXEC Cleaner = (*exec)(nil)

func (n *noexec) Clean(path string) error {
	fmt.Println("noexec delete: " + path)

	return nil
}

func (n *exec) Clean(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	fmt.Println("deleted : " + path)
	return nil
}
