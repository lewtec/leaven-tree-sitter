package grammar_json

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type TSFieldMapSlice struct {
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
type TSLexMode struct {
	F0 int16
	F1 int16
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

var tree_sitter_json_language struct {
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
var ts_small_parse_table [360]int16 = [360]int16{2, 3, 1, 14, 37, 11, 0, 1, 2, 3, 5, 6, 7, 10, 11, 12, 13, 2, 3, 1, 14, 39, 11, 0, 1, 2, 3, 5, 6, 7, 10, 11, 12, 13, 2, 3, 1, 14, 41, 11, 0, 1, 2, 3, 5, 6, 7, 10, 11, 12, 13, 7, 3, 1, 14, 7, 1, 1, 9, 1, 5, 11, 1, 7, 29, 1, 16, 8, 3, 17, 19, 20, 13, 4, 10, 11, 12, 13, 2, 3, 1, 14, 43, 11, 0, 1, 2, 3, 5, 6, 7, 10, 11, 12, 13, 2, 3, 1, 14, 45, 11, 0, 1, 2, 3, 5, 6, 7, 10, 11, 12, 13, 7, 3, 1, 14, 7, 1, 1, 9, 1, 5, 11, 1, 7, 28, 1, 16, 8, 3, 17, 19, 20, 13, 4, 10, 11, 12, 13, 2, 3, 1, 14, 47, 11, 0, 1, 2, 3, 5, 6, 7, 10, 11, 12, 13, 2, 3, 1, 14, 49, 11, 0, 1, 2, 3, 5, 6, 7, 10, 11, 12, 13, 5, 3, 1, 14, 11, 1, 7, 51, 1, 3, 20, 1, 18, 31, 1, 20, 4, 53, 1, 7, 57, 1, 14, 18, 1, 21, 55, 2, 8, 9, 4, 57, 1, 14, 59, 1, 7, 19, 1, 21, 61, 2, 8, 9, 4, 57, 1, 14, 63, 1, 7, 19, 1, 21, 65, 2, 8, 9, 4, 3, 1, 14, 68, 1, 2, 70, 1, 3, 22, 1, 23, 4, 3, 1, 14, 72, 1, 2, 74, 1, 6, 24, 1, 24, 4, 3, 1, 14, 68, 1, 2, 76, 1, 3, 25, 1, 23, 4, 3, 1, 14, 11, 1, 7, 27, 1, 18, 31, 1, 20, 4, 3, 1, 14, 72, 1, 2, 78, 1, 6, 26, 1, 24, 4, 3, 1, 14, 80, 1, 2, 83, 1, 3, 25, 1, 23, 4, 3, 1, 14, 85, 1, 2, 88, 1, 6, 26, 1, 24, 2, 3, 1, 14, 83, 2, 2, 3, 2, 3, 1, 14, 90, 2, 2, 3, 2, 3, 1, 14, 88, 2, 2, 6, 2, 3, 1, 14, 92, 1, 0, 2, 3, 1, 14, 94, 1, 4}
var ts_small_parse_table_map [25]int32 = [25]int32{0, 17, 34, 51, 78, 95, 112, 139, 156, 173, 189, 203, 217, 231, 244, 257, 270, 283, 296, 309, 322, 330, 338, 346, 353}
var ts_symbol_names [25]unsafe.Pointer = [25]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7), libc.Ptr(&_str_8), libc.Ptr(&_str_9), libc.Ptr(&_str_10), libc.Ptr(&_str_11), libc.Ptr(&_str_12), libc.Ptr(&_str_13), libc.Ptr(&_str_14), libc.Ptr(&_str_15), libc.Ptr(&_str_16), libc.Ptr(&_str_17), libc.Ptr(&_str_18), libc.Ptr(&_str_19), libc.Ptr(&_str_20), libc.Ptr(&_str_21), libc.Ptr(&_str_22), libc.Ptr(&_str_23), libc.Ptr(&_str_24), libc.Ptr(&_str_25), libc.Ptr(&_str_26)}
var ts_field_names [3]unsafe.Pointer = [3]unsafe.Pointer{nil, libc.Ptr(&_str_27), libc.Ptr(&_str_28)}
var ts_field_map_slices [2]TSFieldMapSlice = [2]TSFieldMapSlice{TSFieldMapSlice{}, TSFieldMapSlice{0, 2}}
var ts_field_map_entries [2]TSFieldMapEntry = [2]TSFieldMapEntry{TSFieldMapEntry{1, 0, 0}, TSFieldMapEntry{2, 2, 0}}
var ts_symbol_metadata [25]TSSymbolMetadata = [25]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 1}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}}
var ts_symbol_map [25]int16 = [25]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [2][4]int16 = [2][4]int16{}
var ts_primary_state_ids [32]int16 = [32]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
var ts_parse_table struct {
	F0 struct {
		F0 [15]int16
		F1 [10]int16
	}
	F1 [25]int16
	F2 [25]int16
	F3 [25]int16
	F4 [25]int16
	F5 struct {
		F0 [15]int16
		F1 [10]int16
	}
	F6 struct {
		F0 [15]int16
		F1 [10]int16
	}
} = struct {
	F0 struct {
		F0 [15]int16
		F1 [10]int16
	}
	F1 [25]int16
	F2 [25]int16
	F3 [25]int16
	F4 [25]int16
	F5 struct {
		F0 [15]int16
		F1 [10]int16
	}
	F6 struct {
		F0 [15]int16
		F1 [10]int16
	}
}{struct {
	F0 [15]int16
	F1 [10]int16
}{[15]int16{1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 3}, [10]int16{}}, [25]int16{5, 7, 0, 0, 0, 9, 0, 11, 0, 0, 13, 13, 13, 13, 3, 30, 2, 8, 0, 8, 8, 0, 2, 0, 0}, [25]int16{15, 7, 0, 0, 0, 9, 0, 11, 0, 0, 13, 13, 13, 13, 3, 0, 3, 8, 0, 8, 8, 0, 3, 0, 0}, [25]int16{17, 19, 0, 0, 0, 22, 0, 25, 0, 0, 28, 28, 28, 28, 3, 0, 3, 8, 0, 8, 8, 0, 3, 0, 0}, [25]int16{0, 7, 0, 0, 0, 9, 31, 11, 0, 0, 13, 13, 13, 13, 3, 0, 21, 8, 0, 8, 8, 0, 0, 0, 0}, struct {
	F0 [15]int16
	F1 [10]int16
}{[15]int16{33, 33, 33, 33, 33, 33, 33, 33, 0, 0, 33, 33, 33, 33, 3}, [10]int16{}}, struct {
	F0 [15]int16
	F1 [10]int16
}{[15]int16{35, 35, 35, 35, 35, 35, 35, 35, 0, 0, 35, 35, 35, 35, 3}, [10]int16{}}}
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
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F66 TSParseActionEntry
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
			F0 struct {
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F83 struct {
		F0 anon_2
		F1 [6]byte
	}
	F84 TSParseActionEntry
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
	F93 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
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
			F0 struct {
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
		F0 anon_2
		F1 [6]byte
	}
	F66 TSParseActionEntry
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
			F0 struct {
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
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F83 struct {
		F0 anon_2
		F1 [6]byte
	}
	F84 TSParseActionEntry
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
	F93 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 0, 15, 0, 0}}}, struct {
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
}{0, 0, 8, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 15, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 16, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 22, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 20, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 17, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 16, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 19, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 17, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 17, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 19, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 4, 19, 0, 0}}}, struct {
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
}{0, 0, 0, 1, 0}, [2]byte{}}}, struct {
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
}{anon_2{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 21, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 21, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
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
}{anon_2{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 23, 0, 0}}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 23, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 23, 0, 0}}}, struct {
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
}{0, 0, 10, 0, 1}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 2, 24, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 3, 18, 0, 1}}}, struct {
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
}{0, 0, 13, 0, 0}, [2]byte{}}}}
var _str [4]byte = [4]byte{101, 110, 100, 0}
var _str_3 [2]byte = [2]byte{123, 0}
var _str_4 [2]byte = [2]byte{44, 0}
var _str_5 [2]byte = [2]byte{125, 0}
var _str_6 [2]byte = [2]byte{58, 0}
var _str_7 [2]byte = [2]byte{91, 0}
var _str_8 [2]byte = [2]byte{93, 0}
var _str_9 [2]byte = [2]byte{34, 0}
var _str_10 [15]byte = [15]byte{115, 116, 114, 105, 110, 103, 95, 99, 111, 110, 116, 101, 110, 116, 0}
var _str_11 [16]byte = [16]byte{101, 115, 99, 97, 112, 101, 95, 115, 101, 113, 117, 101, 110, 99, 101, 0}
var _str_12 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}
var _str_13 [5]byte = [5]byte{116, 114, 117, 101, 0}
var _str_14 [6]byte = [6]byte{102, 97, 108, 115, 101, 0}
var _str_15 [5]byte = [5]byte{110, 117, 108, 108, 0}
var _str_16 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}
var _str_17 [9]byte = [9]byte{100, 111, 99, 117, 109, 101, 110, 116, 0}
var _str_18 [7]byte = [7]byte{95, 118, 97, 108, 117, 101, 0}
var _str_19 [7]byte = [7]byte{111, 98, 106, 101, 99, 116, 0}
var _str_20 [5]byte = [5]byte{112, 97, 105, 114, 0}
var _str_21 [6]byte = [6]byte{97, 114, 114, 97, 121, 0}
var _str_22 [7]byte = [7]byte{115, 116, 114, 105, 110, 103, 0}
var _str_23 [16]byte = [16]byte{95, 115, 116, 114, 105, 110, 103, 95, 99, 111, 110, 116, 101, 110, 116, 0}
var _str_24 [17]byte = [17]byte{100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_25 [15]byte = [15]byte{111, 98, 106, 101, 99, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_26 [14]byte = [14]byte{97, 114, 114, 97, 121, 95, 114, 101, 112, 101, 97, 116, 49, 0}
var _str_27 [4]byte = [4]byte{107, 101, 121, 0}
var _str_28 [6]byte = [6]byte{118, 97, 108, 117, 101, 0}
var ts_lex_modes struct {
	F0 [20]TSLexMode
	F1 [12]TSLexMode
} = struct {
	F0 [20]TSLexMode
	F1 [12]TSLexMode
}{[20]TSLexMode{TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{1, 0}, TSLexMode{1, 0}, TSLexMode{1, 0}}, [12]TSLexMode{}}
var ts_lex_map [28]int16 = [28]int16{34, 28, 44, 23, 45, 6, 47, 3, 48, 35, 58, 25, 91, 26, 92, 18, 93, 27, 102, 7, 110, 16, 116, 13, 123, 22, 125, 24}
var ts_lex_map_30 [18]int16 = [18]int16{34, 34, 47, 34, 92, 34, 98, 34, 102, 34, 110, 34, 114, 34, 116, 34, 117, 34}
var ts_lex_map_31 [26]int16 = [26]int16{34, 28, 44, 23, 45, 6, 47, 3, 48, 35, 58, 25, 91, 26, 93, 27, 102, 7, 110, 16, 116, 13, 123, 22, 125, 24}

func init() {
	tree_sitter_json_language = struct {
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
	}{14, 25, 0, 15, 0, 32, 7, 2, 2, 4, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), libc.Ptr(&ts_field_names), libc.Ptr(&ts_field_map_slices), libc.Ptr(&ts_field_map_entries), libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids)}
}
func tree_sitter_json() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_json_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, loadedv29, cmp31, cmp35, cmp39, cmp43, cmp47, cmp50, cmp53, cmp57, loadedv61, cmp63, cmp67, cmp71, cmp74, cmp77, loadedv81, cmp83, cmp87, loadedv91, cmp93, cmp97, cmp101, loadedv105, cmp107, cmp111, loadedv115, cmp117, cmp121, cmp124, loadedv128, cmp130, loadedv134, cmp136, loadedv140, cmp142, loadedv146, cmp148, loadedv152, cmp154, loadedv158, cmp160, loadedv164, cmp166, loadedv170, cmp172, loadedv176, cmp178, loadedv182, cmp184, loadedv188, cmp190, cmp193, cmp197, cmp200, loadedv204, cmp209, cmp215, loadedv225, cmp227, cmp230, loadedv234, loadedv236, cmp242, cmp248, cmp258, cmp261, cmp264, cmp268, cmp271, loadedv275, loadedv277, loadedv281, loadedv285, loadedv289, loadedv293, loadedv297, loadedv301, loadedv305, cmp309, cmp313, cmp317, cmp320, cmp323, cmp326, loadedv330, cmp334, cmp338, cmp342, cmp345, cmp348, cmp351, loadedv355, cmp359, cmp363, cmp366, cmp369, cmp372, loadedv376, cmp380, cmp384, cmp387, cmp390, cmp393, cmp397, cmp400, cmp403, cmp406, cmp409, loadedv413, cmp417, cmp420, cmp423, cmp426, loadedv430, loadedv434, cmp438, cmp442, cmp445, loadedv449, cmp453, cmp457, cmp460, cmp464, cmp467, loadedv471, cmp475, cmp478, cmp482, cmp485, loadedv489, cmp493, cmp496, loadedv500, loadedv504, loadedv508, loadedv512, loadedv516, cmp520, cmp523, loadedv527, v265 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v80, v83, v92, v95 int16
	var state_addr, arrayidx, arrayidx11, arrayidx213, arrayidx220, arrayidx246, arrayidx253, result_symbol, result_symbol279, result_symbol283, result_symbol287, result_symbol291, result_symbol295, result_symbol299, result_symbol303, result_symbol307, result_symbol332, result_symbol357, result_symbol378, result_symbol415, result_symbol432, result_symbol436, result_symbol451, result_symbol473, result_symbol491, result_symbol502, result_symbol506, result_symbol510, result_symbol514, result_symbol518 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v24, v25, v26, v27, v28, v29, v30, v31, v33, v34, v35, v36, v37, v39, v40, v42, v43, v44, v46, v47, v49, v50, v51, v53, v55, v57, v59, v61, v63, v65, v67, v69, v71, v73, v74, v75, v76, v78, v79, conv214, v81, v82, add218, v84, add223, v86, v87, v90, v91, conv247, v93, v94, add251, v96, add256, v97, v98, v99, v100, v101, v147, v148, v149, v150, v151, v152, v158, v159, v160, v161, v162, v163, v169, v170, v171, v172, v173, v179, v180, v181, v182, v183, v184, v185, v186, v187, v188, v194, v195, v196, v197, v208, v209, v210, v216, v217, v218, v219, v220, v226, v227, v228, v229, v235, v236, v262, v263 int32
	var lookahead, i, i206, i239, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv208, idxprom212, idxprom219, conv241, idxprom245, idxprom252 int64
	var v3, storedv, v10, v23, v32, v38, v41, v45, v48, v52, v54, v56, v58, v60, v62, v64, v66, v68, v70, v72, v77, v85, v88, v89, v102, v107, v112, v117, v122, v127, v132, v137, v142, v153, v164, v174, v189, v198, v203, v211, v221, v230, v237, v242, v247, v252, v257, v264 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v103, v104, v105, v106, v108, v109, v110, v111, v113, v114, v115, v116, v118, v119, v120, v121, v123, v124, v125, v126, v128, v129, v130, v131, v133, v134, v135, v136, v138, v139, v140, v141, v143, v144, v145, v146, v154, v155, v156, v157, v165, v166, v167, v168, v175, v176, v177, v178, v190, v191, v192, v193, v199, v200, v201, v202, v204, v205, v206, v207, v212, v213, v214, v215, v222, v223, v224, v225, v231, v232, v233, v234, v238, v239, v240, v241, v243, v244, v245, v246, v248, v249, v250, v251, v253, v254, v255, v256, v258, v259, v260, v261 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end280, mark_end284, mark_end288, mark_end292, mark_end296, mark_end300, mark_end304, mark_end308, mark_end333, mark_end358, mark_end379, mark_end416, mark_end433, mark_end437, mark_end452, mark_end474, mark_end492, mark_end503, mark_end507, mark_end511, mark_end515, mark_end519 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i206, i239, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, loadedv29, v24, cmp31, v25, cmp35, v26, cmp39, v27, cmp43, v28, cmp47, v29, cmp50, v30, cmp53, v31, cmp57, v32, loadedv61, v33, cmp63, v34, cmp67, v35, cmp71, v36, cmp74, v37, cmp77, v38, loadedv81, v39, cmp83, v40, cmp87, v41, loadedv91, v42, cmp93, v43, cmp97, v44, cmp101, v45, loadedv105, v46, cmp107, v47, cmp111, v48, loadedv115, v49, cmp117, v50, cmp121, v51, cmp124, v52, loadedv128, v53, cmp130, v54, loadedv134, v55, cmp136, v56, loadedv140, v57, cmp142, v58, loadedv146, v59, cmp148, v60, loadedv152, v61, cmp154, v62, loadedv158, v63, cmp160, v64, loadedv164, v65, cmp166, v66, loadedv170, v67, cmp172, v68, loadedv176, v69, cmp178, v70, loadedv182, v71, cmp184, v72, loadedv188, v73, cmp190, v74, cmp193, v75, cmp197, v76, cmp200, v77, loadedv204, v78, conv208, cmp209, v79, idxprom212, arrayidx213, v80, conv214, v81, cmp215, v82, add218, idxprom219, arrayidx220, v83, v84, add223, v85, loadedv225, v86, cmp227, v87, cmp230, v88, loadedv234, v89, loadedv236, v90, conv241, cmp242, v91, idxprom245, arrayidx246, v92, conv247, v93, cmp248, v94, add251, idxprom252, arrayidx253, v95, v96, add256, v97, cmp258, v98, cmp261, v99, cmp264, v100, cmp268, v101, cmp271, v102, loadedv275, v103, result_symbol, v104, mark_end, v105, v106, v107, loadedv277, v108, result_symbol279, v109, mark_end280, v110, v111, v112, loadedv281, v113, result_symbol283, v114, mark_end284, v115, v116, v117, loadedv285, v118, result_symbol287, v119, mark_end288, v120, v121, v122, loadedv289, v123, result_symbol291, v124, mark_end292, v125, v126, v127, loadedv293, v128, result_symbol295, v129, mark_end296, v130, v131, v132, loadedv297, v133, result_symbol299, v134, mark_end300, v135, v136, v137, loadedv301, v138, result_symbol303, v139, mark_end304, v140, v141, v142, loadedv305, v143, result_symbol307, v144, mark_end308, v145, v146, v147, cmp309, v148, cmp313, v149, cmp317, v150, cmp320, v151, cmp323, v152, cmp326, v153, loadedv330, v154, result_symbol332, v155, mark_end333, v156, v157, v158, cmp334, v159, cmp338, v160, cmp342, v161, cmp345, v162, cmp348, v163, cmp351, v164, loadedv355, v165, result_symbol357, v166, mark_end358, v167, v168, v169, cmp359, v170, cmp363, v171, cmp366, v172, cmp369, v173, cmp372, v174, loadedv376, v175, result_symbol378, v176, mark_end379, v177, v178, v179, cmp380, v180, cmp384, v181, cmp387, v182, cmp390, v183, cmp393, v184, cmp397, v185, cmp400, v186, cmp403, v187, cmp406, v188, cmp409, v189, loadedv413, v190, result_symbol415, v191, mark_end416, v192, v193, v194, cmp417, v195, cmp420, v196, cmp423, v197, cmp426, v198, loadedv430, v199, result_symbol432, v200, mark_end433, v201, v202, v203, loadedv434, v204, result_symbol436, v205, mark_end437, v206, v207, v208, cmp438, v209, cmp442, v210, cmp445, v211, loadedv449, v212, result_symbol451, v213, mark_end452, v214, v215, v216, cmp453, v217, cmp457, v218, cmp460, v219, cmp464, v220, cmp467, v221, loadedv471, v222, result_symbol473, v223, mark_end474, v224, v225, v226, cmp475, v227, cmp478, v228, cmp482, v229, cmp485, v230, loadedv489, v231, result_symbol491, v232, mark_end492, v233, v234, v235, cmp493, v236, cmp496, v237, loadedv500, v238, result_symbol502, v239, mark_end503, v240, v241, v242, loadedv504, v243, result_symbol506, v244, mark_end507, v245, v246, v247, loadedv508, v248, result_symbol510, v249, mark_end511, v250, v251, v252, loadedv512, v253, result_symbol514, v254, mark_end515, v255, v256, v257, loadedv516, v258, result_symbol518, v259, mark_end519, v260, v261, v262, cmp520, v263, cmp523, v264, loadedv527, v265

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
	i206 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i239 = libc.Ptr(&new(struct {
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
		goto sw_bb30
	case 2:
		goto sw_bb62
	case 3:
		goto sw_bb82
	case 4:
		goto sw_bb92
	case 5:
		goto sw_bb106
	case 6:
		goto sw_bb116
	case 7:
		goto sw_bb129
	case 8:
		goto sw_bb135
	case 9:
		goto sw_bb141
	case 10:
		goto sw_bb147
	case 11:
		goto sw_bb153
	case 12:
		goto sw_bb159
	case 13:
		goto sw_bb165
	case 14:
		goto sw_bb171
	case 15:
		goto sw_bb177
	case 16:
		goto sw_bb183
	case 17:
		goto sw_bb189
	case 18:
		goto sw_bb205
	case 19:
		goto sw_bb226
	case 20:
		goto sw_bb235
	case 21:
		goto sw_bb276
	case 22:
		goto sw_bb278
	case 23:
		goto sw_bb282
	case 24:
		goto sw_bb286
	case 25:
		goto sw_bb290
	case 26:
		goto sw_bb294
	case 27:
		goto sw_bb298
	case 28:
		goto sw_bb302
	case 29:
		goto sw_bb306
	case 30:
		goto sw_bb331
	case 31:
		goto sw_bb356
	case 32:
		goto sw_bb377
	case 33:
		goto sw_bb414
	case 34:
		goto sw_bb431
	case 35:
		goto sw_bb435
	case 36:
		goto sw_bb450
	case 37:
		goto sw_bb472
	case 38:
		goto sw_bb490
	case 39:
		goto sw_bb501
	case 40:
		goto sw_bb505
	case 41:
		goto sw_bb509
	case 42:
		goto sw_bb513
	case 43:
		goto sw_bb517
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
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v11 = *libc.As[int32](i)
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(28)
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
	*libc.As[int16](state_addr) = 20
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
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end28:
	v23 = *libc.As[byte](result)
	loadedv29 = (v23 & 1) != 0
	*libc.As[bool](retval) = loadedv29
	goto _return

sw_bb30:
	v24 = *libc.As[int32](lookahead)
	cmp31 = v24 == 10
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end34:
	v25 = *libc.As[int32](lookahead)
	cmp35 = v25 == 34
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end38:
	v26 = *libc.As[int32](lookahead)
	cmp39 = v26 == 47
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end42:
	v27 = *libc.As[int32](lookahead)
	cmp43 = v27 == 92
	if cmp43 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end46:
	v28 = *libc.As[int32](lookahead)
	cmp47 = 9 <= v28
	if cmp47 {
		goto land_lhs_true49
	} else {
		goto lor_lhs_false52
	}

land_lhs_true49:
	v29 = *libc.As[int32](lookahead)
	cmp50 = v29 <= 13
	if cmp50 {
		goto if_then55
	} else {
		goto lor_lhs_false52
	}

lor_lhs_false52:
	v30 = *libc.As[int32](lookahead)
	cmp53 = v30 == 32
	if cmp53 {
		goto if_then55
	} else {
		goto if_end56
	}

if_then55:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end56:
	v31 = *libc.As[int32](lookahead)
	cmp57 = v31 != 0
	if cmp57 {
		goto if_then59
	} else {
		goto if_end60
	}

if_then59:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end60:
	v32 = *libc.As[byte](result)
	loadedv61 = (v32 & 1) != 0
	*libc.As[bool](retval) = loadedv61
	goto _return

sw_bb62:
	v33 = *libc.As[int32](lookahead)
	cmp63 = v33 == 34
	if cmp63 {
		goto if_then65
	} else {
		goto if_end66
	}

if_then65:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end66:
	v34 = *libc.As[int32](lookahead)
	cmp67 = v34 == 47
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end70:
	v35 = *libc.As[int32](lookahead)
	cmp71 = 9 <= v35
	if cmp71 {
		goto land_lhs_true73
	} else {
		goto lor_lhs_false76
	}

land_lhs_true73:
	v36 = *libc.As[int32](lookahead)
	cmp74 = v36 <= 13
	if cmp74 {
		goto if_then79
	} else {
		goto lor_lhs_false76
	}

lor_lhs_false76:
	v37 = *libc.As[int32](lookahead)
	cmp77 = v37 == 32
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
	v38 = *libc.As[byte](result)
	loadedv81 = (v38 & 1) != 0
	*libc.As[bool](retval) = loadedv81
	goto _return

sw_bb82:
	v39 = *libc.As[int32](lookahead)
	cmp83 = v39 == 42
	if cmp83 {
		goto if_then85
	} else {
		goto if_end86
	}

if_then85:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end86:
	v40 = *libc.As[int32](lookahead)
	cmp87 = v40 == 47
	if cmp87 {
		goto if_then89
	} else {
		goto if_end90
	}

if_then89:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end90:
	v41 = *libc.As[byte](result)
	loadedv91 = (v41 & 1) != 0
	*libc.As[bool](retval) = loadedv91
	goto _return

sw_bb92:
	v42 = *libc.As[int32](lookahead)
	cmp93 = v42 == 42
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end96:
	v43 = *libc.As[int32](lookahead)
	cmp97 = v43 == 47
	if cmp97 {
		goto if_then99
	} else {
		goto if_end100
	}

if_then99:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end100:
	v44 = *libc.As[int32](lookahead)
	cmp101 = v44 != 0
	if cmp101 {
		goto if_then103
	} else {
		goto if_end104
	}

if_then103:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end104:
	v45 = *libc.As[byte](result)
	loadedv105 = (v45 & 1) != 0
	*libc.As[bool](retval) = loadedv105
	goto _return

sw_bb106:
	v46 = *libc.As[int32](lookahead)
	cmp107 = v46 == 42
	if cmp107 {
		goto if_then109
	} else {
		goto if_end110
	}

if_then109:
	*libc.As[int16](state_addr) = 4
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
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end114:
	v48 = *libc.As[byte](result)
	loadedv115 = (v48 & 1) != 0
	*libc.As[bool](retval) = loadedv115
	goto _return

sw_bb116:
	v49 = *libc.As[int32](lookahead)
	cmp117 = v49 == 48
	if cmp117 {
		goto if_then119
	} else {
		goto if_end120
	}

if_then119:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end120:
	v50 = *libc.As[int32](lookahead)
	cmp121 = 49 <= v50
	if cmp121 {
		goto land_lhs_true123
	} else {
		goto if_end127
	}

land_lhs_true123:
	v51 = *libc.As[int32](lookahead)
	cmp124 = v51 <= 57
	if cmp124 {
		goto if_then126
	} else {
		goto if_end127
	}

if_then126:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end127:
	v52 = *libc.As[byte](result)
	loadedv128 = (v52 & 1) != 0
	*libc.As[bool](retval) = loadedv128
	goto _return

sw_bb129:
	v53 = *libc.As[int32](lookahead)
	cmp130 = v53 == 97
	if cmp130 {
		goto if_then132
	} else {
		goto if_end133
	}

if_then132:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end133:
	v54 = *libc.As[byte](result)
	loadedv134 = (v54 & 1) != 0
	*libc.As[bool](retval) = loadedv134
	goto _return

sw_bb135:
	v55 = *libc.As[int32](lookahead)
	cmp136 = v55 == 101
	if cmp136 {
		goto if_then138
	} else {
		goto if_end139
	}

if_then138:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end139:
	v56 = *libc.As[byte](result)
	loadedv140 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv140
	goto _return

sw_bb141:
	v57 = *libc.As[int32](lookahead)
	cmp142 = v57 == 101
	if cmp142 {
		goto if_then144
	} else {
		goto if_end145
	}

if_then144:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end145:
	v58 = *libc.As[byte](result)
	loadedv146 = (v58 & 1) != 0
	*libc.As[bool](retval) = loadedv146
	goto _return

sw_bb147:
	v59 = *libc.As[int32](lookahead)
	cmp148 = v59 == 108
	if cmp148 {
		goto if_then150
	} else {
		goto if_end151
	}

if_then150:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end151:
	v60 = *libc.As[byte](result)
	loadedv152 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv152
	goto _return

sw_bb153:
	v61 = *libc.As[int32](lookahead)
	cmp154 = v61 == 108
	if cmp154 {
		goto if_then156
	} else {
		goto if_end157
	}

if_then156:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end157:
	v62 = *libc.As[byte](result)
	loadedv158 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv158
	goto _return

sw_bb159:
	v63 = *libc.As[int32](lookahead)
	cmp160 = v63 == 108
	if cmp160 {
		goto if_then162
	} else {
		goto if_end163
	}

if_then162:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end163:
	v64 = *libc.As[byte](result)
	loadedv164 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv164
	goto _return

sw_bb165:
	v65 = *libc.As[int32](lookahead)
	cmp166 = v65 == 114
	if cmp166 {
		goto if_then168
	} else {
		goto if_end169
	}

if_then168:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end169:
	v66 = *libc.As[byte](result)
	loadedv170 = (v66 & 1) != 0
	*libc.As[bool](retval) = loadedv170
	goto _return

sw_bb171:
	v67 = *libc.As[int32](lookahead)
	cmp172 = v67 == 115
	if cmp172 {
		goto if_then174
	} else {
		goto if_end175
	}

if_then174:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end175:
	v68 = *libc.As[byte](result)
	loadedv176 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv176
	goto _return

sw_bb177:
	v69 = *libc.As[int32](lookahead)
	cmp178 = v69 == 117
	if cmp178 {
		goto if_then180
	} else {
		goto if_end181
	}

if_then180:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end181:
	v70 = *libc.As[byte](result)
	loadedv182 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv182
	goto _return

sw_bb183:
	v71 = *libc.As[int32](lookahead)
	cmp184 = v71 == 117
	if cmp184 {
		goto if_then186
	} else {
		goto if_end187
	}

if_then186:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end187:
	v72 = *libc.As[byte](result)
	loadedv188 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv188
	goto _return

sw_bb189:
	v73 = *libc.As[int32](lookahead)
	cmp190 = v73 == 43
	if cmp190 {
		goto if_then195
	} else {
		goto lor_lhs_false192
	}

lor_lhs_false192:
	v74 = *libc.As[int32](lookahead)
	cmp193 = v74 == 45
	if cmp193 {
		goto if_then195
	} else {
		goto if_end196
	}

if_then195:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end196:
	v75 = *libc.As[int32](lookahead)
	cmp197 = 48 <= v75
	if cmp197 {
		goto land_lhs_true199
	} else {
		goto if_end203
	}

land_lhs_true199:
	v76 = *libc.As[int32](lookahead)
	cmp200 = v76 <= 57
	if cmp200 {
		goto if_then202
	} else {
		goto if_end203
	}

if_then202:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end203:
	v77 = *libc.As[byte](result)
	loadedv204 = (v77 & 1) != 0
	*libc.As[bool](retval) = loadedv204
	goto _return

sw_bb205:
	*libc.As[int32](i206) = 0
	goto for_cond207

for_cond207:
	v78 = *libc.As[int32](i206)
	conv208 = int64(uint64(uint32(v78)))
	cmp209 = uint64(conv208) < uint64(18)
	if cmp209 {
		goto for_body211
	} else {
		goto for_end224
	}

for_body211:
	v79 = *libc.As[int32](i206)
	idxprom212 = int64(uint64(uint32(v79)))
	arrayidx213 = libc.Ptr(&ts_lex_map_30[idxprom212])
	v80 = *libc.As[int16](arrayidx213)
	conv214 = int32(uint32(uint16(v80)))
	v81 = *libc.As[int32](lookahead)
	cmp215 = conv214 == v81
	if cmp215 {
		goto if_then217
	} else {
		goto if_end221
	}

if_then217:
	v82 = *libc.As[int32](i206)
	add218 = v82 + 1
	idxprom219 = int64(uint64(uint32(add218)))
	arrayidx220 = libc.Ptr(&ts_lex_map_30[idxprom219])
	v83 = *libc.As[int16](arrayidx220)
	*libc.As[int16](state_addr) = v83
	goto next_state

if_end221:
	goto for_inc222

for_inc222:
	v84 = *libc.As[int32](i206)
	add223 = v84 + 2
	*libc.As[int32](i206) = add223
	goto for_cond207

for_end224:
	v85 = *libc.As[byte](result)
	loadedv225 = (v85 & 1) != 0
	*libc.As[bool](retval) = loadedv225
	goto _return

sw_bb226:
	v86 = *libc.As[int32](lookahead)
	cmp227 = 48 <= v86
	if cmp227 {
		goto land_lhs_true229
	} else {
		goto if_end233
	}

land_lhs_true229:
	v87 = *libc.As[int32](lookahead)
	cmp230 = v87 <= 57
	if cmp230 {
		goto if_then232
	} else {
		goto if_end233
	}

if_then232:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end233:
	v88 = *libc.As[byte](result)
	loadedv234 = (v88 & 1) != 0
	*libc.As[bool](retval) = loadedv234
	goto _return

sw_bb235:
	v89 = *libc.As[byte](eof)
	loadedv236 = (v89 & 1) != 0
	if loadedv236 {
		goto if_then237
	} else {
		goto if_end238
	}

if_then237:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end238:
	*libc.As[int32](i239) = 0
	goto for_cond240

for_cond240:
	v90 = *libc.As[int32](i239)
	conv241 = int64(uint64(uint32(v90)))
	cmp242 = uint64(conv241) < uint64(26)
	if cmp242 {
		goto for_body244
	} else {
		goto for_end257
	}

for_body244:
	v91 = *libc.As[int32](i239)
	idxprom245 = int64(uint64(uint32(v91)))
	arrayidx246 = libc.Ptr(&ts_lex_map_31[idxprom245])
	v92 = *libc.As[int16](arrayidx246)
	conv247 = int32(uint32(uint16(v92)))
	v93 = *libc.As[int32](lookahead)
	cmp248 = conv247 == v93
	if cmp248 {
		goto if_then250
	} else {
		goto if_end254
	}

if_then250:
	v94 = *libc.As[int32](i239)
	add251 = v94 + 1
	idxprom252 = int64(uint64(uint32(add251)))
	arrayidx253 = libc.Ptr(&ts_lex_map_31[idxprom252])
	v95 = *libc.As[int16](arrayidx253)
	*libc.As[int16](state_addr) = v95
	goto next_state

if_end254:
	goto for_inc255

for_inc255:
	v96 = *libc.As[int32](i239)
	add256 = v96 + 2
	*libc.As[int32](i239) = add256
	goto for_cond240

for_end257:
	v97 = *libc.As[int32](lookahead)
	cmp258 = 9 <= v97
	if cmp258 {
		goto land_lhs_true260
	} else {
		goto lor_lhs_false263
	}

land_lhs_true260:
	v98 = *libc.As[int32](lookahead)
	cmp261 = v98 <= 13
	if cmp261 {
		goto if_then266
	} else {
		goto lor_lhs_false263
	}

lor_lhs_false263:
	v99 = *libc.As[int32](lookahead)
	cmp264 = v99 == 32
	if cmp264 {
		goto if_then266
	} else {
		goto if_end267
	}

if_then266:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end267:
	v100 = *libc.As[int32](lookahead)
	cmp268 = 49 <= v100
	if cmp268 {
		goto land_lhs_true270
	} else {
		goto if_end274
	}

land_lhs_true270:
	v101 = *libc.As[int32](lookahead)
	cmp271 = v101 <= 57
	if cmp271 {
		goto if_then273
	} else {
		goto if_end274
	}

if_then273:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end274:
	v102 = *libc.As[byte](result)
	loadedv275 = (v102 & 1) != 0
	*libc.As[bool](retval) = loadedv275
	goto _return

sw_bb276:
	*libc.As[byte](result) = 1
	v103 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v103).F1)
	*libc.As[int16](result_symbol) = 0
	v104 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v104).F3)
	v105 = *libc.As[unsafe.Pointer](mark_end)
	v106 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v105)(v106)
	v107 = *libc.As[byte](result)
	loadedv277 = (v107 & 1) != 0
	*libc.As[bool](retval) = loadedv277
	goto _return

