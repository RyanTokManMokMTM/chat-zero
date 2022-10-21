package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/config"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/handler"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/logic/serverWs"
	"gtihub.com/ryantokmanmokmtm/chat-zero/internal/svc"
	"net/http"
)

var configFile = flag.String("f", "etc/chat-zero.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/ws",
		Handler: serverWs.NewServerWS(ctx),
	}, rest.WithJwt(ctx.Config.Auth.AccessSecret))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
