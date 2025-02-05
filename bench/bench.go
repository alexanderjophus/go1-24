package bench

func sum(data ...int) int {
	var total int
	for _, d := range data {
		total += d
	}
	return total
}
