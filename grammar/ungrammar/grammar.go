package grammar_ungrammar

import "unsafe"

type TSFieldMapEntry struct {
	F0 int16
	F1 byte
	F2 byte
}

type TSFieldMapSlice struct {
	F0 int16
	F1 int16
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
	F16 *TSFieldMapSlice
	F17 *TSFieldMapEntry
	F18 *TSSymbolMetadata
	F19 *int16
	F20 *int16
	F21 *int16
	F22 *TSLexMode
	F23 func(*TSLexer, int16) bool
	F24 func(*TSLexer, int16) bool
	F25 int16
	F26 anon_2
	F27 *int16
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

var tree_sitter_ungrammar_language TSLanguage = TSLanguage{14, 26, 2, 12, 0, 33, 5, 3, 0, 3, &(*[5][26]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[90]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, ts_lex_keywords, 1, anon_2{}, &ts_primary_state_ids[0]}

var ts_small_parse_table [407]int16 = [407]int16{
	7, 3, 1, 11, 11, 1, 3, 13, 1, 7, 28, 1, 1, 3, 1, 24,
	26, 2, 0, 9, 11, 6, 15, 16, 17, 18, 19, 20, 7, 3, 1, 11,
	9, 1, 1, 11, 1, 3, 13, 1, 7, 3, 1, 24, 26, 2, 4, 9,
	11, 6, 15, 16, 17, 18, 19, 20, 7, 3, 1, 11, 9, 1, 1, 11,
	1, 3, 13, 1, 7, 5, 1, 24, 21, 1, 21, 11, 6, 15, 16, 17,
	18, 19, 20, 7, 3, 1, 11, 9, 1, 1, 11, 1, 3, 13, 1, 7,
	6, 1, 24, 21, 1, 21, 11, 6, 15, 16, 17, 18, 19, 20, 5, 3,
	1, 11, 9, 1, 1, 11, 1, 3, 13, 1, 7, 16, 6, 15, 16, 17,
	18, 19, 20, 3, 3, 1, 11, 33, 1, 8, 31, 8, 0, 3, 4, 5,
	6, 7, 9, 1, 4, 3, 1, 11, 37, 1, 5, 39, 1, 6, 35, 6,
	0, 3, 4, 7, 9, 1, 2, 3, 1, 11, 41, 8, 0, 3, 4, 5,
	6, 7, 9, 1, 2, 3, 1, 11, 43, 8, 0, 3, 4, 5, 6, 7,
	9, 1, 2, 3, 1, 11, 45, 8, 0, 3, 4, 5, 6, 7, 9, 1,
	2, 3, 1, 11, 47, 8, 0, 3, 4, 5, 6, 7, 9, 1, 2, 3,
	1, 11, 49, 8, 0, 3, 4, 5, 6, 7, 9, 1, 4, 3, 1, 11,
	53, 1, 9, 20, 1, 25, 51, 2, 0, 1, 4, 3, 1, 11, 7, 1,
	1, 55, 1, 0, 19, 2, 13, 23, 4, 3, 1, 11, 57, 1, 0, 59,
	1, 1, 19, 2, 13, 23, 4, 3, 1, 11, 53, 1, 9, 22, 1, 25,
	62, 2, 0, 1, 2, 3, 1, 11, 64, 4, 0, 4, 9, 1, 4, 3,
	1, 11, 66, 1, 9, 22, 1, 25, 64, 2, 0, 1, 2, 3, 1, 11,
	69, 3, 0, 4, 1, 4, 3, 1, 11, 62, 1, 4, 71, 1, 9, 25,
	1, 25, 4, 3, 1, 11, 64, 1, 4, 73, 1, 9, 25, 1, 25, 4,
	3, 1, 11, 51, 1, 4, 71, 1, 9, 24, 1, 25, 2, 3, 1, 11,
	76, 2, 0, 1, 2, 78, 1, 10, 80, 1, 11, 2, 3, 1, 11, 82,
	1, 2, 2, 3, 1, 11, 84, 1, 7, 2, 3, 1, 11, 86, 1, 4,
	2, 3, 1, 11, 88, 1, 0,
}

var ts_small_parse_table_map [28]int32 = [28]int32{
	0, 28, 56, 83, 110, 131, 148, 166, 180, 194, 208, 222, 236, 250, 264, 278,
	292, 302, 316, 325, 338, 351, 364, 372, 379, 386, 393, 400,
}

var ts_symbol_names [28]*byte = [28]*byte{
	&_str[0], &_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0], &_str_17[0],
	&_str_18[0], &_str_19[0], &_str_20[0], &_str_21[0], &_str_22[0], &_str_23[0], &_str_24[0], &_str_25[0], &_str_26[0], &_str_27[0], &_str_28[0], &_str_29[0],
}

var ts_symbol_metadata [28]TSSymbolMetadata = [28]TSSymbolMetadata{
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0},
	TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0},
}

