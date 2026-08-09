package grammar_test

import (
	"unsafe"
	"github.com/lewtec/leaven/libc"
)

type Scanner struct {
	F0 int32
	F1 int32
	F2 byte
}

type TSFieldMapEntry struct {
	F0 int16
	F1 byte
	F2 byte
}

type TSLanguage struct {
	F0 int32
	F1 int32
	F2 int32
	F3 int32
	F4 int32
	F5 int32
	F6 int32
	F7 int32
	F8 int32
	F9 int16
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

type TSLexMode struct {
	F0 int16
	F1 int16
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

var tree_sitter_test_language TSLanguage = TSLanguage{14, 26, 3, 16, 3, 33, 2, 9, 3, 5, &(*[2][26]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[110]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], &ts_field_names[0], &ts_field_map_slices[0], &ts_field_map_entries[0], &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], (*TSLexerMode)(unsafe.Pointer(&ts_lex_modes)), ts_lex, nil, 0, anon_2{&ts_external_scanner_states[0][0], &ts_external_scanner_symbol_map[0], tree_sitter_test_external_scanner_create, tree_sitter_test_external_scanner_destroy, tree_sitter_test_external_scanner_scan, tree_sitter_test_external_scanner_serialize, tree_sitter_test_external_scanner_deserialize}, &ts_primary_state_ids[0], nil, nil, 0, 0, nil, nil, nil, TSLanguageMetadata{}}

var ts_small_parse_table [325]int16 = [325]int16{
	10, 9, 1, 2, 13, 1, 6, 15, 1, 9, 17, 1, 12, 19, 1, 14,
	7, 1, 23, 9, 1, 22, 26, 1, 20, 3, 2, 21, 25, 11, 3, 3,
	4, 5, 9, 9, 1, 2, 13, 1, 6, 15, 1, 9, 21, 1, 12, 23,
	1, 14, 7, 1, 23, 9, 1, 22, 4, 2, 21, 25, 11, 3, 3, 4,
	5, 9, 25, 1, 2, 31, 1, 6, 34, 1, 9, 37, 1, 12, 40, 1,
	14, 7, 1, 23, 9, 1, 22, 4, 2, 21, 25, 28, 3, 3, 4, 5,
	1, 42, 8, 14, 2, 3, 4, 5, 6, 9, 12, 1, 44, 8, 14, 2,
	3, 4, 5, 6, 9, 12, 1, 46, 8, 14, 2, 3, 4, 5, 6, 9,
	12, 1, 48, 8, 14, 2, 3, 4, 5, 6, 9, 12, 1, 50, 8, 14,
	2, 3, 4, 5, 6, 9, 12, 1, 52, 8, 14, 2, 3, 4, 5, 6,
	9, 12, 5, 7, 1, 13, 54, 1, 0, 56, 1, 12, 19, 1, 19, 12,
	2, 17, 24, 5, 58, 1, 0, 60, 1, 12, 63, 1, 13, 19, 1, 19,
	12, 2, 17, 24, 3, 14, 1, 18, 66, 2, 13, 0, 68, 2, 1, 12,
	3, 15, 1, 18, 68, 2, 1, 12, 70, 2, 13, 0, 3, 15, 1, 18,
	72, 2, 13, 0, 74, 2, 1, 12, 3, 79, 1, 15, 17, 1, 18, 77,
	2, 1, 12, 3, 72, 1, 15, 17, 1, 18, 81, 2, 1, 12, 1, 84,
	4, 13, 0, 1, 12, 2, 16, 1, 18, 77, 2, 1, 12, 1, 84, 3,
	15, 1, 12, 1, 86, 2, 1, 12, 1, 88, 2, 1, 12, 1, 90, 1,
	8, 1, 92, 1, 11, 1, 94, 1, 7, 1, 96, 1, 14, 1, 98, 1,
	8, 1, 100, 1, 1, 1, 102, 1, 0, 1, 104, 1, 12, 1, 106, 1,
	7, 1, 108, 1, 10,
}

var ts_small_parse_table_map [31]int32 = [31]int32{
	0, 34, 65, 96, 107, 118, 129, 140, 151, 162, 179, 196, 208, 220, 232, 243,
	254, 261, 269, 275, 280, 285, 289, 293, 297, 301, 305, 309, 313, 317, 321,
}

var ts_symbol_names [29]*byte = [29]*byte{
	&_str[0], &_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_14[0], &_str_14[0],
	&_str_15[0], &_str_16[0], &_str_17[0], &_str_18[0], &_str_19[0], &_str_20[0], &_str_21[0], &_str_22[0], &_str_23[0], &_str_24[0], &_str_25[0], &_str_26[0], &_str_27[0],
}

var ts_field_names [4]*byte = [4]*byte{nil, &_str_28[0], &_str_29[0], &_str_30[0]}

var ts_field_map_slices [9]TSMapSlice = [9]TSMapSlice{TSMapSlice{}, TSMapSlice{}, TSMapSlice{}, TSMapSlice{0, 1}, TSMapSlice{1, 1}, TSMapSlice{2, 1}, TSMapSlice{}, TSMapSlice{3, 1}, TSMapSlice{4, 1}}

var ts_field_map_entries [5]TSFieldMapEntry = [5]TSFieldMapEntry{TSFieldMapEntry{1, 0, 0}, TSFieldMapEntry{2, 0, 1}, TSFieldMapEntry{3, 0, 1}, TSFieldMapEntry{2, 2, 0}, TSFieldMapEntry{3, 2, 0}}

var ts_symbol_metadata [29]TSSymbolMetadata = [29]TSSymbolMetadata{
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0},
	TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0},
}

var ts_symbol_map [29]int16 = [29]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 12, 13, 13, 13,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
}

var ts_non_terminal_alias_map [6]int16 = [6]int16{18, 3, 18, 26, 28, 0}

var ts_alias_sequences [9][5]int16 = [9][5]int16{[5]int16{}, [5]int16{0, 26, 0, 0, 0}, [5]int16{0, 27, 0, 0, 0}, [5]int16{}, [5]int16{}, [5]int16{}, [5]int16{0, 26, 0, 28, 0}, [5]int16{}, [5]int16{}}

var ts_lex_modes [33]TSLexMode = [33]TSLexMode{
	TSLexMode{0, 1}, TSLexMode{0, 2}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 3}, TSLexMode{0, 2}, TSLexMode{0, 2}, TSLexMode{77, 2}, TSLexMode{77, 2}, TSLexMode{77, 2},
	TSLexMode{77, 4}, TSLexMode{77, 4}, TSLexMode{77, 2}, TSLexMode{77, 0}, TSLexMode{77, 4}, TSLexMode{77, 0}, TSLexMode{77, 0}, TSLexMode{}, TSLexMode{6, 0}, TSLexMode{}, TSLexMode{0, 3}, TSLexMode{}, TSLexMode{77, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{},
	TSLexMode{76, 0},
}

var ts_external_scanner_states [5][3]byte = [5][3]byte{[3]byte{}, [3]byte{1, 1, 1}, [3]byte{1, 0, 0}, [3]byte{0, 1, 0}, [3]byte{0, 0, 1}}

var ts_external_scanner_symbol_map [3]int16 = [3]int16{13, 14, 15}

var ts_primary_state_ids [33]int16 = [33]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 15, 18, 19, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32,
}

var ts_parse_table struct {
	F0 struct {
	F0 [16]int16
	F1 [10]int16
}
	F1 [26]int16
} = struct {
	F0 struct {
	F0 [16]int16
	F1 [10]int16
}
	F1 [26]int16
}{struct {
	F0 [16]int16
	F1 [10]int16
}{[16]int16{
	1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}, [10]int16{}}, [26]int16{
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 7, 0, 0,
	29, 11, 0, 19, 0, 0, 0, 0, 11, 0,
}}

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
	F4 TSParseActionEntry
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
	F24 TSParseActionEntry
	F25 struct {
	F0 anon_1
	F1 [6]byte
}
	F26 TSParseActionEntry
	F27 struct {
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
	F28 struct {
	F0 anon_1
	F1 [6]byte
}
	F29 TSParseActionEntry
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
	F34 struct {
	F0 anon_1
	F1 [6]byte
}
	F35 TSParseActionEntry
	F36 struct {
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
	F37 struct {
	F0 anon_1
	F1 [6]byte
}
	F38 TSParseActionEntry
	F39 struct {
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
	F40 struct {
	F0 anon_1
	F1 [6]byte
}
	F41 TSParseActionEntry
	F42 struct {
	F0 anon_1
	F1 [6]byte
}
	F43 TSParseActionEntry
	F44 struct {
	F0 anon_1
	F1 [6]byte
}
	F45 TSParseActionEntry
	F46 struct {
	F0 anon_1
	F1 [6]byte
}
	F47 TSParseActionEntry
	F48 struct {
	F0 anon_1
	F1 [6]byte
}
	F49 TSParseActionEntry
	F50 struct {
	F0 anon_1
	F1 [6]byte
}
	F51 TSParseActionEntry
	F52 struct {
	F0 anon_1
	F1 [6]byte
}
	F53 TSParseActionEntry
	F54 struct {
	F0 anon_1
	F1 [6]byte
}
	F55 TSParseActionEntry
	F56 struct {
	F0 anon_1
	F1 [6]byte
}
	F57 struct {
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
	F58 struct {
	F0 anon_1
	F1 [6]byte
}
	F59 TSParseActionEntry
	F60 struct {
	F0 anon_1
	F1 [6]byte
}
	F61 TSParseActionEntry
	F62 struct {
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
	F63 struct {
	F0 anon_1
	F1 [6]byte
}
	F64 TSParseActionEntry
	F65 struct {
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
	F66 struct {
	F0 anon_1
	F1 [6]byte
}
	F67 TSParseActionEntry
	F68 struct {
	F0 anon_1
	F1 [6]byte
}
	F69 struct {
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
	F70 struct {
	F0 anon_1
	F1 [6]byte
}
	F71 TSParseActionEntry
	F72 struct {
	F0 anon_1
	F1 [6]byte
}
	F73 TSParseActionEntry
	F74 struct {
	F0 anon_1
	F1 [6]byte
}
	F75 TSParseActionEntry
	F76 struct {
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
	F77 struct {
	F0 anon_1
	F1 [6]byte
}
	F78 struct {
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
	F79 struct {
	F0 anon_1
	F1 [6]byte
}
	F80 struct {
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
	F81 struct {
	F0 anon_1
	F1 [6]byte
}
	F82 TSParseActionEntry
	F83 struct {
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
	F84 struct {
	F0 anon_1
	F1 [6]byte
}
	F85 TSParseActionEntry
	F86 struct {
	F0 anon_1
	F1 [6]byte
}
	F87 TSParseActionEntry
	F88 struct {
	F0 anon_1
	F1 [6]byte
}
	F89 TSParseActionEntry
	F90 struct {
	F0 anon_1
	F1 [6]byte
}
	F91 struct {
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
	F92 struct {
	F0 anon_1
	F1 [6]byte
}
	F93 struct {
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
	F94 struct {
	F0 anon_1
	F1 [6]byte
}
	F95 struct {
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
	F96 struct {
	F0 anon_1
	F1 [6]byte
}
	F97 struct {
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
	F98 struct {
	F0 anon_1
	F1 [6]byte
}
	F99 struct {
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
	F100 struct {
	F0 anon_1
	F1 [6]byte
}
	F101 struct {
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
	F102 struct {
	F0 anon_1
	F1 [6]byte
}
	F103 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F104 struct {
	F0 anon_1
	F1 [6]byte
}
	F105 struct {
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
	F106 struct {
	F0 anon_1
	F1 [6]byte
}
	F107 struct {
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
	F108 struct {
	F0 anon_1
	F1 [6]byte
}
	F109 struct {
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
	F4 TSParseActionEntry
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
	F24 TSParseActionEntry
	F25 struct {
	F0 anon_1
	F1 [6]byte
}
	F26 TSParseActionEntry
	F27 struct {
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
	F28 struct {
	F0 anon_1
	F1 [6]byte
}
	F29 TSParseActionEntry
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
	F34 struct {
	F0 anon_1
	F1 [6]byte
}
	F35 TSParseActionEntry
	F36 struct {
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
	F37 struct {
	F0 anon_1
	F1 [6]byte
}
	F38 TSParseActionEntry
	F39 struct {
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
	F40 struct {
	F0 anon_1
	F1 [6]byte
}
	F41 TSParseActionEntry
	F42 struct {
	F0 anon_1
	F1 [6]byte
}
	F43 TSParseActionEntry
	F44 struct {
	F0 anon_1
	F1 [6]byte
}
	F45 TSParseActionEntry
	F46 struct {
	F0 anon_1
	F1 [6]byte
}
	F47 TSParseActionEntry
	F48 struct {
	F0 anon_1
	F1 [6]byte
}
	F49 TSParseActionEntry
	F50 struct {
	F0 anon_1
	F1 [6]byte
}
	F51 TSParseActionEntry
	F52 struct {
	F0 anon_1
	F1 [6]byte
}
	F53 TSParseActionEntry
	F54 struct {
	F0 anon_1
	F1 [6]byte
}
	F55 TSParseActionEntry
	F56 struct {
	F0 anon_1
	F1 [6]byte
}
	F57 struct {
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
	F58 struct {
	F0 anon_1
	F1 [6]byte
}
	F59 TSParseActionEntry
	F60 struct {
	F0 anon_1
	F1 [6]byte
}
	F61 TSParseActionEntry
	F62 struct {
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
	F63 struct {
	F0 anon_1
	F1 [6]byte
}
	F64 TSParseActionEntry
	F65 struct {
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
	F66 struct {
	F0 anon_1
	F1 [6]byte
}
	F67 TSParseActionEntry
	F68 struct {
	F0 anon_1
	F1 [6]byte
}
	F69 struct {
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
	F70 struct {
	F0 anon_1
	F1 [6]byte
}
	F71 TSParseActionEntry
	F72 struct {
	F0 anon_1
	F1 [6]byte
}
	F73 TSParseActionEntry
	F74 struct {
	F0 anon_1
	F1 [6]byte
}
	F75 TSParseActionEntry
	F76 struct {
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
	F77 struct {
	F0 anon_1
	F1 [6]byte
}
	F78 struct {
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
	F79 struct {
	F0 anon_1
	F1 [6]byte
}
	F80 struct {
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
	F81 struct {
	F0 anon_1
	F1 [6]byte
}
	F82 TSParseActionEntry
	F83 struct {
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
	F84 struct {
	F0 anon_1
	F1 [6]byte
}
	F85 TSParseActionEntry
	F86 struct {
	F0 anon_1
	F1 [6]byte
}
	F87 TSParseActionEntry
	F88 struct {
	F0 anon_1
	F1 [6]byte
}
	F89 TSParseActionEntry
	F90 struct {
	F0 anon_1
	F1 [6]byte
}
	F91 struct {
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
	F92 struct {
	F0 anon_1
	F1 [6]byte
}
	F93 struct {
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
	F94 struct {
	F0 anon_1
	F1 [6]byte
}
	F95 struct {
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
	F96 struct {
	F0 anon_1
	F1 [6]byte
}
	F97 struct {
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
	F98 struct {
	F0 anon_1
	F1 [6]byte
}
	F99 struct {
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
	F100 struct {
	F0 anon_1
	F1 [6]byte
}
	F101 struct {
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
	F102 struct {
	F0 anon_1
	F1 [6]byte
}
	F103 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F104 struct {
	F0 anon_1
	F1 [6]byte
}
	F105 struct {
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
	F106 struct {
	F0 anon_1
	F1 [6]byte
}
	F107 struct {
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
	F108 struct {
	F0 anon_1
	F1 [6]byte
}
	F109 struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 16, 0, 0}}}, struct {
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
}{0, 28, 0, 0}, [2]byte{}}}, struct {
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
}{0, 6, 0, 0}, [2]byte{}}}, struct {
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
}{0, 31, 0, 0}, [2]byte{}}}, struct {
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
}{0, 21, 0, 0}, [2]byte{}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 20, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 25, 0, 0}}}, struct {
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
}{0, 5, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 25, 0, 0}}}, struct {
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
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 25, 0, 0}}}, struct {
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
}{0, 31, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 25, 0, 0}}}, struct {
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
}{0, 25, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 25, 0, 0}}}, struct {
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
}{0, 4, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 25, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 3}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 5}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 22, 0, 7}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 4}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 23, 0, 8}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 16, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 24, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 24, 0, 0}}}, struct {
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
}{0, 12, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 24, 0, 0}}}, struct {
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
}{0, 28, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 17, 0, 1}}}, struct {
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
}{0, 18, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 17, 0, 6}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 18, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 18, 0, 0}}}, struct {
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
}{0, 18, 0, 1}, [2]byte{}}}, struct {
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
}{0, 20, 0, 0}, [2]byte{}}}, struct {
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
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 18, 0, 0}}}, struct {
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
}{0, 20, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 18, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 19, 0, 2}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 19, 0, 2}}}, struct {
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
}{0, 23, 0, 0}, [2]byte{}}}, struct {
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
}{0, 24, 0, 0}, [2]byte{}}}, struct {
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
}{0, 22, 0, 0}, [2]byte{}}}, struct {
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
}{0, 30, 0, 0}, [2]byte{}}}, struct {
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
}{0, 32, 0, 0}, [2]byte{}}}, struct {
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
}{0, 27, 0, 0}, [2]byte{}}}}

var _str [4]byte = [4]byte{101, 110, 100, 0}

var _str_3 [6]byte = [6]byte{95, 108, 105, 110, 101, 0}

var _str_4 [5]byte = [5]byte{58, 99, 115, 116, 0}

var _str_5 [7]byte = [7]byte{58, 101, 114, 114, 111, 114, 0}

var _str_6 [11]byte = [11]byte{58, 102, 97, 105, 108, 45, 102, 97, 115, 116, 0}

var _str_7 [6]byte = [6]byte{58, 115, 107, 105, 112, 0}

var _str_8 [10]byte = [10]byte{58, 108, 97, 110, 103, 117, 97, 103, 101, 0}

var _str_9 [2]byte = [2]byte{40, 0}

var _str_10 [2]byte = [2]byte{41, 0}

var _str_11 [10]byte = [10]byte{58, 112, 108, 97, 116, 102, 111, 114, 109, 0}

var _str_12 [10]byte = [10]byte{112, 97, 114, 97, 109, 101, 116, 101, 114, 0}

var _str_13 [5]byte = [5]byte{95, 101, 111, 108, 0}

var _str_14 [10]byte = [10]byte{115, 101, 112, 97, 114, 97, 116, 111, 114, 0}

var _str_15 [5]byte = [5]byte{102, 105, 108, 101, 0}

var _str_16 [5]byte = [5]byte{116, 101, 115, 116, 0}

var _str_17 [6]byte = [6]byte{95, 98, 111, 100, 121, 0}

var _str_18 [7]byte = [7]byte{104, 101, 97, 100, 101, 114, 0}

var _str_19 [11]byte = [11]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 0}

var _str_20 [10]byte = [10]byte{97, 116, 116, 114, 105, 98, 117, 116, 101, 0}

var _str_21 [10]byte = [10]byte{95, 108, 97, 110, 103, 117, 97, 103, 101, 0}

var _str_22 [10]byte = [10]byte{95, 112, 108, 97, 116, 102, 111, 114, 109, 0}

var _str_23 [13]byte = [13]byte{102, 105, 108, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}

var _str_24 [19]byte = [19]byte{
	97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 95, 114, 101, 112, 101, 97,
	116, 49, 0,
}

var _str_25 [6]byte = [6]byte{105, 110, 112, 117, 116, 0}

var _str_26 [5]byte = [5]byte{110, 97, 109, 101, 0}

var _str_27 [7]byte = [7]byte{111, 117, 116, 112, 117, 116, 0}

var _str_28 [4]byte = [4]byte{99, 115, 116, 0}

var _str_29 [9]byte = [9]byte{108, 97, 110, 103, 117, 97, 103, 101, 0}

var _str_30 [9]byte = [9]byte{112, 108, 97, 116, 102, 111, 114, 109, 0}

var ts_lex_map [30]int16 = [30]int16{
	10, 132, 13, 133, 40, 85, 41, 86, 58, 12, 97, 108, 100, 119, 102, 120,
	105, 113, 108, 103, 109, 88, 110, 96, 111, 118, 115, 114, 119, 105,
}

var ts_lex_map_31 [20]int16 = [20]int16{
	97, 41, 100, 62, 102, 59, 105, 47, 108, 31, 109, 4, 110, 19, 111, 54,
	115, 52, 119, 33,
}

func tree_sitter_test_external_scanner_create() *byte {
	var scanner **Scanner
	var call, v1, v2, v3, v4 *Scanner
	var initialized, v5 *byte
	var length, suffix *int32

	_, _, _, _, _, _, _, _, _, _ = scanner, call, v1, length, v2, suffix, v3, initialized, v4, v5

	scanner = new(*Scanner)
	call = libc.Malloc[Scanner](int64(12))
	*scanner = call
	v1 = *scanner
	length = &v1.F0
	*length = 0
	v2 = *scanner
	suffix = &v2.F1
	*suffix = 0
	v3 = *scanner
	initialized = &v3.F2
	*initialized = 0
	v4 = *scanner
	v5 = (*byte)(unsafe.Pointer(v4))
	return v5
}

func tree_sitter_test_external_scanner_serialize(payload *byte, buffer *byte) int32 {
	var scanner **Scanner
	var payload_addr, buffer_addr **byte
	var v1, v2, v6, v9, v11 *Scanner
	var v0, v4, v5, add_ptr *byte
	var retval, length, suffix, length1, length2 *int32
	var cmp bool
	var v8 byte
	var v3, add, v7, v10, v12, add3, v13 int32
	var conv int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, payload_addr, buffer_addr, scanner, v0, v1, v2, length, v3, add, cmp, v4, v5, add_ptr, v6, suffix, v7, v8, v9, length1, v10, conv, v11, length2, v12, add3, v13

	retval = new(int32)
	payload_addr = new(*byte)
	buffer_addr = new(*byte)
	scanner = new(*Scanner)
	*payload_addr = payload
	*buffer_addr = buffer
	v0 = *payload_addr
	v1 = (*Scanner)(unsafe.Pointer(v0))
	*scanner = v1
	v2 = *scanner
	length = &v2.F0
	v3 = *length
	add = v3 + 1
	cmp = uint32(add) > 1024
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*retval = 0
	goto _return

if_end:
	v4 = *buffer_addr
	libc.Memset(v4, 1, int64(1))
	v5 = *buffer_addr
	add_ptr = libc.AddPointer(v5, int(int64(1)))
	v6 = *scanner
	suffix = &v6.F1
	v7 = *suffix
	v8 = byte(v7)
	v9 = *scanner
	length1 = &v9.F0
	v10 = *length1
	conv = int64(uint64(uint32(v10)))
	libc.Memset(add_ptr, v8, conv)
	v11 = *scanner
	length2 = &v11.F0
	v12 = *length2
	add3 = v12 + 1
	*retval = add3
	goto _return

_return:
	v13 = *retval
	return v13
}

