package user

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"gateway/internal/logic/user"
	"gateway/internal/svc"
)

func UserLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserLogoutLogic(r.Context(), svcCtx)
		resp, err := l.UserLogout()
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
