package tasscommon

func truncate(s string, l int) string {
	runes := []rune(s)
	if len(runes) <= l {
		return s
	}
	return string(runes[0:l])
}
