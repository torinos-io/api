create table dependencies (
    project_uuid varchar(255) not null references projects(uuid),
    dependent_project_uuid varchar(255) not null references projects(uuid),
    source_cd smallint not null default 0
);
