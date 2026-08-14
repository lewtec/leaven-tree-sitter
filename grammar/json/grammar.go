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
var ts_lex_map [28]int16 = [28]int16{34, 28, 44, 23, 45, 7, 47, 3, 48, 35, 58, 25, 91, 26, 92, 18, 93, 27, 102, 8, 110, 17, 116, 14, 123, 22, 125, 24}
var ts_lex_map_30 [18]int16 = [18]int16{34, 34, 47, 34, 92, 34, 98, 34, 102, 34, 110, 34, 114, 34, 116, 34, 117, 34}
var ts_lex_map_31 [26]int16 = [26]int16{34, 28, 44, 23, 45, 7, 47, 3, 48, 35, 58, 25, 91, 26, 93, 27, 102, 8, 110, 17, 116, 14, 123, 22, 125, 24}

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
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, loadedv29, cmp31, cmp35, cmp39, cmp43, cmp47, cmp50, cmp53, cmp57, loadedv61, cmp63, cmp67, cmp71, cmp74, cmp77, loadedv81, cmp83, cmp87, loadedv91, cmp93, cmp97, cmp101, loadedv105, cmp107, cmp111, loadedv115, cmp117, cmp121, cmp124, loadedv128, cmp130, cmp134, cmp137, loadedv141, cmp143, loadedv147, cmp149, loadedv153, cmp155, loadedv159, cmp161, loadedv165, cmp167, loadedv171, cmp173, loadedv177, cmp179, loadedv183, cmp185, loadedv189, cmp191, loadedv195, cmp197, loadedv201, cmp206, cmp212, loadedv222, cmp224, cmp227, loadedv231, loadedv233, cmp239, cmp245, cmp255, cmp258, cmp261, cmp265, cmp268, loadedv272, loadedv274, loadedv278, loadedv282, loadedv286, loadedv290, loadedv294, loadedv298, loadedv302, cmp306, cmp310, cmp314, cmp317, cmp320, cmp323, loadedv327, cmp331, cmp335, cmp339, cmp342, cmp345, cmp348, loadedv352, cmp356, cmp360, cmp363, cmp366, cmp369, loadedv373, cmp377, cmp381, cmp384, cmp387, cmp390, cmp394, cmp397, cmp400, cmp403, cmp406, loadedv410, cmp414, cmp417, cmp420, cmp423, loadedv427, loadedv431, cmp435, cmp439, cmp442, loadedv446, cmp450, cmp454, cmp457, cmp461, cmp464, loadedv468, cmp472, cmp475, cmp479, cmp482, loadedv486, cmp490, cmp493, loadedv497, loadedv501, loadedv505, loadedv509, loadedv513, cmp517, cmp520, loadedv524, v264 bool
	var retval unsafe.Pointer
	var v9, v13, v16, v79, v82, v91, v94 int16
	var state_addr, arrayidx, arrayidx11, arrayidx210, arrayidx217, arrayidx243, arrayidx250, result_symbol, result_symbol276, result_symbol280, result_symbol284, result_symbol288, result_symbol292, result_symbol296, result_symbol300, result_symbol304, result_symbol329, result_symbol354, result_symbol375, result_symbol412, result_symbol429, result_symbol433, result_symbol448, result_symbol470, result_symbol488, result_symbol499, result_symbol503, result_symbol507, result_symbol511, result_symbol515 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v24, v25, v26, v27, v28, v29, v30, v31, v33, v34, v35, v36, v37, v39, v40, v42, v43, v44, v46, v47, v49, v50, v51, v53, v54, v55, v57, v59, v61, v63, v65, v67, v69, v71, v73, v75, v77, v78, conv211, v80, v81, add215, v83, add220, v85, v86, v89, v90, conv244, v92, v93, add248, v95, add253, v96, v97, v98, v99, v100, v146, v147, v148, v149, v150, v151, v157, v158, v159, v160, v161, v162, v168, v169, v170, v171, v172, v178, v179, v180, v181, v182, v183, v184, v185, v186, v187, v193, v194, v195, v196, v207, v208, v209, v215, v216, v217, v218, v219, v225, v226, v227, v228, v234, v235, v261, v262 int32
	var lookahead, i, i203, i236, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10, conv205, idxprom209, idxprom216, conv238, idxprom242, idxprom249 int64
	var v3, storedv, v10, v23, v32, v38, v41, v45, v48, v52, v56, v58, v60, v62, v64, v66, v68, v70, v72, v74, v76, v84, v87, v88, v101, v106, v111, v116, v121, v126, v131, v136, v141, v152, v163, v173, v188, v197, v202, v210, v220, v229, v236, v241, v246, v251, v256, v263 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v102, v103, v104, v105, v107, v108, v109, v110, v112, v113, v114, v115, v117, v118, v119, v120, v122, v123, v124, v125, v127, v128, v129, v130, v132, v133, v134, v135, v137, v138, v139, v140, v142, v143, v144, v145, v153, v154, v155, v156, v164, v165, v166, v167, v174, v175, v176, v177, v189, v190, v191, v192, v198, v199, v200, v201, v203, v204, v205, v206, v211, v212, v213, v214, v221, v222, v223, v224, v230, v231, v232, v233, v237, v238, v239, v240, v242, v243, v244, v245, v247, v248, v249, v250, v252, v253, v254, v255, v257, v258, v259, v260 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end277, mark_end281, mark_end285, mark_end289, mark_end293, mark_end297, mark_end301, mark_end305, mark_end330, mark_end355, mark_end376, mark_end413, mark_end430, mark_end434, mark_end449, mark_end471, mark_end489, mark_end500, mark_end504, mark_end508, mark_end512, mark_end516 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i203, i236, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, loadedv29, v24, cmp31, v25, cmp35, v26, cmp39, v27, cmp43, v28, cmp47, v29, cmp50, v30, cmp53, v31, cmp57, v32, loadedv61, v33, cmp63, v34, cmp67, v35, cmp71, v36, cmp74, v37, cmp77, v38, loadedv81, v39, cmp83, v40, cmp87, v41, loadedv91, v42, cmp93, v43, cmp97, v44, cmp101, v45, loadedv105, v46, cmp107, v47, cmp111, v48, loadedv115, v49, cmp117, v50, cmp121, v51, cmp124, v52, loadedv128, v53, cmp130, v54, cmp134, v55, cmp137, v56, loadedv141, v57, cmp143, v58, loadedv147, v59, cmp149, v60, loadedv153, v61, cmp155, v62, loadedv159, v63, cmp161, v64, loadedv165, v65, cmp167, v66, loadedv171, v67, cmp173, v68, loadedv177, v69, cmp179, v70, loadedv183, v71, cmp185, v72, loadedv189, v73, cmp191, v74, loadedv195, v75, cmp197, v76, loadedv201, v77, conv205, cmp206, v78, idxprom209, arrayidx210, v79, conv211, v80, cmp212, v81, add215, idxprom216, arrayidx217, v82, v83, add220, v84, loadedv222, v85, cmp224, v86, cmp227, v87, loadedv231, v88, loadedv233, v89, conv238, cmp239, v90, idxprom242, arrayidx243, v91, conv244, v92, cmp245, v93, add248, idxprom249, arrayidx250, v94, v95, add253, v96, cmp255, v97, cmp258, v98, cmp261, v99, cmp265, v100, cmp268, v101, loadedv272, v102, result_symbol, v103, mark_end, v104, v105, v106, loadedv274, v107, result_symbol276, v108, mark_end277, v109, v110, v111, loadedv278, v112, result_symbol280, v113, mark_end281, v114, v115, v116, loadedv282, v117, result_symbol284, v118, mark_end285, v119, v120, v121, loadedv286, v122, result_symbol288, v123, mark_end289, v124, v125, v126, loadedv290, v127, result_symbol292, v128, mark_end293, v129, v130, v131, loadedv294, v132, result_symbol296, v133, mark_end297, v134, v135, v136, loadedv298, v137, result_symbol300, v138, mark_end301, v139, v140, v141, loadedv302, v142, result_symbol304, v143, mark_end305, v144, v145, v146, cmp306, v147, cmp310, v148, cmp314, v149, cmp317, v150, cmp320, v151, cmp323, v152, loadedv327, v153, result_symbol329, v154, mark_end330, v155, v156, v157, cmp331, v158, cmp335, v159, cmp339, v160, cmp342, v161, cmp345, v162, cmp348, v163, loadedv352, v164, result_symbol354, v165, mark_end355, v166, v167, v168, cmp356, v169, cmp360, v170, cmp363, v171, cmp366, v172, cmp369, v173, loadedv373, v174, result_symbol375, v175, mark_end376, v176, v177, v178, cmp377, v179, cmp381, v180, cmp384, v181, cmp387, v182, cmp390, v183, cmp394, v184, cmp397, v185, cmp400, v186, cmp403, v187, cmp406, v188, loadedv410, v189, result_symbol412, v190, mark_end413, v191, v192, v193, cmp414, v194, cmp417, v195, cmp420, v196, cmp423, v197, loadedv427, v198, result_symbol429, v199, mark_end430, v200, v201, v202, loadedv431, v203, result_symbol433, v204, mark_end434, v205, v206, v207, cmp435, v208, cmp439, v209, cmp442, v210, loadedv446, v211, result_symbol448, v212, mark_end449, v213, v214, v215, cmp450, v216, cmp454, v217, cmp457, v218, cmp461, v219, cmp464, v220, loadedv468, v221, result_symbol470, v222, mark_end471, v223, v224, v225, cmp472, v226, cmp475, v227, cmp479, v228, cmp482, v229, loadedv486, v230, result_symbol488, v231, mark_end489, v232, v233, v234, cmp490, v235, cmp493, v236, loadedv497, v237, result_symbol499, v238, mark_end500, v239, v240, v241, loadedv501, v242, result_symbol503, v243, mark_end504, v244, v245, v246, loadedv505, v247, result_symbol507, v248, mark_end508, v249, v250, v251, loadedv509, v252, result_symbol511, v253, mark_end512, v254, v255, v256, loadedv513, v257, result_symbol515, v258, mark_end516, v259, v260, v261, cmp517, v262, cmp520, v263, loadedv524, v264

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
	i203 = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i236 = libc.Ptr(&new(struct {
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
		goto sw_bb142
	case 9:
		goto sw_bb148
	case 10:
		goto sw_bb154
	case 11:
		goto sw_bb160
	case 12:
		goto sw_bb166
	case 13:
		goto sw_bb172
	case 14:
		goto sw_bb178
	case 15:
		goto sw_bb184
	case 16:
		goto sw_bb190
	case 17:
		goto sw_bb196
	case 18:
		goto sw_bb202
	case 19:
		goto sw_bb223
	case 20:
		goto sw_bb232
	case 21:
		goto sw_bb273
	case 22:
		goto sw_bb275
	case 23:
		goto sw_bb279
	case 24:
		goto sw_bb283
	case 25:
		goto sw_bb287
	case 26:
		goto sw_bb291
	case 27:
		goto sw_bb295
	case 28:
		goto sw_bb299
	case 29:
		goto sw_bb303
	case 30:
		goto sw_bb328
	case 31:
		goto sw_bb353
	case 32:
		goto sw_bb374
	case 33:
		goto sw_bb411
	case 34:
		goto sw_bb428
	case 35:
		goto sw_bb432
	case 36:
		goto sw_bb447
	case 37:
		goto sw_bb469
	case 38:
		goto sw_bb487
	case 39:
		goto sw_bb498
	case 40:
		goto sw_bb502
	case 41:
		goto sw_bb506
	case 42:
		goto sw_bb510
	case 43:
		goto sw_bb514
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
	cmp117 = v49 == 45
	if cmp117 {
		goto if_then119
	} else {
		goto if_end120
	}

if_then119:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end120:
	v50 = *libc.As[int32](lookahead)
	cmp121 = 48 <= v50
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
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end127:
	v52 = *libc.As[byte](result)
	loadedv128 = (v52 & 1) != 0
	*libc.As[bool](retval) = loadedv128
	goto _return

sw_bb129:
	v53 = *libc.As[int32](lookahead)
	cmp130 = v53 == 48
	if cmp130 {
		goto if_then132
	} else {
		goto if_end133
	}

if_then132:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end133:
	v54 = *libc.As[int32](lookahead)
	cmp134 = 49 <= v54
	if cmp134 {
		goto land_lhs_true136
	} else {
		goto if_end140
	}

land_lhs_true136:
	v55 = *libc.As[int32](lookahead)
	cmp137 = v55 <= 57
	if cmp137 {
		goto if_then139
	} else {
		goto if_end140
	}

if_then139:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end140:
	v56 = *libc.As[byte](result)
	loadedv141 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv141
	goto _return

sw_bb142:
	v57 = *libc.As[int32](lookahead)
	cmp143 = v57 == 97
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end146:
	v58 = *libc.As[byte](result)
	loadedv147 = (v58 & 1) != 0
	*libc.As[bool](retval) = loadedv147
	goto _return

sw_bb148:
	v59 = *libc.As[int32](lookahead)
	cmp149 = v59 == 101
	if cmp149 {
		goto if_then151
	} else {
		goto if_end152
	}

if_then151:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end152:
	v60 = *libc.As[byte](result)
	loadedv153 = (v60 & 1) != 0
	*libc.As[bool](retval) = loadedv153
	goto _return

sw_bb154:
	v61 = *libc.As[int32](lookahead)
	cmp155 = v61 == 101
	if cmp155 {
		goto if_then157
	} else {
		goto if_end158
	}

if_then157:
	*libc.As[int16](state_addr) = 40
	goto next_state

if_end158:
	v62 = *libc.As[byte](result)
	loadedv159 = (v62 & 1) != 0
	*libc.As[bool](retval) = loadedv159
	goto _return

sw_bb160:
	v63 = *libc.As[int32](lookahead)
	cmp161 = v63 == 108
	if cmp161 {
		goto if_then163
	} else {
		goto if_end164
	}

if_then163:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end164:
	v64 = *libc.As[byte](result)
	loadedv165 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv165
	goto _return

sw_bb166:
	v65 = *libc.As[int32](lookahead)
	cmp167 = v65 == 108
	if cmp167 {
		goto if_then169
	} else {
		goto if_end170
	}

if_then169:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end170:
	v66 = *libc.As[byte](result)
	loadedv171 = (v66 & 1) != 0
	*libc.As[bool](retval) = loadedv171
	goto _return

sw_bb172:
	v67 = *libc.As[int32](lookahead)
	cmp173 = v67 == 108
	if cmp173 {
		goto if_then175
	} else {
		goto if_end176
	}

if_then175:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end176:
	v68 = *libc.As[byte](result)
	loadedv177 = (v68 & 1) != 0
	*libc.As[bool](retval) = loadedv177
	goto _return

sw_bb178:
	v69 = *libc.As[int32](lookahead)
	cmp179 = v69 == 114
	if cmp179 {
		goto if_then181
	} else {
		goto if_end182
	}

if_then181:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end182:
	v70 = *libc.As[byte](result)
	loadedv183 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv183
	goto _return

sw_bb184:
	v71 = *libc.As[int32](lookahead)
	cmp185 = v71 == 115
	if cmp185 {
		goto if_then187
	} else {
		goto if_end188
	}

if_then187:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end188:
	v72 = *libc.As[byte](result)
	loadedv189 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv189
	goto _return

sw_bb190:
	v73 = *libc.As[int32](lookahead)
	cmp191 = v73 == 117
	if cmp191 {
		goto if_then193
	} else {
		goto if_end194
	}

if_then193:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end194:
	v74 = *libc.As[byte](result)
	loadedv195 = (v74 & 1) != 0
	*libc.As[bool](retval) = loadedv195
	goto _return

sw_bb196:
	v75 = *libc.As[int32](lookahead)
	cmp197 = v75 == 117
	if cmp197 {
		goto if_then199
	} else {
		goto if_end200
	}

if_then199:
	*libc.As[int16](state_addr) = 13
	goto next_state

if_end200:
	v76 = *libc.As[byte](result)
	loadedv201 = (v76 & 1) != 0
	*libc.As[bool](retval) = loadedv201
	goto _return

sw_bb202:
	*libc.As[int32](i203) = 0
	goto for_cond204

for_cond204:
	v77 = *libc.As[int32](i203)
	conv205 = int64(uint64(uint32(v77)))
	cmp206 = uint64(conv205) < uint64(18)
	if cmp206 {
		goto for_body208
	} else {
		goto for_end221
	}

for_body208:
	v78 = *libc.As[int32](i203)
	idxprom209 = int64(uint64(uint32(v78)))
	arrayidx210 = libc.Ptr(&ts_lex_map_30[idxprom209])
	v79 = *libc.As[int16](arrayidx210)
	conv211 = int32(uint32(uint16(v79)))
	v80 = *libc.As[int32](lookahead)
	cmp212 = conv211 == v80
	if cmp212 {
		goto if_then214
	} else {
		goto if_end218
	}

if_then214:
	v81 = *libc.As[int32](i203)
	add215 = v81 + 1
	idxprom216 = int64(uint64(uint32(add215)))
	arrayidx217 = libc.Ptr(&ts_lex_map_30[idxprom216])
	v82 = *libc.As[int16](arrayidx217)
	*libc.As[int16](state_addr) = v82
	goto next_state

if_end218:
	goto for_inc219

for_inc219:
	v83 = *libc.As[int32](i203)
	add220 = v83 + 2
	*libc.As[int32](i203) = add220
	goto for_cond204

for_end221:
	v84 = *libc.As[byte](result)
	loadedv222 = (v84 & 1) != 0
	*libc.As[bool](retval) = loadedv222
	goto _return

sw_bb223:
	v85 = *libc.As[int32](lookahead)
	cmp224 = 48 <= v85
	if cmp224 {
		goto land_lhs_true226
	} else {
		goto if_end230
	}

land_lhs_true226:
	v86 = *libc.As[int32](lookahead)
	cmp227 = v86 <= 57
	if cmp227 {
		goto if_then229
	} else {
		goto if_end230
	}

if_then229:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end230:
	v87 = *libc.As[byte](result)
	loadedv231 = (v87 & 1) != 0
	*libc.As[bool](retval) = loadedv231
	goto _return

sw_bb232:
	v88 = *libc.As[byte](eof)
	loadedv233 = (v88 & 1) != 0
	if loadedv233 {
		goto if_then234
	} else {
		goto if_end235
	}

if_then234:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end235:
	*libc.As[int32](i236) = 0
	goto for_cond237

for_cond237:
	v89 = *libc.As[int32](i236)
	conv238 = int64(uint64(uint32(v89)))
	cmp239 = uint64(conv238) < uint64(26)
	if cmp239 {
		goto for_body241
	} else {
		goto for_end254
	}

for_body241:
	v90 = *libc.As[int32](i236)
	idxprom242 = int64(uint64(uint32(v90)))
	arrayidx243 = libc.Ptr(&ts_lex_map_31[idxprom242])
	v91 = *libc.As[int16](arrayidx243)
	conv244 = int32(uint32(uint16(v91)))
	v92 = *libc.As[int32](lookahead)
	cmp245 = conv244 == v92
	if cmp245 {
		goto if_then247
	} else {
		goto if_end251
	}

if_then247:
	v93 = *libc.As[int32](i236)
	add248 = v93 + 1
	idxprom249 = int64(uint64(uint32(add248)))
	arrayidx250 = libc.Ptr(&ts_lex_map_31[idxprom249])
	v94 = *libc.As[int16](arrayidx250)
	*libc.As[int16](state_addr) = v94
	goto next_state

if_end251:
	goto for_inc252

for_inc252:
	v95 = *libc.As[int32](i236)
	add253 = v95 + 2
	*libc.As[int32](i236) = add253
	goto for_cond237

for_end254:
	v96 = *libc.As[int32](lookahead)
	cmp255 = 9 <= v96
	if cmp255 {
		goto land_lhs_true257
	} else {
		goto lor_lhs_false260
	}

land_lhs_true257:
	v97 = *libc.As[int32](lookahead)
	cmp258 = v97 <= 13
	if cmp258 {
		goto if_then263
	} else {
		goto lor_lhs_false260
	}

lor_lhs_false260:
	v98 = *libc.As[int32](lookahead)
	cmp261 = v98 == 32
	if cmp261 {
		goto if_then263
	} else {
		goto if_end264
	}

if_then263:
	*libc.As[byte](skip) = 1
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end264:
	v99 = *libc.As[int32](lookahead)
	cmp265 = 49 <= v99
	if cmp265 {
		goto land_lhs_true267
	} else {
		goto if_end271
	}

land_lhs_true267:
	v100 = *libc.As[int32](lookahead)
	cmp268 = v100 <= 57
	if cmp268 {
		goto if_then270
	} else {
		goto if_end271
	}

if_then270:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end271:
	v101 = *libc.As[byte](result)
	loadedv272 = (v101 & 1) != 0
	*libc.As[bool](retval) = loadedv272
	goto _return

sw_bb273:
	*libc.As[byte](result) = 1
	v102 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v102).F1)
	*libc.As[int16](result_symbol) = 0
	v103 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v103).F3)
	v104 = *libc.As[unsafe.Pointer](mark_end)
	v105 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v104)(v105)
	v106 = *libc.As[byte](result)
	loadedv274 = (v106 & 1) != 0
	*libc.As[bool](retval) = loadedv274
	goto _return

