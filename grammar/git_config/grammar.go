package grammar_git_config

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type TSCharacterRange struct {
	F0 int32
	F1 int32
}
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

var sym_integer_character_set_1 [15]TSCharacterRange = [15]TSCharacterRange{TSCharacterRange{48, 57}, TSCharacterRange{69, 69}, TSCharacterRange{71, 71}, TSCharacterRange{75, 75}, TSCharacterRange{77, 77}, TSCharacterRange{80, 80}, TSCharacterRange{84, 84}, TSCharacterRange{89, 90}, TSCharacterRange{101, 101}, TSCharacterRange{103, 103}, TSCharacterRange{107, 107}, TSCharacterRange{109, 109}, TSCharacterRange{112, 112}, TSCharacterRange{116, 116}, TSCharacterRange{121, 122}}
var tree_sitter_git_config_language struct {
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
var ts_small_parse_table [630]int16 = [630]int16{10, 11, 1, 3, 15, 1, 17, 17, 1, 18, 19, 1, 19, 21, 1, 20, 25, 1, 37, 37, 1, 32, 39, 3, 29, 30, 31, 5, 5, 33, 34, 36, 38, 43, 13, 7, 9, 10, 11, 12, 13, 14, 15, 5, 21, 1, 20, 25, 1, 3, 23, 2, 1, 21, 27, 3, 17, 18, 19, 4, 5, 33, 34, 36, 38, 43, 5, 31, 1, 3, 37, 1, 20, 29, 2, 1, 21, 34, 3, 17, 18, 19, 4, 5, 33, 34, 36, 38, 43, 5, 21, 1, 20, 25, 1, 3, 40, 2, 1, 21, 27, 3, 17, 18, 19, 4, 5, 33, 34, 36, 38, 43, 5, 21, 1, 20, 25, 1, 3, 42, 2, 1, 21, 44, 3, 17, 18, 19, 3, 5, 33, 34, 36, 38, 43, 5, 21, 1, 20, 25, 1, 3, 46, 2, 1, 21, 48, 3, 17, 18, 19, 8, 5, 33, 34, 36, 38, 43, 5, 21, 1, 20, 25, 1, 3, 50, 2, 1, 21, 27, 3, 17, 18, 19, 4, 5, 33, 34, 36, 38, 43, 7, 7, 1, 2, 9, 1, 21, 52, 1, 0, 54, 1, 1, 40, 1, 25, 50, 1, 39, 10, 2, 24, 40, 7, 56, 1, 0, 58, 1, 1, 61, 1, 2, 64, 1, 21, 40, 1, 25, 50, 1, 39, 10, 2, 24, 40, 7, 9, 1, 21, 69, 1, 1, 71, 1, 8, 13, 1, 41, 35, 1, 28, 49, 1, 39, 67, 2, 0, 2, 7, 9, 1, 21, 71, 1, 8, 75, 1, 1, 11, 1, 41, 35, 1, 28, 49, 1, 39, 73, 2, 0, 2, 7, 79, 1, 1, 82, 1, 8, 85, 1, 21, 13, 1, 41, 35, 1, 28, 49, 1, 39, 77, 2, 0, 2, 6, 88, 1, 3, 92, 1, 18, 94, 1, 20, 26, 1, 37, 90, 2, 16, 19, 21, 2, 35, 38, 3, 100, 1, 20, 96, 3, 1, 19, 21, 98, 3, 3, 17, 18, 2, 104, 1, 20, 102, 6, 1, 3, 17, 18, 19, 21, 2, 108, 1, 20, 106, 6, 1, 3, 17, 18, 19, 21, 2, 112, 1, 20, 110, 6, 1, 3, 17, 18, 19, 21, 4, 88, 1, 3, 94, 1, 20, 90, 2, 16, 19, 21, 2, 35, 38, 4, 94, 1, 20, 114, 1, 3, 116, 2, 16, 19, 22, 2, 35, 38, 4, 94, 1, 20, 118, 1, 3, 116, 2, 16, 19, 22, 2, 35, 38, 4, 120, 1, 3, 125, 1, 20, 122, 2, 16, 19, 22, 2, 35, 38, 1, 128, 5, 0, 1, 2, 8, 21, 1, 77, 5, 0, 1, 2, 8, 21, 3, 25, 1, 3, 130, 2, 17, 18, 6, 2, 34, 36, 3, 94, 1, 20, 132, 2, 16, 19, 20, 2, 35, 38, 4, 134, 1, 3, 136, 1, 6, 139, 1, 19, 27, 1, 42, 4, 142, 1, 3, 144, 1, 6, 146, 1, 19, 27, 1, 42, 1, 148, 4, 0, 1, 2, 21, 4, 150, 1, 6, 152, 1, 19, 28, 1, 42, 45, 1, 27, 1, 56, 4, 0, 1, 2, 21, 1, 108, 4, 3, 16, 19, 20, 1, 154, 3, 16, 19, 20, 2, 158, 1, 7, 156, 2, 1, 21, 3, 9, 1, 21, 160, 1, 1, 51, 1, 39, 2, 162, 1, 3, 164, 1, 4, 1, 40, 2, 1, 21, 2, 166, 1, 1, 168, 1, 22, 1, 170, 2, 1, 21, 2, 172, 1, 1, 29, 1, 26, 1, 174, 1, 1, 1, 176, 1, 0, 1, 178, 1, 1, 1, 180, 1, 1, 1, 182, 1, 3, 1, 184, 1, 1, 1, 186, 1, 5, 1, 188, 1, 4, 1, 160, 1, 1, 1, 190, 1, 1, 1, 192, 1, 1, 1, 194, 1, 1}
var ts_small_parse_table_map [51]int32 = [51]int32{0, 43, 66, 89, 112, 135, 158, 181, 204, 227, 250, 273, 296, 317, 331, 343, 355, 367, 382, 397, 412, 427, 435, 443, 455, 467, 480, 493, 500, 513, 520, 527, 533, 541, 551, 558, 563, 570, 575, 582, 586, 590, 594, 598, 602, 606, 610, 614, 618, 622, 626}
var ts_symbol_names [44]unsafe.Pointer = [44]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_12), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_30), libc.Ptr(&_str_31), libc.Ptr(&_str_32), libc.Ptr(&_str_33), libc.Ptr(&_str_34), libc.Ptr(&_str_35), libc.Ptr(&_str_36), libc.Ptr(&_str_37), libc.Ptr(&_str_38), libc.Ptr(&_str_39), libc.Ptr(&_str_40), libc.Ptr(&_str_41)}
var ts_field_names [2]unsafe.Pointer = [2]unsafe.Pointer{nil, libc.Ptr(&_str_42)}
var ts_field_map_slices [2]TSMapSlice = [2]TSMapSlice{TSMapSlice{}, TSMapSlice{0, 1}}
var ts_field_map_entries [1]TSFieldMapEntry = [1]TSFieldMapEntry{TSFieldMapEntry{1, 2, 0}}
var ts_symbol_metadata [44]TSSymbolMetadata = [44]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [44]int16 = [44]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 11, 11, 9, 11, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [2][6]int16 = [2][6]int16{}
var ts_lex_modes [53]TSLexMode = [53]TSLexMode{TSLexMode{}, TSLexMode{}, TSLexMode{4, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{18, 0}, TSLexMode{18, 0}, TSLexMode{18, 0}, TSLexMode{5, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{2, 0}, TSLexMode{6, 0}, TSLexMode{6, 0}, TSLexMode{6, 0}, TSLexMode{6, 0}, TSLexMode{18, 0}, TSLexMode{18, 0}, TSLexMode{2, 0}, TSLexMode{6, 0}, TSLexMode{7, 0}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{6, 0}, TSLexMode{6, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{3, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{9, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}}
var ts_primary_state_ids [53]int16 = [53]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 17, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 46}
var ts_parse_table struct {
	F0 struct {
		F0 [22]int16
		F1 [22]int16
	}
	F1 [44]int16
} = struct {
	F0 struct {
		F0 [22]int16
		F1 [22]int16
	}
	F1 [44]int16
}{struct {
	F0 [22]int16
	F1 [22]int16
}{[22]int16{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1}, [22]int16{}}, [44]int16{3, 5, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 42, 9, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 50, 9, 0, 0, 0}}
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
	F28 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
	F53 TSParseActionEntry
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
	F59 TSParseActionEntry
	F60 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F61 struct {
		F0 anon_2
		F1 [6]byte
	}
	F62 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F70 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F71 struct {
		F0 anon_2
		F1 [6]byte
	}
	F72 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F73 struct {
		F0 anon_2
		F1 [6]byte
	}
	F74 TSParseActionEntry
	F75 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F78 TSParseActionEntry
	F79 struct {
		F0 anon_2
		F1 [6]byte
	}
	F80 TSParseActionEntry
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
	F83 TSParseActionEntry
	F84 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F85 struct {
		F0 anon_2
		F1 [6]byte
	}
	F86 TSParseActionEntry
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
	F121 TSParseActionEntry
	F122 struct {
		F0 anon_2
		F1 [6]byte
	}
	F123 TSParseActionEntry
	F124 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F125 struct {
		F0 anon_2
		F1 [6]byte
	}
	F126 TSParseActionEntry
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
	F129 TSParseActionEntry
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
	F135 TSParseActionEntry
	F136 struct {
		F0 anon_2
		F1 [6]byte
	}
	F137 TSParseActionEntry
	F138 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F139 struct {
		F0 anon_2
		F1 [6]byte
	}
	F140 TSParseActionEntry
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
	F143 TSParseActionEntry
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
	F148 struct {
		F0 anon_2
		F1 [6]byte
	}
	F149 TSParseActionEntry
	F150 struct {
		F0 anon_2
		F1 [6]byte
	}
	F151 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F152 struct {
		F0 anon_2
		F1 [6]byte
	}
	F153 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F154 struct {
		F0 anon_2
		F1 [6]byte
	}
	F155 TSParseActionEntry
	F156 struct {
		F0 anon_2
		F1 [6]byte
	}
	F157 TSParseActionEntry
	F158 struct {
		F0 anon_2
		F1 [6]byte
	}
	F159 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F160 struct {
		F0 anon_2
		F1 [6]byte
	}
	F161 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F162 struct {
		F0 anon_2
		F1 [6]byte
	}
	F163 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F164 struct {
		F0 anon_2
		F1 [6]byte
	}
	F165 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F166 struct {
		F0 anon_2
		F1 [6]byte
	}
	F167 TSParseActionEntry
	F168 struct {
		F0 anon_2
		F1 [6]byte
	}
	F169 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F170 struct {
		F0 anon_2
		F1 [6]byte
	}
	F171 TSParseActionEntry
	F172 struct {
		F0 anon_2
		F1 [6]byte
	}
	F173 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F174 struct {
		F0 anon_2
		F1 [6]byte
	}
	F175 TSParseActionEntry
	F176 struct {
		F0 anon_2
		F1 [6]byte
	}
	F177 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F178 struct {
		F0 anon_2
		F1 [6]byte
	}
	F179 TSParseActionEntry
	F180 struct {
		F0 anon_2
		F1 [6]byte
	}
	F181 TSParseActionEntry
	F182 struct {
		F0 anon_2
		F1 [6]byte
	}
	F183 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F184 struct {
		F0 anon_2
		F1 [6]byte
	}
	F185 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F186 struct {
		F0 anon_2
		F1 [6]byte
	}
	F187 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F188 struct {
		F0 anon_2
		F1 [6]byte
	}
	F189 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F190 struct {
		F0 anon_2
		F1 [6]byte
	}
	F191 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F192 struct {
		F0 anon_2
		F1 [6]byte
	}
	F193 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F194 struct {
		F0 anon_2
		F1 [6]byte
	}
	F195 struct {
		F0 struct {
			F0 struct {
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
	F28 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
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
		F0 struct {
			F0 struct {
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
	F53 TSParseActionEntry
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
	F59 TSParseActionEntry
	F60 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F61 struct {
		F0 anon_2
		F1 [6]byte
	}
	F62 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F70 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F71 struct {
		F0 anon_2
		F1 [6]byte
	}
	F72 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F73 struct {
		F0 anon_2
		F1 [6]byte
	}
	F74 TSParseActionEntry
	F75 struct {
		F0 anon_2
		F1 [6]byte
	}
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
	F78 TSParseActionEntry
	F79 struct {
		F0 anon_2
		F1 [6]byte
	}
	F80 TSParseActionEntry
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
	F83 TSParseActionEntry
	F84 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F85 struct {
		F0 anon_2
		F1 [6]byte
	}
	F86 TSParseActionEntry
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
	F121 TSParseActionEntry
	F122 struct {
		F0 anon_2
		F1 [6]byte
	}
	F123 TSParseActionEntry
	F124 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F125 struct {
		F0 anon_2
		F1 [6]byte
	}
	F126 TSParseActionEntry
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
	F129 TSParseActionEntry
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
	F135 TSParseActionEntry
	F136 struct {
		F0 anon_2
		F1 [6]byte
	}
	F137 TSParseActionEntry
	F138 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F139 struct {
		F0 anon_2
		F1 [6]byte
	}
	F140 TSParseActionEntry
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
	F143 TSParseActionEntry
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
	F148 struct {
		F0 anon_2
		F1 [6]byte
	}
	F149 TSParseActionEntry
	F150 struct {
		F0 anon_2
		F1 [6]byte
	}
	F151 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F152 struct {
		F0 anon_2
		F1 [6]byte
	}
	F153 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F154 struct {
		F0 anon_2
		F1 [6]byte
	}
	F155 TSParseActionEntry
	F156 struct {
		F0 anon_2
		F1 [6]byte
	}
	F157 TSParseActionEntry
	F158 struct {
		F0 anon_2
		F1 [6]byte
	}
	F159 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F160 struct {
		F0 anon_2
		F1 [6]byte
	}
	F161 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F162 struct {
		F0 anon_2
		F1 [6]byte
	}
	F163 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F164 struct {
		F0 anon_2
		F1 [6]byte
	}
	F165 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F166 struct {
		F0 anon_2
		F1 [6]byte
	}
	F167 TSParseActionEntry
	F168 struct {
		F0 anon_2
		F1 [6]byte
	}
	F169 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F170 struct {
		F0 anon_2
		F1 [6]byte
	}
	F171 TSParseActionEntry
	F172 struct {
		F0 anon_2
		F1 [6]byte
	}
	F173 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F174 struct {
		F0 anon_2
		F1 [6]byte
	}
	F175 TSParseActionEntry
	F176 struct {
		F0 anon_2
		F1 [6]byte
	}
	F177 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F178 struct {
		F0 anon_2
		F1 [6]byte
	}
	F179 TSParseActionEntry
	F180 struct {
		F0 anon_2
		F1 [6]byte
	}
	F181 TSParseActionEntry
	F182 struct {
		F0 anon_2
		F1 [6]byte
	}
	F183 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F184 struct {
		F0 anon_2
		F1 [6]byte
	}
	F185 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F186 struct {
		F0 anon_2
		F1 [6]byte
	}
	F187 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F188 struct {
		F0 anon_2
		F1 [6]byte
	}
	F189 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F190 struct {
		F0 anon_2
		F1 [6]byte
	}
	F191 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F192 struct {
		F0 anon_2
		F1 [6]byte
	}
	F193 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F194 struct {
		F0 anon_2
		F1 [6]byte
	}
	F195 struct {
		F0 struct {
			F0 struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 23, 0, 0}}}, struct {
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
}{0, 0, 39, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 46, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 32, 0, 0}}}, struct {
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
}{0, 0, 4, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 43, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 43, 0, 0}}}, struct {
	F0 struct {
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 43, 0, 0}}}, struct {
	F0 struct {
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 43, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 46, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 31, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 32, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 32, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 32, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 23, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 40, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 40, 0, 0}}}, struct {
	F0 struct {
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 40, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 47, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 40, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 38, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 26, 0, 0}}}, struct {
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
}{0, 0, 34, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 26, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 41, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 41, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 13, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 41, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 34, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 41, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 38, 0, 1}, [2]byte{}}}, struct {
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
}{0, 0, 33, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 52, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 36, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 37, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 36, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 38, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 38, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 34, 0, 0}}}, struct {
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
}{0, 0, 22, 0, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 35, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 35, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 22, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 35, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 52, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 41, 0, 0}}}, struct {
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
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 42, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 42, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 27, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 42, 0, 0}}}, struct {
	F0 struct {
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
}{0, 0, 27, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 27, 0, 0}}}, struct {
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
}{0, 0, 27, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
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
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 37, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 28, 0, 0}}}, struct {
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
}{0, 0, 44, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 39, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 28, 0, 1}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 6, 25, 0, 0}}}, struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 39, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 25, 0, 0}}}, struct {
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
}{0, 0, 32, 0, 0}, [2]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [14]byte = [14]byte{99, 111, 110, 102, 105, 103, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_4 [2]byte = [2]byte{91, 0}
var _str_5 [2]byte = [2]byte{34, 0}
var _str_6 [2]byte = [2]byte{93, 0}
var _str_7 [13]byte = [13]byte{115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 0}
var _str_8 [23]byte = [23]byte{115, 117, 98, 115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_9 [2]byte = [2]byte{61, 0}
var _str_10 [5]byte = [5]byte{110, 97, 109, 101, 0}
var _str_11 [5]byte = [5]byte{116, 114, 117, 101, 0}
var _str_12 [6]byte = [6]byte{102, 97, 108, 115, 101, 0}
var _str_13 [8]byte = [8]byte{105, 110, 116, 101, 103, 101, 114, 0}
var _str_14 [30]byte = [30]byte{95, 113, 117, 111, 116, 101, 100, 95, 115, 116, 114, 105, 110, 103, 95, 99, 111, 110, 116, 101, 110, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_15 [24]byte = [24]byte{95, 117, 110, 113, 117, 111, 116, 101, 100, 95, 115, 116, 114, 105, 110, 103, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_16 [2]byte = [2]byte{33, 0}
var _str_17 [16]byte = [16]byte{101, 115, 99, 97, 112, 101, 95, 115, 101, 113, 117, 101, 110, 99, 101, 0}
var _str_18 [2]byte = [2]byte{92, 0}
var _str_19 [15]byte = [15]byte{99, 111, 109, 109, 101, 110, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_20 [15]byte = [15]byte{99, 111, 109, 109, 101, 110, 116, 95, 116, 111, 107, 101, 110, 50, 0}
var _str_21 [7]byte = [7]byte{99, 111, 110, 102, 105, 103, 0}
var _str_22 [8]byte = [8]byte{115, 101, 99, 116, 105, 111, 110, 0}
var _str_23 [15]byte = [15]byte{115, 101, 99, 116, 105, 111, 110, 95, 104, 101, 97, 100, 101, 114, 0}
var _str_24 [14]byte = [14]byte{95, 115, 101, 99, 116, 105, 111, 110, 95, 98, 111, 100, 121, 0}
var _str_25 [16]byte = [16]byte{115, 117, 98, 115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 0}
var _str_26 [9]byte = [9]byte{118, 97, 114, 105, 97, 98, 108, 101, 0}
var _str_27 [7]byte = [7]byte{95, 118, 97, 108, 117, 101, 0}
var _str_28 [9]byte = [9]byte{95, 98, 111, 111, 108, 101, 97, 110, 0}
var _str_29 [7]byte = [7]byte{115, 116, 114, 105, 110, 103, 0}
var _str_30 [22]byte = [22]byte{95, 115, 104, 101, 108, 108, 95, 99, 111, 109, 109, 97, 110, 100, 95, 115, 116, 114, 105, 110, 103, 0}
var _str_31 [17]byte = [17]byte{95, 115, 116, 114, 105, 110, 103, 95, 102, 114, 97, 103, 109, 101, 110, 116, 0}
var _str_32 [15]byte = [15]byte{95, 113, 117, 111, 116, 101, 100, 95, 115, 116, 114, 105, 110, 103, 0}
var _str_33 [23]byte = [23]byte{95, 113, 117, 111, 116, 101, 100, 95, 115, 116, 114, 105, 110, 103, 95, 99, 111, 110, 116, 101, 110, 116, 0}
var _str_34 [17]byte = [17]byte{95, 117, 110, 113, 117, 111, 116, 101, 100, 95, 115, 116, 114, 105, 110, 103, 0}
var _str_35 [14]byte = [14]byte{115, 104, 101, 108, 108, 95, 99, 111, 109, 109, 97, 110, 100, 0}
var _str_36 [19]byte = [19]byte{95, 108, 105, 110, 101, 95, 99, 111, 110, 116, 105, 110, 117, 97, 116, 105, 111, 110, 0}
var _str_37 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_38 [15]byte = [15]byte{99, 111, 110, 102, 105, 103, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_39 [22]byte = [22]byte{95, 115, 101, 99, 116, 105, 111, 110, 95, 98, 111, 100, 121, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_40 [24]byte = [24]byte{115, 117, 98, 115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_41 [15]byte = [15]byte{115, 116, 114, 105, 110, 103, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_42 [6]byte = [6]byte{118, 97, 108, 117, 101, 0}
var ts_lex_map [34]int16 = [34]int16{10, 20, 13, 1, 33, 73, 34, 22, 61, 41, 91, 21, 92, 75, 93, 23, 102, 24, 110, 31, 111, 28, 116, 32, 121, 25, 35, 76, 59, 76, 46, 38, 95, 38}
var ts_lex_map_43 [16]int16 = [16]int16{33, 73, 34, 22, 92, 75, 102, 60, 110, 67, 111, 64, 116, 68, 121, 61}
var ts_lex_map_44 [18]int16 = [18]int16{85, 17, 117, 13, 34, 74, 92, 74, 98, 74, 102, 74, 110, 74, 114, 74, 116, 74}
var ts_lex_map_45 [18]int16 = [18]int16{85, 17, 117, 13, 34, 74, 92, 74, 98, 74, 102, 74, 110, 74, 114, 74, 116, 74}

func init() {
	tree_sitter_git_config_language = struct {
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
	}{14, 44, 0, 23, 0, 53, 2, 2, 1, 6, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), libc.Ptr(&ts_field_names), libc.Ptr(&ts_field_map_slices), libc.Ptr(&ts_field_map_entries), libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), nil, nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{}, [5]byte{}}
}
func tree_sitter_git_config() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_git_config_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, cmp29, cmp32, cmp35, cmp38, loadedv42, cmp44, loadedv48, cmp50, cmp54, cmp58, cmp62, cmp66, cmp70, cmp73, cmp77, cmp80, cmp83, cmp87, loadedv91, cmp93, cmp97, cmp101, cmp104, cmp107, cmp111, loadedv115, cmp120, cmp126, cmp136, cmp139, cmp142, cmp145, cmp149, cmp152, cmp156, cmp159, cmp162, cmp165, cmp168, cmp171, loadedv175, cmp177, cmp181, cmp185, cmp189, cmp192, cmp195, cmp198, cmp202, cmp205, cmp208, loadedv212, cmp214, cmp218, cmp222, cmp225, cmp228, cmp231, cmp235, cmp238, cmp241, loadedv245, cmp247, cmp251, cmp255, cmp258, cmp261, cmp264, cmp268, cmp271, cmp274, loadedv278, cmp283, cmp289, loadedv299, cmp301, cmp304, cmp307, cmp310, cmp314, cmp317, cmp320, cmp323, cmp326, cmp329, cmp332, cmp335, loadedv339, cmp341, cmp344, cmp347, cmp350, cmp353, cmp356, loadedv360, cmp362, cmp365, cmp368, cmp371, cmp374, cmp377, loadedv381, cmp383, cmp386, cmp389, cmp392, cmp395, cmp398, loadedv402, cmp404, cmp407, cmp410, cmp413, cmp416, cmp419, loadedv423, cmp425, cmp428, cmp431, cmp434, cmp437, cmp440, loadedv444, cmp446, cmp449, cmp452, cmp455, cmp458, cmp461, loadedv465, cmp467, cmp470, cmp473, cmp476, cmp479, cmp482, loadedv486, cmp488, cmp491, cmp494, cmp497, cmp500, cmp503, loadedv507, loadedv509, cmp512, cmp516, cmp520, cmp524, cmp527, cmp531, cmp534, cmp537, cmp541, cmp544, cmp547, cmp550, loadedv554, loadedv556, loadedv560, loadedv564, loadedv568, loadedv572, cmp576, cmp580, cmp584, cmp588, cmp591, cmp594, cmp597, cmp600, cmp603, cmp606, loadedv610, cmp614, cmp618, cmp622, cmp626, cmp629, cmp632, cmp635, cmp638, cmp641, cmp644, loadedv648, cmp652, cmp656, cmp660, cmp664, cmp667, cmp670, cmp673, cmp676, cmp679, cmp682, loadedv686, cmp690, cmp694, cmp698, cmp702, cmp705, cmp708, cmp711, cmp714, cmp717, cmp720, loadedv724, cmp728, cmp732, cmp736, cmp740, cmp744, cmp747, cmp750, cmp753, cmp756, cmp759, cmp762, loadedv766, cmp770, cmp774, cmp778, cmp782, cmp785, cmp788, cmp791, cmp794, cmp797, cmp800, loadedv804, cmp808, cmp812, cmp816, cmp820, cmp823, cmp826, cmp829, cmp832, cmp835, cmp838, loadedv842, cmp846, cmp850, cmp854, cmp858, cmp861, cmp864, cmp867, cmp870, cmp873, cmp876, loadedv880, cmp884, cmp888, cmp892, cmp896, cmp899, cmp902, cmp905, cmp908, cmp911, cmp914, loadedv918, cmp922, cmp926, cmp930, cmp934, cmp937, cmp940, cmp943, cmp946, cmp949, cmp952, loadedv956, cmp960, cmp964, cmp968, cmp972, cmp975, cmp978, cmp981, cmp984, cmp987, cmp990, loadedv994, cmp998, cmp1002, cmp1006, cmp1010, cmp1013, cmp1016, cmp1019, cmp1022, cmp1025, cmp1028, loadedv1032, cmp1036, cmp1040, cmp1044, cmp1047, cmp1050, cmp1053, cmp1056, cmp1059, cmp1062, loadedv1066, cmp1070, cmp1073, call1077, cmp1080, cmp1083, cmp1086, cmp1089, cmp1092, cmp1095, loadedv1099, cmp1103, cmp1106, cmp1109, cmp1112, cmp1115, cmp1118, cmp1121, cmp1124, loadedv1128, cmp1132, cmp1135, cmp1138, cmp1141, cmp1145, cmp1148, cmp1151, cmp1154, cmp1157, loadedv1161, cmp1165, cmp1168, cmp1171, cmp1174, cmp1177, loadedv1181, loadedv1185, cmp1189, cmp1192, cmp1195, cmp1198, cmp1201, cmp1204, cmp1207, cmp1210, loadedv1214, cmp1218, cmp1222, cmp1226, cmp1229, cmp1232, cmp1235, cmp1238, cmp1241, cmp1244, loadedv1248, cmp1252, cmp1255, cmp1258, cmp1261, cmp1264, cmp1267, cmp1270, loadedv1274, cmp1278, cmp1282, cmp1286, cmp1289, cmp1292, cmp1295, cmp1298, cmp1301, cmp1304, loadedv1308, cmp1312, cmp1315, cmp1318, cmp1321, cmp1324, cmp1327, cmp1330, loadedv1334, cmp1338, cmp1342, cmp1346, cmp1349, cmp1352, cmp1355, cmp1358, cmp1361, cmp1364, loadedv1368, cmp1372, cmp1375, cmp1378, cmp1381, cmp1384, cmp1387, cmp1390, loadedv1394, cmp1398, cmp1402, cmp1406, cmp1409, cmp1412, cmp1415, cmp1418, cmp1421, cmp1424, loadedv1428, cmp1432, cmp1435, cmp1438, cmp1441, cmp1444, cmp1447, cmp1450, loadedv1454, cmp1458, cmp1462, cmp1466, cmp1469, cmp1472, cmp1475, cmp1478, cmp1481, cmp1484, loadedv1488, cmp1492, cmp1495, cmp1498, cmp1501, cmp1504, cmp1507, cmp1510, loadedv1514, cmp1518, cmp1522, cmp1526, cmp1529, cmp1532, cmp1535, cmp1538, cmp1541, cmp1544, loadedv1548, cmp1552, cmp1555, cmp1558, cmp1561, cmp1564, cmp1567, cmp1570, loadedv1574, cmp1578, cmp1581, call1585, cmp1588, cmp1591, cmp1594, cmp1597, cmp1600, cmp1603, cmp1606, loadedv1610, cmp1614, cmp1617, cmp1620, cmp1623, cmp1626, cmp1629, cmp1632, loadedv1636, loadedv1640, cmp1644, cmp1648, cmp1652, cmp1655, cmp1658, cmp1661, cmp1665, cmp1668, cmp1671, cmp1674, cmp1677, loadedv1681, cmp1685, cmp1689, cmp1692, cmp1695, cmp1698, cmp1702, cmp1705, cmp1708, cmp1711, loadedv1715, cmp1719, cmp1723, cmp1726, cmp1729, cmp1732, cmp1735, cmp1738, cmp1741, loadedv1745, cmp1749, cmp1753, cmp1756, cmp1759, cmp1762, cmp1765, cmp1768, cmp1771, loadedv1775, cmp1779, cmp1783, cmp1786, cmp1789, cmp1792, cmp1795, cmp1798, cmp1801, loadedv1805, cmp1809, cmp1813, cmp1816, cmp1819, cmp1822, cmp1825, cmp1828, cmp1831, loadedv1835, cmp1839, cmp1843, cmp1847, cmp1850, cmp1853, cmp1856, cmp1859, cmp1862, cmp1865, loadedv1869, cmp1873, cmp1877, cmp1880, cmp1883, cmp1886, cmp1889, cmp1892, cmp1895, loadedv1899, cmp1903, cmp1907, cmp1910, cmp1913, cmp1916, cmp1919, cmp1922, cmp1925, loadedv1929, cmp1933, cmp1937, cmp1940, cmp1943, cmp1946, cmp1949, cmp1952, cmp1955, loadedv1959, cmp1963, cmp1967, cmp1970, cmp1973, cmp1976, cmp1979, cmp1982, cmp1985, loadedv1989, cmp1993, cmp1997, cmp2000, cmp2003, cmp2006, cmp2009, cmp2012, cmp2015, loadedv2019, cmp2023, cmp2027, cmp2030, cmp2033, cmp2036, cmp2039, cmp2042, cmp2045, loadedv2049, cmp2053, cmp2057, cmp2060, cmp2063, cmp2066, cmp2069, cmp2072, cmp2075, loadedv2079, cmp2083, cmp2086, cmp2089, cmp2092, cmp2095, cmp2098, cmp2101, loadedv2105, loadedv2109, loadedv2113, cmp2120, cmp2126, loadedv2136, loadedv2140, cmp2144, cmp2147, cmp2150, cmp2153, cmp2157, cmp2160, cmp2163, loadedv2167, cmp2171, cmp2174, cmp2177, loadedv2181, v914 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v51, v54, v102, v105, v883, v886 int16
	var state_addr, arrayidx, arrayidx11, arrayidx124, arrayidx131, arrayidx287, arrayidx294, result_symbol, result_symbol558, result_symbol562, result_symbol566, result_symbol570, result_symbol574, result_symbol612, result_symbol650, result_symbol688, result_symbol726, result_symbol768, result_symbol806, result_symbol844, result_symbol882, result_symbol920, result_symbol958, result_symbol996, result_symbol1034, result_symbol1068, result_symbol1101, result_symbol1130, result_symbol1163, result_symbol1183, result_symbol1187, result_symbol1216, result_symbol1250, result_symbol1276, result_symbol1310, result_symbol1336, result_symbol1370, result_symbol1396, result_symbol1430, result_symbol1456, result_symbol1490, result_symbol1516, result_symbol1550, result_symbol1576, result_symbol1612, result_symbol1638, result_symbol1642, result_symbol1683, result_symbol1717, result_symbol1747, result_symbol1777, result_symbol1807, result_symbol1837, result_symbol1871, result_symbol1901, result_symbol1931, result_symbol1961, result_symbol1991, result_symbol2021, result_symbol2051, result_symbol2081, result_symbol2107, result_symbol2111, result_symbol2115, arrayidx2124, arrayidx2131, result_symbol2138, result_symbol2142, result_symbol2169 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v24, v25, v26, v28, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v42, v43, v44, v45, v46, v47, v49, v50, conv125, v52, v53, add129, v55, add134, v56, v57, v58, v59, v60, v61, v62, v63, v64, v65, v66, v67, v69, v70, v71, v72, v73, v74, v75, v76, v77, v78, v80, v81, v82, v83, v84, v85, v86, v87, v88, v90, v91, v92, v93, v94, v95, v96, v97, v98, v100, v101, conv288, v103, v104, add292, v106, add297, v108, v109, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v121, v122, v123, v124, v125, v126, v128, v129, v130, v131, v132, v133, v135, v136, v137, v138, v139, v140, v142, v143, v144, v145, v146, v147, v149, v150, v151, v152, v153, v154, v156, v157, v158, v159, v160, v161, v163, v164, v165, v166, v167, v168, v170, v171, v172, v173, v174, v175, v178, v179, v180, v181, v182, v183, v184, v185, v186, v187, v188, v189, v220, v221, v222, v223, v224, v225, v226, v227, v228, v229, v235, v236, v237, v238, v239, v240, v241, v242, v243, v244, v250, v251, v252, v253, v254, v255, v256, v257, v258, v259, v265, v266, v267, v268, v269, v270, v271, v272, v273, v274, v280, v281, v282, v283, v284, v285, v286, v287, v288, v289, v290, v296, v297, v298, v299, v300, v301, v302, v303, v304, v305, v311, v312, v313, v314, v315, v316, v317, v318, v319, v320, v326, v327, v328, v329, v330, v331, v332, v333, v334, v335, v341, v342, v343, v344, v345, v346, v347, v348, v349, v350, v356, v357, v358, v359, v360, v361, v362, v363, v364, v365, v371, v372, v373, v374, v375, v376, v377, v378, v379, v380, v386, v387, v388, v389, v390, v391, v392, v393, v394, v395, v401, v402, v403, v404, v405, v406, v407, v408, v409, v415, v416, v417, v418, v419, v420, v421, v422, v423, v429, v430, v431, v432, v433, v434, v435, v436, v442, v443, v444, v445, v446, v447, v448, v449, v450, v456, v457, v458, v459, v460, v471, v472, v473, v474, v475, v476, v477, v478, v484, v485, v486, v487, v488, v489, v490, v491, v492, v498, v499, v500, v501, v502, v503, v504, v510, v511, v512, v513, v514, v515, v516, v517, v518, v524, v525, v526, v527, v528, v529, v530, v536, v537, v538, v539, v540, v541, v542, v543, v544, v550, v551, v552, v553, v554, v555, v556, v562, v563, v564, v565, v566, v567, v568, v569, v570, v576, v577, v578, v579, v580, v581, v582, v588, v589, v590, v591, v592, v593, v594, v595, v596, v602, v603, v604, v605, v606, v607, v608, v614, v615, v616, v617, v618, v619, v620, v621, v622, v628, v629, v630, v631, v632, v633, v634, v640, v641, v642, v643, v644, v645, v646, v647, v648, v649, v655, v656, v657, v658, v659, v660, v661, v672, v673, v674, v675, v676, v677, v678, v679, v680, v681, v682, v688, v689, v690, v691, v692, v693, v694, v695, v696, v702, v703, v704, v705, v706, v707, v708, v709, v715, v716, v717, v718, v719, v720, v721, v722, v728, v729, v730, v731, v732, v733, v734, v735, v741, v742, v743, v744, v745, v746, v747, v748, v754, v755, v756, v757, v758, v759, v760, v761, v762, v768, v769, v770, v771, v772, v773, v774, v775, v781, v782, v783, v784, v785, v786, v787, v788, v794, v795, v796, v797, v798, v799, v800, v801, v807, v808, v809, v810, v811, v812, v813, v814, v820, v821, v822, v823, v824, v825, v826, v827, v833, v834, v835, v836, v837, v838, v839, v840, v846, v847, v848, v849, v850, v851, v852, v853, v859, v860, v861, v862, v863, v864, v865, v881, v882, conv2125, v884, v885, add2129, v887, add2134, v898, v899, v900, v901, v902, v903, v904, v910, v911, v912 int32
	var lookahead, i, i117, i280, i2117, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv119, idxprom123, idxprom130, conv282, idxprom286, idxprom293, conv2119, idxprom2123, idxprom2130 int64
	var v3, storedv, v10, v27, v29, v41, v48, v68, v79, v89, v99, v107, v120, v127, v134, v141, v148, v155, v162, v169, v176, v177, v190, v195, v200, v205, v210, v215, v230, v245, v260, v275, v291, v306, v321, v336, v351, v366, v381, v396, v410, v424, v437, v451, v461, v466, v479, v493, v505, v519, v531, v545, v557, v571, v583, v597, v609, v623, v635, v650, v662, v667, v683, v697, v710, v723, v736, v749, v763, v776, v789, v802, v815, v828, v841, v854, v866, v871, v876, v888, v893, v905, v913 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v191, v192, v193, v194, v196, v197, v198, v199, v201, v202, v203, v204, v206, v207, v208, v209, v211, v212, v213, v214, v216, v217, v218, v219, v231, v232, v233, v234, v246, v247, v248, v249, v261, v262, v263, v264, v276, v277, v278, v279, v292, v293, v294, v295, v307, v308, v309, v310, v322, v323, v324, v325, v337, v338, v339, v340, v352, v353, v354, v355, v367, v368, v369, v370, v382, v383, v384, v385, v397, v398, v399, v400, v411, v412, v413, v414, v425, v426, v427, v428, v438, v439, v440, v441, v452, v453, v454, v455, v462, v463, v464, v465, v467, v468, v469, v470, v480, v481, v482, v483, v494, v495, v496, v497, v506, v507, v508, v509, v520, v521, v522, v523, v532, v533, v534, v535, v546, v547, v548, v549, v558, v559, v560, v561, v572, v573, v574, v575, v584, v585, v586, v587, v598, v599, v600, v601, v610, v611, v612, v613, v624, v625, v626, v627, v636, v637, v638, v639, v651, v652, v653, v654, v663, v664, v665, v666, v668, v669, v670, v671, v684, v685, v686, v687, v698, v699, v700, v701, v711, v712, v713, v714, v724, v725, v726, v727, v737, v738, v739, v740, v750, v751, v752, v753, v764, v765, v766, v767, v777, v778, v779, v780, v790, v791, v792, v793, v803, v804, v805, v806, v816, v817, v818, v819, v829, v830, v831, v832, v842, v843, v844, v845, v855, v856, v857, v858, v867, v868, v869, v870, v872, v873, v874, v875, v877, v878, v879, v880, v889, v890, v891, v892, v894, v895, v896, v897, v906, v907, v908, v909 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end559, mark_end563, mark_end567, mark_end571, mark_end575, mark_end613, mark_end651, mark_end689, mark_end727, mark_end769, mark_end807, mark_end845, mark_end883, mark_end921, mark_end959, mark_end997, mark_end1035, mark_end1069, mark_end1102, mark_end1131, mark_end1164, mark_end1184, mark_end1188, mark_end1217, mark_end1251, mark_end1277, mark_end1311, mark_end1337, mark_end1371, mark_end1397, mark_end1431, mark_end1457, mark_end1491, mark_end1517, mark_end1551, mark_end1577, mark_end1613, mark_end1639, mark_end1643, mark_end1684, mark_end1718, mark_end1748, mark_end1778, mark_end1808, mark_end1838, mark_end1872, mark_end1902, mark_end1932, mark_end1962, mark_end1992, mark_end2022, mark_end2052, mark_end2082, mark_end2108, mark_end2112, mark_end2116, mark_end2139, mark_end2143, mark_end2170 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i117, i280, i2117, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, cmp29, v24, cmp32, v25, cmp35, v26, cmp38, v27, loadedv42, v28, cmp44, v29, loadedv48, v30, cmp50, v31, cmp54, v32, cmp58, v33, cmp62, v34, cmp66, v35, cmp70, v36, cmp73, v37, cmp77, v38, cmp80, v39, cmp83, v40, cmp87, v41, loadedv91, v42, cmp93, v43, cmp97, v44, cmp101, v45, cmp104, v46, cmp107, v47, cmp111, v48, loadedv115, v49, conv119, cmp120, v50, idxprom123, arrayidx124, v51, conv125, v52, cmp126, v53, add129, idxprom130, arrayidx131, v54, v55, add134, v56, cmp136, v57, cmp139, v58, cmp142, v59, cmp145, v60, cmp149, v61, cmp152, v62, cmp156, v63, cmp159, v64, cmp162, v65, cmp165, v66, cmp168, v67, cmp171, v68, loadedv175, v69, cmp177, v70, cmp181, v71, cmp185, v72, cmp189, v73, cmp192, v74, cmp195, v75, cmp198, v76, cmp202, v77, cmp205, v78, cmp208, v79, loadedv212, v80, cmp214, v81, cmp218, v82, cmp222, v83, cmp225, v84, cmp228, v85, cmp231, v86, cmp235, v87, cmp238, v88, cmp241, v89, loadedv245, v90, cmp247, v91, cmp251, v92, cmp255, v93, cmp258, v94, cmp261, v95, cmp264, v96, cmp268, v97, cmp271, v98, cmp274, v99, loadedv278, v100, conv282, cmp283, v101, idxprom286, arrayidx287, v102, conv288, v103, cmp289, v104, add292, idxprom293, arrayidx294, v105, v106, add297, v107, loadedv299, v108, cmp301, v109, cmp304, v110, cmp307, v111, cmp310, v112, cmp314, v113, cmp317, v114, cmp320, v115, cmp323, v116, cmp326, v117, cmp329, v118, cmp332, v119, cmp335, v120, loadedv339, v121, cmp341, v122, cmp344, v123, cmp347, v124, cmp350, v125, cmp353, v126, cmp356, v127, loadedv360, v128, cmp362, v129, cmp365, v130, cmp368, v131, cmp371, v132, cmp374, v133, cmp377, v134, loadedv381, v135, cmp383, v136, cmp386, v137, cmp389, v138, cmp392, v139, cmp395, v140, cmp398, v141, loadedv402, v142, cmp404, v143, cmp407, v144, cmp410, v145, cmp413, v146, cmp416, v147, cmp419, v148, loadedv423, v149, cmp425, v150, cmp428, v151, cmp431, v152, cmp434, v153, cmp437, v154, cmp440, v155, loadedv444, v156, cmp446, v157, cmp449, v158, cmp452, v159, cmp455, v160, cmp458, v161, cmp461, v162, loadedv465, v163, cmp467, v164, cmp470, v165, cmp473, v166, cmp476, v167, cmp479, v168, cmp482, v169, loadedv486, v170, cmp488, v171, cmp491, v172, cmp494, v173, cmp497, v174, cmp500, v175, cmp503, v176, loadedv507, v177, loadedv509, v178, cmp512, v179, cmp516, v180, cmp520, v181, cmp524, v182, cmp527, v183, cmp531, v184, cmp534, v185, cmp537, v186, cmp541, v187, cmp544, v188, cmp547, v189, cmp550, v190, loadedv554, v191, result_symbol, v192, mark_end, v193, v194, v195, loadedv556, v196, result_symbol558, v197, mark_end559, v198, v199, v200, loadedv560, v201, result_symbol562, v202, mark_end563, v203, v204, v205, loadedv564, v206, result_symbol566, v207, mark_end567, v208, v209, v210, loadedv568, v211, result_symbol570, v212, mark_end571, v213, v214, v215, loadedv572, v216, result_symbol574, v217, mark_end575, v218, v219, v220, cmp576, v221, cmp580, v222, cmp584, v223, cmp588, v224, cmp591, v225, cmp594, v226, cmp597, v227, cmp600, v228, cmp603, v229, cmp606, v230, loadedv610, v231, result_symbol612, v232, mark_end613, v233, v234, v235, cmp614, v236, cmp618, v237, cmp622, v238, cmp626, v239, cmp629, v240, cmp632, v241, cmp635, v242, cmp638, v243, cmp641, v244, cmp644, v245, loadedv648, v246, result_symbol650, v247, mark_end651, v248, v249, v250, cmp652, v251, cmp656, v252, cmp660, v253, cmp664, v254, cmp667, v255, cmp670, v256, cmp673, v257, cmp676, v258, cmp679, v259, cmp682, v260, loadedv686, v261, result_symbol688, v262, mark_end689, v263, v264, v265, cmp690, v266, cmp694, v267, cmp698, v268, cmp702, v269, cmp705, v270, cmp708, v271, cmp711, v272, cmp714, v273, cmp717, v274, cmp720, v275, loadedv724, v276, result_symbol726, v277, mark_end727, v278, v279, v280, cmp728, v281, cmp732, v282, cmp736, v283, cmp740, v284, cmp744, v285, cmp747, v286, cmp750, v287, cmp753, v288, cmp756, v289, cmp759, v290, cmp762, v291, loadedv766, v292, result_symbol768, v293, mark_end769, v294, v295, v296, cmp770, v297, cmp774, v298, cmp778, v299, cmp782, v300, cmp785, v301, cmp788, v302, cmp791, v303, cmp794, v304, cmp797, v305, cmp800, v306, loadedv804, v307, result_symbol806, v308, mark_end807, v309, v310, v311, cmp808, v312, cmp812, v313, cmp816, v314, cmp820, v315, cmp823, v316, cmp826, v317, cmp829, v318, cmp832, v319, cmp835, v320, cmp838, v321, loadedv842, v322, result_symbol844, v323, mark_end845, v324, v325, v326, cmp846, v327, cmp850, v328, cmp854, v329, cmp858, v330, cmp861, v331, cmp864, v332, cmp867, v333, cmp870, v334, cmp873, v335, cmp876, v336, loadedv880, v337, result_symbol882, v338, mark_end883, v339, v340, v341, cmp884, v342, cmp888, v343, cmp892, v344, cmp896, v345, cmp899, v346, cmp902, v347, cmp905, v348, cmp908, v349, cmp911, v350, cmp914, v351, loadedv918, v352, result_symbol920, v353, mark_end921, v354, v355, v356, cmp922, v357, cmp926, v358, cmp930, v359, cmp934, v360, cmp937, v361, cmp940, v362, cmp943, v363, cmp946, v364, cmp949, v365, cmp952, v366, loadedv956, v367, result_symbol958, v368, mark_end959, v369, v370, v371, cmp960, v372, cmp964, v373, cmp968, v374, cmp972, v375, cmp975, v376, cmp978, v377, cmp981, v378, cmp984, v379, cmp987, v380, cmp990, v381, loadedv994, v382, result_symbol996, v383, mark_end997, v384, v385, v386, cmp998, v387, cmp1002, v388, cmp1006, v389, cmp1010, v390, cmp1013, v391, cmp1016, v392, cmp1019, v393, cmp1022, v394, cmp1025, v395, cmp1028, v396, loadedv1032, v397, result_symbol1034, v398, mark_end1035, v399, v400, v401, cmp1036, v402, cmp1040, v403, cmp1044, v404, cmp1047, v405, cmp1050, v406, cmp1053, v407, cmp1056, v408, cmp1059, v409, cmp1062, v410, loadedv1066, v411, result_symbol1068, v412, mark_end1069, v413, v414, v415, cmp1070, v416, cmp1073, v417, call1077, v418, cmp1080, v419, cmp1083, v420, cmp1086, v421, cmp1089, v422, cmp1092, v423, cmp1095, v424, loadedv1099, v425, result_symbol1101, v426, mark_end1102, v427, v428, v429, cmp1103, v430, cmp1106, v431, cmp1109, v432, cmp1112, v433, cmp1115, v434, cmp1118, v435, cmp1121, v436, cmp1124, v437, loadedv1128, v438, result_symbol1130, v439, mark_end1131, v440, v441, v442, cmp1132, v443, cmp1135, v444, cmp1138, v445, cmp1141, v446, cmp1145, v447, cmp1148, v448, cmp1151, v449, cmp1154, v450, cmp1157, v451, loadedv1161, v452, result_symbol1163, v453, mark_end1164, v454, v455, v456, cmp1165, v457, cmp1168, v458, cmp1171, v459, cmp1174, v460, cmp1177, v461, loadedv1181, v462, result_symbol1183, v463, mark_end1184, v464, v465, v466, loadedv1185, v467, result_symbol1187, v468, mark_end1188, v469, v470, v471, cmp1189, v472, cmp1192, v473, cmp1195, v474, cmp1198, v475, cmp1201, v476, cmp1204, v477, cmp1207, v478, cmp1210, v479, loadedv1214, v480, result_symbol1216, v481, mark_end1217, v482, v483, v484, cmp1218, v485, cmp1222, v486, cmp1226, v487, cmp1229, v488, cmp1232, v489, cmp1235, v490, cmp1238, v491, cmp1241, v492, cmp1244, v493, loadedv1248, v494, result_symbol1250, v495, mark_end1251, v496, v497, v498, cmp1252, v499, cmp1255, v500, cmp1258, v501, cmp1261, v502, cmp1264, v503, cmp1267, v504, cmp1270, v505, loadedv1274, v506, result_symbol1276, v507, mark_end1277, v508, v509, v510, cmp1278, v511, cmp1282, v512, cmp1286, v513, cmp1289, v514, cmp1292, v515, cmp1295, v516, cmp1298, v517, cmp1301, v518, cmp1304, v519, loadedv1308, v520, result_symbol1310, v521, mark_end1311, v522, v523, v524, cmp1312, v525, cmp1315, v526, cmp1318, v527, cmp1321, v528, cmp1324, v529, cmp1327, v530, cmp1330, v531, loadedv1334, v532, result_symbol1336, v533, mark_end1337, v534, v535, v536, cmp1338, v537, cmp1342, v538, cmp1346, v539, cmp1349, v540, cmp1352, v541, cmp1355, v542, cmp1358, v543, cmp1361, v544, cmp1364, v545, loadedv1368, v546, result_symbol1370, v547, mark_end1371, v548, v549, v550, cmp1372, v551, cmp1375, v552, cmp1378, v553, cmp1381, v554, cmp1384, v555, cmp1387, v556, cmp1390, v557, loadedv1394, v558, result_symbol1396, v559, mark_end1397, v560, v561, v562, cmp1398, v563, cmp1402, v564, cmp1406, v565, cmp1409, v566, cmp1412, v567, cmp1415, v568, cmp1418, v569, cmp1421, v570, cmp1424, v571, loadedv1428, v572, result_symbol1430, v573, mark_end1431, v574, v575, v576, cmp1432, v577, cmp1435, v578, cmp1438, v579, cmp1441, v580, cmp1444, v581, cmp1447, v582, cmp1450, v583, loadedv1454, v584, result_symbol1456, v585, mark_end1457, v586, v587, v588, cmp1458, v589, cmp1462, v590, cmp1466, v591, cmp1469, v592, cmp1472, v593, cmp1475, v594, cmp1478, v595, cmp1481, v596, cmp1484, v597, loadedv1488, v598, result_symbol1490, v599, mark_end1491, v600, v601, v602, cmp1492, v603, cmp1495, v604, cmp1498, v605, cmp1501, v606, cmp1504, v607, cmp1507, v608, cmp1510, v609, loadedv1514, v610, result_symbol1516, v611, mark_end1517, v612, v613, v614, cmp1518, v615, cmp1522, v616, cmp1526, v617, cmp1529, v618, cmp1532, v619, cmp1535, v620, cmp1538, v621, cmp1541, v622, cmp1544, v623, loadedv1548, v624, result_symbol1550, v625, mark_end1551, v626, v627, v628, cmp1552, v629, cmp1555, v630, cmp1558, v631, cmp1561, v632, cmp1564, v633, cmp1567, v634, cmp1570, v635, loadedv1574, v636, result_symbol1576, v637, mark_end1577, v638, v639, v640, cmp1578, v641, cmp1581, v642, call1585, v643, cmp1588, v644, cmp1591, v645, cmp1594, v646, cmp1597, v647, cmp1600, v648, cmp1603, v649, cmp1606, v650, loadedv1610, v651, result_symbol1612, v652, mark_end1613, v653, v654, v655, cmp1614, v656, cmp1617, v657, cmp1620, v658, cmp1623, v659, cmp1626, v660, cmp1629, v661, cmp1632, v662, loadedv1636, v663, result_symbol1638, v664, mark_end1639, v665, v666, v667, loadedv1640, v668, result_symbol1642, v669, mark_end1643, v670, v671, v672, cmp1644, v673, cmp1648, v674, cmp1652, v675, cmp1655, v676, cmp1658, v677, cmp1661, v678, cmp1665, v679, cmp1668, v680, cmp1671, v681, cmp1674, v682, cmp1677, v683, loadedv1681, v684, result_symbol1683, v685, mark_end1684, v686, v687, v688, cmp1685, v689, cmp1689, v690, cmp1692, v691, cmp1695, v692, cmp1698, v693, cmp1702, v694, cmp1705, v695, cmp1708, v696, cmp1711, v697, loadedv1715, v698, result_symbol1717, v699, mark_end1718, v700, v701, v702, cmp1719, v703, cmp1723, v704, cmp1726, v705, cmp1729, v706, cmp1732, v707, cmp1735, v708, cmp1738, v709, cmp1741, v710, loadedv1745, v711, result_symbol1747, v712, mark_end1748, v713, v714, v715, cmp1749, v716, cmp1753, v717, cmp1756, v718, cmp1759, v719, cmp1762, v720, cmp1765, v721, cmp1768, v722, cmp1771, v723, loadedv1775, v724, result_symbol1777, v725, mark_end1778, v726, v727, v728, cmp1779, v729, cmp1783, v730, cmp1786, v731, cmp1789, v732, cmp1792, v733, cmp1795, v734, cmp1798, v735, cmp1801, v736, loadedv1805, v737, result_symbol1807, v738, mark_end1808, v739, v740, v741, cmp1809, v742, cmp1813, v743, cmp1816, v744, cmp1819, v745, cmp1822, v746, cmp1825, v747, cmp1828, v748, cmp1831, v749, loadedv1835, v750, result_symbol1837, v751, mark_end1838, v752, v753, v754, cmp1839, v755, cmp1843, v756, cmp1847, v757, cmp1850, v758, cmp1853, v759, cmp1856, v760, cmp1859, v761, cmp1862, v762, cmp1865, v763, loadedv1869, v764, result_symbol1871, v765, mark_end1872, v766, v767, v768, cmp1873, v769, cmp1877, v770, cmp1880, v771, cmp1883, v772, cmp1886, v773, cmp1889, v774, cmp1892, v775, cmp1895, v776, loadedv1899, v777, result_symbol1901, v778, mark_end1902, v779, v780, v781, cmp1903, v782, cmp1907, v783, cmp1910, v784, cmp1913, v785, cmp1916, v786, cmp1919, v787, cmp1922, v788, cmp1925, v789, loadedv1929, v790, result_symbol1931, v791, mark_end1932, v792, v793, v794, cmp1933, v795, cmp1937, v796, cmp1940, v797, cmp1943, v798, cmp1946, v799, cmp1949, v800, cmp1952, v801, cmp1955, v802, loadedv1959, v803, result_symbol1961, v804, mark_end1962, v805, v806, v807, cmp1963, v808, cmp1967, v809, cmp1970, v810, cmp1973, v811, cmp1976, v812, cmp1979, v813, cmp1982, v814, cmp1985, v815, loadedv1989, v816, result_symbol1991, v817, mark_end1992, v818, v819, v820, cmp1993, v821, cmp1997, v822, cmp2000, v823, cmp2003, v824, cmp2006, v825, cmp2009, v826, cmp2012, v827, cmp2015, v828, loadedv2019, v829, result_symbol2021, v830, mark_end2022, v831, v832, v833, cmp2023, v834, cmp2027, v835, cmp2030, v836, cmp2033, v837, cmp2036, v838, cmp2039, v839, cmp2042, v840, cmp2045, v841, loadedv2049, v842, result_symbol2051, v843, mark_end2052, v844, v845, v846, cmp2053, v847, cmp2057, v848, cmp2060, v849, cmp2063, v850, cmp2066, v851, cmp2069, v852, cmp2072, v853, cmp2075, v854, loadedv2079, v855, result_symbol2081, v856, mark_end2082, v857, v858, v859, cmp2083, v860, cmp2086, v861, cmp2089, v862, cmp2092, v863, cmp2095, v864, cmp2098, v865, cmp2101, v866, loadedv2105, v867, result_symbol2107, v868, mark_end2108, v869, v870, v871, loadedv2109, v872, result_symbol2111, v873, mark_end2112, v874, v875, v876, loadedv2113, v877, result_symbol2115, v878, mark_end2116, v879, v880, v881, conv2119, cmp2120, v882, idxprom2123, arrayidx2124, v883, conv2125, v884, cmp2126, v885, add2129, idxprom2130, arrayidx2131, v886, v887, add2134, v888, loadedv2136, v889, result_symbol2138, v890, mark_end2139, v891, v892, v893, loadedv2140, v894, result_symbol2142, v895, mark_end2143, v896, v897, v898, cmp2144, v899, cmp2147, v900, cmp2150, v901, cmp2153, v902, cmp2157, v903, cmp2160, v904, cmp2163, v905, loadedv2167, v906, result_symbol2169, v907, mark_end2170, v908, v909, v910, cmp2171, v911, cmp2174, v912, cmp2177, v913, loadedv2181, v914

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
	i117 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i280 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i2117 = libc.Ptr(&new(struct {
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
		goto sw_bb49
	case 3:
		goto sw_bb92
	case 4:
		goto sw_bb116
	case 5:
		goto sw_bb176
	case 6:
		goto sw_bb213
	case 7:
		goto sw_bb246
	case 8:
		goto sw_bb279
	case 9:
		goto sw_bb300
	case 10:
		goto sw_bb340
	case 11:
		goto sw_bb361
	case 12:
		goto sw_bb382
	case 13:
		goto sw_bb403
	case 14:
		goto sw_bb424
	case 15:
		goto sw_bb445
	case 16:
		goto sw_bb466
	case 17:
		goto sw_bb487
	case 18:
		goto sw_bb508
	case 19:
		goto sw_bb555
	case 20:
		goto sw_bb557
	case 21:
		goto sw_bb561
	case 22:
		goto sw_bb565
	case 23:
		goto sw_bb569
	case 24:
		goto sw_bb573
	case 25:
		goto sw_bb611
	case 26:
		goto sw_bb649
	case 27:
		goto sw_bb687
	case 28:
		goto sw_bb725
	case 29:
		goto sw_bb767
	case 30:
		goto sw_bb805
	case 31:
		goto sw_bb843
	case 32:
		goto sw_bb881
	case 33:
		goto sw_bb919
	case 34:
		goto sw_bb957
	case 35:
		goto sw_bb995
	case 36:
		goto sw_bb1033
	case 37:
		goto sw_bb1067
	case 38:
		goto sw_bb1100
	case 39:
		goto sw_bb1129
	case 40:
		goto sw_bb1162
	case 41:
		goto sw_bb1182
	case 42:
		goto sw_bb1186
	case 43:
		goto sw_bb1215
	case 44:
		goto sw_bb1249
	case 45:
		goto sw_bb1275
	case 46:
		goto sw_bb1309
	case 47:
		goto sw_bb1335
	case 48:
		goto sw_bb1369
	case 49:
		goto sw_bb1395
	case 50:
		goto sw_bb1429
	case 51:
		goto sw_bb1455
	case 52:
		goto sw_bb1489
	case 53:
		goto sw_bb1515
	case 54:
		goto sw_bb1549
	case 55:
		goto sw_bb1575
	case 56:
		goto sw_bb1611
	case 57:
		goto sw_bb1637
	case 58:
		goto sw_bb1641
	case 59:
		goto sw_bb1682
	case 60:
		goto sw_bb1716
	case 61:
		goto sw_bb1746
	case 62:
		goto sw_bb1776
	case 63:
		goto sw_bb1806
	case 64:
		goto sw_bb1836
	case 65:
		goto sw_bb1870
	case 66:
		goto sw_bb1900
	case 67:
		goto sw_bb1930
	case 68:
		goto sw_bb1960
	case 69:
		goto sw_bb1990
	case 70:
		goto sw_bb2020
	case 71:
		goto sw_bb2050
	case 72:
		goto sw_bb2080
	case 73:
		goto sw_bb2106
	case 74:
		goto sw_bb2110
	case 75:
		goto sw_bb2114
	case 76:
		goto sw_bb2137
	case 77:
		goto sw_bb2141
	case 78:
		goto sw_bb2168
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
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v11 = *libc.As[int32](i)
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(34)
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
	cmp16 = v19 <= 12
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
	*libc.As[int16](state_addr) = 37
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
	cmp38 = v26 <= 122
	if cmp38 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end41:
	v27 = *libc.As[byte](result)
	loadedv42 = (v27 & 1) != 0
	*libc.As[bool](retval) = loadedv42
	goto _return

sw_bb43:
	v28 = *libc.As[int32](lookahead)
	cmp44 = v28 == 10
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end47:
	v29 = *libc.As[byte](result)
	loadedv48 = (v29 & 1) != 0
	*libc.As[bool](retval) = loadedv48
	goto _return

sw_bb49:
	v30 = *libc.As[int32](lookahead)
	cmp50 = v30 == 10
	if cmp50 {
		goto if_then52
	} else {
		goto if_end53
	}

if_then52:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end53:
	v31 = *libc.As[int32](lookahead)
	cmp54 = v31 == 13
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end57:
	v32 = *libc.As[int32](lookahead)
	cmp58 = v32 == 33
	if cmp58 {
		goto if_then60
	} else {
		goto if_end61
	}

if_then60:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end61:
	v33 = *libc.As[int32](lookahead)
	cmp62 = v33 == 34
	if cmp62 {
		goto if_then64
	} else {
		goto if_end65
	}

if_then64:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end65:
	v34 = *libc.As[int32](lookahead)
	cmp66 = v34 == 92
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end69:
	v35 = *libc.As[int32](lookahead)
	cmp70 = v35 == 35
	if cmp70 {
		goto if_then75
	} else {
		goto lor_lhs_false72
	}

lor_lhs_false72:
	v36 = *libc.As[int32](lookahead)
	cmp73 = v36 == 59
	if cmp73 {
		goto if_then75
	} else {
		goto if_end76
	}

if_then75:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end76:
	v37 = *libc.As[int32](lookahead)
	cmp77 = 9 <= v37
	if cmp77 {
		goto land_lhs_true79
	} else {
		goto lor_lhs_false82
	}

land_lhs_true79:
	v38 = *libc.As[int32](lookahead)
	cmp80 = v38 <= 12
	if cmp80 {
		goto if_then85
	} else {
		goto lor_lhs_false82
	}

lor_lhs_false82:
	v39 = *libc.As[int32](lookahead)
	cmp83 = v39 == 32
	if cmp83 {
		goto if_then85
	} else {
		goto if_end86
	}

if_then85:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end86:
	v40 = *libc.As[int32](lookahead)
	cmp87 = v40 != 0
	if cmp87 {
		goto if_then89
	} else {
		goto if_end90
	}

if_then89:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end90:
	v41 = *libc.As[byte](result)
	loadedv91 = (v41 & 1) != 0
	*libc.As[bool](retval) = loadedv91
	goto _return

sw_bb92:
	v42 = *libc.As[int32](lookahead)
	cmp93 = v42 == 10
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end96:
	v43 = *libc.As[int32](lookahead)
	cmp97 = v43 == 13
	if cmp97 {
		goto if_then99
	} else {
		goto if_end100
	}

if_then99:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end100:
	v44 = *libc.As[int32](lookahead)
	cmp101 = 9 <= v44
	if cmp101 {
		goto land_lhs_true103
	} else {
		goto lor_lhs_false106
	}

land_lhs_true103:
	v45 = *libc.As[int32](lookahead)
	cmp104 = v45 <= 12
	if cmp104 {
		goto if_then109
	} else {
		goto lor_lhs_false106
	}

lor_lhs_false106:
	v46 = *libc.As[int32](lookahead)
	cmp107 = v46 == 32
	if cmp107 {
		goto if_then109
	} else {
		goto if_end110
	}

if_then109:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end110:
	v47 = *libc.As[int32](lookahead)
	cmp111 = v47 != 0
	if cmp111 {
		goto if_then113
	} else {
		goto if_end114
	}

if_then113:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end114:
	v48 = *libc.As[byte](result)
	loadedv115 = (v48 & 1) != 0
	*libc.As[bool](retval) = loadedv115
	goto _return

sw_bb116:
	*libc.As[int32](i117) = 0
	goto for_cond118

for_cond118:
	v49 = *libc.As[int32](i117)
	conv119 = int64(uint64(uint32(v49)))
	cmp120 = uint64(conv119) < uint64(16)
	if cmp120 {
		goto for_body122
	} else {
		goto for_end135
	}

for_body122:
	v50 = *libc.As[int32](i117)
	idxprom123 = int64(uint64(uint32(v50)))
	arrayidx124 = libc.Ptr(&ts_lex_map_43[idxprom123])
	v51 = *libc.As[int16](arrayidx124)
	conv125 = int32(uint32(uint16(v51)))
	v52 = *libc.As[int32](lookahead)
	cmp126 = conv125 == v52
	if cmp126 {
		goto if_then128
	} else {
		goto if_end132
	}

if_then128:
	v53 = *libc.As[int32](i117)
	add129 = v53 + 1
	idxprom130 = int64(uint64(uint32(add129)))
	arrayidx131 = libc.Ptr(&ts_lex_map_43[idxprom130])
	v54 = *libc.As[int16](arrayidx131)
	*libc.As[int16](state_addr) = v54
	goto next_state

if_end132:
	goto for_inc133

for_inc133:
	v55 = *libc.As[int32](i117)
	add134 = v55 + 2
	*libc.As[int32](i117) = add134
	goto for_cond118

for_end135:
	v56 = *libc.As[int32](lookahead)
	cmp136 = v56 == 9
	if cmp136 {
		goto if_then147
	} else {
		goto lor_lhs_false138
	}

lor_lhs_false138:
	v57 = *libc.As[int32](lookahead)
	cmp139 = v57 == 11
	if cmp139 {
		goto if_then147
	} else {
		goto lor_lhs_false141
	}

lor_lhs_false141:
	v58 = *libc.As[int32](lookahead)
	cmp142 = v58 == 12
	if cmp142 {
		goto if_then147
	} else {
		goto lor_lhs_false144
	}

lor_lhs_false144:
	v59 = *libc.As[int32](lookahead)
	cmp145 = v59 == 32
	if cmp145 {
		goto if_then147
	} else {
		goto if_end148
	}

if_then147:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end148:
	v60 = *libc.As[int32](lookahead)
	cmp149 = 48 <= v60
	if cmp149 {
		goto land_lhs_true151
	} else {
		goto if_end155
	}

land_lhs_true151:
	v61 = *libc.As[int32](lookahead)
	cmp152 = v61 <= 57
	if cmp152 {
		goto if_then154
	} else {
		goto if_end155
	}

if_then154:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end155:
	v62 = *libc.As[int32](lookahead)
	cmp156 = v62 != 0
	if cmp156 {
		goto land_lhs_true158
	} else {
		goto if_end174
	}

land_lhs_true158:
	v63 = *libc.As[int32](lookahead)
	cmp159 = v63 < 9
	if cmp159 {
		goto land_lhs_true164
	} else {
		goto lor_lhs_false161
	}

lor_lhs_false161:
	v64 = *libc.As[int32](lookahead)
	cmp162 = 13 < v64
	if cmp162 {
		goto land_lhs_true164
	} else {
		goto if_end174
	}

land_lhs_true164:
	v65 = *libc.As[int32](lookahead)
	cmp165 = v65 < 32
	if cmp165 {
		goto land_lhs_true170
	} else {
		goto lor_lhs_false167
	}

lor_lhs_false167:
	v66 = *libc.As[int32](lookahead)
	cmp168 = 35 < v66
	if cmp168 {
		goto land_lhs_true170
	} else {
		goto if_end174
	}

land_lhs_true170:
	v67 = *libc.As[int32](lookahead)
	cmp171 = v67 != 59
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end174:
	v68 = *libc.As[byte](result)
	loadedv175 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv175
	goto _return

sw_bb176:
	v69 = *libc.As[int32](lookahead)
	cmp177 = v69 == 33
	if cmp177 {
		goto if_then179
	} else {
		goto if_end180
	}

if_then179:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end180:
	v70 = *libc.As[int32](lookahead)
	cmp181 = v70 == 34
	if cmp181 {
		goto if_then183
	} else {
		goto if_end184
	}

if_then183:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end184:
	v71 = *libc.As[int32](lookahead)
	cmp185 = v71 == 92
	if cmp185 {
		goto if_then187
	} else {
		goto if_end188
	}

if_then187:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end188:
	v72 = *libc.As[int32](lookahead)
	cmp189 = v72 == 9
	if cmp189 {
		goto if_then200
	} else {
		goto lor_lhs_false191
	}

lor_lhs_false191:
	v73 = *libc.As[int32](lookahead)
	cmp192 = v73 == 11
	if cmp192 {
		goto if_then200
	} else {
		goto lor_lhs_false194
	}

lor_lhs_false194:
	v74 = *libc.As[int32](lookahead)
	cmp195 = v74 == 12
	if cmp195 {
		goto if_then200
	} else {
		goto lor_lhs_false197
	}

lor_lhs_false197:
	v75 = *libc.As[int32](lookahead)
	cmp198 = v75 == 32
	if cmp198 {
		goto if_then200
	} else {
		goto if_end201
	}

if_then200:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end201:
	v76 = *libc.As[int32](lookahead)
	cmp202 = v76 != 0
	if cmp202 {
		goto land_lhs_true204
	} else {
		goto if_end211
	}

land_lhs_true204:
	v77 = *libc.As[int32](lookahead)
	cmp205 = v77 < 9
	if cmp205 {
		goto if_then210
	} else {
		goto lor_lhs_false207
	}

lor_lhs_false207:
	v78 = *libc.As[int32](lookahead)
	cmp208 = 13 < v78
	if cmp208 {
		goto if_then210
	} else {
		goto if_end211
	}

if_then210:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end211:
	v79 = *libc.As[byte](result)
	loadedv212 = (v79 & 1) != 0
	*libc.As[bool](retval) = loadedv212
	goto _return

sw_bb213:
	v80 = *libc.As[int32](lookahead)
	cmp214 = v80 == 34
	if cmp214 {
		goto if_then216
	} else {
		goto if_end217
	}

if_then216:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end217:
	v81 = *libc.As[int32](lookahead)
	cmp218 = v81 == 92
	if cmp218 {
		goto if_then220
	} else {
		goto if_end221
	}

if_then220:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end221:
	v82 = *libc.As[int32](lookahead)
	cmp222 = v82 == 9
	if cmp222 {
		goto if_then233
	} else {
		goto lor_lhs_false224
	}

lor_lhs_false224:
	v83 = *libc.As[int32](lookahead)
	cmp225 = v83 == 11
	if cmp225 {
		goto if_then233
	} else {
		goto lor_lhs_false227
	}

lor_lhs_false227:
	v84 = *libc.As[int32](lookahead)
	cmp228 = v84 == 12
	if cmp228 {
		goto if_then233
	} else {
		goto lor_lhs_false230
	}

lor_lhs_false230:
	v85 = *libc.As[int32](lookahead)
	cmp231 = v85 == 32
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end234:
	v86 = *libc.As[int32](lookahead)
	cmp235 = v86 != 0
	if cmp235 {
		goto land_lhs_true237
	} else {
		goto if_end244
	}

land_lhs_true237:
	v87 = *libc.As[int32](lookahead)
	cmp238 = v87 < 9
	if cmp238 {
		goto if_then243
	} else {
		goto lor_lhs_false240
	}

lor_lhs_false240:
	v88 = *libc.As[int32](lookahead)
	cmp241 = 13 < v88
	if cmp241 {
		goto if_then243
	} else {
		goto if_end244
	}

if_then243:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end244:
	v89 = *libc.As[byte](result)
	loadedv245 = (v89 & 1) != 0
	*libc.As[bool](retval) = loadedv245
	goto _return

sw_bb246:
	v90 = *libc.As[int32](lookahead)
	cmp247 = v90 == 34
	if cmp247 {
		goto if_then249
	} else {
		goto if_end250
	}

if_then249:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end250:
	v91 = *libc.As[int32](lookahead)
	cmp251 = v91 == 92
	if cmp251 {
		goto if_then253
	} else {
		goto if_end254
	}

if_then253:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end254:
	v92 = *libc.As[int32](lookahead)
	cmp255 = v92 == 9
	if cmp255 {
		goto if_then266
	} else {
		goto lor_lhs_false257
	}

lor_lhs_false257:
	v93 = *libc.As[int32](lookahead)
	cmp258 = v93 == 11
	if cmp258 {
		goto if_then266
	} else {
		goto lor_lhs_false260
	}

lor_lhs_false260:
	v94 = *libc.As[int32](lookahead)
	cmp261 = v94 == 12
	if cmp261 {
		goto if_then266
	} else {
		goto lor_lhs_false263
	}

lor_lhs_false263:
	v95 = *libc.As[int32](lookahead)
	cmp264 = v95 == 32
	if cmp264 {
		goto if_then266
	} else {
		goto if_end267
	}

if_then266:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end267:
	v96 = *libc.As[int32](lookahead)
	cmp268 = v96 != 0
	if cmp268 {
		goto land_lhs_true270
	} else {
		goto if_end277
	}

land_lhs_true270:
	v97 = *libc.As[int32](lookahead)
	cmp271 = v97 < 9
	if cmp271 {
		goto if_then276
	} else {
		goto lor_lhs_false273
	}

lor_lhs_false273:
	v98 = *libc.As[int32](lookahead)
	cmp274 = 13 < v98
	if cmp274 {
		goto if_then276
	} else {
		goto if_end277
	}

if_then276:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end277:
	v99 = *libc.As[byte](result)
	loadedv278 = (v99 & 1) != 0
	*libc.As[bool](retval) = loadedv278
	goto _return

sw_bb279:
	*libc.As[int32](i280) = 0
	goto for_cond281

for_cond281:
	v100 = *libc.As[int32](i280)
	conv282 = int64(uint64(uint32(v100)))
	cmp283 = uint64(conv282) < uint64(18)
	if cmp283 {
		goto for_body285
	} else {
		goto for_end298
	}

for_body285:
	v101 = *libc.As[int32](i280)
	idxprom286 = int64(uint64(uint32(v101)))
	arrayidx287 = libc.Ptr(&ts_lex_map_44[idxprom286])
	v102 = *libc.As[int16](arrayidx287)
	conv288 = int32(uint32(uint16(v102)))
	v103 = *libc.As[int32](lookahead)
	cmp289 = conv288 == v103
	if cmp289 {
		goto if_then291
	} else {
		goto if_end295
	}

if_then291:
	v104 = *libc.As[int32](i280)
	add292 = v104 + 1
	idxprom293 = int64(uint64(uint32(add292)))
	arrayidx294 = libc.Ptr(&ts_lex_map_44[idxprom293])
	v105 = *libc.As[int16](arrayidx294)
	*libc.As[int16](state_addr) = v105
	goto next_state

if_end295:
	goto for_inc296

for_inc296:
	v106 = *libc.As[int32](i280)
	add297 = v106 + 2
	*libc.As[int32](i280) = add297
	goto for_cond281

for_end298:
	v107 = *libc.As[byte](result)
	loadedv299 = (v107 & 1) != 0
	*libc.As[bool](retval) = loadedv299
	goto _return

sw_bb300:
	v108 = *libc.As[int32](lookahead)
	cmp301 = v108 == 9
	if cmp301 {
		goto if_then312
	} else {
		goto lor_lhs_false303
	}

lor_lhs_false303:
	v109 = *libc.As[int32](lookahead)
	cmp304 = v109 == 11
	if cmp304 {
		goto if_then312
	} else {
		goto lor_lhs_false306
	}

lor_lhs_false306:
	v110 = *libc.As[int32](lookahead)
	cmp307 = v110 == 12
	if cmp307 {
		goto if_then312
	} else {
		goto lor_lhs_false309
	}

lor_lhs_false309:
	v111 = *libc.As[int32](lookahead)
	cmp310 = v111 == 32
	if cmp310 {
		goto if_then312
	} else {
		goto if_end313
	}

if_then312:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end313:
	v112 = *libc.As[int32](lookahead)
	cmp314 = v112 == 46
	if cmp314 {
		goto if_then337
	} else {
		goto lor_lhs_false316
	}

lor_lhs_false316:
	v113 = *libc.As[int32](lookahead)
	cmp317 = 48 <= v113
	if cmp317 {
		goto land_lhs_true319
	} else {
		goto lor_lhs_false322
	}

land_lhs_true319:
	v114 = *libc.As[int32](lookahead)
	cmp320 = v114 <= 57
	if cmp320 {
		goto if_then337
	} else {
		goto lor_lhs_false322
	}

lor_lhs_false322:
	v115 = *libc.As[int32](lookahead)
	cmp323 = 65 <= v115
	if cmp323 {
		goto land_lhs_true325
	} else {
		goto lor_lhs_false328
	}

land_lhs_true325:
	v116 = *libc.As[int32](lookahead)
	cmp326 = v116 <= 90
	if cmp326 {
		goto if_then337
	} else {
		goto lor_lhs_false328
	}

lor_lhs_false328:
	v117 = *libc.As[int32](lookahead)
	cmp329 = v117 == 95
	if cmp329 {
		goto if_then337
	} else {
		goto lor_lhs_false331
	}

lor_lhs_false331:
	v118 = *libc.As[int32](lookahead)
	cmp332 = 97 <= v118
	if cmp332 {
		goto land_lhs_true334
	} else {
		goto if_end338
	}

land_lhs_true334:
	v119 = *libc.As[int32](lookahead)
	cmp335 = v119 <= 122
	if cmp335 {
		goto if_then337
	} else {
		goto if_end338
	}

if_then337:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end338:
	v120 = *libc.As[byte](result)
	loadedv339 = (v120 & 1) != 0
	*libc.As[bool](retval) = loadedv339
	goto _return

sw_bb340:
	v121 = *libc.As[int32](lookahead)
	cmp341 = 48 <= v121
	if cmp341 {
		goto land_lhs_true343
	} else {
		goto lor_lhs_false346
	}

land_lhs_true343:
	v122 = *libc.As[int32](lookahead)
	cmp344 = v122 <= 57
	if cmp344 {
		goto if_then358
	} else {
		goto lor_lhs_false346
	}

lor_lhs_false346:
	v123 = *libc.As[int32](lookahead)
	cmp347 = 65 <= v123
	if cmp347 {
		goto land_lhs_true349
	} else {
		goto lor_lhs_false352
	}

land_lhs_true349:
	v124 = *libc.As[int32](lookahead)
	cmp350 = v124 <= 70
	if cmp350 {
		goto if_then358
	} else {
		goto lor_lhs_false352
	}

lor_lhs_false352:
	v125 = *libc.As[int32](lookahead)
	cmp353 = 97 <= v125
	if cmp353 {
		goto land_lhs_true355
	} else {
		goto if_end359
	}

land_lhs_true355:
	v126 = *libc.As[int32](lookahead)
	cmp356 = v126 <= 102
	if cmp356 {
		goto if_then358
	} else {
		goto if_end359
	}

if_then358:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end359:
	v127 = *libc.As[byte](result)
	loadedv360 = (v127 & 1) != 0
	*libc.As[bool](retval) = loadedv360
	goto _return

sw_bb361:
	v128 = *libc.As[int32](lookahead)
	cmp362 = 48 <= v128
	if cmp362 {
		goto land_lhs_true364
	} else {
		goto lor_lhs_false367
	}

land_lhs_true364:
	v129 = *libc.As[int32](lookahead)
	cmp365 = v129 <= 57
	if cmp365 {
		goto if_then379
	} else {
		goto lor_lhs_false367
	}

lor_lhs_false367:
	v130 = *libc.As[int32](lookahead)
	cmp368 = 65 <= v130
	if cmp368 {
		goto land_lhs_true370
	} else {
		goto lor_lhs_false373
	}

land_lhs_true370:
	v131 = *libc.As[int32](lookahead)
	cmp371 = v131 <= 70
	if cmp371 {
		goto if_then379
	} else {
		goto lor_lhs_false373
	}

lor_lhs_false373:
	v132 = *libc.As[int32](lookahead)
	cmp374 = 97 <= v132
	if cmp374 {
		goto land_lhs_true376
	} else {
		goto if_end380
	}

land_lhs_true376:
	v133 = *libc.As[int32](lookahead)
	cmp377 = v133 <= 102
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end380:
	v134 = *libc.As[byte](result)
	loadedv381 = (v134 & 1) != 0
	*libc.As[bool](retval) = loadedv381
	goto _return

sw_bb382:
	v135 = *libc.As[int32](lookahead)
	cmp383 = 48 <= v135
	if cmp383 {
		goto land_lhs_true385
	} else {
		goto lor_lhs_false388
	}

land_lhs_true385:
	v136 = *libc.As[int32](lookahead)
	cmp386 = v136 <= 57
	if cmp386 {
		goto if_then400
	} else {
		goto lor_lhs_false388
	}

lor_lhs_false388:
	v137 = *libc.As[int32](lookahead)
	cmp389 = 65 <= v137
	if cmp389 {
		goto land_lhs_true391
	} else {
		goto lor_lhs_false394
	}

land_lhs_true391:
	v138 = *libc.As[int32](lookahead)
	cmp392 = v138 <= 70
	if cmp392 {
		goto if_then400
	} else {
		goto lor_lhs_false394
	}

lor_lhs_false394:
	v139 = *libc.As[int32](lookahead)
	cmp395 = 97 <= v139
	if cmp395 {
		goto land_lhs_true397
	} else {
		goto if_end401
	}

land_lhs_true397:
	v140 = *libc.As[int32](lookahead)
	cmp398 = v140 <= 102
	if cmp398 {
		goto if_then400
	} else {
		goto if_end401
	}

if_then400:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end401:
	v141 = *libc.As[byte](result)
	loadedv402 = (v141 & 1) != 0
	*libc.As[bool](retval) = loadedv402
	goto _return

sw_bb403:
	v142 = *libc.As[int32](lookahead)
	cmp404 = 48 <= v142
	if cmp404 {
		goto land_lhs_true406
	} else {
		goto lor_lhs_false409
	}

land_lhs_true406:
	v143 = *libc.As[int32](lookahead)
	cmp407 = v143 <= 57
	if cmp407 {
		goto if_then421
	} else {
		goto lor_lhs_false409
	}

lor_lhs_false409:
	v144 = *libc.As[int32](lookahead)
	cmp410 = 65 <= v144
	if cmp410 {
		goto land_lhs_true412
	} else {
		goto lor_lhs_false415
	}

land_lhs_true412:
	v145 = *libc.As[int32](lookahead)
	cmp413 = v145 <= 70
	if cmp413 {
		goto if_then421
	} else {
		goto lor_lhs_false415
	}

lor_lhs_false415:
	v146 = *libc.As[int32](lookahead)
	cmp416 = 97 <= v146
	if cmp416 {
		goto land_lhs_true418
	} else {
		goto if_end422
	}

land_lhs_true418:
	v147 = *libc.As[int32](lookahead)
	cmp419 = v147 <= 102
	if cmp419 {
		goto if_then421
	} else {
		goto if_end422
	}

if_then421:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end422:
	v148 = *libc.As[byte](result)
	loadedv423 = (v148 & 1) != 0
	*libc.As[bool](retval) = loadedv423
	goto _return

sw_bb424:
	v149 = *libc.As[int32](lookahead)
	cmp425 = 48 <= v149
	if cmp425 {
		goto land_lhs_true427
	} else {
		goto lor_lhs_false430
	}

land_lhs_true427:
	v150 = *libc.As[int32](lookahead)
	cmp428 = v150 <= 57
	if cmp428 {
		goto if_then442
	} else {
		goto lor_lhs_false430
	}

lor_lhs_false430:
	v151 = *libc.As[int32](lookahead)
	cmp431 = 65 <= v151
	if cmp431 {
		goto land_lhs_true433
	} else {
		goto lor_lhs_false436
	}

land_lhs_true433:
	v152 = *libc.As[int32](lookahead)
	cmp434 = v152 <= 70
	if cmp434 {
		goto if_then442
	} else {
		goto lor_lhs_false436
	}

lor_lhs_false436:
	v153 = *libc.As[int32](lookahead)
	cmp437 = 97 <= v153
	if cmp437 {
		goto land_lhs_true439
	} else {
		goto if_end443
	}

land_lhs_true439:
	v154 = *libc.As[int32](lookahead)
	cmp440 = v154 <= 102
	if cmp440 {
		goto if_then442
	} else {
		goto if_end443
	}

if_then442:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end443:
	v155 = *libc.As[byte](result)
	loadedv444 = (v155 & 1) != 0
	*libc.As[bool](retval) = loadedv444
	goto _return

sw_bb445:
	v156 = *libc.As[int32](lookahead)
	cmp446 = 48 <= v156
	if cmp446 {
		goto land_lhs_true448
	} else {
		goto lor_lhs_false451
	}

land_lhs_true448:
	v157 = *libc.As[int32](lookahead)
	cmp449 = v157 <= 57
	if cmp449 {
		goto if_then463
	} else {
		goto lor_lhs_false451
	}

lor_lhs_false451:
	v158 = *libc.As[int32](lookahead)
	cmp452 = 65 <= v158
	if cmp452 {
		goto land_lhs_true454
	} else {
		goto lor_lhs_false457
	}

land_lhs_true454:
	v159 = *libc.As[int32](lookahead)
	cmp455 = v159 <= 70
	if cmp455 {
		goto if_then463
	} else {
		goto lor_lhs_false457
	}

lor_lhs_false457:
	v160 = *libc.As[int32](lookahead)
	cmp458 = 97 <= v160
	if cmp458 {
		goto land_lhs_true460
	} else {
		goto if_end464
	}

land_lhs_true460:
	v161 = *libc.As[int32](lookahead)
	cmp461 = v161 <= 102
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end464:
	v162 = *libc.As[byte](result)
	loadedv465 = (v162 & 1) != 0
	*libc.As[bool](retval) = loadedv465
	goto _return

sw_bb466:
	v163 = *libc.As[int32](lookahead)
	cmp467 = 48 <= v163
	if cmp467 {
		goto land_lhs_true469
	} else {
		goto lor_lhs_false472
	}

land_lhs_true469:
	v164 = *libc.As[int32](lookahead)
	cmp470 = v164 <= 57
	if cmp470 {
		goto if_then484
	} else {
		goto lor_lhs_false472
	}

lor_lhs_false472:
	v165 = *libc.As[int32](lookahead)
	cmp473 = 65 <= v165
	if cmp473 {
		goto land_lhs_true475
	} else {
		goto lor_lhs_false478
	}

land_lhs_true475:
	v166 = *libc.As[int32](lookahead)
	cmp476 = v166 <= 70
	if cmp476 {
		goto if_then484
	} else {
		goto lor_lhs_false478
	}

lor_lhs_false478:
	v167 = *libc.As[int32](lookahead)
	cmp479 = 97 <= v167
	if cmp479 {
		goto land_lhs_true481
	} else {
		goto if_end485
	}

land_lhs_true481:
	v168 = *libc.As[int32](lookahead)
	cmp482 = v168 <= 102
	if cmp482 {
		goto if_then484
	} else {
		goto if_end485
	}

if_then484:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end485:
	v169 = *libc.As[byte](result)
	loadedv486 = (v169 & 1) != 0
	*libc.As[bool](retval) = loadedv486
	goto _return

sw_bb487:
	v170 = *libc.As[int32](lookahead)
	cmp488 = 48 <= v170
	if cmp488 {
		goto land_lhs_true490
	} else {
		goto lor_lhs_false493
	}

land_lhs_true490:
	v171 = *libc.As[int32](lookahead)
	cmp491 = v171 <= 57
	if cmp491 {
		goto if_then505
	} else {
		goto lor_lhs_false493
	}

lor_lhs_false493:
	v172 = *libc.As[int32](lookahead)
	cmp494 = 65 <= v172
	if cmp494 {
		goto land_lhs_true496
	} else {
		goto lor_lhs_false499
	}

land_lhs_true496:
	v173 = *libc.As[int32](lookahead)
	cmp497 = v173 <= 70
	if cmp497 {
		goto if_then505
	} else {
		goto lor_lhs_false499
	}

lor_lhs_false499:
	v174 = *libc.As[int32](lookahead)
	cmp500 = 97 <= v174
	if cmp500 {
		goto land_lhs_true502
	} else {
		goto if_end506
	}

land_lhs_true502:
	v175 = *libc.As[int32](lookahead)
	cmp503 = v175 <= 102
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end506:
	v176 = *libc.As[byte](result)
	loadedv507 = (v176 & 1) != 0
	*libc.As[bool](retval) = loadedv507
	goto _return

sw_bb508:
	v177 = *libc.As[byte](eof)
	loadedv509 = (v177 & 1) != 0
	if loadedv509 {
		goto if_then510
	} else {
		goto if_end511
	}

if_then510:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end511:
	v178 = *libc.As[int32](lookahead)
	cmp512 = v178 == 10
	if cmp512 {
		goto if_then514
	} else {
		goto if_end515
	}

if_then514:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end515:
	v179 = *libc.As[int32](lookahead)
	cmp516 = v179 == 13
	if cmp516 {
		goto if_then518
	} else {
		goto if_end519
	}

if_then518:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end519:
	v180 = *libc.As[int32](lookahead)
	cmp520 = v180 == 91
	if cmp520 {
		goto if_then522
	} else {
		goto if_end523
	}

if_then522:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end523:
	v181 = *libc.As[int32](lookahead)
	cmp524 = v181 == 35
	if cmp524 {
		goto if_then529
	} else {
		goto lor_lhs_false526
	}

lor_lhs_false526:
	v182 = *libc.As[int32](lookahead)
	cmp527 = v182 == 59
	if cmp527 {
		goto if_then529
	} else {
		goto if_end530
	}

if_then529:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end530:
	v183 = *libc.As[int32](lookahead)
	cmp531 = 9 <= v183
	if cmp531 {
		goto land_lhs_true533
	} else {
		goto lor_lhs_false536
	}

land_lhs_true533:
	v184 = *libc.As[int32](lookahead)
	cmp534 = v184 <= 12
	if cmp534 {
		goto if_then539
	} else {
		goto lor_lhs_false536
	}

lor_lhs_false536:
	v185 = *libc.As[int32](lookahead)
	cmp537 = v185 == 32
	if cmp537 {
		goto if_then539
	} else {
		goto if_end540
	}

if_then539:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end540:
	v186 = *libc.As[int32](lookahead)
	cmp541 = 65 <= v186
	if cmp541 {
		goto land_lhs_true543
	} else {
		goto lor_lhs_false546
	}

land_lhs_true543:
	v187 = *libc.As[int32](lookahead)
	cmp544 = v187 <= 90
	if cmp544 {
		goto if_then552
	} else {
		goto lor_lhs_false546
	}

lor_lhs_false546:
	v188 = *libc.As[int32](lookahead)
	cmp547 = 97 <= v188
	if cmp547 {
		goto land_lhs_true549
	} else {
		goto if_end553
	}

land_lhs_true549:
	v189 = *libc.As[int32](lookahead)
	cmp550 = v189 <= 122
	if cmp550 {
		goto if_then552
	} else {
		goto if_end553
	}

if_then552:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end553:
	v190 = *libc.As[byte](result)
	loadedv554 = (v190 & 1) != 0
	*libc.As[bool](retval) = loadedv554
	goto _return

sw_bb555:
	*libc.As[byte](result) = 1
	v191 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v191).F1)
	*libc.As[int16](result_symbol) = 0
	v192 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v192).F3)
	v193 = *libc.As[unsafe.Pointer](mark_end)
	v194 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v193)(v194)
	v195 = *libc.As[byte](result)
	loadedv556 = (v195 & 1) != 0
	*libc.As[bool](retval) = loadedv556
	goto _return

sw_bb557:
	*libc.As[byte](result) = 1
	v196 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol558 = libc.Ptr(&libc.As[TSLexer](v196).F1)
	*libc.As[int16](result_symbol558) = 1
	v197 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end559 = libc.Ptr(&libc.As[TSLexer](v197).F3)
	v198 = *libc.As[unsafe.Pointer](mark_end559)
	v199 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v198)(v199)
	v200 = *libc.As[byte](result)
	loadedv560 = (v200 & 1) != 0
	*libc.As[bool](retval) = loadedv560
	goto _return

sw_bb561:
	*libc.As[byte](result) = 1
	v201 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol562 = libc.Ptr(&libc.As[TSLexer](v201).F1)
	*libc.As[int16](result_symbol562) = 2
	v202 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end563 = libc.Ptr(&libc.As[TSLexer](v202).F3)
	v203 = *libc.As[unsafe.Pointer](mark_end563)
	v204 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v203)(v204)
	v205 = *libc.As[byte](result)
	loadedv564 = (v205 & 1) != 0
	*libc.As[bool](retval) = loadedv564
	goto _return

sw_bb565:
	*libc.As[byte](result) = 1
	v206 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol566 = libc.Ptr(&libc.As[TSLexer](v206).F1)
	*libc.As[int16](result_symbol566) = 3
	v207 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end567 = libc.Ptr(&libc.As[TSLexer](v207).F3)
	v208 = *libc.As[unsafe.Pointer](mark_end567)
	v209 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v208)(v209)
	v210 = *libc.As[byte](result)
	loadedv568 = (v210 & 1) != 0
	*libc.As[bool](retval) = loadedv568
	goto _return

sw_bb569:
	*libc.As[byte](result) = 1
	v211 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol570 = libc.Ptr(&libc.As[TSLexer](v211).F1)
	*libc.As[int16](result_symbol570) = 4
	v212 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end571 = libc.Ptr(&libc.As[TSLexer](v212).F3)
	v213 = *libc.As[unsafe.Pointer](mark_end571)
	v214 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v213)(v214)
	v215 = *libc.As[byte](result)
	loadedv572 = (v215 & 1) != 0
	*libc.As[bool](retval) = loadedv572
	goto _return

sw_bb573:
	*libc.As[byte](result) = 1
	v216 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol574 = libc.Ptr(&libc.As[TSLexer](v216).F1)
	*libc.As[int16](result_symbol574) = 5
	v217 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end575 = libc.Ptr(&libc.As[TSLexer](v217).F3)
	v218 = *libc.As[unsafe.Pointer](mark_end575)
	v219 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v218)(v219)
	v220 = *libc.As[int32](lookahead)
	cmp576 = v220 == 45
	if cmp576 {
		goto if_then578
	} else {
		goto if_end579
	}

