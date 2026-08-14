package grammar_cst

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

var tree_sitter_cst_language struct {
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
var ts_small_parse_table [529]int16 = [529]int16{10, 11, 1, 3, 13, 1, 4, 15, 1, 6, 17, 1, 9, 19, 1, 12, 21, 1, 13, 8, 1, 28, 20, 1, 24, 30, 1, 26, 41, 5, 19, 20, 21, 22, 23, 10, 11, 1, 3, 13, 1, 4, 15, 1, 6, 17, 1, 9, 19, 1, 12, 21, 1, 13, 8, 1, 28, 20, 1, 24, 30, 1, 26, 55, 5, 19, 20, 21, 22, 23, 8, 5, 1, 11, 7, 1, 13, 23, 1, 0, 25, 1, 14, 3, 1, 28, 34, 1, 18, 35, 1, 17, 6, 2, 16, 27, 8, 5, 1, 11, 7, 1, 13, 27, 1, 0, 29, 1, 14, 3, 1, 28, 34, 1, 18, 35, 1, 17, 6, 2, 16, 27, 7, 33, 1, 11, 36, 1, 13, 3, 1, 28, 34, 1, 18, 35, 1, 17, 31, 2, 0, 14, 6, 2, 16, 27, 8, 5, 1, 11, 7, 1, 13, 27, 1, 0, 29, 1, 14, 3, 1, 28, 34, 1, 18, 35, 1, 17, 4, 2, 16, 27, 4, 43, 1, 13, 8, 1, 28, 41, 2, 6, 9, 39, 3, 3, 4, 12, 4, 46, 1, 8, 48, 1, 10, 11, 1, 29, 46, 1, 25, 4, 50, 1, 8, 52, 1, 10, 15, 1, 29, 48, 1, 25, 4, 48, 1, 10, 54, 1, 3, 56, 1, 8, 14, 1, 29, 1, 58, 4, 0, 11, 13, 14, 1, 60, 4, 0, 11, 13, 14, 4, 62, 1, 3, 64, 1, 8, 67, 1, 10, 14, 1, 29, 4, 52, 1, 10, 54, 1, 4, 70, 1, 8, 16, 1, 29, 4, 62, 1, 4, 72, 1, 8, 75, 1, 10, 16, 1, 29, 2, 80, 1, 8, 78, 2, 3, 10, 3, 11, 1, 3, 82, 1, 14, 45, 1, 20, 3, 84, 1, 7, 86, 1, 13, 88, 1, 14, 3, 90, 1, 9, 92, 1, 12, 58, 1, 26, 3, 11, 1, 3, 94, 1, 14, 51, 1, 20, 3, 96, 1, 1, 98, 1, 13, 26, 1, 28, 3, 11, 1, 3, 100, 1, 14, 53, 1, 20, 2, 80, 1, 8, 78, 2, 4, 10, 3, 11, 1, 3, 102, 1, 14, 54, 1, 20, 3, 39, 1, 1, 104, 1, 13, 26, 1, 28, 2, 107, 1, 13, 109, 1, 14, 2, 111, 1, 13, 113, 1, 14, 2, 5, 1, 11, 36, 1, 18, 2, 115, 1, 5, 117, 1, 9, 2, 13, 1, 4, 50, 1, 21, 1, 119, 2, 9, 12, 2, 121, 1, 13, 123, 1, 14, 2, 125, 1, 13, 22, 1, 28, 2, 127, 1, 13, 2, 1, 28, 1, 129, 1, 13, 1, 131, 1, 0, 1, 133, 1, 13, 1, 135, 1, 0, 1, 137, 1, 14, 1, 139, 1, 14, 1, 141, 1, 14, 1, 143, 1, 13, 1, 145, 1, 7, 1, 147, 1, 14, 1, 149, 1, 3, 1, 151, 1, 11, 1, 153, 1, 4, 1, 155, 1, 13, 1, 157, 1, 14, 1, 159, 1, 14, 1, 161, 1, 13, 1, 163, 1, 14, 1, 165, 1, 14, 1, 167, 1, 14, 1, 23, 1, 0, 1, 169, 1, 2, 1, 171, 1, 9}
var ts_small_parse_table_map [57]int32 = [57]int32{0, 35, 70, 96, 122, 146, 172, 188, 201, 214, 227, 234, 241, 254, 267, 280, 288, 298, 308, 318, 328, 338, 348, 356, 366, 376, 383, 390, 397, 404, 411, 416, 423, 430, 437, 441, 445, 449, 453, 457, 461, 465, 469, 473, 477, 481, 485, 489, 493, 497, 501, 505, 509, 513, 517, 521, 525}
var ts_symbol_names [31]unsafe.Pointer = [31]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_4), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_30), libc.Ptr(&_str_31)}
var ts_field_names [7]unsafe.Pointer = [7]unsafe.Pointer{nil, libc.Ptr(&_str_32), libc.Ptr(&_str), libc.Ptr(&_str_33), libc.Ptr(&_str_34), libc.Ptr(&_str_35), libc.Ptr(&_str_36)}
var ts_field_map_slices [8]TSMapSlice = [8]TSMapSlice{TSMapSlice{}, TSMapSlice{}, TSMapSlice{0, 2}, TSMapSlice{2, 1}, TSMapSlice{}, TSMapSlice{3, 1}, TSMapSlice{}, TSMapSlice{4, 2}}
var ts_field_map_entries [6]TSFieldMapEntry = [6]TSFieldMapEntry{TSFieldMapEntry{1, 2, 0}, TSFieldMapEntry{5, 0, 0}, TSFieldMapEntry{3, 0, 0}, TSFieldMapEntry{4, 0, 0}, TSFieldMapEntry{2, 4, 0}, TSFieldMapEntry{6, 0, 0}}
var ts_symbol_metadata [31]TSSymbolMetadata = [31]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}}
var ts_symbol_map [31]int16 = [31]int16{0, 1, 2, 3, 4, 5, 6, 2, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [8][5]int16 = [8][5]int16{[5]int16{}, [5]int16{30, 0, 0, 0, 0}, [5]int16{}, [5]int16{}, [5]int16{0, 30, 0, 0, 0}, [5]int16{}, [5]int16{0, 0, 30, 0, 0}, [5]int16{}}
var ts_lex_modes [59]TSLexMode = [59]TSLexMode{TSLexMode{}, TSLexMode{7, 0}, TSLexMode{1, 0}, TSLexMode{1, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{1, 0}, TSLexMode{5, 0}, TSLexMode{5, 0}, TSLexMode{4, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{4, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{4, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{6, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{2, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{3, 0}, TSLexMode{}, TSLexMode{6, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{1, 0}, TSLexMode{6, 0}}
var ts_primary_state_ids [59]int16 = [59]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 11, 14, 17, 18, 19, 20, 21, 22, 23, 17, 25, 8, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58}
var ts_parse_table struct {
	F0 struct {
		F0 [15]int16
		F1 [15]int16
	}
	F1 [30]int16
} = struct {
	F0 struct {
		F0 [15]int16
		F1 [15]int16
	}
	F1 [30]int16
}{struct {
	F0 [15]int16
	F1 [15]int16
}{[15]int16{1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1}, [15]int16{}}, [30]int16{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 7, 9, 37, 5, 35, 34, 0, 0, 0, 0, 0, 0, 0, 0, 5, 3, 0}}
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
	F4 TSParseActionEntry
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
	F24 TSParseActionEntry
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
	F32 TSParseActionEntry
	F33 struct {
		F0 anon_2
		F1 [6]byte
	}
	F34 TSParseActionEntry
	F35 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F36 struct {
		F0 anon_2
		F1 [6]byte
	}
	F37 TSParseActionEntry
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
		F0 anon_2
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
	F51 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F55 TSParseActionEntry
	F56 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F59 TSParseActionEntry
	F60 struct {
		F0 anon_2
		F1 [6]byte
	}
	F61 TSParseActionEntry
	F62 struct {
		F0 anon_2
		F1 [6]byte
	}
	F63 TSParseActionEntry
	F64 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F68 TSParseActionEntry
	F69 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F79 TSParseActionEntry
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
	F85 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F89 TSParseActionEntry
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
	F93 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F94 struct {
		F0 anon_2
		F1 [6]byte
	}
	F95 TSParseActionEntry
	F96 struct {
		F0 anon_2
		F1 [6]byte
	}
	F97 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F98 struct {
		F0 anon_2
		F1 [6]byte
	}
	F99 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F108 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F109 struct {
		F0 anon_2
		F1 [6]byte
	}
	F110 TSParseActionEntry
	F111 struct {
		F0 anon_2
		F1 [6]byte
	}
	F112 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F113 struct {
		F0 anon_2
		F1 [6]byte
	}
	F114 TSParseActionEntry
	F115 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F124 TSParseActionEntry
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
	F132 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F133 struct {
		F0 anon_2
		F1 [6]byte
	}
	F134 TSParseActionEntry
	F135 struct {
		F0 anon_2
		F1 [6]byte
	}
	F136 TSParseActionEntry
	F137 struct {
		F0 anon_2
		F1 [6]byte
	}
	F138 TSParseActionEntry
	F139 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F142 TSParseActionEntry
	F143 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F148 TSParseActionEntry
	F149 struct {
		F0 anon_2
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
		F0 anon_2
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
	F158 TSParseActionEntry
	F159 struct {
		F0 anon_2
		F1 [6]byte
	}
	F160 TSParseActionEntry
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
	F4 TSParseActionEntry
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
	F24 TSParseActionEntry
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
	F32 TSParseActionEntry
	F33 struct {
		F0 anon_2
		F1 [6]byte
	}
	F34 TSParseActionEntry
	F35 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F36 struct {
		F0 anon_2
		F1 [6]byte
	}
	F37 TSParseActionEntry
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
		F0 anon_2
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
	F51 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F55 TSParseActionEntry
	F56 struct {
		F0 anon_2
		F1 [6]byte
	}
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
		F0 anon_2
		F1 [6]byte
	}
	F59 TSParseActionEntry
	F60 struct {
		F0 anon_2
		F1 [6]byte
	}
	F61 TSParseActionEntry
	F62 struct {
		F0 anon_2
		F1 [6]byte
	}
	F63 TSParseActionEntry
	F64 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F68 TSParseActionEntry
	F69 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F79 TSParseActionEntry
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
	F85 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F89 TSParseActionEntry
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
	F93 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F94 struct {
		F0 anon_2
		F1 [6]byte
	}
	F95 TSParseActionEntry
	F96 struct {
		F0 anon_2
		F1 [6]byte
	}
	F97 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F98 struct {
		F0 anon_2
		F1 [6]byte
	}
	F99 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F108 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F109 struct {
		F0 anon_2
		F1 [6]byte
	}
	F110 TSParseActionEntry
	F111 struct {
		F0 anon_2
		F1 [6]byte
	}
	F112 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F113 struct {
		F0 anon_2
		F1 [6]byte
	}
	F114 TSParseActionEntry
	F115 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F124 TSParseActionEntry
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
	F132 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F133 struct {
		F0 anon_2
		F1 [6]byte
	}
	F134 TSParseActionEntry
	F135 struct {
		F0 anon_2
		F1 [6]byte
	}
	F136 TSParseActionEntry
	F137 struct {
		F0 anon_2
		F1 [6]byte
	}
	F138 TSParseActionEntry
	F139 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F142 TSParseActionEntry
	F143 struct {
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F148 TSParseActionEntry
	F149 struct {
		F0 anon_2
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
		F0 anon_2
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
	F158 TSParseActionEntry
	F159 struct {
		F0 anon_2
		F1 [6]byte
	}
	F160 TSParseActionEntry
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
		F0 anon_2
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
		F0 anon_2
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
		F0 anon_2
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 15, 0, 0}}}, struct {
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
}{0, 0, 57, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 10, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 19, 0, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 15, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 15, 0, 0}}}, struct {
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
}{0, 0, 56, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 57, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 3, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 28, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 28, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 28, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{0, 0, 17, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 24, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 25, 0, 0}}}, struct {
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
}{0, 0, 14, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 16, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 16, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 29, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 29, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 29, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 29, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 29, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 24, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 29, 0, 3}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 29, 0, 3}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 19, 0, 1}}}, struct {
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
}{0, 0, 52, 0, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 19, 0, 1}}}, struct {
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
}{0, 0, 58, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 19, 0, 4}}}, struct {
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
}{0, 0, 26, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 19, 0, 6}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 28, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 26, 0, 1}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 19, 0, 4}}}, struct {
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
}{0, 0, 27, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 28, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 24, 0, 5}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 19, 0, 6}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 17, 0, 7}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 18, 0, 2}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 15, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 20, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 21, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 19, 0, 1}}}, struct {
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
}{0, 0, 42, 0, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 23, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 19, 0, 4}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 19, 0, 6}}}, struct {
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
}{0, 0, 47, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 33, 0, 0}, [2]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [2]byte = [2]byte{45, 0}
var _str_4 [2]byte = [2]byte{58, 0}
var _str_5 [2]byte = [2]byte{96, 0}
var _str_6 [2]byte = [2]byte{34, 0}
var _str_7 [6]byte = [6]byte{69, 82, 82, 79, 82, 0}
var _str_8 [8]byte = [8]byte{77, 73, 83, 83, 73, 78, 71, 0}
var _str_9 [15]byte = [15]byte{99, 111, 110, 116, 101, 110, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_10 [12]byte = [12]byte{95, 105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}
var _str_11 [8]byte = [8]byte{95, 101, 115, 99, 97, 112, 101, 0}
var _str_12 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}
var _str_13 [4]byte = [4]byte{226, 128, 162, 0}
var _str_14 [7]byte = [7]byte{95, 115, 112, 97, 99, 101, 0}
var _str_15 [5]byte = [5]byte{95, 101, 111, 108, 0}
var _str_16 [4]byte = [4]byte{99, 115, 116, 0}
var _str_17 [6]byte = [6]byte{95, 108, 105, 110, 101, 0}
var _str_18 [6]byte = [6]byte{114, 97, 110, 103, 101, 0}
var _str_19 [9]byte = [9]byte{112, 111, 115, 105, 116, 105, 111, 110, 0}
var _str_20 [5]byte = [5]byte{110, 111, 100, 101, 0}
var _str_21 [5]byte = [5]byte{116, 101, 120, 116, 0}
var _str_22 [8]byte = [8]byte{108, 105, 116, 101, 114, 97, 108, 0}
var _str_23 [6]byte = [6]byte{101, 114, 114, 111, 114, 0}
var _str_24 [8]byte = [8]byte{109, 105, 115, 115, 105, 110, 103, 0}
var _str_25 [6]byte = [6]byte{102, 105, 101, 108, 100, 0}
var _str_26 [8]byte = [8]byte{99, 111, 110, 116, 101, 110, 116, 0}
var _str_27 [6]byte = [6]byte{95, 109, 97, 114, 107, 0}
var _str_28 [12]byte = [12]byte{99, 115, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_29 [14]byte = [14]byte{95, 108, 105, 110, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_30 [16]byte = [16]byte{99, 111, 110, 116, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_31 [5]byte = [5]byte{107, 105, 110, 100, 0}
var _str_32 [7]byte = [7]byte{99, 111, 108, 117, 109, 110, 0}
var _str_33 [7]byte = [7]byte{101, 115, 99, 97, 112, 101, 0}
var _str_34 [5]byte = [5]byte{110, 97, 109, 101, 0}
var _str_35 [4]byte = [4]byte{114, 111, 119, 0}
var _str_36 [6]byte = [6]byte{115, 116, 97, 114, 116, 0}
var ts_lex_map [18]int16 = [18]int16{10, 33, 13, 34, 32, 16, 34, 12, 45, 9, 58, 15, 92, 17, 96, 11, 8226, 31}
var ts_lex_map_37 [18]int16 = [18]int16{34, 29, 48, 29, 92, 29, 96, 29, 102, 29, 110, 29, 114, 29, 116, 29, 118, 29}

func init() {
	tree_sitter_cst_language = struct {
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
	}{14, 30, 1, 15, 0, 59, 2, 8, 6, 5, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), libc.Ptr(&ts_field_names), libc.Ptr(&ts_field_map_slices), libc.Ptr(&ts_field_map_entries), libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), nil, nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{}, [5]byte{}}
}
func tree_sitter_cst() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_cst_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, loadedv18, cmp20, cmp24, cmp28, cmp32, cmp36, cmp40, cmp44, cmp46, cmp48, cmp51, cmp54, cmp57, cmp60, loadedv64, cmp66, cmp70, cmp74, cmp77, cmp80, loadedv84, cmp86, cmp90, cmp93, cmp96, cmp99, cmp102, cmp105, cmp108, cmp111, loadedv115, cmp117, cmp121, cmp125, cmp128, cmp131, loadedv135, cmp137, cmp141, cmp144, cmp147, loadedv151, cmp153, cmp157, cmp160, cmp163, cmp166, cmp169, cmp172, cmp175, cmp178, loadedv182, loadedv184, cmp187, cmp191, cmp195, cmp199, cmp203, cmp207, cmp210, loadedv214, loadedv216, loadedv220, loadedv224, loadedv228, loadedv232, cmp236, cmp239, cmp242, cmp245, cmp248, cmp251, cmp254, cmp257, cmp260, loadedv264, cmp268, cmp271, cmp274, cmp277, cmp280, cmp283, cmp286, cmp289, cmp292, loadedv296, loadedv300, loadedv304, cmp311, cmp317, loadedv327, cmp331, cmp335, cmp338, cmp341, cmp344, cmp347, cmp350, cmp353, cmp356, cmp359, loadedv363, cmp367, cmp371, cmp374, cmp377, cmp380, cmp383, cmp386, cmp389, cmp392, cmp395, loadedv399, cmp403, cmp407, cmp410, cmp413, cmp416, cmp419, cmp422, cmp425, cmp428, cmp431, loadedv435, cmp439, cmp443, cmp446, cmp449, cmp452, cmp455, cmp458, cmp461, cmp464, cmp467, loadedv471, cmp475, cmp479, cmp482, cmp485, cmp488, cmp491, cmp494, cmp497, cmp500, cmp503, loadedv507, cmp511, cmp515, cmp518, cmp521, cmp524, cmp527, cmp530, cmp533, cmp536, cmp539, loadedv543, cmp547, cmp551, cmp554, cmp557, cmp560, cmp563, cmp566, cmp569, cmp572, cmp575, loadedv579, cmp583, cmp587, cmp590, cmp593, cmp596, cmp599, cmp602, cmp605, cmp608, cmp611, loadedv615, cmp619, cmp623, cmp626, cmp629, cmp632, cmp635, cmp638, cmp641, cmp644, cmp647, loadedv651, cmp655, cmp659, cmp662, cmp665, cmp668, cmp671, cmp674, cmp677, cmp680, cmp683, loadedv687, cmp691, cmp694, cmp697, cmp700, cmp703, cmp706, cmp709, cmp712, cmp715, loadedv719, loadedv723, cmp727, cmp730, loadedv734, loadedv738, loadedv742, loadedv746, cmp750, loadedv754, v352 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v149, v152 int16
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol218, result_symbol222, result_symbol226, result_symbol230, result_symbol234, result_symbol266, result_symbol298, result_symbol302, result_symbol306, arrayidx315, arrayidx322, result_symbol329, result_symbol365, result_symbol401, result_symbol437, result_symbol473, result_symbol509, result_symbol545, result_symbol581, result_symbol617, result_symbol653, result_symbol689, result_symbol721, result_symbol725, result_symbol736, result_symbol740, result_symbol744, result_symbol748 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v34, v35, v36, v37, v38, v40, v41, v42, v43, v44, v45, v46, v47, v48, v50, v51, v52, v53, v54, v56, v57, v58, v59, v61, v62, v63, v64, v65, v66, v67, v68, v69, v72, v73, v74, v75, v76, v77, v78, v109, v110, v111, v112, v113, v114, v115, v116, v117, v123, v124, v125, v126, v127, v128, v129, v130, v131, v147, v148, conv316, v150, v151, add320, v153, add325, v159, v160, v161, v162, v163, v164, v165, v166, v167, v168, v174, v175, v176, v177, v178, v179, v180, v181, v182, v183, v189, v190, v191, v192, v193, v194, v195, v196, v197, v198, v204, v205, v206, v207, v208, v209, v210, v211, v212, v213, v219, v220, v221, v222, v223, v224, v225, v226, v227, v228, v234, v235, v236, v237, v238, v239, v240, v241, v242, v243, v249, v250, v251, v252, v253, v254, v255, v256, v257, v258, v264, v265, v266, v267, v268, v269, v270, v271, v272, v273, v279, v280, v281, v282, v283, v284, v285, v286, v287, v288, v294, v295, v296, v297, v298, v299, v300, v301, v302, v303, v309, v310, v311, v312, v313, v314, v315, v316, v317, v328, v329, v350 int32
	var lookahead, i, i308, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv310, idxprom314, idxprom321 int64
	var v3, storedv, v10, v19, v33, v39, v49, v55, v60, v70, v71, v79, v84, v89, v94, v99, v104, v118, v132, v137, v142, v154, v169, v184, v199, v214, v229, v244, v259, v274, v289, v304, v318, v323, v330, v335, v340, v345, v351 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v80, v81, v82, v83, v85, v86, v87, v88, v90, v91, v92, v93, v95, v96, v97, v98, v100, v101, v102, v103, v105, v106, v107, v108, v119, v120, v121, v122, v133, v134, v135, v136, v138, v139, v140, v141, v143, v144, v145, v146, v155, v156, v157, v158, v170, v171, v172, v173, v185, v186, v187, v188, v200, v201, v202, v203, v215, v216, v217, v218, v230, v231, v232, v233, v245, v246, v247, v248, v260, v261, v262, v263, v275, v276, v277, v278, v290, v291, v292, v293, v305, v306, v307, v308, v319, v320, v321, v322, v324, v325, v326, v327, v331, v332, v333, v334, v336, v337, v338, v339, v341, v342, v343, v344, v346, v347, v348, v349 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end219, mark_end223, mark_end227, mark_end231, mark_end235, mark_end267, mark_end299, mark_end303, mark_end307, mark_end330, mark_end366, mark_end402, mark_end438, mark_end474, mark_end510, mark_end546, mark_end582, mark_end618, mark_end654, mark_end690, mark_end722, mark_end726, mark_end737, mark_end741, mark_end745, mark_end749 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i308, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, loadedv18, v20, cmp20, v21, cmp24, v22, cmp28, v23, cmp32, v24, cmp36, v25, cmp40, v26, cmp44, v27, cmp46, v28, cmp48, v29, cmp51, v30, cmp54, v31, cmp57, v32, cmp60, v33, loadedv64, v34, cmp66, v35, cmp70, v36, cmp74, v37, cmp77, v38, cmp80, v39, loadedv84, v40, cmp86, v41, cmp90, v42, cmp93, v43, cmp96, v44, cmp99, v45, cmp102, v46, cmp105, v47, cmp108, v48, cmp111, v49, loadedv115, v50, cmp117, v51, cmp121, v52, cmp125, v53, cmp128, v54, cmp131, v55, loadedv135, v56, cmp137, v57, cmp141, v58, cmp144, v59, cmp147, v60, loadedv151, v61, cmp153, v62, cmp157, v63, cmp160, v64, cmp163, v65, cmp166, v66, cmp169, v67, cmp172, v68, cmp175, v69, cmp178, v70, loadedv182, v71, loadedv184, v72, cmp187, v73, cmp191, v74, cmp195, v75, cmp199, v76, cmp203, v77, cmp207, v78, cmp210, v79, loadedv214, v80, result_symbol, v81, mark_end, v82, v83, v84, loadedv216, v85, result_symbol218, v86, mark_end219, v87, v88, v89, loadedv220, v90, result_symbol222, v91, mark_end223, v92, v93, v94, loadedv224, v95, result_symbol226, v96, mark_end227, v97, v98, v99, loadedv228, v100, result_symbol230, v101, mark_end231, v102, v103, v104, loadedv232, v105, result_symbol234, v106, mark_end235, v107, v108, v109, cmp236, v110, cmp239, v111, cmp242, v112, cmp245, v113, cmp248, v114, cmp251, v115, cmp254, v116, cmp257, v117, cmp260, v118, loadedv264, v119, result_symbol266, v120, mark_end267, v121, v122, v123, cmp268, v124, cmp271, v125, cmp274, v126, cmp277, v127, cmp280, v128, cmp283, v129, cmp286, v130, cmp289, v131, cmp292, v132, loadedv296, v133, result_symbol298, v134, mark_end299, v135, v136, v137, loadedv300, v138, result_symbol302, v139, mark_end303, v140, v141, v142, loadedv304, v143, result_symbol306, v144, mark_end307, v145, v146, v147, conv310, cmp311, v148, idxprom314, arrayidx315, v149, conv316, v150, cmp317, v151, add320, idxprom321, arrayidx322, v152, v153, add325, v154, loadedv327, v155, result_symbol329, v156, mark_end330, v157, v158, v159, cmp331, v160, cmp335, v161, cmp338, v162, cmp341, v163, cmp344, v164, cmp347, v165, cmp350, v166, cmp353, v167, cmp356, v168, cmp359, v169, loadedv363, v170, result_symbol365, v171, mark_end366, v172, v173, v174, cmp367, v175, cmp371, v176, cmp374, v177, cmp377, v178, cmp380, v179, cmp383, v180, cmp386, v181, cmp389, v182, cmp392, v183, cmp395, v184, loadedv399, v185, result_symbol401, v186, mark_end402, v187, v188, v189, cmp403, v190, cmp407, v191, cmp410, v192, cmp413, v193, cmp416, v194, cmp419, v195, cmp422, v196, cmp425, v197, cmp428, v198, cmp431, v199, loadedv435, v200, result_symbol437, v201, mark_end438, v202, v203, v204, cmp439, v205, cmp443, v206, cmp446, v207, cmp449, v208, cmp452, v209, cmp455, v210, cmp458, v211, cmp461, v212, cmp464, v213, cmp467, v214, loadedv471, v215, result_symbol473, v216, mark_end474, v217, v218, v219, cmp475, v220, cmp479, v221, cmp482, v222, cmp485, v223, cmp488, v224, cmp491, v225, cmp494, v226, cmp497, v227, cmp500, v228, cmp503, v229, loadedv507, v230, result_symbol509, v231, mark_end510, v232, v233, v234, cmp511, v235, cmp515, v236, cmp518, v237, cmp521, v238, cmp524, v239, cmp527, v240, cmp530, v241, cmp533, v242, cmp536, v243, cmp539, v244, loadedv543, v245, result_symbol545, v246, mark_end546, v247, v248, v249, cmp547, v250, cmp551, v251, cmp554, v252, cmp557, v253, cmp560, v254, cmp563, v255, cmp566, v256, cmp569, v257, cmp572, v258, cmp575, v259, loadedv579, v260, result_symbol581, v261, mark_end582, v262, v263, v264, cmp583, v265, cmp587, v266, cmp590, v267, cmp593, v268, cmp596, v269, cmp599, v270, cmp602, v271, cmp605, v272, cmp608, v273, cmp611, v274, loadedv615, v275, result_symbol617, v276, mark_end618, v277, v278, v279, cmp619, v280, cmp623, v281, cmp626, v282, cmp629, v283, cmp632, v284, cmp635, v285, cmp638, v286, cmp641, v287, cmp644, v288, cmp647, v289, loadedv651, v290, result_symbol653, v291, mark_end654, v292, v293, v294, cmp655, v295, cmp659, v296, cmp662, v297, cmp665, v298, cmp668, v299, cmp671, v300, cmp674, v301, cmp677, v302, cmp680, v303, cmp683, v304, loadedv687, v305, result_symbol689, v306, mark_end690, v307, v308, v309, cmp691, v310, cmp694, v311, cmp697, v312, cmp700, v313, cmp703, v314, cmp706, v315, cmp709, v316, cmp712, v317, cmp715, v318, loadedv719, v319, result_symbol721, v320, mark_end722, v321, v322, v323, loadedv723, v324, result_symbol725, v325, mark_end726, v326, v327, v328, cmp727, v329, cmp730, v330, loadedv734, v331, result_symbol736, v332, mark_end737, v333, v334, v335, loadedv738, v336, result_symbol740, v337, mark_end741, v338, v339, v340, loadedv742, v341, result_symbol744, v342, mark_end745, v343, v344, v345, loadedv746, v346, result_symbol748, v347, mark_end749, v348, v349, v350, cmp750, v351, loadedv754, v352

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
	i308 = libc.Ptr(&new(struct {
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
		goto sw_bb19
	case 2:
		goto sw_bb65
	case 3:
		goto sw_bb85
	case 4:
		goto sw_bb116
	case 5:
		goto sw_bb136
	case 6:
		goto sw_bb152
	case 7:
		goto sw_bb183
	case 8:
		goto sw_bb215
	case 9:
		goto sw_bb217
	case 10:
		goto sw_bb221
	case 11:
		goto sw_bb225
	case 12:
		goto sw_bb229
	case 13:
		goto sw_bb233
	case 14:
		goto sw_bb265
	case 15:
		goto sw_bb297
	case 16:
		goto sw_bb301
	case 17:
		goto sw_bb305
	case 18:
		goto sw_bb328
	case 19:
		goto sw_bb364
	case 20:
		goto sw_bb400
	case 21:
		goto sw_bb436
	case 22:
		goto sw_bb472
	case 23:
		goto sw_bb508
	case 24:
		goto sw_bb544
	case 25:
		goto sw_bb580
	case 26:
		goto sw_bb616
	case 27:
		goto sw_bb652
	case 28:
		goto sw_bb688
	case 29:
		goto sw_bb720
	case 30:
		goto sw_bb724
	case 31:
		goto sw_bb735
	case 32:
		goto sw_bb739
	case 33:
		goto sw_bb743
	case 34:
		goto sw_bb747
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
	*libc.As[int16](state_addr) = 8
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
	cmp14 = v18 != 0
	if cmp14 {
		goto if_then16
	} else {
		goto if_end17
	}

if_then16:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end17:
	v19 = *libc.As[byte](result)
	loadedv18 = (v19 & 1) != 0
	*libc.As[bool](retval) = loadedv18
	goto _return

sw_bb19:
	v20 = *libc.As[int32](lookahead)
	cmp20 = v20 == 32
	if cmp20 {
		goto if_then22
	} else {
		goto if_end23
	}

if_then22:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end23:
	v21 = *libc.As[int32](lookahead)
	cmp24 = v21 == 34
	if cmp24 {
		goto if_then26
	} else {
		goto if_end27
	}

if_then26:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end27:
	v22 = *libc.As[int32](lookahead)
	cmp28 = v22 == 58
	if cmp28 {
		goto if_then30
	} else {
		goto if_end31
	}

if_then30:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end31:
	v23 = *libc.As[int32](lookahead)
	cmp32 = v23 == 77
	if cmp32 {
		goto if_then34
	} else {
		goto if_end35
	}

if_then34:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end35:
	v24 = *libc.As[int32](lookahead)
	cmp36 = v24 == 96
	if cmp36 {
		goto if_then38
	} else {
		goto if_end39
	}

if_then38:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end39:
	v25 = *libc.As[int32](lookahead)
	cmp40 = v25 == 8226
	if cmp40 {
		goto if_then42
	} else {
		goto if_end43
	}

if_then42:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end43:
	v26 = *libc.As[int32](lookahead)
	cmp44 = v26 == 45
	if cmp44 {
		goto if_then62
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v27 = *libc.As[int32](lookahead)
	cmp46 = 48 <= v27
	if cmp46 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false50
	}

land_lhs_true:
	v28 = *libc.As[int32](lookahead)
	cmp48 = v28 <= 57
	if cmp48 {
		goto if_then62
	} else {
		goto lor_lhs_false50
	}

lor_lhs_false50:
	v29 = *libc.As[int32](lookahead)
	cmp51 = 65 <= v29
	if cmp51 {
		goto land_lhs_true53
	} else {
		goto lor_lhs_false56
	}

land_lhs_true53:
	v30 = *libc.As[int32](lookahead)
	cmp54 = v30 <= 90
	if cmp54 {
		goto if_then62
	} else {
		goto lor_lhs_false56
	}

lor_lhs_false56:
	v31 = *libc.As[int32](lookahead)
	cmp57 = 95 <= v31
	if cmp57 {
		goto land_lhs_true59
	} else {
		goto if_end63
	}

land_lhs_true59:
	v32 = *libc.As[int32](lookahead)
	cmp60 = v32 <= 122
	if cmp60 {
		goto if_then62
	} else {
		goto if_end63
	}

if_then62:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end63:
	v33 = *libc.As[byte](result)
	loadedv64 = (v33 & 1) != 0
	*libc.As[bool](retval) = loadedv64
	goto _return

sw_bb65:
	v34 = *libc.As[int32](lookahead)
	cmp66 = v34 == 34
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end69:
	v35 = *libc.As[int32](lookahead)
	cmp70 = v35 == 92
	if cmp70 {
		goto if_then72
	} else {
		goto if_end73
	}

if_then72:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end73:
	v36 = *libc.As[int32](lookahead)
	cmp74 = v36 != 0
	if cmp74 {
		goto land_lhs_true76
	} else {
		goto if_end83
	}

land_lhs_true76:
	v37 = *libc.As[int32](lookahead)
	cmp77 = v37 != 10
	if cmp77 {
		goto land_lhs_true79
	} else {
		goto if_end83
	}

land_lhs_true79:
	v38 = *libc.As[int32](lookahead)
	cmp80 = v38 != 13
	if cmp80 {
		goto if_then82
	} else {
		goto if_end83
	}

if_then82:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end83:
	v39 = *libc.As[byte](result)
	loadedv84 = (v39 & 1) != 0
	*libc.As[bool](retval) = loadedv84
	goto _return

sw_bb85:
	v40 = *libc.As[int32](lookahead)
	cmp86 = v40 == 69
	if cmp86 {
		goto if_then88
	} else {
		goto if_end89
	}

if_then88:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end89:
	v41 = *libc.As[int32](lookahead)
	cmp90 = v41 == 45
	if cmp90 {
		goto if_then113
	} else {
		goto lor_lhs_false92
	}

lor_lhs_false92:
	v42 = *libc.As[int32](lookahead)
	cmp93 = 48 <= v42
	if cmp93 {
		goto land_lhs_true95
	} else {
		goto lor_lhs_false98
	}

land_lhs_true95:
	v43 = *libc.As[int32](lookahead)
	cmp96 = v43 <= 57
	if cmp96 {
		goto if_then113
	} else {
		goto lor_lhs_false98
	}

lor_lhs_false98:
	v44 = *libc.As[int32](lookahead)
	cmp99 = 65 <= v44
	if cmp99 {
		goto land_lhs_true101
	} else {
		goto lor_lhs_false104
	}

land_lhs_true101:
	v45 = *libc.As[int32](lookahead)
	cmp102 = v45 <= 90
	if cmp102 {
		goto if_then113
	} else {
		goto lor_lhs_false104
	}

lor_lhs_false104:
	v46 = *libc.As[int32](lookahead)
	cmp105 = v46 == 95
	if cmp105 {
		goto if_then113
	} else {
		goto lor_lhs_false107
	}

lor_lhs_false107:
	v47 = *libc.As[int32](lookahead)
	cmp108 = 97 <= v47
	if cmp108 {
		goto land_lhs_true110
	} else {
		goto if_end114
	}

land_lhs_true110:
	v48 = *libc.As[int32](lookahead)
	cmp111 = v48 <= 122
	if cmp111 {
		goto if_then113
	} else {
		goto if_end114
	}

if_then113:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end114:
	v49 = *libc.As[byte](result)
	loadedv115 = (v49 & 1) != 0
	*libc.As[bool](retval) = loadedv115
	goto _return

sw_bb116:
	v50 = *libc.As[int32](lookahead)
	cmp117 = v50 == 92
	if cmp117 {
		goto if_then119
	} else {
		goto if_end120
	}

if_then119:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end120:
	v51 = *libc.As[int32](lookahead)
	cmp121 = v51 == 96
	if cmp121 {
		goto if_then123
	} else {
		goto if_end124
	}

if_then123:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end124:
	v52 = *libc.As[int32](lookahead)
	cmp125 = v52 != 0
	if cmp125 {
		goto land_lhs_true127
	} else {
		goto if_end134
	}

land_lhs_true127:
	v53 = *libc.As[int32](lookahead)
	cmp128 = v53 != 10
	if cmp128 {
		goto land_lhs_true130
	} else {
		goto if_end134
	}

land_lhs_true130:
	v54 = *libc.As[int32](lookahead)
	cmp131 = v54 != 13
	if cmp131 {
		goto if_then133
	} else {
		goto if_end134
	}

if_then133:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end134:
	v55 = *libc.As[byte](result)
	loadedv135 = (v55 & 1) != 0
	*libc.As[bool](retval) = loadedv135
	goto _return

sw_bb136:
	v56 = *libc.As[int32](lookahead)
	cmp137 = v56 == 92
	if cmp137 {
		goto if_then139
	} else {
		goto if_end140
	}

if_then139:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end140:
	v57 = *libc.As[int32](lookahead)
	cmp141 = v57 != 0
	if cmp141 {
		goto land_lhs_true143
	} else {
		goto if_end150
	}

land_lhs_true143:
	v58 = *libc.As[int32](lookahead)
	cmp144 = v58 != 10
	if cmp144 {
		goto land_lhs_true146
	} else {
		goto if_end150
	}

land_lhs_true146:
	v59 = *libc.As[int32](lookahead)
	cmp147 = v59 != 13
	if cmp147 {
		goto if_then149
	} else {
		goto if_end150
	}

if_then149:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end150:
	v60 = *libc.As[byte](result)
	loadedv151 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv151
	goto _return

sw_bb152:
	v61 = *libc.As[int32](lookahead)
	cmp153 = v61 == 8226
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end156:
	v62 = *libc.As[int32](lookahead)
	cmp157 = v62 == 45
	if cmp157 {
		goto if_then180
	} else {
		goto lor_lhs_false159
	}

lor_lhs_false159:
	v63 = *libc.As[int32](lookahead)
	cmp160 = 48 <= v63
	if cmp160 {
		goto land_lhs_true162
	} else {
		goto lor_lhs_false165
	}

land_lhs_true162:
	v64 = *libc.As[int32](lookahead)
	cmp163 = v64 <= 57
	if cmp163 {
		goto if_then180
	} else {
		goto lor_lhs_false165
	}

lor_lhs_false165:
	v65 = *libc.As[int32](lookahead)
	cmp166 = 65 <= v65
	if cmp166 {
		goto land_lhs_true168
	} else {
		goto lor_lhs_false171
	}

land_lhs_true168:
	v66 = *libc.As[int32](lookahead)
	cmp169 = v66 <= 90
	if cmp169 {
		goto if_then180
	} else {
		goto lor_lhs_false171
	}

lor_lhs_false171:
	v67 = *libc.As[int32](lookahead)
	cmp172 = v67 == 95
	if cmp172 {
		goto if_then180
	} else {
		goto lor_lhs_false174
	}

lor_lhs_false174:
	v68 = *libc.As[int32](lookahead)
	cmp175 = 97 <= v68
	if cmp175 {
		goto land_lhs_true177
	} else {
		goto if_end181
	}

land_lhs_true177:
	v69 = *libc.As[int32](lookahead)
	cmp178 = v69 <= 122
	if cmp178 {
		goto if_then180
	} else {
		goto if_end181
	}

if_then180:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end181:
	v70 = *libc.As[byte](result)
	loadedv182 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv182
	goto _return

sw_bb183:
	v71 = *libc.As[byte](eof)
	loadedv184 = (v71 & 1) != 0
	if loadedv184 {
		goto if_then185
	} else {
		goto if_end186
	}

if_then185:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end186:
	v72 = *libc.As[int32](lookahead)
	cmp187 = v72 == 10
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end190:
	v73 = *libc.As[int32](lookahead)
	cmp191 = v73 == 13
	if cmp191 {
		goto if_then193
	} else {
		goto if_end194
	}

if_then193:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end194:
	v74 = *libc.As[int32](lookahead)
	cmp195 = v74 == 32
	if cmp195 {
		goto if_then197
	} else {
		goto if_end198
	}

if_then197:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end198:
	v75 = *libc.As[int32](lookahead)
	cmp199 = v75 == 45
	if cmp199 {
		goto if_then201
	} else {
		goto if_end202
	}

if_then201:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end202:
	v76 = *libc.As[int32](lookahead)
	cmp203 = v76 == 58
	if cmp203 {
		goto if_then205
	} else {
		goto if_end206
	}

if_then205:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end206:
	v77 = *libc.As[int32](lookahead)
	cmp207 = 48 <= v77
	if cmp207 {
		goto land_lhs_true209
	} else {
		goto if_end213
	}

land_lhs_true209:
	v78 = *libc.As[int32](lookahead)
	cmp210 = v78 <= 57
	if cmp210 {
		goto if_then212
	} else {
		goto if_end213
	}

if_then212:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end213:
	v79 = *libc.As[byte](result)
	loadedv214 = (v79 & 1) != 0
	*libc.As[bool](retval) = loadedv214
	goto _return

sw_bb215:
	*libc.As[byte](result) = 1
	v80 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v80).F1)
	*libc.As[int16](result_symbol) = 0
	v81 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v81).F3)
	v82 = *libc.As[unsafe.Pointer](mark_end)
	v83 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v82)(v83)
	v84 = *libc.As[byte](result)
	loadedv216 = (v84 & 1) != 0
	*libc.As[bool](retval) = loadedv216
	goto _return

