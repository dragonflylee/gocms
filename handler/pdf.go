package handler

import (
	"gocms/model"
	"gocms/util"
	"net/http"
)

func PDFInstallRuns(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalPDFInstallRuns(); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if installRuns, err := model.GetPDFInstallRuns(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = installRuns
		}
		if grandTotal, err := model.GetPDFInstallRunsByDate("total"); err == nil {
			data["total"] = grandTotal
		}
		data["page"] = p
	}
	rLayout(w, r, "pdf_install_runs.tpl", data)
}

func PDFRentions(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalRetentions(); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if retentions, err := model.GetPDFRentions(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = retentions
		}
		data["page"] = p
	}
	if result, err := model.GetAvgPDFRetentions(); err == nil {
		data["avg"] = *result
	}
	rLayout(w, r, "pdf_retentions.tpl", data)
}

func MFShowVersions(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	result, err := model.GetMFShowVersions()
	if err == nil {
		data["list"] = result
	}
	rLayout(w, r, "pdf_mfshow_versions.tpl", data)

}

func Feedbacks(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalFeedbacks(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if feedbacks, err := model.GetFeedbacks(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = feedbacks
		}
		data["page"] = p
	}
	rLayout(w, r, "pdf_feedbacks.tpl", data)

}

func UninstallOpts(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalUninstallOpts(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if uninstallOpts, err := model.GetUninstallOpts(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = uninstallOpts
		}
		data["page"] = p
	}
	if results, err := model.GetUninstallResults(); err == nil {
		data["results"] = results
	}
	rLayout(w, r, "pdf_uninstall_opts.tpl", data)
}

func BundleInstall(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalBundleInstalls(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if list, err := model.GetBundleInstalls(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = list
		}
		data["page"] = p
	}
	rLayout(w, r, "bundle_installation.tpl", data)
}

func MiniNewsStats(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalMiniNewsStats(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if list, err := model.GetMiniNewsStats(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = list
		}
		data["page"] = p
	}
	rLayout(w, r, "mininews_stats.tpl", data)
}
