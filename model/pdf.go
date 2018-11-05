package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"sdbackend/domain"
	"sort"
	"time"
)

var conf Config

type PDFInstallRuns struct {
	ID                    int64 `gorm:"primary_key;auto_increment"`
	Date                  string
	InstallStart          int64
	InstallEnd            int64
	UninstallStart        int64
	UninstallEnd          int64
	NewUserUninstallEnd   int64
	UninstallRate         float64
	NewUserUninstallRate  int64
	LoadDoc               int64
	LoadDocDistinctDevice int64
	MFShow                int64     `gorm:"Column:mf_show"`
	MFShow7               int64     `gorm:"Column:mf_show_7"`
	MFShow30              int64     `gorm:"Column:mf_show_30"`
	MFShowOld             int64     `gorm:"Column:mf_show_old"`
	ServerRun             int64     `gorm:"Column:server_run"`
	ServerRun7            int64     `gorm:"Column:server_run_7"`
	ServerRun30           int64     `gorm:"Column:server_run_30"`
	ServerRunOld          int64     `gorm:"Column:server_run_old"`
	Crash                 int64     `gorm:"Column:crash"`
	CrashRate             int64     `gorm:"Column:crash_rate"`
	CreatedAt             time.Time `gorm:"Column:create_time;type(datetime)" json:"-"`
	UpdatedAt             time.Time `gorm:"Column:update_time;type(datetime)" json:"-"`
}

func (*PDFInstallRuns) TableName() string {
	return "install_runs"
}

type Retentions struct {
	ID               int64 `gorm:"primary_key;auto_increment"`
	Date             string
	Round            int64
	Install          int64
	UnInstall        int64
	NewUserUninstall int64
	ServerRun        int64
	MFShow           int64     `gorm:"Column,mf_show"`
	MFShowRate       int64     `gorm:"Column,mf_show_rate"`
	ServerRunRate    int64     `gorm:"Column,server_run_rate"`
	CreatedAt        time.Time `gorm:"Column,create_time;type(datetime)" json:"-"`
	UpdatedAt        time.Time `gorm:"Column,update_time;type(datetime)" json:"-"`
}

type RetentionView struct {
	Date             string
	Install          int64
	UnInstall        int64
	NewUserUninstall int64
	RMFShow          int64
	RMFShow3         int64
	RMFShow7         int64
	RMFShow30        int64
	RServerRun       int64
	RServerRun3      int64
	RServerRun7      int64
	RServerRun30     int64
}

type MFShowVersion struct {
	Version string  `json:"version"`
	Rate    float64 `json:"rate"`
}

type MFShowVersions []MFShowVersion

func (v MFShowVersions) Len() int {
	return len(v)
}

func (v MFShowVersions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v MFShowVersions) Less(i, j int) bool {
	return v[i].Rate < v[j].Rate
}

func (*Retentions) TableName() string {
	return "retentions"
}

type Feedback struct {
	LogTime  time.Time `bson:"log_time"`
	Contact  string    `bson:"contact"`
	ClientIP string    `bson:"client_ip"`
	OS       int       `bson:"os"`
	Version  string    `bson:"version"`
	Feedback string    `bson:"feedback"`
}

type UninstallOpts struct {
	LogTime  time.Time `bson:"log_time"`
	Contact  string    `bson:"contact"`
	ClientIP string    `bson:"client_ip"`
	OS       int       `bson:"os"`
	Version  string    `bson:"version"`
	Content  string    `bson:"content"`
	Feedback string    `bson:"feedback"`
}

type UninstallResult struct {
	Result string
	Count  int
	Rate   int64
}

type BundleInstall struct {
	Date        string
	Show        int
	ClickClose  int
	ClickOK     int
	ClickCancel int
	NoData      int
	NetError    int
	DownPkg     int
	InstallPkg  int
}

type MiniNewsStats struct {
	Date                string `gorm:"unique_index"`
	ForbidMiniNews      int
	SpeedupRun          int
	LoaderDownloadStart int
	LoaderDownloadEnd   int
	LoaderLoadStart     int
	LoaderLoadEnd       int
	CloudDisable        int
}

type CrashVersionRate struct {
	Version string
	Rate    string
}

type CrashInfo struct {
	LogTime  time.Time `bson:"log_time"`
	ClientIP string    `bson:"client_ip"`
	OS       int       `bson:"os"`
	Version  string    `bson:"version"`
}

