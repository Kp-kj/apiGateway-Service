package curatorial

import (
	"net/http"

	"gateway/internal/logic/curatorial"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VoluntarilyTaskScheduleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VoluntarilyTaskScheduleInput
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := curatorial.NewVoluntarilyTaskScheduleLogic(r.Context(), svcCtx)
		resp, err := l.VoluntarilyTaskSchedule(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}