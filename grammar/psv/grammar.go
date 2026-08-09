package grammar_psv

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

var tree_sitter_psv_language TSLanguage = TSLanguage{15, 18, 0, 10, 0, 27, 6, 1, 0, 2, &(*[6][18]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[79]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{1, 2, 0}}

var ts_small_parse_table [173]int16 = [173]int16{
	2, 15, 1, 0, 37, 7, 3, 4, 5, 6, 7, 8, 9, 3, 41, 1,
	2, 8, 1, 17, 39, 2, 0, 1, 3, 41, 1, 2, 9, 1, 17, 43,
	2, 0, 1, 3, 47, 1, 2, 9, 1, 17, 45, 2, 0, 1, 2, 52,
	1, 2, 50, 2, 0, 1, 2, 56, 1, 2, 54, 2, 0, 1, 2, 60,
	1, 2, 58, 2, 0, 1, 2, 64, 1, 2, 62, 2, 0, 1, 3, 43,
	1, 1, 66, 1, 2, 17, 1, 17, 2, 68, 1, 2, 45, 2, 0, 1,
	3, 39, 1, 1, 66, 1, 2, 14, 1, 17, 3, 45, 1, 1, 70, 1,
	2, 17, 1, 17, 2, 13, 1, 0, 73, 1, 1, 2, 50, 1, 1, 52,
	1, 2, 2, 73, 1, 1, 75, 1, 0, 2, 62, 1, 1, 64, 1, 2,
	2, 54, 1, 1, 56, 1, 2, 2, 58, 1, 1, 60, 1, 2, 2, 45,
	1, 1, 68, 1, 2, 1, 77, 1, 0, 1, 73, 1, 1,
}

var ts_small_parse_table_map [21]int32 = [21]int32{
	0, 13, 24, 35, 46, 54, 62, 70, 78, 88, 96, 106, 116, 123, 130, 137,
	144, 151, 158, 165, 169,
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
	TSLexerMode{}, TSLexerMode{43, 0, 0}, TSLexerMode{43, 0, 0}, TSLexerMode{43, 0, 0}, TSLexerMode{43, 0, 0}, TSLexerMode{43, 0, 0}, TSLexerMode{43, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0},
	TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{11, 0, 0}, TSLexerMode{}, TSLexerMode{11, 0, 0},
}

var ts_primary_state_ids [27]int16 = [27]int16{
	0, 1, 2, 3, 4, 4, 6, 7, 8, 9, 10, 11, 12, 13, 8, 15,
	7, 9, 18, 10, 20, 13, 11, 12, 15, 25, 26,
}

var _str [4]byte = [4]byte{112, 115, 118, 0}

var ts_parse_table struct {
	F0 struct {
	F0 [9]int16
	F1 [9]int16
}
	F1 [18]int16
	F2 [18]int16
	F3 [18]int16
	F4 [18]int16
	F5 [18]int16
} = struct {
	F0 struct {
	F0 [9]int16
	F1 [9]int16
}
	F1 [18]int16
	F2 [18]int16
	F3 [18]int16
	F4 [18]int16
	F5 [18]int16
}{struct {
	F0 [9]int16
	F1 [9]int16
}{[9]int16{1, 0, 1, 1, 1, 1, 1, 1, 1}, [9]int16{}}, [18]int16{
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
	F42 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F69 TSParseActionEntry
	F70 struct {
	F0 anon_1
	F1 [6]byte
}
	F71 TSParseActionEntry
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
	F74 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F42 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F69 TSParseActionEntry
	F70 struct {
	F0 anon_1
	F1 [6]byte
}
	F71 TSParseActionEntry
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
	F74 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 17, 0, 0}}}, struct {
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

var _str_5 [2]byte = [2]byte{124, 0}

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

var ts_lex_map [18]int16 = [18]int16{
	34, 31, 46, 40, 48, 17, 102, 33, 116, 37, 9, 29, 11, 29, 12, 29,
	32, 29,
}

var ts_lex_map_21 [18]int16 = [18]int16{
	34, 31, 46, 40, 48, 17, 102, 33, 116, 37, 9, 29, 11, 29, 12, 29,
	32, 29,
}

func tree_sitter_psv() *TSLanguage {
	return &tree_sitter_psv_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v57, v58, v60, v62, v63, v65, v69, v70, v72, v74, v75, v77, v84, v85, v87, v92, v93, v95, v106, v107, v109, v118, v119, v121, v129, v130, v132, v144, v145, v147, v151, v152, v154, v162, v163, v165, v169, v170, v172, v180, v181, v183, v185, v186, v188, v194, v195, v197, v199, v200, v202, v208, v209, v211, v226, v227, v229, v236, v237, v239, v246, v247, v249, v252, v253, v255, v262, v263, v265, v272, v273, v275, v282, v283, v285, v292, v293, v295, v302, v303, v305, v312, v313, v315, v322, v323, v325, v333, v334, v336, v348, v349, v351, v357, v358, v360 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end152, mark_end164, mark_end168, mark_end190, mark_end205, mark_end240, mark_end268, mark_end291, mark_end327, mark_end338, mark_end362, mark_end373, mark_end397, mark_end401, mark_end418, mark_end422, mark_end439, mark_end474, mark_end495, mark_end517, mark_end525, mark_end546, mark_end567, mark_end588, mark_end609, mark_end630, mark_end651, mark_end672, mark_end696, mark_end732, mark_end749 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, result_symbol, result_symbol151, result_symbol163, result_symbol167, result_symbol189, result_symbol204, result_symbol239, result_symbol267, result_symbol290, result_symbol326, result_symbol337, result_symbol361, result_symbol372, result_symbol396, result_symbol400, result_symbol417, result_symbol421, result_symbol438, arrayidx, arrayidx448, result_symbol473, result_symbol494, result_symbol516, result_symbol524, result_symbol545, result_symbol566, result_symbol587, result_symbol608, result_symbol629, result_symbol650, result_symbol671, result_symbol695, result_symbol731, result_symbol748, arrayidx760, arrayidx767 *int16
	var lookahead, i, i753, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp25, cmp27, cmp31, cmp34, tobool38, cmp40, cmp44, tobool48, cmp50, tobool54, cmp56, tobool60, cmp62, tobool66, cmp68, tobool72, cmp74, tobool78, cmp80, tobool84, cmp86, tobool90, cmp92, cmp95, tobool99, cmp101, cmp104, cmp107, cmp110, cmp113, cmp116, tobool120, tobool122, cmp125, cmp129, cmp133, cmp137, cmp140, cmp143, tobool147, tobool149, cmp153, cmp157, tobool161, tobool165, cmp169, cmp173, cmp176, cmp180, cmp183, tobool187, cmp191, cmp195, cmp198, tobool202, cmp206, cmp210, cmp213, cmp217, cmp220, cmp224, cmp227, cmp230, cmp233, tobool237, cmp241, cmp245, cmp248, cmp252, cmp255, cmp258, cmp261, tobool265, cmp269, cmp272, cmp275, cmp278, cmp281, cmp284, tobool288, cmp292, cmp295, cmp298, cmp301, cmp304, cmp307, cmp311, cmp314, cmp317, cmp320, tobool324, cmp328, cmp331, tobool335, cmp339, cmp342, cmp346, cmp349, cmp352, cmp355, tobool359, cmp363, cmp366, tobool370, cmp374, cmp377, cmp381, cmp384, cmp387, cmp390, tobool394, tobool398, cmp402, cmp405, cmp408, cmp411, tobool415, tobool419, cmp423, cmp426, cmp429, cmp432, tobool436, cmp441, cmp444, cmp451, cmp454, cmp458, cmp461, cmp464, cmp467, tobool471, cmp475, cmp479, cmp482, cmp485, cmp488, tobool492, cmp496, cmp500, cmp503, cmp506, cmp510, tobool514, cmp518, tobool522, cmp526, cmp530, cmp533, cmp536, cmp539, tobool543, cmp547, cmp551, cmp554, cmp557, cmp560, tobool564, cmp568, cmp572, cmp575, cmp578, cmp581, tobool585, cmp589, cmp593, cmp596, cmp599, cmp602, tobool606, cmp610, cmp614, cmp617, cmp620, cmp623, tobool627, cmp631, cmp635, cmp638, cmp641, cmp644, tobool648, cmp652, cmp656, cmp659, cmp662, cmp665, tobool669, cmp673, cmp676, cmp680, cmp683, cmp686, cmp689, tobool693, cmp697, cmp700, cmp703, cmp706, cmp709, cmp712, cmp716, cmp719, cmp722, cmp725, tobool729, cmp733, cmp736, cmp739, cmp742, tobool746, tobool750, cmp756, cmp762, cmp772, cmp775, cmp779, cmp782, cmp785, cmp788, tobool792, v376 bool
	var v3, frombool, v10, v21, v24, v26, v28, v30, v32, v34, v36, v38, v41, v48, v49, v56, v61, v68, v73, v83, v91, v105, v117, v128, v143, v150, v161, v168, v179, v184, v193, v198, v207, v225, v235, v245, v251, v261, v271, v281, v291, v301, v311, v321, v332, v347, v356, v361, v375 byte
	var v59, v64, v71, v76, v86, v94, v108, v120, v131, v146, v153, v164, v171, v182, v187, v196, v201, v210, v228, v238, v248, v254, v264, v274, v284, v294, v304, v314, v324, v335, v350, v359 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9, v214, v217, v364, v367 int16
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v22, v23, v25, v27, v29, v31, v33, v35, v37, v39, v40, v42, v43, v44, v45, v46, v47, v50, v51, v52, v53, v54, v55, v66, v67, v78, v79, v80, v81, v82, v88, v89, v90, v96, v97, v98, v99, v100, v101, v102, v103, v104, v110, v111, v112, v113, v114, v115, v116, v122, v123, v124, v125, v126, v127, v133, v134, v135, v136, v137, v138, v139, v140, v141, v142, v148, v149, v155, v156, v157, v158, v159, v160, v166, v167, v173, v174, v175, v176, v177, v178, v189, v190, v191, v192, v203, v204, v205, v206, v212, v213, conv443, v215, v216, add, v218, add450, v219, v220, v221, v222, v223, v224, v230, v231, v232, v233, v234, v240, v241, v242, v243, v244, v250, v256, v257, v258, v259, v260, v266, v267, v268, v269, v270, v276, v277, v278, v279, v280, v286, v287, v288, v289, v290, v296, v297, v298, v299, v300, v306, v307, v308, v309, v310, v316, v317, v318, v319, v320, v326, v327, v328, v329, v330, v331, v337, v338, v339, v340, v341, v342, v343, v344, v345, v346, v352, v353, v354, v355, v362, v363, conv761, v365, v366, add765, v368, add770, v369, v370, v371, v372, v373, v374 int32
	var conv440, idxprom, idxprom447, conv755, idxprom759, idxprom766 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i753, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp25, v18, cmp27, v19, cmp31, v20, cmp34, v21, tobool38, v22, cmp40, v23, cmp44, v24, tobool48, v25, cmp50, v26, tobool54, v27, cmp56, v28, tobool60, v29, cmp62, v30, tobool66, v31, cmp68, v32, tobool72, v33, cmp74, v34, tobool78, v35, cmp80, v36, tobool84, v37, cmp86, v38, tobool90, v39, cmp92, v40, cmp95, v41, tobool99, v42, cmp101, v43, cmp104, v44, cmp107, v45, cmp110, v46, cmp113, v47, cmp116, v48, tobool120, v49, tobool122, v50, cmp125, v51, cmp129, v52, cmp133, v53, cmp137, v54, cmp140, v55, cmp143, v56, tobool147, v57, result_symbol, v58, mark_end, v59, v60, v61, tobool149, v62, result_symbol151, v63, mark_end152, v64, v65, v66, cmp153, v67, cmp157, v68, tobool161, v69, result_symbol163, v70, mark_end164, v71, v72, v73, tobool165, v74, result_symbol167, v75, mark_end168, v76, v77, v78, cmp169, v79, cmp173, v80, cmp176, v81, cmp180, v82, cmp183, v83, tobool187, v84, result_symbol189, v85, mark_end190, v86, v87, v88, cmp191, v89, cmp195, v90, cmp198, v91, tobool202, v92, result_symbol204, v93, mark_end205, v94, v95, v96, cmp206, v97, cmp210, v98, cmp213, v99, cmp217, v100, cmp220, v101, cmp224, v102, cmp227, v103, cmp230, v104, cmp233, v105, tobool237, v106, result_symbol239, v107, mark_end240, v108, v109, v110, cmp241, v111, cmp245, v112, cmp248, v113, cmp252, v114, cmp255, v115, cmp258, v116, cmp261, v117, tobool265, v118, result_symbol267, v119, mark_end268, v120, v121, v122, cmp269, v123, cmp272, v124, cmp275, v125, cmp278, v126, cmp281, v127, cmp284, v128, tobool288, v129, result_symbol290, v130, mark_end291, v131, v132, v133, cmp292, v134, cmp295, v135, cmp298, v136, cmp301, v137, cmp304, v138, cmp307, v139, cmp311, v140, cmp314, v141, cmp317, v142, cmp320, v143, tobool324, v144, result_symbol326, v145, mark_end327, v146, v147, v148, cmp328, v149, cmp331, v150, tobool335, v151, result_symbol337, v152, mark_end338, v153, v154, v155, cmp339, v156, cmp342, v157, cmp346, v158, cmp349, v159, cmp352, v160, cmp355, v161, tobool359, v162, result_symbol361, v163, mark_end362, v164, v165, v166, cmp363, v167, cmp366, v168, tobool370, v169, result_symbol372, v170, mark_end373, v171, v172, v173, cmp374, v174, cmp377, v175, cmp381, v176, cmp384, v177, cmp387, v178, cmp390, v179, tobool394, v180, result_symbol396, v181, mark_end397, v182, v183, v184, tobool398, v185, result_symbol400, v186, mark_end401, v187, v188, v189, cmp402, v190, cmp405, v191, cmp408, v192, cmp411, v193, tobool415, v194, result_symbol417, v195, mark_end418, v196, v197, v198, tobool419, v199, result_symbol421, v200, mark_end422, v201, v202, v203, cmp423, v204, cmp426, v205, cmp429, v206, cmp432, v207, tobool436, v208, result_symbol438, v209, mark_end439, v210, v211, v212, conv440, cmp441, v213, idxprom, arrayidx, v214, conv443, v215, cmp444, v216, add, idxprom447, arrayidx448, v217, v218, add450, v219, cmp451, v220, cmp454, v221, cmp458, v222, cmp461, v223, cmp464, v224, cmp467, v225, tobool471, v226, result_symbol473, v227, mark_end474, v228, v229, v230, cmp475, v231, cmp479, v232, cmp482, v233, cmp485, v234, cmp488, v235, tobool492, v236, result_symbol494, v237, mark_end495, v238, v239, v240, cmp496, v241, cmp500, v242, cmp503, v243, cmp506, v244, cmp510, v245, tobool514, v246, result_symbol516, v247, mark_end517, v248, v249, v250, cmp518, v251, tobool522, v252, result_symbol524, v253, mark_end525, v254, v255, v256, cmp526, v257, cmp530, v258, cmp533, v259, cmp536, v260, cmp539, v261, tobool543, v262, result_symbol545, v263, mark_end546, v264, v265, v266, cmp547, v267, cmp551, v268, cmp554, v269, cmp557, v270, cmp560, v271, tobool564, v272, result_symbol566, v273, mark_end567, v274, v275, v276, cmp568, v277, cmp572, v278, cmp575, v279, cmp578, v280, cmp581, v281, tobool585, v282, result_symbol587, v283, mark_end588, v284, v285, v286, cmp589, v287, cmp593, v288, cmp596, v289, cmp599, v290, cmp602, v291, tobool606, v292, result_symbol608, v293, mark_end609, v294, v295, v296, cmp610, v297, cmp614, v298, cmp617, v299, cmp620, v300, cmp623, v301, tobool627, v302, result_symbol629, v303, mark_end630, v304, v305, v306, cmp631, v307, cmp635, v308, cmp638, v309, cmp641, v310, cmp644, v311, tobool648, v312, result_symbol650, v313, mark_end651, v314, v315, v316, cmp652, v317, cmp656, v318, cmp659, v319, cmp662, v320, cmp665, v321, tobool669, v322, result_symbol671, v323, mark_end672, v324, v325, v326, cmp673, v327, cmp676, v328, cmp680, v329, cmp683, v330, cmp686, v331, cmp689, v332, tobool693, v333, result_symbol695, v334, mark_end696, v335, v336, v337, cmp697, v338, cmp700, v339, cmp703, v340, cmp706, v341, cmp709, v342, cmp712, v343, cmp716, v344, cmp719, v345, cmp722, v346, cmp725, v347, tobool729, v348, result_symbol731, v349, mark_end732, v350, v351, v352, cmp733, v353, cmp736, v354, cmp739, v355, cmp742, v356, tobool746, v357, result_symbol748, v358, mark_end749, v359, v360, v361, tobool750, v362, conv755, cmp756, v363, idxprom759, arrayidx760, v364, conv761, v365, cmp762, v366, add765, idxprom766, arrayidx767, v367, v368, add770, v369, cmp772, v370, cmp775, v371, cmp779, v372, cmp782, v373, cmp785, v374, cmp788, v375, tobool792, v376

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
	i = new(int32)
	i753 = new(int32)
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
		goto sw_bb39
	case 2:
		goto sw_bb49
	case 3:
		goto sw_bb55
	case 4:
		goto sw_bb61
	case 5:
		goto sw_bb67
	case 6:
		goto sw_bb73
	case 7:
		goto sw_bb79
	case 8:
		goto sw_bb85
	case 9:
		goto sw_bb91
	case 10:
		goto sw_bb100
	case 11:
		goto sw_bb121
	case 12:
		goto sw_bb148
	case 13:
		goto sw_bb150
	case 14:
		goto sw_bb162
	case 15:
		goto sw_bb166
	case 16:
		goto sw_bb188
	case 17:
		goto sw_bb203
	case 18:
		goto sw_bb238
	case 19:
		goto sw_bb266
	case 20:
		goto sw_bb289
	case 21:
		goto sw_bb325
	case 22:
		goto sw_bb336
	case 23:
		goto sw_bb360
	case 24:
		goto sw_bb371
	case 25:
		goto sw_bb395
	case 26:
		goto sw_bb399
	case 27:
		goto sw_bb416
	case 28:
		goto sw_bb420
	case 29:
		goto sw_bb437
	case 30:
		goto sw_bb472
	case 31:
		goto sw_bb493
	case 32:
		goto sw_bb515
	case 33:
		goto sw_bb523
	case 34:
		goto sw_bb544
	case 35:
		goto sw_bb565
	case 36:
		goto sw_bb586
	case 37:
		goto sw_bb607
	case 38:
		goto sw_bb628
	case 39:
		goto sw_bb649
	case 40:
		goto sw_bb670
	case 41:
		goto sw_bb694
	case 42:
		goto sw_bb730
	case 43:
		goto sw_bb747
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
	*state_addr = 12
	goto next_state

if_end:
	v11 = *lookahead
	cmp = v11 == 46
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*state_addr = 9
	goto next_state

if_end6:
	v12 = *lookahead
	cmp7 = v12 == 48
	if cmp7 {
		goto if_then9
	} else {
		goto if_end10
	}

if_then9:
	*state_addr = 15
	goto next_state

if_end10:
	v13 = *lookahead
	cmp11 = v13 == 102
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*state_addr = 2
	goto next_state

if_end14:
	v14 = *lookahead
	cmp15 = v14 == 116
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*state_addr = 6
	goto next_state

if_end18:
	v15 = *lookahead
	cmp19 = v15 == 124
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*state_addr = 14
	goto next_state

if_end22:
	v16 = *lookahead
	cmp23 = 9 <= v16
	if cmp23 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false
	}

land_lhs_true:
	v17 = *lookahead
	cmp25 = v17 <= 13
	if cmp25 {
		goto if_then29
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v18 = *lookahead
	cmp27 = v18 == 32
	if cmp27 {
		goto if_then29
	} else {
		goto if_end30
	}

if_then29:
	*skip = 1
	*state_addr = 0
	goto next_state

if_end30:
	v19 = *lookahead
	cmp31 = 49 <= v19
	if cmp31 {
		goto land_lhs_true33
	} else {
		goto if_end37
	}

land_lhs_true33:
	v20 = *lookahead
	cmp34 = v20 <= 57
	if cmp34 {
		goto if_then36
	} else {
		goto if_end37
	}

if_then36:
	*state_addr = 16
	goto next_state

if_end37:
	v21 = *result
	tobool38 = (v21 & 1) != 0
	*retval = tobool38
	goto _return

sw_bb39:
	v22 = *lookahead
	cmp40 = v22 == 34
	if cmp40 {
		goto if_then42
	} else {
		goto if_end43
	}

if_then42:
	*state_addr = 32
	goto next_state

if_end43:
	v23 = *lookahead
	cmp44 = v23 != 0
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*state_addr = 1
	goto next_state

if_end47:
	v24 = *result
	tobool48 = (v24 & 1) != 0
	*retval = tobool48
	goto _return

sw_bb49:
	v25 = *lookahead
	cmp50 = v25 == 97
	if cmp50 {
		goto if_then52
	} else {
		goto if_end53
	}

if_then52:
	*state_addr = 5
	goto next_state

if_end53:
	v26 = *result
	tobool54 = (v26 & 1) != 0
	*retval = tobool54
	goto _return

sw_bb55:
	v27 = *lookahead
	cmp56 = v27 == 101
	if cmp56 {
		goto if_then58
	} else {
		goto if_end59
	}

if_then58:
	*state_addr = 25
	goto next_state

if_end59:
	v28 = *result
	tobool60 = (v28 & 1) != 0
	*retval = tobool60
	goto _return

sw_bb61:
	v29 = *lookahead
	cmp62 = v29 == 101
	if cmp62 {
		goto if_then64
	} else {
		goto if_end65
	}

if_then64:
	*state_addr = 27
	goto next_state

if_end65:
	v30 = *result
	tobool66 = (v30 & 1) != 0
	*retval = tobool66
	goto _return

sw_bb67:
	v31 = *lookahead
	cmp68 = v31 == 108
	if cmp68 {
		goto if_then70
	} else {
		goto if_end71
	}

if_then70:
	*state_addr = 7
	goto next_state

if_end71:
	v32 = *result
	tobool72 = (v32 & 1) != 0
	*retval = tobool72
	goto _return

sw_bb73:
	v33 = *lookahead
	cmp74 = v33 == 114
	if cmp74 {
		goto if_then76
	} else {
		goto if_end77
	}

if_then76:
	*state_addr = 8
	goto next_state

if_end77:
	v34 = *result
	tobool78 = (v34 & 1) != 0
	*retval = tobool78
	goto _return

sw_bb79:
	v35 = *lookahead
	cmp80 = v35 == 115
	if cmp80 {
		goto if_then82
	} else {
		goto if_end83
	}

if_then82:
	*state_addr = 4
	goto next_state

if_end83:
	v36 = *result
	tobool84 = (v36 & 1) != 0
	*retval = tobool84
	goto _return

sw_bb85:
	v37 = *lookahead
	cmp86 = v37 == 117
	if cmp86 {
		goto if_then88
	} else {
		goto if_end89
	}

if_then88:
	*state_addr = 3
	goto next_state

if_end89:
	v38 = *result
	tobool90 = (v38 & 1) != 0
	*retval = tobool90
	goto _return

sw_bb91:
	v39 = *lookahead
	cmp92 = 48 <= v39
	if cmp92 {
		goto land_lhs_true94
	} else {
		goto if_end98
	}

land_lhs_true94:
	v40 = *lookahead
	cmp95 = v40 <= 57
	if cmp95 {
		goto if_then97
	} else {
		goto if_end98
	}

if_then97:
	*state_addr = 21
	goto next_state

if_end98:
	v41 = *result
	tobool99 = (v41 & 1) != 0
	*retval = tobool99
	goto _return

sw_bb100:
	v42 = *lookahead
	cmp101 = 48 <= v42
	if cmp101 {
		goto land_lhs_true103
	} else {
		goto lor_lhs_false106
	}

land_lhs_true103:
	v43 = *lookahead
	cmp104 = v43 <= 57
	if cmp104 {
		goto if_then118
	} else {
		goto lor_lhs_false106
	}

lor_lhs_false106:
	v44 = *lookahead
	cmp107 = 65 <= v44
	if cmp107 {
		goto land_lhs_true109
	} else {
		goto lor_lhs_false112
	}

land_lhs_true109:
	v45 = *lookahead
	cmp110 = v45 <= 70
	if cmp110 {
		goto if_then118
	} else {
		goto lor_lhs_false112
	}

lor_lhs_false112:
	v46 = *lookahead
	cmp113 = 97 <= v46
	if cmp113 {
		goto land_lhs_true115
	} else {
		goto if_end119
	}

land_lhs_true115:
	v47 = *lookahead
	cmp116 = v47 <= 102
	if cmp116 {
		goto if_then118
	} else {
		goto if_end119
	}

if_then118:
	*state_addr = 19
	goto next_state

if_end119:
	v48 = *result
	tobool120 = (v48 & 1) != 0
	*retval = tobool120
	goto _return

sw_bb121:
	v49 = *eof
	tobool122 = (v49 & 1) != 0
	if tobool122 {
		goto if_then123
	} else {
		goto if_end124
	}

if_then123:
	*state_addr = 12
	goto next_state

if_end124:
	v50 = *lookahead
	cmp125 = v50 == 10
	if cmp125 {
		goto if_then127
	} else {
		goto if_end128
	}

if_then127:
	*state_addr = 13
	goto next_state

if_end128:
	v51 = *lookahead
	cmp129 = v51 == 13
	if cmp129 {
		goto if_then131
	} else {
		goto if_end132
	}

if_then131:
	*state_addr = 13
	goto next_state

if_end132:
	v52 = *lookahead
	cmp133 = v52 == 124
	if cmp133 {
		goto if_then135
	} else {
		goto if_end136
	}

if_then135:
	*state_addr = 14
	goto next_state

if_end136:
	v53 = *lookahead
	cmp137 = 9 <= v53
	if cmp137 {
		goto land_lhs_true139
	} else {
		goto lor_lhs_false142
	}

land_lhs_true139:
	v54 = *lookahead
	cmp140 = v54 <= 12
	if cmp140 {
		goto if_then145
	} else {
		goto lor_lhs_false142
	}

lor_lhs_false142:
	v55 = *lookahead
	cmp143 = v55 == 32
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*skip = 1
	*state_addr = 11
	goto next_state

if_end146:
	v56 = *result
	tobool147 = (v56 & 1) != 0
	*retval = tobool147
	goto _return

sw_bb148:
	*result = 1
	v57 = *lexer_addr
	result_symbol = &v57.F1
	*result_symbol = 0
	v58 = *lexer_addr
	mark_end = &v58.F3
	v59 = *mark_end
	v60 = *lexer_addr
	v59(v60)
	v61 = *result
	tobool149 = (v61 & 1) != 0
	*retval = tobool149
	goto _return

sw_bb150:
	*result = 1
	v62 = *lexer_addr
	result_symbol151 = &v62.F1
	*result_symbol151 = 1
	v63 = *lexer_addr
	mark_end152 = &v63.F3
	v64 = *mark_end152
	v65 = *lexer_addr
	v64(v65)
	v66 = *lookahead
	cmp153 = v66 == 10
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*state_addr = 13
	goto next_state

if_end156:
	v67 = *lookahead
	cmp157 = v67 == 13
	if cmp157 {
		goto if_then159
	} else {
		goto if_end160
	}

if_then159:
	*state_addr = 13
	goto next_state

if_end160:
	v68 = *result
	tobool161 = (v68 & 1) != 0
	*retval = tobool161
	goto _return

sw_bb162:
	*result = 1
	v69 = *lexer_addr
	result_symbol163 = &v69.F1
	*result_symbol163 = 2
	v70 = *lexer_addr
	mark_end164 = &v70.F3
	v71 = *mark_end164
	v72 = *lexer_addr
	v71(v72)
	v73 = *result
	tobool165 = (v73 & 1) != 0
	*retval = tobool165
	goto _return

sw_bb166:
	*result = 1
	v74 = *lexer_addr
	result_symbol167 = &v74.F1
	*result_symbol167 = 3
	v75 = *lexer_addr
	mark_end168 = &v75.F3
	v76 = *mark_end168
	v77 = *lexer_addr
	v76(v77)
	v78 = *lookahead
	cmp169 = v78 == 46
	if cmp169 {
		goto if_then171
	} else {
		goto if_end172
	}

if_then171:
	*state_addr = 23
	goto next_state

if_end172:
	v79 = *lookahead
	cmp173 = v79 == 88
	if cmp173 {
		goto if_then178
	} else {
		goto lor_lhs_false175
	}

lor_lhs_false175:
	v80 = *lookahead
	cmp176 = v80 == 120
	if cmp176 {
		goto if_then178
	} else {
		goto if_end179
	}

if_then178:
	*state_addr = 10
	goto next_state

if_end179:
	v81 = *lookahead
	cmp180 = 48 <= v81
	if cmp180 {
		goto land_lhs_true182
	} else {
		goto if_end186
	}

land_lhs_true182:
	v82 = *lookahead
	cmp183 = v82 <= 57
	if cmp183 {
		goto if_then185
	} else {
		goto if_end186
	}

if_then185:
	*state_addr = 16
	goto next_state

if_end186:
	v83 = *result
	tobool187 = (v83 & 1) != 0
	*retval = tobool187
	goto _return

sw_bb188:
	*result = 1
	v84 = *lexer_addr
	result_symbol189 = &v84.F1
	*result_symbol189 = 3
	v85 = *lexer_addr
	mark_end190 = &v85.F3
	v86 = *mark_end190
	v87 = *lexer_addr
	v86(v87)
	v88 = *lookahead
	cmp191 = v88 == 46
	if cmp191 {
		goto if_then193
	} else {
		goto if_end194
	}

if_then193:
	*state_addr = 23
	goto next_state

if_end194:
	v89 = *lookahead
	cmp195 = 48 <= v89
	if cmp195 {
		goto land_lhs_true197
	} else {
		goto if_end201
	}

land_lhs_true197:
	v90 = *lookahead
	cmp198 = v90 <= 57
	if cmp198 {
		goto if_then200
	} else {
		goto if_end201
	}

if_then200:
	*state_addr = 16
	goto next_state

if_end201:
	v91 = *result
	tobool202 = (v91 & 1) != 0
	*retval = tobool202
	goto _return

sw_bb203:
	*result = 1
	v92 = *lexer_addr
	result_symbol204 = &v92.F1
	*result_symbol204 = 3
	v93 = *lexer_addr
	mark_end205 = &v93.F3
	v94 = *mark_end205
	v95 = *lexer_addr
	v94(v95)
	v96 = *lookahead
	cmp206 = v96 == 46
	if cmp206 {
		goto if_then208
	} else {
		goto if_end209
	}

if_then208:
	*state_addr = 24
	goto next_state

if_end209:
	v97 = *lookahead
	cmp210 = v97 == 88
	if cmp210 {
		goto if_then215
	} else {
		goto lor_lhs_false212
	}

lor_lhs_false212:
	v98 = *lookahead
	cmp213 = v98 == 120
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*state_addr = 41
	goto next_state

if_end216:
	v99 = *lookahead
	cmp217 = 48 <= v99
	if cmp217 {
		goto land_lhs_true219
	} else {
		goto if_end223
	}

land_lhs_true219:
	v100 = *lookahead
	cmp220 = v100 <= 57
	if cmp220 {
		goto if_then222
	} else {
		goto if_end223
	}

if_then222:
	*state_addr = 18
	goto next_state

if_end223:
	v101 = *lookahead
	cmp224 = v101 != 0
	if cmp224 {
		goto land_lhs_true226
	} else {
		goto if_end236
	}

land_lhs_true226:
	v102 = *lookahead
	cmp227 = v102 != 10
	if cmp227 {
		goto land_lhs_true229
	} else {
		goto if_end236
	}

land_lhs_true229:
	v103 = *lookahead
	cmp230 = v103 != 13
	if cmp230 {
		goto land_lhs_true232
	} else {
		goto if_end236
	}

land_lhs_true232:
	v104 = *lookahead
	cmp233 = v104 != 124
	if cmp233 {
		goto if_then235
	} else {
		goto if_end236
	}

if_then235:
	*state_addr = 42
	goto next_state

if_end236:
	v105 = *result
	tobool237 = (v105 & 1) != 0
	*retval = tobool237
	goto _return

sw_bb238:
	*result = 1
	v106 = *lexer_addr
	result_symbol239 = &v106.F1
	*result_symbol239 = 3
	v107 = *lexer_addr
	mark_end240 = &v107.F3
	v108 = *mark_end240
	v109 = *lexer_addr
	v108(v109)
	v110 = *lookahead
	cmp241 = v110 == 46
	if cmp241 {
		goto if_then243
	} else {
		goto if_end244
	}

if_then243:
	*state_addr = 24
	goto next_state

if_end244:
	v111 = *lookahead
	cmp245 = 48 <= v111
	if cmp245 {
		goto land_lhs_true247
	} else {
		goto if_end251
	}

land_lhs_true247:
	v112 = *lookahead
	cmp248 = v112 <= 57
	if cmp248 {
		goto if_then250
	} else {
		goto if_end251
	}

if_then250:
	*state_addr = 18
	goto next_state

if_end251:
	v113 = *lookahead
	cmp252 = v113 != 0
	if cmp252 {
		goto land_lhs_true254
	} else {
		goto if_end264
	}

land_lhs_true254:
	v114 = *lookahead
	cmp255 = v114 != 10
	if cmp255 {
		goto land_lhs_true257
	} else {
		goto if_end264
	}

land_lhs_true257:
	v115 = *lookahead
	cmp258 = v115 != 13
	if cmp258 {
		goto land_lhs_true260
	} else {
		goto if_end264
	}

land_lhs_true260:
	v116 = *lookahead
	cmp261 = v116 != 124
	if cmp261 {
		goto if_then263
	} else {
		goto if_end264
	}

if_then263:
	*state_addr = 42
	goto next_state

if_end264:
	v117 = *result
	tobool265 = (v117 & 1) != 0
	*retval = tobool265
	goto _return

sw_bb266:
	*result = 1
	v118 = *lexer_addr
	result_symbol267 = &v118.F1
	*result_symbol267 = 4
	v119 = *lexer_addr
	mark_end268 = &v119.F3
	v120 = *mark_end268
	v121 = *lexer_addr
	v120(v121)
	v122 = *lookahead
	cmp269 = 48 <= v122
	if cmp269 {
		goto land_lhs_true271
	} else {
		goto lor_lhs_false274
	}

land_lhs_true271:
	v123 = *lookahead
	cmp272 = v123 <= 57
	if cmp272 {
		goto if_then286
	} else {
		goto lor_lhs_false274
	}

lor_lhs_false274:
	v124 = *lookahead
	cmp275 = 65 <= v124
	if cmp275 {
		goto land_lhs_true277
	} else {
		goto lor_lhs_false280
	}

land_lhs_true277:
	v125 = *lookahead
	cmp278 = v125 <= 70
	if cmp278 {
		goto if_then286
	} else {
		goto lor_lhs_false280
	}

lor_lhs_false280:
	v126 = *lookahead
	cmp281 = 97 <= v126
	if cmp281 {
		goto land_lhs_true283
	} else {
		goto if_end287
	}

land_lhs_true283:
	v127 = *lookahead
	cmp284 = v127 <= 102
	if cmp284 {
		goto if_then286
	} else {
		goto if_end287
	}

if_then286:
	*state_addr = 19
	goto next_state

if_end287:
	v128 = *result
	tobool288 = (v128 & 1) != 0
	*retval = tobool288
	goto _return

sw_bb289:
	*result = 1
	v129 = *lexer_addr
	result_symbol290 = &v129.F1
	*result_symbol290 = 4
	v130 = *lexer_addr
	mark_end291 = &v130.F3
	v131 = *mark_end291
	v132 = *lexer_addr
	v131(v132)
	v133 = *lookahead
	cmp292 = 48 <= v133
	if cmp292 {
		goto land_lhs_true294
	} else {
		goto lor_lhs_false297
	}

land_lhs_true294:
	v134 = *lookahead
	cmp295 = v134 <= 57
	if cmp295 {
		goto if_then309
	} else {
		goto lor_lhs_false297
	}

lor_lhs_false297:
	v135 = *lookahead
	cmp298 = 65 <= v135
	if cmp298 {
		goto land_lhs_true300
	} else {
		goto lor_lhs_false303
	}

land_lhs_true300:
	v136 = *lookahead
	cmp301 = v136 <= 70
	if cmp301 {
		goto if_then309
	} else {
		goto lor_lhs_false303
	}

lor_lhs_false303:
	v137 = *lookahead
	cmp304 = 97 <= v137
	if cmp304 {
		goto land_lhs_true306
	} else {
		goto if_end310
	}

land_lhs_true306:
	v138 = *lookahead
	cmp307 = v138 <= 102
	if cmp307 {
		goto if_then309
	} else {
		goto if_end310
	}

if_then309:
	*state_addr = 20
	goto next_state

if_end310:
	v139 = *lookahead
	cmp311 = v139 != 0
	if cmp311 {
		goto land_lhs_true313
	} else {
		goto if_end323
	}

land_lhs_true313:
	v140 = *lookahead
	cmp314 = v140 != 10
	if cmp314 {
		goto land_lhs_true316
	} else {
		goto if_end323
	}

land_lhs_true316:
	v141 = *lookahead
	cmp317 = v141 != 13
	if cmp317 {
		goto land_lhs_true319
	} else {
		goto if_end323
	}

land_lhs_true319:
	v142 = *lookahead
	cmp320 = v142 != 124
	if cmp320 {
		goto if_then322
	} else {
		goto if_end323
	}

if_then322:
	*state_addr = 42
	goto next_state

if_end323:
	v143 = *result
	tobool324 = (v143 & 1) != 0
	*retval = tobool324
	goto _return

sw_bb325:
	*result = 1
	v144 = *lexer_addr
	result_symbol326 = &v144.F1
	*result_symbol326 = 5
	v145 = *lexer_addr
	mark_end327 = &v145.F3
	v146 = *mark_end327
	v147 = *lexer_addr
	v146(v147)
	v148 = *lookahead
	cmp328 = 48 <= v148
	if cmp328 {
		goto land_lhs_true330
	} else {
		goto if_end334
	}

land_lhs_true330:
	v149 = *lookahead
	cmp331 = v149 <= 57
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*state_addr = 21
	goto next_state

if_end334:
	v150 = *result
	tobool335 = (v150 & 1) != 0
	*retval = tobool335
	goto _return

sw_bb336:
	*result = 1
	v151 = *lexer_addr
	result_symbol337 = &v151.F1
	*result_symbol337 = 5
	v152 = *lexer_addr
	mark_end338 = &v152.F3
	v153 = *mark_end338
	v154 = *lexer_addr
	v153(v154)
	v155 = *lookahead
	cmp339 = 48 <= v155
	if cmp339 {
		goto land_lhs_true341
	} else {
		goto if_end345
	}

land_lhs_true341:
	v156 = *lookahead
	cmp342 = v156 <= 57
	if cmp342 {
		goto if_then344
	} else {
		goto if_end345
	}

if_then344:
	*state_addr = 22
	goto next_state

if_end345:
	v157 = *lookahead
	cmp346 = v157 != 0
	if cmp346 {
		goto land_lhs_true348
	} else {
		goto if_end358
	}

land_lhs_true348:
	v158 = *lookahead
	cmp349 = v158 != 10
	if cmp349 {
		goto land_lhs_true351
	} else {
		goto if_end358
	}

land_lhs_true351:
	v159 = *lookahead
	cmp352 = v159 != 13
	if cmp352 {
		goto land_lhs_true354
	} else {
		goto if_end358
	}

land_lhs_true354:
	v160 = *lookahead
	cmp355 = v160 != 124
	if cmp355 {
		goto if_then357
	} else {
		goto if_end358
	}

if_then357:
	*state_addr = 42
	goto next_state

if_end358:
	v161 = *result
	tobool359 = (v161 & 1) != 0
	*retval = tobool359
	goto _return

sw_bb360:
	*result = 1
	v162 = *lexer_addr
	result_symbol361 = &v162.F1
	*result_symbol361 = 6
	v163 = *lexer_addr
	mark_end362 = &v163.F3
	v164 = *mark_end362
	v165 = *lexer_addr
	v164(v165)
	v166 = *lookahead
	cmp363 = 48 <= v166
	if cmp363 {
		goto land_lhs_true365
	} else {
		goto if_end369
	}

land_lhs_true365:
	v167 = *lookahead
	cmp366 = v167 <= 57
	if cmp366 {
		goto if_then368
	} else {
		goto if_end369
	}

if_then368:
	*state_addr = 21
	goto next_state

if_end369:
	v168 = *result
	tobool370 = (v168 & 1) != 0
	*retval = tobool370
	goto _return

sw_bb371:
	*result = 1
	v169 = *lexer_addr
	result_symbol372 = &v169.F1
	*result_symbol372 = 6
	v170 = *lexer_addr
	mark_end373 = &v170.F3
	v171 = *mark_end373
	v172 = *lexer_addr
	v171(v172)
	v173 = *lookahead
	cmp374 = 48 <= v173
	if cmp374 {
		goto land_lhs_true376
	} else {
		goto if_end380
	}

land_lhs_true376:
	v174 = *lookahead
	cmp377 = v174 <= 57
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*state_addr = 22
	goto next_state

if_end380:
	v175 = *lookahead
	cmp381 = v175 != 0
	if cmp381 {
		goto land_lhs_true383
	} else {
		goto if_end393
	}

land_lhs_true383:
	v176 = *lookahead
	cmp384 = v176 != 10
	if cmp384 {
		goto land_lhs_true386
	} else {
		goto if_end393
	}

land_lhs_true386:
	v177 = *lookahead
	cmp387 = v177 != 13
	if cmp387 {
		goto land_lhs_true389
	} else {
		goto if_end393
	}

land_lhs_true389:
	v178 = *lookahead
	cmp390 = v178 != 124
	if cmp390 {
		goto if_then392
	} else {
		goto if_end393
	}

if_then392:
	*state_addr = 42
	goto next_state

if_end393:
	v179 = *result
	tobool394 = (v179 & 1) != 0
	*retval = tobool394
	goto _return

sw_bb395:
	*result = 1
	v180 = *lexer_addr
	result_symbol396 = &v180.F1
	*result_symbol396 = 7
	v181 = *lexer_addr
	mark_end397 = &v181.F3
	v182 = *mark_end397
	v183 = *lexer_addr
	v182(v183)
	v184 = *result
	tobool398 = (v184 & 1) != 0
	*retval = tobool398
	goto _return

sw_bb399:
	*result = 1
	v185 = *lexer_addr
	result_symbol400 = &v185.F1
	*result_symbol400 = 7
	v186 = *lexer_addr
	mark_end401 = &v186.F3
	v187 = *mark_end401
	v188 = *lexer_addr
	v187(v188)
	v189 = *lookahead
	cmp402 = v189 != 0
	if cmp402 {
		goto land_lhs_true404
	} else {
		goto if_end414
	}

land_lhs_true404:
	v190 = *lookahead
	cmp405 = v190 != 10
	if cmp405 {
		goto land_lhs_true407
	} else {
		goto if_end414
	}

land_lhs_true407:
	v191 = *lookahead
	cmp408 = v191 != 13
	if cmp408 {
		goto land_lhs_true410
	} else {
		goto if_end414
	}

land_lhs_true410:
	v192 = *lookahead
	cmp411 = v192 != 124
	if cmp411 {
		goto if_then413
	} else {
		goto if_end414
	}

if_then413:
	*state_addr = 42
	goto next_state

if_end414:
	v193 = *result
	tobool415 = (v193 & 1) != 0
	*retval = tobool415
	goto _return

sw_bb416:
	*result = 1
	v194 = *lexer_addr
	result_symbol417 = &v194.F1
	*result_symbol417 = 8
	v195 = *lexer_addr
	mark_end418 = &v195.F3
	v196 = *mark_end418
	v197 = *lexer_addr
	v196(v197)
	v198 = *result
	tobool419 = (v198 & 1) != 0
	*retval = tobool419
	goto _return

sw_bb420:
	*result = 1
	v199 = *lexer_addr
	result_symbol421 = &v199.F1
	*result_symbol421 = 8
	v200 = *lexer_addr
	mark_end422 = &v200.F3
	v201 = *mark_end422
	v202 = *lexer_addr
	v201(v202)
	v203 = *lookahead
	cmp423 = v203 != 0
	if cmp423 {
		goto land_lhs_true425
	} else {
		goto if_end435
	}

land_lhs_true425:
	v204 = *lookahead
	cmp426 = v204 != 10
	if cmp426 {
		goto land_lhs_true428
	} else {
		goto if_end435
	}

land_lhs_true428:
	v205 = *lookahead
	cmp429 = v205 != 13
	if cmp429 {
		goto land_lhs_true431
	} else {
		goto if_end435
	}

land_lhs_true431:
	v206 = *lookahead
	cmp432 = v206 != 124
	if cmp432 {
		goto if_then434
	} else {
		goto if_end435
	}

if_then434:
	*state_addr = 42
	goto next_state

if_end435:
	v207 = *result
	tobool436 = (v207 & 1) != 0
	*retval = tobool436
	goto _return

sw_bb437:
	*result = 1
	v208 = *lexer_addr
	result_symbol438 = &v208.F1
	*result_symbol438 = 9
	v209 = *lexer_addr
	mark_end439 = &v209.F3
	v210 = *mark_end439
	v211 = *lexer_addr
	v210(v211)
	*i = 0
	goto for_cond

for_cond:
	v212 = *i
	conv440 = int64(uint64(uint32(v212)))
	cmp441 = uint64(conv440) < uint64(18)
	if cmp441 {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v213 = *i
	idxprom = int64(uint64(uint32(v213)))
	arrayidx = &ts_lex_map[idxprom]
	v214 = *arrayidx
	conv443 = int32(uint32(uint16(v214)))
	v215 = *lookahead
	cmp444 = conv443 == v215
	if cmp444 {
		goto if_then446
	} else {
		goto if_end449
	}

if_then446:
	v216 = *i
	add = v216 + 1
	idxprom447 = int64(uint64(uint32(add)))
	arrayidx448 = &ts_lex_map[idxprom447]
	v217 = *arrayidx448
	*state_addr = v217
	goto next_state

if_end449:
	goto for_inc

for_inc:
	v218 = *i
	add450 = v218 + 2
	*i = add450
	goto for_cond

for_end:
	v219 = *lookahead
	cmp451 = 49 <= v219
	if cmp451 {
		goto land_lhs_true453
	} else {
		goto if_end457
	}

land_lhs_true453:
	v220 = *lookahead
	cmp454 = v220 <= 57
	if cmp454 {
		goto if_then456
	} else {
		goto if_end457
	}

if_then456:
	*state_addr = 18
	goto next_state

if_end457:
	v221 = *lookahead
	cmp458 = v221 != 0
	if cmp458 {
		goto land_lhs_true460
	} else {
		goto if_end470
	}

land_lhs_true460:
	v222 = *lookahead
	cmp461 = v222 < 9
	if cmp461 {
		goto land_lhs_true466
	} else {
		goto lor_lhs_false463
	}

lor_lhs_false463:
	v223 = *lookahead
	cmp464 = 13 < v223
	if cmp464 {
		goto land_lhs_true466
	} else {
		goto if_end470
	}

land_lhs_true466:
	v224 = *lookahead
	cmp467 = v224 != 124
	if cmp467 {
		goto if_then469
	} else {
		goto if_end470
	}

if_then469:
	*state_addr = 42
	goto next_state

if_end470:
	v225 = *result
	tobool471 = (v225 & 1) != 0
	*retval = tobool471
	goto _return

sw_bb472:
	*result = 1
	v226 = *lexer_addr
	result_symbol473 = &v226.F1
	*result_symbol473 = 9
	v227 = *lexer_addr
	mark_end474 = &v227.F3
	v228 = *mark_end474
	v229 = *lexer_addr
	v228(v229)
	v230 = *lookahead
	cmp475 = v230 == 34
	if cmp475 {
		goto if_then477
	} else {
		goto if_end478
	}

if_then477:
	*state_addr = 31
	goto next_state

if_end478:
	v231 = *lookahead
	cmp479 = v231 != 0
	if cmp479 {
		goto land_lhs_true481
	} else {
		goto if_end491
	}

land_lhs_true481:
	v232 = *lookahead
	cmp482 = v232 != 10
	if cmp482 {
		goto land_lhs_true484
	} else {
		goto if_end491
	}

land_lhs_true484:
	v233 = *lookahead
	cmp485 = v233 != 13
	if cmp485 {
		goto land_lhs_true487
	} else {
		goto if_end491
	}

land_lhs_true487:
	v234 = *lookahead
	cmp488 = v234 != 124
	if cmp488 {
		goto if_then490
	} else {
		goto if_end491
	}

if_then490:
	*state_addr = 42
	goto next_state

if_end491:
	v235 = *result
	tobool492 = (v235 & 1) != 0
	*retval = tobool492
	goto _return

sw_bb493:
	*result = 1
	v236 = *lexer_addr
	result_symbol494 = &v236.F1
	*result_symbol494 = 9
	v237 = *lexer_addr
	mark_end495 = &v237.F3
	v238 = *mark_end495
	v239 = *lexer_addr
	v238(v239)
	v240 = *lookahead
	cmp496 = v240 == 34
	if cmp496 {
		goto if_then498
	} else {
		goto if_end499
	}

if_then498:
	*state_addr = 30
	goto next_state

if_end499:
	v241 = *lookahead
	cmp500 = v241 == 10
	if cmp500 {
		goto if_then508
	} else {
		goto lor_lhs_false502
	}

lor_lhs_false502:
	v242 = *lookahead
	cmp503 = v242 == 13
	if cmp503 {
		goto if_then508
	} else {
		goto lor_lhs_false505
	}

lor_lhs_false505:
	v243 = *lookahead
	cmp506 = v243 == 124
	if cmp506 {
		goto if_then508
	} else {
		goto if_end509
	}

if_then508:
	*state_addr = 1
	goto next_state

if_end509:
	v244 = *lookahead
	cmp510 = v244 != 0
	if cmp510 {
		goto if_then512
	} else {
		goto if_end513
	}

if_then512:
	*state_addr = 31
	goto next_state

if_end513:
	v245 = *result
	tobool514 = (v245 & 1) != 0
	*retval = tobool514
	goto _return

sw_bb515:
	*result = 1
	v246 = *lexer_addr
	result_symbol516 = &v246.F1
	*result_symbol516 = 9
	v247 = *lexer_addr
	mark_end517 = &v247.F3
	v248 = *mark_end517
	v249 = *lexer_addr
	v248(v249)
	v250 = *lookahead
	cmp518 = v250 == 34
	if cmp518 {
		goto if_then520
	} else {
		goto if_end521
	}

if_then520:
	*state_addr = 1
	goto next_state

if_end521:
	v251 = *result
	tobool522 = (v251 & 1) != 0
	*retval = tobool522
	goto _return

sw_bb523:
	*result = 1
	v252 = *lexer_addr
	result_symbol524 = &v252.F1
	*result_symbol524 = 9
	v253 = *lexer_addr
	mark_end525 = &v253.F3
	v254 = *mark_end525
	v255 = *lexer_addr
	v254(v255)
	v256 = *lookahead
	cmp526 = v256 == 97
	if cmp526 {
		goto if_then528
	} else {
		goto if_end529
	}

if_then528:
	*state_addr = 36
	goto next_state

if_end529:
	v257 = *lookahead
	cmp530 = v257 != 0
	if cmp530 {
		goto land_lhs_true532
	} else {
		goto if_end542
	}

land_lhs_true532:
	v258 = *lookahead
	cmp533 = v258 != 10
	if cmp533 {
		goto land_lhs_true535
	} else {
		goto if_end542
	}

land_lhs_true535:
	v259 = *lookahead
	cmp536 = v259 != 13
	if cmp536 {
		goto land_lhs_true538
	} else {
		goto if_end542
	}

land_lhs_true538:
	v260 = *lookahead
	cmp539 = v260 != 124
	if cmp539 {
		goto if_then541
	} else {
		goto if_end542
	}

if_then541:
	*state_addr = 42
	goto next_state

if_end542:
	v261 = *result
	tobool543 = (v261 & 1) != 0
	*retval = tobool543
	goto _return

sw_bb544:
	*result = 1
	v262 = *lexer_addr
	result_symbol545 = &v262.F1
	*result_symbol545 = 9
	v263 = *lexer_addr
	mark_end546 = &v263.F3
	v264 = *mark_end546
	v265 = *lexer_addr
	v264(v265)
	v266 = *lookahead
	cmp547 = v266 == 101
	if cmp547 {
		goto if_then549
	} else {
		goto if_end550
	}

if_then549:
	*state_addr = 26
	goto next_state

if_end550:
	v267 = *lookahead
	cmp551 = v267 != 0
	if cmp551 {
		goto land_lhs_true553
	} else {
		goto if_end563
	}

land_lhs_true553:
	v268 = *lookahead
	cmp554 = v268 != 10
	if cmp554 {
		goto land_lhs_true556
	} else {
		goto if_end563
	}

land_lhs_true556:
	v269 = *lookahead
	cmp557 = v269 != 13
	if cmp557 {
		goto land_lhs_true559
	} else {
		goto if_end563
	}

land_lhs_true559:
	v270 = *lookahead
	cmp560 = v270 != 124
	if cmp560 {
		goto if_then562
	} else {
		goto if_end563
	}

if_then562:
	*state_addr = 42
	goto next_state

if_end563:
	v271 = *result
	tobool564 = (v271 & 1) != 0
	*retval = tobool564
	goto _return

sw_bb565:
	*result = 1
	v272 = *lexer_addr
	result_symbol566 = &v272.F1
	*result_symbol566 = 9
	v273 = *lexer_addr
	mark_end567 = &v273.F3
	v274 = *mark_end567
	v275 = *lexer_addr
	v274(v275)
	v276 = *lookahead
	cmp568 = v276 == 101
	if cmp568 {
		goto if_then570
	} else {
		goto if_end571
	}

if_then570:
	*state_addr = 28
	goto next_state

if_end571:
	v277 = *lookahead
	cmp572 = v277 != 0
	if cmp572 {
		goto land_lhs_true574
	} else {
		goto if_end584
	}

land_lhs_true574:
	v278 = *lookahead
	cmp575 = v278 != 10
	if cmp575 {
		goto land_lhs_true577
	} else {
		goto if_end584
	}

land_lhs_true577:
	v279 = *lookahead
	cmp578 = v279 != 13
	if cmp578 {
		goto land_lhs_true580
	} else {
		goto if_end584
	}

land_lhs_true580:
	v280 = *lookahead
	cmp581 = v280 != 124
	if cmp581 {
		goto if_then583
	} else {
		goto if_end584
	}

if_then583:
	*state_addr = 42
	goto next_state

if_end584:
	v281 = *result
	tobool585 = (v281 & 1) != 0
	*retval = tobool585
	goto _return

sw_bb586:
	*result = 1
	v282 = *lexer_addr
	result_symbol587 = &v282.F1
	*result_symbol587 = 9
	v283 = *lexer_addr
	mark_end588 = &v283.F3
	v284 = *mark_end588
	v285 = *lexer_addr
	v284(v285)
	v286 = *lookahead
	cmp589 = v286 == 108
	if cmp589 {
		goto if_then591
	} else {
		goto if_end592
	}

if_then591:
	*state_addr = 38
	goto next_state

if_end592:
	v287 = *lookahead
	cmp593 = v287 != 0
	if cmp593 {
		goto land_lhs_true595
	} else {
		goto if_end605
	}

land_lhs_true595:
	v288 = *lookahead
	cmp596 = v288 != 10
	if cmp596 {
		goto land_lhs_true598
	} else {
		goto if_end605
	}

land_lhs_true598:
	v289 = *lookahead
	cmp599 = v289 != 13
	if cmp599 {
		goto land_lhs_true601
	} else {
		goto if_end605
	}

land_lhs_true601:
	v290 = *lookahead
	cmp602 = v290 != 124
	if cmp602 {
		goto if_then604
	} else {
		goto if_end605
	}

if_then604:
	*state_addr = 42
	goto next_state

if_end605:
	v291 = *result
	tobool606 = (v291 & 1) != 0
	*retval = tobool606
	goto _return

sw_bb607:
	*result = 1
	v292 = *lexer_addr
	result_symbol608 = &v292.F1
	*result_symbol608 = 9
	v293 = *lexer_addr
	mark_end609 = &v293.F3
	v294 = *mark_end609
	v295 = *lexer_addr
	v294(v295)
	v296 = *lookahead
	cmp610 = v296 == 114
	if cmp610 {
		goto if_then612
	} else {
		goto if_end613
	}

if_then612:
	*state_addr = 39
	goto next_state

if_end613:
	v297 = *lookahead
	cmp614 = v297 != 0
	if cmp614 {
		goto land_lhs_true616
	} else {
		goto if_end626
	}

land_lhs_true616:
	v298 = *lookahead
	cmp617 = v298 != 10
	if cmp617 {
		goto land_lhs_true619
	} else {
		goto if_end626
	}

land_lhs_true619:
	v299 = *lookahead
	cmp620 = v299 != 13
	if cmp620 {
		goto land_lhs_true622
	} else {
		goto if_end626
	}

land_lhs_true622:
	v300 = *lookahead
	cmp623 = v300 != 124
	if cmp623 {
		goto if_then625
	} else {
		goto if_end626
	}

if_then625:
	*state_addr = 42
	goto next_state

if_end626:
	v301 = *result
	tobool627 = (v301 & 1) != 0
	*retval = tobool627
	goto _return

sw_bb628:
	*result = 1
	v302 = *lexer_addr
	result_symbol629 = &v302.F1
	*result_symbol629 = 9
	v303 = *lexer_addr
	mark_end630 = &v303.F3
	v304 = *mark_end630
	v305 = *lexer_addr
	v304(v305)
	v306 = *lookahead
	cmp631 = v306 == 115
	if cmp631 {
		goto if_then633
	} else {
		goto if_end634
	}

if_then633:
	*state_addr = 35
	goto next_state

if_end634:
	v307 = *lookahead
	cmp635 = v307 != 0
	if cmp635 {
		goto land_lhs_true637
	} else {
		goto if_end647
	}

land_lhs_true637:
	v308 = *lookahead
	cmp638 = v308 != 10
	if cmp638 {
		goto land_lhs_true640
	} else {
		goto if_end647
	}

land_lhs_true640:
	v309 = *lookahead
	cmp641 = v309 != 13
	if cmp641 {
		goto land_lhs_true643
	} else {
		goto if_end647
	}

land_lhs_true643:
	v310 = *lookahead
	cmp644 = v310 != 124
	if cmp644 {
		goto if_then646
	} else {
		goto if_end647
	}

if_then646:
	*state_addr = 42
	goto next_state

if_end647:
	v311 = *result
	tobool648 = (v311 & 1) != 0
	*retval = tobool648
	goto _return

sw_bb649:
	*result = 1
	v312 = *lexer_addr
	result_symbol650 = &v312.F1
	*result_symbol650 = 9
	v313 = *lexer_addr
	mark_end651 = &v313.F3
	v314 = *mark_end651
	v315 = *lexer_addr
	v314(v315)
	v316 = *lookahead
	cmp652 = v316 == 117
	if cmp652 {
		goto if_then654
	} else {
		goto if_end655
	}

if_then654:
	*state_addr = 34
	goto next_state

if_end655:
	v317 = *lookahead
	cmp656 = v317 != 0
	if cmp656 {
		goto land_lhs_true658
	} else {
		goto if_end668
	}

land_lhs_true658:
	v318 = *lookahead
	cmp659 = v318 != 10
	if cmp659 {
		goto land_lhs_true661
	} else {
		goto if_end668
	}

land_lhs_true661:
	v319 = *lookahead
	cmp662 = v319 != 13
	if cmp662 {
		goto land_lhs_true664
	} else {
		goto if_end668
	}

land_lhs_true664:
	v320 = *lookahead
	cmp665 = v320 != 124
	if cmp665 {
		goto if_then667
	} else {
		goto if_end668
	}

if_then667:
	*state_addr = 42
	goto next_state

if_end668:
	v321 = *result
	tobool669 = (v321 & 1) != 0
	*retval = tobool669
	goto _return

sw_bb670:
	*result = 1
	v322 = *lexer_addr
	result_symbol671 = &v322.F1
	*result_symbol671 = 9
	v323 = *lexer_addr
	mark_end672 = &v323.F3
	v324 = *mark_end672
	v325 = *lexer_addr
	v324(v325)
	v326 = *lookahead
	cmp673 = 48 <= v326
	if cmp673 {
		goto land_lhs_true675
	} else {
		goto if_end679
	}

land_lhs_true675:
	v327 = *lookahead
	cmp676 = v327 <= 57
	if cmp676 {
		goto if_then678
	} else {
		goto if_end679
	}

if_then678:
	*state_addr = 22
	goto next_state

if_end679:
	v328 = *lookahead
	cmp680 = v328 != 0
	if cmp680 {
		goto land_lhs_true682
	} else {
		goto if_end692
	}

land_lhs_true682:
	v329 = *lookahead
	cmp683 = v329 != 10
	if cmp683 {
		goto land_lhs_true685
	} else {
		goto if_end692
	}

land_lhs_true685:
	v330 = *lookahead
	cmp686 = v330 != 13
	if cmp686 {
		goto land_lhs_true688
	} else {
		goto if_end692
	}

land_lhs_true688:
	v331 = *lookahead
	cmp689 = v331 != 124
	if cmp689 {
		goto if_then691
	} else {
		goto if_end692
	}

if_then691:
	*state_addr = 42
	goto next_state

if_end692:
	v332 = *result
	tobool693 = (v332 & 1) != 0
	*retval = tobool693
	goto _return

sw_bb694:
	*result = 1
	v333 = *lexer_addr
	result_symbol695 = &v333.F1
	*result_symbol695 = 9
	v334 = *lexer_addr
	mark_end696 = &v334.F3
	v335 = *mark_end696
	v336 = *lexer_addr
	v335(v336)
	v337 = *lookahead
	cmp697 = 48 <= v337
	if cmp697 {
		goto land_lhs_true699
	} else {
		goto lor_lhs_false702
	}

land_lhs_true699:
	v338 = *lookahead
	cmp700 = v338 <= 57
	if cmp700 {
		goto if_then714
	} else {
		goto lor_lhs_false702
	}

lor_lhs_false702:
	v339 = *lookahead
	cmp703 = 65 <= v339
	if cmp703 {
		goto land_lhs_true705
	} else {
		goto lor_lhs_false708
	}

land_lhs_true705:
	v340 = *lookahead
	cmp706 = v340 <= 70
	if cmp706 {
		goto if_then714
	} else {
		goto lor_lhs_false708
	}

lor_lhs_false708:
	v341 = *lookahead
	cmp709 = 97 <= v341
	if cmp709 {
		goto land_lhs_true711
	} else {
		goto if_end715
	}

land_lhs_true711:
	v342 = *lookahead
	cmp712 = v342 <= 102
	if cmp712 {
		goto if_then714
	} else {
		goto if_end715
	}

if_then714:
	*state_addr = 20
	goto next_state

if_end715:
	v343 = *lookahead
	cmp716 = v343 != 0
	if cmp716 {
		goto land_lhs_true718
	} else {
		goto if_end728
	}

land_lhs_true718:
	v344 = *lookahead
	cmp719 = v344 != 10
	if cmp719 {
		goto land_lhs_true721
	} else {
		goto if_end728
	}

land_lhs_true721:
	v345 = *lookahead
	cmp722 = v345 != 13
	if cmp722 {
		goto land_lhs_true724
	} else {
		goto if_end728
	}

land_lhs_true724:
	v346 = *lookahead
	cmp725 = v346 != 124
	if cmp725 {
		goto if_then727
	} else {
		goto if_end728
	}

if_then727:
	*state_addr = 42
	goto next_state

if_end728:
	v347 = *result
	tobool729 = (v347 & 1) != 0
	*retval = tobool729
	goto _return

sw_bb730:
	*result = 1
	v348 = *lexer_addr
	result_symbol731 = &v348.F1
	*result_symbol731 = 9
	v349 = *lexer_addr
	mark_end732 = &v349.F3
	v350 = *mark_end732
	v351 = *lexer_addr
	v350(v351)
	v352 = *lookahead
	cmp733 = v352 != 0
	if cmp733 {
		goto land_lhs_true735
	} else {
		goto if_end745
	}

land_lhs_true735:
	v353 = *lookahead
	cmp736 = v353 != 10
	if cmp736 {
		goto land_lhs_true738
	} else {
		goto if_end745
	}

land_lhs_true738:
	v354 = *lookahead
	cmp739 = v354 != 13
	if cmp739 {
		goto land_lhs_true741
	} else {
		goto if_end745
	}

land_lhs_true741:
	v355 = *lookahead
	cmp742 = v355 != 124
	if cmp742 {
		goto if_then744
	} else {
		goto if_end745
	}

if_then744:
	*state_addr = 42
	goto next_state

if_end745:
	v356 = *result
	tobool746 = (v356 & 1) != 0
	*retval = tobool746
	goto _return

sw_bb747:
	*result = 1
	v357 = *lexer_addr
	result_symbol748 = &v357.F1
	*result_symbol748 = 9
	v358 = *lexer_addr
	mark_end749 = &v358.F3
	v359 = *mark_end749
	v360 = *lexer_addr
	v359(v360)
	v361 = *eof
	tobool750 = (v361 & 1) != 0
	if tobool750 {
		goto if_then751
	} else {
		goto if_end752
	}

if_then751:
	*state_addr = 12
	goto next_state

if_end752:
	*i753 = 0
	goto for_cond754

for_cond754:
	v362 = *i753
	conv755 = int64(uint64(uint32(v362)))
	cmp756 = uint64(conv755) < uint64(18)
	if cmp756 {
		goto for_body758
	} else {
		goto for_end771
	}

for_body758:
	v363 = *i753
	idxprom759 = int64(uint64(uint32(v363)))
	arrayidx760 = &ts_lex_map_21[idxprom759]
	v364 = *arrayidx760
	conv761 = int32(uint32(uint16(v364)))
	v365 = *lookahead
	cmp762 = conv761 == v365
	if cmp762 {
		goto if_then764
	} else {
		goto if_end768
	}

if_then764:
	v366 = *i753
	add765 = v366 + 1
	idxprom766 = int64(uint64(uint32(add765)))
	arrayidx767 = &ts_lex_map_21[idxprom766]
	v367 = *arrayidx767
	*state_addr = v367
	goto next_state

if_end768:
	goto for_inc769

for_inc769:
	v368 = *i753
	add770 = v368 + 2
	*i753 = add770
	goto for_cond754

for_end771:
	v369 = *lookahead
	cmp772 = 49 <= v369
	if cmp772 {
		goto land_lhs_true774
	} else {
		goto if_end778
	}

land_lhs_true774:
	v370 = *lookahead
	cmp775 = v370 <= 57
	if cmp775 {
		goto if_then777
	} else {
		goto if_end778
	}

if_then777:
	*state_addr = 18
	goto next_state

if_end778:
	v371 = *lookahead
	cmp779 = v371 != 0
	if cmp779 {
		goto land_lhs_true781
	} else {
		goto if_end791
	}

land_lhs_true781:
	v372 = *lookahead
	cmp782 = v372 < 9
	if cmp782 {
		goto land_lhs_true787
	} else {
		goto lor_lhs_false784
	}

lor_lhs_false784:
	v373 = *lookahead
	cmp785 = 13 < v373
	if cmp785 {
		goto land_lhs_true787
	} else {
		goto if_end791
	}

land_lhs_true787:
	v374 = *lookahead
	cmp788 = v374 != 124
	if cmp788 {
		goto if_then790
	} else {
		goto if_end791
	}

if_then790:
	*state_addr = 42
	goto next_state

if_end791:
	v375 = *result
	tobool792 = (v375 & 1) != 0
	*retval = tobool792
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v376 = *retval
	return v376
}

