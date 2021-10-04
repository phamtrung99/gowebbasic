package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"

	"github.com/phamtrung99/gowebbasic/config"
	"github.com/phamtrung99/gowebbasic/middlewares"
	"github.com/phamtrung99/gowebbasic/models"
	"github.com/phamtrung99/gowebbasic/package/auth"
	checkform "github.com/phamtrung99/gowebbasic/package/checkForm"
	imgvalid "github.com/phamtrung99/gowebbasic/package/fileValid"
	errmsg "github.com/phamtrung99/gowebbasic/package/logMessages"
	"github.com/phamtrung99/gowebbasic/package/pagination"
	"github.com/phamtrung99/gowebbasic/repositories"
)

type UserControl struct {
	UserRep *repositories.UserRepo
}

func NewUserControl() *UserControl {
	return &UserControl{UserRep: repositories.NewUserRepo()}
}

var appErr = errmsg.InitErrMsg()

func (control *UserControl) Login(c echo.Context) error {
	if !checkform.IsFormFullFill(c, []string{"email", "password"}) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	isMail, email := checkform.CheckFormatValue("email", c.FormValue("email"))
	if !isMail {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": email,
		})
	}

	isPass, password := checkform.CheckFormatValue("password", c.FormValue("password"))
	if !isPass {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": password,
		})
	}

	user, count := control.UserRep.CheckAccountExists(email, password)

	if count == -1 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": appErr.AuthMsg.InvalidEmail,
		})
	}

	if count == -2 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": appErr.AuthMsg.InvalidPassword,
		})
	}

	claims := middlewares.NewClaims(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GetSecretKey()))

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (control *UserControl) Register(c echo.Context) error {

	if !checkform.IsFormFullFill(c, []string{"email", "password", "age", "full_name"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}
	isMail, email := checkform.CheckFormatValue("email", c.FormValue("email"))
	if !isMail {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": email,
		})
	}
	isPass, password := checkform.CheckFormatValue("password", c.FormValue("password"))
	if !isPass {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": password,
		})
	}
	isName, fullName := checkform.CheckFormatValue("full_name", c.FormValue("full_name"))
	if !isName {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": fullName,
		})
	}
	isAge, ageStr := checkform.CheckFormatValue("age", c.FormValue("age"))
	if !isAge {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": ageStr,
		})
	}
	age, _ := strconv.Atoi(ageStr)

	passHash, err := auth.HashPassword(password)

	if err != nil {
		log.Fatal(err)
	}

	//avatar
	file, isSend := c.FormFile("avatar")

	var imgFileName = "blank.png"
	var pathFile = "/public/avatar/"
	var filetype = ""

	if isSend == nil {
		isImage, errMess, err := imgvalid.CheckImage(file)

		if !isImage {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": errMess,
			})
		} else {
			filetype = errMess
		}

		if err != nil {
			log.Fatal(errMess)
		}

		imgFileName = control.UserRep.GetNextIDIncrement() + "." + strings.Split(filetype, "/")[1]
		pathFile = "/public/avatar/"

		err = imgvalid.CopyFile(file, imgFileName, "."+pathFile)

		if err != nil {
			log.Fatal(err)
		}
	}

	var userIn = models.User{
		FullName: fullName,
		Email:    email,
		Avatar:   pathFile + imgFileName,
		Password: passHash,
		Age:      age,
		Role:     "user",
	}

	isInserted := control.UserRep.InsertUser(&userIn)

	if !isInserted {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": appErr.AuthMsg.ExistedEmail,
		})
	}

	userInsert, isExists := control.UserRep.GetUserByEmail(userIn.Email)

	if !isExists {
		log.Fatal(appErr.DatabaseMsg.SelectFail)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"user_insert": userInsert,
	})
}

func (control *UserControl) IsRightPassword(password string, c echo.Context) bool {
	UserID := middlewares.GetUserInfFromToken(c).ID

	user, isExisted := control.UserRep.GetUserByID(UserID)

	if !isExisted {
		return false
	}

	return auth.VerifyPassword(password, user.Password)
}

