package person

type Person struct {
	ID      int    `json:"id"`                // encoding/jsonパッケージのタグ付け機能 出力するオブジェクト名を指定
	Name    string `json:"name"`              // "name":として出力される
	Email   string `json:"-"`                 // JSONに格納しない
	Age     int    `json:"age"`               // "age":として出力される
	Address string `json:"address,omitempty"` // "値が空なら無視
	memo    string
}
