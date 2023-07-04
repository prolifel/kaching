include $(PWD)/.env

.PHONY: migrate migratemake

migrate:
	migrate -path database/migration/ -database "${DB_DRIVER}://${DB_USER}:${DB_PWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose ${type}

migratemake:
	migrate create -ext sql -dir database/migration/ -seq ${name}