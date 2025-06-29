package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("health", fx.Provide(ProvideHealth))

type Dependency struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Health struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Deps    []Dependency `json:"deps"`
}

type HealthService interface {
	Liveness(c *gin.Context)
	Readiness(c *gin.Context)
}

type health struct {
	db *gorm.DB
}

type Params struct {
	fx.In
	DB *gorm.DB
}

func ProvideHealth(p Params) HealthService {
	return &health{
		db: p.DB,
	}
}

func (h *health) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, &Health{
		Status:  "healthly",
		Message: "OK",
	})
}

func (h *health) Readiness(c *gin.Context) {
	this := &Health{
		Status:  "healthly",
		Message: "OK",
	}

	deps := make([]Dependency, 0)
	if h.db != nil {
		dep := Dependency{
			Name:    h.db.Name(),
			Status:  "healthly",
			Message: "OK",
		}

		sql, err := h.db.DB()
		if err != nil {
			dep.Status = "unhealthly"
			dep.Message = err.Error()
		}

		if err := sql.Ping(); err != nil {
			dep.Status = "unhealthly"
			dep.Message = err.Error()
		}

		deps = append(deps, dep)
	}

	// if h.redis != nil {
	// 	dep := Dependency{
	// 		Name:    h.db.Name(),
	// 		Status:  "healthly",
	// 		Message: "OK",
	// 	}

	// 	if err := h.redis.Ping(context.Background()).Err(); err != nil {
	// 		dep.Status = "unhealthly"
	// 		dep.Message = err.Error()
	// 	}

	// 	deps = append(deps, dep)
	// }

	this.Deps = deps

	c.JSON(http.StatusOK, this)
}
