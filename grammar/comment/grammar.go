package grammar_comment

import (
	"unsafe"
	"github.com/lewtec/leaven/libc"
)

type TSCharacterRange struct {
	F0 int32
	F1 int32
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

var tree_sitter_comment_language TSLanguage = TSLanguage{15, 35, 0, 27, 2, 16, 9, 1, 0, 3, &(*[9][35]int16)(unsafe.Pointer(&ts_parse_table))[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[71]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], &ts_lex_modes[0], ts_lex, nil, 0, anon_2{&ts_external_scanner_states[0][0], &ts_external_scanner_symbol_map[0], tree_sitter_comment_external_scanner_create, tree_sitter_comment_external_scanner_destroy, tree_sitter_comment_external_scanner_scan, tree_sitter_comment_external_scanner_serialize, tree_sitter_comment_external_scanner_deserialize}, &ts_primary_state_ids[0], &_str[0], nil, 0, 0, nil, nil, nil, TSLanguageMetadata{0, 2, 0}}

var ts_small_parse_table [71]int16 = [71]int16{
	3, 53, 1, 5, 7, 1, 33, 51, 14, 1, 4, 9, 10, 11, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 4, 3, 1, 5, 55, 1, 1, 57, 1,
	2, 13, 1, 29, 2, 3, 1, 5, 59, 1, 0, 2, 61, 1, 3, 63,
	1, 5, 2, 3, 1, 5, 65, 1, 1, 2, 3, 1, 5, 67, 1, 4,
	2, 3, 1, 5, 69, 1, 1,
}

var ts_small_parse_table_map [7]int32 = [7]int32{0, 23, 36, 43, 50, 57, 64}

var ts_symbol_names [35]*byte = [35]*byte{
	&_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0], &_str_8[0], &_str_9[0], &_str_10[0], &_str_11[0], &_str_12[0], &_str_13[0], &_str_14[0], &_str_15[0], &_str_16[0], &_str_17[0], &_str_18[0],
	&_str_19[0], &_str_20[0], &_str_21[0], &_str_22[0], &_str_23[0], &_str_24[0], &_str_25[0], &_str_26[0], &_str_27[0], &_str_28[0], &_str_29[0], &_str_30[0], &_str_31[0], &_str_32[0], &_str_33[0], &_str_34[0],
	&_str_35[0], &_str_34[0], &_str_36[0],
}

var ts_symbol_metadata [35]TSSymbolMetadata = [35]TSSymbolMetadata{
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0},
	TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0},
	TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 0, 0}, TSSymbolMetadata{},
}

var ts_symbol_map [35]int16 = [35]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 31, 34,
}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [1][3]int16 = [1][3]int16{}

var ts_lex_modes [16]TSLexerMode = [16]TSLexerMode{
	TSLexerMode{0, 1, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{0, 2, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{3, 0, 0}, TSLexerMode{}, TSLexerMode{}, TSLexerMode{},
}

var ts_external_scanner_states [3][2]byte = [3][2]byte{[2]byte{}, [2]byte{1, 1}, [2]byte{1, 0}}

var ts_external_scanner_symbol_map [2]int16 = [2]int16{25, 26}

var ts_primary_state_ids [16]int16 = [16]int16{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
}

var _str [8]byte = [8]byte{99, 111, 109, 109, 101, 110, 116, 0}

var __const_is_internal_char_valid_chars [2]int32 = [2]int32{45, 95}

var __const_is_space_space_chars [4]int32 = [4]int32{32, 12, 9, 11}

var __const_is_newline_newline_chars [3]int32 = [3]int32{0, 10, 13}

var ts_parse_table struct {
	F0 struct {
	F0 [27]int16
	F1 [8]int16
}
	F1 [35]int16
	F2 [35]int16
	F3 [35]int16
	F4 struct {
	F0 [26]int16
	F1 [9]int16
}
	F5 struct {
	F0 [26]int16
	F1 [9]int16
}
	F6 struct {
	F0 [26]int16
	F1 [9]int16
}
	F7 struct {
	F0 [26]int16
	F1 [9]int16
}
	F8 struct {
	F0 [26]int16
	F1 [9]int16
}
} = struct {
	F0 struct {
	F0 [27]int16
	F1 [8]int16
}
	F1 [35]int16
	F2 [35]int16
	F3 [35]int16
	F4 struct {
	F0 [26]int16
	F1 [9]int16
}
	F5 struct {
	F0 [26]int16
	F1 [9]int16
}
	F6 struct {
	F0 [26]int16
	F1 [9]int16
}
	F7 struct {
	F0 [26]int16
	F1 [9]int16
}
	F8 struct {
	F0 [26]int16
	F1 [9]int16
}
}{struct {
	F0 [27]int16
	F1 [8]int16
}{[27]int16{
	1, 1, 1, 0, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}, [8]int16{}}, [35]int16{
	5, 7, 7, 0, 7, 3, 9, 11, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 13, 0, 11, 2, 0, 2, 2,
	4, 0, 2,
}, [35]int16{
	15, 7, 7, 0, 7, 3, 9, 11, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 13, 0, 0, 3, 0, 3, 3,
	4, 0, 3,
}, [35]int16{
	17, 19, 19, 0, 19, 3, 22, 25, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 28, 0, 0, 3, 0, 3, 3,
	4, 0, 3,
}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{
	31, 31, 31, 0, 31, 3, 31, 33, 31, 31, 31, 31, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 31, 31, 31, 31, 31,
}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{
	35, 35, 35, 0, 35, 3, 35, 37, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{
	39, 39, 39, 0, 39, 3, 39, 41, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{
	43, 43, 43, 0, 43, 3, 43, 45, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
}, [9]int16{}}, struct {
	F0 [26]int16
	F1 [9]int16
}{[26]int16{
	47, 47, 47, 0, 47, 3, 47, 49, 47, 47, 47, 47, 47, 47, 47, 47,
	47, 47, 47, 47, 47, 47, 47, 47, 47, 47,
}, [9]int16{}}}

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
	F0 anon_1
	F1 [6]byte
}
	F20 TSParseActionEntry
	F21 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F22 struct {
	F0 anon_1
	F1 [6]byte
}
	F23 TSParseActionEntry
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
	F0 struct {
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
	F60 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
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
	F0 anon_1
	F1 [6]byte
}
	F20 TSParseActionEntry
	F21 struct {
	F0 struct {
	F0 struct {
	F0 byte
	F1 int16
	F2 byte
	F3 byte
}
	F1 [2]byte
}
}
	F22 struct {
	F0 anon_1
	F1 [6]byte
}
	F23 TSParseActionEntry
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
	F0 struct {
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
	F60 struct {
	F0 struct {
	F0 byte
	F1 [7]byte
}
}
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 0, 27, 0, 0}}}, struct {
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
}{0, 9, 0, 0}, [2]byte{}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 27, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 34, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 34, 0, 0}}}, struct {
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
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 34, 0, 0}}}, struct {
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
}{0, 9, 0, 1}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{2, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 34, 0, 0}}}, struct {
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
}{anon_1{2, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 34, 0, 0}}}, struct {
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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 31, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 31, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 28, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 28, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 33, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 33, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 30, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 2, 30, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 28, 0, 0}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 0}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 28, 0, 0}}}, struct {
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
}{0, 14, 0, 0}, [2]byte{}}}, struct {
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
}{0, 15, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 3, 29, 0, 0}}}}

var _str_3 [4]byte = [4]byte{101, 110, 100, 0}

var _str_4 [2]byte = [2]byte{58, 0}

var _str_5 [2]byte = [2]byte{40, 0}

var _str_6 [5]byte = [5]byte{117, 115, 101, 114, 0}

var _str_7 [2]byte = [2]byte{41, 0}

var _str_8 [17]byte = [17]byte{
	95, 102, 117, 108, 108, 95, 117, 114, 105, 95, 116, 111, 107, 101, 110, 49,
	0,
}

var _str_9 [4]byte = [4]byte{117, 114, 105, 0}

var _str_10 [13]byte = [13]byte{95, 116, 101, 120, 116, 95, 116, 111, 107, 101, 110, 49, 0}

var _str_11 [2]byte = [2]byte{47, 0}

var _str_12 [2]byte = [2]byte{39, 0}

var _str_13 [2]byte = [2]byte{34, 0}

var _str_14 [2]byte = [2]byte{96, 0}

var _str_15 [2]byte = [2]byte{60, 0}

