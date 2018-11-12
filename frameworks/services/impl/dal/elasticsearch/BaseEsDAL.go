package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/useelk/constants/errs"
	"log"
	"strings"
)

func getEsUrl() string {
	return fmt.Sprintf("http://%s:%d", configs.ES_HOST, configs.ES_PORT)
}

func getClient() (*elastic.Client, error) {
	elastic.SetSniff(configs.ES_ENABLE_SNIFFER)
	url := getEsUrl()
	return elastic.NewClient(elastic.SetURL(url))
}
func mustGetClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(getEsUrl()))
	if nil != err {
		log.Panicln(err)
	}
	return client
}

func Ping() (interface{}, int, error) {
	var (
		result interface{}
		delay  int
		err    error
	)

	client := mustGetClient()
	result, delay, err = client.Ping(getEsUrl()).Do(context.Background())

	return result, delay, err
}

func DeleteIndex(indexName string) (bool, error) {
	var (
		result bool
		err    error
	)

	client, err := getClient()
	if nil != err {
		return result, err
	}

	resp, err := client.DeleteIndex(indexName).Do(context.Background())
	result = resp.Acknowledged

	return result, err
}

func DeleteAll(indexName string, typeName string) (int64, error) {
	var (
		result int64 = 0
		err    error
	)

	// 参数
	if 0 == strings.Compare("", indexName) {
		return result, errs.ERR_INVALID_PARAMETERS
	}
	if 0 == strings.Compare("", typeName) {
		return result, errs.ERR_INVALID_PARAMETERS
	}

	client, err := getClient()
	if nil != err {
		return result, err
	}

	resp, err := client.DeleteByQuery(indexName).Type().QueryString("*").Do(context.Background())
	result = resp.Deleted

	return result, err
}

func Add(indexName string, typeName string, doc interface{}) (bool, error) {
	count, err := BatchAdd(indexName, typeName, []interface{}{doc})
	return 0 < count, err
}

func BatchAdd(indexName string, typeName string, docList []interface{}) (int64, error) {
	var (
		result int64 = 0
		err    error
	)

	// 参数
	if 0 == strings.Compare("", indexName) {
		return result, errs.ERR_INVALID_PARAMETERS
	}
	if 0 == strings.Compare("", typeName) {
		return result, errs.ERR_INVALID_PARAMETERS
	}
	if nil == docList || len(docList) < 1 {
		return result, errs.ERR_INVALID_PARAMETERS
	}

	client, err := getClient()
	if nil != err {
		return result, err
	}

	for _, doc := range docList {
		res, err := client.Index().Index(indexName).Type(typeName).BodyJson(doc).Do(context.Background())
		if nil != err {
			log.Printf("Idx:%s Type:%s got error when inserting.", err.Error())
		} else {
			log.Printf("Idx:%s Type:%s Id:%s Inserted.\n", res.Index, res.Type, res.Result+"|"+res.Id)
			result += 1
		}
	}

	return result, err
}

func Search(indexName string, typeName string, query elastic.Query, pageSize int, pageNo int) ([]*json.RawMessage, int64, error) {
	var (
		result           []*json.RawMessage = make([]*json.RawMessage, 0)
		totalRecordCount int64              = 0
		err              error
	)

	client, err := getClient()
	if nil != err {
		return result, totalRecordCount, err
	}

	//query := elastic.NewQueryStringQuery(queryString)
	searchResult, err := client.Search(indexName).
		Type(typeName).
		Query(query).
		From((pageNo - 1) * pageSize).
		Size(pageSize).
		Do(context.Background())

	if nil != err {
		return result, totalRecordCount, err
	}

	totalRecordCount = searchResult.TotalHits()
	if 0 < totalRecordCount {
		for _, hit := range searchResult.Hits.Hits {
			result = append(result, hit.Source)
		}
	}

	return result, totalRecordCount, err
}

func RawSearch(indexName string, typeName string, queryString string, pageSize int, pageNo int) (*elastic.SearchResult, error) {
	var (
		result *elastic.SearchResult
		err    error
	)

	client, err := getClient()
	if nil != err {
		return result, err
	}

	query := elastic.NewQueryStringQuery(queryString)
	result, err = client.Search(indexName).
		Type(typeName).
		Query(query).
		From((pageNo - 1) * pageSize).
		Size(pageSize).
		Do(context.Background())
	if nil != err {
		return result, err
	}

	return result, err
}
