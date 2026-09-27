package grammar_ini

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
type TSLexerMode struct {
	F0 int16
	F1 int16
	F2 int16
}
type anon_2 struct {
	F0 byte
	F1 byte
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

var tree_sitter_ini_language struct {
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
var ts_small_parse_table [407]int16 = [407]int16{7, 3, 1, 4, 5, 1, 8, 13, 1, 5, 8, 1, 17, 13, 1, 13, 19, 2, 0, 1, 5, 2, 14, 15, 8, 3, 1, 4, 5, 1, 8, 9, 1, 1, 15, 1, 0, 5, 1, 12, 9, 1, 18, 18, 1, 11, 6, 2, 14, 15, 8, 3, 1, 4, 5, 1, 8, 9, 1, 1, 17, 1, 0, 5, 1, 12, 9, 1, 18, 18, 1, 11, 7, 2, 14, 15, 7, 3, 1, 4, 5, 1, 8, 13, 1, 5, 10, 1, 17, 13, 1, 13, 21, 2, 0, 1, 8, 2, 14, 15, 7, 3, 1, 4, 5, 1, 8, 23, 1, 0, 25, 1, 1, 5, 1, 12, 18, 1, 11, 9, 3, 14, 15, 18, 6, 3, 1, 4, 5, 1, 8, 30, 1, 5, 13, 1, 13, 28, 2, 0, 1, 10, 3, 14, 15, 17, 8, 3, 1, 4, 5, 1, 8, 9, 1, 1, 33, 1, 0, 5, 1, 12, 9, 1, 18, 18, 1, 11, 11, 2, 14, 15, 4, 5, 1, 8, 37, 1, 4, 35, 3, 0, 1, 5, 12, 3, 14, 15, 16, 4, 3, 1, 4, 5, 1, 8, 13, 2, 14, 15, 40, 3, 0, 1, 5, 4, 3, 1, 4, 5, 1, 8, 14, 2, 14, 15, 42, 3, 0, 1, 5, 4, 3, 1, 4, 5, 1, 8, 15, 2, 14, 15, 44, 3, 0, 1, 5, 4, 3, 1, 4, 5, 1, 8, 16, 2, 14, 15, 46, 3, 0, 1, 5, 3, 5, 1, 8, 17, 2, 14, 15, 48, 4, 0, 1, 4, 5, 4, 3, 1, 4, 5, 1, 8, 50, 2, 0, 1, 18, 2, 14, 15, 4, 52, 1, 2, 54, 1, 4, 56, 1, 8, 19, 2, 14, 15, 4, 3, 1, 4, 5, 1, 8, 58, 1, 0, 20, 2, 14, 15, 4, 3, 1, 4, 5, 1, 8, 60, 1, 6, 21, 2, 14, 15, 4, 54, 1, 4, 56, 1, 8, 62, 1, 9, 22, 2, 14, 15, 4, 3, 1, 4, 5, 1, 8, 64, 1, 3, 23, 2, 14, 15, 4, 56, 1, 8, 66, 1, 4, 68, 1, 7, 24, 2, 14, 15, 3, 5, 1, 8, 70, 1, 4, 25, 2, 14, 15, 3, 5, 1, 8, 72, 1, 4, 26, 2, 14, 15, 3, 5, 1, 8, 74, 1, 4, 27, 2, 14, 15, 1, 48, 1, 0, 1, 76, 1, 0}
var ts_small_parse_table_map [25]int32 = [25]int32{0, 24, 50, 76, 100, 124, 146, 172, 189, 205, 221, 237, 253, 267, 282, 296, 310, 324, 338, 352, 366, 377, 388, 399, 403}
var ts_symbol_names [19]unsafe.Pointer = [19]unsafe.Pointer{libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_5), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20)}
var ts_field_names [2]unsafe.Pointer = [2]unsafe.Pointer{nil, libc.Ptr(&_str_21)}
var ts_field_map_slices [2]TSMapSlice = [2]TSMapSlice{TSMapSlice{}, TSMapSlice{0, 1}}
var ts_field_map_entries [1]TSFieldMapEntry = [1]TSFieldMapEntry{TSFieldMapEntry{1, 0, 0}}
var ts_symbol_metadata [19]TSSymbolMetadata = [19]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [19]int16 = [19]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 2, 10, 11, 12, 13, 14, 15, 16, 17, 18}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [2][4]int16 = [2][4]int16{}
var ts_lex_modes [30]TSLexerMode = [30]TSLexerMode{TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{20, 0, 0}, TSLexerMode{}, TSLexerMode{2, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{-1, 0, 0}, TSLexerMode{-1, 0, 0}}
var ts_primary_state_ids [30]int16 = [30]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 17, 29}
var _str [4]byte = [4]byte{105, 110, 105, 0}
var ts_parse_table struct {
	F0 struct {
		F0 [9]int16
		F1 [10]int16
	}
	F1 [19]int16
	F2 [19]int16
	F3 [19]int16
	F4 [19]int16
} = struct {
	F0 struct {
		F0 [9]int16
		F1 [10]int16
	}
	F1 [19]int16
	F2 [19]int16
	F3 [19]int16
	F4 [19]int16
}{struct {
	F0 [9]int16
	F1 [10]int16
}{[9]int16{1, 1, 0, 1, 3, 0, 1, 0, 5}, [10]int16{}}, [19]int16{7, 9, 0, 0, 11, 13, 0, 0, 5, 0, 20, 18, 5, 13, 1, 1, 2, 3, 6}, [19]int16{15, 9, 0, 0, 11, 13, 0, 0, 5, 0, 0, 18, 5, 13, 2, 2, 12, 4, 7}, [19]int16{15, 9, 0, 0, 3, 13, 0, 0, 5, 0, 0, 18, 5, 13, 3, 3, 0, 10, 7}, [19]int16{17, 9, 0, 0, 3, 13, 0, 0, 5, 0, 0, 18, 5, 13, 4, 4, 0, 10, 11}}
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
	F8 TSParseActionEntry
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
	F16 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F22 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F31 TSParseActionEntry
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
	F36 TSParseActionEntry
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
	F59 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
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
	F77 TSParseActionEntry
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
	F8 TSParseActionEntry
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
	F16 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F22 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F31 TSParseActionEntry
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
	F36 TSParseActionEntry
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
	F59 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
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
	F77 TSParseActionEntry
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
}{0, 0, 22, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 10, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 11, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 11, 0, 0}}}, struct {
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
}{0, 0, 19, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 17, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 17, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 10, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 16, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 16, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 17, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 12, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 13, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 13, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 15, 0, 1}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 18, 0, 0}}}, struct {
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
}{0, 0, 28, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 24, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 14, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 14, 0, 0}}}}
var _str_3 [4]byte = [4]byte{101, 110, 100, 0}
var _str_4 [2]byte = [2]byte{91, 0}
var _str_5 [5]byte = [5]byte{116, 101, 120, 116, 0}
var _str_6 [2]byte = [2]byte{93, 0}
var _str_7 [20]byte = [20]byte{115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 95, 116, 111, 107, 101, 110, 50, 0}
var _str_8 [13]byte = [13]byte{115, 101, 116, 116, 105, 110, 103, 95, 110, 97, 109, 101, 0}
var _str_9 [2]byte = [2]byte{61, 0}
var _str_10 [14]byte = [14]byte{115, 101, 116, 116, 105, 110, 103, 95, 118, 97, 108, 117, 101, 0}
var _str_11 [15]byte = [15]byte{99, 111, 109, 109, 101, 110, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_12 [9]byte = [9]byte{100, 111, 99, 117, 109, 101, 110, 116, 0}
var _str_13 [8]byte = [8]byte{115, 101, 99, 116, 105, 111, 110, 0}
var _str_14 [13]byte = [13]byte{115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 0}
var _str_15 [8]byte = [8]byte{115, 101, 116, 116, 105, 110, 103, 0}
var _str_16 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_17 [7]byte = [7]byte{95, 98, 108, 97, 110, 107, 0}
var _str_18 [17]byte = [17]byte{100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_19 [17]byte = [17]byte{100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 50, 0}
var _str_20 [17]byte = [17]byte{100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 51, 0}
var _str_21 [6]byte = [6]byte{98, 108, 97, 110, 107, 0}

func init() {
	tree_sitter_ini_language = struct {
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
	}{15, 19, 0, 10, 0, 30, 5, 2, 1, 4, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), libc.Ptr(&ts_field_names), libc.Ptr(&ts_field_map_slices), libc.Ptr(&ts_field_map_entries), libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), libc.Ptr(&_str), nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{1, 4, 0}, [5]byte{}}
}
func tree_sitter_ini() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_ini_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp25, cmp29, cmp32, loadedv36, cmp38, loadedv42, cmp44, cmp48, cmp52, cmp55, cmp59, cmp62, cmp66, loadedv70, cmp72, cmp76, cmp80, cmp83, cmp87, cmp90, cmp94, cmp96, cmp99, loadedv103, cmp105, cmp109, cmp112, cmp115, cmp118, cmp121, cmp124, cmp127, loadedv131, loadedv133, cmp136, cmp140, cmp144, cmp148, cmp151, cmp155, cmp158, cmp162, cmp165, cmp168, cmp171, loadedv175, loadedv177, loadedv181, cmp185, cmp189, cmp193, cmp196, cmp200, cmp203, cmp207, cmp210, cmp213, loadedv217, cmp221, cmp225, cmp228, cmp231, loadedv235, cmp239, cmp242, cmp245, loadedv249, loadedv253, loadedv257, cmp261, cmp265, cmp268, cmp271, cmp274, cmp277, cmp280, cmp283, loadedv287, loadedv291, cmp295, cmp299, loadedv303, cmp307, cmp311, cmp314, cmp318, cmp321, cmp325, cmp328, cmp331, loadedv335, cmp339, cmp342, loadedv346, loadedv350, cmp354, cmp357, cmp360, loadedv364, cmp368, cmp371, cmp375, cmp378, cmp382, cmp385, cmp388, cmp391, loadedv395, cmp399, cmp402, cmp405, loadedv409, v193 bool
	var retval unsafe.Pointer
	var v9 int16
	var state_addr, result_symbol, result_symbol179, result_symbol183, result_symbol219, result_symbol237, result_symbol251, result_symbol255, result_symbol259, result_symbol289, result_symbol293, result_symbol305, result_symbol337, result_symbol348, result_symbol352, result_symbol366, result_symbol397 unsafe.Pointer
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v21, v23, v24, v25, v26, v27, v28, v29, v31, v32, v33, v34, v35, v36, v37, v38, v39, v41, v42, v43, v44, v45, v46, v47, v48, v51, v52, v53, v54, v55, v56, v57, v58, v59, v60, v61, v77, v78, v79, v80, v81, v82, v83, v84, v85, v91, v92, v93, v94, v100, v101, v102, v118, v119, v120, v121, v122, v123, v124, v125, v136, v137, v143, v144, v145, v146, v147, v148, v149, v150, v156, v157, v168, v169, v170, v176, v177, v178, v179, v180, v181, v182, v183, v189, v190, v191 int32
	var lookahead, lookahead1 unsafe.Pointer
	var v3, storedv, v10, v20, v22, v30, v40, v49, v50, v62, v67, v72, v86, v95, v103, v108, v113, v126, v131, v138, v151, v158, v163, v171, v184, v192 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v63, v64, v65, v66, v68, v69, v70, v71, v73, v74, v75, v76, v87, v88, v89, v90, v96, v97, v98, v99, v104, v105, v106, v107, v109, v110, v111, v112, v114, v115, v116, v117, v127, v128, v129, v130, v132, v133, v134, v135, v139, v140, v141, v142, v152, v153, v154, v155, v159, v160, v161, v162, v164, v165, v166, v167, v172, v173, v174, v175, v185, v186, v187, v188 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end180, mark_end184, mark_end220, mark_end238, mark_end252, mark_end256, mark_end260, mark_end290, mark_end294, mark_end306, mark_end338, mark_end349, mark_end353, mark_end367, mark_end398 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp25, v18, cmp29, v19, cmp32, v20, loadedv36, v21, cmp38, v22, loadedv42, v23, cmp44, v24, cmp48, v25, cmp52, v26, cmp55, v27, cmp59, v28, cmp62, v29, cmp66, v30, loadedv70, v31, cmp72, v32, cmp76, v33, cmp80, v34, cmp83, v35, cmp87, v36, cmp90, v37, cmp94, v38, cmp96, v39, cmp99, v40, loadedv103, v41, cmp105, v42, cmp109, v43, cmp112, v44, cmp115, v45, cmp118, v46, cmp121, v47, cmp124, v48, cmp127, v49, loadedv131, v50, loadedv133, v51, cmp136, v52, cmp140, v53, cmp144, v54, cmp148, v55, cmp151, v56, cmp155, v57, cmp158, v58, cmp162, v59, cmp165, v60, cmp168, v61, cmp171, v62, loadedv175, v63, result_symbol, v64, mark_end, v65, v66, v67, loadedv177, v68, result_symbol179, v69, mark_end180, v70, v71, v72, loadedv181, v73, result_symbol183, v74, mark_end184, v75, v76, v77, cmp185, v78, cmp189, v79, cmp193, v80, cmp196, v81, cmp200, v82, cmp203, v83, cmp207, v84, cmp210, v85, cmp213, v86, loadedv217, v87, result_symbol219, v88, mark_end220, v89, v90, v91, cmp221, v92, cmp225, v93, cmp228, v94, cmp231, v95, loadedv235, v96, result_symbol237, v97, mark_end238, v98, v99, v100, cmp239, v101, cmp242, v102, cmp245, v103, loadedv249, v104, result_symbol251, v105, mark_end252, v106, v107, v108, loadedv253, v109, result_symbol255, v110, mark_end256, v111, v112, v113, loadedv257, v114, result_symbol259, v115, mark_end260, v116, v117, v118, cmp261, v119, cmp265, v120, cmp268, v121, cmp271, v122, cmp274, v123, cmp277, v124, cmp280, v125, cmp283, v126, loadedv287, v127, result_symbol289, v128, mark_end290, v129, v130, v131, loadedv291, v132, result_symbol293, v133, mark_end294, v134, v135, v136, cmp295, v137, cmp299, v138, loadedv303, v139, result_symbol305, v140, mark_end306, v141, v142, v143, cmp307, v144, cmp311, v145, cmp314, v146, cmp318, v147, cmp321, v148, cmp325, v149, cmp328, v150, cmp331, v151, loadedv335, v152, result_symbol337, v153, mark_end338, v154, v155, v156, cmp339, v157, cmp342, v158, loadedv346, v159, result_symbol348, v160, mark_end349, v161, v162, v163, loadedv350, v164, result_symbol352, v165, mark_end353, v166, v167, v168, cmp354, v169, cmp357, v170, cmp360, v171, loadedv364, v172, result_symbol366, v173, mark_end367, v174, v175, v176, cmp368, v177, cmp371, v178, cmp375, v179, cmp378, v180, cmp382, v181, cmp385, v182, cmp388, v183, cmp391, v184, loadedv395, v185, result_symbol397, v186, mark_end398, v187, v188, v189, cmp399, v190, cmp402, v191, cmp405, v192, loadedv409, v193

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
		goto sw_bb37
	case 2:
		goto sw_bb43
	case 3:
		goto sw_bb71
	case 4:
		goto sw_bb104
	case 5:
		goto sw_bb132
	case 6:
		goto sw_bb176
	case 7:
		goto sw_bb178
	case 8:
		goto sw_bb182
	case 9:
		goto sw_bb218
	case 10:
		goto sw_bb236
	case 11:
		goto sw_bb250
	case 12:
		goto sw_bb254
	case 13:
		goto sw_bb258
	case 14:
		goto sw_bb288
	case 15:
		goto sw_bb292
	case 16:
		goto sw_bb304
	case 17:
		goto sw_bb336
	case 18:
		goto sw_bb347
	case 19:
		goto sw_bb351
	case 20:
		goto sw_bb365
	case 21:
		goto sw_bb396
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
	v11 = *libc.As[int32](lookahead)
	cmp = v11 == 10
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end6:
	v12 = *libc.As[int32](lookahead)
	cmp7 = v12 == 13
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end10:
	v13 = *libc.As[int32](lookahead)
	cmp11 = v13 == 61
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end14:
	v14 = *libc.As[int32](lookahead)
	cmp15 = v14 == 91
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
	cmp19 = v15 == 93
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
	cmp23 = v16 == 9
	if cmp23 {
		goto if_then27
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v17 = *libc.As[int32](lookahead)
	cmp25 = v17 == 32
	if cmp25 {
		goto if_then27
	} else {
		goto if_end28
	}

if_then27:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end28:
	v18 = *libc.As[int32](lookahead)
	cmp29 = v18 == 35
	if cmp29 {
		goto if_then34
	} else {
		goto lor_lhs_false31
	}

lor_lhs_false31:
	v19 = *libc.As[int32](lookahead)
	cmp32 = v19 == 59
	if cmp32 {
		goto if_then34
	} else {
		goto if_end35
	}

if_then34:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end35:
	v20 = *libc.As[byte](result)
	loadedv36 = (v20 & 1) != 0
	*libc.As[bool](retval) = loadedv36
	goto _return

sw_bb37:
	v21 = *libc.As[int32](lookahead)
	cmp38 = v21 == 10
	if cmp38 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end41:
	v22 = *libc.As[byte](result)
	loadedv42 = (v22 & 1) != 0
	*libc.As[bool](retval) = loadedv42
	goto _return

sw_bb43:
	v23 = *libc.As[int32](lookahead)
	cmp44 = v23 == 10
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end47:
	v24 = *libc.As[int32](lookahead)
	cmp48 = v24 == 13
	if cmp48 {
		goto if_then50
	} else {
		goto if_end51
	}

if_then50:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end51:
	v25 = *libc.As[int32](lookahead)
	cmp52 = v25 == 9
	if cmp52 {
		goto if_then57
	} else {
		goto lor_lhs_false54
	}

lor_lhs_false54:
	v26 = *libc.As[int32](lookahead)
	cmp55 = v26 == 32
	if cmp55 {
		goto if_then57
	} else {
		goto if_end58
	}

if_then57:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end58:
	v27 = *libc.As[int32](lookahead)
	cmp59 = v27 == 35
	if cmp59 {
		goto if_then64
	} else {
		goto lor_lhs_false61
	}

lor_lhs_false61:
	v28 = *libc.As[int32](lookahead)
	cmp62 = v28 == 59
	if cmp62 {
		goto if_then64
	} else {
		goto if_end65
	}

if_then64:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end65:
	v29 = *libc.As[int32](lookahead)
	cmp66 = v29 != 0
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end69:
	v30 = *libc.As[byte](result)
	loadedv70 = (v30 & 1) != 0
	*libc.As[bool](retval) = loadedv70
	goto _return

sw_bb71:
	v31 = *libc.As[int32](lookahead)
	cmp72 = v31 == 10
	if cmp72 {
		goto if_then74
	} else {
		goto if_end75
	}

if_then74:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end75:
	v32 = *libc.As[int32](lookahead)
	cmp76 = v32 == 13
	if cmp76 {
		goto if_then78
	} else {
		goto if_end79
	}

if_then78:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end79:
	v33 = *libc.As[int32](lookahead)
	cmp80 = v33 == 9
	if cmp80 {
		goto if_then85
	} else {
		goto lor_lhs_false82
	}

lor_lhs_false82:
	v34 = *libc.As[int32](lookahead)
	cmp83 = v34 == 32
	if cmp83 {
		goto if_then85
	} else {
		goto if_end86
	}

if_then85:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end86:
	v35 = *libc.As[int32](lookahead)
	cmp87 = v35 == 35
	if cmp87 {
		goto if_then92
	} else {
		goto lor_lhs_false89
	}

lor_lhs_false89:
	v36 = *libc.As[int32](lookahead)
	cmp90 = v36 == 59
	if cmp90 {
		goto if_then92
	} else {
		goto if_end93
	}

if_then92:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end93:
	v37 = *libc.As[int32](lookahead)
	cmp94 = v37 != 0
	if cmp94 {
		goto land_lhs_true
	} else {
		goto if_end102
	}

land_lhs_true:
	v38 = *libc.As[int32](lookahead)
	cmp96 = v38 != 91
	if cmp96 {
		goto land_lhs_true98
	} else {
		goto if_end102
	}

land_lhs_true98:
	v39 = *libc.As[int32](lookahead)
	cmp99 = v39 != 93
	if cmp99 {
		goto if_then101
	} else {
		goto if_end102
	}

if_then101:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end102:
	v40 = *libc.As[byte](result)
	loadedv103 = (v40 & 1) != 0
	*libc.As[bool](retval) = loadedv103
	goto _return

sw_bb104:
	v41 = *libc.As[int32](lookahead)
	cmp105 = v41 == 32
	if cmp105 {
		goto if_then107
	} else {
		goto if_end108
	}

if_then107:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end108:
	v42 = *libc.As[int32](lookahead)
	cmp109 = v42 != 0
	if cmp109 {
		goto land_lhs_true111
	} else {
		goto if_end130
	}

land_lhs_true111:
	v43 = *libc.As[int32](lookahead)
	cmp112 = v43 < 9
	if cmp112 {
		goto land_lhs_true117
	} else {
		goto lor_lhs_false114
	}

lor_lhs_false114:
	v44 = *libc.As[int32](lookahead)
	cmp115 = 13 < v44
	if cmp115 {
		goto land_lhs_true117
	} else {
		goto if_end130
	}

land_lhs_true117:
	v45 = *libc.As[int32](lookahead)
	cmp118 = v45 != 35
	if cmp118 {
		goto land_lhs_true120
	} else {
		goto if_end130
	}

land_lhs_true120:
	v46 = *libc.As[int32](lookahead)
	cmp121 = v46 != 59
	if cmp121 {
		goto land_lhs_true123
	} else {
		goto if_end130
	}

land_lhs_true123:
	v47 = *libc.As[int32](lookahead)
	cmp124 = v47 != 61
	if cmp124 {
		goto land_lhs_true126
	} else {
		goto if_end130
	}

land_lhs_true126:
	v48 = *libc.As[int32](lookahead)
	cmp127 = v48 != 91
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end130:
	v49 = *libc.As[byte](result)
	loadedv131 = (v49 & 1) != 0
	*libc.As[bool](retval) = loadedv131
	goto _return

sw_bb132:
	v50 = *libc.As[byte](eof)
	loadedv133 = (v50 & 1) != 0
	if loadedv133 {
		goto if_then134
	} else {
		goto if_end135
	}

if_then134:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end135:
	v51 = *libc.As[int32](lookahead)
	cmp136 = v51 == 10
	if cmp136 {
		goto if_then138
	} else {
		goto if_end139
	}

if_then138:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end139:
	v52 = *libc.As[int32](lookahead)
	cmp140 = v52 == 13
	if cmp140 {
		goto if_then142
	} else {
		goto if_end143
	}

if_then142:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end143:
	v53 = *libc.As[int32](lookahead)
	cmp144 = v53 == 91
	if cmp144 {
		goto if_then146
	} else {
		goto if_end147
	}

if_then146:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end147:
	v54 = *libc.As[int32](lookahead)
	cmp148 = v54 == 9
	if cmp148 {
		goto if_then153
	} else {
		goto lor_lhs_false150
	}

lor_lhs_false150:
	v55 = *libc.As[int32](lookahead)
	cmp151 = v55 == 32
	if cmp151 {
		goto if_then153
	} else {
		goto if_end154
	}

if_then153:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end154:
	v56 = *libc.As[int32](lookahead)
	cmp155 = v56 == 35
	if cmp155 {
		goto if_then160
	} else {
		goto lor_lhs_false157
	}

lor_lhs_false157:
	v57 = *libc.As[int32](lookahead)
	cmp158 = v57 == 59
	if cmp158 {
		goto if_then160
	} else {
		goto if_end161
	}

if_then160:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end161:
	v58 = *libc.As[int32](lookahead)
	cmp162 = v58 != 0
	if cmp162 {
		goto land_lhs_true164
	} else {
		goto if_end174
	}

land_lhs_true164:
	v59 = *libc.As[int32](lookahead)
	cmp165 = v59 < 9
	if cmp165 {
		goto land_lhs_true170
	} else {
		goto lor_lhs_false167
	}

lor_lhs_false167:
	v60 = *libc.As[int32](lookahead)
	cmp168 = 13 < v60
	if cmp168 {
		goto land_lhs_true170
	} else {
		goto if_end174
	}

land_lhs_true170:
	v61 = *libc.As[int32](lookahead)
	cmp171 = v61 != 61
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end174:
	v62 = *libc.As[byte](result)
	loadedv175 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv175
	goto _return

sw_bb176:
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
	loadedv177 = (v67 & 1) != 0
	*libc.As[bool](retval) = loadedv177
	goto _return

sw_bb178:
	*libc.As[byte](result) = 1
	v68 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol179 = libc.Ptr(&libc.As[TSLexer](v68).F1)
	*libc.As[int16](result_symbol179) = 1
	v69 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end180 = libc.Ptr(&libc.As[TSLexer](v69).F3)
	v70 = *libc.As[unsafe.Pointer](mark_end180)
	v71 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v70)(v71)
	v72 = *libc.As[byte](result)
	loadedv181 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv181
	goto _return

