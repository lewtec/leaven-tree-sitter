package grammar_astro

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type anon_0 struct {
	F0 unsafe.Pointer
	F1 unsafe.Pointer
	F2 unsafe.Pointer
	F3 unsafe.Pointer
	F4 unsafe.Pointer
	F5 unsafe.Pointer
	F6 unsafe.Pointer
}
type TSSymbolMetadata struct {
	F0 byte
	F1 byte
	F2 byte
}
type TSLexMode struct {
	F0 int16
	F1 int16
}
type String struct {
	F0 unsafe.Pointer
	F1 int32
	F2 int32
}
type TagMapEntry struct {
	F0 [16]byte
	F1 int32
}
type TSParseActionEntry struct {
	F0 TSParseAction
}
type TSParseAction struct {
	F0 anon_2
}
type anon_2 struct {
	F0 byte
	F1 byte
	F2 int16
	F3 int16
	F4 int16
}
type anon_3 struct {
	F0 byte
	F1 byte
}
type Tag struct {
	F0 int32
	F1 String
}
type Scanner struct {
	F0 struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}
}
type TSLexer struct {
	F0 int32
	F1 int16
	F2 unsafe.Pointer
	F3 unsafe.Pointer
	F4 unsafe.Pointer
	F5 unsafe.Pointer
	F6 unsafe.Pointer
}
type Array struct {
	F0 unsafe.Pointer
	F1 int32
	F2 int32
}

var tree_sitter_astro_language struct {
	F0  int32
	F1  int32
	F2  int32
	F3  int32
	F4  int32
	F5  int32
	F6  int32
	F7  int32
	F8  int32
	F9  int16
	F10 [2]byte
	F11 unsafe.Pointer
	F12 unsafe.Pointer
	F13 unsafe.Pointer
	F14 unsafe.Pointer
	F15 unsafe.Pointer
	F16 unsafe.Pointer
	F17 unsafe.Pointer
	F18 unsafe.Pointer
	F19 unsafe.Pointer
	F20 unsafe.Pointer
	F21 unsafe.Pointer
	F22 unsafe.Pointer
	F23 unsafe.Pointer
	F24 unsafe.Pointer
	F25 unsafe.Pointer
	F26 int16
	F27 [6]byte
	F28 anon_0
	F29 unsafe.Pointer
}
var ts_small_parse_table [2812]int16 = [2812]int16{15, 3, 1, 28, 19, 1, 1, 21, 1, 5, 23, 1, 7, 25, 1, 16, 27, 1, 26, 29, 1, 29, 5, 1, 42, 82, 1, 55, 83, 1, 54, 85, 1, 45, 91, 1, 46, 116, 1, 43, 117, 1, 44, 10, 8, 37, 38, 39, 40, 41, 47, 53, 56, 15, 3, 1, 28, 19, 1, 1, 21, 1, 5, 25, 1, 16, 29, 1, 29, 31, 1, 7, 33, 1, 26, 5, 1, 42, 71, 1, 46, 82, 1, 55, 83, 1, 54, 85, 1, 45, 116, 1, 43, 117, 1, 44, 10, 8, 37, 38, 39, 40, 41, 47, 53, 56, 15, 3, 1, 28, 19, 1, 1, 21, 1, 5, 29, 1, 29, 35, 1, 7, 37, 1, 16, 39, 1, 26, 5, 1, 42, 46, 1, 46, 82, 1, 55, 83, 1, 54, 85, 1, 45, 116, 1, 43, 117, 1, 44, 7, 8, 37, 38, 39, 40, 41, 47, 53, 56, 15, 3, 1, 28, 19, 1, 1, 21, 1, 5, 29, 1, 29, 31, 1, 7, 41, 1, 16, 43, 1, 26, 5, 1, 42, 80, 1, 46, 82, 1, 55, 83, 1, 54, 85, 1, 45, 116, 1, 43, 117, 1, 44, 3, 8, 37, 38, 39, 40, 41, 47, 53, 56, 15, 3, 1, 28, 19, 1, 1, 21, 1, 5, 23, 1, 7, 29, 1, 29, 45, 1, 16, 47, 1, 26, 5, 1, 42, 82, 1, 55, 83, 1, 54, 85, 1, 45, 102, 1, 46, 116, 1, 43, 117, 1, 44, 2, 8, 37, 38, 39, 40, 41, 47, 53, 56, 15, 3, 1, 28, 19, 1, 1, 21, 1, 5, 25, 1, 16, 29, 1, 29, 35, 1, 7, 49, 1, 26, 5, 1, 42, 57, 1, 46, 82, 1, 55, 83, 1, 54, 85, 1, 45, 116, 1, 43, 117, 1, 44, 10, 8, 37, 38, 39, 40, 41, 47, 53, 56, 14, 3, 1, 28, 51, 1, 0, 53, 1, 1, 56, 1, 5, 59, 1, 7, 62, 1, 16, 65, 1, 29, 4, 1, 42, 64, 1, 45, 66, 1, 54, 74, 1, 55, 118, 1, 44, 119, 1, 43, 8, 8, 37, 38, 39, 40, 41, 47, 53, 56, 14, 3, 1, 28, 7, 1, 1, 9, 1, 5, 11, 1, 7, 17, 1, 29, 68, 1, 0, 70, 1, 16, 4, 1, 42, 64, 1, 45, 66, 1, 54, 74, 1, 55, 118, 1, 44, 119, 1, 43, 8, 8, 37, 38, 39, 40, 41, 47, 53, 56, 14, 3, 1, 28, 51, 1, 26, 72, 1, 1, 75, 1, 5, 78, 1, 7, 81, 1, 16, 84, 1, 29, 5, 1, 42, 82, 1, 55, 83, 1, 54, 85, 1, 45, 116, 1, 43, 117, 1, 44, 10, 8, 37, 38, 39, 40, 41, 47, 53, 56, 14, 3, 1, 28, 7, 1, 1, 9, 1, 5, 11, 1, 7, 17, 1, 29, 87, 1, 0, 89, 1, 16, 4, 1, 42, 64, 1, 45, 66, 1, 54, 74, 1, 55, 118, 1, 44, 119, 1, 43, 9, 8, 37, 38, 39, 40, 41, 47, 53, 56, 14, 3, 1, 28, 7, 1, 1, 9, 1, 5, 11, 1, 7, 17, 1, 29, 70, 1, 16, 87, 1, 0, 4, 1, 42, 64, 1, 45, 66, 1, 54, 74, 1, 55, 118, 1, 44, 119, 1, 43, 8, 8, 37, 38, 39, 40, 41, 47, 53, 56, 13, 3, 1, 28, 91, 1, 1, 93, 1, 5, 95, 1, 29, 97, 1, 30, 99, 1, 34, 6, 1, 42, 87, 1, 55, 88, 1, 54, 96, 1, 45, 121, 1, 43, 122, 1, 44, 18, 7, 37, 39, 40, 41, 50, 53, 58, 13, 3, 1, 28, 101, 1, 1, 104, 1, 5, 107, 1, 29, 110, 1, 30, 112, 1, 34, 6, 1, 42, 87, 1, 55, 88, 1, 54, 96, 1, 45, 121, 1, 43, 122, 1, 44, 14, 7, 37, 39, 40, 41, 50, 53, 58, 13, 3, 1, 28, 91, 1, 1, 93, 1, 5, 95, 1, 29, 115, 1, 30, 117, 1, 34, 6, 1, 42, 87, 1, 55, 88, 1, 54, 96, 1, 45, 121, 1, 43, 122, 1, 44, 19, 7, 37, 39, 40, 41, 50, 53, 58, 13, 3, 1, 28, 91, 1, 1, 93, 1, 5, 95, 1, 29, 119, 1, 30, 121, 1, 34, 6, 1, 42, 87, 1, 55, 88, 1, 54, 96, 1, 45, 121, 1, 43, 122, 1, 44, 14, 7, 37, 39, 40, 41, 50, 53, 58, 13, 3, 1, 28, 91, 1, 1, 93, 1, 5, 95, 1, 29, 123, 1, 30, 125, 1, 34, 6, 1, 42, 87, 1, 55, 88, 1, 54, 96, 1, 45, 121, 1, 43, 122, 1, 44, 16, 7, 37, 39, 40, 41, 50, 53, 58, 13, 3, 1, 28, 91, 1, 1, 93, 1, 5, 95, 1, 29, 121, 1, 34, 127, 1, 30, 6, 1, 42, 87, 1, 55, 88, 1, 54, 96, 1, 45, 121, 1, 43, 122, 1, 44, 14, 7, 37, 39, 40, 41, 50, 53, 58, 13, 3, 1, 28, 91, 1, 1, 93, 1, 5, 95, 1, 29, 121, 1, 34, 129, 1, 30, 6, 1, 42, 87, 1, 55, 88, 1, 54, 96, 1, 45, 121, 1, 43, 122, 1, 44, 14, 7, 37, 39, 40, 41, 50, 53, 58, 7, 3, 1, 28, 131, 1, 3, 133, 1, 6, 135, 1, 9, 137, 1, 19, 113, 1, 52, 35, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 139, 1, 3, 141, 1, 6, 113, 1, 52, 32, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 143, 1, 3, 145, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 147, 1, 3, 149, 1, 6, 113, 1, 52, 33, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 139, 1, 3, 151, 1, 6, 113, 1, 52, 30, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 147, 1, 3, 153, 1, 6, 113, 1, 52, 29, 2, 48, 57, 7, 3, 1, 28, 131, 1, 3, 135, 1, 9, 137, 1, 19, 155, 1, 6, 113, 1, 52, 22, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 157, 1, 3, 159, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 131, 1, 3, 135, 1, 9, 137, 1, 19, 161, 1, 6, 113, 1, 52, 38, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 163, 1, 3, 165, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 157, 1, 3, 167, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 147, 1, 3, 169, 1, 6, 113, 1, 52, 37, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 157, 1, 3, 171, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 163, 1, 3, 173, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 139, 1, 3, 175, 1, 6, 113, 1, 52, 27, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 143, 1, 3, 177, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 137, 1, 19, 179, 1, 10, 181, 1, 12, 183, 1, 14, 185, 1, 33, 111, 2, 49, 52, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 163, 1, 3, 187, 1, 6, 113, 1, 52, 39, 2, 48, 57, 7, 3, 1, 28, 135, 1, 9, 137, 1, 19, 143, 1, 3, 189, 1, 6, 113, 1, 52, 39, 2, 48, 57, 6, 3, 1, 28, 193, 1, 9, 196, 1, 19, 113, 1, 52, 191, 2, 3, 6, 39, 2, 48, 57, 3, 3, 1, 28, 201, 1, 5, 199, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 205, 1, 5, 203, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 209, 1, 5, 207, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 213, 1, 5, 211, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 217, 1, 5, 215, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 221, 1, 5, 219, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 225, 1, 5, 223, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 205, 1, 5, 203, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 229, 1, 5, 227, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 233, 1, 5, 231, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 237, 1, 5, 235, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 241, 1, 5, 239, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 245, 1, 5, 243, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 249, 1, 5, 247, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 253, 1, 5, 251, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 257, 1, 5, 255, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 261, 1, 5, 259, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 265, 1, 5, 263, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 269, 1, 5, 267, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 273, 1, 5, 271, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 277, 1, 5, 275, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 277, 1, 5, 275, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 281, 1, 5, 279, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 257, 1, 5, 255, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 285, 1, 5, 283, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 241, 1, 5, 239, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 289, 1, 5, 287, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 233, 1, 5, 231, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 229, 1, 5, 227, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 273, 1, 5, 271, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 269, 1, 5, 267, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 265, 1, 5, 263, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 261, 1, 5, 259, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 245, 1, 5, 243, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 293, 1, 5, 291, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 297, 1, 5, 295, 5, 29, 0, 1, 7, 16, 3, 3, 1, 28, 237, 1, 5, 235, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 201, 1, 5, 199, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 213, 1, 5, 211, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 209, 1, 5, 207, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 225, 1, 5, 223, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 297, 1, 5, 295, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 293, 1, 5, 291, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 289, 1, 5, 287, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 301, 1, 5, 299, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 285, 1, 5, 283, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 249, 1, 5, 247, 5, 26, 29, 1, 7, 16, 3, 3, 1, 28, 293, 1, 5, 291, 4, 29, 30, 34, 1, 3, 3, 1, 28, 289, 1, 5, 287, 4, 29, 30, 34, 1, 3, 3, 1, 28, 201, 1, 5, 199, 4, 29, 30, 34, 1, 3, 3, 1, 28, 261, 1, 5, 259, 4, 29, 30, 34, 1, 3, 3, 1, 28, 265, 1, 5, 263, 4, 29, 30, 34, 1, 3, 3, 1, 28, 269, 1, 5, 267, 4, 29, 30, 34, 1, 3, 3, 1, 28, 273, 1, 5, 271, 4, 29, 30, 34, 1, 3, 3, 1, 28, 277, 1, 5, 275, 4, 29, 30, 34, 1, 3, 3, 1, 28, 257, 1, 5, 255, 4, 29, 30, 34, 1, 3, 3, 1, 28, 285, 1, 5, 283, 4, 29, 30, 34, 1, 3, 3, 1, 28, 233, 1, 5, 231, 4, 29, 30, 34, 1, 3, 3, 1, 28, 229, 1, 5, 227, 4, 29, 30, 34, 1, 3, 3, 1, 28, 213, 1, 5, 211, 4, 29, 30, 34, 1, 3, 3, 1, 28, 209, 1, 5, 207, 4, 29, 30, 34, 1, 3, 3, 1, 28, 205, 1, 5, 203, 4, 29, 30, 34, 1, 3, 3, 1, 28, 225, 1, 5, 223, 4, 29, 30, 34, 1, 3, 3, 1, 28, 297, 1, 5, 295, 4, 29, 30, 34, 1, 4, 3, 1, 28, 305, 1, 8, 307, 1, 9, 303, 3, 3, 6, 19, 3, 3, 1, 28, 245, 1, 5, 243, 4, 29, 30, 34, 1, 3, 3, 1, 28, 237, 1, 5, 235, 4, 29, 30, 34, 1, 3, 3, 1, 28, 241, 1, 5, 239, 4, 29, 30, 34, 1, 3, 3, 1, 28, 311, 1, 9, 309, 3, 3, 6, 19, 3, 3, 1, 28, 315, 1, 9, 313, 3, 3, 6, 19, 3, 3, 1, 28, 319, 1, 9, 317, 3, 3, 6, 19, 3, 3, 1, 28, 323, 1, 9, 321, 3, 3, 6, 19, 5, 3, 1, 28, 325, 1, 21, 327, 1, 22, 329, 1, 23, 331, 1, 35, 3, 3, 1, 28, 307, 1, 9, 303, 3, 3, 6, 19, 5, 3, 1, 28, 331, 1, 35, 333, 1, 21, 335, 1, 22, 337, 1, 23, 5, 3, 1, 28, 331, 1, 35, 339, 1, 21, 341, 1, 22, 343, 1, 23, 4, 3, 1, 28, 345, 1, 7, 347, 1, 27, 41, 1, 46, 4, 3, 1, 28, 345, 1, 7, 349, 1, 27, 79, 1, 46, 4, 3, 1, 28, 351, 1, 7, 353, 1, 27, 42, 1, 46, 4, 3, 1, 28, 351, 1, 7, 355, 1, 27, 47, 1, 46, 4, 3, 1, 28, 357, 1, 24, 359, 1, 25, 361, 1, 35, 4, 3, 1, 28, 363, 1, 7, 365, 1, 27, 101, 1, 46, 4, 3, 1, 28, 363, 1, 7, 367, 1, 27, 100, 1, 46, 4, 3, 1, 28, 359, 1, 25, 369, 1, 24, 371, 1, 35, 4, 3, 1, 28, 359, 1, 25, 373, 1, 24, 375, 1, 35, 3, 3, 1, 28, 377, 1, 12, 379, 1, 13, 3, 3, 1, 28, 377, 1, 14, 381, 1, 15, 2, 3, 1, 28, 383, 2, 27, 7, 3, 3, 1, 28, 351, 1, 7, 58, 1, 46, 2, 3, 1, 28, 385, 2, 27, 7, 3, 3, 1, 28, 363, 1, 7, 92, 1, 46, 3, 3, 1, 28, 373, 1, 24, 375, 1, 35, 3, 3, 1, 28, 363, 1, 7, 93, 1, 46, 3, 3, 1, 28, 387, 1, 18, 389, 1, 31, 3, 3, 1, 28, 345, 1, 7, 69, 1, 46, 3, 3, 1, 28, 369, 1, 24, 371, 1, 35, 3, 3, 1, 28, 357, 1, 24, 361, 1, 35, 3, 3, 1, 28, 345, 1, 7, 70, 1, 46, 2, 3, 1, 28, 391, 2, 27, 7, 2, 3, 1, 28, 393, 2, 27, 7, 3, 3, 1, 28, 351, 1, 7, 59, 1, 46, 2, 3, 1, 28, 395, 1, 3, 2, 3, 1, 28, 397, 1, 3, 2, 3, 1, 28, 399, 1, 12, 2, 3, 1, 28, 399, 1, 14, 2, 3, 1, 28, 401, 1, 20, 2, 3, 1, 28, 403, 1, 3, 2, 3, 1, 28, 405, 1, 0, 2, 3, 1, 28, 407, 1, 4, 2, 3, 1, 28, 409, 1, 3, 2, 3, 1, 28, 411, 1, 2, 2, 3, 1, 28, 413, 1, 3, 2, 3, 1, 28, 415, 1, 3, 2, 3, 1, 28, 417, 1, 25, 2, 3, 1, 28, 419, 1, 3, 2, 3, 1, 28, 421, 1, 3, 2, 3, 1, 28, 423, 1, 18, 2, 3, 1, 28, 359, 1, 25, 2, 3, 1, 28, 425, 1, 2, 2, 3, 1, 28, 427, 1, 32, 2, 3, 1, 28, 429, 1, 2, 2, 3, 1, 28, 431, 1, 4, 2, 3, 1, 28, 433, 1, 4}
var ts_small_parse_table_map [161]int32 = [161]int32{0, 53, 106, 159, 212, 265, 318, 368, 418, 468, 518, 568, 614, 660, 706, 752, 798, 844, 890, 913, 936, 959, 982, 1005, 1028, 1051, 1074, 1097, 1120, 1143, 1166, 1189, 1212, 1235, 1258, 1281, 1304, 1327, 1348, 1362, 1376, 1390, 1404, 1418, 1432, 1446, 1460, 1474, 1488, 1502, 1516, 1530, 1544, 1558, 1572, 1586, 1600, 1614, 1628, 1642, 1656, 1670, 1684, 1698, 1712, 1726, 1740, 1754, 1768, 1782, 1796, 1810, 1824, 1838, 1852, 1866, 1880, 1894, 1908, 1922, 1936, 1950, 1964, 1978, 1992, 2006, 2019, 2032, 2045, 2058, 2071, 2084, 2097, 2110, 2123, 2136, 2149, 2162, 2175, 2188, 2201, 2214, 2227, 2242, 2255, 2268, 2281, 2293, 2305, 2317, 2329, 2345, 2357, 2373, 2389, 2402, 2415, 2428, 2441, 2454, 2467, 2480, 2493, 2506, 2516, 2526, 2534, 2544, 2552, 2562, 2572, 2582, 2592, 2602, 2612, 2622, 2632, 2640, 2648, 2658, 2665, 2672, 2679, 2686, 2693, 2700, 2707, 2714, 2721, 2728, 2735, 2742, 2749, 2756, 2763, 2770, 2777, 2784, 2791, 2798, 2805}
var ts_symbol_names [59]unsafe.Pointer = [59]unsafe.Pointer{libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_17), libc.Ptr(&_str_20), libc.Ptr(&_str_17), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_25), libc.Ptr(&_str_25), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_30), libc.Ptr(&_str_31), libc.Ptr(&_str_32), libc.Ptr(&_str_33), libc.Ptr(&_str_10), libc.Ptr(&_str_34), libc.Ptr(&_str_11), libc.Ptr(&_str_35), libc.Ptr(&_str_36), libc.Ptr(&_str_37), libc.Ptr(&_str_38), libc.Ptr(&_str_39), libc.Ptr(&_str_39), libc.Ptr(&_str_39), libc.Ptr(&_str_40), libc.Ptr(&_str_41), libc.Ptr(&_str_42), libc.Ptr(&_str_43), libc.Ptr(&_str_44), libc.Ptr(&_str_45), libc.Ptr(&_str_46), libc.Ptr(&_str_47), libc.Ptr(&_str_48), libc.Ptr(&_str_40), libc.Ptr(&_str_40), libc.Ptr(&_str_49), libc.Ptr(&_str_50), libc.Ptr(&_str_51)}
var ts_symbol_metadata [59]TSSymbolMetadata = [59]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [59]int16 = [59]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 10, 14, 10, 16, 17, 17, 19, 20, 21, 21, 21, 21, 25, 26, 27, 28, 19, 20, 31, 32, 33, 34, 3, 36, 37, 38, 39, 40, 41, 42, 42, 42, 45, 46, 47, 48, 49, 50, 51, 52, 53, 45, 45, 56, 57, 58}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][4]int16 = [1][4]int16{}
var ts_lex_modes [163]TSLexMode = [163]TSLexMode{TSLexMode{0, 1}, TSLexMode{22, 2}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{1, 6}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 2}, TSLexMode{23, 2}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{23, 3}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{10, 5}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{0, 4}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{10, 5}, TSLexMode{0, 7}, TSLexMode{10, 5}, TSLexMode{0, 7}, TSLexMode{0, 7}, TSLexMode{0, 8}, TSLexMode{0, 8}, TSLexMode{0, 8}, TSLexMode{0, 8}, TSLexMode{0, 9}, TSLexMode{0, 8}, TSLexMode{0, 8}, TSLexMode{0, 9}, TSLexMode{0, 9}, TSLexMode{4, 10}, TSLexMode{2, 10}, TSLexMode{0, 8}, TSLexMode{0, 10}, TSLexMode{0, 8}, TSLexMode{0, 10}, TSLexMode{0, 11}, TSLexMode{0, 10}, TSLexMode{9, 12}, TSLexMode{0, 10}, TSLexMode{0, 11}, TSLexMode{0, 11}, TSLexMode{0, 10}, TSLexMode{0, 8}, TSLexMode{0, 8}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{20, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{0, 13}, TSLexMode{0, 10}, TSLexMode{0, 10}, TSLexMode{9, 10}, TSLexMode{0, 13}, TSLexMode{20, 10}, TSLexMode{0, 14}, TSLexMode{20, 10}, TSLexMode{0, 10}, TSLexMode{0, 10}}
var ts_external_scanner_states [15][16]byte = [15][16]byte{[16]byte{}, [16]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0}, [16]byte{0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0}, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0}, [16]byte{0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, [16]byte{1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1}, [16]byte{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0}, [16]byte{0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1}, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}, [16]byte{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1}, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0}, [16]byte{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}}
var ts_external_scanner_symbol_map [16]int16 = [16]int16{21, 22, 23, 24, 25, 6, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35}
var ts_primary_state_ids [163]int16 = [163]int16{0, 1, 2, 2, 4, 4, 4, 2, 8, 9, 8, 11, 12, 13, 14, 13, 16, 13, 16, 16, 20, 21, 22, 23, 21, 23, 20, 27, 20, 29, 27, 23, 27, 29, 21, 22, 36, 29, 22, 39, 40, 41, 42, 43, 44, 45, 46, 41, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 60, 62, 55, 64, 51, 66, 49, 48, 59, 58, 57, 56, 52, 74, 75, 50, 40, 43, 42, 46, 75, 74, 66, 84, 64, 53, 74, 66, 40, 56, 57, 58, 59, 60, 55, 64, 49, 48, 43, 42, 41, 46, 75, 104, 52, 50, 51, 108, 109, 110, 111, 112, 113, 112, 112, 116, 117, 117, 116, 120, 116, 117, 120, 120, 125, 126, 127, 128, 129, 128, 131, 132, 133, 132, 131, 131, 128, 138, 139, 132, 141, 142, 143, 144, 145, 146, 147, 148, 146, 150, 141, 142, 153, 141, 146, 156, 153, 150, 159, 150, 148, 148}
var __const_scan_tag struct {
	F0 int32
	F1 [4]byte
	F2 String
} = struct {
	F0 int32
	F1 [4]byte
	F2 String
}{124, [4]byte{}, String{}}
var _str [64]byte = [64]byte{40, 117, 105, 110, 116, 51, 50, 95, 116, 41, 40, 40, 38, 115, 99, 97, 110, 110, 101, 114, 45, 62, 116, 97, 103, 115, 41, 45, 62, 115, 105, 122, 101, 32, 45, 32, 49, 41, 32, 60, 32, 40, 38, 115, 99, 97, 110, 110, 101, 114, 45, 62, 116, 97, 103, 115, 41, 45, 62, 115, 105, 122, 101, 0}
var _str_1 [40]byte = [40]byte{47, 116, 109, 112, 47, 108, 101, 97, 118, 101, 110, 45, 97, 115, 116, 114, 111, 45, 49, 51, 51, 51, 55, 54, 48, 48, 51, 52, 47, 99, 111, 109, 98, 105, 110, 101, 100, 46, 99, 0}
var __PRETTY_FUNCTION___scan [48]byte = [48]byte{95, 66, 111, 111, 108, 32, 115, 99, 97, 110, 40, 83, 99, 97, 110, 110, 101, 114, 32, 42, 44, 32, 84, 83, 76, 101, 120, 101, 114, 32, 42, 44, 32, 99, 111, 110, 115, 116, 32, 95, 66, 111, 111, 108, 32, 42, 41, 0}
var _str_2 [5]byte = [5]byte{10, 45, 45, 45, 0}
var __PRETTY_FUNCTION___scan_raw_text [42]byte = [42]byte{95, 66, 111, 111, 108, 32, 115, 99, 97, 110, 95, 114, 97, 119, 95, 116, 101, 120, 116, 40, 83, 99, 97, 110, 110, 101, 114, 32, 42, 44, 32, 84, 83, 76, 101, 120, 101, 114, 32, 42, 41, 0}
var _str_3 [9]byte = [9]byte{60, 47, 115, 99, 114, 105, 112, 116, 0}
var _str_4 [8]byte = [8]byte{60, 47, 115, 116, 121, 108, 101, 0}
var __PRETTY_FUNCTION___scan_implicit_end_tag [50]byte = [50]byte{95, 66, 111, 111, 108, 32, 115, 99, 97, 110, 95, 105, 109, 112, 108, 105, 99, 105, 116, 95, 101, 110, 100, 95, 116, 97, 103, 40, 83, 99, 97, 110, 110, 101, 114, 32, 42, 44, 32, 84, 83, 76, 101, 120, 101, 114, 32, 42, 41, 0}
var TAG_TYPES_BY_TAG_NAME [126]TagMapEntry = [126]TagMapEntry{TagMapEntry{[16]byte{97, 114, 101, 97, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0}, TagMapEntry{[16]byte{98, 97, 115, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1}, TagMapEntry{[16]byte{98, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 2}, TagMapEntry{[16]byte{99, 111, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 3}, TagMapEntry{[16]byte{99, 111, 109, 109, 97, 110, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 4}, TagMapEntry{[16]byte{101, 109, 98, 101, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 5}, TagMapEntry{[16]byte{102, 114, 97, 109, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 6}, TagMapEntry{[16]byte{104, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 7}, TagMapEntry{[16]byte{105, 109, 97, 103, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 8}, TagMapEntry{[16]byte{105, 109, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 9}, TagMapEntry{[16]byte{105, 110, 112, 117, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10}, TagMapEntry{[16]byte{105, 115, 105, 110, 100, 101, 120, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 11}, TagMapEntry{[16]byte{107, 101, 121, 103, 101, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 12}, TagMapEntry{[16]byte{108, 105, 110, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 13}, TagMapEntry{[16]byte{109, 101, 110, 117, 105, 116, 101, 109, 0, 0, 0, 0, 0, 0, 0, 0}, 14}, TagMapEntry{[16]byte{109, 101, 116, 97, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 15}, TagMapEntry{[16]byte{110, 101, 120, 116, 105, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 16}, TagMapEntry{[16]byte{112, 97, 114, 97, 109, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 17}, TagMapEntry{[16]byte{115, 111, 117, 114, 99, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 18}, TagMapEntry{[16]byte{116, 114, 97, 99, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 19}, TagMapEntry{[16]byte{119, 98, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 20}, TagMapEntry{[16]byte{97, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 22}, TagMapEntry{[16]byte{97, 98, 98, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 23}, TagMapEntry{[16]byte{97, 100, 100, 114, 101, 115, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 24}, TagMapEntry{[16]byte{97, 114, 116, 105, 99, 108, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 25}, TagMapEntry{[16]byte{97, 115, 105, 100, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 26}, TagMapEntry{[16]byte{97, 117, 100, 105, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 27}, TagMapEntry{[16]byte{98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 28}, TagMapEntry{[16]byte{98, 100, 105, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 29}, TagMapEntry{[16]byte{98, 100, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 30}, TagMapEntry{[16]byte{98, 108, 111, 99, 107, 113, 117, 111, 116, 101, 0, 0, 0, 0, 0, 0}, 31}, TagMapEntry{[16]byte{98, 111, 100, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 32}, TagMapEntry{[16]byte{98, 117, 116, 116, 111, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 33}, TagMapEntry{[16]byte{99, 97, 110, 118, 97, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 34}, TagMapEntry{[16]byte{99, 97, 112, 116, 105, 111, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 35}, TagMapEntry{[16]byte{99, 105, 116, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 36}, TagMapEntry{[16]byte{99, 111, 100, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 37}, TagMapEntry{[16]byte{99, 111, 108, 103, 114, 111, 117, 112, 0, 0, 0, 0, 0, 0, 0, 0}, 38}, TagMapEntry{[16]byte{100, 97, 116, 97, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 39}, TagMapEntry{[16]byte{100, 97, 116, 97, 108, 105, 115, 116, 0, 0, 0, 0, 0, 0, 0, 0}, 40}, TagMapEntry{[16]byte{100, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 41}, TagMapEntry{[16]byte{100, 101, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 42}, TagMapEntry{[16]byte{100, 101, 116, 97, 105, 108, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 43}, TagMapEntry{[16]byte{100, 102, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 44}, TagMapEntry{[16]byte{100, 105, 97, 108, 111, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 45}, TagMapEntry{[16]byte{100, 105, 118, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 46}, TagMapEntry{[16]byte{100, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 47}, TagMapEntry{[16]byte{100, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 48}, TagMapEntry{[16]byte{101, 109, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 49}, TagMapEntry{[16]byte{102, 105, 101, 108, 100, 115, 101, 116, 0, 0, 0, 0, 0, 0, 0, 0}, 50}, TagMapEntry{[16]byte{102, 105, 103, 99, 97, 112, 116, 105, 111, 110, 0, 0, 0, 0, 0, 0}, 51}, TagMapEntry{[16]byte{102, 105, 103, 117, 114, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 52}, TagMapEntry{[16]byte{102, 111, 111, 116, 101, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 53}, TagMapEntry{[16]byte{102, 111, 114, 109, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 54}, TagMapEntry{[16]byte{104, 49, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 55}, TagMapEntry{[16]byte{104, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 56}, TagMapEntry{[16]byte{104, 51, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 57}, TagMapEntry{[16]byte{104, 52, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 58}, TagMapEntry{[16]byte{104, 53, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 59}, TagMapEntry{[16]byte{104, 54, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 60}, TagMapEntry{[16]byte{104, 101, 97, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 61}, TagMapEntry{[16]byte{104, 101, 97, 100, 101, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 62}, TagMapEntry{[16]byte{104, 103, 114, 111, 117, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 63}, TagMapEntry{[16]byte{104, 116, 109, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 64}, TagMapEntry{[16]byte{105, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 65}, TagMapEntry{[16]byte{105, 102, 114, 97, 109, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 66}, TagMapEntry{[16]byte{105, 110, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 67}, TagMapEntry{[16]byte{107, 98, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 68}, TagMapEntry{[16]byte{108, 97, 98, 101, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 69}, TagMapEntry{[16]byte{108, 101, 103, 101, 110, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 70}, TagMapEntry{[16]byte{108, 105, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 71}, TagMapEntry{[16]byte{109, 97, 105, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 72}, TagMapEntry{[16]byte{109, 97, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 73}, TagMapEntry{[16]byte{109, 97, 114, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 74}, TagMapEntry{[16]byte{109, 97, 116, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 75}, TagMapEntry{[16]byte{109, 101, 110, 117, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 76}, TagMapEntry{[16]byte{109, 101, 116, 101, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 77}, TagMapEntry{[16]byte{110, 97, 118, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 78}, TagMapEntry{[16]byte{110, 111, 115, 99, 114, 105, 112, 116, 0, 0, 0, 0, 0, 0, 0, 0}, 79}, TagMapEntry{[16]byte{111, 98, 106, 101, 99, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 80}, TagMapEntry{[16]byte{111, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 81}, TagMapEntry{[16]byte{111, 112, 116, 103, 114, 111, 117, 112, 0, 0, 0, 0, 0, 0, 0, 0}, 82}, TagMapEntry{[16]byte{111, 112, 116, 105, 111, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 83}, TagMapEntry{[16]byte{111, 117, 116, 112, 117, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 84}, TagMapEntry{[16]byte{112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 85}, TagMapEntry{[16]byte{112, 105, 99, 116, 117, 114, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 86}, TagMapEntry{[16]byte{112, 114, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 87}, TagMapEntry{[16]byte{112, 114, 111, 103, 114, 101, 115, 115, 0, 0, 0, 0, 0, 0, 0, 0}, 88}, TagMapEntry{[16]byte{113, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 89}, TagMapEntry{[16]byte{114, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 90}, TagMapEntry{[16]byte{114, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 91}, TagMapEntry{[16]byte{114, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 92}, TagMapEntry{[16]byte{114, 116, 99, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 93}, TagMapEntry{[16]byte{114, 117, 98, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 94}, TagMapEntry{[16]byte{115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 95}, TagMapEntry{[16]byte{115, 97, 109, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 96}, TagMapEntry{[16]byte{115, 99, 114, 105, 112, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 97}, TagMapEntry{[16]byte{115, 101, 99, 116, 105, 111, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 98}, TagMapEntry{[16]byte{115, 101, 108, 101, 99, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 99}, TagMapEntry{[16]byte{115, 108, 111, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100}, TagMapEntry{[16]byte{115, 109, 97, 108, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 101}, TagMapEntry{[16]byte{115, 112, 97, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 102}, TagMapEntry{[16]byte{115, 116, 114, 111, 110, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 103}, TagMapEntry{[16]byte{115, 116, 121, 108, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 104}, TagMapEntry{[16]byte{115, 117, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 105}, TagMapEntry{[16]byte{115, 117, 109, 109, 97, 114, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 106}, TagMapEntry{[16]byte{115, 117, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 107}, TagMapEntry{[16]byte{115, 118, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 108}, TagMapEntry{[16]byte{116, 97, 98, 108, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 109}, TagMapEntry{[16]byte{116, 98, 111, 100, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 110}, TagMapEntry{[16]byte{116, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 111}, TagMapEntry{[16]byte{116, 101, 109, 112, 108, 97, 116, 101, 0, 0, 0, 0, 0, 0, 0, 0}, 112}, TagMapEntry{[16]byte{116, 101, 120, 116, 97, 114, 101, 97, 0, 0, 0, 0, 0, 0, 0, 0}, 113}, TagMapEntry{[16]byte{116, 102, 111, 111, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 114}, TagMapEntry{[16]byte{116, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 115}, TagMapEntry{[16]byte{116, 104, 101, 97, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 116}, TagMapEntry{[16]byte{116, 105, 109, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 117}, TagMapEntry{[16]byte{116, 105, 116, 108, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 118}, TagMapEntry{[16]byte{116, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 119}, TagMapEntry{[16]byte{117, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 120}, TagMapEntry{[16]byte{117, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 121}, TagMapEntry{[16]byte{118, 97, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 122}, TagMapEntry{[16]byte{118, 105, 100, 101, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 123}, TagMapEntry{[16]byte{99, 117, 115, 116, 111, 109, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 126}, TagMapEntry{}, TagMapEntry{}}
var TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS [26]int32 = [26]int32{24, 25, 26, 31, 43, 46, 47, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 62, 7, 72, 78, 81, 85, 87, 98}
var __const_scan_start_tag_name_tag struct {
	F0 int32
	F1 [4]byte
	F2 String
} = struct {
	F0 int32
	F1 [4]byte
	F2 String
}{125, [4]byte{}, String{}}
var __PRETTY_FUNCTION___scan_end_tag_name [46]byte = [46]byte{95, 66, 111, 111, 108, 32, 115, 99, 97, 110, 95, 101, 110, 100, 95, 116, 97, 103, 95, 110, 97, 109, 101, 40, 83, 99, 97, 110, 110, 101, 114, 32, 42, 44, 32, 84, 83, 76, 101, 120, 101, 114, 32, 42, 41, 0}
var ts_parse_table struct {
	F0 struct {
		F0 [36]int16
		F1 [23]int16
	}
	F1 [59]int16
} = struct {
	F0 struct {
		F0 [36]int16
		F1 [23]int16
	}
	F1 [59]int16
}{struct {
	F0 [36]int16
	F1 [23]int16
}{[36]int16{1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1}, [23]int16{}}, [59]int16{5, 7, 0, 0, 0, 9, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 13, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 17, 0, 0, 0, 0, 0, 0, 147, 12, 12, 12, 12, 12, 4, 119, 118, 64, 0, 12, 0, 0, 0, 11, 0, 12, 66, 74, 12, 0, 0}}
var ts_parse_actions struct {
	F0 struct {
		F0 anon_3
		F1 [6]byte
	}
	F1 struct {
		F0 anon_3
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_3
		F1 [6]byte
	}
	F4 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F5 struct {
		F0 anon_3
		F1 [6]byte
	}
	F6 TSParseActionEntry
	F7 struct {
		F0 anon_3
		F1 [6]byte
	}
	F8 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F9 struct {
		F0 anon_3
		F1 [6]byte
	}
	F10 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F11 struct {
		F0 anon_3
		F1 [6]byte
	}
	F12 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F13 struct {
		F0 anon_3
		F1 [6]byte
	}
	F14 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F15 struct {
		F0 anon_3
		F1 [6]byte
	}
	F16 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F17 struct {
		F0 anon_3
		F1 [6]byte
	}
	F18 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F19 struct {
		F0 anon_3
		F1 [6]byte
	}
	F20 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F21 struct {
		F0 anon_3
		F1 [6]byte
	}
	F22 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F23 struct {
		F0 anon_3
		F1 [6]byte
	}
	F24 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F25 struct {
		F0 anon_3
		F1 [6]byte
	}
	F26 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F27 struct {
		F0 anon_3
		F1 [6]byte
	}
	F28 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F29 struct {
		F0 anon_3
		F1 [6]byte
	}
	F30 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F31 struct {
		F0 anon_3
		F1 [6]byte
	}
	F32 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F33 struct {
		F0 anon_3
		F1 [6]byte
	}
	F34 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F35 struct {
		F0 anon_3
		F1 [6]byte
	}
	F36 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F37 struct {
		F0 anon_3
		F1 [6]byte
	}
	F38 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F39 struct {
		F0 anon_3
		F1 [6]byte
	}
	F40 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F41 struct {
		F0 anon_3
		F1 [6]byte
	}
	F42 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F43 struct {
		F0 anon_3
		F1 [6]byte
	}
	F44 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F45 struct {
		F0 anon_3
		F1 [6]byte
	}
	F46 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F47 struct {
		F0 anon_3
		F1 [6]byte
	}
	F48 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F49 struct {
		F0 anon_3
		F1 [6]byte
	}
	F50 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F51 struct {
		F0 anon_3
		F1 [6]byte
	}
	F52 TSParseActionEntry
	F53 struct {
		F0 anon_3
		F1 [6]byte
	}
	F54 TSParseActionEntry
	F55 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F56 struct {
		F0 anon_3
		F1 [6]byte
	}
	F57 TSParseActionEntry
	F58 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F59 struct {
		F0 anon_3
		F1 [6]byte
	}
	F60 TSParseActionEntry
	F61 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F62 struct {
		F0 anon_3
		F1 [6]byte
	}
	F63 TSParseActionEntry
	F64 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F65 struct {
		F0 anon_3
		F1 [6]byte
	}
	F66 TSParseActionEntry
	F67 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F68 struct {
		F0 anon_3
		F1 [6]byte
	}
	F69 TSParseActionEntry
	F70 struct {
		F0 anon_3
		F1 [6]byte
	}
	F71 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F72 struct {
		F0 anon_3
		F1 [6]byte
	}
	F73 TSParseActionEntry
	F74 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F75 struct {
		F0 anon_3
		F1 [6]byte
	}
	F76 TSParseActionEntry
	F77 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F78 struct {
		F0 anon_3
		F1 [6]byte
	}
	F79 TSParseActionEntry
	F80 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F81 struct {
		F0 anon_3
		F1 [6]byte
	}
	F82 TSParseActionEntry
	F83 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F84 struct {
		F0 anon_3
		F1 [6]byte
	}
	F85 TSParseActionEntry
	F86 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F87 struct {
		F0 anon_3
		F1 [6]byte
	}
	F88 TSParseActionEntry
	F89 struct {
		F0 anon_3
		F1 [6]byte
	}
	F90 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F91 struct {
		F0 anon_3
		F1 [6]byte
	}
	F92 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F93 struct {
		F0 anon_3
		F1 [6]byte
	}
	F94 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F95 struct {
		F0 anon_3
		F1 [6]byte
	}
	F96 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F97 struct {
		F0 anon_3
		F1 [6]byte
	}
	F98 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F99 struct {
		F0 anon_3
		F1 [6]byte
	}
	F100 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F101 struct {
		F0 anon_3
		F1 [6]byte
	}
	F102 TSParseActionEntry
	F103 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F104 struct {
		F0 anon_3
		F1 [6]byte
	}
	F105 TSParseActionEntry
	F106 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F107 struct {
		F0 anon_3
		F1 [6]byte
	}
	F108 TSParseActionEntry
	F109 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F110 struct {
		F0 anon_3
		F1 [6]byte
	}
	F111 TSParseActionEntry
	F112 struct {
		F0 anon_3
		F1 [6]byte
	}
	F113 TSParseActionEntry
	F114 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F115 struct {
		F0 anon_3
		F1 [6]byte
	}
	F116 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F117 struct {
		F0 anon_3
		F1 [6]byte
	}
	F118 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F119 struct {
		F0 anon_3
		F1 [6]byte
	}
	F120 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F121 struct {
		F0 anon_3
		F1 [6]byte
	}
	F122 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F123 struct {
		F0 anon_3
		F1 [6]byte
	}
	F124 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F125 struct {
		F0 anon_3
		F1 [6]byte
	}
	F126 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F127 struct {
		F0 anon_3
		F1 [6]byte
	}
	F128 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F129 struct {
		F0 anon_3
		F1 [6]byte
	}
	F130 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F131 struct {
		F0 anon_3
		F1 [6]byte
	}
	F132 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F133 struct {
		F0 anon_3
		F1 [6]byte
	}
	F134 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F135 struct {
		F0 anon_3
		F1 [6]byte
	}
	F136 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F137 struct {
		F0 anon_3
		F1 [6]byte
	}
	F138 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F139 struct {
		F0 anon_3
		F1 [6]byte
	}
	F140 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F141 struct {
		F0 anon_3
		F1 [6]byte
	}
	F142 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F143 struct {
		F0 anon_3
		F1 [6]byte
	}
	F144 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F145 struct {
		F0 anon_3
		F1 [6]byte
	}
	F146 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F147 struct {
		F0 anon_3
		F1 [6]byte
	}
	F148 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F149 struct {
		F0 anon_3
		F1 [6]byte
	}
	F150 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F151 struct {
		F0 anon_3
		F1 [6]byte
	}
	F152 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F153 struct {
		F0 anon_3
		F1 [6]byte
	}
	F154 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F155 struct {
		F0 anon_3
		F1 [6]byte
	}
	F156 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F157 struct {
		F0 anon_3
		F1 [6]byte
	}
	F158 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F159 struct {
		F0 anon_3
		F1 [6]byte
	}
	F160 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F161 struct {
		F0 anon_3
		F1 [6]byte
	}
	F162 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F163 struct {
		F0 anon_3
		F1 [6]byte
	}
	F164 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F165 struct {
		F0 anon_3
		F1 [6]byte
	}
	F166 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F167 struct {
		F0 anon_3
		F1 [6]byte
	}
	F168 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F169 struct {
		F0 anon_3
		F1 [6]byte
	}
	F170 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F171 struct {
		F0 anon_3
		F1 [6]byte
	}
	F172 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F173 struct {
		F0 anon_3
		F1 [6]byte
	}
	F174 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F175 struct {
		F0 anon_3
		F1 [6]byte
	}
	F176 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F177 struct {
		F0 anon_3
		F1 [6]byte
	}
	F178 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F179 struct {
		F0 anon_3
		F1 [6]byte
	}
	F180 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F181 struct {
		F0 anon_3
		F1 [6]byte
	}
	F182 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F183 struct {
		F0 anon_3
		F1 [6]byte
	}
	F184 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F185 struct {
		F0 anon_3
		F1 [6]byte
	}
	F186 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F187 struct {
		F0 anon_3
		F1 [6]byte
	}
	F188 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F189 struct {
		F0 anon_3
		F1 [6]byte
	}
	F190 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F191 struct {
		F0 anon_3
		F1 [6]byte
	}
	F192 TSParseActionEntry
	F193 struct {
		F0 anon_3
		F1 [6]byte
	}
	F194 TSParseActionEntry
	F195 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F196 struct {
		F0 anon_3
		F1 [6]byte
	}
	F197 TSParseActionEntry
	F198 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F199 struct {
		F0 anon_3
		F1 [6]byte
	}
	F200 TSParseActionEntry
	F201 struct {
		F0 anon_3
		F1 [6]byte
	}
	F202 TSParseActionEntry
	F203 struct {
		F0 anon_3
		F1 [6]byte
	}
	F204 TSParseActionEntry
	F205 struct {
		F0 anon_3
		F1 [6]byte
	}
	F206 TSParseActionEntry
	F207 struct {
		F0 anon_3
		F1 [6]byte
	}
	F208 TSParseActionEntry
	F209 struct {
		F0 anon_3
		F1 [6]byte
	}
	F210 TSParseActionEntry
	F211 struct {
		F0 anon_3
		F1 [6]byte
	}
	F212 TSParseActionEntry
	F213 struct {
		F0 anon_3
		F1 [6]byte
	}
	F214 TSParseActionEntry
	F215 struct {
		F0 anon_3
		F1 [6]byte
	}
	F216 TSParseActionEntry
	F217 struct {
		F0 anon_3
		F1 [6]byte
	}
	F218 TSParseActionEntry
	F219 struct {
		F0 anon_3
		F1 [6]byte
	}
	F220 TSParseActionEntry
	F221 struct {
		F0 anon_3
		F1 [6]byte
	}
	F222 TSParseActionEntry
	F223 struct {
		F0 anon_3
		F1 [6]byte
	}
	F224 TSParseActionEntry
	F225 struct {
		F0 anon_3
		F1 [6]byte
	}
	F226 TSParseActionEntry
	F227 struct {
		F0 anon_3
		F1 [6]byte
	}
	F228 TSParseActionEntry
	F229 struct {
		F0 anon_3
		F1 [6]byte
	}
	F230 TSParseActionEntry
	F231 struct {
		F0 anon_3
		F1 [6]byte
	}
	F232 TSParseActionEntry
	F233 struct {
		F0 anon_3
		F1 [6]byte
	}
	F234 TSParseActionEntry
	F235 struct {
		F0 anon_3
		F1 [6]byte
	}
	F236 TSParseActionEntry
	F237 struct {
		F0 anon_3
		F1 [6]byte
	}
	F238 TSParseActionEntry
	F239 struct {
		F0 anon_3
		F1 [6]byte
	}
	F240 TSParseActionEntry
	F241 struct {
		F0 anon_3
		F1 [6]byte
	}
	F242 TSParseActionEntry
	F243 struct {
		F0 anon_3
		F1 [6]byte
	}
	F244 TSParseActionEntry
	F245 struct {
		F0 anon_3
		F1 [6]byte
	}
	F246 TSParseActionEntry
	F247 struct {
		F0 anon_3
		F1 [6]byte
	}
	F248 TSParseActionEntry
	F249 struct {
		F0 anon_3
		F1 [6]byte
	}
	F250 TSParseActionEntry
	F251 struct {
		F0 anon_3
		F1 [6]byte
	}
	F252 TSParseActionEntry
	F253 struct {
		F0 anon_3
		F1 [6]byte
	}
	F254 TSParseActionEntry
	F255 struct {
		F0 anon_3
		F1 [6]byte
	}
	F256 TSParseActionEntry
	F257 struct {
		F0 anon_3
		F1 [6]byte
	}
	F258 TSParseActionEntry
	F259 struct {
		F0 anon_3
		F1 [6]byte
	}
	F260 TSParseActionEntry
	F261 struct {
		F0 anon_3
		F1 [6]byte
	}
	F262 TSParseActionEntry
	F263 struct {
		F0 anon_3
		F1 [6]byte
	}
	F264 TSParseActionEntry
	F265 struct {
		F0 anon_3
		F1 [6]byte
	}
	F266 TSParseActionEntry
	F267 struct {
		F0 anon_3
		F1 [6]byte
	}
	F268 TSParseActionEntry
	F269 struct {
		F0 anon_3
		F1 [6]byte
	}
	F270 TSParseActionEntry
	F271 struct {
		F0 anon_3
		F1 [6]byte
	}
	F272 TSParseActionEntry
	F273 struct {
		F0 anon_3
		F1 [6]byte
	}
	F274 TSParseActionEntry
	F275 struct {
		F0 anon_3
		F1 [6]byte
	}
	F276 TSParseActionEntry
	F277 struct {
		F0 anon_3
		F1 [6]byte
	}
	F278 TSParseActionEntry
	F279 struct {
		F0 anon_3
		F1 [6]byte
	}
	F280 TSParseActionEntry
	F281 struct {
		F0 anon_3
		F1 [6]byte
	}
	F282 TSParseActionEntry
	F283 struct {
		F0 anon_3
		F1 [6]byte
	}
	F284 TSParseActionEntry
	F285 struct {
		F0 anon_3
		F1 [6]byte
	}
	F286 TSParseActionEntry
	F287 struct {
		F0 anon_3
		F1 [6]byte
	}
	F288 TSParseActionEntry
	F289 struct {
		F0 anon_3
		F1 [6]byte
	}
	F290 TSParseActionEntry
	F291 struct {
		F0 anon_3
		F1 [6]byte
	}
	F292 TSParseActionEntry
	F293 struct {
		F0 anon_3
		F1 [6]byte
	}
	F294 TSParseActionEntry
	F295 struct {
		F0 anon_3
		F1 [6]byte
	}
	F296 TSParseActionEntry
	F297 struct {
		F0 anon_3
		F1 [6]byte
	}
	F298 TSParseActionEntry
	F299 struct {
		F0 anon_3
		F1 [6]byte
	}
	F300 TSParseActionEntry
	F301 struct {
		F0 anon_3
		F1 [6]byte
	}
	F302 TSParseActionEntry
	F303 struct {
		F0 anon_3
		F1 [6]byte
	}
	F304 TSParseActionEntry
	F305 struct {
		F0 anon_3
		F1 [6]byte
	}
	F306 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F307 struct {
		F0 anon_3
		F1 [6]byte
	}
	F308 TSParseActionEntry
	F309 struct {
		F0 anon_3
		F1 [6]byte
	}
	F310 TSParseActionEntry
	F311 struct {
		F0 anon_3
		F1 [6]byte
	}
	F312 TSParseActionEntry
	F313 struct {
		F0 anon_3
		F1 [6]byte
	}
	F314 TSParseActionEntry
	F315 struct {
		F0 anon_3
		F1 [6]byte
	}
	F316 TSParseActionEntry
	F317 struct {
		F0 anon_3
		F1 [6]byte
	}
	F318 TSParseActionEntry
	F319 struct {
		F0 anon_3
		F1 [6]byte
	}
	F320 TSParseActionEntry
	F321 struct {
		F0 anon_3
		F1 [6]byte
	}
	F322 TSParseActionEntry
	F323 struct {
		F0 anon_3
		F1 [6]byte
	}
	F324 TSParseActionEntry
	F325 struct {
		F0 anon_3
		F1 [6]byte
	}
	F326 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F327 struct {
		F0 anon_3
		F1 [6]byte
	}
	F328 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F329 struct {
		F0 anon_3
		F1 [6]byte
	}
	F330 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F331 struct {
		F0 anon_3
		F1 [6]byte
	}
	F332 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F333 struct {
		F0 anon_3
		F1 [6]byte
	}
	F334 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F335 struct {
		F0 anon_3
		F1 [6]byte
	}
	F336 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F337 struct {
		F0 anon_3
		F1 [6]byte
	}
	F338 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F339 struct {
		F0 anon_3
		F1 [6]byte
	}
	F340 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F341 struct {
		F0 anon_3
		F1 [6]byte
	}
	F342 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F343 struct {
		F0 anon_3
		F1 [6]byte
	}
	F344 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F345 struct {
		F0 anon_3
		F1 [6]byte
	}
	F346 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F347 struct {
		F0 anon_3
		F1 [6]byte
	}
	F348 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F349 struct {
		F0 anon_3
		F1 [6]byte
	}
	F350 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F351 struct {
		F0 anon_3
		F1 [6]byte
	}
	F352 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F353 struct {
		F0 anon_3
		F1 [6]byte
	}
	F354 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F355 struct {
		F0 anon_3
		F1 [6]byte
	}
	F356 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F357 struct {
		F0 anon_3
		F1 [6]byte
	}
	F358 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F359 struct {
		F0 anon_3
		F1 [6]byte
	}
	F360 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F361 struct {
		F0 anon_3
		F1 [6]byte
	}
	F362 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F363 struct {
		F0 anon_3
		F1 [6]byte
	}
	F364 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F365 struct {
		F0 anon_3
		F1 [6]byte
	}
	F366 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F367 struct {
		F0 anon_3
		F1 [6]byte
	}
	F368 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F369 struct {
		F0 anon_3
		F1 [6]byte
	}
	F370 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F371 struct {
		F0 anon_3
		F1 [6]byte
	}
	F372 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F373 struct {
		F0 anon_3
		F1 [6]byte
	}
	F374 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F375 struct {
		F0 anon_3
		F1 [6]byte
	}
	F376 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F377 struct {
		F0 anon_3
		F1 [6]byte
	}
	F378 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F379 struct {
		F0 anon_3
		F1 [6]byte
	}
	F380 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F381 struct {
		F0 anon_3
		F1 [6]byte
	}
	F382 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F383 struct {
		F0 anon_3
		F1 [6]byte
	}
	F384 TSParseActionEntry
	F385 struct {
		F0 anon_3
		F1 [6]byte
	}
	F386 TSParseActionEntry
	F387 struct {
		F0 anon_3
		F1 [6]byte
	}
	F388 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F389 struct {
		F0 anon_3
		F1 [6]byte
	}
	F390 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F391 struct {
		F0 anon_3
		F1 [6]byte
	}
	F392 TSParseActionEntry
	F393 struct {
		F0 anon_3
		F1 [6]byte
	}
	F394 TSParseActionEntry
	F395 struct {
		F0 anon_3
		F1 [6]byte
	}
	F396 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F397 struct {
		F0 anon_3
		F1 [6]byte
	}
	F398 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F399 struct {
		F0 anon_3
		F1 [6]byte
	}
	F400 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F401 struct {
		F0 anon_3
		F1 [6]byte
	}
	F402 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F403 struct {
		F0 anon_3
		F1 [6]byte
	}
	F404 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F405 struct {
		F0 anon_3
		F1 [6]byte
	}
	F406 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F407 struct {
		F0 anon_3
		F1 [6]byte
	}
	F408 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F409 struct {
		F0 anon_3
		F1 [6]byte
	}
	F410 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F411 struct {
		F0 anon_3
		F1 [6]byte
	}
	F412 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F413 struct {
		F0 anon_3
		F1 [6]byte
	}
	F414 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F415 struct {
		F0 anon_3
		F1 [6]byte
	}
	F416 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F417 struct {
		F0 anon_3
		F1 [6]byte
	}
	F418 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F419 struct {
		F0 anon_3
		F1 [6]byte
	}
	F420 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F421 struct {
		F0 anon_3
		F1 [6]byte
	}
	F422 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F423 struct {
		F0 anon_3
		F1 [6]byte
	}
	F424 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F425 struct {
		F0 anon_3
		F1 [6]byte
	}
	F426 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F427 struct {
		F0 anon_3
		F1 [6]byte
	}
	F428 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F429 struct {
		F0 anon_3
		F1 [6]byte
	}
	F430 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F431 struct {
		F0 anon_3
		F1 [6]byte
	}
	F432 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F433 struct {
		F0 anon_3
		F1 [6]byte
	}
	F434 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
} = struct {
	F0 struct {
		F0 anon_3
		F1 [6]byte
	}
	F1 struct {
		F0 anon_3
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_3
		F1 [6]byte
	}
	F4 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F5 struct {
		F0 anon_3
		F1 [6]byte
	}
	F6 TSParseActionEntry
	F7 struct {
		F0 anon_3
		F1 [6]byte
	}
	F8 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F9 struct {
		F0 anon_3
		F1 [6]byte
	}
	F10 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F11 struct {
		F0 anon_3
		F1 [6]byte
	}
	F12 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F13 struct {
		F0 anon_3
		F1 [6]byte
	}
	F14 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F15 struct {
		F0 anon_3
		F1 [6]byte
	}
	F16 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F17 struct {
		F0 anon_3
		F1 [6]byte
	}
	F18 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F19 struct {
		F0 anon_3
		F1 [6]byte
	}
	F20 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F21 struct {
		F0 anon_3
		F1 [6]byte
	}
	F22 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F23 struct {
		F0 anon_3
		F1 [6]byte
	}
	F24 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F25 struct {
		F0 anon_3
		F1 [6]byte
	}
	F26 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F27 struct {
		F0 anon_3
		F1 [6]byte
	}
	F28 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F29 struct {
		F0 anon_3
		F1 [6]byte
	}
	F30 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F31 struct {
		F0 anon_3
		F1 [6]byte
	}
	F32 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F33 struct {
		F0 anon_3
		F1 [6]byte
	}
	F34 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F35 struct {
		F0 anon_3
		F1 [6]byte
	}
	F36 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F37 struct {
		F0 anon_3
		F1 [6]byte
	}
	F38 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F39 struct {
		F0 anon_3
		F1 [6]byte
	}
	F40 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F41 struct {
		F0 anon_3
		F1 [6]byte
	}
	F42 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F43 struct {
		F0 anon_3
		F1 [6]byte
	}
	F44 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F45 struct {
		F0 anon_3
		F1 [6]byte
	}
	F46 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F47 struct {
		F0 anon_3
		F1 [6]byte
	}
	F48 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F49 struct {
		F0 anon_3
		F1 [6]byte
	}
	F50 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F51 struct {
		F0 anon_3
		F1 [6]byte
	}
	F52 TSParseActionEntry
	F53 struct {
		F0 anon_3
		F1 [6]byte
	}
	F54 TSParseActionEntry
	F55 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F56 struct {
		F0 anon_3
		F1 [6]byte
	}
	F57 TSParseActionEntry
	F58 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F59 struct {
		F0 anon_3
		F1 [6]byte
	}
	F60 TSParseActionEntry
	F61 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F62 struct {
		F0 anon_3
		F1 [6]byte
	}
	F63 TSParseActionEntry
	F64 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F65 struct {
		F0 anon_3
		F1 [6]byte
	}
	F66 TSParseActionEntry
	F67 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F68 struct {
		F0 anon_3
		F1 [6]byte
	}
	F69 TSParseActionEntry
	F70 struct {
		F0 anon_3
		F1 [6]byte
	}
	F71 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F72 struct {
		F0 anon_3
		F1 [6]byte
	}
	F73 TSParseActionEntry
	F74 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F75 struct {
		F0 anon_3
		F1 [6]byte
	}
	F76 TSParseActionEntry
	F77 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F78 struct {
		F0 anon_3
		F1 [6]byte
	}
	F79 TSParseActionEntry
	F80 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F81 struct {
		F0 anon_3
		F1 [6]byte
	}
	F82 TSParseActionEntry
	F83 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F84 struct {
		F0 anon_3
		F1 [6]byte
	}
	F85 TSParseActionEntry
	F86 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F87 struct {
		F0 anon_3
		F1 [6]byte
	}
	F88 TSParseActionEntry
	F89 struct {
		F0 anon_3
		F1 [6]byte
	}
	F90 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F91 struct {
		F0 anon_3
		F1 [6]byte
	}
	F92 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F93 struct {
		F0 anon_3
		F1 [6]byte
	}
	F94 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F95 struct {
		F0 anon_3
		F1 [6]byte
	}
	F96 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F97 struct {
		F0 anon_3
		F1 [6]byte
	}
	F98 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F99 struct {
		F0 anon_3
		F1 [6]byte
	}
	F100 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F101 struct {
		F0 anon_3
		F1 [6]byte
	}
	F102 TSParseActionEntry
	F103 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F104 struct {
		F0 anon_3
		F1 [6]byte
	}
	F105 TSParseActionEntry
	F106 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F107 struct {
		F0 anon_3
		F1 [6]byte
	}
	F108 TSParseActionEntry
	F109 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F110 struct {
		F0 anon_3
		F1 [6]byte
	}
	F111 TSParseActionEntry
	F112 struct {
		F0 anon_3
		F1 [6]byte
	}
	F113 TSParseActionEntry
	F114 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F115 struct {
		F0 anon_3
		F1 [6]byte
	}
	F116 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F117 struct {
		F0 anon_3
		F1 [6]byte
	}
	F118 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F119 struct {
		F0 anon_3
		F1 [6]byte
	}
	F120 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F121 struct {
		F0 anon_3
		F1 [6]byte
	}
	F122 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F123 struct {
		F0 anon_3
		F1 [6]byte
	}
	F124 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F125 struct {
		F0 anon_3
		F1 [6]byte
	}
	F126 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F127 struct {
		F0 anon_3
		F1 [6]byte
	}
	F128 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F129 struct {
		F0 anon_3
		F1 [6]byte
	}
	F130 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F131 struct {
		F0 anon_3
		F1 [6]byte
	}
	F132 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F133 struct {
		F0 anon_3
		F1 [6]byte
	}
	F134 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F135 struct {
		F0 anon_3
		F1 [6]byte
	}
	F136 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F137 struct {
		F0 anon_3
		F1 [6]byte
	}
	F138 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F139 struct {
		F0 anon_3
		F1 [6]byte
	}
	F140 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F141 struct {
		F0 anon_3
		F1 [6]byte
	}
	F142 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F143 struct {
		F0 anon_3
		F1 [6]byte
	}
	F144 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F145 struct {
		F0 anon_3
		F1 [6]byte
	}
	F146 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F147 struct {
		F0 anon_3
		F1 [6]byte
	}
	F148 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F149 struct {
		F0 anon_3
		F1 [6]byte
	}
	F150 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F151 struct {
		F0 anon_3
		F1 [6]byte
	}
	F152 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F153 struct {
		F0 anon_3
		F1 [6]byte
	}
	F154 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F155 struct {
		F0 anon_3
		F1 [6]byte
	}
	F156 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F157 struct {
		F0 anon_3
		F1 [6]byte
	}
	F158 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F159 struct {
		F0 anon_3
		F1 [6]byte
	}
	F160 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F161 struct {
		F0 anon_3
		F1 [6]byte
	}
	F162 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F163 struct {
		F0 anon_3
		F1 [6]byte
	}
	F164 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F165 struct {
		F0 anon_3
		F1 [6]byte
	}
	F166 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F167 struct {
		F0 anon_3
		F1 [6]byte
	}
	F168 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F169 struct {
		F0 anon_3
		F1 [6]byte
	}
	F170 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F171 struct {
		F0 anon_3
		F1 [6]byte
	}
	F172 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F173 struct {
		F0 anon_3
		F1 [6]byte
	}
	F174 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F175 struct {
		F0 anon_3
		F1 [6]byte
	}
	F176 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F177 struct {
		F0 anon_3
		F1 [6]byte
	}
	F178 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F179 struct {
		F0 anon_3
		F1 [6]byte
	}
	F180 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F181 struct {
		F0 anon_3
		F1 [6]byte
	}
	F182 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F183 struct {
		F0 anon_3
		F1 [6]byte
	}
	F184 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F185 struct {
		F0 anon_3
		F1 [6]byte
	}
	F186 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F187 struct {
		F0 anon_3
		F1 [6]byte
	}
	F188 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F189 struct {
		F0 anon_3
		F1 [6]byte
	}
	F190 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F191 struct {
		F0 anon_3
		F1 [6]byte
	}
	F192 TSParseActionEntry
	F193 struct {
		F0 anon_3
		F1 [6]byte
	}
	F194 TSParseActionEntry
	F195 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F196 struct {
		F0 anon_3
		F1 [6]byte
	}
	F197 TSParseActionEntry
	F198 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F199 struct {
		F0 anon_3
		F1 [6]byte
	}
	F200 TSParseActionEntry
	F201 struct {
		F0 anon_3
		F1 [6]byte
	}
	F202 TSParseActionEntry
	F203 struct {
		F0 anon_3
		F1 [6]byte
	}
	F204 TSParseActionEntry
	F205 struct {
		F0 anon_3
		F1 [6]byte
	}
	F206 TSParseActionEntry
	F207 struct {
		F0 anon_3
		F1 [6]byte
	}
	F208 TSParseActionEntry
	F209 struct {
		F0 anon_3
		F1 [6]byte
	}
	F210 TSParseActionEntry
	F211 struct {
		F0 anon_3
		F1 [6]byte
	}
	F212 TSParseActionEntry
	F213 struct {
		F0 anon_3
		F1 [6]byte
	}
	F214 TSParseActionEntry
	F215 struct {
		F0 anon_3
		F1 [6]byte
	}
	F216 TSParseActionEntry
	F217 struct {
		F0 anon_3
		F1 [6]byte
	}
	F218 TSParseActionEntry
	F219 struct {
		F0 anon_3
		F1 [6]byte
	}
	F220 TSParseActionEntry
	F221 struct {
		F0 anon_3
		F1 [6]byte
	}
	F222 TSParseActionEntry
	F223 struct {
		F0 anon_3
		F1 [6]byte
	}
	F224 TSParseActionEntry
	F225 struct {
		F0 anon_3
		F1 [6]byte
	}
	F226 TSParseActionEntry
	F227 struct {
		F0 anon_3
		F1 [6]byte
	}
	F228 TSParseActionEntry
	F229 struct {
		F0 anon_3
		F1 [6]byte
	}
	F230 TSParseActionEntry
	F231 struct {
		F0 anon_3
		F1 [6]byte
	}
	F232 TSParseActionEntry
	F233 struct {
		F0 anon_3
		F1 [6]byte
	}
	F234 TSParseActionEntry
	F235 struct {
		F0 anon_3
		F1 [6]byte
	}
	F236 TSParseActionEntry
	F237 struct {
		F0 anon_3
		F1 [6]byte
	}
	F238 TSParseActionEntry
	F239 struct {
		F0 anon_3
		F1 [6]byte
	}
	F240 TSParseActionEntry
	F241 struct {
		F0 anon_3
		F1 [6]byte
	}
	F242 TSParseActionEntry
	F243 struct {
		F0 anon_3
		F1 [6]byte
	}
	F244 TSParseActionEntry
	F245 struct {
		F0 anon_3
		F1 [6]byte
	}
	F246 TSParseActionEntry
	F247 struct {
		F0 anon_3
		F1 [6]byte
	}
	F248 TSParseActionEntry
	F249 struct {
		F0 anon_3
		F1 [6]byte
	}
	F250 TSParseActionEntry
	F251 struct {
		F0 anon_3
		F1 [6]byte
	}
	F252 TSParseActionEntry
	F253 struct {
		F0 anon_3
		F1 [6]byte
	}
	F254 TSParseActionEntry
	F255 struct {
		F0 anon_3
		F1 [6]byte
	}
	F256 TSParseActionEntry
	F257 struct {
		F0 anon_3
		F1 [6]byte
	}
	F258 TSParseActionEntry
	F259 struct {
		F0 anon_3
		F1 [6]byte
	}
	F260 TSParseActionEntry
	F261 struct {
		F0 anon_3
		F1 [6]byte
	}
	F262 TSParseActionEntry
	F263 struct {
		F0 anon_3
		F1 [6]byte
	}
	F264 TSParseActionEntry
	F265 struct {
		F0 anon_3
		F1 [6]byte
	}
	F266 TSParseActionEntry
	F267 struct {
		F0 anon_3
		F1 [6]byte
	}
	F268 TSParseActionEntry
	F269 struct {
		F0 anon_3
		F1 [6]byte
	}
	F270 TSParseActionEntry
	F271 struct {
		F0 anon_3
		F1 [6]byte
	}
	F272 TSParseActionEntry
	F273 struct {
		F0 anon_3
		F1 [6]byte
	}
	F274 TSParseActionEntry
	F275 struct {
		F0 anon_3
		F1 [6]byte
	}
	F276 TSParseActionEntry
	F277 struct {
		F0 anon_3
		F1 [6]byte
	}
	F278 TSParseActionEntry
	F279 struct {
		F0 anon_3
		F1 [6]byte
	}
	F280 TSParseActionEntry
	F281 struct {
		F0 anon_3
		F1 [6]byte
	}
	F282 TSParseActionEntry
	F283 struct {
		F0 anon_3
		F1 [6]byte
	}
	F284 TSParseActionEntry
	F285 struct {
		F0 anon_3
		F1 [6]byte
	}
	F286 TSParseActionEntry
	F287 struct {
		F0 anon_3
		F1 [6]byte
	}
	F288 TSParseActionEntry
	F289 struct {
		F0 anon_3
		F1 [6]byte
	}
	F290 TSParseActionEntry
	F291 struct {
		F0 anon_3
		F1 [6]byte
	}
	F292 TSParseActionEntry
	F293 struct {
		F0 anon_3
		F1 [6]byte
	}
	F294 TSParseActionEntry
	F295 struct {
		F0 anon_3
		F1 [6]byte
	}
	F296 TSParseActionEntry
	F297 struct {
		F0 anon_3
		F1 [6]byte
	}
	F298 TSParseActionEntry
	F299 struct {
		F0 anon_3
		F1 [6]byte
	}
	F300 TSParseActionEntry
	F301 struct {
		F0 anon_3
		F1 [6]byte
	}
	F302 TSParseActionEntry
	F303 struct {
		F0 anon_3
		F1 [6]byte
	}
	F304 TSParseActionEntry
	F305 struct {
		F0 anon_3
		F1 [6]byte
	}
	F306 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F307 struct {
		F0 anon_3
		F1 [6]byte
	}
	F308 TSParseActionEntry
	F309 struct {
		F0 anon_3
		F1 [6]byte
	}
	F310 TSParseActionEntry
	F311 struct {
		F0 anon_3
		F1 [6]byte
	}
	F312 TSParseActionEntry
	F313 struct {
		F0 anon_3
		F1 [6]byte
	}
	F314 TSParseActionEntry
	F315 struct {
		F0 anon_3
		F1 [6]byte
	}
	F316 TSParseActionEntry
	F317 struct {
		F0 anon_3
		F1 [6]byte
	}
	F318 TSParseActionEntry
	F319 struct {
		F0 anon_3
		F1 [6]byte
	}
	F320 TSParseActionEntry
	F321 struct {
		F0 anon_3
		F1 [6]byte
	}
	F322 TSParseActionEntry
	F323 struct {
		F0 anon_3
		F1 [6]byte
	}
	F324 TSParseActionEntry
	F325 struct {
		F0 anon_3
		F1 [6]byte
	}
	F326 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F327 struct {
		F0 anon_3
		F1 [6]byte
	}
	F328 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F329 struct {
		F0 anon_3
		F1 [6]byte
	}
	F330 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F331 struct {
		F0 anon_3
		F1 [6]byte
	}
	F332 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F333 struct {
		F0 anon_3
		F1 [6]byte
	}
	F334 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F335 struct {
		F0 anon_3
		F1 [6]byte
	}
	F336 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F337 struct {
		F0 anon_3
		F1 [6]byte
	}
	F338 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F339 struct {
		F0 anon_3
		F1 [6]byte
	}
	F340 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F341 struct {
		F0 anon_3
		F1 [6]byte
	}
	F342 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F343 struct {
		F0 anon_3
		F1 [6]byte
	}
	F344 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F345 struct {
		F0 anon_3
		F1 [6]byte
	}
	F346 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F347 struct {
		F0 anon_3
		F1 [6]byte
	}
	F348 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F349 struct {
		F0 anon_3
		F1 [6]byte
	}
	F350 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F351 struct {
		F0 anon_3
		F1 [6]byte
	}
	F352 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F353 struct {
		F0 anon_3
		F1 [6]byte
	}
	F354 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F355 struct {
		F0 anon_3
		F1 [6]byte
	}
	F356 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F357 struct {
		F0 anon_3
		F1 [6]byte
	}
	F358 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F359 struct {
		F0 anon_3
		F1 [6]byte
	}
	F360 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F361 struct {
		F0 anon_3
		F1 [6]byte
	}
	F362 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F363 struct {
		F0 anon_3
		F1 [6]byte
	}
	F364 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F365 struct {
		F0 anon_3
		F1 [6]byte
	}
	F366 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F367 struct {
		F0 anon_3
		F1 [6]byte
	}
	F368 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F369 struct {
		F0 anon_3
		F1 [6]byte
	}
	F370 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F371 struct {
		F0 anon_3
		F1 [6]byte
	}
	F372 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F373 struct {
		F0 anon_3
		F1 [6]byte
	}
	F374 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F375 struct {
		F0 anon_3
		F1 [6]byte
	}
	F376 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F377 struct {
		F0 anon_3
		F1 [6]byte
	}
	F378 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F379 struct {
		F0 anon_3
		F1 [6]byte
	}
	F380 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F381 struct {
		F0 anon_3
		F1 [6]byte
	}
	F382 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F383 struct {
		F0 anon_3
		F1 [6]byte
	}
	F384 TSParseActionEntry
	F385 struct {
		F0 anon_3
		F1 [6]byte
	}
	F386 TSParseActionEntry
	F387 struct {
		F0 anon_3
		F1 [6]byte
	}
	F388 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F389 struct {
		F0 anon_3
		F1 [6]byte
	}
	F390 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F391 struct {
		F0 anon_3
		F1 [6]byte
	}
	F392 TSParseActionEntry
	F393 struct {
		F0 anon_3
		F1 [6]byte
	}
	F394 TSParseActionEntry
	F395 struct {
		F0 anon_3
		F1 [6]byte
	}
	F396 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F397 struct {
		F0 anon_3
		F1 [6]byte
	}
	F398 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F399 struct {
		F0 anon_3
		F1 [6]byte
	}
	F400 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F401 struct {
		F0 anon_3
		F1 [6]byte
	}
	F402 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F403 struct {
		F0 anon_3
		F1 [6]byte
	}
	F404 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F405 struct {
		F0 anon_3
		F1 [6]byte
	}
	F406 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F407 struct {
		F0 anon_3
		F1 [6]byte
	}
	F408 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F409 struct {
		F0 anon_3
		F1 [6]byte
	}
	F410 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F411 struct {
		F0 anon_3
		F1 [6]byte
	}
	F412 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F413 struct {
		F0 anon_3
		F1 [6]byte
	}
	F414 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F415 struct {
		F0 anon_3
		F1 [6]byte
	}
	F416 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F417 struct {
		F0 anon_3
		F1 [6]byte
	}
	F418 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F419 struct {
		F0 anon_3
		F1 [6]byte
	}
	F420 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F421 struct {
		F0 anon_3
		F1 [6]byte
	}
	F422 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F423 struct {
		F0 anon_3
		F1 [6]byte
	}
	F424 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F425 struct {
		F0 anon_3
		F1 [6]byte
	}
	F426 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F427 struct {
		F0 anon_3
		F1 [6]byte
	}
	F428 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F429 struct {
		F0 anon_3
		F1 [6]byte
	}
	F430 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F431 struct {
		F0 anon_3
		F1 [6]byte
	}
	F432 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F433 struct {
		F0 anon_3
		F1 [6]byte
	}
	F434 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
}{struct {
	F0 anon_3
	F1 [6]byte
}{}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 byte
		F1 [7]byte
	}
}{struct {
	F0 byte
	F1 [7]byte
}{3, [7]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 0, 1, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 0, 36, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 148, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 112, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 153, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 12, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 133, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 13, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 162, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 115, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 124, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 10, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 91, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 17, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 120, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 71, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 123, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 7, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 46, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 3, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 80, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 102, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 57, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 148, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 112, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 153, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 8, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 13, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 36, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 8, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 162, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 115, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 157, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 10, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 56, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 17, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 36, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 9, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 161, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 114, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 15, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 75, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 18, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 58, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 161, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 58, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 114, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 58, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 15, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 58, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 58, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 14, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 103, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 19, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 78, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 14, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 81, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 16, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 43, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 99, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 45, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 77, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 104, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 159, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 129, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 52, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 62, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 95, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 127, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 50, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 105, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 106, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 89, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 139, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 49, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 40, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 138, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 107, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 97, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 76, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 67, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 65, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 73, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 55, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 111, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 125, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 126, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 111, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 51, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 63, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 57, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 57, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 104, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 57, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 159, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 45, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 45, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 40, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 40, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 41, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 41, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 53, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 53, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 51, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 51, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 42, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 42, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 46, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 46, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 55, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 55, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 54, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 54, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 54, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 54, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 55, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 55, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 47, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 47, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 51, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 51, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 45, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 45, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 46, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 46, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 39, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 39, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 40, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 40, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 41, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 41, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 37, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 37, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 42, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 42, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 39, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 39, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 40, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 40, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 41, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 41, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 53, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 53, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 42, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 42, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 48, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 36, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 48, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 49, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 49, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 49, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 49, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 52, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 52, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 48, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 48, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 28, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 23, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 21, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 84, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 26, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 25, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 24, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 20, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 31, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 34, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 136, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 137, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 134, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 135, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 140, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 128, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 155, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 152, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 72, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 131, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 130, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 132, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 146, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 56, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 149, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 90, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 109, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 143, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 144, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 43, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 44, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 44, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 156, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 43, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 44, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 60, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 53, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 108, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 110, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 68, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 byte
		F1 [7]byte
	}
}{struct {
	F0 byte
	F1 [7]byte
}{2, [7]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 150, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 98, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 141, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 61, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 86, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 142, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 94, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 48, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 54, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 154, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 145, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 151, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 158, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 160, 0, 0}, [2]byte{}}}}
var _str_7 [4]byte = [4]byte{101, 110, 100, 0}
var _str_8 [3]byte = [3]byte{60, 33, 0}
var _str_9 [15]byte = [15]byte{100, 111, 99, 116, 121, 112, 101, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_10 [2]byte = [2]byte{62, 0}
var _str_11 [8]byte = [8]byte{100, 111, 99, 116, 121, 112, 101, 0}
var _str_12 [2]byte = [2]byte{60, 0}
var _str_13 [3]byte = [3]byte{47, 62, 0}
var _str_14 [3]byte = [3]byte{60, 47, 0}
var _str_15 [2]byte = [2]byte{61, 0}
var _str_16 [15]byte = [15]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 110, 97, 109, 101, 0}
var _str_17 [16]byte = [16]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 118, 97, 108, 117, 101, 0}
var _str_18 [7]byte = [7]byte{101, 110, 116, 105, 116, 121, 0}
var _str_19 [2]byte = [2]byte{39, 0}
var _str_20 [2]byte = [2]byte{34, 0}
var _str_21 [5]byte = [5]byte{116, 101, 120, 116, 0}
var _str_22 [4]byte = [4]byte{45, 45, 45, 0}
var _str_23 [2]byte = [2]byte{123, 0}
var _str_24 [2]byte = [2]byte{125, 0}
var _str_25 [9]byte = [9]byte{116, 97, 103, 95, 110, 97, 109, 101, 0}
var _str_26 [23]byte = [23]byte{101, 114, 114, 111, 110, 101, 111, 117, 115, 95, 101, 110, 100, 95, 116, 97, 103, 95, 110, 97, 109, 101, 0}
var _str_27 [18]byte = [18]byte{95, 105, 109, 112, 108, 105, 99, 105, 116, 95, 101, 110, 100, 95, 116, 97, 103, 0}
var _str_28 [9]byte = [9]byte{114, 97, 119, 95, 116, 101, 120, 116, 0}
var _str_29 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_30 [21]byte = [21]byte{102, 114, 111, 110, 116, 109, 97, 116, 116, 101, 114, 95, 106, 115, 95, 98, 108, 111, 99, 107, 0}
var _str_31 [18]byte = [18]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 106, 115, 95, 101, 120, 112, 114, 0}
var _str_32 [26]byte = [26]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 98, 97, 99, 107, 116, 105, 99, 107, 95, 115, 116, 114, 105, 110, 103, 0}
var _str_33 [17]byte = [17]byte{112, 101, 114, 109, 105, 115, 115, 105, 98, 108, 101, 95, 116, 101, 120, 116, 0}
var _str_34 [9]byte = [9]byte{100, 111, 99, 117, 109, 101, 110, 116, 0}
var _str_35 [6]byte = [6]byte{95, 110, 111, 100, 101, 0}
var _str_36 [8]byte = [8]byte{101, 108, 101, 109, 101, 110, 116, 0}
var _str_37 [15]byte = [15]byte{115, 99, 114, 105, 112, 116, 95, 101, 108, 101, 109, 101, 110, 116, 0}
var _str_38 [14]byte = [14]byte{115, 116, 121, 108, 101, 95, 101, 108, 101, 109, 101, 110, 116, 0}
var _str_39 [10]byte = [10]byte{115, 116, 97, 114, 116, 95, 116, 97, 103, 0}
var _str_40 [17]byte = [17]byte{115, 101, 108, 102, 95, 99, 108, 111, 115, 105, 110, 103, 95, 116, 97, 103, 0}
var _str_41 [8]byte = [8]byte{101, 110, 100, 95, 116, 97, 103, 0}
var _str_42 [18]byte = [18]byte{101, 114, 114, 111, 110, 101, 111, 117, 115, 95, 101, 110, 100, 95, 116, 97, 103, 0}
var _str_43 [10]byte = [10]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 0}
var _str_44 [23]byte = [23]byte{113, 117, 111, 116, 101, 100, 95, 97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 118, 97, 108, 117, 101, 0}
var _str_45 [28]byte = [28]byte{95, 110, 111, 100, 101, 95, 119, 105, 116, 104, 95, 112, 101, 114, 109, 105, 115, 115, 105, 98, 108, 101, 95, 116, 101, 120, 116, 0}
var _str_46 [12]byte = [12]byte{102, 114, 111, 110, 116, 109, 97, 116, 116, 101, 114, 0}
var _str_47 [24]byte = [24]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 105, 110, 116, 101, 114, 112, 111, 108, 97, 116, 105, 111, 110, 0}
var _str_48 [19]byte = [19]byte{104, 116, 109, 108, 95, 105, 110, 116, 101, 114, 112, 111, 108, 97, 116, 105, 111, 110, 0}
var _str_49 [17]byte = [17]byte{100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_50 [18]byte = [18]byte{115, 116, 97, 114, 116, 95, 116, 97, 103, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_51 [27]byte = [27]byte{104, 116, 109, 108, 95, 105, 110, 116, 101, 114, 112, 111, 108, 97, 116, 105, 111, 110, 95, 114, 101, 112, 101, 97, 116, 49, 0}

func init() {
	tree_sitter_astro_language = struct {
		F0  int32
		F1  int32
		F2  int32
		F3  int32
		F4  int32
		F5  int32
		F6  int32
		F7  int32
		F8  int32
		F9  int16
		F10 [2]byte
		F11 unsafe.Pointer
		F12 unsafe.Pointer
		F13 unsafe.Pointer
		F14 unsafe.Pointer
		F15 unsafe.Pointer
		F16 unsafe.Pointer
		F17 unsafe.Pointer
		F18 unsafe.Pointer
		F19 unsafe.Pointer
		F20 unsafe.Pointer
		F21 unsafe.Pointer
		F22 unsafe.Pointer
		F23 unsafe.Pointer
		F24 unsafe.Pointer
		F25 unsafe.Pointer
		F26 int16
		F27 [6]byte
		F28 anon_0
		F29 unsafe.Pointer
	}{14, 59, 0, 36, 16, 163, 2, 1, 0, 4, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, anon_0{libc.Ptr(&ts_external_scanner_states), libc.Ptr(&ts_external_scanner_symbol_map), libc.FuncCode(tree_sitter_astro_external_scanner_create), libc.FuncCode(tree_sitter_astro_external_scanner_destroy), libc.FuncCode(tree_sitter_astro_external_scanner_scan), libc.FuncCode(tree_sitter_astro_external_scanner_serialize), libc.FuncCode(tree_sitter_astro_external_scanner_deserialize)}, libc.Ptr(&ts_primary_state_ids)}
}
func tree_sitter_astro_external_scanner_create() unsafe.Pointer {
	var call, v0 unsafe.Pointer
	var scanner unsafe.Pointer
	_, _, _ = scanner, call, v0

	scanner = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	call = libc.Ptr(libc.Calloc[byte](int64(1), int64(16)))
	*libc.As[unsafe.Pointer](scanner) = call
	v0 = *libc.As[unsafe.Pointer](scanner)
	return v0
}
func calloc(a0 int64, a1 int64) unsafe.Pointer {
	panic("unsatisfied: calloc")
}
func tree_sitter_astro_external_scanner_scan(payload unsafe.Pointer, lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var call bool
	var v0, v1, v2, v3 unsafe.Pointer
	var payload_addr, lexer_addr, valid_symbols_addr, scanner unsafe.Pointer
	_, _, _, _, _, _, _, _, _ = payload_addr, lexer_addr, valid_symbols_addr, scanner, v0, v1, v2, v3, call

	payload_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	valid_symbols_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	scanner = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](payload_addr)
	*libc.As[unsafe.Pointer](scanner) = v0
	v1 = *libc.As[unsafe.Pointer](scanner)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	v3 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	call = scan(v1, v2, v3)
	return call
}
func scan(scanner unsafe.Pointer, lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var tag, arrayidx83, arrayidx113, arrayidx122 unsafe.Pointer
	var tags, tags79, tags80, tags81, tags95, tags99, tags101, tags107, tags109, tags117, tags119 unsafe.Pointer
	var loadedv, cmp, loadedv2, loadedv5, loadedv8, call, loadedv12, loadedv17, tobool, call21, tobool25, cmp29, call31, loadedv34, call36, loadedv39, cmp42, cmp45, cmp47, cmp50, cmp53, cmp56, cmp58, v56, loadedv59, loadedv65, call67, loadedv71, call73, loadedv77, loadedv88, call90, loadedv93, cmp97, cmp103, cmp114, loadedv127, loadedv133, loadedv136, loadedv139, loadedv142, call143, call144, tobool146, loadedv148, loadedv151, call154, v125 bool
	var retval unsafe.Pointer
	var result_symbol, result_symbol14, result_symbol84, result_symbol123, result_symbol130 unsafe.Pointer
	var v3, v21, call19, v24, call24, v27, v33, v43, v45, v47, v49, v51, v53, v55, v75, inc, v84, v86, sub, v88, v92, sub111, v93, v100, dec, conv, conv145, cond int32
	var size, lookahead, lookahead23, lookahead27, lookahead28, lookahead41, lookahead44, lookahead46, lookahead49, lookahead52, lookahead55, lookahead57, size82, size96, size100, size102, size110, _type, size120 unsafe.Pointer
	var idxprom, idxprom112, idxprom121 int64
	var v1, v7, v9, v11, v15, v19, v37, v41, storedv, v57, v59, v63, v67, v78, v82, v103, v110, v112, v114, v116, v121, v123 byte
	var definitely_not_permissible_text, invalid, arrayidx, arrayidx1, arrayidx4, arrayidx7, arrayidx11, arrayidx16, arrayidx33, arrayidx38, arrayidx64, arrayidx70, arrayidx76, arrayidx87, arrayidx92, arrayidx126, arrayidx132, arrayidx135, arrayidx138, arrayidx141, arrayidx150 unsafe.Pointer
	var v0, v2, v4, v5, v6, v8, v10, v12, v13, v14, v16, v17, v18, v20, v22, v23, v25, v26, v28, v29, v30, v31, v32, v34, v35, v36, v38, v39, v40, v42, v44, v46, v48, v50, v52, v54, v58, v60, v61, v62, v64, v65, v66, v68, v69, v70, v71, v72, v73, v74, v76, v77, v79, v80, v81, v83, v85, v87, v89, v90, v91, v94, v95, v96, v97, v98, v99, v101, v102, v104, v105, v106, v107, v108, v109, v111, v113, v115, v117, v118, v119, v120, v122, v124 unsafe.Pointer
	var scanner_addr, lexer_addr, valid_symbols_addr, mark_end, local_advance, contents, contents108, advance116, contents118, mark_end129 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, valid_symbols_addr, definitely_not_permissible_text, invalid, tag, v0, arrayidx, v1, loadedv, v2, tags, size, v3, cmp, v4, v5, result_symbol, v6, arrayidx1, v7, loadedv2, v8, arrayidx4, v9, loadedv5, v10, arrayidx7, v11, loadedv8, v12, v13, call, v14, arrayidx11, v15, loadedv12, v16, v17, result_symbol14, v18, arrayidx16, v19, loadedv17, v20, lookahead, v21, call19, tobool, v22, call21, v23, lookahead23, v24, call24, tobool25, v25, v26, lookahead27, v27, v28, mark_end, v29, v30, v31, v32, lookahead28, v33, cmp29, v34, v35, call31, v36, arrayidx33, v37, loadedv34, v38, v39, call36, v40, arrayidx38, v41, loadedv39, v42, lookahead41, v43, cmp42, v44, lookahead44, v45, cmp45, v46, lookahead46, v47, cmp47, v48, lookahead49, v49, cmp50, v50, lookahead52, v51, cmp53, v52, lookahead55, v53, cmp56, v54, lookahead57, v55, cmp58, v56, storedv, v57, loadedv59, v58, arrayidx64, v59, loadedv65, v60, v61, call67, v62, arrayidx70, v63, loadedv71, v64, v65, call73, v66, arrayidx76, v67, loadedv77, v68, local_advance, v69, v70, v71, tags79, v72, tags80, contents, v73, v74, tags81, size82, v75, inc, idxprom, arrayidx83, v76, result_symbol84, v77, arrayidx87, v78, loadedv88, v79, v80, call90, v81, arrayidx92, v82, loadedv93, v83, tags95, size96, v84, cmp97, v85, tags99, size100, v86, sub, v87, tags101, size102, v88, cmp103, v89, tags107, contents108, v90, v91, tags109, size110, v92, sub111, idxprom112, arrayidx113, _type, v93, cmp114, v94, advance116, v95, v96, v97, tags117, contents118, v98, v99, tags119, size120, v100, dec, idxprom121, arrayidx122, v101, result_symbol123, v102, arrayidx126, v103, loadedv127, v104, v105, mark_end129, v106, v107, v108, result_symbol130, v109, arrayidx132, v110, loadedv133, v111, arrayidx135, v112, loadedv136, v113, arrayidx138, v114, loadedv139, v115, arrayidx141, v116, loadedv142, v117, v118, call143, conv, v119, v120, call144, conv145, cond, tobool146, v121, loadedv148, v122, arrayidx150, v123, loadedv151, v124, call154, v125

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	valid_symbols_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	definitely_not_permissible_text = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	invalid = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v0), int(int64(11))*1))
	v1 = *libc.As[byte](arrayidx)
	loadedv = (v1 & 1) != 0
	if loadedv {
		goto land_lhs_true
	} else {
		goto if_end
	}

land_lhs_true:
	v2 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v2).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v3 = *libc.As[int32](size)
	cmp = v3 == 0
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	scan_js_expr_with_delimiter(v4, 0)
	v5 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v5).F1)
	*libc.As[int16](result_symbol) = 11
	*libc.As[bool](retval) = true
	goto _return

if_end:
	v6 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx1 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v6), int(int64(7))*1))
	v7 = *libc.As[byte](arrayidx1)
	loadedv2 = (v7 & 1) != 0
	if loadedv2 {
		goto land_lhs_true3
	} else {
		goto if_end10
	}

land_lhs_true3:
	v8 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx4 = v8
	v9 = *libc.As[byte](arrayidx4)
	loadedv5 = (v9 & 1) != 0
	if loadedv5 {
		goto if_end10
	} else {
		goto land_lhs_true6
	}

land_lhs_true6:
	v10 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx7 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v10), int(int64(3))*1))
	v11 = *libc.As[byte](arrayidx7)
	loadedv8 = (v11 & 1) != 0
	if loadedv8 {
		goto if_end10
	} else {
		goto if_then9
	}

if_then9:
	v12 = *libc.As[unsafe.Pointer](scanner_addr)
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	call = scan_raw_text(v12, v13)
	*libc.As[bool](retval) = call
	goto _return

if_end10:
	v14 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx11 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v14), int(int64(12))*1))
	v15 = *libc.As[byte](arrayidx11)
	loadedv12 = (v15 & 1) != 0
	if loadedv12 {
		goto if_then13
	} else {
		goto if_end15
	}

if_then13:
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	scan_js_expr_with_delimiter(v16, 1)
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol14 = libc.Ptr(&libc.As[TSLexer](v17).F1)
	*libc.As[int16](result_symbol14) = 12
	*libc.As[bool](retval) = true
	goto _return

if_end15:
	v18 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx16 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v18), int(int64(14))*1))
	v19 = *libc.As[byte](arrayidx16)
	loadedv17 = (v19 & 1) != 0
	if loadedv17 {
		goto if_then18
	} else {
		goto if_else
	}

if_then18:
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v20).F0)
	v21 = *libc.As[int32](lookahead)
	call19 = libc.Iswspace(v21)
	tobool = call19 != 0
	if tobool {
		goto if_then20
	} else {
		goto if_end22
	}

if_then20:
	v22 = *libc.As[unsafe.Pointer](lexer_addr)
	call21 = scan_permissible_text(v22)
	*libc.As[bool](retval) = call21
	goto _return

if_end22:
	goto if_end26

if_else:
	goto while_cond

while_cond:
	v23 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead23 = libc.Ptr(&libc.As[TSLexer](v23).F0)
	v24 = *libc.As[int32](lookahead23)
	call24 = libc.Iswspace(v24)
	tobool25 = call24 != 0
	if tobool25 {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	skip(v25)
	goto while_cond

while_end:
	goto if_end26

if_end26:
	*libc.As[byte](definitely_not_permissible_text) = 0
	v26 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead27 = libc.Ptr(&libc.As[TSLexer](v26).F0)
	v27 = *libc.As[int32](lookahead27)
	switch v27 {
	case 60:
		goto sw_bb
	case 0:
		goto sw_bb63
	case 47:
		goto sw_bb69
	case 123:
		goto sw_bb75
	case 125:
		goto sw_bb86
	case 96:
		goto sw_bb125
	default:
		goto sw_default
	}

sw_bb:
	v28 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v28).F3)
	v29 = *libc.As[unsafe.Pointer](mark_end)
	v30 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v29)(v30)
	v31 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v31)
	v32 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead28 = libc.Ptr(&libc.As[TSLexer](v32).F0)
	v33 = *libc.As[int32](lookahead28)
	cmp29 = v33 == 33
	if cmp29 {
		goto if_then30
	} else {
		goto if_end32
	}

if_then30:
	v34 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v34)
	v35 = *libc.As[unsafe.Pointer](lexer_addr)
	call31 = scan_comment(v35)
	*libc.As[bool](retval) = call31
	goto _return

if_end32:
	v36 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx33 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v36), int(int64(6))*1))
	v37 = *libc.As[byte](arrayidx33)
	loadedv34 = (v37 & 1) != 0
	if loadedv34 {
		goto if_then35
	} else {
		goto if_end37
	}

if_then35:
	v38 = *libc.As[unsafe.Pointer](scanner_addr)
	v39 = *libc.As[unsafe.Pointer](lexer_addr)
	call36 = scan_implicit_end_tag(v38, v39)
	*libc.As[bool](retval) = call36
	goto _return

if_end37:
	v40 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx38 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v40), int(int64(14))*1))
	v41 = *libc.As[byte](arrayidx38)
	loadedv39 = (v41 & 1) != 0
	if loadedv39 {
		goto if_then40
	} else {
		goto if_end62
	}

if_then40:
	v42 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead41 = libc.Ptr(&libc.As[TSLexer](v42).F0)
	v43 = *libc.As[int32](lookahead41)
	cmp42 = 97 <= v43
	if cmp42 {
		goto land_lhs_true43
	} else {
		goto lor_lhs_false
	}

land_lhs_true43:
	v44 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead44 = libc.Ptr(&libc.As[TSLexer](v44).F0)
	v45 = *libc.As[int32](lookahead44)
	cmp45 = v45 <= 122
	if cmp45 {
		v56 = true
		goto lor_end
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v46 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead46 = libc.Ptr(&libc.As[TSLexer](v46).F0)
	v47 = *libc.As[int32](lookahead46)
	cmp47 = 65 <= v47
	if cmp47 {
		goto land_lhs_true48
	} else {
		goto lor_lhs_false51
	}

land_lhs_true48:
	v48 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead49 = libc.Ptr(&libc.As[TSLexer](v48).F0)
	v49 = *libc.As[int32](lookahead49)
	cmp50 = v49 <= 90
	if cmp50 {
		v56 = true
		goto lor_end
	} else {
		goto lor_lhs_false51
	}

lor_lhs_false51:
	v50 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead52 = libc.Ptr(&libc.As[TSLexer](v50).F0)
	v51 = *libc.As[int32](lookahead52)
	cmp53 = v51 == 47
	if cmp53 {
		v56 = true
		goto lor_end
	} else {
		goto lor_lhs_false54
	}

lor_lhs_false54:
	v52 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead55 = libc.Ptr(&libc.As[TSLexer](v52).F0)
	v53 = *libc.As[int32](lookahead55)
	cmp56 = v53 == 63
	if cmp56 {
		v56 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v54 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead57 = libc.Ptr(&libc.As[TSLexer](v54).F0)
	v55 = *libc.As[int32](lookahead57)
	cmp58 = v55 == 62
	v56 = cmp58
	goto lor_end

lor_end:
	if v56 {
		storedv = 1
	} else {
		storedv = 0
	}
	*libc.As[byte](invalid) = storedv
	v57 = *libc.As[byte](invalid)
	loadedv59 = (v57 & 1) != 0
	if loadedv59 {
		goto if_then60
	} else {
		goto if_end61
	}

if_then60:
	*libc.As[byte](definitely_not_permissible_text) = 1
	goto if_end61

if_end61:
	goto if_end62

if_end62:
	goto sw_epilog

sw_bb63:
	*libc.As[byte](definitely_not_permissible_text) = 1
	v58 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx64 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v58), int(int64(6))*1))
	v59 = *libc.As[byte](arrayidx64)
	loadedv65 = (v59 & 1) != 0
	if loadedv65 {
		goto if_then66
	} else {
		goto if_end68
	}

if_then66:
	v60 = *libc.As[unsafe.Pointer](scanner_addr)
	v61 = *libc.As[unsafe.Pointer](lexer_addr)
	call67 = scan_implicit_end_tag(v60, v61)
	*libc.As[bool](retval) = call67
	goto _return

if_end68:
	goto sw_epilog

sw_bb69:
	v62 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx70 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v62), int(int64(5))*1))
	v63 = *libc.As[byte](arrayidx70)
	loadedv71 = (v63 & 1) != 0
	if loadedv71 {
		goto if_then72
	} else {
		goto if_end74
	}

if_then72:
	v64 = *libc.As[unsafe.Pointer](scanner_addr)
	v65 = *libc.As[unsafe.Pointer](lexer_addr)
	call73 = scan_self_closing_tag_delimiter(v64, v65)
	*libc.As[bool](retval) = call73
	goto _return

if_end74:
	goto sw_epilog

sw_bb75:
	v66 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx76 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v66), int(int64(9))*1))
	v67 = *libc.As[byte](arrayidx76)
	loadedv77 = (v67 & 1) != 0
	if loadedv77 {
		goto if_then78
	} else {
		goto if_end85
	}

if_then78:
	v68 = *libc.As[unsafe.Pointer](lexer_addr)
	local_advance = libc.Ptr(&libc.As[TSLexer](v68).F2)
	v69 = *libc.As[unsafe.Pointer](local_advance)
	v70 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v69)(v70, false)
	libc.Memmove(libc.As[byte](tag), libc.As[byte](libc.Ptr(&__const_scan_tag)), int64(24))
	v71 = *libc.As[unsafe.Pointer](scanner_addr)
	tags79 = libc.Ptr(&libc.As[Scanner](v71).F0)
	_array__grow(tags79, 1, int64(24))
	v72 = *libc.As[unsafe.Pointer](scanner_addr)
	tags80 = libc.Ptr(&libc.As[Scanner](v72).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags80).F0)
	v73 = *libc.As[unsafe.Pointer](contents)
	v74 = *libc.As[unsafe.Pointer](scanner_addr)
	tags81 = libc.Ptr(&libc.As[Scanner](v74).F0)
	size82 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags81).F1)
	v75 = *libc.As[int32](size82)
	inc = v75 + 1
	*libc.As[int32](size82) = inc
	idxprom = int64(uint64(uint32(v75)))
	arrayidx83 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v73), int(idxprom)*24))
	libc.Memmove(libc.As[byte](arrayidx83), libc.As[byte](tag), int64(24))
	v76 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol84 = libc.Ptr(&libc.As[TSLexer](v76).F1)
	*libc.As[int16](result_symbol84) = 9
	*libc.As[bool](retval) = true
	goto _return

if_end85:
	goto sw_epilog

sw_bb86:
	v77 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx87 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v77), int(int64(6))*1))
	v78 = *libc.As[byte](arrayidx87)
	loadedv88 = (v78 & 1) != 0
	if loadedv88 {
		goto if_then89
	} else {
		goto if_end91
	}

if_then89:
	v79 = *libc.As[unsafe.Pointer](scanner_addr)
	v80 = *libc.As[unsafe.Pointer](lexer_addr)
	call90 = scan_implicit_end_tag(v79, v80)
	*libc.As[bool](retval) = call90
	goto _return

if_end91:
	v81 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx92 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v81), int(int64(10))*1))
	v82 = *libc.As[byte](arrayidx92)
	loadedv93 = (v82 & 1) != 0
	if loadedv93 {
		goto land_lhs_true94
	} else {
		goto if_end124
	}

land_lhs_true94:
	v83 = *libc.As[unsafe.Pointer](scanner_addr)
	tags95 = libc.Ptr(&libc.As[Scanner](v83).F0)
	size96 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags95).F1)
	v84 = *libc.As[int32](size96)
	cmp97 = uint32(v84) > 0
	if cmp97 {
		goto land_lhs_true98
	} else {
		goto if_end124
	}

land_lhs_true98:
	v85 = *libc.As[unsafe.Pointer](scanner_addr)
	tags99 = libc.Ptr(&libc.As[Scanner](v85).F0)
	size100 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags99).F1)
	v86 = *libc.As[int32](size100)
	sub = v86 - 1
	v87 = *libc.As[unsafe.Pointer](scanner_addr)
	tags101 = libc.Ptr(&libc.As[Scanner](v87).F0)
	size102 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags101).F1)
	v88 = *libc.As[int32](size102)
	cmp103 = uint32(sub) < uint32(v88)
	if cmp103 {
		goto if_then104
	} else {
		goto if_else105
	}

if_then104:
	goto if_end106

if_else105:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 650, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan)))
	panic("unreachable")

if_end106:
	v89 = *libc.As[unsafe.Pointer](scanner_addr)
	tags107 = libc.Ptr(&libc.As[Scanner](v89).F0)
	contents108 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags107).F0)
	v90 = *libc.As[unsafe.Pointer](contents108)
	v91 = *libc.As[unsafe.Pointer](scanner_addr)
	tags109 = libc.Ptr(&libc.As[Scanner](v91).F0)
	size110 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags109).F1)
	v92 = *libc.As[int32](size110)
	sub111 = v92 - 1
	idxprom112 = int64(uint64(uint32(sub111)))
	arrayidx113 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v90), int(idxprom112)*24))
	_type = libc.Ptr(&libc.As[Tag](arrayidx113).F0)
	v93 = *libc.As[int32](_type)
	cmp114 = v93 == 124
	if cmp114 {
		goto if_then115
	} else {
		goto if_end124
	}

if_then115:
	v94 = *libc.As[unsafe.Pointer](lexer_addr)
	advance116 = libc.Ptr(&libc.As[TSLexer](v94).F2)
	v95 = *libc.As[unsafe.Pointer](advance116)
	v96 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v95)(v96, false)
	v97 = *libc.As[unsafe.Pointer](scanner_addr)
	tags117 = libc.Ptr(&libc.As[Scanner](v97).F0)
	contents118 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags117).F0)
	v98 = *libc.As[unsafe.Pointer](contents118)
	v99 = *libc.As[unsafe.Pointer](scanner_addr)
	tags119 = libc.Ptr(&libc.As[Scanner](v99).F0)
	size120 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags119).F1)
	v100 = *libc.As[int32](size120)
	dec = v100 - 1
	*libc.As[int32](size120) = dec
	idxprom121 = int64(uint64(uint32(dec)))
	arrayidx122 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v98), int(idxprom121)*24))
	v101 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol123 = libc.Ptr(&libc.As[TSLexer](v101).F1)
	*libc.As[int16](result_symbol123) = 10
	*libc.As[bool](retval) = true
	goto _return

if_end124:
	goto sw_epilog

sw_bb125:
	v102 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx126 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v102), int(int64(13))*1))
	v103 = *libc.As[byte](arrayidx126)
	loadedv127 = (v103 & 1) != 0
	if loadedv127 {
		goto if_then128
	} else {
		goto if_end131
	}

if_then128:
	v104 = *libc.As[unsafe.Pointer](lexer_addr)
	scan_js_backtick_string(v104)
	v105 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end129 = libc.Ptr(&libc.As[TSLexer](v105).F3)
	v106 = *libc.As[unsafe.Pointer](mark_end129)
	v107 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v106)(v107)
	v108 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol130 = libc.Ptr(&libc.As[TSLexer](v108).F1)
	*libc.As[int16](result_symbol130) = 13
	*libc.As[bool](retval) = true
	goto _return

if_end131:
	goto sw_epilog

sw_default:
	v109 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx132 = v109
	v110 = *libc.As[byte](arrayidx132)
	loadedv133 = (v110 & 1) != 0
	if loadedv133 {
		goto land_lhs_true137
	} else {
		goto lor_lhs_false134
	}

lor_lhs_false134:
	v111 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx135 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v111), int(int64(3))*1))
	v112 = *libc.As[byte](arrayidx135)
	loadedv136 = (v112 & 1) != 0
	if loadedv136 {
		goto land_lhs_true137
	} else {
		goto if_end147
	}

land_lhs_true137:
	v113 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx138 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v113), int(int64(7))*1))
	v114 = *libc.As[byte](arrayidx138)
	loadedv139 = (v114 & 1) != 0
	if loadedv139 {
		goto if_end147
	} else {
		goto if_then140
	}

if_then140:
	v115 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx141 = v115
	v116 = *libc.As[byte](arrayidx141)
	loadedv142 = (v116 & 1) != 0
	if loadedv142 {
		goto cond_true
	} else {
		goto cond_false
	}

cond_true:
	v117 = *libc.As[unsafe.Pointer](scanner_addr)
	v118 = *libc.As[unsafe.Pointer](lexer_addr)
	call143 = scan_start_tag_name(v117, v118)
	if call143 {
		conv = 1
	} else {
		conv = 0
	}
	cond = conv
	goto cond_end

cond_false:
	v119 = *libc.As[unsafe.Pointer](scanner_addr)
	v120 = *libc.As[unsafe.Pointer](lexer_addr)
	call144 = scan_end_tag_name(v119, v120)
	if call144 {
		conv145 = 1
	} else {
		conv145 = 0
	}
	cond = conv145
	goto cond_end

cond_end:
	tobool146 = cond != 0
	*libc.As[bool](retval) = tobool146
	goto _return

if_end147:
	goto sw_epilog

sw_epilog:
	v121 = *libc.As[byte](definitely_not_permissible_text)
	loadedv148 = (v121 & 1) != 0
	if loadedv148 {
		goto if_end155
	} else {
		goto land_lhs_true149
	}

land_lhs_true149:
	v122 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx150 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v122), int(int64(14))*1))
	v123 = *libc.As[byte](arrayidx150)
	loadedv151 = (v123 & 1) != 0
	if loadedv151 {
		goto if_then153
	} else {
		goto if_end155
	}

if_then153:
	v124 = *libc.As[unsafe.Pointer](lexer_addr)
	call154 = scan_permissible_text(v124)
	*libc.As[bool](retval) = call154
	goto _return

if_end155:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v125 = *libc.As[bool](retval)
	return v125
}
func tree_sitter_astro_external_scanner_serialize(payload unsafe.Pointer, buffer unsafe.Pointer) int32 {
	var call int32
	var v0, v1, v2 unsafe.Pointer
	var payload_addr, buffer_addr, scanner unsafe.Pointer
	_, _, _, _, _, _, _ = payload_addr, buffer_addr, scanner, v0, v1, v2, call

	payload_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	buffer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	scanner = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](buffer_addr) = buffer
	v0 = *libc.As[unsafe.Pointer](payload_addr)
	*libc.As[unsafe.Pointer](scanner) = v0
	v1 = *libc.As[unsafe.Pointer](scanner)
	v2 = *libc.As[unsafe.Pointer](buffer_addr)
	call = serialize(v1, v2)
	return call
}
func serialize(scanner unsafe.Pointer, buffer unsafe.Pointer) int32 {
	var custom_tag_name, custom_tag_name35 unsafe.Pointer
	var tag, arrayidx12 unsafe.Pointer
	var tags, tags1, tags10 unsafe.Pointer
	var cmp, cmp8, cmp13, cmp16, cmp21, cmp40 bool
	var conv, v7, v8, v11, v33, inc50 int16
	var tag_count, serialized_tag_count unsafe.Pointer
	var v1, v3, cond, v5, v6, conv5, conv6, conv7, v12, v13, v14, v15, add19, v16, add20, v17, v19, inc, v20, v22, inc30, v24, v26, v27, v28, add38, v29, add39, v30, v32, inc46, v35 int32
	var size3, name_length, size, size2, _type, size15, type25, type44 unsafe.Pointer
	var idxprom, conv4, add, idxprom11, idxprom27, idxprom31, idxprom33, conv37, idxprom47 int64
	var conv26, conv29, conv45 byte
	var arrayidx, arrayidx28, arrayidx32, arrayidx34, arrayidx48, arrayidx51 unsafe.Pointer
	var v0, v2, v4, v9, v10, v18, v21, v23, v25, call, v31, v34 unsafe.Pointer
	var scanner_addr, buffer_addr, contents, contents36 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = scanner_addr, buffer_addr, tag_count, serialized_tag_count, size3, tag, name_length, v0, tags, size, v1, cmp, v2, tags1, size2, v3, cond, conv, v4, v5, idxprom, arrayidx, v6, conv4, add, conv5, v7, conv6, v8, conv7, cmp8, v9, tags10, contents, v10, v11, idxprom11, arrayidx12, _type, v12, cmp13, custom_tag_name, size15, v13, v14, cmp16, v15, add19, v16, add20, cmp21, type25, v17, conv26, v18, v19, inc, idxprom27, arrayidx28, v20, conv29, v21, v22, inc30, idxprom31, arrayidx32, v23, v24, idxprom33, arrayidx34, custom_tag_name35, contents36, v25, v26, conv37, call, v27, v28, add38, v29, add39, cmp40, type44, v30, conv45, v31, v32, inc46, idxprom47, arrayidx48, v33, inc50, v34, arrayidx51, v35

	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	buffer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	tag_count = libc.Ptr(&new(struct {
		_ [0]uint64
		v int16
		b byte
	}).v)
	serialized_tag_count = libc.Ptr(&new(struct {
		_ [0]uint64
		v int16
		b byte
	}).v)
	size3 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	name_length = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](buffer_addr) = buffer
	v0 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v0).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v1 = *libc.As[int32](size)
	cmp = uint32(v1) > 65535
	if cmp {
		goto cond_true
	} else {
		goto cond_false
	}

cond_true:
	cond = 65535
	goto cond_end

cond_false:
	v2 = *libc.As[unsafe.Pointer](scanner_addr)
	tags1 = libc.Ptr(&libc.As[Scanner](v2).F0)
	size2 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F1)
	v3 = *libc.As[int32](size2)
	cond = v3
	goto cond_end

cond_end:
	conv = int16(cond)
	*libc.As[int16](tag_count) = conv
	*libc.As[int16](serialized_tag_count) = 0
	*libc.As[int32](size3) = 2
	v4 = *libc.As[unsafe.Pointer](buffer_addr)
	v5 = *libc.As[int32](size3)
	idxprom = int64(uint64(uint32(v5)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v4), int(idxprom)*1))
	libc.Memmove(libc.As[byte](arrayidx), libc.As[byte](tag_count), int64(2))
	v6 = *libc.As[int32](size3)
	conv4 = int64(uint64(uint32(v6)))
	add = conv4 + int64(2)
	conv5 = int32(add)
	*libc.As[int32](size3) = conv5
	goto for_cond

for_cond:
	v7 = *libc.As[int16](serialized_tag_count)
	conv6 = int32(uint32(uint16(v7)))
	v8 = *libc.As[int16](tag_count)
	conv7 = int32(uint32(uint16(v8)))
	cmp8 = conv6 < conv7
	if cmp8 {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v9 = *libc.As[unsafe.Pointer](scanner_addr)
	tags10 = libc.Ptr(&libc.As[Scanner](v9).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags10).F0)
	v10 = *libc.As[unsafe.Pointer](contents)
	v11 = *libc.As[int16](serialized_tag_count)
	idxprom11 = int64(uint64(uint16(v11)))
	arrayidx12 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v10), int(idxprom11)*24))
	libc.Memmove(libc.As[byte](tag), libc.As[byte](arrayidx12), int64(24))
	_type = libc.Ptr(&libc.As[Tag](tag).F0)
	v12 = *libc.As[int32](_type)
	cmp13 = v12 == 126
	if cmp13 {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	custom_tag_name = libc.Ptr(&libc.As[Tag](tag).F1)
	size15 = libc.Ptr(&libc.As[String](custom_tag_name).F1)
	v13 = *libc.As[int32](size15)
	*libc.As[int32](name_length) = v13
	v14 = *libc.As[int32](name_length)
	cmp16 = uint32(v14) > 255
	if cmp16 {
		goto if_then18
	} else {
		goto if_end
	}

if_then18:
	*libc.As[int32](name_length) = 255
	goto if_end

if_end:
	v15 = *libc.As[int32](size3)
	add19 = v15 + 2
	v16 = *libc.As[int32](name_length)
	add20 = add19 + v16
	cmp21 = uint32(add20) >= 1024
	if cmp21 {
		goto if_then23
	} else {
		goto if_end24
	}

if_then23:
	goto for_end

if_end24:
	type25 = libc.Ptr(&libc.As[Tag](tag).F0)
	v17 = *libc.As[int32](type25)
	conv26 = byte(v17)
	v18 = *libc.As[unsafe.Pointer](buffer_addr)
	v19 = *libc.As[int32](size3)
	inc = v19 + 1
	*libc.As[int32](size3) = inc
	idxprom27 = int64(uint64(uint32(v19)))
	arrayidx28 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v18), int(idxprom27)*1))
	*libc.As[byte](arrayidx28) = conv26
	v20 = *libc.As[int32](name_length)
	conv29 = byte(v20)
	v21 = *libc.As[unsafe.Pointer](buffer_addr)
	v22 = *libc.As[int32](size3)
	inc30 = v22 + 1
	*libc.As[int32](size3) = inc30
	idxprom31 = int64(uint64(uint32(v22)))
	arrayidx32 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v21), int(idxprom31)*1))
	*libc.As[byte](arrayidx32) = conv29
	v23 = *libc.As[unsafe.Pointer](buffer_addr)
	v24 = *libc.As[int32](size3)
	idxprom33 = int64(uint64(uint32(v24)))
	arrayidx34 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v23), int(idxprom33)*1))
	custom_tag_name35 = libc.Ptr(&libc.As[Tag](tag).F1)
	contents36 = libc.Ptr(&libc.As[String](custom_tag_name35).F0)
	v25 = *libc.As[unsafe.Pointer](contents36)
	v26 = *libc.As[int32](name_length)
	conv37 = int64(uint64(uint32(v26)))
	call = libc.Ptr(libc.Strncpy(libc.As[byte](arrayidx34), libc.As[byte](v25), conv37))
	v27 = *libc.As[int32](name_length)
	v28 = *libc.As[int32](size3)
	add38 = v28 + v27
	*libc.As[int32](size3) = add38
	goto if_end49

if_else:
	v29 = *libc.As[int32](size3)
	add39 = v29 + 1
	cmp40 = uint32(add39) >= 1024
	if cmp40 {
		goto if_then42
	} else {
		goto if_end43
	}

if_then42:
	goto for_end

if_end43:
	type44 = libc.Ptr(&libc.As[Tag](tag).F0)
	v30 = *libc.As[int32](type44)
	conv45 = byte(v30)
	v31 = *libc.As[unsafe.Pointer](buffer_addr)
	v32 = *libc.As[int32](size3)
	inc46 = v32 + 1
	*libc.As[int32](size3) = inc46
	idxprom47 = int64(uint64(uint32(v32)))
	arrayidx48 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v31), int(idxprom47)*1))
	*libc.As[byte](arrayidx48) = conv45
	goto if_end49

if_end49:
	goto for_inc

for_inc:
	v33 = *libc.As[int16](serialized_tag_count)
	inc50 = v33 + 1
	*libc.As[int16](serialized_tag_count) = inc50
	goto for_cond

for_end:
	v34 = *libc.As[unsafe.Pointer](buffer_addr)
	arrayidx51 = v34
	libc.Memmove(libc.As[byte](arrayidx51), libc.As[byte](serialized_tag_count), int64(2))
	v35 = *libc.As[int32](size3)
	return v35
}
func tree_sitter_astro_external_scanner_deserialize(payload unsafe.Pointer, buffer unsafe.Pointer, length int32) {
	var v3 int32
	var length_addr unsafe.Pointer
	var v0, v1, v2 unsafe.Pointer
	var payload_addr, buffer_addr, scanner unsafe.Pointer
	_, _, _, _, _, _, _, _ = payload_addr, buffer_addr, length_addr, scanner, v0, v1, v2, v3

	payload_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	buffer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	length_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	scanner = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](buffer_addr) = buffer
	*libc.As[int32](length_addr) = length
	v0 = *libc.As[unsafe.Pointer](payload_addr)
	*libc.As[unsafe.Pointer](scanner) = v0
	v1 = *libc.As[unsafe.Pointer](scanner)
	v2 = *libc.As[unsafe.Pointer](buffer_addr)
	v3 = *libc.As[int32](length_addr)
	deserialize(v1, v2, v3)
}
func deserialize(scanner unsafe.Pointer, buffer unsafe.Pointer, length int32) {
	var custom_tag_name, custom_tag_name39, custom_tag_name41 unsafe.Pointer
	var tag, tmp, arrayidx, arrayidx55, arrayidx71 unsafe.Pointer
	var tags, tags1, tags2, tags14, tags48, tags49, tags51, tags64, tags65, tags67 unsafe.Pointer
	var cmp, cmp4, cmp17, cmp22, cmp30, cmp61 bool
	var v16, v17, v19, conv36, v27, v28, v32, v33, v42 int16
	var tag_count, serialized_tag_count, name_length unsafe.Pointer
	var v0, v2, v5, v6, inc, v8, v10, v11, conv8, v13, v14, conv13, conv15, conv16, v18, conv21, v21, inc25, conv28, v23, v25, inc33, conv37, conv38, v31, conv46, v34, add47, v39, inc53, v40, inc57, v41, conv60, v47, inc69, v48, inc73 int32
	var length_addr, i, size5, iter, size, size3, _type, type29, size40, size52, size68 unsafe.Pointer
	var idxprom, idxprom6, conv, add, idxprom9, conv11, add12, idxprom26, idxprom34, idxprom43, conv45, idxprom54, idxprom70 int64
	var v22, v26 byte
	var arrayidx7, arrayidx10, arrayidx27, arrayidx35, arrayidx44 unsafe.Pointer
	var v1, v3, v4, v7, v9, v12, v15, v20, v24, v29, v30, v35, v36, v37, v38, v43, v44, v45, v46 unsafe.Pointer
	var scanner_addr, buffer_addr, contents, contents42, contents50, contents66 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = scanner_addr, buffer_addr, length_addr, i, size5, tag_count, serialized_tag_count, iter, tag, name_length, tmp, v0, v1, tags, size, v2, cmp, v3, tags1, contents, v4, v5, idxprom, arrayidx, v6, inc, v7, tags2, size3, v8, cmp4, v9, v10, idxprom6, arrayidx7, v11, conv, add, conv8, v12, v13, idxprom9, arrayidx10, v14, conv11, add12, conv13, v15, tags14, v16, conv15, v17, conv16, cmp17, v18, v19, conv21, cmp22, v20, v21, inc25, idxprom26, arrayidx27, v22, conv28, _type, type29, v23, cmp30, v24, v25, inc33, idxprom34, arrayidx35, v26, conv36, custom_tag_name, v27, conv37, v28, conv38, custom_tag_name39, size40, custom_tag_name41, contents42, v29, v30, v31, idxprom43, arrayidx44, v32, conv45, v33, conv46, v34, add47, v35, tags48, v36, tags49, contents50, v37, v38, tags51, size52, v39, inc53, idxprom54, arrayidx55, v40, inc57, v41, v42, conv60, cmp61, v43, tags64, v44, tags65, contents66, v45, v46, tags67, size68, v47, inc69, idxprom70, arrayidx71, v48, inc73

	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	buffer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	length_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	size5 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	tag_count = libc.Ptr(&new(struct {
		_ [0]uint64
		v int16
		b byte
	}).v)
	serialized_tag_count = libc.Ptr(&new(struct {
		_ [0]uint64
		v int16
		b byte
	}).v)
	iter = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	name_length = libc.Ptr(&new(struct {
		_ [0]uint64
		v int16
		b byte
	}).v)
	tmp = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](buffer_addr) = buffer
	*libc.As[int32](length_addr) = length
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v0 = *libc.As[int32](i)
	v1 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v1).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v2 = *libc.As[int32](size)
	cmp = uint32(v0) < uint32(v2)
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v3 = *libc.As[unsafe.Pointer](scanner_addr)
	tags1 = libc.Ptr(&libc.As[Scanner](v3).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F0)
	v4 = *libc.As[unsafe.Pointer](contents)
	v5 = *libc.As[int32](i)
	idxprom = int64(uint64(uint32(v5)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v4), int(idxprom)*24))
	tag_free(arrayidx)
	goto for_inc

for_inc:
	v6 = *libc.As[int32](i)
	inc = v6 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	v7 = *libc.As[unsafe.Pointer](scanner_addr)
	tags2 = libc.Ptr(&libc.As[Scanner](v7).F0)
	size3 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags2).F1)
	*libc.As[int32](size3) = 0
	v8 = *libc.As[int32](length_addr)
	cmp4 = uint32(v8) > 0
	if cmp4 {
		goto if_then
	} else {
		goto if_end76
	}

if_then:
	*libc.As[int32](size5) = 0
	*libc.As[int16](tag_count) = 0
	*libc.As[int16](serialized_tag_count) = 0
	v9 = *libc.As[unsafe.Pointer](buffer_addr)
	v10 = *libc.As[int32](size5)
	idxprom6 = int64(uint64(uint32(v10)))
	arrayidx7 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v9), int(idxprom6)*1))
	libc.Memmove(libc.As[byte](serialized_tag_count), libc.As[byte](arrayidx7), int64(2))
	v11 = *libc.As[int32](size5)
	conv = int64(uint64(uint32(v11)))
	add = conv + int64(2)
	conv8 = int32(add)
	*libc.As[int32](size5) = conv8
	v12 = *libc.As[unsafe.Pointer](buffer_addr)
	v13 = *libc.As[int32](size5)
	idxprom9 = int64(uint64(uint32(v13)))
	arrayidx10 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v12), int(idxprom9)*1))
	libc.Memmove(libc.As[byte](tag_count), libc.As[byte](arrayidx10), int64(2))
	v14 = *libc.As[int32](size5)
	conv11 = int64(uint64(uint32(v14)))
	add12 = conv11 + int64(2)
	conv13 = int32(add12)
	*libc.As[int32](size5) = conv13
	v15 = *libc.As[unsafe.Pointer](scanner_addr)
	tags14 = libc.Ptr(&libc.As[Scanner](v15).F0)
	v16 = *libc.As[int16](tag_count)
	conv15 = int32(uint32(uint16(v16)))
	_array__reserve(tags14, int64(24), conv15)
	v17 = *libc.As[int16](tag_count)
	conv16 = int32(uint32(uint16(v17)))
	cmp17 = conv16 > 0
	if cmp17 {
		goto if_then19
	} else {
		goto if_end75
	}

if_then19:
	*libc.As[int32](iter) = 0
	*libc.As[int32](iter) = 0
	goto for_cond20

for_cond20:
	v18 = *libc.As[int32](iter)
	v19 = *libc.As[int16](serialized_tag_count)
	conv21 = int32(uint32(uint16(v19)))
	cmp22 = uint32(v18) < uint32(conv21)
	if cmp22 {
		goto for_body24
	} else {
		goto for_end58
	}

for_body24:
	tag_new(tag)
	v20 = *libc.As[unsafe.Pointer](buffer_addr)
	v21 = *libc.As[int32](size5)
	inc25 = v21 + 1
	*libc.As[int32](size5) = inc25
	idxprom26 = int64(uint64(uint32(v21)))
	arrayidx27 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v20), int(idxprom26)*1))
	v22 = *libc.As[byte](arrayidx27)
	conv28 = int32(int8(v22))
	_type = libc.Ptr(&libc.As[Tag](tag).F0)
	*libc.As[int32](_type) = conv28
	type29 = libc.Ptr(&libc.As[Tag](tag).F0)
	v23 = *libc.As[int32](type29)
	cmp30 = v23 == 126
	if cmp30 {
		goto if_then32
	} else {
		goto if_end
	}

if_then32:
	v24 = *libc.As[unsafe.Pointer](buffer_addr)
	v25 = *libc.As[int32](size5)
	inc33 = v25 + 1
	*libc.As[int32](size5) = inc33
	idxprom34 = int64(uint64(uint32(v25)))
	arrayidx35 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v24), int(idxprom34)*1))
	v26 = *libc.As[byte](arrayidx35)
	conv36 = int16(uint16(v26))
	*libc.As[int16](name_length) = conv36
	custom_tag_name = libc.Ptr(&libc.As[Tag](tag).F1)
	v27 = *libc.As[int16](name_length)
	conv37 = int32(uint32(uint16(v27)))
	_array__reserve(custom_tag_name, int64(1), conv37)
	v28 = *libc.As[int16](name_length)
	conv38 = int32(uint32(uint16(v28)))
	custom_tag_name39 = libc.Ptr(&libc.As[Tag](tag).F1)
	size40 = libc.Ptr(&libc.As[String](custom_tag_name39).F1)
	*libc.As[int32](size40) = conv38
	custom_tag_name41 = libc.Ptr(&libc.As[Tag](tag).F1)
	contents42 = libc.Ptr(&libc.As[String](custom_tag_name41).F0)
	v29 = *libc.As[unsafe.Pointer](contents42)
	v30 = *libc.As[unsafe.Pointer](buffer_addr)
	v31 = *libc.As[int32](size5)
	idxprom43 = int64(uint64(uint32(v31)))
	arrayidx44 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v30), int(idxprom43)*1))
	v32 = *libc.As[int16](name_length)
	conv45 = int64(uint64(uint16(v32)))
	libc.Memmove(libc.As[byte](v29), libc.As[byte](arrayidx44), conv45)
	v33 = *libc.As[int16](name_length)
	conv46 = int32(uint32(uint16(v33)))
	v34 = *libc.As[int32](size5)
	add47 = v34 + conv46
	*libc.As[int32](size5) = add47
	goto if_end

if_end:
	v35 = *libc.As[unsafe.Pointer](scanner_addr)
	tags48 = libc.Ptr(&libc.As[Scanner](v35).F0)
	_array__grow(tags48, 1, int64(24))
	v36 = *libc.As[unsafe.Pointer](scanner_addr)
	tags49 = libc.Ptr(&libc.As[Scanner](v36).F0)
	contents50 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags49).F0)
	v37 = *libc.As[unsafe.Pointer](contents50)
	v38 = *libc.As[unsafe.Pointer](scanner_addr)
	tags51 = libc.Ptr(&libc.As[Scanner](v38).F0)
	size52 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags51).F1)
	v39 = *libc.As[int32](size52)
	inc53 = v39 + 1
	*libc.As[int32](size52) = inc53
	idxprom54 = int64(uint64(uint32(v39)))
	arrayidx55 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v37), int(idxprom54)*24))
	libc.Memmove(libc.As[byte](arrayidx55), libc.As[byte](tag), int64(24))
	goto for_inc56

for_inc56:
	v40 = *libc.As[int32](iter)
	inc57 = v40 + 1
	*libc.As[int32](iter) = inc57
	goto for_cond20

for_end58:
	goto for_cond59

for_cond59:
	v41 = *libc.As[int32](iter)
	v42 = *libc.As[int16](tag_count)
	conv60 = int32(uint32(uint16(v42)))
	cmp61 = uint32(v41) < uint32(conv60)
	if cmp61 {
		goto for_body63
	} else {
		goto for_end74
	}

for_body63:
	v43 = *libc.As[unsafe.Pointer](scanner_addr)
	tags64 = libc.Ptr(&libc.As[Scanner](v43).F0)
	_array__grow(tags64, 1, int64(24))
	v44 = *libc.As[unsafe.Pointer](scanner_addr)
	tags65 = libc.Ptr(&libc.As[Scanner](v44).F0)
	contents66 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags65).F0)
	v45 = *libc.As[unsafe.Pointer](contents66)
	v46 = *libc.As[unsafe.Pointer](scanner_addr)
	tags67 = libc.Ptr(&libc.As[Scanner](v46).F0)
	size68 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags67).F1)
	v47 = *libc.As[int32](size68)
	inc69 = v47 + 1
	*libc.As[int32](size68) = inc69
	idxprom70 = int64(uint64(uint32(v47)))
	arrayidx71 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v45), int(idxprom70)*24))
	tag_new(tmp)
	libc.Memmove(libc.As[byte](arrayidx71), libc.As[byte](tmp), int64(24))
	goto for_inc72

for_inc72:
	v48 = *libc.As[int32](iter)
	inc73 = v48 + 1
	*libc.As[int32](iter) = inc73
	goto for_cond59

for_end74:
	goto if_end75

if_end75:
	goto if_end76

if_end76:
}
func tree_sitter_astro_external_scanner_destroy(payload unsafe.Pointer) {
	var arrayidx unsafe.Pointer
	var tags, tags1, tags2 unsafe.Pointer
	var cmp bool
	var v1, v3, v6, v7, inc int32
	var i, size unsafe.Pointer
	var idxprom int64
	var v0, v2, v4, v5, v8, v9 unsafe.Pointer
	var payload_addr, scanner, contents unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = payload_addr, scanner, i, v0, v1, v2, tags, size, v3, cmp, v4, tags1, contents, v5, v6, idxprom, arrayidx, v7, inc, v8, tags2, v9

	payload_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	scanner = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
	v0 = *libc.As[unsafe.Pointer](payload_addr)
	*libc.As[unsafe.Pointer](scanner) = v0
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v1 = *libc.As[int32](i)
	v2 = *libc.As[unsafe.Pointer](scanner)
	tags = libc.Ptr(&libc.As[Scanner](v2).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v3 = *libc.As[int32](size)
	cmp = uint32(v1) < uint32(v3)
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v4 = *libc.As[unsafe.Pointer](scanner)
	tags1 = libc.Ptr(&libc.As[Scanner](v4).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F0)
	v5 = *libc.As[unsafe.Pointer](contents)
	v6 = *libc.As[int32](i)
	idxprom = int64(uint64(uint32(v6)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v5), int(idxprom)*24))
	tag_free(arrayidx)
	goto for_inc

for_inc:
	v7 = *libc.As[int32](i)
	inc = v7 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	v8 = *libc.As[unsafe.Pointer](scanner)
	tags2 = libc.Ptr(&libc.As[Scanner](v8).F0)
	_array__delete(tags2)
	v9 = *libc.As[unsafe.Pointer](scanner)
	libc.Free(libc.As[byte](v9))
}
func tag_free(tag unsafe.Pointer) {
	var custom_tag_name unsafe.Pointer
	var cmp bool
	var v1 int32
	var _type unsafe.Pointer
	var v0, v2 unsafe.Pointer
	var tag_addr unsafe.Pointer
	_, _, _, _, _, _, _ = tag_addr, v0, _type, v1, cmp, v2, custom_tag_name

	tag_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](tag_addr) = tag
	v0 = *libc.As[unsafe.Pointer](tag_addr)
	_type = libc.Ptr(&libc.As[Tag](v0).F0)
	v1 = *libc.As[int32](_type)
	cmp = v1 == 126
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v2 = *libc.As[unsafe.Pointer](tag_addr)
	custom_tag_name = libc.Ptr(&libc.As[Tag](v2).F1)
	_array__delete(custom_tag_name)
	goto if_end

if_end:
}
func _array__delete(self unsafe.Pointer) {
	var tobool bool
	var size, capacity unsafe.Pointer
	var v0, v1, v2, v3, v4, v5, v6 unsafe.Pointer
	var self_addr, contents, contents1, contents2 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = self_addr, v0, contents, v1, tobool, v2, contents1, v3, v4, contents2, v5, size, v6, capacity

	self_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](self_addr) = self
	v0 = *libc.As[unsafe.Pointer](self_addr)
	contents = libc.Ptr(&libc.As[Array](v0).F0)
	v1 = *libc.As[unsafe.Pointer](contents)
	tobool = uintptr(unsafe.Pointer(v1)) != uintptr(unsafe.Pointer(nil))
	if tobool {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v2 = *libc.As[unsafe.Pointer](self_addr)
	contents1 = libc.Ptr(&libc.As[Array](v2).F0)
	v3 = *libc.As[unsafe.Pointer](contents1)
	libc.Free(libc.As[byte](v3))
	v4 = *libc.As[unsafe.Pointer](self_addr)
	contents2 = libc.Ptr(&libc.As[Array](v4).F0)
	*libc.As[unsafe.Pointer](contents2) = nil
	v5 = *libc.As[unsafe.Pointer](self_addr)
	size = libc.Ptr(&libc.As[Array](v5).F1)
	*libc.As[int32](size) = 0
	v6 = *libc.As[unsafe.Pointer](self_addr)
	capacity = libc.Ptr(&libc.As[Array](v6).F2)
	*libc.As[int32](capacity) = 0
	goto if_end

if_end:
}
func tree_sitter_astro() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_astro_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp27, cmp31, cmp35, cmp39, cmp43, cmp45, cmp49, cmp51, cmp54, loadedv58, cmp60, cmp64, cmp68, cmp72, cmp75, cmp78, cmp82, cmp85, cmp88, loadedv92, cmp94, cmp98, cmp101, cmp104, cmp108, loadedv112, cmp114, cmp118, cmp121, cmp124, cmp127, loadedv131, cmp133, cmp137, cmp140, cmp143, cmp147, loadedv151, cmp153, loadedv157, cmp159, loadedv163, cmp165, loadedv169, cmp171, loadedv175, cmp177, cmp181, cmp184, cmp187, loadedv191, cmp193, cmp197, cmp201, cmp205, cmp209, cmp212, cmp215, cmp219, cmp222, cmp225, cmp228, loadedv232, cmp234, loadedv238, cmp240, cmp243, loadedv247, cmp249, cmp252, loadedv256, cmp258, cmp261, loadedv265, cmp267, cmp270, loadedv274, cmp276, cmp279, loadedv283, cmp285, cmp288, cmp292, cmp295, loadedv299, cmp301, cmp304, loadedv308, cmp310, cmp313, cmp316, cmp320, cmp323, cmp326, cmp329, cmp332, loadedv336, cmp338, cmp341, cmp344, cmp348, cmp351, loadedv355, cmp357, cmp360, cmp363, cmp366, cmp369, cmp372, loadedv376, loadedv378, cmp381, cmp385, cmp389, cmp392, cmp395, cmp399, cmp402, cmp405, cmp408, loadedv412, loadedv414, cmp417, cmp421, cmp424, cmp427, cmp431, cmp434, cmp437, cmp440, loadedv444, loadedv446, loadedv450, cmp454, cmp457, cmp460, cmp464, cmp467, loadedv471, cmp475, cmp478, loadedv482, loadedv486, loadedv490, cmp494, cmp498, loadedv502, loadedv506, loadedv510, loadedv514, cmp518, cmp521, cmp524, cmp527, cmp530, cmp533, cmp536, cmp539, cmp542, loadedv546, cmp550, cmp553, cmp556, cmp559, cmp562, cmp565, cmp568, cmp571, loadedv575, loadedv579, cmp583, loadedv587, cmp591, cmp595, cmp598, loadedv602, cmp606, cmp610, cmp613, loadedv617, cmp621, cmp625, cmp628, loadedv632, cmp636, cmp640, cmp643, loadedv647, cmp651, cmp655, cmp658, cmp661, cmp664, cmp667, cmp670, loadedv674, cmp678, cmp682, cmp685, cmp688, cmp691, cmp694, cmp697, loadedv701, cmp705, cmp709, cmp712, cmp715, cmp718, cmp721, cmp724, loadedv728, cmp732, cmp736, cmp739, cmp742, cmp745, cmp748, cmp751, loadedv755, cmp759, cmp763, cmp766, cmp769, cmp772, cmp775, cmp778, loadedv782, cmp786, cmp790, cmp793, cmp796, cmp799, loadedv803, cmp807, cmp811, cmp814, cmp817, cmp820, loadedv824, cmp828, cmp832, cmp835, cmp838, cmp841, loadedv845, cmp849, cmp853, cmp856, cmp859, cmp862, loadedv866, cmp870, cmp874, cmp877, cmp880, cmp883, loadedv887, cmp891, cmp895, cmp898, cmp901, cmp904, loadedv908, cmp912, cmp916, cmp919, cmp922, cmp925, loadedv929, cmp933, cmp937, cmp940, cmp943, cmp946, loadedv950, cmp954, cmp958, cmp961, cmp964, cmp967, loadedv971, cmp975, cmp979, cmp982, cmp985, cmp988, loadedv992, cmp996, cmp1000, cmp1003, cmp1006, cmp1009, loadedv1013, cmp1017, cmp1021, cmp1024, cmp1027, cmp1030, loadedv1034, cmp1038, cmp1042, cmp1045, cmp1048, cmp1051, loadedv1055, cmp1059, cmp1063, cmp1066, cmp1069, cmp1072, loadedv1076, cmp1080, cmp1084, cmp1087, cmp1090, cmp1093, loadedv1097, cmp1101, cmp1105, cmp1108, cmp1111, cmp1114, loadedv1118, cmp1122, cmp1126, cmp1129, cmp1132, cmp1135, loadedv1139, cmp1143, cmp1147, cmp1150, cmp1153, cmp1156, loadedv1160, cmp1164, cmp1168, cmp1171, cmp1174, cmp1177, loadedv1181, cmp1185, cmp1189, cmp1192, cmp1195, cmp1198, loadedv1202, cmp1206, cmp1210, cmp1213, cmp1216, cmp1219, loadedv1223, cmp1227, cmp1231, cmp1234, cmp1237, cmp1240, loadedv1244, cmp1248, cmp1252, cmp1255, cmp1258, cmp1261, loadedv1265, cmp1269, cmp1273, cmp1276, cmp1279, cmp1282, loadedv1286, cmp1290, cmp1294, cmp1297, cmp1300, cmp1303, loadedv1307, cmp1311, cmp1315, cmp1318, cmp1321, cmp1324, loadedv1328, cmp1332, cmp1336, cmp1339, cmp1342, cmp1345, loadedv1349, cmp1353, cmp1357, cmp1360, cmp1363, cmp1366, loadedv1370, cmp1374, cmp1378, cmp1381, cmp1384, cmp1387, loadedv1391, loadedv1395, cmp1399, cmp1402, cmp1405, cmp1409, cmp1412, loadedv1416, cmp1420, cmp1423, loadedv1427, loadedv1431, cmp1435, cmp1438, cmp1441, cmp1445, cmp1448, loadedv1452, cmp1456, cmp1459, loadedv1463, cmp1467, cmp1471, cmp1474, cmp1477, cmp1481, cmp1484, cmp1487, cmp1490, cmp1493, loadedv1497, cmp1501, cmp1505, cmp1508, cmp1511, cmp1515, cmp1518, cmp1521, cmp1524, cmp1527, loadedv1531, cmp1535, cmp1538, cmp1541, cmp1545, cmp1548, cmp1551, cmp1554, cmp1557, loadedv1561, loadedv1565, loadedv1569, loadedv1573, loadedv1577, v732 bool
	var retval unsafe.Pointer
	var v9 int16
	var state_addr, result_symbol, result_symbol448, result_symbol452, result_symbol473, result_symbol484, result_symbol488, result_symbol492, result_symbol504, result_symbol508, result_symbol512, result_symbol516, result_symbol548, result_symbol577, result_symbol581, result_symbol589, result_symbol604, result_symbol619, result_symbol634, result_symbol649, result_symbol676, result_symbol703, result_symbol730, result_symbol757, result_symbol784, result_symbol805, result_symbol826, result_symbol847, result_symbol868, result_symbol889, result_symbol910, result_symbol931, result_symbol952, result_symbol973, result_symbol994, result_symbol1015, result_symbol1036, result_symbol1057, result_symbol1078, result_symbol1099, result_symbol1120, result_symbol1141, result_symbol1162, result_symbol1183, result_symbol1204, result_symbol1225, result_symbol1246, result_symbol1267, result_symbol1288, result_symbol1309, result_symbol1330, result_symbol1351, result_symbol1372, result_symbol1393, result_symbol1397, result_symbol1418, result_symbol1429, result_symbol1433, result_symbol1454, result_symbol1465, result_symbol1499, result_symbol1533, result_symbol1563, result_symbol1567, result_symbol1571, result_symbol1575 unsafe.Pointer
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v27, v28, v29, v30, v31, v32, v33, v34, v35, v37, v38, v39, v40, v41, v43, v44, v45, v46, v47, v49, v50, v51, v52, v53, v55, v57, v59, v61, v63, v64, v65, v66, v68, v69, v70, v71, v72, v73, v74, v75, v76, v77, v78, v80, v82, v83, v85, v86, v88, v89, v91, v92, v94, v95, v97, v98, v99, v100, v102, v103, v105, v106, v107, v108, v109, v110, v111, v112, v114, v115, v116, v117, v118, v120, v121, v122, v123, v124, v125, v128, v129, v130, v131, v132, v133, v134, v135, v136, v139, v140, v141, v142, v143, v144, v145, v146, v162, v163, v164, v165, v166, v172, v173, v189, v190, v211, v212, v213, v214, v215, v216, v217, v218, v219, v225, v226, v227, v228, v229, v230, v231, v232, v243, v249, v250, v251, v257, v258, v259, v265, v266, v267, v273, v274, v275, v281, v282, v283, v284, v285, v286, v287, v293, v294, v295, v296, v297, v298, v299, v305, v306, v307, v308, v309, v310, v311, v317, v318, v319, v320, v321, v322, v323, v329, v330, v331, v332, v333, v334, v335, v341, v342, v343, v344, v345, v351, v352, v353, v354, v355, v361, v362, v363, v364, v365, v371, v372, v373, v374, v375, v381, v382, v383, v384, v385, v391, v392, v393, v394, v395, v401, v402, v403, v404, v405, v411, v412, v413, v414, v415, v421, v422, v423, v424, v425, v431, v432, v433, v434, v435, v441, v442, v443, v444, v445, v451, v452, v453, v454, v455, v461, v462, v463, v464, v465, v471, v472, v473, v474, v475, v481, v482, v483, v484, v485, v491, v492, v493, v494, v495, v501, v502, v503, v504, v505, v511, v512, v513, v514, v515, v521, v522, v523, v524, v525, v531, v532, v533, v534, v535, v541, v542, v543, v544, v545, v551, v552, v553, v554, v555, v561, v562, v563, v564, v565, v571, v572, v573, v574, v575, v581, v582, v583, v584, v585, v591, v592, v593, v594, v595, v601, v602, v603, v604, v605, v611, v612, v613, v614, v615, v621, v622, v623, v624, v625, v636, v637, v638, v639, v640, v646, v647, v658, v659, v660, v661, v662, v668, v669, v675, v676, v677, v678, v679, v680, v681, v682, v683, v689, v690, v691, v692, v693, v694, v695, v696, v697, v703, v704, v705, v706, v707, v708, v709, v710 int32
	var lookahead, lookahead1 unsafe.Pointer
	var v3, storedv, v10, v26, v36, v42, v48, v54, v56, v58, v60, v62, v67, v79, v81, v84, v87, v90, v93, v96, v101, v104, v113, v119, v126, v127, v137, v138, v147, v152, v157, v167, v174, v179, v184, v191, v196, v201, v206, v220, v233, v238, v244, v252, v260, v268, v276, v288, v300, v312, v324, v336, v346, v356, v366, v376, v386, v396, v406, v416, v426, v436, v446, v456, v466, v476, v486, v496, v506, v516, v526, v536, v546, v556, v566, v576, v586, v596, v606, v616, v626, v631, v641, v648, v653, v663, v670, v684, v698, v711, v716, v721, v726, v731 byte
	var result, local_skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v148, v149, v150, v151, v153, v154, v155, v156, v158, v159, v160, v161, v168, v169, v170, v171, v175, v176, v177, v178, v180, v181, v182, v183, v185, v186, v187, v188, v192, v193, v194, v195, v197, v198, v199, v200, v202, v203, v204, v205, v207, v208, v209, v210, v221, v222, v223, v224, v234, v235, v236, v237, v239, v240, v241, v242, v245, v246, v247, v248, v253, v254, v255, v256, v261, v262, v263, v264, v269, v270, v271, v272, v277, v278, v279, v280, v289, v290, v291, v292, v301, v302, v303, v304, v313, v314, v315, v316, v325, v326, v327, v328, v337, v338, v339, v340, v347, v348, v349, v350, v357, v358, v359, v360, v367, v368, v369, v370, v377, v378, v379, v380, v387, v388, v389, v390, v397, v398, v399, v400, v407, v408, v409, v410, v417, v418, v419, v420, v427, v428, v429, v430, v437, v438, v439, v440, v447, v448, v449, v450, v457, v458, v459, v460, v467, v468, v469, v470, v477, v478, v479, v480, v487, v488, v489, v490, v497, v498, v499, v500, v507, v508, v509, v510, v517, v518, v519, v520, v527, v528, v529, v530, v537, v538, v539, v540, v547, v548, v549, v550, v557, v558, v559, v560, v567, v568, v569, v570, v577, v578, v579, v580, v587, v588, v589, v590, v597, v598, v599, v600, v607, v608, v609, v610, v617, v618, v619, v620, v627, v628, v629, v630, v632, v633, v634, v635, v642, v643, v644, v645, v649, v650, v651, v652, v654, v655, v656, v657, v664, v665, v666, v667, v671, v672, v673, v674, v685, v686, v687, v688, v699, v700, v701, v702, v712, v713, v714, v715, v717, v718, v719, v720, v722, v723, v724, v725, v727, v728, v729, v730 unsafe.Pointer
	var lexer_addr, local_advance, eof2, mark_end, mark_end449, mark_end453, mark_end474, mark_end485, mark_end489, mark_end493, mark_end505, mark_end509, mark_end513, mark_end517, mark_end549, mark_end578, mark_end582, mark_end590, mark_end605, mark_end620, mark_end635, mark_end650, mark_end677, mark_end704, mark_end731, mark_end758, mark_end785, mark_end806, mark_end827, mark_end848, mark_end869, mark_end890, mark_end911, mark_end932, mark_end953, mark_end974, mark_end995, mark_end1016, mark_end1037, mark_end1058, mark_end1079, mark_end1100, mark_end1121, mark_end1142, mark_end1163, mark_end1184, mark_end1205, mark_end1226, mark_end1247, mark_end1268, mark_end1289, mark_end1310, mark_end1331, mark_end1352, mark_end1373, mark_end1394, mark_end1398, mark_end1419, mark_end1430, mark_end1434, mark_end1455, mark_end1466, mark_end1500, mark_end1534, mark_end1564, mark_end1568, mark_end1572, mark_end1576 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, local_skip, eof, lookahead, v0, local_advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp27, v18, cmp31, v19, cmp35, v20, cmp39, v21, cmp43, v22, cmp45, v23, cmp49, v24, cmp51, v25, cmp54, v26, loadedv58, v27, cmp60, v28, cmp64, v29, cmp68, v30, cmp72, v31, cmp75, v32, cmp78, v33, cmp82, v34, cmp85, v35, cmp88, v36, loadedv92, v37, cmp94, v38, cmp98, v39, cmp101, v40, cmp104, v41, cmp108, v42, loadedv112, v43, cmp114, v44, cmp118, v45, cmp121, v46, cmp124, v47, cmp127, v48, loadedv131, v49, cmp133, v50, cmp137, v51, cmp140, v52, cmp143, v53, cmp147, v54, loadedv151, v55, cmp153, v56, loadedv157, v57, cmp159, v58, loadedv163, v59, cmp165, v60, loadedv169, v61, cmp171, v62, loadedv175, v63, cmp177, v64, cmp181, v65, cmp184, v66, cmp187, v67, loadedv191, v68, cmp193, v69, cmp197, v70, cmp201, v71, cmp205, v72, cmp209, v73, cmp212, v74, cmp215, v75, cmp219, v76, cmp222, v77, cmp225, v78, cmp228, v79, loadedv232, v80, cmp234, v81, loadedv238, v82, cmp240, v83, cmp243, v84, loadedv247, v85, cmp249, v86, cmp252, v87, loadedv256, v88, cmp258, v89, cmp261, v90, loadedv265, v91, cmp267, v92, cmp270, v93, loadedv274, v94, cmp276, v95, cmp279, v96, loadedv283, v97, cmp285, v98, cmp288, v99, cmp292, v100, cmp295, v101, loadedv299, v102, cmp301, v103, cmp304, v104, loadedv308, v105, cmp310, v106, cmp313, v107, cmp316, v108, cmp320, v109, cmp323, v110, cmp326, v111, cmp329, v112, cmp332, v113, loadedv336, v114, cmp338, v115, cmp341, v116, cmp344, v117, cmp348, v118, cmp351, v119, loadedv355, v120, cmp357, v121, cmp360, v122, cmp363, v123, cmp366, v124, cmp369, v125, cmp372, v126, loadedv376, v127, loadedv378, v128, cmp381, v129, cmp385, v130, cmp389, v131, cmp392, v132, cmp395, v133, cmp399, v134, cmp402, v135, cmp405, v136, cmp408, v137, loadedv412, v138, loadedv414, v139, cmp417, v140, cmp421, v141, cmp424, v142, cmp427, v143, cmp431, v144, cmp434, v145, cmp437, v146, cmp440, v147, loadedv444, v148, result_symbol, v149, mark_end, v150, v151, v152, loadedv446, v153, result_symbol448, v154, mark_end449, v155, v156, v157, loadedv450, v158, result_symbol452, v159, mark_end453, v160, v161, v162, cmp454, v163, cmp457, v164, cmp460, v165, cmp464, v166, cmp467, v167, loadedv471, v168, result_symbol473, v169, mark_end474, v170, v171, v172, cmp475, v173, cmp478, v174, loadedv482, v175, result_symbol484, v176, mark_end485, v177, v178, v179, loadedv486, v180, result_symbol488, v181, mark_end489, v182, v183, v184, loadedv490, v185, result_symbol492, v186, mark_end493, v187, v188, v189, cmp494, v190, cmp498, v191, loadedv502, v192, result_symbol504, v193, mark_end505, v194, v195, v196, loadedv506, v197, result_symbol508, v198, mark_end509, v199, v200, v201, loadedv510, v202, result_symbol512, v203, mark_end513, v204, v205, v206, loadedv514, v207, result_symbol516, v208, mark_end517, v209, v210, v211, cmp518, v212, cmp521, v213, cmp524, v214, cmp527, v215, cmp530, v216, cmp533, v217, cmp536, v218, cmp539, v219, cmp542, v220, loadedv546, v221, result_symbol548, v222, mark_end549, v223, v224, v225, cmp550, v226, cmp553, v227, cmp556, v228, cmp559, v229, cmp562, v230, cmp565, v231, cmp568, v232, cmp571, v233, loadedv575, v234, result_symbol577, v235, mark_end578, v236, v237, v238, loadedv579, v239, result_symbol581, v240, mark_end582, v241, v242, v243, cmp583, v244, loadedv587, v245, result_symbol589, v246, mark_end590, v247, v248, v249, cmp591, v250, cmp595, v251, cmp598, v252, loadedv602, v253, result_symbol604, v254, mark_end605, v255, v256, v257, cmp606, v258, cmp610, v259, cmp613, v260, loadedv617, v261, result_symbol619, v262, mark_end620, v263, v264, v265, cmp621, v266, cmp625, v267, cmp628, v268, loadedv632, v269, result_symbol634, v270, mark_end635, v271, v272, v273, cmp636, v274, cmp640, v275, cmp643, v276, loadedv647, v277, result_symbol649, v278, mark_end650, v279, v280, v281, cmp651, v282, cmp655, v283, cmp658, v284, cmp661, v285, cmp664, v286, cmp667, v287, cmp670, v288, loadedv674, v289, result_symbol676, v290, mark_end677, v291, v292, v293, cmp678, v294, cmp682, v295, cmp685, v296, cmp688, v297, cmp691, v298, cmp694, v299, cmp697, v300, loadedv701, v301, result_symbol703, v302, mark_end704, v303, v304, v305, cmp705, v306, cmp709, v307, cmp712, v308, cmp715, v309, cmp718, v310, cmp721, v311, cmp724, v312, loadedv728, v313, result_symbol730, v314, mark_end731, v315, v316, v317, cmp732, v318, cmp736, v319, cmp739, v320, cmp742, v321, cmp745, v322, cmp748, v323, cmp751, v324, loadedv755, v325, result_symbol757, v326, mark_end758, v327, v328, v329, cmp759, v330, cmp763, v331, cmp766, v332, cmp769, v333, cmp772, v334, cmp775, v335, cmp778, v336, loadedv782, v337, result_symbol784, v338, mark_end785, v339, v340, v341, cmp786, v342, cmp790, v343, cmp793, v344, cmp796, v345, cmp799, v346, loadedv803, v347, result_symbol805, v348, mark_end806, v349, v350, v351, cmp807, v352, cmp811, v353, cmp814, v354, cmp817, v355, cmp820, v356, loadedv824, v357, result_symbol826, v358, mark_end827, v359, v360, v361, cmp828, v362, cmp832, v363, cmp835, v364, cmp838, v365, cmp841, v366, loadedv845, v367, result_symbol847, v368, mark_end848, v369, v370, v371, cmp849, v372, cmp853, v373, cmp856, v374, cmp859, v375, cmp862, v376, loadedv866, v377, result_symbol868, v378, mark_end869, v379, v380, v381, cmp870, v382, cmp874, v383, cmp877, v384, cmp880, v385, cmp883, v386, loadedv887, v387, result_symbol889, v388, mark_end890, v389, v390, v391, cmp891, v392, cmp895, v393, cmp898, v394, cmp901, v395, cmp904, v396, loadedv908, v397, result_symbol910, v398, mark_end911, v399, v400, v401, cmp912, v402, cmp916, v403, cmp919, v404, cmp922, v405, cmp925, v406, loadedv929, v407, result_symbol931, v408, mark_end932, v409, v410, v411, cmp933, v412, cmp937, v413, cmp940, v414, cmp943, v415, cmp946, v416, loadedv950, v417, result_symbol952, v418, mark_end953, v419, v420, v421, cmp954, v422, cmp958, v423, cmp961, v424, cmp964, v425, cmp967, v426, loadedv971, v427, result_symbol973, v428, mark_end974, v429, v430, v431, cmp975, v432, cmp979, v433, cmp982, v434, cmp985, v435, cmp988, v436, loadedv992, v437, result_symbol994, v438, mark_end995, v439, v440, v441, cmp996, v442, cmp1000, v443, cmp1003, v444, cmp1006, v445, cmp1009, v446, loadedv1013, v447, result_symbol1015, v448, mark_end1016, v449, v450, v451, cmp1017, v452, cmp1021, v453, cmp1024, v454, cmp1027, v455, cmp1030, v456, loadedv1034, v457, result_symbol1036, v458, mark_end1037, v459, v460, v461, cmp1038, v462, cmp1042, v463, cmp1045, v464, cmp1048, v465, cmp1051, v466, loadedv1055, v467, result_symbol1057, v468, mark_end1058, v469, v470, v471, cmp1059, v472, cmp1063, v473, cmp1066, v474, cmp1069, v475, cmp1072, v476, loadedv1076, v477, result_symbol1078, v478, mark_end1079, v479, v480, v481, cmp1080, v482, cmp1084, v483, cmp1087, v484, cmp1090, v485, cmp1093, v486, loadedv1097, v487, result_symbol1099, v488, mark_end1100, v489, v490, v491, cmp1101, v492, cmp1105, v493, cmp1108, v494, cmp1111, v495, cmp1114, v496, loadedv1118, v497, result_symbol1120, v498, mark_end1121, v499, v500, v501, cmp1122, v502, cmp1126, v503, cmp1129, v504, cmp1132, v505, cmp1135, v506, loadedv1139, v507, result_symbol1141, v508, mark_end1142, v509, v510, v511, cmp1143, v512, cmp1147, v513, cmp1150, v514, cmp1153, v515, cmp1156, v516, loadedv1160, v517, result_symbol1162, v518, mark_end1163, v519, v520, v521, cmp1164, v522, cmp1168, v523, cmp1171, v524, cmp1174, v525, cmp1177, v526, loadedv1181, v527, result_symbol1183, v528, mark_end1184, v529, v530, v531, cmp1185, v532, cmp1189, v533, cmp1192, v534, cmp1195, v535, cmp1198, v536, loadedv1202, v537, result_symbol1204, v538, mark_end1205, v539, v540, v541, cmp1206, v542, cmp1210, v543, cmp1213, v544, cmp1216, v545, cmp1219, v546, loadedv1223, v547, result_symbol1225, v548, mark_end1226, v549, v550, v551, cmp1227, v552, cmp1231, v553, cmp1234, v554, cmp1237, v555, cmp1240, v556, loadedv1244, v557, result_symbol1246, v558, mark_end1247, v559, v560, v561, cmp1248, v562, cmp1252, v563, cmp1255, v564, cmp1258, v565, cmp1261, v566, loadedv1265, v567, result_symbol1267, v568, mark_end1268, v569, v570, v571, cmp1269, v572, cmp1273, v573, cmp1276, v574, cmp1279, v575, cmp1282, v576, loadedv1286, v577, result_symbol1288, v578, mark_end1289, v579, v580, v581, cmp1290, v582, cmp1294, v583, cmp1297, v584, cmp1300, v585, cmp1303, v586, loadedv1307, v587, result_symbol1309, v588, mark_end1310, v589, v590, v591, cmp1311, v592, cmp1315, v593, cmp1318, v594, cmp1321, v595, cmp1324, v596, loadedv1328, v597, result_symbol1330, v598, mark_end1331, v599, v600, v601, cmp1332, v602, cmp1336, v603, cmp1339, v604, cmp1342, v605, cmp1345, v606, loadedv1349, v607, result_symbol1351, v608, mark_end1352, v609, v610, v611, cmp1353, v612, cmp1357, v613, cmp1360, v614, cmp1363, v615, cmp1366, v616, loadedv1370, v617, result_symbol1372, v618, mark_end1373, v619, v620, v621, cmp1374, v622, cmp1378, v623, cmp1381, v624, cmp1384, v625, cmp1387, v626, loadedv1391, v627, result_symbol1393, v628, mark_end1394, v629, v630, v631, loadedv1395, v632, result_symbol1397, v633, mark_end1398, v634, v635, v636, cmp1399, v637, cmp1402, v638, cmp1405, v639, cmp1409, v640, cmp1412, v641, loadedv1416, v642, result_symbol1418, v643, mark_end1419, v644, v645, v646, cmp1420, v647, cmp1423, v648, loadedv1427, v649, result_symbol1429, v650, mark_end1430, v651, v652, v653, loadedv1431, v654, result_symbol1433, v655, mark_end1434, v656, v657, v658, cmp1435, v659, cmp1438, v660, cmp1441, v661, cmp1445, v662, cmp1448, v663, loadedv1452, v664, result_symbol1454, v665, mark_end1455, v666, v667, v668, cmp1456, v669, cmp1459, v670, loadedv1463, v671, result_symbol1465, v672, mark_end1466, v673, v674, v675, cmp1467, v676, cmp1471, v677, cmp1474, v678, cmp1477, v679, cmp1481, v680, cmp1484, v681, cmp1487, v682, cmp1490, v683, cmp1493, v684, loadedv1497, v685, result_symbol1499, v686, mark_end1500, v687, v688, v689, cmp1501, v690, cmp1505, v691, cmp1508, v692, cmp1511, v693, cmp1515, v694, cmp1518, v695, cmp1521, v696, cmp1524, v697, cmp1527, v698, loadedv1531, v699, result_symbol1533, v700, mark_end1534, v701, v702, v703, cmp1535, v704, cmp1538, v705, cmp1541, v706, cmp1545, v707, cmp1548, v708, cmp1551, v709, cmp1554, v710, cmp1557, v711, loadedv1561, v712, result_symbol1563, v713, mark_end1564, v714, v715, v716, loadedv1565, v717, result_symbol1567, v718, mark_end1568, v719, v720, v721, loadedv1569, v722, result_symbol1571, v723, mark_end1572, v724, v725, v726, loadedv1573, v727, result_symbol1575, v728, mark_end1576, v729, v730, v731, loadedv1577, v732

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	state_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int16
		b byte
	}).v)
	result = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	local_skip = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	eof = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	lookahead = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[int16](state_addr) = state
	*libc.As[byte](result) = 0
	*libc.As[byte](local_skip) = 0
	*libc.As[byte](eof) = 0
	goto start

next_state:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	local_advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](local_advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	v3 = *libc.As[byte](local_skip)
	loadedv = (v3 & 1) != 0
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, loadedv)
	goto start

start:
	*libc.As[byte](local_skip) = 0
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v4).F0)
	v5 = *libc.As[int32](lookahead1)
	*libc.As[int32](lookahead) = v5
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	eof2 = libc.Ptr(&libc.As[TSLexer](v6).F6)
	v7 = *libc.As[unsafe.Pointer](eof2)
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	call = libc.FuncFromCode[func(unsafe.Pointer) bool](v7)(v8)
	if call {
		storedv = 1
	} else {
		storedv = 0
	}
	*libc.As[byte](eof) = storedv
	v9 = *libc.As[int16](state_addr)
	conv = int32(uint32(uint16(v9)))
	switch conv {
	case 0:
		goto sw_bb
	case 1:
		goto sw_bb59
	case 2:
		goto sw_bb93
	case 3:
		goto sw_bb113
	case 4:
		goto sw_bb132
	case 5:
		goto sw_bb152
	case 6:
		goto sw_bb158
	case 7:
		goto sw_bb164
	case 8:
		goto sw_bb170
	case 9:
		goto sw_bb176
	case 10:
		goto sw_bb192
	case 11:
		goto sw_bb233
	case 12:
		goto sw_bb239
	case 13:
		goto sw_bb248
	case 14:
		goto sw_bb257
	case 15:
		goto sw_bb266
	case 16:
		goto sw_bb275
	case 17:
		goto sw_bb284
	case 18:
		goto sw_bb300
	case 19:
		goto sw_bb309
	case 20:
		goto sw_bb337
	case 21:
		goto sw_bb356
	case 22:
		goto sw_bb377
	case 23:
		goto sw_bb413
	case 24:
		goto sw_bb445
	case 25:
		goto sw_bb447
	case 26:
		goto sw_bb451
	case 27:
		goto sw_bb472
	case 28:
		goto sw_bb483
	case 29:
		goto sw_bb487
	case 30:
		goto sw_bb491
	case 31:
		goto sw_bb503
	case 32:
		goto sw_bb507
	case 33:
		goto sw_bb511
	case 34:
		goto sw_bb515
	case 35:
		goto sw_bb547
	case 36:
		goto sw_bb576
	case 37:
		goto sw_bb580
	case 38:
		goto sw_bb588
	case 39:
		goto sw_bb603
	case 40:
		goto sw_bb618
	case 41:
		goto sw_bb633
	case 42:
		goto sw_bb648
	case 43:
		goto sw_bb675
	case 44:
		goto sw_bb702
	case 45:
		goto sw_bb729
	case 46:
		goto sw_bb756
	case 47:
		goto sw_bb783
	case 48:
		goto sw_bb804
	case 49:
		goto sw_bb825
	case 50:
		goto sw_bb846
	case 51:
		goto sw_bb867
	case 52:
		goto sw_bb888
	case 53:
		goto sw_bb909
	case 54:
		goto sw_bb930
	case 55:
		goto sw_bb951
	case 56:
		goto sw_bb972
	case 57:
		goto sw_bb993
	case 58:
		goto sw_bb1014
	case 59:
		goto sw_bb1035
	case 60:
		goto sw_bb1056
	case 61:
		goto sw_bb1077
	case 62:
		goto sw_bb1098
	case 63:
		goto sw_bb1119
	case 64:
		goto sw_bb1140
	case 65:
		goto sw_bb1161
	case 66:
		goto sw_bb1182
	case 67:
		goto sw_bb1203
	case 68:
		goto sw_bb1224
	case 69:
		goto sw_bb1245
	case 70:
		goto sw_bb1266
	case 71:
		goto sw_bb1287
	case 72:
		goto sw_bb1308
	case 73:
		goto sw_bb1329
	case 74:
		goto sw_bb1350
	case 75:
		goto sw_bb1371
	case 76:
		goto sw_bb1392
	case 77:
		goto sw_bb1396
	case 78:
		goto sw_bb1417
	case 79:
		goto sw_bb1428
	case 80:
		goto sw_bb1432
	case 81:
		goto sw_bb1453
	case 82:
		goto sw_bb1464
	case 83:
		goto sw_bb1498
	case 84:
		goto sw_bb1532
	case 85:
		goto sw_bb1562
	case 86:
		goto sw_bb1566
	case 87:
		goto sw_bb1570
	case 88:
		goto sw_bb1574
	default:
		goto sw_default
	}

sw_bb:
	v10 = *libc.As[byte](eof)
	loadedv3 = (v10 & 1) != 0
	if loadedv3 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end:
	v11 = *libc.As[int32](lookahead)
	cmp = v11 == 34
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end6:
	v12 = *libc.As[int32](lookahead)
	cmp7 = v12 == 38
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end10:
	v13 = *libc.As[int32](lookahead)
	cmp11 = v13 == 39
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end14:
	v14 = *libc.As[int32](lookahead)
	cmp15 = v14 == 45
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end18:
	v15 = *libc.As[int32](lookahead)
	cmp19 = v15 == 47
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end22:
	v16 = *libc.As[int32](lookahead)
	cmp23 = v16 == 60
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end26:
	v17 = *libc.As[int32](lookahead)
	cmp27 = v17 == 61
	if cmp27 {
		goto if_then29
	} else {
		goto if_end30
	}

if_then29:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end30:
	v18 = *libc.As[int32](lookahead)
	cmp31 = v18 == 62
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end34:
	v19 = *libc.As[int32](lookahead)
	cmp35 = v19 == 123
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end38:
	v20 = *libc.As[int32](lookahead)
	cmp39 = v20 == 125
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*libc.As[int16](state_addr) = 88
	goto next_state

if_end42:
	v21 = *libc.As[int32](lookahead)
	cmp43 = v21 == 68
	if cmp43 {
		goto if_then47
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v22 = *libc.As[int32](lookahead)
	cmp45 = v22 == 100
	if cmp45 {
		goto if_then47
	} else {
		goto if_end48
	}

if_then47:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end48:
	v23 = *libc.As[int32](lookahead)
	cmp49 = 9 <= v23
	if cmp49 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false53
	}

land_lhs_true:
	v24 = *libc.As[int32](lookahead)
	cmp51 = v24 <= 13
	if cmp51 {
		goto if_then56
	} else {
		goto lor_lhs_false53
	}

lor_lhs_false53:
	v25 = *libc.As[int32](lookahead)
	cmp54 = v25 == 32
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end57:
	v26 = *libc.As[byte](result)
	loadedv58 = (v26 & 1) != 0
	*libc.As[bool](retval) = loadedv58
	goto _return

sw_bb59:
	v27 = *libc.As[int32](lookahead)
	cmp60 = v27 == 34
	if cmp60 {
		goto if_then62
	} else {
		goto if_end63
	}

if_then62:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end63:
	v28 = *libc.As[int32](lookahead)
	cmp64 = v28 == 39
	if cmp64 {
		goto if_then66
	} else {
		goto if_end67
	}

if_then66:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end67:
	v29 = *libc.As[int32](lookahead)
	cmp68 = v29 == 123
	if cmp68 {
		goto if_then70
	} else {
		goto if_end71
	}

if_then70:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end71:
	v30 = *libc.As[int32](lookahead)
	cmp72 = 9 <= v30
	if cmp72 {
		goto land_lhs_true74
	} else {
		goto lor_lhs_false77
	}

land_lhs_true74:
	v31 = *libc.As[int32](lookahead)
	cmp75 = v31 <= 13
	if cmp75 {
		goto if_then80
	} else {
		goto lor_lhs_false77
	}

lor_lhs_false77:
	v32 = *libc.As[int32](lookahead)
	cmp78 = v32 == 32
	if cmp78 {
		goto if_then80
	} else {
		goto if_end81
	}

if_then80:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end81:
	v33 = *libc.As[int32](lookahead)
	cmp82 = v33 != 0
	if cmp82 {
		goto land_lhs_true84
	} else {
		goto if_end91
	}

land_lhs_true84:
	v34 = *libc.As[int32](lookahead)
	cmp85 = v34 < 60
	if cmp85 {
		goto if_then90
	} else {
		goto lor_lhs_false87
	}

lor_lhs_false87:
	v35 = *libc.As[int32](lookahead)
	cmp88 = 62 < v35
	if cmp88 {
		goto if_then90
	} else {
		goto if_end91
	}

if_then90:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end91:
	v36 = *libc.As[byte](result)
	loadedv92 = (v36 & 1) != 0
	*libc.As[bool](retval) = loadedv92
	goto _return

sw_bb93:
	v37 = *libc.As[int32](lookahead)
	cmp94 = v37 == 34
	if cmp94 {
		goto if_then96
	} else {
		goto if_end97
	}

if_then96:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end97:
	v38 = *libc.As[int32](lookahead)
	cmp98 = 9 <= v38
	if cmp98 {
		goto land_lhs_true100
	} else {
		goto lor_lhs_false103
	}

land_lhs_true100:
	v39 = *libc.As[int32](lookahead)
	cmp101 = v39 <= 13
	if cmp101 {
		goto if_then106
	} else {
		goto lor_lhs_false103
	}

lor_lhs_false103:
	v40 = *libc.As[int32](lookahead)
	cmp104 = v40 == 32
	if cmp104 {
		goto if_then106
	} else {
		goto if_end107
	}

if_then106:
	*libc.As[int16](state_addr) = 80
	goto next_state

if_end107:
	v41 = *libc.As[int32](lookahead)
	cmp108 = v41 != 0
	if cmp108 {
		goto if_then110
	} else {
		goto if_end111
	}

if_then110:
	*libc.As[int16](state_addr) = 81
	goto next_state

if_end111:
	v42 = *libc.As[byte](result)
	loadedv112 = (v42 & 1) != 0
	*libc.As[bool](retval) = loadedv112
	goto _return

sw_bb113:
	v43 = *libc.As[int32](lookahead)
	cmp114 = v43 == 35
	if cmp114 {
		goto if_then116
	} else {
		goto if_end117
	}

if_then116:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end117:
	v44 = *libc.As[int32](lookahead)
	cmp118 = 65 <= v44
	if cmp118 {
		goto land_lhs_true120
	} else {
		goto lor_lhs_false123
	}

land_lhs_true120:
	v45 = *libc.As[int32](lookahead)
	cmp121 = v45 <= 90
	if cmp121 {
		goto if_then129
	} else {
		goto lor_lhs_false123
	}

lor_lhs_false123:
	v46 = *libc.As[int32](lookahead)
	cmp124 = 97 <= v46
	if cmp124 {
		goto land_lhs_true126
	} else {
		goto if_end130
	}

land_lhs_true126:
	v47 = *libc.As[int32](lookahead)
	cmp127 = v47 <= 122
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end130:
	v48 = *libc.As[byte](result)
	loadedv131 = (v48 & 1) != 0
	*libc.As[bool](retval) = loadedv131
	goto _return

sw_bb132:
	v49 = *libc.As[int32](lookahead)
	cmp133 = v49 == 39
	if cmp133 {
		goto if_then135
	} else {
		goto if_end136
	}

if_then135:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end136:
	v50 = *libc.As[int32](lookahead)
	cmp137 = 9 <= v50
	if cmp137 {
		goto land_lhs_true139
	} else {
		goto lor_lhs_false142
	}

land_lhs_true139:
	v51 = *libc.As[int32](lookahead)
	cmp140 = v51 <= 13
	if cmp140 {
		goto if_then145
	} else {
		goto lor_lhs_false142
	}

lor_lhs_false142:
	v52 = *libc.As[int32](lookahead)
	cmp143 = v52 == 32
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end146:
	v53 = *libc.As[int32](lookahead)
	cmp147 = v53 != 0
	if cmp147 {
		goto if_then149
	} else {
		goto if_end150
	}

if_then149:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end150:
	v54 = *libc.As[byte](result)
	loadedv151 = (v54 & 1) != 0
	*libc.As[bool](retval) = loadedv151
	goto _return

sw_bb152:
	v55 = *libc.As[int32](lookahead)
	cmp153 = v55 == 45
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*libc.As[int16](state_addr) = 85
	goto next_state

if_end156:
	v56 = *libc.As[byte](result)
	loadedv157 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv157
	goto _return

sw_bb158:
	v57 = *libc.As[int32](lookahead)
	cmp159 = v57 == 45
	if cmp159 {
		goto if_then161
	} else {
		goto if_end162
	}

if_then161:
	*libc.As[int16](state_addr) = 86
	goto next_state

if_end162:
	v58 = *libc.As[byte](result)
	loadedv163 = (v58 & 1) != 0
	*libc.As[bool](retval) = loadedv163
	goto _return

sw_bb164:
	v59 = *libc.As[int32](lookahead)
	cmp165 = v59 == 45
	if cmp165 {
		goto if_then167
	} else {
		goto if_end168
	}

if_then167:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end168:
	v60 = *libc.As[byte](result)
	loadedv169 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv169
	goto _return

sw_bb170:
	v61 = *libc.As[int32](lookahead)
	cmp171 = v61 == 45
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end174:
	v62 = *libc.As[byte](result)
	loadedv175 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv175
	goto _return

sw_bb176:
	v63 = *libc.As[int32](lookahead)
	cmp177 = v63 == 45
	if cmp177 {
		goto if_then179
	} else {
		goto if_end180
	}

if_then179:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end180:
	v64 = *libc.As[int32](lookahead)
	cmp181 = 9 <= v64
	if cmp181 {
		goto land_lhs_true183
	} else {
		goto lor_lhs_false186
	}

land_lhs_true183:
	v65 = *libc.As[int32](lookahead)
	cmp184 = v65 <= 13
	if cmp184 {
		goto if_then189
	} else {
		goto lor_lhs_false186
	}

lor_lhs_false186:
	v66 = *libc.As[int32](lookahead)
	cmp187 = v66 == 32
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end190:
	v67 = *libc.As[byte](result)
	loadedv191 = (v67 & 1) != 0
	*libc.As[bool](retval) = loadedv191
	goto _return

sw_bb192:
	v68 = *libc.As[int32](lookahead)
	cmp193 = v68 == 47
	if cmp193 {
		goto if_then195
	} else {
		goto if_end196
	}

if_then195:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end196:
	v69 = *libc.As[int32](lookahead)
	cmp197 = v69 == 61
	if cmp197 {
		goto if_then199
	} else {
		goto if_end200
	}

if_then199:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end200:
	v70 = *libc.As[int32](lookahead)
	cmp201 = v70 == 62
	if cmp201 {
		goto if_then203
	} else {
		goto if_end204
	}

if_then203:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end204:
	v71 = *libc.As[int32](lookahead)
	cmp205 = v71 == 123
	if cmp205 {
		goto if_then207
	} else {
		goto if_end208
	}

if_then207:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end208:
	v72 = *libc.As[int32](lookahead)
	cmp209 = 9 <= v72
	if cmp209 {
		goto land_lhs_true211
	} else {
		goto lor_lhs_false214
	}

land_lhs_true211:
	v73 = *libc.As[int32](lookahead)
	cmp212 = v73 <= 13
	if cmp212 {
		goto if_then217
	} else {
		goto lor_lhs_false214
	}

lor_lhs_false214:
	v74 = *libc.As[int32](lookahead)
	cmp215 = v74 == 32
	if cmp215 {
		goto if_then217
	} else {
		goto if_end218
	}

if_then217:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end218:
	v75 = *libc.As[int32](lookahead)
	cmp219 = v75 != 0
	if cmp219 {
		goto land_lhs_true221
	} else {
		goto if_end231
	}

land_lhs_true221:
	v76 = *libc.As[int32](lookahead)
	cmp222 = v76 != 34
	if cmp222 {
		goto land_lhs_true224
	} else {
		goto if_end231
	}

land_lhs_true224:
	v77 = *libc.As[int32](lookahead)
	cmp225 = v77 != 39
	if cmp225 {
		goto land_lhs_true227
	} else {
		goto if_end231
	}

land_lhs_true227:
	v78 = *libc.As[int32](lookahead)
	cmp228 = v78 != 60
	if cmp228 {
		goto if_then230
	} else {
		goto if_end231
	}

if_then230:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end231:
	v79 = *libc.As[byte](result)
	loadedv232 = (v79 & 1) != 0
	*libc.As[bool](retval) = loadedv232
	goto _return

sw_bb233:
	v80 = *libc.As[int32](lookahead)
	cmp234 = v80 == 62
	if cmp234 {
		goto if_then236
	} else {
		goto if_end237
	}

if_then236:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end237:
	v81 = *libc.As[byte](result)
	loadedv238 = (v81 & 1) != 0
	*libc.As[bool](retval) = loadedv238
	goto _return

sw_bb239:
	v82 = *libc.As[int32](lookahead)
	cmp240 = v82 == 67
	if cmp240 {
		goto if_then245
	} else {
		goto lor_lhs_false242
	}

lor_lhs_false242:
	v83 = *libc.As[int32](lookahead)
	cmp243 = v83 == 99
	if cmp243 {
		goto if_then245
	} else {
		goto if_end246
	}

if_then245:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end246:
	v84 = *libc.As[byte](result)
	loadedv247 = (v84 & 1) != 0
	*libc.As[bool](retval) = loadedv247
	goto _return

sw_bb248:
	v85 = *libc.As[int32](lookahead)
	cmp249 = v85 == 69
	if cmp249 {
		goto if_then254
	} else {
		goto lor_lhs_false251
	}

lor_lhs_false251:
	v86 = *libc.As[int32](lookahead)
	cmp252 = v86 == 101
	if cmp252 {
		goto if_then254
	} else {
		goto if_end255
	}

if_then254:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end255:
	v87 = *libc.As[byte](result)
	loadedv256 = (v87 & 1) != 0
	*libc.As[bool](retval) = loadedv256
	goto _return

sw_bb257:
	v88 = *libc.As[int32](lookahead)
	cmp258 = v88 == 79
	if cmp258 {
		goto if_then263
	} else {
		goto lor_lhs_false260
	}

lor_lhs_false260:
	v89 = *libc.As[int32](lookahead)
	cmp261 = v89 == 111
	if cmp261 {
		goto if_then263
	} else {
		goto if_end264
	}

if_then263:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end264:
	v90 = *libc.As[byte](result)
	loadedv265 = (v90 & 1) != 0
	*libc.As[bool](retval) = loadedv265
	goto _return

sw_bb266:
	v91 = *libc.As[int32](lookahead)
	cmp267 = v91 == 80
	if cmp267 {
		goto if_then272
	} else {
		goto lor_lhs_false269
	}

lor_lhs_false269:
	v92 = *libc.As[int32](lookahead)
	cmp270 = v92 == 112
	if cmp270 {
		goto if_then272
	} else {
		goto if_end273
	}

if_then272:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end273:
	v93 = *libc.As[byte](result)
	loadedv274 = (v93 & 1) != 0
	*libc.As[bool](retval) = loadedv274
	goto _return

sw_bb275:
	v94 = *libc.As[int32](lookahead)
	cmp276 = v94 == 84
	if cmp276 {
		goto if_then281
	} else {
		goto lor_lhs_false278
	}

lor_lhs_false278:
	v95 = *libc.As[int32](lookahead)
	cmp279 = v95 == 116
	if cmp279 {
		goto if_then281
	} else {
		goto if_end282
	}

if_then281:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end282:
	v96 = *libc.As[byte](result)
	loadedv283 = (v96 & 1) != 0
	*libc.As[bool](retval) = loadedv283
	goto _return

sw_bb284:
	v97 = *libc.As[int32](lookahead)
	cmp285 = v97 == 88
	if cmp285 {
		goto if_then290
	} else {
		goto lor_lhs_false287
	}

lor_lhs_false287:
	v98 = *libc.As[int32](lookahead)
	cmp288 = v98 == 120
	if cmp288 {
		goto if_then290
	} else {
		goto if_end291
	}

if_then290:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end291:
	v99 = *libc.As[int32](lookahead)
	cmp292 = 48 <= v99
	if cmp292 {
		goto land_lhs_true294
	} else {
		goto if_end298
	}

land_lhs_true294:
	v100 = *libc.As[int32](lookahead)
	cmp295 = v100 <= 57
	if cmp295 {
		goto if_then297
	} else {
		goto if_end298
	}

if_then297:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end298:
	v101 = *libc.As[byte](result)
	loadedv299 = (v101 & 1) != 0
	*libc.As[bool](retval) = loadedv299
	goto _return

sw_bb300:
	v102 = *libc.As[int32](lookahead)
	cmp301 = v102 == 89
	if cmp301 {
		goto if_then306
	} else {
		goto lor_lhs_false303
	}

lor_lhs_false303:
	v103 = *libc.As[int32](lookahead)
	cmp304 = v103 == 121
	if cmp304 {
		goto if_then306
	} else {
		goto if_end307
	}

if_then306:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end307:
	v104 = *libc.As[byte](result)
	loadedv308 = (v104 & 1) != 0
	*libc.As[bool](retval) = loadedv308
	goto _return

sw_bb309:
	v105 = *libc.As[int32](lookahead)
	cmp310 = 9 <= v105
	if cmp310 {
		goto land_lhs_true312
	} else {
		goto lor_lhs_false315
	}

land_lhs_true312:
	v106 = *libc.As[int32](lookahead)
	cmp313 = v106 <= 13
	if cmp313 {
		goto if_then318
	} else {
		goto lor_lhs_false315
	}

lor_lhs_false315:
	v107 = *libc.As[int32](lookahead)
	cmp316 = v107 == 32
	if cmp316 {
		goto if_then318
	} else {
		goto if_end319
	}

if_then318:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end319:
	v108 = *libc.As[int32](lookahead)
	cmp320 = v108 != 0
	if cmp320 {
		goto land_lhs_true322
	} else {
		goto if_end335
	}

land_lhs_true322:
	v109 = *libc.As[int32](lookahead)
	cmp323 = v109 != 60
	if cmp323 {
		goto land_lhs_true325
	} else {
		goto if_end335
	}

land_lhs_true325:
	v110 = *libc.As[int32](lookahead)
	cmp326 = v110 != 62
	if cmp326 {
		goto land_lhs_true328
	} else {
		goto if_end335
	}

land_lhs_true328:
	v111 = *libc.As[int32](lookahead)
	cmp329 = v111 != 123
	if cmp329 {
		goto land_lhs_true331
	} else {
		goto if_end335
	}

land_lhs_true331:
	v112 = *libc.As[int32](lookahead)
	cmp332 = v112 != 125
	if cmp332 {
		goto if_then334
	} else {
		goto if_end335
	}

if_then334:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end335:
	v113 = *libc.As[byte](result)
	loadedv336 = (v113 & 1) != 0
	*libc.As[bool](retval) = loadedv336
	goto _return

sw_bb337:
	v114 = *libc.As[int32](lookahead)
	cmp338 = 9 <= v114
	if cmp338 {
		goto land_lhs_true340
	} else {
		goto lor_lhs_false343
	}

land_lhs_true340:
	v115 = *libc.As[int32](lookahead)
	cmp341 = v115 <= 13
	if cmp341 {
		goto if_then346
	} else {
		goto lor_lhs_false343
	}

lor_lhs_false343:
	v116 = *libc.As[int32](lookahead)
	cmp344 = v116 == 32
	if cmp344 {
		goto if_then346
	} else {
		goto if_end347
	}

if_then346:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end347:
	v117 = *libc.As[int32](lookahead)
	cmp348 = v117 != 0
	if cmp348 {
		goto land_lhs_true350
	} else {
		goto if_end354
	}

land_lhs_true350:
	v118 = *libc.As[int32](lookahead)
	cmp351 = v118 != 62
	if cmp351 {
		goto if_then353
	} else {
		goto if_end354
	}

if_then353:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end354:
	v119 = *libc.As[byte](result)
	loadedv355 = (v119 & 1) != 0
	*libc.As[bool](retval) = loadedv355
	goto _return

sw_bb356:
	v120 = *libc.As[int32](lookahead)
	cmp357 = 48 <= v120
	if cmp357 {
		goto land_lhs_true359
	} else {
		goto lor_lhs_false362
	}

land_lhs_true359:
	v121 = *libc.As[int32](lookahead)
	cmp360 = v121 <= 57
	if cmp360 {
		goto if_then374
	} else {
		goto lor_lhs_false362
	}

lor_lhs_false362:
	v122 = *libc.As[int32](lookahead)
	cmp363 = 65 <= v122
	if cmp363 {
		goto land_lhs_true365
	} else {
		goto lor_lhs_false368
	}

land_lhs_true365:
	v123 = *libc.As[int32](lookahead)
	cmp366 = v123 <= 70
	if cmp366 {
		goto if_then374
	} else {
		goto lor_lhs_false368
	}

lor_lhs_false368:
	v124 = *libc.As[int32](lookahead)
	cmp369 = 97 <= v124
	if cmp369 {
		goto land_lhs_true371
	} else {
		goto if_end375
	}

land_lhs_true371:
	v125 = *libc.As[int32](lookahead)
	cmp372 = v125 <= 102
	if cmp372 {
		goto if_then374
	} else {
		goto if_end375
	}

if_then374:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end375:
	v126 = *libc.As[byte](result)
	loadedv376 = (v126 & 1) != 0
	*libc.As[bool](retval) = loadedv376
	goto _return

sw_bb377:
	v127 = *libc.As[byte](eof)
	loadedv378 = (v127 & 1) != 0
	if loadedv378 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end380:
	v128 = *libc.As[int32](lookahead)
	cmp381 = v128 == 45
	if cmp381 {
		goto if_then383
	} else {
		goto if_end384
	}

if_then383:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end384:
	v129 = *libc.As[int32](lookahead)
	cmp385 = v129 == 60
	if cmp385 {
		goto if_then387
	} else {
		goto if_end388
	}

if_then387:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end388:
	v130 = *libc.As[int32](lookahead)
	cmp389 = 9 <= v130
	if cmp389 {
		goto land_lhs_true391
	} else {
		goto lor_lhs_false394
	}

land_lhs_true391:
	v131 = *libc.As[int32](lookahead)
	cmp392 = v131 <= 13
	if cmp392 {
		goto if_then397
	} else {
		goto lor_lhs_false394
	}

lor_lhs_false394:
	v132 = *libc.As[int32](lookahead)
	cmp395 = v132 == 32
	if cmp395 {
		goto if_then397
	} else {
		goto if_end398
	}

if_then397:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end398:
	v133 = *libc.As[int32](lookahead)
	cmp399 = v133 != 0
	if cmp399 {
		goto land_lhs_true401
	} else {
		goto if_end411
	}

land_lhs_true401:
	v134 = *libc.As[int32](lookahead)
	cmp402 = v134 != 62
	if cmp402 {
		goto land_lhs_true404
	} else {
		goto if_end411
	}

land_lhs_true404:
	v135 = *libc.As[int32](lookahead)
	cmp405 = v135 != 123
	if cmp405 {
		goto land_lhs_true407
	} else {
		goto if_end411
	}

land_lhs_true407:
	v136 = *libc.As[int32](lookahead)
	cmp408 = v136 != 125
	if cmp408 {
		goto if_then410
	} else {
		goto if_end411
	}

if_then410:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end411:
	v137 = *libc.As[byte](result)
	loadedv412 = (v137 & 1) != 0
	*libc.As[bool](retval) = loadedv412
	goto _return

sw_bb413:
	v138 = *libc.As[byte](eof)
	loadedv414 = (v138 & 1) != 0
	if loadedv414 {
		goto if_then415
	} else {
		goto if_end416
	}

if_then415:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end416:
	v139 = *libc.As[int32](lookahead)
	cmp417 = v139 == 60
	if cmp417 {
		goto if_then419
	} else {
		goto if_end420
	}

if_then419:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end420:
	v140 = *libc.As[int32](lookahead)
	cmp421 = 9 <= v140
	if cmp421 {
		goto land_lhs_true423
	} else {
		goto lor_lhs_false426
	}

land_lhs_true423:
	v141 = *libc.As[int32](lookahead)
	cmp424 = v141 <= 13
	if cmp424 {
		goto if_then429
	} else {
		goto lor_lhs_false426
	}

lor_lhs_false426:
	v142 = *libc.As[int32](lookahead)
	cmp427 = v142 == 32
	if cmp427 {
		goto if_then429
	} else {
		goto if_end430
	}

if_then429:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end430:
	v143 = *libc.As[int32](lookahead)
	cmp431 = v143 != 0
	if cmp431 {
		goto land_lhs_true433
	} else {
		goto if_end443
	}

land_lhs_true433:
	v144 = *libc.As[int32](lookahead)
	cmp434 = v144 != 62
	if cmp434 {
		goto land_lhs_true436
	} else {
		goto if_end443
	}

land_lhs_true436:
	v145 = *libc.As[int32](lookahead)
	cmp437 = v145 != 123
	if cmp437 {
		goto land_lhs_true439
	} else {
		goto if_end443
	}

land_lhs_true439:
	v146 = *libc.As[int32](lookahead)
	cmp440 = v146 != 125
	if cmp440 {
		goto if_then442
	} else {
		goto if_end443
	}

if_then442:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end443:
	v147 = *libc.As[byte](result)
	loadedv444 = (v147 & 1) != 0
	*libc.As[bool](retval) = loadedv444
	goto _return

sw_bb445:
	*libc.As[byte](result) = 1
	v148 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v148).F1)
	*libc.As[int16](result_symbol) = 0
	v149 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v149).F3)
	v150 = *libc.As[unsafe.Pointer](mark_end)
	v151 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v150)(v151)
	v152 = *libc.As[byte](result)
	loadedv446 = (v152 & 1) != 0
	*libc.As[bool](retval) = loadedv446
	goto _return

sw_bb447:
	*libc.As[byte](result) = 1
	v153 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol448 = libc.Ptr(&libc.As[TSLexer](v153).F1)
	*libc.As[int16](result_symbol448) = 1
	v154 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end449 = libc.Ptr(&libc.As[TSLexer](v154).F3)
	v155 = *libc.As[unsafe.Pointer](mark_end449)
	v156 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v155)(v156)
	v157 = *libc.As[byte](result)
	loadedv450 = (v157 & 1) != 0
	*libc.As[bool](retval) = loadedv450
	goto _return

sw_bb451:
	*libc.As[byte](result) = 1
	v158 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol452 = libc.Ptr(&libc.As[TSLexer](v158).F1)
	*libc.As[int16](result_symbol452) = 2
	v159 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end453 = libc.Ptr(&libc.As[TSLexer](v159).F3)
	v160 = *libc.As[unsafe.Pointer](mark_end453)
	v161 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v160)(v161)
	v162 = *libc.As[int32](lookahead)
	cmp454 = 9 <= v162
	if cmp454 {
		goto land_lhs_true456
	} else {
		goto lor_lhs_false459
	}

land_lhs_true456:
	v163 = *libc.As[int32](lookahead)
	cmp457 = v163 <= 13
	if cmp457 {
		goto if_then462
	} else {
		goto lor_lhs_false459
	}

lor_lhs_false459:
	v164 = *libc.As[int32](lookahead)
	cmp460 = v164 == 32
	if cmp460 {
		goto if_then462
	} else {
		goto if_end463
	}

if_then462:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end463:
	v165 = *libc.As[int32](lookahead)
	cmp464 = v165 != 0
	if cmp464 {
		goto land_lhs_true466
	} else {
		goto if_end470
	}

land_lhs_true466:
	v166 = *libc.As[int32](lookahead)
	cmp467 = v166 != 62
	if cmp467 {
		goto if_then469
	} else {
		goto if_end470
	}

if_then469:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end470:
	v167 = *libc.As[byte](result)
	loadedv471 = (v167 & 1) != 0
	*libc.As[bool](retval) = loadedv471
	goto _return

sw_bb472:
	*libc.As[byte](result) = 1
	v168 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol473 = libc.Ptr(&libc.As[TSLexer](v168).F1)
	*libc.As[int16](result_symbol473) = 2
	v169 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end474 = libc.Ptr(&libc.As[TSLexer](v169).F3)
	v170 = *libc.As[unsafe.Pointer](mark_end474)
	v171 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v170)(v171)
	v172 = *libc.As[int32](lookahead)
	cmp475 = v172 != 0
	if cmp475 {
		goto land_lhs_true477
	} else {
		goto if_end481
	}

land_lhs_true477:
	v173 = *libc.As[int32](lookahead)
	cmp478 = v173 != 62
	if cmp478 {
		goto if_then480
	} else {
		goto if_end481
	}

if_then480:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end481:
	v174 = *libc.As[byte](result)
	loadedv482 = (v174 & 1) != 0
	*libc.As[bool](retval) = loadedv482
	goto _return

sw_bb483:
	*libc.As[byte](result) = 1
	v175 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol484 = libc.Ptr(&libc.As[TSLexer](v175).F1)
	*libc.As[int16](result_symbol484) = 3
	v176 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end485 = libc.Ptr(&libc.As[TSLexer](v176).F3)
	v177 = *libc.As[unsafe.Pointer](mark_end485)
	v178 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v177)(v178)
	v179 = *libc.As[byte](result)
	loadedv486 = (v179 & 1) != 0
	*libc.As[bool](retval) = loadedv486
	goto _return

sw_bb487:
	*libc.As[byte](result) = 1
	v180 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol488 = libc.Ptr(&libc.As[TSLexer](v180).F1)
	*libc.As[int16](result_symbol488) = 4
	v181 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end489 = libc.Ptr(&libc.As[TSLexer](v181).F3)
	v182 = *libc.As[unsafe.Pointer](mark_end489)
	v183 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v182)(v183)
	v184 = *libc.As[byte](result)
	loadedv490 = (v184 & 1) != 0
	*libc.As[bool](retval) = loadedv490
	goto _return

sw_bb491:
	*libc.As[byte](result) = 1
	v185 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol492 = libc.Ptr(&libc.As[TSLexer](v185).F1)
	*libc.As[int16](result_symbol492) = 5
	v186 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end493 = libc.Ptr(&libc.As[TSLexer](v186).F3)
	v187 = *libc.As[unsafe.Pointer](mark_end493)
	v188 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v187)(v188)
	v189 = *libc.As[int32](lookahead)
	cmp494 = v189 == 33
	if cmp494 {
		goto if_then496
	} else {
		goto if_end497
	}

if_then496:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end497:
	v190 = *libc.As[int32](lookahead)
	cmp498 = v190 == 47
	if cmp498 {
		goto if_then500
	} else {
		goto if_end501
	}

if_then500:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end501:
	v191 = *libc.As[byte](result)
	loadedv502 = (v191 & 1) != 0
	*libc.As[bool](retval) = loadedv502
	goto _return

sw_bb503:
	*libc.As[byte](result) = 1
	v192 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol504 = libc.Ptr(&libc.As[TSLexer](v192).F1)
	*libc.As[int16](result_symbol504) = 6
	v193 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end505 = libc.Ptr(&libc.As[TSLexer](v193).F3)
	v194 = *libc.As[unsafe.Pointer](mark_end505)
	v195 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v194)(v195)
	v196 = *libc.As[byte](result)
	loadedv506 = (v196 & 1) != 0
	*libc.As[bool](retval) = loadedv506
	goto _return

sw_bb507:
	*libc.As[byte](result) = 1
	v197 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol508 = libc.Ptr(&libc.As[TSLexer](v197).F1)
	*libc.As[int16](result_symbol508) = 7
	v198 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end509 = libc.Ptr(&libc.As[TSLexer](v198).F3)
	v199 = *libc.As[unsafe.Pointer](mark_end509)
	v200 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v199)(v200)
	v201 = *libc.As[byte](result)
	loadedv510 = (v201 & 1) != 0
	*libc.As[bool](retval) = loadedv510
	goto _return

sw_bb511:
	*libc.As[byte](result) = 1
	v202 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol512 = libc.Ptr(&libc.As[TSLexer](v202).F1)
	*libc.As[int16](result_symbol512) = 8
	v203 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end513 = libc.Ptr(&libc.As[TSLexer](v203).F3)
	v204 = *libc.As[unsafe.Pointer](mark_end513)
	v205 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v204)(v205)
	v206 = *libc.As[byte](result)
	loadedv514 = (v206 & 1) != 0
	*libc.As[bool](retval) = loadedv514
	goto _return

sw_bb515:
	*libc.As[byte](result) = 1
	v207 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol516 = libc.Ptr(&libc.As[TSLexer](v207).F1)
	*libc.As[int16](result_symbol516) = 9
	v208 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end517 = libc.Ptr(&libc.As[TSLexer](v208).F3)
	v209 = *libc.As[unsafe.Pointer](mark_end517)
	v210 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v209)(v210)
	v211 = *libc.As[int32](lookahead)
	cmp518 = v211 != 0
	if cmp518 {
		goto land_lhs_true520
	} else {
		goto if_end545
	}

land_lhs_true520:
	v212 = *libc.As[int32](lookahead)
	cmp521 = v212 < 9
	if cmp521 {
		goto land_lhs_true526
	} else {
		goto lor_lhs_false523
	}

lor_lhs_false523:
	v213 = *libc.As[int32](lookahead)
	cmp524 = 13 < v213
	if cmp524 {
		goto land_lhs_true526
	} else {
		goto if_end545
	}

land_lhs_true526:
	v214 = *libc.As[int32](lookahead)
	cmp527 = v214 != 32
	if cmp527 {
		goto land_lhs_true529
	} else {
		goto if_end545
	}

land_lhs_true529:
	v215 = *libc.As[int32](lookahead)
	cmp530 = v215 != 34
	if cmp530 {
		goto land_lhs_true532
	} else {
		goto if_end545
	}

land_lhs_true532:
	v216 = *libc.As[int32](lookahead)
	cmp533 = v216 != 39
	if cmp533 {
		goto land_lhs_true535
	} else {
		goto if_end545
	}

land_lhs_true535:
	v217 = *libc.As[int32](lookahead)
	cmp536 = v217 != 47
	if cmp536 {
		goto land_lhs_true538
	} else {
		goto if_end545
	}

land_lhs_true538:
	v218 = *libc.As[int32](lookahead)
	cmp539 = v218 < 60
	if cmp539 {
		goto if_then544
	} else {
		goto lor_lhs_false541
	}

lor_lhs_false541:
	v219 = *libc.As[int32](lookahead)
	cmp542 = 62 < v219
	if cmp542 {
		goto if_then544
	} else {
		goto if_end545
	}

if_then544:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end545:
	v220 = *libc.As[byte](result)
	loadedv546 = (v220 & 1) != 0
	*libc.As[bool](retval) = loadedv546
	goto _return

sw_bb547:
	*libc.As[byte](result) = 1
	v221 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol548 = libc.Ptr(&libc.As[TSLexer](v221).F1)
	*libc.As[int16](result_symbol548) = 10
	v222 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end549 = libc.Ptr(&libc.As[TSLexer](v222).F3)
	v223 = *libc.As[unsafe.Pointer](mark_end549)
	v224 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v223)(v224)
	v225 = *libc.As[int32](lookahead)
	cmp550 = v225 != 0
	if cmp550 {
		goto land_lhs_true552
	} else {
		goto if_end574
	}

land_lhs_true552:
	v226 = *libc.As[int32](lookahead)
	cmp553 = v226 < 9
	if cmp553 {
		goto land_lhs_true558
	} else {
		goto lor_lhs_false555
	}

lor_lhs_false555:
	v227 = *libc.As[int32](lookahead)
	cmp556 = 13 < v227
	if cmp556 {
		goto land_lhs_true558
	} else {
		goto if_end574
	}

land_lhs_true558:
	v228 = *libc.As[int32](lookahead)
	cmp559 = v228 != 32
	if cmp559 {
		goto land_lhs_true561
	} else {
		goto if_end574
	}

land_lhs_true561:
	v229 = *libc.As[int32](lookahead)
	cmp562 = v229 != 34
	if cmp562 {
		goto land_lhs_true564
	} else {
		goto if_end574
	}

land_lhs_true564:
	v230 = *libc.As[int32](lookahead)
	cmp565 = v230 != 39
	if cmp565 {
		goto land_lhs_true567
	} else {
		goto if_end574
	}

land_lhs_true567:
	v231 = *libc.As[int32](lookahead)
	cmp568 = v231 < 60
	if cmp568 {
		goto if_then573
	} else {
		goto lor_lhs_false570
	}

lor_lhs_false570:
	v232 = *libc.As[int32](lookahead)
	cmp571 = 62 < v232
	if cmp571 {
		goto if_then573
	} else {
		goto if_end574
	}

if_then573:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end574:
	v233 = *libc.As[byte](result)
	loadedv575 = (v233 & 1) != 0
	*libc.As[bool](retval) = loadedv575
	goto _return

sw_bb576:
	*libc.As[byte](result) = 1
	v234 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol577 = libc.Ptr(&libc.As[TSLexer](v234).F1)
	*libc.As[int16](result_symbol577) = 11
	v235 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end578 = libc.Ptr(&libc.As[TSLexer](v235).F3)
	v236 = *libc.As[unsafe.Pointer](mark_end578)
	v237 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v236)(v237)
	v238 = *libc.As[byte](result)
	loadedv579 = (v238 & 1) != 0
	*libc.As[bool](retval) = loadedv579
	goto _return

sw_bb580:
	*libc.As[byte](result) = 1
	v239 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol581 = libc.Ptr(&libc.As[TSLexer](v239).F1)
	*libc.As[int16](result_symbol581) = 11
	v240 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end582 = libc.Ptr(&libc.As[TSLexer](v240).F3)
	v241 = *libc.As[unsafe.Pointer](mark_end582)
	v242 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v241)(v242)
	v243 = *libc.As[int32](lookahead)
	cmp583 = v243 == 59
	if cmp583 {
		goto if_then585
	} else {
		goto if_end586
	}

if_then585:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end586:
	v244 = *libc.As[byte](result)
	loadedv587 = (v244 & 1) != 0
	*libc.As[bool](retval) = loadedv587
	goto _return

sw_bb588:
	*libc.As[byte](result) = 1
	v245 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol589 = libc.Ptr(&libc.As[TSLexer](v245).F1)
	*libc.As[int16](result_symbol589) = 11
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end590 = libc.Ptr(&libc.As[TSLexer](v246).F3)
	v247 = *libc.As[unsafe.Pointer](mark_end590)
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v247)(v248)
	v249 = *libc.As[int32](lookahead)
	cmp591 = v249 == 59
	if cmp591 {
		goto if_then593
	} else {
		goto if_end594
	}

if_then593:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end594:
	v250 = *libc.As[int32](lookahead)
	cmp595 = 48 <= v250
	if cmp595 {
		goto land_lhs_true597
	} else {
		goto if_end601
	}

land_lhs_true597:
	v251 = *libc.As[int32](lookahead)
	cmp598 = v251 <= 57
	if cmp598 {
		goto if_then600
	} else {
		goto if_end601
	}

if_then600:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end601:
	v252 = *libc.As[byte](result)
	loadedv602 = (v252 & 1) != 0
	*libc.As[bool](retval) = loadedv602
	goto _return

sw_bb603:
	*libc.As[byte](result) = 1
	v253 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol604 = libc.Ptr(&libc.As[TSLexer](v253).F1)
	*libc.As[int16](result_symbol604) = 11
	v254 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end605 = libc.Ptr(&libc.As[TSLexer](v254).F3)
	v255 = *libc.As[unsafe.Pointer](mark_end605)
	v256 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v255)(v256)
	v257 = *libc.As[int32](lookahead)
	cmp606 = v257 == 59
	if cmp606 {
		goto if_then608
	} else {
		goto if_end609
	}

if_then608:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end609:
	v258 = *libc.As[int32](lookahead)
	cmp610 = 48 <= v258
	if cmp610 {
		goto land_lhs_true612
	} else {
		goto if_end616
	}

land_lhs_true612:
	v259 = *libc.As[int32](lookahead)
	cmp613 = v259 <= 57
	if cmp613 {
		goto if_then615
	} else {
		goto if_end616
	}

if_then615:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end616:
	v260 = *libc.As[byte](result)
	loadedv617 = (v260 & 1) != 0
	*libc.As[bool](retval) = loadedv617
	goto _return

sw_bb618:
	*libc.As[byte](result) = 1
	v261 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol619 = libc.Ptr(&libc.As[TSLexer](v261).F1)
	*libc.As[int16](result_symbol619) = 11
	v262 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end620 = libc.Ptr(&libc.As[TSLexer](v262).F3)
	v263 = *libc.As[unsafe.Pointer](mark_end620)
	v264 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v263)(v264)
	v265 = *libc.As[int32](lookahead)
	cmp621 = v265 == 59
	if cmp621 {
		goto if_then623
	} else {
		goto if_end624
	}

if_then623:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end624:
	v266 = *libc.As[int32](lookahead)
	cmp625 = 48 <= v266
	if cmp625 {
		goto land_lhs_true627
	} else {
		goto if_end631
	}

land_lhs_true627:
	v267 = *libc.As[int32](lookahead)
	cmp628 = v267 <= 57
	if cmp628 {
		goto if_then630
	} else {
		goto if_end631
	}

if_then630:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end631:
	v268 = *libc.As[byte](result)
	loadedv632 = (v268 & 1) != 0
	*libc.As[bool](retval) = loadedv632
	goto _return

sw_bb633:
	*libc.As[byte](result) = 1
	v269 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol634 = libc.Ptr(&libc.As[TSLexer](v269).F1)
	*libc.As[int16](result_symbol634) = 11
	v270 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end635 = libc.Ptr(&libc.As[TSLexer](v270).F3)
	v271 = *libc.As[unsafe.Pointer](mark_end635)
	v272 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v271)(v272)
	v273 = *libc.As[int32](lookahead)
	cmp636 = v273 == 59
	if cmp636 {
		goto if_then638
	} else {
		goto if_end639
	}

if_then638:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end639:
	v274 = *libc.As[int32](lookahead)
	cmp640 = 48 <= v274
	if cmp640 {
		goto land_lhs_true642
	} else {
		goto if_end646
	}

land_lhs_true642:
	v275 = *libc.As[int32](lookahead)
	cmp643 = v275 <= 57
	if cmp643 {
		goto if_then645
	} else {
		goto if_end646
	}

if_then645:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end646:
	v276 = *libc.As[byte](result)
	loadedv647 = (v276 & 1) != 0
	*libc.As[bool](retval) = loadedv647
	goto _return

sw_bb648:
	*libc.As[byte](result) = 1
	v277 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol649 = libc.Ptr(&libc.As[TSLexer](v277).F1)
	*libc.As[int16](result_symbol649) = 11
	v278 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end650 = libc.Ptr(&libc.As[TSLexer](v278).F3)
	v279 = *libc.As[unsafe.Pointer](mark_end650)
	v280 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v279)(v280)
	v281 = *libc.As[int32](lookahead)
	cmp651 = v281 == 59
	if cmp651 {
		goto if_then653
	} else {
		goto if_end654
	}

if_then653:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end654:
	v282 = *libc.As[int32](lookahead)
	cmp655 = 48 <= v282
	if cmp655 {
		goto land_lhs_true657
	} else {
		goto lor_lhs_false660
	}

land_lhs_true657:
	v283 = *libc.As[int32](lookahead)
	cmp658 = v283 <= 57
	if cmp658 {
		goto if_then672
	} else {
		goto lor_lhs_false660
	}

lor_lhs_false660:
	v284 = *libc.As[int32](lookahead)
	cmp661 = 65 <= v284
	if cmp661 {
		goto land_lhs_true663
	} else {
		goto lor_lhs_false666
	}

land_lhs_true663:
	v285 = *libc.As[int32](lookahead)
	cmp664 = v285 <= 70
	if cmp664 {
		goto if_then672
	} else {
		goto lor_lhs_false666
	}

lor_lhs_false666:
	v286 = *libc.As[int32](lookahead)
	cmp667 = 97 <= v286
	if cmp667 {
		goto land_lhs_true669
	} else {
		goto if_end673
	}

land_lhs_true669:
	v287 = *libc.As[int32](lookahead)
	cmp670 = v287 <= 102
	if cmp670 {
		goto if_then672
	} else {
		goto if_end673
	}

if_then672:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end673:
	v288 = *libc.As[byte](result)
	loadedv674 = (v288 & 1) != 0
	*libc.As[bool](retval) = loadedv674
	goto _return

sw_bb675:
	*libc.As[byte](result) = 1
	v289 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol676 = libc.Ptr(&libc.As[TSLexer](v289).F1)
	*libc.As[int16](result_symbol676) = 11
	v290 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end677 = libc.Ptr(&libc.As[TSLexer](v290).F3)
	v291 = *libc.As[unsafe.Pointer](mark_end677)
	v292 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v291)(v292)
	v293 = *libc.As[int32](lookahead)
	cmp678 = v293 == 59
	if cmp678 {
		goto if_then680
	} else {
		goto if_end681
	}

if_then680:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end681:
	v294 = *libc.As[int32](lookahead)
	cmp682 = 48 <= v294
	if cmp682 {
		goto land_lhs_true684
	} else {
		goto lor_lhs_false687
	}

land_lhs_true684:
	v295 = *libc.As[int32](lookahead)
	cmp685 = v295 <= 57
	if cmp685 {
		goto if_then699
	} else {
		goto lor_lhs_false687
	}

lor_lhs_false687:
	v296 = *libc.As[int32](lookahead)
	cmp688 = 65 <= v296
	if cmp688 {
		goto land_lhs_true690
	} else {
		goto lor_lhs_false693
	}

land_lhs_true690:
	v297 = *libc.As[int32](lookahead)
	cmp691 = v297 <= 70
	if cmp691 {
		goto if_then699
	} else {
		goto lor_lhs_false693
	}

lor_lhs_false693:
	v298 = *libc.As[int32](lookahead)
	cmp694 = 97 <= v298
	if cmp694 {
		goto land_lhs_true696
	} else {
		goto if_end700
	}

land_lhs_true696:
	v299 = *libc.As[int32](lookahead)
	cmp697 = v299 <= 102
	if cmp697 {
		goto if_then699
	} else {
		goto if_end700
	}

if_then699:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end700:
	v300 = *libc.As[byte](result)
	loadedv701 = (v300 & 1) != 0
	*libc.As[bool](retval) = loadedv701
	goto _return

sw_bb702:
	*libc.As[byte](result) = 1
	v301 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol703 = libc.Ptr(&libc.As[TSLexer](v301).F1)
	*libc.As[int16](result_symbol703) = 11
	v302 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end704 = libc.Ptr(&libc.As[TSLexer](v302).F3)
	v303 = *libc.As[unsafe.Pointer](mark_end704)
	v304 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v303)(v304)
	v305 = *libc.As[int32](lookahead)
	cmp705 = v305 == 59
	if cmp705 {
		goto if_then707
	} else {
		goto if_end708
	}

if_then707:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end708:
	v306 = *libc.As[int32](lookahead)
	cmp709 = 48 <= v306
	if cmp709 {
		goto land_lhs_true711
	} else {
		goto lor_lhs_false714
	}

land_lhs_true711:
	v307 = *libc.As[int32](lookahead)
	cmp712 = v307 <= 57
	if cmp712 {
		goto if_then726
	} else {
		goto lor_lhs_false714
	}

lor_lhs_false714:
	v308 = *libc.As[int32](lookahead)
	cmp715 = 65 <= v308
	if cmp715 {
		goto land_lhs_true717
	} else {
		goto lor_lhs_false720
	}

land_lhs_true717:
	v309 = *libc.As[int32](lookahead)
	cmp718 = v309 <= 70
	if cmp718 {
		goto if_then726
	} else {
		goto lor_lhs_false720
	}

lor_lhs_false720:
	v310 = *libc.As[int32](lookahead)
	cmp721 = 97 <= v310
	if cmp721 {
		goto land_lhs_true723
	} else {
		goto if_end727
	}

land_lhs_true723:
	v311 = *libc.As[int32](lookahead)
	cmp724 = v311 <= 102
	if cmp724 {
		goto if_then726
	} else {
		goto if_end727
	}

if_then726:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end727:
	v312 = *libc.As[byte](result)
	loadedv728 = (v312 & 1) != 0
	*libc.As[bool](retval) = loadedv728
	goto _return

sw_bb729:
	*libc.As[byte](result) = 1
	v313 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol730 = libc.Ptr(&libc.As[TSLexer](v313).F1)
	*libc.As[int16](result_symbol730) = 11
	v314 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end731 = libc.Ptr(&libc.As[TSLexer](v314).F3)
	v315 = *libc.As[unsafe.Pointer](mark_end731)
	v316 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v315)(v316)
	v317 = *libc.As[int32](lookahead)
	cmp732 = v317 == 59
	if cmp732 {
		goto if_then734
	} else {
		goto if_end735
	}

if_then734:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end735:
	v318 = *libc.As[int32](lookahead)
	cmp736 = 48 <= v318
	if cmp736 {
		goto land_lhs_true738
	} else {
		goto lor_lhs_false741
	}

land_lhs_true738:
	v319 = *libc.As[int32](lookahead)
	cmp739 = v319 <= 57
	if cmp739 {
		goto if_then753
	} else {
		goto lor_lhs_false741
	}

lor_lhs_false741:
	v320 = *libc.As[int32](lookahead)
	cmp742 = 65 <= v320
	if cmp742 {
		goto land_lhs_true744
	} else {
		goto lor_lhs_false747
	}

land_lhs_true744:
	v321 = *libc.As[int32](lookahead)
	cmp745 = v321 <= 70
	if cmp745 {
		goto if_then753
	} else {
		goto lor_lhs_false747
	}

lor_lhs_false747:
	v322 = *libc.As[int32](lookahead)
	cmp748 = 97 <= v322
	if cmp748 {
		goto land_lhs_true750
	} else {
		goto if_end754
	}

land_lhs_true750:
	v323 = *libc.As[int32](lookahead)
	cmp751 = v323 <= 102
	if cmp751 {
		goto if_then753
	} else {
		goto if_end754
	}

if_then753:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end754:
	v324 = *libc.As[byte](result)
	loadedv755 = (v324 & 1) != 0
	*libc.As[bool](retval) = loadedv755
	goto _return

sw_bb756:
	*libc.As[byte](result) = 1
	v325 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol757 = libc.Ptr(&libc.As[TSLexer](v325).F1)
	*libc.As[int16](result_symbol757) = 11
	v326 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end758 = libc.Ptr(&libc.As[TSLexer](v326).F3)
	v327 = *libc.As[unsafe.Pointer](mark_end758)
	v328 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v327)(v328)
	v329 = *libc.As[int32](lookahead)
	cmp759 = v329 == 59
	if cmp759 {
		goto if_then761
	} else {
		goto if_end762
	}

if_then761:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end762:
	v330 = *libc.As[int32](lookahead)
	cmp763 = 48 <= v330
	if cmp763 {
		goto land_lhs_true765
	} else {
		goto lor_lhs_false768
	}

land_lhs_true765:
	v331 = *libc.As[int32](lookahead)
	cmp766 = v331 <= 57
	if cmp766 {
		goto if_then780
	} else {
		goto lor_lhs_false768
	}

lor_lhs_false768:
	v332 = *libc.As[int32](lookahead)
	cmp769 = 65 <= v332
	if cmp769 {
		goto land_lhs_true771
	} else {
		goto lor_lhs_false774
	}

land_lhs_true771:
	v333 = *libc.As[int32](lookahead)
	cmp772 = v333 <= 70
	if cmp772 {
		goto if_then780
	} else {
		goto lor_lhs_false774
	}

lor_lhs_false774:
	v334 = *libc.As[int32](lookahead)
	cmp775 = 97 <= v334
	if cmp775 {
		goto land_lhs_true777
	} else {
		goto if_end781
	}

land_lhs_true777:
	v335 = *libc.As[int32](lookahead)
	cmp778 = v335 <= 102
	if cmp778 {
		goto if_then780
	} else {
		goto if_end781
	}

if_then780:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end781:
	v336 = *libc.As[byte](result)
	loadedv782 = (v336 & 1) != 0
	*libc.As[bool](retval) = loadedv782
	goto _return

sw_bb783:
	*libc.As[byte](result) = 1
	v337 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol784 = libc.Ptr(&libc.As[TSLexer](v337).F1)
	*libc.As[int16](result_symbol784) = 11
	v338 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end785 = libc.Ptr(&libc.As[TSLexer](v338).F3)
	v339 = *libc.As[unsafe.Pointer](mark_end785)
	v340 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v339)(v340)
	v341 = *libc.As[int32](lookahead)
	cmp786 = v341 == 59
	if cmp786 {
		goto if_then788
	} else {
		goto if_end789
	}

if_then788:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end789:
	v342 = *libc.As[int32](lookahead)
	cmp790 = 65 <= v342
	if cmp790 {
		goto land_lhs_true792
	} else {
		goto lor_lhs_false795
	}

land_lhs_true792:
	v343 = *libc.As[int32](lookahead)
	cmp793 = v343 <= 90
	if cmp793 {
		goto if_then801
	} else {
		goto lor_lhs_false795
	}

lor_lhs_false795:
	v344 = *libc.As[int32](lookahead)
	cmp796 = 97 <= v344
	if cmp796 {
		goto land_lhs_true798
	} else {
		goto if_end802
	}

land_lhs_true798:
	v345 = *libc.As[int32](lookahead)
	cmp799 = v345 <= 122
	if cmp799 {
		goto if_then801
	} else {
		goto if_end802
	}

if_then801:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end802:
	v346 = *libc.As[byte](result)
	loadedv803 = (v346 & 1) != 0
	*libc.As[bool](retval) = loadedv803
	goto _return

sw_bb804:
	*libc.As[byte](result) = 1
	v347 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol805 = libc.Ptr(&libc.As[TSLexer](v347).F1)
	*libc.As[int16](result_symbol805) = 11
	v348 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end806 = libc.Ptr(&libc.As[TSLexer](v348).F3)
	v349 = *libc.As[unsafe.Pointer](mark_end806)
	v350 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v349)(v350)
	v351 = *libc.As[int32](lookahead)
	cmp807 = v351 == 59
	if cmp807 {
		goto if_then809
	} else {
		goto if_end810
	}

if_then809:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end810:
	v352 = *libc.As[int32](lookahead)
	cmp811 = 65 <= v352
	if cmp811 {
		goto land_lhs_true813
	} else {
		goto lor_lhs_false816
	}

land_lhs_true813:
	v353 = *libc.As[int32](lookahead)
	cmp814 = v353 <= 90
	if cmp814 {
		goto if_then822
	} else {
		goto lor_lhs_false816
	}

lor_lhs_false816:
	v354 = *libc.As[int32](lookahead)
	cmp817 = 97 <= v354
	if cmp817 {
		goto land_lhs_true819
	} else {
		goto if_end823
	}

land_lhs_true819:
	v355 = *libc.As[int32](lookahead)
	cmp820 = v355 <= 122
	if cmp820 {
		goto if_then822
	} else {
		goto if_end823
	}

if_then822:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end823:
	v356 = *libc.As[byte](result)
	loadedv824 = (v356 & 1) != 0
	*libc.As[bool](retval) = loadedv824
	goto _return

sw_bb825:
	*libc.As[byte](result) = 1
	v357 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol826 = libc.Ptr(&libc.As[TSLexer](v357).F1)
	*libc.As[int16](result_symbol826) = 11
	v358 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end827 = libc.Ptr(&libc.As[TSLexer](v358).F3)
	v359 = *libc.As[unsafe.Pointer](mark_end827)
	v360 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v359)(v360)
	v361 = *libc.As[int32](lookahead)
	cmp828 = v361 == 59
	if cmp828 {
		goto if_then830
	} else {
		goto if_end831
	}

if_then830:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end831:
	v362 = *libc.As[int32](lookahead)
	cmp832 = 65 <= v362
	if cmp832 {
		goto land_lhs_true834
	} else {
		goto lor_lhs_false837
	}

land_lhs_true834:
	v363 = *libc.As[int32](lookahead)
	cmp835 = v363 <= 90
	if cmp835 {
		goto if_then843
	} else {
		goto lor_lhs_false837
	}

lor_lhs_false837:
	v364 = *libc.As[int32](lookahead)
	cmp838 = 97 <= v364
	if cmp838 {
		goto land_lhs_true840
	} else {
		goto if_end844
	}

land_lhs_true840:
	v365 = *libc.As[int32](lookahead)
	cmp841 = v365 <= 122
	if cmp841 {
		goto if_then843
	} else {
		goto if_end844
	}

if_then843:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end844:
	v366 = *libc.As[byte](result)
	loadedv845 = (v366 & 1) != 0
	*libc.As[bool](retval) = loadedv845
	goto _return

sw_bb846:
	*libc.As[byte](result) = 1
	v367 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol847 = libc.Ptr(&libc.As[TSLexer](v367).F1)
	*libc.As[int16](result_symbol847) = 11
	v368 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end848 = libc.Ptr(&libc.As[TSLexer](v368).F3)
	v369 = *libc.As[unsafe.Pointer](mark_end848)
	v370 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v369)(v370)
	v371 = *libc.As[int32](lookahead)
	cmp849 = v371 == 59
	if cmp849 {
		goto if_then851
	} else {
		goto if_end852
	}

if_then851:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end852:
	v372 = *libc.As[int32](lookahead)
	cmp853 = 65 <= v372
	if cmp853 {
		goto land_lhs_true855
	} else {
		goto lor_lhs_false858
	}

land_lhs_true855:
	v373 = *libc.As[int32](lookahead)
	cmp856 = v373 <= 90
	if cmp856 {
		goto if_then864
	} else {
		goto lor_lhs_false858
	}

lor_lhs_false858:
	v374 = *libc.As[int32](lookahead)
	cmp859 = 97 <= v374
	if cmp859 {
		goto land_lhs_true861
	} else {
		goto if_end865
	}

land_lhs_true861:
	v375 = *libc.As[int32](lookahead)
	cmp862 = v375 <= 122
	if cmp862 {
		goto if_then864
	} else {
		goto if_end865
	}

if_then864:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end865:
	v376 = *libc.As[byte](result)
	loadedv866 = (v376 & 1) != 0
	*libc.As[bool](retval) = loadedv866
	goto _return

sw_bb867:
	*libc.As[byte](result) = 1
	v377 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol868 = libc.Ptr(&libc.As[TSLexer](v377).F1)
	*libc.As[int16](result_symbol868) = 11
	v378 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end869 = libc.Ptr(&libc.As[TSLexer](v378).F3)
	v379 = *libc.As[unsafe.Pointer](mark_end869)
	v380 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v379)(v380)
	v381 = *libc.As[int32](lookahead)
	cmp870 = v381 == 59
	if cmp870 {
		goto if_then872
	} else {
		goto if_end873
	}

if_then872:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end873:
	v382 = *libc.As[int32](lookahead)
	cmp874 = 65 <= v382
	if cmp874 {
		goto land_lhs_true876
	} else {
		goto lor_lhs_false879
	}

land_lhs_true876:
	v383 = *libc.As[int32](lookahead)
	cmp877 = v383 <= 90
	if cmp877 {
		goto if_then885
	} else {
		goto lor_lhs_false879
	}

lor_lhs_false879:
	v384 = *libc.As[int32](lookahead)
	cmp880 = 97 <= v384
	if cmp880 {
		goto land_lhs_true882
	} else {
		goto if_end886
	}

land_lhs_true882:
	v385 = *libc.As[int32](lookahead)
	cmp883 = v385 <= 122
	if cmp883 {
		goto if_then885
	} else {
		goto if_end886
	}

if_then885:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end886:
	v386 = *libc.As[byte](result)
	loadedv887 = (v386 & 1) != 0
	*libc.As[bool](retval) = loadedv887
	goto _return

sw_bb888:
	*libc.As[byte](result) = 1
	v387 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol889 = libc.Ptr(&libc.As[TSLexer](v387).F1)
	*libc.As[int16](result_symbol889) = 11
	v388 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end890 = libc.Ptr(&libc.As[TSLexer](v388).F3)
	v389 = *libc.As[unsafe.Pointer](mark_end890)
	v390 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v389)(v390)
	v391 = *libc.As[int32](lookahead)
	cmp891 = v391 == 59
	if cmp891 {
		goto if_then893
	} else {
		goto if_end894
	}

if_then893:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end894:
	v392 = *libc.As[int32](lookahead)
	cmp895 = 65 <= v392
	if cmp895 {
		goto land_lhs_true897
	} else {
		goto lor_lhs_false900
	}

land_lhs_true897:
	v393 = *libc.As[int32](lookahead)
	cmp898 = v393 <= 90
	if cmp898 {
		goto if_then906
	} else {
		goto lor_lhs_false900
	}

lor_lhs_false900:
	v394 = *libc.As[int32](lookahead)
	cmp901 = 97 <= v394
	if cmp901 {
		goto land_lhs_true903
	} else {
		goto if_end907
	}

land_lhs_true903:
	v395 = *libc.As[int32](lookahead)
	cmp904 = v395 <= 122
	if cmp904 {
		goto if_then906
	} else {
		goto if_end907
	}

if_then906:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end907:
	v396 = *libc.As[byte](result)
	loadedv908 = (v396 & 1) != 0
	*libc.As[bool](retval) = loadedv908
	goto _return

sw_bb909:
	*libc.As[byte](result) = 1
	v397 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol910 = libc.Ptr(&libc.As[TSLexer](v397).F1)
	*libc.As[int16](result_symbol910) = 11
	v398 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end911 = libc.Ptr(&libc.As[TSLexer](v398).F3)
	v399 = *libc.As[unsafe.Pointer](mark_end911)
	v400 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v399)(v400)
	v401 = *libc.As[int32](lookahead)
	cmp912 = v401 == 59
	if cmp912 {
		goto if_then914
	} else {
		goto if_end915
	}

if_then914:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end915:
	v402 = *libc.As[int32](lookahead)
	cmp916 = 65 <= v402
	if cmp916 {
		goto land_lhs_true918
	} else {
		goto lor_lhs_false921
	}

land_lhs_true918:
	v403 = *libc.As[int32](lookahead)
	cmp919 = v403 <= 90
	if cmp919 {
		goto if_then927
	} else {
		goto lor_lhs_false921
	}

lor_lhs_false921:
	v404 = *libc.As[int32](lookahead)
	cmp922 = 97 <= v404
	if cmp922 {
		goto land_lhs_true924
	} else {
		goto if_end928
	}

land_lhs_true924:
	v405 = *libc.As[int32](lookahead)
	cmp925 = v405 <= 122
	if cmp925 {
		goto if_then927
	} else {
		goto if_end928
	}

if_then927:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end928:
	v406 = *libc.As[byte](result)
	loadedv929 = (v406 & 1) != 0
	*libc.As[bool](retval) = loadedv929
	goto _return

sw_bb930:
	*libc.As[byte](result) = 1
	v407 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol931 = libc.Ptr(&libc.As[TSLexer](v407).F1)
	*libc.As[int16](result_symbol931) = 11
	v408 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end932 = libc.Ptr(&libc.As[TSLexer](v408).F3)
	v409 = *libc.As[unsafe.Pointer](mark_end932)
	v410 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v409)(v410)
	v411 = *libc.As[int32](lookahead)
	cmp933 = v411 == 59
	if cmp933 {
		goto if_then935
	} else {
		goto if_end936
	}

if_then935:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end936:
	v412 = *libc.As[int32](lookahead)
	cmp937 = 65 <= v412
	if cmp937 {
		goto land_lhs_true939
	} else {
		goto lor_lhs_false942
	}

land_lhs_true939:
	v413 = *libc.As[int32](lookahead)
	cmp940 = v413 <= 90
	if cmp940 {
		goto if_then948
	} else {
		goto lor_lhs_false942
	}

lor_lhs_false942:
	v414 = *libc.As[int32](lookahead)
	cmp943 = 97 <= v414
	if cmp943 {
		goto land_lhs_true945
	} else {
		goto if_end949
	}

land_lhs_true945:
	v415 = *libc.As[int32](lookahead)
	cmp946 = v415 <= 122
	if cmp946 {
		goto if_then948
	} else {
		goto if_end949
	}

if_then948:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end949:
	v416 = *libc.As[byte](result)
	loadedv950 = (v416 & 1) != 0
	*libc.As[bool](retval) = loadedv950
	goto _return

sw_bb951:
	*libc.As[byte](result) = 1
	v417 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol952 = libc.Ptr(&libc.As[TSLexer](v417).F1)
	*libc.As[int16](result_symbol952) = 11
	v418 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end953 = libc.Ptr(&libc.As[TSLexer](v418).F3)
	v419 = *libc.As[unsafe.Pointer](mark_end953)
	v420 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v419)(v420)
	v421 = *libc.As[int32](lookahead)
	cmp954 = v421 == 59
	if cmp954 {
		goto if_then956
	} else {
		goto if_end957
	}

if_then956:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end957:
	v422 = *libc.As[int32](lookahead)
	cmp958 = 65 <= v422
	if cmp958 {
		goto land_lhs_true960
	} else {
		goto lor_lhs_false963
	}

land_lhs_true960:
	v423 = *libc.As[int32](lookahead)
	cmp961 = v423 <= 90
	if cmp961 {
		goto if_then969
	} else {
		goto lor_lhs_false963
	}

lor_lhs_false963:
	v424 = *libc.As[int32](lookahead)
	cmp964 = 97 <= v424
	if cmp964 {
		goto land_lhs_true966
	} else {
		goto if_end970
	}

land_lhs_true966:
	v425 = *libc.As[int32](lookahead)
	cmp967 = v425 <= 122
	if cmp967 {
		goto if_then969
	} else {
		goto if_end970
	}

if_then969:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end970:
	v426 = *libc.As[byte](result)
	loadedv971 = (v426 & 1) != 0
	*libc.As[bool](retval) = loadedv971
	goto _return

sw_bb972:
	*libc.As[byte](result) = 1
	v427 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol973 = libc.Ptr(&libc.As[TSLexer](v427).F1)
	*libc.As[int16](result_symbol973) = 11
	v428 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end974 = libc.Ptr(&libc.As[TSLexer](v428).F3)
	v429 = *libc.As[unsafe.Pointer](mark_end974)
	v430 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v429)(v430)
	v431 = *libc.As[int32](lookahead)
	cmp975 = v431 == 59
	if cmp975 {
		goto if_then977
	} else {
		goto if_end978
	}

if_then977:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end978:
	v432 = *libc.As[int32](lookahead)
	cmp979 = 65 <= v432
	if cmp979 {
		goto land_lhs_true981
	} else {
		goto lor_lhs_false984
	}

land_lhs_true981:
	v433 = *libc.As[int32](lookahead)
	cmp982 = v433 <= 90
	if cmp982 {
		goto if_then990
	} else {
		goto lor_lhs_false984
	}

lor_lhs_false984:
	v434 = *libc.As[int32](lookahead)
	cmp985 = 97 <= v434
	if cmp985 {
		goto land_lhs_true987
	} else {
		goto if_end991
	}

land_lhs_true987:
	v435 = *libc.As[int32](lookahead)
	cmp988 = v435 <= 122
	if cmp988 {
		goto if_then990
	} else {
		goto if_end991
	}

if_then990:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end991:
	v436 = *libc.As[byte](result)
	loadedv992 = (v436 & 1) != 0
	*libc.As[bool](retval) = loadedv992
	goto _return

sw_bb993:
	*libc.As[byte](result) = 1
	v437 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol994 = libc.Ptr(&libc.As[TSLexer](v437).F1)
	*libc.As[int16](result_symbol994) = 11
	v438 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end995 = libc.Ptr(&libc.As[TSLexer](v438).F3)
	v439 = *libc.As[unsafe.Pointer](mark_end995)
	v440 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v439)(v440)
	v441 = *libc.As[int32](lookahead)
	cmp996 = v441 == 59
	if cmp996 {
		goto if_then998
	} else {
		goto if_end999
	}

if_then998:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end999:
	v442 = *libc.As[int32](lookahead)
	cmp1000 = 65 <= v442
	if cmp1000 {
		goto land_lhs_true1002
	} else {
		goto lor_lhs_false1005
	}

land_lhs_true1002:
	v443 = *libc.As[int32](lookahead)
	cmp1003 = v443 <= 90
	if cmp1003 {
		goto if_then1011
	} else {
		goto lor_lhs_false1005
	}

lor_lhs_false1005:
	v444 = *libc.As[int32](lookahead)
	cmp1006 = 97 <= v444
	if cmp1006 {
		goto land_lhs_true1008
	} else {
		goto if_end1012
	}

land_lhs_true1008:
	v445 = *libc.As[int32](lookahead)
	cmp1009 = v445 <= 122
	if cmp1009 {
		goto if_then1011
	} else {
		goto if_end1012
	}

if_then1011:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end1012:
	v446 = *libc.As[byte](result)
	loadedv1013 = (v446 & 1) != 0
	*libc.As[bool](retval) = loadedv1013
	goto _return

sw_bb1014:
	*libc.As[byte](result) = 1
	v447 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1015 = libc.Ptr(&libc.As[TSLexer](v447).F1)
	*libc.As[int16](result_symbol1015) = 11
	v448 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1016 = libc.Ptr(&libc.As[TSLexer](v448).F3)
	v449 = *libc.As[unsafe.Pointer](mark_end1016)
	v450 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v449)(v450)
	v451 = *libc.As[int32](lookahead)
	cmp1017 = v451 == 59
	if cmp1017 {
		goto if_then1019
	} else {
		goto if_end1020
	}

if_then1019:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1020:
	v452 = *libc.As[int32](lookahead)
	cmp1021 = 65 <= v452
	if cmp1021 {
		goto land_lhs_true1023
	} else {
		goto lor_lhs_false1026
	}

land_lhs_true1023:
	v453 = *libc.As[int32](lookahead)
	cmp1024 = v453 <= 90
	if cmp1024 {
		goto if_then1032
	} else {
		goto lor_lhs_false1026
	}

lor_lhs_false1026:
	v454 = *libc.As[int32](lookahead)
	cmp1027 = 97 <= v454
	if cmp1027 {
		goto land_lhs_true1029
	} else {
		goto if_end1033
	}

land_lhs_true1029:
	v455 = *libc.As[int32](lookahead)
	cmp1030 = v455 <= 122
	if cmp1030 {
		goto if_then1032
	} else {
		goto if_end1033
	}

if_then1032:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end1033:
	v456 = *libc.As[byte](result)
	loadedv1034 = (v456 & 1) != 0
	*libc.As[bool](retval) = loadedv1034
	goto _return

sw_bb1035:
	*libc.As[byte](result) = 1
	v457 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1036 = libc.Ptr(&libc.As[TSLexer](v457).F1)
	*libc.As[int16](result_symbol1036) = 11
	v458 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1037 = libc.Ptr(&libc.As[TSLexer](v458).F3)
	v459 = *libc.As[unsafe.Pointer](mark_end1037)
	v460 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v459)(v460)
	v461 = *libc.As[int32](lookahead)
	cmp1038 = v461 == 59
	if cmp1038 {
		goto if_then1040
	} else {
		goto if_end1041
	}

if_then1040:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1041:
	v462 = *libc.As[int32](lookahead)
	cmp1042 = 65 <= v462
	if cmp1042 {
		goto land_lhs_true1044
	} else {
		goto lor_lhs_false1047
	}

land_lhs_true1044:
	v463 = *libc.As[int32](lookahead)
	cmp1045 = v463 <= 90
	if cmp1045 {
		goto if_then1053
	} else {
		goto lor_lhs_false1047
	}

lor_lhs_false1047:
	v464 = *libc.As[int32](lookahead)
	cmp1048 = 97 <= v464
	if cmp1048 {
		goto land_lhs_true1050
	} else {
		goto if_end1054
	}

land_lhs_true1050:
	v465 = *libc.As[int32](lookahead)
	cmp1051 = v465 <= 122
	if cmp1051 {
		goto if_then1053
	} else {
		goto if_end1054
	}

if_then1053:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end1054:
	v466 = *libc.As[byte](result)
	loadedv1055 = (v466 & 1) != 0
	*libc.As[bool](retval) = loadedv1055
	goto _return

sw_bb1056:
	*libc.As[byte](result) = 1
	v467 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1057 = libc.Ptr(&libc.As[TSLexer](v467).F1)
	*libc.As[int16](result_symbol1057) = 11
	v468 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1058 = libc.Ptr(&libc.As[TSLexer](v468).F3)
	v469 = *libc.As[unsafe.Pointer](mark_end1058)
	v470 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v469)(v470)
	v471 = *libc.As[int32](lookahead)
	cmp1059 = v471 == 59
	if cmp1059 {
		goto if_then1061
	} else {
		goto if_end1062
	}

if_then1061:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1062:
	v472 = *libc.As[int32](lookahead)
	cmp1063 = 65 <= v472
	if cmp1063 {
		goto land_lhs_true1065
	} else {
		goto lor_lhs_false1068
	}

land_lhs_true1065:
	v473 = *libc.As[int32](lookahead)
	cmp1066 = v473 <= 90
	if cmp1066 {
		goto if_then1074
	} else {
		goto lor_lhs_false1068
	}

lor_lhs_false1068:
	v474 = *libc.As[int32](lookahead)
	cmp1069 = 97 <= v474
	if cmp1069 {
		goto land_lhs_true1071
	} else {
		goto if_end1075
	}

land_lhs_true1071:
	v475 = *libc.As[int32](lookahead)
	cmp1072 = v475 <= 122
	if cmp1072 {
		goto if_then1074
	} else {
		goto if_end1075
	}

if_then1074:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end1075:
	v476 = *libc.As[byte](result)
	loadedv1076 = (v476 & 1) != 0
	*libc.As[bool](retval) = loadedv1076
	goto _return

sw_bb1077:
	*libc.As[byte](result) = 1
	v477 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1078 = libc.Ptr(&libc.As[TSLexer](v477).F1)
	*libc.As[int16](result_symbol1078) = 11
	v478 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1079 = libc.Ptr(&libc.As[TSLexer](v478).F3)
	v479 = *libc.As[unsafe.Pointer](mark_end1079)
	v480 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v479)(v480)
	v481 = *libc.As[int32](lookahead)
	cmp1080 = v481 == 59
	if cmp1080 {
		goto if_then1082
	} else {
		goto if_end1083
	}

if_then1082:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1083:
	v482 = *libc.As[int32](lookahead)
	cmp1084 = 65 <= v482
	if cmp1084 {
		goto land_lhs_true1086
	} else {
		goto lor_lhs_false1089
	}

land_lhs_true1086:
	v483 = *libc.As[int32](lookahead)
	cmp1087 = v483 <= 90
	if cmp1087 {
		goto if_then1095
	} else {
		goto lor_lhs_false1089
	}

lor_lhs_false1089:
	v484 = *libc.As[int32](lookahead)
	cmp1090 = 97 <= v484
	if cmp1090 {
		goto land_lhs_true1092
	} else {
		goto if_end1096
	}

land_lhs_true1092:
	v485 = *libc.As[int32](lookahead)
	cmp1093 = v485 <= 122
	if cmp1093 {
		goto if_then1095
	} else {
		goto if_end1096
	}

if_then1095:
	*libc.As[int16](state_addr) = 60
	goto next_state

if_end1096:
	v486 = *libc.As[byte](result)
	loadedv1097 = (v486 & 1) != 0
	*libc.As[bool](retval) = loadedv1097
	goto _return

sw_bb1098:
	*libc.As[byte](result) = 1
	v487 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1099 = libc.Ptr(&libc.As[TSLexer](v487).F1)
	*libc.As[int16](result_symbol1099) = 11
	v488 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1100 = libc.Ptr(&libc.As[TSLexer](v488).F3)
	v489 = *libc.As[unsafe.Pointer](mark_end1100)
	v490 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v489)(v490)
	v491 = *libc.As[int32](lookahead)
	cmp1101 = v491 == 59
	if cmp1101 {
		goto if_then1103
	} else {
		goto if_end1104
	}

if_then1103:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1104:
	v492 = *libc.As[int32](lookahead)
	cmp1105 = 65 <= v492
	if cmp1105 {
		goto land_lhs_true1107
	} else {
		goto lor_lhs_false1110
	}

land_lhs_true1107:
	v493 = *libc.As[int32](lookahead)
	cmp1108 = v493 <= 90
	if cmp1108 {
		goto if_then1116
	} else {
		goto lor_lhs_false1110
	}

lor_lhs_false1110:
	v494 = *libc.As[int32](lookahead)
	cmp1111 = 97 <= v494
	if cmp1111 {
		goto land_lhs_true1113
	} else {
		goto if_end1117
	}

land_lhs_true1113:
	v495 = *libc.As[int32](lookahead)
	cmp1114 = v495 <= 122
	if cmp1114 {
		goto if_then1116
	} else {
		goto if_end1117
	}

if_then1116:
	*libc.As[int16](state_addr) = 61
	goto next_state

if_end1117:
	v496 = *libc.As[byte](result)
	loadedv1118 = (v496 & 1) != 0
	*libc.As[bool](retval) = loadedv1118
	goto _return

sw_bb1119:
	*libc.As[byte](result) = 1
	v497 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1120 = libc.Ptr(&libc.As[TSLexer](v497).F1)
	*libc.As[int16](result_symbol1120) = 11
	v498 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1121 = libc.Ptr(&libc.As[TSLexer](v498).F3)
	v499 = *libc.As[unsafe.Pointer](mark_end1121)
	v500 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v499)(v500)
	v501 = *libc.As[int32](lookahead)
	cmp1122 = v501 == 59
	if cmp1122 {
		goto if_then1124
	} else {
		goto if_end1125
	}

if_then1124:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1125:
	v502 = *libc.As[int32](lookahead)
	cmp1126 = 65 <= v502
	if cmp1126 {
		goto land_lhs_true1128
	} else {
		goto lor_lhs_false1131
	}

land_lhs_true1128:
	v503 = *libc.As[int32](lookahead)
	cmp1129 = v503 <= 90
	if cmp1129 {
		goto if_then1137
	} else {
		goto lor_lhs_false1131
	}

lor_lhs_false1131:
	v504 = *libc.As[int32](lookahead)
	cmp1132 = 97 <= v504
	if cmp1132 {
		goto land_lhs_true1134
	} else {
		goto if_end1138
	}

land_lhs_true1134:
	v505 = *libc.As[int32](lookahead)
	cmp1135 = v505 <= 122
	if cmp1135 {
		goto if_then1137
	} else {
		goto if_end1138
	}

if_then1137:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end1138:
	v506 = *libc.As[byte](result)
	loadedv1139 = (v506 & 1) != 0
	*libc.As[bool](retval) = loadedv1139
	goto _return

sw_bb1140:
	*libc.As[byte](result) = 1
	v507 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1141 = libc.Ptr(&libc.As[TSLexer](v507).F1)
	*libc.As[int16](result_symbol1141) = 11
	v508 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1142 = libc.Ptr(&libc.As[TSLexer](v508).F3)
	v509 = *libc.As[unsafe.Pointer](mark_end1142)
	v510 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v509)(v510)
	v511 = *libc.As[int32](lookahead)
	cmp1143 = v511 == 59
	if cmp1143 {
		goto if_then1145
	} else {
		goto if_end1146
	}

if_then1145:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1146:
	v512 = *libc.As[int32](lookahead)
	cmp1147 = 65 <= v512
	if cmp1147 {
		goto land_lhs_true1149
	} else {
		goto lor_lhs_false1152
	}

land_lhs_true1149:
	v513 = *libc.As[int32](lookahead)
	cmp1150 = v513 <= 90
	if cmp1150 {
		goto if_then1158
	} else {
		goto lor_lhs_false1152
	}

lor_lhs_false1152:
	v514 = *libc.As[int32](lookahead)
	cmp1153 = 97 <= v514
	if cmp1153 {
		goto land_lhs_true1155
	} else {
		goto if_end1159
	}

land_lhs_true1155:
	v515 = *libc.As[int32](lookahead)
	cmp1156 = v515 <= 122
	if cmp1156 {
		goto if_then1158
	} else {
		goto if_end1159
	}

if_then1158:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end1159:
	v516 = *libc.As[byte](result)
	loadedv1160 = (v516 & 1) != 0
	*libc.As[bool](retval) = loadedv1160
	goto _return

sw_bb1161:
	*libc.As[byte](result) = 1
	v517 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1162 = libc.Ptr(&libc.As[TSLexer](v517).F1)
	*libc.As[int16](result_symbol1162) = 11
	v518 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1163 = libc.Ptr(&libc.As[TSLexer](v518).F3)
	v519 = *libc.As[unsafe.Pointer](mark_end1163)
	v520 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v519)(v520)
	v521 = *libc.As[int32](lookahead)
	cmp1164 = v521 == 59
	if cmp1164 {
		goto if_then1166
	} else {
		goto if_end1167
	}

if_then1166:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1167:
	v522 = *libc.As[int32](lookahead)
	cmp1168 = 65 <= v522
	if cmp1168 {
		goto land_lhs_true1170
	} else {
		goto lor_lhs_false1173
	}

land_lhs_true1170:
	v523 = *libc.As[int32](lookahead)
	cmp1171 = v523 <= 90
	if cmp1171 {
		goto if_then1179
	} else {
		goto lor_lhs_false1173
	}

lor_lhs_false1173:
	v524 = *libc.As[int32](lookahead)
	cmp1174 = 97 <= v524
	if cmp1174 {
		goto land_lhs_true1176
	} else {
		goto if_end1180
	}

land_lhs_true1176:
	v525 = *libc.As[int32](lookahead)
	cmp1177 = v525 <= 122
	if cmp1177 {
		goto if_then1179
	} else {
		goto if_end1180
	}

if_then1179:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end1180:
	v526 = *libc.As[byte](result)
	loadedv1181 = (v526 & 1) != 0
	*libc.As[bool](retval) = loadedv1181
	goto _return

sw_bb1182:
	*libc.As[byte](result) = 1
	v527 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1183 = libc.Ptr(&libc.As[TSLexer](v527).F1)
	*libc.As[int16](result_symbol1183) = 11
	v528 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1184 = libc.Ptr(&libc.As[TSLexer](v528).F3)
	v529 = *libc.As[unsafe.Pointer](mark_end1184)
	v530 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v529)(v530)
	v531 = *libc.As[int32](lookahead)
	cmp1185 = v531 == 59
	if cmp1185 {
		goto if_then1187
	} else {
		goto if_end1188
	}

if_then1187:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1188:
	v532 = *libc.As[int32](lookahead)
	cmp1189 = 65 <= v532
	if cmp1189 {
		goto land_lhs_true1191
	} else {
		goto lor_lhs_false1194
	}

land_lhs_true1191:
	v533 = *libc.As[int32](lookahead)
	cmp1192 = v533 <= 90
	if cmp1192 {
		goto if_then1200
	} else {
		goto lor_lhs_false1194
	}

lor_lhs_false1194:
	v534 = *libc.As[int32](lookahead)
	cmp1195 = 97 <= v534
	if cmp1195 {
		goto land_lhs_true1197
	} else {
		goto if_end1201
	}

land_lhs_true1197:
	v535 = *libc.As[int32](lookahead)
	cmp1198 = v535 <= 122
	if cmp1198 {
		goto if_then1200
	} else {
		goto if_end1201
	}

if_then1200:
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end1201:
	v536 = *libc.As[byte](result)
	loadedv1202 = (v536 & 1) != 0
	*libc.As[bool](retval) = loadedv1202
	goto _return

sw_bb1203:
	*libc.As[byte](result) = 1
	v537 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1204 = libc.Ptr(&libc.As[TSLexer](v537).F1)
	*libc.As[int16](result_symbol1204) = 11
	v538 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1205 = libc.Ptr(&libc.As[TSLexer](v538).F3)
	v539 = *libc.As[unsafe.Pointer](mark_end1205)
	v540 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v539)(v540)
	v541 = *libc.As[int32](lookahead)
	cmp1206 = v541 == 59
	if cmp1206 {
		goto if_then1208
	} else {
		goto if_end1209
	}

if_then1208:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1209:
	v542 = *libc.As[int32](lookahead)
	cmp1210 = 65 <= v542
	if cmp1210 {
		goto land_lhs_true1212
	} else {
		goto lor_lhs_false1215
	}

land_lhs_true1212:
	v543 = *libc.As[int32](lookahead)
	cmp1213 = v543 <= 90
	if cmp1213 {
		goto if_then1221
	} else {
		goto lor_lhs_false1215
	}

lor_lhs_false1215:
	v544 = *libc.As[int32](lookahead)
	cmp1216 = 97 <= v544
	if cmp1216 {
		goto land_lhs_true1218
	} else {
		goto if_end1222
	}

land_lhs_true1218:
	v545 = *libc.As[int32](lookahead)
	cmp1219 = v545 <= 122
	if cmp1219 {
		goto if_then1221
	} else {
		goto if_end1222
	}

if_then1221:
	*libc.As[int16](state_addr) = 66
	goto next_state

if_end1222:
	v546 = *libc.As[byte](result)
	loadedv1223 = (v546 & 1) != 0
	*libc.As[bool](retval) = loadedv1223
	goto _return

sw_bb1224:
	*libc.As[byte](result) = 1
	v547 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1225 = libc.Ptr(&libc.As[TSLexer](v547).F1)
	*libc.As[int16](result_symbol1225) = 11
	v548 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1226 = libc.Ptr(&libc.As[TSLexer](v548).F3)
	v549 = *libc.As[unsafe.Pointer](mark_end1226)
	v550 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v549)(v550)
	v551 = *libc.As[int32](lookahead)
	cmp1227 = v551 == 59
	if cmp1227 {
		goto if_then1229
	} else {
		goto if_end1230
	}

if_then1229:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1230:
	v552 = *libc.As[int32](lookahead)
	cmp1231 = 65 <= v552
	if cmp1231 {
		goto land_lhs_true1233
	} else {
		goto lor_lhs_false1236
	}

land_lhs_true1233:
	v553 = *libc.As[int32](lookahead)
	cmp1234 = v553 <= 90
	if cmp1234 {
		goto if_then1242
	} else {
		goto lor_lhs_false1236
	}

lor_lhs_false1236:
	v554 = *libc.As[int32](lookahead)
	cmp1237 = 97 <= v554
	if cmp1237 {
		goto land_lhs_true1239
	} else {
		goto if_end1243
	}

land_lhs_true1239:
	v555 = *libc.As[int32](lookahead)
	cmp1240 = v555 <= 122
	if cmp1240 {
		goto if_then1242
	} else {
		goto if_end1243
	}

if_then1242:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end1243:
	v556 = *libc.As[byte](result)
	loadedv1244 = (v556 & 1) != 0
	*libc.As[bool](retval) = loadedv1244
	goto _return

sw_bb1245:
	*libc.As[byte](result) = 1
	v557 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1246 = libc.Ptr(&libc.As[TSLexer](v557).F1)
	*libc.As[int16](result_symbol1246) = 11
	v558 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1247 = libc.Ptr(&libc.As[TSLexer](v558).F3)
	v559 = *libc.As[unsafe.Pointer](mark_end1247)
	v560 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v559)(v560)
	v561 = *libc.As[int32](lookahead)
	cmp1248 = v561 == 59
	if cmp1248 {
		goto if_then1250
	} else {
		goto if_end1251
	}

if_then1250:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1251:
	v562 = *libc.As[int32](lookahead)
	cmp1252 = 65 <= v562
	if cmp1252 {
		goto land_lhs_true1254
	} else {
		goto lor_lhs_false1257
	}

land_lhs_true1254:
	v563 = *libc.As[int32](lookahead)
	cmp1255 = v563 <= 90
	if cmp1255 {
		goto if_then1263
	} else {
		goto lor_lhs_false1257
	}

lor_lhs_false1257:
	v564 = *libc.As[int32](lookahead)
	cmp1258 = 97 <= v564
	if cmp1258 {
		goto land_lhs_true1260
	} else {
		goto if_end1264
	}

land_lhs_true1260:
	v565 = *libc.As[int32](lookahead)
	cmp1261 = v565 <= 122
	if cmp1261 {
		goto if_then1263
	} else {
		goto if_end1264
	}

if_then1263:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end1264:
	v566 = *libc.As[byte](result)
	loadedv1265 = (v566 & 1) != 0
	*libc.As[bool](retval) = loadedv1265
	goto _return

sw_bb1266:
	*libc.As[byte](result) = 1
	v567 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1267 = libc.Ptr(&libc.As[TSLexer](v567).F1)
	*libc.As[int16](result_symbol1267) = 11
	v568 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1268 = libc.Ptr(&libc.As[TSLexer](v568).F3)
	v569 = *libc.As[unsafe.Pointer](mark_end1268)
	v570 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v569)(v570)
	v571 = *libc.As[int32](lookahead)
	cmp1269 = v571 == 59
	if cmp1269 {
		goto if_then1271
	} else {
		goto if_end1272
	}

if_then1271:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1272:
	v572 = *libc.As[int32](lookahead)
	cmp1273 = 65 <= v572
	if cmp1273 {
		goto land_lhs_true1275
	} else {
		goto lor_lhs_false1278
	}

land_lhs_true1275:
	v573 = *libc.As[int32](lookahead)
	cmp1276 = v573 <= 90
	if cmp1276 {
		goto if_then1284
	} else {
		goto lor_lhs_false1278
	}

lor_lhs_false1278:
	v574 = *libc.As[int32](lookahead)
	cmp1279 = 97 <= v574
	if cmp1279 {
		goto land_lhs_true1281
	} else {
		goto if_end1285
	}

land_lhs_true1281:
	v575 = *libc.As[int32](lookahead)
	cmp1282 = v575 <= 122
	if cmp1282 {
		goto if_then1284
	} else {
		goto if_end1285
	}

if_then1284:
	*libc.As[int16](state_addr) = 69
	goto next_state

if_end1285:
	v576 = *libc.As[byte](result)
	loadedv1286 = (v576 & 1) != 0
	*libc.As[bool](retval) = loadedv1286
	goto _return

sw_bb1287:
	*libc.As[byte](result) = 1
	v577 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1288 = libc.Ptr(&libc.As[TSLexer](v577).F1)
	*libc.As[int16](result_symbol1288) = 11
	v578 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1289 = libc.Ptr(&libc.As[TSLexer](v578).F3)
	v579 = *libc.As[unsafe.Pointer](mark_end1289)
	v580 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v579)(v580)
	v581 = *libc.As[int32](lookahead)
	cmp1290 = v581 == 59
	if cmp1290 {
		goto if_then1292
	} else {
		goto if_end1293
	}

if_then1292:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1293:
	v582 = *libc.As[int32](lookahead)
	cmp1294 = 65 <= v582
	if cmp1294 {
		goto land_lhs_true1296
	} else {
		goto lor_lhs_false1299
	}

land_lhs_true1296:
	v583 = *libc.As[int32](lookahead)
	cmp1297 = v583 <= 90
	if cmp1297 {
		goto if_then1305
	} else {
		goto lor_lhs_false1299
	}

lor_lhs_false1299:
	v584 = *libc.As[int32](lookahead)
	cmp1300 = 97 <= v584
	if cmp1300 {
		goto land_lhs_true1302
	} else {
		goto if_end1306
	}

land_lhs_true1302:
	v585 = *libc.As[int32](lookahead)
	cmp1303 = v585 <= 122
	if cmp1303 {
		goto if_then1305
	} else {
		goto if_end1306
	}

if_then1305:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end1306:
	v586 = *libc.As[byte](result)
	loadedv1307 = (v586 & 1) != 0
	*libc.As[bool](retval) = loadedv1307
	goto _return

sw_bb1308:
	*libc.As[byte](result) = 1
	v587 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1309 = libc.Ptr(&libc.As[TSLexer](v587).F1)
	*libc.As[int16](result_symbol1309) = 11
	v588 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1310 = libc.Ptr(&libc.As[TSLexer](v588).F3)
	v589 = *libc.As[unsafe.Pointer](mark_end1310)
	v590 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v589)(v590)
	v591 = *libc.As[int32](lookahead)
	cmp1311 = v591 == 59
	if cmp1311 {
		goto if_then1313
	} else {
		goto if_end1314
	}

if_then1313:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1314:
	v592 = *libc.As[int32](lookahead)
	cmp1315 = 65 <= v592
	if cmp1315 {
		goto land_lhs_true1317
	} else {
		goto lor_lhs_false1320
	}

land_lhs_true1317:
	v593 = *libc.As[int32](lookahead)
	cmp1318 = v593 <= 90
	if cmp1318 {
		goto if_then1326
	} else {
		goto lor_lhs_false1320
	}

lor_lhs_false1320:
	v594 = *libc.As[int32](lookahead)
	cmp1321 = 97 <= v594
	if cmp1321 {
		goto land_lhs_true1323
	} else {
		goto if_end1327
	}

land_lhs_true1323:
	v595 = *libc.As[int32](lookahead)
	cmp1324 = v595 <= 122
	if cmp1324 {
		goto if_then1326
	} else {
		goto if_end1327
	}

if_then1326:
	*libc.As[int16](state_addr) = 71
	goto next_state

if_end1327:
	v596 = *libc.As[byte](result)
	loadedv1328 = (v596 & 1) != 0
	*libc.As[bool](retval) = loadedv1328
	goto _return

sw_bb1329:
	*libc.As[byte](result) = 1
	v597 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1330 = libc.Ptr(&libc.As[TSLexer](v597).F1)
	*libc.As[int16](result_symbol1330) = 11
	v598 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1331 = libc.Ptr(&libc.As[TSLexer](v598).F3)
	v599 = *libc.As[unsafe.Pointer](mark_end1331)
	v600 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v599)(v600)
	v601 = *libc.As[int32](lookahead)
	cmp1332 = v601 == 59
	if cmp1332 {
		goto if_then1334
	} else {
		goto if_end1335
	}

if_then1334:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1335:
	v602 = *libc.As[int32](lookahead)
	cmp1336 = 65 <= v602
	if cmp1336 {
		goto land_lhs_true1338
	} else {
		goto lor_lhs_false1341
	}

land_lhs_true1338:
	v603 = *libc.As[int32](lookahead)
	cmp1339 = v603 <= 90
	if cmp1339 {
		goto if_then1347
	} else {
		goto lor_lhs_false1341
	}

lor_lhs_false1341:
	v604 = *libc.As[int32](lookahead)
	cmp1342 = 97 <= v604
	if cmp1342 {
		goto land_lhs_true1344
	} else {
		goto if_end1348
	}

land_lhs_true1344:
	v605 = *libc.As[int32](lookahead)
	cmp1345 = v605 <= 122
	if cmp1345 {
		goto if_then1347
	} else {
		goto if_end1348
	}

if_then1347:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1348:
	v606 = *libc.As[byte](result)
	loadedv1349 = (v606 & 1) != 0
	*libc.As[bool](retval) = loadedv1349
	goto _return

sw_bb1350:
	*libc.As[byte](result) = 1
	v607 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1351 = libc.Ptr(&libc.As[TSLexer](v607).F1)
	*libc.As[int16](result_symbol1351) = 11
	v608 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1352 = libc.Ptr(&libc.As[TSLexer](v608).F3)
	v609 = *libc.As[unsafe.Pointer](mark_end1352)
	v610 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v609)(v610)
	v611 = *libc.As[int32](lookahead)
	cmp1353 = v611 == 59
	if cmp1353 {
		goto if_then1355
	} else {
		goto if_end1356
	}

if_then1355:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1356:
	v612 = *libc.As[int32](lookahead)
	cmp1357 = 65 <= v612
	if cmp1357 {
		goto land_lhs_true1359
	} else {
		goto lor_lhs_false1362
	}

land_lhs_true1359:
	v613 = *libc.As[int32](lookahead)
	cmp1360 = v613 <= 90
	if cmp1360 {
		goto if_then1368
	} else {
		goto lor_lhs_false1362
	}

lor_lhs_false1362:
	v614 = *libc.As[int32](lookahead)
	cmp1363 = 97 <= v614
	if cmp1363 {
		goto land_lhs_true1365
	} else {
		goto if_end1369
	}

land_lhs_true1365:
	v615 = *libc.As[int32](lookahead)
	cmp1366 = v615 <= 122
	if cmp1366 {
		goto if_then1368
	} else {
		goto if_end1369
	}

if_then1368:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end1369:
	v616 = *libc.As[byte](result)
	loadedv1370 = (v616 & 1) != 0
	*libc.As[bool](retval) = loadedv1370
	goto _return

sw_bb1371:
	*libc.As[byte](result) = 1
	v617 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1372 = libc.Ptr(&libc.As[TSLexer](v617).F1)
	*libc.As[int16](result_symbol1372) = 11
	v618 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1373 = libc.Ptr(&libc.As[TSLexer](v618).F3)
	v619 = *libc.As[unsafe.Pointer](mark_end1373)
	v620 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v619)(v620)
	v621 = *libc.As[int32](lookahead)
	cmp1374 = v621 == 59
	if cmp1374 {
		goto if_then1376
	} else {
		goto if_end1377
	}

if_then1376:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1377:
	v622 = *libc.As[int32](lookahead)
	cmp1378 = 65 <= v622
	if cmp1378 {
		goto land_lhs_true1380
	} else {
		goto lor_lhs_false1383
	}

land_lhs_true1380:
	v623 = *libc.As[int32](lookahead)
	cmp1381 = v623 <= 90
	if cmp1381 {
		goto if_then1389
	} else {
		goto lor_lhs_false1383
	}

lor_lhs_false1383:
	v624 = *libc.As[int32](lookahead)
	cmp1384 = 97 <= v624
	if cmp1384 {
		goto land_lhs_true1386
	} else {
		goto if_end1390
	}

land_lhs_true1386:
	v625 = *libc.As[int32](lookahead)
	cmp1387 = v625 <= 122
	if cmp1387 {
		goto if_then1389
	} else {
		goto if_end1390
	}

if_then1389:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end1390:
	v626 = *libc.As[byte](result)
	loadedv1391 = (v626 & 1) != 0
	*libc.As[bool](retval) = loadedv1391
	goto _return

sw_bb1392:
	*libc.As[byte](result) = 1
	v627 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1393 = libc.Ptr(&libc.As[TSLexer](v627).F1)
	*libc.As[int16](result_symbol1393) = 12
	v628 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1394 = libc.Ptr(&libc.As[TSLexer](v628).F3)
	v629 = *libc.As[unsafe.Pointer](mark_end1394)
	v630 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v629)(v630)
	v631 = *libc.As[byte](result)
	loadedv1395 = (v631 & 1) != 0
	*libc.As[bool](retval) = loadedv1395
	goto _return

sw_bb1396:
	*libc.As[byte](result) = 1
	v632 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1397 = libc.Ptr(&libc.As[TSLexer](v632).F1)
	*libc.As[int16](result_symbol1397) = 13
	v633 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1398 = libc.Ptr(&libc.As[TSLexer](v633).F3)
	v634 = *libc.As[unsafe.Pointer](mark_end1398)
	v635 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v634)(v635)
	v636 = *libc.As[int32](lookahead)
	cmp1399 = 9 <= v636
	if cmp1399 {
		goto land_lhs_true1401
	} else {
		goto lor_lhs_false1404
	}

land_lhs_true1401:
	v637 = *libc.As[int32](lookahead)
	cmp1402 = v637 <= 13
	if cmp1402 {
		goto if_then1407
	} else {
		goto lor_lhs_false1404
	}

lor_lhs_false1404:
	v638 = *libc.As[int32](lookahead)
	cmp1405 = v638 == 32
	if cmp1405 {
		goto if_then1407
	} else {
		goto if_end1408
	}

if_then1407:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end1408:
	v639 = *libc.As[int32](lookahead)
	cmp1409 = v639 != 0
	if cmp1409 {
		goto land_lhs_true1411
	} else {
		goto if_end1415
	}

land_lhs_true1411:
	v640 = *libc.As[int32](lookahead)
	cmp1412 = v640 != 39
	if cmp1412 {
		goto if_then1414
	} else {
		goto if_end1415
	}

if_then1414:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end1415:
	v641 = *libc.As[byte](result)
	loadedv1416 = (v641 & 1) != 0
	*libc.As[bool](retval) = loadedv1416
	goto _return

sw_bb1417:
	*libc.As[byte](result) = 1
	v642 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1418 = libc.Ptr(&libc.As[TSLexer](v642).F1)
	*libc.As[int16](result_symbol1418) = 13
	v643 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1419 = libc.Ptr(&libc.As[TSLexer](v643).F3)
	v644 = *libc.As[unsafe.Pointer](mark_end1419)
	v645 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v644)(v645)
	v646 = *libc.As[int32](lookahead)
	cmp1420 = v646 != 0
	if cmp1420 {
		goto land_lhs_true1422
	} else {
		goto if_end1426
	}

land_lhs_true1422:
	v647 = *libc.As[int32](lookahead)
	cmp1423 = v647 != 39
	if cmp1423 {
		goto if_then1425
	} else {
		goto if_end1426
	}

if_then1425:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end1426:
	v648 = *libc.As[byte](result)
	loadedv1427 = (v648 & 1) != 0
	*libc.As[bool](retval) = loadedv1427
	goto _return

sw_bb1428:
	*libc.As[byte](result) = 1
	v649 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1429 = libc.Ptr(&libc.As[TSLexer](v649).F1)
	*libc.As[int16](result_symbol1429) = 14
	v650 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1430 = libc.Ptr(&libc.As[TSLexer](v650).F3)
	v651 = *libc.As[unsafe.Pointer](mark_end1430)
	v652 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v651)(v652)
	v653 = *libc.As[byte](result)
	loadedv1431 = (v653 & 1) != 0
	*libc.As[bool](retval) = loadedv1431
	goto _return

sw_bb1432:
	*libc.As[byte](result) = 1
	v654 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1433 = libc.Ptr(&libc.As[TSLexer](v654).F1)
	*libc.As[int16](result_symbol1433) = 15
	v655 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1434 = libc.Ptr(&libc.As[TSLexer](v655).F3)
	v656 = *libc.As[unsafe.Pointer](mark_end1434)
	v657 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v656)(v657)
	v658 = *libc.As[int32](lookahead)
	cmp1435 = 9 <= v658
	if cmp1435 {
		goto land_lhs_true1437
	} else {
		goto lor_lhs_false1440
	}

land_lhs_true1437:
	v659 = *libc.As[int32](lookahead)
	cmp1438 = v659 <= 13
	if cmp1438 {
		goto if_then1443
	} else {
		goto lor_lhs_false1440
	}

lor_lhs_false1440:
	v660 = *libc.As[int32](lookahead)
	cmp1441 = v660 == 32
	if cmp1441 {
		goto if_then1443
	} else {
		goto if_end1444
	}

if_then1443:
	*libc.As[int16](state_addr) = 80
	goto next_state

if_end1444:
	v661 = *libc.As[int32](lookahead)
	cmp1445 = v661 != 0
	if cmp1445 {
		goto land_lhs_true1447
	} else {
		goto if_end1451
	}

land_lhs_true1447:
	v662 = *libc.As[int32](lookahead)
	cmp1448 = v662 != 34
	if cmp1448 {
		goto if_then1450
	} else {
		goto if_end1451
	}

if_then1450:
	*libc.As[int16](state_addr) = 81
	goto next_state

if_end1451:
	v663 = *libc.As[byte](result)
	loadedv1452 = (v663 & 1) != 0
	*libc.As[bool](retval) = loadedv1452
	goto _return

sw_bb1453:
	*libc.As[byte](result) = 1
	v664 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1454 = libc.Ptr(&libc.As[TSLexer](v664).F1)
	*libc.As[int16](result_symbol1454) = 15
	v665 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1455 = libc.Ptr(&libc.As[TSLexer](v665).F3)
	v666 = *libc.As[unsafe.Pointer](mark_end1455)
	v667 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v666)(v667)
	v668 = *libc.As[int32](lookahead)
	cmp1456 = v668 != 0
	if cmp1456 {
		goto land_lhs_true1458
	} else {
		goto if_end1462
	}

land_lhs_true1458:
	v669 = *libc.As[int32](lookahead)
	cmp1459 = v669 != 34
	if cmp1459 {
		goto if_then1461
	} else {
		goto if_end1462
	}

if_then1461:
	*libc.As[int16](state_addr) = 81
	goto next_state

if_end1462:
	v670 = *libc.As[byte](result)
	loadedv1463 = (v670 & 1) != 0
	*libc.As[bool](retval) = loadedv1463
	goto _return

sw_bb1464:
	*libc.As[byte](result) = 1
	v671 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1465 = libc.Ptr(&libc.As[TSLexer](v671).F1)
	*libc.As[int16](result_symbol1465) = 16
	v672 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1466 = libc.Ptr(&libc.As[TSLexer](v672).F3)
	v673 = *libc.As[unsafe.Pointer](mark_end1466)
	v674 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v673)(v674)
	v675 = *libc.As[int32](lookahead)
	cmp1467 = v675 == 45
	if cmp1467 {
		goto if_then1469
	} else {
		goto if_end1470
	}

if_then1469:
	*libc.As[int16](state_addr) = 85
	goto next_state

if_end1470:
	v676 = *libc.As[int32](lookahead)
	cmp1471 = 9 <= v676
	if cmp1471 {
		goto land_lhs_true1473
	} else {
		goto lor_lhs_false1476
	}

land_lhs_true1473:
	v677 = *libc.As[int32](lookahead)
	cmp1474 = v677 <= 13
	if cmp1474 {
		goto if_then1479
	} else {
		goto lor_lhs_false1476
	}

lor_lhs_false1476:
	v678 = *libc.As[int32](lookahead)
	cmp1477 = v678 == 32
	if cmp1477 {
		goto if_then1479
	} else {
		goto if_end1480
	}

if_then1479:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end1480:
	v679 = *libc.As[int32](lookahead)
	cmp1481 = v679 != 0
	if cmp1481 {
		goto land_lhs_true1483
	} else {
		goto if_end1496
	}

land_lhs_true1483:
	v680 = *libc.As[int32](lookahead)
	cmp1484 = v680 != 60
	if cmp1484 {
		goto land_lhs_true1486
	} else {
		goto if_end1496
	}

land_lhs_true1486:
	v681 = *libc.As[int32](lookahead)
	cmp1487 = v681 != 62
	if cmp1487 {
		goto land_lhs_true1489
	} else {
		goto if_end1496
	}

land_lhs_true1489:
	v682 = *libc.As[int32](lookahead)
	cmp1490 = v682 != 123
	if cmp1490 {
		goto land_lhs_true1492
	} else {
		goto if_end1496
	}

land_lhs_true1492:
	v683 = *libc.As[int32](lookahead)
	cmp1493 = v683 != 125
	if cmp1493 {
		goto if_then1495
	} else {
		goto if_end1496
	}

if_then1495:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end1496:
	v684 = *libc.As[byte](result)
	loadedv1497 = (v684 & 1) != 0
	*libc.As[bool](retval) = loadedv1497
	goto _return

sw_bb1498:
	*libc.As[byte](result) = 1
	v685 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1499 = libc.Ptr(&libc.As[TSLexer](v685).F1)
	*libc.As[int16](result_symbol1499) = 16
	v686 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1500 = libc.Ptr(&libc.As[TSLexer](v686).F3)
	v687 = *libc.As[unsafe.Pointer](mark_end1500)
	v688 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v687)(v688)
	v689 = *libc.As[int32](lookahead)
	cmp1501 = v689 == 45
	if cmp1501 {
		goto if_then1503
	} else {
		goto if_end1504
	}

if_then1503:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end1504:
	v690 = *libc.As[int32](lookahead)
	cmp1505 = 9 <= v690
	if cmp1505 {
		goto land_lhs_true1507
	} else {
		goto lor_lhs_false1510
	}

land_lhs_true1507:
	v691 = *libc.As[int32](lookahead)
	cmp1508 = v691 <= 13
	if cmp1508 {
		goto if_then1513
	} else {
		goto lor_lhs_false1510
	}

lor_lhs_false1510:
	v692 = *libc.As[int32](lookahead)
	cmp1511 = v692 == 32
	if cmp1511 {
		goto if_then1513
	} else {
		goto if_end1514
	}

if_then1513:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end1514:
	v693 = *libc.As[int32](lookahead)
	cmp1515 = v693 != 0
	if cmp1515 {
		goto land_lhs_true1517
	} else {
		goto if_end1530
	}

land_lhs_true1517:
	v694 = *libc.As[int32](lookahead)
	cmp1518 = v694 != 60
	if cmp1518 {
		goto land_lhs_true1520
	} else {
		goto if_end1530
	}

land_lhs_true1520:
	v695 = *libc.As[int32](lookahead)
	cmp1521 = v695 != 62
	if cmp1521 {
		goto land_lhs_true1523
	} else {
		goto if_end1530
	}

land_lhs_true1523:
	v696 = *libc.As[int32](lookahead)
	cmp1524 = v696 != 123
	if cmp1524 {
		goto land_lhs_true1526
	} else {
		goto if_end1530
	}

land_lhs_true1526:
	v697 = *libc.As[int32](lookahead)
	cmp1527 = v697 != 125
	if cmp1527 {
		goto if_then1529
	} else {
		goto if_end1530
	}

if_then1529:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end1530:
	v698 = *libc.As[byte](result)
	loadedv1531 = (v698 & 1) != 0
	*libc.As[bool](retval) = loadedv1531
	goto _return

sw_bb1532:
	*libc.As[byte](result) = 1
	v699 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1533 = libc.Ptr(&libc.As[TSLexer](v699).F1)
	*libc.As[int16](result_symbol1533) = 16
	v700 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1534 = libc.Ptr(&libc.As[TSLexer](v700).F3)
	v701 = *libc.As[unsafe.Pointer](mark_end1534)
	v702 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v701)(v702)
	v703 = *libc.As[int32](lookahead)
	cmp1535 = 9 <= v703
	if cmp1535 {
		goto land_lhs_true1537
	} else {
		goto lor_lhs_false1540
	}

land_lhs_true1537:
	v704 = *libc.As[int32](lookahead)
	cmp1538 = v704 <= 13
	if cmp1538 {
		goto if_then1543
	} else {
		goto lor_lhs_false1540
	}

lor_lhs_false1540:
	v705 = *libc.As[int32](lookahead)
	cmp1541 = v705 == 32
	if cmp1541 {
		goto if_then1543
	} else {
		goto if_end1544
	}

if_then1543:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end1544:
	v706 = *libc.As[int32](lookahead)
	cmp1545 = v706 != 0
	if cmp1545 {
		goto land_lhs_true1547
	} else {
		goto if_end1560
	}

land_lhs_true1547:
	v707 = *libc.As[int32](lookahead)
	cmp1548 = v707 != 60
	if cmp1548 {
		goto land_lhs_true1550
	} else {
		goto if_end1560
	}

land_lhs_true1550:
	v708 = *libc.As[int32](lookahead)
	cmp1551 = v708 != 62
	if cmp1551 {
		goto land_lhs_true1553
	} else {
		goto if_end1560
	}

land_lhs_true1553:
	v709 = *libc.As[int32](lookahead)
	cmp1554 = v709 != 123
	if cmp1554 {
		goto land_lhs_true1556
	} else {
		goto if_end1560
	}

land_lhs_true1556:
	v710 = *libc.As[int32](lookahead)
	cmp1557 = v710 != 125
	if cmp1557 {
		goto if_then1559
	} else {
		goto if_end1560
	}

if_then1559:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end1560:
	v711 = *libc.As[byte](result)
	loadedv1561 = (v711 & 1) != 0
	*libc.As[bool](retval) = loadedv1561
	goto _return

sw_bb1562:
	*libc.As[byte](result) = 1
	v712 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1563 = libc.Ptr(&libc.As[TSLexer](v712).F1)
	*libc.As[int16](result_symbol1563) = 17
	v713 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1564 = libc.Ptr(&libc.As[TSLexer](v713).F3)
	v714 = *libc.As[unsafe.Pointer](mark_end1564)
	v715 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v714)(v715)
	v716 = *libc.As[byte](result)
	loadedv1565 = (v716 & 1) != 0
	*libc.As[bool](retval) = loadedv1565
	goto _return

sw_bb1566:
	*libc.As[byte](result) = 1
	v717 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1567 = libc.Ptr(&libc.As[TSLexer](v717).F1)
	*libc.As[int16](result_symbol1567) = 18
	v718 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1568 = libc.Ptr(&libc.As[TSLexer](v718).F3)
	v719 = *libc.As[unsafe.Pointer](mark_end1568)
	v720 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v719)(v720)
	v721 = *libc.As[byte](result)
	loadedv1569 = (v721 & 1) != 0
	*libc.As[bool](retval) = loadedv1569
	goto _return

sw_bb1570:
	*libc.As[byte](result) = 1
	v722 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1571 = libc.Ptr(&libc.As[TSLexer](v722).F1)
	*libc.As[int16](result_symbol1571) = 19
	v723 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1572 = libc.Ptr(&libc.As[TSLexer](v723).F3)
	v724 = *libc.As[unsafe.Pointer](mark_end1572)
	v725 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v724)(v725)
	v726 = *libc.As[byte](result)
	loadedv1573 = (v726 & 1) != 0
	*libc.As[bool](retval) = loadedv1573
	goto _return

sw_bb1574:
	*libc.As[byte](result) = 1
	v727 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1575 = libc.Ptr(&libc.As[TSLexer](v727).F1)
	*libc.As[int16](result_symbol1575) = 20
	v728 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1576 = libc.Ptr(&libc.As[TSLexer](v728).F3)
	v729 = *libc.As[unsafe.Pointer](mark_end1576)
	v730 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v729)(v730)
	v731 = *libc.As[byte](result)
	loadedv1577 = (v731 & 1) != 0
	*libc.As[bool](retval) = loadedv1577
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v732 = *libc.As[bool](retval)
	return v732
}
func scan_js_expr_with_delimiter(lexer unsafe.Pointer, end_type int32) {
	var cmp, cmp1, cmp2, cmp4, cmp8, cmp12, cmp18, cmp22, cmp27, cmp33, cmp36, cmp46, cmp49, cmp53, cmp58, cmp62, cmp67, cmp74, cmp78, cmp81, cmp87, cmp91, cmp96 bool
	var v4, v5, v6, v7, v12, v13, conv, v15, inc, v16, v21, cond, v23, v28, v29, inc30, v31, v32, v36, dec, v38, v40, v42, v45, v50, v52, v53, v55, v56, cond83, v61, v63, v68 int32
	var end_type_addr, delimiter_index, curly_count, in_comment, lookahead, lookahead7, lookahead17, lookahead26, lookahead32, lookahead45, lookahead48, lookahead52, lookahead57, lookahead61, lookahead66, lookahead77, lookahead90, lookahead95 unsafe.Pointer
	var idxprom, conv11, v22, v57 int64
	var v14 byte
	var arrayidx unsafe.Pointer
	var v0, v1, v2, v3, v8, v9, v10, v11, v17, v18, v19, v20, v24, v25, v26, v27, v30, v33, v34, v35, v37, v39, v41, v43, v44, v46, v47, v48, v49, v51, v54, v58, v59, v60, v62, v64, v65, v66, v67, v69, v70, v71 unsafe.Pointer
	var lexer_addr, END, mark_end, mark_end6, mark_end16, mark_end25, mark_end39, local_advance, mark_end84, advance94, advance105 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = lexer_addr, end_type_addr, delimiter_index, curly_count, in_comment, END, v0, mark_end, v1, v2, v3, lookahead, v4, cmp, v5, cmp1, v6, cmp2, v7, cmp4, v8, mark_end6, v9, v10, v11, lookahead7, v12, v13, idxprom, arrayidx, v14, conv, cmp8, v15, inc, v16, conv11, cmp12, v17, mark_end16, v18, v19, v20, lookahead17, v21, cmp18, v22, cond, v23, cmp22, v24, mark_end25, v25, v26, v27, lookahead26, v28, cmp27, v29, inc30, v30, lookahead32, v31, cmp33, v32, cmp36, v33, mark_end39, v34, v35, v36, dec, v37, lookahead45, v38, cmp46, v39, lookahead48, v40, cmp49, v41, lookahead52, v42, cmp53, v43, v44, lookahead57, v45, cmp58, v46, local_advance, v47, v48, v49, lookahead61, v50, cmp62, v51, lookahead66, v52, cmp67, v53, cmp74, v54, lookahead77, v55, cmp78, v56, cmp81, v57, cond83, v58, mark_end84, v59, v60, v61, cmp87, v62, lookahead90, v63, cmp91, v64, advance94, v65, v66, v67, lookahead95, v68, cmp96, v69, advance105, v70, v71

	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	end_type_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	delimiter_index = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	curly_count = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	in_comment = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	END = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[int32](end_type_addr) = end_type
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v0).F3)
	v1 = *libc.As[unsafe.Pointer](mark_end)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1)(v2)
	*libc.As[int32](delimiter_index) = 1
	*libc.As[int32](curly_count) = 0
	*libc.As[int32](in_comment) = 0
	goto while_cond

while_cond:
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v3).F0)
	v4 = *libc.As[int32](lookahead)
	cmp = v4 != 0
	if cmp {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v5 = *libc.As[int32](in_comment)
	cmp1 = v5 == 0
	if cmp1 {
		goto if_then
	} else {
		goto if_else73
	}

if_then:
	v6 = *libc.As[int32](end_type_addr)
	cmp2 = v6 == 0
	if cmp2 {
		goto if_then3
	} else {
		goto if_else21
	}

if_then3:
	v7 = *libc.As[int32](delimiter_index)
	cmp4 = v7 == 0
	if cmp4 {
		goto if_then5
	} else {
		goto if_end
	}

if_then5:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end6 = libc.Ptr(&libc.As[TSLexer](v8).F3)
	v9 = *libc.As[unsafe.Pointer](mark_end6)
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v9)(v10)
	goto if_end

if_end:
	*libc.As[unsafe.Pointer](END) = libc.Ptr(&_str_2)
	v11 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead7 = libc.Ptr(&libc.As[TSLexer](v11).F0)
	v12 = *libc.As[int32](lookahead7)
	v13 = *libc.As[int32](delimiter_index)
	idxprom = int64(uint64(uint32(v13)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](libc.Ptr(&_str_2)), int(idxprom)*1))
	v14 = *libc.As[byte](arrayidx)
	conv = int32(int8(v14))
	cmp8 = v12 == conv
	if cmp8 {
		goto if_then10
	} else {
		goto if_else
	}

if_then10:
	v15 = *libc.As[int32](delimiter_index)
	inc = v15 + 1
	*libc.As[int32](delimiter_index) = inc
	v16 = *libc.As[int32](delimiter_index)
	conv11 = int64(uint64(uint32(v16)))
	cmp12 = conv11 == int64(4)
	if cmp12 {
		goto if_then14
	} else {
		goto if_end15
	}

if_then14:
	goto while_end

if_end15:
	goto if_end20

if_else:
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end16 = libc.Ptr(&libc.As[TSLexer](v17).F3)
	v18 = *libc.As[unsafe.Pointer](mark_end16)
	v19 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v18)(v19)
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead17 = libc.Ptr(&libc.As[TSLexer](v20).F0)
	v21 = *libc.As[int32](lookahead17)
	cmp18 = v21 == 10
	if cmp18 {
		v22 = 1
	} else {
		v22 = 0
	}
	if cmp18 {
		cond = 1
	} else {
		cond = 0
	}
	*libc.As[int32](delimiter_index) = cond
	goto if_end20

if_end20:
	goto if_end44

if_else21:
	v23 = *libc.As[int32](end_type_addr)
	cmp22 = v23 == 1
	if cmp22 {
		goto if_then24
	} else {
		goto if_end43
	}

if_then24:
	v24 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end25 = libc.Ptr(&libc.As[TSLexer](v24).F3)
	v25 = *libc.As[unsafe.Pointer](mark_end25)
	v26 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v25)(v26)
	v27 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead26 = libc.Ptr(&libc.As[TSLexer](v27).F0)
	v28 = *libc.As[int32](lookahead26)
	cmp27 = v28 == 123
	if cmp27 {
		goto if_then29
	} else {
		goto if_else31
	}

if_then29:
	v29 = *libc.As[int32](curly_count)
	inc30 = v29 + 1
	*libc.As[int32](curly_count) = inc30
	goto if_end42

if_else31:
	v30 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead32 = libc.Ptr(&libc.As[TSLexer](v30).F0)
	v31 = *libc.As[int32](lookahead32)
	cmp33 = v31 == 125
	if cmp33 {
		goto if_then35
	} else {
		goto if_end41
	}

if_then35:
	v32 = *libc.As[int32](curly_count)
	cmp36 = v32 == 0
	if cmp36 {
		goto if_then38
	} else {
		goto if_end40
	}

if_then38:
	v33 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end39 = libc.Ptr(&libc.As[TSLexer](v33).F3)
	v34 = *libc.As[unsafe.Pointer](mark_end39)
	v35 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v34)(v35)
	goto while_end

if_end40:
	v36 = *libc.As[int32](curly_count)
	dec = v36 - 1
	*libc.As[int32](curly_count) = dec
	goto if_end41

if_end41:
	goto if_end42

if_end42:
	goto if_end43

if_end43:
	goto if_end44

if_end44:
	v37 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead45 = libc.Ptr(&libc.As[TSLexer](v37).F0)
	v38 = *libc.As[int32](lookahead45)
	cmp46 = v38 == 34
	if cmp46 {
		goto if_then55
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v39 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead48 = libc.Ptr(&libc.As[TSLexer](v39).F0)
	v40 = *libc.As[int32](lookahead48)
	cmp49 = v40 == 39
	if cmp49 {
		goto if_then55
	} else {
		goto lor_lhs_false51
	}

lor_lhs_false51:
	v41 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead52 = libc.Ptr(&libc.As[TSLexer](v41).F0)
	v42 = *libc.As[int32](lookahead52)
	cmp53 = v42 == 96
	if cmp53 {
		goto if_then55
	} else {
		goto if_end56
	}

if_then55:
	v43 = *libc.As[unsafe.Pointer](lexer_addr)
	scan_js_string(v43)
	goto while_cond

if_end56:
	v44 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead57 = libc.Ptr(&libc.As[TSLexer](v44).F0)
	v45 = *libc.As[int32](lookahead57)
	cmp58 = v45 == 47
	if cmp58 {
		goto if_then60
	} else {
		goto if_end72
	}

if_then60:
	v46 = *libc.As[unsafe.Pointer](lexer_addr)
	local_advance = libc.Ptr(&libc.As[TSLexer](v46).F2)
	v47 = *libc.As[unsafe.Pointer](local_advance)
	v48 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v47)(v48, false)
	v49 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead61 = libc.Ptr(&libc.As[TSLexer](v49).F0)
	v50 = *libc.As[int32](lookahead61)
	cmp62 = v50 == 47
	if cmp62 {
		goto if_then64
	} else {
		goto if_else65
	}

if_then64:
	*libc.As[int32](in_comment) = 1
	goto if_end71

if_else65:
	v51 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead66 = libc.Ptr(&libc.As[TSLexer](v51).F0)
	v52 = *libc.As[int32](lookahead66)
	cmp67 = v52 == 42
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*libc.As[int32](in_comment) = 2
	goto if_end70

if_end70:
	goto if_end71

if_end71:
	goto while_cond

if_end72:
	goto if_end104

if_else73:
	v53 = *libc.As[int32](in_comment)
	cmp74 = v53 == 1
	if cmp74 {
		goto if_then76
	} else {
		goto if_else86
	}

if_then76:
	v54 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead77 = libc.Ptr(&libc.As[TSLexer](v54).F0)
	v55 = *libc.As[int32](lookahead77)
	cmp78 = v55 == 10
	if cmp78 {
		goto if_then80
	} else {
		goto if_end85
	}

if_then80:
	*libc.As[int32](in_comment) = 0
	v56 = *libc.As[int32](end_type_addr)
	cmp81 = v56 == 0
	if cmp81 {
		v57 = 1
	} else {
		v57 = 0
	}
	if cmp81 {
		cond83 = 1
	} else {
		cond83 = 0
	}
	*libc.As[int32](delimiter_index) = cond83
	v58 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end84 = libc.Ptr(&libc.As[TSLexer](v58).F3)
	v59 = *libc.As[unsafe.Pointer](mark_end84)
	v60 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v59)(v60)
	goto if_end85

if_end85:
	goto if_end103

if_else86:
	v61 = *libc.As[int32](in_comment)
	cmp87 = v61 == 2
	if cmp87 {
		goto if_then89
	} else {
		goto if_end102
	}

if_then89:
	v62 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead90 = libc.Ptr(&libc.As[TSLexer](v62).F0)
	v63 = *libc.As[int32](lookahead90)
	cmp91 = v63 == 42
	if cmp91 {
		goto if_then93
	} else {
		goto if_end101
	}

if_then93:
	v64 = *libc.As[unsafe.Pointer](lexer_addr)
	advance94 = libc.Ptr(&libc.As[TSLexer](v64).F2)
	v65 = *libc.As[unsafe.Pointer](advance94)
	v66 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v65)(v66, false)
	v67 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead95 = libc.Ptr(&libc.As[TSLexer](v67).F0)
	v68 = *libc.As[int32](lookahead95)
	cmp96 = v68 == 47
	if cmp96 {
		goto if_then98
	} else {
		goto if_else99
	}

if_then98:
	*libc.As[int32](in_comment) = 0
	*libc.As[int32](delimiter_index) = 0
	goto if_end100

if_else99:
	goto while_cond

if_end100:
	goto if_end101

if_end101:
	goto if_end102

if_end102:
	goto if_end103

if_end103:
	goto if_end104

if_end104:
	v69 = *libc.As[unsafe.Pointer](lexer_addr)
	advance105 = libc.Ptr(&libc.As[TSLexer](v69).F2)
	v70 = *libc.As[unsafe.Pointer](advance105)
	v71 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v70)(v71, false)
	goto while_cond

while_end:
}
func scan_raw_text(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var arrayidx unsafe.Pointer
	var tags, tags1, tags3, tags4 unsafe.Pointer
	var cond unsafe.Pointer
	var cmp, cmp7, tobool, cmp11, cmp15 bool
	var result_symbol unsafe.Pointer
	var v4, sub, v6, v10, sub6, v11, v14, v16, v18, conv, v20, inc, v21 int32
	var delimiter_index, size, size2, size5, _type, lookahead, lookahead8 unsafe.Pointer
	var idxprom, v12, idxprom9, conv14, call int64
	var v19 byte
	var arrayidx10 unsafe.Pointer
	var v0, v1, v2, v3, v5, v7, v8, v9, v13, v15, v17, v22, v23, v24, v25, v26, v27, v28 unsafe.Pointer
	var scanner_addr, lexer_addr, end_delimiter, mark_end, contents, mark_end20 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = scanner_addr, lexer_addr, end_delimiter, delimiter_index, v0, mark_end, v1, v2, v3, tags, size, v4, sub, v5, tags1, size2, v6, cmp, v7, tags3, contents, v8, v9, tags4, size5, v10, sub6, idxprom, arrayidx, _type, v11, cmp7, v12, cond, v13, lookahead, v14, tobool, v15, lookahead8, v16, v17, v18, idxprom9, arrayidx10, v19, conv, cmp11, v20, inc, v21, conv14, v22, call, cmp15, v23, v24, v25, mark_end20, v26, v27, v28, result_symbol

	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	end_delimiter = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	delimiter_index = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v0).F3)
	v1 = *libc.As[unsafe.Pointer](mark_end)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1)(v2)
	v3 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v3).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v4 = *libc.As[int32](size)
	sub = v4 - 1
	v5 = *libc.As[unsafe.Pointer](scanner_addr)
	tags1 = libc.Ptr(&libc.As[Scanner](v5).F0)
	size2 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F1)
	v6 = *libc.As[int32](size2)
	cmp = uint32(sub) < uint32(v6)
	if cmp {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	goto if_end

if_else:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 315, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_raw_text)))
	panic("unreachable")

if_end:
	v7 = *libc.As[unsafe.Pointer](scanner_addr)
	tags3 = libc.Ptr(&libc.As[Scanner](v7).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags3).F0)
	v8 = *libc.As[unsafe.Pointer](contents)
	v9 = *libc.As[unsafe.Pointer](scanner_addr)
	tags4 = libc.Ptr(&libc.As[Scanner](v9).F0)
	size5 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags4).F1)
	v10 = *libc.As[int32](size5)
	sub6 = v10 - 1
	idxprom = int64(uint64(uint32(sub6)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v8), int(idxprom)*24))
	_type = libc.Ptr(&libc.As[Tag](arrayidx).F0)
	v11 = *libc.As[int32](_type)
	cmp7 = v11 == 97
	if cmp7 {
		v12 = 1
	} else {
		v12 = 0
	}
	if cmp7 {
		cond = libc.Ptr(&_str_3)
	} else {
		cond = libc.Ptr(&_str_4)
	}
	*libc.As[unsafe.Pointer](end_delimiter) = cond
	*libc.As[int32](delimiter_index) = 0
	goto while_cond

while_cond:
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v13).F0)
	v14 = *libc.As[int32](lookahead)
	tobool = v14 != 0
	if tobool {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead8 = libc.Ptr(&libc.As[TSLexer](v15).F0)
	v16 = *libc.As[int32](lookahead8)
	v17 = *libc.As[unsafe.Pointer](end_delimiter)
	v18 = *libc.As[int32](delimiter_index)
	idxprom9 = int64(uint64(uint32(v18)))
	arrayidx10 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v17), int(idxprom9)*1))
	v19 = *libc.As[byte](arrayidx10)
	conv = int32(int8(v19))
	cmp11 = v16 == conv
	if cmp11 {
		goto if_then13
	} else {
		goto if_else19
	}

if_then13:
	v20 = *libc.As[int32](delimiter_index)
	inc = v20 + 1
	*libc.As[int32](delimiter_index) = inc
	v21 = *libc.As[int32](delimiter_index)
	conv14 = int64(uint64(uint32(v21)))
	v22 = *libc.As[unsafe.Pointer](end_delimiter)
	call = libc.Strlen(libc.As[byte](v22))
	cmp15 = conv14 == call
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	goto while_end

if_end18:
	v23 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v23)
	goto if_end21

if_else19:
	*libc.As[int32](delimiter_index) = 0
	v24 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v24)
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end20 = libc.Ptr(&libc.As[TSLexer](v25).F3)
	v26 = *libc.As[unsafe.Pointer](mark_end20)
	v27 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v26)(v27)
	goto if_end21

if_end21:
	goto while_cond

while_end:
	v28 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v28).F1)
	*libc.As[int16](result_symbol) = 7
	return true
}
func scan_permissible_text(lexer unsafe.Pointer) bool {
	var cmp, cmp2, cmp4, cmp6, cmp9, cmp12, cmp16, cmp19, cmp23, cmp25, cmp27, v24, cmp31, cmp35, cmp38, cmp41, cmp49, cmp52, cmp55, cmp58, cmp61, cmp65, cmp69, cmp72, cmp76, loadedv, v62 bool
	var retval unsafe.Pointer
	var result_symbol unsafe.Pointer
	var v1, v3, v5, v7, v9, v11, v14, v17, v19, v21, v23, v27, v29, v32, v35, v38, v41, v43, v45, v47, v49, v51, v53, v55 int32
	var lookahead, lookahead1, lookahead3, lookahead5, lookahead8, lookahead11, lookahead15, lookahead18, lookahead22, lookahead24, lookahead26, lookahead30, lookahead34, lookahead37, lookahead40, lookahead48, lookahead51, lookahead54, lookahead57, lookahead60, lookahead64, lookahead68, lookahead71, lookahead75 unsafe.Pointer
	var v60 byte
	var there_is_text unsafe.Pointer
	var v0, v2, v4, v6, v8, v10, v12, v13, v15, v16, v18, v20, v22, v25, v26, v28, v30, v31, v33, v34, v36, v37, v39, v40, v42, v44, v46, v48, v50, v52, v54, v56, v57, v58, v59, v61 unsafe.Pointer
	var lexer_addr, mark_end unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, there_is_text, v0, lookahead, v1, cmp, v2, lookahead1, v3, cmp2, v4, lookahead3, v5, cmp4, v6, lookahead5, v7, cmp6, v8, lookahead8, v9, cmp9, v10, lookahead11, v11, cmp12, v12, v13, lookahead15, v14, cmp16, v15, v16, lookahead18, v17, cmp19, v18, lookahead22, v19, cmp23, v20, lookahead24, v21, cmp25, v22, lookahead26, v23, cmp27, v24, v25, v26, lookahead30, v27, cmp31, v28, lookahead34, v29, cmp35, v30, v31, lookahead37, v32, cmp38, v33, v34, lookahead40, v35, cmp41, v36, v37, lookahead48, v38, cmp49, v39, v40, lookahead51, v41, cmp52, v42, lookahead54, v43, cmp55, v44, lookahead57, v45, cmp58, v46, lookahead60, v47, cmp61, v48, lookahead64, v49, cmp65, v50, lookahead68, v51, cmp69, v52, lookahead71, v53, cmp72, v54, lookahead75, v55, cmp76, v56, v57, mark_end, v58, v59, v60, loadedv, v61, result_symbol, v62

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	there_is_text = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[byte](there_is_text) = 0
	goto while_cond

while_cond:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v0).F0)
	v1 = *libc.As[int32](lookahead)
	cmp = v1 != 0
	if cmp {
		goto while_body
	} else {
		goto while_end80
	}

while_body:
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v2).F0)
	v3 = *libc.As[int32](lookahead1)
	cmp2 = v3 == 123
	if cmp2 {
		goto if_then
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead3 = libc.Ptr(&libc.As[TSLexer](v4).F0)
	v5 = *libc.As[int32](lookahead3)
	cmp4 = v5 == 125
	if cmp4 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	goto while_end80

if_end:
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead5 = libc.Ptr(&libc.As[TSLexer](v6).F0)
	v7 = *libc.As[int32](lookahead5)
	cmp6 = v7 == 39
	if cmp6 {
		goto if_then13
	} else {
		goto lor_lhs_false7
	}

lor_lhs_false7:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead8 = libc.Ptr(&libc.As[TSLexer](v8).F0)
	v9 = *libc.As[int32](lookahead8)
	cmp9 = v9 == 34
	if cmp9 {
		goto if_then13
	} else {
		goto lor_lhs_false10
	}

lor_lhs_false10:
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead11 = libc.Ptr(&libc.As[TSLexer](v10).F0)
	v11 = *libc.As[int32](lookahead11)
	cmp12 = v11 == 96
	if cmp12 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	scan_js_string(v12)
	*libc.As[byte](there_is_text) = 1
	goto text_found

if_end14:
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead15 = libc.Ptr(&libc.As[TSLexer](v13).F0)
	v14 = *libc.As[int32](lookahead15)
	cmp16 = v14 == 47
	if cmp16 {
		goto if_then17
	} else {
		goto if_end47
	}

if_then17:
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v15)
	*libc.As[byte](there_is_text) = 1
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead18 = libc.Ptr(&libc.As[TSLexer](v16).F0)
	v17 = *libc.As[int32](lookahead18)
	cmp19 = v17 == 47
	if cmp19 {
		goto if_then20
	} else {
		goto if_end29
	}

if_then20:
	goto while_cond21

while_cond21:
	v18 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead22 = libc.Ptr(&libc.As[TSLexer](v18).F0)
	v19 = *libc.As[int32](lookahead22)
	cmp23 = v19 != 13
	if cmp23 {
		goto land_lhs_true
	} else {
		v24 = false
		goto land_end
	}

land_lhs_true:
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead24 = libc.Ptr(&libc.As[TSLexer](v20).F0)
	v21 = *libc.As[int32](lookahead24)
	cmp25 = v21 != 10
	if cmp25 {
		goto land_rhs
	} else {
		v24 = false
		goto land_end
	}

land_rhs:
	v22 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead26 = libc.Ptr(&libc.As[TSLexer](v22).F0)
	v23 = *libc.As[int32](lookahead26)
	cmp27 = v23 != 0
	v24 = cmp27
	goto land_end

land_end:
	if v24 {
		goto while_body28
	} else {
		goto while_end
	}

while_body28:
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v25)
	goto while_cond21

while_end:
	goto if_end29

if_end29:
	v26 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead30 = libc.Ptr(&libc.As[TSLexer](v26).F0)
	v27 = *libc.As[int32](lookahead30)
	cmp31 = v27 == 42
	if cmp31 {
		goto if_then32
	} else {
		goto if_end46
	}

if_then32:
	goto while_cond33

while_cond33:
	v28 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead34 = libc.Ptr(&libc.As[TSLexer](v28).F0)
	v29 = *libc.As[int32](lookahead34)
	cmp35 = v29 != 0
	if cmp35 {
		goto while_body36
	} else {
		goto while_end45
	}

while_body36:
	v30 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v30)
	v31 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead37 = libc.Ptr(&libc.As[TSLexer](v31).F0)
	v32 = *libc.As[int32](lookahead37)
	cmp38 = v32 == 42
	if cmp38 {
		goto if_then39
	} else {
		goto if_end44
	}

if_then39:
	v33 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v33)
	v34 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead40 = libc.Ptr(&libc.As[TSLexer](v34).F0)
	v35 = *libc.As[int32](lookahead40)
	cmp41 = v35 == 47
	if cmp41 {
		goto if_then42
	} else {
		goto if_end43
	}

if_then42:
	v36 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v36)
	goto while_end45

if_end43:
	goto if_end44

if_end44:
	goto while_cond33

while_end45:
	goto if_end46

if_end46:
	goto text_found

if_end47:
	v37 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead48 = libc.Ptr(&libc.As[TSLexer](v37).F0)
	v38 = *libc.As[int32](lookahead48)
	cmp49 = v38 == 60
	if cmp49 {
		goto if_then50
	} else {
		goto if_end79
	}

if_then50:
	v39 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v39)
	v40 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead51 = libc.Ptr(&libc.As[TSLexer](v40).F0)
	v41 = *libc.As[int32](lookahead51)
	cmp52 = 97 <= v41
	if cmp52 {
		goto land_lhs_true53
	} else {
		goto lor_lhs_false56
	}

land_lhs_true53:
	v42 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead54 = libc.Ptr(&libc.As[TSLexer](v42).F0)
	v43 = *libc.As[int32](lookahead54)
	cmp55 = v43 <= 122
	if cmp55 {
		goto if_then62
	} else {
		goto lor_lhs_false56
	}

lor_lhs_false56:
	v44 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead57 = libc.Ptr(&libc.As[TSLexer](v44).F0)
	v45 = *libc.As[int32](lookahead57)
	cmp58 = 65 <= v45
	if cmp58 {
		goto land_lhs_true59
	} else {
		goto if_end63
	}

land_lhs_true59:
	v46 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead60 = libc.Ptr(&libc.As[TSLexer](v46).F0)
	v47 = *libc.As[int32](lookahead60)
	cmp61 = v47 <= 90
	if cmp61 {
		goto if_then62
	} else {
		goto if_end63
	}

if_then62:
	goto while_end80

if_end63:
	v48 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead64 = libc.Ptr(&libc.As[TSLexer](v48).F0)
	v49 = *libc.As[int32](lookahead64)
	cmp65 = v49 == 47
	if cmp65 {
		goto if_then66
	} else {
		goto if_end67
	}

if_then66:
	goto while_end80

if_end67:
	v50 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead68 = libc.Ptr(&libc.As[TSLexer](v50).F0)
	v51 = *libc.As[int32](lookahead68)
	cmp69 = v51 == 47
	if cmp69 {
		goto if_then73
	} else {
		goto lor_lhs_false70
	}

lor_lhs_false70:
	v52 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead71 = libc.Ptr(&libc.As[TSLexer](v52).F0)
	v53 = *libc.As[int32](lookahead71)
	cmp72 = v53 == 63
	if cmp72 {
		goto if_then73
	} else {
		goto if_end74
	}

if_then73:
	goto while_end80

if_end74:
	v54 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead75 = libc.Ptr(&libc.As[TSLexer](v54).F0)
	v55 = *libc.As[int32](lookahead75)
	cmp76 = v55 == 62
	if cmp76 {
		goto if_then77
	} else {
		goto if_end78
	}

if_then77:
	goto while_end80

if_end78:
	goto text_found

if_end79:
	v56 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v56)
	goto text_found

text_found:
	*libc.As[byte](there_is_text) = 1
	v57 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v57).F3)
	v58 = *libc.As[unsafe.Pointer](mark_end)
	v59 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v58)(v59)
	goto while_cond

while_end80:
	v60 = *libc.As[byte](there_is_text)
	loadedv = (v60 & 1) != 0
	if loadedv {
		goto if_then81
	} else {
		goto if_else
	}

if_then81:
	v61 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v61).F1)
	*libc.As[int16](result_symbol) = 14
	*libc.As[bool](retval) = true
	goto _return

if_else:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v62 = *libc.As[bool](retval)
	return v62
}
func skip(lexer unsafe.Pointer) {
	var v0, v1, v2 unsafe.Pointer
	var lexer_addr, local_advance unsafe.Pointer
	_, _, _, _, _ = lexer_addr, v0, local_advance, v1, v2

	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	local_advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](local_advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, true)
}
func advance(lexer unsafe.Pointer) {
	var v0, v1, v2 unsafe.Pointer
	var lexer_addr, local_advance unsafe.Pointer
	_, _, _, _, _ = lexer_addr, v0, local_advance, v1, v2

	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	local_advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](local_advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, false)
}
func scan_comment(lexer unsafe.Pointer) bool {
	var cmp, cmp2, tobool, cmp8, v18 bool
	var retval unsafe.Pointer
	var result_symbol unsafe.Pointer
	var v1, v4, v7, v9, v10, inc, v11 int32
	var dashes, lookahead, lookahead1, lookahead5, lookahead6 unsafe.Pointer
	var v0, v2, v3, v5, v6, v8, v12, v13, v14, v15, v16, v17 unsafe.Pointer
	var lexer_addr, mark_end unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, dashes, v0, lookahead, v1, cmp, v2, v3, lookahead1, v4, cmp2, v5, v6, lookahead5, v7, tobool, v8, lookahead6, v9, v10, inc, v11, cmp8, v12, result_symbol, v13, v14, mark_end, v15, v16, v17, v18

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	dashes = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v0).F0)
	v1 = *libc.As[int32](lookahead)
	cmp = v1 != 45
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v2)
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v3).F0)
	v4 = *libc.As[int32](lookahead1)
	cmp2 = v4 != 45
	if cmp2 {
		goto if_then3
	} else {
		goto if_end4
	}

if_then3:
	*libc.As[bool](retval) = false
	goto _return

if_end4:
	v5 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v5)
	*libc.As[int32](dashes) = 0
	goto while_cond

while_cond:
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead5 = libc.Ptr(&libc.As[TSLexer](v6).F0)
	v7 = *libc.As[int32](lookahead5)
	tobool = v7 != 0
	if tobool {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead6 = libc.Ptr(&libc.As[TSLexer](v8).F0)
	v9 = *libc.As[int32](lookahead6)
	switch v9 {
	case 45:
		goto sw_bb
	case 62:
		goto sw_bb7
	default:
		goto sw_default
	}

sw_bb:
	v10 = *libc.As[int32](dashes)
	inc = v10 + 1
	*libc.As[int32](dashes) = inc
	goto sw_epilog

sw_bb7:
	v11 = *libc.As[int32](dashes)
	cmp8 = uint32(v11) >= 2
	if cmp8 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v12).F1)
	*libc.As[int16](result_symbol) = 8
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v13)
	v14 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v14).F3)
	v15 = *libc.As[unsafe.Pointer](mark_end)
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v15)(v16)
	*libc.As[bool](retval) = true
	goto _return

if_end10:
	goto sw_default

sw_default:
	*libc.As[int32](dashes) = 0
	goto sw_epilog

sw_epilog:
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v17)
	goto while_cond

while_end:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v18 = *libc.As[bool](retval)
	return v18
}
func scan_implicit_end_tag(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var tag_name unsafe.Pointer
	var next_tag, arrayidx, arrayidx43, arrayidx54 unsafe.Pointer
	var tags, tags1, tags3, tags6, tags7, tags24, tags28, tags31, tags37, tags39, tags47, tags50 unsafe.Pointer
	var cmp, cmp5, cmp10, tobool, call, cmp18, call20, loadedv, cmp26, cmp33, call44, cmp49, cmp56, tobool61, call63, cmp65, cmp68, cmp71, call74, v65 bool
	var retval unsafe.Pointer
	var result_symbol, result_symbol58, result_symbol76 unsafe.Pointer
	var v1, v3, sub, v5, v9, sub9, v11, v22, v32, v34, sub30, v36, v40, sub41, v42, v43, v46, sub52, v47, v48, v51, dec, v55, v57, v59 int32
	var i, size, size2, size4, size8, lookahead, size17, size25, size29, size32, size40, size48, _type, type55, type64, type67, type70 unsafe.Pointer
	var idxprom, v21, v29, idxprom42, idxprom53 int64
	var v20, v28 unsafe.Pointer
	var v30 byte
	var is_closing_tag unsafe.Pointer
	var v0, v2, v4, v6, v7, v8, cond, v10, v12, v13, v14, v15, v16, v17, v19, v23, v24, v25, v27, v31, v33, v35, v37, v38, v39, v41, v44, v45, v49, v50, v52, v53, v54, v56, v58, v60, v61, v62, v63, v64 unsafe.Pointer
	var scanner_addr, lexer_addr, parent, contents, v18, eof, v26, contents38, contents51, eof73 unsafe.Pointer
	var call16 struct {
		F0 unsafe.Pointer
		F1 int64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, parent, is_closing_tag, tag_name, next_tag, i, v0, tags, size, v1, cmp, v2, tags1, size2, v3, sub, v4, tags3, size4, v5, cmp5, v6, tags6, contents, v7, v8, tags7, size8, v9, sub9, idxprom, arrayidx, cond, v10, lookahead, v11, cmp10, v12, v13, tobool, v14, call, v15, v16, result_symbol, v17, call16, v18, v19, v20, v21, size17, v22, cmp18, v23, eof, v24, v25, call20, v26, v27, v28, v29, v30, loadedv, v31, tags24, size25, v32, cmp26, v33, tags28, size29, v34, sub30, v35, tags31, size32, v36, cmp33, v37, tags37, contents38, v38, v39, tags39, size40, v40, sub41, idxprom42, arrayidx43, call44, v41, tags47, size48, v42, v43, cmp49, v44, tags50, contents51, v45, v46, sub52, idxprom53, arrayidx54, _type, v47, type55, v48, cmp56, v49, v50, result_symbol58, v51, dec, v52, tobool61, v53, call63, v54, type64, v55, cmp65, v56, type67, v57, cmp68, v58, type70, v59, cmp71, v60, eof73, v61, v62, call74, v63, v64, result_symbol76, v65

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	parent = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	is_closing_tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	tag_name = libc.Ptr(&new(struct {
		_ [0]uint64
		v String
		b byte
	}).v)
	next_tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v0).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v1 = *libc.As[int32](size)
	cmp = v1 == 0
	if cmp {
		goto cond_true
	} else {
		goto cond_false
	}

cond_true:
	cond = nil
	goto cond_end

cond_false:
	v2 = *libc.As[unsafe.Pointer](scanner_addr)
	tags1 = libc.Ptr(&libc.As[Scanner](v2).F0)
	size2 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F1)
	v3 = *libc.As[int32](size2)
	sub = v3 - 1
	v4 = *libc.As[unsafe.Pointer](scanner_addr)
	tags3 = libc.Ptr(&libc.As[Scanner](v4).F0)
	size4 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags3).F1)
	v5 = *libc.As[int32](size4)
	cmp5 = uint32(sub) < uint32(v5)
	if cmp5 {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	goto if_end

if_else:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 343, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_implicit_end_tag)))
	panic("unreachable")

if_end:
	v6 = *libc.As[unsafe.Pointer](scanner_addr)
	tags6 = libc.Ptr(&libc.As[Scanner](v6).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags6).F0)
	v7 = *libc.As[unsafe.Pointer](contents)
	v8 = *libc.As[unsafe.Pointer](scanner_addr)
	tags7 = libc.Ptr(&libc.As[Scanner](v8).F0)
	size8 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags7).F1)
	v9 = *libc.As[int32](size8)
	sub9 = v9 - 1
	idxprom = int64(uint64(uint32(sub9)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v7), int(idxprom)*24))
	cond = arrayidx
	goto cond_end

cond_end:
	*libc.As[unsafe.Pointer](parent) = cond
	*libc.As[byte](is_closing_tag) = 0
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v10).F0)
	v11 = *libc.As[int32](lookahead)
	cmp10 = v11 == 47
	if cmp10 {
		goto if_then11
	} else {
		goto if_else12
	}

if_then11:
	*libc.As[byte](is_closing_tag) = 1
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v12)
	goto if_end15

if_else12:
	v13 = *libc.As[unsafe.Pointer](parent)
	tobool = uintptr(unsafe.Pointer(v13)) != uintptr(unsafe.Pointer(nil))
	if tobool {
		goto land_lhs_true
	} else {
		goto if_end14
	}

land_lhs_true:
	v14 = *libc.As[unsafe.Pointer](parent)
	call = tag_is_void(v14)
	if call {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	v15 = *libc.As[unsafe.Pointer](scanner_addr)
	pop_tag(v15)
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v16).F1)
	*libc.As[int16](result_symbol) = 6
	*libc.As[bool](retval) = true
	goto _return

if_end14:
	goto if_end15

if_end15:
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	call16 = scan_tag_name(v17)
	v18 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v19 = call16.F0
	*libc.As[unsafe.Pointer](v18) = v19
	v20 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v21 = call16.F1
	*libc.As[int64](v20) = v21
	size17 = libc.Ptr(&libc.As[String](tag_name).F1)
	v22 = *libc.As[int32](size17)
	cmp18 = v22 == 0
	if cmp18 {
		goto land_lhs_true19
	} else {
		goto if_end22
	}

land_lhs_true19:
	v23 = *libc.As[unsafe.Pointer](lexer_addr)
	eof = libc.Ptr(&libc.As[TSLexer](v23).F6)
	v24 = *libc.As[unsafe.Pointer](eof)
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	call20 = libc.FuncFromCode[func(unsafe.Pointer) bool](v24)(v25)
	if call20 {
		goto if_end22
	} else {
		goto if_then21
	}

if_then21:
	_array__delete(tag_name)
	*libc.As[bool](retval) = false
	goto _return

if_end22:
	v26 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v27 = *libc.As[unsafe.Pointer](v26)
	v28 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v29 = *libc.As[int64](v28)
	tag_for_name(next_tag, v27, v29)
	v30 = *libc.As[byte](is_closing_tag)
	loadedv = (v30 & 1) != 0
	if loadedv {
		goto if_then23
	} else {
		goto if_else60
	}

if_then23:
	v31 = *libc.As[unsafe.Pointer](scanner_addr)
	tags24 = libc.Ptr(&libc.As[Scanner](v31).F0)
	size25 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags24).F1)
	v32 = *libc.As[int32](size25)
	cmp26 = uint32(v32) > 0
	if cmp26 {
		goto land_lhs_true27
	} else {
		goto if_end46
	}

land_lhs_true27:
	v33 = *libc.As[unsafe.Pointer](scanner_addr)
	tags28 = libc.Ptr(&libc.As[Scanner](v33).F0)
	size29 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags28).F1)
	v34 = *libc.As[int32](size29)
	sub30 = v34 - 1
	v35 = *libc.As[unsafe.Pointer](scanner_addr)
	tags31 = libc.Ptr(&libc.As[Scanner](v35).F0)
	size32 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags31).F1)
	v36 = *libc.As[int32](size32)
	cmp33 = uint32(sub30) < uint32(v36)
	if cmp33 {
		goto if_then34
	} else {
		goto if_else35
	}

if_then34:
	goto if_end36

if_else35:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 367, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_implicit_end_tag)))
	panic("unreachable")

if_end36:
	v37 = *libc.As[unsafe.Pointer](scanner_addr)
	tags37 = libc.Ptr(&libc.As[Scanner](v37).F0)
	contents38 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags37).F0)
	v38 = *libc.As[unsafe.Pointer](contents38)
	v39 = *libc.As[unsafe.Pointer](scanner_addr)
	tags39 = libc.Ptr(&libc.As[Scanner](v39).F0)
	size40 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags39).F1)
	v40 = *libc.As[int32](size40)
	sub41 = v40 - 1
	idxprom42 = int64(uint64(uint32(sub41)))
	arrayidx43 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v38), int(idxprom42)*24))
	call44 = tag_eq(arrayidx43, next_tag)
	if call44 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	tag_free(next_tag)
	*libc.As[bool](retval) = false
	goto _return

if_end46:
	v41 = *libc.As[unsafe.Pointer](scanner_addr)
	tags47 = libc.Ptr(&libc.As[Scanner](v41).F0)
	size48 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags47).F1)
	v42 = *libc.As[int32](size48)
	*libc.As[int32](i) = v42
	goto for_cond

for_cond:
	v43 = *libc.As[int32](i)
	cmp49 = uint32(v43) > 0
	if cmp49 {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v44 = *libc.As[unsafe.Pointer](scanner_addr)
	tags50 = libc.Ptr(&libc.As[Scanner](v44).F0)
	contents51 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags50).F0)
	v45 = *libc.As[unsafe.Pointer](contents51)
	v46 = *libc.As[int32](i)
	sub52 = v46 - 1
	idxprom53 = int64(uint64(uint32(sub52)))
	arrayidx54 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v45), int(idxprom53)*24))
	_type = libc.Ptr(&libc.As[Tag](arrayidx54).F0)
	v47 = *libc.As[int32](_type)
	type55 = libc.Ptr(&libc.As[Tag](next_tag).F0)
	v48 = *libc.As[int32](type55)
	cmp56 = v47 == v48
	if cmp56 {
		goto if_then57
	} else {
		goto if_end59
	}

if_then57:
	v49 = *libc.As[unsafe.Pointer](scanner_addr)
	pop_tag(v49)
	v50 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol58 = libc.Ptr(&libc.As[TSLexer](v50).F1)
	*libc.As[int16](result_symbol58) = 6
	tag_free(next_tag)
	*libc.As[bool](retval) = true
	goto _return

if_end59:
	goto for_inc

for_inc:
	v51 = *libc.As[int32](i)
	dec = v51 - 1
	*libc.As[int32](i) = dec
	goto for_cond

for_end:
	goto if_end78

if_else60:
	v52 = *libc.As[unsafe.Pointer](parent)
	tobool61 = uintptr(unsafe.Pointer(v52)) != uintptr(unsafe.Pointer(nil))
	if tobool61 {
		goto land_lhs_true62
	} else {
		goto if_end77
	}

land_lhs_true62:
	v53 = *libc.As[unsafe.Pointer](parent)
	call63 = tag_can_contain(v53, next_tag)
	if call63 {
		goto lor_lhs_false
	} else {
		goto if_then75
	}

lor_lhs_false:
	v54 = *libc.As[unsafe.Pointer](parent)
	type64 = libc.Ptr(&libc.As[Tag](v54).F0)
	v55 = *libc.As[int32](type64)
	cmp65 = v55 == 64
	if cmp65 {
		goto land_lhs_true72
	} else {
		goto lor_lhs_false66
	}

lor_lhs_false66:
	v56 = *libc.As[unsafe.Pointer](parent)
	type67 = libc.Ptr(&libc.As[Tag](v56).F0)
	v57 = *libc.As[int32](type67)
	cmp68 = v57 == 61
	if cmp68 {
		goto land_lhs_true72
	} else {
		goto lor_lhs_false69
	}

lor_lhs_false69:
	v58 = *libc.As[unsafe.Pointer](parent)
	type70 = libc.Ptr(&libc.As[Tag](v58).F0)
	v59 = *libc.As[int32](type70)
	cmp71 = v59 == 32
	if cmp71 {
		goto land_lhs_true72
	} else {
		goto if_end77
	}

land_lhs_true72:
	v60 = *libc.As[unsafe.Pointer](lexer_addr)
	eof73 = libc.Ptr(&libc.As[TSLexer](v60).F6)
	v61 = *libc.As[unsafe.Pointer](eof73)
	v62 = *libc.As[unsafe.Pointer](lexer_addr)
	call74 = libc.FuncFromCode[func(unsafe.Pointer) bool](v61)(v62)
	if call74 {
		goto if_then75
	} else {
		goto if_end77
	}

if_then75:
	v63 = *libc.As[unsafe.Pointer](scanner_addr)
	pop_tag(v63)
	v64 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol76 = libc.Ptr(&libc.As[TSLexer](v64).F1)
	*libc.As[int16](result_symbol76) = 6
	tag_free(next_tag)
	*libc.As[bool](retval) = true
	goto _return

if_end77:
	goto if_end78

if_end78:
	tag_free(next_tag)
	*libc.As[bool](retval) = false
	goto _return

_return:
	v65 = *libc.As[bool](retval)
	return v65
}
func scan_self_closing_tag_delimiter(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var tags unsafe.Pointer
	var cmp, cmp1, v8 bool
	var retval unsafe.Pointer
	var result_symbol unsafe.Pointer
	var v2, v5 int32
	var lookahead, size unsafe.Pointer
	var v0, v1, v3, v4, v6, v7 unsafe.Pointer
	var scanner_addr, lexer_addr unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, v0, v1, lookahead, v2, cmp, v3, v4, tags, size, v5, cmp1, v6, v7, result_symbol, v8

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v0)
	v1 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v1).F0)
	v2 = *libc.As[int32](lookahead)
	cmp = v2 == 62
	if cmp {
		goto if_then
	} else {
		goto if_end3
	}

if_then:
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v3)
	v4 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v4).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v5 = *libc.As[int32](size)
	cmp1 = uint32(v5) > 0
	if cmp1 {
		goto if_then2
	} else {
		goto if_end
	}

if_then2:
	v6 = *libc.As[unsafe.Pointer](scanner_addr)
	pop_tag(v6)
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v7).F1)
	*libc.As[int16](result_symbol) = 5
	goto if_end

if_end:
	*libc.As[bool](retval) = true
	goto _return

if_end3:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v8 = *libc.As[bool](retval)
	return v8
}
func _array__grow(self unsafe.Pointer, count int32, element_size int64) {
	var cmp, cmp2, cmp4 bool
	var v1, v2, add, v3, v5, v7, mul, v8, v9, v10, v11, v14 int32
	var count_addr, new_size, new_capacity, size, capacity, capacity1 unsafe.Pointer
	var v13 int64
	var element_size_addr unsafe.Pointer
	var v0, v4, v6, v12 unsafe.Pointer
	var self_addr unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = self_addr, count_addr, element_size_addr, new_size, new_capacity, v0, size, v1, v2, add, v3, v4, capacity, v5, cmp, v6, capacity1, v7, mul, v8, cmp2, v9, v10, cmp4, v11, v12, v13, v14

	self_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	count_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	element_size_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int64
		b byte
	}).v)
	new_size = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	new_capacity = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](self_addr) = self
	*libc.As[int32](count_addr) = count
	*libc.As[int64](element_size_addr) = element_size
	v0 = *libc.As[unsafe.Pointer](self_addr)
	size = libc.Ptr(&libc.As[Array](v0).F1)
	v1 = *libc.As[int32](size)
	v2 = *libc.As[int32](count_addr)
	add = v1 + v2
	*libc.As[int32](new_size) = add
	v3 = *libc.As[int32](new_size)
	v4 = *libc.As[unsafe.Pointer](self_addr)
	capacity = libc.Ptr(&libc.As[Array](v4).F2)
	v5 = *libc.As[int32](capacity)
	cmp = uint32(v3) > uint32(v5)
	if cmp {
		goto if_then
	} else {
		goto if_end7
	}

if_then:
	v6 = *libc.As[unsafe.Pointer](self_addr)
	capacity1 = libc.Ptr(&libc.As[Array](v6).F2)
	v7 = *libc.As[int32](capacity1)
	mul = v7 * 2
	*libc.As[int32](new_capacity) = mul
	v8 = *libc.As[int32](new_capacity)
	cmp2 = uint32(v8) < 8
	if cmp2 {
		goto if_then3
	} else {
		goto if_end
	}

if_then3:
	*libc.As[int32](new_capacity) = 8
	goto if_end

if_end:
	v9 = *libc.As[int32](new_capacity)
	v10 = *libc.As[int32](new_size)
	cmp4 = uint32(v9) < uint32(v10)
	if cmp4 {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	v11 = *libc.As[int32](new_size)
	*libc.As[int32](new_capacity) = v11
	goto if_end6

if_end6:
	v12 = *libc.As[unsafe.Pointer](self_addr)
	v13 = *libc.As[int64](element_size_addr)
	v14 = *libc.As[int32](new_capacity)
	_array__reserve(v12, v13, v14)
	goto if_end7

if_end7:
}
func scan_js_backtick_string(lexer unsafe.Pointer) {
	var cmp, cmp2, cmp5, cmp10 bool
	var v4, v6, v11, v17 int32
	var lookahead, lookahead1, lookahead4, lookahead9 unsafe.Pointer
	var v0, v1, v2, v3, v5, v7, v8, v9, v10, v12, v13, v14, v15, v16, v18, v19, v20, v21, v22, v23 unsafe.Pointer
	var lexer_addr, local_advance, advance3, advance7, advance12, advance15 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = lexer_addr, v0, local_advance, v1, v2, v3, lookahead, v4, cmp, v5, lookahead1, v6, cmp2, v7, advance3, v8, v9, v10, lookahead4, v11, cmp5, v12, advance7, v13, v14, v15, v16, lookahead9, v17, cmp10, v18, advance12, v19, v20, v21, advance15, v22, v23

	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	local_advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](local_advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, false)
	goto while_cond

while_cond:
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v3).F0)
	v4 = *libc.As[int32](lookahead)
	cmp = v4 != 0
	if cmp {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v5 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v5).F0)
	v6 = *libc.As[int32](lookahead1)
	cmp2 = v6 == 36
	if cmp2 {
		goto if_then
	} else {
		goto if_else8
	}

if_then:
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	advance3 = libc.Ptr(&libc.As[TSLexer](v7).F2)
	v8 = *libc.As[unsafe.Pointer](advance3)
	v9 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v8)(v9, false)
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead4 = libc.Ptr(&libc.As[TSLexer](v10).F0)
	v11 = *libc.As[int32](lookahead4)
	cmp5 = v11 == 123
	if cmp5 {
		goto if_then6
	} else {
		goto if_else
	}

if_then6:
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	advance7 = libc.Ptr(&libc.As[TSLexer](v12).F2)
	v13 = *libc.As[unsafe.Pointer](advance7)
	v14 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v13)(v14, false)
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	scan_js_expr_with_delimiter(v15, 1)
	goto if_end

if_else:
	goto while_cond

if_end:
	goto if_end14

if_else8:
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead9 = libc.Ptr(&libc.As[TSLexer](v16).F0)
	v17 = *libc.As[int32](lookahead9)
	cmp10 = v17 == 96
	if cmp10 {
		goto if_then11
	} else {
		goto if_end13
	}

if_then11:
	v18 = *libc.As[unsafe.Pointer](lexer_addr)
	advance12 = libc.Ptr(&libc.As[TSLexer](v18).F2)
	v19 = *libc.As[unsafe.Pointer](advance12)
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v19)(v20, false)
	goto while_end

if_end13:
	goto if_end14

if_end14:
	v21 = *libc.As[unsafe.Pointer](lexer_addr)
	advance15 = libc.Ptr(&libc.As[TSLexer](v21).F2)
	v22 = *libc.As[unsafe.Pointer](advance15)
	v23 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v22)(v23, false)
	goto while_cond

while_end:
}
func scan_start_tag_name(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var tag_name unsafe.Pointer
	var tag, tag6, arrayidx, arrayidx14 unsafe.Pointer
	var tags, tags3, tags4, tags7, tags8, tags10 unsafe.Pointer
	var cmp, cmp1, v28 bool
	var retval unsafe.Pointer
	var result_symbol, result_symbol15, result_symbol17, result_symbol18 unsafe.Pointer
	var v5, v7, v13, inc, v23, inc12, v24 int32
	var size, lookahead, size5, size11, _type unsafe.Pointer
	var v4, idxprom, v18, idxprom13 int64
	var v3, v17 unsafe.Pointer
	var v0, v2, v6, v8, v9, v10, v11, v12, v14, v16, v19, v20, v21, v22, v25, v26, v27 unsafe.Pointer
	var scanner_addr, lexer_addr, v1, contents, v15, contents9 unsafe.Pointer
	var call struct {
		F0 unsafe.Pointer
		F1 int64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, tag_name, tag, tag6, v0, call, v1, v2, v3, v4, size, v5, cmp, v6, lookahead, v7, cmp1, v8, v9, tags, v10, tags3, contents, v11, v12, tags4, size5, v13, inc, idxprom, arrayidx, v14, result_symbol, v15, v16, v17, v18, v19, tags7, v20, tags8, contents9, v21, v22, tags10, size11, v23, inc12, idxprom13, arrayidx14, _type, v24, v25, result_symbol15, v26, result_symbol17, v27, result_symbol18, v28

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	tag_name = libc.Ptr(&new(struct {
		_ [0]uint64
		v String
		b byte
	}).v)
	tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	tag6 = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	call = scan_tag_name(v0)
	v1 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v2 = call.F0
	*libc.As[unsafe.Pointer](v1) = v2
	v3 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v4 = call.F1
	*libc.As[int64](v3) = v4
	size = libc.Ptr(&libc.As[String](tag_name).F1)
	v5 = *libc.As[int32](size)
	cmp = v5 == 0
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	_array__delete(tag_name)
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v6).F0)
	v7 = *libc.As[int32](lookahead)
	cmp1 = v7 == 62
	if cmp1 {
		goto if_then2
	} else {
		goto if_else
	}

if_then2:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v8)
	libc.Memmove(libc.As[byte](tag), libc.As[byte](libc.Ptr(&__const_scan_start_tag_name_tag)), int64(24))
	v9 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v9).F0)
	_array__grow(tags, 1, int64(24))
	v10 = *libc.As[unsafe.Pointer](scanner_addr)
	tags3 = libc.Ptr(&libc.As[Scanner](v10).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags3).F0)
	v11 = *libc.As[unsafe.Pointer](contents)
	v12 = *libc.As[unsafe.Pointer](scanner_addr)
	tags4 = libc.Ptr(&libc.As[Scanner](v12).F0)
	size5 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags4).F1)
	v13 = *libc.As[int32](size5)
	inc = v13 + 1
	*libc.As[int32](size5) = inc
	idxprom = int64(uint64(uint32(v13)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v11), int(idxprom)*24))
	libc.Memmove(libc.As[byte](arrayidx), libc.As[byte](tag), int64(24))
	v14 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v14).F1)
	*libc.As[int16](result_symbol) = 15
	*libc.As[bool](retval) = true
	goto _return

if_else:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v15 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v16 = *libc.As[unsafe.Pointer](v15)
	v17 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v18 = *libc.As[int64](v17)
	tag_for_name(tag6, v16, v18)
	v19 = *libc.As[unsafe.Pointer](scanner_addr)
	tags7 = libc.Ptr(&libc.As[Scanner](v19).F0)
	_array__grow(tags7, 1, int64(24))
	v20 = *libc.As[unsafe.Pointer](scanner_addr)
	tags8 = libc.Ptr(&libc.As[Scanner](v20).F0)
	contents9 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags8).F0)
	v21 = *libc.As[unsafe.Pointer](contents9)
	v22 = *libc.As[unsafe.Pointer](scanner_addr)
	tags10 = libc.Ptr(&libc.As[Scanner](v22).F0)
	size11 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags10).F1)
	v23 = *libc.As[int32](size11)
	inc12 = v23 + 1
	*libc.As[int32](size11) = inc12
	idxprom13 = int64(uint64(uint32(v23)))
	arrayidx14 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v21), int(idxprom13)*24))
	libc.Memmove(libc.As[byte](arrayidx14), libc.As[byte](tag6), int64(24))
	_type = libc.Ptr(&libc.As[Tag](tag6).F0)
	v24 = *libc.As[int32](_type)
	switch v24 {
	case 97:
		goto sw_bb
	case 104:
		goto sw_bb16
	default:
		goto sw_default
	}

sw_bb:
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol15 = libc.Ptr(&libc.As[TSLexer](v25).F1)
	*libc.As[int16](result_symbol15) = 1
	goto sw_epilog

sw_bb16:
	v26 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol17 = libc.Ptr(&libc.As[TSLexer](v26).F1)
	*libc.As[int16](result_symbol17) = 2
	goto sw_epilog

sw_default:
	v27 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol18 = libc.Ptr(&libc.As[TSLexer](v27).F1)
	*libc.As[int16](result_symbol18) = 0
	goto sw_epilog

sw_epilog:
	*libc.As[bool](retval) = true
	goto _return

_return:
	v28 = *libc.As[bool](retval)
	return v28
}
func scan_end_tag_name(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var tag_name unsafe.Pointer
	var tag, arrayidx, arrayidx40 unsafe.Pointer
	var tags, tags5, tags7, tags11, tags12, tags21, tags25, tags28, tags34, tags36 unsafe.Pointer
	var cmp, cmp1, cmp4, cmp9, cmp15, cmp23, cmp30, call41, v40 bool
	var retval unsafe.Pointer
	var result_symbol, result_symbol18, result_symbol43, result_symbol45 unsafe.Pointer
	var v5, v7, v10, v12, sub, v14, v18, sub14, v19, v28, v30, sub27, v32, v36, sub38 int32
	var size, lookahead, size3, size6, size8, size13, _type, size22, size26, size29, size37 unsafe.Pointer
	var v4, idxprom, v26, idxprom39 int64
	var v3, v25 unsafe.Pointer
	var v0, v2, v6, v8, v9, v11, v13, v15, v16, v17, v20, v21, v22, v24, v27, v29, v31, v33, v34, v35, v37, v38, v39 unsafe.Pointer
	var scanner_addr, lexer_addr, v1, contents, v23, contents35 unsafe.Pointer
	var call struct {
		F0 unsafe.Pointer
		F1 int64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, tag_name, tag, v0, call, v1, v2, v3, v4, size, v5, cmp, v6, lookahead, v7, cmp1, v8, v9, tags, size3, v10, cmp4, v11, tags5, size6, v12, sub, v13, tags7, size8, v14, cmp9, v15, tags11, contents, v16, v17, tags12, size13, v18, sub14, idxprom, arrayidx, _type, v19, cmp15, v20, v21, result_symbol, v22, result_symbol18, v23, v24, v25, v26, v27, tags21, size22, v28, cmp23, v29, tags25, size26, v30, sub27, v31, tags28, size29, v32, cmp30, v33, tags34, contents35, v34, v35, tags36, size37, v36, sub38, idxprom39, arrayidx40, call41, v37, v38, result_symbol43, v39, result_symbol45, v40

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	tag_name = libc.Ptr(&new(struct {
		_ [0]uint64
		v String
		b byte
	}).v)
	tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	call = scan_tag_name(v0)
	v1 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v2 = call.F0
	*libc.As[unsafe.Pointer](v1) = v2
	v3 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v4 = call.F1
	*libc.As[int64](v3) = v4
	size = libc.Ptr(&libc.As[String](tag_name).F1)
	v5 = *libc.As[int32](size)
	cmp = v5 == 0
	if cmp {
		goto if_then
	} else {
		goto if_end20
	}

if_then:
	_array__delete(tag_name)
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v6).F0)
	v7 = *libc.As[int32](lookahead)
	cmp1 = v7 == 62
	if cmp1 {
		goto if_then2
	} else {
		goto if_else19
	}

if_then2:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v8)
	v9 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v9).F0)
	size3 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v10 = *libc.As[int32](size3)
	cmp4 = uint32(v10) > 0
	if cmp4 {
		goto land_lhs_true
	} else {
		goto if_else17
	}

land_lhs_true:
	v11 = *libc.As[unsafe.Pointer](scanner_addr)
	tags5 = libc.Ptr(&libc.As[Scanner](v11).F0)
	size6 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags5).F1)
	v12 = *libc.As[int32](size6)
	sub = v12 - 1
	v13 = *libc.As[unsafe.Pointer](scanner_addr)
	tags7 = libc.Ptr(&libc.As[Scanner](v13).F0)
	size8 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags7).F1)
	v14 = *libc.As[int32](size8)
	cmp9 = uint32(sub) < uint32(v14)
	if cmp9 {
		goto if_then10
	} else {
		goto if_else
	}

if_then10:
	goto if_end

if_else:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 437, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_end_tag_name)))
	panic("unreachable")

if_end:
	v15 = *libc.As[unsafe.Pointer](scanner_addr)
	tags11 = libc.Ptr(&libc.As[Scanner](v15).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags11).F0)
	v16 = *libc.As[unsafe.Pointer](contents)
	v17 = *libc.As[unsafe.Pointer](scanner_addr)
	tags12 = libc.Ptr(&libc.As[Scanner](v17).F0)
	size13 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags12).F1)
	v18 = *libc.As[int32](size13)
	sub14 = v18 - 1
	idxprom = int64(uint64(uint32(sub14)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v16), int(idxprom)*24))
	_type = libc.Ptr(&libc.As[Tag](arrayidx).F0)
	v19 = *libc.As[int32](_type)
	cmp15 = v19 == 125
	if cmp15 {
		goto if_then16
	} else {
		goto if_else17
	}

if_then16:
	v20 = *libc.As[unsafe.Pointer](scanner_addr)
	pop_tag(v20)
	v21 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v21).F1)
	*libc.As[int16](result_symbol) = 15
	*libc.As[bool](retval) = true
	goto _return

if_else17:
	v22 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol18 = libc.Ptr(&libc.As[TSLexer](v22).F1)
	*libc.As[int16](result_symbol18) = 4
	*libc.As[bool](retval) = true
	goto _return

if_else19:
	*libc.As[bool](retval) = false
	goto _return

if_end20:
	v23 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v24 = *libc.As[unsafe.Pointer](v23)
	v25 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v26 = *libc.As[int64](v25)
	tag_for_name(tag, v24, v26)
	v27 = *libc.As[unsafe.Pointer](scanner_addr)
	tags21 = libc.Ptr(&libc.As[Scanner](v27).F0)
	size22 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags21).F1)
	v28 = *libc.As[int32](size22)
	cmp23 = uint32(v28) > 0
	if cmp23 {
		goto land_lhs_true24
	} else {
		goto if_else44
	}

land_lhs_true24:
	v29 = *libc.As[unsafe.Pointer](scanner_addr)
	tags25 = libc.Ptr(&libc.As[Scanner](v29).F0)
	size26 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags25).F1)
	v30 = *libc.As[int32](size26)
	sub27 = v30 - 1
	v31 = *libc.As[unsafe.Pointer](scanner_addr)
	tags28 = libc.Ptr(&libc.As[Scanner](v31).F0)
	size29 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags28).F1)
	v32 = *libc.As[int32](size29)
	cmp30 = uint32(sub27) < uint32(v32)
	if cmp30 {
		goto if_then31
	} else {
		goto if_else32
	}

if_then31:
	goto if_end33

if_else32:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 452, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_end_tag_name)))
	panic("unreachable")

if_end33:
	v33 = *libc.As[unsafe.Pointer](scanner_addr)
	tags34 = libc.Ptr(&libc.As[Scanner](v33).F0)
	contents35 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags34).F0)
	v34 = *libc.As[unsafe.Pointer](contents35)
	v35 = *libc.As[unsafe.Pointer](scanner_addr)
	tags36 = libc.Ptr(&libc.As[Scanner](v35).F0)
	size37 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags36).F1)
	v36 = *libc.As[int32](size37)
	sub38 = v36 - 1
	idxprom39 = int64(uint64(uint32(sub38)))
	arrayidx40 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v34), int(idxprom39)*24))
	call41 = tag_eq(arrayidx40, tag)
	if call41 {
		goto if_then42
	} else {
		goto if_else44
	}

if_then42:
	v37 = *libc.As[unsafe.Pointer](scanner_addr)
	pop_tag(v37)
	v38 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol43 = libc.Ptr(&libc.As[TSLexer](v38).F1)
	*libc.As[int16](result_symbol43) = 3
	goto if_end46

if_else44:
	v39 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol45 = libc.Ptr(&libc.As[TSLexer](v39).F1)
	*libc.As[int16](result_symbol45) = 4
	goto if_end46

if_end46:
	tag_free(tag)
	*libc.As[bool](retval) = true
	goto _return

_return:
	v40 = *libc.As[bool](retval)
	return v40
}
func scan_js_string(lexer unsafe.Pointer) {
	var cmp, cmp3, cmp6, cmp13 bool
	var v1, v4, v9, v11, v16, conv12 int32
	var lookahead, lookahead1, lookahead2, lookahead5, lookahead11 unsafe.Pointer
	var conv, v17 byte
	var str_end_char unsafe.Pointer
	var v0, v2, v3, v5, v6, v7, v8, v10, v12, v13, v14, v15, v18, v19, v20, v21, v22, v23 unsafe.Pointer
	var lexer_addr, local_advance, advance9, advance16, advance18 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = lexer_addr, str_end_char, v0, lookahead, v1, cmp, v2, v3, lookahead1, v4, conv, v5, local_advance, v6, v7, v8, lookahead2, v9, cmp3, v10, lookahead5, v11, cmp6, v12, advance9, v13, v14, v15, lookahead11, v16, v17, conv12, cmp13, v18, advance16, v19, v20, v21, advance18, v22, v23

	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	str_end_char = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v0).F0)
	v1 = *libc.As[int32](lookahead)
	cmp = v1 == 96
	if cmp {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	scan_js_backtick_string(v2)
	goto if_end19

if_else:
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v3).F0)
	v4 = *libc.As[int32](lookahead1)
	conv = byte(v4)
	*libc.As[byte](str_end_char) = conv
	v5 = *libc.As[unsafe.Pointer](lexer_addr)
	local_advance = libc.Ptr(&libc.As[TSLexer](v5).F2)
	v6 = *libc.As[unsafe.Pointer](local_advance)
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v6)(v7, false)
	goto while_cond

while_cond:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead2 = libc.Ptr(&libc.As[TSLexer](v8).F0)
	v9 = *libc.As[int32](lookahead2)
	cmp3 = v9 != 0
	if cmp3 {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead5 = libc.Ptr(&libc.As[TSLexer](v10).F0)
	v11 = *libc.As[int32](lookahead5)
	cmp6 = v11 == 92
	if cmp6 {
		goto if_then8
	} else {
		goto if_else10
	}

if_then8:
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	advance9 = libc.Ptr(&libc.As[TSLexer](v12).F2)
	v13 = *libc.As[unsafe.Pointer](advance9)
	v14 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v13)(v14, false)
	goto if_end17

if_else10:
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead11 = libc.Ptr(&libc.As[TSLexer](v15).F0)
	v16 = *libc.As[int32](lookahead11)
	v17 = *libc.As[byte](str_end_char)
	conv12 = int32(int8(v17))
	cmp13 = v16 == conv12
	if cmp13 {
		goto if_then15
	} else {
		goto if_end
	}

if_then15:
	v18 = *libc.As[unsafe.Pointer](lexer_addr)
	advance16 = libc.Ptr(&libc.As[TSLexer](v18).F2)
	v19 = *libc.As[unsafe.Pointer](advance16)
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v19)(v20, false)
	goto if_end19

if_end:
	goto if_end17

if_end17:
	v21 = *libc.As[unsafe.Pointer](lexer_addr)
	advance18 = libc.Ptr(&libc.As[TSLexer](v21).F2)
	v22 = *libc.As[unsafe.Pointer](advance18)
	v23 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v22)(v23, false)
	goto while_cond

while_end:
	goto if_end19

if_end19:
}
func tag_is_void(self unsafe.Pointer) bool {
	var cmp bool
	var v1 int32
	var _type unsafe.Pointer
	var v0 unsafe.Pointer
	var self_addr unsafe.Pointer
	_, _, _, _, _ = self_addr, v0, _type, v1, cmp

	self_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](self_addr) = self
	v0 = *libc.As[unsafe.Pointer](self_addr)
	_type = libc.Ptr(&libc.As[Tag](v0).F0)
	v1 = *libc.As[int32](_type)
	cmp = uint32(v1) < 21
	return cmp
}
func pop_tag(scanner unsafe.Pointer) {
	var popped_tag, arrayidx unsafe.Pointer
	var tags, tags1 unsafe.Pointer
	var v3, dec int32
	var size unsafe.Pointer
	var idxprom int64
	var v0, v1, v2 unsafe.Pointer
	var scanner_addr, contents unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _ = scanner_addr, popped_tag, v0, tags, contents, v1, v2, tags1, size, v3, dec, idxprom, arrayidx

	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	popped_tag = libc.Ptr(&new(struct {
		_ [0]uint64
		v Tag
		b byte
	}).v)
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	v0 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v0).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F0)
	v1 = *libc.As[unsafe.Pointer](contents)
	v2 = *libc.As[unsafe.Pointer](scanner_addr)
	tags1 = libc.Ptr(&libc.As[Scanner](v2).F0)
	size = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F1)
	v3 = *libc.As[int32](size)
	dec = v3 - 1
	*libc.As[int32](size) = dec
	idxprom = int64(uint64(uint32(dec)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v1), int(idxprom)*24))
	libc.Memmove(libc.As[byte](popped_tag), libc.As[byte](arrayidx), int64(24))
	tag_free(popped_tag)
}
func scan_tag_name(lexer unsafe.Pointer) struct {
	F0 unsafe.Pointer
	F1 int64
} {
	var retval unsafe.Pointer
	var tobool, cmp, cmp4, cmp6, v8 bool
	var v1, call, v3, v5, v7, v10, v12, inc int32
	var lookahead, lookahead1, lookahead3, lookahead5, lookahead7, size unsafe.Pointer
	var idxprom int64
	var conv byte
	var arrayidx unsafe.Pointer
	var v0, v2, v4, v6, v9, v11, v13 unsafe.Pointer
	var lexer_addr, contents unsafe.Pointer
	var v14 struct {
		F0 unsafe.Pointer
		F1 int64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, v0, lookahead, v1, call, tobool, v2, lookahead1, v3, cmp, v4, lookahead3, v5, cmp4, v6, lookahead5, v7, cmp6, v8, v9, lookahead7, v10, conv, contents, v11, size, v12, inc, idxprom, arrayidx, v13, v14

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v String
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	libc.Memset(libc.As[byte](retval), 0, int64(16))
	goto while_cond

while_cond:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v0).F0)
	v1 = *libc.As[int32](lookahead)
	call = libc.Iswalnum(v1)
	tobool = call != 0
	if tobool {
		v8 = true
		goto lor_end
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v2).F0)
	v3 = *libc.As[int32](lookahead1)
	cmp = v3 == 45
	if cmp {
		v8 = true
		goto lor_end
	} else {
		goto lor_lhs_false2
	}

lor_lhs_false2:
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead3 = libc.Ptr(&libc.As[TSLexer](v4).F0)
	v5 = *libc.As[int32](lookahead3)
	cmp4 = v5 == 58
	if cmp4 {
		v8 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead5 = libc.Ptr(&libc.As[TSLexer](v6).F0)
	v7 = *libc.As[int32](lookahead5)
	cmp6 = v7 == 46
	v8 = cmp6
	goto lor_end

lor_end:
	if v8 {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	_array__grow(retval, 1, int64(1))
	v9 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead7 = libc.Ptr(&libc.As[TSLexer](v9).F0)
	v10 = *libc.As[int32](lookahead7)
	conv = byte(v10)
	contents = libc.Ptr(&libc.As[String](retval).F0)
	v11 = *libc.As[unsafe.Pointer](contents)
	size = libc.Ptr(&libc.As[String](retval).F1)
	v12 = *libc.As[int32](size)
	inc = v12 + 1
	*libc.As[int32](size) = inc
	idxprom = int64(uint64(uint32(v12)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v11), int(idxprom)*1))
	*libc.As[byte](arrayidx) = conv
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v13)
	goto while_cond

while_end:
	v14 = *libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](retval)
	return v14
}
func tag_for_name(agg_result unsafe.Pointer, name_coerce0 unsafe.Pointer, name_coerce1 int64) {
	var name, custom_tag_name unsafe.Pointer
	var cmp bool
	var call, v2 int32
	var _type, type1 unsafe.Pointer
	var v1 unsafe.Pointer
	var v0 unsafe.Pointer
	_, _, _, _, _, _, _, _, _ = name, v0, v1, call, _type, type1, v2, cmp, custom_tag_name

	name = libc.Ptr(&new(struct {
		_ [0]uint64
		v String
		b byte
	}).v)
	v0 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](name).F0)
	*libc.As[unsafe.Pointer](v0) = name_coerce0
	v1 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](name).F1)
	*libc.As[int64](v1) = name_coerce1
	tag_new(agg_result)
	call = tag_type_for_name(name)
	_type = libc.Ptr(&libc.As[Tag](agg_result).F0)
	*libc.As[int32](_type) = call
	type1 = libc.Ptr(&libc.As[Tag](agg_result).F0)
	v2 = *libc.As[int32](type1)
	cmp = v2 == 126
	if cmp {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	custom_tag_name = libc.Ptr(&libc.As[Tag](agg_result).F1)
	libc.Memmove(libc.As[byte](custom_tag_name), libc.As[byte](name), int64(16))
	goto if_end

if_else:
	_array__delete(name)
	goto if_end

if_end:
}
func tag_eq(self unsafe.Pointer, other unsafe.Pointer) bool {
	var custom_tag_name, custom_tag_name5, custom_tag_name10, custom_tag_name11, custom_tag_name13 unsafe.Pointer
	var cmp, cmp3, cmp7, cmp15, v16 bool
	var retval unsafe.Pointer
	var v1, v3, v5, v7, v9, v15, call int32
	var _type, type1, type2, size, size6, size14 unsafe.Pointer
	var conv int64
	var v0, v2, v4, v6, v8, v10, v11, v12, v13, v14 unsafe.Pointer
	var self_addr, other_addr, contents, contents12 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, self_addr, other_addr, v0, _type, v1, v2, type1, v3, cmp, v4, type2, v5, cmp3, v6, custom_tag_name, size, v7, v8, custom_tag_name5, size6, v9, cmp7, v10, custom_tag_name10, contents, v11, v12, custom_tag_name11, contents12, v13, v14, custom_tag_name13, size14, v15, conv, call, cmp15, v16

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	self_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	other_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](self_addr) = self
	*libc.As[unsafe.Pointer](other_addr) = other
	v0 = *libc.As[unsafe.Pointer](self_addr)
	_type = libc.Ptr(&libc.As[Tag](v0).F0)
	v1 = *libc.As[int32](_type)
	v2 = *libc.As[unsafe.Pointer](other_addr)
	type1 = libc.Ptr(&libc.As[Tag](v2).F0)
	v3 = *libc.As[int32](type1)
	cmp = v1 != v3
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v4 = *libc.As[unsafe.Pointer](self_addr)
	type2 = libc.Ptr(&libc.As[Tag](v4).F0)
	v5 = *libc.As[int32](type2)
	cmp3 = v5 == 126
	if cmp3 {
		goto if_then4
	} else {
		goto if_end19
	}

if_then4:
	v6 = *libc.As[unsafe.Pointer](self_addr)
	custom_tag_name = libc.Ptr(&libc.As[Tag](v6).F1)
	size = libc.Ptr(&libc.As[String](custom_tag_name).F1)
	v7 = *libc.As[int32](size)
	v8 = *libc.As[unsafe.Pointer](other_addr)
	custom_tag_name5 = libc.Ptr(&libc.As[Tag](v8).F1)
	size6 = libc.Ptr(&libc.As[String](custom_tag_name5).F1)
	v9 = *libc.As[int32](size6)
	cmp7 = v7 != v9
	if cmp7 {
		goto if_then8
	} else {
		goto if_end9
	}

if_then8:
	*libc.As[bool](retval) = false
	goto _return

if_end9:
	v10 = *libc.As[unsafe.Pointer](self_addr)
	custom_tag_name10 = libc.Ptr(&libc.As[Tag](v10).F1)
	contents = libc.Ptr(&libc.As[String](custom_tag_name10).F0)
	v11 = *libc.As[unsafe.Pointer](contents)
	v12 = *libc.As[unsafe.Pointer](other_addr)
	custom_tag_name11 = libc.Ptr(&libc.As[Tag](v12).F1)
	contents12 = libc.Ptr(&libc.As[String](custom_tag_name11).F0)
	v13 = *libc.As[unsafe.Pointer](contents12)
	v14 = *libc.As[unsafe.Pointer](self_addr)
	custom_tag_name13 = libc.Ptr(&libc.As[Tag](v14).F1)
	size14 = libc.Ptr(&libc.As[String](custom_tag_name13).F1)
	v15 = *libc.As[int32](size14)
	conv = int64(uint64(uint32(v15)))
	call = libc.Memcmp(libc.As[byte](v11), libc.As[byte](v13), conv)
	cmp15 = call != 0
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*libc.As[bool](retval) = false
	goto _return

if_end18:
	goto if_end19

if_end19:
	*libc.As[bool](retval) = true
	goto _return

_return:
	v16 = *libc.As[bool](retval)
	return v16
}
func tag_can_contain(self unsafe.Pointer, other unsafe.Pointer) bool {
	var cmp, cmp2, cmp4, cmp5, v8, cmp7, cmp9, cmp14, cmp17, cmp19, cmp22, v18, cmp26, cmp29, cmp32, cmp35, cmp38, v24, v25 bool
	var retval unsafe.Pointer
	var v1, v2, v4, v5, v6, v7, v9, v10, v11, v12, v13, inc, v14, v15, v16, v17, v19, v20, v21, v22, v23 int32
	var child, i, _type, type1, arrayidx unsafe.Pointer
	var conv, idxprom int64
	var v0, v3 unsafe.Pointer
	var self_addr, other_addr unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, self_addr, other_addr, child, i, v0, _type, v1, v2, cmp, v3, type1, v4, v5, cmp2, v6, cmp4, v7, cmp5, v8, v9, conv, cmp7, v10, v11, idxprom, arrayidx, v12, cmp9, v13, inc, v14, cmp14, v15, cmp17, v16, cmp19, v17, cmp22, v18, v19, cmp26, v20, cmp29, v21, cmp32, v22, cmp35, v23, cmp38, v24, v25

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	self_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	other_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	child = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](self_addr) = self
	*libc.As[unsafe.Pointer](other_addr) = other
	v0 = *libc.As[unsafe.Pointer](other_addr)
	_type = libc.Ptr(&libc.As[Tag](v0).F0)
	v1 = *libc.As[int32](_type)
	*libc.As[int32](child) = v1
	v2 = *libc.As[int32](child)
	cmp = v2 == 124
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = true
	goto _return

if_end:
	v3 = *libc.As[unsafe.Pointer](self_addr)
	type1 = libc.Ptr(&libc.As[Tag](v3).F0)
	v4 = *libc.As[int32](type1)
	switch v4 {
	case 71:
		goto sw_bb
	case 48:
		goto sw_bb3
	case 41:
		goto sw_bb3
	case 85:
		goto sw_bb6
	case 38:
		goto sw_bb13
	case 90:
		goto sw_bb16
	case 92:
		goto sw_bb16
	case 91:
		goto sw_bb16
	case 82:
		goto sw_bb25
	case 119:
		goto sw_bb28
	case 111:
		goto sw_bb31
	case 115:
		goto sw_bb31
	default:
		goto sw_default
	}

sw_bb:
	v5 = *libc.As[int32](child)
	cmp2 = v5 != 71
	*libc.As[bool](retval) = cmp2
	goto _return

sw_bb3:
	v6 = *libc.As[int32](child)
	cmp4 = v6 != 48
	if cmp4 {
		goto land_rhs
	} else {
		v8 = false
		goto land_end
	}

land_rhs:
	v7 = *libc.As[int32](child)
	cmp5 = v7 != 41
	v8 = cmp5
	goto land_end

land_end:
	*libc.As[bool](retval) = v8
	goto _return

sw_bb6:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v9 = *libc.As[int32](i)
	conv = int64(v9)
	cmp7 = uint64(conv) < uint64(26)
	if cmp7 {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v10 = *libc.As[int32](child)
	v11 = *libc.As[int32](i)
	idxprom = int64(v11)
	arrayidx = libc.Ptr(&TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS[idxprom])
	v12 = *libc.As[int32](arrayidx)
	cmp9 = v10 == v12
	if cmp9 {
		goto if_then11
	} else {
		goto if_end12
	}

if_then11:
	*libc.As[bool](retval) = false
	goto _return

if_end12:
	goto for_inc

for_inc:
	v13 = *libc.As[int32](i)
	inc = v13 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	*libc.As[bool](retval) = true
	goto _return

sw_bb13:
	v14 = *libc.As[int32](child)
	cmp14 = v14 == 3
	*libc.As[bool](retval) = cmp14
	goto _return

sw_bb16:
	v15 = *libc.As[int32](child)
	cmp17 = v15 != 90
	if cmp17 {
		goto land_lhs_true
	} else {
		v18 = false
		goto land_end24
	}

land_lhs_true:
	v16 = *libc.As[int32](child)
	cmp19 = v16 != 92
	if cmp19 {
		goto land_rhs21
	} else {
		v18 = false
		goto land_end24
	}

land_rhs21:
	v17 = *libc.As[int32](child)
	cmp22 = v17 != 91
	v18 = cmp22
	goto land_end24

land_end24:
	*libc.As[bool](retval) = v18
	goto _return

sw_bb25:
	v19 = *libc.As[int32](child)
	cmp26 = v19 != 82
	*libc.As[bool](retval) = cmp26
	goto _return

sw_bb28:
	v20 = *libc.As[int32](child)
	cmp29 = v20 != 119
	*libc.As[bool](retval) = cmp29
	goto _return

sw_bb31:
	v21 = *libc.As[int32](child)
	cmp32 = v21 != 111
	if cmp32 {
		goto land_lhs_true34
	} else {
		v24 = false
		goto land_end40
	}

land_lhs_true34:
	v22 = *libc.As[int32](child)
	cmp35 = v22 != 115
	if cmp35 {
		goto land_rhs37
	} else {
		v24 = false
		goto land_end40
	}

land_rhs37:
	v23 = *libc.As[int32](child)
	cmp38 = v23 != 119
	v24 = cmp38
	goto land_end40

land_end40:
	*libc.As[bool](retval) = v24
	goto _return

sw_default:
	*libc.As[bool](retval) = true
	goto _return

_return:
	v25 = *libc.As[bool](retval)
	return v25
}
func tag_type_for_name(tag_name unsafe.Pointer) int32 {
	var arrayidx unsafe.Pointer
	var tag_name2, tag_name5 unsafe.Pointer
	var cmp, cmp3, cmp10 bool
	var v0, v1, v4, v9, call9, v11, v12, inc, v13 int32
	var retval, i, size, size7, tag_type unsafe.Pointer
	var idxprom, call, conv, conv8 int64
	var arraydecay, arraydecay6 unsafe.Pointer
	var v2, v3, v5, v6, v7, v8, v10 unsafe.Pointer
	var tag_name_addr, entry1, contents unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, tag_name_addr, i, entry1, v0, cmp, v1, idxprom, arrayidx, v2, tag_name2, arraydecay, call, v3, size, v4, conv, cmp3, v5, contents, v6, v7, tag_name5, arraydecay6, v8, size7, v9, conv8, call9, cmp10, v10, tag_type, v11, v12, inc, v13

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	tag_name_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	entry1 = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](tag_name_addr) = tag_name
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v0 = *libc.As[int32](i)
	cmp = v0 < 126
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v1 = *libc.As[int32](i)
	idxprom = int64(v1)
	arrayidx = libc.Ptr(&TAG_TYPES_BY_TAG_NAME[idxprom])
	*libc.As[unsafe.Pointer](entry1) = arrayidx
	v2 = *libc.As[unsafe.Pointer](entry1)
	tag_name2 = libc.Ptr(&libc.As[TagMapEntry](v2).F0)
	arraydecay = libc.Ptr(&libc.As[[16]byte](tag_name2)[int64(0)])
	call = libc.Strlen(libc.As[byte](arraydecay))
	v3 = *libc.As[unsafe.Pointer](tag_name_addr)
	size = libc.Ptr(&libc.As[String](v3).F1)
	v4 = *libc.As[int32](size)
	conv = int64(uint64(uint32(v4)))
	cmp3 = call == conv
	if cmp3 {
		goto land_lhs_true
	} else {
		goto if_end
	}

land_lhs_true:
	v5 = *libc.As[unsafe.Pointer](tag_name_addr)
	contents = libc.Ptr(&libc.As[String](v5).F0)
	v6 = *libc.As[unsafe.Pointer](contents)
	v7 = *libc.As[unsafe.Pointer](entry1)
	tag_name5 = libc.Ptr(&libc.As[TagMapEntry](v7).F0)
	arraydecay6 = libc.Ptr(&libc.As[[16]byte](tag_name5)[int64(0)])
	v8 = *libc.As[unsafe.Pointer](tag_name_addr)
	size7 = libc.Ptr(&libc.As[String](v8).F1)
	v9 = *libc.As[int32](size7)
	conv8 = int64(uint64(uint32(v9)))
	call9 = libc.Memcmp(libc.As[byte](v6), libc.As[byte](arraydecay6), conv8)
	cmp10 = call9 == 0
	if cmp10 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v10 = *libc.As[unsafe.Pointer](entry1)
	tag_type = libc.Ptr(&libc.As[TagMapEntry](v10).F1)
	v11 = *libc.As[int32](tag_type)
	*libc.As[int32](retval) = v11
	goto _return

if_end:
	goto for_inc

for_inc:
	v12 = *libc.As[int32](i)
	inc = v12 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	*libc.As[int32](retval) = 126
	goto _return

_return:
	v13 = *libc.As[int32](retval)
	return v13
}
func tag_new(agg_result unsafe.Pointer) {
	var _compoundliteral, custom_tag_name unsafe.Pointer
	var _type, size, capacity unsafe.Pointer
	var contents unsafe.Pointer
	_, _, _, _, _, _ = _compoundliteral, _type, custom_tag_name, contents, size, capacity

	_compoundliteral = libc.Ptr(&new(struct {
		_ [0]uint64
		v String
		b byte
	}).v)
	_type = libc.Ptr(&libc.As[Tag](agg_result).F0)
	*libc.As[int32](_type) = 127
	custom_tag_name = libc.Ptr(&libc.As[Tag](agg_result).F1)
	contents = libc.Ptr(&libc.As[String](_compoundliteral).F0)
	*libc.As[unsafe.Pointer](contents) = nil
	size = libc.Ptr(&libc.As[String](_compoundliteral).F1)
	*libc.As[int32](size) = 0
	capacity = libc.Ptr(&libc.As[String](_compoundliteral).F2)
	*libc.As[int32](capacity) = 0
	libc.Memmove(libc.As[byte](custom_tag_name), libc.As[byte](_compoundliteral), int64(16))
}
func _array__reserve(self unsafe.Pointer, element_size int64, new_capacity int32) {
	var cmp, tobool bool
	var v0, v2, v7, v10, v13 int32
	var new_capacity_addr, capacity, capacity8 unsafe.Pointer
	var conv, v8, mul, conv4, v11, mul5 int64
	var element_size_addr unsafe.Pointer
	var v1, v3, v4, v5, v6, call, v9, call6, v12, v14 unsafe.Pointer
	var self_addr, contents, contents2, contents3, contents7 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = self_addr, element_size_addr, new_capacity_addr, v0, v1, capacity, v2, cmp, v3, contents, v4, tobool, v5, contents2, v6, v7, conv, v8, mul, call, v9, contents3, v10, conv4, v11, mul5, call6, v12, contents7, v13, v14, capacity8

	self_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	element_size_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int64
		b byte
	}).v)
	new_capacity_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](self_addr) = self
	*libc.As[int64](element_size_addr) = element_size
	*libc.As[int32](new_capacity_addr) = new_capacity
	v0 = *libc.As[int32](new_capacity_addr)
	v1 = *libc.As[unsafe.Pointer](self_addr)
	capacity = libc.Ptr(&libc.As[Array](v1).F2)
	v2 = *libc.As[int32](capacity)
	cmp = uint32(v0) > uint32(v2)
	if cmp {
		goto if_then
	} else {
		goto if_end9
	}

if_then:
	v3 = *libc.As[unsafe.Pointer](self_addr)
	contents = libc.Ptr(&libc.As[Array](v3).F0)
	v4 = *libc.As[unsafe.Pointer](contents)
	tobool = uintptr(unsafe.Pointer(v4)) != uintptr(unsafe.Pointer(nil))
	if tobool {
		goto if_then1
	} else {
		goto if_else
	}

if_then1:
	v5 = *libc.As[unsafe.Pointer](self_addr)
	contents2 = libc.Ptr(&libc.As[Array](v5).F0)
	v6 = *libc.As[unsafe.Pointer](contents2)
	v7 = *libc.As[int32](new_capacity_addr)
	conv = int64(uint64(uint32(v7)))
	v8 = *libc.As[int64](element_size_addr)
	mul = conv * v8
	call = libc.Ptr(libc.Realloc(libc.As[byte](v6), mul))
	v9 = *libc.As[unsafe.Pointer](self_addr)
	contents3 = libc.Ptr(&libc.As[Array](v9).F0)
	*libc.As[unsafe.Pointer](contents3) = call
	goto if_end

if_else:
	v10 = *libc.As[int32](new_capacity_addr)
	conv4 = int64(uint64(uint32(v10)))
	v11 = *libc.As[int64](element_size_addr)
	mul5 = conv4 * v11
	call6 = libc.Ptr(libc.Malloc[byte](mul5))
	v12 = *libc.As[unsafe.Pointer](self_addr)
	contents7 = libc.Ptr(&libc.As[Array](v12).F0)
	*libc.As[unsafe.Pointer](contents7) = call6
	goto if_end

if_end:
	v13 = *libc.As[int32](new_capacity_addr)
	v14 = *libc.As[unsafe.Pointer](self_addr)
	capacity8 = libc.Ptr(&libc.As[Array](v14).F2)
	*libc.As[int32](capacity8) = v13
	goto if_end9

if_end9:
}
func malloc(a0 int64) unsafe.Pointer {
	panic("unsatisfied: malloc")
}
