#!/bin/sh

SCRIPT_DIR=$(dirname "$0")

if [ ! -f "$SCRIPT_DIR/../migrations/000001_create_database.sql" ]; then
  echo "Error: File $SCRIPT_DIR/../migrations/000001_create_database.sql not found!"
  exit 1
fi

PGPASSWORD="db" psql -h localhost -d "postgres" -U "db" -f "$SCRIPT_DIR/../migrations/000001_create_database.sql"
