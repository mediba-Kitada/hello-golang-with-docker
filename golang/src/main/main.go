package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"persona"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Golang with Docker!!")
}

// テンプレートのコンパイル
var t = template.Must(template.ParseFiles("assets/index.html")) // Must()関数でエラー時にパニックを発生させる(一度コンパイルを確認できたテンプレートであれば毎回エラー処理をする必要性は低い)

func PersonaHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // 処理の最後に必ずBodyを閉じる

	if r.Method == "GET" {
		// クエリパラメータを取得
		id, err := strconv.Atoi(r.URL.Query().Get("id")) // r.URL.Query().Get("id")の値は文字列(string型)なのでstrconv.Atoi()関数で数値に変換
		if err != nil {
			log.Println(err)
			id = 1
		}

		// idをキーにDBからペルソナ情報を取得
		persona, err := persona.GetPersona(id)
		if err != nil {
			log.Fatal("Could not get persona")
		}

		// レスポンスにエンコーディングしたHTMLを書き込む
		t.Execute(w, persona)
	}

	// GETメソッド以外は許可しない
	if r.Method != "GET" {

		log.Println("Fobidden Request:", r.Method)
		// レスポンスとしてステータスコード403を送信
		w.WriteHeader(http.StatusForbidden)

	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persona", PersonaHandler)
	http.ListenAndServe(":8080", nil)
}
