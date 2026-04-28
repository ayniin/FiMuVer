package models

type Settings struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

func (Settings) TableName() string {
	return "settings"
}
