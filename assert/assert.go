package assert

import (
	"math"
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

// Close64 asserts that two given float64 values, a and b, are close to each
// other within a tolerance of e.
//
// This implements the same "very close with tolerance e" algorithm used by
// Boost Test (http://www.boost.org/doc/libs/1_56_0/libs/test/doc/html/utf/testing-tools/floating_point_comparison.html),
// which works nicely with both very large and very small numbers.
func Close64(t *testing.T, a, b, e float64) bool {
	return True(t, close64Impl(a, b, e),
		"These values were supposed to be close:\n%v\n%v", a, b)
}

func close64Impl(a, b, e float64) bool {
	d := math.Abs(a - b)
	return d == 0 || (d/math.Abs(a) <= e && d/math.Abs(b) <= e)
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
