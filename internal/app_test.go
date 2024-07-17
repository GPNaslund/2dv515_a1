//go:build integration

package internal

import (
	"encoding/json"
	movierecommendations "gn222gq/rec-sys/internal/endpoints/movie-recommendations"
	similarusers "gn222gq/rec-sys/internal/endpoints/similar-users"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpoint_ShouldReturn200Status(t *testing.T) {
	app := NewApp().Create()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Endpoint did not respond with 200 status")
	}
}

func TestSimilarUserEndpoint_TobyTestCase(t *testing.T) {
	app := NewApp().Create()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/similar-users?user=7&algorithm=euclidean&limit=3&page=1", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %s", err)
	}
	defer resp.Body.Close()

	var similarUsers []similarusers.SimilarityScore
	err = json.NewDecoder(resp.Body).Decode(&similarUsers)
	if err != nil {
		t.Fatalf("Failed to decode json body")
	}

	first := similarUsers[0]
	second := similarUsers[1]
	third := similarUsers[2]

	if first.UserId != 5 {
		t.Fatalf("Expected first similar user to have id: %d, but got id: %d", 5, first.UserId)
	}
	if second.UserId != 3 {
		t.Fatalf("Expected second similar user to have id: %d, but got id: %d", 3, second.UserId)
	}
	if third.UserId != 4 {
		t.Fatalf("Expected third similar user to have id: %d, but got id: %d", 4, third.UserId)
	}
}

func TestMovieRecommendationsEndpoint_AngelaTestCase(t *testing.T) {
	app := NewApp().Create()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movie-recommendations?user=4&algorithm=euclidean&limit=3&page=1", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %s", err)
	}
	defer resp.Body.Close()

	var movieRecommendations []movierecommendations.MovieRecommendation
	err = json.NewDecoder(resp.Body).Decode(&movieRecommendations)
	if err != nil {
		t.Fatalf("Failed to decode json")
	}

	first := movieRecommendations[0]
	second := movieRecommendations[1]
	third := movieRecommendations[2]

	if first.MovieId != 1 {
		t.Fatalf("Expected top recommendation to be movie-id: %d, but got movie-id: %d", 1, first.MovieId)
	}

	if second.MovieId != 18 {
		t.Fatalf("Expected second recommendation to be movie-id: %d, but got movie-id: %d", 18, second.MovieId)
	}

	if third.MovieId != 29 {
		t.Fatalf("Expected third recommendation to be movie-id: %d, but got movie-id: %d", 29, third.MovieId)
	}

}
