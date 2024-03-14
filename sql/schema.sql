-- init
create table if not exists users (
    id integer primary key,
    invite text unique not null,
    username text unique not null,
    password text unique not null
);

create table if not exists projects (
    id int primary key,
    name text not null,
    description text,
    leader_id int not null
);

create table if not exists tasks (
    id int primary key,
    project_id int not null,
    mantainer_id int not null,
    status int not null,
    name text not null,
    description text,
    created_at timestamp,
    updated_at timestamp
);