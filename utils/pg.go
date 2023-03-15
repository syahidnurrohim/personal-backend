package utils

import (
	"database/sql"
	"fmt"
	"strings"
)

type pgConnection struct {
	connection *sql.DB
}

func PGConnection(pgc *sql.DB) *pgConnection {
	return &pgConnection{
		connection: pgc,
	}
}

func (pgc *pgConnection) Insert(table string, data map[string]interface{}) (*sql.Rows, error) {
	var keys []string
	var values []interface{}
	var valuesTemplate []string

	kLen := 1
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, v)
		valuesTemplate = append(valuesTemplate, "$"+fmt.Sprintf("%v", kLen))
		kLen++
	}

	queryStr := fmt.Sprintf(`insert into %s (%s) values (%s)`, table, strings.Join(keys, ", "), strings.Join(valuesTemplate, ", "))

	return pgc.connection.Query(queryStr, values...)
}

func (pgc *pgConnection) Query(query string, args ...any) (*sql.Rows, error) {
	return pgc.connection.Query(query, args...)
}

func (pgc *pgConnection) QueryRow(query string, args ...any) *sql.Row {
	return pgc.connection.QueryRow(query, args...)
}

func (pgc *pgConnection) BeginTransaction() {
	pgc.connection.Query(`begin transaction`)
}

func (pgc *pgConnection) Commit() {
	pgc.connection.Query(`commit`)
}

func (pgc *pgConnection) Rollback() {
	pgc.connection.Query(`rollback`)
}
