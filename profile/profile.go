package profile

import (
	"github.com/micro/go-micro/v3/auth/noop"
	"github.com/micro/go-micro/v3/broker/http"
	"github.com/micro/go-micro/v3/registry/mdns"
	"github.com/micro/go-micro/v3/runtime/local"
	"github.com/micro/micro/v3/service/logger"

	mStore "github.com/micro/go-micro/v3/config/store"
	memStream "github.com/micro/go-micro/v3/events/stream/memory"
	mFile "github.com/micro/go-micro/v3/store/file"
	mem "github.com/micro/go-micro/v3/store/memory"

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

var (
	BASE_HERF_PATH = "./"
)

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
		//microStore.DefaultStore = file.NewStore()
		microStore.DefaultStore = mem.NewStore()

		//this is not right, currently,file config is under redoing
		microConfig.DefaultConfig, _ = mStore.NewConfig(mFile.NewStore(), "micro-starter")
		// microConfig.DefaultConfig, _ = config.NewConfig(
		// 	config.WithSource(
		// 		mSrcFile.NewSource(
		// 			mSrcFile.WithPath(BASE_HERF_PATH + "config.yaml"),
		// 		),
		// 	),
		// )
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
