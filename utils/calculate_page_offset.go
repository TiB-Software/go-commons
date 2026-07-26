package utils

func CalculateOffset(page int32, limit int32) int32 {
	return limit * (page - 1)
}
