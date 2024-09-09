-- name: CreateOrderFranchisesDetail :one
INSERT INTO public.order_franchises_detail(ofde_start_date, ofde_end_date, ofde_qty_unit, ofde_price, ofde_total_price, ofde_orfi_id, ofde_frch_id)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
    RETURNING *;

-- name: FindOrderFranchisesDetailById :one
SELECT ofde_id, ofde_start_date, ofde_end_date, ofde_qty_unit, ofde_price, ofde_total_price, ofde_orfi_id, ofde_frch_id
	FROM public.order_franchises_detail
    WHERE ofde_id = $1;

-- name: FindAllOrderFranchisesDetail :many
SELECT ofde_id, ofde_start_date, ofde_end_date, ofde_qty_unit, ofde_price, ofde_total_price, ofde_orfi_id, ofde_frch_id
	FROM public.order_franchises_detail;

-- name: DeleteOrderFranchisesDetail :exec
DELETE FROM public.order_franchises_detail
    WHERE ofde_id = $1
    RETURNING *;