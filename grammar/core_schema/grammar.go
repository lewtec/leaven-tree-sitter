package grammar_core_schema

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

var tree_sitter_core_schema_language TSLanguage = TSLanguage{14, 6, 0, 5, 0, 4, 2, 1, 0, 1, &ts_parse_table[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[11]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], (*TSLexerMode)(unsafe.Pointer(&ts_lex_modes)), ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], nil, nil, 0, 0, nil, nil, nil, TSLanguageMetadata{}}

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

var ts_lex_map [22]int16 = [22]int16{
	46, 6, 48, 37, 70, 2, 78, 16, 84, 13, 102, 17, 110, 29, 116, 26,
	126, 35, 43, 1, 45, 1,
}

func tree_sitter_core_schema() *TSLanguage {
	return &tree_sitter_core_schema_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v112, v113, v115, v117, v118, v120, v122, v123, v125, v127, v128, v130, v139, v140, v142, v149, v150, v152, v156, v157, v159, v167, v168, v170, v172, v173, v175, v181, v182, v184 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end310, mark_end314, mark_end318, mark_end348, mark_end370, mark_end381, mark_end404, mark_end408, mark_end426 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol309, result_symbol313, result_symbol317, result_symbol347, result_symbol369, result_symbol380, result_symbol403, result_symbol407, result_symbol425 *int16
	var lookahead, i, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, tobool20, cmp22, cmp26, cmp29, tobool33, cmp35, cmp39, tobool43, cmp45, cmp49, tobool53, cmp55, tobool59, cmp61, tobool65, cmp67, cmp71, cmp75, cmp79, cmp83, cmp86, tobool90, cmp92, cmp96, cmp100, cmp103, tobool107, cmp109, tobool113, cmp115, tobool119, cmp121, tobool125, cmp127, cmp131, tobool135, cmp137, tobool141, cmp143, cmp147, tobool151, cmp153, tobool157, cmp159, tobool163, cmp165, cmp169, tobool173, cmp175, tobool179, cmp181, tobool185, cmp187, tobool191, cmp193, tobool197, cmp199, tobool203, cmp205, tobool209, cmp211, tobool215, cmp217, tobool221, cmp223, tobool227, cmp229, tobool233, cmp235, tobool239, cmp241, tobool245, cmp247, tobool251, cmp253, cmp255, cmp259, cmp262, tobool266, cmp268, cmp271, tobool275, cmp277, cmp280, tobool284, cmp286, cmp289, cmp292, cmp295, cmp298, cmp301, tobool305, tobool307, tobool311, tobool315, cmp319, cmp323, cmp327, cmp331, cmp334, cmp338, cmp341, tobool345, cmp349, cmp353, cmp356, cmp360, cmp363, tobool367, cmp371, cmp374, tobool378, cmp382, cmp385, cmp388, cmp391, cmp394, cmp397, tobool401, tobool405, cmp409, cmp412, cmp416, cmp419, tobool423, cmp427, cmp430, tobool434, v188 bool
	var v3, frombool, v10, v20, v24, v27, v30, v32, v34, v41, v46, v48, v50, v52, v55, v57, v60, v62, v64, v67, v69, v71, v73, v75, v77, v79, v81, v83, v85, v87, v89, v91, v93, v98, v101, v104, v111, v116, v121, v126, v138, v148, v155, v166, v171, v180, v187 byte
	var v114, v119, v124, v129, v141, v151, v158, v169, v174, v183 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9, v13, v16 int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v21, v22, v23, v25, v26, v28, v29, v31, v33, v35, v36, v37, v38, v39, v40, v42, v43, v44, v45, v47, v49, v51, v53, v54, v56, v58, v59, v61, v63, v65, v66, v68, v70, v72, v74, v76, v78, v80, v82, v84, v86, v88, v90, v92, v94, v95, v96, v97, v99, v100, v102, v103, v105, v106, v107, v108, v109, v110, v131, v132, v133, v134, v135, v136, v137, v143, v144, v145, v146, v147, v153, v154, v160, v161, v162, v163, v164, v165, v176, v177, v178, v179, v185, v186 int32
	var conv4, idxprom, idxprom10 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, tobool20, v21, cmp22, v22, cmp26, v23, cmp29, v24, tobool33, v25, cmp35, v26, cmp39, v27, tobool43, v28, cmp45, v29, cmp49, v30, tobool53, v31, cmp55, v32, tobool59, v33, cmp61, v34, tobool65, v35, cmp67, v36, cmp71, v37, cmp75, v38, cmp79, v39, cmp83, v40, cmp86, v41, tobool90, v42, cmp92, v43, cmp96, v44, cmp100, v45, cmp103, v46, tobool107, v47, cmp109, v48, tobool113, v49, cmp115, v50, tobool119, v51, cmp121, v52, tobool125, v53, cmp127, v54, cmp131, v55, tobool135, v56, cmp137, v57, tobool141, v58, cmp143, v59, cmp147, v60, tobool151, v61, cmp153, v62, tobool157, v63, cmp159, v64, tobool163, v65, cmp165, v66, cmp169, v67, tobool173, v68, cmp175, v69, tobool179, v70, cmp181, v71, tobool185, v72, cmp187, v73, tobool191, v74, cmp193, v75, tobool197, v76, cmp199, v77, tobool203, v78, cmp205, v79, tobool209, v80, cmp211, v81, tobool215, v82, cmp217, v83, tobool221, v84, cmp223, v85, tobool227, v86, cmp229, v87, tobool233, v88, cmp235, v89, tobool239, v90, cmp241, v91, tobool245, v92, cmp247, v93, tobool251, v94, cmp253, v95, cmp255, v96, cmp259, v97, cmp262, v98, tobool266, v99, cmp268, v100, cmp271, v101, tobool275, v102, cmp277, v103, cmp280, v104, tobool284, v105, cmp286, v106, cmp289, v107, cmp292, v108, cmp295, v109, cmp298, v110, cmp301, v111, tobool305, v112, result_symbol, v113, mark_end, v114, v115, v116, tobool307, v117, result_symbol309, v118, mark_end310, v119, v120, v121, tobool311, v122, result_symbol313, v123, mark_end314, v124, v125, v126, tobool315, v127, result_symbol317, v128, mark_end318, v129, v130, v131, cmp319, v132, cmp323, v133, cmp327, v134, cmp331, v135, cmp334, v136, cmp338, v137, cmp341, v138, tobool345, v139, result_symbol347, v140, mark_end348, v141, v142, v143, cmp349, v144, cmp353, v145, cmp356, v146, cmp360, v147, cmp363, v148, tobool367, v149, result_symbol369, v150, mark_end370, v151, v152, v153, cmp371, v154, cmp374, v155, tobool378, v156, result_symbol380, v157, mark_end381, v158, v159, v160, cmp382, v161, cmp385, v162, cmp388, v163, cmp391, v164, cmp394, v165, cmp397, v166, tobool401, v167, result_symbol403, v168, mark_end404, v169, v170, v171, tobool405, v172, result_symbol407, v173, mark_end408, v174, v175, v176, cmp409, v177, cmp412, v178, cmp416, v179, cmp419, v180, tobool423, v181, result_symbol425, v182, mark_end426, v183, v184, v185, cmp427, v186, cmp430, v187, tobool434, v188

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
		goto sw_bb21
	case 2:
		goto sw_bb34
	case 3:
		goto sw_bb44
	case 4:
		goto sw_bb54
	case 5:
		goto sw_bb60
	case 6:
		goto sw_bb66
	case 7:
		goto sw_bb91
	case 8:
		goto sw_bb108
	case 9:
		goto sw_bb114
	case 10:
		goto sw_bb120
	case 11:
		goto sw_bb126
	case 12:
		goto sw_bb136
	case 13:
		goto sw_bb142
	case 14:
		goto sw_bb152
	case 15:
		goto sw_bb158
	case 16:
		goto sw_bb164
	case 17:
		goto sw_bb174
	case 18:
		goto sw_bb180
	case 19:
		goto sw_bb186
	case 20:
		goto sw_bb192
	case 21:
		goto sw_bb198
	case 22:
		goto sw_bb204
	case 23:
		goto sw_bb210
	case 24:
		goto sw_bb216
	case 25:
		goto sw_bb222
	case 26:
		goto sw_bb228
	case 27:
		goto sw_bb234
	case 28:
		goto sw_bb240
	case 29:
		goto sw_bb246
	case 30:
		goto sw_bb252
	case 31:
		goto sw_bb267
	case 32:
		goto sw_bb276
	case 33:
		goto sw_bb285
	case 34:
		goto sw_bb306
	case 35:
		goto sw_bb308
	case 36:
		goto sw_bb312
	case 37:
		goto sw_bb316
	case 38:
		goto sw_bb346
	case 39:
		goto sw_bb368
	case 40:
		goto sw_bb379
	case 41:
		goto sw_bb402
	case 42:
		goto sw_bb406
	case 43:
		goto sw_bb424
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
	*state_addr = 34
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
	cmp14 = 49 <= v18
	if cmp14 {
		goto land_lhs_true
	} else {
		goto if_end19
	}

land_lhs_true:
	v19 = *lookahead
	cmp16 = v19 <= 57
	if cmp16 {
		goto if_then18
	} else {
		goto if_end19
	}

if_then18:
	*state_addr = 38
	goto next_state

if_end19:
	v20 = *result
	tobool20 = (v20 & 1) != 0
	*retval = tobool20
	goto _return

sw_bb21:
	v21 = *lookahead
	cmp22 = v21 == 46
	if cmp22 {
		goto if_then24
	} else {
		goto if_end25
	}

if_then24:
	*state_addr = 7
	goto next_state

if_end25:
	v22 = *lookahead
	cmp26 = 48 <= v22
	if cmp26 {
		goto land_lhs_true28
	} else {
		goto if_end32
	}

land_lhs_true28:
	v23 = *lookahead
	cmp29 = v23 <= 57
	if cmp29 {
		goto if_then31
	} else {
		goto if_end32
	}

if_then31:
	*state_addr = 38
	goto next_state

if_end32:
	v24 = *result
	tobool33 = (v24 & 1) != 0
	*retval = tobool33
	goto _return

sw_bb34:
	v25 = *lookahead
	cmp35 = v25 == 65
	if cmp35 {
		goto if_then37
	} else {
		goto if_end38
	}

if_then37:
	*state_addr = 9
	goto next_state

if_end38:
	v26 = *lookahead
	cmp39 = v26 == 97
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*state_addr = 22
	goto next_state

if_end42:
	v27 = *result
	tobool43 = (v27 & 1) != 0
	*retval = tobool43
	goto _return

sw_bb44:
	v28 = *lookahead
	cmp45 = v28 == 65
	if cmp45 {
		goto if_then47
	} else {
		goto if_end48
	}

if_then47:
	*state_addr = 12
	goto next_state

if_end48:
	v29 = *lookahead
	cmp49 = v29 == 97
	if cmp49 {
		goto if_then51
	} else {
		goto if_end52
	}

if_then51:
	*state_addr = 12
	goto next_state

if_end52:
	v30 = *result
	tobool53 = (v30 & 1) != 0
	*retval = tobool53
	goto _return

sw_bb54:
	v31 = *lookahead
	cmp55 = v31 == 69
	if cmp55 {
		goto if_then57
	} else {
		goto if_end58
	}

if_then57:
	*state_addr = 36
	goto next_state

if_end58:
	v32 = *result
	tobool59 = (v32 & 1) != 0
	*retval = tobool59
	goto _return

sw_bb60:
	v33 = *lookahead
	cmp61 = v33 == 70
	if cmp61 {
		goto if_then63
	} else {
		goto if_end64
	}

if_then63:
	*state_addr = 41
	goto next_state

if_end64:
	v34 = *result
	tobool65 = (v34 & 1) != 0
	*retval = tobool65
	goto _return

sw_bb66:
	v35 = *lookahead
	cmp67 = v35 == 73
	if cmp67 {
		goto if_then69
	} else {
		goto if_end70
	}

if_then69:
	*state_addr = 11
	goto next_state

if_end70:
	v36 = *lookahead
	cmp71 = v36 == 78
	if cmp71 {
		goto if_then73
	} else {
		goto if_end74
	}

if_then73:
	*state_addr = 3
	goto next_state

if_end74:
	v37 = *lookahead
	cmp75 = v37 == 105
	if cmp75 {
		goto if_then77
	} else {
		goto if_end78
	}

if_then77:
	*state_addr = 24
	goto next_state

if_end78:
	v38 = *lookahead
	cmp79 = v38 == 110
	if cmp79 {
		goto if_then81
	} else {
		goto if_end82
	}

if_then81:
	*state_addr = 18
	goto next_state

if_end82:
	v39 = *lookahead
	cmp83 = 48 <= v39
	if cmp83 {
		goto land_lhs_true85
	} else {
		goto if_end89
	}

land_lhs_true85:
	v40 = *lookahead
	cmp86 = v40 <= 57
	if cmp86 {
		goto if_then88
	} else {
		goto if_end89
	}

if_then88:
	*state_addr = 42
	goto next_state

if_end89:
	v41 = *result
	tobool90 = (v41 & 1) != 0
	*retval = tobool90
	goto _return

sw_bb91:
	v42 = *lookahead
	cmp92 = v42 == 73
	if cmp92 {
		goto if_then94
	} else {
		goto if_end95
	}

if_then94:
	*state_addr = 11
	goto next_state

if_end95:
	v43 = *lookahead
	cmp96 = v43 == 105
	if cmp96 {
		goto if_then98
	} else {
		goto if_end99
	}

if_then98:
	*state_addr = 24
	goto next_state

if_end99:
	v44 = *lookahead
	cmp100 = 48 <= v44
	if cmp100 {
		goto land_lhs_true102
	} else {
		goto if_end106
	}

land_lhs_true102:
	v45 = *lookahead
	cmp103 = v45 <= 57
	if cmp103 {
		goto if_then105
	} else {
		goto if_end106
	}

if_then105:
	*state_addr = 42
	goto next_state

if_end106:
	v46 = *result
	tobool107 = (v46 & 1) != 0
	*retval = tobool107
	goto _return

sw_bb108:
	v47 = *lookahead
	cmp109 = v47 == 76
	if cmp109 {
		goto if_then111
	} else {
		goto if_end112
	}

if_then111:
	*state_addr = 35
	goto next_state

if_end112:
	v48 = *result
	tobool113 = (v48 & 1) != 0
	*retval = tobool113
	goto _return

sw_bb114:
	v49 = *lookahead
	cmp115 = v49 == 76
	if cmp115 {
		goto if_then117
	} else {
		goto if_end118
	}

if_then117:
	*state_addr = 14
	goto next_state

if_end118:
	v50 = *result
	tobool119 = (v50 & 1) != 0
	*retval = tobool119
	goto _return

sw_bb120:
	v51 = *lookahead
	cmp121 = v51 == 76
	if cmp121 {
		goto if_then123
	} else {
		goto if_end124
	}

if_then123:
	*state_addr = 8
	goto next_state

if_end124:
	v52 = *result
	tobool125 = (v52 & 1) != 0
	*retval = tobool125
	goto _return

sw_bb126:
	v53 = *lookahead
	cmp127 = v53 == 78
	if cmp127 {
		goto if_then129
	} else {
		goto if_end130
	}

if_then129:
	*state_addr = 5
	goto next_state

if_end130:
	v54 = *lookahead
	cmp131 = v54 == 110
	if cmp131 {
		goto if_then133
	} else {
		goto if_end134
	}

if_then133:
	*state_addr = 20
	goto next_state

if_end134:
	v55 = *result
	tobool135 = (v55 & 1) != 0
	*retval = tobool135
	goto _return

sw_bb136:
	v56 = *lookahead
	cmp137 = v56 == 78
	if cmp137 {
		goto if_then139
	} else {
		goto if_end140
	}

if_then139:
	*state_addr = 41
	goto next_state

if_end140:
	v57 = *result
	tobool141 = (v57 & 1) != 0
	*retval = tobool141
	goto _return

sw_bb142:
	v58 = *lookahead
	cmp143 = v58 == 82
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*state_addr = 15
	goto next_state

if_end146:
	v59 = *lookahead
	cmp147 = v59 == 114
	if cmp147 {
		goto if_then149
	} else {
		goto if_end150
	}

if_then149:
	*state_addr = 28
	goto next_state

if_end150:
	v60 = *result
	tobool151 = (v60 & 1) != 0
	*retval = tobool151
	goto _return

sw_bb152:
	v61 = *lookahead
	cmp153 = v61 == 83
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*state_addr = 4
	goto next_state

if_end156:
	v62 = *result
	tobool157 = (v62 & 1) != 0
	*retval = tobool157
	goto _return

sw_bb158:
	v63 = *lookahead
	cmp159 = v63 == 85
	if cmp159 {
		goto if_then161
	} else {
		goto if_end162
	}

if_then161:
	*state_addr = 4
	goto next_state

if_end162:
	v64 = *result
	tobool163 = (v64 & 1) != 0
	*retval = tobool163
	goto _return

sw_bb164:
	v65 = *lookahead
	cmp165 = v65 == 85
	if cmp165 {
		goto if_then167
	} else {
		goto if_end168
	}

if_then167:
	*state_addr = 10
	goto next_state

if_end168:
	v66 = *lookahead
	cmp169 = v66 == 117
	if cmp169 {
		goto if_then171
	} else {
		goto if_end172
	}

if_then171:
	*state_addr = 23
	goto next_state

if_end172:
	v67 = *result
	tobool173 = (v67 & 1) != 0
	*retval = tobool173
	goto _return

sw_bb174:
	v68 = *lookahead
	cmp175 = v68 == 97
	if cmp175 {
		goto if_then177
	} else {
		goto if_end178
	}

if_then177:
	*state_addr = 22
	goto next_state

if_end178:
	v69 = *result
	tobool179 = (v69 & 1) != 0
	*retval = tobool179
	goto _return

sw_bb180:
	v70 = *lookahead
	cmp181 = v70 == 97
	if cmp181 {
		goto if_then183
	} else {
		goto if_end184
	}

if_then183:
	*state_addr = 25
	goto next_state

if_end184:
	v71 = *result
	tobool185 = (v71 & 1) != 0
	*retval = tobool185
	goto _return

sw_bb186:
	v72 = *lookahead
	cmp187 = v72 == 101
	if cmp187 {
		goto if_then189
	} else {
		goto if_end190
	}

if_then189:
	*state_addr = 36
	goto next_state

if_end190:
	v73 = *result
	tobool191 = (v73 & 1) != 0
	*retval = tobool191
	goto _return

sw_bb192:
	v74 = *lookahead
	cmp193 = v74 == 102
	if cmp193 {
		goto if_then195
	} else {
		goto if_end196
	}

if_then195:
	*state_addr = 41
	goto next_state

if_end196:
	v75 = *result
	tobool197 = (v75 & 1) != 0
	*retval = tobool197
	goto _return

sw_bb198:
	v76 = *lookahead
	cmp199 = v76 == 108
	if cmp199 {
		goto if_then201
	} else {
		goto if_end202
	}

if_then201:
	*state_addr = 35
	goto next_state

if_end202:
	v77 = *result
	tobool203 = (v77 & 1) != 0
	*retval = tobool203
	goto _return

sw_bb204:
	v78 = *lookahead
	cmp205 = v78 == 108
	if cmp205 {
		goto if_then207
	} else {
		goto if_end208
	}

if_then207:
	*state_addr = 27
	goto next_state

if_end208:
	v79 = *result
	tobool209 = (v79 & 1) != 0
	*retval = tobool209
	goto _return

sw_bb210:
	v80 = *lookahead
	cmp211 = v80 == 108
	if cmp211 {
		goto if_then213
	} else {
		goto if_end214
	}

if_then213:
	*state_addr = 21
	goto next_state

if_end214:
	v81 = *result
	tobool215 = (v81 & 1) != 0
	*retval = tobool215
	goto _return

sw_bb216:
	v82 = *lookahead
	cmp217 = v82 == 110
	if cmp217 {
		goto if_then219
	} else {
		goto if_end220
	}

if_then219:
	*state_addr = 20
	goto next_state

if_end220:
	v83 = *result
	tobool221 = (v83 & 1) != 0
	*retval = tobool221
	goto _return

sw_bb222:
	v84 = *lookahead
	cmp223 = v84 == 110
	if cmp223 {
		goto if_then225
	} else {
		goto if_end226
	}

if_then225:
	*state_addr = 41
	goto next_state

if_end226:
	v85 = *result
	tobool227 = (v85 & 1) != 0
	*retval = tobool227
	goto _return

sw_bb228:
	v86 = *lookahead
	cmp229 = v86 == 114
	if cmp229 {
		goto if_then231
	} else {
		goto if_end232
	}

if_then231:
	*state_addr = 28
	goto next_state

if_end232:
	v87 = *result
	tobool233 = (v87 & 1) != 0
	*retval = tobool233
	goto _return

sw_bb234:
	v88 = *lookahead
	cmp235 = v88 == 115
	if cmp235 {
		goto if_then237
	} else {
		goto if_end238
	}

if_then237:
	*state_addr = 19
	goto next_state

if_end238:
	v89 = *result
	tobool239 = (v89 & 1) != 0
	*retval = tobool239
	goto _return

sw_bb240:
	v90 = *lookahead
	cmp241 = v90 == 117
	if cmp241 {
		goto if_then243
	} else {
		goto if_end244
	}

if_then243:
	*state_addr = 19
	goto next_state

if_end244:
	v91 = *result
	tobool245 = (v91 & 1) != 0
	*retval = tobool245
	goto _return

sw_bb246:
	v92 = *lookahead
	cmp247 = v92 == 117
	if cmp247 {
		goto if_then249
	} else {
		goto if_end250
	}

if_then249:
	*state_addr = 23
	goto next_state

if_end250:
	v93 = *result
	tobool251 = (v93 & 1) != 0
	*retval = tobool251
	goto _return

sw_bb252:
	v94 = *lookahead
	cmp253 = v94 == 43
	if cmp253 {
		goto if_then257
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v95 = *lookahead
	cmp255 = v95 == 45
	if cmp255 {
		goto if_then257
	} else {
		goto if_end258
	}

if_then257:
	*state_addr = 32
	goto next_state

if_end258:
	v96 = *lookahead
	cmp259 = 48 <= v96
	if cmp259 {
		goto land_lhs_true261
	} else {
		goto if_end265
	}

land_lhs_true261:
	v97 = *lookahead
	cmp262 = v97 <= 57
	if cmp262 {
		goto if_then264
	} else {
		goto if_end265
	}

if_then264:
	*state_addr = 43
	goto next_state

if_end265:
	v98 = *result
	tobool266 = (v98 & 1) != 0
	*retval = tobool266
	goto _return

sw_bb267:
	v99 = *lookahead
	cmp268 = 48 <= v99
	if cmp268 {
		goto land_lhs_true270
	} else {
		goto if_end274
	}

land_lhs_true270:
	v100 = *lookahead
	cmp271 = v100 <= 55
	if cmp271 {
		goto if_then273
	} else {
		goto if_end274
	}

if_then273:
	*state_addr = 39
	goto next_state

if_end274:
	v101 = *result
	tobool275 = (v101 & 1) != 0
	*retval = tobool275
	goto _return

sw_bb276:
	v102 = *lookahead
	cmp277 = 48 <= v102
	if cmp277 {
		goto land_lhs_true279
	} else {
		goto if_end283
	}

land_lhs_true279:
	v103 = *lookahead
	cmp280 = v103 <= 57
	if cmp280 {
		goto if_then282
	} else {
		goto if_end283
	}

if_then282:
	*state_addr = 43
	goto next_state

if_end283:
	v104 = *result
	tobool284 = (v104 & 1) != 0
	*retval = tobool284
	goto _return

sw_bb285:
	v105 = *lookahead
	cmp286 = 48 <= v105
	if cmp286 {
		goto land_lhs_true288
	} else {
		goto lor_lhs_false291
	}

land_lhs_true288:
	v106 = *lookahead
	cmp289 = v106 <= 57
	if cmp289 {
		goto if_then303
	} else {
		goto lor_lhs_false291
	}

lor_lhs_false291:
	v107 = *lookahead
	cmp292 = 65 <= v107
	if cmp292 {
		goto land_lhs_true294
	} else {
		goto lor_lhs_false297
	}

land_lhs_true294:
	v108 = *lookahead
	cmp295 = v108 <= 70
	if cmp295 {
		goto if_then303
	} else {
		goto lor_lhs_false297
	}

lor_lhs_false297:
	v109 = *lookahead
	cmp298 = 97 <= v109
	if cmp298 {
		goto land_lhs_true300
	} else {
		goto if_end304
	}

land_lhs_true300:
	v110 = *lookahead
	cmp301 = v110 <= 102
	if cmp301 {
		goto if_then303
	} else {
		goto if_end304
	}

if_then303:
	*state_addr = 40
	goto next_state

if_end304:
	v111 = *result
	tobool305 = (v111 & 1) != 0
	*retval = tobool305
	goto _return

sw_bb306:
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
	tobool307 = (v116 & 1) != 0
	*retval = tobool307
	goto _return

sw_bb308:
	*result = 1
	v117 = *lexer_addr
	result_symbol309 = &v117.F1
	*result_symbol309 = 1
	v118 = *lexer_addr
	mark_end310 = &v118.F3
	v119 = *mark_end310
	v120 = *lexer_addr
	v119(v120)
	v121 = *result
	tobool311 = (v121 & 1) != 0
	*retval = tobool311
	goto _return

sw_bb312:
	*result = 1
	v122 = *lexer_addr
	result_symbol313 = &v122.F1
	*result_symbol313 = 2
	v123 = *lexer_addr
	mark_end314 = &v123.F3
	v124 = *mark_end314
	v125 = *lexer_addr
	v124(v125)
	v126 = *result
	tobool315 = (v126 & 1) != 0
	*retval = tobool315
	goto _return

sw_bb316:
	*result = 1
	v127 = *lexer_addr
	result_symbol317 = &v127.F1
	*result_symbol317 = 3
	v128 = *lexer_addr
	mark_end318 = &v128.F3
	v129 = *mark_end318
	v130 = *lexer_addr
	v129(v130)
	v131 = *lookahead
	cmp319 = v131 == 46
	if cmp319 {
		goto if_then321
	} else {
		goto if_end322
	}

if_then321:
	*state_addr = 42
	goto next_state

if_end322:
	v132 = *lookahead
	cmp323 = v132 == 111
	if cmp323 {
		goto if_then325
	} else {
		goto if_end326
	}

if_then325:
	*state_addr = 31
	goto next_state

if_end326:
	v133 = *lookahead
	cmp327 = v133 == 120
	if cmp327 {
		goto if_then329
	} else {
		goto if_end330
	}

if_then329:
	*state_addr = 33
	goto next_state

if_end330:
	v134 = *lookahead
	cmp331 = v134 == 69
	if cmp331 {
		goto if_then336
	} else {
		goto lor_lhs_false333
	}

lor_lhs_false333:
	v135 = *lookahead
	cmp334 = v135 == 101
	if cmp334 {
		goto if_then336
	} else {
		goto if_end337
	}

if_then336:
	*state_addr = 30
	goto next_state

if_end337:
	v136 = *lookahead
	cmp338 = 48 <= v136
	if cmp338 {
		goto land_lhs_true340
	} else {
		goto if_end344
	}

land_lhs_true340:
	v137 = *lookahead
	cmp341 = v137 <= 57
	if cmp341 {
		goto if_then343
	} else {
		goto if_end344
	}

if_then343:
	*state_addr = 38
	goto next_state

if_end344:
	v138 = *result
	tobool345 = (v138 & 1) != 0
	*retval = tobool345
	goto _return

sw_bb346:
	*result = 1
	v139 = *lexer_addr
	result_symbol347 = &v139.F1
	*result_symbol347 = 3
	v140 = *lexer_addr
	mark_end348 = &v140.F3
	v141 = *mark_end348
	v142 = *lexer_addr
	v141(v142)
	v143 = *lookahead
	cmp349 = v143 == 46
	if cmp349 {
		goto if_then351
	} else {
		goto if_end352
	}

if_then351:
	*state_addr = 42
	goto next_state

if_end352:
	v144 = *lookahead
	cmp353 = v144 == 69
	if cmp353 {
		goto if_then358
	} else {
		goto lor_lhs_false355
	}

lor_lhs_false355:
	v145 = *lookahead
	cmp356 = v145 == 101
	if cmp356 {
		goto if_then358
	} else {
		goto if_end359
	}

if_then358:
	*state_addr = 30
	goto next_state

if_end359:
	v146 = *lookahead
	cmp360 = 48 <= v146
	if cmp360 {
		goto land_lhs_true362
	} else {
		goto if_end366
	}

land_lhs_true362:
	v147 = *lookahead
	cmp363 = v147 <= 57
	if cmp363 {
		goto if_then365
	} else {
		goto if_end366
	}

if_then365:
	*state_addr = 38
	goto next_state

if_end366:
	v148 = *result
	tobool367 = (v148 & 1) != 0
	*retval = tobool367
	goto _return

sw_bb368:
	*result = 1
	v149 = *lexer_addr
	result_symbol369 = &v149.F1
	*result_symbol369 = 3
	v150 = *lexer_addr
	mark_end370 = &v150.F3
	v151 = *mark_end370
	v152 = *lexer_addr
	v151(v152)
	v153 = *lookahead
	cmp371 = 48 <= v153
	if cmp371 {
		goto land_lhs_true373
	} else {
		goto if_end377
	}

land_lhs_true373:
	v154 = *lookahead
	cmp374 = v154 <= 55
	if cmp374 {
		goto if_then376
	} else {
		goto if_end377
	}

if_then376:
	*state_addr = 39
	goto next_state

if_end377:
	v155 = *result
	tobool378 = (v155 & 1) != 0
	*retval = tobool378
	goto _return

sw_bb379:
	*result = 1
	v156 = *lexer_addr
	result_symbol380 = &v156.F1
	*result_symbol380 = 3
	v157 = *lexer_addr
	mark_end381 = &v157.F3
	v158 = *mark_end381
	v159 = *lexer_addr
	v158(v159)
	v160 = *lookahead
	cmp382 = 48 <= v160
	if cmp382 {
		goto land_lhs_true384
	} else {
		goto lor_lhs_false387
	}

land_lhs_true384:
	v161 = *lookahead
	cmp385 = v161 <= 57
	if cmp385 {
		goto if_then399
	} else {
		goto lor_lhs_false387
	}

lor_lhs_false387:
	v162 = *lookahead
	cmp388 = 65 <= v162
	if cmp388 {
		goto land_lhs_true390
	} else {
		goto lor_lhs_false393
	}

land_lhs_true390:
	v163 = *lookahead
	cmp391 = v163 <= 70
	if cmp391 {
		goto if_then399
	} else {
		goto lor_lhs_false393
	}

lor_lhs_false393:
	v164 = *lookahead
	cmp394 = 97 <= v164
	if cmp394 {
		goto land_lhs_true396
	} else {
		goto if_end400
	}

land_lhs_true396:
	v165 = *lookahead
	cmp397 = v165 <= 102
	if cmp397 {
		goto if_then399
	} else {
		goto if_end400
	}

if_then399:
	*state_addr = 40
	goto next_state

if_end400:
	v166 = *result
	tobool401 = (v166 & 1) != 0
	*retval = tobool401
	goto _return

sw_bb402:
	*result = 1
	v167 = *lexer_addr
	result_symbol403 = &v167.F1
	*result_symbol403 = 4
	v168 = *lexer_addr
	mark_end404 = &v168.F3
	v169 = *mark_end404
	v170 = *lexer_addr
	v169(v170)
	v171 = *result
	tobool405 = (v171 & 1) != 0
	*retval = tobool405
	goto _return

sw_bb406:
	*result = 1
	v172 = *lexer_addr
	result_symbol407 = &v172.F1
	*result_symbol407 = 4
	v173 = *lexer_addr
	mark_end408 = &v173.F3
	v174 = *mark_end408
	v175 = *lexer_addr
	v174(v175)
	v176 = *lookahead
	cmp409 = v176 == 69
	if cmp409 {
		goto if_then414
	} else {
		goto lor_lhs_false411
	}

lor_lhs_false411:
	v177 = *lookahead
	cmp412 = v177 == 101
	if cmp412 {
		goto if_then414
	} else {
		goto if_end415
	}

if_then414:
	*state_addr = 30
	goto next_state

if_end415:
	v178 = *lookahead
	cmp416 = 48 <= v178
	if cmp416 {
		goto land_lhs_true418
	} else {
		goto if_end422
	}

land_lhs_true418:
	v179 = *lookahead
	cmp419 = v179 <= 57
	if cmp419 {
		goto if_then421
	} else {
		goto if_end422
	}

if_then421:
	*state_addr = 42
	goto next_state

if_end422:
	v180 = *result
	tobool423 = (v180 & 1) != 0
	*retval = tobool423
	goto _return

sw_bb424:
	*result = 1
	v181 = *lexer_addr
	result_symbol425 = &v181.F1
	*result_symbol425 = 4
	v182 = *lexer_addr
	mark_end426 = &v182.F3
	v183 = *mark_end426
	v184 = *lexer_addr
	v183(v184)
	v185 = *lookahead
	cmp427 = 48 <= v185
	if cmp427 {
		goto land_lhs_true429
	} else {
		goto if_end433
	}

land_lhs_true429:
	v186 = *lookahead
	cmp430 = v186 <= 57
	if cmp430 {
		goto if_then432
	} else {
		goto if_end433
	}

if_then432:
	*state_addr = 43
	goto next_state

if_end433:
	v187 = *result
	tobool434 = (v187 & 1) != 0
	*retval = tobool434
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v188 = *retval
	return v188
}