if_then578:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end579:
	v221 = *libc.As[int32](lookahead)
	cmp580 = v221 == 46
	if cmp580 {
		goto if_then582
	} else {
		goto if_end583
	}

if_then582:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end583:
	v222 = *libc.As[int32](lookahead)
	cmp584 = v222 == 97
	if cmp584 {
		goto if_then586
	} else {
		goto if_end587
	}

if_then586:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end587:
	v223 = *libc.As[int32](lookahead)
	cmp588 = 48 <= v223
	if cmp588 {
		goto land_lhs_true590
	} else {
		goto lor_lhs_false593
	}

land_lhs_true590:
	v224 = *libc.As[int32](lookahead)
	cmp591 = v224 <= 57
	if cmp591 {
		goto if_then608
	} else {
		goto lor_lhs_false593
	}

lor_lhs_false593:
	v225 = *libc.As[int32](lookahead)
	cmp594 = 65 <= v225
	if cmp594 {
		goto land_lhs_true596
	} else {
		goto lor_lhs_false599
	}

land_lhs_true596:
	v226 = *libc.As[int32](lookahead)
	cmp597 = v226 <= 90
	if cmp597 {
		goto if_then608
	} else {
		goto lor_lhs_false599
	}

lor_lhs_false599:
	v227 = *libc.As[int32](lookahead)
	cmp600 = v227 == 95
	if cmp600 {
		goto if_then608
	} else {
		goto lor_lhs_false602
	}

