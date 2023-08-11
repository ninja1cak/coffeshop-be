CREATE TABLE public."user" (
	user_id uuid NULL DEFAULT gen_random_uuid() PRIMARY KEY,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	first_name varchar(50) NULL,
	last_name varchar(50) NULL,
	phone_number varchar(20) NOT NULL,
	address text NULL,
	birth_date date NULL,
	gender varchar(25) NULL,
	photo_profile text NULL,
	status varchar(10) NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL,
	"role" varchar(25) NULL,
	CONSTRAINT user_email_key UNIQUE (email),
	CONSTRAINT user_phone_number_key UNIQUE (phone_number)
);