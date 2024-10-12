#!/bin/sh

psql -h localhost -d "postgres" -password "db" -U "db" -f ./migrations/000001_create_databases.sql