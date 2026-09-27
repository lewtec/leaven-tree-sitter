package grammar_eex

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
}

var tree_sitter_eex_language struct {
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
}
var ts_small_parse_table [280]int16 = [280]int16{6, 31, 1, 5, 33, 1, 6, 37, 1, 15, 13, 1, 25, 35, 2, 7, 8, 28, 2, 19, 20, 2, 39, 1, 0, 41, 7, 1, 2, 3, 4, 10, 11, 14, 2, 43, 1, 0, 45, 7, 1, 2, 3, 4, 10, 11, 14, 2, 47, 1, 0, 49, 7, 1, 2, 3, 4, 10, 11, 14, 2, 51, 1, 0, 53, 7, 1, 2, 3, 4, 10, 11, 14, 2, 55, 1, 0, 57, 7, 1, 2, 3, 4, 10, 11, 14, 2, 59, 1, 0, 61, 7, 1, 2, 3, 4, 10, 11, 14, 2, 63, 1, 0, 65, 7, 1, 2, 3, 4, 10, 11, 14, 3, 69, 1, 15, 12, 1, 25, 67, 3, 5, 7, 8, 4, 72, 1, 5, 76, 1, 15, 12, 1, 25, 74, 2, 7, 8, 3, 78, 1, 5, 80, 1, 15, 25, 1, 25, 3, 82, 1, 5, 84, 1, 15, 14, 1, 25, 3, 86, 1, 12, 88, 1, 13, 21, 1, 26, 3, 80, 1, 15, 90, 1, 5, 25, 1, 25, 3, 90, 1, 5, 92, 1, 15, 22, 1, 25, 3, 94, 1, 5, 96, 1, 15, 17, 1, 25, 3, 98, 1, 12, 100, 1, 13, 16, 1, 26, 3, 102, 1, 12, 105, 1, 13, 21, 1, 26, 3, 80, 1, 15, 107, 1, 5, 25, 1, 25, 3, 107, 1, 5, 109, 1, 15, 24, 1, 25, 3, 80, 1, 15, 111, 1, 5, 25, 1, 25, 3, 67, 1, 5, 113, 1, 15, 25, 1, 25, 2, 116, 1, 5, 118, 1, 9, 2, 120, 1, 5, 122, 1, 9, 1, 124, 1, 5, 1, 126, 1, 0}
var ts_small_parse_table_map [26]int32 = [26]int32{0, 21, 34, 47, 60, 73, 86, 99, 112, 124, 138, 148, 158, 168, 178, 188, 198, 208, 218, 228, 238, 248, 258, 265, 272, 276}
var ts_symbol_names [27]unsafe.Pointer = [27]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27), libc.Ptr(&_str_28)}
var ts_symbol_metadata [27]TSSymbolMetadata = [27]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [27]int16 = [27]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][4]int16 = [1][4]int16{}
var ts_lex_modes [30]TSLexMode = [30]TSLexMode{TSLexMode{}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{2, 0}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{15, 0}, TSLexMode{3, 0}, TSLexMode{3, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{7, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}}
var ts_parse_table struct {
	F0 struct {
		F0 [14]int16
		F1 [13]int16
	}
	F1 [27]int16
	F2 [27]int16
	F3 [27]int16
} = struct {
	F0 struct {
		F0 [14]int16
		F1 [13]int16
	}
	F1 [27]int16
	F2 [27]int16
	F3 [27]int16
}{struct {
	F0 [14]int16
	F1 [13]int16
}{[14]int16{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1}, [13]int16{}}, [27]int16{3, 5, 5, 5, 5, 0, 0, 0, 0, 0, 7, 9, 0, 0, 11, 0, 29, 2, 2, 0, 0, 2, 7, 7, 2, 0, 0}, [27]int16{13, 5, 5, 5, 5, 0, 0, 0, 0, 0, 7, 9, 0, 0, 15, 0, 0, 3, 3, 0, 0, 3, 7, 7, 3, 0, 0}, [27]int16{17, 19, 19, 19, 19, 0, 0, 0, 0, 0, 22, 25, 0, 0, 28, 0, 0, 3, 3, 0, 0, 3, 7, 7, 3, 0, 0}}
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
	F14 TSParseActionEntry
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
	F20 TSParseActionEntry
	F21 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F22 struct {
		F0 anon_2
		F1 [6]byte
	}
	F23 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F46 TSParseActionEntry
	F47 struct {
		F0 anon_2
		F1 [6]byte
	}
	F48 TSParseActionEntry
	F49 struct {
		F0 anon_2
		F1 [6]byte
	}
	F50 TSParseActionEntry
	F51 struct {
		F0 anon_2
		F1 [6]byte
	}
	F52 TSParseActionEntry
	F53 struct {
		F0 anon_2
		F1 [6]byte
	}
	F54 TSParseActionEntry
	F55 struct {
		F0 anon_2
		F1 [6]byte
	}
	F56 TSParseActionEntry
	F57 struct {
		F0 anon_2
		F1 [6]byte
	}
	F58 TSParseActionEntry
	F59 struct {
		F0 anon_2
		F1 [6]byte
	}
	F60 TSParseActionEntry
	F61 struct {
		F0 anon_2
		F1 [6]byte
	}
	F62 TSParseActionEntry
	F63 struct {
		F0 anon_2
		F1 [6]byte
	}
	F64 TSParseActionEntry
	F65 struct {
		F0 anon_2
		F1 [6]byte
	}
	F66 TSParseActionEntry
	F67 struct {
		F0 anon_2
		F1 [6]byte
	}
	F68 TSParseActionEntry
	F69 struct {
		F0 anon_2
		F1 [6]byte
	}
	F70 TSParseActionEntry
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
	F81 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F82 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F91 TSParseActionEntry
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
	F103 TSParseActionEntry
	F104 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F110 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
	F117 TSParseActionEntry
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
	F121 TSParseActionEntry
	F122 struct {
		F0 anon_2
		F1 [6]byte
	}
	F123 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
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
	F14 TSParseActionEntry
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
	F20 TSParseActionEntry
	F21 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F22 struct {
		F0 anon_2
		F1 [6]byte
	}
	F23 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F46 TSParseActionEntry
	F47 struct {
		F0 anon_2
		F1 [6]byte
	}
	F48 TSParseActionEntry
	F49 struct {
		F0 anon_2
		F1 [6]byte
	}
	F50 TSParseActionEntry
	F51 struct {
		F0 anon_2
		F1 [6]byte
	}
	F52 TSParseActionEntry
	F53 struct {
		F0 anon_2
		F1 [6]byte
	}
	F54 TSParseActionEntry
	F55 struct {
		F0 anon_2
		F1 [6]byte
	}
	F56 TSParseActionEntry
	F57 struct {
		F0 anon_2
		F1 [6]byte
	}
	F58 TSParseActionEntry
	F59 struct {
		F0 anon_2
		F1 [6]byte
	}
	F60 TSParseActionEntry
	F61 struct {
		F0 anon_2
		F1 [6]byte
	}
	F62 TSParseActionEntry
	F63 struct {
		F0 anon_2
		F1 [6]byte
	}
	F64 TSParseActionEntry
	F65 struct {
		F0 anon_2
		F1 [6]byte
	}
	F66 TSParseActionEntry
	F67 struct {
		F0 anon_2
		F1 [6]byte
	}
	F68 TSParseActionEntry
	F69 struct {
		F0 anon_2
		F1 [6]byte
	}
	F70 TSParseActionEntry
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
	F81 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F82 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F91 TSParseActionEntry
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
	F103 TSParseActionEntry
	F104 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
	F110 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
	F117 TSParseActionEntry
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
	F121 TSParseActionEntry
	F122 struct {
		F0 anon_2
		F1 [6]byte
	}
	F123 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 16, 0, 0}}}, struct {
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
}{0, 0, 4, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 16, 0, 0}}}, struct {
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
}{0, 0, 3, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{0, 0, 8, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 26, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 13, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 23, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 23, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 21, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 21, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 18, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 18, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 18, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 18, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 23, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 23, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 25, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 25, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 20, 0, 0}}}, struct {
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
}{0, 0, 12, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 6, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 25, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 21, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 5, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 19, 0, 0}}}, struct {
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
}{0, 0, 22, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 19, 0, 0}}}, struct {
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
}{0, 0, 10, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 26, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 26, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 19, 0, 0}}}, struct {
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
}{0, 0, 24, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 19, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 25, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 19, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 19, 0, 0}}}, struct {
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
}{0, 0, 9, 0, 0}, [2]byte{}}}, struct {
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
var _str_3 [3]byte = [3]byte{60, 37, 0}
var _str_4 [4]byte = [4]byte{60, 37, 61, 0}
var _str_5 [4]byte = [4]byte{60, 37, 37, 0}
var _str_6 [5]byte = [5]byte{60, 37, 37, 61, 0}
var _str_7 [3]byte = [3]byte{37, 62, 0}
var _str_8 [26]byte = [26]byte{112, 97, 114, 116, 105, 97, 108, 95, 101, 120, 112, 114, 101, 115, 115, 105, 111, 110, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_9 [3]byte = [3]byte{100, 111, 0}
var _str_10 [3]byte = [3]byte{45, 62, 0}
var _str_11 [2]byte = [2]byte{35, 0}
var _str_12 [4]byte = [4]byte{60, 37, 35, 0}
var _str_13 [6]byte = [6]byte{60, 37, 33, 45, 45, 0}
var _str_14 [21]byte = [21]byte{95, 98, 97, 110, 103, 95, 99, 111, 109, 109, 101, 110, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_15 [5]byte = [5]byte{45, 45, 37, 62, 0}
var _str_16 [5]byte = [5]byte{116, 101, 120, 116, 0}
var _str_17 [6]byte = [6]byte{95, 99, 111, 100, 101, 0}
var _str_18 [9]byte = [9]byte{102, 114, 97, 103, 109, 101, 110, 116, 0}
var _str_19 [6]byte = [6]byte{95, 110, 111, 100, 101, 0}
var _str_20 [10]byte = [10]byte{100, 105, 114, 101, 99, 116, 105, 118, 101, 0}
var _str_21 [19]byte = [19]byte{112, 97, 114, 116, 105, 97, 108, 95, 101, 120, 112, 114, 101, 115, 115, 105, 111, 110, 0}
var _str_22 [11]byte = [11]byte{101, 120, 112, 114, 101, 115, 115, 105, 111, 110, 0}
var _str_23 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_24 [14]byte = [14]byte{95, 104, 97, 115, 104, 95, 99, 111, 109, 109, 101, 110, 116, 0}
var _str_25 [14]byte = [14]byte{95, 98, 97, 110, 103, 95, 99, 111, 109, 109, 101, 110, 116, 0}
var _str_26 [17]byte = [17]byte{102, 114, 97, 103, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_27 [27]byte = [27]byte{112, 97, 114, 116, 105, 97, 108, 95, 101, 120, 112, 114, 101, 115, 115, 105, 111, 110, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_28 [22]byte = [22]byte{95, 98, 97, 110, 103, 95, 99, 111, 109, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}

func init() {
	tree_sitter_eex_language = struct {
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
	}{13, 27, 0, 16, 0, 30, 4, 1, 0, 4, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}}
}
func tree_sitter_eex() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_eex_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp27, cmp29, cmp32, cmp35, loadedv39, cmp41, loadedv45, cmp47, cmp51, cmp55, cmp59, cmp63, cmp66, cmp69, cmp72, cmp76, loadedv80, cmp82, cmp86, cmp90, cmp94, cmp97, cmp100, cmp103, cmp107, loadedv111, cmp113, cmp117, cmp120, cmp123, cmp126, cmp130, loadedv134, cmp136, loadedv140, cmp142, loadedv146, cmp148, cmp152, cmp155, cmp158, cmp161, cmp165, loadedv169, cmp171, cmp175, loadedv179, cmp181, loadedv185, cmp187, loadedv191, cmp193, loadedv197, cmp199, loadedv203, cmp205, loadedv209, cmp211, loadedv215, loadedv217, cmp220, cmp224, cmp227, cmp230, cmp233, cmp237, loadedv241, loadedv243, cmp247, cmp251, cmp255, cmp259, loadedv263, loadedv267, cmp271, loadedv275, loadedv279, loadedv283, cmp287, cmp290, cmp293, loadedv297, cmp301, cmp304, cmp307, cmp311, cmp313, cmp316, cmp319, cmp322, cmp325, loadedv329, loadedv333, cmp337, cmp340, cmp343, cmp346, cmp349, cmp352, loadedv356, loadedv360, cmp364, cmp367, cmp370, cmp373, cmp376, cmp379, loadedv383, loadedv387, loadedv391, loadedv395, cmp399, cmp403, cmp406, cmp409, cmp412, cmp416, loadedv420, cmp424, loadedv428, cmp432, cmp435, loadedv439, loadedv443, cmp447, loadedv451, cmp455, cmp459, cmp462, cmp465, cmp468, cmp472, loadedv476, cmp480, cmp483, loadedv487, cmp491, cmp495, cmp499, cmp503, cmp507, cmp510, cmp513, cmp516, cmp520, loadedv524, cmp528, cmp532, cmp536, cmp540, cmp543, cmp546, cmp549, cmp553, loadedv557, cmp561, cmp565, cmp568, cmp571, cmp574, cmp578, loadedv582, cmp586, loadedv590, cmp594, cmp598, cmp601, cmp604, cmp607, cmp610, cmp613, loadedv617, cmp621, cmp625, cmp628, cmp631, cmp634, cmp637, cmp640, loadedv644, cmp648, cmp652, cmp655, cmp658, cmp661, cmp664, cmp667, loadedv671, cmp675, cmp679, cmp682, cmp685, cmp688, cmp691, cmp694, loadedv698, cmp702, cmp705, cmp708, cmp711, cmp714, cmp717, loadedv721, v344 bool
	var retval unsafe.Pointer
	var v9 int16
	var state_addr, result_symbol, result_symbol245, result_symbol265, result_symbol269, result_symbol277, result_symbol281, result_symbol285, result_symbol299, result_symbol331, result_symbol335, result_symbol358, result_symbol362, result_symbol385, result_symbol389, result_symbol393, result_symbol397, result_symbol422, result_symbol430, result_symbol441, result_symbol445, result_symbol453, result_symbol478, result_symbol489, result_symbol526, result_symbol559, result_symbol584, result_symbol592, result_symbol619, result_symbol646, result_symbol673, result_symbol700 unsafe.Pointer
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v22, v24, v25, v26, v27, v28, v29, v30, v31, v32, v34, v35, v36, v37, v38, v39, v40, v41, v43, v44, v45, v46, v47, v48, v50, v52, v54, v55, v56, v57, v58, v59, v61, v62, v64, v66, v68, v70, v72, v74, v77, v78, v79, v80, v81, v82, v93, v94, v95, v96, v107, v123, v124, v125, v131, v132, v133, v134, v135, v136, v137, v138, v139, v150, v151, v152, v153, v154, v155, v166, v167, v168, v169, v170, v171, v192, v193, v194, v195, v196, v197, v203, v209, v210, v221, v227, v228, v229, v230, v231, v232, v238, v239, v245, v246, v247, v248, v249, v250, v251, v252, v253, v259, v260, v261, v262, v263, v264, v265, v266, v272, v273, v274, v275, v276, v277, v283, v289, v290, v291, v292, v293, v294, v295, v301, v302, v303, v304, v305, v306, v307, v313, v314, v315, v316, v317, v318, v319, v325, v326, v327, v328, v329, v330, v331, v337, v338, v339, v340, v341, v342 int32
	var lookahead, lookahead1 unsafe.Pointer
	var v3, storedv, v10, v21, v23, v33, v42, v49, v51, v53, v60, v63, v65, v67, v69, v71, v73, v75, v76, v83, v88, v97, v102, v108, v113, v118, v126, v140, v145, v156, v161, v172, v177, v182, v187, v198, v204, v211, v216, v222, v233, v240, v254, v267, v278, v284, v296, v308, v320, v332, v343 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v84, v85, v86, v87, v89, v90, v91, v92, v98, v99, v100, v101, v103, v104, v105, v106, v109, v110, v111, v112, v114, v115, v116, v117, v119, v120, v121, v122, v127, v128, v129, v130, v141, v142, v143, v144, v146, v147, v148, v149, v157, v158, v159, v160, v162, v163, v164, v165, v173, v174, v175, v176, v178, v179, v180, v181, v183, v184, v185, v186, v188, v189, v190, v191, v199, v200, v201, v202, v205, v206, v207, v208, v212, v213, v214, v215, v217, v218, v219, v220, v223, v224, v225, v226, v234, v235, v236, v237, v241, v242, v243, v244, v255, v256, v257, v258, v268, v269, v270, v271, v279, v280, v281, v282, v285, v286, v287, v288, v297, v298, v299, v300, v309, v310, v311, v312, v321, v322, v323, v324, v333, v334, v335, v336 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end246, mark_end266, mark_end270, mark_end278, mark_end282, mark_end286, mark_end300, mark_end332, mark_end336, mark_end359, mark_end363, mark_end386, mark_end390, mark_end394, mark_end398, mark_end423, mark_end431, mark_end442, mark_end446, mark_end454, mark_end479, mark_end490, mark_end527, mark_end560, mark_end585, mark_end593, mark_end620, mark_end647, mark_end674, mark_end701 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp27, v18, cmp29, v19, cmp32, v20, cmp35, v21, loadedv39, v22, cmp41, v23, loadedv45, v24, cmp47, v25, cmp51, v26, cmp55, v27, cmp59, v28, cmp63, v29, cmp66, v30, cmp69, v31, cmp72, v32, cmp76, v33, loadedv80, v34, cmp82, v35, cmp86, v36, cmp90, v37, cmp94, v38, cmp97, v39, cmp100, v40, cmp103, v41, cmp107, v42, loadedv111, v43, cmp113, v44, cmp117, v45, cmp120, v46, cmp123, v47, cmp126, v48, cmp130, v49, loadedv134, v50, cmp136, v51, loadedv140, v52, cmp142, v53, loadedv146, v54, cmp148, v55, cmp152, v56, cmp155, v57, cmp158, v58, cmp161, v59, cmp165, v60, loadedv169, v61, cmp171, v62, cmp175, v63, loadedv179, v64, cmp181, v65, loadedv185, v66, cmp187, v67, loadedv191, v68, cmp193, v69, loadedv197, v70, cmp199, v71, loadedv203, v72, cmp205, v73, loadedv209, v74, cmp211, v75, loadedv215, v76, loadedv217, v77, cmp220, v78, cmp224, v79, cmp227, v80, cmp230, v81, cmp233, v82, cmp237, v83, loadedv241, v84, result_symbol, v85, mark_end, v86, v87, v88, loadedv243, v89, result_symbol245, v90, mark_end246, v91, v92, v93, cmp247, v94, cmp251, v95, cmp255, v96, cmp259, v97, loadedv263, v98, result_symbol265, v99, mark_end266, v100, v101, v102, loadedv267, v103, result_symbol269, v104, mark_end270, v105, v106, v107, cmp271, v108, loadedv275, v109, result_symbol277, v110, mark_end278, v111, v112, v113, loadedv279, v114, result_symbol281, v115, mark_end282, v116, v117, v118, loadedv283, v119, result_symbol285, v120, mark_end286, v121, v122, v123, cmp287, v124, cmp290, v125, cmp293, v126, loadedv297, v127, result_symbol299, v128, mark_end300, v129, v130, v131, cmp301, v132, cmp304, v133, cmp307, v134, cmp311, v135, cmp313, v136, cmp316, v137, cmp319, v138, cmp322, v139, cmp325, v140, loadedv329, v141, result_symbol331, v142, mark_end332, v143, v144, v145, loadedv333, v146, result_symbol335, v147, mark_end336, v148, v149, v150, cmp337, v151, cmp340, v152, cmp343, v153, cmp346, v154, cmp349, v155, cmp352, v156, loadedv356, v157, result_symbol358, v158, mark_end359, v159, v160, v161, loadedv360, v162, result_symbol362, v163, mark_end363, v164, v165, v166, cmp364, v167, cmp367, v168, cmp370, v169, cmp373, v170, cmp376, v171, cmp379, v172, loadedv383, v173, result_symbol385, v174, mark_end386, v175, v176, v177, loadedv387, v178, result_symbol389, v179, mark_end390, v180, v181, v182, loadedv391, v183, result_symbol393, v184, mark_end394, v185, v186, v187, loadedv395, v188, result_symbol397, v189, mark_end398, v190, v191, v192, cmp399, v193, cmp403, v194, cmp406, v195, cmp409, v196, cmp412, v197, cmp416, v198, loadedv420, v199, result_symbol422, v200, mark_end423, v201, v202, v203, cmp424, v204, loadedv428, v205, result_symbol430, v206, mark_end431, v207, v208, v209, cmp432, v210, cmp435, v211, loadedv439, v212, result_symbol441, v213, mark_end442, v214, v215, v216, loadedv443, v217, result_symbol445, v218, mark_end446, v219, v220, v221, cmp447, v222, loadedv451, v223, result_symbol453, v224, mark_end454, v225, v226, v227, cmp455, v228, cmp459, v229, cmp462, v230, cmp465, v231, cmp468, v232, cmp472, v233, loadedv476, v234, result_symbol478, v235, mark_end479, v236, v237, v238, cmp480, v239, cmp483, v240, loadedv487, v241, result_symbol489, v242, mark_end490, v243, v244, v245, cmp491, v246, cmp495, v247, cmp499, v248, cmp503, v249, cmp507, v250, cmp510, v251, cmp513, v252, cmp516, v253, cmp520, v254, loadedv524, v255, result_symbol526, v256, mark_end527, v257, v258, v259, cmp528, v260, cmp532, v261, cmp536, v262, cmp540, v263, cmp543, v264, cmp546, v265, cmp549, v266, cmp553, v267, loadedv557, v268, result_symbol559, v269, mark_end560, v270, v271, v272, cmp561, v273, cmp565, v274, cmp568, v275, cmp571, v276, cmp574, v277, cmp578, v278, loadedv582, v279, result_symbol584, v280, mark_end585, v281, v282, v283, cmp586, v284, loadedv590, v285, result_symbol592, v286, mark_end593, v287, v288, v289, cmp594, v290, cmp598, v291, cmp601, v292, cmp604, v293, cmp607, v294, cmp610, v295, cmp613, v296, loadedv617, v297, result_symbol619, v298, mark_end620, v299, v300, v301, cmp621, v302, cmp625, v303, cmp628, v304, cmp631, v305, cmp634, v306, cmp637, v307, cmp640, v308, loadedv644, v309, result_symbol646, v310, mark_end647, v311, v312, v313, cmp648, v314, cmp652, v315, cmp655, v316, cmp658, v317, cmp661, v318, cmp664, v319, cmp667, v320, loadedv671, v321, result_symbol673, v322, mark_end674, v323, v324, v325, cmp675, v326, cmp679, v327, cmp682, v328, cmp685, v329, cmp688, v330, cmp691, v331, cmp694, v332, loadedv698, v333, result_symbol700, v334, mark_end701, v335, v336, v337, cmp702, v338, cmp705, v339, cmp708, v340, cmp711, v341, cmp714, v342, cmp717, v343, loadedv721, v344

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
		goto sw_bb40
	case 2:
		goto sw_bb46
	case 3:
		goto sw_bb81
	case 4:
		goto sw_bb112
	case 5:
		goto sw_bb135
	case 6:
		goto sw_bb141
	case 7:
		goto sw_bb147
	case 8:
		goto sw_bb170
	case 9:
		goto sw_bb180
	case 10:
		goto sw_bb186
	case 11:
		goto sw_bb192
	case 12:
		goto sw_bb198
	case 13:
		goto sw_bb204
	case 14:
		goto sw_bb210
	case 15:
		goto sw_bb216
	case 16:
		goto sw_bb242
	case 17:
		goto sw_bb244
	case 18:
		goto sw_bb264
	case 19:
		goto sw_bb268
	case 20:
		goto sw_bb276
	case 21:
		goto sw_bb280
	case 22:
		goto sw_bb284
	case 23:
		goto sw_bb298
	case 24:
		goto sw_bb330
	case 25:
		goto sw_bb334
	case 26:
		goto sw_bb357
	case 27:
		goto sw_bb361
	case 28:
		goto sw_bb384
	case 29:
		goto sw_bb388
	case 30:
		goto sw_bb392
	case 31:
		goto sw_bb396
	case 32:
		goto sw_bb421
	case 33:
		goto sw_bb429
	case 34:
		goto sw_bb440
	case 35:
		goto sw_bb444
	case 36:
		goto sw_bb452
	case 37:
		goto sw_bb477
	case 38:
		goto sw_bb488
	case 39:
		goto sw_bb525
	case 40:
		goto sw_bb558
	case 41:
		goto sw_bb583
	case 42:
		goto sw_bb591
	case 43:
		goto sw_bb618
	case 44:
		goto sw_bb645
	case 45:
		goto sw_bb672
	case 46:
		goto sw_bb699
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
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end:
	v11 = *libc.As[int32](lookahead)
	cmp = v11 == 35
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end6:
	v12 = *libc.As[int32](lookahead)
	cmp7 = v12 == 37
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end10:
	v13 = *libc.As[int32](lookahead)
	cmp11 = v13 == 45
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end14:
	v14 = *libc.As[int32](lookahead)
	cmp15 = v14 == 60
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end18:
	v15 = *libc.As[int32](lookahead)
	cmp19 = v15 == 100
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end22:
	v16 = *libc.As[int32](lookahead)
	cmp23 = v16 == 101
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end26:
	v17 = *libc.As[int32](lookahead)
	cmp27 = v17 == 9
	if cmp27 {
		goto if_then37
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v18 = *libc.As[int32](lookahead)
	cmp29 = v18 == 10
	if cmp29 {
		goto if_then37
	} else {
		goto lor_lhs_false31
	}

lor_lhs_false31:
	v19 = *libc.As[int32](lookahead)
	cmp32 = v19 == 13
	if cmp32 {
		goto if_then37
	} else {
		goto lor_lhs_false34
	}

lor_lhs_false34:
	v20 = *libc.As[int32](lookahead)
	cmp35 = v20 == 32
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
	v21 = *libc.As[byte](result)
	loadedv39 = (v21 & 1) != 0
	*libc.As[bool](retval) = loadedv39
	goto _return

sw_bb40:
	v22 = *libc.As[int32](lookahead)
	cmp41 = v22 == 37
	if cmp41 {
		goto if_then43
	} else {
		goto if_end44
	}

if_then43:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end44:
	v23 = *libc.As[byte](result)
	loadedv45 = (v23 & 1) != 0
	*libc.As[bool](retval) = loadedv45
	goto _return

sw_bb46:
	v24 = *libc.As[int32](lookahead)
	cmp47 = v24 == 37
	if cmp47 {
		goto if_then49
	} else {
		goto if_end50
	}

if_then49:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end50:
	v25 = *libc.As[int32](lookahead)
	cmp51 = v25 == 45
	if cmp51 {
		goto if_then53
	} else {
		goto if_end54
	}

if_then53:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end54:
	v26 = *libc.As[int32](lookahead)
	cmp55 = v26 == 100
	if cmp55 {
		goto if_then57
	} else {
		goto if_end58
	}

if_then57:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end58:
	v27 = *libc.As[int32](lookahead)
	cmp59 = v27 == 101
	if cmp59 {
		goto if_then61
	} else {
		goto if_end62
	}

if_then61:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end62:
	v28 = *libc.As[int32](lookahead)
	cmp63 = v28 == 9
	if cmp63 {
		goto if_then74
	} else {
		goto lor_lhs_false65
	}

lor_lhs_false65:
	v29 = *libc.As[int32](lookahead)
	cmp66 = v29 == 10
	if cmp66 {
		goto if_then74
	} else {
		goto lor_lhs_false68
	}

lor_lhs_false68:
	v30 = *libc.As[int32](lookahead)
	cmp69 = v30 == 13
	if cmp69 {
		goto if_then74
	} else {
		goto lor_lhs_false71
	}

lor_lhs_false71:
	v31 = *libc.As[int32](lookahead)
	cmp72 = v31 == 32
	if cmp72 {
		goto if_then74
	} else {
		goto if_end75
	}

if_then74:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end75:
	v32 = *libc.As[int32](lookahead)
	cmp76 = v32 != 0
	if cmp76 {
		goto if_then78
	} else {
		goto if_end79
	}

if_then78:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end79:
	v33 = *libc.As[byte](result)
	loadedv80 = (v33 & 1) != 0
	*libc.As[bool](retval) = loadedv80
	goto _return

sw_bb81:
	v34 = *libc.As[int32](lookahead)
	cmp82 = v34 == 37
	if cmp82 {
		goto if_then84
	} else {
		goto if_end85
	}

if_then84:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end85:
	v35 = *libc.As[int32](lookahead)
	cmp86 = v35 == 45
	if cmp86 {
		goto if_then88
	} else {
		goto if_end89
	}

if_then88:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end89:
	v36 = *libc.As[int32](lookahead)
	cmp90 = v36 == 100
	if cmp90 {
		goto if_then92
	} else {
		goto if_end93
	}

if_then92:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end93:
	v37 = *libc.As[int32](lookahead)
	cmp94 = v37 == 9
	if cmp94 {
		goto if_then105
	} else {
		goto lor_lhs_false96
	}

lor_lhs_false96:
	v38 = *libc.As[int32](lookahead)
	cmp97 = v38 == 10
	if cmp97 {
		goto if_then105
	} else {
		goto lor_lhs_false99
	}

lor_lhs_false99:
	v39 = *libc.As[int32](lookahead)
	cmp100 = v39 == 13
	if cmp100 {
		goto if_then105
	} else {
		goto lor_lhs_false102
	}

lor_lhs_false102:
	v40 = *libc.As[int32](lookahead)
	cmp103 = v40 == 32
	if cmp103 {
		goto if_then105
	} else {
		goto if_end106
	}

if_then105:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end106:
	v41 = *libc.As[int32](lookahead)
	cmp107 = v41 != 0
	if cmp107 {
		goto if_then109
	} else {
		goto if_end110
	}

if_then109:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end110:
	v42 = *libc.As[byte](result)
	loadedv111 = (v42 & 1) != 0
	*libc.As[bool](retval) = loadedv111
	goto _return

sw_bb112:
	v43 = *libc.As[int32](lookahead)
	cmp113 = v43 == 37
	if cmp113 {
		goto if_then115
	} else {
		goto if_end116
	}

if_then115:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end116:
	v44 = *libc.As[int32](lookahead)
	cmp117 = v44 == 9
	if cmp117 {
		goto if_then128
	} else {
		goto lor_lhs_false119
	}

lor_lhs_false119:
	v45 = *libc.As[int32](lookahead)
	cmp120 = v45 == 10
	if cmp120 {
		goto if_then128
	} else {
		goto lor_lhs_false122
	}

lor_lhs_false122:
	v46 = *libc.As[int32](lookahead)
	cmp123 = v46 == 13
	if cmp123 {
		goto if_then128
	} else {
		goto lor_lhs_false125
	}

lor_lhs_false125:
	v47 = *libc.As[int32](lookahead)
	cmp126 = v47 == 32
	if cmp126 {
		goto if_then128
	} else {
		goto if_end129
	}

if_then128:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end129:
	v48 = *libc.As[int32](lookahead)
	cmp130 = v48 != 0
	if cmp130 {
		goto if_then132
	} else {
		goto if_end133
	}

if_then132:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end133:
	v49 = *libc.As[byte](result)
	loadedv134 = (v49 & 1) != 0
	*libc.As[bool](retval) = loadedv134
	goto _return

sw_bb135:
	v50 = *libc.As[int32](lookahead)
	cmp136 = v50 == 37
	if cmp136 {
		goto if_then138
	} else {
		goto if_end139
	}

if_then138:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end139:
	v51 = *libc.As[byte](result)
	loadedv140 = (v51 & 1) != 0
	*libc.As[bool](retval) = loadedv140
	goto _return

sw_bb141:
	v52 = *libc.As[int32](lookahead)
	cmp142 = v52 == 45
	if cmp142 {
		goto if_then144
	} else {
		goto if_end145
	}

if_then144:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end145:
	v53 = *libc.As[byte](result)
	loadedv146 = (v53 & 1) != 0
	*libc.As[bool](retval) = loadedv146
	goto _return

sw_bb147:
	v54 = *libc.As[int32](lookahead)
	cmp148 = v54 == 45
	if cmp148 {
		goto if_then150
	} else {
		goto if_end151
	}

if_then150:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end151:
	v55 = *libc.As[int32](lookahead)
	cmp152 = v55 == 9
	if cmp152 {
		goto if_then163
	} else {
		goto lor_lhs_false154
	}

lor_lhs_false154:
	v56 = *libc.As[int32](lookahead)
	cmp155 = v56 == 10
	if cmp155 {
		goto if_then163
	} else {
		goto lor_lhs_false157
	}

lor_lhs_false157:
	v57 = *libc.As[int32](lookahead)
	cmp158 = v57 == 13
	if cmp158 {
		goto if_then163
	} else {
		goto lor_lhs_false160
	}

lor_lhs_false160:
	v58 = *libc.As[int32](lookahead)
	cmp161 = v58 == 32
	if cmp161 {
		goto if_then163
	} else {
		goto if_end164
	}

if_then163:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end164:
	v59 = *libc.As[int32](lookahead)
	cmp165 = v59 != 0
	if cmp165 {
		goto if_then167
	} else {
		goto if_end168
	}

if_then167:
	*libc.As[int16](state_addr) = 33
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
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end174:
	v62 = *libc.As[int32](lookahead)
	cmp175 = v62 == 62
	if cmp175 {
		goto if_then177
	} else {
		goto if_end178
	}

if_then177:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end178:
	v63 = *libc.As[byte](result)
	loadedv179 = (v63 & 1) != 0
	*libc.As[bool](retval) = loadedv179
	goto _return

sw_bb180:
	v64 = *libc.As[int32](lookahead)
	cmp181 = v64 == 45
	if cmp181 {
		goto if_then183
	} else {
		goto if_end184
	}

if_then183:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end184:
	v65 = *libc.As[byte](result)
	loadedv185 = (v65 & 1) != 0
	*libc.As[bool](retval) = loadedv185
	goto _return

sw_bb186:
	v66 = *libc.As[int32](lookahead)
	cmp187 = v66 == 62
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end190:
	v67 = *libc.As[byte](result)
	loadedv191 = (v67 & 1) != 0
	*libc.As[bool](retval) = loadedv191
	goto _return

sw_bb192:
	v68 = *libc.As[int32](lookahead)
	cmp193 = v68 == 62
	if cmp193 {
		goto if_then195
	} else {
		goto if_end196
	}

if_then195:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end196:
	v69 = *libc.As[byte](result)
	loadedv197 = (v69 & 1) != 0
	*libc.As[bool](retval) = loadedv197
	goto _return

sw_bb198:
	v70 = *libc.As[int32](lookahead)
	cmp199 = v70 == 100
	if cmp199 {
		goto if_then201
	} else {
		goto if_end202
	}

if_then201:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end202:
	v71 = *libc.As[byte](result)
	loadedv203 = (v71 & 1) != 0
	*libc.As[bool](retval) = loadedv203
	goto _return

sw_bb204:
	v72 = *libc.As[int32](lookahead)
	cmp205 = v72 == 110
	if cmp205 {
		goto if_then207
	} else {
		goto if_end208
	}

if_then207:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end208:
	v73 = *libc.As[byte](result)
	loadedv209 = (v73 & 1) != 0
	*libc.As[bool](retval) = loadedv209
	goto _return

sw_bb210:
	v74 = *libc.As[int32](lookahead)
	cmp211 = v74 == 111
	if cmp211 {
		goto if_then213
	} else {
		goto if_end214
	}

if_then213:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end214:
	v75 = *libc.As[byte](result)
	loadedv215 = (v75 & 1) != 0
	*libc.As[bool](retval) = loadedv215
	goto _return

sw_bb216:
	v76 = *libc.As[byte](eof)
	loadedv217 = (v76 & 1) != 0
	if loadedv217 {
		goto if_then218
	} else {
		goto if_end219
	}

if_then218:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end219:
	v77 = *libc.As[int32](lookahead)
	cmp220 = v77 == 60
	if cmp220 {
		goto if_then222
	} else {
		goto if_end223
	}

if_then222:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end223:
	v78 = *libc.As[int32](lookahead)
	cmp224 = v78 == 9
	if cmp224 {
		goto if_then235
	} else {
		goto lor_lhs_false226
	}

lor_lhs_false226:
	v79 = *libc.As[int32](lookahead)
	cmp227 = v79 == 10
	if cmp227 {
		goto if_then235
	} else {
		goto lor_lhs_false229
	}

lor_lhs_false229:
	v80 = *libc.As[int32](lookahead)
	cmp230 = v80 == 13
	if cmp230 {
		goto if_then235
	} else {
		goto lor_lhs_false232
	}

lor_lhs_false232:
	v81 = *libc.As[int32](lookahead)
	cmp233 = v81 == 32
	if cmp233 {
		goto if_then235
	} else {
		goto if_end236
	}

if_then235:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end236:
	v82 = *libc.As[int32](lookahead)
	cmp237 = v82 != 0
	if cmp237 {
		goto if_then239
	} else {
		goto if_end240
	}

if_then239:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end240:
	v83 = *libc.As[byte](result)
	loadedv241 = (v83 & 1) != 0
	*libc.As[bool](retval) = loadedv241
	goto _return

sw_bb242:
	*libc.As[byte](result) = 1
	v84 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v84).F1)
	*libc.As[int16](result_symbol) = 0
	v85 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v85).F3)
	v86 = *libc.As[unsafe.Pointer](mark_end)
	v87 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v86)(v87)
	v88 = *libc.As[byte](result)
	loadedv243 = (v88 & 1) != 0
	*libc.As[bool](retval) = loadedv243
	goto _return

