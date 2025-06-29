package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *approvalrouter) updateApprovalTemplate(c *gin.Context) {
	c.Status(http.StatusServiceUnavailable)
}
