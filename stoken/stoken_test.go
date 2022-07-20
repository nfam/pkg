package stoken

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	type testcase struct {
		Type   Type
		Client string
		Server string
		Code   int
	}
	cases := []testcase{
		{
			Hash,
			"",
			"",
			200,
		},
		{
			Hash,
			"hash",
			"0EuY9I6Pi8wVxq5awFCAHNbc/UKPtfnmXE4W54BzQPo",
			200,
		},
		{
			Hash,
			"hash",
			"fail",
			401,
		},
		{
			Hash,
			"",
			"fail",
			401,
		},
		{
			Time,
			"",
			"",
			200,
		},
		{
			Time,
			"time",
			"time",
			200,
		},
		{
			Time,
			"time",
			"fail",
			401,
		},
	}
	for i, c := range cases {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", Authorization(c.Type, c.Client))

		rec := httptest.NewRecorder()
		h := Handler(c.Type, c.Server)
		handler := h(http.HandlerFunc(ok))
		handler.ServeHTTP(rec, req)

		if rec.Code != c.Code {
			t.Errorf("\ncase %d\n\texpect: %d\n\tactual: %d", i, c.Code, rec.Code)
		}
	}
}

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
