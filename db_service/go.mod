module github.com/mcpalek/golang-microservices/db_service

go 1.23

// replace github.com/mcpalek/golang-microservices/configloader => ../configloader
replace github.com/mcpalek/golang-microservices/configloader => /app/configloader



require (
	
	github.com/mcpalek/golang-microservices/configloader v0.0.0-00010101000000-000000000000
	github.com/microsoft/go-mssqldb v1.8.0
)

require (
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)
