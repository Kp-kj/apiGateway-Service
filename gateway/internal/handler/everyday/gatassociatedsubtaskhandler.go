package everyday

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"gateway/internal/logic/everyday"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GatAssociatedSubtaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskIDInquireInput
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := everyday.NewGatAssociatedSubtaskLogic(r.Context(), svcCtx)
		resp, err := l.GatAssociatedSubtask(&req)
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
