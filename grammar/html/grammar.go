package grammar_html

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
type TSLexer struct {
	F0 int32
	F1 int16
	F2 unsafe.Pointer
	F3 unsafe.Pointer
	F4 unsafe.Pointer
	F5 unsafe.Pointer
	F6 unsafe.Pointer
	F7 unsafe.Pointer
}
type Tag struct {
	F0 int32
	F1 String
}
type String struct {
	F0 unsafe.Pointer
	F1 int32
	F2 int32
}
type Scanner struct {
	F0 struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}
}
type Array struct {
	F0 unsafe.Pointer
	F1 int32
	F2 int32
}

var tree_sitter_html_language struct {
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
var ts_small_parse_table [1265]int16 = [1265]int16{12, 3, 1, 24, 15, 1, 1, 17, 1, 5, 19, 1, 7, 23, 1, 22, 5, 1, 31, 21, 1, 34, 32, 1, 35, 54, 1, 32, 55, 1, 33, 21, 2, 11, 16, 3, 7, 26, 27, 28, 29, 30, 36, 39, 12, 3, 1, 24, 15, 1, 1, 17, 1, 5, 19, 1, 7, 27, 1, 22, 5, 1, 31, 14, 1, 35, 21, 1, 34, 54, 1, 32, 55, 1, 33, 25, 2, 11, 16, 6, 7, 26, 27, 28, 29, 30, 36, 39, 12, 3, 1, 24, 15, 1, 1, 17, 1, 5, 29, 1, 7, 31, 1, 22, 5, 1, 31, 21, 1, 34, 27, 1, 35, 54, 1, 32, 55, 1, 33, 25, 2, 11, 16, 6, 7, 26, 27, 28, 29, 30, 36, 39, 12, 3, 1, 24, 15, 1, 1, 17, 1, 5, 29, 1, 7, 35, 1, 22, 5, 1, 31, 21, 1, 34, 22, 1, 35, 54, 1, 32, 55, 1, 33, 33, 2, 11, 16, 4, 7, 26, 27, 28, 29, 30, 36, 39, 11, 3, 1, 24, 37, 1, 1, 40, 1, 5, 43, 1, 7, 49, 1, 22, 5, 1, 31, 21, 1, 34, 54, 1, 32, 55, 1, 33, 46, 2, 11, 16, 6, 7, 26, 27, 28, 29, 30, 36, 39, 11, 3, 1, 24, 7, 1, 1, 9, 1, 5, 11, 1, 7, 51, 1, 0, 2, 1, 31, 33, 1, 34, 48, 1, 32, 53, 1, 33, 53, 2, 11, 16, 8, 7, 26, 27, 28, 29, 30, 36, 39, 11, 3, 1, 24, 49, 1, 0, 55, 1, 1, 58, 1, 5, 61, 1, 7, 2, 1, 31, 33, 1, 34, 48, 1, 32, 53, 1, 33, 64, 2, 11, 16, 8, 7, 26, 27, 28, 29, 30, 36, 39, 3, 3, 1, 24, 69, 1, 5, 67, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 73, 1, 5, 71, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 77, 1, 5, 75, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 81, 1, 5, 79, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 85, 1, 5, 83, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 89, 1, 5, 87, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 93, 1, 5, 91, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 97, 1, 5, 95, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 101, 1, 5, 99, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 105, 1, 5, 103, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 109, 1, 5, 107, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 69, 1, 5, 67, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 113, 1, 5, 111, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 117, 1, 5, 115, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 109, 1, 5, 107, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 73, 1, 5, 71, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 81, 1, 5, 79, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 85, 1, 5, 83, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 89, 1, 5, 87, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 93, 1, 5, 91, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 97, 1, 5, 95, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 101, 1, 5, 99, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 121, 1, 5, 119, 5, 22, 1, 7, 11, 16, 3, 3, 1, 24, 117, 1, 5, 115, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 113, 1, 5, 111, 5, 0, 1, 7, 11, 16, 3, 3, 1, 24, 121, 1, 5, 119, 5, 0, 1, 7, 11, 16, 4, 3, 1, 24, 125, 1, 9, 123, 2, 3, 6, 35, 2, 37, 40, 5, 3, 1, 24, 128, 1, 3, 130, 1, 6, 132, 1, 9, 35, 2, 37, 40, 5, 3, 1, 24, 132, 1, 9, 134, 1, 3, 136, 1, 6, 38, 2, 37, 40, 5, 3, 1, 24, 128, 1, 3, 132, 1, 9, 138, 1, 6, 35, 2, 37, 40, 5, 3, 1, 24, 132, 1, 9, 134, 1, 3, 140, 1, 6, 36, 2, 37, 40, 4, 3, 1, 24, 142, 1, 3, 144, 1, 9, 41, 2, 37, 40, 4, 3, 1, 24, 123, 1, 3, 146, 1, 9, 41, 2, 37, 40, 4, 3, 1, 24, 144, 1, 9, 149, 1, 3, 40, 2, 37, 40, 4, 3, 1, 24, 144, 1, 9, 151, 1, 3, 44, 2, 37, 40, 4, 3, 1, 24, 144, 1, 9, 153, 1, 3, 41, 2, 37, 40, 3, 3, 1, 24, 157, 1, 8, 155, 3, 3, 6, 9, 5, 3, 1, 24, 159, 1, 10, 161, 1, 12, 163, 1, 14, 56, 1, 38, 5, 3, 1, 24, 165, 1, 10, 167, 1, 12, 169, 1, 14, 58, 1, 38, 4, 3, 1, 24, 171, 1, 7, 173, 1, 23, 19, 1, 35, 2, 3, 1, 24, 175, 3, 3, 6, 9, 2, 3, 1, 24, 177, 3, 3, 6, 9, 3, 3, 1, 24, 179, 1, 8, 155, 2, 3, 9, 4, 3, 1, 24, 181, 1, 17, 183, 1, 18, 185, 1, 19, 4, 3, 1, 24, 171, 1, 7, 187, 1, 23, 10, 1, 35, 4, 3, 1, 24, 189, 1, 7, 191, 1, 23, 23, 1, 35, 4, 3, 1, 24, 189, 1, 7, 193, 1, 23, 24, 1, 35, 2, 3, 1, 24, 195, 3, 3, 6, 9, 4, 3, 1, 24, 183, 1, 18, 185, 1, 19, 197, 1, 17, 2, 3, 1, 24, 195, 2, 3, 9, 3, 3, 1, 24, 199, 1, 20, 201, 1, 21, 3, 3, 1, 24, 171, 1, 7, 16, 1, 35, 2, 3, 1, 24, 203, 2, 23, 7, 3, 3, 1, 24, 205, 1, 14, 207, 1, 15, 2, 3, 1, 24, 175, 2, 3, 9, 2, 3, 1, 24, 177, 2, 3, 9, 2, 3, 1, 24, 209, 2, 23, 7, 2, 3, 1, 24, 211, 2, 23, 7, 2, 3, 1, 24, 213, 2, 23, 7, 3, 3, 1, 24, 201, 1, 21, 215, 1, 20, 3, 3, 1, 24, 189, 1, 7, 28, 1, 35, 3, 3, 1, 24, 189, 1, 7, 29, 1, 35, 3, 3, 1, 24, 171, 1, 7, 15, 1, 35, 3, 3, 1, 24, 205, 1, 12, 217, 1, 13, 3, 3, 1, 24, 219, 1, 12, 221, 1, 13, 3, 3, 1, 24, 219, 1, 14, 223, 1, 15, 2, 3, 1, 24, 225, 1, 14, 2, 3, 1, 24, 227, 1, 14, 2, 3, 1, 24, 229, 1, 4, 2, 3, 1, 24, 227, 1, 12, 2, 3, 1, 24, 231, 1, 3, 2, 3, 1, 24, 233, 1, 3, 2, 3, 1, 24, 235, 1, 2, 2, 3, 1, 24, 237, 1, 3, 2, 3, 1, 24, 239, 1, 0, 2, 3, 1, 24, 241, 1, 3, 2, 3, 1, 24, 243, 1, 21, 2, 3, 1, 24, 225, 1, 12, 2, 3, 1, 24, 245, 1, 3, 2, 3, 1, 24, 215, 1, 20, 2, 3, 1, 24, 201, 1, 21, 2, 3, 1, 24, 247, 1, 2, 2, 3, 1, 24, 249, 1, 3, 2, 3, 1, 24, 199, 1, 20, 2, 3, 1, 24, 251, 1, 4}
var ts_small_parse_table_map [92]int32 = [92]int32{0, 44, 88, 132, 176, 217, 258, 299, 313, 327, 341, 355, 369, 383, 397, 411, 425, 439, 453, 467, 481, 495, 509, 523, 537, 551, 565, 579, 593, 607, 621, 635, 649, 663, 678, 695, 712, 729, 746, 760, 774, 788, 802, 816, 828, 844, 860, 873, 882, 891, 902, 915, 928, 941, 954, 963, 976, 984, 994, 1004, 1012, 1022, 1030, 1038, 1046, 1054, 1062, 1072, 1082, 1092, 1102, 1112, 1122, 1132, 1139, 1146, 1153, 1160, 1167, 1174, 1181, 1188, 1195, 1202, 1209, 1216, 1223, 1230, 1237, 1244, 1251, 1258}
var ts_symbol_names [41]unsafe.Pointer = [41]unsafe.Pointer{libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_16), libc.Ptr(&_str_19), libc.Ptr(&_str_16), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_21), libc.Ptr(&_str_21), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_10), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_30), libc.Ptr(&_str_31), libc.Ptr(&_str_31), libc.Ptr(&_str_31), libc.Ptr(&_str_32), libc.Ptr(&_str_33), libc.Ptr(&_str_34), libc.Ptr(&_str_35), libc.Ptr(&_str_36), libc.Ptr(&_str_37), libc.Ptr(&_str_38)}
var ts_symbol_metadata [41]TSSymbolMetadata = [41]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [41]int16 = [41]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 10, 14, 10, 16, 17, 17, 17, 17, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 31, 34, 35, 36, 37, 38, 39, 40}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][4]int16 = [1][4]int16{}
var ts_lex_modes [94]TSLexMode = [94]TSLexMode{TSLexMode{0, 1}, TSLexMode{17, 2}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 3}, TSLexMode{17, 2}, TSLexMode{17, 3}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 3}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 3}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{17, 2}, TSLexMode{5, 4}, TSLexMode{5, 4}, TSLexMode{5, 4}, TSLexMode{5, 4}, TSLexMode{5, 4}, TSLexMode{5, 2}, TSLexMode{5, 2}, TSLexMode{5, 2}, TSLexMode{5, 2}, TSLexMode{5, 2}, TSLexMode{5, 4}, TSLexMode{1, 2}, TSLexMode{1, 2}, TSLexMode{0, 5}, TSLexMode{5, 4}, TSLexMode{5, 4}, TSLexMode{5, 2}, TSLexMode{0, 6}, TSLexMode{0, 5}, TSLexMode{0, 5}, TSLexMode{0, 5}, TSLexMode{5, 4}, TSLexMode{0, 6}, TSLexMode{5, 2}, TSLexMode{0, 7}, TSLexMode{0, 2}, TSLexMode{0, 5}, TSLexMode{2, 2}, TSLexMode{5, 2}, TSLexMode{5, 2}, TSLexMode{0, 5}, TSLexMode{0, 5}, TSLexMode{0, 5}, TSLexMode{0, 7}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{4, 2}, TSLexMode{4, 2}, TSLexMode{2, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{15, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 8}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{0, 9}, TSLexMode{0, 8}, TSLexMode{15, 2}, TSLexMode{0, 2}, TSLexMode{0, 9}, TSLexMode{0, 2}}
var ts_external_scanner_states [10][9]byte = [10][9]byte{[9]byte{}, [9]byte{1, 1, 1, 1, 1, 1, 1, 1, 1}, [9]byte{0, 0, 0, 0, 0, 0, 0, 0, 1}, [9]byte{0, 0, 0, 0, 0, 0, 1, 0, 1}, [9]byte{0, 0, 0, 0, 0, 1, 0, 0, 1}, [9]byte{0, 0, 0, 0, 0, 0, 0, 1, 1}, [9]byte{1, 1, 1, 0, 0, 0, 0, 0, 1}, [9]byte{0, 0, 0, 1, 1, 0, 0, 0, 1}, [9]byte{0, 0, 0, 0, 1, 0, 0, 0, 1}, [9]byte{0, 0, 0, 1, 0, 0, 0, 0, 1}}
var ts_external_scanner_symbol_map [9]int16 = [9]int16{17, 18, 19, 20, 21, 6, 22, 23, 24}
var ts_primary_state_ids [94]int16 = [94]int16{0, 1, 2, 3, 3, 2, 6, 7, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 9, 21, 22, 19, 10, 12, 13, 14, 15, 16, 17, 31, 22, 21, 31, 35, 36, 37, 36, 37, 40, 35, 42, 43, 44, 45, 46, 46, 48, 49, 50, 45, 52, 53, 48, 53, 56, 52, 56, 59, 60, 61, 62, 49, 50, 65, 66, 67, 59, 69, 60, 69, 72, 72, 62, 75, 75, 77, 78, 79, 79, 81, 82, 83, 84, 85, 78, 82, 88, 85, 81, 84, 88, 77}
var _str [64]byte = [64]byte{40, 117, 105, 110, 116, 51, 50, 95, 116, 41, 40, 40, 38, 115, 99, 97, 110, 110, 101, 114, 45, 62, 116, 97, 103, 115, 41, 45, 62, 115, 105, 122, 101, 32, 45, 32, 49, 41, 32, 60, 32, 40, 38, 115, 99, 97, 110, 110, 101, 114, 45, 62, 116, 97, 103, 115, 41, 45, 62, 115, 105, 122, 101, 0}
var _str_1 [39]byte = [39]byte{47, 116, 109, 112, 47, 108, 101, 97, 118, 101, 110, 45, 104, 116, 109, 108, 45, 49, 55, 57, 52, 50, 53, 49, 55, 57, 51, 47, 99, 111, 109, 98, 105, 110, 101, 100, 46, 99, 0}
var __PRETTY_FUNCTION___scan_raw_text [42]byte = [42]byte{95, 66, 111, 111, 108, 32, 115, 99, 97, 110, 95, 114, 97, 119, 95, 116, 101, 120, 116, 40, 83, 99, 97, 110, 110, 101, 114, 32, 42, 44, 32, 84, 83, 76, 101, 120, 101, 114, 32, 42, 41, 0}
var _str_2 [9]byte = [9]byte{60, 47, 83, 67, 82, 73, 80, 84, 0}
var _str_3 [8]byte = [8]byte{60, 47, 83, 84, 89, 76, 69, 0}
var __PRETTY_FUNCTION___scan_implicit_end_tag [50]byte = [50]byte{95, 66, 111, 111, 108, 32, 115, 99, 97, 110, 95, 105, 109, 112, 108, 105, 99, 105, 116, 95, 101, 110, 100, 95, 116, 97, 103, 40, 83, 99, 97, 110, 110, 101, 114, 32, 42, 44, 32, 84, 83, 76, 101, 120, 101, 114, 32, 42, 41, 0}
var TAG_TYPES_BY_TAG_NAME [126]TagMapEntry = [126]TagMapEntry{TagMapEntry{[16]byte{65, 82, 69, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0}, TagMapEntry{[16]byte{66, 65, 83, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1}, TagMapEntry{[16]byte{66, 65, 83, 69, 70, 79, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0}, 2}, TagMapEntry{[16]byte{66, 71, 83, 79, 85, 78, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 3}, TagMapEntry{[16]byte{66, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 4}, TagMapEntry{[16]byte{67, 79, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 5}, TagMapEntry{[16]byte{67, 79, 77, 77, 65, 78, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 6}, TagMapEntry{[16]byte{69, 77, 66, 69, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 7}, TagMapEntry{[16]byte{70, 82, 65, 77, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 8}, TagMapEntry{[16]byte{72, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 9}, TagMapEntry{[16]byte{73, 77, 65, 71, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10}, TagMapEntry{[16]byte{73, 77, 71, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 11}, TagMapEntry{[16]byte{73, 78, 80, 85, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 12}, TagMapEntry{[16]byte{73, 83, 73, 78, 68, 69, 88, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 13}, TagMapEntry{[16]byte{75, 69, 89, 71, 69, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 14}, TagMapEntry{[16]byte{76, 73, 78, 75, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 15}, TagMapEntry{[16]byte{77, 69, 78, 85, 73, 84, 69, 77, 0, 0, 0, 0, 0, 0, 0, 0}, 16}, TagMapEntry{[16]byte{77, 69, 84, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 17}, TagMapEntry{[16]byte{78, 69, 88, 84, 73, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 18}, TagMapEntry{[16]byte{80, 65, 82, 65, 77, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 19}, TagMapEntry{[16]byte{83, 79, 85, 82, 67, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 20}, TagMapEntry{[16]byte{84, 82, 65, 67, 75, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 21}, TagMapEntry{[16]byte{87, 66, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 22}, TagMapEntry{[16]byte{65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 24}, TagMapEntry{[16]byte{65, 66, 66, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 25}, TagMapEntry{[16]byte{65, 68, 68, 82, 69, 83, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 26}, TagMapEntry{[16]byte{65, 82, 84, 73, 67, 76, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 27}, TagMapEntry{[16]byte{65, 83, 73, 68, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 28}, TagMapEntry{[16]byte{65, 85, 68, 73, 79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 29}, TagMapEntry{[16]byte{66, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 30}, TagMapEntry{[16]byte{66, 68, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 31}, TagMapEntry{[16]byte{66, 68, 79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 32}, TagMapEntry{[16]byte{66, 76, 79, 67, 75, 81, 85, 79, 84, 69, 0, 0, 0, 0, 0, 0}, 33}, TagMapEntry{[16]byte{66, 79, 68, 89, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 34}, TagMapEntry{[16]byte{66, 85, 84, 84, 79, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 35}, TagMapEntry{[16]byte{67, 65, 78, 86, 65, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 36}, TagMapEntry{[16]byte{67, 65, 80, 84, 73, 79, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 37}, TagMapEntry{[16]byte{67, 73, 84, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 38}, TagMapEntry{[16]byte{67, 79, 68, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 39}, TagMapEntry{[16]byte{67, 79, 76, 71, 82, 79, 85, 80, 0, 0, 0, 0, 0, 0, 0, 0}, 40}, TagMapEntry{[16]byte{68, 65, 84, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 41}, TagMapEntry{[16]byte{68, 65, 84, 65, 76, 73, 83, 84, 0, 0, 0, 0, 0, 0, 0, 0}, 42}, TagMapEntry{[16]byte{68, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 43}, TagMapEntry{[16]byte{68, 69, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 44}, TagMapEntry{[16]byte{68, 69, 84, 65, 73, 76, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 45}, TagMapEntry{[16]byte{68, 70, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 46}, TagMapEntry{[16]byte{68, 73, 65, 76, 79, 71, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 47}, TagMapEntry{[16]byte{68, 73, 86, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 48}, TagMapEntry{[16]byte{68, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 49}, TagMapEntry{[16]byte{68, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 50}, TagMapEntry{[16]byte{69, 77, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 51}, TagMapEntry{[16]byte{70, 73, 69, 76, 68, 83, 69, 84, 0, 0, 0, 0, 0, 0, 0, 0}, 52}, TagMapEntry{[16]byte{70, 73, 71, 67, 65, 80, 84, 73, 79, 78, 0, 0, 0, 0, 0, 0}, 53}, TagMapEntry{[16]byte{70, 73, 71, 85, 82, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 54}, TagMapEntry{[16]byte{70, 79, 79, 84, 69, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 55}, TagMapEntry{[16]byte{70, 79, 82, 77, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 56}, TagMapEntry{[16]byte{72, 49, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 57}, TagMapEntry{[16]byte{72, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 58}, TagMapEntry{[16]byte{72, 51, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 59}, TagMapEntry{[16]byte{72, 52, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 60}, TagMapEntry{[16]byte{72, 53, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 61}, TagMapEntry{[16]byte{72, 54, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 62}, TagMapEntry{[16]byte{72, 69, 65, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 63}, TagMapEntry{[16]byte{72, 69, 65, 68, 69, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 64}, TagMapEntry{[16]byte{72, 71, 82, 79, 85, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 65}, TagMapEntry{[16]byte{72, 84, 77, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 66}, TagMapEntry{[16]byte{73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 67}, TagMapEntry{[16]byte{73, 70, 82, 65, 77, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 68}, TagMapEntry{[16]byte{73, 78, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 69}, TagMapEntry{[16]byte{75, 66, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 70}, TagMapEntry{[16]byte{76, 65, 66, 69, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 71}, TagMapEntry{[16]byte{76, 69, 71, 69, 78, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 72}, TagMapEntry{[16]byte{76, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 73}, TagMapEntry{[16]byte{77, 65, 73, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 74}, TagMapEntry{[16]byte{77, 65, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 75}, TagMapEntry{[16]byte{77, 65, 82, 75, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 76}, TagMapEntry{[16]byte{77, 65, 84, 72, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 77}, TagMapEntry{[16]byte{77, 69, 78, 85, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 78}, TagMapEntry{[16]byte{77, 69, 84, 69, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 79}, TagMapEntry{[16]byte{78, 65, 86, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 80}, TagMapEntry{[16]byte{78, 79, 83, 67, 82, 73, 80, 84, 0, 0, 0, 0, 0, 0, 0, 0}, 81}, TagMapEntry{[16]byte{79, 66, 74, 69, 67, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 82}, TagMapEntry{[16]byte{79, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 83}, TagMapEntry{[16]byte{79, 80, 84, 71, 82, 79, 85, 80, 0, 0, 0, 0, 0, 0, 0, 0}, 84}, TagMapEntry{[16]byte{79, 80, 84, 73, 79, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 85}, TagMapEntry{[16]byte{79, 85, 84, 80, 85, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 86}, TagMapEntry{[16]byte{80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 87}, TagMapEntry{[16]byte{80, 73, 67, 84, 85, 82, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 88}, TagMapEntry{[16]byte{80, 82, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 89}, TagMapEntry{[16]byte{80, 82, 79, 71, 82, 69, 83, 83, 0, 0, 0, 0, 0, 0, 0, 0}, 90}, TagMapEntry{[16]byte{81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 91}, TagMapEntry{[16]byte{82, 66, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 92}, TagMapEntry{[16]byte{82, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 93}, TagMapEntry{[16]byte{82, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 94}, TagMapEntry{[16]byte{82, 84, 67, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 95}, TagMapEntry{[16]byte{82, 85, 66, 89, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 96}, TagMapEntry{[16]byte{83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 97}, TagMapEntry{[16]byte{83, 65, 77, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 98}, TagMapEntry{[16]byte{83, 67, 82, 73, 80, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 99}, TagMapEntry{[16]byte{83, 69, 67, 84, 73, 79, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100}, TagMapEntry{[16]byte{83, 69, 76, 69, 67, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 101}, TagMapEntry{[16]byte{83, 76, 79, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 102}, TagMapEntry{[16]byte{83, 77, 65, 76, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 103}, TagMapEntry{[16]byte{83, 80, 65, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 104}, TagMapEntry{[16]byte{83, 84, 82, 79, 78, 71, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 105}, TagMapEntry{[16]byte{83, 84, 89, 76, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 106}, TagMapEntry{[16]byte{83, 85, 66, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 107}, TagMapEntry{[16]byte{83, 85, 77, 77, 65, 82, 89, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 108}, TagMapEntry{[16]byte{83, 85, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 109}, TagMapEntry{[16]byte{83, 86, 71, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 110}, TagMapEntry{[16]byte{84, 65, 66, 76, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 111}, TagMapEntry{[16]byte{84, 66, 79, 68, 89, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 112}, TagMapEntry{[16]byte{84, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 113}, TagMapEntry{[16]byte{84, 69, 77, 80, 76, 65, 84, 69, 0, 0, 0, 0, 0, 0, 0, 0}, 114}, TagMapEntry{[16]byte{84, 69, 88, 84, 65, 82, 69, 65, 0, 0, 0, 0, 0, 0, 0, 0}, 115}, TagMapEntry{[16]byte{84, 70, 79, 79, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 116}, TagMapEntry{[16]byte{84, 72, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 117}, TagMapEntry{[16]byte{84, 72, 69, 65, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 118}, TagMapEntry{[16]byte{84, 73, 77, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 119}, TagMapEntry{[16]byte{84, 73, 84, 76, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 120}, TagMapEntry{[16]byte{84, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 121}, TagMapEntry{[16]byte{85, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 122}, TagMapEntry{[16]byte{85, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 123}, TagMapEntry{[16]byte{86, 65, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 124}, TagMapEntry{[16]byte{86, 73, 68, 69, 79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 125}, TagMapEntry{[16]byte{67, 85, 83, 84, 79, 77, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 126}}
var TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS [26]int32 = [26]int32{26, 27, 28, 33, 45, 48, 49, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 64, 9, 74, 80, 83, 87, 89, 100}
var __PRETTY_FUNCTION___scan_end_tag_name [46]byte = [46]byte{95, 66, 111, 111, 108, 32, 115, 99, 97, 110, 95, 101, 110, 100, 95, 116, 97, 103, 95, 110, 97, 109, 101, 40, 83, 99, 97, 110, 110, 101, 114, 32, 42, 44, 32, 84, 83, 76, 101, 120, 101, 114, 32, 42, 41, 0}
var ts_parse_table struct {
	F0 struct {
		F0 [25]int16
		F1 [16]int16
	}
	F1 [41]int16
} = struct {
	F0 struct {
		F0 [25]int16
		F1 [16]int16
	}
	F1 [41]int16
}{struct {
	F0 [25]int16
	F1 [16]int16
}{[25]int16{1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 3}, [16]int16{}}, [41]int16{5, 7, 0, 0, 0, 9, 0, 11, 0, 0, 0, 13, 0, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 0, 3, 83, 7, 7, 7, 7, 7, 2, 48, 53, 33, 0, 7, 0, 0, 7, 0}}
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
	F38 TSParseActionEntry
	F39 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F40 struct {
		F0 anon_3
		F1 [6]byte
	}
	F41 TSParseActionEntry
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
	F44 TSParseActionEntry
	F45 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F46 struct {
		F0 anon_3
		F1 [6]byte
	}
	F47 TSParseActionEntry
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
	F50 TSParseActionEntry
	F51 struct {
		F0 anon_3
		F1 [6]byte
	}
	F52 TSParseActionEntry
	F53 struct {
		F0 anon_3
		F1 [6]byte
	}
	F54 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F55 struct {
		F0 anon_3
		F1 [6]byte
	}
	F56 TSParseActionEntry
	F57 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F58 struct {
		F0 anon_3
		F1 [6]byte
	}
	F59 TSParseActionEntry
	F60 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F61 struct {
		F0 anon_3
		F1 [6]byte
	}
	F62 TSParseActionEntry
	F63 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F64 struct {
		F0 anon_3
		F1 [6]byte
	}
	F65 TSParseActionEntry
	F66 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F67 struct {
		F0 anon_3
		F1 [6]byte
	}
	F68 TSParseActionEntry
	F69 struct {
		F0 anon_3
		F1 [6]byte
	}
	F70 TSParseActionEntry
	F71 struct {
		F0 anon_3
		F1 [6]byte
	}
	F72 TSParseActionEntry
	F73 struct {
		F0 anon_3
		F1 [6]byte
	}
	F74 TSParseActionEntry
	F75 struct {
		F0 anon_3
		F1 [6]byte
	}
	F76 TSParseActionEntry
	F77 struct {
		F0 anon_3
		F1 [6]byte
	}
	F78 TSParseActionEntry
	F79 struct {
		F0 anon_3
		F1 [6]byte
	}
	F80 TSParseActionEntry
	F81 struct {
		F0 anon_3
		F1 [6]byte
	}
	F82 TSParseActionEntry
	F83 struct {
		F0 anon_3
		F1 [6]byte
	}
	F84 TSParseActionEntry
	F85 struct {
		F0 anon_3
		F1 [6]byte
	}
	F86 TSParseActionEntry
	F87 struct {
		F0 anon_3
		F1 [6]byte
	}
	F88 TSParseActionEntry
	F89 struct {
		F0 anon_3
		F1 [6]byte
	}
	F90 TSParseActionEntry
	F91 struct {
		F0 anon_3
		F1 [6]byte
	}
	F92 TSParseActionEntry
	F93 struct {
		F0 anon_3
		F1 [6]byte
	}
	F94 TSParseActionEntry
	F95 struct {
		F0 anon_3
		F1 [6]byte
	}
	F96 TSParseActionEntry
	F97 struct {
		F0 anon_3
		F1 [6]byte
	}
	F98 TSParseActionEntry
	F99 struct {
		F0 anon_3
		F1 [6]byte
	}
	F100 TSParseActionEntry
	F101 struct {
		F0 anon_3
		F1 [6]byte
	}
	F102 TSParseActionEntry
	F103 struct {
		F0 anon_3
		F1 [6]byte
	}
	F104 TSParseActionEntry
	F105 struct {
		F0 anon_3
		F1 [6]byte
	}
	F106 TSParseActionEntry
	F107 struct {
		F0 anon_3
		F1 [6]byte
	}
	F108 TSParseActionEntry
	F109 struct {
		F0 anon_3
		F1 [6]byte
	}
	F110 TSParseActionEntry
	F111 struct {
		F0 anon_3
		F1 [6]byte
	}
	F112 TSParseActionEntry
	F113 struct {
		F0 anon_3
		F1 [6]byte
	}
	F114 TSParseActionEntry
	F115 struct {
		F0 anon_3
		F1 [6]byte
	}
	F116 TSParseActionEntry
	F117 struct {
		F0 anon_3
		F1 [6]byte
	}
	F118 TSParseActionEntry
	F119 struct {
		F0 anon_3
		F1 [6]byte
	}
	F120 TSParseActionEntry
	F121 struct {
		F0 anon_3
		F1 [6]byte
	}
	F122 TSParseActionEntry
	F123 struct {
		F0 anon_3
		F1 [6]byte
	}
	F124 TSParseActionEntry
	F125 struct {
		F0 anon_3
		F1 [6]byte
	}
	F126 TSParseActionEntry
	F127 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F128 struct {
		F0 anon_3
		F1 [6]byte
	}
	F129 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F130 struct {
		F0 anon_3
		F1 [6]byte
	}
	F131 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F132 struct {
		F0 anon_3
		F1 [6]byte
	}
	F133 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F134 struct {
		F0 anon_3
		F1 [6]byte
	}
	F135 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F136 struct {
		F0 anon_3
		F1 [6]byte
	}
	F137 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F138 struct {
		F0 anon_3
		F1 [6]byte
	}
	F139 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F140 struct {
		F0 anon_3
		F1 [6]byte
	}
	F141 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F142 struct {
		F0 anon_3
		F1 [6]byte
	}
	F143 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F144 struct {
		F0 anon_3
		F1 [6]byte
	}
	F145 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F146 struct {
		F0 anon_3
		F1 [6]byte
	}
	F147 TSParseActionEntry
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
	F156 TSParseActionEntry
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
	F176 TSParseActionEntry
	F177 struct {
		F0 anon_3
		F1 [6]byte
	}
	F178 TSParseActionEntry
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
	F192 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F193 struct {
		F0 anon_3
		F1 [6]byte
	}
	F194 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F195 struct {
		F0 anon_3
		F1 [6]byte
	}
	F196 TSParseActionEntry
	F197 struct {
		F0 anon_3
		F1 [6]byte
	}
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
	F200 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F201 struct {
		F0 anon_3
		F1 [6]byte
	}
	F202 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F203 struct {
		F0 anon_3
		F1 [6]byte
	}
	F204 TSParseActionEntry
	F205 struct {
		F0 anon_3
		F1 [6]byte
	}
	F206 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F207 struct {
		F0 anon_3
		F1 [6]byte
	}
	F208 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F216 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F217 struct {
		F0 anon_3
		F1 [6]byte
	}
	F218 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F219 struct {
		F0 anon_3
		F1 [6]byte
	}
	F220 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F221 struct {
		F0 anon_3
		F1 [6]byte
	}
	F222 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F223 struct {
		F0 anon_3
		F1 [6]byte
	}
	F224 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F225 struct {
		F0 anon_3
		F1 [6]byte
	}
	F226 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F227 struct {
		F0 anon_3
		F1 [6]byte
	}
	F228 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F229 struct {
		F0 anon_3
		F1 [6]byte
	}
	F230 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F231 struct {
		F0 anon_3
		F1 [6]byte
	}
	F232 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F233 struct {
		F0 anon_3
		F1 [6]byte
	}
	F234 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F235 struct {
		F0 anon_3
		F1 [6]byte
	}
	F236 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F237 struct {
		F0 anon_3
		F1 [6]byte
	}
	F238 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F239 struct {
		F0 anon_3
		F1 [6]byte
	}
	F240 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F241 struct {
		F0 anon_3
		F1 [6]byte
	}
	F242 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F243 struct {
		F0 anon_3
		F1 [6]byte
	}
	F244 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F245 struct {
		F0 anon_3
		F1 [6]byte
	}
	F246 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F247 struct {
		F0 anon_3
		F1 [6]byte
	}
	F248 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F249 struct {
		F0 anon_3
		F1 [6]byte
	}
	F250 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F251 struct {
		F0 anon_3
		F1 [6]byte
	}
	F252 struct {
		F0 struct {
			F0 struct {
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
	F38 TSParseActionEntry
	F39 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F40 struct {
		F0 anon_3
		F1 [6]byte
	}
	F41 TSParseActionEntry
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
	F44 TSParseActionEntry
	F45 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F46 struct {
		F0 anon_3
		F1 [6]byte
	}
	F47 TSParseActionEntry
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
	F50 TSParseActionEntry
	F51 struct {
		F0 anon_3
		F1 [6]byte
	}
	F52 TSParseActionEntry
	F53 struct {
		F0 anon_3
		F1 [6]byte
	}
	F54 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F55 struct {
		F0 anon_3
		F1 [6]byte
	}
	F56 TSParseActionEntry
	F57 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F58 struct {
		F0 anon_3
		F1 [6]byte
	}
	F59 TSParseActionEntry
	F60 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F61 struct {
		F0 anon_3
		F1 [6]byte
	}
	F62 TSParseActionEntry
	F63 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F64 struct {
		F0 anon_3
		F1 [6]byte
	}
	F65 TSParseActionEntry
	F66 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F67 struct {
		F0 anon_3
		F1 [6]byte
	}
	F68 TSParseActionEntry
	F69 struct {
		F0 anon_3
		F1 [6]byte
	}
	F70 TSParseActionEntry
	F71 struct {
		F0 anon_3
		F1 [6]byte
	}
	F72 TSParseActionEntry
	F73 struct {
		F0 anon_3
		F1 [6]byte
	}
	F74 TSParseActionEntry
	F75 struct {
		F0 anon_3
		F1 [6]byte
	}
	F76 TSParseActionEntry
	F77 struct {
		F0 anon_3
		F1 [6]byte
	}
	F78 TSParseActionEntry
	F79 struct {
		F0 anon_3
		F1 [6]byte
	}
	F80 TSParseActionEntry
	F81 struct {
		F0 anon_3
		F1 [6]byte
	}
	F82 TSParseActionEntry
	F83 struct {
		F0 anon_3
		F1 [6]byte
	}
	F84 TSParseActionEntry
	F85 struct {
		F0 anon_3
		F1 [6]byte
	}
	F86 TSParseActionEntry
	F87 struct {
		F0 anon_3
		F1 [6]byte
	}
	F88 TSParseActionEntry
	F89 struct {
		F0 anon_3
		F1 [6]byte
	}
	F90 TSParseActionEntry
	F91 struct {
		F0 anon_3
		F1 [6]byte
	}
	F92 TSParseActionEntry
	F93 struct {
		F0 anon_3
		F1 [6]byte
	}
	F94 TSParseActionEntry
	F95 struct {
		F0 anon_3
		F1 [6]byte
	}
	F96 TSParseActionEntry
	F97 struct {
		F0 anon_3
		F1 [6]byte
	}
	F98 TSParseActionEntry
	F99 struct {
		F0 anon_3
		F1 [6]byte
	}
	F100 TSParseActionEntry
	F101 struct {
		F0 anon_3
		F1 [6]byte
	}
	F102 TSParseActionEntry
	F103 struct {
		F0 anon_3
		F1 [6]byte
	}
	F104 TSParseActionEntry
	F105 struct {
		F0 anon_3
		F1 [6]byte
	}
	F106 TSParseActionEntry
	F107 struct {
		F0 anon_3
		F1 [6]byte
	}
	F108 TSParseActionEntry
	F109 struct {
		F0 anon_3
		F1 [6]byte
	}
	F110 TSParseActionEntry
	F111 struct {
		F0 anon_3
		F1 [6]byte
	}
	F112 TSParseActionEntry
	F113 struct {
		F0 anon_3
		F1 [6]byte
	}
	F114 TSParseActionEntry
	F115 struct {
		F0 anon_3
		F1 [6]byte
	}
	F116 TSParseActionEntry
	F117 struct {
		F0 anon_3
		F1 [6]byte
	}
	F118 TSParseActionEntry
	F119 struct {
		F0 anon_3
		F1 [6]byte
	}
	F120 TSParseActionEntry
	F121 struct {
		F0 anon_3
		F1 [6]byte
	}
	F122 TSParseActionEntry
	F123 struct {
		F0 anon_3
		F1 [6]byte
	}
	F124 TSParseActionEntry
	F125 struct {
		F0 anon_3
		F1 [6]byte
	}
	F126 TSParseActionEntry
	F127 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F128 struct {
		F0 anon_3
		F1 [6]byte
	}
	F129 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F130 struct {
		F0 anon_3
		F1 [6]byte
	}
	F131 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F132 struct {
		F0 anon_3
		F1 [6]byte
	}
	F133 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F134 struct {
		F0 anon_3
		F1 [6]byte
	}
	F135 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F136 struct {
		F0 anon_3
		F1 [6]byte
	}
	F137 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F138 struct {
		F0 anon_3
		F1 [6]byte
	}
	F139 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F140 struct {
		F0 anon_3
		F1 [6]byte
	}
	F141 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F142 struct {
		F0 anon_3
		F1 [6]byte
	}
	F143 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F144 struct {
		F0 anon_3
		F1 [6]byte
	}
	F145 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F146 struct {
		F0 anon_3
		F1 [6]byte
	}
	F147 TSParseActionEntry
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
	F156 TSParseActionEntry
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
	F176 TSParseActionEntry
	F177 struct {
		F0 anon_3
		F1 [6]byte
	}
	F178 TSParseActionEntry
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
	F192 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F193 struct {
		F0 anon_3
		F1 [6]byte
	}
	F194 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F195 struct {
		F0 anon_3
		F1 [6]byte
	}
	F196 TSParseActionEntry
	F197 struct {
		F0 anon_3
		F1 [6]byte
	}
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
	F200 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F201 struct {
		F0 anon_3
		F1 [6]byte
	}
	F202 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F203 struct {
		F0 anon_3
		F1 [6]byte
	}
	F204 TSParseActionEntry
	F205 struct {
		F0 anon_3
		F1 [6]byte
	}
	F206 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F207 struct {
		F0 anon_3
		F1 [6]byte
	}
	F208 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F216 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F217 struct {
		F0 anon_3
		F1 [6]byte
	}
	F218 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F219 struct {
		F0 anon_3
		F1 [6]byte
	}
	F220 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F221 struct {
		F0 anon_3
		F1 [6]byte
	}
	F222 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F223 struct {
		F0 anon_3
		F1 [6]byte
	}
	F224 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F225 struct {
		F0 anon_3
		F1 [6]byte
	}
	F226 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F227 struct {
		F0 anon_3
		F1 [6]byte
	}
	F228 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F229 struct {
		F0 anon_3
		F1 [6]byte
	}
	F230 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F231 struct {
		F0 anon_3
		F1 [6]byte
	}
	F232 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F233 struct {
		F0 anon_3
		F1 [6]byte
	}
	F234 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F235 struct {
		F0 anon_3
		F1 [6]byte
	}
	F236 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F237 struct {
		F0 anon_3
		F1 [6]byte
	}
	F238 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F239 struct {
		F0 anon_3
		F1 [6]byte
	}
	F240 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F241 struct {
		F0 anon_3
		F1 [6]byte
	}
	F242 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F243 struct {
		F0 anon_3
		F1 [6]byte
	}
	F244 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F245 struct {
		F0 anon_3
		F1 [6]byte
	}
	F246 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F247 struct {
		F0 anon_3
		F1 [6]byte
	}
	F248 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F249 struct {
		F0 anon_3
		F1 [6]byte
	}
	F250 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F251 struct {
		F0 anon_3
		F1 [6]byte
	}
	F252 struct {
		F0 struct {
			F0 struct {
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
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 0, 25, 0, 0}}}, struct {
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
}{0, 0, 85, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 93, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 57, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 32, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 6, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 59, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 27, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 4, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 22, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 93, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 57, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 89, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 6, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 25, 0, 0}}}, struct {
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
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 77, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 52, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 85, 0, 1}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 35, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 35, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 30, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 30, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 31, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 31, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 34, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 34, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 36, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 36, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 28, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 28, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 29, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 29, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 30, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 30, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 26, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 26, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 31, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 31, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 29, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 29, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 28, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 28, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 28, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 28, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 34, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 34, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 40, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 40, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 45, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 11, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 51, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 40, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 51, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 66, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 65, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 1, 37, 0, 0}}}, struct {
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
}{0, 0, 58, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 74, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 88, 0, 0}, [2]byte{}}}, struct {
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
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 2, 38, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 38, 0, 0}}}, struct {
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
}{0, 0, 47, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 39, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 42, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 92, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 69, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 70, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 37, 0, 0}}}, struct {
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
}{0, 0, 37, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 82, 0, 0}, [2]byte{}}}, struct {
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
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 32, 0, 0}}}, struct {
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
}{0, 0, 76, 0, 0}, [2]byte{}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 32, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 3, 33, 0, 0}}}, struct {
	F0 anon_3
	F1 [6]byte
}{anon_3{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_2{1, 4, 33, 0, 0}}}, struct {
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
}{0, 0, 87, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 63, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 64, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 30, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 79, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 90, 0, 0}, [2]byte{}}}}
var _str_6 [4]byte = [4]byte{101, 110, 100, 0}
var _str_7 [3]byte = [3]byte{60, 33, 0}
var _str_8 [15]byte = [15]byte{100, 111, 99, 116, 121, 112, 101, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_9 [2]byte = [2]byte{62, 0}
var _str_10 [8]byte = [8]byte{100, 111, 99, 116, 121, 112, 101, 0}
var _str_11 [2]byte = [2]byte{60, 0}
var _str_12 [3]byte = [3]byte{47, 62, 0}
var _str_13 [3]byte = [3]byte{60, 47, 0}
var _str_14 [2]byte = [2]byte{61, 0}
var _str_15 [15]byte = [15]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 110, 97, 109, 101, 0}
var _str_16 [16]byte = [16]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 118, 97, 108, 117, 101, 0}
var _str_17 [7]byte = [7]byte{101, 110, 116, 105, 116, 121, 0}
var _str_18 [2]byte = [2]byte{39, 0}
var _str_19 [2]byte = [2]byte{34, 0}
var _str_20 [5]byte = [5]byte{116, 101, 120, 116, 0}
var _str_21 [9]byte = [9]byte{116, 97, 103, 95, 110, 97, 109, 101, 0}
var _str_22 [23]byte = [23]byte{101, 114, 114, 111, 110, 101, 111, 117, 115, 95, 101, 110, 100, 95, 116, 97, 103, 95, 110, 97, 109, 101, 0}
var _str_23 [18]byte = [18]byte{95, 105, 109, 112, 108, 105, 99, 105, 116, 95, 101, 110, 100, 95, 116, 97, 103, 0}
var _str_24 [9]byte = [9]byte{114, 97, 119, 95, 116, 101, 120, 116, 0}
var _str_25 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_26 [9]byte = [9]byte{100, 111, 99, 117, 109, 101, 110, 116, 0}
var _str_27 [6]byte = [6]byte{95, 110, 111, 100, 101, 0}
var _str_28 [8]byte = [8]byte{101, 108, 101, 109, 101, 110, 116, 0}
var _str_29 [15]byte = [15]byte{115, 99, 114, 105, 112, 116, 95, 101, 108, 101, 109, 101, 110, 116, 0}
var _str_30 [14]byte = [14]byte{115, 116, 121, 108, 101, 95, 101, 108, 101, 109, 101, 110, 116, 0}
var _str_31 [10]byte = [10]byte{115, 116, 97, 114, 116, 95, 116, 97, 103, 0}
var _str_32 [17]byte = [17]byte{115, 101, 108, 102, 95, 99, 108, 111, 115, 105, 110, 103, 95, 116, 97, 103, 0}
var _str_33 [8]byte = [8]byte{101, 110, 100, 95, 116, 97, 103, 0}
var _str_34 [18]byte = [18]byte{101, 114, 114, 111, 110, 101, 111, 117, 115, 95, 101, 110, 100, 95, 116, 97, 103, 0}
var _str_35 [10]byte = [10]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 0}
var _str_36 [23]byte = [23]byte{113, 117, 111, 116, 101, 100, 95, 97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 118, 97, 108, 117, 101, 0}
var _str_37 [17]byte = [17]byte{100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_38 [18]byte = [18]byte{115, 116, 97, 114, 116, 95, 116, 97, 103, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var ts_lex_map [18]int16 = [18]int16{34, 73, 38, 3, 39, 70, 47, 6, 60, 24, 61, 27, 62, 22, 68, 9, 100, 9}

func init() {
	tree_sitter_html_language = struct {
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
	}{14, 41, 0, 25, 9, 94, 2, 1, 0, 4, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, anon_0{libc.Ptr(&ts_external_scanner_states), libc.Ptr(&ts_external_scanner_symbol_map), libc.FuncCode(tree_sitter_html_external_scanner_create), libc.FuncCode(tree_sitter_html_external_scanner_destroy), libc.FuncCode(tree_sitter_html_external_scanner_scan), libc.FuncCode(tree_sitter_html_external_scanner_serialize), libc.FuncCode(tree_sitter_html_external_scanner_deserialize)}, libc.Ptr(&ts_primary_state_ids)}
}
func tree_sitter_html_external_scanner_create() unsafe.Pointer {
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
func tree_sitter_html_external_scanner_scan(payload unsafe.Pointer, lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
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
	var loadedv, loadedv2, loadedv5, call, tobool, cmp, call10, loadedv13, call15, loadedv19, call21, loadedv25, call27, loadedv30, loadedv32, loadedv35, loadedv38, call39, call40, tobool42, v45 bool
	var retval unsafe.Pointer
	var v9, call6, v12, v18, conv, conv41, cond int32
	var lookahead, lookahead7, lookahead8 unsafe.Pointer
	var v1, v3, v5, v22, v26, v30, v34, v36, v38, v40 byte
	var arrayidx, arrayidx1, arrayidx4, arrayidx12, arrayidx18, arrayidx24, arrayidx29, arrayidx31, arrayidx34, arrayidx37 unsafe.Pointer
	var v0, v2, v4, v6, v7, v8, v10, v11, v13, v14, v15, v16, v17, v19, v20, v21, v23, v24, v25, v27, v28, v29, v31, v32, v33, v35, v37, v39, v41, v42, v43, v44 unsafe.Pointer
	var scanner_addr, lexer_addr, valid_symbols_addr, mark_end unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, valid_symbols_addr, v0, arrayidx, v1, loadedv, v2, arrayidx1, v3, loadedv2, v4, arrayidx4, v5, loadedv5, v6, v7, call, v8, lookahead, v9, call6, tobool, v10, v11, lookahead7, v12, v13, mark_end, v14, v15, v16, v17, lookahead8, v18, cmp, v19, v20, call10, v21, arrayidx12, v22, loadedv13, v23, v24, call15, v25, arrayidx18, v26, loadedv19, v27, v28, call21, v29, arrayidx24, v30, loadedv25, v31, v32, call27, v33, arrayidx29, v34, loadedv30, v35, arrayidx31, v36, loadedv32, v37, arrayidx34, v38, loadedv35, v39, arrayidx37, v40, loadedv38, v41, v42, call39, conv, v43, v44, call40, conv41, cond, tobool42, v45

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
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v0), int(int64(7))*1))
	v1 = *libc.As[byte](arrayidx)
	loadedv = (v1 & 1) != 0
	if loadedv {
		goto land_lhs_true
	} else {
		goto if_end
	}

land_lhs_true:
	v2 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx1 = v2
	v3 = *libc.As[byte](arrayidx1)
	loadedv2 = (v3 & 1) != 0
	if loadedv2 {
		goto if_end
	} else {
		goto land_lhs_true3
	}

land_lhs_true3:
	v4 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx4 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v4), int(int64(3))*1))
	v5 = *libc.As[byte](arrayidx4)
	loadedv5 = (v5 & 1) != 0
	if loadedv5 {
		goto if_end
	} else {
		goto if_then
	}

if_then:
	v6 = *libc.As[unsafe.Pointer](scanner_addr)
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	call = scan_raw_text(v6, v7)
	*libc.As[bool](retval) = call
	goto _return

if_end:
	goto while_cond

while_cond:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v8).F0)
	v9 = *libc.As[int32](lookahead)
	call6 = libc.Iswspace(v9)
	tobool = call6 != 0
	if tobool {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	skip(v10)
	goto while_cond

while_end:
	v11 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead7 = libc.Ptr(&libc.As[TSLexer](v11).F0)
	v12 = *libc.As[int32](lookahead7)
	switch v12 {
	case 60:
		goto sw_bb
	case 0:
		goto sw_bb17
	case 47:
		goto sw_bb23
	default:
		goto sw_default
	}

sw_bb:
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v13).F3)
	v14 = *libc.As[unsafe.Pointer](mark_end)
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v14)(v15)
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v16)
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead8 = libc.Ptr(&libc.As[TSLexer](v17).F0)
	v18 = *libc.As[int32](lookahead8)
	cmp = v18 == 33
	if cmp {
		goto if_then9
	} else {
		goto if_end11
	}

if_then9:
	v19 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v19)
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	call10 = scan_comment(v20)
	*libc.As[bool](retval) = call10
	goto _return

if_end11:
	v21 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx12 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v21), int(int64(6))*1))
	v22 = *libc.As[byte](arrayidx12)
	loadedv13 = (v22 & 1) != 0
	if loadedv13 {
		goto if_then14
	} else {
		goto if_end16
	}

if_then14:
	v23 = *libc.As[unsafe.Pointer](scanner_addr)
	v24 = *libc.As[unsafe.Pointer](lexer_addr)
	call15 = scan_implicit_end_tag(v23, v24)
	*libc.As[bool](retval) = call15
	goto _return

if_end16:
	goto sw_epilog

sw_bb17:
	v25 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx18 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v25), int(int64(6))*1))
	v26 = *libc.As[byte](arrayidx18)
	loadedv19 = (v26 & 1) != 0
	if loadedv19 {
		goto if_then20
	} else {
		goto if_end22
	}

if_then20:
	v27 = *libc.As[unsafe.Pointer](scanner_addr)
	v28 = *libc.As[unsafe.Pointer](lexer_addr)
	call21 = scan_implicit_end_tag(v27, v28)
	*libc.As[bool](retval) = call21
	goto _return

if_end22:
	goto sw_epilog

sw_bb23:
	v29 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx24 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v29), int(int64(5))*1))
	v30 = *libc.As[byte](arrayidx24)
	loadedv25 = (v30 & 1) != 0
	if loadedv25 {
		goto if_then26
	} else {
		goto if_end28
	}

if_then26:
	v31 = *libc.As[unsafe.Pointer](scanner_addr)
	v32 = *libc.As[unsafe.Pointer](lexer_addr)
	call27 = scan_self_closing_tag_delimiter(v31, v32)
	*libc.As[bool](retval) = call27
	goto _return

if_end28:
	goto sw_epilog

sw_default:
	v33 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx29 = v33
	v34 = *libc.As[byte](arrayidx29)
	loadedv30 = (v34 & 1) != 0
	if loadedv30 {
		goto land_lhs_true33
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v35 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx31 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v35), int(int64(3))*1))
	v36 = *libc.As[byte](arrayidx31)
	loadedv32 = (v36 & 1) != 0
	if loadedv32 {
		goto land_lhs_true33
	} else {
		goto if_end43
	}

land_lhs_true33:
	v37 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx34 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v37), int(int64(7))*1))
	v38 = *libc.As[byte](arrayidx34)
	loadedv35 = (v38 & 1) != 0
	if loadedv35 {
		goto if_end43
	} else {
		goto if_then36
	}

if_then36:
	v39 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx37 = v39
	v40 = *libc.As[byte](arrayidx37)
	loadedv38 = (v40 & 1) != 0
	if loadedv38 {
		goto cond_true
	} else {
		goto cond_false
	}

cond_true:
	v41 = *libc.As[unsafe.Pointer](scanner_addr)
	v42 = *libc.As[unsafe.Pointer](lexer_addr)
	call39 = scan_start_tag_name(v41, v42)
	if call39 {
		conv = 1
	} else {
		conv = 0
	}
	cond = conv
	goto cond_end

cond_false:
	v43 = *libc.As[unsafe.Pointer](scanner_addr)
	v44 = *libc.As[unsafe.Pointer](lexer_addr)
	call40 = scan_end_tag_name(v43, v44)
	if call40 {
		conv41 = 1
	} else {
		conv41 = 0
	}
	cond = conv41
	goto cond_end

cond_end:
	tobool42 = cond != 0
	*libc.As[bool](retval) = tobool42
	goto _return

if_end43:
	goto sw_epilog

sw_epilog:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v45 = *libc.As[bool](retval)
	return v45
}
func tree_sitter_html_external_scanner_serialize(payload unsafe.Pointer, buffer unsafe.Pointer) int32 {
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
func tree_sitter_html_external_scanner_deserialize(payload unsafe.Pointer, buffer unsafe.Pointer, length int32) {
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
func tree_sitter_html_external_scanner_destroy(payload unsafe.Pointer) {
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
func tree_sitter_html() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_html_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, loadedv22, cmp24, cmp28, cmp32, cmp35, cmp38, cmp42, cmp45, cmp48, loadedv52, cmp54, cmp58, cmp61, cmp64, cmp68, loadedv72, cmp74, cmp78, cmp81, cmp84, cmp87, loadedv91, cmp93, cmp97, cmp100, cmp103, cmp107, loadedv111, cmp113, cmp117, cmp121, cmp125, cmp128, cmp131, cmp135, cmp138, cmp141, cmp144, cmp147, loadedv151, cmp153, loadedv157, cmp159, cmp162, loadedv166, cmp168, cmp171, loadedv175, cmp177, cmp180, loadedv184, cmp186, cmp189, loadedv193, cmp195, cmp198, loadedv202, cmp204, cmp207, cmp211, cmp214, loadedv218, cmp220, cmp223, loadedv227, cmp229, cmp232, cmp235, cmp239, cmp242, cmp245, cmp248, loadedv252, cmp254, cmp257, cmp260, cmp264, cmp267, loadedv271, cmp273, cmp276, cmp279, cmp282, cmp285, cmp288, loadedv292, loadedv294, cmp297, cmp301, cmp305, cmp308, cmp311, cmp315, cmp318, loadedv322, loadedv324, loadedv328, cmp332, cmp335, cmp338, cmp342, cmp345, loadedv349, cmp353, cmp356, loadedv360, loadedv364, loadedv368, cmp372, cmp376, loadedv380, loadedv384, loadedv388, loadedv392, cmp396, cmp399, cmp402, cmp405, cmp408, cmp411, cmp414, cmp417, cmp420, loadedv424, cmp428, cmp431, cmp434, cmp437, cmp440, cmp443, cmp446, cmp449, loadedv453, loadedv457, cmp461, loadedv465, cmp469, cmp473, cmp476, loadedv480, cmp484, cmp488, cmp491, loadedv495, cmp499, cmp503, cmp506, loadedv510, cmp514, cmp518, cmp521, loadedv525, cmp529, cmp533, cmp536, cmp539, cmp542, cmp545, cmp548, loadedv552, cmp556, cmp560, cmp563, cmp566, cmp569, cmp572, cmp575, loadedv579, cmp583, cmp587, cmp590, cmp593, cmp596, cmp599, cmp602, loadedv606, cmp610, cmp614, cmp617, cmp620, cmp623, cmp626, cmp629, loadedv633, cmp637, cmp641, cmp644, cmp647, cmp650, cmp653, cmp656, loadedv660, cmp664, cmp668, cmp671, cmp674, cmp677, loadedv681, cmp685, cmp689, cmp692, cmp695, cmp698, loadedv702, cmp706, cmp710, cmp713, cmp716, cmp719, loadedv723, cmp727, cmp731, cmp734, cmp737, cmp740, loadedv744, cmp748, cmp752, cmp755, cmp758, cmp761, loadedv765, cmp769, cmp773, cmp776, cmp779, cmp782, loadedv786, cmp790, cmp794, cmp797, cmp800, cmp803, loadedv807, cmp811, cmp815, cmp818, cmp821, cmp824, loadedv828, cmp832, cmp836, cmp839, cmp842, cmp845, loadedv849, cmp853, cmp857, cmp860, cmp863, cmp866, loadedv870, cmp874, cmp878, cmp881, cmp884, cmp887, loadedv891, cmp895, cmp899, cmp902, cmp905, cmp908, loadedv912, cmp916, cmp920, cmp923, cmp926, cmp929, loadedv933, cmp937, cmp941, cmp944, cmp947, cmp950, loadedv954, cmp958, cmp962, cmp965, cmp968, cmp971, loadedv975, cmp979, cmp983, cmp986, cmp989, cmp992, loadedv996, cmp1000, cmp1004, cmp1007, cmp1010, cmp1013, loadedv1017, cmp1021, cmp1025, cmp1028, cmp1031, cmp1034, loadedv1038, cmp1042, cmp1046, cmp1049, cmp1052, cmp1055, loadedv1059, cmp1063, cmp1067, cmp1070, cmp1073, cmp1076, loadedv1080, cmp1084, cmp1088, cmp1091, cmp1094, cmp1097, loadedv1101, cmp1105, cmp1109, cmp1112, cmp1115, cmp1118, loadedv1122, cmp1126, cmp1130, cmp1133, cmp1136, cmp1139, loadedv1143, cmp1147, cmp1151, cmp1154, cmp1157, cmp1160, loadedv1164, cmp1168, cmp1172, cmp1175, cmp1178, cmp1181, loadedv1185, cmp1189, cmp1193, cmp1196, cmp1199, cmp1202, loadedv1206, cmp1210, cmp1214, cmp1217, cmp1220, cmp1223, loadedv1227, cmp1231, cmp1235, cmp1238, cmp1241, cmp1244, loadedv1248, cmp1252, cmp1256, cmp1259, cmp1262, cmp1265, loadedv1269, loadedv1273, cmp1277, cmp1280, cmp1283, cmp1287, cmp1290, loadedv1294, cmp1298, cmp1301, loadedv1305, loadedv1309, cmp1313, cmp1316, cmp1319, cmp1323, cmp1326, loadedv1330, cmp1334, cmp1337, loadedv1341, cmp1345, cmp1348, cmp1351, cmp1355, cmp1358, cmp1361, cmp1364, loadedv1368, v651 bool
	var retval unsafe.Pointer
	var v9, v13, v16 int16
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol326, result_symbol330, result_symbol351, result_symbol362, result_symbol366, result_symbol370, result_symbol382, result_symbol386, result_symbol390, result_symbol394, result_symbol426, result_symbol455, result_symbol459, result_symbol467, result_symbol482, result_symbol497, result_symbol512, result_symbol527, result_symbol554, result_symbol581, result_symbol608, result_symbol635, result_symbol662, result_symbol683, result_symbol704, result_symbol725, result_symbol746, result_symbol767, result_symbol788, result_symbol809, result_symbol830, result_symbol851, result_symbol872, result_symbol893, result_symbol914, result_symbol935, result_symbol956, result_symbol977, result_symbol998, result_symbol1019, result_symbol1040, result_symbol1061, result_symbol1082, result_symbol1103, result_symbol1124, result_symbol1145, result_symbol1166, result_symbol1187, result_symbol1208, result_symbol1229, result_symbol1250, result_symbol1271, result_symbol1275, result_symbol1296, result_symbol1307, result_symbol1311, result_symbol1332, result_symbol1343 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v22, v23, v24, v25, v26, v27, v28, v29, v31, v32, v33, v34, v35, v37, v38, v39, v40, v41, v43, v44, v45, v46, v47, v49, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v61, v63, v64, v66, v67, v69, v70, v72, v73, v75, v76, v78, v79, v80, v81, v83, v84, v86, v87, v88, v89, v90, v91, v92, v94, v95, v96, v97, v98, v100, v101, v102, v103, v104, v105, v108, v109, v110, v111, v112, v113, v114, v130, v131, v132, v133, v134, v140, v141, v157, v158, v179, v180, v181, v182, v183, v184, v185, v186, v187, v193, v194, v195, v196, v197, v198, v199, v200, v211, v217, v218, v219, v225, v226, v227, v233, v234, v235, v241, v242, v243, v249, v250, v251, v252, v253, v254, v255, v261, v262, v263, v264, v265, v266, v267, v273, v274, v275, v276, v277, v278, v279, v285, v286, v287, v288, v289, v290, v291, v297, v298, v299, v300, v301, v302, v303, v309, v310, v311, v312, v313, v319, v320, v321, v322, v323, v329, v330, v331, v332, v333, v339, v340, v341, v342, v343, v349, v350, v351, v352, v353, v359, v360, v361, v362, v363, v369, v370, v371, v372, v373, v379, v380, v381, v382, v383, v389, v390, v391, v392, v393, v399, v400, v401, v402, v403, v409, v410, v411, v412, v413, v419, v420, v421, v422, v423, v429, v430, v431, v432, v433, v439, v440, v441, v442, v443, v449, v450, v451, v452, v453, v459, v460, v461, v462, v463, v469, v470, v471, v472, v473, v479, v480, v481, v482, v483, v489, v490, v491, v492, v493, v499, v500, v501, v502, v503, v509, v510, v511, v512, v513, v519, v520, v521, v522, v523, v529, v530, v531, v532, v533, v539, v540, v541, v542, v543, v549, v550, v551, v552, v553, v559, v560, v561, v562, v563, v569, v570, v571, v572, v573, v579, v580, v581, v582, v583, v589, v590, v591, v592, v593, v604, v605, v606, v607, v608, v614, v615, v626, v627, v628, v629, v630, v636, v637, v643, v644, v645, v646, v647, v648, v649 int32
	var lookahead, i, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10 int64
	var v3, storedv, v10, v21, v30, v36, v42, v48, v60, v62, v65, v68, v71, v74, v77, v82, v85, v93, v99, v106, v107, v115, v120, v125, v135, v142, v147, v152, v159, v164, v169, v174, v188, v201, v206, v212, v220, v228, v236, v244, v256, v268, v280, v292, v304, v314, v324, v334, v344, v354, v364, v374, v384, v394, v404, v414, v424, v434, v444, v454, v464, v474, v484, v494, v504, v514, v524, v534, v544, v554, v564, v574, v584, v594, v599, v609, v616, v621, v631, v638, v650 byte
	var result, local_skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v116, v117, v118, v119, v121, v122, v123, v124, v126, v127, v128, v129, v136, v137, v138, v139, v143, v144, v145, v146, v148, v149, v150, v151, v153, v154, v155, v156, v160, v161, v162, v163, v165, v166, v167, v168, v170, v171, v172, v173, v175, v176, v177, v178, v189, v190, v191, v192, v202, v203, v204, v205, v207, v208, v209, v210, v213, v214, v215, v216, v221, v222, v223, v224, v229, v230, v231, v232, v237, v238, v239, v240, v245, v246, v247, v248, v257, v258, v259, v260, v269, v270, v271, v272, v281, v282, v283, v284, v293, v294, v295, v296, v305, v306, v307, v308, v315, v316, v317, v318, v325, v326, v327, v328, v335, v336, v337, v338, v345, v346, v347, v348, v355, v356, v357, v358, v365, v366, v367, v368, v375, v376, v377, v378, v385, v386, v387, v388, v395, v396, v397, v398, v405, v406, v407, v408, v415, v416, v417, v418, v425, v426, v427, v428, v435, v436, v437, v438, v445, v446, v447, v448, v455, v456, v457, v458, v465, v466, v467, v468, v475, v476, v477, v478, v485, v486, v487, v488, v495, v496, v497, v498, v505, v506, v507, v508, v515, v516, v517, v518, v525, v526, v527, v528, v535, v536, v537, v538, v545, v546, v547, v548, v555, v556, v557, v558, v565, v566, v567, v568, v575, v576, v577, v578, v585, v586, v587, v588, v595, v596, v597, v598, v600, v601, v602, v603, v610, v611, v612, v613, v617, v618, v619, v620, v622, v623, v624, v625, v632, v633, v634, v635, v639, v640, v641, v642 unsafe.Pointer
	var lexer_addr, local_advance, eof2, mark_end, mark_end327, mark_end331, mark_end352, mark_end363, mark_end367, mark_end371, mark_end383, mark_end387, mark_end391, mark_end395, mark_end427, mark_end456, mark_end460, mark_end468, mark_end483, mark_end498, mark_end513, mark_end528, mark_end555, mark_end582, mark_end609, mark_end636, mark_end663, mark_end684, mark_end705, mark_end726, mark_end747, mark_end768, mark_end789, mark_end810, mark_end831, mark_end852, mark_end873, mark_end894, mark_end915, mark_end936, mark_end957, mark_end978, mark_end999, mark_end1020, mark_end1041, mark_end1062, mark_end1083, mark_end1104, mark_end1125, mark_end1146, mark_end1167, mark_end1188, mark_end1209, mark_end1230, mark_end1251, mark_end1272, mark_end1276, mark_end1297, mark_end1308, mark_end1312, mark_end1333, mark_end1344 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, local_skip, eof, lookahead, i, v0, local_advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, loadedv22, v22, cmp24, v23, cmp28, v24, cmp32, v25, cmp35, v26, cmp38, v27, cmp42, v28, cmp45, v29, cmp48, v30, loadedv52, v31, cmp54, v32, cmp58, v33, cmp61, v34, cmp64, v35, cmp68, v36, loadedv72, v37, cmp74, v38, cmp78, v39, cmp81, v40, cmp84, v41, cmp87, v42, loadedv91, v43, cmp93, v44, cmp97, v45, cmp100, v46, cmp103, v47, cmp107, v48, loadedv111, v49, cmp113, v50, cmp117, v51, cmp121, v52, cmp125, v53, cmp128, v54, cmp131, v55, cmp135, v56, cmp138, v57, cmp141, v58, cmp144, v59, cmp147, v60, loadedv151, v61, cmp153, v62, loadedv157, v63, cmp159, v64, cmp162, v65, loadedv166, v66, cmp168, v67, cmp171, v68, loadedv175, v69, cmp177, v70, cmp180, v71, loadedv184, v72, cmp186, v73, cmp189, v74, loadedv193, v75, cmp195, v76, cmp198, v77, loadedv202, v78, cmp204, v79, cmp207, v80, cmp211, v81, cmp214, v82, loadedv218, v83, cmp220, v84, cmp223, v85, loadedv227, v86, cmp229, v87, cmp232, v88, cmp235, v89, cmp239, v90, cmp242, v91, cmp245, v92, cmp248, v93, loadedv252, v94, cmp254, v95, cmp257, v96, cmp260, v97, cmp264, v98, cmp267, v99, loadedv271, v100, cmp273, v101, cmp276, v102, cmp279, v103, cmp282, v104, cmp285, v105, cmp288, v106, loadedv292, v107, loadedv294, v108, cmp297, v109, cmp301, v110, cmp305, v111, cmp308, v112, cmp311, v113, cmp315, v114, cmp318, v115, loadedv322, v116, result_symbol, v117, mark_end, v118, v119, v120, loadedv324, v121, result_symbol326, v122, mark_end327, v123, v124, v125, loadedv328, v126, result_symbol330, v127, mark_end331, v128, v129, v130, cmp332, v131, cmp335, v132, cmp338, v133, cmp342, v134, cmp345, v135, loadedv349, v136, result_symbol351, v137, mark_end352, v138, v139, v140, cmp353, v141, cmp356, v142, loadedv360, v143, result_symbol362, v144, mark_end363, v145, v146, v147, loadedv364, v148, result_symbol366, v149, mark_end367, v150, v151, v152, loadedv368, v153, result_symbol370, v154, mark_end371, v155, v156, v157, cmp372, v158, cmp376, v159, loadedv380, v160, result_symbol382, v161, mark_end383, v162, v163, v164, loadedv384, v165, result_symbol386, v166, mark_end387, v167, v168, v169, loadedv388, v170, result_symbol390, v171, mark_end391, v172, v173, v174, loadedv392, v175, result_symbol394, v176, mark_end395, v177, v178, v179, cmp396, v180, cmp399, v181, cmp402, v182, cmp405, v183, cmp408, v184, cmp411, v185, cmp414, v186, cmp417, v187, cmp420, v188, loadedv424, v189, result_symbol426, v190, mark_end427, v191, v192, v193, cmp428, v194, cmp431, v195, cmp434, v196, cmp437, v197, cmp440, v198, cmp443, v199, cmp446, v200, cmp449, v201, loadedv453, v202, result_symbol455, v203, mark_end456, v204, v205, v206, loadedv457, v207, result_symbol459, v208, mark_end460, v209, v210, v211, cmp461, v212, loadedv465, v213, result_symbol467, v214, mark_end468, v215, v216, v217, cmp469, v218, cmp473, v219, cmp476, v220, loadedv480, v221, result_symbol482, v222, mark_end483, v223, v224, v225, cmp484, v226, cmp488, v227, cmp491, v228, loadedv495, v229, result_symbol497, v230, mark_end498, v231, v232, v233, cmp499, v234, cmp503, v235, cmp506, v236, loadedv510, v237, result_symbol512, v238, mark_end513, v239, v240, v241, cmp514, v242, cmp518, v243, cmp521, v244, loadedv525, v245, result_symbol527, v246, mark_end528, v247, v248, v249, cmp529, v250, cmp533, v251, cmp536, v252, cmp539, v253, cmp542, v254, cmp545, v255, cmp548, v256, loadedv552, v257, result_symbol554, v258, mark_end555, v259, v260, v261, cmp556, v262, cmp560, v263, cmp563, v264, cmp566, v265, cmp569, v266, cmp572, v267, cmp575, v268, loadedv579, v269, result_symbol581, v270, mark_end582, v271, v272, v273, cmp583, v274, cmp587, v275, cmp590, v276, cmp593, v277, cmp596, v278, cmp599, v279, cmp602, v280, loadedv606, v281, result_symbol608, v282, mark_end609, v283, v284, v285, cmp610, v286, cmp614, v287, cmp617, v288, cmp620, v289, cmp623, v290, cmp626, v291, cmp629, v292, loadedv633, v293, result_symbol635, v294, mark_end636, v295, v296, v297, cmp637, v298, cmp641, v299, cmp644, v300, cmp647, v301, cmp650, v302, cmp653, v303, cmp656, v304, loadedv660, v305, result_symbol662, v306, mark_end663, v307, v308, v309, cmp664, v310, cmp668, v311, cmp671, v312, cmp674, v313, cmp677, v314, loadedv681, v315, result_symbol683, v316, mark_end684, v317, v318, v319, cmp685, v320, cmp689, v321, cmp692, v322, cmp695, v323, cmp698, v324, loadedv702, v325, result_symbol704, v326, mark_end705, v327, v328, v329, cmp706, v330, cmp710, v331, cmp713, v332, cmp716, v333, cmp719, v334, loadedv723, v335, result_symbol725, v336, mark_end726, v337, v338, v339, cmp727, v340, cmp731, v341, cmp734, v342, cmp737, v343, cmp740, v344, loadedv744, v345, result_symbol746, v346, mark_end747, v347, v348, v349, cmp748, v350, cmp752, v351, cmp755, v352, cmp758, v353, cmp761, v354, loadedv765, v355, result_symbol767, v356, mark_end768, v357, v358, v359, cmp769, v360, cmp773, v361, cmp776, v362, cmp779, v363, cmp782, v364, loadedv786, v365, result_symbol788, v366, mark_end789, v367, v368, v369, cmp790, v370, cmp794, v371, cmp797, v372, cmp800, v373, cmp803, v374, loadedv807, v375, result_symbol809, v376, mark_end810, v377, v378, v379, cmp811, v380, cmp815, v381, cmp818, v382, cmp821, v383, cmp824, v384, loadedv828, v385, result_symbol830, v386, mark_end831, v387, v388, v389, cmp832, v390, cmp836, v391, cmp839, v392, cmp842, v393, cmp845, v394, loadedv849, v395, result_symbol851, v396, mark_end852, v397, v398, v399, cmp853, v400, cmp857, v401, cmp860, v402, cmp863, v403, cmp866, v404, loadedv870, v405, result_symbol872, v406, mark_end873, v407, v408, v409, cmp874, v410, cmp878, v411, cmp881, v412, cmp884, v413, cmp887, v414, loadedv891, v415, result_symbol893, v416, mark_end894, v417, v418, v419, cmp895, v420, cmp899, v421, cmp902, v422, cmp905, v423, cmp908, v424, loadedv912, v425, result_symbol914, v426, mark_end915, v427, v428, v429, cmp916, v430, cmp920, v431, cmp923, v432, cmp926, v433, cmp929, v434, loadedv933, v435, result_symbol935, v436, mark_end936, v437, v438, v439, cmp937, v440, cmp941, v441, cmp944, v442, cmp947, v443, cmp950, v444, loadedv954, v445, result_symbol956, v446, mark_end957, v447, v448, v449, cmp958, v450, cmp962, v451, cmp965, v452, cmp968, v453, cmp971, v454, loadedv975, v455, result_symbol977, v456, mark_end978, v457, v458, v459, cmp979, v460, cmp983, v461, cmp986, v462, cmp989, v463, cmp992, v464, loadedv996, v465, result_symbol998, v466, mark_end999, v467, v468, v469, cmp1000, v470, cmp1004, v471, cmp1007, v472, cmp1010, v473, cmp1013, v474, loadedv1017, v475, result_symbol1019, v476, mark_end1020, v477, v478, v479, cmp1021, v480, cmp1025, v481, cmp1028, v482, cmp1031, v483, cmp1034, v484, loadedv1038, v485, result_symbol1040, v486, mark_end1041, v487, v488, v489, cmp1042, v490, cmp1046, v491, cmp1049, v492, cmp1052, v493, cmp1055, v494, loadedv1059, v495, result_symbol1061, v496, mark_end1062, v497, v498, v499, cmp1063, v500, cmp1067, v501, cmp1070, v502, cmp1073, v503, cmp1076, v504, loadedv1080, v505, result_symbol1082, v506, mark_end1083, v507, v508, v509, cmp1084, v510, cmp1088, v511, cmp1091, v512, cmp1094, v513, cmp1097, v514, loadedv1101, v515, result_symbol1103, v516, mark_end1104, v517, v518, v519, cmp1105, v520, cmp1109, v521, cmp1112, v522, cmp1115, v523, cmp1118, v524, loadedv1122, v525, result_symbol1124, v526, mark_end1125, v527, v528, v529, cmp1126, v530, cmp1130, v531, cmp1133, v532, cmp1136, v533, cmp1139, v534, loadedv1143, v535, result_symbol1145, v536, mark_end1146, v537, v538, v539, cmp1147, v540, cmp1151, v541, cmp1154, v542, cmp1157, v543, cmp1160, v544, loadedv1164, v545, result_symbol1166, v546, mark_end1167, v547, v548, v549, cmp1168, v550, cmp1172, v551, cmp1175, v552, cmp1178, v553, cmp1181, v554, loadedv1185, v555, result_symbol1187, v556, mark_end1188, v557, v558, v559, cmp1189, v560, cmp1193, v561, cmp1196, v562, cmp1199, v563, cmp1202, v564, loadedv1206, v565, result_symbol1208, v566, mark_end1209, v567, v568, v569, cmp1210, v570, cmp1214, v571, cmp1217, v572, cmp1220, v573, cmp1223, v574, loadedv1227, v575, result_symbol1229, v576, mark_end1230, v577, v578, v579, cmp1231, v580, cmp1235, v581, cmp1238, v582, cmp1241, v583, cmp1244, v584, loadedv1248, v585, result_symbol1250, v586, mark_end1251, v587, v588, v589, cmp1252, v590, cmp1256, v591, cmp1259, v592, cmp1262, v593, cmp1265, v594, loadedv1269, v595, result_symbol1271, v596, mark_end1272, v597, v598, v599, loadedv1273, v600, result_symbol1275, v601, mark_end1276, v602, v603, v604, cmp1277, v605, cmp1280, v606, cmp1283, v607, cmp1287, v608, cmp1290, v609, loadedv1294, v610, result_symbol1296, v611, mark_end1297, v612, v613, v614, cmp1298, v615, cmp1301, v616, loadedv1305, v617, result_symbol1307, v618, mark_end1308, v619, v620, v621, loadedv1309, v622, result_symbol1311, v623, mark_end1312, v624, v625, v626, cmp1313, v627, cmp1316, v628, cmp1319, v629, cmp1323, v630, cmp1326, v631, loadedv1330, v632, result_symbol1332, v633, mark_end1333, v634, v635, v636, cmp1334, v637, cmp1337, v638, loadedv1341, v639, result_symbol1343, v640, mark_end1344, v641, v642, v643, cmp1345, v644, cmp1348, v645, cmp1351, v646, cmp1355, v647, cmp1358, v648, cmp1361, v649, cmp1364, v650, loadedv1368, v651

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
	i = libc.Ptr(&new(struct {
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
		goto sw_bb23
	case 2:
		goto sw_bb53
	case 3:
		goto sw_bb73
	case 4:
		goto sw_bb92
	case 5:
		goto sw_bb112
	case 6:
		goto sw_bb152
	case 7:
		goto sw_bb158
	case 8:
		goto sw_bb167
	case 9:
		goto sw_bb176
	case 10:
		goto sw_bb185
	case 11:
		goto sw_bb194
	case 12:
		goto sw_bb203
	case 13:
		goto sw_bb219
	case 14:
		goto sw_bb228
	case 15:
		goto sw_bb253
	case 16:
		goto sw_bb272
	case 17:
		goto sw_bb293
	case 18:
		goto sw_bb323
	case 19:
		goto sw_bb325
	case 20:
		goto sw_bb329
	case 21:
		goto sw_bb350
	case 22:
		goto sw_bb361
	case 23:
		goto sw_bb365
	case 24:
		goto sw_bb369
	case 25:
		goto sw_bb381
	case 26:
		goto sw_bb385
	case 27:
		goto sw_bb389
	case 28:
		goto sw_bb393
	case 29:
		goto sw_bb425
	case 30:
		goto sw_bb454
	case 31:
		goto sw_bb458
	case 32:
		goto sw_bb466
	case 33:
		goto sw_bb481
	case 34:
		goto sw_bb496
	case 35:
		goto sw_bb511
	case 36:
		goto sw_bb526
	case 37:
		goto sw_bb553
	case 38:
		goto sw_bb580
	case 39:
		goto sw_bb607
	case 40:
		goto sw_bb634
	case 41:
		goto sw_bb661
	case 42:
		goto sw_bb682
	case 43:
		goto sw_bb703
	case 44:
		goto sw_bb724
	case 45:
		goto sw_bb745
	case 46:
		goto sw_bb766
	case 47:
		goto sw_bb787
	case 48:
		goto sw_bb808
	case 49:
		goto sw_bb829
	case 50:
		goto sw_bb850
	case 51:
		goto sw_bb871
	case 52:
		goto sw_bb892
	case 53:
		goto sw_bb913
	case 54:
		goto sw_bb934
	case 55:
		goto sw_bb955
	case 56:
		goto sw_bb976
	case 57:
		goto sw_bb997
	case 58:
		goto sw_bb1018
	case 59:
		goto sw_bb1039
	case 60:
		goto sw_bb1060
	case 61:
		goto sw_bb1081
	case 62:
		goto sw_bb1102
	case 63:
		goto sw_bb1123
	case 64:
		goto sw_bb1144
	case 65:
		goto sw_bb1165
	case 66:
		goto sw_bb1186
	case 67:
		goto sw_bb1207
	case 68:
		goto sw_bb1228
	case 69:
		goto sw_bb1249
	case 70:
		goto sw_bb1270
	case 71:
		goto sw_bb1274
	case 72:
		goto sw_bb1295
	case 73:
		goto sw_bb1306
	case 74:
		goto sw_bb1310
	case 75:
		goto sw_bb1331
	case 76:
		goto sw_bb1342
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
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v11 = *libc.As[int32](i)
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(18)
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v12 = *libc.As[int32](i)
	idxprom = int64(uint64(uint32(v12)))
	arrayidx = libc.Ptr(&ts_lex_map[idxprom])
	v13 = *libc.As[int16](arrayidx)
	conv6 = int32(uint32(uint16(v13)))
	v14 = *libc.As[int32](lookahead)
	cmp7 = conv6 == v14
	if cmp7 {
		goto if_then9
	} else {
		goto if_end12
	}

if_then9:
	v15 = *libc.As[int32](i)
	add = v15 + 1
	idxprom10 = int64(uint64(uint32(add)))
	arrayidx11 = libc.Ptr(&ts_lex_map[idxprom10])
	v16 = *libc.As[int16](arrayidx11)
	*libc.As[int16](state_addr) = v16
	goto next_state

if_end12:
	goto for_inc

for_inc:
	v17 = *libc.As[int32](i)
	add13 = v17 + 2
	*libc.As[int32](i) = add13
	goto for_cond

for_end:
	v18 = *libc.As[int32](lookahead)
	cmp14 = 9 <= v18
	if cmp14 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false
	}

land_lhs_true:
	v19 = *libc.As[int32](lookahead)
	cmp16 = v19 <= 13
	if cmp16 {
		goto if_then20
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v20 = *libc.As[int32](lookahead)
	cmp18 = v20 == 32
	if cmp18 {
		goto if_then20
	} else {
		goto if_end21
	}

if_then20:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end21:
	v21 = *libc.As[byte](result)
	loadedv22 = (v21 & 1) != 0
	*libc.As[bool](retval) = loadedv22
	goto _return

sw_bb23:
	v22 = *libc.As[int32](lookahead)
	cmp24 = v22 == 34
	if cmp24 {
		goto if_then26
	} else {
		goto if_end27
	}

if_then26:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end27:
	v23 = *libc.As[int32](lookahead)
	cmp28 = v23 == 39
	if cmp28 {
		goto if_then30
	} else {
		goto if_end31
	}

if_then30:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end31:
	v24 = *libc.As[int32](lookahead)
	cmp32 = 9 <= v24
	if cmp32 {
		goto land_lhs_true34
	} else {
		goto lor_lhs_false37
	}

land_lhs_true34:
	v25 = *libc.As[int32](lookahead)
	cmp35 = v25 <= 13
	if cmp35 {
		goto if_then40
	} else {
		goto lor_lhs_false37
	}

lor_lhs_false37:
	v26 = *libc.As[int32](lookahead)
	cmp38 = v26 == 32
	if cmp38 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end41:
	v27 = *libc.As[int32](lookahead)
	cmp42 = v27 != 0
	if cmp42 {
		goto land_lhs_true44
	} else {
		goto if_end51
	}

land_lhs_true44:
	v28 = *libc.As[int32](lookahead)
	cmp45 = v28 < 60
	if cmp45 {
		goto if_then50
	} else {
		goto lor_lhs_false47
	}

lor_lhs_false47:
	v29 = *libc.As[int32](lookahead)
	cmp48 = 62 < v29
	if cmp48 {
		goto if_then50
	} else {
		goto if_end51
	}

if_then50:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end51:
	v30 = *libc.As[byte](result)
	loadedv52 = (v30 & 1) != 0
	*libc.As[bool](retval) = loadedv52
	goto _return

sw_bb53:
	v31 = *libc.As[int32](lookahead)
	cmp54 = v31 == 34
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end57:
	v32 = *libc.As[int32](lookahead)
	cmp58 = 9 <= v32
	if cmp58 {
		goto land_lhs_true60
	} else {
		goto lor_lhs_false63
	}

land_lhs_true60:
	v33 = *libc.As[int32](lookahead)
	cmp61 = v33 <= 13
	if cmp61 {
		goto if_then66
	} else {
		goto lor_lhs_false63
	}

lor_lhs_false63:
	v34 = *libc.As[int32](lookahead)
	cmp64 = v34 == 32
	if cmp64 {
		goto if_then66
	} else {
		goto if_end67
	}

if_then66:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end67:
	v35 = *libc.As[int32](lookahead)
	cmp68 = v35 != 0
	if cmp68 {
		goto if_then70
	} else {
		goto if_end71
	}

if_then70:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end71:
	v36 = *libc.As[byte](result)
	loadedv72 = (v36 & 1) != 0
	*libc.As[bool](retval) = loadedv72
	goto _return

sw_bb73:
	v37 = *libc.As[int32](lookahead)
	cmp74 = v37 == 35
	if cmp74 {
		goto if_then76
	} else {
		goto if_end77
	}

if_then76:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end77:
	v38 = *libc.As[int32](lookahead)
	cmp78 = 65 <= v38
	if cmp78 {
		goto land_lhs_true80
	} else {
		goto lor_lhs_false83
	}

land_lhs_true80:
	v39 = *libc.As[int32](lookahead)
	cmp81 = v39 <= 90
	if cmp81 {
		goto if_then89
	} else {
		goto lor_lhs_false83
	}

lor_lhs_false83:
	v40 = *libc.As[int32](lookahead)
	cmp84 = 97 <= v40
	if cmp84 {
		goto land_lhs_true86
	} else {
		goto if_end90
	}

land_lhs_true86:
	v41 = *libc.As[int32](lookahead)
	cmp87 = v41 <= 122
	if cmp87 {
		goto if_then89
	} else {
		goto if_end90
	}

if_then89:
	*libc.As[int16](state_addr) = 69
	goto next_state

if_end90:
	v42 = *libc.As[byte](result)
	loadedv91 = (v42 & 1) != 0
	*libc.As[bool](retval) = loadedv91
	goto _return

sw_bb92:
	v43 = *libc.As[int32](lookahead)
	cmp93 = v43 == 39
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end96:
	v44 = *libc.As[int32](lookahead)
	cmp97 = 9 <= v44
	if cmp97 {
		goto land_lhs_true99
	} else {
		goto lor_lhs_false102
	}

land_lhs_true99:
	v45 = *libc.As[int32](lookahead)
	cmp100 = v45 <= 13
	if cmp100 {
		goto if_then105
	} else {
		goto lor_lhs_false102
	}

lor_lhs_false102:
	v46 = *libc.As[int32](lookahead)
	cmp103 = v46 == 32
	if cmp103 {
		goto if_then105
	} else {
		goto if_end106
	}

if_then105:
	*libc.As[int16](state_addr) = 71
	goto next_state

if_end106:
	v47 = *libc.As[int32](lookahead)
	cmp107 = v47 != 0
	if cmp107 {
		goto if_then109
	} else {
		goto if_end110
	}

if_then109:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end110:
	v48 = *libc.As[byte](result)
	loadedv111 = (v48 & 1) != 0
	*libc.As[bool](retval) = loadedv111
	goto _return

sw_bb112:
	v49 = *libc.As[int32](lookahead)
	cmp113 = v49 == 47
	if cmp113 {
		goto if_then115
	} else {
		goto if_end116
	}

if_then115:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end116:
	v50 = *libc.As[int32](lookahead)
	cmp117 = v50 == 61
	if cmp117 {
		goto if_then119
	} else {
		goto if_end120
	}

if_then119:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end120:
	v51 = *libc.As[int32](lookahead)
	cmp121 = v51 == 62
	if cmp121 {
		goto if_then123
	} else {
		goto if_end124
	}

if_then123:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end124:
	v52 = *libc.As[int32](lookahead)
	cmp125 = 9 <= v52
	if cmp125 {
		goto land_lhs_true127
	} else {
		goto lor_lhs_false130
	}

land_lhs_true127:
	v53 = *libc.As[int32](lookahead)
	cmp128 = v53 <= 13
	if cmp128 {
		goto if_then133
	} else {
		goto lor_lhs_false130
	}

lor_lhs_false130:
	v54 = *libc.As[int32](lookahead)
	cmp131 = v54 == 32
	if cmp131 {
		goto if_then133
	} else {
		goto if_end134
	}

if_then133:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end134:
	v55 = *libc.As[int32](lookahead)
	cmp135 = v55 != 0
	if cmp135 {
		goto land_lhs_true137
	} else {
		goto if_end150
	}

land_lhs_true137:
	v56 = *libc.As[int32](lookahead)
	cmp138 = v56 != 34
	if cmp138 {
		goto land_lhs_true140
	} else {
		goto if_end150
	}

land_lhs_true140:
	v57 = *libc.As[int32](lookahead)
	cmp141 = v57 != 39
	if cmp141 {
		goto land_lhs_true143
	} else {
		goto if_end150
	}

land_lhs_true143:
	v58 = *libc.As[int32](lookahead)
	cmp144 = v58 < 60
	if cmp144 {
		goto if_then149
	} else {
		goto lor_lhs_false146
	}

lor_lhs_false146:
	v59 = *libc.As[int32](lookahead)
	cmp147 = 62 < v59
	if cmp147 {
		goto if_then149
	} else {
		goto if_end150
	}

if_then149:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end150:
	v60 = *libc.As[byte](result)
	loadedv151 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv151
	goto _return

sw_bb152:
	v61 = *libc.As[int32](lookahead)
	cmp153 = v61 == 62
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end156:
	v62 = *libc.As[byte](result)
	loadedv157 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv157
	goto _return

sw_bb158:
	v63 = *libc.As[int32](lookahead)
	cmp159 = v63 == 67
	if cmp159 {
		goto if_then164
	} else {
		goto lor_lhs_false161
	}

lor_lhs_false161:
	v64 = *libc.As[int32](lookahead)
	cmp162 = v64 == 99
	if cmp162 {
		goto if_then164
	} else {
		goto if_end165
	}

if_then164:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end165:
	v65 = *libc.As[byte](result)
	loadedv166 = (v65 & 1) != 0
	*libc.As[bool](retval) = loadedv166
	goto _return

sw_bb167:
	v66 = *libc.As[int32](lookahead)
	cmp168 = v66 == 69
	if cmp168 {
		goto if_then173
	} else {
		goto lor_lhs_false170
	}

lor_lhs_false170:
	v67 = *libc.As[int32](lookahead)
	cmp171 = v67 == 101
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end174:
	v68 = *libc.As[byte](result)
	loadedv175 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv175
	goto _return

sw_bb176:
	v69 = *libc.As[int32](lookahead)
	cmp177 = v69 == 79
	if cmp177 {
		goto if_then182
	} else {
		goto lor_lhs_false179
	}

lor_lhs_false179:
	v70 = *libc.As[int32](lookahead)
	cmp180 = v70 == 111
	if cmp180 {
		goto if_then182
	} else {
		goto if_end183
	}

if_then182:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end183:
	v71 = *libc.As[byte](result)
	loadedv184 = (v71 & 1) != 0
	*libc.As[bool](retval) = loadedv184
	goto _return

sw_bb185:
	v72 = *libc.As[int32](lookahead)
	cmp186 = v72 == 80
	if cmp186 {
		goto if_then191
	} else {
		goto lor_lhs_false188
	}

lor_lhs_false188:
	v73 = *libc.As[int32](lookahead)
	cmp189 = v73 == 112
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end192:
	v74 = *libc.As[byte](result)
	loadedv193 = (v74 & 1) != 0
	*libc.As[bool](retval) = loadedv193
	goto _return

sw_bb194:
	v75 = *libc.As[int32](lookahead)
	cmp195 = v75 == 84
	if cmp195 {
		goto if_then200
	} else {
		goto lor_lhs_false197
	}

lor_lhs_false197:
	v76 = *libc.As[int32](lookahead)
	cmp198 = v76 == 116
	if cmp198 {
		goto if_then200
	} else {
		goto if_end201
	}

if_then200:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end201:
	v77 = *libc.As[byte](result)
	loadedv202 = (v77 & 1) != 0
	*libc.As[bool](retval) = loadedv202
	goto _return

sw_bb203:
	v78 = *libc.As[int32](lookahead)
	cmp204 = v78 == 88
	if cmp204 {
		goto if_then209
	} else {
		goto lor_lhs_false206
	}

lor_lhs_false206:
	v79 = *libc.As[int32](lookahead)
	cmp207 = v79 == 120
	if cmp207 {
		goto if_then209
	} else {
		goto if_end210
	}

if_then209:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end210:
	v80 = *libc.As[int32](lookahead)
	cmp211 = 48 <= v80
	if cmp211 {
		goto land_lhs_true213
	} else {
		goto if_end217
	}

land_lhs_true213:
	v81 = *libc.As[int32](lookahead)
	cmp214 = v81 <= 57
	if cmp214 {
		goto if_then216
	} else {
		goto if_end217
	}

if_then216:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end217:
	v82 = *libc.As[byte](result)
	loadedv218 = (v82 & 1) != 0
	*libc.As[bool](retval) = loadedv218
	goto _return

sw_bb219:
	v83 = *libc.As[int32](lookahead)
	cmp220 = v83 == 89
	if cmp220 {
		goto if_then225
	} else {
		goto lor_lhs_false222
	}

lor_lhs_false222:
	v84 = *libc.As[int32](lookahead)
	cmp223 = v84 == 121
	if cmp223 {
		goto if_then225
	} else {
		goto if_end226
	}

if_then225:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end226:
	v85 = *libc.As[byte](result)
	loadedv227 = (v85 & 1) != 0
	*libc.As[bool](retval) = loadedv227
	goto _return

sw_bb228:
	v86 = *libc.As[int32](lookahead)
	cmp229 = 9 <= v86
	if cmp229 {
		goto land_lhs_true231
	} else {
		goto lor_lhs_false234
	}

land_lhs_true231:
	v87 = *libc.As[int32](lookahead)
	cmp232 = v87 <= 13
	if cmp232 {
		goto if_then237
	} else {
		goto lor_lhs_false234
	}

lor_lhs_false234:
	v88 = *libc.As[int32](lookahead)
	cmp235 = v88 == 32
	if cmp235 {
		goto if_then237
	} else {
		goto if_end238
	}

if_then237:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end238:
	v89 = *libc.As[int32](lookahead)
	cmp239 = v89 != 0
	if cmp239 {
		goto land_lhs_true241
	} else {
		goto if_end251
	}

land_lhs_true241:
	v90 = *libc.As[int32](lookahead)
	cmp242 = v90 != 38
	if cmp242 {
		goto land_lhs_true244
	} else {
		goto if_end251
	}

land_lhs_true244:
	v91 = *libc.As[int32](lookahead)
	cmp245 = v91 != 60
	if cmp245 {
		goto land_lhs_true247
	} else {
		goto if_end251
	}

land_lhs_true247:
	v92 = *libc.As[int32](lookahead)
	cmp248 = v92 != 62
	if cmp248 {
		goto if_then250
	} else {
		goto if_end251
	}

if_then250:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end251:
	v93 = *libc.As[byte](result)
	loadedv252 = (v93 & 1) != 0
	*libc.As[bool](retval) = loadedv252
	goto _return

sw_bb253:
	v94 = *libc.As[int32](lookahead)
	cmp254 = 9 <= v94
	if cmp254 {
		goto land_lhs_true256
	} else {
		goto lor_lhs_false259
	}

land_lhs_true256:
	v95 = *libc.As[int32](lookahead)
	cmp257 = v95 <= 13
	if cmp257 {
		goto if_then262
	} else {
		goto lor_lhs_false259
	}

lor_lhs_false259:
	v96 = *libc.As[int32](lookahead)
	cmp260 = v96 == 32
	if cmp260 {
		goto if_then262
	} else {
		goto if_end263
	}

if_then262:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end263:
	v97 = *libc.As[int32](lookahead)
	cmp264 = v97 != 0
	if cmp264 {
		goto land_lhs_true266
	} else {
		goto if_end270
	}

land_lhs_true266:
	v98 = *libc.As[int32](lookahead)
	cmp267 = v98 != 62
	if cmp267 {
		goto if_then269
	} else {
		goto if_end270
	}

if_then269:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end270:
	v99 = *libc.As[byte](result)
	loadedv271 = (v99 & 1) != 0
	*libc.As[bool](retval) = loadedv271
	goto _return

sw_bb272:
	v100 = *libc.As[int32](lookahead)
	cmp273 = 48 <= v100
	if cmp273 {
		goto land_lhs_true275
	} else {
		goto lor_lhs_false278
	}

land_lhs_true275:
	v101 = *libc.As[int32](lookahead)
	cmp276 = v101 <= 57
	if cmp276 {
		goto if_then290
	} else {
		goto lor_lhs_false278
	}

lor_lhs_false278:
	v102 = *libc.As[int32](lookahead)
	cmp279 = 65 <= v102
	if cmp279 {
		goto land_lhs_true281
	} else {
		goto lor_lhs_false284
	}

land_lhs_true281:
	v103 = *libc.As[int32](lookahead)
	cmp282 = v103 <= 70
	if cmp282 {
		goto if_then290
	} else {
		goto lor_lhs_false284
	}

lor_lhs_false284:
	v104 = *libc.As[int32](lookahead)
	cmp285 = 97 <= v104
	if cmp285 {
		goto land_lhs_true287
	} else {
		goto if_end291
	}

land_lhs_true287:
	v105 = *libc.As[int32](lookahead)
	cmp288 = v105 <= 102
	if cmp288 {
		goto if_then290
	} else {
		goto if_end291
	}

if_then290:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end291:
	v106 = *libc.As[byte](result)
	loadedv292 = (v106 & 1) != 0
	*libc.As[bool](retval) = loadedv292
	goto _return

sw_bb293:
	v107 = *libc.As[byte](eof)
	loadedv294 = (v107 & 1) != 0
	if loadedv294 {
		goto if_then295
	} else {
		goto if_end296
	}

if_then295:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end296:
	v108 = *libc.As[int32](lookahead)
	cmp297 = v108 == 38
	if cmp297 {
		goto if_then299
	} else {
		goto if_end300
	}

if_then299:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end300:
	v109 = *libc.As[int32](lookahead)
	cmp301 = v109 == 60
	if cmp301 {
		goto if_then303
	} else {
		goto if_end304
	}

if_then303:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end304:
	v110 = *libc.As[int32](lookahead)
	cmp305 = 9 <= v110
	if cmp305 {
		goto land_lhs_true307
	} else {
		goto lor_lhs_false310
	}

land_lhs_true307:
	v111 = *libc.As[int32](lookahead)
	cmp308 = v111 <= 13
	if cmp308 {
		goto if_then313
	} else {
		goto lor_lhs_false310
	}

lor_lhs_false310:
	v112 = *libc.As[int32](lookahead)
	cmp311 = v112 == 32
	if cmp311 {
		goto if_then313
	} else {
		goto if_end314
	}

if_then313:
	*libc.As[byte](local_skip) = 1
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end314:
	v113 = *libc.As[int32](lookahead)
	cmp315 = v113 != 0
	if cmp315 {
		goto land_lhs_true317
	} else {
		goto if_end321
	}

land_lhs_true317:
	v114 = *libc.As[int32](lookahead)
	cmp318 = v114 != 62
	if cmp318 {
		goto if_then320
	} else {
		goto if_end321
	}

if_then320:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end321:
	v115 = *libc.As[byte](result)
	loadedv322 = (v115 & 1) != 0
	*libc.As[bool](retval) = loadedv322
	goto _return

sw_bb323:
	*libc.As[byte](result) = 1
	v116 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v116).F1)
	*libc.As[int16](result_symbol) = 0
	v117 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v117).F3)
	v118 = *libc.As[unsafe.Pointer](mark_end)
	v119 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v118)(v119)
	v120 = *libc.As[byte](result)
	loadedv324 = (v120 & 1) != 0
	*libc.As[bool](retval) = loadedv324
	goto _return

sw_bb325:
	*libc.As[byte](result) = 1
	v121 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol326 = libc.Ptr(&libc.As[TSLexer](v121).F1)
	*libc.As[int16](result_symbol326) = 1
	v122 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end327 = libc.Ptr(&libc.As[TSLexer](v122).F3)
	v123 = *libc.As[unsafe.Pointer](mark_end327)
	v124 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v123)(v124)
	v125 = *libc.As[byte](result)
	loadedv328 = (v125 & 1) != 0
	*libc.As[bool](retval) = loadedv328
	goto _return

sw_bb329:
	*libc.As[byte](result) = 1
	v126 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol330 = libc.Ptr(&libc.As[TSLexer](v126).F1)
	*libc.As[int16](result_symbol330) = 2
	v127 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end331 = libc.Ptr(&libc.As[TSLexer](v127).F3)
	v128 = *libc.As[unsafe.Pointer](mark_end331)
	v129 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v128)(v129)
	v130 = *libc.As[int32](lookahead)
	cmp332 = 9 <= v130
	if cmp332 {
		goto land_lhs_true334
	} else {
		goto lor_lhs_false337
	}

land_lhs_true334:
	v131 = *libc.As[int32](lookahead)
	cmp335 = v131 <= 13
	if cmp335 {
		goto if_then340
	} else {
		goto lor_lhs_false337
	}

lor_lhs_false337:
	v132 = *libc.As[int32](lookahead)
	cmp338 = v132 == 32
	if cmp338 {
		goto if_then340
	} else {
		goto if_end341
	}

if_then340:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end341:
	v133 = *libc.As[int32](lookahead)
	cmp342 = v133 != 0
	if cmp342 {
		goto land_lhs_true344
	} else {
		goto if_end348
	}

land_lhs_true344:
	v134 = *libc.As[int32](lookahead)
	cmp345 = v134 != 62
	if cmp345 {
		goto if_then347
	} else {
		goto if_end348
	}

if_then347:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end348:
	v135 = *libc.As[byte](result)
	loadedv349 = (v135 & 1) != 0
	*libc.As[bool](retval) = loadedv349
	goto _return

sw_bb350:
	*libc.As[byte](result) = 1
	v136 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol351 = libc.Ptr(&libc.As[TSLexer](v136).F1)
	*libc.As[int16](result_symbol351) = 2
	v137 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end352 = libc.Ptr(&libc.As[TSLexer](v137).F3)
	v138 = *libc.As[unsafe.Pointer](mark_end352)
	v139 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v138)(v139)
	v140 = *libc.As[int32](lookahead)
	cmp353 = v140 != 0
	if cmp353 {
		goto land_lhs_true355
	} else {
		goto if_end359
	}

land_lhs_true355:
	v141 = *libc.As[int32](lookahead)
	cmp356 = v141 != 62
	if cmp356 {
		goto if_then358
	} else {
		goto if_end359
	}

if_then358:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end359:
	v142 = *libc.As[byte](result)
	loadedv360 = (v142 & 1) != 0
	*libc.As[bool](retval) = loadedv360
	goto _return

sw_bb361:
	*libc.As[byte](result) = 1
	v143 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol362 = libc.Ptr(&libc.As[TSLexer](v143).F1)
	*libc.As[int16](result_symbol362) = 3
	v144 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end363 = libc.Ptr(&libc.As[TSLexer](v144).F3)
	v145 = *libc.As[unsafe.Pointer](mark_end363)
	v146 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v145)(v146)
	v147 = *libc.As[byte](result)
	loadedv364 = (v147 & 1) != 0
	*libc.As[bool](retval) = loadedv364
	goto _return

sw_bb365:
	*libc.As[byte](result) = 1
	v148 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol366 = libc.Ptr(&libc.As[TSLexer](v148).F1)
	*libc.As[int16](result_symbol366) = 4
	v149 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end367 = libc.Ptr(&libc.As[TSLexer](v149).F3)
	v150 = *libc.As[unsafe.Pointer](mark_end367)
	v151 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v150)(v151)
	v152 = *libc.As[byte](result)
	loadedv368 = (v152 & 1) != 0
	*libc.As[bool](retval) = loadedv368
	goto _return

sw_bb369:
	*libc.As[byte](result) = 1
	v153 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol370 = libc.Ptr(&libc.As[TSLexer](v153).F1)
	*libc.As[int16](result_symbol370) = 5
	v154 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end371 = libc.Ptr(&libc.As[TSLexer](v154).F3)
	v155 = *libc.As[unsafe.Pointer](mark_end371)
	v156 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v155)(v156)
	v157 = *libc.As[int32](lookahead)
	cmp372 = v157 == 33
	if cmp372 {
		goto if_then374
	} else {
		goto if_end375
	}

if_then374:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end375:
	v158 = *libc.As[int32](lookahead)
	cmp376 = v158 == 47
	if cmp376 {
		goto if_then378
	} else {
		goto if_end379
	}

if_then378:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end379:
	v159 = *libc.As[byte](result)
	loadedv380 = (v159 & 1) != 0
	*libc.As[bool](retval) = loadedv380
	goto _return

sw_bb381:
	*libc.As[byte](result) = 1
	v160 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol382 = libc.Ptr(&libc.As[TSLexer](v160).F1)
	*libc.As[int16](result_symbol382) = 6
	v161 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end383 = libc.Ptr(&libc.As[TSLexer](v161).F3)
	v162 = *libc.As[unsafe.Pointer](mark_end383)
	v163 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v162)(v163)
	v164 = *libc.As[byte](result)
	loadedv384 = (v164 & 1) != 0
	*libc.As[bool](retval) = loadedv384
	goto _return

sw_bb385:
	*libc.As[byte](result) = 1
	v165 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol386 = libc.Ptr(&libc.As[TSLexer](v165).F1)
	*libc.As[int16](result_symbol386) = 7
	v166 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end387 = libc.Ptr(&libc.As[TSLexer](v166).F3)
	v167 = *libc.As[unsafe.Pointer](mark_end387)
	v168 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v167)(v168)
	v169 = *libc.As[byte](result)
	loadedv388 = (v169 & 1) != 0
	*libc.As[bool](retval) = loadedv388
	goto _return

sw_bb389:
	*libc.As[byte](result) = 1
	v170 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol390 = libc.Ptr(&libc.As[TSLexer](v170).F1)
	*libc.As[int16](result_symbol390) = 8
	v171 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end391 = libc.Ptr(&libc.As[TSLexer](v171).F3)
	v172 = *libc.As[unsafe.Pointer](mark_end391)
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v172)(v173)
	v174 = *libc.As[byte](result)
	loadedv392 = (v174 & 1) != 0
	*libc.As[bool](retval) = loadedv392
	goto _return

sw_bb393:
	*libc.As[byte](result) = 1
	v175 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol394 = libc.Ptr(&libc.As[TSLexer](v175).F1)
	*libc.As[int16](result_symbol394) = 9
	v176 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end395 = libc.Ptr(&libc.As[TSLexer](v176).F3)
	v177 = *libc.As[unsafe.Pointer](mark_end395)
	v178 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v177)(v178)
	v179 = *libc.As[int32](lookahead)
	cmp396 = v179 != 0
	if cmp396 {
		goto land_lhs_true398
	} else {
		goto if_end423
	}

land_lhs_true398:
	v180 = *libc.As[int32](lookahead)
	cmp399 = v180 < 9
	if cmp399 {
		goto land_lhs_true404
	} else {
		goto lor_lhs_false401
	}

lor_lhs_false401:
	v181 = *libc.As[int32](lookahead)
	cmp402 = 13 < v181
	if cmp402 {
		goto land_lhs_true404
	} else {
		goto if_end423
	}

land_lhs_true404:
	v182 = *libc.As[int32](lookahead)
	cmp405 = v182 != 32
	if cmp405 {
		goto land_lhs_true407
	} else {
		goto if_end423
	}

land_lhs_true407:
	v183 = *libc.As[int32](lookahead)
	cmp408 = v183 != 34
	if cmp408 {
		goto land_lhs_true410
	} else {
		goto if_end423
	}

land_lhs_true410:
	v184 = *libc.As[int32](lookahead)
	cmp411 = v184 != 39
	if cmp411 {
		goto land_lhs_true413
	} else {
		goto if_end423
	}

land_lhs_true413:
	v185 = *libc.As[int32](lookahead)
	cmp414 = v185 != 47
	if cmp414 {
		goto land_lhs_true416
	} else {
		goto if_end423
	}

land_lhs_true416:
	v186 = *libc.As[int32](lookahead)
	cmp417 = v186 < 60
	if cmp417 {
		goto if_then422
	} else {
		goto lor_lhs_false419
	}

lor_lhs_false419:
	v187 = *libc.As[int32](lookahead)
	cmp420 = 62 < v187
	if cmp420 {
		goto if_then422
	} else {
		goto if_end423
	}

if_then422:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end423:
	v188 = *libc.As[byte](result)
	loadedv424 = (v188 & 1) != 0
	*libc.As[bool](retval) = loadedv424
	goto _return

sw_bb425:
	*libc.As[byte](result) = 1
	v189 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol426 = libc.Ptr(&libc.As[TSLexer](v189).F1)
	*libc.As[int16](result_symbol426) = 10
	v190 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end427 = libc.Ptr(&libc.As[TSLexer](v190).F3)
	v191 = *libc.As[unsafe.Pointer](mark_end427)
	v192 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v191)(v192)
	v193 = *libc.As[int32](lookahead)
	cmp428 = v193 != 0
	if cmp428 {
		goto land_lhs_true430
	} else {
		goto if_end452
	}

land_lhs_true430:
	v194 = *libc.As[int32](lookahead)
	cmp431 = v194 < 9
	if cmp431 {
		goto land_lhs_true436
	} else {
		goto lor_lhs_false433
	}

lor_lhs_false433:
	v195 = *libc.As[int32](lookahead)
	cmp434 = 13 < v195
	if cmp434 {
		goto land_lhs_true436
	} else {
		goto if_end452
	}

land_lhs_true436:
	v196 = *libc.As[int32](lookahead)
	cmp437 = v196 != 32
	if cmp437 {
		goto land_lhs_true439
	} else {
		goto if_end452
	}

land_lhs_true439:
	v197 = *libc.As[int32](lookahead)
	cmp440 = v197 != 34
	if cmp440 {
		goto land_lhs_true442
	} else {
		goto if_end452
	}

land_lhs_true442:
	v198 = *libc.As[int32](lookahead)
	cmp443 = v198 != 39
	if cmp443 {
		goto land_lhs_true445
	} else {
		goto if_end452
	}

land_lhs_true445:
	v199 = *libc.As[int32](lookahead)
	cmp446 = v199 < 60
	if cmp446 {
		goto if_then451
	} else {
		goto lor_lhs_false448
	}

lor_lhs_false448:
	v200 = *libc.As[int32](lookahead)
	cmp449 = 62 < v200
	if cmp449 {
		goto if_then451
	} else {
		goto if_end452
	}

if_then451:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end452:
	v201 = *libc.As[byte](result)
	loadedv453 = (v201 & 1) != 0
	*libc.As[bool](retval) = loadedv453
	goto _return

sw_bb454:
	*libc.As[byte](result) = 1
	v202 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol455 = libc.Ptr(&libc.As[TSLexer](v202).F1)
	*libc.As[int16](result_symbol455) = 11
	v203 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end456 = libc.Ptr(&libc.As[TSLexer](v203).F3)
	v204 = *libc.As[unsafe.Pointer](mark_end456)
	v205 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v204)(v205)
	v206 = *libc.As[byte](result)
	loadedv457 = (v206 & 1) != 0
	*libc.As[bool](retval) = loadedv457
	goto _return

sw_bb458:
	*libc.As[byte](result) = 1
	v207 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol459 = libc.Ptr(&libc.As[TSLexer](v207).F1)
	*libc.As[int16](result_symbol459) = 11
	v208 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end460 = libc.Ptr(&libc.As[TSLexer](v208).F3)
	v209 = *libc.As[unsafe.Pointer](mark_end460)
	v210 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v209)(v210)
	v211 = *libc.As[int32](lookahead)
	cmp461 = v211 == 59
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end464:
	v212 = *libc.As[byte](result)
	loadedv465 = (v212 & 1) != 0
	*libc.As[bool](retval) = loadedv465
	goto _return

sw_bb466:
	*libc.As[byte](result) = 1
	v213 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol467 = libc.Ptr(&libc.As[TSLexer](v213).F1)
	*libc.As[int16](result_symbol467) = 11
	v214 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end468 = libc.Ptr(&libc.As[TSLexer](v214).F3)
	v215 = *libc.As[unsafe.Pointer](mark_end468)
	v216 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v215)(v216)
	v217 = *libc.As[int32](lookahead)
	cmp469 = v217 == 59
	if cmp469 {
		goto if_then471
	} else {
		goto if_end472
	}

if_then471:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end472:
	v218 = *libc.As[int32](lookahead)
	cmp473 = 48 <= v218
	if cmp473 {
		goto land_lhs_true475
	} else {
		goto if_end479
	}

land_lhs_true475:
	v219 = *libc.As[int32](lookahead)
	cmp476 = v219 <= 57
	if cmp476 {
		goto if_then478
	} else {
		goto if_end479
	}

if_then478:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end479:
	v220 = *libc.As[byte](result)
	loadedv480 = (v220 & 1) != 0
	*libc.As[bool](retval) = loadedv480
	goto _return

sw_bb481:
	*libc.As[byte](result) = 1
	v221 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol482 = libc.Ptr(&libc.As[TSLexer](v221).F1)
	*libc.As[int16](result_symbol482) = 11
	v222 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end483 = libc.Ptr(&libc.As[TSLexer](v222).F3)
	v223 = *libc.As[unsafe.Pointer](mark_end483)
	v224 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v223)(v224)
	v225 = *libc.As[int32](lookahead)
	cmp484 = v225 == 59
	if cmp484 {
		goto if_then486
	} else {
		goto if_end487
	}

if_then486:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end487:
	v226 = *libc.As[int32](lookahead)
	cmp488 = 48 <= v226
	if cmp488 {
		goto land_lhs_true490
	} else {
		goto if_end494
	}

land_lhs_true490:
	v227 = *libc.As[int32](lookahead)
	cmp491 = v227 <= 57
	if cmp491 {
		goto if_then493
	} else {
		goto if_end494
	}

if_then493:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end494:
	v228 = *libc.As[byte](result)
	loadedv495 = (v228 & 1) != 0
	*libc.As[bool](retval) = loadedv495
	goto _return

sw_bb496:
	*libc.As[byte](result) = 1
	v229 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol497 = libc.Ptr(&libc.As[TSLexer](v229).F1)
	*libc.As[int16](result_symbol497) = 11
	v230 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end498 = libc.Ptr(&libc.As[TSLexer](v230).F3)
	v231 = *libc.As[unsafe.Pointer](mark_end498)
	v232 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v231)(v232)
	v233 = *libc.As[int32](lookahead)
	cmp499 = v233 == 59
	if cmp499 {
		goto if_then501
	} else {
		goto if_end502
	}

if_then501:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end502:
	v234 = *libc.As[int32](lookahead)
	cmp503 = 48 <= v234
	if cmp503 {
		goto land_lhs_true505
	} else {
		goto if_end509
	}

land_lhs_true505:
	v235 = *libc.As[int32](lookahead)
	cmp506 = v235 <= 57
	if cmp506 {
		goto if_then508
	} else {
		goto if_end509
	}

if_then508:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end509:
	v236 = *libc.As[byte](result)
	loadedv510 = (v236 & 1) != 0
	*libc.As[bool](retval) = loadedv510
	goto _return

sw_bb511:
	*libc.As[byte](result) = 1
	v237 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol512 = libc.Ptr(&libc.As[TSLexer](v237).F1)
	*libc.As[int16](result_symbol512) = 11
	v238 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end513 = libc.Ptr(&libc.As[TSLexer](v238).F3)
	v239 = *libc.As[unsafe.Pointer](mark_end513)
	v240 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v239)(v240)
	v241 = *libc.As[int32](lookahead)
	cmp514 = v241 == 59
	if cmp514 {
		goto if_then516
	} else {
		goto if_end517
	}

if_then516:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end517:
	v242 = *libc.As[int32](lookahead)
	cmp518 = 48 <= v242
	if cmp518 {
		goto land_lhs_true520
	} else {
		goto if_end524
	}

land_lhs_true520:
	v243 = *libc.As[int32](lookahead)
	cmp521 = v243 <= 57
	if cmp521 {
		goto if_then523
	} else {
		goto if_end524
	}

if_then523:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end524:
	v244 = *libc.As[byte](result)
	loadedv525 = (v244 & 1) != 0
	*libc.As[bool](retval) = loadedv525
	goto _return

sw_bb526:
	*libc.As[byte](result) = 1
	v245 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol527 = libc.Ptr(&libc.As[TSLexer](v245).F1)
	*libc.As[int16](result_symbol527) = 11
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end528 = libc.Ptr(&libc.As[TSLexer](v246).F3)
	v247 = *libc.As[unsafe.Pointer](mark_end528)
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v247)(v248)
	v249 = *libc.As[int32](lookahead)
	cmp529 = v249 == 59
	if cmp529 {
		goto if_then531
	} else {
		goto if_end532
	}

if_then531:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end532:
	v250 = *libc.As[int32](lookahead)
	cmp533 = 48 <= v250
	if cmp533 {
		goto land_lhs_true535
	} else {
		goto lor_lhs_false538
	}

land_lhs_true535:
	v251 = *libc.As[int32](lookahead)
	cmp536 = v251 <= 57
	if cmp536 {
		goto if_then550
	} else {
		goto lor_lhs_false538
	}

lor_lhs_false538:
	v252 = *libc.As[int32](lookahead)
	cmp539 = 65 <= v252
	if cmp539 {
		goto land_lhs_true541
	} else {
		goto lor_lhs_false544
	}

land_lhs_true541:
	v253 = *libc.As[int32](lookahead)
	cmp542 = v253 <= 70
	if cmp542 {
		goto if_then550
	} else {
		goto lor_lhs_false544
	}

lor_lhs_false544:
	v254 = *libc.As[int32](lookahead)
	cmp545 = 97 <= v254
	if cmp545 {
		goto land_lhs_true547
	} else {
		goto if_end551
	}

land_lhs_true547:
	v255 = *libc.As[int32](lookahead)
	cmp548 = v255 <= 102
	if cmp548 {
		goto if_then550
	} else {
		goto if_end551
	}

if_then550:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end551:
	v256 = *libc.As[byte](result)
	loadedv552 = (v256 & 1) != 0
	*libc.As[bool](retval) = loadedv552
	goto _return

sw_bb553:
	*libc.As[byte](result) = 1
	v257 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol554 = libc.Ptr(&libc.As[TSLexer](v257).F1)
	*libc.As[int16](result_symbol554) = 11
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end555 = libc.Ptr(&libc.As[TSLexer](v258).F3)
	v259 = *libc.As[unsafe.Pointer](mark_end555)
	v260 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v259)(v260)
	v261 = *libc.As[int32](lookahead)
	cmp556 = v261 == 59
	if cmp556 {
		goto if_then558
	} else {
		goto if_end559
	}

if_then558:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end559:
	v262 = *libc.As[int32](lookahead)
	cmp560 = 48 <= v262
	if cmp560 {
		goto land_lhs_true562
	} else {
		goto lor_lhs_false565
	}

land_lhs_true562:
	v263 = *libc.As[int32](lookahead)
	cmp563 = v263 <= 57
	if cmp563 {
		goto if_then577
	} else {
		goto lor_lhs_false565
	}

lor_lhs_false565:
	v264 = *libc.As[int32](lookahead)
	cmp566 = 65 <= v264
	if cmp566 {
		goto land_lhs_true568
	} else {
		goto lor_lhs_false571
	}

land_lhs_true568:
	v265 = *libc.As[int32](lookahead)
	cmp569 = v265 <= 70
	if cmp569 {
		goto if_then577
	} else {
		goto lor_lhs_false571
	}

lor_lhs_false571:
	v266 = *libc.As[int32](lookahead)
	cmp572 = 97 <= v266
	if cmp572 {
		goto land_lhs_true574
	} else {
		goto if_end578
	}

land_lhs_true574:
	v267 = *libc.As[int32](lookahead)
	cmp575 = v267 <= 102
	if cmp575 {
		goto if_then577
	} else {
		goto if_end578
	}

if_then577:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end578:
	v268 = *libc.As[byte](result)
	loadedv579 = (v268 & 1) != 0
	*libc.As[bool](retval) = loadedv579
	goto _return

sw_bb580:
	*libc.As[byte](result) = 1
	v269 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol581 = libc.Ptr(&libc.As[TSLexer](v269).F1)
	*libc.As[int16](result_symbol581) = 11
	v270 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end582 = libc.Ptr(&libc.As[TSLexer](v270).F3)
	v271 = *libc.As[unsafe.Pointer](mark_end582)
	v272 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v271)(v272)
	v273 = *libc.As[int32](lookahead)
	cmp583 = v273 == 59
	if cmp583 {
		goto if_then585
	} else {
		goto if_end586
	}

if_then585:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end586:
	v274 = *libc.As[int32](lookahead)
	cmp587 = 48 <= v274
	if cmp587 {
		goto land_lhs_true589
	} else {
		goto lor_lhs_false592
	}

land_lhs_true589:
	v275 = *libc.As[int32](lookahead)
	cmp590 = v275 <= 57
	if cmp590 {
		goto if_then604
	} else {
		goto lor_lhs_false592
	}

lor_lhs_false592:
	v276 = *libc.As[int32](lookahead)
	cmp593 = 65 <= v276
	if cmp593 {
		goto land_lhs_true595
	} else {
		goto lor_lhs_false598
	}

land_lhs_true595:
	v277 = *libc.As[int32](lookahead)
	cmp596 = v277 <= 70
	if cmp596 {
		goto if_then604
	} else {
		goto lor_lhs_false598
	}

lor_lhs_false598:
	v278 = *libc.As[int32](lookahead)
	cmp599 = 97 <= v278
	if cmp599 {
		goto land_lhs_true601
	} else {
		goto if_end605
	}

land_lhs_true601:
	v279 = *libc.As[int32](lookahead)
	cmp602 = v279 <= 102
	if cmp602 {
		goto if_then604
	} else {
		goto if_end605
	}

if_then604:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end605:
	v280 = *libc.As[byte](result)
	loadedv606 = (v280 & 1) != 0
	*libc.As[bool](retval) = loadedv606
	goto _return

sw_bb607:
	*libc.As[byte](result) = 1
	v281 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol608 = libc.Ptr(&libc.As[TSLexer](v281).F1)
	*libc.As[int16](result_symbol608) = 11
	v282 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end609 = libc.Ptr(&libc.As[TSLexer](v282).F3)
	v283 = *libc.As[unsafe.Pointer](mark_end609)
	v284 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v283)(v284)
	v285 = *libc.As[int32](lookahead)
	cmp610 = v285 == 59
	if cmp610 {
		goto if_then612
	} else {
		goto if_end613
	}

if_then612:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end613:
	v286 = *libc.As[int32](lookahead)
	cmp614 = 48 <= v286
	if cmp614 {
		goto land_lhs_true616
	} else {
		goto lor_lhs_false619
	}

land_lhs_true616:
	v287 = *libc.As[int32](lookahead)
	cmp617 = v287 <= 57
	if cmp617 {
		goto if_then631
	} else {
		goto lor_lhs_false619
	}

lor_lhs_false619:
	v288 = *libc.As[int32](lookahead)
	cmp620 = 65 <= v288
	if cmp620 {
		goto land_lhs_true622
	} else {
		goto lor_lhs_false625
	}

land_lhs_true622:
	v289 = *libc.As[int32](lookahead)
	cmp623 = v289 <= 70
	if cmp623 {
		goto if_then631
	} else {
		goto lor_lhs_false625
	}

lor_lhs_false625:
	v290 = *libc.As[int32](lookahead)
	cmp626 = 97 <= v290
	if cmp626 {
		goto land_lhs_true628
	} else {
		goto if_end632
	}

land_lhs_true628:
	v291 = *libc.As[int32](lookahead)
	cmp629 = v291 <= 102
	if cmp629 {
		goto if_then631
	} else {
		goto if_end632
	}

if_then631:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end632:
	v292 = *libc.As[byte](result)
	loadedv633 = (v292 & 1) != 0
	*libc.As[bool](retval) = loadedv633
	goto _return

sw_bb634:
	*libc.As[byte](result) = 1
	v293 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol635 = libc.Ptr(&libc.As[TSLexer](v293).F1)
	*libc.As[int16](result_symbol635) = 11
	v294 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end636 = libc.Ptr(&libc.As[TSLexer](v294).F3)
	v295 = *libc.As[unsafe.Pointer](mark_end636)
	v296 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v295)(v296)
	v297 = *libc.As[int32](lookahead)
	cmp637 = v297 == 59
	if cmp637 {
		goto if_then639
	} else {
		goto if_end640
	}

if_then639:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end640:
	v298 = *libc.As[int32](lookahead)
	cmp641 = 48 <= v298
	if cmp641 {
		goto land_lhs_true643
	} else {
		goto lor_lhs_false646
	}

land_lhs_true643:
	v299 = *libc.As[int32](lookahead)
	cmp644 = v299 <= 57
	if cmp644 {
		goto if_then658
	} else {
		goto lor_lhs_false646
	}

lor_lhs_false646:
	v300 = *libc.As[int32](lookahead)
	cmp647 = 65 <= v300
	if cmp647 {
		goto land_lhs_true649
	} else {
		goto lor_lhs_false652
	}

land_lhs_true649:
	v301 = *libc.As[int32](lookahead)
	cmp650 = v301 <= 70
	if cmp650 {
		goto if_then658
	} else {
		goto lor_lhs_false652
	}

lor_lhs_false652:
	v302 = *libc.As[int32](lookahead)
	cmp653 = 97 <= v302
	if cmp653 {
		goto land_lhs_true655
	} else {
		goto if_end659
	}

land_lhs_true655:
	v303 = *libc.As[int32](lookahead)
	cmp656 = v303 <= 102
	if cmp656 {
		goto if_then658
	} else {
		goto if_end659
	}

if_then658:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end659:
	v304 = *libc.As[byte](result)
	loadedv660 = (v304 & 1) != 0
	*libc.As[bool](retval) = loadedv660
	goto _return

sw_bb661:
	*libc.As[byte](result) = 1
	v305 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol662 = libc.Ptr(&libc.As[TSLexer](v305).F1)
	*libc.As[int16](result_symbol662) = 11
	v306 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end663 = libc.Ptr(&libc.As[TSLexer](v306).F3)
	v307 = *libc.As[unsafe.Pointer](mark_end663)
	v308 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v307)(v308)
	v309 = *libc.As[int32](lookahead)
	cmp664 = v309 == 59
	if cmp664 {
		goto if_then666
	} else {
		goto if_end667
	}

if_then666:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end667:
	v310 = *libc.As[int32](lookahead)
	cmp668 = 65 <= v310
	if cmp668 {
		goto land_lhs_true670
	} else {
		goto lor_lhs_false673
	}

land_lhs_true670:
	v311 = *libc.As[int32](lookahead)
	cmp671 = v311 <= 90
	if cmp671 {
		goto if_then679
	} else {
		goto lor_lhs_false673
	}

lor_lhs_false673:
	v312 = *libc.As[int32](lookahead)
	cmp674 = 97 <= v312
	if cmp674 {
		goto land_lhs_true676
	} else {
		goto if_end680
	}

land_lhs_true676:
	v313 = *libc.As[int32](lookahead)
	cmp677 = v313 <= 122
	if cmp677 {
		goto if_then679
	} else {
		goto if_end680
	}

if_then679:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end680:
	v314 = *libc.As[byte](result)
	loadedv681 = (v314 & 1) != 0
	*libc.As[bool](retval) = loadedv681
	goto _return

sw_bb682:
	*libc.As[byte](result) = 1
	v315 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol683 = libc.Ptr(&libc.As[TSLexer](v315).F1)
	*libc.As[int16](result_symbol683) = 11
	v316 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end684 = libc.Ptr(&libc.As[TSLexer](v316).F3)
	v317 = *libc.As[unsafe.Pointer](mark_end684)
	v318 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v317)(v318)
	v319 = *libc.As[int32](lookahead)
	cmp685 = v319 == 59
	if cmp685 {
		goto if_then687
	} else {
		goto if_end688
	}

if_then687:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end688:
	v320 = *libc.As[int32](lookahead)
	cmp689 = 65 <= v320
	if cmp689 {
		goto land_lhs_true691
	} else {
		goto lor_lhs_false694
	}

land_lhs_true691:
	v321 = *libc.As[int32](lookahead)
	cmp692 = v321 <= 90
	if cmp692 {
		goto if_then700
	} else {
		goto lor_lhs_false694
	}

lor_lhs_false694:
	v322 = *libc.As[int32](lookahead)
	cmp695 = 97 <= v322
	if cmp695 {
		goto land_lhs_true697
	} else {
		goto if_end701
	}

land_lhs_true697:
	v323 = *libc.As[int32](lookahead)
	cmp698 = v323 <= 122
	if cmp698 {
		goto if_then700
	} else {
		goto if_end701
	}

if_then700:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end701:
	v324 = *libc.As[byte](result)
	loadedv702 = (v324 & 1) != 0
	*libc.As[bool](retval) = loadedv702
	goto _return

sw_bb703:
	*libc.As[byte](result) = 1
	v325 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol704 = libc.Ptr(&libc.As[TSLexer](v325).F1)
	*libc.As[int16](result_symbol704) = 11
	v326 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end705 = libc.Ptr(&libc.As[TSLexer](v326).F3)
	v327 = *libc.As[unsafe.Pointer](mark_end705)
	v328 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v327)(v328)
	v329 = *libc.As[int32](lookahead)
	cmp706 = v329 == 59
	if cmp706 {
		goto if_then708
	} else {
		goto if_end709
	}

if_then708:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end709:
	v330 = *libc.As[int32](lookahead)
	cmp710 = 65 <= v330
	if cmp710 {
		goto land_lhs_true712
	} else {
		goto lor_lhs_false715
	}

land_lhs_true712:
	v331 = *libc.As[int32](lookahead)
	cmp713 = v331 <= 90
	if cmp713 {
		goto if_then721
	} else {
		goto lor_lhs_false715
	}

lor_lhs_false715:
	v332 = *libc.As[int32](lookahead)
	cmp716 = 97 <= v332
	if cmp716 {
		goto land_lhs_true718
	} else {
		goto if_end722
	}

land_lhs_true718:
	v333 = *libc.As[int32](lookahead)
	cmp719 = v333 <= 122
	if cmp719 {
		goto if_then721
	} else {
		goto if_end722
	}

if_then721:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end722:
	v334 = *libc.As[byte](result)
	loadedv723 = (v334 & 1) != 0
	*libc.As[bool](retval) = loadedv723
	goto _return

sw_bb724:
	*libc.As[byte](result) = 1
	v335 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol725 = libc.Ptr(&libc.As[TSLexer](v335).F1)
	*libc.As[int16](result_symbol725) = 11
	v336 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end726 = libc.Ptr(&libc.As[TSLexer](v336).F3)
	v337 = *libc.As[unsafe.Pointer](mark_end726)
	v338 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v337)(v338)
	v339 = *libc.As[int32](lookahead)
	cmp727 = v339 == 59
	if cmp727 {
		goto if_then729
	} else {
		goto if_end730
	}

if_then729:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end730:
	v340 = *libc.As[int32](lookahead)
	cmp731 = 65 <= v340
	if cmp731 {
		goto land_lhs_true733
	} else {
		goto lor_lhs_false736
	}

land_lhs_true733:
	v341 = *libc.As[int32](lookahead)
	cmp734 = v341 <= 90
	if cmp734 {
		goto if_then742
	} else {
		goto lor_lhs_false736
	}

lor_lhs_false736:
	v342 = *libc.As[int32](lookahead)
	cmp737 = 97 <= v342
	if cmp737 {
		goto land_lhs_true739
	} else {
		goto if_end743
	}

land_lhs_true739:
	v343 = *libc.As[int32](lookahead)
	cmp740 = v343 <= 122
	if cmp740 {
		goto if_then742
	} else {
		goto if_end743
	}

if_then742:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end743:
	v344 = *libc.As[byte](result)
	loadedv744 = (v344 & 1) != 0
	*libc.As[bool](retval) = loadedv744
	goto _return

sw_bb745:
	*libc.As[byte](result) = 1
	v345 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol746 = libc.Ptr(&libc.As[TSLexer](v345).F1)
	*libc.As[int16](result_symbol746) = 11
	v346 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end747 = libc.Ptr(&libc.As[TSLexer](v346).F3)
	v347 = *libc.As[unsafe.Pointer](mark_end747)
	v348 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v347)(v348)
	v349 = *libc.As[int32](lookahead)
	cmp748 = v349 == 59
	if cmp748 {
		goto if_then750
	} else {
		goto if_end751
	}

if_then750:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end751:
	v350 = *libc.As[int32](lookahead)
	cmp752 = 65 <= v350
	if cmp752 {
		goto land_lhs_true754
	} else {
		goto lor_lhs_false757
	}

land_lhs_true754:
	v351 = *libc.As[int32](lookahead)
	cmp755 = v351 <= 90
	if cmp755 {
		goto if_then763
	} else {
		goto lor_lhs_false757
	}

lor_lhs_false757:
	v352 = *libc.As[int32](lookahead)
	cmp758 = 97 <= v352
	if cmp758 {
		goto land_lhs_true760
	} else {
		goto if_end764
	}

land_lhs_true760:
	v353 = *libc.As[int32](lookahead)
	cmp761 = v353 <= 122
	if cmp761 {
		goto if_then763
	} else {
		goto if_end764
	}

if_then763:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end764:
	v354 = *libc.As[byte](result)
	loadedv765 = (v354 & 1) != 0
	*libc.As[bool](retval) = loadedv765
	goto _return

sw_bb766:
	*libc.As[byte](result) = 1
	v355 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol767 = libc.Ptr(&libc.As[TSLexer](v355).F1)
	*libc.As[int16](result_symbol767) = 11
	v356 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end768 = libc.Ptr(&libc.As[TSLexer](v356).F3)
	v357 = *libc.As[unsafe.Pointer](mark_end768)
	v358 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v357)(v358)
	v359 = *libc.As[int32](lookahead)
	cmp769 = v359 == 59
	if cmp769 {
		goto if_then771
	} else {
		goto if_end772
	}

if_then771:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end772:
	v360 = *libc.As[int32](lookahead)
	cmp773 = 65 <= v360
	if cmp773 {
		goto land_lhs_true775
	} else {
		goto lor_lhs_false778
	}

land_lhs_true775:
	v361 = *libc.As[int32](lookahead)
	cmp776 = v361 <= 90
	if cmp776 {
		goto if_then784
	} else {
		goto lor_lhs_false778
	}

lor_lhs_false778:
	v362 = *libc.As[int32](lookahead)
	cmp779 = 97 <= v362
	if cmp779 {
		goto land_lhs_true781
	} else {
		goto if_end785
	}

land_lhs_true781:
	v363 = *libc.As[int32](lookahead)
	cmp782 = v363 <= 122
	if cmp782 {
		goto if_then784
	} else {
		goto if_end785
	}

if_then784:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end785:
	v364 = *libc.As[byte](result)
	loadedv786 = (v364 & 1) != 0
	*libc.As[bool](retval) = loadedv786
	goto _return

sw_bb787:
	*libc.As[byte](result) = 1
	v365 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol788 = libc.Ptr(&libc.As[TSLexer](v365).F1)
	*libc.As[int16](result_symbol788) = 11
	v366 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end789 = libc.Ptr(&libc.As[TSLexer](v366).F3)
	v367 = *libc.As[unsafe.Pointer](mark_end789)
	v368 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v367)(v368)
	v369 = *libc.As[int32](lookahead)
	cmp790 = v369 == 59
	if cmp790 {
		goto if_then792
	} else {
		goto if_end793
	}

if_then792:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end793:
	v370 = *libc.As[int32](lookahead)
	cmp794 = 65 <= v370
	if cmp794 {
		goto land_lhs_true796
	} else {
		goto lor_lhs_false799
	}

land_lhs_true796:
	v371 = *libc.As[int32](lookahead)
	cmp797 = v371 <= 90
	if cmp797 {
		goto if_then805
	} else {
		goto lor_lhs_false799
	}

lor_lhs_false799:
	v372 = *libc.As[int32](lookahead)
	cmp800 = 97 <= v372
	if cmp800 {
		goto land_lhs_true802
	} else {
		goto if_end806
	}

land_lhs_true802:
	v373 = *libc.As[int32](lookahead)
	cmp803 = v373 <= 122
	if cmp803 {
		goto if_then805
	} else {
		goto if_end806
	}

if_then805:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end806:
	v374 = *libc.As[byte](result)
	loadedv807 = (v374 & 1) != 0
	*libc.As[bool](retval) = loadedv807
	goto _return

sw_bb808:
	*libc.As[byte](result) = 1
	v375 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol809 = libc.Ptr(&libc.As[TSLexer](v375).F1)
	*libc.As[int16](result_symbol809) = 11
	v376 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end810 = libc.Ptr(&libc.As[TSLexer](v376).F3)
	v377 = *libc.As[unsafe.Pointer](mark_end810)
	v378 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v377)(v378)
	v379 = *libc.As[int32](lookahead)
	cmp811 = v379 == 59
	if cmp811 {
		goto if_then813
	} else {
		goto if_end814
	}

if_then813:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end814:
	v380 = *libc.As[int32](lookahead)
	cmp815 = 65 <= v380
	if cmp815 {
		goto land_lhs_true817
	} else {
		goto lor_lhs_false820
	}

land_lhs_true817:
	v381 = *libc.As[int32](lookahead)
	cmp818 = v381 <= 90
	if cmp818 {
		goto if_then826
	} else {
		goto lor_lhs_false820
	}

lor_lhs_false820:
	v382 = *libc.As[int32](lookahead)
	cmp821 = 97 <= v382
	if cmp821 {
		goto land_lhs_true823
	} else {
		goto if_end827
	}

land_lhs_true823:
	v383 = *libc.As[int32](lookahead)
	cmp824 = v383 <= 122
	if cmp824 {
		goto if_then826
	} else {
		goto if_end827
	}

if_then826:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end827:
	v384 = *libc.As[byte](result)
	loadedv828 = (v384 & 1) != 0
	*libc.As[bool](retval) = loadedv828
	goto _return

sw_bb829:
	*libc.As[byte](result) = 1
	v385 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol830 = libc.Ptr(&libc.As[TSLexer](v385).F1)
	*libc.As[int16](result_symbol830) = 11
	v386 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end831 = libc.Ptr(&libc.As[TSLexer](v386).F3)
	v387 = *libc.As[unsafe.Pointer](mark_end831)
	v388 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v387)(v388)
	v389 = *libc.As[int32](lookahead)
	cmp832 = v389 == 59
	if cmp832 {
		goto if_then834
	} else {
		goto if_end835
	}

if_then834:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end835:
	v390 = *libc.As[int32](lookahead)
	cmp836 = 65 <= v390
	if cmp836 {
		goto land_lhs_true838
	} else {
		goto lor_lhs_false841
	}

land_lhs_true838:
	v391 = *libc.As[int32](lookahead)
	cmp839 = v391 <= 90
	if cmp839 {
		goto if_then847
	} else {
		goto lor_lhs_false841
	}

lor_lhs_false841:
	v392 = *libc.As[int32](lookahead)
	cmp842 = 97 <= v392
	if cmp842 {
		goto land_lhs_true844
	} else {
		goto if_end848
	}

land_lhs_true844:
	v393 = *libc.As[int32](lookahead)
	cmp845 = v393 <= 122
	if cmp845 {
		goto if_then847
	} else {
		goto if_end848
	}

if_then847:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end848:
	v394 = *libc.As[byte](result)
	loadedv849 = (v394 & 1) != 0
	*libc.As[bool](retval) = loadedv849
	goto _return

sw_bb850:
	*libc.As[byte](result) = 1
	v395 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol851 = libc.Ptr(&libc.As[TSLexer](v395).F1)
	*libc.As[int16](result_symbol851) = 11
	v396 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end852 = libc.Ptr(&libc.As[TSLexer](v396).F3)
	v397 = *libc.As[unsafe.Pointer](mark_end852)
	v398 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v397)(v398)
	v399 = *libc.As[int32](lookahead)
	cmp853 = v399 == 59
	if cmp853 {
		goto if_then855
	} else {
		goto if_end856
	}

if_then855:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end856:
	v400 = *libc.As[int32](lookahead)
	cmp857 = 65 <= v400
	if cmp857 {
		goto land_lhs_true859
	} else {
		goto lor_lhs_false862
	}

land_lhs_true859:
	v401 = *libc.As[int32](lookahead)
	cmp860 = v401 <= 90
	if cmp860 {
		goto if_then868
	} else {
		goto lor_lhs_false862
	}

lor_lhs_false862:
	v402 = *libc.As[int32](lookahead)
	cmp863 = 97 <= v402
	if cmp863 {
		goto land_lhs_true865
	} else {
		goto if_end869
	}

land_lhs_true865:
	v403 = *libc.As[int32](lookahead)
	cmp866 = v403 <= 122
	if cmp866 {
		goto if_then868
	} else {
		goto if_end869
	}

if_then868:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end869:
	v404 = *libc.As[byte](result)
	loadedv870 = (v404 & 1) != 0
	*libc.As[bool](retval) = loadedv870
	goto _return

sw_bb871:
	*libc.As[byte](result) = 1
	v405 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol872 = libc.Ptr(&libc.As[TSLexer](v405).F1)
	*libc.As[int16](result_symbol872) = 11
	v406 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end873 = libc.Ptr(&libc.As[TSLexer](v406).F3)
	v407 = *libc.As[unsafe.Pointer](mark_end873)
	v408 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v407)(v408)
	v409 = *libc.As[int32](lookahead)
	cmp874 = v409 == 59
	if cmp874 {
		goto if_then876
	} else {
		goto if_end877
	}

if_then876:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end877:
	v410 = *libc.As[int32](lookahead)
	cmp878 = 65 <= v410
	if cmp878 {
		goto land_lhs_true880
	} else {
		goto lor_lhs_false883
	}

land_lhs_true880:
	v411 = *libc.As[int32](lookahead)
	cmp881 = v411 <= 90
	if cmp881 {
		goto if_then889
	} else {
		goto lor_lhs_false883
	}

lor_lhs_false883:
	v412 = *libc.As[int32](lookahead)
	cmp884 = 97 <= v412
	if cmp884 {
		goto land_lhs_true886
	} else {
		goto if_end890
	}

land_lhs_true886:
	v413 = *libc.As[int32](lookahead)
	cmp887 = v413 <= 122
	if cmp887 {
		goto if_then889
	} else {
		goto if_end890
	}

if_then889:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end890:
	v414 = *libc.As[byte](result)
	loadedv891 = (v414 & 1) != 0
	*libc.As[bool](retval) = loadedv891
	goto _return

sw_bb892:
	*libc.As[byte](result) = 1
	v415 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol893 = libc.Ptr(&libc.As[TSLexer](v415).F1)
	*libc.As[int16](result_symbol893) = 11
	v416 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end894 = libc.Ptr(&libc.As[TSLexer](v416).F3)
	v417 = *libc.As[unsafe.Pointer](mark_end894)
	v418 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v417)(v418)
	v419 = *libc.As[int32](lookahead)
	cmp895 = v419 == 59
	if cmp895 {
		goto if_then897
	} else {
		goto if_end898
	}

if_then897:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end898:
	v420 = *libc.As[int32](lookahead)
	cmp899 = 65 <= v420
	if cmp899 {
		goto land_lhs_true901
	} else {
		goto lor_lhs_false904
	}

land_lhs_true901:
	v421 = *libc.As[int32](lookahead)
	cmp902 = v421 <= 90
	if cmp902 {
		goto if_then910
	} else {
		goto lor_lhs_false904
	}

lor_lhs_false904:
	v422 = *libc.As[int32](lookahead)
	cmp905 = 97 <= v422
	if cmp905 {
		goto land_lhs_true907
	} else {
		goto if_end911
	}

land_lhs_true907:
	v423 = *libc.As[int32](lookahead)
	cmp908 = v423 <= 122
	if cmp908 {
		goto if_then910
	} else {
		goto if_end911
	}

if_then910:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end911:
	v424 = *libc.As[byte](result)
	loadedv912 = (v424 & 1) != 0
	*libc.As[bool](retval) = loadedv912
	goto _return

sw_bb913:
	*libc.As[byte](result) = 1
	v425 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol914 = libc.Ptr(&libc.As[TSLexer](v425).F1)
	*libc.As[int16](result_symbol914) = 11
	v426 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end915 = libc.Ptr(&libc.As[TSLexer](v426).F3)
	v427 = *libc.As[unsafe.Pointer](mark_end915)
	v428 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v427)(v428)
	v429 = *libc.As[int32](lookahead)
	cmp916 = v429 == 59
	if cmp916 {
		goto if_then918
	} else {
		goto if_end919
	}

if_then918:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end919:
	v430 = *libc.As[int32](lookahead)
	cmp920 = 65 <= v430
	if cmp920 {
		goto land_lhs_true922
	} else {
		goto lor_lhs_false925
	}

land_lhs_true922:
	v431 = *libc.As[int32](lookahead)
	cmp923 = v431 <= 90
	if cmp923 {
		goto if_then931
	} else {
		goto lor_lhs_false925
	}

lor_lhs_false925:
	v432 = *libc.As[int32](lookahead)
	cmp926 = 97 <= v432
	if cmp926 {
		goto land_lhs_true928
	} else {
		goto if_end932
	}

land_lhs_true928:
	v433 = *libc.As[int32](lookahead)
	cmp929 = v433 <= 122
	if cmp929 {
		goto if_then931
	} else {
		goto if_end932
	}

if_then931:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end932:
	v434 = *libc.As[byte](result)
	loadedv933 = (v434 & 1) != 0
	*libc.As[bool](retval) = loadedv933
	goto _return

sw_bb934:
	*libc.As[byte](result) = 1
	v435 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol935 = libc.Ptr(&libc.As[TSLexer](v435).F1)
	*libc.As[int16](result_symbol935) = 11
	v436 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end936 = libc.Ptr(&libc.As[TSLexer](v436).F3)
	v437 = *libc.As[unsafe.Pointer](mark_end936)
	v438 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v437)(v438)
	v439 = *libc.As[int32](lookahead)
	cmp937 = v439 == 59
	if cmp937 {
		goto if_then939
	} else {
		goto if_end940
	}

if_then939:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end940:
	v440 = *libc.As[int32](lookahead)
	cmp941 = 65 <= v440
	if cmp941 {
		goto land_lhs_true943
	} else {
		goto lor_lhs_false946
	}

land_lhs_true943:
	v441 = *libc.As[int32](lookahead)
	cmp944 = v441 <= 90
	if cmp944 {
		goto if_then952
	} else {
		goto lor_lhs_false946
	}

lor_lhs_false946:
	v442 = *libc.As[int32](lookahead)
	cmp947 = 97 <= v442
	if cmp947 {
		goto land_lhs_true949
	} else {
		goto if_end953
	}

land_lhs_true949:
	v443 = *libc.As[int32](lookahead)
	cmp950 = v443 <= 122
	if cmp950 {
		goto if_then952
	} else {
		goto if_end953
	}

if_then952:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end953:
	v444 = *libc.As[byte](result)
	loadedv954 = (v444 & 1) != 0
	*libc.As[bool](retval) = loadedv954
	goto _return

sw_bb955:
	*libc.As[byte](result) = 1
	v445 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol956 = libc.Ptr(&libc.As[TSLexer](v445).F1)
	*libc.As[int16](result_symbol956) = 11
	v446 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end957 = libc.Ptr(&libc.As[TSLexer](v446).F3)
	v447 = *libc.As[unsafe.Pointer](mark_end957)
	v448 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v447)(v448)
	v449 = *libc.As[int32](lookahead)
	cmp958 = v449 == 59
	if cmp958 {
		goto if_then960
	} else {
		goto if_end961
	}

if_then960:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end961:
	v450 = *libc.As[int32](lookahead)
	cmp962 = 65 <= v450
	if cmp962 {
		goto land_lhs_true964
	} else {
		goto lor_lhs_false967
	}

land_lhs_true964:
	v451 = *libc.As[int32](lookahead)
	cmp965 = v451 <= 90
	if cmp965 {
		goto if_then973
	} else {
		goto lor_lhs_false967
	}

lor_lhs_false967:
	v452 = *libc.As[int32](lookahead)
	cmp968 = 97 <= v452
	if cmp968 {
		goto land_lhs_true970
	} else {
		goto if_end974
	}

land_lhs_true970:
	v453 = *libc.As[int32](lookahead)
	cmp971 = v453 <= 122
	if cmp971 {
		goto if_then973
	} else {
		goto if_end974
	}

if_then973:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end974:
	v454 = *libc.As[byte](result)
	loadedv975 = (v454 & 1) != 0
	*libc.As[bool](retval) = loadedv975
	goto _return

sw_bb976:
	*libc.As[byte](result) = 1
	v455 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol977 = libc.Ptr(&libc.As[TSLexer](v455).F1)
	*libc.As[int16](result_symbol977) = 11
	v456 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end978 = libc.Ptr(&libc.As[TSLexer](v456).F3)
	v457 = *libc.As[unsafe.Pointer](mark_end978)
	v458 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v457)(v458)
	v459 = *libc.As[int32](lookahead)
	cmp979 = v459 == 59
	if cmp979 {
		goto if_then981
	} else {
		goto if_end982
	}

if_then981:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end982:
	v460 = *libc.As[int32](lookahead)
	cmp983 = 65 <= v460
	if cmp983 {
		goto land_lhs_true985
	} else {
		goto lor_lhs_false988
	}

land_lhs_true985:
	v461 = *libc.As[int32](lookahead)
	cmp986 = v461 <= 90
	if cmp986 {
		goto if_then994
	} else {
		goto lor_lhs_false988
	}

lor_lhs_false988:
	v462 = *libc.As[int32](lookahead)
	cmp989 = 97 <= v462
	if cmp989 {
		goto land_lhs_true991
	} else {
		goto if_end995
	}

land_lhs_true991:
	v463 = *libc.As[int32](lookahead)
	cmp992 = v463 <= 122
	if cmp992 {
		goto if_then994
	} else {
		goto if_end995
	}

if_then994:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end995:
	v464 = *libc.As[byte](result)
	loadedv996 = (v464 & 1) != 0
	*libc.As[bool](retval) = loadedv996
	goto _return

sw_bb997:
	*libc.As[byte](result) = 1
	v465 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol998 = libc.Ptr(&libc.As[TSLexer](v465).F1)
	*libc.As[int16](result_symbol998) = 11
	v466 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end999 = libc.Ptr(&libc.As[TSLexer](v466).F3)
	v467 = *libc.As[unsafe.Pointer](mark_end999)
	v468 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v467)(v468)
	v469 = *libc.As[int32](lookahead)
	cmp1000 = v469 == 59
	if cmp1000 {
		goto if_then1002
	} else {
		goto if_end1003
	}

if_then1002:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1003:
	v470 = *libc.As[int32](lookahead)
	cmp1004 = 65 <= v470
	if cmp1004 {
		goto land_lhs_true1006
	} else {
		goto lor_lhs_false1009
	}

land_lhs_true1006:
	v471 = *libc.As[int32](lookahead)
	cmp1007 = v471 <= 90
	if cmp1007 {
		goto if_then1015
	} else {
		goto lor_lhs_false1009
	}

lor_lhs_false1009:
	v472 = *libc.As[int32](lookahead)
	cmp1010 = 97 <= v472
	if cmp1010 {
		goto land_lhs_true1012
	} else {
		goto if_end1016
	}

land_lhs_true1012:
	v473 = *libc.As[int32](lookahead)
	cmp1013 = v473 <= 122
	if cmp1013 {
		goto if_then1015
	} else {
		goto if_end1016
	}

if_then1015:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end1016:
	v474 = *libc.As[byte](result)
	loadedv1017 = (v474 & 1) != 0
	*libc.As[bool](retval) = loadedv1017
	goto _return

sw_bb1018:
	*libc.As[byte](result) = 1
	v475 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1019 = libc.Ptr(&libc.As[TSLexer](v475).F1)
	*libc.As[int16](result_symbol1019) = 11
	v476 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1020 = libc.Ptr(&libc.As[TSLexer](v476).F3)
	v477 = *libc.As[unsafe.Pointer](mark_end1020)
	v478 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v477)(v478)
	v479 = *libc.As[int32](lookahead)
	cmp1021 = v479 == 59
	if cmp1021 {
		goto if_then1023
	} else {
		goto if_end1024
	}

if_then1023:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1024:
	v480 = *libc.As[int32](lookahead)
	cmp1025 = 65 <= v480
	if cmp1025 {
		goto land_lhs_true1027
	} else {
		goto lor_lhs_false1030
	}

land_lhs_true1027:
	v481 = *libc.As[int32](lookahead)
	cmp1028 = v481 <= 90
	if cmp1028 {
		goto if_then1036
	} else {
		goto lor_lhs_false1030
	}

lor_lhs_false1030:
	v482 = *libc.As[int32](lookahead)
	cmp1031 = 97 <= v482
	if cmp1031 {
		goto land_lhs_true1033
	} else {
		goto if_end1037
	}

land_lhs_true1033:
	v483 = *libc.As[int32](lookahead)
	cmp1034 = v483 <= 122
	if cmp1034 {
		goto if_then1036
	} else {
		goto if_end1037
	}

if_then1036:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end1037:
	v484 = *libc.As[byte](result)
	loadedv1038 = (v484 & 1) != 0
	*libc.As[bool](retval) = loadedv1038
	goto _return

sw_bb1039:
	*libc.As[byte](result) = 1
	v485 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1040 = libc.Ptr(&libc.As[TSLexer](v485).F1)
	*libc.As[int16](result_symbol1040) = 11
	v486 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1041 = libc.Ptr(&libc.As[TSLexer](v486).F3)
	v487 = *libc.As[unsafe.Pointer](mark_end1041)
	v488 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v487)(v488)
	v489 = *libc.As[int32](lookahead)
	cmp1042 = v489 == 59
	if cmp1042 {
		goto if_then1044
	} else {
		goto if_end1045
	}

if_then1044:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1045:
	v490 = *libc.As[int32](lookahead)
	cmp1046 = 65 <= v490
	if cmp1046 {
		goto land_lhs_true1048
	} else {
		goto lor_lhs_false1051
	}

land_lhs_true1048:
	v491 = *libc.As[int32](lookahead)
	cmp1049 = v491 <= 90
	if cmp1049 {
		goto if_then1057
	} else {
		goto lor_lhs_false1051
	}

lor_lhs_false1051:
	v492 = *libc.As[int32](lookahead)
	cmp1052 = 97 <= v492
	if cmp1052 {
		goto land_lhs_true1054
	} else {
		goto if_end1058
	}

land_lhs_true1054:
	v493 = *libc.As[int32](lookahead)
	cmp1055 = v493 <= 122
	if cmp1055 {
		goto if_then1057
	} else {
		goto if_end1058
	}

if_then1057:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end1058:
	v494 = *libc.As[byte](result)
	loadedv1059 = (v494 & 1) != 0
	*libc.As[bool](retval) = loadedv1059
	goto _return

sw_bb1060:
	*libc.As[byte](result) = 1
	v495 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1061 = libc.Ptr(&libc.As[TSLexer](v495).F1)
	*libc.As[int16](result_symbol1061) = 11
	v496 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1062 = libc.Ptr(&libc.As[TSLexer](v496).F3)
	v497 = *libc.As[unsafe.Pointer](mark_end1062)
	v498 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v497)(v498)
	v499 = *libc.As[int32](lookahead)
	cmp1063 = v499 == 59
	if cmp1063 {
		goto if_then1065
	} else {
		goto if_end1066
	}

if_then1065:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1066:
	v500 = *libc.As[int32](lookahead)
	cmp1067 = 65 <= v500
	if cmp1067 {
		goto land_lhs_true1069
	} else {
		goto lor_lhs_false1072
	}

land_lhs_true1069:
	v501 = *libc.As[int32](lookahead)
	cmp1070 = v501 <= 90
	if cmp1070 {
		goto if_then1078
	} else {
		goto lor_lhs_false1072
	}

lor_lhs_false1072:
	v502 = *libc.As[int32](lookahead)
	cmp1073 = 97 <= v502
	if cmp1073 {
		goto land_lhs_true1075
	} else {
		goto if_end1079
	}

land_lhs_true1075:
	v503 = *libc.As[int32](lookahead)
	cmp1076 = v503 <= 122
	if cmp1076 {
		goto if_then1078
	} else {
		goto if_end1079
	}

if_then1078:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end1079:
	v504 = *libc.As[byte](result)
	loadedv1080 = (v504 & 1) != 0
	*libc.As[bool](retval) = loadedv1080
	goto _return

sw_bb1081:
	*libc.As[byte](result) = 1
	v505 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1082 = libc.Ptr(&libc.As[TSLexer](v505).F1)
	*libc.As[int16](result_symbol1082) = 11
	v506 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1083 = libc.Ptr(&libc.As[TSLexer](v506).F3)
	v507 = *libc.As[unsafe.Pointer](mark_end1083)
	v508 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v507)(v508)
	v509 = *libc.As[int32](lookahead)
	cmp1084 = v509 == 59
	if cmp1084 {
		goto if_then1086
	} else {
		goto if_end1087
	}

if_then1086:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1087:
	v510 = *libc.As[int32](lookahead)
	cmp1088 = 65 <= v510
	if cmp1088 {
		goto land_lhs_true1090
	} else {
		goto lor_lhs_false1093
	}

land_lhs_true1090:
	v511 = *libc.As[int32](lookahead)
	cmp1091 = v511 <= 90
	if cmp1091 {
		goto if_then1099
	} else {
		goto lor_lhs_false1093
	}

lor_lhs_false1093:
	v512 = *libc.As[int32](lookahead)
	cmp1094 = 97 <= v512
	if cmp1094 {
		goto land_lhs_true1096
	} else {
		goto if_end1100
	}

land_lhs_true1096:
	v513 = *libc.As[int32](lookahead)
	cmp1097 = v513 <= 122
	if cmp1097 {
		goto if_then1099
	} else {
		goto if_end1100
	}

if_then1099:
	*libc.As[int16](state_addr) = 60
	goto next_state

if_end1100:
	v514 = *libc.As[byte](result)
	loadedv1101 = (v514 & 1) != 0
	*libc.As[bool](retval) = loadedv1101
	goto _return

sw_bb1102:
	*libc.As[byte](result) = 1
	v515 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1103 = libc.Ptr(&libc.As[TSLexer](v515).F1)
	*libc.As[int16](result_symbol1103) = 11
	v516 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1104 = libc.Ptr(&libc.As[TSLexer](v516).F3)
	v517 = *libc.As[unsafe.Pointer](mark_end1104)
	v518 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v517)(v518)
	v519 = *libc.As[int32](lookahead)
	cmp1105 = v519 == 59
	if cmp1105 {
		goto if_then1107
	} else {
		goto if_end1108
	}

if_then1107:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1108:
	v520 = *libc.As[int32](lookahead)
	cmp1109 = 65 <= v520
	if cmp1109 {
		goto land_lhs_true1111
	} else {
		goto lor_lhs_false1114
	}

land_lhs_true1111:
	v521 = *libc.As[int32](lookahead)
	cmp1112 = v521 <= 90
	if cmp1112 {
		goto if_then1120
	} else {
		goto lor_lhs_false1114
	}

lor_lhs_false1114:
	v522 = *libc.As[int32](lookahead)
	cmp1115 = 97 <= v522
	if cmp1115 {
		goto land_lhs_true1117
	} else {
		goto if_end1121
	}

land_lhs_true1117:
	v523 = *libc.As[int32](lookahead)
	cmp1118 = v523 <= 122
	if cmp1118 {
		goto if_then1120
	} else {
		goto if_end1121
	}

if_then1120:
	*libc.As[int16](state_addr) = 61
	goto next_state

if_end1121:
	v524 = *libc.As[byte](result)
	loadedv1122 = (v524 & 1) != 0
	*libc.As[bool](retval) = loadedv1122
	goto _return

sw_bb1123:
	*libc.As[byte](result) = 1
	v525 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1124 = libc.Ptr(&libc.As[TSLexer](v525).F1)
	*libc.As[int16](result_symbol1124) = 11
	v526 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1125 = libc.Ptr(&libc.As[TSLexer](v526).F3)
	v527 = *libc.As[unsafe.Pointer](mark_end1125)
	v528 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v527)(v528)
	v529 = *libc.As[int32](lookahead)
	cmp1126 = v529 == 59
	if cmp1126 {
		goto if_then1128
	} else {
		goto if_end1129
	}

if_then1128:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1129:
	v530 = *libc.As[int32](lookahead)
	cmp1130 = 65 <= v530
	if cmp1130 {
		goto land_lhs_true1132
	} else {
		goto lor_lhs_false1135
	}

land_lhs_true1132:
	v531 = *libc.As[int32](lookahead)
	cmp1133 = v531 <= 90
	if cmp1133 {
		goto if_then1141
	} else {
		goto lor_lhs_false1135
	}

lor_lhs_false1135:
	v532 = *libc.As[int32](lookahead)
	cmp1136 = 97 <= v532
	if cmp1136 {
		goto land_lhs_true1138
	} else {
		goto if_end1142
	}

land_lhs_true1138:
	v533 = *libc.As[int32](lookahead)
	cmp1139 = v533 <= 122
	if cmp1139 {
		goto if_then1141
	} else {
		goto if_end1142
	}

if_then1141:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end1142:
	v534 = *libc.As[byte](result)
	loadedv1143 = (v534 & 1) != 0
	*libc.As[bool](retval) = loadedv1143
	goto _return

sw_bb1144:
	*libc.As[byte](result) = 1
	v535 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1145 = libc.Ptr(&libc.As[TSLexer](v535).F1)
	*libc.As[int16](result_symbol1145) = 11
	v536 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1146 = libc.Ptr(&libc.As[TSLexer](v536).F3)
	v537 = *libc.As[unsafe.Pointer](mark_end1146)
	v538 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v537)(v538)
	v539 = *libc.As[int32](lookahead)
	cmp1147 = v539 == 59
	if cmp1147 {
		goto if_then1149
	} else {
		goto if_end1150
	}

if_then1149:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1150:
	v540 = *libc.As[int32](lookahead)
	cmp1151 = 65 <= v540
	if cmp1151 {
		goto land_lhs_true1153
	} else {
		goto lor_lhs_false1156
	}

land_lhs_true1153:
	v541 = *libc.As[int32](lookahead)
	cmp1154 = v541 <= 90
	if cmp1154 {
		goto if_then1162
	} else {
		goto lor_lhs_false1156
	}

lor_lhs_false1156:
	v542 = *libc.As[int32](lookahead)
	cmp1157 = 97 <= v542
	if cmp1157 {
		goto land_lhs_true1159
	} else {
		goto if_end1163
	}

land_lhs_true1159:
	v543 = *libc.As[int32](lookahead)
	cmp1160 = v543 <= 122
	if cmp1160 {
		goto if_then1162
	} else {
		goto if_end1163
	}

if_then1162:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end1163:
	v544 = *libc.As[byte](result)
	loadedv1164 = (v544 & 1) != 0
	*libc.As[bool](retval) = loadedv1164
	goto _return

sw_bb1165:
	*libc.As[byte](result) = 1
	v545 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1166 = libc.Ptr(&libc.As[TSLexer](v545).F1)
	*libc.As[int16](result_symbol1166) = 11
	v546 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1167 = libc.Ptr(&libc.As[TSLexer](v546).F3)
	v547 = *libc.As[unsafe.Pointer](mark_end1167)
	v548 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v547)(v548)
	v549 = *libc.As[int32](lookahead)
	cmp1168 = v549 == 59
	if cmp1168 {
		goto if_then1170
	} else {
		goto if_end1171
	}

if_then1170:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1171:
	v550 = *libc.As[int32](lookahead)
	cmp1172 = 65 <= v550
	if cmp1172 {
		goto land_lhs_true1174
	} else {
		goto lor_lhs_false1177
	}

land_lhs_true1174:
	v551 = *libc.As[int32](lookahead)
	cmp1175 = v551 <= 90
	if cmp1175 {
		goto if_then1183
	} else {
		goto lor_lhs_false1177
	}

lor_lhs_false1177:
	v552 = *libc.As[int32](lookahead)
	cmp1178 = 97 <= v552
	if cmp1178 {
		goto land_lhs_true1180
	} else {
		goto if_end1184
	}

land_lhs_true1180:
	v553 = *libc.As[int32](lookahead)
	cmp1181 = v553 <= 122
	if cmp1181 {
		goto if_then1183
	} else {
		goto if_end1184
	}

if_then1183:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end1184:
	v554 = *libc.As[byte](result)
	loadedv1185 = (v554 & 1) != 0
	*libc.As[bool](retval) = loadedv1185
	goto _return

sw_bb1186:
	*libc.As[byte](result) = 1
	v555 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1187 = libc.Ptr(&libc.As[TSLexer](v555).F1)
	*libc.As[int16](result_symbol1187) = 11
	v556 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1188 = libc.Ptr(&libc.As[TSLexer](v556).F3)
	v557 = *libc.As[unsafe.Pointer](mark_end1188)
	v558 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v557)(v558)
	v559 = *libc.As[int32](lookahead)
	cmp1189 = v559 == 59
	if cmp1189 {
		goto if_then1191
	} else {
		goto if_end1192
	}

if_then1191:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1192:
	v560 = *libc.As[int32](lookahead)
	cmp1193 = 65 <= v560
	if cmp1193 {
		goto land_lhs_true1195
	} else {
		goto lor_lhs_false1198
	}

land_lhs_true1195:
	v561 = *libc.As[int32](lookahead)
	cmp1196 = v561 <= 90
	if cmp1196 {
		goto if_then1204
	} else {
		goto lor_lhs_false1198
	}

lor_lhs_false1198:
	v562 = *libc.As[int32](lookahead)
	cmp1199 = 97 <= v562
	if cmp1199 {
		goto land_lhs_true1201
	} else {
		goto if_end1205
	}

land_lhs_true1201:
	v563 = *libc.As[int32](lookahead)
	cmp1202 = v563 <= 122
	if cmp1202 {
		goto if_then1204
	} else {
		goto if_end1205
	}

if_then1204:
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end1205:
	v564 = *libc.As[byte](result)
	loadedv1206 = (v564 & 1) != 0
	*libc.As[bool](retval) = loadedv1206
	goto _return

sw_bb1207:
	*libc.As[byte](result) = 1
	v565 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1208 = libc.Ptr(&libc.As[TSLexer](v565).F1)
	*libc.As[int16](result_symbol1208) = 11
	v566 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1209 = libc.Ptr(&libc.As[TSLexer](v566).F3)
	v567 = *libc.As[unsafe.Pointer](mark_end1209)
	v568 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v567)(v568)
	v569 = *libc.As[int32](lookahead)
	cmp1210 = v569 == 59
	if cmp1210 {
		goto if_then1212
	} else {
		goto if_end1213
	}

if_then1212:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1213:
	v570 = *libc.As[int32](lookahead)
	cmp1214 = 65 <= v570
	if cmp1214 {
		goto land_lhs_true1216
	} else {
		goto lor_lhs_false1219
	}

land_lhs_true1216:
	v571 = *libc.As[int32](lookahead)
	cmp1217 = v571 <= 90
	if cmp1217 {
		goto if_then1225
	} else {
		goto lor_lhs_false1219
	}

lor_lhs_false1219:
	v572 = *libc.As[int32](lookahead)
	cmp1220 = 97 <= v572
	if cmp1220 {
		goto land_lhs_true1222
	} else {
		goto if_end1226
	}

land_lhs_true1222:
	v573 = *libc.As[int32](lookahead)
	cmp1223 = v573 <= 122
	if cmp1223 {
		goto if_then1225
	} else {
		goto if_end1226
	}

if_then1225:
	*libc.As[int16](state_addr) = 66
	goto next_state

if_end1226:
	v574 = *libc.As[byte](result)
	loadedv1227 = (v574 & 1) != 0
	*libc.As[bool](retval) = loadedv1227
	goto _return

sw_bb1228:
	*libc.As[byte](result) = 1
	v575 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1229 = libc.Ptr(&libc.As[TSLexer](v575).F1)
	*libc.As[int16](result_symbol1229) = 11
	v576 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1230 = libc.Ptr(&libc.As[TSLexer](v576).F3)
	v577 = *libc.As[unsafe.Pointer](mark_end1230)
	v578 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v577)(v578)
	v579 = *libc.As[int32](lookahead)
	cmp1231 = v579 == 59
	if cmp1231 {
		goto if_then1233
	} else {
		goto if_end1234
	}

if_then1233:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1234:
	v580 = *libc.As[int32](lookahead)
	cmp1235 = 65 <= v580
	if cmp1235 {
		goto land_lhs_true1237
	} else {
		goto lor_lhs_false1240
	}

land_lhs_true1237:
	v581 = *libc.As[int32](lookahead)
	cmp1238 = v581 <= 90
	if cmp1238 {
		goto if_then1246
	} else {
		goto lor_lhs_false1240
	}

lor_lhs_false1240:
	v582 = *libc.As[int32](lookahead)
	cmp1241 = 97 <= v582
	if cmp1241 {
		goto land_lhs_true1243
	} else {
		goto if_end1247
	}

land_lhs_true1243:
	v583 = *libc.As[int32](lookahead)
	cmp1244 = v583 <= 122
	if cmp1244 {
		goto if_then1246
	} else {
		goto if_end1247
	}

if_then1246:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end1247:
	v584 = *libc.As[byte](result)
	loadedv1248 = (v584 & 1) != 0
	*libc.As[bool](retval) = loadedv1248
	goto _return

sw_bb1249:
	*libc.As[byte](result) = 1
	v585 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1250 = libc.Ptr(&libc.As[TSLexer](v585).F1)
	*libc.As[int16](result_symbol1250) = 11
	v586 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1251 = libc.Ptr(&libc.As[TSLexer](v586).F3)
	v587 = *libc.As[unsafe.Pointer](mark_end1251)
	v588 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v587)(v588)
	v589 = *libc.As[int32](lookahead)
	cmp1252 = v589 == 59
	if cmp1252 {
		goto if_then1254
	} else {
		goto if_end1255
	}

if_then1254:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1255:
	v590 = *libc.As[int32](lookahead)
	cmp1256 = 65 <= v590
	if cmp1256 {
		goto land_lhs_true1258
	} else {
		goto lor_lhs_false1261
	}

land_lhs_true1258:
	v591 = *libc.As[int32](lookahead)
	cmp1259 = v591 <= 90
	if cmp1259 {
		goto if_then1267
	} else {
		goto lor_lhs_false1261
	}

lor_lhs_false1261:
	v592 = *libc.As[int32](lookahead)
	cmp1262 = 97 <= v592
	if cmp1262 {
		goto land_lhs_true1264
	} else {
		goto if_end1268
	}

land_lhs_true1264:
	v593 = *libc.As[int32](lookahead)
	cmp1265 = v593 <= 122
	if cmp1265 {
		goto if_then1267
	} else {
		goto if_end1268
	}

if_then1267:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end1268:
	v594 = *libc.As[byte](result)
	loadedv1269 = (v594 & 1) != 0
	*libc.As[bool](retval) = loadedv1269
	goto _return

sw_bb1270:
	*libc.As[byte](result) = 1
	v595 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1271 = libc.Ptr(&libc.As[TSLexer](v595).F1)
	*libc.As[int16](result_symbol1271) = 12
	v596 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1272 = libc.Ptr(&libc.As[TSLexer](v596).F3)
	v597 = *libc.As[unsafe.Pointer](mark_end1272)
	v598 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v597)(v598)
	v599 = *libc.As[byte](result)
	loadedv1273 = (v599 & 1) != 0
	*libc.As[bool](retval) = loadedv1273
	goto _return

sw_bb1274:
	*libc.As[byte](result) = 1
	v600 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1275 = libc.Ptr(&libc.As[TSLexer](v600).F1)
	*libc.As[int16](result_symbol1275) = 13
	v601 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1276 = libc.Ptr(&libc.As[TSLexer](v601).F3)
	v602 = *libc.As[unsafe.Pointer](mark_end1276)
	v603 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v602)(v603)
	v604 = *libc.As[int32](lookahead)
	cmp1277 = 9 <= v604
	if cmp1277 {
		goto land_lhs_true1279
	} else {
		goto lor_lhs_false1282
	}

land_lhs_true1279:
	v605 = *libc.As[int32](lookahead)
	cmp1280 = v605 <= 13
	if cmp1280 {
		goto if_then1285
	} else {
		goto lor_lhs_false1282
	}

lor_lhs_false1282:
	v606 = *libc.As[int32](lookahead)
	cmp1283 = v606 == 32
	if cmp1283 {
		goto if_then1285
	} else {
		goto if_end1286
	}

if_then1285:
	*libc.As[int16](state_addr) = 71
	goto next_state

if_end1286:
	v607 = *libc.As[int32](lookahead)
	cmp1287 = v607 != 0
	if cmp1287 {
		goto land_lhs_true1289
	} else {
		goto if_end1293
	}

land_lhs_true1289:
	v608 = *libc.As[int32](lookahead)
	cmp1290 = v608 != 39
	if cmp1290 {
		goto if_then1292
	} else {
		goto if_end1293
	}

if_then1292:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1293:
	v609 = *libc.As[byte](result)
	loadedv1294 = (v609 & 1) != 0
	*libc.As[bool](retval) = loadedv1294
	goto _return

sw_bb1295:
	*libc.As[byte](result) = 1
	v610 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1296 = libc.Ptr(&libc.As[TSLexer](v610).F1)
	*libc.As[int16](result_symbol1296) = 13
	v611 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1297 = libc.Ptr(&libc.As[TSLexer](v611).F3)
	v612 = *libc.As[unsafe.Pointer](mark_end1297)
	v613 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v612)(v613)
	v614 = *libc.As[int32](lookahead)
	cmp1298 = v614 != 0
	if cmp1298 {
		goto land_lhs_true1300
	} else {
		goto if_end1304
	}

land_lhs_true1300:
	v615 = *libc.As[int32](lookahead)
	cmp1301 = v615 != 39
	if cmp1301 {
		goto if_then1303
	} else {
		goto if_end1304
	}

if_then1303:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1304:
	v616 = *libc.As[byte](result)
	loadedv1305 = (v616 & 1) != 0
	*libc.As[bool](retval) = loadedv1305
	goto _return

sw_bb1306:
	*libc.As[byte](result) = 1
	v617 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1307 = libc.Ptr(&libc.As[TSLexer](v617).F1)
	*libc.As[int16](result_symbol1307) = 14
	v618 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1308 = libc.Ptr(&libc.As[TSLexer](v618).F3)
	v619 = *libc.As[unsafe.Pointer](mark_end1308)
	v620 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v619)(v620)
	v621 = *libc.As[byte](result)
	loadedv1309 = (v621 & 1) != 0
	*libc.As[bool](retval) = loadedv1309
	goto _return

sw_bb1310:
	*libc.As[byte](result) = 1
	v622 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1311 = libc.Ptr(&libc.As[TSLexer](v622).F1)
	*libc.As[int16](result_symbol1311) = 15
	v623 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1312 = libc.Ptr(&libc.As[TSLexer](v623).F3)
	v624 = *libc.As[unsafe.Pointer](mark_end1312)
	v625 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v624)(v625)
	v626 = *libc.As[int32](lookahead)
	cmp1313 = 9 <= v626
	if cmp1313 {
		goto land_lhs_true1315
	} else {
		goto lor_lhs_false1318
	}

land_lhs_true1315:
	v627 = *libc.As[int32](lookahead)
	cmp1316 = v627 <= 13
	if cmp1316 {
		goto if_then1321
	} else {
		goto lor_lhs_false1318
	}

lor_lhs_false1318:
	v628 = *libc.As[int32](lookahead)
	cmp1319 = v628 == 32
	if cmp1319 {
		goto if_then1321
	} else {
		goto if_end1322
	}

if_then1321:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end1322:
	v629 = *libc.As[int32](lookahead)
	cmp1323 = v629 != 0
	if cmp1323 {
		goto land_lhs_true1325
	} else {
		goto if_end1329
	}

land_lhs_true1325:
	v630 = *libc.As[int32](lookahead)
	cmp1326 = v630 != 34
	if cmp1326 {
		goto if_then1328
	} else {
		goto if_end1329
	}

if_then1328:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end1329:
	v631 = *libc.As[byte](result)
	loadedv1330 = (v631 & 1) != 0
	*libc.As[bool](retval) = loadedv1330
	goto _return

sw_bb1331:
	*libc.As[byte](result) = 1
	v632 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1332 = libc.Ptr(&libc.As[TSLexer](v632).F1)
	*libc.As[int16](result_symbol1332) = 15
	v633 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1333 = libc.Ptr(&libc.As[TSLexer](v633).F3)
	v634 = *libc.As[unsafe.Pointer](mark_end1333)
	v635 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v634)(v635)
	v636 = *libc.As[int32](lookahead)
	cmp1334 = v636 != 0
	if cmp1334 {
		goto land_lhs_true1336
	} else {
		goto if_end1340
	}

land_lhs_true1336:
	v637 = *libc.As[int32](lookahead)
	cmp1337 = v637 != 34
	if cmp1337 {
		goto if_then1339
	} else {
		goto if_end1340
	}

if_then1339:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end1340:
	v638 = *libc.As[byte](result)
	loadedv1341 = (v638 & 1) != 0
	*libc.As[bool](retval) = loadedv1341
	goto _return

sw_bb1342:
	*libc.As[byte](result) = 1
	v639 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1343 = libc.Ptr(&libc.As[TSLexer](v639).F1)
	*libc.As[int16](result_symbol1343) = 16
	v640 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1344 = libc.Ptr(&libc.As[TSLexer](v640).F3)
	v641 = *libc.As[unsafe.Pointer](mark_end1344)
	v642 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v641)(v642)
	v643 = *libc.As[int32](lookahead)
	cmp1345 = 9 <= v643
	if cmp1345 {
		goto land_lhs_true1347
	} else {
		goto lor_lhs_false1350
	}

land_lhs_true1347:
	v644 = *libc.As[int32](lookahead)
	cmp1348 = v644 <= 13
	if cmp1348 {
		goto if_then1353
	} else {
		goto lor_lhs_false1350
	}

lor_lhs_false1350:
	v645 = *libc.As[int32](lookahead)
	cmp1351 = v645 == 32
	if cmp1351 {
		goto if_then1353
	} else {
		goto if_end1354
	}

if_then1353:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end1354:
	v646 = *libc.As[int32](lookahead)
	cmp1355 = v646 != 0
	if cmp1355 {
		goto land_lhs_true1357
	} else {
		goto if_end1367
	}

land_lhs_true1357:
	v647 = *libc.As[int32](lookahead)
	cmp1358 = v647 != 38
	if cmp1358 {
		goto land_lhs_true1360
	} else {
		goto if_end1367
	}

land_lhs_true1360:
	v648 = *libc.As[int32](lookahead)
	cmp1361 = v648 != 60
	if cmp1361 {
		goto land_lhs_true1363
	} else {
		goto if_end1367
	}

land_lhs_true1363:
	v649 = *libc.As[int32](lookahead)
	cmp1364 = v649 != 62
	if cmp1364 {
		goto if_then1366
	} else {
		goto if_end1367
	}

if_then1366:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end1367:
	v650 = *libc.As[byte](result)
	loadedv1368 = (v650 & 1) != 0
	*libc.As[bool](retval) = loadedv1368
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v651 = *libc.As[bool](retval)
	return v651
}
func scan_raw_text(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var arrayidx unsafe.Pointer
	var tags, tags1, tags3, tags8, tags9 unsafe.Pointer
	var cond unsafe.Pointer
	var cmp, cmp5, cmp12, tobool, cmp16, cmp21, v31 bool
	var retval unsafe.Pointer
	var result_symbol unsafe.Pointer
	var v1, v6, sub, v8, v12, sub11, v13, v16, v18, call, v20, conv, v22, inc, v23 int32
	var delimiter_index, size, size2, size4, size10, _type, lookahead, lookahead13 unsafe.Pointer
	var idxprom, v14, idxprom14, conv19, call20 int64
	var v21 byte
	var arrayidx15 unsafe.Pointer
	var v0, v2, v3, v4, v5, v7, v9, v10, v11, v15, v17, v19, v24, v25, v26, v27, v28, v29, v30 unsafe.Pointer
	var scanner_addr, lexer_addr, end_delimiter, mark_end, contents, mark_end26 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, end_delimiter, delimiter_index, v0, tags, size, v1, cmp, v2, mark_end, v3, v4, v5, tags1, size2, v6, sub, v7, tags3, size4, v8, cmp5, v9, tags8, contents, v10, v11, tags9, size10, v12, sub11, idxprom, arrayidx, _type, v13, cmp12, v14, cond, v15, lookahead, v16, tobool, v17, lookahead13, v18, call, v19, v20, idxprom14, arrayidx15, v21, conv, cmp16, v22, inc, v23, conv19, v24, call20, cmp21, v25, v26, v27, mark_end26, v28, v29, v30, result_symbol, v31

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
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v2).F3)
	v3 = *libc.As[unsafe.Pointer](mark_end)
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v3)(v4)
	v5 = *libc.As[unsafe.Pointer](scanner_addr)
	tags1 = libc.Ptr(&libc.As[Scanner](v5).F0)
	size2 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F1)
	v6 = *libc.As[int32](size2)
	sub = v6 - 1
	v7 = *libc.As[unsafe.Pointer](scanner_addr)
	tags3 = libc.Ptr(&libc.As[Scanner](v7).F0)
	size4 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags3).F1)
	v8 = *libc.As[int32](size4)
	cmp5 = uint32(sub) < uint32(v8)
	if cmp5 {
		goto if_then6
	} else {
		goto if_else
	}

if_then6:
	goto if_end7

if_else:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 150, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_raw_text)))
	panic("unreachable")

if_end7:
	v9 = *libc.As[unsafe.Pointer](scanner_addr)
	tags8 = libc.Ptr(&libc.As[Scanner](v9).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags8).F0)
	v10 = *libc.As[unsafe.Pointer](contents)
	v11 = *libc.As[unsafe.Pointer](scanner_addr)
	tags9 = libc.Ptr(&libc.As[Scanner](v11).F0)
	size10 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags9).F1)
	v12 = *libc.As[int32](size10)
	sub11 = v12 - 1
	idxprom = int64(uint64(uint32(sub11)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v10), int(idxprom)*24))
	_type = libc.Ptr(&libc.As[Tag](arrayidx).F0)
	v13 = *libc.As[int32](_type)
	cmp12 = v13 == 99
	if cmp12 {
		v14 = 1
	} else {
		v14 = 0
	}
	if cmp12 {
		cond = libc.Ptr(&_str_2)
	} else {
		cond = libc.Ptr(&_str_3)
	}
	*libc.As[unsafe.Pointer](end_delimiter) = cond
	*libc.As[int32](delimiter_index) = 0
	goto while_cond

while_cond:
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v15).F0)
	v16 = *libc.As[int32](lookahead)
	tobool = v16 != 0
	if tobool {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead13 = libc.Ptr(&libc.As[TSLexer](v17).F0)
	v18 = *libc.As[int32](lookahead13)
	call = libc.Towupper(v18)
	v19 = *libc.As[unsafe.Pointer](end_delimiter)
	v20 = *libc.As[int32](delimiter_index)
	idxprom14 = int64(uint64(uint32(v20)))
	arrayidx15 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v19), int(idxprom14)*1))
	v21 = *libc.As[byte](arrayidx15)
	conv = int32(int8(v21))
	cmp16 = call == conv
	if cmp16 {
		goto if_then18
	} else {
		goto if_else25
	}