sw_bb278:
	*libc.As[byte](result) = 1
	v108 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol279 = libc.Ptr(&libc.As[TSLexer](v108).F1)
	*libc.As[int16](result_symbol279) = 1
	v109 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end280 = libc.Ptr(&libc.As[TSLexer](v109).F3)
	v110 = *libc.As[unsafe.Pointer](mark_end280)
	v111 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v110)(v111)
	v112 = *libc.As[byte](result)
	loadedv281 = (v112 & 1) != 0
	*libc.As[bool](retval) = loadedv281
	goto _return

sw_bb282:
	*libc.As[byte](result) = 1
	v113 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol283 = libc.Ptr(&libc.As[TSLexer](v113).F1)
	*libc.As[int16](result_symbol283) = 2
	v114 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end284 = libc.Ptr(&libc.As[TSLexer](v114).F3)
	v115 = *libc.As[unsafe.Pointer](mark_end284)
	v116 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v115)(v116)
	v117 = *libc.As[byte](result)
	loadedv285 = (v117 & 1) != 0
	*libc.As[bool](retval) = loadedv285
	goto _return

sw_bb286:
	*libc.As[byte](result) = 1
	v118 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol287 = libc.Ptr(&libc.As[TSLexer](v118).F1)
	*libc.As[int16](result_symbol287) = 3
	v119 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end288 = libc.Ptr(&libc.As[TSLexer](v119).F3)
	v120 = *libc.As[unsafe.Pointer](mark_end288)
	v121 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v120)(v121)
	v122 = *libc.As[byte](result)
	loadedv289 = (v122 & 1) != 0
	*libc.As[bool](retval) = loadedv289
	goto _return

