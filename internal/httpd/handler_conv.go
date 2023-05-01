package httpd

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pandodao/botastic/api"
	"github.com/pandodao/botastic/models"
)

func (h *Handler) CreateConv(c *gin.Context) {
	var req api.CreateConvRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.respErr(c, http.StatusBadRequest, err)
		return
	}

	conv := &models.Conv{
		BotID: req.BotID,
	}
	if err := h.sh.CreateConv(c, conv); err != nil {
		h.respErr(c, http.StatusInternalServerError, err)
		return
	}

	h.respData(c, api.CreateConvResponse(conv.API()))
}

func (h *Handler) UpdateConv(c *gin.Context) {
	convIDStr := c.Param("conv_id")
	convID, err := uuid.Parse(convIDStr)
	if err != nil {
		h.respErr(c, http.StatusBadRequest, err)
		return
	}
	var req api.UpdateConvRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.respErr(c, http.StatusBadRequest, err)
		return
	}

	rowsAffected, err := h.sh.UpdateConv(c, convID, map[string]any{
		"bot_id": req.BotID,
	})
	if err != nil {
		h.respErr(c, http.StatusInternalServerError, err)
		return
	}
	if rowsAffected == 0 {
		h.respErr(c, http.StatusNotFound, errors.New("conv not found"))
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) GetConv(c *gin.Context) {
	convIDStr := c.Param("conv_id")
	convID, err := uuid.Parse(convIDStr)
	if err != nil {
		h.respErr(c, http.StatusBadRequest, err)
		return
	}
	conv, err := h.sh.GetConv(c, convID)
	if err != nil {
		h.respErr(c, http.StatusInternalServerError, err)
		return
	}
	if conv.ID == uuid.Nil {
		h.respErr(c, http.StatusNotFound, err)
		return
	}

	h.respData(c, api.GetConvResponse(conv.API()))
}

func (h *Handler) DeleteConv(c *gin.Context) {
	convIDStr := c.Param("conv_id")
	convID, err := uuid.Parse(convIDStr)
	if err != nil {
		h.respErr(c, http.StatusBadRequest, err)
		return
	}

	if err := h.sh.DeleteConv(c, convID); err != nil {
		h.respErr(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