sw_bb217:
	*libc.As[byte](result) = 1
	v85 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol218 = libc.Ptr(&libc.As[TSLexer](v85).F1)
	*libc.As[int16](result_symbol218) = 1
	v86 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end219 = libc.Ptr(&libc.As[TSLexer](v86).F3)
	v87 = *libc.As[unsafe.Pointer](mark_end219)
	v88 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v87)(v88)
	v89 = *libc.As[byte](result)
	loadedv220 = (v89 & 1) != 0
	*libc.As[bool](retval) = loadedv220
	goto _return

sw_bb221:
	*libc.As[byte](result) = 1
	v90 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol222 = libc.Ptr(&libc.As[TSLexer](v90).F1)
	*libc.As[int16](result_symbol222) = 2
	v91 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end223 = libc.Ptr(&libc.As[TSLexer](v91).F3)
	v92 = *libc.As[unsafe.Pointer](mark_end223)
	v93 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v92)(v93)
	v94 = *libc.As[byte](result)
	loadedv224 = (v94 & 1) != 0
	*libc.As[bool](retval) = loadedv224
	goto _return

sw_bb225:
	*libc.As[byte](result) = 1
	v95 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol226 = libc.Ptr(&libc.As[TSLexer](v95).F1)
	*libc.As[int16](result_symbol226) = 3
	v96 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end227 = libc.Ptr(&libc.As[TSLexer](v96).F3)
	v97 = *libc.As[unsafe.Pointer](mark_end227)
	v98 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v97)(v98)
	v99 = *libc.As[byte](result)
	loadedv228 = (v99 & 1) != 0
	*libc.As[bool](retval) = loadedv228
	goto _return

