package main

func average(xs []float) (ave float) {
	sum := 0
	switch len(xs) {
	case 0:
		ave = 0
	default:
		for _, v := range xs {
			sum += v
		}
		ave = sum / len(xs)
	}
	return
}

func main() {
	/* ... */
}
