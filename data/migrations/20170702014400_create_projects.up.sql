create table projects (
    uuid uuid primary key,

    cartfile_content text,
    podfile_lock_content text,
    pbxproj_content text,
    supported_swift_version varchar(255),
    repository varchar(255),
    last_fetched_at timestamp without time zone,
    state_cd smallint not null default 0,

    user_id int references users(id) on delete cascade
);

create unique index index_projects_on_repository on projects (repository);
