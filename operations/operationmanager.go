package operations

// Operation :
type Operation struct {
	cOp      string                 // current operations
	argVal   string                 // current args
	commands map[string]interface{} // map of commands
}

// NewOperation :
func NewOperation() IOperation {
	return &Operation{}
}

// IOperation :
type IOperation interface {
	Parse(input string) error
	Execute() (string, error)
}

// Parse :
func (op *Operation) Parse(input string) error {
	// TODO: implement parse method
	op.cOp = "SOME COmmand"
	op.argVal = "SOME Arguments"
	return nil
}

// Execute :
func (op *Operation) Execute() (string, error) {
	// TODO: implement this method
	return "", nil
}
