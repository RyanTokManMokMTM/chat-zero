package user

import (
	"net/http"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/logic/user"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WxMiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewWxMiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.WxMiniAuth(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
