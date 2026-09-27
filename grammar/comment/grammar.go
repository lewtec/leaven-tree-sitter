package grammar_comment

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
type TSCharacterRange struct {
	F0 int32
	F1 int32
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

var tree_sitter_comment_language struct {
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
var ts_small_parse_table [71]int16 = [71]int16{3, 53, 1, 5, 7, 1, 33, 51, 14, 1, 4, 9, 10, 11, 15, 16, 17, 18, 19, 20, 21, 22, 23, 4, 3, 1, 5, 55, 1, 1, 57, 1, 2, 13, 1, 29, 2, 3, 1, 5, 59, 1, 0, 2, 61, 1, 3, 63, 1, 5, 2, 3, 1, 5, 65, 1, 1, 2, 3, 1, 5, 67, 1, 4, 2, 3, 1, 5, 69, 1, 1}
var ts_small_parse_table_map [7]int32 = [7]int32{0, 23, 36, 43, 50, 57, 64}
var ts_symbol_names [35]unsafe.Pointer = [35]unsafe.Pointer{libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_30), libc.Ptr(&_str_31), libc.Ptr(&_str_32), libc.Ptr(&_str_33), libc.Ptr(&_str_34), libc.Ptr(&_str_35), libc.Ptr(&_str_34), libc.Ptr(&_str_36)}
var ts_symbol_metadata [35]TSSymbolMetadata = [35]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}}
var ts_symbol_map [35]int16 = [35]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 31, 34}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][3]int16 = [1][3]int16{}
var ts_lex_modes [16]TSLexerMode = [16]TSLexerMode{TSLexerMode{0, 1, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}}
var ts_external_scanner_states [3][2]byte = [3][2]byte{[2]byte{}, [2]byte{1, 1}, [2]byte{1, 0}}
var ts_external_scanner_symbol_map [2]int16 = [2]int16{25, 26}
var ts_primary_state_ids [16]int16 = [16]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
var _str [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var __const_is_internal_char_valid_chars [2]int32 = [2]int32{45, 95}
var __const_is_space_space_chars [4]int32 = [4]int32{32, 12, 9, 11}
var __const_is_newline_newline_chars [3]int32 = [3]int32{0, 10, 13}
var ts_parse_table struct {
	F0 struct {
		F0 [27]int16
		F1 [8]int16
	}
	F1 [35]int16
	F2 [35]int16
	F3 [35]int16
	F4 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F5 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F6 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F7 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F8 struct {
		F0 [26]int16
		F1 [9]int16
	}
} = struct {
	F0 struct {
		F0 [27]int16
		F1 [8]int16
	}
	F1 [35]int16
	F2 [35]int16
	F3 [35]int16
	F4 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F5 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F6 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F7 struct {
		F0 [26]int16
		F1 [9]int16
	}
	F8 struct {
		F0 [26]int16
		F1 [9]int16
	}
}{struct {
	F0 [27]int16
	F1 [8]int16
}{[27]int16{1, 1, 1, 0, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, [8]int16{}}, [35]int16{5, 7, 7, 0, 7, 3, 9, 11, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 13, 0, 11, 2, 0, 2, 2, 4, 0, 2}, [35]int16{15, 7, 7, 0, 7, 3, 9, 11, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 13, 0, 0, 3, 0, 3, 3, 4, 0, 3}, [35]int16{17, 19, 19, 0, 19, 3, 22, 25, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 28, 0, 0, 3, 0, 3, 3, 4, 0, 3}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{31, 31, 31, 0, 31, 3, 31, 33, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{35, 35, 35, 0, 35, 3, 35, 37, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{39, 39, 39, 0, 39, 3, 39, 41, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{43, 43, 43, 0, 43, 3, 43, 45, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{47, 47, 47, 0, 47, 3, 47, 49, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47}, [9]int16{}}}
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
		F0 struct {
			F0 struct {
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
	F32 TSParseActionEntry
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
			F0 byte
			F1 [7]byte
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
	F70 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
	F32 TSParseActionEntry
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
			F0 byte
			F1 [7]byte
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
	F70 TSParseActionEntry
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 27, 0, 0}}}, struct {
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
}{0, 0, 10, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 27, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 31, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 31, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 28, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 28, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 33, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 33, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 30, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 30, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 28, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 28, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 29, 0, 0}}}}
var _str_3 [4]byte = [4]byte{101, 110, 100, 0}
var _str_4 [2]byte = [2]byte{58, 0}
var _str_5 [2]byte = [2]byte{40, 0}
var _str_6 [5]byte = [5]byte{117, 115, 101, 114, 0}
var _str_7 [2]byte = [2]byte{41, 0}
var _str_8 [17]byte = [17]byte{95, 102, 117, 108, 108, 95, 117, 114, 105, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_9 [4]byte = [4]byte{117, 114, 105, 0}
var _str_10 [13]byte = [13]byte{95, 116, 101, 120, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_11 [2]byte = [2]byte{47, 0}
var _str_12 [2]byte = [2]byte{39, 0}
var _str_13 [2]byte = [2]byte{34, 0}
var _str_14 [2]byte = [2]byte{96, 0}
var _str_15 [2]byte = [2]byte{60, 0}
var _str_16 [2]byte = [2]byte{91, 0}
var _str_17 [2]byte = [2]byte{123, 0}
var _str_18 [2]byte = [2]byte{46, 0}
var _str_19 [2]byte = [2]byte{44, 0}
var _str_20 [2]byte = [2]byte{59, 0}
var _str_21 [2]byte = [2]byte{33, 0}
var _str_22 [2]byte = [2]byte{63, 0}
var _str_23 [2]byte = [2]byte{92, 0}
var _str_24 [2]byte = [2]byte{125, 0}
var _str_25 [2]byte = [2]byte{93, 0}
var _str_26 [2]byte = [2]byte{62, 0}
var _str_27 [2]byte = [2]byte{45, 0}
var _str_28 [5]byte = [5]byte{110, 97, 109, 101, 0}
var _str_29 [14]byte = [14]byte{105, 110, 118, 97, 108, 105, 100, 95, 116, 111, 107, 101, 110, 0}
var _str_30 [7]byte = [7]byte{115, 111, 117, 114, 99, 101, 0}
var _str_31 [4]byte = [4]byte{116, 97, 103, 0}
var _str_32 [6]byte = [6]byte{95, 117, 115, 101, 114, 0}
var _str_33 [10]byte = [10]byte{95, 102, 117, 108, 108, 95, 117, 114, 105, 0}
var _str_34 [5]byte = [5]byte{116, 101, 120, 116, 0}
var _str_35 [11]byte = [11]byte{95, 115, 116, 111, 112, 95, 99, 104, 97, 114, 0}
var _str_36 [15]byte = [15]byte{115, 111, 117, 114, 99, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var ts_lex_map [42]int16 = [42]int16{33, 29, 34, 21, 39, 20, 40, 8, 41, 10, 44, 27, 45, 35, 46, 26, 47, 19, 58, 7, 59, 28, 60, 23, 62, 34, 63, 30, 91, 24, 92, 31, 93, 33, 96, 22, 104, 17, 123, 25, 125, 32}
var ts_lex_map_37 [28]int16 = [28]int16{33, 5, 34, 5, 39, 5, 41, 5, 44, 5, 46, 5, 58, 5, 59, 5, 62, 5, 63, 5, 92, 5, 93, 5, 96, 5, 125, 5}
var sym_uri_character_set_1 [12]TSCharacterRange = [12]TSCharacterRange{TSCharacterRange{0, 8}, TSCharacterRange{14, 31}, TSCharacterRange{35, 38}, TSCharacterRange{40, 40}, TSCharacterRange{42, 43}, TSCharacterRange{45, 45}, TSCharacterRange{47, 57}, TSCharacterRange{60, 61}, TSCharacterRange{64, 91}, TSCharacterRange{94, 95}, TSCharacterRange{97, 124}, TSCharacterRange{126, 1114111}}
var ts_lex_map_38 [28]int16 = [28]int16{33, 5, 34, 5, 39, 5, 41, 5, 44, 5, 46, 5, 58, 5, 59, 5, 62, 5, 63, 5, 92, 5, 93, 5, 96, 5, 125, 5}
var aux_sym__text_token1_character_set_1 [11]TSCharacterRange = [11]TSCharacterRange{TSCharacterRange{0, 8}, TSCharacterRange{14, 31}, TSCharacterRange{35, 38}, TSCharacterRange{42, 43}, TSCharacterRange{48, 57}, TSCharacterRange{61, 61}, TSCharacterRange{64, 90}, TSCharacterRange{94, 95}, TSCharacterRange{97, 122}, TSCharacterRange{124, 124}, TSCharacterRange{126, 1114111}}

func init() {
	tree_sitter_comment_language = struct {
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
	}{15, 35, 0, 27, 2, 16, 9, 1, 0, 3, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{libc.Ptr(&ts_external_scanner_states), libc.Ptr(&ts_external_scanner_symbol_map), libc.FuncCode(tree_sitter_comment_external_scanner_create), libc.FuncCode(tree_sitter_comment_external_scanner_destroy), libc.FuncCode(tree_sitter_comment_external_scanner_scan), libc.FuncCode(tree_sitter_comment_external_scanner_serialize), libc.FuncCode(tree_sitter_comment_external_scanner_deserialize)}, libc.Ptr(&ts_primary_state_ids), libc.Ptr(&_str), nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{0, 2, 0}, [5]byte{}}
}
func tree_sitter_comment_external_scanner_create() unsafe.Pointer {
	return nil
}
func tree_sitter_comment_external_scanner_destroy(payload unsafe.Pointer) {
	var payload_addr unsafe.Pointer
	_ = payload_addr

	payload_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
}
func tree_sitter_comment_external_scanner_serialize(payload unsafe.Pointer, buffer unsafe.Pointer) int32 {
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
func tree_sitter_comment_external_scanner_deserialize(payload unsafe.Pointer, buffer unsafe.Pointer, length int32) {
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
func tree_sitter_comment_external_scanner_scan(payload unsafe.Pointer, lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var call bool
	var v0, v1 unsafe.Pointer
	var payload_addr, lexer_addr, valid_symbols_addr unsafe.Pointer
	_, _, _, _, _, _ = payload_addr, lexer_addr, valid_symbols_addr, v0, v1, call

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
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	v1 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	call = parse(v0, v1)
	return call
}
func parse(lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var loadedv, call, loadedv2, call4, v8 bool
	var retval unsafe.Pointer
	var v3 int32
	var lookahead unsafe.Pointer
	var v1, v5 byte
	var arrayidx, arrayidx1 unsafe.Pointer
	var v0, v2, v4, v6, v7 unsafe.Pointer
	var lexer_addr, valid_symbols_addr unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, valid_symbols_addr, v0, arrayidx, v1, loadedv, v2, lookahead, v3, call, v4, arrayidx1, v5, loadedv2, v6, v7, call4, v8

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
	valid_symbols_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v0), int(int64(1))*1))
	v1 = *libc.As[byte](arrayidx)
	loadedv = (v1 & 1) != 0
	if loadedv {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v2).F0)
	v3 = *libc.As[int32](lookahead)
	call = is_upper(v3)
	if call {
		goto land_lhs_true
	} else {
		goto if_end5
	}

land_lhs_true:
	v4 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx1 = v4
	v5 = *libc.As[byte](arrayidx1)
	loadedv2 = (v5 & 1) != 0
	if loadedv2 {
		goto if_then3
	} else {
		goto if_end5
	}

if_then3:
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	v7 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	call4 = parse_tagname(v6, v7)
	*libc.As[bool](retval) = call4
	goto _return

if_end5:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v8 = *libc.As[bool](retval)
	return v8
}
func tree_sitter_comment() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_comment_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, loadedv26, cmp28, loadedv32, cmp34, loadedv38, cmp40, cmp43, cmp46, cmp50, cmp53, cmp56, loadedv60, cmp65, cmp71, cmp81, cmp84, cmp87, cmp90, cmp93, loadedv97, loadedv99, call101, loadedv105, loadedv107, loadedv111, loadedv115, cmp119, cmp122, cmp125, loadedv129, loadedv133, loadedv137, cmp144, cmp150, cmp160, cmp163, cmp166, cmp169, cmp172, loadedv176, cmp180, cmp184, loadedv188, call190, loadedv194, cmp198, loadedv202, call204, loadedv208, cmp212, loadedv216, call218, loadedv222, cmp226, loadedv230, call232, loadedv236, cmp240, loadedv244, call246, loadedv250, loadedv254, call256, loadedv260, loadedv264, loadedv268, loadedv272, loadedv276, loadedv280, loadedv284, loadedv288, loadedv292, loadedv296, loadedv300, loadedv304, loadedv308, loadedv312, loadedv316, loadedv320, loadedv324, loadedv328, v233 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v36, v39, v89, v92 int16
	var state_addr, arrayidx, arrayidx11, arrayidx69, arrayidx76, result_symbol, result_symbol109, result_symbol113, result_symbol117, result_symbol131, result_symbol135, result_symbol139, arrayidx148, arrayidx155, result_symbol178, result_symbol196, result_symbol210, result_symbol224, result_symbol238, result_symbol252, result_symbol262, result_symbol266, result_symbol270, result_symbol274, result_symbol278, result_symbol282, result_symbol286, result_symbol290, result_symbol294, result_symbol298, result_symbol302, result_symbol306, result_symbol310, result_symbol314, result_symbol318, result_symbol322, result_symbol326 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v23, v25, v27, v28, v29, v30, v31, v32, v34, v35, conv70, v37, v38, add74, v40, add79, v41, v42, v43, v44, v45, v48, v69, v70, v71, v87, v88, conv149, v90, v91, add153, v93, add158, v94, v95, v96, v97, v98, v104, v105, v107, v113, v115, v121, v123, v129, v131, v137, v139, v146 int32
	var lookahead, i, i62, i141, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv64, idxprom68, idxprom75, conv143, idxprom147, idxprom154 int64
	var v3, storedv, v10, v22, v24, v26, v33, v46, v47, v49, v54, v59, v64, v72, v77, v82, v99, v106, v108, v114, v116, v122, v124, v130, v132, v138, v140, v145, v147, v152, v157, v162, v167, v172, v177, v182, v187, v192, v197, v202, v207, v212, v217, v222, v227, v232 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v50, v51, v52, v53, v55, v56, v57, v58, v60, v61, v62, v63, v65, v66, v67, v68, v73, v74, v75, v76, v78, v79, v80, v81, v83, v84, v85, v86, v100, v101, v102, v103, v109, v110, v111, v112, v117, v118, v119, v120, v125, v126, v127, v128, v133, v134, v135, v136, v141, v142, v143, v144, v148, v149, v150, v151, v153, v154, v155, v156, v158, v159, v160, v161, v163, v164, v165, v166, v168, v169, v170, v171, v173, v174, v175, v176, v178, v179, v180, v181, v183, v184, v185, v186, v188, v189, v190, v191, v193, v194, v195, v196, v198, v199, v200, v201, v203, v204, v205, v206, v208, v209, v210, v211, v213, v214, v215, v216, v218, v219, v220, v221, v223, v224, v225, v226, v228, v229, v230, v231 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end110, mark_end114, mark_end118, mark_end132, mark_end136, mark_end140, mark_end179, mark_end197, mark_end211, mark_end225, mark_end239, mark_end253, mark_end263, mark_end267, mark_end271, mark_end275, mark_end279, mark_end283, mark_end287, mark_end291, mark_end295, mark_end299, mark_end303, mark_end307, mark_end311, mark_end315, mark_end319, mark_end323, mark_end327 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i62, i141, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, loadedv26, v23, cmp28, v24, loadedv32, v25, cmp34, v26, loadedv38, v27, cmp40, v28, cmp43, v29, cmp46, v30, cmp50, v31, cmp53, v32, cmp56, v33, loadedv60, v34, conv64, cmp65, v35, idxprom68, arrayidx69, v36, conv70, v37, cmp71, v38, add74, idxprom75, arrayidx76, v39, v40, add79, v41, cmp81, v42, cmp84, v43, cmp87, v44, cmp90, v45, cmp93, v46, loadedv97, v47, loadedv99, v48, call101, v49, loadedv105, v50, result_symbol, v51, mark_end, v52, v53, v54, loadedv107, v55, result_symbol109, v56, mark_end110, v57, v58, v59, loadedv111, v60, result_symbol113, v61, mark_end114, v62, v63, v64, loadedv115, v65, result_symbol117, v66, mark_end118, v67, v68, v69, cmp119, v70, cmp122, v71, cmp125, v72, loadedv129, v73, result_symbol131, v74, mark_end132, v75, v76, v77, loadedv133, v78, result_symbol135, v79, mark_end136, v80, v81, v82, loadedv137, v83, result_symbol139, v84, mark_end140, v85, v86, v87, conv143, cmp144, v88, idxprom147, arrayidx148, v89, conv149, v90, cmp150, v91, add153, idxprom154, arrayidx155, v92, v93, add158, v94, cmp160, v95, cmp163, v96, cmp166, v97, cmp169, v98, cmp172, v99, loadedv176, v100, result_symbol178, v101, mark_end179, v102, v103, v104, cmp180, v105, cmp184, v106, loadedv188, v107, call190, v108, loadedv194, v109, result_symbol196, v110, mark_end197, v111, v112, v113, cmp198, v114, loadedv202, v115, call204, v116, loadedv208, v117, result_symbol210, v118, mark_end211, v119, v120, v121, cmp212, v122, loadedv216, v123, call218, v124, loadedv222, v125, result_symbol224, v126, mark_end225, v127, v128, v129, cmp226, v130, loadedv230, v131, call232, v132, loadedv236, v133, result_symbol238, v134, mark_end239, v135, v136, v137, cmp240, v138, loadedv244, v139, call246, v140, loadedv250, v141, result_symbol252, v142, mark_end253, v143, v144, v145, loadedv254, v146, call256, v147, loadedv260, v148, result_symbol262, v149, mark_end263, v150, v151, v152, loadedv264, v153, result_symbol266, v154, mark_end267, v155, v156, v157, loadedv268, v158, result_symbol270, v159, mark_end271, v160, v161, v162, loadedv272, v163, result_symbol274, v164, mark_end275, v165, v166, v167, loadedv276, v168, result_symbol278, v169, mark_end279, v170, v171, v172, loadedv280, v173, result_symbol282, v174, mark_end283, v175, v176, v177, loadedv284, v178, result_symbol286, v179, mark_end287, v180, v181, v182, loadedv288, v183, result_symbol290, v184, mark_end291, v185, v186, v187, loadedv292, v188, result_symbol294, v189, mark_end295, v190, v191, v192, loadedv296, v193, result_symbol298, v194, mark_end299, v195, v196, v197, loadedv300, v198, result_symbol302, v199, mark_end303, v200, v201, v202, loadedv304, v203, result_symbol306, v204, mark_end307, v205, v206, v207, loadedv308, v208, result_symbol310, v209, mark_end311, v210, v211, v212, loadedv312, v213, result_symbol314, v214, mark_end315, v215, v216, v217, loadedv316, v218, result_symbol318, v219, mark_end319, v220, v221, v222, loadedv320, v223, result_symbol322, v224, mark_end323, v225, v226, v227, loadedv324, v228, result_symbol326, v229, mark_end327, v230, v231, v232, loadedv328, v233

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
	i62 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i141 = libc.Ptr(&new(struct {
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
		goto sw_bb27
	case 2:
		goto sw_bb33
	case 3:
		goto sw_bb39
	case 4:
		goto sw_bb61
	case 5:
		goto sw_bb98
	case 6:
		goto sw_bb106
	case 7:
		goto sw_bb108
	case 8:
		goto sw_bb112
	case 9:
		goto sw_bb116
	case 10:
		goto sw_bb130
	case 11:
		goto sw_bb134
	case 12:
		goto sw_bb138
	case 13:
		goto sw_bb177
	case 14:
		goto sw_bb195
	case 15:
		goto sw_bb209
	case 16:
		goto sw_bb223
	case 17:
		goto sw_bb237
	case 18:
		goto sw_bb251
	case 19:
		goto sw_bb261
	case 20:
		goto sw_bb265
	case 21:
		goto sw_bb269
	case 22:
		goto sw_bb273
	case 23:
		goto sw_bb277
	case 24:
		goto sw_bb281
	case 25:
		goto sw_bb285
	case 26:
		goto sw_bb289
	case 27:
		goto sw_bb293
	case 28:
		goto sw_bb297
	case 29:
		goto sw_bb301
	case 30:
		goto sw_bb305
	case 31:
		goto sw_bb309
	case 32:
		goto sw_bb313
	case 33:
		goto sw_bb317
	case 34:
		goto sw_bb321
	case 35:
		goto sw_bb325
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
	cmp = uint64(conv4) < uint64(42)
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
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end21:
	v21 = *libc.As[int32](lookahead)
	cmp22 = v21 != 0
	if cmp22 {
		goto if_then24
	} else {
		goto if_end25
	}

if_then24:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end25:
	v22 = *libc.As[byte](result)
	loadedv26 = (v22 & 1) != 0
	*libc.As[bool](retval) = loadedv26
	goto _return

sw_bb27:
	v23 = *libc.As[int32](lookahead)
	cmp28 = v23 == 47
	if cmp28 {
		goto if_then30
	} else {
		goto if_end31
	}

if_then30:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end31:
	v24 = *libc.As[byte](result)
	loadedv32 = (v24 & 1) != 0
	*libc.As[bool](retval) = loadedv32
	goto _return

sw_bb33:
	v25 = *libc.As[int32](lookahead)
	cmp34 = v25 == 47
	if cmp34 {
		goto if_then36
	} else {
		goto if_end37
	}

if_then36:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end37:
	v26 = *libc.As[byte](result)
	loadedv38 = (v26 & 1) != 0
	*libc.As[bool](retval) = loadedv38
	goto _return

sw_bb39:
	v27 = *libc.As[int32](lookahead)
	cmp40 = 9 <= v27
	if cmp40 {
		goto land_lhs_true42
	} else {
		goto lor_lhs_false45
	}

land_lhs_true42:
	v28 = *libc.As[int32](lookahead)
	cmp43 = v28 <= 13
	if cmp43 {
		goto if_then48
	} else {
		goto lor_lhs_false45
	}

lor_lhs_false45:
	v29 = *libc.As[int32](lookahead)
	cmp46 = v29 == 32
	if cmp46 {
		goto if_then48
	} else {
		goto if_end49
	}

if_then48:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end49:
	v30 = *libc.As[int32](lookahead)
	cmp50 = v30 != 0
	if cmp50 {
		goto land_lhs_true52
	} else {
		goto if_end59
	}

land_lhs_true52:
	v31 = *libc.As[int32](lookahead)
	cmp53 = v31 != 40
	if cmp53 {
		goto land_lhs_true55
	} else {
		goto if_end59
	}

land_lhs_true55:
	v32 = *libc.As[int32](lookahead)
	cmp56 = v32 != 41
	if cmp56 {
		goto if_then58
	} else {
		goto if_end59
	}

if_then58:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end59:
	v33 = *libc.As[byte](result)
	loadedv60 = (v33 & 1) != 0
	*libc.As[bool](retval) = loadedv60
	goto _return

sw_bb61:
	*libc.As[int32](i62) = 0
	goto for_cond63

for_cond63:
	v34 = *libc.As[int32](i62)
	conv64 = int64(uint64(uint32(v34)))
	cmp65 = uint64(conv64) < uint64(28)
	if cmp65 {
		goto for_body67
	} else {
		goto for_end80
	}

for_body67:
	v35 = *libc.As[int32](i62)
	idxprom68 = int64(uint64(uint32(v35)))
	arrayidx69 = libc.Ptr(&ts_lex_map_37[idxprom68])
	v36 = *libc.As[int16](arrayidx69)
	conv70 = int32(uint32(uint16(v36)))
	v37 = *libc.As[int32](lookahead)
	cmp71 = conv70 == v37
	if cmp71 {
		goto if_then73
	} else {
		goto if_end77
	}

if_then73:
	v38 = *libc.As[int32](i62)
	add74 = v38 + 1
	idxprom75 = int64(uint64(uint32(add74)))
	arrayidx76 = libc.Ptr(&ts_lex_map_37[idxprom75])
	v39 = *libc.As[int16](arrayidx76)
	*libc.As[int16](state_addr) = v39
	goto next_state

if_end77:
	goto for_inc78

for_inc78:
	v40 = *libc.As[int32](i62)
	add79 = v40 + 2
	*libc.As[int32](i62) = add79
	goto for_cond63

for_end80:
	v41 = *libc.As[int32](lookahead)
	cmp81 = v41 != 0
	if cmp81 {
		goto land_lhs_true83
	} else {
		goto if_end96
	}

land_lhs_true83:
	v42 = *libc.As[int32](lookahead)
	cmp84 = v42 < 9
	if cmp84 {
		goto land_lhs_true89
	} else {
		goto lor_lhs_false86
	}

lor_lhs_false86:
	v43 = *libc.As[int32](lookahead)
	cmp87 = 13 < v43
	if cmp87 {
		goto land_lhs_true89
	} else {
		goto if_end96
	}

land_lhs_true89:
	v44 = *libc.As[int32](lookahead)
	cmp90 = v44 < 32
	if cmp90 {
		goto if_then95
	} else {
		goto lor_lhs_false92
	}

lor_lhs_false92:
	v45 = *libc.As[int32](lookahead)
	cmp93 = 34 < v45
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end96:
	v46 = *libc.As[byte](result)
	loadedv97 = (v46 & 1) != 0
	*libc.As[bool](retval) = loadedv97
	goto _return

sw_bb98:
	v47 = *libc.As[byte](eof)
	loadedv99 = (v47 & 1) != 0
	if loadedv99 {
		goto if_end104
	} else {
		goto land_lhs_true100
	}

land_lhs_true100:
	v48 = *libc.As[int32](lookahead)
	call101 = set_contains(libc.Ptr(&sym_uri_character_set_1), 12, v48)
	if call101 {
		goto if_then103
	} else {
		goto if_end104
	}

if_then103:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end104:
	v49 = *libc.As[byte](result)
	loadedv105 = (v49 & 1) != 0
	*libc.As[bool](retval) = loadedv105
	goto _return

sw_bb106:
	*libc.As[byte](result) = 1
	v50 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v50).F1)
	*libc.As[int16](result_symbol) = 0
	v51 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v51).F3)
	v52 = *libc.As[unsafe.Pointer](mark_end)
	v53 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v52)(v53)
	v54 = *libc.As[byte](result)
	loadedv107 = (v54 & 1) != 0
	*libc.As[bool](retval) = loadedv107
	goto _return

