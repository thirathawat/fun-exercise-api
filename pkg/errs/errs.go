package errs

var ErrNotFound = New("not found")

type Err struct {
	Message string `json:"message"`
}

func (e *Err) Error() string {
	return e.Message
}

func New(message string) *Err {
	return &Err{Message: message}
}
