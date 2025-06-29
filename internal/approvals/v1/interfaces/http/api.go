package http

import (
	"github.com/gin-gonic/gin"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/service"
	"github.com/taufiktriantono/api-first-monorepo/pkg/health"
	"go.uber.org/fx"
)

type approvalrouter struct {
	engine  *gin.Engine
	health  health.HealthService
	service service.ApprovalService
}

type Params struct {
	fx.In
	Engine  *gin.Engine
	Health  health.HealthService
	Service service.ApprovalService
}

func ProvideRouter(p Params) *approvalrouter {
	r := &approvalrouter{
		engine:  p.Engine,
		health:  p.Health,
		service: p.Service,
	}

	r.Routes()

	return r
}

func (h *approvalrouter) Routes() {
	if h.engine == nil {
		return
	}

	health := h.engine.Group("/health")
	{
		health.GET("/liveness", h.health.Liveness)
		health.GET("/readiness", h.health.Readiness)
	}

	v1 := h.engine.Group("/v1")
	{
		templates := v1.Group("/approval-templates")
		{
			templates.GET("", h.listApprovalTemplate)
			templates.POST("", h.createApprovalTemplate)
			template := templates.Group("/:id")
			{
				template.GET("", h.getApprovalTemplate)
				template.PUT("", h.updateApprovalTemplate)
			}
		}

		approvals := v1.Group("/approvals")
		{
			approvals.GET("", h.ListApproval)

			approval := approvals.Group("/:id")
			{
				approval.GET("", h.getApproval)
			}
		}

		approvalsteps := v1.Group("/approval-steps")
		{
			approvalsteps.GET("", h.listApprovalStep)
			approvalstep := approvalsteps.Group("/:id")
			{
				approvalstep.GET("", h.getApprovalStep)
			}
		}
	}
}
