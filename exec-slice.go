package main

import (
	"fmt"
	"reflect"
	"errors"
)

func doubleSlice(s []int) []int  {
	t := make([]int, len(s), (cap(s) + 1) * 2 )
	for i := range s {
		t[i] = s[i]
	}
	return t
}

func doubleSlice2(s interface{}) (interface{}, error) {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		err := errors.New("The interface is not a slice.")
		return nil, err
	}

	newLen := reflect.ValueOf(s).Len()
	newCap := (reflect.ValueOf(s).Cap() + 1) * 2
	t := make([]interface{}, newLen, newCap)

	// Get the element type
	elementType := reflect.TypeOf(s).Elem().Kind()
	for i := 0; i < reflect.ValueOf(s).Len(); i++  {
		switch elementType {
		case reflect.Int32:
			t[i] = int32(reflect.ValueOf(s).Index(i).Int())
		case reflect.Float32:
			t[i] = float32(reflect.ValueOf(s).Index(i).Float())
		}
	}

	return t, nil
}

func main() {
	//s1 := []int{10,11,12,13,14}
	//for i := range s1 {
	//	fmt.Println(i)
	//}

	s2 := []int{21, 22}
	fmt.Printf("%v, len(%d), cap(%d)\n", s2, len(s2), cap(s2))

	//s3 := doubleSlice(s2)
	//fmt.Printf("%v, len(%d), cap(%d)\n", s3, len(s3), cap(s3))

	//s4 := []int32{31, 32, 33, 34}

	s5:=[]float32{51, 52, 53, 54}
	s6, _ := doubleSlice2(s5)
	//s4 := doubleSlice2(s2)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Printf("%v, len(%d), cap(%d)\n", s5, len(s5), cap(s5))
	fmt.Printf("%v, len(%d), cap(%d)\n", s6, reflect.ValueOf(s6).Len(), reflect.ValueOf(s6).Cap())
}
