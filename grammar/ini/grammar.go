package grammar_ini

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

var tree_sitter_ini_language TSLanguage = TSLanguage{15, 19, 0, 10, 0, 30, 5, 2, 1, 4, &(*[5][19]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[78]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], &ts_field_names[0], &ts_field_map_slices[0], &ts_field_map_entries[0], &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{1, 4, 0}}

var ts_small_parse_table [407]int16 = [407]int16{
	7, 3, 1, 4, 5, 1, 8, 13, 1, 5, 8, 1, 17, 13, 1, 13,
	19, 2, 0, 1, 5, 2, 14, 15, 8, 3, 1, 4, 5, 1, 8, 9,
	1, 1, 15, 1, 0, 5, 1, 12, 9, 1, 18, 18, 1, 11, 6, 2,
	14, 15, 8, 3, 1, 4, 5, 1, 8, 9, 1, 1, 17, 1, 0, 5,
	1, 12, 9, 1, 18, 18, 1, 11, 7, 2, 14, 15, 7, 3, 1, 4,
	5, 1, 8, 13, 1, 5, 10, 1, 17, 13, 1, 13, 21, 2, 0, 1,
	8, 2, 14, 15, 7, 3, 1, 4, 5, 1, 8, 23, 1, 0, 25, 1,
	1, 5, 1, 12, 18, 1, 11, 9, 3, 14, 15, 18, 6, 3, 1, 4,
	5, 1, 8, 30, 1, 5, 13, 1, 13, 28, 2, 0, 1, 10, 3, 14,
	15, 17, 8, 3, 1, 4, 5, 1, 8, 9, 1, 1, 33, 1, 0, 5,
	1, 12, 9, 1, 18, 18, 1, 11, 11, 2, 14, 15, 4, 5, 1, 8,
	37, 1, 4, 35, 3, 0, 1, 5, 12, 3, 14, 15, 16, 4, 3, 1,
	4, 5, 1, 8, 13, 2, 14, 15, 40, 3, 0, 1, 5, 4, 3, 1,
	4, 5, 1, 8, 14, 2, 14, 15, 42, 3, 0, 1, 5, 4, 3, 1,
	4, 5, 1, 8, 15, 2, 14, 15, 44, 3, 0, 1, 5, 4, 3, 1,
	4, 5, 1, 8, 16, 2, 14, 15, 46, 3, 0, 1, 5, 3, 5, 1,
	8, 17, 2, 14, 15, 48, 4, 0, 1, 4, 5, 4, 3, 1, 4, 5,
	1, 8, 50, 2, 0, 1, 18, 2, 14, 15, 4, 52, 1, 2, 54, 1,
	4, 56, 1, 8, 19, 2, 14, 15, 4, 3, 1, 4, 5, 1, 8, 58,
	1, 0, 20, 2, 14, 15, 4, 3, 1, 4, 5, 1, 8, 60, 1, 6,
	21, 2, 14, 15, 4, 54, 1, 4, 56, 1, 8, 62, 1, 9, 22, 2,
	14, 15, 4, 3, 1, 4, 5, 1, 8, 64, 1, 3, 23, 2, 14, 15,
	4, 56, 1, 8, 66, 1, 4, 68, 1, 7, 24, 2, 14, 15, 3, 5,
	1, 8, 70, 1, 4, 25, 2, 14, 15, 3, 5, 1, 8, 72, 1, 4,
	26, 2, 14, 15, 3, 5, 1, 8, 74, 1, 4, 27, 2, 14, 15, 1,
	48, 1, 0, 1, 76, 1, 0,
}

var ts_small_parse_table_map [25]int32 = [25]int32{
	0, 24, 50, 76, 100, 124, 146, 172, 189, 205, 221, 237, 253, 267, 282, 296,
	310, 324, 338, 352, 366, 377, 388, 399, 403,
}

var ts_symbol_names [19]*byte = [19]*byte{
	&_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_5[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0], &_str_17[0],
	&_str_18[0], &_str_19[0], &_str_20[0],
}

var ts_field_names [2]*byte = [2]*byte{nil, &_str_21[0]}

var ts_field_map_slices [2]TSMapSlice = [2]TSMapSlice{TSMapSlice{}, TSMapSlice{0, 1}}

var ts_field_map_entries [1]TSFieldMapEntry = [1]TSFieldMapEntry{TSFieldMapEntry{1, 0, 0}}

var ts_symbol_metadata [19]TSSymbolMetadata = [19]TSSymbolMetadata{
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0},
	TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{},
}

var ts_symbol_map [19]int16 = [19]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 2, 10, 11, 12, 13, 14, 15,
	16, 17, 18,
}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [2][4]int16 = [2][4]int16{}

var ts_lex_modes [30]TSLexerMode = [30]TSLexerMode{
	TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0},
	TSLexerMode{5, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{20, 0, 0}, TSLexerMode{}, TSLexerMode{2, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{-1, 0, 0}, TSLexerMode{-1, 0, 0},
}

var ts_primary_state_ids [30]int16 = [30]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 17, 29,
}

var _str [4]byte = [4]byte{105, 110, 105, 0}

