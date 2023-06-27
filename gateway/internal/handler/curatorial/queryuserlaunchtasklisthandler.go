package curatorial

import (
	"fmt"
	"net/http"

	"gateway/internal/logic/curatorial"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryUserLaunchTaskListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLaunchTaskListInput
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Printf("1111111111111111111111:%v\n", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Printf("1111111111111111111111\n")
		l := curatorial.NewQueryUserLaunchTaskListLogic(r.Context(), svcCtx)
		resp, err := l.QueryUserLaunchTaskList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}