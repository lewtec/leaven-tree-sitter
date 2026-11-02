package grammar_printf

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type TSLanguageMetadata struct {
	F0 byte
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

var tree_sitter_printf_language struct {
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
var ts_parse_table [4][13]int16 = [4][13]int16{[13]int16{1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0}, [13]int16{3, 5, 7, 0, 0, 0, 0, 0, 9, 16, 2, 2, 2}, [13]int16{11, 13, 7, 0, 0, 0, 0, 0, 15, 0, 3, 3, 3}, [13]int16{17, 19, 22, 0, 0, 0, 0, 0, 25, 0, 3, 3, 3}}
var ts_small_parse_table [135]int16 = [135]int16{5, 28, 1, 3, 30, 1, 4, 32, 1, 5, 34, 1, 6, 36, 1, 7, 4, 38, 1, 3, 40, 1, 5, 42, 1, 6, 44, 1, 7, 2, 46, 2, 0, 8, 48, 2, 1, 2, 2, 50, 2, 0, 8, 52, 2, 1, 2, 2, 54, 2, 0, 8, 56, 2, 1, 2, 2, 58, 2, 0, 8, 60, 2, 1, 2, 2, 62, 2, 0, 8, 64, 2, 1, 2, 3, 38, 1, 3, 42, 1, 6, 44, 1, 7, 3, 66, 1, 3, 68, 1, 6, 70, 1, 7, 2, 38, 1, 3, 44, 1, 7, 2, 66, 1, 3, 70, 1, 7, 2, 72, 1, 3, 74, 1, 7, 1, 76, 1, 0, 1, 78, 1, 3, 1, 80, 1, 3, 1, 82, 1, 3, 1, 84, 1, 3}
var ts_small_parse_table_map [17]int32 = [17]int32{0, 16, 29, 38, 47, 56, 65, 74, 84, 94, 101, 108, 115, 119, 123, 127, 131}
var ts_symbol_names [13]unsafe.Pointer = [13]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_2), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13)}
var ts_symbol_metadata [13]TSSymbolMetadata = [13]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}}
var ts_symbol_map [13]int16 = [13]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][6]int16 = [1][6]int16{}
var ts_lex_modes [21]TSLexMode = [21]TSLexMode{TSLexMode{}, TSLexMode{5, 0}, TSLexMode{5, 0}, TSLexMode{5, 0}, TSLexMode{1, 0}, TSLexMode{}, TSLexMode{5, 0}, TSLexMode{5, 0}, TSLexMode{5, 0}, TSLexMode{5, 0}, TSLexMode{5, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}, TSLexMode{4, 0}}
var ts_primary_state_ids [21]int16 = [21]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
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
	F12 TSParseActionEntry
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
	F29 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F30 struct {
		F0 anon_2
		F1 [6]byte
	}
	F31 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F32 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F37 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F38 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F41 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F42 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F57 TSParseActionEntry
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
	F73 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
			F0 byte
			F1 [7]byte
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
	F12 TSParseActionEntry
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
	F29 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F30 struct {
		F0 anon_2
		F1 [6]byte
	}
	F31 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F32 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F37 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F38 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F41 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F42 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F57 TSParseActionEntry
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
	F73 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
			F0 byte
			F1 [7]byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 9, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 9, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 12, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 12, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 12, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 12, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{0, 0, 5, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 13, 0, 0}, [2]byte{}}}, struct {
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
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 6, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 6, 10, 0, 0}}}, struct {
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
}{0, 0, 18, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 20, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 10, 0, 0}, [2]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_2 [3]byte = [3]byte{37, 37, 0}
var _str_3 [2]byte = [2]byte{37, 0}
var _str_4 [5]byte = [5]byte{116, 121, 112, 101, 0}
var _str_5 [6]byte = [6]byte{102, 108, 97, 103, 115, 0}
var _str_6 [6]byte = [6]byte{119, 105, 100, 116, 104, 0}
var _str_7 [10]byte = [10]byte{112, 114, 101, 99, 105, 115, 105, 111, 110, 0}
var _str_8 [5]byte = [5]byte{115, 105, 122, 101, 0}
var _str_9 [13]byte = [13]byte{95, 116, 101, 120, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_10 [14]byte = [14]byte{102, 111, 114, 109, 97, 116, 95, 115, 116, 114, 105, 110, 103, 0}
var _str_11 [7]byte = [7]byte{102, 111, 114, 109, 97, 116, 0}
var _str_12 [6]byte = [6]byte{95, 116, 101, 120, 116, 0}
var _str_13 [22]byte = [22]byte{102, 111, 114, 109, 97, 116, 95, 115, 116, 114, 105, 110, 103, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var ts_lex_map [24]int16 = [24]int16{37, 8, 42, 13, 46, 16, 48, 14, 73, 19, 104, 20, 108, 21, 76, 18, 106, 18, 116, 18, 119, 18, 122, 18}
var ts_lex_map_14 [32]int16 = [32]int16{32, 11, 42, 13, 46, 16, 48, 12, 73, 19, 104, 20, 108, 21, 35, 10, 39, 10, 43, 10, 45, 10, 76, 18, 106, 18, 116, 18, 119, 18, 122, 18}
var ts_lex_map_15 [28]int16 = [28]int16{32, 11, 48, 12, 73, 19, 104, 20, 108, 21, 35, 10, 39, 10, 43, 10, 45, 10, 76, 18, 106, 18, 116, 18, 119, 18, 122, 18}

func init() {
	tree_sitter_printf_language = struct {
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
	}{14, 13, 0, 9, 0, 21, 4, 1, 0, 6, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), nil, nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{}, [5]byte{}}
}
func tree_sitter_printf() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_printf_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, cmp29, cmp32, cmp35, cmp38, loadedv42, cmp47, cmp53, cmp63, cmp66, cmp70, cmp73, cmp77, cmp80, cmp83, cmp86, loadedv90, cmp92, loadedv96, cmp98, loadedv102, cmp104, cmp107, cmp110, cmp114, cmp117, cmp120, cmp123, loadedv127, loadedv129, cmp132, cmp136, cmp139, cmp142, cmp146, loadedv150, loadedv152, loadedv156, cmp160, loadedv164, loadedv168, loadedv172, cmp179, cmp185, loadedv195, cmp199, cmp202, loadedv206, loadedv210, cmp214, cmp217, loadedv221, loadedv225, cmp229, cmp233, cmp236, loadedv240, cmp244, cmp247, loadedv251, loadedv255, cmp259, cmp263, loadedv267, cmp271, loadedv275, cmp279, loadedv283, cmp287, cmp290, cmp293, cmp297, cmp300, loadedv304, cmp308, cmp311, loadedv315, v181 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v30, v33, v95, v98 int16
	var state_addr, arrayidx, arrayidx11, arrayidx51, arrayidx58, result_symbol, result_symbol154, result_symbol158, result_symbol166, result_symbol170, result_symbol174, arrayidx183, arrayidx190, result_symbol197, result_symbol208, result_symbol212, result_symbol223, result_symbol227, result_symbol242, result_symbol253, result_symbol257, result_symbol269, result_symbol277, result_symbol285, result_symbol306 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v24, v25, v26, v28, v29, conv52, v31, v32, add56, v34, add61, v35, v36, v37, v38, v39, v40, v41, v42, v44, v46, v48, v49, v50, v51, v52, v53, v54, v57, v58, v59, v60, v61, v77, v93, v94, conv184, v96, v97, add188, v99, add193, v105, v106, v117, v118, v129, v130, v131, v137, v138, v149, v150, v156, v162, v168, v169, v170, v171, v172, v178, v179 int32
	var lookahead, i, i44, i176, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv46, idxprom50, idxprom57, conv178, idxprom182, idxprom189 int64
	var v3, storedv, v10, v27, v43, v45, v47, v55, v56, v62, v67, v72, v78, v83, v88, v100, v107, v112, v119, v124, v132, v139, v144, v151, v157, v163, v173, v180 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v63, v64, v65, v66, v68, v69, v70, v71, v73, v74, v75, v76, v79, v80, v81, v82, v84, v85, v86, v87, v89, v90, v91, v92, v101, v102, v103, v104, v108, v109, v110, v111, v113, v114, v115, v116, v120, v121, v122, v123, v125, v126, v127, v128, v133, v134, v135, v136, v140, v141, v142, v143, v145, v146, v147, v148, v152, v153, v154, v155, v158, v159, v160, v161, v164, v165, v166, v167, v174, v175, v176, v177 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end155, mark_end159, mark_end167, mark_end171, mark_end175, mark_end198, mark_end209, mark_end213, mark_end224, mark_end228, mark_end243, mark_end254, mark_end258, mark_end270, mark_end278, mark_end286, mark_end307 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i44, i176, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, cmp29, v24, cmp32, v25, cmp35, v26, cmp38, v27, loadedv42, v28, conv46, cmp47, v29, idxprom50, arrayidx51, v30, conv52, v31, cmp53, v32, add56, idxprom57, arrayidx58, v33, v34, add61, v35, cmp63, v36, cmp66, v37, cmp70, v38, cmp73, v39, cmp77, v40, cmp80, v41, cmp83, v42, cmp86, v43, loadedv90, v44, cmp92, v45, loadedv96, v46, cmp98, v47, loadedv102, v48, cmp104, v49, cmp107, v50, cmp110, v51, cmp114, v52, cmp117, v53, cmp120, v54, cmp123, v55, loadedv127, v56, loadedv129, v57, cmp132, v58, cmp136, v59, cmp139, v60, cmp142, v61, cmp146, v62, loadedv150, v63, result_symbol, v64, mark_end, v65, v66, v67, loadedv152, v68, result_symbol154, v69, mark_end155, v70, v71, v72, loadedv156, v73, result_symbol158, v74, mark_end159, v75, v76, v77, cmp160, v78, loadedv164, v79, result_symbol166, v80, mark_end167, v81, v82, v83, loadedv168, v84, result_symbol170, v85, mark_end171, v86, v87, v88, loadedv172, v89, result_symbol174, v90, mark_end175, v91, v92, v93, conv178, cmp179, v94, idxprom182, arrayidx183, v95, conv184, v96, cmp185, v97, add188, idxprom189, arrayidx190, v98, v99, add193, v100, loadedv195, v101, result_symbol197, v102, mark_end198, v103, v104, v105, cmp199, v106, cmp202, v107, loadedv206, v108, result_symbol208, v109, mark_end209, v110, v111, v112, loadedv210, v113, result_symbol212, v114, mark_end213, v115, v116, v117, cmp214, v118, cmp217, v119, loadedv221, v120, result_symbol223, v121, mark_end224, v122, v123, v124, loadedv225, v125, result_symbol227, v126, mark_end228, v127, v128, v129, cmp229, v130, cmp233, v131, cmp236, v132, loadedv240, v133, result_symbol242, v134, mark_end243, v135, v136, v137, cmp244, v138, cmp247, v139, loadedv251, v140, result_symbol253, v141, mark_end254, v142, v143, v144, loadedv255, v145, result_symbol257, v146, mark_end258, v147, v148, v149, cmp259, v150, cmp263, v151, loadedv267, v152, result_symbol269, v153, mark_end270, v154, v155, v156, cmp271, v157, loadedv275, v158, result_symbol277, v159, mark_end278, v160, v161, v162, cmp279, v163, loadedv283, v164, result_symbol285, v165, mark_end286, v166, v167, v168, cmp287, v169, cmp290, v170, cmp293, v171, cmp297, v172, cmp300, v173, loadedv304, v174, result_symbol306, v175, mark_end307, v176, v177, v178, cmp308, v179, cmp311, v180, loadedv315, v181

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
	i44 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i176 = libc.Ptr(&new(struct {
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
		goto sw_bb91
	case 3:
		goto sw_bb97
	case 4:
		goto sw_bb103
	case 5:
		goto sw_bb128
	case 6:
		goto sw_bb151
	case 7:
		goto sw_bb153
	case 8:
		goto sw_bb157
	case 9:
		goto sw_bb165
	case 10:
		goto sw_bb169
	case 11:
		goto sw_bb173
	case 12:
		goto sw_bb196
	case 13:
		goto sw_bb207
	case 14:
		goto sw_bb211
	case 15:
		goto sw_bb222
	case 16:
		goto sw_bb226
	case 17:
		goto sw_bb241
	case 18:
		goto sw_bb252
	case 19:
		goto sw_bb256
	case 20:
		goto sw_bb268
	case 21:
		goto sw_bb276
	case 22:
		goto sw_bb284
	case 23:
		goto sw_bb305
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
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v11 = *libc.As[int32](i)
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(24)
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
	cmp22 = 49 <= v21
	if cmp22 {
		goto land_lhs_true24
	} else {
		goto if_end28
	}

land_lhs_true24:
	v22 = *libc.As[int32](lookahead)
	cmp25 = v22 <= 57
	if cmp25 {
		goto if_then27
	} else {
		goto if_end28
	}

if_then27:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end28:
	v23 = *libc.As[int32](lookahead)
	cmp29 = 65 <= v23
	if cmp29 {
		goto land_lhs_true31
	} else {
		goto lor_lhs_false34
	}

land_lhs_true31:
	v24 = *libc.As[int32](lookahead)
	cmp32 = v24 <= 90
	if cmp32 {
		goto if_then40
	} else {
		goto lor_lhs_false34
	}

lor_lhs_false34:
	v25 = *libc.As[int32](lookahead)
	cmp35 = 97 <= v25
	if cmp35 {
		goto land_lhs_true37
	} else {
		goto if_end41
	}

land_lhs_true37:
	v26 = *libc.As[int32](lookahead)
	cmp38 = v26 <= 121
	if cmp38 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end41:
	v27 = *libc.As[byte](result)
	loadedv42 = (v27 & 1) != 0
	*libc.As[bool](retval) = loadedv42
	goto _return

sw_bb43:
	*libc.As[int32](i44) = 0
	goto for_cond45

for_cond45:
	v28 = *libc.As[int32](i44)
	conv46 = int64(uint64(uint32(v28)))
	cmp47 = uint64(conv46) < uint64(32)
	if cmp47 {
		goto for_body49
	} else {
		goto for_end62
	}

for_body49:
	v29 = *libc.As[int32](i44)
	idxprom50 = int64(uint64(uint32(v29)))
	arrayidx51 = libc.Ptr(&ts_lex_map_14[idxprom50])
	v30 = *libc.As[int16](arrayidx51)
	conv52 = int32(uint32(uint16(v30)))
	v31 = *libc.As[int32](lookahead)
	cmp53 = conv52 == v31
	if cmp53 {
		goto if_then55
	} else {
		goto if_end59
	}

if_then55:
	v32 = *libc.As[int32](i44)
	add56 = v32 + 1
	idxprom57 = int64(uint64(uint32(add56)))
	arrayidx58 = libc.Ptr(&ts_lex_map_14[idxprom57])
	v33 = *libc.As[int16](arrayidx58)
	*libc.As[int16](state_addr) = v33
	goto next_state

if_end59:
	goto for_inc60

for_inc60:
	v34 = *libc.As[int32](i44)
	add61 = v34 + 2
	*libc.As[int32](i44) = add61
	goto for_cond45

for_end62:
	v35 = *libc.As[int32](lookahead)
	cmp63 = 9 <= v35
	if cmp63 {
		goto land_lhs_true65
	} else {
		goto if_end69
	}

land_lhs_true65:
	v36 = *libc.As[int32](lookahead)
	cmp66 = v36 <= 13
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end69:
	v37 = *libc.As[int32](lookahead)
	cmp70 = 49 <= v37
	if cmp70 {
		goto land_lhs_true72
	} else {
		goto if_end76
	}

land_lhs_true72:
	v38 = *libc.As[int32](lookahead)
	cmp73 = v38 <= 57
	if cmp73 {
		goto if_then75
	} else {
		goto if_end76
	}

if_then75:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end76:
	v39 = *libc.As[int32](lookahead)
	cmp77 = 65 <= v39
	if cmp77 {
		goto land_lhs_true79
	} else {
		goto lor_lhs_false82
	}

land_lhs_true79:
	v40 = *libc.As[int32](lookahead)
	cmp80 = v40 <= 90
	if cmp80 {
		goto if_then88
	} else {
		goto lor_lhs_false82
	}

lor_lhs_false82:
	v41 = *libc.As[int32](lookahead)
	cmp83 = 97 <= v41
	if cmp83 {
		goto land_lhs_true85
	} else {
		goto if_end89
	}

land_lhs_true85:
	v42 = *libc.As[int32](lookahead)
	cmp86 = v42 <= 121
	if cmp86 {
		goto if_then88
	} else {
		goto if_end89
	}

if_then88:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end89:
	v43 = *libc.As[byte](result)
	loadedv90 = (v43 & 1) != 0
	*libc.As[bool](retval) = loadedv90
	goto _return

sw_bb91:
	v44 = *libc.As[int32](lookahead)
	cmp92 = v44 == 50
	if cmp92 {
		goto if_then94
	} else {
		goto if_end95
	}

if_then94:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end95:
	v45 = *libc.As[byte](result)
	loadedv96 = (v45 & 1) != 0
	*libc.As[bool](retval) = loadedv96
	goto _return

sw_bb97:
	v46 = *libc.As[int32](lookahead)
	cmp98 = v46 == 52
	if cmp98 {
		goto if_then100
	} else {
		goto if_end101
	}

if_then100:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end101:
	v47 = *libc.As[byte](result)
	loadedv102 = (v47 & 1) != 0
	*libc.As[bool](retval) = loadedv102
	goto _return

sw_bb103:
	v48 = *libc.As[int32](lookahead)
	cmp104 = 9 <= v48
	if cmp104 {
		goto land_lhs_true106
	} else {
		goto lor_lhs_false109
	}

land_lhs_true106:
	v49 = *libc.As[int32](lookahead)
	cmp107 = v49 <= 13
	if cmp107 {
		goto if_then112
	} else {
		goto lor_lhs_false109
	}

lor_lhs_false109:
	v50 = *libc.As[int32](lookahead)
	cmp110 = v50 == 32
	if cmp110 {
		goto if_then112
	} else {
		goto if_end113
	}

if_then112:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end113:
	v51 = *libc.As[int32](lookahead)
	cmp114 = 65 <= v51
	if cmp114 {
		goto land_lhs_true116
	} else {
		goto lor_lhs_false119
	}

land_lhs_true116:
	v52 = *libc.As[int32](lookahead)
	cmp117 = v52 <= 90
	if cmp117 {
		goto if_then125
	} else {
		goto lor_lhs_false119
	}

lor_lhs_false119:
	v53 = *libc.As[int32](lookahead)
	cmp120 = 97 <= v53
	if cmp120 {
		goto land_lhs_true122
	} else {
		goto if_end126
	}

land_lhs_true122:
	v54 = *libc.As[int32](lookahead)
	cmp123 = v54 <= 122
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end126:
	v55 = *libc.As[byte](result)
	loadedv127 = (v55 & 1) != 0
	*libc.As[bool](retval) = loadedv127
	goto _return

sw_bb128:
	v56 = *libc.As[byte](eof)
	loadedv129 = (v56 & 1) != 0
	if loadedv129 {
		goto if_then130
	} else {
		goto if_end131
	}

if_then130:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end131:
	v57 = *libc.As[int32](lookahead)
	cmp132 = v57 == 37
	if cmp132 {
		goto if_then134
	} else {
		goto if_end135
	}

if_then134:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end135:
	v58 = *libc.As[int32](lookahead)
	cmp136 = 9 <= v58
	if cmp136 {
		goto land_lhs_true138
	} else {
		goto lor_lhs_false141
	}

land_lhs_true138:
	v59 = *libc.As[int32](lookahead)
	cmp139 = v59 <= 13
	if cmp139 {
		goto if_then144
	} else {
		goto lor_lhs_false141
	}

lor_lhs_false141:
	v60 = *libc.As[int32](lookahead)
	cmp142 = v60 == 32
	if cmp142 {
		goto if_then144
	} else {
		goto if_end145
	}

if_then144:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end145:
	v61 = *libc.As[int32](lookahead)
	cmp146 = v61 != 0
	if cmp146 {
		goto if_then148
	} else {
		goto if_end149
	}

if_then148:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end149:
	v62 = *libc.As[byte](result)
	loadedv150 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv150
	goto _return

sw_bb151:
	*libc.As[byte](result) = 1
	v63 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v63).F1)
	*libc.As[int16](result_symbol) = 0
	v64 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v64).F3)
	v65 = *libc.As[unsafe.Pointer](mark_end)
	v66 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v65)(v66)
	v67 = *libc.As[byte](result)
	loadedv152 = (v67 & 1) != 0
	*libc.As[bool](retval) = loadedv152
	goto _return