if_then18:
	v22 = *libc.As[int32](delimiter_index)
	inc = v22 + 1
	*libc.As[int32](delimiter_index) = inc
	v23 = *libc.As[int32](delimiter_index)
	conv19 = int64(uint64(uint32(v23)))
	v24 = *libc.As[unsafe.Pointer](end_delimiter)
	call20 = libc.Strlen(libc.As[byte](v24))
	cmp21 = conv19 == call20
	if cmp21 {
		goto if_then23
	} else {
		goto if_end24
	}

if_then23:
	goto while_end

if_end24:
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v25)
	goto if_end27

if_else25:
	*libc.As[int32](delimiter_index) = 0
	v26 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v26)
	v27 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end26 = libc.Ptr(&libc.As[TSLexer](v27).F3)
	v28 = *libc.As[unsafe.Pointer](mark_end26)
	v29 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v28)(v29)
	goto if_end27

if_end27:
	goto while_cond

while_end:
	v30 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v30).F1)
	*libc.As[int16](result_symbol) = 7
	*libc.As[bool](retval) = true
	goto _return

_return:
	v31 = *libc.As[bool](retval)
	return v31
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
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 177, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_implicit_end_tag)))
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
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 201, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_implicit_end_tag)))
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
	cmp65 = v55 == 66
	if cmp65 {
		goto land_lhs_true72
	} else {
		goto lor_lhs_false66
	}

