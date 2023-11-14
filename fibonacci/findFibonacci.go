package fibonacci

func Find(num int) int {

	fibSlice := make([]int, num+1)

	fibSlice[0] = 0
	fibSlice[1] = 1

	// fmt.Println(fibSlice)

	for i := 2; i <= num; i++ {
		fibSlice[i] = fibSlice[i-1] + fibSlice[i-2]
		// fmt.Println(fibSlice)
	}
	return fibSlice[num]
}
