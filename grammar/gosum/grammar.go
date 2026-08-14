package grammar_gosum

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type TSLanguageMetadata struct {
	F0 byte
	F1 byte
	F2 byte
}
type TSMapSlice struct {
	F0 int16
	F1 int16
}
type TSFieldMapEntry struct {
	F0 int16
	F1 byte
	F2 byte
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

var tree_sitter_gosum_language struct {
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
	F30 unsafe.Pointer
	F31 unsafe.Pointer
	F32 int16
	F33 [2]byte
	F34 int32
	F35 unsafe.Pointer
	F36 unsafe.Pointer
	F37 unsafe.Pointer
	F38 TSLanguageMetadata
	F39 [5]byte
}
var ts_parse_table [2][24]int16 = [2][24]int16{[24]int16{1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0}, [24]int16{0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 38, 5, 0, 0, 5}}
var ts_small_parse_table [314]int16 = [314]int16{2, 7, 1, 17, 5, 5, 7, 8, 9, 10, 11, 4, 11, 1, 5, 13, 1, 6, 15, 1, 12, 9, 2, 1, 14, 4, 19, 1, 5, 21, 1, 6, 23, 1, 12, 17, 2, 1, 14, 3, 3, 1, 3, 25, 1, 0, 8, 2, 20, 23, 3, 29, 1, 6, 31, 1, 12, 27, 2, 1, 14, 3, 35, 1, 6, 37, 1, 12, 33, 2, 1, 14, 3, 39, 1, 0, 41, 1, 3, 8, 2, 20, 23, 3, 46, 1, 6, 48, 1, 12, 44, 2, 1, 14, 3, 52, 1, 6, 54, 1, 12, 50, 2, 1, 14, 2, 58, 1, 12, 56, 2, 1, 14, 3, 60, 1, 1, 62, 1, 14, 25, 1, 22, 2, 66, 1, 12, 64, 2, 1, 14, 2, 70, 1, 12, 68, 2, 1, 14, 2, 74, 1, 12, 72, 2, 1, 14, 2, 78, 1, 12, 76, 2, 1, 14, 1, 80, 2, 1, 14, 1, 82, 2, 1, 14, 1, 84, 2, 1, 14, 2, 86, 1, 5, 88, 1, 16, 2, 90, 1, 4, 12, 1, 21, 1, 92, 2, 1, 14, 1, 94, 2, 1, 14, 1, 96, 2, 1, 14, 1, 98, 2, 0, 3, 1, 100, 2, 1, 14, 2, 62, 1, 14, 31, 1, 22, 1, 102, 2, 1, 14, 1, 104, 2, 1, 14, 1, 106, 2, 1, 14, 1, 108, 2, 0, 3, 1, 110, 2, 0, 3, 1, 112, 2, 1, 14, 1, 114, 1, 18, 1, 116, 1, 16, 1, 118, 1, 18, 1, 120, 1, 16, 1, 122, 1, 0, 1, 124, 1, 2, 1, 126, 1, 15, 1, 128, 1, 18, 1, 130, 1, 5, 1, 132, 1, 17, 1, 134, 1, 18, 1, 136, 1, 18, 1, 138, 1, 16, 1, 140, 1, 13, 1, 142, 1, 17, 1, 144, 1, 16, 1, 146, 1, 5}
var ts_small_parse_table_map [49]int32 = [49]int32{0, 11, 25, 39, 50, 61, 72, 83, 94, 105, 113, 123, 131, 139, 147, 155, 160, 165, 170, 177, 184, 189, 194, 199, 204, 209, 216, 221, 226, 231, 236, 241, 246, 250, 254, 258, 262, 266, 270, 274, 278, 282, 286, 290, 294, 298, 302, 306, 310}
var ts_symbol_names [24]unsafe.Pointer = [24]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_2), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24)}
var ts_field_names [6]unsafe.Pointer = [6]unsafe.Pointer{nil, libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29)}
var ts_field_map_slices [12]TSMapSlice = [12]TSMapSlice{TSMapSlice{}, TSMapSlice{0, 3}, TSMapSlice{3, 4}, TSMapSlice{7, 5}, TSMapSlice{12, 6}, TSMapSlice{18, 5}, TSMapSlice{23, 7}, TSMapSlice{30, 6}, TSMapSlice{36, 8}, TSMapSlice{44, 7}, TSMapSlice{51, 8}, TSMapSlice{59, 9}}
var ts_field_map_entries [68]TSFieldMapEntry = [68]TSFieldMapEntry{TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{5, 9, 0}, TSFieldMapEntry{1, 9, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{5, 9, 0}, TSFieldMapEntry{5, 10, 0}, TSFieldMapEntry{1, 10, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{5, 9, 0}, TSFieldMapEntry{5, 10, 0}, TSFieldMapEntry{5, 11, 0}, TSFieldMapEntry{1, 11, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{5, 9, 0}, TSFieldMapEntry{1, 12, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{5, 9, 0}, TSFieldMapEntry{5, 10, 0}, TSFieldMapEntry{1, 13, 0}, TSFieldMapEntry{2, 1, 0}, TSFieldMapEntry{3, 3, 0}, TSFieldMapEntry{4, 5, 0}, TSFieldMapEntry{5, 7, 0}, TSFieldMapEntry{5, 8, 0}, TSFieldMapEntry{5, 9, 0}, TSFieldMapEntry{5, 10, 0}, TSFieldMapEntry{5, 11, 0}}
var ts_symbol_metadata [24]TSSymbolMetadata = [24]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}}
var ts_symbol_map [24]int16 = [24]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_lex_modes [51]TSLexMode = [51]TSLexMode{TSLexMode{}, TSLexMode{35, 0}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{35, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{35, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{1, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{35, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{35, 0}, TSLexMode{35, 0}, TSLexMode{}, TSLexMode{33, 0}, TSLexMode{1, 0}, TSLexMode{33, 0}, TSLexMode{1, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{34, 0}, TSLexMode{33, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{33, 0}, TSLexMode{33, 0}, TSLexMode{1, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{1, 0}, TSLexMode{}}
var ts_primary_state_ids [51]int16 = [51]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
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
	F6 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F10 TSParseActionEntry
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
	F18 TSParseActionEntry
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
	F26 TSParseActionEntry
	F27 struct {
		F0 anon_2
		F1 [6]byte
	}
	F28 TSParseActionEntry
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
	F34 TSParseActionEntry
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
	F40 TSParseActionEntry
	F41 struct {
		F0 anon_2
		F1 [6]byte
	}
	F42 TSParseActionEntry
	F43 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F44 struct {
		F0 anon_2
		F1 [6]byte
	}
	F45 TSParseActionEntry
	F46 struct {
		F0 anon_2
		F1 [6]byte
	}
	F47 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F48 struct {
		F0 anon_2
		F1 [6]byte
	}
	F49 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F50 struct {
		F0 anon_2
		F1 [6]byte
	}
	F51 TSParseActionEntry
	F52 struct {
		F0 anon_2
		F1 [6]byte
	}
	F53 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F54 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F57 TSParseActionEntry
	F58 struct {
		F0 anon_2
		F1 [6]byte
	}
	F59 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F60 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F65 TSParseActionEntry
	F66 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F69 TSParseActionEntry
	F70 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F73 TSParseActionEntry
	F74 struct {
		F0 anon_2
		F1 [6]byte
	}
	F75 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F76 struct {
		F0 anon_2
		F1 [6]byte
	}
	F77 TSParseActionEntry
	F78 struct {
		F0 anon_2
		F1 [6]byte
	}
	F79 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F80 struct {
		F0 anon_2
		F1 [6]byte
	}
	F81 TSParseActionEntry
	F82 struct {
		F0 anon_2
		F1 [6]byte
	}
	F83 TSParseActionEntry
	F84 struct {
		F0 anon_2
		F1 [6]byte
	}
	F85 TSParseActionEntry
	F86 struct {
		F0 anon_2
		F1 [6]byte
	}
	F87 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F88 struct {
		F0 anon_2
		F1 [6]byte
	}
	F89 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F90 struct {
		F0 anon_2
		F1 [6]byte
	}
	F91 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F92 struct {
		F0 anon_2
		F1 [6]byte
	}
	F93 TSParseActionEntry
	F94 struct {
		F0 anon_2
		F1 [6]byte
	}
	F95 TSParseActionEntry
	F96 struct {
		F0 anon_2
		F1 [6]byte
	}
	F97 TSParseActionEntry
	F98 struct {
		F0 anon_2
		F1 [6]byte
	}
	F99  TSParseActionEntry
	F100 struct {
		F0 anon_2
		F1 [6]byte
	}
	F101 TSParseActionEntry
	F102 struct {
		F0 anon_2
		F1 [6]byte
	}
	F103 TSParseActionEntry
	F104 struct {
		F0 anon_2
		F1 [6]byte
	}
	F105 TSParseActionEntry
	F106 struct {
		F0 anon_2
		F1 [6]byte
	}
	F107 TSParseActionEntry
	F108 struct {
		F0 anon_2
		F1 [6]byte
	}
	F109 TSParseActionEntry
	F110 struct {
		F0 anon_2
		F1 [6]byte
	}
	F111 TSParseActionEntry
	F112 struct {
		F0 anon_2
		F1 [6]byte
	}
	F113 TSParseActionEntry
	F114 struct {
		F0 anon_2
		F1 [6]byte
	}
	F115 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F116 struct {
		F0 anon_2
		F1 [6]byte
	}
	F117 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F118 struct {
		F0 anon_2
		F1 [6]byte
	}
	F119 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F120 struct {
		F0 anon_2
		F1 [6]byte
	}
	F121 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F122 struct {
		F0 anon_2
		F1 [6]byte
	}
	F123 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F124 struct {
		F0 anon_2
		F1 [6]byte
	}
	F125 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F126 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F147 struct {
		F0 struct {
			F0 struct {
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
	F6 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F10 TSParseActionEntry
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
	F18 TSParseActionEntry
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
	F26 TSParseActionEntry
	F27 struct {
		F0 anon_2
		F1 [6]byte
	}
	F28 TSParseActionEntry
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
	F34 TSParseActionEntry
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
	F40 TSParseActionEntry
	F41 struct {
		F0 anon_2
		F1 [6]byte
	}
	F42 TSParseActionEntry
	F43 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F44 struct {
		F0 anon_2
		F1 [6]byte
	}
	F45 TSParseActionEntry
	F46 struct {
		F0 anon_2
		F1 [6]byte
	}
	F47 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F48 struct {
		F0 anon_2
		F1 [6]byte
	}
	F49 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F50 struct {
		F0 anon_2
		F1 [6]byte
	}
	F51 TSParseActionEntry
	F52 struct {
		F0 anon_2
		F1 [6]byte
	}
	F53 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F54 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F57 TSParseActionEntry
	F58 struct {
		F0 anon_2
		F1 [6]byte
	}
	F59 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F60 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F65 TSParseActionEntry
	F66 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F69 TSParseActionEntry
	F70 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F73 TSParseActionEntry
	F74 struct {
		F0 anon_2
		F1 [6]byte
	}
	F75 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F76 struct {
		F0 anon_2
		F1 [6]byte
	}
	F77 TSParseActionEntry
	F78 struct {
		F0 anon_2
		F1 [6]byte
	}
	F79 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F80 struct {
		F0 anon_2
		F1 [6]byte
	}
	F81 TSParseActionEntry
	F82 struct {
		F0 anon_2
		F1 [6]byte
	}
	F83 TSParseActionEntry
	F84 struct {
		F0 anon_2
		F1 [6]byte
	}
	F85 TSParseActionEntry
	F86 struct {
		F0 anon_2
		F1 [6]byte
	}
	F87 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F88 struct {
		F0 anon_2
		F1 [6]byte
	}
	F89 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F90 struct {
		F0 anon_2
		F1 [6]byte
	}
	F91 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F92 struct {
		F0 anon_2
		F1 [6]byte
	}
	F93 TSParseActionEntry
	F94 struct {
		F0 anon_2
		F1 [6]byte
	}
	F95 TSParseActionEntry
	F96 struct {
		F0 anon_2
		F1 [6]byte
	}
	F97 TSParseActionEntry
	F98 struct {
		F0 anon_2
		F1 [6]byte
	}
	F99  TSParseActionEntry
	F100 struct {
		F0 anon_2
		F1 [6]byte
	}
	F101 TSParseActionEntry
	F102 struct {
		F0 anon_2
		F1 [6]byte
	}
	F103 TSParseActionEntry
	F104 struct {
		F0 anon_2
		F1 [6]byte
	}
	F105 TSParseActionEntry
	F106 struct {
		F0 anon_2
		F1 [6]byte
	}
	F107 TSParseActionEntry
	F108 struct {
		F0 anon_2
		F1 [6]byte
	}
	F109 TSParseActionEntry
	F110 struct {
		F0 anon_2
		F1 [6]byte
	}
	F111 TSParseActionEntry
	F112 struct {
		F0 anon_2
		F1 [6]byte
	}
	F113 TSParseActionEntry
	F114 struct {
		F0 anon_2
		F1 [6]byte
	}
	F115 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F116 struct {
		F0 anon_2
		F1 [6]byte
	}
	F117 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F118 struct {
		F0 anon_2
		F1 [6]byte
	}
	F119 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F120 struct {
		F0 anon_2
		F1 [6]byte
	}
	F121 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F122 struct {
		F0 anon_2
		F1 [6]byte
	}
	F123 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F124 struct {
		F0 anon_2
		F1 [6]byte
	}
	F125 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F126 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F147 struct {
		F0 struct {
			F0 struct {
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
}{0, 0, 20, 0, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 10, 21, 0, 4}}}, struct {
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
}{0, 0, 48, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 34, 0, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 9, 21, 0, 3}}}, struct {
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
}{0, 0, 43, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 23, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 19, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 6, 21, 0, 1}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 12, 21, 0, 8}}}, struct {
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
}{0, 0, 29, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 23, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 23, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 21, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 8, 21, 0, 2}}}, struct {
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
}{0, 0, 36, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 22, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 11, 21, 0, 6}}}, struct {
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
}{0, 0, 41, 0, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 10, 21, 0, 5}}}, struct {
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
}{0, 0, 26, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 39, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 47, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 11, 21, 0, 7}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 12, 21, 0, 9}}}, struct {
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
}{0, 0, 30, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 13, 21, 0, 10}}}, struct {
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
}{0, 0, 33, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 14, 21, 0, 11}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 12, 21, 0, 6}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 15, 21, 0, 11}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 7, 21, 0, 1}}}, struct {
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
}{0, 0, 46, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 35, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 9, 21, 0, 2}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 10, 21, 0, 3}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 11, 21, 0, 4}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 11, 21, 0, 5}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 12, 21, 0, 7}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 13, 21, 0, 8}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 13, 21, 0, 9}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 14, 21, 0, 10}}}, struct {
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
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{0, 0, 42, 0, 0}, [2]byte{}}}, struct {
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
}{2, [7]byte{}}}, struct {
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
}{0, 0, 32, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 15, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 40, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 37, 0, 0}, [2]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_2 [2]byte = [2]byte{47, 0}
var _str_3 [7]byte = [7]byte{103, 111, 46, 109, 111, 100, 0}
var _str_4 [12]byte = [12]byte{109, 111, 100, 117, 108, 101, 95, 112, 97, 116, 104, 0}
var _str_5 [15]byte = [15]byte{109, 111, 100, 117, 108, 101, 95, 118, 101, 114, 115, 105, 111, 110, 0}
var _str_6 [2]byte = [2]byte{46, 0}
var _str_7 [2]byte = [2]byte{45, 0}
var _str_8 [6]byte = [6]byte{97, 108, 112, 104, 97, 0}
var _str_9 [5]byte = [5]byte{98, 101, 116, 97, 0}
var _str_10 [4]byte = [4]byte{100, 101, 118, 0}
var _str_11 [4]byte = [4]byte{112, 114, 101, 0}
var _str_12 [3]byte = [3]byte{114, 99, 0}
var _str_13 [14]byte = [14]byte{43, 105, 110, 99, 111, 109, 112, 97, 116, 105, 98, 108, 101, 0}
var _str_14 [2]byte = [2]byte{58, 0}
var _str_15 [13]byte = [13]byte{104, 97, 115, 104, 95, 118, 101, 114, 115, 105, 111, 110, 0}
var _str_16 [5]byte = [5]byte{104, 97, 115, 104, 0}
var _str_17 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}
var _str_18 [20]byte = [20]byte{110, 117, 109, 98, 101, 114, 95, 119, 105, 116, 104, 95, 100, 101, 99, 105, 109, 97, 108, 0}
var _str_19 [11]byte = [11]byte{104, 101, 120, 95, 110, 117, 109, 98, 101, 114, 0}
var _str_20 [18]byte = [18]byte{99, 104, 101, 99, 107, 115, 117, 109, 95, 100, 97, 116, 97, 98, 97, 115, 101, 0}
var _str_21 [9]byte = [9]byte{99, 104, 101, 99, 107, 115, 117, 109, 0}
var _str_22 [8]byte = [8]byte{118, 101, 114, 115, 105, 111, 110, 0}
var _str_23 [15]byte = [15]byte{99, 104, 101, 99, 107, 115, 117, 109, 95, 118, 97, 108, 117, 101, 0}
var _str_24 [26]byte = [26]byte{99, 104, 101, 99, 107, 115, 117, 109, 95, 100, 97, 116, 97, 98, 97, 115, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_25 [6]byte = [6]byte{98, 117, 105, 108, 100, 0}
var _str_26 [6]byte = [6]byte{109, 97, 106, 111, 114, 0}
var _str_27 [6]byte = [6]byte{109, 105, 110, 111, 114, 0}
var _str_28 [6]byte = [6]byte{112, 97, 116, 99, 104, 0}
var _str_29 [12]byte = [12]byte{112, 114, 101, 95, 114, 101, 108, 101, 97, 115, 101, 0}
var ts_alias_sequences struct {
	F0 [15]int16
	F1 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F2 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F3 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F4 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F5 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F6 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F7 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F8 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F9 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F10 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F11 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
} = struct {
	F0 [15]int16
	F1 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F2 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F3 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F4 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F5 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F6 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F7 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F8 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F9 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F10 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
	F11 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 [9]int16
	}
}{[15]int16{}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}, struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 [9]int16
}{0, 4, 4, 4, 4, 4, [9]int16{}}}
var ts_lex_map [26]int16 = [26]int16{43, 17, 45, 42, 46, 41, 47, 37, 58, 49, 97, 59, 98, 57, 100, 58, 103, 24, 104, 3, 112, 29, 114, 9, 118, 40}

