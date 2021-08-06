create table if not exists users (
    id uuid primary key ,
    email varchar not null unique ,
    phone_number varchar not null unique
);

create table if not exists purchases (
    id uuid primary key ,
    user_id uuid not null references users(id),
    goods text not null ,
    total_price numeric(20, 2)  not null
);

create table if not exists logs (
    id uuid primary key ,
    purchase_id uuid not null references purchases(id),
    error text not null
)