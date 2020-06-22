package helpers

// Error defines Error struct
type Error struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *Error) Error() string {
	return error.msg
}
