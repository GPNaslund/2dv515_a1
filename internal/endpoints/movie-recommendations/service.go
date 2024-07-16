package movierecommendations

import (
	"cmp"
	"context"
	"fmt"
	similarusers "gn222gq/rec-sys/internal/endpoints/similar-users"
	endpointsutil "gn222gq/rec-sys/internal/endpoints/util"
	"gn222gq/rec-sys/internal/model"
	"slices"
)

type Repository interface {
	ValidateUserId(ctx context.Context, userId int) (bool, error)
	GetAllRatings(ctx context.Context) ([]model.Rating, error)
  GetMoviesFromIds(ctx context.Context, movieIds []int) ([]model.Movie, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetMovieRecommendations(ctx context.Context, queryParams map[string]string) ([]MovieRecommendation, error) {
	params, err := endpointsutil.ValidateQueryParams(queryParams)
	if err != nil {
		return nil, err
	}

	validUserId, err := s.repo.ValidateUserId(ctx, params.UserId)
	if err != nil {
		return nil, err
	}
	if !validUserId {
		return nil, fmt.Errorf("Invalid user id")
	}

	ratings, err := s.repo.GetAllRatings(ctx)
	if err != nil {
		return nil, err
	}

	similarityScores, err := similarusers.CalculateUserSimilarity(ratings, params.UserId, params.Algorithm)
	if err != nil {
		return nil, err
	}

	movieRecommendations, err := GetMovieRecommendations(similarityScores, ratings)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(movieRecommendations, func(a, b MovieRecommendation) int {
		return cmp.Compare(a.Score, b.Score)
	})

  paginatedResult := paginateMovieRecommendations(movieRecommendations, params.Limit, params.Page)
  var movieIds []int
  for _, recommendation := range paginatedResult {
    movieIds = append(movieIds, recommendation.MovieId)
  }

  movies, err := s.repo.GetMoviesFromIds(ctx, movieIds)
  if err != nil {
    return nil, err
  }

  var result []MovieRecommendation

  for _, movie := range movies {
    for _ , movieScore := range paginatedResult {
      if movie.MovieId == movieScore.MovieId {
        result = append(result, MovieRecommendation{Title: movie.Title, MovieId: movieScore.MovieId, Score: movieScore.Score })
      }
    }
  }

  return result, nil
  
}
