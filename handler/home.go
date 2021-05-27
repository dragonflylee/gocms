package handler

import (
	"math/rand"
	"net/http"
	"time"

	"gocms/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// Dashboard 仪表盘
func Dashboard(c *gin.Context) {
	switch c.Query("action") {
	case "user":
		from, _ := time.Parse(dateFormate, c.Query("from"))
		to, _ := time.Parse(dateFormate, c.Query("to"))
		data := make([]struct {
			Y     string          `json:"y"`
			Item1 decimal.Decimal `json:"item1"`
			Item2 decimal.Decimal `json:"item2"`
		}, int(to.Sub(from).Hours()/24))
		for i := 0; i < len(data); i++ {
			data[i].Y = from.AddDate(0, 0, i).Format(dateFormate)
			data[i].Item1 = decimal.New(rand.Int63n(8000), -2)
			data[i].Item2 = decimal.New(rand.Int63n(4000), -2)
		}
		c.JSON(http.StatusOK, errors.OK(data))

	default:
		c.Set("Data", time.Now())
		c.HTML(http.StatusOK, "index.html", c.Keys)
	}
}