sw_bb108:
	*libc.As[byte](result) = 1
	v55 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol109 = libc.Ptr(&libc.As[TSLexer](v55).F1)
	*libc.As[int16](result_symbol109) = 1
	v56 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end110 = libc.Ptr(&libc.As[TSLexer](v56).F3)
	v57 = *libc.As[unsafe.Pointer](mark_end110)
	v58 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v57)(v58)
	v59 = *libc.As[byte](result)
	loadedv111 = (v59 & 1) != 0
	*libc.As[bool](retval) = loadedv111
	goto _return

sw_bb112:
	*libc.As[byte](result) = 1
	v60 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol113 = libc.Ptr(&libc.As[TSLexer](v60).F1)
	*libc.As[int16](result_symbol113) = 2
	v61 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end114 = libc.Ptr(&libc.As[TSLexer](v61).F3)
	v62 = *libc.As[unsafe.Pointer](mark_end114)
	v63 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v62)(v63)
	v64 = *libc.As[byte](result)
	loadedv115 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv115
	goto _return

sw_bb116:
	*libc.As[byte](result) = 1
	v65 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol117 = libc.Ptr(&libc.As[TSLexer](v65).F1)
	*libc.As[int16](result_symbol117) = 3
	v66 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end118 = libc.Ptr(&libc.As[TSLexer](v66).F3)
	v67 = *libc.As[unsafe.Pointer](mark_end118)
	v68 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v67)(v68)
	v69 = *libc.As[int32](lookahead)
	cmp119 = v69 != 0
	if cmp119 {
		goto land_lhs_true121
	} else {
		goto if_end128
	}

