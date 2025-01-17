package av2bv

import (
	"unsafe"
)

// slice is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
//
// Unlike reflect.SliceHeader, its Data field is sufficient to guarantee the
// data it references will not be garbage collected.
type slice struct {
	data unsafe.Pointer
	len  int
	cap  int
}

// bytesToString 没有内存开销的转换
func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// stringToBytes 没有内存开销的转换
func stringToBytes(s string) (b []byte) {
	bh := (*slice)(unsafe.Pointer(&b))
	sh := (*slice)(unsafe.Pointer(&s)) // 不要访问 sh.cap
	bh.data = sh.data
	bh.len = sh.len
	bh.cap = sh.len
	return b
}
