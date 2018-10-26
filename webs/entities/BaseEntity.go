package entities

import (
	"iamcc.cn/doubanbookapi/utils"
)

func (this *BookInfo) ToJsonString() (string, error) {
	return utils.ToJsonString(this)
}
