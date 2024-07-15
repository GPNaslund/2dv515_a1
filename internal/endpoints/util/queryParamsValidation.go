package endpointsutil

import (
	"fmt"
	"strconv"
)

func ValidateQueryParams(queryParams map[string]string) (QueryParameters, error) {
	var params QueryParameters

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
	params.UserId = id

	limit, err := strconv.Atoi(queryParams["limit"])
	if err != nil {
		return params, fmt.Errorf("Invalid limit value")
	}
	params.Limit = limit

	page, err := strconv.Atoi(queryParams["page"])
	if err != nil {
		return params, fmt.Errorf("Invalid page value")
	}
	params.Page = page

	switch queryParams["algorithm"] {
	case "euclidean":
		params.Algorithm = Euclidean
	case "pearson":
		params.Algorithm = Pearson
	default:
		return params, fmt.Errorf("Invalid value for algorithm")
	}

	return params, nil
}