sw_bb244:
	*libc.As[byte](result) = 1
	v89 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol245 = libc.Ptr(&libc.As[TSLexer](v89).F1)
	*libc.As[int16](result_symbol245) = 1
	v90 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end246 = libc.Ptr(&libc.As[TSLexer](v90).F3)
	v91 = *libc.As[unsafe.Pointer](mark_end246)
	v92 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v91)(v92)
	v93 = *libc.As[int32](lookahead)
	cmp247 = v93 == 33
	if cmp247 {
		goto if_then249
	} else {
		goto if_end250
	}

if_then249:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end250:
	v94 = *libc.As[int32](lookahead)
	cmp251 = v94 == 35
	if cmp251 {
		goto if_then253
	} else {
		goto if_end254
	}

if_then253:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end254:
	v95 = *libc.As[int32](lookahead)
	cmp255 = v95 == 37
	if cmp255 {
		goto if_then257
	} else {
		goto if_end258
	}

if_then257:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end258:
	v96 = *libc.As[int32](lookahead)
	cmp259 = v96 == 61
	if cmp259 {
		goto if_then261
	} else {
		goto if_end262
	}

if_then261:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end262:
	v97 = *libc.As[byte](result)
	loadedv263 = (v97 & 1) != 0
	*libc.As[bool](retval) = loadedv263
	goto _return

