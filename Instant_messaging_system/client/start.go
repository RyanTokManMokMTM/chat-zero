package client

import (
	"context"
	"github.com/spf13/cobra"
)

func RunClientCmd(ctx context.Context, version string) *cobra.Command {
	opts := &StartOpt{}

	cmd := cobra.Command{
		Use:   "client",
		Short: "Start Chat Client",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(ctx, opts)
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.addr, "addr", "a", "ws://127.0.0.1:8000", "server addr")
	cmd.PersistentFlags().StringVarP(&opts.name, "user", "u", "", "user name")

	return &cmd
}
