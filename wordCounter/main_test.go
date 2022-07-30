package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Word1 Word2 Word3 Word4 \n")
	exp := 4
	res := count(b)

	if res != exp {
		t.Errorf("Expected %d got %d instead.\n", exp, res)
	}
}