var _str_16 [2]byte = [2]byte{91, 0}

var _str_17 [2]byte = [2]byte{123, 0}

var _str_18 [2]byte = [2]byte{46, 0}

var _str_19 [2]byte = [2]byte{44, 0}

var _str_20 [2]byte = [2]byte{59, 0}

var _str_21 [2]byte = [2]byte{33, 0}

var _str_22 [2]byte = [2]byte{63, 0}

var _str_23 [2]byte = [2]byte{92, 0}

var _str_24 [2]byte = [2]byte{125, 0}

var _str_25 [2]byte = [2]byte{93, 0}

var _str_26 [2]byte = [2]byte{62, 0}

var _str_27 [2]byte = [2]byte{45, 0}

var _str_28 [5]byte = [5]byte{110, 97, 109, 101, 0}

var _str_29 [14]byte = [14]byte{105, 110, 118, 97, 108, 105, 100, 95, 116, 111, 107, 101, 110, 0}

var _str_30 [7]byte = [7]byte{115, 111, 117, 114, 99, 101, 0}

var _str_31 [4]byte = [4]byte{116, 97, 103, 0}

var _str_32 [6]byte = [6]byte{95, 117, 115, 101, 114, 0}

var _str_33 [10]byte = [10]byte{95, 102, 117, 108, 108, 95, 117, 114, 105, 0}

var _str_34 [5]byte = [5]byte{116, 101, 120, 116, 0}

var _str_35 [11]byte = [11]byte{95, 115, 116, 111, 112, 95, 99, 104, 97, 114, 0}

var _str_36 [15]byte = [15]byte{115, 111, 117, 114, 99, 101, 95, 114, 101, 112, 101, 97, 116, 49, 0}

var ts_lex_map [42]int16 = [42]int16{
	33, 29, 34, 21, 39, 20, 40, 8, 41, 10, 44, 27, 45, 35, 46, 26,
	47, 19, 58, 7, 59, 28, 60, 23, 62, 34, 63, 30, 91, 24, 92, 31,
	93, 33, 96, 22, 104, 17, 123, 25, 125, 32,
}

var ts_lex_map_37 [28]int16 = [28]int16{
	33, 5, 34, 5, 39, 5, 41, 5, 44, 5, 46, 5, 58, 5, 59, 5,
	62, 5, 63, 5, 92, 5, 93, 5, 96, 5, 125, 5,
}

var sym_uri_character_set_1 [12]TSCharacterRange = [12]TSCharacterRange{TSCharacterRange{0, 8}, TSCharacterRange{14, 31}, TSCharacterRange{35, 38}, TSCharacterRange{40, 40}, TSCharacterRange{42, 43}, TSCharacterRange{45, 45}, TSCharacterRange{47, 57}, TSCharacterRange{60, 61}, TSCharacterRange{64, 91}, TSCharacterRange{94, 95}, TSCharacterRange{97, 124}, TSCharacterRange{126, 1114111}}

var ts_lex_map_38 [28]int16 = [28]int16{
	33, 5, 34, 5, 39, 5, 41, 5, 44, 5, 46, 5, 58, 5, 59, 5,
	62, 5, 63, 5, 92, 5, 93, 5, 96, 5, 125, 5,
}

var aux_sym__text_token1_character_set_1 [11]TSCharacterRange = [11]TSCharacterRange{TSCharacterRange{0, 8}, TSCharacterRange{14, 31}, TSCharacterRange{35, 38}, TSCharacterRange{42, 43}, TSCharacterRange{48, 57}, TSCharacterRange{61, 61}, TSCharacterRange{64, 90}, TSCharacterRange{94, 95}, TSCharacterRange{97, 122}, TSCharacterRange{124, 124}, TSCharacterRange{126, 1114111}}

func tree_sitter_comment_external_scanner_create() *byte {
	return nil
}

func tree_sitter_comment_external_scanner_destroy(payload *byte) {
	var payload_addr **byte

	_ = payload_addr

	payload_addr = new(*byte)
	*payload_addr = payload
}

func tree_sitter_comment_external_scanner_serialize(payload *byte, buffer *byte) int32 {
	var payload_addr, buffer_addr **byte

	_, _ = payload_addr, buffer_addr

	payload_addr = new(*byte)
	buffer_addr = new(*byte)
	*payload_addr = payload
	*buffer_addr = buffer
	return 0
}

func tree_sitter_comment_external_scanner_deserialize(payload *byte, buffer *byte, length int32) {
	var payload_addr, buffer_addr **byte
	var length_addr *int32

	_, _, _ = payload_addr, buffer_addr, length_addr

	payload_addr = new(*byte)
	buffer_addr = new(*byte)
	length_addr = new(int32)
	*payload_addr = payload
	*buffer_addr = buffer
	*length_addr = length
}

func tree_sitter_comment_external_scanner_scan(payload *byte, lexer *TSLexer, valid_symbols *byte) bool {
	var lexer_addr **TSLexer
	var payload_addr, valid_symbols_addr **byte
	var v0 *TSLexer
	var v1 *byte
	var call bool

	_, _, _, _, _, _ = payload_addr, lexer_addr, valid_symbols_addr, v0, v1, call

	payload_addr = new(*byte)
	lexer_addr = new(*TSLexer)
	valid_symbols_addr = new(*byte)
	*payload_addr = payload
	*lexer_addr = lexer
	*valid_symbols_addr = valid_symbols
	v0 = *lexer_addr
	v1 = *valid_symbols_addr
	call = parse(v0, v1)
	return call
}

