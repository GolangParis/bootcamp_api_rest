package main

func AddtoSet(badges []Badge, new_badge Badge) []Badge {
	for _, it := range badges {
		if it.Name == new_badge.Name {
			return badges
		}
	}

	return append(badges, new_badge)
}
