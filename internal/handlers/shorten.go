package handlers

import "fmt"
import "net/http"

func Hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w, "hello\n")
}