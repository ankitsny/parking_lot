package operations

import (
	"errors"
	"strings"

	"github.com/anks333/parking_lot/store"
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
	opm.commands = make(map[string]ICommand)
	// Register all operations here
	opm.AddOperation(NewCreateParkingLot())
	opm.AddOperation(NewParkCMD())
	opm.AddOperation(NewLeaveCMD())
	opm.AddOperation(NewRegistrationNoBasedOnColorCMD())
	opm.AddOperation(NewSlotNoBasedOnColorCMD())
	opm.AddOperation(NewSlotNoBasedOnRegNoCMD())
	opm.AddOperation(NewStatusCMD())
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
	opm.cOp = ""
	if input == "" {
		return errors.New("Input args missing")
	}
	argsArray := strings.Fields(input)
	opm.cOp = argsArray[0]
	restArgs := argsArray[1:]
	opm.argVal = strings.Join(restArgs, " ")
	return nil
}

// Execute :
func (opm *OperationManager) Execute(input string) string {
	// TODO: implement this method
	if err := opm.Parse(input); err != nil {
		return err.Error()
	}
	cmd, ok := opm.commands[opm.cOp]
	if !ok {
		return "Invalid Operation"
	}

	if opm.cOp != "create_parking_lot" && store.GetStorage() == nil {
		return "Parking lot is not created"
	}
	return cmd.Execute(opm.argVal)
}
