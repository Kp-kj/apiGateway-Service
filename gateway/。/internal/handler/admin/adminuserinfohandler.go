package admin

import (
	"net/http"

	"gateway/。/internal/logic/admin"
	"gateway/。/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewAdminUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.AdminUserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
