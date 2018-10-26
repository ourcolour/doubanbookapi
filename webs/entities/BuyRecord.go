package entities

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type BuyRecord struct {
	Isbn     string `json:"isbn" bson:"isbn"`
	Store    string `json:"store" bson:"store"`
	Quantity int    `json:"quantity" bson:"quantity"`

	BuyDate    time.Time `json:"buy_date" bson:"buy_date"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}

func NewBuyRecord(isbn string, store string, quantity int, buyDate time.Time, createTime time.Time, updateTime time.Time) *BuyRecord {
	return &BuyRecord{isbn, store, quantity, buyDate, createTime, updateTime}
}
func NewBuyRecordByJson(jsonStr string) *BuyRecord {
	var (
		result BuyRecord
		err    error
	)

	if 0 == strings.Compare("", jsonStr) {
		return nil
	}

	if nil != err {
		log.Printf("%s\n", err.Error())
	} else {
		bytes := []byte(jsonStr)
		err = json.Unmarshal(bytes, &result)
		if nil != err {
			log.Printf("%s\n", err.Error())
		}
	}

	return &result
}
