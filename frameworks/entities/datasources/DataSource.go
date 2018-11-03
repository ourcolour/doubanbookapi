package datasources

type DataSource struct {
	DataList interface{} `json:"dataList"`
}

func NewDataSource(dataList []interface{}) *DataSource {
	var result *DataSource = &DataSource{
		DataList: dataList,
	}
	return result
}

func FromPagedDataSource(pagedDataSource *PagedDataSource) *DataSource {
	if nil == pagedDataSource {
		return nil
	}

	var result *DataSource = &DataSource{
		DataList: pagedDataSource.DataList,
	}
	return result
}
