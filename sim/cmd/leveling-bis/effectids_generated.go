package main

// engineImplementedEffectItemIDs is every item id wowsims-forever
// registers SOME on-hit/on-use/proc effect for -- generated once, from
// the fork's own source (2026-09-28 weights-effects lane), NOT from
// this site's own data: for every file under sim/ whose name does not
// start with "_" (a leading underscore means Go itself does not
// compile the file -- see sim/priest/_items.go, currently excluded
// entirely), every numeric item-id constant used as the first
// argument to ANY registration call (core.NewItemEffect, or an
// itemhelpers.CreateWeapon*Proc* helper called with that constant)
// anywhere in that file counts as implemented. This is a heuristic,
// not a parse of Go's own call graph: it would mis-flag a constant
// that happens to be some OTHER function's first argument for an
// unrelated reason, which this engine build's own file conventions
// (one id, one registration call, per const) make very unlikely in
// practice -- see the lane report for the exact two-pass script.
//
// Regenerate by re-running that script against sim/enginever.Version's
// commit next: a false negative here (an item the engine now
// implements that this map still calls unimplemented) only costs an
// honest "not simulated" flag on a candidate that would otherwise be
// ranked correctly; this map is not consulted for correctness of the
// scored total, only for whether a slot's pick gets an engine-verified
// run (see rank.go's hasImplementedEffect) and whether report.go's
// effect_unmodelled flag is set.
// Generated from wowsims-forever 912a458f7 by sim/scripts/effectids.py.
var engineImplementedEffectItemIDs = map[int]bool{
	647: true, 754: true, 809: true, 810: true, 870: true, 871: true, 1168: true, 1728: true, 1982: true, 2163: true,
	2164: true, 2243: true, 2825: true, 3854: true, 5616: true, 6622: true, 7717: true, 7959: true, 8190: true, 9423: true,
	9425: true, 9449: true, 9511: true, 9639: true, 9651: true, 10626: true, 10696: true, 10761: true, 10797: true, 10847: true,
	11603: true, 11635: true, 11669: true, 11684: true, 11744: true, 11809: true, 11811: true, 11815: true, 11817: true, 11819: true,
	11832: true, 11902: true, 11920: true, 12531: true, 12582: true, 12583: true, 12590: true, 12592: true, 12709: true, 12777: true,
	12790: true, 12791: true, 12792: true, 12794: true, 12795: true, 12797: true, 12798: true, 12969: true, 13035: true, 13060: true,
	13148: true, 13183: true, 13204: true, 13209: true, 13213: true, 13218: true, 13246: true, 13286: true, 13361: true, 13401: true,
	13505: true, 13983: true, 13984: true, 14024: true, 14487: true, 14531: true, 14541: true, 14554: true, 14555: true, 14576: true,
	15814: true, 16004: true, 16403: true, 16463: true, 16484: true, 16530: true, 16548: true, 16571: true, 17054: true, 17066: true,
	17068: true, 17071: true, 17074: true, 17075: true, 17076: true, 17111: true, 17112: true, 17182: true, 17193: true, 17223: true,
	17705: true, 17752: true, 17774: true, 17780: true, 18168: true, 18202: true, 18203: true, 18310: true, 18326: true, 18348: true,
	18671: true, 18815: true, 18816: true, 18820: true, 19019: true, 19099: true, 19100: true, 19169: true, 19170: true, 19287: true,
	19288: true, 19289: true, 19324: true, 19334: true, 19337: true, 19339: true, 19340: true, 19341: true, 19342: true, 19344: true,
	19353: true, 19577: true, 19601: true, 19812: true, 19874: true, 19901: true, 19918: true, 19946: true, 19947: true, 19948: true,
	19949: true, 19950: true, 19951: true, 19953: true, 19954: true, 19956: true, 19957: true, 19959: true, 19961: true, 19962: true,
	19963: true, 19991: true, 19992: true, 20130: true, 20512: true, 20578: true, 21180: true, 21190: true, 21473: true, 21625: true,
	21670: true, 21679: true, 22268: true, 22321: true, 22395: true, 22397: true, 22678: true, 22691: true, 22862: true, 22954: true,
	23027: true, 23040: true, 23041: true, 23046: true, 23078: true, 23081: true, 23082: true, 23084: true, 23085: true, 23087: true,
	23088: true, 23089: true, 23090: true, 23091: true, 23092: true, 23093: true, 23197: true, 23198: true, 23199: true, 23203: true,
	23206: true, 23207: true, 23221: true, 23279: true, 23570: true, 220606: true, 228176: true, 249441: true, 249442: true, 249469: true,
	249470: true, 272427: true, 272432: true, 272433: true, 272435: true, 279248: true,
}

func effectImplemented(itemID int) bool {
	return engineImplementedEffectItemIDs[itemID]
}
