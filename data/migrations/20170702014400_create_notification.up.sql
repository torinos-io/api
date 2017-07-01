create table notification(
    id serial primary key,
    user_id int foreign key references users(id),
    project_uuid varchar(255) foreign key references projects(uuid) not null,
    email varchar(255) not null,
    deleted_at timestamp without time zone
);