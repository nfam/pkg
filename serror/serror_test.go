package serror

import "testing"

func TestNew(t *testing.T) {
	const text = "new"
	e := New(text, 1)
	if e.Error() != text {
		t.Errorf("\nexpect: %v\nactual: %v", e.Error(), text)
		return
	}
	if e.Code != 1 {
		t.Errorf("\nexpect: %v\nactual: %v", e.Code, 1)
		return
	}
}

func TestBadRequest(t *testing.T) {
	const text = "bad request"
	e := BadRequest(text)
	if e.Error() != text {
		t.Errorf("\nexpect: %v\nactual: %v", e.Error(), text)
		return
	}
	if e.Code != 400 {
		t.Errorf("\nexpect: %v\nactual: %v", e.Code, 400)
		return
	}
}

func TestInternal(t *testing.T) {
	const text = "new"
	e := Internal(text)
	if e.Error() != text {
		t.Errorf("\nexpect: %v\nactual: %v", e.Error(), text)
		return
	}
	if e.Code != 500 {
		t.Errorf("\nexpect: %v\nactual: %v", e.Code, 500)
		return
	}
}
