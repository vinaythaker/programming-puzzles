package puzzles

func spiralorder(matrix [][]int) []int {

	i := 0
	j := 0

	n := len(matrix)
	if n == 0 {
		return make([]int, 0)
	}

	m := len(matrix[0])
	if m == 0 {
		m = 1
	}

	count := 0
	lr_direction := 1
	ud_direction := 0
	l_b := 0
	r_b := m
	u_b := 0
	d_b := n
	output := make([]int, n*m)

	for {

		if lr_direction == 1 && ud_direction == 0 && j <= r_b {
			// going left
			output[count] = matrix[i][j]
			j++
			if j == r_b {
				ud_direction = 1
				lr_direction = 0
				r_b--
				j--
				i++
			}
		} else if lr_direction == -1 && ud_direction == 0 && j >= l_b {
			//going right
			output[count] = matrix[i][j]
			j--
			if j < l_b {
				ud_direction = -1
				lr_direction = 0
				l_b++
				j++
				i--
			}
		} else if lr_direction == 0 && ud_direction == 1 && i <= d_b {
			// going down
			output[count] = matrix[i][j]
			i++
			if i == d_b {
				lr_direction = -1
				ud_direction = 0
				d_b--
				i--
				j--
			}
		} else if lr_direction == 0 && ud_direction == -1 && i >= u_b {
			//going up
			output[count] = matrix[i][j]
			i--
			if i == u_b {
				lr_direction = 1
				ud_direction = 0
				u_b++
				i++
				j++
			}
		}

		count++
		if count >= m*n {
			break
		}
	}

	return output
}
