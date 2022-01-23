package router

import (
	"errors"
	"net/http"

	"github.com/zzae/TIL/gin/protocol"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger log.Logger
}

func NewNameHandler() *Handler {
	return &Handler{
		logger: log.New("module", "GroupHandler"),
	}
}

func (h *Handler) GetName(c *gin.Context) {
	logEntry := h.logger
	req := protocol.ReqCreateName{}

	if err := c.ShouldBindJSON(&req); err != nil {
		protocol.RespError(c, http.StatusBadRequest, err, h.logger)
		return
	} else if uid, err := getUserID(c); err != nil {
		protocol.RespError(c, http.StatusBadRequest, err, logEntry)
	} else {
		logEntry = logEntry.New("user", uid, "param", req.Name)
		protocol.RespResponse(c, nil, logEntry)

}
