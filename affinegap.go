package affinegap

import (
	"math"
)

func AffineGapDistance(string_a, string_b string) float64 {

	const matchWeight = 1
	const mismatchWeight = 11
	const gapWeight = 10
	const spaceWeight = 7
	const abbreviation_scale = .125

	string1 := string_b
	string2 := string_a

	length1 := len(string1)
	length2 := len(string2)

	if string1 == string2 &&
		matchWeight == math.Min(math.Min(matchWeight, mismatchWeight), gapWeight) {
		return float64(matchWeight * length1)
	}

	if length1 < length2 {
		string1, string2 = string2, string1
		length1, length2 = length2, length1
	}

	D := make([]float64, length1+1)
	V_current := make([]float64, length1+1)
	V_previous := make([]float64, length1+1)

	var distance float64

	for j := 1; j < (length1 + 1); j++ {
		V_current[j] = gapWeight + float64(spaceWeight*j)
		D[j] = math.MaxInt32 //TODO maybe not 32
	}

	for i := 1; i < (length2 + 1); i++ {
		char2 := string2[i-1]

		//for _ in range(0, length1 + 1) :
		copy(V_previous, V_current)

		V_current[0] = gapWeight + float64(spaceWeight*i)
		I := float64(math.MaxInt32)

		for j := 1; j < (length1 + 1); j++ {
			char1 := string1[j-1]

			if j <= length2 {
				I = math.Min(I, V_current[j-1]+gapWeight) + spaceWeight
			} else {
				I = (math.Min(I, V_current[j-1]+gapWeight*abbreviation_scale) + spaceWeight*abbreviation_scale)
			}
			D[j] = math.Min(D[j], V_previous[j]+gapWeight) + spaceWeight

			var M float64
			if char2 == char1 {
				M = V_previous[j-1] + matchWeight
			} else {
				M = V_previous[j-1] + mismatchWeight
			}

			V_current[j] = math.Min(math.Min(I, D[j]), M)
		}
	}
	distance = V_current[length1]

	return distance
}
func NormalizedAffineGapDistance(string1, string2 string) float64 {

	length1 := len(string1)
	length2 := len(string2)

	normalizer := float64(length1 + length2)

	distance := AffineGapDistance(string1, string2)

	return distance / normalizer
}
