package util

func ArrayContains (array [6]string, value string) bool {
	for _,a := range array {
		if (a == value) {
			return true;
		}
	}

	return false;
}