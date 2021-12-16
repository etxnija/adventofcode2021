package day6

func simulate(scol []int, days int) int {
	// log.Println("start", scol)
	for i := 0; i < days; i++ {
		var newFish []int
		// log.Println("len", len(scol))
		for j := 0; j < len(scol); j++ {
			if scol[j] == 0 {
				scol[j] = 6
				newFish = append(newFish, 8)
				continue
			}
			scol[j] = scol[j] - 1
		}
		scol = append(scol, newFish...)
		// log.Println("day, scol", i, scol)
	}
	return len(scol)
}

func simulateWithCount(scool []int, days int) int {
	fish := make([]int, 9)

	for _, v := range scool {
		fish[v] = fish[v] + 1
	}
	for i := 0; i < days; i++ {
		tempFish := make([]int, 9)
		tempFish[0] = fish[1]
		tempFish[1] = fish[2]
		tempFish[2] = fish[3]
		tempFish[3] = fish[4]
		tempFish[4] = fish[5]
		tempFish[5] = fish[6]
		tempFish[6] = fish[7] + fish[0]
		tempFish[7] = fish[8]
		tempFish[8] = fish[0]

		fish = tempFish
	}

	// sum
	sum := 0
	for _, v := range fish {
		sum = sum + v
	}
	return sum
}
