package user

import (
	"net/http"

	"gateway/。/internal/logic/user"
	"gateway/。/internal/svc"
	"gateway/。/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetContentByHelpDocumentIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetContentByHelpDocumentId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetContentByHelpDocumentIdLogic(r.Context(), svcCtx)
		resp, err := l.GetContentByHelpDocumentId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