sw_bb229:
	*libc.As[byte](result) = 1
	v100 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol230 = libc.Ptr(&libc.As[TSLexer](v100).F1)
	*libc.As[int16](result_symbol230) = 4
	v101 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end231 = libc.Ptr(&libc.As[TSLexer](v101).F3)
	v102 = *libc.As[unsafe.Pointer](mark_end231)
	v103 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v102)(v103)
	v104 = *libc.As[byte](result)
	loadedv232 = (v104 & 1) != 0
	*libc.As[bool](retval) = loadedv232
	goto _return

sw_bb233:
	*libc.As[byte](result) = 1
	v105 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol234 = libc.Ptr(&libc.As[TSLexer](v105).F1)
	*libc.As[int16](result_symbol234) = 5
	v106 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end235 = libc.Ptr(&libc.As[TSLexer](v106).F3)
	v107 = *libc.As[unsafe.Pointer](mark_end235)
	v108 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v107)(v108)
	v109 = *libc.As[int32](lookahead)
	cmp236 = v109 == 45
	if cmp236 {
		goto if_then262
	} else {
		goto lor_lhs_false238
	}

lor_lhs_false238:
	v110 = *libc.As[int32](lookahead)
	cmp239 = v110 == 46
	if cmp239 {
		goto if_then262
	} else {
		goto lor_lhs_false241
	}

