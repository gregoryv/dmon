package main

import (
	"fmt"
	"net/http"
)

func UpdateState(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi back")
}
