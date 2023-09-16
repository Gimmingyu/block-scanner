package repository

import (
	"errors"
	"strconv"
	"strings"
)

type QueryBuilder interface {
	Select(columns ...string) QueryBuilder
	Table(table string) QueryBuilder
	Where(query string, args ...interface{}) QueryBuilder
	And(query string, args ...interface{}) QueryBuilder
	Or(query string, args ...interface{}) QueryBuilder
	Not(query string, args ...interface{}) QueryBuilder
	GroupBy(columns ...string) QueryBuilder
	Having(query string, args ...interface{}) QueryBuilder
	OrderBy(columns ...string) QueryBuilder
	Limit(limit int) QueryBuilder
	Offset(offset int) QueryBuilder
	Joins(table string) QueryBuilder
	SubQuery(query string, args ...interface{}) QueryBuilder
	Distinct() QueryBuilder
	Build() (string, []interface{}, error)
}

type queryBuilder struct {
	query string
	args  []interface{}
}

func (q *queryBuilder) Select(columns ...string) QueryBuilder {
	q.query = "SELECT " + columns[0]
	for _, column := range columns[1:] {
		q.query += ", " + column
	}

	return q
}

func (q *queryBuilder) Table(table string) QueryBuilder {
	if q.query == "" {
		q.query = "SELECT *"
	}
	q.query += " FROM " + table
	return q
}

func (q *queryBuilder) Where(query string, args ...interface{}) QueryBuilder {
	q.query += " WHERE " + query
	q.args = append(q.args, args...)
	return q
}

func (q *queryBuilder) And(query string, args ...interface{}) QueryBuilder {
	q.query += " AND " + query
	q.args = append(q.args, args...)
	return q
}

func (q *queryBuilder) Or(query string, args ...interface{}) QueryBuilder {
	q.query += " OR " + query
	q.args = append(q.args, args...)
	return q
}

func (q *queryBuilder) Not(query string, args ...interface{}) QueryBuilder {
	q.query += " NOT " + query
	q.args = append(q.args, args...)
	return q
}

func (q *queryBuilder) GroupBy(columns ...string) QueryBuilder {
	q.query += " GROUP BY " + columns[0]
	for _, column := range columns[1:] {
		q.query += ", " + column
	}

	return q
}

func (q *queryBuilder) Having(query string, args ...interface{}) QueryBuilder {
	q.query += " HAVING " + query
	q.args = append(q.args, args...)
	return q
}

func (q *queryBuilder) OrderBy(columns ...string) QueryBuilder {
	q.query += " ORDER BY " + columns[0]
	for _, column := range columns[1:] {
		q.query += ", " + column
	}

	return q
}

func (q *queryBuilder) Limit(limit int) QueryBuilder {
	q.query += " LIMIT " + strconv.Itoa(limit)
	return q
}

func (q *queryBuilder) Offset(offset int) QueryBuilder {
	q.query += " OFFSET " + strconv.Itoa(offset)
	return q
}

func (q *queryBuilder) Joins(table string) QueryBuilder {
	q.query += " JOIN " + table
	return q
}

func (q *queryBuilder) SubQuery(query string, args ...interface{}) QueryBuilder {
	q.query += " (" + query + ") "
	q.args = append(q.args, args...)
	return q
}

func (q *queryBuilder) Distinct() QueryBuilder {
	q.query = "SELECT DISTINCT " + strings.TrimPrefix(q.query, "SELECT ")
	return q
}

func (q *queryBuilder) Build() (string, []interface{}, error) {

	if !strings.Contains(q.query, "SELECT") {
		return "", nil, errors.New("query missing SELECT clause")
	}
	if !strings.Contains(q.query, "FROM") {
		return "", nil, errors.New("query missing FROM clause")
	}
	if strings.Count(q.query, "?") != len(q.args) {
		return "", nil, errors.New("mismatch between placeholders and argument count")
	}

	return q.query, q.args, nil
}

func NewQueryBuilder() QueryBuilder {
	return new(queryBuilder)
}
