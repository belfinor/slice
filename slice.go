package slice

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2019-10-01

import (
	"math/rand"
	"reflect"
	"time"
)

func Shuffle(slice interface{}) {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()

	for i := 0; i < length; i++ {
		j := rnd.Intn(length)
		if i != j {
			swap(i, j)
		}
	}
}

func Reverse(slice interface{}) {

	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	j := length - 1

	for i := 0; i < length/2; i++ {
		swap(i, j)
		j--
	}
}

func Contain(list interface{}, obj interface{}) (bool, int) {

	if list == nil {
		return false, -1
	}

	if reflect.TypeOf(list).Kind() == reflect.Slice || reflect.TypeOf(list).Kind() == reflect.Array {

		listvalue := reflect.ValueOf(list)

		for i := 0; i < listvalue.Len(); i++ {
			if obj == listvalue.Index(i).Interface() {
				return true, i
			}
		}
	}

	return false, -1
}
