package config

import (
	"fmt"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/lib/pq"
)

func Casbin() *casbin.Enforcer {
	adapter, err := sqladapter.NewAdapter(PostgreSQLDB, "postgres", "permissions")
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
