package similarusers

import (
	"gn222gq/rec-sys/internal/model"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestGroupRatings_ShouldGroupMovieRatingsWithUser(t *testing.T) {
	user1Ratings := map[int]float32{}
	user1Ratings[1] = 1.0
	user1Ratings[2] = 2.0
	user1Ratings[3] = 3.0
	user1Ratings[4] = 4.0

	user2Ratings := map[int]float32{}
	user2Ratings[1] = 1.0
	user2Ratings[2] = 2.0

	user3Ratings := map[int]float32{}
	user3Ratings[1] = 1.0

	allRatings := []model.Rating{}

	for key, val := range user1Ratings {
		allRatings = append(allRatings, model.Rating{UserId: 1, MovieId: key, Rating: val})
	}

	for key, val := range user2Ratings {
		allRatings = append(allRatings, model.Rating{UserId: 2, MovieId: key, Rating: val})
	}

	for key, val := range user3Ratings {
		allRatings = append(allRatings, model.Rating{UserId: 3, MovieId: key, Rating: val})
	}

  randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
  randGenerator.Shuffle(len(allRatings), func(i, j int) {
    allRatings[i], allRatings[j] = allRatings[j], allRatings[i]
  })

  result := groupUserRatings(allRatings)

  if !reflect.DeepEqual(result[1], user1Ratings) {
    t.Fatalf("Invalid grouping of user1 ratings")
  }

  if !reflect.DeepEqual(result[2], user2Ratings) {
    t.Fatalf("Invalid grouping of user2 ratings")
  }

  if !reflect.DeepEqual(result[3], user3Ratings) {
    t.Fatalf("Invalid grouping of user3 ratings")
  }

}
