package wrap

type Error struct {
	Cause   *Error
	Message string
	Extras  []interface{}
}

func (that *Error) Error() string {
	return that.Message
}

var UnknownError = &Error{
	Message: "unknown error",
	Cause:   nil,
	Extras:  nil,
}