sw_bb153:
	*libc.As[byte](result) = 1
	v68 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol154 = libc.Ptr(&libc.As[TSLexer](v68).F1)
	*libc.As[int16](result_symbol154) = 1
	v69 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end155 = libc.Ptr(&libc.As[TSLexer](v69).F3)
	v70 = *libc.As[unsafe.Pointer](mark_end155)
	v71 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v70)(v71)
	v72 = *libc.As[byte](result)
	loadedv156 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv156
	goto _return

sw_bb157:
	*libc.As[byte](result) = 1
	v73 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol158 = libc.Ptr(&libc.As[TSLexer](v73).F1)
	*libc.As[int16](result_symbol158) = 2
	v74 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end159 = libc.Ptr(&libc.As[TSLexer](v74).F3)
	v75 = *libc.As[unsafe.Pointer](mark_end159)
	v76 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v75)(v76)
	v77 = *libc.As[int32](lookahead)
	cmp160 = v77 == 37
	if cmp160 {
		goto if_then162
	} else {
		goto if_end163
	}

if_then162:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end163:
	v78 = *libc.As[byte](result)
	loadedv164 = (v78 & 1) != 0
	*libc.As[bool](retval) = loadedv164
	goto _return

sw_bb165:
	*libc.As[byte](result) = 1
	v79 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol166 = libc.Ptr(&libc.As[TSLexer](v79).F1)
	*libc.As[int16](result_symbol166) = 3
	v80 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end167 = libc.Ptr(&libc.As[TSLexer](v80).F3)
	v81 = *libc.As[unsafe.Pointer](mark_end167)
	v82 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v81)(v82)
	v83 = *libc.As[byte](result)
	loadedv168 = (v83 & 1) != 0
	*libc.As[bool](retval) = loadedv168
	goto _return

