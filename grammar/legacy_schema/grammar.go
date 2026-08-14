package grammar_legacy_schema

import (
	"github.com/lewtec/leaven/libc"
	"unsafe"
)

type TSLanguageMetadata struct {
	F0 byte
	F1 byte
	F2 byte
}
type TSSymbolMetadata struct {
	F0 byte
	F1 byte
	F2 byte
}
type TSLexMode struct {
	F0 int16
	F1 int16
}
type TSParseActionEntry struct {
	F0 TSParseAction
}
type TSParseAction struct {
	F0 anon_1
}
type anon_1 struct {
	F0 byte
	F1 byte
	F2 int16
	F3 int16
	F4 int16
}
type anon_2 struct {
	F0 byte
	F1 byte
}
type TSLexer struct {
	F0 int32
	F1 int16
	F2 unsafe.Pointer
	F3 unsafe.Pointer
	F4 unsafe.Pointer
	F5 unsafe.Pointer
	F6 unsafe.Pointer
	F7 unsafe.Pointer
}

var tree_sitter_legacy_schema_language struct {
	F0  int32
	F1  int32
	F2  int32
	F3  int32
	F4  int32
	F5  int32
	F6  int32
	F7  int32
	F8  int32
	F9  int16
	F10 [2]byte
	F11 unsafe.Pointer
	F12 unsafe.Pointer
	F13 unsafe.Pointer
	F14 unsafe.Pointer
	F15 unsafe.Pointer
	F16 unsafe.Pointer
	F17 unsafe.Pointer
	F18 unsafe.Pointer
	F19 unsafe.Pointer
	F20 unsafe.Pointer
	F21 unsafe.Pointer
	F22 unsafe.Pointer
	F23 unsafe.Pointer
	F24 unsafe.Pointer
	F25 unsafe.Pointer
	F26 int16
	F27 [6]byte
	F28 struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}
	F29 unsafe.Pointer
	F30 unsafe.Pointer
	F31 unsafe.Pointer
	F32 int16
	F33 [2]byte
	F34 int32
	F35 unsafe.Pointer
	F36 unsafe.Pointer
	F37 unsafe.Pointer
	F38 TSLanguageMetadata
	F39 [5]byte
}
var ts_parse_table [2][7]int16 = [2][7]int16{[7]int16{1, 1, 1, 1, 1, 1, 0}, [7]int16{0, 3, 5, 5, 3, 3, 3}}
var ts_small_parse_table [8]int16 = [8]int16{1, 7, 1, 0, 1, 9, 1, 0}
var ts_small_parse_table_map [2]int32 = [2]int32{0, 4}
var ts_symbol_names [7]unsafe.Pointer = [7]unsafe.Pointer{libc.Ptr(&_str), libc.Ptr(&_str_2), libc.Ptr(&_str_3), libc.Ptr(&_str_4), libc.Ptr(&_str_5), libc.Ptr(&_str_6), libc.Ptr(&_str_7)}
var ts_symbol_metadata [7]TSSymbolMetadata = [7]TSSymbolMetadata{TSSymbolMetadata{0, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}, TSSymbolMetadata{1, 1, 0}}
var ts_symbol_map [7]int16 = [7]int16{0, 1, 2, 3, 4, 5, 6}
var ts_non_terminal_alias_map [1]int16 = [1]int16{}
var ts_alias_sequences [1][1]int16 = [1][1]int16{}
var ts_lex_modes [4]TSLexMode = [4]TSLexMode{}
var ts_primary_state_ids [4]int16 = [4]int16{0, 1, 2, 3}
var ts_parse_actions struct {
	F0 struct {
		F0 anon_2
		F1 [6]byte
	}
	F1 struct {
		F0 anon_2
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_2
		F1 [6]byte
	}
	F4 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F5 struct {
		F0 anon_2
		F1 [6]byte
	}
	F6 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F7 struct {
		F0 anon_2
		F1 [6]byte
	}
	F8 TSParseActionEntry
	F9 struct {
		F0 anon_2
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
		F0 anon_2
		F1 [6]byte
	}
	F1 struct {
		F0 anon_2
		F1 [6]byte
	}
	F2 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
	F3 struct {
		F0 anon_2
		F1 [6]byte
	}
	F4 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F5 struct {
		F0 anon_2
		F1 [6]byte
	}
	F6 struct {
		F0 struct {
			F0 struct {
				F0 byte
				F1 byte
				F2 int16
				F3 byte
				F4 byte
			}
			F1 [2]byte
		}
	}
	F7 struct {
		F0 anon_2
		F1 [6]byte
	}
	F8 TSParseActionEntry
	F9 struct {
		F0 anon_2
		F1 [6]byte
	}
	F10 struct {
		F0 struct {
			F0 byte
			F1 [7]byte
		}
	}
}{struct {
	F0 anon_2
	F1 [6]byte
}{}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 byte
		F1 [7]byte
	}
}{struct {
	F0 byte
	F1 [7]byte
}{3, [7]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 0}, [6]byte{}}, struct {
	F0 struct {
		F0 struct {
			F0 byte
			F1 byte
			F2 int16
			F3 byte
			F4 byte
		}
		F1 [2]byte
	}
}{struct {
	F0 struct {
		F0 byte
		F1 byte
		F2 int16
		F3 byte
		F4 byte
	}
	F1 [2]byte
}{struct {
	F0 byte
	F1 byte
	F2 int16
	F3 byte
	F4 byte
}{0, 0, 2, 0, 0}, [2]byte{}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, TSParseActionEntry{TSParseAction{anon_1{1, 1, 6, 0, 0}}}, struct {
	F0 anon_2
	F1 [6]byte
}{anon_2{1, 1}, [6]byte{}}, struct {
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
var ts_lex_map [30]int16 = [30]int16{46, 89, 48, 75, 70, 13, 78, 70, 79, 17, 84, 24, 89, 69, 102, 29, 110, 72, 111, 33, 116, 40, 121, 71, 126, 67, 43, 7, 45, 7}

func init() {
	tree_sitter_legacy_schema_language = struct {
		F0  int32
		F1  int32
		F2  int32
		F3  int32
		F4  int32
		F5  int32
		F6  int32
		F7  int32
		F8  int32
		F9  int16
		F10 [2]byte
		F11 unsafe.Pointer
		F12 unsafe.Pointer
		F13 unsafe.Pointer
		F14 unsafe.Pointer
		F15 unsafe.Pointer
		F16 unsafe.Pointer
		F17 unsafe.Pointer
		F18 unsafe.Pointer
		F19 unsafe.Pointer
		F20 unsafe.Pointer
		F21 unsafe.Pointer
		F22 unsafe.Pointer
		F23 unsafe.Pointer
		F24 unsafe.Pointer
		F25 unsafe.Pointer
		F26 int16
		F27 [6]byte
		F28 struct {
			F0 unsafe.Pointer
			F1 unsafe.Pointer
			F2 unsafe.Pointer
			F3 unsafe.Pointer
			F4 unsafe.Pointer
			F5 unsafe.Pointer
			F6 unsafe.Pointer
		}
		F29 unsafe.Pointer
		F30 unsafe.Pointer
		F31 unsafe.Pointer
		F32 int16
		F33 [2]byte
		F34 int32
		F35 unsafe.Pointer
		F36 unsafe.Pointer
		F37 unsafe.Pointer
		F38 TSLanguageMetadata
		F39 [5]byte
	}{14, 7, 0, 6, 0, 4, 2, 1, 0, 1, [2]byte{}, libc.Ptr(&ts_parse_table), libc.Ptr(&ts_small_parse_table), libc.Ptr(&ts_small_parse_table_map), libc.Ptr(&ts_parse_actions), libc.Ptr(&ts_symbol_names), nil, nil, nil, libc.Ptr(&ts_symbol_metadata), libc.Ptr(&ts_symbol_map), libc.Ptr(&ts_non_terminal_alias_map), libc.Ptr(&ts_alias_sequences), libc.Ptr(&ts_lex_modes), libc.FuncCode(ts_lex), nil, 0, [6]byte{}, struct {
		F0 unsafe.Pointer
		F1 unsafe.Pointer
		F2 unsafe.Pointer
		F3 unsafe.Pointer
		F4 unsafe.Pointer
		F5 unsafe.Pointer
		F6 unsafe.Pointer
	}{}, libc.Ptr(&ts_primary_state_ids), nil, nil, 0, [2]byte{}, 0, nil, nil, nil, TSLanguageMetadata{}, [5]byte{}}
}
func tree_sitter_legacy_schema() unsafe.Pointer {
	return libc.Ptr(&tree_sitter_legacy_schema_language)
}
func ts_lex(lexer unsafe.Pointer, state int16) bool {
	var loadedv, call, loadedv3, cmp, cmp7, cmp14, cmp16, loadedv20, cmp22, cmp26, cmp30, cmp34, cmp37, cmp39, loadedv43, cmp45, cmp49, cmp52, loadedv56, cmp58, loadedv62, cmp64, cmp68, cmp72, cmp76, cmp79, loadedv83, cmp85, cmp89, cmp93, cmp97, cmp100, loadedv104, cmp106, cmp110, cmp114, cmp117, cmp120, loadedv124, cmp126, cmp130, cmp134, cmp137, loadedv141, cmp143, cmp147, loadedv151, cmp153, cmp157, cmp161, cmp164, loadedv168, cmp170, loadedv174, cmp176, cmp180, cmp183, loadedv187, cmp189, loadedv193, cmp195, cmp199, loadedv203, cmp205, cmp209, loadedv213, cmp215, loadedv219, cmp221, loadedv225, cmp227, cmp231, cmp235, cmp238, loadedv242, cmp244, loadedv248, cmp250, loadedv254, cmp256, loadedv260, cmp262, loadedv266, cmp268, loadedv272, cmp274, cmp278, loadedv282, cmp284, cmp288, loadedv292, cmp294, loadedv298, cmp300, loadedv304, cmp306, loadedv310, cmp312, cmp316, cmp319, loadedv323, cmp325, loadedv329, cmp331, loadedv335, cmp337, loadedv341, cmp343, loadedv347, cmp349, cmp353, loadedv357, cmp359, loadedv363, cmp365, loadedv369, cmp371, loadedv375, cmp377, loadedv381, cmp383, loadedv387, cmp389, loadedv393, cmp395, loadedv399, cmp401, loadedv405, cmp407, loadedv411, cmp413, loadedv417, cmp419, cmp422, cmp426, cmp429, loadedv433, cmp435, cmp438, cmp442, cmp445, cmp449, cmp452, loadedv456, cmp458, cmp461, cmp465, cmp468, cmp472, cmp475, loadedv479, cmp481, cmp484, cmp488, cmp491, loadedv495, cmp497, cmp500, loadedv504, cmp506, cmp509, cmp512, loadedv516, cmp518, cmp521, cmp525, cmp528, loadedv532, cmp534, cmp537, cmp541, cmp544, loadedv548, cmp550, cmp553, loadedv557, cmp559, cmp562, loadedv566, cmp568, cmp571, loadedv575, cmp577, cmp580, loadedv584, cmp586, cmp589, loadedv593, cmp595, cmp598, loadedv602, cmp604, cmp607, loadedv611, cmp613, cmp616, loadedv620, cmp622, cmp625, loadedv629, cmp631, cmp634, loadedv638, cmp640, cmp643, loadedv647, cmp649, cmp652, loadedv656, cmp658, cmp661, loadedv665, cmp667, cmp670, cmp673, cmp676, cmp679, cmp682, cmp685, loadedv689, loadedv691, loadedv695, loadedv699, cmp703, cmp707, loadedv711, cmp715, cmp719, cmp723, cmp726, loadedv730, cmp734, loadedv738, cmp742, cmp746, loadedv750, cmp754, cmp758, cmp762, cmp766, cmp769, cmp773, cmp776, cmp779, loadedv783, cmp787, cmp791, cmp795, cmp799, cmp802, cmp805, loadedv809, cmp813, cmp817, cmp821, cmp825, cmp829, cmp833, cmp836, cmp840, cmp843, loadedv847, cmp851, cmp855, cmp859, cmp863, cmp866, cmp870, cmp873, loadedv877, cmp881, cmp885, cmp889, cmp893, cmp896, cmp900, cmp903, loadedv907, cmp911, cmp915, cmp919, cmp923, cmp927, cmp930, cmp934, cmp937, cmp940, loadedv944, cmp948, cmp952, cmp956, cmp959, cmp963, cmp966, cmp969, loadedv973, cmp977, cmp981, cmp985, cmp989, cmp992, loadedv996, cmp1000, cmp1004, cmp1008, cmp1012, cmp1015, loadedv1019, cmp1023, cmp1027, cmp1031, cmp1035, cmp1038, loadedv1042, cmp1046, cmp1050, cmp1054, cmp1057, cmp1060, loadedv1064, cmp1068, cmp1072, loadedv1076, cmp1080, cmp1084, cmp1088, cmp1091, loadedv1095, cmp1099, cmp1102, cmp1105, loadedv1109, cmp1113, cmp1116, cmp1119, cmp1122, cmp1125, cmp1128, cmp1131, loadedv1135, loadedv1139, cmp1143, cmp1147, cmp1151, cmp1155, cmp1159, cmp1162, cmp1166, cmp1169, cmp1172, loadedv1176, cmp1180, cmp1184, cmp1188, cmp1191, cmp1195, cmp1198, cmp1201, loadedv1205, cmp1209, cmp1212, cmp1216, cmp1219, cmp1222, loadedv1226, cmp1230, cmp1233, loadedv1237, cmp1241, cmp1244, cmp1247, loadedv1251, loadedv1255, cmp1259, cmp1263, cmp1267, cmp1270, cmp1274, cmp1277, loadedv1281, cmp1285, loadedv1289, cmp1293, cmp1297, cmp1300, loadedv1304, cmp1308, cmp1312, cmp1315, cmp1319, cmp1322, cmp1326, cmp1329, loadedv1333, cmp1337, cmp1340, cmp1344, cmp1347, loadedv1351, v548 bool
	var retval unsafe.Pointer
	var v9, v13, v16 int16
	var state_addr, arrayidx, arrayidx11, result_symbol, result_symbol693, result_symbol697, result_symbol701, result_symbol713, result_symbol732, result_symbol740, result_symbol752, result_symbol785, result_symbol811, result_symbol849, result_symbol879, result_symbol909, result_symbol946, result_symbol975, result_symbol998, result_symbol1021, result_symbol1044, result_symbol1066, result_symbol1078, result_symbol1097, result_symbol1111, result_symbol1137, result_symbol1141, result_symbol1178, result_symbol1207, result_symbol1228, result_symbol1239, result_symbol1253, result_symbol1257, result_symbol1283, result_symbol1291, result_symbol1306, result_symbol1335 unsafe.Pointer
	var v5, conv, v11, v12, conv6, v14, v15, add, v17, add13, v18, v19, v21, v22, v23, v24, v25, v26, v28, v29, v30, v32, v34, v35, v36, v37, v38, v40, v41, v42, v43, v44, v46, v47, v48, v49, v50, v52, v53, v54, v55, v57, v58, v60, v61, v62, v63, v65, v67, v68, v69, v71, v73, v74, v76, v77, v79, v81, v83, v84, v85, v86, v88, v90, v92, v94, v96, v98, v99, v101, v102, v104, v106, v108, v110, v111, v112, v114, v116, v118, v120, v122, v123, v125, v127, v129, v131, v133, v135, v137, v139, v141, v143, v145, v146, v147, v148, v150, v151, v152, v153, v154, v155, v157, v158, v159, v160, v161, v162, v164, v165, v166, v167, v169, v170, v172, v173, v174, v176, v177, v178, v179, v181, v182, v183, v184, v186, v187, v189, v190, v192, v193, v195, v196, v198, v199, v201, v202, v204, v205, v207, v208, v210, v211, v213, v214, v216, v217, v219, v220, v222, v223, v225, v226, v227, v228, v229, v230, v231, v252, v253, v259, v260, v261, v262, v268, v274, v275, v281, v282, v283, v284, v285, v286, v287, v288, v294, v295, v296, v297, v298, v299, v305, v306, v307, v308, v309, v310, v311, v312, v313, v319, v320, v321, v322, v323, v324, v325, v331, v332, v333, v334, v335, v336, v337, v343, v344, v345, v346, v347, v348, v349, v350, v351, v357, v358, v359, v360, v361, v362, v363, v369, v370, v371, v372, v373, v379, v380, v381, v382, v383, v389, v390, v391, v392, v393, v399, v400, v401, v402, v403, v409, v410, v416, v417, v418, v419, v425, v426, v427, v433, v434, v435, v436, v437, v438, v439, v450, v451, v452, v453, v454, v455, v456, v457, v458, v464, v465, v466, v467, v468, v469, v470, v476, v477, v478, v479, v480, v486, v487, v493, v494, v495, v506, v507, v508, v509, v510, v511, v517, v523, v524, v525, v531, v532, v533, v534, v535, v536, v537, v543, v544, v545, v546 int32
	var lookahead, i, lookahead1 unsafe.Pointer
	var conv4, idxprom, idxprom10 int64
	var v3, storedv, v10, v20, v27, v31, v33, v39, v45, v51, v56, v59, v64, v66, v70, v72, v75, v78, v80, v82, v87, v89, v91, v93, v95, v97, v100, v103, v105, v107, v109, v113, v115, v117, v119, v121, v124, v126, v128, v130, v132, v134, v136, v138, v140, v142, v144, v149, v156, v163, v168, v171, v175, v180, v185, v188, v191, v194, v197, v200, v203, v206, v209, v212, v215, v218, v221, v224, v232, v237, v242, v247, v254, v263, v269, v276, v289, v300, v314, v326, v338, v352, v364, v374, v384, v394, v404, v411, v420, v428, v440, v445, v459, v471, v481, v488, v496, v501, v512, v518, v526, v538, v547 byte
	var result, skip, eof unsafe.Pointer
	var v0, v1, v2, v4, v6, v7, v8, v233, v234, v235, v236, v238, v239, v240, v241, v243, v244, v245, v246, v248, v249, v250, v251, v255, v256, v257, v258, v264, v265, v266, v267, v270, v271, v272, v273, v277, v278, v279, v280, v290, v291, v292, v293, v301, v302, v303, v304, v315, v316, v317, v318, v327, v328, v329, v330, v339, v340, v341, v342, v353, v354, v355, v356, v365, v366, v367, v368, v375, v376, v377, v378, v385, v386, v387, v388, v395, v396, v397, v398, v405, v406, v407, v408, v412, v413, v414, v415, v421, v422, v423, v424, v429, v430, v431, v432, v441, v442, v443, v444, v446, v447, v448, v449, v460, v461, v462, v463, v472, v473, v474, v475, v482, v483, v484, v485, v489, v490, v491, v492, v497, v498, v499, v500, v502, v503, v504, v505, v513, v514, v515, v516, v519, v520, v521, v522, v527, v528, v529, v530, v539, v540, v541, v542 unsafe.Pointer
	var lexer_addr, advance, eof2, mark_end, mark_end694, mark_end698, mark_end702, mark_end714, mark_end733, mark_end741, mark_end753, mark_end786, mark_end812, mark_end850, mark_end880, mark_end910, mark_end947, mark_end976, mark_end999, mark_end1022, mark_end1045, mark_end1067, mark_end1079, mark_end1098, mark_end1112, mark_end1138, mark_end1142, mark_end1179, mark_end1208, mark_end1229, mark_end1240, mark_end1254, mark_end1258, mark_end1284, mark_end1292, mark_end1307, mark_end1336 unsafe.Pointer
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = retval, lexer_addr, state_addr, result, skip, eof, lookahead, i, v0, advance, v1, v2, v3, loadedv, v4, lookahead1, v5, v6, eof2, v7, v8, call, storedv, v9, conv, v10, loadedv3, v11, conv4, cmp, v12, idxprom, arrayidx, v13, conv6, v14, cmp7, v15, add, idxprom10, arrayidx11, v16, v17, add13, v18, cmp14, v19, cmp16, v20, loadedv20, v21, cmp22, v22, cmp26, v23, cmp30, v24, cmp34, v25, cmp37, v26, cmp39, v27, loadedv43, v28, cmp45, v29, cmp49, v30, cmp52, v31, loadedv56, v32, cmp58, v33, loadedv62, v34, cmp64, v35, cmp68, v36, cmp72, v37, cmp76, v38, cmp79, v39, loadedv83, v40, cmp85, v41, cmp89, v42, cmp93, v43, cmp97, v44, cmp100, v45, loadedv104, v46, cmp106, v47, cmp110, v48, cmp114, v49, cmp117, v50, cmp120, v51, loadedv124, v52, cmp126, v53, cmp130, v54, cmp134, v55, cmp137, v56, loadedv141, v57, cmp143, v58, cmp147, v59, loadedv151, v60, cmp153, v61, cmp157, v62, cmp161, v63, cmp164, v64, loadedv168, v65, cmp170, v66, loadedv174, v67, cmp176, v68, cmp180, v69, cmp183, v70, loadedv187, v71, cmp189, v72, loadedv193, v73, cmp195, v74, cmp199, v75, loadedv203, v76, cmp205, v77, cmp209, v78, loadedv213, v79, cmp215, v80, loadedv219, v81, cmp221, v82, loadedv225, v83, cmp227, v84, cmp231, v85, cmp235, v86, cmp238, v87, loadedv242, v88, cmp244, v89, loadedv248, v90, cmp250, v91, loadedv254, v92, cmp256, v93, loadedv260, v94, cmp262, v95, loadedv266, v96, cmp268, v97, loadedv272, v98, cmp274, v99, cmp278, v100, loadedv282, v101, cmp284, v102, cmp288, v103, loadedv292, v104, cmp294, v105, loadedv298, v106, cmp300, v107, loadedv304, v108, cmp306, v109, loadedv310, v110, cmp312, v111, cmp316, v112, cmp319, v113, loadedv323, v114, cmp325, v115, loadedv329, v116, cmp331, v117, loadedv335, v118, cmp337, v119, loadedv341, v120, cmp343, v121, loadedv347, v122, cmp349, v123, cmp353, v124, loadedv357, v125, cmp359, v126, loadedv363, v127, cmp365, v128, loadedv369, v129, cmp371, v130, loadedv375, v131, cmp377, v132, loadedv381, v133, cmp383, v134, loadedv387, v135, cmp389, v136, loadedv393, v137, cmp395, v138, loadedv399, v139, cmp401, v140, loadedv405, v141, cmp407, v142, loadedv411, v143, cmp413, v144, loadedv417, v145, cmp419, v146, cmp422, v147, cmp426, v148, cmp429, v149, loadedv433, v150, cmp435, v151, cmp438, v152, cmp442, v153, cmp445, v154, cmp449, v155, cmp452, v156, loadedv456, v157, cmp458, v158, cmp461, v159, cmp465, v160, cmp468, v161, cmp472, v162, cmp475, v163, loadedv479, v164, cmp481, v165, cmp484, v166, cmp488, v167, cmp491, v168, loadedv495, v169, cmp497, v170, cmp500, v171, loadedv504, v172, cmp506, v173, cmp509, v174, cmp512, v175, loadedv516, v176, cmp518, v177, cmp521, v178, cmp525, v179, cmp528, v180, loadedv532, v181, cmp534, v182, cmp537, v183, cmp541, v184, cmp544, v185, loadedv548, v186, cmp550, v187, cmp553, v188, loadedv557, v189, cmp559, v190, cmp562, v191, loadedv566, v192, cmp568, v193, cmp571, v194, loadedv575, v195, cmp577, v196, cmp580, v197, loadedv584, v198, cmp586, v199, cmp589, v200, loadedv593, v201, cmp595, v202, cmp598, v203, loadedv602, v204, cmp604, v205, cmp607, v206, loadedv611, v207, cmp613, v208, cmp616, v209, loadedv620, v210, cmp622, v211, cmp625, v212, loadedv629, v213, cmp631, v214, cmp634, v215, loadedv638, v216, cmp640, v217, cmp643, v218, loadedv647, v219, cmp649, v220, cmp652, v221, loadedv656, v222, cmp658, v223, cmp661, v224, loadedv665, v225, cmp667, v226, cmp670, v227, cmp673, v228, cmp676, v229, cmp679, v230, cmp682, v231, cmp685, v232, loadedv689, v233, result_symbol, v234, mark_end, v235, v236, v237, loadedv691, v238, result_symbol693, v239, mark_end694, v240, v241, v242, loadedv695, v243, result_symbol697, v244, mark_end698, v245, v246, v247, loadedv699, v248, result_symbol701, v249, mark_end702, v250, v251, v252, cmp703, v253, cmp707, v254, loadedv711, v255, result_symbol713, v256, mark_end714, v257, v258, v259, cmp715, v260, cmp719, v261, cmp723, v262, cmp726, v263, loadedv730, v264, result_symbol732, v265, mark_end733, v266, v267, v268, cmp734, v269, loadedv738, v270, result_symbol740, v271, mark_end741, v272, v273, v274, cmp742, v275, cmp746, v276, loadedv750, v277, result_symbol752, v278, mark_end753, v279, v280, v281, cmp754, v282, cmp758, v283, cmp762, v284, cmp766, v285, cmp769, v286, cmp773, v287, cmp776, v288, cmp779, v289, loadedv783, v290, result_symbol785, v291, mark_end786, v292, v293, v294, cmp787, v295, cmp791, v296, cmp795, v297, cmp799, v298, cmp802, v299, cmp805, v300, loadedv809, v301, result_symbol811, v302, mark_end812, v303, v304, v305, cmp813, v306, cmp817, v307, cmp821, v308, cmp825, v309, cmp829, v310, cmp833, v311, cmp836, v312, cmp840, v313, cmp843, v314, loadedv847, v315, result_symbol849, v316, mark_end850, v317, v318, v319, cmp851, v320, cmp855, v321, cmp859, v322, cmp863, v323, cmp866, v324, cmp870, v325, cmp873, v326, loadedv877, v327, result_symbol879, v328, mark_end880, v329, v330, v331, cmp881, v332, cmp885, v333, cmp889, v334, cmp893, v335, cmp896, v336, cmp900, v337, cmp903, v338, loadedv907, v339, result_symbol909, v340, mark_end910, v341, v342, v343, cmp911, v344, cmp915, v345, cmp919, v346, cmp923, v347, cmp927, v348, cmp930, v349, cmp934, v350, cmp937, v351, cmp940, v352, loadedv944, v353, result_symbol946, v354, mark_end947, v355, v356, v357, cmp948, v358, cmp952, v359, cmp956, v360, cmp959, v361, cmp963, v362, cmp966, v363, cmp969, v364, loadedv973, v365, result_symbol975, v366, mark_end976, v367, v368, v369, cmp977, v370, cmp981, v371, cmp985, v372, cmp989, v373, cmp992, v374, loadedv996, v375, result_symbol998, v376, mark_end999, v377, v378, v379, cmp1000, v380, cmp1004, v381, cmp1008, v382, cmp1012, v383, cmp1015, v384, loadedv1019, v385, result_symbol1021, v386, mark_end1022, v387, v388, v389, cmp1023, v390, cmp1027, v391, cmp1031, v392, cmp1035, v393, cmp1038, v394, loadedv1042, v395, result_symbol1044, v396, mark_end1045, v397, v398, v399, cmp1046, v400, cmp1050, v401, cmp1054, v402, cmp1057, v403, cmp1060, v404, loadedv1064, v405, result_symbol1066, v406, mark_end1067, v407, v408, v409, cmp1068, v410, cmp1072, v411, loadedv1076, v412, result_symbol1078, v413, mark_end1079, v414, v415, v416, cmp1080, v417, cmp1084, v418, cmp1088, v419, cmp1091, v420, loadedv1095, v421, result_symbol1097, v422, mark_end1098, v423, v424, v425, cmp1099, v426, cmp1102, v427, cmp1105, v428, loadedv1109, v429, result_symbol1111, v430, mark_end1112, v431, v432, v433, cmp1113, v434, cmp1116, v435, cmp1119, v436, cmp1122, v437, cmp1125, v438, cmp1128, v439, cmp1131, v440, loadedv1135, v441, result_symbol1137, v442, mark_end1138, v443, v444, v445, loadedv1139, v446, result_symbol1141, v447, mark_end1142, v448, v449, v450, cmp1143, v451, cmp1147, v452, cmp1151, v453, cmp1155, v454, cmp1159, v455, cmp1162, v456, cmp1166, v457, cmp1169, v458, cmp1172, v459, loadedv1176, v460, result_symbol1178, v461, mark_end1179, v462, v463, v464, cmp1180, v465, cmp1184, v466, cmp1188, v467, cmp1191, v468, cmp1195, v469, cmp1198, v470, cmp1201, v471, loadedv1205, v472, result_symbol1207, v473, mark_end1208, v474, v475, v476, cmp1209, v477, cmp1212, v478, cmp1216, v479, cmp1219, v480, cmp1222, v481, loadedv1226, v482, result_symbol1228, v483, mark_end1229, v484, v485, v486, cmp1230, v487, cmp1233, v488, loadedv1237, v489, result_symbol1239, v490, mark_end1240, v491, v492, v493, cmp1241, v494, cmp1244, v495, cmp1247, v496, loadedv1251, v497, result_symbol1253, v498, mark_end1254, v499, v500, v501, loadedv1255, v502, result_symbol1257, v503, mark_end1258, v504, v505, v506, cmp1259, v507, cmp1263, v508, cmp1267, v509, cmp1270, v510, cmp1274, v511, cmp1277, v512, loadedv1281, v513, result_symbol1283, v514, mark_end1284, v515, v516, v517, cmp1285, v518, loadedv1289, v519, result_symbol1291, v520, mark_end1292, v521, v522, v523, cmp1293, v524, cmp1297, v525, cmp1300, v526, loadedv1304, v527, result_symbol1306, v528, mark_end1307, v529, v530, v531, cmp1308, v532, cmp1312, v533, cmp1315, v534, cmp1319, v535, cmp1322, v536, cmp1326, v537, cmp1329, v538, loadedv1333, v539, result_symbol1335, v540, mark_end1336, v541, v542, v543, cmp1337, v544, cmp1340, v545, cmp1344, v546, cmp1347, v547, loadedv1351, v548

	retval = libc.Ptr(&new(struct {
		_ [0]uint64
		v bool
		b byte
	}).v)
	lexer_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v unsafe.Pointer
		b byte
	}).v)
	state_addr = libc.Ptr(&new(struct {
		_ [0]uint64
		v int16
		b byte
	}).v)
	result = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	skip = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	eof = libc.Ptr(&new(struct {
		_ [0]uint64
		v byte
		b byte
	}).v)
	lookahead = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	i = libc.Ptr(&new(struct {
		_ [0]uint64
		v int32
		b byte
	}).v)
	*libc.As[unsafe.Pointer](lexer_addr) = lexer
	*libc.As[int16](state_addr) = state
	*libc.As[byte](result) = 0
	*libc.As[byte](skip) = 0
	*libc.As[byte](eof) = 0
	goto start

next_state:
	v0 = *libc.As[unsafe.Pointer](lexer_addr)
	advance = libc.Ptr(&libc.As[TSLexer](v0).F2)
	v1 = *libc.As[unsafe.Pointer](advance)
	v2 = *libc.As[unsafe.Pointer](lexer_addr)
	v3 = *libc.As[byte](skip)
	loadedv = (v3 & 1) != 0
	libc.FuncFromCode[func(unsafe.Pointer, bool)](v1)(v2, loadedv)
	goto start

start:
	*libc.As[byte](skip) = 0
	v4 = *libc.As[unsafe.Pointer](lexer_addr)
	lookahead1 = libc.Ptr(&libc.As[TSLexer](v4).F0)
	v5 = *libc.As[int32](lookahead1)
	*libc.As[int32](lookahead) = v5
	v6 = *libc.As[unsafe.Pointer](lexer_addr)
	eof2 = libc.Ptr(&libc.As[TSLexer](v6).F6)
	v7 = *libc.As[unsafe.Pointer](eof2)
	v8 = *libc.As[unsafe.Pointer](lexer_addr)
	call = libc.FuncFromCode[func(unsafe.Pointer) bool](v7)(v8)
	if call {
		storedv = 1
	} else {
		storedv = 0
	}
	*libc.As[byte](eof) = storedv
	v9 = *libc.As[int16](state_addr)
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
	v10 = *libc.As[byte](eof)
	loadedv3 = (v10 & 1) != 0
	if loadedv3 {
		goto if_then
	} else {
		goto if_end
	}

if_then:
	*libc.As[int16](state_addr) = 66
	goto next_state

if_end:
	*libc.As[int32](i) = 0
	goto for_cond

for_cond:
	v11 = *libc.As[int32](i)
	conv4 = int64(uint64(uint32(v11)))
	cmp = uint64(conv4) < uint64(30)
	if cmp {
		goto for_body
	} else {
		goto for_end
	}

for_body:
	v12 = *libc.As[int32](i)
	idxprom = int64(uint64(uint32(v12)))
	arrayidx = libc.Ptr(&ts_lex_map[idxprom])
	v13 = *libc.As[int16](arrayidx)
	conv6 = int32(uint32(uint16(v13)))
	v14 = *libc.As[int32](lookahead)
	cmp7 = conv6 == v14
	if cmp7 {
		goto if_then9
	} else {
		goto if_end12
	}

if_then9:
	v15 = *libc.As[int32](i)
	add = v15 + 1
	idxprom10 = int64(uint64(uint32(add)))
	arrayidx11 = libc.Ptr(&ts_lex_map[idxprom10])
	v16 = *libc.As[int16](arrayidx11)
	*libc.As[int16](state_addr) = v16
	goto next_state

if_end12:
	goto for_inc

for_inc:
	v17 = *libc.As[int32](i)
	add13 = v17 + 2
	*libc.As[int32](i) = add13
	goto for_cond

for_end:
	v18 = *libc.As[int32](lookahead)
	cmp14 = 49 <= v18
	if cmp14 {
		goto land_lhs_true
	} else {
		goto if_end19
	}

land_lhs_true:
	v19 = *libc.As[int32](lookahead)
	cmp16 = v19 <= 57
	if cmp16 {
		goto if_then18
	} else {
		goto if_end19
	}

if_then18:
	*libc.As[int16](state_addr) = 82
	goto next_state

if_end19:
	v20 = *libc.As[byte](result)
	loadedv20 = (v20 & 1) != 0
	*libc.As[bool](retval) = loadedv20
	goto _return

sw_bb21:
	v21 = *libc.As[int32](lookahead)
	cmp22 = v21 == 45
	if cmp22 {
		goto if_then24
	} else {
		goto if_end25
	}

if_then24:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end25:
	v22 = *libc.As[int32](lookahead)
	cmp26 = v22 == 46
	if cmp26 {
		goto if_then28
	} else {
		goto if_end29
	}

if_then28:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end29:
	v23 = *libc.As[int32](lookahead)
	cmp30 = v23 == 58
	if cmp30 {
		goto if_then32
	} else {
		goto if_end33
	}

if_then32:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end33:
	v24 = *libc.As[int32](lookahead)
	cmp34 = 48 <= v24
	if cmp34 {
		goto land_lhs_true36
	} else {
		goto lor_lhs_false
	}

land_lhs_true36:
	v25 = *libc.As[int32](lookahead)
	cmp37 = v25 <= 57
	if cmp37 {
		goto if_then41
	} else {
		goto lor_lhs_false
	}

lor_lhs_false:
	v26 = *libc.As[int32](lookahead)
	cmp39 = v26 == 95
	if cmp39 {
		goto if_then41
	} else {
		goto if_end42
	}

if_then41:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end42:
	v27 = *libc.As[byte](result)
	loadedv43 = (v27 & 1) != 0
	*libc.As[bool](retval) = loadedv43
	goto _return

sw_bb44:
	v28 = *libc.As[int32](lookahead)
	cmp45 = v28 == 45
	if cmp45 {
		goto if_then47
	} else {
		goto if_end48
	}

if_then47:
	*libc.As[int16](state_addr) = 54
	goto next_state

if_end48:
	v29 = *libc.As[int32](lookahead)
	cmp49 = 48 <= v29
	if cmp49 {
		goto land_lhs_true51
	} else {
		goto if_end55
	}

land_lhs_true51:
	v30 = *libc.As[int32](lookahead)
	cmp52 = v30 <= 57
	if cmp52 {
		goto if_then54
	} else {
		goto if_end55
	}

if_then54:
	*libc.As[int16](state_addr) = 3
	goto next_state

if_end55:
	v31 = *libc.As[byte](result)
	loadedv56 = (v31 & 1) != 0
	*libc.As[bool](retval) = loadedv56
	goto _return

sw_bb57:
	v32 = *libc.As[int32](lookahead)
	cmp58 = v32 == 45
	if cmp58 {
		goto if_then60
	} else {
		goto if_end61
	}

if_then60:
	*libc.As[int16](state_addr) = 60
	goto next_state

if_end61:
	v33 = *libc.As[byte](result)
	loadedv62 = (v33 & 1) != 0
	*libc.As[bool](retval) = loadedv62
	goto _return

sw_bb63:
	v34 = *libc.As[int32](lookahead)
	cmp64 = v34 == 46
	if cmp64 {
		goto if_then66
	} else {
		goto if_end67
	}

if_then66:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end67:
	v35 = *libc.As[int32](lookahead)
	cmp68 = v35 == 58
	if cmp68 {
		goto if_then70
	} else {
		goto if_end71
	}

if_then70:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end71:
	v36 = *libc.As[int32](lookahead)
	cmp72 = v36 == 95
	if cmp72 {
		goto if_then74
	} else {
		goto if_end75
	}

if_then74:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end75:
	v37 = *libc.As[int32](lookahead)
	cmp76 = 48 <= v37
	if cmp76 {
		goto land_lhs_true78
	} else {
		goto if_end82
	}

land_lhs_true78:
	v38 = *libc.As[int32](lookahead)
	cmp79 = v38 <= 57
	if cmp79 {
		goto if_then81
	} else {
		goto if_end82
	}

if_then81:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end82:
	v39 = *libc.As[byte](result)
	loadedv83 = (v39 & 1) != 0
	*libc.As[bool](retval) = loadedv83
	goto _return

sw_bb84:
	v40 = *libc.As[int32](lookahead)
	cmp85 = v40 == 46
	if cmp85 {
		goto if_then87
	} else {
		goto if_end88
	}

if_then87:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end88:
	v41 = *libc.As[int32](lookahead)
	cmp89 = v41 == 58
	if cmp89 {
		goto if_then91
	} else {
		goto if_end92
	}

if_then91:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end92:
	v42 = *libc.As[int32](lookahead)
	cmp93 = v42 == 95
	if cmp93 {
		goto if_then95
	} else {
		goto if_end96
	}

if_then95:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end96:
	v43 = *libc.As[int32](lookahead)
	cmp97 = 48 <= v43
	if cmp97 {
		goto land_lhs_true99
	} else {
		goto if_end103
	}

land_lhs_true99:
	v44 = *libc.As[int32](lookahead)
	cmp100 = v44 <= 57
	if cmp100 {
		goto if_then102
	} else {
		goto if_end103
	}

if_then102:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end103:
	v45 = *libc.As[byte](result)
	loadedv104 = (v45 & 1) != 0
	*libc.As[bool](retval) = loadedv104
	goto _return

sw_bb105:
	v46 = *libc.As[int32](lookahead)
	cmp106 = v46 == 46
	if cmp106 {
		goto if_then108
	} else {
		goto if_end109
	}

if_then108:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end109:
	v47 = *libc.As[int32](lookahead)
	cmp110 = v47 == 58
	if cmp110 {
		goto if_then112
	} else {
		goto if_end113
	}

if_then112:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end113:
	v48 = *libc.As[int32](lookahead)
	cmp114 = 48 <= v48
	if cmp114 {
		goto land_lhs_true116
	} else {
		goto lor_lhs_false119
	}

land_lhs_true116:
	v49 = *libc.As[int32](lookahead)
	cmp117 = v49 <= 57
	if cmp117 {
		goto if_then122
	} else {
		goto lor_lhs_false119
	}

lor_lhs_false119:
	v50 = *libc.As[int32](lookahead)
	cmp120 = v50 == 95
	if cmp120 {
		goto if_then122
	} else {
		goto if_end123
	}

if_then122:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end123:
	v51 = *libc.As[byte](result)
	loadedv124 = (v51 & 1) != 0
	*libc.As[bool](retval) = loadedv124
	goto _return

sw_bb125:
	v52 = *libc.As[int32](lookahead)
	cmp126 = v52 == 46
	if cmp126 {
		goto if_then128
	} else {
		goto if_end129
	}

if_then128:
	*libc.As[int16](state_addr) = 90
	goto next_state

if_end129:
	v53 = *libc.As[int32](lookahead)
	cmp130 = v53 == 48
	if cmp130 {
		goto if_then132
	} else {
		goto if_end133
	}

if_then132:
	*libc.As[int16](state_addr) = 78
	goto next_state

if_end133:
	v54 = *libc.As[int32](lookahead)
	cmp134 = 49 <= v54
	if cmp134 {
		goto land_lhs_true136
	} else {
		goto if_end140
	}

land_lhs_true136:
	v55 = *libc.As[int32](lookahead)
	cmp137 = v55 <= 57
	if cmp137 {
		goto if_then139
	} else {
		goto if_end140
	}

if_then139:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end140:
	v56 = *libc.As[byte](result)
	loadedv141 = (v56 & 1) != 0
	*libc.As[bool](retval) = loadedv141
	goto _return

sw_bb142:
	v57 = *libc.As[int32](lookahead)
	cmp143 = v57 == 46
	if cmp143 {
		goto if_then145
	} else {
		goto if_end146
	}

if_then145:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end146:
	v58 = *libc.As[int32](lookahead)
	cmp147 = v58 == 58
	if cmp147 {
		goto if_then149
	} else {
		goto if_end150
	}

if_then149:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end150:
	v59 = *libc.As[byte](result)
	loadedv151 = (v59 & 1) != 0
	*libc.As[bool](retval) = loadedv151
	goto _return

sw_bb152:
	v60 = *libc.As[int32](lookahead)
	cmp153 = v60 == 46
	if cmp153 {
		goto if_then155
	} else {
		goto if_end156
	}

if_then155:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end156:
	v61 = *libc.As[int32](lookahead)
	cmp157 = v61 == 58
	if cmp157 {
		goto if_then159
	} else {
		goto if_end160
	}

if_then159:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end160:
	v62 = *libc.As[int32](lookahead)
	cmp161 = 48 <= v62
	if cmp161 {
		goto land_lhs_true163
	} else {
		goto if_end167
	}

land_lhs_true163:
	v63 = *libc.As[int32](lookahead)
	cmp164 = v63 <= 57
	if cmp164 {
		goto if_then166
	} else {
		goto if_end167
	}

if_then166:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end167:
	v64 = *libc.As[byte](result)
	loadedv168 = (v64 & 1) != 0
	*libc.As[bool](retval) = loadedv168
	goto _return

sw_bb169:
	v65 = *libc.As[int32](lookahead)
	cmp170 = v65 == 58
	if cmp170 {
		goto if_then172
	} else {
		goto if_end173
	}

if_then172:
	*libc.As[int16](state_addr) = 61
	goto next_state

if_end173:
	v66 = *libc.As[byte](result)
	loadedv174 = (v66 & 1) != 0
	*libc.As[bool](retval) = loadedv174
	goto _return

sw_bb175:
	v67 = *libc.As[int32](lookahead)
	cmp176 = v67 == 58
	if cmp176 {
		goto if_then178
	} else {
		goto if_end179
	}

if_then178:
	*libc.As[int16](state_addr) = 61
	goto next_state

if_end179:
	v68 = *libc.As[int32](lookahead)
	cmp180 = 48 <= v68
	if cmp180 {
		goto land_lhs_true182
	} else {
		goto if_end186
	}

land_lhs_true182:
	v69 = *libc.As[int32](lookahead)
	cmp183 = v69 <= 57
	if cmp183 {
		goto if_then185
	} else {
		goto if_end186
	}

if_then185:
	*libc.As[int16](state_addr) = 10
	goto next_state

if_end186:
	v70 = *libc.As[byte](result)
	loadedv187 = (v70 & 1) != 0
	*libc.As[bool](retval) = loadedv187
	goto _return

sw_bb188:
	v71 = *libc.As[int32](lookahead)
	cmp189 = v71 == 58
	if cmp189 {
		goto if_then191
	} else {
		goto if_end192
	}

if_then191:
	*libc.As[int16](state_addr) = 62
	goto next_state

if_end192:
	v72 = *libc.As[byte](result)
	loadedv193 = (v72 & 1) != 0
	*libc.As[bool](retval) = loadedv193
	goto _return

sw_bb194:
	v73 = *libc.As[int32](lookahead)
	cmp195 = v73 == 65
	if cmp195 {
		goto if_then197
	} else {
		goto if_end198
	}

if_then197:
	*libc.As[int16](state_addr) = 20
	goto next_state

if_end198:
	v74 = *libc.As[int32](lookahead)
	cmp199 = v74 == 97
	if cmp199 {
		goto if_then201
	} else {
		goto if_end202
	}

if_then201:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end202:
	v75 = *libc.As[byte](result)
	loadedv203 = (v75 & 1) != 0
	*libc.As[bool](retval) = loadedv203
	goto _return

sw_bb204:
	v76 = *libc.As[int32](lookahead)
	cmp205 = v76 == 65
	if cmp205 {
		goto if_then207
	} else {
		goto if_end208
	}

if_then207:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end208:
	v77 = *libc.As[int32](lookahead)
	cmp209 = v77 == 97
	if cmp209 {
		goto if_then211
	} else {
		goto if_end212
	}

if_then211:
	*libc.As[int16](state_addr) = 22
	goto next_state

if_end212:
	v78 = *libc.As[byte](result)
	loadedv213 = (v78 & 1) != 0
	*libc.As[bool](retval) = loadedv213
	goto _return

sw_bb214:
	v79 = *libc.As[int32](lookahead)
	cmp215 = v79 == 69
	if cmp215 {
		goto if_then217
	} else {
		goto if_end218
	}

if_then217:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end218:
	v80 = *libc.As[byte](result)
	loadedv219 = (v80 & 1) != 0
	*libc.As[bool](retval) = loadedv219
	goto _return

sw_bb220:
	v81 = *libc.As[int32](lookahead)
	cmp221 = v81 == 70
	if cmp221 {
		goto if_then223
	} else {
		goto if_end224
	}

if_then223:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end224:
	v82 = *libc.As[byte](result)
	loadedv225 = (v82 & 1) != 0
	*libc.As[bool](retval) = loadedv225
	goto _return

sw_bb226:
	v83 = *libc.As[int32](lookahead)
	cmp227 = v83 == 70
	if cmp227 {
		goto if_then229
	} else {
		goto if_end230
	}

if_then229:
	*libc.As[int16](state_addr) = 16
	goto next_state

if_end230:
	v84 = *libc.As[int32](lookahead)
	cmp231 = v84 == 102
	if cmp231 {
		goto if_then233
	} else {
		goto if_end234
	}

if_then233:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end234:
	v85 = *libc.As[int32](lookahead)
	cmp235 = v85 == 78
	if cmp235 {
		goto if_then240
	} else {
		goto lor_lhs_false237
	}

lor_lhs_false237:
	v86 = *libc.As[int32](lookahead)
	cmp238 = v86 == 110
	if cmp238 {
		goto if_then240
	} else {
		goto if_end241
	}

if_then240:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end241:
	v87 = *libc.As[byte](result)
	loadedv242 = (v87 & 1) != 0
	*libc.As[bool](retval) = loadedv242
	goto _return

sw_bb243:
	v88 = *libc.As[int32](lookahead)
	cmp244 = v88 == 70
	if cmp244 {
		goto if_then246
	} else {
		goto if_end247
	}

if_then246:
	*libc.As[int16](state_addr) = 88
	goto next_state

if_end247:
	v89 = *libc.As[byte](result)
	loadedv248 = (v89 & 1) != 0
	*libc.As[bool](retval) = loadedv248
	goto _return

sw_bb249:
	v90 = *libc.As[int32](lookahead)
	cmp250 = v90 == 76
	if cmp250 {
		goto if_then252
	} else {
		goto if_end253
	}

if_then252:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end253:
	v91 = *libc.As[byte](result)
	loadedv254 = (v91 & 1) != 0
	*libc.As[bool](retval) = loadedv254
	goto _return

sw_bb255:
	v92 = *libc.As[int32](lookahead)
	cmp256 = v92 == 76
	if cmp256 {
		goto if_then258
	} else {
		goto if_end259
	}

if_then258:
	*libc.As[int16](state_addr) = 26
	goto next_state

if_end259:
	v93 = *libc.As[byte](result)
	loadedv260 = (v93 & 1) != 0
	*libc.As[bool](retval) = loadedv260
	goto _return

sw_bb261:
	v94 = *libc.As[int32](lookahead)
	cmp262 = v94 == 76
	if cmp262 {
		goto if_then264
	} else {
		goto if_end265
	}

if_then264:
	*libc.As[int16](state_addr) = 19
	goto next_state

if_end265:
	v95 = *libc.As[byte](result)
	loadedv266 = (v95 & 1) != 0
	*libc.As[bool](retval) = loadedv266
	goto _return

sw_bb267:
	v96 = *libc.As[int32](lookahead)
	cmp268 = v96 == 78
	if cmp268 {
		goto if_then270
	} else {
		goto if_end271
	}

if_then270:
	*libc.As[int16](state_addr) = 88
	goto next_state

if_end271:
	v97 = *libc.As[byte](result)
	loadedv272 = (v97 & 1) != 0
	*libc.As[bool](retval) = loadedv272
	goto _return

sw_bb273:
	v98 = *libc.As[int32](lookahead)
	cmp274 = v98 == 78
	if cmp274 {
		goto if_then276
	} else {
		goto if_end277
	}

if_then276:
	*libc.As[int16](state_addr) = 18
	goto next_state

if_end277:
	v99 = *libc.As[int32](lookahead)
	cmp278 = v99 == 110
	if cmp278 {
		goto if_then280
	} else {
		goto if_end281
	}

if_then280:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end281:
	v100 = *libc.As[byte](result)
	loadedv282 = (v100 & 1) != 0
	*libc.As[bool](retval) = loadedv282
	goto _return

sw_bb283:
	v101 = *libc.As[int32](lookahead)
	cmp284 = v101 == 82
	if cmp284 {
		goto if_then286
	} else {
		goto if_end287
	}

if_then286:
	*libc.As[int16](state_addr) = 27
	goto next_state

if_end287:
	v102 = *libc.As[int32](lookahead)
	cmp288 = v102 == 114
	if cmp288 {
		goto if_then290
	} else {
		goto if_end291
	}

if_then290:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end291:
	v103 = *libc.As[byte](result)
	loadedv292 = (v103 & 1) != 0
	*libc.As[bool](retval) = loadedv292
	goto _return

sw_bb293:
	v104 = *libc.As[int32](lookahead)
	cmp294 = v104 == 83
	if cmp294 {
		goto if_then296
	} else {
		goto if_end297
	}

if_then296:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end297:
	v105 = *libc.As[byte](result)
	loadedv298 = (v105 & 1) != 0
	*libc.As[bool](retval) = loadedv298
	goto _return

sw_bb299:
	v106 = *libc.As[int32](lookahead)
	cmp300 = v106 == 83
	if cmp300 {
		goto if_then302
	} else {
		goto if_end303
	}

if_then302:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end303:
	v107 = *libc.As[byte](result)
	loadedv304 = (v107 & 1) != 0
	*libc.As[bool](retval) = loadedv304
	goto _return

sw_bb305:
	v108 = *libc.As[int32](lookahead)
	cmp306 = v108 == 85
	if cmp306 {
		goto if_then308
	} else {
		goto if_end309
	}

if_then308:
	*libc.As[int16](state_addr) = 15
	goto next_state

if_end309:
	v109 = *libc.As[byte](result)
	loadedv310 = (v109 & 1) != 0
	*libc.As[bool](retval) = loadedv310
	goto _return

sw_bb311:
	v110 = *libc.As[int32](lookahead)
	cmp312 = v110 == 90
	if cmp312 {
		goto if_then314
	} else {
		goto if_end315
	}

if_then314:
	*libc.As[int16](state_addr) = 94
	goto next_state

if_end315:
	v111 = *libc.As[int32](lookahead)
	cmp316 = v111 == 9
	if cmp316 {
		goto if_then321
	} else {
		goto lor_lhs_false318
	}

lor_lhs_false318:
	v112 = *libc.As[int32](lookahead)
	cmp319 = v112 == 32
	if cmp319 {
		goto if_then321
	} else {
		goto if_end322
	}

if_then321:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end322:
	v113 = *libc.As[byte](result)
	loadedv323 = (v113 & 1) != 0
	*libc.As[bool](retval) = loadedv323
	goto _return

sw_bb324:
	v114 = *libc.As[int32](lookahead)
	cmp325 = v114 == 97
	if cmp325 {
		goto if_then327
	} else {
		goto if_end328
	}

if_then327:
	*libc.As[int16](state_addr) = 36
	goto next_state

if_end328:
	v115 = *libc.As[byte](result)
	loadedv329 = (v115 & 1) != 0
	*libc.As[bool](retval) = loadedv329
	goto _return

sw_bb330:
	v116 = *libc.As[int32](lookahead)
	cmp331 = v116 == 97
	if cmp331 {
		goto if_then333
	} else {
		goto if_end334
	}

if_then333:
	*libc.As[int16](state_addr) = 38
	goto next_state

if_end334:
	v117 = *libc.As[byte](result)
	loadedv335 = (v117 & 1) != 0
	*libc.As[bool](retval) = loadedv335
	goto _return

sw_bb336:
	v118 = *libc.As[int32](lookahead)
	cmp337 = v118 == 101
	if cmp337 {
		goto if_then339
	} else {
		goto if_end340
	}

if_then339:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end340:
	v119 = *libc.As[byte](result)
	loadedv341 = (v119 & 1) != 0
	*libc.As[bool](retval) = loadedv341
	goto _return

sw_bb342:
	v120 = *libc.As[int32](lookahead)
	cmp343 = v120 == 102
	if cmp343 {
		goto if_then345
	} else {
		goto if_end346
	}

if_then345:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end346:
	v121 = *libc.As[byte](result)
	loadedv347 = (v121 & 1) != 0
	*libc.As[bool](retval) = loadedv347
	goto _return

sw_bb348:
	v122 = *libc.As[int32](lookahead)
	cmp349 = v122 == 102
	if cmp349 {
		goto if_then351
	} else {
		goto if_end352
	}

if_then351:
	*libc.As[int16](state_addr) = 32
	goto next_state

if_end352:
	v123 = *libc.As[int32](lookahead)
	cmp353 = v123 == 110
	if cmp353 {
		goto if_then355
	} else {
		goto if_end356
	}

if_then355:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end356:
	v124 = *libc.As[byte](result)
	loadedv357 = (v124 & 1) != 0
	*libc.As[bool](retval) = loadedv357
	goto _return

sw_bb358:
	v125 = *libc.As[int32](lookahead)
	cmp359 = v125 == 102
	if cmp359 {
		goto if_then361
	} else {
		goto if_end362
	}

if_then361:
	*libc.As[int16](state_addr) = 88
	goto next_state

if_end362:
	v126 = *libc.As[byte](result)
	loadedv363 = (v126 & 1) != 0
	*libc.As[bool](retval) = loadedv363
	goto _return

sw_bb364:
	v127 = *libc.As[int32](lookahead)
	cmp365 = v127 == 108
	if cmp365 {
		goto if_then367
	} else {
		goto if_end368
	}

if_then367:
	*libc.As[int16](state_addr) = 67
	goto next_state

if_end368:
	v128 = *libc.As[byte](result)
	loadedv369 = (v128 & 1) != 0
	*libc.As[bool](retval) = loadedv369
	goto _return

sw_bb370:
	v129 = *libc.As[int32](lookahead)
	cmp371 = v129 == 108
	if cmp371 {
		goto if_then373
	} else {
		goto if_end374
	}

if_then373:
	*libc.As[int16](state_addr) = 42
	goto next_state

if_end374:
	v130 = *libc.As[byte](result)
	loadedv375 = (v130 & 1) != 0
	*libc.As[bool](retval) = loadedv375
	goto _return

sw_bb376:
	v131 = *libc.As[int32](lookahead)
	cmp377 = v131 == 108
	if cmp377 {
		goto if_then379
	} else {
		goto if_end380
	}

if_then379:
	*libc.As[int16](state_addr) = 35
	goto next_state

if_end380:
	v132 = *libc.As[byte](result)
	loadedv381 = (v132 & 1) != 0
	*libc.As[bool](retval) = loadedv381
	goto _return

sw_bb382:
	v133 = *libc.As[int32](lookahead)
	cmp383 = v133 == 110
	if cmp383 {
		goto if_then385
	} else {
		goto if_end386
	}

if_then385:
	*libc.As[int16](state_addr) = 88
	goto next_state

if_end386:
	v134 = *libc.As[byte](result)
	loadedv387 = (v134 & 1) != 0
	*libc.As[bool](retval) = loadedv387
	goto _return

sw_bb388:
	v135 = *libc.As[int32](lookahead)
	cmp389 = v135 == 110
	if cmp389 {
		goto if_then391
	} else {
		goto if_end392
	}

if_then391:
	*libc.As[int16](state_addr) = 34
	goto next_state

if_end392:
	v136 = *libc.As[byte](result)
	loadedv393 = (v136 & 1) != 0
	*libc.As[bool](retval) = loadedv393
	goto _return

sw_bb394:
	v137 = *libc.As[int32](lookahead)
	cmp395 = v137 == 114
	if cmp395 {
		goto if_then397
	} else {
		goto if_end398
	}

if_then397:
	*libc.As[int16](state_addr) = 43
	goto next_state

if_end398:
	v138 = *libc.As[byte](result)
	loadedv399 = (v138 & 1) != 0
	*libc.As[bool](retval) = loadedv399
	goto _return

sw_bb400:
	v139 = *libc.As[int32](lookahead)
	cmp401 = v139 == 115
	if cmp401 {
		goto if_then403
	} else {
		goto if_end404
	}

if_then403:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end404:
	v140 = *libc.As[byte](result)
	loadedv405 = (v140 & 1) != 0
	*libc.As[bool](retval) = loadedv405
	goto _return

sw_bb406:
	v141 = *libc.As[int32](lookahead)
	cmp407 = v141 == 115
	if cmp407 {
		goto if_then409
	} else {
		goto if_end410
	}

if_then409:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end410:
	v142 = *libc.As[byte](result)
	loadedv411 = (v142 & 1) != 0
	*libc.As[bool](retval) = loadedv411
	goto _return

sw_bb412:
	v143 = *libc.As[int32](lookahead)
	cmp413 = v143 == 117
	if cmp413 {
		goto if_then415
	} else {
		goto if_end416
	}

if_then415:
	*libc.As[int16](state_addr) = 31
	goto next_state

if_end416:
	v144 = *libc.As[byte](result)
	loadedv417 = (v144 & 1) != 0
	*libc.As[bool](retval) = loadedv417
	goto _return

sw_bb418:
	v145 = *libc.As[int32](lookahead)
	cmp419 = v145 == 9
	if cmp419 {
		goto if_then424
	} else {
		goto lor_lhs_false421
	}

lor_lhs_false421:
	v146 = *libc.As[int32](lookahead)
	cmp422 = v146 == 32
	if cmp422 {
		goto if_then424
	} else {
		goto if_end425
	}

if_then424:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end425:
	v147 = *libc.As[int32](lookahead)
	cmp426 = v147 == 84
	if cmp426 {
		goto if_then431
	} else {
		goto lor_lhs_false428
	}

lor_lhs_false428:
	v148 = *libc.As[int32](lookahead)
	cmp429 = v148 == 116
	if cmp429 {
		goto if_then431
	} else {
		goto if_end432
	}

if_then431:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end432:
	v149 = *libc.As[byte](result)
	loadedv433 = (v149 & 1) != 0
	*libc.As[bool](retval) = loadedv433
	goto _return

sw_bb434:
	v150 = *libc.As[int32](lookahead)
	cmp435 = v150 == 9
	if cmp435 {
		goto if_then440
	} else {
		goto lor_lhs_false437
	}

lor_lhs_false437:
	v151 = *libc.As[int32](lookahead)
	cmp438 = v151 == 32
	if cmp438 {
		goto if_then440
	} else {
		goto if_end441
	}

if_then440:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end441:
	v152 = *libc.As[int32](lookahead)
	cmp442 = v152 == 84
	if cmp442 {
		goto if_then447
	} else {
		goto lor_lhs_false444
	}

lor_lhs_false444:
	v153 = *libc.As[int32](lookahead)
	cmp445 = v153 == 116
	if cmp445 {
		goto if_then447
	} else {
		goto if_end448
	}

if_then447:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end448:
	v154 = *libc.As[int32](lookahead)
	cmp449 = 48 <= v154
	if cmp449 {
		goto land_lhs_true451
	} else {
		goto if_end455
	}

land_lhs_true451:
	v155 = *libc.As[int32](lookahead)
	cmp452 = v155 <= 57
	if cmp452 {
		goto if_then454
	} else {
		goto if_end455
	}

if_then454:
	*libc.As[int16](state_addr) = 44
	goto next_state

if_end455:
	v156 = *libc.As[byte](result)
	loadedv456 = (v156 & 1) != 0
	*libc.As[bool](retval) = loadedv456
	goto _return

sw_bb457:
	v157 = *libc.As[int32](lookahead)
	cmp458 = v157 == 9
	if cmp458 {
		goto if_then463
	} else {
		goto lor_lhs_false460
	}

lor_lhs_false460:
	v158 = *libc.As[int32](lookahead)
	cmp461 = v158 == 32
	if cmp461 {
		goto if_then463
	} else {
		goto if_end464
	}

if_then463:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end464:
	v159 = *libc.As[int32](lookahead)
	cmp465 = v159 == 84
	if cmp465 {
		goto if_then470
	} else {
		goto lor_lhs_false467
	}

lor_lhs_false467:
	v160 = *libc.As[int32](lookahead)
	cmp468 = v160 == 116
	if cmp468 {
		goto if_then470
	} else {
		goto if_end471
	}

if_then470:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end471:
	v161 = *libc.As[int32](lookahead)
	cmp472 = 48 <= v161
	if cmp472 {
		goto land_lhs_true474
	} else {
		goto if_end478
	}

land_lhs_true474:
	v162 = *libc.As[int32](lookahead)
	cmp475 = v162 <= 57
	if cmp475 {
		goto if_then477
	} else {
		goto if_end478
	}

if_then477:
	*libc.As[int16](state_addr) = 99
	goto next_state

if_end478:
	v163 = *libc.As[byte](result)
	loadedv479 = (v163 & 1) != 0
	*libc.As[bool](retval) = loadedv479
	goto _return

sw_bb480:
	v164 = *libc.As[int32](lookahead)
	cmp481 = v164 == 9
	if cmp481 {
		goto if_then486
	} else {
		goto lor_lhs_false483
	}

lor_lhs_false483:
	v165 = *libc.As[int32](lookahead)
	cmp484 = v165 == 32
	if cmp484 {
		goto if_then486
	} else {
		goto if_end487
	}

if_then486:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end487:
	v166 = *libc.As[int32](lookahead)
	cmp488 = 48 <= v166
	if cmp488 {
		goto land_lhs_true490
	} else {
		goto if_end494
	}

land_lhs_true490:
	v167 = *libc.As[int32](lookahead)
	cmp491 = v167 <= 57
	if cmp491 {
		goto if_then493
	} else {
		goto if_end494
	}

if_then493:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end494:
	v168 = *libc.As[byte](result)
	loadedv495 = (v168 & 1) != 0
	*libc.As[bool](retval) = loadedv495
	goto _return

sw_bb496:
	v169 = *libc.As[int32](lookahead)
	cmp497 = v169 == 43
	if cmp497 {
		goto if_then502
	} else {
		goto lor_lhs_false499
	}

lor_lhs_false499:
	v170 = *libc.As[int32](lookahead)
	cmp500 = v170 == 45
	if cmp500 {
		goto if_then502
	} else {
		goto if_end503
	}

if_then502:
	*libc.As[int16](state_addr) = 52
	goto next_state

if_end503:
	v171 = *libc.As[byte](result)
	loadedv504 = (v171 & 1) != 0
	*libc.As[bool](retval) = loadedv504
	goto _return

sw_bb505:
	v172 = *libc.As[int32](lookahead)
	cmp506 = v172 == 48
	if cmp506 {
		goto if_then514
	} else {
		goto lor_lhs_false508
	}

lor_lhs_false508:
	v173 = *libc.As[int32](lookahead)
	cmp509 = v173 == 49
	if cmp509 {
		goto if_then514
	} else {
		goto lor_lhs_false511
	}

lor_lhs_false511:
	v174 = *libc.As[int32](lookahead)
	cmp512 = v174 == 95
	if cmp512 {
		goto if_then514
	} else {
		goto if_end515
	}

if_then514:
	*libc.As[int16](state_addr) = 86
	goto next_state

if_end515:
	v175 = *libc.As[byte](result)
	loadedv516 = (v175 & 1) != 0
	*libc.As[bool](retval) = loadedv516
	goto _return

sw_bb517:
	v176 = *libc.As[int32](lookahead)
	cmp518 = 54 <= v176
	if cmp518 {
		goto land_lhs_true520
	} else {
		goto if_end524
	}

land_lhs_true520:
	v177 = *libc.As[int32](lookahead)
	cmp521 = v177 <= 57
	if cmp521 {
		goto if_then523
	} else {
		goto if_end524
	}

if_then523:
	*libc.As[int16](state_addr) = 8
	goto next_state

if_end524:
	v178 = *libc.As[int32](lookahead)
	cmp525 = 48 <= v178
	if cmp525 {
		goto land_lhs_true527
	} else {
		goto if_end531
	}

land_lhs_true527:
	v179 = *libc.As[int32](lookahead)
	cmp528 = v179 <= 53
	if cmp528 {
		goto if_then530
	} else {
		goto if_end531
	}

if_then530:
	*libc.As[int16](state_addr) = 9
	goto next_state

if_end531:
	v180 = *libc.As[byte](result)
	loadedv532 = (v180 & 1) != 0
	*libc.As[bool](retval) = loadedv532
	goto _return

sw_bb533:
	v181 = *libc.As[int32](lookahead)
	cmp534 = 54 <= v181
	if cmp534 {
		goto land_lhs_true536
	} else {
		goto if_end540
	}

land_lhs_true536:
	v182 = *libc.As[int32](lookahead)
	cmp537 = v182 <= 57
	if cmp537 {
		goto if_then539
	} else {
		goto if_end540
	}

if_then539:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end540:
	v183 = *libc.As[int32](lookahead)
	cmp541 = 48 <= v183
	if cmp541 {
		goto land_lhs_true543
	} else {
		goto if_end547
	}

land_lhs_true543:
	v184 = *libc.As[int32](lookahead)
	cmp544 = v184 <= 53
	if cmp544 {
		goto if_then546
	} else {
		goto if_end547
	}

if_then546:
	*libc.As[int16](state_addr) = 85
	goto next_state

if_end547:
	v185 = *libc.As[byte](result)
	loadedv548 = (v185 & 1) != 0
	*libc.As[bool](retval) = loadedv548
	goto _return

sw_bb549:
	v186 = *libc.As[int32](lookahead)
	cmp550 = 48 <= v186
	if cmp550 {
		goto land_lhs_true552
	} else {
		goto if_end556
	}

land_lhs_true552:
	v187 = *libc.As[int32](lookahead)
	cmp553 = v187 <= 57
	if cmp553 {
		goto if_then555
	} else {
		goto if_end556
	}

if_then555:
	*libc.As[int16](state_addr) = 92
	goto next_state

if_end556:
	v188 = *libc.As[byte](result)
	loadedv557 = (v188 & 1) != 0
	*libc.As[bool](retval) = loadedv557
	goto _return

sw_bb558:
	v189 = *libc.As[int32](lookahead)
	cmp559 = 48 <= v189
	if cmp559 {
		goto land_lhs_true561
	} else {
		goto if_end565
	}

land_lhs_true561:
	v190 = *libc.As[int32](lookahead)
	cmp562 = v190 <= 57
	if cmp562 {
		goto if_then564
	} else {
		goto if_end565
	}

if_then564:
	*libc.As[int16](state_addr) = 2
	goto next_state

if_end565:
	v191 = *libc.As[byte](result)
	loadedv566 = (v191 & 1) != 0
	*libc.As[bool](retval) = loadedv566
	goto _return

sw_bb567:
	v192 = *libc.As[int32](lookahead)
	cmp568 = 48 <= v192
	if cmp568 {
		goto land_lhs_true570
	} else {
		goto if_end574
	}

land_lhs_true570:
	v193 = *libc.As[int32](lookahead)
	cmp571 = v193 <= 57
	if cmp571 {
		goto if_then573
	} else {
		goto if_end574
	}

if_then573:
	*libc.As[int16](state_addr) = 45
	goto next_state

if_end574:
	v194 = *libc.As[byte](result)
	loadedv575 = (v194 & 1) != 0
	*libc.As[bool](retval) = loadedv575
	goto _return

sw_bb576:
	v195 = *libc.As[int32](lookahead)
	cmp577 = 48 <= v195
	if cmp577 {
		goto land_lhs_true579
	} else {
		goto if_end583
	}

land_lhs_true579:
	v196 = *libc.As[int32](lookahead)
	cmp580 = v196 <= 57
	if cmp580 {
		goto if_then582
	} else {
		goto if_end583
	}

if_then582:
	*libc.As[int16](state_addr) = 11
	goto next_state

if_end583:
	v197 = *libc.As[byte](result)
	loadedv584 = (v197 & 1) != 0
	*libc.As[bool](retval) = loadedv584
	goto _return

sw_bb585:
	v198 = *libc.As[int32](lookahead)
	cmp586 = 48 <= v198
	if cmp586 {
		goto land_lhs_true588
	} else {
		goto if_end592
	}

land_lhs_true588:
	v199 = *libc.As[int32](lookahead)
	cmp589 = v199 <= 57
	if cmp589 {
		goto if_then591
	} else {
		goto if_end592
	}

if_then591:
	*libc.As[int16](state_addr) = 95
	goto next_state

if_end592:
	v200 = *libc.As[byte](result)
	loadedv593 = (v200 & 1) != 0
	*libc.As[bool](retval) = loadedv593
	goto _return

sw_bb594:
	v201 = *libc.As[int32](lookahead)
	cmp595 = 48 <= v201
	if cmp595 {
		goto land_lhs_true597
	} else {
		goto if_end601
	}

land_lhs_true597:
	v202 = *libc.As[int32](lookahead)
	cmp598 = v202 <= 57
	if cmp598 {
		goto if_then600
	} else {
		goto if_end601
	}

if_then600:
	*libc.As[int16](state_addr) = 94
	goto next_state

if_end601:
	v203 = *libc.As[byte](result)
	loadedv602 = (v203 & 1) != 0
	*libc.As[bool](retval) = loadedv602
	goto _return

sw_bb603:
	v204 = *libc.As[int32](lookahead)
	cmp604 = 48 <= v204
	if cmp604 {
		goto land_lhs_true606
	} else {
		goto if_end610
	}

land_lhs_true606:
	v205 = *libc.As[int32](lookahead)
	cmp607 = v205 <= 57
	if cmp607 {
		goto if_then609
	} else {
		goto if_end610
	}

if_then609:
	*libc.As[int16](state_addr) = 98
	goto next_state

if_end610:
	v206 = *libc.As[byte](result)
	loadedv611 = (v206 & 1) != 0
	*libc.As[bool](retval) = loadedv611
	goto _return

sw_bb612:
	v207 = *libc.As[int32](lookahead)
	cmp613 = 48 <= v207
	if cmp613 {
		goto land_lhs_true615
	} else {
		goto if_end619
	}

land_lhs_true615:
	v208 = *libc.As[int32](lookahead)
	cmp616 = v208 <= 57
	if cmp616 {
		goto if_then618
	} else {
		goto if_end619
	}

if_then618:
	*libc.As[int16](state_addr) = 97
	goto next_state

if_end619:
	v209 = *libc.As[byte](result)
	loadedv620 = (v209 & 1) != 0
	*libc.As[bool](retval) = loadedv620
	goto _return

sw_bb621:
	v210 = *libc.As[int32](lookahead)
	cmp622 = 48 <= v210
	if cmp622 {
		goto land_lhs_true624
	} else {
		goto if_end628
	}

land_lhs_true624:
	v211 = *libc.As[int32](lookahead)
	cmp625 = v211 <= 57
	if cmp625 {
		goto if_then627
	} else {
		goto if_end628
	}

if_then627:
	*libc.As[int16](state_addr) = 46
	goto next_state

if_end628:
	v212 = *libc.As[byte](result)
	loadedv629 = (v212 & 1) != 0
	*libc.As[bool](retval) = loadedv629
	goto _return

sw_bb630:
	v213 = *libc.As[int32](lookahead)
	cmp631 = 48 <= v213
	if cmp631 {
		goto land_lhs_true633
	} else {
		goto if_end637
	}

land_lhs_true633:
	v214 = *libc.As[int32](lookahead)
	cmp634 = v214 <= 57
	if cmp634 {
		goto if_then636
	} else {
		goto if_end637
	}

if_then636:
	*libc.As[int16](state_addr) = 64
	goto next_state

if_end637:
	v215 = *libc.As[byte](result)
	loadedv638 = (v215 & 1) != 0
	*libc.As[bool](retval) = loadedv638
	goto _return

sw_bb639:
	v216 = *libc.As[int32](lookahead)
	cmp640 = 48 <= v216
	if cmp640 {
		goto land_lhs_true642
	} else {
		goto if_end646
	}

land_lhs_true642:
	v217 = *libc.As[int32](lookahead)
	cmp643 = v217 <= 57
	if cmp643 {
		goto if_then645
	} else {
		goto if_end646
	}

if_then645:
	*libc.As[int16](state_addr) = 56
	goto next_state

if_end646:
	v218 = *libc.As[byte](result)
	loadedv647 = (v218 & 1) != 0
	*libc.As[bool](retval) = loadedv647
	goto _return

sw_bb648:
	v219 = *libc.As[int32](lookahead)
	cmp649 = 48 <= v219
	if cmp649 {
		goto land_lhs_true651
	} else {
		goto if_end655
	}

land_lhs_true651:
	v220 = *libc.As[int32](lookahead)
	cmp652 = v220 <= 57
	if cmp652 {
		goto if_then654
	} else {
		goto if_end655
	}

if_then654:
	*libc.As[int16](state_addr) = 57
	goto next_state

if_end655:
	v221 = *libc.As[byte](result)
	loadedv656 = (v221 & 1) != 0
	*libc.As[bool](retval) = loadedv656
	goto _return

sw_bb657:
	v222 = *libc.As[int32](lookahead)
	cmp658 = 48 <= v222
	if cmp658 {
		goto land_lhs_true660
	} else {
		goto if_end664
	}

land_lhs_true660:
	v223 = *libc.As[int32](lookahead)
	cmp661 = v223 <= 57
	if cmp661 {
		goto if_then663
	} else {
		goto if_end664
	}

if_then663:
	*libc.As[int16](state_addr) = 12
	goto next_state

if_end664:
	v224 = *libc.As[byte](result)
	loadedv665 = (v224 & 1) != 0
	*libc.As[bool](retval) = loadedv665
	goto _return

sw_bb666:
	v225 = *libc.As[int32](lookahead)
	cmp667 = 48 <= v225
	if cmp667 {
		goto land_lhs_true669
	} else {
		goto lor_lhs_false672
	}

land_lhs_true669:
	v226 = *libc.As[int32](lookahead)
	cmp670 = v226 <= 57
	if cmp670 {
		goto if_then687
	} else {
		goto lor_lhs_false672
	}

lor_lhs_false672:
	v227 = *libc.As[int32](lookahead)
	cmp673 = 65 <= v227
	if cmp673 {
		goto land_lhs_true675
	} else {
		goto lor_lhs_false678
	}

land_lhs_true675:
	v228 = *libc.As[int32](lookahead)
	cmp676 = v228 <= 70
	if cmp676 {
		goto if_then687
	} else {
		goto lor_lhs_false678
	}

lor_lhs_false678:
	v229 = *libc.As[int32](lookahead)
	cmp679 = v229 == 95
	if cmp679 {
		goto if_then687
	} else {
		goto lor_lhs_false681
	}

lor_lhs_false681:
	v230 = *libc.As[int32](lookahead)
	cmp682 = 97 <= v230
	if cmp682 {
		goto land_lhs_true684
	} else {
		goto if_end688
	}

land_lhs_true684:
	v231 = *libc.As[int32](lookahead)
	cmp685 = v231 <= 102
	if cmp685 {
		goto if_then687
	} else {
		goto if_end688
	}

if_then687:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end688:
	v232 = *libc.As[byte](result)
	loadedv689 = (v232 & 1) != 0
	*libc.As[bool](retval) = loadedv689
	goto _return

sw_bb690:
	*libc.As[byte](result) = 1
	v233 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol = libc.Ptr(&libc.As[TSLexer](v233).F1)
	*libc.As[int16](result_symbol) = 0
	v234 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end = libc.Ptr(&libc.As[TSLexer](v234).F3)
	v235 = *libc.As[unsafe.Pointer](mark_end)
	v236 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v235)(v236)
	v237 = *libc.As[byte](result)
	loadedv691 = (v237 & 1) != 0
	*libc.As[bool](retval) = loadedv691
	goto _return

