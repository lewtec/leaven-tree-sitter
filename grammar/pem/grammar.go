package grammar_pem

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

var tree_sitter_pem_language struct {
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
var ts_small_parse_table [127]int16 = [127]int16{5, 5, 1, 6, 9, 1, 0, 11, 1, 7, 4, 1, 10, 3, 2, 9, 13, 5, 13, 1, 0, 15, 1, 6, 18, 1, 7, 4, 1, 10, 3, 2, 9, 13, 3, 21, 1, 4, 5, 1, 14, 9, 1, 12, 3, 23, 1, 4, 25, 1, 6, 7, 1, 14, 2, 27, 1, 0, 29, 2, 6, 7, 3, 31, 1, 4, 34, 1, 6, 7, 1, 14, 2, 36, 1, 0, 38, 2, 6, 7, 2, 40, 1, 6, 6, 1, 11, 1, 42, 1, 1, 1, 44, 1, 0, 1, 46, 1, 2, 1, 48, 1, 5, 1, 50, 1, 3, 1, 52, 1, 6, 1, 54, 1, 2, 1, 56, 1, 4, 1, 58, 1, 5, 1, 60, 1, 6}
var ts_small_parse_table_map [18]int32 = [18]int32{0, 17, 34, 44, 54, 62, 72, 80, 87, 91, 95, 99, 103, 107, 111, 115, 119, 123}
var ts_symbol_names [15]unsafe.Pointer = [15]unsafe.Pointer{libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16)}
var ts_symbol_metadata [15]TSSymbolMetadata = [15]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [15]int16 = [15]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][5]int16 = [1][5]int16{}
var ts_lex_modes [20]TSLexerMode = [20]TSLexerMode{TSLexerMode{}, TSLexerMode{18, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{6, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{}}
var ts_primary_state_ids [20]int16 = [20]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
var _str [4]byte = [4]byte{112, 101, 109, 0}
var ts_parse_table struct {
	F0 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 int16
		F7 [8]int16
	}
	F1 [15]int16
} = struct {
	F0 struct {
		F0 int16
		F1 int16
		F2 int16
		F3 int16
		F4 int16
		F5 int16
		F6 int16
		F7 [8]int16
	}
	F1 [15]int16
}{struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 int16
	F7 [8]int16
}{1, 1, 1, 1, 1, 0, 1, [8]int16{}}, [15]int16{3, 0, 0, 0, 0, 0, 5, 7, 11, 2, 4, 0, 0, 2, 0}}
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
	F30 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F37 TSParseActionEntry
	F38 struct {
		F0 anon_2
		F1 [6]byte
	}
	F39 TSParseActionEntry
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
			F0 byte
			F1 [7]byte
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
	F30 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F37 TSParseActionEntry
	F38 struct {
		F0 anon_2
		F1 [6]byte
	}
	F39 TSParseActionEntry
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
			F0 byte
			F1 [7]byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 8, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 8, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 13, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 13, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 10, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 13, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{0, 0, 7, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 12, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 9, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 9, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 14, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 14, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 11, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 11, 0, 0}}}, struct {
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
}{0, 0, 12, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 18, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 10, 0, 0}}}, struct {
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
}{0, 0, 8, 0, 0}, [2]byte{}}}}
var _str_3 [4]byte = [4]byte{101, 110, 100, 0}
var _str_4 [6]byte = [6]byte{66, 69, 71, 73, 78, 0}
var _str_5 [2]byte = [2]byte{32, 0}
var _str_6 [4]byte = [4]byte{69, 78, 68, 0}
var _str_7 [12]byte = [12]byte{100, 97, 116, 97, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_8 [6]byte = [6]byte{108, 97, 98, 101, 108, 0}
var _str_9 [7]byte = [7]byte{100, 97, 115, 104, 101, 115, 0}
var _str_10 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_11 [8]byte = [8]byte{99, 111, 110, 116, 101, 110, 116, 0}
var _str_12 [7]byte = [7]byte{104, 101, 97, 100, 101, 114, 0}
var _str_13 [7]byte = [7]byte{102, 111, 111, 116, 101, 114, 0}
var _str_14 [5]byte = [5]byte{100, 97, 116, 97, 0}
var _str_15 [12]byte = [12]byte{112, 101, 109, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_16 [13]byte = [13]byte{100, 97, 116, 97, 95, 114, 101, 112, 101, 97, 116, 49, 0}

func init() {
	tree_sitter_pem_language = struct {
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
	}{15, 15, 0, 8, 0, 20, 2, 1, 0, 5, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), libc.Ptr(&_str), nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{0, 3, 0}, [5]byte{}}
}
func tree_sitter_pem() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_pem_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp27, cmp29, cmp31, cmp34, cmp37, cmp40, cmp43, loadedv47, cmp49, cmp53, cmp57, cmp61, loadedv65, cmp67, loadedv71, cmp73, cmp77, cmp81, cmp85, cmp88, cmp91, cmp94, cmp97, cmp100, cmp103, loadedv107, cmp109, loadedv113, cmp115, cmp119, cmp123, cmp127, loadedv131, cmp133, cmp137, cmp141, cmp144, loadedv148, cmp150, loadedv154, cmp156, loadedv160, cmp162, loadedv166, cmp168, loadedv172, cmp174, loadedv178, cmp180, loadedv184, cmp186, loadedv190, cmp192, loadedv196, cmp198, loadedv202, cmp204, loadedv208, loadedv210, cmp213, loadedv217, loadedv219, cmp222, cmp226, cmp230, cmp234, loadedv238, loadedv240, loadedv244, cmp248, cmp252, cmp255, cmp258, cmp261, cmp264, cmp267, cmp270, loadedv274, loadedv278, loadedv282, cmp286, cmp290, cmp293, cmp296, cmp299, cmp302, cmp305, cmp308, loadedv312, loadedv316, cmp320, cmp324, cmp328, cmp331, cmp334, cmp337, cmp340, cmp343, cmp346, loadedv350, cmp354, cmp358, cmp362, cmp365, cmp368, cmp371, cmp374, cmp377, cmp380, loadedv384, cmp388, cmp392, cmp396, cmp399, cmp402, cmp405, cmp408, cmp411, cmp414, loadedv418, cmp422, cmp426, cmp430, cmp433, cmp436, cmp439, cmp442, cmp445, cmp448, loadedv452, cmp456, cmp460, cmp464, cmp467, cmp470, cmp473, cmp476, cmp479, cmp482, loadedv486, cmp490, cmp494, cmp498, cmp501, cmp504, cmp507, cmp510, cmp513, cmp516, loadedv520, cmp524, cmp528, cmp531, cmp534, cmp537, cmp540, cmp543, cmp546, loadedv550, cmp554, loadedv558, cmp562, cmp566, cmp570, cmp573, loadedv577, cmp581, cmp585, cmp588, loadedv592, cmp596, cmp599, loadedv603, loadedv607, cmp611, cmp615, loadedv619, cmp623, cmp627, cmp630, loadedv634, cmp638, cmp642, cmp645, loadedv649, cmp653, cmp657, cmp660, loadedv664, cmp668, cmp672, cmp675, loadedv679, cmp683, cmp686, loadedv690, v313 bool
	var retval unsafe.Pointer
	var v9 int16
	var state_addr, result_symbol, result_symbol242, result_symbol246, result_symbol276, result_symbol280, result_symbol284, result_symbol314, result_symbol318, result_symbol352, result_symbol386, result_symbol420, result_symbol454, result_symbol488, result_symbol522, result_symbol552, result_symbol560, result_symbol579, result_symbol594, result_symbol605, result_symbol609, result_symbol621, result_symbol636, result_symbol651, result_symbol666, result_symbol681 unsafe.Pointer
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v25, v26, v27, v28, v30, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v43, v45, v46, v47, v48, v50, v51, v52, v53, v55, v57, v59, v61, v63, v65, v67, v69, v71, v73, v76, v79, v80, v81, v82, v98, v99, v100, v101, v102, v103, v104, v105, v121, v122, v123, v124, v125, v126, v127, v128, v139, v140, v141, v142, v143, v144, v145, v146, v147, v153, v154, v155, v156, v157, v158, v159, v160, v161, v167, v168, v169, v170, v171, v172, v173, v174, v175, v181, v182, v183, v184, v185, v186, v187, v188, v189, v195, v196, v197, v198, v199, v200, v201, v202, v203, v209, v210, v211, v212, v213, v214, v215, v216, v217, v223, v224, v225, v226, v227, v228, v229, v230, v236, v242, v243, v244, v245, v251, v252, v253, v259, v260, v271, v272, v278, v279, v280, v286, v287, v288, v294, v295, v296, v302, v303, v304, v310, v311 int32
	var lookahead, lookahead1 unsafe.Pointer
	var v3, storedv, v10, v24, v29, v31, v42, v44, v49, v54, v56, v58, v60, v62, v64, v66, v68, v70, v72, v74, v75, v77, v78, v83, v88, v93, v106, v111, v116, v129, v134, v148, v162, v176, v190, v204, v218, v231, v237, v246, v254, v261, v266, v273, v281, v289, v297, v305, v312 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v84, v85, v86, v87, v89, v90, v91, v92, v94, v95, v96, v97, v107, v108, v109, v110, v112, v113, v114, v115, v117, v118, v119, v120, v130, v131, v132, v133, v135, v136, v137, v138, v149, v150, v151, v152, v163, v164, v165, v166, v177, v178, v179, v180, v191, v192, v193, v194, v205, v206, v207, v208, v219, v220, v221, v222, v232, v233, v234, v235, v238, v239, v240, v241, v247, v248, v249, v250, v255, v256, v257, v258, v262, v263, v264, v265, v267, v268, v269, v270, v274, v275, v276, v277, v282, v283, v284, v285, v290, v291, v292, v293, v298, v299, v300, v301, v306, v307, v308, v309 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end243, mark_end247, mark_end277, mark_end281, mark_end285, mark_end315, mark_end319, mark_end353, mark_end387, mark_end421, mark_end455, mark_end489, mark_end523, mark_end553, mark_end561, mark_end580, mark_end595, mark_end606, mark_end610, mark_end622, mark_end637, mark_end652, mark_end667, mark_end682 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp27, v18, cmp29, v19, cmp31, v20, cmp34, v21, cmp37, v22, cmp40, v23, cmp43, v24, loadedv47, v25, cmp49, v26, cmp53, v27, cmp57, v28, cmp61, v29, loadedv65, v30, cmp67, v31, loadedv71, v32, cmp73, v33, cmp77, v34, cmp81, v35, cmp85, v36, cmp88, v37, cmp91, v38, cmp94, v39, cmp97, v40, cmp100, v41, cmp103, v42, loadedv107, v43, cmp109, v44, loadedv113, v45, cmp115, v46, cmp119, v47, cmp123, v48, cmp127, v49, loadedv131, v50, cmp133, v51, cmp137, v52, cmp141, v53, cmp144, v54, loadedv148, v55, cmp150, v56, loadedv154, v57, cmp156, v58, loadedv160, v59, cmp162, v60, loadedv166, v61, cmp168, v62, loadedv172, v63, cmp174, v64, loadedv178, v65, cmp180, v66, loadedv184, v67, cmp186, v68, loadedv190, v69, cmp192, v70, loadedv196, v71, cmp198, v72, loadedv202, v73, cmp204, v74, loadedv208, v75, loadedv210, v76, cmp213, v77, loadedv217, v78, loadedv219, v79, cmp222, v80, cmp226, v81, cmp230, v82, cmp234, v83, loadedv238, v84, result_symbol, v85, mark_end, v86, v87, v88, loadedv240, v89, result_symbol242, v90, mark_end243, v91, v92, v93, loadedv244, v94, result_symbol246, v95, mark_end247, v96, v97, v98, cmp248, v99, cmp252, v100, cmp255, v101, cmp258, v102, cmp261, v103, cmp264, v104, cmp267, v105, cmp270, v106, loadedv274, v107, result_symbol276, v108, mark_end277, v109, v110, v111, loadedv278, v112, result_symbol280, v113, mark_end281, v114, v115, v116, loadedv282, v117, result_symbol284, v118, mark_end285, v119, v120, v121, cmp286, v122, cmp290, v123, cmp293, v124, cmp296, v125, cmp299, v126, cmp302, v127, cmp305, v128, cmp308, v129, loadedv312, v130, result_symbol314, v131, mark_end315, v132, v133, v134, loadedv316, v135, result_symbol318, v136, mark_end319, v137, v138, v139, cmp320, v140, cmp324, v141, cmp328, v142, cmp331, v143, cmp334, v144, cmp337, v145, cmp340, v146, cmp343, v147, cmp346, v148, loadedv350, v149, result_symbol352, v150, mark_end353, v151, v152, v153, cmp354, v154, cmp358, v155, cmp362, v156, cmp365, v157, cmp368, v158, cmp371, v159, cmp374, v160, cmp377, v161, cmp380, v162, loadedv384, v163, result_symbol386, v164, mark_end387, v165, v166, v167, cmp388, v168, cmp392, v169, cmp396, v170, cmp399, v171, cmp402, v172, cmp405, v173, cmp408, v174, cmp411, v175, cmp414, v176, loadedv418, v177, result_symbol420, v178, mark_end421, v179, v180, v181, cmp422, v182, cmp426, v183, cmp430, v184, cmp433, v185, cmp436, v186, cmp439, v187, cmp442, v188, cmp445, v189, cmp448, v190, loadedv452, v191, result_symbol454, v192, mark_end455, v193, v194, v195, cmp456, v196, cmp460, v197, cmp464, v198, cmp467, v199, cmp470, v200, cmp473, v201, cmp476, v202, cmp479, v203, cmp482, v204, loadedv486, v205, result_symbol488, v206, mark_end489, v207, v208, v209, cmp490, v210, cmp494, v211, cmp498, v212, cmp501, v213, cmp504, v214, cmp507, v215, cmp510, v216, cmp513, v217, cmp516, v218, loadedv520, v219, result_symbol522, v220, mark_end523, v221, v222, v223, cmp524, v224, cmp528, v225, cmp531, v226, cmp534, v227, cmp537, v228, cmp540, v229, cmp543, v230, cmp546, v231, loadedv550, v232, result_symbol552, v233, mark_end553, v234, v235, v236, cmp554, v237, loadedv558, v238, result_symbol560, v239, mark_end561, v240, v241, v242, cmp562, v243, cmp566, v244, cmp570, v245, cmp573, v246, loadedv577, v247, result_symbol579, v248, mark_end580, v249, v250, v251, cmp581, v252, cmp585, v253, cmp588, v254, loadedv592, v255, result_symbol594, v256, mark_end595, v257, v258, v259, cmp596, v260, cmp599, v261, loadedv603, v262, result_symbol605, v263, mark_end606, v264, v265, v266, loadedv607, v267, result_symbol609, v268, mark_end610, v269, v270, v271, cmp611, v272, cmp615, v273, loadedv619, v274, result_symbol621, v275, mark_end622, v276, v277, v278, cmp623, v279, cmp627, v280, cmp630, v281, loadedv634, v282, result_symbol636, v283, mark_end637, v284, v285, v286, cmp638, v287, cmp642, v288, cmp645, v289, loadedv649, v290, result_symbol651, v291, mark_end652, v292, v293, v294, cmp653, v295, cmp657, v296, cmp660, v297, loadedv664, v298, result_symbol666, v299, mark_end667, v300, v301, v302, cmp668, v303, cmp672, v304, cmp675, v305, loadedv679, v306, result_symbol681, v307, mark_end682, v308, v309, v310, cmp683, v311, cmp686, v312, loadedv690, v313

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
		goto sw_bb48
	case 2:
		goto sw_bb66
	case 3:
		goto sw_bb72
	case 4:
		goto sw_bb108
	case 5:
		goto sw_bb114
	case 6:
		goto sw_bb132
	case 7:
		goto sw_bb149
	case 8:
		goto sw_bb155
	case 9:
		goto sw_bb161
	case 10:
		goto sw_bb167
	case 11:
		goto sw_bb173
	case 12:
		goto sw_bb179
	case 13:
		goto sw_bb185
	case 14:
		goto sw_bb191
	case 15:
		goto sw_bb197
	case 16:
		goto sw_bb203
	case 17:
		goto sw_bb209
	case 18:
		goto sw_bb218
	case 19:
		goto sw_bb239
	case 20:
		goto sw_bb241
	case 21:
		goto sw_bb245
	case 22:
		goto sw_bb275
	case 23:
		goto sw_bb279
	case 24:
		goto sw_bb283
	case 25:
		goto sw_bb313
	case 26:
		goto sw_bb317
	case 27:
		goto sw_bb351
	case 28:
		goto sw_bb385
	case 29:
		goto sw_bb419
	case 30:
		goto sw_bb453
	case 31:
		goto sw_bb487
	case 32:
		goto sw_bb521
	case 33:
		goto sw_bb551
	case 34:
		goto sw_bb559
	case 35:
		goto sw_bb578
	case 36:
		goto sw_bb593
	case 37:
		goto sw_bb604
	case 38:
		goto sw_bb608
	case 39:
		goto sw_bb620
	case 40:
		goto sw_bb635
	case 41:
		goto sw_bb650
	case 42:
		goto sw_bb665
	case 43:
		goto sw_bb680
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
	*libc.As[int16](state_addr) = 19
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
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
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
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end10:
	v13 = *libc.As[int32](lookahead)
	cmp11 = v13 == 32
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end14:
	v14 = *libc.As[int32](lookahead)
	cmp15 = v14 == 45
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end18:
	v15 = *libc.As[int32](lookahead)
	cmp19 = v15 == 66
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end22:
	v16 = *libc.As[int32](lookahead)
	cmp23 = v16 == 69
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end26:
	v17 = *libc.As[int32](lookahead)
	cmp27 = v17 == 43
	if cmp27 {
		goto if_then45
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v18 = *libc.As[int32](lookahead)
	cmp29 = 47 <= v18
	if cmp29 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false33
	}

land_lhs_true:
	v19 = *libc.As[int32](lookahead)
	cmp31 = v19 <= 57
	if cmp31 {
		goto if_then45
	} else {
		goto lor_lhs_false33
	}

lor_lhs_false33:
	v20 = *libc.As[int32](lookahead)
	cmp34 = 65 <= v20
	if cmp34 {
		goto land_lhs_true36
	} else {
		goto lor_lhs_false39
	}

land_lhs_true36:
	v21 = *libc.As[int32](lookahead)
	cmp37 = v21 <= 90
	if cmp37 {
		goto if_then45
	} else {
		goto lor_lhs_false39
	}

lor_lhs_false39:
	v22 = *libc.As[int32](lookahead)
	cmp40 = 97 <= v22
	if cmp40 {
		goto land_lhs_true42
	} else {
		goto if_end46
	}

land_lhs_true42:
	v23 = *libc.As[int32](lookahead)
	cmp43 = v23 <= 122
	if cmp43 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end46:
	v24 = *libc.As[byte](result)
	loadedv47 = (v24 & 1) != 0
	*libc.As[bool](retval) = loadedv47
	goto _return

sw_bb48:
	v25 = *libc.As[int32](lookahead)
	cmp49 = v25 == 10
	if cmp49 {
		goto if_then51
	} else {
		goto if_end52
	}

if_then51:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end52:
	v26 = *libc.As[int32](lookahead)
	cmp53 = v26 == 13
	if cmp53 {
		goto if_then55
	} else {
		goto if_end56
	}

if_then55:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end56:
	v27 = *libc.As[int32](lookahead)
	cmp57 = v27 == 45
	if cmp57 {
		goto if_then59
	} else {
		goto if_end60
	}

if_then59:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end60:
	v28 = *libc.As[int32](lookahead)
	cmp61 = v28 != 0
	if cmp61 {
		goto if_then63
	} else {
		goto if_end64
	}

if_then63:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end64:
	v29 = *libc.As[byte](result)
	loadedv65 = (v29 & 1) != 0
	*libc.As[bool](retval) = loadedv65
	goto _return

sw_bb66:
	v30 = *libc.As[int32](lookahead)
	cmp67 = v30 == 10
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end70:
	v31 = *libc.As[byte](result)
	loadedv71 = (v31 & 1) != 0
	*libc.As[bool](retval) = loadedv71
	goto _return

sw_bb72:
	v32 = *libc.As[int32](lookahead)
	cmp73 = v32 == 10
	if cmp73 {
		goto if_then75
	} else {
		goto if_end76
	}

if_then75:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end76:
	v33 = *libc.As[int32](lookahead)
	cmp77 = v33 == 13
	if cmp77 {
		goto if_then79
	} else {
		goto if_end80
	}

if_then79:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end80:
	v34 = *libc.As[int32](lookahead)
	cmp81 = v34 == 45
	if cmp81 {
		goto if_then83
	} else {
		goto if_end84
	}

if_then83:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end84:
	v35 = *libc.As[int32](lookahead)
	cmp85 = v35 == 43
	if cmp85 {
		goto if_then105
	} else {
		goto lor_lhs_false87
	}

lor_lhs_false87:
	v36 = *libc.As[int32](lookahead)
	cmp88 = 47 <= v36
	if cmp88 {
		goto land_lhs_true90
	} else {
		goto lor_lhs_false93
	}

land_lhs_true90:
	v37 = *libc.As[int32](lookahead)
	cmp91 = v37 <= 57
	if cmp91 {
		goto if_then105
	} else {
		goto lor_lhs_false93
	}

lor_lhs_false93:
	v38 = *libc.As[int32](lookahead)
	cmp94 = 65 <= v38
	if cmp94 {
		goto land_lhs_true96
	} else {
		goto lor_lhs_false99
	}

land_lhs_true96:
	v39 = *libc.As[int32](lookahead)
	cmp97 = v39 <= 90
	if cmp97 {
		goto if_then105
	} else {
		goto lor_lhs_false99
	}

lor_lhs_false99:
	v40 = *libc.As[int32](lookahead)
	cmp100 = 97 <= v40
	if cmp100 {
		goto land_lhs_true102
	} else {
		goto if_end106
	}

land_lhs_true102:
	v41 = *libc.As[int32](lookahead)
	cmp103 = v41 <= 122
	if cmp103 {
		goto if_then105
	} else {
		goto if_end106
	}

if_then105:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end106:
	v42 = *libc.As[byte](result)
	loadedv107 = (v42 & 1) != 0
	*libc.As[bool](retval) = loadedv107
	goto _return

sw_bb108:
	v43 = *libc.As[int32](lookahead)
	cmp109 = v43 == 10
	if cmp109 {
		goto if_then111
	} else {
		goto if_end112
	}

if_then111:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end112:
	v44 = *libc.As[byte](result)
	loadedv113 = (v44 & 1) != 0
	*libc.As[bool](retval) = loadedv113
	goto _return

sw_bb114:
	v45 = *libc.As[int32](lookahead)
	cmp115 = v45 == 10
	if cmp115 {
		goto if_then117
	} else {
		goto if_end118
	}

if_then117:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end118:
	v46 = *libc.As[int32](lookahead)
	cmp119 = v46 == 13
	if cmp119 {
		goto if_then121
	} else {
		goto if_end122
	}

if_then121:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end122:
	v47 = *libc.As[int32](lookahead)
	cmp123 = v47 == 66
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end126:
	v48 = *libc.As[int32](lookahead)
	cmp127 = v48 == 69
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end130:
	v49 = *libc.As[byte](result)
	loadedv131 = (v49 & 1) != 0
	*libc.As[bool](retval) = loadedv131
	goto _return

sw_bb132:
	v50 = *libc.As[int32](lookahead)
	cmp133 = v50 == 10
	if cmp133 {
		goto if_then135
	} else {
		goto if_end136
	}

if_then135:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end136:
	v51 = *libc.As[int32](lookahead)
	cmp137 = v51 == 13
	if cmp137 {
		goto if_then139
	} else {
		goto if_end140
	}

if_then139:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end140:
	v52 = *libc.As[int32](lookahead)
	cmp141 = v52 != 0
	if cmp141 {
		goto land_lhs_true143
	} else {
		goto if_end147
	}

land_lhs_true143:
	v53 = *libc.As[int32](lookahead)
	cmp144 = v53 != 45
	if cmp144 {
		goto if_then146
	} else {
		goto if_end147
	}

if_then146:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end147:
	v54 = *libc.As[byte](result)
	loadedv148 = (v54 & 1) != 0
	*libc.As[bool](retval) = loadedv148
	goto _return

sw_bb149:
	v55 = *libc.As[int32](lookahead)
	cmp150 = v55 == 45
	if cmp150 {
		goto if_then152
	} else {
		goto if_end153
	}

if_then152:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end153:
	v56 = *libc.As[byte](result)
	loadedv154 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv154
	goto _return

sw_bb155:
	v57 = *libc.As[int32](lookahead)
	cmp156 = v57 == 45
	if cmp156 {
		goto if_then158
	} else {
		goto if_end159
	}

if_then158:
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end159:
	v58 = *libc.As[byte](result)
	loadedv160 = (v58 & 1) != 0
	*libc.As[bool](retval) = loadedv160
	goto _return

sw_bb161:
	v59 = *libc.As[int32](lookahead)
	cmp162 = v59 == 45
	if cmp162 {
		goto if_then164
	} else {
		goto if_end165
	}

if_then164:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end165:
	v60 = *libc.As[byte](result)
	loadedv166 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv166
	goto _return

sw_bb167:
	v61 = *libc.As[int32](lookahead)
	cmp168 = v61 == 45
	if cmp168 {
		goto if_then170
	} else {
		goto if_end171
	}

if_then170:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end171:
	v62 = *libc.As[byte](result)
	loadedv172 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv172
	goto _return

sw_bb173:
	v63 = *libc.As[int32](lookahead)
	cmp174 = v63 == 68
	if cmp174 {
		goto if_then176
	} else {
		goto if_end177
	}

if_then176:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end177:
	v64 = *libc.As[byte](result)
	loadedv178 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv178
	goto _return

sw_bb179:
	v65 = *libc.As[int32](lookahead)
	cmp180 = v65 == 69
	if cmp180 {
		goto if_then182
	} else {
		goto if_end183
	}

if_then182:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end183:
	v66 = *libc.As[byte](result)
	loadedv184 = (v66 & 1) != 0
	*libc.As[bool](retval) = loadedv184
	goto _return

sw_bb185:
	v67 = *libc.As[int32](lookahead)
	cmp186 = v67 == 71
	if cmp186 {
		goto if_then188
	} else {
		goto if_end189
	}

if_then188:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end189:
	v68 = *libc.As[byte](result)
	loadedv190 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv190
	goto _return

sw_bb191:
	v69 = *libc.As[int32](lookahead)
	cmp192 = v69 == 73
	if cmp192 {
		goto if_then194
	} else {
		goto if_end195
	}

if_then194:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end195:
	v70 = *libc.As[byte](result)
	loadedv196 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv196
	goto _return

sw_bb197:
	v71 = *libc.As[int32](lookahead)
	cmp198 = v71 == 78
	if cmp198 {
		goto if_then200
	} else {
		goto if_end201
	}

if_then200:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end201:
	v72 = *libc.As[byte](result)
	loadedv202 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv202
	goto _return

sw_bb203:
	v73 = *libc.As[int32](lookahead)
	cmp204 = v73 == 78
	if cmp204 {
		goto if_then206
	} else {
		goto if_end207
	}

if_then206:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end207:
	v74 = *libc.As[byte](result)
	loadedv208 = (v74 & 1) != 0
	*libc.As[bool](retval) = loadedv208
	goto _return

sw_bb209:
	v75 = *libc.As[byte](eof)
	loadedv210 = (v75 & 1) != 0
	if loadedv210 {
		goto if_then211
	} else {
		goto if_end212
	}

if_then211:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end212:
	v76 = *libc.As[int32](lookahead)
	cmp213 = v76 == 10
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end216:
	v77 = *libc.As[byte](result)
	loadedv217 = (v77 & 1) != 0
	*libc.As[bool](retval) = loadedv217
	goto _return

sw_bb218:
	v78 = *libc.As[byte](eof)
	loadedv219 = (v78 & 1) != 0
	if loadedv219 {
		goto if_then220
	} else {
		goto if_end221
	}

if_then220:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end221:
	v79 = *libc.As[int32](lookahead)
	cmp222 = v79 == 10
	if cmp222 {
		goto if_then224
	} else {
		goto if_end225
	}

if_then224:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end225:
	v80 = *libc.As[int32](lookahead)
	cmp226 = v80 == 13
	if cmp226 {
		goto if_then228
	} else {
		goto if_end229
	}

if_then228:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end229:
	v81 = *libc.As[int32](lookahead)
	cmp230 = v81 == 45
	if cmp230 {
		goto if_then232
	} else {
		goto if_end233
	}

if_then232:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end233:
	v82 = *libc.As[int32](lookahead)
	cmp234 = v82 != 0
	if cmp234 {
		goto if_then236
	} else {
		goto if_end237
	}

if_then236:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end237:
	v83 = *libc.As[byte](result)
	loadedv238 = (v83 & 1) != 0
	*libc.As[bool](retval) = loadedv238
	goto _return

sw_bb239:
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
	loadedv240 = (v88 & 1) != 0
	*libc.As[bool](retval) = loadedv240
	goto _return

sw_bb241:
	*libc.As[byte](result) = 1
	v89 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol242 = libc.Ptr(&libc.As[TSLexer](v89).F1)
	*libc.As[int16](result_symbol242) = 1
	v90 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end243 = libc.Ptr(&libc.As[TSLexer](v90).F3)
	v91 = *libc.As[unsafe.Pointer](mark_end243)
	v92 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v91)(v92)
	v93 = *libc.As[byte](result)
	loadedv244 = (v93 & 1) != 0
	*libc.As[bool](retval) = loadedv244
	goto _return

