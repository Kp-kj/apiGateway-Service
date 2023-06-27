package curatorial

import (
	"fmt"
	"net/http"

	"gateway/internal/logic/curatorial"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryTaskDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskDetailsInput
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Printf("!2222222222222222222222:%v\n", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := curatorial.NewQueryTaskDetailsLogic(r.Context(), svcCtx)
		resp, err := l.QueryTaskDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
