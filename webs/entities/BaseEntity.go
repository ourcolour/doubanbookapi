package entities

import (
	"iamcc.cn/doubanbookapi/utils"
)

func (this *Book) ToJsonString() (string, error) {
	return utils.ToJsonString(this)
}
