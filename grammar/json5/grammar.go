package grammar_json5

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

var tree_sitter_json5_language TSLanguage = TSLanguage{15, 21, 0, 14, 0, 31, 7, 2, 2, 5, &ts_parse_table[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[77]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], &ts_field_names[0], &ts_field_map_slices[0], &ts_field_map_entries[0], &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{0, 1, 0}}
var ts_parse_table [7][21]int16 = [7][21]int16{[21]int16{1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}, [21]int16{0, 3, 5, 0, 0, 0, 0, 7, 0, 9, 9, 9, 9, 9, 29, 30, 0, 30, 30, 0, 0}, [21]int16{0, 3, 5, 0, 0, 0, 0, 7, 11, 13, 13, 13, 13, 13, 0, 19, 0, 19, 19, 0, 0}, [21]int16{0, 3, 5, 0, 0, 0, 0, 7, 15, 17, 17, 17, 17, 17, 0, 25, 0, 25, 25, 0, 0}, [21]int16{0, 3, 5, 0, 0, 0, 0, 7, 19, 17, 17, 17, 17, 17, 0, 25, 0, 25, 25, 0, 0}, [21]int16{0, 3, 5, 0, 0, 0, 0, 7, 0, 21, 21, 21, 21, 21, 0, 26, 0, 26, 26, 0, 0}, [21]int16{0, 3, 5, 0, 0, 0, 0, 7, 0, 17, 17, 17, 17, 17, 0, 25, 0, 25, 25, 0, 0}}
var ts_small_parse_table [256]int16 = [256]int16{4, 3, 1, 1, 23, 1, 4, 18, 1, 16, 25, 2, 6, 9, 2, 3, 1, 1, 27, 4, 0, 3, 4, 8, 4, 3, 1, 1, 29, 1, 4, 27, 1, 16, 25, 2, 6, 9, 2, 3, 1, 1, 31, 4, 0, 3, 4, 8, 2, 3, 1, 1, 33, 4, 0, 3, 4, 8, 2, 3, 1, 1, 35, 4, 0, 3, 4, 8, 4, 3, 1, 1, 37, 1, 4, 27, 1, 16, 25, 2, 6, 9, 2, 3, 1, 1, 39, 4, 0, 3, 4, 8, 2, 3, 1, 1, 41, 4, 0, 3, 4, 8, 2, 3, 1, 1, 43, 4, 0, 3, 4, 8, 2, 3, 1, 1, 45, 4, 0, 3, 4, 8, 4, 3, 1, 1, 47, 1, 3, 49, 1, 4, 20, 1, 19, 4, 3, 1, 1, 51, 1, 3, 53, 1, 8, 21, 1, 20, 4, 3, 1, 1, 29, 1, 4, 55, 1, 3, 22, 1, 19, 4, 3, 1, 1, 15, 1, 8, 57, 1, 3, 23, 1, 20, 4, 3, 1, 1, 59, 1, 3, 62, 1, 4, 22, 1, 19, 4, 3, 1, 1, 64, 1, 3, 67, 1, 8, 23, 1, 20, 3, 3, 1, 1, 27, 1, 16, 25, 2, 6, 9, 2, 3, 1, 1, 67, 2, 3, 8, 2, 3, 1, 1, 69, 2, 3, 4, 2, 3, 1, 1, 62, 2, 3, 4, 2, 3, 1, 1, 71, 1, 5, 2, 3, 1, 1, 73, 1, 0, 2, 3, 1, 1, 75, 1, 0}
var ts_small_parse_table_map [24]int32 = [24]int32{0, 14, 24, 38, 48, 58, 68, 82, 92, 102, 112, 122, 135, 148, 161, 174, 187, 200, 211, 219, 227, 235, 242, 249}
var ts_symbol_names [21]*byte = [21]*byte{&_str_2[0], &_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0], &_str_17[0], &_str_18[0], &_str_19[0], &_str_20[0], &_str_21[0], &_str_22[0]}
var ts_field_names [3]*byte = [3]*byte{nil, &_str_23[0], &_str_24[0]}
var ts_field_map_slices [2]TSMapSlice = [2]TSMapSlice{TSMapSlice{}, TSMapSlice{0, 2}}
var ts_field_map_entries [2]TSFieldMapEntry = [2]TSFieldMapEntry{TSFieldMapEntry{1, 0, 0}, TSFieldMapEntry{2, 2, 0}}
var ts_symbol_metadata [21]TSSymbolMetadata = [21]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [21]int16 = [21]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [2][5]int16 = [2][5]int16{}
var ts_lex_modes [31]TSLexerMode = [31]TSLexerMode{TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{}, TSLexerMode{6, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{6, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{6, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}}
var ts_primary_state_ids [31]int16 = [31]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
var _str [6]byte = [6]byte{106, 115, 111, 110, 53, 0}
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
	F4 struct {
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
	F16 struct {
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
	F17 struct {
		F0 anon_1
		F1 [6]byte
	}
	F18 struct {
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
	F19 struct {
		F0 anon_1
		F1 [6]byte
	}
	F20 struct {
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
	F21 struct {
		F0 anon_1
		F1 [6]byte
	}
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
	F24 struct {
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
	F25 struct {
		F0 anon_1
		F1 [6]byte
	}
	F26 struct {
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
	F27 struct {
		F0 anon_1
		F1 [6]byte
	}
	F28 TSParseActionEntry
	F29 struct {
		F0 anon_1
		F1 [6]byte
	}
	F30 struct {
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
	F31 struct {
		F0 anon_1
		F1 [6]byte
	}
	F32 TSParseActionEntry
	F33 struct {
		F0 anon_1
		F1 [6]byte
	}
	F34 TSParseActionEntry
	F35 struct {
		F0 anon_1
		F1 [6]byte
	}
	F36 TSParseActionEntry
	F37 struct {
		F0 anon_1
		F1 [6]byte
	}
	F38 struct {
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
	F39 struct {
		F0 anon_1
		F1 [6]byte
	}
	F40 TSParseActionEntry
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
		F0 anon_1
		F1 [6]byte
	}
	F46 TSParseActionEntry
	F47 struct {
		F0 anon_1
		F1 [6]byte
	}
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
	F60 TSParseActionEntry
	F61 struct {
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
	F62 struct {
		F0 anon_1
		F1 [6]byte
	}
	F63 TSParseActionEntry
	F64 struct {
		F0 anon_1
		F1 [6]byte
	}
	F65 TSParseActionEntry
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
	F68 TSParseActionEntry
	F69 struct {
		F0 anon_1
		F1 [6]byte
	}
	F70 TSParseActionEntry
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
	F74 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F75 struct {
		F0 anon_1
		F1 [6]byte
	}
	F76 TSParseActionEntry
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
	F4 struct {
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
	F16 struct {
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
	F17 struct {
		F0 anon_1
		F1 [6]byte
	}
	F18 struct {
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
	F19 struct {
		F0 anon_1
		F1 [6]byte
	}
	F20 struct {
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
	F21 struct {
		F0 anon_1
		F1 [6]byte
	}
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
	F24 struct {
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
	F25 struct {
		F0 anon_1
		F1 [6]byte
	}
	F26 struct {
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
	F27 struct {
		F0 anon_1
		F1 [6]byte
	}
	F28 TSParseActionEntry
	F29 struct {
		F0 anon_1
		F1 [6]byte
	}
	F30 struct {
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
	F31 struct {
		F0 anon_1
		F1 [6]byte
	}
	F32 TSParseActionEntry
	F33 struct {
		F0 anon_1
		F1 [6]byte
	}
	F34 TSParseActionEntry
	F35 struct {
		F0 anon_1
		F1 [6]byte
	}
	F36 TSParseActionEntry
	F37 struct {
		F0 anon_1
		F1 [6]byte
	}
	F38 struct {
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
	F39 struct {
		F0 anon_1
		F1 [6]byte
	}
	F40 TSParseActionEntry
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
		F0 anon_1
		F1 [6]byte
	}
	F46 TSParseActionEntry
	F47 struct {
		F0 anon_1
		F1 [6]byte
	}
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
	F60 TSParseActionEntry
	F61 struct {
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
	F62 struct {
		F0 anon_1
		F1 [6]byte
	}
	F63 TSParseActionEntry
	F64 struct {
		F0 anon_1
		F1 [6]byte
	}
	F65 TSParseActionEntry
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
	F68 TSParseActionEntry
	F69 struct {
		F0 anon_1
		F1 [6]byte
	}
	F70 TSParseActionEntry
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
	F74 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F75 struct {
		F0 anon_1
		F1 [6]byte
	}
	F76 TSParseActionEntry
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
}{0, 0, 1, 0}, [2]byte{}}}, struct {
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
}{0, 2, 0, 0}, [2]byte{}}}, struct {
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
}{0, 30, 0, 0}, [2]byte{}}}, struct {
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
}{0, 10, 0, 0}, [2]byte{}}}, struct {
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
}{0, 15, 0, 0}, [2]byte{}}}, struct {
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
}{0, 8, 0, 0}, [2]byte{}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 15, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 17, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 17, 0, 0}}}, struct {
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
}{0, 11, 0, 0}, [2]byte{}}}, struct {
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
}{0, 12, 0, 0}, [2]byte{}}}, struct {
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
}{0, 13, 0, 0}, [2]byte{}}}, struct {
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
}{0, 4, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 19, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 19, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 20, 0, 0}}}, struct {
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
}{0, 6, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 20, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 16, 0, 1}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 14, 0, 0}}}}
var _str_2 [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_4 [2]byte = [2]byte{123, 0}
var _str_5 [2]byte = [2]byte{44, 0}
var _str_6 [2]byte = [2]byte{125, 0}
var _str_7 [2]byte = [2]byte{58, 0}
var _str_8 [11]byte = [11]byte{105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}
var _str_9 [2]byte = [2]byte{91, 0}
var _str_10 [2]byte = [2]byte{93, 0}
var _str_11 [7]byte = [7]byte{115, 116, 114, 105, 110, 103, 0}
var _str_12 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}
var _str_13 [5]byte = [5]byte{110, 117, 108, 108, 0}
var _str_14 [5]byte = [5]byte{116, 114, 117, 101, 0}
var _str_15 [6]byte = [6]byte{102, 97, 108, 115, 101, 0}
var _str_16 [5]byte = [5]byte{102, 105, 108, 101, 0}
var _str_17 [7]byte = [7]byte{111, 98, 106, 101, 99, 116, 0}
var _str_18 [7]byte = [7]byte{109, 101, 109, 98, 101, 114, 0}
var _str_19 [6]byte = [6]byte{97, 114, 114, 97, 121, 0}
var _str_20 [7]byte = [7]byte{95, 118, 97, 108, 117, 101, 0}
var _str_21 [15]byte = [15]byte{111, 98, 106, 101, 99, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_22 [14]byte = [14]byte{97, 114, 114, 97, 121, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_23 [5]byte = [5]byte{110, 97, 109, 101, 0}
var _str_24 [6]byte = [6]byte{118, 97, 108, 117, 101, 0}
var ts_lex_map [36]int16 = [36]int16{34, 7, 39, 8, 44, 47, 46, 76, 47, 9, 48, 74, 58, 49, 73, 61, 78, 51, 91, 70, 93, 71, 102, 52, 110, 67, 116, 63, 123, 46, 125, 48, 43, 12, 45, 12}
var sym_identifier_character_set_2 [680]TSCharacterRange = [680]TSCharacterRange{TSCharacterRange{36, 36}, TSCharacterRange{48, 57}, TSCharacterRange{65, 90}, TSCharacterRange{95, 95}, TSCharacterRange{97, 122}, TSCharacterRange{170, 170}, TSCharacterRange{181, 181}, TSCharacterRange{186, 186}, TSCharacterRange{192, 214}, TSCharacterRange{216, 246}, TSCharacterRange{248, 705}, TSCharacterRange{710, 721}, TSCharacterRange{736, 740}, TSCharacterRange{748, 748}, TSCharacterRange{750, 750}, TSCharacterRange{880, 884}, TSCharacterRange{886, 887}, TSCharacterRange{890, 893}, TSCharacterRange{895, 895}, TSCharacterRange{902, 902}, TSCharacterRange{904, 906}, TSCharacterRange{908, 908}, TSCharacterRange{910, 929}, TSCharacterRange{931, 1013}, TSCharacterRange{1015, 1153}, TSCharacterRange{1162, 1327}, TSCharacterRange{1329, 1366}, TSCharacterRange{1369, 1369}, TSCharacterRange{1376, 1416}, TSCharacterRange{1488, 1514}, TSCharacterRange{1519, 1522}, TSCharacterRange{1568, 1610}, TSCharacterRange{1646, 1647}, TSCharacterRange{1649, 1747}, TSCharacterRange{1749, 1749}, TSCharacterRange{1765, 1766}, TSCharacterRange{1774, 1775}, TSCharacterRange{1786, 1788}, TSCharacterRange{1791, 1791}, TSCharacterRange{1808, 1808}, TSCharacterRange{1810, 1839}, TSCharacterRange{1869, 1957}, TSCharacterRange{1969, 1969}, TSCharacterRange{1994, 2026}, TSCharacterRange{2036, 2037}, TSCharacterRange{2042, 2042}, TSCharacterRange{2048, 2069}, TSCharacterRange{2074, 2074}, TSCharacterRange{2084, 2084}, TSCharacterRange{2088, 2088}, TSCharacterRange{2112, 2136}, TSCharacterRange{2144, 2154}, TSCharacterRange{2160, 2183}, TSCharacterRange{2185, 2190}, TSCharacterRange{2208, 2249}, TSCharacterRange{2308, 2361}, TSCharacterRange{2365, 2365}, TSCharacterRange{2384, 2384}, TSCharacterRange{2392, 2401}, TSCharacterRange{2417, 2432}, TSCharacterRange{2437, 2444}, TSCharacterRange{2447, 2448}, TSCharacterRange{2451, 2472}, TSCharacterRange{2474, 2480}, TSCharacterRange{2482, 2482}, TSCharacterRange{2486, 2489}, TSCharacterRange{2493, 2493}, TSCharacterRange{2510, 2510}, TSCharacterRange{2524, 2525}, TSCharacterRange{2527, 2529}, TSCharacterRange{2544, 2545}, TSCharacterRange{2556, 2556}, TSCharacterRange{2565, 2570}, TSCharacterRange{2575, 2576}, TSCharacterRange{2579, 2600}, TSCharacterRange{2602, 2608}, TSCharacterRange{2610, 2611}, TSCharacterRange{2613, 2614}, TSCharacterRange{2616, 2617}, TSCharacterRange{2649, 2652}, TSCharacterRange{2654, 2654}, TSCharacterRange{2674, 2676}, TSCharacterRange{2693, 2701}, TSCharacterRange{2703, 2705}, TSCharacterRange{2707, 2728}, TSCharacterRange{2730, 2736}, TSCharacterRange{2738, 2739}, TSCharacterRange{2741, 2745}, TSCharacterRange{2749, 2749}, TSCharacterRange{2768, 2768}, TSCharacterRange{2784, 2785}, TSCharacterRange{2809, 2809}, TSCharacterRange{2821, 2828}, TSCharacterRange{2831, 2832}, TSCharacterRange{2835, 2856}, TSCharacterRange{2858, 2864}, TSCharacterRange{2866, 2867}, TSCharacterRange{2869, 2873}, TSCharacterRange{2877, 2877}, TSCharacterRange{2908, 2909}, TSCharacterRange{2911, 2913}, TSCharacterRange{2929, 2929}, TSCharacterRange{2947, 2947}, TSCharacterRange{2949, 2954}, TSCharacterRange{2958, 2960}, TSCharacterRange{2962, 2965}, TSCharacterRange{2969, 2970}, TSCharacterRange{2972, 2972}, TSCharacterRange{2974, 2975}, TSCharacterRange{2979, 2980}, TSCharacterRange{2984, 2986}, TSCharacterRange{2990, 3001}, TSCharacterRange{3024, 3024}, TSCharacterRange{3077, 3084}, TSCharacterRange{3086, 3088}, TSCharacterRange{3090, 3112}, TSCharacterRange{3114, 3129}, TSCharacterRange{3133, 3133}, TSCharacterRange{3160, 3162}, TSCharacterRange{3165, 3165}, TSCharacterRange{3168, 3169}, TSCharacterRange{3200, 3200}, TSCharacterRange{3205, 3212}, TSCharacterRange{3214, 3216}, TSCharacterRange{3218, 3240}, TSCharacterRange{3242, 3251}, TSCharacterRange{3253, 3257}, TSCharacterRange{3261, 3261}, TSCharacterRange{3293, 3294}, TSCharacterRange{3296, 3297}, TSCharacterRange{3313, 3314}, TSCharacterRange{3332, 3340}, TSCharacterRange{3342, 3344}, TSCharacterRange{3346, 3386}, TSCharacterRange{3389, 3389}, TSCharacterRange{3406, 3406}, TSCharacterRange{3412, 3414}, TSCharacterRange{3423, 3425}, TSCharacterRange{3450, 3455}, TSCharacterRange{3461, 3478}, TSCharacterRange{3482, 3505}, TSCharacterRange{3507, 3515}, TSCharacterRange{3517, 3517}, TSCharacterRange{3520, 3526}, TSCharacterRange{3585, 3632}, TSCharacterRange{3634, 3635}, TSCharacterRange{3648, 3654}, TSCharacterRange{3713, 3714}, TSCharacterRange{3716, 3716}, TSCharacterRange{3718, 3722}, TSCharacterRange{3724, 3747}, TSCharacterRange{3749, 3749}, TSCharacterRange{3751, 3760}, TSCharacterRange{3762, 3763}, TSCharacterRange{3773, 3773}, TSCharacterRange{3776, 3780}, TSCharacterRange{3782, 3782}, TSCharacterRange{3804, 3807}, TSCharacterRange{3840, 3840}, TSCharacterRange{3904, 3911}, TSCharacterRange{3913, 3948}, TSCharacterRange{3976, 3980}, TSCharacterRange{4096, 4138}, TSCharacterRange{4159, 4159}, TSCharacterRange{4176, 4181}, TSCharacterRange{4186, 4189}, TSCharacterRange{4193, 4193}, TSCharacterRange{4197, 4198}, TSCharacterRange{4206, 4208}, TSCharacterRange{4213, 4225}, TSCharacterRange{4238, 4238}, TSCharacterRange{4256, 4293}, TSCharacterRange{4295, 4295}, TSCharacterRange{4301, 4301}, TSCharacterRange{4304, 4346}, TSCharacterRange{4348, 4680}, TSCharacterRange{4682, 4685}, TSCharacterRange{4688, 4694}, TSCharacterRange{4696, 4696}, TSCharacterRange{4698, 4701}, TSCharacterRange{4704, 4744}, TSCharacterRange{4746, 4749}, TSCharacterRange{4752, 4784}, TSCharacterRange{4786, 4789}, TSCharacterRange{4792, 4798}, TSCharacterRange{4800, 4800}, TSCharacterRange{4802, 4805}, TSCharacterRange{4808, 4822}, TSCharacterRange{4824, 4880}, TSCharacterRange{4882, 4885}, TSCharacterRange{4888, 4954}, TSCharacterRange{4992, 5007}, TSCharacterRange{5024, 5109}, TSCharacterRange{5112, 5117}, TSCharacterRange{5121, 5740}, TSCharacterRange{5743, 5759}, TSCharacterRange{5761, 5786}, TSCharacterRange{5792, 5866}, TSCharacterRange{5873, 5880}, TSCharacterRange{5888, 5905}, TSCharacterRange{5919, 5937}, TSCharacterRange{5952, 5969}, TSCharacterRange{5984, 5996}, TSCharacterRange{5998, 6000}, TSCharacterRange{6016, 6067}, TSCharacterRange{6103, 6103}, TSCharacterRange{6108, 6108}, TSCharacterRange{6176, 6264}, TSCharacterRange{6272, 6276}, TSCharacterRange{6279, 6312}, TSCharacterRange{6314, 6314}, TSCharacterRange{6320, 6389}, TSCharacterRange{6400, 6430}, TSCharacterRange{6480, 6509}, TSCharacterRange{6512, 6516}, TSCharacterRange{6528, 6571}, TSCharacterRange{6576, 6601}, TSCharacterRange{6656, 6678}, TSCharacterRange{6688, 6740}, TSCharacterRange{6823, 6823}, TSCharacterRange{6917, 6963}, TSCharacterRange{6981, 6988}, TSCharacterRange{7043, 7072}, TSCharacterRange{7086, 7087}, TSCharacterRange{7098, 7141}, TSCharacterRange{7168, 7203}, TSCharacterRange{7245, 7247}, TSCharacterRange{7258, 7293}, TSCharacterRange{7296, 7306}, TSCharacterRange{7312, 7354}, TSCharacterRange{7357, 7359}, TSCharacterRange{7401, 7404}, TSCharacterRange{7406, 7411}, TSCharacterRange{7413, 7414}, TSCharacterRange{7418, 7418}, TSCharacterRange{7424, 7615}, TSCharacterRange{7680, 7957}, TSCharacterRange{7960, 7965}, TSCharacterRange{7968, 8005}, TSCharacterRange{8008, 8013}, TSCharacterRange{8016, 8023}, TSCharacterRange{8025, 8025}, TSCharacterRange{8027, 8027}, TSCharacterRange{8029, 8029}, TSCharacterRange{8031, 8061}, TSCharacterRange{8064, 8116}, TSCharacterRange{8118, 8124}, TSCharacterRange{8126, 8126}, TSCharacterRange{8130, 8132}, TSCharacterRange{8134, 8140}, TSCharacterRange{8144, 8147}, TSCharacterRange{8150, 8155}, TSCharacterRange{8160, 8172}, TSCharacterRange{8178, 8180}, TSCharacterRange{8182, 8188}, TSCharacterRange{8305, 8305}, TSCharacterRange{8319, 8319}, TSCharacterRange{8336, 8348}, TSCharacterRange{8450, 8450}, TSCharacterRange{8455, 8455}, TSCharacterRange{8458, 8467}, TSCharacterRange{8469, 8469}, TSCharacterRange{8473, 8477}, TSCharacterRange{8484, 8484}, TSCharacterRange{8486, 8486}, TSCharacterRange{8488, 8488}, TSCharacterRange{8490, 8493}, TSCharacterRange{8495, 8505}, TSCharacterRange{8508, 8511}, TSCharacterRange{8517, 8521}, TSCharacterRange{8526, 8526}, TSCharacterRange{8579, 8580}, TSCharacterRange{11264, 11492}, TSCharacterRange{11499, 11502}, TSCharacterRange{11506, 11507}, TSCharacterRange{11520, 11557}, TSCharacterRange{11559, 11559}, TSCharacterRange{11565, 11565}, TSCharacterRange{11568, 11623}, TSCharacterRange{11631, 11631}, TSCharacterRange{11648, 11670}, TSCharacterRange{11680, 11686}, TSCharacterRange{11688, 11694}, TSCharacterRange{11696, 11702}, TSCharacterRange{11704, 11710}, TSCharacterRange{11712, 11718}, TSCharacterRange{11720, 11726}, TSCharacterRange{11728, 11734}, TSCharacterRange{11736, 11742}, TSCharacterRange{11823, 11823}, TSCharacterRange{12293, 12294}, TSCharacterRange{12337, 12341}, TSCharacterRange{12347, 12348}, TSCharacterRange{12353, 12438}, TSCharacterRange{12445, 12447}, TSCharacterRange{12449, 12538}, TSCharacterRange{12540, 12543}, TSCharacterRange{12549, 12591}, TSCharacterRange{12593, 12686}, TSCharacterRange{12704, 12735}, TSCharacterRange{12784, 12799}, TSCharacterRange{13312, 19903}, TSCharacterRange{19968, 42124}, TSCharacterRange{42192, 42237}, TSCharacterRange{42240, 42508}, TSCharacterRange{42512, 42527}, TSCharacterRange{42538, 42539}, TSCharacterRange{42560, 42606}, TSCharacterRange{42623, 42653}, TSCharacterRange{42656, 42725}, TSCharacterRange{42775, 42783}, TSCharacterRange{42786, 42888}, TSCharacterRange{42891, 42957}, TSCharacterRange{42960, 42961}, TSCharacterRange{42963, 42963}, TSCharacterRange{42965, 42972}, TSCharacterRange{42994, 43009}, TSCharacterRange{43011, 43013}, TSCharacterRange{43015, 43018}, TSCharacterRange{43020, 43042}, TSCharacterRange{43072, 43123}, TSCharacterRange{43138, 43187}, TSCharacterRange{43250, 43255}, TSCharacterRange{43259, 43259}, TSCharacterRange{43261, 43262}, TSCharacterRange{43274, 43301}, TSCharacterRange{43312, 43334}, TSCharacterRange{43360, 43388}, TSCharacterRange{43396, 43442}, TSCharacterRange{43471, 43471}, TSCharacterRange{43488, 43492}, TSCharacterRange{43494, 43503}, TSCharacterRange{43514, 43518}, TSCharacterRange{43520, 43560}, TSCharacterRange{43584, 43586}, TSCharacterRange{43588, 43595}, TSCharacterRange{43616, 43638}, TSCharacterRange{43642, 43642}, TSCharacterRange{43646, 43695}, TSCharacterRange{43697, 43697}, TSCharacterRange{43701, 43702}, TSCharacterRange{43705, 43709}, TSCharacterRange{43712, 43712}, TSCharacterRange{43714, 43714}, TSCharacterRange{43739, 43741}, TSCharacterRange{43744, 43754}, TSCharacterRange{43762, 43764}, TSCharacterRange{43777, 43782}, TSCharacterRange{43785, 43790}, TSCharacterRange{43793, 43798}, TSCharacterRange{43808, 43814}, TSCharacterRange{43816, 43822}, TSCharacterRange{43824, 43866}, TSCharacterRange{43868, 43881}, TSCharacterRange{43888, 44002}, TSCharacterRange{44032, 55203}, TSCharacterRange{55216, 55238}, TSCharacterRange{55243, 55291}, TSCharacterRange{63744, 64109}, TSCharacterRange{64112, 64217}, TSCharacterRange{64256, 64262}, TSCharacterRange{64275, 64279}, TSCharacterRange{64285, 64285}, TSCharacterRange{64287, 64296}, TSCharacterRange{64298, 64310}, TSCharacterRange{64312, 64316}, TSCharacterRange{64318, 64318}, TSCharacterRange{64320, 64321}, TSCharacterRange{64323, 64324}, TSCharacterRange{64326, 64433}, TSCharacterRange{64467, 64829}, TSCharacterRange{64848, 64911}, TSCharacterRange{64914, 64967}, TSCharacterRange{65008, 65019}, TSCharacterRange{65136, 65140}, TSCharacterRange{65142, 65276}, TSCharacterRange{65313, 65338}, TSCharacterRange{65345, 65370}, TSCharacterRange{65382, 65470}, TSCharacterRange{65474, 65479}, TSCharacterRange{65482, 65487}, TSCharacterRange{65490, 65495}, TSCharacterRange{65498, 65500}, TSCharacterRange{65536, 65547}, TSCharacterRange{65549, 65574}, TSCharacterRange{65576, 65594}, TSCharacterRange{65596, 65597}, TSCharacterRange{65599, 65613}, TSCharacterRange{65616, 65629}, TSCharacterRange{65664, 65786}, TSCharacterRange{66176, 66204}, TSCharacterRange{66208, 66256}, TSCharacterRange{66304, 66335}, TSCharacterRange{66349, 66368}, TSCharacterRange{66370, 66377}, TSCharacterRange{66384, 66421}, TSCharacterRange{66432, 66461}, TSCharacterRange{66464, 66499}, TSCharacterRange{66504, 66511}, TSCharacterRange{66560, 66717}, TSCharacterRange{66736, 66771}, TSCharacterRange{66776, 66811}, TSCharacterRange{66816, 66855}, TSCharacterRange{66864, 66915}, TSCharacterRange{66928, 66938}, TSCharacterRange{66940, 66954}, TSCharacterRange{66956, 66962}, TSCharacterRange{66964, 66965}, TSCharacterRange{66967, 66977}, TSCharacterRange{66979, 66993}, TSCharacterRange{66995, 67001}, TSCharacterRange{67003, 67004}, TSCharacterRange{67008, 67059}, TSCharacterRange{67072, 67382}, TSCharacterRange{67392, 67413}, TSCharacterRange{67424, 67431}, TSCharacterRange{67456, 67461}, TSCharacterRange{67463, 67504}, TSCharacterRange{67506, 67514}, TSCharacterRange{67584, 67589}, TSCharacterRange{67592, 67592}, TSCharacterRange{67594, 67637}, TSCharacterRange{67639, 67640}, TSCharacterRange{67644, 67644}, TSCharacterRange{67647, 67669}, TSCharacterRange{67680, 67702}, TSCharacterRange{67712, 67742}, TSCharacterRange{67808, 67826}, TSCharacterRange{67828, 67829}, TSCharacterRange{67840, 67861}, TSCharacterRange{67872, 67897}, TSCharacterRange{67968, 68023}, TSCharacterRange{68030, 68031}, TSCharacterRange{68096, 68096}, TSCharacterRange{68112, 68115}, TSCharacterRange{68117, 68119}, TSCharacterRange{68121, 68149}, TSCharacterRange{68192, 68220}, TSCharacterRange{68224, 68252}, TSCharacterRange{68288, 68295}, TSCharacterRange{68297, 68324}, TSCharacterRange{68352, 68405}, TSCharacterRange{68416, 68437}, TSCharacterRange{68448, 68466}, TSCharacterRange{68480, 68497}, TSCharacterRange{68608, 68680}, TSCharacterRange{68736, 68786}, TSCharacterRange{68800, 68850}, TSCharacterRange{68864, 68899}, TSCharacterRange{68938, 68965}, TSCharacterRange{68975, 68997}, TSCharacterRange{69248, 69289}, TSCharacterRange{69296, 69297}, TSCharacterRange{69314, 69316}, TSCharacterRange{69376, 69404}, TSCharacterRange{69415, 69415}, TSCharacterRange{69424, 69445}, TSCharacterRange{69488, 69505}, TSCharacterRange{69552, 69572}, TSCharacterRange{69600, 69622}, TSCharacterRange{69635, 69687}, TSCharacterRange{69745, 69746}, TSCharacterRange{69749, 69749}, TSCharacterRange{69763, 69807}, TSCharacterRange{69840, 69864}, TSCharacterRange{69891, 69926}, TSCharacterRange{69956, 69956}, TSCharacterRange{69959, 69959}, TSCharacterRange{69968, 70002}, TSCharacterRange{70006, 70006}, TSCharacterRange{70019, 70066}, TSCharacterRange{70081, 70084}, TSCharacterRange{70106, 70106}, TSCharacterRange{70108, 70108}, TSCharacterRange{70144, 70161}, TSCharacterRange{70163, 70187}, TSCharacterRange{70207, 70208}, TSCharacterRange{70272, 70278}, TSCharacterRange{70280, 70280}, TSCharacterRange{70282, 70285}, TSCharacterRange{70287, 70301}, TSCharacterRange{70303, 70312}, TSCharacterRange{70320, 70366}, TSCharacterRange{70405, 70412}, TSCharacterRange{70415, 70416}, TSCharacterRange{70419, 70440}, TSCharacterRange{70442, 70448}, TSCharacterRange{70450, 70451}, TSCharacterRange{70453, 70457}, TSCharacterRange{70461, 70461}, TSCharacterRange{70480, 70480}, TSCharacterRange{70493, 70497}, TSCharacterRange{70528, 70537}, TSCharacterRange{70539, 70539}, TSCharacterRange{70542, 70542}, TSCharacterRange{70544, 70581}, TSCharacterRange{70583, 70583}, TSCharacterRange{70609, 70609}, TSCharacterRange{70611, 70611}, TSCharacterRange{70656, 70708}, TSCharacterRange{70727, 70730}, TSCharacterRange{70751, 70753}, TSCharacterRange{70784, 70831}, TSCharacterRange{70852, 70853}, TSCharacterRange{70855, 70855}, TSCharacterRange{71040, 71086}, TSCharacterRange{71128, 71131}, TSCharacterRange{71168, 71215}, TSCharacterRange{71236, 71236}, TSCharacterRange{71296, 71338}, TSCharacterRange{71352, 71352}, TSCharacterRange{71424, 71450}, TSCharacterRange{71488, 71494}, TSCharacterRange{71680, 71723}, TSCharacterRange{71840, 71903}, TSCharacterRange{71935, 71942}, TSCharacterRange{71945, 71945}, TSCharacterRange{71948, 71955}, TSCharacterRange{71957, 71958}, TSCharacterRange{71960, 71983}, TSCharacterRange{71999, 71999}, TSCharacterRange{72001, 72001}, TSCharacterRange{72096, 72103}, TSCharacterRange{72106, 72144}, TSCharacterRange{72161, 72161}, TSCharacterRange{72163, 72163}, TSCharacterRange{72192, 72192}, TSCharacterRange{72203, 72242}, TSCharacterRange{72250, 72250}, TSCharacterRange{72272, 72272}, TSCharacterRange{72284, 72329}, TSCharacterRange{72349, 72349}, TSCharacterRange{72368, 72440}, TSCharacterRange{72640, 72672}, TSCharacterRange{72704, 72712}, TSCharacterRange{72714, 72750}, TSCharacterRange{72768, 72768}, TSCharacterRange{72818, 72847}, TSCharacterRange{72960, 72966}, TSCharacterRange{72968, 72969}, TSCharacterRange{72971, 73008}, TSCharacterRange{73030, 73030}, TSCharacterRange{73056, 73061}, TSCharacterRange{73063, 73064}, TSCharacterRange{73066, 73097}, TSCharacterRange{73112, 73112}, TSCharacterRange{73440, 73458}, TSCharacterRange{73474, 73474}, TSCharacterRange{73476, 73488}, TSCharacterRange{73490, 73523}, TSCharacterRange{73648, 73648}, TSCharacterRange{73728, 74649}, TSCharacterRange{74880, 75075}, TSCharacterRange{77712, 77808}, TSCharacterRange{77824, 78895}, TSCharacterRange{78913, 78918}, TSCharacterRange{78944, 82938}, TSCharacterRange{82944, 83526}, TSCharacterRange{90368, 90397}, TSCharacterRange{92160, 92728}, TSCharacterRange{92736, 92766}, TSCharacterRange{92784, 92862}, TSCharacterRange{92880, 92909}, TSCharacterRange{92928, 92975}, TSCharacterRange{92992, 92995}, TSCharacterRange{93027, 93047}, TSCharacterRange{93053, 93071}, TSCharacterRange{93504, 93548}, TSCharacterRange{93760, 93823}, TSCharacterRange{93952, 94026}, TSCharacterRange{94032, 94032}, TSCharacterRange{94099, 94111}, TSCharacterRange{94176, 94177}, TSCharacterRange{94179, 94179}, TSCharacterRange{94208, 100343}, TSCharacterRange{100352, 101589}, TSCharacterRange{101631, 101640}, TSCharacterRange{110576, 110579}, TSCharacterRange{110581, 110587}, TSCharacterRange{110589, 110590}, TSCharacterRange{110592, 110882}, TSCharacterRange{110898, 110898}, TSCharacterRange{110928, 110930}, TSCharacterRange{110933, 110933}, TSCharacterRange{110948, 110951}, TSCharacterRange{110960, 111355}, TSCharacterRange{113664, 113770}, TSCharacterRange{113776, 113788}, TSCharacterRange{113792, 113800}, TSCharacterRange{113808, 113817}, TSCharacterRange{119808, 119892}, TSCharacterRange{119894, 119964}, TSCharacterRange{119966, 119967}, TSCharacterRange{119970, 119970}, TSCharacterRange{119973, 119974}, TSCharacterRange{119977, 119980}, TSCharacterRange{119982, 119993}, TSCharacterRange{119995, 119995}, TSCharacterRange{119997, 120003}, TSCharacterRange{120005, 120069}, TSCharacterRange{120071, 120074}, TSCharacterRange{120077, 120084}, TSCharacterRange{120086, 120092}, TSCharacterRange{120094, 120121}, TSCharacterRange{120123, 120126}, TSCharacterRange{120128, 120132}, TSCharacterRange{120134, 120134}, TSCharacterRange{120138, 120144}, TSCharacterRange{120146, 120485}, TSCharacterRange{120488, 120512}, TSCharacterRange{120514, 120538}, TSCharacterRange{120540, 120570}, TSCharacterRange{120572, 120596}, TSCharacterRange{120598, 120628}, TSCharacterRange{120630, 120654}, TSCharacterRange{120656, 120686}, TSCharacterRange{120688, 120712}, TSCharacterRange{120714, 120744}, TSCharacterRange{120746, 120770}, TSCharacterRange{120772, 120779}, TSCharacterRange{122624, 122654}, TSCharacterRange{122661, 122666}, TSCharacterRange{122928, 122989}, TSCharacterRange{123136, 123180}, TSCharacterRange{123191, 123197}, TSCharacterRange{123214, 123214}, TSCharacterRange{123536, 123565}, TSCharacterRange{123584, 123627}, TSCharacterRange{124112, 124139}, TSCharacterRange{124368, 124397}, TSCharacterRange{124400, 124400}, TSCharacterRange{124896, 124902}, TSCharacterRange{124904, 124907}, TSCharacterRange{124909, 124910}, TSCharacterRange{124912, 124926}, TSCharacterRange{124928, 125124}, TSCharacterRange{125184, 125251}, TSCharacterRange{125259, 125259}, TSCharacterRange{126464, 126467}, TSCharacterRange{126469, 126495}, TSCharacterRange{126497, 126498}, TSCharacterRange{126500, 126500}, TSCharacterRange{126503, 126503}, TSCharacterRange{126505, 126514}, TSCharacterRange{126516, 126519}, TSCharacterRange{126521, 126521}, TSCharacterRange{126523, 126523}, TSCharacterRange{126530, 126530}, TSCharacterRange{126535, 126535}, TSCharacterRange{126537, 126537}, TSCharacterRange{126539, 126539}, TSCharacterRange{126541, 126543}, TSCharacterRange{126545, 126546}, TSCharacterRange{126548, 126548}, TSCharacterRange{126551, 126551}, TSCharacterRange{126553, 126553}, TSCharacterRange{126555, 126555}, TSCharacterRange{126557, 126557}, TSCharacterRange{126559, 126559}, TSCharacterRange{126561, 126562}, TSCharacterRange{126564, 126564}, TSCharacterRange{126567, 126570}, TSCharacterRange{126572, 126578}, TSCharacterRange{126580, 126583}, TSCharacterRange{126585, 126588}, TSCharacterRange{126590, 126590}, TSCharacterRange{126592, 126601}, TSCharacterRange{126603, 126619}, TSCharacterRange{126625, 126627}, TSCharacterRange{126629, 126633}, TSCharacterRange{126635, 126651}, TSCharacterRange{131072, 173791}, TSCharacterRange{173824, 177977}, TSCharacterRange{177984, 178205}, TSCharacterRange{178208, 183969}, TSCharacterRange{183984, 191456}, TSCharacterRange{191472, 192093}, TSCharacterRange{194560, 195101}, TSCharacterRange{196608, 201546}, TSCharacterRange{201552, 205743}}
var ts_lex_map_25 [28]int16 = [28]int16{13, 1, 117, 41, 120, 37, 10, 7, 34, 7, 39, 7, 47, 7, 92, 7, 98, 7, 102, 7, 110, 7, 114, 7, 116, 7, 118, 7}
var ts_lex_map_26 [28]int16 = [28]int16{13, 2, 117, 42, 120, 38, 10, 8, 34, 8, 39, 8, 47, 8, 92, 8, 98, 8, 102, 8, 110, 8, 114, 8, 116, 8, 118, 8}
var ts_lex_map_27 [30]int16 = [30]int16{34, 7, 39, 8, 46, 76, 47, 9, 48, 74, 73, 24, 78, 14, 91, 70, 93, 71, 102, 15, 110, 30, 116, 26, 123, 46, 43, 12, 45, 12}
var sym_identifier_character_set_1 [679]TSCharacterRange = [679]TSCharacterRange{TSCharacterRange{36, 36}, TSCharacterRange{65, 90}, TSCharacterRange{95, 95}, TSCharacterRange{97, 122}, TSCharacterRange{170, 170}, TSCharacterRange{181, 181}, TSCharacterRange{186, 186}, TSCharacterRange{192, 214}, TSCharacterRange{216, 246}, TSCharacterRange{248, 705}, TSCharacterRange{710, 721}, TSCharacterRange{736, 740}, TSCharacterRange{748, 748}, TSCharacterRange{750, 750}, TSCharacterRange{880, 884}, TSCharacterRange{886, 887}, TSCharacterRange{890, 893}, TSCharacterRange{895, 895}, TSCharacterRange{902, 902}, TSCharacterRange{904, 906}, TSCharacterRange{908, 908}, TSCharacterRange{910, 929}, TSCharacterRange{931, 1013}, TSCharacterRange{1015, 1153}, TSCharacterRange{1162, 1327}, TSCharacterRange{1329, 1366}, TSCharacterRange{1369, 1369}, TSCharacterRange{1376, 1416}, TSCharacterRange{1488, 1514}, TSCharacterRange{1519, 1522}, TSCharacterRange{1568, 1610}, TSCharacterRange{1646, 1647}, TSCharacterRange{1649, 1747}, TSCharacterRange{1749, 1749}, TSCharacterRange{1765, 1766}, TSCharacterRange{1774, 1775}, TSCharacterRange{1786, 1788}, TSCharacterRange{1791, 1791}, TSCharacterRange{1808, 1808}, TSCharacterRange{1810, 1839}, TSCharacterRange{1869, 1957}, TSCharacterRange{1969, 1969}, TSCharacterRange{1994, 2026}, TSCharacterRange{2036, 2037}, TSCharacterRange{2042, 2042}, TSCharacterRange{2048, 2069}, TSCharacterRange{2074, 2074}, TSCharacterRange{2084, 2084}, TSCharacterRange{2088, 2088}, TSCharacterRange{2112, 2136}, TSCharacterRange{2144, 2154}, TSCharacterRange{2160, 2183}, TSCharacterRange{2185, 2190}, TSCharacterRange{2208, 2249}, TSCharacterRange{2308, 2361}, TSCharacterRange{2365, 2365}, TSCharacterRange{2384, 2384}, TSCharacterRange{2392, 2401}, TSCharacterRange{2417, 2432}, TSCharacterRange{2437, 2444}, TSCharacterRange{2447, 2448}, TSCharacterRange{2451, 2472}, TSCharacterRange{2474, 2480}, TSCharacterRange{2482, 2482}, TSCharacterRange{2486, 2489}, TSCharacterRange{2493, 2493}, TSCharacterRange{2510, 2510}, TSCharacterRange{2524, 2525}, TSCharacterRange{2527, 2529}, TSCharacterRange{2544, 2545}, TSCharacterRange{2556, 2556}, TSCharacterRange{2565, 2570}, TSCharacterRange{2575, 2576}, TSCharacterRange{2579, 2600}, TSCharacterRange{2602, 2608}, TSCharacterRange{2610, 2611}, TSCharacterRange{2613, 2614}, TSCharacterRange{2616, 2617}, TSCharacterRange{2649, 2652}, TSCharacterRange{2654, 2654}, TSCharacterRange{2674, 2676}, TSCharacterRange{2693, 2701}, TSCharacterRange{2703, 2705}, TSCharacterRange{2707, 2728}, TSCharacterRange{2730, 2736}, TSCharacterRange{2738, 2739}, TSCharacterRange{2741, 2745}, TSCharacterRange{2749, 2749}, TSCharacterRange{2768, 2768}, TSCharacterRange{2784, 2785}, TSCharacterRange{2809, 2809}, TSCharacterRange{2821, 2828}, TSCharacterRange{2831, 2832}, TSCharacterRange{2835, 2856}, TSCharacterRange{2858, 2864}, TSCharacterRange{2866, 2867}, TSCharacterRange{2869, 2873}, TSCharacterRange{2877, 2877}, TSCharacterRange{2908, 2909}, TSCharacterRange{2911, 2913}, TSCharacterRange{2929, 2929}, TSCharacterRange{2947, 2947}, TSCharacterRange{2949, 2954}, TSCharacterRange{2958, 2960}, TSCharacterRange{2962, 2965}, TSCharacterRange{2969, 2970}, TSCharacterRange{2972, 2972}, TSCharacterRange{2974, 2975}, TSCharacterRange{2979, 2980}, TSCharacterRange{2984, 2986}, TSCharacterRange{2990, 3001}, TSCharacterRange{3024, 3024}, TSCharacterRange{3077, 3084}, TSCharacterRange{3086, 3088}, TSCharacterRange{3090, 3112}, TSCharacterRange{3114, 3129}, TSCharacterRange{3133, 3133}, TSCharacterRange{3160, 3162}, TSCharacterRange{3165, 3165}, TSCharacterRange{3168, 3169}, TSCharacterRange{3200, 3200}, TSCharacterRange{3205, 3212}, TSCharacterRange{3214, 3216}, TSCharacterRange{3218, 3240}, TSCharacterRange{3242, 3251}, TSCharacterRange{3253, 3257}, TSCharacterRange{3261, 3261}, TSCharacterRange{3293, 3294}, TSCharacterRange{3296, 3297}, TSCharacterRange{3313, 3314}, TSCharacterRange{3332, 3340}, TSCharacterRange{3342, 3344}, TSCharacterRange{3346, 3386}, TSCharacterRange{3389, 3389}, TSCharacterRange{3406, 3406}, TSCharacterRange{3412, 3414}, TSCharacterRange{3423, 3425}, TSCharacterRange{3450, 3455}, TSCharacterRange{3461, 3478}, TSCharacterRange{3482, 3505}, TSCharacterRange{3507, 3515}, TSCharacterRange{3517, 3517}, TSCharacterRange{3520, 3526}, TSCharacterRange{3585, 3632}, TSCharacterRange{3634, 3635}, TSCharacterRange{3648, 3654}, TSCharacterRange{3713, 3714}, TSCharacterRange{3716, 3716}, TSCharacterRange{3718, 3722}, TSCharacterRange{3724, 3747}, TSCharacterRange{3749, 3749}, TSCharacterRange{3751, 3760}, TSCharacterRange{3762, 3763}, TSCharacterRange{3773, 3773}, TSCharacterRange{3776, 3780}, TSCharacterRange{3782, 3782}, TSCharacterRange{3804, 3807}, TSCharacterRange{3840, 3840}, TSCharacterRange{3904, 3911}, TSCharacterRange{3913, 3948}, TSCharacterRange{3976, 3980}, TSCharacterRange{4096, 4138}, TSCharacterRange{4159, 4159}, TSCharacterRange{4176, 4181}, TSCharacterRange{4186, 4189}, TSCharacterRange{4193, 4193}, TSCharacterRange{4197, 4198}, TSCharacterRange{4206, 4208}, TSCharacterRange{4213, 4225}, TSCharacterRange{4238, 4238}, TSCharacterRange{4256, 4293}, TSCharacterRange{4295, 4295}, TSCharacterRange{4301, 4301}, TSCharacterRange{4304, 4346}, TSCharacterRange{4348, 4680}, TSCharacterRange{4682, 4685}, TSCharacterRange{4688, 4694}, TSCharacterRange{4696, 4696}, TSCharacterRange{4698, 4701}, TSCharacterRange{4704, 4744}, TSCharacterRange{4746, 4749}, TSCharacterRange{4752, 4784}, TSCharacterRange{4786, 4789}, TSCharacterRange{4792, 4798}, TSCharacterRange{4800, 4800}, TSCharacterRange{4802, 4805}, TSCharacterRange{4808, 4822}, TSCharacterRange{4824, 4880}, TSCharacterRange{4882, 4885}, TSCharacterRange{4888, 4954}, TSCharacterRange{4992, 5007}, TSCharacterRange{5024, 5109}, TSCharacterRange{5112, 5117}, TSCharacterRange{5121, 5740}, TSCharacterRange{5743, 5759}, TSCharacterRange{5761, 5786}, TSCharacterRange{5792, 5866}, TSCharacterRange{5873, 5880}, TSCharacterRange{5888, 5905}, TSCharacterRange{5919, 5937}, TSCharacterRange{5952, 5969}, TSCharacterRange{5984, 5996}, TSCharacterRange{5998, 6000}, TSCharacterRange{6016, 6067}, TSCharacterRange{6103, 6103}, TSCharacterRange{6108, 6108}, TSCharacterRange{6176, 6264}, TSCharacterRange{6272, 6276}, TSCharacterRange{6279, 6312}, TSCharacterRange{6314, 6314}, TSCharacterRange{6320, 6389}, TSCharacterRange{6400, 6430}, TSCharacterRange{6480, 6509}, TSCharacterRange{6512, 6516}, TSCharacterRange{6528, 6571}, TSCharacterRange{6576, 6601}, TSCharacterRange{6656, 6678}, TSCharacterRange{6688, 6740}, TSCharacterRange{6823, 6823}, TSCharacterRange{6917, 6963}, TSCharacterRange{6981, 6988}, TSCharacterRange{7043, 7072}, TSCharacterRange{7086, 7087}, TSCharacterRange{7098, 7141}, TSCharacterRange{7168, 7203}, TSCharacterRange{7245, 7247}, TSCharacterRange{7258, 7293}, TSCharacterRange{7296, 7306}, TSCharacterRange{7312, 7354}, TSCharacterRange{7357, 7359}, TSCharacterRange{7401, 7404}, TSCharacterRange{7406, 7411}, TSCharacterRange{7413, 7414}, TSCharacterRange{7418, 7418}, TSCharacterRange{7424, 7615}, TSCharacterRange{7680, 7957}, TSCharacterRange{7960, 7965}, TSCharacterRange{7968, 8005}, TSCharacterRange{8008, 8013}, TSCharacterRange{8016, 8023}, TSCharacterRange{8025, 8025}, TSCharacterRange{8027, 8027}, TSCharacterRange{8029, 8029}, TSCharacterRange{8031, 8061}, TSCharacterRange{8064, 8116}, TSCharacterRange{8118, 8124}, TSCharacterRange{8126, 8126}, TSCharacterRange{8130, 8132}, TSCharacterRange{8134, 8140}, TSCharacterRange{8144, 8147}, TSCharacterRange{8150, 8155}, TSCharacterRange{8160, 8172}, TSCharacterRange{8178, 8180}, TSCharacterRange{8182, 8188}, TSCharacterRange{8305, 8305}, TSCharacterRange{8319, 8319}, TSCharacterRange{8336, 8348}, TSCharacterRange{8450, 8450}, TSCharacterRange{8455, 8455}, TSCharacterRange{8458, 8467}, TSCharacterRange{8469, 8469}, TSCharacterRange{8473, 8477}, TSCharacterRange{8484, 8484}, TSCharacterRange{8486, 8486}, TSCharacterRange{8488, 8488}, TSCharacterRange{8490, 8493}, TSCharacterRange{8495, 8505}, TSCharacterRange{8508, 8511}, TSCharacterRange{8517, 8521}, TSCharacterRange{8526, 8526}, TSCharacterRange{8579, 8580}, TSCharacterRange{11264, 11492}, TSCharacterRange{11499, 11502}, TSCharacterRange{11506, 11507}, TSCharacterRange{11520, 11557}, TSCharacterRange{11559, 11559}, TSCharacterRange{11565, 11565}, TSCharacterRange{11568, 11623}, TSCharacterRange{11631, 11631}, TSCharacterRange{11648, 11670}, TSCharacterRange{11680, 11686}, TSCharacterRange{11688, 11694}, TSCharacterRange{11696, 11702}, TSCharacterRange{11704, 11710}, TSCharacterRange{11712, 11718}, TSCharacterRange{11720, 11726}, TSCharacterRange{11728, 11734}, TSCharacterRange{11736, 11742}, TSCharacterRange{11823, 11823}, TSCharacterRange{12293, 12294}, TSCharacterRange{12337, 12341}, TSCharacterRange{12347, 12348}, TSCharacterRange{12353, 12438}, TSCharacterRange{12445, 12447}, TSCharacterRange{12449, 12538}, TSCharacterRange{12540, 12543}, TSCharacterRange{12549, 12591}, TSCharacterRange{12593, 12686}, TSCharacterRange{12704, 12735}, TSCharacterRange{12784, 12799}, TSCharacterRange{13312, 19903}, TSCharacterRange{19968, 42124}, TSCharacterRange{42192, 42237}, TSCharacterRange{42240, 42508}, TSCharacterRange{42512, 42527}, TSCharacterRange{42538, 42539}, TSCharacterRange{42560, 42606}, TSCharacterRange{42623, 42653}, TSCharacterRange{42656, 42725}, TSCharacterRange{42775, 42783}, TSCharacterRange{42786, 42888}, TSCharacterRange{42891, 42957}, TSCharacterRange{42960, 42961}, TSCharacterRange{42963, 42963}, TSCharacterRange{42965, 42972}, TSCharacterRange{42994, 43009}, TSCharacterRange{43011, 43013}, TSCharacterRange{43015, 43018}, TSCharacterRange{43020, 43042}, TSCharacterRange{43072, 43123}, TSCharacterRange{43138, 43187}, TSCharacterRange{43250, 43255}, TSCharacterRange{43259, 43259}, TSCharacterRange{43261, 43262}, TSCharacterRange{43274, 43301}, TSCharacterRange{43312, 43334}, TSCharacterRange{43360, 43388}, TSCharacterRange{43396, 43442}, TSCharacterRange{43471, 43471}, TSCharacterRange{43488, 43492}, TSCharacterRange{43494, 43503}, TSCharacterRange{43514, 43518}, TSCharacterRange{43520, 43560}, TSCharacterRange{43584, 43586}, TSCharacterRange{43588, 43595}, TSCharacterRange{43616, 43638}, TSCharacterRange{43642, 43642}, TSCharacterRange{43646, 43695}, TSCharacterRange{43697, 43697}, TSCharacterRange{43701, 43702}, TSCharacterRange{43705, 43709}, TSCharacterRange{43712, 43712}, TSCharacterRange{43714, 43714}, TSCharacterRange{43739, 43741}, TSCharacterRange{43744, 43754}, TSCharacterRange{43762, 43764}, TSCharacterRange{43777, 43782}, TSCharacterRange{43785, 43790}, TSCharacterRange{43793, 43798}, TSCharacterRange{43808, 43814}, TSCharacterRange{43816, 43822}, TSCharacterRange{43824, 43866}, TSCharacterRange{43868, 43881}, TSCharacterRange{43888, 44002}, TSCharacterRange{44032, 55203}, TSCharacterRange{55216, 55238}, TSCharacterRange{55243, 55291}, TSCharacterRange{63744, 64109}, TSCharacterRange{64112, 64217}, TSCharacterRange{64256, 64262}, TSCharacterRange{64275, 64279}, TSCharacterRange{64285, 64285}, TSCharacterRange{64287, 64296}, TSCharacterRange{64298, 64310}, TSCharacterRange{64312, 64316}, TSCharacterRange{64318, 64318}, TSCharacterRange{64320, 64321}, TSCharacterRange{64323, 64324}, TSCharacterRange{64326, 64433}, TSCharacterRange{64467, 64829}, TSCharacterRange{64848, 64911}, TSCharacterRange{64914, 64967}, TSCharacterRange{65008, 65019}, TSCharacterRange{65136, 65140}, TSCharacterRange{65142, 65276}, TSCharacterRange{65313, 65338}, TSCharacterRange{65345, 65370}, TSCharacterRange{65382, 65470}, TSCharacterRange{65474, 65479}, TSCharacterRange{65482, 65487}, TSCharacterRange{65490, 65495}, TSCharacterRange{65498, 65500}, TSCharacterRange{65536, 65547}, TSCharacterRange{65549, 65574}, TSCharacterRange{65576, 65594}, TSCharacterRange{65596, 65597}, TSCharacterRange{65599, 65613}, TSCharacterRange{65616, 65629}, TSCharacterRange{65664, 65786}, TSCharacterRange{66176, 66204}, TSCharacterRange{66208, 66256}, TSCharacterRange{66304, 66335}, TSCharacterRange{66349, 66368}, TSCharacterRange{66370, 66377}, TSCharacterRange{66384, 66421}, TSCharacterRange{66432, 66461}, TSCharacterRange{66464, 66499}, TSCharacterRange{66504, 66511}, TSCharacterRange{66560, 66717}, TSCharacterRange{66736, 66771}, TSCharacterRange{66776, 66811}, TSCharacterRange{66816, 66855}, TSCharacterRange{66864, 66915}, TSCharacterRange{66928, 66938}, TSCharacterRange{66940, 66954}, TSCharacterRange{66956, 66962}, TSCharacterRange{66964, 66965}, TSCharacterRange{66967, 66977}, TSCharacterRange{66979, 66993}, TSCharacterRange{66995, 67001}, TSCharacterRange{67003, 67004}, TSCharacterRange{67008, 67059}, TSCharacterRange{67072, 67382}, TSCharacterRange{67392, 67413}, TSCharacterRange{67424, 67431}, TSCharacterRange{67456, 67461}, TSCharacterRange{67463, 67504}, TSCharacterRange{67506, 67514}, TSCharacterRange{67584, 67589}, TSCharacterRange{67592, 67592}, TSCharacterRange{67594, 67637}, TSCharacterRange{67639, 67640}, TSCharacterRange{67644, 67644}, TSCharacterRange{67647, 67669}, TSCharacterRange{67680, 67702}, TSCharacterRange{67712, 67742}, TSCharacterRange{67808, 67826}, TSCharacterRange{67828, 67829}, TSCharacterRange{67840, 67861}, TSCharacterRange{67872, 67897}, TSCharacterRange{67968, 68023}, TSCharacterRange{68030, 68031}, TSCharacterRange{68096, 68096}, TSCharacterRange{68112, 68115}, TSCharacterRange{68117, 68119}, TSCharacterRange{68121, 68149}, TSCharacterRange{68192, 68220}, TSCharacterRange{68224, 68252}, TSCharacterRange{68288, 68295}, TSCharacterRange{68297, 68324}, TSCharacterRange{68352, 68405}, TSCharacterRange{68416, 68437}, TSCharacterRange{68448, 68466}, TSCharacterRange{68480, 68497}, TSCharacterRange{68608, 68680}, TSCharacterRange{68736, 68786}, TSCharacterRange{68800, 68850}, TSCharacterRange{68864, 68899}, TSCharacterRange{68938, 68965}, TSCharacterRange{68975, 68997}, TSCharacterRange{69248, 69289}, TSCharacterRange{69296, 69297}, TSCharacterRange{69314, 69316}, TSCharacterRange{69376, 69404}, TSCharacterRange{69415, 69415}, TSCharacterRange{69424, 69445}, TSCharacterRange{69488, 69505}, TSCharacterRange{69552, 69572}, TSCharacterRange{69600, 69622}, TSCharacterRange{69635, 69687}, TSCharacterRange{69745, 69746}, TSCharacterRange{69749, 69749}, TSCharacterRange{69763, 69807}, TSCharacterRange{69840, 69864}, TSCharacterRange{69891, 69926}, TSCharacterRange{69956, 69956}, TSCharacterRange{69959, 69959}, TSCharacterRange{69968, 70002}, TSCharacterRange{70006, 70006}, TSCharacterRange{70019, 70066}, TSCharacterRange{70081, 70084}, TSCharacterRange{70106, 70106}, TSCharacterRange{70108, 70108}, TSCharacterRange{70144, 70161}, TSCharacterRange{70163, 70187}, TSCharacterRange{70207, 70208}, TSCharacterRange{70272, 70278}, TSCharacterRange{70280, 70280}, TSCharacterRange{70282, 70285}, TSCharacterRange{70287, 70301}, TSCharacterRange{70303, 70312}, TSCharacterRange{70320, 70366}, TSCharacterRange{70405, 70412}, TSCharacterRange{70415, 70416}, TSCharacterRange{70419, 70440}, TSCharacterRange{70442, 70448}, TSCharacterRange{70450, 70451}, TSCharacterRange{70453, 70457}, TSCharacterRange{70461, 70461}, TSCharacterRange{70480, 70480}, TSCharacterRange{70493, 70497}, TSCharacterRange{70528, 70537}, TSCharacterRange{70539, 70539}, TSCharacterRange{70542, 70542}, TSCharacterRange{70544, 70581}, TSCharacterRange{70583, 70583}, TSCharacterRange{70609, 70609}, TSCharacterRange{70611, 70611}, TSCharacterRange{70656, 70708}, TSCharacterRange{70727, 70730}, TSCharacterRange{70751, 70753}, TSCharacterRange{70784, 70831}, TSCharacterRange{70852, 70853}, TSCharacterRange{70855, 70855}, TSCharacterRange{71040, 71086}, TSCharacterRange{71128, 71131}, TSCharacterRange{71168, 71215}, TSCharacterRange{71236, 71236}, TSCharacterRange{71296, 71338}, TSCharacterRange{71352, 71352}, TSCharacterRange{71424, 71450}, TSCharacterRange{71488, 71494}, TSCharacterRange{71680, 71723}, TSCharacterRange{71840, 71903}, TSCharacterRange{71935, 71942}, TSCharacterRange{71945, 71945}, TSCharacterRange{71948, 71955}, TSCharacterRange{71957, 71958}, TSCharacterRange{71960, 71983}, TSCharacterRange{71999, 71999}, TSCharacterRange{72001, 72001}, TSCharacterRange{72096, 72103}, TSCharacterRange{72106, 72144}, TSCharacterRange{72161, 72161}, TSCharacterRange{72163, 72163}, TSCharacterRange{72192, 72192}, TSCharacterRange{72203, 72242}, TSCharacterRange{72250, 72250}, TSCharacterRange{72272, 72272}, TSCharacterRange{72284, 72329}, TSCharacterRange{72349, 72349}, TSCharacterRange{72368, 72440}, TSCharacterRange{72640, 72672}, TSCharacterRange{72704, 72712}, TSCharacterRange{72714, 72750}, TSCharacterRange{72768, 72768}, TSCharacterRange{72818, 72847}, TSCharacterRange{72960, 72966}, TSCharacterRange{72968, 72969}, TSCharacterRange{72971, 73008}, TSCharacterRange{73030, 73030}, TSCharacterRange{73056, 73061}, TSCharacterRange{73063, 73064}, TSCharacterRange{73066, 73097}, TSCharacterRange{73112, 73112}, TSCharacterRange{73440, 73458}, TSCharacterRange{73474, 73474}, TSCharacterRange{73476, 73488}, TSCharacterRange{73490, 73523}, TSCharacterRange{73648, 73648}, TSCharacterRange{73728, 74649}, TSCharacterRange{74880, 75075}, TSCharacterRange{77712, 77808}, TSCharacterRange{77824, 78895}, TSCharacterRange{78913, 78918}, TSCharacterRange{78944, 82938}, TSCharacterRange{82944, 83526}, TSCharacterRange{90368, 90397}, TSCharacterRange{92160, 92728}, TSCharacterRange{92736, 92766}, TSCharacterRange{92784, 92862}, TSCharacterRange{92880, 92909}, TSCharacterRange{92928, 92975}, TSCharacterRange{92992, 92995}, TSCharacterRange{93027, 93047}, TSCharacterRange{93053, 93071}, TSCharacterRange{93504, 93548}, TSCharacterRange{93760, 93823}, TSCharacterRange{93952, 94026}, TSCharacterRange{94032, 94032}, TSCharacterRange{94099, 94111}, TSCharacterRange{94176, 94177}, TSCharacterRange{94179, 94179}, TSCharacterRange{94208, 100343}, TSCharacterRange{100352, 101589}, TSCharacterRange{101631, 101640}, TSCharacterRange{110576, 110579}, TSCharacterRange{110581, 110587}, TSCharacterRange{110589, 110590}, TSCharacterRange{110592, 110882}, TSCharacterRange{110898, 110898}, TSCharacterRange{110928, 110930}, TSCharacterRange{110933, 110933}, TSCharacterRange{110948, 110951}, TSCharacterRange{110960, 111355}, TSCharacterRange{113664, 113770}, TSCharacterRange{113776, 113788}, TSCharacterRange{113792, 113800}, TSCharacterRange{113808, 113817}, TSCharacterRange{119808, 119892}, TSCharacterRange{119894, 119964}, TSCharacterRange{119966, 119967}, TSCharacterRange{119970, 119970}, TSCharacterRange{119973, 119974}, TSCharacterRange{119977, 119980}, TSCharacterRange{119982, 119993}, TSCharacterRange{119995, 119995}, TSCharacterRange{119997, 120003}, TSCharacterRange{120005, 120069}, TSCharacterRange{120071, 120074}, TSCharacterRange{120077, 120084}, TSCharacterRange{120086, 120092}, TSCharacterRange{120094, 120121}, TSCharacterRange{120123, 120126}, TSCharacterRange{120128, 120132}, TSCharacterRange{120134, 120134}, TSCharacterRange{120138, 120144}, TSCharacterRange{120146, 120485}, TSCharacterRange{120488, 120512}, TSCharacterRange{120514, 120538}, TSCharacterRange{120540, 120570}, TSCharacterRange{120572, 120596}, TSCharacterRange{120598, 120628}, TSCharacterRange{120630, 120654}, TSCharacterRange{120656, 120686}, TSCharacterRange{120688, 120712}, TSCharacterRange{120714, 120744}, TSCharacterRange{120746, 120770}, TSCharacterRange{120772, 120779}, TSCharacterRange{122624, 122654}, TSCharacterRange{122661, 122666}, TSCharacterRange{122928, 122989}, TSCharacterRange{123136, 123180}, TSCharacterRange{123191, 123197}, TSCharacterRange{123214, 123214}, TSCharacterRange{123536, 123565}, TSCharacterRange{123584, 123627}, TSCharacterRange{124112, 124139}, TSCharacterRange{124368, 124397}, TSCharacterRange{124400, 124400}, TSCharacterRange{124896, 124902}, TSCharacterRange{124904, 124907}, TSCharacterRange{124909, 124910}, TSCharacterRange{124912, 124926}, TSCharacterRange{124928, 125124}, TSCharacterRange{125184, 125251}, TSCharacterRange{125259, 125259}, TSCharacterRange{126464, 126467}, TSCharacterRange{126469, 126495}, TSCharacterRange{126497, 126498}, TSCharacterRange{126500, 126500}, TSCharacterRange{126503, 126503}, TSCharacterRange{126505, 126514}, TSCharacterRange{126516, 126519}, TSCharacterRange{126521, 126521}, TSCharacterRange{126523, 126523}, TSCharacterRange{126530, 126530}, TSCharacterRange{126535, 126535}, TSCharacterRange{126537, 126537}, TSCharacterRange{126539, 126539}, TSCharacterRange{126541, 126543}, TSCharacterRange{126545, 126546}, TSCharacterRange{126548, 126548}, TSCharacterRange{126551, 126551}, TSCharacterRange{126553, 126553}, TSCharacterRange{126555, 126555}, TSCharacterRange{126557, 126557}, TSCharacterRange{126559, 126559}, TSCharacterRange{126561, 126562}, TSCharacterRange{126564, 126564}, TSCharacterRange{126567, 126570}, TSCharacterRange{126572, 126578}, TSCharacterRange{126580, 126583}, TSCharacterRange{126585, 126588}, TSCharacterRange{126590, 126590}, TSCharacterRange{126592, 126601}, TSCharacterRange{126603, 126619}, TSCharacterRange{126625, 126627}, TSCharacterRange{126629, 126633}, TSCharacterRange{126635, 126651}, TSCharacterRange{131072, 173791}, TSCharacterRange{173824, 177977}, TSCharacterRange{177984, 178205}, TSCharacterRange{178208, 183969}, TSCharacterRange{183984, 191456}, TSCharacterRange{191472, 192093}, TSCharacterRange{194560, 195101}, TSCharacterRange{196608, 201546}, TSCharacterRange{201552, 205743}}

func tree_sitter_json5() *TSLanguage {
	return &tree_sitter_json5_language
}
func ts_lex(lexer *TSLexer, state int16) bool {
	var v0, v2, v4, v6, v8, v201, v202, v204, v206, v207, v209, v211, v212, v214, v218, v219, v221, v223, v224, v226, v228, v229, v231, v233, v234, v236, v238, v239, v241, v245, v246, v248, v252, v253, v255, v259, v260, v262, v266, v267, v269, v273, v274, v276, v280, v281, v283, v287, v288, v290, v294, v295, v297, v301, v302, v304, v308, v309, v311, v315, v316, v318, v322, v323, v325, v329, v330, v332, v336, v337, v339, v343, v344, v346, v350, v351, v353, v357, v358, v360, v364, v365, v367, v371, v372, v374, v377, v378, v380, v382, v383, v385, v387, v388, v390, v392, v393, v395, v397, v398, v400, v407, v408, v410, v417, v418, v420, v426, v427, v429, v433, v434, v436, v444, v445, v447, v449, v450, v452, v455, v456, v458, v460, v461, v463, v466, v467, v469, v471, v472, v474 *TSLexer
	var lexer_addr **TSLexer
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, call29, tobool32, cmp34, tobool38, cmp40, tobool44, cmp49, cmp55, tobool65, cmp70, cmp76, tobool86, cmp91, cmp97, cmp107, cmp110, cmp113, cmp117, cmp120, tobool124, cmp126, cmp130, cmp134, cmp138, cmp142, cmp145, cmp148, call152, tobool155, cmp157, cmp161, cmp165, tobool169, cmp171, cmp175, cmp179, tobool183, cmp185, cmp189, tobool193, cmp195, cmp199, cmp203, tobool207, cmp209, cmp213, tobool217, cmp219, cmp223, cmp227, cmp231, cmp235, cmp238, tobool242, cmp244, tobool248, cmp250, tobool254, cmp256, tobool260, cmp262, tobool266, cmp268, tobool272, cmp274, tobool278, cmp280, tobool284, cmp286, tobool290, cmp292, tobool296, cmp298, tobool302, cmp304, tobool308, cmp310, tobool314, cmp316, tobool320, cmp322, tobool326, cmp328, tobool332, cmp334, tobool338, cmp340, tobool344, cmp346, tobool350, cmp352, tobool356, cmp358, cmp361, cmp365, cmp368, tobool372, cmp374, cmp377, tobool381, cmp383, cmp386, cmp389, cmp392, cmp395, cmp398, tobool402, cmp404, cmp407, cmp410, cmp413, cmp416, cmp419, tobool423, cmp425, cmp428, cmp431, cmp434, cmp437, cmp440, tobool444, cmp446, cmp449, cmp452, cmp455, cmp458, cmp461, tobool465, cmp467, cmp470, cmp473, cmp476, cmp479, cmp482, tobool486, cmp488, cmp491, cmp494, cmp497, cmp500, cmp503, tobool507, cmp509, cmp512, cmp515, cmp518, cmp521, cmp524, tobool528, cmp530, cmp533, cmp536, cmp539, cmp542, cmp545, tobool549, cmp551, cmp554, cmp557, cmp560, cmp563, cmp566, tobool570, tobool572, tobool576, cmp580, cmp583, tobool587, tobool591, tobool595, tobool599, tobool603, cmp607, call611, tobool614, cmp618, call622, tobool625, cmp629, call633, tobool636, cmp640, call644, tobool647, cmp651, call655, tobool658, cmp662, call666, tobool669, cmp673, call677, tobool680, cmp684, call688, tobool691, cmp695, call699, tobool702, cmp706, call710, tobool713, cmp717, call721, tobool724, cmp728, call732, tobool735, cmp739, call743, tobool746, cmp750, call754, tobool757, cmp761, call765, tobool768, cmp772, call776, tobool779, cmp783, call787, tobool790, cmp794, call798, tobool801, cmp805, call809, tobool812, call816, tobool819, tobool823, tobool827, tobool831, tobool835, cmp839, cmp843, cmp846, cmp850, cmp853, tobool857, cmp861, cmp865, cmp868, cmp872, cmp875, tobool879, cmp883, cmp886, cmp890, cmp893, tobool897, cmp901, cmp904, tobool908, cmp912, cmp915, cmp918, cmp921, cmp924, cmp927, tobool931, tobool935, call939, tobool942, tobool946, call950, tobool953, tobool957, call961, tobool964, v477 bool
	var v7 func(*TSLexer) bool
	var eof2 *func(*TSLexer) bool
	var retval *bool
	var v9, v13, v16, v31, v34, v39, v42, v47, v50 int16
	var state_addr, arrayidx, arrayidx11, arrayidx53, arrayidx60, arrayidx74, arrayidx81, arrayidx95, arrayidx102, result_symbol, result_symbol574, result_symbol578, result_symbol589, result_symbol593, result_symbol597, result_symbol601, result_symbol605, result_symbol616, result_symbol627, result_symbol638, result_symbol649, result_symbol660, result_symbol671, result_symbol682, result_symbol693, result_symbol704, result_symbol715, result_symbol726, result_symbol737, result_symbol748, result_symbol759, result_symbol770, result_symbol781, result_symbol792, result_symbol803, result_symbol814, result_symbol821, result_symbol825, result_symbol829, result_symbol833, result_symbol837, result_symbol859, result_symbol881, result_symbol899, result_symbol910, result_symbol933, result_symbol937, result_symbol944, result_symbol948, result_symbol955, result_symbol959 *int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v25, v27, v29, v30, conv54, v32, v33, add58, v35, add63, v37, v38, conv75, v40, v41, add79, v43, add84, v45, v46, conv96, v48, v49, add100, v51, add105, v52, v53, v54, v55, v56, v58, v59, v60, v61, v62, v63, v64, v65, v67, v68, v69, v71, v72, v73, v75, v76, v78, v79, v80, v82, v83, v85, v86, v87, v88, v89, v90, v92, v94, v96, v98, v100, v102, v104, v106, v108, v110, v112, v114, v116, v118, v120, v122, v124, v126, v128, v130, v131, v132, v133, v135, v136, v138, v139, v140, v141, v142, v143, v145, v146, v147, v148, v149, v150, v152, v153, v154, v155, v156, v157, v159, v160, v161, v162, v163, v164, v166, v167, v168, v169, v170, v171, v173, v174, v175, v176, v177, v178, v180, v181, v182, v183, v184, v185, v187, v188, v189, v190, v191, v192, v194, v195, v196, v197, v198, v199, v215, v216, v242, v243, v249, v250, v256, v257, v263, v264, v270, v271, v277, v278, v284, v285, v291, v292, v298, v299, v305, v306, v312, v313, v319, v320, v326, v327, v333, v334, v340, v341, v347, v348, v354, v355, v361, v362, v368, v369, v375, v401, v402, v403, v404, v405, v411, v412, v413, v414, v415, v421, v422, v423, v424, v430, v431, v437, v438, v439, v440, v441, v442, v453, v464, v475 int32
	var lookahead, i, i46, i67, i88, lookahead1 *int32
	var conv4, idxprom, idxprom10, conv48, idxprom52, idxprom59, conv69, idxprom73, idxprom80, conv90, idxprom94, idxprom101 int64
	var v3, frombool, v10, v24, v26, v28, v36, v44, v57, v66, v70, v74, v77, v81, v84, v91, v93, v95, v97, v99, v101, v103, v105, v107, v109, v111, v113, v115, v117, v119, v121, v123, v125, v127, v129, v134, v137, v144, v151, v158, v165, v172, v179, v186, v193, v200, v205, v210, v217, v222, v227, v232, v237, v244, v251, v258, v265, v272, v279, v286, v293, v300, v307, v314, v321, v328, v335, v342, v349, v356, v363, v370, v376, v381, v386, v391, v396, v406, v416, v425, v432, v443, v448, v454, v459, v465, v470, v476 byte
	var result, skip, eof *byte
	var v203, v208, v213, v220, v225, v230, v235, v240, v247, v254, v261, v268, v275, v282, v289, v296, v303, v310, v317, v324, v331, v338, v345, v352, v359, v366, v373, v379, v384, v389, v394, v399, v409, v419, v428, v435, v446, v451, v457, v462, v468, v473 func(*TSLexer)
	var mark_end, mark_end575, mark_end579, mark_end590, mark_end594, mark_end598, mark_end602, mark_end606, mark_end617, mark_end628, mark_end639, mark_end650, mark_end661, mark_end672, mark_end683, mark_end694, mark_end705, mark_end716, mark_end727, mark_end738, mark_end749, mark_end760, mark_end771, mark_end782, mark_end793, mark_end804, mark_end815, mark_end822, mark_end826, mark_end830, mark_end834, mark_end838, mark_end860, mark_end882, mark_end900, mark_end911, mark_end934, mark_end938, mark_end945, mark_end949, mark_end956, mark_end960 *func(*TSLexer)
	var v1 func(*TSLexer, bool)
	var advance *func(*TSLexer, bool)
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i46, i67, i88, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, call29, v24, tobool32, v25, cmp34, v26, tobool38, v27, cmp40, v28, tobool44, v29, conv48, cmp49, v30, idxprom52, arrayidx53, v31, conv54, v32, cmp55, v33, add58, idxprom59, arrayidx60, v34, v35, add63, v36, tobool65, v37, conv69, cmp70, v38, idxprom73, arrayidx74, v39, conv75, v40, cmp76, v41, add79, idxprom80, arrayidx81, v42, v43, add84, v44, tobool86, v45, conv90, cmp91, v46, idxprom94, arrayidx95, v47, conv96, v48, cmp97, v49, add100, idxprom101, arrayidx102, v50, v51, add105, v52, cmp107, v53, cmp110, v54, cmp113, v55, cmp117, v56, cmp120, v57, tobool124, v58, cmp126, v59, cmp130, v60, cmp134, v61, cmp138, v62, cmp142, v63, cmp145, v64, cmp148, v65, call152, v66, tobool155, v67, cmp157, v68, cmp161, v69, cmp165, v70, tobool169, v71, cmp171, v72, cmp175, v73, cmp179, v74, tobool183, v75, cmp185, v76, cmp189, v77, tobool193, v78, cmp195, v79, cmp199, v80, cmp203, v81, tobool207, v82, cmp209, v83, cmp213, v84, tobool217, v85, cmp219, v86, cmp223, v87, cmp227, v88, cmp231, v89, cmp235, v90, cmp238, v91, tobool242, v92, cmp244, v93, tobool248, v94, cmp250, v95, tobool254, v96, cmp256, v97, tobool260, v98, cmp262, v99, tobool266, v100, cmp268, v101, tobool272, v102, cmp274, v103, tobool278, v104, cmp280, v105, tobool284, v106, cmp286, v107, tobool290, v108, cmp292, v109, tobool296, v110, cmp298, v111, tobool302, v112, cmp304, v113, tobool308, v114, cmp310, v115, tobool314, v116, cmp316, v117, tobool320, v118, cmp322, v119, tobool326, v120, cmp328, v121, tobool332, v122, cmp334, v123, tobool338, v124, cmp340, v125, tobool344, v126, cmp346, v127, tobool350, v128, cmp352, v129, tobool356, v130, cmp358, v131, cmp361, v132, cmp365, v133, cmp368, v134, tobool372, v135, cmp374, v136, cmp377, v137, tobool381, v138, cmp383, v139, cmp386, v140, cmp389, v141, cmp392, v142, cmp395, v143, cmp398, v144, tobool402, v145, cmp404, v146, cmp407, v147, cmp410, v148, cmp413, v149, cmp416, v150, cmp419, v151, tobool423, v152, cmp425, v153, cmp428, v154, cmp431, v155, cmp434, v156, cmp437, v157, cmp440, v158, tobool444, v159, cmp446, v160, cmp449, v161, cmp452, v162, cmp455, v163, cmp458, v164, cmp461, v165, tobool465, v166, cmp467, v167, cmp470, v168, cmp473, v169, cmp476, v170, cmp479, v171, cmp482, v172, tobool486, v173, cmp488, v174, cmp491, v175, cmp494, v176, cmp497, v177, cmp500, v178, cmp503, v179, tobool507, v180, cmp509, v181, cmp512, v182, cmp515, v183, cmp518, v184, cmp521, v185, cmp524, v186, tobool528, v187, cmp530, v188, cmp533, v189, cmp536, v190, cmp539, v191, cmp542, v192, cmp545, v193, tobool549, v194, cmp551, v195, cmp554, v196, cmp557, v197, cmp560, v198, cmp563, v199, cmp566, v200, tobool570, v201, result_symbol, v202, mark_end, v203, v204, v205, tobool572, v206, result_symbol574, v207, mark_end575, v208, v209, v210, tobool576, v211, result_symbol578, v212, mark_end579, v213, v214, v215, cmp580, v216, cmp583, v217, tobool587, v218, result_symbol589, v219, mark_end590, v220, v221, v222, tobool591, v223, result_symbol593, v224, mark_end594, v225, v226, v227, tobool595, v228, result_symbol597, v229, mark_end598, v230, v231, v232, tobool599, v233, result_symbol601, v234, mark_end602, v235, v236, v237, tobool603, v238, result_symbol605, v239, mark_end606, v240, v241, v242, cmp607, v243, call611, v244, tobool614, v245, result_symbol616, v246, mark_end617, v247, v248, v249, cmp618, v250, call622, v251, tobool625, v252, result_symbol627, v253, mark_end628, v254, v255, v256, cmp629, v257, call633, v258, tobool636, v259, result_symbol638, v260, mark_end639, v261, v262, v263, cmp640, v264, call644, v265, tobool647, v266, result_symbol649, v267, mark_end650, v268, v269, v270, cmp651, v271, call655, v272, tobool658, v273, result_symbol660, v274, mark_end661, v275, v276, v277, cmp662, v278, call666, v279, tobool669, v280, result_symbol671, v281, mark_end672, v282, v283, v284, cmp673, v285, call677, v286, tobool680, v287, result_symbol682, v288, mark_end683, v289, v290, v291, cmp684, v292, call688, v293, tobool691, v294, result_symbol693, v295, mark_end694, v296, v297, v298, cmp695, v299, call699, v300, tobool702, v301, result_symbol704, v302, mark_end705, v303, v304, v305, cmp706, v306, call710, v307, tobool713, v308, result_symbol715, v309, mark_end716, v310, v311, v312, cmp717, v313, call721, v314, tobool724, v315, result_symbol726, v316, mark_end727, v317, v318, v319, cmp728, v320, call732, v321, tobool735, v322, result_symbol737, v323, mark_end738, v324, v325, v326, cmp739, v327, call743, v328, tobool746, v329, result_symbol748, v330, mark_end749, v331, v332, v333, cmp750, v334, call754, v335, tobool757, v336, result_symbol759, v337, mark_end760, v338, v339, v340, cmp761, v341, call765, v342, tobool768, v343, result_symbol770, v344, mark_end771, v345, v346, v347, cmp772, v348, call776, v349, tobool779, v350, result_symbol781, v351, mark_end782, v352, v353, v354, cmp783, v355, call787, v356, tobool790, v357, result_symbol792, v358, mark_end793, v359, v360, v361, cmp794, v362, call798, v363, tobool801, v364, result_symbol803, v365, mark_end804, v366, v367, v368, cmp805, v369, call809, v370, tobool812, v371, result_symbol814, v372, mark_end815, v373, v374, v375, call816, v376, tobool819, v377, result_symbol821, v378, mark_end822, v379, v380, v381, tobool823, v382, result_symbol825, v383, mark_end826, v384, v385, v386, tobool827, v387, result_symbol829, v388, mark_end830, v389, v390, v391, tobool831, v392, result_symbol833, v393, mark_end834, v394, v395, v396, tobool835, v397, result_symbol837, v398, mark_end838, v399, v400, v401, cmp839, v402, cmp843, v403, cmp846, v404, cmp850, v405, cmp853, v406, tobool857, v407, result_symbol859, v408, mark_end860, v409, v410, v411, cmp861, v412, cmp865, v413, cmp868, v414, cmp872, v415, cmp875, v416, tobool879, v417, result_symbol881, v418, mark_end882, v419, v420, v421, cmp883, v422, cmp886, v423, cmp890, v424, cmp893, v425, tobool897, v426, result_symbol899, v427, mark_end900, v428, v429, v430, cmp901, v431, cmp904, v432, tobool908, v433, result_symbol910, v434, mark_end911, v435, v436, v437, cmp912, v438, cmp915, v439, cmp918, v440, cmp921, v441, cmp924, v442, cmp927, v443, tobool931, v444, result_symbol933, v445, mark_end934, v446, v447, v448, tobool935, v449, result_symbol937, v450, mark_end938, v451, v452, v453, call939, v454, tobool942, v455, result_symbol944, v456, mark_end945, v457, v458, v459, tobool946, v460, result_symbol948, v461, mark_end949, v462, v463, v464, call950, v465, tobool953, v466, result_symbol955, v467, mark_end956, v468, v469, v470, tobool957, v471, result_symbol959, v472, mark_end960, v473, v474, v475, call961, v476, tobool964, v477

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
	i = new(int32)
	i46 = new(int32)
	i67 = new(int32)
	i88 = new(int32)
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
		goto sw_bb33
	case 2:
		goto sw_bb39
	case 3:
		goto sw_bb45
	case 4:
		goto sw_bb66
	case 5:
		goto sw_bb87
	case 6:
		goto sw_bb125
	case 7:
		goto sw_bb156
	case 8:
		goto sw_bb170
	case 9:
		goto sw_bb184
	case 10:
		goto sw_bb194
	case 11:
		goto sw_bb208
	case 12:
		goto sw_bb218
	case 13:
		goto sw_bb243
	case 14:
		goto sw_bb249
	case 15:
		goto sw_bb255
	case 16:
		goto sw_bb261
	case 17:
		goto sw_bb267
	case 18:
		goto sw_bb273
	case 19:
		goto sw_bb279
	case 20:
		goto sw_bb285
	case 21:
		goto sw_bb291
	case 22:
		goto sw_bb297
	case 23:
		goto sw_bb303
	case 24:
		goto sw_bb309
	case 25:
		goto sw_bb315
	case 26:
		goto sw_bb321
	case 27:
		goto sw_bb327
	case 28:
		goto sw_bb333
	case 29:
		goto sw_bb339
	case 30:
		goto sw_bb345
	case 31:
		goto sw_bb351
	case 32:
		goto sw_bb357
	case 33:
		goto sw_bb373
	case 34:
		goto sw_bb382
	case 35:
		goto sw_bb403
	case 36:
		goto sw_bb424
	case 37:
		goto sw_bb445
	case 38:
		goto sw_bb466
	case 39:
		goto sw_bb487
	case 40:
		goto sw_bb508
	case 41:
		goto sw_bb529
	case 42:
		goto sw_bb550
	case 43:
		goto sw_bb571
	case 44:
		goto sw_bb573
	case 45:
		goto sw_bb577
	case 46:
		goto sw_bb588
	case 47:
		goto sw_bb592
	case 48:
		goto sw_bb596
	case 49:
		goto sw_bb600
	case 50:
		goto sw_bb604
	case 51:
		goto sw_bb615
	case 52:
		goto sw_bb626
	case 53:
		goto sw_bb637
	case 54:
		goto sw_bb648
	case 55:
		goto sw_bb659
	case 56:
		goto sw_bb670
	case 57:
		goto sw_bb681
	case 58:
		goto sw_bb692
	case 59:
		goto sw_bb703
	case 60:
		goto sw_bb714
	case 61:
		goto sw_bb725
	case 62:
		goto sw_bb736
	case 63:
		goto sw_bb747
	case 64:
		goto sw_bb758
	case 65:
		goto sw_bb769
	case 66:
		goto sw_bb780
	case 67:
		goto sw_bb791
	case 68:
		goto sw_bb802
	case 69:
		goto sw_bb813
	case 70:
		goto sw_bb820
	case 71:
		goto sw_bb824
	case 72:
		goto sw_bb828
	case 73:
		goto sw_bb832
	case 74:
		goto sw_bb836
	case 75:
		goto sw_bb858
	case 76:
		goto sw_bb880
	case 77:
		goto sw_bb898
	case 78:
		goto sw_bb909
	case 79:
		goto sw_bb932
	case 80:
		goto sw_bb936
	case 81:
		goto sw_bb943
	case 82:
		goto sw_bb947
	case 83:
		goto sw_bb954
	case 84:
		goto sw_bb958
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
	*state_addr = 43
	goto next_state

if_end:
	*i = 0
	goto for_cond

for_cond:
	v11 = *i
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(36)
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
	cmp22 = 49 <= v21
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
	*state_addr = 75
	goto next_state

if_end28:
	v23 = *lookahead
	call29 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v23)
	if call29 {
		goto if_then30
	} else {
		goto if_end31
	}

if_then30:
	*state_addr = 69
	goto next_state

if_end31:
	v24 = *result
	tobool32 = (v24 & 1) != 0
	*retval = tobool32
	goto _return

sw_bb33:
	v25 = *lookahead
	cmp34 = v25 == 10
	if cmp34 {
		goto if_then36
	} else {
		goto if_end37
	}

if_then36:
	*state_addr = 7
	goto next_state

if_end37:
	v26 = *result
	tobool38 = (v26 & 1) != 0
	*retval = tobool38
	goto _return

sw_bb39:
	v27 = *lookahead
	cmp40 = v27 == 10
	if cmp40 {
		goto if_then42
	} else {
		goto if_end43
	}

if_then42:
	*state_addr = 8
	goto next_state

if_end43:
	v28 = *result
	tobool44 = (v28 & 1) != 0
	*retval = tobool44
	goto _return

sw_bb45:
	*i46 = 0
	goto for_cond47

for_cond47:
	v29 = *i46
	conv48 = int64(uint64(uint32(v29)))
	cmp49 = uint64(conv48) < uint64(28)
	if cmp49 {
		goto for_body51
	} else {
		goto for_end64
	}

for_body51:
	v30 = *i46
	idxprom52 = int64(uint64(uint32(v30)))
	arrayidx53 = &ts_lex_map_25[idxprom52]
	v31 = *arrayidx53
	conv54 = int32(uint32(uint16(v31)))
	v32 = *lookahead
	cmp55 = conv54 == v32
	if cmp55 {
		goto if_then57
	} else {
		goto if_end61
	}

if_then57:
	v33 = *i46
	add58 = v33 + 1
	idxprom59 = int64(uint64(uint32(add58)))
	arrayidx60 = &ts_lex_map_25[idxprom59]
	v34 = *arrayidx60
	*state_addr = v34
	goto next_state

if_end61:
	goto for_inc62

for_inc62:
	v35 = *i46
	add63 = v35 + 2
	*i46 = add63
	goto for_cond47

for_end64:
	v36 = *result
	tobool65 = (v36 & 1) != 0
	*retval = tobool65
	goto _return

sw_bb66:
	*i67 = 0
	goto for_cond68

for_cond68:
	v37 = *i67
	conv69 = int64(uint64(uint32(v37)))
	cmp70 = uint64(conv69) < uint64(28)
	if cmp70 {
		goto for_body72
	} else {
		goto for_end85
	}

for_body72:
	v38 = *i67
	idxprom73 = int64(uint64(uint32(v38)))
	arrayidx74 = &ts_lex_map_26[idxprom73]
	v39 = *arrayidx74
	conv75 = int32(uint32(uint16(v39)))
	v40 = *lookahead
	cmp76 = conv75 == v40
	if cmp76 {
		goto if_then78
	} else {
		goto if_end82
	}

if_then78:
	v41 = *i67
	add79 = v41 + 1
	idxprom80 = int64(uint64(uint32(add79)))
	arrayidx81 = &ts_lex_map_26[idxprom80]
	v42 = *arrayidx81
	*state_addr = v42
	goto next_state

if_end82:
	goto for_inc83

for_inc83:
	v43 = *i67
	add84 = v43 + 2
	*i67 = add84
	goto for_cond68

for_end85:
	v44 = *result
	tobool86 = (v44 & 1) != 0
	*retval = tobool86
	goto _return

sw_bb87:
	*i88 = 0
	goto for_cond89

for_cond89:
	v45 = *i88
	conv90 = int64(uint64(uint32(v45)))
	cmp91 = uint64(conv90) < uint64(30)
	if cmp91 {
		goto for_body93
	} else {
		goto for_end106
	}

for_body93:
	v46 = *i88
	idxprom94 = int64(uint64(uint32(v46)))
	arrayidx95 = &ts_lex_map_27[idxprom94]
	v47 = *arrayidx95
	conv96 = int32(uint32(uint16(v47)))
	v48 = *lookahead
	cmp97 = conv96 == v48
	if cmp97 {
		goto if_then99
	} else {
		goto if_end103
	}

if_then99:
	v49 = *i88
	add100 = v49 + 1
	idxprom101 = int64(uint64(uint32(add100)))
	arrayidx102 = &ts_lex_map_27[idxprom101]
	v50 = *arrayidx102
	*state_addr = v50
	goto next_state

if_end103:
	goto for_inc104

for_inc104:
	v51 = *i88
	add105 = v51 + 2
	*i88 = add105
	goto for_cond89

for_end106:
	v52 = *lookahead
	cmp107 = 9 <= v52
	if cmp107 {
		goto land_lhs_true109
	} else {
		goto lor_lhs_false112
	}

land_lhs_true109:
	v53 = *lookahead
	cmp110 = v53 <= 13
	if cmp110 {
		goto if_then115
	} else {
		goto lor_lhs_false112
	}

lor_lhs_false112:
	v54 = *lookahead
	cmp113 = v54 == 32
	if cmp113 {
		goto if_then115
	} else {
		goto if_end116
	}

if_then115:
	*skip = 1
	*state_addr = 5
	goto next_state

if_end116:
	v55 = *lookahead
	cmp117 = 49 <= v55
	if cmp117 {
		goto land_lhs_true119
	} else {
		goto if_end123
	}

land_lhs_true119:
	v56 = *lookahead
	cmp120 = v56 <= 57
	if cmp120 {
		goto if_then122
	} else {
		goto if_end123
	}

if_then122:
	*state_addr = 75
	goto next_state

if_end123:
	v57 = *result
	tobool124 = (v57 & 1) != 0
	*retval = tobool124
	goto _return

sw_bb125:
	v58 = *lookahead
	cmp126 = v58 == 34
	if cmp126 {
		goto if_then128
	} else {
		goto if_end129
	}

if_then128:
	*state_addr = 7
	goto next_state

if_end129:
	v59 = *lookahead
	cmp130 = v59 == 39
	if cmp130 {
		goto if_then132
	} else {
		goto if_end133
	}

if_then132:
	*state_addr = 8
	goto next_state

if_end133:
	v60 = *lookahead
	cmp134 = v60 == 47
	if cmp134 {
		goto if_then136
	} else {
		goto if_end137
	}

if_then136:
	*state_addr = 9
	goto next_state

if_end137:
	v61 = *lookahead
	cmp138 = v61 == 125
	if cmp138 {
		goto if_then140
	} else {
		goto if_end141
	}

if_then140:
	*state_addr = 48
	goto next_state

if_end141:
	v62 = *lookahead
	cmp142 = 9 <= v62
	if cmp142 {
		goto land_lhs_true144
	} else {
		goto lor_lhs_false147
	}

land_lhs_true144:
	v63 = *lookahead
	cmp145 = v63 <= 13
	if cmp145 {
		goto if_then150
	} else {
		goto lor_lhs_false147
	}

lor_lhs_false147:
	v64 = *lookahead
	cmp148 = v64 == 32
	if cmp148 {
		goto if_then150
	} else {
		goto if_end151
	}

if_then150:
	*skip = 1
	*state_addr = 6
	goto next_state

if_end151:
	v65 = *lookahead
	call152 = set_contains(&sym_identifier_character_set_1[int64(0)], 679, v65)
	if call152 {
		goto if_then153
	} else {
		goto if_end154
	}

if_then153:
	*state_addr = 69
	goto next_state

if_end154:
	v66 = *result
	tobool155 = (v66 & 1) != 0
	*retval = tobool155
	goto _return

sw_bb156:
	v67 = *lookahead
	cmp157 = v67 == 34
	if cmp157 {
		goto if_then159
	} else {
		goto if_end160
	}

if_then159:
	*state_addr = 72
	goto next_state

if_end160:
	v68 = *lookahead
	cmp161 = v68 == 92
	if cmp161 {
		goto if_then163
	} else {
		goto if_end164
	}

if_then163:
	*state_addr = 3
	goto next_state

if_end164:
	v69 = *lookahead
	cmp165 = v69 != 0
	if cmp165 {
		goto if_then167
	} else {
		goto if_end168
	}

if_then167:
	*state_addr = 7
	goto next_state

if_end168:
	v70 = *result
	tobool169 = (v70 & 1) != 0
	*retval = tobool169
	goto _return

sw_bb170:
	v71 = *lookahead
	cmp171 = v71 == 39
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*state_addr = 72
	goto next_state

if_end174:
	v72 = *lookahead
	cmp175 = v72 == 92
	if cmp175 {
		goto if_then177
	} else {
		goto if_end178
	}

if_then177:
	*state_addr = 4
	goto next_state

if_end178:
	v73 = *lookahead
	cmp179 = v73 != 0
	if cmp179 {
		goto if_then181
	} else {
		goto if_end182
	}

if_then181:
	*state_addr = 8
	goto next_state

if_end182:
	v74 = *result
	tobool183 = (v74 & 1) != 0
	*retval = tobool183
	goto _return

sw_bb184:
	v75 = *lookahead
	cmp185 = v75 == 42
	if cmp185 {
		goto if_then187
	} else {
		goto if_end188
	}

if_then187:
	*state_addr = 11
	goto next_state

if_end188:
	v76 = *lookahead
	cmp189 = v76 == 47
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*state_addr = 45
	goto next_state

if_end192:
	v77 = *result
	tobool193 = (v77 & 1) != 0
	*retval = tobool193
	goto _return

sw_bb194:
	v78 = *lookahead
	cmp195 = v78 == 42
	if cmp195 {
		goto if_then197
	} else {
		goto if_end198
	}

if_then197:
	*state_addr = 10
	goto next_state

if_end198:
	v79 = *lookahead
	cmp199 = v79 == 47
	if cmp199 {
		goto if_then201
	} else {
		goto if_end202
	}

if_then201:
	*state_addr = 44
	goto next_state

if_end202:
	v80 = *lookahead
	cmp203 = v80 != 0
	if cmp203 {
		goto if_then205
	} else {
		goto if_end206
	}

if_then205:
	*state_addr = 11
	goto next_state

if_end206:
	v81 = *result
	tobool207 = (v81 & 1) != 0
	*retval = tobool207
	goto _return

sw_bb208:
	v82 = *lookahead
	cmp209 = v82 == 42
	if cmp209 {
		goto if_then211
	} else {
		goto if_end212
	}

if_then211:
	*state_addr = 10
	goto next_state

if_end212:
	v83 = *lookahead
	cmp213 = v83 != 0
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*state_addr = 11
	goto next_state

if_end216:
	v84 = *result
	tobool217 = (v84 & 1) != 0
	*retval = tobool217
	goto _return

sw_bb218:
	v85 = *lookahead
	cmp219 = v85 == 46
	if cmp219 {
		goto if_then221
	} else {
		goto if_end222
	}

if_then221:
	*state_addr = 76
	goto next_state

if_end222:
	v86 = *lookahead
	cmp223 = v86 == 48
	if cmp223 {
		goto if_then225
	} else {
		goto if_end226
	}

if_then225:
	*state_addr = 74
	goto next_state

if_end226:
	v87 = *lookahead
	cmp227 = v87 == 73
	if cmp227 {
		goto if_then229
	} else {
		goto if_end230
	}

if_then229:
	*state_addr = 24
	goto next_state

if_end230:
	v88 = *lookahead
	cmp231 = v88 == 78
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*state_addr = 14
	goto next_state

if_end234:
	v89 = *lookahead
	cmp235 = 49 <= v89
	if cmp235 {
		goto land_lhs_true237
	} else {
		goto if_end241
	}

land_lhs_true237:
	v90 = *lookahead
	cmp238 = v90 <= 57
	if cmp238 {
		goto if_then240
	} else {
		goto if_end241
	}

if_then240:
	*state_addr = 75
	goto next_state

if_end241:
	v91 = *result
	tobool242 = (v91 & 1) != 0
	*retval = tobool242
	goto _return

sw_bb243:
	v92 = *lookahead
	cmp244 = v92 == 78
	if cmp244 {
		goto if_then246
	} else {
		goto if_end247
	}

if_then246:
	*state_addr = 73
	goto next_state

if_end247:
	v93 = *result
	tobool248 = (v93 & 1) != 0
	*retval = tobool248
	goto _return

sw_bb249:
	v94 = *lookahead
	cmp250 = v94 == 97
	if cmp250 {
		goto if_then252
	} else {
		goto if_end253
	}

if_then252:
	*state_addr = 13
	goto next_state

if_end253:
	v95 = *result
	tobool254 = (v95 & 1) != 0
	*retval = tobool254
	goto _return

sw_bb255:
	v96 = *lookahead
	cmp256 = v96 == 97
	if cmp256 {
		goto if_then258
	} else {
		goto if_end259
	}

if_then258:
	*state_addr = 21
	goto next_state

if_end259:
	v97 = *result
	tobool260 = (v97 & 1) != 0
	*retval = tobool260
	goto _return

sw_bb261:
	v98 = *lookahead
	cmp262 = v98 == 101
	if cmp262 {
		goto if_then264
	} else {
		goto if_end265
	}

if_then264:
	*state_addr = 81
	goto next_state

if_end265:
	v99 = *result
	tobool266 = (v99 & 1) != 0
	*retval = tobool266
	goto _return

sw_bb267:
	v100 = *lookahead
	cmp268 = v100 == 101
	if cmp268 {
		goto if_then270
	} else {
		goto if_end271
	}

if_then270:
	*state_addr = 83
	goto next_state

if_end271:
	v101 = *result
	tobool272 = (v101 & 1) != 0
	*retval = tobool272
	goto _return

sw_bb273:
	v102 = *lookahead
	cmp274 = v102 == 102
	if cmp274 {
		goto if_then276
	} else {
		goto if_end277
	}

if_then276:
	*state_addr = 20
	goto next_state

if_end277:
	v103 = *result
	tobool278 = (v103 & 1) != 0
	*retval = tobool278
	goto _return

sw_bb279:
	v104 = *lookahead
	cmp280 = v104 == 105
	if cmp280 {
		goto if_then282
	} else {
		goto if_end283
	}

if_then282:
	*state_addr = 28
	goto next_state

if_end283:
	v105 = *result
	tobool284 = (v105 & 1) != 0
	*retval = tobool284
	goto _return

sw_bb285:
	v106 = *lookahead
	cmp286 = v106 == 105
	if cmp286 {
		goto if_then288
	} else {
		goto if_end289
	}

if_then288:
	*state_addr = 25
	goto next_state

if_end289:
	v107 = *result
	tobool290 = (v107 & 1) != 0
	*retval = tobool290
	goto _return

sw_bb291:
	v108 = *lookahead
	cmp292 = v108 == 108
	if cmp292 {
		goto if_then294
	} else {
		goto if_end295
	}

if_then294:
	*state_addr = 27
	goto next_state

if_end295:
	v109 = *result
	tobool296 = (v109 & 1) != 0
	*retval = tobool296
	goto _return

sw_bb297:
	v110 = *lookahead
	cmp298 = v110 == 108
	if cmp298 {
		goto if_then300
	} else {
		goto if_end301
	}

if_then300:
	*state_addr = 79
	goto next_state

if_end301:
	v111 = *result
	tobool302 = (v111 & 1) != 0
	*retval = tobool302
	goto _return

sw_bb303:
	v112 = *lookahead
	cmp304 = v112 == 108
	if cmp304 {
		goto if_then306
	} else {
		goto if_end307
	}

if_then306:
	*state_addr = 22
	goto next_state

if_end307:
	v113 = *result
	tobool308 = (v113 & 1) != 0
	*retval = tobool308
	goto _return

sw_bb309:
	v114 = *lookahead
	cmp310 = v114 == 110
	if cmp310 {
		goto if_then312
	} else {
		goto if_end313
	}

if_then312:
	*state_addr = 18
	goto next_state

if_end313:
	v115 = *result
	tobool314 = (v115 & 1) != 0
	*retval = tobool314
	goto _return

sw_bb315:
	v116 = *lookahead
	cmp316 = v116 == 110
	if cmp316 {
		goto if_then318
	} else {
		goto if_end319
	}

if_then318:
	*state_addr = 19
	goto next_state

if_end319:
	v117 = *result
	tobool320 = (v117 & 1) != 0
	*retval = tobool320
	goto _return

sw_bb321:
	v118 = *lookahead
	cmp322 = v118 == 114
	if cmp322 {
		goto if_then324
	} else {
		goto if_end325
	}

if_then324:
	*state_addr = 29
	goto next_state

if_end325:
	v119 = *result
	tobool326 = (v119 & 1) != 0
	*retval = tobool326
	goto _return

sw_bb327:
	v120 = *lookahead
	cmp328 = v120 == 115
	if cmp328 {
		goto if_then330
	} else {
		goto if_end331
	}

if_then330:
	*state_addr = 17
	goto next_state

if_end331:
	v121 = *result
	tobool332 = (v121 & 1) != 0
	*retval = tobool332
	goto _return

sw_bb333:
	v122 = *lookahead
	cmp334 = v122 == 116
	if cmp334 {
		goto if_then336
	} else {
		goto if_end337
	}

if_then336:
	*state_addr = 31
	goto next_state

if_end337:
	v123 = *result
	tobool338 = (v123 & 1) != 0
	*retval = tobool338
	goto _return

sw_bb339:
	v124 = *lookahead
	cmp340 = v124 == 117
	if cmp340 {
		goto if_then342
	} else {
		goto if_end343
	}

if_then342:
	*state_addr = 16
	goto next_state

if_end343:
	v125 = *result
	tobool344 = (v125 & 1) != 0
	*retval = tobool344
	goto _return

sw_bb345:
	v126 = *lookahead
	cmp346 = v126 == 117
	if cmp346 {
		goto if_then348
	} else {
		goto if_end349
	}

if_then348:
	*state_addr = 23
	goto next_state

if_end349:
	v127 = *result
	tobool350 = (v127 & 1) != 0
	*retval = tobool350
	goto _return

sw_bb351:
	v128 = *lookahead
	cmp352 = v128 == 121
	if cmp352 {
		goto if_then354
	} else {
		goto if_end355
	}

if_then354:
	*state_addr = 73
	goto next_state

if_end355:
	v129 = *result
	tobool356 = (v129 & 1) != 0
	*retval = tobool356
	goto _return

sw_bb357:
	v130 = *lookahead
	cmp358 = v130 == 43
	if cmp358 {
		goto if_then363
	} else {
		goto lor_lhs_false360
	}

lor_lhs_false360:
	v131 = *lookahead
	cmp361 = v131 == 45
	if cmp361 {
		goto if_then363
	} else {
		goto if_end364
	}

if_then363:
	*state_addr = 33
	goto next_state

if_end364:
	v132 = *lookahead
	cmp365 = 48 <= v132
	if cmp365 {
		goto land_lhs_true367
	} else {
		goto if_end371
	}

land_lhs_true367:
	v133 = *lookahead
	cmp368 = v133 <= 57
	if cmp368 {
		goto if_then370
	} else {
		goto if_end371
	}

if_then370:
	*state_addr = 77
	goto next_state

if_end371:
	v134 = *result
	tobool372 = (v134 & 1) != 0
	*retval = tobool372
	goto _return

sw_bb373:
	v135 = *lookahead
	cmp374 = 48 <= v135
	if cmp374 {
		goto land_lhs_true376
	} else {
		goto if_end380
	}

land_lhs_true376:
	v136 = *lookahead
	cmp377 = v136 <= 57
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*state_addr = 77
	goto next_state

if_end380:
	v137 = *result
	tobool381 = (v137 & 1) != 0
	*retval = tobool381
	goto _return

sw_bb382:
	v138 = *lookahead
	cmp383 = 48 <= v138
	if cmp383 {
		goto land_lhs_true385
	} else {
		goto lor_lhs_false388
	}

land_lhs_true385:
	v139 = *lookahead
	cmp386 = v139 <= 57
	if cmp386 {
		goto if_then400
	} else {
		goto lor_lhs_false388
	}

lor_lhs_false388:
	v140 = *lookahead
	cmp389 = 65 <= v140
	if cmp389 {
		goto land_lhs_true391
	} else {
		goto lor_lhs_false394
	}

land_lhs_true391:
	v141 = *lookahead
	cmp392 = v141 <= 70
	if cmp392 {
		goto if_then400
	} else {
		goto lor_lhs_false394
	}

lor_lhs_false394:
	v142 = *lookahead
	cmp395 = 97 <= v142
	if cmp395 {
		goto land_lhs_true397
	} else {
		goto if_end401
	}

land_lhs_true397:
	v143 = *lookahead
	cmp398 = v143 <= 102
	if cmp398 {
		goto if_then400
	} else {
		goto if_end401
	}

if_then400:
	*state_addr = 7
	goto next_state

if_end401:
	v144 = *result
	tobool402 = (v144 & 1) != 0
	*retval = tobool402
	goto _return

sw_bb403:
	v145 = *lookahead
	cmp404 = 48 <= v145
	if cmp404 {
		goto land_lhs_true406
	} else {
		goto lor_lhs_false409
	}

land_lhs_true406:
	v146 = *lookahead
	cmp407 = v146 <= 57
	if cmp407 {
		goto if_then421
	} else {
		goto lor_lhs_false409
	}

lor_lhs_false409:
	v147 = *lookahead
	cmp410 = 65 <= v147
	if cmp410 {
		goto land_lhs_true412
	} else {
		goto lor_lhs_false415
	}

land_lhs_true412:
	v148 = *lookahead
	cmp413 = v148 <= 70
	if cmp413 {
		goto if_then421
	} else {
		goto lor_lhs_false415
	}

lor_lhs_false415:
	v149 = *lookahead
	cmp416 = 97 <= v149
	if cmp416 {
		goto land_lhs_true418
	} else {
		goto if_end422
	}

land_lhs_true418:
	v150 = *lookahead
	cmp419 = v150 <= 102
	if cmp419 {
		goto if_then421
	} else {
		goto if_end422
	}

if_then421:
	*state_addr = 8
	goto next_state

if_end422:
	v151 = *result
	tobool423 = (v151 & 1) != 0
	*retval = tobool423
	goto _return

sw_bb424:
	v152 = *lookahead
	cmp425 = 48 <= v152
	if cmp425 {
		goto land_lhs_true427
	} else {
		goto lor_lhs_false430
	}

land_lhs_true427:
	v153 = *lookahead
	cmp428 = v153 <= 57
	if cmp428 {
		goto if_then442
	} else {
		goto lor_lhs_false430
	}

lor_lhs_false430:
	v154 = *lookahead
	cmp431 = 65 <= v154
	if cmp431 {
		goto land_lhs_true433
	} else {
		goto lor_lhs_false436
	}

land_lhs_true433:
	v155 = *lookahead
	cmp434 = v155 <= 70
	if cmp434 {
		goto if_then442
	} else {
		goto lor_lhs_false436
	}

lor_lhs_false436:
	v156 = *lookahead
	cmp437 = 97 <= v156
	if cmp437 {
		goto land_lhs_true439
	} else {
		goto if_end443
	}

land_lhs_true439:
	v157 = *lookahead
	cmp440 = v157 <= 102
	if cmp440 {
		goto if_then442
	} else {
		goto if_end443
	}

if_then442:
	*state_addr = 78
	goto next_state

if_end443:
	v158 = *result
	tobool444 = (v158 & 1) != 0
	*retval = tobool444
	goto _return

sw_bb445:
	v159 = *lookahead
	cmp446 = 48 <= v159
	if cmp446 {
		goto land_lhs_true448
	} else {
		goto lor_lhs_false451
	}

land_lhs_true448:
	v160 = *lookahead
	cmp449 = v160 <= 57
	if cmp449 {
		goto if_then463
	} else {
		goto lor_lhs_false451
	}

lor_lhs_false451:
	v161 = *lookahead
	cmp452 = 65 <= v161
	if cmp452 {
		goto land_lhs_true454
	} else {
		goto lor_lhs_false457
	}

land_lhs_true454:
	v162 = *lookahead
	cmp455 = v162 <= 70
	if cmp455 {
		goto if_then463
	} else {
		goto lor_lhs_false457
	}

lor_lhs_false457:
	v163 = *lookahead
	cmp458 = 97 <= v163
	if cmp458 {
		goto land_lhs_true460
	} else {
		goto if_end464
	}

land_lhs_true460:
	v164 = *lookahead
	cmp461 = v164 <= 102
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*state_addr = 34
	goto next_state

if_end464:
	v165 = *result
	tobool465 = (v165 & 1) != 0
	*retval = tobool465
	goto _return

sw_bb466:
	v166 = *lookahead
	cmp467 = 48 <= v166
	if cmp467 {
		goto land_lhs_true469
	} else {
		goto lor_lhs_false472
	}

land_lhs_true469:
	v167 = *lookahead
	cmp470 = v167 <= 57
	if cmp470 {
		goto if_then484
	} else {
		goto lor_lhs_false472
	}

lor_lhs_false472:
	v168 = *lookahead
	cmp473 = 65 <= v168
	if cmp473 {
		goto land_lhs_true475
	} else {
		goto lor_lhs_false478
	}

land_lhs_true475:
	v169 = *lookahead
	cmp476 = v169 <= 70
	if cmp476 {
		goto if_then484
	} else {
		goto lor_lhs_false478
	}

lor_lhs_false478:
	v170 = *lookahead
	cmp479 = 97 <= v170
	if cmp479 {
		goto land_lhs_true481
	} else {
		goto if_end485
	}

land_lhs_true481:
	v171 = *lookahead
	cmp482 = v171 <= 102
	if cmp482 {
		goto if_then484
	} else {
		goto if_end485
	}

if_then484:
	*state_addr = 35
	goto next_state

if_end485:
	v172 = *result
	tobool486 = (v172 & 1) != 0
	*retval = tobool486
	goto _return

sw_bb487:
	v173 = *lookahead
	cmp488 = 48 <= v173
	if cmp488 {
		goto land_lhs_true490
	} else {
		goto lor_lhs_false493
	}

land_lhs_true490:
	v174 = *lookahead
	cmp491 = v174 <= 57
	if cmp491 {
		goto if_then505
	} else {
		goto lor_lhs_false493
	}

lor_lhs_false493:
	v175 = *lookahead
	cmp494 = 65 <= v175
	if cmp494 {
		goto land_lhs_true496
	} else {
		goto lor_lhs_false499
	}

land_lhs_true496:
	v176 = *lookahead
	cmp497 = v176 <= 70
	if cmp497 {
		goto if_then505
	} else {
		goto lor_lhs_false499
	}

lor_lhs_false499:
	v177 = *lookahead
	cmp500 = 97 <= v177
	if cmp500 {
		goto land_lhs_true502
	} else {
		goto if_end506
	}

land_lhs_true502:
	v178 = *lookahead
	cmp503 = v178 <= 102
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*state_addr = 37
	goto next_state

if_end506:
	v179 = *result
	tobool507 = (v179 & 1) != 0
	*retval = tobool507
	goto _return

sw_bb508:
	v180 = *lookahead
	cmp509 = 48 <= v180
	if cmp509 {
		goto land_lhs_true511
	} else {
		goto lor_lhs_false514
	}

land_lhs_true511:
	v181 = *lookahead
	cmp512 = v181 <= 57
	if cmp512 {
		goto if_then526
	} else {
		goto lor_lhs_false514
	}

lor_lhs_false514:
	v182 = *lookahead
	cmp515 = 65 <= v182
	if cmp515 {
		goto land_lhs_true517
	} else {
		goto lor_lhs_false520
	}

land_lhs_true517:
	v183 = *lookahead
	cmp518 = v183 <= 70
	if cmp518 {
		goto if_then526
	} else {
		goto lor_lhs_false520
	}

lor_lhs_false520:
	v184 = *lookahead
	cmp521 = 97 <= v184
	if cmp521 {
		goto land_lhs_true523
	} else {
		goto if_end527
	}

land_lhs_true523:
	v185 = *lookahead
	cmp524 = v185 <= 102
	if cmp524 {
		goto if_then526
	} else {
		goto if_end527
	}

if_then526:
	*state_addr = 38
	goto next_state

if_end527:
	v186 = *result
	tobool528 = (v186 & 1) != 0
	*retval = tobool528
	goto _return

sw_bb529:
	v187 = *lookahead
	cmp530 = 48 <= v187
	if cmp530 {
		goto land_lhs_true532
	} else {
		goto lor_lhs_false535
	}

land_lhs_true532:
	v188 = *lookahead
	cmp533 = v188 <= 57
	if cmp533 {
		goto if_then547
	} else {
		goto lor_lhs_false535
	}

lor_lhs_false535:
	v189 = *lookahead
	cmp536 = 65 <= v189
	if cmp536 {
		goto land_lhs_true538
	} else {
		goto lor_lhs_false541
	}

land_lhs_true538:
	v190 = *lookahead
	cmp539 = v190 <= 70
	if cmp539 {
		goto if_then547
	} else {
		goto lor_lhs_false541
	}

lor_lhs_false541:
	v191 = *lookahead
	cmp542 = 97 <= v191
	if cmp542 {
		goto land_lhs_true544
	} else {
		goto if_end548
	}

land_lhs_true544:
	v192 = *lookahead
	cmp545 = v192 <= 102
	if cmp545 {
		goto if_then547
	} else {
		goto if_end548
	}

if_then547:
	*state_addr = 39
	goto next_state

if_end548:
	v193 = *result
	tobool549 = (v193 & 1) != 0
	*retval = tobool549
	goto _return

sw_bb550:
	v194 = *lookahead
	cmp551 = 48 <= v194
	if cmp551 {
		goto land_lhs_true553
	} else {
		goto lor_lhs_false556
	}

land_lhs_true553:
	v195 = *lookahead
	cmp554 = v195 <= 57
	if cmp554 {
		goto if_then568
	} else {
		goto lor_lhs_false556
	}

lor_lhs_false556:
	v196 = *lookahead
	cmp557 = 65 <= v196
	if cmp557 {
		goto land_lhs_true559
	} else {
		goto lor_lhs_false562
	}

land_lhs_true559:
	v197 = *lookahead
	cmp560 = v197 <= 70
	if cmp560 {
		goto if_then568
	} else {
		goto lor_lhs_false562
	}

lor_lhs_false562:
	v198 = *lookahead
	cmp563 = 97 <= v198
	if cmp563 {
		goto land_lhs_true565
	} else {
		goto if_end569
	}

land_lhs_true565:
	v199 = *lookahead
	cmp566 = v199 <= 102
	if cmp566 {
		goto if_then568
	} else {
		goto if_end569
	}

if_then568:
	*state_addr = 40
	goto next_state

if_end569:
	v200 = *result
	tobool570 = (v200 & 1) != 0
	*retval = tobool570
	goto _return

sw_bb571:
	*result = 1
	v201 = *lexer_addr
	result_symbol = &v201.F1
	*result_symbol = 0
	v202 = *lexer_addr
	mark_end = &v202.F3
	v203 = *mark_end
	v204 = *lexer_addr
	v203(v204)
	v205 = *result
	tobool572 = (v205 & 1) != 0
	*retval = tobool572
	goto _return

sw_bb573:
	*result = 1
	v206 = *lexer_addr
	result_symbol574 = &v206.F1
	*result_symbol574 = 1
	v207 = *lexer_addr
	mark_end575 = &v207.F3
	v208 = *mark_end575
	v209 = *lexer_addr
	v208(v209)
	v210 = *result
	tobool576 = (v210 & 1) != 0
	*retval = tobool576
	goto _return

sw_bb577:
	*result = 1
	v211 = *lexer_addr
	result_symbol578 = &v211.F1
	*result_symbol578 = 1
	v212 = *lexer_addr
	mark_end579 = &v212.F3
	v213 = *mark_end579
	v214 = *lexer_addr
	v213(v214)
	v215 = *lookahead
	cmp580 = v215 != 0
	if cmp580 {
		goto land_lhs_true582
	} else {
		goto if_end586
	}

land_lhs_true582:
	v216 = *lookahead
	cmp583 = v216 != 10
	if cmp583 {
		goto if_then585
	} else {
		goto if_end586
	}

if_then585:
	*state_addr = 45
	goto next_state

if_end586:
	v217 = *result
	tobool587 = (v217 & 1) != 0
	*retval = tobool587
	goto _return

sw_bb588:
	*result = 1
	v218 = *lexer_addr
	result_symbol589 = &v218.F1
	*result_symbol589 = 2
	v219 = *lexer_addr
	mark_end590 = &v219.F3
	v220 = *mark_end590
	v221 = *lexer_addr
	v220(v221)
	v222 = *result
	tobool591 = (v222 & 1) != 0
	*retval = tobool591
	goto _return

sw_bb592:
	*result = 1
	v223 = *lexer_addr
	result_symbol593 = &v223.F1
	*result_symbol593 = 3
	v224 = *lexer_addr
	mark_end594 = &v224.F3
	v225 = *mark_end594
	v226 = *lexer_addr
	v225(v226)
	v227 = *result
	tobool595 = (v227 & 1) != 0
	*retval = tobool595
	goto _return

sw_bb596:
	*result = 1
	v228 = *lexer_addr
	result_symbol597 = &v228.F1
	*result_symbol597 = 4
	v229 = *lexer_addr
	mark_end598 = &v229.F3
	v230 = *mark_end598
	v231 = *lexer_addr
	v230(v231)
	v232 = *result
	tobool599 = (v232 & 1) != 0
	*retval = tobool599
	goto _return

sw_bb600:
	*result = 1
	v233 = *lexer_addr
	result_symbol601 = &v233.F1
	*result_symbol601 = 5
	v234 = *lexer_addr
	mark_end602 = &v234.F3
	v235 = *mark_end602
	v236 = *lexer_addr
	v235(v236)
	v237 = *result
	tobool603 = (v237 & 1) != 0
	*retval = tobool603
	goto _return

sw_bb604:
	*result = 1
	v238 = *lexer_addr
	result_symbol605 = &v238.F1
	*result_symbol605 = 6
	v239 = *lexer_addr
	mark_end606 = &v239.F3
	v240 = *mark_end606
	v241 = *lexer_addr
	v240(v241)
	v242 = *lookahead
	cmp607 = v242 == 78
	if cmp607 {
		goto if_then609
	} else {
		goto if_end610
	}

if_then609:
	*state_addr = 69
	goto next_state

if_end610:
	v243 = *lookahead
	call611 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v243)
	if call611 {
		goto if_then612
	} else {
		goto if_end613
	}

if_then612:
	*state_addr = 69
	goto next_state

if_end613:
	v244 = *result
	tobool614 = (v244 & 1) != 0
	*retval = tobool614
	goto _return

sw_bb615:
	*result = 1
	v245 = *lexer_addr
	result_symbol616 = &v245.F1
	*result_symbol616 = 6
	v246 = *lexer_addr
	mark_end617 = &v246.F3
	v247 = *mark_end617
	v248 = *lexer_addr
	v247(v248)
	v249 = *lookahead
	cmp618 = v249 == 97
	if cmp618 {
		goto if_then620
	} else {
		goto if_end621
	}

if_then620:
	*state_addr = 50
	goto next_state

if_end621:
	v250 = *lookahead
	call622 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v250)
	if call622 {
		goto if_then623
	} else {
		goto if_end624
	}

