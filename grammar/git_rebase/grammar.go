package grammar_git_rebase

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

var tree_sitter_git_rebase_language TSLanguage = TSLanguage{15, 20, 0, 10, 0, 27, 7, 1, 0, 5, &(*[7][20]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[85]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{0, 1, 0}}

var ts_small_parse_table [232]int16 = [232]int16{
	6, 3, 1, 9, 25, 1, 0, 27, 1, 1, 31, 1, 6, 26, 1, 17,
	29, 2, 4, 5, 5, 13, 1, 9, 33, 1, 0, 35, 1, 1, 8, 1,
	18, 38, 3, 2, 3, 8, 5, 13, 1, 9, 15, 1, 0, 40, 1, 1,
	3, 1, 18, 14, 1, 19, 5, 13, 1, 9, 42, 1, 0, 44, 1, 1,
	6, 1, 18, 10, 1, 19, 5, 3, 1, 9, 47, 1, 0, 49, 1, 1,
	51, 1, 6, 53, 1, 7, 5, 13, 1, 9, 19, 1, 0, 55, 1, 1,
	5, 1, 18, 13, 1, 19, 5, 13, 1, 9, 23, 1, 0, 57, 1, 1,
	4, 1, 18, 10, 1, 19, 5, 13, 1, 9, 19, 1, 0, 55, 1, 1,
	5, 1, 18, 10, 1, 19, 3, 3, 1, 9, 61, 1, 7, 59, 2, 0,
	1, 3, 3, 1, 9, 65, 1, 7, 63, 2, 0, 1, 2, 13, 1, 9,
	67, 2, 0, 1, 2, 13, 1, 9, 69, 2, 0, 1, 2, 13, 1, 9,
	71, 2, 0, 1, 2, 13, 1, 9, 73, 2, 0, 1, 2, 13, 1, 9,
	42, 2, 0, 1, 2, 13, 1, 9, 75, 2, 0, 1, 2, 13, 1, 9,
	77, 1, 0, 2, 3, 1, 9, 79, 1, 6, 2, 3, 1, 9, 81, 1,
	7, 2, 3, 1, 9, 83, 1, 6,
}

var ts_small_parse_table_map [20]int32 = [20]int32{
	0, 20, 38, 54, 70, 86, 102, 118, 134, 145, 156, 164, 172, 180, 188, 196,
	204, 211, 218, 225,
}

var ts_symbol_names [20]*byte = [20]*byte{
	&_str_3[0], &_str_4[0], &_str_5[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_5[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0],
	&_str_17[0], &_str_18[0], &_str_19[0], &_str_20[0],
}

var ts_symbol_metadata [20]TSSymbolMetadata = [20]TSSymbolMetadata{
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0},
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{},
}

var ts_symbol_map [20]int16 = [20]int16{
	0, 1, 8, 8, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19,
}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [1][5]int16 = [1][5]int16{}

var ts_lex_modes [27]TSLexerMode = [27]TSLexerMode{
	TSLexerMode{}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{4, 0, 0},
	TSLexerMode{4, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{1, 0, 0}, TSLexerMode{2, 0, 0},
}

var ts_primary_state_ids [27]int16 = [27]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
}

var _str [11]byte = [11]byte{103, 105, 116, 95, 114, 101, 98, 97, 115, 101, 0}

var ts_parse_table struct {
	F0 struct {
	F0 [10]int16
	F1 [10]int16
}
	F1 [20]int16
	F2 [20]int16
	F3 [20]int16
	F4 [20]int16
	F5 [20]int16
	F6 [20]int16
} = struct {
	F0 struct {
	F0 [10]int16
	F1 [10]int16
}
	F1 [20]int16
	F2 [20]int16
	F3 [20]int16
	F4 [20]int16
	F5 [20]int16
	F6 [20]int16
}{struct {
	F0 [10]int16
	F1 [10]int16
}{[10]int16{1, 1, 1, 1, 1, 1, 0, 1, 0, 3}, [10]int16{}}, [20]int16{
	5, 7, 9, 9, 0, 0, 0, 0, 11, 13, 23, 9, 17, 17, 17, 17,
	17, 0, 2, 0,
}, [20]int16{
	15, 17, 9, 9, 0, 0, 0, 0, 11, 13, 0, 12, 17, 17, 17, 17,
	17, 0, 8, 0,
}, [20]int16{
	19, 17, 9, 9, 0, 0, 0, 0, 11, 13, 0, 21, 17, 17, 17, 17,
	17, 0, 8, 0,
}, [20]int16{
	21, 17, 9, 9, 0, 0, 0, 0, 11, 13, 0, 21, 17, 17, 17, 17,
	17, 0, 8, 0,
}, [20]int16{
	23, 17, 9, 9, 0, 0, 0, 0, 11, 13, 0, 21, 17, 17, 17, 17,
	17, 0, 8, 0,
}, [20]int16{
	0, 17, 9, 9, 0, 0, 0, 0, 11, 13, 0, 21, 17, 17, 17, 17,
	17, 0, 8, 0,
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F38 struct {
	F0 anon_1
	F1 [6]byte
}
	F39 TSParseActionEntry
	F40 struct {
	F0 anon_1
	F1 [6]byte
}
	F41 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F64 TSParseActionEntry
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
	F78 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
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
	F84 struct {
	F0 struct {
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F38 struct {
	F0 anon_1
	F1 [6]byte
}
	F39 TSParseActionEntry
	F40 struct {
	F0 anon_1
	F1 [6]byte
}
	F41 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F64 TSParseActionEntry
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
	F78 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
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
	F84 struct {
	F0 struct {
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
}{0, 2, 0, 0}, [2]byte{}}}, struct {
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
}{0, 0, 1, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 10, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 10, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 10, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 10, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 12, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 12, 0, 0}}}, struct {
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
}{0, 15, 0, 0}, [2]byte{}}}, struct {
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
}{0, 8, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 18, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 19, 0, 0}}}, struct {
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
}{0, 6, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 15, 0, 0}}}, struct {
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
}{0, 16, 0, 0}, [2]byte{}}}, struct {
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
}{0, 4, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 13, 0, 0}}}, struct {
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
}{0, 19, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 14, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 11, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 13, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 4, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 14, 0, 0}}}, struct {
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
}{0, 18, 0, 0}, [2]byte{}}}, struct {
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
}{0, 11, 0, 0}, [2]byte{}}}}

var _str_3 [4]byte = [4]byte{101, 110, 100, 0}

var _str_4 [14]byte = [14]byte{115, 111, 117, 114, 99, 101, 95, 116, 111, 107, 101, 110, 49, 0}

var _str_5 [8]byte = [8]byte{99, 111, 109, 109, 97, 110, 100, 0}

var _str_6 [3]byte = [3]byte{45, 99, 0}

var _str_7 [3]byte = [3]byte{45, 67, 0}

var _str_8 [6]byte = [6]byte{108, 97, 98, 101, 108, 0}

var _str_9 [8]byte = [8]byte{109, 101, 115, 115, 97, 103, 101, 0}

var _str_10 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}

var _str_11 [7]byte = [7]byte{115, 111, 117, 114, 99, 101, 0}

var _str_12 [10]byte = [10]byte{111, 112, 101, 114, 97, 116, 105, 111, 110, 0}

var _str_13 [19]byte = [19]byte{
	95, 110, 117, 108, 108, 97, 114, 121, 95, 111, 112, 101, 114, 97, 116, 105,
	111, 110, 0,
}

var _str_14 [17]byte = [17]byte{
	95, 108, 97, 98, 101, 108, 95, 111, 112, 101, 114, 97, 116, 105, 111, 110,
	0,
}

