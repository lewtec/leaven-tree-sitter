package grammar_tsv

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

var tree_sitter_tsv_language TSLanguage = TSLanguage{15, 18, 0, 10, 0, 27, 6, 1, 0, 2, &(*[6][18]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[83]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{1, 2, 0}}

var ts_small_parse_table [169]int16 = [169]int16{
	2, 15, 1, 0, 37, 7, 3, 4, 5, 6, 7, 8, 9, 4, 39, 1,
	0, 41, 1, 1, 43, 1, 2, 8, 1, 17, 4, 43, 1, 2, 45, 1,
	0, 47, 1, 1, 9, 1, 17, 4, 49, 1, 0, 51, 1, 1, 53, 1,
	2, 9, 1, 17, 2, 56, 1, 0, 58, 2, 1, 2, 2, 60, 1, 0,
	62, 2, 1, 2, 2, 64, 1, 0, 66, 2, 1, 2, 2, 68, 1, 0,
	70, 2, 1, 2, 3, 47, 1, 1, 72, 1, 2, 17, 1, 17, 2, 49,
	1, 0, 51, 2, 1, 2, 3, 41, 1, 1, 72, 1, 2, 14, 1, 17,
	3, 51, 1, 1, 74, 1, 2, 17, 1, 17, 2, 13, 1, 0, 77, 1,
	1, 1, 58, 2, 1, 2, 2, 77, 1, 1, 79, 1, 0, 1, 70, 2,
	1, 2, 1, 62, 2, 1, 2, 1, 66, 2, 1, 2, 1, 51, 2, 1,
	2, 1, 81, 1, 0, 1, 77, 1, 1,
}

var ts_small_parse_table_map [21]int32 = [21]int32{
	0, 13, 26, 39, 52, 60, 68, 76, 84, 94, 102, 112, 122, 129, 134, 141,
	146, 151, 156, 161, 165,
}

var ts_symbol_names [18]*byte = [18]*byte{
	&_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0], &_str_17[0], &_str_18[0],
	&_str_19[0], &_str_20[0],
}

var ts_symbol_metadata [18]TSSymbolMetadata = [18]TSSymbolMetadata{
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0},
	TSSymbolMetadata{}, TSSymbolMetadata{},
}

var ts_symbol_map [18]int16 = [18]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17,
}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [1][2]int16 = [1][2]int16{}

var ts_lex_modes [27]TSLexerMode = [27]TSLexerMode{
	TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0},
	TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{2, 0, 0}, TSLexerMode{}, TSLexerMode{3, 0, 0},
}

var ts_primary_state_ids [27]int16 = [27]int16{
	0, 1, 2, 3, 4, 4, 6, 7, 8, 9, 10, 11, 12, 13, 8, 15,
	7, 9, 18, 10, 20, 13, 11, 12, 15, 25, 26,
}

var _str [4]byte = [4]byte{116, 115, 118, 0}

