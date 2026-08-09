package grammar_cyberchef

import "unsafe"

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

var tree_sitter_cyberchef_language TSLanguage = TSLanguage{15, 25, 0, 14, 0, 48, 2, 13, 4, 7, &(*[2][25]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[139]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], &ts_field_names[0], &ts_field_map_slices[0], &ts_field_map_entries[0], &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{0, 1, 0}}

var ts_small_parse_table [367]int16 = [367]int16{
	9, 9, 1, 3, 13, 1, 8, 15, 1, 12, 17, 1, 13, 8, 1, 17,
	22, 1, 20, 44, 1, 21, 6, 2, 16, 18, 11, 3, 5, 6, 7, 3,
	13, 1, 8, 13, 2, 16, 18, 19, 3, 5, 6, 7, 3, 13, 1, 8,
	29, 2, 16, 18, 21, 3, 5, 6, 7, 4, 23, 1, 0, 25, 1, 1,
	28, 1, 4, 5, 2, 15, 22, 3, 33, 1, 9, 9, 1, 24, 31, 3,
	3, 12, 13, 4, 7, 1, 4, 35, 1, 0, 37, 1, 1, 5, 2, 15,
	22, 5, 15, 1, 12, 17, 1, 13, 39, 1, 3, 17, 1, 20, 35, 1,
	21, 3, 33, 1, 9, 12, 1, 24, 41, 3, 3, 12, 13, 1, 43, 5,
	3, 9, 10, 12, 13, 1, 45, 5, 3, 9, 10, 12, 13, 3, 49, 1,
	9, 12, 1, 24, 47, 3, 3, 12, 13, 1, 47, 4, 3, 9, 12, 13,
	3, 52, 1, 9, 55, 1, 10, 14, 1, 23, 2, 59, 1, 4, 57, 2,
	0, 1, 3, 61, 1, 9, 63, 1, 10, 18, 1, 23, 3, 17, 1, 13,
	65, 1, 3, 37, 1, 21, 3, 61, 1, 9, 67, 1, 10, 14, 1, 23,
	2, 71, 1, 4, 69, 2, 0, 1, 2, 75, 1, 4, 73, 2, 0, 1,
	2, 79, 1, 4, 77, 2, 0, 1, 3, 17, 1, 13, 81, 1, 3, 45,
	1, 21, 2, 85, 1, 4, 83, 2, 0, 1, 2, 89, 1, 4, 87, 2,
	0, 1, 2, 93, 1, 4, 91, 2, 0, 1, 2, 97, 1, 4, 95, 2,
	0, 1, 1, 99, 2, 9, 10, 2, 101, 1, 5, 16, 1, 19, 1, 103,
	2, 9, 10, 1, 105, 2, 3, 13, 2, 101, 1, 5, 27, 1, 19, 1,
	107, 1, 11, 1, 109, 1, 3, 1, 111, 1, 1, 1, 113, 1, 3, 1,
	115, 1, 1, 1, 117, 1, 3, 1, 119, 1, 1, 1, 121, 1, 0, 1,
	123, 1, 2, 1, 125, 1, 1, 1, 127, 1, 1, 1, 129, 1, 1, 1,
	131, 1, 3, 1, 133, 1, 3, 1, 135, 1, 1, 1, 137, 1, 1,
}

var ts_small_parse_table_map [46]int32 = [46]int32{
	0, 31, 44, 57, 71, 83, 97, 113, 125, 133, 141, 153, 160, 170, 178, 188,
	198, 208, 216, 224, 232, 242, 250, 258, 266, 274, 279, 286, 291, 296, 303, 307,
	311, 315, 319, 323, 327, 331, 335, 339, 343, 347, 351, 355, 359, 363,
}

var ts_symbol_names [25]*byte = [25]*byte{
	&_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0], &_str_17[0], &_str_18[0],
	&_str_19[0], &_str_20[0], &_str_21[0], &_str_22[0], &_str_23[0], &_str_24[0], &_str_25[0], &_str_26[0], &_str_27[0],
}

var ts_field_names [5]*byte = [5]*byte{nil, &_str_28[0], &_str_29[0], &_str_30[0], &_str_31[0]}

var ts_field_map_slices [13]TSMapSlice = [13]TSMapSlice{TSMapSlice{}, TSMapSlice{0, 1}, TSMapSlice{1, 1}, TSMapSlice{2, 2}, TSMapSlice{4, 1}, TSMapSlice{5, 1}, TSMapSlice{6, 2}, TSMapSlice{8, 4}, TSMapSlice{12, 4}, TSMapSlice{16, 1}, TSMapSlice{17, 1}, TSMapSlice{18, 2}, TSMapSlice{20, 2}}

var ts_field_map_entries [22]TSFieldMapEntry = [22]TSFieldMapEntry{
	TSFieldMapEntry{2, 0, 0}, TSFieldMapEntry{1, 0, 0}, TSFieldMapEntry{3, 1, 1}, TSFieldMapEntry{4, 1, 1}, TSFieldMapEntry{2, 2, 1}, TSFieldMapEntry{1, 2, 1}, TSFieldMapEntry{3, 0, 0}, TSFieldMapEntry{4, 2, 0}, TSFieldMapEntry{3, 1, 1}, TSFieldMapEntry{3, 2, 1}, TSFieldMapEntry{4, 1, 1}, TSFieldMapEntry{4, 2, 1}, TSFieldMapEntry{3, 0, 1}, TSFieldMapEntry{3, 1, 1}, TSFieldMapEntry{4, 0, 1}, TSFieldMapEntry{4, 1, 1},
	TSFieldMapEntry{2, 3, 1}, TSFieldMapEntry{1, 3, 1}, TSFieldMapEntry{1, 3, 1}, TSFieldMapEntry{2, 2, 1}, TSFieldMapEntry{1, 4, 1}, TSFieldMapEntry{2, 3, 1},
}

var ts_symbol_metadata [25]TSSymbolMetadata = [25]TSSymbolMetadata{
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0},
	TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{},
}

var ts_symbol_map [25]int16 = [25]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24,
}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [13][7]int16 = [13][7]int16{}

var ts_lex_modes [48]TSLexerMode = [48]TSLexerMode{
	TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{29, 0, 0},
	TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{},
	TSLexerMode{}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{29, 0, 0}, TSLexerMode{29, 0, 0},
}

var ts_primary_state_ids [48]int16 = [48]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
}

var _str [10]byte = [10]byte{99, 121, 98, 101, 114, 99, 104, 101, 102, 0}

