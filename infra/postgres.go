package infra

import (
	"fmt"

	"github.com/chanpon2015/authentication/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Open is
func Open(psql *model.Postgresql) (*gorm.DB, error) {
	return gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			psql.Host,
			psql.Port,
			psql.User,
			psql.DbName,
			psql.Password,
			psql.SSLMode,
		),
	)
}
