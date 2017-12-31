CREATE TABLE branches (
    id bigserial NOT NULL PRIMARY KEY,
    bank varchar(128) NOT NULL,
    ifsc varchar(32) NOT NULL,
    micr varchar (32) NOT NULL,
    branch text,
    address text,
    city varchar(64) NOT NULL,
    district varchar (64) NOT NULL,
    state varchar(64) NOT NULL,
    contact varchar(32) NOT NULL,

    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);