var ts_symbol_map [28]int16 = [28]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27,
}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [3][3]int16 = [3][3]int16{[3]int16{}, [3]int16{26, 0, 0}, [3]int16{27, 0, 0}}

var ts_lex_modes [33]TSLexMode = [33]TSLexMode{
	TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{},
	TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{}, TSLexMode{2, 0}, TSLexMode{}, TSLexMode{}, TSLexMode{},
	TSLexMode{},
}

var ts_primary_state_ids [33]int16 = [33]int16{
	0, 1, 2, 3, 4, 5, 5, 7, 7, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 20, 22, 17, 27, 28, 29, 30, 31,
	32,
}

var ts_parse_table struct {
	F0 struct {
	F0 [12]int16
	F1 [14]int16
}
	F1 [26]int16
	F2 [26]int16
	F3 [26]int16
	F4 [26]int16
} = struct {
	F0 struct {
	F0 [12]int16
	F1 [14]int16
}
	F1 [26]int16
	F2 [26]int16
	F3 [26]int16
	F4 [26]int16
}{struct {
	F0 [12]int16
	F1 [14]int16
}{[12]int16{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 3}, [14]int16{}}, [26]int16{
	5, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 32, 18, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 18, 0, 0,
}, [26]int16{
	0, 9, 0, 11, 0, 0, 0, 13, 0, 0, 0, 3, 0, 0, 27, 11,
	11, 11, 11, 11, 11, 17, 23, 0, 5, 0,
}, [26]int16{
	15, 17, 0, 20, 15, 0, 0, 23, 0, 15, 0, 3, 0, 0, 0, 11,
	11, 11, 11, 11, 11, 0, 0, 0, 3, 0,
}, [26]int16{
	0, 9, 0, 11, 0, 0, 0, 13, 0, 0, 0, 3, 0, 0, 31, 11,
	11, 11, 11, 11, 11, 26, 23, 0, 6, 0,
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
	F6 TSParseActionEntry
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
	F16 TSParseActionEntry
	F17 struct {
	F0 anon_1
	F1 [6]byte
}
	F18 TSParseActionEntry
	F19 struct {
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
	F20 struct {
	F0 anon_1
	F1 [6]byte
}
	F21 TSParseActionEntry
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
	F26 struct {
	F0 anon_1
	F1 [6]byte
}
	F27 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F34 struct {
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
	F40 struct {
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
	F48 TSParseActionEntry
	F49 struct {
	F0 anon_1
	F1 [6]byte
}
	F50 TSParseActionEntry
	F51 struct {
	F0 anon_1
	F1 [6]byte
}
	F52 TSParseActionEntry
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
	F56 TSParseActionEntry
	F57 struct {
	F0 anon_1
	F1 [6]byte
}
	F58 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F67 TSParseActionEntry
	F68 struct {
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
	F74 TSParseActionEntry
	F75 struct {
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
	F76 struct {
	F0 anon_1
	F1 [6]byte
}
	F77 TSParseActionEntry
	F78 struct {
	F0 anon_1
	F1 [6]byte
}
	F79 struct {
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
	F80 struct {
	F0 anon_1
	F1 [6]byte
}
	F81 struct {
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
	F82 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F85 struct {
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
	F86 struct {
	F0 anon_1
	F1 [6]byte
}
	F87 struct {
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
	F88 struct {
	F0 anon_1
	F1 [6]byte
}
	F89 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
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
	F6 TSParseActionEntry
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
	F16 TSParseActionEntry
	F17 struct {
	F0 anon_1
	F1 [6]byte
}
	F18 TSParseActionEntry
	F19 struct {
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
	F20 struct {
	F0 anon_1
	F1 [6]byte
}
	F21 TSParseActionEntry
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
	F26 struct {
	F0 anon_1
	F1 [6]byte
}
	F27 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F34 struct {
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
	F40 struct {
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
	F48 TSParseActionEntry
	F49 struct {
	F0 anon_1
	F1 [6]byte
}
	F50 TSParseActionEntry
	F51 struct {
	F0 anon_1
	F1 [6]byte
}
	F52 TSParseActionEntry
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
	F56 TSParseActionEntry
	F57 struct {
	F0 anon_1
	F1 [6]byte
}
	F58 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F67 TSParseActionEntry
	F68 struct {
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
	F74 TSParseActionEntry
	F75 struct {
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
	F76 struct {
	F0 anon_1
	F1 [6]byte
}
	F77 TSParseActionEntry
	F78 struct {
	F0 anon_1
	F1 [6]byte
}
	F79 struct {
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
	F80 struct {
	F0 anon_1
	F1 [6]byte
}
	F81 struct {
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
	F82 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F85 struct {
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
	F86 struct {
	F0 anon_1
	F1 [6]byte
}
	F87 struct {
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
	F88 struct {
	F0 anon_1
	F1 [6]byte
}
	F89 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 12, 0, 0}}}, struct {
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
}{0, 29, 0, 0}, [2]byte{}}}, struct {
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
}{0, 4, 0, 0}, [2]byte{}}}, struct {
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
}{0, 10, 0, 1}, [2]byte{}}}, struct {
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
}{0, 4, 0, 1}, [2]byte{}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 18, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 24, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 19, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 20, 0, 2}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 22, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 12, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 23, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 23, 0, 0}}}, struct {
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
}{0, 29, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 22, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 25, 0, 0}}}, struct {
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
}{0, 7, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 14, 0, 0}}}, struct {
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
}{0, 8, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 13, 0, 1}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, struct {
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
}{anon_1{1, 0}, [6]byte{}}, struct {
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
}{0, 14, 0, 0}, [2]byte{}}}, struct {
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
}{2, [7]byte{}}}}

