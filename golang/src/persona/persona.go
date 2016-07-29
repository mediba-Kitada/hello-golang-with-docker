package persona

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	sq "gopkg.in/Masterminds/squirrel.v1"
	gorp "gopkg.in/gorp.v1"
)

type Persona struct {
	ID             int64 `db:id` // ペルソナIDだで把握出来ればよいのでその他のフィールドはプライベート
	sex            int64
	age_min        int64
	age_max        int64
	priority       int64
	Name           string `db:name`
	public_started int64
	public_ended   int64
	delete_flag    int64
	create_userid  int64
	update_userid  int64
	created        int64
	updated        int64
}

func GetPersona(personaID int) (persona *Persona, err error) {

	dbmap := initDb()      // DB接続 構造体とのマッピング
	defer dbmap.Db.Close() // 必ずコネクションを閉じる

	// SQL生成
	personas := sq.Select("id", "name").From("persona")
	targetPersona := personas.Where(sq.Eq{"id": personaID}) // ペルソナIDをキーに指定
	sql, args, err := targetPersona.ToSql()                 // プレースホルダ付SQL, プレースホルダに渡す値の配列, 処理結果
	log.Println(sql, args)
	if err != nil {
		log.Fatal(err)
	}

	// SQL発行
	err = dbmap.SelectOne(&persona, sql, args[0])
	if err != nil {
		persona = &Persona{
			ID:   1,
			Name: "標準ペルソナ",
		}
		log.Println(err)
	}

	return persona, nil
}

func initDb() (dbmap *gorp.DbMap) {
	db, err := sql.Open("mysql", "user:password@tcp(mysql:3306)/sample")
	if err != nil {
		log.Fatal("sql.Open() failed")
	}
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "apollo:", log.Lmicroseconds)) // 全てのSQLステートメントをトレース
	return
}
