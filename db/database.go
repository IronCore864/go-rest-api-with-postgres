package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"hello/model"
	"os"
)

type DBConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DBName string
}

func LoadConfiguration() DBConf {
	var config DBConf
	filename := "config/dbconf.json"
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

var dbconf = LoadConfiguration()

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	dbconf.Host, dbconf.Port, dbconf.User, dbconf.Pwd, dbconf.DBName)

func GetQuote() *model.Quote {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var q model.Quote
	err = db.QueryRow("SELECT quote, category FROM quote ORDER BY RANDOM() LIMIT 1").Scan(&q.Quote, &q.Category)
	if err != nil {
		fmt.Println(err)
	}

	return &q
}

func InsertQuote(quote *model.Quote) int {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO quote VALUES (DEFAULT, $1, $2) RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, quote.Quote, quote.Category).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	return id
}
