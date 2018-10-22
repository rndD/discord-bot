package main

import (
	"reflect"
	"testing"
)

func TestFindTasksIds(t *testing.T) {
	actual := findTasksIds("hello LS-1")
	expect := []string{"LS-1"}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Not eaqual, exepct %v, actual %v", expect, actual)
	}
}

func TestFindManyTasksIds(t *testing.T) {
	actual := findTasksIds("hello LS-1 and some other SMARTUPSCHOOL-31, tasks tasks HR-30000")
	expect := []string{"LS-1", "SMARTUPSCHOOL-31", "HR-30000"}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Not eaqual, exepct %v, actual %v", expect, actual)
	}
}

// func TestFindTasksIdsNotUrl(t *testing.T) {
// 	actual := findTasksIds("hello http://some/1LS-1 and this is url already")
// 	if len(actual) > 0 {
// 		t.Errorf("Should be empty, actual %v, len %v", actual, len(actual))
// 	}
// }

func TestNotFindTasksIds(t *testing.T) {
	actual := findTasksIds("hello")
	if len(actual) > 0 {
		t.Errorf("Should be empty, actual %v, len %v", actual, len(actual))
	}
}
