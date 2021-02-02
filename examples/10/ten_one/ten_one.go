package ten_one

import "ten/ten_two"

func GetZero() int {
	return ten_two.GetOne() - 1
}

func GetOne() int {
	return 1
}