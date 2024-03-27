package sqlkit

import (
	"fmt"
	"strings"
)

type queryType string

const (
	selectQueryType queryType = "SELECT"
	insertQueryType queryType = "INSERT"
	updateQueryType queryType = "UPDATE"
)

type QueryBuilder interface {
	Select(fields ...string) QueryBuilder
	From(table string) QueryBuilder
	Where(condition, operation string, arg any) QueryBuilder

	Insert() QueryBuilder
	Update() QueryBuilder

	Table(table string) QueryBuilder
	Set(field string, arg any) QueryBuilder
	Returning(fields ...string) QueryBuilder

	Build() (string, []any)
}

type queryBuilder struct {
	queryType queryType

	selectClause    string
	fromClause      string
	whereClause     string
	returningClause string

	table string

	keys []string
	args []any
}

func (b queryBuilder) Select(fields ...string) QueryBuilder {
	b.queryType = selectQueryType
	b.selectClause = "SELECT " + strings.Join(fields, ", ")
	return b
}

func (b queryBuilder) From(table string) QueryBuilder {
	b.fromClause = "FROM " + table
	b.table = table
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

func (b queryBuilder) Insert() QueryBuilder {
	b.queryType = insertQueryType
	return b
}

func (b queryBuilder) Update() QueryBuilder {
	b.queryType = updateQueryType
	return b
}

func (b queryBuilder) Table(table string) QueryBuilder {
	b.table = table
	return b
}

func (b queryBuilder) Set(field string, arg any) QueryBuilder {
	b.keys = append(b.keys, field)
	b.args = append(b.args, arg)
	return b
}

func (b queryBuilder) Returning(fields ...string) QueryBuilder {
	b.returningClause = "RETURNING " + strings.Join(fields, ", ")
	return b
}

func (b queryBuilder) Build() (string, []any) {
	var q strings.Builder
	switch b.queryType {
	case selectQueryType:
		q.WriteString(fmt.Sprintf("%s %s ", b.selectClause, b.fromClause))
	case insertQueryType:
		q.WriteString(fmt.Sprintf("INSERT INTO %s (", b.table))
		for i, key := range b.keys {
			if i == 0 {
				q.WriteString(key)
			} else {
				q.WriteString(fmt.Sprintf(", %s", key))
			}
		}
		q.WriteString(") VALUES (")
		for i := range b.keys {
			if i == 0 {
				q.WriteString(fmt.Sprintf("$%d", i+1))
			} else {
				q.WriteString(fmt.Sprintf(", $%d", i+1))
			}
		}
		q.WriteString(")")
		if b.returningClause != "" {
			q.WriteString(" " + b.returningClause)
		}
	case updateQueryType:
		q.WriteString(fmt.Sprintf("UPDATE %s SET ", b.table))
		for i, key := range b.keys {
			if i == 0 {
				q.WriteString(fmt.Sprintf("%s = $%d", key, i+1))
			} else {
				q.WriteString(fmt.Sprintf(", %s = $%d ", key, i+1))
			}
		}
	}

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
		keys: make([]string, 0),
		args: make([]any, 0),
	}
}
