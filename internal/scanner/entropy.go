package scanner

import (
	"math"
)

// calculateEntropy computes Shannon entropy of a string
func calculateEntropy(s string) float64 {
	freq := make(map[rune]float64)
	length := float64(len(s))

	for _, char := range s {
		freq[char]++
	}

	var entropy float64
	for _, count := range freq {
		p := count / length
		entropy -= p * math.Log2(p)
	}

	return entropy
}

// isHighEntropy checks if a string has high entropy (likely a secret)
func isHighEntropy(s string) bool {
	threshold := 4.5 // Adjust threshold as needed (higher = stricter)
	return len(s) > 10 && calculateEntropy(s) > threshold
}
