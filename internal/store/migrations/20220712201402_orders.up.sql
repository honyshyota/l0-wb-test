CREATE TYPE delivery AS (
    name varchar(255),
    phone varchar(32),
    zip integer,
    city varchar(255),
    address varchar(255),
    region varchar(255),
    email varchar(255)
);

CREATE TYPE payment AS (
    transaction varchar(32),
    request_id varchar(32),
    currency varchar(32),
    provider varchar(32),
    amount integer,
    payment_dt integer,
    bank varchar(32),
    delivery_cost integer,
    goods_total integer,
    custom_fee integer
);

CREATE TYPE item AS (
    chrt_id integer,
    track_number varchar(32),
    price integer,
    rid varchar(255),
    name varchar(32),
    sale integer,
    size varchar(32),
    total_price integer,
    nm_id integer,
    brand varchar(255),
    status integer
);

CREATE TABLE orders (
    id bigserial not null,
    order_uid varchar(32) unique,
    track_num varchar(32) unique,
    entry varchar(32),
    delivery delivery,
    payment payment,
    items item[],
    locale varchar(32),
    internal_signature varchar(32),
    customer_id varchar(32),
    delivery_service varchar(32),
    shard_key varchar(32),
    sm_id integer,
    date_created varchar(32),
    oof_shard varchar(32) 
);