sw_bb245:
	*libc.As[byte](result) = 1
	v94 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol246 = libc.Ptr(&libc.As[TSLexer](v94).F1)
	*libc.As[int16](result_symbol246) = 1
	v95 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end247 = libc.Ptr(&libc.As[TSLexer](v95).F3)
	v96 = *libc.As[unsafe.Pointer](mark_end247)
	v97 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v96)(v97)
	v98 = *libc.As[int32](lookahead)
	cmp248 = v98 == 61
	if cmp248 {
		goto if_then250
	} else {
		goto if_end251
	}

if_then250:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end251:
	v99 = *libc.As[int32](lookahead)
	cmp252 = v99 == 43
	if cmp252 {
		goto if_then272
	} else {
		goto lor_lhs_false254
	}

lor_lhs_false254:
	v100 = *libc.As[int32](lookahead)
	cmp255 = 47 <= v100
	if cmp255 {
		goto land_lhs_true257
	} else {
		goto lor_lhs_false260
	}

land_lhs_true257:
	v101 = *libc.As[int32](lookahead)
	cmp258 = v101 <= 57
	if cmp258 {
		goto if_then272
	} else {
		goto lor_lhs_false260
	}

lor_lhs_false260:
	v102 = *libc.As[int32](lookahead)
	cmp261 = 65 <= v102
	if cmp261 {
		goto land_lhs_true263
	} else {
		goto lor_lhs_false266
	}

