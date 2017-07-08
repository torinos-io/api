create table users (
    id          serial primary key,
    github_uuid varchar(255),
    
    user_name           varchar(255),
    github_access_token varchar(255)
);

create unique index index_users_on_github_uuid on users (github_uuid);