sw_bb692:
	*libc.As[byte](result) = 1
	v238 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol693 = libc.Ptr(&libc.As[TSLexer](v238).F1)
	*libc.As[int16](result_symbol693) = 1
	v239 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end694 = libc.Ptr(&libc.As[TSLexer](v239).F3)
	v240 = *libc.As[unsafe.Pointer](mark_end694)
	v241 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v240)(v241)
	v242 = *libc.As[byte](result)
	loadedv695 = (v242 & 1) != 0
	*libc.As[bool](retval) = loadedv695
	goto _return

sw_bb696:
	*libc.As[byte](result) = 1
	v243 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol697 = libc.Ptr(&libc.As[TSLexer](v243).F1)
	*libc.As[int16](result_symbol697) = 2
	v244 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end698 = libc.Ptr(&libc.As[TSLexer](v244).F3)
	v245 = *libc.As[unsafe.Pointer](mark_end698)
	v246 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v245)(v246)
	v247 = *libc.As[byte](result)
	loadedv699 = (v247 & 1) != 0
	*libc.As[bool](retval) = loadedv699
	goto _return

sw_bb700:
	*libc.As[byte](result) = 1
	v248 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol701 = libc.Ptr(&libc.As[TSLexer](v248).F1)
	*libc.As[int16](result_symbol701) = 2
	v249 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end702 = libc.Ptr(&libc.As[TSLexer](v249).F3)
	v250 = *libc.As[unsafe.Pointer](mark_end702)
	v251 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v250)(v251)
	v252 = *libc.As[int32](lookahead)
	cmp703 = v252 == 69
	if cmp703 {
		goto if_then705
	} else {
		goto if_end706
	}