sw_bb290:
	*libc.As[byte](result) = 1
	v123 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol291 = libc.Ptr(&libc.As[TSLexer](v123).F1)
	*libc.As[int16](result_symbol291) = 4
	v124 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end292 = libc.Ptr(&libc.As[TSLexer](v124).F3)
	v125 = *libc.As[unsafe.Pointer](mark_end292)
	v126 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v125)(v126)
	v127 = *libc.As[byte](result)
	loadedv293 = (v127 & 1) != 0
	*libc.As[bool](retval) = loadedv293
	goto _return

sw_bb294:
	*libc.As[byte](result) = 1
	v128 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol295 = libc.Ptr(&libc.As[TSLexer](v128).F1)
	*libc.As[int16](result_symbol295) = 5
	v129 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end296 = libc.Ptr(&libc.As[TSLexer](v129).F3)
	v130 = *libc.As[unsafe.Pointer](mark_end296)
	v131 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v130)(v131)
	v132 = *libc.As[byte](result)
	loadedv297 = (v132 & 1) != 0
	*libc.As[bool](retval) = loadedv297
	goto _return

sw_bb298:
	*libc.As[byte](result) = 1
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol299 = libc.Ptr(&libc.As[TSLexer](v133).F1)
	*libc.As[int16](result_symbol299) = 6
	v134 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end300 = libc.Ptr(&libc.As[TSLexer](v134).F3)
	v135 = *libc.As[unsafe.Pointer](mark_end300)
	v136 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v135)(v136)
	v137 = *libc.As[byte](result)
	loadedv301 = (v137 & 1) != 0
	*libc.As[bool](retval) = loadedv301
	goto _return

