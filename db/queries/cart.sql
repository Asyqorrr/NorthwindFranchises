-- name: FindCartByCustomerId :many
SELECT c.cart_id, cart_user_id, cart_fr_id, cart_start_date, cart_end_date, cart_qty, 
	cart_price, cart_total_price, cart_modified, cart_status, cart_cart_id, u.user_name, fr.frch_name
	FROM public.cart c
	JOIN users u ON c.cart_user_id = u.user_id
	JOIN franchises fr ON c.cart_fr_id = fr.frch_id
    WHERE c.cart_user_id = $1;

-- name: FindCartByCustomerOrFranchises :one
SELECT c.cart_id, cart_user_id, cart_fr_id, cart_start_date, cart_end_date, cart_qty, 
	cart_price, cart_total_price, cart_modified, cart_status, cart_cart_id, u.user_name, fr.frch_name
	FROM public.cart c
	JOIN users u ON c.cart_user_id = u.user_id
	JOIN franchises fr ON c.cart_fr_id = fr.frch_id
    WHERE c.cart_user_id = $1 OR fr.frch_name = $1;

-- name: FindCartsbyId :one
SELECT cart_id, cart_user_id, cart_fr_id, cart_start_date, cart_end_date, cart_qty, 
	cart_price, cart_total_price, cart_modified, cart_status, cart_cart_id
    FROM cart
    WHERE cart_id = $1;

-- name: CreateCart :one
INSERT INTO public.cart(cart_user_id, cart_fr_id, cart_start_date, cart_end_date, cart_qty, cart_price, cart_total_price, cart_modified, cart_status, cart_cart_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    RETURNING *;

-- name: UpdateCartQTY :one
UPDATE public.cart 
    SET cart_qty = $1 
    WHERE cart_id = $2 
    RETURNING *;

-- name: DeleteCart :exec
DELETE FROM public.cart 
    WHERE cart_id = $1
    RETURNING *;

-- name: FindCartByCustomerAndFranchises :one
SELECT cart_id, cart_user_id, cart_fr_id, cart_start_date, cart_end_date, cart_qty, 
	cart_price, cart_total_price, cart_modified, cart_status, cart_cart_id
    FROM cart
    WHERE cart_user_id = $1 and cart_fr_id = $2 LIMIT 1;































-- -- name: FindCartByCustomerId :many
-- select cart_id,cu.customer_id,company_name,pr.product_id,product_name,cr.unit_price,
-- cr.unit_price,qty,(cr.unit_price*cr.qty)::decimal(18,2) as price,cart_created_on 
-- from carts cr 
-- join customers cu on cr.customer_id=cu.customer_id
-- join products pr on cr.product_id=pr.product_id
-- where cu.customer_id=$1;


-- -- name: FindCartByCustomerAndProduct :one
-- SELECT cart_id,customer_id,product_id,unit_price,qty,(unit_price*qty)::decimal(18,2) as price,cart_created_on 
-- FROM carts WHERE customer_id=$1 and product_id=$2 LIMIT 1;

-- -- name: FindCartByCustomerPaging :many
-- SELECT cart_id,customer_id,product_id,unit_price,qty,(unit_price*qty)::decimal(18,2) as price,cart_created_on
-- FROM carts
-- WHERE customer_id=$1
-- ORDER BY cart_id
-- LIMIT $2 OFFSET $3;	

-- -- name: CreateCart :one
-- INSERT INTO carts(
-- 	customer_id,product_id,unit_price,qty,cart_created_on)
-- 	VALUES ($1, $2, $3, $4, $5)
-- 	RETURNING *;

-- -- name: UpdateCartQty :one
-- UPDATE carts
-- 	SET qty=$1
-- 	WHERE cart_id=$2
-- 	RETURNING *;

-- -- name: DeleteCart :exec
-- DELETE FROM carts
-- 	WHERE cart_id=$1
--     RETURNING *;