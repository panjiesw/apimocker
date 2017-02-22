package template

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validateFn2IntArgs(ttag, tag string) ([]int, bool) {
	r := regexp.MustCompile(fmt.Sprintf(`%s\(([\d,\s]+)\)$`, ttag))
	if match := r.MatchString(tag); !match {
		return nil, false
	}
	ss := strings.Split(r.ReplaceAllString(tag, "$1"), ",")
	if len(ss) != 2 {
		return nil, false
	}
	i1, _ := strconv.Atoi(strings.TrimSpace(ss[0]))
	i2, _ := strconv.Atoi(strings.TrimSpace(ss[1]))
	return []int{i1, i2}, true
}
