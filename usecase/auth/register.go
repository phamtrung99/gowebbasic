package auth

import (
	"context"

	"github.com/phamtrung99/gowebbasic/model"
	"github.com/phamtrung99/gowebbasic/package/auth"
	checkform "github.com/phamtrung99/gowebbasic/package/checkForm"
	"github.com/phamtrung99/gowebbasic/util/myerror"
	"github.com/pkg/errors"
)

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (u *Usecase) Register(ctx context.Context, req RegisterRequest) (*model.User, error) {

	//Check mail format
	isMail, email := checkform.CheckFormatValue("email", req.Email)
	if !isMail {
		return &model.User{}, errors.Wrap(myerror.ErrValueFormat(nil), "email")
	}

	//Check password format
	isPass, password := checkform.CheckFormatValue("password", req.Password)
	if !isPass {
		return &model.User{}, errors.Wrap(myerror.ErrValueFormat(nil), "password")
	}

	//Check fullname format
	isName, fullName := checkform.CheckFormatValue("full_name", req.FullName)
	if !isName {
		return &model.User{}, errors.Wrap(myerror.ErrValueFormat(nil), "full name")
	}

	passwHash, err := auth.HashPassword(password)

	if err != nil {
		return &model.User{}, myerror.ErrHashPassword(nil)
	}

	var user = &model.User{
		FullName: fullName,
		Email:    email,
		Password: passwHash,
		Age:      req.Age,
		Role:     "user",
	}

	result, err := u.userRepo.Create(ctx, user)

	if err != nil {
		return &model.User{}, myerror.ErrUserCreate(err)
	}

	return result, nil
}
