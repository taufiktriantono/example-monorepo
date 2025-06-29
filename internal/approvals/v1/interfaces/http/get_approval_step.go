package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *approvalrouter) getApprovalStep(c *gin.Context) {
	c.Status(http.StatusServiceUnavailable)
}
