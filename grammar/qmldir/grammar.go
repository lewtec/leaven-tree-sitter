package grammar_qmldir

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
type TSLexerMode struct {
	F0 int16
	F1 int16
	F2 int16
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

var tree_sitter_qmldir_language struct {
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
var ts_parse_table [6][23]int16 = [6][23]int16{[23]int16{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 3, 0, 0, 0, 0, 0}, [23]int16{5, 0, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 0, 0, 0, 3, 10, 2, 6, 2, 0}, [23]int16{11, 0, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 0, 0, 0, 3, 0, 3, 6, 3, 0}, [23]int16{13, 0, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 18, 0, 0, 0, 3, 0, 3, 6, 3, 0}, [23]int16{21, 0, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 0, 0, 0, 3, 0, 0, 0, 0, 0}, [23]int16{25, 0, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 0, 0, 0, 3, 0, 0, 0, 0, 0}}
var ts_small_parse_table [64]int16 = [64]int16{4, 29, 1, 1, 33, 1, 17, 7, 1, 22, 31, 3, 14, 15, 16, 4, 33, 1, 17, 35, 1, 1, 8, 1, 22, 37, 3, 14, 15, 16, 4, 33, 1, 17, 39, 1, 1, 8, 1, 22, 41, 3, 14, 15, 16, 3, 33, 1, 17, 44, 1, 1, 46, 3, 14, 15, 16, 2, 3, 1, 17, 48, 1, 0}
var ts_small_parse_table_map [5]int32 = [5]int32{0, 15, 30, 45, 57}
var ts_symbol_names [23]unsafe.Pointer = [23]unsafe.Pointer{libc.Ptr(&_str_2), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24)}
var ts_symbol_metadata [23]TSSymbolMetadata = [23]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [23]int16 = [23]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][3]int16 = [1][3]int16{}
var ts_lex_modes [11]TSLexerMode = [11]TSLexerMode{TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{1, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{}}
var ts_primary_state_ids [11]int16 = [11]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var _str [7]byte = [7]byte{113, 109, 108, 100, 105, 114, 0}
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
	F12 TSParseActionEntry
	F13 struct {
		F0 anon_2
		F1 [6]byte
	}
	F14 TSParseActionEntry
	F15 struct {
		F0 anon_2
		F1 [6]byte
	}
	F16 TSParseActionEntry
	F17 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F18 struct {
		F0 anon_2
		F1 [6]byte
	}
	F19 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
	F47 TSParseActionEntry
	F48 struct {
		F0 anon_2
		F1 [6]byte
	}
	F49 struct {
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
	F12 TSParseActionEntry
	F13 struct {
		F0 anon_2
		F1 [6]byte
	}
	F14 TSParseActionEntry
	F15 struct {
		F0 anon_2
		F1 [6]byte
	}
	F16 TSParseActionEntry
	F17 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F18 struct {
		F0 anon_2
		F1 [6]byte
	}
	F19 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
	F47 TSParseActionEntry
	F48 struct {
		F0 anon_2
		F1 [6]byte
	}
	F49 struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 18, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 18, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 21, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 21, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 9, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 21, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 19, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 19, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 19, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 19, 0, 0}}}, struct {
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
}{0, 0, 0, 1, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 20, 0, 0}}}, struct {
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
var _str_2 [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [2]byte = [2]byte{10, 0}
var _str_4 [10]byte = [10]byte{99, 108, 97, 115, 115, 110, 97, 109, 101, 0}
var _str_5 [8]byte = [8]byte{100, 101, 112, 101, 110, 100, 115, 0}
var _str_6 [18]byte = [18]byte{100, 101, 115, 105, 103, 110, 101, 114, 115, 117, 112, 112, 111, 114, 116, 101, 100, 0}
var _str_7 [9]byte = [9]byte{105, 110, 116, 101, 114, 110, 97, 108, 0}
var _str_8 [11]byte = [11]byte{108, 105, 110, 107, 116, 97, 114, 103, 101, 116, 0}
var _str_9 [7]byte = [7]byte{109, 111, 100, 117, 108, 101, 0}
var _str_10 [9]byte = [9]byte{111, 112, 116, 105, 111, 110, 97, 108, 0}
var _str_11 [7]byte = [7]byte{112, 108, 117, 103, 105, 110, 0}
var _str_12 [7]byte = [7]byte{112, 114, 101, 102, 101, 114, 0}
var _str_13 [10]byte = [10]byte{115, 105, 110, 103, 108, 101, 116, 111, 110, 0}
var _str_14 [9]byte = [9]byte{116, 121, 112, 101, 105, 110, 102, 111, 0}
var _str_15 [11]byte = [11]byte{105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}
var _str_16 [5]byte = [5]byte{117, 110, 105, 116, 0}
var _str_17 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}
var _str_18 [6]byte = [6]byte{102, 108, 111, 97, 116, 0}
var _str_19 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_20 [18]byte = [18]byte{109, 111, 100, 117, 108, 101, 95, 100, 101, 102, 105, 110, 105, 116, 105, 111, 110, 0}
var _str_21 [8]byte = [8]byte{99, 111, 109, 109, 97, 110, 100, 0}
var _str_22 [8]byte = [8]byte{107, 101, 121, 119, 111, 114, 100, 0}
var _str_23 [26]byte = [26]byte{109, 111, 100, 117, 108, 101, 95, 100, 101, 102, 105, 110, 105, 116, 105, 111, 110, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_24 [16]byte = [16]byte{99, 111, 109, 109, 97, 110, 100, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var ts_lex_map [22]int16 = [22]int16{35, 3, 46, 2, 99, 50, 100, 25, 105, 57, 108, 43, 109, 68, 111, 73, 112, 51, 115, 45, 116, 96}

func init() {
	tree_sitter_qmldir_language = struct {
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
	}{15, 23, 0, 18, 0, 11, 6, 1, 0, 3, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), libc.Ptr(&_str), nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{0, 1, 0}, [5]byte{}}
}
func tree_sitter_qmldir() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_qmldir_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, cmp29, cmp32, cmp35, cmp38, cmp41, loadedv45, cmp47, cmp51, cmp55, cmp59, cmp62, cmp65, cmp69, cmp72, cmp76, cmp79, cmp82, cmp85, cmp88, cmp91, cmp94, loadedv98, cmp100, cmp103, loadedv107, cmp109, cmp112, loadedv116, loadedv118, cmp122, loadedv126, cmp130, cmp133, cmp136, cmp139, cmp142, cmp145, cmp148, loadedv152, cmp156, cmp159, cmp162, cmp165, cmp168, cmp171, cmp174, loadedv178, cmp182, cmp185, cmp188, cmp191, cmp194, cmp197, cmp200, loadedv204, cmp208, cmp211, cmp214, cmp217, cmp220, cmp223, cmp226, loadedv230, cmp234, cmp237, cmp240, cmp243, cmp246, cmp249, cmp252, loadedv256, cmp260, cmp263, cmp266, cmp269, cmp272, cmp275, cmp278, loadedv282, cmp286, cmp289, cmp292, cmp295, cmp298, cmp301, cmp304, loadedv308, cmp312, cmp315, cmp318, cmp321, cmp324, cmp327, cmp330, loadedv334, cmp338, cmp341, cmp344, cmp347, cmp350, cmp353, cmp356, loadedv360, cmp364, cmp367, cmp370, cmp373, cmp376, cmp379, cmp382, loadedv386, cmp390, cmp393, cmp396, cmp399, cmp402, cmp405, cmp408, loadedv412, cmp416, cmp420, cmp423, cmp426, cmp429, cmp432, cmp435, cmp438, loadedv442, cmp446, cmp450, cmp453, cmp456, cmp459, cmp462, cmp465, cmp468, loadedv472, cmp476, cmp480, cmp483, cmp486, cmp489, cmp492, cmp495, cmp498, loadedv502, cmp506, cmp510, cmp513, cmp516, cmp519, cmp522, cmp525, cmp528, loadedv532, cmp536, cmp540, cmp543, cmp546, cmp549, cmp552, cmp555, cmp558, loadedv562, cmp566, cmp570, cmp573, cmp576, cmp579, cmp582, cmp585, cmp588, loadedv592, cmp596, cmp600, cmp603, cmp606, cmp609, cmp612, cmp615, cmp618, loadedv622, cmp626, cmp630, cmp633, cmp636, cmp639, cmp642, cmp645, cmp648, loadedv652, cmp656, cmp660, cmp663, cmp666, cmp669, cmp672, cmp675, cmp678, loadedv682, cmp686, cmp690, cmp693, cmp696, cmp699, cmp702, cmp705, cmp708, loadedv712, cmp716, cmp720, cmp723, cmp726, cmp729, cmp732, cmp735, cmp738, loadedv742, cmp746, cmp750, cmp753, cmp756, cmp759, cmp762, cmp765, cmp768, loadedv772, cmp776, cmp780, cmp783, cmp786, cmp789, cmp792, cmp795, cmp798, loadedv802, cmp806, cmp810, cmp813, cmp816, cmp819, cmp822, cmp825, cmp828, loadedv832, cmp836, cmp840, cmp843, cmp846, cmp849, cmp852, cmp855, cmp858, loadedv862, cmp866, cmp870, cmp873, cmp876, cmp879, cmp882, cmp885, cmp888, loadedv892, cmp896, cmp900, cmp903, cmp906, cmp909, cmp912, cmp915, cmp918, loadedv922, cmp926, cmp930, cmp933, cmp936, cmp939, cmp942, cmp945, cmp948, loadedv952, cmp956, cmp960, cmp963, cmp966, cmp969, cmp972, cmp975, cmp978, loadedv982, cmp986, cmp990, cmp993, cmp996, cmp999, cmp1002, cmp1005, cmp1008, loadedv1012, cmp1016, cmp1020, cmp1023, cmp1026, cmp1029, cmp1032, cmp1035, cmp1038, loadedv1042, cmp1046, cmp1050, cmp1053, cmp1056, cmp1059, cmp1062, cmp1065, cmp1068, loadedv1072, cmp1076, cmp1080, cmp1083, cmp1086, cmp1089, cmp1092, cmp1095, cmp1098, loadedv1102, cmp1106, cmp1110, cmp1113, cmp1116, cmp1119, cmp1122, cmp1125, cmp1128, loadedv1132, cmp1136, cmp1140, cmp1143, cmp1146, cmp1149, cmp1152, cmp1155, cmp1158, loadedv1162, cmp1166, cmp1170, cmp1173, cmp1176, cmp1179, cmp1182, cmp1185, cmp1188, loadedv1192, cmp1196, cmp1200, cmp1203, cmp1206, cmp1209, cmp1212, cmp1215, cmp1218, loadedv1222, cmp1226, cmp1230, cmp1233, cmp1236, cmp1239, cmp1242, cmp1245, cmp1248, loadedv1252, cmp1256, cmp1260, cmp1263, cmp1266, cmp1269, cmp1272, cmp1275, cmp1278, loadedv1282, cmp1286, cmp1290, cmp1293, cmp1296, cmp1299, cmp1302, cmp1305, cmp1308, loadedv1312, cmp1316, cmp1320, cmp1323, cmp1326, cmp1329, cmp1332, cmp1335, cmp1338, loadedv1342, cmp1346, cmp1350, cmp1353, cmp1356, cmp1359, cmp1362, cmp1365, cmp1368, loadedv1372, cmp1376, cmp1380, cmp1383, cmp1386, cmp1389, cmp1392, cmp1395, cmp1398, loadedv1402, cmp1406, cmp1410, cmp1413, cmp1416, cmp1419, cmp1422, cmp1425, cmp1428, loadedv1432, cmp1436, cmp1440, cmp1444, cmp1447, cmp1450, cmp1453, cmp1456, cmp1459, cmp1462, loadedv1466, cmp1470, cmp1474, cmp1477, cmp1480, cmp1483, cmp1486, cmp1489, cmp1492, loadedv1496, cmp1500, cmp1504, cmp1507, cmp1510, cmp1513, cmp1516, cmp1519, cmp1522, loadedv1526, cmp1530, cmp1534, cmp1537, cmp1540, cmp1543, cmp1546, cmp1549, cmp1552, loadedv1556, cmp1560, cmp1564, cmp1567, cmp1570, cmp1573, cmp1576, cmp1579, cmp1582, loadedv1586, cmp1590, cmp1594, cmp1597, cmp1600, cmp1603, cmp1606, cmp1609, cmp1612, loadedv1616, cmp1620, cmp1624, cmp1627, cmp1630, cmp1633, cmp1636, cmp1639, cmp1642, loadedv1646, cmp1650, cmp1654, cmp1657, cmp1660, cmp1663, cmp1666, cmp1669, cmp1672, loadedv1676, cmp1680, cmp1684, cmp1687, cmp1690, cmp1693, cmp1696, cmp1699, cmp1702, loadedv1706, cmp1710, cmp1714, cmp1717, cmp1720, cmp1723, cmp1726, cmp1729, cmp1732, loadedv1736, cmp1740, cmp1744, cmp1747, cmp1750, cmp1753, cmp1756, cmp1759, cmp1762, loadedv1766, cmp1770, cmp1774, cmp1777, cmp1780, cmp1783, cmp1786, cmp1789, cmp1792, loadedv1796, cmp1800, cmp1804, cmp1807, cmp1810, cmp1813, cmp1816, cmp1819, cmp1822, loadedv1826, cmp1830, cmp1834, cmp1837, cmp1840, cmp1843, cmp1846, cmp1849, cmp1852, loadedv1856, cmp1860, cmp1864, cmp1867, cmp1870, cmp1873, cmp1876, cmp1879, cmp1882, loadedv1886, cmp1890, cmp1894, cmp1897, cmp1900, cmp1903, cmp1906, cmp1909, cmp1912, loadedv1916, cmp1920, cmp1924, cmp1927, cmp1930, cmp1933, cmp1936, cmp1939, cmp1942, loadedv1946, cmp1950, cmp1954, cmp1957, cmp1960, cmp1963, cmp1966, cmp1969, cmp1972, loadedv1976, cmp1980, cmp1984, cmp1987, cmp1990, cmp1993, cmp1996, cmp1999, cmp2002, loadedv2006, cmp2010, cmp2014, cmp2017, cmp2020, cmp2023, cmp2026, cmp2029, cmp2032, loadedv2036, cmp2040, cmp2044, cmp2047, cmp2050, cmp2053, cmp2056, cmp2059, cmp2062, loadedv2066, cmp2070, cmp2074, cmp2077, cmp2080, cmp2083, cmp2086, cmp2089, cmp2092, loadedv2096, cmp2100, cmp2104, cmp2107, cmp2110, cmp2113, cmp2116, cmp2119, cmp2122, loadedv2126, cmp2130, cmp2134, cmp2138, cmp2141, cmp2144, cmp2147, cmp2150, cmp2153, cmp2156, loadedv2160, cmp2164, cmp2168, cmp2171, cmp2174, cmp2177, cmp2180, cmp2183, cmp2186, loadedv2190, cmp2194, cmp2198, cmp2201, cmp2204, cmp2207, cmp2210, cmp2213, cmp2216, loadedv2220, cmp2224, cmp2228, cmp2231, cmp2234, cmp2237, cmp2240, cmp2243, cmp2246, loadedv2250, cmp2254, cmp2258, cmp2261, cmp2264, cmp2267, cmp2270, cmp2273, cmp2276, loadedv2280, cmp2284, cmp2288, cmp2291, cmp2294, cmp2297, cmp2300, cmp2303, cmp2306, loadedv2310, cmp2314, cmp2318, cmp2321, cmp2324, cmp2327, cmp2330, cmp2333, cmp2336, loadedv2340, cmp2344, cmp2348, cmp2351, cmp2354, cmp2357, cmp2360, cmp2363, cmp2366, loadedv2370, cmp2374, cmp2378, cmp2381, cmp2384, cmp2387, cmp2390, cmp2393, cmp2396, loadedv2400, cmp2404, cmp2408, cmp2411, cmp2414, cmp2417, cmp2420, cmp2423, cmp2426, loadedv2430, cmp2434, cmp2438, cmp2441, cmp2444, cmp2447, cmp2450, cmp2453, cmp2456, loadedv2460, cmp2464, cmp2468, cmp2471, cmp2474, cmp2477, cmp2480, cmp2483, cmp2486, loadedv2490, cmp2494, cmp2498, cmp2501, cmp2504, cmp2507, cmp2510, cmp2513, cmp2516, loadedv2520, cmp2524, cmp2528, cmp2531, cmp2534, cmp2537, cmp2540, cmp2543, cmp2546, loadedv2550, cmp2554, cmp2558, cmp2561, cmp2564, cmp2567, cmp2570, cmp2573, cmp2576, loadedv2580, cmp2584, cmp2588, cmp2591, cmp2594, cmp2597, cmp2600, cmp2603, cmp2606, loadedv2610, cmp2614, cmp2618, cmp2621, cmp2624, cmp2627, cmp2630, cmp2633, cmp2636, loadedv2640, cmp2644, cmp2648, cmp2651, cmp2654, cmp2657, cmp2660, cmp2663, cmp2666, loadedv2670, cmp2674, cmp2678, cmp2681, cmp2684, cmp2687, cmp2690, cmp2693, cmp2696, loadedv2700, cmp2704, cmp2708, cmp2711, cmp2714, cmp2717, cmp2720, cmp2723, cmp2726, loadedv2730, cmp2734, cmp2738, cmp2741, cmp2744, cmp2747, cmp2750, cmp2753, cmp2756, loadedv2760, cmp2764, cmp2768, cmp2771, cmp2774, cmp2777, cmp2780, cmp2783, cmp2786, loadedv2790, cmp2794, cmp2798, cmp2801, cmp2804, cmp2807, cmp2810, cmp2813, cmp2816, loadedv2820, cmp2824, cmp2827, cmp2830, cmp2833, cmp2836, cmp2839, cmp2842, loadedv2846, cmp2850, cmp2853, cmp2857, cmp2860, cmp2863, cmp2866, cmp2869, cmp2872, cmp2875, loadedv2879, cmp2883, cmp2886, cmp2889, cmp2892, cmp2895, cmp2898, cmp2901, loadedv2905, cmp2909, cmp2913, cmp2916, loadedv2920, cmp2924, cmp2927, loadedv2931, cmp2935, cmp2938, loadedv2942, v1296 bool
	var retval unsafe.Pointer
	var v9, v13, v16 int16
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol120, result_symbol128, result_symbol154, result_symbol180, result_symbol206, result_symbol232, result_symbol258, result_symbol284, result_symbol310, result_symbol336, result_symbol362, result_symbol388, result_symbol414, result_symbol444, result_symbol474, result_symbol504, result_symbol534, result_symbol564, result_symbol594, result_symbol624, result_symbol654, result_symbol684, result_symbol714, result_symbol744, result_symbol774, result_symbol804, result_symbol834, result_symbol864, result_symbol894, result_symbol924, result_symbol954, result_symbol984, result_symbol1014, result_symbol1044, result_symbol1074, result_symbol1104, result_symbol1134, result_symbol1164, result_symbol1194, result_symbol1224, result_symbol1254, result_symbol1284, result_symbol1314, result_symbol1344, result_symbol1374, result_symbol1404, result_symbol1434, result_symbol1468, result_symbol1498, result_symbol1528, result_symbol1558, result_symbol1588, result_symbol1618, result_symbol1648, result_symbol1678, result_symbol1708, result_symbol1738, result_symbol1768, result_symbol1798, result_symbol1828, result_symbol1858, result_symbol1888, result_symbol1918, result_symbol1948, result_symbol1978, result_symbol2008, result_symbol2038, result_symbol2068, result_symbol2098, result_symbol2128, result_symbol2162, result_symbol2192, result_symbol2222, result_symbol2252, result_symbol2282, result_symbol2312, result_symbol2342, result_symbol2372, result_symbol2402, result_symbol2432, result_symbol2462, result_symbol2492, result_symbol2522, result_symbol2552, result_symbol2582, result_symbol2612, result_symbol2642, result_symbol2672, result_symbol2702, result_symbol2732, result_symbol2762, result_symbol2792, result_symbol2822, result_symbol2848, result_symbol2881, result_symbol2907, result_symbol2922, result_symbol2933 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v29, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v42, v43, v45, v46, v48, v49, v60, v66, v67, v68, v69, v70, v71, v72, v78, v79, v80, v81, v82, v83, v84, v90, v91, v92, v93, v94, v95, v96, v102, v103, v104, v105, v106, v107, v108, v114, v115, v116, v117, v118, v119, v120, v126, v127, v128, v129, v130, v131, v132, v138, v139, v140, v141, v142, v143, v144, v150, v151, v152, v153, v154, v155, v156, v162, v163, v164, v165, v166, v167, v168, v174, v175, v176, v177, v178, v179, v180, v186, v187, v188, v189, v190, v191, v192, v198, v199, v200, v201, v202, v203, v204, v205, v211, v212, v213, v214, v215, v216, v217, v218, v224, v225, v226, v227, v228, v229, v230, v231, v237, v238, v239, v240, v241, v242, v243, v244, v250, v251, v252, v253, v254, v255, v256, v257, v263, v264, v265, v266, v267, v268, v269, v270, v276, v277, v278, v279, v280, v281, v282, v283, v289, v290, v291, v292, v293, v294, v295, v296, v302, v303, v304, v305, v306, v307, v308, v309, v315, v316, v317, v318, v319, v320, v321, v322, v328, v329, v330, v331, v332, v333, v334, v335, v341, v342, v343, v344, v345, v346, v347, v348, v354, v355, v356, v357, v358, v359, v360, v361, v367, v368, v369, v370, v371, v372, v373, v374, v380, v381, v382, v383, v384, v385, v386, v387, v393, v394, v395, v396, v397, v398, v399, v400, v406, v407, v408, v409, v410, v411, v412, v413, v419, v420, v421, v422, v423, v424, v425, v426, v432, v433, v434, v435, v436, v437, v438, v439, v445, v446, v447, v448, v449, v450, v451, v452, v458, v459, v460, v461, v462, v463, v464, v465, v471, v472, v473, v474, v475, v476, v477, v478, v484, v485, v486, v487, v488, v489, v490, v491, v497, v498, v499, v500, v501, v502, v503, v504, v510, v511, v512, v513, v514, v515, v516, v517, v523, v524, v525, v526, v527, v528, v529, v530, v536, v537, v538, v539, v540, v541, v542, v543, v549, v550, v551, v552, v553, v554, v555, v556, v562, v563, v564, v565, v566, v567, v568, v569, v575, v576, v577, v578, v579, v580, v581, v582, v588, v589, v590, v591, v592, v593, v594, v595, v601, v602, v603, v604, v605, v606, v607, v608, v614, v615, v616, v617, v618, v619, v620, v621, v627, v628, v629, v630, v631, v632, v633, v634, v640, v641, v642, v643, v644, v645, v646, v647, v648, v654, v655, v656, v657, v658, v659, v660, v661, v667, v668, v669, v670, v671, v672, v673, v674, v680, v681, v682, v683, v684, v685, v686, v687, v693, v694, v695, v696, v697, v698, v699, v700, v706, v707, v708, v709, v710, v711, v712, v713, v719, v720, v721, v722, v723, v724, v725, v726, v732, v733, v734, v735, v736, v737, v738, v739, v745, v746, v747, v748, v749, v750, v751, v752, v758, v759, v760, v761, v762, v763, v764, v765, v771, v772, v773, v774, v775, v776, v777, v778, v784, v785, v786, v787, v788, v789, v790, v791, v797, v798, v799, v800, v801, v802, v803, v804, v810, v811, v812, v813, v814, v815, v816, v817, v823, v824, v825, v826, v827, v828, v829, v830, v836, v837, v838, v839, v840, v841, v842, v843, v849, v850, v851, v852, v853, v854, v855, v856, v862, v863, v864, v865, v866, v867, v868, v869, v875, v876, v877, v878, v879, v880, v881, v882, v888, v889, v890, v891, v892, v893, v894, v895, v901, v902, v903, v904, v905, v906, v907, v908, v914, v915, v916, v917, v918, v919, v920, v921, v927, v928, v929, v930, v931, v932, v933, v934, v940, v941, v942, v943, v944, v945, v946, v947, v948, v954, v955, v956, v957, v958, v959, v960, v961, v967, v968, v969, v970, v971, v972, v973, v974, v980, v981, v982, v983, v984, v985, v986, v987, v993, v994, v995, v996, v997, v998, v999, v1000, v1006, v1007, v1008, v1009, v1010, v1011, v1012, v1013, v1019, v1020, v1021, v1022, v1023, v1024, v1025, v1026, v1032, v1033, v1034, v1035, v1036, v1037, v1038, v1039, v1045, v1046, v1047, v1048, v1049, v1050, v1051, v1052, v1058, v1059, v1060, v1061, v1062, v1063, v1064, v1065, v1071, v1072, v1073, v1074, v1075, v1076, v1077, v1078, v1084, v1085, v1086, v1087, v1088, v1089, v1090, v1091, v1097, v1098, v1099, v1100, v1101, v1102, v1103, v1104, v1110, v1111, v1112, v1113, v1114, v1115, v1116, v1117, v1123, v1124, v1125, v1126, v1127, v1128, v1129, v1130, v1136, v1137, v1138, v1139, v1140, v1141, v1142, v1143, v1149, v1150, v1151, v1152, v1153, v1154, v1155, v1156, v1162, v1163, v1164, v1165, v1166, v1167, v1168, v1169, v1175, v1176, v1177, v1178, v1179, v1180, v1181, v1182, v1188, v1189, v1190, v1191, v1192, v1193, v1194, v1195, v1201, v1202, v1203, v1204, v1205, v1206, v1207, v1208, v1214, v1215, v1216, v1217, v1218, v1219, v1220, v1221, v1227, v1228, v1229, v1230, v1231, v1232, v1233, v1234, v1240, v1241, v1242, v1243, v1244, v1245, v1246, v1252, v1253, v1254, v1255, v1256, v1257, v1258, v1259, v1260, v1266, v1267, v1268, v1269, v1270, v1271, v1272, v1278, v1279, v1280, v1286, v1287, v1293, v1294 int32
	var lookahead, i, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10 int64
	var v3, storedv, v10, v28, v44, v47, v50, v55, v61, v73, v85, v97, v109, v121, v133, v145, v157, v169, v181, v193, v206, v219, v232, v245, v258, v271, v284, v297, v310, v323, v336, v349, v362, v375, v388, v401, v414, v427, v440, v453, v466, v479, v492, v505, v518, v531, v544, v557, v570, v583, v596, v609, v622, v635, v649, v662, v675, v688, v701, v714, v727, v740, v753, v766, v779, v792, v805, v818, v831, v844, v857, v870, v883, v896, v909, v922, v935, v949, v962, v975, v988, v1001, v1014, v1027, v1040, v1053, v1066, v1079, v1092, v1105, v1118, v1131, v1144, v1157, v1170, v1183, v1196, v1209, v1222, v1235, v1247, v1261, v1273, v1281, v1288, v1295 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v51, v52, v53, v54, v56, v57, v58, v59, v62, v63, v64, v65, v74, v75, v76, v77, v86, v87, v88, v89, v98, v99, v100, v101, v110, v111, v112, v113, v122, v123, v124, v125, v134, v135, v136, v137, v146, v147, v148, v149, v158, v159, v160, v161, v170, v171, v172, v173, v182, v183, v184, v185, v194, v195, v196, v197, v207, v208, v209, v210, v220, v221, v222, v223, v233, v234, v235, v236, v246, v247, v248, v249, v259, v260, v261, v262, v272, v273, v274, v275, v285, v286, v287, v288, v298, v299, v300, v301, v311, v312, v313, v314, v324, v325, v326, v327, v337, v338, v339, v340, v350, v351, v352, v353, v363, v364, v365, v366, v376, v377, v378, v379, v389, v390, v391, v392, v402, v403, v404, v405, v415, v416, v417, v418, v428, v429, v430, v431, v441, v442, v443, v444, v454, v455, v456, v457, v467, v468, v469, v470, v480, v481, v482, v483, v493, v494, v495, v496, v506, v507, v508, v509, v519, v520, v521, v522, v532, v533, v534, v535, v545, v546, v547, v548, v558, v559, v560, v561, v571, v572, v573, v574, v584, v585, v586, v587, v597, v598, v599, v600, v610, v611, v612, v613, v623, v624, v625, v626, v636, v637, v638, v639, v650, v651, v652, v653, v663, v664, v665, v666, v676, v677, v678, v679, v689, v690, v691, v692, v702, v703, v704, v705, v715, v716, v717, v718, v728, v729, v730, v731, v741, v742, v743, v744, v754, v755, v756, v757, v767, v768, v769, v770, v780, v781, v782, v783, v793, v794, v795, v796, v806, v807, v808, v809, v819, v820, v821, v822, v832, v833, v834, v835, v845, v846, v847, v848, v858, v859, v860, v861, v871, v872, v873, v874, v884, v885, v886, v887, v897, v898, v899, v900, v910, v911, v912, v913, v923, v924, v925, v926, v936, v937, v938, v939, v950, v951, v952, v953, v963, v964, v965, v966, v976, v977, v978, v979, v989, v990, v991, v992, v1002, v1003, v1004, v1005, v1015, v1016, v1017, v1018, v1028, v1029, v1030, v1031, v1041, v1042, v1043, v1044, v1054, v1055, v1056, v1057, v1067, v1068, v1069, v1070, v1080, v1081, v1082, v1083, v1093, v1094, v1095, v1096, v1106, v1107, v1108, v1109, v1119, v1120, v1121, v1122, v1132, v1133, v1134, v1135, v1145, v1146, v1147, v1148, v1158, v1159, v1160, v1161, v1171, v1172, v1173, v1174, v1184, v1185, v1186, v1187, v1197, v1198, v1199, v1200, v1210, v1211, v1212, v1213, v1223, v1224, v1225, v1226, v1236, v1237, v1238, v1239, v1248, v1249, v1250, v1251, v1262, v1263, v1264, v1265, v1274, v1275, v1276, v1277, v1282, v1283, v1284, v1285, v1289, v1290, v1291, v1292 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end121, mark_end129, mark_end155, mark_end181, mark_end207, mark_end233, mark_end259, mark_end285, mark_end311, mark_end337, mark_end363, mark_end389, mark_end415, mark_end445, mark_end475, mark_end505, mark_end535, mark_end565, mark_end595, mark_end625, mark_end655, mark_end685, mark_end715, mark_end745, mark_end775, mark_end805, mark_end835, mark_end865, mark_end895, mark_end925, mark_end955, mark_end985, mark_end1015, mark_end1045, mark_end1075, mark_end1105, mark_end1135, mark_end1165, mark_end1195, mark_end1225, mark_end1255, mark_end1285, mark_end1315, mark_end1345, mark_end1375, mark_end1405, mark_end1435, mark_end1469, mark_end1499, mark_end1529, mark_end1559, mark_end1589, mark_end1619, mark_end1649, mark_end1679, mark_end1709, mark_end1739, mark_end1769, mark_end1799, mark_end1829, mark_end1859, mark_end1889, mark_end1919, mark_end1949, mark_end1979, mark_end2009, mark_end2039, mark_end2069, mark_end2099, mark_end2129, mark_end2163, mark_end2193, mark_end2223, mark_end2253, mark_end2283, mark_end2313, mark_end2343, mark_end2373, mark_end2403, mark_end2433, mark_end2463, mark_end2493, mark_end2523, mark_end2553, mark_end2583, mark_end2613, mark_end2643, mark_end2673, mark_end2703, mark_end2733, mark_end2763, mark_end2793, mark_end2823, mark_end2849, mark_end2882, mark_end2908, mark_end2923, mark_end2934 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, cmp29, v24, cmp32, v25, cmp35, v26, cmp38, v27, cmp41, v28, loadedv45, v29, cmp47, v30, cmp51, v31, cmp55, v32, cmp59, v33, cmp62, v34, cmp65, v35, cmp69, v36, cmp72, v37, cmp76, v38, cmp79, v39, cmp82, v40, cmp85, v41, cmp88, v42, cmp91, v43, cmp94, v44, loadedv98, v45, cmp100, v46, cmp103, v47, loadedv107, v48, cmp109, v49, cmp112, v50, loadedv116, v51, result_symbol, v52, mark_end, v53, v54, v55, loadedv118, v56, result_symbol120, v57, mark_end121, v58, v59, v60, cmp122, v61, loadedv126, v62, result_symbol128, v63, mark_end129, v64, v65, v66, cmp130, v67, cmp133, v68, cmp136, v69, cmp139, v70, cmp142, v71, cmp145, v72, cmp148, v73, loadedv152, v74, result_symbol154, v75, mark_end155, v76, v77, v78, cmp156, v79, cmp159, v80, cmp162, v81, cmp165, v82, cmp168, v83, cmp171, v84, cmp174, v85, loadedv178, v86, result_symbol180, v87, mark_end181, v88, v89, v90, cmp182, v91, cmp185, v92, cmp188, v93, cmp191, v94, cmp194, v95, cmp197, v96, cmp200, v97, loadedv204, v98, result_symbol206, v99, mark_end207, v100, v101, v102, cmp208, v103, cmp211, v104, cmp214, v105, cmp217, v106, cmp220, v107, cmp223, v108, cmp226, v109, loadedv230, v110, result_symbol232, v111, mark_end233, v112, v113, v114, cmp234, v115, cmp237, v116, cmp240, v117, cmp243, v118, cmp246, v119, cmp249, v120, cmp252, v121, loadedv256, v122, result_symbol258, v123, mark_end259, v124, v125, v126, cmp260, v127, cmp263, v128, cmp266, v129, cmp269, v130, cmp272, v131, cmp275, v132, cmp278, v133, loadedv282, v134, result_symbol284, v135, mark_end285, v136, v137, v138, cmp286, v139, cmp289, v140, cmp292, v141, cmp295, v142, cmp298, v143, cmp301, v144, cmp304, v145, loadedv308, v146, result_symbol310, v147, mark_end311, v148, v149, v150, cmp312, v151, cmp315, v152, cmp318, v153, cmp321, v154, cmp324, v155, cmp327, v156, cmp330, v157, loadedv334, v158, result_symbol336, v159, mark_end337, v160, v161, v162, cmp338, v163, cmp341, v164, cmp344, v165, cmp347, v166, cmp350, v167, cmp353, v168, cmp356, v169, loadedv360, v170, result_symbol362, v171, mark_end363, v172, v173, v174, cmp364, v175, cmp367, v176, cmp370, v177, cmp373, v178, cmp376, v179, cmp379, v180, cmp382, v181, loadedv386, v182, result_symbol388, v183, mark_end389, v184, v185, v186, cmp390, v187, cmp393, v188, cmp396, v189, cmp399, v190, cmp402, v191, cmp405, v192, cmp408, v193, loadedv412, v194, result_symbol414, v195, mark_end415, v196, v197, v198, cmp416, v199, cmp420, v200, cmp423, v201, cmp426, v202, cmp429, v203, cmp432, v204, cmp435, v205, cmp438, v206, loadedv442, v207, result_symbol444, v208, mark_end445, v209, v210, v211, cmp446, v212, cmp450, v213, cmp453, v214, cmp456, v215, cmp459, v216, cmp462, v217, cmp465, v218, cmp468, v219, loadedv472, v220, result_symbol474, v221, mark_end475, v222, v223, v224, cmp476, v225, cmp480, v226, cmp483, v227, cmp486, v228, cmp489, v229, cmp492, v230, cmp495, v231, cmp498, v232, loadedv502, v233, result_symbol504, v234, mark_end505, v235, v236, v237, cmp506, v238, cmp510, v239, cmp513, v240, cmp516, v241, cmp519, v242, cmp522, v243, cmp525, v244, cmp528, v245, loadedv532, v246, result_symbol534, v247, mark_end535, v248, v249, v250, cmp536, v251, cmp540, v252, cmp543, v253, cmp546, v254, cmp549, v255, cmp552, v256, cmp555, v257, cmp558, v258, loadedv562, v259, result_symbol564, v260, mark_end565, v261, v262, v263, cmp566, v264, cmp570, v265, cmp573, v266, cmp576, v267, cmp579, v268, cmp582, v269, cmp585, v270, cmp588, v271, loadedv592, v272, result_symbol594, v273, mark_end595, v274, v275, v276, cmp596, v277, cmp600, v278, cmp603, v279, cmp606, v280, cmp609, v281, cmp612, v282, cmp615, v283, cmp618, v284, loadedv622, v285, result_symbol624, v286, mark_end625, v287, v288, v289, cmp626, v290, cmp630, v291, cmp633, v292, cmp636, v293, cmp639, v294, cmp642, v295, cmp645, v296, cmp648, v297, loadedv652, v298, result_symbol654, v299, mark_end655, v300, v301, v302, cmp656, v303, cmp660, v304, cmp663, v305, cmp666, v306, cmp669, v307, cmp672, v308, cmp675, v309, cmp678, v310, loadedv682, v311, result_symbol684, v312, mark_end685, v313, v314, v315, cmp686, v316, cmp690, v317, cmp693, v318, cmp696, v319, cmp699, v320, cmp702, v321, cmp705, v322, cmp708, v323, loadedv712, v324, result_symbol714, v325, mark_end715, v326, v327, v328, cmp716, v329, cmp720, v330, cmp723, v331, cmp726, v332, cmp729, v333, cmp732, v334, cmp735, v335, cmp738, v336, loadedv742, v337, result_symbol744, v338, mark_end745, v339, v340, v341, cmp746, v342, cmp750, v343, cmp753, v344, cmp756, v345, cmp759, v346, cmp762, v347, cmp765, v348, cmp768, v349, loadedv772, v350, result_symbol774, v351, mark_end775, v352, v353, v354, cmp776, v355, cmp780, v356, cmp783, v357, cmp786, v358, cmp789, v359, cmp792, v360, cmp795, v361, cmp798, v362, loadedv802, v363, result_symbol804, v364, mark_end805, v365, v366, v367, cmp806, v368, cmp810, v369, cmp813, v370, cmp816, v371, cmp819, v372, cmp822, v373, cmp825, v374, cmp828, v375, loadedv832, v376, result_symbol834, v377, mark_end835, v378, v379, v380, cmp836, v381, cmp840, v382, cmp843, v383, cmp846, v384, cmp849, v385, cmp852, v386, cmp855, v387, cmp858, v388, loadedv862, v389, result_symbol864, v390, mark_end865, v391, v392, v393, cmp866, v394, cmp870, v395, cmp873, v396, cmp876, v397, cmp879, v398, cmp882, v399, cmp885, v400, cmp888, v401, loadedv892, v402, result_symbol894, v403, mark_end895, v404, v405, v406, cmp896, v407, cmp900, v408, cmp903, v409, cmp906, v410, cmp909, v411, cmp912, v412, cmp915, v413, cmp918, v414, loadedv922, v415, result_symbol924, v416, mark_end925, v417, v418, v419, cmp926, v420, cmp930, v421, cmp933, v422, cmp936, v423, cmp939, v424, cmp942, v425, cmp945, v426, cmp948, v427, loadedv952, v428, result_symbol954, v429, mark_end955, v430, v431, v432, cmp956, v433, cmp960, v434, cmp963, v435, cmp966, v436, cmp969, v437, cmp972, v438, cmp975, v439, cmp978, v440, loadedv982, v441, result_symbol984, v442, mark_end985, v443, v444, v445, cmp986, v446, cmp990, v447, cmp993, v448, cmp996, v449, cmp999, v450, cmp1002, v451, cmp1005, v452, cmp1008, v453, loadedv1012, v454, result_symbol1014, v455, mark_end1015, v456, v457, v458, cmp1016, v459, cmp1020, v460, cmp1023, v461, cmp1026, v462, cmp1029, v463, cmp1032, v464, cmp1035, v465, cmp1038, v466, loadedv1042, v467, result_symbol1044, v468, mark_end1045, v469, v470, v471, cmp1046, v472, cmp1050, v473, cmp1053, v474, cmp1056, v475, cmp1059, v476, cmp1062, v477, cmp1065, v478, cmp1068, v479, loadedv1072, v480, result_symbol1074, v481, mark_end1075, v482, v483, v484, cmp1076, v485, cmp1080, v486, cmp1083, v487, cmp1086, v488, cmp1089, v489, cmp1092, v490, cmp1095, v491, cmp1098, v492, loadedv1102, v493, result_symbol1104, v494, mark_end1105, v495, v496, v497, cmp1106, v498, cmp1110, v499, cmp1113, v500, cmp1116, v501, cmp1119, v502, cmp1122, v503, cmp1125, v504, cmp1128, v505, loadedv1132, v506, result_symbol1134, v507, mark_end1135, v508, v509, v510, cmp1136, v511, cmp1140, v512, cmp1143, v513, cmp1146, v514, cmp1149, v515, cmp1152, v516, cmp1155, v517, cmp1158, v518, loadedv1162, v519, result_symbol1164, v520, mark_end1165, v521, v522, v523, cmp1166, v524, cmp1170, v525, cmp1173, v526, cmp1176, v527, cmp1179, v528, cmp1182, v529, cmp1185, v530, cmp1188, v531, loadedv1192, v532, result_symbol1194, v533, mark_end1195, v534, v535, v536, cmp1196, v537, cmp1200, v538, cmp1203, v539, cmp1206, v540, cmp1209, v541, cmp1212, v542, cmp1215, v543, cmp1218, v544, loadedv1222, v545, result_symbol1224, v546, mark_end1225, v547, v548, v549, cmp1226, v550, cmp1230, v551, cmp1233, v552, cmp1236, v553, cmp1239, v554, cmp1242, v555, cmp1245, v556, cmp1248, v557, loadedv1252, v558, result_symbol1254, v559, mark_end1255, v560, v561, v562, cmp1256, v563, cmp1260, v564, cmp1263, v565, cmp1266, v566, cmp1269, v567, cmp1272, v568, cmp1275, v569, cmp1278, v570, loadedv1282, v571, result_symbol1284, v572, mark_end1285, v573, v574, v575, cmp1286, v576, cmp1290, v577, cmp1293, v578, cmp1296, v579, cmp1299, v580, cmp1302, v581, cmp1305, v582, cmp1308, v583, loadedv1312, v584, result_symbol1314, v585, mark_end1315, v586, v587, v588, cmp1316, v589, cmp1320, v590, cmp1323, v591, cmp1326, v592, cmp1329, v593, cmp1332, v594, cmp1335, v595, cmp1338, v596, loadedv1342, v597, result_symbol1344, v598, mark_end1345, v599, v600, v601, cmp1346, v602, cmp1350, v603, cmp1353, v604, cmp1356, v605, cmp1359, v606, cmp1362, v607, cmp1365, v608, cmp1368, v609, loadedv1372, v610, result_symbol1374, v611, mark_end1375, v612, v613, v614, cmp1376, v615, cmp1380, v616, cmp1383, v617, cmp1386, v618, cmp1389, v619, cmp1392, v620, cmp1395, v621, cmp1398, v622, loadedv1402, v623, result_symbol1404, v624, mark_end1405, v625, v626, v627, cmp1406, v628, cmp1410, v629, cmp1413, v630, cmp1416, v631, cmp1419, v632, cmp1422, v633, cmp1425, v634, cmp1428, v635, loadedv1432, v636, result_symbol1434, v637, mark_end1435, v638, v639, v640, cmp1436, v641, cmp1440, v642, cmp1444, v643, cmp1447, v644, cmp1450, v645, cmp1453, v646, cmp1456, v647, cmp1459, v648, cmp1462, v649, loadedv1466, v650, result_symbol1468, v651, mark_end1469, v652, v653, v654, cmp1470, v655, cmp1474, v656, cmp1477, v657, cmp1480, v658, cmp1483, v659, cmp1486, v660, cmp1489, v661, cmp1492, v662, loadedv1496, v663, result_symbol1498, v664, mark_end1499, v665, v666, v667, cmp1500, v668, cmp1504, v669, cmp1507, v670, cmp1510, v671, cmp1513, v672, cmp1516, v673, cmp1519, v674, cmp1522, v675, loadedv1526, v676, result_symbol1528, v677, mark_end1529, v678, v679, v680, cmp1530, v681, cmp1534, v682, cmp1537, v683, cmp1540, v684, cmp1543, v685, cmp1546, v686, cmp1549, v687, cmp1552, v688, loadedv1556, v689, result_symbol1558, v690, mark_end1559, v691, v692, v693, cmp1560, v694, cmp1564, v695, cmp1567, v696, cmp1570, v697, cmp1573, v698, cmp1576, v699, cmp1579, v700, cmp1582, v701, loadedv1586, v702, result_symbol1588, v703, mark_end1589, v704, v705, v706, cmp1590, v707, cmp1594, v708, cmp1597, v709, cmp1600, v710, cmp1603, v711, cmp1606, v712, cmp1609, v713, cmp1612, v714, loadedv1616, v715, result_symbol1618, v716, mark_end1619, v717, v718, v719, cmp1620, v720, cmp1624, v721, cmp1627, v722, cmp1630, v723, cmp1633, v724, cmp1636, v725, cmp1639, v726, cmp1642, v727, loadedv1646, v728, result_symbol1648, v729, mark_end1649, v730, v731, v732, cmp1650, v733, cmp1654, v734, cmp1657, v735, cmp1660, v736, cmp1663, v737, cmp1666, v738, cmp1669, v739, cmp1672, v740, loadedv1676, v741, result_symbol1678, v742, mark_end1679, v743, v744, v745, cmp1680, v746, cmp1684, v747, cmp1687, v748, cmp1690, v749, cmp1693, v750, cmp1696, v751, cmp1699, v752, cmp1702, v753, loadedv1706, v754, result_symbol1708, v755, mark_end1709, v756, v757, v758, cmp1710, v759, cmp1714, v760, cmp1717, v761, cmp1720, v762, cmp1723, v763, cmp1726, v764, cmp1729, v765, cmp1732, v766, loadedv1736, v767, result_symbol1738, v768, mark_end1739, v769, v770, v771, cmp1740, v772, cmp1744, v773, cmp1747, v774, cmp1750, v775, cmp1753, v776, cmp1756, v777, cmp1759, v778, cmp1762, v779, loadedv1766, v780, result_symbol1768, v781, mark_end1769, v782, v783, v784, cmp1770, v785, cmp1774, v786, cmp1777, v787, cmp1780, v788, cmp1783, v789, cmp1786, v790, cmp1789, v791, cmp1792, v792, loadedv1796, v793, result_symbol1798, v794, mark_end1799, v795, v796, v797, cmp1800, v798, cmp1804, v799, cmp1807, v800, cmp1810, v801, cmp1813, v802, cmp1816, v803, cmp1819, v804, cmp1822, v805, loadedv1826, v806, result_symbol1828, v807, mark_end1829, v808, v809, v810, cmp1830, v811, cmp1834, v812, cmp1837, v813, cmp1840, v814, cmp1843, v815, cmp1846, v816, cmp1849, v817, cmp1852, v818, loadedv1856, v819, result_symbol1858, v820, mark_end1859, v821, v822, v823, cmp1860, v824, cmp1864, v825, cmp1867, v826, cmp1870, v827, cmp1873, v828, cmp1876, v829, cmp1879, v830, cmp1882, v831, loadedv1886, v832, result_symbol1888, v833, mark_end1889, v834, v835, v836, cmp1890, v837, cmp1894, v838, cmp1897, v839, cmp1900, v840, cmp1903, v841, cmp1906, v842, cmp1909, v843, cmp1912, v844, loadedv1916, v845, result_symbol1918, v846, mark_end1919, v847, v848, v849, cmp1920, v850, cmp1924, v851, cmp1927, v852, cmp1930, v853, cmp1933, v854, cmp1936, v855, cmp1939, v856, cmp1942, v857, loadedv1946, v858, result_symbol1948, v859, mark_end1949, v860, v861, v862, cmp1950, v863, cmp1954, v864, cmp1957, v865, cmp1960, v866, cmp1963, v867, cmp1966, v868, cmp1969, v869, cmp1972, v870, loadedv1976, v871, result_symbol1978, v872, mark_end1979, v873, v874, v875, cmp1980, v876, cmp1984, v877, cmp1987, v878, cmp1990, v879, cmp1993, v880, cmp1996, v881, cmp1999, v882, cmp2002, v883, loadedv2006, v884, result_symbol2008, v885, mark_end2009, v886, v887, v888, cmp2010, v889, cmp2014, v890, cmp2017, v891, cmp2020, v892, cmp2023, v893, cmp2026, v894, cmp2029, v895, cmp2032, v896, loadedv2036, v897, result_symbol2038, v898, mark_end2039, v899, v900, v901, cmp2040, v902, cmp2044, v903, cmp2047, v904, cmp2050, v905, cmp2053, v906, cmp2056, v907, cmp2059, v908, cmp2062, v909, loadedv2066, v910, result_symbol2068, v911, mark_end2069, v912, v913, v914, cmp2070, v915, cmp2074, v916, cmp2077, v917, cmp2080, v918, cmp2083, v919, cmp2086, v920, cmp2089, v921, cmp2092, v922, loadedv2096, v923, result_symbol2098, v924, mark_end2099, v925, v926, v927, cmp2100, v928, cmp2104, v929, cmp2107, v930, cmp2110, v931, cmp2113, v932, cmp2116, v933, cmp2119, v934, cmp2122, v935, loadedv2126, v936, result_symbol2128, v937, mark_end2129, v938, v939, v940, cmp2130, v941, cmp2134, v942, cmp2138, v943, cmp2141, v944, cmp2144, v945, cmp2147, v946, cmp2150, v947, cmp2153, v948, cmp2156, v949, loadedv2160, v950, result_symbol2162, v951, mark_end2163, v952, v953, v954, cmp2164, v955, cmp2168, v956, cmp2171, v957, cmp2174, v958, cmp2177, v959, cmp2180, v960, cmp2183, v961, cmp2186, v962, loadedv2190, v963, result_symbol2192, v964, mark_end2193, v965, v966, v967, cmp2194, v968, cmp2198, v969, cmp2201, v970, cmp2204, v971, cmp2207, v972, cmp2210, v973, cmp2213, v974, cmp2216, v975, loadedv2220, v976, result_symbol2222, v977, mark_end2223, v978, v979, v980, cmp2224, v981, cmp2228, v982, cmp2231, v983, cmp2234, v984, cmp2237, v985, cmp2240, v986, cmp2243, v987, cmp2246, v988, loadedv2250, v989, result_symbol2252, v990, mark_end2253, v991, v992, v993, cmp2254, v994, cmp2258, v995, cmp2261, v996, cmp2264, v997, cmp2267, v998, cmp2270, v999, cmp2273, v1000, cmp2276, v1001, loadedv2280, v1002, result_symbol2282, v1003, mark_end2283, v1004, v1005, v1006, cmp2284, v1007, cmp2288, v1008, cmp2291, v1009, cmp2294, v1010, cmp2297, v1011, cmp2300, v1012, cmp2303, v1013, cmp2306, v1014, loadedv2310, v1015, result_symbol2312, v1016, mark_end2313, v1017, v1018, v1019, cmp2314, v1020, cmp2318, v1021, cmp2321, v1022, cmp2324, v1023, cmp2327, v1024, cmp2330, v1025, cmp2333, v1026, cmp2336, v1027, loadedv2340, v1028, result_symbol2342, v1029, mark_end2343, v1030, v1031, v1032, cmp2344, v1033, cmp2348, v1034, cmp2351, v1035, cmp2354, v1036, cmp2357, v1037, cmp2360, v1038, cmp2363, v1039, cmp2366, v1040, loadedv2370, v1041, result_symbol2372, v1042, mark_end2373, v1043, v1044, v1045, cmp2374, v1046, cmp2378, v1047, cmp2381, v1048, cmp2384, v1049, cmp2387, v1050, cmp2390, v1051, cmp2393, v1052, cmp2396, v1053, loadedv2400, v1054, result_symbol2402, v1055, mark_end2403, v1056, v1057, v1058, cmp2404, v1059, cmp2408, v1060, cmp2411, v1061, cmp2414, v1062, cmp2417, v1063, cmp2420, v1064, cmp2423, v1065, cmp2426, v1066, loadedv2430, v1067, result_symbol2432, v1068, mark_end2433, v1069, v1070, v1071, cmp2434, v1072, cmp2438, v1073, cmp2441, v1074, cmp2444, v1075, cmp2447, v1076, cmp2450, v1077, cmp2453, v1078, cmp2456, v1079, loadedv2460, v1080, result_symbol2462, v1081, mark_end2463, v1082, v1083, v1084, cmp2464, v1085, cmp2468, v1086, cmp2471, v1087, cmp2474, v1088, cmp2477, v1089, cmp2480, v1090, cmp2483, v1091, cmp2486, v1092, loadedv2490, v1093, result_symbol2492, v1094, mark_end2493, v1095, v1096, v1097, cmp2494, v1098, cmp2498, v1099, cmp2501, v1100, cmp2504, v1101, cmp2507, v1102, cmp2510, v1103, cmp2513, v1104, cmp2516, v1105, loadedv2520, v1106, result_symbol2522, v1107, mark_end2523, v1108, v1109, v1110, cmp2524, v1111, cmp2528, v1112, cmp2531, v1113, cmp2534, v1114, cmp2537, v1115, cmp2540, v1116, cmp2543, v1117, cmp2546, v1118, loadedv2550, v1119, result_symbol2552, v1120, mark_end2553, v1121, v1122, v1123, cmp2554, v1124, cmp2558, v1125, cmp2561, v1126, cmp2564, v1127, cmp2567, v1128, cmp2570, v1129, cmp2573, v1130, cmp2576, v1131, loadedv2580, v1132, result_symbol2582, v1133, mark_end2583, v1134, v1135, v1136, cmp2584, v1137, cmp2588, v1138, cmp2591, v1139, cmp2594, v1140, cmp2597, v1141, cmp2600, v1142, cmp2603, v1143, cmp2606, v1144, loadedv2610, v1145, result_symbol2612, v1146, mark_end2613, v1147, v1148, v1149, cmp2614, v1150, cmp2618, v1151, cmp2621, v1152, cmp2624, v1153, cmp2627, v1154, cmp2630, v1155, cmp2633, v1156, cmp2636, v1157, loadedv2640, v1158, result_symbol2642, v1159, mark_end2643, v1160, v1161, v1162, cmp2644, v1163, cmp2648, v1164, cmp2651, v1165, cmp2654, v1166, cmp2657, v1167, cmp2660, v1168, cmp2663, v1169, cmp2666, v1170, loadedv2670, v1171, result_symbol2672, v1172, mark_end2673, v1173, v1174, v1175, cmp2674, v1176, cmp2678, v1177, cmp2681, v1178, cmp2684, v1179, cmp2687, v1180, cmp2690, v1181, cmp2693, v1182, cmp2696, v1183, loadedv2700, v1184, result_symbol2702, v1185, mark_end2703, v1186, v1187, v1188, cmp2704, v1189, cmp2708, v1190, cmp2711, v1191, cmp2714, v1192, cmp2717, v1193, cmp2720, v1194, cmp2723, v1195, cmp2726, v1196, loadedv2730, v1197, result_symbol2732, v1198, mark_end2733, v1199, v1200, v1201, cmp2734, v1202, cmp2738, v1203, cmp2741, v1204, cmp2744, v1205, cmp2747, v1206, cmp2750, v1207, cmp2753, v1208, cmp2756, v1209, loadedv2760, v1210, result_symbol2762, v1211, mark_end2763, v1212, v1213, v1214, cmp2764, v1215, cmp2768, v1216, cmp2771, v1217, cmp2774, v1218, cmp2777, v1219, cmp2780, v1220, cmp2783, v1221, cmp2786, v1222, loadedv2790, v1223, result_symbol2792, v1224, mark_end2793, v1225, v1226, v1227, cmp2794, v1228, cmp2798, v1229, cmp2801, v1230, cmp2804, v1231, cmp2807, v1232, cmp2810, v1233, cmp2813, v1234, cmp2816, v1235, loadedv2820, v1236, result_symbol2822, v1237, mark_end2823, v1238, v1239, v1240, cmp2824, v1241, cmp2827, v1242, cmp2830, v1243, cmp2833, v1244, cmp2836, v1245, cmp2839, v1246, cmp2842, v1247, loadedv2846, v1248, result_symbol2848, v1249, mark_end2849, v1250, v1251, v1252, cmp2850, v1253, cmp2853, v1254, cmp2857, v1255, cmp2860, v1256, cmp2863, v1257, cmp2866, v1258, cmp2869, v1259, cmp2872, v1260, cmp2875, v1261, loadedv2879, v1262, result_symbol2881, v1263, mark_end2882, v1264, v1265, v1266, cmp2883, v1267, cmp2886, v1268, cmp2889, v1269, cmp2892, v1270, cmp2895, v1271, cmp2898, v1272, cmp2901, v1273, loadedv2905, v1274, result_symbol2907, v1275, mark_end2908, v1276, v1277, v1278, cmp2909, v1279, cmp2913, v1280, cmp2916, v1281, loadedv2920, v1282, result_symbol2922, v1283, mark_end2923, v1284, v1285, v1286, cmp2924, v1287, cmp2927, v1288, loadedv2931, v1289, result_symbol2933, v1290, mark_end2934, v1291, v1292, v1293, cmp2935, v1294, cmp2938, v1295, loadedv2942, v1296

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
		goto sw_bb46
	case 2:
		goto sw_bb99
	case 3:
		goto sw_bb108
	case 4:
		goto sw_bb117
	case 5:
		goto sw_bb119
	case 6:
		goto sw_bb127
	case 7:
		goto sw_bb153
	case 8:
		goto sw_bb179
	case 9:
		goto sw_bb205
	case 10:
		goto sw_bb231
	case 11:
		goto sw_bb257
	case 12:
		goto sw_bb283
	case 13:
		goto sw_bb309
	case 14:
		goto sw_bb335
	case 15:
		goto sw_bb361
	case 16:
		goto sw_bb387
	case 17:
		goto sw_bb413
	case 18:
		goto sw_bb443
	case 19:
		goto sw_bb473
	case 20:
		goto sw_bb503
	case 21:
		goto sw_bb533
	case 22:
		goto sw_bb563
	case 23:
		goto sw_bb593
	case 24:
		goto sw_bb623
	case 25:
		goto sw_bb653
	case 26:
		goto sw_bb683
	case 27:
		goto sw_bb713
	case 28:
		goto sw_bb743
	case 29:
		goto sw_bb773
	case 30:
		goto sw_bb803
	case 31:
		goto sw_bb833
	case 32:
		goto sw_bb863
	case 33:
		goto sw_bb893
	case 34:
		goto sw_bb923
	case 35:
		goto sw_bb953
	case 36:
		goto sw_bb983
	case 37:
		goto sw_bb1013
	case 38:
		goto sw_bb1043
	case 39:
		goto sw_bb1073
	case 40:
		goto sw_bb1103
	case 41:
		goto sw_bb1133
	case 42:
		goto sw_bb1163
	case 43:
		goto sw_bb1193
	case 44:
		goto sw_bb1223
	case 45:
		goto sw_bb1253
	case 46:
		goto sw_bb1283
	case 47:
		goto sw_bb1313
	case 48:
		goto sw_bb1343
	case 49:
		goto sw_bb1373
	case 50:
		goto sw_bb1403
	case 51:
		goto sw_bb1433
	case 52:
		goto sw_bb1467
	case 53:
		goto sw_bb1497
	case 54:
		goto sw_bb1527
	case 55:
		goto sw_bb1557
	case 56:
		goto sw_bb1587
	case 57:
		goto sw_bb1617
	case 58:
		goto sw_bb1647
	case 59:
		goto sw_bb1677
	case 60:
		goto sw_bb1707
	case 61:
		goto sw_bb1737
	case 62:
		goto sw_bb1767
	case 63:
		goto sw_bb1797
	case 64:
		goto sw_bb1827
	case 65:
		goto sw_bb1857
	case 66:
		goto sw_bb1887
	case 67:
		goto sw_bb1917
	case 68:
		goto sw_bb1947
	case 69:
		goto sw_bb1977
	case 70:
		goto sw_bb2007
	case 71:
		goto sw_bb2037
	case 72:
		goto sw_bb2067
	case 73:
		goto sw_bb2097
	case 74:
		goto sw_bb2127
	case 75:
		goto sw_bb2161
	case 76:
		goto sw_bb2191
	case 77:
		goto sw_bb2221
	case 78:
		goto sw_bb2251
	case 79:
		goto sw_bb2281
	case 80:
		goto sw_bb2311
	case 81:
		goto sw_bb2341
	case 82:
		goto sw_bb2371
	case 83:
		goto sw_bb2401
	case 84:
		goto sw_bb2431
	case 85:
		goto sw_bb2461
	case 86:
		goto sw_bb2491
	case 87:
		goto sw_bb2521
	case 88:
		goto sw_bb2551
	case 89:
		goto sw_bb2581
	case 90:
		goto sw_bb2611
	case 91:
		goto sw_bb2641
	case 92:
		goto sw_bb2671
	case 93:
		goto sw_bb2701
	case 94:
		goto sw_bb2731
	case 95:
		goto sw_bb2761
	case 96:
		goto sw_bb2791
	case 97:
		goto sw_bb2821
	case 98:
		goto sw_bb2847
	case 99:
		goto sw_bb2880
	case 100:
		goto sw_bb2906
	case 101:
		goto sw_bb2921
	case 102:
		goto sw_bb2932
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
	*libc.As[int16](state_addr) = 4
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
	cmp22 = 48 <= v21
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
	*libc.As[int16](state_addr) = 100
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
		goto if_then43
	} else {
		goto lor_lhs_false34
	}

lor_lhs_false34:
	v25 = *libc.As[int32](lookahead)
	cmp35 = v25 == 95
	if cmp35 {
		goto if_then43
	} else {
		goto lor_lhs_false37
	}

lor_lhs_false37:
	v26 = *libc.As[int32](lookahead)
	cmp38 = 97 <= v26
	if cmp38 {
		goto land_lhs_true40
	} else {
		goto if_end44
	}

land_lhs_true40:
	v27 = *libc.As[int32](lookahead)
	cmp41 = v27 <= 122
	if cmp41 {
		goto if_then43
	} else {
		goto if_end44
	}

if_then43:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end44:
	v28 = *libc.As[byte](result)
	loadedv45 = (v28 & 1) != 0
	*libc.As[bool](retval) = loadedv45
	goto _return

sw_bb46:
	v29 = *libc.As[int32](lookahead)
	cmp47 = v29 == 10
	if cmp47 {
		goto if_then49
	} else {
		goto if_end50
	}

if_then49:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end50:
	v30 = *libc.As[int32](lookahead)
	cmp51 = v30 == 35
	if cmp51 {
		goto if_then53
	} else {
		goto if_end54
	}

if_then53:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end54:
	v31 = *libc.As[int32](lookahead)
	cmp55 = v31 == 46
	if cmp55 {
		goto if_then57
	} else {
		goto if_end58
	}

if_then57:
	*libc.As[int16](state_addr) = 98
	goto next_state

if_end58:
	v32 = *libc.As[int32](lookahead)
	cmp59 = 9 <= v32
	if cmp59 {
		goto land_lhs_true61
	} else {
		goto lor_lhs_false64
	}

land_lhs_true61:
	v33 = *libc.As[int32](lookahead)
	cmp62 = v33 <= 13
	if cmp62 {
		goto if_then67
	} else {
		goto lor_lhs_false64
	}

lor_lhs_false64:
	v34 = *libc.As[int32](lookahead)
	cmp65 = v34 == 32
	if cmp65 {
		goto if_then67
	} else {
		goto if_end68
	}

if_then67:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end68:
	v35 = *libc.As[int32](lookahead)
	cmp69 = 48 <= v35
	if cmp69 {
		goto land_lhs_true71
	} else {
		goto if_end75
	}

land_lhs_true71:
	v36 = *libc.As[int32](lookahead)
	cmp72 = v36 <= 57
	if cmp72 {
		goto if_then74
	} else {
		goto if_end75
	}

if_then74:
	*libc.As[int16](state_addr) = 100
	goto next_state

if_end75:
	v37 = *libc.As[int32](lookahead)
	cmp76 = 47 <= v37
	if cmp76 {
		goto land_lhs_true78
	} else {
		goto lor_lhs_false81
	}

land_lhs_true78:
	v38 = *libc.As[int32](lookahead)
	cmp79 = v38 <= 58
	if cmp79 {
		goto if_then96
	} else {
		goto lor_lhs_false81
	}

lor_lhs_false81:
	v39 = *libc.As[int32](lookahead)
	cmp82 = 65 <= v39
	if cmp82 {
		goto land_lhs_true84
	} else {
		goto lor_lhs_false87
	}

land_lhs_true84:
	v40 = *libc.As[int32](lookahead)
	cmp85 = v40 <= 90
	if cmp85 {
		goto if_then96
	} else {
		goto lor_lhs_false87
	}

lor_lhs_false87:
	v41 = *libc.As[int32](lookahead)
	cmp88 = v41 == 95
	if cmp88 {
		goto if_then96
	} else {
		goto lor_lhs_false90
	}

lor_lhs_false90:
	v42 = *libc.As[int32](lookahead)
	cmp91 = 97 <= v42
	if cmp91 {
		goto land_lhs_true93
	} else {
		goto if_end97
	}

land_lhs_true93:
	v43 = *libc.As[int32](lookahead)
	cmp94 = v43 <= 122
	if cmp94 {
		goto if_then96
	} else {
		goto if_end97
	}

if_then96:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end97:
	v44 = *libc.As[byte](result)
	loadedv98 = (v44 & 1) != 0
	*libc.As[bool](retval) = loadedv98
	goto _return

sw_bb99:
	v45 = *libc.As[int32](lookahead)
	cmp100 = 48 <= v45
	if cmp100 {
		goto land_lhs_true102
	} else {
		goto if_end106
	}

land_lhs_true102:
	v46 = *libc.As[int32](lookahead)
	cmp103 = v46 <= 57
	if cmp103 {
		goto if_then105
	} else {
		goto if_end106
	}

if_then105:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end106:
	v47 = *libc.As[byte](result)
	loadedv107 = (v47 & 1) != 0
	*libc.As[bool](retval) = loadedv107
	goto _return

sw_bb108:
	v48 = *libc.As[int32](lookahead)
	cmp109 = v48 != 0
	if cmp109 {
		goto land_lhs_true111
	} else {
		goto if_end115
	}

land_lhs_true111:
	v49 = *libc.As[int32](lookahead)
	cmp112 = v49 != 10
	if cmp112 {
		goto if_then114
	} else {
		goto if_end115
	}

if_then114:
	*libc.As[int16](state_addr) = 102
	goto next_state

if_end115:
	v50 = *libc.As[byte](result)
	loadedv116 = (v50 & 1) != 0
	*libc.As[bool](retval) = loadedv116
	goto _return

sw_bb117:
	*libc.As[byte](result) = 1
	v51 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v51).F1)
	*libc.As[int16](result_symbol) = 0
	v52 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v52).F3)
	v53 = *libc.As[unsafe.Pointer](mark_end)
	v54 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v53)(v54)
	v55 = *libc.As[byte](result)
	loadedv118 = (v55 & 1) != 0
	*libc.As[bool](retval) = loadedv118
	goto _return