func init() {
	tree_sitter_gosum_language = struct {
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
		F30 unsafe.Pointer
		F31 unsafe.Pointer
		F32 int16
		F33 [2]byte
		F34 int32
		F35 unsafe.Pointer
		F36 unsafe.Pointer
		F37 unsafe.Pointer
		F38 TSLanguageMetadata
		F39 [5]byte
	}{14, 24, 0, 19, 0, 51, 2, 12, 5, 15, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), libc.Ptr(&ts_field_names), libc.Ptr(&ts_field_map_slices), libc.Ptr(&ts_field_map_entries), libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), nil, nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{}, [5]byte{}}
}
func tree_sitter_gosum() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_gosum_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, cmp28, cmp31, cmp35, cmp38, loadedv42, cmp44, cmp48, cmp51, cmp54, cmp58, cmp61, loadedv65, cmp67, loadedv71, cmp73, loadedv77, cmp79, loadedv83, cmp85, loadedv89, cmp91, loadedv95, cmp97, cmp101, cmp105, cmp109, cmp113, cmp117, cmp120, cmp123, cmp127, cmp130, cmp133, loadedv137, cmp139, loadedv143, cmp145, loadedv149, cmp151, loadedv155, cmp157, loadedv161, cmp163, loadedv167, cmp169, loadedv173, cmp175, loadedv179, cmp181, loadedv185, cmp187, loadedv191, cmp193, loadedv197, cmp199, loadedv203, cmp205, loadedv209, cmp211, loadedv215, cmp217, loadedv221, cmp223, loadedv227, cmp229, loadedv233, cmp235, loadedv239, cmp241, loadedv245, cmp247, loadedv251, cmp253, loadedv257, cmp259, loadedv263, cmp265, loadedv269, cmp271, loadedv275, cmp277, loadedv281, cmp283, loadedv287, cmp289, cmp292, cmp295, cmp299, cmp302, cmp305, cmp308, cmp311, cmp314, loadedv318, cmp320, cmp323, cmp326, cmp330, cmp333, cmp336, cmp339, cmp342, cmp345, cmp348, loadedv352, loadedv354, cmp357, cmp360, cmp363, cmp367, cmp370, cmp373, cmp376, cmp379, cmp382, cmp385, loadedv389, loadedv391, loadedv395, loadedv399, cmp403, cmp406, cmp409, cmp412, cmp415, cmp418, cmp421, loadedv425, loadedv429, loadedv433, loadedv437, loadedv441, loadedv445, loadedv449, loadedv453, loadedv457, loadedv461, loadedv465, loadedv469, loadedv473, cmp477, cmp481, cmp484, cmp487, cmp490, cmp493, cmp496, cmp499, loadedv503, cmp507, loadedv511, cmp515, cmp518, cmp522, cmp525, cmp528, cmp531, loadedv535, cmp539, cmp542, loadedv546, cmp550, cmp553, cmp556, loadedv560, cmp564, cmp568, cmp571, cmp574, cmp577, cmp580, cmp583, loadedv587, cmp591, cmp595, cmp598, cmp601, cmp604, cmp607, cmp610, loadedv614, cmp618, cmp622, cmp625, cmp628, cmp631, cmp634, cmp637, loadedv641, cmp645, cmp649, cmp652, cmp655, cmp658, cmp661, cmp664, loadedv668, cmp672, cmp676, cmp679, cmp682, cmp685, cmp688, cmp691, loadedv695, cmp699, cmp702, cmp705, cmp708, cmp711, cmp714, loadedv718, v343 bool
	var retval unsafe.Pointer
	var v9, v13, v16 int16
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol393, result_symbol397, result_symbol401, result_symbol427, result_symbol431, result_symbol435, result_symbol439, result_symbol443, result_symbol447, result_symbol451, result_symbol455, result_symbol459, result_symbol463, result_symbol467, result_symbol471, result_symbol475, result_symbol505, result_symbol513, result_symbol537, result_symbol548, result_symbol562, result_symbol589, result_symbol616, result_symbol643, result_symbol670, result_symbol697 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v24, v25, v26, v28, v29, v30, v31, v32, v33, v35, v37, v39, v41, v43, v45, v46, v47, v48, v49, v50, v51, v52, v53, v54, v55, v57, v59, v61, v63, v65, v67, v69, v71, v73, v75, v77, v79, v81, v83, v85, v87, v89, v91, v93, v95, v97, v99, v101, v103, v105, v107, v108, v109, v110, v111, v112, v113, v114, v115, v117, v118, v119, v120, v121, v122, v123, v124, v125, v126, v129, v130, v131, v132, v133, v134, v135, v136, v137, v138, v159, v160, v161, v162, v163, v164, v165, v231, v232, v233, v234, v235, v236, v237, v238, v244, v250, v251, v252, v253, v254, v255, v261, v262, v268, v269, v270, v276, v277, v278, v279, v280, v281, v282, v288, v289, v290, v291, v292, v293, v294, v300, v301, v302, v303, v304, v305, v306, v312, v313, v314, v315, v316, v317, v318, v324, v325, v326, v327, v328, v329, v330, v336, v337, v338, v339, v340, v341 int32
	var lookahead, i, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10 int64
	var v3, storedv, v10, v27, v34, v36, v38, v40, v42, v44, v56, v58, v60, v62, v64, v66, v68, v70, v72, v74, v76, v78, v80, v82, v84, v86, v88, v90, v92, v94, v96, v98, v100, v102, v104, v106, v116, v127, v128, v139, v144, v149, v154, v166, v171, v176, v181, v186, v191, v196, v201, v206, v211, v216, v221, v226, v239, v245, v256, v263, v271, v283, v295, v307, v319, v331, v342 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v140, v141, v142, v143, v145, v146, v147, v148, v150, v151, v152, v153, v155, v156, v157, v158, v167, v168, v169, v170, v172, v173, v174, v175, v177, v178, v179, v180, v182, v183, v184, v185, v187, v188, v189, v190, v192, v193, v194, v195, v197, v198, v199, v200, v202, v203, v204, v205, v207, v208, v209, v210, v212, v213, v214, v215, v217, v218, v219, v220, v222, v223, v224, v225, v227, v228, v229, v230, v240, v241, v242, v243, v246, v247, v248, v249, v257, v258, v259, v260, v264, v265, v266, v267, v272, v273, v274, v275, v284, v285, v286, v287, v296, v297, v298, v299, v308, v309, v310, v311, v320, v321, v322, v323, v332, v333, v334, v335 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end394, mark_end398, mark_end402, mark_end428, mark_end432, mark_end436, mark_end440, mark_end444, mark_end448, mark_end452, mark_end456, mark_end460, mark_end464, mark_end468, mark_end472, mark_end476, mark_end506, mark_end514, mark_end538, mark_end549, mark_end563, mark_end590, mark_end617, mark_end644, mark_end671, mark_end698 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, cmp28, v24, cmp31, v25, cmp35, v26, cmp38, v27, loadedv42, v28, cmp44, v29, cmp48, v30, cmp51, v31, cmp54, v32, cmp58, v33, cmp61, v34, loadedv65, v35, cmp67, v36, loadedv71, v37, cmp73, v38, loadedv77, v39, cmp79, v40, loadedv83, v41, cmp85, v42, loadedv89, v43, cmp91, v44, loadedv95, v45, cmp97, v46, cmp101, v47, cmp105, v48, cmp109, v49, cmp113, v50, cmp117, v51, cmp120, v52, cmp123, v53, cmp127, v54, cmp130, v55, cmp133, v56, loadedv137, v57, cmp139, v58, loadedv143, v59, cmp145, v60, loadedv149, v61, cmp151, v62, loadedv155, v63, cmp157, v64, loadedv161, v65, cmp163, v66, loadedv167, v67, cmp169, v68, loadedv173, v69, cmp175, v70, loadedv179, v71, cmp181, v72, loadedv185, v73, cmp187, v74, loadedv191, v75, cmp193, v76, loadedv197, v77, cmp199, v78, loadedv203, v79, cmp205, v80, loadedv209, v81, cmp211, v82, loadedv215, v83, cmp217, v84, loadedv221, v85, cmp223, v86, loadedv227, v87, cmp229, v88, loadedv233, v89, cmp235, v90, loadedv239, v91, cmp241, v92, loadedv245, v93, cmp247, v94, loadedv251, v95, cmp253, v96, loadedv257, v97, cmp259, v98, loadedv263, v99, cmp265, v100, loadedv269, v101, cmp271, v102, loadedv275, v103, cmp277, v104, loadedv281, v105, cmp283, v106, loadedv287, v107, cmp289, v108, cmp292, v109, cmp295, v110, cmp299, v111, cmp302, v112, cmp305, v113, cmp308, v114, cmp311, v115, cmp314, v116, loadedv318, v117, cmp320, v118, cmp323, v119, cmp326, v120, cmp330, v121, cmp333, v122, cmp336, v123, cmp339, v124, cmp342, v125, cmp345, v126, cmp348, v127, loadedv352, v128, loadedv354, v129, cmp357, v130, cmp360, v131, cmp363, v132, cmp367, v133, cmp370, v134, cmp373, v135, cmp376, v136, cmp379, v137, cmp382, v138, cmp385, v139, loadedv389, v140, result_symbol, v141, mark_end, v142, v143, v144, loadedv391, v145, result_symbol393, v146, mark_end394, v147, v148, v149, loadedv395, v150, result_symbol397, v151, mark_end398, v152, v153, v154, loadedv399, v155, result_symbol401, v156, mark_end402, v157, v158, v159, cmp403, v160, cmp406, v161, cmp409, v162, cmp412, v163, cmp415, v164, cmp418, v165, cmp421, v166, loadedv425, v167, result_symbol427, v168, mark_end428, v169, v170, v171, loadedv429, v172, result_symbol431, v173, mark_end432, v174, v175, v176, loadedv433, v177, result_symbol435, v178, mark_end436, v179, v180, v181, loadedv437, v182, result_symbol439, v183, mark_end440, v184, v185, v186, loadedv441, v187, result_symbol443, v188, mark_end444, v189, v190, v191, loadedv445, v192, result_symbol447, v193, mark_end448, v194, v195, v196, loadedv449, v197, result_symbol451, v198, mark_end452, v199, v200, v201, loadedv453, v202, result_symbol455, v203, mark_end456, v204, v205, v206, loadedv457, v207, result_symbol459, v208, mark_end460, v209, v210, v211, loadedv461, v212, result_symbol463, v213, mark_end464, v214, v215, v216, loadedv465, v217, result_symbol467, v218, mark_end468, v219, v220, v221, loadedv469, v222, result_symbol471, v223, mark_end472, v224, v225, v226, loadedv473, v227, result_symbol475, v228, mark_end476, v229, v230, v231, cmp477, v232, cmp481, v233, cmp484, v234, cmp487, v235, cmp490, v236, cmp493, v237, cmp496, v238, cmp499, v239, loadedv503, v240, result_symbol505, v241, mark_end506, v242, v243, v244, cmp507, v245, loadedv511, v246, result_symbol513, v247, mark_end514, v248, v249, v250, cmp515, v251, cmp518, v252, cmp522, v253, cmp525, v254, cmp528, v255, cmp531, v256, loadedv535, v257, result_symbol537, v258, mark_end538, v259, v260, v261, cmp539, v262, cmp542, v263, loadedv546, v264, result_symbol548, v265, mark_end549, v266, v267, v268, cmp550, v269, cmp553, v270, cmp556, v271, loadedv560, v272, result_symbol562, v273, mark_end563, v274, v275, v276, cmp564, v277, cmp568, v278, cmp571, v279, cmp574, v280, cmp577, v281, cmp580, v282, cmp583, v283, loadedv587, v284, result_symbol589, v285, mark_end590, v286, v287, v288, cmp591, v289, cmp595, v290, cmp598, v291, cmp601, v292, cmp604, v293, cmp607, v294, cmp610, v295, loadedv614, v296, result_symbol616, v297, mark_end617, v298, v299, v300, cmp618, v301, cmp622, v302, cmp625, v303, cmp628, v304, cmp631, v305, cmp634, v306, cmp637, v307, loadedv641, v308, result_symbol643, v309, mark_end644, v310, v311, v312, cmp645, v313, cmp649, v314, cmp652, v315, cmp655, v316, cmp658, v317, cmp661, v318, cmp664, v319, loadedv668, v320, result_symbol670, v321, mark_end671, v322, v323, v324, cmp672, v325, cmp676, v326, cmp679, v327, cmp682, v328, cmp685, v329, cmp688, v330, cmp691, v331, loadedv695, v332, result_symbol697, v333, mark_end698, v334, v335, v336, cmp699, v337, cmp702, v338, cmp705, v339, cmp708, v340, cmp711, v341, cmp714, v342, loadedv718, v343

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
		goto sw_bb43
	case 2:
		goto sw_bb66
	case 3:
		goto sw_bb72
	case 4:
		goto sw_bb78
	case 5:
		goto sw_bb84
	case 6:
		goto sw_bb90
	case 7:
		goto sw_bb96
	case 8:
		goto sw_bb138
	case 9:
		goto sw_bb144
	case 10:
		goto sw_bb150
	case 11:
		goto sw_bb156
	case 12:
		goto sw_bb162
	case 13:
		goto sw_bb168
	case 14:
		goto sw_bb174
	case 15:
		goto sw_bb180
	case 16:
		goto sw_bb186
	case 17:
		goto sw_bb192
	case 18:
		goto sw_bb198
	case 19:
		goto sw_bb204
	case 20:
		goto sw_bb210
	case 21:
		goto sw_bb216
	case 22:
		goto sw_bb222
	case 23:
		goto sw_bb228
	case 24:
		goto sw_bb234
	case 25:
		goto sw_bb240
	case 26:
		goto sw_bb246
	case 27:
		goto sw_bb252
	case 28:
		goto sw_bb258
	case 29:
		goto sw_bb264
	case 30:
		goto sw_bb270
	case 31:
		goto sw_bb276
	case 32:
		goto sw_bb282
	case 33:
		goto sw_bb288
	case 34:
		goto sw_bb319
	case 35:
		goto sw_bb353
	case 36:
		goto sw_bb390
	case 37:
		goto sw_bb392
	case 38:
		goto sw_bb396
	case 39:
		goto sw_bb400
	case 40:
		goto sw_bb426
	case 41:
		goto sw_bb430
	case 42:
		goto sw_bb434
	case 43:
		goto sw_bb438
	case 44:
		goto sw_bb442
	case 45:
		goto sw_bb446
	case 46:
		goto sw_bb450
	case 47:
		goto sw_bb454
	case 48:
		goto sw_bb458
	case 49:
		goto sw_bb462
	case 50:
		goto sw_bb466
	case 51:
		goto sw_bb470
	case 52:
		goto sw_bb474
	case 53:
		goto sw_bb504
	case 54:
		goto sw_bb512
	case 55:
		goto sw_bb536
	case 56:
		goto sw_bb547
	case 57:
		goto sw_bb561
	case 58:
		goto sw_bb588
	case 59:
		goto sw_bb615
	case 60:
		goto sw_bb642
	case 61:
		goto sw_bb669
	case 62:
		goto sw_bb696
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
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v11 = *libc.As[int32](i)
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(26)
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
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end21:
	v21 = *libc.As[int32](lookahead)
	cmp22 = 65 <= v21
	if cmp22 {
		goto land_lhs_true24
	} else {
		goto lor_lhs_false27
	}

