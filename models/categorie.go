package models

type Categorie struct {
	ID   int    `gorm:"column:id;primary_key"  json:"id"`
	Name string `gorm:"column:name"  json:"name"`
}
