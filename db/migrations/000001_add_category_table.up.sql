CREATE TABLE IF NOT EXISTS public.category
(
	cate_id serial,
	cate_name VARCHAR(25),
	constraint cate_id_pk primary key(cate_id),
	constraint cate_name_uq unique(cate_name)
);