package user

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpDocumentTitleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHelpDocumentTitleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpDocumentTitleListLogic {
	return &GetHelpDocumentTitleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHelpDocumentTitleListLogic) GetHelpDocumentTitleList(req *types.GetHelpDocumentTitleList) (resp *types.GetHelpDocumentTitleListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
