package sql

import (
	"strings"
)

type SQLBuilder struct {
	query      strings.Builder
	args       []interface{}
	conditions []string
	orderBy    []string
	limit      *int
	offset     *int
}

func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{
		args:       make([]interface{}, 0),
		conditions: make([]string, 0),
		orderBy:    make([]string, 0),
	}
}

func (b *SQLBuilder) Select(columns ...string) *SQLBuilder {
	b.query.WriteString("SELECT ")
	b.query.WriteString(strings.Join(columns, ", "))
	return b
}

func (b *SQLBuilder) From(table string) *SQLBuilder {
	b.query.WriteString(" FROM ")
	b.query.WriteString(table)
	return b
}

func (b *SQLBuilder) Where(condition string, args ...interface{}) *SQLBuilder {
	b.conditions = append(b.conditions, condition)
	b.args = append(b.args, args...)
	return b
}

func (b *SQLBuilder) Build() (string, []interface{}) {
	if len(b.conditions) > 0 {
		b.query.WriteString(" WHERE ")
		b.query.WriteString(strings.Join(b.conditions, " AND "))
	}

	if len(b.orderBy) > 0 {
		b.query.WriteString(" ORDER BY ")
		b.query.WriteString(strings.Join(b.orderBy, ", "))
	}

	if b.limit != nil {
		b.query.WriteString(" LIMIT ")
		b.query.WriteString(string(*b.limit))
	}

	if b.offset != nil {
		b.query.WriteString(" OFFSET ")
		b.query.WriteString(string(*b.offset))
	}

	return b.query.String(), b.args
}
