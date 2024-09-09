CREATE TABLE IF NOT EXISTS franchises (
    frch_id SERIAL,
    frch_name VARCHAR(55),
    frch_desc VARCHAR(125),
    frch_price_buyout DECIMAL(18,2),
    frch_price_yearly DECIMAL (18,2),
    frch_modified DATE,
    frch_cate_id INT,
    CONSTRAINT frch_id_pk PRIMARY KEY (frch_id),
    CONSTRAINT frch_name_uq UNIQUE (frch_name),
    CONSTRAINT frch_cate_id_fk  FOREIGN KEY (frch_cate_id) REFERENCES category(cate_id)
)