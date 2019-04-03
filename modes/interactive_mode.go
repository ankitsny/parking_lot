package modes

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anks333/parking_lot/operations"
)

// InteractiveMode :
type InteractiveMode struct {
}

// NewInterActiveMode :
func NewInterActiveMode() *InteractiveMode {
	return &InteractiveMode{}
}

// Start :
func (im *InteractiveMode) Start() {
	reader := bufio.NewReader(os.Stdin)
	opm := operations.NewOperationManager()
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		out := opm.Execute(input)
		fmt.Println(out)
	}
}
