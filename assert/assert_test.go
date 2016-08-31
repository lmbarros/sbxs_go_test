package assert

import "testing"

// Tests if true is really recognized as true
func TestTrue(t *testing.T) {
	if !True(t, true, "Ooops") {
		t.Error("True should have returned true")
	}
}

// Tests if equal things are recognized as equal.
func TestEqual(t *testing.T) {
	// Test the positive cases with the real, exported function
	if !Equal(t, 1234, 1234) {
		t.Error("Equal should have returned true")
	}

	if !Equal(t, 3.14, 3.14) {
		t.Error("Equal should have returned true")
	}

	if !Equal(t, uint(12345678901234567890), uint(12345678901234567890)) {
		t.Error("Equal should have returned true")
	}

	if !Equal(t, true, true) {
		t.Error("Equal should have returned true")
	}

	if !Equal(t, "Same thing here", "Same thing here") {
		t.Error("Equal should have returned true")
	}

	array12345a := [...]int{1, 2, 3, 4, 5}
	array12345b := [...]int{1, 2, 3, 4, 5}
	if !Equal(t, array12345a, array12345b) {
		t.Error("Equal should have returned true")
	}

	slice123a := array12345a[0:4]
	slice123b := array12345a[0:4]
	if !Equal(t, slice123a, slice123b) {
		t.Error("Equal should have returned true")
	}

	// Test the negative cases through equalImpl, which is the best I can do
	if equalImpl(1234, -1234) {
		t.Error("equalImpl should have returned false")
	}

	if equalImpl(3.14, 3.14159) {
		t.Error("equalImpl should have returned false")
	}

	if equalImpl(uint(12345678901234567890), uint(12345678901234567891)) {
		t.Error("equalImpl should have returned false")
	}

	if equalImpl(true, false) {
		t.Error("equalImpl should have returned false")
	}

	if equalImpl("Something here", "Something else here") {
		t.Error("equalImpl should have returned false")
	}

	array123456 := [...]int{1, 2, 3, 4, 5, 6}
	if equalImpl(array12345a, array123456) {
		t.Error("equalImpl should have returned false")
	}

	array12346 := [...]int{1, 2, 3, 4, 6}
	if equalImpl(array12345a, array12346) {
		t.Error("equalImpl should have returned false")
	}

	slice234 := array12345a[1:5]
	slice1234 := array12345a[0:5]
	if equalImpl(slice123a, slice234) {
		t.Error("equalImpl should have returned false")
	}
	if equalImpl(slice123a, slice1234) {
		t.Error("equalImpl should have returned false")
	}
}
