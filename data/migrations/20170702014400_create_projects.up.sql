create table projects (
    uuid varchar(255) primary key,
    user_id int references users(id),
    cart_file_content text,
    pods_file_content text,
    xcode_xml_content text,
    supported_swift_version varchar(255),
    repository varchar(255),
    last_fetched_at timestamp without time zone,
    state_cd smallint not null default 0
);

create unique index index_projects_on_repository on projects (repository);
