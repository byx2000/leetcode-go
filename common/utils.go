package common

func SliceToSet[E comparable](slice []E) map[E]bool {
	set := map[E]bool{}
	for _, e := range slice {
		set[e] = true
	}
	return set
}
