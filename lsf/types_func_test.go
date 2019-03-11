package lsf

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestSortOrderString(t *testing.T) {
	if SortAscending.String() != "asc" {
		t.Fatalf("error in String() function in type SortOrder")
		return
	}
}

func TestEquateString(t *testing.T) {
	if EqEqual.String() != "eq" {
		t.Fatalf("error in String() function in type Equate")
		return
	}
}
