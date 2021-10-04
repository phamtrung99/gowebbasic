package controllers

import (
	"net/http"
	"strconv"
	"trungpham/gowebbasic/middlewares"
	"trungpham/gowebbasic/models"
	checkform "trungpham/gowebbasic/package/checkForm"
	"trungpham/gowebbasic/package/namestand"
	"trungpham/gowebbasic/repositories"

	echo "github.com/labstack/echo/v4"
)

type CmtControl struct {
	CmtRepo *repositories.CmtRepo
}

func NewCmtControl() *CmtControl {
	return &CmtControl{CmtRepo: repositories.NewCmtRepo()}
}

func (control *CmtControl) AddComment(c echo.Context) error {

	if !checkform.IsFormFullFill(c, []string{"movie_id", "content"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	var comment models.Comment
	arrTemp := []string{"movie_id", "content"}

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
				comment.MovieID, _ = strconv.Atoi(result)
			case 1:
				if namestand.IsContainBadWord(result) {
					return c.JSON(http.StatusBadRequest, echo.Map{
						"message": "content" + appErr.AuthMsg.IsHaveBadWord,
					})
				}
				comment.Content = result
			}
		}
	}

	//Check if sub comment
	if c.FormValue("parent_id") != "" {
		isID, result := checkform.CheckFormatValue("parent_id", c.FormValue("parent_id"))
		if !isID {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": result,
			})
		}
		comment.ParentID, _ = strconv.Atoi(result)
	}

	userID := middlewares.GetUserInfFromToken(c).ID

	comment.ActorID = userID

	isInsert := control.CmtRepo.InsertComment(&comment)
	if !isInsert {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": appErr.DatabaseMsg.InsertFail,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"result": comment,
	})
}

func (control *CmtControl) GetComment(c echo.Context) error {
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

	//Check if sub comment
	if c.FormValue("parent_id") != "" {
		isID, result := checkform.CheckFormatValue("parent_id", c.FormValue("parent_id"))
		if !isID {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": result,
			})
		}
		parentID, _ := strconv.Atoi(result)
		listSubComment := control.CmtRepo.GetSubCommentByParentId(parentID)

		if len(listSubComment) == 0 {
			return c.JSON(http.StatusOK, echo.Map{
				"success": false,
				"message": appErr.QueryMsg.ResourceNotFound,
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"result":       listSubComment,
			"total_result": len(listSubComment),
		})
	}

	listComment := control.CmtRepo.GetParentCommentByIDMovie(movieID)

	if len(listComment) == 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"success": false,
			"message": appErr.QueryMsg.ResourceNotFound,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"result":       listComment,
		"total_result": len(listComment),
	})
}

func (control *CmtControl) DeleteComment(c echo.Context) error {
	if !checkform.IsFormFullFill(c, []string{"comment_id"}) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": appErr.AuthMsg.NotEnoughInfo,
		})
	}

	isTrue, result := checkform.CheckFormatValue("comment_id", c.FormValue("comment_id"))
	if !isTrue {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": result,
		})
	}

	commentID, _ := strconv.Atoi(result)
	userID := middlewares.GetUserInfFromToken(c).ID
	userRole := middlewares.GetUserInfFromToken(c).Role

	comment := control.CmtRepo.GetCommentByID(commentID)
	if (models.Comment{}) == comment {
		return c.JSON(http.StatusOK, echo.Map{
			"success": false,
			"message": appErr.QueryMsg.ResourceNotFound,
		})
	}

	//Only comment author or admin have delete permission
	if userRole != "admin" {
		if comment.ActorID != userID {
			return c.JSON(http.StatusOK, echo.Map{
				"success": false,
				"message": appErr.QueryMsg.ResourceNotFound,
			})
		}
	}

	// check is contain subcomment and delete it
	if comment.ParentID == 1 {
		listSubComment := control.CmtRepo.GetSubCommentByParentId(commentID)
		if len(listSubComment) != 0 {
			for i := 0; i < len(listSubComment); i++ {
				isDel := control.CmtRepo.DeleteCommentById(listSubComment[i].ID)
				if !isDel {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"message": appErr.DatabaseMsg.DeleteFail,
					})
				}

			}
		}
	}

	isDel := control.CmtRepo.DeleteCommentById(commentID)

	if !isDel {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": appErr.DatabaseMsg.DeleteFail,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": appErr.DatabaseMsg.DeleteSuccess,
	})
}
