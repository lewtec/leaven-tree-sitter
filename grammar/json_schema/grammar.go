package grammar_json_schema

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

var tree_sitter_json_schema_language TSLanguage = TSLanguage{14, 6, 0, 5, 0, 4, 2, 1, 0, 1, &ts_parse_table[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[11]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], (*TSLexerMode)(unsafe.Pointer(&ts_lex_modes)), ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], nil, nil, 0, 0, nil, nil, nil, TSLanguageMetadata{}}

var ts_parse_table [2][6]int16 = [2][6]int16{[6]int16{1, 1, 1, 1, 1, 0}, [6]int16{0, 3, 3, 5, 5, 3}}

var ts_small_parse_table [8]int16 = [8]int16{1, 7, 1, 0, 1, 9, 1, 0}

var ts_small_parse_table_map [2]int32 = [2]int32{0, 4}

var ts_symbol_names [6]*byte = [6]*byte{&_str[0], &_str_2[0], &_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0]}

var ts_symbol_metadata [6]TSSymbolMetadata = [6]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}}

var ts_symbol_map [6]int16 = [6]int16{0, 1, 2, 3, 4, 5}

var ts_non_terminal_alias_map [1]int16 = [1]int16{}

var ts_alias_sequences [1][1]int16 = [1][1]int16{}

var ts_lex_modes [4]TSLexMode = [4]TSLexMode{}

