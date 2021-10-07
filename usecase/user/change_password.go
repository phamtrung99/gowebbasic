package user

import (
	"context"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/gowebbasic/package/auth"
	checkform "github.com/phamtrung99/gowebbasic/package/checkForm"
	"github.com/phamtrung99/gowebbasic/util/myerror"
	"github.com/pkg/errors"
)

type ChangePasswRequest struct {
	OldPassword   string `json:"old_password"`
	NewPassword   string `json:"new_password"`
	RenewPassword string `json:"renew_password"`
}

func (u *Usecase) ChangePassword(ctx context.Context, req ChangePasswRequest) error {

	if req.NewPassword != req.RenewPassword {
		return myerror.ErrPwdMatching(nil)
	}

	isPass, newPwd := checkform.CheckFormatValue("new_password", req.NewPassword)
	if !isPass {
		return errors.Wrap(myerror.ErrValueFormat(nil), "new password")
	}

	isPass, oldPwd := checkform.CheckFormatValue("old_password", req.OldPassword)
	if !isPass {
		return errors.Wrap(myerror.ErrValueFormat(nil), "old password")
	}

	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	//Get current user info from userID
	user, err := u.userRepo.GetByID(ctx, userID)

	if err != nil {
		return err
	}

	//Check old password is true.
	isPassTrue := auth.VerifyPassword(oldPwd, user.Password)

	if !isPassTrue {
		return errors.Wrap(myerror.ErrInvalid(nil), "old password")
	}

	passHash, err := auth.HashPassword(newPwd)

	if err != nil {
		return myerror.ErrHashPassword(err)
	}

	user.Password = passHash

	_, err = u.userRepo.Update(ctx, user)

	if err != nil {
		return errors.Wrap(myerror.ErrUserUpdate(nil), "update password")
	}
	
	return nil
}
