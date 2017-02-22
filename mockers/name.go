package mockers

import (
	"github.com/panjiesw/apimocker/mockers/name"
)

// GenNameFirstFemale generate random US popular female first name
func GenNameFirstFemale() string {
	return name.FirstFemale()
}

// GenNameFirstMale generate random US popular male first name
func GenNameFirstMale() string {
	return name.FirstMale()
}

// GenNameSure generate random US popular surename
func GenNameSure() string {
	return name.Surename()
}

// GenNameFull generate random fullname from GenNameFirstFemale/GenNameFirstMale + GenNameSure
func GenNameFull() string {
	return name.Fullname()
}
