package models

type MovieCategorie struct {
	ID      int `gorm:"column:id;primary_key"  json:"id"`
	MovieID int `gorm:"column:movie_id"  json:"movie_id"`
	CateID  int `gorm:"column:cate_id"  json:"cate_id"`
}
