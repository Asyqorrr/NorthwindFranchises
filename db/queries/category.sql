-- name: FindCategoryById :one
SELECT cate_id, cate_name
	FROM category WHERE cate_id=$1;

-- name: FindAllCategory :many
SELECT cate_id, cate_name FROM category;

-- name: CreateCategory :one
INSERT INTO category(
	 cate_name)
	VALUES ($1) RETURNING *;

-- name: UpdateCategory :one
UPDATE category
	SET cate_name=$2
	WHERE cate_id=$1    
    RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category
	WHERE cate_id=$1
    RETURNING *;