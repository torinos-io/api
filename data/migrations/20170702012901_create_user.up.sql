create table user (
    id serial primary key,
    user_name varchar(255),
    github_uuid varchar(255) UNIQUE,
    github_access_token varchar(255)
);