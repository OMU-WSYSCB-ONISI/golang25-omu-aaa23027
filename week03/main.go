package main
import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Week 03
/*
now同様、webfortuneでアクセスすると答えるWebおみくじ
今の運勢は大吉（中吉、吉、今日）です
*/
	http.HandleFunc("/webfortune", fortunehandler)
	http.ListenAndServe(":8080", nil)
}

func fortunehandler(w http.ResponseWriter, r *http.Request) {
	fortunes := []string{"大吉", "中吉", "吉", "凶"}

	rand.Seed(time.Now().UnixNano())

	fmt.Fprintf(w, "今の運勢は %s です！", fortunes[rand.Intn(len(fortunes))])

}