sw_bb119:
	*libc.As[byte](result) = 1
	v56 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol120 = libc.Ptr(&libc.As[TSLexer](v56).F1)
	*libc.As[int16](result_symbol120) = 1
	v57 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end121 = libc.Ptr(&libc.As[TSLexer](v57).F3)
	v58 = *libc.As[unsafe.Pointer](mark_end121)
	v59 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v58)(v59)
	v60 = *libc.As[int32](lookahead)
	cmp122 = v60 == 10
	if cmp122 {
		goto if_then124
	} else {
		goto if_end125
	}

if_then124:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end125:
	v61 = *libc.As[byte](result)
	loadedv126 = (v61 & 1) != 0
	*libc.As[bool](retval) = loadedv126
	goto _return

sw_bb127:
	*libc.As[byte](result) = 1
	v62 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol128 = libc.Ptr(&libc.As[TSLexer](v62).F1)
	*libc.As[int16](result_symbol128) = 2
	v63 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end129 = libc.Ptr(&libc.As[TSLexer](v63).F3)
	v64 = *libc.As[unsafe.Pointer](mark_end129)
	v65 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v64)(v65)
	v66 = *libc.As[int32](lookahead)
	cmp130 = 48 <= v66
	if cmp130 {
		goto land_lhs_true132
	} else {
		goto lor_lhs_false135
	}

land_lhs_true132:
	v67 = *libc.As[int32](lookahead)
	cmp133 = v67 <= 57
	if cmp133 {
		goto if_then150
	} else {
		goto lor_lhs_false135
	}

lor_lhs_false135:
	v68 = *libc.As[int32](lookahead)
	cmp136 = 65 <= v68
	if cmp136 {
		goto land_lhs_true138
	} else {
		goto lor_lhs_false141
	}

land_lhs_true138:
	v69 = *libc.As[int32](lookahead)
	cmp139 = v69 <= 90
	if cmp139 {
		goto if_then150
	} else {
		goto lor_lhs_false141
	}

lor_lhs_false141:
	v70 = *libc.As[int32](lookahead)
	cmp142 = v70 == 95
	if cmp142 {
		goto if_then150
	} else {
		goto lor_lhs_false144
	}

lor_lhs_false144:
	v71 = *libc.As[int32](lookahead)
	cmp145 = 97 <= v71
	if cmp145 {
		goto land_lhs_true147
	} else {
		goto if_end151
	}

land_lhs_true147:
	v72 = *libc.As[int32](lookahead)
	cmp148 = v72 <= 122
	if cmp148 {
		goto if_then150
	} else {
		goto if_end151
	}

if_then150:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end151:
	v73 = *libc.As[byte](result)
	loadedv152 = (v73 & 1) != 0
	*libc.As[bool](retval) = loadedv152
	goto _return