lor_lhs_false241:
	v111 = *libc.As[int32](lookahead)
	cmp242 = 48 <= v111
	if cmp242 {
		goto land_lhs_true244
	} else {
		goto lor_lhs_false247
	}

land_lhs_true244:
	v112 = *libc.As[int32](lookahead)
	cmp245 = v112 <= 57
	if cmp245 {
		goto if_then262
	} else {
		goto lor_lhs_false247
	}

lor_lhs_false247:
	v113 = *libc.As[int32](lookahead)
	cmp248 = 65 <= v113
	if cmp248 {
		goto land_lhs_true250
	} else {
		goto lor_lhs_false253
	}

land_lhs_true250:
	v114 = *libc.As[int32](lookahead)
	cmp251 = v114 <= 90
	if cmp251 {
		goto if_then262
	} else {
		goto lor_lhs_false253
	}

lor_lhs_false253:
	v115 = *libc.As[int32](lookahead)
	cmp254 = v115 == 95
	if cmp254 {
		goto if_then262
	} else {
		goto lor_lhs_false256
	}

lor_lhs_false256:
	v116 = *libc.As[int32](lookahead)
	cmp257 = 97 <= v116
	if cmp257 {
		goto land_lhs_true259
	} else {
		goto if_end263
	}

land_lhs_true259:
	v117 = *libc.As[int32](lookahead)
	cmp260 = v117 <= 122
	if cmp260 {
		goto if_then262
	} else {
		goto if_end263
	}

