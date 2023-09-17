package repository

import (
	"database/sql"
	"fmt"
)

func FindByID[T Table](tx *sql.Tx, id uint) (*T, error) {
	var (
		result = new(T)
		table  T
		err    error
	)

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table.Table())
	err = tx.QueryRow(query, id).Scan(&result)
	if err != nil {
		return nil, tx.Rollback()
	}

	return result, nil
}

func FindOne[T Table](tx *sql.Tx, where string, args ...interface{}) (*T, error) {
	var (
		result = new(T)
		table  T
		err    error
	)

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT 1", table.Table(), where)
	err = tx.QueryRow(query, args...).Scan(&result) // Adjust Scan to fit your table structure
	if err != nil {
		return nil, tx.Rollback()
	}

	return result, nil
}

func FindMany[T Table](tx *sql.Tx, where string, args ...interface{}) ([]*T, error) {
	var (
		results []*T
		table   T
		query   string
		rows    *sql.Rows
		err     error
	)

	query = fmt.Sprintf("SELECT * FROM %s WHERE %s", table.Table(), where)
	if rows, err = tx.Query(query, args...); err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var result = new(T)
		if err := rows.Scan(&result); err != nil { // Adjust Scan to fit your table structure
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func Insert[T Table](tx *sql.Tx, model T) error {
	var (
		table T
		err   error
		query string
	)

	query = fmt.Sprintf("INSERT INTO %s (%s) VALUES ?",
		table.Table(),
		model.Columns(),
		model.Values(),
	)

	_, err = tx.Exec(query)
	return err
}
