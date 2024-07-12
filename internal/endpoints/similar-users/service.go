package similarusers

import (
	"fmt"
	"strconv"
  "gn222gq/rec-sys/internal/model"
)

type Repository interface {
  getAllRatings() ([]model.Rating, error)
  validateUserId(userId string) bool
}

type Service struct {
  repo Repository
}

func NewService(repo Repository) *Service {
  return &Service{
    repo: repo,
  }
}

func (s *Service) getSimilarUsers(queryParams map[string]string) ([]model.SimilarityScore, error) {
  err := s.validateQueryParams(queryParams)
  if err != nil {
    return nil, err
  }
  
  validUserId := s.repo.validateUserId(queryParams["user"])
  if !validUserId {
    return nil, fmt.Errorf("invalid user id")
  }

  allRatings, err := s.repo.getAllRatings()
  if err != nil {
    return nil, err
  }

  allSimilarityScores := 
}

func (s *Service) calculateSimilarity([]model.Rating) {

}


func (s *Service) validateQueryParams(queryParams map[string]string) error {
  userId := queryParams["user"]
  algorithm := queryParams["algorithm"]
  limit := queryParams["limit"]
  page := queryParams["page"]

  if userId == "" || algorithm == "" || limit == "" || page == "" {
    return fmt.Errorf("Missing query parameter")
  }

  _, err := strconv.Atoi(userId)
  if err != nil {
    return fmt.Errorf("Invalid user id format")
  }

  if algorithm != "euclidean" && algorithm != "pearson" {
    return fmt.Errorf("Invalid algorithm parameter")
  }

  if len(queryParams) < 4 {
    return fmt.Errorf("Invalid query parameter present")
  }

  return nil
} 

