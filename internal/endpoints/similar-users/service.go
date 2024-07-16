package similarusers

import (
	"cmp"
	"context"
	"fmt"
	endpointsutil "gn222gq/rec-sys/internal/endpoints/util"
	"gn222gq/rec-sys/internal/model"
	"slices"
)

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
	params, err := endpointsutil.ValidateQueryParams(queryParams)
	if err != nil {
		return nil, err
	}

	validUserId, err := s.repo.ValidateUserId(ctx, params.UserId)
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
	similarityScores, err := CalculateUserSimilarity(allRatings, params.UserId, params.Algorithm)
	if err != nil {
		return nil, err
	}

  slices.SortFunc(similarityScores, func(a, b SimilarityScore) int {
    return cmp.Compare(a.Score, b.Score)
  })

	paginatedSimilarityScores := paginateSimilarityScores(similarityScores, params.Limit, params.Page)
	completeSimilarityScores, err := s.getUsernames(ctx, paginatedSimilarityScores)
	if err != nil {
		return nil, err
	}
	return completeSimilarityScores, nil
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
				completeSimilarityScores = append(completeSimilarityScores, SimilarityScore{UserId: user.Id, UserName: user.Name, Score: score.Score})
			}
		}
	}
	return completeSimilarityScores, nil
}
