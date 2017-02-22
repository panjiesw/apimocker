package name

import "testing"
import "strconv"

func TestFirstFemale(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(FirstFemale())
		})
	}
}

func TestFirstMale(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(FirstMale())
		})
	}
}

func TestSurename(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(Surename())
		})
	}
}

func TestFullname(t *testing.T) {
	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(Fullname())
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
