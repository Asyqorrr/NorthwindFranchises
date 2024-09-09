-- name: FindImageById :one
SELECT frim_filename, frim_default, frim_frch_id
	FROM franchises_images WHERE frim_id=$1;

-- name: FindAllImages :many
SELECT frim_filename, frim_default FROM franchises_images;

-- name: UploadImages :one
INSERT INTO franchises_images(
	 frim_filename, frim_default, frim_frch_id)
	VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateImages :one
UPDATE franchises_images
	SET frim_filename=$2,
    frim_default=$3
	WHERE frim_id=$1    
    RETURNING *;

-- name: DeleteImages :exec
DELETE FROM franchises_images
	WHERE frim_id=$1
    RETURNING *;