package db

import (
	"context"
	"os"
	"strings"

	"github.com/taufiktriantono/api-first-monorepo/pkg/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

var Module = fx.Module("database",
	fx.Provide(
		Dialect,
		New,
	),
)

func New(lc fx.Lifecycle, cfg *config.Config, dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
	// Initialize the GORM DB connection
	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		zap.L().Fatal("[DB] ❌ Failed to connect to database", zap.Error(err))
		return nil, err
	}

	if os.Getenv("ENV") != "production" {
		db = db.Debug()
		zap.L().Info("[DB] 🔍 Database is running in DEBUG mode")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.ConnectionPool.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.ConnectionPool.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnectionPool.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.ConnectionPool.ConnMaxIdleTime)

	// db.Callback().Create().Before("gorm:create").Register("audit_before_create", BeforeCreate)
	// db.Callback().Update().Before("gorm:update").Register("audit_before_update", BeforeUpdate)

	zap.L().Info("[DB] ✅ Database connection successfully configured with connection pooling.")

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			zap.L().Info("[DB] Closing connection pool...")
			return sqlDB.Close()
		},
	})

	return db, nil
}

func NewTest() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Otel(db *gorm.DB) error {
	// Register the OpenTelemetry plugin with GORM
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		zap.L().Fatal("❌ Failed to register db telemetry", zap.Error(err))
		return err
	}

	return nil
}

func Metric(db *gorm.DB) error {
	if err := db.Use(prometheus.New(prometheus.Config{
		DBName:          getDBNameFromDialector(db.Dialector), // use `DBName` as metrics label
		RefreshInterval: 15,                                   // Refresh metrics interval (default 15 seconds)
		PushAddr:        "localhost:9090",                     // push metrics if `PushAddr` configured
		StartServer:     true,                                 // start http server to expose metrics
		HTTPServerPort:  8080,                                 // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.Postgres{
				VariableNames: []string{"Threads_running"},
			},
		}, // user defined metrics
	})); err != nil {
		zap.L().Fatal("❌ Failed to register db metrics", zap.Error(err))
		return err
	}
	return nil
}

// Helper function to extract DB name from DSN string
func extractDBNameFromDSN(dsn string) string {
	// Split the DSN into parts (space-separated for PostgreSQL, semicolon-separated for MySQL)
	parts := strings.Fields(dsn) // Fields splits by any whitespace (e.g., spaces)
	for _, part := range parts {
		// Look for the "dbname=" parameter and extract the database name
		if strings.HasPrefix(part, "dbname=") {
			return strings.TrimPrefix(part, "dbname=")
		}
	}
	return "unknown"
}

// Function to get the DB name based on the Dialector type (Postgres/MySQL)
func getDBNameFromDialector(dialector gorm.Dialector) string {
	switch d := dialector.(type) {
	case *postgres.Dialector:
		// For PostgreSQL, extract the DB name from DSN
		return extractDBNameFromDSN(d.Config.DSN)
	case *mysql.Dialector:
		// For MySQL, extract the DB name from DSN
		return extractDBNameFromDSN(d.Config.DSN)
	default:
		return "unknown"
	}
}
