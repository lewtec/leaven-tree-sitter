package grammar_test

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
type Scanner struct {
	F0 int32
	F1 int32
	F2 byte
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

var tree_sitter_test_language struct {
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
var ts_small_parse_table [325]int16 = [325]int16{10, 9, 1, 2, 13, 1, 6, 15, 1, 9, 17, 1, 12, 19, 1, 14, 7, 1, 23, 9, 1, 22, 26, 1, 20, 3, 2, 21, 25, 11, 3, 3, 4, 5, 9, 9, 1, 2, 13, 1, 6, 15, 1, 9, 21, 1, 12, 23, 1, 14, 7, 1, 23, 9, 1, 22, 4, 2, 21, 25, 11, 3, 3, 4, 5, 9, 25, 1, 2, 31, 1, 6, 34, 1, 9, 37, 1, 12, 40, 1, 14, 7, 1, 23, 9, 1, 22, 4, 2, 21, 25, 28, 3, 3, 4, 5, 1, 42, 8, 14, 2, 3, 4, 5, 6, 9, 12, 1, 44, 8, 14, 2, 3, 4, 5, 6, 9, 12, 1, 46, 8, 14, 2, 3, 4, 5, 6, 9, 12, 1, 48, 8, 14, 2, 3, 4, 5, 6, 9, 12, 1, 50, 8, 14, 2, 3, 4, 5, 6, 9, 12, 1, 52, 8, 14, 2, 3, 4, 5, 6, 9, 12, 5, 7, 1, 13, 54, 1, 0, 56, 1, 12, 19, 1, 19, 12, 2, 17, 24, 5, 58, 1, 0, 60, 1, 12, 63, 1, 13, 19, 1, 19, 12, 2, 17, 24, 3, 14, 1, 18, 66, 2, 13, 0, 68, 2, 1, 12, 3, 15, 1, 18, 68, 2, 1, 12, 70, 2, 13, 0, 3, 15, 1, 18, 72, 2, 13, 0, 74, 2, 1, 12, 3, 79, 1, 15, 17, 1, 18, 77, 2, 1, 12, 3, 72, 1, 15, 17, 1, 18, 81, 2, 1, 12, 1, 84, 4, 13, 0, 1, 12, 2, 16, 1, 18, 77, 2, 1, 12, 1, 84, 3, 15, 1, 12, 1, 86, 2, 1, 12, 1, 88, 2, 1, 12, 1, 90, 1, 8, 1, 92, 1, 11, 1, 94, 1, 7, 1, 96, 1, 14, 1, 98, 1, 8, 1, 100, 1, 1, 1, 102, 1, 0, 1, 104, 1, 12, 1, 106, 1, 7, 1, 108, 1, 10}
var ts_small_parse_table_map [31]int32 = [31]int32{0, 34, 65, 96, 107, 118, 129, 140, 151, 162, 179, 196, 208, 220, 232, 243, 254, 261, 269, 275, 280, 285, 289, 293, 297, 301, 305, 309, 313, 317, 321}
var ts_symbol_names [29]unsafe.Pointer = [29]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_14), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27)}
var ts_field_names [4]unsafe.Pointer = [4]unsafe.Pointer{nil, libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_30)}
var ts_field_map_slices [9]TSMapSlice = [9]TSMapSlice{TSMapSlice{}, TSMapSlice{}, TSMapSlice{}, TSMapSlice{0, 1}, TSMapSlice{1, 1}, TSMapSlice{2, 1}, TSMapSlice{}, TSMapSlice{3, 1}, TSMapSlice{4, 1}}
var ts_field_map_entries [5]TSFieldMapEntry = [5]TSFieldMapEntry{TSFieldMapEntry{1, 0, 0}, TSFieldMapEntry{2, 0, 1}, TSFieldMapEntry{3, 0, 1}, TSFieldMapEntry{2, 2, 0}, TSFieldMapEntry{3, 2, 0}}
var ts_symbol_metadata [29]TSSymbolMetadata = [29]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}}
var ts_symbol_map [29]int16 = [29]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 12, 13, 13, 13, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28}
var ts_non_terminal_alias_map [6]int16 = [6]int16{18, 3, 18, 26, 28, 0}
var ts_alias_sequences [9][5]int16 = [9][5]int16{[5]int16{}, [5]int16{0, 26, 0, 0, 0}, [5]int16{0, 27, 0, 0, 0}, [5]int16{}, [5]int16{}, [5]int16{}, [5]int16{0, 26, 0, 28, 0}, [5]int16{}, [5]int16{}}
var ts_lex_modes [33]TSLexMode = [33]TSLexMode{TSLexMode{0, 1}, TSLexMode{0, 2}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{77, 2}, TSLexMode{77, 2}, TSLexMode{77, 2}, TSLexMode{77, 4}, TSLexMode{77, 4}, TSLexMode{77, 2}, TSLexMode{77, 0}, TSLexMode{77, 4}, TSLexMode{77, 0}, TSLexMode{77, 0}, TSLexMode{}, TSLexMode{6, 0}, TSLexMode{}, TSLexMode{0, 3}, TSLexMode{}, TSLexMode{77, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{76, 0}}
var ts_external_scanner_states [5][3]byte = [5][3]byte{[3]byte{}, [3]byte{1, 1, 1}, [3]byte{1, 0, 0}, [3]byte{0, 1, 0}, [3]byte{0, 0, 1}}
var ts_external_scanner_symbol_map [3]int16 = [3]int16{13, 14, 15}
var ts_primary_state_ids [33]int16 = [33]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 15, 18, 19, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
var ts_parse_table struct {
	F0 struct {
		F0 [16]int16
		F1 [10]int16
	}
	F1 [26]int16
} = struct {
	F0 struct {
		F0 [16]int16
		F1 [10]int16
	}
	F1 [26]int16
}{struct {
	F0 [16]int16
	F1 [10]int16
}{[16]int16{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, [10]int16{}}, [26]int16{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 7, 0, 0, 29, 11, 0, 19, 0, 0, 0, 0, 11, 0}}
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
	F26 TSParseActionEntry
	F27 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F28 struct {
		F0 anon_2
		F1 [6]byte
	}
	F29 TSParseActionEntry
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F34 struct {
		F0 anon_2
		F1 [6]byte
	}
	F35 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F41 TSParseActionEntry
	F42 struct {
		F0 anon_2
		F1 [6]byte
	}
	F43 TSParseActionEntry
	F44 struct {
		F0 anon_2
		F1 [6]byte
	}
	F45 TSParseActionEntry
	F46 struct {
		F0 anon_2
		F1 [6]byte
	}
	F47 TSParseActionEntry
	F48 struct {
		F0 anon_2
		F1 [6]byte
	}
	F49 TSParseActionEntry
	F50 struct {
		F0 anon_2
		F1 [6]byte
	}
	F51 TSParseActionEntry
	F52 struct {
		F0 anon_2
		F1 [6]byte
	}
	F53 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
	F64 TSParseActionEntry
	F65 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F66 struct {
		F0 anon_2
		F1 [6]byte
	}
	F67 TSParseActionEntry
	F68 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F71 TSParseActionEntry
	F72 struct {
		F0 anon_2
		F1 [6]byte
	}
	F73 TSParseActionEntry
	F74 struct {
		F0 anon_2
		F1 [6]byte
	}
	F75 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F85 TSParseActionEntry
	F86 struct {
		F0 anon_2
		F1 [6]byte
	}
	F87 TSParseActionEntry
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
	F95 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F101 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F102 struct {
		F0 anon_2
		F1 [6]byte
	}
	F103 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F104 struct {
		F0 anon_2
		F1 [6]byte
	}
	F105 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F106 struct {
		F0 anon_2
		F1 [6]byte
	}
	F107 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F108 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F26 TSParseActionEntry
	F27 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F28 struct {
		F0 anon_2
		F1 [6]byte
	}
	F29 TSParseActionEntry
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F34 struct {
		F0 anon_2
		F1 [6]byte
	}
	F35 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F41 TSParseActionEntry
	F42 struct {
		F0 anon_2
		F1 [6]byte
	}
	F43 TSParseActionEntry
	F44 struct {
		F0 anon_2
		F1 [6]byte
	}
	F45 TSParseActionEntry
	F46 struct {
		F0 anon_2
		F1 [6]byte
	}
	F47 TSParseActionEntry
	F48 struct {
		F0 anon_2
		F1 [6]byte
	}
	F49 TSParseActionEntry
	F50 struct {
		F0 anon_2
		F1 [6]byte
	}
	F51 TSParseActionEntry
	F52 struct {
		F0 anon_2
		F1 [6]byte
	}
	F53 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
	F64 TSParseActionEntry
	F65 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F66 struct {
		F0 anon_2
		F1 [6]byte
	}
	F67 TSParseActionEntry
	F68 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F71 TSParseActionEntry
	F72 struct {
		F0 anon_2
		F1 [6]byte
	}
	F73 TSParseActionEntry
	F74 struct {
		F0 anon_2
		F1 [6]byte
	}
	F75 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F85 TSParseActionEntry
	F86 struct {
		F0 anon_2
		F1 [6]byte
	}
	F87 TSParseActionEntry
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
	F95 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F101 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F102 struct {
		F0 anon_2
		F1 [6]byte
	}
	F103 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F104 struct {
		F0 anon_2
		F1 [6]byte
	}
	F105 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F106 struct {
		F0 anon_2
		F1 [6]byte
	}
	F107 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F108 struct {
		F0 anon_2
		F1 [6]byte
	}
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 16, 0, 0}}}, struct {
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
}{0, 0, 5, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 31, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 4, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 20, 0, 0}}}, struct {
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
}{0, 0, 5, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 6, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 31, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 25, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 4, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 25, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 21, 0, 3}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 21, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 21, 0, 5}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 22, 0, 7}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 21, 0, 4}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 23, 0, 8}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 16, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 12, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 28, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 17, 0, 1}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 17, 0, 6}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 18, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 18, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 18, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 13, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 18, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 20, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 18, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 19, 0, 2}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 19, 0, 2}}}, struct {
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
}{0, 0, 23, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 30, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 27, 0, 0}, [2]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [6]byte = [6]byte{95, 108, 105, 110, 101, 0}
var _str_4 [5]byte = [5]byte{58, 99, 115, 116, 0}
var _str_5 [7]byte = [7]byte{58, 101, 114, 114, 111, 114, 0}
var _str_6 [11]byte = [11]byte{58, 102, 97, 105, 108, 45, 102, 97, 115, 116, 0}
var _str_7 [6]byte = [6]byte{58, 115, 107, 105, 112, 0}
var _str_8 [10]byte = [10]byte{58, 108, 97, 110, 103, 117, 97, 103, 101, 0}
var _str_9 [2]byte = [2]byte{40, 0}
var _str_10 [2]byte = [2]byte{41, 0}
var _str_11 [10]byte = [10]byte{58, 112, 108, 97, 116, 102, 111, 114, 109, 0}
var _str_12 [10]byte = [10]byte{112, 97, 114, 97, 109, 101, 116, 101, 114, 0}
var _str_13 [5]byte = [5]byte{95, 101, 111, 108, 0}
var _str_14 [10]byte = [10]byte{115, 101, 112, 97, 114, 97, 116, 111, 114, 0}
var _str_15 [5]byte = [5]byte{102, 105, 108, 101, 0}
var _str_16 [5]byte = [5]byte{116, 101, 115, 116, 0}
var _str_17 [6]byte = [6]byte{95, 98, 111, 100, 121, 0}
var _str_18 [7]byte = [7]byte{104, 101, 97, 100, 101, 114, 0}
var _str_19 [11]byte = [11]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 0}
var _str_20 [10]byte = [10]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 0}
var _str_21 [10]byte = [10]byte{95, 108, 97, 110, 103, 117, 97, 103, 101, 0}
var _str_22 [10]byte = [10]byte{95, 112, 108, 97, 116, 102, 111, 114, 109, 0}
var _str_23 [13]byte = [13]byte{102, 105, 108, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_24 [19]byte = [19]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_25 [6]byte = [6]byte{105, 110, 112, 117, 116, 0}
var _str_26 [5]byte = [5]byte{110, 97, 109, 101, 0}
var _str_27 [7]byte = [7]byte{111, 117, 116, 112, 117, 116, 0}
var _str_28 [4]byte = [4]byte{99, 115, 116, 0}
var _str_29 [9]byte = [9]byte{108, 97, 110, 103, 117, 97, 103, 101, 0}
var _str_30 [9]byte = [9]byte{112, 108, 97, 116, 102, 111, 114, 109, 0}
var ts_lex_map [30]int16 = [30]int16{10, 132, 13, 133, 40, 85, 41, 86, 58, 12, 97, 108, 100, 119, 102, 120, 105, 113, 108, 103, 109, 88, 110, 96, 111, 118, 115, 114, 119, 105}
var ts_lex_map_31 [20]int16 = [20]int16{97, 41, 100, 62, 102, 59, 105, 47, 108, 31, 109, 4, 110, 19, 111, 54, 115, 52, 119, 33}

func init() {
	tree_sitter_test_language = struct {
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
	}{14, 26, 3, 16, 3, 33, 2, 9, 3, 5, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), libc.Ptr(&ts_field_names), libc.Ptr(&ts_field_map_slices), libc.Ptr(&ts_field_map_entries), libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{libc.Ptr(&ts_external_scanner_states), libc.Ptr(&ts_external_scanner_symbol_map), libc.FuncCode(tree_sitter_test_external_scanner_create), libc.FuncCode(tree_sitter_test_external_scanner_destroy), libc.FuncCode(tree_sitter_test_external_scanner_scan), libc.FuncCode(tree_sitter_test_external_scanner_serialize), libc.FuncCode(tree_sitter_test_external_scanner_deserialize)}, libc.Ptr(&ts_primary_state_ids), nil, nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{}, [5]byte{}}
}
func tree_sitter_test_external_scanner_create() unsafe.Pointer {
	var length, suffix unsafe.Pointer
	var initialized unsafe.Pointer
	var call, v0, v1, v2, v3 unsafe.Pointer
	var scanner unsafe.Pointer
	_, _, _, _, _, _, _, _, _ = scanner, call, v0, length, v1, suffix, v2, initialized, v3

	scanner = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	call = libc.Ptr(libc.Malloc[byte](int64(12)))
	*libc.As[unsafe.Pointer](scanner) = call
	v0 = *libc.As[unsafe.Pointer](scanner)
	length = libc.Ptr(&libc.As[Scanner](v0).F0)
	*libc.As[int32](length) = 0
	v1 = *libc.As[unsafe.Pointer](scanner)
	suffix = libc.Ptr(&libc.As[Scanner](v1).F1)
	*libc.As[int32](suffix) = 0
	v2 = *libc.As[unsafe.Pointer](scanner)
	initialized = libc.Ptr(&libc.As[Scanner](v2).F2)
	*libc.As[byte](initialized) = 0
	v3 = *libc.As[unsafe.Pointer](scanner)
	return v3
}
func malloc(a0 int64) unsafe.Pointer {
	panic("unsatisfied: malloc")
}
func tree_sitter_test_external_scanner_serialize(payload unsafe.Pointer, buffer unsafe.Pointer) int32 {
	var cmp bool
	var v2, add, v6, v9, v11, add3, v12 int32
	var retval, length, suffix, length1, length2 unsafe.Pointer
	var conv int64
	var v7 byte
	var add_ptr unsafe.Pointer
	var v0, v1, v3, v4, v5, v8, v10 unsafe.Pointer
	var payload_addr, buffer_addr, scanner unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, payload_addr, buffer_addr, scanner, v0, v1, length, v2, add, cmp, v3, v4, add_ptr, v5, suffix, v6, v7, v8, length1, v9, conv, v10, length2, v11, add3, v12

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
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
	length = libc.Ptr(&libc.As[Scanner](v1).F0)
	v2 = *libc.As[int32](length)
	add = v2 + 1
	cmp = uint32(add) > 1024
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[int32](retval) = 0
	goto _return

if_end:
	v3 = *libc.As[unsafe.Pointer](buffer_addr)
	libc.Memset(libc.As[byte](v3), 1, int64(1))
	v4 = *libc.As[unsafe.Pointer](buffer_addr)
	add_ptr = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v4), int(int64(1))*1))
	v5 = *libc.As[unsafe.Pointer](scanner)
	suffix = libc.Ptr(&libc.As[Scanner](v5).F1)
	v6 = *libc.As[int32](suffix)
	v7 = byte(v6)
	v8 = *libc.As[unsafe.Pointer](scanner)
	length1 = libc.Ptr(&libc.As[Scanner](v8).F0)
	v9 = *libc.As[int32](length1)
	conv = int64(uint64(uint32(v9)))
	libc.Memset(libc.As[byte](add_ptr), v7, conv)
	v10 = *libc.As[unsafe.Pointer](scanner)
	length2 = libc.Ptr(&libc.As[Scanner](v10).F0)
	v11 = *libc.As[int32](length2)
	add3 = v11 + 1
	*libc.As[int32](retval) = add3
	goto _return

_return:
	v12 = *libc.As[int32](retval)
	return v12
}
func tree_sitter_test_external_scanner_deserialize(payload unsafe.Pointer, buffer unsafe.Pointer, length int32) {
	var cmp, tobool bool
	var v0, v2, sub, conv int32
	var length_addr, length1, suffix unsafe.Pointer
	var v5, v8, storedv byte
	var arrayidx, arrayidx2, initialized unsafe.Pointer
	var v1, v3, v4, v6, v7, v9 unsafe.Pointer
	var payload_addr, buffer_addr, scanner unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = payload_addr, buffer_addr, length_addr, scanner, v0, cmp, v1, v2, sub, v3, length1, v4, arrayidx, v5, conv, v6, suffix, v7, arrayidx2, v8, tobool, v9, initialized, storedv

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
	v0 = *libc.As[int32](length_addr)
	cmp = v0 == 0
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	goto _return

if_end:
	v1 = *libc.As[unsafe.Pointer](payload_addr)
	*libc.As[unsafe.Pointer](scanner) = v1
	v2 = *libc.As[int32](length_addr)
	sub = v2 - 1
	v3 = *libc.As[unsafe.Pointer](scanner)
	length1 = libc.Ptr(&libc.As[Scanner](v3).F0)
	*libc.As[int32](length1) = sub
	v4 = *libc.As[unsafe.Pointer](buffer_addr)
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v4), int(int64(1))*1))
	v5 = *libc.As[byte](arrayidx)
	conv = int32(int8(v5))
	v6 = *libc.As[unsafe.Pointer](scanner)
	suffix = libc.Ptr(&libc.As[Scanner](v6).F1)
	*libc.As[int32](suffix) = conv
	v7 = *libc.As[unsafe.Pointer](buffer_addr)
	arrayidx2 = v7
	v8 = *libc.As[byte](arrayidx2)
	tobool = v8 != 0
	v9 = *libc.As[unsafe.Pointer](scanner)
	initialized = libc.Ptr(&libc.As[Scanner](v9).F2)
	if tobool {
		storedv = 1
	} else {
		storedv = 0
	}
	*libc.As[byte](initialized) = storedv
	goto _return

_return:
}
func tree_sitter_test_external_scanner_scan(payload unsafe.Pointer, lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var loadedv, call, loadedv2, call4, loadedv7, call9, v13 bool
	var retval unsafe.Pointer
	var v2, v6, v10 byte
	var arrayidx, arrayidx1, arrayidx6 unsafe.Pointer
	var v0, v1, v3, v4, v5, v7, v8, v9, v11, v12 unsafe.Pointer
	var payload_addr, lexer_addr, valid_symbols_addr, scanner unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, payload_addr, lexer_addr, valid_symbols_addr, scanner, v0, v1, arrayidx, v2, loadedv, v3, v4, call, v5, arrayidx1, v6, loadedv2, v7, v8, call4, v9, arrayidx6, v10, loadedv7, v11, v12, call9, v13

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
	v1 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx = v1
	v2 = *libc.As[byte](arrayidx)
	loadedv = (v2 & 1) != 0
	if loadedv {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	v4 = *libc.As[unsafe.Pointer](scanner)
	call = scan(v3, v4, 61, 0)
	*libc.As[bool](retval) = call
	goto _return

if_end:
	v5 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx1 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v5), int(int64(1))*1))
	v6 = *libc.As[byte](arrayidx1)
	loadedv2 = (v6 & 1) != 0
	if loadedv2 {
		goto if_then3
	} else {
		goto if_end5
	}

if_then3:
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	v8 = *libc.As[unsafe.Pointer](scanner)
	call4 = scan(v7, v8, 61, 1)
	*libc.As[bool](retval) = call4
	goto _return

if_end5:
	v9 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx6 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v9), int(int64(2))*1))
	v10 = *libc.As[byte](arrayidx6)
	loadedv7 = (v10 & 1) != 0
	if loadedv7 {
		goto if_then8
	} else {
		goto if_end10
	}

if_then8:
	v11 = *libc.As[unsafe.Pointer](lexer_addr)
	v12 = *libc.As[unsafe.Pointer](scanner)
	call9 = scan(v11, v12, 45, 2)
	*libc.As[bool](retval) = call9
	goto _return

if_end10:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v13 = *libc.As[bool](retval)
	return v13
}
func scan(lexer unsafe.Pointer, scanner unsafe.Pointer, chr byte, symbol int32) bool {
	var call, cmp, v6, cmp2, call6, cmp9, cmp12, v17, lnot, v18, cmp16, loadedv, cmp23, cmp30, loadedv34, cmp39, cmp42, cmp46, cmp51, v48 bool
	var retval unsafe.Pointer
	var conv55 int16
	var result_symbol unsafe.Pointer
	var v4, conv, v8, inc, v9, v14, v16, v19, v23, v26, v28, v30, inc28, v31, v34, v37, v38, v39, v41, v44, v46 int32
	var symbol_addr, length, lookahead, lookahead8, lookahead11, lookahead19, suffix, suffix21, lookahead22, length36, length38, lookahead45, lookahead50 unsafe.Pointer
	var v5, v21, v33 byte
	var chr_addr, initialized, initialized33 unsafe.Pointer
	var v0, v1, v2, v3, v7, v10, v11, v12, v13, v15, v20, v22, v24, v25, v27, v29, v32, v35, v36, v40, v42, v43, v45, v47 unsafe.Pointer
	var lexer_addr, scanner_addr, eof, eof5 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, scanner_addr, chr_addr, symbol_addr, length, v0, eof, v1, v2, call, v3, lookahead, v4, v5, conv, cmp, v6, v7, v8, inc, v9, cmp2, v10, eof5, v11, v12, call6, v13, lookahead8, v14, cmp9, v15, lookahead11, v16, cmp12, v17, lnot, v18, v19, cmp16, v20, initialized, v21, loadedv, v22, lookahead19, v23, v24, suffix, v25, suffix21, v26, v27, lookahead22, v28, cmp23, v29, v30, inc28, v31, cmp30, v32, initialized33, v33, loadedv34, v34, v35, length36, v36, length38, v37, v38, cmp39, v39, cmp42, v40, lookahead45, v41, cmp46, v42, v43, lookahead50, v44, cmp51, v45, v46, conv55, v47, result_symbol, v48

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
	scanner_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	chr_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	symbol_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	length = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](scanner_addr) = scanner
	*libc.As[byte](chr_addr) = chr
	*libc.As[int32](symbol_addr) = symbol
	*libc.As[int32](length) = 0
	goto for_cond

for_cond:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	eof = libc.Ptr(&libc.As[TSLexer](v0).F6)
	v1 = *libc.As[unsafe.Pointer](eof)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	call = libc.FuncFromCode[func(unsafe.Pointer) bool](v1)(v2)
	if call {
		v6 = false
		goto land_end
	} else {
		goto land_rhs
	}

land_rhs:
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v3).F0)
	v4 = *libc.As[int32](lookahead)
	v5 = *libc.As[byte](chr_addr)
	conv = int32(int8(v5))
	cmp = v4 == conv
	v6 = cmp
	goto land_end

land_end:
	if v6 {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v7)
	goto for_inc

for_inc:
	v8 = *libc.As[int32](length)
	inc = v8 + 1
	*libc.As[int32](length) = inc
	goto for_cond

for_end:
	v9 = *libc.As[int32](length)
	cmp2 = uint32(v9) < 3
	if cmp2 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	*libc.As[int32](length) = 0
	goto for_cond4

for_cond4:
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	eof5 = libc.Ptr(&libc.As[TSLexer](v10).F6)
	v11 = *libc.As[unsafe.Pointer](eof5)
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	call6 = libc.FuncFromCode[func(unsafe.Pointer) bool](v11)(v12)
	if call6 {
		v18 = false
		goto land_end14
	} else {
		goto land_rhs7
	}

land_rhs7:
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead8 = libc.Ptr(&libc.As[TSLexer](v13).F0)
	v14 = *libc.As[int32](lookahead8)
	cmp9 = v14 == 13
	if cmp9 {
		v17 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead11 = libc.Ptr(&libc.As[TSLexer](v15).F0)
	v16 = *libc.As[int32](lookahead11)
	cmp12 = v16 == 10
	v17 = cmp12
	goto lor_end

lor_end:
	lnot = v17 != true
	v18 = lnot
	goto land_end14

land_end14:
	if v18 {
		goto for_body15
	} else {
		goto for_end29
	}

for_body15:
	v19 = *libc.As[int32](symbol_addr)
	cmp16 = v19 == 0
	if cmp16 {
		goto land_lhs_true
	} else {
		goto if_end20
	}

land_lhs_true:
	v20 = *libc.As[unsafe.Pointer](scanner_addr)
	initialized = libc.Ptr(&libc.As[Scanner](v20).F2)
	v21 = *libc.As[byte](initialized)
	loadedv = (v21 & 1) != 0
	if loadedv {
		goto if_end20
	} else {
		goto if_then18
	}

if_then18:
	v22 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead19 = libc.Ptr(&libc.As[TSLexer](v22).F0)
	v23 = *libc.As[int32](lookahead19)
	v24 = *libc.As[unsafe.Pointer](scanner_addr)
	suffix = libc.Ptr(&libc.As[Scanner](v24).F1)
	*libc.As[int32](suffix) = v23
	goto if_end20

if_end20:
	v25 = *libc.As[unsafe.Pointer](scanner_addr)
	suffix21 = libc.Ptr(&libc.As[Scanner](v25).F1)
	v26 = *libc.As[int32](suffix21)
	v27 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead22 = libc.Ptr(&libc.As[TSLexer](v27).F0)
	v28 = *libc.As[int32](lookahead22)
	cmp23 = v26 != v28
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*libc.As[bool](retval) = false
	goto _return

if_end26:
	v29 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v29)
	goto for_inc27

for_inc27:
	v30 = *libc.As[int32](length)
	inc28 = v30 + 1
	*libc.As[int32](length) = inc28
	goto for_cond4

for_end29:
	v31 = *libc.As[int32](symbol_addr)
	cmp30 = v31 == 0
	if cmp30 {
		goto land_lhs_true32
	} else {
		goto if_end37
	}

land_lhs_true32:
	v32 = *libc.As[unsafe.Pointer](scanner_addr)
	initialized33 = libc.Ptr(&libc.As[Scanner](v32).F2)
	v33 = *libc.As[byte](initialized33)
	loadedv34 = (v33 & 1) != 0
	if loadedv34 {
		goto if_end37
	} else {
		goto if_then35
	}

if_then35:
	v34 = *libc.As[int32](length)
	v35 = *libc.As[unsafe.Pointer](scanner_addr)
	length36 = libc.Ptr(&libc.As[Scanner](v35).F0)
	*libc.As[int32](length36) = v34
	goto if_end37

if_end37:
	v36 = *libc.As[unsafe.Pointer](scanner_addr)
	length38 = libc.Ptr(&libc.As[Scanner](v36).F0)
	v37 = *libc.As[int32](length38)
	v38 = *libc.As[int32](length)
	cmp39 = v37 != v38
	if cmp39 {
		goto if_then41
	} else {
		goto if_end44
	}

if_then41:
	v39 = *libc.As[int32](symbol_addr)
	cmp42 = v39 == 2
	*libc.As[bool](retval) = cmp42
	goto _return

if_end44:
	v40 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead45 = libc.Ptr(&libc.As[TSLexer](v40).F0)
	v41 = *libc.As[int32](lookahead45)
	cmp46 = v41 == 13
	if cmp46 {
		goto if_then48
	} else {
		goto if_end49
	}

if_then48:
	v42 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v42)
	goto if_end49

