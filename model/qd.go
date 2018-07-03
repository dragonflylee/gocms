package model

type AdminQD struct {
	ID int64  `gorm:"primary_key"`
	QD string `gorm:"primary_key"`
}

type QDInstallRuns struct {
	ID             int64 `gorm:"primary_key;auto_increment"`
	Date           string
	QD             string
	InstallStart   int64
	InstallEnd     int64
	UninstallStart int64
	UninstallEnd   int64
	MFShow         int64
	ServerRun      int64
	Coefficient    int64
	Price          int64
	Total          int64 `gorm:"-"`
}

type GroupCoefficient struct {
	ID          int64 `gorm:"primary_key;auto_increment"`
	GroupName   string
	Coefficient int
	Price       int64
	Start       string
}

func AdmindQDs(id int64) ([]string, error) {
	var qds []string
	if err := db.New().Model(new(AdminQD)).Select("qd").Where("id = ?", id).Pluck("qd", &qds).Error; err != nil {
		return nil, err
	}
	if len(qds) == 0 {
		return nil, nil
	}

	return qds, nil
}

func AllQDs() ([]string, error) {
	var qds []string
	if err := db.New().Model(new(QDInstallRuns)).Select("qd").Group("qd").Pluck("qd", &qds).Error; err != nil {
		return nil, err
	}
	if len(qds) == 0 {
		return nil, nil
	}
	return qds, nil
}

func InstallRunsByQD(qds []string, limit, offset int) ([]QDInstallRuns, error) {
	var ins []QDInstallRuns
	if err := db.New().Where("qd in (?)", qds).Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func TotalInstallRunsByQD(qds []string) (int64, error) {
	var total int64
	if err := db.New().Model(new(QDInstallRuns)).Where("qd in (?)", qds).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func GetGroupCoefficient(groupName string) (*GroupCoefficient, error) {
	var c GroupCoefficient
	if err := db.New().Model(new(GroupCoefficient)).Where("group_name = ?", groupName).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
