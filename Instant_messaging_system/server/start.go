package server

import (
	"context"
	"github.com/spf13/cobra"
)

type ServerStartOptions struct {
	id     string
	listen string
}

func RunServer(ctx context.Context, opt *ServerStartOptions) error {
	server := NewServer(opt.id, opt.listen)
	defer server.Shutdown()
	return server.Start()
}

func NewServerCmd(ctx context.Context, version string) *cobra.Command {
	opts := &ServerStartOptions{}

	cmd := &cobra.Command{
		Use:     "chat",
		Short:   "Start a chat server",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServer(ctx, opts)
		},
	}
	//flags
	cmd.PersistentFlags().StringVarP(&opts.id, "serverId", "i", "demo", "serverId")
	cmd.PersistentFlags().StringVarP(&opts.listen, "listen", "l", ":8000", "server listen address")
	return cmd
}
