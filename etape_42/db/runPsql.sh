#!/bin/bash

docker run -it --rm \
			   --link postgres-sandbox-instance:postgres \
			   --env PGPASSWORD=postgres postgres        \
			   psql -h postgres -U postgres -d badges_db