sw_bb302:
	*libc.As[byte](result) = 1
	v138 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol303 = libc.Ptr(&libc.As[TSLexer](v138).F1)
	*libc.As[int16](result_symbol303) = 7
	v139 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end304 = libc.Ptr(&libc.As[TSLexer](v139).F3)
	v140 = *libc.As[unsafe.Pointer](mark_end304)
	v141 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v140)(v141)
	v142 = *libc.As[byte](result)
	loadedv305 = (v142 & 1) != 0
	*libc.As[bool](retval) = loadedv305
	goto _return

sw_bb306:
	*libc.As[byte](result) = 1
	v143 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol307 = libc.Ptr(&libc.As[TSLexer](v143).F1)
	*libc.As[int16](result_symbol307) = 8
	v144 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end308 = libc.Ptr(&libc.As[TSLexer](v144).F3)
	v145 = *libc.As[unsafe.Pointer](mark_end308)
	v146 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v145)(v146)
	v147 = *libc.As[int32](lookahead)
	cmp309 = v147 == 42
	if cmp309 {
		goto if_then311
	} else {
		goto if_end312
	}

if_then311:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end312:
	v148 = *libc.As[int32](lookahead)
	cmp313 = v148 == 47
	if cmp313 {
		goto if_then315
	} else {
		goto if_end316
	}

