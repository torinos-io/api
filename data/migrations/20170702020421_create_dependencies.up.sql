create table dependencies (
    source_cd smallint not null default 0,

    project_uuid            uuid not null references projects(uuid),
    dependent_project_uuid  uuid not null references projects(uuid)
);

create index index_dependencies_on_project_uuid_and_dependent_project_uuid on dependencies (project_uuid, dependent_project_uuid);
