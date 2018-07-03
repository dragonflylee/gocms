package model

type P2WInstallRuns struct {
	ID            int64 `gorm:"primary_key;auto_increment"`
	Date          string
	DownloadStart int64
	DownloadEnd   int64
	InstallStart  int64
	InstallEnd    int64
	Run           int64
}

func (*P2WInstallRuns) TableName() string {
	return "pdf2word"
}

type QDP2WInstallRuns struct {
	ID            int64 `gorm:"primary_key;auto_increment"`
	QD            string
	Date          string
	DownloadStart int64
	DownloadEnd   int64
	InstallStart  int64
	InstallEnd    int64
	Run           int64
}

func (*QDP2WInstallRuns) TableName() string {
	return "qd_pdf2word"
}

func GetP2WInstallRuns(limit, offset int) ([]P2WInstallRuns, error) {
	var ins []P2WInstallRuns
	if err := db.New().Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func GetP2WInstallRunsGroupByQD(qd string, limit, offset int) ([]QDP2WInstallRuns, error) {
	var ins []QDP2WInstallRuns
	db := db.New()
	if qd != "" {
		db = db.Where("qd = ?", qd)
	}
	if err := db.Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func TotalP2WInstallRuns() (int64, error) {
	var total int64
	if err := db.New().Model(new(P2WInstallRuns)).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func TotalP2WInstallRunsGroupByQD(qd string) (int64, error) {
	var total int64
	db := db.New()
	if qd != "" {
		db = db.Where("qd = ?", qd)
	}
	if err := db.Model(new(QDP2WInstallRuns)).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
