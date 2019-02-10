package helloworld

import "golang.org/x/sys/unix"

// GetFHelloWorld will get the hello world string
func GetHelloWorld() string {
	return "Hello, World!"
}

// GetUserID gets the ID of the current user
func GetUserID() int {
	return unix.Getuid()
}

func GetAbsValue(i int) int {
	if i > 0 {
		return i
	}
	return -1 * i
}
