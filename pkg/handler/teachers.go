package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func (h *Handler) createTeacher(c *gin.Context) {
  id, _ := c.Get(userCtx)

  c.JSON(http.StatusOK, map[string]interface{}{
    "id": id,
  })
}

func (h *Handler) getAllTeachers(c *gin.Context) {

}

func (h *Handler) getTeacherById(c *gin.Context) {

}

func (h *Handler) updateTeacher(c *gin.Context) {

}

func (h *Handler) deleteTeacher(c *gin.Context) {

}
