package repo

func CalculateSkip(page, size int64) int64 {
	if page < 1 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	return (page - 1) * size
}
