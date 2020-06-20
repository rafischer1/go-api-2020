package main

// MyError is used to Define Error struct
type MyError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *MyError) Error() string {
	return error.msg
}
