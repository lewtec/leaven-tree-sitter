package grammar_objdump

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

var tree_sitter_objdump_language struct {
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
}
var ts_small_parse_table [582]int16 = [582]int16{6, 13, 1, 0, 15, 1, 18, 18, 1, 21, 21, 1, 23, 24, 1, 24, 2, 8, 31, 32, 33, 34, 36, 43, 46, 47, 6, 5, 1, 18, 7, 1, 21, 9, 1, 23, 11, 1, 24, 27, 1, 0, 2, 8, 31, 32, 33, 34, 36, 43, 46, 47, 8, 31, 1, 5, 33, 1, 11, 35, 1, 12, 8, 1, 42, 18, 1, 45, 20, 1, 39, 29, 2, 0, 21, 37, 3, 18, 23, 24, 5, 41, 1, 18, 44, 1, 23, 47, 1, 24, 39, 2, 0, 21, 5, 4, 34, 36, 43, 48, 5, 52, 1, 18, 55, 1, 23, 58, 1, 24, 50, 2, 0, 21, 5, 4, 34, 36, 43, 48, 5, 63, 1, 18, 66, 1, 23, 69, 1, 24, 61, 2, 0, 21, 5, 4, 34, 36, 43, 48, 4, 31, 1, 5, 16, 1, 45, 72, 2, 0, 21, 74, 3, 18, 23, 24, 4, 11, 1, 24, 76, 1, 18, 78, 1, 23, 7, 4, 34, 36, 43, 48, 4, 31, 1, 5, 26, 1, 45, 80, 2, 0, 21, 82, 3, 18, 23, 24, 4, 86, 1, 5, 27, 1, 35, 84, 2, 0, 21, 88, 3, 18, 23, 24, 4, 11, 1, 24, 76, 1, 18, 78, 1, 23, 6, 4, 34, 36, 43, 48, 2, 90, 3, 0, 1, 21, 92, 3, 18, 23, 24, 2, 94, 3, 0, 5, 21, 96, 3, 18, 23, 24, 2, 98, 3, 0, 5, 21, 100, 3, 18, 23, 24, 2, 102, 2, 0, 21, 104, 3, 18, 23, 24, 2, 106, 2, 0, 21, 108, 3, 18, 23, 24, 2, 72, 2, 0, 21, 74, 3, 18, 23, 24, 2, 110, 2, 0, 21, 112, 3, 18, 23, 24, 2, 114, 2, 0, 21, 116, 3, 18, 23, 24, 2, 118, 2, 0, 21, 120, 3, 18, 23, 24, 2, 122, 2, 0, 21, 124, 3, 18, 23, 24, 2, 126, 2, 0, 21, 128, 3, 18, 23, 24, 2, 130, 2, 0, 21, 132, 3, 18, 23, 24, 2, 134, 2, 0, 21, 136, 3, 18, 23, 24, 2, 138, 2, 0, 21, 140, 3, 18, 23, 24, 2, 142, 2, 0, 21, 144, 3, 18, 23, 24, 3, 146, 1, 15, 148, 1, 18, 17, 2, 40, 41, 4, 150, 1, 8, 152, 1, 17, 155, 1, 28, 29, 1, 49, 4, 157, 1, 8, 159, 1, 17, 161, 1, 28, 31, 1, 49, 4, 159, 1, 17, 163, 1, 8, 165, 1, 28, 29, 1, 49, 3, 167, 1, 9, 169, 1, 10, 24, 2, 37, 38, 3, 31, 1, 5, 171, 1, 1, 44, 1, 45, 2, 155, 1, 28, 150, 2, 8, 17, 2, 173, 1, 13, 175, 1, 14, 2, 177, 1, 8, 179, 1, 28, 2, 181, 1, 1, 183, 1, 4, 2, 185, 1, 16, 36, 1, 44, 2, 187, 1, 2, 189, 1, 22, 2, 35, 1, 12, 10, 1, 42, 1, 191, 1, 20, 1, 193, 1, 15, 1, 195, 1, 22, 1, 197, 1, 1, 1, 199, 1, 27, 1, 201, 1, 19, 1, 181, 1, 1, 1, 203, 1, 6, 1, 205, 1, 7, 1, 207, 1, 7, 1, 209, 1, 16, 1, 211, 1, 3, 1, 213, 1, 1, 1, 215, 1, 1, 1, 217, 1, 25, 1, 189, 1, 22, 1, 219, 1, 0, 1, 221, 1, 15, 1, 223, 1, 1, 1, 225, 1, 1, 1, 227, 1, 14, 1, 229, 1, 26}
var ts_small_parse_table_map [61]int32 = [61]int32{0, 26, 52, 80, 100, 120, 140, 156, 172, 188, 204, 220, 231, 242, 253, 263, 273, 283, 293, 303, 313, 323, 333, 343, 353, 363, 373, 384, 397, 410, 423, 434, 444, 452, 459, 466, 473, 480, 487, 494, 498, 502, 506, 510, 514, 518, 522, 526, 530, 534, 538, 542, 546, 550, 554, 558, 562, 566, 570, 574, 578}
var ts_symbol_names [51]unsafe.Pointer = [51]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26), libc.Ptr(&_str_6), libc.Ptr(&_str_6), libc.Ptr(&_str_6), libc.Ptr(&_str_27), libc.Ptr(&_str_28), libc.Ptr(&_str_29), libc.Ptr(&_str_30), libc.Ptr(&_str_31), libc.Ptr(&_str_32), libc.Ptr(&_str_33), libc.Ptr(&_str_8), libc.Ptr(&_str_34), libc.Ptr(&_str_35), libc.Ptr(&_str_36), libc.Ptr(&_str_37), libc.Ptr(&_str_38), libc.Ptr(&_str_39), libc.Ptr(&_str_40), libc.Ptr(&_str_41), libc.Ptr(&_str_42), libc.Ptr(&_str_43), libc.Ptr(&_str_44), libc.Ptr(&_str_45), libc.Ptr(&_str_46), libc.Ptr(&_str_47), libc.Ptr(&_str_48)}
var ts_symbol_metadata [51]TSSymbolMetadata = [51]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}}
var ts_symbol_map [51]int16 = [51]int16{0, 1, 2, 3, 25, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 25, 25, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [2][5]int16 = [2][5]int16{[5]int16{}, [5]int16{50, 0, 0, 0, 0}}
var ts_lex_modes [63]TSLexMode = [63]TSLexMode{TSLexMode{0, 1}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{8, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{8, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{65, 0}, TSLexMode{7, 0}, TSLexMode{75, 2}, TSLexMode{75, 2}, TSLexMode{75, 2}, TSLexMode{1, 0}, TSLexMode{65, 0}, TSLexMode{75, 2}, TSLexMode{}, TSLexMode{76, 2}, TSLexMode{8, 0}, TSLexMode{11, 0}, TSLexMode{28, 0}, TSLexMode{65, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{28, 0}, TSLexMode{}, TSLexMode{0, 3}, TSLexMode{11, 0}, TSLexMode{}, TSLexMode{11, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{11, 0}, TSLexMode{11, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{2, 0}, TSLexMode{28, 0}, TSLexMode{}, TSLexMode{7, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{61, 0}}
var ts_external_scanner_states [4][3]byte = [4][3]byte{[3]byte{}, [3]byte{1, 1, 1}, [3]byte{0, 1, 0}, [3]byte{1, 0, 0}}
var ts_external_scanner_symbol_map [3]int16 = [3]int16{27, 28, 29}
var ts_primary_state_ids [63]int16 = [63]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62}
var __const_scan_code_identifier_next_token_text [13]byte = [13]byte{40, 70, 105, 108, 101, 79, 102, 102, 115, 101, 116, 58, 0}
var ts_parse_table struct {
	F0 struct {
		F0 [30]int16
		F1 [20]int16
	}
	F1 [50]int16
} = struct {
	F0 struct {
		F0 [30]int16
		F1 [20]int16
	}
	F1 [50]int16
}{struct {
	F0 [30]int16
	F1 [20]int16
}{[30]int16{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1}, [20]int16{}}, [50]int16{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 7, 0, 9, 11, 0, 0, 0, 0, 0, 57, 3, 3, 3, 3, 0, 3, 0, 0, 0, 0, 0, 0, 3, 0, 0, 3, 3, 0, 0}}
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F24 struct {
		F0 anon_2
		F1 [6]byte
	}
	F25 TSParseActionEntry
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
	F30 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F47 struct {
		F0 anon_2
		F1 [6]byte
	}
	F48 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
	F56 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 struct {
			F0 struct {
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
	F75 TSParseActionEntry
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
	F85 TSParseActionEntry
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
	F91 TSParseActionEntry
	F92 struct {
		F0 anon_2
		F1 [6]byte
	}
	F93 TSParseActionEntry
	F94 struct {
		F0 anon_2
		F1 [6]byte
	}
	F95 TSParseActionEntry
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
	F115 TSParseActionEntry
	F116 struct {
		F0 anon_2
		F1 [6]byte
	}
	F117 TSParseActionEntry
	F118 struct {
		F0 anon_2
		F1 [6]byte
	}
	F119 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F125 TSParseActionEntry
	F126 struct {
		F0 anon_2
		F1 [6]byte
	}
	F127 TSParseActionEntry
	F128 struct {
		F0 anon_2
		F1 [6]byte
	}
	F129 TSParseActionEntry
	F130 struct {
		F0 anon_2
		F1 [6]byte
	}
	F131 TSParseActionEntry
	F132 struct {
		F0 anon_2
		F1 [6]byte
	}
	F133 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F139 TSParseActionEntry
	F140 struct {
		F0 anon_2
		F1 [6]byte
	}
	F141 TSParseActionEntry
	F142 struct {
		F0 anon_2
		F1 [6]byte
	}
	F143 TSParseActionEntry
	F144 struct {
		F0 anon_2
		F1 [6]byte
	}
	F145 TSParseActionEntry
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
	F149 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F150 struct {
		F0 anon_2
		F1 [6]byte
	}
	F151 TSParseActionEntry
	F152 struct {
		F0 anon_2
		F1 [6]byte
	}
	F153 TSParseActionEntry
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
	F156 TSParseActionEntry
	F157 struct {
		F0 anon_2
		F1 [6]byte
	}
	F158 TSParseActionEntry
	F159 struct {
		F0 anon_2
		F1 [6]byte
	}
	F160 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F161 struct {
		F0 anon_2
		F1 [6]byte
	}
	F162 TSParseActionEntry
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
	F173 struct {
		F0 anon_2
		F1 [6]byte
	}
	F174 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F175 struct {
		F0 anon_2
		F1 [6]byte
	}
	F176 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F177 struct {
		F0 anon_2
		F1 [6]byte
	}
	F178 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F179 struct {
		F0 anon_2
		F1 [6]byte
	}
	F180 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F181 struct {
		F0 anon_2
		F1 [6]byte
	}
	F182 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F183 struct {
		F0 anon_2
		F1 [6]byte
	}
	F184 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F185 struct {
		F0 anon_2
		F1 [6]byte
	}
	F186 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F187 struct {
		F0 anon_2
		F1 [6]byte
	}
	F188 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F189 struct {
		F0 anon_2
		F1 [6]byte
	}
	F190 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F191 struct {
		F0 anon_2
		F1 [6]byte
	}
	F192 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F193 struct {
		F0 anon_2
		F1 [6]byte
	}
	F194 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F195 struct {
		F0 anon_2
		F1 [6]byte
	}
	F196 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F197 struct {
		F0 anon_2
		F1 [6]byte
	}
	F198 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F199 struct {
		F0 anon_2
		F1 [6]byte
	}
	F200 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F201 struct {
		F0 anon_2
		F1 [6]byte
	}
	F202 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F203 struct {
		F0 anon_2
		F1 [6]byte
	}
	F204 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F205 struct {
		F0 anon_2
		F1 [6]byte
	}
	F206 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F207 struct {
		F0 anon_2
		F1 [6]byte
	}
	F208 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F209 struct {
		F0 anon_2
		F1 [6]byte
	}
	F210 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F211 struct {
		F0 anon_2
		F1 [6]byte
	}
	F212 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F213 struct {
		F0 anon_2
		F1 [6]byte
	}
	F214 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F215 struct {
		F0 anon_2
		F1 [6]byte
	}
	F216 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F217 struct {
		F0 anon_2
		F1 [6]byte
	}
	F218 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F219 struct {
		F0 anon_2
		F1 [6]byte
	}
	F220 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F221 struct {
		F0 anon_2
		F1 [6]byte
	}
	F222 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F223 struct {
		F0 anon_2
		F1 [6]byte
	}
	F224 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F225 struct {
		F0 anon_2
		F1 [6]byte
	}
	F226 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F227 struct {
		F0 anon_2
		F1 [6]byte
	}
	F228 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F229 struct {
		F0 anon_2
		F1 [6]byte
	}
	F230 struct {
		F0 struct {
			F0 struct {
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F24 struct {
		F0 anon_2
		F1 [6]byte
	}
	F25 TSParseActionEntry
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
	F30 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F47 struct {
		F0 anon_2
		F1 [6]byte
	}
	F48 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
	F56 TSParseActionEntry
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
		F0 struct {
			F0 struct {
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
		F0 struct {
			F0 struct {
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
	F75 TSParseActionEntry
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
	F85 TSParseActionEntry
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
	F91 TSParseActionEntry
	F92 struct {
		F0 anon_2
		F1 [6]byte
	}
	F93 TSParseActionEntry
	F94 struct {
		F0 anon_2
		F1 [6]byte
	}
	F95 TSParseActionEntry
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
	F115 TSParseActionEntry
	F116 struct {
		F0 anon_2
		F1 [6]byte
	}
	F117 TSParseActionEntry
	F118 struct {
		F0 anon_2
		F1 [6]byte
	}
	F119 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F125 TSParseActionEntry
	F126 struct {
		F0 anon_2
		F1 [6]byte
	}
	F127 TSParseActionEntry
	F128 struct {
		F0 anon_2
		F1 [6]byte
	}
	F129 TSParseActionEntry
	F130 struct {
		F0 anon_2
		F1 [6]byte
	}
	F131 TSParseActionEntry
	F132 struct {
		F0 anon_2
		F1 [6]byte
	}
	F133 TSParseActionEntry
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
		F0 anon_2
		F1 [6]byte
	}
	F139 TSParseActionEntry
	F140 struct {
		F0 anon_2
		F1 [6]byte
	}
	F141 TSParseActionEntry
	F142 struct {
		F0 anon_2
		F1 [6]byte
	}
	F143 TSParseActionEntry
	F144 struct {
		F0 anon_2
		F1 [6]byte
	}
	F145 TSParseActionEntry
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
	F149 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F150 struct {
		F0 anon_2
		F1 [6]byte
	}
	F151 TSParseActionEntry
	F152 struct {
		F0 anon_2
		F1 [6]byte
	}
	F153 TSParseActionEntry
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
	F156 TSParseActionEntry
	F157 struct {
		F0 anon_2
		F1 [6]byte
	}
	F158 TSParseActionEntry
	F159 struct {
		F0 anon_2
		F1 [6]byte
	}
	F160 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F161 struct {
		F0 anon_2
		F1 [6]byte
	}
	F162 TSParseActionEntry
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
	F173 struct {
		F0 anon_2
		F1 [6]byte
	}
	F174 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F175 struct {
		F0 anon_2
		F1 [6]byte
	}
	F176 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F177 struct {
		F0 anon_2
		F1 [6]byte
	}
	F178 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F179 struct {
		F0 anon_2
		F1 [6]byte
	}
	F180 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F181 struct {
		F0 anon_2
		F1 [6]byte
	}
	F182 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F183 struct {
		F0 anon_2
		F1 [6]byte
	}
	F184 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F185 struct {
		F0 anon_2
		F1 [6]byte
	}
	F186 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F187 struct {
		F0 anon_2
		F1 [6]byte
	}
	F188 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F189 struct {
		F0 anon_2
		F1 [6]byte
	}
	F190 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F191 struct {
		F0 anon_2
		F1 [6]byte
	}
	F192 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F193 struct {
		F0 anon_2
		F1 [6]byte
	}
	F194 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F195 struct {
		F0 anon_2
		F1 [6]byte
	}
	F196 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F197 struct {
		F0 anon_2
		F1 [6]byte
	}
	F198 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F199 struct {
		F0 anon_2
		F1 [6]byte
	}
	F200 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F201 struct {
		F0 anon_2
		F1 [6]byte
	}
	F202 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F203 struct {
		F0 anon_2
		F1 [6]byte
	}
	F204 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F205 struct {
		F0 anon_2
		F1 [6]byte
	}
	F206 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F207 struct {
		F0 anon_2
		F1 [6]byte
	}
	F208 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F209 struct {
		F0 anon_2
		F1 [6]byte
	}
	F210 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F211 struct {
		F0 anon_2
		F1 [6]byte
	}
	F212 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F213 struct {
		F0 anon_2
		F1 [6]byte
	}
	F214 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F215 struct {
		F0 anon_2
		F1 [6]byte
	}
	F216 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F217 struct {
		F0 anon_2
		F1 [6]byte
	}
	F218 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F219 struct {
		F0 anon_2
		F1 [6]byte
	}
	F220 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F221 struct {
		F0 anon_2
		F1 [6]byte
	}
	F222 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F223 struct {
		F0 anon_2
		F1 [6]byte
	}
	F224 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F225 struct {
		F0 anon_2
		F1 [6]byte
	}
	F226 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F227 struct {
		F0 anon_2
		F1 [6]byte
	}
	F228 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F229 struct {
		F0 anon_2
		F1 [6]byte
	}
	F230 struct {
		F0 struct {
			F0 struct {
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 30, 0, 0}}}, struct {
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
}{0, 0, 37, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 62, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 60, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 59, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 47, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 47, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 37, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 47, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 62, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 47, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 60, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 47, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 59, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 30, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 38, 0, 0}}}, struct {
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
}{0, 0, 46, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 45, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 38, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 48, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 48, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 48, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 53, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 48, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 59, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 33, 0, 1}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 33, 0, 1}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 33, 0, 1}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 53, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 33, 0, 1}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 59, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 33, 0, 1}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 33, 0, 1}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 33, 0, 1}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 53, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 33, 0, 1}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 59, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 38, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 38, 0, 0}}}, struct {
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
}{0, 0, 47, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 53, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 40, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 40, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 34, 0, 0}}}, struct {
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
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 45, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 45, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 42, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 42, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 42, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 42, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 38, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 38, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 39, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 39, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 43, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 43, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 37, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 37, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 35, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 35, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 32, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 32, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 46, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 46, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 36, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 5, 36, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 36, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 36, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 40, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 40, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 34, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 34, 0, 0}}}, struct {
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
}{0, 0, 40, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 49, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 49, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 51, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 49, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 44, 0, 0}}}, struct {
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
}{0, 0, 51, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 44, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 44, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 44, 0, 0}}}, struct {
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
}{0, 0, 58, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 33, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 49, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 50, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 35, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 34, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 55, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 61, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 39, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 54, 0, 0}, [2]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [2]byte = [2]byte{58, 0}
var _str_4 [5]byte = [5]byte{102, 105, 108, 101, 0}
var _str_5 [7]byte = [7]byte{102, 111, 114, 109, 97, 116, 0}
var _str_6 [11]byte = [11]byte{105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}
var _str_7 [2]byte = [2]byte{40, 0}
var _str_8 [14]byte = [14]byte{100, 105, 115, 99, 114, 105, 109, 105, 110, 97, 116, 111, 114, 0}
var _str_9 [2]byte = [2]byte{41, 0}
var _str_10 [21]byte = [21]byte{109, 101, 109, 111, 114, 121, 95, 111, 102, 102, 115, 101, 116, 95, 116, 111, 107, 101, 110, 49, 0}
var _str_11 [12]byte = [12]byte{105, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 0}
var _str_12 [16]byte = [16]byte{98, 97, 100, 95, 105, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 0}
var _str_13 [2]byte = [2]byte{35, 0}
var _str_14 [2]byte = [2]byte{60, 0}
var _str_15 [2]byte = [2]byte{43, 0}
var _str_16 [2]byte = [2]byte{62, 0}
var _str_17 [12]byte = [12]byte{104, 101, 120, 97, 100, 101, 99, 105, 109, 97, 108, 0}
var _str_18 [5]byte = [5]byte{98, 121, 116, 101, 0}
var _str_19 [2]byte = [2]byte{32, 0}
var _str_20 [8]byte = [8]byte{97, 100, 100, 114, 101, 115, 115, 0}
var _str_21 [5]byte = [5]byte{70, 105, 108, 101, 0}
var _str_22 [8]byte = [8]byte{79, 102, 102, 115, 101, 116, 58, 0}
var _str_23 [24]byte = [24]byte{68, 105, 115, 97, 115, 115, 101, 109, 98, 108, 121, 32, 111, 102, 32, 115, 101, 99, 116, 105, 111, 110, 32, 0}
var _str_24 [8]byte = [8]byte{105, 110, 116, 101, 103, 101, 114, 0}
var _str_25 [10]byte = [10]byte{102, 105, 108, 101, 95, 112, 97, 116, 104, 0}
var _str_26 [6]byte = [6]byte{108, 97, 98, 101, 108, 0}
var _str_27 [23]byte = [23]byte{95, 119, 104, 105, 116, 101, 115, 112, 97, 99, 101, 95, 110, 111, 95, 110, 101, 119, 108, 105, 110, 101, 0}
var _str_28 [16]byte = [16]byte{95, 101, 114, 114, 111, 114, 95, 115, 101, 110, 116, 105, 110, 101, 108, 0}
var _str_29 [7]byte = [7]byte{115, 111, 117, 114, 99, 101, 0}
var _str_30 [6]byte = [6]byte{95, 108, 105, 110, 101, 0}
var _str_31 [7]byte = [7]byte{104, 101, 97, 100, 101, 114, 0}
var _str_32 [20]byte = [20]byte{100, 105, 115, 97, 115, 115, 101, 109, 98, 108, 121, 95, 115, 101, 99, 116, 105, 111, 110, 0}
var _str_33 [16]byte = [16]byte{115, 111, 117, 114, 99, 101, 95, 108, 111, 99, 97, 116, 105, 111, 110, 0}
var _str_34 [14]byte = [14]byte{109, 101, 109, 111, 114, 121, 95, 111, 102, 102, 115, 101, 116, 0}
var _str_35 [25]byte = [25]byte{95, 105, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 95, 97, 110, 100, 95, 99, 111, 109, 109, 101, 110, 116, 0}
var _str_36 [26]byte = [26]byte{95, 105, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 95, 97, 110, 100, 95, 108, 111, 99, 97, 116, 105, 111, 110, 0}
var _str_37 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_38 [20]byte = [20]byte{95, 99, 111, 109, 109, 101, 110, 116, 95, 119, 105, 116, 104, 95, 108, 97, 98, 101, 108, 0}
var _str_39 [22]byte = [22]byte{95, 99, 111, 109, 109, 101, 110, 116, 95, 119, 105, 116, 104, 95, 97, 100, 100, 114, 101, 115, 115, 0}
var _str_40 [14]byte = [14]byte{99, 111, 100, 101, 95, 108, 111, 99, 97, 116, 105, 111, 110, 0}
var _str_41 [11]byte = [11]byte{108, 97, 98, 101, 108, 95, 108, 105, 110, 101, 0}
var _str_42 [19]byte = [19]byte{109, 97, 99, 104, 105, 110, 101, 95, 99, 111, 100, 101, 95, 98, 121, 116, 101, 115, 0}
var _str_43 [12]byte = [12]byte{102, 105, 108, 101, 95, 111, 102, 102, 115, 101, 116, 0}
var _str_44 [26]byte = [26]byte{100, 105, 115, 97, 115, 115, 101, 109, 98, 108, 121, 95, 115, 101, 99, 116, 105, 111, 110, 95, 108, 97, 98, 101, 108, 0}
var _str_45 [15]byte = [15]byte{115, 111, 117, 114, 99, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_46 [28]byte = [28]byte{100, 105, 115, 97, 115, 115, 101, 109, 98, 108, 121, 95, 115, 101, 99, 116, 105, 111, 110, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_47 [27]byte = [27]byte{109, 97, 99, 104, 105, 110, 101, 95, 99, 111, 100, 101, 95, 98, 121, 116, 101, 115, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_48 [16]byte = [16]byte{115, 101, 99, 116, 105, 111, 110, 95, 97, 100, 100, 114, 101, 115, 115, 0}

func init() {
	tree_sitter_objdump_language = struct {
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
	}{14, 50, 1, 30, 3, 63, 2, 2, 0, 5, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{libc.Ptr(&ts_external_scanner_states), libc.Ptr(&ts_external_scanner_symbol_map), libc.FuncCode(tree_sitter_objdump_external_scanner_create), libc.FuncCode(tree_sitter_objdump_external_scanner_destroy), libc.FuncCode(tree_sitter_objdump_external_scanner_scan), libc.FuncCode(tree_sitter_objdump_external_scanner_serialize), libc.FuncCode(tree_sitter_objdump_external_scanner_deserialize)}, libc.Ptr(&ts_primary_state_ids)}
}
func tree_sitter_objdump_external_scanner_create() unsafe.Pointer {
	return nil
}
func tree_sitter_objdump_external_scanner_deserialize(payload unsafe.Pointer, buffer unsafe.Pointer, length int32) {
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
func tree_sitter_objdump_external_scanner_destroy(payload unsafe.Pointer) {
	var payload_addr unsafe.Pointer
	_ = payload_addr

	payload_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	*libc.As[unsafe.Pointer](payload_addr) = payload
}
func tree_sitter_objdump_external_scanner_scan(payload unsafe.Pointer, lexer unsafe.Pointer, valid_symbols unsafe.Pointer) bool {
	var loadedv, loadedv2, call, loadedv6, call8, v8 bool
	var retval unsafe.Pointer
	var v1, v3, v6 byte
	var arrayidx, arrayidx1, arrayidx5 unsafe.Pointer
	var v0, v2, v4, v5, v7 unsafe.Pointer
	var payload_addr, lexer_addr, valid_symbols_addr unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, payload_addr, lexer_addr, valid_symbols_addr, v0, arrayidx, v1, loadedv, v2, arrayidx1, v3, loadedv2, v4, call, v5, arrayidx5, v6, loadedv6, v7, call8, v8

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
	*libc.As[unsafe.Pointer](payload_addr) = payload
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[unsafe.Pointer](valid_symbols_addr) = valid_symbols
	v0 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v0), int(int64(2))*1))
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
	v2 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx1 = libc.Ptr(libc.AddPointer[byte](libc.As[byte](v2), int(int64(1))*1))
	v3 = *libc.As[byte](arrayidx1)
	loadedv2 = (v3 & 1) != 0
	if loadedv2 {
		goto if_then3
	} else {
		goto if_end4
	}

if_then3:
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	call = scan_whitespace_no_newline(v4)
	*libc.As[bool](retval) = call
	goto _return

if_end4:
	v5 = *libc.As[unsafe.Pointer](valid_symbols_addr)
	arrayidx5 = v5
	v6 = *libc.As[byte](arrayidx5)
	loadedv6 = (v6 & 1) != 0
	if loadedv6 {
		goto if_then7
	} else {
		goto if_end9
	}

if_then7:
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	call8 = scan_code_identifier(v7)
	*libc.As[bool](retval) = call8
	goto _return

if_end9:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v8 = *libc.As[bool](retval)
	return v8
}
func scan_whitespace_no_newline(lexer unsafe.Pointer) bool {
	var call, loadedv, v16 bool
	var retval unsafe.Pointer
	var result_symbol unsafe.Pointer
	var v8 int32
	var lookahead unsafe.Pointer
	var v6 byte
	var has_text unsafe.Pointer
	var v0, v1, v2, v3, v4, v5, v7, v9, v10, v11, v12, v13, v14, v15 unsafe.Pointer
	var lexer_addr, mark_end, eof, mark_end2, advance unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, has_text, v0, mark_end, v1, v2, v3, eof, v4, v5, call, v6, loadedv, v7, lookahead, v8, v9, result_symbol, v10, mark_end2, v11, v12, v13, advance, v14, v15, v16

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
	has_text = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v0).F3)
	v1 = *libc.As[unsafe.Pointer](mark_end)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v1)(v2)
	*libc.As[byte](has_text) = 0
	goto while_body

while_body:
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	eof = libc.Ptr(&libc.As[TSLexer](v3).F6)
	v4 = *libc.As[unsafe.Pointer](eof)
	v5 = *libc.As[unsafe.Pointer](lexer_addr)
	call = libc.FuncFromCode[func(unsafe.Pointer) bool](v4)(v5)
	if call {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v6 = *libc.As[byte](has_text)
	loadedv = (v6 & 1) != 0
	*libc.As[bool](retval) = loadedv
	goto _return

if_end:
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v7).F0)
	v8 = *libc.As[int32](lookahead)
	switch v8 {
	case 10:
		goto sw_bb
	case 32:
		goto sw_bb1
	case 9:
		goto sw_bb1
	default:
		goto sw_default
	}

sw_bb:
	*libc.As[bool](retval) = true
	goto _return

sw_bb1:
	*libc.As[byte](has_text) = 1
	v9 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v9).F1)
	*libc.As[int16](result_symbol) = 1
	v10 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2 = libc.Ptr(&libc.As[TSLexer](v10).F3)
	v11 = *libc.As[unsafe.Pointer](mark_end2)
	v12 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v11)(v12)
	goto sw_epilog

sw_default:
	*libc.As[bool](retval) = false
	goto _return

sw_epilog:
	v13 = *libc.As[unsafe.Pointer](lexer_addr)
	advance = libc.Ptr(&libc.As[TSLexer](v13).F2)
	v14 = *libc.As[unsafe.Pointer](advance)
	v15 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v14)(v15, false)
	goto while_body