var ts_parse_table struct {
	F0 struct {
	F0 [14]int16
	F1 [11]int16
}
	F1 [25]int16
} = struct {
	F0 struct {
	F0 [14]int16
	F1 [11]int16
}
	F1 [25]int16
}{struct {
	F0 [14]int16
	F1 [11]int16
}{[14]int16{1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, [11]int16{}}, [25]int16{
	3, 5, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 39, 7,
	0, 0, 0, 0, 0, 0, 7, 0, 0,
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
	F52 struct {
	F0 anon_1
	F1 [6]byte
}
	F53 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
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
	F64 struct {
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
	F65 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F72 TSParseActionEntry
	F73 struct {
	F0 anon_1
	F1 [6]byte
}
	F74 TSParseActionEntry
	F75 struct {
	F0 anon_1
	F1 [6]byte
}
	F76 TSParseActionEntry
	F77 struct {
	F0 anon_1
	F1 [6]byte
}
	F78 TSParseActionEntry
	F79 struct {
	F0 anon_1
	F1 [6]byte
}
	F80 TSParseActionEntry
	F81 struct {
	F0 anon_1
	F1 [6]byte
}
	F82 struct {
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
	F83 struct {
	F0 anon_1
	F1 [6]byte
}
	F84 TSParseActionEntry
	F85 struct {
	F0 anon_1
	F1 [6]byte
}
	F86 TSParseActionEntry
	F87 struct {
	F0 anon_1
	F1 [6]byte
}
	F88 TSParseActionEntry
	F89 struct {
	F0 anon_1
	F1 [6]byte
}
	F90 TSParseActionEntry
	F91 struct {
	F0 anon_1
	F1 [6]byte
}
	F92 TSParseActionEntry
	F93 struct {
	F0 anon_1
	F1 [6]byte
}
	F94 TSParseActionEntry
	F95 struct {
	F0 anon_1
	F1 [6]byte
}
	F96 TSParseActionEntry
	F97 struct {
	F0 anon_1
	F1 [6]byte
}
	F98 TSParseActionEntry
	F99 struct {
	F0 anon_1
	F1 [6]byte
}
	F100 TSParseActionEntry
	F101 struct {
	F0 anon_1
	F1 [6]byte
}
	F102 struct {
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
	F103 struct {
	F0 anon_1
	F1 [6]byte
}
	F104 TSParseActionEntry
	F105 struct {
	F0 anon_1
	F1 [6]byte
}
	F106 TSParseActionEntry
	F107 struct {
	F0 anon_1
	F1 [6]byte
}
	F108 struct {
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
	F109 struct {
	F0 anon_1
	F1 [6]byte
}
	F110 TSParseActionEntry
	F111 struct {
	F0 anon_1
	F1 [6]byte
}
	F112 struct {
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
	F113 struct {
	F0 anon_1
	F1 [6]byte
}
	F114 struct {
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
	F115 struct {
	F0 anon_1
	F1 [6]byte
}
	F116 struct {
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
	F117 struct {
	F0 anon_1
	F1 [6]byte
}
	F118 struct {
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
	F119 struct {
	F0 anon_1
	F1 [6]byte
}
	F120 struct {
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
	F121 struct {
	F0 anon_1
	F1 [6]byte
}
	F122 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F123 struct {
	F0 anon_1
	F1 [6]byte
}
	F124 struct {
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
	F125 struct {
	F0 anon_1
	F1 [6]byte
}
	F126 struct {
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
	F127 struct {
	F0 anon_1
	F1 [6]byte
}
	F128 struct {
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
	F129 struct {
	F0 anon_1
	F1 [6]byte
}
	F130 struct {
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
	F131 struct {
	F0 anon_1
	F1 [6]byte
}
	F132 struct {
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
	F133 struct {
	F0 anon_1
	F1 [6]byte
}
	F134 struct {
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
	F135 struct {
	F0 anon_1
	F1 [6]byte
}
	F136 struct {
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
	F137 struct {
	F0 anon_1
	F1 [6]byte
}
	F138 struct {
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
	F52 struct {
	F0 anon_1
	F1 [6]byte
}
	F53 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
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
	F64 struct {
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
	F65 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F72 TSParseActionEntry
	F73 struct {
	F0 anon_1
	F1 [6]byte
}
	F74 TSParseActionEntry
	F75 struct {
	F0 anon_1
	F1 [6]byte
}
	F76 TSParseActionEntry
	F77 struct {
	F0 anon_1
	F1 [6]byte
}
	F78 TSParseActionEntry
	F79 struct {
	F0 anon_1
	F1 [6]byte
}
	F80 TSParseActionEntry
	F81 struct {
	F0 anon_1
	F1 [6]byte
}
	F82 struct {
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
	F83 struct {
	F0 anon_1
	F1 [6]byte
}
	F84 TSParseActionEntry
	F85 struct {
	F0 anon_1
	F1 [6]byte
}
	F86 TSParseActionEntry
	F87 struct {
	F0 anon_1
	F1 [6]byte
}
	F88 TSParseActionEntry
	F89 struct {
	F0 anon_1
	F1 [6]byte
}
	F90 TSParseActionEntry
	F91 struct {
	F0 anon_1
	F1 [6]byte
}
	F92 TSParseActionEntry
	F93 struct {
	F0 anon_1
	F1 [6]byte
}
	F94 TSParseActionEntry
	F95 struct {
	F0 anon_1
	F1 [6]byte
}
	F96 TSParseActionEntry
	F97 struct {
	F0 anon_1
	F1 [6]byte
}
	F98 TSParseActionEntry
	F99 struct {
	F0 anon_1
	F1 [6]byte
}
	F100 TSParseActionEntry
	F101 struct {
	F0 anon_1
	F1 [6]byte
}
	F102 struct {
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
	F103 struct {
	F0 anon_1
	F1 [6]byte
}
	F104 TSParseActionEntry
	F105 struct {
	F0 anon_1
	F1 [6]byte
}
	F106 TSParseActionEntry
	F107 struct {
	F0 anon_1
	F1 [6]byte
}
	F108 struct {
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
	F109 struct {
	F0 anon_1
	F1 [6]byte
}
	F110 TSParseActionEntry
	F111 struct {
	F0 anon_1
	F1 [6]byte
}
	F112 struct {
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
	F113 struct {
	F0 anon_1
	F1 [6]byte
}
	F114 struct {
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
	F115 struct {
	F0 anon_1
	F1 [6]byte
}
	F116 struct {
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
	F117 struct {
	F0 anon_1
	F1 [6]byte
}
	F118 struct {
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
	F119 struct {
	F0 anon_1
	F1 [6]byte
}
	F120 struct {
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
	F121 struct {
	F0 anon_1
	F1 [6]byte
}
	F122 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F123 struct {
	F0 anon_1
	F1 [6]byte
}
	F124 struct {
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
	F125 struct {
	F0 anon_1
	F1 [6]byte
}
	F126 struct {
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
	F127 struct {
	F0 anon_1
	F1 [6]byte
}
	F128 struct {
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
	F129 struct {
	F0 anon_1
	F1 [6]byte
}
	F130 struct {
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
	F131 struct {
	F0 anon_1
	F1 [6]byte
}
	F132 struct {
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
	F133 struct {
	F0 anon_1
	F1 [6]byte
}
	F134 struct {
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
	F135 struct {
	F0 anon_1
	F1 [6]byte
}
	F136 struct {
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
	F137 struct {
	F0 anon_1
	F1 [6]byte
}
	F138 struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 14, 0, 0}}}, struct {
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
}{0, 40, 0, 0}, [2]byte{}}}, struct {
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
}{0, 41, 0, 0}, [2]byte{}}}, struct {
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
}{0, 33, 0, 0}, [2]byte{}}}, struct {
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
}{0, 29, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 22, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 22, 0, 0}}}, struct {
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
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 22, 0, 0}}}, struct {
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
}{0, 40, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 17, 0, 0}}}, struct {
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
}{0, 34, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 16, 0, 3}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 16, 0, 7}}}, struct {
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
}{0, 3, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 23, 0, 8}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 23, 0, 8}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 7, 15, 0, 12}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 7, 15, 0, 12}}}, struct {
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
}{0, 36, 0, 0}, [2]byte{}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 15, 0, 4}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 15, 0, 4}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 15, 0, 5}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 15, 0, 5}}}, struct {
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
}{0, 43, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 15, 0, 9}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 15, 0, 9}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 15, 0, 10}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 15, 0, 10}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 15, 0, 11}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 6, 15, 0, 11}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 23, 0, 3}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 19, 0, 6}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 20, 0, 1}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 21, 0, 2}}}, struct {
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
}{0, 38, 0, 0}, [2]byte{}}}, struct {
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
}{0, 46, 0, 0}, [2]byte{}}}, struct {
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
}{0, 47, 0, 0}, [2]byte{}}}, struct {
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
}{0, 42, 0, 0}, [2]byte{}}}, struct {
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
}{0, 21, 0, 0}, [2]byte{}}}}

