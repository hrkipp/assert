package assert_test

import (
	"testing"

	"github.com/hrkipp/assert"
)

type mockT struct {
	testing.TB
	count   int
	message string
}

func (m *mockT) Fatal(args ...interface{}) {
	m.count++
	m.message = args[0].(string)
}

func (*mockT) Helper() {
}

func TestAssertsPass(t *testing.T) {
	assert.CheckOverride = func(args ...interface{}) (bool, string) {
		return true, "should't be set"
	}

	tmock := mockT{}
	assert.Assert(&tmock, "foo", "bar")

	if tmock.count != 0 {
		t.Fatal("checker was called")
	}
	if tmock.message != "" {
		t.Fatal("failure messsage should be empty")
	}
	assert.CheckOverride = assert.Check

}

func TestAssertsFail(t *testing.T) {
	assert.CheckOverride = func(args ...interface{}) (bool, string) {
		if len(args) != 2 {
			t.Fatal("not passing the right number of args")
		}
		if args[0] != "foo" {
			t.Fatal("not passing the right args")
		}
		if args[1] != "bar" {
			t.Fatal("not passing the right args")
		}
		return false, "fail message"
	}

	tmock := mockT{}
	assert.Assert(&tmock, "foo", "bar")
	if tmock.count != 1 {
		t.Fatal("checker wasn't called")
	}
	if tmock.message != "fail message" {
		t.Fatal("wrong failure method")
	}
	assert.CheckOverride = assert.Check

}
