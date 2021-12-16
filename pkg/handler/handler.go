package handler

import (
   "github.com/gin-gonic/gin"
   "github.com/Sedokovka/simple-to-do-app/pkg/service"
)

type Handler struct {
  services *service.Service
}

func NewHandler(services *service.Service) *Handler {
  return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
  router := gin.New()

  auth := router.Group("/auth")
  {
    auth.POST("/sign-up", h.signUp)
    auth.POST("/sign-in", h.signIn)
  }

  api := router.Group("/api")
  {
    users := api.Group("/users")
    {
    users.POST("/", h.createUser)
    users.GET("/", h.getAllUsers)
    users.GET("/:id", h.getUserById)
    users.PUT("/:id", h.updateUser)
    users.DELETE("/:id", h.deleteUser)
    }

    teachers := api.Group("/teachers")
    {
      teachers.POST("/", h.createTeacher)
      teachers.GET("/", h.getAllTeachers)
      teachers.GET("/:id", h.getTeacherById)
      teachers.PUT("/:id", h.updateTeacher)
      teachers.DELETE("/:id", h.deleteTeacher)
    }

    pupils := api.Group("/pupils")
    {
      pupils.POST("/", h.createPupil)
      pupils.GET("/", h.getAllPupils)
      pupils.GET("/:id", h.getPupilById)
      pupils.PUT("/:id", h.updatePupil)
      pupils.DELETE("/:id", h.deletePupil)
    }

    parents := api.Group("/parents")
    {
      parents.POST("/", h.createParent)
      parents.GET("/", h.getAllParents)
      parents.GET("/:id", h.getParentById)
      parents.PUT("/:id", h.updateParent)
      parents.DELETE("/:id", h.deleteParent)
    }
  }

  return router
}
