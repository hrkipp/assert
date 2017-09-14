package assert_test

import (
	"testing"

	"github.com/hrkipp/assert"
)

func TestCheckFailsWithNoArgs(t *testing.T) {
	passes, message := assert.Check()
	if passes {
		t.Fatal("shouldn't have passed with no arguments")
	}
	if message != assert.NotEnoughArgs {
		t.Fatal("should have given us the correct error message")
	}
}

func TestCheckPassesOnTrueBools(t *testing.T) {
	passes, message := assert.Check(true)
	if !passes {
		t.Fatal("should have passed with a true argument")
	}
	if message != "" {
		t.Fatal("there shouldn't have been any error message")
	}
}

func TestCheckFailsOnFalseBools(t *testing.T) {
	passes, message := assert.Check(false)
	if passes {
		t.Fatal("shouldn't have passed with a false argument")
	}
	if message != assert.FalseAssertion {
		t.Fatal("should have given us the correct error message")
	}
}

func TestCheckFailsWhenSecondArgumentIsntAChecker(t *testing.T) {
	passes, message := assert.Check("foo", "bar")
	if passes {
		t.Fatal("shouldn't have passed without the second argument being a checker")
	}
	if message != assert.NoChecker {
		t.Fatal("should have given us the correct error message")
	}
}

type checker struct {
	passes      bool
	message     string
	checkargs   []interface{}
	messageargs []interface{}
}

func (c *checker) Check(args ...interface{}) bool {
	c.checkargs = args
	return c.passes
}

func (c *checker) Message(args ...interface{}) string {
	c.messageargs = args
	return c.message
}

func TestCheckPassesWhenCheckerPasses(t *testing.T) {
	mock := &checker{
		passes:  true,
		message: "should't see me",
	}
	passes, message := assert.Check("foo", mock)
	if !passes {
		t.Fatal("should have passed")
	}
	if message != "" {
		t.Fatal("shouldn't have been an error message")
	}
}

func TestCheckFailsWhenCheckerFails(t *testing.T) {
	mock := &checker{
		passes:  false,
		message: "failure message",
	}
	passes, message := assert.Check("foo", mock)
	if passes {
		t.Fatal("shouldn't have passed")
	}
	if message != "failure message" {
		t.Fatal("shouldn't have been an error message")
	}
	if len(mock.checkargs) != 1 {
		t.Fatal("should have gotten one argument")
	}
	if len(mock.messageargs) != 1 {
		t.Fatal("should have gotten one argument")
	}
	if mock.checkargs[0].(string) != "foo" {
		t.Fatal("wrong argument passed to check")
	}
	if mock.messageargs[0].(string) != "foo" {
		t.Fatal("wrong argument passed to check")
	}
}

func TestCheckPassesCorrectArgs(t *testing.T) {
	mock := &checker{
		passes:  false,
		message: "failure message",
	}
	assert.Check("foo", mock, "bar")
	if len(mock.checkargs) != 2 {
		t.Fatal("should have gotten one argument")
	}
	if len(mock.messageargs) != 2 {
		t.Fatal("should have gotten one argument")
	}
	if mock.checkargs[0].(string) != "foo" {
		t.Fatal("wrong argument passed to check")
	}
	if mock.messageargs[0].(string) != "foo" {
		t.Fatal("wrong argument passed to check")
	}
	if mock.checkargs[1].(string) != "bar" {
		t.Fatal("wrong argument passed to check")
	}
	if mock.messageargs[1].(string) != "bar" {
		t.Fatal("wrong argument passed to check")
	}
}
