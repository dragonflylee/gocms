package model

import (
	"math"
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

// Paginator 分页类
type Paginator struct {
	PerPageNums int `form:"perPage"`
	Page        int `form:"p"`

	rawLink   string
	nums      int64
	pageRange []int
	pageNums  int
}

// PageNums 分页总页数
func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}
	p.pageNums = int(math.Ceil(float64(p.nums) / float64(p.PerPageNums)))
	return p.pageNums
}

// Nums 总个数
func (p *Paginator) Nums() int64 {
	return p.nums
}

// Pages 页数
func (p *Paginator) Pages() []int {
	if p.pageRange == nil && p.nums > 0 {
		var pages []int
		pageNums := p.PageNums()
		page := p.Page
		switch {
		case page >= pageNums-4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page+4+1))))
			for i := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRange
}

// PageLink 分页链接
func (p *Paginator) PageLink(page int) string {
	query, _ := url.ParseQuery(p.rawLink)
	if page == 1 {
		query.Del("p")
	} else {
		query.Set("p", strconv.Itoa(page))
	}
	return "?" + query.Encode()
}

// PageLinkPrev 上一页
func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page - 1)
	}
	return
}

// PageLinkNext 下一页
func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page + 1)
	}
	return
}

// PageLinkFirst 第一页
func (p *Paginator) PageLinkFirst() (link string) {
	return p.PageLink(1)
}

// PageLinkLast 最后一页
func (p *Paginator) PageLinkLast() (link string) {
	return p.PageLink(p.PageNums())
}

// HasPrev 是否存在上一页
func (p *Paginator) HasPrev() bool {
	return p.Page > 1
}

// HasNext 是否存在下一页
func (p *Paginator) HasNext() bool {
	return p.Page < p.PageNums()
}

// IsActive 是否为当前页
func (p *Paginator) IsActive(page int) bool {
	return p.Page == page
}

// HasPages 存在分页
func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

// End 最后元素
func (p *Paginator) End() int {
	if p.Page < p.PageNums() {
		return p.Page * p.PerPageNums
	}
	return int(p.Nums())
}

// Scope implment of model.Scope
func (p *Paginator) Scope(x *gorm.DB) *gorm.DB {
	return x.Limit(p.PerPageNums).Offset((p.Page - 1) * p.PerPageNums)
}

// NewPaginator 创建分页对象
func NewPaginator(r *url.URL, nums int64) *Paginator {
	p := &Paginator{nums: nums, rawLink: r.RawQuery}
	if p.Page, _ = strconv.Atoi(r.Query().Get("p")); p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPageNums, _ = strconv.Atoi(r.Query().Get("perPage")); p.PerPageNums <= 0 {
		p.PerPageNums = 12
	}
	if p.Page > p.PageNums() {
		p.Page = p.PageNums()
	}
	return p
}