sw_bb153:
	*libc.As[byte](result) = 1
	v74 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol154 = libc.Ptr(&libc.As[TSLexer](v74).F1)
	*libc.As[int16](result_symbol154) = 3
	v75 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end155 = libc.Ptr(&libc.As[TSLexer](v75).F3)
	v76 = *libc.As[unsafe.Pointer](mark_end155)
	v77 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v76)(v77)
	v78 = *libc.As[int32](lookahead)
	cmp156 = 48 <= v78
	if cmp156 {
		goto land_lhs_true158
	} else {
		goto lor_lhs_false161
	}

land_lhs_true158:
	v79 = *libc.As[int32](lookahead)
	cmp159 = v79 <= 57
	if cmp159 {
		goto if_then176
	} else {
		goto lor_lhs_false161
	}

lor_lhs_false161:
	v80 = *libc.As[int32](lookahead)
	cmp162 = 65 <= v80
	if cmp162 {
		goto land_lhs_true164
	} else {
		goto lor_lhs_false167
	}

land_lhs_true164:
	v81 = *libc.As[int32](lookahead)
	cmp165 = v81 <= 90
	if cmp165 {
		goto if_then176
	} else {
		goto lor_lhs_false167
	}

lor_lhs_false167:
	v82 = *libc.As[int32](lookahead)
	cmp168 = v82 == 95
	if cmp168 {
		goto if_then176
	} else {
		goto lor_lhs_false170
	}

lor_lhs_false170:
	v83 = *libc.As[int32](lookahead)
	cmp171 = 97 <= v83
	if cmp171 {
		goto land_lhs_true173
	} else {
		goto if_end177
	}

land_lhs_true173:
	v84 = *libc.As[int32](lookahead)
	cmp174 = v84 <= 122
	if cmp174 {
		goto if_then176
	} else {
		goto if_end177
	}

if_then176:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end177:
	v85 = *libc.As[byte](result)
	loadedv178 = (v85 & 1) != 0
	*libc.As[bool](retval) = loadedv178
	goto _return

sw_bb179:
	*libc.As[byte](result) = 1
	v86 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol180 = libc.Ptr(&libc.As[TSLexer](v86).F1)
	*libc.As[int16](result_symbol180) = 4
	v87 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end181 = libc.Ptr(&libc.As[TSLexer](v87).F3)
	v88 = *libc.As[unsafe.Pointer](mark_end181)
	v89 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v88)(v89)
	v90 = *libc.As[int32](lookahead)
	cmp182 = 48 <= v90
	if cmp182 {
		goto land_lhs_true184
	} else {
		goto lor_lhs_false187
	}

land_lhs_true184:
	v91 = *libc.As[int32](lookahead)
	cmp185 = v91 <= 57
	if cmp185 {
		goto if_then202
	} else {
		goto lor_lhs_false187
	}

lor_lhs_false187:
	v92 = *libc.As[int32](lookahead)
	cmp188 = 65 <= v92
	if cmp188 {
		goto land_lhs_true190
	} else {
		goto lor_lhs_false193
	}

land_lhs_true190:
	v93 = *libc.As[int32](lookahead)
	cmp191 = v93 <= 90
	if cmp191 {
		goto if_then202
	} else {
		goto lor_lhs_false193
	}

lor_lhs_false193:
	v94 = *libc.As[int32](lookahead)
	cmp194 = v94 == 95
	if cmp194 {
		goto if_then202
	} else {
		goto lor_lhs_false196
	}

lor_lhs_false196:
	v95 = *libc.As[int32](lookahead)
	cmp197 = 97 <= v95
	if cmp197 {
		goto land_lhs_true199
	} else {
		goto if_end203
	}

land_lhs_true199:
	v96 = *libc.As[int32](lookahead)
	cmp200 = v96 <= 122
	if cmp200 {
		goto if_then202
	} else {
		goto if_end203
	}

if_then202:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end203:
	v97 = *libc.As[byte](result)
	loadedv204 = (v97 & 1) != 0
	*libc.As[bool](retval) = loadedv204
	goto _return

sw_bb205:
	*libc.As[byte](result) = 1
	v98 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol206 = libc.Ptr(&libc.As[TSLexer](v98).F1)
	*libc.As[int16](result_symbol206) = 5
	v99 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end207 = libc.Ptr(&libc.As[TSLexer](v99).F3)
	v100 = *libc.As[unsafe.Pointer](mark_end207)
	v101 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v100)(v101)
	v102 = *libc.As[int32](lookahead)
	cmp208 = 48 <= v102
	if cmp208 {
		goto land_lhs_true210
	} else {
		goto lor_lhs_false213
	}

land_lhs_true210:
	v103 = *libc.As[int32](lookahead)
	cmp211 = v103 <= 57
	if cmp211 {
		goto if_then228
	} else {
		goto lor_lhs_false213
	}

lor_lhs_false213:
	v104 = *libc.As[int32](lookahead)
	cmp214 = 65 <= v104
	if cmp214 {
		goto land_lhs_true216
	} else {
		goto lor_lhs_false219
	}

land_lhs_true216:
	v105 = *libc.As[int32](lookahead)
	cmp217 = v105 <= 90
	if cmp217 {
		goto if_then228
	} else {
		goto lor_lhs_false219
	}

lor_lhs_false219:
	v106 = *libc.As[int32](lookahead)
	cmp220 = v106 == 95
	if cmp220 {
		goto if_then228
	} else {
		goto lor_lhs_false222
	}

lor_lhs_false222:
	v107 = *libc.As[int32](lookahead)
	cmp223 = 97 <= v107
	if cmp223 {
		goto land_lhs_true225
	} else {
		goto if_end229
	}

land_lhs_true225:
	v108 = *libc.As[int32](lookahead)
	cmp226 = v108 <= 122
	if cmp226 {
		goto if_then228
	} else {
		goto if_end229
	}

if_then228:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end229:
	v109 = *libc.As[byte](result)
	loadedv230 = (v109 & 1) != 0
	*libc.As[bool](retval) = loadedv230
	goto _return

sw_bb231:
	*libc.As[byte](result) = 1
	v110 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol232 = libc.Ptr(&libc.As[TSLexer](v110).F1)
	*libc.As[int16](result_symbol232) = 6
	v111 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end233 = libc.Ptr(&libc.As[TSLexer](v111).F3)
	v112 = *libc.As[unsafe.Pointer](mark_end233)
	v113 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v112)(v113)
	v114 = *libc.As[int32](lookahead)
	cmp234 = 48 <= v114
	if cmp234 {
		goto land_lhs_true236
	} else {
		goto lor_lhs_false239
	}

land_lhs_true236:
	v115 = *libc.As[int32](lookahead)
	cmp237 = v115 <= 57
	if cmp237 {
		goto if_then254
	} else {
		goto lor_lhs_false239
	}

lor_lhs_false239:
	v116 = *libc.As[int32](lookahead)
	cmp240 = 65 <= v116
	if cmp240 {
		goto land_lhs_true242
	} else {
		goto lor_lhs_false245
	}

land_lhs_true242:
	v117 = *libc.As[int32](lookahead)
	cmp243 = v117 <= 90
	if cmp243 {
		goto if_then254
	} else {
		goto lor_lhs_false245
	}

lor_lhs_false245:
	v118 = *libc.As[int32](lookahead)
	cmp246 = v118 == 95
	if cmp246 {
		goto if_then254
	} else {
		goto lor_lhs_false248
	}

lor_lhs_false248:
	v119 = *libc.As[int32](lookahead)
	cmp249 = 97 <= v119
	if cmp249 {
		goto land_lhs_true251
	} else {
		goto if_end255
	}

land_lhs_true251:
	v120 = *libc.As[int32](lookahead)
	cmp252 = v120 <= 122
	if cmp252 {
		goto if_then254
	} else {
		goto if_end255
	}

if_then254:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end255:
	v121 = *libc.As[byte](result)
	loadedv256 = (v121 & 1) != 0
	*libc.As[bool](retval) = loadedv256
	goto _return

sw_bb257:
	*libc.As[byte](result) = 1
	v122 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol258 = libc.Ptr(&libc.As[TSLexer](v122).F1)
	*libc.As[int16](result_symbol258) = 7
	v123 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end259 = libc.Ptr(&libc.As[TSLexer](v123).F3)
	v124 = *libc.As[unsafe.Pointer](mark_end259)
	v125 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v124)(v125)
	v126 = *libc.As[int32](lookahead)
	cmp260 = 48 <= v126
	if cmp260 {
		goto land_lhs_true262
	} else {
		goto lor_lhs_false265
	}

land_lhs_true262:
	v127 = *libc.As[int32](lookahead)
	cmp263 = v127 <= 57
	if cmp263 {
		goto if_then280
	} else {
		goto lor_lhs_false265
	}

lor_lhs_false265:
	v128 = *libc.As[int32](lookahead)
	cmp266 = 65 <= v128
	if cmp266 {
		goto land_lhs_true268
	} else {
		goto lor_lhs_false271
	}

land_lhs_true268:
	v129 = *libc.As[int32](lookahead)
	cmp269 = v129 <= 90
	if cmp269 {
		goto if_then280
	} else {
		goto lor_lhs_false271
	}

lor_lhs_false271:
	v130 = *libc.As[int32](lookahead)
	cmp272 = v130 == 95
	if cmp272 {
		goto if_then280
	} else {
		goto lor_lhs_false274
	}

lor_lhs_false274:
	v131 = *libc.As[int32](lookahead)
	cmp275 = 97 <= v131
	if cmp275 {
		goto land_lhs_true277
	} else {
		goto if_end281
	}

land_lhs_true277:
	v132 = *libc.As[int32](lookahead)
	cmp278 = v132 <= 122
	if cmp278 {
		goto if_then280
	} else {
		goto if_end281
	}

if_then280:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end281:
	v133 = *libc.As[byte](result)
	loadedv282 = (v133 & 1) != 0
	*libc.As[bool](retval) = loadedv282
	goto _return

sw_bb283:
	*libc.As[byte](result) = 1
	v134 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol284 = libc.Ptr(&libc.As[TSLexer](v134).F1)
	*libc.As[int16](result_symbol284) = 8
	v135 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end285 = libc.Ptr(&libc.As[TSLexer](v135).F3)
	v136 = *libc.As[unsafe.Pointer](mark_end285)
	v137 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v136)(v137)
	v138 = *libc.As[int32](lookahead)
	cmp286 = 48 <= v138
	if cmp286 {
		goto land_lhs_true288
	} else {
		goto lor_lhs_false291
	}

land_lhs_true288:
	v139 = *libc.As[int32](lookahead)
	cmp289 = v139 <= 57
	if cmp289 {
		goto if_then306
	} else {
		goto lor_lhs_false291
	}

lor_lhs_false291:
	v140 = *libc.As[int32](lookahead)
	cmp292 = 65 <= v140
	if cmp292 {
		goto land_lhs_true294
	} else {
		goto lor_lhs_false297
	}

land_lhs_true294:
	v141 = *libc.As[int32](lookahead)
	cmp295 = v141 <= 90
	if cmp295 {
		goto if_then306
	} else {
		goto lor_lhs_false297
	}

lor_lhs_false297:
	v142 = *libc.As[int32](lookahead)
	cmp298 = v142 == 95
	if cmp298 {
		goto if_then306
	} else {
		goto lor_lhs_false300
	}

lor_lhs_false300:
	v143 = *libc.As[int32](lookahead)
	cmp301 = 97 <= v143
	if cmp301 {
		goto land_lhs_true303
	} else {
		goto if_end307
	}

land_lhs_true303:
	v144 = *libc.As[int32](lookahead)
	cmp304 = v144 <= 122
	if cmp304 {
		goto if_then306
	} else {
		goto if_end307
	}

if_then306:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end307:
	v145 = *libc.As[byte](result)
	loadedv308 = (v145 & 1) != 0
	*libc.As[bool](retval) = loadedv308
	goto _return

sw_bb309:
	*libc.As[byte](result) = 1
	v146 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol310 = libc.Ptr(&libc.As[TSLexer](v146).F1)
	*libc.As[int16](result_symbol310) = 9
	v147 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end311 = libc.Ptr(&libc.As[TSLexer](v147).F3)
	v148 = *libc.As[unsafe.Pointer](mark_end311)
	v149 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v148)(v149)
	v150 = *libc.As[int32](lookahead)
	cmp312 = 48 <= v150
	if cmp312 {
		goto land_lhs_true314
	} else {
		goto lor_lhs_false317
	}

land_lhs_true314:
	v151 = *libc.As[int32](lookahead)
	cmp315 = v151 <= 57
	if cmp315 {
		goto if_then332
	} else {
		goto lor_lhs_false317
	}

lor_lhs_false317:
	v152 = *libc.As[int32](lookahead)
	cmp318 = 65 <= v152
	if cmp318 {
		goto land_lhs_true320
	} else {
		goto lor_lhs_false323
	}

land_lhs_true320:
	v153 = *libc.As[int32](lookahead)
	cmp321 = v153 <= 90
	if cmp321 {
		goto if_then332
	} else {
		goto lor_lhs_false323
	}

lor_lhs_false323:
	v154 = *libc.As[int32](lookahead)
	cmp324 = v154 == 95
	if cmp324 {
		goto if_then332
	} else {
		goto lor_lhs_false326
	}

lor_lhs_false326:
	v155 = *libc.As[int32](lookahead)
	cmp327 = 97 <= v155
	if cmp327 {
		goto land_lhs_true329
	} else {
		goto if_end333
	}

land_lhs_true329:
	v156 = *libc.As[int32](lookahead)
	cmp330 = v156 <= 122
	if cmp330 {
		goto if_then332
	} else {
		goto if_end333
	}

if_then332:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end333:
	v157 = *libc.As[byte](result)
	loadedv334 = (v157 & 1) != 0
	*libc.As[bool](retval) = loadedv334
	goto _return

sw_bb335:
	*libc.As[byte](result) = 1
	v158 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol336 = libc.Ptr(&libc.As[TSLexer](v158).F1)
	*libc.As[int16](result_symbol336) = 10
	v159 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end337 = libc.Ptr(&libc.As[TSLexer](v159).F3)
	v160 = *libc.As[unsafe.Pointer](mark_end337)
	v161 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v160)(v161)
	v162 = *libc.As[int32](lookahead)
	cmp338 = 48 <= v162
	if cmp338 {
		goto land_lhs_true340
	} else {
		goto lor_lhs_false343
	}

land_lhs_true340:
	v163 = *libc.As[int32](lookahead)
	cmp341 = v163 <= 57
	if cmp341 {
		goto if_then358
	} else {
		goto lor_lhs_false343
	}

lor_lhs_false343:
	v164 = *libc.As[int32](lookahead)
	cmp344 = 65 <= v164
	if cmp344 {
		goto land_lhs_true346
	} else {
		goto lor_lhs_false349
	}

land_lhs_true346:
	v165 = *libc.As[int32](lookahead)
	cmp347 = v165 <= 90
	if cmp347 {
		goto if_then358
	} else {
		goto lor_lhs_false349
	}

lor_lhs_false349:
	v166 = *libc.As[int32](lookahead)
	cmp350 = v166 == 95
	if cmp350 {
		goto if_then358
	} else {
		goto lor_lhs_false352
	}

lor_lhs_false352:
	v167 = *libc.As[int32](lookahead)
	cmp353 = 97 <= v167
	if cmp353 {
		goto land_lhs_true355
	} else {
		goto if_end359
	}

land_lhs_true355:
	v168 = *libc.As[int32](lookahead)
	cmp356 = v168 <= 122
	if cmp356 {
		goto if_then358
	} else {
		goto if_end359
	}

if_then358:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end359:
	v169 = *libc.As[byte](result)
	loadedv360 = (v169 & 1) != 0
	*libc.As[bool](retval) = loadedv360
	goto _return

sw_bb361:
	*libc.As[byte](result) = 1
	v170 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol362 = libc.Ptr(&libc.As[TSLexer](v170).F1)
	*libc.As[int16](result_symbol362) = 11
	v171 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end363 = libc.Ptr(&libc.As[TSLexer](v171).F3)
	v172 = *libc.As[unsafe.Pointer](mark_end363)
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v172)(v173)
	v174 = *libc.As[int32](lookahead)
	cmp364 = 48 <= v174
	if cmp364 {
		goto land_lhs_true366
	} else {
		goto lor_lhs_false369
	}

land_lhs_true366:
	v175 = *libc.As[int32](lookahead)
	cmp367 = v175 <= 57
	if cmp367 {
		goto if_then384
	} else {
		goto lor_lhs_false369
	}

lor_lhs_false369:
	v176 = *libc.As[int32](lookahead)
	cmp370 = 65 <= v176
	if cmp370 {
		goto land_lhs_true372
	} else {
		goto lor_lhs_false375
	}

land_lhs_true372:
	v177 = *libc.As[int32](lookahead)
	cmp373 = v177 <= 90
	if cmp373 {
		goto if_then384
	} else {
		goto lor_lhs_false375
	}

lor_lhs_false375:
	v178 = *libc.As[int32](lookahead)
	cmp376 = v178 == 95
	if cmp376 {
		goto if_then384
	} else {
		goto lor_lhs_false378
	}

lor_lhs_false378:
	v179 = *libc.As[int32](lookahead)
	cmp379 = 97 <= v179
	if cmp379 {
		goto land_lhs_true381
	} else {
		goto if_end385
	}

land_lhs_true381:
	v180 = *libc.As[int32](lookahead)
	cmp382 = v180 <= 122
	if cmp382 {
		goto if_then384
	} else {
		goto if_end385
	}

if_then384:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end385:
	v181 = *libc.As[byte](result)
	loadedv386 = (v181 & 1) != 0
	*libc.As[bool](retval) = loadedv386
	goto _return

sw_bb387:
	*libc.As[byte](result) = 1
	v182 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol388 = libc.Ptr(&libc.As[TSLexer](v182).F1)
	*libc.As[int16](result_symbol388) = 12
	v183 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end389 = libc.Ptr(&libc.As[TSLexer](v183).F3)
	v184 = *libc.As[unsafe.Pointer](mark_end389)
	v185 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v184)(v185)
	v186 = *libc.As[int32](lookahead)
	cmp390 = 48 <= v186
	if cmp390 {
		goto land_lhs_true392
	} else {
		goto lor_lhs_false395
	}

land_lhs_true392:
	v187 = *libc.As[int32](lookahead)
	cmp393 = v187 <= 57
	if cmp393 {
		goto if_then410
	} else {
		goto lor_lhs_false395
	}

lor_lhs_false395:
	v188 = *libc.As[int32](lookahead)
	cmp396 = 65 <= v188
	if cmp396 {
		goto land_lhs_true398
	} else {
		goto lor_lhs_false401
	}

land_lhs_true398:
	v189 = *libc.As[int32](lookahead)
	cmp399 = v189 <= 90
	if cmp399 {
		goto if_then410
	} else {
		goto lor_lhs_false401
	}

lor_lhs_false401:
	v190 = *libc.As[int32](lookahead)
	cmp402 = v190 == 95
	if cmp402 {
		goto if_then410
	} else {
		goto lor_lhs_false404
	}

lor_lhs_false404:
	v191 = *libc.As[int32](lookahead)
	cmp405 = 97 <= v191
	if cmp405 {
		goto land_lhs_true407
	} else {
		goto if_end411
	}

land_lhs_true407:
	v192 = *libc.As[int32](lookahead)
	cmp408 = v192 <= 122
	if cmp408 {
		goto if_then410
	} else {
		goto if_end411
	}

if_then410:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end411:
	v193 = *libc.As[byte](result)
	loadedv412 = (v193 & 1) != 0
	*libc.As[bool](retval) = loadedv412
	goto _return

sw_bb413:
	*libc.As[byte](result) = 1
	v194 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol414 = libc.Ptr(&libc.As[TSLexer](v194).F1)
	*libc.As[int16](result_symbol414) = 13
	v195 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end415 = libc.Ptr(&libc.As[TSLexer](v195).F3)
	v196 = *libc.As[unsafe.Pointer](mark_end415)
	v197 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v196)(v197)
	v198 = *libc.As[int32](lookahead)
	cmp416 = v198 == 97
	if cmp416 {
		goto if_then418
	} else {
		goto if_end419
	}

if_then418:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end419:
	v199 = *libc.As[int32](lookahead)
	cmp420 = 48 <= v199
	if cmp420 {
		goto land_lhs_true422
	} else {
		goto lor_lhs_false425
	}

land_lhs_true422:
	v200 = *libc.As[int32](lookahead)
	cmp423 = v200 <= 57
	if cmp423 {
		goto if_then440
	} else {
		goto lor_lhs_false425
	}

lor_lhs_false425:
	v201 = *libc.As[int32](lookahead)
	cmp426 = 65 <= v201
	if cmp426 {
		goto land_lhs_true428
	} else {
		goto lor_lhs_false431
	}

land_lhs_true428:
	v202 = *libc.As[int32](lookahead)
	cmp429 = v202 <= 90
	if cmp429 {
		goto if_then440
	} else {
		goto lor_lhs_false431
	}

lor_lhs_false431:
	v203 = *libc.As[int32](lookahead)
	cmp432 = v203 == 95
	if cmp432 {
		goto if_then440
	} else {
		goto lor_lhs_false434
	}

lor_lhs_false434:
	v204 = *libc.As[int32](lookahead)
	cmp435 = 98 <= v204
	if cmp435 {
		goto land_lhs_true437
	} else {
		goto if_end441
	}

land_lhs_true437:
	v205 = *libc.As[int32](lookahead)
	cmp438 = v205 <= 122
	if cmp438 {
		goto if_then440
	} else {
		goto if_end441
	}

if_then440:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end441:
	v206 = *libc.As[byte](result)
	loadedv442 = (v206 & 1) != 0
	*libc.As[bool](retval) = loadedv442
	goto _return

sw_bb443:
	*libc.As[byte](result) = 1
	v207 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol444 = libc.Ptr(&libc.As[TSLexer](v207).F1)
	*libc.As[int16](result_symbol444) = 13
	v208 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end445 = libc.Ptr(&libc.As[TSLexer](v208).F3)
	v209 = *libc.As[unsafe.Pointer](mark_end445)
	v210 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v209)(v210)
	v211 = *libc.As[int32](lookahead)
	cmp446 = v211 == 97
	if cmp446 {
		goto if_then448
	} else {
		goto if_end449
	}

if_then448:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end449:
	v212 = *libc.As[int32](lookahead)
	cmp450 = 48 <= v212
	if cmp450 {
		goto land_lhs_true452
	} else {
		goto lor_lhs_false455
	}

land_lhs_true452:
	v213 = *libc.As[int32](lookahead)
	cmp453 = v213 <= 57
	if cmp453 {
		goto if_then470
	} else {
		goto lor_lhs_false455
	}

lor_lhs_false455:
	v214 = *libc.As[int32](lookahead)
	cmp456 = 65 <= v214
	if cmp456 {
		goto land_lhs_true458
	} else {
		goto lor_lhs_false461
	}

land_lhs_true458:
	v215 = *libc.As[int32](lookahead)
	cmp459 = v215 <= 90
	if cmp459 {
		goto if_then470
	} else {
		goto lor_lhs_false461
	}

lor_lhs_false461:
	v216 = *libc.As[int32](lookahead)
	cmp462 = v216 == 95
	if cmp462 {
		goto if_then470
	} else {
		goto lor_lhs_false464
	}

lor_lhs_false464:
	v217 = *libc.As[int32](lookahead)
	cmp465 = 98 <= v217
	if cmp465 {
		goto land_lhs_true467
	} else {
		goto if_end471
	}

land_lhs_true467:
	v218 = *libc.As[int32](lookahead)
	cmp468 = v218 <= 122
	if cmp468 {
		goto if_then470
	} else {
		goto if_end471
	}

if_then470:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end471:
	v219 = *libc.As[byte](result)
	loadedv472 = (v219 & 1) != 0
	*libc.As[bool](retval) = loadedv472
	goto _return

sw_bb473:
	*libc.As[byte](result) = 1
	v220 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol474 = libc.Ptr(&libc.As[TSLexer](v220).F1)
	*libc.As[int16](result_symbol474) = 13
	v221 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end475 = libc.Ptr(&libc.As[TSLexer](v221).F3)
	v222 = *libc.As[unsafe.Pointer](mark_end475)
	v223 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v222)(v223)
	v224 = *libc.As[int32](lookahead)
	cmp476 = v224 == 97
	if cmp476 {
		goto if_then478
	} else {
		goto if_end479
	}

if_then478:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end479:
	v225 = *libc.As[int32](lookahead)
	cmp480 = 48 <= v225
	if cmp480 {
		goto land_lhs_true482
	} else {
		goto lor_lhs_false485
	}

land_lhs_true482:
	v226 = *libc.As[int32](lookahead)
	cmp483 = v226 <= 57
	if cmp483 {
		goto if_then500
	} else {
		goto lor_lhs_false485
	}

lor_lhs_false485:
	v227 = *libc.As[int32](lookahead)
	cmp486 = 65 <= v227
	if cmp486 {
		goto land_lhs_true488
	} else {
		goto lor_lhs_false491
	}

land_lhs_true488:
	v228 = *libc.As[int32](lookahead)
	cmp489 = v228 <= 90
	if cmp489 {
		goto if_then500
	} else {
		goto lor_lhs_false491
	}

lor_lhs_false491:
	v229 = *libc.As[int32](lookahead)
	cmp492 = v229 == 95
	if cmp492 {
		goto if_then500
	} else {
		goto lor_lhs_false494
	}

lor_lhs_false494:
	v230 = *libc.As[int32](lookahead)
	cmp495 = 98 <= v230
	if cmp495 {
		goto land_lhs_true497
	} else {
		goto if_end501
	}

land_lhs_true497:
	v231 = *libc.As[int32](lookahead)
	cmp498 = v231 <= 122
	if cmp498 {
		goto if_then500
	} else {
		goto if_end501
	}

if_then500:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end501:
	v232 = *libc.As[byte](result)
	loadedv502 = (v232 & 1) != 0
	*libc.As[bool](retval) = loadedv502
	goto _return

sw_bb503:
	*libc.As[byte](result) = 1
	v233 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol504 = libc.Ptr(&libc.As[TSLexer](v233).F1)
	*libc.As[int16](result_symbol504) = 13
	v234 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end505 = libc.Ptr(&libc.As[TSLexer](v234).F3)
	v235 = *libc.As[unsafe.Pointer](mark_end505)
	v236 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v235)(v236)
	v237 = *libc.As[int32](lookahead)
	cmp506 = v237 == 97
	if cmp506 {
		goto if_then508
	} else {
		goto if_end509
	}

if_then508:
	*libc.As[int16](state_addr) = 80
	goto next_state

if_end509:
	v238 = *libc.As[int32](lookahead)
	cmp510 = 48 <= v238
	if cmp510 {
		goto land_lhs_true512
	} else {
		goto lor_lhs_false515
	}

land_lhs_true512:
	v239 = *libc.As[int32](lookahead)
	cmp513 = v239 <= 57
	if cmp513 {
		goto if_then530
	} else {
		goto lor_lhs_false515
	}

lor_lhs_false515:
	v240 = *libc.As[int32](lookahead)
	cmp516 = 65 <= v240
	if cmp516 {
		goto land_lhs_true518
	} else {
		goto lor_lhs_false521
	}

land_lhs_true518:
	v241 = *libc.As[int32](lookahead)
	cmp519 = v241 <= 90
	if cmp519 {
		goto if_then530
	} else {
		goto lor_lhs_false521
	}

lor_lhs_false521:
	v242 = *libc.As[int32](lookahead)
	cmp522 = v242 == 95
	if cmp522 {
		goto if_then530
	} else {
		goto lor_lhs_false524
	}

lor_lhs_false524:
	v243 = *libc.As[int32](lookahead)
	cmp525 = 98 <= v243
	if cmp525 {
		goto land_lhs_true527
	} else {
		goto if_end531
	}

land_lhs_true527:
	v244 = *libc.As[int32](lookahead)
	cmp528 = v244 <= 122
	if cmp528 {
		goto if_then530
	} else {
		goto if_end531
	}

if_then530:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end531:
	v245 = *libc.As[byte](result)
	loadedv532 = (v245 & 1) != 0
	*libc.As[bool](retval) = loadedv532
	goto _return

sw_bb533:
	*libc.As[byte](result) = 1
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol534 = libc.Ptr(&libc.As[TSLexer](v246).F1)
	*libc.As[int16](result_symbol534) = 13
	v247 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end535 = libc.Ptr(&libc.As[TSLexer](v247).F3)
	v248 = *libc.As[unsafe.Pointer](mark_end535)
	v249 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v248)(v249)
	v250 = *libc.As[int32](lookahead)
	cmp536 = v250 == 97
	if cmp536 {
		goto if_then538
	} else {
		goto if_end539
	}

if_then538:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end539:
	v251 = *libc.As[int32](lookahead)
	cmp540 = 48 <= v251
	if cmp540 {
		goto land_lhs_true542
	} else {
		goto lor_lhs_false545
	}

land_lhs_true542:
	v252 = *libc.As[int32](lookahead)
	cmp543 = v252 <= 57
	if cmp543 {
		goto if_then560
	} else {
		goto lor_lhs_false545
	}

lor_lhs_false545:
	v253 = *libc.As[int32](lookahead)
	cmp546 = 65 <= v253
	if cmp546 {
		goto land_lhs_true548
	} else {
		goto lor_lhs_false551
	}

land_lhs_true548:
	v254 = *libc.As[int32](lookahead)
	cmp549 = v254 <= 90
	if cmp549 {
		goto if_then560
	} else {
		goto lor_lhs_false551
	}

lor_lhs_false551:
	v255 = *libc.As[int32](lookahead)
	cmp552 = v255 == 95
	if cmp552 {
		goto if_then560
	} else {
		goto lor_lhs_false554
	}

lor_lhs_false554:
	v256 = *libc.As[int32](lookahead)
	cmp555 = 98 <= v256
	if cmp555 {
		goto land_lhs_true557
	} else {
		goto if_end561
	}

land_lhs_true557:
	v257 = *libc.As[int32](lookahead)
	cmp558 = v257 <= 122
	if cmp558 {
		goto if_then560
	} else {
		goto if_end561
	}

if_then560:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end561:
	v258 = *libc.As[byte](result)
	loadedv562 = (v258 & 1) != 0
	*libc.As[bool](retval) = loadedv562
	goto _return

sw_bb563:
	*libc.As[byte](result) = 1
	v259 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol564 = libc.Ptr(&libc.As[TSLexer](v259).F1)
	*libc.As[int16](result_symbol564) = 13
	v260 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end565 = libc.Ptr(&libc.As[TSLexer](v260).F3)
	v261 = *libc.As[unsafe.Pointer](mark_end565)
	v262 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v261)(v262)
	v263 = *libc.As[int32](lookahead)
	cmp566 = v263 == 100
	if cmp566 {
		goto if_then568
	} else {
		goto if_end569
	}

if_then568:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end569:
	v264 = *libc.As[int32](lookahead)
	cmp570 = 48 <= v264
	if cmp570 {
		goto land_lhs_true572
	} else {
		goto lor_lhs_false575
	}

land_lhs_true572:
	v265 = *libc.As[int32](lookahead)
	cmp573 = v265 <= 57
	if cmp573 {
		goto if_then590
	} else {
		goto lor_lhs_false575
	}

lor_lhs_false575:
	v266 = *libc.As[int32](lookahead)
	cmp576 = 65 <= v266
	if cmp576 {
		goto land_lhs_true578
	} else {
		goto lor_lhs_false581
	}

land_lhs_true578:
	v267 = *libc.As[int32](lookahead)
	cmp579 = v267 <= 90
	if cmp579 {
		goto if_then590
	} else {
		goto lor_lhs_false581
	}

lor_lhs_false581:
	v268 = *libc.As[int32](lookahead)
	cmp582 = v268 == 95
	if cmp582 {
		goto if_then590
	} else {
		goto lor_lhs_false584
	}

lor_lhs_false584:
	v269 = *libc.As[int32](lookahead)
	cmp585 = 97 <= v269
	if cmp585 {
		goto land_lhs_true587
	} else {
		goto if_end591
	}

land_lhs_true587:
	v270 = *libc.As[int32](lookahead)
	cmp588 = v270 <= 122
	if cmp588 {
		goto if_then590
	} else {
		goto if_end591
	}

if_then590:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end591:
	v271 = *libc.As[byte](result)
	loadedv592 = (v271 & 1) != 0
	*libc.As[bool](retval) = loadedv592
	goto _return

sw_bb593:
	*libc.As[byte](result) = 1
	v272 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol594 = libc.Ptr(&libc.As[TSLexer](v272).F1)
	*libc.As[int16](result_symbol594) = 13
	v273 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end595 = libc.Ptr(&libc.As[TSLexer](v273).F3)
	v274 = *libc.As[unsafe.Pointer](mark_end595)
	v275 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v274)(v275)
	v276 = *libc.As[int32](lookahead)
	cmp596 = v276 == 100
	if cmp596 {
		goto if_then598
	} else {
		goto if_end599
	}

if_then598:
	*libc.As[int16](state_addr) = 94
	goto next_state

if_end599:
	v277 = *libc.As[int32](lookahead)
	cmp600 = 48 <= v277
	if cmp600 {
		goto land_lhs_true602
	} else {
		goto lor_lhs_false605
	}

land_lhs_true602:
	v278 = *libc.As[int32](lookahead)
	cmp603 = v278 <= 57
	if cmp603 {
		goto if_then620
	} else {
		goto lor_lhs_false605
	}

lor_lhs_false605:
	v279 = *libc.As[int32](lookahead)
	cmp606 = 65 <= v279
	if cmp606 {
		goto land_lhs_true608
	} else {
		goto lor_lhs_false611
	}

land_lhs_true608:
	v280 = *libc.As[int32](lookahead)
	cmp609 = v280 <= 90
	if cmp609 {
		goto if_then620
	} else {
		goto lor_lhs_false611
	}

lor_lhs_false611:
	v281 = *libc.As[int32](lookahead)
	cmp612 = v281 == 95
	if cmp612 {
		goto if_then620
	} else {
		goto lor_lhs_false614
	}

lor_lhs_false614:
	v282 = *libc.As[int32](lookahead)
	cmp615 = 97 <= v282
	if cmp615 {
		goto land_lhs_true617
	} else {
		goto if_end621
	}

land_lhs_true617:
	v283 = *libc.As[int32](lookahead)
	cmp618 = v283 <= 122
	if cmp618 {
		goto if_then620
	} else {
		goto if_end621
	}

if_then620:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end621:
	v284 = *libc.As[byte](result)
	loadedv622 = (v284 & 1) != 0
	*libc.As[bool](retval) = loadedv622
	goto _return

sw_bb623:
	*libc.As[byte](result) = 1
	v285 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol624 = libc.Ptr(&libc.As[TSLexer](v285).F1)
	*libc.As[int16](result_symbol624) = 13
	v286 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end625 = libc.Ptr(&libc.As[TSLexer](v286).F3)
	v287 = *libc.As[unsafe.Pointer](mark_end625)
	v288 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v287)(v288)
	v289 = *libc.As[int32](lookahead)
	cmp626 = v289 == 100
	if cmp626 {
		goto if_then628
	} else {
		goto if_end629
	}