lor_lhs_false66:
	v56 = *libc.As[unsafe.Pointer](parent)
	type67 = libc.Ptr(&libc.As[Tag](v56).F0)
	v57 = *libc.As[int32](type67)
	cmp68 = v57 == 63
	if cmp68 {
		goto land_lhs_true72
	} else {
		goto lor_lhs_false69
	}

lor_lhs_false69:
	v58 = *libc.As[unsafe.Pointer](parent)
	type70 = libc.Ptr(&libc.As[Tag](v58).F0)
	v59 = *libc.As[int32](type70)
	cmp71 = v59 == 34
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
func scan_start_tag_name(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var tag_name unsafe.Pointer
	var tag, arrayidx unsafe.Pointer
	var tags, tags1, tags2 unsafe.Pointer
	var cmp, v19 bool
	var retval unsafe.Pointer
	var result_symbol, result_symbol5, result_symbol6 unsafe.Pointer
	var v5, v14, inc, v15 int32
	var size, size3, _type unsafe.Pointer
	var v4, v9, idxprom int64
	var v3, v8 unsafe.Pointer
	var v0, v2, v7, v10, v11, v12, v13, v16, v17, v18 unsafe.Pointer
	var scanner_addr, lexer_addr, v1, v6, contents unsafe.Pointer
	var call struct {
		F0 unsafe.Pointer
		F1 int64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, tag_name, tag, v0, call, v1, v2, v3, v4, size, v5, cmp, v6, v7, v8, v9, v10, tags, v11, tags1, contents, v12, v13, tags2, size3, v14, inc, idxprom, arrayidx, _type, v15, v16, result_symbol, v17, result_symbol5, v18, result_symbol6, v19

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
		goto if_end
	}

if_then:
	_array__delete(tag_name)
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v6 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v7 = *libc.As[unsafe.Pointer](v6)
	v8 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v9 = *libc.As[int64](v8)
	tag_for_name(tag, v7, v9)
	v10 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v10).F0)
	_array__grow(tags, 1, int64(24))
	v11 = *libc.As[unsafe.Pointer](scanner_addr)
	tags1 = libc.Ptr(&libc.As[Scanner](v11).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags1).F0)
	v12 = *libc.As[unsafe.Pointer](contents)
	v13 = *libc.As[unsafe.Pointer](scanner_addr)
	tags2 = libc.Ptr(&libc.As[Scanner](v13).F0)
	size3 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags2).F1)
	v14 = *libc.As[int32](size3)
	inc = v14 + 1
	*libc.As[int32](size3) = inc
	idxprom = int64(uint64(uint32(v14)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v12), int(idxprom)*24))
	libc.Memmove(libc.As[byte](arrayidx), libc.As[byte](tag), int64(24))
	_type = libc.Ptr(&libc.As[Tag](tag).F0)
	v15 = *libc.As[int32](_type)
	switch v15 {
	case 99:
		goto sw_bb
	case 106:
		goto sw_bb4
	default:
		goto sw_default
	}

