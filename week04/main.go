package main
import (
	"fmt"
	"net/http"
	"time"
)
// Week 04: ここに課題のコードを記述してください
func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe(":8080", nil)
}

func info(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst).Format("2006年01月02日 15:04")

	ua := r.Header.Get("User-Agent")
	fmt.Fprintf(w, "現在の日付と時刻は%sで、利用しているブラウザは「%s」ですね！", now, ua)

}
