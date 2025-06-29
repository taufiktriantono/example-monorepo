package approvals

import (
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/interfaces/http"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/repository"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/service"
	"github.com/taufiktriantono/api-first-monorepo/pkg/config"
	"github.com/taufiktriantono/api-first-monorepo/pkg/db"
	"github.com/taufiktriantono/api-first-monorepo/pkg/logger"
	"github.com/taufiktriantono/api-first-monorepo/pkg/pprof"
	"github.com/taufiktriantono/api-first-monorepo/pkg/server"
	"go.uber.org/fx"
)

var Service = fx.Module("approval.svc",
	config.Module,
	logger.Module,
	db.Module,
	repository.Module,
	service.Module,
	http.Module,
	pprof.Module,
	server.Module,
)