sw_bb:
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v16).F1)
	*libc.As[int16](result_symbol) = 1
	goto sw_epilog

sw_bb4:
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol5 = libc.Ptr(&libc.As[TSLexer](v17).F1)
	*libc.As[int16](result_symbol5) = 2
	goto sw_epilog

sw_default:
	v18 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol6 = libc.Ptr(&libc.As[TSLexer](v18).F1)
	*libc.As[int16](result_symbol6) = 0
	goto sw_epilog

sw_epilog:
	*libc.As[bool](retval) = true
	goto _return

_return:
	v19 = *libc.As[bool](retval)
	return v19
}
func scan_end_tag_name(scanner unsafe.Pointer, lexer unsafe.Pointer) bool {
	var tag_name unsafe.Pointer
	var tag, arrayidx unsafe.Pointer
	var tags, tags3, tags5, tags10, tags11 unsafe.Pointer
	var cmp, cmp2, cmp7, call14, v23 bool
	var retval unsafe.Pointer
	var result_symbol, result_symbol17 unsafe.Pointer
	var v5, v11, v13, sub, v15, v19, sub13 int32
	var size, size1, size4, size6, size12 unsafe.Pointer
	var v4, v9, idxprom int64
	var v3, v8 unsafe.Pointer
	var v0, v2, v7, v10, v12, v14, v16, v17, v18, v20, v21, v22 unsafe.Pointer
	var scanner_addr, lexer_addr, v1, v6, contents unsafe.Pointer
	var call struct {
		F0 unsafe.Pointer
		F1 int64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, scanner_addr, lexer_addr, tag_name, tag, v0, call, v1, v2, v3, v4, size, v5, cmp, v6, v7, v8, v9, v10, tags, size1, v11, cmp2, v12, tags3, size4, v13, sub, v14, tags5, size6, v15, cmp7, v16, tags10, contents, v17, v18, tags11, size12, v19, sub13, idxprom, arrayidx, call14, v20, v21, result_symbol, v22, result_symbol17, v23

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
		goto if_end
	}

if_then:
	_array__delete(tag_name)
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v6 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F0)
	v7 = *libc.As[unsafe.Pointer](v6)
	v8 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](tag_name).F1)
	v9 = *libc.As[int64](v8)
	tag_for_name(tag, v7, v9)
	v10 = *libc.As[unsafe.Pointer](scanner_addr)
	tags = libc.Ptr(&libc.As[Scanner](v10).F0)
	size1 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags).F1)
	v11 = *libc.As[int32](size1)
	cmp2 = uint32(v11) > 0
	if cmp2 {
		goto land_lhs_true
	} else {
		goto if_else16
	}

