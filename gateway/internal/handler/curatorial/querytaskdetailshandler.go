package curatorial

import (
	"net/http"

	"gateway/internal/logic/curatorial"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func QueryTaskDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskDetailsInput
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := curatorial.NewQueryTaskDetailsLogic(r.Context(), svcCtx)
		resp, err := l.QueryTaskDetails(&req)
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
