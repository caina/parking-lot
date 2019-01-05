package initialization_test

import (
	"github.com/parking-lot/app/commands/initialization"
	"strings"
	"testing"
)

func TestInitialization(t *testing.T) {
	lot, _ := initialization.Command(strings.Fields("2"))
	if lot.GetSlots() != 2 {
		t.Error("should had 2 slots")
	}
}

func TestShouldGiveErrorIfNoNumber(t *testing.T) {
	if _, errors := initialization.Command(strings.Fields("two")); errors == nil {
		t.Error("should return error if argument is not a number")
	}
}
