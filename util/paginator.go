package util

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
)

// Paginator 分页类
type Paginator struct {
	Request     *http.Request
	PerPageNums int
	MaxPages    int

	nums      int64
	pageRange []int
	pageNums  int
	page      int
}

// PageNums 分页总页数
func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}
	pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
	if p.MaxPages > 0 {
		pageNums = math.Min(pageNums, float64(p.MaxPages))
	}
	p.pageNums = int(pageNums)
	return p.pageNums
}

// Nums 总个数
func (p *Paginator) Nums() int64 {
	return p.nums
}

// Page 当前页
func (p *Paginator) Page() int {
	if p.page != 0 {
		return p.page
	}
	if p.Request.Form == nil {
		p.Request.ParseForm()
	}
	p.page, _ = strconv.Atoi(p.Request.Form.Get("p"))
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

// Pages 页数
func (p *Paginator) Pages() []int {
	if p.pageRange == nil && p.nums > 0 {
		var pages []int
		pageNums := p.PageNums()
		page := p.Page()
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
	link, _ := url.ParseRequestURI(p.Request.RequestURI)
	values := link.Query()
	if page == 1 {
		values.Del("p")
	} else {
		values.Set("p", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	return link.String()
}

// PageLinkPrev 上一页
func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page() - 1)
	}
	return
}

// PageLinkNext 下一页
func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page() + 1)
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
	return p.Page() > 1
}

// HasNext 是否存在下一页
func (p *Paginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

// IsActive 是否为当前页
func (p *Paginator) IsActive(page int) bool {
	return p.Page() == page
}

// Offset 偏移量
func (p *Paginator) Offset() int {
	return (p.Page() - 1) * p.PerPageNums
}

// HasPages 存在分页
func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

// End 最后元素
func (p *Paginator) End() int {
	if p.Page() < p.PageNums() {
		return p.Page() * p.PerPageNums
	}
	return int(p.Nums())
}

// NewPaginator 创建分页对象
func NewPaginator(r *http.Request, nums int64) *Paginator {
	p := Paginator{}
	p.Request = r
	if p.PerPageNums, _ = strconv.Atoi(r.FormValue("perPage")); p.PerPageNums <= 0 {
		p.PerPageNums = 12
	}
	p.nums = nums
	return &p
}