land_lhs_true24:
	v22 = *libc.As[int32](lookahead)
	cmp25 = v22 <= 70
	if cmp25 {
		goto if_then33
	} else {
		goto lor_lhs_false27
	}

lor_lhs_false27:
	v23 = *libc.As[int32](lookahead)
	cmp28 = 99 <= v23
	if cmp28 {
		goto land_lhs_true30
	} else {
		goto if_end34
	}

land_lhs_true30:
	v24 = *libc.As[int32](lookahead)
	cmp31 = v24 <= 102
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end34:
	v25 = *libc.As[int32](lookahead)
	cmp35 = 48 <= v25
	if cmp35 {
		goto land_lhs_true37
	} else {
		goto if_end41
	}

land_lhs_true37:
	v26 = *libc.As[int32](lookahead)
	cmp38 = v26 <= 57
	if cmp38 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end41:
	v27 = *libc.As[byte](result)
	loadedv42 = (v27 & 1) != 0
	*libc.As[bool](retval) = loadedv42
	goto _return

sw_bb43:
	v28 = *libc.As[int32](lookahead)
	cmp44 = v28 == 46
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end47:
	v29 = *libc.As[int32](lookahead)
	cmp48 = 9 <= v29
	if cmp48 {
		goto land_lhs_true50
	} else {
		goto lor_lhs_false53
	}

