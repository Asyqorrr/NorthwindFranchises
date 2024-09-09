// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: order_franchises.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrderFranchises = `-- name: CreateOrderFranchises :one
INSERT INTO public.order_franchises(orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING orfi_id, orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id
`

type CreateOrderFranchisesParams struct {
	OrfiPurchaseNo *string     `json:"orfi_purchase_no"`
	OrfiTax        *float64    `json:"orfi_tax"`
	OrfiSubtotal   *float64    `json:"orfi_subtotal"`
	OrfiPatrxNo    *string     `json:"orfi_patrx_no"`
	OrfiType       *string     `json:"orfi_type"`
	OrfiModified   pgtype.Date `json:"orfi_modified"`
	OrfiUserID     *int32      `json:"orfi_user_id"`
}

func (q *Queries) CreateOrderFranchises(ctx context.Context, arg CreateOrderFranchisesParams) (*OrderFranchise, error) {
	row := q.db.QueryRow(ctx, createOrderFranchises,
		arg.OrfiPurchaseNo,
		arg.OrfiTax,
		arg.OrfiSubtotal,
		arg.OrfiPatrxNo,
		arg.OrfiType,
		arg.OrfiModified,
		arg.OrfiUserID,
	)
	var i OrderFranchise
	err := row.Scan(
		&i.OrfiID,
		&i.OrfiPurchaseNo,
		&i.OrfiTax,
		&i.OrfiSubtotal,
		&i.OrfiPatrxNo,
		&i.OrfiType,
		&i.OrfiModified,
		&i.OrfiUserID,
	)
	return &i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM public.order_franchises
	WHERE orfi_id = $1
    RETURNING orfi_id, orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id
`

func (q *Queries) DeleteOrder(ctx context.Context, orfiID int32) error {
	_, err := q.db.Exec(ctx, deleteOrder, orfiID)
	return err
}

const findAllOrderFranchises = `-- name: FindAllOrderFranchises :many
SELECT orfi_id, orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id
	FROM public.order_franchises
`

func (q *Queries) FindAllOrderFranchises(ctx context.Context) ([]*OrderFranchise, error) {
	rows, err := q.db.Query(ctx, findAllOrderFranchises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*OrderFranchise
	for rows.Next() {
		var i OrderFranchise
		if err := rows.Scan(
			&i.OrfiID,
			&i.OrfiPurchaseNo,
			&i.OrfiTax,
			&i.OrfiSubtotal,
			&i.OrfiPatrxNo,
			&i.OrfiType,
			&i.OrfiModified,
			&i.OrfiUserID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOrderFranchisesById = `-- name: FindOrderFranchisesById :one
SELECT orfi_id, orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id
	FROM public.order_franchises
    WHERE orfi_id = $1
`

func (q *Queries) FindOrderFranchisesById(ctx context.Context, orfiID int32) (*OrderFranchise, error) {
	row := q.db.QueryRow(ctx, findOrderFranchisesById, orfiID)
	var i OrderFranchise
	err := row.Scan(
		&i.OrfiID,
		&i.OrfiPurchaseNo,
		&i.OrfiTax,
		&i.OrfiSubtotal,
		&i.OrfiPatrxNo,
		&i.OrfiType,
		&i.OrfiModified,
		&i.OrfiUserID,
	)
	return &i, err
}

const updateOrder = `-- name: UpdateOrder :one
UPDATE public.order_franchises
	SET orfi_purchase_no=$2, orfi_tax=$3, orfi_subtotal=$4, orfi_patrx_no=$5, orfi_type=$6
	WHERE orfi_id = $1
	RETURNING orfi_id, orfi_purchase_no, orfi_tax, orfi_subtotal, orfi_patrx_no, orfi_type, orfi_modified, orfi_user_id
`

type UpdateOrderParams struct {
	OrfiID         int32    `json:"orfi_id"`
	OrfiPurchaseNo *string  `json:"orfi_purchase_no"`
	OrfiTax        *float64 `json:"orfi_tax"`
	OrfiSubtotal   *float64 `json:"orfi_subtotal"`
	OrfiPatrxNo    *string  `json:"orfi_patrx_no"`
	OrfiType       *string  `json:"orfi_type"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (*OrderFranchise, error) {
	row := q.db.QueryRow(ctx, updateOrder,
		arg.OrfiID,
		arg.OrfiPurchaseNo,
		arg.OrfiTax,
		arg.OrfiSubtotal,
		arg.OrfiPatrxNo,
		arg.OrfiType,
	)
	var i OrderFranchise
	err := row.Scan(
		&i.OrfiID,
		&i.OrfiPurchaseNo,
		&i.OrfiTax,
		&i.OrfiSubtotal,
		&i.OrfiPatrxNo,
		&i.OrfiType,
		&i.OrfiModified,
		&i.OrfiUserID,
	)
	return &i, err
}