if_then315:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end316:
	v149 = *libc.As[int32](lookahead)
	cmp317 = v149 != 0
	if cmp317 {
		goto land_lhs_true319
	} else {
		goto if_end329
	}

land_lhs_true319:
	v150 = *libc.As[int32](lookahead)
	cmp320 = v150 != 10
	if cmp320 {
		goto land_lhs_true322
	} else {
		goto if_end329
	}

land_lhs_true322:
	v151 = *libc.As[int32](lookahead)
	cmp323 = v151 != 34
	if cmp323 {
		goto land_lhs_true325
	} else {
		goto if_end329
	}

land_lhs_true325:
	v152 = *libc.As[int32](lookahead)
	cmp326 = v152 != 92
	if cmp326 {
		goto if_then328
	} else {
		goto if_end329
	}

if_then328:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end329:
	v153 = *libc.As[byte](result)
	loadedv330 = (v153 & 1) != 0
	*libc.As[bool](retval) = loadedv330
	goto _return

sw_bb331:
	*libc.As[byte](result) = 1
	v154 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol332 = libc.Ptr(&libc.As[TSLexer](v154).F1)
	*libc.As[int16](result_symbol332) = 8
	v155 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end333 = libc.Ptr(&libc.As[TSLexer](v155).F3)
	v156 = *libc.As[unsafe.Pointer](mark_end333)
	v157 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v156)(v157)
	v158 = *libc.As[int32](lookahead)
	cmp334 = v158 == 42
	if cmp334 {
		goto if_then336
	} else {
		goto if_end337
	}

