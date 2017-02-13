package main

import (
	"fmt"
	"reflect"
	"errors"
	"regexp"
	"io/ioutil"
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


// http://stackoverflow.com/questions/42151307/how-to-determine-the-element-type-of-slice-interface#answer-42151765
// Call reflect.MakeSlice() and reflect.SliceOf()
func doubleSlice3(s interface{}) (interface{}, error) {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		err := errors.New("The interface is not a slice.")
		return nil, err
	}

	v := reflect.ValueOf(s)
	newLength := v.Len()
	newCapacity := (v.Cap() + 1) * 2
	elementType := reflect.TypeOf(s).Elem()

	t := reflect.MakeSlice(reflect.SliceOf(elementType), newLength, newCapacity)
	reflect.Copy(t, v)
	return t.Interface(), nil
}

const digitFormat = "[0-9]+"

func CopyDigit(filename string) ([]byte, error) {
	digitRegexp := regexp.MustCompile(digitFormat)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m := digitRegexp.Find(content)

	var p []byte
	for _, v := range m {
		p = append(p, v)
	}

	return p, nil
}

func CopyAllDigit(file string) ([][]byte, error)  {
	digitRegexp := regexp.MustCompile(digitFormat)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	matches := digitRegexp.FindAll(content, 10)
	var m [][]byte
	for _, v1 := range matches {
		m1 := make([]byte, len(v1))
		copy(m1, v1)
		m = append(m, m1)
	}

	return m, nil
}

// https://blog.golang.org/slices
type sliceHeader struct {
	Length int
	Capacity int
	ZeroerElement *byte
}

func AddOneToEachElement(s []byte) {
	for i := range s {
		s[i]++
	}
}

func SubtractOneFromLength(s []byte) []byte {
	return s[:len(s) - 1]
}

func PtrSubtractOneFromLength(ps *[]byte) {
	slices := *ps
	*ps = slices[:len(slices) - 1]
}

func main() {
	//s1 := []int{10,11,12,13,14}
	//for i := range s1 {
	//	fmt.Println(i)
	//}

	//s2 := []int{21, 22}
	//fmt.Printf("%v, len(%d), cap(%d)\n", s2, len(s2), cap(s2))

	//s3 := doubleSlice(s2)
	//fmt.Printf("%v, len(%d), cap(%d)\n", s3, len(s3), cap(s3))

	//s4 := []int32{31, 32, 33, 34}

	//s5:=[]float32{51, 52, 53, 54}
	//s6, _ := doubleSlice3(s5)
	////s4 := doubleSlice2(s2)
	//fmt.Println(s5)
	//fmt.Println(s6)
	//fmt.Printf("%v, len(%d), cap(%d)\n", s5, len(s5), cap(s5))
	//fmt.Printf("%v, len(%d), cap(%d)\n", s6, reflect.ValueOf(s6).Len(), reflect.ValueOf(s6).Cap())

	//file := "numbers.txt"
	//p, _ := CopyDigit(file)
	//fmt.Printf("CopyDigit(%s) is %s\n", file, p)
	//
	//m, _ := CopyAllDigit(file)
	//fmt.Printf("CopyAllDigit(%s) is %s\n", file, m)

	var buf [256]byte
	s := buf[100:110]
	for i, _ := range s{
		s[i] = byte(i)
	}
	fmt.Println("Init slice", s)
	//AddOneToEachElement(s)
	//fmt.Println("Add one for each element", s)

	fmt.Println("Before: len(slice) =", len(s))
	PtrSubtractOneFromLength(&s)
	fmt.Println("After:  len(slice) =", len(s))
	//fmt.Println("After:  len(newSlice) =", len(newSlice))
}
