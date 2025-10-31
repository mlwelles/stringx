package stringx

func FormatCurrency(v string) string {
	r := []rune(Reverse(v))
	out := ""
	for i := 0; i < len(r); i = i + 1 {
		if i > 0 && i%3 == 0 {
			out += ","
		}
		out += string(r[i])
	}
	return "$" + Reverse(out)
}
