package http

import (
	"github.com/taufiktriantono/api-first-monorepo/pkg/health"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	health.Module,
	fx.Invoke(ProvideRouter),
)