sw_bb169:
	*libc.As[byte](result) = 1
	v84 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol170 = libc.Ptr(&libc.As[TSLexer](v84).F1)
	*libc.As[int16](result_symbol170) = 4
	v85 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end171 = libc.Ptr(&libc.As[TSLexer](v85).F3)
	v86 = *libc.As[unsafe.Pointer](mark_end171)
	v87 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v86)(v87)
	v88 = *libc.As[byte](result)
	loadedv172 = (v88 & 1) != 0
	*libc.As[bool](retval) = loadedv172
	goto _return

sw_bb173:
	*libc.As[byte](result) = 1
	v89 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol174 = libc.Ptr(&libc.As[TSLexer](v89).F1)
	*libc.As[int16](result_symbol174) = 4
	v90 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end175 = libc.Ptr(&libc.As[TSLexer](v90).F3)
	v91 = *libc.As[unsafe.Pointer](mark_end175)
	v92 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v91)(v92)
	*libc.As[int32](i176) = 0
	goto for_cond177

for_cond177:
	v93 = *libc.As[int32](i176)
	conv178 = int64(uint64(uint32(v93)))
	cmp179 = uint64(conv178) < uint64(28)
	if cmp179 {
		goto for_body181
	} else {
		goto for_end194
	}

for_body181:
	v94 = *libc.As[int32](i176)
	idxprom182 = int64(uint64(uint32(v94)))
	arrayidx183 = libc.Ptr(&ts_lex_map_15[idxprom182])
	v95 = *libc.As[int16](arrayidx183)
	conv184 = int32(uint32(uint16(v95)))
	v96 = *libc.As[int32](lookahead)
	cmp185 = conv184 == v96
	if cmp185 {
		goto if_then187
	} else {
		goto if_end191
	}

