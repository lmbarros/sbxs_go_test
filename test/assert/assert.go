package assert

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

// True asserts that a given Boolean value is true.
func True(t *testing.T, value bool, format string, args ...interface{}) bool {
	if !value {
		t.Errorf(errorInfo()+"\n"+format, args...)
		return false
	}

	return true
}

// Equal asserts that two given values, a and b, are equal. reflect.DeepEqual is
// used to determine if the values are equal.
func Equal(t *testing.T, a, b interface{}) bool {
	return True(t, equalImpl(a, b),
		"These values were supposed to be equal:\n%v\n%v", a, b)
}

func equalImpl(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// errorInfo returns a string with some useful information about the error.
//
// That hardcoded index to the call stack is really ugly: it implies in strong
// assumptions on where and how this function is called.
func errorInfo() string {

	info := ""

	// Location of the failing assertion
	_, file, line, ok := runtime.Caller(3)
	if ok {
		info = "Error at: "
		info += filepath.Base(file) + ":" + strconv.Itoa(line) + "\n"
	}

	return info
}
