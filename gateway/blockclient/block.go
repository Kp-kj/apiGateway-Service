// Code generated by goctl. DO NOT EDIT.
// Source: block.proto

package blockclient

import (
	"context"

	"gateway/blockclient/block"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Activity                   = block.Activity
	AdminActivityListRequest   = block.AdminActivityListRequest
	AdminActivityListResponse  = block.AdminActivityListResponse
	AdminGood                  = block.AdminGood
	AdminGoodListRequest       = block.AdminGoodListRequest
	AdminGoodListResponse      = block.AdminGoodListResponse
	AssistorBargainRequest     = block.AssistorBargainRequest
	AssistorBargainResponse    = block.AssistorBargainResponse
	BargainPayRequest          = block.BargainPayRequest
	BargainRecord              = block.BargainRecord
	CreateActivityRequest      = block.CreateActivityRequest
	CreateCryptominerRequest   = block.CreateCryptominerRequest
	CreatePropRequest          = block.CreatePropRequest
	Cryptominer                = block.Cryptominer
	CryptominerBargainRequest  = block.CryptominerBargainRequest
	CryptominerBargainResponse = block.CryptominerBargainResponse
	CryptominerPurchaseRequest = block.CryptominerPurchaseRequest
	GetBargainRecordRequest    = block.GetBargainRecordRequest
	GetBargainRecordResponse   = block.GetBargainRecordResponse
	GetGoodsListRequest        = block.GetGoodsListRequest
	GetGoodsListResponse       = block.GetGoodsListResponse
	GetPurchaseRecordRequest   = block.GetPurchaseRecordRequest
	GetPurchaseRecordResponse  = block.GetPurchaseRecordResponse
	IsSuccessResponse          = block.IsSuccessResponse
	JudgeBargainRequest        = block.JudgeBargainRequest
	JudgeBargainResponse       = block.JudgeBargainResponse
	JudgeGoodsPurchaseRequest  = block.JudgeGoodsPurchaseRequest
	JudgeGoodsPurchaseResponse = block.JudgeGoodsPurchaseResponse
	Prop                       = block.Prop
	PropPurchaseRequest        = block.PropPurchaseRequest
	PurchaseRecord             = block.PurchaseRecord
	Request                    = block.Request
	Response                   = block.Response
	StartActivityRequest       = block.StartActivityRequest
	StartGoodRequest           = block.StartGoodRequest
	SupportUser                = block.SupportUser

	Block interface {
		// 后台接口
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		CreateCryptominer(ctx context.Context, in *CreateCryptominerRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		CreateProp(ctx context.Context, in *CreatePropRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		AdminGoodList(ctx context.Context, in *AdminGoodListRequest, opts ...grpc.CallOption) (*AdminGoodListResponse, error)
		StartGood(ctx context.Context, in *StartGoodRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		AdminActivityList(ctx context.Context, in *AdminActivityListRequest, opts ...grpc.CallOption) (*AdminActivityListResponse, error)
		StartActivity(ctx context.Context, in *StartActivityRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		// 前台接口
		GetGoodsList(ctx context.Context, in *GetGoodsListRequest, opts ...grpc.CallOption) (*GetGoodsListResponse, error)
		GetPurchaseRecord(ctx context.Context, in *GetPurchaseRecordRequest, opts ...grpc.CallOption) (*GetPurchaseRecordResponse, error)
		JudgeBargain(ctx context.Context, in *JudgeBargainRequest, opts ...grpc.CallOption) (*JudgeBargainResponse, error)
		CryptominerFullPurchase(ctx context.Context, in *CryptominerPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		PropPurchase(ctx context.Context, in *PropPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		CryptominerBargainPurchase(ctx context.Context, in *CryptominerBargainRequest, opts ...grpc.CallOption) (*CryptominerBargainResponse, error)
		AssistorBargain(ctx context.Context, in *AssistorBargainRequest, opts ...grpc.CallOption) (*AssistorBargainResponse, error)
		GetBargainRecord(ctx context.Context, in *GetBargainRecordRequest, opts ...grpc.CallOption) (*GetBargainRecordResponse, error)
		BargainPay(ctx context.Context, in *BargainPayRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		// 外部rpc接口
		JudgeGoodsPurchase(ctx context.Context, in *JudgeGoodsPurchaseRequest, opts ...grpc.CallOption) (*JudgeGoodsPurchaseResponse, error)
	}

	defaultBlock struct {
		cli zrpc.Client
	}
)

func NewBlock(cli zrpc.Client) Block {
	return &defaultBlock{
		cli: cli,
	}
}

// 后台接口
func (m *defaultBlock) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultBlock) CreateCryptominer(ctx context.Context, in *CreateCryptominerRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CreateCryptominer(ctx, in, opts...)
}

func (m *defaultBlock) CreateProp(ctx context.Context, in *CreatePropRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CreateProp(ctx, in, opts...)
}

func (m *defaultBlock) AdminGoodList(ctx context.Context, in *AdminGoodListRequest, opts ...grpc.CallOption) (*AdminGoodListResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.AdminGoodList(ctx, in, opts...)
}

func (m *defaultBlock) StartGood(ctx context.Context, in *StartGoodRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.StartGood(ctx, in, opts...)
}

func (m *defaultBlock) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CreateActivity(ctx, in, opts...)
}

func (m *defaultBlock) AdminActivityList(ctx context.Context, in *AdminActivityListRequest, opts ...grpc.CallOption) (*AdminActivityListResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.AdminActivityList(ctx, in, opts...)
}

func (m *defaultBlock) StartActivity(ctx context.Context, in *StartActivityRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.StartActivity(ctx, in, opts...)
}

// 前台接口
func (m *defaultBlock) GetGoodsList(ctx context.Context, in *GetGoodsListRequest, opts ...grpc.CallOption) (*GetGoodsListResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetGoodsList(ctx, in, opts...)
}

func (m *defaultBlock) GetPurchaseRecord(ctx context.Context, in *GetPurchaseRecordRequest, opts ...grpc.CallOption) (*GetPurchaseRecordResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetPurchaseRecord(ctx, in, opts...)
}

func (m *defaultBlock) JudgeBargain(ctx context.Context, in *JudgeBargainRequest, opts ...grpc.CallOption) (*JudgeBargainResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.JudgeBargain(ctx, in, opts...)
}

func (m *defaultBlock) CryptominerFullPurchase(ctx context.Context, in *CryptominerPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CryptominerFullPurchase(ctx, in, opts...)
}

func (m *defaultBlock) PropPurchase(ctx context.Context, in *PropPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.PropPurchase(ctx, in, opts...)
}

func (m *defaultBlock) CryptominerBargainPurchase(ctx context.Context, in *CryptominerBargainRequest, opts ...grpc.CallOption) (*CryptominerBargainResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CryptominerBargainPurchase(ctx, in, opts...)
}

func (m *defaultBlock) AssistorBargain(ctx context.Context, in *AssistorBargainRequest, opts ...grpc.CallOption) (*AssistorBargainResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.AssistorBargain(ctx, in, opts...)
}

func (m *defaultBlock) GetBargainRecord(ctx context.Context, in *GetBargainRecordRequest, opts ...grpc.CallOption) (*GetBargainRecordResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetBargainRecord(ctx, in, opts...)
}

func (m *defaultBlock) BargainPay(ctx context.Context, in *BargainPayRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.BargainPay(ctx, in, opts...)
}

// 外部rpc接口
func (m *defaultBlock) JudgeGoodsPurchase(ctx context.Context, in *JudgeGoodsPurchaseRequest, opts ...grpc.CallOption) (*JudgeGoodsPurchaseResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.JudgeGoodsPurchase(ctx, in, opts...)
}
