package admin

import (
	"context"
	"gateway/userclient/user"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditDocumentContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditDocumentContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditDocumentContentLogic {
	return &EditDocumentContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// EditDocumentContent 修改文档内容
func (l *EditDocumentContentLogic) EditDocumentContent(req *types.EditDocumentContent) (resp *types.EditDocumentContentReply, err error) {
	if req.Helpdocumentid == 0 {
		//新增
		dataDocument, err := l.svcCtx.UserRpcClient.CreateHelpDocument(l.ctx, &user.CreateHelpDocumentRequest{
			HelpCategoryId: req.HelpCategoryId,
			DocumentStatus: 1, //0 上架，1 下架
		})
		if err != nil {
			return nil, err
		}
		//新增文档内容
		_, err = l.svcCtx.UserRpcClient.CreateHelpDocumentTranslation(l.ctx, &user.CreateHelpDocumentTranslationRequest{
			HelpDocumentId:  dataDocument.HelpDocumentId,
			Language:        "zh",
			DocumentContent: req.ContentZh,
			DocumentTitle:   req.TitleZh,
		})
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.UserRpcClient.CreateHelpDocumentTranslation(l.ctx, &user.CreateHelpDocumentTranslationRequest{
			HelpDocumentId:  dataDocument.HelpDocumentId,
			Language:        "en",
			DocumentContent: req.ContentZh,
			DocumentTitle:   req.TitleEn,
		})
		if err != nil {
			return nil, err
		}

	} else {
		//修改
		_, err = l.svcCtx.UserRpcClient.EditHelpDocumentTranslation(l.ctx, &user.EditHelpDocumentTranslationRequest{
			HelpDocumentId:  req.Helpdocumentid,
			Language:        "zh",
			DocumentTitle:   req.TitleZh,
			DocumentContent: req.ContentZh,
		})
		if err != nil {
			return nil, err
		}

		_, err = l.svcCtx.UserRpcClient.EditHelpDocumentTranslation(l.ctx, &user.EditHelpDocumentTranslationRequest{
			HelpDocumentId:  req.Helpdocumentid,
			Language:        "zh",
			DocumentTitle:   req.TitleEn,
			DocumentContent: req.ContentZh,
		})
		if err != nil {
			return nil, err
		}

	}

	return &types.EditDocumentContentReply{
		IsSuccess: true,
	}, nil
}
