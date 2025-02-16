module web_service

go 1.23

// replace github.com/mcpalek/golang-microservices/configloader => ../configloader
replace github.com/mcpalek/golang-microservices/configloader => /app/configloader

require (
	github.com/denisenkom/go-mssqldb v0.12.3
	github.com/mcpalek/golang-microservices/configloader v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
)
