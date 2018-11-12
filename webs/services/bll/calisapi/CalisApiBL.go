package calisapi

import (
	"fmt"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/utils"
	"log"
	"strings"
)

const (
	ISBN_STYLE_31441 = iota
	ISBN_STYLE_31531

	CALIS_DOMAIN    = "opac.calis.edu.cn"
	API_GATEWAY_URL = "http://" + CALIS_DOMAIN + "/"
)

var DEFAULT_REQUEST_HEADER = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Encoding": "gzip, deflate, br",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Cache-Control":   "max-age=0",
	"Connection":      "keep-alive",
	"Cookie":          "bid=Ozs1l7okmjs; gr_user_id=d2af9175-e992-40d5-a7e8-6a84a90bdb2b; _pk_id.100001.3ac3=bf60c1dab497fa73.1536047844.23.1539575968.1539570750.; _vwo_uuid_v2=DE8AA116596CB1C2A49172C9CE7175CE8|99b0ea88210915872510d833be19d4f7; __yadk_uid=3DUeAWFZVDYza6eZVsqEKtQqnXqMWhYZ; __utma=30149280.1480626694.1536047847.1539570720.1539575960.28; __utmz=30149280.1539251716.21.6.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utma=81379588.2076642786.1536047848.1539570720.1539575960.22; __utmz=81379588.1539071650.15.4.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; viewed=\"30237215_30274766_26434970_30288491_27620808_27191009_25752592\"; douban-fav-remind=1; ue=\"ourcolour@qq.com\"; push_noty_num=0; push_doumail_num=0; __utmv=30149280.5237; _pk_ref.100001.3ac3=%5B%22%22%2C%22%22%2C1539575958%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _ga=GA1.2.1480626694.1536047847; douban-profile-remind=1; ck=zbDo; __utmc=30149280; __utmc=81379588; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=9b860930-619e-4762-80a9-5cf3eb24c47e; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_9b860930-619e-4762-80a9-5cf3eb24c47e=true; gr_cs1_9b860930-619e-4762-80a9-5cf3eb24c47e=user_id%3A0; ap_v=0,6.0; _pk_ses.100001.3ac3=*; __utmb=30149280.2.10.1539575960; __utmt_douban=1; __utmb=81379588.2.10.1539575960; __utmt=1",
	"Host":            CALIS_DOMAIN,
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:62.0) Gecko/20100101 Firefox/62.0",
	//"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0",

	"Referer": API_GATEWAY_URL,
}

func ParseIsbn(oldIsbn string, style int) (string, error) {
	var (
		format string = "%s-%s-%s-%s-%s"

		result string
		err    error
	)

	// Clear src
	result = strings.TrimSpace(strings.Replace(oldIsbn, "-", "", -1))

	// Reformat
	if ISBN_STYLE_31531 == style {
		result = fmt.Sprintf(format, result[0:3], result[3:4], result[4:9], result[9:12], result[12:])
	} else {
		result = fmt.Sprintf(format, result[0:3], result[3:4], result[4:8], result[8:12], result[12:])
	}

	return result, err
}

func getDetail(oid string, isbn string) ([]string, error) {
	//http://opac.calis.edu.cn/opac/showDetails.do?recIndex=1&recTotal=1&fullTextType=1&dbselect=all&oid=597d9701198909fb24b38a757efffdfe
	// 定义
	var (
		link     string              = fmt.Sprintf("%sopac/showDetails.do?recIndex=1&recTotal=1&fullTextType=1&dbselect=all&oid=%s", API_GATEWAY_URL, oid)
		postData map[string][]string = map[string][]string{
			"actionType":      {""},
			"CLC2":            nil,
			"codelanguage":    nil,
			"codetype":        nil,
			"creator2":        nil,
			"datepublication": nil,
			"dbselect": {
				0: "all",
				1: "all",
			},
			"fromTree":       {"false"},
			"groupkey":       nil,
			"langBase":       nil,
			"maximumRecords": {"50"},
			"operation":      {"searchRetrieve"},
			"pageno":         {"1"},
			"pageno2":        {"1"},
			"pagingType":     {"1"},
			"query":          {"(bath.isbn=\"" + isbn + "*\")"},
			//"query":          {"(cql.anywhere=\"*克苏鲁神话合集*\")"},
			"queryType":  {"0"},
			"searchlang": nil,
			"selInvalid": {"请选择记录"},
			"series2":    nil,
			"shw_cql":    {fmt.Sprintf("ISBN检索+=+*%s*", isbn)},
			"sortagainname": {
				0: "title",
				1: "title",
			},
			"sortkey":     {"title"},
			"startRecord": {"1"},
			"unititle2":   nil,
			"version":     {"1.1"},
		}

		//pattern  string = "javascript:dosearch\\('bath.localClassification','([^']+)'\\)"
		pattern string = "<a href=\"javascript:dosearch\\('bath.localClassification','(?:[^']+)'\\)\">([^<]+)</a>"

		result []string
		err    error
	)

	// 参数
	if "" == strings.TrimSpace(oid) {
		return result, errs.ERR_INVALID_PARAMETERS
	}
	if "" == strings.TrimSpace(isbn) || (13+4 != len(isbn) && 10+4 != len(isbn)) {
		return result, errs.ERR_INVALID_PARAMETERS
	}

	// 抓取
	itf, err := utils.HttpPost(link, nil, postData, DEFAULT_REQUEST_HEADER, func(data []byte) (result interface{}, err error) {
		if nil == data {
			return result, errs.ERR_INVALID_PARAMETERS
		}
		return string(data), nil
	})

	// 页面内容
	var html string = ""
	if nil == err {
		html = itf.(string)
	}

	lst, err := utils.FindMatchedGroupByGroupIndex(html, pattern, 1)
	if nil != err {
		return result, err
	}

	result = []string{}
	for item := lst.Front(); item != nil; item = item.Next() {
		curCip := strings.Replace(strings.Replace(item.Value.(string), "*", "", -1), " ", "", -1)
		result = append(result, curCip)
	}

	return result, err
}

