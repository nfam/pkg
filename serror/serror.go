package serror

// Error represents a status error.
type Error struct {
	Text string `json:"error"`
	Code int    `json:"-"`
}

// Error implement error interface.
func (e *Error) Error() string {
	return e.Text
}

// New returns a new error.
func New(text string, code int) *Error {
	return &Error{text, code}
}

// BadRequest returns a new bad request error.
func BadRequest(text string) *Error {
	return New(text, 400)
}

// Internal returns a new internal server error.
func Internal(text string) *Error {
	return New(text, 500)
}