land_lhs_true50:
	v30 = *libc.As[int32](lookahead)
	cmp51 = v30 <= 13
	if cmp51 {
		goto if_then56
	} else {
		goto lor_lhs_false53
	}

lor_lhs_false53:
	v31 = *libc.As[int32](lookahead)
	cmp54 = v31 == 32
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end57:
	v32 = *libc.As[int32](lookahead)
	cmp58 = 48 <= v32
	if cmp58 {
		goto land_lhs_true60
	} else {
		goto if_end64
	}

land_lhs_true60:
	v33 = *libc.As[int32](lookahead)
	cmp61 = v33 <= 57
	if cmp61 {
		goto if_then63
	} else {
		goto if_end64
	}

if_then63:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end64:
	v34 = *libc.As[byte](result)
	loadedv65 = (v34 & 1) != 0
	*libc.As[bool](retval) = loadedv65
	goto _return

sw_bb66:
	v35 = *libc.As[int32](lookahead)
	cmp67 = v35 == 46
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end70:
	v36 = *libc.As[byte](result)
	loadedv71 = (v36 & 1) != 0
	*libc.As[bool](retval) = loadedv71
	goto _return

sw_bb72:
	v37 = *libc.As[int32](lookahead)
	cmp73 = v37 == 49
	if cmp73 {
		goto if_then75
	} else {
		goto if_end76
	}

if_then75:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end76:
	v38 = *libc.As[byte](result)
	loadedv77 = (v38 & 1) != 0
	*libc.As[bool](retval) = loadedv77
	goto _return

sw_bb78:
	v39 = *libc.As[int32](lookahead)
	cmp79 = v39 == 97
	if cmp79 {
		goto if_then81
	} else {
		goto if_end82
	}

if_then81:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end82:
	v40 = *libc.As[byte](result)
	loadedv83 = (v40 & 1) != 0
	*libc.As[bool](retval) = loadedv83
	goto _return

sw_bb84:
	v41 = *libc.As[int32](lookahead)
	cmp85 = v41 == 97
	if cmp85 {
		goto if_then87
	} else {
		goto if_end88
	}

if_then87:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end88:
	v42 = *libc.As[byte](result)
	loadedv89 = (v42 & 1) != 0
	*libc.As[bool](retval) = loadedv89
	goto _return

sw_bb90:
	v43 = *libc.As[int32](lookahead)
	cmp91 = v43 == 97
	if cmp91 {
		goto if_then93
	} else {
		goto if_end94
	}

if_then93:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end94:
	v44 = *libc.As[byte](result)
	loadedv95 = (v44 & 1) != 0
	*libc.As[bool](retval) = loadedv95
	goto _return

sw_bb96:
	v45 = *libc.As[int32](lookahead)
	cmp97 = v45 == 97
	if cmp97 {
		goto if_then99
	} else {
		goto if_end100
	}

if_then99:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end100:
	v46 = *libc.As[int32](lookahead)
	cmp101 = v46 == 98
	if cmp101 {
		goto if_then103
	} else {
		goto if_end104
	}

if_then103:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end104:
	v47 = *libc.As[int32](lookahead)
	cmp105 = v47 == 100
	if cmp105 {
		goto if_then107
	} else {
		goto if_end108
	}

if_then107:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end108:
	v48 = *libc.As[int32](lookahead)
	cmp109 = v48 == 112
	if cmp109 {
		goto if_then111
	} else {
		goto if_end112
	}

if_then111:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end112:
	v49 = *libc.As[int32](lookahead)
	cmp113 = v49 == 114
	if cmp113 {
		goto if_then115
	} else {
		goto if_end116
	}

if_then115:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end116:
	v50 = *libc.As[int32](lookahead)
	cmp117 = 9 <= v50
	if cmp117 {
		goto land_lhs_true119
	} else {
		goto lor_lhs_false122
	}

land_lhs_true119:
	v51 = *libc.As[int32](lookahead)
	cmp120 = v51 <= 13
	if cmp120 {
		goto if_then125
	} else {
		goto lor_lhs_false122
	}

lor_lhs_false122:
	v52 = *libc.As[int32](lookahead)
	cmp123 = v52 == 32
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end126:
	v53 = *libc.As[int32](lookahead)
	cmp127 = v53 == 46
	if cmp127 {
		goto if_then135
	} else {
		goto lor_lhs_false129
	}

lor_lhs_false129:
	v54 = *libc.As[int32](lookahead)
	cmp130 = 48 <= v54
	if cmp130 {
		goto land_lhs_true132
	} else {
		goto if_end136
	}

land_lhs_true132:
	v55 = *libc.As[int32](lookahead)
	cmp133 = v55 <= 57
	if cmp133 {
		goto if_then135
	} else {
		goto if_end136
	}

if_then135:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end136:
	v56 = *libc.As[byte](result)
	loadedv137 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv137
	goto _return

sw_bb138:
	v57 = *libc.As[int32](lookahead)
	cmp139 = v57 == 98
	if cmp139 {
		goto if_then141
	} else {
		goto if_end142
	}

if_then141:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end142:
	v58 = *libc.As[byte](result)
	loadedv143 = (v58 & 1) != 0
	*libc.As[bool](retval) = loadedv143
	goto _return

sw_bb144:
	v59 = *libc.As[int32](lookahead)
	cmp145 = v59 == 99
	if cmp145 {
		goto if_then147
	} else {
		goto if_end148
	}

if_then147:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end148:
	v60 = *libc.As[byte](result)
	loadedv149 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv149
	goto _return

sw_bb150:
	v61 = *libc.As[int32](lookahead)
	cmp151 = v61 == 99
	if cmp151 {
		goto if_then153
	} else {
		goto if_end154
	}

if_then153:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end154:
	v62 = *libc.As[byte](result)
	loadedv155 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv155
	goto _return

sw_bb156:
	v63 = *libc.As[int32](lookahead)
	cmp157 = v63 == 100
	if cmp157 {
		goto if_then159
	} else {
		goto if_end160
	}

if_then159:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end160:
	v64 = *libc.As[byte](result)
	loadedv161 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv161
	goto _return

sw_bb162:
	v65 = *libc.As[int32](lookahead)
	cmp163 = v65 == 101
	if cmp163 {
		goto if_then165
	} else {
		goto if_end166
	}

