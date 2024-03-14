-- user
-- name: GetUserById :one
select * from users
where id = ?;
-- name: GetUserByCred :one
select id from users
where username = ? and password = ?;
-- name: GetUserByInvite :one
select * from users
where invite = ?;
-- name: CreateUser :exec
insert into users (username, password, invite)
values (?, ?, ?);

-- project
-- name: CreateProject :exec
insert into projects (name, description, leader_id)
values (?, ?, ?);
-- name: GetProjectById :one
select * from projects
where id = ?;
-- name: GetProjectsByLeader :many
select * from projects
where leader_id = ?
order by name;
-- name: GetProjectsByName :many
select * from projects
where name like '%' || ? || '%';
-- name: UpdateProjectDescriptionById :exec
update projects set description = ?
where id = ?;
-- name: UpdateProjectNameById :exec
update projects set name = ?
where id = ?;

-- tasks
-- name: CreateTask :exec
insert into tasks (project_id, mantainer_id, name, description, created_at)
values (?, ?, ?, ?, ?);
-- name: GetTaskById :one
select * from tasks
where id = ?;
-- name: GetTasksByProject :many
select * from tasks
where project_id = ?;
-- name: GetTasksByMaintainer :many
select * from tasks
where mantainer_id = ?;
-- name: GetTasksByName :many
select * from tasks
where name like '%' || ? || '%';
-- name: UpdateTaskStatusById :exec
update tasks set status = ?, updated_at = ?
where id = ?;
-- name: 