if_then336:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end337:
	v159 = *libc.As[int32](lookahead)
	cmp338 = v159 == 47
	if cmp338 {
		goto if_then340
	} else {
		goto if_end341
	}

if_then340:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end341:
	v160 = *libc.As[int32](lookahead)
	cmp342 = v160 != 0
	if cmp342 {
		goto land_lhs_true344
	} else {
		goto if_end354
	}

land_lhs_true344:
	v161 = *libc.As[int32](lookahead)
	cmp345 = v161 != 10
	if cmp345 {
		goto land_lhs_true347
	} else {
		goto if_end354
	}

land_lhs_true347:
	v162 = *libc.As[int32](lookahead)
	cmp348 = v162 != 34
	if cmp348 {
		goto land_lhs_true350
	} else {
		goto if_end354
	}

land_lhs_true350:
	v163 = *libc.As[int32](lookahead)
	cmp351 = v163 != 92
	if cmp351 {
		goto if_then353
	} else {
		goto if_end354
	}

if_then353:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end354:
	v164 = *libc.As[byte](result)
	loadedv355 = (v164 & 1) != 0
	*libc.As[bool](retval) = loadedv355
	goto _return

sw_bb356:
	*libc.As[byte](result) = 1
	v165 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol357 = libc.Ptr(&libc.As[TSLexer](v165).F1)
	*libc.As[int16](result_symbol357) = 8
	v166 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end358 = libc.Ptr(&libc.As[TSLexer](v166).F3)
	v167 = *libc.As[unsafe.Pointer](mark_end358)
	v168 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v167)(v168)
	v169 = *libc.As[int32](lookahead)
	cmp359 = v169 == 42
	if cmp359 {
		goto if_then361
	} else {
		goto if_end362
	}