land_lhs_true121:
	v70 = *libc.As[int32](lookahead)
	cmp122 = v70 != 40
	if cmp122 {
		goto land_lhs_true124
	} else {
		goto if_end128
	}

land_lhs_true124:
	v71 = *libc.As[int32](lookahead)
	cmp125 = v71 != 41
	if cmp125 {
		goto if_then127
	} else {
		goto if_end128
	}

if_then127:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end128:
	v72 = *libc.As[byte](result)
	loadedv129 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv129
	goto _return

sw_bb130:
	*libc.As[byte](result) = 1
	v73 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol131 = libc.Ptr(&libc.As[TSLexer](v73).F1)
	*libc.As[int16](result_symbol131) = 4
	v74 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end132 = libc.Ptr(&libc.As[TSLexer](v74).F3)
	v75 = *libc.As[unsafe.Pointer](mark_end132)
	v76 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v75)(v76)
	v77 = *libc.As[byte](result)
	loadedv133 = (v77 & 1) != 0
	*libc.As[bool](retval) = loadedv133
	goto _return

sw_bb134:
	*libc.As[byte](result) = 1
	v78 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol135 = libc.Ptr(&libc.As[TSLexer](v78).F1)
	*libc.As[int16](result_symbol135) = 5
	v79 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end136 = libc.Ptr(&libc.As[TSLexer](v79).F3)
	v80 = *libc.As[unsafe.Pointer](mark_end136)
	v81 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v80)(v81)
	v82 = *libc.As[byte](result)
	loadedv137 = (v82 & 1) != 0
	*libc.As[bool](retval) = loadedv137
	goto _return

