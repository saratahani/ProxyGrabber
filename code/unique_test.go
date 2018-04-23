package code

import (
	"testing"
)

func TestUnique(t *testing.T) {

	testString := []string{`000.00.000.000:0000`, `000.00.000.000:0000`, `001.01.001.001:0001`}

	if val := Unique(testString); len(val) != 2 {
		t.Fail()
	}
}
