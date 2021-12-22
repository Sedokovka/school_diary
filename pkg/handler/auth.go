package handler

import (
  "github.com/gin-gonic/gin"
   "github.com/Sedokovka/simple-to-do-app"
   "net/http"
   "log"
)

func (h *Handler) signUp (c *gin.Context) {
  var input crud.User
  if err := c.BindJSON(&input); err != nil {
    // newErrorResponse(c, http.StatusBadRequest, err.Error())
    log.Fatal(err.Error())
    return
  }

  id, err := h.services.Authorization.CreateUser(input)
  if err != nil {
    log.Fatal(err.Error())
    return
  }

  c.JSON(http.StatusOK, map[string]interface{}{
    "id": id,
  })
}
type signInInput struct{
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn (c *gin.Context) {
  var input signInInput
  if err := c.BindJSON(&input); err != nil {

    log.Fatal(err.Error())
    return
  }

  token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
  if err != nil {
    log.Fatal(err.Error())
    return
  }

  c.JSON(http.StatusOK, map[string]interface{}{
    "token": token,
  })
}
