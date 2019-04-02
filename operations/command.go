package operations

// ICommand :
type ICommand interface {
	Parse(argsVal string) error
	Execute(string) (string, error)
	GetName() string
}