var _str [4]byte = [4]byte{101, 110, 100, 0}

var _str_3 [11]byte = [11]byte{105, 100, 101, 110, 116, 105, 102, 105, 101, 114, 0}

var _str_4 [2]byte = [2]byte{61, 0}

var _str_5 [2]byte = [2]byte{40, 0}

var _str_6 [2]byte = [2]byte{41, 0}

var _str_7 [2]byte = [2]byte{42, 0}

var _str_8 [2]byte = [2]byte{63, 0}

var _str_9 [2]byte = [2]byte{39, 0}

var _str_10 [2]byte = [2]byte{58, 0}

var _str_11 [2]byte = [2]byte{124, 0}

var _str_12 [17]byte = [17]byte{
	116, 111, 107, 101, 110, 95, 105, 100, 101, 110, 116, 105, 102, 105, 101, 114,
	0,
}

var _str_13 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}

var _str_14 [8]byte = [8]byte{103, 114, 97, 109, 109, 97, 114, 0}

var _str_15 [5]byte = [5]byte{110, 111, 100, 101, 0}

var _str_16 [5]byte = [5]byte{114, 117, 108, 101, 0}

var _str_17 [11]byte = [11]byte{95, 97, 116, 111, 109, 95, 114, 117, 108, 101, 0}

var _str_18 [11]byte = [11]byte{114, 101, 112, 101, 116, 105, 116, 105, 111, 110, 0}

var _str_19 [9]byte = [9]byte{111, 112, 116, 105, 111, 110, 97, 108, 0}

var _str_20 [10]byte = [10]byte{110, 111, 100, 101, 95, 114, 117, 108, 101, 0}

var _str_21 [6]byte = [6]byte{116, 111, 107, 101, 110, 0}

var _str_22 [6]byte = [6]byte{108, 97, 98, 101, 108, 0}

var _str_23 [9]byte = [9]byte{115, 101, 113, 117, 101, 110, 99, 101, 0}

var _str_24 [12]byte = [12]byte{97, 108, 116, 101, 114, 110, 97, 116, 105, 111, 110, 0}

var _str_25 [16]byte = [16]byte{
	103, 114, 97, 109, 109, 97, 114, 95, 114, 101, 112, 101, 97, 116, 49, 0,
}

var _str_26 [17]byte = [17]byte{
	115, 101, 113, 117, 101, 110, 99, 101, 95, 114, 101, 112, 101, 97, 116, 49,
	0,
}

var _str_27 [20]byte = [20]byte{
	97, 108, 116, 101, 114, 110, 97, 116, 105, 111, 110, 95, 114, 101, 112, 101,
	97, 116, 49, 0,
}

var _str_28 [11]byte = [11]byte{100, 101, 102, 105, 110, 105, 116, 105, 111, 110, 0}

var _str_29 [11]byte = [11]byte{108, 97, 98, 101, 108, 95, 110, 97, 109, 101, 0}

