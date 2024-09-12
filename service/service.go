package service

func ValidateCard(num string) bool {
	var count int
	for index, value := range num {
		valueInt := int(value - '0')
		if index%2 == 0 {
			a := valueInt * 2
			if a > 9 {
				a -= 9
			}
			count += a
		} else {
			count += valueInt
		}
	}
	return count%10 == 0
}