if_end49:
	v43 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead50 = libc.Ptr(&libc.As[TSLexer](v43).F0)
	v44 = *libc.As[int32](lookahead50)
	cmp51 = v44 == 10
	if cmp51 {
		goto if_then53
	} else {
		goto if_end54
	}

if_then53:
	v45 = *libc.As[unsafe.Pointer](lexer_addr)
	advance(v45)
	goto if_end54

if_end54:
	v46 = *libc.As[int32](symbol_addr)
	conv55 = int16(v46)
	v47 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v47).F1)
	*libc.As[int16](result_symbol) = conv55
	*libc.As[bool](retval) = true
	goto _return

_return:
	v48 = *libc.As[bool](retval)
	return v48
}
func tree_sitter_test_external_scanner_destroy(payload unsafe.Pointer) {
	var v0, v1 unsafe.Pointer
	var payload_addr, scanner unsafe.Pointer
	_, _, _, _ = payload_addr, scanner, v0, v1

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
	*libc.As[unsafe.Pointer](payload_addr) = payload
	v0 = *libc.As[unsafe.Pointer](payload_addr)
	*libc.As[unsafe.Pointer](scanner) = v0
	v1 = *libc.As[unsafe.Pointer](scanner)
	libc.Free(libc.As[byte](v1))
}
func tree_sitter_test() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_test_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp19, cmp21, cmp24, cmp27, cmp30, cmp33, cmp36, loadedv40, cmp42, loadedv46, cmp48, loadedv52, cmp54, loadedv58, cmp60, loadedv64, cmp66, loadedv70, cmp75, cmp81, loadedv91, cmp93, loadedv97, cmp99, loadedv103, cmp105, loadedv109, cmp111, loadedv115, cmp117, loadedv121, cmp123, cmp127, cmp131, cmp135, cmp139, cmp143, loadedv147, cmp149, loadedv153, cmp155, loadedv159, cmp161, loadedv165, cmp167, loadedv171, cmp173, loadedv177, cmp179, loadedv183, cmp185, loadedv189, cmp191, loadedv195, cmp197, loadedv201, cmp203, loadedv207, cmp209, loadedv213, cmp215, loadedv219, cmp221, loadedv225, cmp227, loadedv231, cmp233, loadedv237, cmp239, loadedv243, cmp245, loadedv249, cmp251, loadedv255, cmp257, loadedv261, cmp263, loadedv267, cmp269, loadedv273, cmp275, loadedv279, cmp281, loadedv285, cmp287, loadedv291, cmp293, loadedv297, cmp299, loadedv303, cmp305, loadedv309, cmp311, loadedv315, cmp317, loadedv321, cmp323, loadedv327, cmp329, loadedv333, cmp335, loadedv339, cmp341, loadedv345, cmp347, loadedv351, cmp353, loadedv357, cmp359, loadedv363, cmp365, loadedv369, cmp371, loadedv375, cmp377, loadedv381, cmp383, loadedv387, cmp389, loadedv393, cmp395, loadedv399, cmp401, loadedv405, cmp407, loadedv411, cmp413, loadedv417, cmp419, loadedv423, cmp425, loadedv429, cmp431, loadedv435, cmp437, loadedv441, cmp443, loadedv447, cmp449, loadedv453, cmp455, loadedv459, cmp461, loadedv465, cmp467, loadedv471, cmp473, loadedv477, cmp479, loadedv483, cmp485, loadedv489, cmp491, loadedv495, cmp497, loadedv501, cmp503, loadedv507, cmp509, loadedv513, cmp515, loadedv519, cmp521, loadedv525, cmp527, cmp530, cmp533, cmp536, cmp539, cmp542, cmp545, cmp548, cmp551, loadedv555, loadedv557, cmp560, cmp564, cmp568, loadedv572, loadedv574, cmp578, cmp581, cmp584, loadedv588, loadedv592, loadedv596, loadedv600, loadedv604, loadedv608, loadedv612, loadedv616, loadedv620, cmp624, cmp628, cmp631, cmp634, cmp637, cmp640, cmp643, cmp646, cmp649, cmp652, loadedv656, cmp660, cmp664, cmp667, cmp670, cmp673, cmp676, cmp679, cmp682, cmp685, cmp688, loadedv692, cmp696, cmp700, cmp703, cmp706, cmp709, cmp712, cmp715, cmp718, cmp721, cmp724, loadedv728, cmp732, cmp736, cmp739, cmp742, cmp745, cmp748, cmp751, cmp754, cmp757, cmp760, loadedv764, cmp768, cmp772, cmp775, cmp778, cmp781, cmp784, cmp787, cmp790, cmp793, cmp796, loadedv800, cmp804, cmp808, cmp811, cmp814, cmp817, cmp820, cmp823, cmp826, cmp829, cmp832, loadedv836, cmp840, cmp844, cmp847, cmp850, cmp853, cmp856, cmp859, cmp862, cmp865, cmp868, loadedv872, cmp876, cmp880, cmp883, cmp886, cmp889, cmp892, cmp895, cmp898, cmp901, cmp904, loadedv908, cmp912, cmp916, cmp919, cmp922, cmp925, cmp928, cmp931, cmp934, cmp937, cmp940, loadedv944, cmp948, cmp952, cmp955, cmp958, cmp961, cmp964, cmp967, cmp970, cmp973, cmp976, loadedv980, cmp984, cmp988, cmp991, cmp994, cmp997, cmp1000, cmp1003, cmp1006, cmp1009, cmp1012, loadedv1016, cmp1020, cmp1024, cmp1027, cmp1030, cmp1033, cmp1036, cmp1039, cmp1042, cmp1045, cmp1048, loadedv1052, cmp1056, cmp1060, cmp1063, cmp1066, cmp1069, cmp1072, cmp1075, cmp1078, cmp1081, cmp1084, loadedv1088, cmp1092, cmp1096, cmp1099, cmp1102, cmp1105, cmp1108, cmp1111, cmp1114, cmp1117, cmp1120, loadedv1124, cmp1128, cmp1132, cmp1135, cmp1138, cmp1141, cmp1144, cmp1147, cmp1150, cmp1153, cmp1156, loadedv1160, cmp1164, cmp1168, cmp1171, cmp1174, cmp1177, cmp1180, cmp1183, cmp1186, cmp1189, cmp1192, loadedv1196, cmp1200, cmp1204, cmp1207, cmp1210, cmp1213, cmp1216, cmp1219, cmp1222, cmp1225, cmp1228, loadedv1232, cmp1236, cmp1240, cmp1243, cmp1246, cmp1249, cmp1252, cmp1255, cmp1258, cmp1261, cmp1264, loadedv1268, cmp1272, cmp1276, cmp1279, cmp1282, cmp1285, cmp1288, cmp1291, cmp1294, cmp1297, cmp1300, loadedv1304, cmp1308, cmp1312, cmp1315, cmp1318, cmp1321, cmp1324, cmp1327, cmp1330, cmp1333, cmp1336, loadedv1340, cmp1344, cmp1348, cmp1351, cmp1354, cmp1357, cmp1360, cmp1363, cmp1366, cmp1369, cmp1372, loadedv1376, cmp1380, cmp1384, cmp1387, cmp1390, cmp1393, cmp1396, cmp1399, cmp1402, cmp1405, cmp1408, loadedv1412, cmp1416, cmp1420, cmp1423, cmp1426, cmp1429, cmp1432, cmp1435, cmp1438, cmp1441, cmp1444, loadedv1448, cmp1452, cmp1456, cmp1459, cmp1462, cmp1465, cmp1468, cmp1471, cmp1474, cmp1477, cmp1480, loadedv1484, cmp1488, cmp1492, cmp1495, cmp1498, cmp1501, cmp1504, cmp1507, cmp1510, cmp1513, cmp1516, loadedv1520, cmp1524, cmp1528, cmp1531, cmp1534, cmp1537, cmp1540, cmp1543, cmp1546, cmp1549, cmp1552, loadedv1556, cmp1560, cmp1564, cmp1567, cmp1570, cmp1573, cmp1576, cmp1579, cmp1582, cmp1585, cmp1588, loadedv1592, cmp1596, cmp1600, cmp1603, cmp1606, cmp1609, cmp1612, cmp1615, cmp1618, cmp1621, cmp1624, loadedv1628, cmp1632, cmp1636, cmp1639, cmp1642, cmp1645, cmp1648, cmp1651, cmp1654, cmp1657, cmp1660, loadedv1664, cmp1668, cmp1672, cmp1675, cmp1678, cmp1681, cmp1684, cmp1687, cmp1690, cmp1693, cmp1696, loadedv1700, cmp1704, cmp1708, cmp1711, cmp1714, cmp1717, cmp1720, cmp1723, cmp1726, cmp1729, cmp1732, loadedv1736, cmp1740, cmp1744, cmp1747, cmp1750, cmp1753, cmp1756, cmp1759, cmp1762, cmp1765, cmp1768, loadedv1772, cmp1776, cmp1780, cmp1783, cmp1786, cmp1789, cmp1792, cmp1795, cmp1798, cmp1801, cmp1804, loadedv1808, cmp1812, cmp1816, cmp1819, cmp1822, cmp1825, cmp1828, cmp1831, cmp1834, cmp1837, cmp1840, loadedv1844, cmp1848, cmp1852, cmp1855, cmp1858, cmp1861, cmp1864, cmp1867, cmp1870, cmp1873, cmp1876, loadedv1880, cmp1884, cmp1888, cmp1891, cmp1894, cmp1897, cmp1900, cmp1903, cmp1906, cmp1909, cmp1912, loadedv1916, cmp1920, cmp1924, cmp1927, cmp1930, cmp1933, cmp1936, cmp1939, cmp1942, cmp1945, cmp1948, loadedv1952, cmp1956, cmp1960, cmp1963, cmp1966, cmp1969, cmp1972, cmp1975, cmp1978, cmp1981, cmp1984, loadedv1988, cmp1992, cmp1996, cmp1999, cmp2002, cmp2005, cmp2008, cmp2011, cmp2014, cmp2017, cmp2020, loadedv2024, cmp2028, cmp2032, cmp2035, cmp2038, cmp2041, cmp2044, cmp2047, cmp2050, cmp2053, cmp2056, loadedv2060, cmp2064, cmp2068, cmp2071, cmp2074, cmp2077, cmp2080, cmp2083, cmp2086, cmp2089, cmp2092, loadedv2096, cmp2100, cmp2104, cmp2107, cmp2110, cmp2113, cmp2116, cmp2119, cmp2122, cmp2125, cmp2128, loadedv2132, cmp2136, cmp2139, cmp2142, cmp2145, cmp2148, cmp2151, cmp2154, cmp2157, cmp2160, loadedv2164, loadedv2168, loadedv2172, cmp2176, loadedv2180, v917 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v40, v43 int16
	var state_addr, arrayidx, arrayidx11, arrayidx79, arrayidx86, result_symbol, result_symbol576, result_symbol590, result_symbol594, result_symbol598, result_symbol602, result_symbol606, result_symbol610, result_symbol614, result_symbol618, result_symbol622, result_symbol658, result_symbol694, result_symbol730, result_symbol766, result_symbol802, result_symbol838, result_symbol874, result_symbol910, result_symbol946, result_symbol982, result_symbol1018, result_symbol1054, result_symbol1090, result_symbol1126, result_symbol1162, result_symbol1198, result_symbol1234, result_symbol1270, result_symbol1306, result_symbol1342, result_symbol1378, result_symbol1414, result_symbol1450, result_symbol1486, result_symbol1522, result_symbol1558, result_symbol1594, result_symbol1630, result_symbol1666, result_symbol1702, result_symbol1738, result_symbol1774, result_symbol1810, result_symbol1846, result_symbol1882, result_symbol1918, result_symbol1954, result_symbol1990, result_symbol2026, result_symbol2062, result_symbol2098, result_symbol2134, result_symbol2166, result_symbol2170, result_symbol2174 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v24, v25, v26, v28, v30, v32, v34, v36, v38, v39, conv80, v41, v42, add84, v44, add89, v46, v48, v50, v52, v54, v56, v57, v58, v59, v60, v61, v63, v65, v67, v69, v71, v73, v75, v77, v79, v81, v83, v85, v87, v89, v91, v93, v95, v97, v99, v101, v103, v105, v107, v109, v111, v113, v115, v117, v119, v121, v123, v125, v127, v129, v131, v133, v135, v137, v139, v141, v143, v145, v147, v149, v151, v153, v155, v157, v159, v161, v163, v165, v167, v169, v171, v173, v175, v177, v179, v181, v183, v185, v187, v189, v190, v191, v192, v193, v194, v195, v196, v197, v200, v201, v202, v213, v214, v215, v261, v262, v263, v264, v265, v266, v267, v268, v269, v270, v276, v277, v278, v279, v280, v281, v282, v283, v284, v285, v291, v292, v293, v294, v295, v296, v297, v298, v299, v300, v306, v307, v308, v309, v310, v311, v312, v313, v314, v315, v321, v322, v323, v324, v325, v326, v327, v328, v329, v330, v336, v337, v338, v339, v340, v341, v342, v343, v344, v345, v351, v352, v353, v354, v355, v356, v357, v358, v359, v360, v366, v367, v368, v369, v370, v371, v372, v373, v374, v375, v381, v382, v383, v384, v385, v386, v387, v388, v389, v390, v396, v397, v398, v399, v400, v401, v402, v403, v404, v405, v411, v412, v413, v414, v415, v416, v417, v418, v419, v420, v426, v427, v428, v429, v430, v431, v432, v433, v434, v435, v441, v442, v443, v444, v445, v446, v447, v448, v449, v450, v456, v457, v458, v459, v460, v461, v462, v463, v464, v465, v471, v472, v473, v474, v475, v476, v477, v478, v479, v480, v486, v487, v488, v489, v490, v491, v492, v493, v494, v495, v501, v502, v503, v504, v505, v506, v507, v508, v509, v510, v516, v517, v518, v519, v520, v521, v522, v523, v524, v525, v531, v532, v533, v534, v535, v536, v537, v538, v539, v540, v546, v547, v548, v549, v550, v551, v552, v553, v554, v555, v561, v562, v563, v564, v565, v566, v567, v568, v569, v570, v576, v577, v578, v579, v580, v581, v582, v583, v584, v585, v591, v592, v593, v594, v595, v596, v597, v598, v599, v600, v606, v607, v608, v609, v610, v611, v612, v613, v614, v615, v621, v622, v623, v624, v625, v626, v627, v628, v629, v630, v636, v637, v638, v639, v640, v641, v642, v643, v644, v645, v651, v652, v653, v654, v655, v656, v657, v658, v659, v660, v666, v667, v668, v669, v670, v671, v672, v673, v674, v675, v681, v682, v683, v684, v685, v686, v687, v688, v689, v690, v696, v697, v698, v699, v700, v701, v702, v703, v704, v705, v711, v712, v713, v714, v715, v716, v717, v718, v719, v720, v726, v727, v728, v729, v730, v731, v732, v733, v734, v735, v741, v742, v743, v744, v745, v746, v747, v748, v749, v750, v756, v757, v758, v759, v760, v761, v762, v763, v764, v765, v771, v772, v773, v774, v775, v776, v777, v778, v779, v780, v786, v787, v788, v789, v790, v791, v792, v793, v794, v795, v801, v802, v803, v804, v805, v806, v807, v808, v809, v810, v816, v817, v818, v819, v820, v821, v822, v823, v824, v825, v831, v832, v833, v834, v835, v836, v837, v838, v839, v840, v846, v847, v848, v849, v850, v851, v852, v853, v854, v855, v861, v862, v863, v864, v865, v866, v867, v868, v869, v870, v876, v877, v878, v879, v880, v881, v882, v883, v884, v885, v891, v892, v893, v894, v895, v896, v897, v898, v899, v915 int32
	var lookahead, i, i72, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv74, idxprom78, idxprom85 int64
	var v3, storedv, v10, v27, v29, v31, v33, v35, v37, v45, v47, v49, v51, v53, v55, v62, v64, v66, v68, v70, v72, v74, v76, v78, v80, v82, v84, v86, v88, v90, v92, v94, v96, v98, v100, v102, v104, v106, v108, v110, v112, v114, v116, v118, v120, v122, v124, v126, v128, v130, v132, v134, v136, v138, v140, v142, v144, v146, v148, v150, v152, v154, v156, v158, v160, v162, v164, v166, v168, v170, v172, v174, v176, v178, v180, v182, v184, v186, v188, v198, v199, v203, v208, v216, v221, v226, v231, v236, v241, v246, v251, v256, v271, v286, v301, v316, v331, v346, v361, v376, v391, v406, v421, v436, v451, v466, v481, v496, v511, v526, v541, v556, v571, v586, v601, v616, v631, v646, v661, v676, v691, v706, v721, v736, v751, v766, v781, v796, v811, v826, v841, v856, v871, v886, v900, v905, v910, v916 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v204, v205, v206, v207, v209, v210, v211, v212, v217, v218, v219, v220, v222, v223, v224, v225, v227, v228, v229, v230, v232, v233, v234, v235, v237, v238, v239, v240, v242, v243, v244, v245, v247, v248, v249, v250, v252, v253, v254, v255, v257, v258, v259, v260, v272, v273, v274, v275, v287, v288, v289, v290, v302, v303, v304, v305, v317, v318, v319, v320, v332, v333, v334, v335, v347, v348, v349, v350, v362, v363, v364, v365, v377, v378, v379, v380, v392, v393, v394, v395, v407, v408, v409, v410, v422, v423, v424, v425, v437, v438, v439, v440, v452, v453, v454, v455, v467, v468, v469, v470, v482, v483, v484, v485, v497, v498, v499, v500, v512, v513, v514, v515, v527, v528, v529, v530, v542, v543, v544, v545, v557, v558, v559, v560, v572, v573, v574, v575, v587, v588, v589, v590, v602, v603, v604, v605, v617, v618, v619, v620, v632, v633, v634, v635, v647, v648, v649, v650, v662, v663, v664, v665, v677, v678, v679, v680, v692, v693, v694, v695, v707, v708, v709, v710, v722, v723, v724, v725, v737, v738, v739, v740, v752, v753, v754, v755, v767, v768, v769, v770, v782, v783, v784, v785, v797, v798, v799, v800, v812, v813, v814, v815, v827, v828, v829, v830, v842, v843, v844, v845, v857, v858, v859, v860, v872, v873, v874, v875, v887, v888, v889, v890, v901, v902, v903, v904, v906, v907, v908, v909, v911, v912, v913, v914 unsafe.Pointer
	var lexer_addr, local_advance, eof2, mark_end, mark_end577, mark_end591, mark_end595, mark_end599, mark_end603, mark_end607, mark_end611, mark_end615, mark_end619, mark_end623, mark_end659, mark_end695, mark_end731, mark_end767, mark_end803, mark_end839, mark_end875, mark_end911, mark_end947, mark_end983, mark_end1019, mark_end1055, mark_end1091, mark_end1127, mark_end1163, mark_end1199, mark_end1235, mark_end1271, mark_end1307, mark_end1343, mark_end1379, mark_end1415, mark_end1451, mark_end1487, mark_end1523, mark_end1559, mark_end1595, mark_end1631, mark_end1667, mark_end1703, mark_end1739, mark_end1775, mark_end1811, mark_end1847, mark_end1883, mark_end1919, mark_end1955, mark_end1991, mark_end2027, mark_end2063, mark_end2099, mark_end2135, mark_end2167, mark_end2171, mark_end2175 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i72, v0, local_advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp19, v21, cmp21, v22, cmp24, v23, cmp27, v24, cmp30, v25, cmp33, v26, cmp36, v27, loadedv40, v28, cmp42, v29, loadedv46, v30, cmp48, v31, loadedv52, v32, cmp54, v33, loadedv58, v34, cmp60, v35, loadedv64, v36, cmp66, v37, loadedv70, v38, conv74, cmp75, v39, idxprom78, arrayidx79, v40, conv80, v41, cmp81, v42, add84, idxprom85, arrayidx86, v43, v44, add89, v45, loadedv91, v46, cmp93, v47, loadedv97, v48, cmp99, v49, loadedv103, v50, cmp105, v51, loadedv109, v52, cmp111, v53, loadedv115, v54, cmp117, v55, loadedv121, v56, cmp123, v57, cmp127, v58, cmp131, v59, cmp135, v60, cmp139, v61, cmp143, v62, loadedv147, v63, cmp149, v64, loadedv153, v65, cmp155, v66, loadedv159, v67, cmp161, v68, loadedv165, v69, cmp167, v70, loadedv171, v71, cmp173, v72, loadedv177, v73, cmp179, v74, loadedv183, v75, cmp185, v76, loadedv189, v77, cmp191, v78, loadedv195, v79, cmp197, v80, loadedv201, v81, cmp203, v82, loadedv207, v83, cmp209, v84, loadedv213, v85, cmp215, v86, loadedv219, v87, cmp221, v88, loadedv225, v89, cmp227, v90, loadedv231, v91, cmp233, v92, loadedv237, v93, cmp239, v94, loadedv243, v95, cmp245, v96, loadedv249, v97, cmp251, v98, loadedv255, v99, cmp257, v100, loadedv261, v101, cmp263, v102, loadedv267, v103, cmp269, v104, loadedv273, v105, cmp275, v106, loadedv279, v107, cmp281, v108, loadedv285, v109, cmp287, v110, loadedv291, v111, cmp293, v112, loadedv297, v113, cmp299, v114, loadedv303, v115, cmp305, v116, loadedv309, v117, cmp311, v118, loadedv315, v119, cmp317, v120, loadedv321, v121, cmp323, v122, loadedv327, v123, cmp329, v124, loadedv333, v125, cmp335, v126, loadedv339, v127, cmp341, v128, loadedv345, v129, cmp347, v130, loadedv351, v131, cmp353, v132, loadedv357, v133, cmp359, v134, loadedv363, v135, cmp365, v136, loadedv369, v137, cmp371, v138, loadedv375, v139, cmp377, v140, loadedv381, v141, cmp383, v142, loadedv387, v143, cmp389, v144, loadedv393, v145, cmp395, v146, loadedv399, v147, cmp401, v148, loadedv405, v149, cmp407, v150, loadedv411, v151, cmp413, v152, loadedv417, v153, cmp419, v154, loadedv423, v155, cmp425, v156, loadedv429, v157, cmp431, v158, loadedv435, v159, cmp437, v160, loadedv441, v161, cmp443, v162, loadedv447, v163, cmp449, v164, loadedv453, v165, cmp455, v166, loadedv459, v167, cmp461, v168, loadedv465, v169, cmp467, v170, loadedv471, v171, cmp473, v172, loadedv477, v173, cmp479, v174, loadedv483, v175, cmp485, v176, loadedv489, v177, cmp491, v178, loadedv495, v179, cmp497, v180, loadedv501, v181, cmp503, v182, loadedv507, v183, cmp509, v184, loadedv513, v185, cmp515, v186, loadedv519, v187, cmp521, v188, loadedv525, v189, cmp527, v190, cmp530, v191, cmp533, v192, cmp536, v193, cmp539, v194, cmp542, v195, cmp545, v196, cmp548, v197, cmp551, v198, loadedv555, v199, loadedv557, v200, cmp560, v201, cmp564, v202, cmp568, v203, loadedv572, v204, result_symbol, v205, mark_end, v206, v207, v208, loadedv574, v209, result_symbol576, v210, mark_end577, v211, v212, v213, cmp578, v214, cmp581, v215, cmp584, v216, loadedv588, v217, result_symbol590, v218, mark_end591, v219, v220, v221, loadedv592, v222, result_symbol594, v223, mark_end595, v224, v225, v226, loadedv596, v227, result_symbol598, v228, mark_end599, v229, v230, v231, loadedv600, v232, result_symbol602, v233, mark_end603, v234, v235, v236, loadedv604, v237, result_symbol606, v238, mark_end607, v239, v240, v241, loadedv608, v242, result_symbol610, v243, mark_end611, v244, v245, v246, loadedv612, v247, result_symbol614, v248, mark_end615, v249, v250, v251, loadedv616, v252, result_symbol618, v253, mark_end619, v254, v255, v256, loadedv620, v257, result_symbol622, v258, mark_end623, v259, v260, v261, cmp624, v262, cmp628, v263, cmp631, v264, cmp634, v265, cmp637, v266, cmp640, v267, cmp643, v268, cmp646, v269, cmp649, v270, cmp652, v271, loadedv656, v272, result_symbol658, v273, mark_end659, v274, v275, v276, cmp660, v277, cmp664, v278, cmp667, v279, cmp670, v280, cmp673, v281, cmp676, v282, cmp679, v283, cmp682, v284, cmp685, v285, cmp688, v286, loadedv692, v287, result_symbol694, v288, mark_end695, v289, v290, v291, cmp696, v292, cmp700, v293, cmp703, v294, cmp706, v295, cmp709, v296, cmp712, v297, cmp715, v298, cmp718, v299, cmp721, v300, cmp724, v301, loadedv728, v302, result_symbol730, v303, mark_end731, v304, v305, v306, cmp732, v307, cmp736, v308, cmp739, v309, cmp742, v310, cmp745, v311, cmp748, v312, cmp751, v313, cmp754, v314, cmp757, v315, cmp760, v316, loadedv764, v317, result_symbol766, v318, mark_end767, v319, v320, v321, cmp768, v322, cmp772, v323, cmp775, v324, cmp778, v325, cmp781, v326, cmp784, v327, cmp787, v328, cmp790, v329, cmp793, v330, cmp796, v331, loadedv800, v332, result_symbol802, v333, mark_end803, v334, v335, v336, cmp804, v337, cmp808, v338, cmp811, v339, cmp814, v340, cmp817, v341, cmp820, v342, cmp823, v343, cmp826, v344, cmp829, v345, cmp832, v346, loadedv836, v347, result_symbol838, v348, mark_end839, v349, v350, v351, cmp840, v352, cmp844, v353, cmp847, v354, cmp850, v355, cmp853, v356, cmp856, v357, cmp859, v358, cmp862, v359, cmp865, v360, cmp868, v361, loadedv872, v362, result_symbol874, v363, mark_end875, v364, v365, v366, cmp876, v367, cmp880, v368, cmp883, v369, cmp886, v370, cmp889, v371, cmp892, v372, cmp895, v373, cmp898, v374, cmp901, v375, cmp904, v376, loadedv908, v377, result_symbol910, v378, mark_end911, v379, v380, v381, cmp912, v382, cmp916, v383, cmp919, v384, cmp922, v385, cmp925, v386, cmp928, v387, cmp931, v388, cmp934, v389, cmp937, v390, cmp940, v391, loadedv944, v392, result_symbol946, v393, mark_end947, v394, v395, v396, cmp948, v397, cmp952, v398, cmp955, v399, cmp958, v400, cmp961, v401, cmp964, v402, cmp967, v403, cmp970, v404, cmp973, v405, cmp976, v406, loadedv980, v407, result_symbol982, v408, mark_end983, v409, v410, v411, cmp984, v412, cmp988, v413, cmp991, v414, cmp994, v415, cmp997, v416, cmp1000, v417, cmp1003, v418, cmp1006, v419, cmp1009, v420, cmp1012, v421, loadedv1016, v422, result_symbol1018, v423, mark_end1019, v424, v425, v426, cmp1020, v427, cmp1024, v428, cmp1027, v429, cmp1030, v430, cmp1033, v431, cmp1036, v432, cmp1039, v433, cmp1042, v434, cmp1045, v435, cmp1048, v436, loadedv1052, v437, result_symbol1054, v438, mark_end1055, v439, v440, v441, cmp1056, v442, cmp1060, v443, cmp1063, v444, cmp1066, v445, cmp1069, v446, cmp1072, v447, cmp1075, v448, cmp1078, v449, cmp1081, v450, cmp1084, v451, loadedv1088, v452, result_symbol1090, v453, mark_end1091, v454, v455, v456, cmp1092, v457, cmp1096, v458, cmp1099, v459, cmp1102, v460, cmp1105, v461, cmp1108, v462, cmp1111, v463, cmp1114, v464, cmp1117, v465, cmp1120, v466, loadedv1124, v467, result_symbol1126, v468, mark_end1127, v469, v470, v471, cmp1128, v472, cmp1132, v473, cmp1135, v474, cmp1138, v475, cmp1141, v476, cmp1144, v477, cmp1147, v478, cmp1150, v479, cmp1153, v480, cmp1156, v481, loadedv1160, v482, result_symbol1162, v483, mark_end1163, v484, v485, v486, cmp1164, v487, cmp1168, v488, cmp1171, v489, cmp1174, v490, cmp1177, v491, cmp1180, v492, cmp1183, v493, cmp1186, v494, cmp1189, v495, cmp1192, v496, loadedv1196, v497, result_symbol1198, v498, mark_end1199, v499, v500, v501, cmp1200, v502, cmp1204, v503, cmp1207, v504, cmp1210, v505, cmp1213, v506, cmp1216, v507, cmp1219, v508, cmp1222, v509, cmp1225, v510, cmp1228, v511, loadedv1232, v512, result_symbol1234, v513, mark_end1235, v514, v515, v516, cmp1236, v517, cmp1240, v518, cmp1243, v519, cmp1246, v520, cmp1249, v521, cmp1252, v522, cmp1255, v523, cmp1258, v524, cmp1261, v525, cmp1264, v526, loadedv1268, v527, result_symbol1270, v528, mark_end1271, v529, v530, v531, cmp1272, v532, cmp1276, v533, cmp1279, v534, cmp1282, v535, cmp1285, v536, cmp1288, v537, cmp1291, v538, cmp1294, v539, cmp1297, v540, cmp1300, v541, loadedv1304, v542, result_symbol1306, v543, mark_end1307, v544, v545, v546, cmp1308, v547, cmp1312, v548, cmp1315, v549, cmp1318, v550, cmp1321, v551, cmp1324, v552, cmp1327, v553, cmp1330, v554, cmp1333, v555, cmp1336, v556, loadedv1340, v557, result_symbol1342, v558, mark_end1343, v559, v560, v561, cmp1344, v562, cmp1348, v563, cmp1351, v564, cmp1354, v565, cmp1357, v566, cmp1360, v567, cmp1363, v568, cmp1366, v569, cmp1369, v570, cmp1372, v571, loadedv1376, v572, result_symbol1378, v573, mark_end1379, v574, v575, v576, cmp1380, v577, cmp1384, v578, cmp1387, v579, cmp1390, v580, cmp1393, v581, cmp1396, v582, cmp1399, v583, cmp1402, v584, cmp1405, v585, cmp1408, v586, loadedv1412, v587, result_symbol1414, v588, mark_end1415, v589, v590, v591, cmp1416, v592, cmp1420, v593, cmp1423, v594, cmp1426, v595, cmp1429, v596, cmp1432, v597, cmp1435, v598, cmp1438, v599, cmp1441, v600, cmp1444, v601, loadedv1448, v602, result_symbol1450, v603, mark_end1451, v604, v605, v606, cmp1452, v607, cmp1456, v608, cmp1459, v609, cmp1462, v610, cmp1465, v611, cmp1468, v612, cmp1471, v613, cmp1474, v614, cmp1477, v615, cmp1480, v616, loadedv1484, v617, result_symbol1486, v618, mark_end1487, v619, v620, v621, cmp1488, v622, cmp1492, v623, cmp1495, v624, cmp1498, v625, cmp1501, v626, cmp1504, v627, cmp1507, v628, cmp1510, v629, cmp1513, v630, cmp1516, v631, loadedv1520, v632, result_symbol1522, v633, mark_end1523, v634, v635, v636, cmp1524, v637, cmp1528, v638, cmp1531, v639, cmp1534, v640, cmp1537, v641, cmp1540, v642, cmp1543, v643, cmp1546, v644, cmp1549, v645, cmp1552, v646, loadedv1556, v647, result_symbol1558, v648, mark_end1559, v649, v650, v651, cmp1560, v652, cmp1564, v653, cmp1567, v654, cmp1570, v655, cmp1573, v656, cmp1576, v657, cmp1579, v658, cmp1582, v659, cmp1585, v660, cmp1588, v661, loadedv1592, v662, result_symbol1594, v663, mark_end1595, v664, v665, v666, cmp1596, v667, cmp1600, v668, cmp1603, v669, cmp1606, v670, cmp1609, v671, cmp1612, v672, cmp1615, v673, cmp1618, v674, cmp1621, v675, cmp1624, v676, loadedv1628, v677, result_symbol1630, v678, mark_end1631, v679, v680, v681, cmp1632, v682, cmp1636, v683, cmp1639, v684, cmp1642, v685, cmp1645, v686, cmp1648, v687, cmp1651, v688, cmp1654, v689, cmp1657, v690, cmp1660, v691, loadedv1664, v692, result_symbol1666, v693, mark_end1667, v694, v695, v696, cmp1668, v697, cmp1672, v698, cmp1675, v699, cmp1678, v700, cmp1681, v701, cmp1684, v702, cmp1687, v703, cmp1690, v704, cmp1693, v705, cmp1696, v706, loadedv1700, v707, result_symbol1702, v708, mark_end1703, v709, v710, v711, cmp1704, v712, cmp1708, v713, cmp1711, v714, cmp1714, v715, cmp1717, v716, cmp1720, v717, cmp1723, v718, cmp1726, v719, cmp1729, v720, cmp1732, v721, loadedv1736, v722, result_symbol1738, v723, mark_end1739, v724, v725, v726, cmp1740, v727, cmp1744, v728, cmp1747, v729, cmp1750, v730, cmp1753, v731, cmp1756, v732, cmp1759, v733, cmp1762, v734, cmp1765, v735, cmp1768, v736, loadedv1772, v737, result_symbol1774, v738, mark_end1775, v739, v740, v741, cmp1776, v742, cmp1780, v743, cmp1783, v744, cmp1786, v745, cmp1789, v746, cmp1792, v747, cmp1795, v748, cmp1798, v749, cmp1801, v750, cmp1804, v751, loadedv1808, v752, result_symbol1810, v753, mark_end1811, v754, v755, v756, cmp1812, v757, cmp1816, v758, cmp1819, v759, cmp1822, v760, cmp1825, v761, cmp1828, v762, cmp1831, v763, cmp1834, v764, cmp1837, v765, cmp1840, v766, loadedv1844, v767, result_symbol1846, v768, mark_end1847, v769, v770, v771, cmp1848, v772, cmp1852, v773, cmp1855, v774, cmp1858, v775, cmp1861, v776, cmp1864, v777, cmp1867, v778, cmp1870, v779, cmp1873, v780, cmp1876, v781, loadedv1880, v782, result_symbol1882, v783, mark_end1883, v784, v785, v786, cmp1884, v787, cmp1888, v788, cmp1891, v789, cmp1894, v790, cmp1897, v791, cmp1900, v792, cmp1903, v793, cmp1906, v794, cmp1909, v795, cmp1912, v796, loadedv1916, v797, result_symbol1918, v798, mark_end1919, v799, v800, v801, cmp1920, v802, cmp1924, v803, cmp1927, v804, cmp1930, v805, cmp1933, v806, cmp1936, v807, cmp1939, v808, cmp1942, v809, cmp1945, v810, cmp1948, v811, loadedv1952, v812, result_symbol1954, v813, mark_end1955, v814, v815, v816, cmp1956, v817, cmp1960, v818, cmp1963, v819, cmp1966, v820, cmp1969, v821, cmp1972, v822, cmp1975, v823, cmp1978, v824, cmp1981, v825, cmp1984, v826, loadedv1988, v827, result_symbol1990, v828, mark_end1991, v829, v830, v831, cmp1992, v832, cmp1996, v833, cmp1999, v834, cmp2002, v835, cmp2005, v836, cmp2008, v837, cmp2011, v838, cmp2014, v839, cmp2017, v840, cmp2020, v841, loadedv2024, v842, result_symbol2026, v843, mark_end2027, v844, v845, v846, cmp2028, v847, cmp2032, v848, cmp2035, v849, cmp2038, v850, cmp2041, v851, cmp2044, v852, cmp2047, v853, cmp2050, v854, cmp2053, v855, cmp2056, v856, loadedv2060, v857, result_symbol2062, v858, mark_end2063, v859, v860, v861, cmp2064, v862, cmp2068, v863, cmp2071, v864, cmp2074, v865, cmp2077, v866, cmp2080, v867, cmp2083, v868, cmp2086, v869, cmp2089, v870, cmp2092, v871, loadedv2096, v872, result_symbol2098, v873, mark_end2099, v874, v875, v876, cmp2100, v877, cmp2104, v878, cmp2107, v879, cmp2110, v880, cmp2113, v881, cmp2116, v882, cmp2119, v883, cmp2122, v884, cmp2125, v885, cmp2128, v886, loadedv2132, v887, result_symbol2134, v888, mark_end2135, v889, v890, v891, cmp2136, v892, cmp2139, v893, cmp2142, v894, cmp2145, v895, cmp2148, v896, cmp2151, v897, cmp2154, v898, cmp2157, v899, cmp2160, v900, loadedv2164, v901, result_symbol2166, v902, mark_end2167, v903, v904, v905, loadedv2168, v906, result_symbol2170, v907, mark_end2171, v908, v909, v910, loadedv2172, v911, result_symbol2174, v912, mark_end2175, v913, v914, v915, cmp2176, v916, loadedv2180, v917

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
	i72 = libc.Ptr(&new(struct {
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
	local_advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](local_advance)
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
		goto sw_bb41
	case 2:
		goto sw_bb47
	case 3:
		goto sw_bb53
	case 4:
		goto sw_bb59
	case 5:
		goto sw_bb65
	case 6:
		goto sw_bb71
	case 7:
		goto sw_bb92
	case 8:
		goto sw_bb98
	case 9:
		goto sw_bb104
	case 10:
		goto sw_bb110
	case 11:
		goto sw_bb116
	case 12:
		goto sw_bb122
	case 13:
		goto sw_bb148
	case 14:
		goto sw_bb154
	case 15:
		goto sw_bb160
	case 16:
		goto sw_bb166
	case 17:
		goto sw_bb172
	case 18:
		goto sw_bb178
	case 19:
		goto sw_bb184
	case 20:
		goto sw_bb190
	case 21:
		goto sw_bb196
	case 22:
		goto sw_bb202
	case 23:
		goto sw_bb208
	case 24:
		goto sw_bb214
	case 25:
		goto sw_bb220
	case 26:
		goto sw_bb226
	case 27:
		goto sw_bb232
	case 28:
		goto sw_bb238
	case 29:
		goto sw_bb244
	case 30:
		goto sw_bb250
	case 31:
		goto sw_bb256
	case 32:
		goto sw_bb262
	case 33:
		goto sw_bb268
	case 34:
		goto sw_bb274
	case 35:
		goto sw_bb280
	case 36:
		goto sw_bb286
	case 37:
		goto sw_bb292
	case 38:
		goto sw_bb298
	case 39:
		goto sw_bb304
	case 40:
		goto sw_bb310
	case 41:
		goto sw_bb316
	case 42:
		goto sw_bb322
	case 43:
		goto sw_bb328
	case 44:
		goto sw_bb334
	case 45:
		goto sw_bb340
	case 46:
		goto sw_bb346
	case 47:
		goto sw_bb352
	case 48:
		goto sw_bb358
	case 49:
		goto sw_bb364
	case 50:
		goto sw_bb370
	case 51:
		goto sw_bb376
	case 52:
		goto sw_bb382
	case 53:
		goto sw_bb388
	case 54:
		goto sw_bb394
	case 55:
		goto sw_bb400
	case 56:
		goto sw_bb406
	case 57:
		goto sw_bb412
	case 58:
		goto sw_bb418
	case 59:
		goto sw_bb424
	case 60:
		goto sw_bb430
	case 61:
		goto sw_bb436
	case 62:
		goto sw_bb442
	case 63:
		goto sw_bb448
	case 64:
		goto sw_bb454
	case 65:
		goto sw_bb460
	case 66:
		goto sw_bb466
	case 67:
		goto sw_bb472
	case 68:
		goto sw_bb478
	case 69:
		goto sw_bb484
	case 70:
		goto sw_bb490
	case 71:
		goto sw_bb496
	case 72:
		goto sw_bb502
	case 73:
		goto sw_bb508
	case 74:
		goto sw_bb514
	case 75:
		goto sw_bb520
	case 76:
		goto sw_bb526
	case 77:
		goto sw_bb556
	case 78:
		goto sw_bb573
	case 79:
		goto sw_bb575
	case 80:
		goto sw_bb589
	case 81:
		goto sw_bb593
	case 82:
		goto sw_bb597
	case 83:
		goto sw_bb601
	case 84:
		goto sw_bb605
	case 85:
		goto sw_bb609
	case 86:
		goto sw_bb613
	case 87:
		goto sw_bb617
	case 88:
		goto sw_bb621
	case 89:
		goto sw_bb657
	case 90:
		goto sw_bb693
	case 91:
		goto sw_bb729
	case 92:
		goto sw_bb765
	case 93:
		goto sw_bb801
	case 94:
		goto sw_bb837
	case 95:
		goto sw_bb873
	case 96:
		goto sw_bb909
	case 97:
		goto sw_bb945
	case 98:
		goto sw_bb981
	case 99:
		goto sw_bb1017
	case 100:
		goto sw_bb1053
	case 101:
		goto sw_bb1089
	case 102:
		goto sw_bb1125
	case 103:
		goto sw_bb1161
	case 104:
		goto sw_bb1197
	case 105:
		goto sw_bb1233
	case 106:
		goto sw_bb1269
	case 107:
		goto sw_bb1305
	case 108:
		goto sw_bb1341
	case 109:
		goto sw_bb1377
	case 110:
		goto sw_bb1413
	case 111:
		goto sw_bb1449
	case 112:
		goto sw_bb1485
	case 113:
		goto sw_bb1521
	case 114:
		goto sw_bb1557
	case 115:
		goto sw_bb1593
	case 116:
		goto sw_bb1629
	case 117:
		goto sw_bb1665
	case 118:
		goto sw_bb1701
	case 119:
		goto sw_bb1737
	case 120:
		goto sw_bb1773
	case 121:
		goto sw_bb1809
	case 122:
		goto sw_bb1845
	case 123:
		goto sw_bb1881
	case 124:
		goto sw_bb1917
	case 125:
		goto sw_bb1953
	case 126:
		goto sw_bb1989
	case 127:
		goto sw_bb2025
	case 128:
		goto sw_bb2061
	case 129:
		goto sw_bb2097
	case 130:
		goto sw_bb2133
	case 131:
		goto sw_bb2165
	case 132:
		goto sw_bb2169
	case 133:
		goto sw_bb2173
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
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v11 = *libc.As[int32](i)
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(30)
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
	cmp14 = v18 == 45
	if cmp14 {
		goto if_then38
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v19 = *libc.As[int32](lookahead)
	cmp16 = v19 == 46
	if cmp16 {
		goto if_then38
	} else {
		goto lor_lhs_false18
	}

lor_lhs_false18:
	v20 = *libc.As[int32](lookahead)
	cmp19 = 48 <= v20
	if cmp19 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false23
	}

land_lhs_true:
	v21 = *libc.As[int32](lookahead)
	cmp21 = v21 <= 57
	if cmp21 {
		goto if_then38
	} else {
		goto lor_lhs_false23
	}

lor_lhs_false23:
	v22 = *libc.As[int32](lookahead)
	cmp24 = 65 <= v22
	if cmp24 {
		goto land_lhs_true26
	} else {
		goto lor_lhs_false29
	}

land_lhs_true26:
	v23 = *libc.As[int32](lookahead)
	cmp27 = v23 <= 90
	if cmp27 {
		goto if_then38
	} else {
		goto lor_lhs_false29
	}

lor_lhs_false29:
	v24 = *libc.As[int32](lookahead)
	cmp30 = v24 == 95
	if cmp30 {
		goto if_then38
	} else {
		goto lor_lhs_false32
	}

lor_lhs_false32:
	v25 = *libc.As[int32](lookahead)
	cmp33 = 98 <= v25
	if cmp33 {
		goto land_lhs_true35
	} else {
		goto if_end39
	}

land_lhs_true35:
	v26 = *libc.As[int32](lookahead)
	cmp36 = v26 <= 122
	if cmp36 {
		goto if_then38
	} else {
		goto if_end39
	}

if_then38:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end39:
	v27 = *libc.As[byte](result)
	loadedv40 = (v27 & 1) != 0
	*libc.As[bool](retval) = loadedv40
	goto _return

sw_bb41:
	v28 = *libc.As[int32](lookahead)
	cmp42 = v28 == 45
	if cmp42 {
		goto if_then44
	} else {
		goto if_end45
	}

if_then44:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end45:
	v29 = *libc.As[byte](result)
	loadedv46 = (v29 & 1) != 0
	*libc.As[bool](retval) = loadedv46
	goto _return

sw_bb47:
	v30 = *libc.As[int32](lookahead)
	cmp48 = v30 == 97
	if cmp48 {
		goto if_then50
	} else {
		goto if_end51
	}

if_then50:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end51:
	v31 = *libc.As[byte](result)
	loadedv52 = (v31 & 1) != 0
	*libc.As[bool](retval) = loadedv52
	goto _return

sw_bb53:
	v32 = *libc.As[int32](lookahead)
	cmp54 = v32 == 97
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end57:
	v33 = *libc.As[byte](result)
	loadedv58 = (v33 & 1) != 0
	*libc.As[bool](retval) = loadedv58
	goto _return

sw_bb59:
	v34 = *libc.As[int32](lookahead)
	cmp60 = v34 == 97
	if cmp60 {
		goto if_then62
	} else {
		goto if_end63
	}

if_then62:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end63:
	v35 = *libc.As[byte](result)
	loadedv64 = (v35 & 1) != 0
	*libc.As[bool](retval) = loadedv64
	goto _return

sw_bb65:
	v36 = *libc.As[int32](lookahead)
	cmp66 = v36 == 97
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end69:
	v37 = *libc.As[byte](result)
	loadedv70 = (v37 & 1) != 0
	*libc.As[bool](retval) = loadedv70
	goto _return

sw_bb71:
	*libc.As[int32](i72) = 0
	goto for_cond73

for_cond73:
	v38 = *libc.As[int32](i72)
	conv74 = int64(uint64(uint32(v38)))
	cmp75 = uint64(conv74) < uint64(20)
	if cmp75 {
		goto for_body77
	} else {
		goto for_end90
	}

for_body77:
	v39 = *libc.As[int32](i72)
	idxprom78 = int64(uint64(uint32(v39)))
	arrayidx79 = libc.Ptr(&ts_lex_map_31[idxprom78])
	v40 = *libc.As[int16](arrayidx79)
	conv80 = int32(uint32(uint16(v40)))
	v41 = *libc.As[int32](lookahead)
	cmp81 = conv80 == v41
	if cmp81 {
		goto if_then83
	} else {
		goto if_end87
	}

if_then83:
	v42 = *libc.As[int32](i72)
	add84 = v42 + 1
	idxprom85 = int64(uint64(uint32(add84)))
	arrayidx86 = libc.Ptr(&ts_lex_map_31[idxprom85])
	v43 = *libc.As[int16](arrayidx86)
	*libc.As[int16](state_addr) = v43
	goto next_state

if_end87:
	goto for_inc88

for_inc88:
	v44 = *libc.As[int32](i72)
	add89 = v44 + 2
	*libc.As[int32](i72) = add89
	goto for_cond73

for_end90:
	v45 = *libc.As[byte](result)
	loadedv91 = (v45 & 1) != 0
	*libc.As[bool](retval) = loadedv91
	goto _return

sw_bb92:
	v46 = *libc.As[int32](lookahead)
	cmp93 = v46 == 97
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end96:
	v47 = *libc.As[byte](result)
	loadedv97 = (v47 & 1) != 0
	*libc.As[bool](retval) = loadedv97
	goto _return

sw_bb98:
	v48 = *libc.As[int32](lookahead)
	cmp99 = v48 == 97
	if cmp99 {
		goto if_then101
	} else {
		goto if_end102
	}

if_then101:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end102:
	v49 = *libc.As[byte](result)
	loadedv103 = (v49 & 1) != 0
	*libc.As[bool](retval) = loadedv103
	goto _return

sw_bb104:
	v50 = *libc.As[int32](lookahead)
	cmp105 = v50 == 97
	if cmp105 {
		goto if_then107
	} else {
		goto if_end108
	}

if_then107:
	*libc.As[int16](state_addr) = 61
	goto next_state

if_end108:
	v51 = *libc.As[byte](result)
	loadedv109 = (v51 & 1) != 0
	*libc.As[bool](retval) = loadedv109
	goto _return

sw_bb110:
	v52 = *libc.As[int32](lookahead)
	cmp111 = v52 == 97
	if cmp111 {
		goto if_then113
	} else {
		goto if_end114
	}

if_then113:
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end114:
	v53 = *libc.As[byte](result)
	loadedv115 = (v53 & 1) != 0
	*libc.As[bool](retval) = loadedv115
	goto _return

sw_bb116:
	v54 = *libc.As[int32](lookahead)
	cmp117 = v54 == 98
	if cmp117 {
		goto if_then119
	} else {
		goto if_end120
	}

if_then119:
	*libc.As[int16](state_addr) = 66
	goto next_state

if_end120:
	v55 = *libc.As[byte](result)
	loadedv121 = (v55 & 1) != 0
	*libc.As[bool](retval) = loadedv121
	goto _return

sw_bb122:
	v56 = *libc.As[int32](lookahead)
	cmp123 = v56 == 99
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end126:
	v57 = *libc.As[int32](lookahead)
	cmp127 = v57 == 101
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end130:
	v58 = *libc.As[int32](lookahead)
	cmp131 = v58 == 102
	if cmp131 {
		goto if_then133
	} else {
		goto if_end134
	}

if_then133:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end134:
	v59 = *libc.As[int32](lookahead)
	cmp135 = v59 == 108
	if cmp135 {
		goto if_then137
	} else {
		goto if_end138
	}

if_then137:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end138:
	v60 = *libc.As[int32](lookahead)
	cmp139 = v60 == 112
	if cmp139 {
		goto if_then141
	} else {
		goto if_end142
	}

if_then141:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end142:
	v61 = *libc.As[int32](lookahead)
	cmp143 = v61 == 115
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end146:
	v62 = *libc.As[byte](result)
	loadedv147 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv147
	goto _return

sw_bb148:
	v63 = *libc.As[int32](lookahead)
	cmp149 = v63 == 99
	if cmp149 {
		goto if_then151
	} else {
		goto if_end152
	}

if_then151:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end152:
	v64 = *libc.As[byte](result)
	loadedv153 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv153
	goto _return

sw_bb154:
	v65 = *libc.As[int32](lookahead)
	cmp155 = v65 == 100
	if cmp155 {
		goto if_then157
	} else {
		goto if_end158
	}

if_then157:
	*libc.As[int16](state_addr) = 131
	goto next_state

if_end158:
	v66 = *libc.As[byte](result)
	loadedv159 = (v66 & 1) != 0
	*libc.As[bool](retval) = loadedv159
	goto _return

sw_bb160:
	v67 = *libc.As[int32](lookahead)
	cmp161 = v67 == 100
	if cmp161 {
		goto if_then163
	} else {
		goto if_end164
	}

if_then163:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end164:
	v68 = *libc.As[byte](result)
	loadedv165 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv165
	goto _return

sw_bb166:
	v69 = *libc.As[int32](lookahead)
	cmp167 = v69 == 100
	if cmp167 {
		goto if_then169
	} else {
		goto if_end170
	}

if_then169:
	*libc.As[int16](state_addr) = 60
	goto next_state

if_end170:
	v70 = *libc.As[byte](result)
	loadedv171 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv171
	goto _return

sw_bb172:
	v71 = *libc.As[int32](lookahead)
	cmp173 = v71 == 101
	if cmp173 {
		goto if_then175
	} else {
		goto if_end176
	}

if_then175:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end176:
	v72 = *libc.As[byte](result)
	loadedv177 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv177
	goto _return

sw_bb178:
	v73 = *libc.As[int32](lookahead)
	cmp179 = v73 == 101
	if cmp179 {
		goto if_then181
	} else {
		goto if_end182
	}

if_then181:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end182:
	v74 = *libc.As[byte](result)
	loadedv183 = (v74 & 1) != 0
	*libc.As[bool](retval) = loadedv183
	goto _return

sw_bb184:
	v75 = *libc.As[int32](lookahead)
	cmp185 = v75 == 101
	if cmp185 {
		goto if_then187
	} else {
		goto if_end188
	}

if_then187:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end188:
	v76 = *libc.As[byte](result)
	loadedv189 = (v76 & 1) != 0
	*libc.As[bool](retval) = loadedv189
	goto _return

sw_bb190:
	v77 = *libc.As[int32](lookahead)
	cmp191 = v77 == 101
	if cmp191 {
		goto if_then193
	} else {
		goto if_end194
	}

if_then193:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end194:
	v78 = *libc.As[byte](result)
	loadedv195 = (v78 & 1) != 0
	*libc.As[bool](retval) = loadedv195
	goto _return

sw_bb196:
	v79 = *libc.As[int32](lookahead)
	cmp197 = v79 == 101
	if cmp197 {
		goto if_then199
	} else {
		goto if_end200
	}

if_then199:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end200:
	v80 = *libc.As[byte](result)
	loadedv201 = (v80 & 1) != 0
	*libc.As[bool](retval) = loadedv201
	goto _return

sw_bb202:
	v81 = *libc.As[int32](lookahead)
	cmp203 = v81 == 102
	if cmp203 {
		goto if_then205
	} else {
		goto if_end206
	}

if_then205:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end206:
	v82 = *libc.As[byte](result)
	loadedv207 = (v82 & 1) != 0
	*libc.As[bool](retval) = loadedv207
	goto _return

sw_bb208:
	v83 = *libc.As[int32](lookahead)
	cmp209 = v83 == 102
	if cmp209 {
		goto if_then211
	} else {
		goto if_end212
	}

if_then211:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end212:
	v84 = *libc.As[byte](result)
	loadedv213 = (v84 & 1) != 0
	*libc.As[bool](retval) = loadedv213
	goto _return

sw_bb214:
	v85 = *libc.As[int32](lookahead)
	cmp215 = v85 == 102
	if cmp215 {
		goto if_then217
	} else {
		goto if_end218
	}

if_then217:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end218:
	v86 = *libc.As[byte](result)
	loadedv219 = (v86 & 1) != 0
	*libc.As[bool](retval) = loadedv219
	goto _return

sw_bb220:
	v87 = *libc.As[int32](lookahead)
	cmp221 = v87 == 103
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end224:
	v88 = *libc.As[byte](result)
	loadedv225 = (v88 & 1) != 0
	*libc.As[bool](retval) = loadedv225
	goto _return

sw_bb226:
	v89 = *libc.As[int32](lookahead)
	cmp227 = v89 == 103
	if cmp227 {
		goto if_then229
	} else {
		goto if_end230
	}

if_then229:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end230:
	v90 = *libc.As[byte](result)
	loadedv231 = (v90 & 1) != 0
	*libc.As[bool](retval) = loadedv231
	goto _return

sw_bb232:
	v91 = *libc.As[int32](lookahead)
	cmp233 = v91 == 103
	if cmp233 {
		goto if_then235
	} else {
		goto if_end236
	}

if_then235:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end236:
	v92 = *libc.As[byte](result)
	loadedv237 = (v92 & 1) != 0
	*libc.As[bool](retval) = loadedv237
	goto _return

sw_bb238:
	v93 = *libc.As[int32](lookahead)
	cmp239 = v93 == 105
	if cmp239 {
		goto if_then241
	} else {
		goto if_end242
	}

if_then241:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end242:
	v94 = *libc.As[byte](result)
	loadedv243 = (v94 & 1) != 0
	*libc.As[bool](retval) = loadedv243
	goto _return

sw_bb244:
	v95 = *libc.As[int32](lookahead)
	cmp245 = v95 == 105
	if cmp245 {
		goto if_then247
	} else {
		goto if_end248
	}

if_then247:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end248:
	v96 = *libc.As[byte](result)
	loadedv249 = (v96 & 1) != 0
	*libc.As[bool](retval) = loadedv249
	goto _return

sw_bb250:
	v97 = *libc.As[int32](lookahead)
	cmp251 = v97 == 105
	if cmp251 {
		goto if_then253
	} else {
		goto if_end254
	}

if_then253:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end254:
	v98 = *libc.As[byte](result)
	loadedv255 = (v98 & 1) != 0
	*libc.As[bool](retval) = loadedv255
	goto _return

sw_bb256:
	v99 = *libc.As[int32](lookahead)
	cmp257 = v99 == 105
	if cmp257 {
		goto if_then259
	} else {
		goto if_end260
	}

if_then259:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end260:
	v100 = *libc.As[byte](result)
	loadedv261 = (v100 & 1) != 0
	*libc.As[bool](retval) = loadedv261
	goto _return

sw_bb262:
	v101 = *libc.As[int32](lookahead)
	cmp263 = v101 == 105
	if cmp263 {
		goto if_then265
	} else {
		goto if_end266
	}

if_then265:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end266:
	v102 = *libc.As[byte](result)
	loadedv267 = (v102 & 1) != 0
	*libc.As[bool](retval) = loadedv267
	goto _return

sw_bb268:
	v103 = *libc.As[int32](lookahead)
	cmp269 = v103 == 105
	if cmp269 {
		goto if_then271
	} else {
		goto if_end272
	}

if_then271:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end272:
	v104 = *libc.As[byte](result)
	loadedv273 = (v104 & 1) != 0
	*libc.As[bool](retval) = loadedv273
	goto _return

sw_bb274:
	v105 = *libc.As[int32](lookahead)
	cmp275 = v105 == 107
	if cmp275 {
		goto if_then277
	} else {
		goto if_end278
	}

if_then277:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end278:
	v106 = *libc.As[byte](result)
	loadedv279 = (v106 & 1) != 0
	*libc.As[bool](retval) = loadedv279
	goto _return

sw_bb280:
	v107 = *libc.As[int32](lookahead)
	cmp281 = v107 == 108
	if cmp281 {
		goto if_then283
	} else {
		goto if_end284
	}

if_then283:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end284:
	v108 = *libc.As[byte](result)
	loadedv285 = (v108 & 1) != 0
	*libc.As[bool](retval) = loadedv285
	goto _return

sw_bb286:
	v109 = *libc.As[int32](lookahead)
	cmp287 = v109 == 108
	if cmp287 {
		goto if_then289
	} else {
		goto if_end290
	}

if_then289:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end290:
	v110 = *libc.As[byte](result)
	loadedv291 = (v110 & 1) != 0
	*libc.As[bool](retval) = loadedv291
	goto _return

sw_bb292:
	v111 = *libc.As[int32](lookahead)
	cmp293 = v111 == 108
	if cmp293 {
		goto if_then295
	} else {
		goto if_end296
	}

if_then295:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end296:
	v112 = *libc.As[byte](result)
	loadedv297 = (v112 & 1) != 0
	*libc.As[bool](retval) = loadedv297
	goto _return

sw_bb298:
	v113 = *libc.As[int32](lookahead)
	cmp299 = v113 == 108
	if cmp299 {
		goto if_then301
	} else {
		goto if_end302
	}

if_then301:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end302:
	v114 = *libc.As[byte](result)
	loadedv303 = (v114 & 1) != 0
	*libc.As[bool](retval) = loadedv303
	goto _return

sw_bb304:
	v115 = *libc.As[int32](lookahead)
	cmp305 = v115 == 109
	if cmp305 {
		goto if_then307
	} else {
		goto if_end308
	}

if_then307:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end308:
	v116 = *libc.As[byte](result)
	loadedv309 = (v116 & 1) != 0
	*libc.As[bool](retval) = loadedv309
	goto _return

sw_bb310:
	v117 = *libc.As[int32](lookahead)
	cmp311 = v117 == 110
	if cmp311 {
		goto if_then313
	} else {
		goto if_end314
	}

if_then313:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end314:
	v118 = *libc.As[byte](result)
	loadedv315 = (v118 & 1) != 0
	*libc.As[bool](retval) = loadedv315
	goto _return

sw_bb316:
	v119 = *libc.As[int32](lookahead)
	cmp317 = v119 == 110
	if cmp317 {
		goto if_then319
	} else {
		goto if_end320
	}

if_then319:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end320:
	v120 = *libc.As[byte](result)
	loadedv321 = (v120 & 1) != 0
	*libc.As[bool](retval) = loadedv321
	goto _return

sw_bb322:
	v121 = *libc.As[int32](lookahead)
	cmp323 = v121 == 110
	if cmp323 {
		goto if_then325
	} else {
		goto if_end326
	}

if_then325:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end326:
	v122 = *libc.As[byte](result)
	loadedv327 = (v122 & 1) != 0
	*libc.As[bool](retval) = loadedv327
	goto _return

sw_bb328:
	v123 = *libc.As[int32](lookahead)
	cmp329 = v123 == 110
	if cmp329 {
		goto if_then331
	} else {
		goto if_end332
	}

if_then331:
	*libc.As[int16](state_addr) = 71
	goto next_state

if_end332:
	v124 = *libc.As[byte](result)
	loadedv333 = (v124 & 1) != 0
	*libc.As[bool](retval) = loadedv333
	goto _return

sw_bb334:
	v125 = *libc.As[int32](lookahead)
	cmp335 = v125 == 110
	if cmp335 {
		goto if_then337
	} else {
		goto if_end338
	}

if_then337:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end338:
	v126 = *libc.As[byte](result)
	loadedv339 = (v126 & 1) != 0
	*libc.As[bool](retval) = loadedv339
	goto _return

sw_bb340:
	v127 = *libc.As[int32](lookahead)
	cmp341 = v127 == 110
	if cmp341 {
		goto if_then343
	} else {
		goto if_end344
	}

if_then343:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end344:
	v128 = *libc.As[byte](result)
	loadedv345 = (v128 & 1) != 0
	*libc.As[bool](retval) = loadedv345
	goto _return

sw_bb346:
	v129 = *libc.As[int32](lookahead)
	cmp347 = v129 == 111
	if cmp347 {
		goto if_then349
	} else {
		goto if_end350
	}

if_then349:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end350:
	v130 = *libc.As[byte](result)
	loadedv351 = (v130 & 1) != 0
	*libc.As[bool](retval) = loadedv351
	goto _return

sw_bb352:
	v131 = *libc.As[int32](lookahead)
	cmp353 = v131 == 111
	if cmp353 {
		goto if_then355
	} else {
		goto if_end356
	}

if_then355:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end356:
	v132 = *libc.As[byte](result)
	loadedv357 = (v132 & 1) != 0
	*libc.As[bool](retval) = loadedv357
	goto _return

sw_bb358:
	v133 = *libc.As[int32](lookahead)
	cmp359 = v133 == 111
	if cmp359 {
		goto if_then361
	} else {
		goto if_end362
	}

if_then361:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end362:
	v134 = *libc.As[byte](result)
	loadedv363 = (v134 & 1) != 0
	*libc.As[bool](retval) = loadedv363
	goto _return

sw_bb364:
	v135 = *libc.As[int32](lookahead)
	cmp365 = v135 == 111
	if cmp365 {
		goto if_then367
	} else {
		goto if_end368
	}

if_then367:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end368:
	v136 = *libc.As[byte](result)
	loadedv369 = (v136 & 1) != 0
	*libc.As[bool](retval) = loadedv369
	goto _return

sw_bb370:
	v137 = *libc.As[int32](lookahead)
	cmp371 = v137 == 111
	if cmp371 {
		goto if_then373
	} else {
		goto if_end374
	}

if_then373:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end374:
	v138 = *libc.As[byte](result)
	loadedv375 = (v138 & 1) != 0
	*libc.As[bool](retval) = loadedv375
	goto _return

sw_bb376:
	v139 = *libc.As[int32](lookahead)
	cmp377 = v139 == 111
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end380:
	v140 = *libc.As[byte](result)
	loadedv381 = (v140 & 1) != 0
	*libc.As[bool](retval) = loadedv381
	goto _return

sw_bb382:
	v141 = *libc.As[int32](lookahead)
	cmp383 = v141 == 111
	if cmp383 {
		goto if_then385
	} else {
		goto if_end386
	}

if_then385:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end386:
	v142 = *libc.As[byte](result)
	loadedv387 = (v142 & 1) != 0
	*libc.As[bool](retval) = loadedv387
	goto _return

sw_bb388:
	v143 = *libc.As[int32](lookahead)
	cmp389 = v143 == 112
	if cmp389 {
		goto if_then391
	} else {
		goto if_end392
	}

if_then391:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end392:
	v144 = *libc.As[byte](result)
	loadedv393 = (v144 & 1) != 0
	*libc.As[bool](retval) = loadedv393
	goto _return

sw_bb394:
	v145 = *libc.As[int32](lookahead)
	cmp395 = v145 == 112
	if cmp395 {
		goto if_then397
	} else {
		goto if_end398
	}

if_then397:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end398:
	v146 = *libc.As[byte](result)
	loadedv399 = (v146 & 1) != 0
	*libc.As[bool](retval) = loadedv399
	goto _return

sw_bb400:
	v147 = *libc.As[int32](lookahead)
	cmp401 = v147 == 114
	if cmp401 {
		goto if_then403
	} else {
		goto if_end404
	}

if_then403:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end404:
	v148 = *libc.As[byte](result)
	loadedv405 = (v148 & 1) != 0
	*libc.As[bool](retval) = loadedv405
	goto _return

sw_bb406:
	v149 = *libc.As[int32](lookahead)
	cmp407 = v149 == 114
	if cmp407 {
		goto if_then409
	} else {
		goto if_end410
	}

if_then409:
	*libc.As[int16](state_addr) = 81
	goto next_state

if_end410:
	v150 = *libc.As[byte](result)
	loadedv411 = (v150 & 1) != 0
	*libc.As[bool](retval) = loadedv411
	goto _return

sw_bb412:
	v151 = *libc.As[int32](lookahead)
	cmp413 = v151 == 114
	if cmp413 {
		goto if_then415
	} else {
		goto if_end416
	}

if_then415:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end416:
	v152 = *libc.As[byte](result)
	loadedv417 = (v152 & 1) != 0
	*libc.As[bool](retval) = loadedv417
	goto _return

sw_bb418:
	v153 = *libc.As[int32](lookahead)
	cmp419 = v153 == 114
	if cmp419 {
		goto if_then421
	} else {
		goto if_end422
	}

if_then421:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end422:
	v154 = *libc.As[byte](result)
	loadedv423 = (v154 & 1) != 0
	*libc.As[bool](retval) = loadedv423
	goto _return

sw_bb424:
	v155 = *libc.As[int32](lookahead)
	cmp425 = v155 == 114
	if cmp425 {
		goto if_then427
	} else {
		goto if_end428
	}

if_then427:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end428:
	v156 = *libc.As[byte](result)
	loadedv429 = (v156 & 1) != 0
	*libc.As[bool](retval) = loadedv429
	goto _return

sw_bb430:
	v157 = *libc.As[int32](lookahead)
	cmp431 = v157 == 114
	if cmp431 {
		goto if_then433
	} else {
		goto if_end434
	}

if_then433:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end434:
	v158 = *libc.As[byte](result)
	loadedv435 = (v158 & 1) != 0
	*libc.As[bool](retval) = loadedv435
	goto _return

sw_bb436:
	v159 = *libc.As[int32](lookahead)
	cmp437 = v159 == 114
	if cmp437 {
		goto if_then439
	} else {
		goto if_end440
	}

if_then439:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end440:
	v160 = *libc.As[byte](result)
	loadedv441 = (v160 & 1) != 0
	*libc.As[bool](retval) = loadedv441
	goto _return

sw_bb442:
	v161 = *libc.As[int32](lookahead)
	cmp443 = v161 == 114
	if cmp443 {
		goto if_then445
	} else {
		goto if_end446
	}

if_then445:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end446:
	v162 = *libc.As[byte](result)
	loadedv447 = (v162 & 1) != 0
	*libc.As[bool](retval) = loadedv447
	goto _return

sw_bb448:
	v163 = *libc.As[int32](lookahead)
	cmp449 = v163 == 115
	if cmp449 {
		goto if_then451
	} else {
		goto if_end452
	}

if_then451:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end452:
	v164 = *libc.As[byte](result)
	loadedv453 = (v164 & 1) != 0
	*libc.As[bool](retval) = loadedv453
	goto _return

sw_bb454:
	v165 = *libc.As[int32](lookahead)
	cmp455 = v165 == 115
	if cmp455 {
		goto if_then457
	} else {
		goto if_end458
	}

if_then457:
	*libc.As[int16](state_addr) = 131
	goto next_state

if_end458:
	v166 = *libc.As[byte](result)
	loadedv459 = (v166 & 1) != 0
	*libc.As[bool](retval) = loadedv459
	goto _return

sw_bb460:
	v167 = *libc.As[int32](lookahead)
	cmp461 = v167 == 115
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*libc.As[int16](state_addr) = 69
	goto next_state

if_end464:
	v168 = *libc.As[byte](result)
	loadedv465 = (v168 & 1) != 0
	*libc.As[bool](retval) = loadedv465
	goto _return

sw_bb466:
	v169 = *libc.As[int32](lookahead)
	cmp467 = v169 == 115
	if cmp467 {
		goto if_then469
	} else {
		goto if_end470
	}

if_then469:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end470:
	v170 = *libc.As[byte](result)
	loadedv471 = (v170 & 1) != 0
	*libc.As[bool](retval) = loadedv471
	goto _return

sw_bb472:
	v171 = *libc.As[int32](lookahead)
	cmp473 = v171 == 116
	if cmp473 {
		goto if_then475
	} else {
		goto if_end476
	}

if_then475:
	*libc.As[int16](state_addr) = 80
	goto next_state

if_end476:
	v172 = *libc.As[byte](result)
	loadedv477 = (v172 & 1) != 0
	*libc.As[bool](retval) = loadedv477
	goto _return

sw_bb478:
	v173 = *libc.As[int32](lookahead)
	cmp479 = v173 == 116
	if cmp479 {
		goto if_then481
	} else {
		goto if_end482
	}

if_then481:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end482:
	v174 = *libc.As[byte](result)
	loadedv483 = (v174 & 1) != 0
	*libc.As[bool](retval) = loadedv483
	goto _return

sw_bb484:
	v175 = *libc.As[int32](lookahead)
	cmp485 = v175 == 116
	if cmp485 {
		goto if_then487
	} else {
		goto if_end488
	}

if_then487:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end488:
	v176 = *libc.As[byte](result)
	loadedv489 = (v176 & 1) != 0
	*libc.As[bool](retval) = loadedv489
	goto _return

sw_bb490:
	v177 = *libc.As[int32](lookahead)
	cmp491 = v177 == 116
	if cmp491 {
		goto if_then493
	} else {
		goto if_end494
	}

if_then493:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end494:
	v178 = *libc.As[byte](result)
	loadedv495 = (v178 & 1) != 0
	*libc.As[bool](retval) = loadedv495
	goto _return

sw_bb496:
	v179 = *libc.As[int32](lookahead)
	cmp497 = v179 == 117
	if cmp497 {
		goto if_then499
	} else {
		goto if_end500
	}

if_then499:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end500:
	v180 = *libc.As[byte](result)
	loadedv501 = (v180 & 1) != 0
	*libc.As[bool](retval) = loadedv501
	goto _return

sw_bb502:
	v181 = *libc.As[int32](lookahead)
	cmp503 = v181 == 117
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end506:
	v182 = *libc.As[byte](result)
	loadedv507 = (v182 & 1) != 0
	*libc.As[bool](retval) = loadedv507
	goto _return

sw_bb508:
	v183 = *libc.As[int32](lookahead)
	cmp509 = v183 == 119
	if cmp509 {
		goto if_then511
	} else {
		goto if_end512
	}

if_then511:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end512:
	v184 = *libc.As[byte](result)
	loadedv513 = (v184 & 1) != 0
	*libc.As[bool](retval) = loadedv513
	goto _return

sw_bb514:
	v185 = *libc.As[int32](lookahead)
	cmp515 = v185 == 120
	if cmp515 {
		goto if_then517
	} else {
		goto if_end518
	}

if_then517:
	*libc.As[int16](state_addr) = 131
	goto next_state

if_end518:
	v186 = *libc.As[byte](result)
	loadedv519 = (v186 & 1) != 0
	*libc.As[bool](retval) = loadedv519
	goto _return

sw_bb520:
	v187 = *libc.As[int32](lookahead)
	cmp521 = v187 == 121
	if cmp521 {
		goto if_then523
	} else {
		goto if_end524
	}

if_then523:
	*libc.As[int16](state_addr) = 131
	goto next_state

if_end524:
	v188 = *libc.As[byte](result)
	loadedv525 = (v188 & 1) != 0
	*libc.As[bool](retval) = loadedv525
	goto _return

sw_bb526:
	v189 = *libc.As[int32](lookahead)
	cmp527 = v189 == 45
	if cmp527 {
		goto if_then553
	} else {
		goto lor_lhs_false529
	}

lor_lhs_false529:
	v190 = *libc.As[int32](lookahead)
	cmp530 = v190 == 46
	if cmp530 {
		goto if_then553
	} else {
		goto lor_lhs_false532
	}

lor_lhs_false532:
	v191 = *libc.As[int32](lookahead)
	cmp533 = 48 <= v191
	if cmp533 {
		goto land_lhs_true535
	} else {
		goto lor_lhs_false538
	}

land_lhs_true535:
	v192 = *libc.As[int32](lookahead)
	cmp536 = v192 <= 57
	if cmp536 {
		goto if_then553
	} else {
		goto lor_lhs_false538
	}

lor_lhs_false538:
	v193 = *libc.As[int32](lookahead)
	cmp539 = 65 <= v193
	if cmp539 {
		goto land_lhs_true541
	} else {
		goto lor_lhs_false544
	}

land_lhs_true541:
	v194 = *libc.As[int32](lookahead)
	cmp542 = v194 <= 90
	if cmp542 {
		goto if_then553
	} else {
		goto lor_lhs_false544
	}

lor_lhs_false544:
	v195 = *libc.As[int32](lookahead)
	cmp545 = v195 == 95
	if cmp545 {
		goto if_then553
	} else {
		goto lor_lhs_false547
	}

lor_lhs_false547:
	v196 = *libc.As[int32](lookahead)
	cmp548 = 97 <= v196
	if cmp548 {
		goto land_lhs_true550
	} else {
		goto if_end554
	}

land_lhs_true550:
	v197 = *libc.As[int32](lookahead)
	cmp551 = v197 <= 122
	if cmp551 {
		goto if_then553
	} else {
		goto if_end554
	}

if_then553:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end554:
	v198 = *libc.As[byte](result)
	loadedv555 = (v198 & 1) != 0
	*libc.As[bool](retval) = loadedv555
	goto _return

sw_bb556:
	v199 = *libc.As[byte](eof)
	loadedv557 = (v199 & 1) != 0
	if loadedv557 {
		goto if_then558
	} else {
		goto if_end559
	}

if_then558:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end559:
	v200 = *libc.As[int32](lookahead)
	cmp560 = v200 == 10
	if cmp560 {
		goto if_then562
	} else {
		goto if_end563
	}

if_then562:
	*libc.As[int16](state_addr) = 132
	goto next_state

if_end563:
	v201 = *libc.As[int32](lookahead)
	cmp564 = v201 == 13
	if cmp564 {
		goto if_then566
	} else {
		goto if_end567
	}

if_then566:
	*libc.As[int16](state_addr) = 133
	goto next_state

if_end567:
	v202 = *libc.As[int32](lookahead)
	cmp568 = v202 != 0
	if cmp568 {
		goto if_then570
	} else {
		goto if_end571
	}

if_then570:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end571:
	v203 = *libc.As[byte](result)
	loadedv572 = (v203 & 1) != 0
	*libc.As[bool](retval) = loadedv572
	goto _return

sw_bb573:
	*libc.As[byte](result) = 1
	v204 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v204).F1)
	*libc.As[int16](result_symbol) = 0
	v205 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v205).F3)
	v206 = *libc.As[unsafe.Pointer](mark_end)
	v207 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v206)(v207)
	v208 = *libc.As[byte](result)
	loadedv574 = (v208 & 1) != 0
	*libc.As[bool](retval) = loadedv574
	goto _return

