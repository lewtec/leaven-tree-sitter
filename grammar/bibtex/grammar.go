package grammar_bibtex

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type TSCharacterRange struct {
	F0 int32
	F1 int32
}
type TSFieldMapEntry struct {
	F0 int16
	F1 byte
	F2 byte
}
type TSLanguage struct {
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
	F10 *int16
	F11 *int16
	F12 *int32
	F13 *TSParseActionEntry
	F14 **byte
	F15 **byte
	F16 *TSMapSlice
	F17 *TSFieldMapEntry
	F18 *TSSymbolMetadata
	F19 *int16
	F20 *int16
	F21 *int16
	F22 *TSLexerMode
	F23 func(*TSLexer, int16) bool
	F24 func(*TSLexer, int16) bool
	F25 int16
	F26 anon_2
	F27 *int16
	F28 *byte
	F29 *int16
	F30 int16
	F31 int32
	F32 *int16
	F33 *TSMapSlice
	F34 *int16
	F35 TSLanguageMetadata
}
type TSLanguageMetadata struct {
	F0 byte
	F1 byte
	F2 byte
}
type TSLexer struct {
	F0 int32
	F1 int16
	F2 func(*TSLexer, bool)
	F3 func(*TSLexer)
	F4 func(*TSLexer) int32
	F5 func(*TSLexer) bool
	F6 func(*TSLexer) bool
	F7 func(*TSLexer, *byte, ...interface{})
}
type TSLexerMode struct {
	F0 int16
	F1 int16
	F2 int16
}
type TSMapSlice struct {
	F0 int16
	F1 int16
}
type TSSymbolMetadata struct {
	F0 byte
	F1 byte
	F2 byte
}
type anon_0 struct {
	F0 byte
	F1 byte
	F2 int16
	F3 int16
	F4 int16
}
type anon_1 struct {
	F0 byte
	F1 byte
}
type anon_2 struct {
	F0 *byte
	F1 *int16
	F2 func() *byte
	F3 func(*byte)
	F4 func(*byte, *TSLexer, *byte) bool
	F5 func(*byte, *byte) int32
	F6 func(*byte, *byte, int32)
}
type TSParseAction struct {
	F0 anon_0
}
type TSParseActionEntry struct {
	F0 TSParseAction
}

