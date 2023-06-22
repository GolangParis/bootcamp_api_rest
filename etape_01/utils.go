package main

func AddtoSet(badges []Badge, new_badge Badge) []Badge {
	for _, it := range badges {
		if it == new_badge {
			return badges
		}
	}

	return append(badges, new_badge)
}