_return:
	v16 = *libc.As[bool](retval)
	return v16
}
func scan_code_identifier(lexer unsafe.Pointer) bool {
	var next_token_text unsafe.Pointer
	var cmp, call, cmp2, tobool, loadedv, call9, loadedv13, cmp16, cmp23, cmp26, loadedv36, loadedv38, v37 bool
	var retval unsafe.Pointer
	var result_symbol, result_symbol29 unsafe.Pointer
	var v4, v10, v12, call4, v15, v18, v19, inc, v21, v22, conv22, v24, add, v26, inc31, v28 int32
	var offset_counter, size, lookahead, lookahead1, lookahead3, lookahead8, lookahead15, lookahead21, lookahead34 unsafe.Pointer
	var idxprom int64
	var v13, conv, v16, v23, v29, v30 byte
	var has_text, has_hexadecimal_data, possibly_in_next_hexadecimal_token, possibly_in_next_file_offset_token, arrayidx unsafe.Pointer
	var v0, v1, v2, v3, v5, v6, v7, v8, v9, v11, v14, v17, v20, v25, v27, v31, v32, v33, v34, v35, v36 unsafe.Pointer
	var lexer_addr, advance, eof, mark_end, mark_end42 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, has_text, offset_counter, has_hexadecimal_data, possibly_in_next_hexadecimal_token, possibly_in_next_file_offset_token, next_token_text, size, v0, advance, v1, v2, v3, lookahead, v4, cmp, v5, eof, v6, v7, call, v8, result_symbol, v9, lookahead1, v10, cmp2, v11, lookahead3, v12, call4, tobool, v13, loadedv, v14, lookahead8, v15, conv, call9, v16, loadedv13, v17, lookahead15, v18, cmp16, v19, inc, v20, lookahead21, v21, v22, idxprom, arrayidx, v23, conv22, cmp23, v24, add, cmp26, v25, result_symbol29, v26, inc31, v27, lookahead34, v28, v29, loadedv36, v30, loadedv38, v31, mark_end, v32, v33, v34, mark_end42, v35, v36, v37

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
	has_text = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	offset_counter = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	has_hexadecimal_data = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	possibly_in_next_hexadecimal_token = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	possibly_in_next_file_offset_token = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	next_token_text = libc.Ptr(&new(struct {
		_ [0]uint64
		v [13]byte
		b byte
	}).v)
	size = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[byte](has_text) = 0
	*libc.As[int32](offset_counter) = -1
	*libc.As[byte](has_hexadecimal_data) = 0
	*libc.As[byte](possibly_in_next_hexadecimal_token) = 0
	*libc.As[byte](possibly_in_next_file_offset_token) = 0
	libc.Memmove(libc.As[byte](next_token_text), libc.As[byte](libc.Ptr(&__const_scan_code_identifier_next_token_text)), int64(13))
	*libc.As[int32](size) = 12
	goto while_body

while_body:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, false)
	v3 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead = libc.Ptr(&libc.As[TSLexer](v3).F0)
	v4 = *libc.As[int32](lookahead)
	cmp = v4 == 10
	if cmp {
		goto if_then
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v5 = *libc.As[unsafe.Pointer](lexer_addr)
	eof = libc.Ptr(&libc.As[TSLexer](v5).F6)
	v6 = *libc.As[unsafe.Pointer](eof)
	v7 = *libc.As[unsafe.Pointer](lexer_addr)
	call = libc.FuncFromCode[func(unsafe.Pointer) bool](v6)(v7)
	if call {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v8).F1)
	*libc.As[int16](result_symbol) = 0
	*libc.As[bool](retval) = true
	goto _return

if_end:
	v9 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v9).F0)
	v10 = *libc.As[int32](lookahead1)
	cmp2 = v10 != 10
	if cmp2 {
		goto land_lhs_true
	} else {
		goto if_end6
	}

land_lhs_true:
	v11 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead3 = libc.Ptr(&libc.As[TSLexer](v11).F0)
	v12 = *libc.As[int32](lookahead3)
	call4 = libc.Iswspace(v12)
	tobool = call4 != 0
	if tobool {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	goto while_body

if_end6:
	v13 = *libc.As[byte](possibly_in_next_hexadecimal_token)
	loadedv = (v13 & 1) != 0
	if loadedv {
		goto if_then7
	} else {
		goto if_end12
	}

if_then7:
	v14 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead8 = libc.Ptr(&libc.As[TSLexer](v14).F0)
	v15 = *libc.As[int32](lookahead8)
	conv = byte(v15)
	call9 = is_hexadecimal_character(conv)
	if call9 {
		goto if_then10
	} else {
		goto if_else
	}

if_then10:
	*libc.As[byte](has_hexadecimal_data) = 1
	goto if_end11

if_else:
	*libc.As[byte](possibly_in_next_hexadecimal_token) = 0
	goto if_end11

if_end11:
	goto if_end12

if_end12:
	*libc.As[byte](has_text) = 1
	v16 = *libc.As[byte](possibly_in_next_file_offset_token)
	loadedv13 = (v16 & 1) != 0
	if loadedv13 {
		goto if_else20
	} else {
		goto if_then14
	}

if_then14:
	v17 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead15 = libc.Ptr(&libc.As[TSLexer](v17).F0)
	v18 = *libc.As[int32](lookahead15)
	cmp16 = v18 == 40
	if cmp16 {
		goto if_then18
	} else {
		goto if_end19
	}

if_then18:
	*libc.As[byte](possibly_in_next_file_offset_token) = 1
	v19 = *libc.As[int32](offset_counter)
	inc = v19 + 1
	*libc.As[int32](offset_counter) = inc
	goto while_body

if_end19:
	goto if_end33

if_else20:
	v20 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead21 = libc.Ptr(&libc.As[TSLexer](v20).F0)
	v21 = *libc.As[int32](lookahead21)
	v22 = *libc.As[int32](offset_counter)
	idxprom = int64(uint64(uint32(v22)))
	arrayidx = libc.Ptr(&libc.As[[13]byte](next_token_text)[idxprom])
	v23 = *libc.As[byte](arrayidx)
	conv22 = int32(int8(v23))
	cmp23 = v21 == conv22
	if cmp23 {
		goto if_then25
	} else {
		goto if_else32
	}

if_then25:
	v24 = *libc.As[int32](offset_counter)
	add = v24 + 1
	cmp26 = uint32(add) >= 12
	if cmp26 {
		goto if_then28
	} else {
		goto if_end30
	}

if_then28:
	v25 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol29 = libc.Ptr(&libc.As[TSLexer](v25).F1)
	*libc.As[int16](result_symbol29) = 0
	*libc.As[bool](retval) = true
	goto _return

if_end30:
	v26 = *libc.As[int32](offset_counter)
	inc31 = v26 + 1
	*libc.As[int32](offset_counter) = inc31
	goto while_body

if_else32:
	*libc.As[byte](possibly_in_next_file_offset_token) = 0
	goto while_body

if_end33:
	v27 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead34 = libc.Ptr(&libc.As[TSLexer](v27).F0)
	v28 = *libc.As[int32](lookahead34)
	switch v28 {
	case 10:
		goto sw_bb
	case 62:
		goto sw_bb35
	case 43:
		goto sw_bb41
	default:
		goto sw_epilog
	}

sw_bb:
	*libc.As[bool](retval) = false
	goto _return

sw_bb35:
	v29 = *libc.As[byte](has_hexadecimal_data)
	loadedv36 = (v29 & 1) != 0
	if loadedv36 {
		goto if_end40
	} else {
		goto land_lhs_true37
	}

land_lhs_true37:
	v30 = *libc.As[byte](possibly_in_next_hexadecimal_token)
	loadedv38 = (v30 & 1) != 0
	if loadedv38 {
		goto if_end40
	} else {
		goto if_then39
	}

if_then39:
	v31 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v31).F3)
	v32 = *libc.As[unsafe.Pointer](mark_end)
	v33 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v32)(v33)
	goto if_end40

if_end40:
	goto sw_epilog

sw_bb41:
	v34 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end42 = libc.Ptr(&libc.As[TSLexer](v34).F3)
	v35 = *libc.As[unsafe.Pointer](mark_end42)
	v36 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v35)(v36)
	*libc.As[byte](possibly_in_next_hexadecimal_token) = 1
	goto sw_epilog

sw_epilog:
	goto while_body