sw_bb275:
	*libc.As[byte](result) = 1
	v107 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol276 = libc.Ptr(&libc.As[TSLexer](v107).F1)
	*libc.As[int16](result_symbol276) = 1
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
	*libc.As[int16](result_symbol280) = 2
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
	v121 = *libc.As[byte](result)
	loadedv286 = (v121 & 1) != 0
	*libc.As[bool](retval) = loadedv286
	goto _return

sw_bb287:
	*libc.As[byte](result) = 1
	v122 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol288 = libc.Ptr(&libc.As[TSLexer](v122).F1)
	*libc.As[int16](result_symbol288) = 4
	v123 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end289 = libc.Ptr(&libc.As[TSLexer](v123).F3)
	v124 = *libc.As[unsafe.Pointer](mark_end289)
	v125 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v124)(v125)
	v126 = *libc.As[byte](result)
	loadedv290 = (v126 & 1) != 0
	*libc.As[bool](retval) = loadedv290
	goto _return

sw_bb291:
	*libc.As[byte](result) = 1
	v127 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol292 = libc.Ptr(&libc.As[TSLexer](v127).F1)
	*libc.As[int16](result_symbol292) = 5
	v128 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end293 = libc.Ptr(&libc.As[TSLexer](v128).F3)
	v129 = *libc.As[unsafe.Pointer](mark_end293)
	v130 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v129)(v130)
	v131 = *libc.As[byte](result)
	loadedv294 = (v131 & 1) != 0
	*libc.As[bool](retval) = loadedv294
	goto _return

