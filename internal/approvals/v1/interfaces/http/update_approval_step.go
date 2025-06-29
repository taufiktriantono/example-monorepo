package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *approvalrouter) updateApprovalStep(c *gin.Context) {
	c.Status(http.StatusServiceUnavailable)
}
