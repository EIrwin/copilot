package copilot

import "fmt"

const (
	ErrorInvalidRunArgument = "invalid argument provided"
)

type ErrInvalidRunArgument struct {
	message string
}

func (e ErrInvalidRunArgument) Error() string {
	return fmt.Sprintf("copilot: %v %v", ErrorInvalidRunArgument, e.message)
}

func NewErrInvalidRunArgument(msg string) ErrInvalidRunArgument {
	return ErrInvalidRunArgument{
		message: msg,
	}
}
