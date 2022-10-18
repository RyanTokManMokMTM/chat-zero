package room

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/logic/room"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/types"
)

func LeaveRoomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LeaveRoomReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := room.NewLeaveRoomLogic(r.Context(), svcCtx)
		resp, err := l.LeaveRoom(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