if_then705:
	*libc.As[int16](state_addr) = 25
	goto next_state

if_end706:
	v253 = *libc.As[int32](lookahead)
	cmp707 = v253 == 101
	if cmp707 {
		goto if_then709
	} else {
		goto if_end710
	}

if_then709:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end710:
	v254 = *libc.As[byte](result)
	loadedv711 = (v254 & 1) != 0
	*libc.As[bool](retval) = loadedv711
	goto _return

sw_bb712:
	*libc.As[byte](result) = 1
	v255 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol713 = libc.Ptr(&libc.As[TSLexer](v255).F1)
	*libc.As[int16](result_symbol713) = 2
	v256 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end714 = libc.Ptr(&libc.As[TSLexer](v256).F3)
	v257 = *libc.As[unsafe.Pointer](mark_end714)
	v258 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v257)(v258)
	v259 = *libc.As[int32](lookahead)
	cmp715 = v259 == 85
	if cmp715 {
		goto if_then717
	} else {
		goto if_end718
	}

if_then717:
	*libc.As[int16](state_addr) = 21
	goto next_state

if_end718:
	v260 = *libc.As[int32](lookahead)
	cmp719 = v260 == 117
	if cmp719 {
		goto if_then721
	} else {
		goto if_end722
	}

if_then721:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end722:
	v261 = *libc.As[int32](lookahead)
	cmp723 = v261 == 79
	if cmp723 {
		goto if_then728
	} else {
		goto lor_lhs_false725
	}