sw_bb264:
	*libc.As[byte](result) = 1
	v98 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol265 = libc.Ptr(&libc.As[TSLexer](v98).F1)
	*libc.As[int16](result_symbol265) = 2
	v99 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end266 = libc.Ptr(&libc.As[TSLexer](v99).F3)
	v100 = *libc.As[unsafe.Pointer](mark_end266)
	v101 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v100)(v101)
	v102 = *libc.As[byte](result)
	loadedv267 = (v102 & 1) != 0
	*libc.As[bool](retval) = loadedv267
	goto _return

sw_bb268:
	*libc.As[byte](result) = 1
	v103 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol269 = libc.Ptr(&libc.As[TSLexer](v103).F1)
	*libc.As[int16](result_symbol269) = 3
	v104 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end270 = libc.Ptr(&libc.As[TSLexer](v104).F3)
	v105 = *libc.As[unsafe.Pointer](mark_end270)
	v106 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v105)(v106)
	v107 = *libc.As[int32](lookahead)
	cmp271 = v107 == 61
	if cmp271 {
		goto if_then273
	} else {
		goto if_end274
	}

if_then273:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end274:
	v108 = *libc.As[byte](result)
	loadedv275 = (v108 & 1) != 0
	*libc.As[bool](retval) = loadedv275
	goto _return

