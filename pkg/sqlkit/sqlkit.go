package sqlkit

import (
	"fmt"
	"strings"
)

type QueryBuilder interface {
	Select(fields ...string) QueryBuilder
	From(table string) QueryBuilder
	Where(condition, operation string, arg any) QueryBuilder
	Build() (string, []any)
}

type queryBuilder struct {
	selectClause string
	fromClause   string
	whereClause  string

	args []any
}

func (b queryBuilder) Select(fields ...string) QueryBuilder {
	b.selectClause = "SELECT " + strings.Join(fields, ", ")
	return b
}

func (b queryBuilder) From(table string) QueryBuilder {
	b.fromClause = "FROM " + table
	return b
}

func (b queryBuilder) Where(condition, operation string, arg any) QueryBuilder {
	if b.whereClause == "" {
		b.whereClause = fmt.Sprintf("WHERE %s %s $%d", condition, operation, b.idx())
	} else {
		b.whereClause += fmt.Sprintf(" AND %s %s $%d", condition, operation, b.idx())
	}
	b.args = append(b.args, arg)
	return b
}

func (b queryBuilder) Build() (string, []any) {
	var q strings.Builder
	q.WriteString(fmt.Sprintf("%s %s ", b.selectClause, b.fromClause))

	if b.whereClause != "" {
		q.WriteString(b.whereClause)
	}

	return q.String(), b.args
}

func (b queryBuilder) idx() int {
	return len(b.args) + 1
}

func NewQueryBuilder() QueryBuilder {
	return queryBuilder{
		args: make([]any, 0),
	}
}