var tree_sitter_bibtex_language TSLanguage = TSLanguage{15, 39, 0, 21, 0, 80, 2, 9, 5, 6, &(*[2][39]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[244]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], &ts_field_names[0], &ts_field_map_slices[0], &ts_field_map_entries[0], &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{0, 1, 0}}
var ts_small_parse_table [918]int16 = [918]int16{7, 15, 1, 0, 17, 1, 1, 20, 1, 2, 23, 1, 3, 26, 1, 9, 29, 1, 10, 2, 5, 22, 23, 24, 25, 32, 7, 9, 1, 3, 11, 1, 9, 13, 1, 10, 32, 1, 0, 34, 1, 1, 36, 1, 2, 2, 5, 22, 23, 24, 25, 32, 6, 38, 1, 4, 43, 1, 19, 46, 1, 20, 41, 2, 6, 16, 4, 2, 30, 36, 19, 2, 31, 38, 6, 49, 1, 4, 51, 1, 6, 53, 1, 19, 55, 1, 20, 15, 2, 30, 36, 19, 2, 31, 38, 6, 49, 1, 4, 53, 1, 19, 55, 1, 20, 57, 1, 6, 4, 2, 30, 36, 19, 2, 31, 38, 6, 59, 1, 4, 61, 1, 6, 63, 1, 18, 65, 1, 20, 11, 2, 29, 35, 21, 2, 31, 37, 6, 59, 1, 4, 63, 1, 18, 65, 1, 20, 67, 1, 6, 12, 2, 29, 35, 21, 2, 31, 37, 6, 59, 1, 4, 63, 1, 18, 65, 1, 20, 69, 1, 6, 13, 2, 29, 35, 21, 2, 31, 37, 6, 49, 1, 4, 53, 1, 19, 55, 1, 20, 67, 1, 16, 4, 2, 30, 36, 19, 2, 31, 38, 6, 59, 1, 4, 63, 1, 18, 65, 1, 20, 71, 1, 6, 12, 2, 29, 35, 21, 2, 31, 37, 6, 73, 1, 4, 76, 1, 6, 78, 1, 18, 81, 1, 20, 12, 2, 29, 35, 21, 2, 31, 37, 6, 59, 1, 4, 63, 1, 18, 65, 1, 20, 84, 1, 6, 12, 2, 29, 35, 21, 2, 31, 37, 6, 49, 1, 4, 53, 1, 19, 55, 1, 20, 86, 1, 16, 10, 2, 30, 36, 19, 2, 31, 38, 6, 49, 1, 4, 53, 1, 19, 55, 1, 20, 88, 1, 6, 4, 2, 30, 36, 19, 2, 31, 38, 6, 49, 1, 4, 53, 1, 19, 55, 1, 20, 90, 1, 6, 6, 2, 30, 36, 19, 2, 31, 38, 6, 59, 1, 4, 63, 1, 18, 65, 1, 20, 86, 1, 6, 8, 2, 29, 35, 21, 2, 31, 37, 4, 94, 1, 19, 97, 1, 20, 18, 2, 31, 38, 92, 3, 4, 6, 16, 4, 55, 1, 20, 102, 1, 19, 18, 2, 31, 38, 100, 3, 4, 6, 16, 2, 104, 2, 0, 1, 106, 4, 2, 3, 9, 10, 4, 65, 1, 20, 110, 1, 18, 108, 2, 4, 6, 24, 2, 31, 37, 5, 112, 1, 4, 116, 1, 16, 37, 1, 28, 79, 1, 27, 114, 2, 14, 17, 5, 112, 1, 4, 116, 1, 16, 37, 1, 28, 77, 1, 27, 114, 2, 14, 17, 4, 120, 1, 18, 123, 1, 20, 118, 2, 4, 6, 24, 2, 31, 37, 5, 112, 1, 4, 116, 1, 16, 37, 1, 28, 73, 1, 27, 114, 2, 14, 17, 5, 112, 1, 4, 116, 1, 16, 37, 1, 28, 78, 1, 27, 114, 2, 14, 17, 2, 126, 2, 0, 1, 128, 4, 2, 3, 9, 10, 2, 130, 2, 0, 1, 132, 4, 2, 3, 9, 10, 2, 130, 2, 0, 1, 132, 4, 2, 3, 9, 10, 2, 134, 2, 0, 1, 136, 4, 2, 3, 9, 10, 2, 138, 2, 0, 1, 140, 4, 2, 3, 9, 10, 2, 142, 2, 0, 1, 144, 4, 2, 3, 9, 10, 5, 112, 1, 4, 116, 1, 16, 37, 1, 28, 63, 1, 27, 114, 2, 14, 17, 2, 142, 2, 0, 1, 144, 4, 2, 3, 9, 10, 1, 146, 5, 4, 6, 16, 19, 20, 3, 150, 1, 15, 36, 1, 34, 148, 3, 6, 8, 11, 3, 155, 1, 15, 43, 1, 34, 153, 3, 6, 8, 11, 1, 157, 5, 4, 6, 16, 19, 20, 2, 159, 1, 4, 161, 4, 6, 16, 19, 20, 1, 163, 5, 4, 6, 16, 19, 20, 1, 165, 5, 4, 6, 16, 19, 20, 4, 112, 1, 4, 116, 1, 16, 49, 1, 28, 114, 2, 14, 17, 3, 155, 1, 15, 36, 1, 34, 167, 3, 6, 8, 11, 1, 169, 4, 4, 6, 18, 20, 1, 165, 4, 4, 6, 18, 20, 1, 171, 4, 6, 8, 11, 15, 1, 173, 4, 6, 8, 11, 15, 2, 175, 1, 4, 161, 3, 6, 18, 20, 1, 148, 4, 6, 8, 11, 15, 1, 177, 4, 6, 8, 11, 15, 1, 179, 4, 4, 6, 18, 20, 1, 163, 4, 4, 6, 18, 20, 3, 183, 1, 11, 53, 1, 33, 181, 2, 6, 8, 3, 186, 1, 8, 188, 1, 11, 61, 1, 33, 3, 190, 1, 6, 192, 1, 14, 58, 1, 26, 3, 192, 1, 14, 194, 1, 8, 58, 1, 26, 3, 186, 1, 6, 196, 1, 11, 59, 1, 33, 1, 198, 3, 6, 8, 11, 3, 200, 1, 6, 202, 1, 11, 53, 1, 33, 3, 192, 1, 14, 204, 1, 8, 58, 1, 26, 3, 206, 1, 8, 208, 1, 11, 53, 1, 33, 3, 192, 1, 14, 204, 1, 6, 58, 1, 26, 1, 210, 3, 6, 8, 11, 2, 212, 1, 4, 214, 1, 7, 2, 216, 1, 4, 218, 1, 7, 2, 220, 1, 4, 222, 1, 7, 2, 192, 1, 14, 58, 1, 26, 1, 224, 1, 5, 1, 226, 1, 5, 1, 228, 1, 0, 1, 230, 1, 5, 1, 232, 1, 13, 1, 234, 1, 6, 1, 236, 1, 14, 1, 238, 1, 12, 1, 240, 1, 14, 1, 242, 1, 6, 1, 242, 1, 8, 1, 234, 1, 8}
var ts_small_parse_table_map [78]int32 = [78]int32{0, 26, 52, 74, 95, 116, 137, 158, 179, 200, 221, 242, 263, 284, 305, 326, 347, 363, 379, 390, 405, 422, 439, 454, 471, 488, 499, 510, 521, 532, 543, 554, 571, 582, 590, 602, 614, 622, 632, 640, 648, 662, 674, 681, 688, 695, 702, 711, 718, 725, 732, 739, 750, 760, 770, 780, 790, 796, 806, 816, 826, 836, 842, 849, 856, 863, 870, 874, 878, 882, 886, 890, 894, 898, 902, 906, 910, 914}
var ts_symbol_names [39]*byte = [39]*byte{&_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0], &_str_17[0], &_str_18[0], &_str_19[0], &_str_20[0], &_str_21[0], &_str_22[0], &_str_23[0], &_str_24[0], &_str_25[0], &_str_26[0], &_str_27[0], &_str_28[0], &_str_29[0], &_str_30[0], &_str_31[0], &_str_32[0], &_str_33[0], &_str_34[0], &_str_35[0], &_str_36[0], &_str_37[0], &_str_38[0], &_str_39[0], &_str_40[0], &_str_41[0]}
var ts_field_names [6]*byte = [6]*byte{nil, &_str_29[0], &_str_42[0], &_str_43[0], &_str_44[0], &_str_30[0]}
var ts_field_map_slices [9]TSMapSlice = [9]TSMapSlice{TSMapSlice{}, TSMapSlice{0, 1}, TSMapSlice{1, 2}, TSMapSlice{3, 2}, TSMapSlice{5, 1}, TSMapSlice{6, 3}, TSMapSlice{9, 2}, TSMapSlice{11, 3}, TSMapSlice{14, 2}}
var ts_field_map_entries [16]TSFieldMapEntry = [16]TSFieldMapEntry{TSFieldMapEntry{3, 0, 0}, TSFieldMapEntry{4, 0, 0}, TSFieldMapEntry{5, 2, 0}, TSFieldMapEntry{2, 2, 0}, TSFieldMapEntry{4, 0, 0}, TSFieldMapEntry{1, 1, 0}, TSFieldMapEntry{1, 3, 1}, TSFieldMapEntry{2, 2, 0}, TSFieldMapEntry{4, 0, 0}, TSFieldMapEntry{1, 0, 1}, TSFieldMapEntry{1, 1, 1}, TSFieldMapEntry{3, 2, 0}, TSFieldMapEntry{4, 0, 0}, TSFieldMapEntry{5, 4, 0}, TSFieldMapEntry{3, 0, 0}, TSFieldMapEntry{5, 2, 0}}
var ts_symbol_metadata [39]TSSymbolMetadata = [39]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [39]int16 = [39]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [9][6]int16 = [9][6]int16{}
var ts_lex_modes [80]TSLexerMode = [80]TSLexerMode{TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{1, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{1, 0, 0}, TSLexerMode{}, TSLexerMode{1, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{1, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{37, 0, 0}, TSLexerMode{}, TSLexerMode{1, 0, 0}, TSLexerMode{36, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}}
var ts_primary_state_ids [80]int16 = [80]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 6, 5, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 41, 46, 47, 39, 49, 50, 51, 40, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79}
var _str [7]byte = [7]byte{98, 105, 98, 116, 101, 120, 0}
var ts_parse_table struct {
	F0 struct {
		F0 [21]int16
		F1 [18]int16
	}
	F1 [39]int16
} = struct {
	F0 struct {
		F0 [21]int16
		F1 [18]int16
	}
	F1 [39]int16
}{struct {
	F0 [21]int16
	F1 [18]int16
}{[21]int16{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1}, [18]int16{}}, [39]int16{3, 5, 7, 9, 0, 0, 0, 0, 0, 11, 13, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 70, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}}
var ts_parse_actions struct {
	F0 struct {
		F0 anon_1
		F1 [6]byte
	}
	F1 struct {
		F0 anon_1
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_1
		F1 [6]byte
	}
	F4 TSParseActionEntry
	F5 struct {
		F0 anon_1
		F1 [6]byte
	}
	F6 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F7 struct {
		F0 anon_1
		F1 [6]byte
	}
	F8 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F9 struct {
		F0 anon_1
		F1 [6]byte
	}
	F10 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F11 struct {
		F0 anon_1
		F1 [6]byte
	}
	F12 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F13 struct {
		F0 anon_1
		F1 [6]byte
	}
	F14 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F15 struct {
		F0 anon_1
		F1 [6]byte
	}
	F16 TSParseActionEntry
	F17 struct {
		F0 anon_1
		F1 [6]byte
	}
	F18 TSParseActionEntry
	F19 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F20 struct {
		F0 anon_1
		F1 [6]byte
	}
	F21 TSParseActionEntry
	F22 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F23 struct {
		F0 anon_1
		F1 [6]byte
	}
	F24 TSParseActionEntry
	F25 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F26 struct {
		F0 anon_1
		F1 [6]byte
	}
	F27 TSParseActionEntry
	F28 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F29 struct {
		F0 anon_1
		F1 [6]byte
	}
	F30 TSParseActionEntry
	F31 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F32 struct {
		F0 anon_1
		F1 [6]byte
	}
	F33 TSParseActionEntry
	F34 struct {
		F0 anon_1
		F1 [6]byte
	}
	F35 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F36 struct {
		F0 anon_1
		F1 [6]byte
	}
	F37 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F38 struct {
		F0 anon_1
		F1 [6]byte
	}
	F39 TSParseActionEntry
	F40 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F41 struct {
		F0 anon_1
		F1 [6]byte
	}
	F42 TSParseActionEntry
	F43 struct {
		F0 anon_1
		F1 [6]byte
	}
	F44 TSParseActionEntry
	F45 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F46 struct {
		F0 anon_1
		F1 [6]byte
	}
	F47 TSParseActionEntry
	F48 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F49 struct {
		F0 anon_1
		F1 [6]byte
	}
	F50 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F51 struct {
		F0 anon_1
		F1 [6]byte
	}
	F52 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F53 struct {
		F0 anon_1
		F1 [6]byte
	}
	F54 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F55 struct {
		F0 anon_1
		F1 [6]byte
	}
	F56 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F57 struct {
		F0 anon_1
		F1 [6]byte
	}
	F58 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F59 struct {
		F0 anon_1
		F1 [6]byte
	}
	F60 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F61 struct {
		F0 anon_1
		F1 [6]byte
	}
	F62 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F63 struct {
		F0 anon_1
		F1 [6]byte
	}
	F64 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F65 struct {
		F0 anon_1
		F1 [6]byte
	}
	F66 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F67 struct {
		F0 anon_1
		F1 [6]byte
	}
	F68 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F69 struct {
		F0 anon_1
		F1 [6]byte
	}
	F70 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F71 struct {
		F0 anon_1
		F1 [6]byte
	}
	F72 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F73 struct {
		F0 anon_1
		F1 [6]byte
	}
	F74 TSParseActionEntry
	F75 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F76 struct {
		F0 anon_1
		F1 [6]byte
	}
	F77 TSParseActionEntry
	F78 struct {
		F0 anon_1
		F1 [6]byte
	}
	F79 TSParseActionEntry
	F80 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F81 struct {
		F0 anon_1
		F1 [6]byte
	}
	F82 TSParseActionEntry
	F83 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F84 struct {
		F0 anon_1
		F1 [6]byte
	}
	F85 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F86 struct {
		F0 anon_1
		F1 [6]byte
	}
	F87 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F88 struct {
		F0 anon_1
		F1 [6]byte
	}
	F89 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F90 struct {
		F0 anon_1
		F1 [6]byte
	}
	F91 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F92 struct {
		F0 anon_1
		F1 [6]byte
	}
	F93 TSParseActionEntry
	F94 struct {
		F0 anon_1
		F1 [6]byte
	}
	F95 TSParseActionEntry
	F96 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F97 struct {
		F0 anon_1
		F1 [6]byte
	}
	F98 TSParseActionEntry
	F99 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F100 struct {
		F0 anon_1
		F1 [6]byte
	}
	F101 TSParseActionEntry
	F102 struct {
		F0 anon_1
		F1 [6]byte
	}
	F103 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F104 struct {
		F0 anon_1
		F1 [6]byte
	}
	F105 TSParseActionEntry
	F106 struct {
		F0 anon_1
		F1 [6]byte
	}
	F107 TSParseActionEntry
	F108 struct {
		F0 anon_1
		F1 [6]byte
	}
	F109 TSParseActionEntry
	F110 struct {
		F0 anon_1
		F1 [6]byte
	}
	F111 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F112 struct {
		F0 anon_1
		F1 [6]byte
	}
	F113 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F114 struct {
		F0 anon_1
		F1 [6]byte
	}
	F115 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F116 struct {
		F0 anon_1
		F1 [6]byte
	}
	F117 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F118 struct {
		F0 anon_1
		F1 [6]byte
	}
	F119 TSParseActionEntry
	F120 struct {
		F0 anon_1
		F1 [6]byte
	}
	F121 TSParseActionEntry
	F122 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F123 struct {
		F0 anon_1
		F1 [6]byte
	}
	F124 TSParseActionEntry
	F125 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F126 struct {
		F0 anon_1
		F1 [6]byte
	}
	F127 TSParseActionEntry
	F128 struct {
		F0 anon_1
		F1 [6]byte
	}
	F129 TSParseActionEntry
	F130 struct {
		F0 anon_1
		F1 [6]byte
	}
	F131 TSParseActionEntry
	F132 struct {
		F0 anon_1
		F1 [6]byte
	}
	F133 TSParseActionEntry
	F134 struct {
		F0 anon_1
		F1 [6]byte
	}
	F135 TSParseActionEntry
	F136 struct {
		F0 anon_1
		F1 [6]byte
	}
	F137 TSParseActionEntry
	F138 struct {
		F0 anon_1
		F1 [6]byte
	}
	F139 TSParseActionEntry
	F140 struct {
		F0 anon_1
		F1 [6]byte
	}
	F141 TSParseActionEntry
	F142 struct {
		F0 anon_1
		F1 [6]byte
	}
	F143 TSParseActionEntry
	F144 struct {
		F0 anon_1
		F1 [6]byte
	}
	F145 TSParseActionEntry
	F146 struct {
		F0 anon_1
		F1 [6]byte
	}
	F147 TSParseActionEntry
	F148 struct {
		F0 anon_1
		F1 [6]byte
	}
	F149 TSParseActionEntry
	F150 struct {
		F0 anon_1
		F1 [6]byte
	}
	F151 TSParseActionEntry
	F152 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F153 struct {
		F0 anon_1
		F1 [6]byte
	}
	F154 TSParseActionEntry
	F155 struct {
		F0 anon_1
		F1 [6]byte
	}
	F156 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F157 struct {
		F0 anon_1
		F1 [6]byte
	}
	F158 TSParseActionEntry
	F159 struct {
		F0 anon_1
		F1 [6]byte
	}
	F160 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F161 struct {
		F0 anon_1
		F1 [6]byte
	}
	F162 TSParseActionEntry
	F163 struct {
		F0 anon_1
		F1 [6]byte
	}
	F164 TSParseActionEntry
	F165 struct {
		F0 anon_1
		F1 [6]byte
	}
	F166 TSParseActionEntry
	F167 struct {
		F0 anon_1
		F1 [6]byte
	}
	F168 TSParseActionEntry
	F169 struct {
		F0 anon_1
		F1 [6]byte
	}
	F170 TSParseActionEntry
	F171 struct {
		F0 anon_1
		F1 [6]byte
	}
	F172 TSParseActionEntry
	F173 struct {
		F0 anon_1
		F1 [6]byte
	}
	F174 TSParseActionEntry
	F175 struct {
		F0 anon_1
		F1 [6]byte
	}
	F176 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F177 struct {
		F0 anon_1
		F1 [6]byte
	}
	F178 TSParseActionEntry
	F179 struct {
		F0 anon_1
		F1 [6]byte
	}
	F180 TSParseActionEntry
	F181 struct {
		F0 anon_1
		F1 [6]byte
	}
	F182 TSParseActionEntry
	F183 struct {
		F0 anon_1
		F1 [6]byte
	}
	F184 TSParseActionEntry
	F185 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F186 struct {
		F0 anon_1
		F1 [6]byte
	}
	F187 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F188 struct {
		F0 anon_1
		F1 [6]byte
	}
	F189 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F190 struct {
		F0 anon_1
		F1 [6]byte
	}
	F191 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F192 struct {
		F0 anon_1
		F1 [6]byte
	}
	F193 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F194 struct {
		F0 anon_1
		F1 [6]byte
	}
	F195 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F196 struct {
		F0 anon_1
		F1 [6]byte
	}
	F197 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F198 struct {
		F0 anon_1
		F1 [6]byte
	}
	F199 TSParseActionEntry
	F200 struct {
		F0 anon_1
		F1 [6]byte
	}
	F201 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F202 struct {
		F0 anon_1
		F1 [6]byte
	}
	F203 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F204 struct {
		F0 anon_1
		F1 [6]byte
	}
	F205 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F206 struct {
		F0 anon_1
		F1 [6]byte
	}
	F207 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F208 struct {
		F0 anon_1
		F1 [6]byte
	}
	F209 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F210 struct {
		F0 anon_1
		F1 [6]byte
	}
	F211 TSParseActionEntry
	F212 struct {
		F0 anon_1
		F1 [6]byte
	}
	F213 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F214 struct {
		F0 anon_1
		F1 [6]byte
	}
	F215 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F216 struct {
		F0 anon_1
		F1 [6]byte
	}
	F217 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F218 struct {
		F0 anon_1
		F1 [6]byte
	}
	F219 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F220 struct {
		F0 anon_1
		F1 [6]byte
	}
	F221 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F222 struct {
		F0 anon_1
		F1 [6]byte
	}
	F223 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F224 struct {
		F0 anon_1
		F1 [6]byte
	}
	F225 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F226 struct {
		F0 anon_1
		F1 [6]byte
	}
	F227 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F228 struct {
		F0 anon_1
		F1 [6]byte
	}
	F229 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F230 struct {
		F0 anon_1
		F1 [6]byte
	}
	F231 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F232 struct {
		F0 anon_1
		F1 [6]byte
	}
	F233 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F234 struct {
		F0 anon_1
		F1 [6]byte
	}
	F235 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F236 struct {
		F0 anon_1
		F1 [6]byte
	}
	F237 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F238 struct {
		F0 anon_1
		F1 [6]byte
	}
	F239 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F240 struct {
		F0 anon_1
		F1 [6]byte
	}
	F241 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F242 struct {
		F0 anon_1
		F1 [6]byte
	}
	F243 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
} = struct {
	F0 struct {
		F0 anon_1
		F1 [6]byte
	}
	F1 struct {
		F0 anon_1
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_1
		F1 [6]byte
	}
	F4 TSParseActionEntry
	F5 struct {
		F0 anon_1
		F1 [6]byte
	}
	F6 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F7 struct {
		F0 anon_1
		F1 [6]byte
	}
	F8 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F9 struct {
		F0 anon_1
		F1 [6]byte
	}
	F10 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F11 struct {
		F0 anon_1
		F1 [6]byte
	}
	F12 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F13 struct {
		F0 anon_1
		F1 [6]byte
	}
	F14 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F15 struct {
		F0 anon_1
		F1 [6]byte
	}
	F16 TSParseActionEntry
	F17 struct {
		F0 anon_1
		F1 [6]byte
	}
	F18 TSParseActionEntry
	F19 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F20 struct {
		F0 anon_1
		F1 [6]byte
	}
	F21 TSParseActionEntry
	F22 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F23 struct {
		F0 anon_1
		F1 [6]byte
	}
	F24 TSParseActionEntry
	F25 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F26 struct {
		F0 anon_1
		F1 [6]byte
	}
	F27 TSParseActionEntry
	F28 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F29 struct {
		F0 anon_1
		F1 [6]byte
	}
	F30 TSParseActionEntry
	F31 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F32 struct {
		F0 anon_1
		F1 [6]byte
	}
	F33 TSParseActionEntry
	F34 struct {
		F0 anon_1
		F1 [6]byte
	}
	F35 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F36 struct {
		F0 anon_1
		F1 [6]byte
	}
	F37 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F38 struct {
		F0 anon_1
		F1 [6]byte
	}
	F39 TSParseActionEntry
	F40 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F41 struct {
		F0 anon_1
		F1 [6]byte
	}
	F42 TSParseActionEntry
	F43 struct {
		F0 anon_1
		F1 [6]byte
	}
	F44 TSParseActionEntry
	F45 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F46 struct {
		F0 anon_1
		F1 [6]byte
	}
	F47 TSParseActionEntry
	F48 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F49 struct {
		F0 anon_1
		F1 [6]byte
	}
	F50 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F51 struct {
		F0 anon_1
		F1 [6]byte
	}
	F52 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F53 struct {
		F0 anon_1
		F1 [6]byte
	}
	F54 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F55 struct {
		F0 anon_1
		F1 [6]byte
	}
	F56 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F57 struct {
		F0 anon_1
		F1 [6]byte
	}
	F58 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F59 struct {
		F0 anon_1
		F1 [6]byte
	}
	F60 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F61 struct {
		F0 anon_1
		F1 [6]byte
	}
	F62 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F63 struct {
		F0 anon_1
		F1 [6]byte
	}
	F64 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F65 struct {
		F0 anon_1
		F1 [6]byte
	}
	F66 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F67 struct {
		F0 anon_1
		F1 [6]byte
	}
	F68 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F69 struct {
		F0 anon_1
		F1 [6]byte
	}
	F70 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F71 struct {
		F0 anon_1
		F1 [6]byte
	}
	F72 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F73 struct {
		F0 anon_1
		F1 [6]byte
	}
	F74 TSParseActionEntry
	F75 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F76 struct {
		F0 anon_1
		F1 [6]byte
	}
	F77 TSParseActionEntry
	F78 struct {
		F0 anon_1
		F1 [6]byte
	}
	F79 TSParseActionEntry
	F80 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F81 struct {
		F0 anon_1
		F1 [6]byte
	}
	F82 TSParseActionEntry
	F83 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F84 struct {
		F0 anon_1
		F1 [6]byte
	}
	F85 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F86 struct {
		F0 anon_1
		F1 [6]byte
	}
	F87 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F88 struct {
		F0 anon_1
		F1 [6]byte
	}
	F89 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F90 struct {
		F0 anon_1
		F1 [6]byte
	}
	F91 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F92 struct {
		F0 anon_1
		F1 [6]byte
	}
	F93 TSParseActionEntry
	F94 struct {
		F0 anon_1
		F1 [6]byte
	}
	F95 TSParseActionEntry
	F96 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F97 struct {
		F0 anon_1
		F1 [6]byte
	}
	F98 TSParseActionEntry
	F99 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F100 struct {
		F0 anon_1
		F1 [6]byte
	}
	F101 TSParseActionEntry
	F102 struct {
		F0 anon_1
		F1 [6]byte
	}
	F103 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F104 struct {
		F0 anon_1
		F1 [6]byte
	}
	F105 TSParseActionEntry
	F106 struct {
		F0 anon_1
		F1 [6]byte
	}
	F107 TSParseActionEntry
	F108 struct {
		F0 anon_1
		F1 [6]byte
	}
	F109 TSParseActionEntry
	F110 struct {
		F0 anon_1
		F1 [6]byte
	}
	F111 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F112 struct {
		F0 anon_1
		F1 [6]byte
	}
	F113 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F114 struct {
		F0 anon_1
		F1 [6]byte
	}
	F115 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F116 struct {
		F0 anon_1
		F1 [6]byte
	}
	F117 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F118 struct {
		F0 anon_1
		F1 [6]byte
	}
	F119 TSParseActionEntry
	F120 struct {
		F0 anon_1
		F1 [6]byte
	}
	F121 TSParseActionEntry
	F122 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F123 struct {
		F0 anon_1
		F1 [6]byte
	}
	F124 TSParseActionEntry
	F125 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F126 struct {
		F0 anon_1
		F1 [6]byte
	}
	F127 TSParseActionEntry
	F128 struct {
		F0 anon_1
		F1 [6]byte
	}
	F129 TSParseActionEntry
	F130 struct {
		F0 anon_1
		F1 [6]byte
	}
	F131 TSParseActionEntry
	F132 struct {
		F0 anon_1
		F1 [6]byte
	}
	F133 TSParseActionEntry
	F134 struct {
		F0 anon_1
		F1 [6]byte
	}
	F135 TSParseActionEntry
	F136 struct {
		F0 anon_1
		F1 [6]byte
	}
	F137 TSParseActionEntry
	F138 struct {
		F0 anon_1
		F1 [6]byte
	}
	F139 TSParseActionEntry
	F140 struct {
		F0 anon_1
		F1 [6]byte
	}
	F141 TSParseActionEntry
	F142 struct {
		F0 anon_1
		F1 [6]byte
	}
	F143 TSParseActionEntry
	F144 struct {
		F0 anon_1
		F1 [6]byte
	}
	F145 TSParseActionEntry
	F146 struct {
		F0 anon_1
		F1 [6]byte
	}
	F147 TSParseActionEntry
	F148 struct {
		F0 anon_1
		F1 [6]byte
	}
	F149 TSParseActionEntry
	F150 struct {
		F0 anon_1
		F1 [6]byte
	}
	F151 TSParseActionEntry
	F152 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F153 struct {
		F0 anon_1
		F1 [6]byte
	}
	F154 TSParseActionEntry
	F155 struct {
		F0 anon_1
		F1 [6]byte
	}
	F156 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F157 struct {
		F0 anon_1
		F1 [6]byte
	}
	F158 TSParseActionEntry
	F159 struct {
		F0 anon_1
		F1 [6]byte
	}
	F160 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F161 struct {
		F0 anon_1
		F1 [6]byte
	}
	F162 TSParseActionEntry
	F163 struct {
		F0 anon_1
		F1 [6]byte
	}
	F164 TSParseActionEntry
	F165 struct {
		F0 anon_1
		F1 [6]byte
	}
	F166 TSParseActionEntry
	F167 struct {
		F0 anon_1
		F1 [6]byte
	}
	F168 TSParseActionEntry
	F169 struct {
		F0 anon_1
		F1 [6]byte
	}
	F170 TSParseActionEntry
	F171 struct {
		F0 anon_1
		F1 [6]byte
	}
	F172 TSParseActionEntry
	F173 struct {
		F0 anon_1
		F1 [6]byte
	}
	F174 TSParseActionEntry
	F175 struct {
		F0 anon_1
		F1 [6]byte
	}
	F176 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F177 struct {
		F0 anon_1
		F1 [6]byte
	}
	F178 TSParseActionEntry
	F179 struct {
		F0 anon_1
		F1 [6]byte
	}
	F180 TSParseActionEntry
	F181 struct {
		F0 anon_1
		F1 [6]byte
	}
	F182 TSParseActionEntry
	F183 struct {
		F0 anon_1
		F1 [6]byte
	}
	F184 TSParseActionEntry
	F185 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F186 struct {
		F0 anon_1
		F1 [6]byte
	}
	F187 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F188 struct {
		F0 anon_1
		F1 [6]byte
	}
	F189 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F190 struct {
		F0 anon_1
		F1 [6]byte
	}
	F191 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F192 struct {
		F0 anon_1
		F1 [6]byte
	}
	F193 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F194 struct {
		F0 anon_1
		F1 [6]byte
	}
	F195 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F196 struct {
		F0 anon_1
		F1 [6]byte
	}
	F197 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F198 struct {
		F0 anon_1
		F1 [6]byte
	}
	F199 TSParseActionEntry
	F200 struct {
		F0 anon_1
		F1 [6]byte
	}
	F201 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F202 struct {
		F0 anon_1
		F1 [6]byte
	}
	F203 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F204 struct {
		F0 anon_1
		F1 [6]byte
	}
	F205 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F206 struct {
		F0 anon_1
		F1 [6]byte
	}
	F207 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F208 struct {
		F0 anon_1
		F1 [6]byte
	}
	F209 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F210 struct {
		F0 anon_1
		F1 [6]byte
	}
	F211 TSParseActionEntry
	F212 struct {
		F0 anon_1
		F1 [6]byte
	}
	F213 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F214 struct {
		F0 anon_1
		F1 [6]byte
	}
	F215 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F216 struct {
		F0 anon_1
		F1 [6]byte
	}
	F217 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F218 struct {
		F0 anon_1
		F1 [6]byte
	}
	F219 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F220 struct {
		F0 anon_1
		F1 [6]byte
	}
	F221 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F222 struct {
		F0 anon_1
		F1 [6]byte
	}
	F223 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F224 struct {
		F0 anon_1
		F1 [6]byte
	}
	F225 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F226 struct {
		F0 anon_1
		F1 [6]byte
	}
	F227 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F228 struct {
		F0 anon_1
		F1 [6]byte
	}
	F229 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F230 struct {
		F0 anon_1
		F1 [6]byte
	}
	F231 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F232 struct {
		F0 anon_1
		F1 [6]byte
	}
	F233 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F234 struct {
		F0 anon_1
		F1 [6]byte
	}
	F235 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F236 struct {
		F0 anon_1
		F1 [6]byte
	}
	F237 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F238 struct {
		F0 anon_1
		F1 [6]byte
	}
	F239 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F240 struct {
		F0 anon_1
		F1 [6]byte
	}
	F241 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
	F242 struct {
		F0 anon_1
		F1 [6]byte
	}
	F243 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 int16
				F2 byte
				F3 byte
			}
			F1 [2]byte
		}
	}
}{struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{}, [6]byte{}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 byte
		F1 [7]byte
	}
}{struct {
	F0 byte
	F1 [7]byte
}{3, [7]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 21, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 3, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 3, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 66, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 65, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 64, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 32, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 32, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 2, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 32, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 2, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 32, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 66, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 32, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 65, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 32, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 64, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 36, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 9, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 36, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 36, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 19, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 36, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 39, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 9, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 52, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 19, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 39, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 41, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 7, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 44, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 21, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 48, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 46, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 35, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 51, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 35, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 7, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 35, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 35, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 21, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 35, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 48, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 38, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 47, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 45, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 40, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 38, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 38, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 18, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 38, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 39, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 30, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 18, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 25, 0, 3}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 25, 0, 3}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 29, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 24, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 17, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 50, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 14, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 37, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 37, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 24, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 37, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 48, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 25, 0, 3}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 25, 0, 3}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 25, 0, 5}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 25, 0, 5}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 23, 0, 7}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 23, 0, 7}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 24, 0, 2}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 24, 0, 2}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 25, 0, 5}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 25, 0, 5}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 30, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 34, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 34, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 42, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 27, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 42, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 30, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 16, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 31, 0, 1}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 31, 0, 1}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 31, 0, 1}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 27, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 29, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 28, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 28, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 5, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 28, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 29, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 33, 0, 6}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 33, 0, 6}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 67, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 27, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 60, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 32, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 68, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 34, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 62, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 33, 0, 4}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 28, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 55, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 20, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 29, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 56, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 26, 0, 8}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 75, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 72, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 25, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 22, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 74, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 76, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 33, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 23, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 byte
		F1 [7]byte
	}
}{struct {
	F0 byte
	F1 [7]byte
}{2, [7]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 26, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 54, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 31, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 69, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 57, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 71, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 int16
			F2 byte
			F3 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 int16
		F2 byte
		F3 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}{0, 30, 0, 0}, [2]byte{}}}}