sw_bb138:
	*libc.As[byte](result) = 1
	v83 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol139 = libc.Ptr(&libc.As[TSLexer](v83).F1)
	*libc.As[int16](result_symbol139) = 6
	v84 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end140 = libc.Ptr(&libc.As[TSLexer](v84).F3)
	v85 = *libc.As[unsafe.Pointer](mark_end140)
	v86 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v85)(v86)
	*libc.As[int32](i141) = 0
	goto for_cond142

for_cond142:
	v87 = *libc.As[int32](i141)
	conv143 = int64(uint64(uint32(v87)))
	cmp144 = uint64(conv143) < uint64(28)
	if cmp144 {
		goto for_body146
	} else {
		goto for_end159
	}

for_body146:
	v88 = *libc.As[int32](i141)
	idxprom147 = int64(uint64(uint32(v88)))
	arrayidx148 = libc.Ptr(&ts_lex_map_38[idxprom147])
	v89 = *libc.As[int16](arrayidx148)
	conv149 = int32(uint32(uint16(v89)))
	v90 = *libc.As[int32](lookahead)
	cmp150 = conv149 == v90
	if cmp150 {
		goto if_then152
	} else {
		goto if_end156
	}

if_then152:
	v91 = *libc.As[int32](i141)
	add153 = v91 + 1
	idxprom154 = int64(uint64(uint32(add153)))
	arrayidx155 = libc.Ptr(&ts_lex_map_38[idxprom154])
	v92 = *libc.As[int16](arrayidx155)
	*libc.As[int16](state_addr) = v92
	goto next_state

if_end156:
	goto for_inc157

for_inc157:
	v93 = *libc.As[int32](i141)
	add158 = v93 + 2
	*libc.As[int32](i141) = add158
	goto for_cond142

for_end159:
	v94 = *libc.As[int32](lookahead)
	cmp160 = v94 != 0
	if cmp160 {
		goto land_lhs_true162
	} else {
		goto if_end175
	}

land_lhs_true162:
	v95 = *libc.As[int32](lookahead)
	cmp163 = v95 < 9
	if cmp163 {
		goto land_lhs_true168
	} else {
		goto lor_lhs_false165
	}

lor_lhs_false165:
	v96 = *libc.As[int32](lookahead)
	cmp166 = 13 < v96
	if cmp166 {
		goto land_lhs_true168
	} else {
		goto if_end175
	}

land_lhs_true168:
	v97 = *libc.As[int32](lookahead)
	cmp169 = v97 < 32
	if cmp169 {
		goto if_then174
	} else {
		goto lor_lhs_false171
	}

lor_lhs_false171:
	v98 = *libc.As[int32](lookahead)
	cmp172 = 34 < v98
	if cmp172 {
		goto if_then174
	} else {
		goto if_end175
	}

if_then174:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end175:
	v99 = *libc.As[byte](result)
	loadedv176 = (v99 & 1) != 0
	*libc.As[bool](retval) = loadedv176
	goto _return

sw_bb177:
	*libc.As[byte](result) = 1
	v100 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol178 = libc.Ptr(&libc.As[TSLexer](v100).F1)
	*libc.As[int16](result_symbol178) = 7
	v101 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end179 = libc.Ptr(&libc.As[TSLexer](v101).F3)
	v102 = *libc.As[unsafe.Pointer](mark_end179)
	v103 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v102)(v103)
	v104 = *libc.As[int32](lookahead)
	cmp180 = v104 == 58
	if cmp180 {
		goto if_then182
	} else {
		goto if_end183
	}

if_then182:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end183:
	v105 = *libc.As[int32](lookahead)
	cmp184 = v105 == 115
	if cmp184 {
		goto if_then186
	} else {
		goto if_end187
	}

if_then186:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end187:
	v106 = *libc.As[byte](eof)
	loadedv188 = (v106 & 1) != 0
	if loadedv188 {
		goto if_end193
	} else {
		goto land_lhs_true189
	}

land_lhs_true189:
	v107 = *libc.As[int32](lookahead)
	call190 = set_contains(libc.Ptr(&aux_sym__text_token1_character_set_1), 11, v107)
	if call190 {
		goto if_then192
	} else {
		goto if_end193
	}

if_then192:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end193:
	v108 = *libc.As[byte](result)
	loadedv194 = (v108 & 1) != 0
	*libc.As[bool](retval) = loadedv194
	goto _return

sw_bb195:
	*libc.As[byte](result) = 1
	v109 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol196 = libc.Ptr(&libc.As[TSLexer](v109).F1)
	*libc.As[int16](result_symbol196) = 7
	v110 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end197 = libc.Ptr(&libc.As[TSLexer](v110).F3)
	v111 = *libc.As[unsafe.Pointer](mark_end197)
	v112 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v111)(v112)
	v113 = *libc.As[int32](lookahead)
	cmp198 = v113 == 58
	if cmp198 {
		goto if_then200
	} else {
		goto if_end201
	}

if_then200:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end201:
	v114 = *libc.As[byte](eof)
	loadedv202 = (v114 & 1) != 0
	if loadedv202 {
		goto if_end207
	} else {
		goto land_lhs_true203
	}

