package postgres

import "github.com/Masterminds/squirrel"

// psql holds a reference to squirrel.StatementBuilderType
// which is used to build SQL queries that compatible with PostgreSQL syntax
var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
