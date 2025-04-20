package router

import(
  "github.com/Nandgopal-R/todoPilot/middleware"
  "github.com/gin-gonic/gin"
)

func Router(h *middleware.Handler) *gin.Engine {
  r:=gin.Default()

  r.GET("/todos",h.GetTodoList)

  r.POST("/addTask", h.AddTask)

  r.POST("/todos/:id", h.CompleteTask)

  return r
}