var _str_3 [4]byte = [4]byte{101, 110, 100, 0}

var _str_4 [14]byte = [14]byte{114, 101, 99, 105, 112, 101, 95, 116, 111, 107, 101, 110, 49, 0}

var _str_5 [2]byte = [2]byte{40, 0}

var _str_6 [2]byte = [2]byte{41, 0}

var _str_7 [5]byte = [5]byte{110, 97, 109, 101, 0}

var _str_8 [7]byte = [7]byte{115, 116, 114, 105, 110, 103, 0}

var _str_9 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}

var _str_10 [8]byte = [8]byte{98, 111, 111, 108, 101, 97, 110, 0}

var _str_11 [2]byte = [2]byte{123, 0}

var _str_12 [2]byte = [2]byte{44, 0}

var _str_13 [2]byte = [2]byte{125, 0}

var _str_14 [2]byte = [2]byte{58, 0}

var _str_15 [10]byte = [10]byte{47, 100, 105, 115, 97, 98, 108, 101, 100, 0}

var _str_16 [12]byte = [12]byte{47, 98, 114, 101, 97, 107, 112, 111, 105, 110, 116, 0}

var _str_17 [7]byte = [7]byte{114, 101, 99, 105, 112, 101, 0}

var _str_18 [10]byte = [10]byte{111, 112, 101, 114, 97, 116, 105, 111, 110, 0}

var _str_19 [7]byte = [7]byte{111, 98, 106, 101, 99, 116, 0}

var _str_20 [10]byte = [10]byte{97, 114, 103, 117, 109, 101, 110, 116, 115, 0}

var _str_21 [10]byte = [10]byte{95, 97, 114, 103, 117, 109, 101, 110, 116, 0}

var _str_22 [6]byte = [6]byte{95, 112, 97, 105, 114, 0}

var _str_23 [10]byte = [10]byte{95, 100, 105, 115, 97, 98, 108, 101, 100, 0}

var _str_24 [12]byte = [12]byte{95, 98, 114, 101, 97, 107, 112, 111, 105, 110, 116, 0}

var _str_25 [15]byte = [15]byte{114, 101, 99, 105, 112, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}

var _str_26 [15]byte = [15]byte{111, 98, 106, 101, 99, 116, 95, 114, 101, 112, 101, 97, 116, 49, 0}

var _str_27 [18]byte = [18]byte{
	97, 114, 103, 117, 109, 101, 110, 116, 115, 95, 114, 101, 112, 101, 97, 116,
	49, 0,
}

var _str_28 [11]byte = [11]byte{98, 114, 101, 97, 107, 112, 111, 105, 110, 116, 0}

var _str_29 [9]byte = [9]byte{100, 105, 115, 97, 98, 108, 101, 100, 0}

var _str_30 [4]byte = [4]byte{107, 101, 121, 0}

var _str_31 [6]byte = [6]byte{118, 97, 108, 117, 101, 0}

var ts_lex_map [22]int16 = [22]int16{
	39, 2, 40, 32, 41, 33, 44, 41, 45, 27, 47, 7, 58, 43, 102, 4,
	116, 21, 123, 40, 125, 42,
}

