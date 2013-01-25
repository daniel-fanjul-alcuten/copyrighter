package main

import (
	"testing"
)

func TestReplace(t *testing.T) {
	assertReplace(t, "Copyright 2009\nCopyright 2010 Spotify\nCopyright 2011-2012 Spotify AB\n", "Copyright 2009\nCopyright 2013 Spotify\nCopyright 2011-2013 Spotify AB\n")
}

func assertReplace(t *testing.T, source, target string) {
	c := NewCopyrigher("Spotify", "2013")
	v := c.Replace(source)
	if v != target {
		t.Errorf("Expected: '%v', Found: '%v'", target, v)
	}
}
