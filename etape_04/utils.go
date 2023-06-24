package main

func AddToSet(vv []int, val int) []int {
	for _, v := range vv {
		if v == val {
			return vv
		}
	}

	return append(vv, val)
}

func RemoveFromSet(vv []int, val int) []int {
	pos := 0
	for _, v := range vv {
		if v == val {
			return append(vv[:pos], vv[pos:]...)
		}
		pos++
	}

	return vv
}
