CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL,
    user_name VARCHAR(25),
    user_password VARCHAR(85),
    user_email VARCHAR(25),
    user_phone VARCHAR(15),
    user_token VARCHAR(255),
    CONSTRAINT user_id_pk PRIMARY KEY (user_id),
    CONSTRAINT user_name_uq UNIQUE (user_name),
    CONSTRAINT user_email_uq  UNIQUE (user_email),
    CONSTRAINT user_phone_uq UNIQUE (user_phone)
)