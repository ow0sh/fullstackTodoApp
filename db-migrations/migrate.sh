#! /bin/sh

echo "starting migrations in path $MIGRATIONS_PATH for $DB_URL"
migrate --path $MIGRATIONS_PATH -database $DB_URL -verbose up