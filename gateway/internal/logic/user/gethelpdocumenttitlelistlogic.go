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

// 获取帮助文档标题列表
func (l *GetHelpDocumentTitleListLogic) GetHelpDocumentTitleList(req *types.GetHelpDocumentTitleList) (resp *types.GetHelpDocumentTitleListReply, err error) {
	////获取该分类下上架的帮助文档列表
	helpCategoryId, err := strconv.ParseInt(req.HelpCategoryId, 10, 64)
	helpResp, err := l.svcCtx.UserRpcClient.GetHelpDocuments(l.ctx, &userclient.GetHelpDocumentsRequest{
		HelpCategoryId: helpCategoryId,
		DocumentStatus: 1, //上架
	})
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	DocumentsList := make([]types.HelpDocumentTitleList, len(helpResp.HelpDocuments))
	for i, documents := range helpResp.HelpDocuments {
		// 查找翻译结果
		translationResp, err := l.svcCtx.UserRpcClient.GetHelpDocumentTranslations(l.ctx, &userclient.GetHelpDocumentTranslationsRequest{
			HelpDocumentId: documents.HelpDocumentId,
			Language:       req.LanguageCode,
		})
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}
		// 获取翻译结果
		translation := ""
		if translationResp != nil {
			translation = translationResp.DocumentTitle
		}

		// 创建分类列表项
		DocumentsList[i] = types.HelpDocumentTitleList{
			HelpDocumentId: documents.HelpDocumentId,
			DocumentTitle:  translation,
			LanguageCode:   req.LanguageCode,
		}
	}

	return &types.GetHelpDocumentTitleListReply{HelpDocumentTitleList: DocumentsList}, nil
}
