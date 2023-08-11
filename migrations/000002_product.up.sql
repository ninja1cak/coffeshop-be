CREATE TABLE public.product (
	product_id uuid NOT NULL DEFAULT gen_random_uuid(),
	product_name varchar(255) NOT NULL,
	product_desc text NULL,
	product_stock int4 NULL,
	product_type varchar(255) NULL,
	product_slug varchar(255) NOT NULL,
	isfavorite bool NULL DEFAULT false,
	delivery_method _varchar NULL,
	delivery_hour_start varchar(255) NULL,
	delivery_hour_end varchar(255) NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL,
	product_image text NULL,
	CONSTRAINT product_pkey PRIMARY KEY (product_id),
	CONSTRAINT product_product_slug_key UNIQUE (product_slug)
);