if_then623:
	*state_addr = 69
	goto next_state

if_end624:
	v251 = *result
	tobool625 = (v251 & 1) != 0
	*retval = tobool625
	goto _return

sw_bb626:
	*result = 1
	v252 = *lexer_addr
	result_symbol627 = &v252.F1
	*result_symbol627 = 6
	v253 = *lexer_addr
	mark_end628 = &v253.F3
	v254 = *mark_end628
	v255 = *lexer_addr
	v254(v255)
	v256 = *lookahead
	cmp629 = v256 == 97
	if cmp629 {
		goto if_then631
	} else {
		goto if_end632
	}

if_then631:
	*state_addr = 58
	goto next_state

if_end632:
	v257 = *lookahead
	call633 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v257)
	if call633 {
		goto if_then634
	} else {
		goto if_end635
	}

if_then634:
	*state_addr = 69
	goto next_state

if_end635:
	v258 = *result
	tobool636 = (v258 & 1) != 0
	*retval = tobool636
	goto _return

sw_bb637:
	*result = 1
	v259 = *lexer_addr
	result_symbol638 = &v259.F1
	*result_symbol638 = 6
	v260 = *lexer_addr
	mark_end639 = &v260.F3
	v261 = *mark_end639
	v262 = *lexer_addr
	v261(v262)
	v263 = *lookahead
	cmp640 = v263 == 101
	if cmp640 {
		goto if_then642
	} else {
		goto if_end643
	}

