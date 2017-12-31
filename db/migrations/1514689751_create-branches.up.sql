CREATE TABLE branches (
    id bigserial NOT NULL PRIMARY KEY,
    bank varchar(100) NOT NULL,
    ifsc varchar(20) NOT NULL,
    micr varchar (20) NOT NULL,
    branch text,
    address text,
    city varchar(20) NOT NULL,
    district varchar (20) NOT NULL,
    state varchar(20) NOT NULL,
    contact varchar(20) NOT NULL,

    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);