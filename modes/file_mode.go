package modes

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anks333/parking_lot/operations"
)

// FileMode :
type FileMode struct {
	filePath string
}

// NewFileMode :
func NewFileMode(path string) *FileMode {
	return &FileMode{
		filePath: path,
	}
}

// Start :
func (fm *FileMode) Start() {
	f, err := os.Open(fm.filePath)
	if err != nil {
		fmt.Println("Failed open file")
		return
	}
	defer f.Close()

	inpScanner := bufio.NewScanner(f)
	opm := operations.NewOperationManager()
	for inpScanner.Scan() {
		input := inpScanner.Text()
		out := opm.Execute(input)
		fmt.Println(out)
	}
}
