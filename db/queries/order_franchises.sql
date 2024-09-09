-- name: FindOrderFranchisesById :one
SELECT orfi_id, orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id
	FROM public.order_franchises
    WHERE orfi_id = $1;

-- name: FindAllOrderFranchises :many
SELECT orfi_id, orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id
	FROM public.order_franchises;

-- name: CreateOrderFranchises :one
INSERT INTO public.order_franchises(orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM public.order_franchises
	WHERE orfi_id = $1
    RETURNING *;

-- name: UpdateOrder :one
UPDATE public.order_franchises
	SET orfi_purchase_no=$2, orfi_tax=$3, orfi_subtotal=$4, orfi_patrx_no=$5, orfi_type=$6
	WHERE orfi_id = $1
	RETURNING *;