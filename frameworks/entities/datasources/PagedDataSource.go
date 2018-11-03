package datasources

type PagedDataSource struct {
	Pagination *Pagination `json:"pagination"`
	DataList   interface{} `json:"dataList"`
}

func NewPagedDataSource(pagination *Pagination, dataList interface{}) *PagedDataSource {
	return &PagedDataSource{
		Pagination: pagination,
		DataList:   dataList,
	}
}