_return:
	v37 = *libc.As[bool](retval)
	return v37
}
func tree_sitter_objdump_external_scanner_serialize(payload unsafe.Pointer, buffer unsafe.Pointer) int32 {
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
func tree_sitter_objdump() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_objdump_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp27, cmp31, cmp35, cmp39, cmp43, cmp47, cmp51, cmp53, cmp56, cmp59, cmp63, cmp65, cmp68, cmp71, cmp75, cmp78, loadedv82, cmp84, cmp88, cmp92, cmp95, cmp98, cmp102, cmp105, cmp108, loadedv112, cmp114, cmp118, cmp121, cmp124, cmp128, loadedv132, cmp134, loadedv138, cmp140, loadedv144, cmp146, loadedv150, cmp152, loadedv156, cmp158, cmp162, cmp165, cmp168, cmp171, cmp175, cmp178, cmp181, cmp184, cmp187, cmp190, loadedv194, cmp196, cmp200, cmp204, cmp208, cmp211, cmp214, cmp218, cmp221, cmp224, cmp227, cmp231, cmp234, cmp238, cmp241, cmp244, cmp247, cmp251, cmp254, cmp257, cmp260, cmp263, cmp266, loadedv270, cmp272, loadedv276, cmp278, cmp282, cmp285, loadedv289, cmp291, cmp295, cmp299, cmp303, cmp306, cmp309, cmp312, cmp316, cmp319, cmp322, cmp325, cmp328, cmp331, loadedv335, cmp337, loadedv341, cmp343, loadedv347, cmp349, loadedv353, cmp355, loadedv359, cmp361, loadedv365, cmp367, loadedv371, cmp373, loadedv377, cmp379, loadedv383, cmp385, loadedv389, cmp391, loadedv395, cmp397, loadedv401, cmp403, loadedv407, cmp409, loadedv413, cmp415, loadedv419, cmp421, loadedv425, cmp427, loadedv431, cmp433, cmp437, cmp440, cmp443, cmp446, cmp450, cmp453, loadedv457, cmp459, cmp463, cmp466, cmp469, cmp472, cmp475, cmp478, loadedv482, cmp484, loadedv488, cmp490, cmp494, cmp497, cmp500, cmp503, cmp506, cmp509, loadedv513, cmp515, loadedv519, cmp521, loadedv525, cmp527, loadedv531, cmp533, loadedv537, cmp539, loadedv543, cmp545, loadedv549, cmp551, loadedv555, cmp557, loadedv561, cmp563, loadedv567, cmp569, loadedv573, cmp575, loadedv579, cmp581, cmp585, cmp588, cmp591, cmp594, cmp597, cmp600, loadedv604, cmp606, loadedv610, cmp612, loadedv616, cmp618, loadedv622, cmp624, loadedv628, cmp630, loadedv634, cmp636, loadedv640, cmp642, loadedv646, cmp648, loadedv652, cmp654, loadedv658, cmp660, loadedv664, cmp666, loadedv670, cmp672, loadedv676, cmp678, loadedv682, cmp684, loadedv688, cmp690, loadedv694, cmp696, loadedv700, cmp702, loadedv706, cmp708, cmp711, cmp714, cmp717, cmp721, cmp724, loadedv728, cmp730, cmp733, cmp736, cmp739, cmp742, cmp745, loadedv749, cmp751, cmp754, cmp757, cmp760, cmp763, cmp766, loadedv770, cmp772, cmp775, loadedv779, loadedv781, cmp784, cmp788, cmp792, cmp796, cmp800, cmp804, cmp808, cmp811, cmp814, cmp818, cmp821, cmp824, cmp827, cmp831, cmp834, cmp838, cmp841, cmp844, cmp847, cmp851, cmp854, cmp857, cmp860, cmp863, cmp866, loadedv870, loadedv872, loadedv876, loadedv880, loadedv884, cmp888, cmp892, cmp895, loadedv899, loadedv903, cmp907, loadedv911, loadedv915, loadedv919, cmp923, cmp927, cmp930, cmp933, loadedv937, cmp941, cmp944, cmp947, cmp950, loadedv954, cmp958, cmp962, cmp965, cmp968, cmp972, cmp975, cmp978, cmp981, loadedv985, cmp989, cmp993, cmp996, cmp999, cmp1002, loadedv1006, cmp1010, cmp1014, cmp1017, cmp1020, cmp1023, loadedv1027, cmp1031, cmp1035, cmp1038, cmp1041, cmp1044, loadedv1048, cmp1052, cmp1056, cmp1059, cmp1062, cmp1065, loadedv1069, cmp1073, cmp1076, cmp1079, cmp1082, loadedv1086, loadedv1090, cmp1094, cmp1097, cmp1100, cmp1103, loadedv1107, loadedv1111, loadedv1115, cmp1119, cmp1122, loadedv1126, loadedv1130, loadedv1134, cmp1138, cmp1141, cmp1144, cmp1147, cmp1150, cmp1153, loadedv1157, loadedv1161, cmp1165, cmp1169, cmp1172, cmp1175, loadedv1179, cmp1183, cmp1187, cmp1190, cmp1193, cmp1196, cmp1199, cmp1202, loadedv1206, cmp1210, cmp1214, cmp1217, cmp1220, cmp1223, cmp1226, cmp1229, loadedv1233, cmp1237, cmp1241, cmp1244, cmp1247, cmp1251, cmp1254, cmp1257, cmp1261, cmp1264, cmp1267, cmp1270, cmp1273, cmp1276, cmp1280, cmp1283, cmp1286, cmp1289, cmp1292, cmp1295, loadedv1299, cmp1303, cmp1307, cmp1310, cmp1313, cmp1316, cmp1319, cmp1322, loadedv1326, cmp1330, cmp1334, cmp1338, cmp1341, cmp1344, cmp1347, cmp1350, cmp1353, loadedv1357, cmp1361, cmp1364, cmp1368, cmp1371, cmp1374, cmp1377, cmp1380, cmp1383, loadedv1387, cmp1391, cmp1394, cmp1397, cmp1401, cmp1404, cmp1407, cmp1411, cmp1414, cmp1417, cmp1420, cmp1423, cmp1426, cmp1430, cmp1433, cmp1436, cmp1439, cmp1442, cmp1445, loadedv1449, cmp1453, cmp1456, cmp1460, cmp1463, cmp1466, cmp1469, loadedv1473, cmp1477, cmp1480, cmp1483, cmp1486, cmp1489, cmp1492, loadedv1496, cmp1500, cmp1503, cmp1506, cmp1509, cmp1512, cmp1515, cmp1519, cmp1522, cmp1525, cmp1528, cmp1531, cmp1534, cmp1537, cmp1540, loadedv1544, loadedv1548, loadedv1552, loadedv1556, cmp1560, cmp1563, loadedv1567, cmp1571, cmp1575, cmp1578, cmp1581, cmp1585, cmp1588, cmp1591, cmp1595, cmp1598, cmp1601, cmp1604, cmp1607, cmp1610, cmp1613, loadedv1617, cmp1621, cmp1625, cmp1628, cmp1631, cmp1635, cmp1638, cmp1641, cmp1645, cmp1648, cmp1651, cmp1654, cmp1657, cmp1660, cmp1663, loadedv1667, cmp1671, cmp1675, cmp1678, cmp1681, cmp1685, cmp1688, cmp1691, cmp1695, cmp1698, cmp1701, cmp1704, cmp1707, cmp1710, cmp1713, loadedv1717, cmp1721, cmp1725, cmp1728, cmp1731, cmp1735, cmp1738, cmp1741, cmp1745, cmp1748, cmp1751, cmp1754, cmp1757, cmp1760, cmp1763, loadedv1767, cmp1771, cmp1775, cmp1778, cmp1781, cmp1785, cmp1788, cmp1791, cmp1795, cmp1798, cmp1801, cmp1804, cmp1807, cmp1810, cmp1813, loadedv1817, cmp1821, cmp1825, cmp1828, cmp1831, cmp1835, cmp1838, cmp1841, cmp1845, cmp1848, cmp1851, cmp1854, cmp1857, cmp1860, cmp1863, loadedv1867, cmp1871, cmp1875, cmp1878, cmp1881, cmp1885, cmp1888, cmp1891, cmp1895, cmp1898, cmp1901, cmp1904, cmp1907, cmp1910, cmp1913, loadedv1917, cmp1921, cmp1925, cmp1928, cmp1931, cmp1935, cmp1938, cmp1941, cmp1945, cmp1948, cmp1951, cmp1954, cmp1957, cmp1960, cmp1963, loadedv1967, cmp1971, cmp1975, cmp1978, cmp1981, cmp1985, cmp1988, cmp1991, cmp1995, cmp1998, cmp2001, cmp2004, cmp2007, cmp2010, cmp2013, loadedv2017, cmp2021, cmp2025, cmp2028, cmp2031, cmp2035, cmp2038, cmp2041, cmp2045, cmp2048, cmp2051, cmp2054, cmp2057, cmp2060, cmp2063, loadedv2067, cmp2071, cmp2074, cmp2077, cmp2081, cmp2084, cmp2087, cmp2091, cmp2094, cmp2097, cmp2100, cmp2103, cmp2106, cmp2109, loadedv2113, cmp2117, cmp2120, cmp2123, cmp2126, cmp2129, cmp2132, cmp2135, cmp2138, loadedv2142, cmp2146, cmp2149, cmp2152, cmp2155, cmp2158, cmp2161, cmp2164, cmp2167, cmp2170, cmp2173, loadedv2177, cmp2181, cmp2184, cmp2187, cmp2191, cmp2194, loadedv2198, cmp2202, cmp2205, loadedv2209, cmp2213, cmp2216, cmp2219, cmp2222, cmp2226, cmp2229, loadedv2233, cmp2237, cmp2240, loadedv2244, v923 bool
	var retval unsafe.Pointer
	var v9 int16
	var state_addr, result_symbol, result_symbol874, result_symbol878, result_symbol882, result_symbol886, result_symbol901, result_symbol905, result_symbol913, result_symbol917, result_symbol921, result_symbol939, result_symbol956, result_symbol987, result_symbol1008, result_symbol1029, result_symbol1050, result_symbol1071, result_symbol1088, result_symbol1092, result_symbol1109, result_symbol1113, result_symbol1117, result_symbol1128, result_symbol1132, result_symbol1136, result_symbol1159, result_symbol1163, result_symbol1181, result_symbol1208, result_symbol1235, result_symbol1301, result_symbol1328, result_symbol1359, result_symbol1389, result_symbol1451, result_symbol1475, result_symbol1498, result_symbol1546, result_symbol1550, result_symbol1554, result_symbol1558, result_symbol1569, result_symbol1619, result_symbol1669, result_symbol1719, result_symbol1769, result_symbol1819, result_symbol1869, result_symbol1919, result_symbol1969, result_symbol2019, result_symbol2069, result_symbol2115, result_symbol2144, result_symbol2179, result_symbol2200, result_symbol2211, result_symbol2235 unsafe.Pointer
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v34, v35, v36, v37, v38, v39, v40, v41, v43, v44, v45, v46, v47, v49, v51, v53, v55, v57, v58, v59, v60, v61, v62, v63, v64, v65, v66, v67, v69, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v90, v92, v94, v95, v96, v98, v99, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v110, v112, v114, v116, v118, v120, v122, v124, v126, v128, v130, v132, v134, v136, v138, v140, v142, v144, v145, v146, v147, v148, v149, v150, v152, v153, v154, v155, v156, v157, v158, v160, v162, v163, v164, v165, v166, v167, v168, v170, v172, v174, v176, v178, v180, v182, v184, v186, v188, v190, v192, v193, v194, v195, v196, v197, v198, v200, v202, v204, v206, v208, v210, v212, v214, v216, v218, v220, v222, v224, v226, v228, v230, v232, v234, v235, v236, v237, v238, v239, v241, v242, v243, v244, v245, v246, v248, v249, v250, v251, v252, v253, v255, v256, v259, v260, v261, v262, v263, v264, v265, v266, v267, v268, v269, v270, v271, v272, v273, v274, v275, v276, v277, v278, v279, v280, v281, v282, v283, v309, v310, v311, v322, v338, v339, v340, v341, v347, v348, v349, v350, v356, v357, v358, v359, v360, v361, v362, v363, v369, v370, v371, v372, v373, v379, v380, v381, v382, v383, v389, v390, v391, v392, v393, v399, v400, v401, v402, v403, v409, v410, v411, v412, v423, v424, v425, v426, v442, v443, v459, v460, v461, v462, v463, v464, v475, v476, v477, v478, v484, v485, v486, v487, v488, v489, v490, v496, v497, v498, v499, v500, v501, v502, v508, v509, v510, v511, v512, v513, v514, v515, v516, v517, v518, v519, v520, v521, v522, v523, v524, v525, v526, v532, v533, v534, v535, v536, v537, v538, v544, v545, v546, v547, v548, v549, v550, v551, v557, v558, v559, v560, v561, v562, v563, v564, v570, v571, v572, v573, v574, v575, v576, v577, v578, v579, v580, v581, v582, v583, v584, v585, v586, v587, v593, v594, v595, v596, v597, v598, v604, v605, v606, v607, v608, v609, v615, v616, v617, v618, v619, v620, v621, v622, v623, v624, v625, v626, v627, v628, v649, v650, v656, v657, v658, v659, v660, v661, v662, v663, v664, v665, v666, v667, v668, v669, v675, v676, v677, v678, v679, v680, v681, v682, v683, v684, v685, v686, v687, v688, v694, v695, v696, v697, v698, v699, v700, v701, v702, v703, v704, v705, v706, v707, v713, v714, v715, v716, v717, v718, v719, v720, v721, v722, v723, v724, v725, v726, v732, v733, v734, v735, v736, v737, v738, v739, v740, v741, v742, v743, v744, v745, v751, v752, v753, v754, v755, v756, v757, v758, v759, v760, v761, v762, v763, v764, v770, v771, v772, v773, v774, v775, v776, v777, v778, v779, v780, v781, v782, v783, v789, v790, v791, v792, v793, v794, v795, v796, v797, v798, v799, v800, v801, v802, v808, v809, v810, v811, v812, v813, v814, v815, v816, v817, v818, v819, v820, v821, v827, v828, v829, v830, v831, v832, v833, v834, v835, v836, v837, v838, v839, v840, v846, v847, v848, v849, v850, v851, v852, v853, v854, v855, v856, v857, v858, v864, v865, v866, v867, v868, v869, v870, v871, v877, v878, v879, v880, v881, v882, v883, v884, v885, v886, v892, v893, v894, v895, v896, v902, v903, v909, v910, v911, v912, v913, v914, v920, v921 int32
	var lookahead, lookahead1 unsafe.Pointer
	var v3, storedv, v10, v33, v42, v48, v50, v52, v54, v56, v68, v91, v93, v97, v111, v113, v115, v117, v119, v121, v123, v125, v127, v129, v131, v133, v135, v137, v139, v141, v143, v151, v159, v161, v169, v171, v173, v175, v177, v179, v181, v183, v185, v187, v189, v191, v199, v201, v203, v205, v207, v209, v211, v213, v215, v217, v219, v221, v223, v225, v227, v229, v231, v233, v240, v247, v254, v257, v258, v284, v289, v294, v299, v304, v312, v317, v323, v328, v333, v342, v351, v364, v374, v384, v394, v404, v413, v418, v427, v432, v437, v444, v449, v454, v465, v470, v479, v491, v503, v527, v539, v552, v565, v588, v599, v610, v629, v634, v639, v644, v651, v670, v689, v708, v727, v746, v765, v784, v803, v822, v841, v859, v872, v887, v897, v904, v915, v922 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v285, v286, v287, v288, v290, v291, v292, v293, v295, v296, v297, v298, v300, v301, v302, v303, v305, v306, v307, v308, v313, v314, v315, v316, v318, v319, v320, v321, v324, v325, v326, v327, v329, v330, v331, v332, v334, v335, v336, v337, v343, v344, v345, v346, v352, v353, v354, v355, v365, v366, v367, v368, v375, v376, v377, v378, v385, v386, v387, v388, v395, v396, v397, v398, v405, v406, v407, v408, v414, v415, v416, v417, v419, v420, v421, v422, v428, v429, v430, v431, v433, v434, v435, v436, v438, v439, v440, v441, v445, v446, v447, v448, v450, v451, v452, v453, v455, v456, v457, v458, v466, v467, v468, v469, v471, v472, v473, v474, v480, v481, v482, v483, v492, v493, v494, v495, v504, v505, v506, v507, v528, v529, v530, v531, v540, v541, v542, v543, v553, v554, v555, v556, v566, v567, v568, v569, v589, v590, v591, v592, v600, v601, v602, v603, v611, v612, v613, v614, v630, v631, v632, v633, v635, v636, v637, v638, v640, v641, v642, v643, v645, v646, v647, v648, v652, v653, v654, v655, v671, v672, v673, v674, v690, v691, v692, v693, v709, v710, v711, v712, v728, v729, v730, v731, v747, v748, v749, v750, v766, v767, v768, v769, v785, v786, v787, v788, v804, v805, v806, v807, v823, v824, v825, v826, v842, v843, v844, v845, v860, v861, v862, v863, v873, v874, v875, v876, v888, v889, v890, v891, v898, v899, v900, v901, v905, v906, v907, v908, v916, v917, v918, v919 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end875, mark_end879, mark_end883, mark_end887, mark_end902, mark_end906, mark_end914, mark_end918, mark_end922, mark_end940, mark_end957, mark_end988, mark_end1009, mark_end1030, mark_end1051, mark_end1072, mark_end1089, mark_end1093, mark_end1110, mark_end1114, mark_end1118, mark_end1129, mark_end1133, mark_end1137, mark_end1160, mark_end1164, mark_end1182, mark_end1209, mark_end1236, mark_end1302, mark_end1329, mark_end1360, mark_end1390, mark_end1452, mark_end1476, mark_end1499, mark_end1547, mark_end1551, mark_end1555, mark_end1559, mark_end1570, mark_end1620, mark_end1670, mark_end1720, mark_end1770, mark_end1820, mark_end1870, mark_end1920, mark_end1970, mark_end2020, mark_end2070, mark_end2116, mark_end2145, mark_end2180, mark_end2201, mark_end2212, mark_end2236 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp27, v18, cmp31, v19, cmp35, v20, cmp39, v21, cmp43, v22, cmp47, v23, cmp51, v24, cmp53, v25, cmp56, v26, cmp59, v27, cmp63, v28, cmp65, v29, cmp68, v30, cmp71, v31, cmp75, v32, cmp78, v33, loadedv82, v34, cmp84, v35, cmp88, v36, cmp92, v37, cmp95, v38, cmp98, v39, cmp102, v40, cmp105, v41, cmp108, v42, loadedv112, v43, cmp114, v44, cmp118, v45, cmp121, v46, cmp124, v47, cmp128, v48, loadedv132, v49, cmp134, v50, loadedv138, v51, cmp140, v52, loadedv144, v53, cmp146, v54, loadedv150, v55, cmp152, v56, loadedv156, v57, cmp158, v58, cmp162, v59, cmp165, v60, cmp168, v61, cmp171, v62, cmp175, v63, cmp178, v64, cmp181, v65, cmp184, v66, cmp187, v67, cmp190, v68, loadedv194, v69, cmp196, v70, cmp200, v71, cmp204, v72, cmp208, v73, cmp211, v74, cmp214, v75, cmp218, v76, cmp221, v77, cmp224, v78, cmp227, v79, cmp231, v80, cmp234, v81, cmp238, v82, cmp241, v83, cmp244, v84, cmp247, v85, cmp251, v86, cmp254, v87, cmp257, v88, cmp260, v89, cmp263, v90, cmp266, v91, loadedv270, v92, cmp272, v93, loadedv276, v94, cmp278, v95, cmp282, v96, cmp285, v97, loadedv289, v98, cmp291, v99, cmp295, v100, cmp299, v101, cmp303, v102, cmp306, v103, cmp309, v104, cmp312, v105, cmp316, v106, cmp319, v107, cmp322, v108, cmp325, v109, cmp328, v110, cmp331, v111, loadedv335, v112, cmp337, v113, loadedv341, v114, cmp343, v115, loadedv347, v116, cmp349, v117, loadedv353, v118, cmp355, v119, loadedv359, v120, cmp361, v121, loadedv365, v122, cmp367, v123, loadedv371, v124, cmp373, v125, loadedv377, v126, cmp379, v127, loadedv383, v128, cmp385, v129, loadedv389, v130, cmp391, v131, loadedv395, v132, cmp397, v133, loadedv401, v134, cmp403, v135, loadedv407, v136, cmp409, v137, loadedv413, v138, cmp415, v139, loadedv419, v140, cmp421, v141, loadedv425, v142, cmp427, v143, loadedv431, v144, cmp433, v145, cmp437, v146, cmp440, v147, cmp443, v148, cmp446, v149, cmp450, v150, cmp453, v151, loadedv457, v152, cmp459, v153, cmp463, v154, cmp466, v155, cmp469, v156, cmp472, v157, cmp475, v158, cmp478, v159, loadedv482, v160, cmp484, v161, loadedv488, v162, cmp490, v163, cmp494, v164, cmp497, v165, cmp500, v166, cmp503, v167, cmp506, v168, cmp509, v169, loadedv513, v170, cmp515, v171, loadedv519, v172, cmp521, v173, loadedv525, v174, cmp527, v175, loadedv531, v176, cmp533, v177, loadedv537, v178, cmp539, v179, loadedv543, v180, cmp545, v181, loadedv549, v182, cmp551, v183, loadedv555, v184, cmp557, v185, loadedv561, v186, cmp563, v187, loadedv567, v188, cmp569, v189, loadedv573, v190, cmp575, v191, loadedv579, v192, cmp581, v193, cmp585, v194, cmp588, v195, cmp591, v196, cmp594, v197, cmp597, v198, cmp600, v199, loadedv604, v200, cmp606, v201, loadedv610, v202, cmp612, v203, loadedv616, v204, cmp618, v205, loadedv622, v206, cmp624, v207, loadedv628, v208, cmp630, v209, loadedv634, v210, cmp636, v211, loadedv640, v212, cmp642, v213, loadedv646, v214, cmp648, v215, loadedv652, v216, cmp654, v217, loadedv658, v218, cmp660, v219, loadedv664, v220, cmp666, v221, loadedv670, v222, cmp672, v223, loadedv676, v224, cmp678, v225, loadedv682, v226, cmp684, v227, loadedv688, v228, cmp690, v229, loadedv694, v230, cmp696, v231, loadedv700, v232, cmp702, v233, loadedv706, v234, cmp708, v235, cmp711, v236, cmp714, v237, cmp717, v238, cmp721, v239, cmp724, v240, loadedv728, v241, cmp730, v242, cmp733, v243, cmp736, v244, cmp739, v245, cmp742, v246, cmp745, v247, loadedv749, v248, cmp751, v249, cmp754, v250, cmp757, v251, cmp760, v252, cmp763, v253, cmp766, v254, loadedv770, v255, cmp772, v256, cmp775, v257, loadedv779, v258, loadedv781, v259, cmp784, v260, cmp788, v261, cmp792, v262, cmp796, v263, cmp800, v264, cmp804, v265, cmp808, v266, cmp811, v267, cmp814, v268, cmp818, v269, cmp821, v270, cmp824, v271, cmp827, v272, cmp831, v273, cmp834, v274, cmp838, v275, cmp841, v276, cmp844, v277, cmp847, v278, cmp851, v279, cmp854, v280, cmp857, v281, cmp860, v282, cmp863, v283, cmp866, v284, loadedv870, v285, result_symbol, v286, mark_end, v287, v288, v289, loadedv872, v290, result_symbol874, v291, mark_end875, v292, v293, v294, loadedv876, v295, result_symbol878, v296, mark_end879, v297, v298, v299, loadedv880, v300, result_symbol882, v301, mark_end883, v302, v303, v304, loadedv884, v305, result_symbol886, v306, mark_end887, v307, v308, v309, cmp888, v310, cmp892, v311, cmp895, v312, loadedv899, v313, result_symbol901, v314, mark_end902, v315, v316, v317, loadedv903, v318, result_symbol905, v319, mark_end906, v320, v321, v322, cmp907, v323, loadedv911, v324, result_symbol913, v325, mark_end914, v326, v327, v328, loadedv915, v329, result_symbol917, v330, mark_end918, v331, v332, v333, loadedv919, v334, result_symbol921, v335, mark_end922, v336, v337, v338, cmp923, v339, cmp927, v340, cmp930, v341, cmp933, v342, loadedv937, v343, result_symbol939, v344, mark_end940, v345, v346, v347, cmp941, v348, cmp944, v349, cmp947, v350, cmp950, v351, loadedv954, v352, result_symbol956, v353, mark_end957, v354, v355, v356, cmp958, v357, cmp962, v358, cmp965, v359, cmp968, v360, cmp972, v361, cmp975, v362, cmp978, v363, cmp981, v364, loadedv985, v365, result_symbol987, v366, mark_end988, v367, v368, v369, cmp989, v370, cmp993, v371, cmp996, v372, cmp999, v373, cmp1002, v374, loadedv1006, v375, result_symbol1008, v376, mark_end1009, v377, v378, v379, cmp1010, v380, cmp1014, v381, cmp1017, v382, cmp1020, v383, cmp1023, v384, loadedv1027, v385, result_symbol1029, v386, mark_end1030, v387, v388, v389, cmp1031, v390, cmp1035, v391, cmp1038, v392, cmp1041, v393, cmp1044, v394, loadedv1048, v395, result_symbol1050, v396, mark_end1051, v397, v398, v399, cmp1052, v400, cmp1056, v401, cmp1059, v402, cmp1062, v403, cmp1065, v404, loadedv1069, v405, result_symbol1071, v406, mark_end1072, v407, v408, v409, cmp1073, v410, cmp1076, v411, cmp1079, v412, cmp1082, v413, loadedv1086, v414, result_symbol1088, v415, mark_end1089, v416, v417, v418, loadedv1090, v419, result_symbol1092, v420, mark_end1093, v421, v422, v423, cmp1094, v424, cmp1097, v425, cmp1100, v426, cmp1103, v427, loadedv1107, v428, result_symbol1109, v429, mark_end1110, v430, v431, v432, loadedv1111, v433, result_symbol1113, v434, mark_end1114, v435, v436, v437, loadedv1115, v438, result_symbol1117, v439, mark_end1118, v440, v441, v442, cmp1119, v443, cmp1122, v444, loadedv1126, v445, result_symbol1128, v446, mark_end1129, v447, v448, v449, loadedv1130, v450, result_symbol1132, v451, mark_end1133, v452, v453, v454, loadedv1134, v455, result_symbol1136, v456, mark_end1137, v457, v458, v459, cmp1138, v460, cmp1141, v461, cmp1144, v462, cmp1147, v463, cmp1150, v464, cmp1153, v465, loadedv1157, v466, result_symbol1159, v467, mark_end1160, v468, v469, v470, loadedv1161, v471, result_symbol1163, v472, mark_end1164, v473, v474, v475, cmp1165, v476, cmp1169, v477, cmp1172, v478, cmp1175, v479, loadedv1179, v480, result_symbol1181, v481, mark_end1182, v482, v483, v484, cmp1183, v485, cmp1187, v486, cmp1190, v487, cmp1193, v488, cmp1196, v489, cmp1199, v490, cmp1202, v491, loadedv1206, v492, result_symbol1208, v493, mark_end1209, v494, v495, v496, cmp1210, v497, cmp1214, v498, cmp1217, v499, cmp1220, v500, cmp1223, v501, cmp1226, v502, cmp1229, v503, loadedv1233, v504, result_symbol1235, v505, mark_end1236, v506, v507, v508, cmp1237, v509, cmp1241, v510, cmp1244, v511, cmp1247, v512, cmp1251, v513, cmp1254, v514, cmp1257, v515, cmp1261, v516, cmp1264, v517, cmp1267, v518, cmp1270, v519, cmp1273, v520, cmp1276, v521, cmp1280, v522, cmp1283, v523, cmp1286, v524, cmp1289, v525, cmp1292, v526, cmp1295, v527, loadedv1299, v528, result_symbol1301, v529, mark_end1302, v530, v531, v532, cmp1303, v533, cmp1307, v534, cmp1310, v535, cmp1313, v536, cmp1316, v537, cmp1319, v538, cmp1322, v539, loadedv1326, v540, result_symbol1328, v541, mark_end1329, v542, v543, v544, cmp1330, v545, cmp1334, v546, cmp1338, v547, cmp1341, v548, cmp1344, v549, cmp1347, v550, cmp1350, v551, cmp1353, v552, loadedv1357, v553, result_symbol1359, v554, mark_end1360, v555, v556, v557, cmp1361, v558, cmp1364, v559, cmp1368, v560, cmp1371, v561, cmp1374, v562, cmp1377, v563, cmp1380, v564, cmp1383, v565, loadedv1387, v566, result_symbol1389, v567, mark_end1390, v568, v569, v570, cmp1391, v571, cmp1394, v572, cmp1397, v573, cmp1401, v574, cmp1404, v575, cmp1407, v576, cmp1411, v577, cmp1414, v578, cmp1417, v579, cmp1420, v580, cmp1423, v581, cmp1426, v582, cmp1430, v583, cmp1433, v584, cmp1436, v585, cmp1439, v586, cmp1442, v587, cmp1445, v588, loadedv1449, v589, result_symbol1451, v590, mark_end1452, v591, v592, v593, cmp1453, v594, cmp1456, v595, cmp1460, v596, cmp1463, v597, cmp1466, v598, cmp1469, v599, loadedv1473, v600, result_symbol1475, v601, mark_end1476, v602, v603, v604, cmp1477, v605, cmp1480, v606, cmp1483, v607, cmp1486, v608, cmp1489, v609, cmp1492, v610, loadedv1496, v611, result_symbol1498, v612, mark_end1499, v613, v614, v615, cmp1500, v616, cmp1503, v617, cmp1506, v618, cmp1509, v619, cmp1512, v620, cmp1515, v621, cmp1519, v622, cmp1522, v623, cmp1525, v624, cmp1528, v625, cmp1531, v626, cmp1534, v627, cmp1537, v628, cmp1540, v629, loadedv1544, v630, result_symbol1546, v631, mark_end1547, v632, v633, v634, loadedv1548, v635, result_symbol1550, v636, mark_end1551, v637, v638, v639, loadedv1552, v640, result_symbol1554, v641, mark_end1555, v642, v643, v644, loadedv1556, v645, result_symbol1558, v646, mark_end1559, v647, v648, v649, cmp1560, v650, cmp1563, v651, loadedv1567, v652, result_symbol1569, v653, mark_end1570, v654, v655, v656, cmp1571, v657, cmp1575, v658, cmp1578, v659, cmp1581, v660, cmp1585, v661, cmp1588, v662, cmp1591, v663, cmp1595, v664, cmp1598, v665, cmp1601, v666, cmp1604, v667, cmp1607, v668, cmp1610, v669, cmp1613, v670, loadedv1617, v671, result_symbol1619, v672, mark_end1620, v673, v674, v675, cmp1621, v676, cmp1625, v677, cmp1628, v678, cmp1631, v679, cmp1635, v680, cmp1638, v681, cmp1641, v682, cmp1645, v683, cmp1648, v684, cmp1651, v685, cmp1654, v686, cmp1657, v687, cmp1660, v688, cmp1663, v689, loadedv1667, v690, result_symbol1669, v691, mark_end1670, v692, v693, v694, cmp1671, v695, cmp1675, v696, cmp1678, v697, cmp1681, v698, cmp1685, v699, cmp1688, v700, cmp1691, v701, cmp1695, v702, cmp1698, v703, cmp1701, v704, cmp1704, v705, cmp1707, v706, cmp1710, v707, cmp1713, v708, loadedv1717, v709, result_symbol1719, v710, mark_end1720, v711, v712, v713, cmp1721, v714, cmp1725, v715, cmp1728, v716, cmp1731, v717, cmp1735, v718, cmp1738, v719, cmp1741, v720, cmp1745, v721, cmp1748, v722, cmp1751, v723, cmp1754, v724, cmp1757, v725, cmp1760, v726, cmp1763, v727, loadedv1767, v728, result_symbol1769, v729, mark_end1770, v730, v731, v732, cmp1771, v733, cmp1775, v734, cmp1778, v735, cmp1781, v736, cmp1785, v737, cmp1788, v738, cmp1791, v739, cmp1795, v740, cmp1798, v741, cmp1801, v742, cmp1804, v743, cmp1807, v744, cmp1810, v745, cmp1813, v746, loadedv1817, v747, result_symbol1819, v748, mark_end1820, v749, v750, v751, cmp1821, v752, cmp1825, v753, cmp1828, v754, cmp1831, v755, cmp1835, v756, cmp1838, v757, cmp1841, v758, cmp1845, v759, cmp1848, v760, cmp1851, v761, cmp1854, v762, cmp1857, v763, cmp1860, v764, cmp1863, v765, loadedv1867, v766, result_symbol1869, v767, mark_end1870, v768, v769, v770, cmp1871, v771, cmp1875, v772, cmp1878, v773, cmp1881, v774, cmp1885, v775, cmp1888, v776, cmp1891, v777, cmp1895, v778, cmp1898, v779, cmp1901, v780, cmp1904, v781, cmp1907, v782, cmp1910, v783, cmp1913, v784, loadedv1917, v785, result_symbol1919, v786, mark_end1920, v787, v788, v789, cmp1921, v790, cmp1925, v791, cmp1928, v792, cmp1931, v793, cmp1935, v794, cmp1938, v795, cmp1941, v796, cmp1945, v797, cmp1948, v798, cmp1951, v799, cmp1954, v800, cmp1957, v801, cmp1960, v802, cmp1963, v803, loadedv1967, v804, result_symbol1969, v805, mark_end1970, v806, v807, v808, cmp1971, v809, cmp1975, v810, cmp1978, v811, cmp1981, v812, cmp1985, v813, cmp1988, v814, cmp1991, v815, cmp1995, v816, cmp1998, v817, cmp2001, v818, cmp2004, v819, cmp2007, v820, cmp2010, v821, cmp2013, v822, loadedv2017, v823, result_symbol2019, v824, mark_end2020, v825, v826, v827, cmp2021, v828, cmp2025, v829, cmp2028, v830, cmp2031, v831, cmp2035, v832, cmp2038, v833, cmp2041, v834, cmp2045, v835, cmp2048, v836, cmp2051, v837, cmp2054, v838, cmp2057, v839, cmp2060, v840, cmp2063, v841, loadedv2067, v842, result_symbol2069, v843, mark_end2070, v844, v845, v846, cmp2071, v847, cmp2074, v848, cmp2077, v849, cmp2081, v850, cmp2084, v851, cmp2087, v852, cmp2091, v853, cmp2094, v854, cmp2097, v855, cmp2100, v856, cmp2103, v857, cmp2106, v858, cmp2109, v859, loadedv2113, v860, result_symbol2115, v861, mark_end2116, v862, v863, v864, cmp2117, v865, cmp2120, v866, cmp2123, v867, cmp2126, v868, cmp2129, v869, cmp2132, v870, cmp2135, v871, cmp2138, v872, loadedv2142, v873, result_symbol2144, v874, mark_end2145, v875, v876, v877, cmp2146, v878, cmp2149, v879, cmp2152, v880, cmp2155, v881, cmp2158, v882, cmp2161, v883, cmp2164, v884, cmp2167, v885, cmp2170, v886, cmp2173, v887, loadedv2177, v888, result_symbol2179, v889, mark_end2180, v890, v891, v892, cmp2181, v893, cmp2184, v894, cmp2187, v895, cmp2191, v896, cmp2194, v897, loadedv2198, v898, result_symbol2200, v899, mark_end2201, v900, v901, v902, cmp2202, v903, cmp2205, v904, loadedv2209, v905, result_symbol2211, v906, mark_end2212, v907, v908, v909, cmp2213, v910, cmp2216, v911, cmp2219, v912, cmp2222, v913, cmp2226, v914, cmp2229, v915, loadedv2233, v916, result_symbol2235, v917, mark_end2236, v918, v919, v920, cmp2237, v921, cmp2240, v922, loadedv2244, v923

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
		goto sw_bb83
	case 2:
		goto sw_bb113
	case 3:
		goto sw_bb133
	case 4:
		goto sw_bb139
	case 5:
		goto sw_bb145
	case 6:
		goto sw_bb151
	case 7:
		goto sw_bb157
	case 8:
		goto sw_bb195
	case 9:
		goto sw_bb271
	case 10:
		goto sw_bb277
	case 11:
		goto sw_bb290
	case 12:
		goto sw_bb336
	case 13:
		goto sw_bb342
	case 14:
		goto sw_bb348
	case 15:
		goto sw_bb354
	case 16:
		goto sw_bb360
	case 17:
		goto sw_bb366
	case 18:
		goto sw_bb372
	case 19:
		goto sw_bb378
	case 20:
		goto sw_bb384
	case 21:
		goto sw_bb390
	case 22:
		goto sw_bb396
	case 23:
		goto sw_bb402
	case 24:
		goto sw_bb408
	case 25:
		goto sw_bb414
	case 26:
		goto sw_bb420
	case 27:
		goto sw_bb426
	case 28:
		goto sw_bb432
	case 29:
		goto sw_bb458
	case 30:
		goto sw_bb483
	case 31:
		goto sw_bb489
	case 32:
		goto sw_bb514
	case 33:
		goto sw_bb520
	case 34:
		goto sw_bb526
	case 35:
		goto sw_bb532
	case 36:
		goto sw_bb538
	case 37:
		goto sw_bb544
	case 38:
		goto sw_bb550
	case 39:
		goto sw_bb556
	case 40:
		goto sw_bb562
	case 41:
		goto sw_bb568
	case 42:
		goto sw_bb574
	case 43:
		goto sw_bb580
	case 44:
		goto sw_bb605
	case 45:
		goto sw_bb611
	case 46:
		goto sw_bb617
	case 47:
		goto sw_bb623
	case 48:
		goto sw_bb629
	case 49:
		goto sw_bb635
	case 50:
		goto sw_bb641
	case 51:
		goto sw_bb647
	case 52:
		goto sw_bb653
	case 53:
		goto sw_bb659
	case 54:
		goto sw_bb665
	case 55:
		goto sw_bb671
	case 56:
		goto sw_bb677
	case 57:
		goto sw_bb683
	case 58:
		goto sw_bb689
	case 59:
		goto sw_bb695
	case 60:
		goto sw_bb701
	case 61:
		goto sw_bb707
	case 62:
		goto sw_bb729
	case 63:
		goto sw_bb750
	case 64:
		goto sw_bb771
	case 65:
		goto sw_bb780
	case 66:
		goto sw_bb871
	case 67:
		goto sw_bb873
	case 68:
		goto sw_bb877
	case 69:
		goto sw_bb881
	case 70:
		goto sw_bb885
	case 71:
		goto sw_bb900
	case 72:
		goto sw_bb904
	case 73:
		goto sw_bb912
	case 74:
		goto sw_bb916
	case 75:
		goto sw_bb920
	case 76:
		goto sw_bb938
	case 77:
		goto sw_bb955
	case 78:
		goto sw_bb986
	case 79:
		goto sw_bb1007
	case 80:
		goto sw_bb1028
	case 81:
		goto sw_bb1049
	case 82:
		goto sw_bb1070
	case 83:
		goto sw_bb1087
	case 84:
		goto sw_bb1091
	case 85:
		goto sw_bb1108
	case 86:
		goto sw_bb1112
	case 87:
		goto sw_bb1116
	case 88:
		goto sw_bb1127
	case 89:
		goto sw_bb1131
	case 90:
		goto sw_bb1135
	case 91:
		goto sw_bb1158
	case 92:
		goto sw_bb1162
	case 93:
		goto sw_bb1180
	case 94:
		goto sw_bb1207
	case 95:
		goto sw_bb1234
	case 96:
		goto sw_bb1300
	case 97:
		goto sw_bb1327
	case 98:
		goto sw_bb1358
	case 99:
		goto sw_bb1388
	case 100:
		goto sw_bb1450
	case 101:
		goto sw_bb1474
	case 102:
		goto sw_bb1497
	case 103:
		goto sw_bb1545
	case 104:
		goto sw_bb1549
	case 105:
		goto sw_bb1553
	case 106:
		goto sw_bb1557
	case 107:
		goto sw_bb1568
	case 108:
		goto sw_bb1618
	case 109:
		goto sw_bb1668
	case 110:
		goto sw_bb1718
	case 111:
		goto sw_bb1768
	case 112:
		goto sw_bb1818
	case 113:
		goto sw_bb1868
	case 114:
		goto sw_bb1918
	case 115:
		goto sw_bb1968
	case 116:
		goto sw_bb2018
	case 117:
		goto sw_bb2068
	case 118:
		goto sw_bb2114
	case 119:
		goto sw_bb2143
	case 120:
		goto sw_bb2178
	case 121:
		goto sw_bb2199
	case 122:
		goto sw_bb2210
	case 123:
		goto sw_bb2234
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
	*libc.As[int16](state_addr) = 66
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
	*libc.As[int16](state_addr) = 85
	goto next_state

if_end6:
	v12 = *libc.As[int32](lookahead)
	cmp7 = v12 == 40
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*libc.As[int16](state_addr) = 72
	goto next_state

if_end10:
	v13 = *libc.As[int32](lookahead)
	cmp11 = v13 == 41
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end14:
	v14 = *libc.As[int32](lookahead)
	cmp15 = v14 == 43
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*libc.As[int16](state_addr) = 88
	goto next_state

if_end18:
	v15 = *libc.As[int32](lookahead)
	cmp19 = v15 == 58
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end22:
	v16 = *libc.As[int32](lookahead)
	cmp23 = v16 == 60
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end26:
	v17 = *libc.As[int32](lookahead)
	cmp27 = v17 == 62
	if cmp27 {
		goto if_then29
	} else {
		goto if_end30
	}

if_then29:
	*libc.As[int16](state_addr) = 89
	goto next_state

if_end30:
	v18 = *libc.As[int32](lookahead)
	cmp31 = v18 == 68
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end34:
	v19 = *libc.As[int32](lookahead)
	cmp35 = v19 == 70
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*libc.As[int16](state_addr) = 94
	goto next_state

if_end38:
	v20 = *libc.As[int32](lookahead)
	cmp39 = v20 == 79
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end42:
	v21 = *libc.As[int32](lookahead)
	cmp43 = v21 == 100
	if cmp43 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*libc.As[int16](state_addr) = 96
	goto next_state

if_end46:
	v22 = *libc.As[int32](lookahead)
	cmp47 = v22 == 102
	if cmp47 {
		goto if_then49
	} else {
		goto if_end50
	}

if_then49:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end50:
	v23 = *libc.As[int32](lookahead)
	cmp51 = v23 == 9
	if cmp51 {
		goto if_then61
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v24 = *libc.As[int32](lookahead)
	cmp53 = v24 == 10
	if cmp53 {
		goto if_then61
	} else {
		goto lor_lhs_false55
	}

lor_lhs_false55:
	v25 = *libc.As[int32](lookahead)
	cmp56 = v25 == 13
	if cmp56 {
		goto if_then61
	} else {
		goto lor_lhs_false58
	}

lor_lhs_false58:
	v26 = *libc.As[int32](lookahead)
	cmp59 = v26 == 32
	if cmp59 {
		goto if_then61
	} else {
		goto if_end62
	}

if_then61:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 0
	goto next_state

if_end62:
	v27 = *libc.As[int32](lookahead)
	cmp63 = 65 <= v27
	if cmp63 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false67
	}

land_lhs_true:
	v28 = *libc.As[int32](lookahead)
	cmp65 = v28 <= 69
	if cmp65 {
		goto if_then73
	} else {
		goto lor_lhs_false67
	}

lor_lhs_false67:
	v29 = *libc.As[int32](lookahead)
	cmp68 = 97 <= v29
	if cmp68 {
		goto land_lhs_true70
	} else {
		goto if_end74
	}

land_lhs_true70:
	v30 = *libc.As[int32](lookahead)
	cmp71 = v30 <= 101
	if cmp71 {
		goto if_then73
	} else {
		goto if_end74
	}

if_then73:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end74:
	v31 = *libc.As[int32](lookahead)
	cmp75 = 48 <= v31
	if cmp75 {
		goto land_lhs_true77
	} else {
		goto if_end81
	}

land_lhs_true77:
	v32 = *libc.As[int32](lookahead)
	cmp78 = v32 <= 57
	if cmp78 {
		goto if_then80
	} else {
		goto if_end81
	}

if_then80:
	*libc.As[int16](state_addr) = 100
	goto next_state

if_end81:
	v33 = *libc.As[byte](result)
	loadedv82 = (v33 & 1) != 0
	*libc.As[bool](retval) = loadedv82
	goto _return

sw_bb83:
	v34 = *libc.As[int32](lookahead)
	cmp84 = v34 == 10
	if cmp84 {
		goto if_then86
	} else {
		goto if_end87
	}

if_then86:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end87:
	v35 = *libc.As[int32](lookahead)
	cmp88 = v35 == 40
	if cmp88 {
		goto if_then90
	} else {
		goto if_end91
	}

if_then90:
	*libc.As[int16](state_addr) = 80
	goto next_state

if_end91:
	v36 = *libc.As[int32](lookahead)
	cmp92 = v36 == 9
	if cmp92 {
		goto if_then100
	} else {
		goto lor_lhs_false94
	}

lor_lhs_false94:
	v37 = *libc.As[int32](lookahead)
	cmp95 = v37 == 13
	if cmp95 {
		goto if_then100
	} else {
		goto lor_lhs_false97
	}

lor_lhs_false97:
	v38 = *libc.As[int32](lookahead)
	cmp98 = v38 == 32
	if cmp98 {
		goto if_then100
	} else {
		goto if_end101
	}

if_then100:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end101:
	v39 = *libc.As[int32](lookahead)
	cmp102 = v39 != 0
	if cmp102 {
		goto land_lhs_true104
	} else {
		goto if_end111
	}

land_lhs_true104:
	v40 = *libc.As[int32](lookahead)
	cmp105 = v40 != 35
	if cmp105 {
		goto land_lhs_true107
	} else {
		goto if_end111
	}

land_lhs_true107:
	v41 = *libc.As[int32](lookahead)
	cmp108 = v41 != 60
	if cmp108 {
		goto if_then110
	} else {
		goto if_end111
	}

if_then110:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end111:
	v42 = *libc.As[byte](result)
	loadedv112 = (v42 & 1) != 0
	*libc.As[bool](retval) = loadedv112
	goto _return

sw_bb113:
	v43 = *libc.As[int32](lookahead)
	cmp114 = v43 == 10
	if cmp114 {
		goto if_then116
	} else {
		goto if_end117
	}

if_then116:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end117:
	v44 = *libc.As[int32](lookahead)
	cmp118 = v44 == 9
	if cmp118 {
		goto if_then126
	} else {
		goto lor_lhs_false120
	}

lor_lhs_false120:
	v45 = *libc.As[int32](lookahead)
	cmp121 = v45 == 13
	if cmp121 {
		goto if_then126
	} else {
		goto lor_lhs_false123
	}

lor_lhs_false123:
	v46 = *libc.As[int32](lookahead)
	cmp124 = v46 == 32
	if cmp124 {
		goto if_then126
	} else {
		goto if_end127
	}

if_then126:
	*libc.As[int16](state_addr) = 120
	goto next_state

if_end127:
	v47 = *libc.As[int32](lookahead)
	cmp128 = v47 != 0
	if cmp128 {
		goto if_then130
	} else {
		goto if_end131
	}

if_then130:
	*libc.As[int16](state_addr) = 121
	goto next_state

if_end131:
	v48 = *libc.As[byte](result)
	loadedv132 = (v48 & 1) != 0
	*libc.As[bool](retval) = loadedv132
	goto _return

sw_bb133:
	v49 = *libc.As[int32](lookahead)
	cmp134 = v49 == 32
	if cmp134 {
		goto if_then136
	} else {
		goto if_end137
	}

if_then136:
	*libc.As[int16](state_addr) = 105
	goto next_state

if_end137:
	v50 = *libc.As[byte](result)
	loadedv138 = (v50 & 1) != 0
	*libc.As[bool](retval) = loadedv138
	goto _return

sw_bb139:
	v51 = *libc.As[int32](lookahead)
	cmp140 = v51 == 32
	if cmp140 {
		goto if_then142
	} else {
		goto if_end143
	}

if_then142:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end143:
	v52 = *libc.As[byte](result)
	loadedv144 = (v52 & 1) != 0
	*libc.As[bool](retval) = loadedv144
	goto _return

sw_bb145:
	v53 = *libc.As[int32](lookahead)
	cmp146 = v53 == 32
	if cmp146 {
		goto if_then148
	} else {
		goto if_end149
	}

if_then148:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end149:
	v54 = *libc.As[byte](result)
	loadedv150 = (v54 & 1) != 0
	*libc.As[bool](retval) = loadedv150
	goto _return

sw_bb151:
	v55 = *libc.As[int32](lookahead)
	cmp152 = v55 == 41
	if cmp152 {
		goto if_then154
	} else {
		goto if_end155
	}

if_then154:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end155:
	v56 = *libc.As[byte](result)
	loadedv156 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv156
	goto _return

sw_bb157:
	v57 = *libc.As[int32](lookahead)
	cmp158 = v57 == 48
	if cmp158 {
		goto if_then160
	} else {
		goto if_end161
	}

if_then160:
	*libc.As[int16](state_addr) = 98
	goto next_state

if_end161:
	v58 = *libc.As[int32](lookahead)
	cmp162 = v58 == 9
	if cmp162 {
		goto if_then173
	} else {
		goto lor_lhs_false164
	}

lor_lhs_false164:
	v59 = *libc.As[int32](lookahead)
	cmp165 = v59 == 10
	if cmp165 {
		goto if_then173
	} else {
		goto lor_lhs_false167
	}

lor_lhs_false167:
	v60 = *libc.As[int32](lookahead)
	cmp168 = v60 == 13
	if cmp168 {
		goto if_then173
	} else {
		goto lor_lhs_false170
	}

lor_lhs_false170:
	v61 = *libc.As[int32](lookahead)
	cmp171 = v61 == 32
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 7
	goto next_state

if_end174:
	v62 = *libc.As[int32](lookahead)
	cmp175 = 49 <= v62
	if cmp175 {
		goto land_lhs_true177
	} else {
		goto lor_lhs_false180
	}

land_lhs_true177:
	v63 = *libc.As[int32](lookahead)
	cmp178 = v63 <= 57
	if cmp178 {
		goto if_then192
	} else {
		goto lor_lhs_false180
	}

lor_lhs_false180:
	v64 = *libc.As[int32](lookahead)
	cmp181 = 65 <= v64
	if cmp181 {
		goto land_lhs_true183
	} else {
		goto lor_lhs_false186
	}

land_lhs_true183:
	v65 = *libc.As[int32](lookahead)
	cmp184 = v65 <= 70
	if cmp184 {
		goto if_then192
	} else {
		goto lor_lhs_false186
	}

lor_lhs_false186:
	v66 = *libc.As[int32](lookahead)
	cmp187 = 97 <= v66
	if cmp187 {
		goto land_lhs_true189
	} else {
		goto if_end193
	}

land_lhs_true189:
	v67 = *libc.As[int32](lookahead)
	cmp190 = v67 <= 102
	if cmp190 {
		goto if_then192
	} else {
		goto if_end193
	}

if_then192:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end193:
	v68 = *libc.As[byte](result)
	loadedv194 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv194
	goto _return

sw_bb195:
	v69 = *libc.As[int32](lookahead)
	cmp196 = v69 == 58
	if cmp196 {
		goto if_then198
	} else {
		goto if_end199
	}

if_then198:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end199:
	v70 = *libc.As[int32](lookahead)
	cmp200 = v70 == 60
	if cmp200 {
		goto if_then202
	} else {
		goto if_end203
	}

if_then202:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end203:
	v71 = *libc.As[int32](lookahead)
	cmp204 = v71 == 64
	if cmp204 {
		goto if_then206
	} else {
		goto if_end207
	}

if_then206:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end207:
	v72 = *libc.As[int32](lookahead)
	cmp208 = v72 == 43
	if cmp208 {
		goto if_then216
	} else {
		goto lor_lhs_false210
	}

lor_lhs_false210:
	v73 = *libc.As[int32](lookahead)
	cmp211 = v73 == 45
	if cmp211 {
		goto if_then216
	} else {
		goto lor_lhs_false213
	}

lor_lhs_false213:
	v74 = *libc.As[int32](lookahead)
	cmp214 = v74 == 47
	if cmp214 {
		goto if_then216
	} else {
		goto if_end217
	}

if_then216:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end217:
	v75 = *libc.As[int32](lookahead)
	cmp218 = v75 == 9
	if cmp218 {
		goto if_then229
	} else {
		goto lor_lhs_false220
	}

lor_lhs_false220:
	v76 = *libc.As[int32](lookahead)
	cmp221 = v76 == 10
	if cmp221 {
		goto if_then229
	} else {
		goto lor_lhs_false223
	}

lor_lhs_false223:
	v77 = *libc.As[int32](lookahead)
	cmp224 = v77 == 13
	if cmp224 {
		goto if_then229
	} else {
		goto lor_lhs_false226
	}

lor_lhs_false226:
	v78 = *libc.As[int32](lookahead)
	cmp227 = v78 == 32
	if cmp227 {
		goto if_then229
	} else {
		goto if_end230
	}

if_then229:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end230:
	v79 = *libc.As[int32](lookahead)
	cmp231 = 48 <= v79
	if cmp231 {
		goto land_lhs_true233
	} else {
		goto if_end237
	}

land_lhs_true233:
	v80 = *libc.As[int32](lookahead)
	cmp234 = v80 <= 57
	if cmp234 {
		goto if_then236
	} else {
		goto if_end237
	}

if_then236:
	*libc.As[int16](state_addr) = 102
	goto next_state

if_end237:
	v81 = *libc.As[int32](lookahead)
	cmp238 = 65 <= v81
	if cmp238 {
		goto land_lhs_true240
	} else {
		goto lor_lhs_false243
	}

land_lhs_true240:
	v82 = *libc.As[int32](lookahead)
	cmp241 = v82 <= 70
	if cmp241 {
		goto if_then249
	} else {
		goto lor_lhs_false243
	}

lor_lhs_false243:
	v83 = *libc.As[int32](lookahead)
	cmp244 = 97 <= v83
	if cmp244 {
		goto land_lhs_true246
	} else {
		goto if_end250
	}

land_lhs_true246:
	v84 = *libc.As[int32](lookahead)
	cmp247 = v84 <= 102
	if cmp247 {
		goto if_then249
	} else {
		goto if_end250
	}

if_then249:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end250:
	v85 = *libc.As[int32](lookahead)
	cmp251 = v85 == 46
	if cmp251 {
		goto if_then268
	} else {
		goto lor_lhs_false253
	}

lor_lhs_false253:
	v86 = *libc.As[int32](lookahead)
	cmp254 = 71 <= v86
	if cmp254 {
		goto land_lhs_true256
	} else {
		goto lor_lhs_false259
	}

land_lhs_true256:
	v87 = *libc.As[int32](lookahead)
	cmp257 = v87 <= 90
	if cmp257 {
		goto if_then268
	} else {
		goto lor_lhs_false259
	}

lor_lhs_false259:
	v88 = *libc.As[int32](lookahead)
	cmp260 = v88 == 95
	if cmp260 {
		goto if_then268
	} else {
		goto lor_lhs_false262
	}

lor_lhs_false262:
	v89 = *libc.As[int32](lookahead)
	cmp263 = 103 <= v89
	if cmp263 {
		goto land_lhs_true265
	} else {
		goto if_end269
	}

land_lhs_true265:
	v90 = *libc.As[int32](lookahead)
	cmp266 = v90 <= 122
	if cmp266 {
		goto if_then268
	} else {
		goto if_end269
	}

if_then268:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end269:
	v91 = *libc.As[byte](result)
	loadedv270 = (v91 & 1) != 0
	*libc.As[bool](retval) = loadedv270
	goto _return

sw_bb271:
	v92 = *libc.As[int32](lookahead)
	cmp272 = v92 == 58
	if cmp272 {
		goto if_then274
	} else {
		goto if_end275
	}

if_then274:
	*libc.As[int16](state_addr) = 104
	goto next_state

if_end275:
	v93 = *libc.As[byte](result)
	loadedv276 = (v93 & 1) != 0
	*libc.As[bool](retval) = loadedv276
	goto _return

sw_bb277:
	v94 = *libc.As[int32](lookahead)
	cmp278 = v94 == 62
	if cmp278 {
		goto if_then280
	} else {
		goto if_end281
	}

if_then280:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end281:
	v95 = *libc.As[int32](lookahead)
	cmp282 = v95 != 0
	if cmp282 {
		goto land_lhs_true284
	} else {
		goto if_end288
	}

land_lhs_true284:
	v96 = *libc.As[int32](lookahead)
	cmp285 = v96 != 10
	if cmp285 {
		goto if_then287
	} else {
		goto if_end288
	}

if_then287:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end288:
	v97 = *libc.As[byte](result)
	loadedv289 = (v97 & 1) != 0
	*libc.As[bool](retval) = loadedv289
	goto _return

sw_bb290:
	v98 = *libc.As[int32](lookahead)
	cmp291 = v98 == 70
	if cmp291 {
		goto if_then293
	} else {
		goto if_end294
	}

if_then293:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end294:
	v99 = *libc.As[int32](lookahead)
	cmp295 = v99 == 100
	if cmp295 {
		goto if_then297
	} else {
		goto if_end298
	}

if_then297:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end298:
	v100 = *libc.As[int32](lookahead)
	cmp299 = v100 == 102
	if cmp299 {
		goto if_then301
	} else {
		goto if_end302
	}

if_then301:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end302:
	v101 = *libc.As[int32](lookahead)
	cmp303 = v101 == 9
	if cmp303 {
		goto if_then314
	} else {
		goto lor_lhs_false305
	}

lor_lhs_false305:
	v102 = *libc.As[int32](lookahead)
	cmp306 = v102 == 10
	if cmp306 {
		goto if_then314
	} else {
		goto lor_lhs_false308
	}

lor_lhs_false308:
	v103 = *libc.As[int32](lookahead)
	cmp309 = v103 == 13
	if cmp309 {
		goto if_then314
	} else {
		goto lor_lhs_false311
	}

lor_lhs_false311:
	v104 = *libc.As[int32](lookahead)
	cmp312 = v104 == 32
	if cmp312 {
		goto if_then314
	} else {
		goto if_end315
	}

if_then314:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end315:
	v105 = *libc.As[int32](lookahead)
	cmp316 = 48 <= v105
	if cmp316 {
		goto land_lhs_true318
	} else {
		goto lor_lhs_false321
	}

land_lhs_true318:
	v106 = *libc.As[int32](lookahead)
	cmp319 = v106 <= 57
	if cmp319 {
		goto if_then333
	} else {
		goto lor_lhs_false321
	}

lor_lhs_false321:
	v107 = *libc.As[int32](lookahead)
	cmp322 = 65 <= v107
	if cmp322 {
		goto land_lhs_true324
	} else {
		goto lor_lhs_false327
	}

land_lhs_true324:
	v108 = *libc.As[int32](lookahead)
	cmp325 = v108 <= 69
	if cmp325 {
		goto if_then333
	} else {
		goto lor_lhs_false327
	}

lor_lhs_false327:
	v109 = *libc.As[int32](lookahead)
	cmp328 = 97 <= v109
	if cmp328 {
		goto land_lhs_true330
	} else {
		goto if_end334
	}

land_lhs_true330:
	v110 = *libc.As[int32](lookahead)
	cmp331 = v110 <= 101
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end334:
	v111 = *libc.As[byte](result)
	loadedv335 = (v111 & 1) != 0
	*libc.As[bool](retval) = loadedv335
	goto _return

sw_bb336:
	v112 = *libc.As[int32](lookahead)
	cmp337 = v112 == 97
	if cmp337 {
		goto if_then339
	} else {
		goto if_end340
	}

if_then339:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end340:
	v113 = *libc.As[byte](result)
	loadedv341 = (v113 & 1) != 0
	*libc.As[bool](retval) = loadedv341
	goto _return

sw_bb342:
	v114 = *libc.As[int32](lookahead)
	cmp343 = v114 == 97
	if cmp343 {
		goto if_then345
	} else {
		goto if_end346
	}

if_then345:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end346:
	v115 = *libc.As[byte](result)
	loadedv347 = (v115 & 1) != 0
	*libc.As[bool](retval) = loadedv347
	goto _return

sw_bb348:
	v116 = *libc.As[int32](lookahead)
	cmp349 = v116 == 97
	if cmp349 {
		goto if_then351
	} else {
		goto if_end352
	}

if_then351:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end352:
	v117 = *libc.As[byte](result)
	loadedv353 = (v117 & 1) != 0
	*libc.As[bool](retval) = loadedv353
	goto _return

sw_bb354:
	v118 = *libc.As[int32](lookahead)
	cmp355 = v118 == 97
	if cmp355 {
		goto if_then357
	} else {
		goto if_end358
	}

if_then357:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end358:
	v119 = *libc.As[byte](result)
	loadedv359 = (v119 & 1) != 0
	*libc.As[bool](retval) = loadedv359
	goto _return

sw_bb360:
	v120 = *libc.As[int32](lookahead)
	cmp361 = v120 == 98
	if cmp361 {
		goto if_then363
	} else {
		goto if_end364
	}

if_then363:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end364:
	v121 = *libc.As[byte](result)
	loadedv365 = (v121 & 1) != 0
	*libc.As[bool](retval) = loadedv365
	goto _return

sw_bb366:
	v122 = *libc.As[int32](lookahead)
	cmp367 = v122 == 99
	if cmp367 {
		goto if_then369
	} else {
		goto if_end370
	}

if_then369:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end370:
	v123 = *libc.As[byte](result)
	loadedv371 = (v123 & 1) != 0
	*libc.As[bool](retval) = loadedv371
	goto _return

sw_bb372:
	v124 = *libc.As[int32](lookahead)
	cmp373 = v124 == 99
	if cmp373 {
		goto if_then375
	} else {
		goto if_end376
	}

if_then375:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end376:
	v125 = *libc.As[byte](result)
	loadedv377 = (v125 & 1) != 0
	*libc.As[bool](retval) = loadedv377
	goto _return

sw_bb378:
	v126 = *libc.As[int32](lookahead)
	cmp379 = v126 == 100
	if cmp379 {
		goto if_then381
	} else {
		goto if_end382
	}

if_then381:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end382:
	v127 = *libc.As[byte](result)
	loadedv383 = (v127 & 1) != 0
	*libc.As[bool](retval) = loadedv383
	goto _return

sw_bb384:
	v128 = *libc.As[int32](lookahead)
	cmp385 = v128 == 101
	if cmp385 {
		goto if_then387
	} else {
		goto if_end388
	}

if_then387:
	*libc.As[int16](state_addr) = 103
	goto next_state

if_end388:
	v129 = *libc.As[byte](result)
	loadedv389 = (v129 & 1) != 0
	*libc.As[bool](retval) = loadedv389
	goto _return

sw_bb390:
	v130 = *libc.As[int32](lookahead)
	cmp391 = v130 == 101
	if cmp391 {
		goto if_then393
	} else {
		goto if_end394
	}

if_then393:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end394:
	v131 = *libc.As[byte](result)
	loadedv395 = (v131 & 1) != 0
	*libc.As[bool](retval) = loadedv395
	goto _return

sw_bb396:
	v132 = *libc.As[int32](lookahead)
	cmp397 = v132 == 101
	if cmp397 {
		goto if_then399
	} else {
		goto if_end400
	}

if_then399:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end400:
	v133 = *libc.As[byte](result)
	loadedv401 = (v133 & 1) != 0
	*libc.As[bool](retval) = loadedv401
	goto _return

sw_bb402:
	v134 = *libc.As[int32](lookahead)
	cmp403 = v134 == 101
	if cmp403 {
		goto if_then405
	} else {
		goto if_end406
	}

if_then405:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end406:
	v135 = *libc.As[byte](result)
	loadedv407 = (v135 & 1) != 0
	*libc.As[bool](retval) = loadedv407
	goto _return

sw_bb408:
	v136 = *libc.As[int32](lookahead)
	cmp409 = v136 == 101
	if cmp409 {
		goto if_then411
	} else {
		goto if_end412
	}

if_then411:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end412:
	v137 = *libc.As[byte](result)
	loadedv413 = (v137 & 1) != 0
	*libc.As[bool](retval) = loadedv413
	goto _return

sw_bb414:
	v138 = *libc.As[int32](lookahead)
	cmp415 = v138 == 102
	if cmp415 {
		goto if_then417
	} else {
		goto if_end418
	}

if_then417:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end418:
	v139 = *libc.As[byte](result)
	loadedv419 = (v139 & 1) != 0
	*libc.As[bool](retval) = loadedv419
	goto _return

sw_bb420:
	v140 = *libc.As[int32](lookahead)
	cmp421 = v140 == 102
	if cmp421 {
		goto if_then423
	} else {
		goto if_end424
	}

if_then423:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end424:
	v141 = *libc.As[byte](result)
	loadedv425 = (v141 & 1) != 0
	*libc.As[bool](retval) = loadedv425
	goto _return

sw_bb426:
	v142 = *libc.As[int32](lookahead)
	cmp427 = v142 == 102
	if cmp427 {
		goto if_then429
	} else {
		goto if_end430
	}

if_then429:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end430:
	v143 = *libc.As[byte](result)
	loadedv431 = (v143 & 1) != 0
	*libc.As[bool](retval) = loadedv431
	goto _return

sw_bb432:
	v144 = *libc.As[int32](lookahead)
	cmp433 = v144 == 102
	if cmp433 {
		goto if_then435
	} else {
		goto if_end436
	}

if_then435:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end436:
	v145 = *libc.As[int32](lookahead)
	cmp437 = v145 == 9
	if cmp437 {
		goto if_then448
	} else {
		goto lor_lhs_false439
	}

lor_lhs_false439:
	v146 = *libc.As[int32](lookahead)
	cmp440 = v146 == 10
	if cmp440 {
		goto if_then448
	} else {
		goto lor_lhs_false442
	}

lor_lhs_false442:
	v147 = *libc.As[int32](lookahead)
	cmp443 = v147 == 13
	if cmp443 {
		goto if_then448
	} else {
		goto lor_lhs_false445
	}

lor_lhs_false445:
	v148 = *libc.As[int32](lookahead)
	cmp446 = v148 == 32
	if cmp446 {
		goto if_then448
	} else {
		goto if_end449
	}

if_then448:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end449:
	v149 = *libc.As[int32](lookahead)
	cmp450 = 48 <= v149
	if cmp450 {
		goto land_lhs_true452
	} else {
		goto if_end456
	}

land_lhs_true452:
	v150 = *libc.As[int32](lookahead)
	cmp453 = v150 <= 57
	if cmp453 {
		goto if_then455
	} else {
		goto if_end456
	}

if_then455:
	*libc.As[int16](state_addr) = 106
	goto next_state

if_end456:
	v151 = *libc.As[byte](result)
	loadedv457 = (v151 & 1) != 0
	*libc.As[bool](retval) = loadedv457
	goto _return

sw_bb458:
	v152 = *libc.As[int32](lookahead)
	cmp459 = v152 == 105
	if cmp459 {
		goto if_then461
	} else {
		goto if_end462
	}

if_then461:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end462:
	v153 = *libc.As[int32](lookahead)
	cmp463 = 48 <= v153
	if cmp463 {
		goto land_lhs_true465
	} else {
		goto lor_lhs_false468
	}

land_lhs_true465:
	v154 = *libc.As[int32](lookahead)
	cmp466 = v154 <= 57
	if cmp466 {
		goto if_then480
	} else {
		goto lor_lhs_false468
	}

lor_lhs_false468:
	v155 = *libc.As[int32](lookahead)
	cmp469 = 65 <= v155
	if cmp469 {
		goto land_lhs_true471
	} else {
		goto lor_lhs_false474
	}

land_lhs_true471:
	v156 = *libc.As[int32](lookahead)
	cmp472 = v156 <= 70
	if cmp472 {
		goto if_then480
	} else {
		goto lor_lhs_false474
	}

lor_lhs_false474:
	v157 = *libc.As[int32](lookahead)
	cmp475 = 97 <= v157
	if cmp475 {
		goto land_lhs_true477
	} else {
		goto if_end481
	}

land_lhs_true477:
	v158 = *libc.As[int32](lookahead)
	cmp478 = v158 <= 102
	if cmp478 {
		goto if_then480
	} else {
		goto if_end481
	}

if_then480:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end481:
	v159 = *libc.As[byte](result)
	loadedv482 = (v159 & 1) != 0
	*libc.As[bool](retval) = loadedv482
	goto _return

sw_bb483:
	v160 = *libc.As[int32](lookahead)
	cmp484 = v160 == 105
	if cmp484 {
		goto if_then486
	} else {
		goto if_end487
	}

if_then486:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end487:
	v161 = *libc.As[byte](result)
	loadedv488 = (v161 & 1) != 0
	*libc.As[bool](retval) = loadedv488
	goto _return

sw_bb489:
	v162 = *libc.As[int32](lookahead)
	cmp490 = v162 == 105
	if cmp490 {
		goto if_then492
	} else {
		goto if_end493
	}

if_then492:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end493:
	v163 = *libc.As[int32](lookahead)
	cmp494 = 48 <= v163
	if cmp494 {
		goto land_lhs_true496
	} else {
		goto lor_lhs_false499
	}

land_lhs_true496:
	v164 = *libc.As[int32](lookahead)
	cmp497 = v164 <= 57
	if cmp497 {
		goto if_then511
	} else {
		goto lor_lhs_false499
	}

lor_lhs_false499:
	v165 = *libc.As[int32](lookahead)
	cmp500 = 65 <= v165
	if cmp500 {
		goto land_lhs_true502
	} else {
		goto lor_lhs_false505
	}

land_lhs_true502:
	v166 = *libc.As[int32](lookahead)
	cmp503 = v166 <= 70
	if cmp503 {
		goto if_then511
	} else {
		goto lor_lhs_false505
	}

lor_lhs_false505:
	v167 = *libc.As[int32](lookahead)
	cmp506 = 97 <= v167
	if cmp506 {
		goto land_lhs_true508
	} else {
		goto if_end512
	}

land_lhs_true508:
	v168 = *libc.As[int32](lookahead)
	cmp509 = v168 <= 102
	if cmp509 {
		goto if_then511
	} else {
		goto if_end512
	}

if_then511:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end512:
	v169 = *libc.As[byte](result)
	loadedv513 = (v169 & 1) != 0
	*libc.As[bool](retval) = loadedv513
	goto _return

sw_bb514:
	v170 = *libc.As[int32](lookahead)
	cmp515 = v170 == 105
	if cmp515 {
		goto if_then517
	} else {
		goto if_end518
	}

if_then517:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end518:
	v171 = *libc.As[byte](result)
	loadedv519 = (v171 & 1) != 0
	*libc.As[bool](retval) = loadedv519
	goto _return

sw_bb520:
	v172 = *libc.As[int32](lookahead)
	cmp521 = v172 == 105
	if cmp521 {
		goto if_then523
	} else {
		goto if_end524
	}

if_then523:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end524:
	v173 = *libc.As[byte](result)
	loadedv525 = (v173 & 1) != 0
	*libc.As[bool](retval) = loadedv525
	goto _return

sw_bb526:
	v174 = *libc.As[int32](lookahead)
	cmp527 = v174 == 105
	if cmp527 {
		goto if_then529
	} else {
		goto if_end530
	}

if_then529:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end530:
	v175 = *libc.As[byte](result)
	loadedv531 = (v175 & 1) != 0
	*libc.As[bool](retval) = loadedv531
	goto _return

sw_bb532:
	v176 = *libc.As[int32](lookahead)
	cmp533 = v176 == 108
	if cmp533 {
		goto if_then535
	} else {
		goto if_end536
	}

if_then535:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end536:
	v177 = *libc.As[byte](result)
	loadedv537 = (v177 & 1) != 0
	*libc.As[bool](retval) = loadedv537
	goto _return

sw_bb538:
	v178 = *libc.As[int32](lookahead)
	cmp539 = v178 == 108
	if cmp539 {
		goto if_then541
	} else {
		goto if_end542
	}

if_then541:
	*libc.As[int16](state_addr) = 60
	goto next_state

if_end542:
	v179 = *libc.As[byte](result)
	loadedv543 = (v179 & 1) != 0
	*libc.As[bool](retval) = loadedv543
	goto _return

sw_bb544:
	v180 = *libc.As[int32](lookahead)
	cmp545 = v180 == 108
	if cmp545 {
		goto if_then547
	} else {
		goto if_end548
	}

if_then547:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end548:
	v181 = *libc.As[byte](result)
	loadedv549 = (v181 & 1) != 0
	*libc.As[bool](retval) = loadedv549
	goto _return

sw_bb550:
	v182 = *libc.As[int32](lookahead)
	cmp551 = v182 == 109
	if cmp551 {
		goto if_then553
	} else {
		goto if_end554
	}

if_then553:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end554:
	v183 = *libc.As[byte](result)
	loadedv555 = (v183 & 1) != 0
	*libc.As[bool](retval) = loadedv555
	goto _return

sw_bb556:
	v184 = *libc.As[int32](lookahead)
	cmp557 = v184 == 109
	if cmp557 {
		goto if_then559
	} else {
		goto if_end560
	}

if_then559:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end560:
	v185 = *libc.As[byte](result)
	loadedv561 = (v185 & 1) != 0
	*libc.As[bool](retval) = loadedv561
	goto _return

sw_bb562:
	v186 = *libc.As[int32](lookahead)
	cmp563 = v186 == 109
	if cmp563 {
		goto if_then565
	} else {
		goto if_end566
	}

if_then565:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end566:
	v187 = *libc.As[byte](result)
	loadedv567 = (v187 & 1) != 0
	*libc.As[bool](retval) = loadedv567
	goto _return

sw_bb568:
	v188 = *libc.As[int32](lookahead)
	cmp569 = v188 == 110
	if cmp569 {
		goto if_then571
	} else {
		goto if_end572
	}

if_then571:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end572:
	v189 = *libc.As[byte](result)
	loadedv573 = (v189 & 1) != 0
	*libc.As[bool](retval) = loadedv573
	goto _return

sw_bb574:
	v190 = *libc.As[int32](lookahead)
	cmp575 = v190 == 110
	if cmp575 {
		goto if_then577
	} else {
		goto if_end578
	}

if_then577:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end578:
	v191 = *libc.As[byte](result)
	loadedv579 = (v191 & 1) != 0
	*libc.As[bool](retval) = loadedv579
	goto _return

sw_bb580:
	v192 = *libc.As[int32](lookahead)
	cmp581 = v192 == 111
	if cmp581 {
		goto if_then583
	} else {
		goto if_end584
	}

if_then583:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end584:
	v193 = *libc.As[int32](lookahead)
	cmp585 = 48 <= v193
	if cmp585 {
		goto land_lhs_true587
	} else {
		goto lor_lhs_false590
	}

land_lhs_true587:
	v194 = *libc.As[int32](lookahead)
	cmp588 = v194 <= 57
	if cmp588 {
		goto if_then602
	} else {
		goto lor_lhs_false590
	}

lor_lhs_false590:
	v195 = *libc.As[int32](lookahead)
	cmp591 = 65 <= v195
	if cmp591 {
		goto land_lhs_true593
	} else {
		goto lor_lhs_false596
	}

land_lhs_true593:
	v196 = *libc.As[int32](lookahead)
	cmp594 = v196 <= 70
	if cmp594 {
		goto if_then602
	} else {
		goto lor_lhs_false596
	}

lor_lhs_false596:
	v197 = *libc.As[int32](lookahead)
	cmp597 = 97 <= v197
	if cmp597 {
		goto land_lhs_true599
	} else {
		goto if_end603
	}

land_lhs_true599:
	v198 = *libc.As[int32](lookahead)
	cmp600 = v198 <= 102
	if cmp600 {
		goto if_then602
	} else {
		goto if_end603
	}

if_then602:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end603:
	v199 = *libc.As[byte](result)
	loadedv604 = (v199 & 1) != 0
	*libc.As[bool](retval) = loadedv604
	goto _return

sw_bb605:
	v200 = *libc.As[int32](lookahead)
	cmp606 = v200 == 111
	if cmp606 {
		goto if_then608
	} else {
		goto if_end609
	}

if_then608:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end609:
	v201 = *libc.As[byte](result)
	loadedv610 = (v201 & 1) != 0
	*libc.As[bool](retval) = loadedv610
	goto _return

sw_bb611:
	v202 = *libc.As[int32](lookahead)
	cmp612 = v202 == 111
	if cmp612 {
		goto if_then614
	} else {
		goto if_end615
	}

if_then614:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end615:
	v203 = *libc.As[byte](result)
	loadedv616 = (v203 & 1) != 0
	*libc.As[bool](retval) = loadedv616
	goto _return

sw_bb617:
	v204 = *libc.As[int32](lookahead)
	cmp618 = v204 == 111
	if cmp618 {
		goto if_then620
	} else {
		goto if_end621
	}

if_then620:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end621:
	v205 = *libc.As[byte](result)
	loadedv622 = (v205 & 1) != 0
	*libc.As[bool](retval) = loadedv622
	goto _return

sw_bb623:
	v206 = *libc.As[int32](lookahead)
	cmp624 = v206 == 114
	if cmp624 {
		goto if_then626
	} else {
		goto if_end627
	}

if_then626:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end627:
	v207 = *libc.As[byte](result)
	loadedv628 = (v207 & 1) != 0
	*libc.As[bool](retval) = loadedv628
	goto _return

sw_bb629:
	v208 = *libc.As[int32](lookahead)
	cmp630 = v208 == 114
	if cmp630 {
		goto if_then632
	} else {
		goto if_end633
	}

if_then632:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end633:
	v209 = *libc.As[byte](result)
	loadedv634 = (v209 & 1) != 0
	*libc.As[bool](retval) = loadedv634
	goto _return

sw_bb635:
	v210 = *libc.As[int32](lookahead)
	cmp636 = v210 == 114
	if cmp636 {
		goto if_then638
	} else {
		goto if_end639
	}

if_then638:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end639:
	v211 = *libc.As[byte](result)
	loadedv640 = (v211 & 1) != 0
	*libc.As[bool](retval) = loadedv640
	goto _return

sw_bb641:
	v212 = *libc.As[int32](lookahead)
	cmp642 = v212 == 115
	if cmp642 {
		goto if_then644
	} else {
		goto if_end645
	}

if_then644:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end645:
	v213 = *libc.As[byte](result)
	loadedv646 = (v213 & 1) != 0
	*libc.As[bool](retval) = loadedv646
	goto _return

sw_bb647:
	v214 = *libc.As[int32](lookahead)
	cmp648 = v214 == 115
	if cmp648 {
		goto if_then650
	} else {
		goto if_end651
	}

if_then650:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end651:
	v215 = *libc.As[byte](result)
	loadedv652 = (v215 & 1) != 0
	*libc.As[bool](retval) = loadedv652
	goto _return

sw_bb653:
	v216 = *libc.As[int32](lookahead)
	cmp654 = v216 == 115
	if cmp654 {
		goto if_then656
	} else {
		goto if_end657
	}

if_then656:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end657:
	v217 = *libc.As[byte](result)
	loadedv658 = (v217 & 1) != 0
	*libc.As[bool](retval) = loadedv658
	goto _return

sw_bb659:
	v218 = *libc.As[int32](lookahead)
	cmp660 = v218 == 115
	if cmp660 {
		goto if_then662
	} else {
		goto if_end663
	}

if_then662:
	*libc.As[int16](state_addr) = 24
	goto next_state

if_end663:
	v219 = *libc.As[byte](result)
	loadedv664 = (v219 & 1) != 0
	*libc.As[bool](retval) = loadedv664
	goto _return

sw_bb665:
	v220 = *libc.As[int32](lookahead)
	cmp666 = v220 == 115
	if cmp666 {
		goto if_then668
	} else {
		goto if_end669
	}

if_then668:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end669:
	v221 = *libc.As[byte](result)
	loadedv670 = (v221 & 1) != 0
	*libc.As[bool](retval) = loadedv670
	goto _return

sw_bb671:
	v222 = *libc.As[int32](lookahead)
	cmp672 = v222 == 115
	if cmp672 {
		goto if_then674
	} else {
		goto if_end675
	}

if_then674:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end675:
	v223 = *libc.As[byte](result)
	loadedv676 = (v223 & 1) != 0
	*libc.As[bool](retval) = loadedv676
	goto _return

sw_bb677:
	v224 = *libc.As[int32](lookahead)
	cmp678 = v224 == 116
	if cmp678 {
		goto if_then680
	} else {
		goto if_end681
	}

if_then680:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end681:
	v225 = *libc.As[byte](result)
	loadedv682 = (v225 & 1) != 0
	*libc.As[bool](retval) = loadedv682
	goto _return

sw_bb683:
	v226 = *libc.As[int32](lookahead)
	cmp684 = v226 == 116
	if cmp684 {
		goto if_then686
	} else {
		goto if_end687
	}

if_then686:
	*libc.As[int16](state_addr) = 69
	goto next_state

if_end687:
	v227 = *libc.As[byte](result)
	loadedv688 = (v227 & 1) != 0
	*libc.As[bool](retval) = loadedv688
	goto _return

sw_bb689:
	v228 = *libc.As[int32](lookahead)
	cmp690 = v228 == 116
	if cmp690 {
		goto if_then692
	} else {
		goto if_end693
	}

if_then692:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end693:
	v229 = *libc.As[byte](result)
	loadedv694 = (v229 & 1) != 0
	*libc.As[bool](retval) = loadedv694
	goto _return

sw_bb695:
	v230 = *libc.As[int32](lookahead)
	cmp696 = v230 == 116
	if cmp696 {
		goto if_then698
	} else {
		goto if_end699
	}

if_then698:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end699:
	v231 = *libc.As[byte](result)
	loadedv700 = (v231 & 1) != 0
	*libc.As[bool](retval) = loadedv700
	goto _return

sw_bb701:
	v232 = *libc.As[int32](lookahead)
	cmp702 = v232 == 121
	if cmp702 {
		goto if_then704
	} else {
		goto if_end705
	}

if_then704:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end705:
	v233 = *libc.As[byte](result)
	loadedv706 = (v233 & 1) != 0
	*libc.As[bool](retval) = loadedv706
	goto _return

sw_bb707:
	v234 = *libc.As[int32](lookahead)
	cmp708 = v234 == 9
	if cmp708 {
		goto if_then719
	} else {
		goto lor_lhs_false710
	}

lor_lhs_false710:
	v235 = *libc.As[int32](lookahead)
	cmp711 = v235 == 10
	if cmp711 {
		goto if_then719
	} else {
		goto lor_lhs_false713
	}

lor_lhs_false713:
	v236 = *libc.As[int32](lookahead)
	cmp714 = v236 == 13
	if cmp714 {
		goto if_then719
	} else {
		goto lor_lhs_false716
	}

lor_lhs_false716:
	v237 = *libc.As[int32](lookahead)
	cmp717 = v237 == 32
	if cmp717 {
		goto if_then719
	} else {
		goto if_end720
	}

if_then719:
	*libc.As[int16](state_addr) = 122
	goto next_state

if_end720:
	v238 = *libc.As[int32](lookahead)
	cmp721 = v238 != 0
	if cmp721 {
		goto land_lhs_true723
	} else {
		goto if_end727
	}

land_lhs_true723:
	v239 = *libc.As[int32](lookahead)
	cmp724 = v239 != 58
	if cmp724 {
		goto if_then726
	} else {
		goto if_end727
	}

if_then726:
	*libc.As[int16](state_addr) = 123
	goto next_state

if_end727:
	v240 = *libc.As[byte](result)
	loadedv728 = (v240 & 1) != 0
	*libc.As[bool](retval) = loadedv728
	goto _return

sw_bb729:
	v241 = *libc.As[int32](lookahead)
	cmp730 = 48 <= v241
	if cmp730 {
		goto land_lhs_true732
	} else {
		goto lor_lhs_false735
	}

land_lhs_true732:
	v242 = *libc.As[int32](lookahead)
	cmp733 = v242 <= 57
	if cmp733 {
		goto if_then747
	} else {
		goto lor_lhs_false735
	}

lor_lhs_false735:
	v243 = *libc.As[int32](lookahead)
	cmp736 = 65 <= v243
	if cmp736 {
		goto land_lhs_true738
	} else {
		goto lor_lhs_false741
	}

land_lhs_true738:
	v244 = *libc.As[int32](lookahead)
	cmp739 = v244 <= 70
	if cmp739 {
		goto if_then747
	} else {
		goto lor_lhs_false741
	}

lor_lhs_false741:
	v245 = *libc.As[int32](lookahead)
	cmp742 = 97 <= v245
	if cmp742 {
		goto land_lhs_true744
	} else {
		goto if_end748
	}

land_lhs_true744:
	v246 = *libc.As[int32](lookahead)
	cmp745 = v246 <= 102
	if cmp745 {
		goto if_then747
	} else {
		goto if_end748
	}

if_then747:
	*libc.As[int16](state_addr) = 90
	goto next_state

if_end748:
	v247 = *libc.As[byte](result)
	loadedv749 = (v247 & 1) != 0
	*libc.As[bool](retval) = loadedv749
	goto _return

sw_bb750:
	v248 = *libc.As[int32](lookahead)
	cmp751 = 48 <= v248
	if cmp751 {
		goto land_lhs_true753
	} else {
		goto lor_lhs_false756
	}

land_lhs_true753:
	v249 = *libc.As[int32](lookahead)
	cmp754 = v249 <= 57
	if cmp754 {
		goto if_then768
	} else {
		goto lor_lhs_false756
	}

lor_lhs_false756:
	v250 = *libc.As[int32](lookahead)
	cmp757 = 65 <= v250
	if cmp757 {
		goto land_lhs_true759
	} else {
		goto lor_lhs_false762
	}

land_lhs_true759:
	v251 = *libc.As[int32](lookahead)
	cmp760 = v251 <= 70
	if cmp760 {
		goto if_then768
	} else {
		goto lor_lhs_false762
	}

lor_lhs_false762:
	v252 = *libc.As[int32](lookahead)
	cmp763 = 97 <= v252
	if cmp763 {
		goto land_lhs_true765
	} else {
		goto if_end769
	}

land_lhs_true765:
	v253 = *libc.As[int32](lookahead)
	cmp766 = v253 <= 102
	if cmp766 {
		goto if_then768
	} else {
		goto if_end769
	}

if_then768:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end769:
	v254 = *libc.As[byte](result)
	loadedv770 = (v254 & 1) != 0
	*libc.As[bool](retval) = loadedv770
	goto _return

sw_bb771:
	v255 = *libc.As[int32](lookahead)
	cmp772 = v255 != 0
	if cmp772 {
		goto land_lhs_true774
	} else {
		goto if_end778
	}

land_lhs_true774:
	v256 = *libc.As[int32](lookahead)
	cmp775 = v256 != 10
	if cmp775 {
		goto if_then777
	} else {
		goto if_end778
	}

if_then777:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end778:
	v257 = *libc.As[byte](result)
	loadedv779 = (v257 & 1) != 0
	*libc.As[bool](retval) = loadedv779
	goto _return

sw_bb780:
	v258 = *libc.As[byte](eof)
	loadedv781 = (v258 & 1) != 0
	if loadedv781 {
		goto if_then782
	} else {
		goto if_end783
	}

if_then782:
	*libc.As[int16](state_addr) = 66
	goto next_state

if_end783:
	v259 = *libc.As[int32](lookahead)
	cmp784 = v259 == 35
	if cmp784 {
		goto if_then786
	} else {
		goto if_end787
	}

if_then786:
	*libc.As[int16](state_addr) = 85
	goto next_state

if_end787:
	v260 = *libc.As[int32](lookahead)
	cmp788 = v260 == 40
	if cmp788 {
		goto if_then790
	} else {
		goto if_end791
	}

if_then790:
	*libc.As[int16](state_addr) = 71
	goto next_state

if_end791:
	v261 = *libc.As[int32](lookahead)
	cmp792 = v261 == 58
	if cmp792 {
		goto if_then794
	} else {
		goto if_end795
	}

if_then794:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end795:
	v262 = *libc.As[int32](lookahead)
	cmp796 = v262 == 60
	if cmp796 {
		goto if_then798
	} else {
		goto if_end799
	}

if_then798:
	*libc.As[int16](state_addr) = 86
	goto next_state

if_end799:
	v263 = *libc.As[int32](lookahead)
	cmp800 = v263 == 64
	if cmp800 {
		goto if_then802
	} else {
		goto if_end803
	}

if_then802:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end803:
	v264 = *libc.As[int32](lookahead)
	cmp804 = v264 == 68
	if cmp804 {
		goto if_then806
	} else {
		goto if_end807
	}

if_then806:
	*libc.As[int16](state_addr) = 95
	goto next_state

if_end807:
	v265 = *libc.As[int32](lookahead)
	cmp808 = v265 == 43
	if cmp808 {
		goto if_then816
	} else {
		goto lor_lhs_false810
	}

lor_lhs_false810:
	v266 = *libc.As[int32](lookahead)
	cmp811 = v266 == 45
	if cmp811 {
		goto if_then816
	} else {
		goto lor_lhs_false813
	}

lor_lhs_false813:
	v267 = *libc.As[int32](lookahead)
	cmp814 = v267 == 47
	if cmp814 {
		goto if_then816
	} else {
		goto if_end817
	}

if_then816:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end817:
	v268 = *libc.As[int32](lookahead)
	cmp818 = v268 == 9
	if cmp818 {
		goto if_then829
	} else {
		goto lor_lhs_false820
	}

lor_lhs_false820:
	v269 = *libc.As[int32](lookahead)
	cmp821 = v269 == 10
	if cmp821 {
		goto if_then829
	} else {
		goto lor_lhs_false823
	}

lor_lhs_false823:
	v270 = *libc.As[int32](lookahead)
	cmp824 = v270 == 13
	if cmp824 {
		goto if_then829
	} else {
		goto lor_lhs_false826
	}

lor_lhs_false826:
	v271 = *libc.As[int32](lookahead)
	cmp827 = v271 == 32
	if cmp827 {
		goto if_then829
	} else {
		goto if_end830
	}

if_then829:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end830:
	v272 = *libc.As[int32](lookahead)
	cmp831 = 48 <= v272
	if cmp831 {
		goto land_lhs_true833
	} else {
		goto if_end837
	}

land_lhs_true833:
	v273 = *libc.As[int32](lookahead)
	cmp834 = v273 <= 57
	if cmp834 {
		goto if_then836
	} else {
		goto if_end837
	}

if_then836:
	*libc.As[int16](state_addr) = 102
	goto next_state

if_end837:
	v274 = *libc.As[int32](lookahead)
	cmp838 = 65 <= v274
	if cmp838 {
		goto land_lhs_true840
	} else {
		goto lor_lhs_false843
	}

land_lhs_true840:
	v275 = *libc.As[int32](lookahead)
	cmp841 = v275 <= 70
	if cmp841 {
		goto if_then849
	} else {
		goto lor_lhs_false843
	}

lor_lhs_false843:
	v276 = *libc.As[int32](lookahead)
	cmp844 = 97 <= v276
	if cmp844 {
		goto land_lhs_true846
	} else {
		goto if_end850
	}

land_lhs_true846:
	v277 = *libc.As[int32](lookahead)
	cmp847 = v277 <= 102
	if cmp847 {
		goto if_then849
	} else {
		goto if_end850
	}

if_then849:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end850:
	v278 = *libc.As[int32](lookahead)
	cmp851 = v278 == 46
	if cmp851 {
		goto if_then868
	} else {
		goto lor_lhs_false853
	}

lor_lhs_false853:
	v279 = *libc.As[int32](lookahead)
	cmp854 = 71 <= v279
	if cmp854 {
		goto land_lhs_true856
	} else {
		goto lor_lhs_false859
	}

land_lhs_true856:
	v280 = *libc.As[int32](lookahead)
	cmp857 = v280 <= 90
	if cmp857 {
		goto if_then868
	} else {
		goto lor_lhs_false859
	}

lor_lhs_false859:
	v281 = *libc.As[int32](lookahead)
	cmp860 = v281 == 95
	if cmp860 {
		goto if_then868
	} else {
		goto lor_lhs_false862
	}

lor_lhs_false862:
	v282 = *libc.As[int32](lookahead)
	cmp863 = 103 <= v282
	if cmp863 {
		goto land_lhs_true865
	} else {
		goto if_end869
	}

land_lhs_true865:
	v283 = *libc.As[int32](lookahead)
	cmp866 = v283 <= 122
	if cmp866 {
		goto if_then868
	} else {
		goto if_end869
	}

if_then868:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end869:
	v284 = *libc.As[byte](result)
	loadedv870 = (v284 & 1) != 0
	*libc.As[bool](retval) = loadedv870
	goto _return

sw_bb871:
	*libc.As[byte](result) = 1
	v285 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v285).F1)
	*libc.As[int16](result_symbol) = 0
	v286 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v286).F3)
	v287 = *libc.As[unsafe.Pointer](mark_end)
	v288 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v287)(v288)
	v289 = *libc.As[byte](result)
	loadedv872 = (v289 & 1) != 0
	*libc.As[bool](retval) = loadedv872
	goto _return