lor_lhs_false725:
	v262 = *libc.As[int32](lookahead)
	cmp726 = v262 == 111
	if cmp726 {
		goto if_then728
	} else {
		goto if_end729
	}

if_then728:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end729:
	v263 = *libc.As[byte](result)
	loadedv730 = (v263 & 1) != 0
	*libc.As[bool](retval) = loadedv730
	goto _return

sw_bb731:
	*libc.As[byte](result) = 1
	v264 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol732 = libc.Ptr(&libc.As[TSLexer](v264).F1)
	*libc.As[int16](result_symbol732) = 2
	v265 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end733 = libc.Ptr(&libc.As[TSLexer](v265).F3)
	v266 = *libc.As[unsafe.Pointer](mark_end733)
	v267 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v266)(v267)
	v268 = *libc.As[int32](lookahead)
	cmp734 = v268 == 101
	if cmp734 {
		goto if_then736
	} else {
		goto if_end737
	}

if_then736:
	*libc.As[int16](state_addr) = 41
	goto next_state

if_end737:
	v269 = *libc.As[byte](result)
	loadedv738 = (v269 & 1) != 0
	*libc.As[bool](retval) = loadedv738
	goto _return

sw_bb739:
	*libc.As[byte](result) = 1
	v270 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol740 = libc.Ptr(&libc.As[TSLexer](v270).F1)
	*libc.As[int16](result_symbol740) = 2
	v271 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end741 = libc.Ptr(&libc.As[TSLexer](v271).F3)
	v272 = *libc.As[unsafe.Pointer](mark_end741)
	v273 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v272)(v273)
	v274 = *libc.As[int32](lookahead)
	cmp742 = v274 == 111
	if cmp742 {
		goto if_then744
	} else {
		goto if_end745
	}

