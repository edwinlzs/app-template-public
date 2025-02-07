package utils

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Err  error
	Code int
}

// satisfies `error` interface
func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}
