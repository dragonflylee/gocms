package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"fmt"
	"gocms/model"
	"gocms/util"
)

func QDList(w http.ResponseWriter, r *http.Request) {
	var (
		user *model.Admin
		ok   bool
		qds  []string
		err  error
	)
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
		return
	} else if cookie, exist := session.Values["user"]; !exist {
		http.NotFound(w, r)
		return
	} else if user, ok = cookie.(*model.Admin); !ok {
		http.NotFound(w, r)
		return
	}

	if user.Group.ID == 1 || user.Group.Name == "data_admin" {
		qds, err = model.AllQDs()
		fmt.Println(qds)
		if err != nil {
			http.NotFound(w, r)
			return
		}
	} else {
		qds, err = model.AdmindQDs(user.ID)
		if err != nil {
			http.NotFound(w, r)
			return
		}
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

func QDSettleDay(w http.ResponseWriter, r *http.Request) {
	var (
		user *model.Admin
		ok   bool
		qds  []string
		err  error
	)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
		return
	} else if cookie, exist := session.Values["user"]; !exist {
		http.NotFound(w, r)
		return
	} else if user, ok = cookie.(*model.Admin); !ok {
		http.NotFound(w, r)
		return
	}

	tpl := "settle_day_all.tpl"
	if user.Group.ID == 1 || user.Group.Name == "data_admin" {
		qds, err = model.AllQDs()
		if err != nil {
			http.NotFound(w, r)
			return
		}
	} else {
		qds, err = model.AdmindQDs(user.ID)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		tpl = "settle_day_qd.tpl"
	}

	data := make(map[string]interface{})
	if qds != nil {
		if qd := strings.TrimSpace(r.Form.Get("qd")); len(qd) > 0 && qd != "all" {
			check := false
			for i := range qds {
				if (qds)[i] == qd {
					qds = []string{qd}
					check = true
					break
				}
			}
			if !check {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}

		var viewList []model.QDInstallRuns
		if nums, err := model.TotalInstallRunsByQD(qds); err == nil && nums > 0 {
			p := util.NewPaginator(r, nums)
			if qdStats, err := model.InstallRunsByQD(qds, p.PerPageNums, p.Offset()); err == nil {
				for i := range qdStats {
					coefficient := qdStats[i].Coefficient
					if user.Group.ID == 1 || user.Group.Name == "data_admin" {
						coefficient = 100
					}
					qdStats[i].Total = (qdStats[i].InstallEnd * qdStats[i].Coefficient / 100) * qdStats[i].Price
					qdStats[i].InstallStart = qdStats[i].InstallStart * coefficient / 100
					qdStats[i].UninstallStart = qdStats[i].UninstallStart * coefficient / 100
					qdStats[i].UninstallEnd = qdStats[i].UninstallEnd * coefficient / 100
					qdStats[i].MFShow = qdStats[i].MFShow * coefficient / 100
					qdStats[i].ServerRun = qdStats[i].ServerRun * coefficient / 100
					qdStats[i].InstallEnd = qdStats[i].InstallEnd * coefficient / 100
					viewList = append(viewList, qdStats[i])
				}
				data["list"] = viewList
			}
			data["page"] = p
		}
	}
	rLayout(w, r, tpl, data)
}

func QDSettleMonth(w http.ResponseWriter, r *http.Request) {
	var (
		user *model.Admin
		ok   bool
		qds  []string
		err  error
	)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
		return
	} else if cookie, exist := session.Values["user"]; !exist {
		http.NotFound(w, r)
		return
	} else if user, ok = cookie.(*model.Admin); !ok {
		http.NotFound(w, r)
		return
	}

	tpl := "settle_month_all.tpl"
	if user.Group.ID == 1 || user.Group.Name == "data_admin" {
		qds, err = model.AllQDs()
		if err != nil {
			http.NotFound(w, r)
			return
		}
	} else {
		qds, err = model.AdmindQDs(user.ID)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		tpl = "settle_month_qd.tpl"
	}

	data := make(map[string]interface{})
	fmt.Println(qds)
	if qds != nil {
		if qd := strings.TrimSpace(r.Form.Get("qd")); len(qd) > 0 && qd != "all" {
			check := false
			for i := range qds {
				if (qds)[i] == qd {
					qds = []string{qd}
					check = true
					break
				}
			}
			if !check {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}

		if nums, err := model.TotalMonthSettleByQD(qds); err == nil && nums > 0 {
			p := util.NewPaginator(r, nums)
			if qdStats, err := model.MonthSettleByQD(qds, p.PerPageNums, p.Offset()); err == nil {
				data["list"] = qdStats
			}
			data["page"] = p
		}
	}
	rLayout(w, r, tpl, data)
}
