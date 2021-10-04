package models

type MovieProducer struct {
	ID         int `gorm:"column:id;primary_key"  json:"id"`
	MovieID    int `gorm:"column:movie_id"  json:"movie_id"`
	ProducerID int `gorm:"column:producer_id"  json:"producer_id"`
}
