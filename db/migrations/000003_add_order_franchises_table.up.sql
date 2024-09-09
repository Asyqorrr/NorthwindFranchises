CREATE TABLE IF NOT EXISTS order_franchises(
    orfi_id SERIAL, 
    orfi_purchase_no VARCHAR(25),
    orfi_tax DECIMAL(3,2),
    orfi_subtotal DECIMAL(18,2),
    orfi_patrx_no VARCHAR(55),
    orfi_type VARCHAR(15),
    orfi_modified DATE,
    orfi_user_id INT,
    CONSTRAINT orfi_id_pk PRIMARY KEY (orfi_id),
    CONSTRAINT orfi_purchase_no_uq UNIQUE (orfi_purchase_no),
    CONSTRAINT orfi_user_id_fk FOREIGN KEY (orfi_user_id) REFERENCES users(user_id)
)