sw_bb182:
	*libc.As[byte](result) = 1
	v73 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol183 = libc.Ptr(&libc.As[TSLexer](v73).F1)
	*libc.As[int16](result_symbol183) = 2
	v74 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end184 = libc.Ptr(&libc.As[TSLexer](v74).F3)
	v75 = *libc.As[unsafe.Pointer](mark_end184)
	v76 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v75)(v76)
	v77 = *libc.As[int32](lookahead)
	cmp185 = v77 == 10
	if cmp185 {
		goto if_then187
	} else {
		goto if_end188
	}

if_then187:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end188:
	v78 = *libc.As[int32](lookahead)
	cmp189 = v78 == 13
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end192:
	v79 = *libc.As[int32](lookahead)
	cmp193 = v79 == 9
	if cmp193 {
		goto if_then198
	} else {
		goto lor_lhs_false195
	}

lor_lhs_false195:
	v80 = *libc.As[int32](lookahead)
	cmp196 = v80 == 32
	if cmp196 {
		goto if_then198
	} else {
		goto if_end199
	}

if_then198:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end199:
	v81 = *libc.As[int32](lookahead)
	cmp200 = v81 == 35
	if cmp200 {
		goto if_then205
	} else {
		goto lor_lhs_false202
	}

lor_lhs_false202:
	v82 = *libc.As[int32](lookahead)
	cmp203 = v82 == 59
	if cmp203 {
		goto if_then205
	} else {
		goto if_end206
	}

