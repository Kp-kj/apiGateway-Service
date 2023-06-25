package admin

import (
	"net/http"

	"gateway/internal/logic/admin"
	"gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRegisteredPopulationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewGetRegisteredPopulationLogic(r.Context(), svcCtx)
		resp, err := l.GetRegisteredPopulation()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