land_lhs_true:
	v12 = *libc.As[unsafe.Pointer](scanner_addr)
	tags3 = libc.Ptr(&libc.As[Scanner](v12).F0)
	size4 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags3).F1)
	v13 = *libc.As[int32](size4)
	sub = v13 - 1
	v14 = *libc.As[unsafe.Pointer](scanner_addr)
	tags5 = libc.Ptr(&libc.As[Scanner](v14).F0)
	size6 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags5).F1)
	v15 = *libc.As[int32](size6)
	cmp7 = uint32(sub) < uint32(v15)
	if cmp7 {
		goto if_then8
	} else {
		goto if_else
	}

if_then8:
	goto if_end9

if_else:
	libc.AssertFail(libc.As[byte](libc.Ptr(&_str)), libc.As[byte](libc.Ptr(&_str_1)), 265, libc.As[byte](libc.Ptr(&__PRETTY_FUNCTION___scan_end_tag_name)))
	panic("unreachable")

if_end9:
	v16 = *libc.As[unsafe.Pointer](scanner_addr)
	tags10 = libc.Ptr(&libc.As[Scanner](v16).F0)
	contents = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags10).F0)
	v17 = *libc.As[unsafe.Pointer](contents)
	v18 = *libc.As[unsafe.Pointer](scanner_addr)
	tags11 = libc.Ptr(&libc.As[Scanner](v18).F0)
	size12 = libc.Ptr(&libc.As[struct {
		F0 unsafe.Pointer
		F1 int32
		F2 int32
	}](tags11).F1)
	v19 = *libc.As[int32](size12)
	sub13 = v19 - 1
	idxprom = int64(uint64(uint32(sub13)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v17), int(idxprom)*24))
	call14 = tag_eq(arrayidx, tag)
	if call14 {
		goto if_then15
	} else {
		goto if_else16
	}

