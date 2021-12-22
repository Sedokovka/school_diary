package handler

import ("github.com/gin-gonic/gin"
        "text/template")

var tmpl = template.Must(template.ParseGlob("templates/*"))


func (h *Handler) getAllUsers(c *gin.Context) {

  // tmpl.ExecuteTemplate(w, "Index")
}

func (h *Handler) getUserById(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}