if_then262:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end263:
	v118 = *libc.As[byte](result)
	loadedv264 = (v118 & 1) != 0
	*libc.As[bool](retval) = loadedv264
	goto _return

sw_bb265:
	*libc.As[byte](result) = 1
	v119 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol266 = libc.Ptr(&libc.As[TSLexer](v119).F1)
	*libc.As[int16](result_symbol266) = 6
	v120 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end267 = libc.Ptr(&libc.As[TSLexer](v120).F3)
	v121 = *libc.As[unsafe.Pointer](mark_end267)
	v122 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v121)(v122)
	v123 = *libc.As[int32](lookahead)
	cmp268 = v123 == 45
	if cmp268 {
		goto if_then294
	} else {
		goto lor_lhs_false270
	}

lor_lhs_false270:
	v124 = *libc.As[int32](lookahead)
	cmp271 = v124 == 46
	if cmp271 {
		goto if_then294
	} else {
		goto lor_lhs_false273
	}

lor_lhs_false273:
	v125 = *libc.As[int32](lookahead)
	cmp274 = 48 <= v125
	if cmp274 {
		goto land_lhs_true276
	} else {
		goto lor_lhs_false279
	}

land_lhs_true276:
	v126 = *libc.As[int32](lookahead)
	cmp277 = v126 <= 57
	if cmp277 {
		goto if_then294
	} else {
		goto lor_lhs_false279
	}

lor_lhs_false279:
	v127 = *libc.As[int32](lookahead)
	cmp280 = 65 <= v127
	if cmp280 {
		goto land_lhs_true282
	} else {
		goto lor_lhs_false285
	}

land_lhs_true282:
	v128 = *libc.As[int32](lookahead)
	cmp283 = v128 <= 90
	if cmp283 {
		goto if_then294
	} else {
		goto lor_lhs_false285
	}

lor_lhs_false285:
	v129 = *libc.As[int32](lookahead)
	cmp286 = v129 == 95
	if cmp286 {
		goto if_then294
	} else {
		goto lor_lhs_false288
	}

lor_lhs_false288:
	v130 = *libc.As[int32](lookahead)
	cmp289 = 97 <= v130
	if cmp289 {
		goto land_lhs_true291
	} else {
		goto if_end295
	}

land_lhs_true291:
	v131 = *libc.As[int32](lookahead)
	cmp292 = v131 <= 122
	if cmp292 {
		goto if_then294
	} else {
		goto if_end295
	}

if_then294:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end295:
	v132 = *libc.As[byte](result)
	loadedv296 = (v132 & 1) != 0
	*libc.As[bool](retval) = loadedv296
	goto _return

sw_bb297:
	*libc.As[byte](result) = 1
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol298 = libc.Ptr(&libc.As[TSLexer](v133).F1)
	*libc.As[int16](result_symbol298) = 7
	v134 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end299 = libc.Ptr(&libc.As[TSLexer](v134).F3)
	v135 = *libc.As[unsafe.Pointer](mark_end299)
	v136 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v135)(v136)
	v137 = *libc.As[byte](result)
	loadedv300 = (v137 & 1) != 0
	*libc.As[bool](retval) = loadedv300
	goto _return

sw_bb301:
	*libc.As[byte](result) = 1
	v138 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol302 = libc.Ptr(&libc.As[TSLexer](v138).F1)
	*libc.As[int16](result_symbol302) = 8
	v139 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end303 = libc.Ptr(&libc.As[TSLexer](v139).F3)
	v140 = *libc.As[unsafe.Pointer](mark_end303)
	v141 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v140)(v141)
	v142 = *libc.As[byte](result)
	loadedv304 = (v142 & 1) != 0
	*libc.As[bool](retval) = loadedv304
	goto _return

sw_bb305:
	*libc.As[byte](result) = 1
	v143 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol306 = libc.Ptr(&libc.As[TSLexer](v143).F1)
	*libc.As[int16](result_symbol306) = 8
	v144 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end307 = libc.Ptr(&libc.As[TSLexer](v144).F3)
	v145 = *libc.As[unsafe.Pointer](mark_end307)
	v146 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v145)(v146)
	*libc.As[int32](i308) = 0
	goto for_cond309

for_cond309:
	v147 = *libc.As[int32](i308)
	conv310 = int64(uint64(uint32(v147)))
	cmp311 = uint64(conv310) < uint64(18)
	if cmp311 {
		goto for_body313
	} else {
		goto for_end326
	}

for_body313:
	v148 = *libc.As[int32](i308)
	idxprom314 = int64(uint64(uint32(v148)))
	arrayidx315 = libc.Ptr(&ts_lex_map_37[idxprom314])
	v149 = *libc.As[int16](arrayidx315)
	conv316 = int32(uint32(uint16(v149)))
	v150 = *libc.As[int32](lookahead)
	cmp317 = conv316 == v150
	if cmp317 {
		goto if_then319
	} else {
		goto if_end323
	}

if_then319:
	v151 = *libc.As[int32](i308)
	add320 = v151 + 1
	idxprom321 = int64(uint64(uint32(add320)))
	arrayidx322 = libc.Ptr(&ts_lex_map_37[idxprom321])
	v152 = *libc.As[int16](arrayidx322)
	*libc.As[int16](state_addr) = v152
	goto next_state

if_end323:
	goto for_inc324

for_inc324:
	v153 = *libc.As[int32](i308)
	add325 = v153 + 2
	*libc.As[int32](i308) = add325
	goto for_cond309

for_end326:
	v154 = *libc.As[byte](result)
	loadedv327 = (v154 & 1) != 0
	*libc.As[bool](retval) = loadedv327
	goto _return

sw_bb328:
	*libc.As[byte](result) = 1
	v155 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol329 = libc.Ptr(&libc.As[TSLexer](v155).F1)
	*libc.As[int16](result_symbol329) = 9
	v156 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end330 = libc.Ptr(&libc.As[TSLexer](v156).F3)
	v157 = *libc.As[unsafe.Pointer](mark_end330)
	v158 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v157)(v158)
	v159 = *libc.As[int32](lookahead)
	cmp331 = v159 == 71
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end334:
	v160 = *libc.As[int32](lookahead)
	cmp335 = v160 == 45
	if cmp335 {
		goto if_then361
	} else {
		goto lor_lhs_false337
	}

