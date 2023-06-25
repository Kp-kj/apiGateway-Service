package user

import (
	"context"

	"gateway/。/internal/svc"
	"gateway/。/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentByHelpDocumentIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContentByHelpDocumentIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentByHelpDocumentIdLogic {
	return &GetContentByHelpDocumentIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContentByHelpDocumentIdLogic) GetContentByHelpDocumentId(req *types.GetContentByHelpDocumentId) (resp *types.GetContentByHelpDocumentIdReply, err error) {
	// todo: add your logic here and delete this line

	return
}
