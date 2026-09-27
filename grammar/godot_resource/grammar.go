package grammar_godot_resource

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type TSSymbolMetadata struct {
	F0 byte
	F1 byte
	F2 byte
}
type TSLexMode struct {
	F0 int16
	F1 int16
}
type TSParseActionEntry struct {
	F0 TSParseAction
}
type TSParseAction struct {
	F0 anon_1
}
type anon_1 struct {
	F0 byte
	F1 byte
	F2 int16
	F3 int16
	F4 int16
}
type anon_2 struct {
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

var tree_sitter_godot_resource_language struct {
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
	F28 struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}
	F29 unsafe.Pointer
}
var ts_small_parse_table [1442]int16 = [1442]int16{11, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 23, 1, 18, 15, 1, 21, 46, 1, 29, 15, 2, 19, 6, 13, 4, 3, 4, 5, 7, 44, 5, 22, 23, 30, 31, 34, 11, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 29, 1, 16, 15, 1, 21, 58, 1, 29, 27, 2, 19, 6, 25, 4, 3, 4, 5, 7, 85, 5, 22, 23, 30, 31, 34, 11, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 35, 1, 18, 15, 1, 21, 65, 1, 29, 33, 2, 19, 6, 31, 4, 3, 4, 5, 7, 45, 5, 22, 23, 30, 31, 34, 11, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 37, 1, 16, 15, 1, 21, 61, 1, 29, 27, 2, 19, 6, 25, 4, 3, 4, 5, 7, 85, 5, 22, 23, 30, 31, 34, 10, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 43, 1, 10, 15, 1, 21, 41, 2, 19, 6, 39, 4, 3, 4, 5, 7, 53, 5, 22, 23, 30, 31, 34, 10, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 15, 1, 21, 79, 1, 29, 27, 2, 19, 6, 25, 4, 3, 4, 5, 7, 85, 5, 22, 23, 30, 31, 34, 10, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 15, 1, 21, 76, 1, 29, 47, 2, 19, 6, 45, 4, 3, 4, 5, 7, 55, 5, 22, 23, 30, 31, 34, 10, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 53, 1, 10, 15, 1, 21, 51, 2, 19, 6, 49, 4, 3, 4, 5, 7, 60, 5, 22, 23, 30, 31, 34, 9, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 15, 1, 21, 57, 2, 19, 6, 55, 4, 3, 4, 5, 7, 49, 5, 22, 23, 30, 31, 34, 9, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 15, 1, 21, 61, 2, 19, 6, 59, 4, 3, 4, 5, 7, 62, 5, 22, 23, 30, 31, 34, 9, 3, 1, 2, 17, 1, 8, 19, 1, 9, 21, 1, 14, 63, 1, 1, 15, 1, 21, 67, 2, 19, 6, 65, 4, 3, 4, 5, 7, 50, 5, 22, 23, 30, 31, 34, 9, 3, 1, 2, 11, 1, 1, 17, 1, 8, 19, 1, 9, 21, 1, 14, 15, 1, 21, 71, 2, 19, 6, 69, 4, 3, 4, 5, 7, 67, 5, 22, 23, 30, 31, 34, 9, 3, 1, 2, 11, 1, 1, 77, 1, 8, 79, 1, 9, 81, 1, 14, 30, 1, 21, 75, 2, 19, 6, 73, 4, 3, 4, 5, 7, 78, 5, 22, 23, 30, 31, 34, 6, 3, 1, 2, 85, 1, 9, 87, 1, 17, 22, 1, 32, 75, 1, 33, 83, 7, 0, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 89, 9, 1, 9, 10, 11, 13, 15, 16, 17, 18, 2, 3, 1, 2, 91, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 93, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 95, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 97, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 99, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 101, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 103, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 105, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 107, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 109, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 111, 8, 0, 9, 10, 12, 13, 15, 16, 18, 2, 3, 1, 2, 113, 8, 0, 9, 10, 12, 13, 15, 16, 18, 6, 3, 1, 2, 7, 1, 9, 9, 1, 12, 115, 1, 0, 36, 2, 28, 35, 43, 2, 24, 36, 6, 3, 1, 2, 85, 1, 9, 117, 1, 17, 80, 1, 33, 84, 1, 32, 83, 2, 1, 10, 4, 3, 1, 2, 9, 1, 12, 119, 2, 0, 9, 39, 2, 27, 28, 4, 3, 1, 2, 9, 1, 12, 121, 2, 0, 9, 35, 2, 27, 28, 5, 3, 1, 2, 123, 1, 1, 125, 1, 10, 91, 1, 21, 37, 2, 25, 26, 5, 3, 1, 2, 123, 1, 1, 127, 1, 10, 91, 1, 21, 33, 2, 25, 26, 4, 3, 1, 2, 131, 1, 12, 129, 2, 0, 9, 35, 2, 27, 28, 4, 3, 1, 2, 136, 1, 12, 134, 2, 0, 9, 36, 2, 28, 35, 5, 3, 1, 2, 139, 1, 1, 142, 1, 10, 91, 1, 21, 37, 2, 25, 26, 4, 3, 1, 2, 9, 1, 12, 144, 2, 0, 9, 32, 2, 27, 28, 4, 3, 1, 2, 9, 1, 12, 144, 2, 0, 9, 35, 2, 27, 28, 4, 3, 1, 2, 146, 1, 0, 148, 1, 9, 40, 2, 24, 36, 4, 3, 1, 2, 7, 1, 9, 115, 1, 0, 40, 2, 24, 36, 2, 3, 1, 2, 89, 4, 0, 9, 12, 17, 4, 3, 1, 2, 7, 1, 9, 151, 1, 0, 40, 2, 24, 36, 5, 3, 1, 2, 153, 1, 13, 155, 1, 15, 157, 1, 18, 54, 1, 39, 5, 3, 1, 2, 153, 1, 13, 155, 1, 15, 159, 1, 18, 47, 1, 39, 4, 3, 1, 2, 155, 1, 15, 157, 1, 18, 54, 1, 39, 4, 3, 1, 2, 155, 1, 15, 161, 1, 18, 57, 1, 39, 4, 3, 1, 2, 163, 1, 10, 165, 1, 15, 48, 1, 38, 2, 3, 1, 2, 168, 3, 15, 16, 18, 2, 3, 1, 2, 170, 3, 0, 9, 12, 4, 3, 1, 2, 172, 1, 15, 175, 1, 16, 51, 1, 37, 4, 3, 1, 2, 177, 1, 10, 179, 1, 15, 48, 1, 38, 4, 3, 1, 2, 179, 1, 15, 181, 1, 10, 56, 1, 38, 4, 3, 1, 2, 155, 1, 15, 183, 1, 18, 57, 1, 39, 3, 3, 1, 2, 153, 1, 13, 185, 2, 15, 18, 4, 3, 1, 2, 179, 1, 15, 187, 1, 10, 48, 1, 38, 4, 3, 1, 2, 185, 1, 18, 189, 1, 15, 57, 1, 39, 4, 3, 1, 2, 192, 1, 15, 194, 1, 16, 59, 1, 37, 4, 3, 1, 2, 192, 1, 15, 196, 1, 16, 51, 1, 37, 4, 3, 1, 2, 179, 1, 15, 198, 1, 10, 63, 1, 38, 4, 3, 1, 2, 192, 1, 15, 200, 1, 16, 64, 1, 37, 4, 3, 1, 2, 179, 1, 15, 202, 1, 10, 52, 1, 38, 4, 3, 1, 2, 179, 1, 15, 204, 1, 10, 48, 1, 38, 4, 3, 1, 2, 192, 1, 15, 206, 1, 16, 51, 1, 37, 4, 3, 1, 2, 155, 1, 15, 159, 1, 18, 47, 1, 39, 2, 3, 1, 2, 111, 2, 1, 10, 2, 3, 1, 2, 163, 2, 10, 15, 2, 3, 1, 2, 99, 2, 1, 10, 2, 3, 1, 2, 93, 2, 1, 10, 2, 3, 1, 2, 103, 2, 1, 10, 2, 3, 1, 2, 109, 2, 1, 10, 2, 3, 1, 2, 91, 2, 1, 10, 2, 3, 1, 2, 113, 2, 1, 10, 2, 3, 1, 2, 107, 2, 1, 10, 3, 3, 1, 2, 87, 1, 17, 26, 1, 32, 2, 3, 1, 2, 185, 2, 15, 18, 3, 3, 1, 2, 123, 1, 1, 34, 1, 21, 2, 3, 1, 2, 208, 2, 1, 10, 2, 3, 1, 2, 175, 2, 15, 16, 3, 3, 1, 2, 117, 1, 17, 71, 1, 32, 2, 3, 1, 2, 105, 2, 1, 10, 2, 3, 1, 2, 95, 2, 1, 10, 2, 3, 1, 2, 97, 2, 1, 10, 2, 3, 1, 2, 101, 2, 1, 10, 2, 3, 1, 2, 153, 1, 13, 2, 3, 1, 2, 210, 1, 17, 2, 3, 1, 2, 212, 1, 19, 2, 3, 1, 2, 214, 1, 19, 2, 3, 1, 2, 216, 1, 17, 2, 3, 1, 2, 218, 1, 11, 2, 3, 1, 2, 220, 1, 11, 2, 3, 1, 2, 222, 1, 0}
var ts_small_parse_table_map [91]int32 = [91]int32{0, 42, 84, 126, 168, 207, 246, 285, 324, 360, 396, 432, 468, 504, 529, 544, 558, 572, 586, 600, 614, 628, 642, 656, 670, 684, 698, 712, 733, 753, 768, 783, 800, 817, 832, 847, 864, 879, 894, 908, 922, 932, 946, 962, 978, 991, 1004, 1017, 1026, 1035, 1048, 1061, 1074, 1087, 1098, 1111, 1124, 1137, 1150, 1163, 1176, 1189, 1202, 1215, 1228, 1236, 1244, 1252, 1260, 1268, 1276, 1284, 1292, 1300, 1310, 1318, 1328, 1336, 1344, 1354, 1362, 1370, 1378, 1386, 1393, 1400, 1407, 1414, 1421, 1428, 1435}
var ts_symbol_names [41]unsafe.Pointer = [41]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_30), libc.Ptr(&_str_31), libc.Ptr(&_str_32), libc.Ptr(&_str_33), libc.Ptr(&_str_34), libc.Ptr(&_str_35), libc.Ptr(&_str_36), libc.Ptr(&_str_37), libc.Ptr(&_str_38), libc.Ptr(&_str_39), libc.Ptr(&_str_40), libc.Ptr(&_str_41), libc.Ptr(&_str_21)}
var ts_symbol_metadata [41]TSSymbolMetadata = [41]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}}
var ts_symbol_map [41]int16 = [41]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [2][5]int16 = [2][5]int16{[5]int16{}, [5]int16{0, 40, 0, 0, 0}}
var ts_lex_modes [93]TSLexMode = [93]TSLexMode{TSLexMode{0, 1}, TSLexMode{17, 0}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{17, 0}, TSLexMode{}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{}, TSLexMode{17, 0}, TSLexMode{17, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{17, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{17, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{0, 1}, TSLexMode{0, 1}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}}
var ts_external_scanner_states [2][1]byte = [2][1]byte{[1]byte{}, [1]byte{1}}
var ts_external_scanner_symbol_map [1]int16 = [1]int16{19}
var ts_primary_state_ids [93]int16 = [93]int16{0, 1, 2, 3, 2, 3, 6, 7, 8, 6, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 15, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 16, 43, 44, 44, 46, 47, 48, 49, 50, 51, 52, 53, 47, 55, 56, 57, 58, 59, 53, 58, 62, 56, 59, 46, 27, 67, 21, 18, 23, 26, 17, 28, 25, 75, 76, 77, 78, 79, 75, 24, 19, 20, 22, 85, 86, 87, 87, 89, 90, 91, 92}
var ts_parse_table struct {
	F0 struct {
		F0 [20]int16
		F1 [20]int16
	}
	F1 [40]int16
} = struct {
	F0 struct {
		F0 [20]int16
		F1 [20]int16
	}
	F1 [40]int16
}{struct {
	F0 [20]int16
	F1 [20]int16
}{[20]int16{1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1}, [20]int16{}}, [40]int16{5, 0, 3, 0, 0, 0, 0, 0, 0, 7, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 92, 0, 0, 0, 41, 0, 0, 0, 29, 0, 0, 0, 0, 0, 0, 29, 41, 0, 0, 0}}
var ts_parse_actions struct {
	F0 struct {
		F0 anon_2
		F1 [6]byte
	}
	F1 struct {
		F0 anon_2
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F6 TSParseActionEntry
	F7 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F52 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F53 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F56 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F57 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F62 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F63 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F68 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F69 struct {
		F0 anon_2
		F1 [6]byte
	}
	F70 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F71 struct {
		F0 anon_2
		F1 [6]byte
	}
	F72 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F73 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F76 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F77 struct {
		F0 anon_2
		F1 [6]byte
	}
	F78 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F79 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F82 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F83 struct {
		F0 anon_2
		F1 [6]byte
	}
	F84 TSParseActionEntry
	F85 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F88 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F89 struct {
		F0 anon_2
		F1 [6]byte
	}
	F90 TSParseActionEntry
	F91 struct {
		F0 anon_2
		F1 [6]byte
	}
	F92 TSParseActionEntry
	F93 struct {
		F0 anon_2
		F1 [6]byte
	}
	F94 TSParseActionEntry
	F95 struct {
		F0 anon_2
		F1 [6]byte
	}
	F96 TSParseActionEntry
	F97 struct {
		F0 anon_2
		F1 [6]byte
	}
	F98 TSParseActionEntry
	F99 struct {
		F0 anon_2
		F1 [6]byte
	}
	F100 TSParseActionEntry
	F101 struct {
		F0 anon_2
		F1 [6]byte
	}
	F102 TSParseActionEntry
	F103 struct {
		F0 anon_2
		F1 [6]byte
	}
	F104 TSParseActionEntry
	F105 struct {
		F0 anon_2
		F1 [6]byte
	}
	F106 TSParseActionEntry
	F107 struct {
		F0 anon_2
		F1 [6]byte
	}
	F108 TSParseActionEntry
	F109 struct {
		F0 anon_2
		F1 [6]byte
	}
	F110 TSParseActionEntry
	F111 struct {
		F0 anon_2
		F1 [6]byte
	}
	F112 TSParseActionEntry
	F113 struct {
		F0 anon_2
		F1 [6]byte
	}
	F114 TSParseActionEntry
	F115 struct {
		F0 anon_2
		F1 [6]byte
	}
	F116 TSParseActionEntry
	F117 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F120 TSParseActionEntry
	F121 struct {
		F0 anon_2
		F1 [6]byte
	}
	F122 TSParseActionEntry
	F123 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F130 TSParseActionEntry
	F131 struct {
		F0 anon_2
		F1 [6]byte
	}
	F132 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F135 TSParseActionEntry
	F136 struct {
		F0 anon_2
		F1 [6]byte
	}
	F137 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F140 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F143 TSParseActionEntry
	F144 struct {
		F0 anon_2
		F1 [6]byte
	}
	F145 TSParseActionEntry
	F146 struct {
		F0 anon_2
		F1 [6]byte
	}
	F147 TSParseActionEntry
	F148 struct {
		F0 anon_2
		F1 [6]byte
	}
	F149 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F152 TSParseActionEntry
	F153 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F164 TSParseActionEntry
	F165 struct {
		F0 anon_2
		F1 [6]byte
	}
	F166 TSParseActionEntry
	F167 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F168 struct {
		F0 anon_2
		F1 [6]byte
	}
	F169 TSParseActionEntry
	F170 struct {
		F0 anon_2
		F1 [6]byte
	}
	F171 TSParseActionEntry
	F172 struct {
		F0 anon_2
		F1 [6]byte
	}
	F173 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F176 TSParseActionEntry
	F177 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F186 TSParseActionEntry
	F187 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F190 TSParseActionEntry
	F191 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F192 struct {
		F0 anon_2
		F1 [6]byte
	}
	F193 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F194 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F197 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F198 struct {
		F0 anon_2
		F1 [6]byte
	}
	F199 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F200 struct {
		F0 anon_2
		F1 [6]byte
	}
	F201 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F202 struct {
		F0 anon_2
		F1 [6]byte
	}
	F203 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F204 struct {
		F0 anon_2
		F1 [6]byte
	}
	F205 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F206 struct {
		F0 anon_2
		F1 [6]byte
	}
	F207 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F208 struct {
		F0 anon_2
		F1 [6]byte
	}
	F209 TSParseActionEntry
	F210 struct {
		F0 anon_2
		F1 [6]byte
	}
	F211 TSParseActionEntry
	F212 struct {
		F0 anon_2
		F1 [6]byte
	}
	F213 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F214 struct {
		F0 anon_2
		F1 [6]byte
	}
	F215 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F216 struct {
		F0 anon_2
		F1 [6]byte
	}
	F217 TSParseActionEntry
	F218 struct {
		F0 anon_2
		F1 [6]byte
	}
	F219 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F220 struct {
		F0 anon_2
		F1 [6]byte
	}
	F221 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F222 struct {
		F0 anon_2
		F1 [6]byte
	}
	F223 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
} = struct {
	F0 struct {
		F0 anon_2
		F1 [6]byte
	}
	F1 struct {
		F0 anon_2
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F6 TSParseActionEntry
	F7 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F52 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F53 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F56 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F57 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F62 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F63 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F68 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F69 struct {
		F0 anon_2
		F1 [6]byte
	}
	F70 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F71 struct {
		F0 anon_2
		F1 [6]byte
	}
	F72 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F73 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F76 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F77 struct {
		F0 anon_2
		F1 [6]byte
	}
	F78 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F79 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F82 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F83 struct {
		F0 anon_2
		F1 [6]byte
	}
	F84 TSParseActionEntry
	F85 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F88 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F89 struct {
		F0 anon_2
		F1 [6]byte
	}
	F90 TSParseActionEntry
	F91 struct {
		F0 anon_2
		F1 [6]byte
	}
	F92 TSParseActionEntry
	F93 struct {
		F0 anon_2
		F1 [6]byte
	}
	F94 TSParseActionEntry
	F95 struct {
		F0 anon_2
		F1 [6]byte
	}
	F96 TSParseActionEntry
	F97 struct {
		F0 anon_2
		F1 [6]byte
	}
	F98 TSParseActionEntry
	F99 struct {
		F0 anon_2
		F1 [6]byte
	}
	F100 TSParseActionEntry
	F101 struct {
		F0 anon_2
		F1 [6]byte
	}
	F102 TSParseActionEntry
	F103 struct {
		F0 anon_2
		F1 [6]byte
	}
	F104 TSParseActionEntry
	F105 struct {
		F0 anon_2
		F1 [6]byte
	}
	F106 TSParseActionEntry
	F107 struct {
		F0 anon_2
		F1 [6]byte
	}
	F108 TSParseActionEntry
	F109 struct {
		F0 anon_2
		F1 [6]byte
	}
	F110 TSParseActionEntry
	F111 struct {
		F0 anon_2
		F1 [6]byte
	}
	F112 TSParseActionEntry
	F113 struct {
		F0 anon_2
		F1 [6]byte
	}
	F114 TSParseActionEntry
	F115 struct {
		F0 anon_2
		F1 [6]byte
	}
	F116 TSParseActionEntry
	F117 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F120 TSParseActionEntry
	F121 struct {
		F0 anon_2
		F1 [6]byte
	}
	F122 TSParseActionEntry
	F123 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F130 TSParseActionEntry
	F131 struct {
		F0 anon_2
		F1 [6]byte
	}
	F132 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F135 TSParseActionEntry
	F136 struct {
		F0 anon_2
		F1 [6]byte
	}
	F137 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F140 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F143 TSParseActionEntry
	F144 struct {
		F0 anon_2
		F1 [6]byte
	}
	F145 TSParseActionEntry
	F146 struct {
		F0 anon_2
		F1 [6]byte
	}
	F147 TSParseActionEntry
	F148 struct {
		F0 anon_2
		F1 [6]byte
	}
	F149 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F152 TSParseActionEntry
	F153 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F164 TSParseActionEntry
	F165 struct {
		F0 anon_2
		F1 [6]byte
	}
	F166 TSParseActionEntry
	F167 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F168 struct {
		F0 anon_2
		F1 [6]byte
	}
	F169 TSParseActionEntry
	F170 struct {
		F0 anon_2
		F1 [6]byte
	}
	F171 TSParseActionEntry
	F172 struct {
		F0 anon_2
		F1 [6]byte
	}
	F173 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F176 TSParseActionEntry
	F177 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F186 TSParseActionEntry
	F187 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F190 TSParseActionEntry
	F191 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F192 struct {
		F0 anon_2
		F1 [6]byte
	}
	F193 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F194 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F197 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F198 struct {
		F0 anon_2
		F1 [6]byte
	}
	F199 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F200 struct {
		F0 anon_2
		F1 [6]byte
	}
	F201 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F202 struct {
		F0 anon_2
		F1 [6]byte
	}
	F203 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F204 struct {
		F0 anon_2
		F1 [6]byte
	}
	F205 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F206 struct {
		F0 anon_2
		F1 [6]byte
	}
	F207 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F208 struct {
		F0 anon_2
		F1 [6]byte
	}
	F209 TSParseActionEntry
	F210 struct {
		F0 anon_2
		F1 [6]byte
	}
	F211 TSParseActionEntry
	F212 struct {
		F0 anon_2
		F1 [6]byte
	}
	F213 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F214 struct {
		F0 anon_2
		F1 [6]byte
	}
	F215 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F216 struct {
		F0 anon_2
		F1 [6]byte
	}
	F217 TSParseActionEntry
	F218 struct {
		F0 anon_2
		F1 [6]byte
	}
	F219 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F220 struct {
		F0 anon_2
		F1 [6]byte
	}
	F221 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F222 struct {
		F0 anon_2
		F1 [6]byte
	}
	F223 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
}{struct {
	F0 anon_2
	F1 [6]byte
}{}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 byte
		F1 [7]byte
	}
}{struct {
	F0 byte
	F1 [7]byte
}{3, [7]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
}{0, 0, 83, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
}{0, 0, 5, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 23, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 21, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 31, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 30, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 31, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 30, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 31, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 32, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 1}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 32, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 32, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 30, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 24, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 24, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
}{0, 0, 38, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 27, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 27, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 90, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 35, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 35, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 90, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 25, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 16, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 25, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 24, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 36, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 36, 0, 0}}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 38, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 38, 0, 0}}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 29, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 28, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 37, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 7, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 37, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 39, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 39, 0, 0}}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 26, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 33, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 33, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
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
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 byte
		F1 [7]byte
	}
}{struct {
	F0 byte
	F1 [7]byte
}{2, [7]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [12]byte = [12]byte{95, 105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}
var _str_4 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_5 [5]byte = [5]byte{116, 114, 117, 101, 0}
var _str_6 [6]byte = [6]byte{102, 97, 108, 115, 101, 0}
var _str_7 [5]byte = [5]byte{110, 117, 108, 108, 0}
var _str_8 [6]byte = [6]byte{102, 108, 111, 97, 116, 0}
var _str_9 [8]byte = [8]byte{105, 110, 116, 101, 103, 101, 114, 0}
var _str_10 [2]byte = [2]byte{38, 0}
var _str_11 [2]byte = [2]byte{91, 0}
var _str_12 [2]byte = [2]byte{93, 0}
var _str_13 [2]byte = [2]byte{61, 0}
var _str_14 [5]byte = [5]byte{112, 97, 116, 104, 0}
var _str_15 [2]byte = [2]byte{58, 0}
var _str_16 [2]byte = [2]byte{123, 0}
var _str_17 [2]byte = [2]byte{44, 0}
var _str_18 [2]byte = [2]byte{125, 0}
var _str_19 [2]byte = [2]byte{40, 0}
var _str_20 [2]byte = [2]byte{41, 0}
var _str_21 [7]byte = [7]byte{115, 116, 114, 105, 110, 103, 0}
var _str_22 [9]byte = [9]byte{114, 101, 115, 111, 117, 114, 99, 101, 0}
var _str_23 [11]byte = [11]byte{105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}
var _str_24 [12]byte = [12]byte{115, 116, 114, 105, 110, 103, 95, 110, 97, 109, 101, 0}
var _str_25 [7]byte = [7]byte{95, 118, 97, 108, 117, 101, 0}
var _str_26 [8]byte = [8]byte{115, 101, 99, 116, 105, 111, 110, 0}
var _str_27 [12]byte = [12]byte{95, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 0}
var _str_28 [10]byte = [10]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 0}
var _str_29 [12]byte = [12]byte{95, 112, 114, 111, 112, 101, 114, 116, 105, 101, 115, 0}
var _str_30 [9]byte = [9]byte{112, 114, 111, 112, 101, 114, 116, 121, 0}
var _str_31 [5]byte = [5]byte{112, 97, 105, 114, 0}
var _str_32 [11]byte = [11]byte{100, 105, 99, 116, 105, 111, 110, 97, 114, 121, 0}
var _str_33 [6]byte = [6]byte{97, 114, 114, 97, 121, 0}
var _str_34 [10]byte = [10]byte{97, 114, 103, 117, 109, 101, 110, 116, 115, 0}
var _str_35 [11]byte = [11]byte{95, 116, 121, 112, 101, 95, 97, 114, 103, 115, 0}
var _str_36 [12]byte = [12]byte{99, 111, 110, 115, 116, 114, 117, 99, 116, 111, 114, 0}
var _str_37 [17]byte = [17]byte{114, 101, 115, 111, 117, 114, 99, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_38 [17]byte = [17]byte{114, 101, 115, 111, 117, 114, 99, 101, 95, 114, 101, 112, 101, 97, 116, 50, 0}
var _str_39 [19]byte = [19]byte{100, 105, 99, 116, 105, 111, 110, 97, 114, 121, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_40 [14]byte = [14]byte{97, 114, 114, 97, 121, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_41 [18]byte = [18]byte{97, 114, 103, 117, 109, 101, 110, 116, 115, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var ts_lex_map [22]int16 = [22]int16{38, 31, 40, 40, 41, 41, 44, 38, 45, 9, 46, 2, 48, 25, 58, 36, 59, 20, 61, 34, 91, 32}
var ts_lex_map_42 [22]int16 = [22]int16{45, 9, 46, 22, 95, 27, 66, 4, 98, 4, 69, 1, 101, 1, 79, 5, 111, 5, 88, 6, 120, 6}

func init() {
	tree_sitter_godot_resource_language = struct {
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
		F28 struct {
			F0 unsafe.Pointer
			F1 unsafe.Pointer
			F2 unsafe.Pointer
			F3 unsafe.Pointer
			F4 unsafe.Pointer
			F5 unsafe.Pointer
			F6 unsafe.Pointer
		}
		F29 unsafe.Pointer
	}{14, 40, 1, 20, 1, 93, 2, 2, 0, 5, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), libc.FuncCode(ts_lex_keywords), 1, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{libc.Ptr(&ts_external_scanner_states), libc.Ptr(&ts_external_scanner_symbol_map), libc.FuncCode(tree_sitter_godot_resource_external_scanner_create), libc.FuncCode(tree_sitter_godot_resource_external_scanner_destroy), libc.FuncCode(tree_sitter_godot_resource_external_scanner_scan), libc.FuncCode(tree_sitter_godot_resource_external_scanner_serialize), libc.FuncCode(tree_sitter_godot_resource_external_scanner_deserialize)}, libc.Ptr(&ts_primary_state_ids)}
}
func tree_sitter_godot_resource_external_scanner_scan(payload unsafe.Pointer, lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var loadedv, tobool, cmp, tobool7, cmp9, cmp11, v26 bool
	var retval unsafe.Pointer
	var result_symbol unsafe.Pointer
	var v3, call, v8, v13, v14, v16, v22 int32
	var last_char, lookahead, lookahead1, lookahead6, lookahead10, lookahead15 unsafe.Pointer
	var v1 byte
	var arrayidx unsafe.Pointer
	var v0, v2, v4, v5, v6, v7, v9, v10, v11, v12, v15, v17, v18, v19, v20, v21, v23, v24, v25 unsafe.Pointer
	var payload_addr, lexer_addr, valid_symbols_addr, advance, advance4, advance13, advance16 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, payload_addr, lexer_addr, valid_symbols_addr, last_char, v0, arrayidx, v1, loadedv, v2, lookahead, v3, call, tobool, v4, advance, v5, v6, v7, lookahead1, v8, cmp, v9, advance4, v10, v11, v12, lookahead6, v13, tobool7, v14, cmp9, v15, lookahead10, v16, cmp11, v17, advance13, v18, v19, v20, result_symbol, v21, lookahead15, v22, v23, advance16, v24, v25, v26

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
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
	last_char = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx = v0
	v1 = *libc.As[byte](arrayidx)
	loadedv = (v1 & 1) != 0
	if loadedv {
		goto if_end
	} else {
		goto if_then
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	goto while_cond

while_cond:
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v2).F0)
	v3 = *libc.As[int32](lookahead)
	call = libc.Iswspace(v3)
	tobool = call != 0
	if tobool {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	advance = libc.Ptr(&libc.As[TSLexer](v4).F2)
	v5 = *libc.As[unsafe.Pointer](advance)
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v5)(v6, true)
	goto while_cond

while_end:
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v7).F0)
	v8 = *libc.As[int32](lookahead1)
	cmp = v8 != 34
	if cmp {
		goto if_then2
	} else {
		goto if_end3
	}

if_then2:
	*libc.As[bool](retval) = false
	goto _return

if_end3:
	*libc.As[int32](last_char) = 34
	v9 = *libc.As[unsafe.Pointer](lexer_addr)
	advance4 = libc.Ptr(&libc.As[TSLexer](v9).F2)
	v10 = *libc.As[unsafe.Pointer](advance4)
	v11 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v10)(v11, false)
	goto while_cond5

while_cond5:
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead6 = libc.Ptr(&libc.As[TSLexer](v12).F0)
	v13 = *libc.As[int32](lookahead6)
	tobool7 = v13 != 0
	if tobool7 {
		goto while_body8
	} else {
		goto while_end17
	}

while_body8:
	v14 = *libc.As[int32](last_char)
	cmp9 = v14 != 92
	if cmp9 {
		goto land_lhs_true
	} else {
		goto if_end14
	}

land_lhs_true:
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead10 = libc.Ptr(&libc.As[TSLexer](v15).F0)
	v16 = *libc.As[int32](lookahead10)
	cmp11 = v16 == 34
	if cmp11 {
		goto if_then12
	} else {
		goto if_end14
	}

if_then12:
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	advance13 = libc.Ptr(&libc.As[TSLexer](v17).F2)
	v18 = *libc.As[unsafe.Pointer](advance13)
	v19 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v18)(v19, false)
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v20).F1)
	*libc.As[int16](result_symbol) = 0
	*libc.As[bool](retval) = true
	goto _return