if_then628:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end629:
	v290 = *libc.As[int32](lookahead)
	cmp630 = 48 <= v290
	if cmp630 {
		goto land_lhs_true632
	} else {
		goto lor_lhs_false635
	}

land_lhs_true632:
	v291 = *libc.As[int32](lookahead)
	cmp633 = v291 <= 57
	if cmp633 {
		goto if_then650
	} else {
		goto lor_lhs_false635
	}

lor_lhs_false635:
	v292 = *libc.As[int32](lookahead)
	cmp636 = 65 <= v292
	if cmp636 {
		goto land_lhs_true638
	} else {
		goto lor_lhs_false641
	}

land_lhs_true638:
	v293 = *libc.As[int32](lookahead)
	cmp639 = v293 <= 90
	if cmp639 {
		goto if_then650
	} else {
		goto lor_lhs_false641
	}

lor_lhs_false641:
	v294 = *libc.As[int32](lookahead)
	cmp642 = v294 == 95
	if cmp642 {
		goto if_then650
	} else {
		goto lor_lhs_false644
	}

lor_lhs_false644:
	v295 = *libc.As[int32](lookahead)
	cmp645 = 97 <= v295
	if cmp645 {
		goto land_lhs_true647
	} else {
		goto if_end651
	}

land_lhs_true647:
	v296 = *libc.As[int32](lookahead)
	cmp648 = v296 <= 122
	if cmp648 {
		goto if_then650
	} else {
		goto if_end651
	}

if_then650:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end651:
	v297 = *libc.As[byte](result)
	loadedv652 = (v297 & 1) != 0
	*libc.As[bool](retval) = loadedv652
	goto _return

sw_bb653:
	*libc.As[byte](result) = 1
	v298 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol654 = libc.Ptr(&libc.As[TSLexer](v298).F1)
	*libc.As[int16](result_symbol654) = 13
	v299 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end655 = libc.Ptr(&libc.As[TSLexer](v299).F3)
	v300 = *libc.As[unsafe.Pointer](mark_end655)
	v301 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v300)(v301)
	v302 = *libc.As[int32](lookahead)
	cmp656 = v302 == 101
	if cmp656 {
		goto if_then658
	} else {
		goto if_end659
	}

if_then658:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end659:
	v303 = *libc.As[int32](lookahead)
	cmp660 = 48 <= v303
	if cmp660 {
		goto land_lhs_true662
	} else {
		goto lor_lhs_false665
	}

land_lhs_true662:
	v304 = *libc.As[int32](lookahead)
	cmp663 = v304 <= 57
	if cmp663 {
		goto if_then680
	} else {
		goto lor_lhs_false665
	}

lor_lhs_false665:
	v305 = *libc.As[int32](lookahead)
	cmp666 = 65 <= v305
	if cmp666 {
		goto land_lhs_true668
	} else {
		goto lor_lhs_false671
	}

land_lhs_true668:
	v306 = *libc.As[int32](lookahead)
	cmp669 = v306 <= 90
	if cmp669 {
		goto if_then680
	} else {
		goto lor_lhs_false671
	}

lor_lhs_false671:
	v307 = *libc.As[int32](lookahead)
	cmp672 = v307 == 95
	if cmp672 {
		goto if_then680
	} else {
		goto lor_lhs_false674
	}

lor_lhs_false674:
	v308 = *libc.As[int32](lookahead)
	cmp675 = 97 <= v308
	if cmp675 {
		goto land_lhs_true677
	} else {
		goto if_end681
	}

land_lhs_true677:
	v309 = *libc.As[int32](lookahead)
	cmp678 = v309 <= 122
	if cmp678 {
		goto if_then680
	} else {
		goto if_end681
	}

if_then680:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end681:
	v310 = *libc.As[byte](result)
	loadedv682 = (v310 & 1) != 0
	*libc.As[bool](retval) = loadedv682
	goto _return

sw_bb683:
	*libc.As[byte](result) = 1
	v311 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol684 = libc.Ptr(&libc.As[TSLexer](v311).F1)
	*libc.As[int16](result_symbol684) = 13
	v312 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end685 = libc.Ptr(&libc.As[TSLexer](v312).F3)
	v313 = *libc.As[unsafe.Pointer](mark_end685)
	v314 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v313)(v314)
	v315 = *libc.As[int32](lookahead)
	cmp686 = v315 == 101
	if cmp686 {
		goto if_then688
	} else {
		goto if_end689
	}

if_then688:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end689:
	v316 = *libc.As[int32](lookahead)
	cmp690 = 48 <= v316
	if cmp690 {
		goto land_lhs_true692
	} else {
		goto lor_lhs_false695
	}

land_lhs_true692:
	v317 = *libc.As[int32](lookahead)
	cmp693 = v317 <= 57
	if cmp693 {
		goto if_then710
	} else {
		goto lor_lhs_false695
	}

lor_lhs_false695:
	v318 = *libc.As[int32](lookahead)
	cmp696 = 65 <= v318
	if cmp696 {
		goto land_lhs_true698
	} else {
		goto lor_lhs_false701
	}

land_lhs_true698:
	v319 = *libc.As[int32](lookahead)
	cmp699 = v319 <= 90
	if cmp699 {
		goto if_then710
	} else {
		goto lor_lhs_false701
	}

lor_lhs_false701:
	v320 = *libc.As[int32](lookahead)
	cmp702 = v320 == 95
	if cmp702 {
		goto if_then710
	} else {
		goto lor_lhs_false704
	}

lor_lhs_false704:
	v321 = *libc.As[int32](lookahead)
	cmp705 = 97 <= v321
	if cmp705 {
		goto land_lhs_true707
	} else {
		goto if_end711
	}

land_lhs_true707:
	v322 = *libc.As[int32](lookahead)
	cmp708 = v322 <= 122
	if cmp708 {
		goto if_then710
	} else {
		goto if_end711
	}

if_then710:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end711:
	v323 = *libc.As[byte](result)
	loadedv712 = (v323 & 1) != 0
	*libc.As[bool](retval) = loadedv712
	goto _return

sw_bb713:
	*libc.As[byte](result) = 1
	v324 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol714 = libc.Ptr(&libc.As[TSLexer](v324).F1)
	*libc.As[int16](result_symbol714) = 13
	v325 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end715 = libc.Ptr(&libc.As[TSLexer](v325).F3)
	v326 = *libc.As[unsafe.Pointer](mark_end715)
	v327 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v326)(v327)
	v328 = *libc.As[int32](lookahead)
	cmp716 = v328 == 101
	if cmp716 {
		goto if_then718
	} else {
		goto if_end719
	}

if_then718:
	*libc.As[int16](state_addr) = 81
	goto next_state

if_end719:
	v329 = *libc.As[int32](lookahead)
	cmp720 = 48 <= v329
	if cmp720 {
		goto land_lhs_true722
	} else {
		goto lor_lhs_false725
	}

land_lhs_true722:
	v330 = *libc.As[int32](lookahead)
	cmp723 = v330 <= 57
	if cmp723 {
		goto if_then740
	} else {
		goto lor_lhs_false725
	}

lor_lhs_false725:
	v331 = *libc.As[int32](lookahead)
	cmp726 = 65 <= v331
	if cmp726 {
		goto land_lhs_true728
	} else {
		goto lor_lhs_false731
	}

land_lhs_true728:
	v332 = *libc.As[int32](lookahead)
	cmp729 = v332 <= 90
	if cmp729 {
		goto if_then740
	} else {
		goto lor_lhs_false731
	}

lor_lhs_false731:
	v333 = *libc.As[int32](lookahead)
	cmp732 = v333 == 95
	if cmp732 {
		goto if_then740
	} else {
		goto lor_lhs_false734
	}

lor_lhs_false734:
	v334 = *libc.As[int32](lookahead)
	cmp735 = 97 <= v334
	if cmp735 {
		goto land_lhs_true737
	} else {
		goto if_end741
	}

land_lhs_true737:
	v335 = *libc.As[int32](lookahead)
	cmp738 = v335 <= 122
	if cmp738 {
		goto if_then740
	} else {
		goto if_end741
	}

if_then740:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end741:
	v336 = *libc.As[byte](result)
	loadedv742 = (v336 & 1) != 0
	*libc.As[bool](retval) = loadedv742
	goto _return

sw_bb743:
	*libc.As[byte](result) = 1
	v337 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol744 = libc.Ptr(&libc.As[TSLexer](v337).F1)
	*libc.As[int16](result_symbol744) = 13
	v338 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end745 = libc.Ptr(&libc.As[TSLexer](v338).F3)
	v339 = *libc.As[unsafe.Pointer](mark_end745)
	v340 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v339)(v340)
	v341 = *libc.As[int32](lookahead)
	cmp746 = v341 == 101
	if cmp746 {
		goto if_then748
	} else {
		goto if_end749
	}

if_then748:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end749:
	v342 = *libc.As[int32](lookahead)
	cmp750 = 48 <= v342
	if cmp750 {
		goto land_lhs_true752
	} else {
		goto lor_lhs_false755
	}

land_lhs_true752:
	v343 = *libc.As[int32](lookahead)
	cmp753 = v343 <= 57
	if cmp753 {
		goto if_then770
	} else {
		goto lor_lhs_false755
	}

lor_lhs_false755:
	v344 = *libc.As[int32](lookahead)
	cmp756 = 65 <= v344
	if cmp756 {
		goto land_lhs_true758
	} else {
		goto lor_lhs_false761
	}

land_lhs_true758:
	v345 = *libc.As[int32](lookahead)
	cmp759 = v345 <= 90
	if cmp759 {
		goto if_then770
	} else {
		goto lor_lhs_false761
	}

lor_lhs_false761:
	v346 = *libc.As[int32](lookahead)
	cmp762 = v346 == 95
	if cmp762 {
		goto if_then770
	} else {
		goto lor_lhs_false764
	}

lor_lhs_false764:
	v347 = *libc.As[int32](lookahead)
	cmp765 = 97 <= v347
	if cmp765 {
		goto land_lhs_true767
	} else {
		goto if_end771
	}

land_lhs_true767:
	v348 = *libc.As[int32](lookahead)
	cmp768 = v348 <= 122
	if cmp768 {
		goto if_then770
	} else {
		goto if_end771
	}

if_then770:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end771:
	v349 = *libc.As[byte](result)
	loadedv772 = (v349 & 1) != 0
	*libc.As[bool](retval) = loadedv772
	goto _return

sw_bb773:
	*libc.As[byte](result) = 1
	v350 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol774 = libc.Ptr(&libc.As[TSLexer](v350).F1)
	*libc.As[int16](result_symbol774) = 13
	v351 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end775 = libc.Ptr(&libc.As[TSLexer](v351).F3)
	v352 = *libc.As[unsafe.Pointer](mark_end775)
	v353 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v352)(v353)
	v354 = *libc.As[int32](lookahead)
	cmp776 = v354 == 101
	if cmp776 {
		goto if_then778
	} else {
		goto if_end779
	}

if_then778:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end779:
	v355 = *libc.As[int32](lookahead)
	cmp780 = 48 <= v355
	if cmp780 {
		goto land_lhs_true782
	} else {
		goto lor_lhs_false785
	}

land_lhs_true782:
	v356 = *libc.As[int32](lookahead)
	cmp783 = v356 <= 57
	if cmp783 {
		goto if_then800
	} else {
		goto lor_lhs_false785
	}

lor_lhs_false785:
	v357 = *libc.As[int32](lookahead)
	cmp786 = 65 <= v357
	if cmp786 {
		goto land_lhs_true788
	} else {
		goto lor_lhs_false791
	}

land_lhs_true788:
	v358 = *libc.As[int32](lookahead)
	cmp789 = v358 <= 90
	if cmp789 {
		goto if_then800
	} else {
		goto lor_lhs_false791
	}

lor_lhs_false791:
	v359 = *libc.As[int32](lookahead)
	cmp792 = v359 == 95
	if cmp792 {
		goto if_then800
	} else {
		goto lor_lhs_false794
	}

lor_lhs_false794:
	v360 = *libc.As[int32](lookahead)
	cmp795 = 97 <= v360
	if cmp795 {
		goto land_lhs_true797
	} else {
		goto if_end801
	}

land_lhs_true797:
	v361 = *libc.As[int32](lookahead)
	cmp798 = v361 <= 122
	if cmp798 {
		goto if_then800
	} else {
		goto if_end801
	}

if_then800:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end801:
	v362 = *libc.As[byte](result)
	loadedv802 = (v362 & 1) != 0
	*libc.As[bool](retval) = loadedv802
	goto _return

sw_bb803:
	*libc.As[byte](result) = 1
	v363 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol804 = libc.Ptr(&libc.As[TSLexer](v363).F1)
	*libc.As[int16](result_symbol804) = 13
	v364 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end805 = libc.Ptr(&libc.As[TSLexer](v364).F3)
	v365 = *libc.As[unsafe.Pointer](mark_end805)
	v366 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v365)(v366)
	v367 = *libc.As[int32](lookahead)
	cmp806 = v367 == 101
	if cmp806 {
		goto if_then808
	} else {
		goto if_end809
	}

if_then808:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end809:
	v368 = *libc.As[int32](lookahead)
	cmp810 = 48 <= v368
	if cmp810 {
		goto land_lhs_true812
	} else {
		goto lor_lhs_false815
	}

land_lhs_true812:
	v369 = *libc.As[int32](lookahead)
	cmp813 = v369 <= 57
	if cmp813 {
		goto if_then830
	} else {
		goto lor_lhs_false815
	}

lor_lhs_false815:
	v370 = *libc.As[int32](lookahead)
	cmp816 = 65 <= v370
	if cmp816 {
		goto land_lhs_true818
	} else {
		goto lor_lhs_false821
	}

land_lhs_true818:
	v371 = *libc.As[int32](lookahead)
	cmp819 = v371 <= 90
	if cmp819 {
		goto if_then830
	} else {
		goto lor_lhs_false821
	}

lor_lhs_false821:
	v372 = *libc.As[int32](lookahead)
	cmp822 = v372 == 95
	if cmp822 {
		goto if_then830
	} else {
		goto lor_lhs_false824
	}

lor_lhs_false824:
	v373 = *libc.As[int32](lookahead)
	cmp825 = 97 <= v373
	if cmp825 {
		goto land_lhs_true827
	} else {
		goto if_end831
	}

land_lhs_true827:
	v374 = *libc.As[int32](lookahead)
	cmp828 = v374 <= 122
	if cmp828 {
		goto if_then830
	} else {
		goto if_end831
	}

if_then830:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end831:
	v375 = *libc.As[byte](result)
	loadedv832 = (v375 & 1) != 0
	*libc.As[bool](retval) = loadedv832
	goto _return

sw_bb833:
	*libc.As[byte](result) = 1
	v376 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol834 = libc.Ptr(&libc.As[TSLexer](v376).F1)
	*libc.As[int16](result_symbol834) = 13
	v377 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end835 = libc.Ptr(&libc.As[TSLexer](v377).F3)
	v378 = *libc.As[unsafe.Pointer](mark_end835)
	v379 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v378)(v379)
	v380 = *libc.As[int32](lookahead)
	cmp836 = v380 == 101
	if cmp836 {
		goto if_then838
	} else {
		goto if_end839
	}

if_then838:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end839:
	v381 = *libc.As[int32](lookahead)
	cmp840 = 48 <= v381
	if cmp840 {
		goto land_lhs_true842
	} else {
		goto lor_lhs_false845
	}

land_lhs_true842:
	v382 = *libc.As[int32](lookahead)
	cmp843 = v382 <= 57
	if cmp843 {
		goto if_then860
	} else {
		goto lor_lhs_false845
	}

lor_lhs_false845:
	v383 = *libc.As[int32](lookahead)
	cmp846 = 65 <= v383
	if cmp846 {
		goto land_lhs_true848
	} else {
		goto lor_lhs_false851
	}

land_lhs_true848:
	v384 = *libc.As[int32](lookahead)
	cmp849 = v384 <= 90
	if cmp849 {
		goto if_then860
	} else {
		goto lor_lhs_false851
	}

lor_lhs_false851:
	v385 = *libc.As[int32](lookahead)
	cmp852 = v385 == 95
	if cmp852 {
		goto if_then860
	} else {
		goto lor_lhs_false854
	}

lor_lhs_false854:
	v386 = *libc.As[int32](lookahead)
	cmp855 = 97 <= v386
	if cmp855 {
		goto land_lhs_true857
	} else {
		goto if_end861
	}

land_lhs_true857:
	v387 = *libc.As[int32](lookahead)
	cmp858 = v387 <= 122
	if cmp858 {
		goto if_then860
	} else {
		goto if_end861
	}

if_then860:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end861:
	v388 = *libc.As[byte](result)
	loadedv862 = (v388 & 1) != 0
	*libc.As[bool](retval) = loadedv862
	goto _return

sw_bb863:
	*libc.As[byte](result) = 1
	v389 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol864 = libc.Ptr(&libc.As[TSLexer](v389).F1)
	*libc.As[int16](result_symbol864) = 13
	v390 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end865 = libc.Ptr(&libc.As[TSLexer](v390).F3)
	v391 = *libc.As[unsafe.Pointer](mark_end865)
	v392 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v391)(v392)
	v393 = *libc.As[int32](lookahead)
	cmp866 = v393 == 101
	if cmp866 {
		goto if_then868
	} else {
		goto if_end869
	}

if_then868:
	*libc.As[int16](state_addr) = 61
	goto next_state

if_end869:
	v394 = *libc.As[int32](lookahead)
	cmp870 = 48 <= v394
	if cmp870 {
		goto land_lhs_true872
	} else {
		goto lor_lhs_false875
	}

land_lhs_true872:
	v395 = *libc.As[int32](lookahead)
	cmp873 = v395 <= 57
	if cmp873 {
		goto if_then890
	} else {
		goto lor_lhs_false875
	}

lor_lhs_false875:
	v396 = *libc.As[int32](lookahead)
	cmp876 = 65 <= v396
	if cmp876 {
		goto land_lhs_true878
	} else {
		goto lor_lhs_false881
	}

land_lhs_true878:
	v397 = *libc.As[int32](lookahead)
	cmp879 = v397 <= 90
	if cmp879 {
		goto if_then890
	} else {
		goto lor_lhs_false881
	}

lor_lhs_false881:
	v398 = *libc.As[int32](lookahead)
	cmp882 = v398 == 95
	if cmp882 {
		goto if_then890
	} else {
		goto lor_lhs_false884
	}

lor_lhs_false884:
	v399 = *libc.As[int32](lookahead)
	cmp885 = 97 <= v399
	if cmp885 {
		goto land_lhs_true887
	} else {
		goto if_end891
	}

land_lhs_true887:
	v400 = *libc.As[int32](lookahead)
	cmp888 = v400 <= 122
	if cmp888 {
		goto if_then890
	} else {
		goto if_end891
	}

if_then890:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end891:
	v401 = *libc.As[byte](result)
	loadedv892 = (v401 & 1) != 0
	*libc.As[bool](retval) = loadedv892
	goto _return

sw_bb893:
	*libc.As[byte](result) = 1
	v402 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol894 = libc.Ptr(&libc.As[TSLexer](v402).F1)
	*libc.As[int16](result_symbol894) = 13
	v403 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end895 = libc.Ptr(&libc.As[TSLexer](v403).F3)
	v404 = *libc.As[unsafe.Pointer](mark_end895)
	v405 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v404)(v405)
	v406 = *libc.As[int32](lookahead)
	cmp896 = v406 == 101
	if cmp896 {
		goto if_then898
	} else {
		goto if_end899
	}

if_then898:
	*libc.As[int16](state_addr) = 92
	goto next_state

if_end899:
	v407 = *libc.As[int32](lookahead)
	cmp900 = 48 <= v407
	if cmp900 {
		goto land_lhs_true902
	} else {
		goto lor_lhs_false905
	}

land_lhs_true902:
	v408 = *libc.As[int32](lookahead)
	cmp903 = v408 <= 57
	if cmp903 {
		goto if_then920
	} else {
		goto lor_lhs_false905
	}

lor_lhs_false905:
	v409 = *libc.As[int32](lookahead)
	cmp906 = 65 <= v409
	if cmp906 {
		goto land_lhs_true908
	} else {
		goto lor_lhs_false911
	}

land_lhs_true908:
	v410 = *libc.As[int32](lookahead)
	cmp909 = v410 <= 90
	if cmp909 {
		goto if_then920
	} else {
		goto lor_lhs_false911
	}

lor_lhs_false911:
	v411 = *libc.As[int32](lookahead)
	cmp912 = v411 == 95
	if cmp912 {
		goto if_then920
	} else {
		goto lor_lhs_false914
	}

lor_lhs_false914:
	v412 = *libc.As[int32](lookahead)
	cmp915 = 97 <= v412
	if cmp915 {
		goto land_lhs_true917
	} else {
		goto if_end921
	}

land_lhs_true917:
	v413 = *libc.As[int32](lookahead)
	cmp918 = v413 <= 122
	if cmp918 {
		goto if_then920
	} else {
		goto if_end921
	}

if_then920:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end921:
	v414 = *libc.As[byte](result)
	loadedv922 = (v414 & 1) != 0
	*libc.As[bool](retval) = loadedv922
	goto _return

sw_bb923:
	*libc.As[byte](result) = 1
	v415 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol924 = libc.Ptr(&libc.As[TSLexer](v415).F1)
	*libc.As[int16](result_symbol924) = 13
	v416 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end925 = libc.Ptr(&libc.As[TSLexer](v416).F3)
	v417 = *libc.As[unsafe.Pointer](mark_end925)
	v418 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v417)(v418)
	v419 = *libc.As[int32](lookahead)
	cmp926 = v419 == 101
	if cmp926 {
		goto if_then928
	} else {
		goto if_end929
	}

if_then928:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end929:
	v420 = *libc.As[int32](lookahead)
	cmp930 = 48 <= v420
	if cmp930 {
		goto land_lhs_true932
	} else {
		goto lor_lhs_false935
	}

land_lhs_true932:
	v421 = *libc.As[int32](lookahead)
	cmp933 = v421 <= 57
	if cmp933 {
		goto if_then950
	} else {
		goto lor_lhs_false935
	}

lor_lhs_false935:
	v422 = *libc.As[int32](lookahead)
	cmp936 = 65 <= v422
	if cmp936 {
		goto land_lhs_true938
	} else {
		goto lor_lhs_false941
	}

land_lhs_true938:
	v423 = *libc.As[int32](lookahead)
	cmp939 = v423 <= 90
	if cmp939 {
		goto if_then950
	} else {
		goto lor_lhs_false941
	}

lor_lhs_false941:
	v424 = *libc.As[int32](lookahead)
	cmp942 = v424 == 95
	if cmp942 {
		goto if_then950
	} else {
		goto lor_lhs_false944
	}

lor_lhs_false944:
	v425 = *libc.As[int32](lookahead)
	cmp945 = 97 <= v425
	if cmp945 {
		goto land_lhs_true947
	} else {
		goto if_end951
	}

land_lhs_true947:
	v426 = *libc.As[int32](lookahead)
	cmp948 = v426 <= 122
	if cmp948 {
		goto if_then950
	} else {
		goto if_end951
	}

if_then950:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end951:
	v427 = *libc.As[byte](result)
	loadedv952 = (v427 & 1) != 0
	*libc.As[bool](retval) = loadedv952
	goto _return

sw_bb953:
	*libc.As[byte](result) = 1
	v428 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol954 = libc.Ptr(&libc.As[TSLexer](v428).F1)
	*libc.As[int16](result_symbol954) = 13
	v429 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end955 = libc.Ptr(&libc.As[TSLexer](v429).F3)
	v430 = *libc.As[unsafe.Pointer](mark_end955)
	v431 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v430)(v431)
	v432 = *libc.As[int32](lookahead)
	cmp956 = v432 == 101
	if cmp956 {
		goto if_then958
	} else {
		goto if_end959
	}

if_then958:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end959:
	v433 = *libc.As[int32](lookahead)
	cmp960 = 48 <= v433
	if cmp960 {
		goto land_lhs_true962
	} else {
		goto lor_lhs_false965
	}

land_lhs_true962:
	v434 = *libc.As[int32](lookahead)
	cmp963 = v434 <= 57
	if cmp963 {
		goto if_then980
	} else {
		goto lor_lhs_false965
	}

lor_lhs_false965:
	v435 = *libc.As[int32](lookahead)
	cmp966 = 65 <= v435
	if cmp966 {
		goto land_lhs_true968
	} else {
		goto lor_lhs_false971
	}

land_lhs_true968:
	v436 = *libc.As[int32](lookahead)
	cmp969 = v436 <= 90
	if cmp969 {
		goto if_then980
	} else {
		goto lor_lhs_false971
	}

lor_lhs_false971:
	v437 = *libc.As[int32](lookahead)
	cmp972 = v437 == 95
	if cmp972 {
		goto if_then980
	} else {
		goto lor_lhs_false974
	}

lor_lhs_false974:
	v438 = *libc.As[int32](lookahead)
	cmp975 = 97 <= v438
	if cmp975 {
		goto land_lhs_true977
	} else {
		goto if_end981
	}

land_lhs_true977:
	v439 = *libc.As[int32](lookahead)
	cmp978 = v439 <= 122
	if cmp978 {
		goto if_then980
	} else {
		goto if_end981
	}

if_then980:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end981:
	v440 = *libc.As[byte](result)
	loadedv982 = (v440 & 1) != 0
	*libc.As[bool](retval) = loadedv982
	goto _return

sw_bb983:
	*libc.As[byte](result) = 1
	v441 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol984 = libc.Ptr(&libc.As[TSLexer](v441).F1)
	*libc.As[int16](result_symbol984) = 13
	v442 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end985 = libc.Ptr(&libc.As[TSLexer](v442).F3)
	v443 = *libc.As[unsafe.Pointer](mark_end985)
	v444 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v443)(v444)
	v445 = *libc.As[int32](lookahead)
	cmp986 = v445 == 101
	if cmp986 {
		goto if_then988
	} else {
		goto if_end989
	}

if_then988:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end989:
	v446 = *libc.As[int32](lookahead)
	cmp990 = 48 <= v446
	if cmp990 {
		goto land_lhs_true992
	} else {
		goto lor_lhs_false995
	}

land_lhs_true992:
	v447 = *libc.As[int32](lookahead)
	cmp993 = v447 <= 57
	if cmp993 {
		goto if_then1010
	} else {
		goto lor_lhs_false995
	}

lor_lhs_false995:
	v448 = *libc.As[int32](lookahead)
	cmp996 = 65 <= v448
	if cmp996 {
		goto land_lhs_true998
	} else {
		goto lor_lhs_false1001
	}

land_lhs_true998:
	v449 = *libc.As[int32](lookahead)
	cmp999 = v449 <= 90
	if cmp999 {
		goto if_then1010
	} else {
		goto lor_lhs_false1001
	}

lor_lhs_false1001:
	v450 = *libc.As[int32](lookahead)
	cmp1002 = v450 == 95
	if cmp1002 {
		goto if_then1010
	} else {
		goto lor_lhs_false1004
	}

lor_lhs_false1004:
	v451 = *libc.As[int32](lookahead)
	cmp1005 = 97 <= v451
	if cmp1005 {
		goto land_lhs_true1007
	} else {
		goto if_end1011
	}

land_lhs_true1007:
	v452 = *libc.As[int32](lookahead)
	cmp1008 = v452 <= 122
	if cmp1008 {
		goto if_then1010
	} else {
		goto if_end1011
	}

if_then1010:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1011:
	v453 = *libc.As[byte](result)
	loadedv1012 = (v453 & 1) != 0
	*libc.As[bool](retval) = loadedv1012
	goto _return

sw_bb1013:
	*libc.As[byte](result) = 1
	v454 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1014 = libc.Ptr(&libc.As[TSLexer](v454).F1)
	*libc.As[int16](result_symbol1014) = 13
	v455 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1015 = libc.Ptr(&libc.As[TSLexer](v455).F3)
	v456 = *libc.As[unsafe.Pointer](mark_end1015)
	v457 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v456)(v457)
	v458 = *libc.As[int32](lookahead)
	cmp1016 = v458 == 102
	if cmp1016 {
		goto if_then1018
	} else {
		goto if_end1019
	}

if_then1018:
	*libc.As[int16](state_addr) = 69
	goto next_state

if_end1019:
	v459 = *libc.As[int32](lookahead)
	cmp1020 = 48 <= v459
	if cmp1020 {
		goto land_lhs_true1022
	} else {
		goto lor_lhs_false1025
	}

land_lhs_true1022:
	v460 = *libc.As[int32](lookahead)
	cmp1023 = v460 <= 57
	if cmp1023 {
		goto if_then1040
	} else {
		goto lor_lhs_false1025
	}

lor_lhs_false1025:
	v461 = *libc.As[int32](lookahead)
	cmp1026 = 65 <= v461
	if cmp1026 {
		goto land_lhs_true1028
	} else {
		goto lor_lhs_false1031
	}

land_lhs_true1028:
	v462 = *libc.As[int32](lookahead)
	cmp1029 = v462 <= 90
	if cmp1029 {
		goto if_then1040
	} else {
		goto lor_lhs_false1031
	}

lor_lhs_false1031:
	v463 = *libc.As[int32](lookahead)
	cmp1032 = v463 == 95
	if cmp1032 {
		goto if_then1040
	} else {
		goto lor_lhs_false1034
	}

lor_lhs_false1034:
	v464 = *libc.As[int32](lookahead)
	cmp1035 = 97 <= v464
	if cmp1035 {
		goto land_lhs_true1037
	} else {
		goto if_end1041
	}

land_lhs_true1037:
	v465 = *libc.As[int32](lookahead)
	cmp1038 = v465 <= 122
	if cmp1038 {
		goto if_then1040
	} else {
		goto if_end1041
	}

if_then1040:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1041:
	v466 = *libc.As[byte](result)
	loadedv1042 = (v466 & 1) != 0
	*libc.As[bool](retval) = loadedv1042
	goto _return

sw_bb1043:
	*libc.As[byte](result) = 1
	v467 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1044 = libc.Ptr(&libc.As[TSLexer](v467).F1)
	*libc.As[int16](result_symbol1044) = 13
	v468 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1045 = libc.Ptr(&libc.As[TSLexer](v468).F3)
	v469 = *libc.As[unsafe.Pointer](mark_end1045)
	v470 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v469)(v470)
	v471 = *libc.As[int32](lookahead)
	cmp1046 = v471 == 102
	if cmp1046 {
		goto if_then1048
	} else {
		goto if_end1049
	}

if_then1048:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1049:
	v472 = *libc.As[int32](lookahead)
	cmp1050 = 48 <= v472
	if cmp1050 {
		goto land_lhs_true1052
	} else {
		goto lor_lhs_false1055
	}

land_lhs_true1052:
	v473 = *libc.As[int32](lookahead)
	cmp1053 = v473 <= 57
	if cmp1053 {
		goto if_then1070
	} else {
		goto lor_lhs_false1055
	}

lor_lhs_false1055:
	v474 = *libc.As[int32](lookahead)
	cmp1056 = 65 <= v474
	if cmp1056 {
		goto land_lhs_true1058
	} else {
		goto lor_lhs_false1061
	}

land_lhs_true1058:
	v475 = *libc.As[int32](lookahead)
	cmp1059 = v475 <= 90
	if cmp1059 {
		goto if_then1070
	} else {
		goto lor_lhs_false1061
	}

lor_lhs_false1061:
	v476 = *libc.As[int32](lookahead)
	cmp1062 = v476 == 95
	if cmp1062 {
		goto if_then1070
	} else {
		goto lor_lhs_false1064
	}

lor_lhs_false1064:
	v477 = *libc.As[int32](lookahead)
	cmp1065 = 97 <= v477
	if cmp1065 {
		goto land_lhs_true1067
	} else {
		goto if_end1071
	}

land_lhs_true1067:
	v478 = *libc.As[int32](lookahead)
	cmp1068 = v478 <= 122
	if cmp1068 {
		goto if_then1070
	} else {
		goto if_end1071
	}

if_then1070:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1071:
	v479 = *libc.As[byte](result)
	loadedv1072 = (v479 & 1) != 0
	*libc.As[bool](retval) = loadedv1072
	goto _return

sw_bb1073:
	*libc.As[byte](result) = 1
	v480 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1074 = libc.Ptr(&libc.As[TSLexer](v480).F1)
	*libc.As[int16](result_symbol1074) = 13
	v481 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1075 = libc.Ptr(&libc.As[TSLexer](v481).F3)
	v482 = *libc.As[unsafe.Pointer](mark_end1075)
	v483 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v482)(v483)
	v484 = *libc.As[int32](lookahead)
	cmp1076 = v484 == 103
	if cmp1076 {
		goto if_then1078
	} else {
		goto if_end1079
	}

if_then1078:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end1079:
	v485 = *libc.As[int32](lookahead)
	cmp1080 = 48 <= v485
	if cmp1080 {
		goto land_lhs_true1082
	} else {
		goto lor_lhs_false1085
	}

land_lhs_true1082:
	v486 = *libc.As[int32](lookahead)
	cmp1083 = v486 <= 57
	if cmp1083 {
		goto if_then1100
	} else {
		goto lor_lhs_false1085
	}

lor_lhs_false1085:
	v487 = *libc.As[int32](lookahead)
	cmp1086 = 65 <= v487
	if cmp1086 {
		goto land_lhs_true1088
	} else {
		goto lor_lhs_false1091
	}

land_lhs_true1088:
	v488 = *libc.As[int32](lookahead)
	cmp1089 = v488 <= 90
	if cmp1089 {
		goto if_then1100
	} else {
		goto lor_lhs_false1091
	}

lor_lhs_false1091:
	v489 = *libc.As[int32](lookahead)
	cmp1092 = v489 == 95
	if cmp1092 {
		goto if_then1100
	} else {
		goto lor_lhs_false1094
	}

lor_lhs_false1094:
	v490 = *libc.As[int32](lookahead)
	cmp1095 = 97 <= v490
	if cmp1095 {
		goto land_lhs_true1097
	} else {
		goto if_end1101
	}

land_lhs_true1097:
	v491 = *libc.As[int32](lookahead)
	cmp1098 = v491 <= 122
	if cmp1098 {
		goto if_then1100
	} else {
		goto if_end1101
	}

if_then1100:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1101:
	v492 = *libc.As[byte](result)
	loadedv1102 = (v492 & 1) != 0
	*libc.As[bool](retval) = loadedv1102
	goto _return

sw_bb1103:
	*libc.As[byte](result) = 1
	v493 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1104 = libc.Ptr(&libc.As[TSLexer](v493).F1)
	*libc.As[int16](result_symbol1104) = 13
	v494 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1105 = libc.Ptr(&libc.As[TSLexer](v494).F3)
	v495 = *libc.As[unsafe.Pointer](mark_end1105)
	v496 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v495)(v496)
	v497 = *libc.As[int32](lookahead)
	cmp1106 = v497 == 103
	if cmp1106 {
		goto if_then1108
	} else {
		goto if_end1109
	}