land_lhs_true263:
	v103 = *libc.As[int32](lookahead)
	cmp264 = v103 <= 90
	if cmp264 {
		goto if_then272
	} else {
		goto lor_lhs_false266
	}

lor_lhs_false266:
	v104 = *libc.As[int32](lookahead)
	cmp267 = 97 <= v104
	if cmp267 {
		goto land_lhs_true269
	} else {
		goto if_end273
	}

land_lhs_true269:
	v105 = *libc.As[int32](lookahead)
	cmp270 = v105 <= 122
	if cmp270 {
		goto if_then272
	} else {
		goto if_end273
	}

if_then272:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end273:
	v106 = *libc.As[byte](result)
	loadedv274 = (v106 & 1) != 0
	*libc.As[bool](retval) = loadedv274
	goto _return

sw_bb275:
	*libc.As[byte](result) = 1
	v107 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol276 = libc.Ptr(&libc.As[TSLexer](v107).F1)
	*libc.As[int16](result_symbol276) = 2
	v108 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end277 = libc.Ptr(&libc.As[TSLexer](v108).F3)
	v109 = *libc.As[unsafe.Pointer](mark_end277)
	v110 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v109)(v110)
	v111 = *libc.As[byte](result)
	loadedv278 = (v111 & 1) != 0
	*libc.As[bool](retval) = loadedv278
	goto _return

