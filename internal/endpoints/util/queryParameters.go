package endpointsutil



type SimilarityAlgorithm int

func (s SimilarityAlgorithm) ToString() string {
	switch s {
	case Euclidean:
		return "euclidean"
	case Pearson:
		return "pearson"
	default:
		return "Invalid"
	}
}


const (
	Euclidean SimilarityAlgorithm = iota
	Pearson
	Unkown
)

type QueryParameters struct {
	UserId    int
	Algorithm SimilarityAlgorithm
	Limit     int
	Page      int
}
