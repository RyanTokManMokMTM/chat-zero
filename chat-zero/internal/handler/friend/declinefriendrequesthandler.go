package friend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/logic/friend"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"
)

func DeclineFriendRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeclineFriendNotificationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewDeclineFriendRequestLogic(r.Context(), svcCtx)
		resp, err := l.DeclineFriendRequest(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
