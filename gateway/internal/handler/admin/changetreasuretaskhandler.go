package admin

import (
	"net/http"

	"gateway/internal/logic/admin"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func ChangeTreasureTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TreasureTaskInput
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewChangeTreasureTaskLogic(r.Context(), svcCtx)
		resp, err := l.ChangeTreasureTask(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			w.WriteHeader(http.StatusOK)
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
