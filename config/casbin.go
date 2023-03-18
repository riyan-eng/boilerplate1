package config

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(Database)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}
	enforce, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	if hasPolicy := enforce.HasPolicy("admin", "/api/admin/*", "(GET)|(POST)|(PUT)|(DELETE)"); !hasPolicy {
		enforce.AddPolicy("admin", "/api/admin/*", "(GET)|(POST)|(PUT)|(DELETE)")
	}
	if hasPolicy := enforce.HasPolicy("user", "/api/users/:id/*", "(GET)|(PUT)"); !hasPolicy {
		enforce.AddPolicy("user", "/api/users/:id/*", "(GET)|(PUT)")
	}
	enforce.LoadPolicy()
	return enforce
}
