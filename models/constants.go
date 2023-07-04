package models

const (
	EnvAppTimezone    = "APP_TIMEZONE"
	EnvDBHost         = "DB_HOST"
	EnvDBPort         = "DB_PORT"
	EnvDBUser         = "DB_USER"
	EnvDBPwd          = "DB_PWD"
	EnvDBName         = "DB_NAME"
	EnvDBSSLMode      = "DB_SSL_MODE"
	EnvDBConnLifetime = "DB_CONN_LIFETIME"
	EnvDBConnMaxIdle  = "DB_CONN_MAX_IDLE"
	EnvDBConnMaxOpen  = "DB_CONN_MAX_OPEN"
)

const (
	ConstDatabaseTypePostgresql = "postgres"
	ConstDatabaseTypeMysql      = "mysql"
	ConstDefaultDateType        = "Y-m-d"
)

var AppAvaialbleDatabase = []string{
	ConstDatabaseTypePostgresql,
	ConstDatabaseTypeMysql,
}
