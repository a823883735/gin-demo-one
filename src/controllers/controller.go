package controllers

import (
	"math"
	"strconv"
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
	IsLastPage      bool        `json:"isLastPage"`      //是否为最后一页
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

func NewPage(num, size string) (page *Page) {
	page = new(Page)
	var err error
	page.PageNum, err = strconv.Atoi(num)
	if err != nil {
		page.PageNum = 1
	}
	page.PageSize, err = strconv.Atoi(size)
	if err != nil {
		page.PageSize = 10
	}
	if page.PageNum == 0 {
		page.PageNum = 1
	}
	if page.PageSize == 0 {
		page.PageSize = 10
	}
	page.FirstPage = 1
	return
}

func (page *Page) GetListSplitPage(list interface{}, total, count int) {
	page.Total = total
	page.List = list
	if page.Total == 0 {
		page.IsFirstPage = true
		page.IsLastPage = true
		page.PageNum = 1
		page.LastPage = 1
		page.Pages = 1
		page.NextPage = 1
		page.PrePage = 1
	} else {
		page.Pages = int(math.Ceil(float64(page.Total) / float64(page.PageSize)))
		page.LastPage = page.Pages
		page.StartRow = (page.PageNum-1)*page.PageSize + 1
		page.List = list
		if count == 0 {
			page.EndRow = page.StartRow
		} else {
			page.EndRow = page.StartRow + count - 1
		}
		if page.EndRow < page.Total {
			page.NextPage = page.PageNum + 1
			page.IsLastPage = false
			page.HasNextPage = true
		} else {
			page.NextPage = page.PageNum
			page.IsLastPage = true
			page.HasNextPage = false
		}
		if page.PageNum == 1 {
			page.IsFirstPage = true
			page.PrePage = 1
		} else {
			page.IsFirstPage = false
			page.HasPreviousPage = true
			page.PrePage = page.PageNum - 1
		}
	}
}
