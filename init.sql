-- Enable dblink extension
CREATE EXTENSION IF NOT EXISTS dblink;

DROP DATABASE IF EXISTS csudh_test;
DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'csudh_dev') THEN
        PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE csudh_dev');
    END IF;

    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'csudh_test') THEN
        PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE csudh_test');
    END IF;
END
$$;
