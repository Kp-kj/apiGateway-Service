package admin

import (
	"context"
	"fmt"
	"gateway/userclient/user"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DocumentContentlistStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocumentContentlistStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DocumentContentlistStatusLogic {
	return &DocumentContentlistStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DocumentContentlistStatus 文档内容列表启用状态
func (l *DocumentContentlistStatusLogic) DocumentContentlistStatus(req *types.DocumentContentlistStatus) (resp *types.DocumentContentlistStatusReply, err error) {

	fmt.Println(req.Helpdocumentids, req.Status)
	for _, v := range req.Helpdocumentids {
		_, err = l.svcCtx.UserRpcClient.EditHelpDocument(l.ctx, &user.EditHelpDocumentRequest{HelpDocumentId: v, DocumentStatus: req.Status})
		if err != nil {
			return nil, err
		}
	}

	return &types.DocumentContentlistStatusReply{IsSuccess: true}, nil
}
