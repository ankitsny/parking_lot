package operations

import (
	"errors"
	"strings"
)

// OperationManager :
type OperationManager struct {
	cOp      string              // current operations
	argVal   string              // current args
	commands map[string]ICommand // map of commands
}

// NewOperationManager :
func NewOperationManager() *OperationManager {
	opm := &OperationManager{}

	return opm
}

// IOperation :
type IOperation interface {
	Parse(input string) error
	Execute() (string, error)
	AddOperation(ICommand)
}

// AddOperation :
func (opm *OperationManager) AddOperation(cmd ICommand) {
	opm.commands[cmd.GetName()] = cmd
}

// Parse :
func (opm *OperationManager) Parse(input string) error {
	if input == "" {
		return errors.New("Input args missing")
	}
	argsArray := strings.Fields(input)
	opm.cOp = argsArray[0]
	if len(argsArray) == 2 {
		opm.argVal = argsArray[1]
	}
	return nil
}

// Execute :
func (opm *OperationManager) Execute(input string) (string, error) {
	// TODO: implement this method
	if err := opm.Parse(input); err != nil {
		return "", err
	}
	cmd, ok := opm.commands[opm.cOp]
	if !ok {
		return "", errors.New("Invalid Operation")
	}
	return cmd.Execute(opm.argVal)
}
