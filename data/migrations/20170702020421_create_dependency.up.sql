create table dependency(
    project_uuid primary key foreign key not null references projects(uuid),
    dependent_project_uuid primary key foreign key not null references projects(uuid),
    source_cd smallint not null default 0
);