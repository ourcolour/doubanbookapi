package main

import (
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/doubanbookapi/webs/portal"
)

func main() {
	webLauncher, _ := portal.NewWebLauncherWithHostAndPort(configs.SERVICE_ARRD, configs.SERVICE_PORT)
	webLauncher.Run()
}