if_then642:
	*state_addr = 82
	goto next_state

if_end643:
	v264 = *lookahead
	call644 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v264)
	if call644 {
		goto if_then645
	} else {
		goto if_end646
	}

if_then645:
	*state_addr = 69
	goto next_state

if_end646:
	v265 = *result
	tobool647 = (v265 & 1) != 0
	*retval = tobool647
	goto _return

sw_bb648:
	*result = 1
	v266 = *lexer_addr
	result_symbol649 = &v266.F1
	*result_symbol649 = 6
	v267 = *lexer_addr
	mark_end650 = &v267.F3
	v268 = *mark_end650
	v269 = *lexer_addr
	v268(v269)
	v270 = *lookahead
	cmp651 = v270 == 101
	if cmp651 {
		goto if_then653
	} else {
		goto if_end654
	}

if_then653:
	*state_addr = 84
	goto next_state

if_end654:
	v271 = *lookahead
	call655 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v271)
	if call655 {
		goto if_then656
	} else {
		goto if_end657
	}

if_then656:
	*state_addr = 69
	goto next_state

if_end657:
	v272 = *result
	tobool658 = (v272 & 1) != 0
	*retval = tobool658
	goto _return

sw_bb659:
	*result = 1
	v273 = *lexer_addr
	result_symbol660 = &v273.F1
	*result_symbol660 = 6
	v274 = *lexer_addr
	mark_end661 = &v274.F3
	v275 = *mark_end661
	v276 = *lexer_addr
	v275(v276)
	v277 = *lookahead
	cmp662 = v277 == 102
	if cmp662 {
		goto if_then664
	} else {
		goto if_end665
	}