lor_lhs_false602:
	v228 = *libc.As[int32](lookahead)
	cmp603 = 98 <= v228
	if cmp603 {
		goto land_lhs_true605
	} else {
		goto if_end609
	}

land_lhs_true605:
	v229 = *libc.As[int32](lookahead)
	cmp606 = v229 <= 122
	if cmp606 {
		goto if_then608
	} else {
		goto if_end609
	}

if_then608:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end609:
	v230 = *libc.As[byte](result)
	loadedv610 = (v230 & 1) != 0
	*libc.As[bool](retval) = loadedv610
	goto _return

sw_bb611:
	*libc.As[byte](result) = 1
	v231 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol612 = libc.Ptr(&libc.As[TSLexer](v231).F1)
	*libc.As[int16](result_symbol612) = 5
	v232 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end613 = libc.Ptr(&libc.As[TSLexer](v232).F3)
	v233 = *libc.As[unsafe.Pointer](mark_end613)
	v234 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v233)(v234)
	v235 = *libc.As[int32](lookahead)
	cmp614 = v235 == 45
	if cmp614 {
		goto if_then616
	} else {
		goto if_end617
	}

if_then616:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end617:
	v236 = *libc.As[int32](lookahead)
	cmp618 = v236 == 46
	if cmp618 {
		goto if_then620
	} else {
		goto if_end621
	}

