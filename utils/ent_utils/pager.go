package ent_utils

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/cjrd/allocate"
	"reflect"
)

//分页表单
type PageOptions struct {
	Page int `json:"page" form:"page"; query:"page"`
	Size int `json:"size" form:"size"; query:"size"`
}

//排序表单
type PageSort struct {
	SortColumn string `json:"sort_column" form:"sort_column"`
	SortOrder  string `json:"sort_order" form:"sort_order"`
}

//排序方法
func PageSortDB(client interface{}, pageSort PageSort) {
	t := reflect.ValueOf(client)
	if pageSort.SortOrder == "ascend" && pageSort.SortColumn != "" {
		params := []reflect.Value{reflect.ValueOf(ent.Asc(pageSort.SortColumn))}
		t.MethodByName("Order").Call(params)
	}
	if pageSort.SortOrder == "descend" && pageSort.SortColumn != "" {
		params := []reflect.Value{reflect.ValueOf(ent.Desc(pageSort.SortColumn))}
		t.MethodByName("Order").Call(params)
	}
}

//
func GetDefaultPager(num int, defaultNum int) int {
	if num <= 0 {
		return defaultNum
	}
	return num
}


//填充默认分页
func PipePagerFn(dataSource interface{}) {
	allocate.Zero(dataSource)
	t := reflect.ValueOf(dataSource).Elem()
	page := t.FieldByName("Page")
	size := t.FieldByName("Size")
	if page.Int() < 1 {
		page.SetInt(1)
	}
	if size.Int() < 1 {
		size.SetInt(10)
	}
}

//分页方法
func PipeLimitFn(db interface{}, dataSource interface{}) {
	data := reflect.ValueOf(dataSource).Elem()
	page := data.FieldByName("Page").Int()
	size := data.FieldByName("Size").Int()
	offset := (page - 1) * size
	refDb := reflect.ValueOf(db)
	refDb.MethodByName("Limit").Call([]reflect.Value{reflect.ValueOf(int(size))})
	refDb.MethodByName("Offset").Call([]reflect.Value{reflect.ValueOf(int(offset))})
}