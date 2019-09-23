package model

// Postgreql is
type Postgresql struct {
	Host     string
	DbName   string
	User     string
	Password string
	Port     int
	SSLMode  string
}
