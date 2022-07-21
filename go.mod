module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/mattn/go-sqlite3 v1.14.14 // indirect
	github.com/rs/xid v1.4.0
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	gorm.io/driver/mysql v1.3.5
	gorm.io/driver/postgres v1.3.8
	gorm.io/driver/sqlite v1.3.6
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.8
)

replace gorm.io/gorm => ./gorm
