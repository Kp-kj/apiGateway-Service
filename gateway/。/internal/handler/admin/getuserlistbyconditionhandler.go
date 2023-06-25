package admin

import (
	"net/http"

	"gateway/。/internal/logic/admin"
	"gateway/。/internal/svc"
	"gateway/。/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserListByConditionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserListByCondition
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewGetUserListByConditionLogic(r.Context(), svcCtx)
		resp, err := l.GetUserListByCondition(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
