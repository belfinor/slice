package slice

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
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

func TestEqual(t *testing.T) {

	tF := func(l1 interface{}, l2 interface{}, wait bool) {

		if Equal(l1, l2) != wait {
			t.Fatal(fmt.Sprintf("Equal failed on %v and %v", l1, l2))
		}
	}

	var list []string
	list = append(list, "1", "3", "5", "7")

	tF(nil, nil, true)
	tF(nil, []string{}, true)
	tF(nil, "", false)
	tF([]string{}, nil, true)
	tF(0, nil, false)
	tF([]string{}, []string{}, true)
	tF([]int{}, []int{}, true)
	tF([]int{1, 2, 3}, []int{1, 2, 3}, true)
	tF([]int{1, 2, 3}, []int{1, 2, 4}, false)
	tF([]int{1, 2, 3}, []int{1, 2, 3, 4}, false)
	tF([]string{"1", "3", "5", "7"}, []string{"1", "3", "5", "7"}, true)
	tF(list, []string{"1", "3", "5", "7"}, true)
	tF(list, []string{"1", "3", "6", "7"}, false)
}
