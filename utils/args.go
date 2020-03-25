package utils

import "os"

func HaveArg(value string) bool {
	osArgs := os.Args[1:]

	for _, v := range osArgs {
		if value == v {
			return true
		}

		if len(v) >= 2 {
			if value == v[1:] {
				return true
			}

			if value == v[2:] {
				return true
			}
		}
	}
	return false
}
