create extension if not exists "uuid-ossp";

create table if not exists people (
    id uuid default uuid_generate_v4() primary key,
    name text not null,
    last_name text not null,
    patronymic text,
    gender text not null default '',
    age int not null default 0,
    nationalize text not null default ''
)