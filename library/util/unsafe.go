package util

import "unsafe"

func Btos(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func Stob(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}
