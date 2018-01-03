CREATE TABLE search (
    id bigserial NOT NULL PRIMARY KEY,
    key varchar(256) NOT NULL,
    branch bigint NOT NULL,
    weight integer NOT NULL,
    UNIQUE (key, branch),

    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);