sw_bb279:
	*libc.As[byte](result) = 1
	v112 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol280 = libc.Ptr(&libc.As[TSLexer](v112).F1)
	*libc.As[int16](result_symbol280) = 3
	v113 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end281 = libc.Ptr(&libc.As[TSLexer](v113).F3)
	v114 = *libc.As[unsafe.Pointer](mark_end281)
	v115 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v114)(v115)
	v116 = *libc.As[byte](result)
	loadedv282 = (v116 & 1) != 0
	*libc.As[bool](retval) = loadedv282
	goto _return

sw_bb283:
	*libc.As[byte](result) = 1
	v117 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol284 = libc.Ptr(&libc.As[TSLexer](v117).F1)
	*libc.As[int16](result_symbol284) = 3
	v118 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end285 = libc.Ptr(&libc.As[TSLexer](v118).F3)
	v119 = *libc.As[unsafe.Pointer](mark_end285)
	v120 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v119)(v120)
	v121 = *libc.As[int32](lookahead)
	cmp286 = v121 == 61
	if cmp286 {
		goto if_then288
	} else {
		goto if_end289
	}

if_then288:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end289:
	v122 = *libc.As[int32](lookahead)
	cmp290 = v122 == 43
	if cmp290 {
		goto if_then310
	} else {
		goto lor_lhs_false292
	}

