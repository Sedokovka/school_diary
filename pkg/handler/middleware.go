package handler

import (
  "github.com/gin-gonic/gin"
  "log"
  "strings"
  "errors"
)

const (
    authorizationHeader = "Authorization"
    userCtx = "userId"
)

func(h *Handler) userIdentity(c *gin.Context) {
  header := c.GetHeader(authorizationHeader)

  if header == "" {
    log.Println("empty header")
    return
  }

  headerParts := strings.Split(header, " ")
  if len(headerParts) != 2 {
    log.Println("invalid header")
    return
  }

  userId, err := h.services.Authorization.ParseToken(headerParts[1])
  if err != nil {
    log.Fatal("status unauthorized")
  }

  c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
  id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