if_end14:
	v21 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead15 = libc.Ptr(&libc.As[TSLexer](v21).F0)
	v22 = *libc.As[int32](lookahead15)
	*libc.As[int32](last_char) = v22
	v23 = *libc.As[unsafe.Pointer](lexer_addr)
	advance16 = libc.Ptr(&libc.As[TSLexer](v23).F2)
	v24 = *libc.As[unsafe.Pointer](advance16)
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v24)(v25, false)
	goto while_cond5

while_end17:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v26 = *libc.As[bool](retval)
	return v26
}
func tree_sitter_godot_resource_external_scanner_create() unsafe.Pointer {
	return nil
}
func tree_sitter_godot_resource_external_scanner_serialize(payload unsafe.Pointer, buffer unsafe.Pointer) int32 {
	var payload_addr, buffer_addr unsafe.Pointer
	_, _ = payload_addr, buffer_addr

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
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](buffer_addr) = buffer
	return 0
}
func tree_sitter_godot_resource_external_scanner_deserialize(payload unsafe.Pointer, buffer unsafe.Pointer, length int32) {
	var length_addr unsafe.Pointer
	var payload_addr, buffer_addr unsafe.Pointer
	_, _, _ = payload_addr, buffer_addr, length_addr

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
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](buffer_addr) = buffer
	*libc.As[int32](length_addr) = length
}
func tree_sitter_godot_resource_external_scanner_destroy(payload unsafe.Pointer) {
	var payload_addr unsafe.Pointer
	_ = payload_addr

	payload_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
}
func tree_sitter_godot_resource() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_godot_resource_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp18, cmp22, cmp26, cmp30, cmp32, cmp34, cmp37, cmp40, cmp43, cmp47, cmp50, cmp54, cmp57, cmp60, cmp63, cmp66, loadedv70, cmp72, cmp76, cmp80, cmp83, loadedv87, cmp89, cmp93, cmp96, loadedv100, cmp102, cmp106, cmp109, loadedv113, cmp115, cmp119, cmp122, loadedv126, cmp128, cmp132, cmp135, loadedv139, cmp141, cmp145, cmp148, cmp151, cmp154, cmp157, cmp160, loadedv164, cmp166, cmp169, loadedv173, cmp175, cmp178, loadedv182, cmp184, cmp187, loadedv191, cmp193, cmp196, loadedv200, cmp202, cmp205, loadedv209, cmp211, cmp214, cmp217, cmp220, cmp223, cmp226, loadedv230, loadedv232, cmp235, loadedv239, loadedv241, cmp244, cmp248, loadedv252, loadedv254, cmp257, loadedv261, loadedv263, cmp266, cmp270, loadedv274, loadedv276, cmp279, cmp283, cmp287, cmp291, cmp295, cmp299, cmp303, cmp307, cmp311, cmp315, cmp318, cmp321, cmp324, cmp327, cmp330, cmp334, cmp337, cmp340, cmp343, cmp346, cmp349, cmp352, loadedv356, loadedv358, cmp362, cmp365, cmp368, cmp371, cmp374, cmp377, cmp380, loadedv384, cmp388, cmp391, loadedv395, cmp399, cmp403, cmp407, cmp410, cmp414, cmp417, loadedv421, cmp425, cmp429, cmp432, cmp436, cmp439, loadedv443, cmp447, cmp451, cmp455, cmp458, loadedv462, cmp466, cmp470, cmp473, loadedv477, cmp484, cmp490, cmp500, cmp503, loadedv507, cmp511, cmp515, cmp519, cmp523, cmp526, cmp530, cmp533, loadedv537, cmp541, cmp545, cmp549, cmp552, cmp556, cmp559, loadedv563, cmp567, cmp571, cmp574, loadedv578, cmp582, cmp586, cmp589, loadedv593, cmp597, cmp601, cmp604, cmp607, cmp610, cmp613, cmp616, loadedv620, loadedv624, loadedv628, loadedv632, loadedv636, cmp640, cmp643, cmp646, cmp649, cmp652, cmp655, cmp658, loadedv662, loadedv666, loadedv670, loadedv674, loadedv678, loadedv682, loadedv686, v314 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v193, v196 int16
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol360, result_symbol386, result_symbol397, result_symbol423, result_symbol445, result_symbol464, result_symbol479, arrayidx488, arrayidx495, result_symbol509, result_symbol539, result_symbol565, result_symbol580, result_symbol595, result_symbol622, result_symbol626, result_symbol630, result_symbol634, result_symbol638, result_symbol664, result_symbol668, result_symbol672, result_symbol676, result_symbol680, result_symbol684 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v50, v51, v53, v54, v55, v57, v58, v59, v60, v61, v62, v63, v65, v66, v68, v69, v71, v72, v74, v75, v77, v78, v80, v81, v82, v83, v84, v85, v88, v91, v92, v95, v98, v99, v102, v103, v104, v105, v106, v107, v108, v109, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v120, v121, v122, v123, v134, v135, v136, v137, v138, v139, v140, v146, v147, v153, v154, v155, v156, v157, v158, v164, v165, v166, v167, v168, v174, v175, v176, v177, v183, v184, v185, v191, v192, conv489, v194, v195, add493, v197, add498, v198, v199, v205, v206, v207, v208, v209, v210, v211, v217, v218, v219, v220, v221, v222, v228, v229, v230, v236, v237, v238, v244, v245, v246, v247, v248, v249, v250, v276, v277, v278, v279, v280, v281, v282 int32
	var lookahead, i, i481, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv483, idxprom487, idxprom494 int64
	var v3, storedv, v10, v35, v40, v44, v48, v52, v56, v64, v67, v70, v73, v76, v79, v86, v87, v89, v90, v93, v94, v96, v97, v100, v101, v124, v129, v141, v148, v159, v169, v178, v186, v200, v212, v223, v231, v239, v251, v256, v261, v266, v271, v283, v288, v293, v298, v303, v308, v313 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v125, v126, v127, v128, v130, v131, v132, v133, v142, v143, v144, v145, v149, v150, v151, v152, v160, v161, v162, v163, v170, v171, v172, v173, v179, v180, v181, v182, v187, v188, v189, v190, v201, v202, v203, v204, v213, v214, v215, v216, v224, v225, v226, v227, v232, v233, v234, v235, v240, v241, v242, v243, v252, v253, v254, v255, v257, v258, v259, v260, v262, v263, v264, v265, v267, v268, v269, v270, v272, v273, v274, v275, v284, v285, v286, v287, v289, v290, v291, v292, v294, v295, v296, v297, v299, v300, v301, v302, v304, v305, v306, v307, v309, v310, v311, v312 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end361, mark_end387, mark_end398, mark_end424, mark_end446, mark_end465, mark_end480, mark_end510, mark_end540, mark_end566, mark_end581, mark_end596, mark_end623, mark_end627, mark_end631, mark_end635, mark_end639, mark_end665, mark_end669, mark_end673, mark_end677, mark_end681, mark_end685 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i481, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp18, v20, cmp22, v21, cmp26, v22, cmp30, v23, cmp32, v24, cmp34, v25, cmp37, v26, cmp40, v27, cmp43, v28, cmp47, v29, cmp50, v30, cmp54, v31, cmp57, v32, cmp60, v33, cmp63, v34, cmp66, v35, loadedv70, v36, cmp72, v37, cmp76, v38, cmp80, v39, cmp83, v40, loadedv87, v41, cmp89, v42, cmp93, v43, cmp96, v44, loadedv100, v45, cmp102, v46, cmp106, v47, cmp109, v48, loadedv113, v49, cmp115, v50, cmp119, v51, cmp122, v52, loadedv126, v53, cmp128, v54, cmp132, v55, cmp135, v56, loadedv139, v57, cmp141, v58, cmp145, v59, cmp148, v60, cmp151, v61, cmp154, v62, cmp157, v63, cmp160, v64, loadedv164, v65, cmp166, v66, cmp169, v67, loadedv173, v68, cmp175, v69, cmp178, v70, loadedv182, v71, cmp184, v72, cmp187, v73, loadedv191, v74, cmp193, v75, cmp196, v76, loadedv200, v77, cmp202, v78, cmp205, v79, loadedv209, v80, cmp211, v81, cmp214, v82, cmp217, v83, cmp220, v84, cmp223, v85, cmp226, v86, loadedv230, v87, loadedv232, v88, cmp235, v89, loadedv239, v90, loadedv241, v91, cmp244, v92, cmp248, v93, loadedv252, v94, loadedv254, v95, cmp257, v96, loadedv261, v97, loadedv263, v98, cmp266, v99, cmp270, v100, loadedv274, v101, loadedv276, v102, cmp279, v103, cmp283, v104, cmp287, v105, cmp291, v106, cmp295, v107, cmp299, v108, cmp303, v109, cmp307, v110, cmp311, v111, cmp315, v112, cmp318, v113, cmp321, v114, cmp324, v115, cmp327, v116, cmp330, v117, cmp334, v118, cmp337, v119, cmp340, v120, cmp343, v121, cmp346, v122, cmp349, v123, cmp352, v124, loadedv356, v125, result_symbol, v126, mark_end, v127, v128, v129, loadedv358, v130, result_symbol360, v131, mark_end361, v132, v133, v134, cmp362, v135, cmp365, v136, cmp368, v137, cmp371, v138, cmp374, v139, cmp377, v140, cmp380, v141, loadedv384, v142, result_symbol386, v143, mark_end387, v144, v145, v146, cmp388, v147, cmp391, v148, loadedv395, v149, result_symbol397, v150, mark_end398, v151, v152, v153, cmp399, v154, cmp403, v155, cmp407, v156, cmp410, v157, cmp414, v158, cmp417, v159, loadedv421, v160, result_symbol423, v161, mark_end424, v162, v163, v164, cmp425, v165, cmp429, v166, cmp432, v167, cmp436, v168, cmp439, v169, loadedv443, v170, result_symbol445, v171, mark_end446, v172, v173, v174, cmp447, v175, cmp451, v176, cmp455, v177, cmp458, v178, loadedv462, v179, result_symbol464, v180, mark_end465, v181, v182, v183, cmp466, v184, cmp470, v185, cmp473, v186, loadedv477, v187, result_symbol479, v188, mark_end480, v189, v190, v191, conv483, cmp484, v192, idxprom487, arrayidx488, v193, conv489, v194, cmp490, v195, add493, idxprom494, arrayidx495, v196, v197, add498, v198, cmp500, v199, cmp503, v200, loadedv507, v201, result_symbol509, v202, mark_end510, v203, v204, v205, cmp511, v206, cmp515, v207, cmp519, v208, cmp523, v209, cmp526, v210, cmp530, v211, cmp533, v212, loadedv537, v213, result_symbol539, v214, mark_end540, v215, v216, v217, cmp541, v218, cmp545, v219, cmp549, v220, cmp552, v221, cmp556, v222, cmp559, v223, loadedv563, v224, result_symbol565, v225, mark_end566, v226, v227, v228, cmp567, v229, cmp571, v230, cmp574, v231, loadedv578, v232, result_symbol580, v233, mark_end581, v234, v235, v236, cmp582, v237, cmp586, v238, cmp589, v239, loadedv593, v240, result_symbol595, v241, mark_end596, v242, v243, v244, cmp597, v245, cmp601, v246, cmp604, v247, cmp607, v248, cmp610, v249, cmp613, v250, cmp616, v251, loadedv620, v252, result_symbol622, v253, mark_end623, v254, v255, v256, loadedv624, v257, result_symbol626, v258, mark_end627, v259, v260, v261, loadedv628, v262, result_symbol630, v263, mark_end631, v264, v265, v266, loadedv632, v267, result_symbol634, v268, mark_end635, v269, v270, v271, loadedv636, v272, result_symbol638, v273, mark_end639, v274, v275, v276, cmp640, v277, cmp643, v278, cmp646, v279, cmp649, v280, cmp652, v281, cmp655, v282, cmp658, v283, loadedv662, v284, result_symbol664, v285, mark_end665, v286, v287, v288, loadedv666, v289, result_symbol668, v290, mark_end669, v291, v292, v293, loadedv670, v294, result_symbol672, v295, mark_end673, v296, v297, v298, loadedv674, v299, result_symbol676, v300, mark_end677, v301, v302, v303, loadedv678, v304, result_symbol680, v305, mark_end681, v306, v307, v308, loadedv682, v309, result_symbol684, v310, mark_end685, v311, v312, v313, loadedv686, v314

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
	skip = libc.Ptr(&new(struct {
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
	i481 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[int16](state_addr) = state
	*libc.As[byte](result) = 0
	*libc.As[byte](skip) = 0
	*libc.As[byte](eof) = 0
	goto start

next_state:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	v3 = *libc.As[byte](skip)
	loadedv = (v3 & 1) != 0
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, loadedv)
	goto start

start:
	*libc.As[byte](skip) = 0
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
		goto sw_bb71
	case 2:
		goto sw_bb88
	case 3:
		goto sw_bb101
	case 4:
		goto sw_bb114
	case 5:
		goto sw_bb127
	case 6:
		goto sw_bb140
	case 7:
		goto sw_bb165
	case 8:
		goto sw_bb174
	case 9:
		goto sw_bb183
	case 10:
		goto sw_bb192
	case 11:
		goto sw_bb201
	case 12:
		goto sw_bb210
	case 13:
		goto sw_bb231
	case 14:
		goto sw_bb240
	case 15:
		goto sw_bb253
	case 16:
		goto sw_bb262
	case 17:
		goto sw_bb275
	case 18:
		goto sw_bb357
	case 19:
		goto sw_bb359
	case 20:
		goto sw_bb385
	case 21:
		goto sw_bb396
	case 22:
		goto sw_bb422
	case 23:
		goto sw_bb444
	case 24:
		goto sw_bb463
	case 25:
		goto sw_bb478
	case 26:
		goto sw_bb508
	case 27:
		goto sw_bb538
	case 28:
		goto sw_bb564
	case 29:
		goto sw_bb579
	case 30:
		goto sw_bb594
	case 31:
		goto sw_bb621
	case 32:
		goto sw_bb625
	case 33:
		goto sw_bb629
	case 34:
		goto sw_bb633
	case 35:
		goto sw_bb637
	case 36:
		goto sw_bb663
	case 37:
		goto sw_bb667
	case 38:
		goto sw_bb671
	case 39:
		goto sw_bb675
	case 40:
		goto sw_bb679
	case 41:
		goto sw_bb683
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
	cmp = uint64(conv4) < uint64(22)
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
	cmp14 = v18 == 92
	if cmp14 {
		goto if_then16
	} else {
		goto if_end17
	}

if_then16:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end17:
	v19 = *libc.As[int32](lookahead)
	cmp18 = v19 == 93
	if cmp18 {
		goto if_then20
	} else {
		goto if_end21
	}

if_then20:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end21:
	v20 = *libc.As[int32](lookahead)
	cmp22 = v20 == 123
	if cmp22 {
		goto if_then24
	} else {
		goto if_end25
	}

if_then24:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end25:
	v21 = *libc.As[int32](lookahead)
	cmp26 = v21 == 125
	if cmp26 {
		goto if_then28
	} else {
		goto if_end29
	}

if_then28:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end29:
	v22 = *libc.As[int32](lookahead)
	cmp30 = 9 <= v22
	if cmp30 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false
	}

land_lhs_true:
	v23 = *libc.As[int32](lookahead)
	cmp32 = v23 <= 13
	if cmp32 {
		goto if_then45
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v24 = *libc.As[int32](lookahead)
	cmp34 = v24 == 32
	if cmp34 {
		goto if_then45
	} else {
		goto lor_lhs_false36
	}

lor_lhs_false36:
	v25 = *libc.As[int32](lookahead)
	cmp37 = v25 == 8203
	if cmp37 {
		goto if_then45
	} else {
		goto lor_lhs_false39
	}

lor_lhs_false39:
	v26 = *libc.As[int32](lookahead)
	cmp40 = v26 == 8288
	if cmp40 {
		goto if_then45
	} else {
		goto lor_lhs_false42
	}

lor_lhs_false42:
	v27 = *libc.As[int32](lookahead)
	cmp43 = v27 == 65279
	if cmp43 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end46:
	v28 = *libc.As[int32](lookahead)
	cmp47 = 49 <= v28
	if cmp47 {
		goto land_lhs_true49
	} else {
		goto if_end53
	}

land_lhs_true49:
	v29 = *libc.As[int32](lookahead)
	cmp50 = v29 <= 57
	if cmp50 {
		goto if_then52
	} else {
		goto if_end53
	}

if_then52:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end53:
	v30 = *libc.As[int32](lookahead)
	cmp54 = 65 <= v30
	if cmp54 {
		goto land_lhs_true56
	} else {
		goto lor_lhs_false59
	}

land_lhs_true56:
	v31 = *libc.As[int32](lookahead)
	cmp57 = v31 <= 90
	if cmp57 {
		goto if_then68
	} else {
		goto lor_lhs_false59
	}

lor_lhs_false59:
	v32 = *libc.As[int32](lookahead)
	cmp60 = v32 == 95
	if cmp60 {
		goto if_then68
	} else {
		goto lor_lhs_false62
	}

lor_lhs_false62:
	v33 = *libc.As[int32](lookahead)
	cmp63 = 97 <= v33
	if cmp63 {
		goto land_lhs_true65
	} else {
		goto if_end69
	}

land_lhs_true65:
	v34 = *libc.As[int32](lookahead)
	cmp66 = v34 <= 122
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end69:
	v35 = *libc.As[byte](result)
	loadedv70 = (v35 & 1) != 0
	*libc.As[bool](retval) = loadedv70
	goto _return

sw_bb71:
	v36 = *libc.As[int32](lookahead)
	cmp72 = v36 == 43
	if cmp72 {
		goto if_then74
	} else {
		goto if_end75
	}

if_then74:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end75:
	v37 = *libc.As[int32](lookahead)
	cmp76 = v37 == 45
	if cmp76 {
		goto if_then78
	} else {
		goto if_end79
	}

if_then78:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end79:
	v38 = *libc.As[int32](lookahead)
	cmp80 = 48 <= v38
	if cmp80 {
		goto land_lhs_true82
	} else {
		goto if_end86
	}

land_lhs_true82:
	v39 = *libc.As[int32](lookahead)
	cmp83 = v39 <= 57
	if cmp83 {
		goto if_then85
	} else {
		goto if_end86
	}

if_then85:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end86:
	v40 = *libc.As[byte](result)
	loadedv87 = (v40 & 1) != 0
	*libc.As[bool](retval) = loadedv87
	goto _return

sw_bb88:
	v41 = *libc.As[int32](lookahead)
	cmp89 = v41 == 45
	if cmp89 {
		goto if_then91
	} else {
		goto if_end92
	}

if_then91:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end92:
	v42 = *libc.As[int32](lookahead)
	cmp93 = 48 <= v42
	if cmp93 {
		goto land_lhs_true95
	} else {
		goto if_end99
	}

land_lhs_true95:
	v43 = *libc.As[int32](lookahead)
	cmp96 = v43 <= 57
	if cmp96 {
		goto if_then98
	} else {
		goto if_end99
	}

if_then98:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end99:
	v44 = *libc.As[byte](result)
	loadedv100 = (v44 & 1) != 0
	*libc.As[bool](retval) = loadedv100
	goto _return

sw_bb101:
	v45 = *libc.As[int32](lookahead)
	cmp102 = v45 == 45
	if cmp102 {
		goto if_then104
	} else {
		goto if_end105
	}

if_then104:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end105:
	v46 = *libc.As[int32](lookahead)
	cmp106 = 48 <= v46
	if cmp106 {
		goto land_lhs_true108
	} else {
		goto if_end112
	}

land_lhs_true108:
	v47 = *libc.As[int32](lookahead)
	cmp109 = v47 <= 57
	if cmp109 {
		goto if_then111
	} else {
		goto if_end112
	}

if_then111:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end112:
	v48 = *libc.As[byte](result)
	loadedv113 = (v48 & 1) != 0
	*libc.As[bool](retval) = loadedv113
	goto _return

sw_bb114:
	v49 = *libc.As[int32](lookahead)
	cmp115 = v49 == 95
	if cmp115 {
		goto if_then117
	} else {
		goto if_end118
	}

if_then117:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end118:
	v50 = *libc.As[int32](lookahead)
	cmp119 = v50 == 48
	if cmp119 {
		goto if_then124
	} else {
		goto lor_lhs_false121
	}

lor_lhs_false121:
	v51 = *libc.As[int32](lookahead)
	cmp122 = v51 == 49
	if cmp122 {
		goto if_then124
	} else {
		goto if_end125
	}

if_then124:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end125:
	v52 = *libc.As[byte](result)
	loadedv126 = (v52 & 1) != 0
	*libc.As[bool](retval) = loadedv126
	goto _return

sw_bb127:
	v53 = *libc.As[int32](lookahead)
	cmp128 = v53 == 95
	if cmp128 {
		goto if_then130
	} else {
		goto if_end131
	}

if_then130:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end131:
	v54 = *libc.As[int32](lookahead)
	cmp132 = 48 <= v54
	if cmp132 {
		goto land_lhs_true134
	} else {
		goto if_end138
	}

land_lhs_true134:
	v55 = *libc.As[int32](lookahead)
	cmp135 = v55 <= 55
	if cmp135 {
		goto if_then137
	} else {
		goto if_end138
	}

if_then137:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end138:
	v56 = *libc.As[byte](result)
	loadedv139 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv139
	goto _return

sw_bb140:
	v57 = *libc.As[int32](lookahead)
	cmp141 = v57 == 95
	if cmp141 {
		goto if_then143
	} else {
		goto if_end144
	}

if_then143:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end144:
	v58 = *libc.As[int32](lookahead)
	cmp145 = 48 <= v58
	if cmp145 {
		goto land_lhs_true147
	} else {
		goto lor_lhs_false150
	}

land_lhs_true147:
	v59 = *libc.As[int32](lookahead)
	cmp148 = v59 <= 57
	if cmp148 {
		goto if_then162
	} else {
		goto lor_lhs_false150
	}

lor_lhs_false150:
	v60 = *libc.As[int32](lookahead)
	cmp151 = 65 <= v60
	if cmp151 {
		goto land_lhs_true153
	} else {
		goto lor_lhs_false156
	}

land_lhs_true153:
	v61 = *libc.As[int32](lookahead)
	cmp154 = v61 <= 70
	if cmp154 {
		goto if_then162
	} else {
		goto lor_lhs_false156
	}

lor_lhs_false156:
	v62 = *libc.As[int32](lookahead)
	cmp157 = 97 <= v62
	if cmp157 {
		goto land_lhs_true159
	} else {
		goto if_end163
	}

land_lhs_true159:
	v63 = *libc.As[int32](lookahead)
	cmp160 = v63 <= 102
	if cmp160 {
		goto if_then162
	} else {
		goto if_end163
	}

if_then162:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end163:
	v64 = *libc.As[byte](result)
	loadedv164 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv164
	goto _return

sw_bb165:
	v65 = *libc.As[int32](lookahead)
	cmp166 = v65 == 48
	if cmp166 {
		goto if_then171
	} else {
		goto lor_lhs_false168
	}

lor_lhs_false168:
	v66 = *libc.As[int32](lookahead)
	cmp169 = v66 == 49
	if cmp169 {
		goto if_then171
	} else {
		goto if_end172
	}

if_then171:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end172:
	v67 = *libc.As[byte](result)
	loadedv173 = (v67 & 1) != 0
	*libc.As[bool](retval) = loadedv173
	goto _return

sw_bb174:
	v68 = *libc.As[int32](lookahead)
	cmp175 = 48 <= v68
	if cmp175 {
		goto land_lhs_true177
	} else {
		goto if_end181
	}

land_lhs_true177:
	v69 = *libc.As[int32](lookahead)
	cmp178 = v69 <= 55
	if cmp178 {
		goto if_then180
	} else {
		goto if_end181
	}

if_then180:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end181:
	v70 = *libc.As[byte](result)
	loadedv182 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv182
	goto _return

sw_bb183:
	v71 = *libc.As[int32](lookahead)
	cmp184 = 48 <= v71
	if cmp184 {
		goto land_lhs_true186
	} else {
		goto if_end190
	}

land_lhs_true186:
	v72 = *libc.As[int32](lookahead)
	cmp187 = v72 <= 57
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end190:
	v73 = *libc.As[byte](result)
	loadedv191 = (v73 & 1) != 0
	*libc.As[bool](retval) = loadedv191
	goto _return

sw_bb192:
	v74 = *libc.As[int32](lookahead)
	cmp193 = 48 <= v74
	if cmp193 {
		goto land_lhs_true195
	} else {
		goto if_end199
	}

land_lhs_true195:
	v75 = *libc.As[int32](lookahead)
	cmp196 = v75 <= 57
	if cmp196 {
		goto if_then198
	} else {
		goto if_end199
	}

if_then198:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end199:
	v76 = *libc.As[byte](result)
	loadedv200 = (v76 & 1) != 0
	*libc.As[bool](retval) = loadedv200
	goto _return

sw_bb201:
	v77 = *libc.As[int32](lookahead)
	cmp202 = 48 <= v77
	if cmp202 {
		goto land_lhs_true204
	} else {
		goto if_end208
	}

land_lhs_true204:
	v78 = *libc.As[int32](lookahead)
	cmp205 = v78 <= 57
	if cmp205 {
		goto if_then207
	} else {
		goto if_end208
	}

if_then207:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end208:
	v79 = *libc.As[byte](result)
	loadedv209 = (v79 & 1) != 0
	*libc.As[bool](retval) = loadedv209
	goto _return

sw_bb210:
	v80 = *libc.As[int32](lookahead)
	cmp211 = 48 <= v80
	if cmp211 {
		goto land_lhs_true213
	} else {
		goto lor_lhs_false216
	}

land_lhs_true213:
	v81 = *libc.As[int32](lookahead)
	cmp214 = v81 <= 57
	if cmp214 {
		goto if_then228
	} else {
		goto lor_lhs_false216
	}

lor_lhs_false216:
	v82 = *libc.As[int32](lookahead)
	cmp217 = 65 <= v82
	if cmp217 {
		goto land_lhs_true219
	} else {
		goto lor_lhs_false222
	}

land_lhs_true219:
	v83 = *libc.As[int32](lookahead)
	cmp220 = v83 <= 70
	if cmp220 {
		goto if_then228
	} else {
		goto lor_lhs_false222
	}

lor_lhs_false222:
	v84 = *libc.As[int32](lookahead)
	cmp223 = 97 <= v84
	if cmp223 {
		goto land_lhs_true225
	} else {
		goto if_end229
	}

land_lhs_true225:
	v85 = *libc.As[int32](lookahead)
	cmp226 = v85 <= 102
	if cmp226 {
		goto if_then228
	} else {
		goto if_end229
	}

if_then228:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end229:
	v86 = *libc.As[byte](result)
	loadedv230 = (v86 & 1) != 0
	*libc.As[bool](retval) = loadedv230
	goto _return

sw_bb231:
	v87 = *libc.As[byte](eof)
	loadedv232 = (v87 & 1) != 0
	if loadedv232 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end234:
	v88 = *libc.As[int32](lookahead)
	cmp235 = v88 == 10
	if cmp235 {
		goto if_then237
	} else {
		goto if_end238
	}

if_then237:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end238:
	v89 = *libc.As[byte](result)
	loadedv239 = (v89 & 1) != 0
	*libc.As[bool](retval) = loadedv239
	goto _return

sw_bb240:
	v90 = *libc.As[byte](eof)
	loadedv241 = (v90 & 1) != 0
	if loadedv241 {
		goto if_then242
	} else {
		goto if_end243
	}

if_then242:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end243:
	v91 = *libc.As[int32](lookahead)
	cmp244 = v91 == 10
	if cmp244 {
		goto if_then246
	} else {
		goto if_end247
	}

if_then246:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end247:
	v92 = *libc.As[int32](lookahead)
	cmp248 = v92 == 13
	if cmp248 {
		goto if_then250
	} else {
		goto if_end251
	}

if_then250:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end251:
	v93 = *libc.As[byte](result)
	loadedv252 = (v93 & 1) != 0
	*libc.As[bool](retval) = loadedv252
	goto _return

sw_bb253:
	v94 = *libc.As[byte](eof)
	loadedv254 = (v94 & 1) != 0
	if loadedv254 {
		goto if_then255
	} else {
		goto if_end256
	}

if_then255:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end256:
	v95 = *libc.As[int32](lookahead)
	cmp257 = v95 == 10
	if cmp257 {
		goto if_then259
	} else {
		goto if_end260
	}

if_then259:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end260:
	v96 = *libc.As[byte](result)
	loadedv261 = (v96 & 1) != 0
	*libc.As[bool](retval) = loadedv261
	goto _return

sw_bb262:
	v97 = *libc.As[byte](eof)
	loadedv263 = (v97 & 1) != 0
	if loadedv263 {
		goto if_then264
	} else {
		goto if_end265
	}

if_then264:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end265:
	v98 = *libc.As[int32](lookahead)
	cmp266 = v98 == 10
	if cmp266 {
		goto if_then268
	} else {
		goto if_end269
	}

if_then268:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end269:
	v99 = *libc.As[int32](lookahead)
	cmp270 = v99 == 13
	if cmp270 {
		goto if_then272
	} else {
		goto if_end273
	}

if_then272:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end273:
	v100 = *libc.As[byte](result)
	loadedv274 = (v100 & 1) != 0
	*libc.As[bool](retval) = loadedv274
	goto _return

sw_bb275:
	v101 = *libc.As[byte](eof)
	loadedv276 = (v101 & 1) != 0
	if loadedv276 {
		goto if_then277
	} else {
		goto if_end278
	}

if_then277:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end278:
	v102 = *libc.As[int32](lookahead)
	cmp279 = v102 == 40
	if cmp279 {
		goto if_then281
	} else {
		goto if_end282
	}

if_then281:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end282:
	v103 = *libc.As[int32](lookahead)
	cmp283 = v103 == 41
	if cmp283 {
		goto if_then285
	} else {
		goto if_end286
	}

if_then285:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end286:
	v104 = *libc.As[int32](lookahead)
	cmp287 = v104 == 44
	if cmp287 {
		goto if_then289
	} else {
		goto if_end290
	}

if_then289:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end290:
	v105 = *libc.As[int32](lookahead)
	cmp291 = v105 == 58
	if cmp291 {
		goto if_then293
	} else {
		goto if_end294
	}

if_then293:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end294:
	v106 = *libc.As[int32](lookahead)
	cmp295 = v106 == 59
	if cmp295 {
		goto if_then297
	} else {
		goto if_end298
	}

if_then297:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end298:
	v107 = *libc.As[int32](lookahead)
	cmp299 = v107 == 91
	if cmp299 {
		goto if_then301
	} else {
		goto if_end302
	}

if_then301:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end302:
	v108 = *libc.As[int32](lookahead)
	cmp303 = v108 == 92
	if cmp303 {
		goto if_then305
	} else {
		goto if_end306
	}

if_then305:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end306:
	v109 = *libc.As[int32](lookahead)
	cmp307 = v109 == 93
	if cmp307 {
		goto if_then309
	} else {
		goto if_end310
	}

if_then309:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end310:
	v110 = *libc.As[int32](lookahead)
	cmp311 = v110 == 125
	if cmp311 {
		goto if_then313
	} else {
		goto if_end314
	}

if_then313:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end314:
	v111 = *libc.As[int32](lookahead)
	cmp315 = 9 <= v111
	if cmp315 {
		goto land_lhs_true317
	} else {
		goto lor_lhs_false320
	}

land_lhs_true317:
	v112 = *libc.As[int32](lookahead)
	cmp318 = v112 <= 13
	if cmp318 {
		goto if_then332
	} else {
		goto lor_lhs_false320
	}

lor_lhs_false320:
	v113 = *libc.As[int32](lookahead)
	cmp321 = v113 == 32
	if cmp321 {
		goto if_then332
	} else {
		goto lor_lhs_false323
	}

lor_lhs_false323:
	v114 = *libc.As[int32](lookahead)
	cmp324 = v114 == 8203
	if cmp324 {
		goto if_then332
	} else {
		goto lor_lhs_false326
	}

lor_lhs_false326:
	v115 = *libc.As[int32](lookahead)
	cmp327 = v115 == 8288
	if cmp327 {
		goto if_then332
	} else {
		goto lor_lhs_false329
	}

lor_lhs_false329:
	v116 = *libc.As[int32](lookahead)
	cmp330 = v116 == 65279
	if cmp330 {
		goto if_then332
	} else {
		goto if_end333
	}

if_then332:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end333:
	v117 = *libc.As[int32](lookahead)
	cmp334 = 48 <= v117
	if cmp334 {
		goto land_lhs_true336
	} else {
		goto lor_lhs_false339
	}

land_lhs_true336:
	v118 = *libc.As[int32](lookahead)
	cmp337 = v118 <= 57
	if cmp337 {
		goto if_then354
	} else {
		goto lor_lhs_false339
	}

lor_lhs_false339:
	v119 = *libc.As[int32](lookahead)
	cmp340 = 65 <= v119
	if cmp340 {
		goto land_lhs_true342
	} else {
		goto lor_lhs_false345
	}

land_lhs_true342:
	v120 = *libc.As[int32](lookahead)
	cmp343 = v120 <= 90
	if cmp343 {
		goto if_then354
	} else {
		goto lor_lhs_false345
	}

lor_lhs_false345:
	v121 = *libc.As[int32](lookahead)
	cmp346 = v121 == 95
	if cmp346 {
		goto if_then354
	} else {
		goto lor_lhs_false348
	}

lor_lhs_false348:
	v122 = *libc.As[int32](lookahead)
	cmp349 = 97 <= v122
	if cmp349 {
		goto land_lhs_true351
	} else {
		goto if_end355
	}

land_lhs_true351:
	v123 = *libc.As[int32](lookahead)
	cmp352 = v123 <= 122
	if cmp352 {
		goto if_then354
	} else {
		goto if_end355
	}

if_then354:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end355:
	v124 = *libc.As[byte](result)
	loadedv356 = (v124 & 1) != 0
	*libc.As[bool](retval) = loadedv356
	goto _return

sw_bb357:
	*libc.As[byte](result) = 1
	v125 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v125).F1)
	*libc.As[int16](result_symbol) = 0
	v126 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v126).F3)
	v127 = *libc.As[unsafe.Pointer](mark_end)
	v128 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v127)(v128)
	v129 = *libc.As[byte](result)
	loadedv358 = (v129 & 1) != 0
	*libc.As[bool](retval) = loadedv358
	goto _return