if_then205:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end206:
	v83 = *libc.As[int32](lookahead)
	cmp207 = v83 != 0
	if cmp207 {
		goto land_lhs_true209
	} else {
		goto if_end216
	}

land_lhs_true209:
	v84 = *libc.As[int32](lookahead)
	cmp210 = v84 != 91
	if cmp210 {
		goto land_lhs_true212
	} else {
		goto if_end216
	}

land_lhs_true212:
	v85 = *libc.As[int32](lookahead)
	cmp213 = v85 != 93
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end216:
	v86 = *libc.As[byte](result)
	loadedv217 = (v86 & 1) != 0
	*libc.As[bool](retval) = loadedv217
	goto _return

sw_bb218:
	*libc.As[byte](result) = 1
	v87 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol219 = libc.Ptr(&libc.As[TSLexer](v87).F1)
	*libc.As[int16](result_symbol219) = 2
	v88 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end220 = libc.Ptr(&libc.As[TSLexer](v88).F3)
	v89 = *libc.As[unsafe.Pointer](mark_end220)
	v90 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v89)(v90)
	v91 = *libc.As[int32](lookahead)
	cmp221 = v91 == 10
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end224:
	v92 = *libc.As[int32](lookahead)
	cmp225 = v92 != 0
	if cmp225 {
		goto land_lhs_true227
	} else {
		goto if_end234
	}

