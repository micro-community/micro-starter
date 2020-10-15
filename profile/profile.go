package profile

import (
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/auth/noop"
	"github.com/micro/micro/v3/service/broker/http"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/config/env"
	"github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/events/stream/memory"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/registry/mdns"
	"github.com/micro/micro/v3/service/runtime"
	"github.com/micro/micro/v3/service/runtime/local"
	"github.com/micro/micro/v3/service/store"
	"github.com/micro/micro/v3/service/store/file"

	"github.com/micro/micro/v3/profile"
	"github.com/urfave/cli/v2"
)

func init() {
	profile.Register("dev", Dev)
}

/*
	config.WithSource(
			mSrcFile.NewSource(
				mSrcFile.WithPath(BASE_HERF_PATH + "config.yaml"),
			),
*/

// Dev profile to run develop env
var Dev = &profile.Profile{
	Name: "dev",
	Setup: func(ctx *cli.Context) error {
		auth.DefaultAuth = noop.NewAuth()
		runtime.DefaultRuntime = local.NewRuntime()
		store.DefaultStore = file.NewStore()
		//store.DefaultStore = mem.NewStore()
		config.DefaultConfig, _ = env.NewConfig()
		profile.SetupBroker(http.NewBroker())
		profile.SetupRegistry(mdns.NewRegistry())
		//	profile.SetupJWTRules()
		var err error
		events.DefaultStream, err = memory.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream for dev: %v", err)
		}

		return nil
	},
}
