package similarusers

import (
	endpointsutil "gn222gq/rec-sys/internal/endpoints/util"
	"gn222gq/rec-sys/internal/model"
	"gn222gq/rec-sys/internal/util"
	"math"
)

type SimilarityScore struct {
	UserId   int
	UserName string
	Score    float64
}

func CalculateUserSimilarity(ratings []model.Rating, userId int, algorithm endpointsutil.SimilarityAlgorithm) ([]SimilarityScore, error) {
	groupedRatings := groupUserRatings(ratings)
	similarityScores := []SimilarityScore{}
	for id, userRatings := range groupedRatings {
		if id == userId {
			continue
		}
		baseUserRatings := []float64{}
		otherUserRatings := []float64{}
		for movieId, rating := range userRatings {
			user1Rating, ok := groupedRatings[userId][movieId]
			if ok {
				baseUserRatings = append(baseUserRatings, float64(user1Rating))
				otherUserRatings = append(otherUserRatings, float64(rating))
			}
		}
		if algorithm == endpointsutil.Euclidean {
			similarityScore, err := util.EuclideanDistance(baseUserRatings, otherUserRatings)
			if err != nil {
				return nil, err
			}
			similarityScores = append(similarityScores, SimilarityScore{UserId: id, Score: similarityScore})
		} else if algorithm == endpointsutil.Pearson {
      similarityScore, err := util.PearsonScore(baseUserRatings, otherUserRatings)
      if err != nil {
        return nil, err
      }
      similarityScores = append(similarityScores, SimilarityScore{ UserId: id, Score: similarityScore})
		}
	}
	return similarityScores, nil
}

func groupUserRatings(ratings []model.Rating) map[int]map[int]float32 {
	sortedRatings := map[int]map[int]float32{}
	for _, rating := range ratings {
		userId := rating.UserId
		movieId := rating.MovieId
		rating := rating.Rating

		val, ok := sortedRatings[userId]
		if ok {
			val[movieId] = rating
		} else {
			sortedRatings[userId] = map[int]float32{movieId: rating}
		}
	}

	return sortedRatings
}

func paginateSimilarityScores(scores []SimilarityScore, limit, page int) []SimilarityScore {
	amountScores := len(scores)
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= amountScores {
		return []SimilarityScore{}
	}

	endIndex = int(math.Min(float64(endIndex), float64(amountScores)))
	return scores[startIndex:endIndex]
}
