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
    http.HandleFunc("/bmi", bmihandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}



func bmihandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	//小数対応
	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	//BMIの計算
	heightM := height / 100.0
	bmiValue := weight / (heightM * heightM)

	fmt.Fprintf(w, "あなたのBMIは、%.2f です\n", bmiValue)

	// 判定
	var hantei string
	if bmiValue < 18.5 {
		hantei = "低体重（やせ型）"
	} else if bmiValue < 25 {
		hantei = "普通体重"
	} else {
		hantei = "肥満気味"
	}

	fmt.Fprintf(w, "判定：%s です", hantei)
}
