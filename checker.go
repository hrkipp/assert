package assert

import (
	"fmt"
	"reflect"
)

type Checker interface {
	Check(args ...interface{}) bool
	Message(args ...interface{}) string
}

// Not

func Not(checker Checker) Checker {
	return not{checker}
}

type not struct {
	checker Checker
}

func (c not) Check(args ...interface{}) bool {
	return !c.checker.Check(args...)
}

func (c not) Message(args ...interface{}) string {
	return c.checker.Message(args...)
}

// IsNil
var IsNil Checker = isNil{}
var IsNotNil = Not(IsNil)

type isNil struct {
}

func (isNil) Check(args ...interface{}) bool {
	if args[0] == nil {
		return true
	}
	switch v := reflect.ValueOf(args[0]); v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice:
		return v.IsNil()
	}
	return false
}

func (isNil) Message(args ...interface{}) string {
	return fmt.Sprintf("\nexpected: <nil>\nobtained: %v", args[0])

}

// Equals
var Equals = equals{}

type equals struct {
}

func (equals) Check(args ...interface{}) bool {
	return args[0] == args[1]
}

func (equals) Message(args ...interface{}) string {
	return fmt.Sprintf("\nobtained: %v\nexpected: %v", args[0], args[1])
}

// Deep Equals
var DeepEquals = deepEquals{}

type deepEquals struct {
}

func (deepEquals) Check(args ...interface{}) bool {
	return reflect.DeepEqual(args[0], args[1])
}

func (deepEquals) Message(args ...interface{}) string {
	return fmt.Sprintf("\nobtained: %v\nexpected: %v", args[0], args[1])
}

// HasSameElementsAs

var HasSameElementsAs = hasSameElementsAs{}

type hasSameElementsAs struct {
}

func (hasSameElementsAs) Check(args ...interface{}) bool {
	val0 := reflect.ValueOf(args[0])
	val1 := reflect.ValueOf(args[1])

	if val0.Len() != val1.Len() {
		return false
	}

outer:
	for i := 0; i < val0.Len(); i++ {
		for j := 0; j < val1.Len(); j++ {
			if reflect.DeepEqual(
				val0.Index(i).Interface(),
				val1.Index(j).Interface()) {
				continue outer
			}
		}
		return false
	}
	return true
}

func (hasSameElementsAs) Message(args ...interface{}) string {
	return fmt.Sprintf("\nobtained: %v\nexpected: %v", args[0], args[1])
}

// Contains
var Contains = contains{}

type contains struct {
}

func (contains) Check(args ...interface{}) bool {
	val0 := reflect.ValueOf(args[0])
	for i := 0; i < val0.Len(); i++ {
		if reflect.DeepEqual(val0.Index(i).Interface(), args[1]) {
			return true
		}
	}
	return false
}

func (contains) Message(args ...interface{}) string {
	return fmt.Sprintf("\nobtained: %v\nneeded:  %v", args[0], args[1])
}
