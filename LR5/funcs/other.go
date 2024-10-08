package funcs

func AreaForSlice(s []Shape) {
	for _, v := range s {
		v.Area()
	}
}
