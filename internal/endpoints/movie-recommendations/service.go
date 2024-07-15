package movierecommendations

import (
	"context"
	"fmt"
	similarusers "gn222gq/rec-sys/internal/endpoints/similar-users"
	endpointsutil "gn222gq/rec-sys/internal/endpoints/util"
	"gn222gq/rec-sys/internal/model"
)

type Repository interface {
  ValidateUserId(ctx context.Context, userId int) (bool, error)
	GetAllRatings(ctx context.Context) ([]model.Rating, error)
}

type Service struct {
  repo Repository
}

func NewService(repo Repository) *Service {
  return &Service{
    repo: repo,
  }
}

func (s *Service) GetMovieRecommendations(ctx context.Context, queryParams map[string]string) (string, error) {
  params, err := endpointsutil.ValidateQueryParams(queryParams)
  if err != nil {
    return "", err
  }

  validUserId, err := s.repo.ValidateUserId(ctx, params.UserId)
  if err != nil {
    return "", err
  }
  if !validUserId {
    return "", fmt.Errorf("Invalid user id")
  }

  ratings, err := s.repo.GetAllRatings(ctx)
  if err != nil {
    return "", err
  }

  similarityScores, err := similarusers.CalculateUserSimilarity(ratings, params.UserId, params.Algorithm)
  if err != nil {
    return "", err
  }

}
