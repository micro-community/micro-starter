package main

import (
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/pubsub"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/cmd"
	"github.com/micro/micro/v3/service"
	"github.com/urfave/cli/v2"

	//load profile
	"github.com/micro-community/auth/profile"
)

func main() {

	srv := service.New(
		service.Name("micro-v3-starter"),
		service.Version("latest"),
	)

	// add customer Flags
	cmdFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "configpath",
			EnvVars: []string{"MICRO_STARTER_CONFIG_PATH"},
			Usage:   "config path of current app",
			Value:   "./",
		},
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "run in debug mode",
			EnvVars: []string{"MICRO_STARTER_DEBUG_MODE"},
			Value:   true,
		},
	}

	cmdFlags = append(cmd.DefaultCmd.App().Flags, cmdFlags...)
	cmdOption := cmd.Flags(cmdFlags...)

	cmd.DefaultCmd.Init(cmdOption)

	srv.Init()

	config.LoadConfigWithDefault(nil)

	profile.BuildingStartupService(srv, config.Default)
	//handle pub/sub message
	pubsub.RegisterSubscription(srv, config.Default.Pubsub)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
