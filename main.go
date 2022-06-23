package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}

//docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html