sw_bb276:
	*libc.As[byte](result) = 1
	v109 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol277 = libc.Ptr(&libc.As[TSLexer](v109).F1)
	*libc.As[int16](result_symbol277) = 4
	v110 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end278 = libc.Ptr(&libc.As[TSLexer](v110).F3)
	v111 = *libc.As[unsafe.Pointer](mark_end278)
	v112 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v111)(v112)
	v113 = *libc.As[byte](result)
	loadedv279 = (v113 & 1) != 0
	*libc.As[bool](retval) = loadedv279
	goto _return

sw_bb280:
	*libc.As[byte](result) = 1
	v114 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol281 = libc.Ptr(&libc.As[TSLexer](v114).F1)
	*libc.As[int16](result_symbol281) = 5
	v115 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end282 = libc.Ptr(&libc.As[TSLexer](v115).F3)
	v116 = *libc.As[unsafe.Pointer](mark_end282)
	v117 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v116)(v117)
	v118 = *libc.As[byte](result)
	loadedv283 = (v118 & 1) != 0
	*libc.As[bool](retval) = loadedv283
	goto _return

sw_bb284:
	*libc.As[byte](result) = 1
	v119 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol285 = libc.Ptr(&libc.As[TSLexer](v119).F1)
	*libc.As[int16](result_symbol285) = 6
	v120 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end286 = libc.Ptr(&libc.As[TSLexer](v120).F3)
	v121 = *libc.As[unsafe.Pointer](mark_end286)
	v122 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v121)(v122)
	v123 = *libc.As[int32](lookahead)
	cmp287 = v123 == 41
	if cmp287 {
		goto if_then295
	} else {
		goto lor_lhs_false289
	}