if_then744:
	*libc.As[int16](state_addr) = 68
	goto next_state

if_end745:
	v275 = *libc.As[int32](lookahead)
	cmp746 = v275 == 117
	if cmp746 {
		goto if_then748
	} else {
		goto if_end749
	}

if_then748:
	*libc.As[int16](state_addr) = 37
	goto next_state

if_end749:
	v276 = *libc.As[byte](result)
	loadedv750 = (v276 & 1) != 0
	*libc.As[bool](retval) = loadedv750
	goto _return

sw_bb751:
	*libc.As[byte](result) = 1
	v277 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol752 = libc.Ptr(&libc.As[TSLexer](v277).F1)
	*libc.As[int16](result_symbol752) = 3
	v278 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end753 = libc.Ptr(&libc.As[TSLexer](v278).F3)
	v279 = *libc.As[unsafe.Pointer](mark_end753)
	v280 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v279)(v280)
	v281 = *libc.As[int32](lookahead)
	cmp754 = v281 == 45
	if cmp754 {
		goto if_then756
	} else {
		goto if_end757
	}

if_then756:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end757:
	v282 = *libc.As[int32](lookahead)
	cmp758 = v282 == 46
	if cmp758 {
		goto if_then760
	} else {
		goto if_end761
	}

if_then760:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end761:
	v283 = *libc.As[int32](lookahead)
	cmp762 = v283 == 58
	if cmp762 {
		goto if_then764
	} else {
		goto if_end765
	}

