-- +goose Up
-- +goose StatementBegin
CREATE TABLE todoList(
  id SERIAL PRIMARY KEY,
  task VARCHAR(255) NOT NULL,
  completed BOOLEAN DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todoList;
-- +goose StatementEnd