lor_lhs_false289:
	v124 = *libc.As[int32](lookahead)
	cmp290 = v124 == 93
	if cmp290 {
		goto if_then295
	} else {
		goto lor_lhs_false292
	}

lor_lhs_false292:
	v125 = *libc.As[int32](lookahead)
	cmp293 = v125 == 125
	if cmp293 {
		goto if_then295
	} else {
		goto if_end296
	}

if_then295:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end296:
	v126 = *libc.As[byte](result)
	loadedv297 = (v126 & 1) != 0
	*libc.As[bool](retval) = loadedv297
	goto _return

sw_bb298:
	*libc.As[byte](result) = 1
	v127 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol299 = libc.Ptr(&libc.As[TSLexer](v127).F1)
	*libc.As[int16](result_symbol299) = 6
	v128 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end300 = libc.Ptr(&libc.As[TSLexer](v128).F3)
	v129 = *libc.As[unsafe.Pointer](mark_end300)
	v130 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v129)(v130)
	v131 = *libc.As[int32](lookahead)
	cmp301 = v131 == 41
	if cmp301 {
		goto if_then309
	} else {
		goto lor_lhs_false303
	}

lor_lhs_false303:
	v132 = *libc.As[int32](lookahead)
	cmp304 = v132 == 93
	if cmp304 {
		goto if_then309
	} else {
		goto lor_lhs_false306
	}

lor_lhs_false306:
	v133 = *libc.As[int32](lookahead)
	cmp307 = v133 == 125
	if cmp307 {
		goto if_then309
	} else {
		goto if_end310
	}

if_then309:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end310:
	v134 = *libc.As[int32](lookahead)
	cmp311 = v134 != 0
	if cmp311 {
		goto land_lhs_true
	} else {
		goto if_end328
	}

land_lhs_true:
	v135 = *libc.As[int32](lookahead)
	cmp313 = v135 != 9
	if cmp313 {
		goto land_lhs_true315
	} else {
		goto if_end328
	}

land_lhs_true315:
	v136 = *libc.As[int32](lookahead)
	cmp316 = v136 != 10
	if cmp316 {
		goto land_lhs_true318
	} else {
		goto if_end328
	}

land_lhs_true318:
	v137 = *libc.As[int32](lookahead)
	cmp319 = v137 != 13
	if cmp319 {
		goto land_lhs_true321
	} else {
		goto if_end328
	}

land_lhs_true321:
	v138 = *libc.As[int32](lookahead)
	cmp322 = v138 != 32
	if cmp322 {
		goto land_lhs_true324
	} else {
		goto if_end328
	}

land_lhs_true324:
	v139 = *libc.As[int32](lookahead)
	cmp325 = v139 != 37
	if cmp325 {
		goto if_then327
	} else {
		goto if_end328
	}

if_then327:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end328:
	v140 = *libc.As[byte](result)
	loadedv329 = (v140 & 1) != 0
	*libc.As[bool](retval) = loadedv329
	goto _return

sw_bb330:
	*libc.As[byte](result) = 1
	v141 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol331 = libc.Ptr(&libc.As[TSLexer](v141).F1)
	*libc.As[int16](result_symbol331) = 7
	v142 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end332 = libc.Ptr(&libc.As[TSLexer](v142).F3)
	v143 = *libc.As[unsafe.Pointer](mark_end332)
	v144 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v143)(v144)
	v145 = *libc.As[byte](result)
	loadedv333 = (v145 & 1) != 0
	*libc.As[bool](retval) = loadedv333
	goto _return

sw_bb334:
	*libc.As[byte](result) = 1
	v146 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol335 = libc.Ptr(&libc.As[TSLexer](v146).F1)
	*libc.As[int16](result_symbol335) = 7
	v147 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end336 = libc.Ptr(&libc.As[TSLexer](v147).F3)
	v148 = *libc.As[unsafe.Pointer](mark_end336)
	v149 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v148)(v149)
	v150 = *libc.As[int32](lookahead)
	cmp337 = v150 != 0
	if cmp337 {
		goto land_lhs_true339
	} else {
		goto if_end355
	}

land_lhs_true339:
	v151 = *libc.As[int32](lookahead)
	cmp340 = v151 != 9
	if cmp340 {
		goto land_lhs_true342
	} else {
		goto if_end355
	}

land_lhs_true342:
	v152 = *libc.As[int32](lookahead)
	cmp343 = v152 != 10
	if cmp343 {
		goto land_lhs_true345
	} else {
		goto if_end355
	}