func GetPDFInstallRuns(limit, offset int) ([]PDFInstallRuns, error) {
	var ins []PDFInstallRuns
	if err := db.New().Where("date != ?", "total").Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func GetPDFInstallRunsByDate(date string) (*PDFInstallRuns, error) {
	var ins PDFInstallRuns
	if err := db.New().Where("date = ?", date).Order("date desc").First(&ins).Error; err != nil {
		return nil, err
	}
	return &ins, nil
}

func TotalPDFInstallRuns() (int64, error) {
	var total int64
	if err := db.New().Model(new(PDFInstallRuns)).Where("date != ?", "total").Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func GetPDFRentions(limit, offset int) ([]RetentionView, error) {
	var (
		dates      []string
		ins        []Retentions
		retentions []RetentionView
	)
	if err := db.New().Model(new(Retentions)).Select("distinct(date) as date").Order("date desc").Limit(limit).Offset(offset).Pluck("date", &dates).Error; err != nil {
		return nil, err
	}
	if err := db.New().Where("date in (?)", dates).Order("date desc").Find(&ins).Error; err != nil {
		return nil, err
	}

	for i := range ins {
		if len(retentions) == 0 || retentions[len(retentions)-1].Date != ins[i].Date {
			retentions = append(retentions, RetentionView{
				Date:             ins[i].Date,
				Install:          ins[i].Install,
				UnInstall:        ins[i].UnInstall,
				NewUserUninstall: ins[i].NewUserUninstall,
			})
		}
		index := len(retentions) - 1
		if ins[i].Round == 1 {
			retentions[index].RMFShow = ins[i].MFShowRate
			retentions[index].RServerRun = ins[i].ServerRunRate
		}
		if ins[i].Round == 3 {
			retentions[index].RMFShow3 = ins[i].MFShowRate
			retentions[index].RServerRun3 = ins[i].ServerRunRate
		}
		if ins[i].Round == 7 {
			retentions[index].RMFShow7 = ins[i].MFShowRate
			retentions[index].RServerRun7 = ins[i].ServerRunRate
		}
		if ins[i].Round == 30 {
			retentions[index].RMFShow30 = ins[i].MFShowRate
			retentions[index].RServerRun30 = ins[i].ServerRunRate
		}

	}

	return retentions, nil
}

func GetAvgPDFRetentions() (*RetentionView, error) {
	var (
		retentions []Retentions
		result     RetentionView
	)
	if err := db.New().Model(new(Retentions)).
		Select("round, SUM(install) as install,  SUM(mf_show) as mf_show, SUM(server_run) as server_run").
		Group("round").Find(&retentions).Error; err != nil {
		return nil, err
	}
	for i := range retentions {
		if retentions[i].Round == 1 {
			result.RMFShow = retentions[i].MFShow * 10000 / retentions[i].Install
			result.RServerRun = retentions[i].ServerRun * 10000 / retentions[i].Install
		}

		if retentions[i].Round == 3 {
			result.RMFShow3 = retentions[i].MFShow * 10000 / retentions[i].Install
			result.RServerRun3 = retentions[i].ServerRun * 10000 / retentions[i].Install
		}

		if retentions[i].Round == 7 {
			result.RMFShow7 = retentions[i].MFShow * 10000 / retentions[i].Install
			result.RServerRun7 = retentions[i].ServerRun * 10000 / retentions[i].Install
		}

		if retentions[i].Round == 30 {
			result.RMFShow30 = retentions[i].MFShow * 10000 / retentions[i].Install
			result.RServerRun30 = retentions[i].ServerRun * 10000 / retentions[i].Install
		}
	}
	return &result, nil
}

func TotalRetentions() (int64, error) {
	var total int64
	if err := db.New().Model(new(Retentions)).Select("count(distinct(date))").Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func GetMFShowVersions() (MFShowVersions, error) {
	conn := redisPool.NewConn()
	key := "mf_show_versions"
	index := time.Now().Format("2006-01-02")
	v, err := conn.HGet(key, index).Result()
	if err != nil {
		return nil, err
	}
	var versions MFShowVersions
	var results []interface{}
	if err := json.Unmarshal([]byte(v), &results); err != nil {
		return nil, err
	}
	for i := range results {
		versions = append(versions, MFShowVersion{
			Version: results[i].([]interface{})[0].(string),
			Rate:    results[i].([]interface{})[1].(float64),
		})
	}
	sort.Sort(sort.Reverse(versions))
	return versions, nil
}

func GetFeedbacks(limit, offset int) ([]Feedback, error) {
	conn := mgo.Clone()
	defer conn.Close()
	col := conn.DB(mgoDBName).C("feedbacks")
	var feedbacks []Feedback
	if err := col.Find(nil).Limit(limit).Skip(offset).Sort("-log_time").All(&feedbacks); err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func GetTotalFeedbacks() (int, error) {
	conn := mgo.Clone()
	defer conn.Close()
	var (
		total int
		err   error
	)
	col := conn.DB(mgoDBName).C("feedbacks")
	if total, err = col.Count(); err != nil {
		return 0, err
	}
	return total, nil
}

func GetUninstallOpts(limit, offset int) ([]UninstallOpts, error) {
	conn := mgo.Clone()
	defer conn.Close()
	col := conn.DB(mgoDBName).C("uninstall_opts")
	var feedbacks []UninstallOpts
	if err := col.Find(nil).Limit(limit).Skip(offset).Sort("-log_time").All(&feedbacks); err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func GetTotalUninstallOpts() (int, error) {
	conn := mgo.Clone()
	defer conn.Close()
	var (
		total int
		err   error
	)
	col := conn.DB(mgoDBName).C("uninstall_opts")
	if total, err = col.Count(); err != nil {
		return 0, err
	}
	return total, nil
}

func GetUninstallResults() ([]UninstallResult, error) {
	conn := redisPool.NewConn()
	key := "uninstall_results"
	v, err := conn.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	var (
		results []UninstallResult
		source  map[string]int
	)
	if err := json.Unmarshal(v, &source); err != nil {
		return nil, err
	}
	var total = 0
	for _, v := range source {
		total += v
	}
	if total == 0 {
		return nil, errors.New("Total invalid")
	}
	for k, v := range source {
		results = append(results, UninstallResult{
			Result: k,
			Count:  v,
			Rate:   int64(v) * int64(10000) / int64(total),
		})
	}
	return results, nil
}

func GetTotalBundleInstalls() (int64, error) {
	var total int64
	if err := db.New().Model(new(BundleInstall)).Where("date != ?", "total").Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func GetBundleInstalls(limit, offset int) ([]BundleInstall, error) {
	var ins []BundleInstall
	if err := db.New().Where("date != ?", "total").Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func GetTotalMiniNewsStats() (int64, error) {
	var total int64
	if err := db.New().Model(new(MiniNewsStats)).Where("date != ?", "total").Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func GetMiniNewsStats(limit, offset int) ([]domain.MiniNewsStats, error) {
	var ins []domain.MiniNewsStats
	if err := db.New().Where("date != ?", "total").Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func GetCrashsTotal(start, end *time.Time) (int64, error) {
	conn := mgo.Clone()
	defer conn.Close()
	col := conn.DB(mgoDBName).C("crashs")
	query := bson.M{}
	if start != nil {
		query["$gte"] = start
	}
	if end != nil {
		query["$lt"] = end
	}
	c, err := col.Find(query).Count()
	return int64(c), err
}

func GetCrashsByDay(limit, offset int, start, end *time.Time) ([]CrashInfo, error) {
	conn := mgo.Clone()
	defer conn.Close()
	col := conn.DB(mgoDBName).C("crashs")
	var crashs []CrashInfo
	query := bson.M{}
	if start != nil {
		query["$gte"] = start
	}
	if end != nil {
		query["$lt"] = end
	}
	if err := col.Find(query).Limit(limit).Skip(offset).Sort("-log_time").All(&crashs); err != nil {
		return nil, err
	}

	return crashs, nil
}

func GetCrashVersioRate() ([]CrashVersionRate, error) {
	conn := mgo.Clone()
	defer conn.Close()
	col := conn.DB(mgoDBName).C("crashs")
	pipeLine := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id": "$version",
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
	}
	var results []struct {
		Version string `bson:"_id"`
		Count   int    `bson:"count"`
	}
	pipe := col.Pipe(pipeLine)
	if err := pipe.All(&results); err != nil {
		return nil, err
	}

	var (
		total int
		res   []CrashVersionRate
	)
	for _, ins := range results {
		total += ins.Count
	}
	for _, ins := range results {
		res = append(res, CrashVersionRate{
			Version: ins.Version,
			Rate:    fmt.Sprintf("%.02f", (float32(ins.Count*100) / float32(total))),
		})
	}
	return res, nil
}
