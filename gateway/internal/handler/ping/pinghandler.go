package ping

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"gateway/internal/logic/ping"
	"gateway/internal/svc"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ping.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			w.WriteHeader(http.StatusOK)
			xhttp.JsonBaseResponseCtx(r.Context(), w, "pong")
		}
	}
}