sw_bb873:
	*libc.As[byte](result) = 1
	v290 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol874 = libc.Ptr(&libc.As[TSLexer](v290).F1)
	*libc.As[int16](result_symbol874) = 1
	v291 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end875 = libc.Ptr(&libc.As[TSLexer](v291).F3)
	v292 = *libc.As[unsafe.Pointer](mark_end875)
	v293 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v292)(v293)
	v294 = *libc.As[byte](result)
	loadedv876 = (v294 & 1) != 0
	*libc.As[bool](retval) = loadedv876
	goto _return

sw_bb877:
	*libc.As[byte](result) = 1
	v295 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol878 = libc.Ptr(&libc.As[TSLexer](v295).F1)
	*libc.As[int16](result_symbol878) = 2
	v296 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end879 = libc.Ptr(&libc.As[TSLexer](v296).F3)
	v297 = *libc.As[unsafe.Pointer](mark_end879)
	v298 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v297)(v298)
	v299 = *libc.As[byte](result)
	loadedv880 = (v299 & 1) != 0
	*libc.As[bool](retval) = loadedv880
	goto _return

sw_bb881:
	*libc.As[byte](result) = 1
	v300 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol882 = libc.Ptr(&libc.As[TSLexer](v300).F1)
	*libc.As[int16](result_symbol882) = 3
	v301 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end883 = libc.Ptr(&libc.As[TSLexer](v301).F3)
	v302 = *libc.As[unsafe.Pointer](mark_end883)
	v303 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v302)(v303)
	v304 = *libc.As[byte](result)
	loadedv884 = (v304 & 1) != 0
	*libc.As[bool](retval) = loadedv884
	goto _return

