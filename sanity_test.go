package scream

import (
	"testing"
)

func TestSanityCheck(t *testing.T) {
	err := SanityCheck()
	if err != nil {
		t.Error(err.Error())
	}
}