lor_lhs_false337:
	v161 = *libc.As[int32](lookahead)
	cmp338 = v161 == 46
	if cmp338 {
		goto if_then361
	} else {
		goto lor_lhs_false340
	}

lor_lhs_false340:
	v162 = *libc.As[int32](lookahead)
	cmp341 = 48 <= v162
	if cmp341 {
		goto land_lhs_true343
	} else {
		goto lor_lhs_false346
	}

land_lhs_true343:
	v163 = *libc.As[int32](lookahead)
	cmp344 = v163 <= 57
	if cmp344 {
		goto if_then361
	} else {
		goto lor_lhs_false346
	}

lor_lhs_false346:
	v164 = *libc.As[int32](lookahead)
	cmp347 = 65 <= v164
	if cmp347 {
		goto land_lhs_true349
	} else {
		goto lor_lhs_false352
	}

land_lhs_true349:
	v165 = *libc.As[int32](lookahead)
	cmp350 = v165 <= 90
	if cmp350 {
		goto if_then361
	} else {
		goto lor_lhs_false352
	}

lor_lhs_false352:
	v166 = *libc.As[int32](lookahead)
	cmp353 = v166 == 95
	if cmp353 {
		goto if_then361
	} else {
		goto lor_lhs_false355
	}

lor_lhs_false355:
	v167 = *libc.As[int32](lookahead)
	cmp356 = 97 <= v167
	if cmp356 {
		goto land_lhs_true358
	} else {
		goto if_end362
	}

land_lhs_true358:
	v168 = *libc.As[int32](lookahead)
	cmp359 = v168 <= 122
	if cmp359 {
		goto if_then361
	} else {
		goto if_end362
	}

if_then361:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end362:
	v169 = *libc.As[byte](result)
	loadedv363 = (v169 & 1) != 0
	*libc.As[bool](retval) = loadedv363
	goto _return

sw_bb364:
	*libc.As[byte](result) = 1
	v170 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol365 = libc.Ptr(&libc.As[TSLexer](v170).F1)
	*libc.As[int16](result_symbol365) = 9
	v171 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end366 = libc.Ptr(&libc.As[TSLexer](v171).F3)
	v172 = *libc.As[unsafe.Pointer](mark_end366)
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v172)(v173)
	v174 = *libc.As[int32](lookahead)
	cmp367 = v174 == 73
	if cmp367 {
		goto if_then369
	} else {
		goto if_end370
	}

if_then369:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end370:
	v175 = *libc.As[int32](lookahead)
	cmp371 = v175 == 45
	if cmp371 {
		goto if_then397
	} else {
		goto lor_lhs_false373
	}

lor_lhs_false373:
	v176 = *libc.As[int32](lookahead)
	cmp374 = v176 == 46
	if cmp374 {
		goto if_then397
	} else {
		goto lor_lhs_false376
	}

lor_lhs_false376:
	v177 = *libc.As[int32](lookahead)
	cmp377 = 48 <= v177
	if cmp377 {
		goto land_lhs_true379
	} else {
		goto lor_lhs_false382
	}

land_lhs_true379:
	v178 = *libc.As[int32](lookahead)
	cmp380 = v178 <= 57
	if cmp380 {
		goto if_then397
	} else {
		goto lor_lhs_false382
	}

lor_lhs_false382:
	v179 = *libc.As[int32](lookahead)
	cmp383 = 65 <= v179
	if cmp383 {
		goto land_lhs_true385
	} else {
		goto lor_lhs_false388
	}

land_lhs_true385:
	v180 = *libc.As[int32](lookahead)
	cmp386 = v180 <= 90
	if cmp386 {
		goto if_then397
	} else {
		goto lor_lhs_false388
	}

lor_lhs_false388:
	v181 = *libc.As[int32](lookahead)
	cmp389 = v181 == 95
	if cmp389 {
		goto if_then397
	} else {
		goto lor_lhs_false391
	}

lor_lhs_false391:
	v182 = *libc.As[int32](lookahead)
	cmp392 = 97 <= v182
	if cmp392 {
		goto land_lhs_true394
	} else {
		goto if_end398
	}

land_lhs_true394:
	v183 = *libc.As[int32](lookahead)
	cmp395 = v183 <= 122
	if cmp395 {
		goto if_then397
	} else {
		goto if_end398
	}

if_then397:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end398:
	v184 = *libc.As[byte](result)
	loadedv399 = (v184 & 1) != 0
	*libc.As[bool](retval) = loadedv399
	goto _return

sw_bb400:
	*libc.As[byte](result) = 1
	v185 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol401 = libc.Ptr(&libc.As[TSLexer](v185).F1)
	*libc.As[int16](result_symbol401) = 9
	v186 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end402 = libc.Ptr(&libc.As[TSLexer](v186).F3)
	v187 = *libc.As[unsafe.Pointer](mark_end402)
	v188 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v187)(v188)
	v189 = *libc.As[int32](lookahead)
	cmp403 = v189 == 73
	if cmp403 {
		goto if_then405
	} else {
		goto if_end406
	}

if_then405:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end406:
	v190 = *libc.As[int32](lookahead)
	cmp407 = v190 == 45
	if cmp407 {
		goto if_then433
	} else {
		goto lor_lhs_false409
	}

lor_lhs_false409:
	v191 = *libc.As[int32](lookahead)
	cmp410 = v191 == 46
	if cmp410 {
		goto if_then433
	} else {
		goto lor_lhs_false412
	}

lor_lhs_false412:
	v192 = *libc.As[int32](lookahead)
	cmp413 = 48 <= v192
	if cmp413 {
		goto land_lhs_true415
	} else {
		goto lor_lhs_false418
	}

land_lhs_true415:
	v193 = *libc.As[int32](lookahead)
	cmp416 = v193 <= 57
	if cmp416 {
		goto if_then433
	} else {
		goto lor_lhs_false418
	}

lor_lhs_false418:
	v194 = *libc.As[int32](lookahead)
	cmp419 = 65 <= v194
	if cmp419 {
		goto land_lhs_true421
	} else {
		goto lor_lhs_false424
	}

land_lhs_true421:
	v195 = *libc.As[int32](lookahead)
	cmp422 = v195 <= 90
	if cmp422 {
		goto if_then433
	} else {
		goto lor_lhs_false424
	}

lor_lhs_false424:
	v196 = *libc.As[int32](lookahead)
	cmp425 = v196 == 95
	if cmp425 {
		goto if_then433
	} else {
		goto lor_lhs_false427
	}

lor_lhs_false427:
	v197 = *libc.As[int32](lookahead)
	cmp428 = 97 <= v197
	if cmp428 {
		goto land_lhs_true430
	} else {
		goto if_end434
	}

land_lhs_true430:
	v198 = *libc.As[int32](lookahead)
	cmp431 = v198 <= 122
	if cmp431 {
		goto if_then433
	} else {
		goto if_end434
	}

if_then433:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end434:
	v199 = *libc.As[byte](result)
	loadedv435 = (v199 & 1) != 0
	*libc.As[bool](retval) = loadedv435
	goto _return

sw_bb436:
	*libc.As[byte](result) = 1
	v200 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol437 = libc.Ptr(&libc.As[TSLexer](v200).F1)
	*libc.As[int16](result_symbol437) = 9
	v201 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end438 = libc.Ptr(&libc.As[TSLexer](v201).F3)
	v202 = *libc.As[unsafe.Pointer](mark_end438)
	v203 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v202)(v203)
	v204 = *libc.As[int32](lookahead)
	cmp439 = v204 == 78
	if cmp439 {
		goto if_then441
	} else {
		goto if_end442
	}

if_then441:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end442:
	v205 = *libc.As[int32](lookahead)
	cmp443 = v205 == 45
	if cmp443 {
		goto if_then469
	} else {
		goto lor_lhs_false445
	}

lor_lhs_false445:
	v206 = *libc.As[int32](lookahead)
	cmp446 = v206 == 46
	if cmp446 {
		goto if_then469
	} else {
		goto lor_lhs_false448
	}

lor_lhs_false448:
	v207 = *libc.As[int32](lookahead)
	cmp449 = 48 <= v207
	if cmp449 {
		goto land_lhs_true451
	} else {
		goto lor_lhs_false454
	}

land_lhs_true451:
	v208 = *libc.As[int32](lookahead)
	cmp452 = v208 <= 57
	if cmp452 {
		goto if_then469
	} else {
		goto lor_lhs_false454
	}

lor_lhs_false454:
	v209 = *libc.As[int32](lookahead)
	cmp455 = 65 <= v209
	if cmp455 {
		goto land_lhs_true457
	} else {
		goto lor_lhs_false460
	}

land_lhs_true457:
	v210 = *libc.As[int32](lookahead)
	cmp458 = v210 <= 90
	if cmp458 {
		goto if_then469
	} else {
		goto lor_lhs_false460
	}

lor_lhs_false460:
	v211 = *libc.As[int32](lookahead)
	cmp461 = v211 == 95
	if cmp461 {
		goto if_then469
	} else {
		goto lor_lhs_false463
	}

lor_lhs_false463:
	v212 = *libc.As[int32](lookahead)
	cmp464 = 97 <= v212
	if cmp464 {
		goto land_lhs_true466
	} else {
		goto if_end470
	}

land_lhs_true466:
	v213 = *libc.As[int32](lookahead)
	cmp467 = v213 <= 122
	if cmp467 {
		goto if_then469
	} else {
		goto if_end470
	}

if_then469:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end470:
	v214 = *libc.As[byte](result)
	loadedv471 = (v214 & 1) != 0
	*libc.As[bool](retval) = loadedv471
	goto _return

sw_bb472:
	*libc.As[byte](result) = 1
	v215 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol473 = libc.Ptr(&libc.As[TSLexer](v215).F1)
	*libc.As[int16](result_symbol473) = 9
	v216 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end474 = libc.Ptr(&libc.As[TSLexer](v216).F3)
	v217 = *libc.As[unsafe.Pointer](mark_end474)
	v218 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v217)(v218)
	v219 = *libc.As[int32](lookahead)
	cmp475 = v219 == 79
	if cmp475 {
		goto if_then477
	} else {
		goto if_end478
	}

if_then477:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end478:
	v220 = *libc.As[int32](lookahead)
	cmp479 = v220 == 45
	if cmp479 {
		goto if_then505
	} else {
		goto lor_lhs_false481
	}

