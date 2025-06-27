package logging

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("zap", fx.Provide(New))

func New(fs []zap.Field) *zap.Logger {

	log := zap.Must(zap.NewDevelopment())
	if os.Getenv("ENV") == "production" {
		log = zap.Must(zap.NewProduction())
	}

	log = log.With(fs...)

	zap.ReplaceGlobals(log)

	return log
}
