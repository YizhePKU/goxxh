package xxh

// #cgo LDFLAGS: -L${SRCDIR}/lib -lxxhash
// #include <stdlib.h>
// unsigned long long XXH3_64bits(const void* data, size_t len);
import "C"

func XXH3(data []byte) uint64 {
	// Copy data to a buffer on the heap and pass it to C.
	ptr := C.CBytes(data)
	len := C.ulong(len(data))
	hash := C.XXH3_64bits(ptr, len)

	// Release memory
	C.free(ptr)

	return uint64(hash)
}