func tree_sitter_cyberchef() *TSLanguage {
	return &tree_sitter_cyberchef_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v112, v113, v115, v117, v118, v120, v124, v125, v127, v129, v130, v132, v134, v135, v137, v146, v147, v149, v151, v152, v154, v159, v160, v162, v167, v168, v170, v174, v175, v177, v179, v180, v182, v184, v185, v187, v189, v190, v192, v194, v195, v197, v199, v200, v202, v204, v205, v207 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end309, mark_end321, mark_end325, mark_end329, mark_end355, mark_end359, mark_end375, mark_end390, mark_end401, mark_end405, mark_end409, mark_end413, mark_end417, mark_end421, mark_end425 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol308, result_symbol320, result_symbol324, result_symbol328, result_symbol354, result_symbol358, result_symbol374, result_symbol389, result_symbol400, result_symbol404, result_symbol408, result_symbol412, result_symbol416, result_symbol420, result_symbol424 *int16
	var lookahead, i, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, cmp25, tobool29, cmp31, cmp35, cmp39, cmp42, cmp45, cmp49, cmp52, cmp55, cmp58, cmp61, cmp64, cmp67, tobool71, cmp73, cmp77, cmp81, tobool85, cmp87, cmp91, cmp95, tobool99, cmp101, tobool105, cmp107, tobool111, cmp113, tobool117, cmp119, cmp123, tobool127, cmp129, tobool133, cmp135, tobool139, cmp141, tobool145, cmp147, tobool151, cmp153, tobool157, cmp159, tobool163, cmp165, tobool169, cmp171, tobool175, cmp177, tobool181, cmp183, tobool187, cmp189, tobool193, cmp195, tobool199, cmp201, tobool205, cmp207, tobool211, cmp213, tobool217, cmp219, tobool223, cmp225, tobool229, cmp231, tobool235, cmp237, tobool241, cmp243, cmp246, tobool250, cmp252, cmp255, tobool259, tobool261, cmp264, cmp268, cmp272, cmp275, cmp278, cmp282, cmp285, cmp288, cmp291, cmp294, cmp297, cmp300, tobool304, tobool306, cmp310, cmp314, tobool318, tobool322, tobool326, cmp330, cmp333, cmp336, cmp339, cmp342, cmp345, cmp348, tobool352, tobool356, cmp360, cmp364, cmp368, tobool372, cmp376, cmp380, cmp383, tobool387, cmp391, cmp394, tobool398, tobool402, tobool406, tobool410, tobool414, tobool418, tobool422, tobool426, v209 bool
	var v3, frombool, v10, v23, v36, v40, v44, v46, v48, v50, v53, v55, v57, v59, v61, v63, v65, v67, v69, v71, v73, v75, v77, v79, v81, v83, v85, v87, v89, v91, v94, v97, v98, v111, v116, v123, v128, v133, v145, v150, v158, v166, v173, v178, v183, v188, v193, v198, v203, v208 byte
	var v114, v119, v126, v131, v136, v148, v153, v161, v169, v176, v181, v186, v191, v196, v201, v206 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9, v13, v16 int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v22, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v37, v38, v39, v41, v42, v43, v45, v47, v49, v51, v52, v54, v56, v58, v60, v62, v64, v66, v68, v70, v72, v74, v76, v78, v80, v82, v84, v86, v88, v90, v92, v93, v95, v96, v99, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v110, v121, v122, v138, v139, v140, v141, v142, v143, v144, v155, v156, v157, v163, v164, v165, v171, v172 int32
	var conv4, idxprom, idxprom10 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, cmp25, v23, tobool29, v24, cmp31, v25, cmp35, v26, cmp39, v27, cmp42, v28, cmp45, v29, cmp49, v30, cmp52, v31, cmp55, v32, cmp58, v33, cmp61, v34, cmp64, v35, cmp67, v36, tobool71, v37, cmp73, v38, cmp77, v39, cmp81, v40, tobool85, v41, cmp87, v42, cmp91, v43, cmp95, v44, tobool99, v45, cmp101, v46, tobool105, v47, cmp107, v48, tobool111, v49, cmp113, v50, tobool117, v51, cmp119, v52, cmp123, v53, tobool127, v54, cmp129, v55, tobool133, v56, cmp135, v57, tobool139, v58, cmp141, v59, tobool145, v60, cmp147, v61, tobool151, v62, cmp153, v63, tobool157, v64, cmp159, v65, tobool163, v66, cmp165, v67, tobool169, v68, cmp171, v69, tobool175, v70, cmp177, v71, tobool181, v72, cmp183, v73, tobool187, v74, cmp189, v75, tobool193, v76, cmp195, v77, tobool199, v78, cmp201, v79, tobool205, v80, cmp207, v81, tobool211, v82, cmp213, v83, tobool217, v84, cmp219, v85, tobool223, v86, cmp225, v87, tobool229, v88, cmp231, v89, tobool235, v90, cmp237, v91, tobool241, v92, cmp243, v93, cmp246, v94, tobool250, v95, cmp252, v96, cmp255, v97, tobool259, v98, tobool261, v99, cmp264, v100, cmp268, v101, cmp272, v102, cmp275, v103, cmp278, v104, cmp282, v105, cmp285, v106, cmp288, v107, cmp291, v108, cmp294, v109, cmp297, v110, cmp300, v111, tobool304, v112, result_symbol, v113, mark_end, v114, v115, v116, tobool306, v117, result_symbol308, v118, mark_end309, v119, v120, v121, cmp310, v122, cmp314, v123, tobool318, v124, result_symbol320, v125, mark_end321, v126, v127, v128, tobool322, v129, result_symbol324, v130, mark_end325, v131, v132, v133, tobool326, v134, result_symbol328, v135, mark_end329, v136, v137, v138, cmp330, v139, cmp333, v140, cmp336, v141, cmp339, v142, cmp342, v143, cmp345, v144, cmp348, v145, tobool352, v146, result_symbol354, v147, mark_end355, v148, v149, v150, tobool356, v151, result_symbol358, v152, mark_end359, v153, v154, v155, cmp360, v156, cmp364, v157, cmp368, v158, tobool372, v159, result_symbol374, v160, mark_end375, v161, v162, v163, cmp376, v164, cmp380, v165, cmp383, v166, tobool387, v167, result_symbol389, v168, mark_end390, v169, v170, v171, cmp391, v172, cmp394, v173, tobool398, v174, result_symbol400, v175, mark_end401, v176, v177, v178, tobool402, v179, result_symbol404, v180, mark_end405, v181, v182, v183, tobool406, v184, result_symbol408, v185, mark_end409, v186, v187, v188, tobool410, v189, result_symbol412, v190, mark_end413, v191, v192, v193, tobool414, v194, result_symbol416, v195, mark_end417, v196, v197, v198, tobool418, v199, result_symbol420, v200, mark_end421, v201, v202, v203, tobool422, v204, result_symbol424, v205, mark_end425, v206, v207, v208, tobool426, v209

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
	i = new(int32)
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
		goto sw_bb30
	case 2:
		goto sw_bb72
	case 3:
		goto sw_bb86
	case 4:
		goto sw_bb100
	case 5:
		goto sw_bb106
	case 6:
		goto sw_bb112
	case 7:
		goto sw_bb118
	case 8:
		goto sw_bb128
	case 9:
		goto sw_bb134
	case 10:
		goto sw_bb140
	case 11:
		goto sw_bb146
	case 12:
		goto sw_bb152
	case 13:
		goto sw_bb158
	case 14:
		goto sw_bb164
	case 15:
		goto sw_bb170
	case 16:
		goto sw_bb176
	case 17:
		goto sw_bb182
	case 18:
		goto sw_bb188
	case 19:
		goto sw_bb194
	case 20:
		goto sw_bb200
	case 21:
		goto sw_bb206
	case 22:
		goto sw_bb212
	case 23:
		goto sw_bb218
	case 24:
		goto sw_bb224
	case 25:
		goto sw_bb230
	case 26:
		goto sw_bb236
	case 27:
		goto sw_bb242
	case 28:
		goto sw_bb251
	case 29:
		goto sw_bb260
	case 30:
		goto sw_bb305
	case 31:
		goto sw_bb307
	case 32:
		goto sw_bb319
	case 33:
		goto sw_bb323
	case 34:
		goto sw_bb327
	case 35:
		goto sw_bb353
	case 36:
		goto sw_bb357
	case 37:
		goto sw_bb373
	case 38:
		goto sw_bb388
	case 39:
		goto sw_bb399
	case 40:
		goto sw_bb403
	case 41:
		goto sw_bb407
	case 42:
		goto sw_bb411
	case 43:
		goto sw_bb415
	case 44:
		goto sw_bb419
	case 45:
		goto sw_bb423
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
	*state_addr = 30
	goto next_state

if_end:
	*i = 0
	goto for_cond

for_cond:
	v11 = *i
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(22)
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
	cmp22 = 48 <= v21
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
	*state_addr = 37
	goto next_state

if_end28:
	v23 = *result
	tobool29 = (v23 & 1) != 0
	*retval = tobool29
	goto _return

sw_bb30:
	v24 = *lookahead
	cmp31 = v24 == 10
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*state_addr = 31
	goto next_state

if_end34:
	v25 = *lookahead
	cmp35 = v25 == 13
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*state_addr = 1
	goto next_state

if_end38:
	v26 = *lookahead
	cmp39 = 9 <= v26
	if cmp39 {
		goto land_lhs_true41
	} else {
		goto lor_lhs_false44
	}

land_lhs_true41:
	v27 = *lookahead
	cmp42 = v27 <= 12
	if cmp42 {
		goto if_then47
	} else {
		goto lor_lhs_false44
	}

lor_lhs_false44:
	v28 = *lookahead
	cmp45 = v28 == 32
	if cmp45 {
		goto if_then47
	} else {
		goto if_end48
	}

if_then47:
	*skip = 1
	*state_addr = 1
	goto next_state

if_end48:
	v29 = *lookahead
	cmp49 = 47 <= v29
	if cmp49 {
		goto land_lhs_true51
	} else {
		goto lor_lhs_false54
	}

land_lhs_true51:
	v30 = *lookahead
	cmp52 = v30 <= 57
	if cmp52 {
		goto if_then69
	} else {
		goto lor_lhs_false54
	}

lor_lhs_false54:
	v31 = *lookahead
	cmp55 = 65 <= v31
	if cmp55 {
		goto land_lhs_true57
	} else {
		goto lor_lhs_false60
	}

land_lhs_true57:
	v32 = *lookahead
	cmp58 = v32 <= 90
	if cmp58 {
		goto if_then69
	} else {
		goto lor_lhs_false60
	}

lor_lhs_false60:
	v33 = *lookahead
	cmp61 = v33 == 95
	if cmp61 {
		goto if_then69
	} else {
		goto lor_lhs_false63
	}

lor_lhs_false63:
	v34 = *lookahead
	cmp64 = 97 <= v34
	if cmp64 {
		goto land_lhs_true66
	} else {
		goto if_end70
	}

land_lhs_true66:
	v35 = *lookahead
	cmp67 = v35 <= 122
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*state_addr = 34
	goto next_state

if_end70:
	v36 = *result
	tobool71 = (v36 & 1) != 0
	*retval = tobool71
	goto _return

sw_bb72:
	v37 = *lookahead
	cmp73 = v37 == 39
	if cmp73 {
		goto if_then75
	} else {
		goto if_end76
	}

if_then75:
	*state_addr = 35
	goto next_state

if_end76:
	v38 = *lookahead
	cmp77 = v38 == 92
	if cmp77 {
		goto if_then79
	} else {
		goto if_end80
	}

if_then79:
	*state_addr = 3
	goto next_state

if_end80:
	v39 = *lookahead
	cmp81 = v39 != 0
	if cmp81 {
		goto if_then83
	} else {
		goto if_end84
	}

if_then83:
	*state_addr = 2
	goto next_state

if_end84:
	v40 = *result
	tobool85 = (v40 & 1) != 0
	*retval = tobool85
	goto _return

sw_bb86:
	v41 = *lookahead
	cmp87 = v41 == 39
	if cmp87 {
		goto if_then89
	} else {
		goto if_end90
	}

if_then89:
	*state_addr = 36
	goto next_state

if_end90:
	v42 = *lookahead
	cmp91 = v42 == 92
	if cmp91 {
		goto if_then93
	} else {
		goto if_end94
	}

if_then93:
	*state_addr = 3
	goto next_state

if_end94:
	v43 = *lookahead
	cmp95 = v43 != 0
	if cmp95 {
		goto if_then97
	} else {
		goto if_end98
	}

if_then97:
	*state_addr = 2
	goto next_state

if_end98:
	v44 = *result
	tobool99 = (v44 & 1) != 0
	*retval = tobool99
	goto _return

sw_bb100:
	v45 = *lookahead
	cmp101 = v45 == 97
	if cmp101 {
		goto if_then103
	} else {
		goto if_end104
	}

if_then103:
	*state_addr = 16
	goto next_state

if_end104:
	v46 = *result
	tobool105 = (v46 & 1) != 0
	*retval = tobool105
	goto _return

sw_bb106:
	v47 = *lookahead
	cmp107 = v47 == 97
	if cmp107 {
		goto if_then109
	} else {
		goto if_end110
	}

if_then109:
	*state_addr = 15
	goto next_state

if_end110:
	v48 = *result
	tobool111 = (v48 & 1) != 0
	*retval = tobool111
	goto _return

sw_bb112:
	v49 = *lookahead
	cmp113 = v49 == 97
	if cmp113 {
		goto if_then115
	} else {
		goto if_end116
	}

if_then115:
	*state_addr = 8
	goto next_state

if_end116:
	v50 = *result
	tobool117 = (v50 & 1) != 0
	*retval = tobool117
	goto _return

sw_bb118:
	v51 = *lookahead
	cmp119 = v51 == 98
	if cmp119 {
		goto if_then121
	} else {
		goto if_end122
	}

if_then121:
	*state_addr = 22
	goto next_state

if_end122:
	v52 = *lookahead
	cmp123 = v52 == 100
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*state_addr = 13
	goto next_state

if_end126:
	v53 = *result
	tobool127 = (v53 & 1) != 0
	*retval = tobool127
	goto _return

sw_bb128:
	v54 = *lookahead
	cmp129 = v54 == 98
	if cmp129 {
		goto if_then131
	} else {
		goto if_end132
	}

if_then131:
	*state_addr = 17
	goto next_state

if_end132:
	v55 = *result
	tobool133 = (v55 & 1) != 0
	*retval = tobool133
	goto _return

sw_bb134:
	v56 = *lookahead
	cmp135 = v56 == 100
	if cmp135 {
		goto if_then137
	} else {
		goto if_end138
	}

if_then137:
	*state_addr = 44
	goto next_state

if_end138:
	v57 = *result
	tobool139 = (v57 & 1) != 0
	*retval = tobool139
	goto _return

sw_bb140:
	v58 = *lookahead
	cmp141 = v58 == 101
	if cmp141 {
		goto if_then143
	} else {
		goto if_end144
	}

if_then143:
	*state_addr = 39
	goto next_state

if_end144:
	v59 = *result
	tobool145 = (v59 & 1) != 0
	*retval = tobool145
	goto _return

sw_bb146:
	v60 = *lookahead
	cmp147 = v60 == 101
	if cmp147 {
		goto if_then149
	} else {
		goto if_end150
	}

if_then149:
	*state_addr = 9
	goto next_state

if_end150:
	v61 = *result
	tobool151 = (v61 & 1) != 0
	*retval = tobool151
	goto _return

sw_bb152:
	v62 = *lookahead
	cmp153 = v62 == 101
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*state_addr = 5
	goto next_state

if_end156:
	v63 = *result
	tobool157 = (v63 & 1) != 0
	*retval = tobool157
	goto _return

sw_bb158:
	v64 = *lookahead
	cmp159 = v64 == 105
	if cmp159 {
		goto if_then161
	} else {
		goto if_end162
	}

if_then161:
	*state_addr = 24
	goto next_state

if_end162:
	v65 = *result
	tobool163 = (v65 & 1) != 0
	*retval = tobool163
	goto _return

sw_bb164:
	v66 = *lookahead
	cmp165 = v66 == 105
	if cmp165 {
		goto if_then167
	} else {
		goto if_end168
	}

if_then167:
	*state_addr = 18
	goto next_state

if_end168:
	v67 = *result
	tobool169 = (v67 & 1) != 0
	*retval = tobool169
	goto _return

sw_bb170:
	v68 = *lookahead
	cmp171 = v68 == 107
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*state_addr = 20
	goto next_state

if_end174:
	v69 = *result
	tobool175 = (v69 & 1) != 0
	*retval = tobool175
	goto _return

sw_bb176:
	v70 = *lookahead
	cmp177 = v70 == 108
	if cmp177 {
		goto if_then179
	} else {
		goto if_end180
	}

if_then179:
	*state_addr = 23
	goto next_state

if_end180:
	v71 = *result
	tobool181 = (v71 & 1) != 0
	*retval = tobool181
	goto _return

sw_bb182:
	v72 = *lookahead
	cmp183 = v72 == 108
	if cmp183 {
		goto if_then185
	} else {
		goto if_end186
	}

if_then185:
	*state_addr = 11
	goto next_state

if_end186:
	v73 = *result
	tobool187 = (v73 & 1) != 0
	*retval = tobool187
	goto _return

sw_bb188:
	v74 = *lookahead
	cmp189 = v74 == 110
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*state_addr = 25
	goto next_state

if_end192:
	v75 = *result
	tobool193 = (v75 & 1) != 0
	*retval = tobool193
	goto _return

sw_bb194:
	v76 = *lookahead
	cmp195 = v76 == 111
	if cmp195 {
		goto if_then197
	} else {
		goto if_end198
	}

if_then197:
	*state_addr = 14
	goto next_state

if_end198:
	v77 = *result
	tobool199 = (v77 & 1) != 0
	*retval = tobool199
	goto _return

sw_bb200:
	v78 = *lookahead
	cmp201 = v78 == 112
	if cmp201 {
		goto if_then203
	} else {
		goto if_end204
	}

if_then203:
	*state_addr = 19
	goto next_state

if_end204:
	v79 = *result
	tobool205 = (v79 & 1) != 0
	*retval = tobool205
	goto _return

sw_bb206:
	v80 = *lookahead
	cmp207 = v80 == 114
	if cmp207 {
		goto if_then209
	} else {
		goto if_end210
	}

if_then209:
	*state_addr = 26
	goto next_state

if_end210:
	v81 = *result
	tobool211 = (v81 & 1) != 0
	*retval = tobool211
	goto _return

sw_bb212:
	v82 = *lookahead
	cmp213 = v82 == 114
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*state_addr = 12
	goto next_state

if_end216:
	v83 = *result
	tobool217 = (v83 & 1) != 0
	*retval = tobool217
	goto _return

sw_bb218:
	v84 = *lookahead
	cmp219 = v84 == 115
	if cmp219 {
		goto if_then221
	} else {
		goto if_end222
	}

if_then221:
	*state_addr = 10
	goto next_state

if_end222:
	v85 = *result
	tobool223 = (v85 & 1) != 0
	*retval = tobool223
	goto _return

sw_bb224:
	v86 = *lookahead
	cmp225 = v86 == 115
	if cmp225 {
		goto if_then227
	} else {
		goto if_end228
	}

if_then227:
	*state_addr = 6
	goto next_state

if_end228:
	v87 = *result
	tobool229 = (v87 & 1) != 0
	*retval = tobool229
	goto _return

sw_bb230:
	v88 = *lookahead
	cmp231 = v88 == 116
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*state_addr = 45
	goto next_state

if_end234:
	v89 = *result
	tobool235 = (v89 & 1) != 0
	*retval = tobool235
	goto _return

sw_bb236:
	v90 = *lookahead
	cmp237 = v90 == 117
	if cmp237 {
		goto if_then239
	} else {
		goto if_end240
	}

if_then239:
	*state_addr = 10
	goto next_state

if_end240:
	v91 = *result
	tobool241 = (v91 & 1) != 0
	*retval = tobool241
	goto _return

sw_bb242:
	v92 = *lookahead
	cmp243 = 48 <= v92
	if cmp243 {
		goto land_lhs_true245
	} else {
		goto if_end249
	}

land_lhs_true245:
	v93 = *lookahead
	cmp246 = v93 <= 57
	if cmp246 {
		goto if_then248
	} else {
		goto if_end249
	}

if_then248:
	*state_addr = 37
	goto next_state

if_end249:
	v94 = *result
	tobool250 = (v94 & 1) != 0
	*retval = tobool250
	goto _return

sw_bb251:
	v95 = *lookahead
	cmp252 = 48 <= v95
	if cmp252 {
		goto land_lhs_true254
	} else {
		goto if_end258
	}

land_lhs_true254:
	v96 = *lookahead
	cmp255 = v96 <= 57
	if cmp255 {
		goto if_then257
	} else {
		goto if_end258
	}

if_then257:
	*state_addr = 38
	goto next_state

if_end258:
	v97 = *result
	tobool259 = (v97 & 1) != 0
	*retval = tobool259
	goto _return

sw_bb260:
	v98 = *eof
	tobool261 = (v98 & 1) != 0
	if tobool261 {
		goto if_then262
	} else {
		goto if_end263
	}

if_then262:
	*state_addr = 30
	goto next_state

if_end263:
	v99 = *lookahead
	cmp264 = v99 == 10
	if cmp264 {
		goto if_then266
	} else {
		goto if_end267
	}

if_then266:
	*state_addr = 31
	goto next_state

if_end267:
	v100 = *lookahead
	cmp268 = v100 == 13
	if cmp268 {
		goto if_then270
	} else {
		goto if_end271
	}

if_then270:
	*state_addr = 1
	goto next_state

if_end271:
	v101 = *lookahead
	cmp272 = 9 <= v101
	if cmp272 {
		goto land_lhs_true274
	} else {
		goto lor_lhs_false277
	}

land_lhs_true274:
	v102 = *lookahead
	cmp275 = v102 <= 12
	if cmp275 {
		goto if_then280
	} else {
		goto lor_lhs_false277
	}

lor_lhs_false277:
	v103 = *lookahead
	cmp278 = v103 == 32
	if cmp278 {
		goto if_then280
	} else {
		goto if_end281
	}

if_then280:
	*skip = 1
	*state_addr = 29
	goto next_state

if_end281:
	v104 = *lookahead
	cmp282 = 47 <= v104
	if cmp282 {
		goto land_lhs_true284
	} else {
		goto lor_lhs_false287
	}

land_lhs_true284:
	v105 = *lookahead
	cmp285 = v105 <= 57
	if cmp285 {
		goto if_then302
	} else {
		goto lor_lhs_false287
	}

lor_lhs_false287:
	v106 = *lookahead
	cmp288 = 65 <= v106
	if cmp288 {
		goto land_lhs_true290
	} else {
		goto lor_lhs_false293
	}

land_lhs_true290:
	v107 = *lookahead
	cmp291 = v107 <= 90
	if cmp291 {
		goto if_then302
	} else {
		goto lor_lhs_false293
	}

lor_lhs_false293:
	v108 = *lookahead
	cmp294 = v108 == 95
	if cmp294 {
		goto if_then302
	} else {
		goto lor_lhs_false296
	}

lor_lhs_false296:
	v109 = *lookahead
	cmp297 = 97 <= v109
	if cmp297 {
		goto land_lhs_true299
	} else {
		goto if_end303
	}

land_lhs_true299:
	v110 = *lookahead
	cmp300 = v110 <= 122
	if cmp300 {
		goto if_then302
	} else {
		goto if_end303
	}

if_then302:
	*state_addr = 34
	goto next_state

if_end303:
	v111 = *result
	tobool304 = (v111 & 1) != 0
	*retval = tobool304
	goto _return

sw_bb305:
	*result = 1
	v112 = *lexer_addr
	result_symbol = &v112.F1
	*result_symbol = 0
	v113 = *lexer_addr
	mark_end = &v113.F3
	v114 = *mark_end
	v115 = *lexer_addr
	v114(v115)
	v116 = *result
	tobool306 = (v116 & 1) != 0
	*retval = tobool306
	goto _return

sw_bb307:
	*result = 1
	v117 = *lexer_addr
	result_symbol308 = &v117.F1
	*result_symbol308 = 1
	v118 = *lexer_addr
	mark_end309 = &v118.F3
	v119 = *mark_end309
	v120 = *lexer_addr
	v119(v120)
	v121 = *lookahead
	cmp310 = v121 == 10
	if cmp310 {
		goto if_then312
	} else {
		goto if_end313
	}

if_then312:
	*state_addr = 31
	goto next_state

if_end313:
	v122 = *lookahead
	cmp314 = v122 == 13
	if cmp314 {
		goto if_then316
	} else {
		goto if_end317
	}

if_then316:
	*state_addr = 1
	goto next_state

if_end317:
	v123 = *result
	tobool318 = (v123 & 1) != 0
	*retval = tobool318
	goto _return

sw_bb319:
	*result = 1
	v124 = *lexer_addr
	result_symbol320 = &v124.F1
	*result_symbol320 = 2
	v125 = *lexer_addr
	mark_end321 = &v125.F3
	v126 = *mark_end321
	v127 = *lexer_addr
	v126(v127)
	v128 = *result
	tobool322 = (v128 & 1) != 0
	*retval = tobool322
	goto _return

sw_bb323:
	*result = 1
	v129 = *lexer_addr
	result_symbol324 = &v129.F1
	*result_symbol324 = 3
	v130 = *lexer_addr
	mark_end325 = &v130.F3
	v131 = *mark_end325
	v132 = *lexer_addr
	v131(v132)
	v133 = *result
	tobool326 = (v133 & 1) != 0
	*retval = tobool326
	goto _return

sw_bb327:
	*result = 1
	v134 = *lexer_addr
	result_symbol328 = &v134.F1
	*result_symbol328 = 4
	v135 = *lexer_addr
	mark_end329 = &v135.F3
	v136 = *mark_end329
	v137 = *lexer_addr
	v136(v137)
	v138 = *lookahead
	cmp330 = 47 <= v138
	if cmp330 {
		goto land_lhs_true332
	} else {
		goto lor_lhs_false335
	}

land_lhs_true332:
	v139 = *lookahead
	cmp333 = v139 <= 57
	if cmp333 {
		goto if_then350
	} else {
		goto lor_lhs_false335
	}

lor_lhs_false335:
	v140 = *lookahead
	cmp336 = 65 <= v140
	if cmp336 {
		goto land_lhs_true338
	} else {
		goto lor_lhs_false341
	}

land_lhs_true338:
	v141 = *lookahead
	cmp339 = v141 <= 90
	if cmp339 {
		goto if_then350
	} else {
		goto lor_lhs_false341
	}

lor_lhs_false341:
	v142 = *lookahead
	cmp342 = v142 == 95
	if cmp342 {
		goto if_then350
	} else {
		goto lor_lhs_false344
	}

lor_lhs_false344:
	v143 = *lookahead
	cmp345 = 97 <= v143
	if cmp345 {
		goto land_lhs_true347
	} else {
		goto if_end351
	}

land_lhs_true347:
	v144 = *lookahead
	cmp348 = v144 <= 122
	if cmp348 {
		goto if_then350
	} else {
		goto if_end351
	}

if_then350:
	*state_addr = 34
	goto next_state

if_end351:
	v145 = *result
	tobool352 = (v145 & 1) != 0
	*retval = tobool352
	goto _return

sw_bb353:
	*result = 1
	v146 = *lexer_addr
	result_symbol354 = &v146.F1
	*result_symbol354 = 5
	v147 = *lexer_addr
	mark_end355 = &v147.F3
	v148 = *mark_end355
	v149 = *lexer_addr
	v148(v149)
	v150 = *result
	tobool356 = (v150 & 1) != 0
	*retval = tobool356
	goto _return

sw_bb357:
	*result = 1
	v151 = *lexer_addr
	result_symbol358 = &v151.F1
	*result_symbol358 = 5
	v152 = *lexer_addr
	mark_end359 = &v152.F3
	v153 = *mark_end359
	v154 = *lexer_addr
	v153(v154)
	v155 = *lookahead
	cmp360 = v155 == 39
	if cmp360 {
		goto if_then362
	} else {
		goto if_end363
	}

if_then362:
	*state_addr = 35
	goto next_state

if_end363:
	v156 = *lookahead
	cmp364 = v156 == 92
	if cmp364 {
		goto if_then366
	} else {
		goto if_end367
	}

if_then366:
	*state_addr = 3
	goto next_state

if_end367:
	v157 = *lookahead
	cmp368 = v157 != 0
	if cmp368 {
		goto if_then370
	} else {
		goto if_end371
	}

if_then370:
	*state_addr = 2
	goto next_state

if_end371:
	v158 = *result
	tobool372 = (v158 & 1) != 0
	*retval = tobool372
	goto _return

sw_bb373:
	*result = 1
	v159 = *lexer_addr
	result_symbol374 = &v159.F1
	*result_symbol374 = 6
	v160 = *lexer_addr
	mark_end375 = &v160.F3
	v161 = *mark_end375
	v162 = *lexer_addr
	v161(v162)
	v163 = *lookahead
	cmp376 = v163 == 46
	if cmp376 {
		goto if_then378
	} else {
		goto if_end379
	}

if_then378:
	*state_addr = 28
	goto next_state

if_end379:
	v164 = *lookahead
	cmp380 = 48 <= v164
	if cmp380 {
		goto land_lhs_true382
	} else {
		goto if_end386
	}

land_lhs_true382:
	v165 = *lookahead
	cmp383 = v165 <= 57
	if cmp383 {
		goto if_then385
	} else {
		goto if_end386
	}

if_then385:
	*state_addr = 37
	goto next_state

if_end386:
	v166 = *result
	tobool387 = (v166 & 1) != 0
	*retval = tobool387
	goto _return

sw_bb388:
	*result = 1
	v167 = *lexer_addr
	result_symbol389 = &v167.F1
	*result_symbol389 = 6
	v168 = *lexer_addr
	mark_end390 = &v168.F3
	v169 = *mark_end390
	v170 = *lexer_addr
	v169(v170)
	v171 = *lookahead
	cmp391 = 48 <= v171
	if cmp391 {
		goto land_lhs_true393
	} else {
		goto if_end397
	}

land_lhs_true393:
	v172 = *lookahead
	cmp394 = v172 <= 57
	if cmp394 {
		goto if_then396
	} else {
		goto if_end397
	}

if_then396:
	*state_addr = 38
	goto next_state

if_end397:
	v173 = *result
	tobool398 = (v173 & 1) != 0
	*retval = tobool398
	goto _return

sw_bb399:
	*result = 1
	v174 = *lexer_addr
	result_symbol400 = &v174.F1
	*result_symbol400 = 7
	v175 = *lexer_addr
	mark_end401 = &v175.F3
	v176 = *mark_end401
	v177 = *lexer_addr
	v176(v177)
	v178 = *result
	tobool402 = (v178 & 1) != 0
	*retval = tobool402
	goto _return

sw_bb403:
	*result = 1
	v179 = *lexer_addr
	result_symbol404 = &v179.F1
	*result_symbol404 = 8
	v180 = *lexer_addr
	mark_end405 = &v180.F3
	v181 = *mark_end405
	v182 = *lexer_addr
	v181(v182)
	v183 = *result
	tobool406 = (v183 & 1) != 0
	*retval = tobool406
	goto _return

sw_bb407:
	*result = 1
	v184 = *lexer_addr
	result_symbol408 = &v184.F1
	*result_symbol408 = 9
	v185 = *lexer_addr
	mark_end409 = &v185.F3
	v186 = *mark_end409
	v187 = *lexer_addr
	v186(v187)
	v188 = *result
	tobool410 = (v188 & 1) != 0
	*retval = tobool410
	goto _return

sw_bb411:
	*result = 1
	v189 = *lexer_addr
	result_symbol412 = &v189.F1
	*result_symbol412 = 10
	v190 = *lexer_addr
	mark_end413 = &v190.F3
	v191 = *mark_end413
	v192 = *lexer_addr
	v191(v192)
	v193 = *result
	tobool414 = (v193 & 1) != 0
	*retval = tobool414
	goto _return

sw_bb415:
	*result = 1
	v194 = *lexer_addr
	result_symbol416 = &v194.F1
	*result_symbol416 = 11
	v195 = *lexer_addr
	mark_end417 = &v195.F3
	v196 = *mark_end417
	v197 = *lexer_addr
	v196(v197)
	v198 = *result
	tobool418 = (v198 & 1) != 0
	*retval = tobool418
	goto _return

sw_bb419:
	*result = 1
	v199 = *lexer_addr
	result_symbol420 = &v199.F1
	*result_symbol420 = 12
	v200 = *lexer_addr
	mark_end421 = &v200.F3
	v201 = *mark_end421
	v202 = *lexer_addr
	v201(v202)
	v203 = *result
	tobool422 = (v203 & 1) != 0
	*retval = tobool422
	goto _return

sw_bb423:
	*result = 1
	v204 = *lexer_addr
	result_symbol424 = &v204.F1
	*result_symbol424 = 13
	v205 = *lexer_addr
	mark_end425 = &v205.F3
	v206 = *mark_end425
	v207 = *lexer_addr
	v206(v207)
	v208 = *result
	tobool426 = (v208 & 1) != 0
	*retval = tobool426
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v209 = *retval
	return v209
}