sw_bb359:
	*libc.As[byte](result) = 1
	v130 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol360 = libc.Ptr(&libc.As[TSLexer](v130).F1)
	*libc.As[int16](result_symbol360) = 1
	v131 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end361 = libc.Ptr(&libc.As[TSLexer](v131).F3)
	v132 = *libc.As[unsafe.Pointer](mark_end361)
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v132)(v133)
	v134 = *libc.As[int32](lookahead)
	cmp362 = 48 <= v134
	if cmp362 {
		goto land_lhs_true364
	} else {
		goto lor_lhs_false367
	}

land_lhs_true364:
	v135 = *libc.As[int32](lookahead)
	cmp365 = v135 <= 57
	if cmp365 {
		goto if_then382
	} else {
		goto lor_lhs_false367
	}

lor_lhs_false367:
	v136 = *libc.As[int32](lookahead)
	cmp368 = 65 <= v136
	if cmp368 {
		goto land_lhs_true370
	} else {
		goto lor_lhs_false373
	}

land_lhs_true370:
	v137 = *libc.As[int32](lookahead)
	cmp371 = v137 <= 90
	if cmp371 {
		goto if_then382
	} else {
		goto lor_lhs_false373
	}

lor_lhs_false373:
	v138 = *libc.As[int32](lookahead)
	cmp374 = v138 == 95
	if cmp374 {
		goto if_then382
	} else {
		goto lor_lhs_false376
	}

