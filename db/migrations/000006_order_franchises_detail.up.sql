CREATE TABLE IF NOT EXISTS order_franchises_detail (
    ofde_id SERIAL,
    ofde_start_date TIMESTAMP,
    ofde_end_date TIMESTAMP,
    ofde_qty_unit INT,
    ofde_price DOUBLE PRECISION,
    ofde_total_price DOUBLE PRECISION,
    ofde_orfi_id INT,
    ofde_frch_id INT,
    CONSTRAINT ofde_id_pk PRIMARY KEY (ofde_id),
    CONSTRAINT ofde_orfi_id_fk FOREIGN KEY (ofde_orfi_id) 
        REFERENCES order_franchises(orfi_id),
    CONSTRAINT ofde_frch_id_fk FOREIGN KEY (ofde_frch_id)
        REFERENCES franchises(frch_id)
)