if_then187:
	v97 = *libc.As[int32](i176)
	add188 = v97 + 1
	idxprom189 = int64(uint64(uint32(add188)))
	arrayidx190 = libc.Ptr(&ts_lex_map_15[idxprom189])
	v98 = *libc.As[int16](arrayidx190)
	*libc.As[int16](state_addr) = v98
	goto next_state

if_end191:
	goto for_inc192

for_inc192:
	v99 = *libc.As[int32](i176)
	add193 = v99 + 2
	*libc.As[int32](i176) = add193
	goto for_cond177

for_end194:
	v100 = *libc.As[byte](result)
	loadedv195 = (v100 & 1) != 0
	*libc.As[bool](retval) = loadedv195
	goto _return

sw_bb196:
	*libc.As[byte](result) = 1
	v101 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol197 = libc.Ptr(&libc.As[TSLexer](v101).F1)
	*libc.As[int16](result_symbol197) = 4
	v102 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end198 = libc.Ptr(&libc.As[TSLexer](v102).F3)
	v103 = *libc.As[unsafe.Pointer](mark_end198)
	v104 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v103)(v104)
	v105 = *libc.As[int32](lookahead)
	cmp199 = 48 <= v105
	if cmp199 {
		goto land_lhs_true201
	} else {
		goto if_end205
	}

