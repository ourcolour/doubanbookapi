package errs

import "errors"

var (
	ERR_NONE            = errors.New("没有错误")
	ERR_UNKNOWN         = errors.New("未知错误")
	ERR_NOT_IMPLEMENTED = errors.New("暂未实现")
	ERR_OK              = errors.New("执行成功")

	ERR_NOT_FOUND          = errors.New("内容未找到")
	ERR_INVALID_PARAMETERS = errors.New("无效的参数")
	ERR_INVALID_VALUES     = errors.New("数据库操作失败")
	ERR_DATABASE_OPERATION = errors.New("数据库操作失败")
)