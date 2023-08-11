
--migration 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  

--table user
create table public.user(
	user_id uuid DEFAULT uuid_generate_v4(),
	email varchar(255) not null UNIQUE,
	password varchar(255) not null,
	first_name varchar(50) NULL,
	last_name varchar(50) NULL,
	phone_number varchar(20) not null unique,
	address text,
	birth_date date,
	gender varchar(25),
	photo_profile text,
	status varchar(10),
	created_at timestamp without time zone null default now(),
	updated_at timestamp without time zone null	
)

--table product
create table public.product(
	product_id uuid DEFAULT uuid_generate_v4() primary key,
	product_name varchar(255) not null,
	product_desc  text,
	product_stock int,
	product_type varchar(255),
	product_slug varchar(255) not null unique,
	isFavorite bool default false,
	delivery_method varchar[],
	delivery_hour_start varchar(255),
	delivery_hour_end varchar(255),
	created_at timestamp without time zone null default now(),
	updated_at timestamp without time zone null	
)

--table product size
create table public.product_size(
	psize_id uuid DEFAULT uuid_generate_v4() primary key,
	product_id uuid,
	product_size varchar(255),
	product_price int,
	updated_at timestamp without time zone null,
	constraint product_size_fk1 foreign key(product_id) references public.product(product_id) on delete cascade
)