land_lhs_true201:
	v106 = *libc.As[int32](lookahead)
	cmp202 = v106 <= 57
	if cmp202 {
		goto if_then204
	} else {
		goto if_end205
	}

if_then204:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end205:
	v107 = *libc.As[byte](result)
	loadedv206 = (v107 & 1) != 0
	*libc.As[bool](retval) = loadedv206
	goto _return

sw_bb207:
	*libc.As[byte](result) = 1
	v108 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol208 = libc.Ptr(&libc.As[TSLexer](v108).F1)
	*libc.As[int16](result_symbol208) = 5
	v109 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end209 = libc.Ptr(&libc.As[TSLexer](v109).F3)
	v110 = *libc.As[unsafe.Pointer](mark_end209)
	v111 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v110)(v111)
	v112 = *libc.As[byte](result)
	loadedv210 = (v112 & 1) != 0
	*libc.As[bool](retval) = loadedv210
	goto _return

sw_bb211:
	*libc.As[byte](result) = 1
	v113 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol212 = libc.Ptr(&libc.As[TSLexer](v113).F1)
	*libc.As[int16](result_symbol212) = 5
	v114 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end213 = libc.Ptr(&libc.As[TSLexer](v114).F3)
	v115 = *libc.As[unsafe.Pointer](mark_end213)
	v116 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v115)(v116)
	v117 = *libc.As[int32](lookahead)
	cmp214 = 48 <= v117
	if cmp214 {
		goto land_lhs_true216
	} else {
		goto if_end220
	}

