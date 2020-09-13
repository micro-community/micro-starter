// Package k8s profile is for specific profiles
// @todo this package is the definition of cruft and
// should be rewritten in a more elegant way
package k8s

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v3/auth/jwt"
	"github.com/micro/go-micro/v3/router/static"
	"github.com/micro/go-micro/v3/runtime/kubernetes"

	//	microMetrics "github.com/micro/micro/v3/service/metrics"
	mProfile "github.com/micro/micro/v3/profile"
	microAuth "github.com/micro/micro/v3/service/auth"
	microRouter "github.com/micro/micro/v3/service/router"
	microRuntime "github.com/micro/micro/v3/service/runtime"
)

func init() {
	mProfile.Register("Kubernetes", Kubernetes)
}

// Kubernetes profile to run on kubernetes
var Kubernetes = &mProfile.Profile{
	Name: "kubernetes",
	Setup: func(ctx *cli.Context) error {
		// TODO: implement
		// using a static router so queries are routed based on service name
		microRouter.DefaultRouter = static.NewRouter()
		// Using the kubernetes runtime
		microRuntime.DefaultRuntime = kubernetes.NewRuntime()
		// registry kubernetes
		// config configmap
		// store ...
		microAuth.DefaultAuth = jwt.NewAuth()
		//	setupJWTRules()
		// Set up a default metrics reporter (being careful not to clash with any that have already been set):
		// if !microMetrics.IsSet() {
		// 	prometheusReporter, err := metricsPrometheus.New()
		// 	if err != nil {
		// 		return err
		// 	}
		// 	microMetrics.SetDefaultMetricsReporter(prometheusReporter)
		// }

		return nil
	},
}