lor_lhs_false292:
	v123 = *libc.As[int32](lookahead)
	cmp293 = 47 <= v123
	if cmp293 {
		goto land_lhs_true295
	} else {
		goto lor_lhs_false298
	}

land_lhs_true295:
	v124 = *libc.As[int32](lookahead)
	cmp296 = v124 <= 57
	if cmp296 {
		goto if_then310
	} else {
		goto lor_lhs_false298
	}

lor_lhs_false298:
	v125 = *libc.As[int32](lookahead)
	cmp299 = 65 <= v125
	if cmp299 {
		goto land_lhs_true301
	} else {
		goto lor_lhs_false304
	}

land_lhs_true301:
	v126 = *libc.As[int32](lookahead)
	cmp302 = v126 <= 90
	if cmp302 {
		goto if_then310
	} else {
		goto lor_lhs_false304
	}

lor_lhs_false304:
	v127 = *libc.As[int32](lookahead)
	cmp305 = 97 <= v127
	if cmp305 {
		goto land_lhs_true307
	} else {
		goto if_end311
	}

land_lhs_true307:
	v128 = *libc.As[int32](lookahead)
	cmp308 = v128 <= 122
	if cmp308 {
		goto if_then310
	} else {
		goto if_end311
	}

if_then310:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end311:
	v129 = *libc.As[byte](result)
	loadedv312 = (v129 & 1) != 0
	*libc.As[bool](retval) = loadedv312
	goto _return

sw_bb313:
	*libc.As[byte](result) = 1
	v130 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol314 = libc.Ptr(&libc.As[TSLexer](v130).F1)
	*libc.As[int16](result_symbol314) = 4
	v131 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end315 = libc.Ptr(&libc.As[TSLexer](v131).F3)
	v132 = *libc.As[unsafe.Pointer](mark_end315)
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v132)(v133)
	v134 = *libc.As[byte](result)
	loadedv316 = (v134 & 1) != 0
	*libc.As[bool](retval) = loadedv316
	goto _return

sw_bb317:
	*libc.As[byte](result) = 1
	v135 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol318 = libc.Ptr(&libc.As[TSLexer](v135).F1)
	*libc.As[int16](result_symbol318) = 4
	v136 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end319 = libc.Ptr(&libc.As[TSLexer](v136).F3)
	v137 = *libc.As[unsafe.Pointer](mark_end319)
	v138 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v137)(v138)
	v139 = *libc.As[int32](lookahead)
	cmp320 = v139 == 61
	if cmp320 {
		goto if_then322
	} else {
		goto if_end323
	}

if_then322:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end323:
	v140 = *libc.As[int32](lookahead)
	cmp324 = v140 == 68
	if cmp324 {
		goto if_then326
	} else {
		goto if_end327
	}

if_then326:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end327:
	v141 = *libc.As[int32](lookahead)
	cmp328 = v141 == 43
	if cmp328 {
		goto if_then348
	} else {
		goto lor_lhs_false330
	}

lor_lhs_false330:
	v142 = *libc.As[int32](lookahead)
	cmp331 = 47 <= v142
	if cmp331 {
		goto land_lhs_true333
	} else {
		goto lor_lhs_false336
	}

land_lhs_true333:
	v143 = *libc.As[int32](lookahead)
	cmp334 = v143 <= 57
	if cmp334 {
		goto if_then348
	} else {
		goto lor_lhs_false336
	}

lor_lhs_false336:
	v144 = *libc.As[int32](lookahead)
	cmp337 = 65 <= v144
	if cmp337 {
		goto land_lhs_true339
	} else {
		goto lor_lhs_false342
	}

land_lhs_true339:
	v145 = *libc.As[int32](lookahead)
	cmp340 = v145 <= 90
	if cmp340 {
		goto if_then348
	} else {
		goto lor_lhs_false342
	}

lor_lhs_false342:
	v146 = *libc.As[int32](lookahead)
	cmp343 = 97 <= v146
	if cmp343 {
		goto land_lhs_true345
	} else {
		goto if_end349
	}

land_lhs_true345:
	v147 = *libc.As[int32](lookahead)
	cmp346 = v147 <= 122
	if cmp346 {
		goto if_then348
	} else {
		goto if_end349
	}

if_then348:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end349:
	v148 = *libc.As[byte](result)
	loadedv350 = (v148 & 1) != 0
	*libc.As[bool](retval) = loadedv350
	goto _return

sw_bb351:
	*libc.As[byte](result) = 1
	v149 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol352 = libc.Ptr(&libc.As[TSLexer](v149).F1)
	*libc.As[int16](result_symbol352) = 4
	v150 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end353 = libc.Ptr(&libc.As[TSLexer](v150).F3)
	v151 = *libc.As[unsafe.Pointer](mark_end353)
	v152 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v151)(v152)
	v153 = *libc.As[int32](lookahead)
	cmp354 = v153 == 61
	if cmp354 {
		goto if_then356
	} else {
		goto if_end357
	}

if_then356:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end357:
	v154 = *libc.As[int32](lookahead)
	cmp358 = v154 == 69
	if cmp358 {
		goto if_then360
	} else {
		goto if_end361
	}

if_then360:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end361:
	v155 = *libc.As[int32](lookahead)
	cmp362 = v155 == 43
	if cmp362 {
		goto if_then382
	} else {
		goto lor_lhs_false364
	}

lor_lhs_false364:
	v156 = *libc.As[int32](lookahead)
	cmp365 = 47 <= v156
	if cmp365 {
		goto land_lhs_true367
	} else {
		goto lor_lhs_false370
	}

