package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *approvalrouter) listApprovalStep(c *gin.Context) {
	c.Status(http.StatusServiceUnavailable)
}
