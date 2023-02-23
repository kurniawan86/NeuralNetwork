package main

func LinearEq(m []float32, x []float32) float32 {
	var tot float32
	for i := 0; i < len(m); i++ {
		tot = (m[i] * x[i]) + tot
	}
	return tot
}
