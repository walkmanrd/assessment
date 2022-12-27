package configs

import (
	"database/sql"
	"testing"
)

type mockDB struct {
	query        string
	lastInsertId int64
	rowsAffected int64
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}

func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsertId, nil
}

func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func TestAutoMigrate(t *testing.T) {
	mock := &mockDB{}

	AutoMigrate(mock)

	if mock.query != `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	` {
		t.Error("should have been call db.Exec with but it's not")
	}
}
