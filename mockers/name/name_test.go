package name_test

import (
	"strconv"
	"testing"

	"github.com/panjiesw/apimocker/mockers/name"
)

func TestFirstFemale(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(name.FirstFemale())
		})
	}
}

func TestFirstMale(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(name.FirstMale())
		})
	}
}

func TestSurename(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(name.Surename())
		})
	}
}

func TestFullname(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(name.Fullname())
		})
	}
}

func panicValue(fn func()) (recovered interface{}) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return
}
