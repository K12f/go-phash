package text

import "math"

// 首先要分词
// 文字相似性检测
// cos余弦定理
func CosineSimilar(srcWords, dstWords []string) float64 {
	// get all words
	allWordsMap := make(map[string]int, 0)
	for _, word := range srcWords {
		if _, found := allWordsMap[word]; !found {
			allWordsMap[word] = 1
		} else {
			allWordsMap[word] += 1
		}
	}
	for _, word := range dstWords {
		if _, found := allWordsMap[word]; !found {
			allWordsMap[word] = 1
		} else {
			allWordsMap[word] += 1
		}
	}

	// stable the sort
	allWordsSlice := make([]string, 0)
	for word, _ := range allWordsMap {
		allWordsSlice = append(allWordsSlice, word)
	}

	// assemble vector
	srcVector := make([]int, len(allWordsSlice))
	dstVector := make([]int, len(allWordsSlice))
	for _, word := range srcWords {
		if index := indexOfSlice(allWordsSlice, word); index != -1 {
			srcVector[index] += 1
		}
	}
	for _, word := range dstWords {
		if index := indexOfSlice(allWordsSlice, word); index != -1 {
			dstVector[index] += 1
		}
	}
	// calc cos
	numerator := float64(0)
	srcSq := 0
	dstSq := 0
	for i, srcCount := range srcVector {
		dstCount := dstVector[i]
		numerator += float64(srcCount * dstCount)
		srcSq += srcCount * srcCount
		dstSq += dstCount * dstCount
	}
	denominator := math.Sqrt(float64(srcSq * dstSq))

	return numerator / denominator
}
func indexOfSlice(ss []string, s string) (index int) {
	index = -1
	for k, v := range ss {
		if s == v {
			index = k
			break
		}
	}

	return
}
