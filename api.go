package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func UpdateState(w http.ResponseWriter, r *http.Request) {
	data, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(data))
	fmt.Fprint(w, "hi back")
}