sw_bb295:
	*libc.As[byte](result) = 1
	v132 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol296 = libc.Ptr(&libc.As[TSLexer](v132).F1)
	*libc.As[int16](result_symbol296) = 6
	v133 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end297 = libc.Ptr(&libc.As[TSLexer](v133).F3)
	v134 = *libc.As[unsafe.Pointer](mark_end297)
	v135 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v134)(v135)
	v136 = *libc.As[byte](result)
	loadedv298 = (v136 & 1) != 0
	*libc.As[bool](retval) = loadedv298
	goto _return

sw_bb299:
	*libc.As[byte](result) = 1
	v137 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol300 = libc.Ptr(&libc.As[TSLexer](v137).F1)
	*libc.As[int16](result_symbol300) = 7
	v138 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end301 = libc.Ptr(&libc.As[TSLexer](v138).F3)
	v139 = *libc.As[unsafe.Pointer](mark_end301)
	v140 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v139)(v140)
	v141 = *libc.As[byte](result)
	loadedv302 = (v141 & 1) != 0
	*libc.As[bool](retval) = loadedv302
	goto _return

sw_bb303:
	*libc.As[byte](result) = 1
	v142 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol304 = libc.Ptr(&libc.As[TSLexer](v142).F1)
	*libc.As[int16](result_symbol304) = 8
	v143 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end305 = libc.Ptr(&libc.As[TSLexer](v143).F3)
	v144 = *libc.As[unsafe.Pointer](mark_end305)
	v145 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v144)(v145)
	v146 = *libc.As[int32](lookahead)
	cmp306 = v146 == 42
	if cmp306 {
		goto if_then308
	} else {
		goto if_end309
	}

