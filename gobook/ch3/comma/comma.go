package comma

// comma は負でない 10 進表記整数文字列にカンマを挿入します。
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
