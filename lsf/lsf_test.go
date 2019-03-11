package lsf

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"net/url"
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

func TestBuild(t *testing.T) {
	const (
		baseURL     = `https://example.com/api/v1.0/administration/action/method`
		expectation = `https://example.com/api/v1.0/administration/action/method?filter=id%3Age%3A1&filter=name%3Ane%3A&filter=name%3Ake%3A%2Aion&limit=0%3A2&order=updateAt%3Adesc&order=createAt%3Aasc&order=name%3Aasc&order=id%3Aasc`
	)
	var err error
	var uri *url.URL
	var qfo Interface

	if uri, err = url.Parse(baseURL); err != nil {
		t.Fatalf("net/url.Parse(%q) error: %s", baseURL, err)
		return
	}
	qfo = New().
		Limit(0, 2).
		Order(`updateAt`, SortDescending).
		Order(`createAt`, SortAscending).
		Order(`name`, SortAscending).
		Order(`id`, SortAscending).
		Filter(`id`, EqGreaterEqual, "1").
		Filter(`name`, EqNotEqual, "").
		Filter(`name`, EqLikeThan, "*ion")
	if uri, err = qfo.URL(uri); err != nil {
		t.Fatalf("create URL error: %s", err)
		return
	}
	if uri.String() != expectation {
		t.Fatalf("expected %q, return %q", expectation, uri.String())
		return
	}
}