if_then308:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end309:
	v147 = *libc.As[int32](lookahead)
	cmp310 = v147 == 47
	if cmp310 {
		goto if_then312
	} else {
		goto if_end313
	}

if_then312:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end313:
	v148 = *libc.As[int32](lookahead)
	cmp314 = v148 != 0
	if cmp314 {
		goto land_lhs_true316
	} else {
		goto if_end326
	}

land_lhs_true316:
	v149 = *libc.As[int32](lookahead)
	cmp317 = v149 != 10
	if cmp317 {
		goto land_lhs_true319
	} else {
		goto if_end326
	}

land_lhs_true319:
	v150 = *libc.As[int32](lookahead)
	cmp320 = v150 != 34
	if cmp320 {
		goto land_lhs_true322
	} else {
		goto if_end326
	}

land_lhs_true322:
	v151 = *libc.As[int32](lookahead)
	cmp323 = v151 != 92
	if cmp323 {
		goto if_then325
	} else {
		goto if_end326
	}

if_then325:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end326:
	v152 = *libc.As[byte](result)
	loadedv327 = (v152 & 1) != 0
	*libc.As[bool](retval) = loadedv327
	goto _return

sw_bb328:
	*libc.As[byte](result) = 1
	v153 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol329 = libc.Ptr(&libc.As[TSLexer](v153).F1)
	*libc.As[int16](result_symbol329) = 8
	v154 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end330 = libc.Ptr(&libc.As[TSLexer](v154).F3)
	v155 = *libc.As[unsafe.Pointer](mark_end330)
	v156 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v155)(v156)
	v157 = *libc.As[int32](lookahead)
	cmp331 = v157 == 42
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end334:
	v158 = *libc.As[int32](lookahead)
	cmp335 = v158 == 47
	if cmp335 {
		goto if_then337
	} else {
		goto if_end338
	}