land_lhs_true216:
	v118 = *libc.As[int32](lookahead)
	cmp217 = v118 <= 57
	if cmp217 {
		goto if_then219
	} else {
		goto if_end220
	}

if_then219:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end220:
	v119 = *libc.As[byte](result)
	loadedv221 = (v119 & 1) != 0
	*libc.As[bool](retval) = loadedv221
	goto _return

sw_bb222:
	*libc.As[byte](result) = 1
	v120 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol223 = libc.Ptr(&libc.As[TSLexer](v120).F1)
	*libc.As[int16](result_symbol223) = 6
	v121 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end224 = libc.Ptr(&libc.As[TSLexer](v121).F3)
	v122 = *libc.As[unsafe.Pointer](mark_end224)
	v123 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v122)(v123)
	v124 = *libc.As[byte](result)
	loadedv225 = (v124 & 1) != 0
	*libc.As[bool](retval) = loadedv225
	goto _return

sw_bb226:
	*libc.As[byte](result) = 1
	v125 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol227 = libc.Ptr(&libc.As[TSLexer](v125).F1)
	*libc.As[int16](result_symbol227) = 6
	v126 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end228 = libc.Ptr(&libc.As[TSLexer](v126).F3)
	v127 = *libc.As[unsafe.Pointer](mark_end228)
	v128 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v127)(v128)
	v129 = *libc.As[int32](lookahead)
	cmp229 = v129 == 42
	if cmp229 {
		goto if_then231
	} else {
		goto if_end232
	}

