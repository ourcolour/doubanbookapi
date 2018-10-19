package doubanapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"iamcc.cn/doubanbookapi/utils"
	"iamcc.cn/doubanbookapi/webs/entities"
	"regexp"
	"strings"
)

const DOUBAN_BOOK_DOMAIN string = "book.douban.com"
const API_GATEWAY_URL = "https://api.douban.com/v2/"

type ErrorMessage struct {
	Code     string
	HttpCode int
	Message  string
	Comment  string
}

var ERROR_CODE_ARRAY = map[string]*ErrorMessage{
	"999999": &ErrorMessage{Code: "999999", HttpCode: 400, Message: "unknown_error", Comment: "未知错误"},
	"6000":   &ErrorMessage{Code: "6000", HttpCode: 404, Message: "book_not_found", Comment: "图书不存在"},
	"6002":   &ErrorMessage{Code: "6002", HttpCode: 403, Message: "unauthorized_error", Comment: "没有修改权限"},
	"6004":   &ErrorMessage{Code: "6004", HttpCode: 400, Message: "review_content_short(should more than 150)", Comment: "书评内容过短（需多于150字）"},
	"6006":   &ErrorMessage{Code: "6006", HttpCode: 404, Message: "review_not_found", Comment: "书评不存在"},
	"6007":   &ErrorMessage{Code: "6007", HttpCode: 403, Message: "not_book_request", Comment: "不是豆瓣读书相关请求"},
	"6008":   &ErrorMessage{Code: "6008", HttpCode: 404, Message: "people_not_found", Comment: "用户不存在"},
	"6009":   &ErrorMessage{Code: "6009", HttpCode: 400, Message: "function_error", Comment: "服务器调用异常"},
	"6010":   &ErrorMessage{Code: "6010", HttpCode: 400, Message: "comment_too_long(should less than 350)", Comment: "短评字数过长（需少于350字）"},
	"6011":   &ErrorMessage{Code: "6011", HttpCode: 409, Message: "collection_exist(try PUT if you want to update)", Comment: "该图书已被收藏（如需更新请用PUT方法而不是POST）"},
	"6012":   &ErrorMessage{Code: "6012", HttpCode: 400, Message: "invalid_page_number(should be digit less than 1000000)", Comment: "非法页码（页码需要是小于1000000的数字）"},
	"6013":   &ErrorMessage{Code: "6013", HttpCode: 400, Message: "chapter_too_long(should less than 100)", Comment: "章节名过长（需小于100字）"},
}

var DEFAULT_REQUEST_HEADER = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Encoding": "gzip, deflate, br",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Cache-Control":   "max-age=0",
	"Connection":      "keep-alive",
	"Cookie":          "bid=Ozs1l7okmjs; gr_user_id=d2af9175-e992-40d5-a7e8-6a84a90bdb2b; _pk_id.100001.3ac3=bf60c1dab497fa73.1536047844.23.1539575968.1539570750.; _vwo_uuid_v2=DE8AA116596CB1C2A49172C9CE7175CE8|99b0ea88210915872510d833be19d4f7; __yadk_uid=3DUeAWFZVDYza6eZVsqEKtQqnXqMWhYZ; __utma=30149280.1480626694.1536047847.1539570720.1539575960.28; __utmz=30149280.1539251716.21.6.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utma=81379588.2076642786.1536047848.1539570720.1539575960.22; __utmz=81379588.1539071650.15.4.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; viewed=\"30237215_30274766_26434970_30288491_27620808_27191009_25752592\"; douban-fav-remind=1; ue=\"ourcolour@qq.com\"; push_noty_num=0; push_doumail_num=0; __utmv=30149280.5237; _pk_ref.100001.3ac3=%5B%22%22%2C%22%22%2C1539575958%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _ga=GA1.2.1480626694.1536047847; douban-profile-remind=1; ck=zbDo; __utmc=30149280; __utmc=81379588; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=9b860930-619e-4762-80a9-5cf3eb24c47e; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_9b860930-619e-4762-80a9-5cf3eb24c47e=true; gr_cs1_9b860930-619e-4762-80a9-5cf3eb24c47e=user_id%3A0; ap_v=0,6.0; _pk_ses.100001.3ac3=*; __utmb=30149280.2.10.1539575960; __utmt_douban=1; __utmb=81379588.2.10.1539575960; __utmt=1",
	"Host":            DOUBAN_BOOK_DOMAIN,
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:62.0) Gecko/20100101 Firefox/62.0",
	//"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0",

	"Referer": "https://" + DOUBAN_BOOK_DOMAIN + "/",
}

func GetBookByIsbn(isbn string) (*entities.BookInfo, error) {
	var (
		result *entities.BookInfo
		err    error
	)

	apiUrl := fmt.Sprintf(API_GATEWAY_URL+"book/isbn/%s", isbn)
	headers := DEFAULT_REQUEST_HEADER

	itf, err := utils.HttpGet(apiUrl, nil, headers, func(data []byte) (result interface{}, err error) {
		if nil == data {
			err = errors.New("Invalid parameters.")
			return result, err
		}

		result = &entities.BookInfo{}
		data = fixBookInfoJson(data)
		json.Unmarshal(data, result)

		return result, err
	})

	if nil == err {
		result = itf.(*entities.BookInfo)
	}

	return result, err
}

func fixBookInfoJson(inBytes []byte) []byte {
	var result string = string(inBytes)

	// "average":"8.7"
	// "price":"49.00"

	keyArray := []string{
		"average", "price",
	}

	for _, key := range keyArray {
		reg := regexp.MustCompile(strings.Replace(`("${key}":)("(\d+(?:.\d+)?)")`, "${key}", key, -1))
		result = reg.ReplaceAllString(result, "$1$3")
	}

	return []byte(result)
}