if_then620:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end621:
	v237 = *libc.As[int32](lookahead)
	cmp622 = v237 == 101
	if cmp622 {
		goto if_then624
	} else {
		goto if_end625
	}

if_then624:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end625:
	v238 = *libc.As[int32](lookahead)
	cmp626 = 48 <= v238
	if cmp626 {
		goto land_lhs_true628
	} else {
		goto lor_lhs_false631
	}

land_lhs_true628:
	v239 = *libc.As[int32](lookahead)
	cmp629 = v239 <= 57
	if cmp629 {
		goto if_then646
	} else {
		goto lor_lhs_false631
	}

lor_lhs_false631:
	v240 = *libc.As[int32](lookahead)
	cmp632 = 65 <= v240
	if cmp632 {
		goto land_lhs_true634
	} else {
		goto lor_lhs_false637
	}

land_lhs_true634:
	v241 = *libc.As[int32](lookahead)
	cmp635 = v241 <= 90
	if cmp635 {
		goto if_then646
	} else {
		goto lor_lhs_false637
	}

lor_lhs_false637:
	v242 = *libc.As[int32](lookahead)
	cmp638 = v242 == 95
	if cmp638 {
		goto if_then646
	} else {
		goto lor_lhs_false640
	}

lor_lhs_false640:
	v243 = *libc.As[int32](lookahead)
	cmp641 = 97 <= v243
	if cmp641 {
		goto land_lhs_true643
	} else {
		goto if_end647
	}

land_lhs_true643:
	v244 = *libc.As[int32](lookahead)
	cmp644 = v244 <= 122
	if cmp644 {
		goto if_then646
	} else {
		goto if_end647
	}

if_then646:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end647:
	v245 = *libc.As[byte](result)
	loadedv648 = (v245 & 1) != 0
	*libc.As[bool](retval) = loadedv648
	goto _return

sw_bb649:
	*libc.As[byte](result) = 1
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol650 = libc.Ptr(&libc.As[TSLexer](v246).F1)
	*libc.As[int16](result_symbol650) = 5
	v247 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end651 = libc.Ptr(&libc.As[TSLexer](v247).F3)
	v248 = *libc.As[unsafe.Pointer](mark_end651)
	v249 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v248)(v249)
	v250 = *libc.As[int32](lookahead)
	cmp652 = v250 == 45
	if cmp652 {
		goto if_then654
	} else {
		goto if_end655
	}

if_then654:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end655:
	v251 = *libc.As[int32](lookahead)
	cmp656 = v251 == 46
	if cmp656 {
		goto if_then658
	} else {
		goto if_end659
	}

if_then658:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end659:
	v252 = *libc.As[int32](lookahead)
	cmp660 = v252 == 101
	if cmp660 {
		goto if_then662
	} else {
		goto if_end663
	}

if_then662:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end663:
	v253 = *libc.As[int32](lookahead)
	cmp664 = 48 <= v253
	if cmp664 {
		goto land_lhs_true666
	} else {
		goto lor_lhs_false669
	}

land_lhs_true666:
	v254 = *libc.As[int32](lookahead)
	cmp667 = v254 <= 57
	if cmp667 {
		goto if_then684
	} else {
		goto lor_lhs_false669
	}

lor_lhs_false669:
	v255 = *libc.As[int32](lookahead)
	cmp670 = 65 <= v255
	if cmp670 {
		goto land_lhs_true672
	} else {
		goto lor_lhs_false675
	}

land_lhs_true672:
	v256 = *libc.As[int32](lookahead)
	cmp673 = v256 <= 90
	if cmp673 {
		goto if_then684
	} else {
		goto lor_lhs_false675
	}

lor_lhs_false675:
	v257 = *libc.As[int32](lookahead)
	cmp676 = v257 == 95
	if cmp676 {
		goto if_then684
	} else {
		goto lor_lhs_false678
	}

lor_lhs_false678:
	v258 = *libc.As[int32](lookahead)
	cmp679 = 97 <= v258
	if cmp679 {
		goto land_lhs_true681
	} else {
		goto if_end685
	}

land_lhs_true681:
	v259 = *libc.As[int32](lookahead)
	cmp682 = v259 <= 122
	if cmp682 {
		goto if_then684
	} else {
		goto if_end685
	}

if_then684:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end685:
	v260 = *libc.As[byte](result)
	loadedv686 = (v260 & 1) != 0
	*libc.As[bool](retval) = loadedv686
	goto _return

sw_bb687:
	*libc.As[byte](result) = 1
	v261 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol688 = libc.Ptr(&libc.As[TSLexer](v261).F1)
	*libc.As[int16](result_symbol688) = 5
	v262 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end689 = libc.Ptr(&libc.As[TSLexer](v262).F3)
	v263 = *libc.As[unsafe.Pointer](mark_end689)
	v264 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v263)(v264)
	v265 = *libc.As[int32](lookahead)
	cmp690 = v265 == 45
	if cmp690 {
		goto if_then692
	} else {
		goto if_end693
	}

if_then692:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end693:
	v266 = *libc.As[int32](lookahead)
	cmp694 = v266 == 46
	if cmp694 {
		goto if_then696
	} else {
		goto if_end697
	}

if_then696:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end697:
	v267 = *libc.As[int32](lookahead)
	cmp698 = v267 == 101
	if cmp698 {
		goto if_then700
	} else {
		goto if_end701
	}

if_then700:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end701:
	v268 = *libc.As[int32](lookahead)
	cmp702 = 48 <= v268
	if cmp702 {
		goto land_lhs_true704
	} else {
		goto lor_lhs_false707
	}

land_lhs_true704:
	v269 = *libc.As[int32](lookahead)
	cmp705 = v269 <= 57
	if cmp705 {
		goto if_then722
	} else {
		goto lor_lhs_false707
	}

lor_lhs_false707:
	v270 = *libc.As[int32](lookahead)
	cmp708 = 65 <= v270
	if cmp708 {
		goto land_lhs_true710
	} else {
		goto lor_lhs_false713
	}

land_lhs_true710:
	v271 = *libc.As[int32](lookahead)
	cmp711 = v271 <= 90
	if cmp711 {
		goto if_then722
	} else {
		goto lor_lhs_false713
	}

lor_lhs_false713:
	v272 = *libc.As[int32](lookahead)
	cmp714 = v272 == 95
	if cmp714 {
		goto if_then722
	} else {
		goto lor_lhs_false716
	}

lor_lhs_false716:
	v273 = *libc.As[int32](lookahead)
	cmp717 = 97 <= v273
	if cmp717 {
		goto land_lhs_true719
	} else {
		goto if_end723
	}

land_lhs_true719:
	v274 = *libc.As[int32](lookahead)
	cmp720 = v274 <= 122
	if cmp720 {
		goto if_then722
	} else {
		goto if_end723
	}

if_then722:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end723:
	v275 = *libc.As[byte](result)
	loadedv724 = (v275 & 1) != 0
	*libc.As[bool](retval) = loadedv724
	goto _return

sw_bb725:
	*libc.As[byte](result) = 1
	v276 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol726 = libc.Ptr(&libc.As[TSLexer](v276).F1)
	*libc.As[int16](result_symbol726) = 5
	v277 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end727 = libc.Ptr(&libc.As[TSLexer](v277).F3)
	v278 = *libc.As[unsafe.Pointer](mark_end727)
	v279 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v278)(v279)
	v280 = *libc.As[int32](lookahead)
	cmp728 = v280 == 45
	if cmp728 {
		goto if_then730
	} else {
		goto if_end731
	}

if_then730:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end731:
	v281 = *libc.As[int32](lookahead)
	cmp732 = v281 == 46
	if cmp732 {
		goto if_then734
	} else {
		goto if_end735
	}

if_then734:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end735:
	v282 = *libc.As[int32](lookahead)
	cmp736 = v282 == 102
	if cmp736 {
		goto if_then738
	} else {
		goto if_end739
	}

if_then738:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end739:
	v283 = *libc.As[int32](lookahead)
	cmp740 = v283 == 110
	if cmp740 {
		goto if_then742
	} else {
		goto if_end743
	}

if_then742:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end743:
	v284 = *libc.As[int32](lookahead)
	cmp744 = 48 <= v284
	if cmp744 {
		goto land_lhs_true746
	} else {
		goto lor_lhs_false749
	}

land_lhs_true746:
	v285 = *libc.As[int32](lookahead)
	cmp747 = v285 <= 57
	if cmp747 {
		goto if_then764
	} else {
		goto lor_lhs_false749
	}