var _str_3 [4]byte = [4]byte{101, 110, 100, 0}
var _str_4 [5]byte = [5]byte{106, 117, 110, 107, 0}
var _str_5 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_6 [12]byte = [12]byte{115, 116, 114, 105, 110, 103, 95, 116, 121, 112, 101, 0}
var _str_7 [2]byte = [2]byte{123, 0}
var _str_8 [2]byte = [2]byte{61, 0}
var _str_9 [2]byte = [2]byte{125, 0}
var _str_10 [2]byte = [2]byte{40, 0}
var _str_11 [2]byte = [2]byte{41, 0}
var _str_12 [14]byte = [14]byte{112, 114, 101, 97, 109, 98, 108, 101, 95, 116, 121, 112, 101, 0}
var _str_13 [11]byte = [11]byte{101, 110, 116, 114, 121, 95, 116, 121, 112, 101, 0}
var _str_14 [2]byte = [2]byte{44, 0}
var _str_15 [10]byte = [10]byte{107, 101, 121, 95, 98, 114, 97, 99, 101, 0}
var _str_16 [10]byte = [10]byte{107, 101, 121, 95, 112, 97, 114, 101, 110, 0}
var _str_17 [11]byte = [11]byte{105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}
var _str_18 [2]byte = [2]byte{35, 0}
var _str_19 [2]byte = [2]byte{34, 0}
var _str_20 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}
var _str_21 [11]byte = [11]byte{98, 114, 97, 99, 101, 95, 119, 111, 114, 100, 0}
var _str_22 [11]byte = [11]byte{113, 117, 111, 116, 101, 95, 119, 111, 114, 100, 0}
var _str_23 [13]byte = [13]byte{99, 111, 109, 109, 97, 110, 100, 95, 110, 97, 109, 101, 0}
var _str_24 [9]byte = [9]byte{100, 111, 99, 117, 109, 101, 110, 116, 0}
var _str_25 [18]byte = [18]byte{95, 99, 111, 109, 109, 97, 110, 100, 95, 111, 114, 95, 101, 110, 116, 114, 121, 0}
var _str_26 [7]byte = [7]byte{115, 116, 114, 105, 110, 103, 0}
var _str_27 [9]byte = [9]byte{112, 114, 101, 97, 109, 98, 108, 101, 0}
var _str_28 [6]byte = [6]byte{101, 110, 116, 114, 121, 0}
var _str_29 [6]byte = [6]byte{102, 105, 101, 108, 100, 0}
var _str_30 [6]byte = [6]byte{118, 97, 108, 117, 101, 0}
var _str_31 [6]byte = [6]byte{116, 111, 107, 101, 110, 0}
var _str_32 [16]byte = [16]byte{95, 98, 114, 97, 99, 101, 95, 98, 97, 108, 97, 110, 99, 101, 100, 0}
var _str_33 [16]byte = [16]byte{95, 113, 117, 111, 116, 101, 95, 98, 97, 108, 97, 110, 99, 101, 100, 0}
var _str_34 [8]byte = [8]byte{99, 111, 109, 109, 97, 110, 100, 0}
var _str_35 [17]byte = [17]byte{100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_36 [14]byte = [14]byte{101, 110, 116, 114, 121, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_37 [14]byte = [14]byte{118, 97, 108, 117, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_38 [14]byte = [14]byte{116, 111, 107, 101, 110, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_39 [14]byte = [14]byte{116, 111, 107, 101, 110, 95, 114, 101, 112, 101, 97, 116, 50, 0}
var _str_40 [24]byte = [24]byte{95, 98, 114, 97, 99, 101, 95, 98, 97, 108, 97, 110, 99, 101, 100, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_41 [24]byte = [24]byte{95, 113, 117, 111, 116, 101, 95, 98, 97, 108, 97, 110, 99, 101, 100, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_42 [4]byte = [4]byte{107, 101, 121, 0}
var _str_43 [5]byte = [5]byte{110, 97, 109, 101, 0}
var _str_44 [3]byte = [3]byte{116, 121, 0}
var ts_lex_map [20]int16 = [20]int16{34, 40, 35, 39, 40, 13, 41, 14, 44, 35, 61, 11, 64, 4, 92, 46, 123, 10, 125, 12}
var sym_entry_type_character_set_1 [9]TSCharacterRange = [9]TSCharacterRange{TSCharacterRange{33, 33}, TSCharacterRange{36, 36}, TSCharacterRange{38, 38}, TSCharacterRange{42, 43}, TSCharacterRange{45, 47}, TSCharacterRange{58, 60}, TSCharacterRange{62, 122}, TSCharacterRange{124, 124}, TSCharacterRange{126, 126}}

func tree_sitter_bibtex() *TSLanguage {
	return &tree_sitter_bibtex_language
}
func ts_lex(lexer *TSLexer, state int16) bool {
	var v0, v2, v4, v6, v8, v75, v76, v78, v80, v81, v83, v87, v88, v90, v95, v96, v98, v103, v104, v106, v108, v109, v111, v113, v114, v116, v118, v119, v121, v123, v124, v126, v128, v129, v131, v136, v137, v139, v146, v147, v149, v156, v157, v159, v166, v167, v169, v176, v177, v179, v186, v187, v189, v196, v197, v199, v206, v207, v209, v216, v217, v219, v226, v227, v229, v236, v237, v239, v246, v247, v249, v256, v257, v259, v266, v267, v269, v276, v277, v279, v286, v287, v289, v296, v297, v299, v306, v307, v309, v316, v317, v319, v324, v325, v327, v329, v330, v332, v340, v341, v343, v350, v351, v353, v358, v359, v361, v363, v364, v366, v368, v369, v371, v375, v376, v378, v387, v388, v390, v400, v401, v403, v405, v406, v408, v415, v416, v418 *TSLexer
	var lexer_addr **TSLexer
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, tobool29, cmp31, cmp35, cmp39, cmp43, cmp47, cmp50, cmp53, cmp57, cmp60, cmp64, cmp67, cmp70, cmp73, cmp76, cmp79, cmp82, cmp85, cmp88, tobool92, cmp94, cmp98, cmp102, cmp106, cmp110, cmp113, cmp116, cmp120, tobool124, cmp126, cmp130, cmp134, cmp138, cmp141, cmp144, cmp148, tobool152, cmp154, cmp157, cmp161, cmp164, cmp168, cmp171, call175, tobool178, tobool180, cmp183, cmp187, cmp190, cmp193, cmp197, tobool201, tobool203, cmp207, cmp210, tobool214, call218, cmp221, cmp224, tobool228, call232, cmp235, cmp238, tobool242, tobool246, tobool250, tobool254, tobool258, tobool262, call266, cmp269, cmp272, tobool276, cmp280, cmp283, call287, cmp290, cmp293, tobool297, cmp301, cmp304, call308, cmp311, cmp314, tobool318, cmp322, cmp325, call329, cmp332, cmp335, tobool339, cmp343, cmp346, call350, cmp353, cmp356, tobool360, cmp364, cmp367, call371, cmp374, cmp377, tobool381, cmp385, cmp388, call392, cmp395, cmp398, tobool402, cmp406, cmp409, call413, cmp416, cmp419, tobool423, cmp427, cmp430, call434, cmp437, cmp440, tobool444, cmp448, cmp451, call455, cmp458, cmp461, tobool465, cmp469, cmp472, call476, cmp479, cmp482, tobool486, cmp490, cmp493, call497, cmp500, cmp503, tobool507, cmp511, cmp514, call518, cmp521, cmp524, tobool528, cmp532, cmp535, call539, cmp542, cmp545, tobool549, cmp553, cmp556, call560, cmp563, cmp566, tobool570, cmp574, cmp577, call581, cmp584, cmp587, tobool591, cmp595, cmp598, call602, cmp605, cmp608, tobool612, cmp616, cmp619, call623, cmp626, cmp629, tobool633, cmp637, cmp640, call644, cmp647, cmp650, tobool654, call658, cmp661, cmp664, tobool668, tobool672, cmp676, cmp679, cmp682, cmp685, cmp688, cmp691, tobool695, cmp699, cmp702, cmp705, cmp708, cmp711, tobool715, call719, cmp722, cmp725, tobool729, tobool733, tobool737, cmp741, cmp744, tobool748, cmp752, cmp755, cmp758, cmp761, cmp764, cmp767, cmp770, tobool774, cmp778, cmp781, cmp784, cmp787, cmp790, cmp793, cmp796, cmp799, tobool803, tobool807, cmp811, cmp815, cmp818, cmp821, cmp824, tobool828, cmp832, cmp835, cmp838, cmp841, cmp845, cmp848, cmp851, tobool855, v427 bool
	var v7 func(*TSLexer) bool
	var eof2 *func(*TSLexer) bool
	var retval *bool
	var v9, v13, v16 int16
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol205, result_symbol216, result_symbol230, result_symbol244, result_symbol248, result_symbol252, result_symbol256, result_symbol260, result_symbol264, result_symbol278, result_symbol299, result_symbol320, result_symbol341, result_symbol362, result_symbol383, result_symbol404, result_symbol425, result_symbol446, result_symbol467, result_symbol488, result_symbol509, result_symbol530, result_symbol551, result_symbol572, result_symbol593, result_symbol614, result_symbol635, result_symbol656, result_symbol670, result_symbol674, result_symbol697, result_symbol717, result_symbol731, result_symbol735, result_symbol739, result_symbol750, result_symbol776, result_symbol805, result_symbol809, result_symbol830 *int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v43, v44, v45, v46, v47, v48, v49, v50, v52, v53, v54, v55, v56, v57, v58, v60, v61, v62, v63, v64, v65, v66, v69, v70, v71, v72, v73, v84, v85, v91, v92, v93, v99, v100, v101, v132, v133, v134, v140, v141, v142, v143, v144, v150, v151, v152, v153, v154, v160, v161, v162, v163, v164, v170, v171, v172, v173, v174, v180, v181, v182, v183, v184, v190, v191, v192, v193, v194, v200, v201, v202, v203, v204, v210, v211, v212, v213, v214, v220, v221, v222, v223, v224, v230, v231, v232, v233, v234, v240, v241, v242, v243, v244, v250, v251, v252, v253, v254, v260, v261, v262, v263, v264, v270, v271, v272, v273, v274, v280, v281, v282, v283, v284, v290, v291, v292, v293, v294, v300, v301, v302, v303, v304, v310, v311, v312, v313, v314, v320, v321, v322, v333, v334, v335, v336, v337, v338, v344, v345, v346, v347, v348, v354, v355, v356, v372, v373, v379, v380, v381, v382, v383, v384, v385, v391, v392, v393, v394, v395, v396, v397, v398, v409, v410, v411, v412, v413, v419, v420, v421, v422, v423, v424, v425 int32
	var lookahead, i, lookahead1 *int32
	var conv4, idxprom, idxprom10 int64
	var v3, frombool, v10, v23, v42, v51, v59, v67, v68, v74, v79, v86, v94, v102, v107, v112, v117, v122, v127, v135, v145, v155, v165, v175, v185, v195, v205, v215, v225, v235, v245, v255, v265, v275, v285, v295, v305, v315, v323, v328, v339, v349, v357, v362, v367, v374, v386, v399, v404, v414, v426 byte
	var result, skip, eof *byte
	var v77, v82, v89, v97, v105, v110, v115, v120, v125, v130, v138, v148, v158, v168, v178, v188, v198, v208, v218, v228, v238, v248, v258, v268, v278, v288, v298, v308, v318, v326, v331, v342, v352, v360, v365, v370, v377, v389, v402, v407, v417 func(*TSLexer)
	var mark_end, mark_end206, mark_end217, mark_end231, mark_end245, mark_end249, mark_end253, mark_end257, mark_end261, mark_end265, mark_end279, mark_end300, mark_end321, mark_end342, mark_end363, mark_end384, mark_end405, mark_end426, mark_end447, mark_end468, mark_end489, mark_end510, mark_end531, mark_end552, mark_end573, mark_end594, mark_end615, mark_end636, mark_end657, mark_end671, mark_end675, mark_end698, mark_end718, mark_end732, mark_end736, mark_end740, mark_end751, mark_end777, mark_end806, mark_end810, mark_end831 *func(*TSLexer)
	var v1 func(*TSLexer, bool)
	var advance *func(*TSLexer, bool)
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, tobool29, v24, cmp31, v25, cmp35, v26, cmp39, v27, cmp43, v28, cmp47, v29, cmp50, v30, cmp53, v31, cmp57, v32, cmp60, v33, cmp64, v34, cmp67, v35, cmp70, v36, cmp73, v37, cmp76, v38, cmp79, v39, cmp82, v40, cmp85, v41, cmp88, v42, tobool92, v43, cmp94, v44, cmp98, v45, cmp102, v46, cmp106, v47, cmp110, v48, cmp113, v49, cmp116, v50, cmp120, v51, tobool124, v52, cmp126, v53, cmp130, v54, cmp134, v55, cmp138, v56, cmp141, v57, cmp144, v58, cmp148, v59, tobool152, v60, cmp154, v61, cmp157, v62, cmp161, v63, cmp164, v64, cmp168, v65, cmp171, v66, call175, v67, tobool178, v68, tobool180, v69, cmp183, v70, cmp187, v71, cmp190, v72, cmp193, v73, cmp197, v74, tobool201, v75, result_symbol, v76, mark_end, v77, v78, v79, tobool203, v80, result_symbol205, v81, mark_end206, v82, v83, v84, cmp207, v85, cmp210, v86, tobool214, v87, result_symbol216, v88, mark_end217, v89, v90, v91, call218, v92, cmp221, v93, cmp224, v94, tobool228, v95, result_symbol230, v96, mark_end231, v97, v98, v99, call232, v100, cmp235, v101, cmp238, v102, tobool242, v103, result_symbol244, v104, mark_end245, v105, v106, v107, tobool246, v108, result_symbol248, v109, mark_end249, v110, v111, v112, tobool250, v113, result_symbol252, v114, mark_end253, v115, v116, v117, tobool254, v118, result_symbol256, v119, mark_end257, v120, v121, v122, tobool258, v123, result_symbol260, v124, mark_end261, v125, v126, v127, tobool262, v128, result_symbol264, v129, mark_end265, v130, v131, v132, call266, v133, cmp269, v134, cmp272, v135, tobool276, v136, result_symbol278, v137, mark_end279, v138, v139, v140, cmp280, v141, cmp283, v142, call287, v143, cmp290, v144, cmp293, v145, tobool297, v146, result_symbol299, v147, mark_end300, v148, v149, v150, cmp301, v151, cmp304, v152, call308, v153, cmp311, v154, cmp314, v155, tobool318, v156, result_symbol320, v157, mark_end321, v158, v159, v160, cmp322, v161, cmp325, v162, call329, v163, cmp332, v164, cmp335, v165, tobool339, v166, result_symbol341, v167, mark_end342, v168, v169, v170, cmp343, v171, cmp346, v172, call350, v173, cmp353, v174, cmp356, v175, tobool360, v176, result_symbol362, v177, mark_end363, v178, v179, v180, cmp364, v181, cmp367, v182, call371, v183, cmp374, v184, cmp377, v185, tobool381, v186, result_symbol383, v187, mark_end384, v188, v189, v190, cmp385, v191, cmp388, v192, call392, v193, cmp395, v194, cmp398, v195, tobool402, v196, result_symbol404, v197, mark_end405, v198, v199, v200, cmp406, v201, cmp409, v202, call413, v203, cmp416, v204, cmp419, v205, tobool423, v206, result_symbol425, v207, mark_end426, v208, v209, v210, cmp427, v211, cmp430, v212, call434, v213, cmp437, v214, cmp440, v215, tobool444, v216, result_symbol446, v217, mark_end447, v218, v219, v220, cmp448, v221, cmp451, v222, call455, v223, cmp458, v224, cmp461, v225, tobool465, v226, result_symbol467, v227, mark_end468, v228, v229, v230, cmp469, v231, cmp472, v232, call476, v233, cmp479, v234, cmp482, v235, tobool486, v236, result_symbol488, v237, mark_end489, v238, v239, v240, cmp490, v241, cmp493, v242, call497, v243, cmp500, v244, cmp503, v245, tobool507, v246, result_symbol509, v247, mark_end510, v248, v249, v250, cmp511, v251, cmp514, v252, call518, v253, cmp521, v254, cmp524, v255, tobool528, v256, result_symbol530, v257, mark_end531, v258, v259, v260, cmp532, v261, cmp535, v262, call539, v263, cmp542, v264, cmp545, v265, tobool549, v266, result_symbol551, v267, mark_end552, v268, v269, v270, cmp553, v271, cmp556, v272, call560, v273, cmp563, v274, cmp566, v275, tobool570, v276, result_symbol572, v277, mark_end573, v278, v279, v280, cmp574, v281, cmp577, v282, call581, v283, cmp584, v284, cmp587, v285, tobool591, v286, result_symbol593, v287, mark_end594, v288, v289, v290, cmp595, v291, cmp598, v292, call602, v293, cmp605, v294, cmp608, v295, tobool612, v296, result_symbol614, v297, mark_end615, v298, v299, v300, cmp616, v301, cmp619, v302, call623, v303, cmp626, v304, cmp629, v305, tobool633, v306, result_symbol635, v307, mark_end636, v308, v309, v310, cmp637, v311, cmp640, v312, call644, v313, cmp647, v314, cmp650, v315, tobool654, v316, result_symbol656, v317, mark_end657, v318, v319, v320, call658, v321, cmp661, v322, cmp664, v323, tobool668, v324, result_symbol670, v325, mark_end671, v326, v327, v328, tobool672, v329, result_symbol674, v330, mark_end675, v331, v332, v333, cmp676, v334, cmp679, v335, cmp682, v336, cmp685, v337, cmp688, v338, cmp691, v339, tobool695, v340, result_symbol697, v341, mark_end698, v342, v343, v344, cmp699, v345, cmp702, v346, cmp705, v347, cmp708, v348, cmp711, v349, tobool715, v350, result_symbol717, v351, mark_end718, v352, v353, v354, call719, v355, cmp722, v356, cmp725, v357, tobool729, v358, result_symbol731, v359, mark_end732, v360, v361, v362, tobool733, v363, result_symbol735, v364, mark_end736, v365, v366, v367, tobool737, v368, result_symbol739, v369, mark_end740, v370, v371, v372, cmp741, v373, cmp744, v374, tobool748, v375, result_symbol750, v376, mark_end751, v377, v378, v379, cmp752, v380, cmp755, v381, cmp758, v382, cmp761, v383, cmp764, v384, cmp767, v385, cmp770, v386, tobool774, v387, result_symbol776, v388, mark_end777, v389, v390, v391, cmp778, v392, cmp781, v393, cmp784, v394, cmp787, v395, cmp790, v396, cmp793, v397, cmp796, v398, cmp799, v399, tobool803, v400, result_symbol805, v401, mark_end806, v402, v403, v404, tobool807, v405, result_symbol809, v406, mark_end810, v407, v408, v409, cmp811, v410, cmp815, v411, cmp818, v412, cmp821, v413, cmp824, v414, tobool828, v415, result_symbol830, v416, mark_end831, v417, v418, v419, cmp832, v420, cmp835, v421, cmp838, v422, cmp841, v423, cmp845, v424, cmp848, v425, cmp851, v426, tobool855, v427

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
	i = new(int32)
	*lexer_addr = lexer
	*state_addr = state
	*result = 0
	*skip = 0
	*eof = 0
	goto start

next_state:
	v0 = *lexer_addr
	advance = &v0.F2
	v1 = *advance
	v2 = *lexer_addr
	v3 = *skip
	tobool = (v3 & 1) != 0
	v1(v2, tobool)
	goto start

start:
	*skip = 0
	v4 = *lexer_addr
	lookahead1 = &v4.F0
	v5 = *lookahead1
	*lookahead = v5
	v6 = *lexer_addr
	eof2 = &v6.F6
	v7 = *eof2
	v8 = *lexer_addr
	call = v7(v8)
	if call {
		frombool = 1
	} else {
		frombool = 0
	}
	*eof = frombool
	v9 = *state_addr
	conv = int32(uint32(uint16(v9)))
	switch conv {
	case 0:
		goto sw_bb
	case 1:
		goto sw_bb30
	case 2:
		goto sw_bb93
	case 3:
		goto sw_bb125
	case 4:
		goto sw_bb153
	case 5:
		goto sw_bb179
	case 6:
		goto sw_bb202
	case 7:
		goto sw_bb204
	case 8:
		goto sw_bb215
	case 9:
		goto sw_bb229
	case 10:
		goto sw_bb243
	case 11:
		goto sw_bb247
	case 12:
		goto sw_bb251
	case 13:
		goto sw_bb255
	case 14:
		goto sw_bb259
	case 15:
		goto sw_bb263
	case 16:
		goto sw_bb277
	case 17:
		goto sw_bb298
	case 18:
		goto sw_bb319
	case 19:
		goto sw_bb340
	case 20:
		goto sw_bb361
	case 21:
		goto sw_bb382
	case 22:
		goto sw_bb403
	case 23:
		goto sw_bb424
	case 24:
		goto sw_bb445
	case 25:
		goto sw_bb466
	case 26:
		goto sw_bb487
	case 27:
		goto sw_bb508
	case 28:
		goto sw_bb529
	case 29:
		goto sw_bb550
	case 30:
		goto sw_bb571
	case 31:
		goto sw_bb592
	case 32:
		goto sw_bb613
	case 33:
		goto sw_bb634
	case 34:
		goto sw_bb655
	case 35:
		goto sw_bb669
	case 36:
		goto sw_bb673
	case 37:
		goto sw_bb696
	case 38:
		goto sw_bb716
	case 39:
		goto sw_bb730
	case 40:
		goto sw_bb734
	case 41:
		goto sw_bb738
	case 42:
		goto sw_bb749
	case 43:
		goto sw_bb775
	case 44:
		goto sw_bb804
	case 45:
		goto sw_bb808
	case 46:
		goto sw_bb829
	default:
		goto sw_default
	}

sw_bb:
	v10 = *eof
	tobool3 = (v10 & 1) != 0
	if tobool3 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*state_addr = 6
	goto next_state

if_end:
	*i = 0
	goto for_cond

for_cond:
	v11 = *i
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(20)
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v12 = *i
	idxprom = int64(uint64(uint32(v12)))
	arrayidx = &ts_lex_map[idxprom]
	v13 = *arrayidx
	conv6 = int32(uint32(uint16(v13)))
	v14 = *lookahead
	cmp7 = conv6 == v14
	if cmp7 {
		goto if_then9
	} else {
		goto if_end12
	}

if_then9:
	v15 = *i
	add = v15 + 1
	idxprom10 = int64(uint64(uint32(add)))
	arrayidx11 = &ts_lex_map[idxprom10]
	v16 = *arrayidx11
	*state_addr = v16
	goto next_state

if_end12:
	goto for_inc

for_inc:
	v17 = *i
	add13 = v17 + 2
	*i = add13
	goto for_cond

for_end:
	v18 = *lookahead
	cmp14 = 9 <= v18
	if cmp14 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false
	}

land_lhs_true:
	v19 = *lookahead
	cmp16 = v19 <= 13
	if cmp16 {
		goto if_then20
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v20 = *lookahead
	cmp18 = v20 == 32
	if cmp18 {
		goto if_then20
	} else {
		goto if_end21
	}

if_then20:
	*skip = 1
	*state_addr = 0
	goto next_state

if_end21:
	v21 = *lookahead
	cmp22 = 48 <= v21
	if cmp22 {
		goto land_lhs_true24
	} else {
		goto if_end28
	}

land_lhs_true24:
	v22 = *lookahead
	cmp25 = v22 <= 57
	if cmp25 {
		goto if_then27
	} else {
		goto if_end28
	}

if_then27:
	*state_addr = 41
	goto next_state

if_end28:
	v23 = *result
	tobool29 = (v23 & 1) != 0
	*retval = tobool29
	goto _return

sw_bb30:
	v24 = *lookahead
	cmp31 = v24 == 34
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*state_addr = 40
	goto next_state

if_end34:
	v25 = *lookahead
	cmp35 = v25 == 41
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*state_addr = 14
	goto next_state

if_end38:
	v26 = *lookahead
	cmp39 = v26 == 123
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*state_addr = 10
	goto next_state

if_end42:
	v27 = *lookahead
	cmp43 = v27 == 125
	if cmp43 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*state_addr = 12
	goto next_state

if_end46:
	v28 = *lookahead
	cmp47 = 9 <= v28
	if cmp47 {
		goto land_lhs_true49
	} else {
		goto lor_lhs_false52
	}

land_lhs_true49:
	v29 = *lookahead
	cmp50 = v29 <= 13
	if cmp50 {
		goto if_then55
	} else {
		goto lor_lhs_false52
	}

lor_lhs_false52:
	v30 = *lookahead
	cmp53 = v30 == 32
	if cmp53 {
		goto if_then55
	} else {
		goto if_end56
	}

if_then55:
	*skip = 1
	*state_addr = 1
	goto next_state

if_end56:
	v31 = *lookahead
	cmp57 = 48 <= v31
	if cmp57 {
		goto land_lhs_true59
	} else {
		goto if_end63
	}

land_lhs_true59:
	v32 = *lookahead
	cmp60 = v32 <= 57
	if cmp60 {
		goto if_then62
	} else {
		goto if_end63
	}

if_then62:
	*state_addr = 41
	goto next_state

if_end63:
	v33 = *lookahead
	cmp64 = v33 == 33
	if cmp64 {
		goto if_then90
	} else {
		goto lor_lhs_false66
	}

lor_lhs_false66:
	v34 = *lookahead
	cmp67 = v34 == 36
	if cmp67 {
		goto if_then90
	} else {
		goto lor_lhs_false69
	}

lor_lhs_false69:
	v35 = *lookahead
	cmp70 = v35 == 38
	if cmp70 {
		goto if_then90
	} else {
		goto lor_lhs_false72
	}

lor_lhs_false72:
	v36 = *lookahead
	cmp73 = v36 == 42
	if cmp73 {
		goto if_then90
	} else {
		goto lor_lhs_false75
	}

lor_lhs_false75:
	v37 = *lookahead
	cmp76 = v37 == 43
	if cmp76 {
		goto if_then90
	} else {
		goto lor_lhs_false78
	}

lor_lhs_false78:
	v38 = *lookahead
	cmp79 = 45 <= v38
	if cmp79 {
		goto land_lhs_true81
	} else {
		goto lor_lhs_false84
	}

land_lhs_true81:
	v39 = *lookahead
	cmp82 = v39 <= 60
	if cmp82 {
		goto if_then90
	} else {
		goto lor_lhs_false84
	}

lor_lhs_false84:
	v40 = *lookahead
	cmp85 = 62 <= v40
	if cmp85 {
		goto land_lhs_true87
	} else {
		goto if_end91
	}

land_lhs_true87:
	v41 = *lookahead
	cmp88 = v41 <= 126
	if cmp88 {
		goto if_then90
	} else {
		goto if_end91
	}

if_then90:
	*state_addr = 38
	goto next_state

if_end91:
	v42 = *result
	tobool92 = (v42 & 1) != 0
	*retval = tobool92
	goto _return

sw_bb93:
	v43 = *lookahead
	cmp94 = v43 == 34
	if cmp94 {
		goto if_then96
	} else {
		goto if_end97
	}

if_then96:
	*state_addr = 40
	goto next_state

if_end97:
	v44 = *lookahead
	cmp98 = v44 == 92
	if cmp98 {
		goto if_then100
	} else {
		goto if_end101
	}

if_then100:
	*state_addr = 46
	goto next_state

if_end101:
	v45 = *lookahead
	cmp102 = v45 == 123
	if cmp102 {
		goto if_then104
	} else {
		goto if_end105
	}

if_then104:
	*state_addr = 10
	goto next_state

if_end105:
	v46 = *lookahead
	cmp106 = v46 == 125
	if cmp106 {
		goto if_then108
	} else {
		goto if_end109
	}

if_then108:
	*state_addr = 12
	goto next_state

if_end109:
	v47 = *lookahead
	cmp110 = 9 <= v47
	if cmp110 {
		goto land_lhs_true112
	} else {
		goto lor_lhs_false115
	}

land_lhs_true112:
	v48 = *lookahead
	cmp113 = v48 <= 13
	if cmp113 {
		goto if_then118
	} else {
		goto lor_lhs_false115
	}

lor_lhs_false115:
	v49 = *lookahead
	cmp116 = v49 == 32
	if cmp116 {
		goto if_then118
	} else {
		goto if_end119
	}

if_then118:
	*skip = 1
	*state_addr = 2
	goto next_state

if_end119:
	v50 = *lookahead
	cmp120 = v50 != 0
	if cmp120 {
		goto if_then122
	} else {
		goto if_end123
	}

if_then122:
	*state_addr = 43
	goto next_state

if_end123:
	v51 = *result
	tobool124 = (v51 & 1) != 0
	*retval = tobool124
	goto _return

sw_bb125:
	v52 = *lookahead
	cmp126 = v52 == 92
	if cmp126 {
		goto if_then128
	} else {
		goto if_end129
	}

if_then128:
	*state_addr = 46
	goto next_state

if_end129:
	v53 = *lookahead
	cmp130 = v53 == 123
	if cmp130 {
		goto if_then132
	} else {
		goto if_end133
	}

if_then132:
	*state_addr = 10
	goto next_state

if_end133:
	v54 = *lookahead
	cmp134 = v54 == 125
	if cmp134 {
		goto if_then136
	} else {
		goto if_end137
	}

if_then136:
	*state_addr = 12
	goto next_state

if_end137:
	v55 = *lookahead
	cmp138 = 9 <= v55
	if cmp138 {
		goto land_lhs_true140
	} else {
		goto lor_lhs_false143
	}

land_lhs_true140:
	v56 = *lookahead
	cmp141 = v56 <= 13
	if cmp141 {
		goto if_then146
	} else {
		goto lor_lhs_false143
	}

lor_lhs_false143:
	v57 = *lookahead
	cmp144 = v57 == 32
	if cmp144 {
		goto if_then146
	} else {
		goto if_end147
	}

if_then146:
	*skip = 1
	*state_addr = 3
	goto next_state

if_end147:
	v58 = *lookahead
	cmp148 = v58 != 0
	if cmp148 {
		goto if_then150
	} else {
		goto if_end151
	}

if_then150:
	*state_addr = 42
	goto next_state

if_end151:
	v59 = *result
	tobool152 = (v59 & 1) != 0
	*retval = tobool152
	goto _return

sw_bb153:
	v60 = *lookahead
	cmp154 = v60 == 67
	if cmp154 {
		goto if_then159
	} else {
		goto lor_lhs_false156
	}

lor_lhs_false156:
	v61 = *lookahead
	cmp157 = v61 == 99
	if cmp157 {
		goto if_then159
	} else {
		goto if_end160
	}

if_then159:
	*state_addr = 29
	goto next_state

if_end160:
	v62 = *lookahead
	cmp161 = v62 == 80
	if cmp161 {
		goto if_then166
	} else {
		goto lor_lhs_false163
	}

lor_lhs_false163:
	v63 = *lookahead
	cmp164 = v63 == 112
	if cmp164 {
		goto if_then166
	} else {
		goto if_end167
	}

if_then166:
	*state_addr = 30
	goto next_state

if_end167:
	v64 = *lookahead
	cmp168 = v64 == 83
	if cmp168 {
		goto if_then173
	} else {
		goto lor_lhs_false170
	}

lor_lhs_false170:
	v65 = *lookahead
	cmp171 = v65 == 115
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*state_addr = 33
	goto next_state

if_end174:
	v66 = *lookahead
	call175 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v66)
	if call175 {
		goto if_then176
	} else {
		goto if_end177
	}

if_then176:
	*state_addr = 34
	goto next_state

if_end177:
	v67 = *result
	tobool178 = (v67 & 1) != 0
	*retval = tobool178
	goto _return

sw_bb179:
	v68 = *eof
	tobool180 = (v68 & 1) != 0
	if tobool180 {
		goto if_then181
	} else {
		goto if_end182
	}

if_then181:
	*state_addr = 6
	goto next_state

if_end182:
	v69 = *lookahead
	cmp183 = v69 == 64
	if cmp183 {
		goto if_then185
	} else {
		goto if_end186
	}

if_then185:
	*state_addr = 4
	goto next_state

if_end186:
	v70 = *lookahead
	cmp187 = 9 <= v70
	if cmp187 {
		goto land_lhs_true189
	} else {
		goto lor_lhs_false192
	}

land_lhs_true189:
	v71 = *lookahead
	cmp190 = v71 <= 13
	if cmp190 {
		goto if_then195
	} else {
		goto lor_lhs_false192
	}

lor_lhs_false192:
	v72 = *lookahead
	cmp193 = v72 == 32
	if cmp193 {
		goto if_then195
	} else {
		goto if_end196
	}

if_then195:
	*skip = 1
	*state_addr = 5
	goto next_state

if_end196:
	v73 = *lookahead
	cmp197 = v73 != 0
	if cmp197 {
		goto if_then199
	} else {
		goto if_end200
	}

if_then199:
	*state_addr = 7
	goto next_state

if_end200:
	v74 = *result
	tobool201 = (v74 & 1) != 0
	*retval = tobool201
	goto _return

sw_bb202:
	*result = 1
	v75 = *lexer_addr
	result_symbol = &v75.F1
	*result_symbol = 0
	v76 = *lexer_addr
	mark_end = &v76.F3
	v77 = *mark_end
	v78 = *lexer_addr
	v77(v78)
	v79 = *result
	tobool203 = (v79 & 1) != 0
	*retval = tobool203
	goto _return

sw_bb204:
	*result = 1
	v80 = *lexer_addr
	result_symbol205 = &v80.F1
	*result_symbol205 = 1
	v81 = *lexer_addr
	mark_end206 = &v81.F3
	v82 = *mark_end206
	v83 = *lexer_addr
	v82(v83)
	v84 = *lookahead
	cmp207 = v84 != 0
	if cmp207 {
		goto land_lhs_true209
	} else {
		goto if_end213
	}

land_lhs_true209:
	v85 = *lookahead
	cmp210 = v85 != 64
	if cmp210 {
		goto if_then212
	} else {
		goto if_end213
	}

if_then212:
	*state_addr = 7
	goto next_state

if_end213:
	v86 = *result
	tobool214 = (v86 & 1) != 0
	*retval = tobool214
	goto _return

sw_bb215:
	*result = 1
	v87 = *lexer_addr
	result_symbol216 = &v87.F1
	*result_symbol216 = 2
	v88 = *lexer_addr
	mark_end217 = &v88.F3
	v89 = *mark_end217
	v90 = *lexer_addr
	v89(v90)
	v91 = *lookahead
	call218 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v91)
	if call218 {
		goto if_then226
	} else {
		goto lor_lhs_false220
	}

lor_lhs_false220:
	v92 = *lookahead
	cmp221 = 48 <= v92
	if cmp221 {
		goto land_lhs_true223
	} else {
		goto if_end227
	}

land_lhs_true223:
	v93 = *lookahead
	cmp224 = v93 <= 57
	if cmp224 {
		goto if_then226
	} else {
		goto if_end227
	}

if_then226:
	*state_addr = 34
	goto next_state

if_end227:
	v94 = *result
	tobool228 = (v94 & 1) != 0
	*retval = tobool228
	goto _return

sw_bb229:
	*result = 1
	v95 = *lexer_addr
	result_symbol230 = &v95.F1
	*result_symbol230 = 3
	v96 = *lexer_addr
	mark_end231 = &v96.F3
	v97 = *mark_end231
	v98 = *lexer_addr
	v97(v98)
	v99 = *lookahead
	call232 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v99)
	if call232 {
		goto if_then240
	} else {
		goto lor_lhs_false234
	}

lor_lhs_false234:
	v100 = *lookahead
	cmp235 = 48 <= v100
	if cmp235 {
		goto land_lhs_true237
	} else {
		goto if_end241
	}

land_lhs_true237:
	v101 = *lookahead
	cmp238 = v101 <= 57
	if cmp238 {
		goto if_then240
	} else {
		goto if_end241
	}

if_then240:
	*state_addr = 34
	goto next_state

if_end241:
	v102 = *result
	tobool242 = (v102 & 1) != 0
	*retval = tobool242
	goto _return

sw_bb243:
	*result = 1
	v103 = *lexer_addr
	result_symbol244 = &v103.F1
	*result_symbol244 = 4
	v104 = *lexer_addr
	mark_end245 = &v104.F3
	v105 = *mark_end245
	v106 = *lexer_addr
	v105(v106)
	v107 = *result
	tobool246 = (v107 & 1) != 0
	*retval = tobool246
	goto _return

sw_bb247:
	*result = 1
	v108 = *lexer_addr
	result_symbol248 = &v108.F1
	*result_symbol248 = 5
	v109 = *lexer_addr
	mark_end249 = &v109.F3
	v110 = *mark_end249
	v111 = *lexer_addr
	v110(v111)
	v112 = *result
	tobool250 = (v112 & 1) != 0
	*retval = tobool250
	goto _return

sw_bb251:
	*result = 1
	v113 = *lexer_addr
	result_symbol252 = &v113.F1
	*result_symbol252 = 6
	v114 = *lexer_addr
	mark_end253 = &v114.F3
	v115 = *mark_end253
	v116 = *lexer_addr
	v115(v116)
	v117 = *result
	tobool254 = (v117 & 1) != 0
	*retval = tobool254
	goto _return

sw_bb255:
	*result = 1
	v118 = *lexer_addr
	result_symbol256 = &v118.F1
	*result_symbol256 = 7
	v119 = *lexer_addr
	mark_end257 = &v119.F3
	v120 = *mark_end257
	v121 = *lexer_addr
	v120(v121)
	v122 = *result
	tobool258 = (v122 & 1) != 0
	*retval = tobool258
	goto _return

sw_bb259:
	*result = 1
	v123 = *lexer_addr
	result_symbol260 = &v123.F1
	*result_symbol260 = 8
	v124 = *lexer_addr
	mark_end261 = &v124.F3
	v125 = *mark_end261
	v126 = *lexer_addr
	v125(v126)
	v127 = *result
	tobool262 = (v127 & 1) != 0
	*retval = tobool262
	goto _return

sw_bb263:
	*result = 1
	v128 = *lexer_addr
	result_symbol264 = &v128.F1
	*result_symbol264 = 9
	v129 = *lexer_addr
	mark_end265 = &v129.F3
	v130 = *mark_end265
	v131 = *lexer_addr
	v130(v131)
	v132 = *lookahead
	call266 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v132)
	if call266 {
		goto if_then274
	} else {
		goto lor_lhs_false268
	}

lor_lhs_false268:
	v133 = *lookahead
	cmp269 = 48 <= v133
	if cmp269 {
		goto land_lhs_true271
	} else {
		goto if_end275
	}

land_lhs_true271:
	v134 = *lookahead
	cmp272 = v134 <= 57
	if cmp272 {
		goto if_then274
	} else {
		goto if_end275
	}

if_then274:
	*state_addr = 34
	goto next_state

if_end275:
	v135 = *result
	tobool276 = (v135 & 1) != 0
	*retval = tobool276
	goto _return

sw_bb277:
	*result = 1
	v136 = *lexer_addr
	result_symbol278 = &v136.F1
	*result_symbol278 = 10
	v137 = *lexer_addr
	mark_end279 = &v137.F3
	v138 = *mark_end279
	v139 = *lexer_addr
	v138(v139)
	v140 = *lookahead
	cmp280 = v140 == 65
	if cmp280 {
		goto if_then285
	} else {
		goto lor_lhs_false282
	}

lor_lhs_false282:
	v141 = *lookahead
	cmp283 = v141 == 97
	if cmp283 {
		goto if_then285
	} else {
		goto if_end286
	}

if_then285:
	*state_addr = 24
	goto next_state

if_end286:
	v142 = *lookahead
	call287 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v142)
	if call287 {
		goto if_then295
	} else {
		goto lor_lhs_false289
	}

lor_lhs_false289:
	v143 = *lookahead
	cmp290 = 48 <= v143
	if cmp290 {
		goto land_lhs_true292
	} else {
		goto if_end296
	}

land_lhs_true292:
	v144 = *lookahead
	cmp293 = v144 <= 57
	if cmp293 {
		goto if_then295
	} else {
		goto if_end296
	}

if_then295:
	*state_addr = 34
	goto next_state

if_end296:
	v145 = *result
	tobool297 = (v145 & 1) != 0
	*retval = tobool297
	goto _return

sw_bb298:
	*result = 1
	v146 = *lexer_addr
	result_symbol299 = &v146.F1
	*result_symbol299 = 10
	v147 = *lexer_addr
	mark_end300 = &v147.F3
	v148 = *mark_end300
	v149 = *lexer_addr
	v148(v149)
	v150 = *lookahead
	cmp301 = v150 == 66
	if cmp301 {
		goto if_then306
	} else {
		goto lor_lhs_false303
	}

lor_lhs_false303:
	v151 = *lookahead
	cmp304 = v151 == 98
	if cmp304 {
		goto if_then306
	} else {
		goto if_end307
	}

if_then306:
	*state_addr = 23
	goto next_state

if_end307:
	v152 = *lookahead
	call308 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v152)
	if call308 {
		goto if_then316
	} else {
		goto lor_lhs_false310
	}

lor_lhs_false310:
	v153 = *lookahead
	cmp311 = 48 <= v153
	if cmp311 {
		goto land_lhs_true313
	} else {
		goto if_end317
	}

land_lhs_true313:
	v154 = *lookahead
	cmp314 = v154 <= 57
	if cmp314 {
		goto if_then316
	} else {
		goto if_end317
	}

if_then316:
	*state_addr = 34
	goto next_state

if_end317:
	v155 = *result
	tobool318 = (v155 & 1) != 0
	*retval = tobool318
	goto _return

sw_bb319:
	*result = 1
	v156 = *lexer_addr
	result_symbol320 = &v156.F1
	*result_symbol320 = 10
	v157 = *lexer_addr
	mark_end321 = &v157.F3
	v158 = *mark_end321
	v159 = *lexer_addr
	v158(v159)
	v160 = *lookahead
	cmp322 = v160 == 69
	if cmp322 {
		goto if_then327
	} else {
		goto lor_lhs_false324
	}

lor_lhs_false324:
	v161 = *lookahead
	cmp325 = v161 == 101
	if cmp325 {
		goto if_then327
	} else {
		goto if_end328
	}

if_then327:
	*state_addr = 16
	goto next_state

if_end328:
	v162 = *lookahead
	call329 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v162)
	if call329 {
		goto if_then337
	} else {
		goto lor_lhs_false331
	}

