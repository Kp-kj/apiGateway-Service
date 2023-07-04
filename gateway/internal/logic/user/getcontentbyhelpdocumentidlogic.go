package user

import (
	"context"
	"fmt"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

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

func (l *GetContentByHelpDocumentIdLogic) GetContentByHelpDocumentId(req *types.GetContentByHelpDocumentId) (resp *types.ContentByHelpDocumentIdReply, err error) {
	// todo: add your logic here and delete this line
	helpResp, err := l.svcCtx.UserRpcClient.GetHelpDocumentTranslations(l.ctx, &userclient.GetHelpDocumentTranslationsRequest{
		HelpDocumentId: req.HelpDocumentId,
		Language:       req.LanguageCode,
	})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	fmt.Println(helpResp)

	return &types.ContentByHelpDocumentIdReply{HelpDocumentId: helpResp.HelpDocumentId, DocumentContent: helpResp.DocumentContent, DocumentTitle: helpResp.DocumentTitle, LanguageCode: helpResp.Language}, nil

}
