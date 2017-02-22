package mockers

import (
	"github.com/panjiesw/apimocker/mockers/name"
)

func GenNameFirstFemale() string {
	return name.FirstFemale()
}

func GenNameFirstMale() string {
	return name.FirstMale()
}

func GenNameSure() string {
	return name.Surename()
}

func GenNameFull() string {
	return name.Fullname()
}
