package profiling

import (
	"context"

	"github.com/grafana/pyroscope-go"
	"github.com/taufiktriantono/api-first-monorepo/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("profiling", fx.Invoke(ProvideProfiling))

func ProvideProfiling(lc fx.Lifecycle, c *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			_, err := pyroscope.Start(pyroscope.Config{
				ApplicationName: c.AppName,
				ServerAddress:   c.Pyroscope.Addr,
				ProfileTypes: []pyroscope.ProfileType{
					pyroscope.ProfileCPU,
					pyroscope.ProfileAllocObjects,
					pyroscope.ProfileAllocSpace,
					pyroscope.ProfileInuseObjects,
					pyroscope.ProfileInuseSpace,
					pyroscope.ProfileGoroutines,
				},
			})
			if err != nil {
				zap.L().Error("failed to start pyroscope", zap.Error(err))
				return err
			}

			zap.L().Info("pyroscope started", zap.String("app", c.AppName))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Pyroscope-go doesn't expose Stop(), but if future versions support it,
			// handle graceful shutdown here.
			zap.L().Info("Shutting down Pyroscope (noop for now)")
			return nil
		},
	})
}
