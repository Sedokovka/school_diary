package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/Sedokovka/simple-to-do-app"
  "net/http"
  )

func (h *Handler) CreatePupil(c *gin.Context) {
    userId, err := getUserId(c)
    if err!= nil {
      return
    }

    var input crud.Pupil
    if err := c.BindJSON(&input); err != nil {
      log.Fatal(err)
    }

    id, err := h.services.Pupil.CreatePupil(id.(int), input)
    if err != nil {
      log.Fatal("error")
    }

    c.JSON(http.StatusOK, map[string]interface{}{
      "id": id,
    })

}

func (h *Handler) GetAllPupils(c *gin.Context) {

}

func (h *Handler) GetPupilById(c *gin.Context) {

}

func (h *Handler) UpdatePupil(c *gin.Context) {

}

func (h *Handler) DeletePupil(c *gin.Context) {

}
