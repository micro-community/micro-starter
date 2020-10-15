package profile

import (
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/auth/noop"
	"github.com/micro/micro/v3/service/broker/http"

	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/config/env"

	//	mConfStore "github.com/micro/micro/v3/service/config/store"

	"github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/events/stream/memory"
	"github.com/micro/micro/v3/service/logger"

	//	"github.com/micro/micro/v3/service/registry/mdns"
	mregistry "github.com/micro/micro/v3/service/registry/memory"

	"github.com/micro/micro/v3/service/runtime"
	"github.com/micro/micro/v3/service/runtime/local"

	"github.com/micro/micro/v3/service/store"
	//	fstore "github.com/micro/micro/v3/service/store/file"
	mstore "github.com/micro/micro/v3/service/store/memory"

	"github.com/micro/micro/v3/profile"
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
		//	config.DefaultConfig, _ = mConfStore.NewConfig(store.DefaultStore, "")
		profile.SetupBroker(http.NewBroker())
		profile.SetupRegistry(mregistry.NewRegistry())
		//	profile.SetupJWTRules()
		var err error
		events.DefaultStream, err = memory.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream for dev: %v", err)
		}

		// store.DefaultBlobStore, err = fstore.NewBlobStore()
		// if err != nil {
		// 	logger.Fatalf("Error configuring file blob store: %v", err)
		// }

		return nil
	},
}
