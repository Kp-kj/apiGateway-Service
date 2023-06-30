package ping

import (
	"net/http"

	"gateway/internal/logic/ping"
	"gateway/internal/svc"

	xhttp "github.com/zeromicro/x/http"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ping.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			w.WriteHeader(http.StatusOK)
			xhttp.JsonBaseResponseCtx(r.Context(), w, "pong")
		}
	}
}