land_lhs_true203:
	v115 = *libc.As[int32](lookahead)
	call204 = set_contains(libc.Ptr(&aux_sym__text_token1_character_set_1), 11, v115)
	if call204 {
		goto if_then206
	} else {
		goto if_end207
	}

if_then206:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end207:
	v116 = *libc.As[byte](result)
	loadedv208 = (v116 & 1) != 0
	*libc.As[bool](retval) = loadedv208
	goto _return

sw_bb209:
	*libc.As[byte](result) = 1
	v117 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol210 = libc.Ptr(&libc.As[TSLexer](v117).F1)
	*libc.As[int16](result_symbol210) = 7
	v118 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end211 = libc.Ptr(&libc.As[TSLexer](v118).F3)
	v119 = *libc.As[unsafe.Pointer](mark_end211)
	v120 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v119)(v120)
	v121 = *libc.As[int32](lookahead)
	cmp212 = v121 == 112
	if cmp212 {
		goto if_then214
	} else {
		goto if_end215
	}

if_then214:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end215:
	v122 = *libc.As[byte](eof)
	loadedv216 = (v122 & 1) != 0
	if loadedv216 {
		goto if_end221
	} else {
		goto land_lhs_true217
	}

land_lhs_true217:
	v123 = *libc.As[int32](lookahead)
	call218 = set_contains(libc.Ptr(&aux_sym__text_token1_character_set_1), 11, v123)
	if call218 {
		goto if_then220
	} else {
		goto if_end221
	}

if_then220:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end221:
	v124 = *libc.As[byte](result)
	loadedv222 = (v124 & 1) != 0
	*libc.As[bool](retval) = loadedv222
	goto _return

sw_bb223:
	*libc.As[byte](result) = 1
	v125 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol224 = libc.Ptr(&libc.As[TSLexer](v125).F1)
	*libc.As[int16](result_symbol224) = 7
	v126 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end225 = libc.Ptr(&libc.As[TSLexer](v126).F3)
	v127 = *libc.As[unsafe.Pointer](mark_end225)
	v128 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v127)(v128)
	v129 = *libc.As[int32](lookahead)
	cmp226 = v129 == 116
	if cmp226 {
		goto if_then228
	} else {
		goto if_end229
	}

if_then228:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end229:
	v130 = *libc.As[byte](eof)
	loadedv230 = (v130 & 1) != 0
	if loadedv230 {
		goto if_end235
	} else {
		goto land_lhs_true231
	}

land_lhs_true231:
	v131 = *libc.As[int32](lookahead)
	call232 = set_contains(libc.Ptr(&aux_sym__text_token1_character_set_1), 11, v131)
	if call232 {
		goto if_then234
	} else {
		goto if_end235
	}

if_then234:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end235:
	v132 = *libc.As[byte](result)
	loadedv236 = (v132 & 1) != 0
	*libc.As[bool](retval) = loadedv236
	goto _return

sw_bb237:
	*libc.As[byte](result) = 1
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol238 = libc.Ptr(&libc.As[TSLexer](v133).F1)
	*libc.As[int16](result_symbol238) = 7
	v134 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end239 = libc.Ptr(&libc.As[TSLexer](v134).F3)
	v135 = *libc.As[unsafe.Pointer](mark_end239)
	v136 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v135)(v136)
	v137 = *libc.As[int32](lookahead)
	cmp240 = v137 == 116
	if cmp240 {
		goto if_then242
	} else {
		goto if_end243
	}

if_then242:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end243:
	v138 = *libc.As[byte](eof)
	loadedv244 = (v138 & 1) != 0
	if loadedv244 {
		goto if_end249
	} else {
		goto land_lhs_true245
	}

land_lhs_true245:
	v139 = *libc.As[int32](lookahead)
	call246 = set_contains(libc.Ptr(&aux_sym__text_token1_character_set_1), 11, v139)
	if call246 {
		goto if_then248
	} else {
		goto if_end249
	}

if_then248:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end249:
	v140 = *libc.As[byte](result)
	loadedv250 = (v140 & 1) != 0
	*libc.As[bool](retval) = loadedv250
	goto _return

sw_bb251:
	*libc.As[byte](result) = 1
	v141 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol252 = libc.Ptr(&libc.As[TSLexer](v141).F1)
	*libc.As[int16](result_symbol252) = 7
	v142 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end253 = libc.Ptr(&libc.As[TSLexer](v142).F3)
	v143 = *libc.As[unsafe.Pointer](mark_end253)
	v144 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v143)(v144)
	v145 = *libc.As[byte](eof)
	loadedv254 = (v145 & 1) != 0
	if loadedv254 {
		goto if_end259
	} else {
		goto land_lhs_true255
	}

land_lhs_true255:
	v146 = *libc.As[int32](lookahead)
	call256 = set_contains(libc.Ptr(&aux_sym__text_token1_character_set_1), 11, v146)
	if call256 {
		goto if_then258
	} else {
		goto if_end259
	}

if_then258:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end259:
	v147 = *libc.As[byte](result)
	loadedv260 = (v147 & 1) != 0
	*libc.As[bool](retval) = loadedv260
	goto _return

sw_bb261:
	*libc.As[byte](result) = 1
	v148 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol262 = libc.Ptr(&libc.As[TSLexer](v148).F1)
	*libc.As[int16](result_symbol262) = 8
	v149 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end263 = libc.Ptr(&libc.As[TSLexer](v149).F3)
	v150 = *libc.As[unsafe.Pointer](mark_end263)
	v151 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v150)(v151)
	v152 = *libc.As[byte](result)
	loadedv264 = (v152 & 1) != 0
	*libc.As[bool](retval) = loadedv264
	goto _return

sw_bb265:
	*libc.As[byte](result) = 1
	v153 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol266 = libc.Ptr(&libc.As[TSLexer](v153).F1)
	*libc.As[int16](result_symbol266) = 9
	v154 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end267 = libc.Ptr(&libc.As[TSLexer](v154).F3)
	v155 = *libc.As[unsafe.Pointer](mark_end267)
	v156 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v155)(v156)
	v157 = *libc.As[byte](result)
	loadedv268 = (v157 & 1) != 0
	*libc.As[bool](retval) = loadedv268
	goto _return

sw_bb269:
	*libc.As[byte](result) = 1
	v158 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol270 = libc.Ptr(&libc.As[TSLexer](v158).F1)
	*libc.As[int16](result_symbol270) = 10
	v159 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end271 = libc.Ptr(&libc.As[TSLexer](v159).F3)
	v160 = *libc.As[unsafe.Pointer](mark_end271)
	v161 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v160)(v161)
	v162 = *libc.As[byte](result)
	loadedv272 = (v162 & 1) != 0
	*libc.As[bool](retval) = loadedv272
	goto _return

sw_bb273:
	*libc.As[byte](result) = 1
	v163 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol274 = libc.Ptr(&libc.As[TSLexer](v163).F1)
	*libc.As[int16](result_symbol274) = 11
	v164 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end275 = libc.Ptr(&libc.As[TSLexer](v164).F3)
	v165 = *libc.As[unsafe.Pointer](mark_end275)
	v166 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v165)(v166)
	v167 = *libc.As[byte](result)
	loadedv276 = (v167 & 1) != 0
	*libc.As[bool](retval) = loadedv276
	goto _return

sw_bb277:
	*libc.As[byte](result) = 1
	v168 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol278 = libc.Ptr(&libc.As[TSLexer](v168).F1)
	*libc.As[int16](result_symbol278) = 12
	v169 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end279 = libc.Ptr(&libc.As[TSLexer](v169).F3)
	v170 = *libc.As[unsafe.Pointer](mark_end279)
	v171 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v170)(v171)
	v172 = *libc.As[byte](result)
	loadedv280 = (v172 & 1) != 0
	*libc.As[bool](retval) = loadedv280
	goto _return

sw_bb281:
	*libc.As[byte](result) = 1
	v173 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol282 = libc.Ptr(&libc.As[TSLexer](v173).F1)
	*libc.As[int16](result_symbol282) = 13
	v174 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end283 = libc.Ptr(&libc.As[TSLexer](v174).F3)
	v175 = *libc.As[unsafe.Pointer](mark_end283)
	v176 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v175)(v176)
	v177 = *libc.As[byte](result)
	loadedv284 = (v177 & 1) != 0
	*libc.As[bool](retval) = loadedv284
	goto _return

sw_bb285:
	*libc.As[byte](result) = 1
	v178 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol286 = libc.Ptr(&libc.As[TSLexer](v178).F1)
	*libc.As[int16](result_symbol286) = 14
	v179 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end287 = libc.Ptr(&libc.As[TSLexer](v179).F3)
	v180 = *libc.As[unsafe.Pointer](mark_end287)
	v181 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v180)(v181)
	v182 = *libc.As[byte](result)
	loadedv288 = (v182 & 1) != 0
	*libc.As[bool](retval) = loadedv288
	goto _return

