package admin

import (
	"net/http"

	"gateway/。/internal/logic/admin"
	"gateway/。/internal/svc"
	"gateway/。/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BatchUpdateUserBlackStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchUpdateUserBlackStatus
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewBatchUpdateUserBlackStatusLogic(r.Context(), svcCtx)
		resp, err := l.BatchUpdateUserBlackStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
