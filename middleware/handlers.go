package middleware

import (
  "net/http"

	"github.com/gin-gonic/gin"
  db "github.com/Nandgopal-R/todoPilot/db/sqlc"
)

type Handler struct {
	Queries *db.Queries
}

func NewHandler(q *db.Queries) *Handler {
	return &Handler{Queries: q}
}

func (h *Handler) GetTodoList(c* gin.Context){
  todos, err := h.Queries.GetTodos(c)
  if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}
