package ten_two

import "ten/ten_one"

func GetZero() int {
	return ten_one.GetOne() - 1
}

func GetOne() int {
	return 1
}

