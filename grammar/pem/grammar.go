package grammar_pem

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

var tree_sitter_pem_language TSLanguage = TSLanguage{15, 15, 0, 8, 0, 20, 2, 1, 0, 5, &(*[2][15]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[62]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{0, 3, 0}}

var ts_small_parse_table [127]int16 = [127]int16{
	5, 5, 1, 6, 9, 1, 0, 11, 1, 7, 4, 1, 10, 3, 2, 9,
	13, 5, 13, 1, 0, 15, 1, 6, 18, 1, 7, 4, 1, 10, 3, 2,
	9, 13, 3, 21, 1, 4, 5, 1, 14, 9, 1, 12, 3, 23, 1, 4,
	25, 1, 6, 7, 1, 14, 2, 27, 1, 0, 29, 2, 6, 7, 3, 31,
	1, 4, 34, 1, 6, 7, 1, 14, 2, 36, 1, 0, 38, 2, 6, 7,
	2, 40, 1, 6, 6, 1, 11, 1, 42, 1, 1, 1, 44, 1, 0, 1,
	46, 1, 2, 1, 48, 1, 5, 1, 50, 1, 3, 1, 52, 1, 6, 1,
	54, 1, 2, 1, 56, 1, 4, 1, 58, 1, 5, 1, 60, 1, 6,
}

var ts_small_parse_table_map [18]int32 = [18]int32{
	0, 17, 34, 44, 54, 62, 72, 80, 87, 91, 95, 99, 103, 107, 111, 115,
	119, 123,
}

var ts_symbol_names [15]*byte = [15]*byte{&_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0]}

var ts_symbol_metadata [15]TSSymbolMetadata = [15]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{}}

var ts_symbol_map [15]int16 = [15]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [1][5]int16 = [1][5]int16{}

var ts_lex_modes [20]TSLexerMode = [20]TSLexerMode{
	TSLexerMode{}, TSLexerMode{18, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{3, 0, 0}, TSLexerMode{18, 0, 0}, TSLexerMode{}, TSLexerMode{5, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{6, 0, 0}, TSLexerMode{5, 0, 0}, TSLexerMode{},
	TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{6, 0, 0}, TSLexerMode{},
}

var ts_primary_state_ids [20]int16 = [20]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19,
}

var _str [4]byte = [4]byte{112, 101, 109, 0}

var ts_parse_table struct {
	F0 struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 int16
	F7 [8]int16
}
	F1 [15]int16
} = struct {
	F0 struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 int16
	F7 [8]int16
}
	F1 [15]int16
}{struct {
	F0 int16
	F1 int16
	F2 int16
	F3 int16
	F4 int16
	F5 int16
	F6 int16
	F7 [8]int16
}{1, 1, 1, 1, 1, 0, 1, [8]int16{}}, [15]int16{3, 0, 0, 0, 0, 0, 5, 7, 11, 2, 4, 0, 0, 2, 0}}

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
	F10 TSParseActionEntry
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F18 struct {
	F0 anon_1
	F1 [6]byte
}
	F19 TSParseActionEntry
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
	F24 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F30 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F37 TSParseActionEntry
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
	F43 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F44 struct {
	F0 anon_1
	F1 [6]byte
}
	F45 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F46 struct {
	F0 anon_1
	F1 [6]byte
}
	F47 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F48 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F57 TSParseActionEntry
	F58 struct {
	F0 anon_1
	F1 [6]byte
}
	F59 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
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
	F10 TSParseActionEntry
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
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F18 struct {
	F0 anon_1
	F1 [6]byte
}
	F19 TSParseActionEntry
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
	F24 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
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
	F30 TSParseActionEntry
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
	F0 anon_1
	F1 [6]byte
}
	F37 TSParseActionEntry
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
	F43 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F44 struct {
	F0 anon_1
	F1 [6]byte
}
	F45 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
	F46 struct {
	F0 anon_1
	F1 [6]byte
}
	F47 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F48 struct {
	F0 anon_1
	F1 [6]byte
}
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
	F57 TSParseActionEntry
	F58 struct {
	F0 anon_1
	F1 [6]byte
}
	F59 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 8, 0, 0}}}, struct {
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
}{0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 8, 0, 0}}}, struct {
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
}{0, 3, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 13, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 13, 0, 0}}}, struct {
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
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 13, 0, 0}}}, struct {
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
}{0, 7, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 12, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 9, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 9, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 14, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 14, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 11, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 11, 0, 0}}}, struct {
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
}{0, 18, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 5, 10, 0, 0}}}, struct {
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
}{0, 8, 0, 0}, [2]byte{}}}}

var _str_3 [4]byte = [4]byte{101, 110, 100, 0}

var _str_4 [6]byte = [6]byte{66, 69, 71, 73, 78, 0}

var _str_5 [2]byte = [2]byte{32, 0}

var _str_6 [4]byte = [4]byte{69, 78, 68, 0}

var _str_7 [12]byte = [12]byte{100, 97, 116, 97, 95, 116, 111, 107, 101, 110, 49, 0}

var _str_8 [6]byte = [6]byte{108, 97, 98, 101, 108, 0}

var _str_9 [7]byte = [7]byte{100, 97, 115, 104, 101, 115, 0}

var _str_10 [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}

var _str_11 [8]byte = [8]byte{99, 111, 110, 116, 101, 110, 116, 0}

var _str_12 [7]byte = [7]byte{104, 101, 97, 100, 101, 114, 0}

var _str_13 [7]byte = [7]byte{102, 111, 111, 116, 101, 114, 0}

var _str_14 [5]byte = [5]byte{100, 97, 116, 97, 0}

var _str_15 [12]byte = [12]byte{112, 101, 109, 95, 114, 101, 112, 101, 97, 116, 49, 0}

var _str_16 [13]byte = [13]byte{100, 97, 116, 97, 95, 114, 101, 112, 101, 97, 116, 49, 0}

