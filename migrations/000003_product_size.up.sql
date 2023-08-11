CREATE TABLE public.product_size (
	psize_id uuid NOT NULL DEFAULT gen_random_uuid(),
	product_id uuid NULL,
	product_size varchar(255) NULL,
	product_price int4 NULL,
	updated_at timestamp NULL,
	CONSTRAINT product_size_pkey PRIMARY KEY (psize_id)
);


ALTER TABLE public.product_size ADD CONSTRAINT product_size_fk1 FOREIGN KEY (product_id) REFERENCES public.product(product_id) ON DELETE CASCADE;