var _str_15 [17]byte = [17]byte{
	95, 109, 101, 114, 103, 101, 95, 111, 112, 101, 114, 97, 116, 105, 111, 110,
	0,
}

var _str_16 [17]byte = [17]byte{
	95, 102, 105, 120, 117, 112, 95, 111, 112, 101, 114, 97, 116, 105, 111, 110,
	0,
}

var _str_17 [16]byte = [16]byte{
	95, 101, 120, 101, 99, 95, 111, 112, 101, 114, 97, 116, 105, 111, 110, 0,
}

var _str_18 [7]byte = [7]byte{111, 112, 116, 105, 111, 110, 0}

var _str_19 [15]byte = [15]byte{115, 111, 117, 114, 99, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}

var _str_20 [15]byte = [15]byte{115, 111, 117, 114, 99, 101, 95, 114, 101, 112, 101, 97, 116, 50, 0}

var ts_lex_map [16]int16 = [16]int16{
	45, 21, 63, 8, 101, 24, 120, 11, 10, 10, 13, 10, 35, 25, 59, 25,
}

func tree_sitter_git_rebase() *TSLanguage {
	return &tree_sitter_git_rebase_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v94, v95, v97, v99, v100, v102, v107, v108, v110, v119, v120, v122, v127, v128, v130, v132, v133, v135, v142, v143, v145, v147, v148, v150, v157, v158, v160, v162, v163, v165, v171, v172, v174, v176, v177, v179, v185, v186, v188, v196, v197, v199, v205, v206, v208, v215, v216, v218, v224, v225, v227, v233, v234, v236, v242, v243, v245, v250, v251, v253, v261, v262, v264, v272, v273, v275, v283, v284, v286, v293, v294, v296 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end266, mark_end281, mark_end309, mark_end323, mark_end327, mark_end347, mark_end351, mark_end371, mark_end375, mark_end392, mark_end396, mark_end413, mark_end438, mark_end455, mark_end477, mark_end495, mark_end513, mark_end531, mark_end545, mark_end569, mark_end593, mark_end617, mark_end637 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol265, result_symbol280, result_symbol308, result_symbol322, result_symbol326, result_symbol346, result_symbol350, result_symbol370, result_symbol374, result_symbol391, result_symbol395, result_symbol412, result_symbol437, result_symbol454, result_symbol476, result_symbol494, result_symbol512, result_symbol530, result_symbol544, result_symbol568, result_symbol592, result_symbol616, result_symbol636 *int16
	var lookahead, i, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, tobool26, cmp28, cmp31, cmp35, cmp38, cmp41, cmp44, cmp48, cmp51, cmp54, tobool58, cmp60, cmp63, cmp67, cmp70, cmp73, cmp76, cmp80, cmp83, cmp86, tobool90, tobool92, cmp95, cmp99, cmp103, cmp106, cmp110, cmp113, cmp117, cmp120, cmp123, cmp127, tobool131, tobool133, cmp136, cmp140, cmp143, cmp147, cmp150, cmp154, cmp157, cmp160, cmp164, tobool168, tobool170, cmp173, cmp177, cmp180, cmp184, cmp187, cmp191, cmp194, cmp197, cmp201, tobool205, tobool207, cmp210, cmp214, cmp218, cmp221, cmp225, cmp228, cmp231, cmp235, cmp238, cmp241, cmp245, cmp248, cmp251, cmp254, cmp257, tobool261, tobool263, cmp267, cmp271, cmp274, tobool278, cmp282, cmp286, cmp289, cmp293, cmp296, cmp299, cmp302, tobool306, cmp310, cmp313, cmp316, tobool320, tobool324, cmp328, cmp331, cmp334, cmp337, cmp340, tobool344, tobool348, cmp352, cmp355, cmp358, cmp361, cmp364, tobool368, tobool372, cmp376, cmp379, cmp382, cmp385, tobool389, tobool393, cmp397, cmp400, cmp403, cmp406, tobool410, cmp414, cmp418, cmp422, cmp425, cmp428, cmp431, tobool435, cmp439, cmp442, cmp445, cmp448, tobool452, cmp456, cmp460, cmp464, cmp467, cmp470, tobool474, cmp478, cmp482, cmp485, cmp488, tobool492, cmp496, cmp500, cmp503, cmp506, tobool510, cmp514, cmp518, cmp521, cmp524, tobool528, cmp532, cmp535, cmp538, tobool542, cmp546, cmp550, cmp553, cmp556, cmp559, cmp562, tobool566, cmp570, cmp574, cmp577, cmp580, cmp583, cmp586, tobool590, cmp594, cmp598, cmp601, cmp604, cmp607, cmp610, tobool614, cmp618, cmp621, cmp624, cmp627, cmp630, tobool634, cmp638, cmp641, cmp644, tobool648, v301 bool
	var v3, frombool, v10, v22, v32, v42, v43, v54, v55, v65, v66, v76, v77, v93, v98, v106, v118, v126, v131, v141, v146, v156, v161, v170, v175, v184, v195, v204, v214, v223, v232, v241, v249, v260, v271, v282, v292, v300 byte
	var v96, v101, v109, v121, v129, v134, v144, v149, v159, v164, v173, v178, v187, v198, v207, v217, v226, v235, v244, v252, v263, v274, v285, v295 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9, v13, v16 int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v23, v24, v25, v26, v27, v28, v29, v30, v31, v33, v34, v35, v36, v37, v38, v39, v40, v41, v44, v45, v46, v47, v48, v49, v50, v51, v52, v53, v56, v57, v58, v59, v60, v61, v62, v63, v64, v67, v68, v69, v70, v71, v72, v73, v74, v75, v78, v79, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v90, v91, v92, v103, v104, v105, v111, v112, v113, v114, v115, v116, v117, v123, v124, v125, v136, v137, v138, v139, v140, v151, v152, v153, v154, v155, v166, v167, v168, v169, v180, v181, v182, v183, v189, v190, v191, v192, v193, v194, v200, v201, v202, v203, v209, v210, v211, v212, v213, v219, v220, v221, v222, v228, v229, v230, v231, v237, v238, v239, v240, v246, v247, v248, v254, v255, v256, v257, v258, v259, v265, v266, v267, v268, v269, v270, v276, v277, v278, v279, v280, v281, v287, v288, v289, v290, v291, v297, v298, v299 int32
	var conv4, idxprom, idxprom10 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, tobool26, v23, cmp28, v24, cmp31, v25, cmp35, v26, cmp38, v27, cmp41, v28, cmp44, v29, cmp48, v30, cmp51, v31, cmp54, v32, tobool58, v33, cmp60, v34, cmp63, v35, cmp67, v36, cmp70, v37, cmp73, v38, cmp76, v39, cmp80, v40, cmp83, v41, cmp86, v42, tobool90, v43, tobool92, v44, cmp95, v45, cmp99, v46, cmp103, v47, cmp106, v48, cmp110, v49, cmp113, v50, cmp117, v51, cmp120, v52, cmp123, v53, cmp127, v54, tobool131, v55, tobool133, v56, cmp136, v57, cmp140, v58, cmp143, v59, cmp147, v60, cmp150, v61, cmp154, v62, cmp157, v63, cmp160, v64, cmp164, v65, tobool168, v66, tobool170, v67, cmp173, v68, cmp177, v69, cmp180, v70, cmp184, v71, cmp187, v72, cmp191, v73, cmp194, v74, cmp197, v75, cmp201, v76, tobool205, v77, tobool207, v78, cmp210, v79, cmp214, v80, cmp218, v81, cmp221, v82, cmp225, v83, cmp228, v84, cmp231, v85, cmp235, v86, cmp238, v87, cmp241, v88, cmp245, v89, cmp248, v90, cmp251, v91, cmp254, v92, cmp257, v93, tobool261, v94, result_symbol, v95, mark_end, v96, v97, v98, tobool263, v99, result_symbol265, v100, mark_end266, v101, v102, v103, cmp267, v104, cmp271, v105, cmp274, v106, tobool278, v107, result_symbol280, v108, mark_end281, v109, v110, v111, cmp282, v112, cmp286, v113, cmp289, v114, cmp293, v115, cmp296, v116, cmp299, v117, cmp302, v118, tobool306, v119, result_symbol308, v120, mark_end309, v121, v122, v123, cmp310, v124, cmp313, v125, cmp316, v126, tobool320, v127, result_symbol322, v128, mark_end323, v129, v130, v131, tobool324, v132, result_symbol326, v133, mark_end327, v134, v135, v136, cmp328, v137, cmp331, v138, cmp334, v139, cmp337, v140, cmp340, v141, tobool344, v142, result_symbol346, v143, mark_end347, v144, v145, v146, tobool348, v147, result_symbol350, v148, mark_end351, v149, v150, v151, cmp352, v152, cmp355, v153, cmp358, v154, cmp361, v155, cmp364, v156, tobool368, v157, result_symbol370, v158, mark_end371, v159, v160, v161, tobool372, v162, result_symbol374, v163, mark_end375, v164, v165, v166, cmp376, v167, cmp379, v168, cmp382, v169, cmp385, v170, tobool389, v171, result_symbol391, v172, mark_end392, v173, v174, v175, tobool393, v176, result_symbol395, v177, mark_end396, v178, v179, v180, cmp397, v181, cmp400, v182, cmp403, v183, cmp406, v184, tobool410, v185, result_symbol412, v186, mark_end413, v187, v188, v189, cmp414, v190, cmp418, v191, cmp422, v192, cmp425, v193, cmp428, v194, cmp431, v195, tobool435, v196, result_symbol437, v197, mark_end438, v198, v199, v200, cmp439, v201, cmp442, v202, cmp445, v203, cmp448, v204, tobool452, v205, result_symbol454, v206, mark_end455, v207, v208, v209, cmp456, v210, cmp460, v211, cmp464, v212, cmp467, v213, cmp470, v214, tobool474, v215, result_symbol476, v216, mark_end477, v217, v218, v219, cmp478, v220, cmp482, v221, cmp485, v222, cmp488, v223, tobool492, v224, result_symbol494, v225, mark_end495, v226, v227, v228, cmp496, v229, cmp500, v230, cmp503, v231, cmp506, v232, tobool510, v233, result_symbol512, v234, mark_end513, v235, v236, v237, cmp514, v238, cmp518, v239, cmp521, v240, cmp524, v241, tobool528, v242, result_symbol530, v243, mark_end531, v244, v245, v246, cmp532, v247, cmp535, v248, cmp538, v249, tobool542, v250, result_symbol544, v251, mark_end545, v252, v253, v254, cmp546, v255, cmp550, v256, cmp553, v257, cmp556, v258, cmp559, v259, cmp562, v260, tobool566, v261, result_symbol568, v262, mark_end569, v263, v264, v265, cmp570, v266, cmp574, v267, cmp577, v268, cmp580, v269, cmp583, v270, cmp586, v271, tobool590, v272, result_symbol592, v273, mark_end593, v274, v275, v276, cmp594, v277, cmp598, v278, cmp601, v279, cmp604, v280, cmp607, v281, cmp610, v282, tobool614, v283, result_symbol616, v284, mark_end617, v285, v286, v287, cmp618, v288, cmp621, v289, cmp624, v290, cmp627, v291, cmp630, v292, tobool634, v293, result_symbol636, v294, mark_end637, v295, v296, v297, cmp638, v298, cmp641, v299, cmp644, v300, tobool648, v301

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
		goto sw_bb27
	case 2:
		goto sw_bb59
	case 3:
		goto sw_bb91
	case 4:
		goto sw_bb132
	case 5:
		goto sw_bb169
	case 6:
		goto sw_bb206
	case 7:
		goto sw_bb262
	case 8:
		goto sw_bb264
	case 9:
		goto sw_bb279
	case 10:
		goto sw_bb307
	case 11:
		goto sw_bb321
	case 12:
		goto sw_bb325
	case 13:
		goto sw_bb345
	case 14:
		goto sw_bb349
	case 15:
		goto sw_bb369
	case 16:
		goto sw_bb373
	case 17:
		goto sw_bb390
	case 18:
		goto sw_bb394
	case 19:
		goto sw_bb411
	case 20:
		goto sw_bb436
	case 21:
		goto sw_bb453
	case 22:
		goto sw_bb475
	case 23:
		goto sw_bb493
	case 24:
		goto sw_bb511
	case 25:
		goto sw_bb529
	case 26:
		goto sw_bb543
	case 27:
		goto sw_bb567
	case 28:
		goto sw_bb591
	case 29:
		goto sw_bb615
	case 30:
		goto sw_bb635
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
	*state_addr = 7
	goto next_state

if_end:
	*i = 0
	goto for_cond

for_cond:
	v11 = *i
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(16)
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
	cmp16 = v19 <= 12
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
	cmp22 = v21 != 0
	if cmp22 {
		goto if_then24
	} else {
		goto if_end25
	}

if_then24:
	*state_addr = 25
	goto next_state

if_end25:
	v22 = *result
	tobool26 = (v22 & 1) != 0
	*retval = tobool26
	goto _return

sw_bb27:
	v23 = *lookahead
	cmp28 = v23 == 35
	if cmp28 {
		goto if_then33
	} else {
		goto lor_lhs_false30
	}

lor_lhs_false30:
	v24 = *lookahead
	cmp31 = v24 == 59
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*state_addr = 25
	goto next_state

if_end34:
	v25 = *lookahead
	cmp35 = v25 == 9
	if cmp35 {
		goto if_then46
	} else {
		goto lor_lhs_false37
	}

lor_lhs_false37:
	v26 = *lookahead
	cmp38 = v26 == 11
	if cmp38 {
		goto if_then46
	} else {
		goto lor_lhs_false40
	}

lor_lhs_false40:
	v27 = *lookahead
	cmp41 = v27 == 12
	if cmp41 {
		goto if_then46
	} else {
		goto lor_lhs_false43
	}

lor_lhs_false43:
	v28 = *lookahead
	cmp44 = v28 == 32
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*skip = 1
	*state_addr = 1
	goto next_state

if_end47:
	v29 = *lookahead
	cmp48 = v29 != 0
	if cmp48 {
		goto land_lhs_true50
	} else {
		goto if_end57
	}

land_lhs_true50:
	v30 = *lookahead
	cmp51 = v30 < 9
	if cmp51 {
		goto if_then56
	} else {
		goto lor_lhs_false53
	}

lor_lhs_false53:
	v31 = *lookahead
	cmp54 = 13 < v31
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*state_addr = 25
	goto next_state

if_end57:
	v32 = *result
	tobool58 = (v32 & 1) != 0
	*retval = tobool58
	goto _return

sw_bb59:
	v33 = *lookahead
	cmp60 = v33 == 35
	if cmp60 {
		goto if_then65
	} else {
		goto lor_lhs_false62
	}

lor_lhs_false62:
	v34 = *lookahead
	cmp63 = v34 == 59
	if cmp63 {
		goto if_then65
	} else {
		goto if_end66
	}

if_then65:
	*state_addr = 20
	goto next_state

if_end66:
	v35 = *lookahead
	cmp67 = v35 == 9
	if cmp67 {
		goto if_then78
	} else {
		goto lor_lhs_false69
	}

lor_lhs_false69:
	v36 = *lookahead
	cmp70 = v36 == 11
	if cmp70 {
		goto if_then78
	} else {
		goto lor_lhs_false72
	}

lor_lhs_false72:
	v37 = *lookahead
	cmp73 = v37 == 12
	if cmp73 {
		goto if_then78
	} else {
		goto lor_lhs_false75
	}

lor_lhs_false75:
	v38 = *lookahead
	cmp76 = v38 == 32
	if cmp76 {
		goto if_then78
	} else {
		goto if_end79
	}

if_then78:
	*skip = 1
	*state_addr = 2
	goto next_state

if_end79:
	v39 = *lookahead
	cmp80 = v39 != 0
	if cmp80 {
		goto land_lhs_true82
	} else {
		goto if_end89
	}

land_lhs_true82:
	v40 = *lookahead
	cmp83 = v40 < 9
	if cmp83 {
		goto if_then88
	} else {
		goto lor_lhs_false85
	}

lor_lhs_false85:
	v41 = *lookahead
	cmp86 = 13 < v41
	if cmp86 {
		goto if_then88
	} else {
		goto if_end89
	}

if_then88:
	*state_addr = 20
	goto next_state

if_end89:
	v42 = *result
	tobool90 = (v42 & 1) != 0
	*retval = tobool90
	goto _return

sw_bb91:
	v43 = *eof
	tobool92 = (v43 & 1) != 0
	if tobool92 {
		goto if_then93
	} else {
		goto if_end94
	}

if_then93:
	*state_addr = 7
	goto next_state

if_end94:
	v44 = *lookahead
	cmp95 = v44 == 45
	if cmp95 {
		goto if_then97
	} else {
		goto if_end98
	}

if_then97:
	*state_addr = 19
	goto next_state

if_end98:
	v45 = *lookahead
	cmp99 = v45 == 63
	if cmp99 {
		goto if_then101
	} else {
		goto if_end102
	}

if_then101:
	*state_addr = 9
	goto next_state

if_end102:
	v46 = *lookahead
	cmp103 = v46 == 10
	if cmp103 {
		goto if_then108
	} else {
		goto lor_lhs_false105
	}

lor_lhs_false105:
	v47 = *lookahead
	cmp106 = v47 == 13
	if cmp106 {
		goto if_then108
	} else {
		goto if_end109
	}

if_then108:
	*state_addr = 10
	goto next_state

if_end109:
	v48 = *lookahead
	cmp110 = v48 == 35
	if cmp110 {
		goto if_then115
	} else {
		goto lor_lhs_false112
	}

lor_lhs_false112:
	v49 = *lookahead
	cmp113 = v49 == 59
	if cmp113 {
		goto if_then115
	} else {
		goto if_end116
	}

if_then115:
	*state_addr = 20
	goto next_state

if_end116:
	v50 = *lookahead
	cmp117 = 9 <= v50
	if cmp117 {
		goto land_lhs_true119
	} else {
		goto lor_lhs_false122
	}

land_lhs_true119:
	v51 = *lookahead
	cmp120 = v51 <= 12
	if cmp120 {
		goto if_then125
	} else {
		goto lor_lhs_false122
	}

lor_lhs_false122:
	v52 = *lookahead
	cmp123 = v52 == 32
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*skip = 1
	*state_addr = 3
	goto next_state

if_end126:
	v53 = *lookahead
	cmp127 = v53 != 0
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*state_addr = 20
	goto next_state

if_end130:
	v54 = *result
	tobool131 = (v54 & 1) != 0
	*retval = tobool131
	goto _return

sw_bb132:
	v55 = *eof
	tobool133 = (v55 & 1) != 0
	if tobool133 {
		goto if_then134
	} else {
		goto if_end135
	}

if_then134:
	*state_addr = 7
	goto next_state

if_end135:
	v56 = *lookahead
	cmp136 = v56 == 63
	if cmp136 {
		goto if_then138
	} else {
		goto if_end139
	}

if_then138:
	*state_addr = 8
	goto next_state

if_end139:
	v57 = *lookahead
	cmp140 = v57 == 10
	if cmp140 {
		goto if_then145
	} else {
		goto lor_lhs_false142
	}

lor_lhs_false142:
	v58 = *lookahead
	cmp143 = v58 == 13
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*state_addr = 10
	goto next_state

if_end146:
	v59 = *lookahead
	cmp147 = v59 == 35
	if cmp147 {
		goto if_then152
	} else {
		goto lor_lhs_false149
	}

lor_lhs_false149:
	v60 = *lookahead
	cmp150 = v60 == 59
	if cmp150 {
		goto if_then152
	} else {
		goto if_end153
	}

if_then152:
	*state_addr = 25
	goto next_state

if_end153:
	v61 = *lookahead
	cmp154 = 9 <= v61
	if cmp154 {
		goto land_lhs_true156
	} else {
		goto lor_lhs_false159
	}

land_lhs_true156:
	v62 = *lookahead
	cmp157 = v62 <= 12
	if cmp157 {
		goto if_then162
	} else {
		goto lor_lhs_false159
	}

lor_lhs_false159:
	v63 = *lookahead
	cmp160 = v63 == 32
	if cmp160 {
		goto if_then162
	} else {
		goto if_end163
	}

if_then162:
	*skip = 1
	*state_addr = 4
	goto next_state

if_end163:
	v64 = *lookahead
	cmp164 = v64 != 0
	if cmp164 {
		goto if_then166
	} else {
		goto if_end167
	}

if_then166:
	*state_addr = 25
	goto next_state

if_end167:
	v65 = *result
	tobool168 = (v65 & 1) != 0
	*retval = tobool168
	goto _return

sw_bb169:
	v66 = *eof
	tobool170 = (v66 & 1) != 0
	if tobool170 {
		goto if_then171
	} else {
		goto if_end172
	}

if_then171:
	*state_addr = 7
	goto next_state

if_end172:
	v67 = *lookahead
	cmp173 = v67 == 63
	if cmp173 {
		goto if_then175
	} else {
		goto if_end176
	}

if_then175:
	*state_addr = 9
	goto next_state

if_end176:
	v68 = *lookahead
	cmp177 = v68 == 10
	if cmp177 {
		goto if_then182
	} else {
		goto lor_lhs_false179
	}

lor_lhs_false179:
	v69 = *lookahead
	cmp180 = v69 == 13
	if cmp180 {
		goto if_then182
	} else {
		goto if_end183
	}

if_then182:
	*state_addr = 10
	goto next_state

if_end183:
	v70 = *lookahead
	cmp184 = v70 == 35
	if cmp184 {
		goto if_then189
	} else {
		goto lor_lhs_false186
	}

lor_lhs_false186:
	v71 = *lookahead
	cmp187 = v71 == 59
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*state_addr = 20
	goto next_state

if_end190:
	v72 = *lookahead
	cmp191 = 9 <= v72
	if cmp191 {
		goto land_lhs_true193
	} else {
		goto lor_lhs_false196
	}

land_lhs_true193:
	v73 = *lookahead
	cmp194 = v73 <= 12
	if cmp194 {
		goto if_then199
	} else {
		goto lor_lhs_false196
	}

lor_lhs_false196:
	v74 = *lookahead
	cmp197 = v74 == 32
	if cmp197 {
		goto if_then199
	} else {
		goto if_end200
	}

if_then199:
	*skip = 1
	*state_addr = 5
	goto next_state

if_end200:
	v75 = *lookahead
	cmp201 = v75 != 0
	if cmp201 {
		goto if_then203
	} else {
		goto if_end204
	}

if_then203:
	*state_addr = 20
	goto next_state

if_end204:
	v76 = *result
	tobool205 = (v76 & 1) != 0
	*retval = tobool205
	goto _return

sw_bb206:
	v77 = *eof
	tobool207 = (v77 & 1) != 0
	if tobool207 {
		goto if_then208
	} else {
		goto if_end209
	}

if_then208:
	*state_addr = 7
	goto next_state

if_end209:
	v78 = *lookahead
	cmp210 = v78 == 101
	if cmp210 {
		goto if_then212
	} else {
		goto if_end213
	}

if_then212:
	*state_addr = 28
	goto next_state

if_end213:
	v79 = *lookahead
	cmp214 = v79 == 120
	if cmp214 {
		goto if_then216
	} else {
		goto if_end217
	}

if_then216:
	*state_addr = 12
	goto next_state

if_end217:
	v80 = *lookahead
	cmp218 = v80 == 35
	if cmp218 {
		goto if_then223
	} else {
		goto lor_lhs_false220
	}

lor_lhs_false220:
	v81 = *lookahead
	cmp221 = v81 == 59
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*state_addr = 30
	goto next_state

if_end224:
	v82 = *lookahead
	cmp225 = v82 == 10
	if cmp225 {
		goto if_then233
	} else {
		goto lor_lhs_false227
	}

lor_lhs_false227:
	v83 = *lookahead
	cmp228 = v83 == 13
	if cmp228 {
		goto if_then233
	} else {
		goto lor_lhs_false230
	}

lor_lhs_false230:
	v84 = *lookahead
	cmp231 = v84 == 63
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*state_addr = 10
	goto next_state

if_end234:
	v85 = *lookahead
	cmp235 = 9 <= v85
	if cmp235 {
		goto land_lhs_true237
	} else {
		goto lor_lhs_false240
	}

land_lhs_true237:
	v86 = *lookahead
	cmp238 = v86 <= 12
	if cmp238 {
		goto if_then243
	} else {
		goto lor_lhs_false240
	}

lor_lhs_false240:
	v87 = *lookahead
	cmp241 = v87 == 32
	if cmp241 {
		goto if_then243
	} else {
		goto if_end244
	}

if_then243:
	*skip = 1
	*state_addr = 6
	goto next_state

if_end244:
	v88 = *lookahead
	cmp245 = v88 == 45
	if cmp245 {
		goto if_then259
	} else {
		goto lor_lhs_false247
	}

lor_lhs_false247:
	v89 = *lookahead
	cmp248 = 65 <= v89
	if cmp248 {
		goto land_lhs_true250
	} else {
		goto lor_lhs_false253
	}

land_lhs_true250:
	v90 = *lookahead
	cmp251 = v90 <= 90
	if cmp251 {
		goto if_then259
	} else {
		goto lor_lhs_false253
	}

lor_lhs_false253:
	v91 = *lookahead
	cmp254 = 97 <= v91
	if cmp254 {
		goto land_lhs_true256
	} else {
		goto if_end260
	}

land_lhs_true256:
	v92 = *lookahead
	cmp257 = v92 <= 122
	if cmp257 {
		goto if_then259
	} else {
		goto if_end260
	}

if_then259:
	*state_addr = 29
	goto next_state

if_end260:
	v93 = *result
	tobool261 = (v93 & 1) != 0
	*retval = tobool261
	goto _return

sw_bb262:
	*result = 1
	v94 = *lexer_addr
	result_symbol = &v94.F1
	*result_symbol = 0
	v95 = *lexer_addr
	mark_end = &v95.F3
	v96 = *mark_end
	v97 = *lexer_addr
	v96(v97)
	v98 = *result
	tobool263 = (v98 & 1) != 0
	*retval = tobool263
	goto _return

sw_bb264:
	*result = 1
	v99 = *lexer_addr
	result_symbol265 = &v99.F1
	*result_symbol265 = 1
	v100 = *lexer_addr
	mark_end266 = &v100.F3
	v101 = *mark_end266
	v102 = *lexer_addr
	v101(v102)
	v103 = *lookahead
	cmp267 = v103 == 63
	if cmp267 {
		goto if_then269
	} else {
		goto if_end270
	}

if_then269:
	*state_addr = 8
	goto next_state

if_end270:
	v104 = *lookahead
	cmp271 = v104 == 10
	if cmp271 {
		goto if_then276
	} else {
		goto lor_lhs_false273
	}

lor_lhs_false273:
	v105 = *lookahead
	cmp274 = v105 == 13
	if cmp274 {
		goto if_then276
	} else {
		goto if_end277
	}

if_then276:
	*state_addr = 10
	goto next_state

if_end277:
	v106 = *result
	tobool278 = (v106 & 1) != 0
	*retval = tobool278
	goto _return

sw_bb279:
	*result = 1
	v107 = *lexer_addr
	result_symbol280 = &v107.F1
	*result_symbol280 = 1
	v108 = *lexer_addr
	mark_end281 = &v108.F3
	v109 = *mark_end281
	v110 = *lexer_addr
	v109(v110)
	v111 = *lookahead
	cmp282 = v111 == 63
	if cmp282 {
		goto if_then284
	} else {
		goto if_end285
	}

if_then284:
	*state_addr = 9
	goto next_state

if_end285:
	v112 = *lookahead
	cmp286 = v112 == 10
	if cmp286 {
		goto if_then291
	} else {
		goto lor_lhs_false288
	}

lor_lhs_false288:
	v113 = *lookahead
	cmp289 = v113 == 13
	if cmp289 {
		goto if_then291
	} else {
		goto if_end292
	}

if_then291:
	*state_addr = 10
	goto next_state

if_end292:
	v114 = *lookahead
	cmp293 = v114 != 0
	if cmp293 {
		goto land_lhs_true295
	} else {
		goto if_end305
	}

land_lhs_true295:
	v115 = *lookahead
	cmp296 = v115 < 9
	if cmp296 {
		goto land_lhs_true301
	} else {
		goto lor_lhs_false298
	}

lor_lhs_false298:
	v116 = *lookahead
	cmp299 = 13 < v116
	if cmp299 {
		goto land_lhs_true301
	} else {
		goto if_end305
	}

land_lhs_true301:
	v117 = *lookahead
	cmp302 = v117 != 32
	if cmp302 {
		goto if_then304
	} else {
		goto if_end305
	}

if_then304:
	*state_addr = 20
	goto next_state

if_end305:
	v118 = *result
	tobool306 = (v118 & 1) != 0
	*retval = tobool306
	goto _return

sw_bb307:
	*result = 1
	v119 = *lexer_addr
	result_symbol308 = &v119.F1
	*result_symbol308 = 1
	v120 = *lexer_addr
	mark_end309 = &v120.F3
	v121 = *mark_end309
	v122 = *lexer_addr
	v121(v122)
	v123 = *lookahead
	cmp310 = v123 == 10
	if cmp310 {
		goto if_then318
	} else {
		goto lor_lhs_false312
	}

lor_lhs_false312:
	v124 = *lookahead
	cmp313 = v124 == 13
	if cmp313 {
		goto if_then318
	} else {
		goto lor_lhs_false315
	}

lor_lhs_false315:
	v125 = *lookahead
	cmp316 = v125 == 63
	if cmp316 {
		goto if_then318
	} else {
		goto if_end319
	}

if_then318:
	*state_addr = 10
	goto next_state

if_end319:
	v126 = *result
	tobool320 = (v126 & 1) != 0
	*retval = tobool320
	goto _return

sw_bb321:
	*result = 1
	v127 = *lexer_addr
	result_symbol322 = &v127.F1
	*result_symbol322 = 2
	v128 = *lexer_addr
	mark_end323 = &v128.F3
	v129 = *mark_end323
	v130 = *lexer_addr
	v129(v130)
	v131 = *result
	tobool324 = (v131 & 1) != 0
	*retval = tobool324
	goto _return

sw_bb325:
	*result = 1
	v132 = *lexer_addr
	result_symbol326 = &v132.F1
	*result_symbol326 = 2
	v133 = *lexer_addr
	mark_end327 = &v133.F3
	v134 = *mark_end327
	v135 = *lexer_addr
	v134(v135)
	v136 = *lookahead
	cmp328 = v136 == 45
	if cmp328 {
		goto if_then342
	} else {
		goto lor_lhs_false330
	}

lor_lhs_false330:
	v137 = *lookahead
	cmp331 = 65 <= v137
	if cmp331 {
		goto land_lhs_true333
	} else {
		goto lor_lhs_false336
	}

land_lhs_true333:
	v138 = *lookahead
	cmp334 = v138 <= 90
	if cmp334 {
		goto if_then342
	} else {
		goto lor_lhs_false336
	}

lor_lhs_false336:
	v139 = *lookahead
	cmp337 = 97 <= v139
	if cmp337 {
		goto land_lhs_true339
	} else {
		goto if_end343
	}

land_lhs_true339:
	v140 = *lookahead
	cmp340 = v140 <= 122
	if cmp340 {
		goto if_then342
	} else {
		goto if_end343
	}

if_then342:
	*state_addr = 29
	goto next_state

if_end343:
	v141 = *result
	tobool344 = (v141 & 1) != 0
	*retval = tobool344
	goto _return

sw_bb345:
	*result = 1
	v142 = *lexer_addr
	result_symbol346 = &v142.F1
	*result_symbol346 = 3
	v143 = *lexer_addr
	mark_end347 = &v143.F3
	v144 = *mark_end347
	v145 = *lexer_addr
	v144(v145)
	v146 = *result
	tobool348 = (v146 & 1) != 0
	*retval = tobool348
	goto _return

sw_bb349:
	*result = 1
	v147 = *lexer_addr
	result_symbol350 = &v147.F1
	*result_symbol350 = 3
	v148 = *lexer_addr
	mark_end351 = &v148.F3
	v149 = *mark_end351
	v150 = *lexer_addr
	v149(v150)
	v151 = *lookahead
	cmp352 = v151 == 45
	if cmp352 {
		goto if_then366
	} else {
		goto lor_lhs_false354
	}

lor_lhs_false354:
	v152 = *lookahead
	cmp355 = 65 <= v152
	if cmp355 {
		goto land_lhs_true357
	} else {
		goto lor_lhs_false360
	}

land_lhs_true357:
	v153 = *lookahead
	cmp358 = v153 <= 90
	if cmp358 {
		goto if_then366
	} else {
		goto lor_lhs_false360
	}

lor_lhs_false360:
	v154 = *lookahead
	cmp361 = 97 <= v154
	if cmp361 {
		goto land_lhs_true363
	} else {
		goto if_end367
	}

land_lhs_true363:
	v155 = *lookahead
	cmp364 = v155 <= 122
	if cmp364 {
		goto if_then366
	} else {
		goto if_end367
	}

if_then366:
	*state_addr = 29
	goto next_state

if_end367:
	v156 = *result
	tobool368 = (v156 & 1) != 0
	*retval = tobool368
	goto _return

sw_bb369:
	*result = 1
	v157 = *lexer_addr
	result_symbol370 = &v157.F1
	*result_symbol370 = 4
	v158 = *lexer_addr
	mark_end371 = &v158.F3
	v159 = *mark_end371
	v160 = *lexer_addr
	v159(v160)
	v161 = *result
	tobool372 = (v161 & 1) != 0
	*retval = tobool372
	goto _return

sw_bb373:
	*result = 1
	v162 = *lexer_addr
	result_symbol374 = &v162.F1
	*result_symbol374 = 4
	v163 = *lexer_addr
	mark_end375 = &v163.F3
	v164 = *mark_end375
	v165 = *lexer_addr
	v164(v165)
	v166 = *lookahead
	cmp376 = v166 != 0
	if cmp376 {
		goto land_lhs_true378
	} else {
		goto if_end388
	}

land_lhs_true378:
	v167 = *lookahead
	cmp379 = v167 < 9
	if cmp379 {
		goto land_lhs_true384
	} else {
		goto lor_lhs_false381
	}

lor_lhs_false381:
	v168 = *lookahead
	cmp382 = 13 < v168
	if cmp382 {
		goto land_lhs_true384
	} else {
		goto if_end388
	}

land_lhs_true384:
	v169 = *lookahead
	cmp385 = v169 != 32
	if cmp385 {
		goto if_then387
	} else {
		goto if_end388
	}

if_then387:
	*state_addr = 20
	goto next_state

if_end388:
	v170 = *result
	tobool389 = (v170 & 1) != 0
	*retval = tobool389
	goto _return

sw_bb390:
	*result = 1
	v171 = *lexer_addr
	result_symbol391 = &v171.F1
	*result_symbol391 = 5
	v172 = *lexer_addr
	mark_end392 = &v172.F3
	v173 = *mark_end392
	v174 = *lexer_addr
	v173(v174)
	v175 = *result
	tobool393 = (v175 & 1) != 0
	*retval = tobool393
	goto _return

sw_bb394:
	*result = 1
	v176 = *lexer_addr
	result_symbol395 = &v176.F1
	*result_symbol395 = 5
	v177 = *lexer_addr
	mark_end396 = &v177.F3
	v178 = *mark_end396
	v179 = *lexer_addr
	v178(v179)
	v180 = *lookahead
	cmp397 = v180 != 0
	if cmp397 {
		goto land_lhs_true399
	} else {
		goto if_end409
	}

land_lhs_true399:
	v181 = *lookahead
	cmp400 = v181 < 9
	if cmp400 {
		goto land_lhs_true405
	} else {
		goto lor_lhs_false402
	}

lor_lhs_false402:
	v182 = *lookahead
	cmp403 = 13 < v182
	if cmp403 {
		goto land_lhs_true405
	} else {
		goto if_end409
	}

land_lhs_true405:
	v183 = *lookahead
	cmp406 = v183 != 32
	if cmp406 {
		goto if_then408
	} else {
		goto if_end409
	}

if_then408:
	*state_addr = 20
	goto next_state

if_end409:
	v184 = *result
	tobool410 = (v184 & 1) != 0
	*retval = tobool410
	goto _return

sw_bb411:
	*result = 1
	v185 = *lexer_addr
	result_symbol412 = &v185.F1
	*result_symbol412 = 6
	v186 = *lexer_addr
	mark_end413 = &v186.F3
	v187 = *mark_end413
	v188 = *lexer_addr
	v187(v188)
	v189 = *lookahead
	cmp414 = v189 == 67
	if cmp414 {
		goto if_then416
	} else {
		goto if_end417
	}

if_then416:
	*state_addr = 18
	goto next_state

if_end417:
	v190 = *lookahead
	cmp418 = v190 == 99
	if cmp418 {
		goto if_then420
	} else {
		goto if_end421
	}

if_then420:
	*state_addr = 16
	goto next_state

if_end421:
	v191 = *lookahead
	cmp422 = v191 != 0
	if cmp422 {
		goto land_lhs_true424
	} else {
		goto if_end434
	}

land_lhs_true424:
	v192 = *lookahead
	cmp425 = v192 < 9
	if cmp425 {
		goto land_lhs_true430
	} else {
		goto lor_lhs_false427
	}

lor_lhs_false427:
	v193 = *lookahead
	cmp428 = 13 < v193
	if cmp428 {
		goto land_lhs_true430
	} else {
		goto if_end434
	}

land_lhs_true430:
	v194 = *lookahead
	cmp431 = v194 != 32
	if cmp431 {
		goto if_then433
	} else {
		goto if_end434
	}

if_then433:
	*state_addr = 20
	goto next_state

if_end434:
	v195 = *result
	tobool435 = (v195 & 1) != 0
	*retval = tobool435
	goto _return

sw_bb436:
	*result = 1
	v196 = *lexer_addr
	result_symbol437 = &v196.F1
	*result_symbol437 = 6
	v197 = *lexer_addr
	mark_end438 = &v197.F3
	v198 = *mark_end438
	v199 = *lexer_addr
	v198(v199)
	v200 = *lookahead
	cmp439 = v200 != 0
	if cmp439 {
		goto land_lhs_true441
	} else {
		goto if_end451
	}

land_lhs_true441:
	v201 = *lookahead
	cmp442 = v201 < 9
	if cmp442 {
		goto land_lhs_true447
	} else {
		goto lor_lhs_false444
	}

lor_lhs_false444:
	v202 = *lookahead
	cmp445 = 13 < v202
	if cmp445 {
		goto land_lhs_true447
	} else {
		goto if_end451
	}

land_lhs_true447:
	v203 = *lookahead
	cmp448 = v203 != 32
	if cmp448 {
		goto if_then450
	} else {
		goto if_end451
	}

if_then450:
	*state_addr = 20
	goto next_state

if_end451:
	v204 = *result
	tobool452 = (v204 & 1) != 0
	*retval = tobool452
	goto _return

sw_bb453:
	*result = 1
	v205 = *lexer_addr
	result_symbol454 = &v205.F1
	*result_symbol454 = 7
	v206 = *lexer_addr
	mark_end455 = &v206.F3
	v207 = *mark_end455
	v208 = *lexer_addr
	v207(v208)
	v209 = *lookahead
	cmp456 = v209 == 67
	if cmp456 {
		goto if_then458
	} else {
		goto if_end459
	}

if_then458:
	*state_addr = 17
	goto next_state

if_end459:
	v210 = *lookahead
	cmp460 = v210 == 99
	if cmp460 {
		goto if_then462
	} else {
		goto if_end463
	}

if_then462:
	*state_addr = 15
	goto next_state

if_end463:
	v211 = *lookahead
	cmp464 = v211 != 0
	if cmp464 {
		goto land_lhs_true466
	} else {
		goto if_end473
	}

land_lhs_true466:
	v212 = *lookahead
	cmp467 = v212 != 10
	if cmp467 {
		goto land_lhs_true469
	} else {
		goto if_end473
	}

land_lhs_true469:
	v213 = *lookahead
	cmp470 = v213 != 13
	if cmp470 {
		goto if_then472
	} else {
		goto if_end473
	}

if_then472:
	*state_addr = 25
	goto next_state

if_end473:
	v214 = *result
	tobool474 = (v214 & 1) != 0
	*retval = tobool474
	goto _return

sw_bb475:
	*result = 1
	v215 = *lexer_addr
	result_symbol476 = &v215.F1
	*result_symbol476 = 7
	v216 = *lexer_addr
	mark_end477 = &v216.F3
	v217 = *mark_end477
	v218 = *lexer_addr
	v217(v218)
	v219 = *lookahead
	cmp478 = v219 == 99
	if cmp478 {
		goto if_then480
	} else {
		goto if_end481
	}

if_then480:
	*state_addr = 13
	goto next_state

if_end481:
	v220 = *lookahead
	cmp482 = v220 != 0
	if cmp482 {
		goto land_lhs_true484
	} else {
		goto if_end491
	}

land_lhs_true484:
	v221 = *lookahead
	cmp485 = v221 != 10
	if cmp485 {
		goto land_lhs_true487
	} else {
		goto if_end491
	}

land_lhs_true487:
	v222 = *lookahead
	cmp488 = v222 != 13
	if cmp488 {
		goto if_then490
	} else {
		goto if_end491
	}

if_then490:
	*state_addr = 25
	goto next_state

if_end491:
	v223 = *result
	tobool492 = (v223 & 1) != 0
	*retval = tobool492
	goto _return

sw_bb493:
	*result = 1
	v224 = *lexer_addr
	result_symbol494 = &v224.F1
	*result_symbol494 = 7
	v225 = *lexer_addr
	mark_end495 = &v225.F3
	v226 = *mark_end495
	v227 = *lexer_addr
	v226(v227)
	v228 = *lookahead
	cmp496 = v228 == 101
	if cmp496 {
		goto if_then498
	} else {
		goto if_end499
	}

if_then498:
	*state_addr = 22
	goto next_state

if_end499:
	v229 = *lookahead
	cmp500 = v229 != 0
	if cmp500 {
		goto land_lhs_true502
	} else {
		goto if_end509
	}

land_lhs_true502:
	v230 = *lookahead
	cmp503 = v230 != 10
	if cmp503 {
		goto land_lhs_true505
	} else {
		goto if_end509
	}

land_lhs_true505:
	v231 = *lookahead
	cmp506 = v231 != 13
	if cmp506 {
		goto if_then508
	} else {
		goto if_end509
	}

if_then508:
	*state_addr = 25
	goto next_state

if_end509:
	v232 = *result
	tobool510 = (v232 & 1) != 0
	*retval = tobool510
	goto _return

sw_bb511:
	*result = 1
	v233 = *lexer_addr
	result_symbol512 = &v233.F1
	*result_symbol512 = 7
	v234 = *lexer_addr
	mark_end513 = &v234.F3
	v235 = *mark_end513
	v236 = *lexer_addr
	v235(v236)
	v237 = *lookahead
	cmp514 = v237 == 120
	if cmp514 {
		goto if_then516
	} else {
		goto if_end517
	}

if_then516:
	*state_addr = 23
	goto next_state

if_end517:
	v238 = *lookahead
	cmp518 = v238 != 0
	if cmp518 {
		goto land_lhs_true520
	} else {
		goto if_end527
	}

land_lhs_true520:
	v239 = *lookahead
	cmp521 = v239 != 10
	if cmp521 {
		goto land_lhs_true523
	} else {
		goto if_end527
	}

land_lhs_true523:
	v240 = *lookahead
	cmp524 = v240 != 13
	if cmp524 {
		goto if_then526
	} else {
		goto if_end527
	}

if_then526:
	*state_addr = 25
	goto next_state

if_end527:
	v241 = *result
	tobool528 = (v241 & 1) != 0
	*retval = tobool528
	goto _return

sw_bb529:
	*result = 1
	v242 = *lexer_addr
	result_symbol530 = &v242.F1
	*result_symbol530 = 7
	v243 = *lexer_addr
	mark_end531 = &v243.F3
	v244 = *mark_end531
	v245 = *lexer_addr
	v244(v245)
	v246 = *lookahead
	cmp532 = v246 != 0
	if cmp532 {
		goto land_lhs_true534
	} else {
		goto if_end541
	}

land_lhs_true534:
	v247 = *lookahead
	cmp535 = v247 != 10
	if cmp535 {
		goto land_lhs_true537
	} else {
		goto if_end541
	}

land_lhs_true537:
	v248 = *lookahead
	cmp538 = v248 != 13
	if cmp538 {
		goto if_then540
	} else {
		goto if_end541
	}

if_then540:
	*state_addr = 25
	goto next_state

if_end541:
	v249 = *result
	tobool542 = (v249 & 1) != 0
	*retval = tobool542
	goto _return

sw_bb543:
	*result = 1
	v250 = *lexer_addr
	result_symbol544 = &v250.F1
	*result_symbol544 = 8
	v251 = *lexer_addr
	mark_end545 = &v251.F3
	v252 = *mark_end545
	v253 = *lexer_addr
	v252(v253)
	v254 = *lookahead
	cmp546 = v254 == 99
	if cmp546 {
		goto if_then548
	} else {
		goto if_end549
	}

if_then548:
	*state_addr = 14
	goto next_state

if_end549:
	v255 = *lookahead
	cmp550 = v255 == 45
	if cmp550 {
		goto if_then564
	} else {
		goto lor_lhs_false552
	}

lor_lhs_false552:
	v256 = *lookahead
	cmp553 = 65 <= v256
	if cmp553 {
		goto land_lhs_true555
	} else {
		goto lor_lhs_false558
	}

land_lhs_true555:
	v257 = *lookahead
	cmp556 = v257 <= 90
	if cmp556 {
		goto if_then564
	} else {
		goto lor_lhs_false558
	}

lor_lhs_false558:
	v258 = *lookahead
	cmp559 = 97 <= v258
	if cmp559 {
		goto land_lhs_true561
	} else {
		goto if_end565
	}

land_lhs_true561:
	v259 = *lookahead
	cmp562 = v259 <= 122
	if cmp562 {
		goto if_then564
	} else {
		goto if_end565
	}

if_then564:
	*state_addr = 29
	goto next_state

if_end565:
	v260 = *result
	tobool566 = (v260 & 1) != 0
	*retval = tobool566
	goto _return

sw_bb567:
	*result = 1
	v261 = *lexer_addr
	result_symbol568 = &v261.F1
	*result_symbol568 = 8
	v262 = *lexer_addr
	mark_end569 = &v262.F3
	v263 = *mark_end569
	v264 = *lexer_addr
	v263(v264)
	v265 = *lookahead
	cmp570 = v265 == 101
	if cmp570 {
		goto if_then572
	} else {
		goto if_end573
	}

if_then572:
	*state_addr = 26
	goto next_state

if_end573:
	v266 = *lookahead
	cmp574 = v266 == 45
	if cmp574 {
		goto if_then588
	} else {
		goto lor_lhs_false576
	}

lor_lhs_false576:
	v267 = *lookahead
	cmp577 = 65 <= v267
	if cmp577 {
		goto land_lhs_true579
	} else {
		goto lor_lhs_false582
	}

land_lhs_true579:
	v268 = *lookahead
	cmp580 = v268 <= 90
	if cmp580 {
		goto if_then588
	} else {
		goto lor_lhs_false582
	}

lor_lhs_false582:
	v269 = *lookahead
	cmp583 = 97 <= v269
	if cmp583 {
		goto land_lhs_true585
	} else {
		goto if_end589
	}

land_lhs_true585:
	v270 = *lookahead
	cmp586 = v270 <= 122
	if cmp586 {
		goto if_then588
	} else {
		goto if_end589
	}

if_then588:
	*state_addr = 29
	goto next_state

if_end589:
	v271 = *result
	tobool590 = (v271 & 1) != 0
	*retval = tobool590
	goto _return

sw_bb591:
	*result = 1
	v272 = *lexer_addr
	result_symbol592 = &v272.F1
	*result_symbol592 = 8
	v273 = *lexer_addr
	mark_end593 = &v273.F3
	v274 = *mark_end593
	v275 = *lexer_addr
	v274(v275)
	v276 = *lookahead
	cmp594 = v276 == 120
	if cmp594 {
		goto if_then596
	} else {
		goto if_end597
	}

if_then596:
	*state_addr = 27
	goto next_state

if_end597:
	v277 = *lookahead
	cmp598 = v277 == 45
	if cmp598 {
		goto if_then612
	} else {
		goto lor_lhs_false600
	}

lor_lhs_false600:
	v278 = *lookahead
	cmp601 = 65 <= v278
	if cmp601 {
		goto land_lhs_true603
	} else {
		goto lor_lhs_false606
	}

land_lhs_true603:
	v279 = *lookahead
	cmp604 = v279 <= 90
	if cmp604 {
		goto if_then612
	} else {
		goto lor_lhs_false606
	}

lor_lhs_false606:
	v280 = *lookahead
	cmp607 = 97 <= v280
	if cmp607 {
		goto land_lhs_true609
	} else {
		goto if_end613
	}

land_lhs_true609:
	v281 = *lookahead
	cmp610 = v281 <= 122
	if cmp610 {
		goto if_then612
	} else {
		goto if_end613
	}

if_then612:
	*state_addr = 29
	goto next_state

if_end613:
	v282 = *result
	tobool614 = (v282 & 1) != 0
	*retval = tobool614
	goto _return

sw_bb615:
	*result = 1
	v283 = *lexer_addr
	result_symbol616 = &v283.F1
	*result_symbol616 = 8
	v284 = *lexer_addr
	mark_end617 = &v284.F3
	v285 = *mark_end617
	v286 = *lexer_addr
	v285(v286)
	v287 = *lookahead
	cmp618 = v287 == 45
	if cmp618 {
		goto if_then632
	} else {
		goto lor_lhs_false620
	}

lor_lhs_false620:
	v288 = *lookahead
	cmp621 = 65 <= v288
	if cmp621 {
		goto land_lhs_true623
	} else {
		goto lor_lhs_false626
	}

land_lhs_true623:
	v289 = *lookahead
	cmp624 = v289 <= 90
	if cmp624 {
		goto if_then632
	} else {
		goto lor_lhs_false626
	}

lor_lhs_false626:
	v290 = *lookahead
	cmp627 = 97 <= v290
	if cmp627 {
		goto land_lhs_true629
	} else {
		goto if_end633
	}

land_lhs_true629:
	v291 = *lookahead
	cmp630 = v291 <= 122
	if cmp630 {
		goto if_then632
	} else {
		goto if_end633
	}

if_then632:
	*state_addr = 29
	goto next_state

if_end633:
	v292 = *result
	tobool634 = (v292 & 1) != 0
	*retval = tobool634
	goto _return

sw_bb635:
	*result = 1
	v293 = *lexer_addr
	result_symbol636 = &v293.F1
	*result_symbol636 = 9
	v294 = *lexer_addr
	mark_end637 = &v294.F3
	v295 = *mark_end637
	v296 = *lexer_addr
	v295(v296)
	v297 = *lookahead
	cmp638 = v297 != 0
	if cmp638 {
		goto land_lhs_true640
	} else {
		goto if_end647
	}

land_lhs_true640:
	v298 = *lookahead
	cmp641 = v298 != 10
	if cmp641 {
		goto land_lhs_true643
	} else {
		goto if_end647
	}

land_lhs_true643:
	v299 = *lookahead
	cmp644 = v299 != 13
	if cmp644 {
		goto if_then646
	} else {
		goto if_end647
	}

if_then646:
	*state_addr = 30
	goto next_state

if_end647:
	v300 = *result
	tobool648 = (v300 & 1) != 0
	*retval = tobool648
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v301 = *retval
	return v301
}

