package util

func CompareInt(mode string, int1 int, int2 int) int {
	switch mode {
	case "max":
		if int1 < int2 {
			return int2
		} else {
			return int1
		}
	case "min":
		if int1 < int2 {
			return int1
		} else {
			return int2
		}
	default:
		panic("Unknown mode")
	}
}
