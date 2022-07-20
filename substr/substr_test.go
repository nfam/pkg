package substr

import "testing"

func TestSub(t *testing.T) {
	type testcase struct {
		text   string
		trail  string
		head   []string
		expect string
	}
	var cases = []testcase{
		{
			"abcdef",
			"",
			nil,
			"abcdef",
		},
		{
			"abcdef",
			"c",
			nil,
			"ab",
		},
		{
			"abcdef",
			"",
			[]string{"a", "c"},
			"def",
		},
		{
			"abcdef",
			"f",
			[]string{"a", "c"},
			"de",
		},
		{
			"abcdef",
			"x",
			nil,
			"",
		},
		{
			"abcdef",
			"",
			[]string{"x"},
			"",
		},
	}
	for i, c := range cases {
		v := Sub(c.text, c.trail, c.head...)
		if v != c.expect {
			t.Errorf("\ncase %d:\n\texpected: %s\n\tactual: %s", i, c.expect, v)
		}
	}
}

func TestRSub(t *testing.T) {
	type testcase struct {
		text   string
		head   string
		trail  []string
		expect string
	}
	var cases = []testcase{
		{
			"abcdef",
			"",
			nil,
			"abcdef",
		},
		{
			"acbcdef",
			"c",
			nil,
			"def",
		},
		{
			"abcdefc",
			"",
			[]string{"f", "c"},
			"ab",
		},
		{
			"abcdefc",
			"a",
			[]string{"f", "c"},
			"b",
		},
		{
			"abcdef",
			"x",
			nil,
			"",
		},
		{
			"abcdef",
			"",
			[]string{"x"},
			"",
		},
	}
	for i, c := range cases {
		v := RSub(c.text, c.head, c.trail...)
		if v != c.expect {
			t.Errorf("\ncase %d:\n\texpected: %s\n\tactual: %s", i, c.expect, v)
		}
	}
}