if_then165:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end166:
	v66 = *libc.As[byte](result)
	loadedv167 = (v66 & 1) != 0
	*libc.As[bool](retval) = loadedv167
	goto _return

sw_bb168:
	v67 = *libc.As[int32](lookahead)
	cmp169 = v67 == 101
	if cmp169 {
		goto if_then171
	} else {
		goto if_end172
	}

if_then171:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end172:
	v68 = *libc.As[byte](result)
	loadedv173 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv173
	goto _return

sw_bb174:
	v69 = *libc.As[int32](lookahead)
	cmp175 = v69 == 101
	if cmp175 {
		goto if_then177
	} else {
		goto if_end178
	}

if_then177:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end178:
	v70 = *libc.As[byte](result)
	loadedv179 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv179
	goto _return

sw_bb180:
	v71 = *libc.As[int32](lookahead)
	cmp181 = v71 == 101
	if cmp181 {
		goto if_then183
	} else {
		goto if_end184
	}

if_then183:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end184:
	v72 = *libc.As[byte](result)
	loadedv185 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv185
	goto _return

sw_bb186:
	v73 = *libc.As[int32](lookahead)
	cmp187 = v73 == 104
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end190:
	v74 = *libc.As[byte](result)
	loadedv191 = (v74 & 1) != 0
	*libc.As[bool](retval) = loadedv191
	goto _return

sw_bb192:
	v75 = *libc.As[int32](lookahead)
	cmp193 = v75 == 105
	if cmp193 {
		goto if_then195
	} else {
		goto if_end196
	}

if_then195:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end196:
	v76 = *libc.As[byte](result)
	loadedv197 = (v76 & 1) != 0
	*libc.As[bool](retval) = loadedv197
	goto _return

sw_bb198:
	v77 = *libc.As[int32](lookahead)
	cmp199 = v77 == 105
	if cmp199 {
		goto if_then201
	} else {
		goto if_end202
	}

if_then201:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end202:
	v78 = *libc.As[byte](result)
	loadedv203 = (v78 & 1) != 0
	*libc.As[bool](retval) = loadedv203
	goto _return

sw_bb204:
	v79 = *libc.As[int32](lookahead)
	cmp205 = v79 == 108
	if cmp205 {
		goto if_then207
	} else {
		goto if_end208
	}

if_then207:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end208:
	v80 = *libc.As[byte](result)
	loadedv209 = (v80 & 1) != 0
	*libc.As[bool](retval) = loadedv209
	goto _return

sw_bb210:
	v81 = *libc.As[int32](lookahead)
	cmp211 = v81 == 108
	if cmp211 {
		goto if_then213
	} else {
		goto if_end214
	}

if_then213:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end214:
	v82 = *libc.As[byte](result)
	loadedv215 = (v82 & 1) != 0
	*libc.As[bool](retval) = loadedv215
	goto _return

sw_bb216:
	v83 = *libc.As[int32](lookahead)
	cmp217 = v83 == 109
	if cmp217 {
		goto if_then219
	} else {
		goto if_end220
	}

if_then219:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end220:
	v84 = *libc.As[byte](result)
	loadedv221 = (v84 & 1) != 0
	*libc.As[bool](retval) = loadedv221
	goto _return

sw_bb222:
	v85 = *libc.As[int32](lookahead)
	cmp223 = v85 == 109
	if cmp223 {
		goto if_then225
	} else {
		goto if_end226
	}

if_then225:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end226:
	v86 = *libc.As[byte](result)
	loadedv227 = (v86 & 1) != 0
	*libc.As[bool](retval) = loadedv227
	goto _return

sw_bb228:
	v87 = *libc.As[int32](lookahead)
	cmp229 = v87 == 110
	if cmp229 {
		goto if_then231
	} else {
		goto if_end232
	}

if_then231:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end232:
	v88 = *libc.As[byte](result)
	loadedv233 = (v88 & 1) != 0
	*libc.As[bool](retval) = loadedv233
	goto _return

sw_bb234:
	v89 = *libc.As[int32](lookahead)
	cmp235 = v89 == 111
	if cmp235 {
		goto if_then237
	} else {
		goto if_end238
	}

if_then237:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end238:
	v90 = *libc.As[byte](result)
	loadedv239 = (v90 & 1) != 0
	*libc.As[bool](retval) = loadedv239
	goto _return

sw_bb240:
	v91 = *libc.As[int32](lookahead)
	cmp241 = v91 == 111
	if cmp241 {
		goto if_then243
	} else {
		goto if_end244
	}

if_then243:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end244:
	v92 = *libc.As[byte](result)
	loadedv245 = (v92 & 1) != 0
	*libc.As[bool](retval) = loadedv245
	goto _return

sw_bb246:
	v93 = *libc.As[int32](lookahead)
	cmp247 = v93 == 111
	if cmp247 {
		goto if_then249
	} else {
		goto if_end250
	}

if_then249:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end250:
	v94 = *libc.As[byte](result)
	loadedv251 = (v94 & 1) != 0
	*libc.As[bool](retval) = loadedv251
	goto _return

sw_bb252:
	v95 = *libc.As[int32](lookahead)
	cmp253 = v95 == 112
	if cmp253 {
		goto if_then255
	} else {
		goto if_end256
	}

if_then255:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end256:
	v96 = *libc.As[byte](result)
	loadedv257 = (v96 & 1) != 0
	*libc.As[bool](retval) = loadedv257
	goto _return

sw_bb258:
	v97 = *libc.As[int32](lookahead)
	cmp259 = v97 == 112
	if cmp259 {
		goto if_then261
	} else {
		goto if_end262
	}

if_then261:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end262:
	v98 = *libc.As[byte](result)
	loadedv263 = (v98 & 1) != 0
	*libc.As[bool](retval) = loadedv263
	goto _return

sw_bb264:
	v99 = *libc.As[int32](lookahead)
	cmp265 = v99 == 114
	if cmp265 {
		goto if_then267
	} else {
		goto if_end268
	}

if_then267:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end268:
	v100 = *libc.As[byte](result)
	loadedv269 = (v100 & 1) != 0
	*libc.As[bool](retval) = loadedv269
	goto _return

sw_bb270:
	v101 = *libc.As[int32](lookahead)
	cmp271 = v101 == 116
	if cmp271 {
		goto if_then273
	} else {
		goto if_end274
	}

if_then273:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end274:
	v102 = *libc.As[byte](result)
	loadedv275 = (v102 & 1) != 0
	*libc.As[bool](retval) = loadedv275
	goto _return

sw_bb276:
	v103 = *libc.As[int32](lookahead)
	cmp277 = v103 == 116
	if cmp277 {
		goto if_then279
	} else {
		goto if_end280
	}

if_then279:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end280:
	v104 = *libc.As[byte](result)
	loadedv281 = (v104 & 1) != 0
	*libc.As[bool](retval) = loadedv281
	goto _return

sw_bb282:
	v105 = *libc.As[int32](lookahead)
	cmp283 = v105 == 118
	if cmp283 {
		goto if_then285
	} else {
		goto if_end286
	}

if_then285:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end286:
	v106 = *libc.As[byte](result)
	loadedv287 = (v106 & 1) != 0
	*libc.As[bool](retval) = loadedv287
	goto _return

sw_bb288:
	v107 = *libc.As[int32](lookahead)
	cmp289 = 9 <= v107
	if cmp289 {
		goto land_lhs_true291
	} else {
		goto lor_lhs_false294
	}

land_lhs_true291:
	v108 = *libc.As[int32](lookahead)
	cmp292 = v108 <= 13
	if cmp292 {
		goto if_then297
	} else {
		goto lor_lhs_false294
	}

lor_lhs_false294:
	v109 = *libc.As[int32](lookahead)
	cmp295 = v109 == 32
	if cmp295 {
		goto if_then297
	} else {
		goto if_end298
	}

if_then297:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end298:
	v110 = *libc.As[int32](lookahead)
	cmp299 = 48 <= v110
	if cmp299 {
		goto land_lhs_true301
	} else {
		goto lor_lhs_false304
	}

land_lhs_true301:
	v111 = *libc.As[int32](lookahead)
	cmp302 = v111 <= 57
	if cmp302 {
		goto if_then316
	} else {
		goto lor_lhs_false304
	}

lor_lhs_false304:
	v112 = *libc.As[int32](lookahead)
	cmp305 = 65 <= v112
	if cmp305 {
		goto land_lhs_true307
	} else {
		goto lor_lhs_false310
	}

land_lhs_true307:
	v113 = *libc.As[int32](lookahead)
	cmp308 = v113 <= 70
	if cmp308 {
		goto if_then316
	} else {
		goto lor_lhs_false310
	}

lor_lhs_false310:
	v114 = *libc.As[int32](lookahead)
	cmp311 = 97 <= v114
	if cmp311 {
		goto land_lhs_true313
	} else {
		goto if_end317
	}

land_lhs_true313:
	v115 = *libc.As[int32](lookahead)
	cmp314 = v115 <= 102
	if cmp314 {
		goto if_then316
	} else {
		goto if_end317
	}

if_then316:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end317:
	v116 = *libc.As[byte](result)
	loadedv318 = (v116 & 1) != 0
	*libc.As[bool](retval) = loadedv318
	goto _return

sw_bb319:
	v117 = *libc.As[int32](lookahead)
	cmp320 = 9 <= v117
	if cmp320 {
		goto land_lhs_true322
	} else {
		goto lor_lhs_false325
	}

land_lhs_true322:
	v118 = *libc.As[int32](lookahead)
	cmp323 = v118 <= 13
	if cmp323 {
		goto if_then328
	} else {
		goto lor_lhs_false325
	}

lor_lhs_false325:
	v119 = *libc.As[int32](lookahead)
	cmp326 = v119 == 32
	if cmp326 {
		goto if_then328
	} else {
		goto if_end329
	}

if_then328:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end329:
	v120 = *libc.As[int32](lookahead)
	cmp330 = v120 == 43
	if cmp330 {
		goto if_then350
	} else {
		goto lor_lhs_false332
	}

lor_lhs_false332:
	v121 = *libc.As[int32](lookahead)
	cmp333 = 47 <= v121
	if cmp333 {
		goto land_lhs_true335
	} else {
		goto lor_lhs_false338
	}

land_lhs_true335:
	v122 = *libc.As[int32](lookahead)
	cmp336 = v122 <= 57
	if cmp336 {
		goto if_then350
	} else {
		goto lor_lhs_false338
	}

