package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	. "person"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

// テンプレートのコンパイル
var t = template.Must(template.ParseFiles("assets/index.html")) // Must()関数でエラー時にパニックを発生させる(一度コンパイルを確認できたテンプレートであれば毎回エラー処理をする必要性は低い)

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // 処理の最後に必ずBodyを閉じる

	if r.Method == "GET" {
		// クエリパラメータを取得
		id, err := strconv.Atoi(r.URL.Query().Get("id")) // r.URL.Query().Get("id")の値は文字列(string型)なのでstrconv.Atoi()関数で数値に変換
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("%d.txt", id) // {id].txtのファイル名を特定
		b, err := ioutil.ReadFile(filename)   // {id}.txtの中身を[]byte型の変数bに代入
		if err != nil {
			log.Fatal(err)
		}

		// personを生成
		person := Person{
			ID:   id,
			Name: string(b),
		}

		// レスポンスにエンコーディングしたHTMLを書き込む
		t.Execute(w, person)
	}

	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person Person                  // Person型(構造体)の変数を宣言
		decoder := json.NewDecoder(r.Body) // リクエストボディを用いてデコーダを取得
		err := decoder.Decode(&person)     // JSON形式にデコードし、person変数に代入
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", person.ID) //ファイル名を {id}.txtとする
		file, err := os.Create(filename)             // {id}.txtファイルを作成
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close() // ファイルを必ず閉じる

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name) // io.WriteString()関数でstring型のデータをファイルに書き込む
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)

	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":8080", nil)
}
