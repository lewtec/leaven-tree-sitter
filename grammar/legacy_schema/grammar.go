package grammar_legacy_schema

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

var tree_sitter_legacy_schema_language TSLanguage = TSLanguage{14, 7, 0, 6, 0, 4, 2, 1, 0, 1, &ts_parse_table[0][0], &ts_small_parse_table[0], &ts_small_parse_table_map[0], &(*[11]TSParseActionEntry)(unsafe.Pointer(&ts_parse_actions))[0], &ts_symbol_names[0], nil, nil, nil, &ts_symbol_metadata[0], &ts_symbol_map[0], &ts_non_terminal_alias_map[0], &ts_alias_sequences[0][0], (*TSLexerMode)(unsafe.Pointer(&ts_lex_modes)), ts_lex, nil, 0, anon_2{}, &ts_primary_state_ids[0], nil, nil, 0, 0, nil, nil, nil, TSLanguageMetadata{}}

var ts_parse_table [2][7]int16 = [2][7]int16{[7]int16{1, 1, 1, 1, 1, 1, 0}, [7]int16{0, 3, 5, 5, 3, 3, 3}}

var ts_small_parse_table [8]int16 = [8]int16{1, 7, 1, 0, 1, 9, 1, 0}

var ts_small_parse_table_map [2]int32 = [2]int32{0, 4}

var ts_symbol_names [7]*byte = [7]*byte{&_str[0], &_str_2[0], &_str_3[0], &_str_4[0], &_str_5[0], &_str_6[0], &_str_7[0]}

var ts_symbol_metadata [7]TSSymbolMetadata = [7]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}}

