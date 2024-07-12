package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gn222gq/rec-sys/internal/model"
)

type Repository struct {
  dbConn *sql.Conn
}

func NewRepository(dbConn *sql.Conn) *Repository {
  return &Repository{
    dbConn: dbConn,
  }
}


func (r *Repository) GetAllRatings(ctx context.Context) ([]model.Rating, error) {
  allRows, err := r.dbConn.QueryContext(ctx, "SELECT * FROM ratings")
  if err != nil {
    return nil, err
  }
  defer allRows.Close()

  var ratings []model.Rating

  for allRows.Next() {
    var rating model.Rating
    if err := allRows.Scan(&rating.UserId, &rating.MovieId, &rating.Rating); err != nil {
      return ratings, err
    }
    ratings = append(ratings, rating)
  }
  
  return ratings, nil
}

func (r *Repository) ValidateUserId(ctx context.Context, userId int) (bool, error) {
  var user model.User
  if err := r.dbConn.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", userId).Scan(&user.Id, &user.Name); err != nil {
    return false, err
  }
  return true, nil
}

func (r *Repository) GetUsersFromIds(ctx context.Context, userIds []int) ([]model.User, error) {
  var users []model.User
  queryStr := "SELECT * FROM users WHERE "
  for i, val := range userIds {
    if i == len(userIds) - 1 {
      queryStr += fmt.Sprintf("id = %d", val)
    } else {
      queryStr += fmt.Sprintf("id = %d OR ", val)
    }
  }
  allRows, err := r.dbConn.QueryContext(ctx, queryStr)
  if err != nil {
    return nil, err
  }
  defer allRows.Close()

  for allRows.Next() {
    var user model.User
    if err := allRows.Scan(&user.Id, &user.Name); err != nil {
      return users, nil
    }
    users = append(users, user)
  }
  return users, nil
}