land_lhs_true227:
	v93 = *libc.As[int32](lookahead)
	cmp228 = v93 != 91
	if cmp228 {
		goto land_lhs_true230
	} else {
		goto if_end234
	}

land_lhs_true230:
	v94 = *libc.As[int32](lookahead)
	cmp231 = v94 != 93
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end234:
	v95 = *libc.As[byte](result)
	loadedv235 = (v95 & 1) != 0
	*libc.As[bool](retval) = loadedv235
	goto _return

sw_bb236:
	*libc.As[byte](result) = 1
	v96 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol237 = libc.Ptr(&libc.As[TSLexer](v96).F1)
	*libc.As[int16](result_symbol237) = 2
	v97 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end238 = libc.Ptr(&libc.As[TSLexer](v97).F3)
	v98 = *libc.As[unsafe.Pointer](mark_end238)
	v99 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v98)(v99)
	v100 = *libc.As[int32](lookahead)
	cmp239 = v100 != 0
	if cmp239 {
		goto land_lhs_true241
	} else {
		goto if_end248
	}

land_lhs_true241:
	v101 = *libc.As[int32](lookahead)
	cmp242 = v101 != 91
	if cmp242 {
		goto land_lhs_true244
	} else {
		goto if_end248
	}

land_lhs_true244:
	v102 = *libc.As[int32](lookahead)
	cmp245 = v102 != 93
	if cmp245 {
		goto if_then247
	} else {
		goto if_end248
	}