func parse(lexer *TSLexer, valid_symbols *byte) bool {
	var lexer_addr **TSLexer
	var valid_symbols_addr **byte
	var v2, v6 *TSLexer
	var retval *bool
	var v0, arrayidx, v4, arrayidx1, v7 *byte
	var lookahead *int32
	var tobool, call, tobool2, call4, v8 bool
	var v1, v5 byte
	var v3 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, valid_symbols_addr, v0, arrayidx, v1, tobool, v2, lookahead, v3, call, v4, arrayidx1, v5, tobool2, v6, v7, call4, v8

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	valid_symbols_addr = new(*byte)
	*lexer_addr = lexer
	*valid_symbols_addr = valid_symbols
	v0 = *valid_symbols_addr
	arrayidx = libc.AddPointer(v0, int(int64(1)))
	v1 = *arrayidx
	tobool = (v1 & 1) != 0
	if tobool {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*retval = false
	goto _return

if_end:
	v2 = *lexer_addr
	lookahead = &v2.F0
	v3 = *lookahead
	call = is_upper(v3)
	if call {
		goto land_lhs_true
	} else {
		goto if_end5
	}

land_lhs_true:
	v4 = *valid_symbols_addr
	arrayidx1 = v4
	v5 = *arrayidx1
	tobool2 = (v5 & 1) != 0
	if tobool2 {
		goto if_then3
	} else {
		goto if_end5
	}

if_then3:
	v6 = *lexer_addr
	v7 = *valid_symbols_addr
	call4 = parse_tagname(v6, v7)
	*retval = call4
	goto _return

if_end5:
	*retval = false
	goto _return

_return:
	v8 = *retval
	return v8
}

func tree_sitter_comment() *TSLanguage {
	return &tree_sitter_comment_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v50, v51, v53, v55, v56, v58, v60, v61, v63, v65, v66, v68, v73, v74, v76, v78, v79, v81, v83, v84, v86, v100, v101, v103, v109, v110, v112, v117, v118, v120, v125, v126, v128, v133, v134, v136, v141, v142, v144, v148, v149, v151, v153, v154, v156, v158, v159, v161, v163, v164, v166, v168, v169, v171, v173, v174, v176, v178, v179, v181, v183, v184, v186, v188, v189, v191, v193, v194, v196, v198, v199, v201, v203, v204, v206, v208, v209, v211, v213, v214, v216, v218, v219, v221, v223, v224, v226, v228, v229, v231 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end110, mark_end114, mark_end118, mark_end132, mark_end136, mark_end140, mark_end179, mark_end197, mark_end211, mark_end225, mark_end239, mark_end253, mark_end263, mark_end267, mark_end271, mark_end275, mark_end279, mark_end283, mark_end287, mark_end291, mark_end295, mark_end299, mark_end303, mark_end307, mark_end311, mark_end315, mark_end319, mark_end323, mark_end327 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, arrayidx, arrayidx11, arrayidx69, arrayidx76, result_symbol, result_symbol109, result_symbol113, result_symbol117, result_symbol131, result_symbol135, result_symbol139, arrayidx148, arrayidx155, result_symbol178, result_symbol196, result_symbol210, result_symbol224, result_symbol238, result_symbol252, result_symbol262, result_symbol266, result_symbol270, result_symbol274, result_symbol278, result_symbol282, result_symbol286, result_symbol290, result_symbol294, result_symbol298, result_symbol302, result_symbol306, result_symbol310, result_symbol314, result_symbol318, result_symbol322, result_symbol326 *int16
	var lookahead, i, i62, i141, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, cmp18, cmp22, tobool26, cmp28, tobool32, cmp34, tobool38, cmp40, cmp43, cmp46, cmp50, cmp53, cmp56, tobool60, cmp65, cmp71, cmp81, cmp84, cmp87, cmp90, cmp93, tobool97, tobool99, call101, tobool105, tobool107, tobool111, tobool115, cmp119, cmp122, cmp125, tobool129, tobool133, tobool137, cmp144, cmp150, cmp160, cmp163, cmp166, cmp169, cmp172, tobool176, cmp180, cmp184, tobool188, call190, tobool194, cmp198, tobool202, call204, tobool208, cmp212, tobool216, call218, tobool222, cmp226, tobool230, call232, tobool236, cmp240, tobool244, call246, tobool250, tobool254, call256, tobool260, tobool264, tobool268, tobool272, tobool276, tobool280, tobool284, tobool288, tobool292, tobool296, tobool300, tobool304, tobool308, tobool312, tobool316, tobool320, tobool324, tobool328, v233 bool
	var v3, frombool, v10, v22, v24, v26, v33, v46, v47, v49, v54, v59, v64, v72, v77, v82, v99, v106, v108, v114, v116, v122, v124, v130, v132, v138, v140, v145, v147, v152, v157, v162, v167, v172, v177, v182, v187, v192, v197, v202, v207, v212, v217, v222, v227, v232 byte
	var v52, v57, v62, v67, v75, v80, v85, v102, v111, v119, v127, v135, v143, v150, v155, v160, v165, v170, v175, v180, v185, v190, v195, v200, v205, v210, v215, v220, v225, v230 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9, v13, v16, v36, v39, v89, v92 int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v20, v21, v23, v25, v27, v28, v29, v30, v31, v32, v34, v35, conv70, v37, v38, add74, v40, add79, v41, v42, v43, v44, v45, v48, v69, v70, v71, v87, v88, conv149, v90, v91, add153, v93, add158, v94, v95, v96, v97, v98, v104, v105, v107, v113, v115, v121, v123, v129, v131, v137, v139, v146 int32
	var conv4, idxprom, idxprom10, conv64, idxprom68, idxprom75, conv143, idxprom147, idxprom154 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, i62, i141, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, cmp18, v21, cmp22, v22, tobool26, v23, cmp28, v24, tobool32, v25, cmp34, v26, tobool38, v27, cmp40, v28, cmp43, v29, cmp46, v30, cmp50, v31, cmp53, v32, cmp56, v33, tobool60, v34, conv64, cmp65, v35, idxprom68, arrayidx69, v36, conv70, v37, cmp71, v38, add74, idxprom75, arrayidx76, v39, v40, add79, v41, cmp81, v42, cmp84, v43, cmp87, v44, cmp90, v45, cmp93, v46, tobool97, v47, tobool99, v48, call101, v49, tobool105, v50, result_symbol, v51, mark_end, v52, v53, v54, tobool107, v55, result_symbol109, v56, mark_end110, v57, v58, v59, tobool111, v60, result_symbol113, v61, mark_end114, v62, v63, v64, tobool115, v65, result_symbol117, v66, mark_end118, v67, v68, v69, cmp119, v70, cmp122, v71, cmp125, v72, tobool129, v73, result_symbol131, v74, mark_end132, v75, v76, v77, tobool133, v78, result_symbol135, v79, mark_end136, v80, v81, v82, tobool137, v83, result_symbol139, v84, mark_end140, v85, v86, v87, conv143, cmp144, v88, idxprom147, arrayidx148, v89, conv149, v90, cmp150, v91, add153, idxprom154, arrayidx155, v92, v93, add158, v94, cmp160, v95, cmp163, v96, cmp166, v97, cmp169, v98, cmp172, v99, tobool176, v100, result_symbol178, v101, mark_end179, v102, v103, v104, cmp180, v105, cmp184, v106, tobool188, v107, call190, v108, tobool194, v109, result_symbol196, v110, mark_end197, v111, v112, v113, cmp198, v114, tobool202, v115, call204, v116, tobool208, v117, result_symbol210, v118, mark_end211, v119, v120, v121, cmp212, v122, tobool216, v123, call218, v124, tobool222, v125, result_symbol224, v126, mark_end225, v127, v128, v129, cmp226, v130, tobool230, v131, call232, v132, tobool236, v133, result_symbol238, v134, mark_end239, v135, v136, v137, cmp240, v138, tobool244, v139, call246, v140, tobool250, v141, result_symbol252, v142, mark_end253, v143, v144, v145, tobool254, v146, call256, v147, tobool260, v148, result_symbol262, v149, mark_end263, v150, v151, v152, tobool264, v153, result_symbol266, v154, mark_end267, v155, v156, v157, tobool268, v158, result_symbol270, v159, mark_end271, v160, v161, v162, tobool272, v163, result_symbol274, v164, mark_end275, v165, v166, v167, tobool276, v168, result_symbol278, v169, mark_end279, v170, v171, v172, tobool280, v173, result_symbol282, v174, mark_end283, v175, v176, v177, tobool284, v178, result_symbol286, v179, mark_end287, v180, v181, v182, tobool288, v183, result_symbol290, v184, mark_end291, v185, v186, v187, tobool292, v188, result_symbol294, v189, mark_end295, v190, v191, v192, tobool296, v193, result_symbol298, v194, mark_end299, v195, v196, v197, tobool300, v198, result_symbol302, v199, mark_end303, v200, v201, v202, tobool304, v203, result_symbol306, v204, mark_end307, v205, v206, v207, tobool308, v208, result_symbol310, v209, mark_end311, v210, v211, v212, tobool312, v213, result_symbol314, v214, mark_end315, v215, v216, v217, tobool316, v218, result_symbol318, v219, mark_end319, v220, v221, v222, tobool320, v223, result_symbol322, v224, mark_end323, v225, v226, v227, tobool324, v228, result_symbol326, v229, mark_end327, v230, v231, v232, tobool328, v233

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	state_addr = new(int16)
	result = new(byte)
	skip = new(byte)
	eof = new(byte)
	lookahead = new(int32)
	i = new(int32)
	i62 = new(int32)
	i141 = new(int32)
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
		goto sw_bb33
	case 3:
		goto sw_bb39
	case 4:
		goto sw_bb61
	case 5:
		goto sw_bb98
	case 6:
		goto sw_bb106
	case 7:
		goto sw_bb108
	case 8:
		goto sw_bb112
	case 9:
		goto sw_bb116
	case 10:
		goto sw_bb130
	case 11:
		goto sw_bb134
	case 12:
		goto sw_bb138
	case 13:
		goto sw_bb177
	case 14:
		goto sw_bb195
	case 15:
		goto sw_bb209
	case 16:
		goto sw_bb223
	case 17:
		goto sw_bb237
	case 18:
		goto sw_bb251
	case 19:
		goto sw_bb261
	case 20:
		goto sw_bb265
	case 21:
		goto sw_bb269
	case 22:
		goto sw_bb273
	case 23:
		goto sw_bb277
	case 24:
		goto sw_bb281
	case 25:
		goto sw_bb285
	case 26:
		goto sw_bb289
	case 27:
		goto sw_bb293
	case 28:
		goto sw_bb297
	case 29:
		goto sw_bb301
	case 30:
		goto sw_bb305
	case 31:
		goto sw_bb309
	case 32:
		goto sw_bb313
	case 33:
		goto sw_bb317
	case 34:
		goto sw_bb321
	case 35:
		goto sw_bb325
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
	*i = 0
	goto for_cond

for_cond:
	v11 = *i
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(42)
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
	*state_addr = 11
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
	*state_addr = 18
	goto next_state

if_end25:
	v22 = *result
	tobool26 = (v22 & 1) != 0
	*retval = tobool26
	goto _return

sw_bb27:
	v23 = *lookahead
	cmp28 = v23 == 47
	if cmp28 {
		goto if_then30
	} else {
		goto if_end31
	}

if_then30:
	*state_addr = 4
	goto next_state

if_end31:
	v24 = *result
	tobool32 = (v24 & 1) != 0
	*retval = tobool32
	goto _return

sw_bb33:
	v25 = *lookahead
	cmp34 = v25 == 47
	if cmp34 {
		goto if_then36
	} else {
		goto if_end37
	}

if_then36:
	*state_addr = 1
	goto next_state

if_end37:
	v26 = *result
	tobool38 = (v26 & 1) != 0
	*retval = tobool38
	goto _return

sw_bb39:
	v27 = *lookahead
	cmp40 = 9 <= v27
	if cmp40 {
		goto land_lhs_true42
	} else {
		goto lor_lhs_false45
	}

land_lhs_true42:
	v28 = *lookahead
	cmp43 = v28 <= 13
	if cmp43 {
		goto if_then48
	} else {
		goto lor_lhs_false45
	}

lor_lhs_false45:
	v29 = *lookahead
	cmp46 = v29 == 32
	if cmp46 {
		goto if_then48
	} else {
		goto if_end49
	}

if_then48:
	*state_addr = 9
	goto next_state

if_end49:
	v30 = *lookahead
	cmp50 = v30 != 0
	if cmp50 {
		goto land_lhs_true52
	} else {
		goto if_end59
	}

land_lhs_true52:
	v31 = *lookahead
	cmp53 = v31 != 40
	if cmp53 {
		goto land_lhs_true55
	} else {
		goto if_end59
	}

land_lhs_true55:
	v32 = *lookahead
	cmp56 = v32 != 41
	if cmp56 {
		goto if_then58
	} else {
		goto if_end59
	}

if_then58:
	*state_addr = 9
	goto next_state

if_end59:
	v33 = *result
	tobool60 = (v33 & 1) != 0
	*retval = tobool60
	goto _return

sw_bb61:
	*i62 = 0
	goto for_cond63

for_cond63:
	v34 = *i62
	conv64 = int64(uint64(uint32(v34)))
	cmp65 = uint64(conv64) < uint64(28)
	if cmp65 {
		goto for_body67
	} else {
		goto for_end80
	}

for_body67:
	v35 = *i62
	idxprom68 = int64(uint64(uint32(v35)))
	arrayidx69 = &ts_lex_map_37[idxprom68]
	v36 = *arrayidx69
	conv70 = int32(uint32(uint16(v36)))
	v37 = *lookahead
	cmp71 = conv70 == v37
	if cmp71 {
		goto if_then73
	} else {
		goto if_end77
	}

if_then73:
	v38 = *i62
	add74 = v38 + 1
	idxprom75 = int64(uint64(uint32(add74)))
	arrayidx76 = &ts_lex_map_37[idxprom75]
	v39 = *arrayidx76
	*state_addr = v39
	goto next_state

if_end77:
	goto for_inc78

for_inc78:
	v40 = *i62
	add79 = v40 + 2
	*i62 = add79
	goto for_cond63

for_end80:
	v41 = *lookahead
	cmp81 = v41 != 0
	if cmp81 {
		goto land_lhs_true83
	} else {
		goto if_end96
	}

land_lhs_true83:
	v42 = *lookahead
	cmp84 = v42 < 9
	if cmp84 {
		goto land_lhs_true89
	} else {
		goto lor_lhs_false86
	}

lor_lhs_false86:
	v43 = *lookahead
	cmp87 = 13 < v43
	if cmp87 {
		goto land_lhs_true89
	} else {
		goto if_end96
	}

land_lhs_true89:
	v44 = *lookahead
	cmp90 = v44 < 32
	if cmp90 {
		goto if_then95
	} else {
		goto lor_lhs_false92
	}

lor_lhs_false92:
	v45 = *lookahead
	cmp93 = 34 < v45
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*state_addr = 12
	goto next_state

if_end96:
	v46 = *result
	tobool97 = (v46 & 1) != 0
	*retval = tobool97
	goto _return

sw_bb98:
	v47 = *eof
	tobool99 = (v47 & 1) != 0
	if tobool99 {
		goto if_end104
	} else {
		goto land_lhs_true100
	}

land_lhs_true100:
	v48 = *lookahead
	call101 = set_contains(&sym_uri_character_set_1[int64(0)], 12, v48)
	if call101 {
		goto if_then103
	} else {
		goto if_end104
	}

if_then103:
	*state_addr = 12
	goto next_state

if_end104:
	v49 = *result
	tobool105 = (v49 & 1) != 0
	*retval = tobool105
	goto _return

sw_bb106:
	*result = 1
	v50 = *lexer_addr
	result_symbol = &v50.F1
	*result_symbol = 0
	v51 = *lexer_addr
	mark_end = &v51.F3
	v52 = *mark_end
	v53 = *lexer_addr
	v52(v53)
	v54 = *result
	tobool107 = (v54 & 1) != 0
	*retval = tobool107
	goto _return

sw_bb108:
	*result = 1
	v55 = *lexer_addr
	result_symbol109 = &v55.F1
	*result_symbol109 = 1
	v56 = *lexer_addr
	mark_end110 = &v56.F3
	v57 = *mark_end110
	v58 = *lexer_addr
	v57(v58)
	v59 = *result
	tobool111 = (v59 & 1) != 0
	*retval = tobool111
	goto _return

sw_bb112:
	*result = 1
	v60 = *lexer_addr
	result_symbol113 = &v60.F1
	*result_symbol113 = 2
	v61 = *lexer_addr
	mark_end114 = &v61.F3
	v62 = *mark_end114
	v63 = *lexer_addr
	v62(v63)
	v64 = *result
	tobool115 = (v64 & 1) != 0
	*retval = tobool115
	goto _return

sw_bb116:
	*result = 1
	v65 = *lexer_addr
	result_symbol117 = &v65.F1
	*result_symbol117 = 3
	v66 = *lexer_addr
	mark_end118 = &v66.F3
	v67 = *mark_end118
	v68 = *lexer_addr
	v67(v68)
	v69 = *lookahead
	cmp119 = v69 != 0
	if cmp119 {
		goto land_lhs_true121
	} else {
		goto if_end128
	}

land_lhs_true121:
	v70 = *lookahead
	cmp122 = v70 != 40
	if cmp122 {
		goto land_lhs_true124
	} else {
		goto if_end128
	}

land_lhs_true124:
	v71 = *lookahead
	cmp125 = v71 != 41
	if cmp125 {
		goto if_then127
	} else {
		goto if_end128
	}

if_then127:
	*state_addr = 9
	goto next_state

if_end128:
	v72 = *result
	tobool129 = (v72 & 1) != 0
	*retval = tobool129
	goto _return

sw_bb130:
	*result = 1
	v73 = *lexer_addr
	result_symbol131 = &v73.F1
	*result_symbol131 = 4
	v74 = *lexer_addr
	mark_end132 = &v74.F3
	v75 = *mark_end132
	v76 = *lexer_addr
	v75(v76)
	v77 = *result
	tobool133 = (v77 & 1) != 0
	*retval = tobool133
	goto _return

sw_bb134:
	*result = 1
	v78 = *lexer_addr
	result_symbol135 = &v78.F1
	*result_symbol135 = 5
	v79 = *lexer_addr
	mark_end136 = &v79.F3
	v80 = *mark_end136
	v81 = *lexer_addr
	v80(v81)
	v82 = *result
	tobool137 = (v82 & 1) != 0
	*retval = tobool137
	goto _return

sw_bb138:
	*result = 1
	v83 = *lexer_addr
	result_symbol139 = &v83.F1
	*result_symbol139 = 6
	v84 = *lexer_addr
	mark_end140 = &v84.F3
	v85 = *mark_end140
	v86 = *lexer_addr
	v85(v86)
	*i141 = 0
	goto for_cond142

for_cond142:
	v87 = *i141
	conv143 = int64(uint64(uint32(v87)))
	cmp144 = uint64(conv143) < uint64(28)
	if cmp144 {
		goto for_body146
	} else {
		goto for_end159
	}

for_body146:
	v88 = *i141
	idxprom147 = int64(uint64(uint32(v88)))
	arrayidx148 = &ts_lex_map_38[idxprom147]
	v89 = *arrayidx148
	conv149 = int32(uint32(uint16(v89)))
	v90 = *lookahead
	cmp150 = conv149 == v90
	if cmp150 {
		goto if_then152
	} else {
		goto if_end156
	}

if_then152:
	v91 = *i141
	add153 = v91 + 1
	idxprom154 = int64(uint64(uint32(add153)))
	arrayidx155 = &ts_lex_map_38[idxprom154]
	v92 = *arrayidx155
	*state_addr = v92
	goto next_state

if_end156:
	goto for_inc157

for_inc157:
	v93 = *i141
	add158 = v93 + 2
	*i141 = add158
	goto for_cond142

for_end159:
	v94 = *lookahead
	cmp160 = v94 != 0
	if cmp160 {
		goto land_lhs_true162
	} else {
		goto if_end175
	}

land_lhs_true162:
	v95 = *lookahead
	cmp163 = v95 < 9
	if cmp163 {
		goto land_lhs_true168
	} else {
		goto lor_lhs_false165
	}

lor_lhs_false165:
	v96 = *lookahead
	cmp166 = 13 < v96
	if cmp166 {
		goto land_lhs_true168
	} else {
		goto if_end175
	}

land_lhs_true168:
	v97 = *lookahead
	cmp169 = v97 < 32
	if cmp169 {
		goto if_then174
	} else {
		goto lor_lhs_false171
	}

lor_lhs_false171:
	v98 = *lookahead
	cmp172 = 34 < v98
	if cmp172 {
		goto if_then174
	} else {
		goto if_end175
	}

if_then174:
	*state_addr = 12
	goto next_state

if_end175:
	v99 = *result
	tobool176 = (v99 & 1) != 0
	*retval = tobool176
	goto _return

sw_bb177:
	*result = 1
	v100 = *lexer_addr
	result_symbol178 = &v100.F1
	*result_symbol178 = 7
	v101 = *lexer_addr
	mark_end179 = &v101.F3
	v102 = *mark_end179
	v103 = *lexer_addr
	v102(v103)
	v104 = *lookahead
	cmp180 = v104 == 58
	if cmp180 {
		goto if_then182
	} else {
		goto if_end183
	}

if_then182:
	*state_addr = 2
	goto next_state

if_end183:
	v105 = *lookahead
	cmp184 = v105 == 115
	if cmp184 {
		goto if_then186
	} else {
		goto if_end187
	}

if_then186:
	*state_addr = 14
	goto next_state

if_end187:
	v106 = *eof
	tobool188 = (v106 & 1) != 0
	if tobool188 {
		goto if_end193
	} else {
		goto land_lhs_true189
	}

land_lhs_true189:
	v107 = *lookahead
	call190 = set_contains(&aux_sym__text_token1_character_set_1[int64(0)], 11, v107)
	if call190 {
		goto if_then192
	} else {
		goto if_end193
	}

if_then192:
	*state_addr = 18
	goto next_state

if_end193:
	v108 = *result
	tobool194 = (v108 & 1) != 0
	*retval = tobool194
	goto _return

sw_bb195:
	*result = 1
	v109 = *lexer_addr
	result_symbol196 = &v109.F1
	*result_symbol196 = 7
	v110 = *lexer_addr
	mark_end197 = &v110.F3
	v111 = *mark_end197
	v112 = *lexer_addr
	v111(v112)
	v113 = *lookahead
	cmp198 = v113 == 58
	if cmp198 {
		goto if_then200
	} else {
		goto if_end201
	}

if_then200:
	*state_addr = 2
	goto next_state

if_end201:
	v114 = *eof
	tobool202 = (v114 & 1) != 0
	if tobool202 {
		goto if_end207
	} else {
		goto land_lhs_true203
	}

land_lhs_true203:
	v115 = *lookahead
	call204 = set_contains(&aux_sym__text_token1_character_set_1[int64(0)], 11, v115)
	if call204 {
		goto if_then206
	} else {
		goto if_end207
	}

if_then206:
	*state_addr = 18
	goto next_state

if_end207:
	v116 = *result
	tobool208 = (v116 & 1) != 0
	*retval = tobool208
	goto _return

sw_bb209:
	*result = 1
	v117 = *lexer_addr
	result_symbol210 = &v117.F1
	*result_symbol210 = 7
	v118 = *lexer_addr
	mark_end211 = &v118.F3
	v119 = *mark_end211
	v120 = *lexer_addr
	v119(v120)
	v121 = *lookahead
	cmp212 = v121 == 112
	if cmp212 {
		goto if_then214
	} else {
		goto if_end215
	}

if_then214:
	*state_addr = 13
	goto next_state

if_end215:
	v122 = *eof
	tobool216 = (v122 & 1) != 0
	if tobool216 {
		goto if_end221
	} else {
		goto land_lhs_true217
	}

land_lhs_true217:
	v123 = *lookahead
	call218 = set_contains(&aux_sym__text_token1_character_set_1[int64(0)], 11, v123)
	if call218 {
		goto if_then220
	} else {
		goto if_end221
	}

if_then220:
	*state_addr = 18
	goto next_state

if_end221:
	v124 = *result
	tobool222 = (v124 & 1) != 0
	*retval = tobool222
	goto _return

sw_bb223:
	*result = 1
	v125 = *lexer_addr
	result_symbol224 = &v125.F1
	*result_symbol224 = 7
	v126 = *lexer_addr
	mark_end225 = &v126.F3
	v127 = *mark_end225
	v128 = *lexer_addr
	v127(v128)
	v129 = *lookahead
	cmp226 = v129 == 116
	if cmp226 {
		goto if_then228
	} else {
		goto if_end229
	}

if_then228:
	*state_addr = 15
	goto next_state

if_end229:
	v130 = *eof
	tobool230 = (v130 & 1) != 0
	if tobool230 {
		goto if_end235
	} else {
		goto land_lhs_true231
	}

land_lhs_true231:
	v131 = *lookahead
	call232 = set_contains(&aux_sym__text_token1_character_set_1[int64(0)], 11, v131)
	if call232 {
		goto if_then234
	} else {
		goto if_end235
	}

if_then234:
	*state_addr = 18
	goto next_state

if_end235:
	v132 = *result
	tobool236 = (v132 & 1) != 0
	*retval = tobool236
	goto _return

sw_bb237:
	*result = 1
	v133 = *lexer_addr
	result_symbol238 = &v133.F1
	*result_symbol238 = 7
	v134 = *lexer_addr
	mark_end239 = &v134.F3
	v135 = *mark_end239
	v136 = *lexer_addr
	v135(v136)
	v137 = *lookahead
	cmp240 = v137 == 116
	if cmp240 {
		goto if_then242
	} else {
		goto if_end243
	}

if_then242:
	*state_addr = 16
	goto next_state

if_end243:
	v138 = *eof
	tobool244 = (v138 & 1) != 0
	if tobool244 {
		goto if_end249
	} else {
		goto land_lhs_true245
	}

land_lhs_true245:
	v139 = *lookahead
	call246 = set_contains(&aux_sym__text_token1_character_set_1[int64(0)], 11, v139)
	if call246 {
		goto if_then248
	} else {
		goto if_end249
	}

if_then248:
	*state_addr = 18
	goto next_state

if_end249:
	v140 = *result
	tobool250 = (v140 & 1) != 0
	*retval = tobool250
	goto _return

sw_bb251:
	*result = 1
	v141 = *lexer_addr
	result_symbol252 = &v141.F1
	*result_symbol252 = 7
	v142 = *lexer_addr
	mark_end253 = &v142.F3
	v143 = *mark_end253
	v144 = *lexer_addr
	v143(v144)
	v145 = *eof
	tobool254 = (v145 & 1) != 0
	if tobool254 {
		goto if_end259
	} else {
		goto land_lhs_true255
	}

land_lhs_true255:
	v146 = *lookahead
	call256 = set_contains(&aux_sym__text_token1_character_set_1[int64(0)], 11, v146)
	if call256 {
		goto if_then258
	} else {
		goto if_end259
	}

if_then258:
	*state_addr = 18
	goto next_state

if_end259:
	v147 = *result
	tobool260 = (v147 & 1) != 0
	*retval = tobool260
	goto _return

sw_bb261:
	*result = 1
	v148 = *lexer_addr
	result_symbol262 = &v148.F1
	*result_symbol262 = 8
	v149 = *lexer_addr
	mark_end263 = &v149.F3
	v150 = *mark_end263
	v151 = *lexer_addr
	v150(v151)
	v152 = *result
	tobool264 = (v152 & 1) != 0
	*retval = tobool264
	goto _return

sw_bb265:
	*result = 1
	v153 = *lexer_addr
	result_symbol266 = &v153.F1
	*result_symbol266 = 9
	v154 = *lexer_addr
	mark_end267 = &v154.F3
	v155 = *mark_end267
	v156 = *lexer_addr
	v155(v156)
	v157 = *result
	tobool268 = (v157 & 1) != 0
	*retval = tobool268
	goto _return

sw_bb269:
	*result = 1
	v158 = *lexer_addr
	result_symbol270 = &v158.F1
	*result_symbol270 = 10
	v159 = *lexer_addr
	mark_end271 = &v159.F3
	v160 = *mark_end271
	v161 = *lexer_addr
	v160(v161)
	v162 = *result
	tobool272 = (v162 & 1) != 0
	*retval = tobool272
	goto _return

sw_bb273:
	*result = 1
	v163 = *lexer_addr
	result_symbol274 = &v163.F1
	*result_symbol274 = 11
	v164 = *lexer_addr
	mark_end275 = &v164.F3
	v165 = *mark_end275
	v166 = *lexer_addr
	v165(v166)
	v167 = *result
	tobool276 = (v167 & 1) != 0
	*retval = tobool276
	goto _return

sw_bb277:
	*result = 1
	v168 = *lexer_addr
	result_symbol278 = &v168.F1
	*result_symbol278 = 12
	v169 = *lexer_addr
	mark_end279 = &v169.F3
	v170 = *mark_end279
	v171 = *lexer_addr
	v170(v171)
	v172 = *result
	tobool280 = (v172 & 1) != 0
	*retval = tobool280
	goto _return

sw_bb281:
	*result = 1
	v173 = *lexer_addr
	result_symbol282 = &v173.F1
	*result_symbol282 = 13
	v174 = *lexer_addr
	mark_end283 = &v174.F3
	v175 = *mark_end283
	v176 = *lexer_addr
	v175(v176)
	v177 = *result
	tobool284 = (v177 & 1) != 0
	*retval = tobool284
	goto _return

sw_bb285:
	*result = 1
	v178 = *lexer_addr
	result_symbol286 = &v178.F1
	*result_symbol286 = 14
	v179 = *lexer_addr
	mark_end287 = &v179.F3
	v180 = *mark_end287
	v181 = *lexer_addr
	v180(v181)
	v182 = *result
	tobool288 = (v182 & 1) != 0
	*retval = tobool288
	goto _return

sw_bb289:
	*result = 1
	v183 = *lexer_addr
	result_symbol290 = &v183.F1
	*result_symbol290 = 15
	v184 = *lexer_addr
	mark_end291 = &v184.F3
	v185 = *mark_end291
	v186 = *lexer_addr
	v185(v186)
	v187 = *result
	tobool292 = (v187 & 1) != 0
	*retval = tobool292
	goto _return

sw_bb293:
	*result = 1
	v188 = *lexer_addr
	result_symbol294 = &v188.F1
	*result_symbol294 = 16
	v189 = *lexer_addr
	mark_end295 = &v189.F3
	v190 = *mark_end295
	v191 = *lexer_addr
	v190(v191)
	v192 = *result
	tobool296 = (v192 & 1) != 0
	*retval = tobool296
	goto _return

sw_bb297:
	*result = 1
	v193 = *lexer_addr
	result_symbol298 = &v193.F1
	*result_symbol298 = 17
	v194 = *lexer_addr
	mark_end299 = &v194.F3
	v195 = *mark_end299
	v196 = *lexer_addr
	v195(v196)
	v197 = *result
	tobool300 = (v197 & 1) != 0
	*retval = tobool300
	goto _return

sw_bb301:
	*result = 1
	v198 = *lexer_addr
	result_symbol302 = &v198.F1
	*result_symbol302 = 18
	v199 = *lexer_addr
	mark_end303 = &v199.F3
	v200 = *mark_end303
	v201 = *lexer_addr
	v200(v201)
	v202 = *result
	tobool304 = (v202 & 1) != 0
	*retval = tobool304
	goto _return

sw_bb305:
	*result = 1
	v203 = *lexer_addr
	result_symbol306 = &v203.F1
	*result_symbol306 = 19
	v204 = *lexer_addr
	mark_end307 = &v204.F3
	v205 = *mark_end307
	v206 = *lexer_addr
	v205(v206)
	v207 = *result
	tobool308 = (v207 & 1) != 0
	*retval = tobool308
	goto _return

sw_bb309:
	*result = 1
	v208 = *lexer_addr
	result_symbol310 = &v208.F1
	*result_symbol310 = 20
	v209 = *lexer_addr
	mark_end311 = &v209.F3
	v210 = *mark_end311
	v211 = *lexer_addr
	v210(v211)
	v212 = *result
	tobool312 = (v212 & 1) != 0
	*retval = tobool312
	goto _return

sw_bb313:
	*result = 1
	v213 = *lexer_addr
	result_symbol314 = &v213.F1
	*result_symbol314 = 21
	v214 = *lexer_addr
	mark_end315 = &v214.F3
	v215 = *mark_end315
	v216 = *lexer_addr
	v215(v216)
	v217 = *result
	tobool316 = (v217 & 1) != 0
	*retval = tobool316
	goto _return

sw_bb317:
	*result = 1
	v218 = *lexer_addr
	result_symbol318 = &v218.F1
	*result_symbol318 = 22
	v219 = *lexer_addr
	mark_end319 = &v219.F3
	v220 = *mark_end319
	v221 = *lexer_addr
	v220(v221)
	v222 = *result
	tobool320 = (v222 & 1) != 0
	*retval = tobool320
	goto _return

sw_bb321:
	*result = 1
	v223 = *lexer_addr
	result_symbol322 = &v223.F1
	*result_symbol322 = 23
	v224 = *lexer_addr
	mark_end323 = &v224.F3
	v225 = *mark_end323
	v226 = *lexer_addr
	v225(v226)
	v227 = *result
	tobool324 = (v227 & 1) != 0
	*retval = tobool324
	goto _return

sw_bb325:
	*result = 1
	v228 = *lexer_addr
	result_symbol326 = &v228.F1
	*result_symbol326 = 24
	v229 = *lexer_addr
	mark_end327 = &v229.F3
	v230 = *mark_end327
	v231 = *lexer_addr
	v230(v231)
	v232 = *result
	tobool328 = (v232 & 1) != 0
	*retval = tobool328
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v233 = *retval
	return v233
}

func is_upper(c int32) bool {
	var c_addr, upper, lower *int32
	var cmp, cmp1, v2 bool
	var v0, v1 int32

	_, _, _, _, _, _, _, _ = c_addr, upper, lower, v0, cmp, v1, cmp1, v2

	c_addr = new(int32)
	upper = new(int32)
	lower = new(int32)
	*c_addr = c
	*upper = 65
	*lower = 90
	v0 = *c_addr
	cmp = v0 >= 65
	if cmp {
		goto land_rhs
	} else {
		v2 = false
		goto land_end
	}

land_rhs:
	v1 = *c_addr
	cmp1 = v1 <= 90
	v2 = cmp1
	goto land_end

land_end:
	return v2
}

func parse_tagname(lexer *TSLexer, valid_symbols *byte) bool {
	var lexer_addr **TSLexer
	var valid_symbols_addr **byte
	var v0, v4, v6, v8, v9, v11, v13, v16, v18, v20, v21, v23, v25, v27, v29, v31, v33, v36, v38, v39, v41, v43, v44, v46, v48, v50, v53, v55, v56, v58, v60, v61, v63 *TSLexer
	var retval *bool
	var v2, arrayidx *byte
	var mark_end *func(*TSLexer)
	var advance, advance10, advance27, advance33, advance42, advance47, advance53 *func(*TSLexer, bool)
	var result_symbol *int16
	var previous, user_length, lookahead, lookahead1, lookahead2, lookahead5, lookahead7, lookahead9, lookahead14, lookahead16, lookahead19, lookahead22, lookahead24, lookahead29, lookahead35, lookahead38, lookahead49, lookahead54 *int32
	var call, tobool, call3, call6, call8, v15, call11, call15, call17, cmp, call23, call25, lnot, v35, cmp30, cmp36, call39, cmp44, cmp50, call55, v64 bool
	var v3 byte
	var v22 func(*TSLexer)
	var v7, v19, v37, v42, v49, v54, v59 func(*TSLexer, bool)
	var v1, v5, v10, v12, v14, v17, v24, v26, v28, v30, v32, v34, v40, v45, v47, v51, inc, v52, v57, v62 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, valid_symbols_addr, previous, user_length, v0, lookahead, v1, call, v2, arrayidx, v3, tobool, v4, lookahead1, v5, v6, advance, v7, v8, v9, lookahead2, v10, call3, v11, lookahead5, v12, call6, v13, lookahead7, v14, call8, v15, v16, lookahead9, v17, v18, advance10, v19, v20, v21, mark_end, v22, v23, v24, call11, v25, lookahead14, v26, call15, v27, lookahead16, v28, call17, v29, lookahead19, v30, cmp, v31, lookahead22, v32, call23, v33, lookahead24, v34, call25, lnot, v35, v36, advance27, v37, v38, v39, lookahead29, v40, cmp30, v41, advance33, v42, v43, v44, lookahead35, v45, cmp36, v46, lookahead38, v47, call39, v48, advance42, v49, v50, v51, inc, v52, cmp44, v53, advance47, v54, v55, v56, lookahead49, v57, cmp50, v58, advance53, v59, v60, v61, lookahead54, v62, call55, v63, result_symbol, v64

	retval = new(bool)
	lexer_addr = new(*TSLexer)
	valid_symbols_addr = new(*byte)
	previous = new(int32)
	user_length = new(int32)
	*lexer_addr = lexer
	*valid_symbols_addr = valid_symbols
	v0 = *lexer_addr
	lookahead = &v0.F0
	v1 = *lookahead
	call = is_upper(v1)
	if call {
		goto lor_lhs_false
	} else {
		goto if_then
	}

lor_lhs_false:
	v2 = *valid_symbols_addr
	arrayidx = v2
	v3 = *arrayidx
	tobool = (v3 & 1) != 0
	if tobool {
		goto if_end
	} else {
		goto if_then
	}

if_then:
	*retval = false
	goto _return

if_end:
	v4 = *lexer_addr
	lookahead1 = &v4.F0
	v5 = *lookahead1
	*previous = v5
	v6 = *lexer_addr
	advance = &v6.F2
	v7 = *advance
	v8 = *lexer_addr
	v7(v8, false)
	goto while_cond

while_cond:
	v9 = *lexer_addr
	lookahead2 = &v9.F0
	v10 = *lookahead2
	call3 = is_upper(v10)
	if call3 {
		v15 = true
		goto lor_end
	} else {
		goto lor_lhs_false4
	}

lor_lhs_false4:
	v11 = *lexer_addr
	lookahead5 = &v11.F0
	v12 = *lookahead5
	call6 = is_digit(v12)
	if call6 {
		v15 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v13 = *lexer_addr
	lookahead7 = &v13.F0
	v14 = *lookahead7
	call8 = is_internal_char(v14)
	v15 = call8
	goto lor_end

lor_end:
	if v15 {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v16 = *lexer_addr
	lookahead9 = &v16.F0
	v17 = *lookahead9
	*previous = v17
	v18 = *lexer_addr
	advance10 = &v18.F2
	v19 = *advance10
	v20 = *lexer_addr
	v19(v20, false)
	goto while_cond

while_end:
	v21 = *lexer_addr
	mark_end = &v21.F3
	v22 = *mark_end
	v23 = *lexer_addr
	v22(v23)
	v24 = *previous
	call11 = is_internal_char(v24)
	if call11 {
		goto if_then12
	} else {
		goto if_end13
	}

if_then12:
	*retval = false
	goto _return

if_end13:
	v25 = *lexer_addr
	lookahead14 = &v25.F0
	v26 = *lookahead14
	call15 = is_space(v26)
	if call15 {
		goto land_lhs_true
	} else {
		goto lor_lhs_false18
	}

land_lhs_true:
	v27 = *lexer_addr
	lookahead16 = &v27.F0
	v28 = *lookahead16
	call17 = is_newline(v28)
	if call17 {
		goto lor_lhs_false18
	} else {
		goto if_then20
	}

lor_lhs_false18:
	v29 = *lexer_addr
	lookahead19 = &v29.F0
	v30 = *lookahead19
	cmp = v30 == 40
	if cmp {
		goto if_then20
	} else {
		goto if_end48
	}

if_then20:
	goto while_cond21

while_cond21:
	v31 = *lexer_addr
	lookahead22 = &v31.F0
	v32 = *lookahead22
	call23 = is_space(v32)
	if call23 {
		goto land_rhs
	} else {
		v35 = false
		goto land_end
	}

land_rhs:
	v33 = *lexer_addr
	lookahead24 = &v33.F0
	v34 = *lookahead24
	call25 = is_newline(v34)
	lnot = call25 != true
	v35 = lnot
	goto land_end

land_end:
	if v35 {
		goto while_body26
	} else {
		goto while_end28
	}

while_body26:
	v36 = *lexer_addr
	advance27 = &v36.F2
	v37 = *advance27
	v38 = *lexer_addr
	v37(v38, false)
	goto while_cond21

while_end28:
	v39 = *lexer_addr
	lookahead29 = &v39.F0
	v40 = *lookahead29
	cmp30 = v40 != 40
	if cmp30 {
		goto if_then31
	} else {
		goto if_end32
	}

if_then31:
	*retval = false
	goto _return

if_end32:
	v41 = *lexer_addr
	advance33 = &v41.F2
	v42 = *advance33
	v43 = *lexer_addr
	v42(v43, false)
	*user_length = 0
	goto while_cond34

while_cond34:
	v44 = *lexer_addr
	lookahead35 = &v44.F0
	v45 = *lookahead35
	cmp36 = v45 != 41
	if cmp36 {
		goto while_body37
	} else {
		goto while_end43
	}

while_body37:
	v46 = *lexer_addr
	lookahead38 = &v46.F0
	v47 = *lookahead38
	call39 = is_newline(v47)
	if call39 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*retval = false
	goto _return

if_end41:
	v48 = *lexer_addr
	advance42 = &v48.F2
	v49 = *advance42
	v50 = *lexer_addr
	v49(v50, false)
	v51 = *user_length
	inc = v51 + 1
	*user_length = inc
	goto while_cond34

while_end43:
	v52 = *user_length
	cmp44 = v52 <= 0
	if cmp44 {
		goto if_then45
	} else {
		goto if_end46
	}

if_then45:
	*retval = false
	goto _return

if_end46:
	v53 = *lexer_addr
	advance47 = &v53.F2
	v54 = *advance47
	v55 = *lexer_addr
	v54(v55, false)
	goto if_end48

if_end48:
	v56 = *lexer_addr
	lookahead49 = &v56.F0
	v57 = *lookahead49
	cmp50 = v57 != 58
	if cmp50 {
		goto if_then51
	} else {
		goto if_end52
	}

if_then51:
	*retval = false
	goto _return

if_end52:
	v58 = *lexer_addr
	advance53 = &v58.F2
	v59 = *advance53
	v60 = *lexer_addr
	v59(v60, false)
	v61 = *lexer_addr
	lookahead54 = &v61.F0
	v62 = *lookahead54
	call55 = is_space(v62)
	if call55 {
		goto if_end57
	} else {
		goto if_then56
	}

if_then56:
	*retval = false
	goto _return

if_end57:
	v63 = *lexer_addr
	result_symbol = &v63.F1
	*result_symbol = 0
	*retval = true
	goto _return

_return:
	v64 = *retval
	return v64
}

func is_digit(c int32) bool {
	var c_addr, upper, lower *int32
	var cmp, cmp1, v2 bool
	var v0, v1 int32

	_, _, _, _, _, _, _, _ = c_addr, upper, lower, v0, cmp, v1, cmp1, v2

	c_addr = new(int32)
	upper = new(int32)
	lower = new(int32)
	*c_addr = c
	*upper = 48
	*lower = 57
	v0 = *c_addr
	cmp = v0 >= 48
	if cmp {
		goto land_rhs
	} else {
		v2 = false
		goto land_end
	}

land_rhs:
	v1 = *c_addr
	cmp1 = v1 <= 57
	v2 = cmp1
	goto land_end

land_end:
	return v2
}

func is_internal_char(c int32) bool {
	var valid_chars *[2]int32
	var retval *bool
	var v0 *byte
	var c_addr, length, i, arrayidx *int32
	var cmp, cmp1, v6 bool
	var v1, v2, v3, v4, v5, inc int32
	var idxprom int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, c_addr, valid_chars, length, i, v0, v1, cmp, v2, v3, idxprom, arrayidx, v4, cmp1, v5, inc, v6

	retval = new(bool)
	c_addr = new(int32)
	valid_chars = &new(struct{v [2]int32; b byte}).v
	length = new(int32)
	i = new(int32)
	*c_addr = c
	v0 = (*byte)(unsafe.Pointer(valid_chars))
	libc.Memmove(v0, (*byte)(unsafe.Pointer(&__const_is_internal_char_valid_chars)), int64(8))
	*length = 2
	*i = 0
	goto for_cond

for_cond:
	v1 = *i
	cmp = v1 < 2
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v2 = *c_addr
	v3 = *i
	idxprom = int64(v3)
	arrayidx = &valid_chars[idxprom]
	v4 = *arrayidx
	cmp1 = v2 == v4
	if cmp1 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*retval = true
	goto _return

if_end:
	goto for_inc

for_inc:
	v5 = *i
	inc = v5 + 1
	*i = inc
	goto for_cond

for_end:
	*retval = false
	goto _return

_return:
	v6 = *retval
	return v6
}

func is_space(c int32) bool {
	var space_chars *[4]int32
	var is_space_char, v0 *byte
	var c_addr, length, i, arrayidx *int32
	var cmp, cmp1, tobool, call, v8 bool
	var v6 byte
	var v1, v2, v3, v4, v5, inc, v7 int32
	var idxprom int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c_addr, space_chars, length, is_space_char, i, v0, v1, cmp, v2, v3, idxprom, arrayidx, v4, cmp1, v5, inc, v6, tobool, v7, call, v8

	c_addr = new(int32)
	space_chars = &new(struct{v [4]int32; b byte}).v
	length = new(int32)
	is_space_char = new(byte)
	i = new(int32)
	*c_addr = c
	v0 = (*byte)(unsafe.Pointer(space_chars))
	libc.Memmove(v0, (*byte)(unsafe.Pointer(&__const_is_space_space_chars)), int64(16))
	*length = 4
	*is_space_char = 0
	*i = 0
	goto for_cond

for_cond:
	v1 = *i
	cmp = v1 < 4
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v2 = *c_addr
	v3 = *i
	idxprom = int64(v3)
	arrayidx = &space_chars[idxprom]
	v4 = *arrayidx
	cmp1 = v2 == v4
	if cmp1 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*is_space_char = 1
	goto for_end

if_end:
	goto for_inc

for_inc:
	v5 = *i
	inc = v5 + 1
	*i = inc
	goto for_cond

for_end:
	v6 = *is_space_char
	tobool = (v6 & 1) != 0
	if tobool {
		v8 = true
		goto lor_end
	} else {
		goto lor_rhs
	}

lor_rhs:
	v7 = *c_addr
	call = is_newline(v7)
	v8 = call
	goto lor_end

lor_end:
	return v8
}

func is_newline(c int32) bool {
	var newline_chars *[3]int32
	var retval *bool
	var v0 *byte
	var c_addr, length, i, arrayidx *int32
	var cmp, cmp1, v6 bool
	var v1, v2, v3, v4, v5, inc int32
	var idxprom int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, c_addr, newline_chars, length, i, v0, v1, cmp, v2, v3, idxprom, arrayidx, v4, cmp1, v5, inc, v6

	retval = new(bool)
	c_addr = new(int32)
	newline_chars = &new(struct{v [3]int32; b byte}).v
	length = new(int32)
	i = new(int32)
	*c_addr = c
	v0 = (*byte)(unsafe.Pointer(newline_chars))
	libc.Memmove(v0, (*byte)(unsafe.Pointer(&__const_is_newline_newline_chars)), int64(12))
	*length = 3
	*i = 0
	goto for_cond

for_cond:
	v1 = *i
	cmp = v1 < 3
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v2 = *c_addr
	v3 = *i
	idxprom = int64(v3)
	arrayidx = &newline_chars[idxprom]
	v4 = *arrayidx
	cmp1 = v2 == v4
	if cmp1 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*retval = true
	goto _return

if_end:
	goto for_inc

for_inc:
	v5 = *i
	inc = v5 + 1
	*i = inc
	goto for_cond

for_end:
	*retval = false
	goto _return

_return:
	v6 = *retval
	return v6
}

func set_contains(ranges *TSCharacterRange, len int32, lookahead int32) bool {
	var ranges_addr, _range, range8 **TSCharacterRange
	var v6, arrayidx, v9, v12, v15, v20, arrayidx10, v23, v26 *TSCharacterRange
	var retval *bool
	var len_addr, lookahead_addr, index, size, half_size, mid_index, start, end, end3, start11, end13 *int32
	var cmp, cmp1, cmp2, cmp4, cmp12, cmp14, v28, v29 bool
	var v0, v1, sub, v2, v3, div, v4, v5, add, v7, v8, v10, v11, v13, v14, v16, v17, v18, v19, sub7, v21, v22, v24, v25, v27 int32
	var idxprom, idxprom9 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, ranges_addr, len_addr, lookahead_addr, index, size, half_size, mid_index, _range, range8, v0, v1, sub, v2, cmp, v3, div, v4, v5, add, v6, v7, idxprom, arrayidx, v8, v9, start, v10, cmp1, v11, v12, end, v13, cmp2, v14, v15, end3, v16, cmp4, v17, v18, v19, sub7, v20, v21, idxprom9, arrayidx10, v22, v23, start11, v24, cmp12, v25, v26, end13, v27, cmp14, v28, v29

	retval = new(bool)
	ranges_addr = new(*TSCharacterRange)
	len_addr = new(int32)
	lookahead_addr = new(int32)
	index = new(int32)
	size = new(int32)
	half_size = new(int32)
	mid_index = new(int32)
	_range = new(*TSCharacterRange)
	range8 = new(*TSCharacterRange)
	*ranges_addr = ranges
	*len_addr = len
	*lookahead_addr = lookahead
	*index = 0
	v0 = *len_addr
	v1 = *index
	sub = v0 - v1
	*size = sub
	goto while_cond

while_cond:
	v2 = *size
	cmp = uint32(v2) > 1
	if cmp {
		goto while_body
	} else {
		goto while_end
	}

while_body:
	v3 = *size
	div = int32(uint32(v3) / 2)
	*half_size = div
	v4 = *index
	v5 = *half_size
	add = v4 + v5
	*mid_index = add
	v6 = *ranges_addr
	v7 = *mid_index
	idxprom = int64(uint64(uint32(v7)))
	arrayidx = libc.AddPointer(v6, int(idxprom))
	*_range = arrayidx
	v8 = *lookahead_addr
	v9 = *_range
	start = &v9.F0
	v10 = *start
	cmp1 = v8 >= v10
	if cmp1 {
		goto land_lhs_true
	} else {
		goto if_else
	}

land_lhs_true:
	v11 = *lookahead_addr
	v12 = *_range
	end = &v12.F1
	v13 = *end
	cmp2 = v11 <= v13
	if cmp2 {
		goto if_then
	} else {
		goto if_else
	}

if_then:
	*retval = true
	goto _return

if_else:
	v14 = *lookahead_addr
	v15 = *_range
	end3 = &v15.F1
	v16 = *end3
	cmp4 = v14 > v16
	if cmp4 {
		goto if_then5
	} else {
		goto if_end
	}

if_then5:
	v17 = *mid_index
	*index = v17
	goto if_end

if_end:
	goto if_end6

if_end6:
	v18 = *half_size
	v19 = *size
	sub7 = v19 - v18
	*size = sub7
	goto while_cond

while_end:
	v20 = *ranges_addr
	v21 = *index
	idxprom9 = int64(uint64(uint32(v21)))
	arrayidx10 = libc.AddPointer(v20, int(idxprom9))
	*range8 = arrayidx10
	v22 = *lookahead_addr
	v23 = *range8
	start11 = &v23.F0
	v24 = *start11
	cmp12 = v22 >= v24
	if cmp12 {
		goto land_rhs
	} else {
		v28 = false
		goto land_end
	}

land_rhs:
	v25 = *lookahead_addr
	v26 = *range8
	end13 = &v26.F1
	v27 = *end13
	cmp14 = v25 <= v27
	v28 = cmp14
	goto land_end

land_end:
	*retval = v28
	goto _return

_return:
	v29 = *retval
	return v29
}

