package admin

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"gateway/internal/logic/admin"
	"gateway/internal/svc"
)

func AdminLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewAdminLogoutLogic(r.Context(), svcCtx)
		resp, err := l.AdminLogout()
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			w.WriteHeader(http.StatusOK)
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