if_then337:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end338:
	v159 = *libc.As[int32](lookahead)
	cmp339 = v159 != 0
	if cmp339 {
		goto land_lhs_true341
	} else {
		goto if_end351
	}

land_lhs_true341:
	v160 = *libc.As[int32](lookahead)
	cmp342 = v160 != 10
	if cmp342 {
		goto land_lhs_true344
	} else {
		goto if_end351
	}

land_lhs_true344:
	v161 = *libc.As[int32](lookahead)
	cmp345 = v161 != 34
	if cmp345 {
		goto land_lhs_true347
	} else {
		goto if_end351
	}

land_lhs_true347:
	v162 = *libc.As[int32](lookahead)
	cmp348 = v162 != 92
	if cmp348 {
		goto if_then350
	} else {
		goto if_end351
	}

if_then350:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end351:
	v163 = *libc.As[byte](result)
	loadedv352 = (v163 & 1) != 0
	*libc.As[bool](retval) = loadedv352
	goto _return

sw_bb353:
	*libc.As[byte](result) = 1
	v164 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol354 = libc.Ptr(&libc.As[TSLexer](v164).F1)
	*libc.As[int16](result_symbol354) = 8
	v165 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end355 = libc.Ptr(&libc.As[TSLexer](v165).F3)
	v166 = *libc.As[unsafe.Pointer](mark_end355)
	v167 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v166)(v167)
	v168 = *libc.As[int32](lookahead)
	cmp356 = v168 == 42
	if cmp356 {
		goto if_then358
	} else {
		goto if_end359
	}

