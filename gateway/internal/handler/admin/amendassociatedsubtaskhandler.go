package admin

import (
	"fmt"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"gateway/internal/logic/admin"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AmendAssociatedSubtaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssociatedSubtaskSrt
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Printf("11111:%v\n", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewAmendAssociatedSubtaskLogic(r.Context(), svcCtx)
		resp, err := l.AmendAssociatedSubtask(&req)
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

// 	if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
//		} else {
//			w.WriteHeader(http.StatusOK)
//			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
//		}
