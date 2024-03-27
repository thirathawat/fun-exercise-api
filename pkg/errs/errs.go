package errs

type Err struct {
	Message string `json:"message"`
}

func (e *Err) Error() string {
	return e.Message
}

func New(message string) *Err {
	return &Err{Message: message}
}