if_then231:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end232:
	v130 = *libc.As[int32](lookahead)
	cmp233 = 48 <= v130
	if cmp233 {
		goto land_lhs_true235
	} else {
		goto if_end239
	}

land_lhs_true235:
	v131 = *libc.As[int32](lookahead)
	cmp236 = v131 <= 57
	if cmp236 {
		goto if_then238
	} else {
		goto if_end239
	}

if_then238:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end239:
	v132 = *libc.As[byte](result)
	loadedv240 = (v132 & 1) != 0
	*libc.As[bool](retval) = loadedv240
	goto _return

sw_bb241:
	*libc.As[byte](result) = 1
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol242 = libc.Ptr(&libc.As[TSLexer](v133).F1)
	*libc.As[int16](result_symbol242) = 6
	v134 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end243 = libc.Ptr(&libc.As[TSLexer](v134).F3)
	v135 = *libc.As[unsafe.Pointer](mark_end243)
	v136 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v135)(v136)
	v137 = *libc.As[int32](lookahead)
	cmp244 = 48 <= v137
	if cmp244 {
		goto land_lhs_true246
	} else {
		goto if_end250
	}

land_lhs_true246:
	v138 = *libc.As[int32](lookahead)
	cmp247 = v138 <= 57
	if cmp247 {
		goto if_then249
	} else {
		goto if_end250
	}

if_then249:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end250:
	v139 = *libc.As[byte](result)
	loadedv251 = (v139 & 1) != 0
	*libc.As[bool](retval) = loadedv251
	goto _return

sw_bb252:
	*libc.As[byte](result) = 1
	v140 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol253 = libc.Ptr(&libc.As[TSLexer](v140).F1)
	*libc.As[int16](result_symbol253) = 7
	v141 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end254 = libc.Ptr(&libc.As[TSLexer](v141).F3)
	v142 = *libc.As[unsafe.Pointer](mark_end254)
	v143 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v142)(v143)
	v144 = *libc.As[byte](result)
	loadedv255 = (v144 & 1) != 0
	*libc.As[bool](retval) = loadedv255
	goto _return

sw_bb256:
	*libc.As[byte](result) = 1
	v145 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol257 = libc.Ptr(&libc.As[TSLexer](v145).F1)
	*libc.As[int16](result_symbol257) = 7
	v146 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end258 = libc.Ptr(&libc.As[TSLexer](v146).F3)
	v147 = *libc.As[unsafe.Pointer](mark_end258)
	v148 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v147)(v148)
	v149 = *libc.As[int32](lookahead)
	cmp259 = v149 == 51
	if cmp259 {
		goto if_then261
	} else {
		goto if_end262
	}