sw_bb289:
	*libc.As[byte](result) = 1
	v183 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol290 = libc.Ptr(&libc.As[TSLexer](v183).F1)
	*libc.As[int16](result_symbol290) = 15
	v184 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end291 = libc.Ptr(&libc.As[TSLexer](v184).F3)
	v185 = *libc.As[unsafe.Pointer](mark_end291)
	v186 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v185)(v186)
	v187 = *libc.As[byte](result)
	loadedv292 = (v187 & 1) != 0
	*libc.As[bool](retval) = loadedv292
	goto _return

sw_bb293:
	*libc.As[byte](result) = 1
	v188 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol294 = libc.Ptr(&libc.As[TSLexer](v188).F1)
	*libc.As[int16](result_symbol294) = 16
	v189 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end295 = libc.Ptr(&libc.As[TSLexer](v189).F3)
	v190 = *libc.As[unsafe.Pointer](mark_end295)
	v191 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v190)(v191)
	v192 = *libc.As[byte](result)
	loadedv296 = (v192 & 1) != 0
	*libc.As[bool](retval) = loadedv296
	goto _return

sw_bb297:
	*libc.As[byte](result) = 1
	v193 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol298 = libc.Ptr(&libc.As[TSLexer](v193).F1)
	*libc.As[int16](result_symbol298) = 17
	v194 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end299 = libc.Ptr(&libc.As[TSLexer](v194).F3)
	v195 = *libc.As[unsafe.Pointer](mark_end299)
	v196 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v195)(v196)
	v197 = *libc.As[byte](result)
	loadedv300 = (v197 & 1) != 0
	*libc.As[bool](retval) = loadedv300
	goto _return

sw_bb301:
	*libc.As[byte](result) = 1
	v198 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol302 = libc.Ptr(&libc.As[TSLexer](v198).F1)
	*libc.As[int16](result_symbol302) = 18
	v199 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end303 = libc.Ptr(&libc.As[TSLexer](v199).F3)
	v200 = *libc.As[unsafe.Pointer](mark_end303)
	v201 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v200)(v201)
	v202 = *libc.As[byte](result)
	loadedv304 = (v202 & 1) != 0
	*libc.As[bool](retval) = loadedv304
	goto _return

sw_bb305:
	*libc.As[byte](result) = 1
	v203 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol306 = libc.Ptr(&libc.As[TSLexer](v203).F1)
	*libc.As[int16](result_symbol306) = 19
	v204 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end307 = libc.Ptr(&libc.As[TSLexer](v204).F3)
	v205 = *libc.As[unsafe.Pointer](mark_end307)
	v206 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v205)(v206)
	v207 = *libc.As[byte](result)
	loadedv308 = (v207 & 1) != 0
	*libc.As[bool](retval) = loadedv308
	goto _return

sw_bb309:
	*libc.As[byte](result) = 1
	v208 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol310 = libc.Ptr(&libc.As[TSLexer](v208).F1)
	*libc.As[int16](result_symbol310) = 20
	v209 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end311 = libc.Ptr(&libc.As[TSLexer](v209).F3)
	v210 = *libc.As[unsafe.Pointer](mark_end311)
	v211 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v210)(v211)
	v212 = *libc.As[byte](result)
	loadedv312 = (v212 & 1) != 0
	*libc.As[bool](retval) = loadedv312
	goto _return

sw_bb313:
	*libc.As[byte](result) = 1
	v213 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol314 = libc.Ptr(&libc.As[TSLexer](v213).F1)
	*libc.As[int16](result_symbol314) = 21
	v214 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end315 = libc.Ptr(&libc.As[TSLexer](v214).F3)
	v215 = *libc.As[unsafe.Pointer](mark_end315)
	v216 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v215)(v216)
	v217 = *libc.As[byte](result)
	loadedv316 = (v217 & 1) != 0
	*libc.As[bool](retval) = loadedv316
	goto _return

sw_bb317:
	*libc.As[byte](result) = 1
	v218 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol318 = libc.Ptr(&libc.As[TSLexer](v218).F1)
	*libc.As[int16](result_symbol318) = 22
	v219 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end319 = libc.Ptr(&libc.As[TSLexer](v219).F3)
	v220 = *libc.As[unsafe.Pointer](mark_end319)
	v221 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v220)(v221)
	v222 = *libc.As[byte](result)
	loadedv320 = (v222 & 1) != 0
	*libc.As[bool](retval) = loadedv320
	goto _return

sw_bb321:
	*libc.As[byte](result) = 1
	v223 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol322 = libc.Ptr(&libc.As[TSLexer](v223).F1)
	*libc.As[int16](result_symbol322) = 23
	v224 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end323 = libc.Ptr(&libc.As[TSLexer](v224).F3)
	v225 = *libc.As[unsafe.Pointer](mark_end323)
	v226 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v225)(v226)
	v227 = *libc.As[byte](result)
	loadedv324 = (v227 & 1) != 0
	*libc.As[bool](retval) = loadedv324
	goto _return

sw_bb325:
	*libc.As[byte](result) = 1
	v228 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol326 = libc.Ptr(&libc.As[TSLexer](v228).F1)
	*libc.As[int16](result_symbol326) = 24
	v229 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end327 = libc.Ptr(&libc.As[TSLexer](v229).F3)
	v230 = *libc.As[unsafe.Pointer](mark_end327)
	v231 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v230)(v231)
	v232 = *libc.As[byte](result)
	loadedv328 = (v232 & 1) != 0
	*libc.As[bool](retval) = loadedv328
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v233 = *libc.As[bool](retval)
	return v233
}
func is_upper(c int32) bool {
	var cmp, cmp1, v2 bool
	var v0, v1 int32
	var c_addr, upper, lower unsafe.Pointer
	_, _, _, _, _, _, _, _ = c_addr, upper, lower, v0, cmp, v1, cmp1, v2

	c_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	upper = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	lower = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[int32](c_addr) = c
	*libc.As[int32](upper) = 65
	*libc.As[int32](lower) = 90
	v0 = *libc.As[int32](c_addr)
	cmp = v0 >= 65
	if cmp {
		goto land_rhs
	} else {
		v2 = false
		goto land_end
	}

land_rhs:
	v1 = *libc.As[int32](c_addr)
	cmp1 = v1 <= 90
	v2 = cmp1
	goto land_end

land_end:
	return v2
}
func parse_tagname(lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var call, loadedv, call3, call6, call8, v15, call11, call15, call17, cmp, call23, call25, lnot, v35, cmp30, cmp36, call39, cmp44, cmp50, call55, v64 bool
	var retval unsafe.Pointer
	var result_symbol unsafe.Pointer
	var v1, v5, v10, v12, v14, v17, v24, v26, v28, v30, v32, v34, v40, v45, v47, v51, inc, v52, v57, v62 int32
	var previous, user_length, lookahead, lookahead1, lookahead2, lookahead5, lookahead7, lookahead9, lookahead14, lookahead16, lookahead19, lookahead22, lookahead24, lookahead29, lookahead35, lookahead38, lookahead49, lookahead54 unsafe.Pointer
	var v3 byte
	var arrayidx unsafe.Pointer
	var v0, v2, v4, v6, v7, v8, v9, v11, v13, v16, v18, v19, v20, v21, v22, v23, v25, v27, v29, v31, v33, v36, v37, v38, v39, v41, v42, v43, v44, v46, v48, v49, v50, v53, v54, v55, v56, v58, v59, v60, v61, v63 unsafe.Pointer
	var lexer_addr, valid_symbols_addr, advance, advance10, mark_end, advance27, advance33, advance42, advance47, advance53 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, valid_symbols_addr, previous, user_length, v0, lookahead, v1, call, v2, arrayidx, v3, loadedv, v4, lookahead1, v5, v6, advance, v7, v8, v9, lookahead2, v10, call3, v11, lookahead5, v12, call6, v13, lookahead7, v14, call8, v15, v16, lookahead9, v17, v18, advance10, v19, v20, v21, mark_end, v22, v23, v24, call11, v25, lookahead14, v26, call15, v27, lookahead16, v28, call17, v29, lookahead19, v30, cmp, v31, lookahead22, v32, call23, v33, lookahead24, v34, call25, lnot, v35, v36, advance27, v37, v38, v39, lookahead29, v40, cmp30, v41, advance33, v42, v43, v44, lookahead35, v45, cmp36, v46, lookahead38, v47, call39, v48, advance42, v49, v50, v51, inc, v52, cmp44, v53, advance47, v54, v55, v56, lookahead49, v57, cmp50, v58, advance53, v59, v60, v61, lookahead54, v62, call55, v63, result_symbol, v64

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
	valid_symbols_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	previous = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	user_length = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v0).F0)
	v1 = *libc.As[int32](lookahead)
	call = is_upper(v1)
	if call {
		goto lor_lhs_false
	} else {
		goto if_then
	}

