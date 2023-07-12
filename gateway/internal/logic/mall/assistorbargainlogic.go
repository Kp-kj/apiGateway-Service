package mall

import (
	"context"
	"encoding/json"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/userclient"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
)

type AssistorBargainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssistorBargainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssistorBargainLogic {
	return &AssistorBargainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssistorBargainLogic) AssistorBargain(req *types.AssistorBargainInput) (resp *types.IsSuccessReply, err error) {

	userId := l.ctx.Value("userId")
	var userID int64
	if v, ok := userId.(json.Number); ok {
		userID, _ = v.Int64()
	}
	user, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{
		UserId: strconv.Itoa(int(userID)),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var redisBargainPrice float32

	member, _ := l.svcCtx.Redis.Spop(strconv.FormatInt(req.BargainID, 10))
	if member != "" {
		floatPrice, _ := strconv.ParseFloat(member, 64)
		redisBargainPrice = float32(floatPrice)
	} else {
		redisBargainPrice = 0
	}

	bargain, err := l.svcCtx.BlockClient.AssistorBargain(l.ctx, &blockclient.AssistorBargainRequest{
		UserId:       userID,
		UserName:     user.UserName,
		TwitterName:  user.TwitterName,
		Avatar:       user.UserAvatar,
		BargainPrice: redisBargainPrice,
		BargainId:    req.BargainID,
	})

	if bargain.RedisExist == false {
		split := strings.Split(bargain.UnuserArr, ",")
		for _, amount := range split {
			_, err = l.svcCtx.Redis.Sadd(strconv.FormatInt(req.BargainID, 10), amount)
		}
		l.svcCtx.Redis.Expire(strconv.FormatInt(req.BargainID, 10), 180)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
	} else {
		if member != "" {
			_, err := l.svcCtx.Redis.Srem(strconv.FormatInt(req.BargainID, 10), member)
			if err != nil {
				logx.Error(err)
				return nil, err
			}
			err = l.svcCtx.Redis.Expire(strconv.FormatInt(req.BargainID, 10), 180)
			if err != nil {
				logx.Error(err)
				return nil, err
			}
		} else {
			logx.Error(err)
			return nil, err
		}
	}

	return &types.IsSuccessReply{IsSuccess: true}, err
}