var ts_primary_state_ids [4]int16 = [4]int16{0, 1, 2, 3}

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
}{0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_1
	F1 [6]byte
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 5, 0, 0}}}, struct {
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

var _str_2 [5]byte = [5]byte{110, 117, 108, 108, 0}

var _str_3 [5]byte = [5]byte{98, 111, 111, 108, 0}

var _str_4 [4]byte = [4]byte{105, 110, 116, 0}

var _str_5 [6]byte = [6]byte{102, 108, 111, 97, 116, 0}

var _str_6 [7]byte = [7]byte{115, 99, 97, 108, 97, 114, 0}

func tree_sitter_json_schema() *TSLanguage {
	return &tree_sitter_json_schema_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v49, v50, v52, v54, v55, v57, v59, v60, v62, v64, v65, v67, v72, v73, v75, v82, v83, v85, v91, v92, v94 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end125, mark_end129, mark_end133, mark_end148, mark_end170, mark_end188 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, result_symbol, result_symbol124, result_symbol128, result_symbol132, result_symbol147, result_symbol169, result_symbol187 *int16
	var lookahead, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp11, cmp15, cmp19, cmp23, cmp25, tobool29, cmp31, cmp35, cmp38, tobool42, cmp44, tobool48, cmp50, tobool54, cmp56, tobool60, cmp62, tobool66, cmp68, tobool72, cmp74, tobool78, cmp80, tobool84, cmp86, tobool90, cmp92, tobool96, cmp98, cmp100, cmp104, cmp107, tobool111, cmp113, cmp116, tobool120, tobool122, tobool126, tobool130, cmp134, cmp138, cmp141, tobool145, cmp149, cmp153, cmp156, cmp160, cmp163, tobool167, cmp171, cmp174, cmp178, cmp181, tobool185, cmp189, cmp192, tobool196, v98 bool
	var v3, frombool, v10, v18, v22, v24, v26, v28, v30, v32, v34, v36, v38, v40, v45, v48, v53, v58, v63, v71, v81, v90, v97 byte
	var v51, v56, v61, v66, v74, v84, v93 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9 int16
	var v5, conv, v11, v12, v13, v14, v15, v16, v17, v19, v20, v21, v23, v25, v27, v29, v31, v33, v35, v37, v39, v41, v42, v43, v44, v46, v47, v68, v69, v70, v76, v77, v78, v79, v80, v86, v87, v88, v89, v95, v96 int32

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, cmp, v12, cmp7, v13, cmp11, v14, cmp15, v15, cmp19, v16, cmp23, v17, cmp25, v18, tobool29, v19, cmp31, v20, cmp35, v21, cmp38, v22, tobool42, v23, cmp44, v24, tobool48, v25, cmp50, v26, tobool54, v27, cmp56, v28, tobool60, v29, cmp62, v30, tobool66, v31, cmp68, v32, tobool72, v33, cmp74, v34, tobool78, v35, cmp80, v36, tobool84, v37, cmp86, v38, tobool90, v39, cmp92, v40, tobool96, v41, cmp98, v42, cmp100, v43, cmp104, v44, cmp107, v45, tobool111, v46, cmp113, v47, cmp116, v48, tobool120, v49, result_symbol, v50, mark_end, v51, v52, v53, tobool122, v54, result_symbol124, v55, mark_end125, v56, v57, v58, tobool126, v59, result_symbol128, v60, mark_end129, v61, v62, v63, tobool130, v64, result_symbol132, v65, mark_end133, v66, v67, v68, cmp134, v69, cmp138, v70, cmp141, v71, tobool145, v72, result_symbol147, v73, mark_end148, v74, v75, v76, cmp149, v77, cmp153, v78, cmp156, v79, cmp160, v80, cmp163, v81, tobool167, v82, result_symbol169, v83, mark_end170, v84, v85, v86, cmp171, v87, cmp174, v88, cmp178, v89, cmp181, v90, tobool185, v91, result_symbol187, v92, mark_end188, v93, v94, v95, cmp189, v96, cmp192, v97, tobool196, v98

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
		goto sw_bb30
	case 2:
		goto sw_bb43
	case 3:
		goto sw_bb49
	case 4:
		goto sw_bb55
	case 5:
		goto sw_bb61
	case 6:
		goto sw_bb67
	case 7:
		goto sw_bb73
	case 8:
		goto sw_bb79
	case 9:
		goto sw_bb85
	case 10:
		goto sw_bb91
	case 11:
		goto sw_bb97
	case 12:
		goto sw_bb112
	case 13:
		goto sw_bb121
	case 14:
		goto sw_bb123
	case 15:
		goto sw_bb127
	case 16:
		goto sw_bb131
	case 17:
		goto sw_bb146
	case 18:
		goto sw_bb168
	case 19:
		goto sw_bb186
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
	*state_addr = 13
	goto next_state

if_end:
	v11 = *lookahead
	cmp = v11 == 45
	if cmp {
		goto if_then5
	} else {
		goto if_end6
	}

if_then5:
	*state_addr = 1
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
	*state_addr = 16
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
	cmp15 = v14 == 110
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
	cmp19 = v15 == 116
	if cmp19 {
		goto if_then21
	} else {
		goto if_end22
	}

if_then21:
	*state_addr = 7
	goto next_state

if_end22:
	v16 = *lookahead
	cmp23 = 49 <= v16
	if cmp23 {
		goto land_lhs_true
	} else {
		goto if_end28
	}

land_lhs_true:
	v17 = *lookahead
	cmp25 = v17 <= 57
	if cmp25 {
		goto if_then27
	} else {
		goto if_end28
	}

if_then27:
	*state_addr = 17
	goto next_state

if_end28:
	v18 = *result
	tobool29 = (v18 & 1) != 0
	*retval = tobool29
	goto _return

sw_bb30:
	v19 = *lookahead
	cmp31 = v19 == 48
	if cmp31 {
		goto if_then33
	} else {
		goto if_end34
	}

if_then33:
	*state_addr = 16
	goto next_state

if_end34:
	v20 = *lookahead
	cmp35 = 49 <= v20
	if cmp35 {
		goto land_lhs_true37
	} else {
		goto if_end41
	}

land_lhs_true37:
	v21 = *lookahead
	cmp38 = v21 <= 57
	if cmp38 {
		goto if_then40
	} else {
		goto if_end41
	}

if_then40:
	*state_addr = 17
	goto next_state

if_end41:
	v22 = *result
	tobool42 = (v22 & 1) != 0
	*retval = tobool42
	goto _return

sw_bb43:
	v23 = *lookahead
	cmp44 = v23 == 97
	if cmp44 {
		goto if_then46
	} else {
		goto if_end47
	}

if_then46:
	*state_addr = 4
	goto next_state

if_end47:
	v24 = *result
	tobool48 = (v24 & 1) != 0
	*retval = tobool48
	goto _return

sw_bb49:
	v25 = *lookahead
	cmp50 = v25 == 101
	if cmp50 {
		goto if_then52
	} else {
		goto if_end53
	}

if_then52:
	*state_addr = 15
	goto next_state

if_end53:
	v26 = *result
	tobool54 = (v26 & 1) != 0
	*retval = tobool54
	goto _return

sw_bb55:
	v27 = *lookahead
	cmp56 = v27 == 108
	if cmp56 {
		goto if_then58
	} else {
		goto if_end59
	}

if_then58:
	*state_addr = 8
	goto next_state

if_end59:
	v28 = *result
	tobool60 = (v28 & 1) != 0
	*retval = tobool60
	goto _return

sw_bb61:
	v29 = *lookahead
	cmp62 = v29 == 108
	if cmp62 {
		goto if_then64
	} else {
		goto if_end65
	}

if_then64:
	*state_addr = 14
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
	*state_addr = 5
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
	*state_addr = 9
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
	*state_addr = 3
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
	cmp92 = v39 == 117
	if cmp92 {
		goto if_then94
	} else {
		goto if_end95
	}

if_then94:
	*state_addr = 6
	goto next_state

if_end95:
	v40 = *result
	tobool96 = (v40 & 1) != 0
	*retval = tobool96
	goto _return

sw_bb97:
	v41 = *lookahead
	cmp98 = v41 == 43
	if cmp98 {
		goto if_then102
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v42 = *lookahead
	cmp100 = v42 == 45
	if cmp100 {
		goto if_then102
	} else {
		goto if_end103
	}

if_then102:
	*state_addr = 12
	goto next_state

if_end103:
	v43 = *lookahead
	cmp104 = 48 <= v43
	if cmp104 {
		goto land_lhs_true106
	} else {
		goto if_end110
	}

land_lhs_true106:
	v44 = *lookahead
	cmp107 = v44 <= 57
	if cmp107 {
		goto if_then109
	} else {
		goto if_end110
	}

if_then109:
	*state_addr = 19
	goto next_state

if_end110:
	v45 = *result
	tobool111 = (v45 & 1) != 0
	*retval = tobool111
	goto _return

sw_bb112:
	v46 = *lookahead
	cmp113 = 48 <= v46
	if cmp113 {
		goto land_lhs_true115
	} else {
		goto if_end119
	}

land_lhs_true115:
	v47 = *lookahead
	cmp116 = v47 <= 57
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
	*result = 1
	v49 = *lexer_addr
	result_symbol = &v49.F1
	*result_symbol = 0
	v50 = *lexer_addr
	mark_end = &v50.F3
	v51 = *mark_end
	v52 = *lexer_addr
	v51(v52)
	v53 = *result
	tobool122 = (v53 & 1) != 0
	*retval = tobool122
	goto _return

sw_bb123:
	*result = 1
	v54 = *lexer_addr
	result_symbol124 = &v54.F1
	*result_symbol124 = 1
	v55 = *lexer_addr
	mark_end125 = &v55.F3
	v56 = *mark_end125
	v57 = *lexer_addr
	v56(v57)
	v58 = *result
	tobool126 = (v58 & 1) != 0
	*retval = tobool126
	goto _return

sw_bb127:
	*result = 1
	v59 = *lexer_addr
	result_symbol128 = &v59.F1
	*result_symbol128 = 2
	v60 = *lexer_addr
	mark_end129 = &v60.F3
	v61 = *mark_end129
	v62 = *lexer_addr
	v61(v62)
	v63 = *result
	tobool130 = (v63 & 1) != 0
	*retval = tobool130
	goto _return

sw_bb131:
	*result = 1
	v64 = *lexer_addr
	result_symbol132 = &v64.F1
	*result_symbol132 = 3
	v65 = *lexer_addr
	mark_end133 = &v65.F3
	v66 = *mark_end133
	v67 = *lexer_addr
	v66(v67)
	v68 = *lookahead
	cmp134 = v68 == 46
	if cmp134 {
		goto if_then136
	} else {
		goto if_end137
	}

if_then136:
	*state_addr = 18
	goto next_state

if_end137:
	v69 = *lookahead
	cmp138 = v69 == 69
	if cmp138 {
		goto if_then143
	} else {
		goto lor_lhs_false140
	}

lor_lhs_false140:
	v70 = *lookahead
	cmp141 = v70 == 101
	if cmp141 {
		goto if_then143
	} else {
		goto if_end144
	}

if_then143:
	*state_addr = 11
	goto next_state

if_end144:
	v71 = *result
	tobool145 = (v71 & 1) != 0
	*retval = tobool145
	goto _return

sw_bb146:
	*result = 1
	v72 = *lexer_addr
	result_symbol147 = &v72.F1
	*result_symbol147 = 3
	v73 = *lexer_addr
	mark_end148 = &v73.F3
	v74 = *mark_end148
	v75 = *lexer_addr
	v74(v75)
	v76 = *lookahead
	cmp149 = v76 == 46
	if cmp149 {
		goto if_then151
	} else {
		goto if_end152
	}

if_then151:
	*state_addr = 18
	goto next_state

if_end152:
	v77 = *lookahead
	cmp153 = v77 == 69
	if cmp153 {
		goto if_then158
	} else {
		goto lor_lhs_false155
	}

lor_lhs_false155:
	v78 = *lookahead
	cmp156 = v78 == 101
	if cmp156 {
		goto if_then158
	} else {
		goto if_end159
	}

if_then158:
	*state_addr = 11
	goto next_state

if_end159:
	v79 = *lookahead
	cmp160 = 48 <= v79
	if cmp160 {
		goto land_lhs_true162
	} else {
		goto if_end166
	}

land_lhs_true162:
	v80 = *lookahead
	cmp163 = v80 <= 57
	if cmp163 {
		goto if_then165
	} else {
		goto if_end166
	}

if_then165:
	*state_addr = 17
	goto next_state

if_end166:
	v81 = *result
	tobool167 = (v81 & 1) != 0
	*retval = tobool167
	goto _return

sw_bb168:
	*result = 1
	v82 = *lexer_addr
	result_symbol169 = &v82.F1
	*result_symbol169 = 4
	v83 = *lexer_addr
	mark_end170 = &v83.F3
	v84 = *mark_end170
	v85 = *lexer_addr
	v84(v85)
	v86 = *lookahead
	cmp171 = v86 == 69
	if cmp171 {
		goto if_then176
	} else {
		goto lor_lhs_false173
	}

lor_lhs_false173:
	v87 = *lookahead
	cmp174 = v87 == 101
	if cmp174 {
		goto if_then176
	} else {
		goto if_end177
	}

if_then176:
	*state_addr = 11
	goto next_state

if_end177:
	v88 = *lookahead
	cmp178 = 48 <= v88
	if cmp178 {
		goto land_lhs_true180
	} else {
		goto if_end184
	}

land_lhs_true180:
	v89 = *lookahead
	cmp181 = v89 <= 57
	if cmp181 {
		goto if_then183
	} else {
		goto if_end184
	}

if_then183:
	*state_addr = 18
	goto next_state

if_end184:
	v90 = *result
	tobool185 = (v90 & 1) != 0
	*retval = tobool185
	goto _return

sw_bb186:
	*result = 1
	v91 = *lexer_addr
	result_symbol187 = &v91.F1
	*result_symbol187 = 4
	v92 = *lexer_addr
	mark_end188 = &v92.F3
	v93 = *mark_end188
	v94 = *lexer_addr
	v93(v94)
	v95 = *lookahead
	cmp189 = 48 <= v95
	if cmp189 {
		goto land_lhs_true191
	} else {
		goto if_end195
	}

land_lhs_true191:
	v96 = *lookahead
	cmp192 = v96 <= 57
	if cmp192 {
		goto if_then194
	} else {
		goto if_end195
	}

if_then194:
	*state_addr = 19
	goto next_state

if_end195:
	v97 = *result
	tobool196 = (v97 & 1) != 0
	*retval = tobool196
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v98 = *retval
	return v98
}