var ts_parse_table struct {
	F0 struct {
	F0 [9]int16
	F1 [10]int16
}
	F1 [19]int16
	F2 [19]int16
	F3 [19]int16
	F4 [19]int16
} = struct {
	F0 struct {
	F0 [9]int16
	F1 [10]int16
}
	F1 [19]int16
	F2 [19]int16
	F3 [19]int16
	F4 [19]int16
}{struct {
	F0 [9]int16
	F1 [10]int16
}{[9]int16{1, 1, 0, 1, 3, 0, 1, 0, 5}, [10]int16{}}, [19]int16{
	7, 9, 0, 0, 11, 13, 0, 0, 5, 0, 20, 18, 5, 13, 1, 1,
	2, 3, 6,
}, [19]int16{
	15, 9, 0, 0, 11, 13, 0, 0, 5, 0, 0, 18, 5, 13, 2, 2,
	12, 4, 7,
}, [19]int16{
	15, 9, 0, 0, 3, 13, 0, 0, 5, 0, 0, 18, 5, 13, 3, 3,
	0, 10, 7,
}, [19]int16{
	17, 9, 0, 0, 3, 13, 0, 0, 5, 0, 0, 18, 5, 13, 4, 4,
	0, 10, 11,
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
	F8 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F20 TSParseActionEntry
	F21 struct {
	F0 anon_1
	F1 [6]byte
}
	F22 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F31 TSParseActionEntry
	F32 struct {
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
	F53 struct {
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
	F54 struct {
	F0 anon_1
	F1 [6]byte
}
	F55 struct {
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
	F59 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F60 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F63 struct {
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
	F64 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F67 struct {
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
	F71 struct {
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
	F72 struct {
	F0 anon_1
	F1 [6]byte
}
	F73 struct {
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
	F74 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F8 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F20 TSParseActionEntry
	F21 struct {
	F0 anon_1
	F1 [6]byte
}
	F22 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F31 TSParseActionEntry
	F32 struct {
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
	F53 struct {
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
	F54 struct {
	F0 anon_1
	F1 [6]byte
}
	F55 struct {
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
	F59 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F60 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F63 struct {
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
	F64 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F67 struct {
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
	F71 struct {
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
	F72 struct {
	F0 anon_1
	F1 [6]byte
}
	F73 struct {
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
	F74 struct {
	F0 anon_1
	F1 [6]byte
}
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
}{0, 22, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 10, 0, 0}}}, struct {
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
}{0, 21, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 10, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 10, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 11, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 11, 0, 0}}}, struct {
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
}{0, 19, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
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
}{0, 21, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 10, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
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
}{0, 17, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 12, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 13, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 13, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 15, 0, 1}}}, struct {
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
}{0, 23, 0, 0}, [2]byte{}}}, struct {
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
}{0, 28, 0, 0}, [2]byte{}}}, struct {
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
}{0, 22, 0, 0}, [2]byte{}}}, struct {
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
}{0, 24, 0, 0}, [2]byte{}}}, struct {
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
}{0, 27, 0, 0}, [2]byte{}}}, struct {
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
}{0, 15, 0, 0}, [2]byte{}}}, struct {
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
}{0, 16, 0, 0}, [2]byte{}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 14, 0, 0}}}}

var _str_3 [4]byte = [4]byte{101, 110, 100, 0}

var _str_4 [2]byte = [2]byte{91, 0}

var _str_5 [5]byte = [5]byte{116, 101, 120, 116, 0}

var _str_6 [2]byte = [2]byte{93, 0}

var _str_7 [20]byte = [20]byte{
	115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 95, 116, 111, 107,
	101, 110, 50, 0,
}

var _str_8 [13]byte = [13]byte{115, 101, 116, 116, 105, 110, 103, 95, 110, 97, 109, 101, 0}

var _str_9 [2]byte = [2]byte{61, 0}

var _str_10 [14]byte = [14]byte{115, 101, 116, 116, 105, 110, 103, 95, 118, 97, 108, 117, 101, 0}

var _str_11 [15]byte = [15]byte{99, 111, 109, 109, 101, 110, 116, 95, 116, 111, 107, 101, 110, 49, 0}

var _str_12 [9]byte = [9]byte{100, 111, 99, 117, 109, 101, 110, 116, 0}

var _str_13 [8]byte = [8]byte{115, 101, 99, 116, 105, 111, 110, 0}

var _str_14 [13]byte = [13]byte{115, 101, 99, 116, 105, 111, 110, 95, 110, 97, 109, 101, 0}

var _str_15 [8]byte = [8]byte{115, 101, 116, 116, 105, 110, 103, 0}

var _str_16 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}

var _str_17 [7]byte = [7]byte{95, 98, 108, 97, 110, 107, 0}

var _str_18 [17]byte = [17]byte{
	100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49,
	0,
}

var _str_19 [17]byte = [17]byte{
	100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 50,
	0,
}

var _str_20 [17]byte = [17]byte{
	100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 51,
	0,
}

var _str_21 [6]byte = [6]byte{98, 108, 97, 110, 107, 0}

