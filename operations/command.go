package operations

// ICommand :
type ICommand interface {
	Parse(argsVal string) error
	Execute(string) string
	GetName() string
}
