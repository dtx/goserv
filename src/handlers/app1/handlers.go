package app1

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello, world\n")
}
