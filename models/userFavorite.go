package models

type UserFavorite struct {
	ID      int `gorm:"column:id;primary_key"  json:"id"`
	UserID  int `gorm:"column:user_id"  json:"user_id"`
	MovieID int `gorm:"column:movie_id"  json:"movie_id"`
}
