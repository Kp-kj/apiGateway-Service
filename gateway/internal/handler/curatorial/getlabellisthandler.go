package curatorial

import (
	"fmt"
	"net/http"

	"gateway/internal/logic/curatorial"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetLabelListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIDInquireInput
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Printf("222222222222222\n")
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Printf("111111111111111111111\n")
		l := curatorial.NewGetLabelListLogic(r.Context(), svcCtx)
		resp, err := l.GetLabelList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}