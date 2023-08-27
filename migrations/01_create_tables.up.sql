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
