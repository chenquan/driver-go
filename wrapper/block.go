package wrapper

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <taos.h>
*/
import "C"
import (
	"unsafe"
)

// TaosFetchRawBlock  int         taos_fetch_raw_block(TAOS_RES *res, int* numOfRows, void** pData);
func TaosFetchRawBlock(result unsafe.Pointer) (int, int, unsafe.Pointer) {
	var cSize int
	size := unsafe.Pointer(&cSize)
	var block unsafe.Pointer
	errCode := int(C.taos_fetch_raw_block(result, (*C.int)(size), &block))
	return cSize, errCode, block
}

// TaosWriteRawBlock DLL_EXPORT int           taos_write_raw_block(TAOS *taos, int numOfRows, char *pData, const char* tbname);
func TaosWriteRawBlock(conn unsafe.Pointer, numOfRows int, pData unsafe.Pointer, tableName string) int {
	cStr := C.CString(tableName)
	defer C.free(unsafe.Pointer(cStr))
	return int(C.taos_write_raw_block(conn, (C.int)(numOfRows), (*C.char)(pData), cStr))
}
