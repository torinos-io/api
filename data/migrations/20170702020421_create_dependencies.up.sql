create table dependencies (
    project_uuid varchar(255) not null references projects(uuid),
    dependent_project_uuid varchar(255) not null references projects(uuid),
    source_cd smallint not null default 0
);

create index index_dependencies_on_project_uuid_and_dependent_project_uuid on (project_uuid, dependent_project_uuid);
