package profile

import (
	"github.com/micro/go-micro/v3/auth/noop"
	"github.com/micro/go-micro/v3/broker/http"
	"github.com/micro/go-micro/v3/registry/mdns"
	"github.com/micro/go-micro/v3/runtime/local"
	"github.com/micro/go-micro/v3/store/file"
	"github.com/micro/micro/v3/service/logger"

	mStore "github.com/micro/go-micro/v3/config/store"
	//mEnv "github.com/micro/go-micro/v3/config/env"
	memStream "github.com/micro/go-micro/v3/events/stream/memory"

	microProfile "github.com/micro/micro/v3/profile"
	microAuth "github.com/micro/micro/v3/service/auth"
	microConfig "github.com/micro/micro/v3/service/config"
	microEvents "github.com/micro/micro/v3/service/events"
	microRuntime "github.com/micro/micro/v3/service/runtime"
	microStore "github.com/micro/micro/v3/service/store"

	"github.com/urfave/cli/v2"
)

func init() {
	microProfile.Register("dev", Dev)
}

/*
	config.WithSource(
			mSrcFile.NewSource(
				mSrcFile.WithPath(BASE_HERF_PATH + "config.yaml"),
			),
*/

// Dev profile to run develop env
var Dev = &microProfile.Profile{
	Name: "dev",
	Setup: func(ctx *cli.Context) error {
		microAuth.DefaultAuth = noop.NewAuth()
		microRuntime.DefaultRuntime = local.NewRuntime()
		microStore.DefaultStore = file.NewStore()
		//microStore.DefaultStore = mem.NewStore()
		microConfig.DefaultConfig, _ = mStore.NewConfig(microStore.DefaultStore, "")
		microProfile.SetupBroker(http.NewBroker())
		microProfile.SetupRegistry(mdns.NewRegistry())
		//	microProfile.SetupJWTRules()
		var err error
		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream for dev: %v", err)
		}

		return nil
	},
}
