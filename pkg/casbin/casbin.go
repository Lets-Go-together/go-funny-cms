package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gocms/pkg/config"
	"gocms/pkg/logger"
)

func Initialize() {
	fmt.Println(config.Dsn)
	a, err := gormadapter.NewAdapter("mysql", "funy_cms:JaSRrkmjDnbyr2id@tcp(175.24.233.215:3306)/", "funy_cms", "casbin")
	if err != nil {
		logger.PanicError(err, "new model Adapter", true)
	}
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		logger.PanicError(err, "init model Adapter", true)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		logger.PanicError(err, "init casbin Adapter", true)
	}

	config.Casbin = e
}