sw_bb575:
	*libc.As[byte](result) = 1
	v209 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol576 = libc.Ptr(&libc.As[TSLexer](v209).F1)
	*libc.As[int16](result_symbol576) = 1
	v210 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end577 = libc.Ptr(&libc.As[TSLexer](v210).F3)
	v211 = *libc.As[unsafe.Pointer](mark_end577)
	v212 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v211)(v212)
	v213 = *libc.As[int32](lookahead)
	cmp578 = v213 != 0
	if cmp578 {
		goto land_lhs_true580
	} else {
		goto if_end587
	}

land_lhs_true580:
	v214 = *libc.As[int32](lookahead)
	cmp581 = v214 != 10
	if cmp581 {
		goto land_lhs_true583
	} else {
		goto if_end587
	}

land_lhs_true583:
	v215 = *libc.As[int32](lookahead)
	cmp584 = v215 != 13
	if cmp584 {
		goto if_then586
	} else {
		goto if_end587
	}

if_then586:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end587:
	v216 = *libc.As[byte](result)
	loadedv588 = (v216 & 1) != 0
	*libc.As[bool](retval) = loadedv588
	goto _return

sw_bb589:
	*libc.As[byte](result) = 1
	v217 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol590 = libc.Ptr(&libc.As[TSLexer](v217).F1)
	*libc.As[int16](result_symbol590) = 2
	v218 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end591 = libc.Ptr(&libc.As[TSLexer](v218).F3)
	v219 = *libc.As[unsafe.Pointer](mark_end591)
	v220 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v219)(v220)
	v221 = *libc.As[byte](result)
	loadedv592 = (v221 & 1) != 0
	*libc.As[bool](retval) = loadedv592
	goto _return

