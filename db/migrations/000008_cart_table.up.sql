CREATE TABLE IF NOT EXISTS cart (
    cart_id SERIAL,
    cart_user_id INT,
    cart_fr_id INT,
    cart_start_date TIMESTAMP,
    cart_end_date TIMESTAMP,
    cart_qty INTEGER,
    cart_price DOUBLE PRECISION,
    cart_total_price DOUBLE PRECISION,
    cart_modified TIMESTAMP,
    cart_status VARCHAR(15),
    cart_cart_id INT,
    CONSTRAINT cart_id_pk PRIMARY KEY (cart_id),
    CONSTRAINT cart_user_id_uq UNIQUE (cart_user_id),
    CONSTRAINT cart_fr_id_uq UNIQUE (cart_fr_id),
    CONSTRAINT cart_user_id_fk FOREIGN KEY (cart_user_id) 
        REFERENCES users(user_id),
    CONSTRAINT cart_fr_id_fk FOREIGN KEY (cart_fr_id)
        REFERENCES franchises(frch_id),
    CONSTRAINT cart_cart_id_fk FOREIGN KEY (cart_cart_id)
        REFERENCES cart(cart_id)
)

