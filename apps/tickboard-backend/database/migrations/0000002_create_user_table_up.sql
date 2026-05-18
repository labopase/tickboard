CREATE TABLE users.users (
    id uuid primary key default gen_random_uuid(),
    code varchar(15) not null,
    email varchar(100) unique not null,
    phone_area varchar(8) not null,
    phone_number varchar(25) not null,
    password varchar(255) not null,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    verified_date timestamp,
    created_at timestamp not null default current_timestamp,
    created_by varchar(100) not null,
    modified_at timestamp not null default current_timestamp,
    modified_by varchar(100) not null,
    deleted_at timestamp,
    deleted_by varchar(100)
);