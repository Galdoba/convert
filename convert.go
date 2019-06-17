package convert

import "strconv"

//IntToStr - Int в String
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

//StrToInt - String в Int
func StrToInt(s string) int {
	in, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return in
}