lor_lhs_false376:
	v139 = *libc.As[int32](lookahead)
	cmp377 = 97 <= v139
	if cmp377 {
		goto land_lhs_true379
	} else {
		goto if_end383
	}

land_lhs_true379:
	v140 = *libc.As[int32](lookahead)
	cmp380 = v140 <= 122
	if cmp380 {
		goto if_then382
	} else {
		goto if_end383
	}

if_then382:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end383:
	v141 = *libc.As[byte](result)
	loadedv384 = (v141 & 1) != 0
	*libc.As[bool](retval) = loadedv384
	goto _return

sw_bb385:
	*libc.As[byte](result) = 1
	v142 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol386 = libc.Ptr(&libc.As[TSLexer](v142).F1)
	*libc.As[int16](result_symbol386) = 2
	v143 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end387 = libc.Ptr(&libc.As[TSLexer](v143).F3)
	v144 = *libc.As[unsafe.Pointer](mark_end387)
	v145 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v144)(v145)
	v146 = *libc.As[int32](lookahead)
	cmp388 = v146 != 0
	if cmp388 {
		goto land_lhs_true390
	} else {
		goto if_end394
	}

land_lhs_true390:
	v147 = *libc.As[int32](lookahead)
	cmp391 = v147 != 10
	if cmp391 {
		goto if_then393
	} else {
		goto if_end394
	}

