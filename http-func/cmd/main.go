package main

import (
	"log"
	"net/http"

	function "github.com/yuta-ron/my-cloudfunction-tutorial/http-func"
)

func main() {
	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	//

	http.HandleFunc("/", function.MyFunctionTutorial)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
