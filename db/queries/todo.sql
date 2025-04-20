-- name: GetTodos :many
SELECT * FROM todoList ORDER BY id;

-- name: AddTodo :exec
INSERT INTO todoList (
  task
) VALUES ( $1 );

-- name: CompleteTask :exec
UPDATE todoList
  SET completed = TRUE
  WHERE id = ($1);

-- name: GetTodosById :one
SELECT * FROM todoList WHERE id=$1;
