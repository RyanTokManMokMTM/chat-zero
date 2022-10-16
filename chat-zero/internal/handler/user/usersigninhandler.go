package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/logic/user"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"
)

func UsersigninHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUsersigninLogic(r.Context(), svcCtx)
		resp, err := l.Usersignin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
