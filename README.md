## Migrate
Before migrate install in global env tool for migrations: `migrate`
```
go get -u github.com/golang-migrate/migrate/v4
```

```
migrate -source file://migrations -database postgresql://postgres:postgres@localhost:5432/gotest up
```