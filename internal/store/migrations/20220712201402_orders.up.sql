CREATE TABLE orders (
    id bigserial not null,
    data jsonb
);

CREATE TABLE bad_messages (
    id bigserial not null,
    data jsonb
);