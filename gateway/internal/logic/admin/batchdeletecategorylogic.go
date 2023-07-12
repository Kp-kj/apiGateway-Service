package admin

import (
	"context"
	"gateway/userclient/user"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteCategoryLogic {
	return &BatchDeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BatchDeleteCategory 批量删除分类
func (l *BatchDeleteCategoryLogic) BatchDeleteCategory(req *types.BatchDeleteCategory) (resp *types.BatchDeleteCategoryReply, err error) {
	//批量删除接口
	//遍历所有的分类id
	for _, v := range req.CategoryIds {
		//调用删除接口
		_, err = l.svcCtx.UserRpcClient.DeleteHelpCategory(l.ctx, &user.DeleteHelpCategoryRequest{HelpCategoryId: v})
		if err != nil {
			return nil, err
		}
		//删除分类的翻译
		_, err = l.svcCtx.UserRpcClient.DeleteHelpCategoryTranslation(l.ctx, &user.DeleteHelpCategoryTranslationRequest{HelpCategoryId: v, Language: "zh"})
		if err != nil {
			return nil, err
		}

		_, err = l.svcCtx.UserRpcClient.DeleteHelpCategoryTranslation(l.ctx, &user.DeleteHelpCategoryTranslationRequest{HelpCategoryId: v, Language: "en"})
		if err != nil {
			return nil, err
		}
	}
	return &types.BatchDeleteCategoryReply{
		IsSuccess: true,
	}, nil
}