if_then664:
	*state_addr = 57
	goto next_state

if_end665:
	v278 = *lookahead
	call666 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v278)
	if call666 {
		goto if_then667
	} else {
		goto if_end668
	}

if_then667:
	*state_addr = 69
	goto next_state

if_end668:
	v279 = *result
	tobool669 = (v279 & 1) != 0
	*retval = tobool669
	goto _return

sw_bb670:
	*result = 1
	v280 = *lexer_addr
	result_symbol671 = &v280.F1
	*result_symbol671 = 6
	v281 = *lexer_addr
	mark_end672 = &v281.F3
	v282 = *mark_end672
	v283 = *lexer_addr
	v282(v283)
	v284 = *lookahead
	cmp673 = v284 == 105
	if cmp673 {
		goto if_then675
	} else {
		goto if_end676
	}

if_then675:
	*state_addr = 65
	goto next_state

if_end676:
	v285 = *lookahead
	call677 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v285)
	if call677 {
		goto if_then678
	} else {
		goto if_end679
	}

if_then678:
	*state_addr = 69
	goto next_state

if_end679:
	v286 = *result
	tobool680 = (v286 & 1) != 0
	*retval = tobool680
	goto _return

sw_bb681:
	*result = 1
	v287 = *lexer_addr
	result_symbol682 = &v287.F1
	*result_symbol682 = 6
	v288 = *lexer_addr
	mark_end683 = &v288.F3
	v289 = *mark_end683
	v290 = *lexer_addr
	v289(v290)
	v291 = *lookahead
	cmp684 = v291 == 105
	if cmp684 {
		goto if_then686
	} else {
		goto if_end687
	}