if_then764:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end765:
	v284 = *libc.As[int32](lookahead)
	cmp766 = v284 == 56
	if cmp766 {
		goto if_then771
	} else {
		goto lor_lhs_false768
	}

lor_lhs_false768:
	v285 = *libc.As[int32](lookahead)
	cmp769 = v285 == 57
	if cmp769 {
		goto if_then771
	} else {
		goto if_end772
	}

if_then771:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end772:
	v286 = *libc.As[int32](lookahead)
	cmp773 = 48 <= v286
	if cmp773 {
		goto land_lhs_true775
	} else {
		goto lor_lhs_false778
	}

land_lhs_true775:
	v287 = *libc.As[int32](lookahead)
	cmp776 = v287 <= 55
	if cmp776 {
		goto if_then781
	} else {
		goto lor_lhs_false778
	}

lor_lhs_false778:
	v288 = *libc.As[int32](lookahead)
	cmp779 = v288 == 95
	if cmp779 {
		goto if_then781
	} else {
		goto if_end782
	}

if_then781:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end782:
	v289 = *libc.As[byte](result)
	loadedv783 = (v289 & 1) != 0
	*libc.As[bool](retval) = loadedv783
	goto _return

sw_bb784:
	*libc.As[byte](result) = 1
	v290 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol785 = libc.Ptr(&libc.As[TSLexer](v290).F1)
	*libc.As[int16](result_symbol785) = 3
	v291 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end786 = libc.Ptr(&libc.As[TSLexer](v291).F3)
	v292 = *libc.As[unsafe.Pointer](mark_end786)
	v293 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v292)(v293)
	v294 = *libc.As[int32](lookahead)
	cmp787 = v294 == 45
	if cmp787 {
		goto if_then789
	} else {
		goto if_end790
	}

if_then789:
	*libc.As[int16](state_addr) = 53
	goto next_state

if_end790:
	v295 = *libc.As[int32](lookahead)
	cmp791 = v295 == 46
	if cmp791 {
		goto if_then793
	} else {
		goto if_end794
	}

if_then793:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end794:
	v296 = *libc.As[int32](lookahead)
	cmp795 = v296 == 58
	if cmp795 {
		goto if_then797
	} else {
		goto if_end798
	}

if_then797:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end798:
	v297 = *libc.As[int32](lookahead)
	cmp799 = 48 <= v297
	if cmp799 {
		goto land_lhs_true801
	} else {
		goto lor_lhs_false804
	}

land_lhs_true801:
	v298 = *libc.As[int32](lookahead)
	cmp802 = v298 <= 57
	if cmp802 {
		goto if_then807
	} else {
		goto lor_lhs_false804
	}

lor_lhs_false804:
	v299 = *libc.As[int32](lookahead)
	cmp805 = v299 == 95
	if cmp805 {
		goto if_then807
	} else {
		goto if_end808
	}

if_then807:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end808:
	v300 = *libc.As[byte](result)
	loadedv809 = (v300 & 1) != 0
	*libc.As[bool](retval) = loadedv809
	goto _return

sw_bb810:
	*libc.As[byte](result) = 1
	v301 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol811 = libc.Ptr(&libc.As[TSLexer](v301).F1)
	*libc.As[int16](result_symbol811) = 3
	v302 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end812 = libc.Ptr(&libc.As[TSLexer](v302).F3)
	v303 = *libc.As[unsafe.Pointer](mark_end812)
	v304 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v303)(v304)
	v305 = *libc.As[int32](lookahead)
	cmp813 = v305 == 46
	if cmp813 {
		goto if_then815
	} else {
		goto if_end816
	}

if_then815:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end816:
	v306 = *libc.As[int32](lookahead)
	cmp817 = v306 == 58
	if cmp817 {
		goto if_then819
	} else {
		goto if_end820
	}

if_then819:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end820:
	v307 = *libc.As[int32](lookahead)
	cmp821 = v307 == 95
	if cmp821 {
		goto if_then823
	} else {
		goto if_end824
	}

if_then823:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end824:
	v308 = *libc.As[int32](lookahead)
	cmp825 = v308 == 98
	if cmp825 {
		goto if_then827
	} else {
		goto if_end828
	}

if_then827:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end828:
	v309 = *libc.As[int32](lookahead)
	cmp829 = v309 == 120
	if cmp829 {
		goto if_then831
	} else {
		goto if_end832
	}

if_then831:
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end832:
	v310 = *libc.As[int32](lookahead)
	cmp833 = v310 == 56
	if cmp833 {
		goto if_then838
	} else {
		goto lor_lhs_false835
	}

lor_lhs_false835:
	v311 = *libc.As[int32](lookahead)
	cmp836 = v311 == 57
	if cmp836 {
		goto if_then838
	} else {
		goto if_end839
	}

if_then838:
	*libc.As[int16](state_addr) = 5
	goto next_state

if_end839:
	v312 = *libc.As[int32](lookahead)
	cmp840 = 48 <= v312
	if cmp840 {
		goto land_lhs_true842
	} else {
		goto if_end846
	}

land_lhs_true842:
	v313 = *libc.As[int32](lookahead)
	cmp843 = v313 <= 55
	if cmp843 {
		goto if_then845
	} else {
		goto if_end846
	}

if_then845:
	*libc.As[int16](state_addr) = 77
	goto next_state

if_end846:
	v314 = *libc.As[byte](result)
	loadedv847 = (v314 & 1) != 0
	*libc.As[bool](retval) = loadedv847
	goto _return

sw_bb848:
	*libc.As[byte](result) = 1
	v315 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol849 = libc.Ptr(&libc.As[TSLexer](v315).F1)
	*libc.As[int16](result_symbol849) = 3
	v316 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end850 = libc.Ptr(&libc.As[TSLexer](v316).F3)
	v317 = *libc.As[unsafe.Pointer](mark_end850)
	v318 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v317)(v318)
	v319 = *libc.As[int32](lookahead)
	cmp851 = v319 == 46
	if cmp851 {
		goto if_then853
	} else {
		goto if_end854
	}

if_then853:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end854:
	v320 = *libc.As[int32](lookahead)
	cmp855 = v320 == 58
	if cmp855 {
		goto if_then857
	} else {
		goto if_end858
	}

if_then857:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end858:
	v321 = *libc.As[int32](lookahead)
	cmp859 = v321 == 95
	if cmp859 {
		goto if_then861
	} else {
		goto if_end862
	}

if_then861:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end862:
	v322 = *libc.As[int32](lookahead)
	cmp863 = v322 == 56
	if cmp863 {
		goto if_then868
	} else {
		goto lor_lhs_false865
	}

lor_lhs_false865:
	v323 = *libc.As[int32](lookahead)
	cmp866 = v323 == 57
	if cmp866 {
		goto if_then868
	} else {
		goto if_end869
	}

if_then868:
	*libc.As[int16](state_addr) = 1
	goto next_state

if_end869:
	v324 = *libc.As[int32](lookahead)
	cmp870 = 48 <= v324
	if cmp870 {
		goto land_lhs_true872
	} else {
		goto if_end876
	}

land_lhs_true872:
	v325 = *libc.As[int32](lookahead)
	cmp873 = v325 <= 55
	if cmp873 {
		goto if_then875
	} else {
		goto if_end876
	}

if_then875:
	*libc.As[int16](state_addr) = 73
	goto next_state

if_end876:
	v326 = *libc.As[byte](result)
	loadedv877 = (v326 & 1) != 0
	*libc.As[bool](retval) = loadedv877
	goto _return

sw_bb878:
	*libc.As[byte](result) = 1
	v327 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol879 = libc.Ptr(&libc.As[TSLexer](v327).F1)
	*libc.As[int16](result_symbol879) = 3
	v328 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end880 = libc.Ptr(&libc.As[TSLexer](v328).F3)
	v329 = *libc.As[unsafe.Pointer](mark_end880)
	v330 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v329)(v330)
	v331 = *libc.As[int32](lookahead)
	cmp881 = v331 == 46
	if cmp881 {
		goto if_then883
	} else {
		goto if_end884
	}

if_then883:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end884:
	v332 = *libc.As[int32](lookahead)
	cmp885 = v332 == 58
	if cmp885 {
		goto if_then887
	} else {
		goto if_end888
	}

if_then887:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end888:
	v333 = *libc.As[int32](lookahead)
	cmp889 = v333 == 95
	if cmp889 {
		goto if_then891
	} else {
		goto if_end892
	}

if_then891:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end892:
	v334 = *libc.As[int32](lookahead)
	cmp893 = v334 == 56
	if cmp893 {
		goto if_then898
	} else {
		goto lor_lhs_false895
	}

lor_lhs_false895:
	v335 = *libc.As[int32](lookahead)
	cmp896 = v335 == 57
	if cmp896 {
		goto if_then898
	} else {
		goto if_end899
	}

if_then898:
	*libc.As[int16](state_addr) = 4
	goto next_state

if_end899:
	v336 = *libc.As[int32](lookahead)
	cmp900 = 48 <= v336
	if cmp900 {
		goto land_lhs_true902
	} else {
		goto if_end906
	}

land_lhs_true902:
	v337 = *libc.As[int32](lookahead)
	cmp903 = v337 <= 55
	if cmp903 {
		goto if_then905
	} else {
		goto if_end906
	}

if_then905:
	*libc.As[int16](state_addr) = 76
	goto next_state

if_end906:
	v338 = *libc.As[byte](result)
	loadedv907 = (v338 & 1) != 0
	*libc.As[bool](retval) = loadedv907
	goto _return

sw_bb908:
	*libc.As[byte](result) = 1
	v339 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol909 = libc.Ptr(&libc.As[TSLexer](v339).F1)
	*libc.As[int16](result_symbol909) = 3
	v340 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end910 = libc.Ptr(&libc.As[TSLexer](v340).F3)
	v341 = *libc.As[unsafe.Pointer](mark_end910)
	v342 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v341)(v342)
	v343 = *libc.As[int32](lookahead)
	cmp911 = v343 == 46
	if cmp911 {
		goto if_then913
	} else {
		goto if_end914
	}

if_then913:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end914:
	v344 = *libc.As[int32](lookahead)
	cmp915 = v344 == 58
	if cmp915 {
		goto if_then917
	} else {
		goto if_end918
	}

if_then917:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end918:
	v345 = *libc.As[int32](lookahead)
	cmp919 = v345 == 98
	if cmp919 {
		goto if_then921
	} else {
		goto if_end922
	}

if_then921:
	*libc.As[int16](state_addr) = 49
	goto next_state

if_end922:
	v346 = *libc.As[int32](lookahead)
	cmp923 = v346 == 120
	if cmp923 {
		goto if_then925
	} else {
		goto if_end926
	}

if_then925:
	*libc.As[int16](state_addr) = 65
	goto next_state

if_end926:
	v347 = *libc.As[int32](lookahead)
	cmp927 = v347 == 56
	if cmp927 {
		goto if_then932
	} else {
		goto lor_lhs_false929
	}

lor_lhs_false929:
	v348 = *libc.As[int32](lookahead)
	cmp930 = v348 == 57
	if cmp930 {
		goto if_then932
	} else {
		goto if_end933
	}

if_then932:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end933:
	v349 = *libc.As[int32](lookahead)
	cmp934 = 48 <= v349
	if cmp934 {
		goto land_lhs_true936
	} else {
		goto lor_lhs_false939
	}

land_lhs_true936:
	v350 = *libc.As[int32](lookahead)
	cmp937 = v350 <= 55
	if cmp937 {
		goto if_then942
	} else {
		goto lor_lhs_false939
	}

lor_lhs_false939:
	v351 = *libc.As[int32](lookahead)
	cmp940 = v351 == 95
	if cmp940 {
		goto if_then942
	} else {
		goto if_end943
	}

if_then942:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end943:
	v352 = *libc.As[byte](result)
	loadedv944 = (v352 & 1) != 0
	*libc.As[bool](retval) = loadedv944
	goto _return

sw_bb945:
	*libc.As[byte](result) = 1
	v353 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol946 = libc.Ptr(&libc.As[TSLexer](v353).F1)
	*libc.As[int16](result_symbol946) = 3
	v354 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end947 = libc.Ptr(&libc.As[TSLexer](v354).F3)
	v355 = *libc.As[unsafe.Pointer](mark_end947)
	v356 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v355)(v356)
	v357 = *libc.As[int32](lookahead)
	cmp948 = v357 == 46
	if cmp948 {
		goto if_then950
	} else {
		goto if_end951
	}