land_lhs_true345:
	v153 = *libc.As[int32](lookahead)
	cmp346 = v153 != 13
	if cmp346 {
		goto land_lhs_true348
	} else {
		goto if_end355
	}

land_lhs_true348:
	v154 = *libc.As[int32](lookahead)
	cmp349 = v154 != 32
	if cmp349 {
		goto land_lhs_true351
	} else {
		goto if_end355
	}

land_lhs_true351:
	v155 = *libc.As[int32](lookahead)
	cmp352 = v155 != 37
	if cmp352 {
		goto if_then354
	} else {
		goto if_end355
	}

if_then354:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end355:
	v156 = *libc.As[byte](result)
	loadedv356 = (v156 & 1) != 0
	*libc.As[bool](retval) = loadedv356
	goto _return

sw_bb357:
	*libc.As[byte](result) = 1
	v157 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol358 = libc.Ptr(&libc.As[TSLexer](v157).F1)
	*libc.As[int16](result_symbol358) = 8
	v158 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end359 = libc.Ptr(&libc.As[TSLexer](v158).F3)
	v159 = *libc.As[unsafe.Pointer](mark_end359)
	v160 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v159)(v160)
	v161 = *libc.As[byte](result)
	loadedv360 = (v161 & 1) != 0
	*libc.As[bool](retval) = loadedv360
	goto _return

sw_bb361:
	*libc.As[byte](result) = 1
	v162 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol362 = libc.Ptr(&libc.As[TSLexer](v162).F1)
	*libc.As[int16](result_symbol362) = 8
	v163 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end363 = libc.Ptr(&libc.As[TSLexer](v163).F3)
	v164 = *libc.As[unsafe.Pointer](mark_end363)
	v165 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v164)(v165)
	v166 = *libc.As[int32](lookahead)
	cmp364 = v166 != 0
	if cmp364 {
		goto land_lhs_true366
	} else {
		goto if_end382
	}

land_lhs_true366:
	v167 = *libc.As[int32](lookahead)
	cmp367 = v167 != 9
	if cmp367 {
		goto land_lhs_true369
	} else {
		goto if_end382
	}

land_lhs_true369:
	v168 = *libc.As[int32](lookahead)
	cmp370 = v168 != 10
	if cmp370 {
		goto land_lhs_true372
	} else {
		goto if_end382
	}

land_lhs_true372:
	v169 = *libc.As[int32](lookahead)
	cmp373 = v169 != 13
	if cmp373 {
		goto land_lhs_true375
	} else {
		goto if_end382
	}

land_lhs_true375:
	v170 = *libc.As[int32](lookahead)
	cmp376 = v170 != 32
	if cmp376 {
		goto land_lhs_true378
	} else {
		goto if_end382
	}

land_lhs_true378:
	v171 = *libc.As[int32](lookahead)
	cmp379 = v171 != 37
	if cmp379 {
		goto if_then381
	} else {
		goto if_end382
	}

if_then381:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end382:
	v172 = *libc.As[byte](result)
	loadedv383 = (v172 & 1) != 0
	*libc.As[bool](retval) = loadedv383
	goto _return

sw_bb384:
	*libc.As[byte](result) = 1
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol385 = libc.Ptr(&libc.As[TSLexer](v173).F1)
	*libc.As[int16](result_symbol385) = 9
	v174 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end386 = libc.Ptr(&libc.As[TSLexer](v174).F3)
	v175 = *libc.As[unsafe.Pointer](mark_end386)
	v176 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v175)(v176)
	v177 = *libc.As[byte](result)
	loadedv387 = (v177 & 1) != 0
	*libc.As[bool](retval) = loadedv387
	goto _return

sw_bb388:
	*libc.As[byte](result) = 1
	v178 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol389 = libc.Ptr(&libc.As[TSLexer](v178).F1)
	*libc.As[int16](result_symbol389) = 10
	v179 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end390 = libc.Ptr(&libc.As[TSLexer](v179).F3)
	v180 = *libc.As[unsafe.Pointer](mark_end390)
	v181 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v180)(v181)
	v182 = *libc.As[byte](result)
	loadedv391 = (v182 & 1) != 0
	*libc.As[bool](retval) = loadedv391
	goto _return

sw_bb392:
	*libc.As[byte](result) = 1
	v183 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol393 = libc.Ptr(&libc.As[TSLexer](v183).F1)
	*libc.As[int16](result_symbol393) = 11
	v184 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end394 = libc.Ptr(&libc.As[TSLexer](v184).F3)
	v185 = *libc.As[unsafe.Pointer](mark_end394)
	v186 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v185)(v186)
	v187 = *libc.As[byte](result)
	loadedv395 = (v187 & 1) != 0
	*libc.As[bool](retval) = loadedv395
	goto _return

sw_bb396:
	*libc.As[byte](result) = 1
	v188 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol397 = libc.Ptr(&libc.As[TSLexer](v188).F1)
	*libc.As[int16](result_symbol397) = 12
	v189 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end398 = libc.Ptr(&libc.As[TSLexer](v189).F3)
	v190 = *libc.As[unsafe.Pointer](mark_end398)
	v191 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v190)(v191)
	v192 = *libc.As[int32](lookahead)
	cmp399 = v192 == 45
	if cmp399 {
		goto if_then401
	} else {
		goto if_end402
	}

if_then401:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end402:
	v193 = *libc.As[int32](lookahead)
	cmp403 = v193 == 9
	if cmp403 {
		goto if_then414
	} else {
		goto lor_lhs_false405
	}

lor_lhs_false405:
	v194 = *libc.As[int32](lookahead)
	cmp406 = v194 == 10
	if cmp406 {
		goto if_then414
	} else {
		goto lor_lhs_false408
	}

lor_lhs_false408:
	v195 = *libc.As[int32](lookahead)
	cmp409 = v195 == 13
	if cmp409 {
		goto if_then414
	} else {
		goto lor_lhs_false411
	}

lor_lhs_false411:
	v196 = *libc.As[int32](lookahead)
	cmp412 = v196 == 32
	if cmp412 {
		goto if_then414
	} else {
		goto if_end415
	}

if_then414:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end415:
	v197 = *libc.As[int32](lookahead)
	cmp416 = v197 != 0
	if cmp416 {
		goto if_then418
	} else {
		goto if_end419
	}

if_then418:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end419:
	v198 = *libc.As[byte](result)
	loadedv420 = (v198 & 1) != 0
	*libc.As[bool](retval) = loadedv420
	goto _return

sw_bb421:
	*libc.As[byte](result) = 1
	v199 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol422 = libc.Ptr(&libc.As[TSLexer](v199).F1)
	*libc.As[int16](result_symbol422) = 12
	v200 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end423 = libc.Ptr(&libc.As[TSLexer](v200).F3)
	v201 = *libc.As[unsafe.Pointer](mark_end423)
	v202 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v201)(v202)
	v203 = *libc.As[int32](lookahead)
	cmp424 = v203 == 45
	if cmp424 {
		goto if_then426
	} else {
		goto if_end427
	}

if_then426:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end427:
	v204 = *libc.As[byte](result)
	loadedv428 = (v204 & 1) != 0
	*libc.As[bool](retval) = loadedv428
	goto _return

sw_bb429:
	*libc.As[byte](result) = 1
	v205 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol430 = libc.Ptr(&libc.As[TSLexer](v205).F1)
	*libc.As[int16](result_symbol430) = 12
	v206 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end431 = libc.Ptr(&libc.As[TSLexer](v206).F3)
	v207 = *libc.As[unsafe.Pointer](mark_end431)
	v208 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v207)(v208)
	v209 = *libc.As[int32](lookahead)
	cmp432 = v209 != 0
	if cmp432 {
		goto land_lhs_true434
	} else {
		goto if_end438
	}

land_lhs_true434:
	v210 = *libc.As[int32](lookahead)
	cmp435 = v210 != 45
	if cmp435 {
		goto if_then437
	} else {
		goto if_end438
	}

if_then437:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end438:
	v211 = *libc.As[byte](result)
	loadedv439 = (v211 & 1) != 0
	*libc.As[bool](retval) = loadedv439
	goto _return

sw_bb440:
	*libc.As[byte](result) = 1
	v212 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol441 = libc.Ptr(&libc.As[TSLexer](v212).F1)
	*libc.As[int16](result_symbol441) = 13
	v213 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end442 = libc.Ptr(&libc.As[TSLexer](v213).F3)
	v214 = *libc.As[unsafe.Pointer](mark_end442)
	v215 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v214)(v215)
	v216 = *libc.As[byte](result)
	loadedv443 = (v216 & 1) != 0
	*libc.As[bool](retval) = loadedv443
	goto _return

sw_bb444:
	*libc.As[byte](result) = 1
	v217 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol445 = libc.Ptr(&libc.As[TSLexer](v217).F1)
	*libc.As[int16](result_symbol445) = 14
	v218 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end446 = libc.Ptr(&libc.As[TSLexer](v218).F3)
	v219 = *libc.As[unsafe.Pointer](mark_end446)
	v220 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v219)(v220)
	v221 = *libc.As[int32](lookahead)
	cmp447 = v221 == 37
	if cmp447 {
		goto if_then449
	} else {
		goto if_end450
	}

if_then449:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end450:
	v222 = *libc.As[byte](result)
	loadedv451 = (v222 & 1) != 0
	*libc.As[bool](retval) = loadedv451
	goto _return