if_then686:
	*state_addr = 62
	goto next_state

if_end687:
	v292 = *lookahead
	call688 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v292)
	if call688 {
		goto if_then689
	} else {
		goto if_end690
	}

if_then689:
	*state_addr = 69
	goto next_state

if_end690:
	v293 = *result
	tobool691 = (v293 & 1) != 0
	*retval = tobool691
	goto _return

sw_bb692:
	*result = 1
	v294 = *lexer_addr
	result_symbol693 = &v294.F1
	*result_symbol693 = 6
	v295 = *lexer_addr
	mark_end694 = &v295.F3
	v296 = *mark_end694
	v297 = *lexer_addr
	v296(v297)
	v298 = *lookahead
	cmp695 = v298 == 108
	if cmp695 {
		goto if_then697
	} else {
		goto if_end698
	}

if_then697:
	*state_addr = 64
	goto next_state

if_end698:
	v299 = *lookahead
	call699 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v299)
	if call699 {
		goto if_then700
	} else {
		goto if_end701
	}

if_then700:
	*state_addr = 69
	goto next_state

if_end701:
	v300 = *result
	tobool702 = (v300 & 1) != 0
	*retval = tobool702
	goto _return

sw_bb703:
	*result = 1
	v301 = *lexer_addr
	result_symbol704 = &v301.F1
	*result_symbol704 = 6
	v302 = *lexer_addr
	mark_end705 = &v302.F3
	v303 = *mark_end705
	v304 = *lexer_addr
	v303(v304)
	v305 = *lookahead
	cmp706 = v305 == 108
	if cmp706 {
		goto if_then708
	} else {
		goto if_end709
	}