if_then247:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end248:
	v103 = *libc.As[byte](result)
	loadedv249 = (v103 & 1) != 0
	*libc.As[bool](retval) = loadedv249
	goto _return

sw_bb250:
	*libc.As[byte](result) = 1
	v104 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol251 = libc.Ptr(&libc.As[TSLexer](v104).F1)
	*libc.As[int16](result_symbol251) = 3
	v105 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end252 = libc.Ptr(&libc.As[TSLexer](v105).F3)
	v106 = *libc.As[unsafe.Pointer](mark_end252)
	v107 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v106)(v107)
	v108 = *libc.As[byte](result)
	loadedv253 = (v108 & 1) != 0
	*libc.As[bool](retval) = loadedv253
	goto _return

sw_bb254:
	*libc.As[byte](result) = 1
	v109 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol255 = libc.Ptr(&libc.As[TSLexer](v109).F1)
	*libc.As[int16](result_symbol255) = 4
	v110 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end256 = libc.Ptr(&libc.As[TSLexer](v110).F3)
	v111 = *libc.As[unsafe.Pointer](mark_end256)
	v112 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v111)(v112)
	v113 = *libc.As[byte](result)
	loadedv257 = (v113 & 1) != 0
	*libc.As[bool](retval) = loadedv257
	goto _return