func removeDuplicateCip(inCips []string) ([]string, error) {
	// 定义
	var (
		cipLength int
		result    []string
		err       error
	)

	// 参数
	if nil == inCips || len(inCips) < 1 {
		return nil, errs.ERR_INVALID_PARAMETERS
	}

	// 逐个检查并添加
	cipLength = len(inCips)
	if cipLength < 2 {
		result = inCips
	} else {
		for i := 0; i < cipLength; i++ {
			if (i > 0 && inCips[i-1] == inCips[i]) || len(inCips[i]) == 0 {
				continue
			}
			result = append(result, inCips[i])
		}
	}

	return result, err
}

func GetCipByIsbn(isbn string) ([]string, error) {
	// 定义
	var (
		link     string              = fmt.Sprintf("%sopac/doSimpleQuery.do", API_GATEWAY_URL)
		postData map[string][]string = map[string][]string{
			"actionType":     {"doSimpleQuery"},
			"condition":      {isbn},
			"conInvalid":     {"检索条件不能为空"},
			"dbselect":       {"all"},
			"indexkey":       {"bath.isbn|frt"},
			"langBase":       {"default"},
			"maximumRecords": {"50"},
			"operation":      {"searchRetrieve"},
			"pageno":         {"1"},
			"pagingType":     {"0"},
			"query":          {"(bath.isbn=\"" + isbn + "*\")"},
			"sortkey":        {"title"},
			"startRecord":    {"1"},
			"version":        {"1.1"},
		}

		pattern string = "javascript:doShowDetails\\(([^)]+),([^)]+),([^)]+),([^)]+),([^)]+),([^)]+)\\)"
		oid     string
		cips    []string

		result []string
		err    error
	)

	// 参数
	if "" == strings.TrimSpace(isbn) || (13+4 != len(isbn) && 10+4 != len(isbn)) {
		return result, errs.ERR_INVALID_PARAMETERS
	}

	// 请求头
	var requestHeader map[string]string = DEFAULT_REQUEST_HEADER
	requestHeader["Content-Type"] = "application/x-www-form-urlencoded"

	// 抓取
	itf, err := utils.HttpPost(link, nil, postData, requestHeader, func(data []byte) (result interface{}, err error) {
		if nil == data {
			return result, errs.ERR_INVALID_PARAMETERS
		}
		return string(data), nil
	})

	// 页面内容
	var html string = ""
	if nil == err {
		html = itf.(string)
	}

	lst, err := utils.FindMatchedGroupByGroupIndex(html, pattern, 5)
	if nil != err {
		return result, err
	} else if lst.Len() < 1 {
		return result, err
	}

	for item := lst.Front(); item != nil; item = item.Next() {
		oid = strings.Replace(item.Value.(string), "'", "", -1)

		log.Printf("OID: %s\n", oid)

		// 根据 oid 找到页面后抓取 cip
		cips, err = getDetail(oid, isbn)
		if nil != err {
			continue
		}

		for _, cip := range cips {
			result = append(result, cip)
		}
	}

	// 去重
	result, err = removeDuplicateCip(result)

	return result, err
}
