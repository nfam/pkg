package crc32c

import "testing"

func TestHash(t *testing.T) {
	var expect = "86a072c0"
	v := Hash([]byte("test"))
	if v != expect {
		t.Errorf("\nexpect: %s\nactual: %s", expect, v)
	}
}
