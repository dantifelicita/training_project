package talk_training

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Dependency to connect postgres
	// "log"
)

const (
	host     = "192.168.100.126"
	port     = 5432
	user     = "techacademy"
	password = "123qwe!@#QWE"
	dbname   = "tokopedia-talk"
)

var (
	db *sql.DB
)

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbnew, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = dbnew.Ping()
	if err != nil {
		panic(err)
	}

	db = dbnew
}

func GetTalks(product_id int64) []Talks {
	initDB()

	ListTalks := []Talks{}

	query := fmt.Sprintf("SELECT talk_id, product_id, message, create_time "+
		"FROM ws_talk where product_id=%d", product_id)
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		Talk := Talks{}
		rows.Scan(&Talk.ID, &Talk.ProductID, &Talk.Message, &Talk.CreateTime)
		ListTalks = append(ListTalks, Talk)
	}

	return ListTalks
}

func PostTalk(m Messages) {
	initDB()

	_, err := db.Query("INSERT INTO ws_talk (product_id, shop_id, message, user_id, create_by) VALUES "+
		"($1, $2, $3, $4, $5)", m.ProductID, m.ShopID, m.Message, m.UserID, m.UserID)

	// log.Printf(query)

	// _, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}