lor_lhs_false481:
	v221 = *libc.As[int32](lookahead)
	cmp482 = v221 == 46
	if cmp482 {
		goto if_then505
	} else {
		goto lor_lhs_false484
	}

lor_lhs_false484:
	v222 = *libc.As[int32](lookahead)
	cmp485 = 48 <= v222
	if cmp485 {
		goto land_lhs_true487
	} else {
		goto lor_lhs_false490
	}

land_lhs_true487:
	v223 = *libc.As[int32](lookahead)
	cmp488 = v223 <= 57
	if cmp488 {
		goto if_then505
	} else {
		goto lor_lhs_false490
	}

lor_lhs_false490:
	v224 = *libc.As[int32](lookahead)
	cmp491 = 65 <= v224
	if cmp491 {
		goto land_lhs_true493
	} else {
		goto lor_lhs_false496
	}

land_lhs_true493:
	v225 = *libc.As[int32](lookahead)
	cmp494 = v225 <= 90
	if cmp494 {
		goto if_then505
	} else {
		goto lor_lhs_false496
	}

lor_lhs_false496:
	v226 = *libc.As[int32](lookahead)
	cmp497 = v226 == 95
	if cmp497 {
		goto if_then505
	} else {
		goto lor_lhs_false499
	}

lor_lhs_false499:
	v227 = *libc.As[int32](lookahead)
	cmp500 = 97 <= v227
	if cmp500 {
		goto land_lhs_true502
	} else {
		goto if_end506
	}

land_lhs_true502:
	v228 = *libc.As[int32](lookahead)
	cmp503 = v228 <= 122
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end506:
	v229 = *libc.As[byte](result)
	loadedv507 = (v229 & 1) != 0
	*libc.As[bool](retval) = loadedv507
	goto _return

sw_bb508:
	*libc.As[byte](result) = 1
	v230 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol509 = libc.Ptr(&libc.As[TSLexer](v230).F1)
	*libc.As[int16](result_symbol509) = 9
	v231 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end510 = libc.Ptr(&libc.As[TSLexer](v231).F3)
	v232 = *libc.As[unsafe.Pointer](mark_end510)
	v233 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v232)(v233)
	v234 = *libc.As[int32](lookahead)
	cmp511 = v234 == 82
	if cmp511 {
		goto if_then513
	} else {
		goto if_end514
	}

if_then513:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end514:
	v235 = *libc.As[int32](lookahead)
	cmp515 = v235 == 45
	if cmp515 {
		goto if_then541
	} else {
		goto lor_lhs_false517
	}

lor_lhs_false517:
	v236 = *libc.As[int32](lookahead)
	cmp518 = v236 == 46
	if cmp518 {
		goto if_then541
	} else {
		goto lor_lhs_false520
	}

lor_lhs_false520:
	v237 = *libc.As[int32](lookahead)
	cmp521 = 48 <= v237
	if cmp521 {
		goto land_lhs_true523
	} else {
		goto lor_lhs_false526
	}

land_lhs_true523:
	v238 = *libc.As[int32](lookahead)
	cmp524 = v238 <= 57
	if cmp524 {
		goto if_then541
	} else {
		goto lor_lhs_false526
	}

lor_lhs_false526:
	v239 = *libc.As[int32](lookahead)
	cmp527 = 65 <= v239
	if cmp527 {
		goto land_lhs_true529
	} else {
		goto lor_lhs_false532
	}

land_lhs_true529:
	v240 = *libc.As[int32](lookahead)
	cmp530 = v240 <= 90
	if cmp530 {
		goto if_then541
	} else {
		goto lor_lhs_false532
	}

lor_lhs_false532:
	v241 = *libc.As[int32](lookahead)
	cmp533 = v241 == 95
	if cmp533 {
		goto if_then541
	} else {
		goto lor_lhs_false535
	}

lor_lhs_false535:
	v242 = *libc.As[int32](lookahead)
	cmp536 = 97 <= v242
	if cmp536 {
		goto land_lhs_true538
	} else {
		goto if_end542
	}

land_lhs_true538:
	v243 = *libc.As[int32](lookahead)
	cmp539 = v243 <= 122
	if cmp539 {
		goto if_then541
	} else {
		goto if_end542
	}

if_then541:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end542:
	v244 = *libc.As[byte](result)
	loadedv543 = (v244 & 1) != 0
	*libc.As[bool](retval) = loadedv543
	goto _return

sw_bb544:
	*libc.As[byte](result) = 1
	v245 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol545 = libc.Ptr(&libc.As[TSLexer](v245).F1)
	*libc.As[int16](result_symbol545) = 9
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end546 = libc.Ptr(&libc.As[TSLexer](v246).F3)
	v247 = *libc.As[unsafe.Pointer](mark_end546)
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v247)(v248)
	v249 = *libc.As[int32](lookahead)
	cmp547 = v249 == 82
	if cmp547 {
		goto if_then549
	} else {
		goto if_end550
	}

if_then549:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end550:
	v250 = *libc.As[int32](lookahead)
	cmp551 = v250 == 45
	if cmp551 {
		goto if_then577
	} else {
		goto lor_lhs_false553
	}

lor_lhs_false553:
	v251 = *libc.As[int32](lookahead)
	cmp554 = v251 == 46
	if cmp554 {
		goto if_then577
	} else {
		goto lor_lhs_false556
	}

lor_lhs_false556:
	v252 = *libc.As[int32](lookahead)
	cmp557 = 48 <= v252
	if cmp557 {
		goto land_lhs_true559
	} else {
		goto lor_lhs_false562
	}

land_lhs_true559:
	v253 = *libc.As[int32](lookahead)
	cmp560 = v253 <= 57
	if cmp560 {
		goto if_then577
	} else {
		goto lor_lhs_false562
	}

lor_lhs_false562:
	v254 = *libc.As[int32](lookahead)
	cmp563 = 65 <= v254
	if cmp563 {
		goto land_lhs_true565
	} else {
		goto lor_lhs_false568
	}

land_lhs_true565:
	v255 = *libc.As[int32](lookahead)
	cmp566 = v255 <= 90
	if cmp566 {
		goto if_then577
	} else {
		goto lor_lhs_false568
	}

lor_lhs_false568:
	v256 = *libc.As[int32](lookahead)
	cmp569 = v256 == 95
	if cmp569 {
		goto if_then577
	} else {
		goto lor_lhs_false571
	}

lor_lhs_false571:
	v257 = *libc.As[int32](lookahead)
	cmp572 = 97 <= v257
	if cmp572 {
		goto land_lhs_true574
	} else {
		goto if_end578
	}

land_lhs_true574:
	v258 = *libc.As[int32](lookahead)
	cmp575 = v258 <= 122
	if cmp575 {
		goto if_then577
	} else {
		goto if_end578
	}

if_then577:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end578:
	v259 = *libc.As[byte](result)
	loadedv579 = (v259 & 1) != 0
	*libc.As[bool](retval) = loadedv579
	goto _return

sw_bb580:
	*libc.As[byte](result) = 1
	v260 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol581 = libc.Ptr(&libc.As[TSLexer](v260).F1)
	*libc.As[int16](result_symbol581) = 9
	v261 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end582 = libc.Ptr(&libc.As[TSLexer](v261).F3)
	v262 = *libc.As[unsafe.Pointer](mark_end582)
	v263 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v262)(v263)
	v264 = *libc.As[int32](lookahead)
	cmp583 = v264 == 82
	if cmp583 {
		goto if_then585
	} else {
		goto if_end586
	}

if_then585:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end586:
	v265 = *libc.As[int32](lookahead)
	cmp587 = v265 == 45
	if cmp587 {
		goto if_then613
	} else {
		goto lor_lhs_false589
	}

lor_lhs_false589:
	v266 = *libc.As[int32](lookahead)
	cmp590 = v266 == 46
	if cmp590 {
		goto if_then613
	} else {
		goto lor_lhs_false592
	}

lor_lhs_false592:
	v267 = *libc.As[int32](lookahead)
	cmp593 = 48 <= v267
	if cmp593 {
		goto land_lhs_true595
	} else {
		goto lor_lhs_false598
	}

land_lhs_true595:
	v268 = *libc.As[int32](lookahead)
	cmp596 = v268 <= 57
	if cmp596 {
		goto if_then613
	} else {
		goto lor_lhs_false598
	}

lor_lhs_false598:
	v269 = *libc.As[int32](lookahead)
	cmp599 = 65 <= v269
	if cmp599 {
		goto land_lhs_true601
	} else {
		goto lor_lhs_false604
	}

land_lhs_true601:
	v270 = *libc.As[int32](lookahead)
	cmp602 = v270 <= 90
	if cmp602 {
		goto if_then613
	} else {
		goto lor_lhs_false604
	}

lor_lhs_false604:
	v271 = *libc.As[int32](lookahead)
	cmp605 = v271 == 95
	if cmp605 {
		goto if_then613
	} else {
		goto lor_lhs_false607
	}

lor_lhs_false607:
	v272 = *libc.As[int32](lookahead)
	cmp608 = 97 <= v272
	if cmp608 {
		goto land_lhs_true610
	} else {
		goto if_end614
	}

land_lhs_true610:
	v273 = *libc.As[int32](lookahead)
	cmp611 = v273 <= 122
	if cmp611 {
		goto if_then613
	} else {
		goto if_end614
	}

if_then613:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end614:
	v274 = *libc.As[byte](result)
	loadedv615 = (v274 & 1) != 0
	*libc.As[bool](retval) = loadedv615
	goto _return

sw_bb616:
	*libc.As[byte](result) = 1
	v275 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol617 = libc.Ptr(&libc.As[TSLexer](v275).F1)
	*libc.As[int16](result_symbol617) = 9
	v276 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end618 = libc.Ptr(&libc.As[TSLexer](v276).F3)
	v277 = *libc.As[unsafe.Pointer](mark_end618)
	v278 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v277)(v278)
	v279 = *libc.As[int32](lookahead)
	cmp619 = v279 == 83
	if cmp619 {
		goto if_then621
	} else {
		goto if_end622
	}

if_then621:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end622:
	v280 = *libc.As[int32](lookahead)
	cmp623 = v280 == 45
	if cmp623 {
		goto if_then649
	} else {
		goto lor_lhs_false625
	}

lor_lhs_false625:
	v281 = *libc.As[int32](lookahead)
	cmp626 = v281 == 46
	if cmp626 {
		goto if_then649
	} else {
		goto lor_lhs_false628
	}