if_then950:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end951:
	v358 = *libc.As[int32](lookahead)
	cmp952 = v358 == 58
	if cmp952 {
		goto if_then954
	} else {
		goto if_end955
	}

if_then954:
	*libc.As[int16](state_addr) = 50
	goto next_state

if_end955:
	v359 = *libc.As[int32](lookahead)
	cmp956 = v359 == 56
	if cmp956 {
		goto if_then961
	} else {
		goto lor_lhs_false958
	}

lor_lhs_false958:
	v360 = *libc.As[int32](lookahead)
	cmp959 = v360 == 57
	if cmp959 {
		goto if_then961
	} else {
		goto if_end962
	}

if_then961:
	*libc.As[int16](state_addr) = 6
	goto next_state

if_end962:
	v361 = *libc.As[int32](lookahead)
	cmp963 = 48 <= v361
	if cmp963 {
		goto land_lhs_true965
	} else {
		goto lor_lhs_false968
	}

land_lhs_true965:
	v362 = *libc.As[int32](lookahead)
	cmp966 = v362 <= 55
	if cmp966 {
		goto if_then971
	} else {
		goto lor_lhs_false968
	}

lor_lhs_false968:
	v363 = *libc.As[int32](lookahead)
	cmp969 = v363 == 95
	if cmp969 {
		goto if_then971
	} else {
		goto if_end972
	}

if_then971:
	*libc.As[int16](state_addr) = 79
	goto next_state

if_end972:
	v364 = *libc.As[byte](result)
	loadedv973 = (v364 & 1) != 0
	*libc.As[bool](retval) = loadedv973
	goto _return

sw_bb974:
	*libc.As[byte](result) = 1
	v365 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol975 = libc.Ptr(&libc.As[TSLexer](v365).F1)
	*libc.As[int16](result_symbol975) = 3
	v366 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end976 = libc.Ptr(&libc.As[TSLexer](v366).F3)
	v367 = *libc.As[unsafe.Pointer](mark_end976)
	v368 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v367)(v368)
	v369 = *libc.As[int32](lookahead)
	cmp977 = v369 == 46
	if cmp977 {
		goto if_then979
	} else {
		goto if_end980
	}

if_then979:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end980:
	v370 = *libc.As[int32](lookahead)
	cmp981 = v370 == 58
	if cmp981 {
		goto if_then983
	} else {
		goto if_end984
	}

if_then983:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end984:
	v371 = *libc.As[int32](lookahead)
	cmp985 = v371 == 95
	if cmp985 {
		goto if_then987
	} else {
		goto if_end988
	}

if_then987:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end988:
	v372 = *libc.As[int32](lookahead)
	cmp989 = 48 <= v372
	if cmp989 {
		goto land_lhs_true991
	} else {
		goto if_end995
	}

land_lhs_true991:
	v373 = *libc.As[int32](lookahead)
	cmp992 = v373 <= 57
	if cmp992 {
		goto if_then994
	} else {
		goto if_end995
	}

if_then994:
	*libc.As[int16](state_addr) = 74
	goto next_state

if_end995:
	v374 = *libc.As[byte](result)
	loadedv996 = (v374 & 1) != 0
	*libc.As[bool](retval) = loadedv996
	goto _return

sw_bb997:
	*libc.As[byte](result) = 1
	v375 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol998 = libc.Ptr(&libc.As[TSLexer](v375).F1)
	*libc.As[int16](result_symbol998) = 3
	v376 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end999 = libc.Ptr(&libc.As[TSLexer](v376).F3)
	v377 = *libc.As[unsafe.Pointer](mark_end999)
	v378 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v377)(v378)
	v379 = *libc.As[int32](lookahead)
	cmp1000 = v379 == 46
	if cmp1000 {
		goto if_then1002
	} else {
		goto if_end1003
	}

if_then1002:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1003:
	v380 = *libc.As[int32](lookahead)
	cmp1004 = v380 == 58
	if cmp1004 {
		goto if_then1006
	} else {
		goto if_end1007
	}

if_then1006:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end1007:
	v381 = *libc.As[int32](lookahead)
	cmp1008 = v381 == 95
	if cmp1008 {
		goto if_then1010
	} else {
		goto if_end1011
	}

if_then1010:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end1011:
	v382 = *libc.As[int32](lookahead)
	cmp1012 = 48 <= v382
	if cmp1012 {
		goto land_lhs_true1014
	} else {
		goto if_end1018
	}

land_lhs_true1014:
	v383 = *libc.As[int32](lookahead)
	cmp1015 = v383 <= 57
	if cmp1015 {
		goto if_then1017
	} else {
		goto if_end1018
	}

if_then1017:
	*libc.As[int16](state_addr) = 80
	goto next_state

if_end1018:
	v384 = *libc.As[byte](result)
	loadedv1019 = (v384 & 1) != 0
	*libc.As[bool](retval) = loadedv1019
	goto _return

sw_bb1020:
	*libc.As[byte](result) = 1
	v385 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1021 = libc.Ptr(&libc.As[TSLexer](v385).F1)
	*libc.As[int16](result_symbol1021) = 3
	v386 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1022 = libc.Ptr(&libc.As[TSLexer](v386).F3)
	v387 = *libc.As[unsafe.Pointer](mark_end1022)
	v388 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v387)(v388)
	v389 = *libc.As[int32](lookahead)
	cmp1023 = v389 == 46
	if cmp1023 {
		goto if_then1025
	} else {
		goto if_end1026
	}

if_then1025:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1026:
	v390 = *libc.As[int32](lookahead)
	cmp1027 = v390 == 58
	if cmp1027 {
		goto if_then1029
	} else {
		goto if_end1030
	}

if_then1029:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end1030:
	v391 = *libc.As[int32](lookahead)
	cmp1031 = v391 == 95
	if cmp1031 {
		goto if_then1033
	} else {
		goto if_end1034
	}

if_then1033:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end1034:
	v392 = *libc.As[int32](lookahead)
	cmp1035 = 48 <= v392
	if cmp1035 {
		goto land_lhs_true1037
	} else {
		goto if_end1041
	}

land_lhs_true1037:
	v393 = *libc.As[int32](lookahead)
	cmp1038 = v393 <= 57
	if cmp1038 {
		goto if_then1040
	} else {
		goto if_end1041
	}

if_then1040:
	*libc.As[int16](state_addr) = 81
	goto next_state

if_end1041:
	v394 = *libc.As[byte](result)
	loadedv1042 = (v394 & 1) != 0
	*libc.As[bool](retval) = loadedv1042
	goto _return

sw_bb1043:
	*libc.As[byte](result) = 1
	v395 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1044 = libc.Ptr(&libc.As[TSLexer](v395).F1)
	*libc.As[int16](result_symbol1044) = 3
	v396 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1045 = libc.Ptr(&libc.As[TSLexer](v396).F3)
	v397 = *libc.As[unsafe.Pointer](mark_end1045)
	v398 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v397)(v398)
	v399 = *libc.As[int32](lookahead)
	cmp1046 = v399 == 46
	if cmp1046 {
		goto if_then1048
	} else {
		goto if_end1049
	}

if_then1048:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1049:
	v400 = *libc.As[int32](lookahead)
	cmp1050 = v400 == 58
	if cmp1050 {
		goto if_then1052
	} else {
		goto if_end1053
	}

if_then1052:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end1053:
	v401 = *libc.As[int32](lookahead)
	cmp1054 = 48 <= v401
	if cmp1054 {
		goto land_lhs_true1056
	} else {
		goto lor_lhs_false1059
	}

land_lhs_true1056:
	v402 = *libc.As[int32](lookahead)
	cmp1057 = v402 <= 57
	if cmp1057 {
		goto if_then1062
	} else {
		goto lor_lhs_false1059
	}

lor_lhs_false1059:
	v403 = *libc.As[int32](lookahead)
	cmp1060 = v403 == 95
	if cmp1060 {
		goto if_then1062
	} else {
		goto if_end1063
	}

if_then1062:
	*libc.As[int16](state_addr) = 83
	goto next_state

if_end1063:
	v404 = *libc.As[byte](result)
	loadedv1064 = (v404 & 1) != 0
	*libc.As[bool](retval) = loadedv1064
	goto _return

sw_bb1065:
	*libc.As[byte](result) = 1
	v405 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1066 = libc.Ptr(&libc.As[TSLexer](v405).F1)
	*libc.As[int16](result_symbol1066) = 3
	v406 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1067 = libc.Ptr(&libc.As[TSLexer](v406).F3)
	v407 = *libc.As[unsafe.Pointer](mark_end1067)
	v408 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v407)(v408)
	v409 = *libc.As[int32](lookahead)
	cmp1068 = v409 == 46
	if cmp1068 {
		goto if_then1070
	} else {
		goto if_end1071
	}

if_then1070:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end1071:
	v410 = *libc.As[int32](lookahead)
	cmp1072 = v410 == 58
	if cmp1072 {
		goto if_then1074
	} else {
		goto if_end1075
	}

if_then1074:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end1075:
	v411 = *libc.As[byte](result)
	loadedv1076 = (v411 & 1) != 0
	*libc.As[bool](retval) = loadedv1076
	goto _return

sw_bb1077:
	*libc.As[byte](result) = 1
	v412 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1078 = libc.Ptr(&libc.As[TSLexer](v412).F1)
	*libc.As[int16](result_symbol1078) = 3
	v413 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1079 = libc.Ptr(&libc.As[TSLexer](v413).F3)
	v414 = *libc.As[unsafe.Pointer](mark_end1079)
	v415 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v414)(v415)
	v416 = *libc.As[int32](lookahead)
	cmp1080 = v416 == 46
	if cmp1080 {
		goto if_then1082
	} else {
		goto if_end1083
	}

if_then1082:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end1083:
	v417 = *libc.As[int32](lookahead)
	cmp1084 = v417 == 58
	if cmp1084 {
		goto if_then1086
	} else {
		goto if_end1087
	}

if_then1086:
	*libc.As[int16](state_addr) = 51
	goto next_state

if_end1087:
	v418 = *libc.As[int32](lookahead)
	cmp1088 = 48 <= v418
	if cmp1088 {
		goto land_lhs_true1090
	} else {
		goto if_end1094
	}

land_lhs_true1090:
	v419 = *libc.As[int32](lookahead)
	cmp1091 = v419 <= 57
	if cmp1091 {
		goto if_then1093
	} else {
		goto if_end1094
	}

if_then1093:
	*libc.As[int16](state_addr) = 84
	goto next_state

if_end1094:
	v420 = *libc.As[byte](result)
	loadedv1095 = (v420 & 1) != 0
	*libc.As[bool](retval) = loadedv1095
	goto _return

sw_bb1096:
	*libc.As[byte](result) = 1
	v421 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1097 = libc.Ptr(&libc.As[TSLexer](v421).F1)
	*libc.As[int16](result_symbol1097) = 3
	v422 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1098 = libc.Ptr(&libc.As[TSLexer](v422).F3)
	v423 = *libc.As[unsafe.Pointer](mark_end1098)
	v424 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v423)(v424)
	v425 = *libc.As[int32](lookahead)
	cmp1099 = v425 == 48
	if cmp1099 {
		goto if_then1107
	} else {
		goto lor_lhs_false1101
	}

lor_lhs_false1101:
	v426 = *libc.As[int32](lookahead)
	cmp1102 = v426 == 49
	if cmp1102 {
		goto if_then1107
	} else {
		goto lor_lhs_false1104
	}

lor_lhs_false1104:
	v427 = *libc.As[int32](lookahead)
	cmp1105 = v427 == 95
	if cmp1105 {
		goto if_then1107
	} else {
		goto if_end1108
	}

if_then1107:
	*libc.As[int16](state_addr) = 86
	goto next_state

if_end1108:
	v428 = *libc.As[byte](result)
	loadedv1109 = (v428 & 1) != 0
	*libc.As[bool](retval) = loadedv1109
	goto _return

sw_bb1110:
	*libc.As[byte](result) = 1
	v429 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1111 = libc.Ptr(&libc.As[TSLexer](v429).F1)
	*libc.As[int16](result_symbol1111) = 3
	v430 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1112 = libc.Ptr(&libc.As[TSLexer](v430).F3)
	v431 = *libc.As[unsafe.Pointer](mark_end1112)
	v432 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v431)(v432)
	v433 = *libc.As[int32](lookahead)
	cmp1113 = 48 <= v433
	if cmp1113 {
		goto land_lhs_true1115
	} else {
		goto lor_lhs_false1118
	}

land_lhs_true1115:
	v434 = *libc.As[int32](lookahead)
	cmp1116 = v434 <= 57
	if cmp1116 {
		goto if_then1133
	} else {
		goto lor_lhs_false1118
	}

lor_lhs_false1118:
	v435 = *libc.As[int32](lookahead)
	cmp1119 = 65 <= v435
	if cmp1119 {
		goto land_lhs_true1121
	} else {
		goto lor_lhs_false1124
	}

land_lhs_true1121:
	v436 = *libc.As[int32](lookahead)
	cmp1122 = v436 <= 70
	if cmp1122 {
		goto if_then1133
	} else {
		goto lor_lhs_false1124
	}

lor_lhs_false1124:
	v437 = *libc.As[int32](lookahead)
	cmp1125 = v437 == 95
	if cmp1125 {
		goto if_then1133
	} else {
		goto lor_lhs_false1127
	}

lor_lhs_false1127:
	v438 = *libc.As[int32](lookahead)
	cmp1128 = 97 <= v438
	if cmp1128 {
		goto land_lhs_true1130
	} else {
		goto if_end1134
	}

land_lhs_true1130:
	v439 = *libc.As[int32](lookahead)
	cmp1131 = v439 <= 102
	if cmp1131 {
		goto if_then1133
	} else {
		goto if_end1134
	}

if_then1133:
	*libc.As[int16](state_addr) = 87
	goto next_state

if_end1134:
	v440 = *libc.As[byte](result)
	loadedv1135 = (v440 & 1) != 0
	*libc.As[bool](retval) = loadedv1135
	goto _return

sw_bb1136:
	*libc.As[byte](result) = 1
	v441 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1137 = libc.Ptr(&libc.As[TSLexer](v441).F1)
	*libc.As[int16](result_symbol1137) = 4
	v442 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1138 = libc.Ptr(&libc.As[TSLexer](v442).F3)
	v443 = *libc.As[unsafe.Pointer](mark_end1138)
	v444 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v443)(v444)
	v445 = *libc.As[byte](result)
	loadedv1139 = (v445 & 1) != 0
	*libc.As[bool](retval) = loadedv1139
	goto _return

sw_bb1140:
	*libc.As[byte](result) = 1
	v446 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1141 = libc.Ptr(&libc.As[TSLexer](v446).F1)
	*libc.As[int16](result_symbol1141) = 4
	v447 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1142 = libc.Ptr(&libc.As[TSLexer](v447).F3)
	v448 = *libc.As[unsafe.Pointer](mark_end1142)
	v449 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v448)(v449)
	v450 = *libc.As[int32](lookahead)
	cmp1143 = v450 == 73
	if cmp1143 {
		goto if_then1145
	} else {
		goto if_end1146
	}

if_then1145:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end1146:
	v451 = *libc.As[int32](lookahead)
	cmp1147 = v451 == 78
	if cmp1147 {
		goto if_then1149
	} else {
		goto if_end1150
	}

if_then1149:
	*libc.As[int16](state_addr) = 14
	goto next_state

if_end1150:
	v452 = *libc.As[int32](lookahead)
	cmp1151 = v452 == 105
	if cmp1151 {
		goto if_then1153
	} else {
		goto if_end1154
	}

if_then1153:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end1154:
	v453 = *libc.As[int32](lookahead)
	cmp1155 = v453 == 110
	if cmp1155 {
		goto if_then1157
	} else {
		goto if_end1158
	}

if_then1157:
	*libc.As[int16](state_addr) = 30
	goto next_state

