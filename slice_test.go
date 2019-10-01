package slice

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-10-01

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	orig := []int{1, 2, 3, 4, 5}

	Reverse(orig)

	if len(orig) != 5 {
		t.Fatal("slice.Reverse corrupt data")
	}

	for i, v := range orig {
		if v != 5-i {
			t.Fatal("slice.Reverse not work")
		}
	}

	Shuffle(orig)

	if len(orig) != 5 {
		t.Fatal("slice.Shuffle corrupt data")
	}
}

func TestContain(t *testing.T) {

	tF := func(list interface{}, obj interface{}, has bool, place int) {

		rh, rp := Contain(list, obj)

		if rh != has {
			t.Fatal(fmt.Sprintf("Wrong result slice %v obj %v", list, obj))
		}

		if rp != place {
			t.Fatal(fmt.Sprintf("Wrong index slice %v obj %v", list, obj))
		}
	}

	tF([]string{"1", "2", "3", "10", "12", "13"}, "2", true, 1)
	tF([]string{"1", "2", "3", "10", "12", "13"}, "9", false, -1)
	tF([]string{"1", "2", "3", "10", "12", "13"}, "12", true, 4)

	var list []string

	tF(list, "12", false, -1)
	tF(nil, "12", false, -1)

	tF([]int{1, 2, 3, 4, 5, 6}, 4, true, 3)
	tF([]int{1, 2, 3, 4, 5, 6}, 0, false, -1)
}