if_then361:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end362:
	v170 = *libc.As[int32](lookahead)
	cmp363 = v170 != 0
	if cmp363 {
		goto land_lhs_true365
	} else {
		goto if_end375
	}

land_lhs_true365:
	v171 = *libc.As[int32](lookahead)
	cmp366 = v171 != 10
	if cmp366 {
		goto land_lhs_true368
	} else {
		goto if_end375
	}

land_lhs_true368:
	v172 = *libc.As[int32](lookahead)
	cmp369 = v172 != 34
	if cmp369 {
		goto land_lhs_true371
	} else {
		goto if_end375
	}

land_lhs_true371:
	v173 = *libc.As[int32](lookahead)
	cmp372 = v173 != 92
	if cmp372 {
		goto if_then374
	} else {
		goto if_end375
	}

if_then374:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end375:
	v174 = *libc.As[byte](result)
	loadedv376 = (v174 & 1) != 0
	*libc.As[bool](retval) = loadedv376
	goto _return

sw_bb377:
	*libc.As[byte](result) = 1
	v175 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol378 = libc.Ptr(&libc.As[TSLexer](v175).F1)
	*libc.As[int16](result_symbol378) = 8
	v176 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end379 = libc.Ptr(&libc.As[TSLexer](v176).F3)
	v177 = *libc.As[unsafe.Pointer](mark_end379)
	v178 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v177)(v178)
	v179 = *libc.As[int32](lookahead)
	cmp380 = v179 == 47
	if cmp380 {
		goto if_then382
	} else {
		goto if_end383
	}

if_then382:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end383:
	v180 = *libc.As[int32](lookahead)
	cmp384 = v180 == 9
	if cmp384 {
		goto if_then395
	} else {
		goto lor_lhs_false386
	}

lor_lhs_false386:
	v181 = *libc.As[int32](lookahead)
	cmp387 = 11 <= v181
	if cmp387 {
		goto land_lhs_true389
	} else {
		goto lor_lhs_false392
	}

land_lhs_true389:
	v182 = *libc.As[int32](lookahead)
	cmp390 = v182 <= 13
	if cmp390 {
		goto if_then395
	} else {
		goto lor_lhs_false392
	}

lor_lhs_false392:
	v183 = *libc.As[int32](lookahead)
	cmp393 = v183 == 32
	if cmp393 {
		goto if_then395
	} else {
		goto if_end396
	}

if_then395:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end396:
	v184 = *libc.As[int32](lookahead)
	cmp397 = v184 != 0
	if cmp397 {
		goto land_lhs_true399
	} else {
		goto if_end412
	}

land_lhs_true399:
	v185 = *libc.As[int32](lookahead)
	cmp400 = v185 < 9
	if cmp400 {
		goto land_lhs_true405
	} else {
		goto lor_lhs_false402
	}

lor_lhs_false402:
	v186 = *libc.As[int32](lookahead)
	cmp403 = 13 < v186
	if cmp403 {
		goto land_lhs_true405
	} else {
		goto if_end412
	}

land_lhs_true405:
	v187 = *libc.As[int32](lookahead)
	cmp406 = v187 != 34
	if cmp406 {
		goto land_lhs_true408
	} else {
		goto if_end412
	}

land_lhs_true408:
	v188 = *libc.As[int32](lookahead)
	cmp409 = v188 != 92
	if cmp409 {
		goto if_then411
	} else {
		goto if_end412
	}

if_then411:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end412:
	v189 = *libc.As[byte](result)
	loadedv413 = (v189 & 1) != 0
	*libc.As[bool](retval) = loadedv413
	goto _return

sw_bb414:
	*libc.As[byte](result) = 1
	v190 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol415 = libc.Ptr(&libc.As[TSLexer](v190).F1)
	*libc.As[int16](result_symbol415) = 8
	v191 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end416 = libc.Ptr(&libc.As[TSLexer](v191).F3)
	v192 = *libc.As[unsafe.Pointer](mark_end416)
	v193 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v192)(v193)
	v194 = *libc.As[int32](lookahead)
	cmp417 = v194 != 0
	if cmp417 {
		goto land_lhs_true419
	} else {
		goto if_end429
	}

land_lhs_true419:
	v195 = *libc.As[int32](lookahead)
	cmp420 = v195 != 10
	if cmp420 {
		goto land_lhs_true422
	} else {
		goto if_end429
	}

land_lhs_true422:
	v196 = *libc.As[int32](lookahead)
	cmp423 = v196 != 34
	if cmp423 {
		goto land_lhs_true425
	} else {
		goto if_end429
	}

land_lhs_true425:
	v197 = *libc.As[int32](lookahead)
	cmp426 = v197 != 92
	if cmp426 {
		goto if_then428
	} else {
		goto if_end429
	}

if_then428:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end429:
	v198 = *libc.As[byte](result)
	loadedv430 = (v198 & 1) != 0
	*libc.As[bool](retval) = loadedv430
	goto _return

sw_bb431:
	*libc.As[byte](result) = 1
	v199 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol432 = libc.Ptr(&libc.As[TSLexer](v199).F1)
	*libc.As[int16](result_symbol432) = 9
	v200 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end433 = libc.Ptr(&libc.As[TSLexer](v200).F3)
	v201 = *libc.As[unsafe.Pointer](mark_end433)
	v202 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v201)(v202)
	v203 = *libc.As[byte](result)
	loadedv434 = (v203 & 1) != 0
	*libc.As[bool](retval) = loadedv434
	goto _return

sw_bb435:
	*libc.As[byte](result) = 1
	v204 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol436 = libc.Ptr(&libc.As[TSLexer](v204).F1)
	*libc.As[int16](result_symbol436) = 10
	v205 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end437 = libc.Ptr(&libc.As[TSLexer](v205).F3)
	v206 = *libc.As[unsafe.Pointer](mark_end437)
	v207 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v206)(v207)
	v208 = *libc.As[int32](lookahead)
	cmp438 = v208 == 46
	if cmp438 {
		goto if_then440
	} else {
		goto if_end441
	}

if_then440:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end441:
	v209 = *libc.As[int32](lookahead)
	cmp442 = v209 == 69
	if cmp442 {
		goto if_then447
	} else {
		goto lor_lhs_false444
	}

lor_lhs_false444:
	v210 = *libc.As[int32](lookahead)
	cmp445 = v210 == 101
	if cmp445 {
		goto if_then447
	} else {
		goto if_end448
	}

