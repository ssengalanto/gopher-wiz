package main

import "testing"

func TestHello(t *testing.T) {
	s := "Gopher Wiz"
	got := hello(s)
	want := "Hello Gopher Wiz!"

	if got != want {
		t.Errorf("\nfn: hello(%q)\ngot: %q\nwant: %q\n", s, got, want)
	}
}
