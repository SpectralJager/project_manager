-- init
create table if not exists users (
    id integer primary key,
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
    name text not null,
    description text,
    started_at timestamp,
    updated_at timestamp,
    finished_at timestamp
);

create table if not exists users_projects (
    id int primary key,
    user_id int not null,
    project_id int not null
);

create table if not exists users_tasks (
    id int primary key,
    user_id int not null,
    task_id int not null
);