sw_bb258:
	*libc.As[byte](result) = 1
	v114 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol259 = libc.Ptr(&libc.As[TSLexer](v114).F1)
	*libc.As[int16](result_symbol259) = 5
	v115 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end260 = libc.Ptr(&libc.As[TSLexer](v115).F3)
	v116 = *libc.As[unsafe.Pointer](mark_end260)
	v117 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v116)(v117)
	v118 = *libc.As[int32](lookahead)
	cmp261 = v118 == 32
	if cmp261 {
		goto if_then263
	} else {
		goto if_end264
	}

if_then263:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end264:
	v119 = *libc.As[int32](lookahead)
	cmp265 = v119 != 0
	if cmp265 {
		goto land_lhs_true267
	} else {
		goto if_end286
	}

land_lhs_true267:
	v120 = *libc.As[int32](lookahead)
	cmp268 = v120 < 9
	if cmp268 {
		goto land_lhs_true273
	} else {
		goto lor_lhs_false270
	}

lor_lhs_false270:
	v121 = *libc.As[int32](lookahead)
	cmp271 = 13 < v121
	if cmp271 {
		goto land_lhs_true273
	} else {
		goto if_end286
	}

land_lhs_true273:
	v122 = *libc.As[int32](lookahead)
	cmp274 = v122 != 35
	if cmp274 {
		goto land_lhs_true276
	} else {
		goto if_end286
	}

