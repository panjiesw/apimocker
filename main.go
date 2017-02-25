package main

import "github.com/panjiesw/apimocker/server"
import "net/http"

func main() {
	s := server.New()
	http.ListenAndServe(":3000", s)
}
