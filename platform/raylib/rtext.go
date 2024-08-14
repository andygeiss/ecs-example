package rl

/*
#cgo CFLAGS: -Iinclude
#include "raylib.h"
#include "raymath.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func DrawText(text string, posX, posY, fontSize int32, color Color) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	cPosX := (C.int)(posX)
	cPosY := (C.int)(posY)
	cFontSize := (C.int)(fontSize)
	cColor := ptr[Color, C.Color](color)
	C.DrawText(cText, cPosX, cPosY, cFontSize, *cColor)
}