sw_bb885:
	*libc.As[byte](result) = 1
	v305 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol886 = libc.Ptr(&libc.As[TSLexer](v305).F1)
	*libc.As[int16](result_symbol886) = 4
	v306 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end887 = libc.Ptr(&libc.As[TSLexer](v306).F3)
	v307 = *libc.As[unsafe.Pointer](mark_end887)
	v308 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v307)(v308)
	v309 = *libc.As[int32](lookahead)
	cmp888 = v309 == 62
	if cmp888 {
		goto if_then890
	} else {
		goto if_end891
	}

if_then890:
	*libc.As[int16](state_addr) = 70
	goto next_state

if_end891:
	v310 = *libc.As[int32](lookahead)
	cmp892 = v310 != 0
	if cmp892 {
		goto land_lhs_true894
	} else {
		goto if_end898
	}

land_lhs_true894:
	v311 = *libc.As[int32](lookahead)
	cmp895 = v311 != 10
	if cmp895 {
		goto if_then897
	} else {
		goto if_end898
	}

if_then897:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end898:
	v312 = *libc.As[byte](result)
	loadedv899 = (v312 & 1) != 0
	*libc.As[bool](retval) = loadedv899
	goto _return

sw_bb900:
	*libc.As[byte](result) = 1
	v313 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol901 = libc.Ptr(&libc.As[TSLexer](v313).F1)
	*libc.As[int16](result_symbol901) = 5
	v314 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end902 = libc.Ptr(&libc.As[TSLexer](v314).F3)
	v315 = *libc.As[unsafe.Pointer](mark_end902)
	v316 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v315)(v316)
	v317 = *libc.As[byte](result)
	loadedv903 = (v317 & 1) != 0
	*libc.As[bool](retval) = loadedv903
	goto _return

sw_bb904:
	*libc.As[byte](result) = 1
	v318 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol905 = libc.Ptr(&libc.As[TSLexer](v318).F1)
	*libc.As[int16](result_symbol905) = 5
	v319 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end906 = libc.Ptr(&libc.As[TSLexer](v319).F3)
	v320 = *libc.As[unsafe.Pointer](mark_end906)
	v321 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v320)(v321)
	v322 = *libc.As[int32](lookahead)
	cmp907 = v322 == 98
	if cmp907 {
		goto if_then909
	} else {
		goto if_end910
	}

if_then909:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end910:
	v323 = *libc.As[byte](result)
	loadedv911 = (v323 & 1) != 0
	*libc.As[bool](retval) = loadedv911
	goto _return

sw_bb912:
	*libc.As[byte](result) = 1
	v324 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol913 = libc.Ptr(&libc.As[TSLexer](v324).F1)
	*libc.As[int16](result_symbol913) = 6
	v325 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end914 = libc.Ptr(&libc.As[TSLexer](v325).F3)
	v326 = *libc.As[unsafe.Pointer](mark_end914)
	v327 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v326)(v327)
	v328 = *libc.As[byte](result)
	loadedv915 = (v328 & 1) != 0
	*libc.As[bool](retval) = loadedv915
	goto _return

sw_bb916:
	*libc.As[byte](result) = 1
	v329 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol917 = libc.Ptr(&libc.As[TSLexer](v329).F1)
	*libc.As[int16](result_symbol917) = 7
	v330 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end918 = libc.Ptr(&libc.As[TSLexer](v330).F3)
	v331 = *libc.As[unsafe.Pointer](mark_end918)
	v332 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v331)(v332)
	v333 = *libc.As[byte](result)
	loadedv919 = (v333 & 1) != 0
	*libc.As[bool](retval) = loadedv919
	goto _return

sw_bb920:
	*libc.As[byte](result) = 1
	v334 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol921 = libc.Ptr(&libc.As[TSLexer](v334).F1)
	*libc.As[int16](result_symbol921) = 8
	v335 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end922 = libc.Ptr(&libc.As[TSLexer](v335).F3)
	v336 = *libc.As[unsafe.Pointer](mark_end922)
	v337 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v336)(v337)
	v338 = *libc.As[int32](lookahead)
	cmp923 = v338 == 32
	if cmp923 {
		goto if_then925
	} else {
		goto if_end926
	}

if_then925:
	*libc.As[int16](state_addr) = 92
	goto next_state

if_end926:
	v339 = *libc.As[int32](lookahead)
	cmp927 = v339 == 9
	if cmp927 {
		goto if_then935
	} else {
		goto lor_lhs_false929
	}

lor_lhs_false929:
	v340 = *libc.As[int32](lookahead)
	cmp930 = v340 == 10
	if cmp930 {
		goto if_then935
	} else {
		goto lor_lhs_false932
	}

lor_lhs_false932:
	v341 = *libc.As[int32](lookahead)
	cmp933 = v341 == 13
	if cmp933 {
		goto if_then935
	} else {
		goto if_end936
	}

if_then935:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end936:
	v342 = *libc.As[byte](result)
	loadedv937 = (v342 & 1) != 0
	*libc.As[bool](retval) = loadedv937
	goto _return

sw_bb938:
	*libc.As[byte](result) = 1
	v343 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol939 = libc.Ptr(&libc.As[TSLexer](v343).F1)
	*libc.As[int16](result_symbol939) = 8
	v344 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end940 = libc.Ptr(&libc.As[TSLexer](v344).F3)
	v345 = *libc.As[unsafe.Pointer](mark_end940)
	v346 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v345)(v346)
	v347 = *libc.As[int32](lookahead)
	cmp941 = v347 == 9
	if cmp941 {
		goto if_then952
	} else {
		goto lor_lhs_false943
	}

lor_lhs_false943:
	v348 = *libc.As[int32](lookahead)
	cmp944 = v348 == 10
	if cmp944 {
		goto if_then952
	} else {
		goto lor_lhs_false946
	}

lor_lhs_false946:
	v349 = *libc.As[int32](lookahead)
	cmp947 = v349 == 13
	if cmp947 {
		goto if_then952
	} else {
		goto lor_lhs_false949
	}

lor_lhs_false949:
	v350 = *libc.As[int32](lookahead)
	cmp950 = v350 == 32
	if cmp950 {
		goto if_then952
	} else {
		goto if_end953
	}

if_then952:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end953:
	v351 = *libc.As[byte](result)
	loadedv954 = (v351 & 1) != 0
	*libc.As[bool](retval) = loadedv954
	goto _return

sw_bb955:
	*libc.As[byte](result) = 1
	v352 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol956 = libc.Ptr(&libc.As[TSLexer](v352).F1)
	*libc.As[int16](result_symbol956) = 9
	v353 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end957 = libc.Ptr(&libc.As[TSLexer](v353).F3)
	v354 = *libc.As[unsafe.Pointer](mark_end957)
	v355 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v354)(v355)
	v356 = *libc.As[int32](lookahead)
	cmp958 = v356 == 40
	if cmp958 {
		goto if_then960
	} else {
		goto if_end961
	}

if_then960:
	*libc.As[int16](state_addr) = 80
	goto next_state

if_end961:
	v357 = *libc.As[int32](lookahead)
	cmp962 = v357 == 9
	if cmp962 {
		goto if_then970
	} else {
		goto lor_lhs_false964
	}

lor_lhs_false964:
	v358 = *libc.As[int32](lookahead)
	cmp965 = v358 == 13
	if cmp965 {
		goto if_then970
	} else {
		goto lor_lhs_false967
	}

lor_lhs_false967:
	v359 = *libc.As[int32](lookahead)
	cmp968 = v359 == 32
	if cmp968 {
		goto if_then970
	} else {
		goto if_end971
	}

if_then970:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end971:
	v360 = *libc.As[int32](lookahead)
	cmp972 = v360 != 0
	if cmp972 {
		goto land_lhs_true974
	} else {
		goto if_end984
	}

land_lhs_true974:
	v361 = *libc.As[int32](lookahead)
	cmp975 = v361 != 10
	if cmp975 {
		goto land_lhs_true977
	} else {
		goto if_end984
	}

land_lhs_true977:
	v362 = *libc.As[int32](lookahead)
	cmp978 = v362 != 35
	if cmp978 {
		goto land_lhs_true980
	} else {
		goto if_end984
	}

land_lhs_true980:
	v363 = *libc.As[int32](lookahead)
	cmp981 = v363 != 60
	if cmp981 {
		goto if_then983
	} else {
		goto if_end984
	}

if_then983:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end984:
	v364 = *libc.As[byte](result)
	loadedv985 = (v364 & 1) != 0
	*libc.As[bool](retval) = loadedv985
	goto _return

sw_bb986:
	*libc.As[byte](result) = 1
	v365 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol987 = libc.Ptr(&libc.As[TSLexer](v365).F1)
	*libc.As[int16](result_symbol987) = 9
	v366 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end988 = libc.Ptr(&libc.As[TSLexer](v366).F3)
	v367 = *libc.As[unsafe.Pointer](mark_end988)
	v368 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v367)(v368)
	v369 = *libc.As[int32](lookahead)
	cmp989 = v369 == 41
	if cmp989 {
		goto if_then991
	} else {
		goto if_end992
	}

if_then991:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end992:
	v370 = *libc.As[int32](lookahead)
	cmp993 = v370 != 0
	if cmp993 {
		goto land_lhs_true995
	} else {
		goto if_end1005
	}

land_lhs_true995:
	v371 = *libc.As[int32](lookahead)
	cmp996 = v371 != 10
	if cmp996 {
		goto land_lhs_true998
	} else {
		goto if_end1005
	}

land_lhs_true998:
	v372 = *libc.As[int32](lookahead)
	cmp999 = v372 != 35
	if cmp999 {
		goto land_lhs_true1001
	} else {
		goto if_end1005
	}

land_lhs_true1001:
	v373 = *libc.As[int32](lookahead)
	cmp1002 = v373 != 60
	if cmp1002 {
		goto if_then1004
	} else {
		goto if_end1005
	}

if_then1004:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end1005:
	v374 = *libc.As[byte](result)
	loadedv1006 = (v374 & 1) != 0
	*libc.As[bool](retval) = loadedv1006
	goto _return

sw_bb1007:
	*libc.As[byte](result) = 1
	v375 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1008 = libc.Ptr(&libc.As[TSLexer](v375).F1)
	*libc.As[int16](result_symbol1008) = 9
	v376 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1009 = libc.Ptr(&libc.As[TSLexer](v376).F3)
	v377 = *libc.As[unsafe.Pointer](mark_end1009)
	v378 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v377)(v378)
	v379 = *libc.As[int32](lookahead)
	cmp1010 = v379 == 97
	if cmp1010 {
		goto if_then1012
	} else {
		goto if_end1013
	}

if_then1012:
	*libc.As[int16](state_addr) = 81
	goto next_state

if_end1013:
	v380 = *libc.As[int32](lookahead)
	cmp1014 = v380 != 0
	if cmp1014 {
		goto land_lhs_true1016
	} else {
		goto if_end1026
	}

land_lhs_true1016:
	v381 = *libc.As[int32](lookahead)
	cmp1017 = v381 != 10
	if cmp1017 {
		goto land_lhs_true1019
	} else {
		goto if_end1026
	}

land_lhs_true1019:
	v382 = *libc.As[int32](lookahead)
	cmp1020 = v382 != 35
	if cmp1020 {
		goto land_lhs_true1022
	} else {
		goto if_end1026
	}

land_lhs_true1022:
	v383 = *libc.As[int32](lookahead)
	cmp1023 = v383 != 60
	if cmp1023 {
		goto if_then1025
	} else {
		goto if_end1026
	}

if_then1025:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end1026:
	v384 = *libc.As[byte](result)
	loadedv1027 = (v384 & 1) != 0
	*libc.As[bool](retval) = loadedv1027
	goto _return

sw_bb1028:
	*libc.As[byte](result) = 1
	v385 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1029 = libc.Ptr(&libc.As[TSLexer](v385).F1)
	*libc.As[int16](result_symbol1029) = 9
	v386 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1030 = libc.Ptr(&libc.As[TSLexer](v386).F3)
	v387 = *libc.As[unsafe.Pointer](mark_end1030)
	v388 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v387)(v388)
	v389 = *libc.As[int32](lookahead)
	cmp1031 = v389 == 98
	if cmp1031 {
		goto if_then1033
	} else {
		goto if_end1034
	}

if_then1033:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end1034:
	v390 = *libc.As[int32](lookahead)
	cmp1035 = v390 != 0
	if cmp1035 {
		goto land_lhs_true1037
	} else {
		goto if_end1047
	}

land_lhs_true1037:
	v391 = *libc.As[int32](lookahead)
	cmp1038 = v391 != 10
	if cmp1038 {
		goto land_lhs_true1040
	} else {
		goto if_end1047
	}

land_lhs_true1040:
	v392 = *libc.As[int32](lookahead)
	cmp1041 = v392 != 35
	if cmp1041 {
		goto land_lhs_true1043
	} else {
		goto if_end1047
	}

land_lhs_true1043:
	v393 = *libc.As[int32](lookahead)
	cmp1044 = v393 != 60
	if cmp1044 {
		goto if_then1046
	} else {
		goto if_end1047
	}

if_then1046:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end1047:
	v394 = *libc.As[byte](result)
	loadedv1048 = (v394 & 1) != 0
	*libc.As[bool](retval) = loadedv1048
	goto _return

sw_bb1049:
	*libc.As[byte](result) = 1
	v395 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1050 = libc.Ptr(&libc.As[TSLexer](v395).F1)
	*libc.As[int16](result_symbol1050) = 9
	v396 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1051 = libc.Ptr(&libc.As[TSLexer](v396).F3)
	v397 = *libc.As[unsafe.Pointer](mark_end1051)
	v398 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v397)(v398)
	v399 = *libc.As[int32](lookahead)
	cmp1052 = v399 == 100
	if cmp1052 {
		goto if_then1054
	} else {
		goto if_end1055
	}

if_then1054:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end1055:
	v400 = *libc.As[int32](lookahead)
	cmp1056 = v400 != 0
	if cmp1056 {
		goto land_lhs_true1058
	} else {
		goto if_end1068
	}

land_lhs_true1058:
	v401 = *libc.As[int32](lookahead)
	cmp1059 = v401 != 10
	if cmp1059 {
		goto land_lhs_true1061
	} else {
		goto if_end1068
	}

land_lhs_true1061:
	v402 = *libc.As[int32](lookahead)
	cmp1062 = v402 != 35
	if cmp1062 {
		goto land_lhs_true1064
	} else {
		goto if_end1068
	}

land_lhs_true1064:
	v403 = *libc.As[int32](lookahead)
	cmp1065 = v403 != 60
	if cmp1065 {
		goto if_then1067
	} else {
		goto if_end1068
	}

if_then1067:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end1068:
	v404 = *libc.As[byte](result)
	loadedv1069 = (v404 & 1) != 0
	*libc.As[bool](retval) = loadedv1069
	goto _return

sw_bb1070:
	*libc.As[byte](result) = 1
	v405 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1071 = libc.Ptr(&libc.As[TSLexer](v405).F1)
	*libc.As[int16](result_symbol1071) = 9
	v406 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1072 = libc.Ptr(&libc.As[TSLexer](v406).F3)
	v407 = *libc.As[unsafe.Pointer](mark_end1072)
	v408 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v407)(v408)
	v409 = *libc.As[int32](lookahead)
	cmp1073 = v409 != 0
	if cmp1073 {
		goto land_lhs_true1075
	} else {
		goto if_end1085
	}

land_lhs_true1075:
	v410 = *libc.As[int32](lookahead)
	cmp1076 = v410 != 10
	if cmp1076 {
		goto land_lhs_true1078
	} else {
		goto if_end1085
	}

land_lhs_true1078:
	v411 = *libc.As[int32](lookahead)
	cmp1079 = v411 != 35
	if cmp1079 {
		goto land_lhs_true1081
	} else {
		goto if_end1085
	}

land_lhs_true1081:
	v412 = *libc.As[int32](lookahead)
	cmp1082 = v412 != 60
	if cmp1082 {
		goto if_then1084
	} else {
		goto if_end1085
	}

if_then1084:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end1085:
	v413 = *libc.As[byte](result)
	loadedv1086 = (v413 & 1) != 0
	*libc.As[bool](retval) = loadedv1086
	goto _return

sw_bb1087:
	*libc.As[byte](result) = 1
	v414 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1088 = libc.Ptr(&libc.As[TSLexer](v414).F1)
	*libc.As[int16](result_symbol1088) = 10
	v415 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1089 = libc.Ptr(&libc.As[TSLexer](v415).F3)
	v416 = *libc.As[unsafe.Pointer](mark_end1089)
	v417 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v416)(v417)
	v418 = *libc.As[byte](result)
	loadedv1090 = (v418 & 1) != 0
	*libc.As[bool](retval) = loadedv1090
	goto _return

