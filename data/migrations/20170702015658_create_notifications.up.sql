create table notifications (
    id serial primary key,

    email       varchar(255) not null,
    deleted_at  timestamp without time zone,

    user_id         int references users(id),
    project_uuid    uuid not null references projects(uuid)
);
