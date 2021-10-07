package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64          `gorm:"column:id;primary_key"  json:"id"`
	ParentID  int64          `gorm:"column:parent_id; default:1"  json:"parent_id"`
	ActorID   int64          `gorm:"column:actor_id"  json:"actor_id"`
	MovieID   int64          `gorm:"column:movie_id"  json:"movie_id"`
	Content   string         `gorm:"column:content"  json:"content"`
	CreatedAt time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}

type CommentResult struct {
	Data []Comment `json:"data"`
	Paginator
}

type CommentFilter struct {
	MovieID  int64 `json:"movie_id"`
	ParentID int64 `json:"parent_id"`
}
