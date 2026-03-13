package example1

func CheckDivisibility(input int) string {
	if input % 3 == 0 {
		return "Three!"
	}
	if input % 5 == 0 {
		return "Five!"
	}
	if input % 2 == 0 {
		return "Two"
	}
	return "invalid number"
}