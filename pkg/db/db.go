package db

import (
	cr "../Credential"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Config struct {
	DB        *sql.DB
	Ticker    *time.Ticker
	MaxTranId int
}

func New() Config {
	return Config{
		DB:     ConnectPGDB(),
		Ticker: time.NewTicker(time.Minute * 5),
	}
}

//Connect to DB
func ConnectPGDB() *sql.DB {
	user, err := cr.GetCredential(fmt.Sprintf("/etc/secret-volume/User"))
	if err != nil {
		log.Println("Error read secret file User", err)
	}

	pass, err := cr.GetCredential(fmt.Sprintf("/etc/secret-volume/Pass"))
	if err != nil {
		log.Println("Error read secret file Pass", err)
	}

	DB, err := sql.Open("postgres", fmt.Sprintf("host=external-orderpg.dp.wb.ru user=%s password=%s dbname=UserInfo", user, pass))
	if err != nil {
		log.Fatal("Error open Postgres DB:", err.Error())
		return nil
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error ping Postgres DB:", err.Error())
	}

	DB.SetConnMaxLifetime(time.Minute * 5)
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)

	log.Println("Connect to PG")

	return DB
}

func (c *Config) PingDb() {
	log.Println("Ошибка подключения к БД")
	ticker := time.NewTicker(time.Minute * 10)
	for range ticker.C {
		c.DB = ConnectPGDB()
		if err := c.DB.Ping(); err == nil {
			log.Println("Подключение выполнено!")
			ticker.Stop()
			return
		}
	}
}