lor_lhs_false749:
	v286 = *libc.As[int32](lookahead)
	cmp750 = 65 <= v286
	if cmp750 {
		goto land_lhs_true752
	} else {
		goto lor_lhs_false755
	}

land_lhs_true752:
	v287 = *libc.As[int32](lookahead)
	cmp753 = v287 <= 90
	if cmp753 {
		goto if_then764
	} else {
		goto lor_lhs_false755
	}

lor_lhs_false755:
	v288 = *libc.As[int32](lookahead)
	cmp756 = v288 == 95
	if cmp756 {
		goto if_then764
	} else {
		goto lor_lhs_false758
	}

lor_lhs_false758:
	v289 = *libc.As[int32](lookahead)
	cmp759 = 97 <= v289
	if cmp759 {
		goto land_lhs_true761
	} else {
		goto if_end765
	}

land_lhs_true761:
	v290 = *libc.As[int32](lookahead)
	cmp762 = v290 <= 122
	if cmp762 {
		goto if_then764
	} else {
		goto if_end765
	}

if_then764:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end765:
	v291 = *libc.As[byte](result)
	loadedv766 = (v291 & 1) != 0
	*libc.As[bool](retval) = loadedv766
	goto _return

sw_bb767:
	*libc.As[byte](result) = 1
	v292 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol768 = libc.Ptr(&libc.As[TSLexer](v292).F1)
	*libc.As[int16](result_symbol768) = 5
	v293 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end769 = libc.Ptr(&libc.As[TSLexer](v293).F3)
	v294 = *libc.As[unsafe.Pointer](mark_end769)
	v295 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v294)(v295)
	v296 = *libc.As[int32](lookahead)
	cmp770 = v296 == 45
	if cmp770 {
		goto if_then772
	} else {
		goto if_end773
	}

if_then772:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end773:
	v297 = *libc.As[int32](lookahead)
	cmp774 = v297 == 46
	if cmp774 {
		goto if_then776
	} else {
		goto if_end777
	}

if_then776:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end777:
	v298 = *libc.As[int32](lookahead)
	cmp778 = v298 == 102
	if cmp778 {
		goto if_then780
	} else {
		goto if_end781
	}

if_then780:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end781:
	v299 = *libc.As[int32](lookahead)
	cmp782 = 48 <= v299
	if cmp782 {
		goto land_lhs_true784
	} else {
		goto lor_lhs_false787
	}

land_lhs_true784:
	v300 = *libc.As[int32](lookahead)
	cmp785 = v300 <= 57
	if cmp785 {
		goto if_then802
	} else {
		goto lor_lhs_false787
	}

lor_lhs_false787:
	v301 = *libc.As[int32](lookahead)
	cmp788 = 65 <= v301
	if cmp788 {
		goto land_lhs_true790
	} else {
		goto lor_lhs_false793
	}

land_lhs_true790:
	v302 = *libc.As[int32](lookahead)
	cmp791 = v302 <= 90
	if cmp791 {
		goto if_then802
	} else {
		goto lor_lhs_false793
	}

lor_lhs_false793:
	v303 = *libc.As[int32](lookahead)
	cmp794 = v303 == 95
	if cmp794 {
		goto if_then802
	} else {
		goto lor_lhs_false796
	}

lor_lhs_false796:
	v304 = *libc.As[int32](lookahead)
	cmp797 = 97 <= v304
	if cmp797 {
		goto land_lhs_true799
	} else {
		goto if_end803
	}

land_lhs_true799:
	v305 = *libc.As[int32](lookahead)
	cmp800 = v305 <= 122
	if cmp800 {
		goto if_then802
	} else {
		goto if_end803
	}

if_then802:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end803:
	v306 = *libc.As[byte](result)
	loadedv804 = (v306 & 1) != 0
	*libc.As[bool](retval) = loadedv804
	goto _return

sw_bb805:
	*libc.As[byte](result) = 1
	v307 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol806 = libc.Ptr(&libc.As[TSLexer](v307).F1)
	*libc.As[int16](result_symbol806) = 5
	v308 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end807 = libc.Ptr(&libc.As[TSLexer](v308).F3)
	v309 = *libc.As[unsafe.Pointer](mark_end807)
	v310 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v309)(v310)
	v311 = *libc.As[int32](lookahead)
	cmp808 = v311 == 45
	if cmp808 {
		goto if_then810
	} else {
		goto if_end811
	}

if_then810:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end811:
	v312 = *libc.As[int32](lookahead)
	cmp812 = v312 == 46
	if cmp812 {
		goto if_then814
	} else {
		goto if_end815
	}

if_then814:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end815:
	v313 = *libc.As[int32](lookahead)
	cmp816 = v313 == 108
	if cmp816 {
		goto if_then818
	} else {
		goto if_end819
	}

if_then818:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end819:
	v314 = *libc.As[int32](lookahead)
	cmp820 = 48 <= v314
	if cmp820 {
		goto land_lhs_true822
	} else {
		goto lor_lhs_false825
	}

land_lhs_true822:
	v315 = *libc.As[int32](lookahead)
	cmp823 = v315 <= 57
	if cmp823 {
		goto if_then840
	} else {
		goto lor_lhs_false825
	}

lor_lhs_false825:
	v316 = *libc.As[int32](lookahead)
	cmp826 = 65 <= v316
	if cmp826 {
		goto land_lhs_true828
	} else {
		goto lor_lhs_false831
	}

land_lhs_true828:
	v317 = *libc.As[int32](lookahead)
	cmp829 = v317 <= 90
	if cmp829 {
		goto if_then840
	} else {
		goto lor_lhs_false831
	}

lor_lhs_false831:
	v318 = *libc.As[int32](lookahead)
	cmp832 = v318 == 95
	if cmp832 {
		goto if_then840
	} else {
		goto lor_lhs_false834
	}

lor_lhs_false834:
	v319 = *libc.As[int32](lookahead)
	cmp835 = 97 <= v319
	if cmp835 {
		goto land_lhs_true837
	} else {
		goto if_end841
	}

land_lhs_true837:
	v320 = *libc.As[int32](lookahead)
	cmp838 = v320 <= 122
	if cmp838 {
		goto if_then840
	} else {
		goto if_end841
	}

if_then840:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end841:
	v321 = *libc.As[byte](result)
	loadedv842 = (v321 & 1) != 0
	*libc.As[bool](retval) = loadedv842
	goto _return

sw_bb843:
	*libc.As[byte](result) = 1
	v322 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol844 = libc.Ptr(&libc.As[TSLexer](v322).F1)
	*libc.As[int16](result_symbol844) = 5
	v323 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end845 = libc.Ptr(&libc.As[TSLexer](v323).F3)
	v324 = *libc.As[unsafe.Pointer](mark_end845)
	v325 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v324)(v325)
	v326 = *libc.As[int32](lookahead)
	cmp846 = v326 == 45
	if cmp846 {
		goto if_then848
	} else {
		goto if_end849
	}

if_then848:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end849:
	v327 = *libc.As[int32](lookahead)
	cmp850 = v327 == 46
	if cmp850 {
		goto if_then852
	} else {
		goto if_end853
	}

if_then852:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end853:
	v328 = *libc.As[int32](lookahead)
	cmp854 = v328 == 111
	if cmp854 {
		goto if_then856
	} else {
		goto if_end857
	}

if_then856:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end857:
	v329 = *libc.As[int32](lookahead)
	cmp858 = 48 <= v329
	if cmp858 {
		goto land_lhs_true860
	} else {
		goto lor_lhs_false863
	}

land_lhs_true860:
	v330 = *libc.As[int32](lookahead)
	cmp861 = v330 <= 57
	if cmp861 {
		goto if_then878
	} else {
		goto lor_lhs_false863
	}

lor_lhs_false863:
	v331 = *libc.As[int32](lookahead)
	cmp864 = 65 <= v331
	if cmp864 {
		goto land_lhs_true866
	} else {
		goto lor_lhs_false869
	}

land_lhs_true866:
	v332 = *libc.As[int32](lookahead)
	cmp867 = v332 <= 90
	if cmp867 {
		goto if_then878
	} else {
		goto lor_lhs_false869
	}

lor_lhs_false869:
	v333 = *libc.As[int32](lookahead)
	cmp870 = v333 == 95
	if cmp870 {
		goto if_then878
	} else {
		goto lor_lhs_false872
	}

lor_lhs_false872:
	v334 = *libc.As[int32](lookahead)
	cmp873 = 97 <= v334
	if cmp873 {
		goto land_lhs_true875
	} else {
		goto if_end879
	}

land_lhs_true875:
	v335 = *libc.As[int32](lookahead)
	cmp876 = v335 <= 122
	if cmp876 {
		goto if_then878
	} else {
		goto if_end879
	}

if_then878:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end879:
	v336 = *libc.As[byte](result)
	loadedv880 = (v336 & 1) != 0
	*libc.As[bool](retval) = loadedv880
	goto _return

sw_bb881:
	*libc.As[byte](result) = 1
	v337 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol882 = libc.Ptr(&libc.As[TSLexer](v337).F1)
	*libc.As[int16](result_symbol882) = 5
	v338 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end883 = libc.Ptr(&libc.As[TSLexer](v338).F3)
	v339 = *libc.As[unsafe.Pointer](mark_end883)
	v340 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v339)(v340)
	v341 = *libc.As[int32](lookahead)
	cmp884 = v341 == 45
	if cmp884 {
		goto if_then886
	} else {
		goto if_end887
	}

if_then886:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end887:
	v342 = *libc.As[int32](lookahead)
	cmp888 = v342 == 46
	if cmp888 {
		goto if_then890
	} else {
		goto if_end891
	}

if_then890:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end891:
	v343 = *libc.As[int32](lookahead)
	cmp892 = v343 == 114
	if cmp892 {
		goto if_then894
	} else {
		goto if_end895
	}

if_then894:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end895:
	v344 = *libc.As[int32](lookahead)
	cmp896 = 48 <= v344
	if cmp896 {
		goto land_lhs_true898
	} else {
		goto lor_lhs_false901
	}

land_lhs_true898:
	v345 = *libc.As[int32](lookahead)
	cmp899 = v345 <= 57
	if cmp899 {
		goto if_then916
	} else {
		goto lor_lhs_false901
	}

lor_lhs_false901:
	v346 = *libc.As[int32](lookahead)
	cmp902 = 65 <= v346
	if cmp902 {
		goto land_lhs_true904
	} else {
		goto lor_lhs_false907
	}

land_lhs_true904:
	v347 = *libc.As[int32](lookahead)
	cmp905 = v347 <= 90
	if cmp905 {
		goto if_then916
	} else {
		goto lor_lhs_false907
	}

lor_lhs_false907:
	v348 = *libc.As[int32](lookahead)
	cmp908 = v348 == 95
	if cmp908 {
		goto if_then916
	} else {
		goto lor_lhs_false910
	}

lor_lhs_false910:
	v349 = *libc.As[int32](lookahead)
	cmp911 = 97 <= v349
	if cmp911 {
		goto land_lhs_true913
	} else {
		goto if_end917
	}

land_lhs_true913:
	v350 = *libc.As[int32](lookahead)
	cmp914 = v350 <= 122
	if cmp914 {
		goto if_then916
	} else {
		goto if_end917
	}

if_then916:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end917:
	v351 = *libc.As[byte](result)
	loadedv918 = (v351 & 1) != 0
	*libc.As[bool](retval) = loadedv918
	goto _return

sw_bb919:
	*libc.As[byte](result) = 1
	v352 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol920 = libc.Ptr(&libc.As[TSLexer](v352).F1)
	*libc.As[int16](result_symbol920) = 5
	v353 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end921 = libc.Ptr(&libc.As[TSLexer](v353).F3)
	v354 = *libc.As[unsafe.Pointer](mark_end921)
	v355 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v354)(v355)
	v356 = *libc.As[int32](lookahead)
	cmp922 = v356 == 45
	if cmp922 {
		goto if_then924
	} else {
		goto if_end925
	}

if_then924:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end925:
	v357 = *libc.As[int32](lookahead)
	cmp926 = v357 == 46
	if cmp926 {
		goto if_then928
	} else {
		goto if_end929
	}

if_then928:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end929:
	v358 = *libc.As[int32](lookahead)
	cmp930 = v358 == 115
	if cmp930 {
		goto if_then932
	} else {
		goto if_end933
	}

if_then932:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end933:
	v359 = *libc.As[int32](lookahead)
	cmp934 = 48 <= v359
	if cmp934 {
		goto land_lhs_true936
	} else {
		goto lor_lhs_false939
	}

land_lhs_true936:
	v360 = *libc.As[int32](lookahead)
	cmp937 = v360 <= 57
	if cmp937 {
		goto if_then954
	} else {
		goto lor_lhs_false939
	}

lor_lhs_false939:
	v361 = *libc.As[int32](lookahead)
	cmp940 = 65 <= v361
	if cmp940 {
		goto land_lhs_true942
	} else {
		goto lor_lhs_false945
	}

land_lhs_true942:
	v362 = *libc.As[int32](lookahead)
	cmp943 = v362 <= 90
	if cmp943 {
		goto if_then954
	} else {
		goto lor_lhs_false945
	}

lor_lhs_false945:
	v363 = *libc.As[int32](lookahead)
	cmp946 = v363 == 95
	if cmp946 {
		goto if_then954
	} else {
		goto lor_lhs_false948
	}

lor_lhs_false948:
	v364 = *libc.As[int32](lookahead)
	cmp949 = 97 <= v364
	if cmp949 {
		goto land_lhs_true951
	} else {
		goto if_end955
	}

land_lhs_true951:
	v365 = *libc.As[int32](lookahead)
	cmp952 = v365 <= 122
	if cmp952 {
		goto if_then954
	} else {
		goto if_end955
	}

if_then954:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end955:
	v366 = *libc.As[byte](result)
	loadedv956 = (v366 & 1) != 0
	*libc.As[bool](retval) = loadedv956
	goto _return

sw_bb957:
	*libc.As[byte](result) = 1
	v367 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol958 = libc.Ptr(&libc.As[TSLexer](v367).F1)
	*libc.As[int16](result_symbol958) = 5
	v368 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end959 = libc.Ptr(&libc.As[TSLexer](v368).F3)
	v369 = *libc.As[unsafe.Pointer](mark_end959)
	v370 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v369)(v370)
	v371 = *libc.As[int32](lookahead)
	cmp960 = v371 == 45
	if cmp960 {
		goto if_then962
	} else {
		goto if_end963
	}

if_then962:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end963:
	v372 = *libc.As[int32](lookahead)
	cmp964 = v372 == 46
	if cmp964 {
		goto if_then966
	} else {
		goto if_end967
	}

if_then966:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end967:
	v373 = *libc.As[int32](lookahead)
	cmp968 = v373 == 115
	if cmp968 {
		goto if_then970
	} else {
		goto if_end971
	}

if_then970:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end971:
	v374 = *libc.As[int32](lookahead)
	cmp972 = 48 <= v374
	if cmp972 {
		goto land_lhs_true974
	} else {
		goto lor_lhs_false977
	}

land_lhs_true974:
	v375 = *libc.As[int32](lookahead)
	cmp975 = v375 <= 57
	if cmp975 {
		goto if_then992
	} else {
		goto lor_lhs_false977
	}

lor_lhs_false977:
	v376 = *libc.As[int32](lookahead)
	cmp978 = 65 <= v376
	if cmp978 {
		goto land_lhs_true980
	} else {
		goto lor_lhs_false983
	}

land_lhs_true980:
	v377 = *libc.As[int32](lookahead)
	cmp981 = v377 <= 90
	if cmp981 {
		goto if_then992
	} else {
		goto lor_lhs_false983
	}

lor_lhs_false983:
	v378 = *libc.As[int32](lookahead)
	cmp984 = v378 == 95
	if cmp984 {
		goto if_then992
	} else {
		goto lor_lhs_false986
	}

lor_lhs_false986:
	v379 = *libc.As[int32](lookahead)
	cmp987 = 97 <= v379
	if cmp987 {
		goto land_lhs_true989
	} else {
		goto if_end993
	}

land_lhs_true989:
	v380 = *libc.As[int32](lookahead)
	cmp990 = v380 <= 122
	if cmp990 {
		goto if_then992
	} else {
		goto if_end993
	}

if_then992:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end993:
	v381 = *libc.As[byte](result)
	loadedv994 = (v381 & 1) != 0
	*libc.As[bool](retval) = loadedv994
	goto _return

sw_bb995:
	*libc.As[byte](result) = 1
	v382 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol996 = libc.Ptr(&libc.As[TSLexer](v382).F1)
	*libc.As[int16](result_symbol996) = 5
	v383 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end997 = libc.Ptr(&libc.As[TSLexer](v383).F3)
	v384 = *libc.As[unsafe.Pointer](mark_end997)
	v385 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v384)(v385)
	v386 = *libc.As[int32](lookahead)
	cmp998 = v386 == 45
	if cmp998 {
		goto if_then1000
	} else {
		goto if_end1001
	}

if_then1000:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1001:
	v387 = *libc.As[int32](lookahead)
	cmp1002 = v387 == 46
	if cmp1002 {
		goto if_then1004
	} else {
		goto if_end1005
	}

if_then1004:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1005:
	v388 = *libc.As[int32](lookahead)
	cmp1006 = v388 == 117
	if cmp1006 {
		goto if_then1008
	} else {
		goto if_end1009
	}

if_then1008:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end1009:
	v389 = *libc.As[int32](lookahead)
	cmp1010 = 48 <= v389
	if cmp1010 {
		goto land_lhs_true1012
	} else {
		goto lor_lhs_false1015
	}

land_lhs_true1012:
	v390 = *libc.As[int32](lookahead)
	cmp1013 = v390 <= 57
	if cmp1013 {
		goto if_then1030
	} else {
		goto lor_lhs_false1015
	}

lor_lhs_false1015:
	v391 = *libc.As[int32](lookahead)
	cmp1016 = 65 <= v391
	if cmp1016 {
		goto land_lhs_true1018
	} else {
		goto lor_lhs_false1021
	}

land_lhs_true1018:
	v392 = *libc.As[int32](lookahead)
	cmp1019 = v392 <= 90
	if cmp1019 {
		goto if_then1030
	} else {
		goto lor_lhs_false1021
	}

lor_lhs_false1021:
	v393 = *libc.As[int32](lookahead)
	cmp1022 = v393 == 95
	if cmp1022 {
		goto if_then1030
	} else {
		goto lor_lhs_false1024
	}

lor_lhs_false1024:
	v394 = *libc.As[int32](lookahead)
	cmp1025 = 97 <= v394
	if cmp1025 {
		goto land_lhs_true1027
	} else {
		goto if_end1031
	}

land_lhs_true1027:
	v395 = *libc.As[int32](lookahead)
	cmp1028 = v395 <= 122
	if cmp1028 {
		goto if_then1030
	} else {
		goto if_end1031
	}

if_then1030:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1031:
	v396 = *libc.As[byte](result)
	loadedv1032 = (v396 & 1) != 0
	*libc.As[bool](retval) = loadedv1032
	goto _return

sw_bb1033:
	*libc.As[byte](result) = 1
	v397 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1034 = libc.Ptr(&libc.As[TSLexer](v397).F1)
	*libc.As[int16](result_symbol1034) = 5
	v398 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1035 = libc.Ptr(&libc.As[TSLexer](v398).F3)
	v399 = *libc.As[unsafe.Pointer](mark_end1035)
	v400 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v399)(v400)
	v401 = *libc.As[int32](lookahead)
	cmp1036 = v401 == 45
	if cmp1036 {
		goto if_then1038
	} else {
		goto if_end1039
	}

if_then1038:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1039:
	v402 = *libc.As[int32](lookahead)
	cmp1040 = v402 == 46
	if cmp1040 {
		goto if_then1042
	} else {
		goto if_end1043
	}

if_then1042:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1043:
	v403 = *libc.As[int32](lookahead)
	cmp1044 = 48 <= v403
	if cmp1044 {
		goto land_lhs_true1046
	} else {
		goto lor_lhs_false1049
	}

land_lhs_true1046:
	v404 = *libc.As[int32](lookahead)
	cmp1047 = v404 <= 57
	if cmp1047 {
		goto if_then1064
	} else {
		goto lor_lhs_false1049
	}

lor_lhs_false1049:
	v405 = *libc.As[int32](lookahead)
	cmp1050 = 65 <= v405
	if cmp1050 {
		goto land_lhs_true1052
	} else {
		goto lor_lhs_false1055
	}

land_lhs_true1052:
	v406 = *libc.As[int32](lookahead)
	cmp1053 = v406 <= 90
	if cmp1053 {
		goto if_then1064
	} else {
		goto lor_lhs_false1055
	}

lor_lhs_false1055:
	v407 = *libc.As[int32](lookahead)
	cmp1056 = v407 == 95
	if cmp1056 {
		goto if_then1064
	} else {
		goto lor_lhs_false1058
	}

lor_lhs_false1058:
	v408 = *libc.As[int32](lookahead)
	cmp1059 = 97 <= v408
	if cmp1059 {
		goto land_lhs_true1061
	} else {
		goto if_end1065
	}

land_lhs_true1061:
	v409 = *libc.As[int32](lookahead)
	cmp1062 = v409 <= 122
	if cmp1062 {
		goto if_then1064
	} else {
		goto if_end1065
	}

if_then1064:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1065:
	v410 = *libc.As[byte](result)
	loadedv1066 = (v410 & 1) != 0
	*libc.As[bool](retval) = loadedv1066
	goto _return

sw_bb1067:
	*libc.As[byte](result) = 1
	v411 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1068 = libc.Ptr(&libc.As[TSLexer](v411).F1)
	*libc.As[int16](result_symbol1068) = 5
	v412 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1069 = libc.Ptr(&libc.As[TSLexer](v412).F3)
	v413 = *libc.As[unsafe.Pointer](mark_end1069)
	v414 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v413)(v414)
	v415 = *libc.As[int32](lookahead)
	cmp1070 = 48 <= v415
	if cmp1070 {
		goto land_lhs_true1072
	} else {
		goto if_end1076
	}

land_lhs_true1072:
	v416 = *libc.As[int32](lookahead)
	cmp1073 = v416 <= 57
	if cmp1073 {
		goto if_then1075
	} else {
		goto if_end1076
	}

if_then1075:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end1076:
	v417 = *libc.As[int32](lookahead)
	call1077 = set_contains(libc.Ptr(&sym_integer_character_set_1), 15, v417)
	if call1077 {
		goto if_then1078
	} else {
		goto if_end1079
	}

if_then1078:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1079:
	v418 = *libc.As[int32](lookahead)
	cmp1080 = v418 == 46
	if cmp1080 {
		goto if_then1097
	} else {
		goto lor_lhs_false1082
	}

lor_lhs_false1082:
	v419 = *libc.As[int32](lookahead)
	cmp1083 = 65 <= v419
	if cmp1083 {
		goto land_lhs_true1085
	} else {
		goto lor_lhs_false1088
	}

land_lhs_true1085:
	v420 = *libc.As[int32](lookahead)
	cmp1086 = v420 <= 88
	if cmp1086 {
		goto if_then1097
	} else {
		goto lor_lhs_false1088
	}

lor_lhs_false1088:
	v421 = *libc.As[int32](lookahead)
	cmp1089 = v421 == 95
	if cmp1089 {
		goto if_then1097
	} else {
		goto lor_lhs_false1091
	}

lor_lhs_false1091:
	v422 = *libc.As[int32](lookahead)
	cmp1092 = 97 <= v422
	if cmp1092 {
		goto land_lhs_true1094
	} else {
		goto if_end1098
	}

land_lhs_true1094:
	v423 = *libc.As[int32](lookahead)
	cmp1095 = v423 <= 120
	if cmp1095 {
		goto if_then1097
	} else {
		goto if_end1098
	}

if_then1097:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1098:
	v424 = *libc.As[byte](result)
	loadedv1099 = (v424 & 1) != 0
	*libc.As[bool](retval) = loadedv1099
	goto _return

sw_bb1100:
	*libc.As[byte](result) = 1
	v425 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1101 = libc.Ptr(&libc.As[TSLexer](v425).F1)
	*libc.As[int16](result_symbol1101) = 5
	v426 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1102 = libc.Ptr(&libc.As[TSLexer](v426).F3)
	v427 = *libc.As[unsafe.Pointer](mark_end1102)
	v428 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v427)(v428)
	v429 = *libc.As[int32](lookahead)
	cmp1103 = v429 == 46
	if cmp1103 {
		goto if_then1126
	} else {
		goto lor_lhs_false1105
	}

lor_lhs_false1105:
	v430 = *libc.As[int32](lookahead)
	cmp1106 = 48 <= v430
	if cmp1106 {
		goto land_lhs_true1108
	} else {
		goto lor_lhs_false1111
	}

land_lhs_true1108:
	v431 = *libc.As[int32](lookahead)
	cmp1109 = v431 <= 57
	if cmp1109 {
		goto if_then1126
	} else {
		goto lor_lhs_false1111
	}

lor_lhs_false1111:
	v432 = *libc.As[int32](lookahead)
	cmp1112 = 65 <= v432
	if cmp1112 {
		goto land_lhs_true1114
	} else {
		goto lor_lhs_false1117
	}

land_lhs_true1114:
	v433 = *libc.As[int32](lookahead)
	cmp1115 = v433 <= 90
	if cmp1115 {
		goto if_then1126
	} else {
		goto lor_lhs_false1117
	}

lor_lhs_false1117:
	v434 = *libc.As[int32](lookahead)
	cmp1118 = v434 == 95
	if cmp1118 {
		goto if_then1126
	} else {
		goto lor_lhs_false1120
	}

lor_lhs_false1120:
	v435 = *libc.As[int32](lookahead)
	cmp1121 = 97 <= v435
	if cmp1121 {
		goto land_lhs_true1123
	} else {
		goto if_end1127
	}

land_lhs_true1123:
	v436 = *libc.As[int32](lookahead)
	cmp1124 = v436 <= 122
	if cmp1124 {
		goto if_then1126
	} else {
		goto if_end1127
	}

if_then1126:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1127:
	v437 = *libc.As[byte](result)
	loadedv1128 = (v437 & 1) != 0
	*libc.As[bool](retval) = loadedv1128
	goto _return

sw_bb1129:
	*libc.As[byte](result) = 1
	v438 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1130 = libc.Ptr(&libc.As[TSLexer](v438).F1)
	*libc.As[int16](result_symbol1130) = 6
	v439 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1131 = libc.Ptr(&libc.As[TSLexer](v439).F3)
	v440 = *libc.As[unsafe.Pointer](mark_end1131)
	v441 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v440)(v441)
	v442 = *libc.As[int32](lookahead)
	cmp1132 = v442 == 9
	if cmp1132 {
		goto if_then1143
	} else {
		goto lor_lhs_false1134
	}

