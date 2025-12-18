package main
import (
	"fmt"
	"net/http"
)
func main() {
    http.HandleFunc("/hello", hellohandler)
	http.ListenAndServe(":8080", nil)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Glitch !")
}
