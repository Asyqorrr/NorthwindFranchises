ALTER TABLE cart
DROP CONSTRAINT cart_cart_id_fk,
DROP CONSTRAINT cart_id_pk,
DROP CONSTRAINT cart_user_id_uq;

ALTER TABLE cart
ADD CONSTRAINT cart_user_id_pk PRIMARY KEY (cart_id, cart_user_id);