package name

import (
	"fmt"

	"github.com/panjiesw/apimocker/mockers/utils"
)

func FirstFemale() string {
	return femaleNameList[utils.IntRange(0, len(femaleNameList))]
}

func FirstMale() string {
	return maleNameList[utils.IntRange(0, len(maleNameList))]
}

func Surename() string {
	return surenameList[utils.IntRange(0, len(surenameList))]
}

type nameFn func() string

var nameFns = []nameFn{
	FirstFemale,
	FirstMale,
}

func Fullname() string {
	i := utils.IntRange(0, len(nameFns))
	return fmt.Sprintf("%s %s", nameFns[i](), Surename())
}