if_then1108:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end1109:
	v498 = *libc.As[int32](lookahead)
	cmp1110 = 48 <= v498
	if cmp1110 {
		goto land_lhs_true1112
	} else {
		goto lor_lhs_false1115
	}

land_lhs_true1112:
	v499 = *libc.As[int32](lookahead)
	cmp1113 = v499 <= 57
	if cmp1113 {
		goto if_then1130
	} else {
		goto lor_lhs_false1115
	}

lor_lhs_false1115:
	v500 = *libc.As[int32](lookahead)
	cmp1116 = 65 <= v500
	if cmp1116 {
		goto land_lhs_true1118
	} else {
		goto lor_lhs_false1121
	}

land_lhs_true1118:
	v501 = *libc.As[int32](lookahead)
	cmp1119 = v501 <= 90
	if cmp1119 {
		goto if_then1130
	} else {
		goto lor_lhs_false1121
	}

lor_lhs_false1121:
	v502 = *libc.As[int32](lookahead)
	cmp1122 = v502 == 95
	if cmp1122 {
		goto if_then1130
	} else {
		goto lor_lhs_false1124
	}

lor_lhs_false1124:
	v503 = *libc.As[int32](lookahead)
	cmp1125 = 97 <= v503
	if cmp1125 {
		goto land_lhs_true1127
	} else {
		goto if_end1131
	}

land_lhs_true1127:
	v504 = *libc.As[int32](lookahead)
	cmp1128 = v504 <= 122
	if cmp1128 {
		goto if_then1130
	} else {
		goto if_end1131
	}

if_then1130:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1131:
	v505 = *libc.As[byte](result)
	loadedv1132 = (v505 & 1) != 0
	*libc.As[bool](retval) = loadedv1132
	goto _return

sw_bb1133:
	*libc.As[byte](result) = 1
	v506 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1134 = libc.Ptr(&libc.As[TSLexer](v506).F1)
	*libc.As[int16](result_symbol1134) = 13
	v507 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1135 = libc.Ptr(&libc.As[TSLexer](v507).F3)
	v508 = *libc.As[unsafe.Pointer](mark_end1135)
	v509 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v508)(v509)
	v510 = *libc.As[int32](lookahead)
	cmp1136 = v510 == 103
	if cmp1136 {
		goto if_then1138
	} else {
		goto if_end1139
	}

if_then1138:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end1139:
	v511 = *libc.As[int32](lookahead)
	cmp1140 = 48 <= v511
	if cmp1140 {
		goto land_lhs_true1142
	} else {
		goto lor_lhs_false1145
	}

land_lhs_true1142:
	v512 = *libc.As[int32](lookahead)
	cmp1143 = v512 <= 57
	if cmp1143 {
		goto if_then1160
	} else {
		goto lor_lhs_false1145
	}

lor_lhs_false1145:
	v513 = *libc.As[int32](lookahead)
	cmp1146 = 65 <= v513
	if cmp1146 {
		goto land_lhs_true1148
	} else {
		goto lor_lhs_false1151
	}

land_lhs_true1148:
	v514 = *libc.As[int32](lookahead)
	cmp1149 = v514 <= 90
	if cmp1149 {
		goto if_then1160
	} else {
		goto lor_lhs_false1151
	}

lor_lhs_false1151:
	v515 = *libc.As[int32](lookahead)
	cmp1152 = v515 == 95
	if cmp1152 {
		goto if_then1160
	} else {
		goto lor_lhs_false1154
	}

lor_lhs_false1154:
	v516 = *libc.As[int32](lookahead)
	cmp1155 = 97 <= v516
	if cmp1155 {
		goto land_lhs_true1157
	} else {
		goto if_end1161
	}

land_lhs_true1157:
	v517 = *libc.As[int32](lookahead)
	cmp1158 = v517 <= 122
	if cmp1158 {
		goto if_then1160
	} else {
		goto if_end1161
	}

if_then1160:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1161:
	v518 = *libc.As[byte](result)
	loadedv1162 = (v518 & 1) != 0
	*libc.As[bool](retval) = loadedv1162
	goto _return

sw_bb1163:
	*libc.As[byte](result) = 1
	v519 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1164 = libc.Ptr(&libc.As[TSLexer](v519).F1)
	*libc.As[int16](result_symbol1164) = 13
	v520 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1165 = libc.Ptr(&libc.As[TSLexer](v520).F3)
	v521 = *libc.As[unsafe.Pointer](mark_end1165)
	v522 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v521)(v522)
	v523 = *libc.As[int32](lookahead)
	cmp1166 = v523 == 103
	if cmp1166 {
		goto if_then1168
	} else {
		goto if_end1169
	}

if_then1168:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end1169:
	v524 = *libc.As[int32](lookahead)
	cmp1170 = 48 <= v524
	if cmp1170 {
		goto land_lhs_true1172
	} else {
		goto lor_lhs_false1175
	}

land_lhs_true1172:
	v525 = *libc.As[int32](lookahead)
	cmp1173 = v525 <= 57
	if cmp1173 {
		goto if_then1190
	} else {
		goto lor_lhs_false1175
	}

lor_lhs_false1175:
	v526 = *libc.As[int32](lookahead)
	cmp1176 = 65 <= v526
	if cmp1176 {
		goto land_lhs_true1178
	} else {
		goto lor_lhs_false1181
	}

land_lhs_true1178:
	v527 = *libc.As[int32](lookahead)
	cmp1179 = v527 <= 90
	if cmp1179 {
		goto if_then1190
	} else {
		goto lor_lhs_false1181
	}

lor_lhs_false1181:
	v528 = *libc.As[int32](lookahead)
	cmp1182 = v528 == 95
	if cmp1182 {
		goto if_then1190
	} else {
		goto lor_lhs_false1184
	}

lor_lhs_false1184:
	v529 = *libc.As[int32](lookahead)
	cmp1185 = 97 <= v529
	if cmp1185 {
		goto land_lhs_true1187
	} else {
		goto if_end1191
	}

land_lhs_true1187:
	v530 = *libc.As[int32](lookahead)
	cmp1188 = v530 <= 122
	if cmp1188 {
		goto if_then1190
	} else {
		goto if_end1191
	}

if_then1190:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1191:
	v531 = *libc.As[byte](result)
	loadedv1192 = (v531 & 1) != 0
	*libc.As[bool](retval) = loadedv1192
	goto _return

sw_bb1193:
	*libc.As[byte](result) = 1
	v532 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1194 = libc.Ptr(&libc.As[TSLexer](v532).F1)
	*libc.As[int16](result_symbol1194) = 13
	v533 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1195 = libc.Ptr(&libc.As[TSLexer](v533).F3)
	v534 = *libc.As[unsafe.Pointer](mark_end1195)
	v535 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v534)(v535)
	v536 = *libc.As[int32](lookahead)
	cmp1196 = v536 == 105
	if cmp1196 {
		goto if_then1198
	} else {
		goto if_end1199
	}

if_then1198:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end1199:
	v537 = *libc.As[int32](lookahead)
	cmp1200 = 48 <= v537
	if cmp1200 {
		goto land_lhs_true1202
	} else {
		goto lor_lhs_false1205
	}

land_lhs_true1202:
	v538 = *libc.As[int32](lookahead)
	cmp1203 = v538 <= 57
	if cmp1203 {
		goto if_then1220
	} else {
		goto lor_lhs_false1205
	}

lor_lhs_false1205:
	v539 = *libc.As[int32](lookahead)
	cmp1206 = 65 <= v539
	if cmp1206 {
		goto land_lhs_true1208
	} else {
		goto lor_lhs_false1211
	}

land_lhs_true1208:
	v540 = *libc.As[int32](lookahead)
	cmp1209 = v540 <= 90
	if cmp1209 {
		goto if_then1220
	} else {
		goto lor_lhs_false1211
	}

lor_lhs_false1211:
	v541 = *libc.As[int32](lookahead)
	cmp1212 = v541 == 95
	if cmp1212 {
		goto if_then1220
	} else {
		goto lor_lhs_false1214
	}

lor_lhs_false1214:
	v542 = *libc.As[int32](lookahead)
	cmp1215 = 97 <= v542
	if cmp1215 {
		goto land_lhs_true1217
	} else {
		goto if_end1221
	}

land_lhs_true1217:
	v543 = *libc.As[int32](lookahead)
	cmp1218 = v543 <= 122
	if cmp1218 {
		goto if_then1220
	} else {
		goto if_end1221
	}

if_then1220:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1221:
	v544 = *libc.As[byte](result)
	loadedv1222 = (v544 & 1) != 0
	*libc.As[bool](retval) = loadedv1222
	goto _return

sw_bb1223:
	*libc.As[byte](result) = 1
	v545 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1224 = libc.Ptr(&libc.As[TSLexer](v545).F1)
	*libc.As[int16](result_symbol1224) = 13
	v546 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1225 = libc.Ptr(&libc.As[TSLexer](v546).F3)
	v547 = *libc.As[unsafe.Pointer](mark_end1225)
	v548 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v547)(v548)
	v549 = *libc.As[int32](lookahead)
	cmp1226 = v549 == 105
	if cmp1226 {
		goto if_then1228
	} else {
		goto if_end1229
	}

if_then1228:
	*libc.As[int16](state_addr) = 71
	goto next_state

if_end1229:
	v550 = *libc.As[int32](lookahead)
	cmp1230 = 48 <= v550
	if cmp1230 {
		goto land_lhs_true1232
	} else {
		goto lor_lhs_false1235
	}

land_lhs_true1232:
	v551 = *libc.As[int32](lookahead)
	cmp1233 = v551 <= 57
	if cmp1233 {
		goto if_then1250
	} else {
		goto lor_lhs_false1235
	}

lor_lhs_false1235:
	v552 = *libc.As[int32](lookahead)
	cmp1236 = 65 <= v552
	if cmp1236 {
		goto land_lhs_true1238
	} else {
		goto lor_lhs_false1241
	}

land_lhs_true1238:
	v553 = *libc.As[int32](lookahead)
	cmp1239 = v553 <= 90
	if cmp1239 {
		goto if_then1250
	} else {
		goto lor_lhs_false1241
	}

lor_lhs_false1241:
	v554 = *libc.As[int32](lookahead)
	cmp1242 = v554 == 95
	if cmp1242 {
		goto if_then1250
	} else {
		goto lor_lhs_false1244
	}

lor_lhs_false1244:
	v555 = *libc.As[int32](lookahead)
	cmp1245 = 97 <= v555
	if cmp1245 {
		goto land_lhs_true1247
	} else {
		goto if_end1251
	}

land_lhs_true1247:
	v556 = *libc.As[int32](lookahead)
	cmp1248 = v556 <= 122
	if cmp1248 {
		goto if_then1250
	} else {
		goto if_end1251
	}

if_then1250:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1251:
	v557 = *libc.As[byte](result)
	loadedv1252 = (v557 & 1) != 0
	*libc.As[bool](retval) = loadedv1252
	goto _return

sw_bb1253:
	*libc.As[byte](result) = 1
	v558 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1254 = libc.Ptr(&libc.As[TSLexer](v558).F1)
	*libc.As[int16](result_symbol1254) = 13
	v559 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1255 = libc.Ptr(&libc.As[TSLexer](v559).F3)
	v560 = *libc.As[unsafe.Pointer](mark_end1255)
	v561 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v560)(v561)
	v562 = *libc.As[int32](lookahead)
	cmp1256 = v562 == 105
	if cmp1256 {
		goto if_then1258
	} else {
		goto if_end1259
	}

if_then1258:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end1259:
	v563 = *libc.As[int32](lookahead)
	cmp1260 = 48 <= v563
	if cmp1260 {
		goto land_lhs_true1262
	} else {
		goto lor_lhs_false1265
	}

land_lhs_true1262:
	v564 = *libc.As[int32](lookahead)
	cmp1263 = v564 <= 57
	if cmp1263 {
		goto if_then1280
	} else {
		goto lor_lhs_false1265
	}

lor_lhs_false1265:
	v565 = *libc.As[int32](lookahead)
	cmp1266 = 65 <= v565
	if cmp1266 {
		goto land_lhs_true1268
	} else {
		goto lor_lhs_false1271
	}

land_lhs_true1268:
	v566 = *libc.As[int32](lookahead)
	cmp1269 = v566 <= 90
	if cmp1269 {
		goto if_then1280
	} else {
		goto lor_lhs_false1271
	}

lor_lhs_false1271:
	v567 = *libc.As[int32](lookahead)
	cmp1272 = v567 == 95
	if cmp1272 {
		goto if_then1280
	} else {
		goto lor_lhs_false1274
	}

lor_lhs_false1274:
	v568 = *libc.As[int32](lookahead)
	cmp1275 = 97 <= v568
	if cmp1275 {
		goto land_lhs_true1277
	} else {
		goto if_end1281
	}

land_lhs_true1277:
	v569 = *libc.As[int32](lookahead)
	cmp1278 = v569 <= 122
	if cmp1278 {
		goto if_then1280
	} else {
		goto if_end1281
	}

if_then1280:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1281:
	v570 = *libc.As[byte](result)
	loadedv1282 = (v570 & 1) != 0
	*libc.As[bool](retval) = loadedv1282
	goto _return

sw_bb1283:
	*libc.As[byte](result) = 1
	v571 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1284 = libc.Ptr(&libc.As[TSLexer](v571).F1)
	*libc.As[int16](result_symbol1284) = 13
	v572 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1285 = libc.Ptr(&libc.As[TSLexer](v572).F3)
	v573 = *libc.As[unsafe.Pointer](mark_end1285)
	v574 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v573)(v574)
	v575 = *libc.As[int32](lookahead)
	cmp1286 = v575 == 105
	if cmp1286 {
		goto if_then1288
	} else {
		goto if_end1289
	}

if_then1288:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end1289:
	v576 = *libc.As[int32](lookahead)
	cmp1290 = 48 <= v576
	if cmp1290 {
		goto land_lhs_true1292
	} else {
		goto lor_lhs_false1295
	}

land_lhs_true1292:
	v577 = *libc.As[int32](lookahead)
	cmp1293 = v577 <= 57
	if cmp1293 {
		goto if_then1310
	} else {
		goto lor_lhs_false1295
	}

lor_lhs_false1295:
	v578 = *libc.As[int32](lookahead)
	cmp1296 = 65 <= v578
	if cmp1296 {
		goto land_lhs_true1298
	} else {
		goto lor_lhs_false1301
	}

land_lhs_true1298:
	v579 = *libc.As[int32](lookahead)
	cmp1299 = v579 <= 90
	if cmp1299 {
		goto if_then1310
	} else {
		goto lor_lhs_false1301
	}

lor_lhs_false1301:
	v580 = *libc.As[int32](lookahead)
	cmp1302 = v580 == 95
	if cmp1302 {
		goto if_then1310
	} else {
		goto lor_lhs_false1304
	}

lor_lhs_false1304:
	v581 = *libc.As[int32](lookahead)
	cmp1305 = 97 <= v581
	if cmp1305 {
		goto land_lhs_true1307
	} else {
		goto if_end1311
	}

land_lhs_true1307:
	v582 = *libc.As[int32](lookahead)
	cmp1308 = v582 <= 122
	if cmp1308 {
		goto if_then1310
	} else {
		goto if_end1311
	}

if_then1310:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1311:
	v583 = *libc.As[byte](result)
	loadedv1312 = (v583 & 1) != 0
	*libc.As[bool](retval) = loadedv1312
	goto _return

sw_bb1313:
	*libc.As[byte](result) = 1
	v584 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1314 = libc.Ptr(&libc.As[TSLexer](v584).F1)
	*libc.As[int16](result_symbol1314) = 13
	v585 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1315 = libc.Ptr(&libc.As[TSLexer](v585).F3)
	v586 = *libc.As[unsafe.Pointer](mark_end1315)
	v587 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v586)(v587)
	v588 = *libc.As[int32](lookahead)
	cmp1316 = v588 == 105
	if cmp1316 {
		goto if_then1318
	} else {
		goto if_end1319
	}

if_then1318:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end1319:
	v589 = *libc.As[int32](lookahead)
	cmp1320 = 48 <= v589
	if cmp1320 {
		goto land_lhs_true1322
	} else {
		goto lor_lhs_false1325
	}

land_lhs_true1322:
	v590 = *libc.As[int32](lookahead)
	cmp1323 = v590 <= 57
	if cmp1323 {
		goto if_then1340
	} else {
		goto lor_lhs_false1325
	}

lor_lhs_false1325:
	v591 = *libc.As[int32](lookahead)
	cmp1326 = 65 <= v591
	if cmp1326 {
		goto land_lhs_true1328
	} else {
		goto lor_lhs_false1331
	}

land_lhs_true1328:
	v592 = *libc.As[int32](lookahead)
	cmp1329 = v592 <= 90
	if cmp1329 {
		goto if_then1340
	} else {
		goto lor_lhs_false1331
	}

lor_lhs_false1331:
	v593 = *libc.As[int32](lookahead)
	cmp1332 = v593 == 95
	if cmp1332 {
		goto if_then1340
	} else {
		goto lor_lhs_false1334
	}

lor_lhs_false1334:
	v594 = *libc.As[int32](lookahead)
	cmp1335 = 97 <= v594
	if cmp1335 {
		goto land_lhs_true1337
	} else {
		goto if_end1341
	}

land_lhs_true1337:
	v595 = *libc.As[int32](lookahead)
	cmp1338 = v595 <= 122
	if cmp1338 {
		goto if_then1340
	} else {
		goto if_end1341
	}

if_then1340:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1341:
	v596 = *libc.As[byte](result)
	loadedv1342 = (v596 & 1) != 0
	*libc.As[bool](retval) = loadedv1342
	goto _return

sw_bb1343:
	*libc.As[byte](result) = 1
	v597 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1344 = libc.Ptr(&libc.As[TSLexer](v597).F1)
	*libc.As[int16](result_symbol1344) = 13
	v598 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1345 = libc.Ptr(&libc.As[TSLexer](v598).F3)
	v599 = *libc.As[unsafe.Pointer](mark_end1345)
	v600 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v599)(v600)
	v601 = *libc.As[int32](lookahead)
	cmp1346 = v601 == 105
	if cmp1346 {
		goto if_then1348
	} else {
		goto if_end1349
	}

if_then1348:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end1349:
	v602 = *libc.As[int32](lookahead)
	cmp1350 = 48 <= v602
	if cmp1350 {
		goto land_lhs_true1352
	} else {
		goto lor_lhs_false1355
	}

land_lhs_true1352:
	v603 = *libc.As[int32](lookahead)
	cmp1353 = v603 <= 57
	if cmp1353 {
		goto if_then1370
	} else {
		goto lor_lhs_false1355
	}

lor_lhs_false1355:
	v604 = *libc.As[int32](lookahead)
	cmp1356 = 65 <= v604
	if cmp1356 {
		goto land_lhs_true1358
	} else {
		goto lor_lhs_false1361
	}

land_lhs_true1358:
	v605 = *libc.As[int32](lookahead)
	cmp1359 = v605 <= 90
	if cmp1359 {
		goto if_then1370
	} else {
		goto lor_lhs_false1361
	}

lor_lhs_false1361:
	v606 = *libc.As[int32](lookahead)
	cmp1362 = v606 == 95
	if cmp1362 {
		goto if_then1370
	} else {
		goto lor_lhs_false1364
	}

lor_lhs_false1364:
	v607 = *libc.As[int32](lookahead)
	cmp1365 = 97 <= v607
	if cmp1365 {
		goto land_lhs_true1367
	} else {
		goto if_end1371
	}

land_lhs_true1367:
	v608 = *libc.As[int32](lookahead)
	cmp1368 = v608 <= 122
	if cmp1368 {
		goto if_then1370
	} else {
		goto if_end1371
	}

if_then1370:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1371:
	v609 = *libc.As[byte](result)
	loadedv1372 = (v609 & 1) != 0
	*libc.As[bool](retval) = loadedv1372
	goto _return

sw_bb1373:
	*libc.As[byte](result) = 1
	v610 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1374 = libc.Ptr(&libc.As[TSLexer](v610).F1)
	*libc.As[int16](result_symbol1374) = 13
	v611 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1375 = libc.Ptr(&libc.As[TSLexer](v611).F3)
	v612 = *libc.As[unsafe.Pointer](mark_end1375)
	v613 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v612)(v613)
	v614 = *libc.As[int32](lookahead)
	cmp1376 = v614 == 107
	if cmp1376 {
		goto if_then1378
	} else {
		goto if_end1379
	}

if_then1378:
	*libc.As[int16](state_addr) = 88
	goto next_state

if_end1379:
	v615 = *libc.As[int32](lookahead)
	cmp1380 = 48 <= v615
	if cmp1380 {
		goto land_lhs_true1382
	} else {
		goto lor_lhs_false1385
	}

land_lhs_true1382:
	v616 = *libc.As[int32](lookahead)
	cmp1383 = v616 <= 57
	if cmp1383 {
		goto if_then1400
	} else {
		goto lor_lhs_false1385
	}

lor_lhs_false1385:
	v617 = *libc.As[int32](lookahead)
	cmp1386 = 65 <= v617
	if cmp1386 {
		goto land_lhs_true1388
	} else {
		goto lor_lhs_false1391
	}

land_lhs_true1388:
	v618 = *libc.As[int32](lookahead)
	cmp1389 = v618 <= 90
	if cmp1389 {
		goto if_then1400
	} else {
		goto lor_lhs_false1391
	}

lor_lhs_false1391:
	v619 = *libc.As[int32](lookahead)
	cmp1392 = v619 == 95
	if cmp1392 {
		goto if_then1400
	} else {
		goto lor_lhs_false1394
	}

lor_lhs_false1394:
	v620 = *libc.As[int32](lookahead)
	cmp1395 = 97 <= v620
	if cmp1395 {
		goto land_lhs_true1397
	} else {
		goto if_end1401
	}

land_lhs_true1397:
	v621 = *libc.As[int32](lookahead)
	cmp1398 = v621 <= 122
	if cmp1398 {
		goto if_then1400
	} else {
		goto if_end1401
	}

if_then1400:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1401:
	v622 = *libc.As[byte](result)
	loadedv1402 = (v622 & 1) != 0
	*libc.As[bool](retval) = loadedv1402
	goto _return

sw_bb1403:
	*libc.As[byte](result) = 1
	v623 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1404 = libc.Ptr(&libc.As[TSLexer](v623).F1)
	*libc.As[int16](result_symbol1404) = 13
	v624 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1405 = libc.Ptr(&libc.As[TSLexer](v624).F3)
	v625 = *libc.As[unsafe.Pointer](mark_end1405)
	v626 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v625)(v626)
	v627 = *libc.As[int32](lookahead)
	cmp1406 = v627 == 108
	if cmp1406 {
		goto if_then1408
	} else {
		goto if_end1409
	}

if_then1408:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end1409:
	v628 = *libc.As[int32](lookahead)
	cmp1410 = 48 <= v628
	if cmp1410 {
		goto land_lhs_true1412
	} else {
		goto lor_lhs_false1415
	}

land_lhs_true1412:
	v629 = *libc.As[int32](lookahead)
	cmp1413 = v629 <= 57
	if cmp1413 {
		goto if_then1430
	} else {
		goto lor_lhs_false1415
	}

lor_lhs_false1415:
	v630 = *libc.As[int32](lookahead)
	cmp1416 = 65 <= v630
	if cmp1416 {
		goto land_lhs_true1418
	} else {
		goto lor_lhs_false1421
	}

land_lhs_true1418:
	v631 = *libc.As[int32](lookahead)
	cmp1419 = v631 <= 90
	if cmp1419 {
		goto if_then1430
	} else {
		goto lor_lhs_false1421
	}

lor_lhs_false1421:
	v632 = *libc.As[int32](lookahead)
	cmp1422 = v632 == 95
	if cmp1422 {
		goto if_then1430
	} else {
		goto lor_lhs_false1424
	}

lor_lhs_false1424:
	v633 = *libc.As[int32](lookahead)
	cmp1425 = 97 <= v633
	if cmp1425 {
		goto land_lhs_true1427
	} else {
		goto if_end1431
	}

land_lhs_true1427:
	v634 = *libc.As[int32](lookahead)
	cmp1428 = v634 <= 122
	if cmp1428 {
		goto if_then1430
	} else {
		goto if_end1431
	}

if_then1430:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1431:
	v635 = *libc.As[byte](result)
	loadedv1432 = (v635 & 1) != 0
	*libc.As[bool](retval) = loadedv1432
	goto _return

sw_bb1433:
	*libc.As[byte](result) = 1
	v636 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1434 = libc.Ptr(&libc.As[TSLexer](v636).F1)
	*libc.As[int16](result_symbol1434) = 13
	v637 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1435 = libc.Ptr(&libc.As[TSLexer](v637).F3)
	v638 = *libc.As[unsafe.Pointer](mark_end1435)
	v639 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v638)(v639)
	v640 = *libc.As[int32](lookahead)
	cmp1436 = v640 == 108
	if cmp1436 {
		goto if_then1438
	} else {
		goto if_end1439
	}

if_then1438:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end1439:
	v641 = *libc.As[int32](lookahead)
	cmp1440 = v641 == 114
	if cmp1440 {
		goto if_then1442
	} else {
		goto if_end1443
	}

if_then1442:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end1443:
	v642 = *libc.As[int32](lookahead)
	cmp1444 = 48 <= v642
	if cmp1444 {
		goto land_lhs_true1446
	} else {
		goto lor_lhs_false1449
	}

land_lhs_true1446:
	v643 = *libc.As[int32](lookahead)
	cmp1447 = v643 <= 57
	if cmp1447 {
		goto if_then1464
	} else {
		goto lor_lhs_false1449
	}

lor_lhs_false1449:
	v644 = *libc.As[int32](lookahead)
	cmp1450 = 65 <= v644
	if cmp1450 {
		goto land_lhs_true1452
	} else {
		goto lor_lhs_false1455
	}

land_lhs_true1452:
	v645 = *libc.As[int32](lookahead)
	cmp1453 = v645 <= 90
	if cmp1453 {
		goto if_then1464
	} else {
		goto lor_lhs_false1455
	}

lor_lhs_false1455:
	v646 = *libc.As[int32](lookahead)
	cmp1456 = v646 == 95
	if cmp1456 {
		goto if_then1464
	} else {
		goto lor_lhs_false1458
	}

lor_lhs_false1458:
	v647 = *libc.As[int32](lookahead)
	cmp1459 = 97 <= v647
	if cmp1459 {
		goto land_lhs_true1461
	} else {
		goto if_end1465
	}

land_lhs_true1461:
	v648 = *libc.As[int32](lookahead)
	cmp1462 = v648 <= 122
	if cmp1462 {
		goto if_then1464
	} else {
		goto if_end1465
	}

if_then1464:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1465:
	v649 = *libc.As[byte](result)
	loadedv1466 = (v649 & 1) != 0
	*libc.As[bool](retval) = loadedv1466
	goto _return

sw_bb1467:
	*libc.As[byte](result) = 1
	v650 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1468 = libc.Ptr(&libc.As[TSLexer](v650).F1)
	*libc.As[int16](result_symbol1468) = 13
	v651 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1469 = libc.Ptr(&libc.As[TSLexer](v651).F3)
	v652 = *libc.As[unsafe.Pointer](mark_end1469)
	v653 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v652)(v653)
	v654 = *libc.As[int32](lookahead)
	cmp1470 = v654 == 108
	if cmp1470 {
		goto if_then1472
	} else {
		goto if_end1473
	}

if_then1472:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end1473:
	v655 = *libc.As[int32](lookahead)
	cmp1474 = 48 <= v655
	if cmp1474 {
		goto land_lhs_true1476
	} else {
		goto lor_lhs_false1479
	}

land_lhs_true1476:
	v656 = *libc.As[int32](lookahead)
	cmp1477 = v656 <= 57
	if cmp1477 {
		goto if_then1494
	} else {
		goto lor_lhs_false1479
	}

lor_lhs_false1479:
	v657 = *libc.As[int32](lookahead)
	cmp1480 = 65 <= v657
	if cmp1480 {
		goto land_lhs_true1482
	} else {
		goto lor_lhs_false1485
	}

land_lhs_true1482:
	v658 = *libc.As[int32](lookahead)
	cmp1483 = v658 <= 90
	if cmp1483 {
		goto if_then1494
	} else {
		goto lor_lhs_false1485
	}

lor_lhs_false1485:
	v659 = *libc.As[int32](lookahead)
	cmp1486 = v659 == 95
	if cmp1486 {
		goto if_then1494
	} else {
		goto lor_lhs_false1488
	}

lor_lhs_false1488:
	v660 = *libc.As[int32](lookahead)
	cmp1489 = 97 <= v660
	if cmp1489 {
		goto land_lhs_true1491
	} else {
		goto if_end1495
	}

land_lhs_true1491:
	v661 = *libc.As[int32](lookahead)
	cmp1492 = v661 <= 122
	if cmp1492 {
		goto if_then1494
	} else {
		goto if_end1495
	}

if_then1494:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1495:
	v662 = *libc.As[byte](result)
	loadedv1496 = (v662 & 1) != 0
	*libc.As[bool](retval) = loadedv1496
	goto _return

sw_bb1497:
	*libc.As[byte](result) = 1
	v663 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1498 = libc.Ptr(&libc.As[TSLexer](v663).F1)
	*libc.As[int16](result_symbol1498) = 13
	v664 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1499 = libc.Ptr(&libc.As[TSLexer](v664).F3)
	v665 = *libc.As[unsafe.Pointer](mark_end1499)
	v666 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v665)(v666)
	v667 = *libc.As[int32](lookahead)
	cmp1500 = v667 == 108
	if cmp1500 {
		goto if_then1502
	} else {
		goto if_end1503
	}

if_then1502:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end1503:
	v668 = *libc.As[int32](lookahead)
	cmp1504 = 48 <= v668
	if cmp1504 {
		goto land_lhs_true1506
	} else {
		goto lor_lhs_false1509
	}

land_lhs_true1506:
	v669 = *libc.As[int32](lookahead)
	cmp1507 = v669 <= 57
	if cmp1507 {
		goto if_then1524
	} else {
		goto lor_lhs_false1509
	}

lor_lhs_false1509:
	v670 = *libc.As[int32](lookahead)
	cmp1510 = 65 <= v670
	if cmp1510 {
		goto land_lhs_true1512
	} else {
		goto lor_lhs_false1515
	}

land_lhs_true1512:
	v671 = *libc.As[int32](lookahead)
	cmp1513 = v671 <= 90
	if cmp1513 {
		goto if_then1524
	} else {
		goto lor_lhs_false1515
	}

lor_lhs_false1515:
	v672 = *libc.As[int32](lookahead)
	cmp1516 = v672 == 95
	if cmp1516 {
		goto if_then1524
	} else {
		goto lor_lhs_false1518
	}

lor_lhs_false1518:
	v673 = *libc.As[int32](lookahead)
	cmp1519 = 97 <= v673
	if cmp1519 {
		goto land_lhs_true1521
	} else {
		goto if_end1525
	}

land_lhs_true1521:
	v674 = *libc.As[int32](lookahead)
	cmp1522 = v674 <= 122
	if cmp1522 {
		goto if_then1524
	} else {
		goto if_end1525
	}

if_then1524:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1525:
	v675 = *libc.As[byte](result)
	loadedv1526 = (v675 & 1) != 0
	*libc.As[bool](retval) = loadedv1526
	goto _return

sw_bb1527:
	*libc.As[byte](result) = 1
	v676 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1528 = libc.Ptr(&libc.As[TSLexer](v676).F1)
	*libc.As[int16](result_symbol1528) = 13
	v677 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1529 = libc.Ptr(&libc.As[TSLexer](v677).F3)
	v678 = *libc.As[unsafe.Pointer](mark_end1529)
	v679 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v678)(v679)
	v680 = *libc.As[int32](lookahead)
	cmp1530 = v680 == 108
	if cmp1530 {
		goto if_then1532
	} else {
		goto if_end1533
	}

if_then1532:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end1533:
	v681 = *libc.As[int32](lookahead)
	cmp1534 = 48 <= v681
	if cmp1534 {
		goto land_lhs_true1536
	} else {
		goto lor_lhs_false1539
	}

land_lhs_true1536:
	v682 = *libc.As[int32](lookahead)
	cmp1537 = v682 <= 57
	if cmp1537 {
		goto if_then1554
	} else {
		goto lor_lhs_false1539
	}

lor_lhs_false1539:
	v683 = *libc.As[int32](lookahead)
	cmp1540 = 65 <= v683
	if cmp1540 {
		goto land_lhs_true1542
	} else {
		goto lor_lhs_false1545
	}

land_lhs_true1542:
	v684 = *libc.As[int32](lookahead)
	cmp1543 = v684 <= 90
	if cmp1543 {
		goto if_then1554
	} else {
		goto lor_lhs_false1545
	}

lor_lhs_false1545:
	v685 = *libc.As[int32](lookahead)
	cmp1546 = v685 == 95
	if cmp1546 {
		goto if_then1554
	} else {
		goto lor_lhs_false1548
	}

lor_lhs_false1548:
	v686 = *libc.As[int32](lookahead)
	cmp1549 = 97 <= v686
	if cmp1549 {
		goto land_lhs_true1551
	} else {
		goto if_end1555
	}

land_lhs_true1551:
	v687 = *libc.As[int32](lookahead)
	cmp1552 = v687 <= 122
	if cmp1552 {
		goto if_then1554
	} else {
		goto if_end1555
	}

if_then1554:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1555:
	v688 = *libc.As[byte](result)
	loadedv1556 = (v688 & 1) != 0
	*libc.As[bool](retval) = loadedv1556
	goto _return

sw_bb1557:
	*libc.As[byte](result) = 1
	v689 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1558 = libc.Ptr(&libc.As[TSLexer](v689).F1)
	*libc.As[int16](result_symbol1558) = 13
	v690 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1559 = libc.Ptr(&libc.As[TSLexer](v690).F3)
	v691 = *libc.As[unsafe.Pointer](mark_end1559)
	v692 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v691)(v692)
	v693 = *libc.As[int32](lookahead)
	cmp1560 = v693 == 108
	if cmp1560 {
		goto if_then1562
	} else {
		goto if_end1563
	}

if_then1562:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end1563:
	v694 = *libc.As[int32](lookahead)
	cmp1564 = 48 <= v694
	if cmp1564 {
		goto land_lhs_true1566
	} else {
		goto lor_lhs_false1569
	}

land_lhs_true1566:
	v695 = *libc.As[int32](lookahead)
	cmp1567 = v695 <= 57
	if cmp1567 {
		goto if_then1584
	} else {
		goto lor_lhs_false1569
	}

lor_lhs_false1569:
	v696 = *libc.As[int32](lookahead)
	cmp1570 = 65 <= v696
	if cmp1570 {
		goto land_lhs_true1572
	} else {
		goto lor_lhs_false1575
	}

land_lhs_true1572:
	v697 = *libc.As[int32](lookahead)
	cmp1573 = v697 <= 90
	if cmp1573 {
		goto if_then1584
	} else {
		goto lor_lhs_false1575
	}