if_then393:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end394:
	v148 = *libc.As[byte](result)
	loadedv395 = (v148 & 1) != 0
	*libc.As[bool](retval) = loadedv395
	goto _return

sw_bb396:
	*libc.As[byte](result) = 1
	v149 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol397 = libc.Ptr(&libc.As[TSLexer](v149).F1)
	*libc.As[int16](result_symbol397) = 6
	v150 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end398 = libc.Ptr(&libc.As[TSLexer](v150).F3)
	v151 = *libc.As[unsafe.Pointer](mark_end398)
	v152 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v151)(v152)
	v153 = *libc.As[int32](lookahead)
	cmp399 = v153 == 45
	if cmp399 {
		goto if_then401
	} else {
		goto if_end402
	}

if_then401:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end402:
	v154 = *libc.As[int32](lookahead)
	cmp403 = v154 == 95
	if cmp403 {
		goto if_then405
	} else {
		goto if_end406
	}

if_then405:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end406:
	v155 = *libc.As[int32](lookahead)
	cmp407 = v155 == 69
	if cmp407 {
		goto if_then412
	} else {
		goto lor_lhs_false409
	}

lor_lhs_false409:
	v156 = *libc.As[int32](lookahead)
	cmp410 = v156 == 101
	if cmp410 {
		goto if_then412
	} else {
		goto if_end413
	}

