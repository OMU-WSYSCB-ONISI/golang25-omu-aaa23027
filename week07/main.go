package main
import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	http.Handle("/", http.FileServer(http.Dir("public/")))
    http.HandleFunc("/cal02", cal02handler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}

}

func cal02handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		fmt.Fprintln(w, x/y)
	}
}
