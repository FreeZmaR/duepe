package types

const (
	DatabaseSQLite   DatabaseType = "sqlite"
	DatabasePostgres DatabaseType = "postgres"
)

type Database struct {
	Host     string       `json:"host"`
	Port     string       `json:"port"`
	User     string       `json:"user"`
	Password string       `json:"password"`
	DBName   string       `json:"db_name"`
	DBType   DatabaseType `json:"db_type"`
}

type DatabaseType string

func NewDefaultDatabase() *Database {
	return &Database{
		Host:     "localhost",
		Port:     "5432",
		User:     "admin",
		Password: "admin",
		DBName:   "duepe",
		DBType:   DatabaseSQLite,
	}
}
