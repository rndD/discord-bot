package tracker

import (
	"reflect"
	"testing"
)

func TestFindTasksIds(t *testing.T) {
	actual := FindTasksIds("hello LS-1")
	expect := []string{"LS-1"}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Not eaqual, exepct %v, actual %v", expect, actual)
	}
}

func TestFindManyTasksIds(t *testing.T) {
	actual := FindTasksIds("hello LS-1 and some other SMARTUPSCHOOL-31, tasks tasks HR-30000")
	expect := []string{"LS-1", "SMARTUPSCHOOL-31", "HR-30000"}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Not eaqual, exepct %v, actual %v", expect, actual)
	}
}

func TestFindNoTasksIds(t *testing.T) {
	actual := FindTasksIds("hello 1-1")
	if len(actual) > 0 {
		t.Errorf("Should be empty, actual %v, len %v", actual, len(actual))
	}
}

func TestFindTasksIdsNotUrl(t *testing.T) {
	actual := FindTasksIds("hello http://some/LS-1 and this is url already")
	if len(actual) > 0 {
		t.Errorf("Should be empty, actual %v, len %v", actual, len(actual))
	}
}

func TestFindTasksIdsOnEdge(t *testing.T) {
	actual := FindTasksIds("LS-1 and some other tasks tasks HR-30000")
	expect := []string{"LS-1", "HR-30000"}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Not eaqual, exepct %v, actual %v", expect, actual)
	}
}

func TestNotFindTasksIds(t *testing.T) {
	actual := FindTasksIds("hello")
	if len(actual) > 0 {
		t.Errorf("Should be empty, actual %v, len %v", actual, len(actual))
	}
}