if_then708:
	*state_addr = 80
	goto next_state

if_end709:
	v306 = *lookahead
	call710 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v306)
	if call710 {
		goto if_then711
	} else {
		goto if_end712
	}

if_then711:
	*state_addr = 69
	goto next_state

if_end712:
	v307 = *result
	tobool713 = (v307 & 1) != 0
	*retval = tobool713
	goto _return

sw_bb714:
	*result = 1
	v308 = *lexer_addr
	result_symbol715 = &v308.F1
	*result_symbol715 = 6
	v309 = *lexer_addr
	mark_end716 = &v309.F3
	v310 = *mark_end716
	v311 = *lexer_addr
	v310(v311)
	v312 = *lookahead
	cmp717 = v312 == 108
	if cmp717 {
		goto if_then719
	} else {
		goto if_end720
	}

if_then719:
	*state_addr = 59
	goto next_state

if_end720:
	v313 = *lookahead
	call721 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v313)
	if call721 {
		goto if_then722
	} else {
		goto if_end723
	}

if_then722:
	*state_addr = 69
	goto next_state

if_end723:
	v314 = *result
	tobool724 = (v314 & 1) != 0
	*retval = tobool724
	goto _return

sw_bb725:
	*result = 1
	v315 = *lexer_addr
	result_symbol726 = &v315.F1
	*result_symbol726 = 6
	v316 = *lexer_addr
	mark_end727 = &v316.F3
	v317 = *mark_end727
	v318 = *lexer_addr
	v317(v318)
	v319 = *lookahead
	cmp728 = v319 == 110
	if cmp728 {
		goto if_then730
	} else {
		goto if_end731
	}