if_then358:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end359:
	v169 = *libc.As[int32](lookahead)
	cmp360 = v169 != 0
	if cmp360 {
		goto land_lhs_true362
	} else {
		goto if_end372
	}

land_lhs_true362:
	v170 = *libc.As[int32](lookahead)
	cmp363 = v170 != 10
	if cmp363 {
		goto land_lhs_true365
	} else {
		goto if_end372
	}

land_lhs_true365:
	v171 = *libc.As[int32](lookahead)
	cmp366 = v171 != 34
	if cmp366 {
		goto land_lhs_true368
	} else {
		goto if_end372
	}

land_lhs_true368:
	v172 = *libc.As[int32](lookahead)
	cmp369 = v172 != 92
	if cmp369 {
		goto if_then371
	} else {
		goto if_end372
	}

if_then371:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end372:
	v173 = *libc.As[byte](result)
	loadedv373 = (v173 & 1) != 0
	*libc.As[bool](retval) = loadedv373
	goto _return

sw_bb374:
	*libc.As[byte](result) = 1
	v174 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol375 = libc.Ptr(&libc.As[TSLexer](v174).F1)
	*libc.As[int16](result_symbol375) = 8
	v175 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end376 = libc.Ptr(&libc.As[TSLexer](v175).F3)
	v176 = *libc.As[unsafe.Pointer](mark_end376)
	v177 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v176)(v177)
	v178 = *libc.As[int32](lookahead)
	cmp377 = v178 == 47
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*libc.As[int16](state_addr) = 29
	goto next_state

if_end380:
	v179 = *libc.As[int32](lookahead)
	cmp381 = v179 == 9
	if cmp381 {
		goto if_then392
	} else {
		goto lor_lhs_false383
	}

lor_lhs_false383:
	v180 = *libc.As[int32](lookahead)
	cmp384 = 11 <= v180
	if cmp384 {
		goto land_lhs_true386
	} else {
		goto lor_lhs_false389
	}

land_lhs_true386:
	v181 = *libc.As[int32](lookahead)
	cmp387 = v181 <= 13
	if cmp387 {
		goto if_then392
	} else {
		goto lor_lhs_false389
	}

lor_lhs_false389:
	v182 = *libc.As[int32](lookahead)
	cmp390 = v182 == 32
	if cmp390 {
		goto if_then392
	} else {
		goto if_end393
	}

if_then392:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end393:
	v183 = *libc.As[int32](lookahead)
	cmp394 = v183 != 0
	if cmp394 {
		goto land_lhs_true396
	} else {
		goto if_end409
	}

land_lhs_true396:
	v184 = *libc.As[int32](lookahead)
	cmp397 = v184 < 9
	if cmp397 {
		goto land_lhs_true402
	} else {
		goto lor_lhs_false399
	}

lor_lhs_false399:
	v185 = *libc.As[int32](lookahead)
	cmp400 = 13 < v185
	if cmp400 {
		goto land_lhs_true402
	} else {
		goto if_end409
	}

land_lhs_true402:
	v186 = *libc.As[int32](lookahead)
	cmp403 = v186 != 34
	if cmp403 {
		goto land_lhs_true405
	} else {
		goto if_end409
	}