lor_lhs_false1575:
	v698 = *libc.As[int32](lookahead)
	cmp1576 = v698 == 95
	if cmp1576 {
		goto if_then1584
	} else {
		goto lor_lhs_false1578
	}

lor_lhs_false1578:
	v699 = *libc.As[int32](lookahead)
	cmp1579 = 97 <= v699
	if cmp1579 {
		goto land_lhs_true1581
	} else {
		goto if_end1585
	}

land_lhs_true1581:
	v700 = *libc.As[int32](lookahead)
	cmp1582 = v700 <= 122
	if cmp1582 {
		goto if_then1584
	} else {
		goto if_end1585
	}

if_then1584:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1585:
	v701 = *libc.As[byte](result)
	loadedv1586 = (v701 & 1) != 0
	*libc.As[bool](retval) = loadedv1586
	goto _return

sw_bb1587:
	*libc.As[byte](result) = 1
	v702 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1588 = libc.Ptr(&libc.As[TSLexer](v702).F1)
	*libc.As[int16](result_symbol1588) = 13
	v703 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1589 = libc.Ptr(&libc.As[TSLexer](v703).F3)
	v704 = *libc.As[unsafe.Pointer](mark_end1589)
	v705 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v704)(v705)
	v706 = *libc.As[int32](lookahead)
	cmp1590 = v706 == 109
	if cmp1590 {
		goto if_then1592
	} else {
		goto if_end1593
	}

if_then1592:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end1593:
	v707 = *libc.As[int32](lookahead)
	cmp1594 = 48 <= v707
	if cmp1594 {
		goto land_lhs_true1596
	} else {
		goto lor_lhs_false1599
	}

land_lhs_true1596:
	v708 = *libc.As[int32](lookahead)
	cmp1597 = v708 <= 57
	if cmp1597 {
		goto if_then1614
	} else {
		goto lor_lhs_false1599
	}

lor_lhs_false1599:
	v709 = *libc.As[int32](lookahead)
	cmp1600 = 65 <= v709
	if cmp1600 {
		goto land_lhs_true1602
	} else {
		goto lor_lhs_false1605
	}

land_lhs_true1602:
	v710 = *libc.As[int32](lookahead)
	cmp1603 = v710 <= 90
	if cmp1603 {
		goto if_then1614
	} else {
		goto lor_lhs_false1605
	}

lor_lhs_false1605:
	v711 = *libc.As[int32](lookahead)
	cmp1606 = v711 == 95
	if cmp1606 {
		goto if_then1614
	} else {
		goto lor_lhs_false1608
	}

lor_lhs_false1608:
	v712 = *libc.As[int32](lookahead)
	cmp1609 = 97 <= v712
	if cmp1609 {
		goto land_lhs_true1611
	} else {
		goto if_end1615
	}

land_lhs_true1611:
	v713 = *libc.As[int32](lookahead)
	cmp1612 = v713 <= 122
	if cmp1612 {
		goto if_then1614
	} else {
		goto if_end1615
	}

if_then1614:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1615:
	v714 = *libc.As[byte](result)
	loadedv1616 = (v714 & 1) != 0
	*libc.As[bool](retval) = loadedv1616
	goto _return

sw_bb1617:
	*libc.As[byte](result) = 1
	v715 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1618 = libc.Ptr(&libc.As[TSLexer](v715).F1)
	*libc.As[int16](result_symbol1618) = 13
	v716 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1619 = libc.Ptr(&libc.As[TSLexer](v716).F3)
	v717 = *libc.As[unsafe.Pointer](mark_end1619)
	v718 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v717)(v718)
	v719 = *libc.As[int32](lookahead)
	cmp1620 = v719 == 110
	if cmp1620 {
		goto if_then1622
	} else {
		goto if_end1623
	}

if_then1622:
	*libc.As[int16](state_addr) = 90
	goto next_state

if_end1623:
	v720 = *libc.As[int32](lookahead)
	cmp1624 = 48 <= v720
	if cmp1624 {
		goto land_lhs_true1626
	} else {
		goto lor_lhs_false1629
	}

land_lhs_true1626:
	v721 = *libc.As[int32](lookahead)
	cmp1627 = v721 <= 57
	if cmp1627 {
		goto if_then1644
	} else {
		goto lor_lhs_false1629
	}

lor_lhs_false1629:
	v722 = *libc.As[int32](lookahead)
	cmp1630 = 65 <= v722
	if cmp1630 {
		goto land_lhs_true1632
	} else {
		goto lor_lhs_false1635
	}

land_lhs_true1632:
	v723 = *libc.As[int32](lookahead)
	cmp1633 = v723 <= 90
	if cmp1633 {
		goto if_then1644
	} else {
		goto lor_lhs_false1635
	}

lor_lhs_false1635:
	v724 = *libc.As[int32](lookahead)
	cmp1636 = v724 == 95
	if cmp1636 {
		goto if_then1644
	} else {
		goto lor_lhs_false1638
	}

lor_lhs_false1638:
	v725 = *libc.As[int32](lookahead)
	cmp1639 = 97 <= v725
	if cmp1639 {
		goto land_lhs_true1641
	} else {
		goto if_end1645
	}

land_lhs_true1641:
	v726 = *libc.As[int32](lookahead)
	cmp1642 = v726 <= 122
	if cmp1642 {
		goto if_then1644
	} else {
		goto if_end1645
	}

if_then1644:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1645:
	v727 = *libc.As[byte](result)
	loadedv1646 = (v727 & 1) != 0
	*libc.As[bool](retval) = loadedv1646
	goto _return

sw_bb1647:
	*libc.As[byte](result) = 1
	v728 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1648 = libc.Ptr(&libc.As[TSLexer](v728).F1)
	*libc.As[int16](result_symbol1648) = 13
	v729 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1649 = libc.Ptr(&libc.As[TSLexer](v729).F3)
	v730 = *libc.As[unsafe.Pointer](mark_end1649)
	v731 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v730)(v731)
	v732 = *libc.As[int32](lookahead)
	cmp1650 = v732 == 110
	if cmp1650 {
		goto if_then1652
	} else {
		goto if_end1653
	}

if_then1652:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end1653:
	v733 = *libc.As[int32](lookahead)
	cmp1654 = 48 <= v733
	if cmp1654 {
		goto land_lhs_true1656
	} else {
		goto lor_lhs_false1659
	}

land_lhs_true1656:
	v734 = *libc.As[int32](lookahead)
	cmp1657 = v734 <= 57
	if cmp1657 {
		goto if_then1674
	} else {
		goto lor_lhs_false1659
	}

lor_lhs_false1659:
	v735 = *libc.As[int32](lookahead)
	cmp1660 = 65 <= v735
	if cmp1660 {
		goto land_lhs_true1662
	} else {
		goto lor_lhs_false1665
	}

land_lhs_true1662:
	v736 = *libc.As[int32](lookahead)
	cmp1663 = v736 <= 90
	if cmp1663 {
		goto if_then1674
	} else {
		goto lor_lhs_false1665
	}

lor_lhs_false1665:
	v737 = *libc.As[int32](lookahead)
	cmp1666 = v737 == 95
	if cmp1666 {
		goto if_then1674
	} else {
		goto lor_lhs_false1668
	}

lor_lhs_false1668:
	v738 = *libc.As[int32](lookahead)
	cmp1669 = 97 <= v738
	if cmp1669 {
		goto land_lhs_true1671
	} else {
		goto if_end1675
	}

land_lhs_true1671:
	v739 = *libc.As[int32](lookahead)
	cmp1672 = v739 <= 122
	if cmp1672 {
		goto if_then1674
	} else {
		goto if_end1675
	}

if_then1674:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1675:
	v740 = *libc.As[byte](result)
	loadedv1676 = (v740 & 1) != 0
	*libc.As[bool](retval) = loadedv1676
	goto _return

sw_bb1677:
	*libc.As[byte](result) = 1
	v741 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1678 = libc.Ptr(&libc.As[TSLexer](v741).F1)
	*libc.As[int16](result_symbol1678) = 13
	v742 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1679 = libc.Ptr(&libc.As[TSLexer](v742).F3)
	v743 = *libc.As[unsafe.Pointer](mark_end1679)
	v744 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v743)(v744)
	v745 = *libc.As[int32](lookahead)
	cmp1680 = v745 == 110
	if cmp1680 {
		goto if_then1682
	} else {
		goto if_end1683
	}

if_then1682:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end1683:
	v746 = *libc.As[int32](lookahead)
	cmp1684 = 48 <= v746
	if cmp1684 {
		goto land_lhs_true1686
	} else {
		goto lor_lhs_false1689
	}

land_lhs_true1686:
	v747 = *libc.As[int32](lookahead)
	cmp1687 = v747 <= 57
	if cmp1687 {
		goto if_then1704
	} else {
		goto lor_lhs_false1689
	}

lor_lhs_false1689:
	v748 = *libc.As[int32](lookahead)
	cmp1690 = 65 <= v748
	if cmp1690 {
		goto land_lhs_true1692
	} else {
		goto lor_lhs_false1695
	}

land_lhs_true1692:
	v749 = *libc.As[int32](lookahead)
	cmp1693 = v749 <= 90
	if cmp1693 {
		goto if_then1704
	} else {
		goto lor_lhs_false1695
	}

lor_lhs_false1695:
	v750 = *libc.As[int32](lookahead)
	cmp1696 = v750 == 95
	if cmp1696 {
		goto if_then1704
	} else {
		goto lor_lhs_false1698
	}

lor_lhs_false1698:
	v751 = *libc.As[int32](lookahead)
	cmp1699 = 97 <= v751
	if cmp1699 {
		goto land_lhs_true1701
	} else {
		goto if_end1705
	}

land_lhs_true1701:
	v752 = *libc.As[int32](lookahead)
	cmp1702 = v752 <= 122
	if cmp1702 {
		goto if_then1704
	} else {
		goto if_end1705
	}

if_then1704:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1705:
	v753 = *libc.As[byte](result)
	loadedv1706 = (v753 & 1) != 0
	*libc.As[bool](retval) = loadedv1706
	goto _return

sw_bb1707:
	*libc.As[byte](result) = 1
	v754 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1708 = libc.Ptr(&libc.As[TSLexer](v754).F1)
	*libc.As[int16](result_symbol1708) = 13
	v755 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1709 = libc.Ptr(&libc.As[TSLexer](v755).F3)
	v756 = *libc.As[unsafe.Pointer](mark_end1709)
	v757 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v756)(v757)
	v758 = *libc.As[int32](lookahead)
	cmp1710 = v758 == 110
	if cmp1710 {
		goto if_then1712
	} else {
		goto if_end1713
	}

if_then1712:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end1713:
	v759 = *libc.As[int32](lookahead)
	cmp1714 = 48 <= v759
	if cmp1714 {
		goto land_lhs_true1716
	} else {
		goto lor_lhs_false1719
	}

land_lhs_true1716:
	v760 = *libc.As[int32](lookahead)
	cmp1717 = v760 <= 57
	if cmp1717 {
		goto if_then1734
	} else {
		goto lor_lhs_false1719
	}

lor_lhs_false1719:
	v761 = *libc.As[int32](lookahead)
	cmp1720 = 65 <= v761
	if cmp1720 {
		goto land_lhs_true1722
	} else {
		goto lor_lhs_false1725
	}

land_lhs_true1722:
	v762 = *libc.As[int32](lookahead)
	cmp1723 = v762 <= 90
	if cmp1723 {
		goto if_then1734
	} else {
		goto lor_lhs_false1725
	}

lor_lhs_false1725:
	v763 = *libc.As[int32](lookahead)
	cmp1726 = v763 == 95
	if cmp1726 {
		goto if_then1734
	} else {
		goto lor_lhs_false1728
	}

lor_lhs_false1728:
	v764 = *libc.As[int32](lookahead)
	cmp1729 = 97 <= v764
	if cmp1729 {
		goto land_lhs_true1731
	} else {
		goto if_end1735
	}

land_lhs_true1731:
	v765 = *libc.As[int32](lookahead)
	cmp1732 = v765 <= 122
	if cmp1732 {
		goto if_then1734
	} else {
		goto if_end1735
	}

if_then1734:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1735:
	v766 = *libc.As[byte](result)
	loadedv1736 = (v766 & 1) != 0
	*libc.As[bool](retval) = loadedv1736
	goto _return

sw_bb1737:
	*libc.As[byte](result) = 1
	v767 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1738 = libc.Ptr(&libc.As[TSLexer](v767).F1)
	*libc.As[int16](result_symbol1738) = 13
	v768 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1739 = libc.Ptr(&libc.As[TSLexer](v768).F3)
	v769 = *libc.As[unsafe.Pointer](mark_end1739)
	v770 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v769)(v770)
	v771 = *libc.As[int32](lookahead)
	cmp1740 = v771 == 110
	if cmp1740 {
		goto if_then1742
	} else {
		goto if_end1743
	}

if_then1742:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end1743:
	v772 = *libc.As[int32](lookahead)
	cmp1744 = 48 <= v772
	if cmp1744 {
		goto land_lhs_true1746
	} else {
		goto lor_lhs_false1749
	}

land_lhs_true1746:
	v773 = *libc.As[int32](lookahead)
	cmp1747 = v773 <= 57
	if cmp1747 {
		goto if_then1764
	} else {
		goto lor_lhs_false1749
	}

lor_lhs_false1749:
	v774 = *libc.As[int32](lookahead)
	cmp1750 = 65 <= v774
	if cmp1750 {
		goto land_lhs_true1752
	} else {
		goto lor_lhs_false1755
	}

land_lhs_true1752:
	v775 = *libc.As[int32](lookahead)
	cmp1753 = v775 <= 90
	if cmp1753 {
		goto if_then1764
	} else {
		goto lor_lhs_false1755
	}

lor_lhs_false1755:
	v776 = *libc.As[int32](lookahead)
	cmp1756 = v776 == 95
	if cmp1756 {
		goto if_then1764
	} else {
		goto lor_lhs_false1758
	}

lor_lhs_false1758:
	v777 = *libc.As[int32](lookahead)
	cmp1759 = 97 <= v777
	if cmp1759 {
		goto land_lhs_true1761
	} else {
		goto if_end1765
	}

land_lhs_true1761:
	v778 = *libc.As[int32](lookahead)
	cmp1762 = v778 <= 122
	if cmp1762 {
		goto if_then1764
	} else {
		goto if_end1765
	}

if_then1764:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1765:
	v779 = *libc.As[byte](result)
	loadedv1766 = (v779 & 1) != 0
	*libc.As[bool](retval) = loadedv1766
	goto _return

sw_bb1767:
	*libc.As[byte](result) = 1
	v780 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1768 = libc.Ptr(&libc.As[TSLexer](v780).F1)
	*libc.As[int16](result_symbol1768) = 13
	v781 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1769 = libc.Ptr(&libc.As[TSLexer](v781).F3)
	v782 = *libc.As[unsafe.Pointer](mark_end1769)
	v783 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v782)(v783)
	v784 = *libc.As[int32](lookahead)
	cmp1770 = v784 == 110
	if cmp1770 {
		goto if_then1772
	} else {
		goto if_end1773
	}

if_then1772:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end1773:
	v785 = *libc.As[int32](lookahead)
	cmp1774 = 48 <= v785
	if cmp1774 {
		goto land_lhs_true1776
	} else {
		goto lor_lhs_false1779
	}

land_lhs_true1776:
	v786 = *libc.As[int32](lookahead)
	cmp1777 = v786 <= 57
	if cmp1777 {
		goto if_then1794
	} else {
		goto lor_lhs_false1779
	}

lor_lhs_false1779:
	v787 = *libc.As[int32](lookahead)
	cmp1780 = 65 <= v787
	if cmp1780 {
		goto land_lhs_true1782
	} else {
		goto lor_lhs_false1785
	}

land_lhs_true1782:
	v788 = *libc.As[int32](lookahead)
	cmp1783 = v788 <= 90
	if cmp1783 {
		goto if_then1794
	} else {
		goto lor_lhs_false1785
	}

lor_lhs_false1785:
	v789 = *libc.As[int32](lookahead)
	cmp1786 = v789 == 95
	if cmp1786 {
		goto if_then1794
	} else {
		goto lor_lhs_false1788
	}

lor_lhs_false1788:
	v790 = *libc.As[int32](lookahead)
	cmp1789 = 97 <= v790
	if cmp1789 {
		goto land_lhs_true1791
	} else {
		goto if_end1795
	}

land_lhs_true1791:
	v791 = *libc.As[int32](lookahead)
	cmp1792 = v791 <= 122
	if cmp1792 {
		goto if_then1794
	} else {
		goto if_end1795
	}

if_then1794:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1795:
	v792 = *libc.As[byte](result)
	loadedv1796 = (v792 & 1) != 0
	*libc.As[bool](retval) = loadedv1796
	goto _return

sw_bb1797:
	*libc.As[byte](result) = 1
	v793 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1798 = libc.Ptr(&libc.As[TSLexer](v793).F1)
	*libc.As[int16](result_symbol1798) = 13
	v794 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1799 = libc.Ptr(&libc.As[TSLexer](v794).F3)
	v795 = *libc.As[unsafe.Pointer](mark_end1799)
	v796 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v795)(v796)
	v797 = *libc.As[int32](lookahead)
	cmp1800 = v797 == 110
	if cmp1800 {
		goto if_then1802
	} else {
		goto if_end1803
	}

if_then1802:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end1803:
	v798 = *libc.As[int32](lookahead)
	cmp1804 = 48 <= v798
	if cmp1804 {
		goto land_lhs_true1806
	} else {
		goto lor_lhs_false1809
	}

land_lhs_true1806:
	v799 = *libc.As[int32](lookahead)
	cmp1807 = v799 <= 57
	if cmp1807 {
		goto if_then1824
	} else {
		goto lor_lhs_false1809
	}

lor_lhs_false1809:
	v800 = *libc.As[int32](lookahead)
	cmp1810 = 65 <= v800
	if cmp1810 {
		goto land_lhs_true1812
	} else {
		goto lor_lhs_false1815
	}

land_lhs_true1812:
	v801 = *libc.As[int32](lookahead)
	cmp1813 = v801 <= 90
	if cmp1813 {
		goto if_then1824
	} else {
		goto lor_lhs_false1815
	}

lor_lhs_false1815:
	v802 = *libc.As[int32](lookahead)
	cmp1816 = v802 == 95
	if cmp1816 {
		goto if_then1824
	} else {
		goto lor_lhs_false1818
	}

lor_lhs_false1818:
	v803 = *libc.As[int32](lookahead)
	cmp1819 = 97 <= v803
	if cmp1819 {
		goto land_lhs_true1821
	} else {
		goto if_end1825
	}

land_lhs_true1821:
	v804 = *libc.As[int32](lookahead)
	cmp1822 = v804 <= 122
	if cmp1822 {
		goto if_then1824
	} else {
		goto if_end1825
	}

if_then1824:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1825:
	v805 = *libc.As[byte](result)
	loadedv1826 = (v805 & 1) != 0
	*libc.As[bool](retval) = loadedv1826
	goto _return

sw_bb1827:
	*libc.As[byte](result) = 1
	v806 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1828 = libc.Ptr(&libc.As[TSLexer](v806).F1)
	*libc.As[int16](result_symbol1828) = 13
	v807 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1829 = libc.Ptr(&libc.As[TSLexer](v807).F3)
	v808 = *libc.As[unsafe.Pointer](mark_end1829)
	v809 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v808)(v809)
	v810 = *libc.As[int32](lookahead)
	cmp1830 = v810 == 110
	if cmp1830 {
		goto if_then1832
	} else {
		goto if_end1833
	}

if_then1832:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end1833:
	v811 = *libc.As[int32](lookahead)
	cmp1834 = 48 <= v811
	if cmp1834 {
		goto land_lhs_true1836
	} else {
		goto lor_lhs_false1839
	}

land_lhs_true1836:
	v812 = *libc.As[int32](lookahead)
	cmp1837 = v812 <= 57
	if cmp1837 {
		goto if_then1854
	} else {
		goto lor_lhs_false1839
	}

lor_lhs_false1839:
	v813 = *libc.As[int32](lookahead)
	cmp1840 = 65 <= v813
	if cmp1840 {
		goto land_lhs_true1842
	} else {
		goto lor_lhs_false1845
	}

land_lhs_true1842:
	v814 = *libc.As[int32](lookahead)
	cmp1843 = v814 <= 90
	if cmp1843 {
		goto if_then1854
	} else {
		goto lor_lhs_false1845
	}

lor_lhs_false1845:
	v815 = *libc.As[int32](lookahead)
	cmp1846 = v815 == 95
	if cmp1846 {
		goto if_then1854
	} else {
		goto lor_lhs_false1848
	}

lor_lhs_false1848:
	v816 = *libc.As[int32](lookahead)
	cmp1849 = 97 <= v816
	if cmp1849 {
		goto land_lhs_true1851
	} else {
		goto if_end1855
	}

land_lhs_true1851:
	v817 = *libc.As[int32](lookahead)
	cmp1852 = v817 <= 122
	if cmp1852 {
		goto if_then1854
	} else {
		goto if_end1855
	}

if_then1854:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1855:
	v818 = *libc.As[byte](result)
	loadedv1856 = (v818 & 1) != 0
	*libc.As[bool](retval) = loadedv1856
	goto _return

sw_bb1857:
	*libc.As[byte](result) = 1
	v819 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1858 = libc.Ptr(&libc.As[TSLexer](v819).F1)
	*libc.As[int16](result_symbol1858) = 13
	v820 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1859 = libc.Ptr(&libc.As[TSLexer](v820).F3)
	v821 = *libc.As[unsafe.Pointer](mark_end1859)
	v822 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v821)(v822)
	v823 = *libc.As[int32](lookahead)
	cmp1860 = v823 == 110
	if cmp1860 {
		goto if_then1862
	} else {
		goto if_end1863
	}

if_then1862:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end1863:
	v824 = *libc.As[int32](lookahead)
	cmp1864 = 48 <= v824
	if cmp1864 {
		goto land_lhs_true1866
	} else {
		goto lor_lhs_false1869
	}

land_lhs_true1866:
	v825 = *libc.As[int32](lookahead)
	cmp1867 = v825 <= 57
	if cmp1867 {
		goto if_then1884
	} else {
		goto lor_lhs_false1869
	}

lor_lhs_false1869:
	v826 = *libc.As[int32](lookahead)
	cmp1870 = 65 <= v826
	if cmp1870 {
		goto land_lhs_true1872
	} else {
		goto lor_lhs_false1875
	}

land_lhs_true1872:
	v827 = *libc.As[int32](lookahead)
	cmp1873 = v827 <= 90
	if cmp1873 {
		goto if_then1884
	} else {
		goto lor_lhs_false1875
	}

lor_lhs_false1875:
	v828 = *libc.As[int32](lookahead)
	cmp1876 = v828 == 95
	if cmp1876 {
		goto if_then1884
	} else {
		goto lor_lhs_false1878
	}

lor_lhs_false1878:
	v829 = *libc.As[int32](lookahead)
	cmp1879 = 97 <= v829
	if cmp1879 {
		goto land_lhs_true1881
	} else {
		goto if_end1885
	}

land_lhs_true1881:
	v830 = *libc.As[int32](lookahead)
	cmp1882 = v830 <= 122
	if cmp1882 {
		goto if_then1884
	} else {
		goto if_end1885
	}

if_then1884:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1885:
	v831 = *libc.As[byte](result)
	loadedv1886 = (v831 & 1) != 0
	*libc.As[bool](retval) = loadedv1886
	goto _return

sw_bb1887:
	*libc.As[byte](result) = 1
	v832 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1888 = libc.Ptr(&libc.As[TSLexer](v832).F1)
	*libc.As[int16](result_symbol1888) = 13
	v833 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1889 = libc.Ptr(&libc.As[TSLexer](v833).F3)
	v834 = *libc.As[unsafe.Pointer](mark_end1889)
	v835 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v834)(v835)
	v836 = *libc.As[int32](lookahead)
	cmp1890 = v836 == 110
	if cmp1890 {
		goto if_then1892
	} else {
		goto if_end1893
	}

if_then1892:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end1893:
	v837 = *libc.As[int32](lookahead)
	cmp1894 = 48 <= v837
	if cmp1894 {
		goto land_lhs_true1896
	} else {
		goto lor_lhs_false1899
	}

land_lhs_true1896:
	v838 = *libc.As[int32](lookahead)
	cmp1897 = v838 <= 57
	if cmp1897 {
		goto if_then1914
	} else {
		goto lor_lhs_false1899
	}

lor_lhs_false1899:
	v839 = *libc.As[int32](lookahead)
	cmp1900 = 65 <= v839
	if cmp1900 {
		goto land_lhs_true1902
	} else {
		goto lor_lhs_false1905
	}

land_lhs_true1902:
	v840 = *libc.As[int32](lookahead)
	cmp1903 = v840 <= 90
	if cmp1903 {
		goto if_then1914
	} else {
		goto lor_lhs_false1905
	}

lor_lhs_false1905:
	v841 = *libc.As[int32](lookahead)
	cmp1906 = v841 == 95
	if cmp1906 {
		goto if_then1914
	} else {
		goto lor_lhs_false1908
	}

lor_lhs_false1908:
	v842 = *libc.As[int32](lookahead)
	cmp1909 = 97 <= v842
	if cmp1909 {
		goto land_lhs_true1911
	} else {
		goto if_end1915
	}

land_lhs_true1911:
	v843 = *libc.As[int32](lookahead)
	cmp1912 = v843 <= 122
	if cmp1912 {
		goto if_then1914
	} else {
		goto if_end1915
	}

if_then1914:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1915:
	v844 = *libc.As[byte](result)
	loadedv1916 = (v844 & 1) != 0
	*libc.As[bool](retval) = loadedv1916
	goto _return

sw_bb1917:
	*libc.As[byte](result) = 1
	v845 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1918 = libc.Ptr(&libc.As[TSLexer](v845).F1)
	*libc.As[int16](result_symbol1918) = 13
	v846 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1919 = libc.Ptr(&libc.As[TSLexer](v846).F3)
	v847 = *libc.As[unsafe.Pointer](mark_end1919)
	v848 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v847)(v848)
	v849 = *libc.As[int32](lookahead)
	cmp1920 = v849 == 110
	if cmp1920 {
		goto if_then1922
	} else {
		goto if_end1923
	}

if_then1922:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end1923:
	v850 = *libc.As[int32](lookahead)
	cmp1924 = 48 <= v850
	if cmp1924 {
		goto land_lhs_true1926
	} else {
		goto lor_lhs_false1929
	}

land_lhs_true1926:
	v851 = *libc.As[int32](lookahead)
	cmp1927 = v851 <= 57
	if cmp1927 {
		goto if_then1944
	} else {
		goto lor_lhs_false1929
	}

lor_lhs_false1929:
	v852 = *libc.As[int32](lookahead)
	cmp1930 = 65 <= v852
	if cmp1930 {
		goto land_lhs_true1932
	} else {
		goto lor_lhs_false1935
	}

land_lhs_true1932:
	v853 = *libc.As[int32](lookahead)
	cmp1933 = v853 <= 90
	if cmp1933 {
		goto if_then1944
	} else {
		goto lor_lhs_false1935
	}

lor_lhs_false1935:
	v854 = *libc.As[int32](lookahead)
	cmp1936 = v854 == 95
	if cmp1936 {
		goto if_then1944
	} else {
		goto lor_lhs_false1938
	}

lor_lhs_false1938:
	v855 = *libc.As[int32](lookahead)
	cmp1939 = 97 <= v855
	if cmp1939 {
		goto land_lhs_true1941
	} else {
		goto if_end1945
	}

land_lhs_true1941:
	v856 = *libc.As[int32](lookahead)
	cmp1942 = v856 <= 122
	if cmp1942 {
		goto if_then1944
	} else {
		goto if_end1945
	}

if_then1944:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1945:
	v857 = *libc.As[byte](result)
	loadedv1946 = (v857 & 1) != 0
	*libc.As[bool](retval) = loadedv1946
	goto _return

sw_bb1947:
	*libc.As[byte](result) = 1
	v858 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1948 = libc.Ptr(&libc.As[TSLexer](v858).F1)
	*libc.As[int16](result_symbol1948) = 13
	v859 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1949 = libc.Ptr(&libc.As[TSLexer](v859).F3)
	v860 = *libc.As[unsafe.Pointer](mark_end1949)
	v861 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v860)(v861)
	v862 = *libc.As[int32](lookahead)
	cmp1950 = v862 == 111
	if cmp1950 {
		goto if_then1952
	} else {
		goto if_end1953
	}

if_then1952:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end1953:
	v863 = *libc.As[int32](lookahead)
	cmp1954 = 48 <= v863
	if cmp1954 {
		goto land_lhs_true1956
	} else {
		goto lor_lhs_false1959
	}

land_lhs_true1956:
	v864 = *libc.As[int32](lookahead)
	cmp1957 = v864 <= 57
	if cmp1957 {
		goto if_then1974
	} else {
		goto lor_lhs_false1959
	}

lor_lhs_false1959:
	v865 = *libc.As[int32](lookahead)
	cmp1960 = 65 <= v865
	if cmp1960 {
		goto land_lhs_true1962
	} else {
		goto lor_lhs_false1965
	}

land_lhs_true1962:
	v866 = *libc.As[int32](lookahead)
	cmp1963 = v866 <= 90
	if cmp1963 {
		goto if_then1974
	} else {
		goto lor_lhs_false1965
	}

lor_lhs_false1965:
	v867 = *libc.As[int32](lookahead)
	cmp1966 = v867 == 95
	if cmp1966 {
		goto if_then1974
	} else {
		goto lor_lhs_false1968
	}

lor_lhs_false1968:
	v868 = *libc.As[int32](lookahead)
	cmp1969 = 97 <= v868
	if cmp1969 {
		goto land_lhs_true1971
	} else {
		goto if_end1975
	}

land_lhs_true1971:
	v869 = *libc.As[int32](lookahead)
	cmp1972 = v869 <= 122
	if cmp1972 {
		goto if_then1974
	} else {
		goto if_end1975
	}

if_then1974:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end1975:
	v870 = *libc.As[byte](result)
	loadedv1976 = (v870 & 1) != 0
	*libc.As[bool](retval) = loadedv1976
	goto _return

sw_bb1977:
	*libc.As[byte](result) = 1
	v871 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1978 = libc.Ptr(&libc.As[TSLexer](v871).F1)
	*libc.As[int16](result_symbol1978) = 13
	v872 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1979 = libc.Ptr(&libc.As[TSLexer](v872).F3)
	v873 = *libc.As[unsafe.Pointer](mark_end1979)
	v874 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v873)(v874)
	v875 = *libc.As[int32](lookahead)
	cmp1980 = v875 == 111
	if cmp1980 {
		goto if_then1982
	} else {
		goto if_end1983
	}

if_then1982:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end1983:
	v876 = *libc.As[int32](lookahead)
	cmp1984 = 48 <= v876
	if cmp1984 {
		goto land_lhs_true1986
	} else {
		goto lor_lhs_false1989
	}

land_lhs_true1986:
	v877 = *libc.As[int32](lookahead)
	cmp1987 = v877 <= 57
	if cmp1987 {
		goto if_then2004
	} else {
		goto lor_lhs_false1989
	}

lor_lhs_false1989:
	v878 = *libc.As[int32](lookahead)
	cmp1990 = 65 <= v878
	if cmp1990 {
		goto land_lhs_true1992
	} else {
		goto lor_lhs_false1995
	}

land_lhs_true1992:
	v879 = *libc.As[int32](lookahead)
	cmp1993 = v879 <= 90
	if cmp1993 {
		goto if_then2004
	} else {
		goto lor_lhs_false1995
	}

lor_lhs_false1995:
	v880 = *libc.As[int32](lookahead)
	cmp1996 = v880 == 95
	if cmp1996 {
		goto if_then2004
	} else {
		goto lor_lhs_false1998
	}

lor_lhs_false1998:
	v881 = *libc.As[int32](lookahead)
	cmp1999 = 97 <= v881
	if cmp1999 {
		goto land_lhs_true2001
	} else {
		goto if_end2005
	}

land_lhs_true2001:
	v882 = *libc.As[int32](lookahead)
	cmp2002 = v882 <= 122
	if cmp2002 {
		goto if_then2004
	} else {
		goto if_end2005
	}

if_then2004:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2005:
	v883 = *libc.As[byte](result)
	loadedv2006 = (v883 & 1) != 0
	*libc.As[bool](retval) = loadedv2006
	goto _return

sw_bb2007:
	*libc.As[byte](result) = 1
	v884 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2008 = libc.Ptr(&libc.As[TSLexer](v884).F1)
	*libc.As[int16](result_symbol2008) = 13
	v885 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2009 = libc.Ptr(&libc.As[TSLexer](v885).F3)
	v886 = *libc.As[unsafe.Pointer](mark_end2009)
	v887 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v886)(v887)
	v888 = *libc.As[int32](lookahead)
	cmp2010 = v888 == 111
	if cmp2010 {
		goto if_then2012
	} else {
		goto if_end2013
	}

if_then2012:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end2013:
	v889 = *libc.As[int32](lookahead)
	cmp2014 = 48 <= v889
	if cmp2014 {
		goto land_lhs_true2016
	} else {
		goto lor_lhs_false2019
	}

land_lhs_true2016:
	v890 = *libc.As[int32](lookahead)
	cmp2017 = v890 <= 57
	if cmp2017 {
		goto if_then2034
	} else {
		goto lor_lhs_false2019
	}

lor_lhs_false2019:
	v891 = *libc.As[int32](lookahead)
	cmp2020 = 65 <= v891
	if cmp2020 {
		goto land_lhs_true2022
	} else {
		goto lor_lhs_false2025
	}

land_lhs_true2022:
	v892 = *libc.As[int32](lookahead)
	cmp2023 = v892 <= 90
	if cmp2023 {
		goto if_then2034
	} else {
		goto lor_lhs_false2025
	}

lor_lhs_false2025:
	v893 = *libc.As[int32](lookahead)
	cmp2026 = v893 == 95
	if cmp2026 {
		goto if_then2034
	} else {
		goto lor_lhs_false2028
	}

lor_lhs_false2028:
	v894 = *libc.As[int32](lookahead)
	cmp2029 = 97 <= v894
	if cmp2029 {
		goto land_lhs_true2031
	} else {
		goto if_end2035
	}

land_lhs_true2031:
	v895 = *libc.As[int32](lookahead)
	cmp2032 = v895 <= 122
	if cmp2032 {
		goto if_then2034
	} else {
		goto if_end2035
	}

if_then2034:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2035:
	v896 = *libc.As[byte](result)
	loadedv2036 = (v896 & 1) != 0
	*libc.As[bool](retval) = loadedv2036
	goto _return

sw_bb2037:
	*libc.As[byte](result) = 1
	v897 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2038 = libc.Ptr(&libc.As[TSLexer](v897).F1)
	*libc.As[int16](result_symbol2038) = 13
	v898 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2039 = libc.Ptr(&libc.As[TSLexer](v898).F3)
	v899 = *libc.As[unsafe.Pointer](mark_end2039)
	v900 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v899)(v900)
	v901 = *libc.As[int32](lookahead)
	cmp2040 = v901 == 111
	if cmp2040 {
		goto if_then2042
	} else {
		goto if_end2043
	}

