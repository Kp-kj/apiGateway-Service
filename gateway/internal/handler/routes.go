// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	admin "gateway/internal/handler/admin"
	adminmall "gateway/internal/handler/adminmall"
	curatorial "gateway/internal/handler/curatorial"
	everyday "gateway/internal/handler/everyday"
	mall "gateway/internal/handler/mall"
	ping "gateway/internal/handler/ping"
	user "gateway/internal/handler/user"
	"gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: ping.PingHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/userLogin",
				Handler: user.UserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/getHelpCategoryList",
				Handler: user.GetHelpCategoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/getHelpDocumentTitleList",
				Handler: user.GetHelpDocumentTitleListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/getContentByHelpDocumentId",
				Handler: user.GetContentByHelpDocumentIdHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.BlackMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/userLogout",
					Handler: user.UserLogoutHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getUserInfo",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getMessageList",
					Handler: user.GetMessageListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/admin/adminLogin",
				Handler: admin.AdminLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/admin/createAdminUser",
				Handler: admin.CreateAdminUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/admin/adminLogout",
				Handler: admin.AdminLogoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/admin/getCurrentOnlinePerson",
				Handler: admin.GetCurrentOnlinePersonHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/admin/getRegisteredPopulation",
				Handler: admin.GetRegisteredPopulationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/admin/getUserInfo",
				Handler: admin.AdminUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/admin/getUserListByCondition",
				Handler: admin.GetUserListByConditionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/admin/batchUpdateUserBlackStatus",
				Handler: admin.BatchUpdateUserBlackStatusHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.AdminAuth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.BlackMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/curatorial/create",
					Handler: curatorial.CreateCuratorialTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/curatorial/list",
					Handler: curatorial.QueryTaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/curatorial/details",
					Handler: curatorial.QueryTaskDetailsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/curatorial/user/list",
					Handler: curatorial.QueryUserLaunchTaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/curatorial/label/create",
					Handler: curatorial.CreateLabelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/curatorial/label/delete",
					Handler: curatorial.DeleteLabelHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/curatorial/label/list",
					Handler: curatorial.GetLabelListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/curatorial/verify",
					Handler: curatorial.PerformTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/curatorial/voluntarily",
					Handler: curatorial.VoluntarilyTaskScheduleHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.BlackMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/everyday/treasure/amend",
					Handler: everyday.AmendTreasureTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/everyday/treasure/change",
					Handler: everyday.ChangeTreasureTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/everyday/treasure/list",
					Handler: everyday.QueryTreasureTaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/everyday/subtask/list",
					Handler: everyday.QuerySubtaskStyleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/everyday/subtask/amend",
					Handler: everyday.AmendAssociatedSubtaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/everyday/subtask/delete",
					Handler: everyday.DeleteAssociatedSubtaskHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/everyday/subtask/treasureId",
					Handler: everyday.QueryAssociatedSubtaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/everyday/chest/amemd",
					Handler: everyday.AmendChestCollectionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/everyday/chest/schedule",
					Handler: everyday.QueryChestCollectionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/everyday/subtask/power",
					Handler: everyday.CreateUserPowerTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/everyday/subtask/create",
					Handler: everyday.CreateSubtaskStyleHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/adminmall/create/cryptominer",
				Handler: adminmall.CreateCryptominerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/adminmall/create/prop",
				Handler: adminmall.CreatePropHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/mall/goodlist",
				Handler: mall.GetGoodsListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/mall/judgebargain",
				Handler: mall.JudgeBargainHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mall/purchase/full",
				Handler: mall.CryptominerFullPurchaseHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mall/purchase/bargain",
				Handler: mall.CryptominerBargainPurchaseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/mall/activity/rule",
				Handler: mall.GetBargainRuleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/mall/activity/cryptominer",
				Handler: mall.GetBargainCryptominerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/mall/activity/progress",
				Handler: mall.GetBargainProgressHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/mall/activity/bargainlist",
				Handler: mall.GetBargainRecordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)
}
