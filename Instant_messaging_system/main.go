package main

import (
	"context"
	"flag"
	"github.com/ryantokmanmokmtm/Instant_messaging_system/client"
	"github.com/ryantokmanmokmtm/Instant_messaging_system/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	flag.Parsed()

	r := &cobra.Command{
		Use:     "chat",
		Version: "v1",
		Short:   "chat demo",
	}

	//Server Command
	r.AddCommand(server.NewServerCmd(context.Background(), "v1"))
	r.AddCommand(client.RunClientCmd(context.Background(), "v1"))
	//execute
	if err := r.Execute(); err != nil {
		logrus.WithError(err).Fatal("Run Command failed")
	}
}
