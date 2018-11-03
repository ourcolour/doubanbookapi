package _default

import (
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	mongoDAL "iamcc.cn/doubanbookapi/frameworks/services/impl/dal/mongodb"
)

func VerifyDB() error {
	var (
		err error = nil
	)

	// Check DB
	existed, err := mongoDAL.ExistsDabatase(configs.MGO_DATABASE)
	if nil != err {
		return err
	} else if !existed {
		return errs.ERR_DATABASE_NOT_INITIALIZED
	}

	// Check Collection
	colArray := []string{
		"sl_book_new",
	}
	for _, curName := range colArray {
		existed, err = mongoDAL.ExistsCollection(curName)
		if nil != err {
			return err
		} else if !existed {
			return errs.ERR_DATABASE_NOT_INITIALIZED
		}
	}

	return nil
}
