package grabber

import (
	"reflect"
	"testing"
)

func TestSearchTag(t *testing.T) {
	if tag := getTag(); reflect.TypeOf(tag).Kind() != reflect.Slice {
		t.Fail()
	}
}
