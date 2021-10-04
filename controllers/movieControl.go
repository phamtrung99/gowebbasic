package controllers

import (
	"net/http"
	"strconv"
	"time"
	"trungpham/gowebbasic/config"
	"trungpham/gowebbasic/models"
	checkform "trungpham/gowebbasic/package/checkForm"
	"trungpham/gowebbasic/package/pagination"
	"trungpham/gowebbasic/repositories"

	echo "github.com/labstack/echo/v4"
)

type MovieControl struct {
	MovieRepo *repositories.MovieRepo
}

func NewMovieControl() *MovieControl {
	return &MovieControl{MovieRepo: repositories.NewMovieRepo()}
}

func (control *MovieControl) SearchMovie(c echo.Context) error {
	isPage, msg, page := pagination.GetPageQueryParam(c)
	if !isPage {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": msg,
		})
	}
	var filter models.Filter
	searchText := ""
	arrTemp := []string{"is_adult", "actor_id", "cate_id", "min_rating", "search"}

	for i := 0; i < 5; i++ {
		if c.QueryParam(arrTemp[i]) != "" {
			isTrue, result := checkform.CheckFormatValue(arrTemp[i], c.QueryParam(arrTemp[i]))
			if !isTrue {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": result,
				})
			}
			switch i {
			case 0:
				filter.IsAdult = result
			case 1:
				filter.ActorID = result
			case 2:
				filter.CateID = result
			case 3:
				filter.MinRating = result
			case 4:
				searchText = result
			}
		}
	}

	movieList, totalRow := control.MovieRepo.SearchMovie(searchText, filter, page)

	if len(movieList) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success":        false,
			"status_message": appErr.QueryMsg.ResourceNotFound,
		})
	}

	rowPerPage := int64(config.GetPagination().MovieRowPerPage)

	return c.JSON(http.StatusCreated, echo.Map{
		"page":       page,
		"result":     movieList,
		"total_page": pagination.CountTotalPage(totalRow, rowPerPage),
	})
}

func (control *MovieControl) InsertMovie(c echo.Context) error {
	arrTemp := []string{"is_adult", "image",
		"original_language", "original_title", "popularity", "overview", "movie_link",
		"release_date", "duration", "spoken_language", "rating_average", "status"}

	if !checkform.IsFormFullFill(c, arrTemp) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	var movie models.Movie

	for i := 0; i < len(arrTemp); i++ {
		if c.FormValue(arrTemp[i]) != "" {
			isTrue, result := checkform.CheckFormatValue(arrTemp[i], c.FormValue(arrTemp[i]))
			if !isTrue {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": result,
				})
			}
			switch i {
			case 0:
				movie.IsAdult, _ = strconv.Atoi(result)
			case 1:
				movie.Image = result
			case 2:
				movie.OriginalLanguage = result
			case 3:
				movie.OriginalTitle = result
			case 4:
				movie.Popularity, _ = strconv.ParseFloat(result, 64)
			case 5:
				movie.Overview = result
			case 6:
				movie.MovieLink = result
			case 7:
				t, err := time.Parse(time.RFC3339, result)
				if err != nil {
					panic(err)
				}
				movie.ReleaseDate = t
			case 9:
				movie.Duration, _ = strconv.Atoi(result)
			case 10:
				movie.SpokenLanguage = result
			case 11:
				movie.RatingAverage, _ = strconv.ParseFloat(result, 64)
			case 12:
				movie.Status, _ = strconv.Atoi(result)
			}
		}
	}

	isInsert := control.MovieRepo.InsertMovie(movie)

	if !isInsert {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success":        false,
			"status_message": appErr.DatabaseMsg.InsertFail,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"result": movie,
	})
}

func (control *MovieControl) DeleteMovie(c echo.Context) error {

	if !checkform.IsFormFullFill(c, []string{"movie_id"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	isTrue, result := checkform.CheckFormatValue("movie_id", c.FormValue("movie_id"))
	if !isTrue {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": result,
		})
	}
	movieID, _ := strconv.Atoi(result)

	movie := control.MovieRepo.GetMovieById(movieID)

	if (&models.Movie{}) == movie {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success":        false,
			"status_message": appErr.QueryMsg.ResourceNotFound,
		})
	}

	isDel := control.MovieRepo.DeleteMovie(movie)

	if !isDel {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success":        false,
			"status_message": appErr.DatabaseMsg.DeleteFail,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": appErr.DatabaseMsg.DeleteSuccess,
	})
}

func (control *MovieControl) UpdateMovie(c echo.Context) error {
	arrTemp := []string{"is_adult", "image",
		"original_language", "original_title", "popularity", "overview", "movie_link",
		/*"release_date",*/ "duration", "spoken_language", "rating_average", "status"}

	if !checkform.IsFormFullFill(c, []string{"movie_id"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	isID, result := checkform.CheckFormatValue("movie_id", c.FormValue("movie_id"))

	if !isID {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": result,
		})
	}

	movieID, _ := strconv.Atoi(result)

	movie := control.MovieRepo.GetMovieById(movieID)

	if (&models.Movie{}) == movie {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success":        false,
			"status_message": appErr.QueryMsg.ResourceNotFound,
		})
	}

	for i := 0; i < len(arrTemp); i++ {
		if c.FormValue(arrTemp[i]) != "" {
			isTrue, result := checkform.CheckFormatValue(arrTemp[i], c.FormValue(arrTemp[i]))
			if !isTrue {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": result,
				})
			}
			switch i {
			case 0:
				movie.IsAdult, _ = strconv.Atoi(result)
			case 1:
				movie.Image = result
			case 2:
				movie.OriginalLanguage = result
			case 3:
				movie.OriginalTitle = result
			case 4:
				movie.Popularity, _ = strconv.ParseFloat(result, 64)
			case 5:
				movie.Overview = result
			case 6:
				movie.MovieLink = result
			//case 7:
			// movie.ReleaseDate, _ = time.Parse("30-12-2020", result)
			// log.Fatal(time.Parse("2021-09-29 02:26:52", result))

			case 7:
				movie.Duration, _ = strconv.Atoi(result)
			case 8:
				movie.SpokenLanguage = result
			case 9:
				movie.RatingAverage, _ = strconv.ParseFloat(result, 64)
			case 10:
				movie.Status, _ = strconv.Atoi(result)
			}
		}
	}
	isUpdate := control.MovieRepo.UdpateMovie(movie)

	if !isUpdate {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success":        false,
			"status_message": appErr.DatabaseMsg.UpdateFail,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"success": true,
		"result":  movie,
	})
}