lor_lhs_false:
	v2 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx = v2
	v3 = *libc.As[byte](arrayidx)
	loadedv = (v3 & 1) != 0
	if loadedv {
		goto if_end
	} else {
		goto if_then
	}

if_then:
	*libc.As[bool](retval) = false
	goto _return

if_end:
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v4).F0)
	v5 = *libc.As[int32](lookahead1)
	*libc.As[int32](previous) = v5
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	advance = libc.Ptr(&libc.As[TSLexer](v6).F2)
	v7 = *libc.As[unsafe.Pointer](advance)
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v7)(v8, false)
	goto while_cond

while_cond:
	v9 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead2 = libc.Ptr(&libc.As[TSLexer](v9).F0)
	v10 = *libc.As[int32](lookahead2)
	call3 = is_upper(v10)
	if call3 {
		v15 = true
		goto lor_end
	} else {
		goto lor_lhs_false4
	}

lor_lhs_false4:
	v11 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead5 = libc.Ptr(&libc.As[TSLexer](v11).F0)
	v12 = *libc.As[int32](lookahead5)
	call6 = is_digit(v12)
	if call6 {
		v15 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead7 = libc.Ptr(&libc.As[TSLexer](v13).F0)
	v14 = *libc.As[int32](lookahead7)
	call8 = is_internal_char(v14)
	v15 = call8
	goto lor_end

lor_end:
	if v15 {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v16 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead9 = libc.Ptr(&libc.As[TSLexer](v16).F0)
	v17 = *libc.As[int32](lookahead9)
	*libc.As[int32](previous) = v17
	v18 = *libc.As[unsafe.Pointer](lexer_addr)
	advance10 = libc.Ptr(&libc.As[TSLexer](v18).F2)
	v19 = *libc.As[unsafe.Pointer](advance10)
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v19)(v20, false)
	goto while_cond

while_end:
	v21 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v21).F3)
	v22 = *libc.As[unsafe.Pointer](mark_end)
	v23 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v22)(v23)
	v24 = *libc.As[int32](previous)
	call11 = is_internal_char(v24)
	if call11 {
		goto if_then12
	} else {
		goto if_end13
	}

if_then12:
	*libc.As[bool](retval) = false
	goto _return

if_end13:
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead14 = libc.Ptr(&libc.As[TSLexer](v25).F0)
	v26 = *libc.As[int32](lookahead14)
	call15 = is_space(v26)
	if call15 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false18
	}

land_lhs_true:
	v27 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead16 = libc.Ptr(&libc.As[TSLexer](v27).F0)
	v28 = *libc.As[int32](lookahead16)
	call17 = is_newline(v28)
	if call17 {
		goto lor_lhs_false18
	} else {
		goto if_then20
	}

lor_lhs_false18:
	v29 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead19 = libc.Ptr(&libc.As[TSLexer](v29).F0)
	v30 = *libc.As[int32](lookahead19)
	cmp = v30 == 40
	if cmp {
		goto if_then20
	} else {
		goto if_end48
	}

if_then20:
	goto while_cond21

while_cond21:
	v31 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead22 = libc.Ptr(&libc.As[TSLexer](v31).F0)
	v32 = *libc.As[int32](lookahead22)
	call23 = is_space(v32)
	if call23 {
		goto land_rhs
	} else {
		v35 = false
		goto land_end
	}

land_rhs:
	v33 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead24 = libc.Ptr(&libc.As[TSLexer](v33).F0)
	v34 = *libc.As[int32](lookahead24)
	call25 = is_newline(v34)
	lnot = call25 != true
	v35 = lnot
	goto land_end

land_end:
	if v35 {
		goto while_body26
	} else {
		goto while_end28
	}

while_body26:
	v36 = *libc.As[unsafe.Pointer](lexer_addr)
	advance27 = libc.Ptr(&libc.As[TSLexer](v36).F2)
	v37 = *libc.As[unsafe.Pointer](advance27)
	v38 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v37)(v38, false)
	goto while_cond21

while_end28:
	v39 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead29 = libc.Ptr(&libc.As[TSLexer](v39).F0)
	v40 = *libc.As[int32](lookahead29)
	cmp30 = v40 != 40
	if cmp30 {
		goto if_then31
	} else {
		goto if_end32
	}

if_then31:
	*libc.As[bool](retval) = false
	goto _return

if_end32:
	v41 = *libc.As[unsafe.Pointer](lexer_addr)
	advance33 = libc.Ptr(&libc.As[TSLexer](v41).F2)
	v42 = *libc.As[unsafe.Pointer](advance33)
	v43 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v42)(v43, false)
	*libc.As[int32](user_length) = 0
	goto while_cond34

while_cond34:
	v44 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead35 = libc.Ptr(&libc.As[TSLexer](v44).F0)
	v45 = *libc.As[int32](lookahead35)
	cmp36 = v45 != 41
	if cmp36 {
		goto while_body37
	} else {
		goto while_end43
	}

while_body37:
	v46 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead38 = libc.Ptr(&libc.As[TSLexer](v46).F0)
	v47 = *libc.As[int32](lookahead38)
	call39 = is_newline(v47)
	if call39 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*libc.As[bool](retval) = false
	goto _return

if_end41:
	v48 = *libc.As[unsafe.Pointer](lexer_addr)
	advance42 = libc.Ptr(&libc.As[TSLexer](v48).F2)
	v49 = *libc.As[unsafe.Pointer](advance42)
	v50 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v49)(v50, false)
	v51 = *libc.As[int32](user_length)
	inc = v51 + 1
	*libc.As[int32](user_length) = inc
	goto while_cond34

while_end43:
	v52 = *libc.As[int32](user_length)
	cmp44 = v52 <= 0
	if cmp44 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*libc.As[bool](retval) = false
	goto _return

if_end46:
	v53 = *libc.As[unsafe.Pointer](lexer_addr)
	advance47 = libc.Ptr(&libc.As[TSLexer](v53).F2)
	v54 = *libc.As[unsafe.Pointer](advance47)
	v55 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v54)(v55, false)
	goto if_end48

if_end48:
	v56 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead49 = libc.Ptr(&libc.As[TSLexer](v56).F0)
	v57 = *libc.As[int32](lookahead49)
	cmp50 = v57 != 58
	if cmp50 {
		goto if_then51
	} else {
		goto if_end52
	}

if_then51:
	*libc.As[bool](retval) = false
	goto _return

if_end52:
	v58 = *libc.As[unsafe.Pointer](lexer_addr)
	advance53 = libc.Ptr(&libc.As[TSLexer](v58).F2)
	v59 = *libc.As[unsafe.Pointer](advance53)
	v60 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v59)(v60, false)
	v61 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead54 = libc.Ptr(&libc.As[TSLexer](v61).F0)
	v62 = *libc.As[int32](lookahead54)
	call55 = is_space(v62)
	if call55 {
		goto if_end57
	} else {
		goto if_then56
	}

if_then56:
	*libc.As[bool](retval) = false
	goto _return

if_end57:
	v63 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v63).F1)
	*libc.As[int16](result_symbol) = 0
	*libc.As[bool](retval) = true
	goto _return