sw_bb1091:
	*libc.As[byte](result) = 1
	v419 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1092 = libc.Ptr(&libc.As[TSLexer](v419).F1)
	*libc.As[int16](result_symbol1092) = 10
	v420 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1093 = libc.Ptr(&libc.As[TSLexer](v420).F3)
	v421 = *libc.As[unsafe.Pointer](mark_end1093)
	v422 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v421)(v422)
	v423 = *libc.As[int32](lookahead)
	cmp1094 = v423 != 0
	if cmp1094 {
		goto land_lhs_true1096
	} else {
		goto if_end1106
	}

land_lhs_true1096:
	v424 = *libc.As[int32](lookahead)
	cmp1097 = v424 != 10
	if cmp1097 {
		goto land_lhs_true1099
	} else {
		goto if_end1106
	}

land_lhs_true1099:
	v425 = *libc.As[int32](lookahead)
	cmp1100 = v425 != 35
	if cmp1100 {
		goto land_lhs_true1102
	} else {
		goto if_end1106
	}

land_lhs_true1102:
	v426 = *libc.As[int32](lookahead)
	cmp1103 = v426 != 60
	if cmp1103 {
		goto if_then1105
	} else {
		goto if_end1106
	}

if_then1105:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end1106:
	v427 = *libc.As[byte](result)
	loadedv1107 = (v427 & 1) != 0
	*libc.As[bool](retval) = loadedv1107
	goto _return

sw_bb1108:
	*libc.As[byte](result) = 1
	v428 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1109 = libc.Ptr(&libc.As[TSLexer](v428).F1)
	*libc.As[int16](result_symbol1109) = 11
	v429 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1110 = libc.Ptr(&libc.As[TSLexer](v429).F3)
	v430 = *libc.As[unsafe.Pointer](mark_end1110)
	v431 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v430)(v431)
	v432 = *libc.As[byte](result)
	loadedv1111 = (v432 & 1) != 0
	*libc.As[bool](retval) = loadedv1111
	goto _return

sw_bb1112:
	*libc.As[byte](result) = 1
	v433 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1113 = libc.Ptr(&libc.As[TSLexer](v433).F1)
	*libc.As[int16](result_symbol1113) = 12
	v434 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1114 = libc.Ptr(&libc.As[TSLexer](v434).F3)
	v435 = *libc.As[unsafe.Pointer](mark_end1114)
	v436 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v435)(v436)
	v437 = *libc.As[byte](result)
	loadedv1115 = (v437 & 1) != 0
	*libc.As[bool](retval) = loadedv1115
	goto _return

sw_bb1116:
	*libc.As[byte](result) = 1
	v438 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1117 = libc.Ptr(&libc.As[TSLexer](v438).F1)
	*libc.As[int16](result_symbol1117) = 12
	v439 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1118 = libc.Ptr(&libc.As[TSLexer](v439).F3)
	v440 = *libc.As[unsafe.Pointer](mark_end1118)
	v441 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v440)(v441)
	v442 = *libc.As[int32](lookahead)
	cmp1119 = v442 != 0
	if cmp1119 {
		goto land_lhs_true1121
	} else {
		goto if_end1125
	}

land_lhs_true1121:
	v443 = *libc.As[int32](lookahead)
	cmp1122 = v443 != 10
	if cmp1122 {
		goto if_then1124
	} else {
		goto if_end1125
	}

if_then1124:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end1125:
	v444 = *libc.As[byte](result)
	loadedv1126 = (v444 & 1) != 0
	*libc.As[bool](retval) = loadedv1126
	goto _return

sw_bb1127:
	*libc.As[byte](result) = 1
	v445 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1128 = libc.Ptr(&libc.As[TSLexer](v445).F1)
	*libc.As[int16](result_symbol1128) = 13
	v446 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1129 = libc.Ptr(&libc.As[TSLexer](v446).F3)
	v447 = *libc.As[unsafe.Pointer](mark_end1129)
	v448 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v447)(v448)
	v449 = *libc.As[byte](result)
	loadedv1130 = (v449 & 1) != 0
	*libc.As[bool](retval) = loadedv1130
	goto _return

sw_bb1131:
	*libc.As[byte](result) = 1
	v450 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1132 = libc.Ptr(&libc.As[TSLexer](v450).F1)
	*libc.As[int16](result_symbol1132) = 14
	v451 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1133 = libc.Ptr(&libc.As[TSLexer](v451).F3)
	v452 = *libc.As[unsafe.Pointer](mark_end1133)
	v453 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v452)(v453)
	v454 = *libc.As[byte](result)
	loadedv1134 = (v454 & 1) != 0
	*libc.As[bool](retval) = loadedv1134
	goto _return

sw_bb1135:
	*libc.As[byte](result) = 1
	v455 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1136 = libc.Ptr(&libc.As[TSLexer](v455).F1)
	*libc.As[int16](result_symbol1136) = 15
	v456 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1137 = libc.Ptr(&libc.As[TSLexer](v456).F3)
	v457 = *libc.As[unsafe.Pointer](mark_end1137)
	v458 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v457)(v458)
	v459 = *libc.As[int32](lookahead)
	cmp1138 = 48 <= v459
	if cmp1138 {
		goto land_lhs_true1140
	} else {
		goto lor_lhs_false1143
	}

land_lhs_true1140:
	v460 = *libc.As[int32](lookahead)
	cmp1141 = v460 <= 57
	if cmp1141 {
		goto if_then1155
	} else {
		goto lor_lhs_false1143
	}

lor_lhs_false1143:
	v461 = *libc.As[int32](lookahead)
	cmp1144 = 65 <= v461
	if cmp1144 {
		goto land_lhs_true1146
	} else {
		goto lor_lhs_false1149
	}

land_lhs_true1146:
	v462 = *libc.As[int32](lookahead)
	cmp1147 = v462 <= 70
	if cmp1147 {
		goto if_then1155
	} else {
		goto lor_lhs_false1149
	}

lor_lhs_false1149:
	v463 = *libc.As[int32](lookahead)
	cmp1150 = 97 <= v463
	if cmp1150 {
		goto land_lhs_true1152
	} else {
		goto if_end1156
	}

land_lhs_true1152:
	v464 = *libc.As[int32](lookahead)
	cmp1153 = v464 <= 102
	if cmp1153 {
		goto if_then1155
	} else {
		goto if_end1156
	}

if_then1155:
	*libc.As[int16](state_addr) = 90
	goto next_state

if_end1156:
	v465 = *libc.As[byte](result)
	loadedv1157 = (v465 & 1) != 0
	*libc.As[bool](retval) = loadedv1157
	goto _return

sw_bb1158:
	*libc.As[byte](result) = 1
	v466 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1159 = libc.Ptr(&libc.As[TSLexer](v466).F1)
	*libc.As[int16](result_symbol1159) = 16
	v467 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1160 = libc.Ptr(&libc.As[TSLexer](v467).F3)
	v468 = *libc.As[unsafe.Pointer](mark_end1160)
	v469 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v468)(v469)
	v470 = *libc.As[byte](result)
	loadedv1161 = (v470 & 1) != 0
	*libc.As[bool](retval) = loadedv1161
	goto _return

sw_bb1162:
	*libc.As[byte](result) = 1
	v471 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1163 = libc.Ptr(&libc.As[TSLexer](v471).F1)
	*libc.As[int16](result_symbol1163) = 17
	v472 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1164 = libc.Ptr(&libc.As[TSLexer](v472).F3)
	v473 = *libc.As[unsafe.Pointer](mark_end1164)
	v474 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v473)(v474)
	v475 = *libc.As[int32](lookahead)
	cmp1165 = v475 == 32
	if cmp1165 {
		goto if_then1167
	} else {
		goto if_end1168
	}

if_then1167:
	*libc.As[int16](state_addr) = 92
	goto next_state

if_end1168:
	v476 = *libc.As[int32](lookahead)
	cmp1169 = v476 == 9
	if cmp1169 {
		goto if_then1177
	} else {
		goto lor_lhs_false1171
	}

lor_lhs_false1171:
	v477 = *libc.As[int32](lookahead)
	cmp1172 = v477 == 10
	if cmp1172 {
		goto if_then1177
	} else {
		goto lor_lhs_false1174
	}

lor_lhs_false1174:
	v478 = *libc.As[int32](lookahead)
	cmp1175 = v478 == 13
	if cmp1175 {
		goto if_then1177
	} else {
		goto if_end1178
	}

if_then1177:
	*libc.As[int16](state_addr) = 75
	goto next_state

if_end1178:
	v479 = *libc.As[byte](result)
	loadedv1179 = (v479 & 1) != 0
	*libc.As[bool](retval) = loadedv1179
	goto _return

sw_bb1180:
	*libc.As[byte](result) = 1
	v480 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1181 = libc.Ptr(&libc.As[TSLexer](v480).F1)
	*libc.As[int16](result_symbol1181) = 18
	v481 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1182 = libc.Ptr(&libc.As[TSLexer](v481).F3)
	v482 = *libc.As[unsafe.Pointer](mark_end1182)
	v483 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v482)(v483)
	v484 = *libc.As[int32](lookahead)
	cmp1183 = v484 == 105
	if cmp1183 {
		goto if_then1185
	} else {
		goto if_end1186
	}

if_then1185:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end1186:
	v485 = *libc.As[int32](lookahead)
	cmp1187 = 48 <= v485
	if cmp1187 {
		goto land_lhs_true1189
	} else {
		goto lor_lhs_false1192
	}

land_lhs_true1189:
	v486 = *libc.As[int32](lookahead)
	cmp1190 = v486 <= 57
	if cmp1190 {
		goto if_then1204
	} else {
		goto lor_lhs_false1192
	}

lor_lhs_false1192:
	v487 = *libc.As[int32](lookahead)
	cmp1193 = 65 <= v487
	if cmp1193 {
		goto land_lhs_true1195
	} else {
		goto lor_lhs_false1198
	}

land_lhs_true1195:
	v488 = *libc.As[int32](lookahead)
	cmp1196 = v488 <= 70
	if cmp1196 {
		goto if_then1204
	} else {
		goto lor_lhs_false1198
	}

lor_lhs_false1198:
	v489 = *libc.As[int32](lookahead)
	cmp1199 = 97 <= v489
	if cmp1199 {
		goto land_lhs_true1201
	} else {
		goto if_end1205
	}

land_lhs_true1201:
	v490 = *libc.As[int32](lookahead)
	cmp1202 = v490 <= 102
	if cmp1202 {
		goto if_then1204
	} else {
		goto if_end1205
	}

if_then1204:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end1205:
	v491 = *libc.As[byte](result)
	loadedv1206 = (v491 & 1) != 0
	*libc.As[bool](retval) = loadedv1206
	goto _return

sw_bb1207:
	*libc.As[byte](result) = 1
	v492 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1208 = libc.Ptr(&libc.As[TSLexer](v492).F1)
	*libc.As[int16](result_symbol1208) = 18
	v493 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1209 = libc.Ptr(&libc.As[TSLexer](v493).F3)
	v494 = *libc.As[unsafe.Pointer](mark_end1209)
	v495 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v494)(v495)
	v496 = *libc.As[int32](lookahead)
	cmp1210 = v496 == 105
	if cmp1210 {
		goto if_then1212
	} else {
		goto if_end1213
	}

if_then1212:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end1213:
	v497 = *libc.As[int32](lookahead)
	cmp1214 = 48 <= v497
	if cmp1214 {
		goto land_lhs_true1216
	} else {
		goto lor_lhs_false1219
	}

land_lhs_true1216:
	v498 = *libc.As[int32](lookahead)
	cmp1217 = v498 <= 57
	if cmp1217 {
		goto if_then1231
	} else {
		goto lor_lhs_false1219
	}

lor_lhs_false1219:
	v499 = *libc.As[int32](lookahead)
	cmp1220 = 65 <= v499
	if cmp1220 {
		goto land_lhs_true1222
	} else {
		goto lor_lhs_false1225
	}

land_lhs_true1222:
	v500 = *libc.As[int32](lookahead)
	cmp1223 = v500 <= 70
	if cmp1223 {
		goto if_then1231
	} else {
		goto lor_lhs_false1225
	}

lor_lhs_false1225:
	v501 = *libc.As[int32](lookahead)
	cmp1226 = 97 <= v501
	if cmp1226 {
		goto land_lhs_true1228
	} else {
		goto if_end1232
	}

land_lhs_true1228:
	v502 = *libc.As[int32](lookahead)
	cmp1229 = v502 <= 102
	if cmp1229 {
		goto if_then1231
	} else {
		goto if_end1232
	}

if_then1231:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end1232:
	v503 = *libc.As[byte](result)
	loadedv1233 = (v503 & 1) != 0
	*libc.As[bool](retval) = loadedv1233
	goto _return

sw_bb1234:
	*libc.As[byte](result) = 1
	v504 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1235 = libc.Ptr(&libc.As[TSLexer](v504).F1)
	*libc.As[int16](result_symbol1235) = 18
	v505 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1236 = libc.Ptr(&libc.As[TSLexer](v505).F3)
	v506 = *libc.As[unsafe.Pointer](mark_end1236)
	v507 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v506)(v507)
	v508 = *libc.As[int32](lookahead)
	cmp1237 = v508 == 105
	if cmp1237 {
		goto if_then1239
	} else {
		goto if_end1240
	}

if_then1239:
	*libc.As[int16](state_addr) = 113
	goto next_state

if_end1240:
	v509 = *libc.As[int32](lookahead)
	cmp1241 = v509 == 43
	if cmp1241 {
		goto if_then1249
	} else {
		goto lor_lhs_false1243
	}

lor_lhs_false1243:
	v510 = *libc.As[int32](lookahead)
	cmp1244 = v510 == 45
	if cmp1244 {
		goto if_then1249
	} else {
		goto lor_lhs_false1246
	}

lor_lhs_false1246:
	v511 = *libc.As[int32](lookahead)
	cmp1247 = v511 == 47
	if cmp1247 {
		goto if_then1249
	} else {
		goto if_end1250
	}

if_then1249:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1250:
	v512 = *libc.As[int32](lookahead)
	cmp1251 = 36 <= v512
	if cmp1251 {
		goto land_lhs_true1253
	} else {
		goto lor_lhs_false1256
	}

land_lhs_true1253:
	v513 = *libc.As[int32](lookahead)
	cmp1254 = v513 <= 41
	if cmp1254 {
		goto if_then1259
	} else {
		goto lor_lhs_false1256
	}

lor_lhs_false1256:
	v514 = *libc.As[int32](lookahead)
	cmp1257 = v514 == 64
	if cmp1257 {
		goto if_then1259
	} else {
		goto if_end1260
	}

if_then1259:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1260:
	v515 = *libc.As[int32](lookahead)
	cmp1261 = 48 <= v515
	if cmp1261 {
		goto land_lhs_true1263
	} else {
		goto lor_lhs_false1266
	}

land_lhs_true1263:
	v516 = *libc.As[int32](lookahead)
	cmp1264 = v516 <= 57
	if cmp1264 {
		goto if_then1278
	} else {
		goto lor_lhs_false1266
	}

lor_lhs_false1266:
	v517 = *libc.As[int32](lookahead)
	cmp1267 = 65 <= v517
	if cmp1267 {
		goto land_lhs_true1269
	} else {
		goto lor_lhs_false1272
	}

land_lhs_true1269:
	v518 = *libc.As[int32](lookahead)
	cmp1270 = v518 <= 70
	if cmp1270 {
		goto if_then1278
	} else {
		goto lor_lhs_false1272
	}

lor_lhs_false1272:
	v519 = *libc.As[int32](lookahead)
	cmp1273 = 97 <= v519
	if cmp1273 {
		goto land_lhs_true1275
	} else {
		goto if_end1279
	}

land_lhs_true1275:
	v520 = *libc.As[int32](lookahead)
	cmp1276 = v520 <= 102
	if cmp1276 {
		goto if_then1278
	} else {
		goto if_end1279
	}

if_then1278:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end1279:
	v521 = *libc.As[int32](lookahead)
	cmp1280 = v521 == 46
	if cmp1280 {
		goto if_then1297
	} else {
		goto lor_lhs_false1282
	}

lor_lhs_false1282:
	v522 = *libc.As[int32](lookahead)
	cmp1283 = 71 <= v522
	if cmp1283 {
		goto land_lhs_true1285
	} else {
		goto lor_lhs_false1288
	}

land_lhs_true1285:
	v523 = *libc.As[int32](lookahead)
	cmp1286 = v523 <= 90
	if cmp1286 {
		goto if_then1297
	} else {
		goto lor_lhs_false1288
	}

lor_lhs_false1288:
	v524 = *libc.As[int32](lookahead)
	cmp1289 = v524 == 95
	if cmp1289 {
		goto if_then1297
	} else {
		goto lor_lhs_false1291
	}

lor_lhs_false1291:
	v525 = *libc.As[int32](lookahead)
	cmp1292 = 103 <= v525
	if cmp1292 {
		goto land_lhs_true1294
	} else {
		goto if_end1298
	}

land_lhs_true1294:
	v526 = *libc.As[int32](lookahead)
	cmp1295 = v526 <= 122
	if cmp1295 {
		goto if_then1297
	} else {
		goto if_end1298
	}

if_then1297:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1298:
	v527 = *libc.As[byte](result)
	loadedv1299 = (v527 & 1) != 0
	*libc.As[bool](retval) = loadedv1299
	goto _return

sw_bb1300:
	*libc.As[byte](result) = 1
	v528 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1301 = libc.Ptr(&libc.As[TSLexer](v528).F1)
	*libc.As[int16](result_symbol1301) = 18
	v529 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1302 = libc.Ptr(&libc.As[TSLexer](v529).F3)
	v530 = *libc.As[unsafe.Pointer](mark_end1302)
	v531 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v530)(v531)
	v532 = *libc.As[int32](lookahead)
	cmp1303 = v532 == 105
	if cmp1303 {
		goto if_then1305
	} else {
		goto if_end1306
	}

if_then1305:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end1306:
	v533 = *libc.As[int32](lookahead)
	cmp1307 = 48 <= v533
	if cmp1307 {
		goto land_lhs_true1309
	} else {
		goto lor_lhs_false1312
	}

land_lhs_true1309:
	v534 = *libc.As[int32](lookahead)
	cmp1310 = v534 <= 57
	if cmp1310 {
		goto if_then1324
	} else {
		goto lor_lhs_false1312
	}

lor_lhs_false1312:
	v535 = *libc.As[int32](lookahead)
	cmp1313 = 65 <= v535
	if cmp1313 {
		goto land_lhs_true1315
	} else {
		goto lor_lhs_false1318
	}

land_lhs_true1315:
	v536 = *libc.As[int32](lookahead)
	cmp1316 = v536 <= 70
	if cmp1316 {
		goto if_then1324
	} else {
		goto lor_lhs_false1318
	}

lor_lhs_false1318:
	v537 = *libc.As[int32](lookahead)
	cmp1319 = 97 <= v537
	if cmp1319 {
		goto land_lhs_true1321
	} else {
		goto if_end1325
	}

land_lhs_true1321:
	v538 = *libc.As[int32](lookahead)
	cmp1322 = v538 <= 102
	if cmp1322 {
		goto if_then1324
	} else {
		goto if_end1325
	}

if_then1324:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end1325:
	v539 = *libc.As[byte](result)
	loadedv1326 = (v539 & 1) != 0
	*libc.As[bool](retval) = loadedv1326
	goto _return

sw_bb1327:
	*libc.As[byte](result) = 1
	v540 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1328 = libc.Ptr(&libc.As[TSLexer](v540).F1)
	*libc.As[int16](result_symbol1328) = 18
	v541 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1329 = libc.Ptr(&libc.As[TSLexer](v541).F3)
	v542 = *libc.As[unsafe.Pointer](mark_end1329)
	v543 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v542)(v543)
	v544 = *libc.As[int32](lookahead)
	cmp1330 = v544 == 105
	if cmp1330 {
		goto if_then1332
	} else {
		goto if_end1333
	}

if_then1332:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end1333:
	v545 = *libc.As[int32](lookahead)
	cmp1334 = v545 == 111
	if cmp1334 {
		goto if_then1336
	} else {
		goto if_end1337
	}

if_then1336:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end1337:
	v546 = *libc.As[int32](lookahead)
	cmp1338 = 48 <= v546
	if cmp1338 {
		goto land_lhs_true1340
	} else {
		goto lor_lhs_false1343
	}

land_lhs_true1340:
	v547 = *libc.As[int32](lookahead)
	cmp1341 = v547 <= 57
	if cmp1341 {
		goto if_then1355
	} else {
		goto lor_lhs_false1343
	}

lor_lhs_false1343:
	v548 = *libc.As[int32](lookahead)
	cmp1344 = 65 <= v548
	if cmp1344 {
		goto land_lhs_true1346
	} else {
		goto lor_lhs_false1349
	}

land_lhs_true1346:
	v549 = *libc.As[int32](lookahead)
	cmp1347 = v549 <= 70
	if cmp1347 {
		goto if_then1355
	} else {
		goto lor_lhs_false1349
	}

lor_lhs_false1349:
	v550 = *libc.As[int32](lookahead)
	cmp1350 = 97 <= v550
	if cmp1350 {
		goto land_lhs_true1352
	} else {
		goto if_end1356
	}

land_lhs_true1352:
	v551 = *libc.As[int32](lookahead)
	cmp1353 = v551 <= 102
	if cmp1353 {
		goto if_then1355
	} else {
		goto if_end1356
	}

if_then1355:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end1356:
	v552 = *libc.As[byte](result)
	loadedv1357 = (v552 & 1) != 0
	*libc.As[bool](retval) = loadedv1357
	goto _return

sw_bb1358:
	*libc.As[byte](result) = 1
	v553 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1359 = libc.Ptr(&libc.As[TSLexer](v553).F1)
	*libc.As[int16](result_symbol1359) = 18
	v554 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1360 = libc.Ptr(&libc.As[TSLexer](v554).F3)
	v555 = *libc.As[unsafe.Pointer](mark_end1360)
	v556 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v555)(v556)
	v557 = *libc.As[int32](lookahead)
	cmp1361 = v557 == 104
	if cmp1361 {
		goto if_then1366
	} else {
		goto lor_lhs_false1363
	}

lor_lhs_false1363:
	v558 = *libc.As[int32](lookahead)
	cmp1364 = v558 == 120
	if cmp1364 {
		goto if_then1366
	} else {
		goto if_end1367
	}

if_then1366:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end1367:
	v559 = *libc.As[int32](lookahead)
	cmp1368 = 48 <= v559
	if cmp1368 {
		goto land_lhs_true1370
	} else {
		goto lor_lhs_false1373
	}

land_lhs_true1370:
	v560 = *libc.As[int32](lookahead)
	cmp1371 = v560 <= 57
	if cmp1371 {
		goto if_then1385
	} else {
		goto lor_lhs_false1373
	}

lor_lhs_false1373:
	v561 = *libc.As[int32](lookahead)
	cmp1374 = 65 <= v561
	if cmp1374 {
		goto land_lhs_true1376
	} else {
		goto lor_lhs_false1379
	}

land_lhs_true1376:
	v562 = *libc.As[int32](lookahead)
	cmp1377 = v562 <= 70
	if cmp1377 {
		goto if_then1385
	} else {
		goto lor_lhs_false1379
	}

lor_lhs_false1379:
	v563 = *libc.As[int32](lookahead)
	cmp1380 = 97 <= v563
	if cmp1380 {
		goto land_lhs_true1382
	} else {
		goto if_end1386
	}

land_lhs_true1382:
	v564 = *libc.As[int32](lookahead)
	cmp1383 = v564 <= 102
	if cmp1383 {
		goto if_then1385
	} else {
		goto if_end1386
	}

if_then1385:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end1386:
	v565 = *libc.As[byte](result)
	loadedv1387 = (v565 & 1) != 0
	*libc.As[bool](retval) = loadedv1387
	goto _return

sw_bb1388:
	*libc.As[byte](result) = 1
	v566 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1389 = libc.Ptr(&libc.As[TSLexer](v566).F1)
	*libc.As[int16](result_symbol1389) = 18
	v567 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1390 = libc.Ptr(&libc.As[TSLexer](v567).F3)
	v568 = *libc.As[unsafe.Pointer](mark_end1390)
	v569 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v568)(v569)
	v570 = *libc.As[int32](lookahead)
	cmp1391 = v570 == 43
	if cmp1391 {
		goto if_then1399
	} else {
		goto lor_lhs_false1393
	}

lor_lhs_false1393:
	v571 = *libc.As[int32](lookahead)
	cmp1394 = v571 == 45
	if cmp1394 {
		goto if_then1399
	} else {
		goto lor_lhs_false1396
	}

lor_lhs_false1396:
	v572 = *libc.As[int32](lookahead)
	cmp1397 = v572 == 47
	if cmp1397 {
		goto if_then1399
	} else {
		goto if_end1400
	}

if_then1399:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1400:
	v573 = *libc.As[int32](lookahead)
	cmp1401 = 36 <= v573
	if cmp1401 {
		goto land_lhs_true1403
	} else {
		goto lor_lhs_false1406
	}

land_lhs_true1403:
	v574 = *libc.As[int32](lookahead)
	cmp1404 = v574 <= 41
	if cmp1404 {
		goto if_then1409
	} else {
		goto lor_lhs_false1406
	}

lor_lhs_false1406:
	v575 = *libc.As[int32](lookahead)
	cmp1407 = v575 == 64
	if cmp1407 {
		goto if_then1409
	} else {
		goto if_end1410
	}

if_then1409:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1410:
	v576 = *libc.As[int32](lookahead)
	cmp1411 = 48 <= v576
	if cmp1411 {
		goto land_lhs_true1413
	} else {
		goto lor_lhs_false1416
	}

land_lhs_true1413:
	v577 = *libc.As[int32](lookahead)
	cmp1414 = v577 <= 57
	if cmp1414 {
		goto if_then1428
	} else {
		goto lor_lhs_false1416
	}

lor_lhs_false1416:
	v578 = *libc.As[int32](lookahead)
	cmp1417 = 65 <= v578
	if cmp1417 {
		goto land_lhs_true1419
	} else {
		goto lor_lhs_false1422
	}

land_lhs_true1419:
	v579 = *libc.As[int32](lookahead)
	cmp1420 = v579 <= 70
	if cmp1420 {
		goto if_then1428
	} else {
		goto lor_lhs_false1422
	}

lor_lhs_false1422:
	v580 = *libc.As[int32](lookahead)
	cmp1423 = 97 <= v580
	if cmp1423 {
		goto land_lhs_true1425
	} else {
		goto if_end1429
	}

land_lhs_true1425:
	v581 = *libc.As[int32](lookahead)
	cmp1426 = v581 <= 102
	if cmp1426 {
		goto if_then1428
	} else {
		goto if_end1429
	}

if_then1428:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end1429:
	v582 = *libc.As[int32](lookahead)
	cmp1430 = v582 == 46
	if cmp1430 {
		goto if_then1447
	} else {
		goto lor_lhs_false1432
	}

lor_lhs_false1432:
	v583 = *libc.As[int32](lookahead)
	cmp1433 = 71 <= v583
	if cmp1433 {
		goto land_lhs_true1435
	} else {
		goto lor_lhs_false1438
	}

land_lhs_true1435:
	v584 = *libc.As[int32](lookahead)
	cmp1436 = v584 <= 90
	if cmp1436 {
		goto if_then1447
	} else {
		goto lor_lhs_false1438
	}

lor_lhs_false1438:
	v585 = *libc.As[int32](lookahead)
	cmp1439 = v585 == 95
	if cmp1439 {
		goto if_then1447
	} else {
		goto lor_lhs_false1441
	}

lor_lhs_false1441:
	v586 = *libc.As[int32](lookahead)
	cmp1442 = 103 <= v586
	if cmp1442 {
		goto land_lhs_true1444
	} else {
		goto if_end1448
	}

land_lhs_true1444:
	v587 = *libc.As[int32](lookahead)
	cmp1445 = v587 <= 122
	if cmp1445 {
		goto if_then1447
	} else {
		goto if_end1448
	}

if_then1447:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1448:
	v588 = *libc.As[byte](result)
	loadedv1449 = (v588 & 1) != 0
	*libc.As[bool](retval) = loadedv1449
	goto _return

sw_bb1450:
	*libc.As[byte](result) = 1
	v589 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1451 = libc.Ptr(&libc.As[TSLexer](v589).F1)
	*libc.As[int16](result_symbol1451) = 18
	v590 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1452 = libc.Ptr(&libc.As[TSLexer](v590).F3)
	v591 = *libc.As[unsafe.Pointer](mark_end1452)
	v592 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v591)(v592)
	v593 = *libc.As[int32](lookahead)
	cmp1453 = 48 <= v593
	if cmp1453 {
		goto land_lhs_true1455
	} else {
		goto if_end1459
	}

land_lhs_true1455:
	v594 = *libc.As[int32](lookahead)
	cmp1456 = v594 <= 57
	if cmp1456 {
		goto if_then1458
	} else {
		goto if_end1459
	}

if_then1458:
	*libc.As[int16](state_addr) = 100
	goto next_state

if_end1459:
	v595 = *libc.As[int32](lookahead)
	cmp1460 = 65 <= v595
	if cmp1460 {
		goto land_lhs_true1462
	} else {
		goto lor_lhs_false1465
	}

land_lhs_true1462:
	v596 = *libc.As[int32](lookahead)
	cmp1463 = v596 <= 70
	if cmp1463 {
		goto if_then1471
	} else {
		goto lor_lhs_false1465
	}

lor_lhs_false1465:
	v597 = *libc.As[int32](lookahead)
	cmp1466 = 97 <= v597
	if cmp1466 {
		goto land_lhs_true1468
	} else {
		goto if_end1472
	}

land_lhs_true1468:
	v598 = *libc.As[int32](lookahead)
	cmp1469 = v598 <= 102
	if cmp1469 {
		goto if_then1471
	} else {
		goto if_end1472
	}

