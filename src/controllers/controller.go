package controllers

import (
	"github.com/go-xorm/xorm"
	"reflect"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Page struct {
	FirstPage       int         `json:"firstpage"`       //第一页页码
	IsFirstPage     bool        `json:"isFirstPage"`     //是否为第一页
	LastPage        int         `json:"lastPage"`        //最后一页页码
	IsLastPage      int         `json:"isLastPage"`      //是否为最后一页
	HasNextPage     bool        `json:"hasNextPage"`     //是否有下一页
	HasPreviousPage bool        `json:"hasPreviousPage"` //是否有上一页
	Total           int         `json:"total"`           //数据总条数
	Pages           int         `json:"pages"`           //页数
	PageNum         int         `json:"pageNum"`         //当前页码
	PageSize        int         `json:"pageSize"`        //每页条数
	StartRow        int         `json:"startRow"`        //当前页第一条行号
	EndRow          int         `json:"endRow"`          //当前页最后一条行号
	NextPage        int         `json:"nextPage"`        //下一页页码
	PrePage         int         `json:"prePage"`         //上一页页码
	List            interface{} `json:"list"`            //列表数据
}

func GetListSplitPage(session xorm.Session, obj []interface{}, pageNum, pageSize int) (page Page) {
	if err := session.Find(obj); err == nil {
		page.Total = len(obj)
		if pageNum == 0 {
			page.PageNum = 1
		}
		if pageSize == 0 {
			page.PageSize = 10
		}
		if page.Total == 0 {

		} else {
			page.Pages = int(page.Total / page.PageSize)
		}
	}
	return
}
