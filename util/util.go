package util

import "strconv"


func StringToInt(str string) int{
	if str == "" {
		return 0
	}
	//convert string to int
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}


func IntToString(i int) string{
	//convert int to string
	str := strconv.Itoa(i)
	return str
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func BoolPtr(b bool) *bool {
	return &b
}