if_then1471:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end1472:
	v599 = *libc.As[byte](result)
	loadedv1473 = (v599 & 1) != 0
	*libc.As[bool](retval) = loadedv1473
	goto _return

sw_bb1474:
	*libc.As[byte](result) = 1
	v600 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1475 = libc.Ptr(&libc.As[TSLexer](v600).F1)
	*libc.As[int16](result_symbol1475) = 18
	v601 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1476 = libc.Ptr(&libc.As[TSLexer](v601).F3)
	v602 = *libc.As[unsafe.Pointer](mark_end1476)
	v603 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v602)(v603)
	v604 = *libc.As[int32](lookahead)
	cmp1477 = 48 <= v604
	if cmp1477 {
		goto land_lhs_true1479
	} else {
		goto lor_lhs_false1482
	}

land_lhs_true1479:
	v605 = *libc.As[int32](lookahead)
	cmp1480 = v605 <= 57
	if cmp1480 {
		goto if_then1494
	} else {
		goto lor_lhs_false1482
	}

lor_lhs_false1482:
	v606 = *libc.As[int32](lookahead)
	cmp1483 = 65 <= v606
	if cmp1483 {
		goto land_lhs_true1485
	} else {
		goto lor_lhs_false1488
	}

land_lhs_true1485:
	v607 = *libc.As[int32](lookahead)
	cmp1486 = v607 <= 70
	if cmp1486 {
		goto if_then1494
	} else {
		goto lor_lhs_false1488
	}

lor_lhs_false1488:
	v608 = *libc.As[int32](lookahead)
	cmp1489 = 97 <= v608
	if cmp1489 {
		goto land_lhs_true1491
	} else {
		goto if_end1495
	}

land_lhs_true1491:
	v609 = *libc.As[int32](lookahead)
	cmp1492 = v609 <= 102
	if cmp1492 {
		goto if_then1494
	} else {
		goto if_end1495
	}

if_then1494:
	*libc.As[int16](state_addr) = 101
	goto next_state

if_end1495:
	v610 = *libc.As[byte](result)
	loadedv1496 = (v610 & 1) != 0
	*libc.As[bool](retval) = loadedv1496
	goto _return

sw_bb1497:
	*libc.As[byte](result) = 1
	v611 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1498 = libc.Ptr(&libc.As[TSLexer](v611).F1)
	*libc.As[int16](result_symbol1498) = 18
	v612 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1499 = libc.Ptr(&libc.As[TSLexer](v612).F3)
	v613 = *libc.As[unsafe.Pointer](mark_end1499)
	v614 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v613)(v614)
	v615 = *libc.As[int32](lookahead)
	cmp1500 = 48 <= v615
	if cmp1500 {
		goto land_lhs_true1502
	} else {
		goto lor_lhs_false1505
	}

land_lhs_true1502:
	v616 = *libc.As[int32](lookahead)
	cmp1503 = v616 <= 57
	if cmp1503 {
		goto if_then1517
	} else {
		goto lor_lhs_false1505
	}

lor_lhs_false1505:
	v617 = *libc.As[int32](lookahead)
	cmp1506 = 65 <= v617
	if cmp1506 {
		goto land_lhs_true1508
	} else {
		goto lor_lhs_false1511
	}

land_lhs_true1508:
	v618 = *libc.As[int32](lookahead)
	cmp1509 = v618 <= 70
	if cmp1509 {
		goto if_then1517
	} else {
		goto lor_lhs_false1511
	}

lor_lhs_false1511:
	v619 = *libc.As[int32](lookahead)
	cmp1512 = 97 <= v619
	if cmp1512 {
		goto land_lhs_true1514
	} else {
		goto if_end1518
	}

land_lhs_true1514:
	v620 = *libc.As[int32](lookahead)
	cmp1515 = v620 <= 102
	if cmp1515 {
		goto if_then1517
	} else {
		goto if_end1518
	}

if_then1517:
	*libc.As[int16](state_addr) = 102
	goto next_state

if_end1518:
	v621 = *libc.As[int32](lookahead)
	cmp1519 = v621 == 43
	if cmp1519 {
		goto if_then1542
	} else {
		goto lor_lhs_false1521
	}

lor_lhs_false1521:
	v622 = *libc.As[int32](lookahead)
	cmp1522 = 45 <= v622
	if cmp1522 {
		goto land_lhs_true1524
	} else {
		goto lor_lhs_false1527
	}

land_lhs_true1524:
	v623 = *libc.As[int32](lookahead)
	cmp1525 = v623 <= 47
	if cmp1525 {
		goto if_then1542
	} else {
		goto lor_lhs_false1527
	}

lor_lhs_false1527:
	v624 = *libc.As[int32](lookahead)
	cmp1528 = 71 <= v624
	if cmp1528 {
		goto land_lhs_true1530
	} else {
		goto lor_lhs_false1533
	}

land_lhs_true1530:
	v625 = *libc.As[int32](lookahead)
	cmp1531 = v625 <= 90
	if cmp1531 {
		goto if_then1542
	} else {
		goto lor_lhs_false1533
	}

lor_lhs_false1533:
	v626 = *libc.As[int32](lookahead)
	cmp1534 = v626 == 95
	if cmp1534 {
		goto if_then1542
	} else {
		goto lor_lhs_false1536
	}

lor_lhs_false1536:
	v627 = *libc.As[int32](lookahead)
	cmp1537 = 103 <= v627
	if cmp1537 {
		goto land_lhs_true1539
	} else {
		goto if_end1543
	}

land_lhs_true1539:
	v628 = *libc.As[int32](lookahead)
	cmp1540 = v628 <= 122
	if cmp1540 {
		goto if_then1542
	} else {
		goto if_end1543
	}

if_then1542:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1543:
	v629 = *libc.As[byte](result)
	loadedv1544 = (v629 & 1) != 0
	*libc.As[bool](retval) = loadedv1544
	goto _return

sw_bb1545:
	*libc.As[byte](result) = 1
	v630 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1546 = libc.Ptr(&libc.As[TSLexer](v630).F1)
	*libc.As[int16](result_symbol1546) = 19
	v631 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1547 = libc.Ptr(&libc.As[TSLexer](v631).F3)
	v632 = *libc.As[unsafe.Pointer](mark_end1547)
	v633 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v632)(v633)
	v634 = *libc.As[byte](result)
	loadedv1548 = (v634 & 1) != 0
	*libc.As[bool](retval) = loadedv1548
	goto _return

sw_bb1549:
	*libc.As[byte](result) = 1
	v635 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1550 = libc.Ptr(&libc.As[TSLexer](v635).F1)
	*libc.As[int16](result_symbol1550) = 20
	v636 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1551 = libc.Ptr(&libc.As[TSLexer](v636).F3)
	v637 = *libc.As[unsafe.Pointer](mark_end1551)
	v638 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v637)(v638)
	v639 = *libc.As[byte](result)
	loadedv1552 = (v639 & 1) != 0
	*libc.As[bool](retval) = loadedv1552
	goto _return

sw_bb1553:
	*libc.As[byte](result) = 1
	v640 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1554 = libc.Ptr(&libc.As[TSLexer](v640).F1)
	*libc.As[int16](result_symbol1554) = 21
	v641 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1555 = libc.Ptr(&libc.As[TSLexer](v641).F3)
	v642 = *libc.As[unsafe.Pointer](mark_end1555)
	v643 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v642)(v643)
	v644 = *libc.As[byte](result)
	loadedv1556 = (v644 & 1) != 0
	*libc.As[bool](retval) = loadedv1556
	goto _return

sw_bb1557:
	*libc.As[byte](result) = 1
	v645 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1558 = libc.Ptr(&libc.As[TSLexer](v645).F1)
	*libc.As[int16](result_symbol1558) = 22
	v646 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1559 = libc.Ptr(&libc.As[TSLexer](v646).F3)
	v647 = *libc.As[unsafe.Pointer](mark_end1559)
	v648 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v647)(v648)
	v649 = *libc.As[int32](lookahead)
	cmp1560 = 48 <= v649
	if cmp1560 {
		goto land_lhs_true1562
	} else {
		goto if_end1566
	}

land_lhs_true1562:
	v650 = *libc.As[int32](lookahead)
	cmp1563 = v650 <= 57
	if cmp1563 {
		goto if_then1565
	} else {
		goto if_end1566
	}

if_then1565:
	*libc.As[int16](state_addr) = 106
	goto next_state

if_end1566:
	v651 = *libc.As[byte](result)
	loadedv1567 = (v651 & 1) != 0
	*libc.As[bool](retval) = loadedv1567
	goto _return

sw_bb1568:
	*libc.As[byte](result) = 1
	v652 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1569 = libc.Ptr(&libc.As[TSLexer](v652).F1)
	*libc.As[int16](result_symbol1569) = 23
	v653 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1570 = libc.Ptr(&libc.As[TSLexer](v653).F3)
	v654 = *libc.As[unsafe.Pointer](mark_end1570)
	v655 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v654)(v655)
	v656 = *libc.As[int32](lookahead)
	cmp1571 = v656 == 32
	if cmp1571 {
		goto if_then1573
	} else {
		goto if_end1574
	}

if_then1573:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end1574:
	v657 = *libc.As[int32](lookahead)
	cmp1575 = v657 == 43
	if cmp1575 {
		goto if_then1583
	} else {
		goto lor_lhs_false1577
	}

lor_lhs_false1577:
	v658 = *libc.As[int32](lookahead)
	cmp1578 = v658 == 45
	if cmp1578 {
		goto if_then1583
	} else {
		goto lor_lhs_false1580
	}

lor_lhs_false1580:
	v659 = *libc.As[int32](lookahead)
	cmp1581 = v659 == 47
	if cmp1581 {
		goto if_then1583
	} else {
		goto if_end1584
	}

if_then1583:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1584:
	v660 = *libc.As[int32](lookahead)
	cmp1585 = 36 <= v660
	if cmp1585 {
		goto land_lhs_true1587
	} else {
		goto lor_lhs_false1590
	}

land_lhs_true1587:
	v661 = *libc.As[int32](lookahead)
	cmp1588 = v661 <= 41
	if cmp1588 {
		goto if_then1593
	} else {
		goto lor_lhs_false1590
	}

lor_lhs_false1590:
	v662 = *libc.As[int32](lookahead)
	cmp1591 = v662 == 64
	if cmp1591 {
		goto if_then1593
	} else {
		goto if_end1594
	}

if_then1593:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1594:
	v663 = *libc.As[int32](lookahead)
	cmp1595 = 46 <= v663
	if cmp1595 {
		goto land_lhs_true1597
	} else {
		goto lor_lhs_false1600
	}

land_lhs_true1597:
	v664 = *libc.As[int32](lookahead)
	cmp1598 = v664 <= 57
	if cmp1598 {
		goto if_then1615
	} else {
		goto lor_lhs_false1600
	}

lor_lhs_false1600:
	v665 = *libc.As[int32](lookahead)
	cmp1601 = 65 <= v665
	if cmp1601 {
		goto land_lhs_true1603
	} else {
		goto lor_lhs_false1606
	}

land_lhs_true1603:
	v666 = *libc.As[int32](lookahead)
	cmp1604 = v666 <= 90
	if cmp1604 {
		goto if_then1615
	} else {
		goto lor_lhs_false1606
	}

lor_lhs_false1606:
	v667 = *libc.As[int32](lookahead)
	cmp1607 = v667 == 95
	if cmp1607 {
		goto if_then1615
	} else {
		goto lor_lhs_false1609
	}

lor_lhs_false1609:
	v668 = *libc.As[int32](lookahead)
	cmp1610 = 97 <= v668
	if cmp1610 {
		goto land_lhs_true1612
	} else {
		goto if_end1616
	}

land_lhs_true1612:
	v669 = *libc.As[int32](lookahead)
	cmp1613 = v669 <= 122
	if cmp1613 {
		goto if_then1615
	} else {
		goto if_end1616
	}

if_then1615:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1616:
	v670 = *libc.As[byte](result)
	loadedv1617 = (v670 & 1) != 0
	*libc.As[bool](retval) = loadedv1617
	goto _return

sw_bb1618:
	*libc.As[byte](result) = 1
	v671 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1619 = libc.Ptr(&libc.As[TSLexer](v671).F1)
	*libc.As[int16](result_symbol1619) = 23
	v672 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1620 = libc.Ptr(&libc.As[TSLexer](v672).F3)
	v673 = *libc.As[unsafe.Pointer](mark_end1620)
	v674 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v673)(v674)
	v675 = *libc.As[int32](lookahead)
	cmp1621 = v675 == 97
	if cmp1621 {
		goto if_then1623
	} else {
		goto if_end1624
	}

if_then1623:
	*libc.As[int16](state_addr) = 115
	goto next_state

if_end1624:
	v676 = *libc.As[int32](lookahead)
	cmp1625 = v676 == 43
	if cmp1625 {
		goto if_then1633
	} else {
		goto lor_lhs_false1627
	}

lor_lhs_false1627:
	v677 = *libc.As[int32](lookahead)
	cmp1628 = v677 == 45
	if cmp1628 {
		goto if_then1633
	} else {
		goto lor_lhs_false1630
	}

lor_lhs_false1630:
	v678 = *libc.As[int32](lookahead)
	cmp1631 = v678 == 47
	if cmp1631 {
		goto if_then1633
	} else {
		goto if_end1634
	}

if_then1633:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1634:
	v679 = *libc.As[int32](lookahead)
	cmp1635 = 36 <= v679
	if cmp1635 {
		goto land_lhs_true1637
	} else {
		goto lor_lhs_false1640
	}

land_lhs_true1637:
	v680 = *libc.As[int32](lookahead)
	cmp1638 = v680 <= 41
	if cmp1638 {
		goto if_then1643
	} else {
		goto lor_lhs_false1640
	}

lor_lhs_false1640:
	v681 = *libc.As[int32](lookahead)
	cmp1641 = v681 == 64
	if cmp1641 {
		goto if_then1643
	} else {
		goto if_end1644
	}

if_then1643:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1644:
	v682 = *libc.As[int32](lookahead)
	cmp1645 = 46 <= v682
	if cmp1645 {
		goto land_lhs_true1647
	} else {
		goto lor_lhs_false1650
	}

land_lhs_true1647:
	v683 = *libc.As[int32](lookahead)
	cmp1648 = v683 <= 57
	if cmp1648 {
		goto if_then1665
	} else {
		goto lor_lhs_false1650
	}

lor_lhs_false1650:
	v684 = *libc.As[int32](lookahead)
	cmp1651 = 65 <= v684
	if cmp1651 {
		goto land_lhs_true1653
	} else {
		goto lor_lhs_false1656
	}

land_lhs_true1653:
	v685 = *libc.As[int32](lookahead)
	cmp1654 = v685 <= 90
	if cmp1654 {
		goto if_then1665
	} else {
		goto lor_lhs_false1656
	}

lor_lhs_false1656:
	v686 = *libc.As[int32](lookahead)
	cmp1657 = v686 == 95
	if cmp1657 {
		goto if_then1665
	} else {
		goto lor_lhs_false1659
	}

lor_lhs_false1659:
	v687 = *libc.As[int32](lookahead)
	cmp1660 = 98 <= v687
	if cmp1660 {
		goto land_lhs_true1662
	} else {
		goto if_end1666
	}

land_lhs_true1662:
	v688 = *libc.As[int32](lookahead)
	cmp1663 = v688 <= 122
	if cmp1663 {
		goto if_then1665
	} else {
		goto if_end1666
	}

if_then1665:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1666:
	v689 = *libc.As[byte](result)
	loadedv1667 = (v689 & 1) != 0
	*libc.As[bool](retval) = loadedv1667
	goto _return

sw_bb1668:
	*libc.As[byte](result) = 1
	v690 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1669 = libc.Ptr(&libc.As[TSLexer](v690).F1)
	*libc.As[int16](result_symbol1669) = 23
	v691 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1670 = libc.Ptr(&libc.As[TSLexer](v691).F3)
	v692 = *libc.As[unsafe.Pointer](mark_end1670)
	v693 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v692)(v693)
	v694 = *libc.As[int32](lookahead)
	cmp1671 = v694 == 98
	if cmp1671 {
		goto if_then1673
	} else {
		goto if_end1674
	}

if_then1673:
	*libc.As[int16](state_addr) = 111
	goto next_state

if_end1674:
	v695 = *libc.As[int32](lookahead)
	cmp1675 = v695 == 43
	if cmp1675 {
		goto if_then1683
	} else {
		goto lor_lhs_false1677
	}

lor_lhs_false1677:
	v696 = *libc.As[int32](lookahead)
	cmp1678 = v696 == 45
	if cmp1678 {
		goto if_then1683
	} else {
		goto lor_lhs_false1680
	}

lor_lhs_false1680:
	v697 = *libc.As[int32](lookahead)
	cmp1681 = v697 == 47
	if cmp1681 {
		goto if_then1683
	} else {
		goto if_end1684
	}

if_then1683:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1684:
	v698 = *libc.As[int32](lookahead)
	cmp1685 = 36 <= v698
	if cmp1685 {
		goto land_lhs_true1687
	} else {
		goto lor_lhs_false1690
	}

land_lhs_true1687:
	v699 = *libc.As[int32](lookahead)
	cmp1688 = v699 <= 41
	if cmp1688 {
		goto if_then1693
	} else {
		goto lor_lhs_false1690
	}

lor_lhs_false1690:
	v700 = *libc.As[int32](lookahead)
	cmp1691 = v700 == 64
	if cmp1691 {
		goto if_then1693
	} else {
		goto if_end1694
	}

if_then1693:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1694:
	v701 = *libc.As[int32](lookahead)
	cmp1695 = 46 <= v701
	if cmp1695 {
		goto land_lhs_true1697
	} else {
		goto lor_lhs_false1700
	}

land_lhs_true1697:
	v702 = *libc.As[int32](lookahead)
	cmp1698 = v702 <= 57
	if cmp1698 {
		goto if_then1715
	} else {
		goto lor_lhs_false1700
	}

lor_lhs_false1700:
	v703 = *libc.As[int32](lookahead)
	cmp1701 = 65 <= v703
	if cmp1701 {
		goto land_lhs_true1703
	} else {
		goto lor_lhs_false1706
	}

land_lhs_true1703:
	v704 = *libc.As[int32](lookahead)
	cmp1704 = v704 <= 90
	if cmp1704 {
		goto if_then1715
	} else {
		goto lor_lhs_false1706
	}

lor_lhs_false1706:
	v705 = *libc.As[int32](lookahead)
	cmp1707 = v705 == 95
	if cmp1707 {
		goto if_then1715
	} else {
		goto lor_lhs_false1709
	}

lor_lhs_false1709:
	v706 = *libc.As[int32](lookahead)
	cmp1710 = 97 <= v706
	if cmp1710 {
		goto land_lhs_true1712
	} else {
		goto if_end1716
	}

land_lhs_true1712:
	v707 = *libc.As[int32](lookahead)
	cmp1713 = v707 <= 122
	if cmp1713 {
		goto if_then1715
	} else {
		goto if_end1716
	}

if_then1715:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1716:
	v708 = *libc.As[byte](result)
	loadedv1717 = (v708 & 1) != 0
	*libc.As[bool](retval) = loadedv1717
	goto _return

sw_bb1718:
	*libc.As[byte](result) = 1
	v709 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1719 = libc.Ptr(&libc.As[TSLexer](v709).F1)
	*libc.As[int16](result_symbol1719) = 23
	v710 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1720 = libc.Ptr(&libc.As[TSLexer](v710).F3)
	v711 = *libc.As[unsafe.Pointer](mark_end1720)
	v712 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v711)(v712)
	v713 = *libc.As[int32](lookahead)
	cmp1721 = v713 == 101
	if cmp1721 {
		goto if_then1723
	} else {
		goto if_end1724
	}

if_then1723:
	*libc.As[int16](state_addr) = 112
	goto next_state

if_end1724:
	v714 = *libc.As[int32](lookahead)
	cmp1725 = v714 == 43
	if cmp1725 {
		goto if_then1733
	} else {
		goto lor_lhs_false1727
	}

lor_lhs_false1727:
	v715 = *libc.As[int32](lookahead)
	cmp1728 = v715 == 45
	if cmp1728 {
		goto if_then1733
	} else {
		goto lor_lhs_false1730
	}

lor_lhs_false1730:
	v716 = *libc.As[int32](lookahead)
	cmp1731 = v716 == 47
	if cmp1731 {
		goto if_then1733
	} else {
		goto if_end1734
	}

if_then1733:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1734:
	v717 = *libc.As[int32](lookahead)
	cmp1735 = 36 <= v717
	if cmp1735 {
		goto land_lhs_true1737
	} else {
		goto lor_lhs_false1740
	}

land_lhs_true1737:
	v718 = *libc.As[int32](lookahead)
	cmp1738 = v718 <= 41
	if cmp1738 {
		goto if_then1743
	} else {
		goto lor_lhs_false1740
	}

lor_lhs_false1740:
	v719 = *libc.As[int32](lookahead)
	cmp1741 = v719 == 64
	if cmp1741 {
		goto if_then1743
	} else {
		goto if_end1744
	}

if_then1743:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1744:
	v720 = *libc.As[int32](lookahead)
	cmp1745 = 46 <= v720
	if cmp1745 {
		goto land_lhs_true1747
	} else {
		goto lor_lhs_false1750
	}

land_lhs_true1747:
	v721 = *libc.As[int32](lookahead)
	cmp1748 = v721 <= 57
	if cmp1748 {
		goto if_then1765
	} else {
		goto lor_lhs_false1750
	}

lor_lhs_false1750:
	v722 = *libc.As[int32](lookahead)
	cmp1751 = 65 <= v722
	if cmp1751 {
		goto land_lhs_true1753
	} else {
		goto lor_lhs_false1756
	}

land_lhs_true1753:
	v723 = *libc.As[int32](lookahead)
	cmp1754 = v723 <= 90
	if cmp1754 {
		goto if_then1765
	} else {
		goto lor_lhs_false1756
	}

lor_lhs_false1756:
	v724 = *libc.As[int32](lookahead)
	cmp1757 = v724 == 95
	if cmp1757 {
		goto if_then1765
	} else {
		goto lor_lhs_false1759
	}

lor_lhs_false1759:
	v725 = *libc.As[int32](lookahead)
	cmp1760 = 97 <= v725
	if cmp1760 {
		goto land_lhs_true1762
	} else {
		goto if_end1766
	}

land_lhs_true1762:
	v726 = *libc.As[int32](lookahead)
	cmp1763 = v726 <= 122
	if cmp1763 {
		goto if_then1765
	} else {
		goto if_end1766
	}

if_then1765:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1766:
	v727 = *libc.As[byte](result)
	loadedv1767 = (v727 & 1) != 0
	*libc.As[bool](retval) = loadedv1767
	goto _return

sw_bb1768:
	*libc.As[byte](result) = 1
	v728 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1769 = libc.Ptr(&libc.As[TSLexer](v728).F1)
	*libc.As[int16](result_symbol1769) = 23
	v729 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1770 = libc.Ptr(&libc.As[TSLexer](v729).F3)
	v730 = *libc.As[unsafe.Pointer](mark_end1770)
	v731 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v730)(v731)
	v732 = *libc.As[int32](lookahead)
	cmp1771 = v732 == 108
	if cmp1771 {
		goto if_then1773
	} else {
		goto if_end1774
	}

if_then1773:
	*libc.As[int16](state_addr) = 116
	goto next_state

if_end1774:
	v733 = *libc.As[int32](lookahead)
	cmp1775 = v733 == 43
	if cmp1775 {
		goto if_then1783
	} else {
		goto lor_lhs_false1777
	}

lor_lhs_false1777:
	v734 = *libc.As[int32](lookahead)
	cmp1778 = v734 == 45
	if cmp1778 {
		goto if_then1783
	} else {
		goto lor_lhs_false1780
	}

lor_lhs_false1780:
	v735 = *libc.As[int32](lookahead)
	cmp1781 = v735 == 47
	if cmp1781 {
		goto if_then1783
	} else {
		goto if_end1784
	}

if_then1783:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1784:
	v736 = *libc.As[int32](lookahead)
	cmp1785 = 36 <= v736
	if cmp1785 {
		goto land_lhs_true1787
	} else {
		goto lor_lhs_false1790
	}

land_lhs_true1787:
	v737 = *libc.As[int32](lookahead)
	cmp1788 = v737 <= 41
	if cmp1788 {
		goto if_then1793
	} else {
		goto lor_lhs_false1790
	}

lor_lhs_false1790:
	v738 = *libc.As[int32](lookahead)
	cmp1791 = v738 == 64
	if cmp1791 {
		goto if_then1793
	} else {
		goto if_end1794
	}

if_then1793:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1794:
	v739 = *libc.As[int32](lookahead)
	cmp1795 = 46 <= v739
	if cmp1795 {
		goto land_lhs_true1797
	} else {
		goto lor_lhs_false1800
	}

land_lhs_true1797:
	v740 = *libc.As[int32](lookahead)
	cmp1798 = v740 <= 57
	if cmp1798 {
		goto if_then1815
	} else {
		goto lor_lhs_false1800
	}

lor_lhs_false1800:
	v741 = *libc.As[int32](lookahead)
	cmp1801 = 65 <= v741
	if cmp1801 {
		goto land_lhs_true1803
	} else {
		goto lor_lhs_false1806
	}

land_lhs_true1803:
	v742 = *libc.As[int32](lookahead)
	cmp1804 = v742 <= 90
	if cmp1804 {
		goto if_then1815
	} else {
		goto lor_lhs_false1806
	}

lor_lhs_false1806:
	v743 = *libc.As[int32](lookahead)
	cmp1807 = v743 == 95
	if cmp1807 {
		goto if_then1815
	} else {
		goto lor_lhs_false1809
	}

lor_lhs_false1809:
	v744 = *libc.As[int32](lookahead)
	cmp1810 = 97 <= v744
	if cmp1810 {
		goto land_lhs_true1812
	} else {
		goto if_end1816
	}

land_lhs_true1812:
	v745 = *libc.As[int32](lookahead)
	cmp1813 = v745 <= 122
	if cmp1813 {
		goto if_then1815
	} else {
		goto if_end1816
	}

if_then1815:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1816:
	v746 = *libc.As[byte](result)
	loadedv1817 = (v746 & 1) != 0
	*libc.As[bool](retval) = loadedv1817
	goto _return

sw_bb1818:
	*libc.As[byte](result) = 1
	v747 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1819 = libc.Ptr(&libc.As[TSLexer](v747).F1)
	*libc.As[int16](result_symbol1819) = 23
	v748 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1820 = libc.Ptr(&libc.As[TSLexer](v748).F3)
	v749 = *libc.As[unsafe.Pointer](mark_end1820)
	v750 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v749)(v750)
	v751 = *libc.As[int32](lookahead)
	cmp1821 = v751 == 109
	if cmp1821 {
		goto if_then1823
	} else {
		goto if_end1824
	}

if_then1823:
	*libc.As[int16](state_addr) = 109
	goto next_state

if_end1824:
	v752 = *libc.As[int32](lookahead)
	cmp1825 = v752 == 43
	if cmp1825 {
		goto if_then1833
	} else {
		goto lor_lhs_false1827
	}

lor_lhs_false1827:
	v753 = *libc.As[int32](lookahead)
	cmp1828 = v753 == 45
	if cmp1828 {
		goto if_then1833
	} else {
		goto lor_lhs_false1830
	}

lor_lhs_false1830:
	v754 = *libc.As[int32](lookahead)
	cmp1831 = v754 == 47
	if cmp1831 {
		goto if_then1833
	} else {
		goto if_end1834
	}

if_then1833:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1834:
	v755 = *libc.As[int32](lookahead)
	cmp1835 = 36 <= v755
	if cmp1835 {
		goto land_lhs_true1837
	} else {
		goto lor_lhs_false1840
	}

land_lhs_true1837:
	v756 = *libc.As[int32](lookahead)
	cmp1838 = v756 <= 41
	if cmp1838 {
		goto if_then1843
	} else {
		goto lor_lhs_false1840
	}

lor_lhs_false1840:
	v757 = *libc.As[int32](lookahead)
	cmp1841 = v757 == 64
	if cmp1841 {
		goto if_then1843
	} else {
		goto if_end1844
	}

if_then1843:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1844:
	v758 = *libc.As[int32](lookahead)
	cmp1845 = 46 <= v758
	if cmp1845 {
		goto land_lhs_true1847
	} else {
		goto lor_lhs_false1850
	}

land_lhs_true1847:
	v759 = *libc.As[int32](lookahead)
	cmp1848 = v759 <= 57
	if cmp1848 {
		goto if_then1865
	} else {
		goto lor_lhs_false1850
	}

lor_lhs_false1850:
	v760 = *libc.As[int32](lookahead)
	cmp1851 = 65 <= v760
	if cmp1851 {
		goto land_lhs_true1853
	} else {
		goto lor_lhs_false1856
	}

land_lhs_true1853:
	v761 = *libc.As[int32](lookahead)
	cmp1854 = v761 <= 90
	if cmp1854 {
		goto if_then1865
	} else {
		goto lor_lhs_false1856
	}

lor_lhs_false1856:
	v762 = *libc.As[int32](lookahead)
	cmp1857 = v762 == 95
	if cmp1857 {
		goto if_then1865
	} else {
		goto lor_lhs_false1859
	}

lor_lhs_false1859:
	v763 = *libc.As[int32](lookahead)
	cmp1860 = 97 <= v763
	if cmp1860 {
		goto land_lhs_true1862
	} else {
		goto if_end1866
	}

land_lhs_true1862:
	v764 = *libc.As[int32](lookahead)
	cmp1863 = v764 <= 122
	if cmp1863 {
		goto if_then1865
	} else {
		goto if_end1866
	}

if_then1865:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1866:
	v765 = *libc.As[byte](result)
	loadedv1867 = (v765 & 1) != 0
	*libc.As[bool](retval) = loadedv1867
	goto _return

