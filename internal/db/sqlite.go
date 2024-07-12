package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type SqliteDb struct {
  connectionString string
  db *sql.DB
}

func NewSqliteDb(connString string) *SqliteDb {
  db, err := sql.Open("sqlite", connString)
  if err != nil {
    fmt.Printf("could not open sqlite datbase connection: %s", err)
    os.Exit(1)
  }
  return &SqliteDb{
    db: db,
  }
}

func (s *SqliteDb) GetConnection(c context.Context) (*sql.Conn, error) {
  return s.db.Conn(c)
}