land_lhs_true276:
	v123 = *libc.As[int32](lookahead)
	cmp277 = v123 != 59
	if cmp277 {
		goto land_lhs_true279
	} else {
		goto if_end286
	}

land_lhs_true279:
	v124 = *libc.As[int32](lookahead)
	cmp280 = v124 != 61
	if cmp280 {
		goto land_lhs_true282
	} else {
		goto if_end286
	}

land_lhs_true282:
	v125 = *libc.As[int32](lookahead)
	cmp283 = v125 != 91
	if cmp283 {
		goto if_then285
	} else {
		goto if_end286
	}

if_then285:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end286:
	v126 = *libc.As[byte](result)
	loadedv287 = (v126 & 1) != 0
	*libc.As[bool](retval) = loadedv287
	goto _return

sw_bb288:
	*libc.As[byte](result) = 1
	v127 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol289 = libc.Ptr(&libc.As[TSLexer](v127).F1)
	*libc.As[int16](result_symbol289) = 6
	v128 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end290 = libc.Ptr(&libc.As[TSLexer](v128).F3)
	v129 = *libc.As[unsafe.Pointer](mark_end290)
	v130 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v129)(v130)
	v131 = *libc.As[byte](result)
	loadedv291 = (v131 & 1) != 0
	*libc.As[bool](retval) = loadedv291
	goto _return

sw_bb292:
	*libc.As[byte](result) = 1
	v132 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol293 = libc.Ptr(&libc.As[TSLexer](v132).F1)
	*libc.As[int16](result_symbol293) = 7
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end294 = libc.Ptr(&libc.As[TSLexer](v133).F3)
	v134 = *libc.As[unsafe.Pointer](mark_end294)
	v135 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v134)(v135)
	v136 = *libc.As[int32](lookahead)
	cmp295 = v136 == 10
	if cmp295 {
		goto if_then297
	} else {
		goto if_end298
	}

if_then297:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end298:
	v137 = *libc.As[int32](lookahead)
	cmp299 = v137 != 0
	if cmp299 {
		goto if_then301
	} else {
		goto if_end302
	}

if_then301:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end302:
	v138 = *libc.As[byte](result)
	loadedv303 = (v138 & 1) != 0
	*libc.As[bool](retval) = loadedv303
	goto _return

sw_bb304:
	*libc.As[byte](result) = 1
	v139 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol305 = libc.Ptr(&libc.As[TSLexer](v139).F1)
	*libc.As[int16](result_symbol305) = 7
	v140 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end306 = libc.Ptr(&libc.As[TSLexer](v140).F3)
	v141 = *libc.As[unsafe.Pointer](mark_end306)
	v142 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v141)(v142)
	v143 = *libc.As[int32](lookahead)
	cmp307 = v143 == 13
	if cmp307 {
		goto if_then309
	} else {
		goto if_end310
	}

if_then309:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end310:
	v144 = *libc.As[int32](lookahead)
	cmp311 = v144 == 9
	if cmp311 {
		goto if_then316
	} else {
		goto lor_lhs_false313
	}

lor_lhs_false313:
	v145 = *libc.As[int32](lookahead)
	cmp314 = v145 == 32
	if cmp314 {
		goto if_then316
	} else {
		goto if_end317
	}

if_then316:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end317:
	v146 = *libc.As[int32](lookahead)
	cmp318 = v146 == 35
	if cmp318 {
		goto if_then323
	} else {
		goto lor_lhs_false320
	}

lor_lhs_false320:
	v147 = *libc.As[int32](lookahead)
	cmp321 = v147 == 59
	if cmp321 {
		goto if_then323
	} else {
		goto if_end324
	}

if_then323:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end324:
	v148 = *libc.As[int32](lookahead)
	cmp325 = v148 != 0
	if cmp325 {
		goto land_lhs_true327
	} else {
		goto if_end334
	}

land_lhs_true327:
	v149 = *libc.As[int32](lookahead)
	cmp328 = v149 != 9
	if cmp328 {
		goto land_lhs_true330
	} else {
		goto if_end334
	}

land_lhs_true330:
	v150 = *libc.As[int32](lookahead)
	cmp331 = v150 != 10
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end334:
	v151 = *libc.As[byte](result)
	loadedv335 = (v151 & 1) != 0
	*libc.As[bool](retval) = loadedv335
	goto _return

sw_bb336:
	*libc.As[byte](result) = 1
	v152 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol337 = libc.Ptr(&libc.As[TSLexer](v152).F1)
	*libc.As[int16](result_symbol337) = 7
	v153 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end338 = libc.Ptr(&libc.As[TSLexer](v153).F3)
	v154 = *libc.As[unsafe.Pointer](mark_end338)
	v155 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v154)(v155)
	v156 = *libc.As[int32](lookahead)
	cmp339 = v156 != 0
	if cmp339 {
		goto land_lhs_true341
	} else {
		goto if_end345
	}

