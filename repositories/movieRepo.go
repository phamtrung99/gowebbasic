package repositories

import (
	"fmt"
	"log"

	"github.com/phamtrung99/gowebbasic/config"
	"github.com/phamtrung99/gowebbasic/models"

	"gorm.io/gorm"
)

type MovieRepo struct {
	db *gorm.DB
}

func NewMovieRepo() *MovieRepo {
	return &MovieRepo{db: ConnectMysqlInit()}
}

func (repo *MovieRepo) GetMovieById(ID int) *models.Movie {
	query := repo.db.Table("movies")
	var movie = models.Movie{}

	result := query.Where("id = ?", ID).
		Find(&movie)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return &movie
}

func (repo *MovieRepo) GetListMovieByListId(listMovieID []int, page int) ([]models.Movie, int64) {
	rowPerPage := config.GetPagination().MovieRowPerPage
	query := repo.db.Table("movies")
	var movie = []models.Movie{}

	totalResult := query.Find(&movie, listMovieID).RowsAffected

	if totalResult == 0 {
		return []models.Movie{}, 0
	}

	result := query.Limit(rowPerPage).Offset((page-1)*rowPerPage).
		Find(&movie, listMovieID)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return movie, totalResult
}

func (repo *MovieRepo) GetMovieByName(str string, page int) ([]models.Movie, int64) {

	rowPerPage := config.GetPagination().MovieRowPerPage
	query := repo.db.Table("movies")
	var movie = []models.Movie{}

	totalResult := query.Where("original_title LIKE ?", "%"+str+"%").Scan(&[]models.Movie{}).RowsAffected

	if totalResult == 0 {
		return []models.Movie{}, 0
	}

	result := query.Limit(rowPerPage).Offset((page-1)*rowPerPage).
		Where("original_title LIKE ?", "%"+str+"%").
		Find(&movie)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return movie, totalResult
}

func (repo *MovieRepo) SearchMovie(str string, filter models.Filter, page int) ([]models.Movie, int64) {
	rowPerPage := config.GetPagination().MovieRowPerPage
	var movie = []models.Movie{}

	filterActor := ""
	filterCate := ""
	filterRate := ""
	filterAdult := ""
	filterName := ""

	if filter.ActorID != "" {
		filterActor = "JOIN movie_actors ON movie_actors.movie_id = movies.id AND movie_actors.actor_id = " + filter.ActorID
	}

	if filter.CateID != "" {
		filterCate = "JOIN movie_categories ON movie_categories.movie_id = movies.id AND movie_categories.cate_id = " + filter.CateID
	}

	if filter.IsAdult != "" {
		filterAdult = "AND is_adult = " + filter.IsAdult
	}

	if filter.MinRating != "" {
		filterRate = "AND rating_average > " + filter.MinRating
	}

	if str != "" {
		filterName = "AND original_title LIKE " + "'%" + str + "%'"
	}

	limitQuery := " LIMIT " + fmt.Sprint(rowPerPage) + " OFFSET " + fmt.Sprint((page-1)*rowPerPage)
	queryStr := "SELECT * FROM movies " +
		filterActor + " " +
		filterCate + " WHERE isnull(movies.deleted_at) " +
		filterAdult + " " +
		filterRate + " " +
		filterName

	totalResult := repo.db.Raw(queryStr).Scan(&[]models.Movie{}).RowsAffected

	if totalResult == 0 {
		return []models.Movie{}, 0
	}

	result := repo.db.
		Raw(queryStr + " " + limitQuery).Find(&movie)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return movie, totalResult
}

func (repo *MovieRepo) InsertMovie(movie models.Movie) bool {
	query := repo.db.Table("movies")

	resultIns := query.
		Create(&movie)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	if resultIns.RowsAffected == 0 {
		return false
	}

	return true
}

func (repo *MovieRepo) DeleteMovie(movie *models.Movie) bool {
	query := repo.db.Table("movies")

	resultIns := query.
		Delete(&movie)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	if resultIns.RowsAffected == 0 {
		return false
	}

	return true
}

func (repo *MovieRepo) UdpateMovie(movie *models.Movie) bool {
	query := repo.db.Table("movies")

	resultIns := query.
		Save(&movie)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	if resultIns.RowsAffected == 0 {
		return false
	}

	return true
}
