#!/bin/bash

dropdb $PG_DB_NAME || true
dropuser $PG_USER || true

createuser $PG_USER
psql -c "ALTER USER $PG_USER PASSWORD '$PG_PASS';"
createdb -O $PG_USER $PG_DB_NAME

./bin/out/fixture_db