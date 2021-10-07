package auth

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/phamtrung99/gowebbasic/config"
	"github.com/phamtrung99/gowebbasic/middlewares"
	"github.com/phamtrung99/gowebbasic/model"
	"github.com/phamtrung99/gowebbasic/package/auth"
	checkform "github.com/phamtrung99/gowebbasic/package/checkForm"
	"github.com/phamtrung99/gowebbasic/util/myerror"
	"github.com/pkg/errors"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *Usecase) Login(ctx context.Context, req LoginRequest) (string, error) {
	//Check mail format
	isMail, email := checkform.CheckFormatValue("email", req.Email)
	if !isMail {
		return "", errors.Wrap(myerror.ErrValueFormat(nil), "email")
	}

	//Check password format
	isPass, password := checkform.CheckFormatValue("password", req.Password)
	if !isPass {
		return "", errors.Wrap(myerror.ErrValueFormat(nil), "password")
	}

	//Check mail exists
	user, err := u.userRepo.GetByEmail(ctx, email)

	if err != nil {
		return "", myerror.ErrLogin(err)
	}

	if (&model.User{}) == user {
		return "", myerror.ErrInvalid(nil)
	}

	//Check password is true.
	isPassTrue := auth.VerifyPassword(password, user.Password)

	if !isPassTrue {
		return "", myerror.ErrInvalid(nil)
	}

	claims := middlewares.NewClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GetConfig().Jwt.Key))

	if err != nil {
		return "", myerror.ErrToken(nil)
	}

	return t, nil
}
