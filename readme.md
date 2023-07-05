# kaching

## prerequisites
- install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

## database migration
```bash
# up or down the migration
make migrate type={up|down}

# make new migration
make migratemake name={migration_name}
```