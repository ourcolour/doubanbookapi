package entities

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type RatingInfo struct {
	NumRaters int     `json:"numRaters" bson:"numRaters"`
	Max       float64 `json:"max" bson:"max"`
	Average   float64 `json:"average" bson:"average"`
	Min       float64 `json:"min" bson:"min"`
}

type ImageInfo struct {
	Small  string `json:"small" bson:"small"`
	Medium string `json:"medium" bson:"medium"`
	Large  string `json:"large" bson:"large"`
}

type TagInfo struct {
	Count int    `json:"count" bson:"count"`
	Name  string `json:"name" bson:"name"`
	Title string `json:"title" bson:"title"`
}

type BookInfo struct {
	Rating *RatingInfo `json:"rating" bson:"rating"`

	SubTitle string   `json:"subtitle" bson:"subtitle"`
	Authors  []string `json:"author" bson:"author"`
	PubDate  string   `json:"pubdate" bson:"pubdate"`

	Tags        []*TagInfo `json:"tags" bson:"tags"`
	OriginTitle string     `json:"origin_title" bson:"origin_title"`
	Image       string     `json:"image" bson:"image"`
	Binding     string     `json:"binding" bson:"binding"`
	Translator  string     `json:"translator" bson:"translator"`
	Catalog     string     `json:"catalog" bson:"catalog"`
	Pages       int        `json:"pages" bson:"pages"`
	Images      *ImageInfo `json:"images" bson:"images"`
	Alt         string     `json:"alt" bson:"alt"`
	Id          string     `json:"id" bson:"id"`
	Publisher   string     `json:"publisher" bson:"publisher"`
	Isbn10      string     `json:"isbn10" bson:"isbn10"`
	Isbn13      string     `json:"isbn13" bson:"isbn13"`
	Title       string     `json:"title" bson:"title"`
	Url         string     `json:"url" bson:"url"`
	AltTitle    string     `json:"alt_title" bson:"alt_title"`
	AuthorIntro string     `json:"author_intro" bson:"author_intro"`
	Summary     string     `json:"summary" bson:"summary"`
	Price       float64    `json:"price" bson:"price"`

	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}

func NewBookInfoByJson(jsonStr string) *BookInfo {
	var (
		result BookInfo
		err    error
	)

	if 0 == strings.Compare("", strings.TrimSpace(jsonStr)) {
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
