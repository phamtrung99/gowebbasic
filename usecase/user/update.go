package user

import (
	"context"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/gowebbasic/model"
	"github.com/phamtrung99/gowebbasic/package/auth"
	checkform "github.com/phamtrung99/gowebbasic/package/checkForm"
	imgvalid "github.com/phamtrung99/gowebbasic/package/fileValid"
	"github.com/phamtrung99/gowebbasic/util/myerror"
	"github.com/pkg/errors"
)

type UpdateInfoRequest struct {
	FormData *multipart.Form
}

func (u *Usecase) Update(ctx context.Context, req UpdateInfoRequest) (*model.User, error) {
	//Check password format
	isPass, password := checkform.CheckFormatValue("password", req.FormData.Value["password"][0])
	if !isPass {
		return &model.User{}, errors.Wrap(myerror.ErrValueFormat(nil), "password")
	}

	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	//Get current user info from userID
	user, err := u.userRepo.GetByID(ctx, userID)

	if err != nil {
		return &model.User{}, myerror.ErrUserGet(err)
	}

	//Check password is true.
	isPassTrue := auth.VerifyPassword(password, user.Password)

	if !isPassTrue {
		return &model.User{}, myerror.ErrInvalid(nil)
	}

	//Check email
	if req.FormData.Value["email"][0] != "" {
		formEmail := req.FormData.Value["email"][0]
		if user.Email != formEmail {
			isMail, email := checkform.CheckFormatValue("email", formEmail)
			if !isMail {
				return &model.User{}, errors.Wrap(myerror.ErrValueFormat(nil), "email")
			}

			userTemp, err := u.userRepo.GetByEmail(ctx, email)

			if err != nil {
				return &model.User{}, myerror.ErrUserGet(err)
			}

			if (&model.User{}) != userTemp {
				return &model.User{}, myerror.ErrExistedEmail(nil)
			}

			user.Email = email
		}
	}

	//Check full name
	if req.FormData.Value["full_name"][0] != "" {
		isName, fullName := checkform.CheckFormatValue("full_name", req.FormData.Value["full_name"][0])

		if !isName {
			return &model.User{}, errors.Wrap(myerror.ErrValueFormat(nil), "full_name")
		}

		user.FullName = fullName
	}

	if req.FormData.Value["age"][0] != "" {
		isAge, ageStr := checkform.CheckFormatValue("age", req.FormData.Value["age"][0])
		if !isAge {
			return &model.User{}, errors.Wrap(myerror.ErrValueFormat(nil), "age")
		}

		age, _ := strconv.Atoi(ageStr)
		user.Age = age
	}

	file := req.FormData.File["avatar"][0]

	if len(req.FormData.File["avatar"]) != 0 {
		var imgFileName = "blank.png"
		var pathFile = "/public/avatar/"
		var filetype = ""

		isImage, errMess, err := imgvalid.CheckImage(file)

		if !isImage {
			return &model.User{}, errors.Wrap(myerror.ErrImage(err), errMess)
		}

		imgFileName = strconv.FormatInt(user.ID, 10) + "." + strings.Split(filetype, "/")[1]

		err = imgvalid.CopyFile(file, imgFileName, "."+pathFile)

		if err != nil {
			log.Fatal(err)
		}

		user.Avatar = pathFile + imgFileName
	}

	result, err := u.userRepo.Update(ctx, user)

	if err != nil {
		return &model.User{}, myerror.ErrUserUpdate(err)
	}

	return result, nil
}
