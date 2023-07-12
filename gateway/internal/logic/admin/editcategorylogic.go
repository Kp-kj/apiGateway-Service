package admin

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditCategoryLogic {
	return &EditCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 新建和编辑分类
func (l *EditCategoryLogic) EditCategory(req *types.EditCategory) (resp *types.EditCategoryReply, err error) {
	if req.CategoryId == 0 {
		// 新建分类
		caregoryId, err := l.svcCtx.UserRpcClient.CreateHelpCategory(l.ctx, &userclient.CreateHelpCategoryRequest{
			CreatedAt:      0,
			CategoryStatus: 1,
		}) // 创建分类 返回分类id
		if err != nil {
			return nil, err
		}
		// 那分类id取创建分类的翻译
		_, err = l.svcCtx.UserRpcClient.CreateHelpCategoryTranslation(l.ctx, &userclient.CreateHelpCategoryTranslationRequest{
			HelpCategoryId: caregoryId.HelpCategoryId,
			CategoryName:   req.CategoryNameZh,
			Language:       "zh",
		})
		if err != nil {
			return nil, err
		}

		_, err = l.svcCtx.UserRpcClient.CreateHelpCategoryTranslation(l.ctx, &userclient.CreateHelpCategoryTranslationRequest{
			HelpCategoryId: caregoryId.GetHelpCategoryId(),
			CategoryName:   req.CategoryNameZh,
			Language:       "en",
		})
		if err != nil {
			return nil, err
		}
	} else {
		// 编辑分类翻译
		//中文
		_, err = l.svcCtx.UserRpcClient.EditHelpCategoryTranslation(l.ctx, &userclient.EditHelpCategoryTranslationRequest{
			HelpCategoryId: req.CategoryId,
			CategoryName:   req.CategoryNameZh,
			Language:       "zh",
		})
		if err != nil {
			return nil, err
		}
		//英文
		_, err = l.svcCtx.UserRpcClient.EditHelpCategoryTranslation(l.ctx, &userclient.EditHelpCategoryTranslationRequest{
			HelpCategoryId: req.CategoryId,
			CategoryName:   req.CategoryNameEn,
			Language:       "en",
		})
		if err != nil {
			return nil, err
		}
	}
	return &types.EditCategoryReply{
		IsSuccess: true,
	}, nil
}
