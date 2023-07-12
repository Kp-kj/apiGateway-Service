package admin

import (
	"context"
	"gateway/userclient/user"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteDocumentContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchDeleteDocumentContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteDocumentContentLogic {
	return &BatchDeleteDocumentContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//   BatchDeleteDocumentContent 批量删除文档内容
func (l *BatchDeleteDocumentContentLogic) BatchDeleteDocumentContent(req *types.BatchDeleteDocumentContent) (resp *types.BatchDeleteDocumentContentReply, err error) {
	for _, v := range req.Helpdocumentids {
		//批量删除文档内容
		_, err = l.svcCtx.UserRpcClient.DeleteHelpDocument(l.ctx, &user.DeleteHelpDocumentRequest{HelpDocumentId: v})
		if err != nil {
			return nil, err
		}
		// 批量删除文档内容翻译
		_, err = l.svcCtx.UserRpcClient.DeleteHelpDocumentTranslation(l.ctx, &user.DeleteHelpDocumentTranslationRequest{HelpDocumentId: v, Language: "zh"})
		if err != nil {
			return nil, err
		}

		_, err = l.svcCtx.UserRpcClient.DeleteHelpDocumentTranslation(l.ctx, &user.DeleteHelpDocumentTranslationRequest{HelpDocumentId: v, Language: "en"})
		if err != nil {
			return nil, err
		}
	}

	return &types.BatchDeleteDocumentContentReply{
		IsSuccess: true,
	}, nil
}