lor_lhs_false331:
	v163 = *lookahead
	cmp332 = 48 <= v163
	if cmp332 {
		goto land_lhs_true334
	} else {
		goto if_end338
	}

land_lhs_true334:
	v164 = *lookahead
	cmp335 = v164 <= 57
	if cmp335 {
		goto if_then337
	} else {
		goto if_end338
	}

if_then337:
	*state_addr = 34
	goto next_state

if_end338:
	v165 = *result
	tobool339 = (v165 & 1) != 0
	*retval = tobool339
	goto _return

sw_bb340:
	*result = 1
	v166 = *lexer_addr
	result_symbol341 = &v166.F1
	*result_symbol341 = 10
	v167 = *lexer_addr
	mark_end342 = &v167.F3
	v168 = *mark_end342
	v169 = *lexer_addr
	v168(v169)
	v170 = *lookahead
	cmp343 = v170 == 69
	if cmp343 {
		goto if_then348
	} else {
		goto lor_lhs_false345
	}

lor_lhs_false345:
	v171 = *lookahead
	cmp346 = v171 == 101
	if cmp346 {
		goto if_then348
	} else {
		goto if_end349
	}

if_then348:
	*state_addr = 15
	goto next_state

if_end349:
	v172 = *lookahead
	call350 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v172)
	if call350 {
		goto if_then358
	} else {
		goto lor_lhs_false352
	}