if_then412:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end413:
	v157 = *libc.As[int32](lookahead)
	cmp414 = 48 <= v157
	if cmp414 {
		goto land_lhs_true416
	} else {
		goto if_end420
	}

land_lhs_true416:
	v158 = *libc.As[int32](lookahead)
	cmp417 = v158 <= 57
	if cmp417 {
		goto if_then419
	} else {
		goto if_end420
	}

if_then419:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end420:
	v159 = *libc.As[byte](result)
	loadedv421 = (v159 & 1) != 0
	*libc.As[bool](retval) = loadedv421
	goto _return

sw_bb422:
	*libc.As[byte](result) = 1
	v160 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol423 = libc.Ptr(&libc.As[TSLexer](v160).F1)
	*libc.As[int16](result_symbol423) = 6
	v161 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end424 = libc.Ptr(&libc.As[TSLexer](v161).F3)
	v162 = *libc.As[unsafe.Pointer](mark_end424)
	v163 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v162)(v163)
	v164 = *libc.As[int32](lookahead)
	cmp425 = v164 == 45
	if cmp425 {
		goto if_then427
	} else {
		goto if_end428
	}

if_then427:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end428:
	v165 = *libc.As[int32](lookahead)
	cmp429 = v165 == 69
	if cmp429 {
		goto if_then434
	} else {
		goto lor_lhs_false431
	}

lor_lhs_false431:
	v166 = *libc.As[int32](lookahead)
	cmp432 = v166 == 101
	if cmp432 {
		goto if_then434
	} else {
		goto if_end435
	}

if_then434:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end435:
	v167 = *libc.As[int32](lookahead)
	cmp436 = 48 <= v167
	if cmp436 {
		goto land_lhs_true438
	} else {
		goto if_end442
	}

land_lhs_true438:
	v168 = *libc.As[int32](lookahead)
	cmp439 = v168 <= 57
	if cmp439 {
		goto if_then441
	} else {
		goto if_end442
	}

if_then441:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end442:
	v169 = *libc.As[byte](result)
	loadedv443 = (v169 & 1) != 0
	*libc.As[bool](retval) = loadedv443
	goto _return

sw_bb444:
	*libc.As[byte](result) = 1
	v170 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol445 = libc.Ptr(&libc.As[TSLexer](v170).F1)
	*libc.As[int16](result_symbol445) = 6
	v171 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end446 = libc.Ptr(&libc.As[TSLexer](v171).F3)
	v172 = *libc.As[unsafe.Pointer](mark_end446)
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v172)(v173)
	v174 = *libc.As[int32](lookahead)
	cmp447 = v174 == 45
	if cmp447 {
		goto if_then449
	} else {
		goto if_end450
	}

if_then449:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end450:
	v175 = *libc.As[int32](lookahead)
	cmp451 = v175 == 95
	if cmp451 {
		goto if_then453
	} else {
		goto if_end454
	}

if_then453:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end454:
	v176 = *libc.As[int32](lookahead)
	cmp455 = 48 <= v176
	if cmp455 {
		goto land_lhs_true457
	} else {
		goto if_end461
	}

land_lhs_true457:
	v177 = *libc.As[int32](lookahead)
	cmp458 = v177 <= 57
	if cmp458 {
		goto if_then460
	} else {
		goto if_end461
	}

if_then460:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end461:
	v178 = *libc.As[byte](result)
	loadedv462 = (v178 & 1) != 0
	*libc.As[bool](retval) = loadedv462
	goto _return

sw_bb463:
	*libc.As[byte](result) = 1
	v179 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol464 = libc.Ptr(&libc.As[TSLexer](v179).F1)
	*libc.As[int16](result_symbol464) = 6
	v180 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end465 = libc.Ptr(&libc.As[TSLexer](v180).F3)
	v181 = *libc.As[unsafe.Pointer](mark_end465)
	v182 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v181)(v182)
	v183 = *libc.As[int32](lookahead)
	cmp466 = v183 == 45
	if cmp466 {
		goto if_then468
	} else {
		goto if_end469
	}

if_then468:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end469:
	v184 = *libc.As[int32](lookahead)
	cmp470 = 48 <= v184
	if cmp470 {
		goto land_lhs_true472
	} else {
		goto if_end476
	}

land_lhs_true472:
	v185 = *libc.As[int32](lookahead)
	cmp473 = v185 <= 57
	if cmp473 {
		goto if_then475
	} else {
		goto if_end476
	}

if_then475:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end476:
	v186 = *libc.As[byte](result)
	loadedv477 = (v186 & 1) != 0
	*libc.As[bool](retval) = loadedv477
	goto _return

sw_bb478:
	*libc.As[byte](result) = 1
	v187 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol479 = libc.Ptr(&libc.As[TSLexer](v187).F1)
	*libc.As[int16](result_symbol479) = 7
	v188 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end480 = libc.Ptr(&libc.As[TSLexer](v188).F3)
	v189 = *libc.As[unsafe.Pointer](mark_end480)
	v190 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v189)(v190)
	*libc.As[int32](i481) = 0
	goto for_cond482

for_cond482:
	v191 = *libc.As[int32](i481)
	conv483 = int64(uint64(uint32(v191)))
	cmp484 = uint64(conv483) < uint64(22)
	if cmp484 {
		goto for_body486
	} else {
		goto for_end499
	}

for_body486:
	v192 = *libc.As[int32](i481)
	idxprom487 = int64(uint64(uint32(v192)))
	arrayidx488 = libc.Ptr(&ts_lex_map_42[idxprom487])
	v193 = *libc.As[int16](arrayidx488)
	conv489 = int32(uint32(uint16(v193)))
	v194 = *libc.As[int32](lookahead)
	cmp490 = conv489 == v194
	if cmp490 {
		goto if_then492
	} else {
		goto if_end496
	}

if_then492:
	v195 = *libc.As[int32](i481)
	add493 = v195 + 1
	idxprom494 = int64(uint64(uint32(add493)))
	arrayidx495 = libc.Ptr(&ts_lex_map_42[idxprom494])
	v196 = *libc.As[int16](arrayidx495)
	*libc.As[int16](state_addr) = v196
	goto next_state

if_end496:
	goto for_inc497

for_inc497:
	v197 = *libc.As[int32](i481)
	add498 = v197 + 2
	*libc.As[int32](i481) = add498
	goto for_cond482

for_end499:
	v198 = *libc.As[int32](lookahead)
	cmp500 = 48 <= v198
	if cmp500 {
		goto land_lhs_true502
	} else {
		goto if_end506
	}

land_lhs_true502:
	v199 = *libc.As[int32](lookahead)
	cmp503 = v199 <= 57
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end506:
	v200 = *libc.As[byte](result)
	loadedv507 = (v200 & 1) != 0
	*libc.As[bool](retval) = loadedv507
	goto _return

sw_bb508:
	*libc.As[byte](result) = 1
	v201 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol509 = libc.Ptr(&libc.As[TSLexer](v201).F1)
	*libc.As[int16](result_symbol509) = 7
	v202 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end510 = libc.Ptr(&libc.As[TSLexer](v202).F3)
	v203 = *libc.As[unsafe.Pointer](mark_end510)
	v204 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v203)(v204)
	v205 = *libc.As[int32](lookahead)
	cmp511 = v205 == 45
	if cmp511 {
		goto if_then513
	} else {
		goto if_end514
	}

