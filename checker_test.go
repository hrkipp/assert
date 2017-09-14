package assert_test

import (
	"testing"

	"github.com/hrkipp/assert"
)

func TestIsNilWithNilLiteralIsTrue(t *testing.T) {
	assert.That(t, assert.IsNil.Check(nil))
}

func TestIsNilWithIntLiteralIsFalse(t *testing.T) {
	assert.That(t, !assert.IsNil.Check(1))
}

func TestIsNilWithNilInterfaceIsTrue(t *testing.T) {
	assert.That(t, assert.IsNil.Check(assert.Checker(nil)))
}

func TestThatNotInvertsCheck(t *testing.T) {
	assert.That(t, assert.Not(assert.IsNil).Check(nil) != assert.IsNil.Check(nil))
}

func TestThatNotPassesMessage(t *testing.T) {
	assert.That(t, assert.Not(assert.IsNil).Message(nil) == assert.IsNil.Message(nil))
}

func TestThat1Equals1(t *testing.T) {
	assert.That(t, assert.Equals.Check(1, 1))
}

func TestThat1DoesNotEqual2(t *testing.T) {
	assert.That(t, !assert.Equals.Check(1, 2))
}

func TestSameElementsPass(t *testing.T) {
	empty := []string{}
	foo := []string{"foo"}
	foobar := []string{"foo", "bar"}
	barfoo := []string{"bar", "foo"}
	foobaz := []string{"foo", "baz"}
	foobarbaz := []string{"foo", "bar", "baz"}
	assert.That(t, assert.HasSameElementsAs.Check(empty, empty))
	assert.That(t, assert.HasSameElementsAs.Check(foo, foo))
	assert.That(t, assert.HasSameElementsAs.Check(foobar, barfoo))
	assert.That(t, !assert.HasSameElementsAs.Check(foobar, foobaz))
	assert.That(t, !assert.HasSameElementsAs.Check(foobarbaz, barfoo))
	assert.That(t, !assert.HasSameElementsAs.Check(foobar, foobarbaz))
}

func TestContains(t *testing.T) {
	assert.That(t, !assert.Contains.Check([]string{""}, "foo"))
	assert.That(t, assert.Contains.Check([]string{"foo"}, "foo"))
	assert.That(t, assert.Contains.Check([]string{"foo", "bar"}, "foo"))

}