func tree_sitter_pem() *TSLanguage {
	return &tree_sitter_pem_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v84, v85, v87, v89, v90, v92, v94, v95, v97, v107, v108, v110, v112, v113, v115, v117, v118, v120, v130, v131, v133, v135, v136, v138, v149, v150, v152, v163, v164, v166, v177, v178, v180, v191, v192, v194, v205, v206, v208, v219, v220, v222, v232, v233, v235, v238, v239, v241, v247, v248, v250, v255, v256, v258, v262, v263, v265, v267, v268, v270, v274, v275, v277, v282, v283, v285, v290, v291, v293, v298, v299, v301, v306, v307, v309 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end243, mark_end247, mark_end277, mark_end281, mark_end285, mark_end315, mark_end319, mark_end353, mark_end387, mark_end421, mark_end455, mark_end489, mark_end523, mark_end553, mark_end561, mark_end580, mark_end595, mark_end606, mark_end610, mark_end622, mark_end637, mark_end652, mark_end667, mark_end682 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, result_symbol, result_symbol242, result_symbol246, result_symbol276, result_symbol280, result_symbol284, result_symbol314, result_symbol318, result_symbol352, result_symbol386, result_symbol420, result_symbol454, result_symbol488, result_symbol522, result_symbol552, result_symbol560, result_symbol579, result_symbol594, result_symbol605, result_symbol609, result_symbol621, result_symbol636, result_symbol651, result_symbol666, result_symbol681 *int16
	var lookahead, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp27, cmp29, cmp31, cmp34, cmp37, cmp40, cmp43, tobool47, cmp49, cmp53, cmp57, cmp61, tobool65, cmp67, tobool71, cmp73, cmp77, cmp81, cmp85, cmp88, cmp91, cmp94, cmp97, cmp100, cmp103, tobool107, cmp109, tobool113, cmp115, cmp119, cmp123, cmp127, tobool131, cmp133, cmp137, cmp141, cmp144, tobool148, cmp150, tobool154, cmp156, tobool160, cmp162, tobool166, cmp168, tobool172, cmp174, tobool178, cmp180, tobool184, cmp186, tobool190, cmp192, tobool196, cmp198, tobool202, cmp204, tobool208, tobool210, cmp213, tobool217, tobool219, cmp222, cmp226, cmp230, cmp234, tobool238, tobool240, tobool244, cmp248, cmp252, cmp255, cmp258, cmp261, cmp264, cmp267, cmp270, tobool274, tobool278, tobool282, cmp286, cmp290, cmp293, cmp296, cmp299, cmp302, cmp305, cmp308, tobool312, tobool316, cmp320, cmp324, cmp328, cmp331, cmp334, cmp337, cmp340, cmp343, cmp346, tobool350, cmp354, cmp358, cmp362, cmp365, cmp368, cmp371, cmp374, cmp377, cmp380, tobool384, cmp388, cmp392, cmp396, cmp399, cmp402, cmp405, cmp408, cmp411, cmp414, tobool418, cmp422, cmp426, cmp430, cmp433, cmp436, cmp439, cmp442, cmp445, cmp448, tobool452, cmp456, cmp460, cmp464, cmp467, cmp470, cmp473, cmp476, cmp479, cmp482, tobool486, cmp490, cmp494, cmp498, cmp501, cmp504, cmp507, cmp510, cmp513, cmp516, tobool520, cmp524, cmp528, cmp531, cmp534, cmp537, cmp540, cmp543, cmp546, tobool550, cmp554, tobool558, cmp562, cmp566, cmp570, cmp573, tobool577, cmp581, cmp585, cmp588, tobool592, cmp596, cmp599, tobool603, tobool607, cmp611, cmp615, tobool619, cmp623, cmp627, cmp630, tobool634, cmp638, cmp642, cmp645, tobool649, cmp653, cmp657, cmp660, tobool664, cmp668, cmp672, cmp675, tobool679, cmp683, cmp686, tobool690, v313 bool
	var v3, frombool, v10, v24, v29, v31, v42, v44, v49, v54, v56, v58, v60, v62, v64, v66, v68, v70, v72, v74, v75, v77, v78, v83, v88, v93, v106, v111, v116, v129, v134, v148, v162, v176, v190, v204, v218, v231, v237, v246, v254, v261, v266, v273, v281, v289, v297, v305, v312 byte
	var v86, v91, v96, v109, v114, v119, v132, v137, v151, v165, v179, v193, v207, v221, v234, v240, v249, v257, v264, v269, v276, v284, v292, v300, v308 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9 int16
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v25, v26, v27, v28, v30, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v43, v45, v46, v47, v48, v50, v51, v52, v53, v55, v57, v59, v61, v63, v65, v67, v69, v71, v73, v76, v79, v80, v81, v82, v98, v99, v100, v101, v102, v103, v104, v105, v121, v122, v123, v124, v125, v126, v127, v128, v139, v140, v141, v142, v143, v144, v145, v146, v147, v153, v154, v155, v156, v157, v158, v159, v160, v161, v167, v168, v169, v170, v171, v172, v173, v174, v175, v181, v182, v183, v184, v185, v186, v187, v188, v189, v195, v196, v197, v198, v199, v200, v201, v202, v203, v209, v210, v211, v212, v213, v214, v215, v216, v217, v223, v224, v225, v226, v227, v228, v229, v230, v236, v242, v243, v244, v245, v251, v252, v253, v259, v260, v271, v272, v278, v279, v280, v286, v287, v288, v294, v295, v296, v302, v303, v304, v310, v311 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp27, v18, cmp29, v19, cmp31, v20, cmp34, v21, cmp37, v22, cmp40, v23, cmp43, v24, tobool47, v25, cmp49, v26, cmp53, v27, cmp57, v28, cmp61, v29, tobool65, v30, cmp67, v31, tobool71, v32, cmp73, v33, cmp77, v34, cmp81, v35, cmp85, v36, cmp88, v37, cmp91, v38, cmp94, v39, cmp97, v40, cmp100, v41, cmp103, v42, tobool107, v43, cmp109, v44, tobool113, v45, cmp115, v46, cmp119, v47, cmp123, v48, cmp127, v49, tobool131, v50, cmp133, v51, cmp137, v52, cmp141, v53, cmp144, v54, tobool148, v55, cmp150, v56, tobool154, v57, cmp156, v58, tobool160, v59, cmp162, v60, tobool166, v61, cmp168, v62, tobool172, v63, cmp174, v64, tobool178, v65, cmp180, v66, tobool184, v67, cmp186, v68, tobool190, v69, cmp192, v70, tobool196, v71, cmp198, v72, tobool202, v73, cmp204, v74, tobool208, v75, tobool210, v76, cmp213, v77, tobool217, v78, tobool219, v79, cmp222, v80, cmp226, v81, cmp230, v82, cmp234, v83, tobool238, v84, result_symbol, v85, mark_end, v86, v87, v88, tobool240, v89, result_symbol242, v90, mark_end243, v91, v92, v93, tobool244, v94, result_symbol246, v95, mark_end247, v96, v97, v98, cmp248, v99, cmp252, v100, cmp255, v101, cmp258, v102, cmp261, v103, cmp264, v104, cmp267, v105, cmp270, v106, tobool274, v107, result_symbol276, v108, mark_end277, v109, v110, v111, tobool278, v112, result_symbol280, v113, mark_end281, v114, v115, v116, tobool282, v117, result_symbol284, v118, mark_end285, v119, v120, v121, cmp286, v122, cmp290, v123, cmp293, v124, cmp296, v125, cmp299, v126, cmp302, v127, cmp305, v128, cmp308, v129, tobool312, v130, result_symbol314, v131, mark_end315, v132, v133, v134, tobool316, v135, result_symbol318, v136, mark_end319, v137, v138, v139, cmp320, v140, cmp324, v141, cmp328, v142, cmp331, v143, cmp334, v144, cmp337, v145, cmp340, v146, cmp343, v147, cmp346, v148, tobool350, v149, result_symbol352, v150, mark_end353, v151, v152, v153, cmp354, v154, cmp358, v155, cmp362, v156, cmp365, v157, cmp368, v158, cmp371, v159, cmp374, v160, cmp377, v161, cmp380, v162, tobool384, v163, result_symbol386, v164, mark_end387, v165, v166, v167, cmp388, v168, cmp392, v169, cmp396, v170, cmp399, v171, cmp402, v172, cmp405, v173, cmp408, v174, cmp411, v175, cmp414, v176, tobool418, v177, result_symbol420, v178, mark_end421, v179, v180, v181, cmp422, v182, cmp426, v183, cmp430, v184, cmp433, v185, cmp436, v186, cmp439, v187, cmp442, v188, cmp445, v189, cmp448, v190, tobool452, v191, result_symbol454, v192, mark_end455, v193, v194, v195, cmp456, v196, cmp460, v197, cmp464, v198, cmp467, v199, cmp470, v200, cmp473, v201, cmp476, v202, cmp479, v203, cmp482, v204, tobool486, v205, result_symbol488, v206, mark_end489, v207, v208, v209, cmp490, v210, cmp494, v211, cmp498, v212, cmp501, v213, cmp504, v214, cmp507, v215, cmp510, v216, cmp513, v217, cmp516, v218, tobool520, v219, result_symbol522, v220, mark_end523, v221, v222, v223, cmp524, v224, cmp528, v225, cmp531, v226, cmp534, v227, cmp537, v228, cmp540, v229, cmp543, v230, cmp546, v231, tobool550, v232, result_symbol552, v233, mark_end553, v234, v235, v236, cmp554, v237, tobool558, v238, result_symbol560, v239, mark_end561, v240, v241, v242, cmp562, v243, cmp566, v244, cmp570, v245, cmp573, v246, tobool577, v247, result_symbol579, v248, mark_end580, v249, v250, v251, cmp581, v252, cmp585, v253, cmp588, v254, tobool592, v255, result_symbol594, v256, mark_end595, v257, v258, v259, cmp596, v260, cmp599, v261, tobool603, v262, result_symbol605, v263, mark_end606, v264, v265, v266, tobool607, v267, result_symbol609, v268, mark_end610, v269, v270, v271, cmp611, v272, cmp615, v273, tobool619, v274, result_symbol621, v275, mark_end622, v276, v277, v278, cmp623, v279, cmp627, v280, cmp630, v281, tobool634, v282, result_symbol636, v283, mark_end637, v284, v285, v286, cmp638, v287, cmp642, v288, cmp645, v289, tobool649, v290, result_symbol651, v291, mark_end652, v292, v293, v294, cmp653, v295, cmp657, v296, cmp660, v297, tobool664, v298, result_symbol666, v299, mark_end667, v300, v301, v302, cmp668, v303, cmp672, v304, cmp675, v305, tobool679, v306, result_symbol681, v307, mark_end682, v308, v309, v310, cmp683, v311, cmp686, v312, tobool690, v313

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
		goto sw_bb48
	case 2:
		goto sw_bb66
	case 3:
		goto sw_bb72
	case 4:
		goto sw_bb108
	case 5:
		goto sw_bb114
	case 6:
		goto sw_bb132
	case 7:
		goto sw_bb149
	case 8:
		goto sw_bb155
	case 9:
		goto sw_bb161
	case 10:
		goto sw_bb167
	case 11:
		goto sw_bb173
	case 12:
		goto sw_bb179
	case 13:
		goto sw_bb185
	case 14:
		goto sw_bb191
	case 15:
		goto sw_bb197
	case 16:
		goto sw_bb203
	case 17:
		goto sw_bb209
	case 18:
		goto sw_bb218
	case 19:
		goto sw_bb239
	case 20:
		goto sw_bb241
	case 21:
		goto sw_bb245
	case 22:
		goto sw_bb275
	case 23:
		goto sw_bb279
	case 24:
		goto sw_bb283
	case 25:
		goto sw_bb313
	case 26:
		goto sw_bb317
	case 27:
		goto sw_bb351
	case 28:
		goto sw_bb385
	case 29:
		goto sw_bb419
	case 30:
		goto sw_bb453
	case 31:
		goto sw_bb487
	case 32:
		goto sw_bb521
	case 33:
		goto sw_bb551
	case 34:
		goto sw_bb559
	case 35:
		goto sw_bb578
	case 36:
		goto sw_bb593
	case 37:
		goto sw_bb604
	case 38:
		goto sw_bb608
	case 39:
		goto sw_bb620
	case 40:
		goto sw_bb635
	case 41:
		goto sw_bb650
	case 42:
		goto sw_bb665
	case 43:
		goto sw_bb680
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
	*state_addr = 19
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
	*skip = 1
	*state_addr = 0
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
	*skip = 1
	*state_addr = 17
	goto next_state

if_end10:
	v13 = *lookahead
	cmp11 = v13 == 32
	if cmp11 {
		goto if_then13
	} else {
		goto if_end14
	}

if_then13:
	*state_addr = 22
	goto next_state

if_end14:
	v14 = *lookahead
	cmp15 = v14 == 45
	if cmp15 {
		goto if_then17
	} else {
		goto if_end18
	}

if_then17:
	*state_addr = 10
	goto next_state

if_end18:
	v15 = *lookahead
	cmp19 = v15 == 66
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*state_addr = 27
	goto next_state

if_end22:
	v16 = *lookahead
	cmp23 = v16 == 69
	if cmp23 {
		goto if_then25
	} else {
		goto if_end26
	}

if_then25:
	*state_addr = 30
	goto next_state

if_end26:
	v17 = *lookahead
	cmp27 = v17 == 43
	if cmp27 {
		goto if_then45
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v18 = *lookahead
	cmp29 = 47 <= v18
	if cmp29 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false33
	}

land_lhs_true:
	v19 = *lookahead
	cmp31 = v19 <= 57
	if cmp31 {
		goto if_then45
	} else {
		goto lor_lhs_false33
	}

lor_lhs_false33:
	v20 = *lookahead
	cmp34 = 65 <= v20
	if cmp34 {
		goto land_lhs_true36
	} else {
		goto lor_lhs_false39
	}

land_lhs_true36:
	v21 = *lookahead
	cmp37 = v21 <= 90
	if cmp37 {
		goto if_then45
	} else {
		goto lor_lhs_false39
	}

lor_lhs_false39:
	v22 = *lookahead
	cmp40 = 97 <= v22
	if cmp40 {
		goto land_lhs_true42
	} else {
		goto if_end46
	}

land_lhs_true42:
	v23 = *lookahead
	cmp43 = v23 <= 122
	if cmp43 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*state_addr = 32
	goto next_state

if_end46:
	v24 = *result
	tobool47 = (v24 & 1) != 0
	*retval = tobool47
	goto _return

sw_bb48:
	v25 = *lookahead
	cmp49 = v25 == 10
	if cmp49 {
		goto if_then51
	} else {
		goto if_end52
	}

if_then51:
	*skip = 1
	*state_addr = 1
	goto next_state

if_end52:
	v26 = *lookahead
	cmp53 = v26 == 13
	if cmp53 {
		goto if_then55
	} else {
		goto if_end56
	}

if_then55:
	*state_addr = 38
	goto next_state

if_end56:
	v27 = *lookahead
	cmp57 = v27 == 45
	if cmp57 {
		goto if_then59
	} else {
		goto if_end60
	}

if_then59:
	*state_addr = 42
	goto next_state

if_end60:
	v28 = *lookahead
	cmp61 = v28 != 0
	if cmp61 {
		goto if_then63
	} else {
		goto if_end64
	}

if_then63:
	*state_addr = 43
	goto next_state

if_end64:
	v29 = *result
	tobool65 = (v29 & 1) != 0
	*retval = tobool65
	goto _return

sw_bb66:
	v30 = *lookahead
	cmp67 = v30 == 10
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*skip = 1
	*state_addr = 3
	goto next_state

if_end70:
	v31 = *result
	tobool71 = (v31 & 1) != 0
	*retval = tobool71
	goto _return

sw_bb72:
	v32 = *lookahead
	cmp73 = v32 == 10
	if cmp73 {
		goto if_then75
	} else {
		goto if_end76
	}

if_then75:
	*skip = 1
	*state_addr = 3
	goto next_state

if_end76:
	v33 = *lookahead
	cmp77 = v33 == 13
	if cmp77 {
		goto if_then79
	} else {
		goto if_end80
	}

if_then79:
	*skip = 1
	*state_addr = 2
	goto next_state

if_end80:
	v34 = *lookahead
	cmp81 = v34 == 45
	if cmp81 {
		goto if_then83
	} else {
		goto if_end84
	}

if_then83:
	*state_addr = 10
	goto next_state

if_end84:
	v35 = *lookahead
	cmp85 = v35 == 43
	if cmp85 {
		goto if_then105
	} else {
		goto lor_lhs_false87
	}

lor_lhs_false87:
	v36 = *lookahead
	cmp88 = 47 <= v36
	if cmp88 {
		goto land_lhs_true90
	} else {
		goto lor_lhs_false93
	}

land_lhs_true90:
	v37 = *lookahead
	cmp91 = v37 <= 57
	if cmp91 {
		goto if_then105
	} else {
		goto lor_lhs_false93
	}

lor_lhs_false93:
	v38 = *lookahead
	cmp94 = 65 <= v38
	if cmp94 {
		goto land_lhs_true96
	} else {
		goto lor_lhs_false99
	}

land_lhs_true96:
	v39 = *lookahead
	cmp97 = v39 <= 90
	if cmp97 {
		goto if_then105
	} else {
		goto lor_lhs_false99
	}

lor_lhs_false99:
	v40 = *lookahead
	cmp100 = 97 <= v40
	if cmp100 {
		goto land_lhs_true102
	} else {
		goto if_end106
	}

land_lhs_true102:
	v41 = *lookahead
	cmp103 = v41 <= 122
	if cmp103 {
		goto if_then105
	} else {
		goto if_end106
	}

if_then105:
	*state_addr = 32
	goto next_state

if_end106:
	v42 = *result
	tobool107 = (v42 & 1) != 0
	*retval = tobool107
	goto _return

sw_bb108:
	v43 = *lookahead
	cmp109 = v43 == 10
	if cmp109 {
		goto if_then111
	} else {
		goto if_end112
	}

if_then111:
	*skip = 1
	*state_addr = 5
	goto next_state

if_end112:
	v44 = *result
	tobool113 = (v44 & 1) != 0
	*retval = tobool113
	goto _return

sw_bb114:
	v45 = *lookahead
	cmp115 = v45 == 10
	if cmp115 {
		goto if_then117
	} else {
		goto if_end118
	}

if_then117:
	*skip = 1
	*state_addr = 5
	goto next_state

if_end118:
	v46 = *lookahead
	cmp119 = v46 == 13
	if cmp119 {
		goto if_then121
	} else {
		goto if_end122
	}

if_then121:
	*skip = 1
	*state_addr = 4
	goto next_state

if_end122:
	v47 = *lookahead
	cmp123 = v47 == 66
	if cmp123 {
		goto if_then125
	} else {
		goto if_end126
	}

if_then125:
	*state_addr = 12
	goto next_state

if_end126:
	v48 = *lookahead
	cmp127 = v48 == 69
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*state_addr = 15
	goto next_state

if_end130:
	v49 = *result
	tobool131 = (v49 & 1) != 0
	*retval = tobool131
	goto _return

sw_bb132:
	v50 = *lookahead
	cmp133 = v50 == 10
	if cmp133 {
		goto if_then135
	} else {
		goto if_end136
	}

if_then135:
	*state_addr = 34
	goto next_state

if_end136:
	v51 = *lookahead
	cmp137 = v51 == 13
	if cmp137 {
		goto if_then139
	} else {
		goto if_end140
	}

if_then139:
	*state_addr = 35
	goto next_state

if_end140:
	v52 = *lookahead
	cmp141 = v52 != 0
	if cmp141 {
		goto land_lhs_true143
	} else {
		goto if_end147
	}

land_lhs_true143:
	v53 = *lookahead
	cmp144 = v53 != 45
	if cmp144 {
		goto if_then146
	} else {
		goto if_end147
	}

if_then146:
	*state_addr = 36
	goto next_state

if_end147:
	v54 = *result
	tobool148 = (v54 & 1) != 0
	*retval = tobool148
	goto _return

sw_bb149:
	v55 = *lookahead
	cmp150 = v55 == 45
	if cmp150 {
		goto if_then152
	} else {
		goto if_end153
	}

if_then152:
	*state_addr = 37
	goto next_state

if_end153:
	v56 = *result
	tobool154 = (v56 & 1) != 0
	*retval = tobool154
	goto _return

sw_bb155:
	v57 = *lookahead
	cmp156 = v57 == 45
	if cmp156 {
		goto if_then158
	} else {
		goto if_end159
	}

if_then158:
	*state_addr = 7
	goto next_state

if_end159:
	v58 = *result
	tobool160 = (v58 & 1) != 0
	*retval = tobool160
	goto _return

sw_bb161:
	v59 = *lookahead
	cmp162 = v59 == 45
	if cmp162 {
		goto if_then164
	} else {
		goto if_end165
	}

if_then164:
	*state_addr = 8
	goto next_state

if_end165:
	v60 = *result
	tobool166 = (v60 & 1) != 0
	*retval = tobool166
	goto _return

sw_bb167:
	v61 = *lookahead
	cmp168 = v61 == 45
	if cmp168 {
		goto if_then170
	} else {
		goto if_end171
	}

if_then170:
	*state_addr = 9
	goto next_state

if_end171:
	v62 = *result
	tobool172 = (v62 & 1) != 0
	*retval = tobool172
	goto _return

sw_bb173:
	v63 = *lookahead
	cmp174 = v63 == 68
	if cmp174 {
		goto if_then176
	} else {
		goto if_end177
	}

if_then176:
	*state_addr = 23
	goto next_state

if_end177:
	v64 = *result
	tobool178 = (v64 & 1) != 0
	*retval = tobool178
	goto _return

sw_bb179:
	v65 = *lookahead
	cmp180 = v65 == 69
	if cmp180 {
		goto if_then182
	} else {
		goto if_end183
	}

if_then182:
	*state_addr = 13
	goto next_state

if_end183:
	v66 = *result
	tobool184 = (v66 & 1) != 0
	*retval = tobool184
	goto _return

sw_bb185:
	v67 = *lookahead
	cmp186 = v67 == 71
	if cmp186 {
		goto if_then188
	} else {
		goto if_end189
	}

if_then188:
	*state_addr = 14
	goto next_state

if_end189:
	v68 = *result
	tobool190 = (v68 & 1) != 0
	*retval = tobool190
	goto _return

sw_bb191:
	v69 = *lookahead
	cmp192 = v69 == 73
	if cmp192 {
		goto if_then194
	} else {
		goto if_end195
	}

if_then194:
	*state_addr = 16
	goto next_state

if_end195:
	v70 = *result
	tobool196 = (v70 & 1) != 0
	*retval = tobool196
	goto _return

sw_bb197:
	v71 = *lookahead
	cmp198 = v71 == 78
	if cmp198 {
		goto if_then200
	} else {
		goto if_end201
	}

if_then200:
	*state_addr = 11
	goto next_state

if_end201:
	v72 = *result
	tobool202 = (v72 & 1) != 0
	*retval = tobool202
	goto _return

sw_bb203:
	v73 = *lookahead
	cmp204 = v73 == 78
	if cmp204 {
		goto if_then206
	} else {
		goto if_end207
	}

if_then206:
	*state_addr = 20
	goto next_state

if_end207:
	v74 = *result
	tobool208 = (v74 & 1) != 0
	*retval = tobool208
	goto _return

sw_bb209:
	v75 = *eof
	tobool210 = (v75 & 1) != 0
	if tobool210 {
		goto if_then211
	} else {
		goto if_end212
	}

if_then211:
	*state_addr = 19
	goto next_state

if_end212:
	v76 = *lookahead
	cmp213 = v76 == 10
	if cmp213 {
		goto if_then215
	} else {
		goto if_end216
	}

if_then215:
	*skip = 1
	*state_addr = 0
	goto next_state

if_end216:
	v77 = *result
	tobool217 = (v77 & 1) != 0
	*retval = tobool217
	goto _return

sw_bb218:
	v78 = *eof
	tobool219 = (v78 & 1) != 0
	if tobool219 {
		goto if_then220
	} else {
		goto if_end221
	}

if_then220:
	*state_addr = 19
	goto next_state

if_end221:
	v79 = *lookahead
	cmp222 = v79 == 10
	if cmp222 {
		goto if_then224
	} else {
		goto if_end225
	}

if_then224:
	*skip = 1
	*state_addr = 18
	goto next_state

if_end225:
	v80 = *lookahead
	cmp226 = v80 == 13
	if cmp226 {
		goto if_then228
	} else {
		goto if_end229
	}

if_then228:
	*state_addr = 38
	goto next_state

if_end229:
	v81 = *lookahead
	cmp230 = v81 == 45
	if cmp230 {
		goto if_then232
	} else {
		goto if_end233
	}

if_then232:
	*state_addr = 42
	goto next_state

if_end233:
	v82 = *lookahead
	cmp234 = v82 != 0
	if cmp234 {
		goto if_then236
	} else {
		goto if_end237
	}

if_then236:
	*state_addr = 43
	goto next_state

if_end237:
	v83 = *result
	tobool238 = (v83 & 1) != 0
	*retval = tobool238
	goto _return

sw_bb239:
	*result = 1
	v84 = *lexer_addr
	result_symbol = &v84.F1
	*result_symbol = 0
	v85 = *lexer_addr
	mark_end = &v85.F3
	v86 = *mark_end
	v87 = *lexer_addr
	v86(v87)
	v88 = *result
	tobool240 = (v88 & 1) != 0
	*retval = tobool240
	goto _return

sw_bb241:
	*result = 1
	v89 = *lexer_addr
	result_symbol242 = &v89.F1
	*result_symbol242 = 1
	v90 = *lexer_addr
	mark_end243 = &v90.F3
	v91 = *mark_end243
	v92 = *lexer_addr
	v91(v92)
	v93 = *result
	tobool244 = (v93 & 1) != 0
	*retval = tobool244
	goto _return

sw_bb245:
	*result = 1
	v94 = *lexer_addr
	result_symbol246 = &v94.F1
	*result_symbol246 = 1
	v95 = *lexer_addr
	mark_end247 = &v95.F3
	v96 = *mark_end247
	v97 = *lexer_addr
	v96(v97)
	v98 = *lookahead
	cmp248 = v98 == 61
	if cmp248 {
		goto if_then250
	} else {
		goto if_end251
	}

if_then250:
	*state_addr = 33
	goto next_state

if_end251:
	v99 = *lookahead
	cmp252 = v99 == 43
	if cmp252 {
		goto if_then272
	} else {
		goto lor_lhs_false254
	}

lor_lhs_false254:
	v100 = *lookahead
	cmp255 = 47 <= v100
	if cmp255 {
		goto land_lhs_true257
	} else {
		goto lor_lhs_false260
	}

land_lhs_true257:
	v101 = *lookahead
	cmp258 = v101 <= 57
	if cmp258 {
		goto if_then272
	} else {
		goto lor_lhs_false260
	}

lor_lhs_false260:
	v102 = *lookahead
	cmp261 = 65 <= v102
	if cmp261 {
		goto land_lhs_true263
	} else {
		goto lor_lhs_false266
	}

land_lhs_true263:
	v103 = *lookahead
	cmp264 = v103 <= 90
	if cmp264 {
		goto if_then272
	} else {
		goto lor_lhs_false266
	}

lor_lhs_false266:
	v104 = *lookahead
	cmp267 = 97 <= v104
	if cmp267 {
		goto land_lhs_true269
	} else {
		goto if_end273
	}

land_lhs_true269:
	v105 = *lookahead
	cmp270 = v105 <= 122
	if cmp270 {
		goto if_then272
	} else {
		goto if_end273
	}

if_then272:
	*state_addr = 32
	goto next_state

if_end273:
	v106 = *result
	tobool274 = (v106 & 1) != 0
	*retval = tobool274
	goto _return

sw_bb275:
	*result = 1
	v107 = *lexer_addr
	result_symbol276 = &v107.F1
	*result_symbol276 = 2
	v108 = *lexer_addr
	mark_end277 = &v108.F3
	v109 = *mark_end277
	v110 = *lexer_addr
	v109(v110)
	v111 = *result
	tobool278 = (v111 & 1) != 0
	*retval = tobool278
	goto _return

sw_bb279:
	*result = 1
	v112 = *lexer_addr
	result_symbol280 = &v112.F1
	*result_symbol280 = 3
	v113 = *lexer_addr
	mark_end281 = &v113.F3
	v114 = *mark_end281
	v115 = *lexer_addr
	v114(v115)
	v116 = *result
	tobool282 = (v116 & 1) != 0
	*retval = tobool282
	goto _return

sw_bb283:
	*result = 1
	v117 = *lexer_addr
	result_symbol284 = &v117.F1
	*result_symbol284 = 3
	v118 = *lexer_addr
	mark_end285 = &v118.F3
	v119 = *mark_end285
	v120 = *lexer_addr
	v119(v120)
	v121 = *lookahead
	cmp286 = v121 == 61
	if cmp286 {
		goto if_then288
	} else {
		goto if_end289
	}

if_then288:
	*state_addr = 33
	goto next_state

if_end289:
	v122 = *lookahead
	cmp290 = v122 == 43
	if cmp290 {
		goto if_then310
	} else {
		goto lor_lhs_false292
	}

lor_lhs_false292:
	v123 = *lookahead
	cmp293 = 47 <= v123
	if cmp293 {
		goto land_lhs_true295
	} else {
		goto lor_lhs_false298
	}

land_lhs_true295:
	v124 = *lookahead
	cmp296 = v124 <= 57
	if cmp296 {
		goto if_then310
	} else {
		goto lor_lhs_false298
	}

lor_lhs_false298:
	v125 = *lookahead
	cmp299 = 65 <= v125
	if cmp299 {
		goto land_lhs_true301
	} else {
		goto lor_lhs_false304
	}

land_lhs_true301:
	v126 = *lookahead
	cmp302 = v126 <= 90
	if cmp302 {
		goto if_then310
	} else {
		goto lor_lhs_false304
	}

lor_lhs_false304:
	v127 = *lookahead
	cmp305 = 97 <= v127
	if cmp305 {
		goto land_lhs_true307
	} else {
		goto if_end311
	}

land_lhs_true307:
	v128 = *lookahead
	cmp308 = v128 <= 122
	if cmp308 {
		goto if_then310
	} else {
		goto if_end311
	}

if_then310:
	*state_addr = 32
	goto next_state

if_end311:
	v129 = *result
	tobool312 = (v129 & 1) != 0
	*retval = tobool312
	goto _return

sw_bb313:
	*result = 1
	v130 = *lexer_addr
	result_symbol314 = &v130.F1
	*result_symbol314 = 4
	v131 = *lexer_addr
	mark_end315 = &v131.F3
	v132 = *mark_end315
	v133 = *lexer_addr
	v132(v133)
	v134 = *result
	tobool316 = (v134 & 1) != 0
	*retval = tobool316
	goto _return

sw_bb317:
	*result = 1
	v135 = *lexer_addr
	result_symbol318 = &v135.F1
	*result_symbol318 = 4
	v136 = *lexer_addr
	mark_end319 = &v136.F3
	v137 = *mark_end319
	v138 = *lexer_addr
	v137(v138)
	v139 = *lookahead
	cmp320 = v139 == 61
	if cmp320 {
		goto if_then322
	} else {
		goto if_end323
	}

if_then322:
	*state_addr = 33
	goto next_state

if_end323:
	v140 = *lookahead
	cmp324 = v140 == 68
	if cmp324 {
		goto if_then326
	} else {
		goto if_end327
	}

if_then326:
	*state_addr = 24
	goto next_state

if_end327:
	v141 = *lookahead
	cmp328 = v141 == 43
	if cmp328 {
		goto if_then348
	} else {
		goto lor_lhs_false330
	}

lor_lhs_false330:
	v142 = *lookahead
	cmp331 = 47 <= v142
	if cmp331 {
		goto land_lhs_true333
	} else {
		goto lor_lhs_false336
	}

land_lhs_true333:
	v143 = *lookahead
	cmp334 = v143 <= 57
	if cmp334 {
		goto if_then348
	} else {
		goto lor_lhs_false336
	}

lor_lhs_false336:
	v144 = *lookahead
	cmp337 = 65 <= v144
	if cmp337 {
		goto land_lhs_true339
	} else {
		goto lor_lhs_false342
	}

land_lhs_true339:
	v145 = *lookahead
	cmp340 = v145 <= 90
	if cmp340 {
		goto if_then348
	} else {
		goto lor_lhs_false342
	}

lor_lhs_false342:
	v146 = *lookahead
	cmp343 = 97 <= v146
	if cmp343 {
		goto land_lhs_true345
	} else {
		goto if_end349
	}

land_lhs_true345:
	v147 = *lookahead
	cmp346 = v147 <= 122
	if cmp346 {
		goto if_then348
	} else {
		goto if_end349
	}

if_then348:
	*state_addr = 32
	goto next_state

if_end349:
	v148 = *result
	tobool350 = (v148 & 1) != 0
	*retval = tobool350
	goto _return

sw_bb351:
	*result = 1
	v149 = *lexer_addr
	result_symbol352 = &v149.F1
	*result_symbol352 = 4
	v150 = *lexer_addr
	mark_end353 = &v150.F3
	v151 = *mark_end353
	v152 = *lexer_addr
	v151(v152)
	v153 = *lookahead
	cmp354 = v153 == 61
	if cmp354 {
		goto if_then356
	} else {
		goto if_end357
	}

if_then356:
	*state_addr = 33
	goto next_state

if_end357:
	v154 = *lookahead
	cmp358 = v154 == 69
	if cmp358 {
		goto if_then360
	} else {
		goto if_end361
	}

if_then360:
	*state_addr = 28
	goto next_state

if_end361:
	v155 = *lookahead
	cmp362 = v155 == 43
	if cmp362 {
		goto if_then382
	} else {
		goto lor_lhs_false364
	}

lor_lhs_false364:
	v156 = *lookahead
	cmp365 = 47 <= v156
	if cmp365 {
		goto land_lhs_true367
	} else {
		goto lor_lhs_false370
	}

land_lhs_true367:
	v157 = *lookahead
	cmp368 = v157 <= 57
	if cmp368 {
		goto if_then382
	} else {
		goto lor_lhs_false370
	}

lor_lhs_false370:
	v158 = *lookahead
	cmp371 = 65 <= v158
	if cmp371 {
		goto land_lhs_true373
	} else {
		goto lor_lhs_false376
	}

land_lhs_true373:
	v159 = *lookahead
	cmp374 = v159 <= 90
	if cmp374 {
		goto if_then382
	} else {
		goto lor_lhs_false376
	}

lor_lhs_false376:
	v160 = *lookahead
	cmp377 = 97 <= v160
	if cmp377 {
		goto land_lhs_true379
	} else {
		goto if_end383
	}

land_lhs_true379:
	v161 = *lookahead
	cmp380 = v161 <= 122
	if cmp380 {
		goto if_then382
	} else {
		goto if_end383
	}

if_then382:
	*state_addr = 32
	goto next_state

if_end383:
	v162 = *result
	tobool384 = (v162 & 1) != 0
	*retval = tobool384
	goto _return

sw_bb385:
	*result = 1
	v163 = *lexer_addr
	result_symbol386 = &v163.F1
	*result_symbol386 = 4
	v164 = *lexer_addr
	mark_end387 = &v164.F3
	v165 = *mark_end387
	v166 = *lexer_addr
	v165(v166)
	v167 = *lookahead
	cmp388 = v167 == 61
	if cmp388 {
		goto if_then390
	} else {
		goto if_end391
	}

if_then390:
	*state_addr = 33
	goto next_state

if_end391:
	v168 = *lookahead
	cmp392 = v168 == 71
	if cmp392 {
		goto if_then394
	} else {
		goto if_end395
	}

if_then394:
	*state_addr = 29
	goto next_state

if_end395:
	v169 = *lookahead
	cmp396 = v169 == 43
	if cmp396 {
		goto if_then416
	} else {
		goto lor_lhs_false398
	}

lor_lhs_false398:
	v170 = *lookahead
	cmp399 = 47 <= v170
	if cmp399 {
		goto land_lhs_true401
	} else {
		goto lor_lhs_false404
	}

land_lhs_true401:
	v171 = *lookahead
	cmp402 = v171 <= 57
	if cmp402 {
		goto if_then416
	} else {
		goto lor_lhs_false404
	}

lor_lhs_false404:
	v172 = *lookahead
	cmp405 = 65 <= v172
	if cmp405 {
		goto land_lhs_true407
	} else {
		goto lor_lhs_false410
	}

land_lhs_true407:
	v173 = *lookahead
	cmp408 = v173 <= 90
	if cmp408 {
		goto if_then416
	} else {
		goto lor_lhs_false410
	}

lor_lhs_false410:
	v174 = *lookahead
	cmp411 = 97 <= v174
	if cmp411 {
		goto land_lhs_true413
	} else {
		goto if_end417
	}

land_lhs_true413:
	v175 = *lookahead
	cmp414 = v175 <= 122
	if cmp414 {
		goto if_then416
	} else {
		goto if_end417
	}

if_then416:
	*state_addr = 32
	goto next_state

if_end417:
	v176 = *result
	tobool418 = (v176 & 1) != 0
	*retval = tobool418
	goto _return

sw_bb419:
	*result = 1
	v177 = *lexer_addr
	result_symbol420 = &v177.F1
	*result_symbol420 = 4
	v178 = *lexer_addr
	mark_end421 = &v178.F3
	v179 = *mark_end421
	v180 = *lexer_addr
	v179(v180)
	v181 = *lookahead
	cmp422 = v181 == 61
	if cmp422 {
		goto if_then424
	} else {
		goto if_end425
	}

if_then424:
	*state_addr = 33
	goto next_state

if_end425:
	v182 = *lookahead
	cmp426 = v182 == 73
	if cmp426 {
		goto if_then428
	} else {
		goto if_end429
	}

if_then428:
	*state_addr = 31
	goto next_state

if_end429:
	v183 = *lookahead
	cmp430 = v183 == 43
	if cmp430 {
		goto if_then450
	} else {
		goto lor_lhs_false432
	}

lor_lhs_false432:
	v184 = *lookahead
	cmp433 = 47 <= v184
	if cmp433 {
		goto land_lhs_true435
	} else {
		goto lor_lhs_false438
	}

land_lhs_true435:
	v185 = *lookahead
	cmp436 = v185 <= 57
	if cmp436 {
		goto if_then450
	} else {
		goto lor_lhs_false438
	}

lor_lhs_false438:
	v186 = *lookahead
	cmp439 = 65 <= v186
	if cmp439 {
		goto land_lhs_true441
	} else {
		goto lor_lhs_false444
	}

land_lhs_true441:
	v187 = *lookahead
	cmp442 = v187 <= 90
	if cmp442 {
		goto if_then450
	} else {
		goto lor_lhs_false444
	}

lor_lhs_false444:
	v188 = *lookahead
	cmp445 = 97 <= v188
	if cmp445 {
		goto land_lhs_true447
	} else {
		goto if_end451
	}

land_lhs_true447:
	v189 = *lookahead
	cmp448 = v189 <= 122
	if cmp448 {
		goto if_then450
	} else {
		goto if_end451
	}

if_then450:
	*state_addr = 32
	goto next_state

if_end451:
	v190 = *result
	tobool452 = (v190 & 1) != 0
	*retval = tobool452
	goto _return

sw_bb453:
	*result = 1
	v191 = *lexer_addr
	result_symbol454 = &v191.F1
	*result_symbol454 = 4
	v192 = *lexer_addr
	mark_end455 = &v192.F3
	v193 = *mark_end455
	v194 = *lexer_addr
	v193(v194)
	v195 = *lookahead
	cmp456 = v195 == 61
	if cmp456 {
		goto if_then458
	} else {
		goto if_end459
	}

if_then458:
	*state_addr = 33
	goto next_state

if_end459:
	v196 = *lookahead
	cmp460 = v196 == 78
	if cmp460 {
		goto if_then462
	} else {
		goto if_end463
	}

if_then462:
	*state_addr = 26
	goto next_state

if_end463:
	v197 = *lookahead
	cmp464 = v197 == 43
	if cmp464 {
		goto if_then484
	} else {
		goto lor_lhs_false466
	}

lor_lhs_false466:
	v198 = *lookahead
	cmp467 = 47 <= v198
	if cmp467 {
		goto land_lhs_true469
	} else {
		goto lor_lhs_false472
	}

land_lhs_true469:
	v199 = *lookahead
	cmp470 = v199 <= 57
	if cmp470 {
		goto if_then484
	} else {
		goto lor_lhs_false472
	}

lor_lhs_false472:
	v200 = *lookahead
	cmp473 = 65 <= v200
	if cmp473 {
		goto land_lhs_true475
	} else {
		goto lor_lhs_false478
	}

land_lhs_true475:
	v201 = *lookahead
	cmp476 = v201 <= 90
	if cmp476 {
		goto if_then484
	} else {
		goto lor_lhs_false478
	}

lor_lhs_false478:
	v202 = *lookahead
	cmp479 = 97 <= v202
	if cmp479 {
		goto land_lhs_true481
	} else {
		goto if_end485
	}

land_lhs_true481:
	v203 = *lookahead
	cmp482 = v203 <= 122
	if cmp482 {
		goto if_then484
	} else {
		goto if_end485
	}

if_then484:
	*state_addr = 32
	goto next_state

if_end485:
	v204 = *result
	tobool486 = (v204 & 1) != 0
	*retval = tobool486
	goto _return

sw_bb487:
	*result = 1
	v205 = *lexer_addr
	result_symbol488 = &v205.F1
	*result_symbol488 = 4
	v206 = *lexer_addr
	mark_end489 = &v206.F3
	v207 = *mark_end489
	v208 = *lexer_addr
	v207(v208)
	v209 = *lookahead
	cmp490 = v209 == 61
	if cmp490 {
		goto if_then492
	} else {
		goto if_end493
	}

if_then492:
	*state_addr = 33
	goto next_state

if_end493:
	v210 = *lookahead
	cmp494 = v210 == 78
	if cmp494 {
		goto if_then496
	} else {
		goto if_end497
	}

if_then496:
	*state_addr = 21
	goto next_state

if_end497:
	v211 = *lookahead
	cmp498 = v211 == 43
	if cmp498 {
		goto if_then518
	} else {
		goto lor_lhs_false500
	}

lor_lhs_false500:
	v212 = *lookahead
	cmp501 = 47 <= v212
	if cmp501 {
		goto land_lhs_true503
	} else {
		goto lor_lhs_false506
	}

land_lhs_true503:
	v213 = *lookahead
	cmp504 = v213 <= 57
	if cmp504 {
		goto if_then518
	} else {
		goto lor_lhs_false506
	}

lor_lhs_false506:
	v214 = *lookahead
	cmp507 = 65 <= v214
	if cmp507 {
		goto land_lhs_true509
	} else {
		goto lor_lhs_false512
	}

land_lhs_true509:
	v215 = *lookahead
	cmp510 = v215 <= 90
	if cmp510 {
		goto if_then518
	} else {
		goto lor_lhs_false512
	}

lor_lhs_false512:
	v216 = *lookahead
	cmp513 = 97 <= v216
	if cmp513 {
		goto land_lhs_true515
	} else {
		goto if_end519
	}

land_lhs_true515:
	v217 = *lookahead
	cmp516 = v217 <= 122
	if cmp516 {
		goto if_then518
	} else {
		goto if_end519
	}

if_then518:
	*state_addr = 32
	goto next_state

if_end519:
	v218 = *result
	tobool520 = (v218 & 1) != 0
	*retval = tobool520
	goto _return

sw_bb521:
	*result = 1
	v219 = *lexer_addr
	result_symbol522 = &v219.F1
	*result_symbol522 = 4
	v220 = *lexer_addr
	mark_end523 = &v220.F3
	v221 = *mark_end523
	v222 = *lexer_addr
	v221(v222)
	v223 = *lookahead
	cmp524 = v223 == 61
	if cmp524 {
		goto if_then526
	} else {
		goto if_end527
	}

if_then526:
	*state_addr = 33
	goto next_state

if_end527:
	v224 = *lookahead
	cmp528 = v224 == 43
	if cmp528 {
		goto if_then548
	} else {
		goto lor_lhs_false530
	}

lor_lhs_false530:
	v225 = *lookahead
	cmp531 = 47 <= v225
	if cmp531 {
		goto land_lhs_true533
	} else {
		goto lor_lhs_false536
	}

land_lhs_true533:
	v226 = *lookahead
	cmp534 = v226 <= 57
	if cmp534 {
		goto if_then548
	} else {
		goto lor_lhs_false536
	}

lor_lhs_false536:
	v227 = *lookahead
	cmp537 = 65 <= v227
	if cmp537 {
		goto land_lhs_true539
	} else {
		goto lor_lhs_false542
	}

land_lhs_true539:
	v228 = *lookahead
	cmp540 = v228 <= 90
	if cmp540 {
		goto if_then548
	} else {
		goto lor_lhs_false542
	}

lor_lhs_false542:
	v229 = *lookahead
	cmp543 = 97 <= v229
	if cmp543 {
		goto land_lhs_true545
	} else {
		goto if_end549
	}

land_lhs_true545:
	v230 = *lookahead
	cmp546 = v230 <= 122
	if cmp546 {
		goto if_then548
	} else {
		goto if_end549
	}

if_then548:
	*state_addr = 32
	goto next_state

if_end549:
	v231 = *result
	tobool550 = (v231 & 1) != 0
	*retval = tobool550
	goto _return

sw_bb551:
	*result = 1
	v232 = *lexer_addr
	result_symbol552 = &v232.F1
	*result_symbol552 = 4
	v233 = *lexer_addr
	mark_end553 = &v233.F3
	v234 = *mark_end553
	v235 = *lexer_addr
	v234(v235)
	v236 = *lookahead
	cmp554 = v236 == 61
	if cmp554 {
		goto if_then556
	} else {
		goto if_end557
	}

if_then556:
	*state_addr = 25
	goto next_state

if_end557:
	v237 = *result
	tobool558 = (v237 & 1) != 0
	*retval = tobool558
	goto _return

sw_bb559:
	*result = 1
	v238 = *lexer_addr
	result_symbol560 = &v238.F1
	*result_symbol560 = 5
	v239 = *lexer_addr
	mark_end561 = &v239.F3
	v240 = *mark_end561
	v241 = *lexer_addr
	v240(v241)
	v242 = *lookahead
	cmp562 = v242 == 10
	if cmp562 {
		goto if_then564
	} else {
		goto if_end565
	}

if_then564:
	*state_addr = 34
	goto next_state

if_end565:
	v243 = *lookahead
	cmp566 = v243 == 13
	if cmp566 {
		goto if_then568
	} else {
		goto if_end569
	}

if_then568:
	*state_addr = 35
	goto next_state

if_end569:
	v244 = *lookahead
	cmp570 = v244 != 0
	if cmp570 {
		goto land_lhs_true572
	} else {
		goto if_end576
	}

land_lhs_true572:
	v245 = *lookahead
	cmp573 = v245 != 45
	if cmp573 {
		goto if_then575
	} else {
		goto if_end576
	}

if_then575:
	*state_addr = 36
	goto next_state

if_end576:
	v246 = *result
	tobool577 = (v246 & 1) != 0
	*retval = tobool577
	goto _return

sw_bb578:
	*result = 1
	v247 = *lexer_addr
	result_symbol579 = &v247.F1
	*result_symbol579 = 5
	v248 = *lexer_addr
	mark_end580 = &v248.F3
	v249 = *mark_end580
	v250 = *lexer_addr
	v249(v250)
	v251 = *lookahead
	cmp581 = v251 == 10
	if cmp581 {
		goto if_then583
	} else {
		goto if_end584
	}

if_then583:
	*state_addr = 34
	goto next_state

if_end584:
	v252 = *lookahead
	cmp585 = v252 != 0
	if cmp585 {
		goto land_lhs_true587
	} else {
		goto if_end591
	}

land_lhs_true587:
	v253 = *lookahead
	cmp588 = v253 != 45
	if cmp588 {
		goto if_then590
	} else {
		goto if_end591
	}

if_then590:
	*state_addr = 36
	goto next_state

if_end591:
	v254 = *result
	tobool592 = (v254 & 1) != 0
	*retval = tobool592
	goto _return

sw_bb593:
	*result = 1
	v255 = *lexer_addr
	result_symbol594 = &v255.F1
	*result_symbol594 = 5
	v256 = *lexer_addr
	mark_end595 = &v256.F3
	v257 = *mark_end595
	v258 = *lexer_addr
	v257(v258)
	v259 = *lookahead
	cmp596 = v259 != 0
	if cmp596 {
		goto land_lhs_true598
	} else {
		goto if_end602
	}

land_lhs_true598:
	v260 = *lookahead
	cmp599 = v260 != 45
	if cmp599 {
		goto if_then601
	} else {
		goto if_end602
	}

if_then601:
	*state_addr = 36
	goto next_state

if_end602:
	v261 = *result
	tobool603 = (v261 & 1) != 0
	*retval = tobool603
	goto _return

sw_bb604:
	*result = 1
	v262 = *lexer_addr
	result_symbol605 = &v262.F1
	*result_symbol605 = 6
	v263 = *lexer_addr
	mark_end606 = &v263.F3
	v264 = *mark_end606
	v265 = *lexer_addr
	v264(v265)
	v266 = *result
	tobool607 = (v266 & 1) != 0
	*retval = tobool607
	goto _return

sw_bb608:
	*result = 1
	v267 = *lexer_addr
	result_symbol609 = &v267.F1
	*result_symbol609 = 7
	v268 = *lexer_addr
	mark_end610 = &v268.F3
	v269 = *mark_end610
	v270 = *lexer_addr
	v269(v270)
	v271 = *lookahead
	cmp611 = v271 == 10
	if cmp611 {
		goto if_then613
	} else {
		goto if_end614
	}

if_then613:
	*skip = 1
	*state_addr = 1
	goto next_state

if_end614:
	v272 = *lookahead
	cmp615 = v272 != 0
	if cmp615 {
		goto if_then617
	} else {
		goto if_end618
	}

if_then617:
	*state_addr = 43
	goto next_state

if_end618:
	v273 = *result
	tobool619 = (v273 & 1) != 0
	*retval = tobool619
	goto _return

sw_bb620:
	*result = 1
	v274 = *lexer_addr
	result_symbol621 = &v274.F1
	*result_symbol621 = 7
	v275 = *lexer_addr
	mark_end622 = &v275.F3
	v276 = *mark_end622
	v277 = *lexer_addr
	v276(v277)
	v278 = *lookahead
	cmp623 = v278 == 45
	if cmp623 {
		goto if_then625
	} else {
		goto if_end626
	}

if_then625:
	*state_addr = 37
	goto next_state

if_end626:
	v279 = *lookahead
	cmp627 = v279 != 0
	if cmp627 {
		goto land_lhs_true629
	} else {
		goto if_end633
	}

land_lhs_true629:
	v280 = *lookahead
	cmp630 = v280 != 10
	if cmp630 {
		goto if_then632
	} else {
		goto if_end633
	}

if_then632:
	*state_addr = 43
	goto next_state

if_end633:
	v281 = *result
	tobool634 = (v281 & 1) != 0
	*retval = tobool634
	goto _return

sw_bb635:
	*result = 1
	v282 = *lexer_addr
	result_symbol636 = &v282.F1
	*result_symbol636 = 7
	v283 = *lexer_addr
	mark_end637 = &v283.F3
	v284 = *mark_end637
	v285 = *lexer_addr
	v284(v285)
	v286 = *lookahead
	cmp638 = v286 == 45
	if cmp638 {
		goto if_then640
	} else {
		goto if_end641
	}

if_then640:
	*state_addr = 39
	goto next_state

if_end641:
	v287 = *lookahead
	cmp642 = v287 != 0
	if cmp642 {
		goto land_lhs_true644
	} else {
		goto if_end648
	}

land_lhs_true644:
	v288 = *lookahead
	cmp645 = v288 != 10
	if cmp645 {
		goto if_then647
	} else {
		goto if_end648
	}

if_then647:
	*state_addr = 43
	goto next_state

if_end648:
	v289 = *result
	tobool649 = (v289 & 1) != 0
	*retval = tobool649
	goto _return

sw_bb650:
	*result = 1
	v290 = *lexer_addr
	result_symbol651 = &v290.F1
	*result_symbol651 = 7
	v291 = *lexer_addr
	mark_end652 = &v291.F3
	v292 = *mark_end652
	v293 = *lexer_addr
	v292(v293)
	v294 = *lookahead
	cmp653 = v294 == 45
	if cmp653 {
		goto if_then655
	} else {
		goto if_end656
	}

if_then655:
	*state_addr = 40
	goto next_state

if_end656:
	v295 = *lookahead
	cmp657 = v295 != 0
	if cmp657 {
		goto land_lhs_true659
	} else {
		goto if_end663
	}

land_lhs_true659:
	v296 = *lookahead
	cmp660 = v296 != 10
	if cmp660 {
		goto if_then662
	} else {
		goto if_end663
	}

if_then662:
	*state_addr = 43
	goto next_state

if_end663:
	v297 = *result
	tobool664 = (v297 & 1) != 0
	*retval = tobool664
	goto _return

sw_bb665:
	*result = 1
	v298 = *lexer_addr
	result_symbol666 = &v298.F1
	*result_symbol666 = 7
	v299 = *lexer_addr
	mark_end667 = &v299.F3
	v300 = *mark_end667
	v301 = *lexer_addr
	v300(v301)
	v302 = *lookahead
	cmp668 = v302 == 45
	if cmp668 {
		goto if_then670
	} else {
		goto if_end671
	}

if_then670:
	*state_addr = 41
	goto next_state

if_end671:
	v303 = *lookahead
	cmp672 = v303 != 0
	if cmp672 {
		goto land_lhs_true674
	} else {
		goto if_end678
	}

land_lhs_true674:
	v304 = *lookahead
	cmp675 = v304 != 10
	if cmp675 {
		goto if_then677
	} else {
		goto if_end678
	}

if_then677:
	*state_addr = 43
	goto next_state

if_end678:
	v305 = *result
	tobool679 = (v305 & 1) != 0
	*retval = tobool679
	goto _return

sw_bb680:
	*result = 1
	v306 = *lexer_addr
	result_symbol681 = &v306.F1
	*result_symbol681 = 7
	v307 = *lexer_addr
	mark_end682 = &v307.F3
	v308 = *mark_end682
	v309 = *lexer_addr
	v308(v309)
	v310 = *lookahead
	cmp683 = v310 != 0
	if cmp683 {
		goto land_lhs_true685
	} else {
		goto if_end689
	}

land_lhs_true685:
	v311 = *lookahead
	cmp686 = v311 != 10
	if cmp686 {
		goto if_then688
	} else {
		goto if_end689
	}

if_then688:
	*state_addr = 43
	goto next_state

if_end689:
	v312 = *result
	tobool690 = (v312 & 1) != 0
	*retval = tobool690
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v313 = *retval
	return v313
}

