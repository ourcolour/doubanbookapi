package main

import (
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/doubanbookapi/webs/portal"
	"log"
)

func main() {
	//keyword := "雍正帝"
	//pageNo := 1
	//pageSize := 10
	//
	//searchService := impl.NewSearchService()
	//dataList, pagination, err := searchService.SearchBook(keyword, pageNo, pageSize)
	//if nil != err {
	//	panic(err)
	//}
	//
	//log.Printf("Current page: %d", pagination.CurrentPage)
	//for _, cur := range dataList {
	//	log.Printf("%-20s %-5.2f %-5d", cur.Title, cur.Rating.Average, cur.Rating.NumRaters)
	//
	//	jsonString, _ := utils.ToJsonString(cur)
	//	log.Println(jsonString)
	//	break
	//}
	//
	//return
	webLauncher, err := portal.NewWebLauncherWithHostAndPort(configs.SERVICE_ARRD, configs.SERVICE_PORT)
	if nil != err {
		log.Panic(err)
	}
	webLauncher.Run()
}
