#!/bin/bash
set -eux

migrate \
    -path "database/mysql/migration" \
    -database "mysql://$TODO_API_DATABASE_DSN" \
    "$MIGRATION_DIRECTION"
