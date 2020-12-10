package answer

import "regexp"

func isNumber(s string) bool {
	reg, _ := regexp.Compile(`^(\s)*(-|\+)?((\d+\.?)|(\d+e(-|\+)?\d+)|(\d*\.\d+)|(\d*\.\d+e(-|\+)?\d+)|(\d+\.\d*e(-|\+)?\d+))(\s)*$`)
	ans := reg.MatchString(s)
	return ans
}
