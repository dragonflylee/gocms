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

func GetP2WInstallRuns(limit, offset int) ([]P2WInstallRuns, error) {
	var ins []P2WInstallRuns
	if err := db.New().Order("date desc").Limit(limit).Offset(offset).Find(&ins).Error; err != nil {
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
