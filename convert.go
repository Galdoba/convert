package convert

import "strconv"

//ItoS - Int в String
func ItoS(i int) string {
	return strconv.Itoa(i)
}

//StoI - String в Int
func StoI(s string) int {
	in, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return in
}
