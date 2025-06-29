package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/dto"
	"github.com/taufiktriantono/api-first-monorepo/pkg/errutil"
)

func (h *approvalrouter) listApprovalTemplate(c *gin.Context) {
	var req dto.ListTemplatRequest
	if err := c.ShouldBind(&req); err != nil {
		c.Error(err)
		return
	}

	resp, err := h.service.ListTemplate(c.Request.Context(), &req)
	if err != nil {
		if v, ok := err.(errutil.BaseError); ok {
			c.JSON(v.Code.HTTPStatus(), v.Err)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp)
}
