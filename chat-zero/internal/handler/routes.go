// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	check "gtihub.com/ryantokmanmokmtm/chat-zero/internal/handler/check"
	user "gtihub.com/ryantokmanmokmtm/chat-zero/internal/handler/user"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: check.PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/signup",
				Handler: user.UsersignupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/signin",
				Handler: user.UsersigninHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/profile",
				Handler: user.GetuserprofileHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