if_then513:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end514:
	v206 = *libc.As[int32](lookahead)
	cmp515 = v206 == 46
	if cmp515 {
		goto if_then517
	} else {
		goto if_end518
	}

if_then517:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end518:
	v207 = *libc.As[int32](lookahead)
	cmp519 = v207 == 95
	if cmp519 {
		goto if_then521
	} else {
		goto if_end522
	}

if_then521:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end522:
	v208 = *libc.As[int32](lookahead)
	cmp523 = v208 == 69
	if cmp523 {
		goto if_then528
	} else {
		goto lor_lhs_false525
	}

lor_lhs_false525:
	v209 = *libc.As[int32](lookahead)
	cmp526 = v209 == 101
	if cmp526 {
		goto if_then528
	} else {
		goto if_end529
	}

if_then528:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end529:
	v210 = *libc.As[int32](lookahead)
	cmp530 = 48 <= v210
	if cmp530 {
		goto land_lhs_true532
	} else {
		goto if_end536
	}

land_lhs_true532:
	v211 = *libc.As[int32](lookahead)
	cmp533 = v211 <= 57
	if cmp533 {
		goto if_then535
	} else {
		goto if_end536
	}

if_then535:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end536:
	v212 = *libc.As[byte](result)
	loadedv537 = (v212 & 1) != 0
	*libc.As[bool](retval) = loadedv537
	goto _return

sw_bb538:
	*libc.As[byte](result) = 1
	v213 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol539 = libc.Ptr(&libc.As[TSLexer](v213).F1)
	*libc.As[int16](result_symbol539) = 7
	v214 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end540 = libc.Ptr(&libc.As[TSLexer](v214).F3)
	v215 = *libc.As[unsafe.Pointer](mark_end540)
	v216 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v215)(v216)
	v217 = *libc.As[int32](lookahead)
	cmp541 = v217 == 45
	if cmp541 {
		goto if_then543
	} else {
		goto if_end544
	}

if_then543:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end544:
	v218 = *libc.As[int32](lookahead)
	cmp545 = v218 == 46
	if cmp545 {
		goto if_then547
	} else {
		goto if_end548
	}

if_then547:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end548:
	v219 = *libc.As[int32](lookahead)
	cmp549 = v219 == 69
	if cmp549 {
		goto if_then554
	} else {
		goto lor_lhs_false551
	}

lor_lhs_false551:
	v220 = *libc.As[int32](lookahead)
	cmp552 = v220 == 101
	if cmp552 {
		goto if_then554
	} else {
		goto if_end555
	}

if_then554:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end555:
	v221 = *libc.As[int32](lookahead)
	cmp556 = 48 <= v221
	if cmp556 {
		goto land_lhs_true558
	} else {
		goto if_end562
	}

land_lhs_true558:
	v222 = *libc.As[int32](lookahead)
	cmp559 = v222 <= 57
	if cmp559 {
		goto if_then561
	} else {
		goto if_end562
	}

if_then561:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end562:
	v223 = *libc.As[byte](result)
	loadedv563 = (v223 & 1) != 0
	*libc.As[bool](retval) = loadedv563
	goto _return

sw_bb564:
	*libc.As[byte](result) = 1
	v224 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol565 = libc.Ptr(&libc.As[TSLexer](v224).F1)
	*libc.As[int16](result_symbol565) = 7
	v225 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end566 = libc.Ptr(&libc.As[TSLexer](v225).F3)
	v226 = *libc.As[unsafe.Pointer](mark_end566)
	v227 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v226)(v227)
	v228 = *libc.As[int32](lookahead)
	cmp567 = v228 == 95
	if cmp567 {
		goto if_then569
	} else {
		goto if_end570
	}

if_then569:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end570:
	v229 = *libc.As[int32](lookahead)
	cmp571 = v229 == 48
	if cmp571 {
		goto if_then576
	} else {
		goto lor_lhs_false573
	}

lor_lhs_false573:
	v230 = *libc.As[int32](lookahead)
	cmp574 = v230 == 49
	if cmp574 {
		goto if_then576
	} else {
		goto if_end577
	}

if_then576:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end577:
	v231 = *libc.As[byte](result)
	loadedv578 = (v231 & 1) != 0
	*libc.As[bool](retval) = loadedv578
	goto _return

sw_bb579:
	*libc.As[byte](result) = 1
	v232 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol580 = libc.Ptr(&libc.As[TSLexer](v232).F1)
	*libc.As[int16](result_symbol580) = 7
	v233 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end581 = libc.Ptr(&libc.As[TSLexer](v233).F3)
	v234 = *libc.As[unsafe.Pointer](mark_end581)
	v235 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v234)(v235)
	v236 = *libc.As[int32](lookahead)
	cmp582 = v236 == 95
	if cmp582 {
		goto if_then584
	} else {
		goto if_end585
	}

if_then584:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end585:
	v237 = *libc.As[int32](lookahead)
	cmp586 = 48 <= v237
	if cmp586 {
		goto land_lhs_true588
	} else {
		goto if_end592
	}

land_lhs_true588:
	v238 = *libc.As[int32](lookahead)
	cmp589 = v238 <= 55
	if cmp589 {
		goto if_then591
	} else {
		goto if_end592
	}

if_then591:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end592:
	v239 = *libc.As[byte](result)
	loadedv593 = (v239 & 1) != 0
	*libc.As[bool](retval) = loadedv593
	goto _return

sw_bb594:
	*libc.As[byte](result) = 1
	v240 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol595 = libc.Ptr(&libc.As[TSLexer](v240).F1)
	*libc.As[int16](result_symbol595) = 7
	v241 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end596 = libc.Ptr(&libc.As[TSLexer](v241).F3)
	v242 = *libc.As[unsafe.Pointer](mark_end596)
	v243 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v242)(v243)
	v244 = *libc.As[int32](lookahead)
	cmp597 = v244 == 95
	if cmp597 {
		goto if_then599
	} else {
		goto if_end600
	}

if_then599:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end600:
	v245 = *libc.As[int32](lookahead)
	cmp601 = 48 <= v245
	if cmp601 {
		goto land_lhs_true603
	} else {
		goto lor_lhs_false606
	}

land_lhs_true603:
	v246 = *libc.As[int32](lookahead)
	cmp604 = v246 <= 57
	if cmp604 {
		goto if_then618
	} else {
		goto lor_lhs_false606
	}

lor_lhs_false606:
	v247 = *libc.As[int32](lookahead)
	cmp607 = 65 <= v247
	if cmp607 {
		goto land_lhs_true609
	} else {
		goto lor_lhs_false612
	}

land_lhs_true609:
	v248 = *libc.As[int32](lookahead)
	cmp610 = v248 <= 70
	if cmp610 {
		goto if_then618
	} else {
		goto lor_lhs_false612
	}

lor_lhs_false612:
	v249 = *libc.As[int32](lookahead)
	cmp613 = 97 <= v249
	if cmp613 {
		goto land_lhs_true615
	} else {
		goto if_end619
	}

land_lhs_true615:
	v250 = *libc.As[int32](lookahead)
	cmp616 = v250 <= 102
	if cmp616 {
		goto if_then618
	} else {
		goto if_end619
	}

if_then618:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end619:
	v251 = *libc.As[byte](result)
	loadedv620 = (v251 & 1) != 0
	*libc.As[bool](retval) = loadedv620
	goto _return

sw_bb621:
	*libc.As[byte](result) = 1
	v252 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol622 = libc.Ptr(&libc.As[TSLexer](v252).F1)
	*libc.As[int16](result_symbol622) = 8
	v253 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end623 = libc.Ptr(&libc.As[TSLexer](v253).F3)
	v254 = *libc.As[unsafe.Pointer](mark_end623)
	v255 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v254)(v255)
	v256 = *libc.As[byte](result)
	loadedv624 = (v256 & 1) != 0
	*libc.As[bool](retval) = loadedv624
	goto _return

sw_bb625:
	*libc.As[byte](result) = 1
	v257 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol626 = libc.Ptr(&libc.As[TSLexer](v257).F1)
	*libc.As[int16](result_symbol626) = 9
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end627 = libc.Ptr(&libc.As[TSLexer](v258).F3)
	v259 = *libc.As[unsafe.Pointer](mark_end627)
	v260 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v259)(v260)
	v261 = *libc.As[byte](result)
	loadedv628 = (v261 & 1) != 0
	*libc.As[bool](retval) = loadedv628
	goto _return

sw_bb629:
	*libc.As[byte](result) = 1
	v262 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol630 = libc.Ptr(&libc.As[TSLexer](v262).F1)
	*libc.As[int16](result_symbol630) = 10
	v263 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end631 = libc.Ptr(&libc.As[TSLexer](v263).F3)
	v264 = *libc.As[unsafe.Pointer](mark_end631)
	v265 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v264)(v265)
	v266 = *libc.As[byte](result)
	loadedv632 = (v266 & 1) != 0
	*libc.As[bool](retval) = loadedv632
	goto _return

sw_bb633:
	*libc.As[byte](result) = 1
	v267 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol634 = libc.Ptr(&libc.As[TSLexer](v267).F1)
	*libc.As[int16](result_symbol634) = 11
	v268 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end635 = libc.Ptr(&libc.As[TSLexer](v268).F3)
	v269 = *libc.As[unsafe.Pointer](mark_end635)
	v270 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v269)(v270)
	v271 = *libc.As[byte](result)
	loadedv636 = (v271 & 1) != 0
	*libc.As[bool](retval) = loadedv636
	goto _return

sw_bb637:
	*libc.As[byte](result) = 1
	v272 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol638 = libc.Ptr(&libc.As[TSLexer](v272).F1)
	*libc.As[int16](result_symbol638) = 12
	v273 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end639 = libc.Ptr(&libc.As[TSLexer](v273).F3)
	v274 = *libc.As[unsafe.Pointer](mark_end639)
	v275 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v274)(v275)
	v276 = *libc.As[int32](lookahead)
	cmp640 = 47 <= v276
	if cmp640 {
		goto land_lhs_true642
	} else {
		goto lor_lhs_false645
	}

land_lhs_true642:
	v277 = *libc.As[int32](lookahead)
	cmp643 = v277 <= 58
	if cmp643 {
		goto if_then660
	} else {
		goto lor_lhs_false645
	}

lor_lhs_false645:
	v278 = *libc.As[int32](lookahead)
	cmp646 = 65 <= v278
	if cmp646 {
		goto land_lhs_true648
	} else {
		goto lor_lhs_false651
	}

land_lhs_true648:
	v279 = *libc.As[int32](lookahead)
	cmp649 = v279 <= 90
	if cmp649 {
		goto if_then660
	} else {
		goto lor_lhs_false651
	}

lor_lhs_false651:
	v280 = *libc.As[int32](lookahead)
	cmp652 = v280 == 95
	if cmp652 {
		goto if_then660
	} else {
		goto lor_lhs_false654
	}

lor_lhs_false654:
	v281 = *libc.As[int32](lookahead)
	cmp655 = 97 <= v281
	if cmp655 {
		goto land_lhs_true657
	} else {
		goto if_end661
	}

land_lhs_true657:
	v282 = *libc.As[int32](lookahead)
	cmp658 = v282 <= 122
	if cmp658 {
		goto if_then660
	} else {
		goto if_end661
	}

if_then660:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end661:
	v283 = *libc.As[byte](result)
	loadedv662 = (v283 & 1) != 0
	*libc.As[bool](retval) = loadedv662
	goto _return

sw_bb663:
	*libc.As[byte](result) = 1
	v284 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol664 = libc.Ptr(&libc.As[TSLexer](v284).F1)
	*libc.As[int16](result_symbol664) = 13
	v285 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end665 = libc.Ptr(&libc.As[TSLexer](v285).F3)
	v286 = *libc.As[unsafe.Pointer](mark_end665)
	v287 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v286)(v287)
	v288 = *libc.As[byte](result)
	loadedv666 = (v288 & 1) != 0
	*libc.As[bool](retval) = loadedv666
	goto _return

sw_bb667:
	*libc.As[byte](result) = 1
	v289 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol668 = libc.Ptr(&libc.As[TSLexer](v289).F1)
	*libc.As[int16](result_symbol668) = 14
	v290 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end669 = libc.Ptr(&libc.As[TSLexer](v290).F3)
	v291 = *libc.As[unsafe.Pointer](mark_end669)
	v292 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v291)(v292)
	v293 = *libc.As[byte](result)
	loadedv670 = (v293 & 1) != 0
	*libc.As[bool](retval) = loadedv670
	goto _return