land_lhs_true341:
	v157 = *libc.As[int32](lookahead)
	cmp342 = v157 != 10
	if cmp342 {
		goto if_then344
	} else {
		goto if_end345
	}

if_then344:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end345:
	v158 = *libc.As[byte](result)
	loadedv346 = (v158 & 1) != 0
	*libc.As[bool](retval) = loadedv346
	goto _return

sw_bb347:
	*libc.As[byte](result) = 1
	v159 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol348 = libc.Ptr(&libc.As[TSLexer](v159).F1)
	*libc.As[int16](result_symbol348) = 8
	v160 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end349 = libc.Ptr(&libc.As[TSLexer](v160).F3)
	v161 = *libc.As[unsafe.Pointer](mark_end349)
	v162 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v161)(v162)
	v163 = *libc.As[byte](result)
	loadedv350 = (v163 & 1) != 0
	*libc.As[bool](retval) = loadedv350
	goto _return

sw_bb351:
	*libc.As[byte](result) = 1
	v164 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol352 = libc.Ptr(&libc.As[TSLexer](v164).F1)
	*libc.As[int16](result_symbol352) = 8
	v165 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end353 = libc.Ptr(&libc.As[TSLexer](v165).F3)
	v166 = *libc.As[unsafe.Pointer](mark_end353)
	v167 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v166)(v167)
	v168 = *libc.As[int32](lookahead)
	cmp354 = v168 != 0
	if cmp354 {
		goto land_lhs_true356
	} else {
		goto if_end363
	}

land_lhs_true356:
	v169 = *libc.As[int32](lookahead)
	cmp357 = v169 != 10
	if cmp357 {
		goto land_lhs_true359
	} else {
		goto if_end363
	}

land_lhs_true359:
	v170 = *libc.As[int32](lookahead)
	cmp360 = v170 != 13
	if cmp360 {
		goto if_then362
	} else {
		goto if_end363
	}

if_then362:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end363:
	v171 = *libc.As[byte](result)
	loadedv364 = (v171 & 1) != 0
	*libc.As[bool](retval) = loadedv364
	goto _return

sw_bb365:
	*libc.As[byte](result) = 1
	v172 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol366 = libc.Ptr(&libc.As[TSLexer](v172).F1)
	*libc.As[int16](result_symbol366) = 9
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end367 = libc.Ptr(&libc.As[TSLexer](v173).F3)
	v174 = *libc.As[unsafe.Pointer](mark_end367)
	v175 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v174)(v175)
	v176 = *libc.As[int32](lookahead)
	cmp368 = v176 == 9
	if cmp368 {
		goto if_then373
	} else {
		goto lor_lhs_false370
	}

lor_lhs_false370:
	v177 = *libc.As[int32](lookahead)
	cmp371 = v177 == 32
	if cmp371 {
		goto if_then373
	} else {
		goto if_end374
	}

if_then373:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end374:
	v178 = *libc.As[int32](lookahead)
	cmp375 = v178 == 35
	if cmp375 {
		goto if_then380
	} else {
		goto lor_lhs_false377
	}

lor_lhs_false377:
	v179 = *libc.As[int32](lookahead)
	cmp378 = v179 == 59
	if cmp378 {
		goto if_then380
	} else {
		goto if_end381
	}

if_then380:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end381:
	v180 = *libc.As[int32](lookahead)
	cmp382 = v180 != 0
	if cmp382 {
		goto land_lhs_true384
	} else {
		goto if_end394
	}

land_lhs_true384:
	v181 = *libc.As[int32](lookahead)
	cmp385 = v181 != 9
	if cmp385 {
		goto land_lhs_true387
	} else {
		goto if_end394
	}

land_lhs_true387:
	v182 = *libc.As[int32](lookahead)
	cmp388 = v182 != 10
	if cmp388 {
		goto land_lhs_true390
	} else {
		goto if_end394
	}

land_lhs_true390:
	v183 = *libc.As[int32](lookahead)
	cmp391 = v183 != 13
	if cmp391 {
		goto if_then393
	} else {
		goto if_end394
	}

if_then393:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end394:
	v184 = *libc.As[byte](result)
	loadedv395 = (v184 & 1) != 0
	*libc.As[bool](retval) = loadedv395
	goto _return

sw_bb396:
	*libc.As[byte](result) = 1
	v185 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol397 = libc.Ptr(&libc.As[TSLexer](v185).F1)
	*libc.As[int16](result_symbol397) = 9
	v186 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end398 = libc.Ptr(&libc.As[TSLexer](v186).F3)
	v187 = *libc.As[unsafe.Pointer](mark_end398)
	v188 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v187)(v188)
	v189 = *libc.As[int32](lookahead)
	cmp399 = v189 != 0
	if cmp399 {
		goto land_lhs_true401
	} else {
		goto if_end408
	}

land_lhs_true401:
	v190 = *libc.As[int32](lookahead)
	cmp402 = v190 != 10
	if cmp402 {
		goto land_lhs_true404
	} else {
		goto if_end408
	}

land_lhs_true404:
	v191 = *libc.As[int32](lookahead)
	cmp405 = v191 != 13
	if cmp405 {
		goto if_then407
	} else {
		goto if_end408
	}

if_then407:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end408:
	v192 = *libc.As[byte](result)
	loadedv409 = (v192 & 1) != 0
	*libc.As[bool](retval) = loadedv409
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v193 = *libc.As[bool](retval)
	return v193
}
