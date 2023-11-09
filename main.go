package main

import "net/http"

func main() {
	http.HandleFunc("/", UpdateState)

	http.ListenAndServe(":7171", nil)
}
