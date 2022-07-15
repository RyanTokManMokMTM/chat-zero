package user

import (
	"net/http"

	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/logic/user"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/svc"
	"github.com/ryantokmanmokmtm/house-booking-service/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewSignupLogic(r.Context(), svcCtx)
		resp, err := l.Signup(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