lor_lhs_false1134:
	v443 = *libc.As[int32](lookahead)
	cmp1135 = v443 == 11
	if cmp1135 {
		goto if_then1143
	} else {
		goto lor_lhs_false1137
	}

lor_lhs_false1137:
	v444 = *libc.As[int32](lookahead)
	cmp1138 = v444 == 12
	if cmp1138 {
		goto if_then1143
	} else {
		goto lor_lhs_false1140
	}

lor_lhs_false1140:
	v445 = *libc.As[int32](lookahead)
	cmp1141 = v445 == 32
	if cmp1141 {
		goto if_then1143
	} else {
		goto if_end1144
	}

if_then1143:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end1144:
	v446 = *libc.As[int32](lookahead)
	cmp1145 = v446 != 0
	if cmp1145 {
		goto land_lhs_true1147
	} else {
		goto if_end1160
	}

land_lhs_true1147:
	v447 = *libc.As[int32](lookahead)
	cmp1148 = v447 < 9
	if cmp1148 {
		goto land_lhs_true1153
	} else {
		goto lor_lhs_false1150
	}

lor_lhs_false1150:
	v448 = *libc.As[int32](lookahead)
	cmp1151 = 13 < v448
	if cmp1151 {
		goto land_lhs_true1153
	} else {
		goto if_end1160
	}

land_lhs_true1153:
	v449 = *libc.As[int32](lookahead)
	cmp1154 = v449 != 34
	if cmp1154 {
		goto land_lhs_true1156
	} else {
		goto if_end1160
	}

land_lhs_true1156:
	v450 = *libc.As[int32](lookahead)
	cmp1157 = v450 != 92
	if cmp1157 {
		goto if_then1159
	} else {
		goto if_end1160
	}

if_then1159:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end1160:
	v451 = *libc.As[byte](result)
	loadedv1161 = (v451 & 1) != 0
	*libc.As[bool](retval) = loadedv1161
	goto _return

sw_bb1162:
	*libc.As[byte](result) = 1
	v452 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1163 = libc.Ptr(&libc.As[TSLexer](v452).F1)
	*libc.As[int16](result_symbol1163) = 6
	v453 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1164 = libc.Ptr(&libc.As[TSLexer](v453).F3)
	v454 = *libc.As[unsafe.Pointer](mark_end1164)
	v455 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v454)(v455)
	v456 = *libc.As[int32](lookahead)
	cmp1165 = v456 != 0
	if cmp1165 {
		goto land_lhs_true1167
	} else {
		goto if_end1180
	}

land_lhs_true1167:
	v457 = *libc.As[int32](lookahead)
	cmp1168 = v457 != 10
	if cmp1168 {
		goto land_lhs_true1170
	} else {
		goto if_end1180
	}

land_lhs_true1170:
	v458 = *libc.As[int32](lookahead)
	cmp1171 = v458 != 13
	if cmp1171 {
		goto land_lhs_true1173
	} else {
		goto if_end1180
	}

land_lhs_true1173:
	v459 = *libc.As[int32](lookahead)
	cmp1174 = v459 != 34
	if cmp1174 {
		goto land_lhs_true1176
	} else {
		goto if_end1180
	}

land_lhs_true1176:
	v460 = *libc.As[int32](lookahead)
	cmp1177 = v460 != 92
	if cmp1177 {
		goto if_then1179
	} else {
		goto if_end1180
	}

if_then1179:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end1180:
	v461 = *libc.As[byte](result)
	loadedv1181 = (v461 & 1) != 0
	*libc.As[bool](retval) = loadedv1181
	goto _return

sw_bb1182:
	*libc.As[byte](result) = 1
	v462 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1183 = libc.Ptr(&libc.As[TSLexer](v462).F1)
	*libc.As[int16](result_symbol1183) = 7
	v463 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1184 = libc.Ptr(&libc.As[TSLexer](v463).F3)
	v464 = *libc.As[unsafe.Pointer](mark_end1184)
	v465 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v464)(v465)
	v466 = *libc.As[byte](result)
	loadedv1185 = (v466 & 1) != 0
	*libc.As[bool](retval) = loadedv1185
	goto _return

sw_bb1186:
	*libc.As[byte](result) = 1
	v467 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1187 = libc.Ptr(&libc.As[TSLexer](v467).F1)
	*libc.As[int16](result_symbol1187) = 8
	v468 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1188 = libc.Ptr(&libc.As[TSLexer](v468).F3)
	v469 = *libc.As[unsafe.Pointer](mark_end1188)
	v470 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v469)(v470)
	v471 = *libc.As[int32](lookahead)
	cmp1189 = v471 == 45
	if cmp1189 {
		goto if_then1212
	} else {
		goto lor_lhs_false1191
	}

lor_lhs_false1191:
	v472 = *libc.As[int32](lookahead)
	cmp1192 = 48 <= v472
	if cmp1192 {
		goto land_lhs_true1194
	} else {
		goto lor_lhs_false1197
	}

land_lhs_true1194:
	v473 = *libc.As[int32](lookahead)
	cmp1195 = v473 <= 57
	if cmp1195 {
		goto if_then1212
	} else {
		goto lor_lhs_false1197
	}

lor_lhs_false1197:
	v474 = *libc.As[int32](lookahead)
	cmp1198 = 65 <= v474
	if cmp1198 {
		goto land_lhs_true1200
	} else {
		goto lor_lhs_false1203
	}

land_lhs_true1200:
	v475 = *libc.As[int32](lookahead)
	cmp1201 = v475 <= 90
	if cmp1201 {
		goto if_then1212
	} else {
		goto lor_lhs_false1203
	}

lor_lhs_false1203:
	v476 = *libc.As[int32](lookahead)
	cmp1204 = v476 == 95
	if cmp1204 {
		goto if_then1212
	} else {
		goto lor_lhs_false1206
	}

lor_lhs_false1206:
	v477 = *libc.As[int32](lookahead)
	cmp1207 = 97 <= v477
	if cmp1207 {
		goto land_lhs_true1209
	} else {
		goto if_end1213
	}

land_lhs_true1209:
	v478 = *libc.As[int32](lookahead)
	cmp1210 = v478 <= 122
	if cmp1210 {
		goto if_then1212
	} else {
		goto if_end1213
	}

if_then1212:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1213:
	v479 = *libc.As[byte](result)
	loadedv1214 = (v479 & 1) != 0
	*libc.As[bool](retval) = loadedv1214
	goto _return

sw_bb1215:
	*libc.As[byte](result) = 1
	v480 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1216 = libc.Ptr(&libc.As[TSLexer](v480).F1)
	*libc.As[int16](result_symbol1216) = 9
	v481 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1217 = libc.Ptr(&libc.As[TSLexer](v481).F3)
	v482 = *libc.As[unsafe.Pointer](mark_end1217)
	v483 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v482)(v483)
	v484 = *libc.As[int32](lookahead)
	cmp1218 = v484 == 45
	if cmp1218 {
		goto if_then1220
	} else {
		goto if_end1221
	}

if_then1220:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1221:
	v485 = *libc.As[int32](lookahead)
	cmp1222 = v485 == 46
	if cmp1222 {
		goto if_then1224
	} else {
		goto if_end1225
	}

if_then1224:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1225:
	v486 = *libc.As[int32](lookahead)
	cmp1226 = 48 <= v486
	if cmp1226 {
		goto land_lhs_true1228
	} else {
		goto lor_lhs_false1231
	}

land_lhs_true1228:
	v487 = *libc.As[int32](lookahead)
	cmp1229 = v487 <= 57
	if cmp1229 {
		goto if_then1246
	} else {
		goto lor_lhs_false1231
	}

lor_lhs_false1231:
	v488 = *libc.As[int32](lookahead)
	cmp1232 = 65 <= v488
	if cmp1232 {
		goto land_lhs_true1234
	} else {
		goto lor_lhs_false1237
	}

land_lhs_true1234:
	v489 = *libc.As[int32](lookahead)
	cmp1235 = v489 <= 90
	if cmp1235 {
		goto if_then1246
	} else {
		goto lor_lhs_false1237
	}

lor_lhs_false1237:
	v490 = *libc.As[int32](lookahead)
	cmp1238 = v490 == 95
	if cmp1238 {
		goto if_then1246
	} else {
		goto lor_lhs_false1240
	}

lor_lhs_false1240:
	v491 = *libc.As[int32](lookahead)
	cmp1241 = 97 <= v491
	if cmp1241 {
		goto land_lhs_true1243
	} else {
		goto if_end1247
	}

land_lhs_true1243:
	v492 = *libc.As[int32](lookahead)
	cmp1244 = v492 <= 122
	if cmp1244 {
		goto if_then1246
	} else {
		goto if_end1247
	}

if_then1246:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1247:
	v493 = *libc.As[byte](result)
	loadedv1248 = (v493 & 1) != 0
	*libc.As[bool](retval) = loadedv1248
	goto _return

sw_bb1249:
	*libc.As[byte](result) = 1
	v494 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1250 = libc.Ptr(&libc.As[TSLexer](v494).F1)
	*libc.As[int16](result_symbol1250) = 9
	v495 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1251 = libc.Ptr(&libc.As[TSLexer](v495).F3)
	v496 = *libc.As[unsafe.Pointer](mark_end1251)
	v497 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v496)(v497)
	v498 = *libc.As[int32](lookahead)
	cmp1252 = v498 != 0
	if cmp1252 {
		goto land_lhs_true1254
	} else {
		goto if_end1273
	}

land_lhs_true1254:
	v499 = *libc.As[int32](lookahead)
	cmp1255 = v499 != 10
	if cmp1255 {
		goto land_lhs_true1257
	} else {
		goto if_end1273
	}

land_lhs_true1257:
	v500 = *libc.As[int32](lookahead)
	cmp1258 = v500 != 13
	if cmp1258 {
		goto land_lhs_true1260
	} else {
		goto if_end1273
	}

land_lhs_true1260:
	v501 = *libc.As[int32](lookahead)
	cmp1261 = v501 != 34
	if cmp1261 {
		goto land_lhs_true1263
	} else {
		goto if_end1273
	}

land_lhs_true1263:
	v502 = *libc.As[int32](lookahead)
	cmp1264 = v502 != 35
	if cmp1264 {
		goto land_lhs_true1266
	} else {
		goto if_end1273
	}

land_lhs_true1266:
	v503 = *libc.As[int32](lookahead)
	cmp1267 = v503 != 59
	if cmp1267 {
		goto land_lhs_true1269
	} else {
		goto if_end1273
	}

land_lhs_true1269:
	v504 = *libc.As[int32](lookahead)
	cmp1270 = v504 != 92
	if cmp1270 {
		goto if_then1272
	} else {
		goto if_end1273
	}

if_then1272:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1273:
	v505 = *libc.As[byte](result)
	loadedv1274 = (v505 & 1) != 0
	*libc.As[bool](retval) = loadedv1274
	goto _return

sw_bb1275:
	*libc.As[byte](result) = 1
	v506 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1276 = libc.Ptr(&libc.As[TSLexer](v506).F1)
	*libc.As[int16](result_symbol1276) = 10
	v507 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1277 = libc.Ptr(&libc.As[TSLexer](v507).F3)
	v508 = *libc.As[unsafe.Pointer](mark_end1277)
	v509 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v508)(v509)
	v510 = *libc.As[int32](lookahead)
	cmp1278 = v510 == 45
	if cmp1278 {
		goto if_then1280
	} else {
		goto if_end1281
	}

if_then1280:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1281:
	v511 = *libc.As[int32](lookahead)
	cmp1282 = v511 == 46
	if cmp1282 {
		goto if_then1284
	} else {
		goto if_end1285
	}

if_then1284:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1285:
	v512 = *libc.As[int32](lookahead)
	cmp1286 = 48 <= v512
	if cmp1286 {
		goto land_lhs_true1288
	} else {
		goto lor_lhs_false1291
	}

land_lhs_true1288:
	v513 = *libc.As[int32](lookahead)
	cmp1289 = v513 <= 57
	if cmp1289 {
		goto if_then1306
	} else {
		goto lor_lhs_false1291
	}

lor_lhs_false1291:
	v514 = *libc.As[int32](lookahead)
	cmp1292 = 65 <= v514
	if cmp1292 {
		goto land_lhs_true1294
	} else {
		goto lor_lhs_false1297
	}

land_lhs_true1294:
	v515 = *libc.As[int32](lookahead)
	cmp1295 = v515 <= 90
	if cmp1295 {
		goto if_then1306
	} else {
		goto lor_lhs_false1297
	}

lor_lhs_false1297:
	v516 = *libc.As[int32](lookahead)
	cmp1298 = v516 == 95
	if cmp1298 {
		goto if_then1306
	} else {
		goto lor_lhs_false1300
	}

lor_lhs_false1300:
	v517 = *libc.As[int32](lookahead)
	cmp1301 = 97 <= v517
	if cmp1301 {
		goto land_lhs_true1303
	} else {
		goto if_end1307
	}

land_lhs_true1303:
	v518 = *libc.As[int32](lookahead)
	cmp1304 = v518 <= 122
	if cmp1304 {
		goto if_then1306
	} else {
		goto if_end1307
	}

if_then1306:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1307:
	v519 = *libc.As[byte](result)
	loadedv1308 = (v519 & 1) != 0
	*libc.As[bool](retval) = loadedv1308
	goto _return

sw_bb1309:
	*libc.As[byte](result) = 1
	v520 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1310 = libc.Ptr(&libc.As[TSLexer](v520).F1)
	*libc.As[int16](result_symbol1310) = 10
	v521 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1311 = libc.Ptr(&libc.As[TSLexer](v521).F3)
	v522 = *libc.As[unsafe.Pointer](mark_end1311)
	v523 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v522)(v523)
	v524 = *libc.As[int32](lookahead)
	cmp1312 = v524 != 0
	if cmp1312 {
		goto land_lhs_true1314
	} else {
		goto if_end1333
	}

land_lhs_true1314:
	v525 = *libc.As[int32](lookahead)
	cmp1315 = v525 != 10
	if cmp1315 {
		goto land_lhs_true1317
	} else {
		goto if_end1333
	}

land_lhs_true1317:
	v526 = *libc.As[int32](lookahead)
	cmp1318 = v526 != 13
	if cmp1318 {
		goto land_lhs_true1320
	} else {
		goto if_end1333
	}

land_lhs_true1320:
	v527 = *libc.As[int32](lookahead)
	cmp1321 = v527 != 34
	if cmp1321 {
		goto land_lhs_true1323
	} else {
		goto if_end1333
	}

land_lhs_true1323:
	v528 = *libc.As[int32](lookahead)
	cmp1324 = v528 != 35
	if cmp1324 {
		goto land_lhs_true1326
	} else {
		goto if_end1333
	}

land_lhs_true1326:
	v529 = *libc.As[int32](lookahead)
	cmp1327 = v529 != 59
	if cmp1327 {
		goto land_lhs_true1329
	} else {
		goto if_end1333
	}

land_lhs_true1329:
	v530 = *libc.As[int32](lookahead)
	cmp1330 = v530 != 92
	if cmp1330 {
		goto if_then1332
	} else {
		goto if_end1333
	}

if_then1332:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1333:
	v531 = *libc.As[byte](result)
	loadedv1334 = (v531 & 1) != 0
	*libc.As[bool](retval) = loadedv1334
	goto _return

sw_bb1335:
	*libc.As[byte](result) = 1
	v532 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1336 = libc.Ptr(&libc.As[TSLexer](v532).F1)
	*libc.As[int16](result_symbol1336) = 11
	v533 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1337 = libc.Ptr(&libc.As[TSLexer](v533).F3)
	v534 = *libc.As[unsafe.Pointer](mark_end1337)
	v535 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v534)(v535)
	v536 = *libc.As[int32](lookahead)
	cmp1338 = v536 == 45
	if cmp1338 {
		goto if_then1340
	} else {
		goto if_end1341
	}

if_then1340:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1341:
	v537 = *libc.As[int32](lookahead)
	cmp1342 = v537 == 46
	if cmp1342 {
		goto if_then1344
	} else {
		goto if_end1345
	}

if_then1344:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1345:
	v538 = *libc.As[int32](lookahead)
	cmp1346 = 48 <= v538
	if cmp1346 {
		goto land_lhs_true1348
	} else {
		goto lor_lhs_false1351
	}

land_lhs_true1348:
	v539 = *libc.As[int32](lookahead)
	cmp1349 = v539 <= 57
	if cmp1349 {
		goto if_then1366
	} else {
		goto lor_lhs_false1351
	}

lor_lhs_false1351:
	v540 = *libc.As[int32](lookahead)
	cmp1352 = 65 <= v540
	if cmp1352 {
		goto land_lhs_true1354
	} else {
		goto lor_lhs_false1357
	}

land_lhs_true1354:
	v541 = *libc.As[int32](lookahead)
	cmp1355 = v541 <= 90
	if cmp1355 {
		goto if_then1366
	} else {
		goto lor_lhs_false1357
	}

lor_lhs_false1357:
	v542 = *libc.As[int32](lookahead)
	cmp1358 = v542 == 95
	if cmp1358 {
		goto if_then1366
	} else {
		goto lor_lhs_false1360
	}

lor_lhs_false1360:
	v543 = *libc.As[int32](lookahead)
	cmp1361 = 97 <= v543
	if cmp1361 {
		goto land_lhs_true1363
	} else {
		goto if_end1367
	}

land_lhs_true1363:
	v544 = *libc.As[int32](lookahead)
	cmp1364 = v544 <= 122
	if cmp1364 {
		goto if_then1366
	} else {
		goto if_end1367
	}

if_then1366:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1367:
	v545 = *libc.As[byte](result)
	loadedv1368 = (v545 & 1) != 0
	*libc.As[bool](retval) = loadedv1368
	goto _return

sw_bb1369:
	*libc.As[byte](result) = 1
	v546 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1370 = libc.Ptr(&libc.As[TSLexer](v546).F1)
	*libc.As[int16](result_symbol1370) = 11
	v547 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1371 = libc.Ptr(&libc.As[TSLexer](v547).F3)
	v548 = *libc.As[unsafe.Pointer](mark_end1371)
	v549 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v548)(v549)
	v550 = *libc.As[int32](lookahead)
	cmp1372 = v550 != 0
	if cmp1372 {
		goto land_lhs_true1374
	} else {
		goto if_end1393
	}

land_lhs_true1374:
	v551 = *libc.As[int32](lookahead)
	cmp1375 = v551 != 10
	if cmp1375 {
		goto land_lhs_true1377
	} else {
		goto if_end1393
	}

land_lhs_true1377:
	v552 = *libc.As[int32](lookahead)
	cmp1378 = v552 != 13
	if cmp1378 {
		goto land_lhs_true1380
	} else {
		goto if_end1393
	}

land_lhs_true1380:
	v553 = *libc.As[int32](lookahead)
	cmp1381 = v553 != 34
	if cmp1381 {
		goto land_lhs_true1383
	} else {
		goto if_end1393
	}

land_lhs_true1383:
	v554 = *libc.As[int32](lookahead)
	cmp1384 = v554 != 35
	if cmp1384 {
		goto land_lhs_true1386
	} else {
		goto if_end1393
	}

land_lhs_true1386:
	v555 = *libc.As[int32](lookahead)
	cmp1387 = v555 != 59
	if cmp1387 {
		goto land_lhs_true1389
	} else {
		goto if_end1393
	}

land_lhs_true1389:
	v556 = *libc.As[int32](lookahead)
	cmp1390 = v556 != 92
	if cmp1390 {
		goto if_then1392
	} else {
		goto if_end1393
	}

if_then1392:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1393:
	v557 = *libc.As[byte](result)
	loadedv1394 = (v557 & 1) != 0
	*libc.As[bool](retval) = loadedv1394
	goto _return

sw_bb1395:
	*libc.As[byte](result) = 1
	v558 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1396 = libc.Ptr(&libc.As[TSLexer](v558).F1)
	*libc.As[int16](result_symbol1396) = 12
	v559 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1397 = libc.Ptr(&libc.As[TSLexer](v559).F3)
	v560 = *libc.As[unsafe.Pointer](mark_end1397)
	v561 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v560)(v561)
	v562 = *libc.As[int32](lookahead)
	cmp1398 = v562 == 45
	if cmp1398 {
		goto if_then1400
	} else {
		goto if_end1401
	}

if_then1400:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1401:
	v563 = *libc.As[int32](lookahead)
	cmp1402 = v563 == 46
	if cmp1402 {
		goto if_then1404
	} else {
		goto if_end1405
	}

if_then1404:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1405:
	v564 = *libc.As[int32](lookahead)
	cmp1406 = 48 <= v564
	if cmp1406 {
		goto land_lhs_true1408
	} else {
		goto lor_lhs_false1411
	}

land_lhs_true1408:
	v565 = *libc.As[int32](lookahead)
	cmp1409 = v565 <= 57
	if cmp1409 {
		goto if_then1426
	} else {
		goto lor_lhs_false1411
	}

lor_lhs_false1411:
	v566 = *libc.As[int32](lookahead)
	cmp1412 = 65 <= v566
	if cmp1412 {
		goto land_lhs_true1414
	} else {
		goto lor_lhs_false1417
	}

land_lhs_true1414:
	v567 = *libc.As[int32](lookahead)
	cmp1415 = v567 <= 90
	if cmp1415 {
		goto if_then1426
	} else {
		goto lor_lhs_false1417
	}

lor_lhs_false1417:
	v568 = *libc.As[int32](lookahead)
	cmp1418 = v568 == 95
	if cmp1418 {
		goto if_then1426
	} else {
		goto lor_lhs_false1420
	}

lor_lhs_false1420:
	v569 = *libc.As[int32](lookahead)
	cmp1421 = 97 <= v569
	if cmp1421 {
		goto land_lhs_true1423
	} else {
		goto if_end1427
	}

land_lhs_true1423:
	v570 = *libc.As[int32](lookahead)
	cmp1424 = v570 <= 122
	if cmp1424 {
		goto if_then1426
	} else {
		goto if_end1427
	}

if_then1426:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1427:
	v571 = *libc.As[byte](result)
	loadedv1428 = (v571 & 1) != 0
	*libc.As[bool](retval) = loadedv1428
	goto _return

sw_bb1429:
	*libc.As[byte](result) = 1
	v572 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1430 = libc.Ptr(&libc.As[TSLexer](v572).F1)
	*libc.As[int16](result_symbol1430) = 12
	v573 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1431 = libc.Ptr(&libc.As[TSLexer](v573).F3)
	v574 = *libc.As[unsafe.Pointer](mark_end1431)
	v575 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v574)(v575)
	v576 = *libc.As[int32](lookahead)
	cmp1432 = v576 != 0
	if cmp1432 {
		goto land_lhs_true1434
	} else {
		goto if_end1453
	}

land_lhs_true1434:
	v577 = *libc.As[int32](lookahead)
	cmp1435 = v577 != 10
	if cmp1435 {
		goto land_lhs_true1437
	} else {
		goto if_end1453
	}

land_lhs_true1437:
	v578 = *libc.As[int32](lookahead)
	cmp1438 = v578 != 13
	if cmp1438 {
		goto land_lhs_true1440
	} else {
		goto if_end1453
	}

land_lhs_true1440:
	v579 = *libc.As[int32](lookahead)
	cmp1441 = v579 != 34
	if cmp1441 {
		goto land_lhs_true1443
	} else {
		goto if_end1453
	}

land_lhs_true1443:
	v580 = *libc.As[int32](lookahead)
	cmp1444 = v580 != 35
	if cmp1444 {
		goto land_lhs_true1446
	} else {
		goto if_end1453
	}

land_lhs_true1446:
	v581 = *libc.As[int32](lookahead)
	cmp1447 = v581 != 59
	if cmp1447 {
		goto land_lhs_true1449
	} else {
		goto if_end1453
	}

land_lhs_true1449:
	v582 = *libc.As[int32](lookahead)
	cmp1450 = v582 != 92
	if cmp1450 {
		goto if_then1452
	} else {
		goto if_end1453
	}

if_then1452:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1453:
	v583 = *libc.As[byte](result)
	loadedv1454 = (v583 & 1) != 0
	*libc.As[bool](retval) = loadedv1454
	goto _return

sw_bb1455:
	*libc.As[byte](result) = 1
	v584 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1456 = libc.Ptr(&libc.As[TSLexer](v584).F1)
	*libc.As[int16](result_symbol1456) = 13
	v585 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1457 = libc.Ptr(&libc.As[TSLexer](v585).F3)
	v586 = *libc.As[unsafe.Pointer](mark_end1457)
	v587 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v586)(v587)
	v588 = *libc.As[int32](lookahead)
	cmp1458 = v588 == 45
	if cmp1458 {
		goto if_then1460
	} else {
		goto if_end1461
	}

if_then1460:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1461:
	v589 = *libc.As[int32](lookahead)
	cmp1462 = v589 == 46
	if cmp1462 {
		goto if_then1464
	} else {
		goto if_end1465
	}

if_then1464:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1465:
	v590 = *libc.As[int32](lookahead)
	cmp1466 = 48 <= v590
	if cmp1466 {
		goto land_lhs_true1468
	} else {
		goto lor_lhs_false1471
	}

land_lhs_true1468:
	v591 = *libc.As[int32](lookahead)
	cmp1469 = v591 <= 57
	if cmp1469 {
		goto if_then1486
	} else {
		goto lor_lhs_false1471
	}

lor_lhs_false1471:
	v592 = *libc.As[int32](lookahead)
	cmp1472 = 65 <= v592
	if cmp1472 {
		goto land_lhs_true1474
	} else {
		goto lor_lhs_false1477
	}

land_lhs_true1474:
	v593 = *libc.As[int32](lookahead)
	cmp1475 = v593 <= 90
	if cmp1475 {
		goto if_then1486
	} else {
		goto lor_lhs_false1477
	}

lor_lhs_false1477:
	v594 = *libc.As[int32](lookahead)
	cmp1478 = v594 == 95
	if cmp1478 {
		goto if_then1486
	} else {
		goto lor_lhs_false1480
	}

lor_lhs_false1480:
	v595 = *libc.As[int32](lookahead)
	cmp1481 = 97 <= v595
	if cmp1481 {
		goto land_lhs_true1483
	} else {
		goto if_end1487
	}