land_lhs_true367:
	v157 = *libc.As[int32](lookahead)
	cmp368 = v157 <= 57
	if cmp368 {
		goto if_then382
	} else {
		goto lor_lhs_false370
	}

lor_lhs_false370:
	v158 = *libc.As[int32](lookahead)
	cmp371 = 65 <= v158
	if cmp371 {
		goto land_lhs_true373
	} else {
		goto lor_lhs_false376
	}

land_lhs_true373:
	v159 = *libc.As[int32](lookahead)
	cmp374 = v159 <= 90
	if cmp374 {
		goto if_then382
	} else {
		goto lor_lhs_false376
	}

lor_lhs_false376:
	v160 = *libc.As[int32](lookahead)
	cmp377 = 97 <= v160
	if cmp377 {
		goto land_lhs_true379
	} else {
		goto if_end383
	}

land_lhs_true379:
	v161 = *libc.As[int32](lookahead)
	cmp380 = v161 <= 122
	if cmp380 {
		goto if_then382
	} else {
		goto if_end383
	}

if_then382:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end383:
	v162 = *libc.As[byte](result)
	loadedv384 = (v162 & 1) != 0
	*libc.As[bool](retval) = loadedv384
	goto _return

sw_bb385:
	*libc.As[byte](result) = 1
	v163 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol386 = libc.Ptr(&libc.As[TSLexer](v163).F1)
	*libc.As[int16](result_symbol386) = 4
	v164 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end387 = libc.Ptr(&libc.As[TSLexer](v164).F3)
	v165 = *libc.As[unsafe.Pointer](mark_end387)
	v166 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v165)(v166)
	v167 = *libc.As[int32](lookahead)
	cmp388 = v167 == 61
	if cmp388 {
		goto if_then390
	} else {
		goto if_end391
	}

if_then390:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end391:
	v168 = *libc.As[int32](lookahead)
	cmp392 = v168 == 71
	if cmp392 {
		goto if_then394
	} else {
		goto if_end395
	}

if_then394:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end395:
	v169 = *libc.As[int32](lookahead)
	cmp396 = v169 == 43
	if cmp396 {
		goto if_then416
	} else {
		goto lor_lhs_false398
	}

lor_lhs_false398:
	v170 = *libc.As[int32](lookahead)
	cmp399 = 47 <= v170
	if cmp399 {
		goto land_lhs_true401
	} else {
		goto lor_lhs_false404
	}

land_lhs_true401:
	v171 = *libc.As[int32](lookahead)
	cmp402 = v171 <= 57
	if cmp402 {
		goto if_then416
	} else {
		goto lor_lhs_false404
	}

lor_lhs_false404:
	v172 = *libc.As[int32](lookahead)
	cmp405 = 65 <= v172
	if cmp405 {
		goto land_lhs_true407
	} else {
		goto lor_lhs_false410
	}

land_lhs_true407:
	v173 = *libc.As[int32](lookahead)
	cmp408 = v173 <= 90
	if cmp408 {
		goto if_then416
	} else {
		goto lor_lhs_false410
	}

lor_lhs_false410:
	v174 = *libc.As[int32](lookahead)
	cmp411 = 97 <= v174
	if cmp411 {
		goto land_lhs_true413
	} else {
		goto if_end417
	}

land_lhs_true413:
	v175 = *libc.As[int32](lookahead)
	cmp414 = v175 <= 122
	if cmp414 {
		goto if_then416
	} else {
		goto if_end417
	}

if_then416:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end417:
	v176 = *libc.As[byte](result)
	loadedv418 = (v176 & 1) != 0
	*libc.As[bool](retval) = loadedv418
	goto _return

sw_bb419:
	*libc.As[byte](result) = 1
	v177 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol420 = libc.Ptr(&libc.As[TSLexer](v177).F1)
	*libc.As[int16](result_symbol420) = 4
	v178 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end421 = libc.Ptr(&libc.As[TSLexer](v178).F3)
	v179 = *libc.As[unsafe.Pointer](mark_end421)
	v180 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v179)(v180)
	v181 = *libc.As[int32](lookahead)
	cmp422 = v181 == 61
	if cmp422 {
		goto if_then424
	} else {
		goto if_end425
	}

if_then424:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end425:
	v182 = *libc.As[int32](lookahead)
	cmp426 = v182 == 73
	if cmp426 {
		goto if_then428
	} else {
		goto if_end429
	}

if_then428:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end429:
	v183 = *libc.As[int32](lookahead)
	cmp430 = v183 == 43
	if cmp430 {
		goto if_then450
	} else {
		goto lor_lhs_false432
	}

lor_lhs_false432:
	v184 = *libc.As[int32](lookahead)
	cmp433 = 47 <= v184
	if cmp433 {
		goto land_lhs_true435
	} else {
		goto lor_lhs_false438
	}

land_lhs_true435:
	v185 = *libc.As[int32](lookahead)
	cmp436 = v185 <= 57
	if cmp436 {
		goto if_then450
	} else {
		goto lor_lhs_false438
	}

lor_lhs_false438:
	v186 = *libc.As[int32](lookahead)
	cmp439 = 65 <= v186
	if cmp439 {
		goto land_lhs_true441
	} else {
		goto lor_lhs_false444
	}

land_lhs_true441:
	v187 = *libc.As[int32](lookahead)
	cmp442 = v187 <= 90
	if cmp442 {
		goto if_then450
	} else {
		goto lor_lhs_false444
	}

lor_lhs_false444:
	v188 = *libc.As[int32](lookahead)
	cmp445 = 97 <= v188
	if cmp445 {
		goto land_lhs_true447
	} else {
		goto if_end451
	}

land_lhs_true447:
	v189 = *libc.As[int32](lookahead)
	cmp448 = v189 <= 122
	if cmp448 {
		goto if_then450
	} else {
		goto if_end451
	}

if_then450:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end451:
	v190 = *libc.As[byte](result)
	loadedv452 = (v190 & 1) != 0
	*libc.As[bool](retval) = loadedv452
	goto _return

sw_bb453:
	*libc.As[byte](result) = 1
	v191 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol454 = libc.Ptr(&libc.As[TSLexer](v191).F1)
	*libc.As[int16](result_symbol454) = 4
	v192 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end455 = libc.Ptr(&libc.As[TSLexer](v192).F3)
	v193 = *libc.As[unsafe.Pointer](mark_end455)
	v194 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v193)(v194)
	v195 = *libc.As[int32](lookahead)
	cmp456 = v195 == 61
	if cmp456 {
		goto if_then458
	} else {
		goto if_end459
	}

if_then458:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end459:
	v196 = *libc.As[int32](lookahead)
	cmp460 = v196 == 78
	if cmp460 {
		goto if_then462
	} else {
		goto if_end463
	}

if_then462:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end463:
	v197 = *libc.As[int32](lookahead)
	cmp464 = v197 == 43
	if cmp464 {
		goto if_then484
	} else {
		goto lor_lhs_false466
	}

lor_lhs_false466:
	v198 = *libc.As[int32](lookahead)
	cmp467 = 47 <= v198
	if cmp467 {
		goto land_lhs_true469
	} else {
		goto lor_lhs_false472
	}

land_lhs_true469:
	v199 = *libc.As[int32](lookahead)
	cmp470 = v199 <= 57
	if cmp470 {
		goto if_then484
	} else {
		goto lor_lhs_false472
	}

lor_lhs_false472:
	v200 = *libc.As[int32](lookahead)
	cmp473 = 65 <= v200
	if cmp473 {
		goto land_lhs_true475
	} else {
		goto lor_lhs_false478
	}

land_lhs_true475:
	v201 = *libc.As[int32](lookahead)
	cmp476 = v201 <= 90
	if cmp476 {
		goto if_then484
	} else {
		goto lor_lhs_false478
	}

lor_lhs_false478:
	v202 = *libc.As[int32](lookahead)
	cmp479 = 97 <= v202
	if cmp479 {
		goto land_lhs_true481
	} else {
		goto if_end485
	}

land_lhs_true481:
	v203 = *libc.As[int32](lookahead)
	cmp482 = v203 <= 122
	if cmp482 {
		goto if_then484
	} else {
		goto if_end485
	}

