create table dependency (
    project_uuid foreign key not null references projects(uuid),
    dependent_project_uuid foreign key not null references projects(uuid),
    source_cd smallint not null default 0
    primary key(project_uuid, dependent_project_uuid)
);
