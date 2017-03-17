package main

import (
	"github.com/firelyu/go_sample/encode"
	"fmt"
)

func main() {
	raw := "common"
	code := encode.EncodeBase64([]byte(raw))
	fmt.Println(string(code))

	c := "N2RhNGRlMTFhZjY0MzJhN2RmNTRiM2FiY2UxNzczYzNlNzYyNWZmODp7InVzZXJuYW1lIjoiY29tbW9uIiwidWlkIjoiZmRkODUxMzQtMDFmZC00MGQzLWFiYWItYWYwMWRiYjc4MTg2IiwicHJpdmlsZWdlcyI6IndyaXRlIiwidG9rZW4iOiJkMzcxMjcyYmI0ZmM0ODcxYjQzMmU0NTllYmYzYTMwYzZmZGQ4NTEzNDgyODQ5ZTJjMDFmZDQwZDM1OWFhNGJmNWFiYWJhZjAxYjExYmQ5ZjdkYmI3ODE4NmU4ZGE5YmIyIiwicm9sZSI6Im5vcm1hbCIsImlzc3VlX3RpbWUiOiIxNDg5NzMyMTI4IiwiZXhwaXJlX3RpbWUiOiIxODAwIiwicmVmcmVzaF90b2tlbiI6ImJlODIzMGEzZjRlMzRiNjI5YzRjZTJkZDBjMjI4ZmYyNjUwMjBhZGU2MzFmNGJmNDg4NTM1ZDQifQ=="
	p, _ := encode.DecodeBase64([]byte(c))
	fmt.Println(string(p))


}