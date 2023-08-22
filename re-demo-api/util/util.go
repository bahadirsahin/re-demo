package util

// for given items, get smallest pack possible from various pack sizes
func GetSmallestPack(items int, packSizes []int) int {
	if items == 0 || len(packSizes) == 0 {
		return 0
	}

	for i := 0; i < len(packSizes); i++ {
		if items <= packSizes[i] {
			return packSizes[i]
		}
	}

	return packSizes[len(packSizes)-1]
}