_return:
	v64 = *libc.As[bool](retval)
	return v64
}
func is_digit(c int32) bool {
	var cmp, cmp1, v2 bool
	var v0, v1 int32
	var c_addr, upper, lower unsafe.Pointer
	_, _, _, _, _, _, _, _ = c_addr, upper, lower, v0, cmp, v1, cmp1, v2

	c_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	upper = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	lower = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[int32](c_addr) = c
	*libc.As[int32](upper) = 48
	*libc.As[int32](lower) = 57
	v0 = *libc.As[int32](c_addr)
	cmp = v0 >= 48
	if cmp {
		goto land_rhs
	} else {
		v2 = false
		goto land_end
	}

land_rhs:
	v1 = *libc.As[int32](c_addr)
	cmp1 = v1 <= 57
	v2 = cmp1
	goto land_end

land_end:
	return v2
}
func is_internal_char(c int32) bool {
	var valid_chars unsafe.Pointer
	var cmp, cmp1, v5 bool
	var retval unsafe.Pointer
	var v0, v1, v2, v3, v4, inc int32
	var c_addr, length, i, arrayidx unsafe.Pointer
	var idxprom int64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, c_addr, valid_chars, length, i, v0, cmp, v1, v2, idxprom, arrayidx, v3, cmp1, v4, inc, v5

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	c_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	valid_chars = libc.Ptr(&new(struct {
		_ [0]uint64
		v [2]int32
		b byte
	}).v)
	length = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[int32](c_addr) = c
	libc.Memmove(libc.As[byte](valid_chars), libc.As[byte](libc.Ptr(&__const_is_internal_char_valid_chars)), int64(8))
	*libc.As[int32](length) = 2
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v0 = *libc.As[int32](i)
	cmp = v0 < 2
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v1 = *libc.As[int32](c_addr)
	v2 = *libc.As[int32](i)
	idxprom = int64(v2)
	arrayidx = libc.Ptr(&libc.As[[2]int32](valid_chars)[idxprom])
	v3 = *libc.As[int32](arrayidx)
	cmp1 = v1 == v3
	if cmp1 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = true
	goto _return

if_end:
	goto for_inc

for_inc:
	v4 = *libc.As[int32](i)
	inc = v4 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v5 = *libc.As[bool](retval)
	return v5
}
func is_space(c int32) bool {
	var space_chars unsafe.Pointer
	var cmp, cmp1, loadedv, call, v7 bool
	var v0, v1, v2, v3, v4, inc, v6 int32
	var c_addr, length, i, arrayidx unsafe.Pointer
	var idxprom int64
	var v5 byte
	var is_space_char unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c_addr, space_chars, length, is_space_char, i, v0, cmp, v1, v2, idxprom, arrayidx, v3, cmp1, v4, inc, v5, loadedv, v6, call, v7

	c_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	space_chars = libc.Ptr(&new(struct {
		_ [0]uint64
		v [4]int32
		b byte
	}).v)
	length = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	is_space_char = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[int32](c_addr) = c
	libc.Memmove(libc.As[byte](space_chars), libc.As[byte](libc.Ptr(&__const_is_space_space_chars)), int64(16))
	*libc.As[int32](length) = 4
	*libc.As[byte](is_space_char) = 0
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v0 = *libc.As[int32](i)
	cmp = v0 < 4
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v1 = *libc.As[int32](c_addr)
	v2 = *libc.As[int32](i)
	idxprom = int64(v2)
	arrayidx = libc.Ptr(&libc.As[[4]int32](space_chars)[idxprom])
	v3 = *libc.As[int32](arrayidx)
	cmp1 = v1 == v3
	if cmp1 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[byte](is_space_char) = 1
	goto for_end

if_end:
	goto for_inc

for_inc:
	v4 = *libc.As[int32](i)
	inc = v4 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	v5 = *libc.As[byte](is_space_char)
	loadedv = (v5 & 1) != 0
	if loadedv {
		v7 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v6 = *libc.As[int32](c_addr)
	call = is_newline(v6)
	v7 = call
	goto lor_end

lor_end:
	return v7
}
func is_newline(c int32) bool {
	var newline_chars unsafe.Pointer
	var cmp, cmp1, v5 bool
	var retval unsafe.Pointer
	var v0, v1, v2, v3, v4, inc int32
	var c_addr, length, i, arrayidx unsafe.Pointer
	var idxprom int64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, c_addr, newline_chars, length, i, v0, cmp, v1, v2, idxprom, arrayidx, v3, cmp1, v4, inc, v5

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	c_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	newline_chars = libc.Ptr(&new(struct {
		_ [0]uint64
		v [3]int32
		b byte
	}).v)
	length = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[int32](c_addr) = c
	libc.Memmove(libc.As[byte](newline_chars), libc.As[byte](libc.Ptr(&__const_is_newline_newline_chars)), int64(12))
	*libc.As[int32](length) = 3
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v0 = *libc.As[int32](i)
	cmp = v0 < 3
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v1 = *libc.As[int32](c_addr)
	v2 = *libc.As[int32](i)
	idxprom = int64(v2)
	arrayidx = libc.Ptr(&libc.As[[3]int32](newline_chars)[idxprom])
	v3 = *libc.As[int32](arrayidx)
	cmp1 = v1 == v3
	if cmp1 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[bool](retval) = true
	goto _return

if_end:
	goto for_inc

for_inc:
	v4 = *libc.As[int32](i)
	inc = v4 + 1
	*libc.As[int32](i) = inc
	goto for_cond

for_end:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v5 = *libc.As[bool](retval)
	return v5
}
func set_contains(ranges unsafe.Pointer, len int32, lookahead int32) bool {
	var arrayidx, arrayidx10 unsafe.Pointer
	var cmp, cmp1, cmp2, cmp4, cmp12, cmp14, v28, v29 bool
	var retval unsafe.Pointer
	var v0, v1, sub, v2, v3, div, v4, v5, add, v7, v8, v10, v11, v13, v14, v16, v17, v18, v19, sub7, v21, v22, v24, v25, v27 int32
	var len_addr, lookahead_addr, index, size, half_size, mid_index, start, end, end3, start11, end13 unsafe.Pointer
	var idxprom, idxprom9 int64
	var v6, v9, v12, v15, v20, v23, v26 unsafe.Pointer
	var ranges_addr, _range, range8 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, ranges_addr, len_addr, lookahead_addr, index, size, half_size, mid_index, _range, range8, v0, v1, sub, v2, cmp, v3, div, v4, v5, add, v6, v7, idxprom, arrayidx, v8, v9, start, v10, cmp1, v11, v12, end, v13, cmp2, v14, v15, end3, v16, cmp4, v17, v18, v19, sub7, v20, v21, idxprom9, arrayidx10, v22, v23, start11, v24, cmp12, v25, v26, end13, v27, cmp14, v28, v29

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	ranges_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	len_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	lookahead_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	index = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	size = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	half_size = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	mid_index = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	_range = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	range8 = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](ranges_addr) = ranges
	*libc.As[int32](len_addr) = len
	*libc.As[int32](lookahead_addr) = lookahead
	*libc.As[int32](index) = 0
	v0 = *libc.As[int32](len_addr)
	v1 = *libc.As[int32](index)
	sub = v0 - v1
	*libc.As[int32](size) = sub
	goto while_cond

while_cond:
	v2 = *libc.As[int32](size)
	cmp = uint32(v2) > 1
	if cmp {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v3 = *libc.As[int32](size)
	div = int32(uint32(v3) / 2)
	*libc.As[int32](half_size) = div
	v4 = *libc.As[int32](index)
	v5 = *libc.As[int32](half_size)
	add = v4 + v5
	*libc.As[int32](mid_index) = add
	v6 = *libc.As[unsafe.Pointer](ranges_addr)
	v7 = *libc.As[int32](mid_index)
	idxprom = int64(uint64(uint32(v7)))
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v6), int(idxprom)*8))
	*libc.As[unsafe.Pointer](_range) = arrayidx
	v8 = *libc.As[int32](lookahead_addr)
	v9 = *libc.As[unsafe.Pointer](_range)
	start = libc.Ptr(&libc.As[TSCharacterRange](v9).F0)
	v10 = *libc.As[int32](start)
	cmp1 = v8 >= v10
	if cmp1 {
		goto land_lhs_true
	} else {
		goto if_else
	}

land_lhs_true:
	v11 = *libc.As[int32](lookahead_addr)
	v12 = *libc.As[unsafe.Pointer](_range)
	end = libc.Ptr(&libc.As[TSCharacterRange](v12).F1)
	v13 = *libc.As[int32](end)
	cmp2 = v11 <= v13
	if cmp2 {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	*libc.As[bool](retval) = true
	goto _return

if_else:
	v14 = *libc.As[int32](lookahead_addr)
	v15 = *libc.As[unsafe.Pointer](_range)
	end3 = libc.Ptr(&libc.As[TSCharacterRange](v15).F1)
	v16 = *libc.As[int32](end3)
	cmp4 = v14 > v16
	if cmp4 {
		goto if_then5
	} else {
		goto if_end
	}

if_then5:
	v17 = *libc.As[int32](mid_index)
	*libc.As[int32](index) = v17
	goto if_end

if_end:
	goto if_end6

if_end6:
	v18 = *libc.As[int32](half_size)
	v19 = *libc.As[int32](size)
	sub7 = v19 - v18
	*libc.As[int32](size) = sub7
	goto while_cond

while_end:
	v20 = *libc.As[unsafe.Pointer](ranges_addr)
	v21 = *libc.As[int32](index)
	idxprom9 = int64(uint64(uint32(v21)))
	arrayidx10 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v20), int(idxprom9)*8))
	*libc.As[unsafe.Pointer](range8) = arrayidx10
	v22 = *libc.As[int32](lookahead_addr)
	v23 = *libc.As[unsafe.Pointer](range8)
	start11 = libc.Ptr(&libc.As[TSCharacterRange](v23).F0)
	v24 = *libc.As[int32](start11)
	cmp12 = v22 >= v24
	if cmp12 {
		goto land_rhs
	} else {
		v28 = false
		goto land_end
	}

land_rhs:
	v25 = *libc.As[int32](lookahead_addr)
	v26 = *libc.As[unsafe.Pointer](range8)
	end13 = libc.Ptr(&libc.As[TSCharacterRange](v26).F1)
	v27 = *libc.As[int32](end13)
	cmp14 = v25 <= v27
	v28 = cmp14
	goto land_end

land_end:
	*libc.As[bool](retval) = v28
	goto _return

_return:
	v29 = *libc.As[bool](retval)
	return v29
}