lor_lhs_false338:
	v123 = *libc.As[int32](lookahead)
	cmp339 = 65 <= v123
	if cmp339 {
		goto land_lhs_true341
	} else {
		goto lor_lhs_false344
	}

land_lhs_true341:
	v124 = *libc.As[int32](lookahead)
	cmp342 = v124 <= 90
	if cmp342 {
		goto if_then350
	} else {
		goto lor_lhs_false344
	}

lor_lhs_false344:
	v125 = *libc.As[int32](lookahead)
	cmp345 = 97 <= v125
	if cmp345 {
		goto land_lhs_true347
	} else {
		goto if_end351
	}

land_lhs_true347:
	v126 = *libc.As[int32](lookahead)
	cmp348 = v126 <= 122
	if cmp348 {
		goto if_then350
	} else {
		goto if_end351
	}

if_then350:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end351:
	v127 = *libc.As[byte](result)
	loadedv352 = (v127 & 1) != 0
	*libc.As[bool](retval) = loadedv352
	goto _return

sw_bb353:
	v128 = *libc.As[byte](eof)
	loadedv354 = (v128 & 1) != 0
	if loadedv354 {
		goto if_then355
	} else {
		goto if_end356
	}

if_then355:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end356:
	v129 = *libc.As[int32](lookahead)
	cmp357 = 9 <= v129
	if cmp357 {
		goto land_lhs_true359
	} else {
		goto lor_lhs_false362
	}

land_lhs_true359:
	v130 = *libc.As[int32](lookahead)
	cmp360 = v130 <= 13
	if cmp360 {
		goto if_then365
	} else {
		goto lor_lhs_false362
	}

lor_lhs_false362:
	v131 = *libc.As[int32](lookahead)
	cmp363 = v131 == 32
	if cmp363 {
		goto if_then365
	} else {
		goto if_end366
	}

if_then365:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end366:
	v132 = *libc.As[int32](lookahead)
	cmp367 = 45 <= v132
	if cmp367 {
		goto land_lhs_true369
	} else {
		goto lor_lhs_false372
	}

land_lhs_true369:
	v133 = *libc.As[int32](lookahead)
	cmp370 = v133 <= 57
	if cmp370 {
		goto if_then387
	} else {
		goto lor_lhs_false372
	}

lor_lhs_false372:
	v134 = *libc.As[int32](lookahead)
	cmp373 = 65 <= v134
	if cmp373 {
		goto land_lhs_true375
	} else {
		goto lor_lhs_false378
	}

land_lhs_true375:
	v135 = *libc.As[int32](lookahead)
	cmp376 = v135 <= 90
	if cmp376 {
		goto if_then387
	} else {
		goto lor_lhs_false378
	}

lor_lhs_false378:
	v136 = *libc.As[int32](lookahead)
	cmp379 = v136 == 95
	if cmp379 {
		goto if_then387
	} else {
		goto lor_lhs_false381
	}

lor_lhs_false381:
	v137 = *libc.As[int32](lookahead)
	cmp382 = 97 <= v137
	if cmp382 {
		goto land_lhs_true384
	} else {
		goto if_end388
	}

land_lhs_true384:
	v138 = *libc.As[int32](lookahead)
	cmp385 = v138 <= 122
	if cmp385 {
		goto if_then387
	} else {
		goto if_end388
	}

if_then387:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end388:
	v139 = *libc.As[byte](result)
	loadedv389 = (v139 & 1) != 0
	*libc.As[bool](retval) = loadedv389
	goto _return

sw_bb390:
	*libc.As[byte](result) = 1
	v140 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v140).F1)
	*libc.As[int16](result_symbol) = 0
	v141 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v141).F3)
	v142 = *libc.As[unsafe.Pointer](mark_end)
	v143 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v142)(v143)
	v144 = *libc.As[byte](result)
	loadedv391 = (v144 & 1) != 0
	*libc.As[bool](retval) = loadedv391
	goto _return

sw_bb392:
	*libc.As[byte](result) = 1
	v145 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol393 = libc.Ptr(&libc.As[TSLexer](v145).F1)
	*libc.As[int16](result_symbol393) = 1
	v146 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end394 = libc.Ptr(&libc.As[TSLexer](v146).F3)
	v147 = *libc.As[unsafe.Pointer](mark_end394)
	v148 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v147)(v148)
	v149 = *libc.As[byte](result)
	loadedv395 = (v149 & 1) != 0
	*libc.As[bool](retval) = loadedv395
	goto _return

sw_bb396:
	*libc.As[byte](result) = 1
	v150 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol397 = libc.Ptr(&libc.As[TSLexer](v150).F1)
	*libc.As[int16](result_symbol397) = 2
	v151 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end398 = libc.Ptr(&libc.As[TSLexer](v151).F3)
	v152 = *libc.As[unsafe.Pointer](mark_end398)
	v153 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v152)(v153)
	v154 = *libc.As[byte](result)
	loadedv399 = (v154 & 1) != 0
	*libc.As[bool](retval) = loadedv399
	goto _return

sw_bb400:
	*libc.As[byte](result) = 1
	v155 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol401 = libc.Ptr(&libc.As[TSLexer](v155).F1)
	*libc.As[int16](result_symbol401) = 3
	v156 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end402 = libc.Ptr(&libc.As[TSLexer](v156).F3)
	v157 = *libc.As[unsafe.Pointer](mark_end402)
	v158 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v157)(v158)
	v159 = *libc.As[int32](lookahead)
	cmp403 = 45 <= v159
	if cmp403 {
		goto land_lhs_true405
	} else {
		goto lor_lhs_false408
	}

land_lhs_true405:
	v160 = *libc.As[int32](lookahead)
	cmp406 = v160 <= 57
	if cmp406 {
		goto if_then423
	} else {
		goto lor_lhs_false408
	}

lor_lhs_false408:
	v161 = *libc.As[int32](lookahead)
	cmp409 = 65 <= v161
	if cmp409 {
		goto land_lhs_true411
	} else {
		goto lor_lhs_false414
	}

land_lhs_true411:
	v162 = *libc.As[int32](lookahead)
	cmp412 = v162 <= 90
	if cmp412 {
		goto if_then423
	} else {
		goto lor_lhs_false414
	}

lor_lhs_false414:
	v163 = *libc.As[int32](lookahead)
	cmp415 = v163 == 95
	if cmp415 {
		goto if_then423
	} else {
		goto lor_lhs_false417
	}

lor_lhs_false417:
	v164 = *libc.As[int32](lookahead)
	cmp418 = 97 <= v164
	if cmp418 {
		goto land_lhs_true420
	} else {
		goto if_end424
	}

land_lhs_true420:
	v165 = *libc.As[int32](lookahead)
	cmp421 = v165 <= 122
	if cmp421 {
		goto if_then423
	} else {
		goto if_end424
	}

if_then423:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end424:
	v166 = *libc.As[byte](result)
	loadedv425 = (v166 & 1) != 0
	*libc.As[bool](retval) = loadedv425
	goto _return

sw_bb426:
	*libc.As[byte](result) = 1
	v167 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol427 = libc.Ptr(&libc.As[TSLexer](v167).F1)
	*libc.As[int16](result_symbol427) = 4
	v168 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end428 = libc.Ptr(&libc.As[TSLexer](v168).F3)
	v169 = *libc.As[unsafe.Pointer](mark_end428)
	v170 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v169)(v170)
	v171 = *libc.As[byte](result)
	loadedv429 = (v171 & 1) != 0
	*libc.As[bool](retval) = loadedv429
	goto _return

sw_bb430:
	*libc.As[byte](result) = 1
	v172 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol431 = libc.Ptr(&libc.As[TSLexer](v172).F1)
	*libc.As[int16](result_symbol431) = 5
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end432 = libc.Ptr(&libc.As[TSLexer](v173).F3)
	v174 = *libc.As[unsafe.Pointer](mark_end432)
	v175 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v174)(v175)
	v176 = *libc.As[byte](result)
	loadedv433 = (v176 & 1) != 0
	*libc.As[bool](retval) = loadedv433
	goto _return

sw_bb434:
	*libc.As[byte](result) = 1
	v177 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol435 = libc.Ptr(&libc.As[TSLexer](v177).F1)
	*libc.As[int16](result_symbol435) = 6
	v178 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end436 = libc.Ptr(&libc.As[TSLexer](v178).F3)
	v179 = *libc.As[unsafe.Pointer](mark_end436)
	v180 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v179)(v180)
	v181 = *libc.As[byte](result)
	loadedv437 = (v181 & 1) != 0
	*libc.As[bool](retval) = loadedv437
	goto _return

sw_bb438:
	*libc.As[byte](result) = 1
	v182 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol439 = libc.Ptr(&libc.As[TSLexer](v182).F1)
	*libc.As[int16](result_symbol439) = 7
	v183 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end440 = libc.Ptr(&libc.As[TSLexer](v183).F3)
	v184 = *libc.As[unsafe.Pointer](mark_end440)
	v185 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v184)(v185)
	v186 = *libc.As[byte](result)
	loadedv441 = (v186 & 1) != 0
	*libc.As[bool](retval) = loadedv441
	goto _return

sw_bb442:
	*libc.As[byte](result) = 1
	v187 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol443 = libc.Ptr(&libc.As[TSLexer](v187).F1)
	*libc.As[int16](result_symbol443) = 8
	v188 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end444 = libc.Ptr(&libc.As[TSLexer](v188).F3)
	v189 = *libc.As[unsafe.Pointer](mark_end444)
	v190 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v189)(v190)
	v191 = *libc.As[byte](result)
	loadedv445 = (v191 & 1) != 0
	*libc.As[bool](retval) = loadedv445
	goto _return

sw_bb446:
	*libc.As[byte](result) = 1
	v192 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol447 = libc.Ptr(&libc.As[TSLexer](v192).F1)
	*libc.As[int16](result_symbol447) = 9
	v193 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end448 = libc.Ptr(&libc.As[TSLexer](v193).F3)
	v194 = *libc.As[unsafe.Pointer](mark_end448)
	v195 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v194)(v195)
	v196 = *libc.As[byte](result)
	loadedv449 = (v196 & 1) != 0
	*libc.As[bool](retval) = loadedv449
	goto _return

