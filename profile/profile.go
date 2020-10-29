package profile

import (
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/auth/noop"

	//"github.com/micro/micro/v3/service/broker/http"
	mbroker "github.com/micro/micro/v3/service/broker/memory"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/config/env"
	"github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/events/stream/memory"
	"github.com/micro/micro/v3/service/logger"

	"github.com/micro/micro/v3/service/runtime"
	"github.com/micro/micro/v3/service/runtime/local"

	"github.com/micro/micro/v3/profile"

	//mregistry "github.com/micro/micro/v3/service/registry/memory"
	mregistry "github.com/micro/micro/v3/service/registry/mdns"
	"github.com/micro/micro/v3/service/store"
	mstore "github.com/micro/micro/v3/service/store/memory"
	"github.com/urfave/cli/v2"
)

func init() {
	_ = profile.Register("dev", Dev)
}

// Dev profile to run develop env
var Dev = &profile.Profile{
	Name: "dev",
	Setup: func(ctx *cli.Context) error {
		auth.DefaultAuth = noop.NewAuth()
		runtime.DefaultRuntime = local.NewRuntime()
		//store.DefaultStore = fstore.NewStore()
		store.DefaultStore = mstore.NewStore()
		config.DefaultConfig, _ = env.NewConfig()
		//config.DefaultConfig, _ = storeConfig.NewConfig(store.DefaultStore, "")

		//	profile.SetupJWTRules()
		var err error
		events.DefaultStream, err = memory.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream for dev: %v", err)
		}

		profile.SetupBroker(mbroker.NewBroker())
		profile.SetupRegistry(mregistry.NewRegistry())
		// store.DefaultBlobStore, err = fstore.NewBlobStore()
		// if err != nil {
		// 	logger.Fatalf("Error configuring file blob store: %v", err)
		// }

		return nil
	},
}
