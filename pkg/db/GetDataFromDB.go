package db

import (
	"../models"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type DataFromDB struct {
	tranid     int
	dataJson   string
	dateInsert string
}

func GetDataFromDB(ctx context.Context) {
	conf := New()
	conf.MaxTranId = GetMaxTranId(conf.DB)
	ChecksUpdate(ctx, conf)
}

func ChecksUpdate(ctx context.Context, conf Config) {
	for {
		select {
		case <-conf.Ticker.C:
			data := GetOrdersFromDB(conf)
			for i := range data {
				models.ChanOrderPay <- fmt.Sprintf("Новый заказ (номер заказа = %v) за %v от %v", data[i].tranid, data[i].dateInsert, data[i].dataJson)
				if data[i].tranid > conf.MaxTranId {
					conf.MaxTranId = data[i].tranid
				}
			}
		case <-ctx.Done():
			log.Println("Exit db loop")
			return
		}

	}
}

func GetOrdersFromDB(conf Config) []DataFromDB {
	var data []DataFromDB

	rows, err := conf.DB.Query(fmt.Sprintf("select tran_id, data, data_insert from orders.orders where tran_id > %v", conf.MaxTranId))
	if err != nil {
		log.Println(err)
		conf.PingDb()
	}
	for rows.Next() {
		var dataS DataFromDB
		err := rows.Scan(&dataS.tranid, &dataS.dataJson, &dataS.dateInsert)
		if err != nil {
			log.Println(err)
		}
		data = append(data, dataS)
	}
	return data
}

func GetMaxTranId(db *sql.DB) int {
	var MaxTranId int
	row := db.QueryRow("select max(tran_id) from orders.orders")
	err := row.Scan(&MaxTranId)
	if err != nil {
		log.Println("Error scan max tran-id", err)
	}
	return MaxTranId
}
