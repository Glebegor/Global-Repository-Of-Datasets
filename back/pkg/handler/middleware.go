package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userCtx2            = "userUsername"
	userCtx3            = "userCountRequests"
)

func (h *Handler) userIndentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}
	userId, userName, userCountRequests, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
	c.Set(userCtx2, userName)
	c.Set(userCtx3, userCountRequests)
}

func GetUserById(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newResponse(c, http.StatusInternalServerError, "user id is not found")
		return 0, errors.New("user id is not found")
	}
	idint, ok := id.(int)
	if !ok {
		newResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	return idint, nil
}
