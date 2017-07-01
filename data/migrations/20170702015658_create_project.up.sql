create table project (
    uuid varchar(255) primary key,
    user_id int foreign key references users(id),
    cart_file_content text,
    pods_file_content text,
    xcode_xml_content text,
    supported_swift_version varchar(255),
    repository varchar(255),
    last_fetched_at timestamp without time zone,
    state_cd smallint default 0 not null
);

create uique index index_project_on_repository on project (repository);