if_then484:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end485:
	v204 = *libc.As[byte](result)
	loadedv486 = (v204 & 1) != 0
	*libc.As[bool](retval) = loadedv486
	goto _return

sw_bb487:
	*libc.As[byte](result) = 1
	v205 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol488 = libc.Ptr(&libc.As[TSLexer](v205).F1)
	*libc.As[int16](result_symbol488) = 4
	v206 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end489 = libc.Ptr(&libc.As[TSLexer](v206).F3)
	v207 = *libc.As[unsafe.Pointer](mark_end489)
	v208 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v207)(v208)
	v209 = *libc.As[int32](lookahead)
	cmp490 = v209 == 61
	if cmp490 {
		goto if_then492
	} else {
		goto if_end493
	}

if_then492:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end493:
	v210 = *libc.As[int32](lookahead)
	cmp494 = v210 == 78
	if cmp494 {
		goto if_then496
	} else {
		goto if_end497
	}

if_then496:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end497:
	v211 = *libc.As[int32](lookahead)
	cmp498 = v211 == 43
	if cmp498 {
		goto if_then518
	} else {
		goto lor_lhs_false500
	}

lor_lhs_false500:
	v212 = *libc.As[int32](lookahead)
	cmp501 = 47 <= v212
	if cmp501 {
		goto land_lhs_true503
	} else {
		goto lor_lhs_false506
	}

land_lhs_true503:
	v213 = *libc.As[int32](lookahead)
	cmp504 = v213 <= 57
	if cmp504 {
		goto if_then518
	} else {
		goto lor_lhs_false506
	}

lor_lhs_false506:
	v214 = *libc.As[int32](lookahead)
	cmp507 = 65 <= v214
	if cmp507 {
		goto land_lhs_true509
	} else {
		goto lor_lhs_false512
	}

land_lhs_true509:
	v215 = *libc.As[int32](lookahead)
	cmp510 = v215 <= 90
	if cmp510 {
		goto if_then518
	} else {
		goto lor_lhs_false512
	}

lor_lhs_false512:
	v216 = *libc.As[int32](lookahead)
	cmp513 = 97 <= v216
	if cmp513 {
		goto land_lhs_true515
	} else {
		goto if_end519
	}

land_lhs_true515:
	v217 = *libc.As[int32](lookahead)
	cmp516 = v217 <= 122
	if cmp516 {
		goto if_then518
	} else {
		goto if_end519
	}

if_then518:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end519:
	v218 = *libc.As[byte](result)
	loadedv520 = (v218 & 1) != 0
	*libc.As[bool](retval) = loadedv520
	goto _return

sw_bb521:
	*libc.As[byte](result) = 1
	v219 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol522 = libc.Ptr(&libc.As[TSLexer](v219).F1)
	*libc.As[int16](result_symbol522) = 4
	v220 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end523 = libc.Ptr(&libc.As[TSLexer](v220).F3)
	v221 = *libc.As[unsafe.Pointer](mark_end523)
	v222 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v221)(v222)
	v223 = *libc.As[int32](lookahead)
	cmp524 = v223 == 61
	if cmp524 {
		goto if_then526
	} else {
		goto if_end527
	}

if_then526:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end527:
	v224 = *libc.As[int32](lookahead)
	cmp528 = v224 == 43
	if cmp528 {
		goto if_then548
	} else {
		goto lor_lhs_false530
	}

lor_lhs_false530:
	v225 = *libc.As[int32](lookahead)
	cmp531 = 47 <= v225
	if cmp531 {
		goto land_lhs_true533
	} else {
		goto lor_lhs_false536
	}

land_lhs_true533:
	v226 = *libc.As[int32](lookahead)
	cmp534 = v226 <= 57
	if cmp534 {
		goto if_then548
	} else {
		goto lor_lhs_false536
	}

lor_lhs_false536:
	v227 = *libc.As[int32](lookahead)
	cmp537 = 65 <= v227
	if cmp537 {
		goto land_lhs_true539
	} else {
		goto lor_lhs_false542
	}

land_lhs_true539:
	v228 = *libc.As[int32](lookahead)
	cmp540 = v228 <= 90
	if cmp540 {
		goto if_then548
	} else {
		goto lor_lhs_false542
	}

lor_lhs_false542:
	v229 = *libc.As[int32](lookahead)
	cmp543 = 97 <= v229
	if cmp543 {
		goto land_lhs_true545
	} else {
		goto if_end549
	}

land_lhs_true545:
	v230 = *libc.As[int32](lookahead)
	cmp546 = v230 <= 122
	if cmp546 {
		goto if_then548
	} else {
		goto if_end549
	}

if_then548:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end549:
	v231 = *libc.As[byte](result)
	loadedv550 = (v231 & 1) != 0
	*libc.As[bool](retval) = loadedv550
	goto _return

sw_bb551:
	*libc.As[byte](result) = 1
	v232 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol552 = libc.Ptr(&libc.As[TSLexer](v232).F1)
	*libc.As[int16](result_symbol552) = 4
	v233 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end553 = libc.Ptr(&libc.As[TSLexer](v233).F3)
	v234 = *libc.As[unsafe.Pointer](mark_end553)
	v235 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v234)(v235)
	v236 = *libc.As[int32](lookahead)
	cmp554 = v236 == 61
	if cmp554 {
		goto if_then556
	} else {
		goto if_end557
	}

if_then556:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end557:
	v237 = *libc.As[byte](result)
	loadedv558 = (v237 & 1) != 0
	*libc.As[bool](retval) = loadedv558
	goto _return

sw_bb559:
	*libc.As[byte](result) = 1
	v238 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol560 = libc.Ptr(&libc.As[TSLexer](v238).F1)
	*libc.As[int16](result_symbol560) = 5
	v239 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end561 = libc.Ptr(&libc.As[TSLexer](v239).F3)
	v240 = *libc.As[unsafe.Pointer](mark_end561)
	v241 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v240)(v241)
	v242 = *libc.As[int32](lookahead)
	cmp562 = v242 == 10
	if cmp562 {
		goto if_then564
	} else {
		goto if_end565
	}

if_then564:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end565:
	v243 = *libc.As[int32](lookahead)
	cmp566 = v243 == 13
	if cmp566 {
		goto if_then568
	} else {
		goto if_end569
	}

if_then568:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end569:
	v244 = *libc.As[int32](lookahead)
	cmp570 = v244 != 0
	if cmp570 {
		goto land_lhs_true572
	} else {
		goto if_end576
	}

land_lhs_true572:
	v245 = *libc.As[int32](lookahead)
	cmp573 = v245 != 45
	if cmp573 {
		goto if_then575
	} else {
		goto if_end576
	}

if_then575:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end576:
	v246 = *libc.As[byte](result)
	loadedv577 = (v246 & 1) != 0
	*libc.As[bool](retval) = loadedv577
	goto _return

sw_bb578:
	*libc.As[byte](result) = 1
	v247 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol579 = libc.Ptr(&libc.As[TSLexer](v247).F1)
	*libc.As[int16](result_symbol579) = 5
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end580 = libc.Ptr(&libc.As[TSLexer](v248).F3)
	v249 = *libc.As[unsafe.Pointer](mark_end580)
	v250 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v249)(v250)
	v251 = *libc.As[int32](lookahead)
	cmp581 = v251 == 10
	if cmp581 {
		goto if_then583
	} else {
		goto if_end584
	}

if_then583:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end584:
	v252 = *libc.As[int32](lookahead)
	cmp585 = v252 != 0
	if cmp585 {
		goto land_lhs_true587
	} else {
		goto if_end591
	}

land_lhs_true587:
	v253 = *libc.As[int32](lookahead)
	cmp588 = v253 != 45
	if cmp588 {
		goto if_then590
	} else {
		goto if_end591
	}

if_then590:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end591:
	v254 = *libc.As[byte](result)
	loadedv592 = (v254 & 1) != 0
	*libc.As[bool](retval) = loadedv592
	goto _return

