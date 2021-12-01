package services

import (
	"apiGo/collections"
	"reflect"
)

func Comparation(a, n collections.Product) collections.Product {
	c := collections.Product{}
	aType := reflect.ValueOf(a)
	cType := reflect.ValueOf(c)
	nType := reflect.ValueOf(n)
	rType := reflect.TypeOf(a)

	newP := reflect.New(rType)
	for i := 0; i < aType.NumField(); i++ {
		nValue := aType.Field(i)
		ceroValue := cType.Field(i)
		currentValue := nType.Field(i)
		if i != 4 {
			if nValue.Interface() != ceroValue.Interface() && nValue.Interface() != currentValue.Interface() {
				newP.Elem().Field(i).Set(nValue)
			} else {
				newP.Elem().Field(i).Set(currentValue)
			}
		}

	}
	g3 := newP.Interface().(*collections.Product)
	g3.Tags = n.Tags
	return *g3
}
