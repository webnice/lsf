package lsf

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestNew(t *testing.T) {
	var nbj Interface
	var qfo *impl
	var ok bool

	if nbj = New(); nbj == nil {
		t.Fatalf("New return nil object")
		return
	}
	if qfo, ok = nbj.(*impl); !ok {
		t.Fatalf("New return wrong object")
		return
	}
	if qfo.filter == nil {
		t.Fatalf("New not created filter slice")
		return
	}
	if qfo.order == nil {
		t.Fatalf("New not created oreder slice")
		return
	}
}
