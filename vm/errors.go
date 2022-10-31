package vm

var (
	NoEntryError = NewRuntimeError("Cannot find entry function")
)

type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

type RuntimeError struct {
	message string
}

func (e RuntimeError) Error() string {
	return e.message
}

func NewRuntimeError(message string) RuntimeError {
	return RuntimeError{message}
}