func tree_sitter_ini() *TSLanguage {
	return &tree_sitter_ini_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v63, v64, v66, v68, v69, v71, v73, v74, v76, v87, v88, v90, v96, v97, v99, v104, v105, v107, v109, v110, v112, v114, v115, v117, v127, v128, v130, v132, v133, v135, v139, v140, v142, v152, v153, v155, v159, v160, v162, v164, v165, v167, v172, v173, v175, v185, v186, v188 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end180, mark_end184, mark_end220, mark_end238, mark_end252, mark_end256, mark_end260, mark_end290, mark_end294, mark_end306, mark_end338, mark_end349, mark_end353, mark_end367, mark_end398 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, result_symbol, result_symbol179, result_symbol183, result_symbol219, result_symbol237, result_symbol251, result_symbol255, result_symbol259, result_symbol289, result_symbol293, result_symbol305, result_symbol337, result_symbol348, result_symbol352, result_symbol366, result_symbol397 *int16
	var lookahead, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp25, cmp29, cmp32, tobool36, cmp38, tobool42, cmp44, cmp48, cmp52, cmp55, cmp59, cmp62, cmp66, tobool70, cmp72, cmp76, cmp80, cmp83, cmp87, cmp90, cmp94, cmp96, cmp99, tobool103, cmp105, cmp109, cmp112, cmp115, cmp118, cmp121, cmp124, cmp127, tobool131, tobool133, cmp136, cmp140, cmp144, cmp148, cmp151, cmp155, cmp158, cmp162, cmp165, cmp168, cmp171, tobool175, tobool177, tobool181, cmp185, cmp189, cmp193, cmp196, cmp200, cmp203, cmp207, cmp210, cmp213, tobool217, cmp221, cmp225, cmp228, cmp231, tobool235, cmp239, cmp242, cmp245, tobool249, tobool253, tobool257, cmp261, cmp265, cmp268, cmp271, cmp274, cmp277, cmp280, cmp283, tobool287, tobool291, cmp295, cmp299, tobool303, cmp307, cmp311, cmp314, cmp318, cmp321, cmp325, cmp328, cmp331, tobool335, cmp339, cmp342, tobool346, tobool350, cmp354, cmp357, cmp360, tobool364, cmp368, cmp371, cmp375, cmp378, cmp382, cmp385, cmp388, cmp391, tobool395, cmp399, cmp402, cmp405, tobool409, v193 bool
	var v3, frombool, v10, v20, v22, v30, v40, v49, v50, v62, v67, v72, v86, v95, v103, v108, v113, v126, v131, v138, v151, v158, v163, v171, v184, v192 byte
	var v65, v70, v75, v89, v98, v106, v111, v116, v129, v134, v141, v154, v161, v166, v174, v187 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9 int16
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v21, v23, v24, v25, v26, v27, v28, v29, v31, v32, v33, v34, v35, v36, v37, v38, v39, v41, v42, v43, v44, v45, v46, v47, v48, v51, v52, v53, v54, v55, v56, v57, v58, v59, v60, v61, v77, v78, v79, v80, v81, v82, v83, v84, v85, v91, v92, v93, v94, v100, v101, v102, v118, v119, v120, v121, v122, v123, v124, v125, v136, v137, v143, v144, v145, v146, v147, v148, v149, v150, v156, v157, v168, v169, v170, v176, v177, v178, v179, v180, v181, v182, v183, v189, v190, v191 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp25, v18, cmp29, v19, cmp32, v20, tobool36, v21, cmp38, v22, tobool42, v23, cmp44, v24, cmp48, v25, cmp52, v26, cmp55, v27, cmp59, v28, cmp62, v29, cmp66, v30, tobool70, v31, cmp72, v32, cmp76, v33, cmp80, v34, cmp83, v35, cmp87, v36, cmp90, v37, cmp94, v38, cmp96, v39, cmp99, v40, tobool103, v41, cmp105, v42, cmp109, v43, cmp112, v44, cmp115, v45, cmp118, v46, cmp121, v47, cmp124, v48, cmp127, v49, tobool131, v50, tobool133, v51, cmp136, v52, cmp140, v53, cmp144, v54, cmp148, v55, cmp151, v56, cmp155, v57, cmp158, v58, cmp162, v59, cmp165, v60, cmp168, v61, cmp171, v62, tobool175, v63, result_symbol, v64, mark_end, v65, v66, v67, tobool177, v68, result_symbol179, v69, mark_end180, v70, v71, v72, tobool181, v73, result_symbol183, v74, mark_end184, v75, v76, v77, cmp185, v78, cmp189, v79, cmp193, v80, cmp196, v81, cmp200, v82, cmp203, v83, cmp207, v84, cmp210, v85, cmp213, v86, tobool217, v87, result_symbol219, v88, mark_end220, v89, v90, v91, cmp221, v92, cmp225, v93, cmp228, v94, cmp231, v95, tobool235, v96, result_symbol237, v97, mark_end238, v98, v99, v100, cmp239, v101, cmp242, v102, cmp245, v103, tobool249, v104, result_symbol251, v105, mark_end252, v106, v107, v108, tobool253, v109, result_symbol255, v110, mark_end256, v111, v112, v113, tobool257, v114, result_symbol259, v115, mark_end260, v116, v117, v118, cmp261, v119, cmp265, v120, cmp268, v121, cmp271, v122, cmp274, v123, cmp277, v124, cmp280, v125, cmp283, v126, tobool287, v127, result_symbol289, v128, mark_end290, v129, v130, v131, tobool291, v132, result_symbol293, v133, mark_end294, v134, v135, v136, cmp295, v137, cmp299, v138, tobool303, v139, result_symbol305, v140, mark_end306, v141, v142, v143, cmp307, v144, cmp311, v145, cmp314, v146, cmp318, v147, cmp321, v148, cmp325, v149, cmp328, v150, cmp331, v151, tobool335, v152, result_symbol337, v153, mark_end338, v154, v155, v156, cmp339, v157, cmp342, v158, tobool346, v159, result_symbol348, v160, mark_end349, v161, v162, v163, tobool350, v164, result_symbol352, v165, mark_end353, v166, v167, v168, cmp354, v169, cmp357, v170, cmp360, v171, tobool364, v172, result_symbol366, v173, mark_end367, v174, v175, v176, cmp368, v177, cmp371, v178, cmp375, v179, cmp378, v180, cmp382, v181, cmp385, v182, cmp388, v183, cmp391, v184, tobool395, v185, result_symbol397, v186, mark_end398, v187, v188, v189, cmp399, v190, cmp402, v191, cmp405, v192, tobool409, v193

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
		goto sw_bb37
	case 2:
		goto sw_bb43
	case 3:
		goto sw_bb71
	case 4:
		goto sw_bb104
	case 5:
		goto sw_bb132
	case 6:
		goto sw_bb176
	case 7:
		goto sw_bb178
	case 8:
		goto sw_bb182
	case 9:
		goto sw_bb218
	case 10:
		goto sw_bb236
	case 11:
		goto sw_bb250
	case 12:
		goto sw_bb254
	case 13:
		goto sw_bb258
	case 14:
		goto sw_bb288
	case 15:
		goto sw_bb292
	case 16:
		goto sw_bb304
	case 17:
		goto sw_bb336
	case 18:
		goto sw_bb347
	case 19:
		goto sw_bb351
	case 20:
		goto sw_bb365
	case 21:
		goto sw_bb396
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
	*state_addr = 6
	goto next_state

if_end:
	v11 = *lookahead
	cmp = v11 == 10
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*state_addr = 12
	goto next_state

if_end6:
	v12 = *lookahead
	cmp7 = v12 == 13
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*state_addr = 1
	goto next_state

if_end10:
	v13 = *lookahead
	cmp11 = v13 == 61
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*state_addr = 14
	goto next_state

if_end14:
	v14 = *lookahead
	cmp15 = v14 == 91
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*state_addr = 7
	goto next_state

if_end18:
	v15 = *lookahead
	cmp19 = v15 == 93
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*state_addr = 11
	goto next_state

if_end22:
	v16 = *lookahead
	cmp23 = v16 == 9
	if cmp23 {
		goto if_then27
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v17 = *lookahead
	cmp25 = v17 == 32
	if cmp25 {
		goto if_then27
	} else {
		goto if_end28
	}

if_then27:
	*skip = 1
	*state_addr = 0
	goto next_state

if_end28:
	v18 = *lookahead
	cmp29 = v18 == 35
	if cmp29 {
		goto if_then34
	} else {
		goto lor_lhs_false31
	}

lor_lhs_false31:
	v19 = *lookahead
	cmp32 = v19 == 59
	if cmp32 {
		goto if_then34
	} else {
		goto if_end35
	}

if_then34:
	*state_addr = 18
	goto next_state

if_end35:
	v20 = *result
	tobool36 = (v20 & 1) != 0
	*retval = tobool36
	goto _return

sw_bb37:
	v21 = *lookahead
	cmp38 = v21 == 10
	if cmp38 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*state_addr = 12
	goto next_state

if_end41:
	v22 = *result
	tobool42 = (v22 & 1) != 0
	*retval = tobool42
	goto _return

sw_bb43:
	v23 = *lookahead
	cmp44 = v23 == 10
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*state_addr = 12
	goto next_state

if_end47:
	v24 = *lookahead
	cmp48 = v24 == 13
	if cmp48 {
		goto if_then50
	} else {
		goto if_end51
	}

if_then50:
	*state_addr = 15
	goto next_state

if_end51:
	v25 = *lookahead
	cmp52 = v25 == 9
	if cmp52 {
		goto if_then57
	} else {
		goto lor_lhs_false54
	}

lor_lhs_false54:
	v26 = *lookahead
	cmp55 = v26 == 32
	if cmp55 {
		goto if_then57
	} else {
		goto if_end58
	}

if_then57:
	*state_addr = 16
	goto next_state

if_end58:
	v27 = *lookahead
	cmp59 = v27 == 35
	if cmp59 {
		goto if_then64
	} else {
		goto lor_lhs_false61
	}

lor_lhs_false61:
	v28 = *lookahead
	cmp62 = v28 == 59
	if cmp62 {
		goto if_then64
	} else {
		goto if_end65
	}

if_then64:
	*state_addr = 17
	goto next_state

if_end65:
	v29 = *lookahead
	cmp66 = v29 != 0
	if cmp66 {
		goto if_then68
	} else {
		goto if_end69
	}

if_then68:
	*state_addr = 17
	goto next_state

if_end69:
	v30 = *result
	tobool70 = (v30 & 1) != 0
	*retval = tobool70
	goto _return

sw_bb71:
	v31 = *lookahead
	cmp72 = v31 == 10
	if cmp72 {
		goto if_then74
	} else {
		goto if_end75
	}

if_then74:
	*state_addr = 10
	goto next_state

if_end75:
	v32 = *lookahead
	cmp76 = v32 == 13
	if cmp76 {
		goto if_then78
	} else {
		goto if_end79
	}

if_then78:
	*state_addr = 9
	goto next_state

if_end79:
	v33 = *lookahead
	cmp80 = v33 == 9
	if cmp80 {
		goto if_then85
	} else {
		goto lor_lhs_false82
	}

lor_lhs_false82:
	v34 = *lookahead
	cmp83 = v34 == 32
	if cmp83 {
		goto if_then85
	} else {
		goto if_end86
	}

if_then85:
	*state_addr = 8
	goto next_state

if_end86:
	v35 = *lookahead
	cmp87 = v35 == 35
	if cmp87 {
		goto if_then92
	} else {
		goto lor_lhs_false89
	}

lor_lhs_false89:
	v36 = *lookahead
	cmp90 = v36 == 59
	if cmp90 {
		goto if_then92
	} else {
		goto if_end93
	}

if_then92:
	*state_addr = 10
	goto next_state

if_end93:
	v37 = *lookahead
	cmp94 = v37 != 0
	if cmp94 {
		goto land_lhs_true
	} else {
		goto if_end102
	}

land_lhs_true:
	v38 = *lookahead
	cmp96 = v38 != 91
	if cmp96 {
		goto land_lhs_true98
	} else {
		goto if_end102
	}

land_lhs_true98:
	v39 = *lookahead
	cmp99 = v39 != 93
	if cmp99 {
		goto if_then101
	} else {
		goto if_end102
	}

if_then101:
	*state_addr = 10
	goto next_state

if_end102:
	v40 = *result
	tobool103 = (v40 & 1) != 0
	*retval = tobool103
	goto _return

sw_bb104:
	v41 = *lookahead
	cmp105 = v41 == 32
	if cmp105 {
		goto if_then107
	} else {
		goto if_end108
	}

if_then107:
	*state_addr = 4
	goto next_state

if_end108:
	v42 = *lookahead
	cmp109 = v42 != 0
	if cmp109 {
		goto land_lhs_true111
	} else {
		goto if_end130
	}

land_lhs_true111:
	v43 = *lookahead
	cmp112 = v43 < 9
	if cmp112 {
		goto land_lhs_true117
	} else {
		goto lor_lhs_false114
	}

lor_lhs_false114:
	v44 = *lookahead
	cmp115 = 13 < v44
	if cmp115 {
		goto land_lhs_true117
	} else {
		goto if_end130
	}

land_lhs_true117:
	v45 = *lookahead
	cmp118 = v45 != 35
	if cmp118 {
		goto land_lhs_true120
	} else {
		goto if_end130
	}

land_lhs_true120:
	v46 = *lookahead
	cmp121 = v46 != 59
	if cmp121 {
		goto land_lhs_true123
	} else {
		goto if_end130
	}

land_lhs_true123:
	v47 = *lookahead
	cmp124 = v47 != 61
	if cmp124 {
		goto land_lhs_true126
	} else {
		goto if_end130
	}

land_lhs_true126:
	v48 = *lookahead
	cmp127 = v48 != 91
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*state_addr = 13
	goto next_state

if_end130:
	v49 = *result
	tobool131 = (v49 & 1) != 0
	*retval = tobool131
	goto _return

sw_bb132:
	v50 = *eof
	tobool133 = (v50 & 1) != 0
	if tobool133 {
		goto if_then134
	} else {
		goto if_end135
	}

if_then134:
	*state_addr = 6
	goto next_state

if_end135:
	v51 = *lookahead
	cmp136 = v51 == 10
	if cmp136 {
		goto if_then138
	} else {
		goto if_end139
	}

if_then138:
	*state_addr = 12
	goto next_state

if_end139:
	v52 = *lookahead
	cmp140 = v52 == 13
	if cmp140 {
		goto if_then142
	} else {
		goto if_end143
	}

if_then142:
	*state_addr = 1
	goto next_state

if_end143:
	v53 = *lookahead
	cmp144 = v53 == 91
	if cmp144 {
		goto if_then146
	} else {
		goto if_end147
	}

if_then146:
	*state_addr = 7
	goto next_state

if_end147:
	v54 = *lookahead
	cmp148 = v54 == 9
	if cmp148 {
		goto if_then153
	} else {
		goto lor_lhs_false150
	}

lor_lhs_false150:
	v55 = *lookahead
	cmp151 = v55 == 32
	if cmp151 {
		goto if_then153
	} else {
		goto if_end154
	}

if_then153:
	*skip = 1
	*state_addr = 5
	goto next_state

if_end154:
	v56 = *lookahead
	cmp155 = v56 == 35
	if cmp155 {
		goto if_then160
	} else {
		goto lor_lhs_false157
	}

lor_lhs_false157:
	v57 = *lookahead
	cmp158 = v57 == 59
	if cmp158 {
		goto if_then160
	} else {
		goto if_end161
	}

if_then160:
	*state_addr = 18
	goto next_state

if_end161:
	v58 = *lookahead
	cmp162 = v58 != 0
	if cmp162 {
		goto land_lhs_true164
	} else {
		goto if_end174
	}

land_lhs_true164:
	v59 = *lookahead
	cmp165 = v59 < 9
	if cmp165 {
		goto land_lhs_true170
	} else {
		goto lor_lhs_false167
	}

lor_lhs_false167:
	v60 = *lookahead
	cmp168 = 13 < v60
	if cmp168 {
		goto land_lhs_true170
	} else {
		goto if_end174
	}

land_lhs_true170:
	v61 = *lookahead
	cmp171 = v61 != 61
	if cmp171 {
		goto if_then173
	} else {
		goto if_end174
	}

if_then173:
	*state_addr = 13
	goto next_state

if_end174:
	v62 = *result
	tobool175 = (v62 & 1) != 0
	*retval = tobool175
	goto _return

sw_bb176:
	*result = 1
	v63 = *lexer_addr
	result_symbol = &v63.F1
	*result_symbol = 0
	v64 = *lexer_addr
	mark_end = &v64.F3
	v65 = *mark_end
	v66 = *lexer_addr
	v65(v66)
	v67 = *result
	tobool177 = (v67 & 1) != 0
	*retval = tobool177
	goto _return

sw_bb178:
	*result = 1
	v68 = *lexer_addr
	result_symbol179 = &v68.F1
	*result_symbol179 = 1
	v69 = *lexer_addr
	mark_end180 = &v69.F3
	v70 = *mark_end180
	v71 = *lexer_addr
	v70(v71)
	v72 = *result
	tobool181 = (v72 & 1) != 0
	*retval = tobool181
	goto _return

sw_bb182:
	*result = 1
	v73 = *lexer_addr
	result_symbol183 = &v73.F1
	*result_symbol183 = 2
	v74 = *lexer_addr
	mark_end184 = &v74.F3
	v75 = *mark_end184
	v76 = *lexer_addr
	v75(v76)
	v77 = *lookahead
	cmp185 = v77 == 10
	if cmp185 {
		goto if_then187
	} else {
		goto if_end188
	}

if_then187:
	*state_addr = 10
	goto next_state

if_end188:
	v78 = *lookahead
	cmp189 = v78 == 13
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*state_addr = 9
	goto next_state

if_end192:
	v79 = *lookahead
	cmp193 = v79 == 9
	if cmp193 {
		goto if_then198
	} else {
		goto lor_lhs_false195
	}

lor_lhs_false195:
	v80 = *lookahead
	cmp196 = v80 == 32
	if cmp196 {
		goto if_then198
	} else {
		goto if_end199
	}

if_then198:
	*state_addr = 8
	goto next_state

if_end199:
	v81 = *lookahead
	cmp200 = v81 == 35
	if cmp200 {
		goto if_then205
	} else {
		goto lor_lhs_false202
	}

lor_lhs_false202:
	v82 = *lookahead
	cmp203 = v82 == 59
	if cmp203 {
		goto if_then205
	} else {
		goto if_end206
	}

if_then205:
	*state_addr = 10
	goto next_state

if_end206:
	v83 = *lookahead
	cmp207 = v83 != 0
	if cmp207 {
		goto land_lhs_true209
	} else {
		goto if_end216
	}

land_lhs_true209:
	v84 = *lookahead
	cmp210 = v84 != 91
	if cmp210 {
		goto land_lhs_true212
	} else {
		goto if_end216
	}

land_lhs_true212:
	v85 = *lookahead
	cmp213 = v85 != 93
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*state_addr = 10
	goto next_state

if_end216:
	v86 = *result
	tobool217 = (v86 & 1) != 0
	*retval = tobool217
	goto _return

sw_bb218:
	*result = 1
	v87 = *lexer_addr
	result_symbol219 = &v87.F1
	*result_symbol219 = 2
	v88 = *lexer_addr
	mark_end220 = &v88.F3
	v89 = *mark_end220
	v90 = *lexer_addr
	v89(v90)
	v91 = *lookahead
	cmp221 = v91 == 10
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*state_addr = 10
	goto next_state

if_end224:
	v92 = *lookahead
	cmp225 = v92 != 0
	if cmp225 {
		goto land_lhs_true227
	} else {
		goto if_end234
	}

land_lhs_true227:
	v93 = *lookahead
	cmp228 = v93 != 91
	if cmp228 {
		goto land_lhs_true230
	} else {
		goto if_end234
	}

land_lhs_true230:
	v94 = *lookahead
	cmp231 = v94 != 93
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*state_addr = 10
	goto next_state

if_end234:
	v95 = *result
	tobool235 = (v95 & 1) != 0
	*retval = tobool235
	goto _return

sw_bb236:
	*result = 1
	v96 = *lexer_addr
	result_symbol237 = &v96.F1
	*result_symbol237 = 2
	v97 = *lexer_addr
	mark_end238 = &v97.F3
	v98 = *mark_end238
	v99 = *lexer_addr
	v98(v99)
	v100 = *lookahead
	cmp239 = v100 != 0
	if cmp239 {
		goto land_lhs_true241
	} else {
		goto if_end248
	}

land_lhs_true241:
	v101 = *lookahead
	cmp242 = v101 != 91
	if cmp242 {
		goto land_lhs_true244
	} else {
		goto if_end248
	}

land_lhs_true244:
	v102 = *lookahead
	cmp245 = v102 != 93
	if cmp245 {
		goto if_then247
	} else {
		goto if_end248
	}

if_then247:
	*state_addr = 10
	goto next_state

if_end248:
	v103 = *result
	tobool249 = (v103 & 1) != 0
	*retval = tobool249
	goto _return

sw_bb250:
	*result = 1
	v104 = *lexer_addr
	result_symbol251 = &v104.F1
	*result_symbol251 = 3
	v105 = *lexer_addr
	mark_end252 = &v105.F3
	v106 = *mark_end252
	v107 = *lexer_addr
	v106(v107)
	v108 = *result
	tobool253 = (v108 & 1) != 0
	*retval = tobool253
	goto _return

sw_bb254:
	*result = 1
	v109 = *lexer_addr
	result_symbol255 = &v109.F1
	*result_symbol255 = 4
	v110 = *lexer_addr
	mark_end256 = &v110.F3
	v111 = *mark_end256
	v112 = *lexer_addr
	v111(v112)
	v113 = *result
	tobool257 = (v113 & 1) != 0
	*retval = tobool257
	goto _return

sw_bb258:
	*result = 1
	v114 = *lexer_addr
	result_symbol259 = &v114.F1
	*result_symbol259 = 5
	v115 = *lexer_addr
	mark_end260 = &v115.F3
	v116 = *mark_end260
	v117 = *lexer_addr
	v116(v117)
	v118 = *lookahead
	cmp261 = v118 == 32
	if cmp261 {
		goto if_then263
	} else {
		goto if_end264
	}

if_then263:
	*state_addr = 4
	goto next_state

if_end264:
	v119 = *lookahead
	cmp265 = v119 != 0
	if cmp265 {
		goto land_lhs_true267
	} else {
		goto if_end286
	}

land_lhs_true267:
	v120 = *lookahead
	cmp268 = v120 < 9
	if cmp268 {
		goto land_lhs_true273
	} else {
		goto lor_lhs_false270
	}

lor_lhs_false270:
	v121 = *lookahead
	cmp271 = 13 < v121
	if cmp271 {
		goto land_lhs_true273
	} else {
		goto if_end286
	}

land_lhs_true273:
	v122 = *lookahead
	cmp274 = v122 != 35
	if cmp274 {
		goto land_lhs_true276
	} else {
		goto if_end286
	}

land_lhs_true276:
	v123 = *lookahead
	cmp277 = v123 != 59
	if cmp277 {
		goto land_lhs_true279
	} else {
		goto if_end286
	}

land_lhs_true279:
	v124 = *lookahead
	cmp280 = v124 != 61
	if cmp280 {
		goto land_lhs_true282
	} else {
		goto if_end286
	}

land_lhs_true282:
	v125 = *lookahead
	cmp283 = v125 != 91
	if cmp283 {
		goto if_then285
	} else {
		goto if_end286
	}

if_then285:
	*state_addr = 13
	goto next_state

if_end286:
	v126 = *result
	tobool287 = (v126 & 1) != 0
	*retval = tobool287
	goto _return

sw_bb288:
	*result = 1
	v127 = *lexer_addr
	result_symbol289 = &v127.F1
	*result_symbol289 = 6
	v128 = *lexer_addr
	mark_end290 = &v128.F3
	v129 = *mark_end290
	v130 = *lexer_addr
	v129(v130)
	v131 = *result
	tobool291 = (v131 & 1) != 0
	*retval = tobool291
	goto _return

sw_bb292:
	*result = 1
	v132 = *lexer_addr
	result_symbol293 = &v132.F1
	*result_symbol293 = 7
	v133 = *lexer_addr
	mark_end294 = &v133.F3
	v134 = *mark_end294
	v135 = *lexer_addr
	v134(v135)
	v136 = *lookahead
	cmp295 = v136 == 10
	if cmp295 {
		goto if_then297
	} else {
		goto if_end298
	}

if_then297:
	*state_addr = 12
	goto next_state

if_end298:
	v137 = *lookahead
	cmp299 = v137 != 0
	if cmp299 {
		goto if_then301
	} else {
		goto if_end302
	}

if_then301:
	*state_addr = 17
	goto next_state

if_end302:
	v138 = *result
	tobool303 = (v138 & 1) != 0
	*retval = tobool303
	goto _return

sw_bb304:
	*result = 1
	v139 = *lexer_addr
	result_symbol305 = &v139.F1
	*result_symbol305 = 7
	v140 = *lexer_addr
	mark_end306 = &v140.F3
	v141 = *mark_end306
	v142 = *lexer_addr
	v141(v142)
	v143 = *lookahead
	cmp307 = v143 == 13
	if cmp307 {
		goto if_then309
	} else {
		goto if_end310
	}

if_then309:
	*state_addr = 15
	goto next_state

if_end310:
	v144 = *lookahead
	cmp311 = v144 == 9
	if cmp311 {
		goto if_then316
	} else {
		goto lor_lhs_false313
	}

lor_lhs_false313:
	v145 = *lookahead
	cmp314 = v145 == 32
	if cmp314 {
		goto if_then316
	} else {
		goto if_end317
	}

if_then316:
	*state_addr = 16
	goto next_state

if_end317:
	v146 = *lookahead
	cmp318 = v146 == 35
	if cmp318 {
		goto if_then323
	} else {
		goto lor_lhs_false320
	}

lor_lhs_false320:
	v147 = *lookahead
	cmp321 = v147 == 59
	if cmp321 {
		goto if_then323
	} else {
		goto if_end324
	}

if_then323:
	*state_addr = 17
	goto next_state

if_end324:
	v148 = *lookahead
	cmp325 = v148 != 0
	if cmp325 {
		goto land_lhs_true327
	} else {
		goto if_end334
	}

land_lhs_true327:
	v149 = *lookahead
	cmp328 = v149 != 9
	if cmp328 {
		goto land_lhs_true330
	} else {
		goto if_end334
	}

land_lhs_true330:
	v150 = *lookahead
	cmp331 = v150 != 10
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*state_addr = 17
	goto next_state

if_end334:
	v151 = *result
	tobool335 = (v151 & 1) != 0
	*retval = tobool335
	goto _return

sw_bb336:
	*result = 1
	v152 = *lexer_addr
	result_symbol337 = &v152.F1
	*result_symbol337 = 7
	v153 = *lexer_addr
	mark_end338 = &v153.F3
	v154 = *mark_end338
	v155 = *lexer_addr
	v154(v155)
	v156 = *lookahead
	cmp339 = v156 != 0
	if cmp339 {
		goto land_lhs_true341
	} else {
		goto if_end345
	}

land_lhs_true341:
	v157 = *lookahead
	cmp342 = v157 != 10
	if cmp342 {
		goto if_then344
	} else {
		goto if_end345
	}

if_then344:
	*state_addr = 17
	goto next_state

if_end345:
	v158 = *result
	tobool346 = (v158 & 1) != 0
	*retval = tobool346
	goto _return

sw_bb347:
	*result = 1
	v159 = *lexer_addr
	result_symbol348 = &v159.F1
	*result_symbol348 = 8
	v160 = *lexer_addr
	mark_end349 = &v160.F3
	v161 = *mark_end349
	v162 = *lexer_addr
	v161(v162)
	v163 = *result
	tobool350 = (v163 & 1) != 0
	*retval = tobool350
	goto _return

sw_bb351:
	*result = 1
	v164 = *lexer_addr
	result_symbol352 = &v164.F1
	*result_symbol352 = 8
	v165 = *lexer_addr
	mark_end353 = &v165.F3
	v166 = *mark_end353
	v167 = *lexer_addr
	v166(v167)
	v168 = *lookahead
	cmp354 = v168 != 0
	if cmp354 {
		goto land_lhs_true356
	} else {
		goto if_end363
	}

land_lhs_true356:
	v169 = *lookahead
	cmp357 = v169 != 10
	if cmp357 {
		goto land_lhs_true359
	} else {
		goto if_end363
	}

land_lhs_true359:
	v170 = *lookahead
	cmp360 = v170 != 13
	if cmp360 {
		goto if_then362
	} else {
		goto if_end363
	}

if_then362:
	*state_addr = 21
	goto next_state

if_end363:
	v171 = *result
	tobool364 = (v171 & 1) != 0
	*retval = tobool364
	goto _return

sw_bb365:
	*result = 1
	v172 = *lexer_addr
	result_symbol366 = &v172.F1
	*result_symbol366 = 9
	v173 = *lexer_addr
	mark_end367 = &v173.F3
	v174 = *mark_end367
	v175 = *lexer_addr
	v174(v175)
	v176 = *lookahead
	cmp368 = v176 == 9
	if cmp368 {
		goto if_then373
	} else {
		goto lor_lhs_false370
	}

lor_lhs_false370:
	v177 = *lookahead
	cmp371 = v177 == 32
	if cmp371 {
		goto if_then373
	} else {
		goto if_end374
	}

if_then373:
	*state_addr = 20
	goto next_state

if_end374:
	v178 = *lookahead
	cmp375 = v178 == 35
	if cmp375 {
		goto if_then380
	} else {
		goto lor_lhs_false377
	}

lor_lhs_false377:
	v179 = *lookahead
	cmp378 = v179 == 59
	if cmp378 {
		goto if_then380
	} else {
		goto if_end381
	}

if_then380:
	*state_addr = 19
	goto next_state

if_end381:
	v180 = *lookahead
	cmp382 = v180 != 0
	if cmp382 {
		goto land_lhs_true384
	} else {
		goto if_end394
	}

land_lhs_true384:
	v181 = *lookahead
	cmp385 = v181 != 9
	if cmp385 {
		goto land_lhs_true387
	} else {
		goto if_end394
	}

land_lhs_true387:
	v182 = *lookahead
	cmp388 = v182 != 10
	if cmp388 {
		goto land_lhs_true390
	} else {
		goto if_end394
	}

land_lhs_true390:
	v183 = *lookahead
	cmp391 = v183 != 13
	if cmp391 {
		goto if_then393
	} else {
		goto if_end394
	}

if_then393:
	*state_addr = 21
	goto next_state

if_end394:
	v184 = *result
	tobool395 = (v184 & 1) != 0
	*retval = tobool395
	goto _return

sw_bb396:
	*result = 1
	v185 = *lexer_addr
	result_symbol397 = &v185.F1
	*result_symbol397 = 9
	v186 = *lexer_addr
	mark_end398 = &v186.F3
	v187 = *mark_end398
	v188 = *lexer_addr
	v187(v188)
	v189 = *lookahead
	cmp399 = v189 != 0
	if cmp399 {
		goto land_lhs_true401
	} else {
		goto if_end408
	}

land_lhs_true401:
	v190 = *lookahead
	cmp402 = v190 != 10
	if cmp402 {
		goto land_lhs_true404
	} else {
		goto if_end408
	}

land_lhs_true404:
	v191 = *lookahead
	cmp405 = v191 != 13
	if cmp405 {
		goto if_then407
	} else {
		goto if_end408
	}

if_then407:
	*state_addr = 21
	goto next_state

if_end408:
	v192 = *result
	tobool409 = (v192 & 1) != 0
	*retval = tobool409
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v193 = *retval
	return v193
}

