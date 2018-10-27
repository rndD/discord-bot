package gitlab

import (
	"reflect"
	"testing"
)

func TestFindMrIds(t *testing.T) {
	actual := FindMRIds("hello !42")
	expect := []string{"42"}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Not eaqual, exepct %v, actual %v", expect, actual)
	}
}

func TestNotFindMrIds(t *testing.T) {
	actual := FindMRIds("hello 42")
	if len(actual) > 0 {
		t.Errorf("Should be empty, actual %v, len %v", actual, len(actual))
	}
}

func TestNotFindMrIds2(t *testing.T) {
	actual := FindMRIds("hello 42!")
	if len(actual) > 0 {
		t.Errorf("Should be empty, actual %v, len %v", actual, len(actual))
	}
}

func TestGetTextForMr(t *testing.T) {
	actual := GetTextForMr([]string{"51", "42"})
	expect := MR_URL + "51" + "\n" + MR_URL + "42"
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Not eaqual, exepct %v, actual %v", expect, actual)
	}
}
