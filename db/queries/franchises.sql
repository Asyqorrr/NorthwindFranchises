-- name: GetAllFranchises :many
SELECT frch_id, frch_name, frch_desc, frch_price_buyout, frch_price_yearly, frch_modified, frch_cate_id
	FROM franchises;

-- name: GetFranchisesById :one
SELECT frch_id, frch_name, frch_desc, frch_price_buyout, frch_price_yearly, frch_modified, frch_cate_id
	FROM franchises
    WHERE frch_id = $1;

-- name: GetFranchisesByNameAndCateId :one
SELECT frch_id, frch_name, frch_desc, frch_price_buyout, frch_price_yearly, frch_modified, frch_cate_id
	FROM franchises
    WHERE frch_name = $1 AND frch_cate_id = $2;

-- name: CreateFranchises :one
INSERT INTO franchises(
    frch_name, frch_desc, frch_price_buyout, frch_price_yearly, frch_modified, frch_cate_id)
	VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: DeleteFranchises :exec
DELETE FROM franchises
	WHERE frch_id = $1
    RETURNING *;