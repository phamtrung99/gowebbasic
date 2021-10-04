package repositories

import (
	"log"
	"strconv"

	"gorm.io/gorm"

	"trungpham/gowebbasic/models"
	"trungpham/gowebbasic/package/auth"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{db: ConnectMysqlInit()}
}

func (repo *UserRepo) CheckAccountExists(email string, password string) (models.User, int64) {
	query := repo.db.Table("users")
	var usr models.User

	result := query.Where("email = ?", email).
		Find(&usr)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		return models.User{}, -1
	}

	checkPass := auth.VerifyPassword(password, usr.Password)

	if !checkPass {
		return models.User{}, -2
	}

	return usr, result.RowsAffected
}

func (repo *UserRepo) DelUserByID(id int) bool {
	query := repo.db.Table("users")

	result := query.Where("ID = ?", id).
		Delete(&models.User{})

	if result.Error != nil {
		log.Fatal("Delete Error: ", result.Error)
	}

	log.Fatal(result)

	return true
}

func (repo *UserRepo) InsertUser(userIn *models.User) bool {
	query := repo.db.Table("users")

	if repo.IsExistedEmail(userIn.Email) {
		return false
	}

	resultIns := query.
		Create(&userIn)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	return true
}

func (repo *UserRepo) GetUserByEmail(email string) (*models.User, bool) {
	query := repo.db.Table("users")
	var user = models.User{}

	result := query.Where("email = ?", email).
		Find(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		return &models.User{}, false
	}

	return &user, true
}

func (repo *UserRepo) IsExistedEmail(email string) bool {
	query := repo.db.Table("users")
	var user = models.User{}

	result := query.Where("email = ?", email).
		Find(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (repo *UserRepo) GetUserByID(id int) (*models.User, bool) {
	query := repo.db.Table("users")
	var user = models.User{}

	result := query.Where("id = ?", id).
		Find(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		return &models.User{}, false
	}

	return &user, true
}

func (repo *UserRepo) UpdateUser(userUpd *models.User, id int) (*models.User, bool) {
	query := repo.db.Table("users")

	user, isExisted := repo.GetUserByID(id)

	if !isExisted {
		return &models.User{}, false
	}

	resultIns := query.Model(&user).
		Omit("role", "id", "deleted_at", "password", "ref_token").
		Updates(&userUpd)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	user, isExisted = repo.GetUserByID(id)

	if !isExisted {
		return &models.User{}, false
	}

	return user, isExisted
}

func (repo *UserRepo) GetNextIDIncrement() string {
	var nextID int
	result := repo.db.Raw(`
	SELECT AUTO_INCREMENT
	FROM information_schema.TABLES
	WHERE TABLE_SCHEMA = "the_movie_db"
	AND TABLE_NAME = "users"`).Scan(&nextID)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	nextID += 1

	return strconv.Itoa(nextID)
}

func (repo *UserRepo) UpdateUserPassword(passwHash string, id int) bool {
	query := repo.db.Table("users")

	result := query.Where("id = ?", id).
		Updates(&models.User{Password: passwHash})

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (repo *UserRepo) InsertFavoriteUserMovies(userID int, movieID int) (*models.UserFavorite, bool) {
	query := repo.db.Table("user_favorites")
	var userFavor = models.UserFavorite{
		UserID:  userID,
		MovieID: movieID,
	}

	isExists := query.Where("user_id = ? AND movie_id = ? ", userID, movieID).
		Find(&models.UserFavorite{}).RowsAffected

	if isExists != 0 {
		return &models.UserFavorite{}, false
	}

	resultIns := query.
		Create(&userFavor)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	return &userFavor, true
}

func (repo *UserRepo) GetFavoriteUserMovies(userID int) []int {
	query := repo.db.Table("user_favorites")
	var userFavor []models.UserFavorite
	var result = []int{}

	resultIns := query.Where("user_id = ? ", userID).
		Find(&userFavor)

	if resultIns.Error != nil {
		log.Fatal(resultIns.Error)
	}

	if resultIns.RowsAffected == 0 {
		return result
	}

	for i := 0; i < len(userFavor); i++ {
		result = append(result, userFavor[i].MovieID) 
	}
	

	return result
}