func tree_sitter_test_external_scanner_deserialize(payload *byte, buffer *byte, length int32) {
	var scanner **Scanner
	var payload_addr, buffer_addr **byte
	var v2, v4, v7, v10 *Scanner
	var v1, v5, arrayidx, v8, arrayidx2, initialized *byte
	var length_addr, length1, suffix *int32
	var cmp, tobool bool
	var v6, v9, frombool byte
	var v0, v3, sub, conv int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = payload_addr, buffer_addr, length_addr, scanner, v0, cmp, v1, v2, v3, sub, v4, length1, v5, arrayidx, v6, conv, v7, suffix, v8, arrayidx2, v9, tobool, v10, initialized, frombool

	payload_addr = new(*byte)
	buffer_addr = new(*byte)
	length_addr = new(int32)
	scanner = new(*Scanner)
	*payload_addr = payload
	*buffer_addr = buffer
	*length_addr = length
	v0 = *length_addr
	cmp = v0 == 0
	if cmp {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	goto _return

if_end:
	v1 = *payload_addr
	v2 = (*Scanner)(unsafe.Pointer(v1))
	*scanner = v2
	v3 = *length_addr
	sub = v3 - 1
	v4 = *scanner
	length1 = &v4.F0
	*length1 = sub
	v5 = *buffer_addr
	arrayidx = libc.AddPointer(v5, int(int64(1)))
	v6 = *arrayidx
	conv = int32(int8(v6))
	v7 = *scanner
	suffix = &v7.F1
	*suffix = conv
	v8 = *buffer_addr
	arrayidx2 = v8
	v9 = *arrayidx2
	tobool = v9 != 0
	v10 = *scanner
	initialized = &v10.F2
	if tobool { frombool = 1 } else { frombool = 0 }
	*initialized = frombool
	goto _return

_return:
}

func tree_sitter_test_external_scanner_scan(payload *byte, lexer *TSLexer, valid_symbols *byte) bool {
	var scanner **Scanner
	var lexer_addr **TSLexer
	var payload_addr, valid_symbols_addr **byte
	var v1, v5, v9, v13 *Scanner
	var v4, v8, v12 *TSLexer
	var retval *bool
	var v0, v2, arrayidx, v6, arrayidx1, v10, arrayidx6 *byte
	var tobool, call, tobool2, call4, tobool7, call9, v14 bool
	var v3, v7, v11 byte

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, payload_addr, lexer_addr, valid_symbols_addr, scanner, v0, v1, v2, arrayidx, v3, tobool, v4, v5, call, v6, arrayidx1, v7, tobool2, v8, v9, call4, v10, arrayidx6, v11, tobool7, v12, v13, call9, v14

	retval = new(bool)
	payload_addr = new(*byte)
	lexer_addr = new(*TSLexer)
	valid_symbols_addr = new(*byte)
	scanner = new(*Scanner)
	*payload_addr = payload
	*lexer_addr = lexer
	*valid_symbols_addr = valid_symbols
	v0 = *payload_addr
	v1 = (*Scanner)(unsafe.Pointer(v0))
	*scanner = v1
	v2 = *valid_symbols_addr
	arrayidx = v2
	v3 = *arrayidx
	tobool = (v3 & 1) != 0
	if tobool {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	v4 = *lexer_addr
	v5 = *scanner
	call = scan(v4, v5, 61, 0)
	*retval = call
	goto _return

if_end:
	v6 = *valid_symbols_addr
	arrayidx1 = libc.AddPointer(v6, int(int64(1)))
	v7 = *arrayidx1
	tobool2 = (v7 & 1) != 0
	if tobool2 {
		goto if_then3
	} else {
		goto if_end5
	}

if_then3:
	v8 = *lexer_addr
	v9 = *scanner
	call4 = scan(v8, v9, 61, 1)
	*retval = call4
	goto _return

if_end5:
	v10 = *valid_symbols_addr
	arrayidx6 = libc.AddPointer(v10, int(int64(2)))
	v11 = *arrayidx6
	tobool7 = (v11 & 1) != 0
	if tobool7 {
		goto if_then8
	} else {
		goto if_end10
	}

if_then8:
	v12 = *lexer_addr
	v13 = *scanner
	call9 = scan(v12, v13, 45, 2)
	*retval = call9
	goto _return

if_end10:
	*retval = false
	goto _return

_return:
	v14 = *retval
	return v14
}

func scan(lexer *TSLexer, scanner *Scanner, chr byte, symbol int32) bool {
	var scanner_addr **Scanner
	var lexer_addr **TSLexer
	var v20, v24, v25, v32, v35, v36 *Scanner
	var v0, v2, v3, v7, v10, v12, v13, v15, v22, v27, v29, v40, v42, v43, v45, v47 *TSLexer
	var retval *bool
	var chr_addr, initialized, initialized33 *byte
	var eof, eof5 *func(*TSLexer) bool
	var result_symbol *int16
	var symbol_addr, length, lookahead, lookahead8, lookahead11, lookahead19, suffix, suffix21, lookahead22, length36, length38, lookahead45, lookahead50 *int32
	var call, cmp, v6, cmp2, call6, cmp9, cmp12, v17, lnot, v18, cmp16, tobool, cmp23, cmp30, tobool34, cmp39, cmp42, cmp46, cmp51, v48 bool
	var v5, v21, v33 byte
	var v1, v11 func(*TSLexer) bool
	var conv55 int16
	var v4, conv, v8, inc, v9, v14, v16, v19, v23, v26, v28, v30, inc28, v31, v34, v37, v38, v39, v41, v44, v46 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, scanner_addr, chr_addr, symbol_addr, length, v0, eof, v1, v2, call, v3, lookahead, v4, v5, conv, cmp, v6, v7, v8, inc, v9, cmp2, v10, eof5, v11, v12, call6, v13, lookahead8, v14, cmp9, v15, lookahead11, v16, cmp12, v17, lnot, v18, v19, cmp16, v20, initialized, v21, tobool, v22, lookahead19, v23, v24, suffix, v25, suffix21, v26, v27, lookahead22, v28, cmp23, v29, v30, inc28, v31, cmp30, v32, initialized33, v33, tobool34, v34, v35, length36, v36, length38, v37, v38, cmp39, v39, cmp42, v40, lookahead45, v41, cmp46, v42, v43, lookahead50, v44, cmp51, v45, v46, conv55, v47, result_symbol, v48

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	scanner_addr = new(*Scanner)
	chr_addr = new(byte)
	symbol_addr = new(int32)
	length = new(int32)
	*lexer_addr = lexer
	*scanner_addr = scanner
	*chr_addr = chr
	*symbol_addr = symbol
	*length = 0
	goto for_cond

for_cond:
	v0 = *lexer_addr
	eof = &v0.F6
	v1 = *eof
	v2 = *lexer_addr
	call = v1(v2)
	if call {
		v6 = false
		goto land_end
	} else {
		goto land_rhs
	}

land_rhs:
	v3 = *lexer_addr
	lookahead = &v3.F0
	v4 = *lookahead
	v5 = *chr_addr
	conv = int32(int8(v5))
	cmp = v4 == conv
	v6 = cmp
	goto land_end

land_end:
	if v6 {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v7 = *lexer_addr
	advance(v7)
	goto for_inc

for_inc:
	v8 = *length
	inc = v8 + 1
	*length = inc
	goto for_cond

for_end:
	v9 = *length
	cmp2 = uint32(v9) < 3
	if cmp2 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*retval = false
	goto _return

if_end:
	*length = 0
	goto for_cond4

for_cond4:
	v10 = *lexer_addr
	eof5 = &v10.F6
	v11 = *eof5
	v12 = *lexer_addr
	call6 = v11(v12)
	if call6 {
		v18 = false
		goto land_end14
	} else {
		goto land_rhs7
	}

land_rhs7:
	v13 = *lexer_addr
	lookahead8 = &v13.F0
	v14 = *lookahead8
	cmp9 = v14 == 13
	if cmp9 {
		v17 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v15 = *lexer_addr
	lookahead11 = &v15.F0
	v16 = *lookahead11
	cmp12 = v16 == 10
	v17 = cmp12
	goto lor_end

lor_end:
	lnot = v17 != true
	v18 = lnot
	goto land_end14

land_end14:
	if v18 {
		goto for_body15
	} else {
		goto for_end29
	}

for_body15:
	v19 = *symbol_addr
	cmp16 = v19 == 0
	if cmp16 {
		goto land_lhs_true
	} else {
		goto if_end20
	}

land_lhs_true:
	v20 = *scanner_addr
	initialized = &v20.F2
	v21 = *initialized
	tobool = (v21 & 1) != 0
	if tobool {
		goto if_end20
	} else {
		goto if_then18
	}

if_then18:
	v22 = *lexer_addr
	lookahead19 = &v22.F0
	v23 = *lookahead19
	v24 = *scanner_addr
	suffix = &v24.F1
	*suffix = v23
	goto if_end20

if_end20:
	v25 = *scanner_addr
	suffix21 = &v25.F1
	v26 = *suffix21
	v27 = *lexer_addr
	lookahead22 = &v27.F0
	v28 = *lookahead22
	cmp23 = v26 != v28
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*retval = false
	goto _return

if_end26:
	v29 = *lexer_addr
	advance(v29)
	goto for_inc27

for_inc27:
	v30 = *length
	inc28 = v30 + 1
	*length = inc28
	goto for_cond4

for_end29:
	v31 = *symbol_addr
	cmp30 = v31 == 0
	if cmp30 {
		goto land_lhs_true32
	} else {
		goto if_end37
	}

land_lhs_true32:
	v32 = *scanner_addr
	initialized33 = &v32.F2
	v33 = *initialized33
	tobool34 = (v33 & 1) != 0
	if tobool34 {
		goto if_end37
	} else {
		goto if_then35
	}

if_then35:
	v34 = *length
	v35 = *scanner_addr
	length36 = &v35.F0
	*length36 = v34
	goto if_end37

if_end37:
	v36 = *scanner_addr
	length38 = &v36.F0
	v37 = *length38
	v38 = *length
	cmp39 = v37 != v38
	if cmp39 {
		goto if_then41
	} else {
		goto if_end44
	}

if_then41:
	v39 = *symbol_addr
	cmp42 = v39 == 2
	*retval = cmp42
	goto _return

if_end44:
	v40 = *lexer_addr
	lookahead45 = &v40.F0
	v41 = *lookahead45
	cmp46 = v41 == 13
	if cmp46 {
		goto if_then48
	} else {
		goto if_end49
	}

if_then48:
	v42 = *lexer_addr
	advance(v42)
	goto if_end49

if_end49:
	v43 = *lexer_addr
	lookahead50 = &v43.F0
	v44 = *lookahead50
	cmp51 = v44 == 10
	if cmp51 {
		goto if_then53
	} else {
		goto if_end54
	}

if_then53:
	v45 = *lexer_addr
	advance(v45)
	goto if_end54

if_end54:
	v46 = *symbol_addr
	conv55 = int16(v46)
	v47 = *lexer_addr
	result_symbol = &v47.F1
	*result_symbol = conv55
	*retval = true
	goto _return

_return:
	v48 = *retval
	return v48
}

func tree_sitter_test_external_scanner_destroy(payload *byte) {
	var scanner **Scanner
	var payload_addr **byte
	var v1, v2 *Scanner
	var v0, v3 *byte

	_, _, _, _, _, _ = payload_addr, scanner, v0, v1, v2, v3

	payload_addr = new(*byte)
	scanner = new(*Scanner)
	*payload_addr = payload
	v0 = *payload_addr
	v1 = (*Scanner)(unsafe.Pointer(v0))
	*scanner = v1
	v2 = *scanner
	v3 = (*byte)(unsafe.Pointer(v2))
	libc.Free(v3)
}

func tree_sitter_test() *TSLanguage {
	return &tree_sitter_test_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v204, v205, v207, v209, v210, v212, v217, v218, v220, v222, v223, v225, v227, v228, v230, v232, v233, v235, v237, v238, v240, v242, v243, v245, v247, v248, v250, v252, v253, v255, v257, v258, v260, v272, v273, v275, v287, v288, v290, v302, v303, v305, v317, v318, v320, v332, v333, v335, v347, v348, v350, v362, v363, v365, v377, v378, v380, v392, v393, v395, v407, v408, v410, v422, v423, v425, v437, v438, v440, v452, v453, v455, v467, v468, v470, v482, v483, v485, v497, v498, v500, v512, v513, v515, v527, v528, v530, v542, v543, v545, v557, v558, v560, v572, v573, v575, v587, v588, v590, v602, v603, v605, v617, v618, v620, v632, v633, v635, v647, v648, v650, v662, v663, v665, v677, v678, v680, v692, v693, v695, v707, v708, v710, v722, v723, v725, v737, v738, v740, v752, v753, v755, v767, v768, v770, v782, v783, v785, v797, v798, v800, v812, v813, v815, v827, v828, v830, v842, v843, v845, v857, v858, v860, v872, v873, v875, v887, v888, v890, v901, v902, v904, v906, v907, v909, v911, v912, v914 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end577, mark_end591, mark_end595, mark_end599, mark_end603, mark_end607, mark_end611, mark_end615, mark_end619, mark_end623, mark_end659, mark_end695, mark_end731, mark_end767, mark_end803, mark_end839, mark_end875, mark_end911, mark_end947, mark_end983, mark_end1019, mark_end1055, mark_end1091, mark_end1127, mark_end1163, mark_end1199, mark_end1235, mark_end1271, mark_end1307, mark_end1343, mark_end1379, mark_end1415, mark_end1451, mark_end1487, mark_end1523, mark_end1559, mark_end1595, mark_end1631, mark_end1667, mark_end1703, mark_end1739, mark_end1775, mark_end1811, mark_end1847, mark_end1883, mark_end1919, mark_end1955, mark_end1991, mark_end2027, mark_end2063, mark_end2099, mark_end2135, mark_end2167, mark_end2171, mark_end2175 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var local_advance *func(*TSLexer, bool)
	var state_addr, arrayidx, arrayidx11, arrayidx79, arrayidx86, result_symbol, result_symbol576, result_symbol590, result_symbol594, result_symbol598, result_symbol602, result_symbol606, result_symbol610, result_symbol614, result_symbol618, result_symbol622, result_symbol658, result_symbol694, result_symbol730, result_symbol766, result_symbol802, result_symbol838, result_symbol874, result_symbol910, result_symbol946, result_symbol982, result_symbol1018, result_symbol1054, result_symbol1090, result_symbol1126, result_symbol1162, result_symbol1198, result_symbol1234, result_symbol1270, result_symbol1306, result_symbol1342, result_symbol1378, result_symbol1414, result_symbol1450, result_symbol1486, result_symbol1522, result_symbol1558, result_symbol1594, result_symbol1630, result_symbol1666, result_symbol1702, result_symbol1738, result_symbol1774, result_symbol1810, result_symbol1846, result_symbol1882, result_symbol1918, result_symbol1954, result_symbol1990, result_symbol2026, result_symbol2062, result_symbol2098, result_symbol2134, result_symbol2166, result_symbol2170, result_symbol2174 *int16
	var lookahead, i, i72, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, cmp19, cmp21, cmp24, cmp27, cmp30, cmp33, cmp36, tobool40, cmp42, tobool46, cmp48, tobool52, cmp54, tobool58, cmp60, tobool64, cmp66, tobool70, cmp75, cmp81, tobool91, cmp93, tobool97, cmp99, tobool103, cmp105, tobool109, cmp111, tobool115, cmp117, tobool121, cmp123, cmp127, cmp131, cmp135, cmp139, cmp143, tobool147, cmp149, tobool153, cmp155, tobool159, cmp161, tobool165, cmp167, tobool171, cmp173, tobool177, cmp179, tobool183, cmp185, tobool189, cmp191, tobool195, cmp197, tobool201, cmp203, tobool207, cmp209, tobool213, cmp215, tobool219, cmp221, tobool225, cmp227, tobool231, cmp233, tobool237, cmp239, tobool243, cmp245, tobool249, cmp251, tobool255, cmp257, tobool261, cmp263, tobool267, cmp269, tobool273, cmp275, tobool279, cmp281, tobool285, cmp287, tobool291, cmp293, tobool297, cmp299, tobool303, cmp305, tobool309, cmp311, tobool315, cmp317, tobool321, cmp323, tobool327, cmp329, tobool333, cmp335, tobool339, cmp341, tobool345, cmp347, tobool351, cmp353, tobool357, cmp359, tobool363, cmp365, tobool369, cmp371, tobool375, cmp377, tobool381, cmp383, tobool387, cmp389, tobool393, cmp395, tobool399, cmp401, tobool405, cmp407, tobool411, cmp413, tobool417, cmp419, tobool423, cmp425, tobool429, cmp431, tobool435, cmp437, tobool441, cmp443, tobool447, cmp449, tobool453, cmp455, tobool459, cmp461, tobool465, cmp467, tobool471, cmp473, tobool477, cmp479, tobool483, cmp485, tobool489, cmp491, tobool495, cmp497, tobool501, cmp503, tobool507, cmp509, tobool513, cmp515, tobool519, cmp521, tobool525, cmp527, cmp530, cmp533, cmp536, cmp539, cmp542, cmp545, cmp548, cmp551, tobool555, tobool557, cmp560, cmp564, cmp568, tobool572, tobool574, cmp578, cmp581, cmp584, tobool588, tobool592, tobool596, tobool600, tobool604, tobool608, tobool612, tobool616, tobool620, cmp624, cmp628, cmp631, cmp634, cmp637, cmp640, cmp643, cmp646, cmp649, cmp652, tobool656, cmp660, cmp664, cmp667, cmp670, cmp673, cmp676, cmp679, cmp682, cmp685, cmp688, tobool692, cmp696, cmp700, cmp703, cmp706, cmp709, cmp712, cmp715, cmp718, cmp721, cmp724, tobool728, cmp732, cmp736, cmp739, cmp742, cmp745, cmp748, cmp751, cmp754, cmp757, cmp760, tobool764, cmp768, cmp772, cmp775, cmp778, cmp781, cmp784, cmp787, cmp790, cmp793, cmp796, tobool800, cmp804, cmp808, cmp811, cmp814, cmp817, cmp820, cmp823, cmp826, cmp829, cmp832, tobool836, cmp840, cmp844, cmp847, cmp850, cmp853, cmp856, cmp859, cmp862, cmp865, cmp868, tobool872, cmp876, cmp880, cmp883, cmp886, cmp889, cmp892, cmp895, cmp898, cmp901, cmp904, tobool908, cmp912, cmp916, cmp919, cmp922, cmp925, cmp928, cmp931, cmp934, cmp937, cmp940, tobool944, cmp948, cmp952, cmp955, cmp958, cmp961, cmp964, cmp967, cmp970, cmp973, cmp976, tobool980, cmp984, cmp988, cmp991, cmp994, cmp997, cmp1000, cmp1003, cmp1006, cmp1009, cmp1012, tobool1016, cmp1020, cmp1024, cmp1027, cmp1030, cmp1033, cmp1036, cmp1039, cmp1042, cmp1045, cmp1048, tobool1052, cmp1056, cmp1060, cmp1063, cmp1066, cmp1069, cmp1072, cmp1075, cmp1078, cmp1081, cmp1084, tobool1088, cmp1092, cmp1096, cmp1099, cmp1102, cmp1105, cmp1108, cmp1111, cmp1114, cmp1117, cmp1120, tobool1124, cmp1128, cmp1132, cmp1135, cmp1138, cmp1141, cmp1144, cmp1147, cmp1150, cmp1153, cmp1156, tobool1160, cmp1164, cmp1168, cmp1171, cmp1174, cmp1177, cmp1180, cmp1183, cmp1186, cmp1189, cmp1192, tobool1196, cmp1200, cmp1204, cmp1207, cmp1210, cmp1213, cmp1216, cmp1219, cmp1222, cmp1225, cmp1228, tobool1232, cmp1236, cmp1240, cmp1243, cmp1246, cmp1249, cmp1252, cmp1255, cmp1258, cmp1261, cmp1264, tobool1268, cmp1272, cmp1276, cmp1279, cmp1282, cmp1285, cmp1288, cmp1291, cmp1294, cmp1297, cmp1300, tobool1304, cmp1308, cmp1312, cmp1315, cmp1318, cmp1321, cmp1324, cmp1327, cmp1330, cmp1333, cmp1336, tobool1340, cmp1344, cmp1348, cmp1351, cmp1354, cmp1357, cmp1360, cmp1363, cmp1366, cmp1369, cmp1372, tobool1376, cmp1380, cmp1384, cmp1387, cmp1390, cmp1393, cmp1396, cmp1399, cmp1402, cmp1405, cmp1408, tobool1412, cmp1416, cmp1420, cmp1423, cmp1426, cmp1429, cmp1432, cmp1435, cmp1438, cmp1441, cmp1444, tobool1448, cmp1452, cmp1456, cmp1459, cmp1462, cmp1465, cmp1468, cmp1471, cmp1474, cmp1477, cmp1480, tobool1484, cmp1488, cmp1492, cmp1495, cmp1498, cmp1501, cmp1504, cmp1507, cmp1510, cmp1513, cmp1516, tobool1520, cmp1524, cmp1528, cmp1531, cmp1534, cmp1537, cmp1540, cmp1543, cmp1546, cmp1549, cmp1552, tobool1556, cmp1560, cmp1564, cmp1567, cmp1570, cmp1573, cmp1576, cmp1579, cmp1582, cmp1585, cmp1588, tobool1592, cmp1596, cmp1600, cmp1603, cmp1606, cmp1609, cmp1612, cmp1615, cmp1618, cmp1621, cmp1624, tobool1628, cmp1632, cmp1636, cmp1639, cmp1642, cmp1645, cmp1648, cmp1651, cmp1654, cmp1657, cmp1660, tobool1664, cmp1668, cmp1672, cmp1675, cmp1678, cmp1681, cmp1684, cmp1687, cmp1690, cmp1693, cmp1696, tobool1700, cmp1704, cmp1708, cmp1711, cmp1714, cmp1717, cmp1720, cmp1723, cmp1726, cmp1729, cmp1732, tobool1736, cmp1740, cmp1744, cmp1747, cmp1750, cmp1753, cmp1756, cmp1759, cmp1762, cmp1765, cmp1768, tobool1772, cmp1776, cmp1780, cmp1783, cmp1786, cmp1789, cmp1792, cmp1795, cmp1798, cmp1801, cmp1804, tobool1808, cmp1812, cmp1816, cmp1819, cmp1822, cmp1825, cmp1828, cmp1831, cmp1834, cmp1837, cmp1840, tobool1844, cmp1848, cmp1852, cmp1855, cmp1858, cmp1861, cmp1864, cmp1867, cmp1870, cmp1873, cmp1876, tobool1880, cmp1884, cmp1888, cmp1891, cmp1894, cmp1897, cmp1900, cmp1903, cmp1906, cmp1909, cmp1912, tobool1916, cmp1920, cmp1924, cmp1927, cmp1930, cmp1933, cmp1936, cmp1939, cmp1942, cmp1945, cmp1948, tobool1952, cmp1956, cmp1960, cmp1963, cmp1966, cmp1969, cmp1972, cmp1975, cmp1978, cmp1981, cmp1984, tobool1988, cmp1992, cmp1996, cmp1999, cmp2002, cmp2005, cmp2008, cmp2011, cmp2014, cmp2017, cmp2020, tobool2024, cmp2028, cmp2032, cmp2035, cmp2038, cmp2041, cmp2044, cmp2047, cmp2050, cmp2053, cmp2056, tobool2060, cmp2064, cmp2068, cmp2071, cmp2074, cmp2077, cmp2080, cmp2083, cmp2086, cmp2089, cmp2092, tobool2096, cmp2100, cmp2104, cmp2107, cmp2110, cmp2113, cmp2116, cmp2119, cmp2122, cmp2125, cmp2128, tobool2132, cmp2136, cmp2139, cmp2142, cmp2145, cmp2148, cmp2151, cmp2154, cmp2157, cmp2160, tobool2164, tobool2168, tobool2172, cmp2176, tobool2180, v917 bool
	var v3, frombool, v10, v27, v29, v31, v33, v35, v37, v45, v47, v49, v51, v53, v55, v62, v64, v66, v68, v70, v72, v74, v76, v78, v80, v82, v84, v86, v88, v90, v92, v94, v96, v98, v100, v102, v104, v106, v108, v110, v112, v114, v116, v118, v120, v122, v124, v126, v128, v130, v132, v134, v136, v138, v140, v142, v144, v146, v148, v150, v152, v154, v156, v158, v160, v162, v164, v166, v168, v170, v172, v174, v176, v178, v180, v182, v184, v186, v188, v198, v199, v203, v208, v216, v221, v226, v231, v236, v241, v246, v251, v256, v271, v286, v301, v316, v331, v346, v361, v376, v391, v406, v421, v436, v451, v466, v481, v496, v511, v526, v541, v556, v571, v586, v601, v616, v631, v646, v661, v676, v691, v706, v721, v736, v751, v766, v781, v796, v811, v826, v841, v856, v871, v886, v900, v905, v910, v916 byte
	var v206, v211, v219, v224, v229, v234, v239, v244, v249, v254, v259, v274, v289, v304, v319, v334, v349, v364, v379, v394, v409, v424, v439, v454, v469, v484, v499, v514, v529, v544, v559, v574, v589, v604, v619, v634, v649, v664, v679, v694, v709, v724, v739, v754, v769, v784, v799, v814, v829, v844, v859, v874, v889, v903, v908, v913 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9, v13, v16, v40, v43 int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v23, v24, v25, v26, v28, v30, v32, v34, v36, v38, v39, conv80, v41, v42, add84, v44, add89, v46, v48, v50, v52, v54, v56, v57, v58, v59, v60, v61, v63, v65, v67, v69, v71, v73, v75, v77, v79, v81, v83, v85, v87, v89, v91, v93, v95, v97, v99, v101, v103, v105, v107, v109, v111, v113, v115, v117, v119, v121, v123, v125, v127, v129, v131, v133, v135, v137, v139, v141, v143, v145, v147, v149, v151, v153, v155, v157, v159, v161, v163, v165, v167, v169, v171, v173, v175, v177, v179, v181, v183, v185, v187, v189, v190, v191, v192, v193, v194, v195, v196, v197, v200, v201, v202, v213, v214, v215, v261, v262, v263, v264, v265, v266, v267, v268, v269, v270, v276, v277, v278, v279, v280, v281, v282, v283, v284, v285, v291, v292, v293, v294, v295, v296, v297, v298, v299, v300, v306, v307, v308, v309, v310, v311, v312, v313, v314, v315, v321, v322, v323, v324, v325, v326, v327, v328, v329, v330, v336, v337, v338, v339, v340, v341, v342, v343, v344, v345, v351, v352, v353, v354, v355, v356, v357, v358, v359, v360, v366, v367, v368, v369, v370, v371, v372, v373, v374, v375, v381, v382, v383, v384, v385, v386, v387, v388, v389, v390, v396, v397, v398, v399, v400, v401, v402, v403, v404, v405, v411, v412, v413, v414, v415, v416, v417, v418, v419, v420, v426, v427, v428, v429, v430, v431, v432, v433, v434, v435, v441, v442, v443, v444, v445, v446, v447, v448, v449, v450, v456, v457, v458, v459, v460, v461, v462, v463, v464, v465, v471, v472, v473, v474, v475, v476, v477, v478, v479, v480, v486, v487, v488, v489, v490, v491, v492, v493, v494, v495, v501, v502, v503, v504, v505, v506, v507, v508, v509, v510, v516, v517, v518, v519, v520, v521, v522, v523, v524, v525, v531, v532, v533, v534, v535, v536, v537, v538, v539, v540, v546, v547, v548, v549, v550, v551, v552, v553, v554, v555, v561, v562, v563, v564, v565, v566, v567, v568, v569, v570, v576, v577, v578, v579, v580, v581, v582, v583, v584, v585, v591, v592, v593, v594, v595, v596, v597, v598, v599, v600, v606, v607, v608, v609, v610, v611, v612, v613, v614, v615, v621, v622, v623, v624, v625, v626, v627, v628, v629, v630, v636, v637, v638, v639, v640, v641, v642, v643, v644, v645, v651, v652, v653, v654, v655, v656, v657, v658, v659, v660, v666, v667, v668, v669, v670, v671, v672, v673, v674, v675, v681, v682, v683, v684, v685, v686, v687, v688, v689, v690, v696, v697, v698, v699, v700, v701, v702, v703, v704, v705, v711, v712, v713, v714, v715, v716, v717, v718, v719, v720, v726, v727, v728, v729, v730, v731, v732, v733, v734, v735, v741, v742, v743, v744, v745, v746, v747, v748, v749, v750, v756, v757, v758, v759, v760, v761, v762, v763, v764, v765, v771, v772, v773, v774, v775, v776, v777, v778, v779, v780, v786, v787, v788, v789, v790, v791, v792, v793, v794, v795, v801, v802, v803, v804, v805, v806, v807, v808, v809, v810, v816, v817, v818, v819, v820, v821, v822, v823, v824, v825, v831, v832, v833, v834, v835, v836, v837, v838, v839, v840, v846, v847, v848, v849, v850, v851, v852, v853, v854, v855, v861, v862, v863, v864, v865, v866, v867, v868, v869, v870, v876, v877, v878, v879, v880, v881, v882, v883, v884, v885, v891, v892, v893, v894, v895, v896, v897, v898, v899, v915 int32
	var conv4, idxprom, idxprom10, conv74, idxprom78, idxprom85 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i72, v0, local_advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp19, v21, cmp21, v22, cmp24, v23, cmp27, v24, cmp30, v25, cmp33, v26, cmp36, v27, tobool40, v28, cmp42, v29, tobool46, v30, cmp48, v31, tobool52, v32, cmp54, v33, tobool58, v34, cmp60, v35, tobool64, v36, cmp66, v37, tobool70, v38, conv74, cmp75, v39, idxprom78, arrayidx79, v40, conv80, v41, cmp81, v42, add84, idxprom85, arrayidx86, v43, v44, add89, v45, tobool91, v46, cmp93, v47, tobool97, v48, cmp99, v49, tobool103, v50, cmp105, v51, tobool109, v52, cmp111, v53, tobool115, v54, cmp117, v55, tobool121, v56, cmp123, v57, cmp127, v58, cmp131, v59, cmp135, v60, cmp139, v61, cmp143, v62, tobool147, v63, cmp149, v64, tobool153, v65, cmp155, v66, tobool159, v67, cmp161, v68, tobool165, v69, cmp167, v70, tobool171, v71, cmp173, v72, tobool177, v73, cmp179, v74, tobool183, v75, cmp185, v76, tobool189, v77, cmp191, v78, tobool195, v79, cmp197, v80, tobool201, v81, cmp203, v82, tobool207, v83, cmp209, v84, tobool213, v85, cmp215, v86, tobool219, v87, cmp221, v88, tobool225, v89, cmp227, v90, tobool231, v91, cmp233, v92, tobool237, v93, cmp239, v94, tobool243, v95, cmp245, v96, tobool249, v97, cmp251, v98, tobool255, v99, cmp257, v100, tobool261, v101, cmp263, v102, tobool267, v103, cmp269, v104, tobool273, v105, cmp275, v106, tobool279, v107, cmp281, v108, tobool285, v109, cmp287, v110, tobool291, v111, cmp293, v112, tobool297, v113, cmp299, v114, tobool303, v115, cmp305, v116, tobool309, v117, cmp311, v118, tobool315, v119, cmp317, v120, tobool321, v121, cmp323, v122, tobool327, v123, cmp329, v124, tobool333, v125, cmp335, v126, tobool339, v127, cmp341, v128, tobool345, v129, cmp347, v130, tobool351, v131, cmp353, v132, tobool357, v133, cmp359, v134, tobool363, v135, cmp365, v136, tobool369, v137, cmp371, v138, tobool375, v139, cmp377, v140, tobool381, v141, cmp383, v142, tobool387, v143, cmp389, v144, tobool393, v145, cmp395, v146, tobool399, v147, cmp401, v148, tobool405, v149, cmp407, v150, tobool411, v151, cmp413, v152, tobool417, v153, cmp419, v154, tobool423, v155, cmp425, v156, tobool429, v157, cmp431, v158, tobool435, v159, cmp437, v160, tobool441, v161, cmp443, v162, tobool447, v163, cmp449, v164, tobool453, v165, cmp455, v166, tobool459, v167, cmp461, v168, tobool465, v169, cmp467, v170, tobool471, v171, cmp473, v172, tobool477, v173, cmp479, v174, tobool483, v175, cmp485, v176, tobool489, v177, cmp491, v178, tobool495, v179, cmp497, v180, tobool501, v181, cmp503, v182, tobool507, v183, cmp509, v184, tobool513, v185, cmp515, v186, tobool519, v187, cmp521, v188, tobool525, v189, cmp527, v190, cmp530, v191, cmp533, v192, cmp536, v193, cmp539, v194, cmp542, v195, cmp545, v196, cmp548, v197, cmp551, v198, tobool555, v199, tobool557, v200, cmp560, v201, cmp564, v202, cmp568, v203, tobool572, v204, result_symbol, v205, mark_end, v206, v207, v208, tobool574, v209, result_symbol576, v210, mark_end577, v211, v212, v213, cmp578, v214, cmp581, v215, cmp584, v216, tobool588, v217, result_symbol590, v218, mark_end591, v219, v220, v221, tobool592, v222, result_symbol594, v223, mark_end595, v224, v225, v226, tobool596, v227, result_symbol598, v228, mark_end599, v229, v230, v231, tobool600, v232, result_symbol602, v233, mark_end603, v234, v235, v236, tobool604, v237, result_symbol606, v238, mark_end607, v239, v240, v241, tobool608, v242, result_symbol610, v243, mark_end611, v244, v245, v246, tobool612, v247, result_symbol614, v248, mark_end615, v249, v250, v251, tobool616, v252, result_symbol618, v253, mark_end619, v254, v255, v256, tobool620, v257, result_symbol622, v258, mark_end623, v259, v260, v261, cmp624, v262, cmp628, v263, cmp631, v264, cmp634, v265, cmp637, v266, cmp640, v267, cmp643, v268, cmp646, v269, cmp649, v270, cmp652, v271, tobool656, v272, result_symbol658, v273, mark_end659, v274, v275, v276, cmp660, v277, cmp664, v278, cmp667, v279, cmp670, v280, cmp673, v281, cmp676, v282, cmp679, v283, cmp682, v284, cmp685, v285, cmp688, v286, tobool692, v287, result_symbol694, v288, mark_end695, v289, v290, v291, cmp696, v292, cmp700, v293, cmp703, v294, cmp706, v295, cmp709, v296, cmp712, v297, cmp715, v298, cmp718, v299, cmp721, v300, cmp724, v301, tobool728, v302, result_symbol730, v303, mark_end731, v304, v305, v306, cmp732, v307, cmp736, v308, cmp739, v309, cmp742, v310, cmp745, v311, cmp748, v312, cmp751, v313, cmp754, v314, cmp757, v315, cmp760, v316, tobool764, v317, result_symbol766, v318, mark_end767, v319, v320, v321, cmp768, v322, cmp772, v323, cmp775, v324, cmp778, v325, cmp781, v326, cmp784, v327, cmp787, v328, cmp790, v329, cmp793, v330, cmp796, v331, tobool800, v332, result_symbol802, v333, mark_end803, v334, v335, v336, cmp804, v337, cmp808, v338, cmp811, v339, cmp814, v340, cmp817, v341, cmp820, v342, cmp823, v343, cmp826, v344, cmp829, v345, cmp832, v346, tobool836, v347, result_symbol838, v348, mark_end839, v349, v350, v351, cmp840, v352, cmp844, v353, cmp847, v354, cmp850, v355, cmp853, v356, cmp856, v357, cmp859, v358, cmp862, v359, cmp865, v360, cmp868, v361, tobool872, v362, result_symbol874, v363, mark_end875, v364, v365, v366, cmp876, v367, cmp880, v368, cmp883, v369, cmp886, v370, cmp889, v371, cmp892, v372, cmp895, v373, cmp898, v374, cmp901, v375, cmp904, v376, tobool908, v377, result_symbol910, v378, mark_end911, v379, v380, v381, cmp912, v382, cmp916, v383, cmp919, v384, cmp922, v385, cmp925, v386, cmp928, v387, cmp931, v388, cmp934, v389, cmp937, v390, cmp940, v391, tobool944, v392, result_symbol946, v393, mark_end947, v394, v395, v396, cmp948, v397, cmp952, v398, cmp955, v399, cmp958, v400, cmp961, v401, cmp964, v402, cmp967, v403, cmp970, v404, cmp973, v405, cmp976, v406, tobool980, v407, result_symbol982, v408, mark_end983, v409, v410, v411, cmp984, v412, cmp988, v413, cmp991, v414, cmp994, v415, cmp997, v416, cmp1000, v417, cmp1003, v418, cmp1006, v419, cmp1009, v420, cmp1012, v421, tobool1016, v422, result_symbol1018, v423, mark_end1019, v424, v425, v426, cmp1020, v427, cmp1024, v428, cmp1027, v429, cmp1030, v430, cmp1033, v431, cmp1036, v432, cmp1039, v433, cmp1042, v434, cmp1045, v435, cmp1048, v436, tobool1052, v437, result_symbol1054, v438, mark_end1055, v439, v440, v441, cmp1056, v442, cmp1060, v443, cmp1063, v444, cmp1066, v445, cmp1069, v446, cmp1072, v447, cmp1075, v448, cmp1078, v449, cmp1081, v450, cmp1084, v451, tobool1088, v452, result_symbol1090, v453, mark_end1091, v454, v455, v456, cmp1092, v457, cmp1096, v458, cmp1099, v459, cmp1102, v460, cmp1105, v461, cmp1108, v462, cmp1111, v463, cmp1114, v464, cmp1117, v465, cmp1120, v466, tobool1124, v467, result_symbol1126, v468, mark_end1127, v469, v470, v471, cmp1128, v472, cmp1132, v473, cmp1135, v474, cmp1138, v475, cmp1141, v476, cmp1144, v477, cmp1147, v478, cmp1150, v479, cmp1153, v480, cmp1156, v481, tobool1160, v482, result_symbol1162, v483, mark_end1163, v484, v485, v486, cmp1164, v487, cmp1168, v488, cmp1171, v489, cmp1174, v490, cmp1177, v491, cmp1180, v492, cmp1183, v493, cmp1186, v494, cmp1189, v495, cmp1192, v496, tobool1196, v497, result_symbol1198, v498, mark_end1199, v499, v500, v501, cmp1200, v502, cmp1204, v503, cmp1207, v504, cmp1210, v505, cmp1213, v506, cmp1216, v507, cmp1219, v508, cmp1222, v509, cmp1225, v510, cmp1228, v511, tobool1232, v512, result_symbol1234, v513, mark_end1235, v514, v515, v516, cmp1236, v517, cmp1240, v518, cmp1243, v519, cmp1246, v520, cmp1249, v521, cmp1252, v522, cmp1255, v523, cmp1258, v524, cmp1261, v525, cmp1264, v526, tobool1268, v527, result_symbol1270, v528, mark_end1271, v529, v530, v531, cmp1272, v532, cmp1276, v533, cmp1279, v534, cmp1282, v535, cmp1285, v536, cmp1288, v537, cmp1291, v538, cmp1294, v539, cmp1297, v540, cmp1300, v541, tobool1304, v542, result_symbol1306, v543, mark_end1307, v544, v545, v546, cmp1308, v547, cmp1312, v548, cmp1315, v549, cmp1318, v550, cmp1321, v551, cmp1324, v552, cmp1327, v553, cmp1330, v554, cmp1333, v555, cmp1336, v556, tobool1340, v557, result_symbol1342, v558, mark_end1343, v559, v560, v561, cmp1344, v562, cmp1348, v563, cmp1351, v564, cmp1354, v565, cmp1357, v566, cmp1360, v567, cmp1363, v568, cmp1366, v569, cmp1369, v570, cmp1372, v571, tobool1376, v572, result_symbol1378, v573, mark_end1379, v574, v575, v576, cmp1380, v577, cmp1384, v578, cmp1387, v579, cmp1390, v580, cmp1393, v581, cmp1396, v582, cmp1399, v583, cmp1402, v584, cmp1405, v585, cmp1408, v586, tobool1412, v587, result_symbol1414, v588, mark_end1415, v589, v590, v591, cmp1416, v592, cmp1420, v593, cmp1423, v594, cmp1426, v595, cmp1429, v596, cmp1432, v597, cmp1435, v598, cmp1438, v599, cmp1441, v600, cmp1444, v601, tobool1448, v602, result_symbol1450, v603, mark_end1451, v604, v605, v606, cmp1452, v607, cmp1456, v608, cmp1459, v609, cmp1462, v610, cmp1465, v611, cmp1468, v612, cmp1471, v613, cmp1474, v614, cmp1477, v615, cmp1480, v616, tobool1484, v617, result_symbol1486, v618, mark_end1487, v619, v620, v621, cmp1488, v622, cmp1492, v623, cmp1495, v624, cmp1498, v625, cmp1501, v626, cmp1504, v627, cmp1507, v628, cmp1510, v629, cmp1513, v630, cmp1516, v631, tobool1520, v632, result_symbol1522, v633, mark_end1523, v634, v635, v636, cmp1524, v637, cmp1528, v638, cmp1531, v639, cmp1534, v640, cmp1537, v641, cmp1540, v642, cmp1543, v643, cmp1546, v644, cmp1549, v645, cmp1552, v646, tobool1556, v647, result_symbol1558, v648, mark_end1559, v649, v650, v651, cmp1560, v652, cmp1564, v653, cmp1567, v654, cmp1570, v655, cmp1573, v656, cmp1576, v657, cmp1579, v658, cmp1582, v659, cmp1585, v660, cmp1588, v661, tobool1592, v662, result_symbol1594, v663, mark_end1595, v664, v665, v666, cmp1596, v667, cmp1600, v668, cmp1603, v669, cmp1606, v670, cmp1609, v671, cmp1612, v672, cmp1615, v673, cmp1618, v674, cmp1621, v675, cmp1624, v676, tobool1628, v677, result_symbol1630, v678, mark_end1631, v679, v680, v681, cmp1632, v682, cmp1636, v683, cmp1639, v684, cmp1642, v685, cmp1645, v686, cmp1648, v687, cmp1651, v688, cmp1654, v689, cmp1657, v690, cmp1660, v691, tobool1664, v692, result_symbol1666, v693, mark_end1667, v694, v695, v696, cmp1668, v697, cmp1672, v698, cmp1675, v699, cmp1678, v700, cmp1681, v701, cmp1684, v702, cmp1687, v703, cmp1690, v704, cmp1693, v705, cmp1696, v706, tobool1700, v707, result_symbol1702, v708, mark_end1703, v709, v710, v711, cmp1704, v712, cmp1708, v713, cmp1711, v714, cmp1714, v715, cmp1717, v716, cmp1720, v717, cmp1723, v718, cmp1726, v719, cmp1729, v720, cmp1732, v721, tobool1736, v722, result_symbol1738, v723, mark_end1739, v724, v725, v726, cmp1740, v727, cmp1744, v728, cmp1747, v729, cmp1750, v730, cmp1753, v731, cmp1756, v732, cmp1759, v733, cmp1762, v734, cmp1765, v735, cmp1768, v736, tobool1772, v737, result_symbol1774, v738, mark_end1775, v739, v740, v741, cmp1776, v742, cmp1780, v743, cmp1783, v744, cmp1786, v745, cmp1789, v746, cmp1792, v747, cmp1795, v748, cmp1798, v749, cmp1801, v750, cmp1804, v751, tobool1808, v752, result_symbol1810, v753, mark_end1811, v754, v755, v756, cmp1812, v757, cmp1816, v758, cmp1819, v759, cmp1822, v760, cmp1825, v761, cmp1828, v762, cmp1831, v763, cmp1834, v764, cmp1837, v765, cmp1840, v766, tobool1844, v767, result_symbol1846, v768, mark_end1847, v769, v770, v771, cmp1848, v772, cmp1852, v773, cmp1855, v774, cmp1858, v775, cmp1861, v776, cmp1864, v777, cmp1867, v778, cmp1870, v779, cmp1873, v780, cmp1876, v781, tobool1880, v782, result_symbol1882, v783, mark_end1883, v784, v785, v786, cmp1884, v787, cmp1888, v788, cmp1891, v789, cmp1894, v790, cmp1897, v791, cmp1900, v792, cmp1903, v793, cmp1906, v794, cmp1909, v795, cmp1912, v796, tobool1916, v797, result_symbol1918, v798, mark_end1919, v799, v800, v801, cmp1920, v802, cmp1924, v803, cmp1927, v804, cmp1930, v805, cmp1933, v806, cmp1936, v807, cmp1939, v808, cmp1942, v809, cmp1945, v810, cmp1948, v811, tobool1952, v812, result_symbol1954, v813, mark_end1955, v814, v815, v816, cmp1956, v817, cmp1960, v818, cmp1963, v819, cmp1966, v820, cmp1969, v821, cmp1972, v822, cmp1975, v823, cmp1978, v824, cmp1981, v825, cmp1984, v826, tobool1988, v827, result_symbol1990, v828, mark_end1991, v829, v830, v831, cmp1992, v832, cmp1996, v833, cmp1999, v834, cmp2002, v835, cmp2005, v836, cmp2008, v837, cmp2011, v838, cmp2014, v839, cmp2017, v840, cmp2020, v841, tobool2024, v842, result_symbol2026, v843, mark_end2027, v844, v845, v846, cmp2028, v847, cmp2032, v848, cmp2035, v849, cmp2038, v850, cmp2041, v851, cmp2044, v852, cmp2047, v853, cmp2050, v854, cmp2053, v855, cmp2056, v856, tobool2060, v857, result_symbol2062, v858, mark_end2063, v859, v860, v861, cmp2064, v862, cmp2068, v863, cmp2071, v864, cmp2074, v865, cmp2077, v866, cmp2080, v867, cmp2083, v868, cmp2086, v869, cmp2089, v870, cmp2092, v871, tobool2096, v872, result_symbol2098, v873, mark_end2099, v874, v875, v876, cmp2100, v877, cmp2104, v878, cmp2107, v879, cmp2110, v880, cmp2113, v881, cmp2116, v882, cmp2119, v883, cmp2122, v884, cmp2125, v885, cmp2128, v886, tobool2132, v887, result_symbol2134, v888, mark_end2135, v889, v890, v891, cmp2136, v892, cmp2139, v893, cmp2142, v894, cmp2145, v895, cmp2148, v896, cmp2151, v897, cmp2154, v898, cmp2157, v899, cmp2160, v900, tobool2164, v901, result_symbol2166, v902, mark_end2167, v903, v904, v905, tobool2168, v906, result_symbol2170, v907, mark_end2171, v908, v909, v910, tobool2172, v911, result_symbol2174, v912, mark_end2175, v913, v914, v915, cmp2176, v916, tobool2180, v917

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
	i = new(int32)
	i72 = new(int32)
	*lexer_addr = lexer
	*state_addr = state
	*result = 0
	*skip = 0
	*eof = 0
	goto start

next_state:
	v0 = *lexer_addr
	local_advance = &v0.F2
	v1 = *local_advance
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
	if call { frombool = 1 } else { frombool = 0 }
	*eof = frombool
	v9 = *state_addr
	conv = int32(uint32(uint16(v9)))
	switch conv {
	case 0:
		goto sw_bb
	case 1:
		goto sw_bb41
	case 2:
		goto sw_bb47
	case 3:
		goto sw_bb53
	case 4:
		goto sw_bb59
	case 5:
		goto sw_bb65
	case 6:
		goto sw_bb71
	case 7:
		goto sw_bb92
	case 8:
		goto sw_bb98
	case 9:
		goto sw_bb104
	case 10:
		goto sw_bb110
	case 11:
		goto sw_bb116
	case 12:
		goto sw_bb122
	case 13:
		goto sw_bb148
	case 14:
		goto sw_bb154
	case 15:
		goto sw_bb160
	case 16:
		goto sw_bb166
	case 17:
		goto sw_bb172
	case 18:
		goto sw_bb178
	case 19:
		goto sw_bb184
	case 20:
		goto sw_bb190
	case 21:
		goto sw_bb196
	case 22:
		goto sw_bb202
	case 23:
		goto sw_bb208
	case 24:
		goto sw_bb214
	case 25:
		goto sw_bb220
	case 26:
		goto sw_bb226
	case 27:
		goto sw_bb232
	case 28:
		goto sw_bb238
	case 29:
		goto sw_bb244
	case 30:
		goto sw_bb250
	case 31:
		goto sw_bb256
	case 32:
		goto sw_bb262
	case 33:
		goto sw_bb268
	case 34:
		goto sw_bb274
	case 35:
		goto sw_bb280
	case 36:
		goto sw_bb286
	case 37:
		goto sw_bb292
	case 38:
		goto sw_bb298
	case 39:
		goto sw_bb304
	case 40:
		goto sw_bb310
	case 41:
		goto sw_bb316
	case 42:
		goto sw_bb322
	case 43:
		goto sw_bb328
	case 44:
		goto sw_bb334
	case 45:
		goto sw_bb340
	case 46:
		goto sw_bb346
	case 47:
		goto sw_bb352
	case 48:
		goto sw_bb358
	case 49:
		goto sw_bb364
	case 50:
		goto sw_bb370
	case 51:
		goto sw_bb376
	case 52:
		goto sw_bb382
	case 53:
		goto sw_bb388
	case 54:
		goto sw_bb394
	case 55:
		goto sw_bb400
	case 56:
		goto sw_bb406
	case 57:
		goto sw_bb412
	case 58:
		goto sw_bb418
	case 59:
		goto sw_bb424
	case 60:
		goto sw_bb430
	case 61:
		goto sw_bb436
	case 62:
		goto sw_bb442
	case 63:
		goto sw_bb448
	case 64:
		goto sw_bb454
	case 65:
		goto sw_bb460
	case 66:
		goto sw_bb466
	case 67:
		goto sw_bb472
	case 68:
		goto sw_bb478
	case 69:
		goto sw_bb484
	case 70:
		goto sw_bb490
	case 71:
		goto sw_bb496
	case 72:
		goto sw_bb502
	case 73:
		goto sw_bb508
	case 74:
		goto sw_bb514
	case 75:
		goto sw_bb520
	case 76:
		goto sw_bb526
	case 77:
		goto sw_bb556
	case 78:
		goto sw_bb573
	case 79:
		goto sw_bb575
	case 80:
		goto sw_bb589
	case 81:
		goto sw_bb593
	case 82:
		goto sw_bb597
	case 83:
		goto sw_bb601
	case 84:
		goto sw_bb605
	case 85:
		goto sw_bb609
	case 86:
		goto sw_bb613
	case 87:
		goto sw_bb617
	case 88:
		goto sw_bb621
	case 89:
		goto sw_bb657
	case 90:
		goto sw_bb693
	case 91:
		goto sw_bb729
	case 92:
		goto sw_bb765
	case 93:
		goto sw_bb801
	case 94:
		goto sw_bb837
	case 95:
		goto sw_bb873
	case 96:
		goto sw_bb909
	case 97:
		goto sw_bb945
	case 98:
		goto sw_bb981
	case 99:
		goto sw_bb1017
	case 100:
		goto sw_bb1053
	case 101:
		goto sw_bb1089
	case 102:
		goto sw_bb1125
	case 103:
		goto sw_bb1161
	case 104:
		goto sw_bb1197
	case 105:
		goto sw_bb1233
	case 106:
		goto sw_bb1269
	case 107:
		goto sw_bb1305
	case 108:
		goto sw_bb1341
	case 109:
		goto sw_bb1377
	case 110:
		goto sw_bb1413
	case 111:
		goto sw_bb1449
	case 112:
		goto sw_bb1485
	case 113:
		goto sw_bb1521
	case 114:
		goto sw_bb1557
	case 115:
		goto sw_bb1593
	case 116:
		goto sw_bb1629
	case 117:
		goto sw_bb1665
	case 118:
		goto sw_bb1701
	case 119:
		goto sw_bb1737
	case 120:
		goto sw_bb1773
	case 121:
		goto sw_bb1809
	case 122:
		goto sw_bb1845
	case 123:
		goto sw_bb1881
	case 124:
		goto sw_bb1917
	case 125:
		goto sw_bb1953
	case 126:
		goto sw_bb1989
	case 127:
		goto sw_bb2025
	case 128:
		goto sw_bb2061
	case 129:
		goto sw_bb2097
	case 130:
		goto sw_bb2133
	case 131:
		goto sw_bb2165
	case 132:
		goto sw_bb2169
	case 133:
		goto sw_bb2173
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
	*state_addr = 78
	goto next_state

if_end:
	*i = 0
	goto for_cond

for_cond:
	v11 = *i
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(30)
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
	cmp14 = v18 == 45
	if cmp14 {
		goto if_then38
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v19 = *lookahead
	cmp16 = v19 == 46
	if cmp16 {
		goto if_then38
	} else {
		goto lor_lhs_false18
	}

lor_lhs_false18:
	v20 = *lookahead
	cmp19 = 48 <= v20
	if cmp19 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false23
	}

land_lhs_true:
	v21 = *lookahead
	cmp21 = v21 <= 57
	if cmp21 {
		goto if_then38
	} else {
		goto lor_lhs_false23
	}

lor_lhs_false23:
	v22 = *lookahead
	cmp24 = 65 <= v22
	if cmp24 {
		goto land_lhs_true26
	} else {
		goto lor_lhs_false29
	}

land_lhs_true26:
	v23 = *lookahead
	cmp27 = v23 <= 90
	if cmp27 {
		goto if_then38
	} else {
		goto lor_lhs_false29
	}

lor_lhs_false29:
	v24 = *lookahead
	cmp30 = v24 == 95
	if cmp30 {
		goto if_then38
	} else {
		goto lor_lhs_false32
	}

lor_lhs_false32:
	v25 = *lookahead
	cmp33 = 98 <= v25
	if cmp33 {
		goto land_lhs_true35
	} else {
		goto if_end39
	}

land_lhs_true35:
	v26 = *lookahead
	cmp36 = v26 <= 122
	if cmp36 {
		goto if_then38
	} else {
		goto if_end39
	}

if_then38:
	*state_addr = 130
	goto next_state

if_end39:
	v27 = *result
	tobool40 = (v27 & 1) != 0
	*retval = tobool40
	goto _return

sw_bb41:
	v28 = *lookahead
	cmp42 = v28 == 45
	if cmp42 {
		goto if_then44
	} else {
		goto if_end45
	}

if_then44:
	*state_addr = 23
	goto next_state

if_end45:
	v29 = *result
	tobool46 = (v29 & 1) != 0
	*retval = tobool46
	goto _return

sw_bb47:
	v30 = *lookahead
	cmp48 = v30 == 97
	if cmp48 {
		goto if_then50
	} else {
		goto if_end51
	}

if_then50:
	*state_addr = 30
	goto next_state

if_end51:
	v31 = *result
	tobool52 = (v31 & 1) != 0
	*retval = tobool52
	goto _return

sw_bb53:
	v32 = *lookahead
	cmp54 = v32 == 97
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*state_addr = 40
	goto next_state

if_end57:
	v33 = *result
	tobool58 = (v33 & 1) != 0
	*retval = tobool58
	goto _return

sw_bb59:
	v34 = *lookahead
	cmp60 = v34 == 97
	if cmp60 {
		goto if_then62
	} else {
		goto if_end63
	}

if_then62:
	*state_addr = 13
	goto next_state

if_end63:
	v35 = *result
	tobool64 = (v35 & 1) != 0
	*retval = tobool64
	goto _return

sw_bb65:
	v36 = *lookahead
	cmp66 = v36 == 97
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*state_addr = 68
	goto next_state

if_end69:
	v37 = *result
	tobool70 = (v37 & 1) != 0
	*retval = tobool70
	goto _return

sw_bb71:
	*i72 = 0
	goto for_cond73

for_cond73:
	v38 = *i72
	conv74 = int64(uint64(uint32(v38)))
	cmp75 = uint64(conv74) < uint64(20)
	if cmp75 {
		goto for_body77
	} else {
		goto for_end90
	}

for_body77:
	v39 = *i72
	idxprom78 = int64(uint64(uint32(v39)))
	arrayidx79 = &ts_lex_map_31[idxprom78]
	v40 = *arrayidx79
	conv80 = int32(uint32(uint16(v40)))
	v41 = *lookahead
	cmp81 = conv80 == v41
	if cmp81 {
		goto if_then83
	} else {
		goto if_end87
	}

if_then83:
	v42 = *i72
	add84 = v42 + 1
	idxprom85 = int64(uint64(uint32(add84)))
	arrayidx86 = &ts_lex_map_31[idxprom85]
	v43 = *arrayidx86
	*state_addr = v43
	goto next_state

if_end87:
	goto for_inc88

for_inc88:
	v44 = *i72
	add89 = v44 + 2
	*i72 = add89
	goto for_cond73

for_end90:
	v45 = *result
	tobool91 = (v45 & 1) != 0
	*retval = tobool91
	goto _return

sw_bb92:
	v46 = *lookahead
	cmp93 = v46 == 97
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*state_addr = 26
	goto next_state

if_end96:
	v47 = *result
	tobool97 = (v47 & 1) != 0
	*retval = tobool97
	goto _return

sw_bb98:
	v48 = *lookahead
	cmp99 = v48 == 97
	if cmp99 {
		goto if_then101
	} else {
		goto if_end102
	}

if_then101:
	*state_addr = 27
	goto next_state

if_end102:
	v49 = *result
	tobool103 = (v49 & 1) != 0
	*retval = tobool103
	goto _return

sw_bb104:
	v50 = *lookahead
	cmp105 = v50 == 97
	if cmp105 {
		goto if_then107
	} else {
		goto if_end108
	}

if_then107:
	*state_addr = 61
	goto next_state

if_end108:
	v51 = *result
	tobool109 = (v51 & 1) != 0
	*retval = tobool109
	goto _return

sw_bb110:
	v52 = *lookahead
	cmp111 = v52 == 97
	if cmp111 {
		goto if_then113
	} else {
		goto if_end114
	}

if_then113:
	*state_addr = 65
	goto next_state

if_end114:
	v53 = *result
	tobool115 = (v53 & 1) != 0
	*retval = tobool115
	goto _return

sw_bb116:
	v54 = *lookahead
	cmp117 = v54 == 98
	if cmp117 {
		goto if_then119
	} else {
		goto if_end120
	}

if_then119:
	*state_addr = 66
	goto next_state

if_end120:
	v55 = *result
	tobool121 = (v55 & 1) != 0
	*retval = tobool121
	goto _return

sw_bb122:
	v56 = *lookahead
	cmp123 = v56 == 99
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*state_addr = 63
	goto next_state

if_end126:
	v57 = *lookahead
	cmp127 = v57 == 101
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*state_addr = 58
	goto next_state

if_end130:
	v58 = *lookahead
	cmp131 = v58 == 102
	if cmp131 {
		goto if_then133
	} else {
		goto if_end134
	}

if_then133:
	*state_addr = 2
	goto next_state

if_end134:
	v59 = *lookahead
	cmp135 = v59 == 108
	if cmp135 {
		goto if_then137
	} else {
		goto if_end138
	}

if_then137:
	*state_addr = 3
	goto next_state

if_end138:
	v60 = *lookahead
	cmp139 = v60 == 112
	if cmp139 {
		goto if_then141
	} else {
		goto if_end142
	}

if_then141:
	*state_addr = 37
	goto next_state

if_end142:
	v61 = *lookahead
	cmp143 = v61 == 115
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*state_addr = 34
	goto next_state

if_end146:
	v62 = *result
	tobool147 = (v62 & 1) != 0
	*retval = tobool147
	goto _return

sw_bb148:
	v63 = *lookahead
	cmp149 = v63 == 99
	if cmp149 {
		goto if_then151
	} else {
		goto if_end152
	}

if_then151:
	*state_addr = 47
	goto next_state

if_end152:
	v64 = *result
	tobool153 = (v64 & 1) != 0
	*retval = tobool153
	goto _return

sw_bb154:
	v65 = *lookahead
	cmp155 = v65 == 100
	if cmp155 {
		goto if_then157
	} else {
		goto if_end158
	}

if_then157:
	*state_addr = 131
	goto next_state

if_end158:
	v66 = *result
	tobool159 = (v66 & 1) != 0
	*retval = tobool159
	goto _return

sw_bb160:
	v67 = *lookahead
	cmp161 = v67 == 100
	if cmp161 {
		goto if_then163
	} else {
		goto if_end164
	}

if_then163:
	*state_addr = 46
	goto next_state

if_end164:
	v68 = *result
	tobool165 = (v68 & 1) != 0
	*retval = tobool165
	goto _return

sw_bb166:
	v69 = *lookahead
	cmp167 = v69 == 100
	if cmp167 {
		goto if_then169
	} else {
		goto if_end170
	}

if_then169:
	*state_addr = 60
	goto next_state

if_end170:
	v70 = *result
	tobool171 = (v70 & 1) != 0
	*retval = tobool171
	goto _return

sw_bb172:
	v71 = *lookahead
	cmp173 = v71 == 101
	if cmp173 {
		goto if_then175
	} else {
		goto if_end176
	}

if_then175:
	*state_addr = 84
	goto next_state

if_end176:
	v72 = *result
	tobool177 = (v72 & 1) != 0
	*retval = tobool177
	goto _return

sw_bb178:
	v73 = *lookahead
	cmp179 = v73 == 101
	if cmp179 {
		goto if_then181
	} else {
		goto if_end182
	}

if_then181:
	*state_addr = 11
	goto next_state

if_end182:
	v74 = *result
	tobool183 = (v74 & 1) != 0
	*retval = tobool183
	goto _return

sw_bb184:
	v75 = *lookahead
	cmp185 = v75 == 101
	if cmp185 {
		goto if_then187
	} else {
		goto if_end188
	}

if_then187:
	*state_addr = 70
	goto next_state

if_end188:
	v76 = *result
	tobool189 = (v76 & 1) != 0
	*retval = tobool189
	goto _return

sw_bb190:
	v77 = *lookahead
	cmp191 = v77 == 101
	if cmp191 {
		goto if_then193
	} else {
		goto if_end194
	}

if_then193:
	*state_addr = 42
	goto next_state

if_end194:
	v78 = *result
	tobool195 = (v78 & 1) != 0
	*retval = tobool195
	goto _return

sw_bb196:
	v79 = *lookahead
	cmp197 = v79 == 101
	if cmp197 {
		goto if_then199
	} else {
		goto if_end200
	}

if_then199:
	*state_addr = 18
	goto next_state

if_end200:
	v80 = *result
	tobool201 = (v80 & 1) != 0
	*retval = tobool201
	goto _return

sw_bb202:
	v81 = *lookahead
	cmp203 = v81 == 102
	if cmp203 {
		goto if_then205
	} else {
		goto if_end206
	}

if_then205:
	*state_addr = 36
	goto next_state

if_end206:
	v82 = *result
	tobool207 = (v82 & 1) != 0
	*retval = tobool207
	goto _return

sw_bb208:
	v83 = *lookahead
	cmp209 = v83 == 102
	if cmp209 {
		goto if_then211
	} else {
		goto if_end212
	}

if_then211:
	*state_addr = 10
	goto next_state

if_end212:
	v84 = *result
	tobool213 = (v84 & 1) != 0
	*retval = tobool213
	goto _return

sw_bb214:
	v85 = *lookahead
	cmp215 = v85 == 102
	if cmp215 {
		goto if_then217
	} else {
		goto if_end218
	}

if_then217:
	*state_addr = 49
	goto next_state

if_end218:
	v86 = *result
	tobool219 = (v86 & 1) != 0
	*retval = tobool219
	goto _return

sw_bb220:
	v87 = *lookahead
	cmp221 = v87 == 103
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*state_addr = 72
	goto next_state

if_end224:
	v88 = *result
	tobool225 = (v88 & 1) != 0
	*retval = tobool225
	goto _return

sw_bb226:
	v89 = *lookahead
	cmp227 = v89 == 103
	if cmp227 {
		goto if_then229
	} else {
		goto if_end230
	}

if_then229:
	*state_addr = 17
	goto next_state

if_end230:
	v90 = *result
	tobool231 = (v90 & 1) != 0
	*retval = tobool231
	goto _return

sw_bb232:
	v91 = *lookahead
	cmp233 = v91 == 103
	if cmp233 {
		goto if_then235
	} else {
		goto if_end236
	}

if_then235:
	*state_addr = 51
	goto next_state

if_end236:
	v92 = *result
	tobool237 = (v92 & 1) != 0
	*retval = tobool237
	goto _return

sw_bb238:
	v93 = *lookahead
	cmp239 = v93 == 105
	if cmp239 {
		goto if_then241
	} else {
		goto if_end242
	}

if_then241:
	*state_addr = 53
	goto next_state

if_end242:
	v94 = *result
	tobool243 = (v94 & 1) != 0
	*retval = tobool243
	goto _return

sw_bb244:
	v95 = *lookahead
	cmp245 = v95 == 105
	if cmp245 {
		goto if_then247
	} else {
		goto if_end248
	}

if_then247:
	*state_addr = 64
	goto next_state

if_end248:
	v96 = *result
	tobool249 = (v96 & 1) != 0
	*retval = tobool249
	goto _return

sw_bb250:
	v97 = *lookahead
	cmp251 = v97 == 105
	if cmp251 {
		goto if_then253
	} else {
		goto if_end254
	}

if_then253:
	*state_addr = 35
	goto next_state

if_end254:
	v98 = *result
	tobool255 = (v98 & 1) != 0
	*retval = tobool255
	goto _return

sw_bb256:
	v99 = *lookahead
	cmp257 = v99 == 105
	if cmp257 {
		goto if_then259
	} else {
		goto if_end260
	}

if_then259:
	*state_addr = 43
	goto next_state

if_end260:
	v100 = *result
	tobool261 = (v100 & 1) != 0
	*retval = tobool261
	goto _return

sw_bb262:
	v101 = *lookahead
	cmp263 = v101 == 105
	if cmp263 {
		goto if_then265
	} else {
		goto if_end266
	}

if_then265:
	*state_addr = 14
	goto next_state

if_end266:
	v102 = *result
	tobool267 = (v102 & 1) != 0
	*retval = tobool267
	goto _return

sw_bb268:
	v103 = *lookahead
	cmp269 = v103 == 105
	if cmp269 {
		goto if_then271
	} else {
		goto if_end272
	}

if_then271:
	*state_addr = 44
	goto next_state

if_end272:
	v104 = *result
	tobool273 = (v104 & 1) != 0
	*retval = tobool273
	goto _return

sw_bb274:
	v105 = *lookahead
	cmp275 = v105 == 107
	if cmp275 {
		goto if_then277
	} else {
		goto if_end278
	}

if_then277:
	*state_addr = 28
	goto next_state

if_end278:
	v106 = *result
	tobool279 = (v106 & 1) != 0
	*retval = tobool279
	goto _return

sw_bb280:
	v107 = *lookahead
	cmp281 = v107 == 108
	if cmp281 {
		goto if_then283
	} else {
		goto if_end284
	}

if_then283:
	*state_addr = 1
	goto next_state

if_end284:
	v108 = *result
	tobool285 = (v108 & 1) != 0
	*retval = tobool285
	goto _return

sw_bb286:
	v109 = *lookahead
	cmp287 = v109 == 108
	if cmp287 {
		goto if_then289
	} else {
		goto if_end290
	}

if_then289:
	*state_addr = 75
	goto next_state

if_end290:
	v110 = *result
	tobool291 = (v110 & 1) != 0
	*retval = tobool291
	goto _return

sw_bb292:
	v111 = *lookahead
	cmp293 = v111 == 108
	if cmp293 {
		goto if_then295
	} else {
		goto if_end296
	}

if_then295:
	*state_addr = 5
	goto next_state

if_end296:
	v112 = *result
	tobool297 = (v112 & 1) != 0
	*retval = tobool297
	goto _return

sw_bb298:
	v113 = *lookahead
	cmp299 = v113 == 108
	if cmp299 {
		goto if_then301
	} else {
		goto if_end302
	}

if_then301:
	*state_addr = 9
	goto next_state

if_end302:
	v114 = *result
	tobool303 = (v114 & 1) != 0
	*retval = tobool303
	goto _return

sw_bb304:
	v115 = *lookahead
	cmp305 = v115 == 109
	if cmp305 {
		goto if_then307
	} else {
		goto if_end308
	}

if_then307:
	*state_addr = 87
	goto next_state

if_end308:
	v116 = *result
	tobool309 = (v116 & 1) != 0
	*retval = tobool309
	goto _return

sw_bb310:
	v117 = *lookahead
	cmp311 = v117 == 110
	if cmp311 {
		goto if_then313
	} else {
		goto if_end314
	}

if_then313:
	*state_addr = 25
	goto next_state

if_end314:
	v118 = *result
	tobool315 = (v118 & 1) != 0
	*retval = tobool315
	goto _return

sw_bb316:
	v119 = *lookahead
	cmp317 = v119 == 110
	if cmp317 {
		goto if_then319
	} else {
		goto if_end320
	}

if_then319:
	*state_addr = 16
	goto next_state

if_end320:
	v120 = *result
	tobool321 = (v120 & 1) != 0
	*retval = tobool321
	goto _return

sw_bb322:
	v121 = *lookahead
	cmp323 = v121 == 110
	if cmp323 {
		goto if_then325
	} else {
		goto if_end326
	}

if_then325:
	*state_addr = 11
	goto next_state

if_end326:
	v122 = *result
	tobool327 = (v122 & 1) != 0
	*retval = tobool327
	goto _return

sw_bb328:
	v123 = *lookahead
	cmp329 = v123 == 110
	if cmp329 {
		goto if_then331
	} else {
		goto if_end332
	}

if_then331:
	*state_addr = 71
	goto next_state

if_end332:
	v124 = *result
	tobool333 = (v124 & 1) != 0
	*retval = tobool333
	goto _return

sw_bb334:
	v125 = *lookahead
	cmp335 = v125 == 110
	if cmp335 {
		goto if_then337
	} else {
		goto if_end338
	}

if_then337:
	*state_addr = 15
	goto next_state

if_end338:
	v126 = *result
	tobool339 = (v126 & 1) != 0
	*retval = tobool339
	goto _return

sw_bb340:
	v127 = *lookahead
	cmp341 = v127 == 110
	if cmp341 {
		goto if_then343
	} else {
		goto if_end344
	}

if_then343:
	*state_addr = 22
	goto next_state

if_end344:
	v128 = *result
	tobool345 = (v128 & 1) != 0
	*retval = tobool345
	goto _return

sw_bb346:
	v129 = *lookahead
	cmp347 = v129 == 111
	if cmp347 {
		goto if_then349
	} else {
		goto if_end350
	}

if_then349:
	*state_addr = 73
	goto next_state

if_end350:
	v130 = *result
	tobool351 = (v130 & 1) != 0
	*retval = tobool351
	goto _return

sw_bb352:
	v131 = *lookahead
	cmp353 = v131 == 111
	if cmp353 {
		goto if_then355
	} else {
		goto if_end356
	}

if_then355:
	*state_addr = 64
	goto next_state

if_end356:
	v132 = *result
	tobool357 = (v132 & 1) != 0
	*retval = tobool357
	goto _return

sw_bb358:
	v133 = *lookahead
	cmp359 = v133 == 111
	if cmp359 {
		goto if_then361
	} else {
		goto if_end362
	}

if_then361:
	*state_addr = 56
	goto next_state

if_end362:
	v134 = *result
	tobool363 = (v134 & 1) != 0
	*retval = tobool363
	goto _return

sw_bb364:
	v135 = *lookahead
	cmp365 = v135 == 111
	if cmp365 {
		goto if_then367
	} else {
		goto if_end368
	}

if_then367:
	*state_addr = 57
	goto next_state

if_end368:
	v136 = *result
	tobool369 = (v136 & 1) != 0
	*retval = tobool369
	goto _return

sw_bb370:
	v137 = *lookahead
	cmp371 = v137 == 111
	if cmp371 {
		goto if_then373
	} else {
		goto if_end374
	}

if_then373:
	*state_addr = 32
	goto next_state

if_end374:
	v138 = *result
	tobool375 = (v138 & 1) != 0
	*retval = tobool375
	goto _return

sw_bb376:
	v139 = *lookahead
	cmp377 = v139 == 111
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*state_addr = 45
	goto next_state

if_end380:
	v140 = *result
	tobool381 = (v140 & 1) != 0
	*retval = tobool381
	goto _return

sw_bb382:
	v141 = *lookahead
	cmp383 = v141 == 111
	if cmp383 {
		goto if_then385
	} else {
		goto if_end386
	}

if_then385:
	*state_addr = 38
	goto next_state

if_end386:
	v142 = *result
	tobool387 = (v142 & 1) != 0
	*retval = tobool387
	goto _return

sw_bb388:
	v143 = *lookahead
	cmp389 = v143 == 112
	if cmp389 {
		goto if_then391
	} else {
		goto if_end392
	}

if_then391:
	*state_addr = 83
	goto next_state

if_end392:
	v144 = *result
	tobool393 = (v144 & 1) != 0
	*retval = tobool393
	goto _return

sw_bb394:
	v145 = *lookahead
	cmp395 = v145 == 112
	if cmp395 {
		goto if_then397
	} else {
		goto if_end398
	}

if_then397:
	*state_addr = 20
	goto next_state

if_end398:
	v146 = *result
	tobool399 = (v146 & 1) != 0
	*retval = tobool399
	goto _return

sw_bb400:
	v147 = *lookahead
	cmp401 = v147 == 114
	if cmp401 {
		goto if_then403
	} else {
		goto if_end404
	}

if_then403:
	*state_addr = 48
	goto next_state

if_end404:
	v148 = *result
	tobool405 = (v148 & 1) != 0
	*retval = tobool405
	goto _return

sw_bb406:
	v149 = *lookahead
	cmp407 = v149 == 114
	if cmp407 {
		goto if_then409
	} else {
		goto if_end410
	}

if_then409:
	*state_addr = 81
	goto next_state

if_end410:
	v150 = *result
	tobool411 = (v150 & 1) != 0
	*retval = tobool411
	goto _return

sw_bb412:
	v151 = *lookahead
	cmp413 = v151 == 114
	if cmp413 {
		goto if_then415
	} else {
		goto if_end416
	}

if_then415:
	*state_addr = 39
	goto next_state

if_end416:
	v152 = *result
	tobool417 = (v152 & 1) != 0
	*retval = tobool417
	goto _return

sw_bb418:
	v153 = *lookahead
	cmp419 = v153 == 114
	if cmp419 {
		goto if_then421
	} else {
		goto if_end422
	}

if_then421:
	*state_addr = 55
	goto next_state

if_end422:
	v154 = *result
	tobool423 = (v154 & 1) != 0
	*retval = tobool423
	goto _return

sw_bb424:
	v155 = *lookahead
	cmp425 = v155 == 114
	if cmp425 {
		goto if_then427
	} else {
		goto if_end428
	}

if_then427:
	*state_addr = 21
	goto next_state

if_end428:
	v156 = *result
	tobool429 = (v156 & 1) != 0
	*retval = tobool429
	goto _return

sw_bb430:
	v157 = *lookahead
	cmp431 = v157 == 114
	if cmp431 {
		goto if_then433
	} else {
		goto if_end434
	}

if_then433:
	*state_addr = 50
	goto next_state

if_end434:
	v158 = *result
	tobool435 = (v158 & 1) != 0
	*retval = tobool435
	goto _return

sw_bb436:
	v159 = *lookahead
	cmp437 = v159 == 114
	if cmp437 {
		goto if_then439
	} else {
		goto if_end440
	}

if_then439:
	*state_addr = 29
	goto next_state

if_end440:
	v160 = *result
	tobool441 = (v160 & 1) != 0
	*retval = tobool441
	goto _return

sw_bb442:
	v161 = *lookahead
	cmp443 = v161 == 114
	if cmp443 {
		goto if_then445
	} else {
		goto if_end446
	}

if_then445:
	*state_addr = 8
	goto next_state

if_end446:
	v162 = *result
	tobool447 = (v162 & 1) != 0
	*retval = tobool447
	goto _return

sw_bb448:
	v163 = *lookahead
	cmp449 = v163 == 115
	if cmp449 {
		goto if_then451
	} else {
		goto if_end452
	}

if_then451:
	*state_addr = 67
	goto next_state

if_end452:
	v164 = *result
	tobool453 = (v164 & 1) != 0
	*retval = tobool453
	goto _return

sw_bb454:
	v165 = *lookahead
	cmp455 = v165 == 115
	if cmp455 {
		goto if_then457
	} else {
		goto if_end458
	}

if_then457:
	*state_addr = 131
	goto next_state

if_end458:
	v166 = *result
	tobool459 = (v166 & 1) != 0
	*retval = tobool459
	goto _return

sw_bb460:
	v167 = *lookahead
	cmp461 = v167 == 115
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*state_addr = 69
	goto next_state

if_end464:
	v168 = *result
	tobool465 = (v168 & 1) != 0
	*retval = tobool465
	goto _return

sw_bb466:
	v169 = *lookahead
	cmp467 = v169 == 115
	if cmp467 {
		goto if_then469
	} else {
		goto if_end470
	}

if_then469:
	*state_addr = 14
	goto next_state

if_end470:
	v170 = *result
	tobool471 = (v170 & 1) != 0
	*retval = tobool471
	goto _return

sw_bb472:
	v171 = *lookahead
	cmp473 = v171 == 116
	if cmp473 {
		goto if_then475
	} else {
		goto if_end476
	}

if_then475:
	*state_addr = 80
	goto next_state

if_end476:
	v172 = *result
	tobool477 = (v172 & 1) != 0
	*retval = tobool477
	goto _return

sw_bb478:
	v173 = *lookahead
	cmp479 = v173 == 116
	if cmp479 {
		goto if_then481
	} else {
		goto if_end482
	}

if_then481:
	*state_addr = 24
	goto next_state

if_end482:
	v174 = *result
	tobool483 = (v174 & 1) != 0
	*retval = tobool483
	goto _return

sw_bb484:
	v175 = *lookahead
	cmp485 = v175 == 116
	if cmp485 {
		goto if_then487
	} else {
		goto if_end488
	}

if_then487:
	*state_addr = 82
	goto next_state

if_end488:
	v176 = *result
	tobool489 = (v176 & 1) != 0
	*retval = tobool489
	goto _return

sw_bb490:
	v177 = *lookahead
	cmp491 = v177 == 116
	if cmp491 {
		goto if_then493
	} else {
		goto if_end494
	}

if_then493:
	*state_addr = 11
	goto next_state

if_end494:
	v178 = *result
	tobool495 = (v178 & 1) != 0
	*retval = tobool495
	goto _return

sw_bb496:
	v179 = *lookahead
	cmp497 = v179 == 117
	if cmp497 {
		goto if_then499
	} else {
		goto if_end500
	}

if_then499:
	*state_addr = 74
	goto next_state

if_end500:
	v180 = *result
	tobool501 = (v180 & 1) != 0
	*retval = tobool501
	goto _return

sw_bb502:
	v181 = *lookahead
	cmp503 = v181 == 117
	if cmp503 {
		goto if_then505
	} else {
		goto if_end506
	}

if_then505:
	*state_addr = 7
	goto next_state

if_end506:
	v182 = *result
	tobool507 = (v182 & 1) != 0
	*retval = tobool507
	goto _return

sw_bb508:
	v183 = *lookahead
	cmp509 = v183 == 119
	if cmp509 {
		goto if_then511
	} else {
		goto if_end512
	}

if_then511:
	*state_addr = 64
	goto next_state

if_end512:
	v184 = *result
	tobool513 = (v184 & 1) != 0
	*retval = tobool513
	goto _return

sw_bb514:
	v185 = *lookahead
	cmp515 = v185 == 120
	if cmp515 {
		goto if_then517
	} else {
		goto if_end518
	}

if_then517:
	*state_addr = 131
	goto next_state

if_end518:
	v186 = *result
	tobool519 = (v186 & 1) != 0
	*retval = tobool519
	goto _return

sw_bb520:
	v187 = *lookahead
	cmp521 = v187 == 121
	if cmp521 {
		goto if_then523
	} else {
		goto if_end524
	}

if_then523:
	*state_addr = 131
	goto next_state

if_end524:
	v188 = *result
	tobool525 = (v188 & 1) != 0
	*retval = tobool525
	goto _return

sw_bb526:
	v189 = *lookahead
	cmp527 = v189 == 45
	if cmp527 {
		goto if_then553
	} else {
		goto lor_lhs_false529
	}

lor_lhs_false529:
	v190 = *lookahead
	cmp530 = v190 == 46
	if cmp530 {
		goto if_then553
	} else {
		goto lor_lhs_false532
	}

lor_lhs_false532:
	v191 = *lookahead
	cmp533 = 48 <= v191
	if cmp533 {
		goto land_lhs_true535
	} else {
		goto lor_lhs_false538
	}

land_lhs_true535:
	v192 = *lookahead
	cmp536 = v192 <= 57
	if cmp536 {
		goto if_then553
	} else {
		goto lor_lhs_false538
	}

lor_lhs_false538:
	v193 = *lookahead
	cmp539 = 65 <= v193
	if cmp539 {
		goto land_lhs_true541
	} else {
		goto lor_lhs_false544
	}

land_lhs_true541:
	v194 = *lookahead
	cmp542 = v194 <= 90
	if cmp542 {
		goto if_then553
	} else {
		goto lor_lhs_false544
	}

lor_lhs_false544:
	v195 = *lookahead
	cmp545 = v195 == 95
	if cmp545 {
		goto if_then553
	} else {
		goto lor_lhs_false547
	}

lor_lhs_false547:
	v196 = *lookahead
	cmp548 = 97 <= v196
	if cmp548 {
		goto land_lhs_true550
	} else {
		goto if_end554
	}

land_lhs_true550:
	v197 = *lookahead
	cmp551 = v197 <= 122
	if cmp551 {
		goto if_then553
	} else {
		goto if_end554
	}

if_then553:
	*state_addr = 130
	goto next_state

if_end554:
	v198 = *result
	tobool555 = (v198 & 1) != 0
	*retval = tobool555
	goto _return

sw_bb556:
	v199 = *eof
	tobool557 = (v199 & 1) != 0
	if tobool557 {
		goto if_then558
	} else {
		goto if_end559
	}

if_then558:
	*state_addr = 78
	goto next_state

if_end559:
	v200 = *lookahead
	cmp560 = v200 == 10
	if cmp560 {
		goto if_then562
	} else {
		goto if_end563
	}

if_then562:
	*state_addr = 132
	goto next_state

if_end563:
	v201 = *lookahead
	cmp564 = v201 == 13
	if cmp564 {
		goto if_then566
	} else {
		goto if_end567
	}

if_then566:
	*state_addr = 133
	goto next_state

if_end567:
	v202 = *lookahead
	cmp568 = v202 != 0
	if cmp568 {
		goto if_then570
	} else {
		goto if_end571
	}

if_then570:
	*state_addr = 79
	goto next_state

if_end571:
	v203 = *result
	tobool572 = (v203 & 1) != 0
	*retval = tobool572
	goto _return

sw_bb573:
	*result = 1
	v204 = *lexer_addr
	result_symbol = &v204.F1
	*result_symbol = 0
	v205 = *lexer_addr
	mark_end = &v205.F3
	v206 = *mark_end
	v207 = *lexer_addr
	v206(v207)
	v208 = *result
	tobool574 = (v208 & 1) != 0
	*retval = tobool574
	goto _return

sw_bb575:
	*result = 1
	v209 = *lexer_addr
	result_symbol576 = &v209.F1
	*result_symbol576 = 1
	v210 = *lexer_addr
	mark_end577 = &v210.F3
	v211 = *mark_end577
	v212 = *lexer_addr
	v211(v212)
	v213 = *lookahead
	cmp578 = v213 != 0
	if cmp578 {
		goto land_lhs_true580
	} else {
		goto if_end587
	}

land_lhs_true580:
	v214 = *lookahead
	cmp581 = v214 != 10
	if cmp581 {
		goto land_lhs_true583
	} else {
		goto if_end587
	}

land_lhs_true583:
	v215 = *lookahead
	cmp584 = v215 != 13
	if cmp584 {
		goto if_then586
	} else {
		goto if_end587
	}

if_then586:
	*state_addr = 79
	goto next_state

if_end587:
	v216 = *result
	tobool588 = (v216 & 1) != 0
	*retval = tobool588
	goto _return

sw_bb589:
	*result = 1
	v217 = *lexer_addr
	result_symbol590 = &v217.F1
	*result_symbol590 = 2
	v218 = *lexer_addr
	mark_end591 = &v218.F3
	v219 = *mark_end591
	v220 = *lexer_addr
	v219(v220)
	v221 = *result
	tobool592 = (v221 & 1) != 0
	*retval = tobool592
	goto _return

sw_bb593:
	*result = 1
	v222 = *lexer_addr
	result_symbol594 = &v222.F1
	*result_symbol594 = 3
	v223 = *lexer_addr
	mark_end595 = &v223.F3
	v224 = *mark_end595
	v225 = *lexer_addr
	v224(v225)
	v226 = *result
	tobool596 = (v226 & 1) != 0
	*retval = tobool596
	goto _return

sw_bb597:
	*result = 1
	v227 = *lexer_addr
	result_symbol598 = &v227.F1
	*result_symbol598 = 4
	v228 = *lexer_addr
	mark_end599 = &v228.F3
	v229 = *mark_end599
	v230 = *lexer_addr
	v229(v230)
	v231 = *result
	tobool600 = (v231 & 1) != 0
	*retval = tobool600
	goto _return

sw_bb601:
	*result = 1
	v232 = *lexer_addr
	result_symbol602 = &v232.F1
	*result_symbol602 = 5
	v233 = *lexer_addr
	mark_end603 = &v233.F3
	v234 = *mark_end603
	v235 = *lexer_addr
	v234(v235)
	v236 = *result
	tobool604 = (v236 & 1) != 0
	*retval = tobool604
	goto _return

sw_bb605:
	*result = 1
	v237 = *lexer_addr
	result_symbol606 = &v237.F1
	*result_symbol606 = 6
	v238 = *lexer_addr
	mark_end607 = &v238.F3
	v239 = *mark_end607
	v240 = *lexer_addr
	v239(v240)
	v241 = *result
	tobool608 = (v241 & 1) != 0
	*retval = tobool608
	goto _return

sw_bb609:
	*result = 1
	v242 = *lexer_addr
	result_symbol610 = &v242.F1
	*result_symbol610 = 7
	v243 = *lexer_addr
	mark_end611 = &v243.F3
	v244 = *mark_end611
	v245 = *lexer_addr
	v244(v245)
	v246 = *result
	tobool612 = (v246 & 1) != 0
	*retval = tobool612
	goto _return

sw_bb613:
	*result = 1
	v247 = *lexer_addr
	result_symbol614 = &v247.F1
	*result_symbol614 = 8
	v248 = *lexer_addr
	mark_end615 = &v248.F3
	v249 = *mark_end615
	v250 = *lexer_addr
	v249(v250)
	v251 = *result
	tobool616 = (v251 & 1) != 0
	*retval = tobool616
	goto _return

sw_bb617:
	*result = 1
	v252 = *lexer_addr
	result_symbol618 = &v252.F1
	*result_symbol618 = 9
	v253 = *lexer_addr
	mark_end619 = &v253.F3
	v254 = *mark_end619
	v255 = *lexer_addr
	v254(v255)
	v256 = *result
	tobool620 = (v256 & 1) != 0
	*retval = tobool620
	goto _return

sw_bb621:
	*result = 1
	v257 = *lexer_addr
	result_symbol622 = &v257.F1
	*result_symbol622 = 10
	v258 = *lexer_addr
	mark_end623 = &v258.F3
	v259 = *mark_end623
	v260 = *lexer_addr
	v259(v260)
	v261 = *lookahead
	cmp624 = v261 == 97
	if cmp624 {
		goto if_then626
	} else {
		goto if_end627
	}

if_then626:
	*state_addr = 92
	goto next_state

if_end627:
	v262 = *lookahead
	cmp628 = v262 == 45
	if cmp628 {
		goto if_then654
	} else {
		goto lor_lhs_false630
	}

lor_lhs_false630:
	v263 = *lookahead
	cmp631 = v263 == 46
	if cmp631 {
		goto if_then654
	} else {
		goto lor_lhs_false633
	}

lor_lhs_false633:
	v264 = *lookahead
	cmp634 = 48 <= v264
	if cmp634 {
		goto land_lhs_true636
	} else {
		goto lor_lhs_false639
	}

land_lhs_true636:
	v265 = *lookahead
	cmp637 = v265 <= 57
	if cmp637 {
		goto if_then654
	} else {
		goto lor_lhs_false639
	}

lor_lhs_false639:
	v266 = *lookahead
	cmp640 = 65 <= v266
	if cmp640 {
		goto land_lhs_true642
	} else {
		goto lor_lhs_false645
	}

land_lhs_true642:
	v267 = *lookahead
	cmp643 = v267 <= 90
	if cmp643 {
		goto if_then654
	} else {
		goto lor_lhs_false645
	}

lor_lhs_false645:
	v268 = *lookahead
	cmp646 = v268 == 95
	if cmp646 {
		goto if_then654
	} else {
		goto lor_lhs_false648
	}

lor_lhs_false648:
	v269 = *lookahead
	cmp649 = 98 <= v269
	if cmp649 {
		goto land_lhs_true651
	} else {
		goto if_end655
	}

land_lhs_true651:
	v270 = *lookahead
	cmp652 = v270 <= 122
	if cmp652 {
		goto if_then654
	} else {
		goto if_end655
	}

if_then654:
	*state_addr = 130
	goto next_state

if_end655:
	v271 = *result
	tobool656 = (v271 & 1) != 0
	*retval = tobool656
	goto _return

sw_bb657:
	*result = 1
	v272 = *lexer_addr
	result_symbol658 = &v272.F1
	*result_symbol658 = 10
	v273 = *lexer_addr
	mark_end659 = &v273.F3
	v274 = *mark_end659
	v275 = *lexer_addr
	v274(v275)
	v276 = *lookahead
	cmp660 = v276 == 97
	if cmp660 {
		goto if_then662
	} else {
		goto if_end663
	}

if_then662:
	*state_addr = 101
	goto next_state

if_end663:
	v277 = *lookahead
	cmp664 = v277 == 45
	if cmp664 {
		goto if_then690
	} else {
		goto lor_lhs_false666
	}

lor_lhs_false666:
	v278 = *lookahead
	cmp667 = v278 == 46
	if cmp667 {
		goto if_then690
	} else {
		goto lor_lhs_false669
	}

lor_lhs_false669:
	v279 = *lookahead
	cmp670 = 48 <= v279
	if cmp670 {
		goto land_lhs_true672
	} else {
		goto lor_lhs_false675
	}

land_lhs_true672:
	v280 = *lookahead
	cmp673 = v280 <= 57
	if cmp673 {
		goto if_then690
	} else {
		goto lor_lhs_false675
	}

lor_lhs_false675:
	v281 = *lookahead
	cmp676 = 65 <= v281
	if cmp676 {
		goto land_lhs_true678
	} else {
		goto lor_lhs_false681
	}

land_lhs_true678:
	v282 = *lookahead
	cmp679 = v282 <= 90
	if cmp679 {
		goto if_then690
	} else {
		goto lor_lhs_false681
	}

lor_lhs_false681:
	v283 = *lookahead
	cmp682 = v283 == 95
	if cmp682 {
		goto if_then690
	} else {
		goto lor_lhs_false684
	}

lor_lhs_false684:
	v284 = *lookahead
	cmp685 = 98 <= v284
	if cmp685 {
		goto land_lhs_true687
	} else {
		goto if_end691
	}

land_lhs_true687:
	v285 = *lookahead
	cmp688 = v285 <= 122
	if cmp688 {
		goto if_then690
	} else {
		goto if_end691
	}

if_then690:
	*state_addr = 130
	goto next_state

if_end691:
	v286 = *result
	tobool692 = (v286 & 1) != 0
	*retval = tobool692
	goto _return

sw_bb693:
	*result = 1
	v287 = *lexer_addr
	result_symbol694 = &v287.F1
	*result_symbol694 = 10
	v288 = *lexer_addr
	mark_end695 = &v288.F3
	v289 = *mark_end695
	v290 = *lexer_addr
	v289(v290)
	v291 = *lookahead
	cmp696 = v291 == 97
	if cmp696 {
		goto if_then698
	} else {
		goto if_end699
	}

if_then698:
	*state_addr = 122
	goto next_state

if_end699:
	v292 = *lookahead
	cmp700 = v292 == 45
	if cmp700 {
		goto if_then726
	} else {
		goto lor_lhs_false702
	}

lor_lhs_false702:
	v293 = *lookahead
	cmp703 = v293 == 46
	if cmp703 {
		goto if_then726
	} else {
		goto lor_lhs_false705
	}

lor_lhs_false705:
	v294 = *lookahead
	cmp706 = 48 <= v294
	if cmp706 {
		goto land_lhs_true708
	} else {
		goto lor_lhs_false711
	}

land_lhs_true708:
	v295 = *lookahead
	cmp709 = v295 <= 57
	if cmp709 {
		goto if_then726
	} else {
		goto lor_lhs_false711
	}

lor_lhs_false711:
	v296 = *lookahead
	cmp712 = 65 <= v296
	if cmp712 {
		goto land_lhs_true714
	} else {
		goto lor_lhs_false717
	}

land_lhs_true714:
	v297 = *lookahead
	cmp715 = v297 <= 90
	if cmp715 {
		goto if_then726
	} else {
		goto lor_lhs_false717
	}

lor_lhs_false717:
	v298 = *lookahead
	cmp718 = v298 == 95
	if cmp718 {
		goto if_then726
	} else {
		goto lor_lhs_false720
	}

lor_lhs_false720:
	v299 = *lookahead
	cmp721 = 98 <= v299
	if cmp721 {
		goto land_lhs_true723
	} else {
		goto if_end727
	}

land_lhs_true723:
	v300 = *lookahead
	cmp724 = v300 <= 122
	if cmp724 {
		goto if_then726
	} else {
		goto if_end727
	}

if_then726:
	*state_addr = 130
	goto next_state

if_end727:
	v301 = *result
	tobool728 = (v301 & 1) != 0
	*retval = tobool728
	goto _return

sw_bb729:
	*result = 1
	v302 = *lexer_addr
	result_symbol730 = &v302.F1
	*result_symbol730 = 10
	v303 = *lexer_addr
	mark_end731 = &v303.F3
	v304 = *mark_end731
	v305 = *lexer_addr
	v304(v305)
	v306 = *lookahead
	cmp732 = v306 == 98
	if cmp732 {
		goto if_then734
	} else {
		goto if_end735
	}

if_then734:
	*state_addr = 124
	goto next_state

if_end735:
	v307 = *lookahead
	cmp736 = v307 == 45
	if cmp736 {
		goto if_then762
	} else {
		goto lor_lhs_false738
	}

lor_lhs_false738:
	v308 = *lookahead
	cmp739 = v308 == 46
	if cmp739 {
		goto if_then762
	} else {
		goto lor_lhs_false741
	}

lor_lhs_false741:
	v309 = *lookahead
	cmp742 = 48 <= v309
	if cmp742 {
		goto land_lhs_true744
	} else {
		goto lor_lhs_false747
	}

land_lhs_true744:
	v310 = *lookahead
	cmp745 = v310 <= 57
	if cmp745 {
		goto if_then762
	} else {
		goto lor_lhs_false747
	}

lor_lhs_false747:
	v311 = *lookahead
	cmp748 = 65 <= v311
	if cmp748 {
		goto land_lhs_true750
	} else {
		goto lor_lhs_false753
	}

land_lhs_true750:
	v312 = *lookahead
	cmp751 = v312 <= 90
	if cmp751 {
		goto if_then762
	} else {
		goto lor_lhs_false753
	}

lor_lhs_false753:
	v313 = *lookahead
	cmp754 = v313 == 95
	if cmp754 {
		goto if_then762
	} else {
		goto lor_lhs_false756
	}

lor_lhs_false756:
	v314 = *lookahead
	cmp757 = 97 <= v314
	if cmp757 {
		goto land_lhs_true759
	} else {
		goto if_end763
	}

land_lhs_true759:
	v315 = *lookahead
	cmp760 = v315 <= 122
	if cmp760 {
		goto if_then762
	} else {
		goto if_end763
	}

if_then762:
	*state_addr = 130
	goto next_state

if_end763:
	v316 = *result
	tobool764 = (v316 & 1) != 0
	*retval = tobool764
	goto _return

sw_bb765:
	*result = 1
	v317 = *lexer_addr
	result_symbol766 = &v317.F1
	*result_symbol766 = 10
	v318 = *lexer_addr
	mark_end767 = &v318.F3
	v319 = *mark_end767
	v320 = *lexer_addr
	v319(v320)
	v321 = *lookahead
	cmp768 = v321 == 99
	if cmp768 {
		goto if_then770
	} else {
		goto if_end771
	}

if_then770:
	*state_addr = 113
	goto next_state

if_end771:
	v322 = *lookahead
	cmp772 = v322 == 45
	if cmp772 {
		goto if_then798
	} else {
		goto lor_lhs_false774
	}

lor_lhs_false774:
	v323 = *lookahead
	cmp775 = v323 == 46
	if cmp775 {
		goto if_then798
	} else {
		goto lor_lhs_false777
	}

lor_lhs_false777:
	v324 = *lookahead
	cmp778 = 48 <= v324
	if cmp778 {
		goto land_lhs_true780
	} else {
		goto lor_lhs_false783
	}

land_lhs_true780:
	v325 = *lookahead
	cmp781 = v325 <= 57
	if cmp781 {
		goto if_then798
	} else {
		goto lor_lhs_false783
	}

lor_lhs_false783:
	v326 = *lookahead
	cmp784 = 65 <= v326
	if cmp784 {
		goto land_lhs_true786
	} else {
		goto lor_lhs_false789
	}

land_lhs_true786:
	v327 = *lookahead
	cmp787 = v327 <= 90
	if cmp787 {
		goto if_then798
	} else {
		goto lor_lhs_false789
	}

lor_lhs_false789:
	v328 = *lookahead
	cmp790 = v328 == 95
	if cmp790 {
		goto if_then798
	} else {
		goto lor_lhs_false792
	}

lor_lhs_false792:
	v329 = *lookahead
	cmp793 = 97 <= v329
	if cmp793 {
		goto land_lhs_true795
	} else {
		goto if_end799
	}

land_lhs_true795:
	v330 = *lookahead
	cmp796 = v330 <= 122
	if cmp796 {
		goto if_then798
	} else {
		goto if_end799
	}

if_then798:
	*state_addr = 130
	goto next_state

if_end799:
	v331 = *result
	tobool800 = (v331 & 1) != 0
	*retval = tobool800
	goto _return

sw_bb801:
	*result = 1
	v332 = *lexer_addr
	result_symbol802 = &v332.F1
	*result_symbol802 = 10
	v333 = *lexer_addr
	mark_end803 = &v333.F3
	v334 = *mark_end803
	v335 = *lexer_addr
	v334(v335)
	v336 = *lookahead
	cmp804 = v336 == 100
	if cmp804 {
		goto if_then806
	} else {
		goto if_end807
	}

if_then806:
	*state_addr = 130
	goto next_state

if_end807:
	v337 = *lookahead
	cmp808 = v337 == 45
	if cmp808 {
		goto if_then834
	} else {
		goto lor_lhs_false810
	}

lor_lhs_false810:
	v338 = *lookahead
	cmp811 = v338 == 46
	if cmp811 {
		goto if_then834
	} else {
		goto lor_lhs_false813
	}

lor_lhs_false813:
	v339 = *lookahead
	cmp814 = 48 <= v339
	if cmp814 {
		goto land_lhs_true816
	} else {
		goto lor_lhs_false819
	}

land_lhs_true816:
	v340 = *lookahead
	cmp817 = v340 <= 57
	if cmp817 {
		goto if_then834
	} else {
		goto lor_lhs_false819
	}

lor_lhs_false819:
	v341 = *lookahead
	cmp820 = 65 <= v341
	if cmp820 {
		goto land_lhs_true822
	} else {
		goto lor_lhs_false825
	}

land_lhs_true822:
	v342 = *lookahead
	cmp823 = v342 <= 90
	if cmp823 {
		goto if_then834
	} else {
		goto lor_lhs_false825
	}

lor_lhs_false825:
	v343 = *lookahead
	cmp826 = v343 == 95
	if cmp826 {
		goto if_then834
	} else {
		goto lor_lhs_false828
	}

lor_lhs_false828:
	v344 = *lookahead
	cmp829 = 97 <= v344
	if cmp829 {
		goto land_lhs_true831
	} else {
		goto if_end835
	}

land_lhs_true831:
	v345 = *lookahead
	cmp832 = v345 <= 122
	if cmp832 {
		goto if_then834
	} else {
		goto if_end835
	}

if_then834:
	*state_addr = 130
	goto next_state

if_end835:
	v346 = *result
	tobool836 = (v346 & 1) != 0
	*retval = tobool836
	goto _return

sw_bb837:
	*result = 1
	v347 = *lexer_addr
	result_symbol838 = &v347.F1
	*result_symbol838 = 10
	v348 = *lexer_addr
	mark_end839 = &v348.F3
	v349 = *mark_end839
	v350 = *lexer_addr
	v349(v350)
	v351 = *lookahead
	cmp840 = v351 == 100
	if cmp840 {
		goto if_then842
	} else {
		goto if_end843
	}

if_then842:
	*state_addr = 121
	goto next_state

if_end843:
	v352 = *lookahead
	cmp844 = v352 == 45
	if cmp844 {
		goto if_then870
	} else {
		goto lor_lhs_false846
	}

lor_lhs_false846:
	v353 = *lookahead
	cmp847 = v353 == 46
	if cmp847 {
		goto if_then870
	} else {
		goto lor_lhs_false849
	}

lor_lhs_false849:
	v354 = *lookahead
	cmp850 = 48 <= v354
	if cmp850 {
		goto land_lhs_true852
	} else {
		goto lor_lhs_false855
	}

land_lhs_true852:
	v355 = *lookahead
	cmp853 = v355 <= 57
	if cmp853 {
		goto if_then870
	} else {
		goto lor_lhs_false855
	}

lor_lhs_false855:
	v356 = *lookahead
	cmp856 = 65 <= v356
	if cmp856 {
		goto land_lhs_true858
	} else {
		goto lor_lhs_false861
	}

land_lhs_true858:
	v357 = *lookahead
	cmp859 = v357 <= 90
	if cmp859 {
		goto if_then870
	} else {
		goto lor_lhs_false861
	}

lor_lhs_false861:
	v358 = *lookahead
	cmp862 = v358 == 95
	if cmp862 {
		goto if_then870
	} else {
		goto lor_lhs_false864
	}

lor_lhs_false864:
	v359 = *lookahead
	cmp865 = 97 <= v359
	if cmp865 {
		goto land_lhs_true867
	} else {
		goto if_end871
	}

land_lhs_true867:
	v360 = *lookahead
	cmp868 = v360 <= 122
	if cmp868 {
		goto if_then870
	} else {
		goto if_end871
	}

if_then870:
	*state_addr = 130
	goto next_state

if_end871:
	v361 = *result
	tobool872 = (v361 & 1) != 0
	*retval = tobool872
	goto _return

sw_bb873:
	*result = 1
	v362 = *lexer_addr
	result_symbol874 = &v362.F1
	*result_symbol874 = 10
	v363 = *lexer_addr
	mark_end875 = &v363.F3
	v364 = *mark_end875
	v365 = *lexer_addr
	v364(v365)
	v366 = *lookahead
	cmp876 = v366 == 100
	if cmp876 {
		goto if_then878
	} else {
		goto if_end879
	}

if_then878:
	*state_addr = 115
	goto next_state

if_end879:
	v367 = *lookahead
	cmp880 = v367 == 45
	if cmp880 {
		goto if_then906
	} else {
		goto lor_lhs_false882
	}

lor_lhs_false882:
	v368 = *lookahead
	cmp883 = v368 == 46
	if cmp883 {
		goto if_then906
	} else {
		goto lor_lhs_false885
	}

lor_lhs_false885:
	v369 = *lookahead
	cmp886 = 48 <= v369
	if cmp886 {
		goto land_lhs_true888
	} else {
		goto lor_lhs_false891
	}

land_lhs_true888:
	v370 = *lookahead
	cmp889 = v370 <= 57
	if cmp889 {
		goto if_then906
	} else {
		goto lor_lhs_false891
	}

lor_lhs_false891:
	v371 = *lookahead
	cmp892 = 65 <= v371
	if cmp892 {
		goto land_lhs_true894
	} else {
		goto lor_lhs_false897
	}

land_lhs_true894:
	v372 = *lookahead
	cmp895 = v372 <= 90
	if cmp895 {
		goto if_then906
	} else {
		goto lor_lhs_false897
	}

lor_lhs_false897:
	v373 = *lookahead
	cmp898 = v373 == 95
	if cmp898 {
		goto if_then906
	} else {
		goto lor_lhs_false900
	}

lor_lhs_false900:
	v374 = *lookahead
	cmp901 = 97 <= v374
	if cmp901 {
		goto land_lhs_true903
	} else {
		goto if_end907
	}

land_lhs_true903:
	v375 = *lookahead
	cmp904 = v375 <= 122
	if cmp904 {
		goto if_then906
	} else {
		goto if_end907
	}

if_then906:
	*state_addr = 130
	goto next_state

if_end907:
	v376 = *result
	tobool908 = (v376 & 1) != 0
	*retval = tobool908
	goto _return

sw_bb909:
	*result = 1
	v377 = *lexer_addr
	result_symbol910 = &v377.F1
	*result_symbol910 = 10
	v378 = *lexer_addr
	mark_end911 = &v378.F3
	v379 = *mark_end911
	v380 = *lexer_addr
	v379(v380)
	v381 = *lookahead
	cmp912 = v381 == 101
	if cmp912 {
		goto if_then914
	} else {
		goto if_end915
	}

if_then914:
	*state_addr = 125
	goto next_state

if_end915:
	v382 = *lookahead
	cmp916 = v382 == 45
	if cmp916 {
		goto if_then942
	} else {
		goto lor_lhs_false918
	}

lor_lhs_false918:
	v383 = *lookahead
	cmp919 = v383 == 46
	if cmp919 {
		goto if_then942
	} else {
		goto lor_lhs_false921
	}

lor_lhs_false921:
	v384 = *lookahead
	cmp922 = 48 <= v384
	if cmp922 {
		goto land_lhs_true924
	} else {
		goto lor_lhs_false927
	}

land_lhs_true924:
	v385 = *lookahead
	cmp925 = v385 <= 57
	if cmp925 {
		goto if_then942
	} else {
		goto lor_lhs_false927
	}

lor_lhs_false927:
	v386 = *lookahead
	cmp928 = 65 <= v386
	if cmp928 {
		goto land_lhs_true930
	} else {
		goto lor_lhs_false933
	}

land_lhs_true930:
	v387 = *lookahead
	cmp931 = v387 <= 90
	if cmp931 {
		goto if_then942
	} else {
		goto lor_lhs_false933
	}

lor_lhs_false933:
	v388 = *lookahead
	cmp934 = v388 == 95
	if cmp934 {
		goto if_then942
	} else {
		goto lor_lhs_false936
	}

lor_lhs_false936:
	v389 = *lookahead
	cmp937 = 97 <= v389
	if cmp937 {
		goto land_lhs_true939
	} else {
		goto if_end943
	}

land_lhs_true939:
	v390 = *lookahead
	cmp940 = v390 <= 122
	if cmp940 {
		goto if_then942
	} else {
		goto if_end943
	}

if_then942:
	*state_addr = 130
	goto next_state

if_end943:
	v391 = *result
	tobool944 = (v391 & 1) != 0
	*retval = tobool944
	goto _return

sw_bb945:
	*result = 1
	v392 = *lexer_addr
	result_symbol946 = &v392.F1
	*result_symbol946 = 10
	v393 = *lexer_addr
	mark_end947 = &v393.F3
	v394 = *mark_end947
	v395 = *lexer_addr
	v394(v395)
	v396 = *lookahead
	cmp948 = v396 == 101
	if cmp948 {
		goto if_then950
	} else {
		goto if_end951
	}

if_then950:
	*state_addr = 91
	goto next_state

if_end951:
	v397 = *lookahead
	cmp952 = v397 == 45
	if cmp952 {
		goto if_then978
	} else {
		goto lor_lhs_false954
	}

lor_lhs_false954:
	v398 = *lookahead
	cmp955 = v398 == 46
	if cmp955 {
		goto if_then978
	} else {
		goto lor_lhs_false957
	}

lor_lhs_false957:
	v399 = *lookahead
	cmp958 = 48 <= v399
	if cmp958 {
		goto land_lhs_true960
	} else {
		goto lor_lhs_false963
	}

land_lhs_true960:
	v400 = *lookahead
	cmp961 = v400 <= 57
	if cmp961 {
		goto if_then978
	} else {
		goto lor_lhs_false963
	}

lor_lhs_false963:
	v401 = *lookahead
	cmp964 = 65 <= v401
	if cmp964 {
		goto land_lhs_true966
	} else {
		goto lor_lhs_false969
	}

land_lhs_true966:
	v402 = *lookahead
	cmp967 = v402 <= 90
	if cmp967 {
		goto if_then978
	} else {
		goto lor_lhs_false969
	}

lor_lhs_false969:
	v403 = *lookahead
	cmp970 = v403 == 95
	if cmp970 {
		goto if_then978
	} else {
		goto lor_lhs_false972
	}

lor_lhs_false972:
	v404 = *lookahead
	cmp973 = 97 <= v404
	if cmp973 {
		goto land_lhs_true975
	} else {
		goto if_end979
	}

land_lhs_true975:
	v405 = *lookahead
	cmp976 = v405 <= 122
	if cmp976 {
		goto if_then978
	} else {
		goto if_end979
	}

if_then978:
	*state_addr = 130
	goto next_state

if_end979:
	v406 = *result
	tobool980 = (v406 & 1) != 0
	*retval = tobool980
	goto _return

sw_bb981:
	*result = 1
	v407 = *lexer_addr
	result_symbol982 = &v407.F1
	*result_symbol982 = 10
	v408 = *lexer_addr
	mark_end983 = &v408.F3
	v409 = *mark_end983
	v410 = *lexer_addr
	v409(v410)
	v411 = *lookahead
	cmp984 = v411 == 101
	if cmp984 {
		goto if_then986
	} else {
		goto if_end987
	}

if_then986:
	*state_addr = 110
	goto next_state

if_end987:
	v412 = *lookahead
	cmp988 = v412 == 45
	if cmp988 {
		goto if_then1014
	} else {
		goto lor_lhs_false990
	}

lor_lhs_false990:
	v413 = *lookahead
	cmp991 = v413 == 46
	if cmp991 {
		goto if_then1014
	} else {
		goto lor_lhs_false993
	}

lor_lhs_false993:
	v414 = *lookahead
	cmp994 = 48 <= v414
	if cmp994 {
		goto land_lhs_true996
	} else {
		goto lor_lhs_false999
	}

land_lhs_true996:
	v415 = *lookahead
	cmp997 = v415 <= 57
	if cmp997 {
		goto if_then1014
	} else {
		goto lor_lhs_false999
	}

lor_lhs_false999:
	v416 = *lookahead
	cmp1000 = 65 <= v416
	if cmp1000 {
		goto land_lhs_true1002
	} else {
		goto lor_lhs_false1005
	}

land_lhs_true1002:
	v417 = *lookahead
	cmp1003 = v417 <= 90
	if cmp1003 {
		goto if_then1014
	} else {
		goto lor_lhs_false1005
	}

lor_lhs_false1005:
	v418 = *lookahead
	cmp1006 = v418 == 95
	if cmp1006 {
		goto if_then1014
	} else {
		goto lor_lhs_false1008
	}

lor_lhs_false1008:
	v419 = *lookahead
	cmp1009 = 97 <= v419
	if cmp1009 {
		goto land_lhs_true1011
	} else {
		goto if_end1015
	}

land_lhs_true1011:
	v420 = *lookahead
	cmp1012 = v420 <= 122
	if cmp1012 {
		goto if_then1014
	} else {
		goto if_end1015
	}

if_then1014:
	*state_addr = 130
	goto next_state

if_end1015:
	v421 = *result
	tobool1016 = (v421 & 1) != 0
	*retval = tobool1016
	goto _return

sw_bb1017:
	*result = 1
	v422 = *lexer_addr
	result_symbol1018 = &v422.F1
	*result_symbol1018 = 10
	v423 = *lexer_addr
	mark_end1019 = &v423.F3
	v424 = *mark_end1019
	v425 = *lexer_addr
	v424(v425)
	v426 = *lookahead
	cmp1020 = v426 == 101
	if cmp1020 {
		goto if_then1022
	} else {
		goto if_end1023
	}

if_then1022:
	*state_addr = 97
	goto next_state

if_end1023:
	v427 = *lookahead
	cmp1024 = v427 == 45
	if cmp1024 {
		goto if_then1050
	} else {
		goto lor_lhs_false1026
	}

lor_lhs_false1026:
	v428 = *lookahead
	cmp1027 = v428 == 46
	if cmp1027 {
		goto if_then1050
	} else {
		goto lor_lhs_false1029
	}

lor_lhs_false1029:
	v429 = *lookahead
	cmp1030 = 48 <= v429
	if cmp1030 {
		goto land_lhs_true1032
	} else {
		goto lor_lhs_false1035
	}

land_lhs_true1032:
	v430 = *lookahead
	cmp1033 = v430 <= 57
	if cmp1033 {
		goto if_then1050
	} else {
		goto lor_lhs_false1035
	}

lor_lhs_false1035:
	v431 = *lookahead
	cmp1036 = 65 <= v431
	if cmp1036 {
		goto land_lhs_true1038
	} else {
		goto lor_lhs_false1041
	}

land_lhs_true1038:
	v432 = *lookahead
	cmp1039 = v432 <= 90
	if cmp1039 {
		goto if_then1050
	} else {
		goto lor_lhs_false1041
	}

lor_lhs_false1041:
	v433 = *lookahead
	cmp1042 = v433 == 95
	if cmp1042 {
		goto if_then1050
	} else {
		goto lor_lhs_false1044
	}

lor_lhs_false1044:
	v434 = *lookahead
	cmp1045 = 97 <= v434
	if cmp1045 {
		goto land_lhs_true1047
	} else {
		goto if_end1051
	}

land_lhs_true1047:
	v435 = *lookahead
	cmp1048 = v435 <= 122
	if cmp1048 {
		goto if_then1050
	} else {
		goto if_end1051
	}

if_then1050:
	*state_addr = 130
	goto next_state

if_end1051:
	v436 = *result
	tobool1052 = (v436 & 1) != 0
	*retval = tobool1052
	goto _return

sw_bb1053:
	*result = 1
	v437 = *lexer_addr
	result_symbol1054 = &v437.F1
	*result_symbol1054 = 10
	v438 = *lexer_addr
	mark_end1055 = &v438.F3
	v439 = *mark_end1055
	v440 = *lexer_addr
	v439(v440)
	v441 = *lookahead
	cmp1056 = v441 == 102
	if cmp1056 {
		goto if_then1058
	} else {
		goto if_end1059
	}

if_then1058:
	*state_addr = 106
	goto next_state

if_end1059:
	v442 = *lookahead
	cmp1060 = v442 == 45
	if cmp1060 {
		goto if_then1086
	} else {
		goto lor_lhs_false1062
	}

lor_lhs_false1062:
	v443 = *lookahead
	cmp1063 = v443 == 46
	if cmp1063 {
		goto if_then1086
	} else {
		goto lor_lhs_false1065
	}

lor_lhs_false1065:
	v444 = *lookahead
	cmp1066 = 48 <= v444
	if cmp1066 {
		goto land_lhs_true1068
	} else {
		goto lor_lhs_false1071
	}

land_lhs_true1068:
	v445 = *lookahead
	cmp1069 = v445 <= 57
	if cmp1069 {
		goto if_then1086
	} else {
		goto lor_lhs_false1071
	}

lor_lhs_false1071:
	v446 = *lookahead
	cmp1072 = 65 <= v446
	if cmp1072 {
		goto land_lhs_true1074
	} else {
		goto lor_lhs_false1077
	}

land_lhs_true1074:
	v447 = *lookahead
	cmp1075 = v447 <= 90
	if cmp1075 {
		goto if_then1086
	} else {
		goto lor_lhs_false1077
	}

lor_lhs_false1077:
	v448 = *lookahead
	cmp1078 = v448 == 95
	if cmp1078 {
		goto if_then1086
	} else {
		goto lor_lhs_false1080
	}

lor_lhs_false1080:
	v449 = *lookahead
	cmp1081 = 97 <= v449
	if cmp1081 {
		goto land_lhs_true1083
	} else {
		goto if_end1087
	}

land_lhs_true1083:
	v450 = *lookahead
	cmp1084 = v450 <= 122
	if cmp1084 {
		goto if_then1086
	} else {
		goto if_end1087
	}

if_then1086:
	*state_addr = 130
	goto next_state

if_end1087:
	v451 = *result
	tobool1088 = (v451 & 1) != 0
	*retval = tobool1088
	goto _return

sw_bb1089:
	*result = 1
	v452 = *lexer_addr
	result_symbol1090 = &v452.F1
	*result_symbol1090 = 10
	v453 = *lexer_addr
	mark_end1091 = &v453.F3
	v454 = *mark_end1091
	v455 = *lexer_addr
	v454(v455)
	v456 = *lookahead
	cmp1092 = v456 == 103
	if cmp1092 {
		goto if_then1094
	} else {
		goto if_end1095
	}

if_then1094:
	*state_addr = 117
	goto next_state

if_end1095:
	v457 = *lookahead
	cmp1096 = v457 == 45
	if cmp1096 {
		goto if_then1122
	} else {
		goto lor_lhs_false1098
	}

lor_lhs_false1098:
	v458 = *lookahead
	cmp1099 = v458 == 46
	if cmp1099 {
		goto if_then1122
	} else {
		goto lor_lhs_false1101
	}

lor_lhs_false1101:
	v459 = *lookahead
	cmp1102 = 48 <= v459
	if cmp1102 {
		goto land_lhs_true1104
	} else {
		goto lor_lhs_false1107
	}

land_lhs_true1104:
	v460 = *lookahead
	cmp1105 = v460 <= 57
	if cmp1105 {
		goto if_then1122
	} else {
		goto lor_lhs_false1107
	}

lor_lhs_false1107:
	v461 = *lookahead
	cmp1108 = 65 <= v461
	if cmp1108 {
		goto land_lhs_true1110
	} else {
		goto lor_lhs_false1113
	}

land_lhs_true1110:
	v462 = *lookahead
	cmp1111 = v462 <= 90
	if cmp1111 {
		goto if_then1122
	} else {
		goto lor_lhs_false1113
	}

lor_lhs_false1113:
	v463 = *lookahead
	cmp1114 = v463 == 95
	if cmp1114 {
		goto if_then1122
	} else {
		goto lor_lhs_false1116
	}

lor_lhs_false1116:
	v464 = *lookahead
	cmp1117 = 97 <= v464
	if cmp1117 {
		goto land_lhs_true1119
	} else {
		goto if_end1123
	}

land_lhs_true1119:
	v465 = *lookahead
	cmp1120 = v465 <= 122
	if cmp1120 {
		goto if_then1122
	} else {
		goto if_end1123
	}

if_then1122:
	*state_addr = 130
	goto next_state

if_end1123:
	v466 = *result
	tobool1124 = (v466 & 1) != 0
	*retval = tobool1124
	goto _return

sw_bb1125:
	*result = 1
	v467 = *lexer_addr
	result_symbol1126 = &v467.F1
	*result_symbol1126 = 10
	v468 = *lexer_addr
	mark_end1127 = &v468.F3
	v469 = *mark_end1127
	v470 = *lexer_addr
	v469(v470)
	v471 = *lookahead
	cmp1128 = v471 == 105
	if cmp1128 {
		goto if_then1130
	} else {
		goto if_end1131
	}

if_then1130:
	*state_addr = 123
	goto next_state

if_end1131:
	v472 = *lookahead
	cmp1132 = v472 == 45
	if cmp1132 {
		goto if_then1158
	} else {
		goto lor_lhs_false1134
	}

lor_lhs_false1134:
	v473 = *lookahead
	cmp1135 = v473 == 46
	if cmp1135 {
		goto if_then1158
	} else {
		goto lor_lhs_false1137
	}

lor_lhs_false1137:
	v474 = *lookahead
	cmp1138 = 48 <= v474
	if cmp1138 {
		goto land_lhs_true1140
	} else {
		goto lor_lhs_false1143
	}

land_lhs_true1140:
	v475 = *lookahead
	cmp1141 = v475 <= 57
	if cmp1141 {
		goto if_then1158
	} else {
		goto lor_lhs_false1143
	}

lor_lhs_false1143:
	v476 = *lookahead
	cmp1144 = 65 <= v476
	if cmp1144 {
		goto land_lhs_true1146
	} else {
		goto lor_lhs_false1149
	}

land_lhs_true1146:
	v477 = *lookahead
	cmp1147 = v477 <= 90
	if cmp1147 {
		goto if_then1158
	} else {
		goto lor_lhs_false1149
	}

lor_lhs_false1149:
	v478 = *lookahead
	cmp1150 = v478 == 95
	if cmp1150 {
		goto if_then1158
	} else {
		goto lor_lhs_false1152
	}

lor_lhs_false1152:
	v479 = *lookahead
	cmp1153 = 97 <= v479
	if cmp1153 {
		goto land_lhs_true1155
	} else {
		goto if_end1159
	}

land_lhs_true1155:
	v480 = *lookahead
	cmp1156 = v480 <= 122
	if cmp1156 {
		goto if_then1158
	} else {
		goto if_end1159
	}

if_then1158:
	*state_addr = 130
	goto next_state

if_end1159:
	v481 = *result
	tobool1160 = (v481 & 1) != 0
	*retval = tobool1160
	goto _return

sw_bb1161:
	*result = 1
	v482 = *lexer_addr
	result_symbol1162 = &v482.F1
	*result_symbol1162 = 10
	v483 = *lexer_addr
	mark_end1163 = &v483.F3
	v484 = *mark_end1163
	v485 = *lexer_addr
	v484(v485)
	v486 = *lookahead
	cmp1164 = v486 == 105
	if cmp1164 {
		goto if_then1166
	} else {
		goto if_end1167
	}

if_then1166:
	*state_addr = 109
	goto next_state

if_end1167:
	v487 = *lookahead
	cmp1168 = v487 == 45
	if cmp1168 {
		goto if_then1194
	} else {
		goto lor_lhs_false1170
	}

lor_lhs_false1170:
	v488 = *lookahead
	cmp1171 = v488 == 46
	if cmp1171 {
		goto if_then1194
	} else {
		goto lor_lhs_false1173
	}

lor_lhs_false1173:
	v489 = *lookahead
	cmp1174 = 48 <= v489
	if cmp1174 {
		goto land_lhs_true1176
	} else {
		goto lor_lhs_false1179
	}

land_lhs_true1176:
	v490 = *lookahead
	cmp1177 = v490 <= 57
	if cmp1177 {
		goto if_then1194
	} else {
		goto lor_lhs_false1179
	}

lor_lhs_false1179:
	v491 = *lookahead
	cmp1180 = 65 <= v491
	if cmp1180 {
		goto land_lhs_true1182
	} else {
		goto lor_lhs_false1185
	}

land_lhs_true1182:
	v492 = *lookahead
	cmp1183 = v492 <= 90
	if cmp1183 {
		goto if_then1194
	} else {
		goto lor_lhs_false1185
	}

lor_lhs_false1185:
	v493 = *lookahead
	cmp1186 = v493 == 95
	if cmp1186 {
		goto if_then1194
	} else {
		goto lor_lhs_false1188
	}

lor_lhs_false1188:
	v494 = *lookahead
	cmp1189 = 97 <= v494
	if cmp1189 {
		goto land_lhs_true1191
	} else {
		goto if_end1195
	}

land_lhs_true1191:
	v495 = *lookahead
	cmp1192 = v495 <= 122
	if cmp1192 {
		goto if_then1194
	} else {
		goto if_end1195
	}

if_then1194:
	*state_addr = 130
	goto next_state

if_end1195:
	v496 = *result
	tobool1196 = (v496 & 1) != 0
	*retval = tobool1196
	goto _return

sw_bb1197:
	*result = 1
	v497 = *lexer_addr
	result_symbol1198 = &v497.F1
	*result_symbol1198 = 10
	v498 = *lexer_addr
	mark_end1199 = &v498.F3
	v499 = *mark_end1199
	v500 = *lexer_addr
	v499(v500)
	v501 = *lookahead
	cmp1200 = v501 == 105
	if cmp1200 {
		goto if_then1202
	} else {
		goto if_end1203
	}

if_then1202:
	*state_addr = 93
	goto next_state

if_end1203:
	v502 = *lookahead
	cmp1204 = v502 == 45
	if cmp1204 {
		goto if_then1230
	} else {
		goto lor_lhs_false1206
	}

lor_lhs_false1206:
	v503 = *lookahead
	cmp1207 = v503 == 46
	if cmp1207 {
		goto if_then1230
	} else {
		goto lor_lhs_false1209
	}

lor_lhs_false1209:
	v504 = *lookahead
	cmp1210 = 48 <= v504
	if cmp1210 {
		goto land_lhs_true1212
	} else {
		goto lor_lhs_false1215
	}

land_lhs_true1212:
	v505 = *lookahead
	cmp1213 = v505 <= 57
	if cmp1213 {
		goto if_then1230
	} else {
		goto lor_lhs_false1215
	}

lor_lhs_false1215:
	v506 = *lookahead
	cmp1216 = 65 <= v506
	if cmp1216 {
		goto land_lhs_true1218
	} else {
		goto lor_lhs_false1221
	}

land_lhs_true1218:
	v507 = *lookahead
	cmp1219 = v507 <= 90
	if cmp1219 {
		goto if_then1230
	} else {
		goto lor_lhs_false1221
	}

lor_lhs_false1221:
	v508 = *lookahead
	cmp1222 = v508 == 95
	if cmp1222 {
		goto if_then1230
	} else {
		goto lor_lhs_false1224
	}

lor_lhs_false1224:
	v509 = *lookahead
	cmp1225 = 97 <= v509
	if cmp1225 {
		goto land_lhs_true1227
	} else {
		goto if_end1231
	}

land_lhs_true1227:
	v510 = *lookahead
	cmp1228 = v510 <= 122
	if cmp1228 {
		goto if_then1230
	} else {
		goto if_end1231
	}

if_then1230:
	*state_addr = 130
	goto next_state

if_end1231:
	v511 = *result
	tobool1232 = (v511 & 1) != 0
	*retval = tobool1232
	goto _return

sw_bb1233:
	*result = 1
	v512 = *lexer_addr
	result_symbol1234 = &v512.F1
	*result_symbol1234 = 10
	v513 = *lexer_addr
	mark_end1235 = &v513.F3
	v514 = *mark_end1235
	v515 = *lexer_addr
	v514(v515)
	v516 = *lookahead
	cmp1236 = v516 == 105
	if cmp1236 {
		goto if_then1238
	} else {
		goto if_end1239
	}

if_then1238:
	*state_addr = 112
	goto next_state

if_end1239:
	v517 = *lookahead
	cmp1240 = v517 == 45
	if cmp1240 {
		goto if_then1266
	} else {
		goto lor_lhs_false1242
	}

lor_lhs_false1242:
	v518 = *lookahead
	cmp1243 = v518 == 46
	if cmp1243 {
		goto if_then1266
	} else {
		goto lor_lhs_false1245
	}

lor_lhs_false1245:
	v519 = *lookahead
	cmp1246 = 48 <= v519
	if cmp1246 {
		goto land_lhs_true1248
	} else {
		goto lor_lhs_false1251
	}

land_lhs_true1248:
	v520 = *lookahead
	cmp1249 = v520 <= 57
	if cmp1249 {
		goto if_then1266
	} else {
		goto lor_lhs_false1251
	}

lor_lhs_false1251:
	v521 = *lookahead
	cmp1252 = 65 <= v521
	if cmp1252 {
		goto land_lhs_true1254
	} else {
		goto lor_lhs_false1257
	}

land_lhs_true1254:
	v522 = *lookahead
	cmp1255 = v522 <= 90
	if cmp1255 {
		goto if_then1266
	} else {
		goto lor_lhs_false1257
	}

lor_lhs_false1257:
	v523 = *lookahead
	cmp1258 = v523 == 95
	if cmp1258 {
		goto if_then1266
	} else {
		goto lor_lhs_false1260
	}

lor_lhs_false1260:
	v524 = *lookahead
	cmp1261 = 97 <= v524
	if cmp1261 {
		goto land_lhs_true1263
	} else {
		goto if_end1267
	}

land_lhs_true1263:
	v525 = *lookahead
	cmp1264 = v525 <= 122
	if cmp1264 {
		goto if_then1266
	} else {
		goto if_end1267
	}

if_then1266:
	*state_addr = 130
	goto next_state

if_end1267:
	v526 = *result
	tobool1268 = (v526 & 1) != 0
	*retval = tobool1268
	goto _return

sw_bb1269:
	*result = 1
	v527 = *lexer_addr
	result_symbol1270 = &v527.F1
	*result_symbol1270 = 10
	v528 = *lexer_addr
	mark_end1271 = &v528.F3
	v529 = *mark_end1271
	v530 = *lexer_addr
	v529(v530)
	v531 = *lookahead
	cmp1272 = v531 == 108
	if cmp1272 {
		goto if_then1274
	} else {
		goto if_end1275
	}

if_then1274:
	*state_addr = 129
	goto next_state

if_end1275:
	v532 = *lookahead
	cmp1276 = v532 == 45
	if cmp1276 {
		goto if_then1302
	} else {
		goto lor_lhs_false1278
	}

lor_lhs_false1278:
	v533 = *lookahead
	cmp1279 = v533 == 46
	if cmp1279 {
		goto if_then1302
	} else {
		goto lor_lhs_false1281
	}

lor_lhs_false1281:
	v534 = *lookahead
	cmp1282 = 48 <= v534
	if cmp1282 {
		goto land_lhs_true1284
	} else {
		goto lor_lhs_false1287
	}

land_lhs_true1284:
	v535 = *lookahead
	cmp1285 = v535 <= 57
	if cmp1285 {
		goto if_then1302
	} else {
		goto lor_lhs_false1287
	}

lor_lhs_false1287:
	v536 = *lookahead
	cmp1288 = 65 <= v536
	if cmp1288 {
		goto land_lhs_true1290
	} else {
		goto lor_lhs_false1293
	}

land_lhs_true1290:
	v537 = *lookahead
	cmp1291 = v537 <= 90
	if cmp1291 {
		goto if_then1302
	} else {
		goto lor_lhs_false1293
	}

lor_lhs_false1293:
	v538 = *lookahead
	cmp1294 = v538 == 95
	if cmp1294 {
		goto if_then1302
	} else {
		goto lor_lhs_false1296
	}

lor_lhs_false1296:
	v539 = *lookahead
	cmp1297 = 97 <= v539
	if cmp1297 {
		goto land_lhs_true1299
	} else {
		goto if_end1303
	}

land_lhs_true1299:
	v540 = *lookahead
	cmp1300 = v540 <= 122
	if cmp1300 {
		goto if_then1302
	} else {
		goto if_end1303
	}

if_then1302:
	*state_addr = 130
	goto next_state

if_end1303:
	v541 = *result
	tobool1304 = (v541 & 1) != 0
	*retval = tobool1304
	goto _return

sw_bb1305:
	*result = 1
	v542 = *lexer_addr
	result_symbol1306 = &v542.F1
	*result_symbol1306 = 10
	v543 = *lexer_addr
	mark_end1307 = &v543.F3
	v544 = *mark_end1307
	v545 = *lexer_addr
	v544(v545)
	v546 = *lookahead
	cmp1308 = v546 == 108
	if cmp1308 {
		goto if_then1310
	} else {
		goto if_end1311
	}

if_then1310:
	*state_addr = 90
	goto next_state

if_end1311:
	v547 = *lookahead
	cmp1312 = v547 == 45
	if cmp1312 {
		goto if_then1338
	} else {
		goto lor_lhs_false1314
	}

lor_lhs_false1314:
	v548 = *lookahead
	cmp1315 = v548 == 46
	if cmp1315 {
		goto if_then1338
	} else {
		goto lor_lhs_false1317
	}

lor_lhs_false1317:
	v549 = *lookahead
	cmp1318 = 48 <= v549
	if cmp1318 {
		goto land_lhs_true1320
	} else {
		goto lor_lhs_false1323
	}

land_lhs_true1320:
	v550 = *lookahead
	cmp1321 = v550 <= 57
	if cmp1321 {
		goto if_then1338
	} else {
		goto lor_lhs_false1323
	}

lor_lhs_false1323:
	v551 = *lookahead
	cmp1324 = 65 <= v551
	if cmp1324 {
		goto land_lhs_true1326
	} else {
		goto lor_lhs_false1329
	}

land_lhs_true1326:
	v552 = *lookahead
	cmp1327 = v552 <= 90
	if cmp1327 {
		goto if_then1338
	} else {
		goto lor_lhs_false1329
	}

lor_lhs_false1329:
	v553 = *lookahead
	cmp1330 = v553 == 95
	if cmp1330 {
		goto if_then1338
	} else {
		goto lor_lhs_false1332
	}

lor_lhs_false1332:
	v554 = *lookahead
	cmp1333 = 97 <= v554
	if cmp1333 {
		goto land_lhs_true1335
	} else {
		goto if_end1339
	}

land_lhs_true1335:
	v555 = *lookahead
	cmp1336 = v555 <= 122
	if cmp1336 {
		goto if_then1338
	} else {
		goto if_end1339
	}

if_then1338:
	*state_addr = 130
	goto next_state

if_end1339:
	v556 = *result
	tobool1340 = (v556 & 1) != 0
	*retval = tobool1340
	goto _return

sw_bb1341:
	*result = 1
	v557 = *lexer_addr
	result_symbol1342 = &v557.F1
	*result_symbol1342 = 10
	v558 = *lexer_addr
	mark_end1343 = &v558.F3
	v559 = *mark_end1343
	v560 = *lexer_addr
	v559(v560)
	v561 = *lookahead
	cmp1344 = v561 == 110
	if cmp1344 {
		goto if_then1346
	} else {
		goto if_end1347
	}

if_then1346:
	*state_addr = 94
	goto next_state

if_end1347:
	v562 = *lookahead
	cmp1348 = v562 == 45
	if cmp1348 {
		goto if_then1374
	} else {
		goto lor_lhs_false1350
	}

lor_lhs_false1350:
	v563 = *lookahead
	cmp1351 = v563 == 46
	if cmp1351 {
		goto if_then1374
	} else {
		goto lor_lhs_false1353
	}

lor_lhs_false1353:
	v564 = *lookahead
	cmp1354 = 48 <= v564
	if cmp1354 {
		goto land_lhs_true1356
	} else {
		goto lor_lhs_false1359
	}

land_lhs_true1356:
	v565 = *lookahead
	cmp1357 = v565 <= 57
	if cmp1357 {
		goto if_then1374
	} else {
		goto lor_lhs_false1359
	}

lor_lhs_false1359:
	v566 = *lookahead
	cmp1360 = 65 <= v566
	if cmp1360 {
		goto land_lhs_true1362
	} else {
		goto lor_lhs_false1365
	}

land_lhs_true1362:
	v567 = *lookahead
	cmp1363 = v567 <= 90
	if cmp1363 {
		goto if_then1374
	} else {
		goto lor_lhs_false1365
	}

lor_lhs_false1365:
	v568 = *lookahead
	cmp1366 = v568 == 95
	if cmp1366 {
		goto if_then1374
	} else {
		goto lor_lhs_false1368
	}

lor_lhs_false1368:
	v569 = *lookahead
	cmp1369 = 97 <= v569
	if cmp1369 {
		goto land_lhs_true1371
	} else {
		goto if_end1375
	}

land_lhs_true1371:
	v570 = *lookahead
	cmp1372 = v570 <= 122
	if cmp1372 {
		goto if_then1374
	} else {
		goto if_end1375
	}

if_then1374:
	*state_addr = 130
	goto next_state

if_end1375:
	v571 = *result
	tobool1376 = (v571 & 1) != 0
	*retval = tobool1376
	goto _return

sw_bb1377:
	*result = 1
	v572 = *lexer_addr
	result_symbol1378 = &v572.F1
	*result_symbol1378 = 10
	v573 = *lexer_addr
	mark_end1379 = &v573.F3
	v574 = *mark_end1379
	v575 = *lexer_addr
	v574(v575)
	v576 = *lookahead
	cmp1380 = v576 == 110
	if cmp1380 {
		goto if_then1382
	} else {
		goto if_end1383
	}

if_then1382:
	*state_addr = 126
	goto next_state

if_end1383:
	v577 = *lookahead
	cmp1384 = v577 == 45
	if cmp1384 {
		goto if_then1410
	} else {
		goto lor_lhs_false1386
	}

lor_lhs_false1386:
	v578 = *lookahead
	cmp1387 = v578 == 46
	if cmp1387 {
		goto if_then1410
	} else {
		goto lor_lhs_false1389
	}

lor_lhs_false1389:
	v579 = *lookahead
	cmp1390 = 48 <= v579
	if cmp1390 {
		goto land_lhs_true1392
	} else {
		goto lor_lhs_false1395
	}

land_lhs_true1392:
	v580 = *lookahead
	cmp1393 = v580 <= 57
	if cmp1393 {
		goto if_then1410
	} else {
		goto lor_lhs_false1395
	}

lor_lhs_false1395:
	v581 = *lookahead
	cmp1396 = 65 <= v581
	if cmp1396 {
		goto land_lhs_true1398
	} else {
		goto lor_lhs_false1401
	}

land_lhs_true1398:
	v582 = *lookahead
	cmp1399 = v582 <= 90
	if cmp1399 {
		goto if_then1410
	} else {
		goto lor_lhs_false1401
	}

lor_lhs_false1401:
	v583 = *lookahead
	cmp1402 = v583 == 95
	if cmp1402 {
		goto if_then1410
	} else {
		goto lor_lhs_false1404
	}

lor_lhs_false1404:
	v584 = *lookahead
	cmp1405 = 97 <= v584
	if cmp1405 {
		goto land_lhs_true1407
	} else {
		goto if_end1411
	}

land_lhs_true1407:
	v585 = *lookahead
	cmp1408 = v585 <= 122
	if cmp1408 {
		goto if_then1410
	} else {
		goto if_end1411
	}

if_then1410:
	*state_addr = 130
	goto next_state

if_end1411:
	v586 = *result
	tobool1412 = (v586 & 1) != 0
	*retval = tobool1412
	goto _return

sw_bb1413:
	*result = 1
	v587 = *lexer_addr
	result_symbol1414 = &v587.F1
	*result_symbol1414 = 10
	v588 = *lexer_addr
	mark_end1415 = &v588.F3
	v589 = *mark_end1415
	v590 = *lexer_addr
	v589(v590)
	v591 = *lookahead
	cmp1416 = v591 == 110
	if cmp1416 {
		goto if_then1418
	} else {
		goto if_end1419
	}

if_then1418:
	*state_addr = 91
	goto next_state

if_end1419:
	v592 = *lookahead
	cmp1420 = v592 == 45
	if cmp1420 {
		goto if_then1446
	} else {
		goto lor_lhs_false1422
	}

lor_lhs_false1422:
	v593 = *lookahead
	cmp1423 = v593 == 46
	if cmp1423 {
		goto if_then1446
	} else {
		goto lor_lhs_false1425
	}

lor_lhs_false1425:
	v594 = *lookahead
	cmp1426 = 48 <= v594
	if cmp1426 {
		goto land_lhs_true1428
	} else {
		goto lor_lhs_false1431
	}

land_lhs_true1428:
	v595 = *lookahead
	cmp1429 = v595 <= 57
	if cmp1429 {
		goto if_then1446
	} else {
		goto lor_lhs_false1431
	}

lor_lhs_false1431:
	v596 = *lookahead
	cmp1432 = 65 <= v596
	if cmp1432 {
		goto land_lhs_true1434
	} else {
		goto lor_lhs_false1437
	}

land_lhs_true1434:
	v597 = *lookahead
	cmp1435 = v597 <= 90
	if cmp1435 {
		goto if_then1446
	} else {
		goto lor_lhs_false1437
	}

lor_lhs_false1437:
	v598 = *lookahead
	cmp1438 = v598 == 95
	if cmp1438 {
		goto if_then1446
	} else {
		goto lor_lhs_false1440
	}

lor_lhs_false1440:
	v599 = *lookahead
	cmp1441 = 97 <= v599
	if cmp1441 {
		goto land_lhs_true1443
	} else {
		goto if_end1447
	}

land_lhs_true1443:
	v600 = *lookahead
	cmp1444 = v600 <= 122
	if cmp1444 {
		goto if_then1446
	} else {
		goto if_end1447
	}

if_then1446:
	*state_addr = 130
	goto next_state

if_end1447:
	v601 = *result
	tobool1448 = (v601 & 1) != 0
	*retval = tobool1448
	goto _return

sw_bb1449:
	*result = 1
	v602 = *lexer_addr
	result_symbol1450 = &v602.F1
	*result_symbol1450 = 10
	v603 = *lexer_addr
	mark_end1451 = &v603.F3
	v604 = *mark_end1451
	v605 = *lexer_addr
	v604(v605)
	v606 = *lookahead
	cmp1452 = v606 == 110
	if cmp1452 {
		goto if_then1454
	} else {
		goto if_end1455
	}

if_then1454:
	*state_addr = 100
	goto next_state

if_end1455:
	v607 = *lookahead
	cmp1456 = v607 == 45
	if cmp1456 {
		goto if_then1482
	} else {
		goto lor_lhs_false1458
	}

lor_lhs_false1458:
	v608 = *lookahead
	cmp1459 = v608 == 46
	if cmp1459 {
		goto if_then1482
	} else {
		goto lor_lhs_false1461
	}

lor_lhs_false1461:
	v609 = *lookahead
	cmp1462 = 48 <= v609
	if cmp1462 {
		goto land_lhs_true1464
	} else {
		goto lor_lhs_false1467
	}

land_lhs_true1464:
	v610 = *lookahead
	cmp1465 = v610 <= 57
	if cmp1465 {
		goto if_then1482
	} else {
		goto lor_lhs_false1467
	}

lor_lhs_false1467:
	v611 = *lookahead
	cmp1468 = 65 <= v611
	if cmp1468 {
		goto land_lhs_true1470
	} else {
		goto lor_lhs_false1473
	}

land_lhs_true1470:
	v612 = *lookahead
	cmp1471 = v612 <= 90
	if cmp1471 {
		goto if_then1482
	} else {
		goto lor_lhs_false1473
	}

lor_lhs_false1473:
	v613 = *lookahead
	cmp1474 = v613 == 95
	if cmp1474 {
		goto if_then1482
	} else {
		goto lor_lhs_false1476
	}

lor_lhs_false1476:
	v614 = *lookahead
	cmp1477 = 97 <= v614
	if cmp1477 {
		goto land_lhs_true1479
	} else {
		goto if_end1483
	}

land_lhs_true1479:
	v615 = *lookahead
	cmp1480 = v615 <= 122
	if cmp1480 {
		goto if_then1482
	} else {
		goto if_end1483
	}

if_then1482:
	*state_addr = 130
	goto next_state

if_end1483:
	v616 = *result
	tobool1484 = (v616 & 1) != 0
	*retval = tobool1484
	goto _return

sw_bb1485:
	*result = 1
	v617 = *lexer_addr
	result_symbol1486 = &v617.F1
	*result_symbol1486 = 10
	v618 = *lexer_addr
	mark_end1487 = &v618.F3
	v619 = *mark_end1487
	v620 = *lexer_addr
	v619(v620)
	v621 = *lookahead
	cmp1488 = v621 == 110
	if cmp1488 {
		goto if_then1490
	} else {
		goto if_end1491
	}

if_then1490:
	*state_addr = 95
	goto next_state

if_end1491:
	v622 = *lookahead
	cmp1492 = v622 == 45
	if cmp1492 {
		goto if_then1518
	} else {
		goto lor_lhs_false1494
	}

lor_lhs_false1494:
	v623 = *lookahead
	cmp1495 = v623 == 46
	if cmp1495 {
		goto if_then1518
	} else {
		goto lor_lhs_false1497
	}

lor_lhs_false1497:
	v624 = *lookahead
	cmp1498 = 48 <= v624
	if cmp1498 {
		goto land_lhs_true1500
	} else {
		goto lor_lhs_false1503
	}

land_lhs_true1500:
	v625 = *lookahead
	cmp1501 = v625 <= 57
	if cmp1501 {
		goto if_then1518
	} else {
		goto lor_lhs_false1503
	}

lor_lhs_false1503:
	v626 = *lookahead
	cmp1504 = 65 <= v626
	if cmp1504 {
		goto land_lhs_true1506
	} else {
		goto lor_lhs_false1509
	}

land_lhs_true1506:
	v627 = *lookahead
	cmp1507 = v627 <= 90
	if cmp1507 {
		goto if_then1518
	} else {
		goto lor_lhs_false1509
	}

lor_lhs_false1509:
	v628 = *lookahead
	cmp1510 = v628 == 95
	if cmp1510 {
		goto if_then1518
	} else {
		goto lor_lhs_false1512
	}

lor_lhs_false1512:
	v629 = *lookahead
	cmp1513 = 97 <= v629
	if cmp1513 {
		goto land_lhs_true1515
	} else {
		goto if_end1519
	}

land_lhs_true1515:
	v630 = *lookahead
	cmp1516 = v630 <= 122
	if cmp1516 {
		goto if_then1518
	} else {
		goto if_end1519
	}

if_then1518:
	*state_addr = 130
	goto next_state

if_end1519:
	v631 = *result
	tobool1520 = (v631 & 1) != 0
	*retval = tobool1520
	goto _return

sw_bb1521:
	*result = 1
	v632 = *lexer_addr
	result_symbol1522 = &v632.F1
	*result_symbol1522 = 10
	v633 = *lexer_addr
	mark_end1523 = &v633.F3
	v634 = *mark_end1523
	v635 = *lexer_addr
	v634(v635)
	v636 = *lookahead
	cmp1524 = v636 == 111
	if cmp1524 {
		goto if_then1526
	} else {
		goto if_end1527
	}

if_then1526:
	*state_addr = 123
	goto next_state

if_end1527:
	v637 = *lookahead
	cmp1528 = v637 == 45
	if cmp1528 {
		goto if_then1554
	} else {
		goto lor_lhs_false1530
	}

lor_lhs_false1530:
	v638 = *lookahead
	cmp1531 = v638 == 46
	if cmp1531 {
		goto if_then1554
	} else {
		goto lor_lhs_false1533
	}

lor_lhs_false1533:
	v639 = *lookahead
	cmp1534 = 48 <= v639
	if cmp1534 {
		goto land_lhs_true1536
	} else {
		goto lor_lhs_false1539
	}

land_lhs_true1536:
	v640 = *lookahead
	cmp1537 = v640 <= 57
	if cmp1537 {
		goto if_then1554
	} else {
		goto lor_lhs_false1539
	}

lor_lhs_false1539:
	v641 = *lookahead
	cmp1540 = 65 <= v641
	if cmp1540 {
		goto land_lhs_true1542
	} else {
		goto lor_lhs_false1545
	}

land_lhs_true1542:
	v642 = *lookahead
	cmp1543 = v642 <= 90
	if cmp1543 {
		goto if_then1554
	} else {
		goto lor_lhs_false1545
	}

lor_lhs_false1545:
	v643 = *lookahead
	cmp1546 = v643 == 95
	if cmp1546 {
		goto if_then1554
	} else {
		goto lor_lhs_false1548
	}

lor_lhs_false1548:
	v644 = *lookahead
	cmp1549 = 97 <= v644
	if cmp1549 {
		goto land_lhs_true1551
	} else {
		goto if_end1555
	}

land_lhs_true1551:
	v645 = *lookahead
	cmp1552 = v645 <= 122
	if cmp1552 {
		goto if_then1554
	} else {
		goto if_end1555
	}

if_then1554:
	*state_addr = 130
	goto next_state

if_end1555:
	v646 = *result
	tobool1556 = (v646 & 1) != 0
	*retval = tobool1556
	goto _return

sw_bb1557:
	*result = 1
	v647 = *lexer_addr
	result_symbol1558 = &v647.F1
	*result_symbol1558 = 10
	v648 = *lexer_addr
	mark_end1559 = &v648.F3
	v649 = *mark_end1559
	v650 = *lexer_addr
	v649(v650)
	v651 = *lookahead
	cmp1560 = v651 == 111
	if cmp1560 {
		goto if_then1562
	} else {
		goto if_end1563
	}

if_then1562:
	*state_addr = 107
	goto next_state

if_end1563:
	v652 = *lookahead
	cmp1564 = v652 == 45
	if cmp1564 {
		goto if_then1590
	} else {
		goto lor_lhs_false1566
	}

lor_lhs_false1566:
	v653 = *lookahead
	cmp1567 = v653 == 46
	if cmp1567 {
		goto if_then1590
	} else {
		goto lor_lhs_false1569
	}

lor_lhs_false1569:
	v654 = *lookahead
	cmp1570 = 48 <= v654
	if cmp1570 {
		goto land_lhs_true1572
	} else {
		goto lor_lhs_false1575
	}

land_lhs_true1572:
	v655 = *lookahead
	cmp1573 = v655 <= 57
	if cmp1573 {
		goto if_then1590
	} else {
		goto lor_lhs_false1575
	}

lor_lhs_false1575:
	v656 = *lookahead
	cmp1576 = 65 <= v656
	if cmp1576 {
		goto land_lhs_true1578
	} else {
		goto lor_lhs_false1581
	}

land_lhs_true1578:
	v657 = *lookahead
	cmp1579 = v657 <= 90
	if cmp1579 {
		goto if_then1590
	} else {
		goto lor_lhs_false1581
	}

lor_lhs_false1581:
	v658 = *lookahead
	cmp1582 = v658 == 95
	if cmp1582 {
		goto if_then1590
	} else {
		goto lor_lhs_false1584
	}

lor_lhs_false1584:
	v659 = *lookahead
	cmp1585 = 97 <= v659
	if cmp1585 {
		goto land_lhs_true1587
	} else {
		goto if_end1591
	}

land_lhs_true1587:
	v660 = *lookahead
	cmp1588 = v660 <= 122
	if cmp1588 {
		goto if_then1590
	} else {
		goto if_end1591
	}

if_then1590:
	*state_addr = 130
	goto next_state

if_end1591:
	v661 = *result
	tobool1592 = (v661 & 1) != 0
	*retval = tobool1592
	goto _return

sw_bb1593:
	*result = 1
	v662 = *lexer_addr
	result_symbol1594 = &v662.F1
	*result_symbol1594 = 10
	v663 = *lexer_addr
	mark_end1595 = &v663.F3
	v664 = *mark_end1595
	v665 = *lexer_addr
	v664(v665)
	v666 = *lookahead
	cmp1596 = v666 == 111
	if cmp1596 {
		goto if_then1598
	} else {
		goto if_end1599
	}

if_then1598:
	*state_addr = 127
	goto next_state

if_end1599:
	v667 = *lookahead
	cmp1600 = v667 == 45
	if cmp1600 {
		goto if_then1626
	} else {
		goto lor_lhs_false1602
	}

lor_lhs_false1602:
	v668 = *lookahead
	cmp1603 = v668 == 46
	if cmp1603 {
		goto if_then1626
	} else {
		goto lor_lhs_false1605
	}

lor_lhs_false1605:
	v669 = *lookahead
	cmp1606 = 48 <= v669
	if cmp1606 {
		goto land_lhs_true1608
	} else {
		goto lor_lhs_false1611
	}

land_lhs_true1608:
	v670 = *lookahead
	cmp1609 = v670 <= 57
	if cmp1609 {
		goto if_then1626
	} else {
		goto lor_lhs_false1611
	}

lor_lhs_false1611:
	v671 = *lookahead
	cmp1612 = 65 <= v671
	if cmp1612 {
		goto land_lhs_true1614
	} else {
		goto lor_lhs_false1617
	}

land_lhs_true1614:
	v672 = *lookahead
	cmp1615 = v672 <= 90
	if cmp1615 {
		goto if_then1626
	} else {
		goto lor_lhs_false1617
	}

lor_lhs_false1617:
	v673 = *lookahead
	cmp1618 = v673 == 95
	if cmp1618 {
		goto if_then1626
	} else {
		goto lor_lhs_false1620
	}

lor_lhs_false1620:
	v674 = *lookahead
	cmp1621 = 97 <= v674
	if cmp1621 {
		goto land_lhs_true1623
	} else {
		goto if_end1627
	}

land_lhs_true1623:
	v675 = *lookahead
	cmp1624 = v675 <= 122
	if cmp1624 {
		goto if_then1626
	} else {
		goto if_end1627
	}

if_then1626:
	*state_addr = 130
	goto next_state

if_end1627:
	v676 = *result
	tobool1628 = (v676 & 1) != 0
	*retval = tobool1628
	goto _return

sw_bb1629:
	*result = 1
	v677 = *lexer_addr
	result_symbol1630 = &v677.F1
	*result_symbol1630 = 10
	v678 = *lexer_addr
	mark_end1631 = &v678.F3
	v679 = *mark_end1631
	v680 = *lexer_addr
	v679(v680)
	v681 = *lookahead
	cmp1632 = v681 == 111
	if cmp1632 {
		goto if_then1634
	} else {
		goto if_end1635
	}

if_then1634:
	*state_addr = 104
	goto next_state

if_end1635:
	v682 = *lookahead
	cmp1636 = v682 == 45
	if cmp1636 {
		goto if_then1662
	} else {
		goto lor_lhs_false1638
	}

lor_lhs_false1638:
	v683 = *lookahead
	cmp1639 = v683 == 46
	if cmp1639 {
		goto if_then1662
	} else {
		goto lor_lhs_false1641
	}

lor_lhs_false1641:
	v684 = *lookahead
	cmp1642 = 48 <= v684
	if cmp1642 {
		goto land_lhs_true1644
	} else {
		goto lor_lhs_false1647
	}

land_lhs_true1644:
	v685 = *lookahead
	cmp1645 = v685 <= 57
	if cmp1645 {
		goto if_then1662
	} else {
		goto lor_lhs_false1647
	}

lor_lhs_false1647:
	v686 = *lookahead
	cmp1648 = 65 <= v686
	if cmp1648 {
		goto land_lhs_true1650
	} else {
		goto lor_lhs_false1653
	}

land_lhs_true1650:
	v687 = *lookahead
	cmp1651 = v687 <= 90
	if cmp1651 {
		goto if_then1662
	} else {
		goto lor_lhs_false1653
	}

lor_lhs_false1653:
	v688 = *lookahead
	cmp1654 = v688 == 95
	if cmp1654 {
		goto if_then1662
	} else {
		goto lor_lhs_false1656
	}

lor_lhs_false1656:
	v689 = *lookahead
	cmp1657 = 97 <= v689
	if cmp1657 {
		goto land_lhs_true1659
	} else {
		goto if_end1663
	}

land_lhs_true1659:
	v690 = *lookahead
	cmp1660 = v690 <= 122
	if cmp1660 {
		goto if_then1662
	} else {
		goto if_end1663
	}

if_then1662:
	*state_addr = 130
	goto next_state

if_end1663:
	v691 = *result
	tobool1664 = (v691 & 1) != 0
	*retval = tobool1664
	goto _return

sw_bb1665:
	*result = 1
	v692 = *lexer_addr
	result_symbol1666 = &v692.F1
	*result_symbol1666 = 10
	v693 = *lexer_addr
	mark_end1667 = &v693.F3
	v694 = *mark_end1667
	v695 = *lexer_addr
	v694(v695)
	v696 = *lookahead
	cmp1668 = v696 == 111
	if cmp1668 {
		goto if_then1670
	} else {
		goto if_end1671
	}

if_then1670:
	*state_addr = 111
	goto next_state

if_end1671:
	v697 = *lookahead
	cmp1672 = v697 == 45
	if cmp1672 {
		goto if_then1698
	} else {
		goto lor_lhs_false1674
	}

lor_lhs_false1674:
	v698 = *lookahead
	cmp1675 = v698 == 46
	if cmp1675 {
		goto if_then1698
	} else {
		goto lor_lhs_false1677
	}

lor_lhs_false1677:
	v699 = *lookahead
	cmp1678 = 48 <= v699
	if cmp1678 {
		goto land_lhs_true1680
	} else {
		goto lor_lhs_false1683
	}

land_lhs_true1680:
	v700 = *lookahead
	cmp1681 = v700 <= 57
	if cmp1681 {
		goto if_then1698
	} else {
		goto lor_lhs_false1683
	}

lor_lhs_false1683:
	v701 = *lookahead
	cmp1684 = 65 <= v701
	if cmp1684 {
		goto land_lhs_true1686
	} else {
		goto lor_lhs_false1689
	}

land_lhs_true1686:
	v702 = *lookahead
	cmp1687 = v702 <= 90
	if cmp1687 {
		goto if_then1698
	} else {
		goto lor_lhs_false1689
	}

lor_lhs_false1689:
	v703 = *lookahead
	cmp1690 = v703 == 95
	if cmp1690 {
		goto if_then1698
	} else {
		goto lor_lhs_false1692
	}

lor_lhs_false1692:
	v704 = *lookahead
	cmp1693 = 97 <= v704
	if cmp1693 {
		goto land_lhs_true1695
	} else {
		goto if_end1699
	}

land_lhs_true1695:
	v705 = *lookahead
	cmp1696 = v705 <= 122
	if cmp1696 {
		goto if_then1698
	} else {
		goto if_end1699
	}

if_then1698:
	*state_addr = 130
	goto next_state

if_end1699:
	v706 = *result
	tobool1700 = (v706 & 1) != 0
	*retval = tobool1700
	goto _return

sw_bb1701:
	*result = 1
	v707 = *lexer_addr
	result_symbol1702 = &v707.F1
	*result_symbol1702 = 10
	v708 = *lexer_addr
	mark_end1703 = &v708.F3
	v709 = *mark_end1703
	v710 = *lexer_addr
	v709(v710)
	v711 = *lookahead
	cmp1704 = v711 == 112
	if cmp1704 {
		goto if_then1706
	} else {
		goto if_end1707
	}

if_then1706:
	*state_addr = 98
	goto next_state

if_end1707:
	v712 = *lookahead
	cmp1708 = v712 == 45
	if cmp1708 {
		goto if_then1734
	} else {
		goto lor_lhs_false1710
	}

lor_lhs_false1710:
	v713 = *lookahead
	cmp1711 = v713 == 46
	if cmp1711 {
		goto if_then1734
	} else {
		goto lor_lhs_false1713
	}

lor_lhs_false1713:
	v714 = *lookahead
	cmp1714 = 48 <= v714
	if cmp1714 {
		goto land_lhs_true1716
	} else {
		goto lor_lhs_false1719
	}

land_lhs_true1716:
	v715 = *lookahead
	cmp1717 = v715 <= 57
	if cmp1717 {
		goto if_then1734
	} else {
		goto lor_lhs_false1719
	}

lor_lhs_false1719:
	v716 = *lookahead
	cmp1720 = 65 <= v716
	if cmp1720 {
		goto land_lhs_true1722
	} else {
		goto lor_lhs_false1725
	}

land_lhs_true1722:
	v717 = *lookahead
	cmp1723 = v717 <= 90
	if cmp1723 {
		goto if_then1734
	} else {
		goto lor_lhs_false1725
	}

lor_lhs_false1725:
	v718 = *lookahead
	cmp1726 = v718 == 95
	if cmp1726 {
		goto if_then1734
	} else {
		goto lor_lhs_false1728
	}

lor_lhs_false1728:
	v719 = *lookahead
	cmp1729 = 97 <= v719
	if cmp1729 {
		goto land_lhs_true1731
	} else {
		goto if_end1735
	}

land_lhs_true1731:
	v720 = *lookahead
	cmp1732 = v720 <= 122
	if cmp1732 {
		goto if_then1734
	} else {
		goto if_end1735
	}

if_then1734:
	*state_addr = 130
	goto next_state

if_end1735:
	v721 = *result
	tobool1736 = (v721 & 1) != 0
	*retval = tobool1736
	goto _return

sw_bb1737:
	*result = 1
	v722 = *lexer_addr
	result_symbol1738 = &v722.F1
	*result_symbol1738 = 10
	v723 = *lexer_addr
	mark_end1739 = &v723.F3
	v724 = *mark_end1739
	v725 = *lexer_addr
	v724(v725)
	v726 = *lookahead
	cmp1740 = v726 == 114
	if cmp1740 {
		goto if_then1742
	} else {
		goto if_end1743
	}

if_then1742:
	*state_addr = 89
	goto next_state

if_end1743:
	v727 = *lookahead
	cmp1744 = v727 == 45
	if cmp1744 {
		goto if_then1770
	} else {
		goto lor_lhs_false1746
	}

lor_lhs_false1746:
	v728 = *lookahead
	cmp1747 = v728 == 46
	if cmp1747 {
		goto if_then1770
	} else {
		goto lor_lhs_false1749
	}

lor_lhs_false1749:
	v729 = *lookahead
	cmp1750 = 48 <= v729
	if cmp1750 {
		goto land_lhs_true1752
	} else {
		goto lor_lhs_false1755
	}

land_lhs_true1752:
	v730 = *lookahead
	cmp1753 = v730 <= 57
	if cmp1753 {
		goto if_then1770
	} else {
		goto lor_lhs_false1755
	}

lor_lhs_false1755:
	v731 = *lookahead
	cmp1756 = 65 <= v731
	if cmp1756 {
		goto land_lhs_true1758
	} else {
		goto lor_lhs_false1761
	}

land_lhs_true1758:
	v732 = *lookahead
	cmp1759 = v732 <= 90
	if cmp1759 {
		goto if_then1770
	} else {
		goto lor_lhs_false1761
	}

lor_lhs_false1761:
	v733 = *lookahead
	cmp1762 = v733 == 95
	if cmp1762 {
		goto if_then1770
	} else {
		goto lor_lhs_false1764
	}

lor_lhs_false1764:
	v734 = *lookahead
	cmp1765 = 97 <= v734
	if cmp1765 {
		goto land_lhs_true1767
	} else {
		goto if_end1771
	}

land_lhs_true1767:
	v735 = *lookahead
	cmp1768 = v735 <= 122
	if cmp1768 {
		goto if_then1770
	} else {
		goto if_end1771
	}

if_then1770:
	*state_addr = 130
	goto next_state

if_end1771:
	v736 = *result
	tobool1772 = (v736 & 1) != 0
	*retval = tobool1772
	goto _return

sw_bb1773:
	*result = 1
	v737 = *lexer_addr
	result_symbol1774 = &v737.F1
	*result_symbol1774 = 10
	v738 = *lexer_addr
	mark_end1775 = &v738.F3
	v739 = *mark_end1775
	v740 = *lexer_addr
	v739(v740)
	v741 = *lookahead
	cmp1776 = v741 == 114
	if cmp1776 {
		goto if_then1778
	} else {
		goto if_end1779
	}

if_then1778:
	*state_addr = 99
	goto next_state

if_end1779:
	v742 = *lookahead
	cmp1780 = v742 == 45
	if cmp1780 {
		goto if_then1806
	} else {
		goto lor_lhs_false1782
	}

lor_lhs_false1782:
	v743 = *lookahead
	cmp1783 = v743 == 46
	if cmp1783 {
		goto if_then1806
	} else {
		goto lor_lhs_false1785
	}

lor_lhs_false1785:
	v744 = *lookahead
	cmp1786 = 48 <= v744
	if cmp1786 {
		goto land_lhs_true1788
	} else {
		goto lor_lhs_false1791
	}

land_lhs_true1788:
	v745 = *lookahead
	cmp1789 = v745 <= 57
	if cmp1789 {
		goto if_then1806
	} else {
		goto lor_lhs_false1791
	}

lor_lhs_false1791:
	v746 = *lookahead
	cmp1792 = 65 <= v746
	if cmp1792 {
		goto land_lhs_true1794
	} else {
		goto lor_lhs_false1797
	}

land_lhs_true1794:
	v747 = *lookahead
	cmp1795 = v747 <= 90
	if cmp1795 {
		goto if_then1806
	} else {
		goto lor_lhs_false1797
	}

lor_lhs_false1797:
	v748 = *lookahead
	cmp1798 = v748 == 95
	if cmp1798 {
		goto if_then1806
	} else {
		goto lor_lhs_false1800
	}

lor_lhs_false1800:
	v749 = *lookahead
	cmp1801 = 97 <= v749
	if cmp1801 {
		goto land_lhs_true1803
	} else {
		goto if_end1807
	}

land_lhs_true1803:
	v750 = *lookahead
	cmp1804 = v750 <= 122
	if cmp1804 {
		goto if_then1806
	} else {
		goto if_end1807
	}

if_then1806:
	*state_addr = 130
	goto next_state

if_end1807:
	v751 = *result
	tobool1808 = (v751 & 1) != 0
	*retval = tobool1808
	goto _return

sw_bb1809:
	*result = 1
	v752 = *lexer_addr
	result_symbol1810 = &v752.F1
	*result_symbol1810 = 10
	v753 = *lexer_addr
	mark_end1811 = &v753.F3
	v754 = *mark_end1811
	v755 = *lexer_addr
	v754(v755)
	v756 = *lookahead
	cmp1812 = v756 == 114
	if cmp1812 {
		goto if_then1814
	} else {
		goto if_end1815
	}

if_then1814:
	*state_addr = 116
	goto next_state

if_end1815:
	v757 = *lookahead
	cmp1816 = v757 == 45
	if cmp1816 {
		goto if_then1842
	} else {
		goto lor_lhs_false1818
	}

lor_lhs_false1818:
	v758 = *lookahead
	cmp1819 = v758 == 46
	if cmp1819 {
		goto if_then1842
	} else {
		goto lor_lhs_false1821
	}

lor_lhs_false1821:
	v759 = *lookahead
	cmp1822 = 48 <= v759
	if cmp1822 {
		goto land_lhs_true1824
	} else {
		goto lor_lhs_false1827
	}

land_lhs_true1824:
	v760 = *lookahead
	cmp1825 = v760 <= 57
	if cmp1825 {
		goto if_then1842
	} else {
		goto lor_lhs_false1827
	}

lor_lhs_false1827:
	v761 = *lookahead
	cmp1828 = 65 <= v761
	if cmp1828 {
		goto land_lhs_true1830
	} else {
		goto lor_lhs_false1833
	}

land_lhs_true1830:
	v762 = *lookahead
	cmp1831 = v762 <= 90
	if cmp1831 {
		goto if_then1842
	} else {
		goto lor_lhs_false1833
	}

lor_lhs_false1833:
	v763 = *lookahead
	cmp1834 = v763 == 95
	if cmp1834 {
		goto if_then1842
	} else {
		goto lor_lhs_false1836
	}

lor_lhs_false1836:
	v764 = *lookahead
	cmp1837 = 97 <= v764
	if cmp1837 {
		goto land_lhs_true1839
	} else {
		goto if_end1843
	}

land_lhs_true1839:
	v765 = *lookahead
	cmp1840 = v765 <= 122
	if cmp1840 {
		goto if_then1842
	} else {
		goto if_end1843
	}

if_then1842:
	*state_addr = 130
	goto next_state

if_end1843:
	v766 = *result
	tobool1844 = (v766 & 1) != 0
	*retval = tobool1844
	goto _return

sw_bb1845:
	*result = 1
	v767 = *lexer_addr
	result_symbol1846 = &v767.F1
	*result_symbol1846 = 10
	v768 = *lexer_addr
	mark_end1847 = &v768.F3
	v769 = *mark_end1847
	v770 = *lexer_addr
	v769(v770)
	v771 = *lookahead
	cmp1848 = v771 == 114
	if cmp1848 {
		goto if_then1850
	} else {
		goto if_end1851
	}

if_then1850:
	*state_addr = 102
	goto next_state

if_end1851:
	v772 = *lookahead
	cmp1852 = v772 == 45
	if cmp1852 {
		goto if_then1878
	} else {
		goto lor_lhs_false1854
	}

lor_lhs_false1854:
	v773 = *lookahead
	cmp1855 = v773 == 46
	if cmp1855 {
		goto if_then1878
	} else {
		goto lor_lhs_false1857
	}

lor_lhs_false1857:
	v774 = *lookahead
	cmp1858 = 48 <= v774
	if cmp1858 {
		goto land_lhs_true1860
	} else {
		goto lor_lhs_false1863
	}

land_lhs_true1860:
	v775 = *lookahead
	cmp1861 = v775 <= 57
	if cmp1861 {
		goto if_then1878
	} else {
		goto lor_lhs_false1863
	}

lor_lhs_false1863:
	v776 = *lookahead
	cmp1864 = 65 <= v776
	if cmp1864 {
		goto land_lhs_true1866
	} else {
		goto lor_lhs_false1869
	}

land_lhs_true1866:
	v777 = *lookahead
	cmp1867 = v777 <= 90
	if cmp1867 {
		goto if_then1878
	} else {
		goto lor_lhs_false1869
	}

lor_lhs_false1869:
	v778 = *lookahead
	cmp1870 = v778 == 95
	if cmp1870 {
		goto if_then1878
	} else {
		goto lor_lhs_false1872
	}

lor_lhs_false1872:
	v779 = *lookahead
	cmp1873 = 97 <= v779
	if cmp1873 {
		goto land_lhs_true1875
	} else {
		goto if_end1879
	}

land_lhs_true1875:
	v780 = *lookahead
	cmp1876 = v780 <= 122
	if cmp1876 {
		goto if_then1878
	} else {
		goto if_end1879
	}

if_then1878:
	*state_addr = 130
	goto next_state

if_end1879:
	v781 = *result
	tobool1880 = (v781 & 1) != 0
	*retval = tobool1880
	goto _return

sw_bb1881:
	*result = 1
	v782 = *lexer_addr
	result_symbol1882 = &v782.F1
	*result_symbol1882 = 10
	v783 = *lexer_addr
	mark_end1883 = &v783.F3
	v784 = *mark_end1883
	v785 = *lexer_addr
	v784(v785)
	v786 = *lookahead
	cmp1884 = v786 == 115
	if cmp1884 {
		goto if_then1886
	} else {
		goto if_end1887
	}

if_then1886:
	*state_addr = 130
	goto next_state

if_end1887:
	v787 = *lookahead
	cmp1888 = v787 == 45
	if cmp1888 {
		goto if_then1914
	} else {
		goto lor_lhs_false1890
	}

lor_lhs_false1890:
	v788 = *lookahead
	cmp1891 = v788 == 46
	if cmp1891 {
		goto if_then1914
	} else {
		goto lor_lhs_false1893
	}

lor_lhs_false1893:
	v789 = *lookahead
	cmp1894 = 48 <= v789
	if cmp1894 {
		goto land_lhs_true1896
	} else {
		goto lor_lhs_false1899
	}

land_lhs_true1896:
	v790 = *lookahead
	cmp1897 = v790 <= 57
	if cmp1897 {
		goto if_then1914
	} else {
		goto lor_lhs_false1899
	}

lor_lhs_false1899:
	v791 = *lookahead
	cmp1900 = 65 <= v791
	if cmp1900 {
		goto land_lhs_true1902
	} else {
		goto lor_lhs_false1905
	}

land_lhs_true1902:
	v792 = *lookahead
	cmp1903 = v792 <= 90
	if cmp1903 {
		goto if_then1914
	} else {
		goto lor_lhs_false1905
	}

lor_lhs_false1905:
	v793 = *lookahead
	cmp1906 = v793 == 95
	if cmp1906 {
		goto if_then1914
	} else {
		goto lor_lhs_false1908
	}

lor_lhs_false1908:
	v794 = *lookahead
	cmp1909 = 97 <= v794
	if cmp1909 {
		goto land_lhs_true1911
	} else {
		goto if_end1915
	}

land_lhs_true1911:
	v795 = *lookahead
	cmp1912 = v795 <= 122
	if cmp1912 {
		goto if_then1914
	} else {
		goto if_end1915
	}

if_then1914:
	*state_addr = 130
	goto next_state

if_end1915:
	v796 = *result
	tobool1916 = (v796 & 1) != 0
	*retval = tobool1916
	goto _return

sw_bb1917:
	*result = 1
	v797 = *lexer_addr
	result_symbol1918 = &v797.F1
	*result_symbol1918 = 10
	v798 = *lexer_addr
	mark_end1919 = &v798.F3
	v799 = *mark_end1919
	v800 = *lexer_addr
	v799(v800)
	v801 = *lookahead
	cmp1920 = v801 == 115
	if cmp1920 {
		goto if_then1922
	} else {
		goto if_end1923
	}

if_then1922:
	*state_addr = 93
	goto next_state

if_end1923:
	v802 = *lookahead
	cmp1924 = v802 == 45
	if cmp1924 {
		goto if_then1950
	} else {
		goto lor_lhs_false1926
	}

lor_lhs_false1926:
	v803 = *lookahead
	cmp1927 = v803 == 46
	if cmp1927 {
		goto if_then1950
	} else {
		goto lor_lhs_false1929
	}

lor_lhs_false1929:
	v804 = *lookahead
	cmp1930 = 48 <= v804
	if cmp1930 {
		goto land_lhs_true1932
	} else {
		goto lor_lhs_false1935
	}

land_lhs_true1932:
	v805 = *lookahead
	cmp1933 = v805 <= 57
	if cmp1933 {
		goto if_then1950
	} else {
		goto lor_lhs_false1935
	}

lor_lhs_false1935:
	v806 = *lookahead
	cmp1936 = 65 <= v806
	if cmp1936 {
		goto land_lhs_true1938
	} else {
		goto lor_lhs_false1941
	}

land_lhs_true1938:
	v807 = *lookahead
	cmp1939 = v807 <= 90
	if cmp1939 {
		goto if_then1950
	} else {
		goto lor_lhs_false1941
	}

lor_lhs_false1941:
	v808 = *lookahead
	cmp1942 = v808 == 95
	if cmp1942 {
		goto if_then1950
	} else {
		goto lor_lhs_false1944
	}

lor_lhs_false1944:
	v809 = *lookahead
	cmp1945 = 97 <= v809
	if cmp1945 {
		goto land_lhs_true1947
	} else {
		goto if_end1951
	}

land_lhs_true1947:
	v810 = *lookahead
	cmp1948 = v810 <= 122
	if cmp1948 {
		goto if_then1950
	} else {
		goto if_end1951
	}

if_then1950:
	*state_addr = 130
	goto next_state

if_end1951:
	v811 = *result
	tobool1952 = (v811 & 1) != 0
	*retval = tobool1952
	goto _return

sw_bb1953:
	*result = 1
	v812 = *lexer_addr
	result_symbol1954 = &v812.F1
	*result_symbol1954 = 10
	v813 = *lexer_addr
	mark_end1955 = &v813.F3
	v814 = *mark_end1955
	v815 = *lexer_addr
	v814(v815)
	v816 = *lookahead
	cmp1956 = v816 == 116
	if cmp1956 {
		goto if_then1958
	} else {
		goto if_end1959
	}

if_then1958:
	*state_addr = 91
	goto next_state

if_end1959:
	v817 = *lookahead
	cmp1960 = v817 == 45
	if cmp1960 {
		goto if_then1986
	} else {
		goto lor_lhs_false1962
	}

lor_lhs_false1962:
	v818 = *lookahead
	cmp1963 = v818 == 46
	if cmp1963 {
		goto if_then1986
	} else {
		goto lor_lhs_false1965
	}

lor_lhs_false1965:
	v819 = *lookahead
	cmp1966 = 48 <= v819
	if cmp1966 {
		goto land_lhs_true1968
	} else {
		goto lor_lhs_false1971
	}

land_lhs_true1968:
	v820 = *lookahead
	cmp1969 = v820 <= 57
	if cmp1969 {
		goto if_then1986
	} else {
		goto lor_lhs_false1971
	}

lor_lhs_false1971:
	v821 = *lookahead
	cmp1972 = 65 <= v821
	if cmp1972 {
		goto land_lhs_true1974
	} else {
		goto lor_lhs_false1977
	}

land_lhs_true1974:
	v822 = *lookahead
	cmp1975 = v822 <= 90
	if cmp1975 {
		goto if_then1986
	} else {
		goto lor_lhs_false1977
	}

lor_lhs_false1977:
	v823 = *lookahead
	cmp1978 = v823 == 95
	if cmp1978 {
		goto if_then1986
	} else {
		goto lor_lhs_false1980
	}

lor_lhs_false1980:
	v824 = *lookahead
	cmp1981 = 97 <= v824
	if cmp1981 {
		goto land_lhs_true1983
	} else {
		goto if_end1987
	}

land_lhs_true1983:
	v825 = *lookahead
	cmp1984 = v825 <= 122
	if cmp1984 {
		goto if_then1986
	} else {
		goto if_end1987
	}

if_then1986:
	*state_addr = 130
	goto next_state

if_end1987:
	v826 = *result
	tobool1988 = (v826 & 1) != 0
	*retval = tobool1988
	goto _return

sw_bb1989:
	*result = 1
	v827 = *lexer_addr
	result_symbol1990 = &v827.F1
	*result_symbol1990 = 10
	v828 = *lexer_addr
	mark_end1991 = &v828.F3
	v829 = *mark_end1991
	v830 = *lexer_addr
	v829(v830)
	v831 = *lookahead
	cmp1992 = v831 == 117
	if cmp1992 {
		goto if_then1994
	} else {
		goto if_end1995
	}

if_then1994:
	*state_addr = 128
	goto next_state

if_end1995:
	v832 = *lookahead
	cmp1996 = v832 == 45
	if cmp1996 {
		goto if_then2022
	} else {
		goto lor_lhs_false1998
	}

lor_lhs_false1998:
	v833 = *lookahead
	cmp1999 = v833 == 46
	if cmp1999 {
		goto if_then2022
	} else {
		goto lor_lhs_false2001
	}

lor_lhs_false2001:
	v834 = *lookahead
	cmp2002 = 48 <= v834
	if cmp2002 {
		goto land_lhs_true2004
	} else {
		goto lor_lhs_false2007
	}

land_lhs_true2004:
	v835 = *lookahead
	cmp2005 = v835 <= 57
	if cmp2005 {
		goto if_then2022
	} else {
		goto lor_lhs_false2007
	}

lor_lhs_false2007:
	v836 = *lookahead
	cmp2008 = 65 <= v836
	if cmp2008 {
		goto land_lhs_true2010
	} else {
		goto lor_lhs_false2013
	}

land_lhs_true2010:
	v837 = *lookahead
	cmp2011 = v837 <= 90
	if cmp2011 {
		goto if_then2022
	} else {
		goto lor_lhs_false2013
	}

lor_lhs_false2013:
	v838 = *lookahead
	cmp2014 = v838 == 95
	if cmp2014 {
		goto if_then2022
	} else {
		goto lor_lhs_false2016
	}

lor_lhs_false2016:
	v839 = *lookahead
	cmp2017 = 97 <= v839
	if cmp2017 {
		goto land_lhs_true2019
	} else {
		goto if_end2023
	}

land_lhs_true2019:
	v840 = *lookahead
	cmp2020 = v840 <= 122
	if cmp2020 {
		goto if_then2022
	} else {
		goto if_end2023
	}

if_then2022:
	*state_addr = 130
	goto next_state

if_end2023:
	v841 = *result
	tobool2024 = (v841 & 1) != 0
	*retval = tobool2024
	goto _return

sw_bb2025:
	*result = 1
	v842 = *lexer_addr
	result_symbol2026 = &v842.F1
	*result_symbol2026 = 10
	v843 = *lexer_addr
	mark_end2027 = &v843.F3
	v844 = *mark_end2027
	v845 = *lexer_addr
	v844(v845)
	v846 = *lookahead
	cmp2028 = v846 == 119
	if cmp2028 {
		goto if_then2030
	} else {
		goto if_end2031
	}

if_then2030:
	*state_addr = 123
	goto next_state

if_end2031:
	v847 = *lookahead
	cmp2032 = v847 == 45
	if cmp2032 {
		goto if_then2058
	} else {
		goto lor_lhs_false2034
	}

lor_lhs_false2034:
	v848 = *lookahead
	cmp2035 = v848 == 46
	if cmp2035 {
		goto if_then2058
	} else {
		goto lor_lhs_false2037
	}

lor_lhs_false2037:
	v849 = *lookahead
	cmp2038 = 48 <= v849
	if cmp2038 {
		goto land_lhs_true2040
	} else {
		goto lor_lhs_false2043
	}

land_lhs_true2040:
	v850 = *lookahead
	cmp2041 = v850 <= 57
	if cmp2041 {
		goto if_then2058
	} else {
		goto lor_lhs_false2043
	}

lor_lhs_false2043:
	v851 = *lookahead
	cmp2044 = 65 <= v851
	if cmp2044 {
		goto land_lhs_true2046
	} else {
		goto lor_lhs_false2049
	}

land_lhs_true2046:
	v852 = *lookahead
	cmp2047 = v852 <= 90
	if cmp2047 {
		goto if_then2058
	} else {
		goto lor_lhs_false2049
	}

lor_lhs_false2049:
	v853 = *lookahead
	cmp2050 = v853 == 95
	if cmp2050 {
		goto if_then2058
	} else {
		goto lor_lhs_false2052
	}

lor_lhs_false2052:
	v854 = *lookahead
	cmp2053 = 97 <= v854
	if cmp2053 {
		goto land_lhs_true2055
	} else {
		goto if_end2059
	}

land_lhs_true2055:
	v855 = *lookahead
	cmp2056 = v855 <= 122
	if cmp2056 {
		goto if_then2058
	} else {
		goto if_end2059
	}

if_then2058:
	*state_addr = 130
	goto next_state

if_end2059:
	v856 = *result
	tobool2060 = (v856 & 1) != 0
	*retval = tobool2060
	goto _return

sw_bb2061:
	*result = 1
	v857 = *lexer_addr
	result_symbol2062 = &v857.F1
	*result_symbol2062 = 10
	v858 = *lexer_addr
	mark_end2063 = &v858.F3
	v859 = *mark_end2063
	v860 = *lexer_addr
	v859(v860)
	v861 = *lookahead
	cmp2064 = v861 == 120
	if cmp2064 {
		goto if_then2066
	} else {
		goto if_end2067
	}

if_then2066:
	*state_addr = 130
	goto next_state

if_end2067:
	v862 = *lookahead
	cmp2068 = v862 == 45
	if cmp2068 {
		goto if_then2094
	} else {
		goto lor_lhs_false2070
	}

lor_lhs_false2070:
	v863 = *lookahead
	cmp2071 = v863 == 46
	if cmp2071 {
		goto if_then2094
	} else {
		goto lor_lhs_false2073
	}

lor_lhs_false2073:
	v864 = *lookahead
	cmp2074 = 48 <= v864
	if cmp2074 {
		goto land_lhs_true2076
	} else {
		goto lor_lhs_false2079
	}

land_lhs_true2076:
	v865 = *lookahead
	cmp2077 = v865 <= 57
	if cmp2077 {
		goto if_then2094
	} else {
		goto lor_lhs_false2079
	}

lor_lhs_false2079:
	v866 = *lookahead
	cmp2080 = 65 <= v866
	if cmp2080 {
		goto land_lhs_true2082
	} else {
		goto lor_lhs_false2085
	}

land_lhs_true2082:
	v867 = *lookahead
	cmp2083 = v867 <= 90
	if cmp2083 {
		goto if_then2094
	} else {
		goto lor_lhs_false2085
	}

lor_lhs_false2085:
	v868 = *lookahead
	cmp2086 = v868 == 95
	if cmp2086 {
		goto if_then2094
	} else {
		goto lor_lhs_false2088
	}

lor_lhs_false2088:
	v869 = *lookahead
	cmp2089 = 97 <= v869
	if cmp2089 {
		goto land_lhs_true2091
	} else {
		goto if_end2095
	}

land_lhs_true2091:
	v870 = *lookahead
	cmp2092 = v870 <= 122
	if cmp2092 {
		goto if_then2094
	} else {
		goto if_end2095
	}

if_then2094:
	*state_addr = 130
	goto next_state

if_end2095:
	v871 = *result
	tobool2096 = (v871 & 1) != 0
	*retval = tobool2096
	goto _return

sw_bb2097:
	*result = 1
	v872 = *lexer_addr
	result_symbol2098 = &v872.F1
	*result_symbol2098 = 10
	v873 = *lexer_addr
	mark_end2099 = &v873.F3
	v874 = *mark_end2099
	v875 = *lexer_addr
	v874(v875)
	v876 = *lookahead
	cmp2100 = v876 == 121
	if cmp2100 {
		goto if_then2102
	} else {
		goto if_end2103
	}

if_then2102:
	*state_addr = 130
	goto next_state

if_end2103:
	v877 = *lookahead
	cmp2104 = v877 == 45
	if cmp2104 {
		goto if_then2130
	} else {
		goto lor_lhs_false2106
	}

lor_lhs_false2106:
	v878 = *lookahead
	cmp2107 = v878 == 46
	if cmp2107 {
		goto if_then2130
	} else {
		goto lor_lhs_false2109
	}

lor_lhs_false2109:
	v879 = *lookahead
	cmp2110 = 48 <= v879
	if cmp2110 {
		goto land_lhs_true2112
	} else {
		goto lor_lhs_false2115
	}

land_lhs_true2112:
	v880 = *lookahead
	cmp2113 = v880 <= 57
	if cmp2113 {
		goto if_then2130
	} else {
		goto lor_lhs_false2115
	}

lor_lhs_false2115:
	v881 = *lookahead
	cmp2116 = 65 <= v881
	if cmp2116 {
		goto land_lhs_true2118
	} else {
		goto lor_lhs_false2121
	}

land_lhs_true2118:
	v882 = *lookahead
	cmp2119 = v882 <= 90
	if cmp2119 {
		goto if_then2130
	} else {
		goto lor_lhs_false2121
	}

lor_lhs_false2121:
	v883 = *lookahead
	cmp2122 = v883 == 95
	if cmp2122 {
		goto if_then2130
	} else {
		goto lor_lhs_false2124
	}

lor_lhs_false2124:
	v884 = *lookahead
	cmp2125 = 97 <= v884
	if cmp2125 {
		goto land_lhs_true2127
	} else {
		goto if_end2131
	}

land_lhs_true2127:
	v885 = *lookahead
	cmp2128 = v885 <= 122
	if cmp2128 {
		goto if_then2130
	} else {
		goto if_end2131
	}

if_then2130:
	*state_addr = 130
	goto next_state

if_end2131:
	v886 = *result
	tobool2132 = (v886 & 1) != 0
	*retval = tobool2132
	goto _return

sw_bb2133:
	*result = 1
	v887 = *lexer_addr
	result_symbol2134 = &v887.F1
	*result_symbol2134 = 10
	v888 = *lexer_addr
	mark_end2135 = &v888.F3
	v889 = *mark_end2135
	v890 = *lexer_addr
	v889(v890)
	v891 = *lookahead
	cmp2136 = v891 == 45
	if cmp2136 {
		goto if_then2162
	} else {
		goto lor_lhs_false2138
	}

lor_lhs_false2138:
	v892 = *lookahead
	cmp2139 = v892 == 46
	if cmp2139 {
		goto if_then2162
	} else {
		goto lor_lhs_false2141
	}

lor_lhs_false2141:
	v893 = *lookahead
	cmp2142 = 48 <= v893
	if cmp2142 {
		goto land_lhs_true2144
	} else {
		goto lor_lhs_false2147
	}

land_lhs_true2144:
	v894 = *lookahead
	cmp2145 = v894 <= 57
	if cmp2145 {
		goto if_then2162
	} else {
		goto lor_lhs_false2147
	}

lor_lhs_false2147:
	v895 = *lookahead
	cmp2148 = 65 <= v895
	if cmp2148 {
		goto land_lhs_true2150
	} else {
		goto lor_lhs_false2153
	}

land_lhs_true2150:
	v896 = *lookahead
	cmp2151 = v896 <= 90
	if cmp2151 {
		goto if_then2162
	} else {
		goto lor_lhs_false2153
	}

lor_lhs_false2153:
	v897 = *lookahead
	cmp2154 = v897 == 95
	if cmp2154 {
		goto if_then2162
	} else {
		goto lor_lhs_false2156
	}

lor_lhs_false2156:
	v898 = *lookahead
	cmp2157 = 97 <= v898
	if cmp2157 {
		goto land_lhs_true2159
	} else {
		goto if_end2163
	}

land_lhs_true2159:
	v899 = *lookahead
	cmp2160 = v899 <= 122
	if cmp2160 {
		goto if_then2162
	} else {
		goto if_end2163
	}

if_then2162:
	*state_addr = 130
	goto next_state

if_end2163:
	v900 = *result
	tobool2164 = (v900 & 1) != 0
	*retval = tobool2164
	goto _return

sw_bb2165:
	*result = 1
	v901 = *lexer_addr
	result_symbol2166 = &v901.F1
	*result_symbol2166 = 11
	v902 = *lexer_addr
	mark_end2167 = &v902.F3
	v903 = *mark_end2167
	v904 = *lexer_addr
	v903(v904)
	v905 = *result
	tobool2168 = (v905 & 1) != 0
	*retval = tobool2168
	goto _return

sw_bb2169:
	*result = 1
	v906 = *lexer_addr
	result_symbol2170 = &v906.F1
	*result_symbol2170 = 12
	v907 = *lexer_addr
	mark_end2171 = &v907.F3
	v908 = *mark_end2171
	v909 = *lexer_addr
	v908(v909)
	v910 = *result
	tobool2172 = (v910 & 1) != 0
	*retval = tobool2172
	goto _return

sw_bb2173:
	*result = 1
	v911 = *lexer_addr
	result_symbol2174 = &v911.F1
	*result_symbol2174 = 12
	v912 = *lexer_addr
	mark_end2175 = &v912.F3
	v913 = *mark_end2175
	v914 = *lexer_addr
	v913(v914)
	v915 = *lookahead
	cmp2176 = v915 == 10
	if cmp2176 {
		goto if_then2178
	} else {
		goto if_end2179
	}

if_then2178:
	*state_addr = 132
	goto next_state

if_end2179:
	v916 = *result
	tobool2180 = (v916 & 1) != 0
	*retval = tobool2180
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v917 = *retval
	return v917
}

func advance(lexer *TSLexer) {
	var lexer_addr **TSLexer
	var v0, v2 *TSLexer
	var local_advance *func(*TSLexer, bool)
	var v1 func(*TSLexer, bool)

	_, _, _, _, _ = lexer_addr, v0, local_advance, v1, v2

	lexer_addr = new(*TSLexer)
	*lexer_addr = lexer
	v0 = *lexer_addr
	local_advance = &v0.F2
	v1 = *local_advance
	v2 = *lexer_addr
	v1(v2, false)
}