if_then730:
	*state_addr = 55
	goto next_state

if_end731:
	v320 = *lookahead
	call732 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v320)
	if call732 {
		goto if_then733
	} else {
		goto if_end734
	}

if_then733:
	*state_addr = 69
	goto next_state

if_end734:
	v321 = *result
	tobool735 = (v321 & 1) != 0
	*retval = tobool735
	goto _return

sw_bb736:
	*result = 1
	v322 = *lexer_addr
	result_symbol737 = &v322.F1
	*result_symbol737 = 6
	v323 = *lexer_addr
	mark_end738 = &v323.F3
	v324 = *mark_end738
	v325 = *lexer_addr
	v324(v325)
	v326 = *lookahead
	cmp739 = v326 == 110
	if cmp739 {
		goto if_then741
	} else {
		goto if_end742
	}

if_then741:
	*state_addr = 56
	goto next_state

if_end742:
	v327 = *lookahead
	call743 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v327)
	if call743 {
		goto if_then744
	} else {
		goto if_end745
	}

if_then744:
	*state_addr = 69
	goto next_state

if_end745:
	v328 = *result
	tobool746 = (v328 & 1) != 0
	*retval = tobool746
	goto _return

sw_bb747:
	*result = 1
	v329 = *lexer_addr
	result_symbol748 = &v329.F1
	*result_symbol748 = 6
	v330 = *lexer_addr
	mark_end749 = &v330.F3
	v331 = *mark_end749
	v332 = *lexer_addr
	v331(v332)
	v333 = *lookahead
	cmp750 = v333 == 114
	if cmp750 {
		goto if_then752
	} else {
		goto if_end753
	}

