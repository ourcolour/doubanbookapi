package es

import (
	"encoding/json"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	"log"
	"strings"
	"time"
)

type ES_Rating struct {
	NumRaters int     `json:"numRaters" bson:"numRaters"`
	Max       float64 `json:"max" bson:"max"`
	Average   float64 `json:"average" bson:"average"`
	Min       float64 `json:"min" bson:"min"`
}

type ES_Image struct {
	Small  string `json:"small" bson:"small"`
	Medium string `json:"medium" bson:"medium"`
	Large  string `json:"large" bson:"large"`
}

type ES_Tag struct {
	Count int    `json:"count" bson:"count"`
	Name  string `json:"name" bson:"name"`
	Title string `json:"title" bson:"title"`
}

type ES_Book struct {
	//ObjectId bson.ObjectId `json:"_id" bson:"_id"`

	Rating      *ES_Rating `json:"rating" bson:"rating"`
	SubTitle    string     `json:"subtitle" bson:"subtitle"`
	Authors     []string   `json:"author" bson:"author"`
	PubDate     string     `json:"pubdate" bson:"pubdate"`
	Cips        []string   `json:"cips" bson:"cips"`
	Tags        []*ES_Tag  `json:"tags" bson:"tags"`
	OriginTitle string     `json:"origin_title" bson:"origin_title"`
	Image       string     `json:"image" bson:"image"`
	Binding     string     `json:"binding" bson:"binding"`
	Translator  string     `json:"translator" bson:"translator"`
	Catalog     string     `json:"catalog" bson:"catalog"`
	Pages       int        `json:"pages" bson:"pages"`
	Images      *ES_Image  `json:"images" bson:"images"`
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

func NewESBookByJson(jsonStr string) *ES_Book {
	var (
		result ES_Book
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

func NewESBookByInterface(itf interface{}) *ES_Book {
	jsonStr := utils.MustToJsonString(itf)
	return NewESBookByJson(jsonStr)
}

func NewESBookByBook(book *entities.Book) *ES_Book {
	if nil == book {
		return nil
	}

	var result *ES_Book = new(ES_Book)
	if nil != book.Rating {
		result.Rating = &ES_Rating{
			NumRaters: book.Rating.NumRaters,
			Average:   book.Rating.Average,
			Max:       book.Rating.Max,
			Min:       book.Rating.Min,
		}
	}
	result.SubTitle = book.SubTitle
	result.Authors = book.Authors
	result.PubDate = book.PubDate
	result.Cips = book.Cips
	result.Tags = []*ES_Tag{}
	for _, tag := range book.Tags {
		if nil == tag {
			continue
		}
		result.Tags = append(result.Tags, &ES_Tag{
			Count: tag.Count,
			Title: tag.Title,
			Name:  tag.Name,
		})
	}
	result.OriginTitle = book.OriginTitle
	result.Image = book.Image
	result.Binding = book.Binding
	result.Translator = book.Translator
	result.Catalog = book.Catalog
	result.Pages = book.Pages
	if nil != book.Images {
		result.Images = &ES_Image{
			Small:  book.Images.Small,
			Medium: book.Images.Medium,
			Large:  book.Images.Large,
		}
	}
	result.Alt = book.Alt
	result.Id = book.Id
	result.Publisher = book.Publisher
	result.Isbn10 = book.Isbn10
	result.Isbn13 = book.Isbn13
	result.Title = book.Title
	result.Url = book.Url
	result.AltTitle = book.AltTitle
	result.AuthorIntro = book.AuthorIntro
	result.Summary = book.Summary
	result.Price = book.Price
	result.CreateTime = book.CreateTime
	result.UpdateTime = book.UpdateTime

	return result
}
