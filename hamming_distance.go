package phash

//
func HammingDistance(h1, h2 []int) (distance int) {
	for i := 0; i < len(h1); i++ {
		if h1[i] != h2[i] {
			distance += 1
		}
	}
	return
}
