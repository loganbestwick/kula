package api

import (
	"fmt"
	"net/http"

	"goji.io/pat"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "hello, %s!", name)
}