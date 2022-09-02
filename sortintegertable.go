package challenge_go

func SortIntegerTable(table []int) {

	size := len(table)
	BoardTemps := make([]int, size)
	CheckExist := make([]int, size)
	copy(BoardTemps, table)
	Id := 0
	for j := 0; j < size; j++ {
		Temps := int(10e16)
		for i := 0; i < size; i++ {
			if BoardTemps[i] <= Temps && CheckExist[i] == 0 {
				Temps = BoardTemps[i]
				Id = i
			}
		}
		table[j] = BoardTemps[Id]
		CheckExist[Id] = 1
	}
}
