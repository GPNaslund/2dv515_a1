package movierecommendations

import (
	similarusers "gn222gq/rec-sys/internal/endpoints/similar-users"
	"gn222gq/rec-sys/internal/model"
	"math"
)

type MovieRecommendation struct {
  Title string `json:"title"`
  MovieId int `json:"movie_id"`
  Score float64 `json:"score"`
}

type MovieSimilarityScore struct {
  userSimilarityScore float64
  movieRecommendationScore float64
}

func GetMovieRecommendations(similarityScores []similarusers.SimilarityScore, ratings []model.Rating) ([]MovieRecommendation, error) {
  movieScores := map[int][]MovieSimilarityScore{}
  
  for _, rating := range ratings {
    for _, similarity := range similarityScores {
      if similarity.UserId == rating.UserId {
        score := MovieSimilarityScore { userSimilarityScore: similarity.Score, movieRecommendationScore: similarity.Score * float64(rating.Rating) }
        _, ok := movieScores[rating.MovieId]
        if !ok {
          movieScores[rating.MovieId] = []MovieSimilarityScore{ score }
        } else {
          movieScores[rating.MovieId] = append(movieScores[rating.MovieId], score)
        }
      }
    }
  }

  recommendations := []MovieRecommendation{}
  for key, values := range movieScores {
    var total float64
    var simTotal float64
    for _, score := range values {
      total += score.movieRecommendationScore
      simTotal += score.userSimilarityScore
    }
    recScore := total / simTotal
    recommendation := MovieRecommendation{ MovieId: key, Score: recScore}
    recommendations = append(recommendations, recommendation)
  }

  return recommendations, nil
}

func paginateMovieRecommendations(scores []MovieRecommendation, limit, page int) []MovieRecommendation {
	amountRecommendations := len(scores)
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= amountRecommendations {
		return []MovieRecommendation{}
	}

	endIndex = int(math.Min(float64(endIndex), float64(amountRecommendations)))
	return scores[startIndex:endIndex]
}


