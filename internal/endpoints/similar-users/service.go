package similarusers

import (
	"context"
	"fmt"
	"gn222gq/rec-sys/internal/model"
	"math"
	"strconv"
)

type Parameters struct {
  userId int
  algorithm string
  limit int
  page int
}

type SimilarityScore struct {
	UserId int
  UserName string
	Score  float64
}

type Repository interface {
	GetAllRatings(ctx context.Context) ([]model.Rating, error)
	ValidateUserId(ctx context.Context, userId int) (bool, error)
  GetUsersFromIds(ctx context.Context, userIds []int) ([]model.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetSimilarUsers(ctx context.Context, queryParams map[string]string) ([]SimilarityScore, error) {
	params, err := s.validateQueryParams(queryParams)
	if err != nil {
		return nil, err
	}

	validUserId, err := s.repo.ValidateUserId(ctx, params.userId)
  if err != nil {
    return nil, err
  }
	if !validUserId {
		return nil, fmt.Errorf("invalid user id")
	}

	allRatings, err := s.repo.GetAllRatings(ctx)
	if err != nil {
		return nil, err
	}
  userId, _ := strconv.Atoi(queryParams["user"])
  similarityScores, err := CalculateSimilarity(allRatings, userId)
  if err != nil {
    return nil, err
  }

}

func (s *Service) getUsernames(ctx context.Context, scores []SimilarityScore) ([]SimilarityScore, error) {
  var userIds []int
  for _, score := range scores {
    userIds = append(userIds, score.UserId)
  }
  users, err := s.repo.GetUsersFromIds(ctx, userIds)
  if err != nil {
    return nil, err
  }

  var completeSimilarityScores []SimilarityScore
  
  for _, user := range users {
    for _, score := range scores {
      if user.Id == score.UserId {
        completeSimilarityScores = append(completeSimilarityScores, SimilarityScore{ UserId: user.Id, UserName: user.Name, Score: score.Score })
      }
    }
  }
  return completeSimilarityScores, nil
}

func (s *Service) paginateSimilarityScores(scores []SimilarityScore, limit, page int) []SimilarityScore {
  amountScores := len(scores)
  startIndex := (page - 1) * limit
  endIndex := startIndex + limit

  if startIndex >= amountScores {
    return []SimilarityScore{}
  }

  endIndex= int(math.Min(float64(endIndex), float64(amountScores)))
  return scores[startIndex:endIndex]
}

func (s *Service) validateQueryParams(queryParams map[string]string) (Parameters, error) {
  var params Parameters

  if len(queryParams) != 4 {
    return params, fmt.Errorf("Invalid amount of query parameters")
  }

	if queryParams["user"] == "" || queryParams["algorithm"] == "" || queryParams["limit"] == "" || queryParams["page"] == "" {
		return params, fmt.Errorf("Missing query parameter")
	}

	id, err := strconv.Atoi(queryParams["user"])
	if err != nil {
		return params, fmt.Errorf("Invalid user id format")
	}
  params.userId = id

  limit, err := strconv.Atoi(queryParams["limit"])
  if err != nil {
    return params, fmt.Errorf("Invalid limit value")
  }
  params.limit = limit

  page, err := strconv.Atoi(queryParams["page"])
  if err != nil {
    return params, fmt.Errorf("Invalid page value")
  }
  params.page = page


	if queryParams["algorithm"] != "euclidean" && queryParams["algorithm"] != "pearson" {
		return params, fmt.Errorf("Invalid algorithm parameter")
	}
  params.algorithm = queryParams["algorithm"]

	return params, nil
}
