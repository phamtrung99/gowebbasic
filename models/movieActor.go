package models

type MovieActor struct {
	ID            int    `gorm:"column:id;primary_key"  json:"id"`
	MovieID       int    `gorm:"column:movie_id"  json:"movie_id"`
	ActorID       int    `gorm:"column:actor_id"  json:"actor_id"`
	CharacterName string `gorm:"column:character_name"  json:"character_name"`
	IsMainActor   int    `gorm:"column:is_main_actor"  json:"is_main_actor"`
}
