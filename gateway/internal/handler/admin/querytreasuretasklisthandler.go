package admin

import (
	"fmt"
	"net/http"

	"gateway/internal/logic/admin"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func QueryTreasureTaskListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TreasureTaskListInput
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Printf("zzzzzzzzzz:%v\n", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewQueryTreasureTaskListLogic(r.Context(), svcCtx)
		resp, err := l.QueryTreasureTaskList(&req)
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
