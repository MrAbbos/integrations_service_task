DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles  -- SELECT list can be empty for this
      WHERE  rolname = 'postgres') THEN

CREATE ROLE postgres SUPERUSER ;
END IF;
END
$do$;

create table if not exists books (
    id uuid primary key not null,
    name varchar not null,
    number_of_pages int,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);
