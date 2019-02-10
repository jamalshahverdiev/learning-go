package helloworld

import (
	"testing"

	"golang.org/x/sys/unix"
)

func TestGetHelloWorld(t *testing.T) {
	retval := GetHelloWorld()
	if retval != "Hello, World!" {
		t.Fail()
	}
}

func TestGetUserID(t *testing.T) {
	retid := GetUserID()
	actualid := unix.Getuid()
	if retid != actualid {
		t.Fail()
	}
}

func TestGetAbsValue(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	retval := GetAbsValue(-5)
	if retval != 5 {
		t.Fail()
	}
	retval = GetAbsValue(10)
	if retval != 10 {
		t.Fail()
	}
}
