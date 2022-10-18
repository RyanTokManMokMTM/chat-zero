package friend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/logic/friend"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"
)

func AddFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddFriendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewAddFriendLogic(r.Context(), svcCtx)
		resp, err := l.AddFriend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