sw_bb593:
	*libc.As[byte](result) = 1
	v255 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol594 = libc.Ptr(&libc.As[TSLexer](v255).F1)
	*libc.As[int16](result_symbol594) = 5
	v256 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end595 = libc.Ptr(&libc.As[TSLexer](v256).F3)
	v257 = *libc.As[unsafe.Pointer](mark_end595)
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v257)(v258)
	v259 = *libc.As[int32](lookahead)
	cmp596 = v259 != 0
	if cmp596 {
		goto land_lhs_true598
	} else {
		goto if_end602
	}

land_lhs_true598:
	v260 = *libc.As[int32](lookahead)
	cmp599 = v260 != 45
	if cmp599 {
		goto if_then601
	} else {
		goto if_end602
	}

if_then601:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end602:
	v261 = *libc.As[byte](result)
	loadedv603 = (v261 & 1) != 0
	*libc.As[bool](retval) = loadedv603
	goto _return

sw_bb604:
	*libc.As[byte](result) = 1
	v262 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol605 = libc.Ptr(&libc.As[TSLexer](v262).F1)
	*libc.As[int16](result_symbol605) = 6
	v263 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end606 = libc.Ptr(&libc.As[TSLexer](v263).F3)
	v264 = *libc.As[unsafe.Pointer](mark_end606)
	v265 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v264)(v265)
	v266 = *libc.As[byte](result)
	loadedv607 = (v266 & 1) != 0
	*libc.As[bool](retval) = loadedv607
	goto _return

sw_bb608:
	*libc.As[byte](result) = 1
	v267 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol609 = libc.Ptr(&libc.As[TSLexer](v267).F1)
	*libc.As[int16](result_symbol609) = 7
	v268 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end610 = libc.Ptr(&libc.As[TSLexer](v268).F3)
	v269 = *libc.As[unsafe.Pointer](mark_end610)
	v270 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v269)(v270)
	v271 = *libc.As[int32](lookahead)
	cmp611 = v271 == 10
	if cmp611 {
		goto if_then613
	} else {
		goto if_end614
	}

if_then613:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end614:
	v272 = *libc.As[int32](lookahead)
	cmp615 = v272 != 0
	if cmp615 {
		goto if_then617
	} else {
		goto if_end618
	}

if_then617:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end618:
	v273 = *libc.As[byte](result)
	loadedv619 = (v273 & 1) != 0
	*libc.As[bool](retval) = loadedv619
	goto _return

sw_bb620:
	*libc.As[byte](result) = 1
	v274 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol621 = libc.Ptr(&libc.As[TSLexer](v274).F1)
	*libc.As[int16](result_symbol621) = 7
	v275 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end622 = libc.Ptr(&libc.As[TSLexer](v275).F3)
	v276 = *libc.As[unsafe.Pointer](mark_end622)
	v277 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v276)(v277)
	v278 = *libc.As[int32](lookahead)
	cmp623 = v278 == 45
	if cmp623 {
		goto if_then625
	} else {
		goto if_end626
	}

if_then625:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end626:
	v279 = *libc.As[int32](lookahead)
	cmp627 = v279 != 0
	if cmp627 {
		goto land_lhs_true629
	} else {
		goto if_end633
	}

land_lhs_true629:
	v280 = *libc.As[int32](lookahead)
	cmp630 = v280 != 10
	if cmp630 {
		goto if_then632
	} else {
		goto if_end633
	}

if_then632:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end633:
	v281 = *libc.As[byte](result)
	loadedv634 = (v281 & 1) != 0
	*libc.As[bool](retval) = loadedv634
	goto _return

sw_bb635:
	*libc.As[byte](result) = 1
	v282 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol636 = libc.Ptr(&libc.As[TSLexer](v282).F1)
	*libc.As[int16](result_symbol636) = 7
	v283 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end637 = libc.Ptr(&libc.As[TSLexer](v283).F3)
	v284 = *libc.As[unsafe.Pointer](mark_end637)
	v285 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v284)(v285)
	v286 = *libc.As[int32](lookahead)
	cmp638 = v286 == 45
	if cmp638 {
		goto if_then640
	} else {
		goto if_end641
	}

if_then640:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end641:
	v287 = *libc.As[int32](lookahead)
	cmp642 = v287 != 0
	if cmp642 {
		goto land_lhs_true644
	} else {
		goto if_end648
	}

land_lhs_true644:
	v288 = *libc.As[int32](lookahead)
	cmp645 = v288 != 10
	if cmp645 {
		goto if_then647
	} else {
		goto if_end648
	}

if_then647:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end648:
	v289 = *libc.As[byte](result)
	loadedv649 = (v289 & 1) != 0
	*libc.As[bool](retval) = loadedv649
	goto _return

sw_bb650:
	*libc.As[byte](result) = 1
	v290 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol651 = libc.Ptr(&libc.As[TSLexer](v290).F1)
	*libc.As[int16](result_symbol651) = 7
	v291 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end652 = libc.Ptr(&libc.As[TSLexer](v291).F3)
	v292 = *libc.As[unsafe.Pointer](mark_end652)
	v293 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v292)(v293)
	v294 = *libc.As[int32](lookahead)
	cmp653 = v294 == 45
	if cmp653 {
		goto if_then655
	} else {
		goto if_end656
	}

if_then655:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end656:
	v295 = *libc.As[int32](lookahead)
	cmp657 = v295 != 0
	if cmp657 {
		goto land_lhs_true659
	} else {
		goto if_end663
	}

land_lhs_true659:
	v296 = *libc.As[int32](lookahead)
	cmp660 = v296 != 10
	if cmp660 {
		goto if_then662
	} else {
		goto if_end663
	}

if_then662:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end663:
	v297 = *libc.As[byte](result)
	loadedv664 = (v297 & 1) != 0
	*libc.As[bool](retval) = loadedv664
	goto _return

sw_bb665:
	*libc.As[byte](result) = 1
	v298 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol666 = libc.Ptr(&libc.As[TSLexer](v298).F1)
	*libc.As[int16](result_symbol666) = 7
	v299 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end667 = libc.Ptr(&libc.As[TSLexer](v299).F3)
	v300 = *libc.As[unsafe.Pointer](mark_end667)
	v301 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v300)(v301)
	v302 = *libc.As[int32](lookahead)
	cmp668 = v302 == 45
	if cmp668 {
		goto if_then670
	} else {
		goto if_end671
	}

if_then670:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end671:
	v303 = *libc.As[int32](lookahead)
	cmp672 = v303 != 0
	if cmp672 {
		goto land_lhs_true674
	} else {
		goto if_end678
	}

land_lhs_true674:
	v304 = *libc.As[int32](lookahead)
	cmp675 = v304 != 10
	if cmp675 {
		goto if_then677
	} else {
		goto if_end678
	}

if_then677:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end678:
	v305 = *libc.As[byte](result)
	loadedv679 = (v305 & 1) != 0
	*libc.As[bool](retval) = loadedv679
	goto _return

sw_bb680:
	*libc.As[byte](result) = 1
	v306 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol681 = libc.Ptr(&libc.As[TSLexer](v306).F1)
	*libc.As[int16](result_symbol681) = 7
	v307 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end682 = libc.Ptr(&libc.As[TSLexer](v307).F3)
	v308 = *libc.As[unsafe.Pointer](mark_end682)
	v309 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v308)(v309)
	v310 = *libc.As[int32](lookahead)
	cmp683 = v310 != 0
	if cmp683 {
		goto land_lhs_true685
	} else {
		goto if_end689
	}

land_lhs_true685:
	v311 = *libc.As[int32](lookahead)
	cmp686 = v311 != 10
	if cmp686 {
		goto if_then688
	} else {
		goto if_end689
	}

if_then688:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end689:
	v312 = *libc.As[byte](result)
	loadedv690 = (v312 & 1) != 0
	*libc.As[bool](retval) = loadedv690
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v313 = *libc.As[bool](retval)
	return v313
}