func tree_sitter_ungrammar() *TSLanguage {
	return &tree_sitter_ungrammar_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v44, v45, v47, v49, v50, v52, v54, v55, v57, v59, v60, v62, v64, v65, v67, v69, v70, v72, v74, v75, v77, v79, v80, v82, v84, v85, v87, v89, v90, v92, v99, v100, v102, v108, v109, v111, v121, v122, v124, v130, v131, v133, v138, v139, v141, v147, v148, v150 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end116, mark_end120, mark_end124, mark_end128, mark_end132, mark_end136, mark_end140, mark_end144, mark_end148, mark_end168, mark_end188, mark_end220, mark_end239, mark_end254, mark_end272 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, result_symbol, result_symbol115, result_symbol119, result_symbol123, result_symbol127, result_symbol131, result_symbol135, result_symbol139, result_symbol143, result_symbol147, result_symbol167, result_symbol187, result_symbol219, result_symbol238, result_symbol253, result_symbol271 *int16
	var lookahead, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp27, cmp31, cmp35, cmp39, cmp41, cmp44, cmp47, cmp51, cmp53, cmp56, cmp59, cmp62, tobool66, cmp68, tobool72, cmp74, cmp78, cmp82, cmp85, cmp88, cmp91, cmp95, cmp98, tobool102, cmp104, cmp107, tobool111, tobool113, tobool117, tobool121, tobool125, tobool129, tobool133, tobool137, tobool141, tobool145, cmp149, cmp152, cmp155, cmp158, cmp161, tobool165, cmp169, cmp173, cmp177, cmp181, tobool185, cmp189, cmp193, cmp197, cmp200, cmp203, cmp206, cmp210, cmp213, tobool217, cmp221, cmp225, cmp229, cmp232, tobool236, cmp240, cmp244, cmp247, tobool251, cmp255, cmp258, cmp262, cmp265, tobool269, cmp273, cmp276, tobool280, v154 bool
	var v3, frombool, v10, v29, v31, v40, v43, v48, v53, v58, v63, v68, v73, v78, v83, v88, v98, v107, v120, v129, v137, v146, v153 byte
	var v46, v51, v56, v61, v66, v71, v76, v81, v86, v91, v101, v110, v123, v132, v140, v149 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9 int16
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v30, v32, v33, v34, v35, v36, v37, v38, v39, v41, v42, v93, v94, v95, v96, v97, v103, v104, v105, v106, v112, v113, v114, v115, v116, v117, v118, v119, v125, v126, v127, v128, v134, v135, v136, v142, v143, v144, v145, v151, v152 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp27, v18, cmp31, v19, cmp35, v20, cmp39, v21, cmp41, v22, cmp44, v23, cmp47, v24, cmp51, v25, cmp53, v26, cmp56, v27, cmp59, v28, cmp62, v29, tobool66, v30, cmp68, v31, tobool72, v32, cmp74, v33, cmp78, v34, cmp82, v35, cmp85, v36, cmp88, v37, cmp91, v38, cmp95, v39, cmp98, v40, tobool102, v41, cmp104, v42, cmp107, v43, tobool111, v44, result_symbol, v45, mark_end, v46, v47, v48, tobool113, v49, result_symbol115, v50, mark_end116, v51, v52, v53, tobool117, v54, result_symbol119, v55, mark_end120, v56, v57, v58, tobool121, v59, result_symbol123, v60, mark_end124, v61, v62, v63, tobool125, v64, result_symbol127, v65, mark_end128, v66, v67, v68, tobool129, v69, result_symbol131, v70, mark_end132, v71, v72, v73, tobool133, v74, result_symbol135, v75, mark_end136, v76, v77, v78, tobool137, v79, result_symbol139, v80, mark_end140, v81, v82, v83, tobool141, v84, result_symbol143, v85, mark_end144, v86, v87, v88, tobool145, v89, result_symbol147, v90, mark_end148, v91, v92, v93, cmp149, v94, cmp152, v95, cmp155, v96, cmp158, v97, cmp161, v98, tobool165, v99, result_symbol167, v100, mark_end168, v101, v102, v103, cmp169, v104, cmp173, v105, cmp177, v106, cmp181, v107, tobool185, v108, result_symbol187, v109, mark_end188, v110, v111, v112, cmp189, v113, cmp193, v114, cmp197, v115, cmp200, v116, cmp203, v117, cmp206, v118, cmp210, v119, cmp213, v120, tobool217, v121, result_symbol219, v122, mark_end220, v123, v124, v125, cmp221, v126, cmp225, v127, cmp229, v128, cmp232, v129, tobool236, v130, result_symbol238, v131, mark_end239, v132, v133, v134, cmp240, v135, cmp244, v136, cmp247, v137, tobool251, v138, result_symbol253, v139, mark_end254, v140, v141, v142, cmp255, v143, cmp258, v144, cmp262, v145, cmp265, v146, tobool269, v147, result_symbol271, v148, mark_end272, v149, v150, v151, cmp273, v152, cmp276, v153, tobool280, v154

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
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
	if call { frombool = 1 } else { frombool = 0 }
	*eof = frombool
	v9 = *state_addr
	conv = int32(uint32(uint16(v9)))
	switch conv {
	case 0:
		goto sw_bb
	case 1:
		goto sw_bb67
	case 2:
		goto sw_bb73
	case 3:
		goto sw_bb103
	case 4:
		goto sw_bb112
	case 5:
		goto sw_bb114
	case 6:
		goto sw_bb118
	case 7:
		goto sw_bb122
	case 8:
		goto sw_bb126
	case 9:
		goto sw_bb130
	case 10:
		goto sw_bb134
	case 11:
		goto sw_bb138
	case 12:
		goto sw_bb142
	case 13:
		goto sw_bb146
	case 14:
		goto sw_bb166
	case 15:
		goto sw_bb186
	case 16:
		goto sw_bb218
	case 17:
		goto sw_bb237
	case 18:
		goto sw_bb252
	case 19:
		goto sw_bb270
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
	*state_addr = 4
	goto next_state

if_end:
	v11 = *lookahead
	cmp = v11 == 39
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*state_addr = 10
	goto next_state

if_end6:
	v12 = *lookahead
	cmp7 = v12 == 40
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*state_addr = 6
	goto next_state

if_end10:
	v13 = *lookahead
	cmp11 = v13 == 41
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*state_addr = 7
	goto next_state

if_end14:
	v14 = *lookahead
	cmp15 = v14 == 42
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*state_addr = 8
	goto next_state

if_end18:
	v15 = *lookahead
	cmp19 = v15 == 47
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*state_addr = 1
	goto next_state

if_end22:
	v16 = *lookahead
	cmp23 = v16 == 58
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*state_addr = 11
	goto next_state

if_end26:
	v17 = *lookahead
	cmp27 = v17 == 61
	if cmp27 {
		goto if_then29
	} else {
		goto if_end30
	}

if_then29:
	*state_addr = 5
	goto next_state

if_end30:
	v18 = *lookahead
	cmp31 = v18 == 63
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*state_addr = 9
	goto next_state

if_end34:
	v19 = *lookahead
	cmp35 = v19 == 124
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*state_addr = 12
	goto next_state

if_end38:
	v20 = *lookahead
	cmp39 = v20 == 9
	if cmp39 {
		goto if_then49
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v21 = *lookahead
	cmp41 = v21 == 10
	if cmp41 {
		goto if_then49
	} else {
		goto lor_lhs_false43
	}

lor_lhs_false43:
	v22 = *lookahead
	cmp44 = v22 == 13
	if cmp44 {
		goto if_then49
	} else {
		goto lor_lhs_false46
	}

lor_lhs_false46:
	v23 = *lookahead
	cmp47 = v23 == 32
	if cmp47 {
		goto if_then49
	} else {
		goto if_end50
	}

if_then49:
	*skip = 1
	*state_addr = 0
	goto next_state

if_end50:
	v24 = *lookahead
	cmp51 = 65 <= v24
	if cmp51 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false55
	}

land_lhs_true:
	v25 = *lookahead
	cmp53 = v25 <= 90
	if cmp53 {
		goto if_then64
	} else {
		goto lor_lhs_false55
	}

lor_lhs_false55:
	v26 = *lookahead
	cmp56 = v26 == 95
	if cmp56 {
		goto if_then64
	} else {
		goto lor_lhs_false58
	}

lor_lhs_false58:
	v27 = *lookahead
	cmp59 = 97 <= v27
	if cmp59 {
		goto land_lhs_true61
	} else {
		goto if_end65
	}

land_lhs_true61:
	v28 = *lookahead
	cmp62 = v28 <= 122
	if cmp62 {
		goto if_then64
	} else {
		goto if_end65
	}

if_then64:
	*state_addr = 13
	goto next_state

if_end65:
	v29 = *result
	tobool66 = (v29 & 1) != 0
	*retval = tobool66
	goto _return

sw_bb67:
	v30 = *lookahead
	cmp68 = v30 == 47
	if cmp68 {
		goto if_then70
	} else {
		goto if_end71
	}

if_then70:
	*state_addr = 19
	goto next_state

if_end71:
	v31 = *result
	tobool72 = (v31 & 1) != 0
	*retval = tobool72
	goto _return

sw_bb73:
	v32 = *lookahead
	cmp74 = v32 == 47
	if cmp74 {
		goto if_then76
	} else {
		goto if_end77
	}

if_then76:
	*state_addr = 16
	goto next_state

if_end77:
	v33 = *lookahead
	cmp78 = v33 == 92
	if cmp78 {
		goto if_then80
	} else {
		goto if_end81
	}

if_then80:
	*state_addr = 3
	goto next_state

if_end81:
	v34 = *lookahead
	cmp82 = v34 == 9
	if cmp82 {
		goto if_then93
	} else {
		goto lor_lhs_false84
	}

lor_lhs_false84:
	v35 = *lookahead
	cmp85 = v35 == 10
	if cmp85 {
		goto if_then93
	} else {
		goto lor_lhs_false87
	}

lor_lhs_false87:
	v36 = *lookahead
	cmp88 = v36 == 13
	if cmp88 {
		goto if_then93
	} else {
		goto lor_lhs_false90
	}

lor_lhs_false90:
	v37 = *lookahead
	cmp91 = v37 == 32
	if cmp91 {
		goto if_then93
	} else {
		goto if_end94
	}

if_then93:
	*state_addr = 15
	goto next_state

if_end94:
	v38 = *lookahead
	cmp95 = v38 != 0
	if cmp95 {
		goto land_lhs_true97
	} else {
		goto if_end101
	}

land_lhs_true97:
	v39 = *lookahead
	cmp98 = v39 != 39
	if cmp98 {
		goto if_then100
	} else {
		goto if_end101
	}

if_then100:
	*state_addr = 17
	goto next_state

if_end101:
	v40 = *result
	tobool102 = (v40 & 1) != 0
	*retval = tobool102
	goto _return

sw_bb103:
	v41 = *lookahead
	cmp104 = v41 == 39
	if cmp104 {
		goto if_then109
	} else {
		goto lor_lhs_false106
	}

lor_lhs_false106:
	v42 = *lookahead
	cmp107 = v42 == 92
	if cmp107 {
		goto if_then109
	} else {
		goto if_end110
	}

if_then109:
	*state_addr = 17
	goto next_state

if_end110:
	v43 = *result
	tobool111 = (v43 & 1) != 0
	*retval = tobool111
	goto _return

sw_bb112:
	*result = 1
	v44 = *lexer_addr
	result_symbol = &v44.F1
	*result_symbol = 0
	v45 = *lexer_addr
	mark_end = &v45.F3
	v46 = *mark_end
	v47 = *lexer_addr
	v46(v47)
	v48 = *result
	tobool113 = (v48 & 1) != 0
	*retval = tobool113
	goto _return

sw_bb114:
	*result = 1
	v49 = *lexer_addr
	result_symbol115 = &v49.F1
	*result_symbol115 = 2
	v50 = *lexer_addr
	mark_end116 = &v50.F3
	v51 = *mark_end116
	v52 = *lexer_addr
	v51(v52)
	v53 = *result
	tobool117 = (v53 & 1) != 0
	*retval = tobool117
	goto _return

sw_bb118:
	*result = 1
	v54 = *lexer_addr
	result_symbol119 = &v54.F1
	*result_symbol119 = 3
	v55 = *lexer_addr
	mark_end120 = &v55.F3
	v56 = *mark_end120
	v57 = *lexer_addr
	v56(v57)
	v58 = *result
	tobool121 = (v58 & 1) != 0
	*retval = tobool121
	goto _return

sw_bb122:
	*result = 1
	v59 = *lexer_addr
	result_symbol123 = &v59.F1
	*result_symbol123 = 4
	v60 = *lexer_addr
	mark_end124 = &v60.F3
	v61 = *mark_end124
	v62 = *lexer_addr
	v61(v62)
	v63 = *result
	tobool125 = (v63 & 1) != 0
	*retval = tobool125
	goto _return

sw_bb126:
	*result = 1
	v64 = *lexer_addr
	result_symbol127 = &v64.F1
	*result_symbol127 = 5
	v65 = *lexer_addr
	mark_end128 = &v65.F3
	v66 = *mark_end128
	v67 = *lexer_addr
	v66(v67)
	v68 = *result
	tobool129 = (v68 & 1) != 0
	*retval = tobool129
	goto _return

sw_bb130:
	*result = 1
	v69 = *lexer_addr
	result_symbol131 = &v69.F1
	*result_symbol131 = 6
	v70 = *lexer_addr
	mark_end132 = &v70.F3
	v71 = *mark_end132
	v72 = *lexer_addr
	v71(v72)
	v73 = *result
	tobool133 = (v73 & 1) != 0
	*retval = tobool133
	goto _return

sw_bb134:
	*result = 1
	v74 = *lexer_addr
	result_symbol135 = &v74.F1
	*result_symbol135 = 7
	v75 = *lexer_addr
	mark_end136 = &v75.F3
	v76 = *mark_end136
	v77 = *lexer_addr
	v76(v77)
	v78 = *result
	tobool137 = (v78 & 1) != 0
	*retval = tobool137
	goto _return

sw_bb138:
	*result = 1
	v79 = *lexer_addr
	result_symbol139 = &v79.F1
	*result_symbol139 = 8
	v80 = *lexer_addr
	mark_end140 = &v80.F3
	v81 = *mark_end140
	v82 = *lexer_addr
	v81(v82)
	v83 = *result
	tobool141 = (v83 & 1) != 0
	*retval = tobool141
	goto _return

sw_bb142:
	*result = 1
	v84 = *lexer_addr
	result_symbol143 = &v84.F1
	*result_symbol143 = 9
	v85 = *lexer_addr
	mark_end144 = &v85.F3
	v86 = *mark_end144
	v87 = *lexer_addr
	v86(v87)
	v88 = *result
	tobool145 = (v88 & 1) != 0
	*retval = tobool145
	goto _return

sw_bb146:
	*result = 1
	v89 = *lexer_addr
	result_symbol147 = &v89.F1
	*result_symbol147 = 1
	v90 = *lexer_addr
	mark_end148 = &v90.F3
	v91 = *mark_end148
	v92 = *lexer_addr
	v91(v92)
	v93 = *lookahead
	cmp149 = 65 <= v93
	if cmp149 {
		goto land_lhs_true151
	} else {
		goto lor_lhs_false154
	}

land_lhs_true151:
	v94 = *lookahead
	cmp152 = v94 <= 90
	if cmp152 {
		goto if_then163
	} else {
		goto lor_lhs_false154
	}

lor_lhs_false154:
	v95 = *lookahead
	cmp155 = v95 == 95
	if cmp155 {
		goto if_then163
	} else {
		goto lor_lhs_false157
	}

lor_lhs_false157:
	v96 = *lookahead
	cmp158 = 97 <= v96
	if cmp158 {
		goto land_lhs_true160
	} else {
		goto if_end164
	}

land_lhs_true160:
	v97 = *lookahead
	cmp161 = v97 <= 122
	if cmp161 {
		goto if_then163
	} else {
		goto if_end164
	}

if_then163:
	*state_addr = 13
	goto next_state

if_end164:
	v98 = *result
	tobool165 = (v98 & 1) != 0
	*retval = tobool165
	goto _return

sw_bb166:
	*result = 1
	v99 = *lexer_addr
	result_symbol167 = &v99.F1
	*result_symbol167 = 10
	v100 = *lexer_addr
	mark_end168 = &v100.F3
	v101 = *mark_end168
	v102 = *lexer_addr
	v101(v102)
	v103 = *lookahead
	cmp169 = v103 == 10
	if cmp169 {
		goto if_then171
	} else {
		goto if_end172
	}

if_then171:
	*state_addr = 17
	goto next_state

if_end172:
	v104 = *lookahead
	cmp173 = v104 == 39
	if cmp173 {
		goto if_then175
	} else {
		goto if_end176
	}

if_then175:
	*state_addr = 19
	goto next_state

if_end176:
	v105 = *lookahead
	cmp177 = v105 == 92
	if cmp177 {
		goto if_then179
	} else {
		goto if_end180
	}

if_then179:
	*state_addr = 18
	goto next_state

if_end180:
	v106 = *lookahead
	cmp181 = v106 != 0
	if cmp181 {
		goto if_then183
	} else {
		goto if_end184
	}

if_then183:
	*state_addr = 14
	goto next_state

if_end184:
	v107 = *result
	tobool185 = (v107 & 1) != 0
	*retval = tobool185
	goto _return

sw_bb186:
	*result = 1
	v108 = *lexer_addr
	result_symbol187 = &v108.F1
	*result_symbol187 = 10
	v109 = *lexer_addr
	mark_end188 = &v109.F3
	v110 = *mark_end188
	v111 = *lexer_addr
	v110(v111)
	v112 = *lookahead
	cmp189 = v112 == 47
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*state_addr = 16
	goto next_state

if_end192:
	v113 = *lookahead
	cmp193 = v113 == 92
	if cmp193 {
		goto if_then195
	} else {
		goto if_end196
	}

if_then195:
	*state_addr = 3
	goto next_state

if_end196:
	v114 = *lookahead
	cmp197 = v114 == 9
	if cmp197 {
		goto if_then208
	} else {
		goto lor_lhs_false199
	}

lor_lhs_false199:
	v115 = *lookahead
	cmp200 = v115 == 10
	if cmp200 {
		goto if_then208
	} else {
		goto lor_lhs_false202
	}

lor_lhs_false202:
	v116 = *lookahead
	cmp203 = v116 == 13
	if cmp203 {
		goto if_then208
	} else {
		goto lor_lhs_false205
	}

lor_lhs_false205:
	v117 = *lookahead
	cmp206 = v117 == 32
	if cmp206 {
		goto if_then208
	} else {
		goto if_end209
	}

if_then208:
	*state_addr = 15
	goto next_state

if_end209:
	v118 = *lookahead
	cmp210 = v118 != 0
	if cmp210 {
		goto land_lhs_true212
	} else {
		goto if_end216
	}

land_lhs_true212:
	v119 = *lookahead
	cmp213 = v119 != 39
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*state_addr = 17
	goto next_state

if_end216:
	v120 = *result
	tobool217 = (v120 & 1) != 0
	*retval = tobool217
	goto _return

sw_bb218:
	*result = 1
	v121 = *lexer_addr
	result_symbol219 = &v121.F1
	*result_symbol219 = 10
	v122 = *lexer_addr
	mark_end220 = &v122.F3
	v123 = *mark_end220
	v124 = *lexer_addr
	v123(v124)
	v125 = *lookahead
	cmp221 = v125 == 47
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*state_addr = 14
	goto next_state

if_end224:
	v126 = *lookahead
	cmp225 = v126 == 92
	if cmp225 {
		goto if_then227
	} else {
		goto if_end228
	}

if_then227:
	*state_addr = 3
	goto next_state

if_end228:
	v127 = *lookahead
	cmp229 = v127 != 0
	if cmp229 {
		goto land_lhs_true231
	} else {
		goto if_end235
	}

land_lhs_true231:
	v128 = *lookahead
	cmp232 = v128 != 39
	if cmp232 {
		goto if_then234
	} else {
		goto if_end235
	}

if_then234:
	*state_addr = 17
	goto next_state

if_end235:
	v129 = *result
	tobool236 = (v129 & 1) != 0
	*retval = tobool236
	goto _return

sw_bb237:
	*result = 1
	v130 = *lexer_addr
	result_symbol238 = &v130.F1
	*result_symbol238 = 10
	v131 = *lexer_addr
	mark_end239 = &v131.F3
	v132 = *mark_end239
	v133 = *lexer_addr
	v132(v133)
	v134 = *lookahead
	cmp240 = v134 == 92
	if cmp240 {
		goto if_then242
	} else {
		goto if_end243
	}

if_then242:
	*state_addr = 3
	goto next_state

if_end243:
	v135 = *lookahead
	cmp244 = v135 != 0
	if cmp244 {
		goto land_lhs_true246
	} else {
		goto if_end250
	}

land_lhs_true246:
	v136 = *lookahead
	cmp247 = v136 != 39
	if cmp247 {
		goto if_then249
	} else {
		goto if_end250
	}

if_then249:
	*state_addr = 17
	goto next_state

if_end250:
	v137 = *result
	tobool251 = (v137 & 1) != 0
	*retval = tobool251
	goto _return

sw_bb252:
	*result = 1
	v138 = *lexer_addr
	result_symbol253 = &v138.F1
	*result_symbol253 = 11
	v139 = *lexer_addr
	mark_end254 = &v139.F3
	v140 = *mark_end254
	v141 = *lexer_addr
	v140(v141)
	v142 = *lookahead
	cmp255 = v142 == 39
	if cmp255 {
		goto if_then260
	} else {
		goto lor_lhs_false257
	}

lor_lhs_false257:
	v143 = *lookahead
	cmp258 = v143 == 92
	if cmp258 {
		goto if_then260
	} else {
		goto if_end261
	}

if_then260:
	*state_addr = 14
	goto next_state

if_end261:
	v144 = *lookahead
	cmp262 = v144 != 0
	if cmp262 {
		goto land_lhs_true264
	} else {
		goto if_end268
	}

land_lhs_true264:
	v145 = *lookahead
	cmp265 = v145 != 10
	if cmp265 {
		goto if_then267
	} else {
		goto if_end268
	}

if_then267:
	*state_addr = 19
	goto next_state

if_end268:
	v146 = *result
	tobool269 = (v146 & 1) != 0
	*retval = tobool269
	goto _return

sw_bb270:
	*result = 1
	v147 = *lexer_addr
	result_symbol271 = &v147.F1
	*result_symbol271 = 11
	v148 = *lexer_addr
	mark_end272 = &v148.F3
	v149 = *mark_end272
	v150 = *lexer_addr
	v149(v150)
	v151 = *lookahead
	cmp273 = v151 != 0
	if cmp273 {
		goto land_lhs_true275
	} else {
		goto if_end279
	}

land_lhs_true275:
	v152 = *lookahead
	cmp276 = v152 != 10
	if cmp276 {
		goto if_then278
	} else {
		goto if_end279
	}

if_then278:
	*state_addr = 19
	goto next_state

if_end279:
	v153 = *result
	tobool280 = (v153 & 1) != 0
	*retval = tobool280
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v154 = *retval
	return v154
}

func ts_lex_keywords(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v10, v11, v13 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, result_symbol *int16
	var lookahead, lookahead1 *int32
	var tobool, call, tobool3, v15 bool
	var v3, frombool, v14 byte
	var v12 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9 int16
	var v5, conv int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, result_symbol, v11, mark_end, v12, v13, v14, tobool3, v15

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
	*lexer_addr = lexer
	*state_addr = state
	*result = 0
	*skip = 0
	*eof = 0
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
	default:
		goto sw_default
	}

sw_bb:
	*result = 1
	v10 = *lexer_addr
	result_symbol = &v10.F1
	*result_symbol = 0
	v11 = *lexer_addr
	mark_end = &v11.F3
	v12 = *mark_end
	v13 = *lexer_addr
	v12(v13)
	v14 = *result
	tobool3 = (v14 & 1) != 0
	*retval = tobool3
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v15 = *retval
	return v15
}

