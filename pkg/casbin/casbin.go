package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gocms/pkg/config"
	"gocms/pkg/logger"
)

func Initialize() {
	username := config.GetString("DB_USERNAME")
	password := config.GetString("DB_PASSWORD")
	host := config.GetString("DB_HOST")
	port := config.GetString("DB_PORT")
	database := config.GetString("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	a, err := gormadapter.NewAdapter("mysql", dsn, database) // Your driver and data source.
	logger.PanicError(err, "new adapter", true)
	Enforcer, err := casbin.NewEnforcer("casbin.conf", a)
	logger.PanicError(err, "new adapter", true)

	config.Enforcer = Enforcer
}
