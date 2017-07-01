create table user (
    id serial primary key,
    
    user_name varchar(255),
    github_uuid varchar(255),
    github_access_token varchar(255)
);

crete unique index index_user_on_github_uuid on user (github_uuid);