lor_lhs_false352:
	v173 = *lookahead
	cmp353 = 48 <= v173
	if cmp353 {
		goto land_lhs_true355
	} else {
		goto if_end359
	}

land_lhs_true355:
	v174 = *lookahead
	cmp356 = v174 <= 57
	if cmp356 {
		goto if_then358
	} else {
		goto if_end359
	}

if_then358:
	*state_addr = 34
	goto next_state

if_end359:
	v175 = *result
	tobool360 = (v175 & 1) != 0
	*retval = tobool360
	goto _return

sw_bb361:
	*result = 1
	v176 = *lexer_addr
	result_symbol362 = &v176.F1
	*result_symbol362 = 10
	v177 = *lexer_addr
	mark_end363 = &v177.F3
	v178 = *mark_end363
	v179 = *lexer_addr
	v178(v179)
	v180 = *lookahead
	cmp364 = v180 == 69
	if cmp364 {
		goto if_then369
	} else {
		goto lor_lhs_false366
	}

lor_lhs_false366:
	v181 = *lookahead
	cmp367 = v181 == 101
	if cmp367 {
		goto if_then369
	} else {
		goto if_end370
	}

if_then369:
	*state_addr = 28
	goto next_state

if_end370:
	v182 = *lookahead
	call371 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v182)
	if call371 {
		goto if_then379
	} else {
		goto lor_lhs_false373
	}

