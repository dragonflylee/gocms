package handler

import (
	"fmt"
	"gocms/model"
	"gocms/util"
	"net/http"
)

func P2WInstallRuns(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalP2WInstallRuns(); err == nil && nums > 0 {
		fmt.Println(nums)
		p := util.NewPaginator(r, nums)
		if p2ws, err := model.GetP2WInstallRuns(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = p2ws
		}
		data["page"] = p
	}
	rLayout(w, r, "p2w.tpl", data)
}
