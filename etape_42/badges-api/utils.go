package main

// AddToSet permet de n'ajouter une valeur à une slice 
// seulement si cette dernière ne la contient pas déja
func AddToSet(vals []int, newV int) []int {
	for _, v := range vals {
		if v == newV {
			return vals
		}
	}
	vals = append(vals, newV)
	return vals
}
