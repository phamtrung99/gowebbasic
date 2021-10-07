package comment

import (
	"context"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/gowebbasic/model"
	checkform "github.com/phamtrung99/gowebbasic/package/checkForm"
	"github.com/phamtrung99/gowebbasic/package/namestand"
	"github.com/phamtrung99/gowebbasic/util/myerror"
	"github.com/pkg/errors"
)

type InsertCmtRequest struct {
	ParentID int64  `json:"parent_id"`
	ActorID  int64  `json:"actor_id"`
	MovieID  int64  `json:"movie_id"`
	Content  string `json:"content"`
}

func (u *Usecase) Insert(ctx context.Context, req InsertCmtRequest) (*model.Comment, error) {

	comment := &model.Comment{}

	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	//Check content format
	isTrue, msg := checkform.CheckFormatValue("content", req.Content)

	if !isTrue {
		return nil, errors.Wrap(myerror.ErrValueFormat(nil), msg)
	}

	//Check content has bad word
	if namestand.IsContainBadWord(msg) {
		return nil, myerror.ErrBadWordContent(nil)
	}

	comment.Content = msg

	//Check if insert subcomment
	if req.ParentID != 0 {
		comment.ParentID = req.ParentID
	}

	comment.ActorID = userID

	result, err := u.cmtRepo.Insert(ctx, comment)

	if err != nil {
		return nil, err
	}

	return result, nil
}
