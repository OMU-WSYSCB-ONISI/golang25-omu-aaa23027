package main
import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
    http.HandleFunc("/ave", avehandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}


func avehandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
		return
	}

	// カンマで区切って配列にする
	tokuten := strings.Split(r.FormValue("dd"), ",")

	// 合計と人数
	sum := 0
	count := 0

	// 分布（0-9点、10-19点...90-100点）
	bunpu := make([]int, 11)

	// 1つずつ処理
	for _, t := range tokuten {
		// 数字に変換
		score, _ := strconv.Atoi(strings.TrimSpace(t))

		// 合計に足す
		sum += score
		count++

		// 分布に記録
		if score == 100 {
			bunpu[10]++
		} else {
			bunpu[score/10]++
		}
	}

	// 平均を計算
	ave := float64(sum) / float64(count)

	// 結果を表示
	fmt.Fprintf(w, "データ数: %d 個\n", count)
	fmt.Fprintf(w, "合計: %d 点\n", sum)
	fmt.Fprintf(w, "平均: %.2f 点\n\n", ave)

	fmt.Fprintln(w, "【得点分布】")
	fmt.Fprintln(w, "0-9点:", bunpu[0], "人")
	fmt.Fprintln(w, "10-19点:", bunpu[1], "人")
	fmt.Fprintln(w, "20-29点:", bunpu[2], "人")
	fmt.Fprintln(w, "30-39点:", bunpu[3], "人")
	fmt.Fprintln(w, "40-49点:", bunpu[4], "人")
	fmt.Fprintln(w, "50-59点:", bunpu[5], "人")
	fmt.Fprintln(w, "60-69点:", bunpu[6], "人")
	fmt.Fprintln(w, "70-79点:", bunpu[7], "人")
	fmt.Fprintln(w, "80-89点:", bunpu[8], "人")
	fmt.Fprintln(w, "90-99点:", bunpu[9], "人")
	fmt.Fprintln(w, "100点:", bunpu[10], "人")
}
