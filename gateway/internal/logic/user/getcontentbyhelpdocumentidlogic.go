package user

import (
	"context"
	"fmt"
	"gateway/userclient"
	"strconv"

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

// 根据帮助文档id获取帮助文档内容
func (l *GetContentByHelpDocumentIdLogic) GetContentByHelpDocumentId(req *types.GetContentByHelpDocumentId) (resp *types.ContentByHelpDocumentIdReply, err error) {

	helpDocumentId, err := strconv.ParseInt(req.HelpDocumentId, 10, 64)
	helpResp, err := l.svcCtx.UserRpcClient.GetHelpDocumentTranslations(l.ctx, &userclient.GetHelpDocumentTranslationsRequest{
		HelpDocumentId: helpDocumentId,
		Language:       req.LanguageCode,
	})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	fmt.Println(helpResp)

	return &types.ContentByHelpDocumentIdReply{HelpDocumentId: helpResp.HelpDocumentId, DocumentContent: helpResp.DocumentContent, DocumentTitle: helpResp.DocumentTitle, LanguageCode: helpResp.Language}, nil

}
