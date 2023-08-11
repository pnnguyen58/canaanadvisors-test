SELECT 'CREATE DATABASE canaanadvisors'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'canaanadvisors');

DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'canaanadvisors') THEN

      RAISE NOTICE 'Role "canaanadvisors" already exists. Skipping.';
ELSE
CREATE ROLE my_user LOGIN PASSWORD '1qazxsw2';
END IF;
END
$do$;

GRANT CONNECT ON DATABASE canaanadvisors TO canaanadvisors;
GRANT USAGE ON SCHEMA public TO canaanadvisors;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO canaanadvisors;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO canaanadvisors;
GRANT ALL PRIVILEGES ON DATABASE canaanadvisors TO canaanadvisors;