land_lhs_true405:
	v187 = *libc.As[int32](lookahead)
	cmp406 = v187 != 92
	if cmp406 {
		goto if_then408
	} else {
		goto if_end409
	}

if_then408:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end409:
	v188 = *libc.As[byte](result)
	loadedv410 = (v188 & 1) != 0
	*libc.As[bool](retval) = loadedv410
	goto _return

sw_bb411:
	*libc.As[byte](result) = 1
	v189 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol412 = libc.Ptr(&libc.As[TSLexer](v189).F1)
	*libc.As[int16](result_symbol412) = 8
	v190 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end413 = libc.Ptr(&libc.As[TSLexer](v190).F3)
	v191 = *libc.As[unsafe.Pointer](mark_end413)
	v192 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v191)(v192)
	v193 = *libc.As[int32](lookahead)
	cmp414 = v193 != 0
	if cmp414 {
		goto land_lhs_true416
	} else {
		goto if_end426
	}

land_lhs_true416:
	v194 = *libc.As[int32](lookahead)
	cmp417 = v194 != 10
	if cmp417 {
		goto land_lhs_true419
	} else {
		goto if_end426
	}

land_lhs_true419:
	v195 = *libc.As[int32](lookahead)
	cmp420 = v195 != 34
	if cmp420 {
		goto land_lhs_true422
	} else {
		goto if_end426
	}

land_lhs_true422:
	v196 = *libc.As[int32](lookahead)
	cmp423 = v196 != 92
	if cmp423 {
		goto if_then425
	} else {
		goto if_end426
	}

if_then425:
	*libc.As[int16](state_addr) = 33
	goto next_state

if_end426:
	v197 = *libc.As[byte](result)
	loadedv427 = (v197 & 1) != 0
	*libc.As[bool](retval) = loadedv427
	goto _return

sw_bb428:
	*libc.As[byte](result) = 1
	v198 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol429 = libc.Ptr(&libc.As[TSLexer](v198).F1)
	*libc.As[int16](result_symbol429) = 9
	v199 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end430 = libc.Ptr(&libc.As[TSLexer](v199).F3)
	v200 = *libc.As[unsafe.Pointer](mark_end430)
	v201 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v200)(v201)
	v202 = *libc.As[byte](result)
	loadedv431 = (v202 & 1) != 0
	*libc.As[bool](retval) = loadedv431
	goto _return

sw_bb432:
	*libc.As[byte](result) = 1
	v203 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol433 = libc.Ptr(&libc.As[TSLexer](v203).F1)
	*libc.As[int16](result_symbol433) = 10
	v204 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end434 = libc.Ptr(&libc.As[TSLexer](v204).F3)
	v205 = *libc.As[unsafe.Pointer](mark_end434)
	v206 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v205)(v206)
	v207 = *libc.As[int32](lookahead)
	cmp435 = v207 == 46
	if cmp435 {
		goto if_then437
	} else {
		goto if_end438
	}

if_then437:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end438:
	v208 = *libc.As[int32](lookahead)
	cmp439 = v208 == 69
	if cmp439 {
		goto if_then444
	} else {
		goto lor_lhs_false441
	}

lor_lhs_false441:
	v209 = *libc.As[int32](lookahead)
	cmp442 = v209 == 101
	if cmp442 {
		goto if_then444
	} else {
		goto if_end445
	}

if_then444:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end445:
	v210 = *libc.As[byte](result)
	loadedv446 = (v210 & 1) != 0
	*libc.As[bool](retval) = loadedv446
	goto _return

sw_bb447:
	*libc.As[byte](result) = 1
	v211 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol448 = libc.Ptr(&libc.As[TSLexer](v211).F1)
	*libc.As[int16](result_symbol448) = 10
	v212 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end449 = libc.Ptr(&libc.As[TSLexer](v212).F3)
	v213 = *libc.As[unsafe.Pointer](mark_end449)
	v214 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v213)(v214)
	v215 = *libc.As[int32](lookahead)
	cmp450 = v215 == 46
	if cmp450 {
		goto if_then452
	} else {
		goto if_end453
	}

if_then452:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end453:
	v216 = *libc.As[int32](lookahead)
	cmp454 = v216 == 69
	if cmp454 {
		goto if_then459
	} else {
		goto lor_lhs_false456
	}

lor_lhs_false456:
	v217 = *libc.As[int32](lookahead)
	cmp457 = v217 == 101
	if cmp457 {
		goto if_then459
	} else {
		goto if_end460
	}

if_then459:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end460:
	v218 = *libc.As[int32](lookahead)
	cmp461 = 48 <= v218
	if cmp461 {
		goto land_lhs_true463
	} else {
		goto if_end467
	}

land_lhs_true463:
	v219 = *libc.As[int32](lookahead)
	cmp464 = v219 <= 57
	if cmp464 {
		goto if_then466
	} else {
		goto if_end467
	}

if_then466:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end467:
	v220 = *libc.As[byte](result)
	loadedv468 = (v220 & 1) != 0
	*libc.As[bool](retval) = loadedv468
	goto _return

