package dal

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"iamcc.cn/doubanbookapi/configs"
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"iamcc.cn/doubanbookapi/utils"
	"log"
	"reflect"
	"strings"
)

func getDialInfo() *mgo.DialInfo {
	var (
		mgoAddrs       []string = strings.Split(strings.Replace(configs.MGO_ADDRESSES, " ", "", -1), ",")
		replicaSetName string   = configs.MGO_REPLICATE_SET_NAME
		database       string   = configs.MGO_DATABASE

		username string = configs.MGO_USERNAME
		password string = configs.MGO_PASSWORD

		dialInfo *mgo.DialInfo
	)

	dialInfo = &mgo.DialInfo{
		Addrs:          mgoAddrs,
		Direct:         len(mgoAddrs) < 2,
		ReplicaSetName: replicaSetName,
		Database:       database,
	}

	if "" != username {
		dialInfo.Username = username
		dialInfo.Password = password
	}

	return dialInfo
}

func connect() (*mgo.Session, error) {
	var (
		dialInfo *mgo.DialInfo = getDialInfo()

		result *mgo.Session
		err    error
	)

	result, err = mgo.DialWithInfo(dialInfo)
	if nil != err {
		log.Fatalf("Failed to connect to mongodb, %s.\n", err.Error())
	} else {
		result.SetMode(mgo.Nearest, true)
	}

	return result, err
}

func FindId(colName string, id interface{}) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return nil, err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	query := col.FindId(id)
	err = query.One(&result)

	// 忽略空记录异常
	if reflect.TypeOf(errs.ERR_NOT_FOUND).Elem() == reflect.TypeOf(err).Elem() {
		err = nil
	}

	return result, err
}

func FindOne(colName string, selector bson.M) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return nil, err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	query := col.Find(selector)
	err = query.One(&result)

	// 忽略空记录异常
	if nil != err && reflect.TypeOf(errs.ERR_NOT_FOUND).Elem() == reflect.TypeOf(err).Elem() {
		err = nil
	}

	return result, err
}

func FindAll(colName string, selector bson.M, typ reflect.Type) ([]interface{}, error) {
	var (
		result []interface{} = make([]interface{}, 0)
		err    error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return result, err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	//log.Println(utils.ToJsonString(selector))
	query := col.Find(selector)

	itr := query.Iter()
	pObj := reflect.New(typ).Interface()
	for itr.Next(pObj) {
		result = append(result, pObj)
		pObj = reflect.New(typ).Interface()
	}

	return result, err
}

func FindList(colName string, selector bson.M, typ reflect.Type, skip int, limit int) ([]interface{}, error) {
	var (
		result []interface{} = make([]interface{}, 0)
		err    error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return nil, err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	log.Println(utils.ToJsonString(selector))
	query := col.Find(selector).Skip(skip).Limit(limit)

	itr := query.Iter()
	pObj := reflect.New(typ).Interface()
	for itr.Next(pObj) {
		result = append(result, pObj)
		pObj = reflect.New(typ).Interface()
	}

	return result, err
}

func Count(colName string, selector bson.M) (int, error) {
	var (
		result int = 0
		err    error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return 0, err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	result, err = col.Find(selector).Count()

	return result, err
}

func MustFindOne(colName string, selector bson.M) interface{} {
	val, err := FindOne(colName, selector)
	if nil != err {
		return nil
	} else {
		return val
	}
}

func MustFindAll(colName string, selector bson.M, typ reflect.Type) []interface{} {
	result, err := FindAll(colName, selector, typ)
	if nil != err {
		log.Panicln(err)
	}
	return result
}

func MustFindList(colName string, selector bson.M, typ reflect.Type, skip int, limit int) []interface{} {
	result, err := FindList(colName, selector, typ, skip, limit)
	if nil != err {
		log.Panicln(err)
	}
	return result
}

// ---

func Insert(colName string, data interface{}) error {
	var (
		err error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	err = col.Insert(data)

	return err
}

func Update(colName string, selector bson.M, data interface{}) error {
	var (
		err error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	err = col.Update(selector, data)

	return err
}

func UpdateId(colName string, id interface{}, data interface{}) error {
	var (
		err error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	err = col.UpdateId(id, data)

	return err
}

func Upsert(colName string, selector bson.M, data interface{}) (*mgo.ChangeInfo, error) {
	var (
		changeInfo *mgo.ChangeInfo
		err        error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return nil, err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	changeInfo, err = col.Upsert(selector, data)

	return changeInfo, err
}

func UpsertId(colName string, id interface{}, data interface{}) error {
	var (
		err error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	err = col.UpdateId(id, data)

	return err
}

func Remove(colName string, selector bson.M) error {
	var (
		err error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	err = col.Remove(selector)

	return err
}

func RemoveAll(colName string, selector bson.M) (*mgo.ChangeInfo, error) {
	var (
		changeInfo *mgo.ChangeInfo
		err        error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return nil, err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	changeInfo, err = col.RemoveAll(selector)

	return changeInfo, err
}

func RemoveId(colName string, id interface{}) error {
	var (
		err error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return err
	}
	defer session.Close()
	col := session.DB(configs.MGO_DATABASE).C(colName)

	err = col.RemoveId(id)

	return err
}

func ExistsDabatase(dbName string) (bool, error) {
	var (
		result bool = false
		err    error
	)

	if "" == dbName {
		err = errs.ERR_INVALID_PARAMETERS
		return result, err
	}

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return result, err
	}
	defer session.Close()

	dbArray, err := session.DatabaseNames()
	if nil != err {
		return result, err
	}

	for _, curDbName := range dbArray {
		if 0 == strings.Compare(curDbName, dbName) {
			result = true
			break
		}
	}

	return result, err
}

func ExistsCollection(colName string) (bool, error) {
	var (
		result bool = false
		err    error
	)

	if "" == colName {
		err = errs.ERR_INVALID_PARAMETERS
		return result, err
	}

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return result, err
	}
	defer session.Close()
	db := session.DB(configs.MGO_DATABASE)

	collectionArray, err := db.CollectionNames()
	if nil != err {
		return result, err
	}

	for _, curColName := range collectionArray {
		if 0 == strings.Compare(curColName, colName) {
			result = true
			break
		}
	}

	return result, err
}

func Ping() error {
	var (
		err error
	)

	session, err := connect()
	if nil != err {
		log.Printf("%s\n", err.Error())
		return err
	}
	defer session.Close()

	err = session.Ping()

	return err
}