lor_lhs_false373:
	v183 = *lookahead
	cmp374 = 48 <= v183
	if cmp374 {
		goto land_lhs_true376
	} else {
		goto if_end380
	}

land_lhs_true376:
	v184 = *lookahead
	cmp377 = v184 <= 57
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*state_addr = 34
	goto next_state

if_end380:
	v185 = *result
	tobool381 = (v185 & 1) != 0
	*retval = tobool381
	goto _return

sw_bb382:
	*result = 1
	v186 = *lexer_addr
	result_symbol383 = &v186.F1
	*result_symbol383 = 10
	v187 = *lexer_addr
	mark_end384 = &v187.F3
	v188 = *mark_end384
	v189 = *lexer_addr
	v188(v189)
	v190 = *lookahead
	cmp385 = v190 == 71
	if cmp385 {
		goto if_then390
	} else {
		goto lor_lhs_false387
	}

lor_lhs_false387:
	v191 = *lookahead
	cmp388 = v191 == 103
	if cmp388 {
		goto if_then390
	} else {
		goto if_end391
	}

if_then390:
	*state_addr = 9
	goto next_state

if_end391:
	v192 = *lookahead
	call392 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v192)
	if call392 {
		goto if_then400
	} else {
		goto lor_lhs_false394
	}

lor_lhs_false394:
	v193 = *lookahead
	cmp395 = 48 <= v193
	if cmp395 {
		goto land_lhs_true397
	} else {
		goto if_end401
	}

land_lhs_true397:
	v194 = *lookahead
	cmp398 = v194 <= 57
	if cmp398 {
		goto if_then400
	} else {
		goto if_end401
	}

if_then400:
	*state_addr = 34
	goto next_state

if_end401:
	v195 = *result
	tobool402 = (v195 & 1) != 0
	*retval = tobool402
	goto _return

sw_bb403:
	*result = 1
	v196 = *lexer_addr
	result_symbol404 = &v196.F1
	*result_symbol404 = 10
	v197 = *lexer_addr
	mark_end405 = &v197.F3
	v198 = *mark_end405
	v199 = *lexer_addr
	v198(v199)
	v200 = *lookahead
	cmp406 = v200 == 73
	if cmp406 {
		goto if_then411
	} else {
		goto lor_lhs_false408
	}

lor_lhs_false408:
	v201 = *lookahead
	cmp409 = v201 == 105
	if cmp409 {
		goto if_then411
	} else {
		goto if_end412
	}

if_then411:
	*state_addr = 27
	goto next_state

if_end412:
	v202 = *lookahead
	call413 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v202)
	if call413 {
		goto if_then421
	} else {
		goto lor_lhs_false415
	}

lor_lhs_false415:
	v203 = *lookahead
	cmp416 = 48 <= v203
	if cmp416 {
		goto land_lhs_true418
	} else {
		goto if_end422
	}

land_lhs_true418:
	v204 = *lookahead
	cmp419 = v204 <= 57
	if cmp419 {
		goto if_then421
	} else {
		goto if_end422
	}

if_then421:
	*state_addr = 34
	goto next_state

if_end422:
	v205 = *result
	tobool423 = (v205 & 1) != 0
	*retval = tobool423
	goto _return

sw_bb424:
	*result = 1
	v206 = *lexer_addr
	result_symbol425 = &v206.F1
	*result_symbol425 = 10
	v207 = *lexer_addr
	mark_end426 = &v207.F3
	v208 = *mark_end426
	v209 = *lexer_addr
	v208(v209)
	v210 = *lookahead
	cmp427 = v210 == 76
	if cmp427 {
		goto if_then432
	} else {
		goto lor_lhs_false429
	}

lor_lhs_false429:
	v211 = *lookahead
	cmp430 = v211 == 108
	if cmp430 {
		goto if_then432
	} else {
		goto if_end433
	}

if_then432:
	*state_addr = 19
	goto next_state

if_end433:
	v212 = *lookahead
	call434 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v212)
	if call434 {
		goto if_then442
	} else {
		goto lor_lhs_false436
	}

lor_lhs_false436:
	v213 = *lookahead
	cmp437 = 48 <= v213
	if cmp437 {
		goto land_lhs_true439
	} else {
		goto if_end443
	}

land_lhs_true439:
	v214 = *lookahead
	cmp440 = v214 <= 57
	if cmp440 {
		goto if_then442
	} else {
		goto if_end443
	}

if_then442:
	*state_addr = 34
	goto next_state

if_end443:
	v215 = *result
	tobool444 = (v215 & 1) != 0
	*retval = tobool444
	goto _return

sw_bb445:
	*result = 1
	v216 = *lexer_addr
	result_symbol446 = &v216.F1
	*result_symbol446 = 10
	v217 = *lexer_addr
	mark_end447 = &v217.F3
	v218 = *mark_end447
	v219 = *lexer_addr
	v218(v219)
	v220 = *lookahead
	cmp448 = v220 == 77
	if cmp448 {
		goto if_then453
	} else {
		goto lor_lhs_false450
	}

lor_lhs_false450:
	v221 = *lookahead
	cmp451 = v221 == 109
	if cmp451 {
		goto if_then453
	} else {
		goto if_end454
	}

if_then453:
	*state_addr = 17
	goto next_state

if_end454:
	v222 = *lookahead
	call455 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v222)
	if call455 {
		goto if_then463
	} else {
		goto lor_lhs_false457
	}

lor_lhs_false457:
	v223 = *lookahead
	cmp458 = 48 <= v223
	if cmp458 {
		goto land_lhs_true460
	} else {
		goto if_end464
	}

land_lhs_true460:
	v224 = *lookahead
	cmp461 = v224 <= 57
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*state_addr = 34
	goto next_state

if_end464:
	v225 = *result
	tobool465 = (v225 & 1) != 0
	*retval = tobool465
	goto _return

sw_bb466:
	*result = 1
	v226 = *lexer_addr
	result_symbol467 = &v226.F1
	*result_symbol467 = 10
	v227 = *lexer_addr
	mark_end468 = &v227.F3
	v228 = *mark_end468
	v229 = *lexer_addr
	v228(v229)
	v230 = *lookahead
	cmp469 = v230 == 77
	if cmp469 {
		goto if_then474
	} else {
		goto lor_lhs_false471
	}