sw_bb452:
	*libc.As[byte](result) = 1
	v223 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol453 = libc.Ptr(&libc.As[TSLexer](v223).F1)
	*libc.As[int16](result_symbol453) = 14
	v224 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end454 = libc.Ptr(&libc.As[TSLexer](v224).F3)
	v225 = *libc.As[unsafe.Pointer](mark_end454)
	v226 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v225)(v226)
	v227 = *libc.As[int32](lookahead)
	cmp455 = v227 == 60
	if cmp455 {
		goto if_then457
	} else {
		goto if_end458
	}

if_then457:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end458:
	v228 = *libc.As[int32](lookahead)
	cmp459 = v228 == 9
	if cmp459 {
		goto if_then470
	} else {
		goto lor_lhs_false461
	}

lor_lhs_false461:
	v229 = *libc.As[int32](lookahead)
	cmp462 = v229 == 10
	if cmp462 {
		goto if_then470
	} else {
		goto lor_lhs_false464
	}

lor_lhs_false464:
	v230 = *libc.As[int32](lookahead)
	cmp465 = v230 == 13
	if cmp465 {
		goto if_then470
	} else {
		goto lor_lhs_false467
	}

lor_lhs_false467:
	v231 = *libc.As[int32](lookahead)
	cmp468 = v231 == 32
	if cmp468 {
		goto if_then470
	} else {
		goto if_end471
	}

if_then470:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end471:
	v232 = *libc.As[int32](lookahead)
	cmp472 = v232 != 0
	if cmp472 {
		goto if_then474
	} else {
		goto if_end475
	}

if_then474:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end475:
	v233 = *libc.As[byte](result)
	loadedv476 = (v233 & 1) != 0
	*libc.As[bool](retval) = loadedv476
	goto _return

sw_bb477:
	*libc.As[byte](result) = 1
	v234 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol478 = libc.Ptr(&libc.As[TSLexer](v234).F1)
	*libc.As[int16](result_symbol478) = 14
	v235 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end479 = libc.Ptr(&libc.As[TSLexer](v235).F3)
	v236 = *libc.As[unsafe.Pointer](mark_end479)
	v237 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v236)(v237)
	v238 = *libc.As[int32](lookahead)
	cmp480 = v238 != 0
	if cmp480 {
		goto land_lhs_true482
	} else {
		goto if_end486
	}

land_lhs_true482:
	v239 = *libc.As[int32](lookahead)
	cmp483 = v239 != 60
	if cmp483 {
		goto if_then485
	} else {
		goto if_end486
	}

if_then485:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end486:
	v240 = *libc.As[byte](result)
	loadedv487 = (v240 & 1) != 0
	*libc.As[bool](retval) = loadedv487
	goto _return

sw_bb488:
	*libc.As[byte](result) = 1
	v241 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol489 = libc.Ptr(&libc.As[TSLexer](v241).F1)
	*libc.As[int16](result_symbol489) = 15
	v242 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end490 = libc.Ptr(&libc.As[TSLexer](v242).F3)
	v243 = *libc.As[unsafe.Pointer](mark_end490)
	v244 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v243)(v244)
	v245 = *libc.As[int32](lookahead)
	cmp491 = v245 == 37
	if cmp491 {
		goto if_then493
	} else {
		goto if_end494
	}

if_then493:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end494:
	v246 = *libc.As[int32](lookahead)
	cmp495 = v246 == 45
	if cmp495 {
		goto if_then497
	} else {
		goto if_end498
	}

if_then497:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end498:
	v247 = *libc.As[int32](lookahead)
	cmp499 = v247 == 100
	if cmp499 {
		goto if_then501
	} else {
		goto if_end502
	}

if_then501:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end502:
	v248 = *libc.As[int32](lookahead)
	cmp503 = v248 == 101
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end506:
	v249 = *libc.As[int32](lookahead)
	cmp507 = v249 == 9
	if cmp507 {
		goto if_then518
	} else {
		goto lor_lhs_false509
	}

lor_lhs_false509:
	v250 = *libc.As[int32](lookahead)
	cmp510 = v250 == 10
	if cmp510 {
		goto if_then518
	} else {
		goto lor_lhs_false512
	}

lor_lhs_false512:
	v251 = *libc.As[int32](lookahead)
	cmp513 = v251 == 13
	if cmp513 {
		goto if_then518
	} else {
		goto lor_lhs_false515
	}

lor_lhs_false515:
	v252 = *libc.As[int32](lookahead)
	cmp516 = v252 == 32
	if cmp516 {
		goto if_then518
	} else {
		goto if_end519
	}

if_then518:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end519:
	v253 = *libc.As[int32](lookahead)
	cmp520 = v253 != 0
	if cmp520 {
		goto if_then522
	} else {
		goto if_end523
	}

if_then522:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end523:
	v254 = *libc.As[byte](result)
	loadedv524 = (v254 & 1) != 0
	*libc.As[bool](retval) = loadedv524
	goto _return

sw_bb525:
	*libc.As[byte](result) = 1
	v255 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol526 = libc.Ptr(&libc.As[TSLexer](v255).F1)
	*libc.As[int16](result_symbol526) = 15
	v256 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end527 = libc.Ptr(&libc.As[TSLexer](v256).F3)
	v257 = *libc.As[unsafe.Pointer](mark_end527)
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v257)(v258)
	v259 = *libc.As[int32](lookahead)
	cmp528 = v259 == 37
	if cmp528 {
		goto if_then530
	} else {
		goto if_end531
	}

if_then530:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end531:
	v260 = *libc.As[int32](lookahead)
	cmp532 = v260 == 45
	if cmp532 {
		goto if_then534
	} else {
		goto if_end535
	}

if_then534:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end535:
	v261 = *libc.As[int32](lookahead)
	cmp536 = v261 == 100
	if cmp536 {
		goto if_then538
	} else {
		goto if_end539
	}

if_then538:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end539:
	v262 = *libc.As[int32](lookahead)
	cmp540 = v262 == 9
	if cmp540 {
		goto if_then551
	} else {
		goto lor_lhs_false542
	}

lor_lhs_false542:
	v263 = *libc.As[int32](lookahead)
	cmp543 = v263 == 10
	if cmp543 {
		goto if_then551
	} else {
		goto lor_lhs_false545
	}

lor_lhs_false545:
	v264 = *libc.As[int32](lookahead)
	cmp546 = v264 == 13
	if cmp546 {
		goto if_then551
	} else {
		goto lor_lhs_false548
	}

lor_lhs_false548:
	v265 = *libc.As[int32](lookahead)
	cmp549 = v265 == 32
	if cmp549 {
		goto if_then551
	} else {
		goto if_end552
	}

if_then551:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end552:
	v266 = *libc.As[int32](lookahead)
	cmp553 = v266 != 0
	if cmp553 {
		goto if_then555
	} else {
		goto if_end556
	}

if_then555:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end556:
	v267 = *libc.As[byte](result)
	loadedv557 = (v267 & 1) != 0
	*libc.As[bool](retval) = loadedv557
	goto _return

sw_bb558:
	*libc.As[byte](result) = 1
	v268 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol559 = libc.Ptr(&libc.As[TSLexer](v268).F1)
	*libc.As[int16](result_symbol559) = 15
	v269 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end560 = libc.Ptr(&libc.As[TSLexer](v269).F3)
	v270 = *libc.As[unsafe.Pointer](mark_end560)
	v271 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v270)(v271)
	v272 = *libc.As[int32](lookahead)
	cmp561 = v272 == 37
	if cmp561 {
		goto if_then563
	} else {
		goto if_end564
	}

if_then563:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end564:
	v273 = *libc.As[int32](lookahead)
	cmp565 = v273 == 9
	if cmp565 {
		goto if_then576
	} else {
		goto lor_lhs_false567
	}

lor_lhs_false567:
	v274 = *libc.As[int32](lookahead)
	cmp568 = v274 == 10
	if cmp568 {
		goto if_then576
	} else {
		goto lor_lhs_false570
	}

lor_lhs_false570:
	v275 = *libc.As[int32](lookahead)
	cmp571 = v275 == 13
	if cmp571 {
		goto if_then576
	} else {
		goto lor_lhs_false573
	}

lor_lhs_false573:
	v276 = *libc.As[int32](lookahead)
	cmp574 = v276 == 32
	if cmp574 {
		goto if_then576
	} else {
		goto if_end577
	}

if_then576:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end577:
	v277 = *libc.As[int32](lookahead)
	cmp578 = v277 != 0
	if cmp578 {
		goto if_then580
	} else {
		goto if_end581
	}

if_then580:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end581:
	v278 = *libc.As[byte](result)
	loadedv582 = (v278 & 1) != 0
	*libc.As[bool](retval) = loadedv582
	goto _return

sw_bb583:
	*libc.As[byte](result) = 1
	v279 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol584 = libc.Ptr(&libc.As[TSLexer](v279).F1)
	*libc.As[int16](result_symbol584) = 15
	v280 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end585 = libc.Ptr(&libc.As[TSLexer](v280).F3)
	v281 = *libc.As[unsafe.Pointer](mark_end585)
	v282 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v281)(v282)
	v283 = *libc.As[int32](lookahead)
	cmp586 = v283 == 62
	if cmp586 {
		goto if_then588
	} else {
		goto if_end589
	}

if_then588:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end589:
	v284 = *libc.As[byte](result)
	loadedv590 = (v284 & 1) != 0
	*libc.As[bool](retval) = loadedv590
	goto _return

