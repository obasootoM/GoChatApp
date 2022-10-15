-- name: CreatePost :one
INSERT INTO post (
  username,content
) VALUES ($1,$2)
RETURNING *;