lor_lhs_false471:
	v231 = *lookahead
	cmp472 = v231 == 109
	if cmp472 {
		goto if_then474
	} else {
		goto if_end475
	}

if_then474:
	*state_addr = 26
	goto next_state

if_end475:
	v232 = *lookahead
	call476 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v232)
	if call476 {
		goto if_then484
	} else {
		goto lor_lhs_false478
	}

lor_lhs_false478:
	v233 = *lookahead
	cmp479 = 48 <= v233
	if cmp479 {
		goto land_lhs_true481
	} else {
		goto if_end485
	}

land_lhs_true481:
	v234 = *lookahead
	cmp482 = v234 <= 57
	if cmp482 {
		goto if_then484
	} else {
		goto if_end485
	}

if_then484:
	*state_addr = 34
	goto next_state

if_end485:
	v235 = *result
	tobool486 = (v235 & 1) != 0
	*retval = tobool486
	goto _return

sw_bb487:
	*result = 1
	v236 = *lexer_addr
	result_symbol488 = &v236.F1
	*result_symbol488 = 10
	v237 = *lexer_addr
	mark_end489 = &v237.F3
	v238 = *mark_end489
	v239 = *lexer_addr
	v238(v239)
	v240 = *lookahead
	cmp490 = v240 == 77
	if cmp490 {
		goto if_then495
	} else {
		goto lor_lhs_false492
	}

lor_lhs_false492:
	v241 = *lookahead
	cmp493 = v241 == 109
	if cmp493 {
		goto if_then495
	} else {
		goto if_end496
	}

if_then495:
	*state_addr = 20
	goto next_state

if_end496:
	v242 = *lookahead
	call497 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v242)
	if call497 {
		goto if_then505
	} else {
		goto lor_lhs_false499
	}

lor_lhs_false499:
	v243 = *lookahead
	cmp500 = 48 <= v243
	if cmp500 {
		goto land_lhs_true502
	} else {
		goto if_end506
	}

land_lhs_true502:
	v244 = *lookahead
	cmp503 = v244 <= 57
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*state_addr = 34
	goto next_state

if_end506:
	v245 = *result
	tobool507 = (v245 & 1) != 0
	*retval = tobool507
	goto _return

sw_bb508:
	*result = 1
	v246 = *lexer_addr
	result_symbol509 = &v246.F1
	*result_symbol509 = 10
	v247 = *lexer_addr
	mark_end510 = &v247.F3
	v248 = *mark_end510
	v249 = *lexer_addr
	v248(v249)
	v250 = *lookahead
	cmp511 = v250 == 78
	if cmp511 {
		goto if_then516
	} else {
		goto lor_lhs_false513
	}

lor_lhs_false513:
	v251 = *lookahead
	cmp514 = v251 == 110
	if cmp514 {
		goto if_then516
	} else {
		goto if_end517
	}

if_then516:
	*state_addr = 21
	goto next_state

if_end517:
	v252 = *lookahead
	call518 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v252)
	if call518 {
		goto if_then526
	} else {
		goto lor_lhs_false520
	}

lor_lhs_false520:
	v253 = *lookahead
	cmp521 = 48 <= v253
	if cmp521 {
		goto land_lhs_true523
	} else {
		goto if_end527
	}

land_lhs_true523:
	v254 = *lookahead
	cmp524 = v254 <= 57
	if cmp524 {
		goto if_then526
	} else {
		goto if_end527
	}

if_then526:
	*state_addr = 34
	goto next_state

if_end527:
	v255 = *result
	tobool528 = (v255 & 1) != 0
	*retval = tobool528
	goto _return

sw_bb529:
	*result = 1
	v256 = *lexer_addr
	result_symbol530 = &v256.F1
	*result_symbol530 = 10
	v257 = *lexer_addr
	mark_end531 = &v257.F3
	v258 = *mark_end531
	v259 = *lexer_addr
	v258(v259)
	v260 = *lookahead
	cmp532 = v260 == 78
	if cmp532 {
		goto if_then537
	} else {
		goto lor_lhs_false534
	}

lor_lhs_false534:
	v261 = *lookahead
	cmp535 = v261 == 110
	if cmp535 {
		goto if_then537
	} else {
		goto if_end538
	}

if_then537:
	*state_addr = 32
	goto next_state

if_end538:
	v262 = *lookahead
	call539 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v262)
	if call539 {
		goto if_then547
	} else {
		goto lor_lhs_false541
	}

lor_lhs_false541:
	v263 = *lookahead
	cmp542 = 48 <= v263
	if cmp542 {
		goto land_lhs_true544
	} else {
		goto if_end548
	}

land_lhs_true544:
	v264 = *lookahead
	cmp545 = v264 <= 57
	if cmp545 {
		goto if_then547
	} else {
		goto if_end548
	}

if_then547:
	*state_addr = 34
	goto next_state

if_end548:
	v265 = *result
	tobool549 = (v265 & 1) != 0
	*retval = tobool549
	goto _return

sw_bb550:
	*result = 1
	v266 = *lexer_addr
	result_symbol551 = &v266.F1
	*result_symbol551 = 10
	v267 = *lexer_addr
	mark_end552 = &v267.F3
	v268 = *mark_end552
	v269 = *lexer_addr
	v268(v269)
	v270 = *lookahead
	cmp553 = v270 == 79
	if cmp553 {
		goto if_then558
	} else {
		goto lor_lhs_false555
	}

lor_lhs_false555:
	v271 = *lookahead
	cmp556 = v271 == 111
	if cmp556 {
		goto if_then558
	} else {
		goto if_end559
	}

if_then558:
	*state_addr = 25
	goto next_state

if_end559:
	v272 = *lookahead
	call560 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v272)
	if call560 {
		goto if_then568
	} else {
		goto lor_lhs_false562
	}

lor_lhs_false562:
	v273 = *lookahead
	cmp563 = 48 <= v273
	if cmp563 {
		goto land_lhs_true565
	} else {
		goto if_end569
	}

land_lhs_true565:
	v274 = *lookahead
	cmp566 = v274 <= 57
	if cmp566 {
		goto if_then568
	} else {
		goto if_end569
	}

if_then568:
	*state_addr = 34
	goto next_state

if_end569:
	v275 = *result
	tobool570 = (v275 & 1) != 0
	*retval = tobool570
	goto _return

sw_bb571:
	*result = 1
	v276 = *lexer_addr
	result_symbol572 = &v276.F1
	*result_symbol572 = 10
	v277 = *lexer_addr
	mark_end573 = &v277.F3
	v278 = *mark_end573
	v279 = *lexer_addr
	v278(v279)
	v280 = *lookahead
	cmp574 = v280 == 82
	if cmp574 {
		goto if_then579
	} else {
		goto lor_lhs_false576
	}

lor_lhs_false576:
	v281 = *lookahead
	cmp577 = v281 == 114
	if cmp577 {
		goto if_then579
	} else {
		goto if_end580
	}

if_then579:
	*state_addr = 18
	goto next_state

if_end580:
	v282 = *lookahead
	call581 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v282)
	if call581 {
		goto if_then589
	} else {
		goto lor_lhs_false583
	}

lor_lhs_false583:
	v283 = *lookahead
	cmp584 = 48 <= v283
	if cmp584 {
		goto land_lhs_true586
	} else {
		goto if_end590
	}

land_lhs_true586:
	v284 = *lookahead
	cmp587 = v284 <= 57
	if cmp587 {
		goto if_then589
	} else {
		goto if_end590
	}

if_then589:
	*state_addr = 34
	goto next_state

if_end590:
	v285 = *result
	tobool591 = (v285 & 1) != 0
	*retval = tobool591
	goto _return

sw_bb592:
	*result = 1
	v286 = *lexer_addr
	result_symbol593 = &v286.F1
	*result_symbol593 = 10
	v287 = *lexer_addr
	mark_end594 = &v287.F3
	v288 = *mark_end594
	v289 = *lexer_addr
	v288(v289)
	v290 = *lookahead
	cmp595 = v290 == 82
	if cmp595 {
		goto if_then600
	} else {
		goto lor_lhs_false597
	}

lor_lhs_false597:
	v291 = *lookahead
	cmp598 = v291 == 114
	if cmp598 {
		goto if_then600
	} else {
		goto if_end601
	}

if_then600:
	*state_addr = 22
	goto next_state

if_end601:
	v292 = *lookahead
	call602 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v292)
	if call602 {
		goto if_then610
	} else {
		goto lor_lhs_false604
	}

lor_lhs_false604:
	v293 = *lookahead
	cmp605 = 48 <= v293
	if cmp605 {
		goto land_lhs_true607
	} else {
		goto if_end611
	}

land_lhs_true607:
	v294 = *lookahead
	cmp608 = v294 <= 57
	if cmp608 {
		goto if_then610
	} else {
		goto if_end611
	}

if_then610:
	*state_addr = 34
	goto next_state

if_end611:
	v295 = *result
	tobool612 = (v295 & 1) != 0
	*retval = tobool612
	goto _return

sw_bb613:
	*result = 1
	v296 = *lexer_addr
	result_symbol614 = &v296.F1
	*result_symbol614 = 10
	v297 = *lexer_addr
	mark_end615 = &v297.F3
	v298 = *mark_end615
	v299 = *lexer_addr
	v298(v299)
	v300 = *lookahead
	cmp616 = v300 == 84
	if cmp616 {
		goto if_then621
	} else {
		goto lor_lhs_false618
	}

lor_lhs_false618:
	v301 = *lookahead
	cmp619 = v301 == 116
	if cmp619 {
		goto if_then621
	} else {
		goto if_end622
	}

if_then621:
	*state_addr = 8
	goto next_state

if_end622:
	v302 = *lookahead
	call623 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v302)
	if call623 {
		goto if_then631
	} else {
		goto lor_lhs_false625
	}

lor_lhs_false625:
	v303 = *lookahead
	cmp626 = 48 <= v303
	if cmp626 {
		goto land_lhs_true628
	} else {
		goto if_end632
	}

land_lhs_true628:
	v304 = *lookahead
	cmp629 = v304 <= 57
	if cmp629 {
		goto if_then631
	} else {
		goto if_end632
	}

if_then631:
	*state_addr = 34
	goto next_state

if_end632:
	v305 = *result
	tobool633 = (v305 & 1) != 0
	*retval = tobool633
	goto _return

sw_bb634:
	*result = 1
	v306 = *lexer_addr
	result_symbol635 = &v306.F1
	*result_symbol635 = 10
	v307 = *lexer_addr
	mark_end636 = &v307.F3
	v308 = *mark_end636
	v309 = *lexer_addr
	v308(v309)
	v310 = *lookahead
	cmp637 = v310 == 84
	if cmp637 {
		goto if_then642
	} else {
		goto lor_lhs_false639
	}

lor_lhs_false639:
	v311 = *lookahead
	cmp640 = v311 == 116
	if cmp640 {
		goto if_then642
	} else {
		goto if_end643
	}

if_then642:
	*state_addr = 31
	goto next_state

if_end643:
	v312 = *lookahead
	call644 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v312)
	if call644 {
		goto if_then652
	} else {
		goto lor_lhs_false646
	}

lor_lhs_false646:
	v313 = *lookahead
	cmp647 = 48 <= v313
	if cmp647 {
		goto land_lhs_true649
	} else {
		goto if_end653
	}

land_lhs_true649:
	v314 = *lookahead
	cmp650 = v314 <= 57
	if cmp650 {
		goto if_then652
	} else {
		goto if_end653
	}

if_then652:
	*state_addr = 34
	goto next_state

if_end653:
	v315 = *result
	tobool654 = (v315 & 1) != 0
	*retval = tobool654
	goto _return

sw_bb655:
	*result = 1
	v316 = *lexer_addr
	result_symbol656 = &v316.F1
	*result_symbol656 = 10
	v317 = *lexer_addr
	mark_end657 = &v317.F3
	v318 = *mark_end657
	v319 = *lexer_addr
	v318(v319)
	v320 = *lookahead
	call658 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v320)
	if call658 {
		goto if_then666
	} else {
		goto lor_lhs_false660
	}

lor_lhs_false660:
	v321 = *lookahead
	cmp661 = 48 <= v321
	if cmp661 {
		goto land_lhs_true663
	} else {
		goto if_end667
	}

land_lhs_true663:
	v322 = *lookahead
	cmp664 = v322 <= 57
	if cmp664 {
		goto if_then666
	} else {
		goto if_end667
	}

if_then666:
	*state_addr = 34
	goto next_state

if_end667:
	v323 = *result
	tobool668 = (v323 & 1) != 0
	*retval = tobool668
	goto _return

sw_bb669:
	*result = 1
	v324 = *lexer_addr
	result_symbol670 = &v324.F1
	*result_symbol670 = 11
	v325 = *lexer_addr
	mark_end671 = &v325.F3
	v326 = *mark_end671
	v327 = *lexer_addr
	v326(v327)
	v328 = *result
	tobool672 = (v328 & 1) != 0
	*retval = tobool672
	goto _return

sw_bb673:
	*result = 1
	v329 = *lexer_addr
	result_symbol674 = &v329.F1
	*result_symbol674 = 12
	v330 = *lexer_addr
	mark_end675 = &v330.F3
	v331 = *mark_end675
	v332 = *lexer_addr
	v331(v332)
	v333 = *lookahead
	cmp676 = v333 != 0
	if cmp676 {
		goto land_lhs_true678
	} else {
		goto if_end694
	}

land_lhs_true678:
	v334 = *lookahead
	cmp679 = v334 < 9
	if cmp679 {
		goto land_lhs_true684
	} else {
		goto lor_lhs_false681
	}

lor_lhs_false681:
	v335 = *lookahead
	cmp682 = 13 < v335
	if cmp682 {
		goto land_lhs_true684
	} else {
		goto if_end694
	}

land_lhs_true684:
	v336 = *lookahead
	cmp685 = v336 != 32
	if cmp685 {
		goto land_lhs_true687
	} else {
		goto if_end694
	}

land_lhs_true687:
	v337 = *lookahead
	cmp688 = v337 != 44
	if cmp688 {
		goto land_lhs_true690
	} else {
		goto if_end694
	}

land_lhs_true690:
	v338 = *lookahead
	cmp691 = v338 != 125
	if cmp691 {
		goto if_then693
	} else {
		goto if_end694
	}

if_then693:
	*state_addr = 36
	goto next_state

if_end694:
	v339 = *result
	tobool695 = (v339 & 1) != 0
	*retval = tobool695
	goto _return