sw_bb450:
	*libc.As[byte](result) = 1
	v197 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol451 = libc.Ptr(&libc.As[TSLexer](v197).F1)
	*libc.As[int16](result_symbol451) = 10
	v198 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end452 = libc.Ptr(&libc.As[TSLexer](v198).F3)
	v199 = *libc.As[unsafe.Pointer](mark_end452)
	v200 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v199)(v200)
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
	*libc.As[int16](result_symbol459) = 12
	v208 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end460 = libc.Ptr(&libc.As[TSLexer](v208).F3)
	v209 = *libc.As[unsafe.Pointer](mark_end460)
	v210 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v209)(v210)
	v211 = *libc.As[byte](result)
	loadedv461 = (v211 & 1) != 0
	*libc.As[bool](retval) = loadedv461
	goto _return

sw_bb462:
	*libc.As[byte](result) = 1
	v212 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol463 = libc.Ptr(&libc.As[TSLexer](v212).F1)
	*libc.As[int16](result_symbol463) = 13
	v213 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end464 = libc.Ptr(&libc.As[TSLexer](v213).F3)
	v214 = *libc.As[unsafe.Pointer](mark_end464)
	v215 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v214)(v215)
	v216 = *libc.As[byte](result)
	loadedv465 = (v216 & 1) != 0
	*libc.As[bool](retval) = loadedv465
	goto _return

sw_bb466:
	*libc.As[byte](result) = 1
	v217 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol467 = libc.Ptr(&libc.As[TSLexer](v217).F1)
	*libc.As[int16](result_symbol467) = 14
	v218 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end468 = libc.Ptr(&libc.As[TSLexer](v218).F3)
	v219 = *libc.As[unsafe.Pointer](mark_end468)
	v220 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v219)(v220)
	v221 = *libc.As[byte](result)
	loadedv469 = (v221 & 1) != 0
	*libc.As[bool](retval) = loadedv469
	goto _return

sw_bb470:
	*libc.As[byte](result) = 1
	v222 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol471 = libc.Ptr(&libc.As[TSLexer](v222).F1)
	*libc.As[int16](result_symbol471) = 15
	v223 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end472 = libc.Ptr(&libc.As[TSLexer](v223).F3)
	v224 = *libc.As[unsafe.Pointer](mark_end472)
	v225 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v224)(v225)
	v226 = *libc.As[byte](result)
	loadedv473 = (v226 & 1) != 0
	*libc.As[bool](retval) = loadedv473
	goto _return

sw_bb474:
	*libc.As[byte](result) = 1
	v227 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol475 = libc.Ptr(&libc.As[TSLexer](v227).F1)
	*libc.As[int16](result_symbol475) = 15
	v228 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end476 = libc.Ptr(&libc.As[TSLexer](v228).F3)
	v229 = *libc.As[unsafe.Pointer](mark_end476)
	v230 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v229)(v230)
	v231 = *libc.As[int32](lookahead)
	cmp477 = v231 == 61
	if cmp477 {
		goto if_then479
	} else {
		goto if_end480
	}

if_then479:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end480:
	v232 = *libc.As[int32](lookahead)
	cmp481 = v232 == 43
	if cmp481 {
		goto if_then501
	} else {
		goto lor_lhs_false483
	}

lor_lhs_false483:
	v233 = *libc.As[int32](lookahead)
	cmp484 = 47 <= v233
	if cmp484 {
		goto land_lhs_true486
	} else {
		goto lor_lhs_false489
	}

land_lhs_true486:
	v234 = *libc.As[int32](lookahead)
	cmp487 = v234 <= 57
	if cmp487 {
		goto if_then501
	} else {
		goto lor_lhs_false489
	}

lor_lhs_false489:
	v235 = *libc.As[int32](lookahead)
	cmp490 = 65 <= v235
	if cmp490 {
		goto land_lhs_true492
	} else {
		goto lor_lhs_false495
	}

land_lhs_true492:
	v236 = *libc.As[int32](lookahead)
	cmp493 = v236 <= 90
	if cmp493 {
		goto if_then501
	} else {
		goto lor_lhs_false495
	}

lor_lhs_false495:
	v237 = *libc.As[int32](lookahead)
	cmp496 = 97 <= v237
	if cmp496 {
		goto land_lhs_true498
	} else {
		goto if_end502
	}

land_lhs_true498:
	v238 = *libc.As[int32](lookahead)
	cmp499 = v238 <= 122
	if cmp499 {
		goto if_then501
	} else {
		goto if_end502
	}

if_then501:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end502:
	v239 = *libc.As[byte](result)
	loadedv503 = (v239 & 1) != 0
	*libc.As[bool](retval) = loadedv503
	goto _return

sw_bb504:
	*libc.As[byte](result) = 1
	v240 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol505 = libc.Ptr(&libc.As[TSLexer](v240).F1)
	*libc.As[int16](result_symbol505) = 15
	v241 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end506 = libc.Ptr(&libc.As[TSLexer](v241).F3)
	v242 = *libc.As[unsafe.Pointer](mark_end506)
	v243 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v242)(v243)
	v244 = *libc.As[int32](lookahead)
	cmp507 = v244 == 61
	if cmp507 {
		goto if_then509
	} else {
		goto if_end510
	}

if_then509:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end510:
	v245 = *libc.As[byte](result)
	loadedv511 = (v245 & 1) != 0
	*libc.As[bool](retval) = loadedv511
	goto _return

sw_bb512:
	*libc.As[byte](result) = 1
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol513 = libc.Ptr(&libc.As[TSLexer](v246).F1)
	*libc.As[int16](result_symbol513) = 16
	v247 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end514 = libc.Ptr(&libc.As[TSLexer](v247).F3)
	v248 = *libc.As[unsafe.Pointer](mark_end514)
	v249 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v248)(v249)
	v250 = *libc.As[int32](lookahead)
	cmp515 = 48 <= v250
	if cmp515 {
		goto land_lhs_true517
	} else {
		goto if_end521
	}

land_lhs_true517:
	v251 = *libc.As[int32](lookahead)
	cmp518 = v251 <= 57
	if cmp518 {
		goto if_then520
	} else {
		goto if_end521
	}

if_then520:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end521:
	v252 = *libc.As[int32](lookahead)
	cmp522 = 65 <= v252
	if cmp522 {
		goto land_lhs_true524
	} else {
		goto lor_lhs_false527
	}

land_lhs_true524:
	v253 = *libc.As[int32](lookahead)
	cmp525 = v253 <= 70
	if cmp525 {
		goto if_then533
	} else {
		goto lor_lhs_false527
	}

lor_lhs_false527:
	v254 = *libc.As[int32](lookahead)
	cmp528 = 97 <= v254
	if cmp528 {
		goto land_lhs_true530
	} else {
		goto if_end534
	}

land_lhs_true530:
	v255 = *libc.As[int32](lookahead)
	cmp531 = v255 <= 102
	if cmp531 {
		goto if_then533
	} else {
		goto if_end534
	}

if_then533:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end534:
	v256 = *libc.As[byte](result)
	loadedv535 = (v256 & 1) != 0
	*libc.As[bool](retval) = loadedv535
	goto _return

sw_bb536:
	*libc.As[byte](result) = 1
	v257 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol537 = libc.Ptr(&libc.As[TSLexer](v257).F1)
	*libc.As[int16](result_symbol537) = 16
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end538 = libc.Ptr(&libc.As[TSLexer](v258).F3)
	v259 = *libc.As[unsafe.Pointer](mark_end538)
	v260 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v259)(v260)
	v261 = *libc.As[int32](lookahead)
	cmp539 = 48 <= v261
	if cmp539 {
		goto land_lhs_true541
	} else {
		goto if_end545
	}

land_lhs_true541:
	v262 = *libc.As[int32](lookahead)
	cmp542 = v262 <= 57
	if cmp542 {
		goto if_then544
	} else {
		goto if_end545
	}

if_then544:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end545:
	v263 = *libc.As[byte](result)
	loadedv546 = (v263 & 1) != 0
	*libc.As[bool](retval) = loadedv546
	goto _return

sw_bb547:
	*libc.As[byte](result) = 1
	v264 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol548 = libc.Ptr(&libc.As[TSLexer](v264).F1)
	*libc.As[int16](result_symbol548) = 17
	v265 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end549 = libc.Ptr(&libc.As[TSLexer](v265).F3)
	v266 = *libc.As[unsafe.Pointer](mark_end549)
	v267 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v266)(v267)
	v268 = *libc.As[int32](lookahead)
	cmp550 = v268 == 46
	if cmp550 {
		goto if_then558
	} else {
		goto lor_lhs_false552
	}

lor_lhs_false552:
	v269 = *libc.As[int32](lookahead)
	cmp553 = 48 <= v269
	if cmp553 {
		goto land_lhs_true555
	} else {
		goto if_end559
	}

land_lhs_true555:
	v270 = *libc.As[int32](lookahead)
	cmp556 = v270 <= 57
	if cmp556 {
		goto if_then558
	} else {
		goto if_end559
	}

if_then558:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end559:
	v271 = *libc.As[byte](result)
	loadedv560 = (v271 & 1) != 0
	*libc.As[bool](retval) = loadedv560
	goto _return

sw_bb561:
	*libc.As[byte](result) = 1
	v272 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol562 = libc.Ptr(&libc.As[TSLexer](v272).F1)
	*libc.As[int16](result_symbol562) = 18
	v273 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end563 = libc.Ptr(&libc.As[TSLexer](v273).F3)
	v274 = *libc.As[unsafe.Pointer](mark_end563)
	v275 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v274)(v275)
	v276 = *libc.As[int32](lookahead)
	cmp564 = v276 == 101
	if cmp564 {
		goto if_then566
	} else {
		goto if_end567
	}

if_then566:
	*libc.As[int16](state_addr) = 60
	goto next_state

if_end567:
	v277 = *libc.As[int32](lookahead)
	cmp568 = 48 <= v277
	if cmp568 {
		goto land_lhs_true570
	} else {
		goto lor_lhs_false573
	}

land_lhs_true570:
	v278 = *libc.As[int32](lookahead)
	cmp571 = v278 <= 57
	if cmp571 {
		goto if_then585
	} else {
		goto lor_lhs_false573
	}

lor_lhs_false573:
	v279 = *libc.As[int32](lookahead)
	cmp574 = 65 <= v279
	if cmp574 {
		goto land_lhs_true576
	} else {
		goto lor_lhs_false579
	}

land_lhs_true576:
	v280 = *libc.As[int32](lookahead)
	cmp577 = v280 <= 70
	if cmp577 {
		goto if_then585
	} else {
		goto lor_lhs_false579
	}

