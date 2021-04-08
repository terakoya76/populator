module github.com/terakoya76/populator

go 1.12

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.12.0

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.3
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.7
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20190712062909-fae7ac547cb7 // indirect
)
