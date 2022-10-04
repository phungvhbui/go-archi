package connector

import (
	"fmt"
	"strings"

	"github.com/phungvhbui/go-archi/internal/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB(
	scheme, host string, port int, database, username, password, params string,
) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if strings.EqualFold(scheme, "mysql") {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?%s",
			username,
			password,
			host,
			port,
			database,
			params,
		)

		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                      dsn,
			DefaultStringSize:        256,
			DisableDatetimePrecision: true,
		}))
	} else {
		err = fmt.Errorf("unknown database scheme specified '%s'", scheme)
	}

	err = db.AutoMigrate(
		&entity.Organization{},
		&entity.User{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
