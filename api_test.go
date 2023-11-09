package main

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateState(t *testing.T) {
	r := httptest.NewRequest("POST", "/", strings.NewReader("hi"))
	w := httptest.NewRecorder()

	UpdateState(w, r)

	resp := w.Result()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := string(data)
	exp := "hi back"
	if got != exp {
		t.Errorf("got %q, expected %s", got, exp)
	}
}
