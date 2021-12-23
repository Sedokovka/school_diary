package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/Sedokovka/simple-to-do-app"
  "net/http"
  "strconv"
  "log"
  )

func (h *Handler) CreatePupil(c *gin.Context) {
    userId, err := getUserId(c)
    if err != nil {
      userId = 1
    }
    userId := 1
    var input crud.Pupil
    if err := c.BindJSON(&input); err != nil {
      log.Fatal(err)
    }

    id, err := h.services.Pupil.CreatePupil(userId, input)
    if err != nil {
      log.Fatal("error")
    }

    c.JSON(http.StatusOK, map[string]interface{}{
      "id": id,
    })

}

func (h *Handler) GetPupilById(c *gin.Context) {
      userId, err := getUserId(c)
      if err != nil {
        userId = 1
      }
      pupilId, err := strconv.Atoi(c.Param("id"))

      if err != nil {
        log.Fatal(err)
      }

      pupil, err := h.services.Pupil.GetPupilById(userId, pupilId)
      if err != nil {
        log.Fatal(err)
      }

      c.JSON(http.StatusOK, map[string]interface{}{
        "id": pupil.Id,
        "name": pupil.Name,
        "class_id": pupil.Class_id,
        "birthdate": pupil.Birthdate,
        "rfid_card_id": pupil.Rfid_card_id,
        "user_id" : pupil.User_id,
      })
}

func (h *Handler) GetAllPupils(c *gin.Context) {
  userId, err := getUserId(c)
  if err!= nil {
    userId = 1
  }
  var pupils []crud.Pupil
  pupils, err = h.services.Pupil.GetAllPupils(userId)
  if err != nil {
    log.Fatal(err)
  }
  returnedData := make([]map[string]interface{}, len(pupils))

  for i, pupil := range pupils {
    returnedData[i] = map[string]interface{}{
      "id": pupil.Id,
      "name": pupil.Name,
      "class_id": pupil.Class_id,
      "birthdate": pupil.Birthdate,
      "rfid_card_id": pupil.Rfid_card_id,
      "user_id" : pupil.User_id,
    }
  }
  c.JSON(http.StatusOK, returnedData)

}

func (h *Handler) UpdatePupil(c *gin.Context) {

}

func (h *Handler) DeletePupil(c *gin.Context) {

}