if_then752:
	*state_addr = 66
	goto next_state

if_end753:
	v334 = *lookahead
	call754 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v334)
	if call754 {
		goto if_then755
	} else {
		goto if_end756
	}

if_then755:
	*state_addr = 69
	goto next_state

if_end756:
	v335 = *result
	tobool757 = (v335 & 1) != 0
	*retval = tobool757
	goto _return

sw_bb758:
	*result = 1
	v336 = *lexer_addr
	result_symbol759 = &v336.F1
	*result_symbol759 = 6
	v337 = *lexer_addr
	mark_end760 = &v337.F3
	v338 = *mark_end760
	v339 = *lexer_addr
	v338(v339)
	v340 = *lookahead
	cmp761 = v340 == 115
	if cmp761 {
		goto if_then763
	} else {
		goto if_end764
	}

if_then763:
	*state_addr = 54
	goto next_state

if_end764:
	v341 = *lookahead
	call765 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v341)
	if call765 {
		goto if_then766
	} else {
		goto if_end767
	}

if_then766:
	*state_addr = 69
	goto next_state

if_end767:
	v342 = *result
	tobool768 = (v342 & 1) != 0
	*retval = tobool768
	goto _return

sw_bb769:
	*result = 1
	v343 = *lexer_addr
	result_symbol770 = &v343.F1
	*result_symbol770 = 6
	v344 = *lexer_addr
	mark_end771 = &v344.F3
	v345 = *mark_end771
	v346 = *lexer_addr
	v345(v346)
	v347 = *lookahead
	cmp772 = v347 == 116
	if cmp772 {
		goto if_then774
	} else {
		goto if_end775
	}

if_then774:
	*state_addr = 68
	goto next_state

if_end775:
	v348 = *lookahead
	call776 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v348)
	if call776 {
		goto if_then777
	} else {
		goto if_end778
	}

if_then777:
	*state_addr = 69
	goto next_state

if_end778:
	v349 = *result
	tobool779 = (v349 & 1) != 0
	*retval = tobool779
	goto _return

sw_bb780:
	*result = 1
	v350 = *lexer_addr
	result_symbol781 = &v350.F1
	*result_symbol781 = 6
	v351 = *lexer_addr
	mark_end782 = &v351.F3
	v352 = *mark_end782
	v353 = *lexer_addr
	v352(v353)
	v354 = *lookahead
	cmp783 = v354 == 117
	if cmp783 {
		goto if_then785
	} else {
		goto if_end786
	}

if_then785:
	*state_addr = 53
	goto next_state

if_end786:
	v355 = *lookahead
	call787 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v355)
	if call787 {
		goto if_then788
	} else {
		goto if_end789
	}

if_then788:
	*state_addr = 69
	goto next_state

if_end789:
	v356 = *result
	tobool790 = (v356 & 1) != 0
	*retval = tobool790
	goto _return

sw_bb791:
	*result = 1
	v357 = *lexer_addr
	result_symbol792 = &v357.F1
	*result_symbol792 = 6
	v358 = *lexer_addr
	mark_end793 = &v358.F3
	v359 = *mark_end793
	v360 = *lexer_addr
	v359(v360)
	v361 = *lookahead
	cmp794 = v361 == 117
	if cmp794 {
		goto if_then796
	} else {
		goto if_end797
	}

if_then796:
	*state_addr = 60
	goto next_state

if_end797:
	v362 = *lookahead
	call798 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v362)
	if call798 {
		goto if_then799
	} else {
		goto if_end800
	}

if_then799:
	*state_addr = 69
	goto next_state

if_end800:
	v363 = *result
	tobool801 = (v363 & 1) != 0
	*retval = tobool801
	goto _return

sw_bb802:
	*result = 1
	v364 = *lexer_addr
	result_symbol803 = &v364.F1
	*result_symbol803 = 6
	v365 = *lexer_addr
	mark_end804 = &v365.F3
	v366 = *mark_end804
	v367 = *lexer_addr
	v366(v367)
	v368 = *lookahead
	cmp805 = v368 == 121
	if cmp805 {
		goto if_then807
	} else {
		goto if_end808
	}

if_then807:
	*state_addr = 69
	goto next_state

if_end808:
	v369 = *lookahead
	call809 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v369)
	if call809 {
		goto if_then810
	} else {
		goto if_end811
	}

if_then810:
	*state_addr = 69
	goto next_state

if_end811:
	v370 = *result
	tobool812 = (v370 & 1) != 0
	*retval = tobool812
	goto _return

sw_bb813:
	*result = 1
	v371 = *lexer_addr
	result_symbol814 = &v371.F1
	*result_symbol814 = 6
	v372 = *lexer_addr
	mark_end815 = &v372.F3
	v373 = *mark_end815
	v374 = *lexer_addr
	v373(v374)
	v375 = *lookahead
	call816 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v375)
	if call816 {
		goto if_then817
	} else {
		goto if_end818
	}

if_then817:
	*state_addr = 69
	goto next_state

if_end818:
	v376 = *result
	tobool819 = (v376 & 1) != 0
	*retval = tobool819
	goto _return

sw_bb820:
	*result = 1
	v377 = *lexer_addr
	result_symbol821 = &v377.F1
	*result_symbol821 = 7
	v378 = *lexer_addr
	mark_end822 = &v378.F3
	v379 = *mark_end822
	v380 = *lexer_addr
	v379(v380)
	v381 = *result
	tobool823 = (v381 & 1) != 0
	*retval = tobool823
	goto _return

sw_bb824:
	*result = 1
	v382 = *lexer_addr
	result_symbol825 = &v382.F1
	*result_symbol825 = 8
	v383 = *lexer_addr
	mark_end826 = &v383.F3
	v384 = *mark_end826
	v385 = *lexer_addr
	v384(v385)
	v386 = *result
	tobool827 = (v386 & 1) != 0
	*retval = tobool827
	goto _return

sw_bb828:
	*result = 1
	v387 = *lexer_addr
	result_symbol829 = &v387.F1
	*result_symbol829 = 9
	v388 = *lexer_addr
	mark_end830 = &v388.F3
	v389 = *mark_end830
	v390 = *lexer_addr
	v389(v390)
	v391 = *result
	tobool831 = (v391 & 1) != 0
	*retval = tobool831
	goto _return

sw_bb832:
	*result = 1
	v392 = *lexer_addr
	result_symbol833 = &v392.F1
	*result_symbol833 = 10
	v393 = *lexer_addr
	mark_end834 = &v393.F3
	v394 = *mark_end834
	v395 = *lexer_addr
	v394(v395)
	v396 = *result
	tobool835 = (v396 & 1) != 0
	*retval = tobool835
	goto _return

sw_bb836:
	*result = 1
	v397 = *lexer_addr
	result_symbol837 = &v397.F1
	*result_symbol837 = 10
	v398 = *lexer_addr
	mark_end838 = &v398.F3
	v399 = *mark_end838
	v400 = *lexer_addr
	v399(v400)
	v401 = *lookahead
	cmp839 = v401 == 46
	if cmp839 {
		goto if_then841
	} else {
		goto if_end842
	}

if_then841:
	*state_addr = 76
	goto next_state

if_end842:
	v402 = *lookahead
	cmp843 = v402 == 69
	if cmp843 {
		goto if_then848
	} else {
		goto lor_lhs_false845
	}

lor_lhs_false845:
	v403 = *lookahead
	cmp846 = v403 == 101
	if cmp846 {
		goto if_then848
	} else {
		goto if_end849
	}

if_then848:
	*state_addr = 32
	goto next_state

if_end849:
	v404 = *lookahead
	cmp850 = v404 == 88
	if cmp850 {
		goto if_then855
	} else {
		goto lor_lhs_false852
	}

lor_lhs_false852:
	v405 = *lookahead
	cmp853 = v405 == 120
	if cmp853 {
		goto if_then855
	} else {
		goto if_end856
	}

if_then855:
	*state_addr = 36
	goto next_state

if_end856:
	v406 = *result
	tobool857 = (v406 & 1) != 0
	*retval = tobool857
	goto _return

sw_bb858:
	*result = 1
	v407 = *lexer_addr
	result_symbol859 = &v407.F1
	*result_symbol859 = 10
	v408 = *lexer_addr
	mark_end860 = &v408.F3
	v409 = *mark_end860
	v410 = *lexer_addr
	v409(v410)
	v411 = *lookahead
	cmp861 = v411 == 46
	if cmp861 {
		goto if_then863
	} else {
		goto if_end864
	}

if_then863:
	*state_addr = 76
	goto next_state

if_end864:
	v412 = *lookahead
	cmp865 = v412 == 69
	if cmp865 {
		goto if_then870
	} else {
		goto lor_lhs_false867
	}

lor_lhs_false867:
	v413 = *lookahead
	cmp868 = v413 == 101
	if cmp868 {
		goto if_then870
	} else {
		goto if_end871
	}

if_then870:
	*state_addr = 32
	goto next_state

if_end871:
	v414 = *lookahead
	cmp872 = 48 <= v414
	if cmp872 {
		goto land_lhs_true874
	} else {
		goto if_end878
	}

land_lhs_true874:
	v415 = *lookahead
	cmp875 = v415 <= 57
	if cmp875 {
		goto if_then877
	} else {
		goto if_end878
	}

if_then877:
	*state_addr = 75
	goto next_state

if_end878:
	v416 = *result
	tobool879 = (v416 & 1) != 0
	*retval = tobool879
	goto _return

sw_bb880:
	*result = 1
	v417 = *lexer_addr
	result_symbol881 = &v417.F1
	*result_symbol881 = 10
	v418 = *lexer_addr
	mark_end882 = &v418.F3
	v419 = *mark_end882
	v420 = *lexer_addr
	v419(v420)
	v421 = *lookahead
	cmp883 = v421 == 69
	if cmp883 {
		goto if_then888
	} else {
		goto lor_lhs_false885
	}

lor_lhs_false885:
	v422 = *lookahead
	cmp886 = v422 == 101
	if cmp886 {
		goto if_then888
	} else {
		goto if_end889
	}

if_then888:
	*state_addr = 32
	goto next_state

if_end889:
	v423 = *lookahead
	cmp890 = 48 <= v423
	if cmp890 {
		goto land_lhs_true892
	} else {
		goto if_end896
	}

land_lhs_true892:
	v424 = *lookahead
	cmp893 = v424 <= 57
	if cmp893 {
		goto if_then895
	} else {
		goto if_end896
	}

if_then895:
	*state_addr = 76
	goto next_state

if_end896:
	v425 = *result
	tobool897 = (v425 & 1) != 0
	*retval = tobool897
	goto _return

sw_bb898:
	*result = 1
	v426 = *lexer_addr
	result_symbol899 = &v426.F1
	*result_symbol899 = 10
	v427 = *lexer_addr
	mark_end900 = &v427.F3
	v428 = *mark_end900
	v429 = *lexer_addr
	v428(v429)
	v430 = *lookahead
	cmp901 = 48 <= v430
	if cmp901 {
		goto land_lhs_true903
	} else {
		goto if_end907
	}

land_lhs_true903:
	v431 = *lookahead
	cmp904 = v431 <= 57
	if cmp904 {
		goto if_then906
	} else {
		goto if_end907
	}

if_then906:
	*state_addr = 77
	goto next_state

if_end907:
	v432 = *result
	tobool908 = (v432 & 1) != 0
	*retval = tobool908
	goto _return

sw_bb909:
	*result = 1
	v433 = *lexer_addr
	result_symbol910 = &v433.F1
	*result_symbol910 = 10
	v434 = *lexer_addr
	mark_end911 = &v434.F3
	v435 = *mark_end911
	v436 = *lexer_addr
	v435(v436)
	v437 = *lookahead
	cmp912 = 48 <= v437
	if cmp912 {
		goto land_lhs_true914
	} else {
		goto lor_lhs_false917
	}

land_lhs_true914:
	v438 = *lookahead
	cmp915 = v438 <= 57
	if cmp915 {
		goto if_then929
	} else {
		goto lor_lhs_false917
	}

lor_lhs_false917:
	v439 = *lookahead
	cmp918 = 65 <= v439
	if cmp918 {
		goto land_lhs_true920
	} else {
		goto lor_lhs_false923
	}

land_lhs_true920:
	v440 = *lookahead
	cmp921 = v440 <= 70
	if cmp921 {
		goto if_then929
	} else {
		goto lor_lhs_false923
	}

lor_lhs_false923:
	v441 = *lookahead
	cmp924 = 97 <= v441
	if cmp924 {
		goto land_lhs_true926
	} else {
		goto if_end930
	}

land_lhs_true926:
	v442 = *lookahead
	cmp927 = v442 <= 102
	if cmp927 {
		goto if_then929
	} else {
		goto if_end930
	}

if_then929:
	*state_addr = 78
	goto next_state

if_end930:
	v443 = *result
	tobool931 = (v443 & 1) != 0
	*retval = tobool931
	goto _return

sw_bb932:
	*result = 1
	v444 = *lexer_addr
	result_symbol933 = &v444.F1
	*result_symbol933 = 11
	v445 = *lexer_addr
	mark_end934 = &v445.F3
	v446 = *mark_end934
	v447 = *lexer_addr
	v446(v447)
	v448 = *result
	tobool935 = (v448 & 1) != 0
	*retval = tobool935
	goto _return

sw_bb936:
	*result = 1
	v449 = *lexer_addr
	result_symbol937 = &v449.F1
	*result_symbol937 = 11
	v450 = *lexer_addr
	mark_end938 = &v450.F3
	v451 = *mark_end938
	v452 = *lexer_addr
	v451(v452)
	v453 = *lookahead
	call939 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v453)
	if call939 {
		goto if_then940
	} else {
		goto if_end941
	}

if_then940:
	*state_addr = 69
	goto next_state

if_end941:
	v454 = *result
	tobool942 = (v454 & 1) != 0
	*retval = tobool942
	goto _return

sw_bb943:
	*result = 1
	v455 = *lexer_addr
	result_symbol944 = &v455.F1
	*result_symbol944 = 12
	v456 = *lexer_addr
	mark_end945 = &v456.F3
	v457 = *mark_end945
	v458 = *lexer_addr
	v457(v458)
	v459 = *result
	tobool946 = (v459 & 1) != 0
	*retval = tobool946
	goto _return

sw_bb947:
	*result = 1
	v460 = *lexer_addr
	result_symbol948 = &v460.F1
	*result_symbol948 = 12
	v461 = *lexer_addr
	mark_end949 = &v461.F3
	v462 = *mark_end949
	v463 = *lexer_addr
	v462(v463)
	v464 = *lookahead
	call950 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v464)
	if call950 {
		goto if_then951
	} else {
		goto if_end952
	}

if_then951:
	*state_addr = 69
	goto next_state

if_end952:
	v465 = *result
	tobool953 = (v465 & 1) != 0
	*retval = tobool953
	goto _return

sw_bb954:
	*result = 1
	v466 = *lexer_addr
	result_symbol955 = &v466.F1
	*result_symbol955 = 13
	v467 = *lexer_addr
	mark_end956 = &v467.F3
	v468 = *mark_end956
	v469 = *lexer_addr
	v468(v469)
	v470 = *result
	tobool957 = (v470 & 1) != 0
	*retval = tobool957
	goto _return

sw_bb958:
	*result = 1
	v471 = *lexer_addr
	result_symbol959 = &v471.F1
	*result_symbol959 = 13
	v472 = *lexer_addr
	mark_end960 = &v472.F3
	v473 = *mark_end960
	v474 = *lexer_addr
	v473(v474)
	v475 = *lookahead
	call961 = set_contains(&sym_identifier_character_set_2[int64(0)], 680, v475)
	if call961 {
		goto if_then962
	} else {
		goto if_end963
	}

if_then962:
	*state_addr = 69
	goto next_state

if_end963:
	v476 = *result
	tobool964 = (v476 & 1) != 0
	*retval = tobool964
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v477 = *retval
	return v477
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
