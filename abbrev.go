package stringx

func Abbrev(str string, maxs ...int) string {
	max := 10
	if len(maxs) > 0 {
		max = maxs[0]
	}
	if len(str) > max {
		str = str[:max]
		str = str + "..."
	}
	return str

}