land_lhs_true1483:
	v596 = *libc.As[int32](lookahead)
	cmp1484 = v596 <= 122
	if cmp1484 {
		goto if_then1486
	} else {
		goto if_end1487
	}

if_then1486:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1487:
	v597 = *libc.As[byte](result)
	loadedv1488 = (v597 & 1) != 0
	*libc.As[bool](retval) = loadedv1488
	goto _return

sw_bb1489:
	*libc.As[byte](result) = 1
	v598 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1490 = libc.Ptr(&libc.As[TSLexer](v598).F1)
	*libc.As[int16](result_symbol1490) = 13
	v599 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1491 = libc.Ptr(&libc.As[TSLexer](v599).F3)
	v600 = *libc.As[unsafe.Pointer](mark_end1491)
	v601 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v600)(v601)
	v602 = *libc.As[int32](lookahead)
	cmp1492 = v602 != 0
	if cmp1492 {
		goto land_lhs_true1494
	} else {
		goto if_end1513
	}

land_lhs_true1494:
	v603 = *libc.As[int32](lookahead)
	cmp1495 = v603 != 10
	if cmp1495 {
		goto land_lhs_true1497
	} else {
		goto if_end1513
	}

land_lhs_true1497:
	v604 = *libc.As[int32](lookahead)
	cmp1498 = v604 != 13
	if cmp1498 {
		goto land_lhs_true1500
	} else {
		goto if_end1513
	}

land_lhs_true1500:
	v605 = *libc.As[int32](lookahead)
	cmp1501 = v605 != 34
	if cmp1501 {
		goto land_lhs_true1503
	} else {
		goto if_end1513
	}

land_lhs_true1503:
	v606 = *libc.As[int32](lookahead)
	cmp1504 = v606 != 35
	if cmp1504 {
		goto land_lhs_true1506
	} else {
		goto if_end1513
	}

land_lhs_true1506:
	v607 = *libc.As[int32](lookahead)
	cmp1507 = v607 != 59
	if cmp1507 {
		goto land_lhs_true1509
	} else {
		goto if_end1513
	}

land_lhs_true1509:
	v608 = *libc.As[int32](lookahead)
	cmp1510 = v608 != 92
	if cmp1510 {
		goto if_then1512
	} else {
		goto if_end1513
	}

if_then1512:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1513:
	v609 = *libc.As[byte](result)
	loadedv1514 = (v609 & 1) != 0
	*libc.As[bool](retval) = loadedv1514
	goto _return

sw_bb1515:
	*libc.As[byte](result) = 1
	v610 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1516 = libc.Ptr(&libc.As[TSLexer](v610).F1)
	*libc.As[int16](result_symbol1516) = 14
	v611 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1517 = libc.Ptr(&libc.As[TSLexer](v611).F3)
	v612 = *libc.As[unsafe.Pointer](mark_end1517)
	v613 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v612)(v613)
	v614 = *libc.As[int32](lookahead)
	cmp1518 = v614 == 45
	if cmp1518 {
		goto if_then1520
	} else {
		goto if_end1521
	}

if_then1520:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end1521:
	v615 = *libc.As[int32](lookahead)
	cmp1522 = v615 == 46
	if cmp1522 {
		goto if_then1524
	} else {
		goto if_end1525
	}

if_then1524:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end1525:
	v616 = *libc.As[int32](lookahead)
	cmp1526 = 48 <= v616
	if cmp1526 {
		goto land_lhs_true1528
	} else {
		goto lor_lhs_false1531
	}

land_lhs_true1528:
	v617 = *libc.As[int32](lookahead)
	cmp1529 = v617 <= 57
	if cmp1529 {
		goto if_then1546
	} else {
		goto lor_lhs_false1531
	}

lor_lhs_false1531:
	v618 = *libc.As[int32](lookahead)
	cmp1532 = 65 <= v618
	if cmp1532 {
		goto land_lhs_true1534
	} else {
		goto lor_lhs_false1537
	}

land_lhs_true1534:
	v619 = *libc.As[int32](lookahead)
	cmp1535 = v619 <= 90
	if cmp1535 {
		goto if_then1546
	} else {
		goto lor_lhs_false1537
	}

lor_lhs_false1537:
	v620 = *libc.As[int32](lookahead)
	cmp1538 = v620 == 95
	if cmp1538 {
		goto if_then1546
	} else {
		goto lor_lhs_false1540
	}

lor_lhs_false1540:
	v621 = *libc.As[int32](lookahead)
	cmp1541 = 97 <= v621
	if cmp1541 {
		goto land_lhs_true1543
	} else {
		goto if_end1547
	}

land_lhs_true1543:
	v622 = *libc.As[int32](lookahead)
	cmp1544 = v622 <= 122
	if cmp1544 {
		goto if_then1546
	} else {
		goto if_end1547
	}

if_then1546:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end1547:
	v623 = *libc.As[byte](result)
	loadedv1548 = (v623 & 1) != 0
	*libc.As[bool](retval) = loadedv1548
	goto _return

sw_bb1549:
	*libc.As[byte](result) = 1
	v624 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1550 = libc.Ptr(&libc.As[TSLexer](v624).F1)
	*libc.As[int16](result_symbol1550) = 14
	v625 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1551 = libc.Ptr(&libc.As[TSLexer](v625).F3)
	v626 = *libc.As[unsafe.Pointer](mark_end1551)
	v627 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v626)(v627)
	v628 = *libc.As[int32](lookahead)
	cmp1552 = v628 != 0
	if cmp1552 {
		goto land_lhs_true1554
	} else {
		goto if_end1573
	}

land_lhs_true1554:
	v629 = *libc.As[int32](lookahead)
	cmp1555 = v629 != 10
	if cmp1555 {
		goto land_lhs_true1557
	} else {
		goto if_end1573
	}

land_lhs_true1557:
	v630 = *libc.As[int32](lookahead)
	cmp1558 = v630 != 13
	if cmp1558 {
		goto land_lhs_true1560
	} else {
		goto if_end1573
	}

land_lhs_true1560:
	v631 = *libc.As[int32](lookahead)
	cmp1561 = v631 != 34
	if cmp1561 {
		goto land_lhs_true1563
	} else {
		goto if_end1573
	}

land_lhs_true1563:
	v632 = *libc.As[int32](lookahead)
	cmp1564 = v632 != 35
	if cmp1564 {
		goto land_lhs_true1566
	} else {
		goto if_end1573
	}

land_lhs_true1566:
	v633 = *libc.As[int32](lookahead)
	cmp1567 = v633 != 59
	if cmp1567 {
		goto land_lhs_true1569
	} else {
		goto if_end1573
	}

land_lhs_true1569:
	v634 = *libc.As[int32](lookahead)
	cmp1570 = v634 != 92
	if cmp1570 {
		goto if_then1572
	} else {
		goto if_end1573
	}

if_then1572:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1573:
	v635 = *libc.As[byte](result)
	loadedv1574 = (v635 & 1) != 0
	*libc.As[bool](retval) = loadedv1574
	goto _return

sw_bb1575:
	*libc.As[byte](result) = 1
	v636 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1576 = libc.Ptr(&libc.As[TSLexer](v636).F1)
	*libc.As[int16](result_symbol1576) = 15
	v637 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1577 = libc.Ptr(&libc.As[TSLexer](v637).F3)
	v638 = *libc.As[unsafe.Pointer](mark_end1577)
	v639 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v638)(v639)
	v640 = *libc.As[int32](lookahead)
	cmp1578 = 48 <= v640
	if cmp1578 {
		goto land_lhs_true1580
	} else {
		goto if_end1584
	}

land_lhs_true1580:
	v641 = *libc.As[int32](lookahead)
	cmp1581 = v641 <= 57
	if cmp1581 {
		goto if_then1583
	} else {
		goto if_end1584
	}

if_then1583:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end1584:
	v642 = *libc.As[int32](lookahead)
	call1585 = set_contains(libc.Ptr(&sym_integer_character_set_1), 15, v642)
	if call1585 {
		goto if_then1586
	} else {
		goto if_end1587
	}

if_then1586:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end1587:
	v643 = *libc.As[int32](lookahead)
	cmp1588 = v643 != 0
	if cmp1588 {
		goto land_lhs_true1590
	} else {
		goto if_end1609
	}

land_lhs_true1590:
	v644 = *libc.As[int32](lookahead)
	cmp1591 = v644 != 10
	if cmp1591 {
		goto land_lhs_true1593
	} else {
		goto if_end1609
	}

land_lhs_true1593:
	v645 = *libc.As[int32](lookahead)
	cmp1594 = v645 != 13
	if cmp1594 {
		goto land_lhs_true1596
	} else {
		goto if_end1609
	}

land_lhs_true1596:
	v646 = *libc.As[int32](lookahead)
	cmp1597 = v646 != 34
	if cmp1597 {
		goto land_lhs_true1599
	} else {
		goto if_end1609
	}

land_lhs_true1599:
	v647 = *libc.As[int32](lookahead)
	cmp1600 = v647 != 35
	if cmp1600 {
		goto land_lhs_true1602
	} else {
		goto if_end1609
	}

land_lhs_true1602:
	v648 = *libc.As[int32](lookahead)
	cmp1603 = v648 != 59
	if cmp1603 {
		goto land_lhs_true1605
	} else {
		goto if_end1609
	}

land_lhs_true1605:
	v649 = *libc.As[int32](lookahead)
	cmp1606 = v649 != 92
	if cmp1606 {
		goto if_then1608
	} else {
		goto if_end1609
	}

if_then1608:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1609:
	v650 = *libc.As[byte](result)
	loadedv1610 = (v650 & 1) != 0
	*libc.As[bool](retval) = loadedv1610
	goto _return

sw_bb1611:
	*libc.As[byte](result) = 1
	v651 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1612 = libc.Ptr(&libc.As[TSLexer](v651).F1)
	*libc.As[int16](result_symbol1612) = 15
	v652 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1613 = libc.Ptr(&libc.As[TSLexer](v652).F3)
	v653 = *libc.As[unsafe.Pointer](mark_end1613)
	v654 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v653)(v654)
	v655 = *libc.As[int32](lookahead)
	cmp1614 = v655 != 0
	if cmp1614 {
		goto land_lhs_true1616
	} else {
		goto if_end1635
	}

land_lhs_true1616:
	v656 = *libc.As[int32](lookahead)
	cmp1617 = v656 != 10
	if cmp1617 {
		goto land_lhs_true1619
	} else {
		goto if_end1635
	}

land_lhs_true1619:
	v657 = *libc.As[int32](lookahead)
	cmp1620 = v657 != 13
	if cmp1620 {
		goto land_lhs_true1622
	} else {
		goto if_end1635
	}

land_lhs_true1622:
	v658 = *libc.As[int32](lookahead)
	cmp1623 = v658 != 34
	if cmp1623 {
		goto land_lhs_true1625
	} else {
		goto if_end1635
	}

land_lhs_true1625:
	v659 = *libc.As[int32](lookahead)
	cmp1626 = v659 != 35
	if cmp1626 {
		goto land_lhs_true1628
	} else {
		goto if_end1635
	}

land_lhs_true1628:
	v660 = *libc.As[int32](lookahead)
	cmp1629 = v660 != 59
	if cmp1629 {
		goto land_lhs_true1631
	} else {
		goto if_end1635
	}

land_lhs_true1631:
	v661 = *libc.As[int32](lookahead)
	cmp1632 = v661 != 92
	if cmp1632 {
		goto if_then1634
	} else {
		goto if_end1635
	}

if_then1634:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1635:
	v662 = *libc.As[byte](result)
	loadedv1636 = (v662 & 1) != 0
	*libc.As[bool](retval) = loadedv1636
	goto _return

sw_bb1637:
	*libc.As[byte](result) = 1
	v663 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1638 = libc.Ptr(&libc.As[TSLexer](v663).F1)
	*libc.As[int16](result_symbol1638) = 16
	v664 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1639 = libc.Ptr(&libc.As[TSLexer](v664).F3)
	v665 = *libc.As[unsafe.Pointer](mark_end1639)
	v666 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v665)(v666)
	v667 = *libc.As[byte](result)
	loadedv1640 = (v667 & 1) != 0
	*libc.As[bool](retval) = loadedv1640
	goto _return

sw_bb1641:
	*libc.As[byte](result) = 1
	v668 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1642 = libc.Ptr(&libc.As[TSLexer](v668).F1)
	*libc.As[int16](result_symbol1642) = 16
	v669 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1643 = libc.Ptr(&libc.As[TSLexer](v669).F3)
	v670 = *libc.As[unsafe.Pointer](mark_end1643)
	v671 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v670)(v671)
	v672 = *libc.As[int32](lookahead)
	cmp1644 = v672 == 33
	if cmp1644 {
		goto if_then1646
	} else {
		goto if_end1647
	}

if_then1646:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end1647:
	v673 = *libc.As[int32](lookahead)
	cmp1648 = v673 == 92
	if cmp1648 {
		goto if_then1650
	} else {
		goto if_end1651
	}

if_then1650:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end1651:
	v674 = *libc.As[int32](lookahead)
	cmp1652 = v674 == 9
	if cmp1652 {
		goto if_then1663
	} else {
		goto lor_lhs_false1654
	}

lor_lhs_false1654:
	v675 = *libc.As[int32](lookahead)
	cmp1655 = v675 == 11
	if cmp1655 {
		goto if_then1663
	} else {
		goto lor_lhs_false1657
	}

lor_lhs_false1657:
	v676 = *libc.As[int32](lookahead)
	cmp1658 = v676 == 12
	if cmp1658 {
		goto if_then1663
	} else {
		goto lor_lhs_false1660
	}

lor_lhs_false1660:
	v677 = *libc.As[int32](lookahead)
	cmp1661 = v677 == 32
	if cmp1661 {
		goto if_then1663
	} else {
		goto if_end1664
	}

if_then1663:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end1664:
	v678 = *libc.As[int32](lookahead)
	cmp1665 = v678 != 0
	if cmp1665 {
		goto land_lhs_true1667
	} else {
		goto if_end1680
	}

land_lhs_true1667:
	v679 = *libc.As[int32](lookahead)
	cmp1668 = v679 < 9
	if cmp1668 {
		goto land_lhs_true1673
	} else {
		goto lor_lhs_false1670
	}

lor_lhs_false1670:
	v680 = *libc.As[int32](lookahead)
	cmp1671 = 13 < v680
	if cmp1671 {
		goto land_lhs_true1673
	} else {
		goto if_end1680
	}

land_lhs_true1673:
	v681 = *libc.As[int32](lookahead)
	cmp1674 = v681 < 32
	if cmp1674 {
		goto if_then1679
	} else {
		goto lor_lhs_false1676
	}

lor_lhs_false1676:
	v682 = *libc.As[int32](lookahead)
	cmp1677 = 34 < v682
	if cmp1677 {
		goto if_then1679
	} else {
		goto if_end1680
	}

if_then1679:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end1680:
	v683 = *libc.As[byte](result)
	loadedv1681 = (v683 & 1) != 0
	*libc.As[bool](retval) = loadedv1681
	goto _return

sw_bb1682:
	*libc.As[byte](result) = 1
	v684 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1683 = libc.Ptr(&libc.As[TSLexer](v684).F1)
	*libc.As[int16](result_symbol1683) = 16
	v685 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1684 = libc.Ptr(&libc.As[TSLexer](v685).F3)
	v686 = *libc.As[unsafe.Pointer](mark_end1684)
	v687 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v686)(v687)
	v688 = *libc.As[int32](lookahead)
	cmp1685 = v688 == 92
	if cmp1685 {
		goto if_then1687
	} else {
		goto if_end1688
	}

if_then1687:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end1688:
	v689 = *libc.As[int32](lookahead)
	cmp1689 = v689 == 9
	if cmp1689 {
		goto if_then1700
	} else {
		goto lor_lhs_false1691
	}

lor_lhs_false1691:
	v690 = *libc.As[int32](lookahead)
	cmp1692 = v690 == 11
	if cmp1692 {
		goto if_then1700
	} else {
		goto lor_lhs_false1694
	}

lor_lhs_false1694:
	v691 = *libc.As[int32](lookahead)
	cmp1695 = v691 == 12
	if cmp1695 {
		goto if_then1700
	} else {
		goto lor_lhs_false1697
	}

lor_lhs_false1697:
	v692 = *libc.As[int32](lookahead)
	cmp1698 = v692 == 32
	if cmp1698 {
		goto if_then1700
	} else {
		goto if_end1701
	}

if_then1700:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end1701:
	v693 = *libc.As[int32](lookahead)
	cmp1702 = v693 != 0
	if cmp1702 {
		goto land_lhs_true1704
	} else {
		goto if_end1714
	}

land_lhs_true1704:
	v694 = *libc.As[int32](lookahead)
	cmp1705 = v694 < 9
	if cmp1705 {
		goto land_lhs_true1710
	} else {
		goto lor_lhs_false1707
	}

lor_lhs_false1707:
	v695 = *libc.As[int32](lookahead)
	cmp1708 = 13 < v695
	if cmp1708 {
		goto land_lhs_true1710
	} else {
		goto if_end1714
	}

land_lhs_true1710:
	v696 = *libc.As[int32](lookahead)
	cmp1711 = v696 != 34
	if cmp1711 {
		goto if_then1713
	} else {
		goto if_end1714
	}

if_then1713:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end1714:
	v697 = *libc.As[byte](result)
	loadedv1715 = (v697 & 1) != 0
	*libc.As[bool](retval) = loadedv1715
	goto _return

sw_bb1716:
	*libc.As[byte](result) = 1
	v698 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1717 = libc.Ptr(&libc.As[TSLexer](v698).F1)
	*libc.As[int16](result_symbol1717) = 17
	v699 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1718 = libc.Ptr(&libc.As[TSLexer](v699).F3)
	v700 = *libc.As[unsafe.Pointer](mark_end1718)
	v701 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v700)(v701)
	v702 = *libc.As[int32](lookahead)
	cmp1719 = v702 == 97
	if cmp1719 {
		goto if_then1721
	} else {
		goto if_end1722
	}

if_then1721:
	*libc.As[int16](state_addr) = 66
	goto next_state

if_end1722:
	v703 = *libc.As[int32](lookahead)
	cmp1723 = v703 != 0
	if cmp1723 {
		goto land_lhs_true1725
	} else {
		goto if_end1744
	}

land_lhs_true1725:
	v704 = *libc.As[int32](lookahead)
	cmp1726 = v704 != 10
	if cmp1726 {
		goto land_lhs_true1728
	} else {
		goto if_end1744
	}

land_lhs_true1728:
	v705 = *libc.As[int32](lookahead)
	cmp1729 = v705 != 13
	if cmp1729 {
		goto land_lhs_true1731
	} else {
		goto if_end1744
	}

land_lhs_true1731:
	v706 = *libc.As[int32](lookahead)
	cmp1732 = v706 != 34
	if cmp1732 {
		goto land_lhs_true1734
	} else {
		goto if_end1744
	}

land_lhs_true1734:
	v707 = *libc.As[int32](lookahead)
	cmp1735 = v707 != 35
	if cmp1735 {
		goto land_lhs_true1737
	} else {
		goto if_end1744
	}

land_lhs_true1737:
	v708 = *libc.As[int32](lookahead)
	cmp1738 = v708 != 59
	if cmp1738 {
		goto land_lhs_true1740
	} else {
		goto if_end1744
	}

land_lhs_true1740:
	v709 = *libc.As[int32](lookahead)
	cmp1741 = v709 != 92
	if cmp1741 {
		goto if_then1743
	} else {
		goto if_end1744
	}

if_then1743:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1744:
	v710 = *libc.As[byte](result)
	loadedv1745 = (v710 & 1) != 0
	*libc.As[bool](retval) = loadedv1745
	goto _return

sw_bb1746:
	*libc.As[byte](result) = 1
	v711 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1747 = libc.Ptr(&libc.As[TSLexer](v711).F1)
	*libc.As[int16](result_symbol1747) = 17
	v712 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1748 = libc.Ptr(&libc.As[TSLexer](v712).F3)
	v713 = *libc.As[unsafe.Pointer](mark_end1748)
	v714 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v713)(v714)
	v715 = *libc.As[int32](lookahead)
	cmp1749 = v715 == 101
	if cmp1749 {
		goto if_then1751
	} else {
		goto if_end1752
	}

if_then1751:
	*libc.As[int16](state_addr) = 69
	goto next_state

if_end1752:
	v716 = *libc.As[int32](lookahead)
	cmp1753 = v716 != 0
	if cmp1753 {
		goto land_lhs_true1755
	} else {
		goto if_end1774
	}

land_lhs_true1755:
	v717 = *libc.As[int32](lookahead)
	cmp1756 = v717 != 10
	if cmp1756 {
		goto land_lhs_true1758
	} else {
		goto if_end1774
	}

land_lhs_true1758:
	v718 = *libc.As[int32](lookahead)
	cmp1759 = v718 != 13
	if cmp1759 {
		goto land_lhs_true1761
	} else {
		goto if_end1774
	}

land_lhs_true1761:
	v719 = *libc.As[int32](lookahead)
	cmp1762 = v719 != 34
	if cmp1762 {
		goto land_lhs_true1764
	} else {
		goto if_end1774
	}

land_lhs_true1764:
	v720 = *libc.As[int32](lookahead)
	cmp1765 = v720 != 35
	if cmp1765 {
		goto land_lhs_true1767
	} else {
		goto if_end1774
	}

land_lhs_true1767:
	v721 = *libc.As[int32](lookahead)
	cmp1768 = v721 != 59
	if cmp1768 {
		goto land_lhs_true1770
	} else {
		goto if_end1774
	}

land_lhs_true1770:
	v722 = *libc.As[int32](lookahead)
	cmp1771 = v722 != 92
	if cmp1771 {
		goto if_then1773
	} else {
		goto if_end1774
	}

if_then1773:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1774:
	v723 = *libc.As[byte](result)
	loadedv1775 = (v723 & 1) != 0
	*libc.As[bool](retval) = loadedv1775
	goto _return

sw_bb1776:
	*libc.As[byte](result) = 1
	v724 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1777 = libc.Ptr(&libc.As[TSLexer](v724).F1)
	*libc.As[int16](result_symbol1777) = 17
	v725 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1778 = libc.Ptr(&libc.As[TSLexer](v725).F3)
	v726 = *libc.As[unsafe.Pointer](mark_end1778)
	v727 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v726)(v727)
	v728 = *libc.As[int32](lookahead)
	cmp1779 = v728 == 101
	if cmp1779 {
		goto if_then1781
	} else {
		goto if_end1782
	}

if_then1781:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end1782:
	v729 = *libc.As[int32](lookahead)
	cmp1783 = v729 != 0
	if cmp1783 {
		goto land_lhs_true1785
	} else {
		goto if_end1804
	}

land_lhs_true1785:
	v730 = *libc.As[int32](lookahead)
	cmp1786 = v730 != 10
	if cmp1786 {
		goto land_lhs_true1788
	} else {
		goto if_end1804
	}

land_lhs_true1788:
	v731 = *libc.As[int32](lookahead)
	cmp1789 = v731 != 13
	if cmp1789 {
		goto land_lhs_true1791
	} else {
		goto if_end1804
	}

land_lhs_true1791:
	v732 = *libc.As[int32](lookahead)
	cmp1792 = v732 != 34
	if cmp1792 {
		goto land_lhs_true1794
	} else {
		goto if_end1804
	}

land_lhs_true1794:
	v733 = *libc.As[int32](lookahead)
	cmp1795 = v733 != 35
	if cmp1795 {
		goto land_lhs_true1797
	} else {
		goto if_end1804
	}

land_lhs_true1797:
	v734 = *libc.As[int32](lookahead)
	cmp1798 = v734 != 59
	if cmp1798 {
		goto land_lhs_true1800
	} else {
		goto if_end1804
	}

land_lhs_true1800:
	v735 = *libc.As[int32](lookahead)
	cmp1801 = v735 != 92
	if cmp1801 {
		goto if_then1803
	} else {
		goto if_end1804
	}

if_then1803:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1804:
	v736 = *libc.As[byte](result)
	loadedv1805 = (v736 & 1) != 0
	*libc.As[bool](retval) = loadedv1805
	goto _return

sw_bb1806:
	*libc.As[byte](result) = 1
	v737 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1807 = libc.Ptr(&libc.As[TSLexer](v737).F1)
	*libc.As[int16](result_symbol1807) = 17
	v738 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1808 = libc.Ptr(&libc.As[TSLexer](v738).F3)
	v739 = *libc.As[unsafe.Pointer](mark_end1808)
	v740 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v739)(v740)
	v741 = *libc.As[int32](lookahead)
	cmp1809 = v741 == 101
	if cmp1809 {
		goto if_then1811
	} else {
		goto if_end1812
	}

if_then1811:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end1812:
	v742 = *libc.As[int32](lookahead)
	cmp1813 = v742 != 0
	if cmp1813 {
		goto land_lhs_true1815
	} else {
		goto if_end1834
	}

land_lhs_true1815:
	v743 = *libc.As[int32](lookahead)
	cmp1816 = v743 != 10
	if cmp1816 {
		goto land_lhs_true1818
	} else {
		goto if_end1834
	}

land_lhs_true1818:
	v744 = *libc.As[int32](lookahead)
	cmp1819 = v744 != 13
	if cmp1819 {
		goto land_lhs_true1821
	} else {
		goto if_end1834
	}

land_lhs_true1821:
	v745 = *libc.As[int32](lookahead)
	cmp1822 = v745 != 34
	if cmp1822 {
		goto land_lhs_true1824
	} else {
		goto if_end1834
	}

land_lhs_true1824:
	v746 = *libc.As[int32](lookahead)
	cmp1825 = v746 != 35
	if cmp1825 {
		goto land_lhs_true1827
	} else {
		goto if_end1834
	}

land_lhs_true1827:
	v747 = *libc.As[int32](lookahead)
	cmp1828 = v747 != 59
	if cmp1828 {
		goto land_lhs_true1830
	} else {
		goto if_end1834
	}

land_lhs_true1830:
	v748 = *libc.As[int32](lookahead)
	cmp1831 = v748 != 92
	if cmp1831 {
		goto if_then1833
	} else {
		goto if_end1834
	}

if_then1833:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1834:
	v749 = *libc.As[byte](result)
	loadedv1835 = (v749 & 1) != 0
	*libc.As[bool](retval) = loadedv1835
	goto _return

