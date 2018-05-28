package handler

import (
	"encoding/json"
	"gocms/model"
	"gocms/util"
	"net/http"
	"strings"
)

func P2WInstallRuns(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalP2WInstallRuns(); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if p2ws, err := model.GetP2WInstallRuns(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = p2ws
		}
		data["page"] = p
	}
	rLayout(w, r, "p2w_all.tpl", data)
}

func P2WInstallRunsGroupByQD(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := make(map[string]interface{})
	qd := strings.TrimSpace(r.Form.Get("qd"))
	if qd == "all" {
		qd = ""
	}
	if nums, err := model.TotalP2WInstallRunsGroupByQD(qd); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if p2ws, err := model.GetP2WInstallRunsGroupByQD(qd, p.PerPageNums, p.Offset()); err == nil {
			data["list"] = p2ws
		}
		data["page"] = p
	}
	rLayout(w, r, "p2w_qd.tpl", data)
}

func P2WQDList(w http.ResponseWriter, r *http.Request) {
	var (
		qds []string
		err error
	)
	qds, err = model.AllQDs()
	if err != nil {
		http.NotFound(w, r)
		return
	}
	data := make(map[string]interface{})
	selects := []select2{
		select2{"all", "all"},
	}
	for i := range qds {
		selects = append(selects, select2{ID: qds[i], Name: qds[i]})
	}
	data["results"] = selects
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