if_then2042:
	*libc.As[int16](state_addr) = 66
	goto next_state

if_end2043:
	v902 = *libc.As[int32](lookahead)
	cmp2044 = 48 <= v902
	if cmp2044 {
		goto land_lhs_true2046
	} else {
		goto lor_lhs_false2049
	}

land_lhs_true2046:
	v903 = *libc.As[int32](lookahead)
	cmp2047 = v903 <= 57
	if cmp2047 {
		goto if_then2064
	} else {
		goto lor_lhs_false2049
	}

lor_lhs_false2049:
	v904 = *libc.As[int32](lookahead)
	cmp2050 = 65 <= v904
	if cmp2050 {
		goto land_lhs_true2052
	} else {
		goto lor_lhs_false2055
	}

land_lhs_true2052:
	v905 = *libc.As[int32](lookahead)
	cmp2053 = v905 <= 90
	if cmp2053 {
		goto if_then2064
	} else {
		goto lor_lhs_false2055
	}

lor_lhs_false2055:
	v906 = *libc.As[int32](lookahead)
	cmp2056 = v906 == 95
	if cmp2056 {
		goto if_then2064
	} else {
		goto lor_lhs_false2058
	}

lor_lhs_false2058:
	v907 = *libc.As[int32](lookahead)
	cmp2059 = 97 <= v907
	if cmp2059 {
		goto land_lhs_true2061
	} else {
		goto if_end2065
	}

land_lhs_true2061:
	v908 = *libc.As[int32](lookahead)
	cmp2062 = v908 <= 122
	if cmp2062 {
		goto if_then2064
	} else {
		goto if_end2065
	}

if_then2064:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2065:
	v909 = *libc.As[byte](result)
	loadedv2066 = (v909 & 1) != 0
	*libc.As[bool](retval) = loadedv2066
	goto _return

sw_bb2067:
	*libc.As[byte](result) = 1
	v910 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2068 = libc.Ptr(&libc.As[TSLexer](v910).F1)
	*libc.As[int16](result_symbol2068) = 13
	v911 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2069 = libc.Ptr(&libc.As[TSLexer](v911).F3)
	v912 = *libc.As[unsafe.Pointer](mark_end2069)
	v913 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v912)(v913)
	v914 = *libc.As[int32](lookahead)
	cmp2070 = v914 == 111
	if cmp2070 {
		goto if_then2072
	} else {
		goto if_end2073
	}

if_then2072:
	*libc.As[int16](state_addr) = 60
	goto next_state

if_end2073:
	v915 = *libc.As[int32](lookahead)
	cmp2074 = 48 <= v915
	if cmp2074 {
		goto land_lhs_true2076
	} else {
		goto lor_lhs_false2079
	}

land_lhs_true2076:
	v916 = *libc.As[int32](lookahead)
	cmp2077 = v916 <= 57
	if cmp2077 {
		goto if_then2094
	} else {
		goto lor_lhs_false2079
	}

lor_lhs_false2079:
	v917 = *libc.As[int32](lookahead)
	cmp2080 = 65 <= v917
	if cmp2080 {
		goto land_lhs_true2082
	} else {
		goto lor_lhs_false2085
	}

land_lhs_true2082:
	v918 = *libc.As[int32](lookahead)
	cmp2083 = v918 <= 90
	if cmp2083 {
		goto if_then2094
	} else {
		goto lor_lhs_false2085
	}

lor_lhs_false2085:
	v919 = *libc.As[int32](lookahead)
	cmp2086 = v919 == 95
	if cmp2086 {
		goto if_then2094
	} else {
		goto lor_lhs_false2088
	}

lor_lhs_false2088:
	v920 = *libc.As[int32](lookahead)
	cmp2089 = 97 <= v920
	if cmp2089 {
		goto land_lhs_true2091
	} else {
		goto if_end2095
	}

land_lhs_true2091:
	v921 = *libc.As[int32](lookahead)
	cmp2092 = v921 <= 122
	if cmp2092 {
		goto if_then2094
	} else {
		goto if_end2095
	}

if_then2094:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2095:
	v922 = *libc.As[byte](result)
	loadedv2096 = (v922 & 1) != 0
	*libc.As[bool](retval) = loadedv2096
	goto _return

sw_bb2097:
	*libc.As[byte](result) = 1
	v923 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2098 = libc.Ptr(&libc.As[TSLexer](v923).F1)
	*libc.As[int16](result_symbol2098) = 13
	v924 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2099 = libc.Ptr(&libc.As[TSLexer](v924).F3)
	v925 = *libc.As[unsafe.Pointer](mark_end2099)
	v926 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v925)(v926)
	v927 = *libc.As[int32](lookahead)
	cmp2100 = v927 == 112
	if cmp2100 {
		goto if_then2102
	} else {
		goto if_end2103
	}

if_then2102:
	*libc.As[int16](state_addr) = 89
	goto next_state

if_end2103:
	v928 = *libc.As[int32](lookahead)
	cmp2104 = 48 <= v928
	if cmp2104 {
		goto land_lhs_true2106
	} else {
		goto lor_lhs_false2109
	}

land_lhs_true2106:
	v929 = *libc.As[int32](lookahead)
	cmp2107 = v929 <= 57
	if cmp2107 {
		goto if_then2124
	} else {
		goto lor_lhs_false2109
	}

lor_lhs_false2109:
	v930 = *libc.As[int32](lookahead)
	cmp2110 = 65 <= v930
	if cmp2110 {
		goto land_lhs_true2112
	} else {
		goto lor_lhs_false2115
	}

land_lhs_true2112:
	v931 = *libc.As[int32](lookahead)
	cmp2113 = v931 <= 90
	if cmp2113 {
		goto if_then2124
	} else {
		goto lor_lhs_false2115
	}

lor_lhs_false2115:
	v932 = *libc.As[int32](lookahead)
	cmp2116 = v932 == 95
	if cmp2116 {
		goto if_then2124
	} else {
		goto lor_lhs_false2118
	}

lor_lhs_false2118:
	v933 = *libc.As[int32](lookahead)
	cmp2119 = 97 <= v933
	if cmp2119 {
		goto land_lhs_true2121
	} else {
		goto if_end2125
	}

land_lhs_true2121:
	v934 = *libc.As[int32](lookahead)
	cmp2122 = v934 <= 122
	if cmp2122 {
		goto if_then2124
	} else {
		goto if_end2125
	}

if_then2124:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2125:
	v935 = *libc.As[byte](result)
	loadedv2126 = (v935 & 1) != 0
	*libc.As[bool](retval) = loadedv2126
	goto _return

sw_bb2127:
	*libc.As[byte](result) = 1
	v936 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2128 = libc.Ptr(&libc.As[TSLexer](v936).F1)
	*libc.As[int16](result_symbol2128) = 13
	v937 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2129 = libc.Ptr(&libc.As[TSLexer](v937).F3)
	v938 = *libc.As[unsafe.Pointer](mark_end2129)
	v939 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v938)(v939)
	v940 = *libc.As[int32](lookahead)
	cmp2130 = v940 == 112
	if cmp2130 {
		goto if_then2132
	} else {
		goto if_end2133
	}

if_then2132:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end2133:
	v941 = *libc.As[int32](lookahead)
	cmp2134 = v941 == 115
	if cmp2134 {
		goto if_then2136
	} else {
		goto if_end2137
	}

if_then2136:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end2137:
	v942 = *libc.As[int32](lookahead)
	cmp2138 = 48 <= v942
	if cmp2138 {
		goto land_lhs_true2140
	} else {
		goto lor_lhs_false2143
	}

land_lhs_true2140:
	v943 = *libc.As[int32](lookahead)
	cmp2141 = v943 <= 57
	if cmp2141 {
		goto if_then2158
	} else {
		goto lor_lhs_false2143
	}

lor_lhs_false2143:
	v944 = *libc.As[int32](lookahead)
	cmp2144 = 65 <= v944
	if cmp2144 {
		goto land_lhs_true2146
	} else {
		goto lor_lhs_false2149
	}

land_lhs_true2146:
	v945 = *libc.As[int32](lookahead)
	cmp2147 = v945 <= 90
	if cmp2147 {
		goto if_then2158
	} else {
		goto lor_lhs_false2149
	}

lor_lhs_false2149:
	v946 = *libc.As[int32](lookahead)
	cmp2150 = v946 == 95
	if cmp2150 {
		goto if_then2158
	} else {
		goto lor_lhs_false2152
	}

lor_lhs_false2152:
	v947 = *libc.As[int32](lookahead)
	cmp2153 = 97 <= v947
	if cmp2153 {
		goto land_lhs_true2155
	} else {
		goto if_end2159
	}

land_lhs_true2155:
	v948 = *libc.As[int32](lookahead)
	cmp2156 = v948 <= 122
	if cmp2156 {
		goto if_then2158
	} else {
		goto if_end2159
	}

if_then2158:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2159:
	v949 = *libc.As[byte](result)
	loadedv2160 = (v949 & 1) != 0
	*libc.As[bool](retval) = loadedv2160
	goto _return

sw_bb2161:
	*libc.As[byte](result) = 1
	v950 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2162 = libc.Ptr(&libc.As[TSLexer](v950).F1)
	*libc.As[int16](result_symbol2162) = 13
	v951 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2163 = libc.Ptr(&libc.As[TSLexer](v951).F3)
	v952 = *libc.As[unsafe.Pointer](mark_end2163)
	v953 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v952)(v953)
	v954 = *libc.As[int32](lookahead)
	cmp2164 = v954 == 112
	if cmp2164 {
		goto if_then2166
	} else {
		goto if_end2167
	}

if_then2166:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end2167:
	v955 = *libc.As[int32](lookahead)
	cmp2168 = 48 <= v955
	if cmp2168 {
		goto land_lhs_true2170
	} else {
		goto lor_lhs_false2173
	}

land_lhs_true2170:
	v956 = *libc.As[int32](lookahead)
	cmp2171 = v956 <= 57
	if cmp2171 {
		goto if_then2188
	} else {
		goto lor_lhs_false2173
	}

lor_lhs_false2173:
	v957 = *libc.As[int32](lookahead)
	cmp2174 = 65 <= v957
	if cmp2174 {
		goto land_lhs_true2176
	} else {
		goto lor_lhs_false2179
	}

land_lhs_true2176:
	v958 = *libc.As[int32](lookahead)
	cmp2177 = v958 <= 90
	if cmp2177 {
		goto if_then2188
	} else {
		goto lor_lhs_false2179
	}

lor_lhs_false2179:
	v959 = *libc.As[int32](lookahead)
	cmp2180 = v959 == 95
	if cmp2180 {
		goto if_then2188
	} else {
		goto lor_lhs_false2182
	}

lor_lhs_false2182:
	v960 = *libc.As[int32](lookahead)
	cmp2183 = 97 <= v960
	if cmp2183 {
		goto land_lhs_true2185
	} else {
		goto if_end2189
	}

land_lhs_true2185:
	v961 = *libc.As[int32](lookahead)
	cmp2186 = v961 <= 122
	if cmp2186 {
		goto if_then2188
	} else {
		goto if_end2189
	}

if_then2188:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2189:
	v962 = *libc.As[byte](result)
	loadedv2190 = (v962 & 1) != 0
	*libc.As[bool](retval) = loadedv2190
	goto _return

sw_bb2191:
	*libc.As[byte](result) = 1
	v963 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2192 = libc.Ptr(&libc.As[TSLexer](v963).F1)
	*libc.As[int16](result_symbol2192) = 13
	v964 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2193 = libc.Ptr(&libc.As[TSLexer](v964).F3)
	v965 = *libc.As[unsafe.Pointer](mark_end2193)
	v966 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v965)(v966)
	v967 = *libc.As[int32](lookahead)
	cmp2194 = v967 == 112
	if cmp2194 {
		goto if_then2196
	} else {
		goto if_end2197
	}

if_then2196:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end2197:
	v968 = *libc.As[int32](lookahead)
	cmp2198 = 48 <= v968
	if cmp2198 {
		goto land_lhs_true2200
	} else {
		goto lor_lhs_false2203
	}

land_lhs_true2200:
	v969 = *libc.As[int32](lookahead)
	cmp2201 = v969 <= 57
	if cmp2201 {
		goto if_then2218
	} else {
		goto lor_lhs_false2203
	}

lor_lhs_false2203:
	v970 = *libc.As[int32](lookahead)
	cmp2204 = 65 <= v970
	if cmp2204 {
		goto land_lhs_true2206
	} else {
		goto lor_lhs_false2209
	}

land_lhs_true2206:
	v971 = *libc.As[int32](lookahead)
	cmp2207 = v971 <= 90
	if cmp2207 {
		goto if_then2218
	} else {
		goto lor_lhs_false2209
	}

lor_lhs_false2209:
	v972 = *libc.As[int32](lookahead)
	cmp2210 = v972 == 95
	if cmp2210 {
		goto if_then2218
	} else {
		goto lor_lhs_false2212
	}

lor_lhs_false2212:
	v973 = *libc.As[int32](lookahead)
	cmp2213 = 97 <= v973
	if cmp2213 {
		goto land_lhs_true2215
	} else {
		goto if_end2219
	}

land_lhs_true2215:
	v974 = *libc.As[int32](lookahead)
	cmp2216 = v974 <= 122
	if cmp2216 {
		goto if_then2218
	} else {
		goto if_end2219
	}

if_then2218:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2219:
	v975 = *libc.As[byte](result)
	loadedv2220 = (v975 & 1) != 0
	*libc.As[bool](retval) = loadedv2220
	goto _return

sw_bb2221:
	*libc.As[byte](result) = 1
	v976 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2222 = libc.Ptr(&libc.As[TSLexer](v976).F1)
	*libc.As[int16](result_symbol2222) = 13
	v977 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2223 = libc.Ptr(&libc.As[TSLexer](v977).F3)
	v978 = *libc.As[unsafe.Pointer](mark_end2223)
	v979 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v978)(v979)
	v980 = *libc.As[int32](lookahead)
	cmp2224 = v980 == 112
	if cmp2224 {
		goto if_then2226
	} else {
		goto if_end2227
	}

if_then2226:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end2227:
	v981 = *libc.As[int32](lookahead)
	cmp2228 = 48 <= v981
	if cmp2228 {
		goto land_lhs_true2230
	} else {
		goto lor_lhs_false2233
	}

land_lhs_true2230:
	v982 = *libc.As[int32](lookahead)
	cmp2231 = v982 <= 57
	if cmp2231 {
		goto if_then2248
	} else {
		goto lor_lhs_false2233
	}

lor_lhs_false2233:
	v983 = *libc.As[int32](lookahead)
	cmp2234 = 65 <= v983
	if cmp2234 {
		goto land_lhs_true2236
	} else {
		goto lor_lhs_false2239
	}

land_lhs_true2236:
	v984 = *libc.As[int32](lookahead)
	cmp2237 = v984 <= 90
	if cmp2237 {
		goto if_then2248
	} else {
		goto lor_lhs_false2239
	}

lor_lhs_false2239:
	v985 = *libc.As[int32](lookahead)
	cmp2240 = v985 == 95
	if cmp2240 {
		goto if_then2248
	} else {
		goto lor_lhs_false2242
	}

lor_lhs_false2242:
	v986 = *libc.As[int32](lookahead)
	cmp2243 = 97 <= v986
	if cmp2243 {
		goto land_lhs_true2245
	} else {
		goto if_end2249
	}

land_lhs_true2245:
	v987 = *libc.As[int32](lookahead)
	cmp2246 = v987 <= 122
	if cmp2246 {
		goto if_then2248
	} else {
		goto if_end2249
	}

if_then2248:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2249:
	v988 = *libc.As[byte](result)
	loadedv2250 = (v988 & 1) != 0
	*libc.As[bool](retval) = loadedv2250
	goto _return

sw_bb2251:
	*libc.As[byte](result) = 1
	v989 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2252 = libc.Ptr(&libc.As[TSLexer](v989).F1)
	*libc.As[int16](result_symbol2252) = 13
	v990 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2253 = libc.Ptr(&libc.As[TSLexer](v990).F3)
	v991 = *libc.As[unsafe.Pointer](mark_end2253)
	v992 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v991)(v992)
	v993 = *libc.As[int32](lookahead)
	cmp2254 = v993 == 114
	if cmp2254 {
		goto if_then2256
	} else {
		goto if_end2257
	}

if_then2256:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end2257:
	v994 = *libc.As[int32](lookahead)
	cmp2258 = 48 <= v994
	if cmp2258 {
		goto land_lhs_true2260
	} else {
		goto lor_lhs_false2263
	}

land_lhs_true2260:
	v995 = *libc.As[int32](lookahead)
	cmp2261 = v995 <= 57
	if cmp2261 {
		goto if_then2278
	} else {
		goto lor_lhs_false2263
	}

lor_lhs_false2263:
	v996 = *libc.As[int32](lookahead)
	cmp2264 = 65 <= v996
	if cmp2264 {
		goto land_lhs_true2266
	} else {
		goto lor_lhs_false2269
	}

land_lhs_true2266:
	v997 = *libc.As[int32](lookahead)
	cmp2267 = v997 <= 90
	if cmp2267 {
		goto if_then2278
	} else {
		goto lor_lhs_false2269
	}

lor_lhs_false2269:
	v998 = *libc.As[int32](lookahead)
	cmp2270 = v998 == 95
	if cmp2270 {
		goto if_then2278
	} else {
		goto lor_lhs_false2272
	}

lor_lhs_false2272:
	v999 = *libc.As[int32](lookahead)
	cmp2273 = 97 <= v999
	if cmp2273 {
		goto land_lhs_true2275
	} else {
		goto if_end2279
	}

land_lhs_true2275:
	v1000 = *libc.As[int32](lookahead)
	cmp2276 = v1000 <= 122
	if cmp2276 {
		goto if_then2278
	} else {
		goto if_end2279
	}

if_then2278:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2279:
	v1001 = *libc.As[byte](result)
	loadedv2280 = (v1001 & 1) != 0
	*libc.As[bool](retval) = loadedv2280
	goto _return

sw_bb2281:
	*libc.As[byte](result) = 1
	v1002 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2282 = libc.Ptr(&libc.As[TSLexer](v1002).F1)
	*libc.As[int16](result_symbol2282) = 13
	v1003 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2283 = libc.Ptr(&libc.As[TSLexer](v1003).F3)
	v1004 = *libc.As[unsafe.Pointer](mark_end2283)
	v1005 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1004)(v1005)
	v1006 = *libc.As[int32](lookahead)
	cmp2284 = v1006 == 114
	if cmp2284 {
		goto if_then2286
	} else {
		goto if_end2287
	}

if_then2286:
	*libc.As[int16](state_addr) = 85
	goto next_state

if_end2287:
	v1007 = *libc.As[int32](lookahead)
	cmp2288 = 48 <= v1007
	if cmp2288 {
		goto land_lhs_true2290
	} else {
		goto lor_lhs_false2293
	}

land_lhs_true2290:
	v1008 = *libc.As[int32](lookahead)
	cmp2291 = v1008 <= 57
	if cmp2291 {
		goto if_then2308
	} else {
		goto lor_lhs_false2293
	}

lor_lhs_false2293:
	v1009 = *libc.As[int32](lookahead)
	cmp2294 = 65 <= v1009
	if cmp2294 {
		goto land_lhs_true2296
	} else {
		goto lor_lhs_false2299
	}

land_lhs_true2296:
	v1010 = *libc.As[int32](lookahead)
	cmp2297 = v1010 <= 90
	if cmp2297 {
		goto if_then2308
	} else {
		goto lor_lhs_false2299
	}

lor_lhs_false2299:
	v1011 = *libc.As[int32](lookahead)
	cmp2300 = v1011 == 95
	if cmp2300 {
		goto if_then2308
	} else {
		goto lor_lhs_false2302
	}

lor_lhs_false2302:
	v1012 = *libc.As[int32](lookahead)
	cmp2303 = 97 <= v1012
	if cmp2303 {
		goto land_lhs_true2305
	} else {
		goto if_end2309
	}

land_lhs_true2305:
	v1013 = *libc.As[int32](lookahead)
	cmp2306 = v1013 <= 122
	if cmp2306 {
		goto if_then2308
	} else {
		goto if_end2309
	}

if_then2308:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2309:
	v1014 = *libc.As[byte](result)
	loadedv2310 = (v1014 & 1) != 0
	*libc.As[bool](retval) = loadedv2310
	goto _return

sw_bb2311:
	*libc.As[byte](result) = 1
	v1015 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2312 = libc.Ptr(&libc.As[TSLexer](v1015).F1)
	*libc.As[int16](result_symbol2312) = 13
	v1016 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2313 = libc.Ptr(&libc.As[TSLexer](v1016).F3)
	v1017 = *libc.As[unsafe.Pointer](mark_end2313)
	v1018 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1017)(v1018)
	v1019 = *libc.As[int32](lookahead)
	cmp2314 = v1019 == 114
	if cmp2314 {
		goto if_then2316
	} else {
		goto if_end2317
	}

if_then2316:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end2317:
	v1020 = *libc.As[int32](lookahead)
	cmp2318 = 48 <= v1020
	if cmp2318 {
		goto land_lhs_true2320
	} else {
		goto lor_lhs_false2323
	}

land_lhs_true2320:
	v1021 = *libc.As[int32](lookahead)
	cmp2321 = v1021 <= 57
	if cmp2321 {
		goto if_then2338
	} else {
		goto lor_lhs_false2323
	}

lor_lhs_false2323:
	v1022 = *libc.As[int32](lookahead)
	cmp2324 = 65 <= v1022
	if cmp2324 {
		goto land_lhs_true2326
	} else {
		goto lor_lhs_false2329
	}

land_lhs_true2326:
	v1023 = *libc.As[int32](lookahead)
	cmp2327 = v1023 <= 90
	if cmp2327 {
		goto if_then2338
	} else {
		goto lor_lhs_false2329
	}

lor_lhs_false2329:
	v1024 = *libc.As[int32](lookahead)
	cmp2330 = v1024 == 95
	if cmp2330 {
		goto if_then2338
	} else {
		goto lor_lhs_false2332
	}

lor_lhs_false2332:
	v1025 = *libc.As[int32](lookahead)
	cmp2333 = 97 <= v1025
	if cmp2333 {
		goto land_lhs_true2335
	} else {
		goto if_end2339
	}

land_lhs_true2335:
	v1026 = *libc.As[int32](lookahead)
	cmp2336 = v1026 <= 122
	if cmp2336 {
		goto if_then2338
	} else {
		goto if_end2339
	}

if_then2338:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2339:
	v1027 = *libc.As[byte](result)
	loadedv2340 = (v1027 & 1) != 0
	*libc.As[bool](retval) = loadedv2340
	goto _return

sw_bb2341:
	*libc.As[byte](result) = 1
	v1028 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2342 = libc.Ptr(&libc.As[TSLexer](v1028).F1)
	*libc.As[int16](result_symbol2342) = 13
	v1029 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2343 = libc.Ptr(&libc.As[TSLexer](v1029).F3)
	v1030 = *libc.As[unsafe.Pointer](mark_end2343)
	v1031 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1030)(v1031)
	v1032 = *libc.As[int32](lookahead)
	cmp2344 = v1032 == 114
	if cmp2344 {
		goto if_then2346
	} else {
		goto if_end2347
	}

if_then2346:
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end2347:
	v1033 = *libc.As[int32](lookahead)
	cmp2348 = 48 <= v1033
	if cmp2348 {
		goto land_lhs_true2350
	} else {
		goto lor_lhs_false2353
	}

land_lhs_true2350:
	v1034 = *libc.As[int32](lookahead)
	cmp2351 = v1034 <= 57
	if cmp2351 {
		goto if_then2368
	} else {
		goto lor_lhs_false2353
	}

lor_lhs_false2353:
	v1035 = *libc.As[int32](lookahead)
	cmp2354 = 65 <= v1035
	if cmp2354 {
		goto land_lhs_true2356
	} else {
		goto lor_lhs_false2359
	}

land_lhs_true2356:
	v1036 = *libc.As[int32](lookahead)
	cmp2357 = v1036 <= 90
	if cmp2357 {
		goto if_then2368
	} else {
		goto lor_lhs_false2359
	}

lor_lhs_false2359:
	v1037 = *libc.As[int32](lookahead)
	cmp2360 = v1037 == 95
	if cmp2360 {
		goto if_then2368
	} else {
		goto lor_lhs_false2362
	}

lor_lhs_false2362:
	v1038 = *libc.As[int32](lookahead)
	cmp2363 = 97 <= v1038
	if cmp2363 {
		goto land_lhs_true2365
	} else {
		goto if_end2369
	}

land_lhs_true2365:
	v1039 = *libc.As[int32](lookahead)
	cmp2366 = v1039 <= 122
	if cmp2366 {
		goto if_then2368
	} else {
		goto if_end2369
	}

if_then2368:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2369:
	v1040 = *libc.As[byte](result)
	loadedv2370 = (v1040 & 1) != 0
	*libc.As[bool](retval) = loadedv2370
	goto _return

sw_bb2371:
	*libc.As[byte](result) = 1
	v1041 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2372 = libc.Ptr(&libc.As[TSLexer](v1041).F1)
	*libc.As[int16](result_symbol2372) = 13
	v1042 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2373 = libc.Ptr(&libc.As[TSLexer](v1042).F3)
	v1043 = *libc.As[unsafe.Pointer](mark_end2373)
	v1044 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1043)(v1044)
	v1045 = *libc.As[int32](lookahead)
	cmp2374 = v1045 == 114
	if cmp2374 {
		goto if_then2376
	} else {
		goto if_end2377
	}

if_then2376:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end2377:
	v1046 = *libc.As[int32](lookahead)
	cmp2378 = 48 <= v1046
	if cmp2378 {
		goto land_lhs_true2380
	} else {
		goto lor_lhs_false2383
	}

land_lhs_true2380:
	v1047 = *libc.As[int32](lookahead)
	cmp2381 = v1047 <= 57
	if cmp2381 {
		goto if_then2398
	} else {
		goto lor_lhs_false2383
	}

lor_lhs_false2383:
	v1048 = *libc.As[int32](lookahead)
	cmp2384 = 65 <= v1048
	if cmp2384 {
		goto land_lhs_true2386
	} else {
		goto lor_lhs_false2389
	}

land_lhs_true2386:
	v1049 = *libc.As[int32](lookahead)
	cmp2387 = v1049 <= 90
	if cmp2387 {
		goto if_then2398
	} else {
		goto lor_lhs_false2389
	}

lor_lhs_false2389:
	v1050 = *libc.As[int32](lookahead)
	cmp2390 = v1050 == 95
	if cmp2390 {
		goto if_then2398
	} else {
		goto lor_lhs_false2392
	}

lor_lhs_false2392:
	v1051 = *libc.As[int32](lookahead)
	cmp2393 = 97 <= v1051
	if cmp2393 {
		goto land_lhs_true2395
	} else {
		goto if_end2399
	}

land_lhs_true2395:
	v1052 = *libc.As[int32](lookahead)
	cmp2396 = v1052 <= 122
	if cmp2396 {
		goto if_then2398
	} else {
		goto if_end2399
	}

if_then2398:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2399:
	v1053 = *libc.As[byte](result)
	loadedv2400 = (v1053 & 1) != 0
	*libc.As[bool](retval) = loadedv2400
	goto _return

sw_bb2401:
	*libc.As[byte](result) = 1
	v1054 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2402 = libc.Ptr(&libc.As[TSLexer](v1054).F1)
	*libc.As[int16](result_symbol2402) = 13
	v1055 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2403 = libc.Ptr(&libc.As[TSLexer](v1055).F3)
	v1056 = *libc.As[unsafe.Pointer](mark_end2403)
	v1057 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1056)(v1057)
	v1058 = *libc.As[int32](lookahead)
	cmp2404 = v1058 == 115
	if cmp2404 {
		goto if_then2406
	} else {
		goto if_end2407
	}

if_then2406:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end2407:
	v1059 = *libc.As[int32](lookahead)
	cmp2408 = 48 <= v1059
	if cmp2408 {
		goto land_lhs_true2410
	} else {
		goto lor_lhs_false2413
	}

land_lhs_true2410:
	v1060 = *libc.As[int32](lookahead)
	cmp2411 = v1060 <= 57
	if cmp2411 {
		goto if_then2428
	} else {
		goto lor_lhs_false2413
	}

lor_lhs_false2413:
	v1061 = *libc.As[int32](lookahead)
	cmp2414 = 65 <= v1061
	if cmp2414 {
		goto land_lhs_true2416
	} else {
		goto lor_lhs_false2419
	}

land_lhs_true2416:
	v1062 = *libc.As[int32](lookahead)
	cmp2417 = v1062 <= 90
	if cmp2417 {
		goto if_then2428
	} else {
		goto lor_lhs_false2419
	}

lor_lhs_false2419:
	v1063 = *libc.As[int32](lookahead)
	cmp2420 = v1063 == 95
	if cmp2420 {
		goto if_then2428
	} else {
		goto lor_lhs_false2422
	}

lor_lhs_false2422:
	v1064 = *libc.As[int32](lookahead)
	cmp2423 = 97 <= v1064
	if cmp2423 {
		goto land_lhs_true2425
	} else {
		goto if_end2429
	}

land_lhs_true2425:
	v1065 = *libc.As[int32](lookahead)
	cmp2426 = v1065 <= 122
	if cmp2426 {
		goto if_then2428
	} else {
		goto if_end2429
	}

if_then2428:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2429:
	v1066 = *libc.As[byte](result)
	loadedv2430 = (v1066 & 1) != 0
	*libc.As[bool](retval) = loadedv2430
	goto _return

sw_bb2431:
	*libc.As[byte](result) = 1
	v1067 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2432 = libc.Ptr(&libc.As[TSLexer](v1067).F1)
	*libc.As[int16](result_symbol2432) = 13
	v1068 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2433 = libc.Ptr(&libc.As[TSLexer](v1068).F3)
	v1069 = *libc.As[unsafe.Pointer](mark_end2433)
	v1070 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1069)(v1070)
	v1071 = *libc.As[int32](lookahead)
	cmp2434 = v1071 == 115
	if cmp2434 {
		goto if_then2436
	} else {
		goto if_end2437
	}

if_then2436:
	*libc.As[int16](state_addr) = 86
	goto next_state

if_end2437:
	v1072 = *libc.As[int32](lookahead)
	cmp2438 = 48 <= v1072
	if cmp2438 {
		goto land_lhs_true2440
	} else {
		goto lor_lhs_false2443
	}

land_lhs_true2440:
	v1073 = *libc.As[int32](lookahead)
	cmp2441 = v1073 <= 57
	if cmp2441 {
		goto if_then2458
	} else {
		goto lor_lhs_false2443
	}

lor_lhs_false2443:
	v1074 = *libc.As[int32](lookahead)
	cmp2444 = 65 <= v1074
	if cmp2444 {
		goto land_lhs_true2446
	} else {
		goto lor_lhs_false2449
	}

land_lhs_true2446:
	v1075 = *libc.As[int32](lookahead)
	cmp2447 = v1075 <= 90
	if cmp2447 {
		goto if_then2458
	} else {
		goto lor_lhs_false2449
	}

lor_lhs_false2449:
	v1076 = *libc.As[int32](lookahead)
	cmp2450 = v1076 == 95
	if cmp2450 {
		goto if_then2458
	} else {
		goto lor_lhs_false2452
	}

lor_lhs_false2452:
	v1077 = *libc.As[int32](lookahead)
	cmp2453 = 97 <= v1077
	if cmp2453 {
		goto land_lhs_true2455
	} else {
		goto if_end2459
	}

land_lhs_true2455:
	v1078 = *libc.As[int32](lookahead)
	cmp2456 = v1078 <= 122
	if cmp2456 {
		goto if_then2458
	} else {
		goto if_end2459
	}

if_then2458:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2459:
	v1079 = *libc.As[byte](result)
	loadedv2460 = (v1079 & 1) != 0
	*libc.As[bool](retval) = loadedv2460
	goto _return

sw_bb2461:
	*libc.As[byte](result) = 1
	v1080 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2462 = libc.Ptr(&libc.As[TSLexer](v1080).F1)
	*libc.As[int16](result_symbol2462) = 13
	v1081 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2463 = libc.Ptr(&libc.As[TSLexer](v1081).F3)
	v1082 = *libc.As[unsafe.Pointer](mark_end2463)
	v1083 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1082)(v1083)
	v1084 = *libc.As[int32](lookahead)
	cmp2464 = v1084 == 115
	if cmp2464 {
		goto if_then2466
	} else {
		goto if_end2467
	}

if_then2466:
	*libc.As[int16](state_addr) = 95
	goto next_state

if_end2467:
	v1085 = *libc.As[int32](lookahead)
	cmp2468 = 48 <= v1085
	if cmp2468 {
		goto land_lhs_true2470
	} else {
		goto lor_lhs_false2473
	}

land_lhs_true2470:
	v1086 = *libc.As[int32](lookahead)
	cmp2471 = v1086 <= 57
	if cmp2471 {
		goto if_then2488
	} else {
		goto lor_lhs_false2473
	}

lor_lhs_false2473:
	v1087 = *libc.As[int32](lookahead)
	cmp2474 = 65 <= v1087
	if cmp2474 {
		goto land_lhs_true2476
	} else {
		goto lor_lhs_false2479
	}

land_lhs_true2476:
	v1088 = *libc.As[int32](lookahead)
	cmp2477 = v1088 <= 90
	if cmp2477 {
		goto if_then2488
	} else {
		goto lor_lhs_false2479
	}

lor_lhs_false2479:
	v1089 = *libc.As[int32](lookahead)
	cmp2480 = v1089 == 95
	if cmp2480 {
		goto if_then2488
	} else {
		goto lor_lhs_false2482
	}

lor_lhs_false2482:
	v1090 = *libc.As[int32](lookahead)
	cmp2483 = 97 <= v1090
	if cmp2483 {
		goto land_lhs_true2485
	} else {
		goto if_end2489
	}

land_lhs_true2485:
	v1091 = *libc.As[int32](lookahead)
	cmp2486 = v1091 <= 122
	if cmp2486 {
		goto if_then2488
	} else {
		goto if_end2489
	}

if_then2488:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2489:
	v1092 = *libc.As[byte](result)
	loadedv2490 = (v1092 & 1) != 0
	*libc.As[bool](retval) = loadedv2490
	goto _return

sw_bb2491:
	*libc.As[byte](result) = 1
	v1093 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2492 = libc.Ptr(&libc.As[TSLexer](v1093).F1)
	*libc.As[int16](result_symbol2492) = 13
	v1094 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2493 = libc.Ptr(&libc.As[TSLexer](v1094).F3)
	v1095 = *libc.As[unsafe.Pointer](mark_end2493)
	v1096 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1095)(v1096)
	v1097 = *libc.As[int32](lookahead)
	cmp2494 = v1097 == 115
	if cmp2494 {
		goto if_then2496
	} else {
		goto if_end2497
	}

