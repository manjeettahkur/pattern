package pattern

func Diamond(rows int) {
	var i, j int
	if rows < 0 {
		rows = 4
	}

	spaces := rows - 1
	stars := 1

	for i = 1; i < rows*2; i++ {
		for j = 1; j <= spaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= stars*2; j++ {
			fmt.Print("*")
		}
		fmt.Println("")

		if i < rows {
			spaces--
			stars++
		} else {
			spaces++
			stars--
		}
	}
}