lor_lhs_false579:
	v281 = *libc.As[int32](lookahead)
	cmp580 = 97 <= v281
	if cmp580 {
		goto land_lhs_true582
	} else {
		goto if_end586
	}

land_lhs_true582:
	v282 = *libc.As[int32](lookahead)
	cmp583 = v282 <= 102
	if cmp583 {
		goto if_then585
	} else {
		goto if_end586
	}

if_then585:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end586:
	v283 = *libc.As[byte](result)
	loadedv587 = (v283 & 1) != 0
	*libc.As[bool](retval) = loadedv587
	goto _return

sw_bb588:
	*libc.As[byte](result) = 1
	v284 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol589 = libc.Ptr(&libc.As[TSLexer](v284).F1)
	*libc.As[int16](result_symbol589) = 18
	v285 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end590 = libc.Ptr(&libc.As[TSLexer](v285).F3)
	v286 = *libc.As[unsafe.Pointer](mark_end590)
	v287 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v286)(v287)
	v288 = *libc.As[int32](lookahead)
	cmp591 = v288 == 101
	if cmp591 {
		goto if_then593
	} else {
		goto if_end594
	}

if_then593:
	*libc.As[int16](state_addr) = 61
	goto next_state

if_end594:
	v289 = *libc.As[int32](lookahead)
	cmp595 = 48 <= v289
	if cmp595 {
		goto land_lhs_true597
	} else {
		goto lor_lhs_false600
	}

land_lhs_true597:
	v290 = *libc.As[int32](lookahead)
	cmp598 = v290 <= 57
	if cmp598 {
		goto if_then612
	} else {
		goto lor_lhs_false600
	}

lor_lhs_false600:
	v291 = *libc.As[int32](lookahead)
	cmp601 = 65 <= v291
	if cmp601 {
		goto land_lhs_true603
	} else {
		goto lor_lhs_false606
	}

land_lhs_true603:
	v292 = *libc.As[int32](lookahead)
	cmp604 = v292 <= 70
	if cmp604 {
		goto if_then612
	} else {
		goto lor_lhs_false606
	}

lor_lhs_false606:
	v293 = *libc.As[int32](lookahead)
	cmp607 = 97 <= v293
	if cmp607 {
		goto land_lhs_true609
	} else {
		goto if_end613
	}

land_lhs_true609:
	v294 = *libc.As[int32](lookahead)
	cmp610 = v294 <= 102
	if cmp610 {
		goto if_then612
	} else {
		goto if_end613
	}

if_then612:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end613:
	v295 = *libc.As[byte](result)
	loadedv614 = (v295 & 1) != 0
	*libc.As[bool](retval) = loadedv614
	goto _return

sw_bb615:
	*libc.As[byte](result) = 1
	v296 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol616 = libc.Ptr(&libc.As[TSLexer](v296).F1)
	*libc.As[int16](result_symbol616) = 18
	v297 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end617 = libc.Ptr(&libc.As[TSLexer](v297).F3)
	v298 = *libc.As[unsafe.Pointer](mark_end617)
	v299 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v298)(v299)
	v300 = *libc.As[int32](lookahead)
	cmp618 = v300 == 108
	if cmp618 {
		goto if_then620
	} else {
		goto if_end621
	}

if_then620:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end621:
	v301 = *libc.As[int32](lookahead)
	cmp622 = 48 <= v301
	if cmp622 {
		goto land_lhs_true624
	} else {
		goto lor_lhs_false627
	}

land_lhs_true624:
	v302 = *libc.As[int32](lookahead)
	cmp625 = v302 <= 57
	if cmp625 {
		goto if_then639
	} else {
		goto lor_lhs_false627
	}

lor_lhs_false627:
	v303 = *libc.As[int32](lookahead)
	cmp628 = 65 <= v303
	if cmp628 {
		goto land_lhs_true630
	} else {
		goto lor_lhs_false633
	}

land_lhs_true630:
	v304 = *libc.As[int32](lookahead)
	cmp631 = v304 <= 70
	if cmp631 {
		goto if_then639
	} else {
		goto lor_lhs_false633
	}

lor_lhs_false633:
	v305 = *libc.As[int32](lookahead)
	cmp634 = 97 <= v305
	if cmp634 {
		goto land_lhs_true636
	} else {
		goto if_end640
	}

land_lhs_true636:
	v306 = *libc.As[int32](lookahead)
	cmp637 = v306 <= 102
	if cmp637 {
		goto if_then639
	} else {
		goto if_end640
	}

if_then639:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end640:
	v307 = *libc.As[byte](result)
	loadedv641 = (v307 & 1) != 0
	*libc.As[bool](retval) = loadedv641
	goto _return

sw_bb642:
	*libc.As[byte](result) = 1
	v308 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol643 = libc.Ptr(&libc.As[TSLexer](v308).F1)
	*libc.As[int16](result_symbol643) = 18
	v309 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end644 = libc.Ptr(&libc.As[TSLexer](v309).F3)
	v310 = *libc.As[unsafe.Pointer](mark_end644)
	v311 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v310)(v311)
	v312 = *libc.As[int32](lookahead)
	cmp645 = v312 == 116
	if cmp645 {
		goto if_then647
	} else {
		goto if_end648
	}

if_then647:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end648:
	v313 = *libc.As[int32](lookahead)
	cmp649 = 48 <= v313
	if cmp649 {
		goto land_lhs_true651
	} else {
		goto lor_lhs_false654
	}

land_lhs_true651:
	v314 = *libc.As[int32](lookahead)
	cmp652 = v314 <= 57
	if cmp652 {
		goto if_then666
	} else {
		goto lor_lhs_false654
	}

lor_lhs_false654:
	v315 = *libc.As[int32](lookahead)
	cmp655 = 65 <= v315
	if cmp655 {
		goto land_lhs_true657
	} else {
		goto lor_lhs_false660
	}

land_lhs_true657:
	v316 = *libc.As[int32](lookahead)
	cmp658 = v316 <= 70
	if cmp658 {
		goto if_then666
	} else {
		goto lor_lhs_false660
	}

lor_lhs_false660:
	v317 = *libc.As[int32](lookahead)
	cmp661 = 97 <= v317
	if cmp661 {
		goto land_lhs_true663
	} else {
		goto if_end667
	}

land_lhs_true663:
	v318 = *libc.As[int32](lookahead)
	cmp664 = v318 <= 102
	if cmp664 {
		goto if_then666
	} else {
		goto if_end667
	}

if_then666:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end667:
	v319 = *libc.As[byte](result)
	loadedv668 = (v319 & 1) != 0
	*libc.As[bool](retval) = loadedv668
	goto _return

sw_bb669:
	*libc.As[byte](result) = 1
	v320 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol670 = libc.Ptr(&libc.As[TSLexer](v320).F1)
	*libc.As[int16](result_symbol670) = 18
	v321 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end671 = libc.Ptr(&libc.As[TSLexer](v321).F3)
	v322 = *libc.As[unsafe.Pointer](mark_end671)
	v323 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v322)(v323)
	v324 = *libc.As[int32](lookahead)
	cmp672 = v324 == 118
	if cmp672 {
		goto if_then674
	} else {
		goto if_end675
	}

if_then674:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end675:
	v325 = *libc.As[int32](lookahead)
	cmp676 = 48 <= v325
	if cmp676 {
		goto land_lhs_true678
	} else {
		goto lor_lhs_false681
	}

land_lhs_true678:
	v326 = *libc.As[int32](lookahead)
	cmp679 = v326 <= 57
	if cmp679 {
		goto if_then693
	} else {
		goto lor_lhs_false681
	}

lor_lhs_false681:
	v327 = *libc.As[int32](lookahead)
	cmp682 = 65 <= v327
	if cmp682 {
		goto land_lhs_true684
	} else {
		goto lor_lhs_false687
	}

land_lhs_true684:
	v328 = *libc.As[int32](lookahead)
	cmp685 = v328 <= 70
	if cmp685 {
		goto if_then693
	} else {
		goto lor_lhs_false687
	}

lor_lhs_false687:
	v329 = *libc.As[int32](lookahead)
	cmp688 = 97 <= v329
	if cmp688 {
		goto land_lhs_true690
	} else {
		goto if_end694
	}

land_lhs_true690:
	v330 = *libc.As[int32](lookahead)
	cmp691 = v330 <= 102
	if cmp691 {
		goto if_then693
	} else {
		goto if_end694
	}

if_then693:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end694:
	v331 = *libc.As[byte](result)
	loadedv695 = (v331 & 1) != 0
	*libc.As[bool](retval) = loadedv695
	goto _return

sw_bb696:
	*libc.As[byte](result) = 1
	v332 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol697 = libc.Ptr(&libc.As[TSLexer](v332).F1)
	*libc.As[int16](result_symbol697) = 18
	v333 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end698 = libc.Ptr(&libc.As[TSLexer](v333).F3)
	v334 = *libc.As[unsafe.Pointer](mark_end698)
	v335 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v334)(v335)
	v336 = *libc.As[int32](lookahead)
	cmp699 = 48 <= v336
	if cmp699 {
		goto land_lhs_true701
	} else {
		goto lor_lhs_false704
	}

land_lhs_true701:
	v337 = *libc.As[int32](lookahead)
	cmp702 = v337 <= 57
	if cmp702 {
		goto if_then716
	} else {
		goto lor_lhs_false704
	}

lor_lhs_false704:
	v338 = *libc.As[int32](lookahead)
	cmp705 = 65 <= v338
	if cmp705 {
		goto land_lhs_true707
	} else {
		goto lor_lhs_false710
	}

land_lhs_true707:
	v339 = *libc.As[int32](lookahead)
	cmp708 = v339 <= 70
	if cmp708 {
		goto if_then716
	} else {
		goto lor_lhs_false710
	}

lor_lhs_false710:
	v340 = *libc.As[int32](lookahead)
	cmp711 = 97 <= v340
	if cmp711 {
		goto land_lhs_true713
	} else {
		goto if_end717
	}

land_lhs_true713:
	v341 = *libc.As[int32](lookahead)
	cmp714 = v341 <= 102
	if cmp714 {
		goto if_then716
	} else {
		goto if_end717
	}

if_then716:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end717:
	v342 = *libc.As[byte](result)
	loadedv718 = (v342 & 1) != 0
	*libc.As[bool](retval) = loadedv718
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v343 = *libc.As[bool](retval)
	return v343
}