sw_bb1836:
	*libc.As[byte](result) = 1
	v750 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1837 = libc.Ptr(&libc.As[TSLexer](v750).F1)
	*libc.As[int16](result_symbol1837) = 17
	v751 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1838 = libc.Ptr(&libc.As[TSLexer](v751).F3)
	v752 = *libc.As[unsafe.Pointer](mark_end1838)
	v753 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v752)(v753)
	v754 = *libc.As[int32](lookahead)
	cmp1839 = v754 == 102
	if cmp1839 {
		goto if_then1841
	} else {
		goto if_end1842
	}

if_then1841:
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end1842:
	v755 = *libc.As[int32](lookahead)
	cmp1843 = v755 == 110
	if cmp1843 {
		goto if_then1845
	} else {
		goto if_end1846
	}

if_then1845:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end1846:
	v756 = *libc.As[int32](lookahead)
	cmp1847 = v756 != 0
	if cmp1847 {
		goto land_lhs_true1849
	} else {
		goto if_end1868
	}

land_lhs_true1849:
	v757 = *libc.As[int32](lookahead)
	cmp1850 = v757 != 10
	if cmp1850 {
		goto land_lhs_true1852
	} else {
		goto if_end1868
	}

land_lhs_true1852:
	v758 = *libc.As[int32](lookahead)
	cmp1853 = v758 != 13
	if cmp1853 {
		goto land_lhs_true1855
	} else {
		goto if_end1868
	}

land_lhs_true1855:
	v759 = *libc.As[int32](lookahead)
	cmp1856 = v759 != 34
	if cmp1856 {
		goto land_lhs_true1858
	} else {
		goto if_end1868
	}

land_lhs_true1858:
	v760 = *libc.As[int32](lookahead)
	cmp1859 = v760 != 35
	if cmp1859 {
		goto land_lhs_true1861
	} else {
		goto if_end1868
	}

land_lhs_true1861:
	v761 = *libc.As[int32](lookahead)
	cmp1862 = v761 != 59
	if cmp1862 {
		goto land_lhs_true1864
	} else {
		goto if_end1868
	}

land_lhs_true1864:
	v762 = *libc.As[int32](lookahead)
	cmp1865 = v762 != 92
	if cmp1865 {
		goto if_then1867
	} else {
		goto if_end1868
	}

if_then1867:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1868:
	v763 = *libc.As[byte](result)
	loadedv1869 = (v763 & 1) != 0
	*libc.As[bool](retval) = loadedv1869
	goto _return

sw_bb1870:
	*libc.As[byte](result) = 1
	v764 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1871 = libc.Ptr(&libc.As[TSLexer](v764).F1)
	*libc.As[int16](result_symbol1871) = 17
	v765 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1872 = libc.Ptr(&libc.As[TSLexer](v765).F3)
	v766 = *libc.As[unsafe.Pointer](mark_end1872)
	v767 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v766)(v767)
	v768 = *libc.As[int32](lookahead)
	cmp1873 = v768 == 102
	if cmp1873 {
		goto if_then1875
	} else {
		goto if_end1876
	}

if_then1875:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end1876:
	v769 = *libc.As[int32](lookahead)
	cmp1877 = v769 != 0
	if cmp1877 {
		goto land_lhs_true1879
	} else {
		goto if_end1898
	}

land_lhs_true1879:
	v770 = *libc.As[int32](lookahead)
	cmp1880 = v770 != 10
	if cmp1880 {
		goto land_lhs_true1882
	} else {
		goto if_end1898
	}

land_lhs_true1882:
	v771 = *libc.As[int32](lookahead)
	cmp1883 = v771 != 13
	if cmp1883 {
		goto land_lhs_true1885
	} else {
		goto if_end1898
	}

land_lhs_true1885:
	v772 = *libc.As[int32](lookahead)
	cmp1886 = v772 != 34
	if cmp1886 {
		goto land_lhs_true1888
	} else {
		goto if_end1898
	}

land_lhs_true1888:
	v773 = *libc.As[int32](lookahead)
	cmp1889 = v773 != 35
	if cmp1889 {
		goto land_lhs_true1891
	} else {
		goto if_end1898
	}

land_lhs_true1891:
	v774 = *libc.As[int32](lookahead)
	cmp1892 = v774 != 59
	if cmp1892 {
		goto land_lhs_true1894
	} else {
		goto if_end1898
	}

land_lhs_true1894:
	v775 = *libc.As[int32](lookahead)
	cmp1895 = v775 != 92
	if cmp1895 {
		goto if_then1897
	} else {
		goto if_end1898
	}

if_then1897:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1898:
	v776 = *libc.As[byte](result)
	loadedv1899 = (v776 & 1) != 0
	*libc.As[bool](retval) = loadedv1899
	goto _return

sw_bb1900:
	*libc.As[byte](result) = 1
	v777 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1901 = libc.Ptr(&libc.As[TSLexer](v777).F1)
	*libc.As[int16](result_symbol1901) = 17
	v778 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1902 = libc.Ptr(&libc.As[TSLexer](v778).F3)
	v779 = *libc.As[unsafe.Pointer](mark_end1902)
	v780 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v779)(v780)
	v781 = *libc.As[int32](lookahead)
	cmp1903 = v781 == 108
	if cmp1903 {
		goto if_then1905
	} else {
		goto if_end1906
	}

if_then1905:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end1906:
	v782 = *libc.As[int32](lookahead)
	cmp1907 = v782 != 0
	if cmp1907 {
		goto land_lhs_true1909
	} else {
		goto if_end1928
	}

land_lhs_true1909:
	v783 = *libc.As[int32](lookahead)
	cmp1910 = v783 != 10
	if cmp1910 {
		goto land_lhs_true1912
	} else {
		goto if_end1928
	}

land_lhs_true1912:
	v784 = *libc.As[int32](lookahead)
	cmp1913 = v784 != 13
	if cmp1913 {
		goto land_lhs_true1915
	} else {
		goto if_end1928
	}

land_lhs_true1915:
	v785 = *libc.As[int32](lookahead)
	cmp1916 = v785 != 34
	if cmp1916 {
		goto land_lhs_true1918
	} else {
		goto if_end1928
	}

land_lhs_true1918:
	v786 = *libc.As[int32](lookahead)
	cmp1919 = v786 != 35
	if cmp1919 {
		goto land_lhs_true1921
	} else {
		goto if_end1928
	}

land_lhs_true1921:
	v787 = *libc.As[int32](lookahead)
	cmp1922 = v787 != 59
	if cmp1922 {
		goto land_lhs_true1924
	} else {
		goto if_end1928
	}

land_lhs_true1924:
	v788 = *libc.As[int32](lookahead)
	cmp1925 = v788 != 92
	if cmp1925 {
		goto if_then1927
	} else {
		goto if_end1928
	}

if_then1927:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1928:
	v789 = *libc.As[byte](result)
	loadedv1929 = (v789 & 1) != 0
	*libc.As[bool](retval) = loadedv1929
	goto _return

sw_bb1930:
	*libc.As[byte](result) = 1
	v790 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1931 = libc.Ptr(&libc.As[TSLexer](v790).F1)
	*libc.As[int16](result_symbol1931) = 17
	v791 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1932 = libc.Ptr(&libc.As[TSLexer](v791).F3)
	v792 = *libc.As[unsafe.Pointer](mark_end1932)
	v793 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v792)(v793)
	v794 = *libc.As[int32](lookahead)
	cmp1933 = v794 == 111
	if cmp1933 {
		goto if_then1935
	} else {
		goto if_end1936
	}

if_then1935:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end1936:
	v795 = *libc.As[int32](lookahead)
	cmp1937 = v795 != 0
	if cmp1937 {
		goto land_lhs_true1939
	} else {
		goto if_end1958
	}

land_lhs_true1939:
	v796 = *libc.As[int32](lookahead)
	cmp1940 = v796 != 10
	if cmp1940 {
		goto land_lhs_true1942
	} else {
		goto if_end1958
	}

land_lhs_true1942:
	v797 = *libc.As[int32](lookahead)
	cmp1943 = v797 != 13
	if cmp1943 {
		goto land_lhs_true1945
	} else {
		goto if_end1958
	}

land_lhs_true1945:
	v798 = *libc.As[int32](lookahead)
	cmp1946 = v798 != 34
	if cmp1946 {
		goto land_lhs_true1948
	} else {
		goto if_end1958
	}

land_lhs_true1948:
	v799 = *libc.As[int32](lookahead)
	cmp1949 = v799 != 35
	if cmp1949 {
		goto land_lhs_true1951
	} else {
		goto if_end1958
	}

land_lhs_true1951:
	v800 = *libc.As[int32](lookahead)
	cmp1952 = v800 != 59
	if cmp1952 {
		goto land_lhs_true1954
	} else {
		goto if_end1958
	}

land_lhs_true1954:
	v801 = *libc.As[int32](lookahead)
	cmp1955 = v801 != 92
	if cmp1955 {
		goto if_then1957
	} else {
		goto if_end1958
	}

if_then1957:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1958:
	v802 = *libc.As[byte](result)
	loadedv1959 = (v802 & 1) != 0
	*libc.As[bool](retval) = loadedv1959
	goto _return

sw_bb1960:
	*libc.As[byte](result) = 1
	v803 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1961 = libc.Ptr(&libc.As[TSLexer](v803).F1)
	*libc.As[int16](result_symbol1961) = 17
	v804 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1962 = libc.Ptr(&libc.As[TSLexer](v804).F3)
	v805 = *libc.As[unsafe.Pointer](mark_end1962)
	v806 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v805)(v806)
	v807 = *libc.As[int32](lookahead)
	cmp1963 = v807 == 114
	if cmp1963 {
		goto if_then1965
	} else {
		goto if_end1966
	}

if_then1965:
	*libc.As[int16](state_addr) = 71
	goto next_state

if_end1966:
	v808 = *libc.As[int32](lookahead)
	cmp1967 = v808 != 0
	if cmp1967 {
		goto land_lhs_true1969
	} else {
		goto if_end1988
	}

land_lhs_true1969:
	v809 = *libc.As[int32](lookahead)
	cmp1970 = v809 != 10
	if cmp1970 {
		goto land_lhs_true1972
	} else {
		goto if_end1988
	}

land_lhs_true1972:
	v810 = *libc.As[int32](lookahead)
	cmp1973 = v810 != 13
	if cmp1973 {
		goto land_lhs_true1975
	} else {
		goto if_end1988
	}

land_lhs_true1975:
	v811 = *libc.As[int32](lookahead)
	cmp1976 = v811 != 34
	if cmp1976 {
		goto land_lhs_true1978
	} else {
		goto if_end1988
	}

land_lhs_true1978:
	v812 = *libc.As[int32](lookahead)
	cmp1979 = v812 != 35
	if cmp1979 {
		goto land_lhs_true1981
	} else {
		goto if_end1988
	}

land_lhs_true1981:
	v813 = *libc.As[int32](lookahead)
	cmp1982 = v813 != 59
	if cmp1982 {
		goto land_lhs_true1984
	} else {
		goto if_end1988
	}

land_lhs_true1984:
	v814 = *libc.As[int32](lookahead)
	cmp1985 = v814 != 92
	if cmp1985 {
		goto if_then1987
	} else {
		goto if_end1988
	}

if_then1987:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end1988:
	v815 = *libc.As[byte](result)
	loadedv1989 = (v815 & 1) != 0
	*libc.As[bool](retval) = loadedv1989
	goto _return

sw_bb1990:
	*libc.As[byte](result) = 1
	v816 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1991 = libc.Ptr(&libc.As[TSLexer](v816).F1)
	*libc.As[int16](result_symbol1991) = 17
	v817 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1992 = libc.Ptr(&libc.As[TSLexer](v817).F3)
	v818 = *libc.As[unsafe.Pointer](mark_end1992)
	v819 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v818)(v819)
	v820 = *libc.As[int32](lookahead)
	cmp1993 = v820 == 115
	if cmp1993 {
		goto if_then1995
	} else {
		goto if_end1996
	}

if_then1995:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end1996:
	v821 = *libc.As[int32](lookahead)
	cmp1997 = v821 != 0
	if cmp1997 {
		goto land_lhs_true1999
	} else {
		goto if_end2018
	}

land_lhs_true1999:
	v822 = *libc.As[int32](lookahead)
	cmp2000 = v822 != 10
	if cmp2000 {
		goto land_lhs_true2002
	} else {
		goto if_end2018
	}

land_lhs_true2002:
	v823 = *libc.As[int32](lookahead)
	cmp2003 = v823 != 13
	if cmp2003 {
		goto land_lhs_true2005
	} else {
		goto if_end2018
	}

land_lhs_true2005:
	v824 = *libc.As[int32](lookahead)
	cmp2006 = v824 != 34
	if cmp2006 {
		goto land_lhs_true2008
	} else {
		goto if_end2018
	}

land_lhs_true2008:
	v825 = *libc.As[int32](lookahead)
	cmp2009 = v825 != 35
	if cmp2009 {
		goto land_lhs_true2011
	} else {
		goto if_end2018
	}

land_lhs_true2011:
	v826 = *libc.As[int32](lookahead)
	cmp2012 = v826 != 59
	if cmp2012 {
		goto land_lhs_true2014
	} else {
		goto if_end2018
	}

land_lhs_true2014:
	v827 = *libc.As[int32](lookahead)
	cmp2015 = v827 != 92
	if cmp2015 {
		goto if_then2017
	} else {
		goto if_end2018
	}

if_then2017:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end2018:
	v828 = *libc.As[byte](result)
	loadedv2019 = (v828 & 1) != 0
	*libc.As[bool](retval) = loadedv2019
	goto _return

sw_bb2020:
	*libc.As[byte](result) = 1
	v829 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2021 = libc.Ptr(&libc.As[TSLexer](v829).F1)
	*libc.As[int16](result_symbol2021) = 17
	v830 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2022 = libc.Ptr(&libc.As[TSLexer](v830).F3)
	v831 = *libc.As[unsafe.Pointer](mark_end2022)
	v832 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v831)(v832)
	v833 = *libc.As[int32](lookahead)
	cmp2023 = v833 == 115
	if cmp2023 {
		goto if_then2025
	} else {
		goto if_end2026
	}

if_then2025:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end2026:
	v834 = *libc.As[int32](lookahead)
	cmp2027 = v834 != 0
	if cmp2027 {
		goto land_lhs_true2029
	} else {
		goto if_end2048
	}

land_lhs_true2029:
	v835 = *libc.As[int32](lookahead)
	cmp2030 = v835 != 10
	if cmp2030 {
		goto land_lhs_true2032
	} else {
		goto if_end2048
	}

land_lhs_true2032:
	v836 = *libc.As[int32](lookahead)
	cmp2033 = v836 != 13
	if cmp2033 {
		goto land_lhs_true2035
	} else {
		goto if_end2048
	}

land_lhs_true2035:
	v837 = *libc.As[int32](lookahead)
	cmp2036 = v837 != 34
	if cmp2036 {
		goto land_lhs_true2038
	} else {
		goto if_end2048
	}

land_lhs_true2038:
	v838 = *libc.As[int32](lookahead)
	cmp2039 = v838 != 35
	if cmp2039 {
		goto land_lhs_true2041
	} else {
		goto if_end2048
	}

land_lhs_true2041:
	v839 = *libc.As[int32](lookahead)
	cmp2042 = v839 != 59
	if cmp2042 {
		goto land_lhs_true2044
	} else {
		goto if_end2048
	}

land_lhs_true2044:
	v840 = *libc.As[int32](lookahead)
	cmp2045 = v840 != 92
	if cmp2045 {
		goto if_then2047
	} else {
		goto if_end2048
	}

if_then2047:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end2048:
	v841 = *libc.As[byte](result)
	loadedv2049 = (v841 & 1) != 0
	*libc.As[bool](retval) = loadedv2049
	goto _return

sw_bb2050:
	*libc.As[byte](result) = 1
	v842 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2051 = libc.Ptr(&libc.As[TSLexer](v842).F1)
	*libc.As[int16](result_symbol2051) = 17
	v843 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2052 = libc.Ptr(&libc.As[TSLexer](v843).F3)
	v844 = *libc.As[unsafe.Pointer](mark_end2052)
	v845 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v844)(v845)
	v846 = *libc.As[int32](lookahead)
	cmp2053 = v846 == 117
	if cmp2053 {
		goto if_then2055
	} else {
		goto if_end2056
	}

if_then2055:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end2056:
	v847 = *libc.As[int32](lookahead)
	cmp2057 = v847 != 0
	if cmp2057 {
		goto land_lhs_true2059
	} else {
		goto if_end2078
	}

land_lhs_true2059:
	v848 = *libc.As[int32](lookahead)
	cmp2060 = v848 != 10
	if cmp2060 {
		goto land_lhs_true2062
	} else {
		goto if_end2078
	}

land_lhs_true2062:
	v849 = *libc.As[int32](lookahead)
	cmp2063 = v849 != 13
	if cmp2063 {
		goto land_lhs_true2065
	} else {
		goto if_end2078
	}

land_lhs_true2065:
	v850 = *libc.As[int32](lookahead)
	cmp2066 = v850 != 34
	if cmp2066 {
		goto land_lhs_true2068
	} else {
		goto if_end2078
	}

land_lhs_true2068:
	v851 = *libc.As[int32](lookahead)
	cmp2069 = v851 != 35
	if cmp2069 {
		goto land_lhs_true2071
	} else {
		goto if_end2078
	}

land_lhs_true2071:
	v852 = *libc.As[int32](lookahead)
	cmp2072 = v852 != 59
	if cmp2072 {
		goto land_lhs_true2074
	} else {
		goto if_end2078
	}

land_lhs_true2074:
	v853 = *libc.As[int32](lookahead)
	cmp2075 = v853 != 92
	if cmp2075 {
		goto if_then2077
	} else {
		goto if_end2078
	}

if_then2077:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end2078:
	v854 = *libc.As[byte](result)
	loadedv2079 = (v854 & 1) != 0
	*libc.As[bool](retval) = loadedv2079
	goto _return

sw_bb2080:
	*libc.As[byte](result) = 1
	v855 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2081 = libc.Ptr(&libc.As[TSLexer](v855).F1)
	*libc.As[int16](result_symbol2081) = 17
	v856 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2082 = libc.Ptr(&libc.As[TSLexer](v856).F3)
	v857 = *libc.As[unsafe.Pointer](mark_end2082)
	v858 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v857)(v858)
	v859 = *libc.As[int32](lookahead)
	cmp2083 = v859 != 0
	if cmp2083 {
		goto land_lhs_true2085
	} else {
		goto if_end2104
	}

land_lhs_true2085:
	v860 = *libc.As[int32](lookahead)
	cmp2086 = v860 != 10
	if cmp2086 {
		goto land_lhs_true2088
	} else {
		goto if_end2104
	}

land_lhs_true2088:
	v861 = *libc.As[int32](lookahead)
	cmp2089 = v861 != 13
	if cmp2089 {
		goto land_lhs_true2091
	} else {
		goto if_end2104
	}

land_lhs_true2091:
	v862 = *libc.As[int32](lookahead)
	cmp2092 = v862 != 34
	if cmp2092 {
		goto land_lhs_true2094
	} else {
		goto if_end2104
	}

land_lhs_true2094:
	v863 = *libc.As[int32](lookahead)
	cmp2095 = v863 != 35
	if cmp2095 {
		goto land_lhs_true2097
	} else {
		goto if_end2104
	}

land_lhs_true2097:
	v864 = *libc.As[int32](lookahead)
	cmp2098 = v864 != 59
	if cmp2098 {
		goto land_lhs_true2100
	} else {
		goto if_end2104
	}

land_lhs_true2100:
	v865 = *libc.As[int32](lookahead)
	cmp2101 = v865 != 92
	if cmp2101 {
		goto if_then2103
	} else {
		goto if_end2104
	}

if_then2103:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end2104:
	v866 = *libc.As[byte](result)
	loadedv2105 = (v866 & 1) != 0
	*libc.As[bool](retval) = loadedv2105
	goto _return

sw_bb2106:
	*libc.As[byte](result) = 1
	v867 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2107 = libc.Ptr(&libc.As[TSLexer](v867).F1)
	*libc.As[int16](result_symbol2107) = 18
	v868 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2108 = libc.Ptr(&libc.As[TSLexer](v868).F3)
	v869 = *libc.As[unsafe.Pointer](mark_end2108)
	v870 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v869)(v870)
	v871 = *libc.As[byte](result)
	loadedv2109 = (v871 & 1) != 0
	*libc.As[bool](retval) = loadedv2109
	goto _return

sw_bb2110:
	*libc.As[byte](result) = 1
	v872 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2111 = libc.Ptr(&libc.As[TSLexer](v872).F1)
	*libc.As[int16](result_symbol2111) = 19
	v873 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2112 = libc.Ptr(&libc.As[TSLexer](v873).F3)
	v874 = *libc.As[unsafe.Pointer](mark_end2112)
	v875 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v874)(v875)
	v876 = *libc.As[byte](result)
	loadedv2113 = (v876 & 1) != 0
	*libc.As[bool](retval) = loadedv2113
	goto _return

sw_bb2114:
	*libc.As[byte](result) = 1
	v877 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2115 = libc.Ptr(&libc.As[TSLexer](v877).F1)
	*libc.As[int16](result_symbol2115) = 20
	v878 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2116 = libc.Ptr(&libc.As[TSLexer](v878).F3)
	v879 = *libc.As[unsafe.Pointer](mark_end2116)
	v880 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v879)(v880)
	*libc.As[int32](i2117) = 0
	goto for_cond2118

for_cond2118:
	v881 = *libc.As[int32](i2117)
	conv2119 = int64(uint64(uint32(v881)))
	cmp2120 = uint64(conv2119) < uint64(18)
	if cmp2120 {
		goto for_body2122
	} else {
		goto for_end2135
	}

for_body2122:
	v882 = *libc.As[int32](i2117)
	idxprom2123 = int64(uint64(uint32(v882)))
	arrayidx2124 = libc.Ptr(&ts_lex_map_45[idxprom2123])
	v883 = *libc.As[int16](arrayidx2124)
	conv2125 = int32(uint32(uint16(v883)))
	v884 = *libc.As[int32](lookahead)
	cmp2126 = conv2125 == v884
	if cmp2126 {
		goto if_then2128
	} else {
		goto if_end2132
	}

if_then2128:
	v885 = *libc.As[int32](i2117)
	add2129 = v885 + 1
	idxprom2130 = int64(uint64(uint32(add2129)))
	arrayidx2131 = libc.Ptr(&ts_lex_map_45[idxprom2130])
	v886 = *libc.As[int16](arrayidx2131)
	*libc.As[int16](state_addr) = v886
	goto next_state

if_end2132:
	goto for_inc2133

for_inc2133:
	v887 = *libc.As[int32](i2117)
	add2134 = v887 + 2
	*libc.As[int32](i2117) = add2134
	goto for_cond2118

for_end2135:
	v888 = *libc.As[byte](result)
	loadedv2136 = (v888 & 1) != 0
	*libc.As[bool](retval) = loadedv2136
	goto _return

sw_bb2137:
	*libc.As[byte](result) = 1
	v889 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2138 = libc.Ptr(&libc.As[TSLexer](v889).F1)
	*libc.As[int16](result_symbol2138) = 21
	v890 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2139 = libc.Ptr(&libc.As[TSLexer](v890).F3)
	v891 = *libc.As[unsafe.Pointer](mark_end2139)
	v892 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v891)(v892)
	v893 = *libc.As[byte](result)
	loadedv2140 = (v893 & 1) != 0
	*libc.As[bool](retval) = loadedv2140
	goto _return

sw_bb2141:
	*libc.As[byte](result) = 1
	v894 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2142 = libc.Ptr(&libc.As[TSLexer](v894).F1)
	*libc.As[int16](result_symbol2142) = 22
	v895 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2143 = libc.Ptr(&libc.As[TSLexer](v895).F3)
	v896 = *libc.As[unsafe.Pointer](mark_end2143)
	v897 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v896)(v897)
	v898 = *libc.As[int32](lookahead)
	cmp2144 = v898 == 9
	if cmp2144 {
		goto if_then2155
	} else {
		goto lor_lhs_false2146
	}

lor_lhs_false2146:
	v899 = *libc.As[int32](lookahead)
	cmp2147 = v899 == 11
	if cmp2147 {
		goto if_then2155
	} else {
		goto lor_lhs_false2149
	}

lor_lhs_false2149:
	v900 = *libc.As[int32](lookahead)
	cmp2150 = v900 == 12
	if cmp2150 {
		goto if_then2155
	} else {
		goto lor_lhs_false2152
	}

lor_lhs_false2152:
	v901 = *libc.As[int32](lookahead)
	cmp2153 = v901 == 32
	if cmp2153 {
		goto if_then2155
	} else {
		goto if_end2156
	}

if_then2155:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end2156:
	v902 = *libc.As[int32](lookahead)
	cmp2157 = v902 != 0
	if cmp2157 {
		goto land_lhs_true2159
	} else {
		goto if_end2166
	}

land_lhs_true2159:
	v903 = *libc.As[int32](lookahead)
	cmp2160 = v903 < 9
	if cmp2160 {
		goto if_then2165
	} else {
		goto lor_lhs_false2162
	}

lor_lhs_false2162:
	v904 = *libc.As[int32](lookahead)
	cmp2163 = 13 < v904
	if cmp2163 {
		goto if_then2165
	} else {
		goto if_end2166
	}

if_then2165:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end2166:
	v905 = *libc.As[byte](result)
	loadedv2167 = (v905 & 1) != 0
	*libc.As[bool](retval) = loadedv2167
	goto _return

sw_bb2168:
	*libc.As[byte](result) = 1
	v906 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2169 = libc.Ptr(&libc.As[TSLexer](v906).F1)
	*libc.As[int16](result_symbol2169) = 22
	v907 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2170 = libc.Ptr(&libc.As[TSLexer](v907).F3)
	v908 = *libc.As[unsafe.Pointer](mark_end2170)
	v909 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v908)(v909)
	v910 = *libc.As[int32](lookahead)
	cmp2171 = v910 != 0
	if cmp2171 {
		goto land_lhs_true2173
	} else {
		goto if_end2180
	}

land_lhs_true2173:
	v911 = *libc.As[int32](lookahead)
	cmp2174 = v911 != 10
	if cmp2174 {
		goto land_lhs_true2176
	} else {
		goto if_end2180
	}

land_lhs_true2176:
	v912 = *libc.As[int32](lookahead)
	cmp2177 = v912 != 13
	if cmp2177 {
		goto if_then2179
	} else {
		goto if_end2180
	}

if_then2179:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end2180:
	v913 = *libc.As[byte](result)
	loadedv2181 = (v913 & 1) != 0
	*libc.As[bool](retval) = loadedv2181
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v914 = *libc.As[bool](retval)
	return v914
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
