package middleware

import (
  "net/http"
  "strconv"

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

func (h *Handler) AddTask(c* gin.Context){
  var input struct {
		Task string `json:"task" binding:"required"`
	}

  if err := c.ShouldBindJSON(&input); err!= nil {
    c.JSON(http.StatusBadRequest, gin.H{"error":"invalid input"})
    return
  }

  err := h.Queries.AddTodo(c,input.Task)

  if err!=nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add task"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message":"task added successfully"})

}

func (h *Handler) CompleteTask(c* gin.Context){
  idParam := c.Param("id")

  id, err := strconv.Atoi(idParam)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error":"invalid task ID"})
    return
  }

  _, err = h.Queries.GetTodosById(c, int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

  err = h.Queries.CompleteTask(c, int32(id))

  if err != nil{
    c.JSON(http.StatusInternalServerError, gin.H{"error":"could not complete task"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message":"task completed"})
}