if_end1158:
	v454 = *libc.As[int32](lookahead)
	cmp1159 = v454 == 69
	if cmp1159 {
		goto if_then1164
	} else {
		goto lor_lhs_false1161
	}

lor_lhs_false1161:
	v455 = *libc.As[int32](lookahead)
	cmp1162 = v455 == 101
	if cmp1162 {
		goto if_then1164
	} else {
		goto if_end1165
	}

if_then1164:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end1165:
	v456 = *libc.As[int32](lookahead)
	cmp1166 = v456 == 46
	if cmp1166 {
		goto if_then1174
	} else {
		goto lor_lhs_false1168
	}

lor_lhs_false1168:
	v457 = *libc.As[int32](lookahead)
	cmp1169 = 48 <= v457
	if cmp1169 {
		goto land_lhs_true1171
	} else {
		goto if_end1175
	}

land_lhs_true1171:
	v458 = *libc.As[int32](lookahead)
	cmp1172 = v458 <= 57
	if cmp1172 {
		goto if_then1174
	} else {
		goto if_end1175
	}

if_then1174:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1175:
	v459 = *libc.As[byte](result)
	loadedv1176 = (v459 & 1) != 0
	*libc.As[bool](retval) = loadedv1176
	goto _return

sw_bb1177:
	*libc.As[byte](result) = 1
	v460 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1178 = libc.Ptr(&libc.As[TSLexer](v460).F1)
	*libc.As[int16](result_symbol1178) = 4
	v461 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1179 = libc.Ptr(&libc.As[TSLexer](v461).F3)
	v462 = *libc.As[unsafe.Pointer](mark_end1179)
	v463 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v462)(v463)
	v464 = *libc.As[int32](lookahead)
	cmp1180 = v464 == 73
	if cmp1180 {
		goto if_then1182
	} else {
		goto if_end1183
	}

if_then1182:
	*libc.As[int16](state_addr) = 23
	goto next_state

if_end1183:
	v465 = *libc.As[int32](lookahead)
	cmp1184 = v465 == 105
	if cmp1184 {
		goto if_then1186
	} else {
		goto if_end1187
	}

if_then1186:
	*libc.As[int16](state_addr) = 39
	goto next_state

if_end1187:
	v466 = *libc.As[int32](lookahead)
	cmp1188 = v466 == 69
	if cmp1188 {
		goto if_then1193
	} else {
		goto lor_lhs_false1190
	}

lor_lhs_false1190:
	v467 = *libc.As[int32](lookahead)
	cmp1191 = v467 == 101
	if cmp1191 {
		goto if_then1193
	} else {
		goto if_end1194
	}

if_then1193:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end1194:
	v468 = *libc.As[int32](lookahead)
	cmp1195 = v468 == 46
	if cmp1195 {
		goto if_then1203
	} else {
		goto lor_lhs_false1197
	}

lor_lhs_false1197:
	v469 = *libc.As[int32](lookahead)
	cmp1198 = 48 <= v469
	if cmp1198 {
		goto land_lhs_true1200
	} else {
		goto if_end1204
	}

land_lhs_true1200:
	v470 = *libc.As[int32](lookahead)
	cmp1201 = v470 <= 57
	if cmp1201 {
		goto if_then1203
	} else {
		goto if_end1204
	}

if_then1203:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1204:
	v471 = *libc.As[byte](result)
	loadedv1205 = (v471 & 1) != 0
	*libc.As[bool](retval) = loadedv1205
	goto _return

sw_bb1206:
	*libc.As[byte](result) = 1
	v472 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1207 = libc.Ptr(&libc.As[TSLexer](v472).F1)
	*libc.As[int16](result_symbol1207) = 4
	v473 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1208 = libc.Ptr(&libc.As[TSLexer](v473).F3)
	v474 = *libc.As[unsafe.Pointer](mark_end1208)
	v475 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v474)(v475)
	v476 = *libc.As[int32](lookahead)
	cmp1209 = v476 == 69
	if cmp1209 {
		goto if_then1214
	} else {
		goto lor_lhs_false1211
	}

lor_lhs_false1211:
	v477 = *libc.As[int32](lookahead)
	cmp1212 = v477 == 101
	if cmp1212 {
		goto if_then1214
	} else {
		goto if_end1215
	}

if_then1214:
	*libc.As[int16](state_addr) = 48
	goto next_state

if_end1215:
	v478 = *libc.As[int32](lookahead)
	cmp1216 = v478 == 46
	if cmp1216 {
		goto if_then1224
	} else {
		goto lor_lhs_false1218
	}

lor_lhs_false1218:
	v479 = *libc.As[int32](lookahead)
	cmp1219 = 48 <= v479
	if cmp1219 {
		goto land_lhs_true1221
	} else {
		goto if_end1225
	}

land_lhs_true1221:
	v480 = *libc.As[int32](lookahead)
	cmp1222 = v480 <= 57
	if cmp1222 {
		goto if_then1224
	} else {
		goto if_end1225
	}

if_then1224:
	*libc.As[int16](state_addr) = 91
	goto next_state

if_end1225:
	v481 = *libc.As[byte](result)
	loadedv1226 = (v481 & 1) != 0
	*libc.As[bool](retval) = loadedv1226
	goto _return

sw_bb1227:
	*libc.As[byte](result) = 1
	v482 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1228 = libc.Ptr(&libc.As[TSLexer](v482).F1)
	*libc.As[int16](result_symbol1228) = 4
	v483 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1229 = libc.Ptr(&libc.As[TSLexer](v483).F3)
	v484 = *libc.As[unsafe.Pointer](mark_end1229)
	v485 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v484)(v485)
	v486 = *libc.As[int32](lookahead)
	cmp1230 = 48 <= v486
	if cmp1230 {
		goto land_lhs_true1232
	} else {
		goto if_end1236
	}

land_lhs_true1232:
	v487 = *libc.As[int32](lookahead)
	cmp1233 = v487 <= 57
	if cmp1233 {
		goto if_then1235
	} else {
		goto if_end1236
	}

if_then1235:
	*libc.As[int16](state_addr) = 92
	goto next_state

if_end1236:
	v488 = *libc.As[byte](result)
	loadedv1237 = (v488 & 1) != 0
	*libc.As[bool](retval) = loadedv1237
	goto _return

sw_bb1238:
	*libc.As[byte](result) = 1
	v489 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1239 = libc.Ptr(&libc.As[TSLexer](v489).F1)
	*libc.As[int16](result_symbol1239) = 4
	v490 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1240 = libc.Ptr(&libc.As[TSLexer](v490).F3)
	v491 = *libc.As[unsafe.Pointer](mark_end1240)
	v492 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v491)(v492)
	v493 = *libc.As[int32](lookahead)
	cmp1241 = 48 <= v493
	if cmp1241 {
		goto land_lhs_true1243
	} else {
		goto lor_lhs_false1246
	}

land_lhs_true1243:
	v494 = *libc.As[int32](lookahead)
	cmp1244 = v494 <= 57
	if cmp1244 {
		goto if_then1249
	} else {
		goto lor_lhs_false1246
	}

lor_lhs_false1246:
	v495 = *libc.As[int32](lookahead)
	cmp1247 = v495 == 95
	if cmp1247 {
		goto if_then1249
	} else {
		goto if_end1250
	}

if_then1249:
	*libc.As[int16](state_addr) = 93
	goto next_state

if_end1250:
	v496 = *libc.As[byte](result)
	loadedv1251 = (v496 & 1) != 0
	*libc.As[bool](retval) = loadedv1251
	goto _return

sw_bb1252:
	*libc.As[byte](result) = 1
	v497 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1253 = libc.Ptr(&libc.As[TSLexer](v497).F1)
	*libc.As[int16](result_symbol1253) = 5
	v498 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1254 = libc.Ptr(&libc.As[TSLexer](v498).F3)
	v499 = *libc.As[unsafe.Pointer](mark_end1254)
	v500 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v499)(v500)
	v501 = *libc.As[byte](result)
	loadedv1255 = (v501 & 1) != 0
	*libc.As[bool](retval) = loadedv1255
	goto _return

sw_bb1256:
	*libc.As[byte](result) = 1
	v502 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1257 = libc.Ptr(&libc.As[TSLexer](v502).F1)
	*libc.As[int16](result_symbol1257) = 5
	v503 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1258 = libc.Ptr(&libc.As[TSLexer](v503).F3)
	v504 = *libc.As[unsafe.Pointer](mark_end1258)
	v505 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v504)(v505)
	v506 = *libc.As[int32](lookahead)
	cmp1259 = v506 == 46
	if cmp1259 {
		goto if_then1261
	} else {
		goto if_end1262
	}

if_then1261:
	*libc.As[int16](state_addr) = 58
	goto next_state

if_end1262:
	v507 = *libc.As[int32](lookahead)
	cmp1263 = v507 == 90
	if cmp1263 {
		goto if_then1265
	} else {
		goto if_end1266
	}

if_then1265:
	*libc.As[int16](state_addr) = 94
	goto next_state

if_end1266:
	v508 = *libc.As[int32](lookahead)
	cmp1267 = v508 == 9
	if cmp1267 {
		goto if_then1272
	} else {
		goto lor_lhs_false1269
	}

lor_lhs_false1269:
	v509 = *libc.As[int32](lookahead)
	cmp1270 = v509 == 32
	if cmp1270 {
		goto if_then1272
	} else {
		goto if_end1273
	}

if_then1272:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end1273:
	v510 = *libc.As[int32](lookahead)
	cmp1274 = v510 == 43
	if cmp1274 {
		goto if_then1279
	} else {
		goto lor_lhs_false1276
	}

lor_lhs_false1276:
	v511 = *libc.As[int32](lookahead)
	cmp1277 = v511 == 45
	if cmp1277 {
		goto if_then1279
	} else {
		goto if_end1280
	}

if_then1279:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end1280:
	v512 = *libc.As[byte](result)
	loadedv1281 = (v512 & 1) != 0
	*libc.As[bool](retval) = loadedv1281
	goto _return

sw_bb1282:
	*libc.As[byte](result) = 1
	v513 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1283 = libc.Ptr(&libc.As[TSLexer](v513).F1)
	*libc.As[int16](result_symbol1283) = 5
	v514 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1284 = libc.Ptr(&libc.As[TSLexer](v514).F3)
	v515 = *libc.As[unsafe.Pointer](mark_end1284)
	v516 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v515)(v516)
	v517 = *libc.As[int32](lookahead)
	cmp1285 = v517 == 58
	if cmp1285 {
		goto if_then1287
	} else {
		goto if_end1288
	}

if_then1287:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end1288:
	v518 = *libc.As[byte](result)
	loadedv1289 = (v518 & 1) != 0
	*libc.As[bool](retval) = loadedv1289
	goto _return

sw_bb1290:
	*libc.As[byte](result) = 1
	v519 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1291 = libc.Ptr(&libc.As[TSLexer](v519).F1)
	*libc.As[int16](result_symbol1291) = 5
	v520 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1292 = libc.Ptr(&libc.As[TSLexer](v520).F3)
	v521 = *libc.As[unsafe.Pointer](mark_end1292)
	v522 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v521)(v522)
	v523 = *libc.As[int32](lookahead)
	cmp1293 = v523 == 58
	if cmp1293 {
		goto if_then1295
	} else {
		goto if_end1296
	}

if_then1295:
	*libc.As[int16](state_addr) = 63
	goto next_state

if_end1296:
	v524 = *libc.As[int32](lookahead)
	cmp1297 = 48 <= v524
	if cmp1297 {
		goto land_lhs_true1299
	} else {
		goto if_end1303
	}

land_lhs_true1299:
	v525 = *libc.As[int32](lookahead)
	cmp1300 = v525 <= 57
	if cmp1300 {
		goto if_then1302
	} else {
		goto if_end1303
	}

if_then1302:
	*libc.As[int16](state_addr) = 96
	goto next_state

if_end1303:
	v526 = *libc.As[byte](result)
	loadedv1304 = (v526 & 1) != 0
	*libc.As[bool](retval) = loadedv1304
	goto _return

sw_bb1305:
	*libc.As[byte](result) = 1
	v527 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1306 = libc.Ptr(&libc.As[TSLexer](v527).F1)
	*libc.As[int16](result_symbol1306) = 5
	v528 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1307 = libc.Ptr(&libc.As[TSLexer](v528).F3)
	v529 = *libc.As[unsafe.Pointer](mark_end1307)
	v530 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v529)(v530)
	v531 = *libc.As[int32](lookahead)
	cmp1308 = v531 == 90
	if cmp1308 {
		goto if_then1310
	} else {
		goto if_end1311
	}

if_then1310:
	*libc.As[int16](state_addr) = 94
	goto next_state

if_end1311:
	v532 = *libc.As[int32](lookahead)
	cmp1312 = v532 == 9
	if cmp1312 {
		goto if_then1317
	} else {
		goto lor_lhs_false1314
	}

lor_lhs_false1314:
	v533 = *libc.As[int32](lookahead)
	cmp1315 = v533 == 32
	if cmp1315 {
		goto if_then1317
	} else {
		goto if_end1318
	}

if_then1317:
	*libc.As[int16](state_addr) = 28
	goto next_state

if_end1318:
	v534 = *libc.As[int32](lookahead)
	cmp1319 = v534 == 43
	if cmp1319 {
		goto if_then1324
	} else {
		goto lor_lhs_false1321
	}

lor_lhs_false1321:
	v535 = *libc.As[int32](lookahead)
	cmp1322 = v535 == 45
	if cmp1322 {
		goto if_then1324
	} else {
		goto if_end1325
	}

if_then1324:
	*libc.As[int16](state_addr) = 59
	goto next_state

if_end1325:
	v536 = *libc.As[int32](lookahead)
	cmp1326 = 48 <= v536
	if cmp1326 {
		goto land_lhs_true1328
	} else {
		goto if_end1332
	}

land_lhs_true1328:
	v537 = *libc.As[int32](lookahead)
	cmp1329 = v537 <= 57
	if cmp1329 {
		goto if_then1331
	} else {
		goto if_end1332
	}

if_then1331:
	*libc.As[int16](state_addr) = 98
	goto next_state

if_end1332:
	v538 = *libc.As[byte](result)
	loadedv1333 = (v538 & 1) != 0
	*libc.As[bool](retval) = loadedv1333
	goto _return

sw_bb1334:
	*libc.As[byte](result) = 1
	v539 = *libc.As[unsafe.Pointer](lexer_addr)
	result_symbol1335 = libc.Ptr(&libc.As[TSLexer](v539).F1)
	*libc.As[int16](result_symbol1335) = 5
	v540 = *libc.As[unsafe.Pointer](lexer_addr)
	mark_end1336 = libc.Ptr(&libc.As[TSLexer](v540).F3)
	v541 = *libc.As[unsafe.Pointer](mark_end1336)
	v542 = *libc.As[unsafe.Pointer](lexer_addr)
	libc.FuncFromCode[func(unsafe.Pointer)](v541)(v542)
	v543 = *libc.As[int32](lookahead)
	cmp1337 = v543 == 9
	if cmp1337 {
		goto if_then1342
	} else {
		goto lor_lhs_false1339
	}

lor_lhs_false1339:
	v544 = *libc.As[int32](lookahead)
	cmp1340 = v544 == 32
	if cmp1340 {
		goto if_then1342
	} else {
		goto if_end1343
	}

if_then1342:
	*libc.As[int16](state_addr) = 47
	goto next_state

if_end1343:
	v545 = *libc.As[int32](lookahead)
	cmp1344 = v545 == 84
	if cmp1344 {
		goto if_then1349
	} else {
		goto lor_lhs_false1346
	}

lor_lhs_false1346:
	v546 = *libc.As[int32](lookahead)
	cmp1347 = v546 == 116
	if cmp1347 {
		goto if_then1349
	} else {
		goto if_end1350
	}

if_then1349:
	*libc.As[int16](state_addr) = 55
	goto next_state

if_end1350:
	v547 = *libc.As[byte](result)
	loadedv1351 = (v547 & 1) != 0
	*libc.As[bool](retval) = loadedv1351
	goto _return

sw_default:
	*libc.As[bool](retval) = false
	goto _return

_return:
	v548 = *libc.As[bool](retval)
	return v548
}