if_then447:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end448:
	v211 = *libc.As[byte](result)
	loadedv449 = (v211 & 1) != 0
	*libc.As[bool](retval) = loadedv449
	goto _return

sw_bb450:
	*libc.As[byte](result) = 1
	v212 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol451 = libc.Ptr(&libc.As[TSLexer](v212).F1)
	*libc.As[int16](result_symbol451) = 10
	v213 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end452 = libc.Ptr(&libc.As[TSLexer](v213).F3)
	v214 = *libc.As[unsafe.Pointer](mark_end452)
	v215 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v214)(v215)
	v216 = *libc.As[int32](lookahead)
	cmp453 = v216 == 46
	if cmp453 {
		goto if_then455
	} else {
		goto if_end456
	}

if_then455:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end456:
	v217 = *libc.As[int32](lookahead)
	cmp457 = v217 == 69
	if cmp457 {
		goto if_then462
	} else {
		goto lor_lhs_false459
	}

lor_lhs_false459:
	v218 = *libc.As[int32](lookahead)
	cmp460 = v218 == 101
	if cmp460 {
		goto if_then462
	} else {
		goto if_end463
	}

if_then462:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end463:
	v219 = *libc.As[int32](lookahead)
	cmp464 = 48 <= v219
	if cmp464 {
		goto land_lhs_true466
	} else {
		goto if_end470
	}

land_lhs_true466:
	v220 = *libc.As[int32](lookahead)
	cmp467 = v220 <= 57
	if cmp467 {
		goto if_then469
	} else {
		goto if_end470
	}

if_then469:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end470:
	v221 = *libc.As[byte](result)
	loadedv471 = (v221 & 1) != 0
	*libc.As[bool](retval) = loadedv471
	goto _return

sw_bb472:
	*libc.As[byte](result) = 1
	v222 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol473 = libc.Ptr(&libc.As[TSLexer](v222).F1)
	*libc.As[int16](result_symbol473) = 10
	v223 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end474 = libc.Ptr(&libc.As[TSLexer](v223).F3)
	v224 = *libc.As[unsafe.Pointer](mark_end474)
	v225 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v224)(v225)
	v226 = *libc.As[int32](lookahead)
	cmp475 = v226 == 69
	if cmp475 {
		goto if_then480
	} else {
		goto lor_lhs_false477
	}

lor_lhs_false477:
	v227 = *libc.As[int32](lookahead)
	cmp478 = v227 == 101
	if cmp478 {
		goto if_then480
	} else {
		goto if_end481
	}

if_then480:
	*libc.As[int16](state_addr) = 17
	goto next_state

if_end481:
	v228 = *libc.As[int32](lookahead)
	cmp482 = 48 <= v228
	if cmp482 {
		goto land_lhs_true484
	} else {
		goto if_end488
	}

land_lhs_true484:
	v229 = *libc.As[int32](lookahead)
	cmp485 = v229 <= 57
	if cmp485 {
		goto if_then487
	} else {
		goto if_end488
	}

if_then487:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end488:
	v230 = *libc.As[byte](result)
	loadedv489 = (v230 & 1) != 0
	*libc.As[bool](retval) = loadedv489
	goto _return

sw_bb490:
	*libc.As[byte](result) = 1
	v231 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol491 = libc.Ptr(&libc.As[TSLexer](v231).F1)
	*libc.As[int16](result_symbol491) = 10
	v232 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end492 = libc.Ptr(&libc.As[TSLexer](v232).F3)
	v233 = *libc.As[unsafe.Pointer](mark_end492)
	v234 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v233)(v234)
	v235 = *libc.As[int32](lookahead)
	cmp493 = 48 <= v235
	if cmp493 {
		goto land_lhs_true495
	} else {
		goto if_end499
	}

land_lhs_true495:
	v236 = *libc.As[int32](lookahead)
	cmp496 = v236 <= 57
	if cmp496 {
		goto if_then498
	} else {
		goto if_end499
	}

if_then498:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end499:
	v237 = *libc.As[byte](result)
	loadedv500 = (v237 & 1) != 0
	*libc.As[bool](retval) = loadedv500
	goto _return

sw_bb501:
	*libc.As[byte](result) = 1
	v238 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol502 = libc.Ptr(&libc.As[TSLexer](v238).F1)
	*libc.As[int16](result_symbol502) = 11
	v239 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end503 = libc.Ptr(&libc.As[TSLexer](v239).F3)
	v240 = *libc.As[unsafe.Pointer](mark_end503)
	v241 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v240)(v241)
	v242 = *libc.As[byte](result)
	loadedv504 = (v242 & 1) != 0
	*libc.As[bool](retval) = loadedv504
	goto _return

sw_bb505:
	*libc.As[byte](result) = 1
	v243 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol506 = libc.Ptr(&libc.As[TSLexer](v243).F1)
	*libc.As[int16](result_symbol506) = 12
	v244 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end507 = libc.Ptr(&libc.As[TSLexer](v244).F3)
	v245 = *libc.As[unsafe.Pointer](mark_end507)
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v245)(v246)
	v247 = *libc.As[byte](result)
	loadedv508 = (v247 & 1) != 0
	*libc.As[bool](retval) = loadedv508
	goto _return

sw_bb509:
	*libc.As[byte](result) = 1
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol510 = libc.Ptr(&libc.As[TSLexer](v248).F1)
	*libc.As[int16](result_symbol510) = 13
	v249 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end511 = libc.Ptr(&libc.As[TSLexer](v249).F3)
	v250 = *libc.As[unsafe.Pointer](mark_end511)
	v251 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v250)(v251)
	v252 = *libc.As[byte](result)
	loadedv512 = (v252 & 1) != 0
	*libc.As[bool](retval) = loadedv512
	goto _return

sw_bb513:
	*libc.As[byte](result) = 1
	v253 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol514 = libc.Ptr(&libc.As[TSLexer](v253).F1)
	*libc.As[int16](result_symbol514) = 14
	v254 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end515 = libc.Ptr(&libc.As[TSLexer](v254).F3)
	v255 = *libc.As[unsafe.Pointer](mark_end515)
	v256 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v255)(v256)
	v257 = *libc.As[byte](result)
	loadedv516 = (v257 & 1) != 0
	*libc.As[bool](retval) = loadedv516
	goto _return

sw_bb517:
	*libc.As[byte](result) = 1
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol518 = libc.Ptr(&libc.As[TSLexer](v258).F1)
	*libc.As[int16](result_symbol518) = 14
	v259 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end519 = libc.Ptr(&libc.As[TSLexer](v259).F3)
	v260 = *libc.As[unsafe.Pointer](mark_end519)
	v261 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v260)(v261)
	v262 = *libc.As[int32](lookahead)
	cmp520 = v262 != 0
	if cmp520 {
		goto land_lhs_true522
	} else {
		goto if_end526
	}

land_lhs_true522:
	v263 = *libc.As[int32](lookahead)
	cmp523 = v263 != 10
	if cmp523 {
		goto if_then525
	} else {
		goto if_end526
	}

if_then525:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end526:
	v264 = *libc.As[byte](result)
	loadedv527 = (v264 & 1) != 0
	*libc.As[bool](retval) = loadedv527
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v265 = *libc.As[bool](retval)
	return v265
}