if_then261:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end262:
	v150 = *libc.As[int32](lookahead)
	cmp263 = v150 == 54
	if cmp263 {
		goto if_then265
	} else {
		goto if_end266
	}

if_then265:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end266:
	v151 = *libc.As[byte](result)
	loadedv267 = (v151 & 1) != 0
	*libc.As[bool](retval) = loadedv267
	goto _return

sw_bb268:
	*libc.As[byte](result) = 1
	v152 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol269 = libc.Ptr(&libc.As[TSLexer](v152).F1)
	*libc.As[int16](result_symbol269) = 7
	v153 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end270 = libc.Ptr(&libc.As[TSLexer](v153).F3)
	v154 = *libc.As[unsafe.Pointer](mark_end270)
	v155 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v154)(v155)
	v156 = *libc.As[int32](lookahead)
	cmp271 = v156 == 104
	if cmp271 {
		goto if_then273
	} else {
		goto if_end274
	}

if_then273:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end274:
	v157 = *libc.As[byte](result)
	loadedv275 = (v157 & 1) != 0
	*libc.As[bool](retval) = loadedv275
	goto _return

sw_bb276:
	*libc.As[byte](result) = 1
	v158 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol277 = libc.Ptr(&libc.As[TSLexer](v158).F1)
	*libc.As[int16](result_symbol277) = 7
	v159 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end278 = libc.Ptr(&libc.As[TSLexer](v159).F3)
	v160 = *libc.As[unsafe.Pointer](mark_end278)
	v161 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v160)(v161)
	v162 = *libc.As[int32](lookahead)
	cmp279 = v162 == 108
	if cmp279 {
		goto if_then281
	} else {
		goto if_end282
	}

if_then281:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end282:
	v163 = *libc.As[byte](result)
	loadedv283 = (v163 & 1) != 0
	*libc.As[bool](retval) = loadedv283
	goto _return

sw_bb284:
	*libc.As[byte](result) = 1
	v164 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol285 = libc.Ptr(&libc.As[TSLexer](v164).F1)
	*libc.As[int16](result_symbol285) = 8
	v165 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end286 = libc.Ptr(&libc.As[TSLexer](v165).F3)
	v166 = *libc.As[unsafe.Pointer](mark_end286)
	v167 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v166)(v167)
	v168 = *libc.As[int32](lookahead)
	cmp287 = 9 <= v168
	if cmp287 {
		goto land_lhs_true289
	} else {
		goto lor_lhs_false292
	}

land_lhs_true289:
	v169 = *libc.As[int32](lookahead)
	cmp290 = v169 <= 13
	if cmp290 {
		goto if_then295
	} else {
		goto lor_lhs_false292
	}

lor_lhs_false292:
	v170 = *libc.As[int32](lookahead)
	cmp293 = v170 == 32
	if cmp293 {
		goto if_then295
	} else {
		goto if_end296
	}

if_then295:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end296:
	v171 = *libc.As[int32](lookahead)
	cmp297 = v171 != 0
	if cmp297 {
		goto land_lhs_true299
	} else {
		goto if_end303
	}

land_lhs_true299:
	v172 = *libc.As[int32](lookahead)
	cmp300 = v172 != 37
	if cmp300 {
		goto if_then302
	} else {
		goto if_end303
	}

if_then302:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end303:
	v173 = *libc.As[byte](result)
	loadedv304 = (v173 & 1) != 0
	*libc.As[bool](retval) = loadedv304
	goto _return

sw_bb305:
	*libc.As[byte](result) = 1
	v174 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol306 = libc.Ptr(&libc.As[TSLexer](v174).F1)
	*libc.As[int16](result_symbol306) = 8
	v175 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end307 = libc.Ptr(&libc.As[TSLexer](v175).F3)
	v176 = *libc.As[unsafe.Pointer](mark_end307)
	v177 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v176)(v177)
	v178 = *libc.As[int32](lookahead)
	cmp308 = v178 != 0
	if cmp308 {
		goto land_lhs_true310
	} else {
		goto if_end314
	}

land_lhs_true310:
	v179 = *libc.As[int32](lookahead)
	cmp311 = v179 != 37
	if cmp311 {
		goto if_then313
	} else {
		goto if_end314
	}

if_then313:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end314:
	v180 = *libc.As[byte](result)
	loadedv315 = (v180 & 1) != 0
	*libc.As[bool](retval) = loadedv315
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v181 = *libc.As[bool](retval)
	return v181
}