sw_bb593:
	*libc.As[byte](result) = 1
	v222 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol594 = libc.Ptr(&libc.As[TSLexer](v222).F1)
	*libc.As[int16](result_symbol594) = 3
	v223 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end595 = libc.Ptr(&libc.As[TSLexer](v223).F3)
	v224 = *libc.As[unsafe.Pointer](mark_end595)
	v225 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v224)(v225)
	v226 = *libc.As[byte](result)
	loadedv596 = (v226 & 1) != 0
	*libc.As[bool](retval) = loadedv596
	goto _return

sw_bb597:
	*libc.As[byte](result) = 1
	v227 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol598 = libc.Ptr(&libc.As[TSLexer](v227).F1)
	*libc.As[int16](result_symbol598) = 4
	v228 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end599 = libc.Ptr(&libc.As[TSLexer](v228).F3)
	v229 = *libc.As[unsafe.Pointer](mark_end599)
	v230 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v229)(v230)
	v231 = *libc.As[byte](result)
	loadedv600 = (v231 & 1) != 0
	*libc.As[bool](retval) = loadedv600
	goto _return

sw_bb601:
	*libc.As[byte](result) = 1
	v232 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol602 = libc.Ptr(&libc.As[TSLexer](v232).F1)
	*libc.As[int16](result_symbol602) = 5
	v233 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end603 = libc.Ptr(&libc.As[TSLexer](v233).F3)
	v234 = *libc.As[unsafe.Pointer](mark_end603)
	v235 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v234)(v235)
	v236 = *libc.As[byte](result)
	loadedv604 = (v236 & 1) != 0
	*libc.As[bool](retval) = loadedv604
	goto _return

sw_bb605:
	*libc.As[byte](result) = 1
	v237 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol606 = libc.Ptr(&libc.As[TSLexer](v237).F1)
	*libc.As[int16](result_symbol606) = 6
	v238 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end607 = libc.Ptr(&libc.As[TSLexer](v238).F3)
	v239 = *libc.As[unsafe.Pointer](mark_end607)
	v240 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v239)(v240)
	v241 = *libc.As[byte](result)
	loadedv608 = (v241 & 1) != 0
	*libc.As[bool](retval) = loadedv608
	goto _return

sw_bb609:
	*libc.As[byte](result) = 1
	v242 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol610 = libc.Ptr(&libc.As[TSLexer](v242).F1)
	*libc.As[int16](result_symbol610) = 7
	v243 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end611 = libc.Ptr(&libc.As[TSLexer](v243).F3)
	v244 = *libc.As[unsafe.Pointer](mark_end611)
	v245 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v244)(v245)
	v246 = *libc.As[byte](result)
	loadedv612 = (v246 & 1) != 0
	*libc.As[bool](retval) = loadedv612
	goto _return

sw_bb613:
	*libc.As[byte](result) = 1
	v247 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol614 = libc.Ptr(&libc.As[TSLexer](v247).F1)
	*libc.As[int16](result_symbol614) = 8
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end615 = libc.Ptr(&libc.As[TSLexer](v248).F3)
	v249 = *libc.As[unsafe.Pointer](mark_end615)
	v250 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v249)(v250)
	v251 = *libc.As[byte](result)
	loadedv616 = (v251 & 1) != 0
	*libc.As[bool](retval) = loadedv616
	goto _return

sw_bb617:
	*libc.As[byte](result) = 1
	v252 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol618 = libc.Ptr(&libc.As[TSLexer](v252).F1)
	*libc.As[int16](result_symbol618) = 9
	v253 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end619 = libc.Ptr(&libc.As[TSLexer](v253).F3)
	v254 = *libc.As[unsafe.Pointer](mark_end619)
	v255 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v254)(v255)
	v256 = *libc.As[byte](result)
	loadedv620 = (v256 & 1) != 0
	*libc.As[bool](retval) = loadedv620
	goto _return

sw_bb621:
	*libc.As[byte](result) = 1
	v257 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol622 = libc.Ptr(&libc.As[TSLexer](v257).F1)
	*libc.As[int16](result_symbol622) = 10
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end623 = libc.Ptr(&libc.As[TSLexer](v258).F3)
	v259 = *libc.As[unsafe.Pointer](mark_end623)
	v260 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v259)(v260)
	v261 = *libc.As[int32](lookahead)
	cmp624 = v261 == 97
	if cmp624 {
		goto if_then626
	} else {
		goto if_end627
	}

if_then626:
	*libc.As[int16](state_addr) = 92
	goto next_state

if_end627:
	v262 = *libc.As[int32](lookahead)
	cmp628 = v262 == 45
	if cmp628 {
		goto if_then654
	} else {
		goto lor_lhs_false630
	}

lor_lhs_false630:
	v263 = *libc.As[int32](lookahead)
	cmp631 = v263 == 46
	if cmp631 {
		goto if_then654
	} else {
		goto lor_lhs_false633
	}

lor_lhs_false633:
	v264 = *libc.As[int32](lookahead)
	cmp634 = 48 <= v264
	if cmp634 {
		goto land_lhs_true636
	} else {
		goto lor_lhs_false639
	}

land_lhs_true636:
	v265 = *libc.As[int32](lookahead)
	cmp637 = v265 <= 57
	if cmp637 {
		goto if_then654
	} else {
		goto lor_lhs_false639
	}

lor_lhs_false639:
	v266 = *libc.As[int32](lookahead)
	cmp640 = 65 <= v266
	if cmp640 {
		goto land_lhs_true642
	} else {
		goto lor_lhs_false645
	}

land_lhs_true642:
	v267 = *libc.As[int32](lookahead)
	cmp643 = v267 <= 90
	if cmp643 {
		goto if_then654
	} else {
		goto lor_lhs_false645
	}

lor_lhs_false645:
	v268 = *libc.As[int32](lookahead)
	cmp646 = v268 == 95
	if cmp646 {
		goto if_then654
	} else {
		goto lor_lhs_false648
	}

lor_lhs_false648:
	v269 = *libc.As[int32](lookahead)
	cmp649 = 98 <= v269
	if cmp649 {
		goto land_lhs_true651
	} else {
		goto if_end655
	}

land_lhs_true651:
	v270 = *libc.As[int32](lookahead)
	cmp652 = v270 <= 122
	if cmp652 {
		goto if_then654
	} else {
		goto if_end655
	}

if_then654:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end655:
	v271 = *libc.As[byte](result)
	loadedv656 = (v271 & 1) != 0
	*libc.As[bool](retval) = loadedv656
	goto _return

sw_bb657:
	*libc.As[byte](result) = 1
	v272 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol658 = libc.Ptr(&libc.As[TSLexer](v272).F1)
	*libc.As[int16](result_symbol658) = 10
	v273 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end659 = libc.Ptr(&libc.As[TSLexer](v273).F3)
	v274 = *libc.As[unsafe.Pointer](mark_end659)
	v275 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v274)(v275)
	v276 = *libc.As[int32](lookahead)
	cmp660 = v276 == 97
	if cmp660 {
		goto if_then662
	} else {
		goto if_end663
	}

if_then662:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end663:
	v277 = *libc.As[int32](lookahead)
	cmp664 = v277 == 45
	if cmp664 {
		goto if_then690
	} else {
		goto lor_lhs_false666
	}

lor_lhs_false666:
	v278 = *libc.As[int32](lookahead)
	cmp667 = v278 == 46
	if cmp667 {
		goto if_then690
	} else {
		goto lor_lhs_false669
	}

lor_lhs_false669:
	v279 = *libc.As[int32](lookahead)
	cmp670 = 48 <= v279
	if cmp670 {
		goto land_lhs_true672
	} else {
		goto lor_lhs_false675
	}

