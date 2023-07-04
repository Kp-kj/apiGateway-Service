package mall

import (
	"context"
	"gateway/blockclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeBargainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJudgeBargainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeBargainLogic {
	return &JudgeBargainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JudgeBargainLogic) JudgeBargain(req *types.JudgeBargainInput) (resp *types.JudgeBargainReply, err error) {

	result, err := l.svcCtx.BlockClient.JudgeBargain(l.ctx, &blockclient.JudgeBargainRequest{
		UserId:            req.UserID,
		CryptominerTypeid: req.CryptominerTypeID,
	})

	return &types.JudgeBargainReply{IsBargain: result.IsBargain}, err

}
