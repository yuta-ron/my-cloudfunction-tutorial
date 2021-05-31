package functions

import (
	"fmt"
	"log"
	"net/http"
)

// F is sample functions
func MyFunctionTutorial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Println("Hello by fmt pkg")
	log.Println("Hello by log pkg")
	// w.Write([]byte(r.Header.Get("X-Forwarded-For")))
	w.Write([]byte("Hello"))
}