if_then2496:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end2497:
	v1098 = *libc.As[int32](lookahead)
	cmp2498 = 48 <= v1098
	if cmp2498 {
		goto land_lhs_true2500
	} else {
		goto lor_lhs_false2503
	}

land_lhs_true2500:
	v1099 = *libc.As[int32](lookahead)
	cmp2501 = v1099 <= 57
	if cmp2501 {
		goto if_then2518
	} else {
		goto lor_lhs_false2503
	}

lor_lhs_false2503:
	v1100 = *libc.As[int32](lookahead)
	cmp2504 = 65 <= v1100
	if cmp2504 {
		goto land_lhs_true2506
	} else {
		goto lor_lhs_false2509
	}

land_lhs_true2506:
	v1101 = *libc.As[int32](lookahead)
	cmp2507 = v1101 <= 90
	if cmp2507 {
		goto if_then2518
	} else {
		goto lor_lhs_false2509
	}

lor_lhs_false2509:
	v1102 = *libc.As[int32](lookahead)
	cmp2510 = v1102 == 95
	if cmp2510 {
		goto if_then2518
	} else {
		goto lor_lhs_false2512
	}

lor_lhs_false2512:
	v1103 = *libc.As[int32](lookahead)
	cmp2513 = 97 <= v1103
	if cmp2513 {
		goto land_lhs_true2515
	} else {
		goto if_end2519
	}

land_lhs_true2515:
	v1104 = *libc.As[int32](lookahead)
	cmp2516 = v1104 <= 122
	if cmp2516 {
		goto if_then2518
	} else {
		goto if_end2519
	}

if_then2518:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2519:
	v1105 = *libc.As[byte](result)
	loadedv2520 = (v1105 & 1) != 0
	*libc.As[bool](retval) = loadedv2520
	goto _return

sw_bb2521:
	*libc.As[byte](result) = 1
	v1106 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2522 = libc.Ptr(&libc.As[TSLexer](v1106).F1)
	*libc.As[int16](result_symbol2522) = 13
	v1107 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2523 = libc.Ptr(&libc.As[TSLexer](v1107).F3)
	v1108 = *libc.As[unsafe.Pointer](mark_end2523)
	v1109 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1108)(v1109)
	v1110 = *libc.As[int32](lookahead)
	cmp2524 = v1110 == 116
	if cmp2524 {
		goto if_then2526
	} else {
		goto if_end2527
	}

if_then2526:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end2527:
	v1111 = *libc.As[int32](lookahead)
	cmp2528 = 48 <= v1111
	if cmp2528 {
		goto land_lhs_true2530
	} else {
		goto lor_lhs_false2533
	}

land_lhs_true2530:
	v1112 = *libc.As[int32](lookahead)
	cmp2531 = v1112 <= 57
	if cmp2531 {
		goto if_then2548
	} else {
		goto lor_lhs_false2533
	}

lor_lhs_false2533:
	v1113 = *libc.As[int32](lookahead)
	cmp2534 = 65 <= v1113
	if cmp2534 {
		goto land_lhs_true2536
	} else {
		goto lor_lhs_false2539
	}

land_lhs_true2536:
	v1114 = *libc.As[int32](lookahead)
	cmp2537 = v1114 <= 90
	if cmp2537 {
		goto if_then2548
	} else {
		goto lor_lhs_false2539
	}

lor_lhs_false2539:
	v1115 = *libc.As[int32](lookahead)
	cmp2540 = v1115 == 95
	if cmp2540 {
		goto if_then2548
	} else {
		goto lor_lhs_false2542
	}

lor_lhs_false2542:
	v1116 = *libc.As[int32](lookahead)
	cmp2543 = 97 <= v1116
	if cmp2543 {
		goto land_lhs_true2545
	} else {
		goto if_end2549
	}

land_lhs_true2545:
	v1117 = *libc.As[int32](lookahead)
	cmp2546 = v1117 <= 122
	if cmp2546 {
		goto if_then2548
	} else {
		goto if_end2549
	}

if_then2548:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2549:
	v1118 = *libc.As[byte](result)
	loadedv2550 = (v1118 & 1) != 0
	*libc.As[bool](retval) = loadedv2550
	goto _return

sw_bb2551:
	*libc.As[byte](result) = 1
	v1119 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2552 = libc.Ptr(&libc.As[TSLexer](v1119).F1)
	*libc.As[int16](result_symbol2552) = 13
	v1120 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2553 = libc.Ptr(&libc.As[TSLexer](v1120).F3)
	v1121 = *libc.As[unsafe.Pointer](mark_end2553)
	v1122 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1121)(v1122)
	v1123 = *libc.As[int32](lookahead)
	cmp2554 = v1123 == 116
	if cmp2554 {
		goto if_then2556
	} else {
		goto if_end2557
	}

if_then2556:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end2557:
	v1124 = *libc.As[int32](lookahead)
	cmp2558 = 48 <= v1124
	if cmp2558 {
		goto land_lhs_true2560
	} else {
		goto lor_lhs_false2563
	}

land_lhs_true2560:
	v1125 = *libc.As[int32](lookahead)
	cmp2561 = v1125 <= 57
	if cmp2561 {
		goto if_then2578
	} else {
		goto lor_lhs_false2563
	}

lor_lhs_false2563:
	v1126 = *libc.As[int32](lookahead)
	cmp2564 = 65 <= v1126
	if cmp2564 {
		goto land_lhs_true2566
	} else {
		goto lor_lhs_false2569
	}

land_lhs_true2566:
	v1127 = *libc.As[int32](lookahead)
	cmp2567 = v1127 <= 90
	if cmp2567 {
		goto if_then2578
	} else {
		goto lor_lhs_false2569
	}

lor_lhs_false2569:
	v1128 = *libc.As[int32](lookahead)
	cmp2570 = v1128 == 95
	if cmp2570 {
		goto if_then2578
	} else {
		goto lor_lhs_false2572
	}

lor_lhs_false2572:
	v1129 = *libc.As[int32](lookahead)
	cmp2573 = 97 <= v1129
	if cmp2573 {
		goto land_lhs_true2575
	} else {
		goto if_end2579
	}

land_lhs_true2575:
	v1130 = *libc.As[int32](lookahead)
	cmp2576 = v1130 <= 122
	if cmp2576 {
		goto if_then2578
	} else {
		goto if_end2579
	}

if_then2578:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2579:
	v1131 = *libc.As[byte](result)
	loadedv2580 = (v1131 & 1) != 0
	*libc.As[bool](retval) = loadedv2580
	goto _return

sw_bb2581:
	*libc.As[byte](result) = 1
	v1132 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2582 = libc.Ptr(&libc.As[TSLexer](v1132).F1)
	*libc.As[int16](result_symbol2582) = 13
	v1133 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2583 = libc.Ptr(&libc.As[TSLexer](v1133).F3)
	v1134 = *libc.As[unsafe.Pointer](mark_end2583)
	v1135 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1134)(v1135)
	v1136 = *libc.As[int32](lookahead)
	cmp2584 = v1136 == 116
	if cmp2584 {
		goto if_then2586
	} else {
		goto if_end2587
	}

if_then2586:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end2587:
	v1137 = *libc.As[int32](lookahead)
	cmp2588 = 48 <= v1137
	if cmp2588 {
		goto land_lhs_true2590
	} else {
		goto lor_lhs_false2593
	}

land_lhs_true2590:
	v1138 = *libc.As[int32](lookahead)
	cmp2591 = v1138 <= 57
	if cmp2591 {
		goto if_then2608
	} else {
		goto lor_lhs_false2593
	}

lor_lhs_false2593:
	v1139 = *libc.As[int32](lookahead)
	cmp2594 = 65 <= v1139
	if cmp2594 {
		goto land_lhs_true2596
	} else {
		goto lor_lhs_false2599
	}

land_lhs_true2596:
	v1140 = *libc.As[int32](lookahead)
	cmp2597 = v1140 <= 90
	if cmp2597 {
		goto if_then2608
	} else {
		goto lor_lhs_false2599
	}

lor_lhs_false2599:
	v1141 = *libc.As[int32](lookahead)
	cmp2600 = v1141 == 95
	if cmp2600 {
		goto if_then2608
	} else {
		goto lor_lhs_false2602
	}

lor_lhs_false2602:
	v1142 = *libc.As[int32](lookahead)
	cmp2603 = 97 <= v1142
	if cmp2603 {
		goto land_lhs_true2605
	} else {
		goto if_end2609
	}

land_lhs_true2605:
	v1143 = *libc.As[int32](lookahead)
	cmp2606 = v1143 <= 122
	if cmp2606 {
		goto if_then2608
	} else {
		goto if_end2609
	}

if_then2608:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2609:
	v1144 = *libc.As[byte](result)
	loadedv2610 = (v1144 & 1) != 0
	*libc.As[bool](retval) = loadedv2610
	goto _return

sw_bb2611:
	*libc.As[byte](result) = 1
	v1145 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2612 = libc.Ptr(&libc.As[TSLexer](v1145).F1)
	*libc.As[int16](result_symbol2612) = 13
	v1146 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2613 = libc.Ptr(&libc.As[TSLexer](v1146).F3)
	v1147 = *libc.As[unsafe.Pointer](mark_end2613)
	v1148 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1147)(v1148)
	v1149 = *libc.As[int32](lookahead)
	cmp2614 = v1149 == 116
	if cmp2614 {
		goto if_then2616
	} else {
		goto if_end2617
	}

if_then2616:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end2617:
	v1150 = *libc.As[int32](lookahead)
	cmp2618 = 48 <= v1150
	if cmp2618 {
		goto land_lhs_true2620
	} else {
		goto lor_lhs_false2623
	}

land_lhs_true2620:
	v1151 = *libc.As[int32](lookahead)
	cmp2621 = v1151 <= 57
	if cmp2621 {
		goto if_then2638
	} else {
		goto lor_lhs_false2623
	}

lor_lhs_false2623:
	v1152 = *libc.As[int32](lookahead)
	cmp2624 = 65 <= v1152
	if cmp2624 {
		goto land_lhs_true2626
	} else {
		goto lor_lhs_false2629
	}

land_lhs_true2626:
	v1153 = *libc.As[int32](lookahead)
	cmp2627 = v1153 <= 90
	if cmp2627 {
		goto if_then2638
	} else {
		goto lor_lhs_false2629
	}

lor_lhs_false2629:
	v1154 = *libc.As[int32](lookahead)
	cmp2630 = v1154 == 95
	if cmp2630 {
		goto if_then2638
	} else {
		goto lor_lhs_false2632
	}

lor_lhs_false2632:
	v1155 = *libc.As[int32](lookahead)
	cmp2633 = 97 <= v1155
	if cmp2633 {
		goto land_lhs_true2635
	} else {
		goto if_end2639
	}

land_lhs_true2635:
	v1156 = *libc.As[int32](lookahead)
	cmp2636 = v1156 <= 122
	if cmp2636 {
		goto if_then2638
	} else {
		goto if_end2639
	}

if_then2638:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2639:
	v1157 = *libc.As[byte](result)
	loadedv2640 = (v1157 & 1) != 0
	*libc.As[bool](retval) = loadedv2640
	goto _return

sw_bb2641:
	*libc.As[byte](result) = 1
	v1158 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2642 = libc.Ptr(&libc.As[TSLexer](v1158).F1)
	*libc.As[int16](result_symbol2642) = 13
	v1159 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2643 = libc.Ptr(&libc.As[TSLexer](v1159).F3)
	v1160 = *libc.As[unsafe.Pointer](mark_end2643)
	v1161 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1160)(v1161)
	v1162 = *libc.As[int32](lookahead)
	cmp2644 = v1162 == 116
	if cmp2644 {
		goto if_then2646
	} else {
		goto if_end2647
	}

if_then2646:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end2647:
	v1163 = *libc.As[int32](lookahead)
	cmp2648 = 48 <= v1163
	if cmp2648 {
		goto land_lhs_true2650
	} else {
		goto lor_lhs_false2653
	}

land_lhs_true2650:
	v1164 = *libc.As[int32](lookahead)
	cmp2651 = v1164 <= 57
	if cmp2651 {
		goto if_then2668
	} else {
		goto lor_lhs_false2653
	}

lor_lhs_false2653:
	v1165 = *libc.As[int32](lookahead)
	cmp2654 = 65 <= v1165
	if cmp2654 {
		goto land_lhs_true2656
	} else {
		goto lor_lhs_false2659
	}

land_lhs_true2656:
	v1166 = *libc.As[int32](lookahead)
	cmp2657 = v1166 <= 90
	if cmp2657 {
		goto if_then2668
	} else {
		goto lor_lhs_false2659
	}

lor_lhs_false2659:
	v1167 = *libc.As[int32](lookahead)
	cmp2660 = v1167 == 95
	if cmp2660 {
		goto if_then2668
	} else {
		goto lor_lhs_false2662
	}

lor_lhs_false2662:
	v1168 = *libc.As[int32](lookahead)
	cmp2663 = 97 <= v1168
	if cmp2663 {
		goto land_lhs_true2665
	} else {
		goto if_end2669
	}

land_lhs_true2665:
	v1169 = *libc.As[int32](lookahead)
	cmp2666 = v1169 <= 122
	if cmp2666 {
		goto if_then2668
	} else {
		goto if_end2669
	}

if_then2668:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2669:
	v1170 = *libc.As[byte](result)
	loadedv2670 = (v1170 & 1) != 0
	*libc.As[bool](retval) = loadedv2670
	goto _return

sw_bb2671:
	*libc.As[byte](result) = 1
	v1171 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2672 = libc.Ptr(&libc.As[TSLexer](v1171).F1)
	*libc.As[int16](result_symbol2672) = 13
	v1172 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2673 = libc.Ptr(&libc.As[TSLexer](v1172).F3)
	v1173 = *libc.As[unsafe.Pointer](mark_end2673)
	v1174 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1173)(v1174)
	v1175 = *libc.As[int32](lookahead)
	cmp2674 = v1175 == 116
	if cmp2674 {
		goto if_then2676
	} else {
		goto if_end2677
	}

if_then2676:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end2677:
	v1176 = *libc.As[int32](lookahead)
	cmp2678 = 48 <= v1176
	if cmp2678 {
		goto land_lhs_true2680
	} else {
		goto lor_lhs_false2683
	}

land_lhs_true2680:
	v1177 = *libc.As[int32](lookahead)
	cmp2681 = v1177 <= 57
	if cmp2681 {
		goto if_then2698
	} else {
		goto lor_lhs_false2683
	}

lor_lhs_false2683:
	v1178 = *libc.As[int32](lookahead)
	cmp2684 = 65 <= v1178
	if cmp2684 {
		goto land_lhs_true2686
	} else {
		goto lor_lhs_false2689
	}

land_lhs_true2686:
	v1179 = *libc.As[int32](lookahead)
	cmp2687 = v1179 <= 90
	if cmp2687 {
		goto if_then2698
	} else {
		goto lor_lhs_false2689
	}

lor_lhs_false2689:
	v1180 = *libc.As[int32](lookahead)
	cmp2690 = v1180 == 95
	if cmp2690 {
		goto if_then2698
	} else {
		goto lor_lhs_false2692
	}

lor_lhs_false2692:
	v1181 = *libc.As[int32](lookahead)
	cmp2693 = 97 <= v1181
	if cmp2693 {
		goto land_lhs_true2695
	} else {
		goto if_end2699
	}

land_lhs_true2695:
	v1182 = *libc.As[int32](lookahead)
	cmp2696 = v1182 <= 122
	if cmp2696 {
		goto if_then2698
	} else {
		goto if_end2699
	}

if_then2698:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2699:
	v1183 = *libc.As[byte](result)
	loadedv2700 = (v1183 & 1) != 0
	*libc.As[bool](retval) = loadedv2700
	goto _return

sw_bb2701:
	*libc.As[byte](result) = 1
	v1184 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2702 = libc.Ptr(&libc.As[TSLexer](v1184).F1)
	*libc.As[int16](result_symbol2702) = 13
	v1185 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2703 = libc.Ptr(&libc.As[TSLexer](v1185).F3)
	v1186 = *libc.As[unsafe.Pointer](mark_end2703)
	v1187 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1186)(v1187)
	v1188 = *libc.As[int32](lookahead)
	cmp2704 = v1188 == 117
	if cmp2704 {
		goto if_then2706
	} else {
		goto if_end2707
	}

if_then2706:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end2707:
	v1189 = *libc.As[int32](lookahead)
	cmp2708 = 48 <= v1189
	if cmp2708 {
		goto land_lhs_true2710
	} else {
		goto lor_lhs_false2713
	}

land_lhs_true2710:
	v1190 = *libc.As[int32](lookahead)
	cmp2711 = v1190 <= 57
	if cmp2711 {
		goto if_then2728
	} else {
		goto lor_lhs_false2713
	}

lor_lhs_false2713:
	v1191 = *libc.As[int32](lookahead)
	cmp2714 = 65 <= v1191
	if cmp2714 {
		goto land_lhs_true2716
	} else {
		goto lor_lhs_false2719
	}

land_lhs_true2716:
	v1192 = *libc.As[int32](lookahead)
	cmp2717 = v1192 <= 90
	if cmp2717 {
		goto if_then2728
	} else {
		goto lor_lhs_false2719
	}

lor_lhs_false2719:
	v1193 = *libc.As[int32](lookahead)
	cmp2720 = v1193 == 95
	if cmp2720 {
		goto if_then2728
	} else {
		goto lor_lhs_false2722
	}

lor_lhs_false2722:
	v1194 = *libc.As[int32](lookahead)
	cmp2723 = 97 <= v1194
	if cmp2723 {
		goto land_lhs_true2725
	} else {
		goto if_end2729
	}

land_lhs_true2725:
	v1195 = *libc.As[int32](lookahead)
	cmp2726 = v1195 <= 122
	if cmp2726 {
		goto if_then2728
	} else {
		goto if_end2729
	}

if_then2728:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2729:
	v1196 = *libc.As[byte](result)
	loadedv2730 = (v1196 & 1) != 0
	*libc.As[bool](retval) = loadedv2730
	goto _return

sw_bb2731:
	*libc.As[byte](result) = 1
	v1197 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2732 = libc.Ptr(&libc.As[TSLexer](v1197).F1)
	*libc.As[int16](result_symbol2732) = 13
	v1198 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2733 = libc.Ptr(&libc.As[TSLexer](v1198).F3)
	v1199 = *libc.As[unsafe.Pointer](mark_end2733)
	v1200 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1199)(v1200)
	v1201 = *libc.As[int32](lookahead)
	cmp2734 = v1201 == 117
	if cmp2734 {
		goto if_then2736
	} else {
		goto if_end2737
	}

if_then2736:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end2737:
	v1202 = *libc.As[int32](lookahead)
	cmp2738 = 48 <= v1202
	if cmp2738 {
		goto land_lhs_true2740
	} else {
		goto lor_lhs_false2743
	}

land_lhs_true2740:
	v1203 = *libc.As[int32](lookahead)
	cmp2741 = v1203 <= 57
	if cmp2741 {
		goto if_then2758
	} else {
		goto lor_lhs_false2743
	}

lor_lhs_false2743:
	v1204 = *libc.As[int32](lookahead)
	cmp2744 = 65 <= v1204
	if cmp2744 {
		goto land_lhs_true2746
	} else {
		goto lor_lhs_false2749
	}

land_lhs_true2746:
	v1205 = *libc.As[int32](lookahead)
	cmp2747 = v1205 <= 90
	if cmp2747 {
		goto if_then2758
	} else {
		goto lor_lhs_false2749
	}

lor_lhs_false2749:
	v1206 = *libc.As[int32](lookahead)
	cmp2750 = v1206 == 95
	if cmp2750 {
		goto if_then2758
	} else {
		goto lor_lhs_false2752
	}

lor_lhs_false2752:
	v1207 = *libc.As[int32](lookahead)
	cmp2753 = 97 <= v1207
	if cmp2753 {
		goto land_lhs_true2755
	} else {
		goto if_end2759
	}

land_lhs_true2755:
	v1208 = *libc.As[int32](lookahead)
	cmp2756 = v1208 <= 122
	if cmp2756 {
		goto if_then2758
	} else {
		goto if_end2759
	}

if_then2758:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2759:
	v1209 = *libc.As[byte](result)
	loadedv2760 = (v1209 & 1) != 0
	*libc.As[bool](retval) = loadedv2760
	goto _return

sw_bb2761:
	*libc.As[byte](result) = 1
	v1210 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2762 = libc.Ptr(&libc.As[TSLexer](v1210).F1)
	*libc.As[int16](result_symbol2762) = 13
	v1211 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2763 = libc.Ptr(&libc.As[TSLexer](v1211).F3)
	v1212 = *libc.As[unsafe.Pointer](mark_end2763)
	v1213 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1212)(v1213)
	v1214 = *libc.As[int32](lookahead)
	cmp2764 = v1214 == 117
	if cmp2764 {
		goto if_then2766
	} else {
		goto if_end2767
	}

if_then2766:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end2767:
	v1215 = *libc.As[int32](lookahead)
	cmp2768 = 48 <= v1215
	if cmp2768 {
		goto land_lhs_true2770
	} else {
		goto lor_lhs_false2773
	}

land_lhs_true2770:
	v1216 = *libc.As[int32](lookahead)
	cmp2771 = v1216 <= 57
	if cmp2771 {
		goto if_then2788
	} else {
		goto lor_lhs_false2773
	}

lor_lhs_false2773:
	v1217 = *libc.As[int32](lookahead)
	cmp2774 = 65 <= v1217
	if cmp2774 {
		goto land_lhs_true2776
	} else {
		goto lor_lhs_false2779
	}

land_lhs_true2776:
	v1218 = *libc.As[int32](lookahead)
	cmp2777 = v1218 <= 90
	if cmp2777 {
		goto if_then2788
	} else {
		goto lor_lhs_false2779
	}

lor_lhs_false2779:
	v1219 = *libc.As[int32](lookahead)
	cmp2780 = v1219 == 95
	if cmp2780 {
		goto if_then2788
	} else {
		goto lor_lhs_false2782
	}

lor_lhs_false2782:
	v1220 = *libc.As[int32](lookahead)
	cmp2783 = 97 <= v1220
	if cmp2783 {
		goto land_lhs_true2785
	} else {
		goto if_end2789
	}

land_lhs_true2785:
	v1221 = *libc.As[int32](lookahead)
	cmp2786 = v1221 <= 122
	if cmp2786 {
		goto if_then2788
	} else {
		goto if_end2789
	}

if_then2788:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2789:
	v1222 = *libc.As[byte](result)
	loadedv2790 = (v1222 & 1) != 0
	*libc.As[bool](retval) = loadedv2790
	goto _return

sw_bb2791:
	*libc.As[byte](result) = 1
	v1223 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2792 = libc.Ptr(&libc.As[TSLexer](v1223).F1)
	*libc.As[int16](result_symbol2792) = 13
	v1224 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2793 = libc.Ptr(&libc.As[TSLexer](v1224).F3)
	v1225 = *libc.As[unsafe.Pointer](mark_end2793)
	v1226 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1225)(v1226)
	v1227 = *libc.As[int32](lookahead)
	cmp2794 = v1227 == 121
	if cmp2794 {
		goto if_then2796
	} else {
		goto if_end2797
	}

if_then2796:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end2797:
	v1228 = *libc.As[int32](lookahead)
	cmp2798 = 48 <= v1228
	if cmp2798 {
		goto land_lhs_true2800
	} else {
		goto lor_lhs_false2803
	}

land_lhs_true2800:
	v1229 = *libc.As[int32](lookahead)
	cmp2801 = v1229 <= 57
	if cmp2801 {
		goto if_then2818
	} else {
		goto lor_lhs_false2803
	}

lor_lhs_false2803:
	v1230 = *libc.As[int32](lookahead)
	cmp2804 = 65 <= v1230
	if cmp2804 {
		goto land_lhs_true2806
	} else {
		goto lor_lhs_false2809
	}

land_lhs_true2806:
	v1231 = *libc.As[int32](lookahead)
	cmp2807 = v1231 <= 90
	if cmp2807 {
		goto if_then2818
	} else {
		goto lor_lhs_false2809
	}

lor_lhs_false2809:
	v1232 = *libc.As[int32](lookahead)
	cmp2810 = v1232 == 95
	if cmp2810 {
		goto if_then2818
	} else {
		goto lor_lhs_false2812
	}

lor_lhs_false2812:
	v1233 = *libc.As[int32](lookahead)
	cmp2813 = 97 <= v1233
	if cmp2813 {
		goto land_lhs_true2815
	} else {
		goto if_end2819
	}

land_lhs_true2815:
	v1234 = *libc.As[int32](lookahead)
	cmp2816 = v1234 <= 122
	if cmp2816 {
		goto if_then2818
	} else {
		goto if_end2819
	}

if_then2818:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2819:
	v1235 = *libc.As[byte](result)
	loadedv2820 = (v1235 & 1) != 0
	*libc.As[bool](retval) = loadedv2820
	goto _return

sw_bb2821:
	*libc.As[byte](result) = 1
	v1236 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2822 = libc.Ptr(&libc.As[TSLexer](v1236).F1)
	*libc.As[int16](result_symbol2822) = 13
	v1237 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2823 = libc.Ptr(&libc.As[TSLexer](v1237).F3)
	v1238 = *libc.As[unsafe.Pointer](mark_end2823)
	v1239 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1238)(v1239)
	v1240 = *libc.As[int32](lookahead)
	cmp2824 = 48 <= v1240
	if cmp2824 {
		goto land_lhs_true2826
	} else {
		goto lor_lhs_false2829
	}

land_lhs_true2826:
	v1241 = *libc.As[int32](lookahead)
	cmp2827 = v1241 <= 57
	if cmp2827 {
		goto if_then2844
	} else {
		goto lor_lhs_false2829
	}

lor_lhs_false2829:
	v1242 = *libc.As[int32](lookahead)
	cmp2830 = 65 <= v1242
	if cmp2830 {
		goto land_lhs_true2832
	} else {
		goto lor_lhs_false2835
	}

land_lhs_true2832:
	v1243 = *libc.As[int32](lookahead)
	cmp2833 = v1243 <= 90
	if cmp2833 {
		goto if_then2844
	} else {
		goto lor_lhs_false2835
	}

lor_lhs_false2835:
	v1244 = *libc.As[int32](lookahead)
	cmp2836 = v1244 == 95
	if cmp2836 {
		goto if_then2844
	} else {
		goto lor_lhs_false2838
	}

lor_lhs_false2838:
	v1245 = *libc.As[int32](lookahead)
	cmp2839 = 97 <= v1245
	if cmp2839 {
		goto land_lhs_true2841
	} else {
		goto if_end2845
	}

land_lhs_true2841:
	v1246 = *libc.As[int32](lookahead)
	cmp2842 = v1246 <= 122
	if cmp2842 {
		goto if_then2844
	} else {
		goto if_end2845
	}

if_then2844:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end2845:
	v1247 = *libc.As[byte](result)
	loadedv2846 = (v1247 & 1) != 0
	*libc.As[bool](retval) = loadedv2846
	goto _return

sw_bb2847:
	*libc.As[byte](result) = 1
	v1248 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2848 = libc.Ptr(&libc.As[TSLexer](v1248).F1)
	*libc.As[int16](result_symbol2848) = 14
	v1249 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2849 = libc.Ptr(&libc.As[TSLexer](v1249).F3)
	v1250 = *libc.As[unsafe.Pointer](mark_end2849)
	v1251 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1250)(v1251)
	v1252 = *libc.As[int32](lookahead)
	cmp2850 = 48 <= v1252
	if cmp2850 {
		goto land_lhs_true2852
	} else {
		goto if_end2856
	}

land_lhs_true2852:
	v1253 = *libc.As[int32](lookahead)
	cmp2853 = v1253 <= 57
	if cmp2853 {
		goto if_then2855
	} else {
		goto if_end2856
	}

if_then2855:
	*libc.As[int16](state_addr) = 98
	goto next_state

if_end2856:
	v1254 = *libc.As[int32](lookahead)
	cmp2857 = 46 <= v1254
	if cmp2857 {
		goto land_lhs_true2859
	} else {
		goto lor_lhs_false2862
	}

land_lhs_true2859:
	v1255 = *libc.As[int32](lookahead)
	cmp2860 = v1255 <= 58
	if cmp2860 {
		goto if_then2877
	} else {
		goto lor_lhs_false2862
	}

lor_lhs_false2862:
	v1256 = *libc.As[int32](lookahead)
	cmp2863 = 65 <= v1256
	if cmp2863 {
		goto land_lhs_true2865
	} else {
		goto lor_lhs_false2868
	}

land_lhs_true2865:
	v1257 = *libc.As[int32](lookahead)
	cmp2866 = v1257 <= 90
	if cmp2866 {
		goto if_then2877
	} else {
		goto lor_lhs_false2868
	}

lor_lhs_false2868:
	v1258 = *libc.As[int32](lookahead)
	cmp2869 = v1258 == 95
	if cmp2869 {
		goto if_then2877
	} else {
		goto lor_lhs_false2871
	}

lor_lhs_false2871:
	v1259 = *libc.As[int32](lookahead)
	cmp2872 = 97 <= v1259
	if cmp2872 {
		goto land_lhs_true2874
	} else {
		goto if_end2878
	}

land_lhs_true2874:
	v1260 = *libc.As[int32](lookahead)
	cmp2875 = v1260 <= 122
	if cmp2875 {
		goto if_then2877
	} else {
		goto if_end2878
	}

if_then2877:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end2878:
	v1261 = *libc.As[byte](result)
	loadedv2879 = (v1261 & 1) != 0
	*libc.As[bool](retval) = loadedv2879
	goto _return

sw_bb2880:
	*libc.As[byte](result) = 1
	v1262 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2881 = libc.Ptr(&libc.As[TSLexer](v1262).F1)
	*libc.As[int16](result_symbol2881) = 14
	v1263 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2882 = libc.Ptr(&libc.As[TSLexer](v1263).F3)
	v1264 = *libc.As[unsafe.Pointer](mark_end2882)
	v1265 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1264)(v1265)
	v1266 = *libc.As[int32](lookahead)
	cmp2883 = 46 <= v1266
	if cmp2883 {
		goto land_lhs_true2885
	} else {
		goto lor_lhs_false2888
	}

land_lhs_true2885:
	v1267 = *libc.As[int32](lookahead)
	cmp2886 = v1267 <= 58
	if cmp2886 {
		goto if_then2903
	} else {
		goto lor_lhs_false2888
	}

lor_lhs_false2888:
	v1268 = *libc.As[int32](lookahead)
	cmp2889 = 65 <= v1268
	if cmp2889 {
		goto land_lhs_true2891
	} else {
		goto lor_lhs_false2894
	}

land_lhs_true2891:
	v1269 = *libc.As[int32](lookahead)
	cmp2892 = v1269 <= 90
	if cmp2892 {
		goto if_then2903
	} else {
		goto lor_lhs_false2894
	}

lor_lhs_false2894:
	v1270 = *libc.As[int32](lookahead)
	cmp2895 = v1270 == 95
	if cmp2895 {
		goto if_then2903
	} else {
		goto lor_lhs_false2897
	}

lor_lhs_false2897:
	v1271 = *libc.As[int32](lookahead)
	cmp2898 = 97 <= v1271
	if cmp2898 {
		goto land_lhs_true2900
	} else {
		goto if_end2904
	}

land_lhs_true2900:
	v1272 = *libc.As[int32](lookahead)
	cmp2901 = v1272 <= 122
	if cmp2901 {
		goto if_then2903
	} else {
		goto if_end2904
	}

if_then2903:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end2904:
	v1273 = *libc.As[byte](result)
	loadedv2905 = (v1273 & 1) != 0
	*libc.As[bool](retval) = loadedv2905
	goto _return

sw_bb2906:
	*libc.As[byte](result) = 1
	v1274 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2907 = libc.Ptr(&libc.As[TSLexer](v1274).F1)
	*libc.As[int16](result_symbol2907) = 15
	v1275 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2908 = libc.Ptr(&libc.As[TSLexer](v1275).F3)
	v1276 = *libc.As[unsafe.Pointer](mark_end2908)
	v1277 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1276)(v1277)
	v1278 = *libc.As[int32](lookahead)
	cmp2909 = v1278 == 46
	if cmp2909 {
		goto if_then2911
	} else {
		goto if_end2912
	}

if_then2911:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end2912:
	v1279 = *libc.As[int32](lookahead)
	cmp2913 = 48 <= v1279
	if cmp2913 {
		goto land_lhs_true2915
	} else {
		goto if_end2919
	}

land_lhs_true2915:
	v1280 = *libc.As[int32](lookahead)
	cmp2916 = v1280 <= 57
	if cmp2916 {
		goto if_then2918
	} else {
		goto if_end2919
	}

if_then2918:
	*libc.As[int16](state_addr) = 100
	goto next_state

if_end2919:
	v1281 = *libc.As[byte](result)
	loadedv2920 = (v1281 & 1) != 0
	*libc.As[bool](retval) = loadedv2920
	goto _return

sw_bb2921:
	*libc.As[byte](result) = 1
	v1282 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2922 = libc.Ptr(&libc.As[TSLexer](v1282).F1)
	*libc.As[int16](result_symbol2922) = 16
	v1283 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2923 = libc.Ptr(&libc.As[TSLexer](v1283).F3)
	v1284 = *libc.As[unsafe.Pointer](mark_end2923)
	v1285 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1284)(v1285)
	v1286 = *libc.As[int32](lookahead)
	cmp2924 = 48 <= v1286
	if cmp2924 {
		goto land_lhs_true2926
	} else {
		goto if_end2930
	}

land_lhs_true2926:
	v1287 = *libc.As[int32](lookahead)
	cmp2927 = v1287 <= 57
	if cmp2927 {
		goto if_then2929
	} else {
		goto if_end2930
	}

if_then2929:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end2930:
	v1288 = *libc.As[byte](result)
	loadedv2931 = (v1288 & 1) != 0
	*libc.As[bool](retval) = loadedv2931
	goto _return

sw_bb2932:
	*libc.As[byte](result) = 1
	v1289 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2933 = libc.Ptr(&libc.As[TSLexer](v1289).F1)
	*libc.As[int16](result_symbol2933) = 17
	v1290 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2934 = libc.Ptr(&libc.As[TSLexer](v1290).F3)
	v1291 = *libc.As[unsafe.Pointer](mark_end2934)
	v1292 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1291)(v1292)
	v1293 = *libc.As[int32](lookahead)
	cmp2935 = v1293 != 0
	if cmp2935 {
		goto land_lhs_true2937
	} else {
		goto if_end2941
	}

land_lhs_true2937:
	v1294 = *libc.As[int32](lookahead)
	cmp2938 = v1294 != 10
	if cmp2938 {
		goto if_then2940
	} else {
		goto if_end2941
	}

if_then2940:
	*libc.As[int16](state_addr) = 102
	goto next_state

if_end2941:
	v1295 = *libc.As[byte](result)
	loadedv2942 = (v1295 & 1) != 0
	*libc.As[bool](retval) = loadedv2942
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v1296 = *libc.As[bool](retval)
	return v1296
}
