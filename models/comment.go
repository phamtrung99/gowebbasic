package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `gorm:"column:id;primary_key"  json:"id"`
	ParentID  int            `gorm:"column:parent_id; default:1"  json:"parent_id"`
	ActorID   int            `gorm:"column:actor_id"  json:"actor_id"`
	MovieID   int            `gorm:"column:movie_id"  json:"movie_id"`
	Content   string         `gorm:"column:content"  json:"content"`
	CreatedAt time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}