sw_bb1868:
	*libc.As[byte](result) = 1
	v766 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1869 = libc.Ptr(&libc.As[TSLexer](v766).F1)
	*libc.As[int16](result_symbol1869) = 23
	v767 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1870 = libc.Ptr(&libc.As[TSLexer](v767).F3)
	v768 = *libc.As[unsafe.Pointer](mark_end1870)
	v769 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v768)(v769)
	v770 = *libc.As[int32](lookahead)
	cmp1871 = v770 == 115
	if cmp1871 {
		goto if_then1873
	} else {
		goto if_end1874
	}

if_then1873:
	*libc.As[int16](state_addr) = 108
	goto next_state

if_end1874:
	v771 = *libc.As[int32](lookahead)
	cmp1875 = v771 == 43
	if cmp1875 {
		goto if_then1883
	} else {
		goto lor_lhs_false1877
	}

lor_lhs_false1877:
	v772 = *libc.As[int32](lookahead)
	cmp1878 = v772 == 45
	if cmp1878 {
		goto if_then1883
	} else {
		goto lor_lhs_false1880
	}

lor_lhs_false1880:
	v773 = *libc.As[int32](lookahead)
	cmp1881 = v773 == 47
	if cmp1881 {
		goto if_then1883
	} else {
		goto if_end1884
	}

if_then1883:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1884:
	v774 = *libc.As[int32](lookahead)
	cmp1885 = 36 <= v774
	if cmp1885 {
		goto land_lhs_true1887
	} else {
		goto lor_lhs_false1890
	}

land_lhs_true1887:
	v775 = *libc.As[int32](lookahead)
	cmp1888 = v775 <= 41
	if cmp1888 {
		goto if_then1893
	} else {
		goto lor_lhs_false1890
	}

lor_lhs_false1890:
	v776 = *libc.As[int32](lookahead)
	cmp1891 = v776 == 64
	if cmp1891 {
		goto if_then1893
	} else {
		goto if_end1894
	}

if_then1893:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1894:
	v777 = *libc.As[int32](lookahead)
	cmp1895 = 46 <= v777
	if cmp1895 {
		goto land_lhs_true1897
	} else {
		goto lor_lhs_false1900
	}

land_lhs_true1897:
	v778 = *libc.As[int32](lookahead)
	cmp1898 = v778 <= 57
	if cmp1898 {
		goto if_then1915
	} else {
		goto lor_lhs_false1900
	}

lor_lhs_false1900:
	v779 = *libc.As[int32](lookahead)
	cmp1901 = 65 <= v779
	if cmp1901 {
		goto land_lhs_true1903
	} else {
		goto lor_lhs_false1906
	}

land_lhs_true1903:
	v780 = *libc.As[int32](lookahead)
	cmp1904 = v780 <= 90
	if cmp1904 {
		goto if_then1915
	} else {
		goto lor_lhs_false1906
	}

lor_lhs_false1906:
	v781 = *libc.As[int32](lookahead)
	cmp1907 = v781 == 95
	if cmp1907 {
		goto if_then1915
	} else {
		goto lor_lhs_false1909
	}

lor_lhs_false1909:
	v782 = *libc.As[int32](lookahead)
	cmp1910 = 97 <= v782
	if cmp1910 {
		goto land_lhs_true1912
	} else {
		goto if_end1916
	}

land_lhs_true1912:
	v783 = *libc.As[int32](lookahead)
	cmp1913 = v783 <= 122
	if cmp1913 {
		goto if_then1915
	} else {
		goto if_end1916
	}

if_then1915:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1916:
	v784 = *libc.As[byte](result)
	loadedv1917 = (v784 & 1) != 0
	*libc.As[bool](retval) = loadedv1917
	goto _return

sw_bb1918:
	*libc.As[byte](result) = 1
	v785 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1919 = libc.Ptr(&libc.As[TSLexer](v785).F1)
	*libc.As[int16](result_symbol1919) = 23
	v786 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1920 = libc.Ptr(&libc.As[TSLexer](v786).F3)
	v787 = *libc.As[unsafe.Pointer](mark_end1920)
	v788 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v787)(v788)
	v789 = *libc.As[int32](lookahead)
	cmp1921 = v789 == 115
	if cmp1921 {
		goto if_then1923
	} else {
		goto if_end1924
	}

if_then1923:
	*libc.As[int16](state_addr) = 110
	goto next_state

if_end1924:
	v790 = *libc.As[int32](lookahead)
	cmp1925 = v790 == 43
	if cmp1925 {
		goto if_then1933
	} else {
		goto lor_lhs_false1927
	}

lor_lhs_false1927:
	v791 = *libc.As[int32](lookahead)
	cmp1928 = v791 == 45
	if cmp1928 {
		goto if_then1933
	} else {
		goto lor_lhs_false1930
	}

lor_lhs_false1930:
	v792 = *libc.As[int32](lookahead)
	cmp1931 = v792 == 47
	if cmp1931 {
		goto if_then1933
	} else {
		goto if_end1934
	}

if_then1933:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1934:
	v793 = *libc.As[int32](lookahead)
	cmp1935 = 36 <= v793
	if cmp1935 {
		goto land_lhs_true1937
	} else {
		goto lor_lhs_false1940
	}

land_lhs_true1937:
	v794 = *libc.As[int32](lookahead)
	cmp1938 = v794 <= 41
	if cmp1938 {
		goto if_then1943
	} else {
		goto lor_lhs_false1940
	}

lor_lhs_false1940:
	v795 = *libc.As[int32](lookahead)
	cmp1941 = v795 == 64
	if cmp1941 {
		goto if_then1943
	} else {
		goto if_end1944
	}

if_then1943:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1944:
	v796 = *libc.As[int32](lookahead)
	cmp1945 = 46 <= v796
	if cmp1945 {
		goto land_lhs_true1947
	} else {
		goto lor_lhs_false1950
	}

land_lhs_true1947:
	v797 = *libc.As[int32](lookahead)
	cmp1948 = v797 <= 57
	if cmp1948 {
		goto if_then1965
	} else {
		goto lor_lhs_false1950
	}

lor_lhs_false1950:
	v798 = *libc.As[int32](lookahead)
	cmp1951 = 65 <= v798
	if cmp1951 {
		goto land_lhs_true1953
	} else {
		goto lor_lhs_false1956
	}

land_lhs_true1953:
	v799 = *libc.As[int32](lookahead)
	cmp1954 = v799 <= 90
	if cmp1954 {
		goto if_then1965
	} else {
		goto lor_lhs_false1956
	}

lor_lhs_false1956:
	v800 = *libc.As[int32](lookahead)
	cmp1957 = v800 == 95
	if cmp1957 {
		goto if_then1965
	} else {
		goto lor_lhs_false1959
	}

lor_lhs_false1959:
	v801 = *libc.As[int32](lookahead)
	cmp1960 = 97 <= v801
	if cmp1960 {
		goto land_lhs_true1962
	} else {
		goto if_end1966
	}

land_lhs_true1962:
	v802 = *libc.As[int32](lookahead)
	cmp1963 = v802 <= 122
	if cmp1963 {
		goto if_then1965
	} else {
		goto if_end1966
	}

if_then1965:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end1966:
	v803 = *libc.As[byte](result)
	loadedv1967 = (v803 & 1) != 0
	*libc.As[bool](retval) = loadedv1967
	goto _return

sw_bb1968:
	*libc.As[byte](result) = 1
	v804 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1969 = libc.Ptr(&libc.As[TSLexer](v804).F1)
	*libc.As[int16](result_symbol1969) = 23
	v805 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1970 = libc.Ptr(&libc.As[TSLexer](v805).F3)
	v806 = *libc.As[unsafe.Pointer](mark_end1970)
	v807 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v806)(v807)
	v808 = *libc.As[int32](lookahead)
	cmp1971 = v808 == 115
	if cmp1971 {
		goto if_then1973
	} else {
		goto if_end1974
	}

if_then1973:
	*libc.As[int16](state_addr) = 114
	goto next_state

if_end1974:
	v809 = *libc.As[int32](lookahead)
	cmp1975 = v809 == 43
	if cmp1975 {
		goto if_then1983
	} else {
		goto lor_lhs_false1977
	}

lor_lhs_false1977:
	v810 = *libc.As[int32](lookahead)
	cmp1978 = v810 == 45
	if cmp1978 {
		goto if_then1983
	} else {
		goto lor_lhs_false1980
	}

lor_lhs_false1980:
	v811 = *libc.As[int32](lookahead)
	cmp1981 = v811 == 47
	if cmp1981 {
		goto if_then1983
	} else {
		goto if_end1984
	}

if_then1983:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end1984:
	v812 = *libc.As[int32](lookahead)
	cmp1985 = 36 <= v812
	if cmp1985 {
		goto land_lhs_true1987
	} else {
		goto lor_lhs_false1990
	}

land_lhs_true1987:
	v813 = *libc.As[int32](lookahead)
	cmp1988 = v813 <= 41
	if cmp1988 {
		goto if_then1993
	} else {
		goto lor_lhs_false1990
	}

lor_lhs_false1990:
	v814 = *libc.As[int32](lookahead)
	cmp1991 = v814 == 64
	if cmp1991 {
		goto if_then1993
	} else {
		goto if_end1994
	}

if_then1993:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end1994:
	v815 = *libc.As[int32](lookahead)
	cmp1995 = 46 <= v815
	if cmp1995 {
		goto land_lhs_true1997
	} else {
		goto lor_lhs_false2000
	}

land_lhs_true1997:
	v816 = *libc.As[int32](lookahead)
	cmp1998 = v816 <= 57
	if cmp1998 {
		goto if_then2015
	} else {
		goto lor_lhs_false2000
	}

lor_lhs_false2000:
	v817 = *libc.As[int32](lookahead)
	cmp2001 = 65 <= v817
	if cmp2001 {
		goto land_lhs_true2003
	} else {
		goto lor_lhs_false2006
	}

land_lhs_true2003:
	v818 = *libc.As[int32](lookahead)
	cmp2004 = v818 <= 90
	if cmp2004 {
		goto if_then2015
	} else {
		goto lor_lhs_false2006
	}

lor_lhs_false2006:
	v819 = *libc.As[int32](lookahead)
	cmp2007 = v819 == 95
	if cmp2007 {
		goto if_then2015
	} else {
		goto lor_lhs_false2009
	}

lor_lhs_false2009:
	v820 = *libc.As[int32](lookahead)
	cmp2010 = 97 <= v820
	if cmp2010 {
		goto land_lhs_true2012
	} else {
		goto if_end2016
	}

land_lhs_true2012:
	v821 = *libc.As[int32](lookahead)
	cmp2013 = v821 <= 122
	if cmp2013 {
		goto if_then2015
	} else {
		goto if_end2016
	}

if_then2015:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end2016:
	v822 = *libc.As[byte](result)
	loadedv2017 = (v822 & 1) != 0
	*libc.As[bool](retval) = loadedv2017
	goto _return

sw_bb2018:
	*libc.As[byte](result) = 1
	v823 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2019 = libc.Ptr(&libc.As[TSLexer](v823).F1)
	*libc.As[int16](result_symbol2019) = 23
	v824 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2020 = libc.Ptr(&libc.As[TSLexer](v824).F3)
	v825 = *libc.As[unsafe.Pointer](mark_end2020)
	v826 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v825)(v826)
	v827 = *libc.As[int32](lookahead)
	cmp2021 = v827 == 121
	if cmp2021 {
		goto if_then2023
	} else {
		goto if_end2024
	}

if_then2023:
	*libc.As[int16](state_addr) = 107
	goto next_state

if_end2024:
	v828 = *libc.As[int32](lookahead)
	cmp2025 = v828 == 43
	if cmp2025 {
		goto if_then2033
	} else {
		goto lor_lhs_false2027
	}

lor_lhs_false2027:
	v829 = *libc.As[int32](lookahead)
	cmp2028 = v829 == 45
	if cmp2028 {
		goto if_then2033
	} else {
		goto lor_lhs_false2030
	}

lor_lhs_false2030:
	v830 = *libc.As[int32](lookahead)
	cmp2031 = v830 == 47
	if cmp2031 {
		goto if_then2033
	} else {
		goto if_end2034
	}

if_then2033:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end2034:
	v831 = *libc.As[int32](lookahead)
	cmp2035 = 36 <= v831
	if cmp2035 {
		goto land_lhs_true2037
	} else {
		goto lor_lhs_false2040
	}

land_lhs_true2037:
	v832 = *libc.As[int32](lookahead)
	cmp2038 = v832 <= 41
	if cmp2038 {
		goto if_then2043
	} else {
		goto lor_lhs_false2040
	}

lor_lhs_false2040:
	v833 = *libc.As[int32](lookahead)
	cmp2041 = v833 == 64
	if cmp2041 {
		goto if_then2043
	} else {
		goto if_end2044
	}

if_then2043:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end2044:
	v834 = *libc.As[int32](lookahead)
	cmp2045 = 46 <= v834
	if cmp2045 {
		goto land_lhs_true2047
	} else {
		goto lor_lhs_false2050
	}

land_lhs_true2047:
	v835 = *libc.As[int32](lookahead)
	cmp2048 = v835 <= 57
	if cmp2048 {
		goto if_then2065
	} else {
		goto lor_lhs_false2050
	}

lor_lhs_false2050:
	v836 = *libc.As[int32](lookahead)
	cmp2051 = 65 <= v836
	if cmp2051 {
		goto land_lhs_true2053
	} else {
		goto lor_lhs_false2056
	}

land_lhs_true2053:
	v837 = *libc.As[int32](lookahead)
	cmp2054 = v837 <= 90
	if cmp2054 {
		goto if_then2065
	} else {
		goto lor_lhs_false2056
	}

lor_lhs_false2056:
	v838 = *libc.As[int32](lookahead)
	cmp2057 = v838 == 95
	if cmp2057 {
		goto if_then2065
	} else {
		goto lor_lhs_false2059
	}

lor_lhs_false2059:
	v839 = *libc.As[int32](lookahead)
	cmp2060 = 97 <= v839
	if cmp2060 {
		goto land_lhs_true2062
	} else {
		goto if_end2066
	}

land_lhs_true2062:
	v840 = *libc.As[int32](lookahead)
	cmp2063 = v840 <= 122
	if cmp2063 {
		goto if_then2065
	} else {
		goto if_end2066
	}

if_then2065:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end2066:
	v841 = *libc.As[byte](result)
	loadedv2067 = (v841 & 1) != 0
	*libc.As[bool](retval) = loadedv2067
	goto _return

sw_bb2068:
	*libc.As[byte](result) = 1
	v842 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2069 = libc.Ptr(&libc.As[TSLexer](v842).F1)
	*libc.As[int16](result_symbol2069) = 23
	v843 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2070 = libc.Ptr(&libc.As[TSLexer](v843).F3)
	v844 = *libc.As[unsafe.Pointer](mark_end2070)
	v845 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v844)(v845)
	v846 = *libc.As[int32](lookahead)
	cmp2071 = v846 == 43
	if cmp2071 {
		goto if_then2079
	} else {
		goto lor_lhs_false2073
	}

lor_lhs_false2073:
	v847 = *libc.As[int32](lookahead)
	cmp2074 = v847 == 45
	if cmp2074 {
		goto if_then2079
	} else {
		goto lor_lhs_false2076
	}

lor_lhs_false2076:
	v848 = *libc.As[int32](lookahead)
	cmp2077 = v848 == 47
	if cmp2077 {
		goto if_then2079
	} else {
		goto if_end2080
	}

if_then2079:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end2080:
	v849 = *libc.As[int32](lookahead)
	cmp2081 = 36 <= v849
	if cmp2081 {
		goto land_lhs_true2083
	} else {
		goto lor_lhs_false2086
	}

land_lhs_true2083:
	v850 = *libc.As[int32](lookahead)
	cmp2084 = v850 <= 41
	if cmp2084 {
		goto if_then2089
	} else {
		goto lor_lhs_false2086
	}

lor_lhs_false2086:
	v851 = *libc.As[int32](lookahead)
	cmp2087 = v851 == 64
	if cmp2087 {
		goto if_then2089
	} else {
		goto if_end2090
	}

if_then2089:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end2090:
	v852 = *libc.As[int32](lookahead)
	cmp2091 = 46 <= v852
	if cmp2091 {
		goto land_lhs_true2093
	} else {
		goto lor_lhs_false2096
	}

land_lhs_true2093:
	v853 = *libc.As[int32](lookahead)
	cmp2094 = v853 <= 57
	if cmp2094 {
		goto if_then2111
	} else {
		goto lor_lhs_false2096
	}

lor_lhs_false2096:
	v854 = *libc.As[int32](lookahead)
	cmp2097 = 65 <= v854
	if cmp2097 {
		goto land_lhs_true2099
	} else {
		goto lor_lhs_false2102
	}

land_lhs_true2099:
	v855 = *libc.As[int32](lookahead)
	cmp2100 = v855 <= 90
	if cmp2100 {
		goto if_then2111
	} else {
		goto lor_lhs_false2102
	}

lor_lhs_false2102:
	v856 = *libc.As[int32](lookahead)
	cmp2103 = v856 == 95
	if cmp2103 {
		goto if_then2111
	} else {
		goto lor_lhs_false2105
	}

lor_lhs_false2105:
	v857 = *libc.As[int32](lookahead)
	cmp2106 = 97 <= v857
	if cmp2106 {
		goto land_lhs_true2108
	} else {
		goto if_end2112
	}

land_lhs_true2108:
	v858 = *libc.As[int32](lookahead)
	cmp2109 = v858 <= 122
	if cmp2109 {
		goto if_then2111
	} else {
		goto if_end2112
	}

if_then2111:
	*libc.As[int16](state_addr) = 117
	goto next_state

if_end2112:
	v859 = *libc.As[byte](result)
	loadedv2113 = (v859 & 1) != 0
	*libc.As[bool](retval) = loadedv2113
	goto _return

sw_bb2114:
	*libc.As[byte](result) = 1
	v860 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2115 = libc.Ptr(&libc.As[TSLexer](v860).F1)
	*libc.As[int16](result_symbol2115) = 23
	v861 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2116 = libc.Ptr(&libc.As[TSLexer](v861).F3)
	v862 = *libc.As[unsafe.Pointer](mark_end2116)
	v863 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v862)(v863)
	v864 = *libc.As[int32](lookahead)
	cmp2117 = v864 == 43
	if cmp2117 {
		goto if_then2140
	} else {
		goto lor_lhs_false2119
	}

lor_lhs_false2119:
	v865 = *libc.As[int32](lookahead)
	cmp2120 = 45 <= v865
	if cmp2120 {
		goto land_lhs_true2122
	} else {
		goto lor_lhs_false2125
	}

land_lhs_true2122:
	v866 = *libc.As[int32](lookahead)
	cmp2123 = v866 <= 57
	if cmp2123 {
		goto if_then2140
	} else {
		goto lor_lhs_false2125
	}

lor_lhs_false2125:
	v867 = *libc.As[int32](lookahead)
	cmp2126 = 65 <= v867
	if cmp2126 {
		goto land_lhs_true2128
	} else {
		goto lor_lhs_false2131
	}

land_lhs_true2128:
	v868 = *libc.As[int32](lookahead)
	cmp2129 = v868 <= 90
	if cmp2129 {
		goto if_then2140
	} else {
		goto lor_lhs_false2131
	}

lor_lhs_false2131:
	v869 = *libc.As[int32](lookahead)
	cmp2132 = v869 == 95
	if cmp2132 {
		goto if_then2140
	} else {
		goto lor_lhs_false2134
	}

lor_lhs_false2134:
	v870 = *libc.As[int32](lookahead)
	cmp2135 = 97 <= v870
	if cmp2135 {
		goto land_lhs_true2137
	} else {
		goto if_end2141
	}

land_lhs_true2137:
	v871 = *libc.As[int32](lookahead)
	cmp2138 = v871 <= 122
	if cmp2138 {
		goto if_then2140
	} else {
		goto if_end2141
	}

if_then2140:
	*libc.As[int16](state_addr) = 118
	goto next_state

if_end2141:
	v872 = *libc.As[byte](result)
	loadedv2142 = (v872 & 1) != 0
	*libc.As[bool](retval) = loadedv2142
	goto _return

sw_bb2143:
	*libc.As[byte](result) = 1
	v873 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2144 = libc.Ptr(&libc.As[TSLexer](v873).F1)
	*libc.As[int16](result_symbol2144) = 24
	v874 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2145 = libc.Ptr(&libc.As[TSLexer](v874).F3)
	v875 = *libc.As[unsafe.Pointer](mark_end2145)
	v876 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v875)(v876)
	v877 = *libc.As[int32](lookahead)
	cmp2146 = 36 <= v877
	if cmp2146 {
		goto land_lhs_true2148
	} else {
		goto lor_lhs_false2151
	}

land_lhs_true2148:
	v878 = *libc.As[int32](lookahead)
	cmp2149 = v878 <= 41
	if cmp2149 {
		goto if_then2175
	} else {
		goto lor_lhs_false2151
	}

lor_lhs_false2151:
	v879 = *libc.As[int32](lookahead)
	cmp2152 = v879 == 46
	if cmp2152 {
		goto if_then2175
	} else {
		goto lor_lhs_false2154
	}

lor_lhs_false2154:
	v880 = *libc.As[int32](lookahead)
	cmp2155 = 48 <= v880
	if cmp2155 {
		goto land_lhs_true2157
	} else {
		goto lor_lhs_false2160
	}

land_lhs_true2157:
	v881 = *libc.As[int32](lookahead)
	cmp2158 = v881 <= 57
	if cmp2158 {
		goto if_then2175
	} else {
		goto lor_lhs_false2160
	}

lor_lhs_false2160:
	v882 = *libc.As[int32](lookahead)
	cmp2161 = 64 <= v882
	if cmp2161 {
		goto land_lhs_true2163
	} else {
		goto lor_lhs_false2166
	}

land_lhs_true2163:
	v883 = *libc.As[int32](lookahead)
	cmp2164 = v883 <= 90
	if cmp2164 {
		goto if_then2175
	} else {
		goto lor_lhs_false2166
	}

lor_lhs_false2166:
	v884 = *libc.As[int32](lookahead)
	cmp2167 = v884 == 95
	if cmp2167 {
		goto if_then2175
	} else {
		goto lor_lhs_false2169
	}

lor_lhs_false2169:
	v885 = *libc.As[int32](lookahead)
	cmp2170 = 97 <= v885
	if cmp2170 {
		goto land_lhs_true2172
	} else {
		goto if_end2176
	}

land_lhs_true2172:
	v886 = *libc.As[int32](lookahead)
	cmp2173 = v886 <= 122
	if cmp2173 {
		goto if_then2175
	} else {
		goto if_end2176
	}

if_then2175:
	*libc.As[int16](state_addr) = 119
	goto next_state

if_end2176:
	v887 = *libc.As[byte](result)
	loadedv2177 = (v887 & 1) != 0
	*libc.As[bool](retval) = loadedv2177
	goto _return

sw_bb2178:
	*libc.As[byte](result) = 1
	v888 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2179 = libc.Ptr(&libc.As[TSLexer](v888).F1)
	*libc.As[int16](result_symbol2179) = 25
	v889 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2180 = libc.Ptr(&libc.As[TSLexer](v889).F3)
	v890 = *libc.As[unsafe.Pointer](mark_end2180)
	v891 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v890)(v891)
	v892 = *libc.As[int32](lookahead)
	cmp2181 = v892 == 9
	if cmp2181 {
		goto if_then2189
	} else {
		goto lor_lhs_false2183
	}

lor_lhs_false2183:
	v893 = *libc.As[int32](lookahead)
	cmp2184 = v893 == 13
	if cmp2184 {
		goto if_then2189
	} else {
		goto lor_lhs_false2186
	}

lor_lhs_false2186:
	v894 = *libc.As[int32](lookahead)
	cmp2187 = v894 == 32
	if cmp2187 {
		goto if_then2189
	} else {
		goto if_end2190
	}

if_then2189:
	*libc.As[int16](state_addr) = 120
	goto next_state

if_end2190:
	v895 = *libc.As[int32](lookahead)
	cmp2191 = v895 != 0
	if cmp2191 {
		goto land_lhs_true2193
	} else {
		goto if_end2197
	}

land_lhs_true2193:
	v896 = *libc.As[int32](lookahead)
	cmp2194 = v896 != 10
	if cmp2194 {
		goto if_then2196
	} else {
		goto if_end2197
	}

if_then2196:
	*libc.As[int16](state_addr) = 121
	goto next_state

if_end2197:
	v897 = *libc.As[byte](result)
	loadedv2198 = (v897 & 1) != 0
	*libc.As[bool](retval) = loadedv2198
	goto _return

sw_bb2199:
	*libc.As[byte](result) = 1
	v898 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2200 = libc.Ptr(&libc.As[TSLexer](v898).F1)
	*libc.As[int16](result_symbol2200) = 25
	v899 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2201 = libc.Ptr(&libc.As[TSLexer](v899).F3)
	v900 = *libc.As[unsafe.Pointer](mark_end2201)
	v901 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v900)(v901)
	v902 = *libc.As[int32](lookahead)
	cmp2202 = v902 != 0
	if cmp2202 {
		goto land_lhs_true2204
	} else {
		goto if_end2208
	}

land_lhs_true2204:
	v903 = *libc.As[int32](lookahead)
	cmp2205 = v903 != 10
	if cmp2205 {
		goto if_then2207
	} else {
		goto if_end2208
	}

if_then2207:
	*libc.As[int16](state_addr) = 121
	goto next_state

if_end2208:
	v904 = *libc.As[byte](result)
	loadedv2209 = (v904 & 1) != 0
	*libc.As[bool](retval) = loadedv2209
	goto _return

sw_bb2210:
	*libc.As[byte](result) = 1
	v905 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2211 = libc.Ptr(&libc.As[TSLexer](v905).F1)
	*libc.As[int16](result_symbol2211) = 26
	v906 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2212 = libc.Ptr(&libc.As[TSLexer](v906).F3)
	v907 = *libc.As[unsafe.Pointer](mark_end2212)
	v908 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v907)(v908)
	v909 = *libc.As[int32](lookahead)
	cmp2213 = v909 == 9
	if cmp2213 {
		goto if_then2224
	} else {
		goto lor_lhs_false2215
	}

lor_lhs_false2215:
	v910 = *libc.As[int32](lookahead)
	cmp2216 = v910 == 10
	if cmp2216 {
		goto if_then2224
	} else {
		goto lor_lhs_false2218
	}

lor_lhs_false2218:
	v911 = *libc.As[int32](lookahead)
	cmp2219 = v911 == 13
	if cmp2219 {
		goto if_then2224
	} else {
		goto lor_lhs_false2221
	}

lor_lhs_false2221:
	v912 = *libc.As[int32](lookahead)
	cmp2222 = v912 == 32
	if cmp2222 {
		goto if_then2224
	} else {
		goto if_end2225
	}

if_then2224:
	*libc.As[int16](state_addr) = 122
	goto next_state

if_end2225:
	v913 = *libc.As[int32](lookahead)
	cmp2226 = v913 != 0
	if cmp2226 {
		goto land_lhs_true2228
	} else {
		goto if_end2232
	}

land_lhs_true2228:
	v914 = *libc.As[int32](lookahead)
	cmp2229 = v914 != 58
	if cmp2229 {
		goto if_then2231
	} else {
		goto if_end2232
	}

if_then2231:
	*libc.As[int16](state_addr) = 123
	goto next_state

if_end2232:
	v915 = *libc.As[byte](result)
	loadedv2233 = (v915 & 1) != 0
	*libc.As[bool](retval) = loadedv2233
	goto _return

sw_bb2234:
	*libc.As[byte](result) = 1
	v916 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol2235 = libc.Ptr(&libc.As[TSLexer](v916).F1)
	*libc.As[int16](result_symbol2235) = 26
	v917 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end2236 = libc.Ptr(&libc.As[TSLexer](v917).F3)
	v918 = *libc.As[unsafe.Pointer](mark_end2236)
	v919 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v918)(v919)
	v920 = *libc.As[int32](lookahead)
	cmp2237 = v920 != 0
	if cmp2237 {
		goto land_lhs_true2239
	} else {
		goto if_end2243
	}

land_lhs_true2239:
	v921 = *libc.As[int32](lookahead)
	cmp2240 = v921 != 58
	if cmp2240 {
		goto if_then2242
	} else {
		goto if_end2243
	}

if_then2242:
	*libc.As[int16](state_addr) = 123
	goto next_state

if_end2243:
	v922 = *libc.As[byte](result)
	loadedv2244 = (v922 & 1) != 0
	*libc.As[bool](retval) = loadedv2244
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v923 = *libc.As[bool](retval)
	return v923
}
func is_hexadecimal_character(character byte) bool {
	var v1 bool
	var retval unsafe.Pointer
	var conv int32
	var v0 byte
	var character_addr unsafe.Pointer
	_, _, _, _, _ = retval, character_addr, v0, conv, v1

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	character_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	*libc.As[byte](character_addr) = character
	v0 = *libc.As[byte](character_addr)
	conv = int32(int8(v0))
	switch conv {
	case 48:
		goto sw_bb
	case 49:
		goto sw_bb
	case 50:
		goto sw_bb
	case 51:
		goto sw_bb
	case 52:
		goto sw_bb
	case 53:
		goto sw_bb
	case 54:
		goto sw_bb
	case 55:
		goto sw_bb
	case 56:
		goto sw_bb
	case 57:
		goto sw_bb
	case 65:
		goto sw_bb
	case 66:
		goto sw_bb
	case 67:
		goto sw_bb
	case 68:
		goto sw_bb
	case 69:
		goto sw_bb
	case 70:
		goto sw_bb
	case 97:
		goto sw_bb
	case 98:
		goto sw_bb
	case 99:
		goto sw_bb
	case 100:
		goto sw_bb
	case 101:
		goto sw_bb
	case 102:
		goto sw_bb
	case 104:
		goto sw_bb
	case 120:
		goto sw_bb
	default:
		goto sw_default
	}

sw_bb:
	*libc.As[bool](retval) = true
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v1 = *libc.As[bool](retval)
	return v1
}