sw_bb696:
	*result = 1
	v340 = *lexer_addr
	result_symbol697 = &v340.F1
	*result_symbol697 = 13
	v341 = *lexer_addr
	mark_end698 = &v341.F3
	v342 = *mark_end698
	v343 = *lexer_addr
	v342(v343)
	v344 = *lookahead
	cmp699 = v344 != 0
	if cmp699 {
		goto land_lhs_true701
	} else {
		goto if_end714
	}

land_lhs_true701:
	v345 = *lookahead
	cmp702 = v345 < 9
	if cmp702 {
		goto land_lhs_true707
	} else {
		goto lor_lhs_false704
	}

lor_lhs_false704:
	v346 = *lookahead
	cmp705 = 13 < v346
	if cmp705 {
		goto land_lhs_true707
	} else {
		goto if_end714
	}

land_lhs_true707:
	v347 = *lookahead
	cmp708 = v347 != 32
	if cmp708 {
		goto land_lhs_true710
	} else {
		goto if_end714
	}

land_lhs_true710:
	v348 = *lookahead
	cmp711 = v348 != 44
	if cmp711 {
		goto if_then713
	} else {
		goto if_end714
	}

if_then713:
	*state_addr = 37
	goto next_state

if_end714:
	v349 = *result
	tobool715 = (v349 & 1) != 0
	*retval = tobool715
	goto _return

sw_bb716:
	*result = 1
	v350 = *lexer_addr
	result_symbol717 = &v350.F1
	*result_symbol717 = 14
	v351 = *lexer_addr
	mark_end718 = &v351.F3
	v352 = *mark_end718
	v353 = *lexer_addr
	v352(v353)
	v354 = *lookahead
	call719 = set_contains(&sym_entry_type_character_set_1[int64(0)], 9, v354)
	if call719 {
		goto if_then727
	} else {
		goto lor_lhs_false721
	}

lor_lhs_false721:
	v355 = *lookahead
	cmp722 = 48 <= v355
	if cmp722 {
		goto land_lhs_true724
	} else {
		goto if_end728
	}

land_lhs_true724:
	v356 = *lookahead
	cmp725 = v356 <= 57
	if cmp725 {
		goto if_then727
	} else {
		goto if_end728
	}

if_then727:
	*state_addr = 38
	goto next_state

if_end728:
	v357 = *result
	tobool729 = (v357 & 1) != 0
	*retval = tobool729
	goto _return

sw_bb730:
	*result = 1
	v358 = *lexer_addr
	result_symbol731 = &v358.F1
	*result_symbol731 = 15
	v359 = *lexer_addr
	mark_end732 = &v359.F3
	v360 = *mark_end732
	v361 = *lexer_addr
	v360(v361)
	v362 = *result
	tobool733 = (v362 & 1) != 0
	*retval = tobool733
	goto _return

sw_bb734:
	*result = 1
	v363 = *lexer_addr
	result_symbol735 = &v363.F1
	*result_symbol735 = 16
	v364 = *lexer_addr
	mark_end736 = &v364.F3
	v365 = *mark_end736
	v366 = *lexer_addr
	v365(v366)
	v367 = *result
	tobool737 = (v367 & 1) != 0
	*retval = tobool737
	goto _return

sw_bb738:
	*result = 1
	v368 = *lexer_addr
	result_symbol739 = &v368.F1
	*result_symbol739 = 17
	v369 = *lexer_addr
	mark_end740 = &v369.F3
	v370 = *mark_end740
	v371 = *lexer_addr
	v370(v371)
	v372 = *lookahead
	cmp741 = 48 <= v372
	if cmp741 {
		goto land_lhs_true743
	} else {
		goto if_end747
	}

land_lhs_true743:
	v373 = *lookahead
	cmp744 = v373 <= 57
	if cmp744 {
		goto if_then746
	} else {
		goto if_end747
	}

if_then746:
	*state_addr = 41
	goto next_state

if_end747:
	v374 = *result
	tobool748 = (v374 & 1) != 0
	*retval = tobool748
	goto _return

sw_bb749:
	*result = 1
	v375 = *lexer_addr
	result_symbol750 = &v375.F1
	*result_symbol750 = 18
	v376 = *lexer_addr
	mark_end751 = &v376.F3
	v377 = *mark_end751
	v378 = *lexer_addr
	v377(v378)
	v379 = *lookahead
	cmp752 = v379 != 0
	if cmp752 {
		goto land_lhs_true754
	} else {
		goto if_end773
	}

land_lhs_true754:
	v380 = *lookahead
	cmp755 = v380 < 9
	if cmp755 {
		goto land_lhs_true760
	} else {
		goto lor_lhs_false757
	}

lor_lhs_false757:
	v381 = *lookahead
	cmp758 = 13 < v381
	if cmp758 {
		goto land_lhs_true760
	} else {
		goto if_end773
	}

land_lhs_true760:
	v382 = *lookahead
	cmp761 = v382 != 32
	if cmp761 {
		goto land_lhs_true763
	} else {
		goto if_end773
	}

land_lhs_true763:
	v383 = *lookahead
	cmp764 = v383 != 92
	if cmp764 {
		goto land_lhs_true766
	} else {
		goto if_end773
	}

land_lhs_true766:
	v384 = *lookahead
	cmp767 = v384 != 123
	if cmp767 {
		goto land_lhs_true769
	} else {
		goto if_end773
	}

land_lhs_true769:
	v385 = *lookahead
	cmp770 = v385 != 125
	if cmp770 {
		goto if_then772
	} else {
		goto if_end773
	}

if_then772:
	*state_addr = 42
	goto next_state

if_end773:
	v386 = *result
	tobool774 = (v386 & 1) != 0
	*retval = tobool774
	goto _return

sw_bb775:
	*result = 1
	v387 = *lexer_addr
	result_symbol776 = &v387.F1
	*result_symbol776 = 19
	v388 = *lexer_addr
	mark_end777 = &v388.F3
	v389 = *mark_end777
	v390 = *lexer_addr
	v389(v390)
	v391 = *lookahead
	cmp778 = v391 != 0
	if cmp778 {
		goto land_lhs_true780
	} else {
		goto if_end802
	}

land_lhs_true780:
	v392 = *lookahead
	cmp781 = v392 < 9
	if cmp781 {
		goto land_lhs_true786
	} else {
		goto lor_lhs_false783
	}

lor_lhs_false783:
	v393 = *lookahead
	cmp784 = 13 < v393
	if cmp784 {
		goto land_lhs_true786
	} else {
		goto if_end802
	}

land_lhs_true786:
	v394 = *lookahead
	cmp787 = v394 != 32
	if cmp787 {
		goto land_lhs_true789
	} else {
		goto if_end802
	}

land_lhs_true789:
	v395 = *lookahead
	cmp790 = v395 != 34
	if cmp790 {
		goto land_lhs_true792
	} else {
		goto if_end802
	}

land_lhs_true792:
	v396 = *lookahead
	cmp793 = v396 != 92
	if cmp793 {
		goto land_lhs_true795
	} else {
		goto if_end802
	}

land_lhs_true795:
	v397 = *lookahead
	cmp796 = v397 != 123
	if cmp796 {
		goto land_lhs_true798
	} else {
		goto if_end802
	}

land_lhs_true798:
	v398 = *lookahead
	cmp799 = v398 != 125
	if cmp799 {
		goto if_then801
	} else {
		goto if_end802
	}

if_then801:
	*state_addr = 43
	goto next_state

if_end802:
	v399 = *result
	tobool803 = (v399 & 1) != 0
	*retval = tobool803
	goto _return

sw_bb804:
	*result = 1
	v400 = *lexer_addr
	result_symbol805 = &v400.F1
	*result_symbol805 = 20
	v401 = *lexer_addr
	mark_end806 = &v401.F3
	v402 = *mark_end806
	v403 = *lexer_addr
	v402(v403)
	v404 = *result
	tobool807 = (v404 & 1) != 0
	*retval = tobool807
	goto _return

sw_bb808:
	*result = 1
	v405 = *lexer_addr
	result_symbol809 = &v405.F1
	*result_symbol809 = 20
	v406 = *lexer_addr
	mark_end810 = &v406.F3
	v407 = *mark_end810
	v408 = *lexer_addr
	v407(v408)
	v409 = *lookahead
	cmp811 = v409 == 42
	if cmp811 {
		goto if_then813
	} else {
		goto if_end814
	}

if_then813:
	*state_addr = 44
	goto next_state

if_end814:
	v410 = *lookahead
	cmp815 = 64 <= v410
	if cmp815 {
		goto land_lhs_true817
	} else {
		goto lor_lhs_false820
	}

land_lhs_true817:
	v411 = *lookahead
	cmp818 = v411 <= 90
	if cmp818 {
		goto if_then826
	} else {
		goto lor_lhs_false820
	}

lor_lhs_false820:
	v412 = *lookahead
	cmp821 = 97 <= v412
	if cmp821 {
		goto land_lhs_true823
	} else {
		goto if_end827
	}

land_lhs_true823:
	v413 = *lookahead
	cmp824 = v413 <= 122
	if cmp824 {
		goto if_then826
	} else {
		goto if_end827
	}

if_then826:
	*state_addr = 45
	goto next_state

if_end827:
	v414 = *result
	tobool828 = (v414 & 1) != 0
	*retval = tobool828
	goto _return

sw_bb829:
	*result = 1
	v415 = *lexer_addr
	result_symbol830 = &v415.F1
	*result_symbol830 = 20
	v416 = *lexer_addr
	mark_end831 = &v416.F3
	v417 = *mark_end831
	v418 = *lexer_addr
	v417(v418)
	v419 = *lookahead
	cmp832 = 64 <= v419
	if cmp832 {
		goto land_lhs_true834
	} else {
		goto lor_lhs_false837
	}

land_lhs_true834:
	v420 = *lookahead
	cmp835 = v420 <= 90
	if cmp835 {
		goto if_then843
	} else {
		goto lor_lhs_false837
	}

lor_lhs_false837:
	v421 = *lookahead
	cmp838 = 97 <= v421
	if cmp838 {
		goto land_lhs_true840
	} else {
		goto if_end844
	}

land_lhs_true840:
	v422 = *lookahead
	cmp841 = v422 <= 122
	if cmp841 {
		goto if_then843
	} else {
		goto if_end844
	}

if_then843:
	*state_addr = 45
	goto next_state

if_end844:
	v423 = *lookahead
	cmp845 = v423 != 0
	if cmp845 {
		goto land_lhs_true847
	} else {
		goto if_end854
	}

land_lhs_true847:
	v424 = *lookahead
	cmp848 = v424 != 10
	if cmp848 {
		goto land_lhs_true850
	} else {
		goto if_end854
	}

land_lhs_true850:
	v425 = *lookahead
	cmp851 = v425 != 13
	if cmp851 {
		goto if_then853
	} else {
		goto if_end854
	}

if_then853:
	*state_addr = 44
	goto next_state

if_end854:
	v426 = *result
	tobool855 = (v426 & 1) != 0
	*retval = tobool855
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v427 = *retval
	return v427
}
func set_contains(ranges *TSCharacterRange, len int32, lookahead int32) bool {
	var v6, arrayidx, v9, v12, v15, v20, arrayidx10, v23, v26 *TSCharacterRange
	var ranges_addr, _range, range8 **TSCharacterRange
	var cmp, cmp1, cmp2, cmp4, cmp12, cmp14, v28, v29 bool
	var retval *bool
	var v0, v1, sub, v2, v3, div, v4, v5, add, v7, v8, v10, v11, v13, v14, v16, v17, v18, v19, sub7, v21, v22, v24, v25, v27 int32
	var len_addr, lookahead_addr, index, size, half_size, mid_index, start, end, end3, start11, end13 *int32
	var idxprom, idxprom9 int64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, ranges_addr, len_addr, lookahead_addr, index, size, half_size, mid_index, _range, range8, v0, v1, sub, v2, cmp, v3, div, v4, v5, add, v6, v7, idxprom, arrayidx, v8, v9, start, v10, cmp1, v11, v12, end, v13, cmp2, v14, v15, end3, v16, cmp4, v17, v18, v19, sub7, v20, v21, idxprom9, arrayidx10, v22, v23, start11, v24, cmp12, v25, v26, end13, v27, cmp14, v28, v29

	retval = new(bool)
	ranges_addr = new(*TSCharacterRange)
	len_addr = new(int32)
	lookahead_addr = new(int32)
	index = new(int32)
	size = new(int32)
	half_size = new(int32)
	mid_index = new(int32)
	_range = new(*TSCharacterRange)
	range8 = new(*TSCharacterRange)
	*ranges_addr = ranges
	*len_addr = len
	*lookahead_addr = lookahead
	*index = 0
	v0 = *len_addr
	v1 = *index
	sub = v0 - v1
	*size = sub
	goto while_cond

while_cond:
	v2 = *size
	cmp = uint32(v2) > 1
	if cmp {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v3 = *size
	div = int32(uint32(v3) / 2)
	*half_size = div
	v4 = *index
	v5 = *half_size
	add = v4 + v5
	*mid_index = add
	v6 = *ranges_addr
	v7 = *mid_index
	idxprom = int64(uint64(uint32(v7)))
	arrayidx = libc.AddPointer(v6, int(idxprom))
	*_range = arrayidx
	v8 = *lookahead_addr
	v9 = *_range
	start = &v9.F0
	v10 = *start
	cmp1 = v8 >= v10
	if cmp1 {
		goto land_lhs_true
	} else {
		goto if_else
	}

land_lhs_true:
	v11 = *lookahead_addr
	v12 = *_range
	end = &v12.F1
	v13 = *end
	cmp2 = v11 <= v13
	if cmp2 {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	*retval = true
	goto _return

if_else:
	v14 = *lookahead_addr
	v15 = *_range
	end3 = &v15.F1
	v16 = *end3
	cmp4 = v14 > v16
	if cmp4 {
		goto if_then5
	} else {
		goto if_end
	}

if_then5:
	v17 = *mid_index
	*index = v17
	goto if_end

if_end:
	goto if_end6

if_end6:
	v18 = *half_size
	v19 = *size
	sub7 = v19 - v18
	*size = sub7
	goto while_cond

while_end:
	v20 = *ranges_addr
	v21 = *index
	idxprom9 = int64(uint64(uint32(v21)))
	arrayidx10 = libc.AddPointer(v20, int(idxprom9))
	*range8 = arrayidx10
	v22 = *lookahead_addr
	v23 = *range8
	start11 = &v23.F0
	v24 = *start11
	cmp12 = v22 >= v24
	if cmp12 {
		goto land_rhs
	} else {
		v28 = false
		goto land_end
	}

land_rhs:
	v25 = *lookahead_addr
	v26 = *range8
	end13 = &v26.F1
	v27 = *end13
	cmp14 = v25 <= v27
	v28 = cmp14
	goto land_end

land_end:
	*retval = v28
	goto _return

_return:
	v29 = *retval
	return v29
}
