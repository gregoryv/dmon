package main

import (
	"net/http"
	"testing"
)

func Test_main(t *testing.T) {
	go main()
	resp, err := http.Get("http://localhost:7171/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode >= 400 {
		t.Error(resp.Status)
	}
}
