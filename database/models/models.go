package models



//create log model for postgres database
type Log struct {
	ID        uint   `gorm:"primaryKey"`
	AppKey    string `gorm:"type:varchar(255)"`
	Logs      string `gorm:"type:text"`
	CreatedAt string `gorm:"type:varchar(255)"`
}

//create foreignkey for Appkey id in Log struct
func (Log) TableName() string {
	return "logs"
}