land_lhs_true672:
	v280 = *libc.As[int32](lookahead)
	cmp673 = v280 <= 57
	if cmp673 {
		goto if_then690
	} else {
		goto lor_lhs_false675
	}

lor_lhs_false675:
	v281 = *libc.As[int32](lookahead)
	cmp676 = 65 <= v281
	if cmp676 {
		goto land_lhs_true678
	} else {
		goto lor_lhs_false681
	}

land_lhs_true678:
	v282 = *libc.As[int32](lookahead)
	cmp679 = v282 <= 90
	if cmp679 {
		goto if_then690
	} else {
		goto lor_lhs_false681
	}

lor_lhs_false681:
	v283 = *libc.As[int32](lookahead)
	cmp682 = v283 == 95
	if cmp682 {
		goto if_then690
	} else {
		goto lor_lhs_false684
	}

lor_lhs_false684:
	v284 = *libc.As[int32](lookahead)
	cmp685 = 98 <= v284
	if cmp685 {
		goto land_lhs_true687
	} else {
		goto if_end691
	}

land_lhs_true687:
	v285 = *libc.As[int32](lookahead)
	cmp688 = v285 <= 122
	if cmp688 {
		goto if_then690
	} else {
		goto if_end691
	}

if_then690:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end691:
	v286 = *libc.As[byte](result)
	loadedv692 = (v286 & 1) != 0
	*libc.As[bool](retval) = loadedv692
	goto _return

sw_bb693:
	*libc.As[byte](result) = 1
	v287 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol694 = libc.Ptr(&libc.As[TSLexer](v287).F1)
	*libc.As[int16](result_symbol694) = 10
	v288 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end695 = libc.Ptr(&libc.As[TSLexer](v288).F3)
	v289 = *libc.As[unsafe.Pointer](mark_end695)
	v290 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v289)(v290)
	v291 = *libc.As[int32](lookahead)
	cmp696 = v291 == 97
	if cmp696 {
		goto if_then698
	} else {
		goto if_end699
	}

if_then698:
	*libc.As[int16](state_addr) = 122
	goto next_state

if_end699:
	v292 = *libc.As[int32](lookahead)
	cmp700 = v292 == 45
	if cmp700 {
		goto if_then726
	} else {
		goto lor_lhs_false702
	}

lor_lhs_false702:
	v293 = *libc.As[int32](lookahead)
	cmp703 = v293 == 46
	if cmp703 {
		goto if_then726
	} else {
		goto lor_lhs_false705
	}

lor_lhs_false705:
	v294 = *libc.As[int32](lookahead)
	cmp706 = 48 <= v294
	if cmp706 {
		goto land_lhs_true708
	} else {
		goto lor_lhs_false711
	}

land_lhs_true708:
	v295 = *libc.As[int32](lookahead)
	cmp709 = v295 <= 57
	if cmp709 {
		goto if_then726
	} else {
		goto lor_lhs_false711
	}

lor_lhs_false711:
	v296 = *libc.As[int32](lookahead)
	cmp712 = 65 <= v296
	if cmp712 {
		goto land_lhs_true714
	} else {
		goto lor_lhs_false717
	}

land_lhs_true714:
	v297 = *libc.As[int32](lookahead)
	cmp715 = v297 <= 90
	if cmp715 {
		goto if_then726
	} else {
		goto lor_lhs_false717
	}

lor_lhs_false717:
	v298 = *libc.As[int32](lookahead)
	cmp718 = v298 == 95
	if cmp718 {
		goto if_then726
	} else {
		goto lor_lhs_false720
	}

lor_lhs_false720:
	v299 = *libc.As[int32](lookahead)
	cmp721 = 98 <= v299
	if cmp721 {
		goto land_lhs_true723
	} else {
		goto if_end727
	}

land_lhs_true723:
	v300 = *libc.As[int32](lookahead)
	cmp724 = v300 <= 122
	if cmp724 {
		goto if_then726
	} else {
		goto if_end727
	}

if_then726:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end727:
	v301 = *libc.As[byte](result)
	loadedv728 = (v301 & 1) != 0
	*libc.As[bool](retval) = loadedv728
	goto _return

sw_bb729:
	*libc.As[byte](result) = 1
	v302 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol730 = libc.Ptr(&libc.As[TSLexer](v302).F1)
	*libc.As[int16](result_symbol730) = 10
	v303 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end731 = libc.Ptr(&libc.As[TSLexer](v303).F3)
	v304 = *libc.As[unsafe.Pointer](mark_end731)
	v305 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v304)(v305)
	v306 = *libc.As[int32](lookahead)
	cmp732 = v306 == 98
	if cmp732 {
		goto if_then734
	} else {
		goto if_end735
	}

if_then734:
	*libc.As[int16](state_addr) = 124
	goto next_state

if_end735:
	v307 = *libc.As[int32](lookahead)
	cmp736 = v307 == 45
	if cmp736 {
		goto if_then762
	} else {
		goto lor_lhs_false738
	}

lor_lhs_false738:
	v308 = *libc.As[int32](lookahead)
	cmp739 = v308 == 46
	if cmp739 {
		goto if_then762
	} else {
		goto lor_lhs_false741
	}

lor_lhs_false741:
	v309 = *libc.As[int32](lookahead)
	cmp742 = 48 <= v309
	if cmp742 {
		goto land_lhs_true744
	} else {
		goto lor_lhs_false747
	}

land_lhs_true744:
	v310 = *libc.As[int32](lookahead)
	cmp745 = v310 <= 57
	if cmp745 {
		goto if_then762
	} else {
		goto lor_lhs_false747
	}

lor_lhs_false747:
	v311 = *libc.As[int32](lookahead)
	cmp748 = 65 <= v311
	if cmp748 {
		goto land_lhs_true750
	} else {
		goto lor_lhs_false753
	}

land_lhs_true750:
	v312 = *libc.As[int32](lookahead)
	cmp751 = v312 <= 90
	if cmp751 {
		goto if_then762
	} else {
		goto lor_lhs_false753
	}

lor_lhs_false753:
	v313 = *libc.As[int32](lookahead)
	cmp754 = v313 == 95
	if cmp754 {
		goto if_then762
	} else {
		goto lor_lhs_false756
	}

lor_lhs_false756:
	v314 = *libc.As[int32](lookahead)
	cmp757 = 97 <= v314
	if cmp757 {
		goto land_lhs_true759
	} else {
		goto if_end763
	}

land_lhs_true759:
	v315 = *libc.As[int32](lookahead)
	cmp760 = v315 <= 122
	if cmp760 {
		goto if_then762
	} else {
		goto if_end763
	}

if_then762:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end763:
	v316 = *libc.As[byte](result)
	loadedv764 = (v316 & 1) != 0
	*libc.As[bool](retval) = loadedv764
	goto _return

sw_bb765:
	*libc.As[byte](result) = 1
	v317 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol766 = libc.Ptr(&libc.As[TSLexer](v317).F1)
	*libc.As[int16](result_symbol766) = 10
	v318 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end767 = libc.Ptr(&libc.As[TSLexer](v318).F3)
	v319 = *libc.As[unsafe.Pointer](mark_end767)
	v320 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v319)(v320)
	v321 = *libc.As[int32](lookahead)
	cmp768 = v321 == 99
	if cmp768 {
		goto if_then770
	} else {
		goto if_end771
	}

if_then770:
	*libc.As[int16](state_addr) = 113
	goto next_state

if_end771:
	v322 = *libc.As[int32](lookahead)
	cmp772 = v322 == 45
	if cmp772 {
		goto if_then798
	} else {
		goto lor_lhs_false774
	}

lor_lhs_false774:
	v323 = *libc.As[int32](lookahead)
	cmp775 = v323 == 46
	if cmp775 {
		goto if_then798
	} else {
		goto lor_lhs_false777
	}

lor_lhs_false777:
	v324 = *libc.As[int32](lookahead)
	cmp778 = 48 <= v324
	if cmp778 {
		goto land_lhs_true780
	} else {
		goto lor_lhs_false783
	}

land_lhs_true780:
	v325 = *libc.As[int32](lookahead)
	cmp781 = v325 <= 57
	if cmp781 {
		goto if_then798
	} else {
		goto lor_lhs_false783
	}

lor_lhs_false783:
	v326 = *libc.As[int32](lookahead)
	cmp784 = 65 <= v326
	if cmp784 {
		goto land_lhs_true786
	} else {
		goto lor_lhs_false789
	}

land_lhs_true786:
	v327 = *libc.As[int32](lookahead)
	cmp787 = v327 <= 90
	if cmp787 {
		goto if_then798
	} else {
		goto lor_lhs_false789
	}

lor_lhs_false789:
	v328 = *libc.As[int32](lookahead)
	cmp790 = v328 == 95
	if cmp790 {
		goto if_then798
	} else {
		goto lor_lhs_false792
	}

lor_lhs_false792:
	v329 = *libc.As[int32](lookahead)
	cmp793 = 97 <= v329
	if cmp793 {
		goto land_lhs_true795
	} else {
		goto if_end799
	}

land_lhs_true795:
	v330 = *libc.As[int32](lookahead)
	cmp796 = v330 <= 122
	if cmp796 {
		goto if_then798
	} else {
		goto if_end799
	}

if_then798:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end799:
	v331 = *libc.As[byte](result)
	loadedv800 = (v331 & 1) != 0
	*libc.As[bool](retval) = loadedv800
	goto _return

sw_bb801:
	*libc.As[byte](result) = 1
	v332 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol802 = libc.Ptr(&libc.As[TSLexer](v332).F1)
	*libc.As[int16](result_symbol802) = 10
	v333 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end803 = libc.Ptr(&libc.As[TSLexer](v333).F3)
	v334 = *libc.As[unsafe.Pointer](mark_end803)
	v335 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v334)(v335)
	v336 = *libc.As[int32](lookahead)
	cmp804 = v336 == 100
	if cmp804 {
		goto if_then806
	} else {
		goto if_end807
	}

if_then806:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end807:
	v337 = *libc.As[int32](lookahead)
	cmp808 = v337 == 45
	if cmp808 {
		goto if_then834
	} else {
		goto lor_lhs_false810
	}

lor_lhs_false810:
	v338 = *libc.As[int32](lookahead)
	cmp811 = v338 == 46
	if cmp811 {
		goto if_then834
	} else {
		goto lor_lhs_false813
	}

lor_lhs_false813:
	v339 = *libc.As[int32](lookahead)
	cmp814 = 48 <= v339
	if cmp814 {
		goto land_lhs_true816
	} else {
		goto lor_lhs_false819
	}

land_lhs_true816:
	v340 = *libc.As[int32](lookahead)
	cmp817 = v340 <= 57
	if cmp817 {
		goto if_then834
	} else {
		goto lor_lhs_false819
	}

lor_lhs_false819:
	v341 = *libc.As[int32](lookahead)
	cmp820 = 65 <= v341
	if cmp820 {
		goto land_lhs_true822
	} else {
		goto lor_lhs_false825
	}

land_lhs_true822:
	v342 = *libc.As[int32](lookahead)
	cmp823 = v342 <= 90
	if cmp823 {
		goto if_then834
	} else {
		goto lor_lhs_false825
	}

lor_lhs_false825:
	v343 = *libc.As[int32](lookahead)
	cmp826 = v343 == 95
	if cmp826 {
		goto if_then834
	} else {
		goto lor_lhs_false828
	}

lor_lhs_false828:
	v344 = *libc.As[int32](lookahead)
	cmp829 = 97 <= v344
	if cmp829 {
		goto land_lhs_true831
	} else {
		goto if_end835
	}

land_lhs_true831:
	v345 = *libc.As[int32](lookahead)
	cmp832 = v345 <= 122
	if cmp832 {
		goto if_then834
	} else {
		goto if_end835
	}

if_then834:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end835:
	v346 = *libc.As[byte](result)
	loadedv836 = (v346 & 1) != 0
	*libc.As[bool](retval) = loadedv836
	goto _return

sw_bb837:
	*libc.As[byte](result) = 1
	v347 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol838 = libc.Ptr(&libc.As[TSLexer](v347).F1)
	*libc.As[int16](result_symbol838) = 10
	v348 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end839 = libc.Ptr(&libc.As[TSLexer](v348).F3)
	v349 = *libc.As[unsafe.Pointer](mark_end839)
	v350 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v349)(v350)
	v351 = *libc.As[int32](lookahead)
	cmp840 = v351 == 100
	if cmp840 {
		goto if_then842
	} else {
		goto if_end843
	}

if_then842:
	*libc.As[int16](state_addr) = 121
	goto next_state

if_end843:
	v352 = *libc.As[int32](lookahead)
	cmp844 = v352 == 45
	if cmp844 {
		goto if_then870
	} else {
		goto lor_lhs_false846
	}

lor_lhs_false846:
	v353 = *libc.As[int32](lookahead)
	cmp847 = v353 == 46
	if cmp847 {
		goto if_then870
	} else {
		goto lor_lhs_false849
	}

lor_lhs_false849:
	v354 = *libc.As[int32](lookahead)
	cmp850 = 48 <= v354
	if cmp850 {
		goto land_lhs_true852
	} else {
		goto lor_lhs_false855
	}

land_lhs_true852:
	v355 = *libc.As[int32](lookahead)
	cmp853 = v355 <= 57
	if cmp853 {
		goto if_then870
	} else {
		goto lor_lhs_false855
	}

lor_lhs_false855:
	v356 = *libc.As[int32](lookahead)
	cmp856 = 65 <= v356
	if cmp856 {
		goto land_lhs_true858
	} else {
		goto lor_lhs_false861
	}

land_lhs_true858:
	v357 = *libc.As[int32](lookahead)
	cmp859 = v357 <= 90
	if cmp859 {
		goto if_then870
	} else {
		goto lor_lhs_false861
	}

lor_lhs_false861:
	v358 = *libc.As[int32](lookahead)
	cmp862 = v358 == 95
	if cmp862 {
		goto if_then870
	} else {
		goto lor_lhs_false864
	}

lor_lhs_false864:
	v359 = *libc.As[int32](lookahead)
	cmp865 = 97 <= v359
	if cmp865 {
		goto land_lhs_true867
	} else {
		goto if_end871
	}

land_lhs_true867:
	v360 = *libc.As[int32](lookahead)
	cmp868 = v360 <= 122
	if cmp868 {
		goto if_then870
	} else {
		goto if_end871
	}

if_then870:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end871:
	v361 = *libc.As[byte](result)
	loadedv872 = (v361 & 1) != 0
	*libc.As[bool](retval) = loadedv872
	goto _return

sw_bb873:
	*libc.As[byte](result) = 1
	v362 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol874 = libc.Ptr(&libc.As[TSLexer](v362).F1)
	*libc.As[int16](result_symbol874) = 10
	v363 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end875 = libc.Ptr(&libc.As[TSLexer](v363).F3)
	v364 = *libc.As[unsafe.Pointer](mark_end875)
	v365 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v364)(v365)
	v366 = *libc.As[int32](lookahead)
	cmp876 = v366 == 100
	if cmp876 {
		goto if_then878
	} else {
		goto if_end879
	}

if_then878:
	*libc.As[int16](state_addr) = 115
	goto next_state

if_end879:
	v367 = *libc.As[int32](lookahead)
	cmp880 = v367 == 45
	if cmp880 {
		goto if_then906
	} else {
		goto lor_lhs_false882
	}

lor_lhs_false882:
	v368 = *libc.As[int32](lookahead)
	cmp883 = v368 == 46
	if cmp883 {
		goto if_then906
	} else {
		goto lor_lhs_false885
	}

lor_lhs_false885:
	v369 = *libc.As[int32](lookahead)
	cmp886 = 48 <= v369
	if cmp886 {
		goto land_lhs_true888
	} else {
		goto lor_lhs_false891
	}

land_lhs_true888:
	v370 = *libc.As[int32](lookahead)
	cmp889 = v370 <= 57
	if cmp889 {
		goto if_then906
	} else {
		goto lor_lhs_false891
	}

lor_lhs_false891:
	v371 = *libc.As[int32](lookahead)
	cmp892 = 65 <= v371
	if cmp892 {
		goto land_lhs_true894
	} else {
		goto lor_lhs_false897
	}

land_lhs_true894:
	v372 = *libc.As[int32](lookahead)
	cmp895 = v372 <= 90
	if cmp895 {
		goto if_then906
	} else {
		goto lor_lhs_false897
	}

lor_lhs_false897:
	v373 = *libc.As[int32](lookahead)
	cmp898 = v373 == 95
	if cmp898 {
		goto if_then906
	} else {
		goto lor_lhs_false900
	}

lor_lhs_false900:
	v374 = *libc.As[int32](lookahead)
	cmp901 = 97 <= v374
	if cmp901 {
		goto land_lhs_true903
	} else {
		goto if_end907
	}

land_lhs_true903:
	v375 = *libc.As[int32](lookahead)
	cmp904 = v375 <= 122
	if cmp904 {
		goto if_then906
	} else {
		goto if_end907
	}

if_then906:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end907:
	v376 = *libc.As[byte](result)
	loadedv908 = (v376 & 1) != 0
	*libc.As[bool](retval) = loadedv908
	goto _return

sw_bb909:
	*libc.As[byte](result) = 1
	v377 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol910 = libc.Ptr(&libc.As[TSLexer](v377).F1)
	*libc.As[int16](result_symbol910) = 10
	v378 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end911 = libc.Ptr(&libc.As[TSLexer](v378).F3)
	v379 = *libc.As[unsafe.Pointer](mark_end911)
	v380 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v379)(v380)
	v381 = *libc.As[int32](lookahead)
	cmp912 = v381 == 101
	if cmp912 {
		goto if_then914
	} else {
		goto if_end915
	}

if_then914:
	*libc.As[int16](state_addr) = 125
	goto next_state

if_end915:
	v382 = *libc.As[int32](lookahead)
	cmp916 = v382 == 45
	if cmp916 {
		goto if_then942
	} else {
		goto lor_lhs_false918
	}

lor_lhs_false918:
	v383 = *libc.As[int32](lookahead)
	cmp919 = v383 == 46
	if cmp919 {
		goto if_then942
	} else {
		goto lor_lhs_false921
	}

lor_lhs_false921:
	v384 = *libc.As[int32](lookahead)
	cmp922 = 48 <= v384
	if cmp922 {
		goto land_lhs_true924
	} else {
		goto lor_lhs_false927
	}

land_lhs_true924:
	v385 = *libc.As[int32](lookahead)
	cmp925 = v385 <= 57
	if cmp925 {
		goto if_then942
	} else {
		goto lor_lhs_false927
	}

lor_lhs_false927:
	v386 = *libc.As[int32](lookahead)
	cmp928 = 65 <= v386
	if cmp928 {
		goto land_lhs_true930
	} else {
		goto lor_lhs_false933
	}

land_lhs_true930:
	v387 = *libc.As[int32](lookahead)
	cmp931 = v387 <= 90
	if cmp931 {
		goto if_then942
	} else {
		goto lor_lhs_false933
	}

lor_lhs_false933:
	v388 = *libc.As[int32](lookahead)
	cmp934 = v388 == 95
	if cmp934 {
		goto if_then942
	} else {
		goto lor_lhs_false936
	}

lor_lhs_false936:
	v389 = *libc.As[int32](lookahead)
	cmp937 = 97 <= v389
	if cmp937 {
		goto land_lhs_true939
	} else {
		goto if_end943
	}

land_lhs_true939:
	v390 = *libc.As[int32](lookahead)
	cmp940 = v390 <= 122
	if cmp940 {
		goto if_then942
	} else {
		goto if_end943
	}

if_then942:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end943:
	v391 = *libc.As[byte](result)
	loadedv944 = (v391 & 1) != 0
	*libc.As[bool](retval) = loadedv944
	goto _return

sw_bb945:
	*libc.As[byte](result) = 1
	v392 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol946 = libc.Ptr(&libc.As[TSLexer](v392).F1)
	*libc.As[int16](result_symbol946) = 10
	v393 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end947 = libc.Ptr(&libc.As[TSLexer](v393).F3)
	v394 = *libc.As[unsafe.Pointer](mark_end947)
	v395 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v394)(v395)
	v396 = *libc.As[int32](lookahead)
	cmp948 = v396 == 101
	if cmp948 {
		goto if_then950
	} else {
		goto if_end951
	}

if_then950:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end951:
	v397 = *libc.As[int32](lookahead)
	cmp952 = v397 == 45
	if cmp952 {
		goto if_then978
	} else {
		goto lor_lhs_false954
	}

lor_lhs_false954:
	v398 = *libc.As[int32](lookahead)
	cmp955 = v398 == 46
	if cmp955 {
		goto if_then978
	} else {
		goto lor_lhs_false957
	}

lor_lhs_false957:
	v399 = *libc.As[int32](lookahead)
	cmp958 = 48 <= v399
	if cmp958 {
		goto land_lhs_true960
	} else {
		goto lor_lhs_false963
	}

land_lhs_true960:
	v400 = *libc.As[int32](lookahead)
	cmp961 = v400 <= 57
	if cmp961 {
		goto if_then978
	} else {
		goto lor_lhs_false963
	}

lor_lhs_false963:
	v401 = *libc.As[int32](lookahead)
	cmp964 = 65 <= v401
	if cmp964 {
		goto land_lhs_true966
	} else {
		goto lor_lhs_false969
	}

land_lhs_true966:
	v402 = *libc.As[int32](lookahead)
	cmp967 = v402 <= 90
	if cmp967 {
		goto if_then978
	} else {
		goto lor_lhs_false969
	}

lor_lhs_false969:
	v403 = *libc.As[int32](lookahead)
	cmp970 = v403 == 95
	if cmp970 {
		goto if_then978
	} else {
		goto lor_lhs_false972
	}

lor_lhs_false972:
	v404 = *libc.As[int32](lookahead)
	cmp973 = 97 <= v404
	if cmp973 {
		goto land_lhs_true975
	} else {
		goto if_end979
	}

land_lhs_true975:
	v405 = *libc.As[int32](lookahead)
	cmp976 = v405 <= 122
	if cmp976 {
		goto if_then978
	} else {
		goto if_end979
	}

if_then978:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end979:
	v406 = *libc.As[byte](result)
	loadedv980 = (v406 & 1) != 0
	*libc.As[bool](retval) = loadedv980
	goto _return

sw_bb981:
	*libc.As[byte](result) = 1
	v407 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol982 = libc.Ptr(&libc.As[TSLexer](v407).F1)
	*libc.As[int16](result_symbol982) = 10
	v408 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end983 = libc.Ptr(&libc.As[TSLexer](v408).F3)
	v409 = *libc.As[unsafe.Pointer](mark_end983)
	v410 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v409)(v410)
	v411 = *libc.As[int32](lookahead)
	cmp984 = v411 == 101
	if cmp984 {
		goto if_then986
	} else {
		goto if_end987
	}

if_then986:
	*libc.As[int16](state_addr) = 110
	goto next_state

if_end987:
	v412 = *libc.As[int32](lookahead)
	cmp988 = v412 == 45
	if cmp988 {
		goto if_then1014
	} else {
		goto lor_lhs_false990
	}

lor_lhs_false990:
	v413 = *libc.As[int32](lookahead)
	cmp991 = v413 == 46
	if cmp991 {
		goto if_then1014
	} else {
		goto lor_lhs_false993
	}

lor_lhs_false993:
	v414 = *libc.As[int32](lookahead)
	cmp994 = 48 <= v414
	if cmp994 {
		goto land_lhs_true996
	} else {
		goto lor_lhs_false999
	}

land_lhs_true996:
	v415 = *libc.As[int32](lookahead)
	cmp997 = v415 <= 57
	if cmp997 {
		goto if_then1014
	} else {
		goto lor_lhs_false999
	}

lor_lhs_false999:
	v416 = *libc.As[int32](lookahead)
	cmp1000 = 65 <= v416
	if cmp1000 {
		goto land_lhs_true1002
	} else {
		goto lor_lhs_false1005
	}

land_lhs_true1002:
	v417 = *libc.As[int32](lookahead)
	cmp1003 = v417 <= 90
	if cmp1003 {
		goto if_then1014
	} else {
		goto lor_lhs_false1005
	}

lor_lhs_false1005:
	v418 = *libc.As[int32](lookahead)
	cmp1006 = v418 == 95
	if cmp1006 {
		goto if_then1014
	} else {
		goto lor_lhs_false1008
	}

lor_lhs_false1008:
	v419 = *libc.As[int32](lookahead)
	cmp1009 = 97 <= v419
	if cmp1009 {
		goto land_lhs_true1011
	} else {
		goto if_end1015
	}

land_lhs_true1011:
	v420 = *libc.As[int32](lookahead)
	cmp1012 = v420 <= 122
	if cmp1012 {
		goto if_then1014
	} else {
		goto if_end1015
	}

if_then1014:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1015:
	v421 = *libc.As[byte](result)
	loadedv1016 = (v421 & 1) != 0
	*libc.As[bool](retval) = loadedv1016
	goto _return

sw_bb1017:
	*libc.As[byte](result) = 1
	v422 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1018 = libc.Ptr(&libc.As[TSLexer](v422).F1)
	*libc.As[int16](result_symbol1018) = 10
	v423 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1019 = libc.Ptr(&libc.As[TSLexer](v423).F3)
	v424 = *libc.As[unsafe.Pointer](mark_end1019)
	v425 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v424)(v425)
	v426 = *libc.As[int32](lookahead)
	cmp1020 = v426 == 101
	if cmp1020 {
		goto if_then1022
	} else {
		goto if_end1023
	}

if_then1022:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1023:
	v427 = *libc.As[int32](lookahead)
	cmp1024 = v427 == 45
	if cmp1024 {
		goto if_then1050
	} else {
		goto lor_lhs_false1026
	}

lor_lhs_false1026:
	v428 = *libc.As[int32](lookahead)
	cmp1027 = v428 == 46
	if cmp1027 {
		goto if_then1050
	} else {
		goto lor_lhs_false1029
	}

lor_lhs_false1029:
	v429 = *libc.As[int32](lookahead)
	cmp1030 = 48 <= v429
	if cmp1030 {
		goto land_lhs_true1032
	} else {
		goto lor_lhs_false1035
	}

land_lhs_true1032:
	v430 = *libc.As[int32](lookahead)
	cmp1033 = v430 <= 57
	if cmp1033 {
		goto if_then1050
	} else {
		goto lor_lhs_false1035
	}

lor_lhs_false1035:
	v431 = *libc.As[int32](lookahead)
	cmp1036 = 65 <= v431
	if cmp1036 {
		goto land_lhs_true1038
	} else {
		goto lor_lhs_false1041
	}

land_lhs_true1038:
	v432 = *libc.As[int32](lookahead)
	cmp1039 = v432 <= 90
	if cmp1039 {
		goto if_then1050
	} else {
		goto lor_lhs_false1041
	}

lor_lhs_false1041:
	v433 = *libc.As[int32](lookahead)
	cmp1042 = v433 == 95
	if cmp1042 {
		goto if_then1050
	} else {
		goto lor_lhs_false1044
	}

lor_lhs_false1044:
	v434 = *libc.As[int32](lookahead)
	cmp1045 = 97 <= v434
	if cmp1045 {
		goto land_lhs_true1047
	} else {
		goto if_end1051
	}

land_lhs_true1047:
	v435 = *libc.As[int32](lookahead)
	cmp1048 = v435 <= 122
	if cmp1048 {
		goto if_then1050
	} else {
		goto if_end1051
	}

if_then1050:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1051:
	v436 = *libc.As[byte](result)
	loadedv1052 = (v436 & 1) != 0
	*libc.As[bool](retval) = loadedv1052
	goto _return

sw_bb1053:
	*libc.As[byte](result) = 1
	v437 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1054 = libc.Ptr(&libc.As[TSLexer](v437).F1)
	*libc.As[int16](result_symbol1054) = 10
	v438 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1055 = libc.Ptr(&libc.As[TSLexer](v438).F3)
	v439 = *libc.As[unsafe.Pointer](mark_end1055)
	v440 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v439)(v440)
	v441 = *libc.As[int32](lookahead)
	cmp1056 = v441 == 102
	if cmp1056 {
		goto if_then1058
	} else {
		goto if_end1059
	}

if_then1058:
	*libc.As[int16](state_addr) = 106
	goto next_state

if_end1059:
	v442 = *libc.As[int32](lookahead)
	cmp1060 = v442 == 45
	if cmp1060 {
		goto if_then1086
	} else {
		goto lor_lhs_false1062
	}