if_then15:
	v20 = *libc.As[unsafe.Pointer](scanner_addr)
	pop_tag(v20)
	v21 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v21).F1)
	*libc.As[int16](result_symbol) = 3
	goto if_end18

if_else16:
	v22 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol17 = libc.Ptr(&libc.As[TSLexer](v22).F1)
	*libc.As[int16](result_symbol17) = 4
	goto if_end18

if_end18:
	tag_free(tag)
	*libc.As[bool](retval) = true
	goto _return

_return:
	v23 = *libc.As[bool](retval)
	return v23
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
	cmp = uint32(v1) < 23
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
	var tobool, cmp, cmp3, v6 bool
	var v1, call, v3, v5, v8, call5, v10, inc int32
	var lookahead, lookahead1, lookahead2, lookahead4, size unsafe.Pointer
	var idxprom int64
	var conv byte
	var arrayidx unsafe.Pointer
	var v0, v2, v4, v7, v9, v11 unsafe.Pointer
	var lexer_addr, contents unsafe.Pointer
	var v12 struct {
		F0 unsafe.Pointer
		F1 int64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, v0, lookahead, v1, call, tobool, v2, lookahead1, v3, cmp, v4, lookahead2, v5, cmp3, v6, v7, lookahead4, v8, call5, conv, contents, v9, size, v10, inc, idxprom, arrayidx, v11, v12

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
		v6 = true
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
		v6 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead2 = libc.Ptr(&libc.As[TSLexer](v4).F0)
	v5 = *libc.As[int32](lookahead2)
	cmp3 = v5 == 58
	v6 = cmp3
	goto lor_end

lor_end:
	if v6 {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	_array__grow(retval, 1, int64(1))
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead4 = libc.Ptr(&libc.As[TSLexer](v7).F0)
	v8 = *libc.As[int32](lookahead4)
	call5 = libc.Towupper(v8)
	conv = byte(call5)
	contents = libc.Ptr(&libc.As[String](retval).F0)
	v9 = *libc.As[unsafe.Pointer](contents)
	size = libc.Ptr(&libc.As[String](retval).F1)
	v10 = *libc.As[int32](size)
	inc = v10 + 1
	*libc.As[int32](size) = inc
	idxprom = int64(uint64(uint32(v10)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v9), int(idxprom)*1))
	*libc.As[byte](arrayidx) = conv
	v11 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v11)
	goto while_cond

