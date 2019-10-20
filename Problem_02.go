
func solution(availableProgrammes []string, carbonToOffset int) []string {
    

   
	costeTotal := []int{}

	for _, prog := range availableProgrammes {

		cost := 0
		if prog == "s1" {

			if carbonToOffset < 500 {

				cost = carbonToOffset * 125
				fmt.Println(cost)
				costeTotal = append(costeTotal, cost)

			} else if carbonToOffset >= 500 {

				cost = carbonToOffset * 1
				costeTotal = append(costeTotal, cost)
			}
		}

		if prog == "s2" {

			if carbonToOffset > 300 {

				cost = 300 * 15
				rest := carbonToOffset - 3000

				cost = cost + (rest * 05)
				costeTotal = append(costeTotal, cost)

			}
		}

		if prog == "b1" {

			cost = carbonToOffset * 08
			costeTotal = append(costeTotal, cost)

		}

		if prog == "h1" {

			cost = 100.0 + (carbonToOffset * 05)
			costeTotal = append(costeTotal, cost)

		}

		if prog == "h2" {

			if carbonToOffset%10 == 0 {

				packs := carbonToOffset / 10

				cost = packs * 990
				costeTotal = append(costeTotal, cost)

			}
		}

	}
	min := costeTotal[0] 

	for _, value := range costeTotal {

		if value < min {
			min = value 
		}
	}

	return min
    
    
}