var ts_symbol_map [7]int16 = [7]int16{0, 1, 2, 3, 4, 5, 6}

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
}{anon_1{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_0{1, 1, 6, 0, 0}}}, struct {
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

var _str_6 [10]byte = [10]byte{116, 105, 109, 101, 115, 116, 97, 109, 112, 0}

var _str_7 [7]byte = [7]byte{115, 99, 97, 108, 97, 114, 0}

var ts_lex_map [30]int16 = [30]int16{
	46, 89, 48, 75, 70, 13, 78, 70, 79, 17, 84, 24, 89, 69, 102, 29,
	110, 72, 111, 33, 116, 40, 121, 71, 126, 67, 43, 7, 45, 7,
}

func tree_sitter_legacy_schema() *TSLanguage {
	return &tree_sitter_legacy_schema_language
}

func ts_lex(lexer *TSLexer, state int16) bool {
	var lexer_addr **TSLexer
	var v0, v2, v4, v6, v8, v233, v234, v236, v238, v239, v241, v243, v244, v246, v248, v249, v251, v255, v256, v258, v264, v265, v267, v270, v271, v273, v277, v278, v280, v290, v291, v293, v301, v302, v304, v315, v316, v318, v327, v328, v330, v339, v340, v342, v353, v354, v356, v365, v366, v368, v375, v376, v378, v385, v386, v388, v395, v396, v398, v405, v406, v408, v412, v413, v415, v421, v422, v424, v429, v430, v432, v441, v442, v444, v446, v447, v449, v460, v461, v463, v472, v473, v475, v482, v483, v485, v489, v490, v492, v497, v498, v500, v502, v503, v505, v513, v514, v516, v519, v520, v522, v527, v528, v530, v539, v540, v542 *TSLexer
	var retval *bool
	var result, skip, eof *byte
	var mark_end, mark_end694, mark_end698, mark_end702, mark_end714, mark_end733, mark_end741, mark_end753, mark_end786, mark_end812, mark_end850, mark_end880, mark_end910, mark_end947, mark_end976, mark_end999, mark_end1022, mark_end1045, mark_end1067, mark_end1079, mark_end1098, mark_end1112, mark_end1138, mark_end1142, mark_end1179, mark_end1208, mark_end1229, mark_end1240, mark_end1254, mark_end1258, mark_end1284, mark_end1292, mark_end1307, mark_end1336 *func(*TSLexer)
	var eof2 *func(*TSLexer) bool
	var advance *func(*TSLexer, bool)
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol693, result_symbol697, result_symbol701, result_symbol713, result_symbol732, result_symbol740, result_symbol752, result_symbol785, result_symbol811, result_symbol849, result_symbol879, result_symbol909, result_symbol946, result_symbol975, result_symbol998, result_symbol1021, result_symbol1044, result_symbol1066, result_symbol1078, result_symbol1097, result_symbol1111, result_symbol1137, result_symbol1141, result_symbol1178, result_symbol1207, result_symbol1228, result_symbol1239, result_symbol1253, result_symbol1257, result_symbol1283, result_symbol1291, result_symbol1306, result_symbol1335 *int16
	var lookahead, i, lookahead1 *int32
	var tobool, call, tobool3, cmp, cmp7, cmp14, cmp16, tobool20, cmp22, cmp26, cmp30, cmp34, cmp37, cmp39, tobool43, cmp45, cmp49, cmp52, tobool56, cmp58, tobool62, cmp64, cmp68, cmp72, cmp76, cmp79, tobool83, cmp85, cmp89, cmp93, cmp97, cmp100, tobool104, cmp106, cmp110, cmp114, cmp117, cmp120, tobool124, cmp126, cmp130, cmp134, cmp137, tobool141, cmp143, cmp147, tobool151, cmp153, cmp157, cmp161, cmp164, tobool168, cmp170, tobool174, cmp176, cmp180, cmp183, tobool187, cmp189, tobool193, cmp195, cmp199, tobool203, cmp205, cmp209, tobool213, cmp215, tobool219, cmp221, tobool225, cmp227, cmp231, cmp235, cmp238, tobool242, cmp244, tobool248, cmp250, tobool254, cmp256, tobool260, cmp262, tobool266, cmp268, tobool272, cmp274, cmp278, tobool282, cmp284, cmp288, tobool292, cmp294, tobool298, cmp300, tobool304, cmp306, tobool310, cmp312, cmp316, cmp319, tobool323, cmp325, tobool329, cmp331, tobool335, cmp337, tobool341, cmp343, tobool347, cmp349, cmp353, tobool357, cmp359, tobool363, cmp365, tobool369, cmp371, tobool375, cmp377, tobool381, cmp383, tobool387, cmp389, tobool393, cmp395, tobool399, cmp401, tobool405, cmp407, tobool411, cmp413, tobool417, cmp419, cmp422, cmp426, cmp429, tobool433, cmp435, cmp438, cmp442, cmp445, cmp449, cmp452, tobool456, cmp458, cmp461, cmp465, cmp468, cmp472, cmp475, tobool479, cmp481, cmp484, cmp488, cmp491, tobool495, cmp497, cmp500, tobool504, cmp506, cmp509, cmp512, tobool516, cmp518, cmp521, cmp525, cmp528, tobool532, cmp534, cmp537, cmp541, cmp544, tobool548, cmp550, cmp553, tobool557, cmp559, cmp562, tobool566, cmp568, cmp571, tobool575, cmp577, cmp580, tobool584, cmp586, cmp589, tobool593, cmp595, cmp598, tobool602, cmp604, cmp607, tobool611, cmp613, cmp616, tobool620, cmp622, cmp625, tobool629, cmp631, cmp634, tobool638, cmp640, cmp643, tobool647, cmp649, cmp652, tobool656, cmp658, cmp661, tobool665, cmp667, cmp670, cmp673, cmp676, cmp679, cmp682, cmp685, tobool689, tobool691, tobool695, tobool699, cmp703, cmp707, tobool711, cmp715, cmp719, cmp723, cmp726, tobool730, cmp734, tobool738, cmp742, cmp746, tobool750, cmp754, cmp758, cmp762, cmp766, cmp769, cmp773, cmp776, cmp779, tobool783, cmp787, cmp791, cmp795, cmp799, cmp802, cmp805, tobool809, cmp813, cmp817, cmp821, cmp825, cmp829, cmp833, cmp836, cmp840, cmp843, tobool847, cmp851, cmp855, cmp859, cmp863, cmp866, cmp870, cmp873, tobool877, cmp881, cmp885, cmp889, cmp893, cmp896, cmp900, cmp903, tobool907, cmp911, cmp915, cmp919, cmp923, cmp927, cmp930, cmp934, cmp937, cmp940, tobool944, cmp948, cmp952, cmp956, cmp959, cmp963, cmp966, cmp969, tobool973, cmp977, cmp981, cmp985, cmp989, cmp992, tobool996, cmp1000, cmp1004, cmp1008, cmp1012, cmp1015, tobool1019, cmp1023, cmp1027, cmp1031, cmp1035, cmp1038, tobool1042, cmp1046, cmp1050, cmp1054, cmp1057, cmp1060, tobool1064, cmp1068, cmp1072, tobool1076, cmp1080, cmp1084, cmp1088, cmp1091, tobool1095, cmp1099, cmp1102, cmp1105, tobool1109, cmp1113, cmp1116, cmp1119, cmp1122, cmp1125, cmp1128, cmp1131, tobool1135, tobool1139, cmp1143, cmp1147, cmp1151, cmp1155, cmp1159, cmp1162, cmp1166, cmp1169, cmp1172, tobool1176, cmp1180, cmp1184, cmp1188, cmp1191, cmp1195, cmp1198, cmp1201, tobool1205, cmp1209, cmp1212, cmp1216, cmp1219, cmp1222, tobool1226, cmp1230, cmp1233, tobool1237, cmp1241, cmp1244, cmp1247, tobool1251, tobool1255, cmp1259, cmp1263, cmp1267, cmp1270, cmp1274, cmp1277, tobool1281, cmp1285, tobool1289, cmp1293, cmp1297, cmp1300, tobool1304, cmp1308, cmp1312, cmp1315, cmp1319, cmp1322, cmp1326, cmp1329, tobool1333, cmp1337, cmp1340, cmp1344, cmp1347, tobool1351, v548 bool
	var v3, frombool, v10, v20, v27, v31, v33, v39, v45, v51, v56, v59, v64, v66, v70, v72, v75, v78, v80, v82, v87, v89, v91, v93, v95, v97, v100, v103, v105, v107, v109, v113, v115, v117, v119, v121, v124, v126, v128, v130, v132, v134, v136, v138, v140, v142, v144, v149, v156, v163, v168, v171, v175, v180, v185, v188, v191, v194, v197, v200, v203, v206, v209, v212, v215, v218, v221, v224, v232, v237, v242, v247, v254, v263, v269, v276, v289, v300, v314, v326, v338, v352, v364, v374, v384, v394, v404, v411, v420, v428, v440, v445, v459, v471, v481, v488, v496, v501, v512, v518, v526, v538, v547 byte
	var v235, v240, v245, v250, v257, v266, v272, v279, v292, v303, v317, v329, v341, v355, v367, v377, v387, v397, v407, v414, v423, v431, v443, v448, v462, v474, v484, v491, v499, v504, v515, v521, v529, v541 func(*TSLexer)
	var v7 func(*TSLexer) bool
	var v1 func(*TSLexer, bool)
	var v9, v13, v16 int16
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v21, v22, v23, v24, v25, v26, v28, v29, v30, v32, v34, v35, v36, v37, v38, v40, v41, v42, v43, v44, v46, v47, v48, v49, v50, v52, v53, v54, v55, v57, v58, v60, v61, v62, v63, v65, v67, v68, v69, v71, v73, v74, v76, v77, v79, v81, v83, v84, v85, v86, v88, v90, v92, v94, v96, v98, v99, v101, v102, v104, v106, v108, v110, v111, v112, v114, v116, v118, v120, v122, v123, v125, v127, v129, v131, v133, v135, v137, v139, v141, v143, v145, v146, v147, v148, v150, v151, v152, v153, v154, v155, v157, v158, v159, v160, v161, v162, v164, v165, v166, v167, v169, v170, v172, v173, v174, v176, v177, v178, v179, v181, v182, v183, v184, v186, v187, v189, v190, v192, v193, v195, v196, v198, v199, v201, v202, v204, v205, v207, v208, v210, v211, v213, v214, v216, v217, v219, v220, v222, v223, v225, v226, v227, v228, v229, v230, v231, v252, v253, v259, v260, v261, v262, v268, v274, v275, v281, v282, v283, v284, v285, v286, v287, v288, v294, v295, v296, v297, v298, v299, v305, v306, v307, v308, v309, v310, v311, v312, v313, v319, v320, v321, v322, v323, v324, v325, v331, v332, v333, v334, v335, v336, v337, v343, v344, v345, v346, v347, v348, v349, v350, v351, v357, v358, v359, v360, v361, v362, v363, v369, v370, v371, v372, v373, v379, v380, v381, v382, v383, v389, v390, v391, v392, v393, v399, v400, v401, v402, v403, v409, v410, v416, v417, v418, v419, v425, v426, v427, v433, v434, v435, v436, v437, v438, v439, v450, v451, v452, v453, v454, v455, v456, v457, v458, v464, v465, v466, v467, v468, v469, v470, v476, v477, v478, v479, v480, v486, v487, v493, v494, v495, v506, v507, v508, v509, v510, v511, v517, v523, v524, v525, v531, v532, v533, v534, v535, v536, v537, v543, v544, v545, v546 int32
	var conv4, idxprom, idxprom10 int64

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, tobool, v4, lookahead1, v5, v6, eof2, v7, v8, call, frombool, v9, conv, v10, tobool3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, tobool20, v21, cmp22, v22, cmp26, v23, cmp30, v24, cmp34, v25, cmp37, v26, cmp39, v27, tobool43, v28, cmp45, v29, cmp49, v30, cmp52, v31, tobool56, v32, cmp58, v33, tobool62, v34, cmp64, v35, cmp68, v36, cmp72, v37, cmp76, v38, cmp79, v39, tobool83, v40, cmp85, v41, cmp89, v42, cmp93, v43, cmp97, v44, cmp100, v45, tobool104, v46, cmp106, v47, cmp110, v48, cmp114, v49, cmp117, v50, cmp120, v51, tobool124, v52, cmp126, v53, cmp130, v54, cmp134, v55, cmp137, v56, tobool141, v57, cmp143, v58, cmp147, v59, tobool151, v60, cmp153, v61, cmp157, v62, cmp161, v63, cmp164, v64, tobool168, v65, cmp170, v66, tobool174, v67, cmp176, v68, cmp180, v69, cmp183, v70, tobool187, v71, cmp189, v72, tobool193, v73, cmp195, v74, cmp199, v75, tobool203, v76, cmp205, v77, cmp209, v78, tobool213, v79, cmp215, v80, tobool219, v81, cmp221, v82, tobool225, v83, cmp227, v84, cmp231, v85, cmp235, v86, cmp238, v87, tobool242, v88, cmp244, v89, tobool248, v90, cmp250, v91, tobool254, v92, cmp256, v93, tobool260, v94, cmp262, v95, tobool266, v96, cmp268, v97, tobool272, v98, cmp274, v99, cmp278, v100, tobool282, v101, cmp284, v102, cmp288, v103, tobool292, v104, cmp294, v105, tobool298, v106, cmp300, v107, tobool304, v108, cmp306, v109, tobool310, v110, cmp312, v111, cmp316, v112, cmp319, v113, tobool323, v114, cmp325, v115, tobool329, v116, cmp331, v117, tobool335, v118, cmp337, v119, tobool341, v120, cmp343, v121, tobool347, v122, cmp349, v123, cmp353, v124, tobool357, v125, cmp359, v126, tobool363, v127, cmp365, v128, tobool369, v129, cmp371, v130, tobool375, v131, cmp377, v132, tobool381, v133, cmp383, v134, tobool387, v135, cmp389, v136, tobool393, v137, cmp395, v138, tobool399, v139, cmp401, v140, tobool405, v141, cmp407, v142, tobool411, v143, cmp413, v144, tobool417, v145, cmp419, v146, cmp422, v147, cmp426, v148, cmp429, v149, tobool433, v150, cmp435, v151, cmp438, v152, cmp442, v153, cmp445, v154, cmp449, v155, cmp452, v156, tobool456, v157, cmp458, v158, cmp461, v159, cmp465, v160, cmp468, v161, cmp472, v162, cmp475, v163, tobool479, v164, cmp481, v165, cmp484, v166, cmp488, v167, cmp491, v168, tobool495, v169, cmp497, v170, cmp500, v171, tobool504, v172, cmp506, v173, cmp509, v174, cmp512, v175, tobool516, v176, cmp518, v177, cmp521, v178, cmp525, v179, cmp528, v180, tobool532, v181, cmp534, v182, cmp537, v183, cmp541, v184, cmp544, v185, tobool548, v186, cmp550, v187, cmp553, v188, tobool557, v189, cmp559, v190, cmp562, v191, tobool566, v192, cmp568, v193, cmp571, v194, tobool575, v195, cmp577, v196, cmp580, v197, tobool584, v198, cmp586, v199, cmp589, v200, tobool593, v201, cmp595, v202, cmp598, v203, tobool602, v204, cmp604, v205, cmp607, v206, tobool611, v207, cmp613, v208, cmp616, v209, tobool620, v210, cmp622, v211, cmp625, v212, tobool629, v213, cmp631, v214, cmp634, v215, tobool638, v216, cmp640, v217, cmp643, v218, tobool647, v219, cmp649, v220, cmp652, v221, tobool656, v222, cmp658, v223, cmp661, v224, tobool665, v225, cmp667, v226, cmp670, v227, cmp673, v228, cmp676, v229, cmp679, v230, cmp682, v231, cmp685, v232, tobool689, v233, result_symbol, v234, mark_end, v235, v236, v237, tobool691, v238, result_symbol693, v239, mark_end694, v240, v241, v242, tobool695, v243, result_symbol697, v244, mark_end698, v245, v246, v247, tobool699, v248, result_symbol701, v249, mark_end702, v250, v251, v252, cmp703, v253, cmp707, v254, tobool711, v255, result_symbol713, v256, mark_end714, v257, v258, v259, cmp715, v260, cmp719, v261, cmp723, v262, cmp726, v263, tobool730, v264, result_symbol732, v265, mark_end733, v266, v267, v268, cmp734, v269, tobool738, v270, result_symbol740, v271, mark_end741, v272, v273, v274, cmp742, v275, cmp746, v276, tobool750, v277, result_symbol752, v278, mark_end753, v279, v280, v281, cmp754, v282, cmp758, v283, cmp762, v284, cmp766, v285, cmp769, v286, cmp773, v287, cmp776, v288, cmp779, v289, tobool783, v290, result_symbol785, v291, mark_end786, v292, v293, v294, cmp787, v295, cmp791, v296, cmp795, v297, cmp799, v298, cmp802, v299, cmp805, v300, tobool809, v301, result_symbol811, v302, mark_end812, v303, v304, v305, cmp813, v306, cmp817, v307, cmp821, v308, cmp825, v309, cmp829, v310, cmp833, v311, cmp836, v312, cmp840, v313, cmp843, v314, tobool847, v315, result_symbol849, v316, mark_end850, v317, v318, v319, cmp851, v320, cmp855, v321, cmp859, v322, cmp863, v323, cmp866, v324, cmp870, v325, cmp873, v326, tobool877, v327, result_symbol879, v328, mark_end880, v329, v330, v331, cmp881, v332, cmp885, v333, cmp889, v334, cmp893, v335, cmp896, v336, cmp900, v337, cmp903, v338, tobool907, v339, result_symbol909, v340, mark_end910, v341, v342, v343, cmp911, v344, cmp915, v345, cmp919, v346, cmp923, v347, cmp927, v348, cmp930, v349, cmp934, v350, cmp937, v351, cmp940, v352, tobool944, v353, result_symbol946, v354, mark_end947, v355, v356, v357, cmp948, v358, cmp952, v359, cmp956, v360, cmp959, v361, cmp963, v362, cmp966, v363, cmp969, v364, tobool973, v365, result_symbol975, v366, mark_end976, v367, v368, v369, cmp977, v370, cmp981, v371, cmp985, v372, cmp989, v373, cmp992, v374, tobool996, v375, result_symbol998, v376, mark_end999, v377, v378, v379, cmp1000, v380, cmp1004, v381, cmp1008, v382, cmp1012, v383, cmp1015, v384, tobool1019, v385, result_symbol1021, v386, mark_end1022, v387, v388, v389, cmp1023, v390, cmp1027, v391, cmp1031, v392, cmp1035, v393, cmp1038, v394, tobool1042, v395, result_symbol1044, v396, mark_end1045, v397, v398, v399, cmp1046, v400, cmp1050, v401, cmp1054, v402, cmp1057, v403, cmp1060, v404, tobool1064, v405, result_symbol1066, v406, mark_end1067, v407, v408, v409, cmp1068, v410, cmp1072, v411, tobool1076, v412, result_symbol1078, v413, mark_end1079, v414, v415, v416, cmp1080, v417, cmp1084, v418, cmp1088, v419, cmp1091, v420, tobool1095, v421, result_symbol1097, v422, mark_end1098, v423, v424, v425, cmp1099, v426, cmp1102, v427, cmp1105, v428, tobool1109, v429, result_symbol1111, v430, mark_end1112, v431, v432, v433, cmp1113, v434, cmp1116, v435, cmp1119, v436, cmp1122, v437, cmp1125, v438, cmp1128, v439, cmp1131, v440, tobool1135, v441, result_symbol1137, v442, mark_end1138, v443, v444, v445, tobool1139, v446, result_symbol1141, v447, mark_end1142, v448, v449, v450, cmp1143, v451, cmp1147, v452, cmp1151, v453, cmp1155, v454, cmp1159, v455, cmp1162, v456, cmp1166, v457, cmp1169, v458, cmp1172, v459, tobool1176, v460, result_symbol1178, v461, mark_end1179, v462, v463, v464, cmp1180, v465, cmp1184, v466, cmp1188, v467, cmp1191, v468, cmp1195, v469, cmp1198, v470, cmp1201, v471, tobool1205, v472, result_symbol1207, v473, mark_end1208, v474, v475, v476, cmp1209, v477, cmp1212, v478, cmp1216, v479, cmp1219, v480, cmp1222, v481, tobool1226, v482, result_symbol1228, v483, mark_end1229, v484, v485, v486, cmp1230, v487, cmp1233, v488, tobool1237, v489, result_symbol1239, v490, mark_end1240, v491, v492, v493, cmp1241, v494, cmp1244, v495, cmp1247, v496, tobool1251, v497, result_symbol1253, v498, mark_end1254, v499, v500, v501, tobool1255, v502, result_symbol1257, v503, mark_end1258, v504, v505, v506, cmp1259, v507, cmp1263, v508, cmp1267, v509, cmp1270, v510, cmp1274, v511, cmp1277, v512, tobool1281, v513, result_symbol1283, v514, mark_end1284, v515, v516, v517, cmp1285, v518, tobool1289, v519, result_symbol1291, v520, mark_end1292, v521, v522, v523, cmp1293, v524, cmp1297, v525, cmp1300, v526, tobool1304, v527, result_symbol1306, v528, mark_end1307, v529, v530, v531, cmp1308, v532, cmp1312, v533, cmp1315, v534, cmp1319, v535, cmp1322, v536, cmp1326, v537, cmp1329, v538, tobool1333, v539, result_symbol1335, v540, mark_end1336, v541, v542, v543, cmp1337, v544, cmp1340, v545, cmp1344, v546, cmp1347, v547, tobool1351, v548

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
		goto sw_bb44
	case 3:
		goto sw_bb57
	case 4:
		goto sw_bb63
	case 5:
		goto sw_bb84
	case 6:
		goto sw_bb105
	case 7:
		goto sw_bb125
	case 8:
		goto sw_bb142
	case 9:
		goto sw_bb152
	case 10:
		goto sw_bb169
	case 11:
		goto sw_bb175
	case 12:
		goto sw_bb188
	case 13:
		goto sw_bb194
	case 14:
		goto sw_bb204
	case 15:
		goto sw_bb214
	case 16:
		goto sw_bb220
	case 17:
		goto sw_bb226
	case 18:
		goto sw_bb243
	case 19:
		goto sw_bb249
	case 20:
		goto sw_bb255
	case 21:
		goto sw_bb261
	case 22:
		goto sw_bb267
	case 23:
		goto sw_bb273
	case 24:
		goto sw_bb283
	case 25:
		goto sw_bb293
	case 26:
		goto sw_bb299
	case 27:
		goto sw_bb305
	case 28:
		goto sw_bb311
	case 29:
		goto sw_bb324
	case 30:
		goto sw_bb330
	case 31:
		goto sw_bb336
	case 32:
		goto sw_bb342
	case 33:
		goto sw_bb348
	case 34:
		goto sw_bb358
	case 35:
		goto sw_bb364
	case 36:
		goto sw_bb370
	case 37:
		goto sw_bb376
	case 38:
		goto sw_bb382
	case 39:
		goto sw_bb388
	case 40:
		goto sw_bb394
	case 41:
		goto sw_bb400
	case 42:
		goto sw_bb406
	case 43:
		goto sw_bb412
	case 44:
		goto sw_bb418
	case 45:
		goto sw_bb434
	case 46:
		goto sw_bb457
	case 47:
		goto sw_bb480
	case 48:
		goto sw_bb496
	case 49:
		goto sw_bb505
	case 50:
		goto sw_bb517
	case 51:
		goto sw_bb533
	case 52:
		goto sw_bb549
	case 53:
		goto sw_bb558
	case 54:
		goto sw_bb567
	case 55:
		goto sw_bb576
	case 56:
		goto sw_bb585
	case 57:
		goto sw_bb594
	case 58:
		goto sw_bb603
	case 59:
		goto sw_bb612
	case 60:
		goto sw_bb621
	case 61:
		goto sw_bb630
	case 62:
		goto sw_bb639
	case 63:
		goto sw_bb648
	case 64:
		goto sw_bb657
	case 65:
		goto sw_bb666
	case 66:
		goto sw_bb690
	case 67:
		goto sw_bb692
	case 68:
		goto sw_bb696
	case 69:
		goto sw_bb700
	case 70:
		goto sw_bb712
	case 71:
		goto sw_bb731
	case 72:
		goto sw_bb739
	case 73:
		goto sw_bb751
	case 74:
		goto sw_bb784
	case 75:
		goto sw_bb810
	case 76:
		goto sw_bb848
	case 77:
		goto sw_bb878
	case 78:
		goto sw_bb908
	case 79:
		goto sw_bb945
	case 80:
		goto sw_bb974
	case 81:
		goto sw_bb997
	case 82:
		goto sw_bb1020
	case 83:
		goto sw_bb1043
	case 84:
		goto sw_bb1065
	case 85:
		goto sw_bb1077
	case 86:
		goto sw_bb1096
	case 87:
		goto sw_bb1110
	case 88:
		goto sw_bb1136
	case 89:
		goto sw_bb1140
	case 90:
		goto sw_bb1177
	case 91:
		goto sw_bb1206
	case 92:
		goto sw_bb1227
	case 93:
		goto sw_bb1238
	case 94:
		goto sw_bb1252
	case 95:
		goto sw_bb1256
	case 96:
		goto sw_bb1282
	case 97:
		goto sw_bb1290
	case 98:
		goto sw_bb1305
	case 99:
		goto sw_bb1334
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
	*state_addr = 66
	goto next_state

if_end:
	*i = 0
	goto for_cond

for_cond:
	v11 = *i
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(30)
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
	*state_addr = 82
	goto next_state

if_end19:
	v20 = *result
	tobool20 = (v20 & 1) != 0
	*retval = tobool20
	goto _return

sw_bb21:
	v21 = *lookahead
	cmp22 = v21 == 45
	if cmp22 {
		goto if_then24
	} else {
		goto if_end25
	}

if_then24:
	*state_addr = 53
	goto next_state

if_end25:
	v22 = *lookahead
	cmp26 = v22 == 46
	if cmp26 {
		goto if_then28
	} else {
		goto if_end29
	}

if_then28:
	*state_addr = 91
	goto next_state

if_end29:
	v23 = *lookahead
	cmp30 = v23 == 58
	if cmp30 {
		goto if_then32
	} else {
		goto if_end33
	}

if_then32:
	*state_addr = 50
	goto next_state

if_end33:
	v24 = *lookahead
	cmp34 = 48 <= v24
	if cmp34 {
		goto land_lhs_true36
	} else {
		goto lor_lhs_false
	}

land_lhs_true36:
	v25 = *lookahead
	cmp37 = v25 <= 57
	if cmp37 {
		goto if_then41
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v26 = *lookahead
	cmp39 = v26 == 95
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*state_addr = 6
	goto next_state

if_end42:
	v27 = *result
	tobool43 = (v27 & 1) != 0
	*retval = tobool43
	goto _return

sw_bb44:
	v28 = *lookahead
	cmp45 = v28 == 45
	if cmp45 {
		goto if_then47
	} else {
		goto if_end48
	}

if_then47:
	*state_addr = 54
	goto next_state

if_end48:
	v29 = *lookahead
	cmp49 = 48 <= v29
	if cmp49 {
		goto land_lhs_true51
	} else {
		goto if_end55
	}

land_lhs_true51:
	v30 = *lookahead
	cmp52 = v30 <= 57
	if cmp52 {
		goto if_then54
	} else {
		goto if_end55
	}

if_then54:
	*state_addr = 3
	goto next_state

if_end55:
	v31 = *result
	tobool56 = (v31 & 1) != 0
	*retval = tobool56
	goto _return

sw_bb57:
	v32 = *lookahead
	cmp58 = v32 == 45
	if cmp58 {
		goto if_then60
	} else {
		goto if_end61
	}

if_then60:
	*state_addr = 60
	goto next_state

if_end61:
	v33 = *result
	tobool62 = (v33 & 1) != 0
	*retval = tobool62
	goto _return

sw_bb63:
	v34 = *lookahead
	cmp64 = v34 == 46
	if cmp64 {
		goto if_then66
	} else {
		goto if_end67
	}

if_then66:
	*state_addr = 91
	goto next_state

if_end67:
	v35 = *lookahead
	cmp68 = v35 == 58
	if cmp68 {
		goto if_then70
	} else {
		goto if_end71
	}

if_then70:
	*state_addr = 50
	goto next_state

if_end71:
	v36 = *lookahead
	cmp72 = v36 == 95
	if cmp72 {
		goto if_then74
	} else {
		goto if_end75
	}

if_then74:
	*state_addr = 6
	goto next_state

if_end75:
	v37 = *lookahead
	cmp76 = 48 <= v37
	if cmp76 {
		goto land_lhs_true78
	} else {
		goto if_end82
	}

land_lhs_true78:
	v38 = *lookahead
	cmp79 = v38 <= 57
	if cmp79 {
		goto if_then81
	} else {
		goto if_end82
	}

if_then81:
	*state_addr = 1
	goto next_state

if_end82:
	v39 = *result
	tobool83 = (v39 & 1) != 0
	*retval = tobool83
	goto _return

sw_bb84:
	v40 = *lookahead
	cmp85 = v40 == 46
	if cmp85 {
		goto if_then87
	} else {
		goto if_end88
	}

if_then87:
	*state_addr = 91
	goto next_state

if_end88:
	v41 = *lookahead
	cmp89 = v41 == 58
	if cmp89 {
		goto if_then91
	} else {
		goto if_end92
	}

if_then91:
	*state_addr = 50
	goto next_state

if_end92:
	v42 = *lookahead
	cmp93 = v42 == 95
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*state_addr = 6
	goto next_state

if_end96:
	v43 = *lookahead
	cmp97 = 48 <= v43
	if cmp97 {
		goto land_lhs_true99
	} else {
		goto if_end103
	}

land_lhs_true99:
	v44 = *lookahead
	cmp100 = v44 <= 57
	if cmp100 {
		goto if_then102
	} else {
		goto if_end103
	}

if_then102:
	*state_addr = 4
	goto next_state

if_end103:
	v45 = *result
	tobool104 = (v45 & 1) != 0
	*retval = tobool104
	goto _return

sw_bb105:
	v46 = *lookahead
	cmp106 = v46 == 46
	if cmp106 {
		goto if_then108
	} else {
		goto if_end109
	}

if_then108:
	*state_addr = 91
	goto next_state

if_end109:
	v47 = *lookahead
	cmp110 = v47 == 58
	if cmp110 {
		goto if_then112
	} else {
		goto if_end113
	}

if_then112:
	*state_addr = 50
	goto next_state

if_end113:
	v48 = *lookahead
	cmp114 = 48 <= v48
	if cmp114 {
		goto land_lhs_true116
	} else {
		goto lor_lhs_false119
	}

land_lhs_true116:
	v49 = *lookahead
	cmp117 = v49 <= 57
	if cmp117 {
		goto if_then122
	} else {
		goto lor_lhs_false119
	}

lor_lhs_false119:
	v50 = *lookahead
	cmp120 = v50 == 95
	if cmp120 {
		goto if_then122
	} else {
		goto if_end123
	}

if_then122:
	*state_addr = 6
	goto next_state

if_end123:
	v51 = *result
	tobool124 = (v51 & 1) != 0
	*retval = tobool124
	goto _return

sw_bb125:
	v52 = *lookahead
	cmp126 = v52 == 46
	if cmp126 {
		goto if_then128
	} else {
		goto if_end129
	}

if_then128:
	*state_addr = 90
	goto next_state

if_end129:
	v53 = *lookahead
	cmp130 = v53 == 48
	if cmp130 {
		goto if_then132
	} else {
		goto if_end133
	}

if_then132:
	*state_addr = 78
	goto next_state

if_end133:
	v54 = *lookahead
	cmp134 = 49 <= v54
	if cmp134 {
		goto land_lhs_true136
	} else {
		goto if_end140
	}

land_lhs_true136:
	v55 = *lookahead
	cmp137 = v55 <= 57
	if cmp137 {
		goto if_then139
	} else {
		goto if_end140
	}

if_then139:
	*state_addr = 83
	goto next_state

if_end140:
	v56 = *result
	tobool141 = (v56 & 1) != 0
	*retval = tobool141
	goto _return

sw_bb142:
	v57 = *lookahead
	cmp143 = v57 == 46
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*state_addr = 93
	goto next_state

if_end146:
	v58 = *lookahead
	cmp147 = v58 == 58
	if cmp147 {
		goto if_then149
	} else {
		goto if_end150
	}

if_then149:
	*state_addr = 50
	goto next_state

if_end150:
	v59 = *result
	tobool151 = (v59 & 1) != 0
	*retval = tobool151
	goto _return

sw_bb152:
	v60 = *lookahead
	cmp153 = v60 == 46
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*state_addr = 93
	goto next_state

if_end156:
	v61 = *lookahead
	cmp157 = v61 == 58
	if cmp157 {
		goto if_then159
	} else {
		goto if_end160
	}

if_then159:
	*state_addr = 50
	goto next_state

if_end160:
	v62 = *lookahead
	cmp161 = 48 <= v62
	if cmp161 {
		goto land_lhs_true163
	} else {
		goto if_end167
	}

land_lhs_true163:
	v63 = *lookahead
	cmp164 = v63 <= 57
	if cmp164 {
		goto if_then166
	} else {
		goto if_end167
	}

if_then166:
	*state_addr = 8
	goto next_state

if_end167:
	v64 = *result
	tobool168 = (v64 & 1) != 0
	*retval = tobool168
	goto _return

sw_bb169:
	v65 = *lookahead
	cmp170 = v65 == 58
	if cmp170 {
		goto if_then172
	} else {
		goto if_end173
	}

if_then172:
	*state_addr = 61
	goto next_state

if_end173:
	v66 = *result
	tobool174 = (v66 & 1) != 0
	*retval = tobool174
	goto _return

sw_bb175:
	v67 = *lookahead
	cmp176 = v67 == 58
	if cmp176 {
		goto if_then178
	} else {
		goto if_end179
	}

if_then178:
	*state_addr = 61
	goto next_state

if_end179:
	v68 = *lookahead
	cmp180 = 48 <= v68
	if cmp180 {
		goto land_lhs_true182
	} else {
		goto if_end186
	}

land_lhs_true182:
	v69 = *lookahead
	cmp183 = v69 <= 57
	if cmp183 {
		goto if_then185
	} else {
		goto if_end186
	}

if_then185:
	*state_addr = 10
	goto next_state

if_end186:
	v70 = *result
	tobool187 = (v70 & 1) != 0
	*retval = tobool187
	goto _return

sw_bb188:
	v71 = *lookahead
	cmp189 = v71 == 58
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*state_addr = 62
	goto next_state

if_end192:
	v72 = *result
	tobool193 = (v72 & 1) != 0
	*retval = tobool193
	goto _return

sw_bb194:
	v73 = *lookahead
	cmp195 = v73 == 65
	if cmp195 {
		goto if_then197
	} else {
		goto if_end198
	}

if_then197:
	*state_addr = 20
	goto next_state

if_end198:
	v74 = *lookahead
	cmp199 = v74 == 97
	if cmp199 {
		goto if_then201
	} else {
		goto if_end202
	}

if_then201:
	*state_addr = 36
	goto next_state

if_end202:
	v75 = *result
	tobool203 = (v75 & 1) != 0
	*retval = tobool203
	goto _return

sw_bb204:
	v76 = *lookahead
	cmp205 = v76 == 65
	if cmp205 {
		goto if_then207
	} else {
		goto if_end208
	}

if_then207:
	*state_addr = 22
	goto next_state

if_end208:
	v77 = *lookahead
	cmp209 = v77 == 97
	if cmp209 {
		goto if_then211
	} else {
		goto if_end212
	}

if_then211:
	*state_addr = 22
	goto next_state

if_end212:
	v78 = *result
	tobool213 = (v78 & 1) != 0
	*retval = tobool213
	goto _return

sw_bb214:
	v79 = *lookahead
	cmp215 = v79 == 69
	if cmp215 {
		goto if_then217
	} else {
		goto if_end218
	}

if_then217:
	*state_addr = 68
	goto next_state

if_end218:
	v80 = *result
	tobool219 = (v80 & 1) != 0
	*retval = tobool219
	goto _return

sw_bb220:
	v81 = *lookahead
	cmp221 = v81 == 70
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*state_addr = 68
	goto next_state

if_end224:
	v82 = *result
	tobool225 = (v82 & 1) != 0
	*retval = tobool225
	goto _return

sw_bb226:
	v83 = *lookahead
	cmp227 = v83 == 70
	if cmp227 {
		goto if_then229
	} else {
		goto if_end230
	}

if_then229:
	*state_addr = 16
	goto next_state

if_end230:
	v84 = *lookahead
	cmp231 = v84 == 102
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*state_addr = 32
	goto next_state

if_end234:
	v85 = *lookahead
	cmp235 = v85 == 78
	if cmp235 {
		goto if_then240
	} else {
		goto lor_lhs_false237
	}

lor_lhs_false237:
	v86 = *lookahead
	cmp238 = v86 == 110
	if cmp238 {
		goto if_then240
	} else {
		goto if_end241
	}

if_then240:
	*state_addr = 68
	goto next_state

if_end241:
	v87 = *result
	tobool242 = (v87 & 1) != 0
	*retval = tobool242
	goto _return

sw_bb243:
	v88 = *lookahead
	cmp244 = v88 == 70
	if cmp244 {
		goto if_then246
	} else {
		goto if_end247
	}

if_then246:
	*state_addr = 88
	goto next_state

if_end247:
	v89 = *result
	tobool248 = (v89 & 1) != 0
	*retval = tobool248
	goto _return

sw_bb249:
	v90 = *lookahead
	cmp250 = v90 == 76
	if cmp250 {
		goto if_then252
	} else {
		goto if_end253
	}

if_then252:
	*state_addr = 67
	goto next_state

if_end253:
	v91 = *result
	tobool254 = (v91 & 1) != 0
	*retval = tobool254
	goto _return

sw_bb255:
	v92 = *lookahead
	cmp256 = v92 == 76
	if cmp256 {
		goto if_then258
	} else {
		goto if_end259
	}

if_then258:
	*state_addr = 26
	goto next_state

if_end259:
	v93 = *result
	tobool260 = (v93 & 1) != 0
	*retval = tobool260
	goto _return

sw_bb261:
	v94 = *lookahead
	cmp262 = v94 == 76
	if cmp262 {
		goto if_then264
	} else {
		goto if_end265
	}

if_then264:
	*state_addr = 19
	goto next_state

if_end265:
	v95 = *result
	tobool266 = (v95 & 1) != 0
	*retval = tobool266
	goto _return

sw_bb267:
	v96 = *lookahead
	cmp268 = v96 == 78
	if cmp268 {
		goto if_then270
	} else {
		goto if_end271
	}

if_then270:
	*state_addr = 88
	goto next_state

if_end271:
	v97 = *result
	tobool272 = (v97 & 1) != 0
	*retval = tobool272
	goto _return

sw_bb273:
	v98 = *lookahead
	cmp274 = v98 == 78
	if cmp274 {
		goto if_then276
	} else {
		goto if_end277
	}

if_then276:
	*state_addr = 18
	goto next_state

if_end277:
	v99 = *lookahead
	cmp278 = v99 == 110
	if cmp278 {
		goto if_then280
	} else {
		goto if_end281
	}

if_then280:
	*state_addr = 34
	goto next_state

if_end281:
	v100 = *result
	tobool282 = (v100 & 1) != 0
	*retval = tobool282
	goto _return

sw_bb283:
	v101 = *lookahead
	cmp284 = v101 == 82
	if cmp284 {
		goto if_then286
	} else {
		goto if_end287
	}

if_then286:
	*state_addr = 27
	goto next_state

if_end287:
	v102 = *lookahead
	cmp288 = v102 == 114
	if cmp288 {
		goto if_then290
	} else {
		goto if_end291
	}

if_then290:
	*state_addr = 43
	goto next_state

if_end291:
	v103 = *result
	tobool292 = (v103 & 1) != 0
	*retval = tobool292
	goto _return

sw_bb293:
	v104 = *lookahead
	cmp294 = v104 == 83
	if cmp294 {
		goto if_then296
	} else {
		goto if_end297
	}

if_then296:
	*state_addr = 68
	goto next_state

if_end297:
	v105 = *result
	tobool298 = (v105 & 1) != 0
	*retval = tobool298
	goto _return

sw_bb299:
	v106 = *lookahead
	cmp300 = v106 == 83
	if cmp300 {
		goto if_then302
	} else {
		goto if_end303
	}

if_then302:
	*state_addr = 15
	goto next_state

if_end303:
	v107 = *result
	tobool304 = (v107 & 1) != 0
	*retval = tobool304
	goto _return

sw_bb305:
	v108 = *lookahead
	cmp306 = v108 == 85
	if cmp306 {
		goto if_then308
	} else {
		goto if_end309
	}

if_then308:
	*state_addr = 15
	goto next_state

if_end309:
	v109 = *result
	tobool310 = (v109 & 1) != 0
	*retval = tobool310
	goto _return

sw_bb311:
	v110 = *lookahead
	cmp312 = v110 == 90
	if cmp312 {
		goto if_then314
	} else {
		goto if_end315
	}

if_then314:
	*state_addr = 94
	goto next_state

if_end315:
	v111 = *lookahead
	cmp316 = v111 == 9
	if cmp316 {
		goto if_then321
	} else {
		goto lor_lhs_false318
	}

lor_lhs_false318:
	v112 = *lookahead
	cmp319 = v112 == 32
	if cmp319 {
		goto if_then321
	} else {
		goto if_end322
	}

if_then321:
	*state_addr = 28
	goto next_state

if_end322:
	v113 = *result
	tobool323 = (v113 & 1) != 0
	*retval = tobool323
	goto _return

sw_bb324:
	v114 = *lookahead
	cmp325 = v114 == 97
	if cmp325 {
		goto if_then327
	} else {
		goto if_end328
	}

if_then327:
	*state_addr = 36
	goto next_state

if_end328:
	v115 = *result
	tobool329 = (v115 & 1) != 0
	*retval = tobool329
	goto _return

sw_bb330:
	v116 = *lookahead
	cmp331 = v116 == 97
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*state_addr = 38
	goto next_state

if_end334:
	v117 = *result
	tobool335 = (v117 & 1) != 0
	*retval = tobool335
	goto _return

sw_bb336:
	v118 = *lookahead
	cmp337 = v118 == 101
	if cmp337 {
		goto if_then339
	} else {
		goto if_end340
	}

if_then339:
	*state_addr = 68
	goto next_state

if_end340:
	v119 = *result
	tobool341 = (v119 & 1) != 0
	*retval = tobool341
	goto _return

sw_bb342:
	v120 = *lookahead
	cmp343 = v120 == 102
	if cmp343 {
		goto if_then345
	} else {
		goto if_end346
	}

if_then345:
	*state_addr = 68
	goto next_state

if_end346:
	v121 = *result
	tobool347 = (v121 & 1) != 0
	*retval = tobool347
	goto _return

sw_bb348:
	v122 = *lookahead
	cmp349 = v122 == 102
	if cmp349 {
		goto if_then351
	} else {
		goto if_end352
	}

if_then351:
	*state_addr = 32
	goto next_state

if_end352:
	v123 = *lookahead
	cmp353 = v123 == 110
	if cmp353 {
		goto if_then355
	} else {
		goto if_end356
	}

if_then355:
	*state_addr = 68
	goto next_state

if_end356:
	v124 = *result
	tobool357 = (v124 & 1) != 0
	*retval = tobool357
	goto _return

sw_bb358:
	v125 = *lookahead
	cmp359 = v125 == 102
	if cmp359 {
		goto if_then361
	} else {
		goto if_end362
	}

if_then361:
	*state_addr = 88
	goto next_state

if_end362:
	v126 = *result
	tobool363 = (v126 & 1) != 0
	*retval = tobool363
	goto _return

sw_bb364:
	v127 = *lookahead
	cmp365 = v127 == 108
	if cmp365 {
		goto if_then367
	} else {
		goto if_end368
	}

if_then367:
	*state_addr = 67
	goto next_state

if_end368:
	v128 = *result
	tobool369 = (v128 & 1) != 0
	*retval = tobool369
	goto _return

sw_bb370:
	v129 = *lookahead
	cmp371 = v129 == 108
	if cmp371 {
		goto if_then373
	} else {
		goto if_end374
	}

if_then373:
	*state_addr = 42
	goto next_state

if_end374:
	v130 = *result
	tobool375 = (v130 & 1) != 0
	*retval = tobool375
	goto _return

sw_bb376:
	v131 = *lookahead
	cmp377 = v131 == 108
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*state_addr = 35
	goto next_state

if_end380:
	v132 = *result
	tobool381 = (v132 & 1) != 0
	*retval = tobool381
	goto _return

sw_bb382:
	v133 = *lookahead
	cmp383 = v133 == 110
	if cmp383 {
		goto if_then385
	} else {
		goto if_end386
	}

if_then385:
	*state_addr = 88
	goto next_state

if_end386:
	v134 = *result
	tobool387 = (v134 & 1) != 0
	*retval = tobool387
	goto _return

sw_bb388:
	v135 = *lookahead
	cmp389 = v135 == 110
	if cmp389 {
		goto if_then391
	} else {
		goto if_end392
	}

if_then391:
	*state_addr = 34
	goto next_state

if_end392:
	v136 = *result
	tobool393 = (v136 & 1) != 0
	*retval = tobool393
	goto _return

sw_bb394:
	v137 = *lookahead
	cmp395 = v137 == 114
	if cmp395 {
		goto if_then397
	} else {
		goto if_end398
	}

if_then397:
	*state_addr = 43
	goto next_state

if_end398:
	v138 = *result
	tobool399 = (v138 & 1) != 0
	*retval = tobool399
	goto _return

sw_bb400:
	v139 = *lookahead
	cmp401 = v139 == 115
	if cmp401 {
		goto if_then403
	} else {
		goto if_end404
	}

if_then403:
	*state_addr = 68
	goto next_state

if_end404:
	v140 = *result
	tobool405 = (v140 & 1) != 0
	*retval = tobool405
	goto _return

sw_bb406:
	v141 = *lookahead
	cmp407 = v141 == 115
	if cmp407 {
		goto if_then409
	} else {
		goto if_end410
	}

if_then409:
	*state_addr = 31
	goto next_state

if_end410:
	v142 = *result
	tobool411 = (v142 & 1) != 0
	*retval = tobool411
	goto _return

sw_bb412:
	v143 = *lookahead
	cmp413 = v143 == 117
	if cmp413 {
		goto if_then415
	} else {
		goto if_end416
	}

if_then415:
	*state_addr = 31
	goto next_state

if_end416:
	v144 = *result
	tobool417 = (v144 & 1) != 0
	*retval = tobool417
	goto _return

sw_bb418:
	v145 = *lookahead
	cmp419 = v145 == 9
	if cmp419 {
		goto if_then424
	} else {
		goto lor_lhs_false421
	}

lor_lhs_false421:
	v146 = *lookahead
	cmp422 = v146 == 32
	if cmp422 {
		goto if_then424
	} else {
		goto if_end425
	}

if_then424:
	*state_addr = 47
	goto next_state

if_end425:
	v147 = *lookahead
	cmp426 = v147 == 84
	if cmp426 {
		goto if_then431
	} else {
		goto lor_lhs_false428
	}

lor_lhs_false428:
	v148 = *lookahead
	cmp429 = v148 == 116
	if cmp429 {
		goto if_then431
	} else {
		goto if_end432
	}

if_then431:
	*state_addr = 55
	goto next_state

if_end432:
	v149 = *result
	tobool433 = (v149 & 1) != 0
	*retval = tobool433
	goto _return

sw_bb434:
	v150 = *lookahead
	cmp435 = v150 == 9
	if cmp435 {
		goto if_then440
	} else {
		goto lor_lhs_false437
	}

lor_lhs_false437:
	v151 = *lookahead
	cmp438 = v151 == 32
	if cmp438 {
		goto if_then440
	} else {
		goto if_end441
	}

if_then440:
	*state_addr = 47
	goto next_state

if_end441:
	v152 = *lookahead
	cmp442 = v152 == 84
	if cmp442 {
		goto if_then447
	} else {
		goto lor_lhs_false444
	}

lor_lhs_false444:
	v153 = *lookahead
	cmp445 = v153 == 116
	if cmp445 {
		goto if_then447
	} else {
		goto if_end448
	}

if_then447:
	*state_addr = 55
	goto next_state

if_end448:
	v154 = *lookahead
	cmp449 = 48 <= v154
	if cmp449 {
		goto land_lhs_true451
	} else {
		goto if_end455
	}

land_lhs_true451:
	v155 = *lookahead
	cmp452 = v155 <= 57
	if cmp452 {
		goto if_then454
	} else {
		goto if_end455
	}

if_then454:
	*state_addr = 44
	goto next_state

if_end455:
	v156 = *result
	tobool456 = (v156 & 1) != 0
	*retval = tobool456
	goto _return

sw_bb457:
	v157 = *lookahead
	cmp458 = v157 == 9
	if cmp458 {
		goto if_then463
	} else {
		goto lor_lhs_false460
	}

lor_lhs_false460:
	v158 = *lookahead
	cmp461 = v158 == 32
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*state_addr = 47
	goto next_state

if_end464:
	v159 = *lookahead
	cmp465 = v159 == 84
	if cmp465 {
		goto if_then470
	} else {
		goto lor_lhs_false467
	}

lor_lhs_false467:
	v160 = *lookahead
	cmp468 = v160 == 116
	if cmp468 {
		goto if_then470
	} else {
		goto if_end471
	}

if_then470:
	*state_addr = 55
	goto next_state

if_end471:
	v161 = *lookahead
	cmp472 = 48 <= v161
	if cmp472 {
		goto land_lhs_true474
	} else {
		goto if_end478
	}

land_lhs_true474:
	v162 = *lookahead
	cmp475 = v162 <= 57
	if cmp475 {
		goto if_then477
	} else {
		goto if_end478
	}

if_then477:
	*state_addr = 99
	goto next_state

if_end478:
	v163 = *result
	tobool479 = (v163 & 1) != 0
	*retval = tobool479
	goto _return

sw_bb480:
	v164 = *lookahead
	cmp481 = v164 == 9
	if cmp481 {
		goto if_then486
	} else {
		goto lor_lhs_false483
	}

lor_lhs_false483:
	v165 = *lookahead
	cmp484 = v165 == 32
	if cmp484 {
		goto if_then486
	} else {
		goto if_end487
	}

if_then486:
	*state_addr = 47
	goto next_state

if_end487:
	v166 = *lookahead
	cmp488 = 48 <= v166
	if cmp488 {
		goto land_lhs_true490
	} else {
		goto if_end494
	}

land_lhs_true490:
	v167 = *lookahead
	cmp491 = v167 <= 57
	if cmp491 {
		goto if_then493
	} else {
		goto if_end494
	}

if_then493:
	*state_addr = 11
	goto next_state

if_end494:
	v168 = *result
	tobool495 = (v168 & 1) != 0
	*retval = tobool495
	goto _return

sw_bb496:
	v169 = *lookahead
	cmp497 = v169 == 43
	if cmp497 {
		goto if_then502
	} else {
		goto lor_lhs_false499
	}

lor_lhs_false499:
	v170 = *lookahead
	cmp500 = v170 == 45
	if cmp500 {
		goto if_then502
	} else {
		goto if_end503
	}

if_then502:
	*state_addr = 52
	goto next_state

if_end503:
	v171 = *result
	tobool504 = (v171 & 1) != 0
	*retval = tobool504
	goto _return

sw_bb505:
	v172 = *lookahead
	cmp506 = v172 == 48
	if cmp506 {
		goto if_then514
	} else {
		goto lor_lhs_false508
	}

lor_lhs_false508:
	v173 = *lookahead
	cmp509 = v173 == 49
	if cmp509 {
		goto if_then514
	} else {
		goto lor_lhs_false511
	}

lor_lhs_false511:
	v174 = *lookahead
	cmp512 = v174 == 95
	if cmp512 {
		goto if_then514
	} else {
		goto if_end515
	}

if_then514:
	*state_addr = 86
	goto next_state

if_end515:
	v175 = *result
	tobool516 = (v175 & 1) != 0
	*retval = tobool516
	goto _return

sw_bb517:
	v176 = *lookahead
	cmp518 = 54 <= v176
	if cmp518 {
		goto land_lhs_true520
	} else {
		goto if_end524
	}

land_lhs_true520:
	v177 = *lookahead
	cmp521 = v177 <= 57
	if cmp521 {
		goto if_then523
	} else {
		goto if_end524
	}

if_then523:
	*state_addr = 8
	goto next_state

if_end524:
	v178 = *lookahead
	cmp525 = 48 <= v178
	if cmp525 {
		goto land_lhs_true527
	} else {
		goto if_end531
	}

land_lhs_true527:
	v179 = *lookahead
	cmp528 = v179 <= 53
	if cmp528 {
		goto if_then530
	} else {
		goto if_end531
	}

if_then530:
	*state_addr = 9
	goto next_state

if_end531:
	v180 = *result
	tobool532 = (v180 & 1) != 0
	*retval = tobool532
	goto _return

sw_bb533:
	v181 = *lookahead
	cmp534 = 54 <= v181
	if cmp534 {
		goto land_lhs_true536
	} else {
		goto if_end540
	}

land_lhs_true536:
	v182 = *lookahead
	cmp537 = v182 <= 57
	if cmp537 {
		goto if_then539
	} else {
		goto if_end540
	}

if_then539:
	*state_addr = 84
	goto next_state

if_end540:
	v183 = *lookahead
	cmp541 = 48 <= v183
	if cmp541 {
		goto land_lhs_true543
	} else {
		goto if_end547
	}

land_lhs_true543:
	v184 = *lookahead
	cmp544 = v184 <= 53
	if cmp544 {
		goto if_then546
	} else {
		goto if_end547
	}

if_then546:
	*state_addr = 85
	goto next_state

if_end547:
	v185 = *result
	tobool548 = (v185 & 1) != 0
	*retval = tobool548
	goto _return

sw_bb549:
	v186 = *lookahead
	cmp550 = 48 <= v186
	if cmp550 {
		goto land_lhs_true552
	} else {
		goto if_end556
	}

land_lhs_true552:
	v187 = *lookahead
	cmp553 = v187 <= 57
	if cmp553 {
		goto if_then555
	} else {
		goto if_end556
	}

if_then555:
	*state_addr = 92
	goto next_state

if_end556:
	v188 = *result
	tobool557 = (v188 & 1) != 0
	*retval = tobool557
	goto _return

sw_bb558:
	v189 = *lookahead
	cmp559 = 48 <= v189
	if cmp559 {
		goto land_lhs_true561
	} else {
		goto if_end565
	}

land_lhs_true561:
	v190 = *lookahead
	cmp562 = v190 <= 57
	if cmp562 {
		goto if_then564
	} else {
		goto if_end565
	}

if_then564:
	*state_addr = 2
	goto next_state

if_end565:
	v191 = *result
	tobool566 = (v191 & 1) != 0
	*retval = tobool566
	goto _return

sw_bb567:
	v192 = *lookahead
	cmp568 = 48 <= v192
	if cmp568 {
		goto land_lhs_true570
	} else {
		goto if_end574
	}

land_lhs_true570:
	v193 = *lookahead
	cmp571 = v193 <= 57
	if cmp571 {
		goto if_then573
	} else {
		goto if_end574
	}

if_then573:
	*state_addr = 45
	goto next_state

if_end574:
	v194 = *result
	tobool575 = (v194 & 1) != 0
	*retval = tobool575
	goto _return

sw_bb576:
	v195 = *lookahead
	cmp577 = 48 <= v195
	if cmp577 {
		goto land_lhs_true579
	} else {
		goto if_end583
	}

land_lhs_true579:
	v196 = *lookahead
	cmp580 = v196 <= 57
	if cmp580 {
		goto if_then582
	} else {
		goto if_end583
	}

if_then582:
	*state_addr = 11
	goto next_state

if_end583:
	v197 = *result
	tobool584 = (v197 & 1) != 0
	*retval = tobool584
	goto _return

sw_bb585:
	v198 = *lookahead
	cmp586 = 48 <= v198
	if cmp586 {
		goto land_lhs_true588
	} else {
		goto if_end592
	}

land_lhs_true588:
	v199 = *lookahead
	cmp589 = v199 <= 57
	if cmp589 {
		goto if_then591
	} else {
		goto if_end592
	}

if_then591:
	*state_addr = 95
	goto next_state

if_end592:
	v200 = *result
	tobool593 = (v200 & 1) != 0
	*retval = tobool593
	goto _return

sw_bb594:
	v201 = *lookahead
	cmp595 = 48 <= v201
	if cmp595 {
		goto land_lhs_true597
	} else {
		goto if_end601
	}

land_lhs_true597:
	v202 = *lookahead
	cmp598 = v202 <= 57
	if cmp598 {
		goto if_then600
	} else {
		goto if_end601
	}

if_then600:
	*state_addr = 94
	goto next_state

if_end601:
	v203 = *result
	tobool602 = (v203 & 1) != 0
	*retval = tobool602
	goto _return

sw_bb603:
	v204 = *lookahead
	cmp604 = 48 <= v204
	if cmp604 {
		goto land_lhs_true606
	} else {
		goto if_end610
	}

land_lhs_true606:
	v205 = *lookahead
	cmp607 = v205 <= 57
	if cmp607 {
		goto if_then609
	} else {
		goto if_end610
	}

if_then609:
	*state_addr = 98
	goto next_state

if_end610:
	v206 = *result
	tobool611 = (v206 & 1) != 0
	*retval = tobool611
	goto _return

sw_bb612:
	v207 = *lookahead
	cmp613 = 48 <= v207
	if cmp613 {
		goto land_lhs_true615
	} else {
		goto if_end619
	}

land_lhs_true615:
	v208 = *lookahead
	cmp616 = v208 <= 57
	if cmp616 {
		goto if_then618
	} else {
		goto if_end619
	}

if_then618:
	*state_addr = 97
	goto next_state

if_end619:
	v209 = *result
	tobool620 = (v209 & 1) != 0
	*retval = tobool620
	goto _return

sw_bb621:
	v210 = *lookahead
	cmp622 = 48 <= v210
	if cmp622 {
		goto land_lhs_true624
	} else {
		goto if_end628
	}

land_lhs_true624:
	v211 = *lookahead
	cmp625 = v211 <= 57
	if cmp625 {
		goto if_then627
	} else {
		goto if_end628
	}

if_then627:
	*state_addr = 46
	goto next_state

if_end628:
	v212 = *result
	tobool629 = (v212 & 1) != 0
	*retval = tobool629
	goto _return

sw_bb630:
	v213 = *lookahead
	cmp631 = 48 <= v213
	if cmp631 {
		goto land_lhs_true633
	} else {
		goto if_end637
	}

land_lhs_true633:
	v214 = *lookahead
	cmp634 = v214 <= 57
	if cmp634 {
		goto if_then636
	} else {
		goto if_end637
	}

if_then636:
	*state_addr = 64
	goto next_state

if_end637:
	v215 = *result
	tobool638 = (v215 & 1) != 0
	*retval = tobool638
	goto _return

sw_bb639:
	v216 = *lookahead
	cmp640 = 48 <= v216
	if cmp640 {
		goto land_lhs_true642
	} else {
		goto if_end646
	}

land_lhs_true642:
	v217 = *lookahead
	cmp643 = v217 <= 57
	if cmp643 {
		goto if_then645
	} else {
		goto if_end646
	}

if_then645:
	*state_addr = 56
	goto next_state

if_end646:
	v218 = *result
	tobool647 = (v218 & 1) != 0
	*retval = tobool647
	goto _return

sw_bb648:
	v219 = *lookahead
	cmp649 = 48 <= v219
	if cmp649 {
		goto land_lhs_true651
	} else {
		goto if_end655
	}

land_lhs_true651:
	v220 = *lookahead
	cmp652 = v220 <= 57
	if cmp652 {
		goto if_then654
	} else {
		goto if_end655
	}

if_then654:
	*state_addr = 57
	goto next_state

if_end655:
	v221 = *result
	tobool656 = (v221 & 1) != 0
	*retval = tobool656
	goto _return

sw_bb657:
	v222 = *lookahead
	cmp658 = 48 <= v222
	if cmp658 {
		goto land_lhs_true660
	} else {
		goto if_end664
	}

land_lhs_true660:
	v223 = *lookahead
	cmp661 = v223 <= 57
	if cmp661 {
		goto if_then663
	} else {
		goto if_end664
	}

if_then663:
	*state_addr = 12
	goto next_state

if_end664:
	v224 = *result
	tobool665 = (v224 & 1) != 0
	*retval = tobool665
	goto _return

sw_bb666:
	v225 = *lookahead
	cmp667 = 48 <= v225
	if cmp667 {
		goto land_lhs_true669
	} else {
		goto lor_lhs_false672
	}

land_lhs_true669:
	v226 = *lookahead
	cmp670 = v226 <= 57
	if cmp670 {
		goto if_then687
	} else {
		goto lor_lhs_false672
	}

lor_lhs_false672:
	v227 = *lookahead
	cmp673 = 65 <= v227
	if cmp673 {
		goto land_lhs_true675
	} else {
		goto lor_lhs_false678
	}

land_lhs_true675:
	v228 = *lookahead
	cmp676 = v228 <= 70
	if cmp676 {
		goto if_then687
	} else {
		goto lor_lhs_false678
	}

lor_lhs_false678:
	v229 = *lookahead
	cmp679 = v229 == 95
	if cmp679 {
		goto if_then687
	} else {
		goto lor_lhs_false681
	}

lor_lhs_false681:
	v230 = *lookahead
	cmp682 = 97 <= v230
	if cmp682 {
		goto land_lhs_true684
	} else {
		goto if_end688
	}

land_lhs_true684:
	v231 = *lookahead
	cmp685 = v231 <= 102
	if cmp685 {
		goto if_then687
	} else {
		goto if_end688
	}

if_then687:
	*state_addr = 87
	goto next_state

if_end688:
	v232 = *result
	tobool689 = (v232 & 1) != 0
	*retval = tobool689
	goto _return

sw_bb690:
	*result = 1
	v233 = *lexer_addr
	result_symbol = &v233.F1
	*result_symbol = 0
	v234 = *lexer_addr
	mark_end = &v234.F3
	v235 = *mark_end
	v236 = *lexer_addr
	v235(v236)
	v237 = *result
	tobool691 = (v237 & 1) != 0
	*retval = tobool691
	goto _return

sw_bb692:
	*result = 1
	v238 = *lexer_addr
	result_symbol693 = &v238.F1
	*result_symbol693 = 1
	v239 = *lexer_addr
	mark_end694 = &v239.F3
	v240 = *mark_end694
	v241 = *lexer_addr
	v240(v241)
	v242 = *result
	tobool695 = (v242 & 1) != 0
	*retval = tobool695
	goto _return

sw_bb696:
	*result = 1
	v243 = *lexer_addr
	result_symbol697 = &v243.F1
	*result_symbol697 = 2
	v244 = *lexer_addr
	mark_end698 = &v244.F3
	v245 = *mark_end698
	v246 = *lexer_addr
	v245(v246)
	v247 = *result
	tobool699 = (v247 & 1) != 0
	*retval = tobool699
	goto _return

sw_bb700:
	*result = 1
	v248 = *lexer_addr
	result_symbol701 = &v248.F1
	*result_symbol701 = 2
	v249 = *lexer_addr
	mark_end702 = &v249.F3
	v250 = *mark_end702
	v251 = *lexer_addr
	v250(v251)
	v252 = *lookahead
	cmp703 = v252 == 69
	if cmp703 {
		goto if_then705
	} else {
		goto if_end706
	}

if_then705:
	*state_addr = 25
	goto next_state

if_end706:
	v253 = *lookahead
	cmp707 = v253 == 101
	if cmp707 {
		goto if_then709
	} else {
		goto if_end710
	}

if_then709:
	*state_addr = 41
	goto next_state

if_end710:
	v254 = *result
	tobool711 = (v254 & 1) != 0
	*retval = tobool711
	goto _return

sw_bb712:
	*result = 1
	v255 = *lexer_addr
	result_symbol713 = &v255.F1
	*result_symbol713 = 2
	v256 = *lexer_addr
	mark_end714 = &v256.F3
	v257 = *mark_end714
	v258 = *lexer_addr
	v257(v258)
	v259 = *lookahead
	cmp715 = v259 == 85
	if cmp715 {
		goto if_then717
	} else {
		goto if_end718
	}

if_then717:
	*state_addr = 21
	goto next_state

if_end718:
	v260 = *lookahead
	cmp719 = v260 == 117
	if cmp719 {
		goto if_then721
	} else {
		goto if_end722
	}

if_then721:
	*state_addr = 37
	goto next_state

if_end722:
	v261 = *lookahead
	cmp723 = v261 == 79
	if cmp723 {
		goto if_then728
	} else {
		goto lor_lhs_false725
	}

lor_lhs_false725:
	v262 = *lookahead
	cmp726 = v262 == 111
	if cmp726 {
		goto if_then728
	} else {
		goto if_end729
	}

if_then728:
	*state_addr = 68
	goto next_state

if_end729:
	v263 = *result
	tobool730 = (v263 & 1) != 0
	*retval = tobool730
	goto _return

sw_bb731:
	*result = 1
	v264 = *lexer_addr
	result_symbol732 = &v264.F1
	*result_symbol732 = 2
	v265 = *lexer_addr
	mark_end733 = &v265.F3
	v266 = *mark_end733
	v267 = *lexer_addr
	v266(v267)
	v268 = *lookahead
	cmp734 = v268 == 101
	if cmp734 {
		goto if_then736
	} else {
		goto if_end737
	}

if_then736:
	*state_addr = 41
	goto next_state

if_end737:
	v269 = *result
	tobool738 = (v269 & 1) != 0
	*retval = tobool738
	goto _return

sw_bb739:
	*result = 1
	v270 = *lexer_addr
	result_symbol740 = &v270.F1
	*result_symbol740 = 2
	v271 = *lexer_addr
	mark_end741 = &v271.F3
	v272 = *mark_end741
	v273 = *lexer_addr
	v272(v273)
	v274 = *lookahead
	cmp742 = v274 == 111
	if cmp742 {
		goto if_then744
	} else {
		goto if_end745
	}

if_then744:
	*state_addr = 68
	goto next_state

if_end745:
	v275 = *lookahead
	cmp746 = v275 == 117
	if cmp746 {
		goto if_then748
	} else {
		goto if_end749
	}

if_then748:
	*state_addr = 37
	goto next_state

if_end749:
	v276 = *result
	tobool750 = (v276 & 1) != 0
	*retval = tobool750
	goto _return

sw_bb751:
	*result = 1
	v277 = *lexer_addr
	result_symbol752 = &v277.F1
	*result_symbol752 = 3
	v278 = *lexer_addr
	mark_end753 = &v278.F3
	v279 = *mark_end753
	v280 = *lexer_addr
	v279(v280)
	v281 = *lookahead
	cmp754 = v281 == 45
	if cmp754 {
		goto if_then756
	} else {
		goto if_end757
	}

if_then756:
	*state_addr = 53
	goto next_state

if_end757:
	v282 = *lookahead
	cmp758 = v282 == 46
	if cmp758 {
		goto if_then760
	} else {
		goto if_end761
	}

if_then760:
	*state_addr = 91
	goto next_state

if_end761:
	v283 = *lookahead
	cmp762 = v283 == 58
	if cmp762 {
		goto if_then764
	} else {
		goto if_end765
	}

if_then764:
	*state_addr = 50
	goto next_state

if_end765:
	v284 = *lookahead
	cmp766 = v284 == 56
	if cmp766 {
		goto if_then771
	} else {
		goto lor_lhs_false768
	}

lor_lhs_false768:
	v285 = *lookahead
	cmp769 = v285 == 57
	if cmp769 {
		goto if_then771
	} else {
		goto if_end772
	}

if_then771:
	*state_addr = 6
	goto next_state

if_end772:
	v286 = *lookahead
	cmp773 = 48 <= v286
	if cmp773 {
		goto land_lhs_true775
	} else {
		goto lor_lhs_false778
	}

land_lhs_true775:
	v287 = *lookahead
	cmp776 = v287 <= 55
	if cmp776 {
		goto if_then781
	} else {
		goto lor_lhs_false778
	}

lor_lhs_false778:
	v288 = *lookahead
	cmp779 = v288 == 95
	if cmp779 {
		goto if_then781
	} else {
		goto if_end782
	}

if_then781:
	*state_addr = 79
	goto next_state

if_end782:
	v289 = *result
	tobool783 = (v289 & 1) != 0
	*retval = tobool783
	goto _return

sw_bb784:
	*result = 1
	v290 = *lexer_addr
	result_symbol785 = &v290.F1
	*result_symbol785 = 3
	v291 = *lexer_addr
	mark_end786 = &v291.F3
	v292 = *mark_end786
	v293 = *lexer_addr
	v292(v293)
	v294 = *lookahead
	cmp787 = v294 == 45
	if cmp787 {
		goto if_then789
	} else {
		goto if_end790
	}

if_then789:
	*state_addr = 53
	goto next_state

if_end790:
	v295 = *lookahead
	cmp791 = v295 == 46
	if cmp791 {
		goto if_then793
	} else {
		goto if_end794
	}

if_then793:
	*state_addr = 91
	goto next_state

if_end794:
	v296 = *lookahead
	cmp795 = v296 == 58
	if cmp795 {
		goto if_then797
	} else {
		goto if_end798
	}

if_then797:
	*state_addr = 51
	goto next_state

if_end798:
	v297 = *lookahead
	cmp799 = 48 <= v297
	if cmp799 {
		goto land_lhs_true801
	} else {
		goto lor_lhs_false804
	}

land_lhs_true801:
	v298 = *lookahead
	cmp802 = v298 <= 57
	if cmp802 {
		goto if_then807
	} else {
		goto lor_lhs_false804
	}

lor_lhs_false804:
	v299 = *lookahead
	cmp805 = v299 == 95
	if cmp805 {
		goto if_then807
	} else {
		goto if_end808
	}

if_then807:
	*state_addr = 83
	goto next_state

if_end808:
	v300 = *result
	tobool809 = (v300 & 1) != 0
	*retval = tobool809
	goto _return

sw_bb810:
	*result = 1
	v301 = *lexer_addr
	result_symbol811 = &v301.F1
	*result_symbol811 = 3
	v302 = *lexer_addr
	mark_end812 = &v302.F3
	v303 = *mark_end812
	v304 = *lexer_addr
	v303(v304)
	v305 = *lookahead
	cmp813 = v305 == 46
	if cmp813 {
		goto if_then815
	} else {
		goto if_end816
	}

if_then815:
	*state_addr = 91
	goto next_state

if_end816:
	v306 = *lookahead
	cmp817 = v306 == 58
	if cmp817 {
		goto if_then819
	} else {
		goto if_end820
	}

if_then819:
	*state_addr = 50
	goto next_state

if_end820:
	v307 = *lookahead
	cmp821 = v307 == 95
	if cmp821 {
		goto if_then823
	} else {
		goto if_end824
	}

if_then823:
	*state_addr = 79
	goto next_state

if_end824:
	v308 = *lookahead
	cmp825 = v308 == 98
	if cmp825 {
		goto if_then827
	} else {
		goto if_end828
	}

if_then827:
	*state_addr = 49
	goto next_state

if_end828:
	v309 = *lookahead
	cmp829 = v309 == 120
	if cmp829 {
		goto if_then831
	} else {
		goto if_end832
	}

if_then831:
	*state_addr = 65
	goto next_state

if_end832:
	v310 = *lookahead
	cmp833 = v310 == 56
	if cmp833 {
		goto if_then838
	} else {
		goto lor_lhs_false835
	}

lor_lhs_false835:
	v311 = *lookahead
	cmp836 = v311 == 57
	if cmp836 {
		goto if_then838
	} else {
		goto if_end839
	}

if_then838:
	*state_addr = 5
	goto next_state

if_end839:
	v312 = *lookahead
	cmp840 = 48 <= v312
	if cmp840 {
		goto land_lhs_true842
	} else {
		goto if_end846
	}

land_lhs_true842:
	v313 = *lookahead
	cmp843 = v313 <= 55
	if cmp843 {
		goto if_then845
	} else {
		goto if_end846
	}

if_then845:
	*state_addr = 77
	goto next_state

if_end846:
	v314 = *result
	tobool847 = (v314 & 1) != 0
	*retval = tobool847
	goto _return

sw_bb848:
	*result = 1
	v315 = *lexer_addr
	result_symbol849 = &v315.F1
	*result_symbol849 = 3
	v316 = *lexer_addr
	mark_end850 = &v316.F3
	v317 = *mark_end850
	v318 = *lexer_addr
	v317(v318)
	v319 = *lookahead
	cmp851 = v319 == 46
	if cmp851 {
		goto if_then853
	} else {
		goto if_end854
	}

if_then853:
	*state_addr = 91
	goto next_state

if_end854:
	v320 = *lookahead
	cmp855 = v320 == 58
	if cmp855 {
		goto if_then857
	} else {
		goto if_end858
	}

if_then857:
	*state_addr = 50
	goto next_state

if_end858:
	v321 = *lookahead
	cmp859 = v321 == 95
	if cmp859 {
		goto if_then861
	} else {
		goto if_end862
	}

if_then861:
	*state_addr = 79
	goto next_state

if_end862:
	v322 = *lookahead
	cmp863 = v322 == 56
	if cmp863 {
		goto if_then868
	} else {
		goto lor_lhs_false865
	}

lor_lhs_false865:
	v323 = *lookahead
	cmp866 = v323 == 57
	if cmp866 {
		goto if_then868
	} else {
		goto if_end869
	}

if_then868:
	*state_addr = 1
	goto next_state

if_end869:
	v324 = *lookahead
	cmp870 = 48 <= v324
	if cmp870 {
		goto land_lhs_true872
	} else {
		goto if_end876
	}

land_lhs_true872:
	v325 = *lookahead
	cmp873 = v325 <= 55
	if cmp873 {
		goto if_then875
	} else {
		goto if_end876
	}

if_then875:
	*state_addr = 73
	goto next_state

if_end876:
	v326 = *result
	tobool877 = (v326 & 1) != 0
	*retval = tobool877
	goto _return

sw_bb878:
	*result = 1
	v327 = *lexer_addr
	result_symbol879 = &v327.F1
	*result_symbol879 = 3
	v328 = *lexer_addr
	mark_end880 = &v328.F3
	v329 = *mark_end880
	v330 = *lexer_addr
	v329(v330)
	v331 = *lookahead
	cmp881 = v331 == 46
	if cmp881 {
		goto if_then883
	} else {
		goto if_end884
	}

if_then883:
	*state_addr = 91
	goto next_state

if_end884:
	v332 = *lookahead
	cmp885 = v332 == 58
	if cmp885 {
		goto if_then887
	} else {
		goto if_end888
	}

if_then887:
	*state_addr = 50
	goto next_state

if_end888:
	v333 = *lookahead
	cmp889 = v333 == 95
	if cmp889 {
		goto if_then891
	} else {
		goto if_end892
	}

if_then891:
	*state_addr = 79
	goto next_state

if_end892:
	v334 = *lookahead
	cmp893 = v334 == 56
	if cmp893 {
		goto if_then898
	} else {
		goto lor_lhs_false895
	}

lor_lhs_false895:
	v335 = *lookahead
	cmp896 = v335 == 57
	if cmp896 {
		goto if_then898
	} else {
		goto if_end899
	}

if_then898:
	*state_addr = 4
	goto next_state

if_end899:
	v336 = *lookahead
	cmp900 = 48 <= v336
	if cmp900 {
		goto land_lhs_true902
	} else {
		goto if_end906
	}

land_lhs_true902:
	v337 = *lookahead
	cmp903 = v337 <= 55
	if cmp903 {
		goto if_then905
	} else {
		goto if_end906
	}

if_then905:
	*state_addr = 76
	goto next_state

if_end906:
	v338 = *result
	tobool907 = (v338 & 1) != 0
	*retval = tobool907
	goto _return

sw_bb908:
	*result = 1
	v339 = *lexer_addr
	result_symbol909 = &v339.F1
	*result_symbol909 = 3
	v340 = *lexer_addr
	mark_end910 = &v340.F3
	v341 = *mark_end910
	v342 = *lexer_addr
	v341(v342)
	v343 = *lookahead
	cmp911 = v343 == 46
	if cmp911 {
		goto if_then913
	} else {
		goto if_end914
	}

if_then913:
	*state_addr = 91
	goto next_state

if_end914:
	v344 = *lookahead
	cmp915 = v344 == 58
	if cmp915 {
		goto if_then917
	} else {
		goto if_end918
	}

if_then917:
	*state_addr = 50
	goto next_state

if_end918:
	v345 = *lookahead
	cmp919 = v345 == 98
	if cmp919 {
		goto if_then921
	} else {
		goto if_end922
	}

if_then921:
	*state_addr = 49
	goto next_state

if_end922:
	v346 = *lookahead
	cmp923 = v346 == 120
	if cmp923 {
		goto if_then925
	} else {
		goto if_end926
	}

if_then925:
	*state_addr = 65
	goto next_state

if_end926:
	v347 = *lookahead
	cmp927 = v347 == 56
	if cmp927 {
		goto if_then932
	} else {
		goto lor_lhs_false929
	}

lor_lhs_false929:
	v348 = *lookahead
	cmp930 = v348 == 57
	if cmp930 {
		goto if_then932
	} else {
		goto if_end933
	}

if_then932:
	*state_addr = 6
	goto next_state

if_end933:
	v349 = *lookahead
	cmp934 = 48 <= v349
	if cmp934 {
		goto land_lhs_true936
	} else {
		goto lor_lhs_false939
	}

land_lhs_true936:
	v350 = *lookahead
	cmp937 = v350 <= 55
	if cmp937 {
		goto if_then942
	} else {
		goto lor_lhs_false939
	}

lor_lhs_false939:
	v351 = *lookahead
	cmp940 = v351 == 95
	if cmp940 {
		goto if_then942
	} else {
		goto if_end943
	}

if_then942:
	*state_addr = 79
	goto next_state

if_end943:
	v352 = *result
	tobool944 = (v352 & 1) != 0
	*retval = tobool944
	goto _return

sw_bb945:
	*result = 1
	v353 = *lexer_addr
	result_symbol946 = &v353.F1
	*result_symbol946 = 3
	v354 = *lexer_addr
	mark_end947 = &v354.F3
	v355 = *mark_end947
	v356 = *lexer_addr
	v355(v356)
	v357 = *lookahead
	cmp948 = v357 == 46
	if cmp948 {
		goto if_then950
	} else {
		goto if_end951
	}

if_then950:
	*state_addr = 91
	goto next_state

if_end951:
	v358 = *lookahead
	cmp952 = v358 == 58
	if cmp952 {
		goto if_then954
	} else {
		goto if_end955
	}

if_then954:
	*state_addr = 50
	goto next_state

if_end955:
	v359 = *lookahead
	cmp956 = v359 == 56
	if cmp956 {
		goto if_then961
	} else {
		goto lor_lhs_false958
	}

lor_lhs_false958:
	v360 = *lookahead
	cmp959 = v360 == 57
	if cmp959 {
		goto if_then961
	} else {
		goto if_end962
	}

if_then961:
	*state_addr = 6
	goto next_state

if_end962:
	v361 = *lookahead
	cmp963 = 48 <= v361
	if cmp963 {
		goto land_lhs_true965
	} else {
		goto lor_lhs_false968
	}

land_lhs_true965:
	v362 = *lookahead
	cmp966 = v362 <= 55
	if cmp966 {
		goto if_then971
	} else {
		goto lor_lhs_false968
	}

lor_lhs_false968:
	v363 = *lookahead
	cmp969 = v363 == 95
	if cmp969 {
		goto if_then971
	} else {
		goto if_end972
	}

if_then971:
	*state_addr = 79
	goto next_state

if_end972:
	v364 = *result
	tobool973 = (v364 & 1) != 0
	*retval = tobool973
	goto _return

sw_bb974:
	*result = 1
	v365 = *lexer_addr
	result_symbol975 = &v365.F1
	*result_symbol975 = 3
	v366 = *lexer_addr
	mark_end976 = &v366.F3
	v367 = *mark_end976
	v368 = *lexer_addr
	v367(v368)
	v369 = *lookahead
	cmp977 = v369 == 46
	if cmp977 {
		goto if_then979
	} else {
		goto if_end980
	}

if_then979:
	*state_addr = 91
	goto next_state

if_end980:
	v370 = *lookahead
	cmp981 = v370 == 58
	if cmp981 {
		goto if_then983
	} else {
		goto if_end984
	}

if_then983:
	*state_addr = 51
	goto next_state

if_end984:
	v371 = *lookahead
	cmp985 = v371 == 95
	if cmp985 {
		goto if_then987
	} else {
		goto if_end988
	}

if_then987:
	*state_addr = 83
	goto next_state

if_end988:
	v372 = *lookahead
	cmp989 = 48 <= v372
	if cmp989 {
		goto land_lhs_true991
	} else {
		goto if_end995
	}

land_lhs_true991:
	v373 = *lookahead
	cmp992 = v373 <= 57
	if cmp992 {
		goto if_then994
	} else {
		goto if_end995
	}

if_then994:
	*state_addr = 74
	goto next_state

if_end995:
	v374 = *result
	tobool996 = (v374 & 1) != 0
	*retval = tobool996
	goto _return

sw_bb997:
	*result = 1
	v375 = *lexer_addr
	result_symbol998 = &v375.F1
	*result_symbol998 = 3
	v376 = *lexer_addr
	mark_end999 = &v376.F3
	v377 = *mark_end999
	v378 = *lexer_addr
	v377(v378)
	v379 = *lookahead
	cmp1000 = v379 == 46
	if cmp1000 {
		goto if_then1002
	} else {
		goto if_end1003
	}

if_then1002:
	*state_addr = 91
	goto next_state

if_end1003:
	v380 = *lookahead
	cmp1004 = v380 == 58
	if cmp1004 {
		goto if_then1006
	} else {
		goto if_end1007
	}

if_then1006:
	*state_addr = 51
	goto next_state

if_end1007:
	v381 = *lookahead
	cmp1008 = v381 == 95
	if cmp1008 {
		goto if_then1010
	} else {
		goto if_end1011
	}

if_then1010:
	*state_addr = 83
	goto next_state

if_end1011:
	v382 = *lookahead
	cmp1012 = 48 <= v382
	if cmp1012 {
		goto land_lhs_true1014
	} else {
		goto if_end1018
	}

land_lhs_true1014:
	v383 = *lookahead
	cmp1015 = v383 <= 57
	if cmp1015 {
		goto if_then1017
	} else {
		goto if_end1018
	}

if_then1017:
	*state_addr = 80
	goto next_state

if_end1018:
	v384 = *result
	tobool1019 = (v384 & 1) != 0
	*retval = tobool1019
	goto _return

sw_bb1020:
	*result = 1
	v385 = *lexer_addr
	result_symbol1021 = &v385.F1
	*result_symbol1021 = 3
	v386 = *lexer_addr
	mark_end1022 = &v386.F3
	v387 = *mark_end1022
	v388 = *lexer_addr
	v387(v388)
	v389 = *lookahead
	cmp1023 = v389 == 46
	if cmp1023 {
		goto if_then1025
	} else {
		goto if_end1026
	}

if_then1025:
	*state_addr = 91
	goto next_state

if_end1026:
	v390 = *lookahead
	cmp1027 = v390 == 58
	if cmp1027 {
		goto if_then1029
	} else {
		goto if_end1030
	}

if_then1029:
	*state_addr = 51
	goto next_state

if_end1030:
	v391 = *lookahead
	cmp1031 = v391 == 95
	if cmp1031 {
		goto if_then1033
	} else {
		goto if_end1034
	}

if_then1033:
	*state_addr = 83
	goto next_state

if_end1034:
	v392 = *lookahead
	cmp1035 = 48 <= v392
	if cmp1035 {
		goto land_lhs_true1037
	} else {
		goto if_end1041
	}

land_lhs_true1037:
	v393 = *lookahead
	cmp1038 = v393 <= 57
	if cmp1038 {
		goto if_then1040
	} else {
		goto if_end1041
	}

if_then1040:
	*state_addr = 81
	goto next_state

if_end1041:
	v394 = *result
	tobool1042 = (v394 & 1) != 0
	*retval = tobool1042
	goto _return

sw_bb1043:
	*result = 1
	v395 = *lexer_addr
	result_symbol1044 = &v395.F1
	*result_symbol1044 = 3
	v396 = *lexer_addr
	mark_end1045 = &v396.F3
	v397 = *mark_end1045
	v398 = *lexer_addr
	v397(v398)
	v399 = *lookahead
	cmp1046 = v399 == 46
	if cmp1046 {
		goto if_then1048
	} else {
		goto if_end1049
	}

if_then1048:
	*state_addr = 91
	goto next_state

if_end1049:
	v400 = *lookahead
	cmp1050 = v400 == 58
	if cmp1050 {
		goto if_then1052
	} else {
		goto if_end1053
	}

if_then1052:
	*state_addr = 51
	goto next_state

if_end1053:
	v401 = *lookahead
	cmp1054 = 48 <= v401
	if cmp1054 {
		goto land_lhs_true1056
	} else {
		goto lor_lhs_false1059
	}

land_lhs_true1056:
	v402 = *lookahead
	cmp1057 = v402 <= 57
	if cmp1057 {
		goto if_then1062
	} else {
		goto lor_lhs_false1059
	}

lor_lhs_false1059:
	v403 = *lookahead
	cmp1060 = v403 == 95
	if cmp1060 {
		goto if_then1062
	} else {
		goto if_end1063
	}

if_then1062:
	*state_addr = 83
	goto next_state

if_end1063:
	v404 = *result
	tobool1064 = (v404 & 1) != 0
	*retval = tobool1064
	goto _return

sw_bb1065:
	*result = 1
	v405 = *lexer_addr
	result_symbol1066 = &v405.F1
	*result_symbol1066 = 3
	v406 = *lexer_addr
	mark_end1067 = &v406.F3
	v407 = *mark_end1067
	v408 = *lexer_addr
	v407(v408)
	v409 = *lookahead
	cmp1068 = v409 == 46
	if cmp1068 {
		goto if_then1070
	} else {
		goto if_end1071
	}

if_then1070:
	*state_addr = 93
	goto next_state

if_end1071:
	v410 = *lookahead
	cmp1072 = v410 == 58
	if cmp1072 {
		goto if_then1074
	} else {
		goto if_end1075
	}

if_then1074:
	*state_addr = 51
	goto next_state

if_end1075:
	v411 = *result
	tobool1076 = (v411 & 1) != 0
	*retval = tobool1076
	goto _return

sw_bb1077:
	*result = 1
	v412 = *lexer_addr
	result_symbol1078 = &v412.F1
	*result_symbol1078 = 3
	v413 = *lexer_addr
	mark_end1079 = &v413.F3
	v414 = *mark_end1079
	v415 = *lexer_addr
	v414(v415)
	v416 = *lookahead
	cmp1080 = v416 == 46
	if cmp1080 {
		goto if_then1082
	} else {
		goto if_end1083
	}

if_then1082:
	*state_addr = 93
	goto next_state

if_end1083:
	v417 = *lookahead
	cmp1084 = v417 == 58
	if cmp1084 {
		goto if_then1086
	} else {
		goto if_end1087
	}

if_then1086:
	*state_addr = 51
	goto next_state

if_end1087:
	v418 = *lookahead
	cmp1088 = 48 <= v418
	if cmp1088 {
		goto land_lhs_true1090
	} else {
		goto if_end1094
	}

land_lhs_true1090:
	v419 = *lookahead
	cmp1091 = v419 <= 57
	if cmp1091 {
		goto if_then1093
	} else {
		goto if_end1094
	}

if_then1093:
	*state_addr = 84
	goto next_state

if_end1094:
	v420 = *result
	tobool1095 = (v420 & 1) != 0
	*retval = tobool1095
	goto _return

sw_bb1096:
	*result = 1
	v421 = *lexer_addr
	result_symbol1097 = &v421.F1
	*result_symbol1097 = 3
	v422 = *lexer_addr
	mark_end1098 = &v422.F3
	v423 = *mark_end1098
	v424 = *lexer_addr
	v423(v424)
	v425 = *lookahead
	cmp1099 = v425 == 48
	if cmp1099 {
		goto if_then1107
	} else {
		goto lor_lhs_false1101
	}

lor_lhs_false1101:
	v426 = *lookahead
	cmp1102 = v426 == 49
	if cmp1102 {
		goto if_then1107
	} else {
		goto lor_lhs_false1104
	}

lor_lhs_false1104:
	v427 = *lookahead
	cmp1105 = v427 == 95
	if cmp1105 {
		goto if_then1107
	} else {
		goto if_end1108
	}

if_then1107:
	*state_addr = 86
	goto next_state

if_end1108:
	v428 = *result
	tobool1109 = (v428 & 1) != 0
	*retval = tobool1109
	goto _return

sw_bb1110:
	*result = 1
	v429 = *lexer_addr
	result_symbol1111 = &v429.F1
	*result_symbol1111 = 3
	v430 = *lexer_addr
	mark_end1112 = &v430.F3
	v431 = *mark_end1112
	v432 = *lexer_addr
	v431(v432)
	v433 = *lookahead
	cmp1113 = 48 <= v433
	if cmp1113 {
		goto land_lhs_true1115
	} else {
		goto lor_lhs_false1118
	}

land_lhs_true1115:
	v434 = *lookahead
	cmp1116 = v434 <= 57
	if cmp1116 {
		goto if_then1133
	} else {
		goto lor_lhs_false1118
	}

lor_lhs_false1118:
	v435 = *lookahead
	cmp1119 = 65 <= v435
	if cmp1119 {
		goto land_lhs_true1121
	} else {
		goto lor_lhs_false1124
	}

land_lhs_true1121:
	v436 = *lookahead
	cmp1122 = v436 <= 70
	if cmp1122 {
		goto if_then1133
	} else {
		goto lor_lhs_false1124
	}

lor_lhs_false1124:
	v437 = *lookahead
	cmp1125 = v437 == 95
	if cmp1125 {
		goto if_then1133
	} else {
		goto lor_lhs_false1127
	}

lor_lhs_false1127:
	v438 = *lookahead
	cmp1128 = 97 <= v438
	if cmp1128 {
		goto land_lhs_true1130
	} else {
		goto if_end1134
	}

land_lhs_true1130:
	v439 = *lookahead
	cmp1131 = v439 <= 102
	if cmp1131 {
		goto if_then1133
	} else {
		goto if_end1134
	}

if_then1133:
	*state_addr = 87
	goto next_state

if_end1134:
	v440 = *result
	tobool1135 = (v440 & 1) != 0
	*retval = tobool1135
	goto _return

sw_bb1136:
	*result = 1
	v441 = *lexer_addr
	result_symbol1137 = &v441.F1
	*result_symbol1137 = 4
	v442 = *lexer_addr
	mark_end1138 = &v442.F3
	v443 = *mark_end1138
	v444 = *lexer_addr
	v443(v444)
	v445 = *result
	tobool1139 = (v445 & 1) != 0
	*retval = tobool1139
	goto _return

sw_bb1140:
	*result = 1
	v446 = *lexer_addr
	result_symbol1141 = &v446.F1
	*result_symbol1141 = 4
	v447 = *lexer_addr
	mark_end1142 = &v447.F3
	v448 = *mark_end1142
	v449 = *lexer_addr
	v448(v449)
	v450 = *lookahead
	cmp1143 = v450 == 73
	if cmp1143 {
		goto if_then1145
	} else {
		goto if_end1146
	}

if_then1145:
	*state_addr = 23
	goto next_state

if_end1146:
	v451 = *lookahead
	cmp1147 = v451 == 78
	if cmp1147 {
		goto if_then1149
	} else {
		goto if_end1150
	}

if_then1149:
	*state_addr = 14
	goto next_state

if_end1150:
	v452 = *lookahead
	cmp1151 = v452 == 105
	if cmp1151 {
		goto if_then1153
	} else {
		goto if_end1154
	}

if_then1153:
	*state_addr = 39
	goto next_state

if_end1154:
	v453 = *lookahead
	cmp1155 = v453 == 110
	if cmp1155 {
		goto if_then1157
	} else {
		goto if_end1158
	}

if_then1157:
	*state_addr = 30
	goto next_state

if_end1158:
	v454 = *lookahead
	cmp1159 = v454 == 69
	if cmp1159 {
		goto if_then1164
	} else {
		goto lor_lhs_false1161
	}

lor_lhs_false1161:
	v455 = *lookahead
	cmp1162 = v455 == 101
	if cmp1162 {
		goto if_then1164
	} else {
		goto if_end1165
	}

if_then1164:
	*state_addr = 48
	goto next_state

if_end1165:
	v456 = *lookahead
	cmp1166 = v456 == 46
	if cmp1166 {
		goto if_then1174
	} else {
		goto lor_lhs_false1168
	}

lor_lhs_false1168:
	v457 = *lookahead
	cmp1169 = 48 <= v457
	if cmp1169 {
		goto land_lhs_true1171
	} else {
		goto if_end1175
	}

land_lhs_true1171:
	v458 = *lookahead
	cmp1172 = v458 <= 57
	if cmp1172 {
		goto if_then1174
	} else {
		goto if_end1175
	}

if_then1174:
	*state_addr = 91
	goto next_state

if_end1175:
	v459 = *result
	tobool1176 = (v459 & 1) != 0
	*retval = tobool1176
	goto _return

sw_bb1177:
	*result = 1
	v460 = *lexer_addr
	result_symbol1178 = &v460.F1
	*result_symbol1178 = 4
	v461 = *lexer_addr
	mark_end1179 = &v461.F3
	v462 = *mark_end1179
	v463 = *lexer_addr
	v462(v463)
	v464 = *lookahead
	cmp1180 = v464 == 73
	if cmp1180 {
		goto if_then1182
	} else {
		goto if_end1183
	}

if_then1182:
	*state_addr = 23
	goto next_state

if_end1183:
	v465 = *lookahead
	cmp1184 = v465 == 105
	if cmp1184 {
		goto if_then1186
	} else {
		goto if_end1187
	}

if_then1186:
	*state_addr = 39
	goto next_state

if_end1187:
	v466 = *lookahead
	cmp1188 = v466 == 69
	if cmp1188 {
		goto if_then1193
	} else {
		goto lor_lhs_false1190
	}

lor_lhs_false1190:
	v467 = *lookahead
	cmp1191 = v467 == 101
	if cmp1191 {
		goto if_then1193
	} else {
		goto if_end1194
	}

if_then1193:
	*state_addr = 48
	goto next_state

if_end1194:
	v468 = *lookahead
	cmp1195 = v468 == 46
	if cmp1195 {
		goto if_then1203
	} else {
		goto lor_lhs_false1197
	}

lor_lhs_false1197:
	v469 = *lookahead
	cmp1198 = 48 <= v469
	if cmp1198 {
		goto land_lhs_true1200
	} else {
		goto if_end1204
	}

land_lhs_true1200:
	v470 = *lookahead
	cmp1201 = v470 <= 57
	if cmp1201 {
		goto if_then1203
	} else {
		goto if_end1204
	}

if_then1203:
	*state_addr = 91
	goto next_state

if_end1204:
	v471 = *result
	tobool1205 = (v471 & 1) != 0
	*retval = tobool1205
	goto _return

sw_bb1206:
	*result = 1
	v472 = *lexer_addr
	result_symbol1207 = &v472.F1
	*result_symbol1207 = 4
	v473 = *lexer_addr
	mark_end1208 = &v473.F3
	v474 = *mark_end1208
	v475 = *lexer_addr
	v474(v475)
	v476 = *lookahead
	cmp1209 = v476 == 69
	if cmp1209 {
		goto if_then1214
	} else {
		goto lor_lhs_false1211
	}

lor_lhs_false1211:
	v477 = *lookahead
	cmp1212 = v477 == 101
	if cmp1212 {
		goto if_then1214
	} else {
		goto if_end1215
	}

if_then1214:
	*state_addr = 48
	goto next_state

if_end1215:
	v478 = *lookahead
	cmp1216 = v478 == 46
	if cmp1216 {
		goto if_then1224
	} else {
		goto lor_lhs_false1218
	}

lor_lhs_false1218:
	v479 = *lookahead
	cmp1219 = 48 <= v479
	if cmp1219 {
		goto land_lhs_true1221
	} else {
		goto if_end1225
	}

land_lhs_true1221:
	v480 = *lookahead
	cmp1222 = v480 <= 57
	if cmp1222 {
		goto if_then1224
	} else {
		goto if_end1225
	}

if_then1224:
	*state_addr = 91
	goto next_state

if_end1225:
	v481 = *result
	tobool1226 = (v481 & 1) != 0
	*retval = tobool1226
	goto _return

sw_bb1227:
	*result = 1
	v482 = *lexer_addr
	result_symbol1228 = &v482.F1
	*result_symbol1228 = 4
	v483 = *lexer_addr
	mark_end1229 = &v483.F3
	v484 = *mark_end1229
	v485 = *lexer_addr
	v484(v485)
	v486 = *lookahead
	cmp1230 = 48 <= v486
	if cmp1230 {
		goto land_lhs_true1232
	} else {
		goto if_end1236
	}

land_lhs_true1232:
	v487 = *lookahead
	cmp1233 = v487 <= 57
	if cmp1233 {
		goto if_then1235
	} else {
		goto if_end1236
	}

if_then1235:
	*state_addr = 92
	goto next_state

if_end1236:
	v488 = *result
	tobool1237 = (v488 & 1) != 0
	*retval = tobool1237
	goto _return

sw_bb1238:
	*result = 1
	v489 = *lexer_addr
	result_symbol1239 = &v489.F1
	*result_symbol1239 = 4
	v490 = *lexer_addr
	mark_end1240 = &v490.F3
	v491 = *mark_end1240
	v492 = *lexer_addr
	v491(v492)
	v493 = *lookahead
	cmp1241 = 48 <= v493
	if cmp1241 {
		goto land_lhs_true1243
	} else {
		goto lor_lhs_false1246
	}

land_lhs_true1243:
	v494 = *lookahead
	cmp1244 = v494 <= 57
	if cmp1244 {
		goto if_then1249
	} else {
		goto lor_lhs_false1246
	}

lor_lhs_false1246:
	v495 = *lookahead
	cmp1247 = v495 == 95
	if cmp1247 {
		goto if_then1249
	} else {
		goto if_end1250
	}

if_then1249:
	*state_addr = 93
	goto next_state

if_end1250:
	v496 = *result
	tobool1251 = (v496 & 1) != 0
	*retval = tobool1251
	goto _return

sw_bb1252:
	*result = 1
	v497 = *lexer_addr
	result_symbol1253 = &v497.F1
	*result_symbol1253 = 5
	v498 = *lexer_addr
	mark_end1254 = &v498.F3
	v499 = *mark_end1254
	v500 = *lexer_addr
	v499(v500)
	v501 = *result
	tobool1255 = (v501 & 1) != 0
	*retval = tobool1255
	goto _return

sw_bb1256:
	*result = 1
	v502 = *lexer_addr
	result_symbol1257 = &v502.F1
	*result_symbol1257 = 5
	v503 = *lexer_addr
	mark_end1258 = &v503.F3
	v504 = *mark_end1258
	v505 = *lexer_addr
	v504(v505)
	v506 = *lookahead
	cmp1259 = v506 == 46
	if cmp1259 {
		goto if_then1261
	} else {
		goto if_end1262
	}

if_then1261:
	*state_addr = 58
	goto next_state

if_end1262:
	v507 = *lookahead
	cmp1263 = v507 == 90
	if cmp1263 {
		goto if_then1265
	} else {
		goto if_end1266
	}

if_then1265:
	*state_addr = 94
	goto next_state

if_end1266:
	v508 = *lookahead
	cmp1267 = v508 == 9
	if cmp1267 {
		goto if_then1272
	} else {
		goto lor_lhs_false1269
	}

lor_lhs_false1269:
	v509 = *lookahead
	cmp1270 = v509 == 32
	if cmp1270 {
		goto if_then1272
	} else {
		goto if_end1273
	}

if_then1272:
	*state_addr = 28
	goto next_state

if_end1273:
	v510 = *lookahead
	cmp1274 = v510 == 43
	if cmp1274 {
		goto if_then1279
	} else {
		goto lor_lhs_false1276
	}

lor_lhs_false1276:
	v511 = *lookahead
	cmp1277 = v511 == 45
	if cmp1277 {
		goto if_then1279
	} else {
		goto if_end1280
	}

if_then1279:
	*state_addr = 59
	goto next_state

if_end1280:
	v512 = *result
	tobool1281 = (v512 & 1) != 0
	*retval = tobool1281
	goto _return

sw_bb1282:
	*result = 1
	v513 = *lexer_addr
	result_symbol1283 = &v513.F1
	*result_symbol1283 = 5
	v514 = *lexer_addr
	mark_end1284 = &v514.F3
	v515 = *mark_end1284
	v516 = *lexer_addr
	v515(v516)
	v517 = *lookahead
	cmp1285 = v517 == 58
	if cmp1285 {
		goto if_then1287
	} else {
		goto if_end1288
	}

if_then1287:
	*state_addr = 63
	goto next_state

if_end1288:
	v518 = *result
	tobool1289 = (v518 & 1) != 0
	*retval = tobool1289
	goto _return

sw_bb1290:
	*result = 1
	v519 = *lexer_addr
	result_symbol1291 = &v519.F1
	*result_symbol1291 = 5
	v520 = *lexer_addr
	mark_end1292 = &v520.F3
	v521 = *mark_end1292
	v522 = *lexer_addr
	v521(v522)
	v523 = *lookahead
	cmp1293 = v523 == 58
	if cmp1293 {
		goto if_then1295
	} else {
		goto if_end1296
	}

if_then1295:
	*state_addr = 63
	goto next_state

if_end1296:
	v524 = *lookahead
	cmp1297 = 48 <= v524
	if cmp1297 {
		goto land_lhs_true1299
	} else {
		goto if_end1303
	}

land_lhs_true1299:
	v525 = *lookahead
	cmp1300 = v525 <= 57
	if cmp1300 {
		goto if_then1302
	} else {
		goto if_end1303
	}

if_then1302:
	*state_addr = 96
	goto next_state

if_end1303:
	v526 = *result
	tobool1304 = (v526 & 1) != 0
	*retval = tobool1304
	goto _return

sw_bb1305:
	*result = 1
	v527 = *lexer_addr
	result_symbol1306 = &v527.F1
	*result_symbol1306 = 5
	v528 = *lexer_addr
	mark_end1307 = &v528.F3
	v529 = *mark_end1307
	v530 = *lexer_addr
	v529(v530)
	v531 = *lookahead
	cmp1308 = v531 == 90
	if cmp1308 {
		goto if_then1310
	} else {
		goto if_end1311
	}

if_then1310:
	*state_addr = 94
	goto next_state

if_end1311:
	v532 = *lookahead
	cmp1312 = v532 == 9
	if cmp1312 {
		goto if_then1317
	} else {
		goto lor_lhs_false1314
	}

lor_lhs_false1314:
	v533 = *lookahead
	cmp1315 = v533 == 32
	if cmp1315 {
		goto if_then1317
	} else {
		goto if_end1318
	}

if_then1317:
	*state_addr = 28
	goto next_state

if_end1318:
	v534 = *lookahead
	cmp1319 = v534 == 43
	if cmp1319 {
		goto if_then1324
	} else {
		goto lor_lhs_false1321
	}

lor_lhs_false1321:
	v535 = *lookahead
	cmp1322 = v535 == 45
	if cmp1322 {
		goto if_then1324
	} else {
		goto if_end1325
	}

if_then1324:
	*state_addr = 59
	goto next_state

if_end1325:
	v536 = *lookahead
	cmp1326 = 48 <= v536
	if cmp1326 {
		goto land_lhs_true1328
	} else {
		goto if_end1332
	}

land_lhs_true1328:
	v537 = *lookahead
	cmp1329 = v537 <= 57
	if cmp1329 {
		goto if_then1331
	} else {
		goto if_end1332
	}

if_then1331:
	*state_addr = 98
	goto next_state

if_end1332:
	v538 = *result
	tobool1333 = (v538 & 1) != 0
	*retval = tobool1333
	goto _return

sw_bb1334:
	*result = 1
	v539 = *lexer_addr
	result_symbol1335 = &v539.F1
	*result_symbol1335 = 5
	v540 = *lexer_addr
	mark_end1336 = &v540.F3
	v541 = *mark_end1336
	v542 = *lexer_addr
	v541(v542)
	v543 = *lookahead
	cmp1337 = v543 == 9
	if cmp1337 {
		goto if_then1342
	} else {
		goto lor_lhs_false1339
	}

lor_lhs_false1339:
	v544 = *lookahead
	cmp1340 = v544 == 32
	if cmp1340 {
		goto if_then1342
	} else {
		goto if_end1343
	}

if_then1342:
	*state_addr = 47
	goto next_state

if_end1343:
	v545 = *lookahead
	cmp1344 = v545 == 84
	if cmp1344 {
		goto if_then1349
	} else {
		goto lor_lhs_false1346
	}

lor_lhs_false1346:
	v546 = *lookahead
	cmp1347 = v546 == 116
	if cmp1347 {
		goto if_then1349
	} else {
		goto if_end1350
	}

if_then1349:
	*state_addr = 55
	goto next_state

if_end1350:
	v547 = *result
	tobool1351 = (v547 & 1) != 0
	*retval = tobool1351
	goto _return

sw_default:
	*retval = false
	goto _return

_return:
	v548 = *retval
	return v548
}