sw_bb591:
	*libc.As[byte](result) = 1
	v285 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol592 = libc.Ptr(&libc.As[TSLexer](v285).F1)
	*libc.As[int16](result_symbol592) = 15
	v286 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end593 = libc.Ptr(&libc.As[TSLexer](v286).F3)
	v287 = *libc.As[unsafe.Pointer](mark_end593)
	v288 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v287)(v288)
	v289 = *libc.As[int32](lookahead)
	cmp594 = v289 == 62
	if cmp594 {
		goto if_then596
	} else {
		goto if_end597
	}

if_then596:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end597:
	v290 = *libc.As[int32](lookahead)
	cmp598 = v290 != 0
	if cmp598 {
		goto land_lhs_true600
	} else {
		goto if_end616
	}

land_lhs_true600:
	v291 = *libc.As[int32](lookahead)
	cmp601 = v291 != 9
	if cmp601 {
		goto land_lhs_true603
	} else {
		goto if_end616
	}

land_lhs_true603:
	v292 = *libc.As[int32](lookahead)
	cmp604 = v292 != 10
	if cmp604 {
		goto land_lhs_true606
	} else {
		goto if_end616
	}

land_lhs_true606:
	v293 = *libc.As[int32](lookahead)
	cmp607 = v293 != 13
	if cmp607 {
		goto land_lhs_true609
	} else {
		goto if_end616
	}

land_lhs_true609:
	v294 = *libc.As[int32](lookahead)
	cmp610 = v294 != 32
	if cmp610 {
		goto land_lhs_true612
	} else {
		goto if_end616
	}

land_lhs_true612:
	v295 = *libc.As[int32](lookahead)
	cmp613 = v295 != 37
	if cmp613 {
		goto if_then615
	} else {
		goto if_end616
	}

if_then615:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end616:
	v296 = *libc.As[byte](result)
	loadedv617 = (v296 & 1) != 0
	*libc.As[bool](retval) = loadedv617
	goto _return

sw_bb618:
	*libc.As[byte](result) = 1
	v297 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol619 = libc.Ptr(&libc.As[TSLexer](v297).F1)
	*libc.As[int16](result_symbol619) = 15
	v298 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end620 = libc.Ptr(&libc.As[TSLexer](v298).F3)
	v299 = *libc.As[unsafe.Pointer](mark_end620)
	v300 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v299)(v300)
	v301 = *libc.As[int32](lookahead)
	cmp621 = v301 == 100
	if cmp621 {
		goto if_then623
	} else {
		goto if_end624
	}

if_then623:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end624:
	v302 = *libc.As[int32](lookahead)
	cmp625 = v302 != 0
	if cmp625 {
		goto land_lhs_true627
	} else {
		goto if_end643
	}

land_lhs_true627:
	v303 = *libc.As[int32](lookahead)
	cmp628 = v303 != 9
	if cmp628 {
		goto land_lhs_true630
	} else {
		goto if_end643
	}

land_lhs_true630:
	v304 = *libc.As[int32](lookahead)
	cmp631 = v304 != 10
	if cmp631 {
		goto land_lhs_true633
	} else {
		goto if_end643
	}

land_lhs_true633:
	v305 = *libc.As[int32](lookahead)
	cmp634 = v305 != 13
	if cmp634 {
		goto land_lhs_true636
	} else {
		goto if_end643
	}

land_lhs_true636:
	v306 = *libc.As[int32](lookahead)
	cmp637 = v306 != 32
	if cmp637 {
		goto land_lhs_true639
	} else {
		goto if_end643
	}

land_lhs_true639:
	v307 = *libc.As[int32](lookahead)
	cmp640 = v307 != 37
	if cmp640 {
		goto if_then642
	} else {
		goto if_end643
	}

if_then642:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end643:
	v308 = *libc.As[byte](result)
	loadedv644 = (v308 & 1) != 0
	*libc.As[bool](retval) = loadedv644
	goto _return

sw_bb645:
	*libc.As[byte](result) = 1
	v309 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol646 = libc.Ptr(&libc.As[TSLexer](v309).F1)
	*libc.As[int16](result_symbol646) = 15
	v310 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end647 = libc.Ptr(&libc.As[TSLexer](v310).F3)
	v311 = *libc.As[unsafe.Pointer](mark_end647)
	v312 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v311)(v312)
	v313 = *libc.As[int32](lookahead)
	cmp648 = v313 == 110
	if cmp648 {
		goto if_then650
	} else {
		goto if_end651
	}

if_then650:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end651:
	v314 = *libc.As[int32](lookahead)
	cmp652 = v314 != 0
	if cmp652 {
		goto land_lhs_true654
	} else {
		goto if_end670
	}

land_lhs_true654:
	v315 = *libc.As[int32](lookahead)
	cmp655 = v315 != 9
	if cmp655 {
		goto land_lhs_true657
	} else {
		goto if_end670
	}

land_lhs_true657:
	v316 = *libc.As[int32](lookahead)
	cmp658 = v316 != 10
	if cmp658 {
		goto land_lhs_true660
	} else {
		goto if_end670
	}

land_lhs_true660:
	v317 = *libc.As[int32](lookahead)
	cmp661 = v317 != 13
	if cmp661 {
		goto land_lhs_true663
	} else {
		goto if_end670
	}

land_lhs_true663:
	v318 = *libc.As[int32](lookahead)
	cmp664 = v318 != 32
	if cmp664 {
		goto land_lhs_true666
	} else {
		goto if_end670
	}

land_lhs_true666:
	v319 = *libc.As[int32](lookahead)
	cmp667 = v319 != 37
	if cmp667 {
		goto if_then669
	} else {
		goto if_end670
	}

if_then669:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end670:
	v320 = *libc.As[byte](result)
	loadedv671 = (v320 & 1) != 0
	*libc.As[bool](retval) = loadedv671
	goto _return

sw_bb672:
	*libc.As[byte](result) = 1
	v321 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol673 = libc.Ptr(&libc.As[TSLexer](v321).F1)
	*libc.As[int16](result_symbol673) = 15
	v322 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end674 = libc.Ptr(&libc.As[TSLexer](v322).F3)
	v323 = *libc.As[unsafe.Pointer](mark_end674)
	v324 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v323)(v324)
	v325 = *libc.As[int32](lookahead)
	cmp675 = v325 == 111
	if cmp675 {
		goto if_then677
	} else {
		goto if_end678
	}

if_then677:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end678:
	v326 = *libc.As[int32](lookahead)
	cmp679 = v326 != 0
	if cmp679 {
		goto land_lhs_true681
	} else {
		goto if_end697
	}

land_lhs_true681:
	v327 = *libc.As[int32](lookahead)
	cmp682 = v327 != 9
	if cmp682 {
		goto land_lhs_true684
	} else {
		goto if_end697
	}

land_lhs_true684:
	v328 = *libc.As[int32](lookahead)
	cmp685 = v328 != 10
	if cmp685 {
		goto land_lhs_true687
	} else {
		goto if_end697
	}

land_lhs_true687:
	v329 = *libc.As[int32](lookahead)
	cmp688 = v329 != 13
	if cmp688 {
		goto land_lhs_true690
	} else {
		goto if_end697
	}

land_lhs_true690:
	v330 = *libc.As[int32](lookahead)
	cmp691 = v330 != 32
	if cmp691 {
		goto land_lhs_true693
	} else {
		goto if_end697
	}

land_lhs_true693:
	v331 = *libc.As[int32](lookahead)
	cmp694 = v331 != 37
	if cmp694 {
		goto if_then696
	} else {
		goto if_end697
	}

if_then696:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end697:
	v332 = *libc.As[byte](result)
	loadedv698 = (v332 & 1) != 0
	*libc.As[bool](retval) = loadedv698
	goto _return

sw_bb699:
	*libc.As[byte](result) = 1
	v333 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol700 = libc.Ptr(&libc.As[TSLexer](v333).F1)
	*libc.As[int16](result_symbol700) = 15
	v334 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end701 = libc.Ptr(&libc.As[TSLexer](v334).F3)
	v335 = *libc.As[unsafe.Pointer](mark_end701)
	v336 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v335)(v336)
	v337 = *libc.As[int32](lookahead)
	cmp702 = v337 != 0
	if cmp702 {
		goto land_lhs_true704
	} else {
		goto if_end720
	}

land_lhs_true704:
	v338 = *libc.As[int32](lookahead)
	cmp705 = v338 != 9
	if cmp705 {
		goto land_lhs_true707
	} else {
		goto if_end720
	}

land_lhs_true707:
	v339 = *libc.As[int32](lookahead)
	cmp708 = v339 != 10
	if cmp708 {
		goto land_lhs_true710
	} else {
		goto if_end720
	}

land_lhs_true710:
	v340 = *libc.As[int32](lookahead)
	cmp711 = v340 != 13
	if cmp711 {
		goto land_lhs_true713
	} else {
		goto if_end720
	}

land_lhs_true713:
	v341 = *libc.As[int32](lookahead)
	cmp714 = v341 != 32
	if cmp714 {
		goto land_lhs_true716
	} else {
		goto if_end720
	}

land_lhs_true716:
	v342 = *libc.As[int32](lookahead)
	cmp717 = v342 != 37
	if cmp717 {
		goto if_then719
	} else {
		goto if_end720
	}

if_then719:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end720:
	v343 = *libc.As[byte](result)
	loadedv721 = (v343 & 1) != 0
	*libc.As[bool](retval) = loadedv721
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v344 = *libc.As[bool](retval)
	return v344
}