while_end:
	v12 = *libc.As[struct {
		F0 unsafe.Pointer
		F1 int64
	}](retval)
	return v12
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
	var cmp, cmp3, cmp4, v7, cmp6, cmp7, cmp9, cmp11, cmp12, cmp14, v17, cmp17, cmp19, cmp21, cmp23, cmp25, v23, v24 bool
	var retval unsafe.Pointer
	var v1, v3, v4, v5, v6, v8, v9, v10, v11, v12, inc, v13, v14, v15, v16, v18, v19, v20, v21, v22 int32
	var child, i, _type, type1, arrayidx unsafe.Pointer
	var idxprom int64
	var v0, v2 unsafe.Pointer
	var self_addr, other_addr unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, self_addr, other_addr, child, i, v0, _type, v1, v2, type1, v3, v4, cmp, v5, cmp3, v6, cmp4, v7, v8, cmp6, v9, v10, idxprom, arrayidx, v11, cmp7, v12, inc, v13, cmp9, v14, cmp11, v15, cmp12, v16, cmp14, v17, v18, cmp17, v19, cmp19, v20, cmp21, v21, cmp23, v22, cmp25, v23, v24

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
	v2 = *libc.As[unsafe.Pointer](self_addr)
	type1 = libc.Ptr(&libc.As[Tag](v2).F0)
	v3 = *libc.As[int32](type1)
	switch v3 {
	case 73:
		goto sw_bb
	case 50:
		goto sw_bb2
	case 43:
		goto sw_bb2
	case 87:
		goto sw_bb5
	case 40:
		goto sw_bb8
	case 92:
		goto sw_bb10
	case 94:
		goto sw_bb10
	case 93:
		goto sw_bb10
	case 84:
		goto sw_bb16
	case 121:
		goto sw_bb18
	case 113:
		goto sw_bb20
	case 117:
		goto sw_bb20
	default:
		goto sw_default
	}

