package profile

import (
	"github.com/micro/go-micro/v3/auth/noop"
	"github.com/micro/go-micro/v3/broker/http"
	"github.com/micro/go-micro/v3/config"

	"github.com/micro/go-micro/v3/registry/mdns"
	"github.com/micro/go-micro/v3/runtime/local"

	//"github.com/micro/go-micro/v3/store/file"
	"github.com/micro/micro/v3/service/logger"

	mSrcFile "github.com/micro/go-micro/v3/config/source/file"
	memStream "github.com/micro/go-micro/v3/events/stream/memory"
	mem "github.com/micro/go-micro/v3/store/memory"

	mProfile "github.com/micro/micro/v3/profile"
	microAuth "github.com/micro/micro/v3/service/auth"
	microConfig "github.com/micro/micro/v3/service/config"
	microEvents "github.com/micro/micro/v3/service/events"
	microRuntime "github.com/micro/micro/v3/service/runtime"
	microStore "github.com/micro/micro/v3/service/store"

	"github.com/urfave/cli/v2"
)

func init() {
	mProfile.Register("dev", Dev)
}

var (
	BASE_HERF_PATH = "./"
)

// Dev profile to run develop env
var Dev = &mProfile.Profile{
	Name: "dev",
	Setup: func(ctx *cli.Context) error {
		microAuth.DefaultAuth = noop.NewAuth()
		microRuntime.DefaultRuntime = local.NewRuntime()
		//microStore.DefaultStore = file.NewStore()
		microStore.DefaultStore = mem.NewStore()
		microConfig.DefaultConfig, _ = config.NewConfig(
			config.WithSource(
				mSrcFile.NewSource(
					mSrcFile.WithPath(BASE_HERF_PATH + "config.yaml"),
				),
			),
		)
		mProfile.SetupBroker(http.NewBroker())
		mProfile.SetupRegistry(mdns.NewRegistry())
		//	mProfile.SetupJWTRules()
		var err error
		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream for dev: %v", err)
		}

		return nil
	},
}
