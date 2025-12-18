package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
	"encoding/json"
	"time"
	"strconv"
)

const logFile = "public/logs.json" // データの保存先 --- (*1)

// Log 掲示板に保存するデータを構造体で定義 --- (*2)
type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Body  string `json:"body"`
	CTime int64  `json:"ctime"`
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/bbs", showHandler)
	http.HandleFunc("/write", writeHandler)
	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

// 書き込みログを画面に表示する --- (*6)
func showHandler(w http.ResponseWriter, r *http.Request) {
	// ログを読み出してHTMLを生成 --- (*7)
	htmlLog := ""
	logs := loadLogs() // データを読み出す

	// エラー表示
	errMsg := r.URL.Query().Get("error")
	if errMsg != ""{
		htmlLog += "<p style='color:red;'>" +
			html.EscapeString(errMsg) + "</p>"
	}
	// 投稿数を表示させる
	htmlLog += "<p>投稿数：" + strconv.Itoa(len(logs)) + "</p>"

	//表示を2行に、さまざまな情報と本文で分ける
	for _, i := range logs {
		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%sさん</span> --- %s<br>%s</p>",
			i.ID,
			html.EscapeString(i.Name),
			time.Unix(i.CTime, 0).Format("2006/1/2 15:04"),
			html.EscapeString(i.Body),
		)
	}
	// HTML全体を出力 --- (*8)
	htmlBody := "<html><head><style>" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style></head><body><h1>BBS</h1>" +
		getForm() + htmlLog + "</body></html>"

	_, _ = w.Write([]byte(htmlBody))
}

// フォームから送信された内容を書き込み --- (*9)
func writeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // フォームを解析 --- (*10)
	if err != nil{
		http.Error(w, "フォームの解析に失敗しました", http.StatusBadRequest)
		return
	}

	var log Log
	log.Name = r.Form["name"][0]
	log.Body = r.Form["body"][0]
	//名前なかったとき
	if log.Name == "" {
		log.Name = "名無し"
	}
	//本文の空白投稿を防ぐ
	if log.Body == ""{
		http.Redirect(w, r, "/bbs?error=本文を入力してください", 302)
		return
	}
	//字数制限 200字まで
	if len([]rune(log.Body)) > 150{
		http.Redirect(w, r, "/bbs?error=本文は150文字以内で入力してください", 302)
		return
	}

	logs := loadLogs() // 既存のデータを読み出し --- (*11)
	log.ID = len(logs) + 1
	log.CTime = time.Now().Unix()
	logs = append(logs, log)         // 追記 --- (*12)
	saveLogs(logs)                   // 保存

	http.Redirect(w, r, "/bbs", 302) // リダイレクト --- (*13)
}

// 書き込みフォームを返す --- (*14)
func getForm() string {
	return "<div><form action='/write' method='get'>" +
	"名前: <input type='text' name='name'><br>" +
	"本文: <input type='text' name='body' style='width:30em;' maxlength='150'><br>" +
	"<input type='submit' value='書込'>" +
	"</form></div><hr>"
}

// ファイルからログファイルの読み込み --- (*15)
func loadLogs() []Log {
	// ファイルを開く
	text, err := os.ReadFile(logFile)
	if err != nil {
		return make([]Log, 0)
	}
	// JSONをパース --- (*16)
	var logs []Log
	// JSONが壊れていたら空配列を返す
	if err := json.Unmarshal(text, &logs); err != nil {
		return make([]Log, 0)
	}
	return logs
}

// ログファイルの書き込み --- (*17)
func saveLogs(logs []Log) {
	// JSONにエンコード
	bytes, _ := json.Marshal(logs)
	// ファイルへ書き込む
	_ =  os.WriteFile(logFile, bytes, 0644)
}