lor_lhs_false628:
	v282 = *libc.As[int32](lookahead)
	cmp629 = 48 <= v282
	if cmp629 {
		goto land_lhs_true631
	} else {
		goto lor_lhs_false634
	}

land_lhs_true631:
	v283 = *libc.As[int32](lookahead)
	cmp632 = v283 <= 57
	if cmp632 {
		goto if_then649
	} else {
		goto lor_lhs_false634
	}

lor_lhs_false634:
	v284 = *libc.As[int32](lookahead)
	cmp635 = 65 <= v284
	if cmp635 {
		goto land_lhs_true637
	} else {
		goto lor_lhs_false640
	}

land_lhs_true637:
	v285 = *libc.As[int32](lookahead)
	cmp638 = v285 <= 90
	if cmp638 {
		goto if_then649
	} else {
		goto lor_lhs_false640
	}

lor_lhs_false640:
	v286 = *libc.As[int32](lookahead)
	cmp641 = v286 == 95
	if cmp641 {
		goto if_then649
	} else {
		goto lor_lhs_false643
	}

lor_lhs_false643:
	v287 = *libc.As[int32](lookahead)
	cmp644 = 97 <= v287
	if cmp644 {
		goto land_lhs_true646
	} else {
		goto if_end650
	}

land_lhs_true646:
	v288 = *libc.As[int32](lookahead)
	cmp647 = v288 <= 122
	if cmp647 {
		goto if_then649
	} else {
		goto if_end650
	}

if_then649:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end650:
	v289 = *libc.As[byte](result)
	loadedv651 = (v289 & 1) != 0
	*libc.As[bool](retval) = loadedv651
	goto _return

sw_bb652:
	*libc.As[byte](result) = 1
	v290 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol653 = libc.Ptr(&libc.As[TSLexer](v290).F1)
	*libc.As[int16](result_symbol653) = 9
	v291 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end654 = libc.Ptr(&libc.As[TSLexer](v291).F3)
	v292 = *libc.As[unsafe.Pointer](mark_end654)
	v293 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v292)(v293)
	v294 = *libc.As[int32](lookahead)
	cmp655 = v294 == 83
	if cmp655 {
		goto if_then657
	} else {
		goto if_end658
	}

if_then657:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end658:
	v295 = *libc.As[int32](lookahead)
	cmp659 = v295 == 45
	if cmp659 {
		goto if_then685
	} else {
		goto lor_lhs_false661
	}

lor_lhs_false661:
	v296 = *libc.As[int32](lookahead)
	cmp662 = v296 == 46
	if cmp662 {
		goto if_then685
	} else {
		goto lor_lhs_false664
	}

lor_lhs_false664:
	v297 = *libc.As[int32](lookahead)
	cmp665 = 48 <= v297
	if cmp665 {
		goto land_lhs_true667
	} else {
		goto lor_lhs_false670
	}

land_lhs_true667:
	v298 = *libc.As[int32](lookahead)
	cmp668 = v298 <= 57
	if cmp668 {
		goto if_then685
	} else {
		goto lor_lhs_false670
	}

lor_lhs_false670:
	v299 = *libc.As[int32](lookahead)
	cmp671 = 65 <= v299
	if cmp671 {
		goto land_lhs_true673
	} else {
		goto lor_lhs_false676
	}

land_lhs_true673:
	v300 = *libc.As[int32](lookahead)
	cmp674 = v300 <= 90
	if cmp674 {
		goto if_then685
	} else {
		goto lor_lhs_false676
	}

lor_lhs_false676:
	v301 = *libc.As[int32](lookahead)
	cmp677 = v301 == 95
	if cmp677 {
		goto if_then685
	} else {
		goto lor_lhs_false679
	}

lor_lhs_false679:
	v302 = *libc.As[int32](lookahead)
	cmp680 = 97 <= v302
	if cmp680 {
		goto land_lhs_true682
	} else {
		goto if_end686
	}

land_lhs_true682:
	v303 = *libc.As[int32](lookahead)
	cmp683 = v303 <= 122
	if cmp683 {
		goto if_then685
	} else {
		goto if_end686
	}

if_then685:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end686:
	v304 = *libc.As[byte](result)
	loadedv687 = (v304 & 1) != 0
	*libc.As[bool](retval) = loadedv687
	goto _return

sw_bb688:
	*libc.As[byte](result) = 1
	v305 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol689 = libc.Ptr(&libc.As[TSLexer](v305).F1)
	*libc.As[int16](result_symbol689) = 9
	v306 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end690 = libc.Ptr(&libc.As[TSLexer](v306).F3)
	v307 = *libc.As[unsafe.Pointer](mark_end690)
	v308 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v307)(v308)
	v309 = *libc.As[int32](lookahead)
	cmp691 = v309 == 45
	if cmp691 {
		goto if_then717
	} else {
		goto lor_lhs_false693
	}

lor_lhs_false693:
	v310 = *libc.As[int32](lookahead)
	cmp694 = v310 == 46
	if cmp694 {
		goto if_then717
	} else {
		goto lor_lhs_false696
	}

lor_lhs_false696:
	v311 = *libc.As[int32](lookahead)
	cmp697 = 48 <= v311
	if cmp697 {
		goto land_lhs_true699
	} else {
		goto lor_lhs_false702
	}

land_lhs_true699:
	v312 = *libc.As[int32](lookahead)
	cmp700 = v312 <= 57
	if cmp700 {
		goto if_then717
	} else {
		goto lor_lhs_false702
	}

lor_lhs_false702:
	v313 = *libc.As[int32](lookahead)
	cmp703 = 65 <= v313
	if cmp703 {
		goto land_lhs_true705
	} else {
		goto lor_lhs_false708
	}

land_lhs_true705:
	v314 = *libc.As[int32](lookahead)
	cmp706 = v314 <= 90
	if cmp706 {
		goto if_then717
	} else {
		goto lor_lhs_false708
	}

lor_lhs_false708:
	v315 = *libc.As[int32](lookahead)
	cmp709 = v315 == 95
	if cmp709 {
		goto if_then717
	} else {
		goto lor_lhs_false711
	}

lor_lhs_false711:
	v316 = *libc.As[int32](lookahead)
	cmp712 = 97 <= v316
	if cmp712 {
		goto land_lhs_true714
	} else {
		goto if_end718
	}

land_lhs_true714:
	v317 = *libc.As[int32](lookahead)
	cmp715 = v317 <= 122
	if cmp715 {
		goto if_then717
	} else {
		goto if_end718
	}

if_then717:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end718:
	v318 = *libc.As[byte](result)
	loadedv719 = (v318 & 1) != 0
	*libc.As[bool](retval) = loadedv719
	goto _return

sw_bb720:
	*libc.As[byte](result) = 1
	v319 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol721 = libc.Ptr(&libc.As[TSLexer](v319).F1)
	*libc.As[int16](result_symbol721) = 10
	v320 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end722 = libc.Ptr(&libc.As[TSLexer](v320).F3)
	v321 = *libc.As[unsafe.Pointer](mark_end722)
	v322 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v321)(v322)
	v323 = *libc.As[byte](result)
	loadedv723 = (v323 & 1) != 0
	*libc.As[bool](retval) = loadedv723
	goto _return

sw_bb724:
	*libc.As[byte](result) = 1
	v324 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol725 = libc.Ptr(&libc.As[TSLexer](v324).F1)
	*libc.As[int16](result_symbol725) = 11
	v325 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end726 = libc.Ptr(&libc.As[TSLexer](v325).F3)
	v326 = *libc.As[unsafe.Pointer](mark_end726)
	v327 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v326)(v327)
	v328 = *libc.As[int32](lookahead)
	cmp727 = 48 <= v328
	if cmp727 {
		goto land_lhs_true729
	} else {
		goto if_end733
	}

land_lhs_true729:
	v329 = *libc.As[int32](lookahead)
	cmp730 = v329 <= 57
	if cmp730 {
		goto if_then732
	} else {
		goto if_end733
	}

if_then732:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end733:
	v330 = *libc.As[byte](result)
	loadedv734 = (v330 & 1) != 0
	*libc.As[bool](retval) = loadedv734
	goto _return

sw_bb735:
	*libc.As[byte](result) = 1
	v331 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol736 = libc.Ptr(&libc.As[TSLexer](v331).F1)
	*libc.As[int16](result_symbol736) = 12
	v332 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end737 = libc.Ptr(&libc.As[TSLexer](v332).F3)
	v333 = *libc.As[unsafe.Pointer](mark_end737)
	v334 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v333)(v334)
	v335 = *libc.As[byte](result)
	loadedv738 = (v335 & 1) != 0
	*libc.As[bool](retval) = loadedv738
	goto _return

sw_bb739:
	*libc.As[byte](result) = 1
	v336 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol740 = libc.Ptr(&libc.As[TSLexer](v336).F1)
	*libc.As[int16](result_symbol740) = 13
	v337 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end741 = libc.Ptr(&libc.As[TSLexer](v337).F3)
	v338 = *libc.As[unsafe.Pointer](mark_end741)
	v339 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v338)(v339)
	v340 = *libc.As[byte](result)
	loadedv742 = (v340 & 1) != 0
	*libc.As[bool](retval) = loadedv742
	goto _return

sw_bb743:
	*libc.As[byte](result) = 1
	v341 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol744 = libc.Ptr(&libc.As[TSLexer](v341).F1)
	*libc.As[int16](result_symbol744) = 14
	v342 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end745 = libc.Ptr(&libc.As[TSLexer](v342).F3)
	v343 = *libc.As[unsafe.Pointer](mark_end745)
	v344 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v343)(v344)
	v345 = *libc.As[byte](result)
	loadedv746 = (v345 & 1) != 0
	*libc.As[bool](retval) = loadedv746
	goto _return

sw_bb747:
	*libc.As[byte](result) = 1
	v346 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol748 = libc.Ptr(&libc.As[TSLexer](v346).F1)
	*libc.As[int16](result_symbol748) = 14
	v347 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end749 = libc.Ptr(&libc.As[TSLexer](v347).F3)
	v348 = *libc.As[unsafe.Pointer](mark_end749)
	v349 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v348)(v349)
	v350 = *libc.As[int32](lookahead)
	cmp750 = v350 == 10
	if cmp750 {
		goto if_then752
	} else {
		goto if_end753
	}

if_then752:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end753:
	v351 = *libc.As[byte](result)
	loadedv754 = (v351 & 1) != 0
	*libc.As[bool](retval) = loadedv754
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v352 = *libc.As[bool](retval)
	return v352
}
