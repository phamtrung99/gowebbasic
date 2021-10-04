package repositories

import (
	"log"

	"github.com/phamtrung99/gowebbasic/models"

	"gorm.io/gorm"
)

type CmtRepo struct {
	db *gorm.DB
}

func NewCmtRepo() *CmtRepo {
	return &CmtRepo{db: ConnectMysqlInit()}
}

func (repo *CmtRepo) InsertComment(comment *models.Comment) bool {
	query := repo.db.Table("comments")

	resultIns := query.
		Create(&comment)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	return true
}

func (repo *CmtRepo) GetParentCommentByIDMovie(movieID int) []models.Comment {
	query := repo.db.Table("comments")
	var listComment []models.Comment

	result := query.Where("movie_id = ? AND parent_id = 1", movieID).
		Find(&listComment)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		return listComment
	}

	return listComment
}

func (repo *CmtRepo) GetSubCommentByParentId(parentID int) []models.Comment {
	query := repo.db.Table("comments")
	var listSubComment []models.Comment

	result := query.Where("parent_id = ?", parentID).
		Find(&listSubComment)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return listSubComment
}

func (repo *CmtRepo) DeleteCommentById(commentID int) bool {
	query := repo.db.Table("comments")
	var comment models.Comment
	comment.ID = commentID

	result := query.Delete(&comment)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (repo *CmtRepo) GetCommentByID(commentID int) models.Comment {
	query := repo.db.Table("comments")
	var comment models.Comment

	result := query.Where("id = ?", commentID).
		Find(&comment)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return comment
}