sw_bb:
	v4 = *libc.As[int32](child)
	cmp = v4 != 73
	*libc.As[bool](retval) = cmp
	goto _return

sw_bb2:
	v5 = *libc.As[int32](child)
	cmp3 = v5 != 50
	if cmp3 {
		goto land_rhs
	} else {
		v7 = false
		goto land_end
	}

land_rhs:
	v6 = *libc.As[int32](child)
	cmp4 = v6 != 43
	v7 = cmp4
	goto land_end

land_end:
	*libc.As[bool](retval) = v7
	goto _return

sw_bb5:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v8 = *libc.As[int32](i)
	cmp6 = v8 < 26
	if cmp6 {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v9 = *libc.As[int32](child)
	v10 = *libc.As[int32](i)
	idxprom = int64(v10)
	arrayidx = libc.Ptr(&TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS[idxprom])
	v11 = *libc.As[int32](arrayidx)
	cmp7 = v9 == v11
	if cmp7 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	goto for_inc

for_inc:
	v12 = *libc.As[int32](i)
	inc = v12 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	*libc.As[bool](retval) = true
	goto _return

sw_bb8:
	v13 = *libc.As[int32](child)
	cmp9 = v13 == 5
	*libc.As[bool](retval) = cmp9
	goto _return

sw_bb10:
	v14 = *libc.As[int32](child)
	cmp11 = v14 != 92
	if cmp11 {
		goto land_lhs_true
	} else {
		v17 = false
		goto land_end15
	}

land_lhs_true:
	v15 = *libc.As[int32](child)
	cmp12 = v15 != 94
	if cmp12 {
		goto land_rhs13
	} else {
		v17 = false
		goto land_end15
	}

land_rhs13:
	v16 = *libc.As[int32](child)
	cmp14 = v16 != 93
	v17 = cmp14
	goto land_end15

land_end15:
	*libc.As[bool](retval) = v17
	goto _return

sw_bb16:
	v18 = *libc.As[int32](child)
	cmp17 = v18 != 84
	*libc.As[bool](retval) = cmp17
	goto _return

sw_bb18:
	v19 = *libc.As[int32](child)
	cmp19 = v19 != 121
	*libc.As[bool](retval) = cmp19
	goto _return

sw_bb20:
	v20 = *libc.As[int32](child)
	cmp21 = v20 != 113
	if cmp21 {
		goto land_lhs_true22
	} else {
		v23 = false
		goto land_end26
	}

land_lhs_true22:
	v21 = *libc.As[int32](child)
	cmp23 = v21 != 117
	if cmp23 {
		goto land_rhs24
	} else {
		v23 = false
		goto land_end26
	}

land_rhs24:
	v22 = *libc.As[int32](child)
	cmp25 = v22 != 121
	v23 = cmp25
	goto land_end26

land_end26:
	*libc.As[bool](retval) = v23
	goto _return

sw_default:
	*libc.As[bool](retval) = true
	goto _return

_return:
	v24 = *libc.As[bool](retval)
	return v24
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
