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

// GetContentByHelpDocumentId 获取帮助文档内容
func (l *GetContentByHelpDocumentIdLogic) GetContentByHelpDocumentId(req *types.GetContentByHelpDocumentId) (resp *types.GetContentByHelpDocumentIdReply, err error) {
	// todo: add your logic here and delete this line
	helpResp, err := l.svcCtx.UserRpcClient.GetHelpCategories(l.ctx, &userclient.GetHelpCategoriesRequest{CategoryStatus: 0})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	fmt.Println(helpResp.HelpCategories)

	categoryList := make([]types.ContentByHelpDocumentIdReply, 0)

	return &types.GetContentByHelpDocumentIdReply{ContextByHelpDocumentId: categoryList}, nil
}
