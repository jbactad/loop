//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/vektra/mockery"
)
