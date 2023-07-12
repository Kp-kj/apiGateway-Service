package admin

import (
	"net/http"

	"gateway/internal/logic/admin"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func AmendAssociatedSubtaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssociatedSubtaskSrt
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewAmendAssociatedSubtaskLogic(r.Context(), svcCtx)
		resp, err := l.AmendAssociatedSubtask(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			w.WriteHeader(http.StatusOK)
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