func (control *UserControl) UpdateUser(c echo.Context) error {

	if !checkform.IsFormFullFill(c, []string{"password"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	isPass, password := checkform.CheckFormatValue("password", c.FormValue("password"))

	if !isPass {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": password,
		})
	}

	if !control.IsRightPassword(password, c) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": appErr.AuthMsg.InvalidPassword,
		})
	}

	//log.Fatal("dasdsdas")

	var user models.User

	if c.FormValue("email") != "" {
		userTemp, isExist := control.UserRep.GetUserByID(middlewares.GetUserInfFromToken(c).ID)

		if isExist && userTemp.Email != c.FormValue("email") {
			if control.UserRep.IsExistedEmail(c.FormValue("email")) {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": appErr.AuthMsg.ExistedEmail,
				})
			}

			isMail, email := checkform.CheckFormatValue("email", c.FormValue("email"))

			if !isMail {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": email,
				})
			}
			user.Email = email
		}

	}

	if c.FormValue("full_name") != "" {
		isName, fullName := checkform.CheckFormatValue("full_name", c.FormValue("full_name"))

		if !isName {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": fullName,
			})
		}
		user.FullName = fullName
	}

	if c.FormValue("age") != "" {
		isAge, ageStr := checkform.CheckFormatValue("age", c.FormValue("age"))
		if !isAge {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": ageStr,
			})
		}

		age, _ := strconv.Atoi(ageStr)
		user.Age = age
	}
	//log.Fatal(user)

	file, isSend := c.FormFile("avatar")

	if isSend == nil {
		var imgFileName = "blank.png"
		var pathFile = "/public/avatar/"
		var filetype = ""
		//err := os.Remove(pathFile + middlewares.GetUserInfFromToken(c).ID)

		isImage, errMess, err := imgvalid.CheckImage(file)

		if !isImage {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": errMess,
			})
		} else {
			filetype = errMess
		}

		if err != nil {
			log.Fatal(err)
		}

		imgFileName = strconv.Itoa(middlewares.GetUserInfFromToken(c).ID) + "." + strings.Split(filetype, "/")[1]
		pathFile = "/public/avatar/"

		err = imgvalid.CopyFile(file, imgFileName, "."+pathFile)

		if err != nil {
			log.Fatal(err)
		}

		user.Avatar = pathFile + imgFileName
	}

	UserID := middlewares.GetUserInfFromToken(c).ID
	result, isUpdate := control.UserRep.UpdateUser(&user, UserID)

	if !isUpdate {
		log.Fatal(appErr.DatabaseMsg.UpdateFail)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"user_update": result,
	})
}

func (control *UserControl) ChangePassword(c echo.Context) error {
	if !checkform.IsFormFullFill(c, []string{"old_password", "new_password", "renew_password"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	if c.FormValue("new_password") != c.FormValue("renew_password") {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NewPasswordNotMatch,
		})
	}

	isPass, newPassw := checkform.CheckFormatValue("new_password", c.FormValue("new_password"))
	if !isPass {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": newPassw,
		})
	}

	isPass, oldPassw := checkform.CheckFormatValue("old_password", c.FormValue("old_password"))
	if !isPass {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": oldPassw,
		})
	}

	if !control.IsRightPassword(oldPassw, c) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": appErr.AuthMsg.InvalidPassword,
		})
	}

	passHash, err := auth.HashPassword(newPassw)

	if err != nil {
		log.Fatal(err)
	}

	isUpdate := control.UserRep.UpdateUserPassword(passHash, middlewares.GetUserInfFromToken(c).ID)

	if !isUpdate {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": appErr.AuthMsg.ChangePassFail,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": appErr.AuthMsg.ChangePassOK,
	})
}

func (control *UserControl) AddMovieToFavorite(c echo.Context) error {

	if !checkform.IsFormFullFill(c, []string{"movie_id"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	isID, ID := checkform.CheckFormatValue("movie_id", c.FormValue("movie_id"))
	if !isID {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": ID,
		})
	}

	userID := middlewares.GetUserInfFromToken(c).ID
	movieID, _ := strconv.Atoi(ID)

	result, isInsert := control.UserRep.InsertFavoriteUserMovies(userID, movieID)
	if !isInsert {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Exists movie in favorite movie list.",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"result": result,
	})
}

func (control *UserControl) GetFavoriteMovie(c echo.Context) error {
	userID := middlewares.GetUserInfFromToken(c).ID
	isPage, msg, page := pagination.GetPageQueryParam(c)
	if !isPage {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": msg,
		})
	}

	listMovieID := control.UserRep.GetFavoriteUserMovies(userID)

	listMovie, totalRow := repositories.NewMovieRepo().GetListMovieByListId(listMovieID, page)

	if len(listMovie) == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": appErr.DatabaseMsg.InsertFail,
		})
	}

	rowPerPage := int64(config.GetPagination().MovieRowPerPage)

	return c.JSON(http.StatusCreated, echo.Map{
		"page":       page,
		"result":     listMovie,
		"total_page": pagination.CountTotalPage(totalRow, rowPerPage),
	})
}
