package main

import (
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/doubanbookapi/webs/portal"
	"log"
)

func main() {
	webLauncher, err := portal.NewWebLauncherWithHostAndPort(configs.SERVICE_ARRD, configs.SERVICE_PORT)
	if nil != err {
		log.Panic(err)
	}
	webLauncher.Run()
}
