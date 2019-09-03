package test

import (
	"net/url"
	"testing"
	"time"

	"github.com/dragonflylee/gocms/util"

	"github.com/shopspring/decimal"
)

func Test_UtilForm(t *testing.T) {
	var m struct {
		ID    []int64         `form:"id"`
		Name  []string        `form:"name"`
		Vol   decimal.Decimal `form:"vol"`
		Date  *time.Time      `form:"date"`
		Count *int            `form:"count"`
	}
	m.Count = new(int)
	n := time.Now().UTC()
	f := url.Values{
		"date":  {n.Format(time.RFC3339)},
		"name":  {"a1", "b2", "c3", "c4"},
		"id":    {"2", "7", "12"},
		"vol":   {"3.1415926"},
		"count": {"15"},
	}
	if err := util.ParseForm(f, &m); err != nil {
		t.Fatal(err)
	}
	if m.Date == nil {
		t.FailNow()
	}
	t.Logf("%v %v %v", m.ID, m.Name, m.Vol)
	t.Logf("%d", *m.Count)

	if !m.Date.Equal(n) {
		t.Logf("time %v", m.Date.Sub(n))
		t.Fail()
	}
}