lor_lhs_false1062:
	v443 = *libc.As[int32](lookahead)
	cmp1063 = v443 == 46
	if cmp1063 {
		goto if_then1086
	} else {
		goto lor_lhs_false1065
	}

lor_lhs_false1065:
	v444 = *libc.As[int32](lookahead)
	cmp1066 = 48 <= v444
	if cmp1066 {
		goto land_lhs_true1068
	} else {
		goto lor_lhs_false1071
	}

land_lhs_true1068:
	v445 = *libc.As[int32](lookahead)
	cmp1069 = v445 <= 57
	if cmp1069 {
		goto if_then1086
	} else {
		goto lor_lhs_false1071
	}

lor_lhs_false1071:
	v446 = *libc.As[int32](lookahead)
	cmp1072 = 65 <= v446
	if cmp1072 {
		goto land_lhs_true1074
	} else {
		goto lor_lhs_false1077
	}

land_lhs_true1074:
	v447 = *libc.As[int32](lookahead)
	cmp1075 = v447 <= 90
	if cmp1075 {
		goto if_then1086
	} else {
		goto lor_lhs_false1077
	}

lor_lhs_false1077:
	v448 = *libc.As[int32](lookahead)
	cmp1078 = v448 == 95
	if cmp1078 {
		goto if_then1086
	} else {
		goto lor_lhs_false1080
	}

lor_lhs_false1080:
	v449 = *libc.As[int32](lookahead)
	cmp1081 = 97 <= v449
	if cmp1081 {
		goto land_lhs_true1083
	} else {
		goto if_end1087
	}

land_lhs_true1083:
	v450 = *libc.As[int32](lookahead)
	cmp1084 = v450 <= 122
	if cmp1084 {
		goto if_then1086
	} else {
		goto if_end1087
	}

if_then1086:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1087:
	v451 = *libc.As[byte](result)
	loadedv1088 = (v451 & 1) != 0
	*libc.As[bool](retval) = loadedv1088
	goto _return

sw_bb1089:
	*libc.As[byte](result) = 1
	v452 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1090 = libc.Ptr(&libc.As[TSLexer](v452).F1)
	*libc.As[int16](result_symbol1090) = 10
	v453 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1091 = libc.Ptr(&libc.As[TSLexer](v453).F3)
	v454 = *libc.As[unsafe.Pointer](mark_end1091)
	v455 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v454)(v455)
	v456 = *libc.As[int32](lookahead)
	cmp1092 = v456 == 103
	if cmp1092 {
		goto if_then1094
	} else {
		goto if_end1095
	}

if_then1094:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1095:
	v457 = *libc.As[int32](lookahead)
	cmp1096 = v457 == 45
	if cmp1096 {
		goto if_then1122
	} else {
		goto lor_lhs_false1098
	}

lor_lhs_false1098:
	v458 = *libc.As[int32](lookahead)
	cmp1099 = v458 == 46
	if cmp1099 {
		goto if_then1122
	} else {
		goto lor_lhs_false1101
	}

lor_lhs_false1101:
	v459 = *libc.As[int32](lookahead)
	cmp1102 = 48 <= v459
	if cmp1102 {
		goto land_lhs_true1104
	} else {
		goto lor_lhs_false1107
	}

land_lhs_true1104:
	v460 = *libc.As[int32](lookahead)
	cmp1105 = v460 <= 57
	if cmp1105 {
		goto if_then1122
	} else {
		goto lor_lhs_false1107
	}

lor_lhs_false1107:
	v461 = *libc.As[int32](lookahead)
	cmp1108 = 65 <= v461
	if cmp1108 {
		goto land_lhs_true1110
	} else {
		goto lor_lhs_false1113
	}

land_lhs_true1110:
	v462 = *libc.As[int32](lookahead)
	cmp1111 = v462 <= 90
	if cmp1111 {
		goto if_then1122
	} else {
		goto lor_lhs_false1113
	}

lor_lhs_false1113:
	v463 = *libc.As[int32](lookahead)
	cmp1114 = v463 == 95
	if cmp1114 {
		goto if_then1122
	} else {
		goto lor_lhs_false1116
	}

lor_lhs_false1116:
	v464 = *libc.As[int32](lookahead)
	cmp1117 = 97 <= v464
	if cmp1117 {
		goto land_lhs_true1119
	} else {
		goto if_end1123
	}

land_lhs_true1119:
	v465 = *libc.As[int32](lookahead)
	cmp1120 = v465 <= 122
	if cmp1120 {
		goto if_then1122
	} else {
		goto if_end1123
	}

if_then1122:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1123:
	v466 = *libc.As[byte](result)
	loadedv1124 = (v466 & 1) != 0
	*libc.As[bool](retval) = loadedv1124
	goto _return

sw_bb1125:
	*libc.As[byte](result) = 1
	v467 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1126 = libc.Ptr(&libc.As[TSLexer](v467).F1)
	*libc.As[int16](result_symbol1126) = 10
	v468 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1127 = libc.Ptr(&libc.As[TSLexer](v468).F3)
	v469 = *libc.As[unsafe.Pointer](mark_end1127)
	v470 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v469)(v470)
	v471 = *libc.As[int32](lookahead)
	cmp1128 = v471 == 105
	if cmp1128 {
		goto if_then1130
	} else {
		goto if_end1131
	}

if_then1130:
	*libc.As[int16](state_addr) = 123
	goto next_state

if_end1131:
	v472 = *libc.As[int32](lookahead)
	cmp1132 = v472 == 45
	if cmp1132 {
		goto if_then1158
	} else {
		goto lor_lhs_false1134
	}

lor_lhs_false1134:
	v473 = *libc.As[int32](lookahead)
	cmp1135 = v473 == 46
	if cmp1135 {
		goto if_then1158
	} else {
		goto lor_lhs_false1137
	}

lor_lhs_false1137:
	v474 = *libc.As[int32](lookahead)
	cmp1138 = 48 <= v474
	if cmp1138 {
		goto land_lhs_true1140
	} else {
		goto lor_lhs_false1143
	}

land_lhs_true1140:
	v475 = *libc.As[int32](lookahead)
	cmp1141 = v475 <= 57
	if cmp1141 {
		goto if_then1158
	} else {
		goto lor_lhs_false1143
	}

lor_lhs_false1143:
	v476 = *libc.As[int32](lookahead)
	cmp1144 = 65 <= v476
	if cmp1144 {
		goto land_lhs_true1146
	} else {
		goto lor_lhs_false1149
	}

land_lhs_true1146:
	v477 = *libc.As[int32](lookahead)
	cmp1147 = v477 <= 90
	if cmp1147 {
		goto if_then1158
	} else {
		goto lor_lhs_false1149
	}

lor_lhs_false1149:
	v478 = *libc.As[int32](lookahead)
	cmp1150 = v478 == 95
	if cmp1150 {
		goto if_then1158
	} else {
		goto lor_lhs_false1152
	}

lor_lhs_false1152:
	v479 = *libc.As[int32](lookahead)
	cmp1153 = 97 <= v479
	if cmp1153 {
		goto land_lhs_true1155
	} else {
		goto if_end1159
	}

land_lhs_true1155:
	v480 = *libc.As[int32](lookahead)
	cmp1156 = v480 <= 122
	if cmp1156 {
		goto if_then1158
	} else {
		goto if_end1159
	}

if_then1158:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1159:
	v481 = *libc.As[byte](result)
	loadedv1160 = (v481 & 1) != 0
	*libc.As[bool](retval) = loadedv1160
	goto _return

sw_bb1161:
	*libc.As[byte](result) = 1
	v482 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1162 = libc.Ptr(&libc.As[TSLexer](v482).F1)
	*libc.As[int16](result_symbol1162) = 10
	v483 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1163 = libc.Ptr(&libc.As[TSLexer](v483).F3)
	v484 = *libc.As[unsafe.Pointer](mark_end1163)
	v485 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v484)(v485)
	v486 = *libc.As[int32](lookahead)
	cmp1164 = v486 == 105
	if cmp1164 {
		goto if_then1166
	} else {
		goto if_end1167
	}

if_then1166:
	*libc.As[int16](state_addr) = 109
	goto next_state

if_end1167:
	v487 = *libc.As[int32](lookahead)
	cmp1168 = v487 == 45
	if cmp1168 {
		goto if_then1194
	} else {
		goto lor_lhs_false1170
	}

lor_lhs_false1170:
	v488 = *libc.As[int32](lookahead)
	cmp1171 = v488 == 46
	if cmp1171 {
		goto if_then1194
	} else {
		goto lor_lhs_false1173
	}

lor_lhs_false1173:
	v489 = *libc.As[int32](lookahead)
	cmp1174 = 48 <= v489
	if cmp1174 {
		goto land_lhs_true1176
	} else {
		goto lor_lhs_false1179
	}

land_lhs_true1176:
	v490 = *libc.As[int32](lookahead)
	cmp1177 = v490 <= 57
	if cmp1177 {
		goto if_then1194
	} else {
		goto lor_lhs_false1179
	}

lor_lhs_false1179:
	v491 = *libc.As[int32](lookahead)
	cmp1180 = 65 <= v491
	if cmp1180 {
		goto land_lhs_true1182
	} else {
		goto lor_lhs_false1185
	}

land_lhs_true1182:
	v492 = *libc.As[int32](lookahead)
	cmp1183 = v492 <= 90
	if cmp1183 {
		goto if_then1194
	} else {
		goto lor_lhs_false1185
	}

lor_lhs_false1185:
	v493 = *libc.As[int32](lookahead)
	cmp1186 = v493 == 95
	if cmp1186 {
		goto if_then1194
	} else {
		goto lor_lhs_false1188
	}

lor_lhs_false1188:
	v494 = *libc.As[int32](lookahead)
	cmp1189 = 97 <= v494
	if cmp1189 {
		goto land_lhs_true1191
	} else {
		goto if_end1195
	}

land_lhs_true1191:
	v495 = *libc.As[int32](lookahead)
	cmp1192 = v495 <= 122
	if cmp1192 {
		goto if_then1194
	} else {
		goto if_end1195
	}

if_then1194:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1195:
	v496 = *libc.As[byte](result)
	loadedv1196 = (v496 & 1) != 0
	*libc.As[bool](retval) = loadedv1196
	goto _return

sw_bb1197:
	*libc.As[byte](result) = 1
	v497 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1198 = libc.Ptr(&libc.As[TSLexer](v497).F1)
	*libc.As[int16](result_symbol1198) = 10
	v498 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1199 = libc.Ptr(&libc.As[TSLexer](v498).F3)
	v499 = *libc.As[unsafe.Pointer](mark_end1199)
	v500 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v499)(v500)
	v501 = *libc.As[int32](lookahead)
	cmp1200 = v501 == 105
	if cmp1200 {
		goto if_then1202
	} else {
		goto if_end1203
	}

if_then1202:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end1203:
	v502 = *libc.As[int32](lookahead)
	cmp1204 = v502 == 45
	if cmp1204 {
		goto if_then1230
	} else {
		goto lor_lhs_false1206
	}

lor_lhs_false1206:
	v503 = *libc.As[int32](lookahead)
	cmp1207 = v503 == 46
	if cmp1207 {
		goto if_then1230
	} else {
		goto lor_lhs_false1209
	}

lor_lhs_false1209:
	v504 = *libc.As[int32](lookahead)
	cmp1210 = 48 <= v504
	if cmp1210 {
		goto land_lhs_true1212
	} else {
		goto lor_lhs_false1215
	}

land_lhs_true1212:
	v505 = *libc.As[int32](lookahead)
	cmp1213 = v505 <= 57
	if cmp1213 {
		goto if_then1230
	} else {
		goto lor_lhs_false1215
	}

lor_lhs_false1215:
	v506 = *libc.As[int32](lookahead)
	cmp1216 = 65 <= v506
	if cmp1216 {
		goto land_lhs_true1218
	} else {
		goto lor_lhs_false1221
	}

land_lhs_true1218:
	v507 = *libc.As[int32](lookahead)
	cmp1219 = v507 <= 90
	if cmp1219 {
		goto if_then1230
	} else {
		goto lor_lhs_false1221
	}

lor_lhs_false1221:
	v508 = *libc.As[int32](lookahead)
	cmp1222 = v508 == 95
	if cmp1222 {
		goto if_then1230
	} else {
		goto lor_lhs_false1224
	}

lor_lhs_false1224:
	v509 = *libc.As[int32](lookahead)
	cmp1225 = 97 <= v509
	if cmp1225 {
		goto land_lhs_true1227
	} else {
		goto if_end1231
	}

land_lhs_true1227:
	v510 = *libc.As[int32](lookahead)
	cmp1228 = v510 <= 122
	if cmp1228 {
		goto if_then1230
	} else {
		goto if_end1231
	}

if_then1230:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1231:
	v511 = *libc.As[byte](result)
	loadedv1232 = (v511 & 1) != 0
	*libc.As[bool](retval) = loadedv1232
	goto _return

sw_bb1233:
	*libc.As[byte](result) = 1
	v512 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1234 = libc.Ptr(&libc.As[TSLexer](v512).F1)
	*libc.As[int16](result_symbol1234) = 10
	v513 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1235 = libc.Ptr(&libc.As[TSLexer](v513).F3)
	v514 = *libc.As[unsafe.Pointer](mark_end1235)
	v515 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v514)(v515)
	v516 = *libc.As[int32](lookahead)
	cmp1236 = v516 == 105
	if cmp1236 {
		goto if_then1238
	} else {
		goto if_end1239
	}

if_then1238:
	*libc.As[int16](state_addr) = 112
	goto next_state

if_end1239:
	v517 = *libc.As[int32](lookahead)
	cmp1240 = v517 == 45
	if cmp1240 {
		goto if_then1266
	} else {
		goto lor_lhs_false1242
	}

lor_lhs_false1242:
	v518 = *libc.As[int32](lookahead)
	cmp1243 = v518 == 46
	if cmp1243 {
		goto if_then1266
	} else {
		goto lor_lhs_false1245
	}

lor_lhs_false1245:
	v519 = *libc.As[int32](lookahead)
	cmp1246 = 48 <= v519
	if cmp1246 {
		goto land_lhs_true1248
	} else {
		goto lor_lhs_false1251
	}

land_lhs_true1248:
	v520 = *libc.As[int32](lookahead)
	cmp1249 = v520 <= 57
	if cmp1249 {
		goto if_then1266
	} else {
		goto lor_lhs_false1251
	}

lor_lhs_false1251:
	v521 = *libc.As[int32](lookahead)
	cmp1252 = 65 <= v521
	if cmp1252 {
		goto land_lhs_true1254
	} else {
		goto lor_lhs_false1257
	}

land_lhs_true1254:
	v522 = *libc.As[int32](lookahead)
	cmp1255 = v522 <= 90
	if cmp1255 {
		goto if_then1266
	} else {
		goto lor_lhs_false1257
	}

lor_lhs_false1257:
	v523 = *libc.As[int32](lookahead)
	cmp1258 = v523 == 95
	if cmp1258 {
		goto if_then1266
	} else {
		goto lor_lhs_false1260
	}

lor_lhs_false1260:
	v524 = *libc.As[int32](lookahead)
	cmp1261 = 97 <= v524
	if cmp1261 {
		goto land_lhs_true1263
	} else {
		goto if_end1267
	}

land_lhs_true1263:
	v525 = *libc.As[int32](lookahead)
	cmp1264 = v525 <= 122
	if cmp1264 {
		goto if_then1266
	} else {
		goto if_end1267
	}

if_then1266:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1267:
	v526 = *libc.As[byte](result)
	loadedv1268 = (v526 & 1) != 0
	*libc.As[bool](retval) = loadedv1268
	goto _return

sw_bb1269:
	*libc.As[byte](result) = 1
	v527 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1270 = libc.Ptr(&libc.As[TSLexer](v527).F1)
	*libc.As[int16](result_symbol1270) = 10
	v528 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1271 = libc.Ptr(&libc.As[TSLexer](v528).F3)
	v529 = *libc.As[unsafe.Pointer](mark_end1271)
	v530 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v529)(v530)
	v531 = *libc.As[int32](lookahead)
	cmp1272 = v531 == 108
	if cmp1272 {
		goto if_then1274
	} else {
		goto if_end1275
	}

if_then1274:
	*libc.As[int16](state_addr) = 129
	goto next_state

if_end1275:
	v532 = *libc.As[int32](lookahead)
	cmp1276 = v532 == 45
	if cmp1276 {
		goto if_then1302
	} else {
		goto lor_lhs_false1278
	}

lor_lhs_false1278:
	v533 = *libc.As[int32](lookahead)
	cmp1279 = v533 == 46
	if cmp1279 {
		goto if_then1302
	} else {
		goto lor_lhs_false1281
	}

lor_lhs_false1281:
	v534 = *libc.As[int32](lookahead)
	cmp1282 = 48 <= v534
	if cmp1282 {
		goto land_lhs_true1284
	} else {
		goto lor_lhs_false1287
	}

land_lhs_true1284:
	v535 = *libc.As[int32](lookahead)
	cmp1285 = v535 <= 57
	if cmp1285 {
		goto if_then1302
	} else {
		goto lor_lhs_false1287
	}

lor_lhs_false1287:
	v536 = *libc.As[int32](lookahead)
	cmp1288 = 65 <= v536
	if cmp1288 {
		goto land_lhs_true1290
	} else {
		goto lor_lhs_false1293
	}

land_lhs_true1290:
	v537 = *libc.As[int32](lookahead)
	cmp1291 = v537 <= 90
	if cmp1291 {
		goto if_then1302
	} else {
		goto lor_lhs_false1293
	}

lor_lhs_false1293:
	v538 = *libc.As[int32](lookahead)
	cmp1294 = v538 == 95
	if cmp1294 {
		goto if_then1302
	} else {
		goto lor_lhs_false1296
	}

lor_lhs_false1296:
	v539 = *libc.As[int32](lookahead)
	cmp1297 = 97 <= v539
	if cmp1297 {
		goto land_lhs_true1299
	} else {
		goto if_end1303
	}

land_lhs_true1299:
	v540 = *libc.As[int32](lookahead)
	cmp1300 = v540 <= 122
	if cmp1300 {
		goto if_then1302
	} else {
		goto if_end1303
	}

if_then1302:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1303:
	v541 = *libc.As[byte](result)
	loadedv1304 = (v541 & 1) != 0
	*libc.As[bool](retval) = loadedv1304
	goto _return

sw_bb1305:
	*libc.As[byte](result) = 1
	v542 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1306 = libc.Ptr(&libc.As[TSLexer](v542).F1)
	*libc.As[int16](result_symbol1306) = 10
	v543 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1307 = libc.Ptr(&libc.As[TSLexer](v543).F3)
	v544 = *libc.As[unsafe.Pointer](mark_end1307)
	v545 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v544)(v545)
	v546 = *libc.As[int32](lookahead)
	cmp1308 = v546 == 108
	if cmp1308 {
		goto if_then1310
	} else {
		goto if_end1311
	}

if_then1310:
	*libc.As[int16](state_addr) = 90
	goto next_state

if_end1311:
	v547 = *libc.As[int32](lookahead)
	cmp1312 = v547 == 45
	if cmp1312 {
		goto if_then1338
	} else {
		goto lor_lhs_false1314
	}

lor_lhs_false1314:
	v548 = *libc.As[int32](lookahead)
	cmp1315 = v548 == 46
	if cmp1315 {
		goto if_then1338
	} else {
		goto lor_lhs_false1317
	}

lor_lhs_false1317:
	v549 = *libc.As[int32](lookahead)
	cmp1318 = 48 <= v549
	if cmp1318 {
		goto land_lhs_true1320
	} else {
		goto lor_lhs_false1323
	}

land_lhs_true1320:
	v550 = *libc.As[int32](lookahead)
	cmp1321 = v550 <= 57
	if cmp1321 {
		goto if_then1338
	} else {
		goto lor_lhs_false1323
	}

lor_lhs_false1323:
	v551 = *libc.As[int32](lookahead)
	cmp1324 = 65 <= v551
	if cmp1324 {
		goto land_lhs_true1326
	} else {
		goto lor_lhs_false1329
	}

land_lhs_true1326:
	v552 = *libc.As[int32](lookahead)
	cmp1327 = v552 <= 90
	if cmp1327 {
		goto if_then1338
	} else {
		goto lor_lhs_false1329
	}

lor_lhs_false1329:
	v553 = *libc.As[int32](lookahead)
	cmp1330 = v553 == 95
	if cmp1330 {
		goto if_then1338
	} else {
		goto lor_lhs_false1332
	}

lor_lhs_false1332:
	v554 = *libc.As[int32](lookahead)
	cmp1333 = 97 <= v554
	if cmp1333 {
		goto land_lhs_true1335
	} else {
		goto if_end1339
	}

land_lhs_true1335:
	v555 = *libc.As[int32](lookahead)
	cmp1336 = v555 <= 122
	if cmp1336 {
		goto if_then1338
	} else {
		goto if_end1339
	}

if_then1338:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1339:
	v556 = *libc.As[byte](result)
	loadedv1340 = (v556 & 1) != 0
	*libc.As[bool](retval) = loadedv1340
	goto _return

sw_bb1341:
	*libc.As[byte](result) = 1
	v557 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1342 = libc.Ptr(&libc.As[TSLexer](v557).F1)
	*libc.As[int16](result_symbol1342) = 10
	v558 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1343 = libc.Ptr(&libc.As[TSLexer](v558).F3)
	v559 = *libc.As[unsafe.Pointer](mark_end1343)
	v560 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v559)(v560)
	v561 = *libc.As[int32](lookahead)
	cmp1344 = v561 == 110
	if cmp1344 {
		goto if_then1346
	} else {
		goto if_end1347
	}

if_then1346:
	*libc.As[int16](state_addr) = 94
	goto next_state

if_end1347:
	v562 = *libc.As[int32](lookahead)
	cmp1348 = v562 == 45
	if cmp1348 {
		goto if_then1374
	} else {
		goto lor_lhs_false1350
	}

lor_lhs_false1350:
	v563 = *libc.As[int32](lookahead)
	cmp1351 = v563 == 46
	if cmp1351 {
		goto if_then1374
	} else {
		goto lor_lhs_false1353
	}

lor_lhs_false1353:
	v564 = *libc.As[int32](lookahead)
	cmp1354 = 48 <= v564
	if cmp1354 {
		goto land_lhs_true1356
	} else {
		goto lor_lhs_false1359
	}

land_lhs_true1356:
	v565 = *libc.As[int32](lookahead)
	cmp1357 = v565 <= 57
	if cmp1357 {
		goto if_then1374
	} else {
		goto lor_lhs_false1359
	}

lor_lhs_false1359:
	v566 = *libc.As[int32](lookahead)
	cmp1360 = 65 <= v566
	if cmp1360 {
		goto land_lhs_true1362
	} else {
		goto lor_lhs_false1365
	}

land_lhs_true1362:
	v567 = *libc.As[int32](lookahead)
	cmp1363 = v567 <= 90
	if cmp1363 {
		goto if_then1374
	} else {
		goto lor_lhs_false1365
	}

lor_lhs_false1365:
	v568 = *libc.As[int32](lookahead)
	cmp1366 = v568 == 95
	if cmp1366 {
		goto if_then1374
	} else {
		goto lor_lhs_false1368
	}

lor_lhs_false1368:
	v569 = *libc.As[int32](lookahead)
	cmp1369 = 97 <= v569
	if cmp1369 {
		goto land_lhs_true1371
	} else {
		goto if_end1375
	}

land_lhs_true1371:
	v570 = *libc.As[int32](lookahead)
	cmp1372 = v570 <= 122
	if cmp1372 {
		goto if_then1374
	} else {
		goto if_end1375
	}

if_then1374:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1375:
	v571 = *libc.As[byte](result)
	loadedv1376 = (v571 & 1) != 0
	*libc.As[bool](retval) = loadedv1376
	goto _return

sw_bb1377:
	*libc.As[byte](result) = 1
	v572 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1378 = libc.Ptr(&libc.As[TSLexer](v572).F1)
	*libc.As[int16](result_symbol1378) = 10
	v573 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1379 = libc.Ptr(&libc.As[TSLexer](v573).F3)
	v574 = *libc.As[unsafe.Pointer](mark_end1379)
	v575 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v574)(v575)
	v576 = *libc.As[int32](lookahead)
	cmp1380 = v576 == 110
	if cmp1380 {
		goto if_then1382
	} else {
		goto if_end1383
	}

if_then1382:
	*libc.As[int16](state_addr) = 126
	goto next_state

if_end1383:
	v577 = *libc.As[int32](lookahead)
	cmp1384 = v577 == 45
	if cmp1384 {
		goto if_then1410
	} else {
		goto lor_lhs_false1386
	}

lor_lhs_false1386:
	v578 = *libc.As[int32](lookahead)
	cmp1387 = v578 == 46
	if cmp1387 {
		goto if_then1410
	} else {
		goto lor_lhs_false1389
	}

lor_lhs_false1389:
	v579 = *libc.As[int32](lookahead)
	cmp1390 = 48 <= v579
	if cmp1390 {
		goto land_lhs_true1392
	} else {
		goto lor_lhs_false1395
	}

land_lhs_true1392:
	v580 = *libc.As[int32](lookahead)
	cmp1393 = v580 <= 57
	if cmp1393 {
		goto if_then1410
	} else {
		goto lor_lhs_false1395
	}

lor_lhs_false1395:
	v581 = *libc.As[int32](lookahead)
	cmp1396 = 65 <= v581
	if cmp1396 {
		goto land_lhs_true1398
	} else {
		goto lor_lhs_false1401
	}

land_lhs_true1398:
	v582 = *libc.As[int32](lookahead)
	cmp1399 = v582 <= 90
	if cmp1399 {
		goto if_then1410
	} else {
		goto lor_lhs_false1401
	}

lor_lhs_false1401:
	v583 = *libc.As[int32](lookahead)
	cmp1402 = v583 == 95
	if cmp1402 {
		goto if_then1410
	} else {
		goto lor_lhs_false1404
	}

lor_lhs_false1404:
	v584 = *libc.As[int32](lookahead)
	cmp1405 = 97 <= v584
	if cmp1405 {
		goto land_lhs_true1407
	} else {
		goto if_end1411
	}

land_lhs_true1407:
	v585 = *libc.As[int32](lookahead)
	cmp1408 = v585 <= 122
	if cmp1408 {
		goto if_then1410
	} else {
		goto if_end1411
	}

if_then1410:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1411:
	v586 = *libc.As[byte](result)
	loadedv1412 = (v586 & 1) != 0
	*libc.As[bool](retval) = loadedv1412
	goto _return

sw_bb1413:
	*libc.As[byte](result) = 1
	v587 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1414 = libc.Ptr(&libc.As[TSLexer](v587).F1)
	*libc.As[int16](result_symbol1414) = 10
	v588 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1415 = libc.Ptr(&libc.As[TSLexer](v588).F3)
	v589 = *libc.As[unsafe.Pointer](mark_end1415)
	v590 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v589)(v590)
	v591 = *libc.As[int32](lookahead)
	cmp1416 = v591 == 110
	if cmp1416 {
		goto if_then1418
	} else {
		goto if_end1419
	}

if_then1418:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1419:
	v592 = *libc.As[int32](lookahead)
	cmp1420 = v592 == 45
	if cmp1420 {
		goto if_then1446
	} else {
		goto lor_lhs_false1422
	}

lor_lhs_false1422:
	v593 = *libc.As[int32](lookahead)
	cmp1423 = v593 == 46
	if cmp1423 {
		goto if_then1446
	} else {
		goto lor_lhs_false1425
	}

lor_lhs_false1425:
	v594 = *libc.As[int32](lookahead)
	cmp1426 = 48 <= v594
	if cmp1426 {
		goto land_lhs_true1428
	} else {
		goto lor_lhs_false1431
	}

land_lhs_true1428:
	v595 = *libc.As[int32](lookahead)
	cmp1429 = v595 <= 57
	if cmp1429 {
		goto if_then1446
	} else {
		goto lor_lhs_false1431
	}

lor_lhs_false1431:
	v596 = *libc.As[int32](lookahead)
	cmp1432 = 65 <= v596
	if cmp1432 {
		goto land_lhs_true1434
	} else {
		goto lor_lhs_false1437
	}

land_lhs_true1434:
	v597 = *libc.As[int32](lookahead)
	cmp1435 = v597 <= 90
	if cmp1435 {
		goto if_then1446
	} else {
		goto lor_lhs_false1437
	}

lor_lhs_false1437:
	v598 = *libc.As[int32](lookahead)
	cmp1438 = v598 == 95
	if cmp1438 {
		goto if_then1446
	} else {
		goto lor_lhs_false1440
	}

