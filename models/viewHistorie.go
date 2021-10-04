package models

import "time"

type ViewHistorie struct {
	ID                int       `gorm:"column:id;primary_key"  json:"id"`
	UserID            int       `gorm:"column:user_id"  json:"user_id"`
	MovieID           int       `gorm:"column:movie_id"  json:"movie_id"`
	LastMoviesRuntime int       `gorm:"column:last_movies_runtime"  json:"last_movies_runtime"`
	LastViewAt        time.Time `gorm:"column:last_view_at"  json:"last_view_at"`
}