var ts_parse_table struct {
	F0 struct {
	F0 [10]int16
	F1 [8]int16
}
	F1 [18]int16
	F2 [18]int16
	F3 [18]int16
	F4 [18]int16
	F5 [18]int16
} = struct {
	F0 struct {
	F0 [10]int16
	F1 [8]int16
}
	F1 [18]int16
	F2 [18]int16
	F3 [18]int16
	F4 [18]int16
	F5 [18]int16
}{struct {
	F0 [10]int16
	F1 [8]int16
}{[10]int16{1, 0, 0, 1, 1, 1, 1, 1, 1, 1}, [8]int16{}}, [18]int16{
	3, 0, 0, 5, 5, 7, 7, 9, 9, 11, 25, 18, 7, 12, 12, 12,
	2, 0,
}, [18]int16{
	13, 0, 0, 5, 5, 7, 7, 9, 9, 11, 0, 20, 7, 12, 12, 12,
	3, 0,
}, [18]int16{
	15, 0, 0, 17, 17, 20, 20, 23, 23, 26, 0, 26, 16, 23, 23, 23,
	3, 0,
}, [18]int16{
	0, 0, 0, 5, 5, 7, 7, 9, 9, 11, 0, 0, 15, 12, 12, 12,
	0, 0,
}, [18]int16{
	0, 0, 0, 29, 29, 31, 31, 33, 33, 35, 0, 0, 24, 23, 23, 23,
	0, 0,
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
	F14 TSParseActionEntry
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F44 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F54 TSParseActionEntry
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
	F57 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F69 TSParseActionEntry
	F70 struct {
	F0 anon_1
	F1 [6]byte
}
	F71 TSParseActionEntry
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
	F80 TSParseActionEntry
	F81 struct {
	F0 anon_1
	F1 [6]byte
}
	F82 struct {
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
	F14 TSParseActionEntry
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F44 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F54 TSParseActionEntry
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
	F57 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F69 TSParseActionEntry
	F70 struct {
	F0 anon_1
	F1 [6]byte
}
	F71 TSParseActionEntry
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
	F80 TSParseActionEntry
	F81 struct {
	F0 anon_1
	F1 [6]byte
}
	F82 struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 10, 0, 0}}}, struct {
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
}{0, 13, 0, 0}, [2]byte{}}}, struct {
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
}{0, 10, 0, 0}, [2]byte{}}}, struct {
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
}{0, 11, 0, 0}, [2]byte{}}}, struct {
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
}{0, 12, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 10, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
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
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
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
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
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
}{0, 22, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
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
}{0, 23, 0, 1}, [2]byte{}}}, struct {
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
}{0, 21, 0, 0}, [2]byte{}}}, struct {
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
}{0, 23, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 16, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 11, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 11, 0, 0}}}, struct {
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
}{0, 4, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 11, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 11, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 14, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 14, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 15, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 12, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 12, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 13, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 13, 0, 0}}}, struct {
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
}{0, 5, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 10, 0, 0}}}, struct {
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

var _str_3 [4]byte = [4]byte{101, 110, 100, 0}

var _str_4 [16]byte = [16]byte{
	100, 111, 99, 117, 109, 101, 110, 116, 95, 116, 111, 107, 101, 110, 49, 0,
}

var _str_5 [2]byte = [2]byte{9, 0}

var _str_6 [14]byte = [14]byte{110, 117, 109, 98, 101, 114, 95, 116, 111, 107, 101, 110, 49, 0}

var _str_7 [14]byte = [14]byte{110, 117, 109, 98, 101, 114, 95, 116, 111, 107, 101, 110, 50, 0}

var _str_8 [13]byte = [13]byte{102, 108, 111, 97, 116, 95, 116, 111, 107, 101, 110, 49, 0}

var _str_9 [13]byte = [13]byte{102, 108, 111, 97, 116, 95, 116, 111, 107, 101, 110, 50, 0}

var _str_10 [5]byte = [5]byte{116, 114, 117, 101, 0}

var _str_11 [6]byte = [6]byte{102, 97, 108, 115, 101, 0}

var _str_12 [5]byte = [5]byte{116, 101, 120, 116, 0}

var _str_13 [9]byte = [9]byte{100, 111, 99, 117, 109, 101, 110, 116, 0}

var _str_14 [4]byte = [4]byte{114, 111, 119, 0}

var _str_15 [6]byte = [6]byte{102, 105, 101, 108, 100, 0}

var _str_16 [7]byte = [7]byte{110, 117, 109, 98, 101, 114, 0}

var _str_17 [6]byte = [6]byte{102, 108, 111, 97, 116, 0}

var _str_18 [8]byte = [8]byte{98, 111, 111, 108, 101, 97, 110, 0}

var _str_19 [17]byte = [17]byte{
	100, 111, 99, 117, 109, 101, 110, 116, 95, 114, 101, 112, 101, 97, 116, 49,
	0,
}

var _str_20 [12]byte = [12]byte{114, 111, 119, 95, 114, 101, 112, 101, 97, 116, 49, 0}

func tree_sitter_tsv() *TSLanguage {
	return &tree_sitter_tsv_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v10, v11, v13, v47, v48, v50, v52, v53, v55, v60, v61, v63, v67, v68, v70, v75, v76, v78, v89, v90, v92, v101, v102, v104, v116, v117, v119, v127, v128, v130, v138, v139, v141, v147, v148, v150, v156, v157, v159, v174, v175, v177, v184, v185, v187, v194, v195, v197, v200, v201, v203, v210, v211, v213, v220, v221, v223, v230, v231, v233, v240, v241, v243, v250, v251, v253, v260, v261, v263, v270, v271, v273, v281, v282, v284, v296, v297, v299 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end111, mark_end115, mark_end131, mark_end143, mark_end159, mark_end194, mark_end222, mark_end258, mark_end282, mark_end306, mark_end323, mark_end340, mark_end391, mark_end412, mark_end434, mark_end442, mark_end463, mark_end484, mark_end505, mark_end526, mark_end547, mark_end568, mark_end589, mark_end613, mark_end649 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, result_symbol, result_symbol110, result_symbol114, result_symbol130, result_symbol142, result_symbol158, result_symbol193, result_symbol221, result_symbol257, result_symbol281, result_symbol305, result_symbol322, result_symbol339, result_symbol390, result_symbol411, result_symbol433, result_symbol441, result_symbol462, result_symbol483, result_symbol504, result_symbol525, result_symbol546, result_symbol567, result_symbol588, result_symbol612, result_symbol648 *int16
	var lookahead, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp25, cmp28, cmp32, cmp34, cmp38, cmp41, cmp44, tobool48, cmp50, cmp54, tobool58, tobool60, cmp63, cmp67, cmp71, cmp75, cmp78, cmp81, tobool85, tobool87, cmp90, cmp94, cmp98, cmp101, cmp104, tobool108, tobool112, cmp116, cmp120, cmp124, tobool128, cmp132, cmp136, tobool140, cmp144, cmp148, cmp152, tobool156, cmp160, cmp164, cmp167, cmp171, cmp174, cmp178, cmp181, cmp184, cmp187, tobool191, cmp195, cmp199, cmp202, cmp206, cmp209, cmp212, cmp215, tobool219, cmp223, cmp226, cmp229, cmp232, cmp235, cmp238, cmp242, cmp245, cmp248, cmp251, tobool255, cmp259, cmp262, cmp266, cmp269, cmp272, cmp275, tobool279, cmp283, cmp286, cmp290, cmp293, cmp296, cmp299, tobool303, cmp307, cmp310, cmp313, cmp316, tobool320, cmp324, cmp327, cmp330, cmp333, tobool337, cmp341, cmp345, cmp349, cmp353, cmp357, cmp361, cmp364, cmp367, cmp371, cmp374, cmp378, cmp381, cmp384, tobool388, cmp392, cmp396, cmp399, cmp402, cmp405, tobool409, cmp413, cmp417, cmp420, cmp423, cmp427, tobool431, cmp435, tobool439, cmp443, cmp447, cmp450, cmp453, cmp456, tobool460, cmp464, cmp468, cmp471, cmp474, cmp477, tobool481, cmp485, cmp489, cmp492, cmp495, cmp498, tobool502, cmp506, cmp510, cmp513, cmp516, cmp519, tobool523, cmp527, cmp531, cmp534, cmp537, cmp540, tobool544, cmp548, cmp552, cmp555, cmp558, cmp561, tobool565, cmp569, cmp573, cmp576, cmp579, cmp582, tobool586, cmp590, cmp593, cmp597, cmp600, cmp603, cmp606, tobool610, cmp614, cmp617, cmp620, cmp623, cmp626, cmp629, cmp633, cmp636, cmp639, cmp642, tobool646, cmp650, cmp653, cmp656, cmp659, tobool663, v305 bool
	var v3, frombool, v14, v28, v31, v32, v39, v40, v46, v51, v59, v66, v74, v88, v100, v115, v126, v137, v146, v155, v173, v183, v193, v199, v209, v219, v229, v239, v249, v259, v269, v280, v295, v304 byte
	var v12, v49, v54, v62, v69, v77, v91, v103, v118, v129, v140, v149, v158, v176, v186, v196, v202, v212, v222, v232, v242, v252, v262, v272, v283, v298 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9 int16
	var v5, conv, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v29, v30, v33, v34, v35, v36, v37, v38, v41, v42, v43, v44, v45, v56, v57, v58, v64, v65, v71, v72, v73, v79, v80, v81, v82, v83, v84, v85, v86, v87, v93, v94, v95, v96, v97, v98, v99, v105, v106, v107, v108, v109, v110, v111, v112, v113, v114, v120, v121, v122, v123, v124, v125, v131, v132, v133, v134, v135, v136, v142, v143, v144, v145, v151, v152, v153, v154, v160, v161, v162, v163, v164, v165, v166, v167, v168, v169, v170, v171, v172, v178, v179, v180, v181, v182, v188, v189, v190, v191, v192, v198, v204, v205, v206, v207, v208, v214, v215, v216, v217, v218, v224, v225, v226, v227, v228, v234, v235, v236, v237, v238, v244, v245, v246, v247, v248, v254, v255, v256, v257, v258, v264, v265, v266, v267, v268, v274, v275, v276, v277, v278, v279, v285, v286, v287, v288, v289, v290, v291, v292, v293, v294, v300, v301, v302, v303 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, result_symbol, v11, mark_end, v12, v13, v14, tobool3, v15, cmp, v16, cmp7, v17, cmp11, v18, cmp15, v19, cmp19, v20, cmp23, v21, cmp25, v22, cmp28, v23, cmp32, v24, cmp34, v25, cmp38, v26, cmp41, v27, cmp44, v28, tobool48, v29, cmp50, v30, cmp54, v31, tobool58, v32, tobool60, v33, cmp63, v34, cmp67, v35, cmp71, v36, cmp75, v37, cmp78, v38, cmp81, v39, tobool85, v40, tobool87, v41, cmp90, v42, cmp94, v43, cmp98, v44, cmp101, v45, cmp104, v46, tobool108, v47, result_symbol110, v48, mark_end111, v49, v50, v51, tobool112, v52, result_symbol114, v53, mark_end115, v54, v55, v56, cmp116, v57, cmp120, v58, cmp124, v59, tobool128, v60, result_symbol130, v61, mark_end131, v62, v63, v64, cmp132, v65, cmp136, v66, tobool140, v67, result_symbol142, v68, mark_end143, v69, v70, v71, cmp144, v72, cmp148, v73, cmp152, v74, tobool156, v75, result_symbol158, v76, mark_end159, v77, v78, v79, cmp160, v80, cmp164, v81, cmp167, v82, cmp171, v83, cmp174, v84, cmp178, v85, cmp181, v86, cmp184, v87, cmp187, v88, tobool191, v89, result_symbol193, v90, mark_end194, v91, v92, v93, cmp195, v94, cmp199, v95, cmp202, v96, cmp206, v97, cmp209, v98, cmp212, v99, cmp215, v100, tobool219, v101, result_symbol221, v102, mark_end222, v103, v104, v105, cmp223, v106, cmp226, v107, cmp229, v108, cmp232, v109, cmp235, v110, cmp238, v111, cmp242, v112, cmp245, v113, cmp248, v114, cmp251, v115, tobool255, v116, result_symbol257, v117, mark_end258, v118, v119, v120, cmp259, v121, cmp262, v122, cmp266, v123, cmp269, v124, cmp272, v125, cmp275, v126, tobool279, v127, result_symbol281, v128, mark_end282, v129, v130, v131, cmp283, v132, cmp286, v133, cmp290, v134, cmp293, v135, cmp296, v136, cmp299, v137, tobool303, v138, result_symbol305, v139, mark_end306, v140, v141, v142, cmp307, v143, cmp310, v144, cmp313, v145, cmp316, v146, tobool320, v147, result_symbol322, v148, mark_end323, v149, v150, v151, cmp324, v152, cmp327, v153, cmp330, v154, cmp333, v155, tobool337, v156, result_symbol339, v157, mark_end340, v158, v159, v160, cmp341, v161, cmp345, v162, cmp349, v163, cmp353, v164, cmp357, v165, cmp361, v166, cmp364, v167, cmp367, v168, cmp371, v169, cmp374, v170, cmp378, v171, cmp381, v172, cmp384, v173, tobool388, v174, result_symbol390, v175, mark_end391, v176, v177, v178, cmp392, v179, cmp396, v180, cmp399, v181, cmp402, v182, cmp405, v183, tobool409, v184, result_symbol411, v185, mark_end412, v186, v187, v188, cmp413, v189, cmp417, v190, cmp420, v191, cmp423, v192, cmp427, v193, tobool431, v194, result_symbol433, v195, mark_end434, v196, v197, v198, cmp435, v199, tobool439, v200, result_symbol441, v201, mark_end442, v202, v203, v204, cmp443, v205, cmp447, v206, cmp450, v207, cmp453, v208, cmp456, v209, tobool460, v210, result_symbol462, v211, mark_end463, v212, v213, v214, cmp464, v215, cmp468, v216, cmp471, v217, cmp474, v218, cmp477, v219, tobool481, v220, result_symbol483, v221, mark_end484, v222, v223, v224, cmp485, v225, cmp489, v226, cmp492, v227, cmp495, v228, cmp498, v229, tobool502, v230, result_symbol504, v231, mark_end505, v232, v233, v234, cmp506, v235, cmp510, v236, cmp513, v237, cmp516, v238, cmp519, v239, tobool523, v240, result_symbol525, v241, mark_end526, v242, v243, v244, cmp527, v245, cmp531, v246, cmp534, v247, cmp537, v248, cmp540, v249, tobool544, v250, result_symbol546, v251, mark_end547, v252, v253, v254, cmp548, v255, cmp552, v256, cmp555, v257, cmp558, v258, cmp561, v259, tobool565, v260, result_symbol567, v261, mark_end568, v262, v263, v264, cmp569, v265, cmp573, v266, cmp576, v267, cmp579, v268, cmp582, v269, tobool586, v270, result_symbol588, v271, mark_end589, v272, v273, v274, cmp590, v275, cmp593, v276, cmp597, v277, cmp600, v278, cmp603, v279, cmp606, v280, tobool610, v281, result_symbol612, v282, mark_end613, v283, v284, v285, cmp614, v286, cmp617, v287, cmp620, v288, cmp623, v289, cmp626, v290, cmp629, v291, cmp633, v292, cmp636, v293, cmp639, v294, cmp642, v295, tobool646, v296, result_symbol648, v297, mark_end649, v298, v299, v300, cmp650, v301, cmp653, v302, cmp656, v303, cmp659, v304, tobool663, v305

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
		goto sw_bb49
	case 2:
		goto sw_bb59
	case 3:
		goto sw_bb86
	case 4:
		goto sw_bb109
	case 5:
		goto sw_bb113
	case 6:
		goto sw_bb129
	case 7:
		goto sw_bb141
	case 8:
		goto sw_bb157
	case 9:
		goto sw_bb192
	case 10:
		goto sw_bb220
	case 11:
		goto sw_bb256
	case 12:
		goto sw_bb280
	case 13:
		goto sw_bb304
	case 14:
		goto sw_bb321
	case 15:
		goto sw_bb338
	case 16:
		goto sw_bb389
	case 17:
		goto sw_bb410
	case 18:
		goto sw_bb432
	case 19:
		goto sw_bb440
	case 20:
		goto sw_bb461
	case 21:
		goto sw_bb482
	case 22:
		goto sw_bb503
	case 23:
		goto sw_bb524
	case 24:
		goto sw_bb545
	case 25:
		goto sw_bb566
	case 26:
		goto sw_bb587
	case 27:
		goto sw_bb611
	case 28:
		goto sw_bb647
	default:
		goto sw_default
	}

sw_bb:
	*result = 1
	v10 = *lexer_addr
	result_symbol = &v10.F1
	*result_symbol = 9
	v11 = *lexer_addr
	mark_end = &v11.F3
	v12 = *mark_end
	v13 = *lexer_addr
	v12(v13)
	v14 = *eof
	tobool3 = (v14 & 1) != 0
	if tobool3 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*state_addr = 4
	goto next_state

if_end:
	v15 = *lookahead
	cmp = v15 == 34
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*state_addr = 17
	goto next_state

if_end6:
	v16 = *lookahead
	cmp7 = v16 == 46
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*state_addr = 26
	goto next_state

if_end10:
	v17 = *lookahead
	cmp11 = v17 == 48
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*state_addr = 8
	goto next_state

if_end14:
	v18 = *lookahead
	cmp15 = v18 == 102
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*state_addr = 19
	goto next_state

if_end18:
	v19 = *lookahead
	cmp19 = v19 == 116
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*state_addr = 23
	goto next_state

if_end22:
	v20 = *lookahead
	cmp23 = v20 == 11
	if cmp23 {
		goto if_then30
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v21 = *lookahead
	cmp25 = v21 == 12
	if cmp25 {
		goto if_then30
	} else {
		goto lor_lhs_false27
	}

lor_lhs_false27:
	v22 = *lookahead
	cmp28 = v22 == 32
	if cmp28 {
		goto if_then30
	} else {
		goto if_end31
	}

if_then30:
	*state_addr = 15
	goto next_state

if_end31:
	v23 = *lookahead
	cmp32 = 49 <= v23
	if cmp32 {
		goto land_lhs_true
	} else {
		goto if_end37
	}

land_lhs_true:
	v24 = *lookahead
	cmp34 = v24 <= 57
	if cmp34 {
		goto if_then36
	} else {
		goto if_end37
	}

if_then36:
	*state_addr = 9
	goto next_state

if_end37:
	v25 = *lookahead
	cmp38 = v25 != 0
	if cmp38 {
		goto land_lhs_true40
	} else {
		goto if_end47
	}

land_lhs_true40:
	v26 = *lookahead
	cmp41 = v26 < 9
	if cmp41 {
		goto if_then46
	} else {
		goto lor_lhs_false43
	}

lor_lhs_false43:
	v27 = *lookahead
	cmp44 = 13 < v27
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*state_addr = 28
	goto next_state

if_end47:
	v28 = *result
	tobool48 = (v28 & 1) != 0
	*retval = tobool48
	goto _return

sw_bb49:
	v29 = *lookahead
	cmp50 = v29 == 34
	if cmp50 {
		goto if_then52
	} else {
		goto if_end53
	}

if_then52:
	*state_addr = 18
	goto next_state

if_end53:
	v30 = *lookahead
	cmp54 = v30 != 0
	if cmp54 {
		goto if_then56
	} else {
		goto if_end57
	}

if_then56:
	*state_addr = 1
	goto next_state

if_end57:
	v31 = *result
	tobool58 = (v31 & 1) != 0
	*retval = tobool58
	goto _return

sw_bb59:
	v32 = *eof
	tobool60 = (v32 & 1) != 0
	if tobool60 {
		goto if_then61
	} else {
		goto if_end62
	}

if_then61:
	*state_addr = 4
	goto next_state

if_end62:
	v33 = *lookahead
	cmp63 = v33 == 9
	if cmp63 {
		goto if_then65
	} else {
		goto if_end66
	}

if_then65:
	*state_addr = 7
	goto next_state

if_end66:
	v34 = *lookahead
	cmp67 = v34 == 10
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*state_addr = 5
	goto next_state

if_end70:
	v35 = *lookahead
	cmp71 = v35 == 13
	if cmp71 {
		goto if_then73
	} else {
		goto if_end74
	}

if_then73:
	*state_addr = 5
	goto next_state

if_end74:
	v36 = *lookahead
	cmp75 = v36 == 11
	if cmp75 {
		goto if_then83
	} else {
		goto lor_lhs_false77
	}

lor_lhs_false77:
	v37 = *lookahead
	cmp78 = v37 == 12
	if cmp78 {
		goto if_then83
	} else {
		goto lor_lhs_false80
	}

lor_lhs_false80:
	v38 = *lookahead
	cmp81 = v38 == 32
	if cmp81 {
		goto if_then83
	} else {
		goto if_end84
	}

if_then83:
	*skip = 1
	*state_addr = 2
	goto next_state

if_end84:
	v39 = *result
	tobool85 = (v39 & 1) != 0
	*retval = tobool85
	goto _return

sw_bb86:
	v40 = *eof
	tobool87 = (v40 & 1) != 0
	if tobool87 {
		goto if_then88
	} else {
		goto if_end89
	}

if_then88:
	*state_addr = 4
	goto next_state

if_end89:
	v41 = *lookahead
	cmp90 = v41 == 10
	if cmp90 {
		goto if_then92
	} else {
		goto if_end93
	}

if_then92:
	*state_addr = 6
	goto next_state

if_end93:
	v42 = *lookahead
	cmp94 = v42 == 13
	if cmp94 {
		goto if_then96
	} else {
		goto if_end97
	}

if_then96:
	*state_addr = 6
	goto next_state

if_end97:
	v43 = *lookahead
	cmp98 = 9 <= v43
	if cmp98 {
		goto land_lhs_true100
	} else {
		goto lor_lhs_false103
	}

land_lhs_true100:
	v44 = *lookahead
	cmp101 = v44 <= 12
	if cmp101 {
		goto if_then106
	} else {
		goto lor_lhs_false103
	}

lor_lhs_false103:
	v45 = *lookahead
	cmp104 = v45 == 32
	if cmp104 {
		goto if_then106
	} else {
		goto if_end107
	}

if_then106:
	*skip = 1
	*state_addr = 3
	goto next_state

if_end107:
	v46 = *result
	tobool108 = (v46 & 1) != 0
	*retval = tobool108
	goto _return

sw_bb109:
	*result = 1
	v47 = *lexer_addr
	result_symbol110 = &v47.F1
	*result_symbol110 = 0
	v48 = *lexer_addr
	mark_end111 = &v48.F3
	v49 = *mark_end111
	v50 = *lexer_addr
	v49(v50)
	v51 = *result
	tobool112 = (v51 & 1) != 0
	*retval = tobool112
	goto _return

sw_bb113:
	*result = 1
	v52 = *lexer_addr
	result_symbol114 = &v52.F1
	*result_symbol114 = 1
	v53 = *lexer_addr
	mark_end115 = &v53.F3
	v54 = *mark_end115
	v55 = *lexer_addr
	v54(v55)
	v56 = *lookahead
	cmp116 = v56 == 9
	if cmp116 {
		goto if_then118
	} else {
		goto if_end119
	}

if_then118:
	*state_addr = 7
	goto next_state

if_end119:
	v57 = *lookahead
	cmp120 = v57 == 10
	if cmp120 {
		goto if_then122
	} else {
		goto if_end123
	}

if_then122:
	*state_addr = 5
	goto next_state

if_end123:
	v58 = *lookahead
	cmp124 = v58 == 13
	if cmp124 {
		goto if_then126
	} else {
		goto if_end127
	}

if_then126:
	*state_addr = 5
	goto next_state

if_end127:
	v59 = *result
	tobool128 = (v59 & 1) != 0
	*retval = tobool128
	goto _return

sw_bb129:
	*result = 1
	v60 = *lexer_addr
	result_symbol130 = &v60.F1
	*result_symbol130 = 1
	v61 = *lexer_addr
	mark_end131 = &v61.F3
	v62 = *mark_end131
	v63 = *lexer_addr
	v62(v63)
	v64 = *lookahead
	cmp132 = v64 == 10
	if cmp132 {
		goto if_then134
	} else {
		goto if_end135
	}

if_then134:
	*state_addr = 6
	goto next_state

if_end135:
	v65 = *lookahead
	cmp136 = v65 == 13
	if cmp136 {
		goto if_then138
	} else {
		goto if_end139
	}

if_then138:
	*state_addr = 6
	goto next_state

if_end139:
	v66 = *result
	tobool140 = (v66 & 1) != 0
	*retval = tobool140
	goto _return

sw_bb141:
	*result = 1
	v67 = *lexer_addr
	result_symbol142 = &v67.F1
	*result_symbol142 = 2
	v68 = *lexer_addr
	mark_end143 = &v68.F3
	v69 = *mark_end143
	v70 = *lexer_addr
	v69(v70)
	v71 = *lookahead
	cmp144 = v71 == 9
	if cmp144 {
		goto if_then146
	} else {
		goto if_end147
	}

if_then146:
	*state_addr = 7
	goto next_state

if_end147:
	v72 = *lookahead
	cmp148 = v72 == 10
	if cmp148 {
		goto if_then150
	} else {
		goto if_end151
	}

if_then150:
	*state_addr = 5
	goto next_state

if_end151:
	v73 = *lookahead
	cmp152 = v73 == 13
	if cmp152 {
		goto if_then154
	} else {
		goto if_end155
	}

if_then154:
	*state_addr = 5
	goto next_state

if_end155:
	v74 = *result
	tobool156 = (v74 & 1) != 0
	*retval = tobool156
	goto _return

sw_bb157:
	*result = 1
	v75 = *lexer_addr
	result_symbol158 = &v75.F1
	*result_symbol158 = 3
	v76 = *lexer_addr
	mark_end159 = &v76.F3
	v77 = *mark_end159
	v78 = *lexer_addr
	v77(v78)
	v79 = *lookahead
	cmp160 = v79 == 46
	if cmp160 {
		goto if_then162
	} else {
		goto if_end163
	}

if_then162:
	*state_addr = 12
	goto next_state

if_end163:
	v80 = *lookahead
	cmp164 = v80 == 88
	if cmp164 {
		goto if_then169
	} else {
		goto lor_lhs_false166
	}

lor_lhs_false166:
	v81 = *lookahead
	cmp167 = v81 == 120
	if cmp167 {
		goto if_then169
	} else {
		goto if_end170
	}

if_then169:
	*state_addr = 27
	goto next_state

if_end170:
	v82 = *lookahead
	cmp171 = 48 <= v82
	if cmp171 {
		goto land_lhs_true173
	} else {
		goto if_end177
	}

land_lhs_true173:
	v83 = *lookahead
	cmp174 = v83 <= 57
	if cmp174 {
		goto if_then176
	} else {
		goto if_end177
	}

if_then176:
	*state_addr = 9
	goto next_state

if_end177:
	v84 = *lookahead
	cmp178 = v84 != 0
	if cmp178 {
		goto land_lhs_true180
	} else {
		goto if_end190
	}

land_lhs_true180:
	v85 = *lookahead
	cmp181 = v85 != 9
	if cmp181 {
		goto land_lhs_true183
	} else {
		goto if_end190
	}

land_lhs_true183:
	v86 = *lookahead
	cmp184 = v86 != 10
	if cmp184 {
		goto land_lhs_true186
	} else {
		goto if_end190
	}

land_lhs_true186:
	v87 = *lookahead
	cmp187 = v87 != 13
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*state_addr = 28
	goto next_state

if_end190:
	v88 = *result
	tobool191 = (v88 & 1) != 0
	*retval = tobool191
	goto _return

sw_bb192:
	*result = 1
	v89 = *lexer_addr
	result_symbol193 = &v89.F1
	*result_symbol193 = 3
	v90 = *lexer_addr
	mark_end194 = &v90.F3
	v91 = *mark_end194
	v92 = *lexer_addr
	v91(v92)
	v93 = *lookahead
	cmp195 = v93 == 46
	if cmp195 {
		goto if_then197
	} else {
		goto if_end198
	}

if_then197:
	*state_addr = 12
	goto next_state

if_end198:
	v94 = *lookahead
	cmp199 = 48 <= v94
	if cmp199 {
		goto land_lhs_true201
	} else {
		goto if_end205
	}

land_lhs_true201:
	v95 = *lookahead
	cmp202 = v95 <= 57
	if cmp202 {
		goto if_then204
	} else {
		goto if_end205
	}

if_then204:
	*state_addr = 9
	goto next_state

if_end205:
	v96 = *lookahead
	cmp206 = v96 != 0
	if cmp206 {
		goto land_lhs_true208
	} else {
		goto if_end218
	}

land_lhs_true208:
	v97 = *lookahead
	cmp209 = v97 != 9
	if cmp209 {
		goto land_lhs_true211
	} else {
		goto if_end218
	}

land_lhs_true211:
	v98 = *lookahead
	cmp212 = v98 != 10
	if cmp212 {
		goto land_lhs_true214
	} else {
		goto if_end218
	}

land_lhs_true214:
	v99 = *lookahead
	cmp215 = v99 != 13
	if cmp215 {
		goto if_then217
	} else {
		goto if_end218
	}

if_then217:
	*state_addr = 28
	goto next_state

if_end218:
	v100 = *result
	tobool219 = (v100 & 1) != 0
	*retval = tobool219
	goto _return

sw_bb220:
	*result = 1
	v101 = *lexer_addr
	result_symbol221 = &v101.F1
	*result_symbol221 = 4
	v102 = *lexer_addr
	mark_end222 = &v102.F3
	v103 = *mark_end222
	v104 = *lexer_addr
	v103(v104)
	v105 = *lookahead
	cmp223 = 48 <= v105
	if cmp223 {
		goto land_lhs_true225
	} else {
		goto lor_lhs_false228
	}

land_lhs_true225:
	v106 = *lookahead
	cmp226 = v106 <= 57
	if cmp226 {
		goto if_then240
	} else {
		goto lor_lhs_false228
	}

lor_lhs_false228:
	v107 = *lookahead
	cmp229 = 65 <= v107
	if cmp229 {
		goto land_lhs_true231
	} else {
		goto lor_lhs_false234
	}

land_lhs_true231:
	v108 = *lookahead
	cmp232 = v108 <= 70
	if cmp232 {
		goto if_then240
	} else {
		goto lor_lhs_false234
	}

lor_lhs_false234:
	v109 = *lookahead
	cmp235 = 97 <= v109
	if cmp235 {
		goto land_lhs_true237
	} else {
		goto if_end241
	}

land_lhs_true237:
	v110 = *lookahead
	cmp238 = v110 <= 102
	if cmp238 {
		goto if_then240
	} else {
		goto if_end241
	}

if_then240:
	*state_addr = 10
	goto next_state

if_end241:
	v111 = *lookahead
	cmp242 = v111 != 0
	if cmp242 {
		goto land_lhs_true244
	} else {
		goto if_end254
	}

land_lhs_true244:
	v112 = *lookahead
	cmp245 = v112 != 9
	if cmp245 {
		goto land_lhs_true247
	} else {
		goto if_end254
	}

land_lhs_true247:
	v113 = *lookahead
	cmp248 = v113 != 10
	if cmp248 {
		goto land_lhs_true250
	} else {
		goto if_end254
	}

land_lhs_true250:
	v114 = *lookahead
	cmp251 = v114 != 13
	if cmp251 {
		goto if_then253
	} else {
		goto if_end254
	}

if_then253:
	*state_addr = 28
	goto next_state

if_end254:
	v115 = *result
	tobool255 = (v115 & 1) != 0
	*retval = tobool255
	goto _return

sw_bb256:
	*result = 1
	v116 = *lexer_addr
	result_symbol257 = &v116.F1
	*result_symbol257 = 5
	v117 = *lexer_addr
	mark_end258 = &v117.F3
	v118 = *mark_end258
	v119 = *lexer_addr
	v118(v119)
	v120 = *lookahead
	cmp259 = 48 <= v120
	if cmp259 {
		goto land_lhs_true261
	} else {
		goto if_end265
	}

land_lhs_true261:
	v121 = *lookahead
	cmp262 = v121 <= 57
	if cmp262 {
		goto if_then264
	} else {
		goto if_end265
	}

if_then264:
	*state_addr = 11
	goto next_state

if_end265:
	v122 = *lookahead
	cmp266 = v122 != 0
	if cmp266 {
		goto land_lhs_true268
	} else {
		goto if_end278
	}

land_lhs_true268:
	v123 = *lookahead
	cmp269 = v123 != 9
	if cmp269 {
		goto land_lhs_true271
	} else {
		goto if_end278
	}

land_lhs_true271:
	v124 = *lookahead
	cmp272 = v124 != 10
	if cmp272 {
		goto land_lhs_true274
	} else {
		goto if_end278
	}

land_lhs_true274:
	v125 = *lookahead
	cmp275 = v125 != 13
	if cmp275 {
		goto if_then277
	} else {
		goto if_end278
	}

if_then277:
	*state_addr = 28
	goto next_state

if_end278:
	v126 = *result
	tobool279 = (v126 & 1) != 0
	*retval = tobool279
	goto _return

sw_bb280:
	*result = 1
	v127 = *lexer_addr
	result_symbol281 = &v127.F1
	*result_symbol281 = 6
	v128 = *lexer_addr
	mark_end282 = &v128.F3
	v129 = *mark_end282
	v130 = *lexer_addr
	v129(v130)
	v131 = *lookahead
	cmp283 = 48 <= v131
	if cmp283 {
		goto land_lhs_true285
	} else {
		goto if_end289
	}

land_lhs_true285:
	v132 = *lookahead
	cmp286 = v132 <= 57
	if cmp286 {
		goto if_then288
	} else {
		goto if_end289
	}

if_then288:
	*state_addr = 11
	goto next_state

if_end289:
	v133 = *lookahead
	cmp290 = v133 != 0
	if cmp290 {
		goto land_lhs_true292
	} else {
		goto if_end302
	}

land_lhs_true292:
	v134 = *lookahead
	cmp293 = v134 != 9
	if cmp293 {
		goto land_lhs_true295
	} else {
		goto if_end302
	}

land_lhs_true295:
	v135 = *lookahead
	cmp296 = v135 != 10
	if cmp296 {
		goto land_lhs_true298
	} else {
		goto if_end302
	}

land_lhs_true298:
	v136 = *lookahead
	cmp299 = v136 != 13
	if cmp299 {
		goto if_then301
	} else {
		goto if_end302
	}

if_then301:
	*state_addr = 28
	goto next_state

if_end302:
	v137 = *result
	tobool303 = (v137 & 1) != 0
	*retval = tobool303
	goto _return

sw_bb304:
	*result = 1
	v138 = *lexer_addr
	result_symbol305 = &v138.F1
	*result_symbol305 = 7
	v139 = *lexer_addr
	mark_end306 = &v139.F3
	v140 = *mark_end306
	v141 = *lexer_addr
	v140(v141)
	v142 = *lookahead
	cmp307 = v142 != 0
	if cmp307 {
		goto land_lhs_true309
	} else {
		goto if_end319
	}

land_lhs_true309:
	v143 = *lookahead
	cmp310 = v143 != 9
	if cmp310 {
		goto land_lhs_true312
	} else {
		goto if_end319
	}

land_lhs_true312:
	v144 = *lookahead
	cmp313 = v144 != 10
	if cmp313 {
		goto land_lhs_true315
	} else {
		goto if_end319
	}

land_lhs_true315:
	v145 = *lookahead
	cmp316 = v145 != 13
	if cmp316 {
		goto if_then318
	} else {
		goto if_end319
	}

if_then318:
	*state_addr = 28
	goto next_state

if_end319:
	v146 = *result
	tobool320 = (v146 & 1) != 0
	*retval = tobool320
	goto _return

sw_bb321:
	*result = 1
	v147 = *lexer_addr
	result_symbol322 = &v147.F1
	*result_symbol322 = 8
	v148 = *lexer_addr
	mark_end323 = &v148.F3
	v149 = *mark_end323
	v150 = *lexer_addr
	v149(v150)
	v151 = *lookahead
	cmp324 = v151 != 0
	if cmp324 {
		goto land_lhs_true326
	} else {
		goto if_end336
	}

land_lhs_true326:
	v152 = *lookahead
	cmp327 = v152 != 9
	if cmp327 {
		goto land_lhs_true329
	} else {
		goto if_end336
	}

land_lhs_true329:
	v153 = *lookahead
	cmp330 = v153 != 10
	if cmp330 {
		goto land_lhs_true332
	} else {
		goto if_end336
	}

land_lhs_true332:
	v154 = *lookahead
	cmp333 = v154 != 13
	if cmp333 {
		goto if_then335
	} else {
		goto if_end336
	}

if_then335:
	*state_addr = 28
	goto next_state

if_end336:
	v155 = *result
	tobool337 = (v155 & 1) != 0
	*retval = tobool337
	goto _return

sw_bb338:
	*result = 1
	v156 = *lexer_addr
	result_symbol339 = &v156.F1
	*result_symbol339 = 9
	v157 = *lexer_addr
	mark_end340 = &v157.F3
	v158 = *mark_end340
	v159 = *lexer_addr
	v158(v159)
	v160 = *lookahead
	cmp341 = v160 == 34
	if cmp341 {
		goto if_then343
	} else {
		goto if_end344
	}

if_then343:
	*state_addr = 17
	goto next_state

if_end344:
	v161 = *lookahead
	cmp345 = v161 == 46
	if cmp345 {
		goto if_then347
	} else {
		goto if_end348
	}

if_then347:
	*state_addr = 26
	goto next_state

if_end348:
	v162 = *lookahead
	cmp349 = v162 == 48
	if cmp349 {
		goto if_then351
	} else {
		goto if_end352
	}

if_then351:
	*state_addr = 8
	goto next_state

if_end352:
	v163 = *lookahead
	cmp353 = v163 == 102
	if cmp353 {
		goto if_then355
	} else {
		goto if_end356
	}

if_then355:
	*state_addr = 19
	goto next_state

if_end356:
	v164 = *lookahead
	cmp357 = v164 == 116
	if cmp357 {
		goto if_then359
	} else {
		goto if_end360
	}

if_then359:
	*state_addr = 23
	goto next_state

if_end360:
	v165 = *lookahead
	cmp361 = v165 == 11
	if cmp361 {
		goto if_then369
	} else {
		goto lor_lhs_false363
	}

lor_lhs_false363:
	v166 = *lookahead
	cmp364 = v166 == 12
	if cmp364 {
		goto if_then369
	} else {
		goto lor_lhs_false366
	}

lor_lhs_false366:
	v167 = *lookahead
	cmp367 = v167 == 32
	if cmp367 {
		goto if_then369
	} else {
		goto if_end370
	}

if_then369:
	*state_addr = 15
	goto next_state

if_end370:
	v168 = *lookahead
	cmp371 = 49 <= v168
	if cmp371 {
		goto land_lhs_true373
	} else {
		goto if_end377
	}

land_lhs_true373:
	v169 = *lookahead
	cmp374 = v169 <= 57
	if cmp374 {
		goto if_then376
	} else {
		goto if_end377
	}

if_then376:
	*state_addr = 9
	goto next_state

if_end377:
	v170 = *lookahead
	cmp378 = v170 != 0
	if cmp378 {
		goto land_lhs_true380
	} else {
		goto if_end387
	}

land_lhs_true380:
	v171 = *lookahead
	cmp381 = v171 < 9
	if cmp381 {
		goto if_then386
	} else {
		goto lor_lhs_false383
	}

lor_lhs_false383:
	v172 = *lookahead
	cmp384 = 13 < v172
	if cmp384 {
		goto if_then386
	} else {
		goto if_end387
	}

if_then386:
	*state_addr = 28
	goto next_state

if_end387:
	v173 = *result
	tobool388 = (v173 & 1) != 0
	*retval = tobool388
	goto _return

sw_bb389:
	*result = 1
	v174 = *lexer_addr
	result_symbol390 = &v174.F1
	*result_symbol390 = 9
	v175 = *lexer_addr
	mark_end391 = &v175.F3
	v176 = *mark_end391
	v177 = *lexer_addr
	v176(v177)
	v178 = *lookahead
	cmp392 = v178 == 34
	if cmp392 {
		goto if_then394
	} else {
		goto if_end395
	}

if_then394:
	*state_addr = 17
	goto next_state

if_end395:
	v179 = *lookahead
	cmp396 = v179 != 0
	if cmp396 {
		goto land_lhs_true398
	} else {
		goto if_end408
	}

land_lhs_true398:
	v180 = *lookahead
	cmp399 = v180 != 9
	if cmp399 {
		goto land_lhs_true401
	} else {
		goto if_end408
	}

land_lhs_true401:
	v181 = *lookahead
	cmp402 = v181 != 10
	if cmp402 {
		goto land_lhs_true404
	} else {
		goto if_end408
	}

land_lhs_true404:
	v182 = *lookahead
	cmp405 = v182 != 13
	if cmp405 {
		goto if_then407
	} else {
		goto if_end408
	}

if_then407:
	*state_addr = 28
	goto next_state

if_end408:
	v183 = *result
	tobool409 = (v183 & 1) != 0
	*retval = tobool409
	goto _return

sw_bb410:
	*result = 1
	v184 = *lexer_addr
	result_symbol411 = &v184.F1
	*result_symbol411 = 9
	v185 = *lexer_addr
	mark_end412 = &v185.F3
	v186 = *mark_end412
	v187 = *lexer_addr
	v186(v187)
	v188 = *lookahead
	cmp413 = v188 == 34
	if cmp413 {
		goto if_then415
	} else {
		goto if_end416
	}

if_then415:
	*state_addr = 16
	goto next_state

if_end416:
	v189 = *lookahead
	cmp417 = v189 == 9
	if cmp417 {
		goto if_then425
	} else {
		goto lor_lhs_false419
	}

lor_lhs_false419:
	v190 = *lookahead
	cmp420 = v190 == 10
	if cmp420 {
		goto if_then425
	} else {
		goto lor_lhs_false422
	}

lor_lhs_false422:
	v191 = *lookahead
	cmp423 = v191 == 13
	if cmp423 {
		goto if_then425
	} else {
		goto if_end426
	}

if_then425:
	*state_addr = 1
	goto next_state

if_end426:
	v192 = *lookahead
	cmp427 = v192 != 0
	if cmp427 {
		goto if_then429
	} else {
		goto if_end430
	}

if_then429:
	*state_addr = 17
	goto next_state

if_end430:
	v193 = *result
	tobool431 = (v193 & 1) != 0
	*retval = tobool431
	goto _return

sw_bb432:
	*result = 1
	v194 = *lexer_addr
	result_symbol433 = &v194.F1
	*result_symbol433 = 9
	v195 = *lexer_addr
	mark_end434 = &v195.F3
	v196 = *mark_end434
	v197 = *lexer_addr
	v196(v197)
	v198 = *lookahead
	cmp435 = v198 == 34
	if cmp435 {
		goto if_then437
	} else {
		goto if_end438
	}

if_then437:
	*state_addr = 1
	goto next_state

if_end438:
	v199 = *result
	tobool439 = (v199 & 1) != 0
	*retval = tobool439
	goto _return

sw_bb440:
	*result = 1
	v200 = *lexer_addr
	result_symbol441 = &v200.F1
	*result_symbol441 = 9
	v201 = *lexer_addr
	mark_end442 = &v201.F3
	v202 = *mark_end442
	v203 = *lexer_addr
	v202(v203)
	v204 = *lookahead
	cmp443 = v204 == 97
	if cmp443 {
		goto if_then445
	} else {
		goto if_end446
	}

if_then445:
	*state_addr = 22
	goto next_state

if_end446:
	v205 = *lookahead
	cmp447 = v205 != 0
	if cmp447 {
		goto land_lhs_true449
	} else {
		goto if_end459
	}

land_lhs_true449:
	v206 = *lookahead
	cmp450 = v206 != 9
	if cmp450 {
		goto land_lhs_true452
	} else {
		goto if_end459
	}

land_lhs_true452:
	v207 = *lookahead
	cmp453 = v207 != 10
	if cmp453 {
		goto land_lhs_true455
	} else {
		goto if_end459
	}

land_lhs_true455:
	v208 = *lookahead
	cmp456 = v208 != 13
	if cmp456 {
		goto if_then458
	} else {
		goto if_end459
	}

if_then458:
	*state_addr = 28
	goto next_state

if_end459:
	v209 = *result
	tobool460 = (v209 & 1) != 0
	*retval = tobool460
	goto _return

sw_bb461:
	*result = 1
	v210 = *lexer_addr
	result_symbol462 = &v210.F1
	*result_symbol462 = 9
	v211 = *lexer_addr
	mark_end463 = &v211.F3
	v212 = *mark_end463
	v213 = *lexer_addr
	v212(v213)
	v214 = *lookahead
	cmp464 = v214 == 101
	if cmp464 {
		goto if_then466
	} else {
		goto if_end467
	}

if_then466:
	*state_addr = 13
	goto next_state

if_end467:
	v215 = *lookahead
	cmp468 = v215 != 0
	if cmp468 {
		goto land_lhs_true470
	} else {
		goto if_end480
	}

land_lhs_true470:
	v216 = *lookahead
	cmp471 = v216 != 9
	if cmp471 {
		goto land_lhs_true473
	} else {
		goto if_end480
	}

land_lhs_true473:
	v217 = *lookahead
	cmp474 = v217 != 10
	if cmp474 {
		goto land_lhs_true476
	} else {
		goto if_end480
	}

land_lhs_true476:
	v218 = *lookahead
	cmp477 = v218 != 13
	if cmp477 {
		goto if_then479
	} else {
		goto if_end480
	}

if_then479:
	*state_addr = 28
	goto next_state

if_end480:
	v219 = *result
	tobool481 = (v219 & 1) != 0
	*retval = tobool481
	goto _return

sw_bb482:
	*result = 1
	v220 = *lexer_addr
	result_symbol483 = &v220.F1
	*result_symbol483 = 9
	v221 = *lexer_addr
	mark_end484 = &v221.F3
	v222 = *mark_end484
	v223 = *lexer_addr
	v222(v223)
	v224 = *lookahead
	cmp485 = v224 == 101
	if cmp485 {
		goto if_then487
	} else {
		goto if_end488
	}

if_then487:
	*state_addr = 14
	goto next_state

if_end488:
	v225 = *lookahead
	cmp489 = v225 != 0
	if cmp489 {
		goto land_lhs_true491
	} else {
		goto if_end501
	}

land_lhs_true491:
	v226 = *lookahead
	cmp492 = v226 != 9
	if cmp492 {
		goto land_lhs_true494
	} else {
		goto if_end501
	}

land_lhs_true494:
	v227 = *lookahead
	cmp495 = v227 != 10
	if cmp495 {
		goto land_lhs_true497
	} else {
		goto if_end501
	}

land_lhs_true497:
	v228 = *lookahead
	cmp498 = v228 != 13
	if cmp498 {
		goto if_then500
	} else {
		goto if_end501
	}

if_then500:
	*state_addr = 28
	goto next_state

if_end501:
	v229 = *result
	tobool502 = (v229 & 1) != 0
	*retval = tobool502
	goto _return

sw_bb503:
	*result = 1
	v230 = *lexer_addr
	result_symbol504 = &v230.F1
	*result_symbol504 = 9
	v231 = *lexer_addr
	mark_end505 = &v231.F3
	v232 = *mark_end505
	v233 = *lexer_addr
	v232(v233)
	v234 = *lookahead
	cmp506 = v234 == 108
	if cmp506 {
		goto if_then508
	} else {
		goto if_end509
	}

if_then508:
	*state_addr = 24
	goto next_state

if_end509:
	v235 = *lookahead
	cmp510 = v235 != 0
	if cmp510 {
		goto land_lhs_true512
	} else {
		goto if_end522
	}

land_lhs_true512:
	v236 = *lookahead
	cmp513 = v236 != 9
	if cmp513 {
		goto land_lhs_true515
	} else {
		goto if_end522
	}

land_lhs_true515:
	v237 = *lookahead
	cmp516 = v237 != 10
	if cmp516 {
		goto land_lhs_true518
	} else {
		goto if_end522
	}

land_lhs_true518:
	v238 = *lookahead
	cmp519 = v238 != 13
	if cmp519 {
		goto if_then521
	} else {
		goto if_end522
	}

if_then521:
	*state_addr = 28
	goto next_state

if_end522:
	v239 = *result
	tobool523 = (v239 & 1) != 0
	*retval = tobool523
	goto _return

sw_bb524:
	*result = 1
	v240 = *lexer_addr
	result_symbol525 = &v240.F1
	*result_symbol525 = 9
	v241 = *lexer_addr
	mark_end526 = &v241.F3
	v242 = *mark_end526
	v243 = *lexer_addr
	v242(v243)
	v244 = *lookahead
	cmp527 = v244 == 114
	if cmp527 {
		goto if_then529
	} else {
		goto if_end530
	}

if_then529:
	*state_addr = 25
	goto next_state

if_end530:
	v245 = *lookahead
	cmp531 = v245 != 0
	if cmp531 {
		goto land_lhs_true533
	} else {
		goto if_end543
	}

land_lhs_true533:
	v246 = *lookahead
	cmp534 = v246 != 9
	if cmp534 {
		goto land_lhs_true536
	} else {
		goto if_end543
	}

land_lhs_true536:
	v247 = *lookahead
	cmp537 = v247 != 10
	if cmp537 {
		goto land_lhs_true539
	} else {
		goto if_end543
	}

land_lhs_true539:
	v248 = *lookahead
	cmp540 = v248 != 13
	if cmp540 {
		goto if_then542
	} else {
		goto if_end543
	}

if_then542:
	*state_addr = 28
	goto next_state

if_end543:
	v249 = *result
	tobool544 = (v249 & 1) != 0
	*retval = tobool544
	goto _return

sw_bb545:
	*result = 1
	v250 = *lexer_addr
	result_symbol546 = &v250.F1
	*result_symbol546 = 9
	v251 = *lexer_addr
	mark_end547 = &v251.F3
	v252 = *mark_end547
	v253 = *lexer_addr
	v252(v253)
	v254 = *lookahead
	cmp548 = v254 == 115
	if cmp548 {
		goto if_then550
	} else {
		goto if_end551
	}

if_then550:
	*state_addr = 21
	goto next_state

if_end551:
	v255 = *lookahead
	cmp552 = v255 != 0
	if cmp552 {
		goto land_lhs_true554
	} else {
		goto if_end564
	}

land_lhs_true554:
	v256 = *lookahead
	cmp555 = v256 != 9
	if cmp555 {
		goto land_lhs_true557
	} else {
		goto if_end564
	}

land_lhs_true557:
	v257 = *lookahead
	cmp558 = v257 != 10
	if cmp558 {
		goto land_lhs_true560
	} else {
		goto if_end564
	}

land_lhs_true560:
	v258 = *lookahead
	cmp561 = v258 != 13
	if cmp561 {
		goto if_then563
	} else {
		goto if_end564
	}

if_then563:
	*state_addr = 28
	goto next_state

if_end564:
	v259 = *result
	tobool565 = (v259 & 1) != 0
	*retval = tobool565
	goto _return

sw_bb566:
	*result = 1
	v260 = *lexer_addr
	result_symbol567 = &v260.F1
	*result_symbol567 = 9
	v261 = *lexer_addr
	mark_end568 = &v261.F3
	v262 = *mark_end568
	v263 = *lexer_addr
	v262(v263)
	v264 = *lookahead
	cmp569 = v264 == 117
	if cmp569 {
		goto if_then571
	} else {
		goto if_end572
	}

if_then571:
	*state_addr = 20
	goto next_state

if_end572:
	v265 = *lookahead
	cmp573 = v265 != 0
	if cmp573 {
		goto land_lhs_true575
	} else {
		goto if_end585
	}

land_lhs_true575:
	v266 = *lookahead
	cmp576 = v266 != 9
	if cmp576 {
		goto land_lhs_true578
	} else {
		goto if_end585
	}

land_lhs_true578:
	v267 = *lookahead
	cmp579 = v267 != 10
	if cmp579 {
		goto land_lhs_true581
	} else {
		goto if_end585
	}

land_lhs_true581:
	v268 = *lookahead
	cmp582 = v268 != 13
	if cmp582 {
		goto if_then584
	} else {
		goto if_end585
	}

if_then584:
	*state_addr = 28
	goto next_state

if_end585:
	v269 = *result
	tobool586 = (v269 & 1) != 0
	*retval = tobool586
	goto _return

sw_bb587:
	*result = 1
	v270 = *lexer_addr
	result_symbol588 = &v270.F1
	*result_symbol588 = 9
	v271 = *lexer_addr
	mark_end589 = &v271.F3
	v272 = *mark_end589
	v273 = *lexer_addr
	v272(v273)
	v274 = *lookahead
	cmp590 = 48 <= v274
	if cmp590 {
		goto land_lhs_true592
	} else {
		goto if_end596
	}

land_lhs_true592:
	v275 = *lookahead
	cmp593 = v275 <= 57
	if cmp593 {
		goto if_then595
	} else {
		goto if_end596
	}

if_then595:
	*state_addr = 11
	goto next_state

if_end596:
	v276 = *lookahead
	cmp597 = v276 != 0
	if cmp597 {
		goto land_lhs_true599
	} else {
		goto if_end609
	}

land_lhs_true599:
	v277 = *lookahead
	cmp600 = v277 != 9
	if cmp600 {
		goto land_lhs_true602
	} else {
		goto if_end609
	}

land_lhs_true602:
	v278 = *lookahead
	cmp603 = v278 != 10
	if cmp603 {
		goto land_lhs_true605
	} else {
		goto if_end609
	}

land_lhs_true605:
	v279 = *lookahead
	cmp606 = v279 != 13
	if cmp606 {
		goto if_then608
	} else {
		goto if_end609
	}

if_then608:
	*state_addr = 28
	goto next_state

if_end609:
	v280 = *result
	tobool610 = (v280 & 1) != 0
	*retval = tobool610
	goto _return

sw_bb611:
	*result = 1
	v281 = *lexer_addr
	result_symbol612 = &v281.F1
	*result_symbol612 = 9
	v282 = *lexer_addr
	mark_end613 = &v282.F3
	v283 = *mark_end613
	v284 = *lexer_addr
	v283(v284)
	v285 = *lookahead
	cmp614 = 48 <= v285
	if cmp614 {
		goto land_lhs_true616
	} else {
		goto lor_lhs_false619
	}

land_lhs_true616:
	v286 = *lookahead
	cmp617 = v286 <= 57
	if cmp617 {
		goto if_then631
	} else {
		goto lor_lhs_false619
	}

lor_lhs_false619:
	v287 = *lookahead
	cmp620 = 65 <= v287
	if cmp620 {
		goto land_lhs_true622
	} else {
		goto lor_lhs_false625
	}

land_lhs_true622:
	v288 = *lookahead
	cmp623 = v288 <= 70
	if cmp623 {
		goto if_then631
	} else {
		goto lor_lhs_false625
	}

lor_lhs_false625:
	v289 = *lookahead
	cmp626 = 97 <= v289
	if cmp626 {
		goto land_lhs_true628
	} else {
		goto if_end632
	}

land_lhs_true628:
	v290 = *lookahead
	cmp629 = v290 <= 102
	if cmp629 {
		goto if_then631
	} else {
		goto if_end632
	}

if_then631:
	*state_addr = 10
	goto next_state

if_end632:
	v291 = *lookahead
	cmp633 = v291 != 0
	if cmp633 {
		goto land_lhs_true635
	} else {
		goto if_end645
	}

land_lhs_true635:
	v292 = *lookahead
	cmp636 = v292 != 9
	if cmp636 {
		goto land_lhs_true638
	} else {
		goto if_end645
	}

land_lhs_true638:
	v293 = *lookahead
	cmp639 = v293 != 10
	if cmp639 {
		goto land_lhs_true641
	} else {
		goto if_end645
	}

land_lhs_true641:
	v294 = *lookahead
	cmp642 = v294 != 13
	if cmp642 {
		goto if_then644
	} else {
		goto if_end645
	}

if_then644:
	*state_addr = 28
	goto next_state

if_end645:
	v295 = *result
	tobool646 = (v295 & 1) != 0
	*retval = tobool646
	goto _return

sw_bb647:
	*result = 1
	v296 = *lexer_addr
	result_symbol648 = &v296.F1
	*result_symbol648 = 9
	v297 = *lexer_addr
	mark_end649 = &v297.F3
	v298 = *mark_end649
	v299 = *lexer_addr
	v298(v299)
	v300 = *lookahead
	cmp650 = v300 != 0
	if cmp650 {
		goto land_lhs_true652
	} else {
		goto if_end662
	}

land_lhs_true652:
	v301 = *lookahead
	cmp653 = v301 != 9
	if cmp653 {
		goto land_lhs_true655
	} else {
		goto if_end662
	}

land_lhs_true655:
	v302 = *lookahead
	cmp656 = v302 != 10
	if cmp656 {
		goto land_lhs_true658
	} else {
		goto if_end662
	}

land_lhs_true658:
	v303 = *lookahead
	cmp659 = v303 != 13
	if cmp659 {
		goto if_then661
	} else {
		goto if_end662
	}

if_then661:
	*state_addr = 28
	goto next_state

if_end662:
	v304 = *result
	tobool663 = (v304 & 1) != 0
	*retval = tobool663
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v305 = *retval
	return v305
}

