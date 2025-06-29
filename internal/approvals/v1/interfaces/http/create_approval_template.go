package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *approvalrouter) createApprovalTemplate(c *gin.Context) {
	c.Status(http.StatusServiceUnavailable)
}