sw_bb671:
	*libc.As[byte](result) = 1
	v294 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol672 = libc.Ptr(&libc.As[TSLexer](v294).F1)
	*libc.As[int16](result_symbol672) = 15
	v295 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end673 = libc.Ptr(&libc.As[TSLexer](v295).F3)
	v296 = *libc.As[unsafe.Pointer](mark_end673)
	v297 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v296)(v297)
	v298 = *libc.As[byte](result)
	loadedv674 = (v298 & 1) != 0
	*libc.As[bool](retval) = loadedv674
	goto _return

sw_bb675:
	*libc.As[byte](result) = 1
	v299 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol676 = libc.Ptr(&libc.As[TSLexer](v299).F1)
	*libc.As[int16](result_symbol676) = 16
	v300 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end677 = libc.Ptr(&libc.As[TSLexer](v300).F3)
	v301 = *libc.As[unsafe.Pointer](mark_end677)
	v302 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v301)(v302)
	v303 = *libc.As[byte](result)
	loadedv678 = (v303 & 1) != 0
	*libc.As[bool](retval) = loadedv678
	goto _return

sw_bb679:
	*libc.As[byte](result) = 1
	v304 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol680 = libc.Ptr(&libc.As[TSLexer](v304).F1)
	*libc.As[int16](result_symbol680) = 17
	v305 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end681 = libc.Ptr(&libc.As[TSLexer](v305).F3)
	v306 = *libc.As[unsafe.Pointer](mark_end681)
	v307 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v306)(v307)
	v308 = *libc.As[byte](result)
	loadedv682 = (v308 & 1) != 0
	*libc.As[bool](retval) = loadedv682
	goto _return

sw_bb683:
	*libc.As[byte](result) = 1
	v309 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol684 = libc.Ptr(&libc.As[TSLexer](v309).F1)
	*libc.As[int16](result_symbol684) = 18
	v310 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end685 = libc.Ptr(&libc.As[TSLexer](v310).F3)
	v311 = *libc.As[unsafe.Pointer](mark_end685)
	v312 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v311)(v312)
	v313 = *libc.As[byte](result)
	loadedv686 = (v313 & 1) != 0
	*libc.As[bool](retval) = loadedv686
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v314 = *libc.As[bool](retval)
	return v314
}
func ts_lex_keywords(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, cmp, cmp4, cmp8, cmp12, cmp16, cmp18, cmp20, cmp23, cmp26, cmp29, loadedv33, cmp35, cmp39, loadedv43, cmp45, loadedv49, cmp51, loadedv55, cmp57, loadedv61, cmp63, loadedv67, cmp69, loadedv73, cmp75, loadedv79, cmp81, loadedv85, cmp87, loadedv91, cmp93, loadedv97, cmp99, loadedv103, cmp105, loadedv109, loadedv111, loadedv115, loadedv119, v61 bool
	var retval unsafe.Pointer
	var v9 int16
	var state_addr, result_symbol, result_symbol113, result_symbol117 unsafe.Pointer
	var v5, conv, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v21, v22, v24, v26, v28, v30, v32, v34, v36, v38, v40, v42, v44 int32
	var lookahead, lookahead1 unsafe.Pointer
	var v3, storedv, v20, v23, v25, v27, v29, v31, v33, v35, v37, v39, v41, v43, v45, v50, v55, v60 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v46, v47, v48, v49, v51, v52, v53, v54, v56, v57, v58, v59 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end114, mark_end118 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, cmp, v11, cmp4, v12, cmp8, v13, cmp12, v14, cmp16, v15, cmp18, v16, cmp20, v17, cmp23, v18, cmp26, v19, cmp29, v20, loadedv33, v21, cmp35, v22, cmp39, v23, loadedv43, v24, cmp45, v25, loadedv49, v26, cmp51, v27, loadedv55, v28, cmp57, v29, loadedv61, v30, cmp63, v31, loadedv67, v32, cmp69, v33, loadedv73, v34, cmp75, v35, loadedv79, v36, cmp81, v37, loadedv85, v38, cmp87, v39, loadedv91, v40, cmp93, v41, loadedv97, v42, cmp99, v43, loadedv103, v44, cmp105, v45, loadedv109, v46, result_symbol, v47, mark_end, v48, v49, v50, loadedv111, v51, result_symbol113, v52, mark_end114, v53, v54, v55, loadedv115, v56, result_symbol117, v57, mark_end118, v58, v59, v60, loadedv119, v61

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
	skip = libc.Ptr(&new(struct {
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
	*libc.As[byte](skip) = 0
	*libc.As[byte](eof) = 0
	goto start

next_state:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	v3 = *libc.As[byte](skip)
	loadedv = (v3 & 1) != 0
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, loadedv)
	goto start

start:
	*libc.As[byte](skip) = 0
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
		goto sw_bb34
	case 2:
		goto sw_bb44
	case 3:
		goto sw_bb50
	case 4:
		goto sw_bb56
	case 5:
		goto sw_bb62
	case 6:
		goto sw_bb68
	case 7:
		goto sw_bb74
	case 8:
		goto sw_bb80
	case 9:
		goto sw_bb86
	case 10:
		goto sw_bb92
	case 11:
		goto sw_bb98
	case 12:
		goto sw_bb104
	case 13:
		goto sw_bb110
	case 14:
		goto sw_bb112
	case 15:
		goto sw_bb116
	default:
		goto sw_default
	}

sw_bb:
	v10 = *libc.As[int32](lookahead)
	cmp = v10 == 92
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end:
	v11 = *libc.As[int32](lookahead)
	cmp4 = v11 == 102
	if cmp4 {
		goto if_then6
	} else {
		goto if_end7
	}

if_then6:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end7:
	v12 = *libc.As[int32](lookahead)
	cmp8 = v12 == 110
	if cmp8 {
		goto if_then10
	} else {
		goto if_end11
	}

if_then10:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end11:
	v13 = *libc.As[int32](lookahead)
	cmp12 = v13 == 116
	if cmp12 {
		goto if_then14
	} else {
		goto if_end15
	}

if_then14:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end15:
	v14 = *libc.As[int32](lookahead)
	cmp16 = 9 <= v14
	if cmp16 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false
	}

land_lhs_true:
	v15 = *libc.As[int32](lookahead)
	cmp18 = v15 <= 13
	if cmp18 {
		goto if_then31
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v16 = *libc.As[int32](lookahead)
	cmp20 = v16 == 32
	if cmp20 {
		goto if_then31
	} else {
		goto lor_lhs_false22
	}

lor_lhs_false22:
	v17 = *libc.As[int32](lookahead)
	cmp23 = v17 == 8203
	if cmp23 {
		goto if_then31
	} else {
		goto lor_lhs_false25
	}

lor_lhs_false25:
	v18 = *libc.As[int32](lookahead)
	cmp26 = v18 == 8288
	if cmp26 {
		goto if_then31
	} else {
		goto lor_lhs_false28
	}

lor_lhs_false28:
	v19 = *libc.As[int32](lookahead)
	cmp29 = v19 == 65279
	if cmp29 {
		goto if_then31
	} else {
		goto if_end32
	}

if_then31:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end32:
	v20 = *libc.As[byte](result)
	loadedv33 = (v20 & 1) != 0
	*libc.As[bool](retval) = loadedv33
	goto _return

sw_bb34:
	v21 = *libc.As[int32](lookahead)
	cmp35 = v21 == 10
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end38:
	v22 = *libc.As[int32](lookahead)
	cmp39 = v22 == 13
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end42:
	v23 = *libc.As[byte](result)
	loadedv43 = (v23 & 1) != 0
	*libc.As[bool](retval) = loadedv43
	goto _return

sw_bb44:
	v24 = *libc.As[int32](lookahead)
	cmp45 = v24 == 97
	if cmp45 {
		goto if_then47
	} else {
		goto if_end48
	}

if_then47:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end48:
	v25 = *libc.As[byte](result)
	loadedv49 = (v25 & 1) != 0
	*libc.As[bool](retval) = loadedv49
	goto _return

sw_bb50:
	v26 = *libc.As[int32](lookahead)
	cmp51 = v26 == 117
	if cmp51 {
		goto if_then53
	} else {
		goto if_end54
	}

if_then53:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end54:
	v27 = *libc.As[byte](result)
	loadedv55 = (v27 & 1) != 0
	*libc.As[bool](retval) = loadedv55
	goto _return

sw_bb56:
	v28 = *libc.As[int32](lookahead)
	cmp57 = v28 == 114
	if cmp57 {
		goto if_then59
	} else {
		goto if_end60
	}

if_then59:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end60:
	v29 = *libc.As[byte](result)
	loadedv61 = (v29 & 1) != 0
	*libc.As[bool](retval) = loadedv61
	goto _return

sw_bb62:
	v30 = *libc.As[int32](lookahead)
	cmp63 = v30 == 10
	if cmp63 {
		goto if_then65
	} else {
		goto if_end66
	}

if_then65:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end66:
	v31 = *libc.As[byte](result)
	loadedv67 = (v31 & 1) != 0
	*libc.As[bool](retval) = loadedv67
	goto _return

sw_bb68:
	v32 = *libc.As[int32](lookahead)
	cmp69 = v32 == 108
	if cmp69 {
		goto if_then71
	} else {
		goto if_end72
	}

if_then71:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end72:
	v33 = *libc.As[byte](result)
	loadedv73 = (v33 & 1) != 0
	*libc.As[bool](retval) = loadedv73
	goto _return

sw_bb74:
	v34 = *libc.As[int32](lookahead)
	cmp75 = v34 == 108
	if cmp75 {
		goto if_then77
	} else {
		goto if_end78
	}

if_then77:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end78:
	v35 = *libc.As[byte](result)
	loadedv79 = (v35 & 1) != 0
	*libc.As[bool](retval) = loadedv79
	goto _return

sw_bb80:
	v36 = *libc.As[int32](lookahead)
	cmp81 = v36 == 117
	if cmp81 {
		goto if_then83
	} else {
		goto if_end84
	}

if_then83:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end84:
	v37 = *libc.As[byte](result)
	loadedv85 = (v37 & 1) != 0
	*libc.As[bool](retval) = loadedv85
	goto _return

sw_bb86:
	v38 = *libc.As[int32](lookahead)
	cmp87 = v38 == 115
	if cmp87 {
		goto if_then89
	} else {
		goto if_end90
	}

if_then89:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end90:
	v39 = *libc.As[byte](result)
	loadedv91 = (v39 & 1) != 0
	*libc.As[bool](retval) = loadedv91
	goto _return

sw_bb92:
	v40 = *libc.As[int32](lookahead)
	cmp93 = v40 == 108
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end96:
	v41 = *libc.As[byte](result)
	loadedv97 = (v41 & 1) != 0
	*libc.As[bool](retval) = loadedv97
	goto _return

sw_bb98:
	v42 = *libc.As[int32](lookahead)
	cmp99 = v42 == 101
	if cmp99 {
		goto if_then101
	} else {
		goto if_end102
	}

if_then101:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end102:
	v43 = *libc.As[byte](result)
	loadedv103 = (v43 & 1) != 0
	*libc.As[bool](retval) = loadedv103
	goto _return

sw_bb104:
	v44 = *libc.As[int32](lookahead)
	cmp105 = v44 == 101
	if cmp105 {
		goto if_then107
	} else {
		goto if_end108
	}

if_then107:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end108:
	v45 = *libc.As[byte](result)
	loadedv109 = (v45 & 1) != 0
	*libc.As[bool](retval) = loadedv109
	goto _return

sw_bb110:
	*libc.As[byte](result) = 1
	v46 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v46).F1)
	*libc.As[int16](result_symbol) = 5
	v47 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v47).F3)
	v48 = *libc.As[unsafe.Pointer](mark_end)
	v49 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v48)(v49)
	v50 = *libc.As[byte](result)
	loadedv111 = (v50 & 1) != 0
	*libc.As[bool](retval) = loadedv111
	goto _return

sw_bb112:
	*libc.As[byte](result) = 1
	v51 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol113 = libc.Ptr(&libc.As[TSLexer](v51).F1)
	*libc.As[int16](result_symbol113) = 3
	v52 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end114 = libc.Ptr(&libc.As[TSLexer](v52).F3)
	v53 = *libc.As[unsafe.Pointer](mark_end114)
	v54 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v53)(v54)
	v55 = *libc.As[byte](result)
	loadedv115 = (v55 & 1) != 0
	*libc.As[bool](retval) = loadedv115
	goto _return

sw_bb116:
	*libc.As[byte](result) = 1
	v56 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol117 = libc.Ptr(&libc.As[TSLexer](v56).F1)
	*libc.As[int16](result_symbol117) = 4
	v57 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end118 = libc.Ptr(&libc.As[TSLexer](v57).F3)
	v58 = *libc.As[unsafe.Pointer](mark_end118)
	v59 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v58)(v59)
	v60 = *libc.As[byte](result)
	loadedv119 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv119
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v61 = *libc.As[bool](retval)
	return v61
}