lor_lhs_false1440:
	v599 = *libc.As[int32](lookahead)
	cmp1441 = 97 <= v599
	if cmp1441 {
		goto land_lhs_true1443
	} else {
		goto if_end1447
	}

land_lhs_true1443:
	v600 = *libc.As[int32](lookahead)
	cmp1444 = v600 <= 122
	if cmp1444 {
		goto if_then1446
	} else {
		goto if_end1447
	}

if_then1446:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1447:
	v601 = *libc.As[byte](result)
	loadedv1448 = (v601 & 1) != 0
	*libc.As[bool](retval) = loadedv1448
	goto _return

sw_bb1449:
	*libc.As[byte](result) = 1
	v602 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1450 = libc.Ptr(&libc.As[TSLexer](v602).F1)
	*libc.As[int16](result_symbol1450) = 10
	v603 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1451 = libc.Ptr(&libc.As[TSLexer](v603).F3)
	v604 = *libc.As[unsafe.Pointer](mark_end1451)
	v605 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v604)(v605)
	v606 = *libc.As[int32](lookahead)
	cmp1452 = v606 == 110
	if cmp1452 {
		goto if_then1454
	} else {
		goto if_end1455
	}

if_then1454:
	*libc.As[int16](state_addr) = 100
	goto next_state

if_end1455:
	v607 = *libc.As[int32](lookahead)
	cmp1456 = v607 == 45
	if cmp1456 {
		goto if_then1482
	} else {
		goto lor_lhs_false1458
	}

lor_lhs_false1458:
	v608 = *libc.As[int32](lookahead)
	cmp1459 = v608 == 46
	if cmp1459 {
		goto if_then1482
	} else {
		goto lor_lhs_false1461
	}

lor_lhs_false1461:
	v609 = *libc.As[int32](lookahead)
	cmp1462 = 48 <= v609
	if cmp1462 {
		goto land_lhs_true1464
	} else {
		goto lor_lhs_false1467
	}

land_lhs_true1464:
	v610 = *libc.As[int32](lookahead)
	cmp1465 = v610 <= 57
	if cmp1465 {
		goto if_then1482
	} else {
		goto lor_lhs_false1467
	}

lor_lhs_false1467:
	v611 = *libc.As[int32](lookahead)
	cmp1468 = 65 <= v611
	if cmp1468 {
		goto land_lhs_true1470
	} else {
		goto lor_lhs_false1473
	}

land_lhs_true1470:
	v612 = *libc.As[int32](lookahead)
	cmp1471 = v612 <= 90
	if cmp1471 {
		goto if_then1482
	} else {
		goto lor_lhs_false1473
	}

lor_lhs_false1473:
	v613 = *libc.As[int32](lookahead)
	cmp1474 = v613 == 95
	if cmp1474 {
		goto if_then1482
	} else {
		goto lor_lhs_false1476
	}

lor_lhs_false1476:
	v614 = *libc.As[int32](lookahead)
	cmp1477 = 97 <= v614
	if cmp1477 {
		goto land_lhs_true1479
	} else {
		goto if_end1483
	}

land_lhs_true1479:
	v615 = *libc.As[int32](lookahead)
	cmp1480 = v615 <= 122
	if cmp1480 {
		goto if_then1482
	} else {
		goto if_end1483
	}

if_then1482:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1483:
	v616 = *libc.As[byte](result)
	loadedv1484 = (v616 & 1) != 0
	*libc.As[bool](retval) = loadedv1484
	goto _return

sw_bb1485:
	*libc.As[byte](result) = 1
	v617 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1486 = libc.Ptr(&libc.As[TSLexer](v617).F1)
	*libc.As[int16](result_symbol1486) = 10
	v618 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1487 = libc.Ptr(&libc.As[TSLexer](v618).F3)
	v619 = *libc.As[unsafe.Pointer](mark_end1487)
	v620 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v619)(v620)
	v621 = *libc.As[int32](lookahead)
	cmp1488 = v621 == 110
	if cmp1488 {
		goto if_then1490
	} else {
		goto if_end1491
	}

if_then1490:
	*libc.As[int16](state_addr) = 95
	goto next_state

if_end1491:
	v622 = *libc.As[int32](lookahead)
	cmp1492 = v622 == 45
	if cmp1492 {
		goto if_then1518
	} else {
		goto lor_lhs_false1494
	}

lor_lhs_false1494:
	v623 = *libc.As[int32](lookahead)
	cmp1495 = v623 == 46
	if cmp1495 {
		goto if_then1518
	} else {
		goto lor_lhs_false1497
	}

lor_lhs_false1497:
	v624 = *libc.As[int32](lookahead)
	cmp1498 = 48 <= v624
	if cmp1498 {
		goto land_lhs_true1500
	} else {
		goto lor_lhs_false1503
	}

land_lhs_true1500:
	v625 = *libc.As[int32](lookahead)
	cmp1501 = v625 <= 57
	if cmp1501 {
		goto if_then1518
	} else {
		goto lor_lhs_false1503
	}

lor_lhs_false1503:
	v626 = *libc.As[int32](lookahead)
	cmp1504 = 65 <= v626
	if cmp1504 {
		goto land_lhs_true1506
	} else {
		goto lor_lhs_false1509
	}

land_lhs_true1506:
	v627 = *libc.As[int32](lookahead)
	cmp1507 = v627 <= 90
	if cmp1507 {
		goto if_then1518
	} else {
		goto lor_lhs_false1509
	}

lor_lhs_false1509:
	v628 = *libc.As[int32](lookahead)
	cmp1510 = v628 == 95
	if cmp1510 {
		goto if_then1518
	} else {
		goto lor_lhs_false1512
	}

lor_lhs_false1512:
	v629 = *libc.As[int32](lookahead)
	cmp1513 = 97 <= v629
	if cmp1513 {
		goto land_lhs_true1515
	} else {
		goto if_end1519
	}

land_lhs_true1515:
	v630 = *libc.As[int32](lookahead)
	cmp1516 = v630 <= 122
	if cmp1516 {
		goto if_then1518
	} else {
		goto if_end1519
	}

if_then1518:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1519:
	v631 = *libc.As[byte](result)
	loadedv1520 = (v631 & 1) != 0
	*libc.As[bool](retval) = loadedv1520
	goto _return

sw_bb1521:
	*libc.As[byte](result) = 1
	v632 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1522 = libc.Ptr(&libc.As[TSLexer](v632).F1)
	*libc.As[int16](result_symbol1522) = 10
	v633 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1523 = libc.Ptr(&libc.As[TSLexer](v633).F3)
	v634 = *libc.As[unsafe.Pointer](mark_end1523)
	v635 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v634)(v635)
	v636 = *libc.As[int32](lookahead)
	cmp1524 = v636 == 111
	if cmp1524 {
		goto if_then1526
	} else {
		goto if_end1527
	}

if_then1526:
	*libc.As[int16](state_addr) = 123
	goto next_state

if_end1527:
	v637 = *libc.As[int32](lookahead)
	cmp1528 = v637 == 45
	if cmp1528 {
		goto if_then1554
	} else {
		goto lor_lhs_false1530
	}

lor_lhs_false1530:
	v638 = *libc.As[int32](lookahead)
	cmp1531 = v638 == 46
	if cmp1531 {
		goto if_then1554
	} else {
		goto lor_lhs_false1533
	}

lor_lhs_false1533:
	v639 = *libc.As[int32](lookahead)
	cmp1534 = 48 <= v639
	if cmp1534 {
		goto land_lhs_true1536
	} else {
		goto lor_lhs_false1539
	}

land_lhs_true1536:
	v640 = *libc.As[int32](lookahead)
	cmp1537 = v640 <= 57
	if cmp1537 {
		goto if_then1554
	} else {
		goto lor_lhs_false1539
	}

lor_lhs_false1539:
	v641 = *libc.As[int32](lookahead)
	cmp1540 = 65 <= v641
	if cmp1540 {
		goto land_lhs_true1542
	} else {
		goto lor_lhs_false1545
	}

land_lhs_true1542:
	v642 = *libc.As[int32](lookahead)
	cmp1543 = v642 <= 90
	if cmp1543 {
		goto if_then1554
	} else {
		goto lor_lhs_false1545
	}

lor_lhs_false1545:
	v643 = *libc.As[int32](lookahead)
	cmp1546 = v643 == 95
	if cmp1546 {
		goto if_then1554
	} else {
		goto lor_lhs_false1548
	}

lor_lhs_false1548:
	v644 = *libc.As[int32](lookahead)
	cmp1549 = 97 <= v644
	if cmp1549 {
		goto land_lhs_true1551
	} else {
		goto if_end1555
	}

land_lhs_true1551:
	v645 = *libc.As[int32](lookahead)
	cmp1552 = v645 <= 122
	if cmp1552 {
		goto if_then1554
	} else {
		goto if_end1555
	}

if_then1554:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1555:
	v646 = *libc.As[byte](result)
	loadedv1556 = (v646 & 1) != 0
	*libc.As[bool](retval) = loadedv1556
	goto _return

sw_bb1557:
	*libc.As[byte](result) = 1
	v647 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1558 = libc.Ptr(&libc.As[TSLexer](v647).F1)
	*libc.As[int16](result_symbol1558) = 10
	v648 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1559 = libc.Ptr(&libc.As[TSLexer](v648).F3)
	v649 = *libc.As[unsafe.Pointer](mark_end1559)
	v650 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v649)(v650)
	v651 = *libc.As[int32](lookahead)
	cmp1560 = v651 == 111
	if cmp1560 {
		goto if_then1562
	} else {
		goto if_end1563
	}

if_then1562:
	*libc.As[int16](state_addr) = 107
	goto next_state

if_end1563:
	v652 = *libc.As[int32](lookahead)
	cmp1564 = v652 == 45
	if cmp1564 {
		goto if_then1590
	} else {
		goto lor_lhs_false1566
	}

lor_lhs_false1566:
	v653 = *libc.As[int32](lookahead)
	cmp1567 = v653 == 46
	if cmp1567 {
		goto if_then1590
	} else {
		goto lor_lhs_false1569
	}

lor_lhs_false1569:
	v654 = *libc.As[int32](lookahead)
	cmp1570 = 48 <= v654
	if cmp1570 {
		goto land_lhs_true1572
	} else {
		goto lor_lhs_false1575
	}

land_lhs_true1572:
	v655 = *libc.As[int32](lookahead)
	cmp1573 = v655 <= 57
	if cmp1573 {
		goto if_then1590
	} else {
		goto lor_lhs_false1575
	}

lor_lhs_false1575:
	v656 = *libc.As[int32](lookahead)
	cmp1576 = 65 <= v656
	if cmp1576 {
		goto land_lhs_true1578
	} else {
		goto lor_lhs_false1581
	}

land_lhs_true1578:
	v657 = *libc.As[int32](lookahead)
	cmp1579 = v657 <= 90
	if cmp1579 {
		goto if_then1590
	} else {
		goto lor_lhs_false1581
	}

lor_lhs_false1581:
	v658 = *libc.As[int32](lookahead)
	cmp1582 = v658 == 95
	if cmp1582 {
		goto if_then1590
	} else {
		goto lor_lhs_false1584
	}

lor_lhs_false1584:
	v659 = *libc.As[int32](lookahead)
	cmp1585 = 97 <= v659
	if cmp1585 {
		goto land_lhs_true1587
	} else {
		goto if_end1591
	}

land_lhs_true1587:
	v660 = *libc.As[int32](lookahead)
	cmp1588 = v660 <= 122
	if cmp1588 {
		goto if_then1590
	} else {
		goto if_end1591
	}

if_then1590:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1591:
	v661 = *libc.As[byte](result)
	loadedv1592 = (v661 & 1) != 0
	*libc.As[bool](retval) = loadedv1592
	goto _return

sw_bb1593:
	*libc.As[byte](result) = 1
	v662 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1594 = libc.Ptr(&libc.As[TSLexer](v662).F1)
	*libc.As[int16](result_symbol1594) = 10
	v663 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1595 = libc.Ptr(&libc.As[TSLexer](v663).F3)
	v664 = *libc.As[unsafe.Pointer](mark_end1595)
	v665 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v664)(v665)
	v666 = *libc.As[int32](lookahead)
	cmp1596 = v666 == 111
	if cmp1596 {
		goto if_then1598
	} else {
		goto if_end1599
	}

if_then1598:
	*libc.As[int16](state_addr) = 127
	goto next_state

if_end1599:
	v667 = *libc.As[int32](lookahead)
	cmp1600 = v667 == 45
	if cmp1600 {
		goto if_then1626
	} else {
		goto lor_lhs_false1602
	}

lor_lhs_false1602:
	v668 = *libc.As[int32](lookahead)
	cmp1603 = v668 == 46
	if cmp1603 {
		goto if_then1626
	} else {
		goto lor_lhs_false1605
	}

lor_lhs_false1605:
	v669 = *libc.As[int32](lookahead)
	cmp1606 = 48 <= v669
	if cmp1606 {
		goto land_lhs_true1608
	} else {
		goto lor_lhs_false1611
	}

land_lhs_true1608:
	v670 = *libc.As[int32](lookahead)
	cmp1609 = v670 <= 57
	if cmp1609 {
		goto if_then1626
	} else {
		goto lor_lhs_false1611
	}

lor_lhs_false1611:
	v671 = *libc.As[int32](lookahead)
	cmp1612 = 65 <= v671
	if cmp1612 {
		goto land_lhs_true1614
	} else {
		goto lor_lhs_false1617
	}

land_lhs_true1614:
	v672 = *libc.As[int32](lookahead)
	cmp1615 = v672 <= 90
	if cmp1615 {
		goto if_then1626
	} else {
		goto lor_lhs_false1617
	}

lor_lhs_false1617:
	v673 = *libc.As[int32](lookahead)
	cmp1618 = v673 == 95
	if cmp1618 {
		goto if_then1626
	} else {
		goto lor_lhs_false1620
	}

lor_lhs_false1620:
	v674 = *libc.As[int32](lookahead)
	cmp1621 = 97 <= v674
	if cmp1621 {
		goto land_lhs_true1623
	} else {
		goto if_end1627
	}

land_lhs_true1623:
	v675 = *libc.As[int32](lookahead)
	cmp1624 = v675 <= 122
	if cmp1624 {
		goto if_then1626
	} else {
		goto if_end1627
	}

if_then1626:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1627:
	v676 = *libc.As[byte](result)
	loadedv1628 = (v676 & 1) != 0
	*libc.As[bool](retval) = loadedv1628
	goto _return

sw_bb1629:
	*libc.As[byte](result) = 1
	v677 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1630 = libc.Ptr(&libc.As[TSLexer](v677).F1)
	*libc.As[int16](result_symbol1630) = 10
	v678 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1631 = libc.Ptr(&libc.As[TSLexer](v678).F3)
	v679 = *libc.As[unsafe.Pointer](mark_end1631)
	v680 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v679)(v680)
	v681 = *libc.As[int32](lookahead)
	cmp1632 = v681 == 111
	if cmp1632 {
		goto if_then1634
	} else {
		goto if_end1635
	}

if_then1634:
	*libc.As[int16](state_addr) = 104
	goto next_state

if_end1635:
	v682 = *libc.As[int32](lookahead)
	cmp1636 = v682 == 45
	if cmp1636 {
		goto if_then1662
	} else {
		goto lor_lhs_false1638
	}

lor_lhs_false1638:
	v683 = *libc.As[int32](lookahead)
	cmp1639 = v683 == 46
	if cmp1639 {
		goto if_then1662
	} else {
		goto lor_lhs_false1641
	}

lor_lhs_false1641:
	v684 = *libc.As[int32](lookahead)
	cmp1642 = 48 <= v684
	if cmp1642 {
		goto land_lhs_true1644
	} else {
		goto lor_lhs_false1647
	}

land_lhs_true1644:
	v685 = *libc.As[int32](lookahead)
	cmp1645 = v685 <= 57
	if cmp1645 {
		goto if_then1662
	} else {
		goto lor_lhs_false1647
	}

lor_lhs_false1647:
	v686 = *libc.As[int32](lookahead)
	cmp1648 = 65 <= v686
	if cmp1648 {
		goto land_lhs_true1650
	} else {
		goto lor_lhs_false1653
	}

land_lhs_true1650:
	v687 = *libc.As[int32](lookahead)
	cmp1651 = v687 <= 90
	if cmp1651 {
		goto if_then1662
	} else {
		goto lor_lhs_false1653
	}

lor_lhs_false1653:
	v688 = *libc.As[int32](lookahead)
	cmp1654 = v688 == 95
	if cmp1654 {
		goto if_then1662
	} else {
		goto lor_lhs_false1656
	}

lor_lhs_false1656:
	v689 = *libc.As[int32](lookahead)
	cmp1657 = 97 <= v689
	if cmp1657 {
		goto land_lhs_true1659
	} else {
		goto if_end1663
	}

land_lhs_true1659:
	v690 = *libc.As[int32](lookahead)
	cmp1660 = v690 <= 122
	if cmp1660 {
		goto if_then1662
	} else {
		goto if_end1663
	}

if_then1662:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1663:
	v691 = *libc.As[byte](result)
	loadedv1664 = (v691 & 1) != 0
	*libc.As[bool](retval) = loadedv1664
	goto _return

sw_bb1665:
	*libc.As[byte](result) = 1
	v692 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1666 = libc.Ptr(&libc.As[TSLexer](v692).F1)
	*libc.As[int16](result_symbol1666) = 10
	v693 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1667 = libc.Ptr(&libc.As[TSLexer](v693).F3)
	v694 = *libc.As[unsafe.Pointer](mark_end1667)
	v695 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v694)(v695)
	v696 = *libc.As[int32](lookahead)
	cmp1668 = v696 == 111
	if cmp1668 {
		goto if_then1670
	} else {
		goto if_end1671
	}

if_then1670:
	*libc.As[int16](state_addr) = 111
	goto next_state

if_end1671:
	v697 = *libc.As[int32](lookahead)
	cmp1672 = v697 == 45
	if cmp1672 {
		goto if_then1698
	} else {
		goto lor_lhs_false1674
	}

lor_lhs_false1674:
	v698 = *libc.As[int32](lookahead)
	cmp1675 = v698 == 46
	if cmp1675 {
		goto if_then1698
	} else {
		goto lor_lhs_false1677
	}

lor_lhs_false1677:
	v699 = *libc.As[int32](lookahead)
	cmp1678 = 48 <= v699
	if cmp1678 {
		goto land_lhs_true1680
	} else {
		goto lor_lhs_false1683
	}

land_lhs_true1680:
	v700 = *libc.As[int32](lookahead)
	cmp1681 = v700 <= 57
	if cmp1681 {
		goto if_then1698
	} else {
		goto lor_lhs_false1683
	}

lor_lhs_false1683:
	v701 = *libc.As[int32](lookahead)
	cmp1684 = 65 <= v701
	if cmp1684 {
		goto land_lhs_true1686
	} else {
		goto lor_lhs_false1689
	}

land_lhs_true1686:
	v702 = *libc.As[int32](lookahead)
	cmp1687 = v702 <= 90
	if cmp1687 {
		goto if_then1698
	} else {
		goto lor_lhs_false1689
	}

lor_lhs_false1689:
	v703 = *libc.As[int32](lookahead)
	cmp1690 = v703 == 95
	if cmp1690 {
		goto if_then1698
	} else {
		goto lor_lhs_false1692
	}

lor_lhs_false1692:
	v704 = *libc.As[int32](lookahead)
	cmp1693 = 97 <= v704
	if cmp1693 {
		goto land_lhs_true1695
	} else {
		goto if_end1699
	}

land_lhs_true1695:
	v705 = *libc.As[int32](lookahead)
	cmp1696 = v705 <= 122
	if cmp1696 {
		goto if_then1698
	} else {
		goto if_end1699
	}

if_then1698:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1699:
	v706 = *libc.As[byte](result)
	loadedv1700 = (v706 & 1) != 0
	*libc.As[bool](retval) = loadedv1700
	goto _return

sw_bb1701:
	*libc.As[byte](result) = 1
	v707 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1702 = libc.Ptr(&libc.As[TSLexer](v707).F1)
	*libc.As[int16](result_symbol1702) = 10
	v708 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1703 = libc.Ptr(&libc.As[TSLexer](v708).F3)
	v709 = *libc.As[unsafe.Pointer](mark_end1703)
	v710 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v709)(v710)
	v711 = *libc.As[int32](lookahead)
	cmp1704 = v711 == 112
	if cmp1704 {
		goto if_then1706
	} else {
		goto if_end1707
	}

if_then1706:
	*libc.As[int16](state_addr) = 98
	goto next_state

if_end1707:
	v712 = *libc.As[int32](lookahead)
	cmp1708 = v712 == 45
	if cmp1708 {
		goto if_then1734
	} else {
		goto lor_lhs_false1710
	}

lor_lhs_false1710:
	v713 = *libc.As[int32](lookahead)
	cmp1711 = v713 == 46
	if cmp1711 {
		goto if_then1734
	} else {
		goto lor_lhs_false1713
	}

lor_lhs_false1713:
	v714 = *libc.As[int32](lookahead)
	cmp1714 = 48 <= v714
	if cmp1714 {
		goto land_lhs_true1716
	} else {
		goto lor_lhs_false1719
	}

land_lhs_true1716:
	v715 = *libc.As[int32](lookahead)
	cmp1717 = v715 <= 57
	if cmp1717 {
		goto if_then1734
	} else {
		goto lor_lhs_false1719
	}

lor_lhs_false1719:
	v716 = *libc.As[int32](lookahead)
	cmp1720 = 65 <= v716
	if cmp1720 {
		goto land_lhs_true1722
	} else {
		goto lor_lhs_false1725
	}

land_lhs_true1722:
	v717 = *libc.As[int32](lookahead)
	cmp1723 = v717 <= 90
	if cmp1723 {
		goto if_then1734
	} else {
		goto lor_lhs_false1725
	}

lor_lhs_false1725:
	v718 = *libc.As[int32](lookahead)
	cmp1726 = v718 == 95
	if cmp1726 {
		goto if_then1734
	} else {
		goto lor_lhs_false1728
	}

lor_lhs_false1728:
	v719 = *libc.As[int32](lookahead)
	cmp1729 = 97 <= v719
	if cmp1729 {
		goto land_lhs_true1731
	} else {
		goto if_end1735
	}

land_lhs_true1731:
	v720 = *libc.As[int32](lookahead)
	cmp1732 = v720 <= 122
	if cmp1732 {
		goto if_then1734
	} else {
		goto if_end1735
	}

if_then1734:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1735:
	v721 = *libc.As[byte](result)
	loadedv1736 = (v721 & 1) != 0
	*libc.As[bool](retval) = loadedv1736
	goto _return

sw_bb1737:
	*libc.As[byte](result) = 1
	v722 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1738 = libc.Ptr(&libc.As[TSLexer](v722).F1)
	*libc.As[int16](result_symbol1738) = 10
	v723 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1739 = libc.Ptr(&libc.As[TSLexer](v723).F3)
	v724 = *libc.As[unsafe.Pointer](mark_end1739)
	v725 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v724)(v725)
	v726 = *libc.As[int32](lookahead)
	cmp1740 = v726 == 114
	if cmp1740 {
		goto if_then1742
	} else {
		goto if_end1743
	}

if_then1742:
	*libc.As[int16](state_addr) = 89
	goto next_state

if_end1743:
	v727 = *libc.As[int32](lookahead)
	cmp1744 = v727 == 45
	if cmp1744 {
		goto if_then1770
	} else {
		goto lor_lhs_false1746
	}

lor_lhs_false1746:
	v728 = *libc.As[int32](lookahead)
	cmp1747 = v728 == 46
	if cmp1747 {
		goto if_then1770
	} else {
		goto lor_lhs_false1749
	}

lor_lhs_false1749:
	v729 = *libc.As[int32](lookahead)
	cmp1750 = 48 <= v729
	if cmp1750 {
		goto land_lhs_true1752
	} else {
		goto lor_lhs_false1755
	}

land_lhs_true1752:
	v730 = *libc.As[int32](lookahead)
	cmp1753 = v730 <= 57
	if cmp1753 {
		goto if_then1770
	} else {
		goto lor_lhs_false1755
	}

lor_lhs_false1755:
	v731 = *libc.As[int32](lookahead)
	cmp1756 = 65 <= v731
	if cmp1756 {
		goto land_lhs_true1758
	} else {
		goto lor_lhs_false1761
	}

land_lhs_true1758:
	v732 = *libc.As[int32](lookahead)
	cmp1759 = v732 <= 90
	if cmp1759 {
		goto if_then1770
	} else {
		goto lor_lhs_false1761
	}

lor_lhs_false1761:
	v733 = *libc.As[int32](lookahead)
	cmp1762 = v733 == 95
	if cmp1762 {
		goto if_then1770
	} else {
		goto lor_lhs_false1764
	}

lor_lhs_false1764:
	v734 = *libc.As[int32](lookahead)
	cmp1765 = 97 <= v734
	if cmp1765 {
		goto land_lhs_true1767
	} else {
		goto if_end1771
	}

land_lhs_true1767:
	v735 = *libc.As[int32](lookahead)
	cmp1768 = v735 <= 122
	if cmp1768 {
		goto if_then1770
	} else {
		goto if_end1771
	}

if_then1770:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1771:
	v736 = *libc.As[byte](result)
	loadedv1772 = (v736 & 1) != 0
	*libc.As[bool](retval) = loadedv1772
	goto _return

sw_bb1773:
	*libc.As[byte](result) = 1
	v737 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1774 = libc.Ptr(&libc.As[TSLexer](v737).F1)
	*libc.As[int16](result_symbol1774) = 10
	v738 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1775 = libc.Ptr(&libc.As[TSLexer](v738).F3)
	v739 = *libc.As[unsafe.Pointer](mark_end1775)
	v740 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v739)(v740)
	v741 = *libc.As[int32](lookahead)
	cmp1776 = v741 == 114
	if cmp1776 {
		goto if_then1778
	} else {
		goto if_end1779
	}

if_then1778:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end1779:
	v742 = *libc.As[int32](lookahead)
	cmp1780 = v742 == 45
	if cmp1780 {
		goto if_then1806
	} else {
		goto lor_lhs_false1782
	}

lor_lhs_false1782:
	v743 = *libc.As[int32](lookahead)
	cmp1783 = v743 == 46
	if cmp1783 {
		goto if_then1806
	} else {
		goto lor_lhs_false1785
	}

lor_lhs_false1785:
	v744 = *libc.As[int32](lookahead)
	cmp1786 = 48 <= v744
	if cmp1786 {
		goto land_lhs_true1788
	} else {
		goto lor_lhs_false1791
	}

land_lhs_true1788:
	v745 = *libc.As[int32](lookahead)
	cmp1789 = v745 <= 57
	if cmp1789 {
		goto if_then1806
	} else {
		goto lor_lhs_false1791
	}

lor_lhs_false1791:
	v746 = *libc.As[int32](lookahead)
	cmp1792 = 65 <= v746
	if cmp1792 {
		goto land_lhs_true1794
	} else {
		goto lor_lhs_false1797
	}

land_lhs_true1794:
	v747 = *libc.As[int32](lookahead)
	cmp1795 = v747 <= 90
	if cmp1795 {
		goto if_then1806
	} else {
		goto lor_lhs_false1797
	}

lor_lhs_false1797:
	v748 = *libc.As[int32](lookahead)
	cmp1798 = v748 == 95
	if cmp1798 {
		goto if_then1806
	} else {
		goto lor_lhs_false1800
	}

lor_lhs_false1800:
	v749 = *libc.As[int32](lookahead)
	cmp1801 = 97 <= v749
	if cmp1801 {
		goto land_lhs_true1803
	} else {
		goto if_end1807
	}

land_lhs_true1803:
	v750 = *libc.As[int32](lookahead)
	cmp1804 = v750 <= 122
	if cmp1804 {
		goto if_then1806
	} else {
		goto if_end1807
	}

if_then1806:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1807:
	v751 = *libc.As[byte](result)
	loadedv1808 = (v751 & 1) != 0
	*libc.As[bool](retval) = loadedv1808
	goto _return

sw_bb1809:
	*libc.As[byte](result) = 1
	v752 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1810 = libc.Ptr(&libc.As[TSLexer](v752).F1)
	*libc.As[int16](result_symbol1810) = 10
	v753 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1811 = libc.Ptr(&libc.As[TSLexer](v753).F3)
	v754 = *libc.As[unsafe.Pointer](mark_end1811)
	v755 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v754)(v755)
	v756 = *libc.As[int32](lookahead)
	cmp1812 = v756 == 114
	if cmp1812 {
		goto if_then1814
	} else {
		goto if_end1815
	}