sw_bb469:
	*libc.As[byte](result) = 1
	v221 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol470 = libc.Ptr(&libc.As[TSLexer](v221).F1)
	*libc.As[int16](result_symbol470) = 10
	v222 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end471 = libc.Ptr(&libc.As[TSLexer](v222).F3)
	v223 = *libc.As[unsafe.Pointer](mark_end471)
	v224 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v223)(v224)
	v225 = *libc.As[int32](lookahead)
	cmp472 = v225 == 69
	if cmp472 {
		goto if_then477
	} else {
		goto lor_lhs_false474
	}

lor_lhs_false474:
	v226 = *libc.As[int32](lookahead)
	cmp475 = v226 == 101
	if cmp475 {
		goto if_then477
	} else {
		goto if_end478
	}

if_then477:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end478:
	v227 = *libc.As[int32](lookahead)
	cmp479 = 48 <= v227
	if cmp479 {
		goto land_lhs_true481
	} else {
		goto if_end485
	}

land_lhs_true481:
	v228 = *libc.As[int32](lookahead)
	cmp482 = v228 <= 57
	if cmp482 {
		goto if_then484
	} else {
		goto if_end485
	}

if_then484:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end485:
	v229 = *libc.As[byte](result)
	loadedv486 = (v229 & 1) != 0
	*libc.As[bool](retval) = loadedv486
	goto _return

sw_bb487:
	*libc.As[byte](result) = 1
	v230 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol488 = libc.Ptr(&libc.As[TSLexer](v230).F1)
	*libc.As[int16](result_symbol488) = 10
	v231 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end489 = libc.Ptr(&libc.As[TSLexer](v231).F3)
	v232 = *libc.As[unsafe.Pointer](mark_end489)
	v233 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v232)(v233)
	v234 = *libc.As[int32](lookahead)
	cmp490 = 48 <= v234
	if cmp490 {
		goto land_lhs_true492
	} else {
		goto if_end496
	}

land_lhs_true492:
	v235 = *libc.As[int32](lookahead)
	cmp493 = v235 <= 57
	if cmp493 {
		goto if_then495
	} else {
		goto if_end496
	}

if_then495:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end496:
	v236 = *libc.As[byte](result)
	loadedv497 = (v236 & 1) != 0
	*libc.As[bool](retval) = loadedv497
	goto _return

sw_bb498:
	*libc.As[byte](result) = 1
	v237 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol499 = libc.Ptr(&libc.As[TSLexer](v237).F1)
	*libc.As[int16](result_symbol499) = 11
	v238 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end500 = libc.Ptr(&libc.As[TSLexer](v238).F3)
	v239 = *libc.As[unsafe.Pointer](mark_end500)
	v240 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v239)(v240)
	v241 = *libc.As[byte](result)
	loadedv501 = (v241 & 1) != 0
	*libc.As[bool](retval) = loadedv501
	goto _return

sw_bb502:
	*libc.As[byte](result) = 1
	v242 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol503 = libc.Ptr(&libc.As[TSLexer](v242).F1)
	*libc.As[int16](result_symbol503) = 12
	v243 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end504 = libc.Ptr(&libc.As[TSLexer](v243).F3)
	v244 = *libc.As[unsafe.Pointer](mark_end504)
	v245 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v244)(v245)
	v246 = *libc.As[byte](result)
	loadedv505 = (v246 & 1) != 0
	*libc.As[bool](retval) = loadedv505
	goto _return

sw_bb506:
	*libc.As[byte](result) = 1
	v247 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol507 = libc.Ptr(&libc.As[TSLexer](v247).F1)
	*libc.As[int16](result_symbol507) = 13
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end508 = libc.Ptr(&libc.As[TSLexer](v248).F3)
	v249 = *libc.As[unsafe.Pointer](mark_end508)
	v250 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v249)(v250)
	v251 = *libc.As[byte](result)
	loadedv509 = (v251 & 1) != 0
	*libc.As[bool](retval) = loadedv509
	goto _return

sw_bb510:
	*libc.As[byte](result) = 1
	v252 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol511 = libc.Ptr(&libc.As[TSLexer](v252).F1)
	*libc.As[int16](result_symbol511) = 14
	v253 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end512 = libc.Ptr(&libc.As[TSLexer](v253).F3)
	v254 = *libc.As[unsafe.Pointer](mark_end512)
	v255 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v254)(v255)
	v256 = *libc.As[byte](result)
	loadedv513 = (v256 & 1) != 0
	*libc.As[bool](retval) = loadedv513
	goto _return

sw_bb514:
	*libc.As[byte](result) = 1
	v257 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol515 = libc.Ptr(&libc.As[TSLexer](v257).F1)
	*libc.As[int16](result_symbol515) = 14
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end516 = libc.Ptr(&libc.As[TSLexer](v258).F3)
	v259 = *libc.As[unsafe.Pointer](mark_end516)
	v260 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v259)(v260)
	v261 = *libc.As[int32](lookahead)
	cmp517 = v261 != 0
	if cmp517 {
		goto land_lhs_true519
	} else {
		goto if_end523
	}

land_lhs_true519:
	v262 = *libc.As[int32](lookahead)
	cmp520 = v262 != 10
	if cmp520 {
		goto if_then522
	} else {
		goto if_end523
	}

if_then522:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end523:
	v263 = *libc.As[byte](result)
	loadedv524 = (v263 & 1) != 0
	*libc.As[bool](retval) = loadedv524
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v264 = *libc.As[bool](retval)
	return v264
}
