package code

import (
	"reflect"
	"testing"
)

func TestSearchTag(t *testing.T) {
	if tag := GetTag(); reflect.TypeOf(tag).Kind() != reflect.Slice {
		t.Fail()
	}
}
