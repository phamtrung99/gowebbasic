package models

type Producer struct {
	ID       int    `gorm:"column:id;primary_key"  json:"id"`
	LogoPath string `gorm:"column:logo_path"  json:"logo_path"`
	Name     string `gorm:"column:name"  json:"name"`
}
