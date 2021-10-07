package comment

import (
	"context"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/gowebbasic/util/myerror"
)

type DeleteCmtRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) Delete(ctx context.Context, req *DeleteCmtRequest) error {

	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	//Get userinfo to get role user
	user, err := u.userRepo.GetByID(ctx, userID)

	if err != nil {
		return err
	}

	//get comment by ID
	comment, err := u.cmtRepo.GetById(ctx, req.ID)

	if err != nil {
		return err
	}

	//if role user - can only delete current user comment.
	//if role admin - can delete all.

	if user.Role == "user" && comment.ActorID != userID {
		//fail if current user isn't actor of comment
		return myerror.ErrNotCmtActor(nil)
	}

	// check if this is root comment and delete subcomment
	if comment.ParentID == 1 {
		err := u.cmtRepo.DeleteSubCmt(ctx, comment.ID)

		if err != nil {
			return err
		}
	}

	//delete this comment
	err = u.cmtRepo.Delete(ctx, comment.ID)

	if err != nil {
		return err
	}

	return nil
}
