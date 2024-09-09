CREATE TABLE IF NOT EXISTS franchises_iamges(
    frim_id SERIAL,
    frim_filename VARCHAR(125),
    frim_default VARCHAR(1),
    frim_frch_id INT,
    CONSTRAINT frim_id_pk PRIMARY KEY (frim_id),
    CONSTRAINT frim_frch_id_fk FOREIGN KEY (frim_frch_id)
        REFERENCES franchises(frch_id)
)