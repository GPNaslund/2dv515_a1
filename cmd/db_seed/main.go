package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"gn222gq/rec-sys/internal/db"
	"log"
	"os"
)

func main() {
  sqliteInstance := db.NewSqliteDb("/home/gpnaslund/GitHub/2dv515_a1/recsys.db")
  conn, err := sqliteInstance.GetConnection(context.Background())
  if err != nil {
    log.Fatalf("Failed to get database connection: %s", err)
  }

  ctx := context.Background()

  seed_users_table(conn, ctx)
  seed_movies_table(conn, ctx)
  seed_ratings_table(conn, ctx)

  println("All tables are initialized")
}

func seed_users_table(conn *sql.Conn, ctx context.Context) {
  query := "CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY, name TEXT)"
  _, err := conn.ExecContext(ctx, query)
  if err != nil {
    log.Fatal("Failed to create users table")
  }

  query = "SELECT * FROM users"
  err = conn.QueryRowContext(ctx, query).Scan()

  if err == nil {
    println("Users table is allready populated")
    return 
  }

  file, err := os.Open("/home/gpnaslund/GitHub/2dv515_a1/data/movies_large/users.csv")
  if err != nil {
    log.Fatalf("Failed to open users csv file: %s", err)
  }
  defer file.Close()

  reader := csv.NewReader(file)
  reader.Comma = ';'

  records, err := reader.ReadAll()
  if err != nil {
    log.Fatalf("Failed to read users csv records: %s", err)
  }

  query = "INSERT INTO users VALUES(?, ?)"
  for i, rec := range records {
    if i == 0 {
      continue
    }
    conn.ExecContext(ctx, query, rec[0], rec[1])
  }

  println("Successfully populated users table")
  return
}

func seed_movies_table(conn *sql.Conn, ctx context.Context) {
  query := "CREATE TABLE IF NOT EXISTS movies (id INT PRIMARY KEY, title TEXT, year INT)"
  _, err := conn.ExecContext(ctx, query)
  if err != nil {
    log.Fatal("Failed to create the movies table")
  }

  query = "SELECT * FROM movies"
  err = conn.QueryRowContext(ctx, query).Scan()

  if err == nil {
    println("Movies table is allready populated")
    return
  }

  file, err := os.Open("/home/gpnaslund/GitHub/2dv515_a1/data/movies_large/movies.csv")
  if err != nil {
    log.Fatalf("Failed to open movies csv file: %s", err)
  }
  defer file.Close()

  reader := csv.NewReader(file)
  reader.LazyQuotes = true
  reader.Comma = ';'
  
  records, err := reader.ReadAll()
  if err != nil {
    log.Fatalf("Failed to read movies csv records: %s", err)
  }

  query = "INSERT INTO movies VALUES (?, ?, ?)"
  for i, rec := range records {
    if i == 0 {
      continue
    }
    conn.ExecContext(ctx, query, rec[0], rec[1], rec[2])
  }

  println("Successfully populated movies table")
  return
}

func seed_ratings_table(conn *sql.Conn, ctx context.Context) {
  query := "CREATE TABLE IF NOT EXISTS ratings (user_id INT, movie_id INT, rating FLOAT, PRIMARY KEY (user_id, movie_id) FOREIGN KEY (user_id) REFERENCES users(id), FOREIGN KEY (movie_id) REFERENCES movies(id))"
  _, err := conn.ExecContext(ctx, query)
  if err != nil {
    log.Fatal("Failed to create ratings table")
  }

  query = "SELECT * FROM ratings"
  err = conn.QueryRowContext(ctx, query).Scan()

  if err == nil {
    println("Ratings table is allready populated")
    return
  }

  file, err := os.Open("/home/gpnaslund/GitHub/2dv515_a1/data/movies_large/ratings.csv")
  if err != nil {
    log.Fatalf("Failed to open ratings csv file: %s", err)
  }
  defer file.Close()

  reader := csv.NewReader(file)
  reader.Comma = ';'

  records, err := reader.ReadAll()
  if err != nil {
    log.Fatalf("Failed to read ratings csv records: %s", err)
  }

  query = "INSERT INTO ratings VALUES (?, ?, ?)"
  for i, rec := range records {
    if i == 0 {
      continue
    }
    conn.ExecContext(ctx, query, rec[0], rec[1], rec[2])
  }

  println("Successfully populated ratings table")
  return
}
