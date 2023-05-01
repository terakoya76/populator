module github.com/terakoya76/populator

go 1.12

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.12.0

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.2
)
