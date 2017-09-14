package assert

import "testing"

var CheckOverride = Check
var That = Assert

func Assert(t testing.TB, args ...interface{}) {
	t.Helper()
	passes, message := CheckOverride(args...)
	if !passes {
		t.Fatal(message)
	}
}