if_then1814:
	*libc.As[int16](state_addr) = 116
	goto next_state

if_end1815:
	v757 = *libc.As[int32](lookahead)
	cmp1816 = v757 == 45
	if cmp1816 {
		goto if_then1842
	} else {
		goto lor_lhs_false1818
	}

lor_lhs_false1818:
	v758 = *libc.As[int32](lookahead)
	cmp1819 = v758 == 46
	if cmp1819 {
		goto if_then1842
	} else {
		goto lor_lhs_false1821
	}

lor_lhs_false1821:
	v759 = *libc.As[int32](lookahead)
	cmp1822 = 48 <= v759
	if cmp1822 {
		goto land_lhs_true1824
	} else {
		goto lor_lhs_false1827
	}

land_lhs_true1824:
	v760 = *libc.As[int32](lookahead)
	cmp1825 = v760 <= 57
	if cmp1825 {
		goto if_then1842
	} else {
		goto lor_lhs_false1827
	}

lor_lhs_false1827:
	v761 = *libc.As[int32](lookahead)
	cmp1828 = 65 <= v761
	if cmp1828 {
		goto land_lhs_true1830
	} else {
		goto lor_lhs_false1833
	}

land_lhs_true1830:
	v762 = *libc.As[int32](lookahead)
	cmp1831 = v762 <= 90
	if cmp1831 {
		goto if_then1842
	} else {
		goto lor_lhs_false1833
	}

lor_lhs_false1833:
	v763 = *libc.As[int32](lookahead)
	cmp1834 = v763 == 95
	if cmp1834 {
		goto if_then1842
	} else {
		goto lor_lhs_false1836
	}

lor_lhs_false1836:
	v764 = *libc.As[int32](lookahead)
	cmp1837 = 97 <= v764
	if cmp1837 {
		goto land_lhs_true1839
	} else {
		goto if_end1843
	}

land_lhs_true1839:
	v765 = *libc.As[int32](lookahead)
	cmp1840 = v765 <= 122
	if cmp1840 {
		goto if_then1842
	} else {
		goto if_end1843
	}

if_then1842:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1843:
	v766 = *libc.As[byte](result)
	loadedv1844 = (v766 & 1) != 0
	*libc.As[bool](retval) = loadedv1844
	goto _return

sw_bb1845:
	*libc.As[byte](result) = 1
	v767 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1846 = libc.Ptr(&libc.As[TSLexer](v767).F1)
	*libc.As[int16](result_symbol1846) = 10
	v768 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1847 = libc.Ptr(&libc.As[TSLexer](v768).F3)
	v769 = *libc.As[unsafe.Pointer](mark_end1847)
	v770 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v769)(v770)
	v771 = *libc.As[int32](lookahead)
	cmp1848 = v771 == 114
	if cmp1848 {
		goto if_then1850
	} else {
		goto if_end1851
	}

if_then1850:
	*libc.As[int16](state_addr) = 102
	goto next_state

if_end1851:
	v772 = *libc.As[int32](lookahead)
	cmp1852 = v772 == 45
	if cmp1852 {
		goto if_then1878
	} else {
		goto lor_lhs_false1854
	}

lor_lhs_false1854:
	v773 = *libc.As[int32](lookahead)
	cmp1855 = v773 == 46
	if cmp1855 {
		goto if_then1878
	} else {
		goto lor_lhs_false1857
	}

lor_lhs_false1857:
	v774 = *libc.As[int32](lookahead)
	cmp1858 = 48 <= v774
	if cmp1858 {
		goto land_lhs_true1860
	} else {
		goto lor_lhs_false1863
	}

land_lhs_true1860:
	v775 = *libc.As[int32](lookahead)
	cmp1861 = v775 <= 57
	if cmp1861 {
		goto if_then1878
	} else {
		goto lor_lhs_false1863
	}

lor_lhs_false1863:
	v776 = *libc.As[int32](lookahead)
	cmp1864 = 65 <= v776
	if cmp1864 {
		goto land_lhs_true1866
	} else {
		goto lor_lhs_false1869
	}

land_lhs_true1866:
	v777 = *libc.As[int32](lookahead)
	cmp1867 = v777 <= 90
	if cmp1867 {
		goto if_then1878
	} else {
		goto lor_lhs_false1869
	}

lor_lhs_false1869:
	v778 = *libc.As[int32](lookahead)
	cmp1870 = v778 == 95
	if cmp1870 {
		goto if_then1878
	} else {
		goto lor_lhs_false1872
	}

lor_lhs_false1872:
	v779 = *libc.As[int32](lookahead)
	cmp1873 = 97 <= v779
	if cmp1873 {
		goto land_lhs_true1875
	} else {
		goto if_end1879
	}

land_lhs_true1875:
	v780 = *libc.As[int32](lookahead)
	cmp1876 = v780 <= 122
	if cmp1876 {
		goto if_then1878
	} else {
		goto if_end1879
	}

if_then1878:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1879:
	v781 = *libc.As[byte](result)
	loadedv1880 = (v781 & 1) != 0
	*libc.As[bool](retval) = loadedv1880
	goto _return

sw_bb1881:
	*libc.As[byte](result) = 1
	v782 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1882 = libc.Ptr(&libc.As[TSLexer](v782).F1)
	*libc.As[int16](result_symbol1882) = 10
	v783 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1883 = libc.Ptr(&libc.As[TSLexer](v783).F3)
	v784 = *libc.As[unsafe.Pointer](mark_end1883)
	v785 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v784)(v785)
	v786 = *libc.As[int32](lookahead)
	cmp1884 = v786 == 115
	if cmp1884 {
		goto if_then1886
	} else {
		goto if_end1887
	}

if_then1886:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1887:
	v787 = *libc.As[int32](lookahead)
	cmp1888 = v787 == 45
	if cmp1888 {
		goto if_then1914
	} else {
		goto lor_lhs_false1890
	}

lor_lhs_false1890:
	v788 = *libc.As[int32](lookahead)
	cmp1891 = v788 == 46
	if cmp1891 {
		goto if_then1914
	} else {
		goto lor_lhs_false1893
	}

lor_lhs_false1893:
	v789 = *libc.As[int32](lookahead)
	cmp1894 = 48 <= v789
	if cmp1894 {
		goto land_lhs_true1896
	} else {
		goto lor_lhs_false1899
	}

land_lhs_true1896:
	v790 = *libc.As[int32](lookahead)
	cmp1897 = v790 <= 57
	if cmp1897 {
		goto if_then1914
	} else {
		goto lor_lhs_false1899
	}

lor_lhs_false1899:
	v791 = *libc.As[int32](lookahead)
	cmp1900 = 65 <= v791
	if cmp1900 {
		goto land_lhs_true1902
	} else {
		goto lor_lhs_false1905
	}

land_lhs_true1902:
	v792 = *libc.As[int32](lookahead)
	cmp1903 = v792 <= 90
	if cmp1903 {
		goto if_then1914
	} else {
		goto lor_lhs_false1905
	}

lor_lhs_false1905:
	v793 = *libc.As[int32](lookahead)
	cmp1906 = v793 == 95
	if cmp1906 {
		goto if_then1914
	} else {
		goto lor_lhs_false1908
	}

lor_lhs_false1908:
	v794 = *libc.As[int32](lookahead)
	cmp1909 = 97 <= v794
	if cmp1909 {
		goto land_lhs_true1911
	} else {
		goto if_end1915
	}

land_lhs_true1911:
	v795 = *libc.As[int32](lookahead)
	cmp1912 = v795 <= 122
	if cmp1912 {
		goto if_then1914
	} else {
		goto if_end1915
	}

if_then1914:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1915:
	v796 = *libc.As[byte](result)
	loadedv1916 = (v796 & 1) != 0
	*libc.As[bool](retval) = loadedv1916
	goto _return

sw_bb1917:
	*libc.As[byte](result) = 1
	v797 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1918 = libc.Ptr(&libc.As[TSLexer](v797).F1)
	*libc.As[int16](result_symbol1918) = 10
	v798 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1919 = libc.Ptr(&libc.As[TSLexer](v798).F3)
	v799 = *libc.As[unsafe.Pointer](mark_end1919)
	v800 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v799)(v800)
	v801 = *libc.As[int32](lookahead)
	cmp1920 = v801 == 115
	if cmp1920 {
		goto if_then1922
	} else {
		goto if_end1923
	}

if_then1922:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end1923:
	v802 = *libc.As[int32](lookahead)
	cmp1924 = v802 == 45
	if cmp1924 {
		goto if_then1950
	} else {
		goto lor_lhs_false1926
	}

lor_lhs_false1926:
	v803 = *libc.As[int32](lookahead)
	cmp1927 = v803 == 46
	if cmp1927 {
		goto if_then1950
	} else {
		goto lor_lhs_false1929
	}

lor_lhs_false1929:
	v804 = *libc.As[int32](lookahead)
	cmp1930 = 48 <= v804
	if cmp1930 {
		goto land_lhs_true1932
	} else {
		goto lor_lhs_false1935
	}

land_lhs_true1932:
	v805 = *libc.As[int32](lookahead)
	cmp1933 = v805 <= 57
	if cmp1933 {
		goto if_then1950
	} else {
		goto lor_lhs_false1935
	}

lor_lhs_false1935:
	v806 = *libc.As[int32](lookahead)
	cmp1936 = 65 <= v806
	if cmp1936 {
		goto land_lhs_true1938
	} else {
		goto lor_lhs_false1941
	}

land_lhs_true1938:
	v807 = *libc.As[int32](lookahead)
	cmp1939 = v807 <= 90
	if cmp1939 {
		goto if_then1950
	} else {
		goto lor_lhs_false1941
	}

lor_lhs_false1941:
	v808 = *libc.As[int32](lookahead)
	cmp1942 = v808 == 95
	if cmp1942 {
		goto if_then1950
	} else {
		goto lor_lhs_false1944
	}

lor_lhs_false1944:
	v809 = *libc.As[int32](lookahead)
	cmp1945 = 97 <= v809
	if cmp1945 {
		goto land_lhs_true1947
	} else {
		goto if_end1951
	}

land_lhs_true1947:
	v810 = *libc.As[int32](lookahead)
	cmp1948 = v810 <= 122
	if cmp1948 {
		goto if_then1950
	} else {
		goto if_end1951
	}

if_then1950:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1951:
	v811 = *libc.As[byte](result)
	loadedv1952 = (v811 & 1) != 0
	*libc.As[bool](retval) = loadedv1952
	goto _return

sw_bb1953:
	*libc.As[byte](result) = 1
	v812 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1954 = libc.Ptr(&libc.As[TSLexer](v812).F1)
	*libc.As[int16](result_symbol1954) = 10
	v813 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1955 = libc.Ptr(&libc.As[TSLexer](v813).F3)
	v814 = *libc.As[unsafe.Pointer](mark_end1955)
	v815 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v814)(v815)
	v816 = *libc.As[int32](lookahead)
	cmp1956 = v816 == 116
	if cmp1956 {
		goto if_then1958
	} else {
		goto if_end1959
	}

if_then1958:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1959:
	v817 = *libc.As[int32](lookahead)
	cmp1960 = v817 == 45
	if cmp1960 {
		goto if_then1986
	} else {
		goto lor_lhs_false1962
	}

lor_lhs_false1962:
	v818 = *libc.As[int32](lookahead)
	cmp1963 = v818 == 46
	if cmp1963 {
		goto if_then1986
	} else {
		goto lor_lhs_false1965
	}

lor_lhs_false1965:
	v819 = *libc.As[int32](lookahead)
	cmp1966 = 48 <= v819
	if cmp1966 {
		goto land_lhs_true1968
	} else {
		goto lor_lhs_false1971
	}

land_lhs_true1968:
	v820 = *libc.As[int32](lookahead)
	cmp1969 = v820 <= 57
	if cmp1969 {
		goto if_then1986
	} else {
		goto lor_lhs_false1971
	}

lor_lhs_false1971:
	v821 = *libc.As[int32](lookahead)
	cmp1972 = 65 <= v821
	if cmp1972 {
		goto land_lhs_true1974
	} else {
		goto lor_lhs_false1977
	}

land_lhs_true1974:
	v822 = *libc.As[int32](lookahead)
	cmp1975 = v822 <= 90
	if cmp1975 {
		goto if_then1986
	} else {
		goto lor_lhs_false1977
	}

lor_lhs_false1977:
	v823 = *libc.As[int32](lookahead)
	cmp1978 = v823 == 95
	if cmp1978 {
		goto if_then1986
	} else {
		goto lor_lhs_false1980
	}

lor_lhs_false1980:
	v824 = *libc.As[int32](lookahead)
	cmp1981 = 97 <= v824
	if cmp1981 {
		goto land_lhs_true1983
	} else {
		goto if_end1987
	}

land_lhs_true1983:
	v825 = *libc.As[int32](lookahead)
	cmp1984 = v825 <= 122
	if cmp1984 {
		goto if_then1986
	} else {
		goto if_end1987
	}

if_then1986:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end1987:
	v826 = *libc.As[byte](result)
	loadedv1988 = (v826 & 1) != 0
	*libc.As[bool](retval) = loadedv1988
	goto _return

sw_bb1989:
	*libc.As[byte](result) = 1
	v827 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1990 = libc.Ptr(&libc.As[TSLexer](v827).F1)
	*libc.As[int16](result_symbol1990) = 10
	v828 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1991 = libc.Ptr(&libc.As[TSLexer](v828).F3)
	v829 = *libc.As[unsafe.Pointer](mark_end1991)
	v830 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v829)(v830)
	v831 = *libc.As[int32](lookahead)
	cmp1992 = v831 == 117
	if cmp1992 {
		goto if_then1994
	} else {
		goto if_end1995
	}

if_then1994:
	*libc.As[int16](state_addr) = 128
	goto next_state

if_end1995:
	v832 = *libc.As[int32](lookahead)
	cmp1996 = v832 == 45
	if cmp1996 {
		goto if_then2022
	} else {
		goto lor_lhs_false1998
	}

lor_lhs_false1998:
	v833 = *libc.As[int32](lookahead)
	cmp1999 = v833 == 46
	if cmp1999 {
		goto if_then2022
	} else {
		goto lor_lhs_false2001
	}

lor_lhs_false2001:
	v834 = *libc.As[int32](lookahead)
	cmp2002 = 48 <= v834
	if cmp2002 {
		goto land_lhs_true2004
	} else {
		goto lor_lhs_false2007
	}

land_lhs_true2004:
	v835 = *libc.As[int32](lookahead)
	cmp2005 = v835 <= 57
	if cmp2005 {
		goto if_then2022
	} else {
		goto lor_lhs_false2007
	}

lor_lhs_false2007:
	v836 = *libc.As[int32](lookahead)
	cmp2008 = 65 <= v836
	if cmp2008 {
		goto land_lhs_true2010
	} else {
		goto lor_lhs_false2013
	}

land_lhs_true2010:
	v837 = *libc.As[int32](lookahead)
	cmp2011 = v837 <= 90
	if cmp2011 {
		goto if_then2022
	} else {
		goto lor_lhs_false2013
	}

lor_lhs_false2013:
	v838 = *libc.As[int32](lookahead)
	cmp2014 = v838 == 95
	if cmp2014 {
		goto if_then2022
	} else {
		goto lor_lhs_false2016
	}

lor_lhs_false2016:
	v839 = *libc.As[int32](lookahead)
	cmp2017 = 97 <= v839
	if cmp2017 {
		goto land_lhs_true2019
	} else {
		goto if_end2023
	}

land_lhs_true2019:
	v840 = *libc.As[int32](lookahead)
	cmp2020 = v840 <= 122
	if cmp2020 {
		goto if_then2022
	} else {
		goto if_end2023
	}

if_then2022:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end2023:
	v841 = *libc.As[byte](result)
	loadedv2024 = (v841 & 1) != 0
	*libc.As[bool](retval) = loadedv2024
	goto _return

sw_bb2025:
	*libc.As[byte](result) = 1
	v842 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2026 = libc.Ptr(&libc.As[TSLexer](v842).F1)
	*libc.As[int16](result_symbol2026) = 10
	v843 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2027 = libc.Ptr(&libc.As[TSLexer](v843).F3)
	v844 = *libc.As[unsafe.Pointer](mark_end2027)
	v845 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v844)(v845)
	v846 = *libc.As[int32](lookahead)
	cmp2028 = v846 == 119
	if cmp2028 {
		goto if_then2030
	} else {
		goto if_end2031
	}

if_then2030:
	*libc.As[int16](state_addr) = 123
	goto next_state

if_end2031:
	v847 = *libc.As[int32](lookahead)
	cmp2032 = v847 == 45
	if cmp2032 {
		goto if_then2058
	} else {
		goto lor_lhs_false2034
	}

lor_lhs_false2034:
	v848 = *libc.As[int32](lookahead)
	cmp2035 = v848 == 46
	if cmp2035 {
		goto if_then2058
	} else {
		goto lor_lhs_false2037
	}

lor_lhs_false2037:
	v849 = *libc.As[int32](lookahead)
	cmp2038 = 48 <= v849
	if cmp2038 {
		goto land_lhs_true2040
	} else {
		goto lor_lhs_false2043
	}

land_lhs_true2040:
	v850 = *libc.As[int32](lookahead)
	cmp2041 = v850 <= 57
	if cmp2041 {
		goto if_then2058
	} else {
		goto lor_lhs_false2043
	}

lor_lhs_false2043:
	v851 = *libc.As[int32](lookahead)
	cmp2044 = 65 <= v851
	if cmp2044 {
		goto land_lhs_true2046
	} else {
		goto lor_lhs_false2049
	}

land_lhs_true2046:
	v852 = *libc.As[int32](lookahead)
	cmp2047 = v852 <= 90
	if cmp2047 {
		goto if_then2058
	} else {
		goto lor_lhs_false2049
	}

lor_lhs_false2049:
	v853 = *libc.As[int32](lookahead)
	cmp2050 = v853 == 95
	if cmp2050 {
		goto if_then2058
	} else {
		goto lor_lhs_false2052
	}

lor_lhs_false2052:
	v854 = *libc.As[int32](lookahead)
	cmp2053 = 97 <= v854
	if cmp2053 {
		goto land_lhs_true2055
	} else {
		goto if_end2059
	}

land_lhs_true2055:
	v855 = *libc.As[int32](lookahead)
	cmp2056 = v855 <= 122
	if cmp2056 {
		goto if_then2058
	} else {
		goto if_end2059
	}

if_then2058:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end2059:
	v856 = *libc.As[byte](result)
	loadedv2060 = (v856 & 1) != 0
	*libc.As[bool](retval) = loadedv2060
	goto _return

sw_bb2061:
	*libc.As[byte](result) = 1
	v857 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2062 = libc.Ptr(&libc.As[TSLexer](v857).F1)
	*libc.As[int16](result_symbol2062) = 10
	v858 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2063 = libc.Ptr(&libc.As[TSLexer](v858).F3)
	v859 = *libc.As[unsafe.Pointer](mark_end2063)
	v860 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v859)(v860)
	v861 = *libc.As[int32](lookahead)
	cmp2064 = v861 == 120
	if cmp2064 {
		goto if_then2066
	} else {
		goto if_end2067
	}

if_then2066:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end2067:
	v862 = *libc.As[int32](lookahead)
	cmp2068 = v862 == 45
	if cmp2068 {
		goto if_then2094
	} else {
		goto lor_lhs_false2070
	}

lor_lhs_false2070:
	v863 = *libc.As[int32](lookahead)
	cmp2071 = v863 == 46
	if cmp2071 {
		goto if_then2094
	} else {
		goto lor_lhs_false2073
	}

lor_lhs_false2073:
	v864 = *libc.As[int32](lookahead)
	cmp2074 = 48 <= v864
	if cmp2074 {
		goto land_lhs_true2076
	} else {
		goto lor_lhs_false2079
	}

land_lhs_true2076:
	v865 = *libc.As[int32](lookahead)
	cmp2077 = v865 <= 57
	if cmp2077 {
		goto if_then2094
	} else {
		goto lor_lhs_false2079
	}

lor_lhs_false2079:
	v866 = *libc.As[int32](lookahead)
	cmp2080 = 65 <= v866
	if cmp2080 {
		goto land_lhs_true2082
	} else {
		goto lor_lhs_false2085
	}

land_lhs_true2082:
	v867 = *libc.As[int32](lookahead)
	cmp2083 = v867 <= 90
	if cmp2083 {
		goto if_then2094
	} else {
		goto lor_lhs_false2085
	}

lor_lhs_false2085:
	v868 = *libc.As[int32](lookahead)
	cmp2086 = v868 == 95
	if cmp2086 {
		goto if_then2094
	} else {
		goto lor_lhs_false2088
	}

lor_lhs_false2088:
	v869 = *libc.As[int32](lookahead)
	cmp2089 = 97 <= v869
	if cmp2089 {
		goto land_lhs_true2091
	} else {
		goto if_end2095
	}

land_lhs_true2091:
	v870 = *libc.As[int32](lookahead)
	cmp2092 = v870 <= 122
	if cmp2092 {
		goto if_then2094
	} else {
		goto if_end2095
	}

if_then2094:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end2095:
	v871 = *libc.As[byte](result)
	loadedv2096 = (v871 & 1) != 0
	*libc.As[bool](retval) = loadedv2096
	goto _return

sw_bb2097:
	*libc.As[byte](result) = 1
	v872 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2098 = libc.Ptr(&libc.As[TSLexer](v872).F1)
	*libc.As[int16](result_symbol2098) = 10
	v873 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2099 = libc.Ptr(&libc.As[TSLexer](v873).F3)
	v874 = *libc.As[unsafe.Pointer](mark_end2099)
	v875 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v874)(v875)
	v876 = *libc.As[int32](lookahead)
	cmp2100 = v876 == 121
	if cmp2100 {
		goto if_then2102
	} else {
		goto if_end2103
	}

if_then2102:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end2103:
	v877 = *libc.As[int32](lookahead)
	cmp2104 = v877 == 45
	if cmp2104 {
		goto if_then2130
	} else {
		goto lor_lhs_false2106
	}

lor_lhs_false2106:
	v878 = *libc.As[int32](lookahead)
	cmp2107 = v878 == 46
	if cmp2107 {
		goto if_then2130
	} else {
		goto lor_lhs_false2109
	}

lor_lhs_false2109:
	v879 = *libc.As[int32](lookahead)
	cmp2110 = 48 <= v879
	if cmp2110 {
		goto land_lhs_true2112
	} else {
		goto lor_lhs_false2115
	}

land_lhs_true2112:
	v880 = *libc.As[int32](lookahead)
	cmp2113 = v880 <= 57
	if cmp2113 {
		goto if_then2130
	} else {
		goto lor_lhs_false2115
	}

lor_lhs_false2115:
	v881 = *libc.As[int32](lookahead)
	cmp2116 = 65 <= v881
	if cmp2116 {
		goto land_lhs_true2118
	} else {
		goto lor_lhs_false2121
	}

land_lhs_true2118:
	v882 = *libc.As[int32](lookahead)
	cmp2119 = v882 <= 90
	if cmp2119 {
		goto if_then2130
	} else {
		goto lor_lhs_false2121
	}

lor_lhs_false2121:
	v883 = *libc.As[int32](lookahead)
	cmp2122 = v883 == 95
	if cmp2122 {
		goto if_then2130
	} else {
		goto lor_lhs_false2124
	}

lor_lhs_false2124:
	v884 = *libc.As[int32](lookahead)
	cmp2125 = 97 <= v884
	if cmp2125 {
		goto land_lhs_true2127
	} else {
		goto if_end2131
	}

land_lhs_true2127:
	v885 = *libc.As[int32](lookahead)
	cmp2128 = v885 <= 122
	if cmp2128 {
		goto if_then2130
	} else {
		goto if_end2131
	}

if_then2130:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end2131:
	v886 = *libc.As[byte](result)
	loadedv2132 = (v886 & 1) != 0
	*libc.As[bool](retval) = loadedv2132
	goto _return

sw_bb2133:
	*libc.As[byte](result) = 1
	v887 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2134 = libc.Ptr(&libc.As[TSLexer](v887).F1)
	*libc.As[int16](result_symbol2134) = 10
	v888 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2135 = libc.Ptr(&libc.As[TSLexer](v888).F3)
	v889 = *libc.As[unsafe.Pointer](mark_end2135)
	v890 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v889)(v890)
	v891 = *libc.As[int32](lookahead)
	cmp2136 = v891 == 45
	if cmp2136 {
		goto if_then2162
	} else {
		goto lor_lhs_false2138
	}

lor_lhs_false2138:
	v892 = *libc.As[int32](lookahead)
	cmp2139 = v892 == 46
	if cmp2139 {
		goto if_then2162
	} else {
		goto lor_lhs_false2141
	}

lor_lhs_false2141:
	v893 = *libc.As[int32](lookahead)
	cmp2142 = 48 <= v893
	if cmp2142 {
		goto land_lhs_true2144
	} else {
		goto lor_lhs_false2147
	}

land_lhs_true2144:
	v894 = *libc.As[int32](lookahead)
	cmp2145 = v894 <= 57
	if cmp2145 {
		goto if_then2162
	} else {
		goto lor_lhs_false2147
	}

lor_lhs_false2147:
	v895 = *libc.As[int32](lookahead)
	cmp2148 = 65 <= v895
	if cmp2148 {
		goto land_lhs_true2150
	} else {
		goto lor_lhs_false2153
	}

land_lhs_true2150:
	v896 = *libc.As[int32](lookahead)
	cmp2151 = v896 <= 90
	if cmp2151 {
		goto if_then2162
	} else {
		goto lor_lhs_false2153
	}

lor_lhs_false2153:
	v897 = *libc.As[int32](lookahead)
	cmp2154 = v897 == 95
	if cmp2154 {
		goto if_then2162
	} else {
		goto lor_lhs_false2156
	}

lor_lhs_false2156:
	v898 = *libc.As[int32](lookahead)
	cmp2157 = 97 <= v898
	if cmp2157 {
		goto land_lhs_true2159
	} else {
		goto if_end2163
	}

land_lhs_true2159:
	v899 = *libc.As[int32](lookahead)
	cmp2160 = v899 <= 122
	if cmp2160 {
		goto if_then2162
	} else {
		goto if_end2163
	}

if_then2162:
	*libc.As[int16](state_addr) = 130
	goto next_state

if_end2163:
	v900 = *libc.As[byte](result)
	loadedv2164 = (v900 & 1) != 0
	*libc.As[bool](retval) = loadedv2164
	goto _return

sw_bb2165:
	*libc.As[byte](result) = 1
	v901 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2166 = libc.Ptr(&libc.As[TSLexer](v901).F1)
	*libc.As[int16](result_symbol2166) = 11
	v902 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2167 = libc.Ptr(&libc.As[TSLexer](v902).F3)
	v903 = *libc.As[unsafe.Pointer](mark_end2167)
	v904 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v903)(v904)
	v905 = *libc.As[byte](result)
	loadedv2168 = (v905 & 1) != 0
	*libc.As[bool](retval) = loadedv2168
	goto _return

sw_bb2169:
	*libc.As[byte](result) = 1
	v906 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2170 = libc.Ptr(&libc.As[TSLexer](v906).F1)
	*libc.As[int16](result_symbol2170) = 12
	v907 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2171 = libc.Ptr(&libc.As[TSLexer](v907).F3)
	v908 = *libc.As[unsafe.Pointer](mark_end2171)
	v909 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v908)(v909)
	v910 = *libc.As[byte](result)
	loadedv2172 = (v910 & 1) != 0
	*libc.As[bool](retval) = loadedv2172
	goto _return

sw_bb2173:
	*libc.As[byte](result) = 1
	v911 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2174 = libc.Ptr(&libc.As[TSLexer](v911).F1)
	*libc.As[int16](result_symbol2174) = 12
	v912 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2175 = libc.Ptr(&libc.As[TSLexer](v912).F3)
	v913 = *libc.As[unsafe.Pointer](mark_end2175)
	v914 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v913)(v914)
	v915 = *libc.As[int32](lookahead)
	cmp2176 = v915 == 10
	if cmp2176 {
		goto if_then2178
	} else {
		goto if_end2179
	}

if_then2178:
	*libc.As[int16](state_addr) = 132
	goto next_state

if_end2179:
	v916 = *libc.As[byte](result)
	loadedv2180 = (v916 & 1) != 0
	*libc.As[bool](retval) = loadedv2180
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v917 = *libc.As[bool](retval)
	return v917
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
